package workflow

import (
	"strings"
	"time"
)

const (
	LifecycleSchemaVersion = "4.enterprise_workflow_lifecycle.v1"
	FeedbackSchemaVersion  = "4.enterprise_feedback_signal.v1"

	StateDiscovered      = "discovered"
	StateTriaged         = "triaged"
	StateAssigned        = "assigned"
	StateUnderValidation = "under_validation"
	StateValidated       = "validated"
	StateRejected        = "rejected"
	StateExceptionActive = "exception_active"
	StateResolved        = "resolved"
	StateReopened        = "reopened"

	ApprovalStateAuto     = "auto_safe"
	ApprovalStatePending  = "pending_review"
	ApprovalStateApproved = "approved"
	ApprovalStateRejected = "rejected"

	ValidationStatePending  = "pending"
	ValidationStateVerified = "verified"
	ValidationStateFailed   = "failed"
	ValidationStateSkipped  = "not_required"
)

type Ownership struct {
	FindingOwner     string `json:"finding_owner,omitempty"`
	RemediationOwner string `json:"remediation_owner,omitempty"`
	Approver         string `json:"approver,omitempty"`
	PartnerOwner     string `json:"partner_owner,omitempty"`
	ComplianceOwner  string `json:"compliance_owner,omitempty"`
}

type RoutingDecision struct {
	PrimaryOwner   string   `json:"primary_owner,omitempty"`
	RouteClass     string   `json:"route_class,omitempty"`
	AlertTargets   []string `json:"alert_targets,omitempty"`
	EscalationPath []string `json:"escalation_path,omitempty"`
}

type FeedbackSignal struct {
	SchemaVersion string    `json:"schema_version"`
	SourceSystem  string    `json:"source_system,omitempty"`
	Actor         string    `json:"actor,omitempty"`
	Decision      string    `json:"decision,omitempty"`
	Reason        string    `json:"reason,omitempty"`
	CapturedAt    time.Time `json:"captured_at"`
}

type LifecycleInput struct {
	WorkflowID         string          `json:"workflow_id,omitempty"`
	ArtifactType       string          `json:"artifact_type,omitempty"`
	SubjectRef         string          `json:"subject_ref,omitempty"`
	Severity           string          `json:"severity,omitempty"`
	RequestedState     string          `json:"requested_state,omitempty"`
	Owners             Ownership       `json:"owners"`
	TicketID           string          `json:"ticket_id,omitempty"`
	ValidationRequired bool            `json:"validation_required,omitempty"`
	ValidationState    string          `json:"validation_state,omitempty"`
	ExceptionActive    bool            `json:"exception_active,omitempty"`
	ApprovalRequired   bool            `json:"approval_required,omitempty"`
	ApprovalState      string          `json:"approval_state,omitempty"`
	Feedback           *FeedbackSignal `json:"feedback,omitempty"`
	EvidenceRefs       []string        `json:"evidence_refs,omitempty"`
	ReasonCodes        []string        `json:"reason_codes,omitempty"`
}

type LifecycleRecord struct {
	SchemaVersion      string          `json:"schema_version"`
	WorkflowID         string          `json:"workflow_id"`
	ArtifactType       string          `json:"artifact_type,omitempty"`
	SubjectRef         string          `json:"subject_ref,omitempty"`
	RequestedState     string          `json:"requested_state,omitempty"`
	CurrentState       string          `json:"current_state"`
	CanonicalState     string          `json:"canonical_state"`
	Owners             Ownership       `json:"owners"`
	Routing            RoutingDecision `json:"routing"`
	TicketID           string          `json:"ticket_id,omitempty"`
	ValidationRequired bool            `json:"validation_required"`
	ValidationState    string          `json:"validation_state,omitempty"`
	ApprovalRequired   bool            `json:"approval_required"`
	ApprovalState      string          `json:"approval_state,omitempty"`
	ExceptionActive    bool            `json:"exception_active"`
	ClosureReady       bool            `json:"closure_ready"`
	Feedback           *FeedbackSignal `json:"feedback,omitempty"`
	EvidenceRefs       []string        `json:"evidence_refs,omitempty"`
	ReasonCodes        []string        `json:"reason_codes,omitempty"`
	ObservedAt         time.Time       `json:"observed_at"`
	Limitations        []string        `json:"limitations,omitempty"`
}

