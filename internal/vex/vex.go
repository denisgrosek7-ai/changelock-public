package vex

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	SourceFormatAPI       = "api"
	SourceFormatCSAF      = "csaf"
	SourceFormatCycloneDX = "cyclonedx"

	StatusNotAffected       = "not_affected"
	StatusAffected          = "affected"
	StatusFixed             = "fixed"
	StatusUnderInvestigation = "under_investigation"

	LegacyDecisionNotAffected        = "NOT_AFFECTED"
	LegacyDecisionAcceptedRisk       = "ACCEPTED_RISK"
	LegacyDecisionFixRequired        = "FIX_REQUIRED"
	LegacyDecisionUnderInvestigation = "UNDER_INVESTIGATION"
)

var (
	ErrInvalidStatement = errors.New("invalid vex statement")
	ErrInvalidFilter    = errors.New("invalid vex filter")
)

type Scope struct {
	ImageDigest string `json:"image_digest,omitempty"`
	PackageName string `json:"package_name,omitempty"`
	PURL        string `json:"purl,omitempty"`
	Repo        string `json:"repo,omitempty"`
	Workload    string `json:"workload,omitempty"`
	TenantID    string `json:"tenant_id,omitempty"`
	ClusterID   string `json:"cluster_id,omitempty"`
	Environment string `json:"environment,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
}

type Statement struct {
	ID              int64           `json:"id"`
	StatementKey    string          `json:"statement_key,omitempty"`
	SourceFormat    string          `json:"source_format"`
	SourceRef       string          `json:"source_ref,omitempty"`
	VulnerabilityID string          `json:"vulnerability_id"`
	Scope           Scope           `json:"scope"`
	Status          string          `json:"status"`
	Justification   string          `json:"justification,omitempty"`
	ActionStatement string          `json:"action_statement,omitempty"`
	ImpactStatement string          `json:"impact_statement,omitempty"`
	FixedVersion    string          `json:"fixed_version,omitempty"`
	CreatedBy       string          `json:"created_by,omitempty"`
	UpdatedBy       string          `json:"updated_by,omitempty"`
	ExpiresAt       *time.Time      `json:"expires_at,omitempty"`
	RevokedAt       *time.Time      `json:"revoked_at,omitempty"`
	RevokedBy       string          `json:"revoked_by,omitempty"`
	Active          bool            `json:"active"`
	Metadata        json.RawMessage `json:"metadata,omitempty"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type CreateRequest struct {
	SourceFormat    string          `json:"source_format,omitempty"`
	SourceRef       string          `json:"source_ref,omitempty"`
	VulnerabilityID string          `json:"vulnerability_id"`
	Scope           Scope           `json:"scope"`
	Status          string          `json:"status"`
	Justification   string          `json:"justification,omitempty"`
	ActionStatement string          `json:"action_statement,omitempty"`
	ImpactStatement string          `json:"impact_statement,omitempty"`
	FixedVersion    string          `json:"fixed_version,omitempty"`
	ExpiresAt       *time.Time      `json:"expires_at,omitempty"`
	Metadata        json.RawMessage `json:"metadata,omitempty"`
}

type IngestRequest struct {
	Format    string          `json:"format,omitempty"`
	SourceRef string          `json:"source_ref,omitempty"`
	Scope     Scope           `json:"scope,omitempty"`
	ExpiresAt *time.Time      `json:"expires_at,omitempty"`
	Payload   json.RawMessage `json:"payload"`
}

type Filter struct {
	VulnerabilityID string
	ImageDigest     string
	PackageName     string
	PURL            string
	Repo            string
	Workload        string
	TenantID        string
	ClusterID       string
	Environment     string
	Namespace       string
	SourceFormat    string
	Status          string
	SourceRef       string
	Active          *bool
	Limit           int
}

type StatusSummary struct {
	ActiveCount     int            `json:"active_count"`
	ExpiringCount   int            `json:"expiring_count"`
	RevokedCount    int            `json:"revoked_count"`
	CountsByStatus  map[string]int `json:"counts_by_status"`
	AppliedFilters  map[string]string `json:"applied_filters,omitempty"`
}

type ImportResult struct {
	SourceFormat string      `json:"source_format"`
	Imported     int         `json:"imported"`
	Statements   []Statement `json:"statements"`
}

type FindingRef struct {
	VulnerabilityID string
	ImageDigest     string
	PackageName     string
	PURL            string
}

type EvaluationScope struct {
	TenantID    string
	ClusterID   string
	Repo        string
	Workload    string
	Environment string
	Namespace   string
}

type Config struct {
	ImportDir string
}

func ParseConfig(getenv func(string) string) (Config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}
	return Config{
		ImportDir: strings.TrimSpace(getenv("CHANGELOCK_VEX_IMPORT_DIR")),
	}, nil
}

