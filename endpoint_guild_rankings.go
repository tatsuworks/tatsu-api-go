package tatsu_api

import "fmt"

type rankingPeriod string

const (
	rankingPeriodAllTime rankingPeriod = "all"
	rankingPeriodMonth   rankingPeriod = "month"
	rankingPeriodWeek    rankingPeriod = "week"
)

func getGuildMemberRanking(period rankingPeriod, guildID string, userID string) string {
	return baseURL + fmt.Sprintf("/guilds/%s/rankings/members/%s/%s", guildID, userID, string(period))
}

func getGuildRankings(period rankingPeriod, guildID string, offset uint64) string {
	return baseURL + fmt.Sprintf("/guilds/%s/rankings/%s?offset=%v", guildID, string(period), offset)
}
