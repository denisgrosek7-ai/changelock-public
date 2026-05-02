package formal

import (
	"encoding/hex"
	"strings"
	"time"
)

const (
	Point11ValAStateActive         = "point11_vala_policy_authority_core_active"
	Point11ValAStateBlocked        = "point11_vala_policy_authority_core_blocked"
	Point11ValAStateReviewRequired = "point11_vala_policy_authority_core_review_required"

	Point11ValADependencyStateActive         = "point11_vala_dependency_active"
	Point11ValADependencyStateBlocked        = "point11_vala_dependency_blocked"
	Point11ValADependencyStateReviewRequired = "point11_vala_dependency_review_required"

	Point11ValARegistryStateActive  = "point11_vala_registry_active"
	Point11ValARegistryStateBlocked = "point11_vala_registry_blocked"

	Point11ValASignatureStateActive  = "point11_vala_signature_active"
	Point11ValASignatureStateBlocked = "point11_vala_signature_blocked"

	Point11ValAAnchorStateActive  = "point11_vala_anchor_active"
	Point11ValAAnchorStateBlocked = "point11_vala_anchor_blocked"

	Point11ValALifecycleTransitionStateActive  = "point11_vala_lifecycle_transition_active"
	Point11ValALifecycleTransitionStateBlocked = "point11_vala_lifecycle_transition_blocked"

	Point11ValALifecycleStateActive  = Point11ValALifecycleTransitionStateActive
	Point11ValALifecycleStateBlocked = Point11ValALifecycleTransitionStateBlocked

	Point11ValAPolicyUseStateActive         = "point11_vala_policy_use_active"
	Point11ValAPolicyUseStateBlocked        = "point11_vala_policy_use_blocked"
	Point11ValAPolicyUseStateHistoricalOnly = "point11_vala_policy_use_historical_only"
	Point11ValAPolicyUseStateNotYetActive   = "point11_vala_policy_use_not_yet_active"

	Point11ValAGraphStateActive  = "point11_vala_graph_active"
	Point11ValAGraphStateBlocked = "point11_vala_graph_blocked"
)

const (
	point11ValAProjectionDisclaimerBaseline = "projection_only not_canonical_truth point11_vala_policy_authority_core"

	point11ValAPolicyLifecycleDraft          = "draft"
	point11ValAPolicyLifecycleReviewRequired = "review_required"
	point11ValAPolicyLifecycleApproved       = "approved"
	point11ValAPolicyLifecycleActive         = "active"
	point11ValAPolicyLifecycleDeprecated     = "deprecated"
	point11ValAPolicyLifecycleSuperseded     = "superseded"
	point11ValAPolicyLifecycleRevoked        = "revoked"
	point11ValAPolicyLifecycleExpired        = "expired"
	point11ValAPolicyLifecycleBlocked        = "blocked"

	point11ValASignatureEnvelopeStateVerified = "signature_state_verified"
	point11ValASignatureVerificationActive    = "signature_verification_active"

	point11ValAAnchorEnvelopeStateVerified = "anchor_state_verified"
	point11ValAAnchorVerificationActive    = "anchor_verification_active"
)

type Point11ValAVal0ReviewContext struct {
	LocalReviewAllowsDependencyReviewRequired bool     `json:"local_review_allows_dependency_review_required"`
	Val0Point11PassEmitted                    bool     `json:"val0_point11_pass_emitted"`
	Val0CreatesAuthorityClaims                bool     `json:"val0_creates_authority_claims"`
	Val0CreatesPublicationSideEffects         bool     `json:"val0_creates_publication_side_effects"`
	OpenCLB0Findings                          int      `json:"open_clb0_findings"`
	OpenCLB1Findings                          int      `json:"open_clb1_findings"`
	OpenCLB2Findings                          int      `json:"open_clb2_findings"`
	ReviewPrerequisites                       []string `json:"review_prerequisites,omitempty"`
}

type Point11ValADependencySnapshot struct {
	Val0CurrentState                  string   `json:"val0_current_state"`
	Val0DependencyState               string   `json:"val0_dependency_state"`
	Val0PolicyContractState           string   `json:"val0_policy_contract_state"`
	Val0ClaimGovernanceState          string   `json:"val0_claim_governance_state"`
	Val0AuthorityMatrixState          string   `json:"val0_authority_matrix_state"`
	Val0ExceptionGovernanceState      string   `json:"val0_exception_governance_state"`
	Val0ABACState                     string   `json:"val0_abac_state"`
	Val0DecisionBindingState          string   `json:"val0_decision_binding_state"`
	Val0NoOverclaimState              string   `json:"val0_no_overclaim_state"`
	Val0CrossDomainCompatibilityState string   `json:"val0_cross_domain_compatibility_state"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
	Val0Point11PassEmitted            bool     `json:"val0_point11_pass_emitted"`
	Val0CreatesAuthorityClaims        bool     `json:"val0_creates_authority_claims"`
	Val0CreatesPublicationSideEffects bool     `json:"val0_creates_publication_side_effects"`
	OpenCLB0Findings                  int      `json:"open_clb0_findings"`
	OpenCLB1Findings                  int      `json:"open_clb1_findings"`
	OpenCLB2Findings                  int      `json:"open_clb2_findings"`
	LocalReviewAllowsReviewRequired   bool     `json:"local_review_allows_review_required"`
	ReviewPrerequisites               []string `json:"review_prerequisites,omitempty"`
}

type Point11ValASignedPolicyRegistry struct {
	CurrentState         string   `json:"current_state"`
	RegistryID           string   `json:"registry_id"`
	RegistryVersion      string   `json:"registry_version"`
	PolicyPackID         string   `json:"policy_pack_id"`
	PolicyID             string   `json:"policy_id"`
	PolicyVersion        string   `json:"policy_version"`
	PolicyScope          string   `json:"policy_scope"`
	PolicyOwner          string   `json:"policy_owner"`
	PolicyIssuer         string   `json:"policy_issuer"`
	ApprovalChainRefs    []string `json:"approval_chain_refs,omitempty"`
	SignatureRef         string   `json:"signature_ref"`
	SigningKeyRef        string   `json:"signing_key_ref"`
	SigningAlgorithm     string   `json:"signing_algorithm"`
	SignatureState       string   `json:"signature_state"`
	AnchorRef            string   `json:"anchor_ref"`
	AnchorState          string   `json:"anchor_state"`
	SchemaVersion        string   `json:"schema_version"`
	CompatibilityVersion string   `json:"compatibility_version"`
	EffectiveFrom        string   `json:"effective_from"`
	EffectiveUntil       string   `json:"effective_until"`
	LifecycleState       string   `json:"lifecycle_state"`
	SupersededBy         string   `json:"superseded_by"`
	RevokedBy            string   `json:"revoked_by"`
	GovernanceEventRef   string   `json:"governance_event_ref"`
	ApprovalEvidenceRefs []string `json:"approval_evidence_refs,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type Point11ValAPolicySignatureEnvelope struct {
	CurrentState             string   `json:"current_state"`
	SignatureRef             string   `json:"signature_ref"`
	SigningKeyRef            string   `json:"signing_key_ref"`
	SigningAlgorithm         string   `json:"signing_algorithm"`
	SignedSubjectRef         string   `json:"signed_subject_ref"`
	SignedSubjectHash        string   `json:"signed_subject_hash"`
	SignerIdentity           string   `json:"signer_identity"`
	IssuedAt                 string   `json:"issued_at"`
	ExpiresAt                string   `json:"expires_at"`
	RevocationRef            string   `json:"revocation_ref"`
	SignatureState           string   `json:"signature_state"`
	VerificationResult       string   `json:"verification_result"`
	VerificationEvidenceRefs []string `json:"verification_evidence_refs,omitempty"`
}

