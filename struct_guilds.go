package tatsu_api

type modifyGuildMemberScoreReq struct {
	Action Action `json:"action"`
	Amount uint32 `json:"amount"`
}

type GuildMemberScore struct {
	GuildID string `json:"guild_id"`
	Score   uint64 `json:"score"`
	UserID  string `json:"user_id"`
}
