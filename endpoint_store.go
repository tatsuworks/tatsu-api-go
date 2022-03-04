package tatsu_api

import "fmt"

func getStoreListing(listingID string) string {
	return baseURL + fmt.Sprintf("/store/listings/%s", listingID)
}
