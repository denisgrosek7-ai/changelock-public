package compliance

import (
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/connectors"
	"github.com/denisgrosek/changelock/internal/handoff"
	"github.com/denisgrosek/changelock/internal/workflow"
)

const (
	ComplianceMappingSchemaVersion = "4.enterprise_compliance_mapping.v1"
	EvidenceVaultSchemaVersion     = "4.enterprise_evidence_vault.v1"
	PolicyDriftSchemaVersion       = "4.enterprise_policy_drift.v1"
	ExecutiveReportSchemaVersion   = "4.enterprise_executive_report.v1"

	CoverageFull     = "full"
	CoveragePartial  = "partial"
	CoverageInferred = "inferred"
	CoverageMissing  = "missing"

	DriftStateStable       = "stable"
	DriftStateSoftened     = "policy_softened"
	DriftStateStrengthened = "policy_strengthened"
	DriftStateUnderReview  = "under_review"
)

type EvidenceVaultRecord struct {
	SchemaVersion   string   `json:"schema_version"`
	VaultID         string   `json:"vault_id"`
	CustodyState    string   `json:"custody_state"`
	RetentionClass  string   `json:"retention_class"`
	TimestampState  string   `json:"timestamp_state"`
	ChainContinuity string   `json:"chain_continuity"`
	EvidenceRefs    []string `json:"evidence_refs,omitempty"`
}

type MappingInput struct {
	SubjectRef         string   `json:"subject_ref,omitempty"`
	ControlFamily      string   `json:"control_family,omitempty"`
	ControlID          string   `json:"control_id,omitempty"`
	CoverageState      string   `json:"coverage_state,omitempty"`
	FreshnessState     string   `json:"freshness_state,omitempty"`
	CustodyState       string   `json:"custody_state,omitempty"`
	RetentionClass     string   `json:"retention_class,omitempty"`
	TimestampState     string   `json:"timestamp_state,omitempty"`
	ChainContinuity    string   `json:"chain_continuity,omitempty"`
	EvidenceRefs       []string `json:"evidence_refs,omitempty"`
	TechnicalEventRefs []string `json:"technical_event_refs,omitempty"`
	ReasonCodes        []string `json:"reason_codes,omitempty"`
}

type ComplianceMappingRecord struct {
	SchemaVersion      string              `json:"schema_version"`
	SubjectRef         string              `json:"subject_ref,omitempty"`
	ControlFamily      string              `json:"control_family,omitempty"`
	ControlID          string              `json:"control_id,omitempty"`
	CurrentState       string              `json:"current_state"`
	CoverageState      string              `json:"coverage_state"`
	FreshnessState     string              `json:"freshness_state,omitempty"`
	EvidenceVault      EvidenceVaultRecord `json:"evidence_vault"`
	EvidenceRefs       []string            `json:"evidence_refs,omitempty"`
	TechnicalEventRefs []string            `json:"technical_event_refs,omitempty"`
	ReasonCodes        []string            `json:"reason_codes,omitempty"`
	ObservedAt         time.Time           `json:"observed_at"`
	Limitations        []string            `json:"limitations,omitempty"`
}

type DriftInput struct {
	SubjectRef       string   `json:"subject_ref,omitempty"`
	Actor            string   `json:"actor,omitempty"`
	PreviousMode     string   `json:"previous_mode,omitempty"`
	CurrentMode      string   `json:"current_mode,omitempty"`
	ChangeReason     string   `json:"change_reason,omitempty"`
	ExceptionID      string   `json:"exception_id,omitempty"`
	ImpactedControls []string `json:"impacted_controls,omitempty"`
	EvidenceRefs     []string `json:"evidence_refs,omitempty"`
}