type Point11ValAPolicyAnchorEnvelope struct {
	CurrentState             string   `json:"current_state"`
	AnchorRef                string   `json:"anchor_ref"`
	AnchorType               string   `json:"anchor_type"`
	AnchoredSubjectRef       string   `json:"anchored_subject_ref"`
	AnchoredSubjectHash      string   `json:"anchored_subject_hash"`
	AnchorTimestamp          string   `json:"anchor_timestamp"`
	AnchorState              string   `json:"anchor_state"`
	AnchorEvidenceRefs       []string `json:"anchor_evidence_refs,omitempty"`
	AnchorVerificationResult string   `json:"anchor_verification_result"`
}

type Point11ValAPolicyLifecycleTransition struct {
	CurrentState                string   `json:"current_state"`
	TransitionID                string   `json:"transition_id"`
	FromState                   string   `json:"from_state"`
	ToState                     string   `json:"to_state"`
	PolicyRef                   string   `json:"policy_ref"`
	ActorRef                    string   `json:"actor_ref"`
	ApproverRef                 string   `json:"approver_ref"`
	GovernanceEventRef          string   `json:"governance_event_ref"`
	Reason                      string   `json:"reason"`
	TransitionTimestamp         string   `json:"transition_timestamp"`
	ApprovalEvidenceRefs        []string `json:"approval_evidence_refs,omitempty"`
	CompatibilityReviewRef      string   `json:"compatibility_review_ref"`
	RollbackRef                 string   `json:"rollback_ref"`
	AuditID                     string   `json:"audit_id"`
	SuccessorPolicyRef          string   `json:"successor_policy_ref"`
	SignatureEnvelopeState      string   `json:"signature_envelope_state"`
	AnchorEnvelopeState         string   `json:"anchor_envelope_state"`
	PolicyRelaxationRequested   bool     `json:"policy_relaxation_requested"`
	AuthorityExpansionRequested bool     `json:"authority_expansion_requested"`
}

type Point11ValAPolicySupersessionRevocationGraph struct {
	CurrentState            string   `json:"current_state"`
	GraphID                 string   `json:"graph_id"`
	SourcePolicyRef         string   `json:"source_policy_ref"`
	SuccessorPolicyRef      string   `json:"successor_policy_ref"`
	RevokedByRef            string   `json:"revoked_by_ref"`
	SupersessionReason      string   `json:"supersession_reason"`
	RevocationReason        string   `json:"revocation_reason"`
	CompatibilityVersion    string   `json:"compatibility_version"`
	CompatibilityReviewRef  string   `json:"compatibility_review_ref"`
	EffectiveFrom           string   `json:"effective_from"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
	AuditID                 string   `json:"audit_id"`
	SuccessorLifecycleState string   `json:"successor_lifecycle_state"`
	LineagePath             []string `json:"lineage_path,omitempty"`
}

type Point11ValALifecycleEvaluation struct {
	TransitionState string   `json:"transition_state"`
	PolicyUseState  string   `json:"policy_use_state"`
	Reason          string   `json:"reason"`
	Diagnostics     []string `json:"diagnostics,omitempty"`
}

type Point11ValADiagnostics struct {
	CurrentState         string   `json:"current_state"`
	BlockingReasons      []string `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites  []string `json:"review_prerequisites,omitempty"`
	ComponentStates      []string `json:"component_states,omitempty"`
	LifecycleReason      string   `json:"lifecycle_reason"`
	LifecycleDiagnostics []string `json:"lifecycle_diagnostics,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type Point11ValAFoundation struct {
	CurrentState                             string                                       `json:"current_state"`
	BlockingReasons                          []string                                     `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                      []string                                     `json:"review_prerequisites,omitempty"`
	DependencyState                          string                                       `json:"dependency_state"`
	RegistryState                            string                                       `json:"registry_state"`
	SignatureState                           string                                       `json:"signature_state"`
	AnchorState                              string                                       `json:"anchor_state"`
	LifecycleState                           string                                       `json:"lifecycle_state"`
	LifecycleTransitionState                 string                                       `json:"lifecycle_transition_state"`
	PolicyUseState                           string                                       `json:"policy_use_state"`
	GraphState                               string                                       `json:"graph_state"`
	LifecycleEvaluation                      Point11ValALifecycleEvaluation               `json:"lifecycle_evaluation"`
	Diagnostics                              Point11ValADiagnostics                       `json:"diagnostics"`
	ProjectionDisclaimer                     string                                       `json:"projection_disclaimer"`
	CreatesLegalRegulatoryCertificationClaim bool                                         `json:"creates_legal_regulatory_certification_claim"`
	CreatesPublicationSideEffects            bool                                         `json:"creates_publication_side_effects"`
	Dependency                               Point11ValADependencySnapshot                `json:"dependency"`
	Registry                                 Point11ValASignedPolicyRegistry              `json:"registry"`
	Signature                                Point11ValAPolicySignatureEnvelope           `json:"signature"`
	Anchor                                   Point11ValAPolicyAnchorEnvelope              `json:"anchor"`
	Lifecycle                                Point11ValAPolicyLifecycleTransition         `json:"lifecycle"`
	Graph                                    Point11ValAPolicySupersessionRevocationGraph `json:"graph"`
}

