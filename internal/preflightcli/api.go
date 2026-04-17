package preflightcli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
)

type APIClient struct {
	baseURL string
	token   string
	client  *http.Client
}

type AuthInfo struct {
	Authenticated bool   `json:"authenticated"`
	AuthMode      string `json:"auth_mode"`
	Subject       string `json:"subject,omitempty"`
	Role          string `json:"role,omitempty"`
	TenantID      string `json:"tenant_id,omitempty"`
	GlobalScope   bool   `json:"global_scope,omitempty"`
}

func NewAPIClient(config Config, client *http.Client) *APIClient {
	baseURL := strings.TrimRight(strings.TrimSpace(config.APIURL), "/")
	if config.Offline || baseURL == "" {
		return nil
	}
	if client == nil {
		client = &http.Client{Timeout: config.Timeout}
	}
	return &APIClient{
		baseURL: baseURL,
		token:   strings.TrimSpace(config.Token),
		client:  client,
	}
}

func (c *APIClient) AuthMe(ctx context.Context) (AuthInfo, error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/v1/auth/me", nil)
	if err != nil {
		return AuthInfo{}, err
	}
	var info AuthInfo
	if err := c.doJSON(req, &info); err != nil {
		return AuthInfo{}, err
	}
	return info, nil
}

func (c *APIClient) ListExceptions(ctx context.Context, filter audit.ExceptionFilter) ([]audit.PolicyException, error) {
	query := url.Values{}
	if filter.Status != "" {
		query.Set("status", filter.Status)
	}
	if filter.ExceptionType != "" {
		query.Set("exception_type", filter.ExceptionType)
	}
	if filter.TenantID != "" {
		query.Set("tenant_id", filter.TenantID)
	}
	if filter.Environment != "" {
		query.Set("environment", filter.Environment)
	}
	if filter.Namespace != "" {
		query.Set("namespace", filter.Namespace)
	}
	if filter.Repo != "" {
		query.Set("repo", filter.Repo)
	}
	if filter.ImageDigest != "" {
		query.Set("image_digest", filter.ImageDigest)
	}
	if filter.CVEID != "" {
		query.Set("cve_id", filter.CVEID)
	}
	if filter.Active != nil {
		query.Set("active", fmt.Sprintf("%t", *filter.Active))
	}
	if filter.Limit > 0 {
		query.Set("limit", fmt.Sprintf("%d", filter.Limit))
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/v1/exceptions?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Exceptions []audit.PolicyException `json:"exceptions"`
	}
	if err := c.doJSON(req, &response); err != nil {
		return nil, err
	}
	return response.Exceptions, nil
}

func (c *APIClient) VulnerabilityNet(ctx context.Context, imageDigest, tenantID, environment, severityThreshold string) (audit.VulnerabilityNetResponse, error) {
	query := url.Values{}
	if strings.TrimSpace(imageDigest) != "" {
		query.Set("image_digest", strings.TrimSpace(imageDigest))
	}
	if strings.TrimSpace(tenantID) != "" {
		query.Set("tenant_id", strings.TrimSpace(tenantID))
	}
	if strings.TrimSpace(environment) != "" {
		query.Set("environment", strings.TrimSpace(environment))
	}
	if strings.TrimSpace(severityThreshold) != "" {
		query.Set("severity_threshold", strings.TrimSpace(severityThreshold))
	}
	query.Set("limit", "100")

	req, err := c.newRequest(ctx, http.MethodGet, "/v1/vulnerabilities/net?"+query.Encode(), nil)
	if err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	var response audit.VulnerabilityNetResponse
	if err := c.doJSON(req, &response); err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	return response, nil
}

func (c *APIClient) newRequest(ctx context.Context, method, path string, body []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	return req, nil
}

func (c *APIClient) doJSON(req *http.Request, dst any) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var failure struct {
			Error string `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&failure); err == nil && strings.TrimSpace(failure.Error) != "" {
			return fmt.Errorf("api request failed with status %d: %s", resp.StatusCode, failure.Error)
		}
		return fmt.Errorf("api request failed with status %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(dst)
}

func apiTimeout(config Config) time.Duration {
	if config.Timeout <= 0 {
		return 2 * time.Minute
	}
	return config.Timeout
}