type PolicyDriftRecord struct {
	SchemaVersion    string    `json:"schema_version"`
	SubjectRef       string    `json:"subject_ref,omitempty"`
	Actor            string    `json:"actor,omitempty"`
	CurrentState     string    `json:"current_state"`
	PreviousMode     string    `json:"previous_mode,omitempty"`
	CurrentMode      string    `json:"current_mode,omitempty"`
	ChangeReason     string    `json:"change_reason,omitempty"`
	ExceptionID      string    `json:"exception_id,omitempty"`
	ImpactedControls []string  `json:"impacted_controls,omitempty"`
	ImpactSummary    []string  `json:"impact_summary,omitempty"`
	EvidenceRefs     []string  `json:"evidence_refs,omitempty"`
	IdentityTrail    []string  `json:"identity_trail,omitempty"`
	ObservedAt       time.Time `json:"observed_at"`
	Limitations      []string  `json:"limitations,omitempty"`
}

type ExecutiveReportInput struct {
	ScopeRef                string                            `json:"scope_ref,omitempty"`
	WorkflowArtifacts       []workflow.LifecycleRecord        `json:"workflow_artifacts,omitempty"`
	ReconciliationArtifacts []connectors.ReconciliationRecord `json:"reconciliation_artifacts,omitempty"`
	PartnerArtifacts        []handoff.IntakeRecord            `json:"partner_artifacts,omitempty"`
	ComplianceArtifacts     []ComplianceMappingRecord         `json:"compliance_artifacts,omitempty"`
	DriftArtifacts          []PolicyDriftRecord               `json:"drift_artifacts,omitempty"`
}

type ExecutiveReport struct {
	SchemaVersion          string         `json:"schema_version"`
	ScopeRef               string         `json:"scope_ref,omitempty"`
	CurrentState           string         `json:"current_state"`
	WorkflowSummary        map[string]int `json:"workflow_summary,omitempty"`
	ConnectorSummary       map[string]int `json:"connector_summary,omitempty"`
	PartnerSummary         map[string]int `json:"partner_summary,omitempty"`
	ControlCoverageSummary map[string]int `json:"control_coverage_summary,omitempty"`
	DriftSummary           map[string]int `json:"drift_summary,omitempty"`
	Highlights             []string       `json:"highlights,omitempty"`
	EvidenceTraceRefs      []string       `json:"evidence_trace_refs,omitempty"`
	GeneratedAt            time.Time      `json:"generated_at"`
	Limitations            []string       `json:"limitations,omitempty"`
}

func EvaluateComplianceMapping(input MappingInput, now func() time.Time) ComplianceMappingRecord {
	if now == nil {
		now = time.Now
	}
	input.SubjectRef = strings.TrimSpace(input.SubjectRef)
	input.ControlFamily = strings.TrimSpace(input.ControlFamily)
	input.ControlID = strings.TrimSpace(input.ControlID)
	input.CoverageState = normalizeCoverageState(input.CoverageState)
	input.FreshnessState = firstNonEmpty(strings.TrimSpace(input.FreshnessState), "fresh")
	input.CustodyState = firstNonEmpty(strings.TrimSpace(input.CustodyState), "sealed_custody")
	input.RetentionClass = firstNonEmpty(strings.TrimSpace(input.RetentionClass), "audit_long_term")
	input.TimestampState = firstNonEmpty(strings.TrimSpace(input.TimestampState), "timestamped")
	input.ChainContinuity = firstNonEmpty(strings.TrimSpace(input.ChainContinuity), "continuous")

	state := "control_mapped"
	reasons := append([]string{}, input.ReasonCodes...)
	switch input.CoverageState {
	case CoverageMissing:
		state = "control_gap_active"
		reasons = append(reasons, "coverage_missing")
	case CoverageInferred:
		state = "control_partially_supported"
		reasons = append(reasons, "coverage_inferred_from_indirect_evidence")
	case CoveragePartial:
		state = "control_partially_supported"
		reasons = append(reasons, "coverage_partial")
	default:
		reasons = append(reasons, "coverage_supported_by_evidence")
	}

	return ComplianceMappingRecord{
		SchemaVersion:  ComplianceMappingSchemaVersion,
		SubjectRef:     input.SubjectRef,
		ControlFamily:  input.ControlFamily,
		ControlID:      input.ControlID,
		CurrentState:   state,
		CoverageState:  input.CoverageState,
		FreshnessState: input.FreshnessState,
		EvidenceVault: EvidenceVaultRecord{
			SchemaVersion:   EvidenceVaultSchemaVersion,
			VaultID:         "vault-" + now().UTC().Format("20060102150405"),
			CustodyState:    input.CustodyState,
			RetentionClass:  input.RetentionClass,
			TimestampState:  input.TimestampState,
			ChainContinuity: input.ChainContinuity,
			EvidenceRefs:    uniqueStrings(input.EvidenceRefs),
		},
		EvidenceRefs:       uniqueStrings(input.EvidenceRefs),
		TechnicalEventRefs: uniqueStrings(input.TechnicalEventRefs),
		ReasonCodes:        uniqueStrings(reasons),
		ObservedAt:         now().UTC(),
		Limitations: []string{
			"Control mapping distinguishes full, partial, inferred, and missing coverage and does not claim stronger assurance than the linked evidence supports.",
			"Evidence vault posture remains custody-aware and traceable to the underlying technical event chain.",
		},
	}
}

