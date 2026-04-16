package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	EventTypeVulnerabilityScanResult          = "vulnerability_scan_result"
	EventTypeVulnerabilityDriftDetected       = "vulnerability_drift_detected"
	EventTypeVulnerabilityDecisionRecorded    = "vulnerability_decision_recorded"
	EventTypeVulnerabilityDecisionDeactivated = "vulnerability_decision_deactivated"

	VulnerabilityFindingStatusOpen       = "OPEN"
	VulnerabilityFindingStatusResolved   = "RESOLVED"
	VulnerabilityFindingStatusSuppressed = "SUPPRESSED"

	VulnerabilityDecisionNotAffected        = "NOT_AFFECTED"
	VulnerabilityDecisionAcceptedRisk       = "ACCEPTED_RISK"
	VulnerabilityDecisionFixRequired        = "FIX_REQUIRED"
	VulnerabilityDecisionUnderInvestigation = "UNDER_INVESTIGATION"

	VulnerabilityScanModeManual   = "manual"
	VulnerabilityScanModePeriodic = "periodic"
	VulnerabilityScanModeOnDemand = "on-demand"

	VulnerabilityScanStatusPending   = "PENDING"
	VulnerabilityScanStatusCompleted = "COMPLETED"
	VulnerabilityScanStatusFailed    = "FAILED"

	SBOMFormatSPDXJSON      = "spdx-json"
	SBOMFormatCycloneDXJSON = "cyclonedx-json"
)

type SBOMIngestRequest struct {
	ImageDigest string          `json:"image_digest"`
	ImageRef    string          `json:"image_ref,omitempty"`
	SBOMFormat  string          `json:"sbom_format"`
	SourceRef   string          `json:"source_ref,omitempty"`
	SBOM        json.RawMessage `json:"sbom"`
}

type SBOMIngestResult struct {
	DocumentStored     bool   `json:"document_stored"`
	DocumentID         int64  `json:"document_id"`
	ImageDigest        string `json:"image_digest"`
	SBOMHash           string `json:"sbom_hash"`
	ComponentsIngested int    `json:"components_ingested"`
}

type SBOMDocument struct {
	ID          int64           `json:"id"`
	ImageDigest string          `json:"image_digest"`
	ImageRef    string          `json:"image_ref,omitempty"`
	SBOMFormat  string          `json:"sbom_format"`
	SourceRef   string          `json:"source_ref,omitempty"`
	SBOMHash    string          `json:"sbom_hash,omitempty"`
	RawSBOM     json.RawMessage `json:"raw_sbom,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
}

type SBOMComponent struct {
	ID               int64           `json:"id"`
	ImageDigest      string          `json:"image_digest"`
	ComponentName    string          `json:"component_name"`
	ComponentVersion string          `json:"component_version,omitempty"`
	ComponentType    string          `json:"component_type,omitempty"`
	License          string          `json:"license,omitempty"`
	PURL             string          `json:"purl,omitempty"`
	Metadata         json.RawMessage `json:"metadata,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
}

type SBOMImageResponse struct {
	Document       SBOMDocument    `json:"document"`
	ComponentCount int             `json:"component_count"`
	Components     []SBOMComponent `json:"components"`
}

type SBOMComponentSearchFilter struct {
	ComponentName string
	PURL          string
	ImageDigest   string
	TenantID      string
	Limit         int
}

type VulnerabilityScanRun struct {
	ID          int64           `json:"id"`
	ImageDigest string          `json:"image_digest"`
	ImageRef    string          `json:"image_ref,omitempty"`
	Scanner     string          `json:"scanner"`
	ScanMode    string          `json:"scan_mode"`
	StartedAt   time.Time       `json:"started_at"`
	CompletedAt *time.Time      `json:"completed_at,omitempty"`
	Status      string          `json:"status"`
	Summary     json.RawMessage `json:"summary,omitempty"`
	SourceRef   string          `json:"source_ref,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
}

type VulnerabilityFinding struct {
	ID             int64                  `json:"id"`
	ImageDigest    string                 `json:"image_digest"`
	ImageRef       string                 `json:"image_ref,omitempty"`
	ScanRunID      int64                  `json:"scan_run_id"`
	CVEID          string                 `json:"cve_id"`
	Severity       string                 `json:"severity,omitempty"`
	PackageName    string                 `json:"package_name,omitempty"`
	PackageVersion string                 `json:"package_version,omitempty"`
	FixedVersion   string                 `json:"fixed_version,omitempty"`
	PURL           string                 `json:"purl,omitempty"`
	Status         string                 `json:"status"`
	Title          string                 `json:"title,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Source         string                 `json:"source,omitempty"`
	Metadata       json.RawMessage        `json:"metadata,omitempty"`
	FirstSeenAt    time.Time              `json:"first_seen_at"`
	LastSeenAt     time.Time              `json:"last_seen_at"`
	Decision       *VulnerabilityDecision `json:"decision,omitempty"`
}

