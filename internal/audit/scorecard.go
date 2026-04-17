package audit

import (
	"fmt"
	"html"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/signingidentity"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

const (
	TrustMetricStatusVerified = "verified"
	TrustMetricStatusPartial  = "partial"
	TrustMetricStatusGap      = "gap"
	TrustMetricStatusUnknown  = "unknown"

	TrustBadgeStateVerified = "verified"
	TrustBadgeStatePartial  = "partial"
	TrustBadgeStateGap      = "gap"
	TrustBadgeStateUnknown  = "unknown"

	TrustGradeA = "A"
	TrustGradeB = "B"
	TrustGradeC = "C"
	TrustGradeD = "D"
	TrustGradeF = "F"

	ScorecardMetricArtifactIntegrity = "artifact_integrity"
	ScorecardMetricVulnerability     = "vulnerability_posture"
	ScorecardMetricSignerGovernance  = "signer_identity_governance"
	ScorecardMetricRuntimeHardening  = "runtime_hardening"
	ScorecardMetricExceptionHygiene  = "exception_hygiene"
	ScorecardMetricPolicyEvidence    = "policy_enforcement"

	ScorecardReasonArtifactVerified = "artifact_integrity_verified"
	ScorecardReasonArtifactPartial  = "artifact_integrity_partial"
	ScorecardReasonArtifactMissing  = "artifact_integrity_missing"

	ScorecardReasonVulnClear      = "vulnerability_posture_clear"
	ScorecardReasonVulnManaged    = "vulnerability_posture_vex_managed"
	ScorecardReasonVulnActionable = "vulnerability_posture_actionable"
	ScorecardReasonVulnCritical   = "vulnerability_posture_critical"
	ScorecardReasonVulnUnknown    = "vulnerability_posture_unknown"

	ScorecardReasonSignerEnforced = "signer_governance_enforced"
	ScorecardReasonSignerMonitor  = "signer_governance_monitored"
	ScorecardReasonSignerGap      = "signer_governance_gap"
	ScorecardReasonSignerUnknown  = "signer_governance_unknown"

	ScorecardReasonRuntimeHealthy = "runtime_hardening_in_sync"
	ScorecardReasonRuntimePartial = "runtime_hardening_partial"
	ScorecardReasonRuntimeGap     = "runtime_hardening_gap"
	ScorecardReasonRuntimeUnknown = "runtime_hardening_unknown"

	ScorecardReasonExceptionsClear   = "exception_hygiene_clear"
	ScorecardReasonExceptionsPending = "exception_hygiene_pending"
	ScorecardReasonExceptionsStale   = "exception_hygiene_stale"
	ScorecardReasonExceptionsActive  = "exception_hygiene_active"

	ScorecardReasonPolicyEvidenced = "policy_enforcement_evidenced"
	ScorecardReasonPolicyPartial   = "policy_enforcement_partial"
	ScorecardReasonPolicyUnknown   = "policy_enforcement_unknown"

	HardeningFindingStaleExceptions          = "stale_exception_active"
	HardeningFindingVulnerabilityDebt        = "actionable_vulnerability_above_threshold"
	HardeningFindingMissingSignerPolicies    = "signer_policy_missing"
	HardeningFindingSignerFindingsActive     = "signer_identity_findings_active"
	HardeningFindingRuntimeContainmentActive = "runtime_quarantine_active"
	HardeningFindingPolicyEvidenceGap        = "policy_evidence_gap"
	HardeningFindingTransparencyGap          = "artifact_transparency_gap"

	TrustPublicationDisabled = "disabled"
	TrustPublicationPreview  = "preview"
	TrustPublicationExport   = "export"
)

type TrustScorecardInput struct {
	ScopeType                  string
	ScopeRef                   string
	TenantID                   string
	ClusterID                  string
	Environment                string
	Repo                       string
	CalculatedAt               time.Time
	ArtifactVerificationEvents []StoredEvent
	PolicyDecisionEvents       []StoredEvent
	Summary                    Summary
	VulnerabilityNet           VulnerabilityNetResponse
	VEXStatus                  internalvex.StatusSummary
	SigningIdentityStatus      signingidentity.StatusSummary
	SigningIdentityFindings    []signingidentity.Finding
	RuntimeStatus              RuntimeClosedLoopStatus
	RuntimeActiveStates        []RuntimeActiveStateView
	ExceptionReport            ExceptionReport
	PublicationMode            string
	StaleExceptionDays         int
}

type TrustScorecard struct {
	ID                           string             `json:"id"`
	ScopeType                    string             `json:"scope_type"`
	ScopeRef                     string             `json:"scope_ref"`
	TenantID                     string             `json:"tenant_id,omitempty"`
	ClusterID                    string             `json:"cluster_id,omitempty"`
	Environment                  string             `json:"environment,omitempty"`
	Repo                         string             `json:"repo,omitempty"`
	CalculatedAt                 time.Time          `json:"calculated_at"`
	OverallGrade                 string             `json:"overall_grade"`
	OverallScore                 int                `json:"overall_score"`
	SigningCoverage              int                `json:"signing_coverage"`
	TransparencyCoverage         int                `json:"transparency_coverage"`
	SBOMCoverage                 int                `json:"sbom_or_provenance_coverage"`
	ActionableVulnerabilityCount int                `json:"actionable_vulnerability_count"`
	StaleExceptionCount          int                `json:"stale_exception_count"`
	PublicationMode              string             `json:"publication_mode"`
	Metrics                      []TrustScoreMetric `json:"metrics"`
	Notes                        []string           `json:"notes,omitempty"`
}

type TrustScoreMetric struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Weight            int      `json:"weight"`
	Score             int      `json:"score"`
	Status            string   `json:"status"`
	ReasonCode        string   `json:"reason_code"`
	ReasonDetail      string   `json:"reason_detail,omitempty"`
	EvidenceRefs      []string `json:"evidence_refs,omitempty"`
	AdvisoryOnly      bool     `json:"advisory_only"`
	PublicPublishable bool     `json:"public_publishable"`
	MappingRefs       []string `json:"mapping_refs,omitempty"`
}

type TrustBadge struct {
	ID                string `json:"id"`
	Label             string `json:"label"`
	State             string `json:"state"`
	Summary           string `json:"summary"`
	PublicPublishable bool   `json:"public_publishable"`
	SVG               string `json:"svg,omitempty"`
}

