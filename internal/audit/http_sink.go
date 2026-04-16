package audit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/metrics"
)

type HTTPSink struct {
	client    *http.Client
	url       string
	token     string
	clusterID string
}

func NewHTTPSink(baseURL string, timeout time.Duration) *HTTPSink {
	return NewHTTPSinkWithConfig(baseURL, timeout, "", "")
}

func NewHTTPSinkWithConfig(baseURL string, timeout time.Duration, token string, clusterID string) *HTTPSink {
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	return &HTTPSink{
		client:    &http.Client{Timeout: timeout},
		url:       baseURL + "/v1/ingest",
		token:     strings.TrimSpace(token),
		clusterID: strings.TrimSpace(clusterID),
	}
}

func (s *HTTPSink) Write(ctx context.Context, event Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		metrics.IncAuditForwardingFailure(event.Component)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.url, bytes.NewReader(payload))
	if err != nil {
		metrics.IncAuditForwardingFailure(event.Component)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if event.RequestID != "" {
		req.Header.Set("X-Request-Id", event.RequestID)
	}
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}
	if s.clusterID != "" {
		req.Header.Set("X-Changelock-Cluster-Id", s.clusterID)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		metrics.IncAuditForwardingFailure(event.Component)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		metrics.IncAuditForwardingFailure(event.Component)
		return fmt.Errorf("audit writer returned status %d", resp.StatusCode)
	}

	return nil
}