type VulnerabilityFindingInput struct {
	CVEID          string          `json:"cve_id"`
	Severity       string          `json:"severity,omitempty"`
	PackageName    string          `json:"package_name,omitempty"`
	PackageVersion string          `json:"package_version,omitempty"`
	FixedVersion   string          `json:"fixed_version,omitempty"`
	PURL           string          `json:"purl,omitempty"`
	Title          string          `json:"title,omitempty"`
	Description    string          `json:"description,omitempty"`
	Source         string          `json:"source,omitempty"`
	Metadata       json.RawMessage `json:"metadata,omitempty"`
}

type VulnerabilityScanRequest struct {
	ImageDigest string                      `json:"image_digest"`
	ImageRef    string                      `json:"image_ref,omitempty"`
	Scanner     string                      `json:"scanner"`
	ScanMode    string                      `json:"scan_mode"`
	StartedAt   time.Time                   `json:"started_at"`
	CompletedAt *time.Time                  `json:"completed_at,omitempty"`
	Status      string                      `json:"status"`
	Summary     json.RawMessage             `json:"summary,omitempty"`
	SourceRef   string                      `json:"source_ref,omitempty"`
	Findings    []VulnerabilityFindingInput `json:"findings,omitempty"`
}

type VulnerabilityScanIngestResult struct {
	Run                   VulnerabilityScanRun   `json:"run"`
	Findings              []VulnerabilityFinding `json:"findings"`
	NewFindings           []VulnerabilityFinding `json:"new_findings,omitempty"`
	HadPriorSuccessfulRun bool                   `json:"had_prior_successful_run"`
}

type VulnerabilityDecision struct {
	ID            int64           `json:"id"`
	ImageDigest   string          `json:"image_digest"`
	CVEID         string          `json:"cve_id"`
	Decision      string          `json:"decision"`
	Justification string          `json:"justification"`
	DecidedBy     string          `json:"decided_by"`
	ExpiresAt     *time.Time      `json:"expires_at,omitempty"`
	Active        bool            `json:"active"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type VulnerabilityDecisionCreateRequest struct {
	ImageDigest   string          `json:"image_digest"`
	CVEID         string          `json:"cve_id"`
	Decision      string          `json:"decision"`
	Justification string          `json:"justification"`
	ExpiresAt     *time.Time      `json:"expires_at,omitempty"`
	TTLHours      int             `json:"ttl_hours,omitempty"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
}

type VulnerabilityDecisionFilter struct {
	ImageDigest string
	CVEID       string
	TenantID    string
	Active      *bool
	Limit       int
}

type VulnerabilityActiveFilter struct {
	Severity          string
	CVEID             string
	ImageDigest       string
	ComponentName     string
	TenantID          string
	Environment       string
	Limit             int
	IncludeSuppressed bool
}