type AuditFinding struct {
	ID                string    `json:"id"`
	Category          string    `json:"category"`
	Severity          string    `json:"severity"`
	Status            string    `json:"status"`
	ReasonCode        string    `json:"reason_code"`
	ReasonDetail      string    `json:"reason_detail,omitempty"`
	ScopeRef          string    `json:"scope_ref,omitempty"`
	EvidenceRefs      []string  `json:"evidence_refs,omitempty"`
	AdvisoryOnly      bool      `json:"advisory_only"`
	PublicPublishable bool      `json:"public_publishable"`
	DetectedAt        time.Time `json:"detected_at"`
}

type StandardsMapping struct {
	Standard     string   `json:"standard"`
	Control      string   `json:"control"`
	Status       string   `json:"status"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
}

type PublishedTrustView struct {
	GeneratedAt  time.Time          `json:"generated_at"`
	ScopeType    string             `json:"scope_type"`
	ScopeRef     string             `json:"scope_ref"`
	OverallGrade string             `json:"overall_grade"`
	OverallScore int                `json:"overall_score"`
	Badges       []TrustBadge       `json:"badges"`
	Metrics      []TrustScoreMetric `json:"metrics"`
	Mapping      []StandardsMapping `json:"mapping"`
	Notes        []string           `json:"notes,omitempty"`
}

type AuditReport struct {
	ID               string              `json:"id"`
	GeneratedAt      time.Time           `json:"generated_at"`
	ScopeType        string              `json:"scope_type"`
	ScopeRef         string              `json:"scope_ref"`
	Scorecard        TrustScorecard      `json:"scorecard"`
	Findings         []AuditFinding      `json:"findings"`
	Badges           []TrustBadge        `json:"badges"`
	StandardsMapping []StandardsMapping  `json:"standards_mapping"`
	PublicView       *PublishedTrustView `json:"public_view,omitempty"`
	Limitations      []string            `json:"limitations,omitempty"`
	Format           string              `json:"format,omitempty"`
	GeneratedBy      string              `json:"generated_by,omitempty"`
}

type TrustEvidenceExport struct {
	ArtifactEvidence  []TrustArtifactEvidenceItem  `json:"artifact_evidence,omitempty"`
	ExceptionEvidence []TrustExceptionEvidenceItem `json:"exception_evidence,omitempty"`
}

type TrustArtifactEvidenceItem struct {
	Digest            string `json:"digest,omitempty"`
	Repo              string `json:"repo,omitempty"`
	SignerIdentity    string `json:"signer_identity,omitempty"`
	VerificationState string `json:"verification_state,omitempty"`
	BundleRef         string `json:"bundle_ref,omitempty"`
	LogEntryID        string `json:"log_entry_id,omitempty"`
	SBOMRef           string `json:"sbom_ref,omitempty"`
}

type TrustExceptionEvidenceItem struct {
	ExceptionID        string `json:"exception_id"`
	Status             string `json:"status"`
	VerificationState  string `json:"verification_state,omitempty"`
	VerificationReason string `json:"verification_reason,omitempty"`
	TicketID           string `json:"ticket_id,omitempty"`
}

type AuditExportBundle struct {
	GeneratedAt time.Time           `json:"generated_at"`
	ScopeType   string              `json:"scope_type"`
	ScopeRef    string              `json:"scope_ref"`
	Scorecard   TrustScorecard      `json:"scorecard"`
	Report      AuditReport         `json:"report"`
	Evidence    TrustEvidenceExport `json:"evidence"`
	PublicView  *PublishedTrustView `json:"public_view,omitempty"`
}

func ComputeTrustScorecard(input TrustScorecardInput) TrustScorecard {
	if input.CalculatedAt.IsZero() {
		input.CalculatedAt = time.Now().UTC()
	}
	if strings.TrimSpace(input.ScopeType) == "" {
		input.ScopeType = "global"
	}
	if strings.TrimSpace(input.ScopeRef) == "" {
		input.ScopeRef = input.ScopeType + ":default"
	}
	if input.StaleExceptionDays <= 0 {
		input.StaleExceptionDays = 14
	}
	card := TrustScorecard{
		ID:              trustScopeID(input.ScopeType, input.ScopeRef),
		ScopeType:       input.ScopeType,
		ScopeRef:        input.ScopeRef,
		TenantID:        strings.TrimSpace(input.TenantID),
		ClusterID:       strings.TrimSpace(input.ClusterID),
		Environment:     strings.TrimSpace(input.Environment),
		Repo:            strings.TrimSpace(input.Repo),
		CalculatedAt:    input.CalculatedAt.UTC(),
		PublicationMode: normalizeTrustPublicationMode(input.PublicationMode),
	}

	artifactSample := collectArtifactEvidence(input.ArtifactVerificationEvents)
	card.SigningCoverage = artifactSample.SigningCoverage
	card.TransparencyCoverage = artifactSample.TransparencyCoverage
	card.SBOMCoverage = artifactSample.SBOMCoverage
	card.ActionableVulnerabilityCount = input.VulnerabilityNet.ActionableCount
	card.StaleExceptionCount = staleExceptions(input.ExceptionReport, input.CalculatedAt, input.StaleExceptionDays)

	metrics := []TrustScoreMetric{
		scoreArtifactIntegrityMetric(artifactSample),
		scoreVulnerabilityMetric(input.VulnerabilityNet),
		scoreSignerGovernanceMetric(input.SigningIdentityStatus),
		scoreRuntimeMetric(input.RuntimeStatus, len(input.RuntimeActiveStates)),
		scoreExceptionMetric(input.ExceptionReport, card.StaleExceptionCount),
		scorePolicyMetric(input.PolicyDecisionEvents),
	}
	card.Metrics = metrics

	total := 0
	for _, metric := range metrics {
		total += metric.Score
	}
	card.OverallScore = clampScore(total)
	card.OverallGrade = trustGradeForScore(card.OverallScore)
	card.Notes = scorecardNotes(input, artifactSample)
	return card
}

func BuildTrustBadges(scorecard TrustScorecard, input TrustScorecardInput) []TrustBadge {
	metric := metricByID(scorecard.Metrics)
	badges := []TrustBadge{
		newTrustBadge("overall", "Trust Grade", badgeStateForMetricScore(scorecard.OverallScore, scorecard.OverallGrade), fmt.Sprintf("Measured trust posture grade %s (%d/100).", scorecard.OverallGrade, scorecard.OverallScore), true),
		newTrustBadge("sbom_available", "SBOM Available", badgeStateForPercent(scorecard.SBOMCoverage), coverageSummary("SBOM / provenance coverage", scorecard.SBOMCoverage), true),
		newTrustBadge("vex_triage", "VEX Triage Active", vexBadgeState(input.VEXStatus, input.VulnerabilityNet), vexBadgeSummary(input.VEXStatus, input.VulnerabilityNet), true),
		newTrustBadge("signer_governed", "Signer Identity Governed", metric[ScorecardMetricSignerGovernance].Status, metric[ScorecardMetricSignerGovernance].ReasonDetail, true),
		newTrustBadge("runtime_hardening", "Runtime Hardening Active", metric[ScorecardMetricRuntimeHardening].Status, metric[ScorecardMetricRuntimeHardening].ReasonDetail, false),
	}
	for i := range badges {
		badges[i].SVG = RenderTrustBadgeSVG(badges[i])
	}
	return badges
}

func BuildHardeningReview(input TrustScorecardInput, scorecard TrustScorecard) []AuditFinding {
	now := input.CalculatedAt
	if now.IsZero() {
		now = time.Now().UTC()
	}
	findings := make([]AuditFinding, 0)
	if scorecard.StaleExceptionCount > 0 {
		findings = append(findings, AuditFinding{
			ID:           trustFindingID(scorecard.ScopeRef, HardeningFindingStaleExceptions),
			Category:     "exceptions",
			Severity:     "medium",
			Status:       "open",
			ReasonCode:   HardeningFindingStaleExceptions,
			ReasonDetail: fmt.Sprintf("%d active exceptions are older than %d days.", scorecard.StaleExceptionCount, max(input.StaleExceptionDays, 14)),
			ScopeRef:     scorecard.ScopeRef,
			EvidenceRefs: []string{"report:/v1/reports/exceptions"},
			AdvisoryOnly: true,
			DetectedAt:   now,
		})
	}
	if input.VulnerabilityNet.ActionableCount > 0 {
		severity := "medium"
		if input.VulnerabilityNet.ThresholdBreached {
			severity = "high"
		}
		findings = append(findings, AuditFinding{
			ID:           trustFindingID(scorecard.ScopeRef, HardeningFindingVulnerabilityDebt),
			Category:     "vulnerabilities",
			Severity:     severity,
			Status:       "open",
			ReasonCode:   HardeningFindingVulnerabilityDebt,
			ReasonDetail: fmt.Sprintf("%d net actionable vulnerabilities remain after VEX merge.", input.VulnerabilityNet.ActionableCount),
			ScopeRef:     scorecard.ScopeRef,
			EvidenceRefs: []string{"report:/v1/vulnerabilities/net"},
			AdvisoryOnly: true,
			DetectedAt:   now,
		})
	}
	if input.SigningIdentityStatus.ObservedIdentities > 0 && input.SigningIdentityStatus.EnabledPolicies == 0 {
		findings = append(findings, AuditFinding{
			ID:           trustFindingID(scorecard.ScopeRef, HardeningFindingMissingSignerPolicies),
			Category:     "signing",
			Severity:     "high",
			Status:       "open",
			ReasonCode:   HardeningFindingMissingSignerPolicies,
			ReasonDetail: "Observed signing identities exist, but no enabled signer policies are recorded for this scope.",
			ScopeRef:     scorecard.ScopeRef,
			EvidenceRefs: []string{"status:/v1/signing-identities/status"},
			AdvisoryOnly: true,
			DetectedAt:   now,
		})
	}
	if input.SigningIdentityStatus.Unauthorized > 0 || input.SigningIdentityStatus.Unknown > 0 || input.SigningIdentityStatus.Findings > 0 {
		findings = append(findings, AuditFinding{
			ID:           trustFindingID(scorecard.ScopeRef, HardeningFindingSignerFindingsActive),
			Category:     "signing",
			Severity:     "high",
			Status:       "open",
			ReasonCode:   HardeningFindingSignerFindingsActive,
			ReasonDetail: fmt.Sprintf("%d unauthorized, %d unknown signer observations and %d findings remain active.", input.SigningIdentityStatus.Unauthorized, input.SigningIdentityStatus.Unknown, input.SigningIdentityStatus.Findings),
			ScopeRef:     scorecard.ScopeRef,
			EvidenceRefs: []string{"status:/v1/signing-identities/status", "status:/v1/signing-identities/findings"},
			AdvisoryOnly: true,
			DetectedAt:   now,
		})
	}
	if input.RuntimeStatus.Quarantined > 0 || input.RuntimeStatus.Failed > 0 {
		findings = append(findings, AuditFinding{
			ID:           trustFindingID(scorecard.ScopeRef, HardeningFindingRuntimeContainmentActive),
			Category:     "runtime",
			Severity:     "medium",
			Status:       "open",
			ReasonCode:   HardeningFindingRuntimeContainmentActive,
			ReasonDetail: fmt.Sprintf("Runtime closed loop reports %d quarantined and %d failed targets.", input.RuntimeStatus.Quarantined, input.RuntimeStatus.Failed),
			ScopeRef:     scorecard.ScopeRef,
			EvidenceRefs: []string{"status:/v1/runtime/closed-loop/status", "status:/v1/runtime/quarantine"},
			AdvisoryOnly: true,
			DetectedAt:   now,
		})
	}
	if metric := findMetric(scorecard.Metrics, ScorecardMetricPolicyEvidence); metric.Status == TrustMetricStatusUnknown || metric.Status == TrustMetricStatusGap {
		findings = append(findings, AuditFinding{
			ID:           trustFindingID(scorecard.ScopeRef, HardeningFindingPolicyEvidenceGap),
			Category:     "policy",
			Severity:     "medium",
			Status:       "open",
			ReasonCode:   HardeningFindingPolicyEvidenceGap,
			ReasonDetail: metric.ReasonDetail,
			ScopeRef:     scorecard.ScopeRef,
			EvidenceRefs: metric.EvidenceRefs,
			AdvisoryOnly: true,
			DetectedAt:   now,
		})
	}
	if metric := findMetric(scorecard.Metrics, ScorecardMetricArtifactIntegrity); metric.Status == TrustMetricStatusGap || metric.Status == TrustMetricStatusUnknown {
		findings = append(findings, AuditFinding{
			ID:           trustFindingID(scorecard.ScopeRef, HardeningFindingTransparencyGap),
			Category:     "evidence",
			Severity:     "medium",
			Status:       "open",
			ReasonCode:   HardeningFindingTransparencyGap,
			ReasonDetail: metric.ReasonDetail,
			ScopeRef:     scorecard.ScopeRef,
			EvidenceRefs: metric.EvidenceRefs,
			AdvisoryOnly: true,
			DetectedAt:   now,
		})
	}
	sort.Slice(findings, func(i, j int) bool {
		if findings[i].Severity == findings[j].Severity {
			return findings[i].ReasonCode < findings[j].ReasonCode
		}
		return findingSeverityRank(findings[i].Severity) > findingSeverityRank(findings[j].Severity)
	})
	return findings
}

func BuildStandardsMapping(scorecard TrustScorecard) []StandardsMapping {
	metric := metricByID(scorecard.Metrics)
	return []StandardsMapping{
		{
			Standard:     "NIST SSDF",
			Control:      "PS.3 / PW.6",
			Status:       mappingStatus(metric[ScorecardMetricArtifactIntegrity].Status),
			Summary:      metric[ScorecardMetricArtifactIntegrity].ReasonDetail,
			EvidenceRefs: metric[ScorecardMetricArtifactIntegrity].EvidenceRefs,
		},
		{
			Standard:     "NIST SSDF",
			Control:      "RV.1 / RV.3",
			Status:       mappingStatus(metric[ScorecardMetricVulnerability].Status),
			Summary:      metric[ScorecardMetricVulnerability].ReasonDetail,
			EvidenceRefs: metric[ScorecardMetricVulnerability].EvidenceRefs,
		},
		{
			Standard:     "SLSA readiness",
			Control:      "provenance_and_identity_governance",
			Status:       mappingStatus(minMetricStatus(metric[ScorecardMetricArtifactIntegrity].Status, metric[ScorecardMetricSignerGovernance].Status)),
			Summary:      "Mapped from measured artifact evidence, transparency verification, and signer identity governance signals.",
			EvidenceRefs: append(append([]string{}, metric[ScorecardMetricArtifactIntegrity].EvidenceRefs...), metric[ScorecardMetricSignerGovernance].EvidenceRefs...),
		},
		{
			Standard:     "Internal controls",
			Control:      "runtime_hardening_and_exception_hygiene",
			Status:       mappingStatus(minMetricStatus(metric[ScorecardMetricRuntimeHardening].Status, metric[ScorecardMetricExceptionHygiene].Status)),
			Summary:      "Mapped from runtime closed-loop posture and active exception hygiene signals.",
			EvidenceRefs: append(append([]string{}, metric[ScorecardMetricRuntimeHardening].EvidenceRefs...), metric[ScorecardMetricExceptionHygiene].EvidenceRefs...),
		},
	}
}

func BuildPublishedTrustView(scorecard TrustScorecard, badges []TrustBadge, mappings []StandardsMapping) *PublishedTrustView {
	if normalizeTrustPublicationMode(scorecard.PublicationMode) == TrustPublicationDisabled {
		return nil
	}
	publicMetrics := make([]TrustScoreMetric, 0, len(scorecard.Metrics))
	for _, metric := range scorecard.Metrics {
		if metric.PublicPublishable {
			publicMetrics = append(publicMetrics, TrustScoreMetric{
				ID:           metric.ID,
				Name:         metric.Name,
				Weight:       metric.Weight,
				Score:        metric.Score,
				Status:       metric.Status,
				ReasonCode:   metric.ReasonCode,
				ReasonDetail: metric.ReasonDetail,
			})
		}
	}
	publicBadges := make([]TrustBadge, 0, len(badges))
	for _, badge := range badges {
		if badge.PublicPublishable {
			publicBadges = append(publicBadges, badge)
		}
	}
	publicMappings := make([]StandardsMapping, 0, len(mappings))
	for _, mapping := range mappings {
		publicMappings = append(publicMappings, StandardsMapping{
			Standard: mapping.Standard,
			Control:  mapping.Control,
			Status:   mapping.Status,
			Summary:  mapping.Summary,
		})
	}
	return &PublishedTrustView{
		GeneratedAt:  scorecard.CalculatedAt,
		ScopeType:    scorecard.ScopeType,
		ScopeRef:     scorecard.ScopeRef,
		OverallGrade: scorecard.OverallGrade,
		OverallScore: scorecard.OverallScore,
		Badges:       publicBadges,
		Metrics:      publicMetrics,
		Mapping:      publicMappings,
		Notes: []string{
			"Sanitized trust view derived from measured internal posture.",
			"This is a readiness and coverage view, not a formal certification claim.",
		},
	}
}

func BuildAuditReport(scorecard TrustScorecard, findings []AuditFinding, badges []TrustBadge, mappings []StandardsMapping, publicView *PublishedTrustView, format string, generatedBy string) AuditReport {
	format = strings.ToLower(strings.TrimSpace(firstNonEmpty(format, "json")))
	return AuditReport{
		ID:               trustScopeID(scorecard.ScopeType, scorecard.ScopeRef) + ":" + scorecard.CalculatedAt.UTC().Format("20060102T150405Z"),
		GeneratedAt:      scorecard.CalculatedAt,
		ScopeType:        scorecard.ScopeType,
		ScopeRef:         scorecard.ScopeRef,
		Scorecard:        scorecard,
		Findings:         findings,
		Badges:           badges,
		StandardsMapping: mappings,
		PublicView:       publicView,
		Limitations: []string{
			"Grades reflect measured repository signals only; missing data stays partial or unknown.",
			"Standards mappings are readiness and evidence mappings, not certification claims.",
			"Current scorecard scope is bounded to the signals ChangeLock already records today.",
		},
		Format:      format,
		GeneratedBy: strings.TrimSpace(generatedBy),
	}
}

func BuildAuditExportBundle(report AuditReport, source TrustEvidenceExport) AuditExportBundle {
	return AuditExportBundle{
		GeneratedAt: report.GeneratedAt,
		ScopeType:   report.ScopeType,
		ScopeRef:    report.ScopeRef,
		Scorecard:   report.Scorecard,
		Report:      report,
		Evidence:    source,
		PublicView:  report.PublicView,
	}
}

func RenderAuditReportHTML(report AuditReport) string {
	var builder strings.Builder
	builder.WriteString("<!doctype html><html><head><meta charset=\"utf-8\"><title>ChangeLock Hardening Audit</title>")
	builder.WriteString("<style>body{font-family:IBM Plex Sans,Segoe UI,sans-serif;background:#f7f8fb;color:#16202a;padding:32px}h1,h2,h3{margin:0 0 12px}section{margin:0 0 24px;padding:20px;background:#fff;border:1px solid #d7dee6;border-radius:16px}table{width:100%;border-collapse:collapse}th,td{text-align:left;padding:10px;border-bottom:1px solid #e3e8ef}small{color:#56697c}</style>")
	builder.WriteString("</head><body>")
	builder.WriteString("<h1>ChangeLock Hardening Audit</h1>")
	builder.WriteString("<p><small>Scope: " + html.EscapeString(report.ScopeRef) + " · Generated at " + html.EscapeString(report.GeneratedAt.Format(time.RFC3339)) + "</small></p>")
	builder.WriteString("<section><h2>Scorecard</h2>")
	builder.WriteString(fmt.Sprintf("<p>Overall grade <strong>%s</strong> (%d/100).</p>", html.EscapeString(report.Scorecard.OverallGrade), report.Scorecard.OverallScore))
	builder.WriteString("<table><thead><tr><th>Metric</th><th>Status</th><th>Score</th><th>Reason</th></tr></thead><tbody>")
	for _, metric := range report.Scorecard.Metrics {
		builder.WriteString("<tr><td>" + html.EscapeString(metric.Name) + "</td><td>" + html.EscapeString(metric.Status) + "</td><td>" + fmt.Sprintf("%d/%d", metric.Score, metric.Weight) + "</td><td>" + html.EscapeString(metric.ReasonDetail) + "</td></tr>")
	}
	builder.WriteString("</tbody></table></section>")
	builder.WriteString("<section><h2>Hardening Review</h2><table><thead><tr><th>Category</th><th>Severity</th><th>Reason</th></tr></thead><tbody>")
	if len(report.Findings) == 0 {
		builder.WriteString("<tr><td colspan=\"3\">No active hardening review findings.</td></tr>")
	} else {
		for _, finding := range report.Findings {
			builder.WriteString("<tr><td>" + html.EscapeString(finding.Category) + "</td><td>" + html.EscapeString(finding.Severity) + "</td><td>" + html.EscapeString(finding.ReasonDetail) + "</td></tr>")
		}
	}
	builder.WriteString("</tbody></table></section>")
	builder.WriteString("<section><h2>Standards Mapping</h2><table><thead><tr><th>Standard</th><th>Control</th><th>Status</th><th>Summary</th></tr></thead><tbody>")
	for _, mapping := range report.StandardsMapping {
		builder.WriteString("<tr><td>" + html.EscapeString(mapping.Standard) + "</td><td>" + html.EscapeString(mapping.Control) + "</td><td>" + html.EscapeString(mapping.Status) + "</td><td>" + html.EscapeString(mapping.Summary) + "</td></tr>")
	}
	builder.WriteString("</tbody></table></section>")
	builder.WriteString("</body></html>")
	return builder.String()
}

func RenderTrustBadgeSVG(badge TrustBadge) string {
	label := html.EscapeString(strings.TrimSpace(firstNonEmpty(badge.Label, badge.ID)))
	value := html.EscapeString(strings.TrimSpace(firstNonEmpty(badge.State, "unknown")))
	color := "#617384"
	switch badge.State {
	case TrustBadgeStateVerified:
		color = "#027a48"
	case TrustBadgeStatePartial:
		color = "#b54708"
	case TrustBadgeStateGap:
		color = "#b42318"
	}
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="220" height="28" role="img" aria-label="%s: %s"><rect width="120" height="28" fill="#12202f"/><rect x="120" width="100" height="28" fill="%s"/><text x="60" y="19" fill="#fff" font-family="IBM Plex Sans,Segoe UI,sans-serif" font-size="12" text-anchor="middle">%s</text><text x="170" y="19" fill="#fff" font-family="IBM Plex Sans,Segoe UI,sans-serif" font-size="12" text-anchor="middle">%s</text></svg>`, label, value, color, label, value)
}

