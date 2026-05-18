package formal

import (
	"strings"
	"time"
)

const (
	Point11ValCStateActive         = "point11_valc_governance_enforcement_core_active"
	Point11ValCStateBlocked        = "point11_valc_governance_enforcement_core_blocked"
	Point11ValCStateReviewRequired = "point11_valc_governance_enforcement_core_review_required"

	Point11ValCDependencyStateActive         = "point11_valc_dependency_active"
	Point11ValCDependencyStateBlocked        = "point11_valc_dependency_blocked"
	Point11ValCDependencyStateReviewRequired = "point11_valc_dependency_review_required"

	Point11ValCEnforcementInputStateActive  = "point11_valc_enforcement_input_active"
	Point11ValCEnforcementInputStateBlocked = "point11_valc_enforcement_input_blocked"

	Point11ValCEnforcementResultStateActive         = "point11_valc_enforcement_result_active"
	Point11ValCEnforcementResultStateBlocked        = "point11_valc_enforcement_result_blocked"
	Point11ValCEnforcementResultStateReviewRequired = "point11_valc_enforcement_result_review_required"

	Point11ValCABACDecisionStateActive  = "point11_valc_abac_decision_active"
	Point11ValCABACDecisionStateBlocked = "point11_valc_abac_decision_blocked"

	Point11ValCExceptionDecisionStateActive         = "point11_valc_exception_decision_active"
	Point11ValCExceptionDecisionStateBlocked        = "point11_valc_exception_decision_blocked"
	Point11ValCExceptionDecisionStateReviewRequired = "point11_valc_exception_decision_review_required"

	Point11ValCPrecedenceStateActive         = "point11_valc_precedence_active"
	Point11ValCPrecedenceStateBlocked        = "point11_valc_precedence_blocked"
	Point11ValCPrecedenceStateReviewRequired = "point11_valc_precedence_review_required"

	Point11ValCMonitoringStateActive         = "point11_valc_monitoring_active"
	Point11ValCMonitoringStateBlocked        = "point11_valc_monitoring_blocked"
	Point11ValCMonitoringStateReviewRequired = "point11_valc_monitoring_review_required"

	Point11ValCDashboardStateActive  = "point11_valc_dashboard_active"
	Point11ValCDashboardStateBlocked = "point11_valc_dashboard_blocked"
)

const (
	point11ValCProjectionDisclaimerBaseline            = "projection_only not_canonical_truth point11_valc_governance_enforcement_core"
	point11ValCRequestedActionEvaluate                 = "evaluate_governance_decision"
	point11ValCRequestedActionObserve                  = "observe_governance_state"
	point11ValCRequestedActionReview                   = "review_governance_decision"
	point11ValCRequestedActionTemporaryOverride        = "temporary_override_candidate"
	point11ValCRequestedActionRenderDashboard          = "render_governance_dashboard"
	point11ValCRequestedActionDeploy                   = "deploy_to_production"
	point11ValCRequestedActionMutateEvidence           = "mutate_canonical_evidence"
	point11ValCRequestedActionExecuteEnforcement       = "execute_enforcement_action"
	point11ValCRequestedSurfaceInternalReview          = "internal_review"
	point11ValCRequestedSurfaceGovernanceDashboard     = "governance_dashboard"
	point11ValCRequestedSurfaceSupportConsole          = "support_console"
	point11ValCRequestedOutcomeAllow                   = "allow_bounded_governance_decision"
	point11ValCRequestedOutcomeDeny                    = "deny_governance_decision"
	point11ValCRequestedOutcomeReview                  = "review_required_governance_decision"
	point11ValCRequestedOutcomeTemporaryOverride       = "temporary_override_candidate"
	point11ValCABACPrecedenceDenyOverAllow             = "deny_over_allow"
	point11ValCABACExplanationVisible                  = "deny_over_allow_precedence_visible"
	point11ValCExceptionTypeScoped                     = "scoped_exception"
	point11ValCExceptionTypeTemporaryOverrideCandidate = "temporary_override_candidate"
	point11ValCExceptionTypeEmergencyOverrideCandidate = "emergency_override_candidate"
	point11ValCPolicyStateActive                       = "policy_state_active"
	point11ValCClaimStateActive                        = "claim_state_active"
	point11ValCEvidenceStateActive                     = "evidence_state_active"
	point11ValCScopeStateActive                        = "scope_state_active"
	point11ValCAuthorityStateActive                    = "authority_state_active"
	point11ValCABACResultAllow                         = "abac_allow"
	point11ValCABACResultDeny                          = "abac_deny"
	point11ValCLocalPolicyResultActive                 = "local_policy_active"
	point11ValCLocalPolicyResultDeny                   = "local_policy_deny"
	point11ValCLocalPolicyResultInvalid                = "local_policy_invalid"
	point11ValCLocalPolicyResultReview                 = "local_policy_review_required"
	point11ValCRemoteClaimResultCompatible             = "remote_claim_compatible"
	point11ValCRemoteClaimResultBlocked                = "remote_claim_blocked"
	point11ValCRemoteClaimResultReview                 = "remote_claim_review_required"
	point11ValCCheckStateActive                        = "check_active"
	point11ValCCheckStateBlocked                       = "check_blocked"
	point11ValCMonitoringSignalFresh                   = "signal_fresh"
	point11ValCMonitoringSignalStale                   = "signal_stale"
	point11ValCMonitoringLinkActive                    = "monitoring_link_active"
	point11ValCDashboardRenderedStateBounded           = "bounded_governance_dashboard"
)

type Point11ValCValBReviewContext struct {
	LocalReviewAllowsDependencyReviewRequired bool     `json:"local_review_allows_dependency_review_required"`
	ValBPoint11PassEmitted                    bool     `json:"valb_point11_pass_emitted"`
	ValBCreatesAuthorityClaims                bool     `json:"valb_creates_authority_claims"`
	ValBCreatesPublicationSideEffects         bool     `json:"valb_creates_publication_side_effects"`
	ValBCreatesSigningSideEffects             bool     `json:"valb_creates_signing_side_effects"`
	ValBCreatesAnchoringSideEffects           bool     `json:"valb_creates_anchoring_side_effects"`
	ValBCreatesExternalAPISideEffects         bool     `json:"valb_creates_external_api_side_effects"`
	ValBCreatesProductionSideEffects          bool     `json:"valb_creates_production_side_effects"`
	OpenCLB0Findings                          int      `json:"open_clb0_findings"`
	OpenCLB1Findings                          int      `json:"open_clb1_findings"`
	OpenCLB2Findings                          int      `json:"open_clb2_findings"`
	ReviewPrerequisites                       []string `json:"review_prerequisites,omitempty"`
}

type Point11ValCDependencySnapshot struct {
	ValBCurrentState                  string   `json:"valb_current_state"`
	ValBDependencyState               string   `json:"valb_dependency_state"`
	ValBClaimTypeState                string   `json:"valb_claim_type_state"`
	ValBIssuanceRequestState          string   `json:"valb_issuance_request_state"`
	ValBIssuedClaimState              string   `json:"valb_issued_claim_state"`
	ValBRegistryState                 string   `json:"valb_registry_state"`
	ValBVerificationState             string   `json:"valb_verification_state"`
	ValBCrossDomainIntakeState        string   `json:"valb_cross_domain_intake_state"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
	ValBPoint11PassEmitted            bool     `json:"valb_point11_pass_emitted"`
	ValBCreatesAuthorityClaims        bool     `json:"valb_creates_authority_claims"`
	ValBCreatesPublicationSideEffects bool     `json:"valb_creates_publication_side_effects"`
	ValBCreatesSigningSideEffects     bool     `json:"valb_creates_signing_side_effects"`
	ValBCreatesAnchoringSideEffects   bool     `json:"valb_creates_anchoring_side_effects"`
	ValBCreatesExternalAPISideEffects bool     `json:"valb_creates_external_api_side_effects"`
	ValBCreatesProductionSideEffects  bool     `json:"valb_creates_production_side_effects"`
	OpenCLB0Findings                  int      `json:"open_clb0_findings"`
	OpenCLB1Findings                  int      `json:"open_clb1_findings"`
	OpenCLB2Findings                  int      `json:"open_clb2_findings"`
	LocalReviewAllowsReviewRequired   bool     `json:"local_review_allows_review_required"`
	ReviewPrerequisites               []string `json:"review_prerequisites,omitempty"`
}

type Point11ValCGovernanceEnforcementInput struct {
	CurrentState           string   `json:"current_state"`
	EnforcementID          string   `json:"enforcement_id"`
	DecisionID             string   `json:"decision_id"`
	SubjectRef             string   `json:"subject_ref"`
	SubjectKind            string   `json:"subject_kind"`
	ActorRef               string   `json:"actor_ref"`
	ActorKind              string   `json:"actor_kind"`
	TenantScope            string   `json:"tenant_scope"`
	EnvironmentRef         string   `json:"environment_ref"`
	ArtifactRef            string   `json:"artifact_ref"`
	PolicyBasisRef         string   `json:"policy_basis_ref"`
	PolicyBasisState       string   `json:"policy_basis_state"`
	PolicyVersion          string   `json:"policy_version"`
	ClaimsRequired         bool     `json:"claims_required"`
	ClaimRefs              []string `json:"claim_refs,omitempty"`
	ClaimVerificationRefs  []string `json:"claim_verification_refs,omitempty"`
	ClaimVerificationState string   `json:"claim_verification_state"`
	RegistryRef            string   `json:"registry_ref"`
	RegistryState          string   `json:"registry_state"`
	ABACContextRef         string   `json:"abac_context_ref"`
	ExceptionRefs          []string `json:"exception_refs,omitempty"`
	EmergencyRefs          []string `json:"emergency_refs,omitempty"`
	RequestedAction        string   `json:"requested_action"`
	RequestedSurface       string   `json:"requested_surface"`
	RequestedOutcome       string   `json:"requested_outcome"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs       []string `json:"evidence_hash_refs,omitempty"`
	GovernanceEventRef     string   `json:"governance_event_ref"`
	AuditID                string   `json:"audit_id"`
	DecisionTimestamp      string   `json:"decision_timestamp"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
}

type Point11ValCPolicyDecisionEnforcementResult struct {
	CurrentState           string   `json:"current_state"`
	EnforcementResultID    string   `json:"enforcement_result_id"`
	EnforcementID          string   `json:"enforcement_id"`
	DecisionID             string   `json:"decision_id"`
	PolicyBasisRef         string   `json:"policy_basis_ref"`
	ClaimRefs              []string `json:"claim_refs,omitempty"`
	ClaimVerificationRefs  []string `json:"claim_verification_refs,omitempty"`
	ABACDecisionRef        string   `json:"abac_decision_ref"`
	ExceptionDecisionRef   string   `json:"exception_decision_ref"`
	EmergencyDecisionRef   string   `json:"emergency_decision_ref"`
	PolicyResultState      string   `json:"policy_result_state"`
	ClaimResultState       string   `json:"claim_result_state"`
	ABACDecisionState      string   `json:"abac_decision_state"`
	ExceptionDecisionState string   `json:"exception_decision_state"`
	EmergencyDecisionState string   `json:"emergency_decision_state"`
	EvidenceState          string   `json:"evidence_state"`
	ScopeState             string   `json:"scope_state"`
	AuthorityState         string   `json:"authority_state"`
	EnforcementState       string   `json:"enforcement_state"`
	EnforcementOutcome     string   `json:"enforcement_outcome"`
	AllowedAction          string   `json:"allowed_action"`
	BlockedActionReason    string   `json:"blocked_action_reason"`
	ReviewRequiredReason   string   `json:"review_required_reason"`
	EffectivePolicyVersion string   `json:"effective_policy_version"`
	EffectiveClaimVersions []string `json:"effective_claim_versions,omitempty"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs       []string `json:"evidence_hash_refs,omitempty"`
	AuditID                string   `json:"audit_id"`
	Diagnostics            []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
}

