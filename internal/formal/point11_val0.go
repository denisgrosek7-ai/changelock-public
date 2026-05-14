package formal

import (
	"strings"
	"sync"
	"time"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point11Val0StateActive         = "point11_val0_governance_foundation_active"
	Point11Val0StateBlocked        = "point11_val0_governance_foundation_blocked"
	Point11Val0StateReviewRequired = "point11_val0_governance_foundation_review_required"

	Point11Val0DependencyStateActive         = "point11_val0_dependency_active"
	Point11Val0DependencyStateBlocked        = "point11_val0_dependency_blocked"
	Point11Val0DependencyStateReviewRequired = "point11_val0_dependency_review_required"

	Point11Val0PolicyContractStateActive  = "point11_val0_policy_contract_active"
	Point11Val0PolicyContractStateBlocked = "point11_val0_policy_contract_blocked"

	Point11Val0ClaimGovernanceStateActive  = "point11_val0_claim_governance_active"
	Point11Val0ClaimGovernanceStateBlocked = "point11_val0_claim_governance_blocked"

	Point11Val0AuthorityMatrixStateActive  = "point11_val0_authority_matrix_active"
	Point11Val0AuthorityMatrixStateBlocked = "point11_val0_authority_matrix_blocked"

	Point11Val0ExceptionGovernanceStateActive  = "point11_val0_exception_governance_active"
	Point11Val0ExceptionGovernanceStateBlocked = "point11_val0_exception_governance_blocked"

	Point11Val0ABACStateActive  = "point11_val0_abac_governance_active"
	Point11Val0ABACStateBlocked = "point11_val0_abac_governance_blocked"

	Point11Val0DecisionBindingStateActive  = "point11_val0_decision_binding_active"
	Point11Val0DecisionBindingStateBlocked = "point11_val0_decision_binding_blocked"

	Point11Val0NoOverclaimStateActive  = "point11_val0_no_overclaim_active"
	Point11Val0NoOverclaimStateBlocked = "point11_val0_no_overclaim_blocked"

	Point11Val0CrossDomainCompatibilityStateActive         = "point11_val0_cross_domain_compatibility_active"
	Point11Val0CrossDomainCompatibilityStateBlocked        = "point11_val0_cross_domain_compatibility_blocked"
	Point11Val0CrossDomainCompatibilityStateReviewRequired = "point11_val0_cross_domain_compatibility_review_required"
)

const (
	Point11Val0ClaimLifecycleDraft          = "draft"
	Point11Val0ClaimLifecycleReviewRequired = "review_required"
	Point11Val0ClaimLifecycleApproved       = "approved"
	Point11Val0ClaimLifecyclePublished      = "published"
	Point11Val0ClaimLifecycleRevoked        = "revoked"
	Point11Val0ClaimLifecycleSuperseded     = "superseded"
	Point11Val0ClaimLifecycleCorrected      = "corrected"
	Point11Val0ClaimLifecycleBlocked        = "blocked"

	Point11Val0ClaimCategoryAllowed         = "allowed"
	Point11Val0ClaimCategoryBlocked         = "blocked"
	Point11Val0ClaimCategoryReviewRequired  = "review_required"
	Point11Val0ClaimCategoryInternalOnly    = "internal_only"
	Point11Val0ClaimCategoryCustomerVisible = "customer_visible"
	Point11Val0ClaimCategoryPublicSafe      = "public_safe"
)

const (
	point11Val0PublicationSurfaceDocs           = "docs"
	point11Val0PublicationSurfacePortal         = "portal"
	point11Val0PublicationSurfaceExport         = "export"
	point11Val0PublicationSurfacePartner        = "partner_material"
	point11Val0PublicationSurfaceDemo           = "demo_material"
	point11Val0PublicationSurfaceSales          = "sales_material"
	point11Val0PublicationSurfaceBuyer          = "buyer_material"
	point11Val0PublicationSurfaceAgentOutput    = "agent_output"
	point11Val0ProjectionDisclaimerBaseline     = "projection_only not_canonical_truth point11_val0_governance_foundation"
	point11Val0PolicySignedState                = "signed_policy_contract"
	point11Val0PolicyAnchoredState              = "anchored_policy_contract"
	point11Val0ExceptionSignedGovernanceState   = "signed_exception_governance_context"
	point11Val0DecisionRefStateActive           = "ref_active"
	point11Val0DecisionRefStateRevoked          = "ref_revoked"
	point11Val0DecisionRefStateExpired          = "ref_expired"
	point11Val0DecisionRefStateSuperseded       = "ref_superseded"
	point11Val0DecisionRefStateUnknown          = "ref_unknown"
	point11Val0CrossDomainScopeCompatible       = "scope_compatible"
	point11Val0CrossDomainScopeIncompatible     = "scope_incompatible"
	point11Val0CrossDomainFreshnessCompatible   = "freshness_compatible"
	point11Val0CrossDomainFreshnessIncompatible = "freshness_incompatible"
	point11Val0CrossDomainIssuerTrusted         = "issuer_trusted_by_local_rule"
	point11Val0CrossDomainIssuerUnknown         = "issuer_trust_rule_unknown"
)

var (
	point11Val0ExpectedDependencyProjectionDisclaimerOnce   sync.Once
	point11Val0ExpectedDependencyProjectionDisclaimerCached string
)

type Point11Val0Point10RepoReview struct {
	LatestValEClosurePatchPresent bool     `json:"latest_vale_closure_patch_present"`
	Point10PassOutsideValE        bool     `json:"point10_pass_outside_vale"`
	CIGreenVisible                bool     `json:"ci_green_visible"`
	CIGreen                       bool     `json:"ci_green"`
	MergeStatusVisible            bool     `json:"merge_status_visible"`
	MergeAccepted                 bool     `json:"merge_accepted"`
	ReviewPrerequisites           []string `json:"review_prerequisites,omitempty"`
}

type Point11Val0DependencySnapshot struct {
	Point10CurrentState             string   `json:"point10_current_state"`
	Point10DependencyState          string   `json:"point10_dependency_state"`
	Point10IntegratedInvariantState string   `json:"point10_integrated_invariant_state"`
	Point10EvidenceQualityState     string   `json:"point10_evidence_quality_state"`
	Point10CLBClosureState          string   `json:"point10_clb_closure_state"`
	Point10PassClosureManifestState string   `json:"point10_pass_closure_manifest_state"`
	Point10NoOverclaimState         string   `json:"point10_no_overclaim_state"`
	Point10ProjectionBoundaryState  string   `json:"point10_projection_boundary_state"`
	Point10CleanRoomIPState         string   `json:"point10_clean_room_ip_state"`
	Point10PassRuleState            string   `json:"point10_pass_rule_state"`
	Point10State                    string   `json:"point10_state"`
	Point10CLB0OpenFindings         int      `json:"point10_clb0_open_findings"`
	Point10CLB1OpenFindings         int      `json:"point10_clb1_open_findings"`
	Point10CLB2OpenFindings         int      `json:"point10_clb2_open_findings"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
	LatestValEClosurePatchPresent   bool     `json:"latest_vale_closure_patch_present"`
	Point10PassOutsideValE          bool     `json:"point10_pass_outside_vale"`
	CIGreenVisible                  bool     `json:"ci_green_visible"`
	CIGreen                         bool     `json:"ci_green"`
	MergeStatusVisible              bool     `json:"merge_status_visible"`
	MergeAccepted                   bool     `json:"merge_accepted"`
	ReviewPrerequisites             []string `json:"review_prerequisites,omitempty"`
}

type Point11Val0PolicyContract struct {
	CurrentState           string   `json:"current_state"`
	PolicyID               string   `json:"policy_id"`
	Version                string   `json:"version"`
	Scope                  string   `json:"scope"`
	Issuer                 string   `json:"issuer"`
	Owner                  string   `json:"owner"`
	ApproverChain          []string `json:"approver_chain,omitempty"`
	SignedState            string   `json:"signed_state"`
	AnchoredState          string   `json:"anchored_state"`
	EffectiveFrom          string   `json:"effective_from"`
	EffectiveUntil         string   `json:"effective_until"`
	SchemaVersion          string   `json:"schema_version"`
	CompatibilityVersion   string   `json:"compatibility_version"`
	SupersededBy           string   `json:"superseded_by"`
	RevokedBy              string   `json:"revoked_by"`
	ApprovalEvidenceRefs   []string `json:"approval_evidence_refs,omitempty"`
	MethodologyOrRuleBasis string   `json:"methodology_or_rule_basis"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
}

