package tatsu_api

import (
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

func (rc *restClient) getAllTimeGuildMemberRanking(ctx context.Context, guildID string,
	userID string) (*GuildMemberRanking, error) {
	// Ensure IDs are not empty.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	if userID == "" {
		return nil, xerrors.New("user id was empty")
	}

	var ranking *GuildMemberRanking

	// Make request.
	err := rc.get(ctx, getGuildMemberRanking(rankingPeriodAllTime, guildID, userID), &ranking)

	return ranking, err
}

func (rc *restClient) getAllTimeGuildRankings(ctx context.Context, guildID string, offset uint64) (*GuildRankings, error) {
	// Ensure guild ID is not empty.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	var rankings *GuildRankings

	// Make request.
	err := rc.get(ctx, getGuildRankings(rankingPeriodAllTime, guildID, offset), &rankings)

	return rankings, err
}

func (rc *restClient) getCurrentMonthGuildMemberRanking(ctx context.Context, guildID string,
	userID string) (*GuildMemberRanking, error) {
	// Ensure IDs are not empty.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	if userID == "" {
		return nil, xerrors.New("user id was empty")
	}

	var ranking *GuildMemberRanking

	// Make request.
	err := rc.get(ctx, getGuildMemberRanking(rankingPeriodMonth, guildID, userID), &ranking)

	return ranking, err
}

func (rc *restClient) getCurrentMonthGuildRankings(ctx context.Context, guildID string, offset uint64) (*GuildRankings, error) {
	// Ensure guild ID is not empty.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	var rankings *GuildRankings

	// Make request.
	err := rc.get(ctx, getGuildRankings(rankingPeriodMonth, guildID, offset), &rankings)

	return rankings, err
}

func (rc *restClient) getCurrentWeekGuildMemberRanking(ctx context.Context, guildID string,
	userID string) (*GuildMemberRanking, error) {
	// Ensure IDs are not empty.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	if userID == "" {
		return nil, xerrors.New("user id was empty")
	}

	var ranking *GuildMemberRanking

	// Make request.
	err := rc.get(ctx, getGuildMemberRanking(rankingPerioidWeek, guildID, userID), &ranking)

	return ranking, err
}

func (rc *restClient) getCurrentWeekGuildRankings(ctx context.Context, guildID string, offset uint64) (*GuildRankings, error) {
	// Ensure guild ID is not empty.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	var rankings *GuildRankings

	// Make request.
	err := rc.get(ctx, getGuildRankings(rankingPerioidWeek, guildID, offset), &rankings)

	return rankings, err
}

func (rc *restClient) getUserProfile(ctx context.Context, userID string) (*User, error) {
	// Ensure user ID is not empty.
	if userID == "" {
		return nil, xerrors.New("user id was empty")
	}

	var user *User

	// Make request.
	err := rc.get(ctx, getUserProfile(userID), &user)

	return user, err
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

func (rc *restClient) setHeaders(r *http.Request) {
	r.Header.Set("Authorization", rc.apikey)
	r.Header.Set("User-Agent", "Tatsu API Go/"+version+" (Hassie, https://github.com/tatsuworks/tatsu-api-go)")
}