type artifactEvidenceSample struct {
	SampleSize           int
	VerifiedCount        int
	TransparentCount     int
	SBOMCount            int
	SigningCoverage      int
	TransparencyCoverage int
	SBOMCoverage         int
}

func collectArtifactEvidence(events []StoredEvent) artifactEvidenceSample {
	sample := artifactEvidenceSample{}
	for _, event := range events {
		if event.EventType != EventTypeArtifactVerificationResult || event.Evidence == nil || event.Evidence.Artifact == nil {
			continue
		}
		sample.SampleSize++
		if event.VerifierSummary != nil && event.VerifierSummary.SignatureValid && event.VerifierSummary.AttestationValid {
			sample.VerifiedCount++
		}
		if event.Evidence.VerificationState == "verified" {
			sample.TransparentCount++
		}
		if strings.TrimSpace(event.Evidence.Artifact.SBOMArtifactRef) != "" || strings.TrimSpace(event.Evidence.Artifact.SBOMDigestRef) != "" || strings.TrimSpace(event.Evidence.Artifact.SBOMHash) != "" {
			sample.SBOMCount++
		}
	}
	if sample.SampleSize > 0 {
		sample.SigningCoverage = int((float64(sample.VerifiedCount) / float64(sample.SampleSize)) * 100)
		sample.TransparencyCoverage = int((float64(sample.TransparentCount) / float64(sample.SampleSize)) * 100)
		sample.SBOMCoverage = int((float64(sample.SBOMCount) / float64(sample.SampleSize)) * 100)
	}
	return sample
}

