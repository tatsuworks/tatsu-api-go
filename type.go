package tatsu_api

import (
	"time"

	"golang.org/x/xerrors"
)

type Timestamp string

func (t Timestamp) Parse() (time.Time, error) {
	timestamp, err := time.Parse(time.RFC3339, string(t))
	if err != nil {
		return time.Time{}, xerrors.Errorf("tatsu_api: parse subscription renewal: %w", err)
	}

	return timestamp, nil
}