func EvaluatePolicyDrift(input DriftInput, now func() time.Time) PolicyDriftRecord {
	if now == nil {
		now = time.Now
	}
	input.SubjectRef = strings.TrimSpace(input.SubjectRef)
	input.Actor = strings.TrimSpace(input.Actor)
	input.PreviousMode = strings.TrimSpace(input.PreviousMode)
	input.CurrentMode = strings.TrimSpace(input.CurrentMode)
	input.ChangeReason = strings.TrimSpace(input.ChangeReason)
	input.ExceptionID = strings.TrimSpace(input.ExceptionID)

	state := DriftStateStable
	impact := []string{}
	switch {
	case softensPolicy(input.PreviousMode, input.CurrentMode):
		state = DriftStateSoftened
		impact = append(impact, "policy_softening_requires_executive_review")
	case strengthensPolicy(input.PreviousMode, input.CurrentMode):
		state = DriftStateStrengthened
		impact = append(impact, "policy_strengthening_recorded")
	default:
		impact = append(impact, "policy_state_under_review")
	}
	if input.ExceptionID != "" {
		impact = append(impact, "exception_linked_to_drift")
	}
	if input.Actor == "" {
		impact = append(impact, "identity_trail_missing")
	}

	return PolicyDriftRecord{
		SchemaVersion:    PolicyDriftSchemaVersion,
		SubjectRef:       input.SubjectRef,
		Actor:            input.Actor,
		CurrentState:     state,
		PreviousMode:     input.PreviousMode,
		CurrentMode:      input.CurrentMode,
		ChangeReason:     input.ChangeReason,
		ExceptionID:      input.ExceptionID,
		ImpactedControls: uniqueStrings(input.ImpactedControls),
		ImpactSummary:    uniqueStrings(impact),
		EvidenceRefs:     uniqueStrings(input.EvidenceRefs),
		IdentityTrail:    uniqueStrings([]string{input.Actor, input.ExceptionID}),
		ObservedAt:       now().UTC(),
		Limitations: []string{
			"Policy drift reporting records who changed policy posture and why, but does not replace the underlying policy bundle history.",
			"Executive drift summaries remain traceable to exception and evidence refs rather than abstract score-only outputs.",
		},
	}
}