func EvaluateLifecycle(input LifecycleInput, now func() time.Time) LifecycleRecord {
	if now == nil {
		now = time.Now
	}
	input.WorkflowID = strings.TrimSpace(input.WorkflowID)
	if input.WorkflowID == "" {
		input.WorkflowID = "wf-" + now().UTC().Format("20060102150405")
	}
	input.ArtifactType = strings.TrimSpace(input.ArtifactType)
	input.SubjectRef = strings.TrimSpace(input.SubjectRef)
	input.Severity = strings.ToLower(strings.TrimSpace(input.Severity))
	input.RequestedState = normalizeState(input.RequestedState)
	if input.RequestedState == "" {
		input.RequestedState = StateDiscovered
	}
	if input.ValidationRequired && strings.TrimSpace(input.ValidationState) == "" {
		input.ValidationState = ValidationStatePending
	}
	if !input.ValidationRequired {
		input.ValidationState = ValidationStateSkipped
	}
	input.ApprovalState = normalizeApprovalState(input.ApprovalState, input.ApprovalRequired)
	input.Owners = normalizeOwners(input.Owners)
	if input.Feedback != nil {
		input.Feedback.SchemaVersion = FeedbackSchemaVersion
		if input.Feedback.CapturedAt.IsZero() {
			input.Feedback.CapturedAt = now().UTC()
		}
		input.Feedback.SourceSystem = strings.TrimSpace(input.Feedback.SourceSystem)
		input.Feedback.Actor = strings.TrimSpace(input.Feedback.Actor)
		input.Feedback.Decision = strings.TrimSpace(input.Feedback.Decision)
		input.Feedback.Reason = strings.TrimSpace(input.Feedback.Reason)
	}

	reasons := append([]string{}, input.ReasonCodes...)
	currentState := input.RequestedState
	canonicalState := input.RequestedState

	switch {
	case input.ExceptionActive:
		currentState = StateExceptionActive
		canonicalState = StateExceptionActive
		reasons = append(reasons, "exception_active_requires_governance_tracking")
	case input.ApprovalRequired && input.ApprovalState == ApprovalStateRejected:
		currentState = StateRejected
		canonicalState = StateRejected
		reasons = append(reasons, "approval_rejected")
	case input.ApprovalRequired && input.ApprovalState == ApprovalStatePending && wantsClosure(input.RequestedState):
		currentState = StateAssigned
		canonicalState = StateAssigned
		reasons = append(reasons, "approval_pending_before_closure")
	case input.ValidationRequired && wantsClosure(input.RequestedState) && input.ValidationState != ValidationStateVerified:
		currentState = StateUnderValidation
		canonicalState = StateUnderValidation
		reasons = append(reasons, "validation_required_before_resolution")
	}
	if input.Feedback != nil {
		if input.Feedback.Actor == "" {
			reasons = append(reasons, "feedback_missing_identity_trail")
			if strings.EqualFold(input.Feedback.Decision, "reject") {
				reasons = append(reasons, "feedback_reject_recorded_without_authority")
			}
		} else if strings.EqualFold(input.Feedback.Decision, "reject") {
			currentState = StateRejected
			canonicalState = StateRejected
			reasons = append(reasons, "feedback_rejected_recommendation")
		}
	}

	routing := buildRoutingDecision(input.Severity, input.Owners)
	if routing.PrimaryOwner == "" {
		reasons = append(reasons, "owner_mapping_missing")
	}

	closureReady := currentState == StateResolved && (!input.ValidationRequired || input.ValidationState == ValidationStateVerified)
	if currentState == StateValidated && input.ValidationState == ValidationStateVerified {
		reasons = append(reasons, "validation_evidence_ready_for_closure")
	}

	return LifecycleRecord{
		SchemaVersion:      LifecycleSchemaVersion,
		WorkflowID:         input.WorkflowID,
		ArtifactType:       input.ArtifactType,
		SubjectRef:         input.SubjectRef,
		RequestedState:     input.RequestedState,
		CurrentState:       currentState,
		CanonicalState:     canonicalState,
		Owners:             input.Owners,
		Routing:            routing,
		TicketID:           strings.TrimSpace(input.TicketID),
		ValidationRequired: input.ValidationRequired,
		ValidationState:    strings.TrimSpace(input.ValidationState),
		ApprovalRequired:   input.ApprovalRequired,
		ApprovalState:      input.ApprovalState,
		ExceptionActive:    input.ExceptionActive,
		ClosureReady:       closureReady,
		Feedback:           input.Feedback,
		EvidenceRefs:       uniqueStrings(input.EvidenceRefs),
		ReasonCodes:        uniqueStrings(reasons),
		ObservedAt:         now().UTC(),
		Limitations: []string{
			"Workflow state remains canonical only when backed by validation, approval, and evidence discipline; external closure alone is not technical truth.",
			"Feedback signals remain audited workflow inputs and do not bypass approval or validation gates.",
		},
	}
}