func NormalizeCreateRequest(request CreateRequest, now func() time.Time) (CreateRequest, error) {
	if now == nil {
		now = time.Now
	}
	request.SourceFormat = normalizeSourceFormat(firstNonEmpty(request.SourceFormat, SourceFormatAPI))
	request.SourceRef = strings.TrimSpace(request.SourceRef)
	request.VulnerabilityID = normalizeVulnerabilityID(request.VulnerabilityID)
	request.Scope = normalizeScope(request.Scope)
	request.Status = normalizeStatus(request.Status)
	request.Justification = strings.TrimSpace(request.Justification)
	request.ActionStatement = strings.TrimSpace(request.ActionStatement)
	request.ImpactStatement = strings.TrimSpace(request.ImpactStatement)
	request.FixedVersion = strings.TrimSpace(request.FixedVersion)
	request.Metadata = normalizeJSON(request.Metadata)

	if request.SourceFormat == "" {
		return request, fmt.Errorf("%w: source_format is required", ErrInvalidStatement)
	}
	if request.VulnerabilityID == "" {
		return request, fmt.Errorf("%w: vulnerability_id is required", ErrInvalidStatement)
	}
	if request.Status == "" {
		return request, fmt.Errorf("%w: status is required", ErrInvalidStatement)
	}
	if !validStatus(request.Status) {
		return request, fmt.Errorf("%w: unsupported status %q", ErrInvalidStatement, request.Status)
	}
	if !scopeHasMatchTarget(request.Scope) {
		return request, fmt.Errorf("%w: scope must include image_digest, purl, or package_name", ErrInvalidStatement)
	}
	if request.ExpiresAt != nil {
		expiresAt := request.ExpiresAt.UTC()
		if !expiresAt.After(now().UTC()) {
			return request, fmt.Errorf("%w: expires_at must be in the future", ErrInvalidStatement)
		}
		request.ExpiresAt = &expiresAt
	}
	return request, nil
}

func NormalizeFilter(filter Filter) (Filter, error) {
	filter.VulnerabilityID = normalizeVulnerabilityID(filter.VulnerabilityID)
	filter.ImageDigest = strings.TrimSpace(filter.ImageDigest)
	filter.PackageName = strings.TrimSpace(filter.PackageName)
	filter.PURL = strings.TrimSpace(filter.PURL)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.Workload = strings.TrimSpace(filter.Workload)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.Namespace = strings.TrimSpace(filter.Namespace)
	filter.SourceFormat = normalizeSourceFormat(filter.SourceFormat)
	filter.Status = normalizeStatus(filter.Status)
	filter.SourceRef = strings.TrimSpace(filter.SourceRef)
	if filter.Status != "" && !validStatus(filter.Status) {
		return filter, fmt.Errorf("%w: unsupported status %q", ErrInvalidFilter, filter.Status)
	}
	if filter.Limit <= 0 {
		filter.Limit = 100
	}
	if filter.Limit > 1000 {
		filter.Limit = 1000
	}
	return filter, nil
}

func NormalizeStatement(statement Statement, now func() time.Time) (Statement, error) {
	request, err := NormalizeCreateRequest(CreateRequest{
		SourceFormat:    statement.SourceFormat,
		SourceRef:       statement.SourceRef,
		VulnerabilityID: statement.VulnerabilityID,
		Scope:           statement.Scope,
		Status:          statement.Status,
		Justification:   statement.Justification,
		ActionStatement: statement.ActionStatement,
		ImpactStatement: statement.ImpactStatement,
		FixedVersion:    statement.FixedVersion,
		ExpiresAt:       statement.ExpiresAt,
		Metadata:        statement.Metadata,
	}, now)
	if err != nil {
		return statement, err
	}
	statement.SourceFormat = request.SourceFormat
	statement.SourceRef = request.SourceRef
	statement.VulnerabilityID = request.VulnerabilityID
	statement.Scope = request.Scope
	statement.Status = request.Status
	statement.Justification = request.Justification
	statement.ActionStatement = request.ActionStatement
	statement.ImpactStatement = request.ImpactStatement
	statement.FixedVersion = request.FixedVersion
	statement.ExpiresAt = request.ExpiresAt
	statement.Metadata = request.Metadata
	statement.StatementKey = StatementIdentityKey(statement)
	return statement, nil
}