type ActiveWorkloadRef struct {
	TenantID    string `json:"tenant_id,omitempty"`
	Environment string `json:"environment,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
	Workload    string `json:"workload,omitempty"`
	Repo        string `json:"repo,omitempty"`
	Image       string `json:"image,omitempty"`
	Digest      string `json:"digest,omitempty"`
}

type VulnerabilityBlastRadiusFilter struct {
	CVEID         string
	ComponentName string
	PURL          string
	TenantID      string
	Limit         int
}

type VulnerabilityBlastRadiusItem struct {
	ImageDigest string                 `json:"image_digest"`
	ImageRef    string                 `json:"image_ref,omitempty"`
	Findings    []VulnerabilityFinding `json:"findings"`
	Workloads   []ActiveWorkloadRef    `json:"workloads"`
}

type VulnerabilityBlastRadiusResponse struct {
	Items          []VulnerabilityBlastRadiusItem `json:"items"`
	AppliedFilters map[string]string              `json:"applied_filters"`
}

type VulnerabilityTimelineFilter struct {
	ImageDigest string
	CVEID       string
	TenantID    string
	WindowDays  int
}

type VulnerabilityTimelineEntry struct {
	ImageDigest    string                 `json:"image_digest"`
	CVEID          string                 `json:"cve_id"`
	PackageName    string                 `json:"package_name,omitempty"`
	PackageVersion string                 `json:"package_version,omitempty"`
	Severity       string                 `json:"severity,omitempty"`
	Status         string                 `json:"status"`
	FirstSeenAt    time.Time              `json:"first_seen_at"`
	LastSeenAt     time.Time              `json:"last_seen_at"`
	Decision       *VulnerabilityDecision `json:"decision,omitempty"`
}

type VulnerabilityTimelineResponse struct {
	Items          []VulnerabilityTimelineEntry `json:"items"`
	AppliedFilters map[string]string            `json:"applied_filters"`
}

type VulnerabilityRescanRequest struct {
	ImageDigest string `json:"image_digest,omitempty"`
	ImageRef    string `json:"image_ref,omitempty"`
}

type VulnerabilityRescanResponse struct {
	Status         string   `json:"status"`
	ScannedDigests []string `json:"scanned_digests"`
	ScanRuns       int      `json:"scan_runs"`
}

type ActiveDigestRef struct {
	ImageDigest string              `json:"image_digest"`
	ImageRef    string              `json:"image_ref,omitempty"`
	Repo        string              `json:"repo,omitempty"`
	Scopes      []ActiveWorkloadRef `json:"scopes,omitempty"`
}

func NormalizeSBOMIngestRequest(request SBOMIngestRequest) (SBOMIngestRequest, error) {
	request.ImageDigest = strings.TrimSpace(request.ImageDigest)
	request.ImageRef = strings.TrimSpace(request.ImageRef)
	request.SourceRef = strings.TrimSpace(request.SourceRef)
	request.SBOMFormat = normalizeSBOMFormat(request.SBOMFormat)
	request.SBOM = normalizeMetadata(request.SBOM)

	switch {
	case request.ImageDigest == "":
		return request, fmt.Errorf("%w: image_digest is required", ErrInvalidException)
	case request.SBOMFormat == "":
		return request, fmt.Errorf("%w: sbom_format is required", ErrInvalidException)
	case len(request.SBOM) == 0:
		return request, fmt.Errorf("%w: sbom is required", ErrInvalidException)
	}

	switch request.SBOMFormat {
	case SBOMFormatSPDXJSON, SBOMFormatCycloneDXJSON:
	default:
		return request, fmt.Errorf("%w: unsupported sbom_format %q", ErrInvalidException, request.SBOMFormat)
	}
	return request, nil
}

func NormalizeSBOMComponentSearchFilter(filter SBOMComponentSearchFilter) (SBOMComponentSearchFilter, error) {
	filter.ComponentName = strings.TrimSpace(filter.ComponentName)
	filter.PURL = strings.TrimSpace(filter.PURL)
	filter.ImageDigest = strings.TrimSpace(filter.ImageDigest)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	if filter.ComponentName == "" && filter.PURL == "" && filter.ImageDigest == "" {
		return filter, fmt.Errorf("%w: at least one component search filter is required", ErrInvalidFilter)
	}
	if filter.Limit <= 0 {
		filter.Limit = 50
	}
	if filter.Limit > 500 {
		filter.Limit = 500
	}
	return filter, nil
}

func NormalizeVulnerabilityScanRequest(request VulnerabilityScanRequest, now func() time.Time) (VulnerabilityScanRequest, error) {
	if now == nil {
		now = time.Now
	}
	request.ImageDigest = strings.TrimSpace(request.ImageDigest)
	request.ImageRef = strings.TrimSpace(request.ImageRef)
	request.Scanner = strings.TrimSpace(strings.ToLower(request.Scanner))
	request.ScanMode = strings.TrimSpace(strings.ToLower(request.ScanMode))
	request.SourceRef = strings.TrimSpace(request.SourceRef)
	request.Summary = normalizeMetadata(request.Summary)

	switch {
	case request.ImageDigest == "":
		return request, fmt.Errorf("%w: image_digest is required", ErrInvalidException)
	case request.Scanner == "":
		return request, fmt.Errorf("%w: scanner is required", ErrInvalidException)
	}

	if request.ScanMode == "" {
		request.ScanMode = VulnerabilityScanModeOnDemand
	}
	if request.StartedAt.IsZero() {
		request.StartedAt = now().UTC()
	}
	if request.Status == "" {
		request.Status = VulnerabilityScanStatusCompleted
	}
	for i := range request.Findings {
		request.Findings[i].CVEID = strings.TrimSpace(strings.ToUpper(request.Findings[i].CVEID))
		request.Findings[i].Severity = strings.TrimSpace(strings.ToUpper(request.Findings[i].Severity))
		request.Findings[i].PackageName = strings.TrimSpace(request.Findings[i].PackageName)
		request.Findings[i].PackageVersion = strings.TrimSpace(request.Findings[i].PackageVersion)
		request.Findings[i].FixedVersion = strings.TrimSpace(request.Findings[i].FixedVersion)
		request.Findings[i].PURL = strings.TrimSpace(request.Findings[i].PURL)
		request.Findings[i].Title = strings.TrimSpace(request.Findings[i].Title)
		request.Findings[i].Description = strings.TrimSpace(request.Findings[i].Description)
		request.Findings[i].Source = strings.TrimSpace(request.Findings[i].Source)
		request.Findings[i].Metadata = normalizeMetadata(request.Findings[i].Metadata)
		if request.Findings[i].CVEID == "" {
			return request, fmt.Errorf("%w: findings[%d].cve_id is required", ErrInvalidException, i)
		}
	}
	return request, nil
}

func NormalizeVulnerabilityDecisionCreateRequest(request VulnerabilityDecisionCreateRequest, now func() time.Time) (VulnerabilityDecisionCreateRequest, error) {
	if now == nil {
		now = time.Now
	}
	request.ImageDigest = strings.TrimSpace(request.ImageDigest)
	request.CVEID = strings.TrimSpace(strings.ToUpper(request.CVEID))
	request.Decision = strings.TrimSpace(strings.ToUpper(request.Decision))
	request.Justification = strings.TrimSpace(request.Justification)
	request.Metadata = normalizeMetadata(request.Metadata)

	switch {
	case request.ImageDigest == "":
		return request, fmt.Errorf("%w: image_digest is required", ErrInvalidException)
	case request.CVEID == "":
		return request, fmt.Errorf("%w: cve_id is required", ErrInvalidException)
	case request.Decision == "":
		return request, fmt.Errorf("%w: decision is required", ErrInvalidException)
	case request.Justification == "":
		return request, fmt.Errorf("%w: justification is required", ErrInvalidException)
	}

	switch request.Decision {
	case VulnerabilityDecisionNotAffected, VulnerabilityDecisionAcceptedRisk, VulnerabilityDecisionFixRequired, VulnerabilityDecisionUnderInvestigation:
	default:
		return request, fmt.Errorf("%w: unsupported vulnerability decision %q", ErrInvalidException, request.Decision)
	}

	if request.ExpiresAt == nil && request.TTLHours > 0 {
		expiresAt := now().UTC().Add(time.Duration(request.TTLHours) * time.Hour)
		request.ExpiresAt = &expiresAt
	}
	if request.ExpiresAt != nil {
		expiresAt := request.ExpiresAt.UTC()
		if !expiresAt.After(now().UTC()) {
			return request, fmt.Errorf("%w: expires_at must be in the future", ErrInvalidException)
		}
		request.ExpiresAt = &expiresAt
	}
	return request, nil
}

func NormalizeVulnerabilityDecisionFilter(filter VulnerabilityDecisionFilter) VulnerabilityDecisionFilter {
	filter.ImageDigest = strings.TrimSpace(filter.ImageDigest)
	filter.CVEID = strings.TrimSpace(strings.ToUpper(filter.CVEID))
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	if filter.Limit <= 0 {
		filter.Limit = 50
	}
	if filter.Limit > 500 {
		filter.Limit = 500
	}
	return filter
}

func NormalizeVulnerabilityActiveFilter(filter VulnerabilityActiveFilter) VulnerabilityActiveFilter {
	filter.Severity = strings.TrimSpace(strings.ToUpper(filter.Severity))
	filter.CVEID = strings.TrimSpace(strings.ToUpper(filter.CVEID))
	filter.ImageDigest = strings.TrimSpace(filter.ImageDigest)
	filter.ComponentName = strings.TrimSpace(filter.ComponentName)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	if filter.Limit <= 0 {
		filter.Limit = 50
	}
	if filter.Limit > 500 {
		filter.Limit = 500
	}
	return filter
}

func NormalizeVulnerabilityBlastRadiusFilter(filter VulnerabilityBlastRadiusFilter) (VulnerabilityBlastRadiusFilter, error) {
	filter.CVEID = strings.TrimSpace(strings.ToUpper(filter.CVEID))
	filter.ComponentName = strings.TrimSpace(filter.ComponentName)
	filter.PURL = strings.TrimSpace(filter.PURL)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	if filter.CVEID == "" && filter.ComponentName == "" && filter.PURL == "" {
		return filter, fmt.Errorf("%w: at least one blast-radius filter is required", ErrInvalidFilter)
	}
	if filter.Limit <= 0 {
		filter.Limit = 50
	}
	if filter.Limit > 500 {
		filter.Limit = 500
	}
	return filter, nil
}

func NormalizeVulnerabilityTimelineFilter(filter VulnerabilityTimelineFilter) (VulnerabilityTimelineFilter, error) {
	filter.ImageDigest = strings.TrimSpace(filter.ImageDigest)
	filter.CVEID = strings.TrimSpace(strings.ToUpper(filter.CVEID))
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	if filter.ImageDigest == "" {
		return filter, fmt.Errorf("%w: image_digest is required", ErrInvalidFilter)
	}
	if filter.CVEID == "" {
		return filter, fmt.Errorf("%w: cve_id is required", ErrInvalidFilter)
	}
	if filter.WindowDays <= 0 {
		filter.WindowDays = 30
	}
	if filter.WindowDays > 365 {
		filter.WindowDays = 365
	}
	return filter, nil
}

func normalizeSBOMFormat(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "spdx-json", "spdxjson", "spdx":
		return SBOMFormatSPDXJSON
	case "cyclonedx-json", "cyclonedxjson", "cyclonedx":
		return SBOMFormatCycloneDXJSON
	default:
		return strings.ToLower(strings.TrimSpace(value))
	}
}

func sbomHash(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	sum := sha256.Sum256(raw)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func cloneSBOMDocument(document SBOMDocument) SBOMDocument {
	if document.RawSBOM != nil {
		document.RawSBOM = append(document.RawSBOM[:0:0], document.RawSBOM...)
	}
	return document
}

func cloneSBOMComponent(component SBOMComponent) SBOMComponent {
	if component.Metadata != nil {
		component.Metadata = append(component.Metadata[:0:0], component.Metadata...)
	}
	return component
}

func cloneVulnerabilityFinding(finding VulnerabilityFinding) VulnerabilityFinding {
	if finding.Metadata != nil {
		finding.Metadata = append(finding.Metadata[:0:0], finding.Metadata...)
	}
	if finding.Decision != nil {
		decision := cloneVulnerabilityDecision(*finding.Decision)
		finding.Decision = &decision
	}
	return finding
}

func cloneVulnerabilityDecision(decision VulnerabilityDecision) VulnerabilityDecision {
	if decision.Metadata != nil {
		decision.Metadata = append(decision.Metadata[:0:0], decision.Metadata...)
	}
	return decision
}

func vulnerabilityFindingKey(imageDigest string, finding VulnerabilityFindingInput) string {
	return strings.Join([]string{
		strings.TrimSpace(imageDigest),
		strings.TrimSpace(strings.ToUpper(finding.CVEID)),
		strings.TrimSpace(finding.PackageName),
		strings.TrimSpace(finding.PackageVersion),
		strings.TrimSpace(finding.PURL),
	}, "|")
}

func activeDecisionApplies(decision *VulnerabilityDecision, now time.Time) bool {
	if decision == nil || !decision.Active {
		return false
	}
	if decision.ExpiresAt != nil && !decision.ExpiresAt.After(now.UTC()) {
		return false
	}
	return true
}
