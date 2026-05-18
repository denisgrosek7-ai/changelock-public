package formal

import (
	"strings"
	"time"
)

const (
	Point11ValBStateActive         = "point11_valb_claims_architecture_core_active"
	Point11ValBStateBlocked        = "point11_valb_claims_architecture_core_blocked"
	Point11ValBStateReviewRequired = "point11_valb_claims_architecture_core_review_required"

	Point11ValBDependencyStateActive         = "point11_valb_dependency_active"
	Point11ValBDependencyStateBlocked        = "point11_valb_dependency_blocked"
	Point11ValBDependencyStateReviewRequired = "point11_valb_dependency_review_required"

	Point11ValBClaimTypeStateActive  = "point11_valb_claim_type_active"
	Point11ValBClaimTypeStateBlocked = "point11_valb_claim_type_blocked"

	Point11ValBIssuanceRequestStateActive  = "point11_valb_issuance_request_active"
	Point11ValBIssuanceRequestStateBlocked = "point11_valb_issuance_request_blocked"

	Point11ValBIssuedClaimStateActive  = "point11_valb_issued_claim_active"
	Point11ValBIssuedClaimStateBlocked = "point11_valb_issued_claim_blocked"

	Point11ValBRegistryStateActive  = "point11_valb_claim_registry_active"
	Point11ValBRegistryStateBlocked = "point11_valb_claim_registry_blocked"

	Point11ValBVerificationStateActive  = "point11_valb_claim_verification_active"
	Point11ValBVerificationStateBlocked = "point11_valb_claim_verification_blocked"

	Point11ValBCrossDomainIntakeStateActive         = "point11_valb_cross_domain_intake_active"
	Point11ValBCrossDomainIntakeStateBlocked        = "point11_valb_cross_domain_intake_blocked"
	Point11ValBCrossDomainIntakeStateReviewRequired = "point11_valb_cross_domain_intake_review_required"
)

const (
	point11ValBProjectionDisclaimerBaseline       = "projection_only not_canonical_truth point11_valb_claims_architecture_core"
	point11ValBPolicyBasisStateActive             = "policy_basis_active"
	point11ValBAuthorityMatrixStateActive         = "authority_matrix_active"
	point11ValBVerificationMethodEvidenceLinked   = "evidence_linked_review"
	point11ValBVerificationMethodRegistryBound    = "registry_bound_verification"
	point11ValBVerificationMethodPolicyBound      = "policy_bound_review"
	point11ValBFreshnessClassCurrent              = "current"
	point11ValBFreshnessClassBounded              = "bounded"
	point11ValBFreshnessClassHistorical           = "historical"
	point11ValBClaimLifecycleActive               = "active"
	point11ValBClaimLifecycleExpired              = "expired"
	point11ValBVerificationResultVerified         = "verification_result_verified"
	point11ValBVerificationResultInactive         = "verification_result_inactive"
	point11ValBFreshnessResultActive              = "freshness_result_active"
	point11ValBFreshnessResultStale               = "freshness_result_stale"
	point11ValBSignatureOrRegistryResultActive    = "signature_or_registry_active"
	point11ValBSignatureOrRegistryResultBlocked   = "signature_or_registry_blocked"
	point11ValBRevocationCheckActive              = "claim_not_revoked"
	point11ValBRevocationCheckBlocked             = "claim_revoked"
	point11ValBSupersessionCheckActive            = "claim_not_superseded"
	point11ValBSupersessionCheckHistoricalOnly    = "claim_superseded_historical_only"
	point11ValBSupersessionCheckBlocked           = "claim_superseded"
	point11ValBScopeCheckActive                   = "scope_check_active"
	point11ValBScopeCheckBlocked                  = "scope_check_blocked"
	point11ValBAudienceCheckActive                = "audience_check_active"
	point11ValBAudienceCheckBlocked               = "audience_check_blocked"
	point11ValBIssuerTrustActive                  = "issuer_trust_active"
	point11ValBIssuerTrustBlocked                 = "issuer_trust_blocked"
	point11ValBCrossDomainCompatibilityActive     = "cross_domain_compatibility_active"
	point11ValBCrossDomainCompatibilityBlocked    = "cross_domain_compatibility_blocked"
	point11ValBCrossDomainCompatibilityReview     = "cross_domain_compatibility_review_required"
	point11ValBGenericCompatibilityActive         = "compatibility_active"
	point11ValBGenericCompatibilityReviewRequired = "compatibility_review_required"
	point11ValBGenericCompatibilityBlocked        = "compatibility_blocked"
	point11ValBRevocationHandlingResultActive     = "revocation_handling_active"
	point11ValBEvidenceTranslationResultActive    = "evidence_translation_active"
	point11ValBEvidenceTranslationResultReview    = "evidence_translation_review_required"
	point11ValBEvidenceTranslationResultBlocked   = "evidence_translation_blocked"
	point11ValBLocalAdmissibilityResultActive     = "local_admissibility_active"
	point11ValBLocalAdmissibilityResultReview     = "local_admissibility_review_required"
	point11ValBLocalAdmissibilityResultBlocked    = "local_admissibility_blocked"
)

type Point11ValBValAReviewContext struct {
	LocalReviewAllowsDependencyReviewRequired bool     `json:"local_review_allows_dependency_review_required"`
	ValAPoint11PassEmitted                    bool     `json:"vala_point11_pass_emitted"`
	ValACreatesAuthorityClaims                bool     `json:"vala_creates_authority_claims"`
	ValACreatesSigningSideEffects             bool     `json:"vala_creates_signing_side_effects"`
	ValACreatesAnchoringSideEffects           bool     `json:"vala_creates_anchoring_side_effects"`
	ValACreatesPublicationSideEffects         bool     `json:"vala_creates_publication_side_effects"`
	ValACreatesExternalAPISideEffects         bool     `json:"vala_creates_external_api_side_effects"`
	ValACreatesProductionSideEffects          bool     `json:"vala_creates_production_side_effects"`
	UnresolvedCrossPointTaxonomyDrift         bool     `json:"unresolved_cross_point_taxonomy_drift"`
	TaxonomyDriftReviewPrerequisite           bool     `json:"taxonomy_drift_review_prerequisite"`
	OpenCLB0Findings                          int      `json:"open_clb0_findings"`
	OpenCLB1Findings                          int      `json:"open_clb1_findings"`
	OpenCLB2Findings                          int      `json:"open_clb2_findings"`
	ReviewPrerequisites                       []string `json:"review_prerequisites,omitempty"`
}

