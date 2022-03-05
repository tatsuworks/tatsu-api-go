package tatsu_api

type ApiError struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
