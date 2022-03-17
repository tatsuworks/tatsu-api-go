package tatsu_api

import "fmt"

func getGuildMemberPoints(guildID string, userID string) string {
	return baseURL + fmt.Sprintf("/guilds/%s/members/%s/points", guildID, userID)
}

func modifyGuildMemberScore(guildID string, userID string) string {
	return baseURL + fmt.Sprintf("/guilds/%s/members/%s/score", guildID, userID)
}