type Point11Val0ClaimGovernance struct {
	CurrentState                   string   `json:"current_state"`
	LifecycleState                 string   `json:"lifecycle_state"`
	ClaimID                        string   `json:"claim_id"`
	ClaimType                      string   `json:"claim_type"`
	Subject                        string   `json:"subject"`
	Issuer                         string   `json:"issuer"`
	Owner                          string   `json:"owner"`
	Audience                       string   `json:"audience"`
	Scope                          string   `json:"scope"`
	PolicyBasis                    string   `json:"policy_basis"`
	PolicyVersion                  string   `json:"policy_version"`
	EvidenceRefs                   []string `json:"evidence_refs,omitempty"`
	Freshness                      string   `json:"freshness"`
	Expiry                         string   `json:"expiry"`
	VerificationMethod             string   `json:"verification_method"`
	PublicationBoundary            string   `json:"publication_boundary"`
	ClaimCategory                  string   `json:"claim_category"`
	CleanRoomIPReview              string   `json:"clean_room_ip_review"`
	RevocationOrSupersessionStatus string   `json:"revocation_or_supersession_status"`
	RevocationPath                 string   `json:"revocation_path"`
	ApprovalStatus                 string   `json:"approval_status"`
	GovernanceEvent                string   `json:"governance_event"`
	ProjectionDisclaimer           string   `json:"projection_disclaimer"`
}

type Point11Val0AuthorityMatrix struct {
	CurrentState                string `json:"current_state"`
	ClaimType                   string `json:"claim_type"`
	Proposer                    string `json:"proposer"`
	Reviewer                    string `json:"reviewer"`
	FinalApprover               string `json:"final_approver"`
	Publisher                   string `json:"publisher"`
	Revoker                     string `json:"revoker"`
	Corrector                   string `json:"corrector"`
	CustomerVisibleOrPublic     bool   `json:"customer_visible_or_public"`
	ExportVisible               bool   `json:"export_visible"`
	PolicyRelaxationRequested   bool   `json:"policy_relaxation_requested"`
	ExceptionRequested          bool   `json:"exception_requested"`
	AuthorityExpansionRequested bool   `json:"authority_expansion_requested"`
	GovernanceEvent             string `json:"governance_event"`
	ProjectionDisclaimer        string `json:"projection_disclaimer"`
}

type Point11Val0ExceptionGovernance struct {
	CurrentState              string   `json:"current_state"`
	ExceptionID               string   `json:"exception_id"`
	EmergencyClaimID          string   `json:"emergency_claim_id"`
	Reason                    string   `json:"reason"`
	Issuer                    string   `json:"issuer"`
	Approver                  string   `json:"approver"`
	Subject                   string   `json:"subject"`
	Scope                     string   `json:"scope"`
	IssuedAt                  string   `json:"issued_at"`
	ExpiresAt                 string   `json:"expires_at"`
	MonitoringRequirements    string   `json:"monitoring_requirements"`
	RollbackOrReviewCondition string   `json:"rollback_or_review_condition"`
	RevocationPath            string   `json:"revocation_path"`
	AuditID                   string   `json:"audit_id"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	SignedGovernanceState     string   `json:"signed_governance_state"`
	PermanentSilentException  bool     `json:"permanent_silent_exception"`
	Revoked                   bool     `json:"revoked"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
}

type Point11Val0ABACGovernance struct {
	CurrentState          string   `json:"current_state"`
	Subject               string   `json:"subject"`
	Role                  string   `json:"role"`
	Tenant                string   `json:"tenant"`
	Environment           string   `json:"environment"`
	Artifact              string   `json:"artifact"`
	PolicyProfile         string   `json:"policy_profile"`
	ClaimRefs             []string `json:"claim_refs,omitempty"`
	PolicyRefs            []string `json:"policy_refs,omitempty"`
	ExceptionState        string   `json:"exception_state"`
	DecisionSurface       string   `json:"decision_surface"`
	AllowedAttributes     []string `json:"allowed_attributes,omitempty"`
	DeniedAttributes      []string `json:"denied_attributes,omitempty"`
	UnknownAttributes     []string `json:"unknown_attributes,omitempty"`
	ExplanationAttributes []string `json:"explanation_attributes,omitempty"`
	ExplanationPolicyRefs []string `json:"explanation_policy_refs,omitempty"`
	ExplanationClaimRefs  []string `json:"explanation_claim_refs,omitempty"`
	ExceptionInteraction  string   `json:"exception_interaction"`
	Diagnostics           []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer  string   `json:"projection_disclaimer"`
}