func point11ValARegistryRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"registry_",
		"point11_registry_",
	})
}

func point11ValAGraphRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"graph_",
		"point11_graph_",
	})
}

func point11ValAPolicyPackRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"policy_pack_",
		"point11_policy_pack_",
	})
}

func point11ValAPolicyRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"policy_",
		"pol_",
		"point11_policy_",
	})
}

func point11ValASignatureRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"signature_",
		"point11_signature_",
	})
}

func point11ValASigningKeyRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"signing_key_",
		"point11_signing_key_",
	})
}

func point11ValAAnchorRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"anchor_",
		"point11_anchor_",
	})
}

func point11ValAGovernanceEventRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"governance_event_",
		"point11_governance_event_",
	})
}

func point11ValATransitionRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"transition_",
		"point11_transition_",
	})
}

func point11ValAAuditRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"audit_",
		"point11_audit_",
	})
}

func point11ValACompatibilityReviewRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"compatibility_review_",
		"point11_compatibility_review_",
	})
}

func point11ValARevokerRefValid(value string) bool {
	return point11ValAPolicyRefValid(value) || point11ValAGovernanceEventRefValid(value)
}

func point11ValAHashValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return false
	}
	if strings.HasPrefix(trimmed, "sha256:") {
		trimmed = strings.TrimPrefix(trimmed, "sha256:")
	}
	if len(trimmed) != 64 {
		return false
	}
	_, err := hex.DecodeString(trimmed)
	return err == nil
}

func point11ValASigningAlgorithmAllowed(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		"ed25519",
		"ecdsa_p256",
	}, value)
}

func point11ValAAnchorTypeAllowed(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		"trusted_timestamp",
		"transparency_log",
		"internal_anchor_record",
	}, value)
}

func point11ValAPolicyLifecycleAllowed(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValAPolicyLifecycleDraft,
		point11ValAPolicyLifecycleReviewRequired,
		point11ValAPolicyLifecycleApproved,
		point11ValAPolicyLifecycleActive,
		point11ValAPolicyLifecycleDeprecated,
		point11ValAPolicyLifecycleSuperseded,
		point11ValAPolicyLifecycleRevoked,
		point11ValAPolicyLifecycleExpired,
		point11ValAPolicyLifecycleBlocked,
	}, value)
}

func point11ValAPolicyLifecycleInvalidated(value string) bool {
	switch strings.TrimSpace(value) {
	case point11ValAPolicyLifecycleRevoked, point11ValAPolicyLifecycleExpired, point11ValAPolicyLifecycleSuperseded, point11ValAPolicyLifecycleBlocked:
		return true
	default:
		return false
	}
}

func point11ValASignatureEnvelopeActive(value string) bool {
	return strings.TrimSpace(value) == Point11ValASignatureStateActive
}

func point11ValAAnchorEnvelopeActive(value string) bool {
	return strings.TrimSpace(value) == Point11ValAAnchorStateActive
}

func point11ValAGraphHasCycle(path []string) bool {
	if len(path) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, ref := range path {
		trimmed := strings.TrimSpace(ref)
		if !point11ValAPolicyRefValid(trimmed) {
			return true
		}
		if _, exists := seen[trimmed]; exists {
			return true
		}
		seen[trimmed] = struct{}{}
	}
	return false
}

func point11ValAComponentStates(model Point11ValAFoundation) []string {
	return []string{
		"dependency:" + model.DependencyState,
		"registry:" + model.RegistryState,
		"signature:" + model.SignatureState,
		"anchor:" + model.AnchorState,
		"lifecycle_transition:" + model.LifecycleTransitionState,
		"policy_use:" + model.PolicyUseState,
		"graph:" + model.GraphState,
	}
}

// Dependency snapshots must copy actual computed upstream output.
// They must not repair, replace, fallback, or regenerate upstream dependency values.
// The dependency evaluator is responsible for fail-closed validation.
func SnapshotPoint11ValADependencyFromComputedVal0(val0 Point11Val0Foundation, review Point11ValAVal0ReviewContext) Point11ValADependencySnapshot {
	reviewPrerequisites := append([]string{}, val0.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point11ValADependencySnapshot{
		Val0CurrentState:                  val0.CurrentState,
		Val0DependencyState:               val0.DependencyState,
		Val0PolicyContractState:           val0.PolicyContractState,
		Val0ClaimGovernanceState:          val0.ClaimGovernanceState,
		Val0AuthorityMatrixState:          val0.AuthorityMatrixState,
		Val0ExceptionGovernanceState:      val0.ExceptionGovernanceState,
		Val0ABACState:                     val0.ABACGovernanceState,
		Val0DecisionBindingState:          val0.DecisionBindingState,
		Val0NoOverclaimState:              val0.NoOverclaimState,
		Val0CrossDomainCompatibilityState: val0.CrossDomainCompatibilityState,
		ProjectionDisclaimer:              val0.ProjectionDisclaimer,
		Val0Point11PassEmitted:            review.Val0Point11PassEmitted,
		Val0CreatesAuthorityClaims:        review.Val0CreatesAuthorityClaims,
		Val0CreatesPublicationSideEffects: review.Val0CreatesPublicationSideEffects,
		OpenCLB0Findings:                  review.OpenCLB0Findings,
		OpenCLB1Findings:                  review.OpenCLB1Findings,
		OpenCLB2Findings:                  review.OpenCLB2Findings,
		LocalReviewAllowsReviewRequired:   review.LocalReviewAllowsDependencyReviewRequired,
		ReviewPrerequisites:               reviewPrerequisites,
	}
}

func point11ValADependencyReviewContextModel() Point11ValAVal0ReviewContext {
	return Point11ValAVal0ReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
		Val0Point11PassEmitted:                    false,
		Val0CreatesAuthorityClaims:                false,
		Val0CreatesPublicationSideEffects:         false,
		OpenCLB0Findings:                          0,
		OpenCLB1Findings:                          0,
		OpenCLB2Findings:                          0,
	}
}

