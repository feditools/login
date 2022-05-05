package http

import (
	"context"
	"io"
	"net/http"
)

// NewRequest calls http.NewRequest with expected http User-Agent
func NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", GetUserAgent())
	return req, nil
}

// Get calls http.Get with expected http User-Agent
func Get(ctx context.Context, url string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, err := NewRequest(ctx, http.MethodGet, url, nil)
	return client.Do(req)
}