func scoreArtifactIntegrityMetric(sample artifactEvidenceSample) TrustScoreMetric {
	metric := TrustScoreMetric{
		ID:                ScorecardMetricArtifactIntegrity,
		Name:              "Artifact Integrity",
		Weight:            25,
		PublicPublishable: true,
		EvidenceRefs:      []string{"event_type:artifact_verification_result"},
		MappingRefs:       []string{"NIST SSDF PS.3", "SLSA readiness provenance_and_signing"},
	}
	if sample.SampleSize == 0 {
		metric.Status = TrustMetricStatusUnknown
		metric.ReasonCode = ScorecardReasonArtifactMissing
		metric.ReasonDetail = "No recent artifact verification evidence is available for this scope."
		return metric
	}
	score := weightedPercent(sample.SigningCoverage, 50) + weightedPercent(sample.TransparencyCoverage, 30) + weightedPercent(sample.SBOMCoverage, 20)
	metric.Score = scaleMetricScore(score, metric.Weight)
	metric.ReasonDetail = fmt.Sprintf("Verified signatures/attestations on %d%% of sampled artifacts, transparency evidence on %d%%, SBOM/provenance references on %d%%.", sample.SigningCoverage, sample.TransparencyCoverage, sample.SBOMCoverage)
	switch {
	case sample.SigningCoverage >= 90 && sample.TransparencyCoverage >= 80:
		metric.Status = TrustMetricStatusVerified
		metric.ReasonCode = ScorecardReasonArtifactVerified
	case sample.SigningCoverage >= 60:
		metric.Status = TrustMetricStatusPartial
		metric.ReasonCode = ScorecardReasonArtifactPartial
	default:
		metric.Status = TrustMetricStatusGap
		metric.ReasonCode = ScorecardReasonArtifactMissing
	}
	return metric
}

