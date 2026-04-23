package claims

import "strings"

const (
	MeasuredPublicProofValDReleaseIssuanceStateActive     = "measured_public_proof_vald_release_issuance_active"
	MeasuredPublicProofValDReleaseIssuanceStatePartial    = "measured_public_proof_vald_release_issuance_partial"
	MeasuredPublicProofValDReleaseIssuanceStateIncomplete = "measured_public_proof_vald_release_issuance_incomplete"

	MeasuredPublicProofValDClaimLifecycleStateActive     = "measured_public_proof_vald_claim_lifecycle_active"
	MeasuredPublicProofValDClaimLifecycleStatePartial    = "measured_public_proof_vald_claim_lifecycle_partial"
	MeasuredPublicProofValDClaimLifecycleStateIncomplete = "measured_public_proof_vald_claim_lifecycle_incomplete"

	MeasuredPublicProofValDPublicationDecisionStateActive     = "measured_public_proof_vald_publication_decision_active"
	MeasuredPublicProofValDPublicationDecisionStatePartial    = "measured_public_proof_vald_publication_decision_partial"
	MeasuredPublicProofValDPublicationDecisionStateIncomplete = "measured_public_proof_vald_publication_decision_incomplete"

	MeasuredPublicProofValDCorrectionWorkflowStateActive     = "measured_public_proof_vald_correction_workflow_active"
	MeasuredPublicProofValDCorrectionWorkflowStatePartial    = "measured_public_proof_vald_correction_workflow_partial"
	MeasuredPublicProofValDCorrectionWorkflowStateIncomplete = "measured_public_proof_vald_correction_workflow_incomplete"

	MeasuredPublicProofValDStateIncomplete  = "measured_public_proof_vald_incomplete"
	MeasuredPublicProofValDStateSubstantial = "measured_public_proof_vald_substantially_ready"
	MeasuredPublicProofValDStateActive      = "measured_public_proof_vald_active"
)

type PublicProofReleaseIssuanceItem struct {
	ClaimID             string   `json:"claim_id"`
	ArtifactID          string   `json:"artifact_id"`
	CurrentState        string   `json:"current_state"`
	ReleaseID           string   `json:"release_id"`
	BuildIdentity       string   `json:"build_identity"`
	ReleaseChannel      string   `json:"release_channel"`
	PriorReleaseRef     string   `json:"prior_release_ref"`
	ReissueDecision     string   `json:"reissue_decision"`
	PublicationDecision string   `json:"publication_decision"`
	RequiredChecks      []string `json:"required_checks,omitempty"`
	SatisfiedChecks     []string `json:"satisfied_checks,omitempty"`
	VerificationRefs    []string `json:"verification_refs,omitempty"`
	AuditRefs           []string `json:"audit_refs,omitempty"`
	FailureStates       []string `json:"failure_states,omitempty"`
	Limitations         []string `json:"limitations,omitempty"`
}