type Point11ValCABACEnforcementDecision struct {
	CurrentState                string   `json:"current_state"`
	ABACDecisionID              string   `json:"abac_decision_id"`
	SubjectRef                  string   `json:"subject_ref"`
	SubjectAttributes           []string `json:"subject_attributes,omitempty"`
	ActorRef                    string   `json:"actor_ref"`
	ActorAttributes             []string `json:"actor_attributes,omitempty"`
	TenantScope                 string   `json:"tenant_scope"`
	EnvironmentAttributes       []string `json:"environment_attributes,omitempty"`
	ArtifactAttributes          []string `json:"artifact_attributes,omitempty"`
	PolicyProfileRef            string   `json:"policy_profile_ref"`
	ClaimRefs                   []string `json:"claim_refs,omitempty"`
	ExceptionRefs               []string `json:"exception_refs,omitempty"`
	RequestedAction             string   `json:"requested_action"`
	RequestedSurface            string   `json:"requested_surface"`
	AllowedAttributes           []string `json:"allowed_attributes,omitempty"`
	DeniedAttributes            []string `json:"denied_attributes,omitempty"`
	UnknownAttributes           []string `json:"unknown_attributes,omitempty"`
	PrecedenceRule              string   `json:"precedence_rule"`
	DecisionState               string   `json:"decision_state"`
	Explanation                 string   `json:"explanation"`
	Diagnostics                 []string `json:"diagnostics,omitempty"`
	AuditID                     string   `json:"audit_id"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
	PolicyState                 string   `json:"policy_state"`
	ClaimState                  string   `json:"claim_state"`
	EvidenceState               string   `json:"evidence_state"`
	ExceptionState              string   `json:"exception_state"`
	ExceptionScoped             bool     `json:"exception_scoped"`
	ExceptionExpired            bool     `json:"exception_expired"`
	ExceptionRevocable          bool     `json:"exception_revocable"`
	ExceptionGovernanceApproved bool     `json:"exception_governance_approved"`
}

type Point11ValCExceptionEmergencyDecision struct {
	CurrentState                  string   `json:"current_state"`
	ExceptionDecisionID           string   `json:"exception_decision_id"`
	ExceptionRef                  string   `json:"exception_ref"`
	EmergencyRef                  string   `json:"emergency_ref"`
	ExceptionType                 string   `json:"exception_type"`
	SubjectRef                    string   `json:"subject_ref"`
	TenantScope                   string   `json:"tenant_scope"`
	Reason                        string   `json:"reason"`
	IssuerRef                     string   `json:"issuer_ref"`
	ApproverRef                   string   `json:"approver_ref"`
	AuthorityBasisRef             string   `json:"authority_basis_ref"`
	GovernanceEventRef            string   `json:"governance_event_ref"`
	IssuedAt                      string   `json:"issued_at"`
	ExpiresAt                     string   `json:"expires_at"`
	RevocationPathRef             string   `json:"revocation_path_ref"`
	MonitoringRequirementRef      string   `json:"monitoring_requirement_ref"`
	RollbackOrReviewConditionRef  string   `json:"rollback_or_review_condition_ref"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	AuditID                       string   `json:"audit_id"`
	DecisionState                 string   `json:"decision_state"`
	Diagnostics                   []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	PermanentSilentException      bool     `json:"permanent_silent_exception"`
	EmergencyClaimState           string   `json:"emergency_claim_state"`
	CreatesProductionApproval     bool     `json:"creates_production_approval"`
	MutatesCanonicalEvidenceSpine bool     `json:"mutates_canonical_evidence_spine"`
}

type Point11ValCOverridePrecedence struct {
	CurrentState                string   `json:"current_state"`
	PrecedenceID                string   `json:"precedence_id"`
	BaseDecisionRef             string   `json:"base_decision_ref"`
	ABACDecisionRef             string   `json:"abac_decision_ref"`
	ExceptionDecisionRef        string   `json:"exception_decision_ref"`
	EmergencyDecisionRef        string   `json:"emergency_decision_ref"`
	ClaimVerificationRefs       []string `json:"claim_verification_refs,omitempty"`
	LocalPolicyResult           string   `json:"local_policy_result"`
	ClaimResultState            string   `json:"claim_result_state"`
	ABACResult                  string   `json:"abac_result"`
	RemoteClaimResult           string   `json:"remote_claim_result"`
	ExceptionResult             string   `json:"exception_result"`
	EmergencyResult             string   `json:"emergency_result"`
	FinalPrecedenceRule         string   `json:"final_precedence_rule"`
	FinalState                  string   `json:"final_state"`
	Diagnostics                 []string `json:"diagnostics,omitempty"`
	AuditID                     string   `json:"audit_id"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
	ExceptionScoped             bool     `json:"exception_scoped"`
	ExceptionExpired            bool     `json:"exception_expired"`
	ExceptionRevoked            bool     `json:"exception_revoked"`
	ExceptionGovernanceApproved bool     `json:"exception_governance_approved"`
	EmergencyScoped             bool     `json:"emergency_scoped"`
	EmergencyTimeBound          bool     `json:"emergency_time_bound"`
	EmergencyMonitored          bool     `json:"emergency_monitored"`
	EmergencyRevocable          bool     `json:"emergency_revocable"`
	GovernanceEventResolved     bool     `json:"governance_event_resolved"`
	RemoteClaimOverridesLocal   bool     `json:"remote_claim_overrides_local"`
}

type Point11ValCMonitoringLinkedEmergencyHandling struct {
	CurrentState             string   `json:"current_state"`
	MonitoringLinkID         string   `json:"monitoring_link_id"`
	EmergencyRef             string   `json:"emergency_ref"`
	MonitoringRequirementRef string   `json:"monitoring_requirement_ref"`
	MonitoringState          string   `json:"monitoring_state"`
	SignalRefs               []string `json:"signal_refs,omitempty"`
	SignalFreshness          string   `json:"signal_freshness"`
	EscalationRef            string   `json:"escalation_ref"`
	ReviewDeadline           string   `json:"review_deadline"`
	ExpiryEnforcementState   string   `json:"expiry_enforcement_state"`
	RevocationCheckState     string   `json:"revocation_check_state"`
	RollbackCheckState       string   `json:"rollback_check_state"`
	AuditID                  string   `json:"audit_id"`
	Diagnostics              []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
	HighRiskEmergency        bool     `json:"high_risk_emergency"`
	CreatesActionSideEffects bool     `json:"creates_action_side_effects"`
}

type Point11ValCGovernanceDashboardReadModel struct {
	CurrentState                  string   `json:"current_state"`
	DashboardViewID               string   `json:"dashboard_view_id"`
	SourceDecisionRefs            []string `json:"source_decision_refs,omitempty"`
	SourceClaimRefs               []string `json:"source_claim_refs,omitempty"`
	SourceExceptionRefs           []string `json:"source_exception_refs,omitempty"`
	SourceEmergencyRefs           []string `json:"source_emergency_refs,omitempty"`
	SourceMonitoringRefs          []string `json:"source_monitoring_refs,omitempty"`
	RenderedState                 string   `json:"rendered_state"`
	VisibleSurfaces               []string `json:"visible_surfaces,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	CreatesPublicationSideEffects bool     `json:"creates_publication_side_effects"`
	CreatesAuthorityClaim         bool     `json:"creates_authority_claim"`
	MutatesCanonicalState         bool     `json:"mutates_canonical_state"`
	Diagnostics                   []string `json:"diagnostics,omitempty"`
}

