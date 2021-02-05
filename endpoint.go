package tatsu_api

import "fmt"

const baseURL = "https://api.tatsu.gg/v1"

func getAllTimeGuildMemberRanking(guildID string, userID string) string {
	return baseURL + fmt.Sprintf("/guilds/%s/rankings/members/%s/all", guildID, userID)
}

func getAllTimeGuildRankings(guildID string, offset uint64) string {
	return baseURL + fmt.Sprintf("/guilds/%s/rankings/all?offset=%v", guildID, offset)
}

func getUserProfile(userID string) string {
	return baseURL + fmt.Sprintf("/users/%s/profile", userID)
}