type PublicProofClaimLifecycleItem struct {
	ClaimID                  string   `json:"claim_id"`
	ArtifactID               string   `json:"artifact_id"`
	CurrentState             string   `json:"current_state"`
	ClaimStatus              string   `json:"claim_status"`
	ReissueState             string   `json:"reissue_state"`
	FreshnessState           string   `json:"freshness_state"`
	PublicationScope         string   `json:"publication_scope"`
	RestrictionState         string   `json:"restriction_state"`
	WithdrawalState          string   `json:"withdrawal_state"`
	SupersessionState        string   `json:"supersession_state"`
	SupportedLifecycleStates []string `json:"supported_lifecycle_states,omitempty"`
	VerifierNoticeRefs       []string `json:"verifier_notice_refs,omitempty"`
	PortalRefs               []string `json:"portal_refs,omitempty"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type PublicProofPublicationDecisionItem struct {
	ClaimID           string   `json:"claim_id"`
	ArtifactID        string   `json:"artifact_id"`
	CurrentState      string   `json:"current_state"`
	PublicationStatus string   `json:"publication_status"`
	ApprovalBoundary  string   `json:"approval_boundary"`
	RedactionTier     string   `json:"redaction_tier"`
	PublicationScope  string   `json:"publication_scope"`
	AutomationState   string   `json:"automation_state"`
	OverridePermitted bool     `json:"override_permitted"`
	DecisionAuditRefs []string `json:"decision_audit_refs,omitempty"`
	ProjectionRefs    []string `json:"projection_refs,omitempty"`
	FailureStates     []string `json:"failure_states,omitempty"`
	Limitations       []string `json:"limitations,omitempty"`
}

type PublicProofCorrectionWorkflowItem struct {
	ClaimID               string   `json:"claim_id"`
	ArtifactID            string   `json:"artifact_id"`
	CurrentState          string   `json:"current_state"`
	TriggerClass          string   `json:"trigger_class"`
	TriggerState          string   `json:"trigger_state"`
	RestrictionActionRef  string   `json:"restriction_action_ref"`
	WithdrawalActionRef   string   `json:"withdrawal_action_ref"`
	SupersessionActionRef string   `json:"supersession_action_ref"`
	CorrectionNoticeRef   string   `json:"correction_notice_ref"`
	AuditRefs             []string `json:"audit_refs,omitempty"`
	VerifierRefs          []string `json:"verifier_refs,omitempty"`
	PortalNoticeRefs      []string `json:"portal_notice_refs,omitempty"`
	FailureStates         []string `json:"failure_states,omitempty"`
	Limitations           []string `json:"limitations,omitempty"`
}

func EvaluateMeasuredPublicProofValDReleaseIssuanceState(items []PublicProofReleaseIssuanceItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValDReleaseIssuanceStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReleaseID) == "" {
			return MeasuredPublicProofValDReleaseIssuanceStateIncomplete
		}
		if strings.TrimSpace(item.BuildIdentity) == "" || strings.TrimSpace(item.ReleaseChannel) == "" || strings.TrimSpace(item.PriorReleaseRef) == "" || strings.TrimSpace(item.ReissueDecision) == "" || strings.TrimSpace(item.PublicationDecision) == "" {
			return MeasuredPublicProofValDReleaseIssuanceStatePartial
		}
		if len(item.RequiredChecks) == 0 || len(item.SatisfiedChecks) == 0 || len(item.VerificationRefs) == 0 || len(item.AuditRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValDReleaseIssuanceStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "issuance_gate_ready" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValDReleaseIssuanceStatePartial
	}
	return MeasuredPublicProofValDReleaseIssuanceStateActive
}

func EvaluateMeasuredPublicProofValDClaimLifecycleState(items []PublicProofClaimLifecycleItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValDClaimLifecycleStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ClaimStatus) == "" {
			return MeasuredPublicProofValDClaimLifecycleStateIncomplete
		}
		if strings.TrimSpace(item.ReissueState) == "" || strings.TrimSpace(item.FreshnessState) == "" || strings.TrimSpace(item.PublicationScope) == "" || strings.TrimSpace(item.RestrictionState) == "" || strings.TrimSpace(item.WithdrawalState) == "" || strings.TrimSpace(item.SupersessionState) == "" {
			return MeasuredPublicProofValDClaimLifecycleStatePartial
		}
		if len(item.SupportedLifecycleStates) == 0 || len(item.VerifierNoticeRefs) == 0 || len(item.PortalRefs) == 0 || len(item.EvidenceRefs) == 0 {
			return MeasuredPublicProofValDClaimLifecycleStatePartial
		}
		if !containsTrimmedString(item.SupportedLifecycleStates, PublicProofStatusRestricted) ||
			!containsTrimmedString(item.SupportedLifecycleStates, PublicProofStatusSuperseded) ||
			!containsTrimmedString(item.SupportedLifecycleStates, PublicProofStatusWithdrawn) ||
			!containsTrimmedString(item.SupportedLifecycleStates, PublicProofStatusClaimNotReissued) {
			return MeasuredPublicProofValDClaimLifecycleStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "claim_lifecycle_governed" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValDClaimLifecycleStatePartial
	}
	return MeasuredPublicProofValDClaimLifecycleStateActive
}

func EvaluateMeasuredPublicProofValDPublicationDecisionState(items []PublicProofPublicationDecisionItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValDPublicationDecisionStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.PublicationStatus) == "" {
			return MeasuredPublicProofValDPublicationDecisionStateIncomplete
		}
		if strings.TrimSpace(item.ApprovalBoundary) == "" || strings.TrimSpace(item.RedactionTier) == "" || strings.TrimSpace(item.PublicationScope) == "" || strings.TrimSpace(item.AutomationState) == "" {
			return MeasuredPublicProofValDPublicationDecisionStatePartial
		}
		if len(item.DecisionAuditRefs) == 0 || len(item.ProjectionRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValDPublicationDecisionStatePartial
		}
		if item.OverridePermitted || strings.TrimSpace(item.CurrentState) != "publication_decision_ready" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValDPublicationDecisionStatePartial
	}
	return MeasuredPublicProofValDPublicationDecisionStateActive
}

func EvaluateMeasuredPublicProofValDCorrectionWorkflowState(items []PublicProofCorrectionWorkflowItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValDCorrectionWorkflowStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.TriggerClass) == "" {
			return MeasuredPublicProofValDCorrectionWorkflowStateIncomplete
		}
		if strings.TrimSpace(item.TriggerState) == "" || strings.TrimSpace(item.RestrictionActionRef) == "" || strings.TrimSpace(item.WithdrawalActionRef) == "" || strings.TrimSpace(item.SupersessionActionRef) == "" || strings.TrimSpace(item.CorrectionNoticeRef) == "" {
			return MeasuredPublicProofValDCorrectionWorkflowStatePartial
		}
		if len(item.AuditRefs) == 0 || len(item.VerifierRefs) == 0 || len(item.PortalNoticeRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValDCorrectionWorkflowStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "correction_workflow_ready" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValDCorrectionWorkflowStatePartial
	}
	return MeasuredPublicProofValDCorrectionWorkflowStateActive
}

func EvaluateMeasuredPublicProofValDState(valCState, releaseIssuanceState, claimLifecycleState, publicationDecisionState, correctionWorkflowState string) string {
	if strings.TrimSpace(valCState) != MeasuredPublicProofValCStateActive {
		return MeasuredPublicProofValDStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(releaseIssuanceState),
		strings.TrimSpace(claimLifecycleState),
		strings.TrimSpace(publicationDecisionState),
		strings.TrimSpace(correctionWorkflowState),
	} {
		switch state {
		case MeasuredPublicProofValDReleaseIssuanceStateActive,
			MeasuredPublicProofValDClaimLifecycleStateActive,
			MeasuredPublicProofValDPublicationDecisionStateActive,
			MeasuredPublicProofValDCorrectionWorkflowStateActive:
		case MeasuredPublicProofValDReleaseIssuanceStatePartial,
			MeasuredPublicProofValDClaimLifecycleStatePartial,
			MeasuredPublicProofValDPublicationDecisionStatePartial,
			MeasuredPublicProofValDCorrectionWorkflowStatePartial:
			hasPartial = true
		default:
			return MeasuredPublicProofValDStateIncomplete
		}
	}
	if hasPartial {
		return MeasuredPublicProofValDStateSubstantial
	}
	return MeasuredPublicProofValDStateActive
}
