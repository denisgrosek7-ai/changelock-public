package signingidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	ProviderGitHubOIDC  = "github-oidc"
	ProviderGenericOIDC = "generic-oidc"

	EnforcementDisabled = "disabled"
	EnforcementMonitor  = "monitor"
	EnforcementEnforce  = "enforce"

	AuthorizationAuthorized   = "authorized"
	AuthorizationUnauthorized = "unauthorized"
	AuthorizationUnknown      = "unknown"

	ReasonAuthorized              = "signer_identity_authorized"
	ReasonEvidenceMissing         = "signer_identity_evidence_missing"
	ReasonPolicyMissing           = "signer_identity_policy_missing"
	ReasonIssuerMismatch          = "signer_identity_issuer_mismatch"
	ReasonSubjectMismatch         = "signer_identity_subject_mismatch"
	ReasonRepositoryMismatch      = "signer_identity_repository_mismatch"
	ReasonWorkflowMismatch        = "signer_identity_workflow_mismatch"
	ReasonRefMismatch             = "signer_identity_ref_mismatch"
	ReasonPolicyDisabled          = "signer_identity_policy_disabled"
	ReasonDistrustedAfterCutoff   = "signer_identity_distrusted_after_cutoff"
	ReasonDistrustTimeUnavailable = "signer_identity_distrust_time_unavailable"
	ReasonRekorRequired           = "signer_identity_rekor_required"
	ReasonTransparencyUnverified  = "signer_identity_transparency_unverified"
	ReasonProviderUnsupported     = "signer_identity_provider_unsupported"
	ReasonUnknown                 = "signer_identity_unknown"

	FindingUnauthorizedIdentity    = "unauthorized_identity_observed"
	FindingDistrustedIdentity      = "distrusted_identity_observed"
	FindingTransparencyMissing     = "transparency_evidence_missing"
	FindingWorkflowTokenNoPolicy   = "workflow_id_token_without_policy"
	FindingPolicyWorkflowMissing   = "policy_workflow_missing"
	FindingWorkflowUnexpectedScope = "workflow_scope_unlinked"
)

var (
	ErrInvalidConfig = errors.New("invalid signing identity config")
	ErrInvalidPolicy = errors.New("invalid signing identity policy")
)

type Config struct {
	Enforcement       string
	RequireRekor      bool
	QuarantineOnDrift bool
	WorkflowsDir      string
}

