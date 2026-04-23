package claims

import "strings"

const (
	MeasuredPublicProofValBTransparencyChainStateActive     = "measured_public_proof_valb_transparency_chain_active"
	MeasuredPublicProofValBTransparencyChainStatePartial    = "measured_public_proof_valb_transparency_chain_partial"
	MeasuredPublicProofValBTransparencyChainStateIncomplete = "measured_public_proof_valb_transparency_chain_incomplete"

	MeasuredPublicProofValBVerifierCapabilityStateActive     = "measured_public_proof_valb_verifier_capability_active"
	MeasuredPublicProofValBVerifierCapabilityStatePartial    = "measured_public_proof_valb_verifier_capability_partial"
	MeasuredPublicProofValBVerifierCapabilityStateIncomplete = "measured_public_proof_valb_verifier_capability_incomplete"

	MeasuredPublicProofValBSignatureVerificationStateActive     = "measured_public_proof_valb_signature_verification_active"
	MeasuredPublicProofValBSignatureVerificationStatePartial    = "measured_public_proof_valb_signature_verification_partial"
	MeasuredPublicProofValBSignatureVerificationStateIncomplete = "measured_public_proof_valb_signature_verification_incomplete"

	MeasuredPublicProofValBReplayVerificationStateActive     = "measured_public_proof_valb_replay_verification_active"
	MeasuredPublicProofValBReplayVerificationStatePartial    = "measured_public_proof_valb_replay_verification_partial"
	MeasuredPublicProofValBReplayVerificationStateIncomplete = "measured_public_proof_valb_replay_verification_incomplete"

	MeasuredPublicProofValBStateIncomplete  = "measured_public_proof_valb_incomplete"
	MeasuredPublicProofValBStateSubstantial = "measured_public_proof_valb_substantially_ready"
	MeasuredPublicProofValBStateActive      = "measured_public_proof_valb_active"
)

type PublicProofTransparencyEntry struct {
	ArtifactID       string   `json:"artifact_id"`
	CurrentState     string   `json:"current_state"`
	ParentAnchorRef  string   `json:"parent_anchor_ref"`
	AnchorID         string   `json:"anchor_id"`
	EntryID          string   `json:"entry_id"`
	EntryHash        string   `json:"entry_hash"`
	AnchoredAt       string   `json:"anchored_at"`
	TransparencyRefs []string `json:"transparency_refs,omitempty"`
	SupersessionRefs []string `json:"supersession_refs,omitempty"`
	Limitations      []string `json:"limitations,omitempty"`
}

type PublicProofTransparencyChain struct {
	CurrentState     string                         `json:"current_state"`
	ChainID          string                         `json:"chain_id"`
	ParentAnchorRef  string                         `json:"parent_anchor_ref"`
	Entries          []PublicProofTransparencyEntry `json:"entries,omitempty"`
	IntegrityRules   []string                       `json:"integrity_rules,omitempty"`
	PublicationRules []string                       `json:"publication_rules,omitempty"`
	Limitations      []string                       `json:"limitations,omitempty"`
}

type PublicProofVerifierCapability struct {
	CurrentState         string   `json:"current_state"`
	SDKRef               string   `json:"sdk_ref"`
	ReferencePackRef     string   `json:"reference_pack_ref"`
	OfflineGuideRef      string   `json:"offline_guide_ref"`
	SupportedSchemaLines []string `json:"supported_schema_lines,omitempty"`
	ResultStates         []string `json:"result_states,omitempty"`
	TrustVerification    []string `json:"trust_verification_rules,omitempty"`
	ReplayVerification   []string `json:"replay_verification_rules,omitempty"`
	CommandHints         []string `json:"command_hints,omitempty"`
	UnsupportedCases     []string `json:"unsupported_cases,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type PublicProofSignatureVerificationItem struct {
	ArtifactID          string   `json:"artifact_id"`
	CurrentState        string   `json:"current_state"`
	VerificationState   string   `json:"verification_state"`
	TrustRootState      string   `json:"trust_root_state"`
	SchemaCompatibility string   `json:"schema_compatibility"`
	VerifierRef         string   `json:"verifier_ref"`
	TrustRootID         string   `json:"trust_root_id"`
	KeyVersion          string   `json:"key_version"`
	SignatureProvider   string   `json:"signature_provider"`
	PayloadDigest       string   `json:"payload_digest"`
	EvidenceRefs        []string `json:"evidence_refs,omitempty"`
	FailureStates       []string `json:"failure_states,omitempty"`
	Limitations         []string `json:"limitations,omitempty"`
}

type PublicProofReplayVerificationItem struct {
	ArtifactID                  string   `json:"artifact_id"`
	CurrentState                string   `json:"current_state"`
	ComparisonMode              string   `json:"comparison_mode"`
	MethodologyRef              string   `json:"methodology_ref"`
	EvaluationState             string   `json:"evaluation_state"`
	EvaluationRef               string   `json:"evaluation_ref"`
	SupportedEnvironmentClasses []string `json:"supported_environment_classes,omitempty"`
	ToleranceBands              []string `json:"tolerance_bands,omitempty"`
	ReplayCommandHints          []string `json:"replay_command_hints,omitempty"`
	UnsupportedReplayCases      []string `json:"unsupported_replay_cases,omitempty"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
}

