package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

type HttpResponse struct {
	Headers    http.Header
	Body       []byte
	StatusCode int
}

type HttpClientConfig struct {
	URL     string
	Headers map[string]string
}

type HttpClienter interface {
	Post(ctx context.Context, config *HttpClientConfig, payload []byte) (*HttpResponse, error)
	Get(ctx context.Context, config *HttpClientConfig) (*HttpResponse, error)
}

type HttpClient struct {
	client http.Client
}

func NewClient(http http.Client) *HttpClient {
	return &HttpClient{
		client: http,
	}
}

func (httpClient *HttpClient) Get(ctx context.Context, config *HttpClientConfig) (*HttpResponse, error) {
	return httpClient.execute(ctx, http.MethodGet, config, nil)
}

func (httpClient *HttpClient) Post(ctx context.Context, config *HttpClientConfig, body []byte) (*HttpResponse, error) {
	return httpClient.execute(ctx, http.MethodPost, config, body)
}

func (httpClient *HttpClient) execute(ctx context.Context, method string, config *HttpClientConfig, body []byte) (*HttpResponse, error) {
	var payload io.Reader
	if body != nil {
		payload = bytes.NewBuffer(body)

	}

	req, err := http.NewRequestWithContext(ctx, method, config.URL, payload)
	if err != nil {
		return nil, err
	}

	for key, value := range config.Headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, _ := io.ReadAll(resp.Body)

	return &HttpResponse{
		Body:       responseBody,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header.Clone(),
	}, nil

}
