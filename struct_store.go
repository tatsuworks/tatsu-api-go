package tatsu_api

type StoreCurrency uint8

const (
	StoreCurrencyCredits StoreCurrency = iota
	StoreCurrencyTokens
	StoreCurrencyEmeralds
	StoreCurrencyCandyCane
	StoreCurrencyUSD
	StoreCurrencyCandyCorn
)

type StoreListing struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Summary     string        `json:"summary"`
	Description string        `json:"description"`
	New         bool          `json:"new"`
	Preview     string        `json:"preview"`
	Prices      []*StorePrice `json:"prices"`
	Categories  []string      `json:"categories"`
	Tags        []string      `json:"tags"`
}

type StorePrice struct {
	Currency StoreCurrency `json:"currency"`
	Amount   float32       `json:"amount"`
}