func point11ValADependencySnapshotModel() Point11ValADependencySnapshot {
	val0 := ComputePoint11Val0Foundation(Point11Val0FoundationModel())
	return SnapshotPoint11ValADependencyFromComputedVal0(val0, point11ValADependencyReviewContextModel())
}

func EvaluatePoint11ValADependencyState(model Point11ValADependencySnapshot) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.Val0Point11PassEmitted ||
		model.Val0CreatesAuthorityClaims ||
		model.Val0CreatesPublicationSideEffects ||
		model.OpenCLB0Findings > 0 ||
		model.OpenCLB1Findings > 0 ||
		model.OpenCLB2Findings > 0 ||
		strings.TrimSpace(model.Val0PolicyContractState) != Point11Val0PolicyContractStateActive ||
		strings.TrimSpace(model.Val0ClaimGovernanceState) != Point11Val0ClaimGovernanceStateActive ||
		strings.TrimSpace(model.Val0AuthorityMatrixState) != Point11Val0AuthorityMatrixStateActive ||
		strings.TrimSpace(model.Val0ExceptionGovernanceState) != Point11Val0ExceptionGovernanceStateActive ||
		strings.TrimSpace(model.Val0ABACState) != Point11Val0ABACStateActive ||
		strings.TrimSpace(model.Val0DecisionBindingState) != Point11Val0DecisionBindingStateActive ||
		strings.TrimSpace(model.Val0NoOverclaimState) != Point11Val0NoOverclaimStateActive ||
		strings.TrimSpace(model.Val0CrossDomainCompatibilityState) != Point11Val0CrossDomainCompatibilityStateActive {
		return Point11ValADependencyStateBlocked
	}
	if strings.TrimSpace(model.Val0CurrentState) == Point11Val0StateActive &&
		strings.TrimSpace(model.Val0DependencyState) == Point11Val0DependencyStateActive {
		return Point11ValADependencyStateActive
	}
	if model.LocalReviewAllowsReviewRequired &&
		(strings.TrimSpace(model.Val0CurrentState) == Point11Val0StateReviewRequired ||
			strings.TrimSpace(model.Val0DependencyState) == Point11Val0DependencyStateReviewRequired) {
		return Point11ValADependencyStateReviewRequired
	}
	return Point11ValADependencyStateBlocked
}

func EvaluatePoint11ValARegistryState(model Point11ValASignedPolicyRegistry) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11ValARegistryRefValid(model.RegistryID) ||
		!point11Val0IdentityValueValid(model.RegistryVersion) ||
		!point11ValAPolicyPackRefValid(model.PolicyPackID) ||
		!point11ValAPolicyRefValid(model.PolicyID) ||
		!point11Val0IdentityValueValid(model.PolicyVersion) ||
		!point11Val0ScopeValid(model.PolicyScope) ||
		!point11Val0IdentityValueValid(model.PolicyOwner) ||
		!point11Val0IdentityValueValid(model.PolicyIssuer) ||
		!point11Val0AllValuesValid(model.ApprovalChainRefs) ||
		!point11ValASignatureRefValid(model.SignatureRef) ||
		!point11ValASigningKeyRefValid(model.SigningKeyRef) ||
		!point11ValASigningAlgorithmAllowed(model.SigningAlgorithm) ||
		strings.TrimSpace(model.SignatureState) != point11ValASignatureEnvelopeStateVerified ||
		!point11ValAAnchorRefValid(model.AnchorRef) ||
		strings.TrimSpace(model.AnchorState) != point11ValAAnchorEnvelopeStateVerified ||
		!point11Val0IdentityValueValid(model.SchemaVersion) ||
		!point11Val0IdentityValueValid(model.CompatibilityVersion) ||
		!point11Val0ValidTimestamp(model.EffectiveFrom) ||
		!point11Val0ValidTimestamp(model.EffectiveUntil) ||
		!point11ValAPolicyLifecycleAllowed(model.LifecycleState) ||
		!point11ValAGovernanceEventRefValid(model.GovernanceEventRef) ||
		!point11Val0EvidenceRefsValid(model.ApprovalEvidenceRefs) {
		return Point11ValARegistryStateBlocked
	}
	expiresAt, _ := time.Parse(time.RFC3339, strings.TrimSpace(model.EffectiveUntil))
	if expiresAt.Before(time.Now().UTC()) || strings.TrimSpace(model.RevokedBy) != "" {
		return Point11ValARegistryStateBlocked
	}
	if strings.TrimSpace(model.LifecycleState) == point11ValAPolicyLifecycleSuperseded {
		if !point11ValAPolicyRefValid(model.SupersededBy) || !point11Val0IdentityValueValid(model.CompatibilityVersion) {
			return Point11ValARegistryStateBlocked
		}
		return Point11ValARegistryStateBlocked
	}
	if strings.TrimSpace(model.LifecycleState) == point11ValAPolicyLifecycleRevoked ||
		strings.TrimSpace(model.LifecycleState) == point11ValAPolicyLifecycleExpired ||
		strings.TrimSpace(model.LifecycleState) == point11ValAPolicyLifecycleBlocked {
		return Point11ValARegistryStateBlocked
	}
	if strings.TrimSpace(model.LifecycleState) != point11ValAPolicyLifecycleActive &&
		strings.TrimSpace(model.LifecycleState) != point11ValAPolicyLifecycleApproved {
		return Point11ValARegistryStateBlocked
	}
	return Point11ValARegistryStateActive
}

