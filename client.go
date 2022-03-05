package tatsu_api

type Client struct {
	rc *restClient
}

func New(apikey string) *Client {
	return &Client{
		rc: newRestClient(apikey),
	}
}