type Point11Val0DecisionBinding struct {
	CurrentState               string   `json:"current_state"`
	DecisionID                 string   `json:"decision_id"`
	PolicyRef                  string   `json:"policy_ref"`
	PolicyRefState             string   `json:"policy_ref_state"`
	ClaimRef                   string   `json:"claim_ref"`
	ClaimRefApplicable         bool     `json:"claim_ref_applicable"`
	ClaimRefState              string   `json:"claim_ref_state"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	DecisionTimestamp          string   `json:"decision_timestamp"`
	DecisionOwnerOrSystemActor string   `json:"decision_owner_or_system_actor"`
	EnforcementOutcome         string   `json:"enforcement_outcome"`
	Diagnostics                []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type Point11Val0NoOverclaimDiscipline struct {
	CurrentState         string   `json:"current_state"`
	ObservedClaims       []string `json:"observed_claims,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type Point11Val0CrossDomainCompatibility struct {
	CurrentState                  string   `json:"current_state"`
	TrustRootRef                  string   `json:"trust_root_ref"`
	IssuerTrustRule               string   `json:"issuer_trust_rule"`
	AcceptedClaimTypes            []string `json:"accepted_claim_types,omitempty"`
	ScopeCompatibility            string   `json:"scope_compatibility"`
	FreshnessCompatibility        string   `json:"freshness_compatibility"`
	RevocationHandling            string   `json:"revocation_handling"`
	IncompatibilityBehavior       string   `json:"incompatibility_behavior"`
	LocalPolicyAdmissibilityRule  string   `json:"local_policy_admissibility_rule"`
	RemoteClaimState              string   `json:"remote_claim_state"`
	RemoteOverridesLocalPolicy    bool     `json:"remote_overrides_local_policy"`
	CreatesPublicAuthority        bool     `json:"creates_public_authority"`
	CreatesRegulatoryAuthority    bool     `json:"creates_regulatory_authority"`
	CreatesCertificationAuthority bool     `json:"creates_certification_authority"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type Point11Val0Foundation struct {
	CurrentState                  string                              `json:"current_state"`
	BlockingReasons               []string                            `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites           []string                            `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer          string                              `json:"projection_disclaimer"`
	DependencyState               string                              `json:"dependency_state"`
	PolicyContractState           string                              `json:"policy_contract_state"`
	ClaimGovernanceState          string                              `json:"claim_governance_state"`
	AuthorityMatrixState          string                              `json:"authority_matrix_state"`
	ExceptionGovernanceState      string                              `json:"exception_governance_state"`
	ABACGovernanceState           string                              `json:"abac_governance_state"`
	DecisionBindingState          string                              `json:"decision_binding_state"`
	NoOverclaimState              string                              `json:"no_overclaim_state"`
	CrossDomainCompatibilityState string                              `json:"cross_domain_compatibility_state"`
	Dependency                    Point11Val0DependencySnapshot       `json:"dependency"`
	PolicyContract                Point11Val0PolicyContract           `json:"policy_contract"`
	ClaimGovernance               Point11Val0ClaimGovernance          `json:"claim_governance"`
	AuthorityMatrix               Point11Val0AuthorityMatrix          `json:"authority_matrix"`
	ExceptionGovernance           Point11Val0ExceptionGovernance      `json:"exception_governance"`
	ABACGovernance                Point11Val0ABACGovernance           `json:"abac_governance"`
	DecisionBinding               Point11Val0DecisionBinding          `json:"decision_binding"`
	NoOverclaim                   Point11Val0NoOverclaimDiscipline    `json:"no_overclaim"`
	CrossDomainCompatibility      Point11Val0CrossDomainCompatibility `json:"cross_domain_compatibility"`
}

func point11Val0NormalizeText(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(strings.TrimSpace(value))), " ")
}

func point11Val0ContainsTrimmed(values []string, target string) bool {
	target = strings.TrimSpace(target)
	for _, value := range values {
		if strings.TrimSpace(value) == target {
			return true
		}
	}
	return false
}

func point11Val0ValidTimestamp(value string) bool {
	_, err := time.Parse(time.RFC3339, strings.TrimSpace(value))
	return err == nil
}

func point11Val0ValidProjectionDisclaimer(value string) bool {
	if value == "" || value != strings.TrimSpace(value) || strings.ContainsAny(value, "\t\r\n") {
		return false
	}
	if strings.ToLower(value) != value {
		return false
	}
	tokens := strings.Fields(value)
	if len(tokens) != 3 || tokens[0] != "projection_only" || tokens[1] != "not_canonical_truth" {
		return false
	}
	if !point11Val0ProjectionDisclaimerSuffixValid(tokens[2]) {
		return false
	}
	// Projection disclaimers must stay advisory-only and must not carry forbidden authority claims.
	return !point11Val0ContainsForbiddenClaim(tokens[0] + " " + tokens[1] + " " + strings.ReplaceAll(tokens[2], "_", " "))
}

func point11Val0ProjectionDisclaimerSuffixValid(value string) bool {
	if value == "" || strings.HasPrefix(value, "_") || strings.HasSuffix(value, "_") || strings.Contains(value, "__") {
		return false
	}
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			continue
		}
		return false
	}
	return true
}

func point11Val0HasFoundationProjectionDisclaimer(value string) bool {
	return value == point11Val0ProjectionDisclaimerBaseline
}

func point11Val0ExpectedDependencyProjectionDisclaimer() string {
	point11Val0ExpectedDependencyProjectionDisclaimerOnce.Do(func() {
		// Deterministic baseline: cache once to avoid rebuilding full Point10 models on every dependency check.
		point11Val0ExpectedDependencyProjectionDisclaimerCached =
			operability.DeploymentMultiTenantValEFoundationModel().Point10PassRule.ProjectionDisclaimer
	})
	return point11Val0ExpectedDependencyProjectionDisclaimerCached
}

func point11Val0ValidDependencyProjectionDisclaimer(value string) bool {
	return value == point11Val0ExpectedDependencyProjectionDisclaimer()
}

func point11Val0IdentityValueValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return false
	}
	normalized := point11Val0NormalizeText(trimmed)
	for _, blocked := range []string{
		"unknown",
		"partial",
		"incomplete",
		"stale",
		"malformed",
		"unsupported",
		"blocked",
		"revoked",
		"expired",
		"duplicate",
		"unrelated",
	} {
		if strings.Contains(normalized, blocked) {
			return false
		}
	}
	return true
}

func point11Val0CanonicalRefWithPrefixes(value string, prefixes []string) bool {
	trimmed := strings.TrimSpace(value)
	if !point11Val0IdentityValueValid(trimmed) {
		return false
	}
	if strings.Contains(trimmed, "/") || strings.Contains(trimmed, " ") {
		return false
	}
	lowerTrimmed := strings.ToLower(trimmed)
	for _, blocked := range []string{
		"unknown",
		"unsupported",
		"invalid",
		"revoked",
		"expired",
		"superseded",
		"malformed",
		"placeholder",
		"<empty>",
		"junk",
		"marker",
		"global",
		"unscoped",
		"all-tenants",
		"wildcard",
	} {
		if strings.Contains(lowerTrimmed, blocked) {
			return false
		}
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(trimmed, prefix) {
			return true
		}
	}
	return false
}

func point11Val0PolicyLineageRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"policy_",
		"pol_",
		"point11_policy_",
	})
}

func point11Val0EmergencyClaimRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"emergency_claim_",
		"eclaim_",
		"point11_emergency_claim_",
	})
}

func point11Val0ClaimInvalidatedStatus(value string) bool {
	normalized := point11Val0NormalizeText(value)
	return strings.Contains(normalized, "revoked") ||
		strings.Contains(normalized, "superseded") ||
		strings.Contains(normalized, "expired")
}

func point11Val0ScopeValid(value string) bool {
	if !point11Val0IdentityValueValid(value) {
		return false
	}
	normalized := point11Val0NormalizeText(value)
	for _, blocked := range []string{"global", "unscoped", "wildcard", "all-tenants", "cross-tenant"} {
		if strings.Contains(normalized, blocked) {
			return false
		}
	}
	return true
}

func point11Val0EvidenceRefsValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if !point11Val0IdentityValueValid(trimmed) {
			return false
		}
		if _, exists := seen[trimmed]; exists {
			return false
		}
		seen[trimmed] = struct{}{}
	}
	return true
}

func point11Val0AllValuesValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if !point11Val0IdentityValueValid(trimmed) {
			return false
		}
		if _, exists := seen[trimmed]; exists {
			return false
		}
		seen[trimmed] = struct{}{}
	}
	return true
}

func point11Val0AllowedClaimLifecycles() []string {
	return []string{
		Point11Val0ClaimLifecycleDraft,
		Point11Val0ClaimLifecycleReviewRequired,
		Point11Val0ClaimLifecycleApproved,
		Point11Val0ClaimLifecyclePublished,
		Point11Val0ClaimLifecycleRevoked,
		Point11Val0ClaimLifecycleSuperseded,
		Point11Val0ClaimLifecycleCorrected,
		Point11Val0ClaimLifecycleBlocked,
	}
}

func point11Val0AllowedClaimCategories() []string {
	return []string{
		Point11Val0ClaimCategoryAllowed,
		Point11Val0ClaimCategoryBlocked,
		Point11Val0ClaimCategoryReviewRequired,
		Point11Val0ClaimCategoryInternalOnly,
		Point11Val0ClaimCategoryCustomerVisible,
		Point11Val0ClaimCategoryPublicSafe,
	}
}

func point11Val0PublicationSurfaces() []string {
	return []string{
		point11Val0PublicationSurfaceDocs,
		point11Val0PublicationSurfacePortal,
		point11Val0PublicationSurfaceExport,
		point11Val0PublicationSurfacePartner,
		point11Val0PublicationSurfaceDemo,
		point11Val0PublicationSurfaceSales,
		point11Val0PublicationSurfaceBuyer,
		point11Val0PublicationSurfaceAgentOutput,
	}
}

func point11Val0PublicFacingSurface(surface string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11Val0PublicationSurfaceDocs,
		point11Val0PublicationSurfacePortal,
		point11Val0PublicationSurfaceExport,
		point11Val0PublicationSurfacePartner,
		point11Val0PublicationSurfaceDemo,
		point11Val0PublicationSurfaceSales,
		point11Val0PublicationSurfaceBuyer,
	}, surface)
}

func point11Val0ActorClass(value string) string {
	normalized := point11Val0NormalizeText(value)
	switch {
	case strings.Contains(normalized, "partner"):
		return "partner"
	case strings.Contains(normalized, "customer"):
		return "customer"
	case strings.Contains(normalized, "agent"):
		return "agent"
	default:
		return "internal"
	}
}

func point11Val0ForbiddenClaims() []string {
	return []string{
		"certified",
		"regulator-approved",
		"compliance guaranteed",
		"production approved",
		"deployment approved",
		"legally certified",
		"legal authority",
		"regulator authority",
		"public badge",
		"official authority",
		"universal authority",
		"global truth",
		"mathematically proves total truth",
		"impossible to violate without detection",
		"supreme authority",
		"supreme arbiter",
		"self-healed and secure",
		"agent approved",
		"ai certified",
		"ai approved",
		"ai-approved",
		"ai decision",
		"ai legal proof",
		"autonomous remediation",
		"continuous compliance attestation",
		"compliance-as-a-service",
		"guaranteed secure",
		"one-click secure",
		"zero-risk",
	}
}

func point11Val0AllowedClaims() []string {
	return []string{
		"signed and versioned policy contract",
		"evidence-linked governance decision",
		"bounded claim",
		"claim pending review",
		"approved for declared scope",
		"audit-ready governance trail",
		"policy-bound decision support",
		"not a certification",
		"not regulator approval",
		"not production approval",
		"not deployment approval",
		"not compliance guarantee",
		"advisory governance signal",
		"evidence support available for review",
	}
}

func point11Val0ContainsForbiddenClaim(values ...string) bool {
	allowed := map[string]struct{}{}
	for _, value := range point11Val0AllowedClaims() {
		allowed[point11Val0NormalizeText(value)] = struct{}{}
	}
	for _, value := range values {
		normalized := point11Val0NormalizeText(value)
		if _, ok := allowed[normalized]; ok {
			continue
		}
		for _, forbidden := range point11Val0ForbiddenClaims() {
			if strings.Contains(normalized, point11Val0NormalizeText(forbidden)) {
				return true
			}
		}
	}
	return false
}

// Dependency snapshots must copy actual computed upstream output.
// They must not repair, replace, fallback, or regenerate upstream dependency values.
// The dependency evaluator is responsible for fail-closed validation.
func SnapshotPoint11Val0DependencyFromComputedPoint10ValE(valE operability.DeploymentMultiTenantValEFoundation, repoReview Point11Val0Point10RepoReview) Point11Val0DependencySnapshot {
	return Point11Val0DependencySnapshot{
		Point10CurrentState:             valE.CurrentState,
		Point10DependencyState:          valE.DependencyState,
		Point10IntegratedInvariantState: valE.IntegratedInvariantState,
		Point10EvidenceQualityState:     valE.EvidenceQualityState,
		Point10CLBClosureState:          valE.CLBClosureState,
		Point10PassClosureManifestState: valE.PassClosureManifestState,
		Point10NoOverclaimState:         valE.NoOverclaimState,
		Point10ProjectionBoundaryState:  valE.ProjectionBoundaryState,
		Point10CleanRoomIPState:         valE.CleanRoomIPState,
		Point10PassRuleState:            valE.Point10PassRuleState,
		Point10State:                    valE.Point10State,
		Point10CLB0OpenFindings:         len(valE.CLBClosureLedger.CLB0OpenFindings),
		Point10CLB1OpenFindings:         len(valE.CLBClosureLedger.CLB1OpenFindings),
		Point10CLB2OpenFindings:         len(valE.CLBClosureLedger.CLB2OpenFindings),
		ProjectionDisclaimer:            valE.Point10PassRule.ProjectionDisclaimer,
		LatestValEClosurePatchPresent:   repoReview.LatestValEClosurePatchPresent,
		Point10PassOutsideValE:          repoReview.Point10PassOutsideValE,
		CIGreenVisible:                  repoReview.CIGreenVisible,
		CIGreen:                         repoReview.CIGreen,
		MergeStatusVisible:              repoReview.MergeStatusVisible,
		MergeAccepted:                   repoReview.MergeAccepted,
		ReviewPrerequisites:             append([]string{}, repoReview.ReviewPrerequisites...),
	}
}

func point11Val0DependencyRepoReviewModel() Point11Val0Point10RepoReview {
	return Point11Val0Point10RepoReview{
		LatestValEClosurePatchPresent: true,
		Point10PassOutsideValE:        false,
		CIGreenVisible:                false,
		CIGreen:                       false,
		MergeStatusVisible:            false,
		MergeAccepted:                 false,
		ReviewPrerequisites: []string{
			"point10_ci_green_not_visible_in_repo_context",
			"point10_merge_state_not_visible_in_repo_context",
		},
	}
}

func point11Val0DependencySnapshotModel() Point11Val0DependencySnapshot {
	valE := operability.ComputeDeploymentMultiTenantValEFoundation(operability.DeploymentMultiTenantValEFoundationModel())
	return SnapshotPoint11Val0DependencyFromComputedPoint10ValE(valE, point11Val0DependencyRepoReviewModel())
}

func EvaluatePoint11Val0DependencyState(model Point11Val0DependencySnapshot) string {
	if !point11Val0ValidDependencyProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!model.LatestValEClosurePatchPresent ||
		model.Point10PassOutsideValE ||
		model.Point10CLB0OpenFindings > 0 ||
		model.Point10CLB1OpenFindings > 0 ||
		model.Point10CLB2OpenFindings > 0 ||
		model.Point10CurrentState != operability.DeploymentMultiTenantValEStatePass ||
		model.Point10DependencyState != operability.DeploymentMultiTenantValEDependencyStateActive ||
		model.Point10IntegratedInvariantState != operability.DeploymentMultiTenantValEIntegratedInvariantStateActive ||
		model.Point10EvidenceQualityState != operability.DeploymentMultiTenantValEEvidenceQualityStateActive ||
		model.Point10CLBClosureState != operability.DeploymentMultiTenantValECLBClosureStateActive ||
		model.Point10PassClosureManifestState != operability.DeploymentMultiTenantValEPassClosureManifestStateActive ||
		model.Point10NoOverclaimState != operability.DeploymentMultiTenantValENoOverclaimStateActive ||
		model.Point10ProjectionBoundaryState != operability.DeploymentMultiTenantValEProjectionBoundaryStateActive ||
		model.Point10CleanRoomIPState != operability.DeploymentMultiTenantValECleanRoomIPStateActive ||
		model.Point10PassRuleState != operability.DeploymentMultiTenantValEPoint10PassRuleStateActive ||
		model.Point10State != operability.DeploymentMultiTenantPoint10StatePass {
		return Point11Val0DependencyStateBlocked
	}
	if !model.CIGreenVisible || !model.MergeStatusVisible || len(model.ReviewPrerequisites) > 0 {
		return Point11Val0DependencyStateReviewRequired
	}
	if !model.CIGreen || !model.MergeAccepted {
		return Point11Val0DependencyStateBlocked
	}
	return Point11Val0DependencyStateActive
}

func EvaluatePoint11Val0PolicyContractState(model Point11Val0PolicyContract) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0IdentityValueValid(model.PolicyID) ||
		!point11Val0IdentityValueValid(model.Version) ||
		!point11Val0ScopeValid(model.Scope) ||
		!point11Val0IdentityValueValid(model.Issuer) ||
		!point11Val0IdentityValueValid(model.Owner) ||
		!point11Val0AllValuesValid(model.ApproverChain) ||
		strings.TrimSpace(model.SignedState) != point11Val0PolicySignedState ||
		strings.TrimSpace(model.AnchoredState) != point11Val0PolicyAnchoredState ||
		!point11Val0ValidTimestamp(model.EffectiveFrom) ||
		!point11Val0IdentityValueValid(model.SchemaVersion) ||
		!point11Val0EvidenceRefsValid(model.ApprovalEvidenceRefs) ||
		!point11Val0IdentityValueValid(model.MethodologyOrRuleBasis) {
		return Point11Val0PolicyContractStateBlocked
	}
	if strings.TrimSpace(model.RevokedBy) != "" {
		return Point11Val0PolicyContractStateBlocked
	}
	if strings.TrimSpace(model.EffectiveUntil) != "" {
		if !point11Val0ValidTimestamp(model.EffectiveUntil) {
			return Point11Val0PolicyContractStateBlocked
		}
		expiresAt, _ := time.Parse(time.RFC3339, strings.TrimSpace(model.EffectiveUntil))
		if expiresAt.Before(time.Now().UTC()) {
			return Point11Val0PolicyContractStateBlocked
		}
	}
	if model.SupersededBy != "" {
		if !point11Val0PolicyLineageRefValid(model.SupersededBy) ||
			!point11Val0IdentityValueValid(model.CompatibilityVersion) {
			return Point11Val0PolicyContractStateBlocked
		}
	}
	return Point11Val0PolicyContractStateActive
}

func EvaluatePoint11Val0ClaimGovernanceState(model Point11Val0ClaimGovernance) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0ContainsTrimmed(point11Val0AllowedClaimLifecycles(), model.LifecycleState) ||
		!point11Val0IdentityValueValid(model.ClaimID) ||
		!point11Val0IdentityValueValid(model.ClaimType) ||
		!point11Val0IdentityValueValid(model.Subject) ||
		!point11Val0IdentityValueValid(model.Issuer) ||
		!point11Val0IdentityValueValid(model.Owner) ||
		!point11Val0IdentityValueValid(model.Audience) ||
		!point11Val0ScopeValid(model.Scope) ||
		!point11Val0IdentityValueValid(model.PolicyBasis) ||
		!point11Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point11Val0IdentityValueValid(model.Freshness) ||
		!point11Val0ValidTimestamp(model.Expiry) ||
		!point11Val0IdentityValueValid(model.VerificationMethod) ||
		!point11Val0ContainsTrimmed(point11Val0PublicationSurfaces(), model.PublicationBoundary) ||
		!point11Val0ContainsTrimmed(point11Val0AllowedClaimCategories(), model.ClaimCategory) ||
		!point11Val0IdentityValueValid(model.RevocationOrSupersessionStatus) ||
		point11Val0ContainsForbiddenClaim(model.ClaimType, model.Audience) {
		return Point11Val0ClaimGovernanceStateBlocked
	}
	if strings.TrimSpace(model.LifecycleState) == Point11Val0ClaimLifecycleBlocked ||
		strings.TrimSpace(model.LifecycleState) == Point11Val0ClaimLifecycleRevoked ||
		strings.TrimSpace(model.LifecycleState) == Point11Val0ClaimLifecycleSuperseded ||
		strings.TrimSpace(model.ClaimCategory) == Point11Val0ClaimCategoryBlocked {
		return Point11Val0ClaimGovernanceStateBlocked
	}
	if point11Val0ClaimInvalidatedStatus(model.RevocationOrSupersessionStatus) {
		return Point11Val0ClaimGovernanceStateBlocked
	}
	expiresAt, _ := time.Parse(time.RFC3339, strings.TrimSpace(model.Expiry))
	if expiresAt.Before(time.Now().UTC()) {
		return Point11Val0ClaimGovernanceStateBlocked
	}
	if model.LifecycleState == Point11Val0ClaimLifecyclePublished {
		if !point11Val0IdentityValueValid(model.PolicyVersion) ||
			!point11Val0ScopeValid(model.Scope) ||
			!point11Val0IdentityValueValid(model.Owner) ||
			!point11Val0IdentityValueValid(model.RevocationPath) ||
			strings.TrimSpace(model.ApprovalStatus) != Point11Val0ClaimLifecycleApproved {
			return Point11Val0ClaimGovernanceStateBlocked
		}
	}
	if point11Val0PublicFacingSurface(model.PublicationBoundary) {
		if model.ClaimCategory == Point11Val0ClaimCategoryInternalOnly {
			return Point11Val0ClaimGovernanceStateBlocked
		}
		if model.ClaimCategory == Point11Val0ClaimCategoryReviewRequired && strings.TrimSpace(model.ApprovalStatus) != Point11Val0ClaimLifecycleApproved {
			return Point11Val0ClaimGovernanceStateBlocked
		}
		if !point11Val0IdentityValueValid(model.PolicyVersion) ||
			!point11Val0IdentityValueValid(model.Owner) ||
			!point11Val0ScopeValid(model.Scope) ||
			!point11Val0IdentityValueValid(model.RevocationPath) {
			return Point11Val0ClaimGovernanceStateBlocked
		}
	}
	if point11Val0PublicFacingSurface(model.PublicationBoundary) &&
		!point11Val0IdentityValueValid(model.CleanRoomIPReview) {
		return Point11Val0ClaimGovernanceStateBlocked
	}
	if model.PublicationBoundary == point11Val0PublicationSurfaceAgentOutput &&
		(model.ClaimCategory == Point11Val0ClaimCategoryCustomerVisible || model.ClaimCategory == Point11Val0ClaimCategoryPublicSafe) &&
		!point11Val0IdentityValueValid(model.GovernanceEvent) {
		return Point11Val0ClaimGovernanceStateBlocked
	}
	if point11Val0PublicFacingSurface(model.PublicationBoundary) && !point11Val0EvidenceRefsValid(model.EvidenceRefs) {
		return Point11Val0ClaimGovernanceStateBlocked
	}
	return Point11Val0ClaimGovernanceStateActive
}

func EvaluatePoint11Val0AuthorityMatrixState(model Point11Val0AuthorityMatrix) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0IdentityValueValid(model.ClaimType) ||
		!point11Val0IdentityValueValid(model.Proposer) ||
		!point11Val0IdentityValueValid(model.Reviewer) ||
		!point11Val0IdentityValueValid(model.FinalApprover) ||
		!point11Val0IdentityValueValid(model.Publisher) ||
		!point11Val0IdentityValueValid(model.Revoker) ||
		!point11Val0IdentityValueValid(model.Corrector) {
		return Point11Val0AuthorityMatrixStateBlocked
	}
	if model.Proposer == model.FinalApprover {
		switch point11Val0ActorClass(model.Proposer) {
		case "partner", "customer", "agent":
			return Point11Val0AuthorityMatrixStateBlocked
		}
		if model.CustomerVisibleOrPublic || model.ExportVisible {
			return Point11Val0AuthorityMatrixStateBlocked
		}
	}
	if (model.PolicyRelaxationRequested || model.ExceptionRequested || model.AuthorityExpansionRequested) && !point11Val0IdentityValueValid(model.GovernanceEvent) {
		return Point11Val0AuthorityMatrixStateBlocked
	}
	return Point11Val0AuthorityMatrixStateActive
}

func EvaluatePoint11Val0ExceptionGovernanceState(model Point11Val0ExceptionGovernance) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0IdentityValueValid(model.ExceptionID) ||
		!point11Val0EmergencyClaimRefValid(model.EmergencyClaimID) ||
		!point11Val0IdentityValueValid(model.Reason) ||
		!point11Val0IdentityValueValid(model.Issuer) ||
		!point11Val0IdentityValueValid(model.Approver) ||
		!point11Val0IdentityValueValid(model.Subject) ||
		!point11Val0ScopeValid(model.Scope) ||
		!point11Val0ValidTimestamp(model.IssuedAt) ||
		!point11Val0ValidTimestamp(model.ExpiresAt) ||
		!point11Val0IdentityValueValid(model.MonitoringRequirements) ||
		!point11Val0IdentityValueValid(model.RollbackOrReviewCondition) ||
		!point11Val0IdentityValueValid(model.RevocationPath) ||
		!point11Val0IdentityValueValid(model.AuditID) ||
		!point11Val0EvidenceRefsValid(model.EvidenceRefs) ||
		strings.TrimSpace(model.SignedGovernanceState) != point11Val0ExceptionSignedGovernanceState ||
		model.PermanentSilentException ||
		model.Revoked {
		return Point11Val0ExceptionGovernanceStateBlocked
	}
	expiresAt, _ := time.Parse(time.RFC3339, strings.TrimSpace(model.ExpiresAt))
	if expiresAt.Before(time.Now().UTC()) {
		return Point11Val0ExceptionGovernanceStateBlocked
	}
	return Point11Val0ExceptionGovernanceStateActive
}

func EvaluatePoint11Val0ABACGovernanceState(model Point11Val0ABACGovernance) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0IdentityValueValid(model.Subject) ||
		!point11Val0IdentityValueValid(model.Role) ||
		!point11Val0ScopeValid(model.Tenant) ||
		!point11Val0IdentityValueValid(model.Environment) ||
		!point11Val0IdentityValueValid(model.Artifact) ||
		!point11Val0IdentityValueValid(model.PolicyProfile) ||
		!point11Val0ContainsTrimmed(point11Val0PublicationSurfaces(), model.DecisionSurface) ||
		!point11Val0AllValuesValid(model.AllowedAttributes) ||
		!point11Val0AllValuesValid(model.PolicyRefs) ||
		!point11Val0AllValuesValid(model.ClaimRefs) ||
		len(model.ExplanationAttributes) == 0 ||
		len(model.ExplanationPolicyRefs) == 0 ||
		len(model.ExplanationClaimRefs) == 0 {
		return Point11Val0ABACStateBlocked
	}
	if len(model.UnknownAttributes) > 0 {
		return Point11Val0ABACStateBlocked
	}
	if len(model.DeniedAttributes) > 0 {
		if !point11Val0ContainsTrimmed(model.Diagnostics, "deny_over_allow_precedence_visible") {
			return Point11Val0ABACStateBlocked
		}
		return Point11Val0ABACStateBlocked
	}
	if strings.TrimSpace(model.ExceptionState) != "" && !point11Val0IdentityValueValid(model.ExceptionInteraction) {
		return Point11Val0ABACStateBlocked
	}
	return Point11Val0ABACStateActive
}

func point11Val0DecisionRefStateIsActive(state string) bool {
	return strings.TrimSpace(state) == point11Val0DecisionRefStateActive
}

func EvaluatePoint11Val0DecisionBindingState(model Point11Val0DecisionBinding) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0IdentityValueValid(model.DecisionID) ||
		!point11Val0IdentityValueValid(model.PolicyRef) ||
		!point11Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point11Val0ValidTimestamp(model.DecisionTimestamp) ||
		!point11Val0IdentityValueValid(model.DecisionOwnerOrSystemActor) ||
		!point11Val0IdentityValueValid(model.EnforcementOutcome) {
		return Point11Val0DecisionBindingStateBlocked
	}
	if !point11Val0DecisionRefStateIsActive(model.PolicyRefState) {
		return Point11Val0DecisionBindingStateBlocked
	}
	if model.ClaimRefApplicable {
		if !point11Val0IdentityValueValid(model.ClaimRef) || !point11Val0DecisionRefStateIsActive(model.ClaimRefState) {
			return Point11Val0DecisionBindingStateBlocked
		}
	}
	values := append([]string{model.EnforcementOutcome}, model.Diagnostics...)
	if point11Val0ContainsForbiddenClaim(values...) {
		return Point11Val0DecisionBindingStateBlocked
	}
	return Point11Val0DecisionBindingStateActive
}

func EvaluatePoint11Val0NoOverclaimState(model Point11Val0NoOverclaimDiscipline) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		point11Val0ContainsForbiddenClaim(model.ObservedClaims...) {
		return Point11Val0NoOverclaimStateBlocked
	}
	return Point11Val0NoOverclaimStateActive
}

func EvaluatePoint11Val0CrossDomainCompatibilityState(model Point11Val0CrossDomainCompatibility) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0IdentityValueValid(model.TrustRootRef) ||
		!point11Val0AllValuesValid(model.AcceptedClaimTypes) ||
		!point11Val0IdentityValueValid(model.RevocationHandling) ||
		!point11Val0IdentityValueValid(model.IncompatibilityBehavior) ||
		!point11Val0IdentityValueValid(model.LocalPolicyAdmissibilityRule) {
		return Point11Val0CrossDomainCompatibilityStateBlocked
	}
	if strings.TrimSpace(model.IssuerTrustRule) == point11Val0CrossDomainIssuerUnknown || !point11Val0IdentityValueValid(model.IssuerTrustRule) {
		return Point11Val0CrossDomainCompatibilityStateBlocked
	}
	if model.RemoteOverridesLocalPolicy || model.CreatesPublicAuthority || model.CreatesRegulatoryAuthority || model.CreatesCertificationAuthority {
		return Point11Val0CrossDomainCompatibilityStateBlocked
	}
	switch strings.TrimSpace(model.RemoteClaimState) {
	case point11Val0DecisionRefStateRevoked, point11Val0DecisionRefStateExpired:
		return Point11Val0CrossDomainCompatibilityStateBlocked
	case point11Val0DecisionRefStateUnknown:
		return Point11Val0CrossDomainCompatibilityStateReviewRequired
	}
	if strings.TrimSpace(model.ScopeCompatibility) == point11Val0CrossDomainScopeIncompatible ||
		strings.TrimSpace(model.FreshnessCompatibility) == point11Val0CrossDomainFreshnessIncompatible {
		return Point11Val0CrossDomainCompatibilityStateReviewRequired
	}
	if strings.TrimSpace(model.ScopeCompatibility) != point11Val0CrossDomainScopeCompatible ||
		strings.TrimSpace(model.FreshnessCompatibility) != point11Val0CrossDomainFreshnessCompatible {
		return Point11Val0CrossDomainCompatibilityStateBlocked
	}
	return Point11Val0CrossDomainCompatibilityStateActive
}

func EvaluatePoint11Val0State(model Point11Val0Foundation) string {
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point11Val0StateBlocked
	}
	states := []string{
		model.PolicyContractState,
		model.ClaimGovernanceState,
		model.AuthorityMatrixState,
		model.ExceptionGovernanceState,
		model.ABACGovernanceState,
		model.DecisionBindingState,
		model.NoOverclaimState,
		model.CrossDomainCompatibilityState,
	}
	if strings.TrimSpace(model.DependencyState) == Point11Val0DependencyStateBlocked {
		return Point11Val0StateBlocked
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point11Val0PolicyContractStateBlocked ||
			strings.TrimSpace(state) == Point11Val0ClaimGovernanceStateBlocked ||
			strings.TrimSpace(state) == Point11Val0AuthorityMatrixStateBlocked ||
			strings.TrimSpace(state) == Point11Val0ExceptionGovernanceStateBlocked ||
			strings.TrimSpace(state) == Point11Val0ABACStateBlocked ||
			strings.TrimSpace(state) == Point11Val0DecisionBindingStateBlocked ||
			strings.TrimSpace(state) == Point11Val0NoOverclaimStateBlocked ||
			strings.TrimSpace(state) == Point11Val0CrossDomainCompatibilityStateBlocked {
			return Point11Val0StateBlocked
		}
	}
	if strings.TrimSpace(model.DependencyState) == Point11Val0DependencyStateReviewRequired ||
		strings.TrimSpace(model.CrossDomainCompatibilityState) == Point11Val0CrossDomainCompatibilityStateReviewRequired {
		return Point11Val0StateReviewRequired
	}
	return Point11Val0StateActive
}

func point11Val0BlockingReasons(model Point11Val0Foundation) []string {
	reasons := []string{}
	if !point11Val0HasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState == Point11Val0DependencyStateBlocked {
		reasons = append(reasons, "point10_dependency_blocked")
	}
	if model.PolicyContractState == Point11Val0PolicyContractStateBlocked {
		reasons = append(reasons, "policy_contract_blocked")
	}
	if model.ClaimGovernanceState == Point11Val0ClaimGovernanceStateBlocked {
		reasons = append(reasons, "claim_governance_blocked")
	}
	if model.AuthorityMatrixState == Point11Val0AuthorityMatrixStateBlocked {
		reasons = append(reasons, "authority_matrix_blocked")
	}
	if model.ExceptionGovernanceState == Point11Val0ExceptionGovernanceStateBlocked {
		reasons = append(reasons, "exception_governance_blocked")
	}
	if model.ABACGovernanceState == Point11Val0ABACStateBlocked {
		reasons = append(reasons, "abac_governance_blocked")
	}
	if model.DecisionBindingState == Point11Val0DecisionBindingStateBlocked {
		reasons = append(reasons, "decision_binding_blocked")
	}
	if model.NoOverclaimState == Point11Val0NoOverclaimStateBlocked {
		reasons = append(reasons, "no_overclaim_blocked")
	}
	if model.CrossDomainCompatibilityState == Point11Val0CrossDomainCompatibilityStateBlocked {
		reasons = append(reasons, "cross_domain_compatibility_blocked")
	}
	return reasons
}

func Point11Val0FoundationModel() Point11Val0Foundation {
	disclaimer := point11Val0ProjectionDisclaimerBaseline
	dependency := point11Val0DependencySnapshotModel()
	return Point11Val0Foundation{
		CurrentState:                  Point11Val0StateReviewRequired,
		ProjectionDisclaimer:          disclaimer,
		DependencyState:               Point11Val0DependencyStateReviewRequired,
		PolicyContractState:           Point11Val0PolicyContractStateActive,
		ClaimGovernanceState:          Point11Val0ClaimGovernanceStateActive,
		AuthorityMatrixState:          Point11Val0AuthorityMatrixStateActive,
		ExceptionGovernanceState:      Point11Val0ExceptionGovernanceStateActive,
		ABACGovernanceState:           Point11Val0ABACStateActive,
		DecisionBindingState:          Point11Val0DecisionBindingStateActive,
		NoOverclaimState:              Point11Val0NoOverclaimStateActive,
		CrossDomainCompatibilityState: Point11Val0CrossDomainCompatibilityStateActive,
		Dependency:                    dependency,
		PolicyContract: Point11Val0PolicyContract{
			PolicyID:               "point11_val0_policy_contract",
			Version:                "point11_val0_policy_v1",
			Scope:                  "tenant_scoped_governance_foundation",
			Issuer:                 "governance_authority_team",
			Owner:                  "policy_owner_team",
			ApproverChain:          []string{"governance_reviewer", "governance_final_approver"},
			SignedState:            point11Val0PolicySignedState,
			AnchoredState:          point11Val0PolicyAnchoredState,
			EffectiveFrom:          "2026-05-02T10:00:00Z",
			EffectiveUntil:         "2099-01-01T00:00:00Z",
			SchemaVersion:          "point11_val0_schema_v1",
			CompatibilityVersion:   "point11_val0_compat_v1",
			ApprovalEvidenceRefs:   []string{"evidence:point11-policy-approval-001"},
			MethodologyOrRuleBasis: "signed_versioned_policy_methodology",
			ProjectionDisclaimer:   disclaimer,
		},
		ClaimGovernance: Point11Val0ClaimGovernance{
			LifecycleState:                 Point11Val0ClaimLifecycleDraft,
			ClaimID:                        "claim:point11-val0-bounded-governance",
			ClaimType:                      "bounded_governance_signal",
			Subject:                        "policy_contract_surface",
			Issuer:                         "governance_authority_team",
			Owner:                          "claims_owner_team",
			Audience:                       "internal_review",
			Scope:                          "tenant_scoped_governance_foundation",
			PolicyBasis:                    "point11_val0_policy_contract",
			PolicyVersion:                  "point11_val0_policy_v1",
			EvidenceRefs:                   []string{"evidence:point11-claim-evidence-001"},
			Freshness:                      "fresh",
			Expiry:                         "2099-01-01T00:00:00Z",
			VerificationMethod:             "manual_and_rule_validation",
			PublicationBoundary:            point11Val0PublicationSurfaceAgentOutput,
			ClaimCategory:                  Point11Val0ClaimCategoryAllowed,
			CleanRoomIPReview:              "clean_room_review_complete",
			RevocationOrSupersessionStatus: "active_current_claim",
			RevocationPath:                 "governance_revocation_path",
			ApprovalStatus:                 Point11Val0ClaimLifecycleApproved,
			GovernanceEvent:                "governance_event_claim_boundary",
			ProjectionDisclaimer:           disclaimer,
		},
		AuthorityMatrix: Point11Val0AuthorityMatrix{
			ClaimType:            "bounded_governance_signal",
			Proposer:             "policy_analyst_internal",
			Reviewer:             "governance_reviewer",
			FinalApprover:        "governance_final_approver",
			Publisher:            "governance_publisher",
			Revoker:              "governance_revoker",
			Corrector:            "governance_corrector",
			GovernanceEvent:      "governance_event_claim_boundary",
			ProjectionDisclaimer: disclaimer,
		},
		ExceptionGovernance: Point11Val0ExceptionGovernance{
			ExceptionID:               "exception:point11-val0-scope-bounded",
			EmergencyClaimID:          "point11_emergency_claim_governance_foundation",
			Reason:                    "bounded_governance_exception_review",
			Issuer:                    "governance_authority_team",
			Approver:                  "governance_final_approver",
			Subject:                   "policy_contract_surface",
			Scope:                     "tenant_scoped_governance_foundation",
			IssuedAt:                  "2026-05-02T10:00:00Z",
			ExpiresAt:                 "2099-01-01T00:00:00Z",
			MonitoringRequirements:    "exception_monitoring_required",
			RollbackOrReviewCondition: "exception_review_or_rollback_required",
			RevocationPath:            "exception_revocation_path",
			AuditID:                   "audit:point11-val0-exception-001",
			EvidenceRefs:              []string{"evidence:point11-exception-evidence-001"},
			SignedGovernanceState:     point11Val0ExceptionSignedGovernanceState,
			ProjectionDisclaimer:      disclaimer,
		},
		ABACGovernance: Point11Val0ABACGovernance{
			Subject:               "governance_subject",
			Role:                  "governance_reviewer",
			Tenant:                "tenant_scope_alpha",
			Environment:           "staging_governance",
			Artifact:              "policy_contract_surface",
			PolicyProfile:         "policy_profile_governance_val0",
			ClaimRefs:             []string{"claim:point11-val0-bounded-governance"},
			PolicyRefs:            []string{"point11_val0_policy_contract"},
			ExceptionState:        "exception_not_required",
			DecisionSurface:       point11Val0PublicationSurfaceAgentOutput,
			AllowedAttributes:     []string{"subject", "role", "tenant", "environment", "artifact", "policy_profile"},
			ExplanationAttributes: []string{"subject", "role", "tenant", "environment", "artifact", "policy_profile"},
			ExplanationPolicyRefs: []string{"point11_val0_policy_contract"},
			ExplanationClaimRefs:  []string{"claim:point11-val0-bounded-governance"},
			ExceptionInteraction:  "no_exception_interaction",
			ProjectionDisclaimer:  disclaimer,
		},
		DecisionBinding: Point11Val0DecisionBinding{
			DecisionID:                 "decision:point11-val0-governance-boundary",
			PolicyRef:                  "point11_val0_policy_contract",
			PolicyRefState:             point11Val0DecisionRefStateActive,
			ClaimRef:                   "claim:point11-val0-bounded-governance",
			ClaimRefApplicable:         true,
			ClaimRefState:              point11Val0DecisionRefStateActive,
			EvidenceRefs:               []string{"evidence:point11-decision-evidence-001"},
			DecisionTimestamp:          "2026-05-02T10:00:00Z",
			DecisionOwnerOrSystemActor: "governance_decision_engine",
			EnforcementOutcome:         "policy-bound decision support",
			Diagnostics:                []string{"advisory governance signal", "evidence support available for review"},
			ProjectionDisclaimer:       disclaimer,
		},
		NoOverclaim: Point11Val0NoOverclaimDiscipline{
			ObservedClaims:       []string{"signed and versioned policy contract", "not production approval"},
			ProjectionDisclaimer: disclaimer,
		},
		CrossDomainCompatibility: Point11Val0CrossDomainCompatibility{
			TrustRootRef:                 "trust_root_local_governance",
			IssuerTrustRule:              point11Val0CrossDomainIssuerTrusted,
			AcceptedClaimTypes:           []string{"bounded_governance_signal"},
			ScopeCompatibility:           point11Val0CrossDomainScopeCompatible,
			FreshnessCompatibility:       point11Val0CrossDomainFreshnessCompatible,
			RevocationHandling:           "remote_claim_revocation_enforced",
			IncompatibilityBehavior:      "review_required_on_incompatibility",
			LocalPolicyAdmissibilityRule: "local_policy_owns_final_admissibility",
			RemoteClaimState:             point11Val0DecisionRefStateActive,
			ProjectionDisclaimer:         disclaimer,
		},
	}
}

func ComputePoint11Val0Foundation(model Point11Val0Foundation) Point11Val0Foundation {
	model.DependencyState = EvaluatePoint11Val0DependencyState(model.Dependency)
	model.PolicyContractState = EvaluatePoint11Val0PolicyContractState(model.PolicyContract)
	model.ClaimGovernanceState = EvaluatePoint11Val0ClaimGovernanceState(model.ClaimGovernance)
	model.AuthorityMatrixState = EvaluatePoint11Val0AuthorityMatrixState(model.AuthorityMatrix)
	model.ExceptionGovernanceState = EvaluatePoint11Val0ExceptionGovernanceState(model.ExceptionGovernance)
	model.ABACGovernanceState = EvaluatePoint11Val0ABACGovernanceState(model.ABACGovernance)
	model.DecisionBindingState = EvaluatePoint11Val0DecisionBindingState(model.DecisionBinding)
	model.NoOverclaimState = EvaluatePoint11Val0NoOverclaimState(model.NoOverclaim)
	model.CrossDomainCompatibilityState = EvaluatePoint11Val0CrossDomainCompatibilityState(model.CrossDomainCompatibility)
	model.CurrentState = EvaluatePoint11Val0State(model)
	model.BlockingReasons = point11Val0BlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	return model
}
