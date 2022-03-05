package tatsu_api

import (
	"context"
	"golang.org/x/xerrors"
)

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
	err := rc.get(ctx, getGuildMemberRanking(rankingPeriodWeek, guildID, userID), &ranking)

	return ranking, err
}

func (rc *restClient) getCurrentWeekGuildRankings(ctx context.Context, guildID string, offset uint64) (*GuildRankings, error) {
	// Ensure guild ID is not empty.
	if guildID == "" {
		return nil, xerrors.New("guild id was empty")
	}

	var rankings *GuildRankings

	// Make request.
	err := rc.get(ctx, getGuildRankings(rankingPeriodWeek, guildID, offset), &rankings)

	return rankings, err
}