func EvaluatePoint11ValASignatureState(model Point11ValAPolicySignatureEnvelope) string {
	if !point11ValASignatureRefValid(model.SignatureRef) ||
		!point11ValASigningKeyRefValid(model.SigningKeyRef) ||
		!point11ValASigningAlgorithmAllowed(model.SigningAlgorithm) ||
		!point11ValAPolicyRefValid(model.SignedSubjectRef) ||
		!point11ValAHashValid(model.SignedSubjectHash) ||
		!point11Val0IdentityValueValid(model.SignerIdentity) ||
		!point11Val0ValidTimestamp(model.IssuedAt) ||
		!point11Val0ValidTimestamp(model.ExpiresAt) ||
		strings.TrimSpace(model.SignatureState) != point11ValASignatureEnvelopeStateVerified ||
		strings.TrimSpace(model.VerificationResult) != point11ValASignatureVerificationActive ||
		!point11Val0EvidenceRefsValid(model.VerificationEvidenceRefs) {
		return Point11ValASignatureStateBlocked
	}
	if strings.TrimSpace(model.RevocationRef) != "" && !point11ValAGovernanceEventRefValid(model.RevocationRef) {
		return Point11ValASignatureStateBlocked
	}
	expiresAt, _ := time.Parse(time.RFC3339, strings.TrimSpace(model.ExpiresAt))
	if expiresAt.Before(time.Now().UTC()) {
		return Point11ValASignatureStateBlocked
	}
	return Point11ValASignatureStateActive
}

func EvaluatePoint11ValAAnchorState(model Point11ValAPolicyAnchorEnvelope) string {
	if !point11ValAAnchorRefValid(model.AnchorRef) ||
		!point11ValAAnchorTypeAllowed(model.AnchorType) ||
		!point11ValAPolicyRefValid(model.AnchoredSubjectRef) ||
		!point11ValAHashValid(model.AnchoredSubjectHash) ||
		!point11Val0ValidTimestamp(model.AnchorTimestamp) ||
		strings.TrimSpace(model.AnchorState) != point11ValAAnchorEnvelopeStateVerified ||
		strings.TrimSpace(model.AnchorVerificationResult) != point11ValAAnchorVerificationActive ||
		!point11Val0EvidenceRefsValid(model.AnchorEvidenceRefs) {
		return Point11ValAAnchorStateBlocked
	}
	return Point11ValAAnchorStateActive
}

func point11ValALifecycleBlockedEvaluation(reason string, diagnostics ...string) Point11ValALifecycleEvaluation {
	return Point11ValALifecycleEvaluation{
		TransitionState: Point11ValALifecycleTransitionStateBlocked,
		PolicyUseState:  Point11ValAPolicyUseStateBlocked,
		Reason:          reason,
		Diagnostics:     append([]string{}, diagnostics...),
	}
}

func point11ValALifecycleActiveEvaluation(reason, policyUseState string, diagnostics ...string) Point11ValALifecycleEvaluation {
	return Point11ValALifecycleEvaluation{
		TransitionState: Point11ValALifecycleTransitionStateActive,
		PolicyUseState:  policyUseState,
		Reason:          reason,
		Diagnostics:     append([]string{}, diagnostics...),
	}
}

