package tatsu_api

import (
	"context"

	"golang.org/x/xerrors"
)

// ModifyGuildMemberScore wraps around ModifyGuildMemberScore using the background context.
func (c *Client) ModifyGuildMemberScore(guildID string, userID string, action Action,
	amount uint32) (*GuildMemberScore, error) {
	return c.ModifyGuildMemberScoreWithContext(context.Background(), guildID, userID, action, amount)
}

// ModifyGuildMemberScoreWithContext modifies a guild member's score using the specified context.
func (c *Client) ModifyGuildMemberScoreWithContext(ctx context.Context, guildID string, userID string,
	action Action, amount uint32) (*GuildMemberScore, error) {
	score, err := c.rc.modifyGuildMemberScore(ctx, guildID, userID, action, amount)
	if err != nil {
		return nil, xerrors.Errorf("failed to modify guild member score: %w", err)
	}

	return score, err
}
