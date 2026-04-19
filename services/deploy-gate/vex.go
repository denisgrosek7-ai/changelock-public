package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

const (
	vexDeployModeDisabled = "disabled"
	vexDeployModeEnforce  = "enforce"
)

type vulnerabilityNetEvaluator interface {
	Enabled() bool
	Mode() string
	NetVulnerabilities(ctx context.Context, tenant, environment, imageDigest, severityThreshold string) (audit.VulnerabilityNetResponse, error)
}

type httpVulnerabilityNetEvaluator struct {
	mode      string
	baseURL   string
	token     string
	client    *http.Client
}

func newVulnerabilityNetEvaluator() vulnerabilityNetEvaluator {
	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_VEX_DEPLOY_MODE"), vexDeployModeDisabled)))
	baseURL := strings.TrimRight(strings.TrimSpace(firstNonEmpty(
		os.Getenv("CHANGELOCK_VEX_URL"),
		os.Getenv("AUDIT_WRITER_URL"),
		os.Getenv("CHANGELOCK_AUDIT_WRITER_URL"),
	)), "/")
	return &httpVulnerabilityNetEvaluator{
		mode:    mode,
		baseURL: baseURL,
		token:   strings.TrimSpace(os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")),
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func validateVulnerabilityNetEvaluatorConfig() error {
	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_VEX_DEPLOY_MODE"), vexDeployModeDisabled)))
	switch mode {
	case vexDeployModeDisabled, vexDeployModeEnforce:
	default:
		return fmt.Errorf("unsupported CHANGELOCK_VEX_DEPLOY_MODE: %s", mode)
	}
	if mode == vexDeployModeDisabled {
		return nil
	}
	baseURL := strings.TrimSpace(firstNonEmpty(
		os.Getenv("CHANGELOCK_VEX_URL"),
		os.Getenv("AUDIT_WRITER_URL"),
		os.Getenv("CHANGELOCK_AUDIT_WRITER_URL"),
	))
	if baseURL == "" {
		return errors.New("CHANGELOCK_VEX_DEPLOY_MODE=enforce requires CHANGELOCK_VEX_URL or AUDIT_WRITER_URL")
	}
	if strings.EqualFold(strings.TrimSpace(os.Getenv("CHANGELOCK_AUTH_MODE")), auth.ModeStaticToken) &&
		strings.TrimSpace(os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")) == "" {
		return errors.New("CHANGELOCK_INTERNAL_SERVICE_TOKEN is required when CHANGELOCK_VEX_DEPLOY_MODE=enforce and CHANGELOCK_AUTH_MODE=static-token")
	}
	return nil
}

func (e *httpVulnerabilityNetEvaluator) Enabled() bool {
	return e != nil && e.mode == vexDeployModeEnforce && e.baseURL != ""
}

func (e *httpVulnerabilityNetEvaluator) Mode() string {
	if e == nil || e.mode == "" {
		return vexDeployModeDisabled
	}
	return e.mode
}

func (e *httpVulnerabilityNetEvaluator) NetVulnerabilities(ctx context.Context, tenant, environment, imageDigest, severityThreshold string) (audit.VulnerabilityNetResponse, error) {
	if !e.Enabled() {
		return audit.VulnerabilityNetResponse{}, errors.New("vex deploy evaluation is disabled")
	}
	query := url.Values{}
	query.Set("image_digest", strings.TrimSpace(imageDigest))
	query.Set("limit", "100")
	if strings.TrimSpace(tenant) != "" {
		query.Set("tenant_id", strings.TrimSpace(tenant))
	}
	if strings.TrimSpace(environment) != "" {
		query.Set("environment", strings.TrimSpace(environment))
	}
	if strings.TrimSpace(severityThreshold) != "" {
		query.Set("severity_threshold", strings.TrimSpace(severityThreshold))
	}
	endpoint := e.baseURL + "/v1/vulnerabilities/net?" + query.Encode()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	request.Header.Set("Accept", "application/json")
	if e.token != "" {
		request.Header.Set("Authorization", "Bearer "+e.token)
	}
	response, err := e.client.Do(request)
	if err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return audit.VulnerabilityNetResponse{}, fmt.Errorf("vex net vulnerability lookup returned %s", response.Status)
	}
	var result audit.VulnerabilityNetResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	return result, nil
}
