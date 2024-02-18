package max

import "net/http"

type Client struct {
	apiKey      string
	apiSecret   string
	apiEndpoint string
	httpClient  *http.Client
}

const (
	baseAPImainURL = "https://max-api.maicoin.com"
)

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		apiEndpoint: baseAPImainURL,
		httpClient:  http.DefaultClient,
	}
}
