package osugo

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// OsuClient acts as a middleman between the client making the requests and the API server.
type OsuClient struct {
	apiKey string
	client *http.Client
}

// InitClient creates a new OsuClient to make requests with.
func InitClient(key string) *OsuClient {
	c := OsuClient{
		apiKey: key,
		client: &http.Client{},
	}
	return &c
}

func (c OsuClient) sendRequest(endpoint string, q query) ([]byte, error) {
	key, err := q.constructQuery(c.apiKey)
	if err != nil {
		return nil, err
	}

	url := []string{"https://osu.ppy.sh/api/", endpoint, "?", key}
	bytes, err := c.client.Get(strings.Join(url, ""))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(bytes.Body)
	if err != nil {
		return nil, err
	}

	defer bytes.Body.Close()
	return data, nil
}
