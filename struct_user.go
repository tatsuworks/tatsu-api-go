package tatsu_api

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
