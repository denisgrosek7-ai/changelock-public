package audit

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type HTTPExceptionClient struct {
	client *http.Client
	url    string
	token  string
}

func NewHTTPExceptionClient(baseURL string, timeout time.Duration, bearerToken ...string) *HTTPExceptionClient {
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")
	if baseURL == "" {
		return nil
	}
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	token := ""
	if len(bearerToken) > 0 {
		token = strings.TrimSpace(bearerToken[0])
	}
	return &HTTPExceptionClient{
		client: &http.Client{Timeout: timeout},
		url:    baseURL + "/v1/exceptions/validate",
		token:  token,
	}
}

func (c *HTTPExceptionClient) Validate(ctx context.Context, request ExceptionValidationRequest) (ExceptionValidationResult, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return ExceptionValidationResult{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewReader(payload))
	if err != nil {
		return ExceptionValidationResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return ExceptionValidationResult{}, err
	}
	defer resp.Body.Close()

	var result ExceptionValidationResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ExceptionValidationResult{}, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		if result.Reason != "" {
			return ExceptionValidationResult{}, ErrInvalidException
		}
		return ExceptionValidationResult{}, ErrInvalidException
	}
	return result, nil
}