func buildRoutingDecision(severity string, owners Ownership) RoutingDecision {
	primary := firstNonEmpty(
		owners.RemediationOwner,
		owners.FindingOwner,
		owners.ComplianceOwner,
		owners.PartnerOwner,
		owners.Approver,
	)
	routeClass := "owner_routed"
	if primary == "" {
		routeClass = "unassigned"
	}
	escalation := []string{}
	if owners.Approver != "" && owners.Approver != primary {
		escalation = append(escalation, owners.Approver)
	}
	if owners.ComplianceOwner != "" && owners.ComplianceOwner != primary && owners.ComplianceOwner != owners.Approver {
		escalation = append(escalation, owners.ComplianceOwner)
	}
	if owners.PartnerOwner != "" && owners.PartnerOwner != primary && owners.PartnerOwner != owners.Approver {
		escalation = append(escalation, owners.PartnerOwner)
	}
	targets := []string{}
	if primary != "" {
		targets = append(targets, primary)
	}
	targets = append(targets, escalation...)
	if severity == "critical" || severity == "high" {
		routeClass = "evidence_aware_escalation"
	}
	return RoutingDecision{
		PrimaryOwner:   primary,
		RouteClass:     routeClass,
		AlertTargets:   uniqueStrings(targets),
		EscalationPath: uniqueStrings(escalation),
	}
}

func normalizeOwners(owners Ownership) Ownership {
	owners.FindingOwner = strings.TrimSpace(owners.FindingOwner)
	owners.RemediationOwner = strings.TrimSpace(owners.RemediationOwner)
	owners.Approver = strings.TrimSpace(owners.Approver)
	owners.PartnerOwner = strings.TrimSpace(owners.PartnerOwner)
	owners.ComplianceOwner = strings.TrimSpace(owners.ComplianceOwner)
	if owners.RemediationOwner == "" {
		owners.RemediationOwner = owners.FindingOwner
	}
	if owners.Approver == "" {
		owners.Approver = firstNonEmpty(owners.ComplianceOwner, owners.FindingOwner)
	}
	return owners
}

func normalizeState(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case StateDiscovered:
		return StateDiscovered
	case StateTriaged:
		return StateTriaged
	case StateAssigned:
		return StateAssigned
	case StateUnderValidation:
		return StateUnderValidation
	case StateValidated:
		return StateValidated
	case StateRejected:
		return StateRejected
	case StateExceptionActive:
		return StateExceptionActive
	case StateResolved:
		return StateResolved
	case StateReopened:
		return StateReopened
	default:
		return ""
	}
}

func normalizeApprovalState(value string, approvalRequired bool) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case ApprovalStatePending:
		return ApprovalStatePending
	case ApprovalStateApproved:
		return ApprovalStateApproved
	case ApprovalStateRejected:
		return ApprovalStateRejected
	case ApprovalStateAuto:
		return ApprovalStateAuto
	default:
		if approvalRequired {
			return ApprovalStatePending
		}
		return ApprovalStateAuto
	}
}

func wantsClosure(state string) bool {
	return state == StateResolved || state == StateValidated
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