type Point11ValBDependencySnapshot struct {
	ValACurrentState                  string   `json:"vala_current_state"`
	ValADependencyState               string   `json:"vala_dependency_state"`
	ValARegistryState                 string   `json:"vala_registry_state"`
	ValASignatureState                string   `json:"vala_signature_state"`
	ValAAnchorState                   string   `json:"vala_anchor_state"`
	ValALifecycleTransitionState      string   `json:"vala_lifecycle_transition_state"`
	ValAPolicyUseState                string   `json:"vala_policy_use_state"`
	ValAGraphState                    string   `json:"vala_graph_state"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
	ValAPoint11PassEmitted            bool     `json:"vala_point11_pass_emitted"`
	ValACreatesAuthorityClaims        bool     `json:"vala_creates_authority_claims"`
	ValACreatesSigningSideEffects     bool     `json:"vala_creates_signing_side_effects"`
	ValACreatesAnchoringSideEffects   bool     `json:"vala_creates_anchoring_side_effects"`
	ValACreatesPublicationSideEffects bool     `json:"vala_creates_publication_side_effects"`
	ValACreatesExternalAPISideEffects bool     `json:"vala_creates_external_api_side_effects"`
	ValACreatesProductionSideEffects  bool     `json:"vala_creates_production_side_effects"`
	UnresolvedCrossPointTaxonomyDrift bool     `json:"unresolved_cross_point_taxonomy_drift"`
	TaxonomyDriftReviewPrerequisite   bool     `json:"taxonomy_drift_review_prerequisite"`
	OpenCLB0Findings                  int      `json:"open_clb0_findings"`
	OpenCLB1Findings                  int      `json:"open_clb1_findings"`
	OpenCLB2Findings                  int      `json:"open_clb2_findings"`
	LocalReviewAllowsReviewRequired   bool     `json:"local_review_allows_review_required"`
	ReviewPrerequisites               []string `json:"review_prerequisites,omitempty"`
}

type Point11ValBClaimTypeDefinition struct {
	CurrentState                          string   `json:"current_state"`
	ClaimTypeID                           string   `json:"claim_type_id"`
	ClaimType                             string   `json:"claim_type"`
	Category                              string   `json:"category"`
	AllowedSubjectKinds                   []string `json:"allowed_subject_kinds,omitempty"`
	AllowedIssuerKinds                    []string `json:"allowed_issuer_kinds,omitempty"`
	AllowedAudiences                      []string `json:"allowed_audiences,omitempty"`
	AllowedPublicationSurfaces            []string `json:"allowed_publication_surfaces,omitempty"`
	RequiredPolicyBasis                   string   `json:"required_policy_basis"`
	RequiredEvidenceKinds                 []string `json:"required_evidence_kinds,omitempty"`
	RequiredVerificationMethod            string   `json:"required_verification_method"`
	RequiredFreshnessClass                string   `json:"required_freshness_class"`
	DefaultExpiryDuration                 string   `json:"default_expiry_duration"`
	RevocationRequired                    bool     `json:"revocation_required"`
	SupersessionAllowed                   bool     `json:"supersession_allowed"`
	CrossDomainAllowed                    bool     `json:"cross_domain_allowed"`
	AgentOriginAllowed                    bool     `json:"agent_origin_allowed"`
	CustomerVisibleAllowed                bool     `json:"customer_visible_allowed"`
	PublicSafeAllowed                     bool     `json:"public_safe_allowed"`
	CleanRoomIPRequired                   bool     `json:"clean_room_ip_required"`
	GovernanceEventRequired               bool     `json:"governance_event_required"`
	CrossDomainTrustCompatibilityRequired bool     `json:"cross_domain_trust_compatibility_required"`
	AgentPublishAuthority                 bool     `json:"agent_publish_authority"`
	AgentApproveAuthority                 bool     `json:"agent_approve_authority"`
	ProjectionDisclaimer                  string   `json:"projection_disclaimer"`
}

type Point11ValBClaimIssuanceRequest struct {
	CurrentState                     string   `json:"current_state"`
	IssuanceRequestID                string   `json:"issuance_request_id"`
	ClaimID                          string   `json:"claim_id"`
	ClaimTypeID                      string   `json:"claim_type_id"`
	ClaimTypeName                    string   `json:"claim_type_name"`
	SubjectRef                       string   `json:"subject_ref"`
	SubjectKind                      string   `json:"subject_kind"`
	IssuerRef                        string   `json:"issuer_ref"`
	IssuerKind                       string   `json:"issuer_kind"`
	ProposerRef                      string   `json:"proposer_ref"`
	ReviewerRef                      string   `json:"reviewer_ref"`
	ApproverRef                      string   `json:"approver_ref"`
	Audience                         string   `json:"audience"`
	PublicationSurface               string   `json:"publication_surface"`
	PolicyBasisRef                   string   `json:"policy_basis_ref"`
	PolicyBasisState                 string   `json:"policy_basis_state"`
	PolicyVersion                    string   `json:"policy_version"`
	EvidenceRefs                     []string `json:"evidence_refs,omitempty"`
	VerificationMethod               string   `json:"verification_method"`
	FreshnessClass                   string   `json:"freshness_class"`
	RequestedLifecycleState          string   `json:"requested_lifecycle_state"`
	RequestedClaimCategory           string   `json:"requested_claim_category"`
	CleanRoomIPReviewRef             string   `json:"clean_room_ip_review_ref"`
	GovernanceEventRef               string   `json:"governance_event_ref"`
	AuthorityMatrixRef               string   `json:"authority_matrix_ref"`
	AuthorityMatrixState             string   `json:"authority_matrix_state"`
	IssuedAt                         string   `json:"issued_at"`
	ExpiresAt                        string   `json:"expires_at"`
	ClaimTypeState                   string   `json:"claim_type_state"`
	ClaimTypeAllowedSubjectKinds     []string `json:"claim_type_allowed_subject_kinds,omitempty"`
	ClaimTypeAllowedIssuerKinds      []string `json:"claim_type_allowed_issuer_kinds,omitempty"`
	ClaimTypeAllowedAudiences        []string `json:"claim_type_allowed_audiences,omitempty"`
	ClaimTypeAllowedSurfaces         []string `json:"claim_type_allowed_surfaces,omitempty"`
	ClaimTypeRequiredVerification    string   `json:"claim_type_required_verification"`
	ClaimTypeRequiredFreshness       string   `json:"claim_type_required_freshness"`
	ClaimTypeGovernanceEventRequired bool     `json:"claim_type_governance_event_required"`
	ClaimTypeCleanRoomIPRequired     bool     `json:"claim_type_clean_room_ip_required"`
	ClaimTypeCustomerVisibleAllowed  bool     `json:"claim_type_customer_visible_allowed"`
	ClaimTypePublicSafeAllowed       bool     `json:"claim_type_public_safe_allowed"`
	ClaimTypeAgentOriginAllowed      bool     `json:"claim_type_agent_origin_allowed"`
	AgentOrigin                      bool     `json:"agent_origin"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type Point11ValBIssuedClaimRecord struct {
	CurrentState         string   `json:"current_state"`
	ClaimID              string   `json:"claim_id"`
	ClaimTypeID          string   `json:"claim_type_id"`
	ClaimVersion         string   `json:"claim_version"`
	SubjectRef           string   `json:"subject_ref"`
	IssuerRef            string   `json:"issuer_ref"`
	IssuerKind           string   `json:"issuer_kind"`
	Audience             string   `json:"audience"`
	Scope                string   `json:"scope"`
	PublicationSurface   string   `json:"publication_surface"`
	PolicyBasisRef       string   `json:"policy_basis_ref"`
	PolicyBasisState     string   `json:"policy_basis_state"`
	PolicyVersion        string   `json:"policy_version"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs     []string `json:"evidence_hash_refs,omitempty"`
	EvidenceHashRequired bool     `json:"evidence_hash_required"`
	VerificationMethod   string   `json:"verification_method"`
	VerificationResult   string   `json:"verification_result"`
	IssuedAt             string   `json:"issued_at"`
	ExpiresAt            string   `json:"expires_at"`
	LifecycleState       string   `json:"lifecycle_state"`
	ClaimCategory        string   `json:"claim_category"`
	RevocationRef        string   `json:"revocation_ref"`
	SupersededBy         string   `json:"superseded_by"`
	CorrectionRef        string   `json:"correction_ref"`
	GovernanceEventRef   string   `json:"governance_event_ref"`
	CleanRoomIPReviewRef string   `json:"clean_room_ip_review_ref"`
	AuditID              string   `json:"audit_id"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type Point11ValBClaimRegistry struct {
	CurrentState             string                         `json:"current_state"`
	RegistryID               string                         `json:"registry_id"`
	RegistryVersion          string                         `json:"registry_version"`
	RegisteredClaims         []Point11ValBIssuedClaimRecord `json:"registered_claims,omitempty"`
	ActiveClaimRefs          []string                       `json:"active_claim_refs,omitempty"`
	RevokedClaimRefs         []string                       `json:"revoked_claim_refs,omitempty"`
	SupersededClaimRefs      []string                       `json:"superseded_claim_refs,omitempty"`
	CorrectedClaimRefs       []string                       `json:"corrected_claim_refs,omitempty"`
	BlockedClaimRefs         []string                       `json:"blocked_claim_refs,omitempty"`
	DuplicateClaimRefs       []string                       `json:"duplicate_claim_refs,omitempty"`
	ConflictingClaimRefs     []string                       `json:"conflicting_claim_refs,omitempty"`
	RegistryPolicyBasisRef   string                         `json:"registry_policy_basis_ref"`
	RegistryPolicyBasisState string                         `json:"registry_policy_basis_state"`
	GovernanceEventRef       string                         `json:"governance_event_ref"`
	AuditID                  string                         `json:"audit_id"`
	ProjectionDisclaimer     string                         `json:"projection_disclaimer"`
}

type Point11ValBClaimVerificationResult struct {
	CurrentState                   string   `json:"current_state"`
	VerificationID                 string   `json:"verification_id"`
	ClaimRef                       string   `json:"claim_ref"`
	ClaimVersion                   string   `json:"claim_version"`
	VerifierRef                    string   `json:"verifier_ref"`
	VerifierKind                   string   `json:"verifier_kind"`
	PolicyBasisRef                 string   `json:"policy_basis_ref"`
	PolicyBasisState               string   `json:"policy_basis_state"`
	EvidenceRefs                   []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs               []string `json:"evidence_hash_refs,omitempty"`
	EvidenceHashRequired           bool     `json:"evidence_hash_required"`
	VerificationMethod             string   `json:"verification_method"`
	VerificationTimestamp          string   `json:"verification_timestamp"`
	FreshnessResult                string   `json:"freshness_result"`
	SignatureOrRegistryResult      string   `json:"signature_or_registry_result"`
	RevocationCheckResult          string   `json:"revocation_check_result"`
	SupersessionCheckResult        string   `json:"supersession_check_result"`
	ScopeCheckResult               string   `json:"scope_check_result"`
	AudienceCheckResult            string   `json:"audience_check_result"`
	IssuerTrustResult              string   `json:"issuer_trust_result"`
	CrossDomainCompatibilityResult string   `json:"cross_domain_compatibility_result"`
	ResultState                    string   `json:"result_state"`
	Diagnostics                    []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer           string   `json:"projection_disclaimer"`
	ClaimRegistered                bool     `json:"claim_registered"`
	ClaimRegistryState             string   `json:"claim_registry_state"`
	ClaimEvidenceRefs              []string `json:"claim_evidence_refs,omitempty"`
	ClaimEvidenceHashRefs          []string `json:"claim_evidence_hash_refs,omitempty"`
	ClaimLifecycleState            string   `json:"claim_lifecycle_state"`
	HistoricalVerificationAllowed  bool     `json:"historical_verification_allowed"`
	CrossDomainRequired            bool     `json:"cross_domain_required"`
}

type Point11ValBCrossDomainClaimIntake struct {
	CurrentState                  string   `json:"current_state"`
	RemoteClaimRef                string   `json:"remote_claim_ref"`
	RemoteClaimType               string   `json:"remote_claim_type"`
	RemoteClaimState              string   `json:"remote_claim_state"`
	RemoteIssuerRef               string   `json:"remote_issuer_ref"`
	RemoteTrustRootRef            string   `json:"remote_trust_root_ref"`
	AcceptedClaimTypes            []string `json:"accepted_claim_types,omitempty"`
	LocalPolicyBasisRef           string   `json:"local_policy_basis_ref"`
	LocalPolicyBasisState         string   `json:"local_policy_basis_state"`
	CompatibilityRuleRef          string   `json:"compatibility_rule_ref"`
	ScopeCompatibilityResult      string   `json:"scope_compatibility_result"`
	FreshnessCompatibilityResult  string   `json:"freshness_compatibility_result"`
	RevocationHandlingResult      string   `json:"revocation_handling_result"`
	IssuerTrustResult             string   `json:"issuer_trust_result"`
	EvidenceTranslationResult     string   `json:"evidence_translation_result"`
	LocalAdmissibilityResult      string   `json:"local_admissibility_result"`
	Diagnostics                   []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	RemoteOverridesLocalPolicy    bool     `json:"remote_overrides_local_policy"`
	TreatsRemoteClaimAsAuthority  bool     `json:"treats_remote_claim_as_authority"`
	CreatesLegalAuthority         bool     `json:"creates_legal_authority"`
	CreatesRegulatoryAuthority    bool     `json:"creates_regulatory_authority"`
	CreatesCertificationAuthority bool     `json:"creates_certification_authority"`
}

type Point11ValBDiagnostics struct {
	CurrentState         string   `json:"current_state"`
	BlockingReasons      []string `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites  []string `json:"review_prerequisites,omitempty"`
	ComponentStates      []string `json:"component_states,omitempty"`
	DependencyReasons    []string `json:"dependency_reasons,omitempty"`
	ClaimTypeReasons     []string `json:"claim_type_reasons,omitempty"`
	IssuanceReasons      []string `json:"issuance_reasons,omitempty"`
	IssuedClaimReasons   []string `json:"issued_claim_reasons,omitempty"`
	RegistryReasons      []string `json:"registry_reasons,omitempty"`
	VerificationReasons  []string `json:"verification_reasons,omitempty"`
	CrossDomainReasons   []string `json:"cross_domain_reasons,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type Point11ValBFoundation struct {
	CurrentState                  string                             `json:"current_state"`
	BlockingReasons               []string                           `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites           []string                           `json:"review_prerequisites,omitempty"`
	DependencyState               string                             `json:"dependency_state"`
	ClaimTypeState                string                             `json:"claim_type_state"`
	IssuanceRequestState          string                             `json:"issuance_request_state"`
	IssuedClaimState              string                             `json:"issued_claim_state"`
	RegistryState                 string                             `json:"registry_state"`
	VerificationState             string                             `json:"verification_state"`
	CrossDomainIntakeState        string                             `json:"cross_domain_intake_state"`
	Diagnostics                   Point11ValBDiagnostics             `json:"diagnostics"`
	ProjectionDisclaimer          string                             `json:"projection_disclaimer"`
	CreatesAuthorityClaims        bool                               `json:"creates_authority_claims"`
	CreatesPublicationSideEffects bool                               `json:"creates_publication_side_effects"`
	Dependency                    Point11ValBDependencySnapshot      `json:"dependency"`
	ClaimTypeDefinition           Point11ValBClaimTypeDefinition     `json:"claim_type_definition"`
	IssuanceRequest               Point11ValBClaimIssuanceRequest    `json:"issuance_request"`
	IssuedClaim                   Point11ValBIssuedClaimRecord       `json:"issued_claim"`
	Registry                      Point11ValBClaimRegistry           `json:"registry"`
	Verification                  Point11ValBClaimVerificationResult `json:"verification"`
	CrossDomainIntake             Point11ValBCrossDomainClaimIntake  `json:"cross_domain_intake"`
}

func point11ValBClaimTypeRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"claim_type_",
		"point11_claim_type_",
	})
}

func point11ValBClaimRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"claim_",
		"point11_claim_",
	})
}

func point11ValBSubjectRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"subject_",
		"point11_subject_",
	})
}

func point11ValBIssuerRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"issuer_",
		"point11_issuer_",
	})
}

func point11ValBVerificationRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"verification_",
		"point11_verification_",
	})
}

func point11ValBClaimRegistryRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"claim_registry_",
		"point11_claim_registry_",
	})
}

func point11ValBAuthorityMatrixRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"authority_matrix_",
		"point11_authority_matrix_",
	})
}

func point11ValBGovernanceEventRefValid(value string) bool {
	return point11ValAGovernanceEventRefValid(value)
}

func point11ValBCleanRoomReviewRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"clean_room_review_",
		"point11_clean_room_review_",
	})
}

func point11ValBAuditRefValid(value string) bool {
	return point11ValAAuditRefValid(value)
}

func point11ValBTrustRootRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"trust_root_",
		"point11_trust_root_",
	})
}

func point11ValBCompatibilityRuleRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"compatibility_rule_",
		"point11_compatibility_rule_",
	})
}

func point11ValBEvidenceHashRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"evidence_hash_",
		"point11_evidence_hash_",
	})
}

func point11ValBExplicitValuesValid(values []string) bool {
	return point11Val0AllValuesValid(values)
}

func point11ValBPublicationSurfacesValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11Val0ContainsTrimmed(point11Val0PublicationSurfaces(), value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValBClaimCategoryAllowed(value string) bool {
	return point11Val0ContainsTrimmed(point11Val0AllowedClaimCategories(), value)
}

func point11ValBVerificationMethodSupported(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValBVerificationMethodEvidenceLinked,
		point11ValBVerificationMethodRegistryBound,
		point11ValBVerificationMethodPolicyBound,
	}, value)
}

func point11ValBFreshnessClassSupported(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValBFreshnessClassCurrent,
		point11ValBFreshnessClassBounded,
		point11ValBFreshnessClassHistorical,
	}, value)
}

func point11ValBFreshnessClassCompatible(required, actual string) bool {
	return required == actual
}

func point11ValBClaimLifecycleAllowed(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		Point11Val0ClaimLifecycleDraft,
		Point11Val0ClaimLifecycleReviewRequired,
		Point11Val0ClaimLifecycleApproved,
		Point11Val0ClaimLifecyclePublished,
		point11ValBClaimLifecycleActive,
		Point11Val0ClaimLifecycleRevoked,
		Point11Val0ClaimLifecycleSuperseded,
		Point11Val0ClaimLifecycleCorrected,
		Point11Val0ClaimLifecycleBlocked,
		point11ValBClaimLifecycleExpired,
	}, value)
}

func point11ValBClaimLifecycleActiveUseEligible(value string) bool {
	switch value {
	case Point11Val0ClaimLifecycleApproved, Point11Val0ClaimLifecyclePublished, point11ValBClaimLifecycleActive:
		return true
	default:
		return false
	}
}

func point11ValBClaimLifecycleInvalidated(value string) bool {
	switch value {
	case Point11Val0ClaimLifecycleRevoked,
		Point11Val0ClaimLifecycleSuperseded,
		Point11Val0ClaimLifecycleCorrected,
		Point11Val0ClaimLifecycleBlocked,
		point11ValBClaimLifecycleExpired:
		return true
	default:
		return false
	}
}

func point11ValBClaimNeedsGovernanceEvent(category, surface string) bool {
	return point11Val0PublicFacingSurface(surface) ||
		category == Point11Val0ClaimCategoryCustomerVisible ||
		category == Point11Val0ClaimCategoryPublicSafe
}

func point11ValBClaimNeedsCleanRoomReview(category, surface string) bool {
	return point11ValBClaimNeedsGovernanceEvent(category, surface)
}

func point11ValBStringSetKey(values []string) string {
	seen := map[string]struct{}{}
	ordered := []string{}
	for _, value := range values {
		if value == "" || value != strings.TrimSpace(value) || strings.ContainsAny(value, "\t\r\n") {
			return ""
		}
		if _, exists := seen[value]; exists {
			return ""
		}
		seen[value] = struct{}{}
		ordered = append(ordered, value)
	}
	if len(ordered) == 0 {
		return ""
	}
	for index := 0; index < len(ordered); index++ {
		for next := index + 1; next < len(ordered); next++ {
			if ordered[next] < ordered[index] {
				ordered[index], ordered[next] = ordered[next], ordered[index]
			}
		}
	}
	return strings.Join(ordered, ",")
}

func point11ValBExactStringSetMatch(left, right []string) bool {
	return point11ValBStringSetKey(left) != "" &&
		point11ValBStringSetKey(left) == point11ValBStringSetKey(right)
}

func point11ValBClaimRefListValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11ValBClaimRefValid(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValBEvidenceHashRefsValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11ValBEvidenceHashRefValid(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValBOptionalEvidenceHashRefsValid(required bool, verificationHashes, claimHashes []string) bool {
	verificationPresent := len(verificationHashes) > 0
	claimPresent := len(claimHashes) > 0
	if required {
		return point11ValBEvidenceHashRefsValid(verificationHashes) &&
			point11ValBEvidenceHashRefsValid(claimHashes) &&
			point11ValBExactStringSetMatch(verificationHashes, claimHashes)
	}
	if verificationPresent != claimPresent {
		return false
	}
	if !verificationPresent {
		return true
	}
	return point11ValBEvidenceHashRefsValid(verificationHashes) &&
		point11ValBEvidenceHashRefsValid(claimHashes) &&
		point11ValBExactStringSetMatch(verificationHashes, claimHashes)
}

func point11ValBComponentStates(model Point11ValBFoundation) []string {
	return []string{
		"dependency:" + model.DependencyState,
		"claim_type:" + model.ClaimTypeState,
		"issuance_request:" + model.IssuanceRequestState,
		"issued_claim:" + model.IssuedClaimState,
		"registry:" + model.RegistryState,
		"verification:" + model.VerificationState,
		"cross_domain_intake:" + model.CrossDomainIntakeState,
	}
}

// Dependency snapshots must copy actual computed upstream output.
// They must not repair, replace, fallback, or regenerate upstream dependency values.
// The dependency evaluator is responsible for fail-closed validation.
func SnapshotPoint11ValBDependencyFromComputedValA(valA Point11ValAFoundation, review Point11ValBValAReviewContext) Point11ValBDependencySnapshot {
	reviewPrerequisites := append([]string{}, valA.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point11ValBDependencySnapshot{
		ValACurrentState:                  valA.CurrentState,
		ValADependencyState:               valA.DependencyState,
		ValARegistryState:                 valA.RegistryState,
		ValASignatureState:                valA.SignatureState,
		ValAAnchorState:                   valA.AnchorState,
		ValALifecycleTransitionState:      valA.LifecycleTransitionState,
		ValAPolicyUseState:                valA.PolicyUseState,
		ValAGraphState:                    valA.GraphState,
		ProjectionDisclaimer:              valA.ProjectionDisclaimer,
		ValAPoint11PassEmitted:            review.ValAPoint11PassEmitted,
		ValACreatesAuthorityClaims:        review.ValACreatesAuthorityClaims,
		ValACreatesSigningSideEffects:     review.ValACreatesSigningSideEffects,
		ValACreatesAnchoringSideEffects:   review.ValACreatesAnchoringSideEffects,
		ValACreatesPublicationSideEffects: review.ValACreatesPublicationSideEffects,
		ValACreatesExternalAPISideEffects: review.ValACreatesExternalAPISideEffects,
		ValACreatesProductionSideEffects:  review.ValACreatesProductionSideEffects,
		UnresolvedCrossPointTaxonomyDrift: review.UnresolvedCrossPointTaxonomyDrift,
		TaxonomyDriftReviewPrerequisite:   review.TaxonomyDriftReviewPrerequisite,
		OpenCLB0Findings:                  review.OpenCLB0Findings,
		OpenCLB1Findings:                  review.OpenCLB1Findings,
		OpenCLB2Findings:                  review.OpenCLB2Findings,
		LocalReviewAllowsReviewRequired:   review.LocalReviewAllowsDependencyReviewRequired,
		ReviewPrerequisites:               reviewPrerequisites,
	}
}

func point11ValBDependencyReviewContextModel() Point11ValBValAReviewContext {
	return Point11ValBValAReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
		ValAPoint11PassEmitted:                    false,
		ValACreatesAuthorityClaims:                false,
		ValACreatesSigningSideEffects:             false,
		ValACreatesAnchoringSideEffects:           false,
		ValACreatesPublicationSideEffects:         false,
		ValACreatesExternalAPISideEffects:         false,
		ValACreatesProductionSideEffects:          false,
		UnresolvedCrossPointTaxonomyDrift:         false,
		TaxonomyDriftReviewPrerequisite:           false,
		OpenCLB0Findings:                          0,
		OpenCLB1Findings:                          0,
		OpenCLB2Findings:                          0,
	}
}

func point11ValBDependencySnapshotModel() Point11ValBDependencySnapshot {
	valA := ComputePoint11ValAFoundation(Point11ValAFoundationModel())
	return SnapshotPoint11ValBDependencyFromComputedValA(valA, point11ValBDependencyReviewContextModel())
}

func point11ValBDependencyStateAndReasons(model Point11ValBDependencySnapshot) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "vala_projection_disclaimer_blocked")
	}
	if model.ValAPoint11PassEmitted {
		reasons = append(reasons, "vala_point11_pass_emitted")
	}
	if model.ValACreatesAuthorityClaims {
		reasons = append(reasons, "vala_authority_claim_surface_blocked")
	}
	if model.ValACreatesSigningSideEffects {
		reasons = append(reasons, "vala_signing_side_effects_blocked")
	}
	if model.ValACreatesAnchoringSideEffects {
		reasons = append(reasons, "vala_anchoring_side_effects_blocked")
	}
	if model.ValACreatesPublicationSideEffects {
		reasons = append(reasons, "vala_publication_side_effects_blocked")
	}
	if model.ValACreatesExternalAPISideEffects {
		reasons = append(reasons, "vala_external_api_side_effects_blocked")
	}
	if model.ValACreatesProductionSideEffects {
		reasons = append(reasons, "vala_production_side_effects_blocked")
	}
	if model.OpenCLB0Findings > 0 {
		reasons = append(reasons, "vala_open_clb0_findings")
	}
	if model.OpenCLB1Findings > 0 {
		reasons = append(reasons, "vala_open_clb1_findings")
	}
	if model.OpenCLB2Findings > 0 {
		reasons = append(reasons, "vala_open_clb2_findings")
	}
	if model.ValARegistryState != Point11ValARegistryStateActive {
		reasons = append(reasons, "vala_registry_not_active")
	}
	if model.ValASignatureState != Point11ValASignatureStateActive {
		reasons = append(reasons, "vala_signature_not_active")
	}
	if model.ValAAnchorState != Point11ValAAnchorStateActive {
		reasons = append(reasons, "vala_anchor_not_active")
	}
	if model.ValALifecycleTransitionState != Point11ValALifecycleTransitionStateActive {
		reasons = append(reasons, "vala_lifecycle_transition_not_active")
	}
	if model.ValAPolicyUseState != Point11ValAPolicyUseStateActive {
		reasons = append(reasons, "vala_policy_use_not_active")
	}
	if model.ValAGraphState != Point11ValAGraphStateActive {
		reasons = append(reasons, "vala_graph_not_active")
	}
	// Hard blockers always win. Review prerequisites are considered only after
	// the upstream dependency snapshot has no local blocking conditions.
	if len(reasons) > 0 {
		return Point11ValBDependencyStateBlocked, reasons
	}
	if model.UnresolvedCrossPointTaxonomyDrift {
		if model.TaxonomyDriftReviewPrerequisite {
			return Point11ValBDependencyStateReviewRequired, []string{"cross_point_taxonomy_drift_review_prerequisite"}
		}
		return Point11ValBDependencyStateBlocked, []string{"cross_point_taxonomy_drift_unresolved"}
	}
	if model.ValACurrentState == Point11ValAStateActive &&
		model.ValADependencyState == Point11ValADependencyStateActive &&
		len(model.ReviewPrerequisites) == 0 {
		return Point11ValBDependencyStateActive, nil
	}
	if model.LocalReviewAllowsReviewRequired &&
		(len(model.ReviewPrerequisites) > 0 ||
			model.ValACurrentState == Point11ValAStateReviewRequired ||
			model.ValADependencyState == Point11ValADependencyStateReviewRequired) {
		return Point11ValBDependencyStateReviewRequired, []string{"vala_dependency_review_required"}
	}
	return Point11ValBDependencyStateBlocked, []string{"vala_dependency_not_active"}
}

func EvaluatePoint11ValBDependencyState(model Point11ValBDependencySnapshot) string {
	state, _ := point11ValBDependencyStateAndReasons(model)
	return state
}

func point11ValBClaimTypeStateAndReasons(model Point11ValBClaimTypeDefinition) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "claim_type_projection_disclaimer_blocked")
	}
	if !point11ValBClaimTypeRefValid(model.ClaimTypeID) {
		reasons = append(reasons, "claim_type_id_invalid")
	}
	if !point11Val0IdentityValueValid(model.ClaimType) || point11Val0ContainsForbiddenClaim(model.ClaimType) {
		reasons = append(reasons, "claim_type_name_invalid_or_overclaim")
	}
	if !point11ValBClaimCategoryAllowed(model.Category) {
		reasons = append(reasons, "claim_type_category_invalid")
	}
	if !point11ValBExplicitValuesValid(model.AllowedSubjectKinds) {
		reasons = append(reasons, "claim_type_allowed_subject_kinds_invalid")
	}
	if !point11ValBExplicitValuesValid(model.AllowedIssuerKinds) {
		reasons = append(reasons, "claim_type_allowed_issuer_kinds_invalid")
	}
	if !point11ValBExplicitValuesValid(model.AllowedAudiences) {
		reasons = append(reasons, "claim_type_allowed_audiences_invalid")
	}
	if !point11ValBPublicationSurfacesValid(model.AllowedPublicationSurfaces) {
		reasons = append(reasons, "claim_type_allowed_publication_surfaces_invalid")
	}
	if !point11ValAPolicyRefValid(model.RequiredPolicyBasis) {
		reasons = append(reasons, "claim_type_required_policy_basis_invalid")
	}
	if !point11ValBExplicitValuesValid(model.RequiredEvidenceKinds) {
		reasons = append(reasons, "claim_type_required_evidence_kinds_invalid")
	}
	if !point11ValBVerificationMethodSupported(model.RequiredVerificationMethod) {
		reasons = append(reasons, "claim_type_required_verification_method_invalid")
	}
	if !point11ValBFreshnessClassSupported(model.RequiredFreshnessClass) {
		reasons = append(reasons, "claim_type_required_freshness_class_invalid")
	}
	if !point11Val0IdentityValueValid(model.DefaultExpiryDuration) {
		reasons = append(reasons, "claim_type_default_expiry_duration_invalid")
	}
	if point11Val0PublicFacingSurfaceSet(model.AllowedPublicationSurfaces) && !model.CleanRoomIPRequired {
		reasons = append(reasons, "claim_type_public_surface_missing_clean_room_requirement")
	}
	if model.CustomerVisibleAllowed && !model.CleanRoomIPRequired {
		reasons = append(reasons, "claim_type_customer_visible_missing_clean_room_requirement")
	}
	if (model.CustomerVisibleAllowed || model.PublicSafeAllowed) && !model.GovernanceEventRequired {
		reasons = append(reasons, "claim_type_customer_or_public_missing_governance_event_requirement")
	}
	if model.CrossDomainAllowed && !model.CrossDomainTrustCompatibilityRequired {
		reasons = append(reasons, "claim_type_cross_domain_missing_trust_compatibility_requirement")
	}
	if model.AgentOriginAllowed && (model.AgentPublishAuthority || model.AgentApproveAuthority) {
		reasons = append(reasons, "claim_type_agent_origin_implies_publish_or_approve_authority")
	}
	if len(reasons) > 0 {
		return Point11ValBClaimTypeStateBlocked, reasons
	}
	return Point11ValBClaimTypeStateActive, nil
}

func EvaluatePoint11ValBClaimTypeState(model Point11ValBClaimTypeDefinition) string {
	state, _ := point11ValBClaimTypeStateAndReasons(model)
	return state
}

func point11ValBIssuanceRequestStateAndReasons(model Point11ValBClaimIssuanceRequest) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "issuance_projection_disclaimer_blocked")
	}
	if !point11ValBVerificationRefValid(model.IssuanceRequestID) {
		reasons = append(reasons, "issuance_request_id_invalid")
	}
	if !point11ValBClaimRefValid(model.ClaimID) {
		reasons = append(reasons, "issuance_claim_id_invalid")
	}
	if !point11ValBClaimTypeRefValid(model.ClaimTypeID) {
		reasons = append(reasons, "issuance_claim_type_id_invalid")
	}
	if !point11Val0IdentityValueValid(model.ClaimTypeName) || point11Val0ContainsForbiddenClaim(model.ClaimTypeName) {
		reasons = append(reasons, "issuance_claim_type_name_invalid_or_overclaim")
	}
	if !point11ValBSubjectRefValid(model.SubjectRef) {
		reasons = append(reasons, "issuance_subject_ref_invalid")
	}
	if !point11ValBIssuerRefValid(model.IssuerRef) {
		reasons = append(reasons, "issuance_issuer_ref_invalid")
	}
	if !point11Val0IdentityValueValid(model.SubjectKind) || !point11Val0ContainsTrimmed(model.ClaimTypeAllowedSubjectKinds, model.SubjectKind) {
		reasons = append(reasons, "issuance_subject_kind_not_allowed")
	}
	if !point11Val0IdentityValueValid(model.IssuerKind) || !point11Val0ContainsTrimmed(model.ClaimTypeAllowedIssuerKinds, model.IssuerKind) {
		reasons = append(reasons, "issuance_issuer_kind_not_allowed")
	}
	if !point11Val0IdentityValueValid(model.ProposerRef) ||
		!point11Val0IdentityValueValid(model.ReviewerRef) ||
		!point11Val0IdentityValueValid(model.ApproverRef) {
		reasons = append(reasons, "issuance_actor_refs_invalid")
	}
	if !point11Val0IdentityValueValid(model.Audience) || !point11Val0ContainsTrimmed(model.ClaimTypeAllowedAudiences, model.Audience) {
		reasons = append(reasons, "issuance_audience_not_allowed")
	}
	if !point11Val0ContainsTrimmed(model.ClaimTypeAllowedSurfaces, model.PublicationSurface) {
		reasons = append(reasons, "issuance_publication_surface_not_allowed")
	}
	if !point11ValAPolicyRefValid(model.PolicyBasisRef) {
		reasons = append(reasons, "issuance_policy_basis_ref_invalid")
	}
	if model.PolicyBasisState != point11ValBPolicyBasisStateActive {
		reasons = append(reasons, "issuance_policy_basis_not_active")
	}
	if !point11Val0IdentityValueValid(model.PolicyVersion) {
		reasons = append(reasons, "issuance_policy_version_invalid")
	}
	if !point11Val0EvidenceRefsValid(model.EvidenceRefs) {
		reasons = append(reasons, "issuance_evidence_refs_invalid")
	}
	if !point11ValBVerificationMethodSupported(model.VerificationMethod) ||
		model.VerificationMethod != model.ClaimTypeRequiredVerification {
		reasons = append(reasons, "issuance_verification_method_invalid")
	}
	if !point11ValBFreshnessClassSupported(model.FreshnessClass) ||
		!point11ValBFreshnessClassCompatible(model.ClaimTypeRequiredFreshness, model.FreshnessClass) {
		reasons = append(reasons, "issuance_freshness_class_invalid")
	}
	if !point11ValBClaimLifecycleAllowed(model.RequestedLifecycleState) ||
		point11ValBClaimLifecycleInvalidated(model.RequestedLifecycleState) {
		reasons = append(reasons, "issuance_requested_lifecycle_invalid")
	}
	if !point11ValBClaimCategoryAllowed(model.RequestedClaimCategory) ||
		point11Val0ContainsForbiddenClaim(model.RequestedClaimCategory, model.Audience) {
		reasons = append(reasons, "issuance_requested_category_or_audience_invalid")
	}
	if !point11ValBAuthorityMatrixRefValid(model.AuthorityMatrixRef) ||
		model.AuthorityMatrixState != point11ValBAuthorityMatrixStateActive {
		reasons = append(reasons, "issuance_authority_matrix_invalid")
	}
	if !point11Val0ValidTimestamp(model.IssuedAt) || !point11Val0ValidTimestamp(model.ExpiresAt) {
		reasons = append(reasons, "issuance_timestamp_invalid")
	} else {
		expiresAt, _ := time.Parse(time.RFC3339, model.ExpiresAt)
		if expiresAt.Before(time.Now().UTC()) {
			reasons = append(reasons, "issuance_expired")
		}
	}
	if model.ClaimTypeState != Point11ValBClaimTypeStateActive {
		reasons = append(reasons, "issuance_claim_type_not_active")
	}
	if model.RequestedClaimCategory == Point11Val0ClaimCategoryCustomerVisible &&
		!model.ClaimTypeCustomerVisibleAllowed {
		reasons = append(reasons, "issuance_customer_visible_category_not_allowed")
	}
	if model.RequestedClaimCategory == Point11Val0ClaimCategoryPublicSafe &&
		!model.ClaimTypePublicSafeAllowed {
		reasons = append(reasons, "issuance_public_safe_category_not_allowed")
	}
	if model.AgentOrigin && !model.ClaimTypeAgentOriginAllowed {
		reasons = append(reasons, "issuance_agent_origin_not_allowed")
	}
	if point11ValBClaimNeedsCleanRoomReview(model.RequestedClaimCategory, model.PublicationSurface) &&
		(!model.ClaimTypeCleanRoomIPRequired || !point11ValBCleanRoomReviewRefValid(model.CleanRoomIPReviewRef)) {
		reasons = append(reasons, "issuance_missing_clean_room_review")
	}
	if point11ValBClaimNeedsGovernanceEvent(model.RequestedClaimCategory, model.PublicationSurface) &&
		(!model.ClaimTypeGovernanceEventRequired || !point11ValBGovernanceEventRefValid(model.GovernanceEventRef)) {
		reasons = append(reasons, "issuance_missing_governance_event")
	}
	if model.ProposerRef == model.ApproverRef {
		if point11ValBClaimNeedsGovernanceEvent(model.RequestedClaimCategory, model.PublicationSurface) ||
			model.PublicationSurface == point11Val0PublicationSurfaceExport {
			reasons = append(reasons, "issuance_proposer_equals_final_approver")
		}
		switch point11Val0ActorClass(model.ProposerRef) {
		case "partner":
			reasons = append(reasons, "issuance_partner_self_approval")
		case "customer":
			reasons = append(reasons, "issuance_customer_self_approval")
		case "agent":
			reasons = append(reasons, "issuance_agent_self_approval")
		}
	}
	if model.AgentOrigin &&
		(model.RequestedClaimCategory == Point11Val0ClaimCategoryCustomerVisible ||
			model.RequestedClaimCategory == Point11Val0ClaimCategoryPublicSafe) &&
		!point11ValBGovernanceEventRefValid(model.GovernanceEventRef) {
		reasons = append(reasons, "issuance_agent_origin_missing_governance_event")
	}
	if len(reasons) > 0 {
		return Point11ValBIssuanceRequestStateBlocked, reasons
	}
	return Point11ValBIssuanceRequestStateActive, nil
}

func EvaluatePoint11ValBIssuanceRequestState(model Point11ValBClaimIssuanceRequest) string {
	state, _ := point11ValBIssuanceRequestStateAndReasons(model)
	return state
}

func point11ValBIssuedClaimStateAndReasons(model Point11ValBIssuedClaimRecord) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "issued_claim_projection_disclaimer_blocked")
	}
	if !point11ValBClaimRefValid(model.ClaimID) {
		reasons = append(reasons, "issued_claim_id_invalid")
	}
	if !point11ValBClaimTypeRefValid(model.ClaimTypeID) {
		reasons = append(reasons, "issued_claim_type_id_invalid")
	}
	if !point11Val0IdentityValueValid(model.ClaimVersion) {
		reasons = append(reasons, "issued_claim_version_invalid")
	}
	if !point11ValBSubjectRefValid(model.SubjectRef) {
		reasons = append(reasons, "issued_subject_ref_invalid")
	}
	if !point11ValBIssuerRefValid(model.IssuerRef) || !point11Val0IdentityValueValid(model.IssuerKind) {
		reasons = append(reasons, "issued_issuer_invalid")
	}
	if !point11Val0IdentityValueValid(model.Audience) {
		reasons = append(reasons, "issued_audience_invalid")
	}
	if !point11Val0ScopeValid(model.Scope) {
		reasons = append(reasons, "issued_scope_invalid")
	}
	if !point11Val0ContainsTrimmed(point11Val0PublicationSurfaces(), model.PublicationSurface) {
		reasons = append(reasons, "issued_publication_surface_invalid")
	}
	if !point11ValAPolicyRefValid(model.PolicyBasisRef) ||
		model.PolicyBasisState != point11ValBPolicyBasisStateActive ||
		!point11Val0IdentityValueValid(model.PolicyVersion) {
		reasons = append(reasons, "issued_policy_basis_invalid")
	}
	if !point11Val0EvidenceRefsValid(model.EvidenceRefs) {
		reasons = append(reasons, "issued_evidence_refs_invalid")
	}
	if model.EvidenceHashRequired && !point11ValBEvidenceHashRefsValid(model.EvidenceHashRefs) {
		reasons = append(reasons, "issued_evidence_hash_refs_invalid")
	}
	if !point11ValBVerificationMethodSupported(model.VerificationMethod) {
		reasons = append(reasons, "issued_verification_method_invalid")
	}
	if model.VerificationResult != point11ValBVerificationResultVerified {
		reasons = append(reasons, "issued_verification_result_not_active")
	}
	if !point11Val0ValidTimestamp(model.IssuedAt) || !point11Val0ValidTimestamp(model.ExpiresAt) {
		reasons = append(reasons, "issued_timestamps_invalid")
	} else {
		expiresAt, _ := time.Parse(time.RFC3339, model.ExpiresAt)
		if expiresAt.Before(time.Now().UTC()) {
			reasons = append(reasons, "issued_claim_expired")
		}
	}
	if !point11ValBClaimLifecycleAllowed(model.LifecycleState) {
		reasons = append(reasons, "issued_lifecycle_invalid")
	}
	switch model.LifecycleState {
	case Point11Val0ClaimLifecycleRevoked:
		if !point11ValBGovernanceEventRefValid(model.RevocationRef) {
			reasons = append(reasons, "issued_revocation_ref_missing")
		}
		reasons = append(reasons, "issued_claim_revoked")
	case Point11Val0ClaimLifecycleSuperseded:
		if !point11ValBClaimRefValid(model.SupersededBy) {
			reasons = append(reasons, "issued_superseded_by_missing")
		}
		reasons = append(reasons, "issued_claim_superseded")
	case Point11Val0ClaimLifecycleCorrected:
		if !point11ValBGovernanceEventRefValid(model.CorrectionRef) {
			reasons = append(reasons, "issued_correction_ref_missing")
		}
		reasons = append(reasons, "issued_claim_corrected")
	case Point11Val0ClaimLifecycleBlocked, point11ValBClaimLifecycleExpired:
		reasons = append(reasons, "issued_claim_not_active_use_eligible")
	default:
		if !point11ValBClaimLifecycleActiveUseEligible(model.LifecycleState) {
			reasons = append(reasons, "issued_claim_lifecycle_not_active_use_eligible")
		}
	}
	if !point11ValBClaimCategoryAllowed(model.ClaimCategory) {
		reasons = append(reasons, "issued_claim_category_invalid")
	}
	if point11ValBClaimNeedsGovernanceEvent(model.ClaimCategory, model.PublicationSurface) &&
		!point11ValBGovernanceEventRefValid(model.GovernanceEventRef) {
		reasons = append(reasons, "issued_governance_event_missing")
	}
	if point11ValBClaimNeedsCleanRoomReview(model.ClaimCategory, model.PublicationSurface) &&
		!point11ValBCleanRoomReviewRefValid(model.CleanRoomIPReviewRef) {
		reasons = append(reasons, "issued_clean_room_review_missing")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "issued_audit_id_invalid")
	}
	if len(reasons) > 0 {
		return Point11ValBIssuedClaimStateBlocked, reasons
	}
	return Point11ValBIssuedClaimStateActive, nil
}

func EvaluatePoint11ValBIssuedClaimState(model Point11ValBIssuedClaimRecord) string {
	state, _ := point11ValBIssuedClaimStateAndReasons(model)
	return state
}

func point11ValBRegisteredClaimMap(claims []Point11ValBIssuedClaimRecord) map[string]Point11ValBIssuedClaimRecord {
	registered := map[string]Point11ValBIssuedClaimRecord{}
	for _, claim := range claims {
		registered[claim.ClaimID] = claim
	}
	return registered
}

func point11ValBClaimIdentityKey(claim Point11ValBIssuedClaimRecord) string {
	return strings.Join([]string{
		claim.ClaimID,
		claim.SubjectRef,
		claim.PolicyBasisRef,
		point11ValBStringSetKey(claim.EvidenceHashRefs),
	}, "|")
}

func point11ValBRegistryStateAndReasons(model Point11ValBClaimRegistry) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "claim_registry_projection_disclaimer_blocked")
	}
	if !point11ValBClaimRegistryRefValid(model.RegistryID) {
		reasons = append(reasons, "claim_registry_id_invalid")
	}
	if !point11Val0IdentityValueValid(model.RegistryVersion) {
		reasons = append(reasons, "claim_registry_version_invalid")
	}
	if !point11ValAPolicyRefValid(model.RegistryPolicyBasisRef) ||
		model.RegistryPolicyBasisState != point11ValBPolicyBasisStateActive {
		reasons = append(reasons, "claim_registry_policy_basis_invalid")
	}
	if !point11ValBGovernanceEventRefValid(model.GovernanceEventRef) {
		reasons = append(reasons, "claim_registry_governance_event_invalid")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "claim_registry_audit_id_invalid")
	}
	claimRefsLists := [][]string{
		model.ActiveClaimRefs,
		model.RevokedClaimRefs,
		model.SupersededClaimRefs,
		model.CorrectedClaimRefs,
		model.BlockedClaimRefs,
	}
	for _, list := range claimRefsLists {
		if len(list) > 0 && !point11ValBClaimRefListValid(list) {
			reasons = append(reasons, "claim_registry_claim_ref_list_invalid")
			break
		}
	}
	if len(model.DuplicateClaimRefs) > 0 {
		if !point11ValBClaimRefListValid(model.DuplicateClaimRefs) {
			reasons = append(reasons, "claim_registry_duplicate_refs_invalid")
		}
		reasons = append(reasons, "claim_registry_duplicate_claim_identity")
	}
	if len(model.ConflictingClaimRefs) > 0 {
		if !point11ValBClaimRefListValid(model.ConflictingClaimRefs) {
			reasons = append(reasons, "claim_registry_conflicting_refs_invalid")
		}
		reasons = append(reasons, "claim_registry_conflicting_claim_identity")
	}
	if len(model.RegisteredClaims) == 0 {
		reasons = append(reasons, "claim_registry_registered_claims_missing")
	}
	registeredClaims := point11ValBRegisteredClaimMap(model.RegisteredClaims)
	activeRefSet := map[string]struct{}{}
	for _, ref := range model.ActiveClaimRefs {
		activeRefSet[ref] = struct{}{}
		claim, exists := registeredClaims[ref]
		if !exists {
			reasons = append(reasons, "claim_registry_active_ref_not_registered")
			continue
		}
		state, _ := point11ValBIssuedClaimStateAndReasons(claim)
		if state != Point11ValBIssuedClaimStateActive {
			reasons = append(reasons, "claim_registry_active_ref_points_to_non_active_claim")
		}
	}
	seenIdentity := map[string]struct{}{}
	claimsByID := map[string][]Point11ValBIssuedClaimRecord{}
	for _, claim := range model.RegisteredClaims {
		identityKey := point11ValBClaimIdentityKey(claim)
		if identityKey == "" {
			reasons = append(reasons, "claim_registry_registered_claim_identity_invalid")
			continue
		}
		if _, exists := seenIdentity[identityKey]; exists {
			reasons = append(reasons, "claim_registry_duplicate_claim_identity")
		}
		seenIdentity[identityKey] = struct{}{}
		claimsByID[claim.ClaimID] = append(claimsByID[claim.ClaimID], claim)
	}
	for _, claims := range claimsByID {
		if len(claims) < 2 {
			continue
		}
		scopeKeys := map[string]struct{}{}
		subjectKeys := map[string]struct{}{}
		policyKeys := map[string]struct{}{}
		evidenceKeys := map[string]struct{}{}
		governed := true
		for _, claim := range claims {
			scopeKeys[claim.Scope] = struct{}{}
			subjectKeys[claim.SubjectRef] = struct{}{}
			policyKeys[claim.PolicyBasisRef] = struct{}{}
			evidenceKeys[point11ValBStringSetKey(claim.EvidenceHashRefs)] = struct{}{}
			lifecycle := claim.LifecycleState
			if lifecycle != Point11Val0ClaimLifecycleCorrected && lifecycle != Point11Val0ClaimLifecycleSuperseded {
				governed = false
			}
			if !point11ValBGovernanceEventRefValid(claim.GovernanceEventRef) {
				governed = false
			}
		}
		if len(scopeKeys) > 1 {
			reasons = append(reasons, "claim_registry_cross_tenant_claim_reuse")
		}
		if len(subjectKeys) > 1 || len(policyKeys) > 1 || len(evidenceKeys) > 1 {
			if !governed {
				reasons = append(reasons, "claim_registry_conflicting_claim_identity")
			}
		}
	}
	if len(reasons) > 0 {
		return Point11ValBRegistryStateBlocked, reasons
	}
	return Point11ValBRegistryStateActive, nil
}

func EvaluatePoint11ValBRegistryState(model Point11ValBClaimRegistry) string {
	state, _ := point11ValBRegistryStateAndReasons(model)
	return state
}

func point11ValBVerificationStateAndReasons(model Point11ValBClaimVerificationResult) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "claim_verification_projection_disclaimer_blocked")
	}
	if !point11ValBVerificationRefValid(model.VerificationID) {
		reasons = append(reasons, "claim_verification_id_invalid")
	}
	if !point11ValBClaimRefValid(model.ClaimRef) || !point11Val0IdentityValueValid(model.ClaimVersion) {
		reasons = append(reasons, "claim_verification_claim_ref_invalid")
	}
	if !point11ValBIssuerRefValid(model.VerifierRef) || !point11Val0IdentityValueValid(model.VerifierKind) {
		reasons = append(reasons, "claim_verification_verifier_invalid")
	}
	if !point11ValAPolicyRefValid(model.PolicyBasisRef) || model.PolicyBasisState != point11ValBPolicyBasisStateActive {
		reasons = append(reasons, "claim_verification_policy_basis_invalid")
	}
	if !model.ClaimRegistered || model.ClaimRegistryState != Point11ValBRegistryStateActive {
		reasons = append(reasons, "claim_verification_claim_not_registered")
	}
	if !point11Val0EvidenceRefsValid(model.EvidenceRefs) || !point11Val0EvidenceRefsValid(model.ClaimEvidenceRefs) ||
		!point11ValBExactStringSetMatch(model.EvidenceRefs, model.ClaimEvidenceRefs) {
		reasons = append(reasons, "claim_verification_evidence_refs_mismatch")
	}
	if !point11ValBOptionalEvidenceHashRefsValid(model.EvidenceHashRequired, model.EvidenceHashRefs, model.ClaimEvidenceHashRefs) {
		reasons = append(reasons, "claim_verification_evidence_hash_refs_mismatch")
	}
	if !point11ValBVerificationMethodSupported(model.VerificationMethod) {
		reasons = append(reasons, "claim_verification_method_invalid")
	}
	if !point11Val0ValidTimestamp(model.VerificationTimestamp) {
		reasons = append(reasons, "claim_verification_timestamp_invalid")
	}
	if model.FreshnessResult != point11ValBFreshnessResultActive {
		reasons = append(reasons, "claim_verification_freshness_blocked")
	}
	if model.SignatureOrRegistryResult != point11ValBSignatureOrRegistryResultActive {
		reasons = append(reasons, "claim_verification_signature_or_registry_blocked")
	}
	if model.RevocationCheckResult != point11ValBRevocationCheckActive {
		reasons = append(reasons, "claim_verification_revocation_blocked")
	}
	if model.SupersessionCheckResult != point11ValBSupersessionCheckActive {
		if !(model.HistoricalVerificationAllowed && model.SupersessionCheckResult == point11ValBSupersessionCheckHistoricalOnly) {
			reasons = append(reasons, "claim_verification_supersession_blocked")
		}
	}
	if model.ScopeCheckResult != point11ValBScopeCheckActive {
		reasons = append(reasons, "claim_verification_scope_blocked")
	}
	if model.AudienceCheckResult != point11ValBAudienceCheckActive {
		reasons = append(reasons, "claim_verification_audience_blocked")
	}
	if model.IssuerTrustResult != point11ValBIssuerTrustActive {
		reasons = append(reasons, "claim_verification_issuer_trust_blocked")
	}
	if model.CrossDomainRequired && model.CrossDomainCompatibilityResult != point11ValBCrossDomainCompatibilityActive {
		reasons = append(reasons, "claim_verification_cross_domain_blocked")
	}
	if model.ResultState != point11ValBVerificationResultVerified {
		reasons = append(reasons, "claim_verification_result_state_blocked")
	}
	if point11ValBClaimLifecycleInvalidated(model.ClaimLifecycleState) && !model.HistoricalVerificationAllowed {
		reasons = append(reasons, "claim_verification_claim_not_active_use_eligible")
	}
	if point11Val0ContainsForbiddenClaim(model.Diagnostics...) {
		reasons = append(reasons, "claim_verification_diagnostics_overclaim")
	}
	if len(reasons) > 0 {
		return Point11ValBVerificationStateBlocked, reasons
	}
	return Point11ValBVerificationStateActive, nil
}

func EvaluatePoint11ValBVerificationState(model Point11ValBClaimVerificationResult) string {
	state, _ := point11ValBVerificationStateAndReasons(model)
	return state
}

func point11ValBCrossDomainStateAndReasons(model Point11ValBCrossDomainClaimIntake) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "cross_domain_projection_disclaimer_blocked")
	}
	if !point11ValBClaimRefValid(model.RemoteClaimRef) {
		reasons = append(reasons, "cross_domain_remote_claim_ref_invalid")
	}
	if !point11Val0IdentityValueValid(model.RemoteClaimType) || !point11Val0ContainsTrimmed(model.AcceptedClaimTypes, model.RemoteClaimType) {
		reasons = append(reasons, "cross_domain_remote_claim_type_not_accepted")
	}
	if !point11ValBIssuerRefValid(model.RemoteIssuerRef) {
		reasons = append(reasons, "cross_domain_remote_issuer_invalid")
	}
	if !point11ValBTrustRootRefValid(model.RemoteTrustRootRef) {
		reasons = append(reasons, "cross_domain_trust_root_invalid")
	}
	if !point11ValAPolicyRefValid(model.LocalPolicyBasisRef) || model.LocalPolicyBasisState != point11ValBPolicyBasisStateActive {
		reasons = append(reasons, "cross_domain_local_policy_basis_invalid")
	}
	if !point11ValBCompatibilityRuleRefValid(model.CompatibilityRuleRef) {
		reasons = append(reasons, "cross_domain_compatibility_rule_invalid")
	}
	if model.RemoteOverridesLocalPolicy {
		reasons = append(reasons, "cross_domain_remote_overrides_local_policy")
	}
	if model.TreatsRemoteClaimAsAuthority {
		reasons = append(reasons, "cross_domain_remote_claim_treated_as_authority")
	}
	if model.CreatesLegalAuthority || model.CreatesRegulatoryAuthority || model.CreatesCertificationAuthority {
		reasons = append(reasons, "cross_domain_creates_forbidden_authority")
	}
	switch model.RemoteClaimState {
	case Point11Val0ClaimLifecycleRevoked, point11ValBClaimLifecycleExpired, Point11Val0ClaimLifecycleSuperseded:
		reasons = append(reasons, "cross_domain_remote_claim_invalidated")
	}
	reviewRequired := false
	switch model.ScopeCompatibilityResult {
	case point11ValBGenericCompatibilityActive:
	case point11ValBGenericCompatibilityReviewRequired:
		reviewRequired = true
	default:
		reasons = append(reasons, "cross_domain_scope_compatibility_blocked")
	}
	switch model.FreshnessCompatibilityResult {
	case point11ValBGenericCompatibilityActive:
	case point11ValBGenericCompatibilityReviewRequired:
		reviewRequired = true
	default:
		reasons = append(reasons, "cross_domain_freshness_compatibility_blocked")
	}
	if model.RevocationHandlingResult != point11ValBRevocationHandlingResultActive {
		reasons = append(reasons, "cross_domain_revocation_handling_blocked")
	}
	if model.IssuerTrustResult != point11ValBIssuerTrustActive {
		reasons = append(reasons, "cross_domain_issuer_trust_blocked")
	}
	switch model.EvidenceTranslationResult {
	case point11ValBEvidenceTranslationResultActive:
	case point11ValBEvidenceTranslationResultReview:
		reviewRequired = true
	default:
		reasons = append(reasons, "cross_domain_evidence_translation_blocked")
	}
	switch model.LocalAdmissibilityResult {
	case point11ValBLocalAdmissibilityResultActive:
	case point11ValBLocalAdmissibilityResultReview:
		reviewRequired = true
	default:
		reasons = append(reasons, "cross_domain_local_admissibility_blocked")
	}
	if point11Val0ContainsForbiddenClaim(model.Diagnostics...) {
		reasons = append(reasons, "cross_domain_diagnostics_overclaim")
	}
	if len(reasons) > 0 {
		return Point11ValBCrossDomainIntakeStateBlocked, reasons
	}
	if reviewRequired {
		return Point11ValBCrossDomainIntakeStateReviewRequired, []string{"cross_domain_review_required"}
	}
	return Point11ValBCrossDomainIntakeStateActive, nil
}

func EvaluatePoint11ValBCrossDomainIntakeState(model Point11ValBCrossDomainClaimIntake) string {
	state, _ := point11ValBCrossDomainStateAndReasons(model)
	return state
}

func point11ValBBlockingReasons(model Point11ValBFoundation) []string {
	reasons := []string{}
	if model.DependencyState != Point11ValBDependencyStateActive &&
		model.DependencyState != Point11ValBDependencyStateReviewRequired {
		reasons = append(reasons, "vala_dependency_blocked")
	}
	if model.ClaimTypeState != Point11ValBClaimTypeStateActive {
		reasons = append(reasons, "claim_type_blocked")
	}
	if model.IssuanceRequestState != Point11ValBIssuanceRequestStateActive {
		reasons = append(reasons, "issuance_request_blocked")
	}
	if model.IssuedClaimState != Point11ValBIssuedClaimStateActive {
		reasons = append(reasons, "issued_claim_blocked")
	}
	if model.RegistryState != Point11ValBRegistryStateActive {
		reasons = append(reasons, "claim_registry_blocked")
	}
	if model.VerificationState != Point11ValBVerificationStateActive {
		reasons = append(reasons, "claim_verification_blocked")
	}
	if model.CrossDomainIntakeState != Point11ValBCrossDomainIntakeStateActive &&
		model.CrossDomainIntakeState != Point11ValBCrossDomainIntakeStateReviewRequired {
		reasons = append(reasons, "cross_domain_intake_blocked")
	}
	if model.CreatesAuthorityClaims {
		reasons = append(reasons, "authority_claim_surface_blocked")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "publication_side_effects_blocked")
	}
	return reasons
}

func point11ValBDiagnosticsModel(
	model Point11ValBFoundation,
	dependencyReasons []string,
	claimTypeReasons []string,
	issuanceReasons []string,
	issuedClaimReasons []string,
	registryReasons []string,
	verificationReasons []string,
	crossDomainReasons []string,
) Point11ValBDiagnostics {
	return Point11ValBDiagnostics{
		CurrentState:         model.CurrentState,
		BlockingReasons:      append([]string{}, model.BlockingReasons...),
		ReviewPrerequisites:  append([]string{}, model.ReviewPrerequisites...),
		ComponentStates:      point11ValBComponentStates(model),
		DependencyReasons:    append([]string{}, dependencyReasons...),
		ClaimTypeReasons:     append([]string{}, claimTypeReasons...),
		IssuanceReasons:      append([]string{}, issuanceReasons...),
		IssuedClaimReasons:   append([]string{}, issuedClaimReasons...),
		RegistryReasons:      append([]string{}, registryReasons...),
		VerificationReasons:  append([]string{}, verificationReasons...),
		CrossDomainReasons:   append([]string{}, crossDomainReasons...),
		ProjectionDisclaimer: model.ProjectionDisclaimer,
	}
}

func EvaluatePoint11ValBFoundationState(model Point11ValBFoundation) string {
	// Any local blocked component overrides review-required dependency or
	// cross-domain states. Review-required may propagate only when no local
	// blocking state is present.
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.CreatesAuthorityClaims ||
		model.CreatesPublicationSideEffects ||
		model.ClaimTypeState == Point11ValBClaimTypeStateBlocked ||
		model.IssuanceRequestState == Point11ValBIssuanceRequestStateBlocked ||
		model.IssuedClaimState == Point11ValBIssuedClaimStateBlocked ||
		model.RegistryState == Point11ValBRegistryStateBlocked ||
		model.VerificationState == Point11ValBVerificationStateBlocked ||
		model.CrossDomainIntakeState == Point11ValBCrossDomainIntakeStateBlocked ||
		model.DependencyState == Point11ValBDependencyStateBlocked {
		return Point11ValBStateBlocked
	}
	if model.DependencyState == Point11ValBDependencyStateReviewRequired ||
		model.CrossDomainIntakeState == Point11ValBCrossDomainIntakeStateReviewRequired {
		return Point11ValBStateReviewRequired
	}
	if model.DependencyState == Point11ValBDependencyStateActive &&
		model.ClaimTypeState == Point11ValBClaimTypeStateActive &&
		model.IssuanceRequestState == Point11ValBIssuanceRequestStateActive &&
		model.IssuedClaimState == Point11ValBIssuedClaimStateActive &&
		model.RegistryState == Point11ValBRegistryStateActive &&
		model.VerificationState == Point11ValBVerificationStateActive &&
		model.CrossDomainIntakeState == Point11ValBCrossDomainIntakeStateActive {
		return Point11ValBStateActive
	}
	return Point11ValBStateBlocked
}

func Point11ValBFoundationModel() Point11ValBFoundation {
	disclaimer := point11ValBProjectionDisclaimerBaseline
	claimTypeID := "claim_type_point11_valb_policy_attested_claim"
	claimID := "claim_point11_valb_customer_scope_001"
	policyRef := "policy_point11_vala_authority_core_v1"
	claimRegistryID := "claim_registry_point11_valb_core"
	return Point11ValBFoundation{
		CurrentState:                  Point11ValBStateReviewRequired,
		DependencyState:               Point11ValBDependencyStateReviewRequired,
		ClaimTypeState:                Point11ValBClaimTypeStateActive,
		IssuanceRequestState:          Point11ValBIssuanceRequestStateActive,
		IssuedClaimState:              Point11ValBIssuedClaimStateActive,
		RegistryState:                 Point11ValBRegistryStateActive,
		VerificationState:             Point11ValBVerificationStateActive,
		CrossDomainIntakeState:        Point11ValBCrossDomainIntakeStateActive,
		ProjectionDisclaimer:          disclaimer,
		CreatesAuthorityClaims:        false,
		CreatesPublicationSideEffects: false,
		Dependency:                    point11ValBDependencySnapshotModel(),
		ClaimTypeDefinition: Point11ValBClaimTypeDefinition{
			ClaimTypeID:                           claimTypeID,
			ClaimType:                             "bounded_governance_claim",
			Category:                              Point11Val0ClaimCategoryCustomerVisible,
			AllowedSubjectKinds:                   []string{"workload", "service", "artifact"},
			AllowedIssuerKinds:                    []string{"governance_authority", "policy_engine"},
			AllowedAudiences:                      []string{"customer", "partner", "internal_review"},
			AllowedPublicationSurfaces:            []string{point11Val0PublicationSurfaceDocs, point11Val0PublicationSurfacePortal, point11Val0PublicationSurfaceExport, point11Val0PublicationSurfacePartner, point11Val0PublicationSurfaceDemo, point11Val0PublicationSurfaceSales, point11Val0PublicationSurfaceBuyer},
			RequiredPolicyBasis:                   policyRef,
			RequiredEvidenceKinds:                 []string{"evidence_bundle", "verification_record"},
			RequiredVerificationMethod:            point11ValBVerificationMethodEvidenceLinked,
			RequiredFreshnessClass:                point11ValBFreshnessClassCurrent,
			DefaultExpiryDuration:                 "duration_p30d",
			RevocationRequired:                    true,
			SupersessionAllowed:                   true,
			CrossDomainAllowed:                    true,
			AgentOriginAllowed:                    true,
			CustomerVisibleAllowed:                true,
			PublicSafeAllowed:                     true,
			CleanRoomIPRequired:                   true,
			GovernanceEventRequired:               true,
			CrossDomainTrustCompatibilityRequired: true,
			AgentPublishAuthority:                 false,
			AgentApproveAuthority:                 false,
			ProjectionDisclaimer:                  disclaimer,
		},
		IssuanceRequest: Point11ValBClaimIssuanceRequest{
			IssuanceRequestID:                "verification_point11_valb_issuance_request_001",
			ClaimID:                          claimID,
			ClaimTypeID:                      claimTypeID,
			ClaimTypeName:                    "bounded_governance_claim",
			SubjectRef:                       "subject_point11_valb_workload_alpha",
			SubjectKind:                      "workload",
			IssuerRef:                        "issuer_point11_valb_governance_team",
			IssuerKind:                       "governance_authority",
			ProposerRef:                      "internal_proposer_team",
			ReviewerRef:                      "internal_reviewer_team",
			ApproverRef:                      "internal_final_approver",
			Audience:                         "customer",
			PublicationSurface:               point11Val0PublicationSurfaceExport,
			PolicyBasisRef:                   policyRef,
			PolicyBasisState:                 point11ValBPolicyBasisStateActive,
			PolicyVersion:                    "point11_vala_policy_v1",
			EvidenceRefs:                     []string{"evidence:point11-valb-issuance-001", "evidence:point11-valb-issuance-002"},
			VerificationMethod:               point11ValBVerificationMethodEvidenceLinked,
			FreshnessClass:                   point11ValBFreshnessClassCurrent,
			RequestedLifecycleState:          Point11Val0ClaimLifecyclePublished,
			RequestedClaimCategory:           Point11Val0ClaimCategoryCustomerVisible,
			CleanRoomIPReviewRef:             "clean_room_review_point11_valb_export_001",
			GovernanceEventRef:               "governance_event_point11_valb_claim_issue_001",
			AuthorityMatrixRef:               "authority_matrix_point11_val0_claim_governance",
			AuthorityMatrixState:             point11ValBAuthorityMatrixStateActive,
			IssuedAt:                         "2026-05-02T12:00:00Z",
			ExpiresAt:                        "2099-01-01T00:00:00Z",
			ClaimTypeState:                   Point11ValBClaimTypeStateActive,
			ClaimTypeAllowedSubjectKinds:     []string{"workload", "service", "artifact"},
			ClaimTypeAllowedIssuerKinds:      []string{"governance_authority", "policy_engine"},
			ClaimTypeAllowedAudiences:        []string{"customer", "partner", "internal_review"},
			ClaimTypeAllowedSurfaces:         []string{point11Val0PublicationSurfaceDocs, point11Val0PublicationSurfacePortal, point11Val0PublicationSurfaceExport, point11Val0PublicationSurfacePartner, point11Val0PublicationSurfaceDemo, point11Val0PublicationSurfaceSales, point11Val0PublicationSurfaceBuyer},
			ClaimTypeRequiredVerification:    point11ValBVerificationMethodEvidenceLinked,
			ClaimTypeRequiredFreshness:       point11ValBFreshnessClassCurrent,
			ClaimTypeGovernanceEventRequired: true,
			ClaimTypeCleanRoomIPRequired:     true,
			ClaimTypeCustomerVisibleAllowed:  true,
			ClaimTypePublicSafeAllowed:       true,
			ClaimTypeAgentOriginAllowed:      true,
			AgentOrigin:                      false,
			ProjectionDisclaimer:             disclaimer,
		},
		IssuedClaim: Point11ValBIssuedClaimRecord{
			ClaimID:              claimID,
			ClaimTypeID:          claimTypeID,
			ClaimVersion:         "claim_version_point11_valb_v1",
			SubjectRef:           "subject_point11_valb_workload_alpha",
			IssuerRef:            "issuer_point11_valb_governance_team",
			IssuerKind:           "governance_authority",
			Audience:             "customer",
			Scope:                "tenant_scope_alpha",
			PublicationSurface:   point11Val0PublicationSurfaceExport,
			PolicyBasisRef:       policyRef,
			PolicyBasisState:     point11ValBPolicyBasisStateActive,
			PolicyVersion:        "point11_vala_policy_v1",
			EvidenceRefs:         []string{"evidence:point11-valb-issuance-001", "evidence:point11-valb-issuance-002"},
			EvidenceHashRefs:     []string{"evidence_hash_point11_valb_claim_001", "evidence_hash_point11_valb_claim_002"},
			EvidenceHashRequired: true,
			VerificationMethod:   point11ValBVerificationMethodEvidenceLinked,
			VerificationResult:   point11ValBVerificationResultVerified,
			IssuedAt:             "2026-05-02T12:00:00Z",
			ExpiresAt:            "2099-01-01T00:00:00Z",
			LifecycleState:       Point11Val0ClaimLifecyclePublished,
			ClaimCategory:        Point11Val0ClaimCategoryCustomerVisible,
			GovernanceEventRef:   "governance_event_point11_valb_claim_issue_001",
			CleanRoomIPReviewRef: "clean_room_review_point11_valb_export_001",
			AuditID:              "audit_point11_valb_claim_001",
			ProjectionDisclaimer: disclaimer,
		},
		Registry: Point11ValBClaimRegistry{
			RegistryID:               claimRegistryID,
			RegistryVersion:          "claim_registry_point11_valb_v1",
			RegisteredClaims:         []Point11ValBIssuedClaimRecord{},
			ActiveClaimRefs:          []string{claimID},
			RegistryPolicyBasisRef:   policyRef,
			RegistryPolicyBasisState: point11ValBPolicyBasisStateActive,
			GovernanceEventRef:       "governance_event_point11_valb_registry_001",
			AuditID:                  "audit_point11_valb_registry_001",
			ProjectionDisclaimer:     disclaimer,
		},
		Verification: Point11ValBClaimVerificationResult{
			VerificationID:                 "verification_point11_valb_claim_001",
			ClaimRef:                       claimID,
			ClaimVersion:                   "claim_version_point11_valb_v1",
			VerifierRef:                    "issuer_point11_valb_verifier_team",
			VerifierKind:                   "verification_service",
			PolicyBasisRef:                 policyRef,
			PolicyBasisState:               point11ValBPolicyBasisStateActive,
			EvidenceRefs:                   []string{"evidence:point11-valb-issuance-001", "evidence:point11-valb-issuance-002"},
			EvidenceHashRefs:               []string{"evidence_hash_point11_valb_claim_001", "evidence_hash_point11_valb_claim_002"},
			EvidenceHashRequired:           true,
			VerificationMethod:             point11ValBVerificationMethodEvidenceLinked,
			VerificationTimestamp:          "2026-05-02T12:05:00Z",
			FreshnessResult:                point11ValBFreshnessResultActive,
			SignatureOrRegistryResult:      point11ValBSignatureOrRegistryResultActive,
			RevocationCheckResult:          point11ValBRevocationCheckActive,
			SupersessionCheckResult:        point11ValBSupersessionCheckActive,
			ScopeCheckResult:               point11ValBScopeCheckActive,
			AudienceCheckResult:            point11ValBAudienceCheckActive,
			IssuerTrustResult:              point11ValBIssuerTrustActive,
			CrossDomainCompatibilityResult: point11ValBCrossDomainCompatibilityActive,
			ResultState:                    point11ValBVerificationResultVerified,
			Diagnostics:                    []string{"claim_verification_bounded"},
			ProjectionDisclaimer:           disclaimer,
			ClaimRegistered:                true,
			ClaimRegistryState:             Point11ValBRegistryStateActive,
			ClaimEvidenceRefs:              []string{"evidence:point11-valb-issuance-001", "evidence:point11-valb-issuance-002"},
			ClaimEvidenceHashRefs:          []string{"evidence_hash_point11_valb_claim_001", "evidence_hash_point11_valb_claim_002"},
			ClaimLifecycleState:            Point11Val0ClaimLifecyclePublished,
			HistoricalVerificationAllowed:  false,
			CrossDomainRequired:            false,
		},
		CrossDomainIntake: Point11ValBCrossDomainClaimIntake{
			RemoteClaimRef:               "claim_point11_valb_remote_001",
			RemoteClaimType:              "bounded_governance_claim",
			RemoteClaimState:             point11ValBClaimLifecycleActive,
			RemoteIssuerRef:              "issuer_point11_valb_remote_partner",
			RemoteTrustRootRef:           "trust_root_point11_valb_partner_001",
			AcceptedClaimTypes:           []string{"bounded_governance_claim", "partner_attested_claim"},
			LocalPolicyBasisRef:          policyRef,
			LocalPolicyBasisState:        point11ValBPolicyBasisStateActive,
			CompatibilityRuleRef:         "compatibility_rule_point11_valb_partner_claims",
			ScopeCompatibilityResult:     point11ValBGenericCompatibilityActive,
			FreshnessCompatibilityResult: point11ValBGenericCompatibilityActive,
			RevocationHandlingResult:     point11ValBRevocationHandlingResultActive,
			IssuerTrustResult:            point11ValBIssuerTrustActive,
			EvidenceTranslationResult:    point11ValBEvidenceTranslationResultActive,
			LocalAdmissibilityResult:     point11ValBLocalAdmissibilityResultActive,
			Diagnostics:                  []string{"cross_domain_claim_intake_bounded"},
			ProjectionDisclaimer:         disclaimer,
		},
	}
}

func ComputePoint11ValBFoundation(model Point11ValBFoundation) Point11ValBFoundation {
	if len(model.Registry.RegisteredClaims) == 0 {
		model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim}
	}
	model.Verification.EvidenceHashRequired = model.IssuedClaim.EvidenceHashRequired
	dependencyState, dependencyReasons := point11ValBDependencyStateAndReasons(model.Dependency)
	claimTypeState, claimTypeReasons := point11ValBClaimTypeStateAndReasons(model.ClaimTypeDefinition)
	issuanceState, issuanceReasons := point11ValBIssuanceRequestStateAndReasons(model.IssuanceRequest)
	issuedClaimState, issuedClaimReasons := point11ValBIssuedClaimStateAndReasons(model.IssuedClaim)
	registryState, registryReasons := point11ValBRegistryStateAndReasons(model.Registry)
	verificationState, verificationReasons := point11ValBVerificationStateAndReasons(model.Verification)
	crossDomainState, crossDomainReasons := point11ValBCrossDomainStateAndReasons(model.CrossDomainIntake)

	model.DependencyState = dependencyState
	model.ClaimTypeState = claimTypeState
	model.IssuanceRequestState = issuanceState
	model.IssuedClaimState = issuedClaimState
	model.RegistryState = registryState
	model.VerificationState = verificationState
	model.CrossDomainIntakeState = crossDomainState
	model.CurrentState = EvaluatePoint11ValBFoundationState(model)
	model.BlockingReasons = point11ValBBlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	if model.CrossDomainIntakeState == Point11ValBCrossDomainIntakeStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "cross_domain_intake_review_required")
	}
	model.Diagnostics = point11ValBDiagnosticsModel(
		model,
		dependencyReasons,
		claimTypeReasons,
		issuanceReasons,
		issuedClaimReasons,
		registryReasons,
		verificationReasons,
		crossDomainReasons,
	)
	return model
}

func point11Val0PublicFacingSurfaceSet(values []string) bool {
	for _, value := range values {
		if point11Val0PublicFacingSurface(value) {
			return true
		}
	}
	return false
}
