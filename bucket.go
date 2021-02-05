package tatsu_api

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

type bucket struct {
	max           uint8
	remaining     uint8
	resetInterval time.Duration
	reset         time.Time
	sync.Mutex
}

func newBucket() *bucket {
	return &bucket{
		max:           60,
		remaining:     60,
		resetInterval: 1 * time.Minute,
		reset:         time.Now().Add(1 * time.Minute),
	}
}

func (b *bucket) acquire() {
	b.Lock()
	defer b.Unlock()

	// Check if bucket needs to be refilled.
	if !time.Now().Before(b.reset) {
		b.refill()
	}

	if b.remaining > 0 {
		b.remaining--
		return
	}

	// Sleep for the time difference.
	time.Sleep(b.reset.Sub(time.Now()))

	b.refill()
	b.remaining--
}

func (b *bucket) refill() {
	b.remaining = b.max
	b.reset = time.Now().Add(b.resetInterval)
}

func (b *bucket) parseHeaders(headers http.Header) {
	if headers.Get("X-RateLimit-Remaining") == "" {
		return
	}

	b.Lock()
	defer b.Unlock()

	// Parse remaining.
	remaining, err := strconv.ParseInt(headers.Get("X-RateLimit-Remaining"), 10, 8)
	if err != nil {
		return
	}

	// Parse reset.
	reset, err := strconv.ParseInt(headers.Get("X-RateLimit-Reset"), 10, 64)
	if err != nil {
		return
	}

	b.remaining = uint8(remaining)
	b.reset = time.Unix(reset, 0)
}
