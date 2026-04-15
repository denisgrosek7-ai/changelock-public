package audit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

	if resp.StatusCode >= http.StatusBadRequest {
		var failure struct {
			Error  string `json:"error"`
			Reason string `json:"reason"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&failure); err == nil {
			message := strings.TrimSpace(FirstNonEmpty(failure.Error, failure.Reason))
			if message != "" {
				return ExceptionValidationResult{}, fmt.Errorf("exception validate request failed with status %d: %s", resp.StatusCode, message)
			}
		}
		return ExceptionValidationResult{}, fmt.Errorf("exception validate request failed with status %d", resp.StatusCode)
	}

	var result ExceptionValidationResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ExceptionValidationResult{}, err
	}
	return result, nil
}
