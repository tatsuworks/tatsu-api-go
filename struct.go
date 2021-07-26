package tatsu_api

import (
	"golang.org/x/xerrors"
	"time"
)

type GuildMemberRanking struct {
	GuildID string `json:"guild_id"`
	Rank    int64  `json:"rank"`
	Score   uint64 `json:"score"`
	UserID  string `json:"user_id"`
}

type GuildRankings struct {
	GuildID  string          `json:"guild_id"`
	Rankings []*GuildRanking `json:"rankings"`
}

type GuildRanking struct {
	Rank   int64  `json:"rank"`
	Score  uint64 `json:"score"`
	UserID string `json:"user_id"`
}

type Timestamp string

func (t Timestamp) Parse() (time.Time, error) {
	timestamp, err := time.Parse(time.RFC3339, string(t))
	if err != nil {
		return time.Time{}, xerrors.Errorf("tatsu_api: parse subscription renewal: %w", err)
	}

	return timestamp, nil
}

type User struct {
	AvatarURL           string    `json:"avatar_url"`
	Credits             uint64    `json:"credits"`
	Discriminator       string    `json:"discriminator"`
	ID                  string    `json:"id"`
	InfoBox             string    `json:"info_box"`
	Reputation          uint64    `json:"reputation"`
	SubscriptionRenewal Timestamp `json:"subscription_renewal"`
	SubscriptionType    uint8     `json:"subscription_type"`
	Title               string    `json:"title"`
	Tokens              uint64    `json:"tokens"`
	Username            string    `json:"username"`
	XP                  uint64    `json:"xp"`
}