func scoreVulnerabilityMetric(response VulnerabilityNetResponse) TrustScoreMetric {
	metric := TrustScoreMetric{
		ID:                ScorecardMetricVulnerability,
		Name:              "Vulnerability Posture",
		Weight:            20,
		PublicPublishable: true,
		EvidenceRefs:      []string{"report:/v1/vulnerabilities/net"},
		MappingRefs:       []string{"NIST SSDF RV.1", "NIST SSDF RV.3"},
	}
	if response.RawCount == 0 && response.ActionableCount == 0 && response.ResolvedByVEXCount == 0 {
		metric.Status = TrustMetricStatusUnknown
		metric.ReasonCode = ScorecardReasonVulnUnknown
		metric.ReasonDetail = "No vulnerability posture signal is currently available for this scope."
		return metric
	}
	score := 100
	switch {
	case response.ActionableCount == 0 && response.RawCount == 0:
		score = 100
	case response.ActionableCount == 0:
		score = 92
	case response.ThresholdBreached:
		score = 35
	default:
		score = max(45, 90-(response.ActionableCount*6))
	}
	metric.Score = scaleMetricScore(score, metric.Weight)
	if response.ActionableCount == 0 && response.ResolvedByVEXCount > 0 {
		metric.Status = TrustMetricStatusVerified
		metric.ReasonCode = ScorecardReasonVulnManaged
		metric.ReasonDetail = fmt.Sprintf("%d findings are resolved by VEX and no net actionable vulnerabilities remain.", response.ResolvedByVEXCount)
		return metric
	}
	if response.ActionableCount == 0 {
		metric.Status = TrustMetricStatusVerified
		metric.ReasonCode = ScorecardReasonVulnClear
		metric.ReasonDetail = "No net actionable vulnerabilities remain for the current scope."
		return metric
	}
	if response.ThresholdBreached {
		metric.Status = TrustMetricStatusGap
		metric.ReasonCode = ScorecardReasonVulnCritical
		metric.ReasonDetail = fmt.Sprintf("%d net actionable vulnerabilities remain and the configured severity threshold is breached.", response.ActionableCount)
		return metric
	}
	metric.Status = TrustMetricStatusPartial
	metric.ReasonCode = ScorecardReasonVulnActionable
	metric.ReasonDetail = fmt.Sprintf("%d net actionable vulnerabilities remain after VEX merge.", response.ActionableCount)
	return metric
}