func BuildExecutiveReport(input ExecutiveReportInput, now func() time.Time) ExecutiveReport {
	if now == nil {
		now = time.Now
	}
	workflowSummary := map[string]int{}
	connectorSummary := map[string]int{}
	partnerSummary := map[string]int{}
	controlSummary := map[string]int{}
	driftSummary := map[string]int{}
	evidenceRefs := []string{}
	highlights := []string{}

	for _, item := range input.WorkflowArtifacts {
		workflowSummary[item.CurrentState]++
		evidenceRefs = append(evidenceRefs, item.EvidenceRefs...)
		if !item.ClosureReady && item.RequestedState == workflow.StateResolved {
			highlights = append(highlights, "workflow closures remain blocked pending validation evidence")
		}
	}
	for _, item := range input.ReconciliationArtifacts {
		connectorSummary[item.CurrentState]++
		evidenceRefs = append(evidenceRefs, item.EvidenceRefs...)
		if item.CurrentState == connectors.StateConnectorDegraded {
			highlights = append(highlights, "connector health degraded but canonical workflow truth stayed local")
		}
	}
	for _, item := range input.PartnerArtifacts {
		partnerSummary[item.CurrentState]++
		evidenceRefs = append(evidenceRefs, item.EvidenceRefs...)
		if item.CurrentState == handoff.IntakeStateRejected || item.CurrentState == handoff.IntakeStateExpired {
			highlights = append(highlights, "partner trust posture requires review or refresh")
		}
	}
	for _, item := range input.ComplianceArtifacts {
		controlSummary[item.CoverageState]++
		evidenceRefs = append(evidenceRefs, item.EvidenceRefs...)
	}
	for _, item := range input.DriftArtifacts {
		driftSummary[item.CurrentState]++
		evidenceRefs = append(evidenceRefs, item.EvidenceRefs...)
		if item.CurrentState == DriftStateSoftened {
			highlights = append(highlights, "policy softening pressure is active and tied to named evidence")
		}
	}

	currentState := "executive_governance_ready"
	if controlSummary[CoverageMissing] > 0 || driftSummary[DriftStateSoftened] > 0 {
		currentState = "executive_governance_attention_required"
	}

	return ExecutiveReport{
		SchemaVersion:          ExecutiveReportSchemaVersion,
		ScopeRef:               strings.TrimSpace(input.ScopeRef),
		CurrentState:           currentState,
		WorkflowSummary:        workflowSummary,
		ConnectorSummary:       connectorSummary,
		PartnerSummary:         partnerSummary,
		ControlCoverageSummary: controlSummary,
		DriftSummary:           driftSummary,
		Highlights:             uniqueStrings(highlights),
		EvidenceTraceRefs:      uniqueStrings(evidenceRefs),
		GeneratedAt:            now().UTC(),
		Limitations: []string{
			"Executive reporting is a bounded summary layer and remains traceable to the underlying evidence refs.",
			"Compliance posture is derived from technical and workflow evidence and does not assert coverage beyond that evidence.",
		},
	}
}

func normalizeCoverageState(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case CoveragePartial:
		return CoveragePartial
	case CoverageInferred:
		return CoverageInferred
	case CoverageMissing:
		return CoverageMissing
	default:
		return CoverageFull
	}
}

func softensPolicy(previousMode, currentMode string) bool {
	previousMode = strings.ToLower(strings.TrimSpace(previousMode))
	currentMode = strings.ToLower(strings.TrimSpace(currentMode))
	if previousMode == "" || currentMode == "" {
		return false
	}
	return (previousMode == "deny" && currentMode != "deny") || (previousMode == "enforce" && (currentMode == "monitor" || currentMode == "exception"))
}

func strengthensPolicy(previousMode, currentMode string) bool {
	previousMode = strings.ToLower(strings.TrimSpace(previousMode))
	currentMode = strings.ToLower(strings.TrimSpace(currentMode))
	if previousMode == "" || currentMode == "" {
		return false
	}
	return (previousMode == "monitor" && currentMode == "enforce") || (previousMode == "exception" && currentMode == "deny")
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		items = append(items, value)
	}
	if len(items) == 0 {
		return nil
	}
	return items
}
