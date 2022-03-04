package tatsu_api

import (
	"context"

	"golang.org/x/xerrors"
)

type Client struct {
	rc *restClient
}

func New(apikey string) *Client {
	return &Client{
		rc: newRestClient(apikey),
	}
}

// GetAllTimeGuildMemberRanking wraps around GetAllTimeGuildMemberRankingWithContext using the background context.
func (c *Client) GetAllTimeGuildMemberRanking(guildID string, userID string) (*GuildMemberRanking, error) {
	return c.GetAllTimeGuildMemberRankingWithContext(context.Background(), guildID, userID)
}

// GetAllTimeGuildMemberRankingWithContext gets a guild member's all time ranking using the specified context.
func (c *Client) GetAllTimeGuildMemberRankingWithContext(ctx context.Context, guildID string,
	userID string) (*GuildMemberRanking, error) {
	ranking, err := c.rc.getAllTimeGuildMemberRanking(ctx, guildID, userID)
	if err != nil {
		return nil, xerrors.Errorf("failed to get all time guild member ranking: %w", err)
	}

	return ranking, err
}

// GetCurrentMonthGuildMemberRanking wraps around GetCurrentMonthGuildMemberRankingWithContext using the background context.
func (c *Client) GetCurrentMonthGuildMemberRanking(guildID string, userID string) (*GuildMemberRanking, error) {
	return c.GetCurrentMonthGuildMemberRankingWithContext(context.Background(), guildID, userID)
}

// GetCurrentMonthGuildMemberRankingWithContext gets a guild member's current month ranking using the specified context.
func (c *Client) GetCurrentMonthGuildMemberRankingWithContext(ctx context.Context, guildID string,
	userID string) (*GuildMemberRanking, error) {
	ranking, err := c.rc.getCurrentMonthGuildMemberRanking(ctx, guildID, userID)
	if err != nil {
		return nil, xerrors.Errorf("failed to get current month guild member ranking: %w", err)
	}

	return ranking, err
}

// GetCurrentWeekGuildMemberRanking wraps around GetCurrentWeekGuildMemberRankingWithContext using the background context.
func (c *Client) GetCurrentWeekGuildMemberRanking(guildID string, userID string) (*GuildMemberRanking, error) {
	return c.GetCurrentWeekGuildMemberRankingWithContext(context.Background(), guildID, userID)
}

// GetCurrentWeekGuildMemberRankingWithContext gets a guild member's current week ranking using the specified context.
func (c *Client) GetCurrentWeekGuildMemberRankingWithContext(ctx context.Context, guildID string,
	userID string) (*GuildMemberRanking, error) {
	ranking, err := c.rc.getCurrentWeekGuildMemberRanking(ctx, guildID, userID)
	if err != nil {
		return nil, xerrors.Errorf("failed to get current week guild member ranking: %w", err)
	}

	return ranking, err
}

// GetAllTimeGuildRankings wraps around GetAllTimeGuildRankingsWithContext using a background context.
func (c *Client) GetAllTimeGuildRankings(guildID string, offset uint64) (*GuildRankings, error) {
	return c.GetAllTimeGuildRankingsWithContext(context.Background(), guildID, offset)
}

// GetAllTimeGuildRankingsWithContext gets a guild's all time rankings using the specified context.
func (c *Client) GetAllTimeGuildRankingsWithContext(ctx context.Context, guildID string,
	offset uint64) (*GuildRankings, error) {
	rankings, err := c.rc.getAllTimeGuildRankings(ctx, guildID, offset)
	if err != nil {
		return nil, xerrors.Errorf("failed to get all time guild rankings: %w", err)
	}

	return rankings, err
}

// GetCurrentMonthGuildRankings wraps around GetCurrentMonthGuildRankingsWithContext using a background context.
func (c *Client) GetCurrentMonthGuildRankings(guildID string, offset uint64) (*GuildRankings, error) {
	return c.GetCurrentMonthGuildRankingsWithContext(context.Background(), guildID, offset)
}

// GetCurrentMonthGuildRankingsWithContext gets a guild's current month rankings using the specified context.
func (c *Client) GetCurrentMonthGuildRankingsWithContext(ctx context.Context, guildID string,
	offset uint64) (*GuildRankings, error) {
	rankings, err := c.rc.getCurrentMonthGuildRankings(ctx, guildID, offset)
	if err != nil {
		return nil, xerrors.Errorf("failed to get current month guild rankings: %w", err)
	}

	return rankings, err
}

// GetCurrentWeekGuildRankings wraps around GetCurrentWeekGuildRankingsWithContext using a background context.
func (c *Client) GetCurrentWeekGuildRankings(guildID string, offset uint64) (*GuildRankings, error) {
	return c.GetCurrentWeekGuildRankingsWithContext(context.Background(), guildID, offset)
}

// GetCurrentWeekGuildRankingsWithContext gets a guild's current week rankings using the specified context.
func (c *Client) GetCurrentWeekGuildRankingsWithContext(ctx context.Context, guildID string,
	offset uint64) (*GuildRankings, error) {
	rankings, err := c.rc.getCurrentWeekGuildRankings(ctx, guildID, offset)
	if err != nil {
		return nil, xerrors.Errorf("failed to get current week guild rankings: %w", err)
	}

	return rankings, err
}

// GetUserProfile wraps around GetUserProfileWithContext using the background context.
func (c *Client) GetUserProfile(userID string) (*User, error) {
	return c.GetUserProfileWithContext(context.Background(), userID)
}

// GetUserProfileWithContext gets a user's profile using the specified context.
func (c *Client) GetUserProfileWithContext(ctx context.Context, userID string) (*User, error) {
	user, err := c.rc.getUserProfile(ctx, userID)
	if err != nil {
		return nil, xerrors.Errorf("failed to get user profile: %w", err)
	}

	return user, nil
}