func scoreSignerGovernanceMetric(status signingidentity.StatusSummary) TrustScoreMetric {
	metric := TrustScoreMetric{
		ID:                ScorecardMetricSignerGovernance,
		Name:              "Signer Identity Governance",
		Weight:            15,
		PublicPublishable: true,
		EvidenceRefs:      []string{"status:/v1/signing-identities/status"},
		MappingRefs:       []string{"SLSA readiness build_identity_governance"},
	}
	if status.ObservedIdentities == 0 && status.TotalPolicies == 0 {
		metric.Status = TrustMetricStatusUnknown
		metric.ReasonCode = ScorecardReasonSignerUnknown
		metric.ReasonDetail = "No signing identity observations or policies are available for this scope."
		return metric
	}
	score := 40
	switch status.EnforcementMode {
	case signingidentity.EnforcementEnforce:
		score = 100
	case signingidentity.EnforcementMonitor:
		score = 80
	default:
		score = 45
	}
	score -= (status.Unauthorized * 15) + (status.Unknown * 10)
	score -= min(status.Findings*5, 20)
	score = clampPercent(score)
	metric.Score = scaleMetricScore(score, metric.Weight)
	switch {
	case status.Unauthorized == 0 && status.Unknown == 0 && status.Findings == 0 && status.EnforcementMode == signingidentity.EnforcementEnforce:
		metric.Status = TrustMetricStatusVerified
		metric.ReasonCode = ScorecardReasonSignerEnforced
		metric.ReasonDetail = "Signer identity policies are enforced and no unauthorized or unknown observations are active."
	case status.Unauthorized == 0 && status.Unknown == 0:
		metric.Status = TrustMetricStatusPartial
		metric.ReasonCode = ScorecardReasonSignerMonitor
		metric.ReasonDetail = fmt.Sprintf("Signer identities are governed in %s mode with %d active findings.", status.EnforcementMode, status.Findings)
	default:
		metric.Status = TrustMetricStatusGap
		metric.ReasonCode = ScorecardReasonSignerGap
		metric.ReasonDetail = fmt.Sprintf("%d unauthorized and %d unknown signer observations remain active.", status.Unauthorized, status.Unknown)
	}
	return metric
}

func scoreRuntimeMetric(status RuntimeClosedLoopStatus, activeTargets int) TrustScoreMetric {
	metric := TrustScoreMetric{
		ID:                ScorecardMetricRuntimeHardening,
		Name:              "Runtime Hardening",
		Weight:            15,
		PublicPublishable: false,
		EvidenceRefs:      []string{"status:/v1/runtime/closed-loop/status"},
		MappingRefs:       []string{"Internal controls runtime_hardening"},
	}
	if activeTargets == 0 || status.TotalTargets == 0 {
		metric.Status = TrustMetricStatusUnknown
		metric.ReasonCode = ScorecardReasonRuntimeUnknown
		metric.ReasonDetail = "No closed-loop runtime state is available for this scope."
		return metric
	}
	inSyncPercent := int((float64(status.InSync) / float64(maxInt64(status.TotalTargets, 1))) * 100)
	score := inSyncPercent
	score -= int(status.Quarantined * 10)
	score -= int(status.Failed * 15)
	score = clampPercent(score)
	metric.Score = scaleMetricScore(score, metric.Weight)
	switch {
	case status.Quarantined == 0 && status.Failed == 0 && inSyncPercent >= 80:
		metric.Status = TrustMetricStatusVerified
		metric.ReasonCode = ScorecardReasonRuntimeHealthy
		metric.ReasonDetail = fmt.Sprintf("%d%% of observed runtime targets are in sync and no failed or quarantined targets remain.", inSyncPercent)
	case status.Failed == 0:
		metric.Status = TrustMetricStatusPartial
		metric.ReasonCode = ScorecardReasonRuntimePartial
		metric.ReasonDetail = fmt.Sprintf("%d%% of observed runtime targets are in sync, with %d quarantined targets under containment.", inSyncPercent, status.Quarantined)
	default:
		metric.Status = TrustMetricStatusGap
		metric.ReasonCode = ScorecardReasonRuntimeGap
		metric.ReasonDetail = fmt.Sprintf("Runtime closed loop reports %d failed and %d quarantined targets.", status.Failed, status.Quarantined)
	}
	return metric
}