func EvaluateMeasuredPublicProofValBTransparencyChainState(model PublicProofTransparencyChain) string {
	if strings.TrimSpace(model.ChainID) == "" || strings.TrimSpace(model.ParentAnchorRef) == "" || len(model.Entries) == 0 {
		return MeasuredPublicProofValBTransparencyChainStateIncomplete
	}
	if len(model.IntegrityRules) == 0 || len(model.PublicationRules) == 0 {
		return MeasuredPublicProofValBTransparencyChainStatePartial
	}
	for _, entry := range model.Entries {
		if strings.TrimSpace(entry.ArtifactID) == "" || strings.TrimSpace(entry.CurrentState) == "" || strings.TrimSpace(entry.AnchorID) == "" || strings.TrimSpace(entry.EntryID) == "" || strings.TrimSpace(entry.EntryHash) == "" || strings.TrimSpace(entry.AnchoredAt) == "" {
			return MeasuredPublicProofValBTransparencyChainStatePartial
		}
		if len(entry.TransparencyRefs) == 0 {
			return MeasuredPublicProofValBTransparencyChainStatePartial
		}
	}
	return MeasuredPublicProofValBTransparencyChainStateActive
}

func EvaluateMeasuredPublicProofValBVerifierCapabilityState(model PublicProofVerifierCapability) string {
	if strings.TrimSpace(model.SDKRef) == "" || strings.TrimSpace(model.ReferencePackRef) == "" || strings.TrimSpace(model.OfflineGuideRef) == "" {
		return MeasuredPublicProofValBVerifierCapabilityStateIncomplete
	}
	if len(model.SupportedSchemaLines) == 0 || len(model.ResultStates) == 0 || len(model.TrustVerification) == 0 || len(model.ReplayVerification) == 0 || len(model.CommandHints) == 0 {
		return MeasuredPublicProofValBVerifierCapabilityStatePartial
	}
	return MeasuredPublicProofValBVerifierCapabilityStateActive
}

func EvaluateMeasuredPublicProofValBSignatureVerificationState(items []PublicProofSignatureVerificationItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValBSignatureVerificationStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.VerificationState) == "" {
			return MeasuredPublicProofValBSignatureVerificationStateIncomplete
		}
		if strings.TrimSpace(item.TrustRootState) == "" || strings.TrimSpace(item.SchemaCompatibility) == "" || strings.TrimSpace(item.VerifierRef) == "" || strings.TrimSpace(item.TrustRootID) == "" || strings.TrimSpace(item.KeyVersion) == "" || strings.TrimSpace(item.SignatureProvider) == "" || strings.TrimSpace(item.PayloadDigest) == "" || len(item.EvidenceRefs) == 0 || len(item.FailureStates) == 0 {
			return MeasuredPublicProofValBSignatureVerificationStatePartial
		}
		if item.VerificationState != "verified" || item.TrustRootState != "trusted" || item.SchemaCompatibility != "supported" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValBSignatureVerificationStatePartial
	}
	return MeasuredPublicProofValBSignatureVerificationStateActive
}

func EvaluateMeasuredPublicProofValBReplayVerificationState(items []PublicProofReplayVerificationItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValBReplayVerificationStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ComparisonMode) == "" {
			return MeasuredPublicProofValBReplayVerificationStateIncomplete
		}
		if strings.TrimSpace(item.MethodologyRef) == "" || strings.TrimSpace(item.EvaluationState) == "" || strings.TrimSpace(item.EvaluationRef) == "" || len(item.SupportedEnvironmentClasses) == 0 || len(item.ToleranceBands) == 0 || len(item.ReplayCommandHints) == 0 || len(item.UnsupportedReplayCases) == 0 || len(item.EvidenceRefs) == 0 {
			return MeasuredPublicProofValBReplayVerificationStatePartial
		}
		if item.EvaluationState != "replay_verified" && item.EvaluationState != "comparison_verified" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValBReplayVerificationStatePartial
	}
	return MeasuredPublicProofValBReplayVerificationStateActive
}

func EvaluateMeasuredPublicProofValBState(valAState, transparencyChainState, verifierCapabilityState, signatureVerificationState, replayVerificationState string) string {
	if strings.TrimSpace(valAState) != MeasuredPublicProofValAStateActive {
		return MeasuredPublicProofValBStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(transparencyChainState),
		strings.TrimSpace(verifierCapabilityState),
		strings.TrimSpace(signatureVerificationState),
		strings.TrimSpace(replayVerificationState),
	} {
		switch state {
		case MeasuredPublicProofValBTransparencyChainStateActive,
			MeasuredPublicProofValBVerifierCapabilityStateActive,
			MeasuredPublicProofValBSignatureVerificationStateActive,
			MeasuredPublicProofValBReplayVerificationStateActive:
		case MeasuredPublicProofValBTransparencyChainStatePartial,
			MeasuredPublicProofValBVerifierCapabilityStatePartial,
			MeasuredPublicProofValBSignatureVerificationStatePartial,
			MeasuredPublicProofValBReplayVerificationStatePartial:
			hasPartial = true
		default:
			return MeasuredPublicProofValBStateIncomplete
		}
	}
	if hasPartial {
		return MeasuredPublicProofValBStateSubstantial
	}
	return MeasuredPublicProofValBStateActive
}