func EvaluatePoint11ValALifecycleTransition(model Point11ValAPolicyLifecycleTransition) Point11ValALifecycleEvaluation {
	if !point11ValATransitionRefValid(model.TransitionID) ||
		!point11ValAPolicyLifecycleAllowed(model.FromState) ||
		!point11ValAPolicyLifecycleAllowed(model.ToState) ||
		!point11ValAPolicyRefValid(model.PolicyRef) ||
		!point11Val0IdentityValueValid(model.ActorRef) ||
		!point11Val0ValidTimestamp(model.TransitionTimestamp) ||
		!point11ValAAuditRefValid(model.AuditID) {
		return point11ValALifecycleBlockedEvaluation(
			"lifecycle_transition_identity_invalid",
			"invalid_transition_identity_or_timestamp",
		)
	}
	if strings.TrimSpace(model.RollbackRef) != "" && !point11ValAGovernanceEventRefValid(model.RollbackRef) {
		return point11ValALifecycleBlockedEvaluation(
			"lifecycle_transition_invalid_rollback_ref",
			"invalid_rollback_ref",
		)
	}
	if (model.PolicyRelaxationRequested || model.AuthorityExpansionRequested) && !point11ValAGovernanceEventRefValid(model.GovernanceEventRef) {
		return point11ValALifecycleBlockedEvaluation(
			"lifecycle_transition_missing_governance_event_for_relaxation_or_expansion",
			"missing_governance_event_for_relaxation_or_expansion",
		)
	}
	switch {
	case strings.TrimSpace(model.FromState) == point11ValAPolicyLifecycleDraft && strings.TrimSpace(model.ToState) == point11ValAPolicyLifecycleReviewRequired:
		if !point11Val0IdentityValueValid(model.Reason) {
			return point11ValALifecycleBlockedEvaluation(
				"draft_to_review_required_missing_reason",
				"missing_transition_reason",
			)
		}
		return point11ValALifecycleActiveEvaluation(
			"draft_to_review_required_transition_valid",
			Point11ValAPolicyUseStateNotYetActive,
			"policy_use_not_yet_active",
		)
	case strings.TrimSpace(model.FromState) == point11ValAPolicyLifecycleReviewRequired && strings.TrimSpace(model.ToState) == point11ValAPolicyLifecycleApproved:
		if !point11Val0IdentityValueValid(model.ApproverRef) ||
			!point11ValAGovernanceEventRefValid(model.GovernanceEventRef) ||
			!point11Val0EvidenceRefsValid(model.ApprovalEvidenceRefs) {
			return point11ValALifecycleBlockedEvaluation(
				"review_required_to_approved_transition_missing_approval_context",
				"missing_approver_or_governance_event_or_approval_evidence",
			)
		}
		return point11ValALifecycleActiveEvaluation(
			"review_required_to_approved_transition_valid",
			Point11ValAPolicyUseStateNotYetActive,
			"policy_use_not_yet_active",
		)
	case strings.TrimSpace(model.FromState) == point11ValAPolicyLifecycleApproved && strings.TrimSpace(model.ToState) == point11ValAPolicyLifecycleActive:
		if !point11ValAGovernanceEventRefValid(model.GovernanceEventRef) ||
			!point11Val0EvidenceRefsValid(model.ApprovalEvidenceRefs) {
			return point11ValALifecycleBlockedEvaluation(
				"approved_to_active_transition_missing_governance_activation_context",
				"missing_governance_event_or_activation_evidence",
			)
		}
		if !point11ValASignatureEnvelopeActive(model.SignatureEnvelopeState) ||
			!point11ValAAnchorEnvelopeActive(model.AnchorEnvelopeState) {
			return point11ValALifecycleBlockedEvaluation(
				"approved_to_active_transition_missing_active_signature_or_anchor",
				"signature_or_anchor_not_active",
			)
		}
		return point11ValALifecycleActiveEvaluation(
			"approved_to_active_transition_valid",
			Point11ValAPolicyUseStateActive,
			"policy_use_active",
		)
	case strings.TrimSpace(model.FromState) == point11ValAPolicyLifecycleActive && strings.TrimSpace(model.ToState) == point11ValAPolicyLifecycleDeprecated:
		if !point11ValAGovernanceEventRefValid(model.GovernanceEventRef) ||
			!point11Val0IdentityValueValid(model.Reason) ||
			!point11Val0EvidenceRefsValid(model.ApprovalEvidenceRefs) {
			return point11ValALifecycleBlockedEvaluation(
				"active_to_deprecated_transition_missing_governance_context",
				"missing_governance_event_or_reason_or_evidence",
			)
		}
		return point11ValALifecycleActiveEvaluation(
			"active_to_deprecated_transition_valid",
			Point11ValAPolicyUseStateHistoricalOnly,
			"policy_use_historical_only_due_to_deprecation",
		)
	case strings.TrimSpace(model.FromState) == point11ValAPolicyLifecycleActive && strings.TrimSpace(model.ToState) == point11ValAPolicyLifecycleSuperseded:
		if !point11ValAPolicyRefValid(model.SuccessorPolicyRef) ||
			!point11ValACompatibilityReviewRefValid(model.CompatibilityReviewRef) ||
			!point11ValAGovernanceEventRefValid(model.GovernanceEventRef) ||
			!point11Val0IdentityValueValid(model.Reason) ||
			!point11Val0EvidenceRefsValid(model.ApprovalEvidenceRefs) {
			return point11ValALifecycleBlockedEvaluation(
				"active_to_superseded_transition_missing_supersession_context",
				"missing_successor_or_compatibility_review_or_governance_event_or_reason_or_evidence",
			)
		}
		return point11ValALifecycleActiveEvaluation(
			"active_to_superseded_transition_valid",
			Point11ValAPolicyUseStateHistoricalOnly,
			"policy_use_historical_only_due_to_supersession",
		)
	case strings.TrimSpace(model.FromState) == point11ValAPolicyLifecycleActive && strings.TrimSpace(model.ToState) == point11ValAPolicyLifecycleRevoked:
		if !point11Val0IdentityValueValid(model.ApproverRef) ||
			!point11ValAGovernanceEventRefValid(model.GovernanceEventRef) ||
			!point11Val0IdentityValueValid(model.Reason) ||
			!point11Val0EvidenceRefsValid(model.ApprovalEvidenceRefs) {
			return point11ValALifecycleBlockedEvaluation(
				"active_to_revoked_transition_missing_revocation_context",
				"missing_approver_or_governance_event_or_reason_or_evidence",
			)
		}
		return point11ValALifecycleActiveEvaluation(
			"active_to_revoked_transition_valid",
			Point11ValAPolicyUseStateBlocked,
			"policy_use_blocked_by_revocation",
		)
	default:
		return point11ValALifecycleBlockedEvaluation(
			"unsupported_lifecycle_transition",
			"unsupported_lifecycle_transition",
		)
	}
}

func EvaluatePoint11ValALifecycleState(model Point11ValAPolicyLifecycleTransition) string {
	return EvaluatePoint11ValALifecycleTransition(model).TransitionState
}

func EvaluatePoint11ValAPolicyUseState(model Point11ValAPolicyLifecycleTransition) string {
	return EvaluatePoint11ValALifecycleTransition(model).PolicyUseState
}

func point11ValALifecycleGraphConsistencyDiagnostics(
	lifecycle Point11ValAPolicyLifecycleTransition,
	graph Point11ValAPolicySupersessionRevocationGraph,
) []string {
	diagnostics := []string{}
	fromState := strings.TrimSpace(lifecycle.FromState)
	toState := strings.TrimSpace(lifecycle.ToState)
	if fromState == point11ValAPolicyLifecycleActive && toState == point11ValAPolicyLifecycleSuperseded {
		if strings.TrimSpace(graph.SourcePolicyRef) != strings.TrimSpace(lifecycle.PolicyRef) {
			diagnostics = append(diagnostics, "graph_source_policy_ref_mismatch")
		}
		if strings.TrimSpace(graph.SuccessorPolicyRef) != strings.TrimSpace(lifecycle.SuccessorPolicyRef) {
			diagnostics = append(diagnostics, "graph_missing_or_mismatched_successor_policy_ref")
		}
		if strings.TrimSpace(graph.CompatibilityReviewRef) != strings.TrimSpace(lifecycle.CompatibilityReviewRef) {
			diagnostics = append(diagnostics, "graph_missing_or_mismatched_compatibility_review_ref")
		}
	}
	if fromState == point11ValAPolicyLifecycleActive && toState == point11ValAPolicyLifecycleRevoked {
		if strings.TrimSpace(graph.SourcePolicyRef) != strings.TrimSpace(lifecycle.PolicyRef) {
			diagnostics = append(diagnostics, "graph_source_policy_ref_mismatch")
		}
		if !point11ValARevokerRefValid(graph.RevokedByRef) {
			diagnostics = append(diagnostics, "graph_missing_or_invalid_revoked_by_ref")
		}
		if !point11Val0IdentityValueValid(graph.RevocationReason) {
			diagnostics = append(diagnostics, "graph_missing_revocation_reason")
		}
	}
	return diagnostics
}

