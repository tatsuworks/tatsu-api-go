package tatsu_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

type restClient struct {
	apikey     string
	bucket     *bucket
	httpClient *http.Client
}

func newRestClient(apikey string) *restClient {
	return &restClient{
		apikey:     apikey,
		bucket:     newBucket(),
		httpClient: http.DefaultClient,
	}
}

func (rc *restClient) get(ctx context.Context, endpoint string, v interface{}) error {
	// Get rate limit clearance.
	rc.bucket.acquire()

	// Create request.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return xerrors.Errorf("create request: %w", err)
	}

	// Set headers.
	rc.setHeaders(req)

	// Execute request.
	resp, err := rc.httpClient.Do(req)
	if err != nil {
		return xerrors.Errorf("execute request: %w", err)
	}

	// Defer closing body.
	defer resp.Body.Close()

	// Pass headers to bucket.
	rc.bucket.parseHeaders(resp.Header)

	// Return error if status is not OK.
	if resp.StatusCode != http.StatusOK {
		// Try parsing error body.
		var errorResp ApiError
		err = json.NewDecoder(resp.Body).Decode(&errorResp)
		if err != nil {
			// Return status as error.
			return xerrors.New(resp.Status)
		}

		return xerrors.New(fmt.Sprintf("%d: %s", errorResp.Code, errorResp.Message))
	}

	// Parse response.
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return xerrors.Errorf("parse response: %w", err)
	}

	return nil
}

func (rc *restClient) patch(ctx context.Context, endpoint string, r interface{}, v interface{}) error {
	// Acquire bucket token.
	rc.bucket.acquire()

	// Marshal request body.
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(r)
	if err != nil {
		return xerrors.Errorf("marshal request body: %w", err)
	}

	// Create request.
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpoint, buf)
	if err != nil {
		return xerrors.Errorf("create request: %w", err)
	}

	// Set headers.
	rc.setHeaders(req)

	// Execute request.
	resp, err := rc.httpClient.Do(req)
	if err != nil {
		return xerrors.Errorf("execute request: %w", err)
	}

	// Defer closing body.
	defer resp.Body.Close()

	// Pass headers to bucket.
	rc.bucket.parseHeaders(resp.Header)

	// Return error if status is not 200 OK or 204 No Content.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		// Try parsing error body.
		var errorResp ApiError
		err = json.NewDecoder(resp.Body).Decode(&errorResp)
		if err != nil {
			// Return status as error.
			return xerrors.New(resp.Status)
		}

		return xerrors.New(fmt.Sprintf("%d: %s", errorResp.Code, errorResp.Message))
	}

	// If status 204 No Content, return.
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	// Status is 200 OK, parse response.
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return xerrors.Errorf("parse response: %w", err)
	}

	return nil
}

func (rc *restClient) setHeaders(r *http.Request) {
	r.Header.Set("Authorization", rc.apikey)
	r.Header.Set("User-Agent", "Tatsu API Go/"+version+" (Hassie, https://github.com/tatsuworks/tatsu-api-go)")
}
