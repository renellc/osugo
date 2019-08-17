package osugo

import "net/http"

type OsuClient struct {
	apiKey string
	client *http.Client
}
