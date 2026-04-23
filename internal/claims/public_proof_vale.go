package claims

import "strings"

const (
	MeasuredPublicProofValEReplayCorrectnessReviewStateActive     = "measured_public_proof_vale_replay_correctness_review_active"
	MeasuredPublicProofValEReplayCorrectnessReviewStatePartial    = "measured_public_proof_vale_replay_correctness_review_partial"
	MeasuredPublicProofValEReplayCorrectnessReviewStateIncomplete = "measured_public_proof_vale_replay_correctness_review_incomplete"

	MeasuredPublicProofValESigningTrustReviewStateActive     = "measured_public_proof_vale_signing_trust_review_active"
	MeasuredPublicProofValESigningTrustReviewStatePartial    = "measured_public_proof_vale_signing_trust_review_partial"
	MeasuredPublicProofValESigningTrustReviewStateIncomplete = "measured_public_proof_vale_signing_trust_review_incomplete"

	MeasuredPublicProofValETransparencyReviewStateActive     = "measured_public_proof_vale_transparency_review_active"
	MeasuredPublicProofValETransparencyReviewStatePartial    = "measured_public_proof_vale_transparency_review_partial"
	MeasuredPublicProofValETransparencyReviewStateIncomplete = "measured_public_proof_vale_transparency_review_incomplete"

	MeasuredPublicProofValERedactionReviewStateActive     = "measured_public_proof_vale_redaction_review_active"
	MeasuredPublicProofValERedactionReviewStatePartial    = "measured_public_proof_vale_redaction_review_partial"
	MeasuredPublicProofValERedactionReviewStateIncomplete = "measured_public_proof_vale_redaction_review_incomplete"

	MeasuredPublicProofValECompatibilityReviewStateActive     = "measured_public_proof_vale_compatibility_review_active"
	MeasuredPublicProofValECompatibilityReviewStatePartial    = "measured_public_proof_vale_compatibility_review_partial"
	MeasuredPublicProofValECompatibilityReviewStateIncomplete = "measured_public_proof_vale_compatibility_review_incomplete"

	MeasuredPublicProofValEIssuanceReviewStateActive     = "measured_public_proof_vale_issuance_review_active"
	MeasuredPublicProofValEIssuanceReviewStatePartial    = "measured_public_proof_vale_issuance_review_partial"
	MeasuredPublicProofValEIssuanceReviewStateIncomplete = "measured_public_proof_vale_issuance_review_incomplete"

	MeasuredPublicProofValEFailureStateReviewStateActive     = "measured_public_proof_vale_failure_state_review_active"
	MeasuredPublicProofValEFailureStateReviewStatePartial    = "measured_public_proof_vale_failure_state_review_partial"
	MeasuredPublicProofValEFailureStateReviewStateIncomplete = "measured_public_proof_vale_failure_state_review_incomplete"

	MeasuredPublicProofValEStateIncomplete  = "measured_public_proof_vale_incomplete"
	MeasuredPublicProofValEStateSubstantial = "measured_public_proof_vale_substantially_ready"
	MeasuredPublicProofValEStateActive      = "measured_public_proof_vale_active"
)

