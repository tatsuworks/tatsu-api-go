package tatsu_api

type modifyGuildMemberScoreReq struct {
	Action Action `json:"action"`
	Amount uint32 `json:"amount"`
}

type GuildMemberPoints struct {
	GuildID string `json:"guild_id"`
	Points  uint64 `json:"points"`
	Rank    int64  `json:"rank"`
	UserID  string `json:"user_id"`
}

type GuildMemberScore struct {
	GuildID string `json:"guild_id"`
	Score   uint64 `json:"score"`
	UserID  string `json:"user_id"`
}