func Matches(statement Statement, finding FindingRef, scope EvaluationScope, now time.Time) bool {
	if !statement.Active {
		return false
	}
	if statement.RevokedAt != nil {
		return false
	}
	if statement.ExpiresAt != nil && !statement.ExpiresAt.After(now.UTC()) {
		return false
	}
	if normalizeVulnerabilityID(statement.VulnerabilityID) != normalizeVulnerabilityID(finding.VulnerabilityID) {
		return false
	}
	statementScope := normalizeScope(statement.Scope)
	if statementScope.ImageDigest != "" && statementScope.ImageDigest != strings.TrimSpace(finding.ImageDigest) {
		return false
	}
	if statementScope.PURL != "" && statementScope.PURL != strings.TrimSpace(finding.PURL) {
		return false
	}
	if statementScope.PackageName != "" && !strings.EqualFold(statementScope.PackageName, strings.TrimSpace(finding.PackageName)) {
		return false
	}
	if statementScope.TenantID != "" && statementScope.TenantID != strings.TrimSpace(scope.TenantID) {
		return false
	}
	if statementScope.ClusterID != "" && statementScope.ClusterID != strings.TrimSpace(scope.ClusterID) {
		return false
	}
	if statementScope.Repo != "" && statementScope.Repo != strings.TrimSpace(scope.Repo) {
		return false
	}
	if statementScope.Workload != "" && statementScope.Workload != strings.TrimSpace(scope.Workload) {
		return false
	}
	if statementScope.Environment != "" && statementScope.Environment != strings.TrimSpace(scope.Environment) {
		return false
	}
	if statementScope.Namespace != "" && statementScope.Namespace != strings.TrimSpace(scope.Namespace) {
		return false
	}
	if statementScope.ImageDigest == "" && statementScope.PURL == "" && statementScope.PackageName == "" {
		return false
	}
	return true
}

func SuppressesFinding(statement Statement) bool {
	return normalizeStatus(statement.Status) == StatusNotAffected
}

func StatementIdentityKey(statement Statement) string {
	scope := normalizeScope(statement.Scope)
	parts := []string{
		normalizeSourceFormat(statement.SourceFormat),
		strings.TrimSpace(statement.SourceRef),
		normalizeVulnerabilityID(statement.VulnerabilityID),
		scope.ImageDigest,
		strings.ToLower(scope.PackageName),
		scope.PURL,
		scope.Repo,
		scope.Workload,
		scope.TenantID,
		scope.ClusterID,
		scope.Environment,
		scope.Namespace,
	}
	sum := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return "vex:" + hex.EncodeToString(sum[:])
}

func normalizeSourceFormat(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case SourceFormatAPI:
		return SourceFormatAPI
	case SourceFormatCSAF:
		return SourceFormatCSAF
	case SourceFormatCycloneDX, "cyclonedx-vex":
		return SourceFormatCycloneDX
	default:
		return strings.ToLower(strings.TrimSpace(value))
	}
}

func normalizeStatus(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func normalizeVulnerabilityID(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func validStatus(value string) bool {
	switch normalizeStatus(value) {
	case StatusNotAffected, StatusAffected, StatusFixed, StatusUnderInvestigation:
		return true
	default:
		return false
	}
}

func normalizeScope(scope Scope) Scope {
	scope.ImageDigest = strings.TrimSpace(scope.ImageDigest)
	scope.PackageName = strings.TrimSpace(scope.PackageName)
	scope.PURL = strings.TrimSpace(scope.PURL)
	scope.Repo = strings.TrimSpace(scope.Repo)
	scope.Workload = strings.TrimSpace(scope.Workload)
	scope.TenantID = strings.TrimSpace(scope.TenantID)
	scope.ClusterID = strings.TrimSpace(scope.ClusterID)
	scope.Environment = strings.TrimSpace(scope.Environment)
	scope.Namespace = strings.TrimSpace(scope.Namespace)
	return scope
}

func scopeHasMatchTarget(scope Scope) bool {
	return scope.ImageDigest != "" || scope.PURL != "" || scope.PackageName != ""
}

func normalizeJSON(value json.RawMessage) json.RawMessage {
	if len(value) == 0 {
		return nil
	}
	return append(json.RawMessage(nil), value...)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
