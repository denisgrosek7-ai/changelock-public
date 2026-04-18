package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/signingidentity"
)

type signerIdentityEvaluator interface {
	Enabled() bool
	Mode() string
	Evaluate(ctx context.Context, request signingIdentityEvaluateRequest) (signingidentity.Decision, error)
}

type signingIdentityEvaluateRequest struct {
	Issuer            string     `json:"issuer"`
	SignerIdentity    string     `json:"signer_identity"`
	Subject           string     `json:"subject,omitempty"`
	Repository        string     `json:"repository,omitempty"`
	Workflow          string     `json:"workflow,omitempty"`
	Ref               string     `json:"ref,omitempty"`
	TenantID          string     `json:"tenant_id,omitempty"`
	ClusterID         string     `json:"cluster_id,omitempty"`
	Environment       string     `json:"environment,omitempty"`
	TransparencyState string     `json:"transparency_state,omitempty"`
	EvidenceAt        *time.Time `json:"evidence_at,omitempty"`
}

type httpSignerIdentityEvaluator struct {
	mode    string
	baseURL string
	token   string
	client  *http.Client
}

func newSignerIdentityEvaluator() signerIdentityEvaluator {
	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT"), signingidentity.EnforcementDisabled)))
	baseURL := strings.TrimRight(strings.TrimSpace(firstNonEmpty(
		os.Getenv("CHANGELOCK_SIGNER_IDENTITY_URL"),
		os.Getenv("AUDIT_WRITER_URL"),
		os.Getenv("CHANGELOCK_AUDIT_WRITER_URL"),
	)), "/")
	return &httpSignerIdentityEvaluator{
		mode:    mode,
		baseURL: baseURL,
		token:   strings.TrimSpace(os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")),
		client:  &http.Client{Timeout: 3 * time.Second},
	}
}

func validateSignerIdentityEvaluatorConfig() error {
	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT"), signingidentity.EnforcementDisabled)))
	switch mode {
	case signingidentity.EnforcementDisabled, signingidentity.EnforcementMonitor, signingidentity.EnforcementEnforce:
	default:
		return fmt.Errorf("unsupported CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT: %s", mode)
	}
	if mode == signingidentity.EnforcementDisabled {
		return nil
	}
	baseURL := strings.TrimSpace(firstNonEmpty(
		os.Getenv("CHANGELOCK_SIGNER_IDENTITY_URL"),
		os.Getenv("AUDIT_WRITER_URL"),
		os.Getenv("CHANGELOCK_AUDIT_WRITER_URL"),
	))
	if baseURL == "" {
		return errors.New("CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT requires CHANGELOCK_SIGNER_IDENTITY_URL or AUDIT_WRITER_URL")
	}
	if strings.EqualFold(strings.TrimSpace(os.Getenv("CHANGELOCK_AUTH_MODE")), auth.ModeStaticToken) &&
		strings.TrimSpace(os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")) == "" {
		return errors.New("CHANGELOCK_INTERNAL_SERVICE_TOKEN is required when CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT is enabled and CHANGELOCK_AUTH_MODE=static-token")
	}
	return nil
}

func (e *httpSignerIdentityEvaluator) Enabled() bool {
	return e != nil && e.mode != signingidentity.EnforcementDisabled && e.baseURL != ""
}

func (e *httpSignerIdentityEvaluator) Mode() string {
	if e == nil || e.mode == "" {
		return signingidentity.EnforcementDisabled
	}
	return e.mode
}

func (e *httpSignerIdentityEvaluator) Evaluate(ctx context.Context, request signingIdentityEvaluateRequest) (signingidentity.Decision, error) {
	if !e.Enabled() {
		return signingidentity.Decision{}, errors.New("signing identity enforcement is disabled")
	}
	endpoint := e.baseURL + "/v1/signing-identities/evaluate"
	body, err := json.Marshal(request)
	if err != nil {
		return signingidentity.Decision{}, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(string(body)))
	if err != nil {
		return signingidentity.Decision{}, err
	}
	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("Content-Type", "application/json")
	if e.token != "" {
		httpRequest.Header.Set("Authorization", "Bearer "+e.token)
	}
	response, err := e.client.Do(httpRequest)
	if err != nil {
		return signingidentity.Decision{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return signingidentity.Decision{}, fmt.Errorf("signing identity evaluation returned %s", response.Status)
	}
	var decision signingidentity.Decision
	if err := json.NewDecoder(response.Body).Decode(&decision); err != nil {
		return signingidentity.Decision{}, err
	}
	return decision, nil
}
