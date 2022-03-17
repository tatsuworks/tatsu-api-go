package tatsu_api

import (
	"context"

	"golang.org/x/xerrors"
)

func (rc *restClient) getGuildMemberPoints(ctx context.Context, guildID string, userID string) (*GuildMemberPoints, error) {
	// Validate input.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}
	if userID == "" {
		return nil, xerrors.New("user id was empty")
	}

	var points *GuildMemberPoints
	err := rc.get(ctx, getGuildMemberPoints(guildID, userID), &points)

	return points, err
}

func (rc *restClient) modifyGuildMemberScore(ctx context.Context, guildID string, userID string, action Action,
	amount uint32) (*GuildMemberScore, error) {
	// Validate input.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	if userID == "" {
		return nil, xerrors.New("user id was empty")
	}

	if amount < 1 || amount > 100000 {
		return nil, xerrors.New("the amount to modify must be between 1 and 100,000 inclusive")
	}

	// Make request.
	var score *GuildMemberScore
	err := rc.patch(ctx, modifyGuildMemberScore(guildID, userID), &modifyGuildMemberScoreReq{
		Action: action,
		Amount: amount,
	}, &score)

	return score, err
}
