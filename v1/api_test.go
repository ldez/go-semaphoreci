package v1

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

func setupTest() (*Client, *http.ServeMux, func()) {
	apiHandler := http.NewServeMux()
	server := httptest.NewServer(apiHandler)

	client := NewClient(nil)

	baseURL, _ := url.Parse(server.URL)
	client.BaseURL = baseURL

	return client, apiHandler, server.Close
}

func mustTimeParse(raw string) *time.Time {
	parse, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		panic(err)
	}
	return &parse
}
