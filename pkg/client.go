package lastfm

import "fmt"

const (
	baseURL = "http://ws.audioscrobbler.com/2.0/?"
)

// Client ...
type Client struct {
	apiKey string
}

// NewClient ...
func NewClient(apiKey string) Client {
	return Client{apiKey: apiKey}
}

func generateURL(method string, query string, key string) string {
	urlFormat := "%smethod=%s&%s&api_key=%s&format=json"
	return fmt.Sprintf(urlFormat, baseURL, method, query, key)
}
