package tatsu_api

import "fmt"

func getUserProfile(userID string) string {
	return baseURL + fmt.Sprintf("/users/%s/profile", userID)
}
