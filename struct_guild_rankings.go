package tatsu_api

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