func scoreExceptionMetric(report ExceptionReport, staleCount int) TrustScoreMetric {
	metric := TrustScoreMetric{
		ID:                ScorecardMetricExceptionHygiene,
		Name:              "Exception Hygiene",
		Weight:            10,
		PublicPublishable: false,
		EvidenceRefs:      []string{"report:/v1/reports/exceptions"},
		MappingRefs:       []string{"Internal controls exception_hygiene"},
	}
	activeCount := len(report.Active)
	pendingCount := len(report.Pending)
	if activeCount == 0 && pendingCount == 0 && staleCount == 0 {
		metric.Status = TrustMetricStatusVerified
		metric.ReasonCode = ScorecardReasonExceptionsClear
		metric.ReasonDetail = "No active, pending, or stale exceptions are currently recorded."
		metric.Score = metric.Weight
		return metric
	}
	score := 100 - (activeCount * 12) - (pendingCount * 10) - (staleCount * 20)
	score = clampPercent(score)
	metric.Score = scaleMetricScore(score, metric.Weight)
	switch {
	case staleCount > 0:
		metric.Status = TrustMetricStatusGap
		metric.ReasonCode = ScorecardReasonExceptionsStale
		metric.ReasonDetail = fmt.Sprintf("%d active exceptions are older than the stale threshold.", staleCount)
	case pendingCount > 0:
		metric.Status = TrustMetricStatusPartial
		metric.ReasonCode = ScorecardReasonExceptionsPending
		metric.ReasonDetail = fmt.Sprintf("%d pending exceptions still require review.", pendingCount)
	default:
		metric.Status = TrustMetricStatusPartial
		metric.ReasonCode = ScorecardReasonExceptionsActive
		metric.ReasonDetail = fmt.Sprintf("%d active exceptions remain open.", activeCount)
	}
	return metric
}

func scorePolicyMetric(events []StoredEvent) TrustScoreMetric {
	metric := TrustScoreMetric{
		ID:                ScorecardMetricPolicyEvidence,
		Name:              "Policy Enforcement Evidence",
		Weight:            15,
		PublicPublishable: true,
		EvidenceRefs:      []string{"event_type:policy_decision", "event_type:deploy_gate_decision"},
		MappingRefs:       []string{"NIST SSDF PS.1", "NIST SSDF PW.9"},
	}
	if len(events) == 0 {
		metric.Status = TrustMetricStatusUnknown
		metric.ReasonCode = ScorecardReasonPolicyUnknown
		metric.ReasonDetail = "No recent policy or deploy-gate decisions are available for this scope."
		return metric
	}
	evidenced := 0
	for _, event := range events {
		if strings.TrimSpace(event.PolicyBundleHash) != "" || strings.TrimSpace(event.PolicyBundleID) != "" || strings.TrimSpace(event.PolicyVersion) != "" {
			evidenced++
		}
	}
	coverage := int((float64(evidenced) / float64(len(events))) * 100)
	metric.Score = scaleMetricScore(coverage, metric.Weight)
	metric.ReasonDetail = fmt.Sprintf("%d%% of sampled policy decisions include bundle/version evidence.", coverage)
	switch {
	case coverage >= 80:
		metric.Status = TrustMetricStatusVerified
		metric.ReasonCode = ScorecardReasonPolicyEvidenced
	case coverage >= 50:
		metric.Status = TrustMetricStatusPartial
		metric.ReasonCode = ScorecardReasonPolicyPartial
	default:
		metric.Status = TrustMetricStatusGap
		metric.ReasonCode = ScorecardReasonPolicyUnknown
	}
	return metric
}

func BuildTrustEvidenceExport(input TrustScorecardInput) TrustEvidenceExport {
	export := TrustEvidenceExport{
		ArtifactEvidence:  make([]TrustArtifactEvidenceItem, 0, len(input.ArtifactVerificationEvents)),
		ExceptionEvidence: make([]TrustExceptionEvidenceItem, 0, len(input.ExceptionReport.Active)+len(input.ExceptionReport.Pending)),
	}
	for _, event := range input.ArtifactVerificationEvents {
		if event.Evidence == nil || event.Evidence.Artifact == nil {
			continue
		}
		item := TrustArtifactEvidenceItem{
			Digest:            firstNonEmpty(event.Evidence.Artifact.Digest, event.Digest),
			Repo:              firstNonEmpty(event.Evidence.Artifact.Repository, event.Repo),
			SignerIdentity:    event.Evidence.Artifact.SignerIdentity,
			VerificationState: firstNonEmpty(event.Evidence.VerificationState, event.Evidence.VerificationReason),
			SBOMRef:           firstNonEmpty(event.Evidence.Artifact.SBOMArtifactRef, event.Evidence.Artifact.SBOMDigestRef),
		}
		if event.Evidence.Bundle != nil {
			item.BundleRef = event.Evidence.Bundle.BundleRef
			item.LogEntryID = event.Evidence.Bundle.LogEntryID
		}
		export.ArtifactEvidence = append(export.ArtifactEvidence, item)
	}
	seen := map[string]struct{}{}
	for _, exception := range append(append([]PolicyException{}, input.ExceptionReport.Active...), input.ExceptionReport.Pending...) {
		if _, ok := seen[exception.ExceptionID]; ok {
			continue
		}
		seen[exception.ExceptionID] = struct{}{}
		export.ExceptionEvidence = append(export.ExceptionEvidence, TrustExceptionEvidenceItem{
			ExceptionID:        exception.ExceptionID,
			Status:             exception.Status,
			VerificationState:  exception.VerificationState,
			VerificationReason: exception.VerificationReason,
			TicketID:           exception.TicketID,
		})
	}
	return export
}

func staleExceptions(report ExceptionReport, now time.Time, staleDays int) int {
	if staleDays <= 0 {
		staleDays = 14
	}
	cutoff := now.UTC().Add(-time.Duration(staleDays) * 24 * time.Hour)
	count := 0
	for _, item := range report.Active {
		updatedAt := item.CreatedAt
		if item.LastUpdatedAt != nil && !item.LastUpdatedAt.IsZero() {
			updatedAt = item.LastUpdatedAt.UTC()
		} else if item.ApprovedAt != nil && !item.ApprovedAt.IsZero() {
			updatedAt = item.ApprovedAt.UTC()
		} else if item.RequestedAt != nil && !item.RequestedAt.IsZero() {
			updatedAt = item.RequestedAt.UTC()
		}
		if updatedAt.Before(cutoff) {
			count++
		}
	}
	return count
}

func newTrustBadge(id, label, state, summary string, publicPublishable bool) TrustBadge {
	state = normalizeMetricStatus(state)
	switch state {
	case TrustMetricStatusVerified:
		state = TrustBadgeStateVerified
	case TrustMetricStatusPartial:
		state = TrustBadgeStatePartial
	case TrustMetricStatusGap:
		state = TrustBadgeStateGap
	default:
		state = TrustBadgeStateUnknown
	}
	return TrustBadge{
		ID:                strings.TrimSpace(id),
		Label:             strings.TrimSpace(label),
		State:             state,
		Summary:           strings.TrimSpace(summary),
		PublicPublishable: publicPublishable,
	}
}