type PublicProofReplayCorrectnessReviewItem struct {
	ClaimID                     string   `json:"claim_id"`
	ArtifactID                  string   `json:"artifact_id"`
	CurrentState                string   `json:"current_state"`
	ReviewOutcome               string   `json:"review_outcome"`
	ReplayState                 string   `json:"replay_state"`
	ComparisonMode              string   `json:"comparison_mode"`
	MethodologyRef              string   `json:"methodology_ref"`
	EvaluationRef               string   `json:"evaluation_ref"`
	ToleranceDecision           string   `json:"tolerance_decision"`
	SupportedEnvironmentClasses []string `json:"supported_environment_classes,omitempty"`
	ToleranceBands              []string `json:"tolerance_bands,omitempty"`
	UnsupportedReplayCases      []string `json:"unsupported_replay_cases,omitempty"`
	ReviewRefs                  []string `json:"review_refs,omitempty"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
}

type PublicProofSigningTrustReviewItem struct {
	ClaimID                     string   `json:"claim_id"`
	ArtifactID                  string   `json:"artifact_id"`
	CurrentState                string   `json:"current_state"`
	ReviewOutcome               string   `json:"review_outcome"`
	VerificationState           string   `json:"verification_state"`
	TrustRootState              string   `json:"trust_root_state"`
	SigningPurposeState         string   `json:"signing_purpose_state"`
	HistoricalVerificationState string   `json:"historical_verification_state"`
	KeyRotationState            string   `json:"key_rotation_state"`
	RevocationState             string   `json:"revocation_state"`
	TimestampState              string   `json:"timestamp_state"`
	SignerMode                  string   `json:"signer_mode"`
	TrustRootID                 string   `json:"trust_root_id"`
	KeyVersion                  string   `json:"key_version"`
	ReviewRefs                  []string `json:"review_refs,omitempty"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	FailureStates               []string `json:"failure_states,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
}

type PublicProofTransparencyReviewItem struct {
	ClaimID                string   `json:"claim_id"`
	ArtifactID             string   `json:"artifact_id"`
	CurrentState           string   `json:"current_state"`
	ReviewOutcome          string   `json:"review_outcome"`
	TransparencyState      string   `json:"transparency_state"`
	EntryHashState         string   `json:"entry_hash_state"`
	AnchorState            string   `json:"anchor_state"`
	SupersessionVisibility string   `json:"supersession_visibility"`
	AnchorRef              string   `json:"anchor_ref"`
	EntryID                string   `json:"entry_id"`
	EntryHash              string   `json:"entry_hash"`
	LineageRef             string   `json:"lineage_ref"`
	ReviewRefs             []string `json:"review_refs,omitempty"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	FailureStates          []string `json:"failure_states,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type PublicProofRedactionReviewItem struct {
	ClaimID               string   `json:"claim_id"`
	ArtifactID            string   `json:"artifact_id"`
	CurrentState          string   `json:"current_state"`
	ReviewOutcome         string   `json:"review_outcome"`
	RedactionTier         string   `json:"redaction_tier"`
	PublicationScope      string   `json:"publication_scope"`
	PortalProjectionState string   `json:"portal_projection_state"`
	RedactionDecision     string   `json:"redaction_decision"`
	ProjectionDiscipline  string   `json:"projection_discipline"`
	RemovedFields         []string `json:"removed_fields,omitempty"`
	NeverPublishedFields  []string `json:"never_published_fields,omitempty"`
	ReviewRefs            []string `json:"review_refs,omitempty"`
	EvidenceRefs          []string `json:"evidence_refs,omitempty"`
	FailureStates         []string `json:"failure_states,omitempty"`
	Limitations           []string `json:"limitations,omitempty"`
}

type PublicProofCompatibilityReviewItem struct {
	ClaimID                  string   `json:"claim_id"`
	ArtifactID               string   `json:"artifact_id"`
	CurrentState             string   `json:"current_state"`
	ReviewOutcome            string   `json:"review_outcome"`
	SchemaCompatibility      string   `json:"schema_compatibility"`
	VerifierCompatibility    string   `json:"verifier_compatibility"`
	DeprecationState         string   `json:"deprecation_state"`
	ReplayCompatibility      string   `json:"replay_compatibility"`
	MethodologyCompatibility string   `json:"methodology_compatibility"`
	SupportedSchemaLines     []string `json:"supported_schema_lines,omitempty"`
	UnsupportedCases         []string `json:"unsupported_cases,omitempty"`
	ReviewRefs               []string `json:"review_refs,omitempty"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	FailureStates            []string `json:"failure_states,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type PublicProofIssuanceReviewItem struct {
	ClaimID                 string   `json:"claim_id"`
	ArtifactID              string   `json:"artifact_id"`
	CurrentState            string   `json:"current_state"`
	ReviewOutcome           string   `json:"review_outcome"`
	ReleaseIssuanceState    string   `json:"release_issuance_state"`
	ClaimLifecycleStatus    string   `json:"claim_lifecycle_status"`
	PublicationDecision     string   `json:"publication_decision"`
	CorrectionWorkflowState string   `json:"correction_workflow_state"`
	OverrideState           string   `json:"override_state"`
	ReviewRefs              []string `json:"review_refs,omitempty"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
	FailureStates           []string `json:"failure_states,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type PublicProofFailureStateReviewItem struct {
	ClaimID                     string   `json:"claim_id"`
	ArtifactID                  string   `json:"artifact_id"`
	CurrentState                string   `json:"current_state"`
	ReviewOutcome               string   `json:"review_outcome"`
	FailureVisibilityState      string   `json:"failure_visibility_state"`
	RestrictionVisibilityState  string   `json:"restriction_visibility_state"`
	WithdrawalVisibilityState   string   `json:"withdrawal_visibility_state"`
	SupersessionVisibilityState string   `json:"supersession_visibility_state"`
	ReissueFailureState         string   `json:"reissue_failure_state"`
	SupportedFailureStates      []string `json:"supported_failure_states,omitempty"`
	ReviewRefs                  []string `json:"review_refs,omitempty"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	FailureStates               []string `json:"failure_states,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
}

func EvaluateMeasuredPublicProofValEReplayCorrectnessReviewState(items []PublicProofReplayCorrectnessReviewItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValEReplayCorrectnessReviewStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReviewOutcome) == "" {
			return MeasuredPublicProofValEReplayCorrectnessReviewStateIncomplete
		}
		if strings.TrimSpace(item.ReplayState) == "" || strings.TrimSpace(item.ComparisonMode) == "" || strings.TrimSpace(item.MethodologyRef) == "" || strings.TrimSpace(item.EvaluationRef) == "" || strings.TrimSpace(item.ToleranceDecision) == "" {
			return MeasuredPublicProofValEReplayCorrectnessReviewStatePartial
		}
		if len(item.SupportedEnvironmentClasses) == 0 || len(item.ToleranceBands) == 0 || len(item.UnsupportedReplayCases) == 0 || len(item.ReviewRefs) == 0 || len(item.EvidenceRefs) == 0 {
			return MeasuredPublicProofValEReplayCorrectnessReviewStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "replay_review_ready" || strings.TrimSpace(item.ReviewOutcome) != "approved" || (strings.TrimSpace(item.ReplayState) != "comparison_verified" && strings.TrimSpace(item.ReplayState) != "replay_verified") || strings.TrimSpace(item.ToleranceDecision) != "within_declared_bands" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValEReplayCorrectnessReviewStatePartial
	}
	return MeasuredPublicProofValEReplayCorrectnessReviewStateActive
}

func EvaluateMeasuredPublicProofValESigningTrustReviewState(items []PublicProofSigningTrustReviewItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValESigningTrustReviewStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReviewOutcome) == "" {
			return MeasuredPublicProofValESigningTrustReviewStateIncomplete
		}
		if strings.TrimSpace(item.VerificationState) == "" || strings.TrimSpace(item.TrustRootState) == "" || strings.TrimSpace(item.SigningPurposeState) == "" || strings.TrimSpace(item.HistoricalVerificationState) == "" || strings.TrimSpace(item.KeyRotationState) == "" || strings.TrimSpace(item.RevocationState) == "" || strings.TrimSpace(item.TimestampState) == "" || strings.TrimSpace(item.SignerMode) == "" || strings.TrimSpace(item.TrustRootID) == "" || strings.TrimSpace(item.KeyVersion) == "" {
			return MeasuredPublicProofValESigningTrustReviewStatePartial
		}
		if len(item.ReviewRefs) == 0 || len(item.EvidenceRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValESigningTrustReviewStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "signing_trust_review_ready" || strings.TrimSpace(item.ReviewOutcome) != "approved" || strings.TrimSpace(item.VerificationState) != "verified" || strings.TrimSpace(item.TrustRootState) != "trusted" || strings.TrimSpace(item.SigningPurposeState) != "purpose_enabled" || strings.TrimSpace(item.HistoricalVerificationState) != "historical_verification_ready" || strings.TrimSpace(item.KeyRotationState) != "rotation_ready" || strings.TrimSpace(item.RevocationState) != "revocation_ready" || strings.TrimSpace(item.TimestampState) != "timestamp_bound" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValESigningTrustReviewStatePartial
	}
	return MeasuredPublicProofValESigningTrustReviewStateActive
}

func EvaluateMeasuredPublicProofValETransparencyReviewState(items []PublicProofTransparencyReviewItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValETransparencyReviewStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReviewOutcome) == "" {
			return MeasuredPublicProofValETransparencyReviewStateIncomplete
		}
		if strings.TrimSpace(item.TransparencyState) == "" || strings.TrimSpace(item.EntryHashState) == "" || strings.TrimSpace(item.AnchorState) == "" || strings.TrimSpace(item.SupersessionVisibility) == "" || strings.TrimSpace(item.AnchorRef) == "" || strings.TrimSpace(item.EntryID) == "" || strings.TrimSpace(item.EntryHash) == "" || strings.TrimSpace(item.LineageRef) == "" {
			return MeasuredPublicProofValETransparencyReviewStatePartial
		}
		if len(item.ReviewRefs) == 0 || len(item.EvidenceRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValETransparencyReviewStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "transparency_review_ready" || strings.TrimSpace(item.ReviewOutcome) != "approved" || strings.TrimSpace(item.TransparencyState) != "anchored_projection_ready" || strings.TrimSpace(item.EntryHashState) != "digest_bound" || strings.TrimSpace(item.AnchorState) != "anchor_active" || strings.TrimSpace(item.SupersessionVisibility) != "visible_not_superseded" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValETransparencyReviewStatePartial
	}
	return MeasuredPublicProofValETransparencyReviewStateActive
}

func EvaluateMeasuredPublicProofValERedactionReviewState(items []PublicProofRedactionReviewItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValERedactionReviewStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReviewOutcome) == "" {
			return MeasuredPublicProofValERedactionReviewStateIncomplete
		}
		if strings.TrimSpace(item.RedactionTier) == "" || strings.TrimSpace(item.PublicationScope) == "" || strings.TrimSpace(item.PortalProjectionState) == "" || strings.TrimSpace(item.RedactionDecision) == "" || strings.TrimSpace(item.ProjectionDiscipline) == "" {
			return MeasuredPublicProofValERedactionReviewStatePartial
		}
		if len(item.RemovedFields) == 0 || len(item.NeverPublishedFields) == 0 || len(item.ReviewRefs) == 0 || len(item.EvidenceRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValERedactionReviewStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "redaction_review_ready" || strings.TrimSpace(item.ReviewOutcome) != "approved" || strings.TrimSpace(item.PortalProjectionState) != "portal_projection_ready" || strings.TrimSpace(item.RedactionDecision) != "tier_scope_match" || strings.TrimSpace(item.ProjectionDiscipline) != "projection_only_enforced" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValERedactionReviewStatePartial
	}
	return MeasuredPublicProofValERedactionReviewStateActive
}

func EvaluateMeasuredPublicProofValECompatibilityReviewState(items []PublicProofCompatibilityReviewItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValECompatibilityReviewStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReviewOutcome) == "" {
			return MeasuredPublicProofValECompatibilityReviewStateIncomplete
		}
		if strings.TrimSpace(item.SchemaCompatibility) == "" || strings.TrimSpace(item.VerifierCompatibility) == "" || strings.TrimSpace(item.DeprecationState) == "" || strings.TrimSpace(item.ReplayCompatibility) == "" || strings.TrimSpace(item.MethodologyCompatibility) == "" {
			return MeasuredPublicProofValECompatibilityReviewStatePartial
		}
		if len(item.SupportedSchemaLines) == 0 || len(item.UnsupportedCases) == 0 || len(item.ReviewRefs) == 0 || len(item.EvidenceRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValECompatibilityReviewStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "compatibility_review_ready" || strings.TrimSpace(item.ReviewOutcome) != "approved" || strings.TrimSpace(item.SchemaCompatibility) != "supported" || strings.TrimSpace(item.VerifierCompatibility) != "supported" || strings.TrimSpace(item.DeprecationState) != "not_deprecated" || strings.TrimSpace(item.ReplayCompatibility) != "supported" || strings.TrimSpace(item.MethodologyCompatibility) != "bounded_supported" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValECompatibilityReviewStatePartial
	}
	return MeasuredPublicProofValECompatibilityReviewStateActive
}

func EvaluateMeasuredPublicProofValEIssuanceReviewState(items []PublicProofIssuanceReviewItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValEIssuanceReviewStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReviewOutcome) == "" {
			return MeasuredPublicProofValEIssuanceReviewStateIncomplete
		}
		if strings.TrimSpace(item.ReleaseIssuanceState) == "" || strings.TrimSpace(item.ClaimLifecycleStatus) == "" || strings.TrimSpace(item.PublicationDecision) == "" || strings.TrimSpace(item.CorrectionWorkflowState) == "" || strings.TrimSpace(item.OverrideState) == "" {
			return MeasuredPublicProofValEIssuanceReviewStatePartial
		}
		if len(item.ReviewRefs) == 0 || len(item.EvidenceRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValEIssuanceReviewStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "issuance_review_ready" || strings.TrimSpace(item.ReviewOutcome) != "approved" || strings.TrimSpace(item.ReleaseIssuanceState) != "issuance_gate_ready" || strings.TrimSpace(item.CorrectionWorkflowState) != "correction_workflow_ready" || strings.TrimSpace(item.OverrideState) != "override_not_permitted" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValEIssuanceReviewStatePartial
	}
	return MeasuredPublicProofValEIssuanceReviewStateActive
}

func EvaluateMeasuredPublicProofValEFailureStateReviewState(items []PublicProofFailureStateReviewItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValEFailureStateReviewStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ReviewOutcome) == "" {
			return MeasuredPublicProofValEFailureStateReviewStateIncomplete
		}
		if strings.TrimSpace(item.FailureVisibilityState) == "" || strings.TrimSpace(item.RestrictionVisibilityState) == "" || strings.TrimSpace(item.WithdrawalVisibilityState) == "" || strings.TrimSpace(item.SupersessionVisibilityState) == "" || strings.TrimSpace(item.ReissueFailureState) == "" {
			return MeasuredPublicProofValEFailureStateReviewStatePartial
		}
		if len(item.SupportedFailureStates) == 0 || len(item.ReviewRefs) == 0 || len(item.EvidenceRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValEFailureStateReviewStatePartial
		}
		if !containsTrimmedString(item.SupportedFailureStates, PublicProofStatusProofFailed) ||
			!containsTrimmedString(item.SupportedFailureStates, PublicProofStatusClaimNotReissued) ||
			!containsTrimmedString(item.SupportedFailureStates, PublicProofStatusRestricted) ||
			!containsTrimmedString(item.SupportedFailureStates, PublicProofStatusSuperseded) ||
			!containsTrimmedString(item.SupportedFailureStates, PublicProofStatusWithdrawn) ||
			!containsTrimmedString(item.SupportedFailureStates, PublicProofStatusStale) {
			return MeasuredPublicProofValEFailureStateReviewStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "failure_state_review_ready" || strings.TrimSpace(item.ReviewOutcome) != "approved" || strings.TrimSpace(item.FailureVisibilityState) != "verifier_visible_failures_modeled" || strings.TrimSpace(item.RestrictionVisibilityState) != "restriction_visible" || strings.TrimSpace(item.WithdrawalVisibilityState) != "withdrawal_visible_not_triggered" || strings.TrimSpace(item.SupersessionVisibilityState) != "supersession_visible_not_triggered" || strings.TrimSpace(item.ReissueFailureState) != "claim_not_reissued_modeled" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValEFailureStateReviewStatePartial
	}
	return MeasuredPublicProofValEFailureStateReviewStateActive
}

func EvaluateMeasuredPublicProofValEState(valDState, replayCorrectnessReviewState, signingTrustReviewState, transparencyReviewState, redactionReviewState, compatibilityReviewState, issuanceReviewState, failureStateReviewState string) string {
	if strings.TrimSpace(valDState) != MeasuredPublicProofValDStateActive {
		return MeasuredPublicProofValEStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(replayCorrectnessReviewState),
		strings.TrimSpace(signingTrustReviewState),
		strings.TrimSpace(transparencyReviewState),
		strings.TrimSpace(redactionReviewState),
		strings.TrimSpace(compatibilityReviewState),
		strings.TrimSpace(issuanceReviewState),
		strings.TrimSpace(failureStateReviewState),
	} {
		switch state {
		case MeasuredPublicProofValEReplayCorrectnessReviewStateActive,
			MeasuredPublicProofValESigningTrustReviewStateActive,
			MeasuredPublicProofValETransparencyReviewStateActive,
			MeasuredPublicProofValERedactionReviewStateActive,
			MeasuredPublicProofValECompatibilityReviewStateActive,
			MeasuredPublicProofValEIssuanceReviewStateActive,
			MeasuredPublicProofValEFailureStateReviewStateActive:
		case MeasuredPublicProofValEReplayCorrectnessReviewStatePartial,
			MeasuredPublicProofValESigningTrustReviewStatePartial,
			MeasuredPublicProofValETransparencyReviewStatePartial,
			MeasuredPublicProofValERedactionReviewStatePartial,
			MeasuredPublicProofValECompatibilityReviewStatePartial,
			MeasuredPublicProofValEIssuanceReviewStatePartial,
			MeasuredPublicProofValEFailureStateReviewStatePartial:
			hasPartial = true
		default:
			return MeasuredPublicProofValEStateIncomplete
		}
	}
	if hasPartial {
		return MeasuredPublicProofValEStateSubstantial
	}
	return MeasuredPublicProofValEStateActive
}