type Point11ValCDiagnostics struct {
	CurrentState             string   `json:"current_state"`
	BlockingReasons          []string `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites      []string `json:"review_prerequisites,omitempty"`
	ComponentStates          []string `json:"component_states,omitempty"`
	DependencyReasons        []string `json:"dependency_reasons,omitempty"`
	EnforcementInputReasons  []string `json:"enforcement_input_reasons,omitempty"`
	EnforcementResultReasons []string `json:"enforcement_result_reasons,omitempty"`
	ABACReasons              []string `json:"abac_reasons,omitempty"`
	ExceptionReasons         []string `json:"exception_reasons,omitempty"`
	PrecedenceReasons        []string `json:"precedence_reasons,omitempty"`
	MonitoringReasons        []string `json:"monitoring_reasons,omitempty"`
	DashboardReasons         []string `json:"dashboard_reasons,omitempty"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type Point11ValCFoundation struct {
	CurrentState                      string                                       `json:"current_state"`
	BlockingReasons                   []string                                     `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites               []string                                     `json:"review_prerequisites,omitempty"`
	DependencyState                   string                                       `json:"dependency_state"`
	EnforcementInputState             string                                       `json:"enforcement_input_state"`
	EnforcementResultState            string                                       `json:"enforcement_result_state"`
	ABACDecisionState                 string                                       `json:"abac_decision_state"`
	ExceptionDecisionState            string                                       `json:"exception_decision_state"`
	PrecedenceState                   string                                       `json:"precedence_state"`
	MonitoringState                   string                                       `json:"monitoring_state"`
	DashboardState                    string                                       `json:"dashboard_state"`
	Diagnostics                       Point11ValCDiagnostics                       `json:"diagnostics"`
	ProjectionDisclaimer              string                                       `json:"projection_disclaimer"`
	CreatesAuthorityClaims            bool                                         `json:"creates_authority_claims"`
	CreatesPublicationSideEffects     bool                                         `json:"creates_publication_side_effects"`
	CreatesRealEnforcementSideEffects bool                                         `json:"creates_real_enforcement_side_effects"`
	Dependency                        Point11ValCDependencySnapshot                `json:"dependency"`
	EnforcementInput                  Point11ValCGovernanceEnforcementInput        `json:"enforcement_input"`
	EnforcementResult                 Point11ValCPolicyDecisionEnforcementResult   `json:"enforcement_result"`
	ABACDecision                      Point11ValCABACEnforcementDecision           `json:"abac_decision"`
	ExceptionDecision                 Point11ValCExceptionEmergencyDecision        `json:"exception_decision"`
	Precedence                        Point11ValCOverridePrecedence                `json:"precedence"`
	Monitoring                        Point11ValCMonitoringLinkedEmergencyHandling `json:"monitoring"`
	Dashboard                         Point11ValCGovernanceDashboardReadModel      `json:"dashboard"`
}

func point11ValCContainsForbiddenText(values ...string) bool {
	if point11Val0ContainsForbiddenClaim(values...) {
		return true
	}
	sourcePhrase := "source of" + " truth"
	canonicalPhrase := "canonical" + " truth"
	normalizedValues := make([]string, 0, len(values))
	allowedValues := make([]bool, 0, len(values))
	for _, value := range values {
		normalized := point11Val0NormalizeText(value)
		if strings.Contains(normalized, sourcePhrase) || strings.Contains(normalized, canonicalPhrase) {
			return true
		}
		if normalized != "" {
			normalizedValues = append(normalizedValues, normalized)
			allowedValues = append(allowedValues, false)
		}
	}
	for _, phrase := range []string{sourcePhrase, canonicalPhrase} {
		if formalNoOverclaimForbiddenPhraseAcrossValues(normalizedValues, allowedValues, phrase) {
			return true
		}
	}
	return false
}

func point11ValCEnforcementRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"enforcement_", "point11_enforcement_"})
}

func point11ValCDecisionRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"decision_", "point11_decision_"})
}

func point11ValCABACRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"abac_", "point11_abac_"})
}

func point11ValCExceptionRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"exception_", "point11_exception_"})
}

func point11ValCEmergencyRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"emergency_", "point11_emergency_"})
}

func point11ValCMonitoringRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"monitoring_", "point11_monitoring_"})
}

func point11ValCDashboardRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"dashboard_", "point11_dashboard_"})
}

func point11ValCActorRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"actor_", "point11_actor_", "internal_", "partner_", "customer_", "agent_"})
}

func point11ValCTenantRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"tenant_", "point11_tenant_"})
}

func point11ValCArtifactRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"artifact_", "point11_artifact_"})
}

func point11ValCEnvironmentRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"environment_", "point11_environment_"})
}

func point11ValCAuthorityBasisRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"authority_basis_", "point11_authority_basis_"})
}

func point11ValCRevocationPathRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"revocation_path_", "point11_revocation_path_"})
}

func point11ValCRollbackReviewRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"rollback_review_", "point11_rollback_review_"})
}

func point11ValCEscalationRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"escalation_", "point11_escalation_"})
}

func point11ValCSignalRefsValid(values []string) bool {
	return point11Val0AllValuesValid(values)
}

func point11ValCExceptionRefListValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11ValCExceptionRefValid(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValCEmergencyRefListValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11ValCEmergencyRefValid(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValCDecisionRefListValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11ValCDecisionRefValid(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValCMonitoringRefListValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11ValCMonitoringRefValid(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValCRequestedActionSupported(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValCRequestedActionEvaluate,
		point11ValCRequestedActionObserve,
		point11ValCRequestedActionReview,
		point11ValCRequestedActionTemporaryOverride,
		point11ValCRequestedActionRenderDashboard,
	}, value)
}

func point11ValCRequestedActionImpliesRealSideEffect(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValCRequestedActionDeploy,
		point11ValCRequestedActionMutateEvidence,
		point11ValCRequestedActionExecuteEnforcement,
	}, value)
}

func point11ValCRequestedSurfaceAllowed(value string) bool {
	return point11Val0ContainsTrimmed(append(point11Val0PublicationSurfaces(),
		point11ValCRequestedSurfaceInternalReview,
		point11ValCRequestedSurfaceGovernanceDashboard,
		point11ValCRequestedSurfaceSupportConsole,
	), value)
}

func point11ValCRequestedOutcomeAllowed(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValCRequestedOutcomeAllow,
		point11ValCRequestedOutcomeDeny,
		point11ValCRequestedOutcomeReview,
		point11ValCRequestedOutcomeTemporaryOverride,
	}, value)
}

func point11ValCExceptionTypeAllowed(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValCExceptionTypeScoped,
		point11ValCExceptionTypeTemporaryOverrideCandidate,
		point11ValCExceptionTypeEmergencyOverrideCandidate,
	}, value)
}

func point11ValCVisibleSurfacesValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11Val0ContainsTrimmed([]string{
			point11ValCRequestedSurfaceInternalReview,
			point11ValCRequestedSurfaceGovernanceDashboard,
			point11ValCRequestedSurfaceSupportConsole,
		}, value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValCGenericStateActive(value string) bool {
	return value == point11ValCCheckStateActive
}

func point11ValCPrecedenceRuleAllowed(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point11ValCABACPrecedenceDenyOverAllow,
		"blocked_over_review_required_over_active",
		"temporary_override_requires_governance_resolution",
	}, value)
}

func point11ValCExceptionOverrideEligible(model Point11ValCOverridePrecedence) bool {
	return model.ExceptionScoped &&
		!model.ExceptionExpired &&
		!model.ExceptionRevoked &&
		model.ExceptionGovernanceApproved
}

func point11ValCEmergencyOverrideEligible(model Point11ValCOverridePrecedence) bool {
	return model.EmergencyScoped &&
		model.EmergencyTimeBound &&
		model.EmergencyMonitored &&
		model.EmergencyRevocable
}

func point11ValCComponentStates(model Point11ValCFoundation) []string {
	return []string{
		"dependency:" + model.DependencyState,
		"enforcement_input:" + model.EnforcementInputState,
		"enforcement_result:" + model.EnforcementResultState,
		"abac:" + model.ABACDecisionState,
		"exception:" + model.ExceptionDecisionState,
		"precedence:" + model.PrecedenceState,
		"monitoring:" + model.MonitoringState,
		"dashboard:" + model.DashboardState,
	}
}

func SnapshotPoint11ValCDependencyFromComputedValB(valB Point11ValBFoundation, review Point11ValCValBReviewContext) Point11ValCDependencySnapshot {
	reviewPrerequisites := append([]string{}, valB.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point11ValCDependencySnapshot{
		ValBCurrentState:                  valB.CurrentState,
		ValBDependencyState:               valB.DependencyState,
		ValBClaimTypeState:                valB.ClaimTypeState,
		ValBIssuanceRequestState:          valB.IssuanceRequestState,
		ValBIssuedClaimState:              valB.IssuedClaimState,
		ValBRegistryState:                 valB.RegistryState,
		ValBVerificationState:             valB.VerificationState,
		ValBCrossDomainIntakeState:        valB.CrossDomainIntakeState,
		ProjectionDisclaimer:              valB.ProjectionDisclaimer,
		ValBPoint11PassEmitted:            review.ValBPoint11PassEmitted,
		ValBCreatesAuthorityClaims:        review.ValBCreatesAuthorityClaims,
		ValBCreatesPublicationSideEffects: review.ValBCreatesPublicationSideEffects,
		ValBCreatesSigningSideEffects:     review.ValBCreatesSigningSideEffects,
		ValBCreatesAnchoringSideEffects:   review.ValBCreatesAnchoringSideEffects,
		ValBCreatesExternalAPISideEffects: review.ValBCreatesExternalAPISideEffects,
		ValBCreatesProductionSideEffects:  review.ValBCreatesProductionSideEffects,
		OpenCLB0Findings:                  review.OpenCLB0Findings,
		OpenCLB1Findings:                  review.OpenCLB1Findings,
		OpenCLB2Findings:                  review.OpenCLB2Findings,
		LocalReviewAllowsReviewRequired:   review.LocalReviewAllowsDependencyReviewRequired,
		ReviewPrerequisites:               reviewPrerequisites,
	}
}

func point11ValCDependencyReviewContextModel() Point11ValCValBReviewContext {
	return Point11ValCValBReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	}
}

func point11ValCDependencySnapshotModel() Point11ValCDependencySnapshot {
	valB := ComputePoint11ValBFoundation(Point11ValBFoundationModel())
	return SnapshotPoint11ValCDependencyFromComputedValB(valB, point11ValCDependencyReviewContextModel())
}

func point11ValCDependencyStateAndReasons(model Point11ValCDependencySnapshot) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "valb_projection_disclaimer_blocked")
	}
	if model.ValBPoint11PassEmitted {
		reasons = append(reasons, "valb_point11_pass_emitted")
	}
	if model.ValBCreatesAuthorityClaims {
		reasons = append(reasons, "valb_authority_claim_surface_blocked")
	}
	if model.ValBCreatesPublicationSideEffects {
		reasons = append(reasons, "valb_publication_side_effects_blocked")
	}
	if model.ValBCreatesSigningSideEffects {
		reasons = append(reasons, "valb_signing_side_effects_blocked")
	}
	if model.ValBCreatesAnchoringSideEffects {
		reasons = append(reasons, "valb_anchoring_side_effects_blocked")
	}
	if model.ValBCreatesExternalAPISideEffects {
		reasons = append(reasons, "valb_external_api_side_effects_blocked")
	}
	if model.ValBCreatesProductionSideEffects {
		reasons = append(reasons, "valb_production_side_effects_blocked")
	}
	if model.OpenCLB0Findings > 0 {
		reasons = append(reasons, "valb_open_clb0_findings")
	}
	if model.OpenCLB1Findings > 0 {
		reasons = append(reasons, "valb_open_clb1_findings")
	}
	if model.OpenCLB2Findings > 0 {
		reasons = append(reasons, "valb_open_clb2_findings")
	}
	if model.ValBClaimTypeState != Point11ValBClaimTypeStateActive {
		reasons = append(reasons, "valb_claim_type_not_active")
	}
	if model.ValBIssuanceRequestState != Point11ValBIssuanceRequestStateActive {
		reasons = append(reasons, "valb_issuance_request_not_active")
	}
	if model.ValBIssuedClaimState != Point11ValBIssuedClaimStateActive {
		reasons = append(reasons, "valb_issued_claim_not_active")
	}
	if model.ValBRegistryState != Point11ValBRegistryStateActive {
		reasons = append(reasons, "valb_registry_not_active")
	}
	if model.ValBVerificationState != Point11ValBVerificationStateActive {
		reasons = append(reasons, "valb_verification_not_active")
	}
	if model.ValBCrossDomainIntakeState == Point11ValBCrossDomainIntakeStateBlocked {
		reasons = append(reasons, "valb_cross_domain_intake_blocked")
	}
	if len(reasons) > 0 {
		return Point11ValCDependencyStateBlocked, reasons
	}
	if model.ValBCurrentState == Point11ValBStateActive &&
		model.ValBDependencyState == Point11ValBDependencyStateActive &&
		model.ValBCrossDomainIntakeState == Point11ValBCrossDomainIntakeStateActive &&
		len(model.ReviewPrerequisites) == 0 {
		return Point11ValCDependencyStateActive, nil
	}
	if model.LocalReviewAllowsReviewRequired &&
		(len(model.ReviewPrerequisites) > 0 ||
			model.ValBCurrentState == Point11ValBStateReviewRequired ||
			model.ValBDependencyState == Point11ValBDependencyStateReviewRequired ||
			model.ValBCrossDomainIntakeState == Point11ValBCrossDomainIntakeStateReviewRequired) {
		return Point11ValCDependencyStateReviewRequired, []string{"valb_dependency_review_required"}
	}
	return Point11ValCDependencyStateBlocked, []string{"valb_dependency_not_active"}
}

func EvaluatePoint11ValCDependencyState(model Point11ValCDependencySnapshot) string {
	state, _ := point11ValCDependencyStateAndReasons(model)
	return state
}

func point11ValCEnforcementInputStateAndReasons(model Point11ValCGovernanceEnforcementInput) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "enforcement_input_projection_disclaimer_blocked")
	}
	if !point11ValCEnforcementRefValid(model.EnforcementID) {
		reasons = append(reasons, "enforcement_input_id_invalid")
	}
	if !point11ValCDecisionRefValid(model.DecisionID) {
		reasons = append(reasons, "enforcement_input_decision_id_invalid")
	}
	if !point11ValBSubjectRefValid(model.SubjectRef) || !point11Val0IdentityValueValid(model.SubjectKind) {
		reasons = append(reasons, "enforcement_input_subject_invalid")
	}
	if !point11ValCActorRefValid(model.ActorRef) || !point11Val0IdentityValueValid(model.ActorKind) {
		reasons = append(reasons, "enforcement_input_actor_invalid")
	}
	if !point11Val0ScopeValid(model.TenantScope) {
		reasons = append(reasons, "enforcement_input_tenant_scope_invalid")
	}
	if !point11ValCEnvironmentRefValid(model.EnvironmentRef) {
		reasons = append(reasons, "enforcement_input_environment_ref_invalid")
	}
	if !point11ValCArtifactRefValid(model.ArtifactRef) {
		reasons = append(reasons, "enforcement_input_artifact_ref_invalid")
	}
	if !point11ValAPolicyRefValid(model.PolicyBasisRef) || model.PolicyBasisState != point11ValBPolicyBasisStateActive {
		reasons = append(reasons, "enforcement_input_policy_basis_invalid")
	}
	if !point11Val0IdentityValueValid(model.PolicyVersion) {
		reasons = append(reasons, "enforcement_input_policy_version_invalid")
	}
	if model.ClaimsRequired && !point11ValBClaimRefListValid(model.ClaimRefs) {
		reasons = append(reasons, "enforcement_input_claim_refs_invalid")
	}
	if !point11ValBVerificationRefValid(model.ABACContextRef) && !point11ValCABACRefValid(model.ABACContextRef) {
		reasons = append(reasons, "enforcement_input_abac_context_ref_invalid")
	}
	if len(model.ClaimVerificationRefs) == 0 || !point11ValBClaimRefListValid(model.ClaimVerificationRefs) {
		reasons = append(reasons, "enforcement_input_claim_verification_refs_invalid")
	}
	if model.ClaimVerificationState != Point11ValBVerificationStateActive {
		reasons = append(reasons, "enforcement_input_claim_verification_not_active")
	}
	if !point11ValBClaimRegistryRefValid(model.RegistryRef) || model.RegistryState != Point11ValBRegistryStateActive {
		reasons = append(reasons, "enforcement_input_registry_invalid")
	}
	if len(model.ExceptionRefs) > 0 && !point11ValCExceptionRefListValid(model.ExceptionRefs) {
		reasons = append(reasons, "enforcement_input_exception_refs_invalid")
	}
	if len(model.EmergencyRefs) > 0 && !point11ValCEmergencyRefListValid(model.EmergencyRefs) {
		reasons = append(reasons, "enforcement_input_emergency_refs_invalid")
	}
	if point11ValCRequestedActionImpliesRealSideEffect(model.RequestedAction) {
		reasons = append(reasons, "enforcement_input_requested_action_real_side_effect")
	} else if !point11ValCRequestedActionSupported(model.RequestedAction) {
		reasons = append(reasons, "enforcement_input_requested_action_unsupported")
	}
	if !point11ValCRequestedSurfaceAllowed(model.RequestedSurface) {
		reasons = append(reasons, "enforcement_input_requested_surface_invalid")
	}
	if !point11ValCRequestedOutcomeAllowed(model.RequestedOutcome) || point11ValCContainsForbiddenText(model.RequestedOutcome) {
		reasons = append(reasons, "enforcement_input_requested_outcome_invalid_or_overclaim")
	}
	if !point11Val0EvidenceRefsValid(model.EvidenceRefs) {
		reasons = append(reasons, "enforcement_input_evidence_refs_invalid")
	}
	if len(model.EvidenceHashRefs) > 0 && !point11ValBEvidenceHashRefsValid(model.EvidenceHashRefs) {
		reasons = append(reasons, "enforcement_input_evidence_hash_refs_invalid")
	}
	if point11Val0PublicFacingSurface(model.RequestedSurface) && !point11ValBGovernanceEventRefValid(model.GovernanceEventRef) {
		reasons = append(reasons, "enforcement_input_governance_event_missing")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "enforcement_input_audit_id_invalid")
	}
	if !point11Val0ValidTimestamp(model.DecisionTimestamp) {
		reasons = append(reasons, "enforcement_input_decision_timestamp_invalid")
	}
	if len(reasons) > 0 {
		return Point11ValCEnforcementInputStateBlocked, reasons
	}
	return Point11ValCEnforcementInputStateActive, nil
}

func EvaluatePoint11ValCEnforcementInputState(model Point11ValCGovernanceEnforcementInput) string {
	state, _ := point11ValCEnforcementInputStateAndReasons(model)
	return state
}

func point11ValCEnforcementResultStateAndReasons(model Point11ValCPolicyDecisionEnforcementResult) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "enforcement_result_projection_disclaimer_blocked")
	}
	if !point11ValCDecisionRefValid(model.EnforcementResultID) ||
		!point11ValCEnforcementRefValid(model.EnforcementID) ||
		!point11ValCDecisionRefValid(model.DecisionID) {
		reasons = append(reasons, "enforcement_result_identity_invalid")
	}
	if !point11ValAPolicyRefValid(model.PolicyBasisRef) {
		reasons = append(reasons, "enforcement_result_policy_basis_ref_invalid")
	}
	if len(model.ClaimRefs) > 0 && !point11ValBClaimRefListValid(model.ClaimRefs) {
		reasons = append(reasons, "enforcement_result_claim_refs_invalid")
	}
	if len(model.ClaimVerificationRefs) == 0 || !point11ValBClaimRefListValid(model.ClaimVerificationRefs) {
		reasons = append(reasons, "enforcement_result_claim_verification_refs_invalid")
	}
	if !point11ValCABACRefValid(model.ABACDecisionRef) {
		reasons = append(reasons, "enforcement_result_abac_decision_ref_invalid")
	}
	if !point11ValCDecisionRefValid(model.ExceptionDecisionRef) {
		reasons = append(reasons, "enforcement_result_exception_decision_ref_invalid")
	}
	if !point11ValCDecisionRefValid(model.EmergencyDecisionRef) {
		reasons = append(reasons, "enforcement_result_emergency_decision_ref_invalid")
	}
	if !point11Val0IdentityValueValid(model.EffectivePolicyVersion) || !point11Val0AllValuesValid(model.EffectiveClaimVersions) {
		reasons = append(reasons, "enforcement_result_effective_versions_invalid")
	}
	if !point11Val0EvidenceRefsValid(model.EvidenceRefs) {
		reasons = append(reasons, "enforcement_result_evidence_refs_invalid")
	}
	if len(model.EvidenceHashRefs) > 0 && !point11ValBEvidenceHashRefsValid(model.EvidenceHashRefs) {
		reasons = append(reasons, "enforcement_result_evidence_hash_refs_invalid")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "enforcement_result_audit_id_invalid")
	}
	if point11ValCContainsForbiddenText(model.EnforcementOutcome, model.AllowedAction, model.BlockedActionReason, model.ReviewRequiredReason) ||
		point11ValCContainsForbiddenText(model.Diagnostics...) {
		reasons = append(reasons, "enforcement_result_overclaim_detected")
	}
	if model.PolicyResultState != point11ValCPolicyStateActive {
		reasons = append(reasons, "enforcement_result_policy_invalid")
	}
	if model.ClaimResultState != point11ValCClaimStateActive {
		reasons = append(reasons, "enforcement_result_claim_invalid")
	}
	if model.ABACDecisionState != Point11ValCABACDecisionStateActive {
		reasons = append(reasons, "enforcement_result_abac_blocked")
	}
	if model.ExceptionDecisionState == Point11ValCExceptionDecisionStateBlocked {
		reasons = append(reasons, "enforcement_result_exception_invalid")
	}
	if model.EmergencyDecisionState == Point11ValCExceptionDecisionStateBlocked {
		reasons = append(reasons, "enforcement_result_emergency_invalid")
	}
	if model.EvidenceState != point11ValCEvidenceStateActive {
		reasons = append(reasons, "enforcement_result_evidence_invalid")
	}
	if model.ScopeState != point11ValCScopeStateActive {
		reasons = append(reasons, "enforcement_result_scope_invalid")
	}
	if model.AuthorityState != point11ValCAuthorityStateActive {
		reasons = append(reasons, "enforcement_result_authority_invalid")
	}
	if point11ValCRequestedActionImpliesRealSideEffect(model.AllowedAction) {
		reasons = append(reasons, "enforcement_result_allowed_action_real_side_effect")
	}
	if !point11ValCRequestedActionSupported(model.AllowedAction) {
		reasons = append(reasons, "enforcement_result_allowed_action_invalid")
	}
	if len(reasons) > 0 {
		if model.BlockedActionReason == "" || model.BlockedActionReason != strings.TrimSpace(model.BlockedActionReason) || strings.ContainsAny(model.BlockedActionReason, "\t\r\n") {
			reasons = append(reasons, "enforcement_result_blocked_action_reason_missing")
		}
		return Point11ValCEnforcementResultStateBlocked, reasons
	}
	if model.ExceptionDecisionState == Point11ValCExceptionDecisionStateReviewRequired ||
		model.EmergencyDecisionState == Point11ValCExceptionDecisionStateReviewRequired ||
		model.EnforcementState == Point11ValCEnforcementResultStateReviewRequired {
		if model.ReviewRequiredReason == "" || model.ReviewRequiredReason != strings.TrimSpace(model.ReviewRequiredReason) || strings.ContainsAny(model.ReviewRequiredReason, "\t\r\n") {
			return Point11ValCEnforcementResultStateBlocked, []string{"enforcement_result_review_required_reason_missing"}
		}
		reviewReasons := append([]string{"enforcement_result_review_required"}, model.Diagnostics...)
		return Point11ValCEnforcementResultStateReviewRequired, reviewReasons
	}
	if model.EnforcementOutcome != point11ValCRequestedOutcomeAllow {
		return Point11ValCEnforcementResultStateBlocked, []string{"enforcement_result_allow_outcome_missing"}
	}
	return Point11ValCEnforcementResultStateActive, append([]string{}, model.Diagnostics...)
}

func EvaluatePoint11ValCEnforcementResultState(model Point11ValCPolicyDecisionEnforcementResult) string {
	state, _ := point11ValCEnforcementResultStateAndReasons(model)
	return state
}

func point11ValCABACDecisionStateAndReasons(model Point11ValCABACEnforcementDecision) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "abac_projection_disclaimer_blocked")
	}
	if !point11ValCABACRefValid(model.ABACDecisionID) {
		reasons = append(reasons, "abac_decision_id_invalid")
	}
	if !point11ValBSubjectRefValid(model.SubjectRef) || !point11Val0AllValuesValid(model.SubjectAttributes) {
		reasons = append(reasons, "abac_subject_invalid")
	}
	if !point11ValCActorRefValid(model.ActorRef) || !point11Val0AllValuesValid(model.ActorAttributes) {
		reasons = append(reasons, "abac_actor_invalid")
	}
	if !point11Val0ScopeValid(model.TenantScope) {
		reasons = append(reasons, "abac_tenant_scope_invalid")
	}
	if !point11Val0AllValuesValid(model.EnvironmentAttributes) || !point11Val0AllValuesValid(model.ArtifactAttributes) {
		reasons = append(reasons, "abac_environment_or_artifact_attributes_invalid")
	}
	if !point11ValAPolicyRefValid(model.PolicyProfileRef) {
		reasons = append(reasons, "abac_policy_profile_invalid")
	}
	if len(model.ClaimRefs) > 0 && !point11ValBClaimRefListValid(model.ClaimRefs) {
		reasons = append(reasons, "abac_claim_refs_invalid")
	}
	if len(model.ExceptionRefs) > 0 && !point11ValCExceptionRefListValid(model.ExceptionRefs) {
		reasons = append(reasons, "abac_exception_refs_invalid")
	}
	if point11ValCRequestedActionImpliesRealSideEffect(model.RequestedAction) {
		reasons = append(reasons, "abac_requested_action_real_side_effect")
	} else if !point11ValCRequestedActionSupported(model.RequestedAction) {
		reasons = append(reasons, "abac_requested_action_unsupported")
	}
	if !point11ValCRequestedSurfaceAllowed(model.RequestedSurface) {
		reasons = append(reasons, "abac_requested_surface_invalid")
	}
	if !point11Val0AllValuesValid(model.AllowedAttributes) {
		reasons = append(reasons, "abac_allowed_attributes_invalid")
	}
	if len(model.DeniedAttributes) > 0 && !point11Val0AllValuesValid(model.DeniedAttributes) {
		reasons = append(reasons, "abac_denied_attributes_invalid")
	}
	if len(model.UnknownAttributes) > 0 {
		reasons = append(reasons, "abac_unknown_attributes_present")
	}
	if !point11ValCPrecedenceRuleAllowed(model.PrecedenceRule) {
		reasons = append(reasons, "abac_precedence_rule_invalid")
	}
	if !strings.Contains(point11Val0NormalizeText(model.Explanation+" "+strings.Join(model.Diagnostics, " ")), point11Val0NormalizeText(point11ValCABACExplanationVisible)) {
		reasons = append(reasons, "abac_deny_over_allow_explanation_missing")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "abac_audit_id_invalid")
	}
	if model.PolicyState != point11ValCPolicyStateActive ||
		model.ClaimState != point11ValCClaimStateActive ||
		model.EvidenceState != point11ValCEvidenceStateActive {
		reasons = append(reasons, "abac_cannot_override_invalid_policy_claim_or_evidence")
	}
	if len(model.DeniedAttributes) > 0 {
		reasons = append(reasons, "abac_deny_overrides_allow")
		if len(model.ExceptionRefs) > 0 &&
			(model.ExceptionState != Point11ValCExceptionDecisionStateActive ||
				!model.ExceptionScoped || model.ExceptionExpired || !model.ExceptionRevocable || !model.ExceptionGovernanceApproved) {
			reasons = append(reasons, "abac_exception_override_not_eligible")
		}
	}
	if point11ValCContainsForbiddenText(model.Explanation, strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "abac_overclaim_detected")
	}
	if len(reasons) > 0 {
		return Point11ValCABACDecisionStateBlocked, reasons
	}
	return Point11ValCABACDecisionStateActive, nil
}

func EvaluatePoint11ValCABACDecisionState(model Point11ValCABACEnforcementDecision) string {
	state, _ := point11ValCABACDecisionStateAndReasons(model)
	return state
}

func point11ValCExceptionDecisionStateAndReasons(model Point11ValCExceptionEmergencyDecision) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "exception_projection_disclaimer_blocked")
	}
	if !point11ValCDecisionRefValid(model.ExceptionDecisionID) {
		reasons = append(reasons, "exception_decision_id_invalid")
	}
	if !point11ValCExceptionRefValid(model.ExceptionRef) {
		reasons = append(reasons, "exception_ref_invalid")
	}
	if !point11ValCEmergencyRefValid(model.EmergencyRef) {
		reasons = append(reasons, "emergency_ref_invalid")
	}
	if !point11ValCExceptionTypeAllowed(model.ExceptionType) {
		reasons = append(reasons, "exception_type_invalid")
	}
	if !point11ValBSubjectRefValid(model.SubjectRef) {
		reasons = append(reasons, "exception_subject_ref_invalid")
	}
	if !point11Val0ScopeValid(model.TenantScope) {
		reasons = append(reasons, "exception_tenant_scope_invalid")
	}
	if !point11Val0IdentityValueValid(model.Reason) || point11ValCContainsForbiddenText(model.Reason) {
		reasons = append(reasons, "exception_reason_invalid")
	}
	if !point11ValBIssuerRefValid(model.IssuerRef) || !point11ValCActorRefValid(model.ApproverRef) {
		reasons = append(reasons, "exception_issuer_or_approver_invalid")
	}
	if model.IssuerRef == model.ApproverRef {
		reasons = append(reasons, "exception_issuer_equals_approver")
	}
	if !point11ValCAuthorityBasisRefValid(model.AuthorityBasisRef) {
		reasons = append(reasons, "exception_authority_basis_invalid")
	}
	if !point11ValBGovernanceEventRefValid(model.GovernanceEventRef) {
		reasons = append(reasons, "exception_governance_event_invalid")
	}
	if !point11Val0ValidTimestamp(model.IssuedAt) || !point11Val0ValidTimestamp(model.ExpiresAt) {
		reasons = append(reasons, "exception_timestamps_invalid")
	} else {
		expiresAt, _ := time.Parse(time.RFC3339, model.ExpiresAt)
		if expiresAt.Before(time.Now().UTC()) {
			reasons = append(reasons, "exception_expired")
		}
	}
	if !point11ValCRevocationPathRefValid(model.RevocationPathRef) {
		reasons = append(reasons, "exception_revocation_path_missing")
	}
	if !point11ValCMonitoringRefValid(model.MonitoringRequirementRef) {
		reasons = append(reasons, "exception_monitoring_requirement_missing")
	}
	if !point11ValCRollbackReviewRefValid(model.RollbackOrReviewConditionRef) {
		reasons = append(reasons, "exception_rollback_or_review_condition_missing")
	}
	if !point11Val0EvidenceRefsValid(model.EvidenceRefs) {
		reasons = append(reasons, "exception_evidence_refs_invalid")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "exception_audit_id_invalid")
	}
	if model.PermanentSilentException {
		reasons = append(reasons, "exception_permanent_silent_override_blocked")
	}
	if point11ValBClaimLifecycleInvalidated(model.EmergencyClaimState) {
		reasons = append(reasons, "exception_emergency_claim_invalidated")
	}
	if model.CreatesProductionApproval {
		reasons = append(reasons, "exception_creates_production_approval")
	}
	if model.MutatesCanonicalEvidenceSpine {
		reasons = append(reasons, "exception_mutates_canonical_evidence_spine")
	}
	if point11ValCContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "exception_diagnostics_overclaim")
	}
	if len(reasons) > 0 {
		return Point11ValCExceptionDecisionStateBlocked, reasons
	}
	switch model.ExceptionType {
	case point11ValCExceptionTypeTemporaryOverrideCandidate, point11ValCExceptionTypeEmergencyOverrideCandidate:
		return Point11ValCExceptionDecisionStateReviewRequired, []string{"exception_temporary_override_candidate"}
	default:
		return Point11ValCExceptionDecisionStateActive, nil
	}
}

func EvaluatePoint11ValCExceptionDecisionState(model Point11ValCExceptionEmergencyDecision) string {
	state, _ := point11ValCExceptionDecisionStateAndReasons(model)
	return state
}

func point11ValCPrecedenceStateAndReasons(model Point11ValCOverridePrecedence) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "precedence_projection_disclaimer_blocked")
	}
	if !point11ValCDecisionRefValid(model.PrecedenceID) ||
		!point11ValCDecisionRefValid(model.BaseDecisionRef) ||
		!point11ValCABACRefValid(model.ABACDecisionRef) ||
		!point11ValCDecisionRefValid(model.ExceptionDecisionRef) ||
		!point11ValCDecisionRefValid(model.EmergencyDecisionRef) {
		reasons = append(reasons, "precedence_identity_invalid")
	}
	if len(model.ClaimVerificationRefs) == 0 || !point11ValBClaimRefListValid(model.ClaimVerificationRefs) {
		reasons = append(reasons, "precedence_claim_verification_refs_invalid")
	}
	if !point11ValCPrecedenceRuleAllowed(model.FinalPrecedenceRule) {
		reasons = append(reasons, "precedence_rule_invalid")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "precedence_audit_id_invalid")
	}
	if point11ValCContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "precedence_overclaim_detected")
	}
	if model.LocalPolicyResult == point11ValCLocalPolicyResultInvalid {
		reasons = append(reasons, "precedence_invalid_local_policy")
	}
	if model.ClaimResultState != point11ValCClaimStateActive {
		reasons = append(reasons, "precedence_invalid_claim")
	}
	if model.ExceptionExpired {
		reasons = append(reasons, "precedence_exception_expired")
	}
	if model.ExceptionRevoked {
		reasons = append(reasons, "precedence_exception_revoked")
	}
	if model.RemoteClaimOverridesLocal {
		reasons = append(reasons, "precedence_remote_claim_overrides_local_policy")
	}
	if model.ABACResult == point11ValCABACResultDeny {
		if model.LocalPolicyResult == point11ValCLocalPolicyResultDeny {
			if point11ValCExceptionOverrideEligible(model) &&
				(model.ExceptionResult == Point11ValCExceptionDecisionStateReviewRequired ||
					model.ExceptionResult == Point11ValCExceptionDecisionStateActive) {
				return Point11ValCPrecedenceStateReviewRequired, []string{"precedence_exception_override_candidate"}
			}
			if point11ValCEmergencyOverrideEligible(model) &&
				model.EmergencyResult == Point11ValCExceptionDecisionStateReviewRequired {
				return Point11ValCPrecedenceStateReviewRequired, []string{"precedence_emergency_override_candidate"}
			}
		}
		reasons = append(reasons, "precedence_abac_deny_overrides_allow")
	}
	if model.LocalPolicyResult == point11ValCLocalPolicyResultDeny && !point11ValCExceptionOverrideEligible(model) {
		reasons = append(reasons, "precedence_local_policy_deny")
	}
	if len(reasons) > 0 {
		return Point11ValCPrecedenceStateBlocked, reasons
	}
	if model.LocalPolicyResult == point11ValCLocalPolicyResultReview ||
		model.RemoteClaimResult == point11ValCRemoteClaimResultReview ||
		model.ExceptionResult == Point11ValCExceptionDecisionStateReviewRequired ||
		model.EmergencyResult == Point11ValCExceptionDecisionStateReviewRequired {
		if model.GovernanceEventResolved {
			return Point11ValCPrecedenceStateActive, nil
		}
		return Point11ValCPrecedenceStateReviewRequired, []string{"precedence_review_required"}
	}
	if model.LocalPolicyResult == point11ValCLocalPolicyResultActive &&
		model.ABACResult == point11ValCABACResultAllow &&
		model.RemoteClaimResult != point11ValCRemoteClaimResultBlocked {
		return Point11ValCPrecedenceStateActive, nil
	}
	return Point11ValCPrecedenceStateBlocked, []string{"precedence_not_active"}
}

func EvaluatePoint11ValCPrecedenceState(model Point11ValCOverridePrecedence) string {
	state, _ := point11ValCPrecedenceStateAndReasons(model)
	return state
}

func point11ValCMonitoringStateAndReasons(model Point11ValCMonitoringLinkedEmergencyHandling) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "monitoring_projection_disclaimer_blocked")
	}
	if !point11ValCMonitoringRefValid(model.MonitoringLinkID) {
		reasons = append(reasons, "monitoring_link_id_invalid")
	}
	if !point11ValCEmergencyRefValid(model.EmergencyRef) {
		reasons = append(reasons, "monitoring_emergency_ref_invalid")
	}
	if !point11ValCMonitoringRefValid(model.MonitoringRequirementRef) {
		reasons = append(reasons, "monitoring_requirement_ref_invalid")
	}
	if model.MonitoringState != point11ValCMonitoringLinkActive {
		reasons = append(reasons, "monitoring_link_not_active")
	}
	if !point11ValCSignalRefsValid(model.SignalRefs) {
		reasons = append(reasons, "monitoring_signal_refs_invalid")
	}
	if !point11Val0ContainsTrimmed([]string{point11ValCMonitoringSignalFresh, point11ValCMonitoringSignalStale}, model.SignalFreshness) {
		reasons = append(reasons, "monitoring_signal_freshness_invalid")
	}
	if !point11Val0ValidTimestamp(model.ReviewDeadline) {
		reasons = append(reasons, "monitoring_review_deadline_invalid")
	} else {
		reviewDeadline, _ := time.Parse(time.RFC3339, model.ReviewDeadline)
		if reviewDeadline.Before(time.Now().UTC()) {
			reasons = append(reasons, "monitoring_review_deadline_expired")
		}
	}
	if !point11ValCGenericStateActive(model.ExpiryEnforcementState) {
		reasons = append(reasons, "monitoring_expiry_enforcement_missing")
	}
	if !point11ValCGenericStateActive(model.RevocationCheckState) {
		reasons = append(reasons, "monitoring_revocation_check_missing")
	}
	if !point11ValCGenericStateActive(model.RollbackCheckState) {
		reasons = append(reasons, "monitoring_rollback_check_missing")
	}
	if !point11ValBAuditRefValid(model.AuditID) {
		reasons = append(reasons, "monitoring_audit_id_invalid")
	}
	if model.HighRiskEmergency && !point11ValCEscalationRefValid(model.EscalationRef) {
		reasons = append(reasons, "monitoring_escalation_ref_missing")
	}
	if model.CreatesActionSideEffects {
		reasons = append(reasons, "monitoring_creates_action_side_effects")
	}
	if point11ValCContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "monitoring_diagnostics_overclaim")
	}
	if len(reasons) > 0 {
		return Point11ValCMonitoringStateBlocked, reasons
	}
	if model.SignalFreshness == point11ValCMonitoringSignalStale {
		return Point11ValCMonitoringStateReviewRequired, []string{"monitoring_signal_stale_review_required"}
	}
	return Point11ValCMonitoringStateActive, nil
}

func EvaluatePoint11ValCMonitoringState(model Point11ValCMonitoringLinkedEmergencyHandling) string {
	state, _ := point11ValCMonitoringStateAndReasons(model)
	return state
}

func point11ValCDashboardStateAndReasons(model Point11ValCGovernanceDashboardReadModel) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "dashboard_projection_disclaimer_blocked")
	}
	if !point11ValCDashboardRefValid(model.DashboardViewID) {
		reasons = append(reasons, "dashboard_view_id_invalid")
	}
	if len(model.SourceDecisionRefs) == 0 || !point11ValCDecisionRefListValid(model.SourceDecisionRefs) {
		reasons = append(reasons, "dashboard_source_decision_refs_invalid")
	}
	if len(model.SourceClaimRefs) == 0 || !point11ValBClaimRefListValid(model.SourceClaimRefs) {
		reasons = append(reasons, "dashboard_source_claim_refs_invalid")
	}
	if len(model.SourceExceptionRefs) == 0 || !point11ValCExceptionRefListValid(model.SourceExceptionRefs) {
		reasons = append(reasons, "dashboard_source_exception_refs_invalid")
	}
	if len(model.SourceEmergencyRefs) == 0 || !point11ValCEmergencyRefListValid(model.SourceEmergencyRefs) {
		reasons = append(reasons, "dashboard_source_emergency_refs_invalid")
	}
	if len(model.SourceMonitoringRefs) == 0 || !point11ValCMonitoringRefListValid(model.SourceMonitoringRefs) {
		reasons = append(reasons, "dashboard_source_monitoring_refs_invalid")
	}
	if !point11Val0IdentityValueValid(model.RenderedState) || point11ValCContainsForbiddenText(model.RenderedState) {
		reasons = append(reasons, "dashboard_rendered_state_invalid_or_overclaim")
	}
	if !point11ValCVisibleSurfacesValid(model.VisibleSurfaces) {
		reasons = append(reasons, "dashboard_visible_surfaces_invalid")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "dashboard_publication_side_effects_blocked")
	}
	if model.CreatesAuthorityClaim {
		reasons = append(reasons, "dashboard_authority_claim_blocked")
	}
	if model.MutatesCanonicalState {
		reasons = append(reasons, "dashboard_mutates_canonical_state")
	}
	if point11ValCContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "dashboard_diagnostics_overclaim")
	}
	if len(reasons) > 0 {
		return Point11ValCDashboardStateBlocked, reasons
	}
	return Point11ValCDashboardStateActive, nil
}

func EvaluatePoint11ValCDashboardState(model Point11ValCGovernanceDashboardReadModel) string {
	state, _ := point11ValCDashboardStateAndReasons(model)
	return state
}

func point11ValCBlockingReasons(model Point11ValCFoundation) []string {
	reasons := []string{}
	if model.DependencyState != Point11ValCDependencyStateActive &&
		model.DependencyState != Point11ValCDependencyStateReviewRequired {
		reasons = append(reasons, "valb_dependency_blocked")
	}
	if model.EnforcementInputState != Point11ValCEnforcementInputStateActive {
		reasons = append(reasons, "enforcement_input_blocked")
	}
	if model.EnforcementResultState != Point11ValCEnforcementResultStateActive &&
		model.EnforcementResultState != Point11ValCEnforcementResultStateReviewRequired {
		reasons = append(reasons, "enforcement_result_blocked")
	}
	if model.ABACDecisionState != Point11ValCABACDecisionStateActive {
		reasons = append(reasons, "abac_decision_blocked")
	}
	if model.ExceptionDecisionState != Point11ValCExceptionDecisionStateActive &&
		model.ExceptionDecisionState != Point11ValCExceptionDecisionStateReviewRequired {
		reasons = append(reasons, "exception_decision_blocked")
	}
	if model.PrecedenceState != Point11ValCPrecedenceStateActive &&
		model.PrecedenceState != Point11ValCPrecedenceStateReviewRequired {
		reasons = append(reasons, "precedence_blocked")
	}
	if model.MonitoringState != Point11ValCMonitoringStateActive &&
		model.MonitoringState != Point11ValCMonitoringStateReviewRequired {
		reasons = append(reasons, "monitoring_blocked")
	}
	if model.DashboardState != Point11ValCDashboardStateActive {
		reasons = append(reasons, "dashboard_blocked")
	}
	if model.CreatesAuthorityClaims {
		reasons = append(reasons, "authority_claim_surface_blocked")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "publication_side_effects_blocked")
	}
	if model.CreatesRealEnforcementSideEffects {
		reasons = append(reasons, "real_enforcement_side_effects_blocked")
	}
	return reasons
}

func point11ValCDiagnosticsModel(
	model Point11ValCFoundation,
	dependencyReasons []string,
	enforcementInputReasons []string,
	enforcementResultReasons []string,
	abacReasons []string,
	exceptionReasons []string,
	precedenceReasons []string,
	monitoringReasons []string,
	dashboardReasons []string,
) Point11ValCDiagnostics {
	return Point11ValCDiagnostics{
		CurrentState:             model.CurrentState,
		BlockingReasons:          append([]string{}, model.BlockingReasons...),
		ReviewPrerequisites:      append([]string{}, model.ReviewPrerequisites...),
		ComponentStates:          point11ValCComponentStates(model),
		DependencyReasons:        append([]string{}, dependencyReasons...),
		EnforcementInputReasons:  append([]string{}, enforcementInputReasons...),
		EnforcementResultReasons: append([]string{}, enforcementResultReasons...),
		ABACReasons:              append([]string{}, abacReasons...),
		ExceptionReasons:         append([]string{}, exceptionReasons...),
		PrecedenceReasons:        append([]string{}, precedenceReasons...),
		MonitoringReasons:        append([]string{}, monitoringReasons...),
		DashboardReasons:         append([]string{}, dashboardReasons...),
		ProjectionDisclaimer:     model.ProjectionDisclaimer,
	}
}

func EvaluatePoint11ValCFoundationState(model Point11ValCFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.CreatesAuthorityClaims ||
		model.CreatesPublicationSideEffects ||
		model.CreatesRealEnforcementSideEffects ||
		model.DependencyState == Point11ValCDependencyStateBlocked ||
		model.EnforcementInputState == Point11ValCEnforcementInputStateBlocked ||
		model.EnforcementResultState == Point11ValCEnforcementResultStateBlocked ||
		model.ABACDecisionState == Point11ValCABACDecisionStateBlocked ||
		model.ExceptionDecisionState == Point11ValCExceptionDecisionStateBlocked ||
		model.PrecedenceState == Point11ValCPrecedenceStateBlocked ||
		model.MonitoringState == Point11ValCMonitoringStateBlocked ||
		model.DashboardState == Point11ValCDashboardStateBlocked {
		return Point11ValCStateBlocked
	}
	if model.DependencyState == Point11ValCDependencyStateReviewRequired ||
		model.ExceptionDecisionState == Point11ValCExceptionDecisionStateReviewRequired ||
		model.EnforcementResultState == Point11ValCEnforcementResultStateReviewRequired ||
		model.PrecedenceState == Point11ValCPrecedenceStateReviewRequired ||
		model.MonitoringState == Point11ValCMonitoringStateReviewRequired {
		return Point11ValCStateReviewRequired
	}
	if model.DependencyState == Point11ValCDependencyStateActive &&
		model.EnforcementInputState == Point11ValCEnforcementInputStateActive &&
		model.EnforcementResultState == Point11ValCEnforcementResultStateActive &&
		model.ABACDecisionState == Point11ValCABACDecisionStateActive &&
		model.ExceptionDecisionState == Point11ValCExceptionDecisionStateActive &&
		model.PrecedenceState == Point11ValCPrecedenceStateActive &&
		model.MonitoringState == Point11ValCMonitoringStateActive &&
		model.DashboardState == Point11ValCDashboardStateActive {
		return Point11ValCStateActive
	}
	return Point11ValCStateBlocked
}

func Point11ValCFoundationModel() Point11ValCFoundation {
	disclaimer := point11ValCProjectionDisclaimerBaseline
	claimID := "claim_point11_valb_customer_scope_001"
	policyRef := "policy_point11_vala_authority_core_v1"
	enforcementID := "enforcement_point11_valc_decision_001"
	decisionID := "decision_point11_valc_governance_001"
	return Point11ValCFoundation{
		CurrentState:                      Point11ValCStateReviewRequired,
		DependencyState:                   Point11ValCDependencyStateReviewRequired,
		EnforcementInputState:             Point11ValCEnforcementInputStateActive,
		EnforcementResultState:            Point11ValCEnforcementResultStateActive,
		ABACDecisionState:                 Point11ValCABACDecisionStateActive,
		ExceptionDecisionState:            Point11ValCExceptionDecisionStateActive,
		PrecedenceState:                   Point11ValCPrecedenceStateActive,
		MonitoringState:                   Point11ValCMonitoringStateActive,
		DashboardState:                    Point11ValCDashboardStateActive,
		ProjectionDisclaimer:              disclaimer,
		CreatesAuthorityClaims:            false,
		CreatesPublicationSideEffects:     false,
		CreatesRealEnforcementSideEffects: false,
		Dependency:                        point11ValCDependencySnapshotModel(),
		EnforcementInput: Point11ValCGovernanceEnforcementInput{
			EnforcementID:          enforcementID,
			DecisionID:             decisionID,
			SubjectRef:             "subject_point11_valb_workload_alpha",
			SubjectKind:            "workload",
			ActorRef:               "actor_point11_valc_governance_operator",
			ActorKind:              "governance_operator",
			TenantScope:            "tenant_scope_alpha",
			EnvironmentRef:         "environment_point11_valc_prod_shadow",
			ArtifactRef:            "artifact_point11_valc_policy_artifact_001",
			PolicyBasisRef:         policyRef,
			PolicyBasisState:       point11ValBPolicyBasisStateActive,
			PolicyVersion:          "point11_vala_policy_v1",
			ClaimsRequired:         true,
			ClaimRefs:              []string{claimID},
			ClaimVerificationRefs:  []string{claimID},
			ClaimVerificationState: Point11ValBVerificationStateActive,
			RegistryRef:            "claim_registry_point11_valb_core",
			RegistryState:          Point11ValBRegistryStateActive,
			ABACContextRef:         "abac_point11_valc_context_001",
			ExceptionRefs:          []string{"exception_point11_valc_scope_override_001"},
			EmergencyRefs:          []string{"emergency_point11_valc_scope_override_001"},
			RequestedAction:        point11ValCRequestedActionEvaluate,
			RequestedSurface:       point11Val0PublicationSurfaceExport,
			RequestedOutcome:       point11ValCRequestedOutcomeAllow,
			EvidenceRefs:           []string{"evidence:point11-valc-enforcement-001", "evidence:point11-valc-enforcement-002"},
			EvidenceHashRefs:       []string{"evidence_hash_point11_valc_enforcement_001"},
			GovernanceEventRef:     "governance_event_point11_valc_enforcement_001",
			AuditID:                "audit_point11_valc_enforcement_001",
			DecisionTimestamp:      "2099-01-01T00:00:00Z",
			ProjectionDisclaimer:   disclaimer,
		},
		EnforcementResult: Point11ValCPolicyDecisionEnforcementResult{
			EnforcementResultID:    "decision_point11_valc_enforcement_result_001",
			EnforcementID:          enforcementID,
			DecisionID:             decisionID,
			PolicyBasisRef:         policyRef,
			ClaimRefs:              []string{claimID},
			ClaimVerificationRefs:  []string{claimID},
			ABACDecisionRef:        "abac_point11_valc_context_001",
			ExceptionDecisionRef:   "decision_point11_valc_exception_001",
			EmergencyDecisionRef:   "decision_point11_valc_emergency_001",
			PolicyResultState:      point11ValCPolicyStateActive,
			ClaimResultState:       point11ValCClaimStateActive,
			ABACDecisionState:      Point11ValCABACDecisionStateActive,
			ExceptionDecisionState: Point11ValCExceptionDecisionStateActive,
			EmergencyDecisionState: Point11ValCExceptionDecisionStateActive,
			EvidenceState:          point11ValCEvidenceStateActive,
			ScopeState:             point11ValCScopeStateActive,
			AuthorityState:         point11ValCAuthorityStateActive,
			EnforcementState:       Point11ValCEnforcementResultStateActive,
			EnforcementOutcome:     point11ValCRequestedOutcomeAllow,
			AllowedAction:          point11ValCRequestedActionEvaluate,
			EffectivePolicyVersion: "point11_vala_policy_v1",
			EffectiveClaimVersions: []string{"claim_version_point11_valb_v1"},
			EvidenceRefs:           []string{"evidence:point11-valc-enforcement-001", "evidence:point11-valc-enforcement-002"},
			EvidenceHashRefs:       []string{"evidence_hash_point11_valc_enforcement_001"},
			AuditID:                "audit_point11_valc_enforcement_result_001",
			Diagnostics:            []string{"policy_reason_present", "claim_reason_present", "abac_reason_present", "exception_reason_present", "evidence_reason_present"},
			ProjectionDisclaimer:   disclaimer,
		},
		ABACDecision: Point11ValCABACEnforcementDecision{
			ABACDecisionID:              "abac_point11_valc_context_001",
			SubjectRef:                  "subject_point11_valb_workload_alpha",
			SubjectAttributes:           []string{"subject:workload", "subject:verified"},
			ActorRef:                    "actor_point11_valc_governance_operator",
			ActorAttributes:             []string{"actor:governance_operator", "actor:approved"},
			TenantScope:                 "tenant_scope_alpha",
			EnvironmentAttributes:       []string{"env:shadow", "env:bounded"},
			ArtifactAttributes:          []string{"artifact:policy_bound", "artifact:evidence_linked"},
			PolicyProfileRef:            policyRef,
			ClaimRefs:                   []string{claimID},
			ExceptionRefs:               []string{"exception_point11_valc_scope_override_001"},
			RequestedAction:             point11ValCRequestedActionEvaluate,
			RequestedSurface:            point11Val0PublicationSurfaceExport,
			AllowedAttributes:           []string{"allow:customer_visible", "allow:governance_bound"},
			DeniedAttributes:            nil,
			UnknownAttributes:           nil,
			PrecedenceRule:              point11ValCABACPrecedenceDenyOverAllow,
			DecisionState:               Point11ValCABACDecisionStateActive,
			Explanation:                 point11ValCABACExplanationVisible,
			Diagnostics:                 []string{point11ValCABACExplanationVisible},
			AuditID:                     "audit_point11_valc_abac_001",
			ProjectionDisclaimer:        disclaimer,
			PolicyState:                 point11ValCPolicyStateActive,
			ClaimState:                  point11ValCClaimStateActive,
			EvidenceState:               point11ValCEvidenceStateActive,
			ExceptionState:              Point11ValCExceptionDecisionStateActive,
			ExceptionScoped:             true,
			ExceptionExpired:            false,
			ExceptionRevocable:          true,
			ExceptionGovernanceApproved: true,
		},
		ExceptionDecision: Point11ValCExceptionEmergencyDecision{
			ExceptionDecisionID:          "decision_point11_valc_exception_001",
			ExceptionRef:                 "exception_point11_valc_scope_override_001",
			EmergencyRef:                 "emergency_point11_valc_scope_override_001",
			ExceptionType:                point11ValCExceptionTypeScoped,
			SubjectRef:                   "subject_point11_valb_workload_alpha",
			TenantScope:                  "tenant_scope_alpha",
			Reason:                       "bounded tenant scoped governance exception",
			IssuerRef:                    "issuer_point11_valb_governance_team",
			ApproverRef:                  "actor_point11_valc_exception_approver",
			AuthorityBasisRef:            "authority_basis_point11_valc_policy_exception_001",
			GovernanceEventRef:           "governance_event_point11_valc_exception_001",
			IssuedAt:                     "2098-01-01T00:00:00Z",
			ExpiresAt:                    "2099-01-01T00:00:00Z",
			RevocationPathRef:            "revocation_path_point11_valc_exception_001",
			MonitoringRequirementRef:     "monitoring_point11_valc_requirement_001",
			RollbackOrReviewConditionRef: "rollback_review_point11_valc_exception_001",
			EvidenceRefs:                 []string{"evidence:point11-valc-exception-001"},
			AuditID:                      "audit_point11_valc_exception_001",
			DecisionState:                Point11ValCExceptionDecisionStateActive,
			Diagnostics:                  []string{"exception_scoped_and_revocable"},
			ProjectionDisclaimer:         disclaimer,
			EmergencyClaimState:          point11ValBClaimLifecycleActive,
		},
		Precedence: Point11ValCOverridePrecedence{
			PrecedenceID:                "decision_point11_valc_precedence_001",
			BaseDecisionRef:             decisionID,
			ABACDecisionRef:             "abac_point11_valc_context_001",
			ExceptionDecisionRef:        "decision_point11_valc_exception_001",
			EmergencyDecisionRef:        "decision_point11_valc_emergency_001",
			ClaimVerificationRefs:       []string{claimID},
			LocalPolicyResult:           point11ValCLocalPolicyResultActive,
			ClaimResultState:            point11ValCClaimStateActive,
			ABACResult:                  point11ValCABACResultAllow,
			RemoteClaimResult:           point11ValCRemoteClaimResultCompatible,
			ExceptionResult:             Point11ValCExceptionDecisionStateActive,
			EmergencyResult:             Point11ValCExceptionDecisionStateActive,
			FinalPrecedenceRule:         "blocked_over_review_required_over_active",
			FinalState:                  Point11ValCPrecedenceStateActive,
			Diagnostics:                 []string{"precedence_policy_reason_present", "precedence_claim_reason_present", "precedence_abac_reason_present"},
			AuditID:                     "audit_point11_valc_precedence_001",
			ProjectionDisclaimer:        disclaimer,
			ExceptionScoped:             true,
			ExceptionExpired:            false,
			ExceptionRevoked:            false,
			ExceptionGovernanceApproved: true,
			EmergencyScoped:             true,
			EmergencyTimeBound:          true,
			EmergencyMonitored:          true,
			EmergencyRevocable:          true,
			GovernanceEventResolved:     true,
		},
		Monitoring: Point11ValCMonitoringLinkedEmergencyHandling{
			MonitoringLinkID:         "monitoring_point11_valc_link_001",
			EmergencyRef:             "emergency_point11_valc_scope_override_001",
			MonitoringRequirementRef: "monitoring_point11_valc_requirement_001",
			MonitoringState:          point11ValCMonitoringLinkActive,
			SignalRefs:               []string{"signal_point11_valc_health_001"},
			SignalFreshness:          point11ValCMonitoringSignalFresh,
			EscalationRef:            "escalation_point11_valc_emergency_001",
			ReviewDeadline:           "2099-01-01T00:00:00Z",
			ExpiryEnforcementState:   point11ValCCheckStateActive,
			RevocationCheckState:     point11ValCCheckStateActive,
			RollbackCheckState:       point11ValCCheckStateActive,
			AuditID:                  "audit_point11_valc_monitoring_001",
			Diagnostics:              []string{"monitoring_projection_only"},
			ProjectionDisclaimer:     disclaimer,
			HighRiskEmergency:        true,
		},
		Dashboard: Point11ValCGovernanceDashboardReadModel{
			DashboardViewID:               "dashboard_point11_valc_governance_001",
			SourceDecisionRefs:            []string{decisionID},
			SourceClaimRefs:               []string{claimID},
			SourceExceptionRefs:           []string{"exception_point11_valc_scope_override_001"},
			SourceEmergencyRefs:           []string{"emergency_point11_valc_scope_override_001"},
			SourceMonitoringRefs:          []string{"monitoring_point11_valc_link_001"},
			RenderedState:                 point11ValCDashboardRenderedStateBounded,
			VisibleSurfaces:               []string{point11ValCRequestedSurfaceGovernanceDashboard, point11ValCRequestedSurfaceInternalReview},
			ProjectionDisclaimer:          disclaimer,
			CreatesPublicationSideEffects: false,
			CreatesAuthorityClaim:         false,
			MutatesCanonicalState:         false,
			Diagnostics:                   []string{"dashboard_projection_only"},
		},
	}
}

func ComputePoint11ValCFoundation(model Point11ValCFoundation) Point11ValCFoundation {
	dependencyState, dependencyReasons := point11ValCDependencyStateAndReasons(model.Dependency)
	enforcementInputState, enforcementInputReasons := point11ValCEnforcementInputStateAndReasons(model.EnforcementInput)
	enforcementResultState, enforcementResultReasons := point11ValCEnforcementResultStateAndReasons(model.EnforcementResult)
	abacState, abacReasons := point11ValCABACDecisionStateAndReasons(model.ABACDecision)
	exceptionState, exceptionReasons := point11ValCExceptionDecisionStateAndReasons(model.ExceptionDecision)
	precedenceState, precedenceReasons := point11ValCPrecedenceStateAndReasons(model.Precedence)
	monitoringState, monitoringReasons := point11ValCMonitoringStateAndReasons(model.Monitoring)
	dashboardState, dashboardReasons := point11ValCDashboardStateAndReasons(model.Dashboard)

	model.DependencyState = dependencyState
	model.EnforcementInputState = enforcementInputState
	model.EnforcementResultState = enforcementResultState
	model.ABACDecisionState = abacState
	model.ExceptionDecisionState = exceptionState
	model.PrecedenceState = precedenceState
	model.MonitoringState = monitoringState
	model.DashboardState = dashboardState
	model.CurrentState = EvaluatePoint11ValCFoundationState(model)
	model.BlockingReasons = point11ValCBlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	if model.ExceptionDecisionState == Point11ValCExceptionDecisionStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "exception_decision_review_required")
	}
	if model.MonitoringState == Point11ValCMonitoringStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "monitoring_review_required")
	}
	model.Diagnostics = point11ValCDiagnosticsModel(
		model,
		dependencyReasons,
		enforcementInputReasons,
		enforcementResultReasons,
		abacReasons,
		exceptionReasons,
		precedenceReasons,
		monitoringReasons,
		dashboardReasons,
	)
	return model
}