func scorecardNotes(input TrustScorecardInput, sample artifactEvidenceSample) []string {
	notes := []string{
		"Scorecards are derived from measured ChangeLock signals and stay conservative when data is partial or unknown.",
	}
	if sample.SampleSize > 0 {
		notes = append(notes, fmt.Sprintf("Artifact integrity is sampled from the latest %d artifact verification events within the selected scope.", sample.SampleSize))
	} else {
		notes = append(notes, "Artifact integrity is currently unknown because no artifact verification events were available in the sampled scope.")
	}
	if normalizeTrustPublicationMode(input.PublicationMode) == TrustPublicationDisabled {
		notes = append(notes, "Public trust publication is disabled by default; sanitized publication requires explicit operator enablement.")
	}
	return notes
}

func metricByID(metrics []TrustScoreMetric) map[string]TrustScoreMetric {
	result := make(map[string]TrustScoreMetric, len(metrics))
	for _, metric := range metrics {
		result[metric.ID] = metric
	}
	return result
}

func findMetric(metrics []TrustScoreMetric, id string) TrustScoreMetric {
	for _, metric := range metrics {
		if metric.ID == id {
			return metric
		}
	}
	return TrustScoreMetric{ID: id, Status: TrustMetricStatusUnknown}
}

func weightedPercent(percent int, weight int) int {
	return int((float64(clampPercent(percent)) / 100.0) * float64(weight))
}

func scaleMetricScore(percent int, weight int) int {
	return int((float64(clampPercent(percent)) / 100.0) * float64(weight))
}

func clampPercent(value int) int {
	if value < 0 {
		return 0
	}
	if value > 100 {
		return 100
	}
	return value
}

func clampScore(value int) int {
	if value < 0 {
		return 0
	}
	if value > 100 {
		return 100
	}
	return value
}

func trustGradeForScore(score int) string {
	switch {
	case score >= 90:
		return TrustGradeA
	case score >= 80:
		return TrustGradeB
	case score >= 70:
		return TrustGradeC
	case score >= 60:
		return TrustGradeD
	default:
		return TrustGradeF
	}
}

func trustScopeID(scopeType, scopeRef string) string {
	scopeType = strings.TrimSpace(firstNonEmpty(scopeType, "global"))
	scopeRef = strings.TrimSpace(firstNonEmpty(scopeRef, scopeType+":default"))
	return scopeType + ":" + scopeRef
}

func trustFindingID(scopeRef, reasonCode string) string {
	scopeRef = strings.TrimSpace(firstNonEmpty(scopeRef, "global:default"))
	reasonCode = strings.TrimSpace(firstNonEmpty(reasonCode, "finding"))
	return strings.ReplaceAll(scopeRef, ":", "_") + ":" + reasonCode
}

func normalizeMetricStatus(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case TrustMetricStatusVerified, TrustMetricStatusPartial, TrustMetricStatusGap:
		return strings.ToLower(strings.TrimSpace(value))
	default:
		return TrustMetricStatusUnknown
	}
}

func normalizeTrustPublicationMode(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case TrustPublicationPreview, TrustPublicationExport:
		return strings.ToLower(strings.TrimSpace(value))
	default:
		return TrustPublicationDisabled
	}
}

func coverageSummary(label string, coverage int) string {
	return fmt.Sprintf("%s is measured at %d%% for the current scope.", strings.TrimSpace(label), clampPercent(coverage))
}

func vexBadgeState(summary internalvex.StatusSummary, response VulnerabilityNetResponse) string {
	switch {
	case response.ResolvedByVEXCount > 0 || summary.ActiveCount > 0:
		return TrustBadgeStateVerified
	case response.RawCount > 0:
		return TrustBadgeStatePartial
	default:
		return TrustBadgeStateUnknown
	}
}

func vexBadgeSummary(summary internalvex.StatusSummary, response VulnerabilityNetResponse) string {
	switch {
	case response.ResolvedByVEXCount > 0:
		return fmt.Sprintf("%d findings are currently resolved by active VEX statements.", response.ResolvedByVEXCount)
	case summary.ActiveCount > 0:
		return fmt.Sprintf("%d active VEX statements are recorded for this scope.", summary.ActiveCount)
	case response.RawCount > 0:
		return "VEX resolution is available, but no current finding is resolved by an active statement."
	default:
		return "No active VEX posture signal is available for this scope."
	}
}

func badgeStateForPercent(percent int) string {
	switch {
	case percent >= 90:
		return TrustBadgeStateVerified
	case percent >= 60:
		return TrustBadgeStatePartial
	case percent > 0:
		return TrustBadgeStateGap
	default:
		return TrustBadgeStateUnknown
	}
}

func badgeStateForMetricScore(score int, grade string) string {
	switch grade {
	case TrustGradeA, TrustGradeB:
		return TrustBadgeStateVerified
	case TrustGradeC:
		return TrustBadgeStatePartial
	case TrustGradeD:
		return TrustBadgeStateGap
	default:
		if score > 0 {
			return TrustBadgeStateGap
		}
		return TrustBadgeStateUnknown
	}
}

func mappingStatus(metricStatus string) string {
	switch normalizeMetricStatus(metricStatus) {
	case TrustMetricStatusVerified:
		return "implemented"
	case TrustMetricStatusPartial:
		return "partial"
	case TrustMetricStatusGap:
		return "gap"
	default:
		return "not_evaluated"
	}
}

func minMetricStatus(left, right string) string {
	statuses := []string{normalizeMetricStatus(left), normalizeMetricStatus(right)}
	sort.Slice(statuses, func(i, j int) bool {
		return metricStatusRank(statuses[i]) < metricStatusRank(statuses[j])
	})
	return statuses[0]
}

func metricStatusRank(status string) int {
	switch normalizeMetricStatus(status) {
	case TrustMetricStatusUnknown:
		return 0
	case TrustMetricStatusGap:
		return 1
	case TrustMetricStatusPartial:
		return 2
	case TrustMetricStatusVerified:
		return 3
	default:
		return 0
	}
}

func findingSeverityRank(value string) int {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "critical":
		return 5
	case "high":
		return 4
	case "medium":
		return 3
	case "low":
		return 2
	default:
		return 1
	}
}

func maxInt64(value, fallback int64) int64 {
	if value > 0 {
		return value
	}
	return fallback
}