type Policy struct {
	ID              string     `json:"id"`
	Name            string     `json:"name,omitempty"`
	ProviderType    string     `json:"provider_type"`
	Issuer          string     `json:"issuer,omitempty"`
	SignerIdentity  string     `json:"signer_identity,omitempty"`
	Subject         string     `json:"subject,omitempty"`
	Repository      string     `json:"repository,omitempty"`
	Workflow        string     `json:"workflow,omitempty"`
	Ref             string     `json:"ref,omitempty"`
	TenantID        string     `json:"tenant_id,omitempty"`
	ClusterID       string     `json:"cluster_id,omitempty"`
	Environment     string     `json:"environment,omitempty"`
	Enabled         bool       `json:"enabled"`
	DistrustedAfter *time.Time `json:"distrusted_after,omitempty"`
	DistrustReason  string     `json:"distrust_reason,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	CreatedBy       string     `json:"created_by,omitempty"`
	UpdatedBy       string     `json:"updated_by,omitempty"`
}

type CreatePolicyRequest struct {
	Name           string `json:"name,omitempty"`
	ProviderType   string `json:"provider_type,omitempty"`
	Issuer         string `json:"issuer,omitempty"`
	SignerIdentity string `json:"signer_identity,omitempty"`
	Subject        string `json:"subject,omitempty"`
	Repository     string `json:"repository,omitempty"`
	Workflow       string `json:"workflow,omitempty"`
	Ref            string `json:"ref,omitempty"`
	TenantID       string `json:"tenant_id,omitempty"`
	ClusterID      string `json:"cluster_id,omitempty"`
	Environment    string `json:"environment,omitempty"`
	Enabled        *bool  `json:"enabled,omitempty"`
}

type DecisionInput struct {
	Issuer            string
	SignerIdentity    string
	Subject           string
	Repository        string
	Workflow          string
	Ref               string
	TenantID          string
	ClusterID         string
	Environment       string
	TransparencyState string
	EvidenceAt        *time.Time
}

type Decision struct {
	Authorized           string     `json:"authorized"`
	EnforcementMode      string     `json:"enforcement_mode"`
	MatchedPolicyID      string     `json:"matched_policy_id,omitempty"`
	MatchedPolicyName    string     `json:"matched_policy_name,omitempty"`
	PolicyMatched        bool       `json:"policy_matched"`
	ReasonCode           string     `json:"reason_code"`
	ReasonDetail         string     `json:"reason_detail,omitempty"`
	Deny                 bool       `json:"deny"`
	DistrustedAfter      *time.Time `json:"distrusted_after,omitempty"`
	TransparencyState    string     `json:"transparency_state,omitempty"`
	TransparencyRequired bool       `json:"transparency_required,omitempty"`
}

type Observation struct {
	ID                string     `json:"id"`
	ProviderType      string     `json:"provider_type,omitempty"`
	Issuer            string     `json:"issuer,omitempty"`
	SignerIdentity    string     `json:"signer_identity,omitempty"`
	Subject           string     `json:"subject,omitempty"`
	Repository        string     `json:"repository,omitempty"`
	Workflow          string     `json:"workflow,omitempty"`
	Ref               string     `json:"ref,omitempty"`
	CommitSHA         string     `json:"commit_sha,omitempty"`
	ImageDigest       string     `json:"image_digest,omitempty"`
	TenantID          string     `json:"tenant_id,omitempty"`
	ClusterID         string     `json:"cluster_id,omitempty"`
	Environment       string     `json:"environment,omitempty"`
	FirstSeenAt       *time.Time `json:"first_seen_at,omitempty"`
	LastSeenAt        *time.Time `json:"last_seen_at,omitempty"`
	EventCount        int        `json:"event_count"`
	ArtifactCount     int        `json:"artifact_count"`
	VerificationState string     `json:"verification_state,omitempty"`
	Authorized        string     `json:"authorized"`
	MatchedPolicyID   string     `json:"matched_policy_id,omitempty"`
	DistrustedAfter   *time.Time `json:"distrusted_after,omitempty"`
	ReasonCode        string     `json:"reason_code,omitempty"`
	ReasonDetail      string     `json:"reason_detail,omitempty"`
}

type Finding struct {
	ID            string     `json:"id"`
	Type          string     `json:"type"`
	Severity      string     `json:"severity"`
	Repository    string     `json:"repository,omitempty"`
	Workflow      string     `json:"workflow,omitempty"`
	Ref           string     `json:"ref,omitempty"`
	PolicyID      string     `json:"policy_id,omitempty"`
	ObservationID string     `json:"observation_id,omitempty"`
	Reason        string     `json:"reason"`
	DetectedAt    *time.Time `json:"detected_at,omitempty"`
	Advisory      bool       `json:"advisory"`
}

type StatusSummary struct {
	EnforcementMode       string         `json:"enforcement_mode"`
	RequireRekor          bool           `json:"require_rekor"`
	TotalPolicies         int            `json:"total_policies"`
	EnabledPolicies       int            `json:"enabled_policies"`
	ObservedIdentities    int            `json:"observed_identities"`
	Authorized            int            `json:"authorized"`
	Unauthorized          int            `json:"unauthorized"`
	Unknown               int            `json:"unknown"`
	Findings              int            `json:"findings"`
	WorkflowDriftFindings int            `json:"workflow_drift_findings"`
	CountsByReasonCode    map[string]int `json:"counts_by_reason_code"`
}

type WorkflowDocument struct {
	Path           string `json:"path"`
	SigningCapable bool   `json:"signing_capable"`
}

func ParseConfig(getenv func(string) string) (Config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}
	enforcement := strings.ToLower(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT"), EnforcementDisabled)))
	switch enforcement {
	case EnforcementDisabled, EnforcementMonitor, EnforcementEnforce:
	default:
		return Config{}, fmt.Errorf("%w: unsupported CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT %q", ErrInvalidConfig, enforcement)
	}

	requireRekor, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_SIGNER_IDENTITY_REQUIRE_REKOR"), "false"))
	if err != nil {
		return Config{}, fmt.Errorf("%w: invalid CHANGELOCK_SIGNER_IDENTITY_REQUIRE_REKOR: %v", ErrInvalidConfig, err)
	}
	quarantineOnDrift, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_SIGNER_IDENTITY_QUARANTINE_ON_DRIFT"), "false"))
	if err != nil {
		return Config{}, fmt.Errorf("%w: invalid CHANGELOCK_SIGNER_IDENTITY_QUARANTINE_ON_DRIFT: %v", ErrInvalidConfig, err)
	}

	return Config{
		Enforcement:       enforcement,
		RequireRekor:      requireRekor,
		QuarantineOnDrift: quarantineOnDrift,
		WorkflowsDir:      strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR"), ".github/workflows")),
	}, nil
}

func NormalizeCreateRequest(request CreatePolicyRequest) (CreatePolicyRequest, error) {
	request.Name = strings.TrimSpace(request.Name)
	request.ProviderType = normalizeProviderType(firstNonEmpty(request.ProviderType, ProviderGitHubOIDC))
	request.Issuer = strings.TrimSpace(request.Issuer)
	request.SignerIdentity = strings.TrimSpace(request.SignerIdentity)
	request.Subject = strings.TrimSpace(request.Subject)
	request.Repository = strings.TrimSpace(request.Repository)
	request.Workflow = normalizeWorkflowPath(request.Workflow)
	request.Ref = strings.TrimSpace(request.Ref)
	request.TenantID = strings.TrimSpace(request.TenantID)
	request.ClusterID = strings.TrimSpace(request.ClusterID)
	request.Environment = strings.TrimSpace(request.Environment)

	if !validProviderType(request.ProviderType) {
		return request, fmt.Errorf("%w: unsupported provider_type %q", ErrInvalidPolicy, request.ProviderType)
	}
	if request.Issuer == "" {
		return request, fmt.Errorf("%w: issuer is required", ErrInvalidPolicy)
	}
	if request.SignerIdentity == "" {
		return request, fmt.Errorf("%w: signer_identity is required", ErrInvalidPolicy)
	}
	if request.Repository == "" {
		return request, fmt.Errorf("%w: repository is required", ErrInvalidPolicy)
	}
	return request, nil
}

func NewPolicy(request CreatePolicyRequest, actor string, now time.Time) (Policy, error) {
	normalized, err := NormalizeCreateRequest(request)
	if err != nil {
		return Policy{}, err
	}
	if now.IsZero() {
		now = time.Now().UTC()
	}
	enabled := true
	if normalized.Enabled != nil {
		enabled = *normalized.Enabled
	}
	policy := Policy{
		Name:           normalized.Name,
		ProviderType:   normalized.ProviderType,
		Issuer:         normalized.Issuer,
		SignerIdentity: normalized.SignerIdentity,
		Subject:        normalized.Subject,
		Repository:     normalized.Repository,
		Workflow:       normalized.Workflow,
		Ref:            normalized.Ref,
		TenantID:       normalized.TenantID,
		ClusterID:      normalized.ClusterID,
		Environment:    normalized.Environment,
		Enabled:        enabled,
		CreatedAt:      now.UTC(),
		UpdatedAt:      now.UTC(),
		CreatedBy:      strings.TrimSpace(actor),
		UpdatedBy:      strings.TrimSpace(actor),
	}
	policy.ID = PolicyID(policy)
	return policy, nil
}

func PolicyID(policy Policy) string {
	parts := []string{
		normalizeProviderType(policy.ProviderType),
		strings.TrimSpace(policy.Issuer),
		strings.TrimSpace(policy.SignerIdentity),
		strings.TrimSpace(policy.Subject),
		strings.TrimSpace(policy.Repository),
		normalizeWorkflowPath(policy.Workflow),
		strings.TrimSpace(policy.Ref),
		strings.TrimSpace(policy.TenantID),
		strings.TrimSpace(policy.ClusterID),
		strings.TrimSpace(policy.Environment),
	}
	sum := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return "sid:" + hex.EncodeToString(sum[:])
}

func ObservationID(input DecisionInput, digest string) string {
	parts := []string{
		strings.TrimSpace(input.Issuer),
		strings.TrimSpace(input.SignerIdentity),
		strings.TrimSpace(input.Subject),
		strings.TrimSpace(input.Repository),
		normalizeWorkflowPath(input.Workflow),
		strings.TrimSpace(input.Ref),
		strings.TrimSpace(digest),
		strings.TrimSpace(input.TenantID),
		strings.TrimSpace(input.ClusterID),
		strings.TrimSpace(input.Environment),
	}
	sum := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return "obs:" + hex.EncodeToString(sum[:])
}

func Evaluate(config Config, policies []Policy, input DecisionInput) Decision {
	decision := Decision{
		Authorized:           AuthorizationUnknown,
		EnforcementMode:      config.Enforcement,
		ReasonCode:           ReasonUnknown,
		ReasonDetail:         "signing identity authorization could not be determined",
		TransparencyState:    strings.TrimSpace(input.TransparencyState),
		TransparencyRequired: config.RequireRekor,
	}

	if strings.TrimSpace(input.Issuer) == "" || strings.TrimSpace(input.SignerIdentity) == "" {
		decision.ReasonCode = ReasonEvidenceMissing
		decision.ReasonDetail = "issuer and signer identity are required for authorization"
		decision.Deny = config.Enforcement == EnforcementEnforce
		return decision
	}

	if config.RequireRekor {
		if decision.TransparencyState == "" || decision.TransparencyState == "disabled" || decision.TransparencyState == "unverified" {
			decision.Authorized = AuthorizationUnauthorized
			decision.ReasonCode = ReasonTransparencyUnverified
			decision.ReasonDetail = "transparency evidence is required and is not verified"
			decision.Deny = config.Enforcement == EnforcementEnforce
			return decision
		}
		if decision.TransparencyState != "verified" {
			decision.Authorized = AuthorizationUnauthorized
			decision.ReasonCode = ReasonRekorRequired
			decision.ReasonDetail = "transparency evidence is required for signer identity authorization"
			decision.Deny = config.Enforcement == EnforcementEnforce
			return decision
		}
	}

	var bestMismatchCode string
	var bestMismatchDetail string
	for _, policy := range policies {
		if policy.ProviderType == "" {
			policy.ProviderType = ProviderGitHubOIDC
		}
		if !policy.Enabled {
			if bestMismatchCode == "" {
				bestMismatchCode = ReasonPolicyDisabled
				bestMismatchDetail = "matching policy is disabled"
			}
			continue
		}
		if !scopeMatches(policy, input) {
			continue
		}
		if !providerMatches(policy.ProviderType, input.SignerIdentity) {
			bestMismatchCode = ReasonProviderUnsupported
			bestMismatchDetail = fmt.Sprintf("provider_type %q is not supported for identity %q", policy.ProviderType, input.SignerIdentity)
			continue
		}
		if mismatchCode, mismatchDetail := exactMatchReason(policy, input); mismatchCode != "" {
			if bestMismatchCode == "" {
				bestMismatchCode = mismatchCode
				bestMismatchDetail = mismatchDetail
			}
			continue
		}

		decision.Authorized = AuthorizationAuthorized
		decision.PolicyMatched = true
		decision.MatchedPolicyID = policy.ID
		decision.MatchedPolicyName = policy.Name
		decision.ReasonCode = ReasonAuthorized
		decision.ReasonDetail = "signer identity matches an enabled policy"

		if policy.DistrustedAfter != nil {
			decision.DistrustedAfter = timePointer(policy.DistrustedAfter.UTC())
			evidenceAt := decisionTime(input)
			if evidenceAt == nil {
				decision.Authorized = AuthorizationUnknown
				decision.ReasonCode = ReasonDistrustTimeUnavailable
				decision.ReasonDetail = "policy has a distrust cutoff but signing time evidence is unavailable"
				decision.Deny = config.Enforcement == EnforcementEnforce
				return decision
			}
			if !evidenceAt.Before(policy.DistrustedAfter.UTC()) {
				decision.Authorized = AuthorizationUnauthorized
				decision.ReasonCode = ReasonDistrustedAfterCutoff
				decision.ReasonDetail = "signer identity is distrusted after the configured cutoff time"
				decision.Deny = config.Enforcement == EnforcementEnforce
				return decision
			}
		}
		return decision
	}

	if bestMismatchCode == "" {
		bestMismatchCode = ReasonPolicyMissing
		bestMismatchDetail = "no enabled signing identity policy matched the observed signer"
	}
	decision.ReasonCode = bestMismatchCode
	decision.ReasonDetail = bestMismatchDetail
	if bestMismatchCode == ReasonPolicyMissing {
		decision.Authorized = AuthorizationUnknown
	} else {
		decision.Authorized = AuthorizationUnauthorized
	}
	decision.Deny = config.Enforcement == EnforcementEnforce
	return decision
}

func ScanWorkflowDocuments(root string) ([]WorkflowDocument, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return nil, nil
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	documents := make([]WorkflowDocument, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.ToLower(entry.Name())
		if !strings.HasSuffix(name, ".yml") && !strings.HasSuffix(name, ".yaml") {
			continue
		}
		path := filepath.Join(root, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		relative := normalizeWorkflowPath(path)
		documents = append(documents, WorkflowDocument{
			Path:           relative,
			SigningCapable: workflowRequestsIDToken(string(content)),
		})
	}
	return documents, nil
}

func BuildWorkflowFindings(workflows []WorkflowDocument, policies []Policy, now time.Time) []Finding {
	if now.IsZero() {
		now = time.Now().UTC()
	}
	covered := map[string]struct{}{}
	for _, policy := range policies {
		if policy.Enabled && policy.Workflow != "" {
			covered[normalizeWorkflowPath(policy.Workflow)] = struct{}{}
		}
	}

	findings := []Finding{}
	for _, workflow := range workflows {
		if workflow.SigningCapable {
			if _, ok := covered[normalizeWorkflowPath(workflow.Path)]; !ok {
				findings = append(findings, Finding{
					ID:         findingID(FindingWorkflowTokenNoPolicy, workflow.Path, ""),
					Type:       FindingWorkflowTokenNoPolicy,
					Severity:   "medium",
					Workflow:   normalizeWorkflowPath(workflow.Path),
					Reason:     "workflow requests id-token: write but no signing identity policy references it",
					DetectedAt: timePointer(now),
					Advisory:   true,
				})
			}
		}
	}

	available := map[string]struct{}{}
	for _, workflow := range workflows {
		available[normalizeWorkflowPath(workflow.Path)] = struct{}{}
	}
	for _, policy := range policies {
		if policy.Workflow == "" {
			continue
		}
		workflow := normalizeWorkflowPath(policy.Workflow)
		if _, ok := available[workflow]; !ok {
			findings = append(findings, Finding{
				ID:         findingID(FindingPolicyWorkflowMissing, workflow, policy.ID),
				Type:       FindingPolicyWorkflowMissing,
				Severity:   "low",
				Workflow:   workflow,
				PolicyID:   policy.ID,
				Reason:     "policy references a workflow path that is not present in the monitored repository",
				DetectedAt: timePointer(now),
				Advisory:   true,
			})
		}
	}
	return findings
}

func workflowRequestsIDToken(content string) bool {
	normalized := strings.ToLower(strings.ReplaceAll(content, "\r\n", "\n"))
	return strings.Contains(normalized, "id-token: write")
}

func exactMatchReason(policy Policy, input DecisionInput) (string, string) {
	if policy.Issuer != "" && policy.Issuer != strings.TrimSpace(input.Issuer) {
		return ReasonIssuerMismatch, "issuer did not match the configured signer policy"
	}
	if policy.SignerIdentity != "" && policy.SignerIdentity != strings.TrimSpace(input.SignerIdentity) {
		return ReasonSubjectMismatch, "certificate identity did not match the configured signer policy"
	}
	if policy.Subject != "" && policy.Subject != strings.TrimSpace(input.Subject) {
		return ReasonSubjectMismatch, "subject did not match the configured signer policy"
	}
	if policy.Repository != "" && policy.Repository != strings.TrimSpace(input.Repository) {
		return ReasonRepositoryMismatch, "repository did not match the configured signer policy"
	}
	if policy.Workflow != "" && normalizeWorkflowPath(policy.Workflow) != normalizeWorkflowPath(input.Workflow) {
		return ReasonWorkflowMismatch, "workflow path did not match the configured signer policy"
	}
	if policy.Ref != "" && policy.Ref != strings.TrimSpace(input.Ref) {
		return ReasonRefMismatch, "ref did not match the configured signer policy"
	}
	return "", ""
}

func scopeMatches(policy Policy, input DecisionInput) bool {
	if policy.TenantID != "" && policy.TenantID != strings.TrimSpace(input.TenantID) {
		return false
	}
	if policy.ClusterID != "" && policy.ClusterID != strings.TrimSpace(input.ClusterID) {
		return false
	}
	if policy.Environment != "" && policy.Environment != strings.TrimSpace(input.Environment) {
		return false
	}
	return true
}

func providerMatches(providerType string, signerIdentity string) bool {
	providerType = normalizeProviderType(providerType)
	switch providerType {
	case ProviderGitHubOIDC:
		return strings.Contains(strings.TrimSpace(signerIdentity), "github.com/")
	case ProviderGenericOIDC:
		return strings.TrimSpace(signerIdentity) != ""
	default:
		return false
	}
}

func decisionTime(input DecisionInput) *time.Time {
	if input.EvidenceAt == nil || input.EvidenceAt.IsZero() {
		return nil
	}
	timestamp := input.EvidenceAt.UTC()
	return &timestamp
}

func normalizeProviderType(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func normalizeWorkflowPath(value string) string {
	value = strings.TrimSpace(strings.ReplaceAll(value, "\\", "/"))
	if value == "" {
		return ""
	}
	idx := strings.Index(value, ".github/workflows/")
	if idx >= 0 {
		return value[idx:]
	}
	return value
}

func validProviderType(value string) bool {
	switch normalizeProviderType(value) {
	case ProviderGitHubOIDC, ProviderGenericOIDC:
		return true
	default:
		return false
	}
}

func parseBool(value string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("unsupported boolean value %q", value)
	}
}

func findingID(parts ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return "idf:" + hex.EncodeToString(sum[:])
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func timePointer(value time.Time) *time.Time {
	if value.IsZero() {
		return nil
	}
	timestamp := value.UTC()
	return &timestamp
}