func EvaluatePoint11ValAGraphState(model Point11ValAPolicySupersessionRevocationGraph) string {
	if !point11ValAGraphRefValid(model.GraphID) ||
		!point11ValAPolicyRefValid(model.SourcePolicyRef) ||
		!point11Val0ValidTimestamp(model.EffectiveFrom) ||
		!point11Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point11ValAAuditRefValid(model.AuditID) ||
		point11ValAGraphHasCycle(model.LineagePath) {
		return Point11ValAGraphStateBlocked
	}
	if model.SuccessorPolicyRef != "" {
		if !point11ValAPolicyRefValid(model.SuccessorPolicyRef) ||
			strings.TrimSpace(model.SuccessorPolicyRef) == strings.TrimSpace(model.SourcePolicyRef) ||
			!point11Val0IdentityValueValid(model.CompatibilityVersion) ||
			!point11ValACompatibilityReviewRefValid(model.CompatibilityReviewRef) ||
			!point11Val0IdentityValueValid(model.SupersessionReason) {
			return Point11ValAGraphStateBlocked
		}
		switch strings.TrimSpace(model.SuccessorLifecycleState) {
		case point11ValAPolicyLifecycleRevoked, point11ValAPolicyLifecycleExpired:
			return Point11ValAGraphStateBlocked
		}
	}
	if model.RevokedByRef != "" {
		if !point11ValARevokerRefValid(model.RevokedByRef) ||
			!point11Val0IdentityValueValid(model.RevocationReason) {
			return Point11ValAGraphStateBlocked
		}
	}
	if model.RevocationReason != "" && model.RevokedByRef == "" {
		return Point11ValAGraphStateBlocked
	}
	if model.SuccessorPolicyRef == "" && (model.CompatibilityVersion != "" || model.CompatibilityReviewRef != "" || model.SupersessionReason != "") {
		return Point11ValAGraphStateBlocked
	}
	return Point11ValAGraphStateActive
}

func point11ValABlockingReasons(model Point11ValAFoundation) []string {
	reasons := []string{}
	if model.DependencyState == Point11ValADependencyStateBlocked {
		reasons = append(reasons, "val0_dependency_blocked")
	}
	if model.RegistryState == Point11ValARegistryStateBlocked {
		reasons = append(reasons, "policy_registry_blocked")
	}
	if model.SignatureState == Point11ValASignatureStateBlocked {
		reasons = append(reasons, "policy_signature_blocked")
	}
	if model.AnchorState == Point11ValAAnchorStateBlocked {
		reasons = append(reasons, "policy_anchor_blocked")
	}
	if model.LifecycleTransitionState == Point11ValALifecycleTransitionStateBlocked {
		reasons = append(reasons, "policy_lifecycle_transition_blocked")
	}
	if model.PolicyUseState == Point11ValAPolicyUseStateHistoricalOnly {
		reasons = append(reasons, "policy_use_historical_only")
	}
	if model.PolicyUseState == Point11ValAPolicyUseStateBlocked ||
		model.PolicyUseState == Point11ValAPolicyUseStateNotYetActive {
		reasons = append(reasons, "policy_use_not_active")
	}
	if model.GraphState == Point11ValAGraphStateBlocked {
		reasons = append(reasons, "policy_graph_blocked")
	}
	if model.CreatesLegalRegulatoryCertificationClaim {
		reasons = append(reasons, "authority_claim_surface_blocked")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "publication_side_effects_blocked")
	}
	return reasons
}

func point11ValADiagnosticsModel(model Point11ValAFoundation) Point11ValADiagnostics {
	return Point11ValADiagnostics{
		CurrentState:         model.CurrentState,
		BlockingReasons:      append([]string{}, model.BlockingReasons...),
		ReviewPrerequisites:  append([]string{}, model.ReviewPrerequisites...),
		ComponentStates:      point11ValAComponentStates(model),
		LifecycleReason:      model.LifecycleEvaluation.Reason,
		LifecycleDiagnostics: append([]string{}, model.LifecycleEvaluation.Diagnostics...),
		ProjectionDisclaimer: model.ProjectionDisclaimer,
	}
}

func EvaluatePoint11ValAFoundationState(model Point11ValAFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.CreatesLegalRegulatoryCertificationClaim ||
		model.CreatesPublicationSideEffects ||
		strings.TrimSpace(model.RegistryState) == Point11ValARegistryStateBlocked ||
		strings.TrimSpace(model.SignatureState) == Point11ValASignatureStateBlocked ||
		strings.TrimSpace(model.AnchorState) == Point11ValAAnchorStateBlocked ||
		strings.TrimSpace(model.LifecycleTransitionState) == Point11ValALifecycleTransitionStateBlocked ||
		strings.TrimSpace(model.PolicyUseState) != Point11ValAPolicyUseStateActive ||
		strings.TrimSpace(model.GraphState) == Point11ValAGraphStateBlocked ||
		strings.TrimSpace(model.DependencyState) == Point11ValADependencyStateBlocked {
		return Point11ValAStateBlocked
	}
	if strings.TrimSpace(model.DependencyState) == Point11ValADependencyStateReviewRequired {
		return Point11ValAStateReviewRequired
	}
	return Point11ValAStateActive
}

