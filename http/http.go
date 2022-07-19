package http

import (
	"net/http"
)

//go:generate mockgen -source=./http.go -destination=../mock/http_mock.go http Http
type Http interface {
	Get(url string) (*http.Response, error)
}

type CustomClient struct {
	client *http.Client
}

func NewCustomClient() Http {
	return &CustomClient{client: http.DefaultClient}
}

func (c *CustomClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