func Point11ValAFoundationModel() Point11ValAFoundation {
	disclaimer := point11ValAProjectionDisclaimerBaseline
	policyRef := "point11_policy_authority_core_v1"
	signatureRef := "signature_point11_vala_policy_20260502"
	signingKeyRef := "signing_key_point11_vala_primary"
	anchorRef := "anchor_point11_vala_policy_20260502"
	return Point11ValAFoundation{
		DependencyState:                          Point11ValADependencyStateReviewRequired,
		RegistryState:                            Point11ValARegistryStateActive,
		SignatureState:                           Point11ValASignatureStateActive,
		AnchorState:                              Point11ValAAnchorStateActive,
		LifecycleState:                           Point11ValALifecycleTransitionStateActive,
		LifecycleTransitionState:                 Point11ValALifecycleTransitionStateActive,
		PolicyUseState:                           Point11ValAPolicyUseStateActive,
		GraphState:                               Point11ValAGraphStateActive,
		ProjectionDisclaimer:                     disclaimer,
		CreatesLegalRegulatoryCertificationClaim: false,
		CreatesPublicationSideEffects:            false,
		Dependency:                               point11ValADependencySnapshotModel(),
		Registry: Point11ValASignedPolicyRegistry{
			RegistryID:           "registry_point11_vala_policy_authority_core",
			RegistryVersion:      "point11_vala_registry_v1",
			PolicyPackID:         "policy_pack_point11_vala_core",
			PolicyID:             policyRef,
			PolicyVersion:        "point11_vala_policy_v1",
			PolicyScope:          "tenant_scoped_policy_authority_core",
			PolicyOwner:          "policy_owner_team",
			PolicyIssuer:         "governance_authority_team",
			ApprovalChainRefs:    []string{"governance_reviewer", "governance_final_approver"},
			SignatureRef:         signatureRef,
			SigningKeyRef:        signingKeyRef,
			SigningAlgorithm:     "ed25519",
			SignatureState:       point11ValASignatureEnvelopeStateVerified,
			AnchorRef:            anchorRef,
			AnchorState:          point11ValAAnchorEnvelopeStateVerified,
			SchemaVersion:        "point11_vala_schema_v1",
			CompatibilityVersion: "point11_vala_compat_v1",
			EffectiveFrom:        "2026-05-02T10:00:00Z",
			EffectiveUntil:       "2099-01-01T00:00:00Z",
			LifecycleState:       point11ValAPolicyLifecycleActive,
			GovernanceEventRef:   "governance_event_point11_vala_activation",
			ApprovalEvidenceRefs: []string{"evidence:point11-vala-policy-approval-001"},
			ProjectionDisclaimer: disclaimer,
		},
		Signature: Point11ValAPolicySignatureEnvelope{
			SignatureRef:             signatureRef,
			SigningKeyRef:            signingKeyRef,
			SigningAlgorithm:         "ed25519",
			SignedSubjectRef:         policyRef,
			SignedSubjectHash:        "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SignerIdentity:           "policy_signer_team",
			IssuedAt:                 "2026-05-02T10:00:00Z",
			ExpiresAt:                "2099-01-01T00:00:00Z",
			SignatureState:           point11ValASignatureEnvelopeStateVerified,
			VerificationResult:       point11ValASignatureVerificationActive,
			VerificationEvidenceRefs: []string{"evidence:point11-vala-signature-verification-001"},
		},
		Anchor: Point11ValAPolicyAnchorEnvelope{
			AnchorRef:                anchorRef,
			AnchorType:               "trusted_timestamp",
			AnchoredSubjectRef:       policyRef,
			AnchoredSubjectHash:      "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			AnchorTimestamp:          "2026-05-02T10:00:00Z",
			AnchorState:              point11ValAAnchorEnvelopeStateVerified,
			AnchorEvidenceRefs:       []string{"evidence:point11-vala-anchor-verification-001"},
			AnchorVerificationResult: point11ValAAnchorVerificationActive,
		},
		Lifecycle: Point11ValAPolicyLifecycleTransition{
			TransitionID:           "transition_point11_vala_approval_to_active",
			FromState:              point11ValAPolicyLifecycleApproved,
			ToState:                point11ValAPolicyLifecycleActive,
			PolicyRef:              policyRef,
			ActorRef:               "governance_operator",
			ApproverRef:            "governance_final_approver",
			GovernanceEventRef:     "governance_event_point11_vala_activation",
			Reason:                 "policy_activation_transition",
			TransitionTimestamp:    "2026-05-02T10:00:00Z",
			ApprovalEvidenceRefs:   []string{"evidence:point11-vala-policy-approval-001"},
			CompatibilityReviewRef: "compatibility_review_point11_vala_activation",
			RollbackRef:            "governance_event_point11_vala_rollback",
			AuditID:                "audit_point11_vala_transition_001",
			SignatureEnvelopeState: Point11ValASignatureStateActive,
			AnchorEnvelopeState:    Point11ValAAnchorStateActive,
		},
		Graph: Point11ValAPolicySupersessionRevocationGraph{
			GraphID:         "graph_point11_vala_policy_lineage",
			SourcePolicyRef: policyRef,
			EffectiveFrom:   "2026-05-02T10:00:00Z",
			EvidenceRefs:    []string{"evidence:point11-vala-graph-001"},
			AuditID:         "audit_point11_vala_graph_001",
			LineagePath:     []string{policyRef},
		},
	}
}

func ComputePoint11ValAFoundation(model Point11ValAFoundation) Point11ValAFoundation {
	model.DependencyState = EvaluatePoint11ValADependencyState(model.Dependency)
	model.RegistryState = EvaluatePoint11ValARegistryState(model.Registry)
	model.SignatureState = EvaluatePoint11ValASignatureState(model.Signature)
	model.AnchorState = EvaluatePoint11ValAAnchorState(model.Anchor)
	model.LifecycleEvaluation = EvaluatePoint11ValALifecycleTransition(model.Lifecycle)
	model.LifecycleTransitionState = model.LifecycleEvaluation.TransitionState
	model.LifecycleState = model.LifecycleEvaluation.TransitionState
	model.PolicyUseState = model.LifecycleEvaluation.PolicyUseState
	model.GraphState = EvaluatePoint11ValAGraphState(model.Graph)
	if consistencyDiagnostics := point11ValALifecycleGraphConsistencyDiagnostics(model.Lifecycle, model.Graph); len(consistencyDiagnostics) > 0 {
		model.GraphState = Point11ValAGraphStateBlocked
		model.LifecycleEvaluation.Diagnostics = append(model.LifecycleEvaluation.Diagnostics, consistencyDiagnostics...)
	}
	model.CurrentState = EvaluatePoint11ValAFoundationState(model)
	model.BlockingReasons = point11ValABlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	model.Diagnostics = point11ValADiagnosticsModel(model)
	return model
}
