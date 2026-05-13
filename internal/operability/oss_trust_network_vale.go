package operability

import (
	"strings"
	"unicode"
)

const (
	OSSTrustNetworkPoint9StatePass = "oss_trust_network_point_9_pass"

	OSSTrustNetworkValEStatePass       = "oss_trust_network_vale_pass"
	OSSTrustNetworkValEStatePartial    = "oss_trust_network_vale_partial"
	OSSTrustNetworkValEStateIncomplete = "oss_trust_network_vale_incomplete"
	OSSTrustNetworkValEStateBlocked    = "oss_trust_network_vale_blocked"
	OSSTrustNetworkValEStateUnknown    = "oss_trust_network_vale_unknown"

	OSSTrustNetworkValESourceStateActive     = "oss_trust_network_vale_source_active"
	OSSTrustNetworkValESourceStatePartial    = "oss_trust_network_vale_source_partial"
	OSSTrustNetworkValESourceStateIncomplete = "oss_trust_network_vale_source_incomplete"
	OSSTrustNetworkValESourceStateBlocked    = "oss_trust_network_vale_source_blocked"
	OSSTrustNetworkValESourceStateUnknown    = "oss_trust_network_vale_source_unknown"

	OSSTrustNetworkValEDependencyStateActive     = "oss_trust_network_vale_dependency_active"
	OSSTrustNetworkValEDependencyStatePartial    = "oss_trust_network_vale_dependency_partial"
	OSSTrustNetworkValEDependencyStateIncomplete = "oss_trust_network_vale_dependency_incomplete"
	OSSTrustNetworkValEDependencyStateBlocked    = "oss_trust_network_vale_dependency_blocked"
	OSSTrustNetworkValEDependencyStateUnknown    = "oss_trust_network_vale_dependency_unknown"

	OSSTrustNetworkValEIntegratedClosureStateActive     = "oss_trust_network_vale_integrated_closure_active"
	OSSTrustNetworkValEIntegratedClosureStatePartial    = "oss_trust_network_vale_integrated_closure_partial"
	OSSTrustNetworkValEIntegratedClosureStateIncomplete = "oss_trust_network_vale_integrated_closure_incomplete"
	OSSTrustNetworkValEIntegratedClosureStateBlocked    = "oss_trust_network_vale_integrated_closure_blocked"
	OSSTrustNetworkValEIntegratedClosureStateUnknown    = "oss_trust_network_vale_integrated_closure_unknown"

	OSSTrustNetworkValECanonicalBoundaryStateActive     = "oss_trust_network_vale_canonical_boundary_active"
	OSSTrustNetworkValECanonicalBoundaryStatePartial    = "oss_trust_network_vale_canonical_boundary_partial"
	OSSTrustNetworkValECanonicalBoundaryStateIncomplete = "oss_trust_network_vale_canonical_boundary_incomplete"
	OSSTrustNetworkValECanonicalBoundaryStateBlocked    = "oss_trust_network_vale_canonical_boundary_blocked"
	OSSTrustNetworkValECanonicalBoundaryStateUnknown    = "oss_trust_network_vale_canonical_boundary_unknown"

	OSSTrustNetworkValEEvidenceQualityStateActive     = "oss_trust_network_vale_evidence_quality_active"
	OSSTrustNetworkValEEvidenceQualityStatePartial    = "oss_trust_network_vale_evidence_quality_partial"
	OSSTrustNetworkValEEvidenceQualityStateIncomplete = "oss_trust_network_vale_evidence_quality_incomplete"
	OSSTrustNetworkValEEvidenceQualityStateBlocked    = "oss_trust_network_vale_evidence_quality_blocked"
	OSSTrustNetworkValEEvidenceQualityStateUnknown    = "oss_trust_network_vale_evidence_quality_unknown"

	OSSTrustNetworkValENoOverclaimStateActive     = "oss_trust_network_vale_no_overclaim_active"
	OSSTrustNetworkValENoOverclaimStatePartial    = "oss_trust_network_vale_no_overclaim_partial"
	OSSTrustNetworkValENoOverclaimStateIncomplete = "oss_trust_network_vale_no_overclaim_incomplete"
	OSSTrustNetworkValENoOverclaimStateBlocked    = "oss_trust_network_vale_no_overclaim_blocked"
	OSSTrustNetworkValENoOverclaimStateUnknown    = "oss_trust_network_vale_no_overclaim_unknown"

	OSSTrustNetworkValEFinalPassRuleStateActive     = "oss_trust_network_vale_final_pass_rule_active"
	OSSTrustNetworkValEFinalPassRuleStatePartial    = "oss_trust_network_vale_final_pass_rule_partial"
	OSSTrustNetworkValEFinalPassRuleStateIncomplete = "oss_trust_network_vale_final_pass_rule_incomplete"
	OSSTrustNetworkValEFinalPassRuleStateBlocked    = "oss_trust_network_vale_final_pass_rule_blocked"
	OSSTrustNetworkValEFinalPassRuleStateUnknown    = "oss_trust_network_vale_final_pass_rule_unknown"

	OSSTrustNetworkValEClosureStateActive     = "oss_trust_network_vale_closure_active"
	OSSTrustNetworkValEClosureStatePartial    = "oss_trust_network_vale_closure_partial"
	OSSTrustNetworkValEClosureStateIncomplete = "oss_trust_network_vale_closure_incomplete"
	OSSTrustNetworkValEClosureStateBlocked    = "oss_trust_network_vale_closure_blocked"
	OSSTrustNetworkValEClosureStateUnknown    = "oss_trust_network_vale_closure_unknown"

	OSSTrustNetworkValEPoint9PassReasonAllowed                  = "point_9_pass through Val E only after actual Val 0 through Val D proof states, exact proof surfaces, exact evidence refs, fresh evidence quality, bounded canonical boundaries, fail-closed closure gates, and no-overclaim discipline all remain active."
	OSSTrustNetworkValEPoint9PassReasonBlocked                  = "point_9_pass remains blocked until actual Val 0 through Val D proof states, exact proof surfaces, exact evidence refs, fresh evidence quality, bounded canonical boundaries, fail-closed closure gates, and no-overclaim discipline all remain active."
	OSSTrustNetworkValEPoint9PassSafeDiagnosticVal0CannotReturn = "Val 0 cannot return point_9_pass before integrated closure."
	OSSTrustNetworkValEPoint9PassSafeDiagnosticValACannotReturn = "Val A cannot return point_9_pass before integrated closure."
	OSSTrustNetworkValEPoint9PassSafeDiagnosticValBCannotReturn = "Val B cannot return point_9_pass before integrated closure."
	OSSTrustNetworkValEPoint9PassSafeDiagnosticValCCannotReturn = "Val C cannot return point_9_pass before integrated closure."
	OSSTrustNetworkValEPoint9PassSafeDiagnosticValDCannotReturn = "Val D cannot return point_9_pass before integrated closure."
)

const (
	ossTrustNetworkValEPoint9PassReasonStateAllowed = "allowed"
	ossTrustNetworkValEPoint9PassReasonStateBlocked = "blocked"
	ossTrustNetworkValEPoint9PassReasonStateUnknown = "unknown"
)

type OSSTrustNetworkValEIntegratedClosureGate struct {
	CurrentState                                 string   `json:"current_state"`
	GateID                                       string   `json:"gate_id"`
	Version                                      string   `json:"version"`
	CanonicalExecutionAuditEvidenceSourceOfTruth bool     `json:"canonical_execution_audit_evidence_source_of_truth"`
	OSSNetworkOutputsRemainAdvisory              bool     `json:"oss_network_outputs_remain_advisory"`
	EvidenceLinkedIntegratedClosure              bool     `json:"evidence_linked_integrated_closure"`
	LocalEnterpriseApplicabilityExplicit         bool     `json:"local_enterprise_applicability_explicit"`
	EvidenceRefs                                 []string `json:"evidence_refs,omitempty"`
	Caveats                                      []string `json:"caveats,omitempty"`
	ProjectionDisclaimer                         string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValECanonicalBoundaryGate struct {
	CurrentState                         string   `json:"current_state"`
	BoundaryID                           string   `json:"boundary_id"`
	LocalOverrideVisibleEvidenceLinked   bool     `json:"local_override_visible_evidence_linked"`
	LocalOverrideVisible                 bool     `json:"local_override_visible"`
	LocalEnterpriseApplicabilityExplicit bool     `json:"local_enterprise_applicability_explicit"`
	RegistrySurfaceAuthorityClaim        bool     `json:"registry_surface_authority_claim"`
	CommunitySurfaceAuthorityClaim       bool     `json:"community_surface_authority_claim"`
	MaintainerSurfaceAuthorityClaim      bool     `json:"maintainer_surface_authority_claim"`
	SharedVEXSurfaceAuthorityClaim       bool     `json:"shared_vex_surface_authority_claim"`
	DashboardSurfaceAuthorityClaim       bool     `json:"dashboard_surface_authority_claim"`
	RemediationSurfaceAuthorityClaim     bool     `json:"remediation_surface_authority_claim"`
	PRProposalSurfaceAuthorityClaim      bool     `json:"pr_proposal_surface_authority_claim"`
	PropagationSurfaceAuthorityClaim     bool     `json:"propagation_surface_authority_claim"`
	OfficialOSSAuthorityClaim            bool     `json:"official_oss_authority_claim"`
	SharedSignalSilentOverride           bool     `json:"shared_signal_silent_override"`
	FinalClosureApprovesDeployment       bool     `json:"final_closure_approves_deployment"`
	FinalClosureApprovesProductionUse    bool     `json:"final_closure_approves_production_use"`
	EvidenceRefs                         []string `json:"evidence_refs,omitempty"`
	Caveats                              []string `json:"caveats,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValEEvidenceQuality struct {
	CurrentState                   string                                   `json:"current_state"`
	EvidenceQualityID              string                                   `json:"evidence_quality_id"`
	Evidence                       []ReferenceArchitectureEvidenceReference `json:"evidence,omitempty"`
	ProofSurfaceRefs               []string                                 `json:"proof_surface_refs,omitempty"`
	EvidenceRefs                   []string                                 `json:"evidence_refs,omitempty"`
	DependencyProofSurfaceRefs     []string                                 `json:"dependency_proof_surface_refs,omitempty"`
	DependencyEvidenceRefs         []string                                 `json:"dependency_evidence_refs,omitempty"`
	DependencyEvidence             []ReferenceArchitectureEvidenceReference `json:"dependency_evidence,omitempty"`
	DependencyProjectionDisclaimer string                                   `json:"dependency_projection_disclaimer"`
	ProjectionDisclaimer           string                                   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValENoOverclaim struct {
	CurrentState              string   `json:"current_state"`
	GateID                    string   `json:"gate_id"`
	Version                   string   `json:"version"`
	ObservedClaims            []string `json:"observed_claims,omitempty"`
	Point9PassClaim           bool     `json:"point_9_pass_claim"`
	CertifiesTrust            bool     `json:"certifies_trust"`
	ApprovesProductionUse     bool     `json:"approves_production_use"`
	ApprovesDeployment        bool     `json:"approves_deployment"`
	RegulatorApprovalClaim    bool     `json:"regulator_approval_claim"`
	LegalIPClearanceClaim     bool     `json:"legal_ip_clearance_claim"`
	PublicBadgeClaim          bool     `json:"public_badge_claim"`
	OfficialOSSAuthorityClaim bool     `json:"official_oss_authority_claim"`
	GlobalTruthClaim          bool     `json:"global_truth_claim"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValEIntegratedClosure struct {
	CurrentState             string                                   `json:"current_state"`
	Point9State              string                                   `json:"point_9_state"`
	Point9PassAllowed        bool                                     `json:"point_9_pass_allowed"`
	Point9PassReason         string                                   `json:"point_9_pass_reason"`
	ClosureState             string                                   `json:"closure_state"`
	DependencyState          string                                   `json:"dependency_state"`
	Val0SourceState          string                                   `json:"val0_source_state"`
	ValASourceState          string                                   `json:"vala_source_state"`
	ValBSourceState          string                                   `json:"valb_source_state"`
	ValCSourceState          string                                   `json:"valc_source_state"`
	ValDSourceState          string                                   `json:"vald_source_state"`
	IntegratedClosureState   string                                   `json:"integrated_closure_state"`
	CanonicalBoundaryState   string                                   `json:"canonical_boundary_state"`
	EvidenceQualityState     string                                   `json:"evidence_quality_state"`
	NoOverclaimState         string                                   `json:"no_overclaim_state"`
	FinalPassRuleState       string                                   `json:"final_pass_rule_state"`
	Val0Source               OSSTrustNetworkVal0Foundation            `json:"val0_source"`
	ValASource               OSSTrustNetworkValACore                  `json:"vala_source"`
	ValBSource               OSSTrustNetworkValBCore                  `json:"valb_source"`
	ValCSource               OSSTrustNetworkValCCore                  `json:"valc_source"`
	ValDSource               OSSTrustNetworkValDCore                  `json:"vald_source"`
	IntegratedClosure        OSSTrustNetworkValEIntegratedClosureGate `json:"integrated_closure"`
	CanonicalBoundary        OSSTrustNetworkValECanonicalBoundaryGate `json:"canonical_boundary"`
	EvidenceQuality          OSSTrustNetworkValEEvidenceQuality       `json:"evidence_quality"`
	NoOverclaim              OSSTrustNetworkValENoOverclaim           `json:"no_overclaim"`
	ProofSurfaceRefs         []string                                 `json:"proof_surface_refs,omitempty"`
	EvidenceRefs             []string                                 `json:"evidence_refs,omitempty"`
	BlockingReasons          []string                                 `json:"blocking_reasons,omitempty"`
	WhyPoint9Pass            []string                                 `json:"why_point_9_pass,omitempty"`
	IntegratedClosureSummary []string                                 `json:"integrated_closure_summary,omitempty"`
	Limitations              []string                                 `json:"limitations,omitempty"`
	ProjectionDisclaimer     string                                   `json:"projection_disclaimer"`
}

type ossTrustNetworkValEExpectedEvidenceMetadata struct {
	EvidenceType string
	Source       string
	Scope        string
}

func ossTrustNetworkValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth oss_trust_network_vale integrated_closure advisory_projection"
}

func ossTrustNetworkValEHasProjectionDisclaimer(value string) bool {
	return value == ossTrustNetworkValEProjectionDisclaimer() ||
		value == ossTrustNetworkValEProjectionDisclaimer()+" aggregate_dependency_snapshot"
}

func ossTrustNetworkValEHasFoundationProjectionDisclaimer(value string) bool {
	return value == ossTrustNetworkValEProjectionDisclaimer()
}

func OSSTrustNetworkValEProofSurfaceRefs() []string {
	return []string{
		"/v1/oss-trust-network/val0/status",
		"/v1/oss-trust-network/val0/proofs",
		"/v1/oss-trust-network/vala/status",
		"/v1/oss-trust-network/vala/proofs",
		"/v1/oss-trust-network/valb/status",
		"/v1/oss-trust-network/valb/proofs",
		"/v1/oss-trust-network/valc/status",
		"/v1/oss-trust-network/valc/proofs",
		"/v1/oss-trust-network/vald/status",
		"/v1/oss-trust-network/vald/proofs",
		"/v1/oss-trust-network/vale/closure",
		"/v1/oss-trust-network/vale/proofs",
	}
}

func OSSTrustNetworkValEProofEvidenceRefs() []string {
	return []string{
		"evidence:ostn-vale-val0-source-001",
		"evidence:ostn-vale-vala-source-001",
		"evidence:ostn-vale-valb-source-001",
		"evidence:ostn-vale-valc-source-001",
		"evidence:ostn-vale-vald-source-001",
		"evidence:ostn-vale-dependency-001",
		"evidence:ostn-vale-integrated-closure-001",
		"evidence:ostn-vale-canonical-boundary-001",
		"evidence:ostn-vale-evidence-quality-001",
		"evidence:ostn-vale-no-overclaim-001",
		"evidence:ostn-vale-point9-governance-001",
	}
}

func ossTrustNetworkValEExpectedEvidenceMetadataByID() map[string]ossTrustNetworkValEExpectedEvidenceMetadata {
	return map[string]ossTrustNetworkValEExpectedEvidenceMetadata{
		"evidence:ostn-vale-val0-source-001":        {EvidenceType: "dependency_state", Source: "oss-trust-network/vale/val0-source", Scope: "val0_source"},
		"evidence:ostn-vale-vala-source-001":        {EvidenceType: "dependency_state", Source: "oss-trust-network/vale/vala-source", Scope: "vala_source"},
		"evidence:ostn-vale-valb-source-001":        {EvidenceType: "dependency_state", Source: "oss-trust-network/vale/valb-source", Scope: "valb_source"},
		"evidence:ostn-vale-valc-source-001":        {EvidenceType: "dependency_state", Source: "oss-trust-network/vale/valc-source", Scope: "valc_source"},
		"evidence:ostn-vale-vald-source-001":        {EvidenceType: "dependency_state", Source: "oss-trust-network/vale/vald-source", Scope: "vald_source"},
		"evidence:ostn-vale-dependency-001":         {EvidenceType: "dependency_state", Source: "oss-trust-network/vale/dependency", Scope: "vald_dependency"},
		"evidence:ostn-vale-integrated-closure-001": {EvidenceType: "integrated_closure", Source: "oss-trust-network/vale/integrated-closure", Scope: "integrated_closure_gate"},
		"evidence:ostn-vale-canonical-boundary-001": {EvidenceType: "canonical_boundary", Source: "oss-trust-network/vale/canonical-boundary", Scope: "canonical_boundary_gate"},
		"evidence:ostn-vale-evidence-quality-001":   {EvidenceType: "evidence_quality", Source: "oss-trust-network/vale/evidence-quality", Scope: "evidence_quality_gate"},
		"evidence:ostn-vale-no-overclaim-001":       {EvidenceType: "no_overclaim", Source: "oss-trust-network/vale/no-overclaim", Scope: "no_overclaim_gate"},
		"evidence:ostn-vale-point9-governance-001":  {EvidenceType: "state_governance", Source: "oss-trust-network/vale/point9-governance", Scope: "point9_governance"},
	}
}

func ossTrustNetworkValEEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:ostn-vale-val0-source-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vale/val0-source", Timestamp: "2026-04-30T11:00:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "val0_source", Caveats: []string{"Val 0 remains an active prerequisite and stays not complete before integrated closure."}},
		{EvidenceID: "evidence:ostn-vale-vala-source-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vale/vala-source", Timestamp: "2026-04-30T11:01:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "vala_source", Caveats: []string{"Val A remains an active prerequisite and stays not complete before integrated closure."}},
		{EvidenceID: "evidence:ostn-vale-valb-source-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vale/valb-source", Timestamp: "2026-04-30T11:02:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "valb_source", Caveats: []string{"Val B remains an active prerequisite and stays not complete before integrated closure."}},
		{EvidenceID: "evidence:ostn-vale-valc-source-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vale/valc-source", Timestamp: "2026-04-30T11:03:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "valc_source", Caveats: []string{"Val C remains an active prerequisite and stays not complete before integrated closure."}},
		{EvidenceID: "evidence:ostn-vale-vald-source-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vale/vald-source", Timestamp: "2026-04-30T11:04:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "vald_source", Caveats: []string{"Val D remains the final readiness prerequisite and stays not complete before integrated closure."}},
		{EvidenceID: "evidence:ostn-vale-dependency-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vale/dependency", Timestamp: "2026-04-30T11:05:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "vald_dependency", Caveats: []string{"Val E depends on exact and active Val D only and route presence never satisfies the gate."}},
		{EvidenceID: "evidence:ostn-vale-integrated-closure-001", EvidenceType: "integrated_closure", Source: "oss-trust-network/vale/integrated-closure", Timestamp: "2026-04-30T11:06:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "integrated_closure_gate", Caveats: []string{"Integrated closure remains evidence-linked, bounded, and non-authoritative outside the canonical execution, audit, and evidence spine."}},
		{EvidenceID: "evidence:ostn-vale-canonical-boundary-001", EvidenceType: "canonical_boundary", Source: "oss-trust-network/vale/canonical-boundary", Timestamp: "2026-04-30T11:07:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "canonical_boundary_gate", Caveats: []string{"Integrated closure does not authorize deployment or production use and does not turn advisory surfaces into authority."}},
		{EvidenceID: "evidence:ostn-vale-evidence-quality-001", EvidenceType: "evidence_quality", Source: "oss-trust-network/vale/evidence-quality", Timestamp: "2026-04-30T11:08:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "evidence_quality_gate", Caveats: []string{"Val E evidence stays exact, fresh, related, and scope-correct."}},
		{EvidenceID: "evidence:ostn-vale-no-overclaim-001", EvidenceType: "no_overclaim", Source: "oss-trust-network/vale/no-overclaim", Timestamp: "2026-04-30T11:09:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_gate", Caveats: []string{"Val E integrated closure cannot create promotional badge semantics, approval authority, or universal truth semantics."}},
		{EvidenceID: "evidence:ostn-vale-point9-governance-001", EvidenceType: "state_governance", Source: "oss-trust-network/vale/point9-governance", Timestamp: "2026-04-30T11:10:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point9_governance", Caveats: []string{"Only Val E may emit point_9_pass, and only after every closure gate remains active."}},
	}
}

func ossTrustNetworkValECopyEvidence(items []ReferenceArchitectureEvidenceReference) []ReferenceArchitectureEvidenceReference {
	cloned := make([]ReferenceArchitectureEvidenceReference, 0, len(items))
	for _, item := range items {
		cloned = append(cloned, ReferenceArchitectureEvidenceReference{
			EvidenceID:     item.EvidenceID,
			EvidenceType:   item.EvidenceType,
			Source:         item.Source,
			Timestamp:      item.Timestamp,
			FreshnessState: item.FreshnessState,
			Scope:          item.Scope,
			Caveats:        append([]string{}, item.Caveats...),
		})
	}
	return cloned
}

func OSSTrustNetworkValEProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, refs []string) bool {
	if !ossTrustNetworkValEContainsExactStringSet(refs, OSSTrustNetworkValEProofEvidenceRefs()...) {
		return false
	}
	expected := ossTrustNetworkValEExpectedEvidenceMetadataByID()
	if len(evidence) != len(expected) {
		return false
	}
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale {
		return false
	}
	seen := map[string]struct{}{}
	for _, evidenceRef := range evidence {
		id := evidenceRef.EvidenceID
		expectedMetadata, exists := expected[id]
		if id == "" || !exists {
			return false
		}
		if _, duplicate := seen[id]; duplicate {
			return false
		}
		seen[id] = struct{}{}
		if evidenceRef.EvidenceType != expectedMetadata.EvidenceType ||
			evidenceRef.Source != expectedMetadata.Source ||
			evidenceRef.Scope != expectedMetadata.Scope {
			return false
		}
	}
	return len(seen) == len(expected)
}

func ossTrustNetworkValEContainsExactStringSet(values []string, expected ...string) bool {
	if len(values) != len(expected) {
		return false
	}
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		if value == "" {
			return false
		}
		if _, duplicate := seen[value]; duplicate {
			return false
		}
		seen[value] = struct{}{}
	}
	for _, item := range expected {
		if _, ok := seen[item]; !ok {
			return false
		}
	}
	return true
}

func ossTrustNetworkValENormalizeText(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(strings.TrimSpace(value))), " ")
}

func ossTrustNetworkValENormalizeClaimText(value string) string {
	var builder strings.Builder
	lastSpace := true
	for _, char := range strings.TrimSpace(deploymentMultiTenantVal0CompatibilityFold(value)) {
		folded := deploymentMultiTenantVal0ConfusableFold(char)
		if unicode.IsLetter(folded) || unicode.IsDigit(folded) {
			builder.WriteRune(folded)
			lastSpace = false
			continue
		}
		if !lastSpace {
			builder.WriteByte(' ')
			lastSpace = true
		}
	}
	return strings.TrimSpace(builder.String())
}

func ossTrustNetworkValECompactClaimText(value string) string {
	var builder strings.Builder
	for _, char := range deploymentMultiTenantVal0CompatibilityFold(value) {
		folded := deploymentMultiTenantVal0ConfusableFold(char)
		if unicode.IsLetter(folded) || unicode.IsDigit(folded) {
			builder.WriteRune(folded)
		}
	}
	return builder.String()
}

func ossTrustNetworkValEPoint9PassReasonState(value string) string {
	switch value {
	case OSSTrustNetworkValEPoint9PassReasonAllowed:
		return ossTrustNetworkValEPoint9PassReasonStateAllowed
	case OSSTrustNetworkValEPoint9PassReasonBlocked:
		return ossTrustNetworkValEPoint9PassReasonStateBlocked
	default:
		return ossTrustNetworkValEPoint9PassReasonStateUnknown
	}
}

func ossTrustNetworkValEPassAllowedClaim(values ...string) bool {
	for _, value := range values {
		if ossTrustNetworkValEPoint9PassReasonState(value) == ossTrustNetworkValEPoint9PassReasonStateAllowed {
			return true
		}
	}
	return false
}

func ossTrustNetworkValEExactSafePoint9PassDiagnostic(value string) bool {
	if value == "" {
		return false
	}
	safe := []string{
		OSSTrustNetworkValEPoint9PassReasonAllowed,
		OSSTrustNetworkValEPoint9PassReasonBlocked,
		OSSTrustNetworkValEPoint9PassSafeDiagnosticVal0CannotReturn,
		OSSTrustNetworkValEPoint9PassSafeDiagnosticValACannotReturn,
		OSSTrustNetworkValEPoint9PassSafeDiagnosticValBCannotReturn,
		OSSTrustNetworkValEPoint9PassSafeDiagnosticValCCannotReturn,
		OSSTrustNetworkValEPoint9PassSafeDiagnosticValDCannotReturn,
	}
	for _, candidate := range safe {
		if value == candidate {
			return true
		}
	}
	return false
}

func ossTrustNetworkValEContainsForbiddenClaim(values ...string) bool {
	allowedPhrases := []string{
		"integrated ostn closure",
		"bounded oss trust signal",
		"reviewed oss trust signal",
		"candidate oss trust signal",
		"source-weighted reviewed signal",
		"bounded reviewed exchange",
		"local applicability context",
		"package trust visibility",
		"advisory remediation suggestion",
		"pr proposal descriptor",
		"reviewer-required proposal",
		"no hidden mutation path",
		"evidence-linked integrated closure",
		"not canonical truth",
		"not formal certification",
		"not production approval",
		"not deployment approval",
		"not official oss authority",
	}
	// Keep this list aligned with the full required Val E blocked-phrase set.
	disallowed := []string{
		"changelock verified",
		"certified package",
		"officially safe package",
		"regulator-approved",
		"audit passed",
		"compliance guaranteed",
		"production approval",
		"deployment approval",
		"production approved",
		"deployment approved",
		"legal certification",
		"patent cleared",
		"fto cleared",
		"de-facto standard",
		"immune system for open source",
		"universal trust score",
		"integrity score",
		"score > 90",
		"globally safe",
		"automatically blocked everywhere",
		"global truth",
		"crowd-sourced truth layer",
		"reviewed means safe",
		"community truth",
		"network truth",
		"auto-remediated",
		"auto-merged",
		"production autopatch",
		"public badge",
		"official oss authority",
		"certified open-source trust network",
		"officially approved oss network",
		"point_9_pass production approved",
		"point_9_pass certified",
		"point_9_pass global truth",
	}
	allowedExact := make(map[string]struct{}, len(allowedPhrases)*2)
	for _, phrase := range allowedPhrases {
		allowedExact[ossTrustNetworkValENormalizeClaimText(phrase)] = struct{}{}
		allowedExact[ossTrustNetworkValECompactClaimText(phrase)] = struct{}{}
	}
	blockedNormalized := make([]string, 0, len(disallowed))
	blockedCompact := make([]string, 0, len(disallowed))
	for _, blocked := range disallowed {
		blockedNormalized = append(blockedNormalized, ossTrustNetworkValENormalizeClaimText(blocked))
		blockedCompact = append(blockedCompact, ossTrustNetworkValECompactClaimText(blocked))
	}
	point9PassCompact := ossTrustNetworkValECompactClaimText("point_9_pass")
	normalizedBuckets := make([]string, 0, len(values))
	compactBuckets := make([]string, 0, len(values))
	for _, value := range values {
		normalized := ossTrustNetworkValENormalizeClaimText(value)
		compact := ossTrustNetworkValECompactClaimText(value)
		if normalized == "" && compact == "" {
			continue
		}
		if _, ok := allowedExact[normalized]; ok {
			continue
		}
		if _, ok := allowedExact[compact]; ok {
			continue
		}
		normalizedBuckets = append(normalizedBuckets, normalized)
		compactBuckets = append(compactBuckets, compact)
		for i := range blockedNormalized {
			if blockedNormalized[i] != "" && strings.Contains(normalized, blockedNormalized[i]) {
				return true
			}
			if blockedCompact[i] != "" && strings.Contains(compact, blockedCompact[i]) {
				return true
			}
		}
		if strings.Contains(compact, "certifi") && strings.Contains(compact, "package") {
			return true
		}
		if strings.Contains(compact, point9PassCompact) && !ossTrustNetworkValEExactSafePoint9PassDiagnostic(value) {
			return true
		}
	}
	for i := range blockedNormalized {
		if blockedNormalized[i] != "" && deploymentMultiTenantVal0BucketsContainForbiddenPhrase(normalizedBuckets, blockedNormalized[i]) {
			return true
		}
		if blockedCompact[i] != "" && deploymentMultiTenantVal0BucketsContainForbiddenPhrase(compactBuckets, blockedCompact[i]) {
			return true
		}
	}
	return false
}

func ossTrustNetworkValECollectText(values []string) []string {
	collected := make([]string, 0, len(values))
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		collected = append(collected, trimmed)
	}
	return collected
}

func ossTrustNetworkValEGenericStateSeverity(value string) int {
	trimmed := strings.TrimSpace(value)
	switch {
	case trimmed == "":
		return 2
	case strings.HasSuffix(trimmed, "_blocked"):
		return 4
	case strings.HasSuffix(trimmed, "_unknown"):
		return 3
	case strings.HasSuffix(trimmed, "_incomplete"):
		return 2
	case strings.HasSuffix(trimmed, "_partial"):
		return 1
	default:
		return 0
	}
}

func ossTrustNetworkValEGenericStateFromSeverity(severity int) string {
	switch severity {
	case 4:
		return OSSTrustNetworkValESourceStateBlocked
	case 3:
		return OSSTrustNetworkValESourceStateUnknown
	case 2:
		return OSSTrustNetworkValESourceStateIncomplete
	default:
		return OSSTrustNetworkValESourceStatePartial
	}
}

func OSSTrustNetworkValEIntegratedClosureGateModel() OSSTrustNetworkValEIntegratedClosureGate {
	return OSSTrustNetworkValEIntegratedClosureGate{
		GateID:  "oss-trust-network-vale-integrated-closure",
		Version: "v0",
		CanonicalExecutionAuditEvidenceSourceOfTruth: true,
		OSSNetworkOutputsRemainAdvisory:              true,
		EvidenceLinkedIntegratedClosure:              true,
		LocalEnterpriseApplicabilityExplicit:         true,
		EvidenceRefs:                                 []string{"evidence:ostn-vale-integrated-closure-001", "evidence:ostn-vale-point9-governance-001"},
		Caveats:                                      []string{"Integrated closure remains bounded and non-authoritative outside the canonical execution, audit, and evidence spine."},
		ProjectionDisclaimer:                         ossTrustNetworkValEProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValECanonicalBoundaryGateModel() OSSTrustNetworkValECanonicalBoundaryGate {
	return OSSTrustNetworkValECanonicalBoundaryGate{
		BoundaryID:                           "oss-trust-network-vale-canonical-boundary",
		LocalOverrideVisibleEvidenceLinked:   true,
		LocalOverrideVisible:                 true,
		LocalEnterpriseApplicabilityExplicit: true,
		EvidenceRefs:                         []string{"evidence:ostn-vale-canonical-boundary-001"},
		Caveats:                              []string{"Integrated closure cannot turn advisory, remediation, proposal, registry, community, or propagation surfaces into authority."},
		ProjectionDisclaimer:                 ossTrustNetworkValEProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValEEvidenceQualityModel() OSSTrustNetworkValEEvidenceQuality {
	valD := ComputeOSSTrustNetworkValDCore(OSSTrustNetworkValDCoreModel())
	return OSSTrustNetworkValEEvidenceQuality{
		EvidenceQualityID:              "oss-trust-network-vale-evidence-quality",
		Evidence:                       ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence()),
		ProofSurfaceRefs:               OSSTrustNetworkValEProofSurfaceRefs(),
		EvidenceRefs:                   OSSTrustNetworkValEProofEvidenceRefs(),
		DependencyProofSurfaceRefs:     append([]string{}, valD.ProofSurfaceRefs...),
		DependencyEvidenceRefs:         append([]string{}, valD.EvidenceRefs...),
		DependencyEvidence:             ossTrustNetworkValECopyEvidence(ossTrustNetworkValDEvidence()),
		DependencyProjectionDisclaimer: valD.ProjectionDisclaimer,
		ProjectionDisclaimer:           ossTrustNetworkValEProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValENoOverclaimModel() OSSTrustNetworkValENoOverclaim {
	return OSSTrustNetworkValENoOverclaim{
		GateID:               "oss-trust-network-vale-no-overclaim",
		Version:              "v0",
		ObservedClaims:       []string{"integrated ostn closure", "evidence-linked integrated closure", "not canonical truth", "not production approval", "not deployment approval", "not official oss authority"},
		ProjectionDisclaimer: ossTrustNetworkValEProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValEIntegratedClosureModel() OSSTrustNetworkValEIntegratedClosure {
	return OSSTrustNetworkValEIntegratedClosure{
		Point9PassAllowed: false,
		Point9PassReason:  OSSTrustNetworkValEPoint9PassReasonBlocked,
		Val0Source:        ComputeOSSTrustNetworkVal0Foundation(OSSTrustNetworkVal0FoundationModel()),
		ValASource:        ComputeOSSTrustNetworkValACore(OSSTrustNetworkValACoreModel()),
		ValBSource:        ComputeOSSTrustNetworkValBCore(OSSTrustNetworkValBCoreModel()),
		ValCSource:        ComputeOSSTrustNetworkValCCore(OSSTrustNetworkValCCoreModel()),
		ValDSource:        ComputeOSSTrustNetworkValDCore(OSSTrustNetworkValDCoreModel()),
		IntegratedClosure: OSSTrustNetworkValEIntegratedClosureGateModel(),
		CanonicalBoundary: OSSTrustNetworkValECanonicalBoundaryGateModel(),
		EvidenceQuality:   OSSTrustNetworkValEEvidenceQualityModel(),
		NoOverclaim:       OSSTrustNetworkValENoOverclaimModel(),
		ProofSurfaceRefs:  OSSTrustNetworkValEProofSurfaceRefs(),
		EvidenceRefs:      OSSTrustNetworkValEProofEvidenceRefs(),
		WhyPoint9Pass: []string{
			"Only Val E may emit point_9_pass and only after exact Val 0 through Val D states, proof refs, evidence refs, evidence quality, canonical boundaries, and no-overclaim discipline remain active.",
			"Registry, maintainer, shared VEX, dashboard, remediation, proposal, and propagation surfaces remain bounded and non-authoritative even when point_9_pass is active.",
		},
		IntegratedClosureSummary: []string{
			"Val E integrates Val 0 through Val D into one evidence-linked closure gate over bounded OSTN behavior.",
			"Val E may emit final Point 9 pass semantics only when every dependency, evidence, canonical boundary, and no-overclaim rule remains active.",
		},
		Limitations: []string{
			"Val E closes Točka 9 only and does not implement Točka 10 or any new live network connector behavior.",
			"Integrated closure does not authorize deployment, production use, regulator-facing approvals, legal or IP clearance, or promotional badge semantics.",
			"External review or audit handoff remains outside this implementation and belongs in later program-close material.",
		},
		ProjectionDisclaimer: ossTrustNetworkValEProjectionDisclaimer(),
	}
}

func EvaluateOSSTrustNetworkValEVal0SourceState(model OSSTrustNetworkVal0Foundation) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CurrentState,
		model.Point9State,
		model.DependencyState,
		model.SignalContractState,
		model.TrustMarkingState,
		model.MaintainerIdentityState,
		model.RegistryFreshnessState,
		model.SharedVEXState,
		model.PropagationState,
		model.LocalApplicabilityState,
		model.NoOverclaimState,
		model.ProjectionDisclaimer,
	) || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return OSSTrustNetworkValESourceStateIncomplete
	}
	if model.ProjectionDisclaimer != ossTrustNetworkVal0ProjectionDisclaimer() {
		return OSSTrustNetworkValESourceStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(model.ProofSurfaceRefs, OSSTrustNetworkVal0ProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.EvidenceRefs, OSSTrustNetworkVal0ProofEvidenceRefs()...) ||
		!OSSTrustNetworkVal0ProofEvidenceQualityValid(ossTrustNetworkVal0Evidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValESourceStateBlocked
	}
	if model.CurrentState == OSSTrustNetworkVal0StateActive &&
		model.Point9State == OSSTrustNetworkPoint9StateNotComplete &&
		model.DependencyState == OSSTrustNetworkVal0DependencyStateActive &&
		model.SignalContractState == OSSTrustNetworkVal0SignalContractStateActive &&
		model.TrustMarkingState == OSSTrustNetworkVal0TrustMarkingStateActive &&
		model.MaintainerIdentityState == OSSTrustNetworkVal0MaintainerIdentityStateActive &&
		model.RegistryFreshnessState == OSSTrustNetworkVal0RegistryFreshnessStateActive &&
		model.SharedVEXState == OSSTrustNetworkVal0SharedVEXStateActive &&
		model.PropagationState == OSSTrustNetworkVal0PropagationStateActive &&
		model.LocalApplicabilityState == OSSTrustNetworkVal0LocalApplicabilityStateActive &&
		model.NoOverclaimState == OSSTrustNetworkVal0NoOverclaimStateActive {
		return OSSTrustNetworkValESourceStateActive
	}
	severity := 0
	for _, state := range []string{
		model.CurrentState,
		model.DependencyState,
		model.SignalContractState,
		model.TrustMarkingState,
		model.MaintainerIdentityState,
		model.RegistryFreshnessState,
		model.SharedVEXState,
		model.PropagationState,
		model.LocalApplicabilityState,
		model.NoOverclaimState,
	} {
		severity = maxStateSeverity(severity, ossTrustNetworkValEGenericStateSeverity(state))
	}
	return ossTrustNetworkValEGenericStateFromSeverity(severity)
}

func EvaluateOSSTrustNetworkValEValASourceState(model OSSTrustNetworkValACore) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CurrentState,
		model.Point9State,
		model.DependencyState,
		model.ReleaseTrustIntakeState,
		model.SigningSignalState,
		model.MaintainerAttestationState,
		model.ProvenanceMaterialState,
		model.RegistryDescriptorState,
		model.RegistryMetadataState,
		model.TypoSquattingWarningState,
		model.DriftSignalState,
		model.NoOverclaimState,
		model.ProjectionDisclaimer,
	) || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return OSSTrustNetworkValESourceStateIncomplete
	}
	if model.ProjectionDisclaimer != ossTrustNetworkValAProjectionDisclaimer() {
		return OSSTrustNetworkValESourceStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValAProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.EvidenceRefs, OSSTrustNetworkValAProofEvidenceRefs()...) ||
		!OSSTrustNetworkValAProofEvidenceQualityValid(ossTrustNetworkValAEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValESourceStateBlocked
	}
	if model.CurrentState == OSSTrustNetworkValAStateActive &&
		model.Point9State == OSSTrustNetworkPoint9StateNotComplete &&
		model.DependencyState == OSSTrustNetworkValADependencyStateActive &&
		model.ReleaseTrustIntakeState == OSSTrustNetworkValAReleaseTrustIntakeStateActive &&
		model.SigningSignalState == OSSTrustNetworkValASigningSignalStateActive &&
		model.MaintainerAttestationState == OSSTrustNetworkValAMaintainerAttestationStateActive &&
		model.ProvenanceMaterialState == OSSTrustNetworkValAProvenanceMaterialStateActive &&
		model.RegistryDescriptorState == OSSTrustNetworkValARegistryDescriptorStateActive &&
		model.RegistryMetadataState == OSSTrustNetworkValARegistryMetadataStateActive &&
		model.TypoSquattingWarningState == OSSTrustNetworkValATypoSquattingWarningStateActive &&
		model.DriftSignalState == OSSTrustNetworkValADriftSignalStateActive &&
		model.NoOverclaimState == OSSTrustNetworkValANoOverclaimStateActive {
		return OSSTrustNetworkValESourceStateActive
	}
	severity := 0
	for _, state := range []string{
		model.CurrentState,
		model.DependencyState,
		model.ReleaseTrustIntakeState,
		model.SigningSignalState,
		model.MaintainerAttestationState,
		model.ProvenanceMaterialState,
		model.RegistryDescriptorState,
		model.RegistryMetadataState,
		model.TypoSquattingWarningState,
		model.DriftSignalState,
		model.NoOverclaimState,
	} {
		severity = maxStateSeverity(severity, ossTrustNetworkValEGenericStateSeverity(state))
	}
	return ossTrustNetworkValEGenericStateFromSeverity(severity)
}

func EvaluateOSSTrustNetworkValEValBSourceState(model OSSTrustNetworkValBCore) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CurrentState,
		model.Point9State,
		model.DependencyState,
		model.CandidateSignalIntakeState,
		model.ReviewWorkflowState,
		model.SharedVEXTriageState,
		model.SourceWeightingState,
		model.LocalApplicabilityState,
		model.PropagationExchangeState,
		model.SupersessionRevocationState,
		model.ReviewerAuditabilityState,
		model.NoOverclaimState,
		model.ProjectionDisclaimer,
	) || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return OSSTrustNetworkValESourceStateIncomplete
	}
	if model.ProjectionDisclaimer != ossTrustNetworkValBProjectionDisclaimer() {
		return OSSTrustNetworkValESourceStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValBProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.EvidenceRefs, OSSTrustNetworkValBProofEvidenceRefs()...) ||
		!OSSTrustNetworkValBProofEvidenceQualityValid(ossTrustNetworkValBEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValESourceStateBlocked
	}
	if model.CurrentState == OSSTrustNetworkValBStateActive &&
		model.Point9State == OSSTrustNetworkPoint9StateNotComplete &&
		model.DependencyState == OSSTrustNetworkValBDependencyStateActive &&
		model.CandidateSignalIntakeState == OSSTrustNetworkValBCandidateSignalIntakeStateActive &&
		model.ReviewWorkflowState == OSSTrustNetworkValBReviewWorkflowStateActive &&
		model.SharedVEXTriageState == OSSTrustNetworkValBSharedVEXTriageStateActive &&
		model.SourceWeightingState == OSSTrustNetworkValBSourceWeightingStateActive &&
		model.LocalApplicabilityState == OSSTrustNetworkValBLocalApplicabilityStateActive &&
		model.PropagationExchangeState == OSSTrustNetworkValBPropagationExchangeStateActive &&
		model.SupersessionRevocationState == OSSTrustNetworkValBSupersessionRevocationStateActive &&
		model.ReviewerAuditabilityState == OSSTrustNetworkValBReviewerAuditabilityStateActive &&
		model.NoOverclaimState == OSSTrustNetworkValBNoOverclaimStateActive {
		return OSSTrustNetworkValESourceStateActive
	}
	severity := 0
	for _, state := range []string{
		model.CurrentState,
		model.DependencyState,
		model.CandidateSignalIntakeState,
		model.ReviewWorkflowState,
		model.SharedVEXTriageState,
		model.SourceWeightingState,
		model.LocalApplicabilityState,
		model.PropagationExchangeState,
		model.SupersessionRevocationState,
		model.ReviewerAuditabilityState,
		model.NoOverclaimState,
	} {
		severity = maxStateSeverity(severity, ossTrustNetworkValEGenericStateSeverity(state))
	}
	return ossTrustNetworkValEGenericStateFromSeverity(severity)
}

func EvaluateOSSTrustNetworkValEValCSourceState(model OSSTrustNetworkValCCore) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CurrentState,
		model.Point9State,
		model.DependencyState,
		model.TrustVisibilityState,
		model.PackageTrustStatusState,
		model.ExportBoundaryState,
		model.RemediationSuggestionState,
		model.PRProposalState,
		model.LocalOverrideState,
		model.RemediationSafetyState,
		model.EcosystemConsistencyState,
		model.NoOverclaimState,
		model.ProjectionDisclaimer,
	) || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return OSSTrustNetworkValESourceStateIncomplete
	}
	if model.ProjectionDisclaimer != ossTrustNetworkValCProjectionDisclaimer() {
		return OSSTrustNetworkValESourceStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValCProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.EvidenceRefs, OSSTrustNetworkValCProofEvidenceRefs()...) ||
		!OSSTrustNetworkValCProofEvidenceQualityValid(ossTrustNetworkValCEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValESourceStateBlocked
	}
	if model.CurrentState == OSSTrustNetworkValCStateActive &&
		model.Point9State == OSSTrustNetworkPoint9StateNotComplete &&
		model.DependencyState == OSSTrustNetworkValCDependencyStateActive &&
		model.TrustVisibilityState == OSSTrustNetworkValCTrustVisibilityStateActive &&
		model.PackageTrustStatusState == OSSTrustNetworkValCPackageTrustStatusStateActive &&
		model.ExportBoundaryState == OSSTrustNetworkValCExportBoundaryStateActive &&
		model.RemediationSuggestionState == OSSTrustNetworkValCRemediationSuggestionStateActive &&
		model.PRProposalState == OSSTrustNetworkValCPRProposalStateActive &&
		model.LocalOverrideState == OSSTrustNetworkValCLocalOverrideStateActive &&
		model.RemediationSafetyState == OSSTrustNetworkValCRemediationSafetyStateActive &&
		model.EcosystemConsistencyState == OSSTrustNetworkValCEcosystemConsistencyStateActive &&
		model.NoOverclaimState == OSSTrustNetworkValCNoOverclaimStateActive {
		return OSSTrustNetworkValESourceStateActive
	}
	severity := 0
	for _, state := range []string{
		model.CurrentState,
		model.DependencyState,
		model.TrustVisibilityState,
		model.PackageTrustStatusState,
		model.ExportBoundaryState,
		model.RemediationSuggestionState,
		model.PRProposalState,
		model.LocalOverrideState,
		model.RemediationSafetyState,
		model.EcosystemConsistencyState,
		model.NoOverclaimState,
	} {
		severity = maxStateSeverity(severity, ossTrustNetworkValEGenericStateSeverity(state))
	}
	return ossTrustNetworkValEGenericStateFromSeverity(severity)
}

func EvaluateOSSTrustNetworkValEValDSourceState(model OSSTrustNetworkValDCore) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CurrentState,
		model.Point9State,
		model.DependencyState,
		model.SignalCorrectnessState,
		model.ReleaseFoundationState,
		model.ReviewedIntelligenceState,
		model.PropagationSafetyState,
		model.RemediationPRSafetyState,
		model.EcosystemVisibilityConsistencyState,
		model.EvidenceQualityState,
		model.NoOverclaimState,
		model.ProjectionDisclaimer,
	) || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return OSSTrustNetworkValESourceStateIncomplete
	}
	if model.ProjectionDisclaimer != ossTrustNetworkValDProjectionDisclaimer() {
		return OSSTrustNetworkValESourceStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValDProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.EvidenceRefs, OSSTrustNetworkValDProofEvidenceRefs()...) ||
		!OSSTrustNetworkValDProofEvidenceQualityValid(ossTrustNetworkValDEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValESourceStateBlocked
	}
	if model.CurrentState == OSSTrustNetworkValDStateActive &&
		model.Point9State == OSSTrustNetworkPoint9StateNotComplete &&
		model.DependencyState == OSSTrustNetworkValDDependencyStateActive &&
		model.SignalCorrectnessState == OSSTrustNetworkValDSignalCorrectnessStateActive &&
		model.ReleaseFoundationState == OSSTrustNetworkValDReleaseFoundationStateActive &&
		model.ReviewedIntelligenceState == OSSTrustNetworkValDReviewedIntelligenceStateActive &&
		model.PropagationSafetyState == OSSTrustNetworkValDPropagationSafetyStateActive &&
		model.RemediationPRSafetyState == OSSTrustNetworkValDRemediationPRSafetyStateActive &&
		model.EcosystemVisibilityConsistencyState == OSSTrustNetworkValDEcosystemVisibilityConsistencyStateActive &&
		model.EvidenceQualityState == OSSTrustNetworkValDEvidenceQualityStateActive &&
		model.NoOverclaimState == OSSTrustNetworkValDNoOverclaimStateActive {
		return OSSTrustNetworkValESourceStateActive
	}
	severity := 0
	for _, state := range []string{
		model.CurrentState,
		model.DependencyState,
		model.SignalCorrectnessState,
		model.ReleaseFoundationState,
		model.ReviewedIntelligenceState,
		model.PropagationSafetyState,
		model.RemediationPRSafetyState,
		model.EcosystemVisibilityConsistencyState,
		model.EvidenceQualityState,
		model.NoOverclaimState,
	} {
		severity = maxStateSeverity(severity, ossTrustNetworkValEGenericStateSeverity(state))
	}
	return ossTrustNetworkValEGenericStateFromSeverity(severity)
}

func EvaluateOSSTrustNetworkValEDependencyState(model OSSTrustNetworkValEIntegratedClosure) string {
	if strings.TrimSpace(model.ProjectionDisclaimer) == "" || strings.TrimSpace(model.Point9PassReason) == "" {
		return OSSTrustNetworkValEDependencyStateIncomplete
	}
	if !ossTrustNetworkValEHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValEDependencyStateUnknown
	}
	if model.ValDSourceState != OSSTrustNetworkValESourceStateActive ||
		!ossTrustNetworkValEContainsExactStringSet(model.ValDSource.ProofSurfaceRefs, OSSTrustNetworkValDProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.ValDSource.EvidenceRefs, OSSTrustNetworkValDProofEvidenceRefs()...) ||
		!OSSTrustNetworkValDProofEvidenceQualityValid(ossTrustNetworkValDEvidence(), model.ValDSource.EvidenceRefs) {
		return OSSTrustNetworkValEDependencyStateBlocked
	}
	if model.ValDSource.CurrentState == OSSTrustNetworkValDStateActive &&
		model.ValDSource.Point9State == OSSTrustNetworkPoint9StateNotComplete &&
		model.ValDSource.DependencyState == OSSTrustNetworkValDDependencyStateActive &&
		model.ValDSource.SignalCorrectnessState == OSSTrustNetworkValDSignalCorrectnessStateActive &&
		model.ValDSource.ReleaseFoundationState == OSSTrustNetworkValDReleaseFoundationStateActive &&
		model.ValDSource.ReviewedIntelligenceState == OSSTrustNetworkValDReviewedIntelligenceStateActive &&
		model.ValDSource.PropagationSafetyState == OSSTrustNetworkValDPropagationSafetyStateActive &&
		model.ValDSource.RemediationPRSafetyState == OSSTrustNetworkValDRemediationPRSafetyStateActive &&
		model.ValDSource.EcosystemVisibilityConsistencyState == OSSTrustNetworkValDEcosystemVisibilityConsistencyStateActive &&
		model.ValDSource.EvidenceQualityState == OSSTrustNetworkValDEvidenceQualityStateActive &&
		model.ValDSource.NoOverclaimState == OSSTrustNetworkValDNoOverclaimStateActive {
		return OSSTrustNetworkValEDependencyStateActive
	}
	return OSSTrustNetworkValEDependencyStateBlocked
}

func EvaluateOSSTrustNetworkValEIntegratedClosureState(model OSSTrustNetworkValEIntegratedClosure) string {
	gate := model.IntegratedClosure
	if !referenceArchitectureValBRequiredRefsPresent(
		gate.GateID,
		gate.Version,
		model.Val0SourceState,
		model.ValASourceState,
		model.ValBSourceState,
		model.ValCSourceState,
		model.ValDSourceState,
		model.DependencyState,
		gate.ProjectionDisclaimer,
	) || len(gate.EvidenceRefs) == 0 || len(gate.Caveats) == 0 {
		return OSSTrustNetworkValEIntegratedClosureStateIncomplete
	}
	if gate.GateID != OSSTrustNetworkValEIntegratedClosureGateModel().GateID ||
		gate.Version != OSSTrustNetworkValEIntegratedClosureGateModel().Version {
		return OSSTrustNetworkValEIntegratedClosureStateBlocked
	}
	if !ossTrustNetworkValEHasFoundationProjectionDisclaimer(gate.ProjectionDisclaimer) {
		return OSSTrustNetworkValEIntegratedClosureStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(gate.EvidenceRefs, "evidence:ostn-vale-integrated-closure-001", "evidence:ostn-vale-point9-governance-001") {
		return OSSTrustNetworkValEIntegratedClosureStateBlocked
	}
	if model.Val0Source.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValASource.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValBSource.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValCSource.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValDSource.Point9State != OSSTrustNetworkPoint9StateNotComplete {
		return OSSTrustNetworkValEIntegratedClosureStateBlocked
	}
	if !gate.CanonicalExecutionAuditEvidenceSourceOfTruth ||
		!gate.OSSNetworkOutputsRemainAdvisory ||
		!gate.EvidenceLinkedIntegratedClosure ||
		!gate.LocalEnterpriseApplicabilityExplicit {
		return OSSTrustNetworkValEIntegratedClosureStateBlocked
	}
	if gate.CanonicalExecutionAuditEvidenceSourceOfTruth &&
		gate.OSSNetworkOutputsRemainAdvisory &&
		gate.EvidenceLinkedIntegratedClosure &&
		gate.LocalEnterpriseApplicabilityExplicit &&
		model.DependencyState == OSSTrustNetworkValEDependencyStateActive &&
		model.Val0SourceState == OSSTrustNetworkValESourceStateActive &&
		model.ValASourceState == OSSTrustNetworkValESourceStateActive &&
		model.ValBSourceState == OSSTrustNetworkValESourceStateActive &&
		model.ValCSourceState == OSSTrustNetworkValESourceStateActive &&
		model.ValDSourceState == OSSTrustNetworkValESourceStateActive {
		return OSSTrustNetworkValEIntegratedClosureStateActive
	}
	severity := 0
	for _, state := range []string{
		model.Val0SourceState,
		model.ValASourceState,
		model.ValBSourceState,
		model.ValCSourceState,
		model.ValDSourceState,
		model.DependencyState,
	} {
		severity = maxStateSeverity(severity, ossTrustNetworkValEGenericStateSeverity(state))
	}
	if result, ok := map[int]string{
		4: OSSTrustNetworkValEIntegratedClosureStateBlocked,
		3: OSSTrustNetworkValEIntegratedClosureStateUnknown,
		2: OSSTrustNetworkValEIntegratedClosureStateIncomplete,
	}[severity]; ok {
		return result
	}
	return OSSTrustNetworkValEIntegratedClosureStatePartial
}

func EvaluateOSSTrustNetworkValECanonicalBoundaryState(model OSSTrustNetworkValEIntegratedClosure) string {
	gate := model.CanonicalBoundary
	if !referenceArchitectureValBRequiredRefsPresent(
		gate.BoundaryID,
		model.ValBSourceState,
		model.ValCSourceState,
		model.ValDSourceState,
		gate.ProjectionDisclaimer,
	) || len(gate.EvidenceRefs) == 0 || len(gate.Caveats) == 0 {
		return OSSTrustNetworkValECanonicalBoundaryStateIncomplete
	}
	if gate.BoundaryID != OSSTrustNetworkValECanonicalBoundaryGateModel().BoundaryID {
		return OSSTrustNetworkValECanonicalBoundaryStateBlocked
	}
	if !ossTrustNetworkValEHasFoundationProjectionDisclaimer(gate.ProjectionDisclaimer) {
		return OSSTrustNetworkValECanonicalBoundaryStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(gate.EvidenceRefs, "evidence:ostn-vale-canonical-boundary-001") {
		return OSSTrustNetworkValECanonicalBoundaryStateBlocked
	}
	if gate.RegistrySurfaceAuthorityClaim ||
		gate.CommunitySurfaceAuthorityClaim ||
		gate.MaintainerSurfaceAuthorityClaim ||
		gate.SharedVEXSurfaceAuthorityClaim ||
		gate.DashboardSurfaceAuthorityClaim ||
		gate.RemediationSurfaceAuthorityClaim ||
		gate.PRProposalSurfaceAuthorityClaim ||
		gate.PropagationSurfaceAuthorityClaim ||
		gate.OfficialOSSAuthorityClaim ||
		gate.SharedSignalSilentOverride ||
		gate.FinalClosureApprovesDeployment ||
		gate.FinalClosureApprovesProductionUse ||
		!gate.LocalOverrideVisibleEvidenceLinked ||
		!gate.LocalOverrideVisible ||
		!gate.LocalEnterpriseApplicabilityExplicit ||
		model.ValBSource.LocalApplicabilityState != OSSTrustNetworkValBLocalApplicabilityStateActive ||
		model.ValCSource.LocalOverrideState != OSSTrustNetworkValCLocalOverrideStateActive ||
		!model.ValCSource.LocalOverride.LocalOnlyBoundary ||
		len(model.ValCSource.LocalOverride.EvidenceRefs) == 0 ||
		strings.TrimSpace(model.ValCSource.LocalOverride.Rationale) == "" ||
		model.ValCSource.LocalOverride.RewriteCanonicalEvidence ||
		model.ValCSource.LocalOverride.SilentlySuppressNetworkIntelligence ||
		model.ValCSource.LocalOverride.SharedSignalOverridesLocalDecision {
		return OSSTrustNetworkValECanonicalBoundaryStateBlocked
	}
	return OSSTrustNetworkValECanonicalBoundaryStateActive
}

func EvaluateOSSTrustNetworkValEEvidenceQualityState(model OSSTrustNetworkValEEvidenceQuality) string {
	if len(model.Evidence) == 0 || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 ||
		len(model.DependencyProofSurfaceRefs) == 0 || len(model.DependencyEvidenceRefs) == 0 || len(model.DependencyEvidence) == 0 ||
		strings.TrimSpace(model.DependencyProjectionDisclaimer) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValEEvidenceQualityStateIncomplete
	}
	if !ossTrustNetworkValEHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!ossTrustNetworkValDHasProjectionDisclaimer(model.DependencyProjectionDisclaimer) {
		return OSSTrustNetworkValEEvidenceQualityStateUnknown
	}
	if !ossTrustNetworkValEContainsExactStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValEProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.EvidenceRefs, OSSTrustNetworkValEProofEvidenceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.DependencyProofSurfaceRefs, OSSTrustNetworkValDProofSurfaceRefs()...) ||
		!ossTrustNetworkValEContainsExactStringSet(model.DependencyEvidenceRefs, OSSTrustNetworkValDProofEvidenceRefs()...) ||
		!OSSTrustNetworkValEProofEvidenceQualityValid(model.Evidence, model.EvidenceRefs) ||
		!OSSTrustNetworkValDProofEvidenceQualityValid(model.DependencyEvidence, model.DependencyEvidenceRefs) {
		return OSSTrustNetworkValEEvidenceQualityStateBlocked
	}
	return OSSTrustNetworkValEEvidenceQualityStateActive
}

func EvaluateOSSTrustNetworkValENoOverclaimState(model OSSTrustNetworkValEIntegratedClosure) string {
	if strings.TrimSpace(model.Point9PassReason) == "" || strings.TrimSpace(model.NoOverclaim.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValENoOverclaimStateIncomplete
	}
	if !ossTrustNetworkValEHasFoundationProjectionDisclaimer(model.NoOverclaim.ProjectionDisclaimer) {
		return OSSTrustNetworkValENoOverclaimStateUnknown
	}
	if model.NoOverclaim.Point9PassClaim ||
		model.NoOverclaim.CertifiesTrust ||
		model.NoOverclaim.ApprovesProductionUse ||
		model.NoOverclaim.ApprovesDeployment ||
		model.NoOverclaim.RegulatorApprovalClaim ||
		model.NoOverclaim.LegalIPClearanceClaim ||
		model.NoOverclaim.PublicBadgeClaim ||
		model.NoOverclaim.OfficialOSSAuthorityClaim ||
		model.NoOverclaim.GlobalTruthClaim {
		return OSSTrustNetworkValENoOverclaimStateBlocked
	}
	claims := append([]string{model.Point9PassReason}, model.NoOverclaim.ObservedClaims...)
	if ossTrustNetworkValEContainsForbiddenClaim(claims...) {
		return OSSTrustNetworkValENoOverclaimStateBlocked
	}
	if model.Val0Source.NoOverclaimState != OSSTrustNetworkVal0NoOverclaimStateActive ||
		model.ValASource.NoOverclaimState != OSSTrustNetworkValANoOverclaimStateActive ||
		model.ValBSource.NoOverclaimState != OSSTrustNetworkValBNoOverclaimStateActive ||
		model.ValCSource.NoOverclaimState != OSSTrustNetworkValCNoOverclaimStateActive ||
		model.ValDSource.NoOverclaimState != OSSTrustNetworkValDNoOverclaimStateActive {
		return OSSTrustNetworkValENoOverclaimStateBlocked
	}
	return OSSTrustNetworkValENoOverclaimStateActive
}

func EvaluateOSSTrustNetworkValEFinalPassRuleState(model OSSTrustNetworkValEIntegratedClosure) string {
	if strings.TrimSpace(model.Point9PassReason) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValEFinalPassRuleStateIncomplete
	}
	reasonState := ossTrustNetworkValEPoint9PassReasonState(model.Point9PassReason)
	if reasonState == ossTrustNetworkValEPoint9PassReasonStateUnknown &&
		ossTrustNetworkValEContainsForbiddenClaim(model.Point9PassReason) {
		return OSSTrustNetworkValEFinalPassRuleStateBlocked
	}
	if model.DependencyState == OSSTrustNetworkValEDependencyStateActive &&
		model.IntegratedClosureState == OSSTrustNetworkValEIntegratedClosureStateActive &&
		model.CanonicalBoundaryState == OSSTrustNetworkValECanonicalBoundaryStateActive &&
		model.EvidenceQualityState == OSSTrustNetworkValEEvidenceQualityStateActive &&
		model.NoOverclaimState == OSSTrustNetworkValENoOverclaimStateActive &&
		!model.CanonicalBoundary.RegistrySurfaceAuthorityClaim &&
		!model.CanonicalBoundary.CommunitySurfaceAuthorityClaim &&
		!model.CanonicalBoundary.MaintainerSurfaceAuthorityClaim &&
		!model.CanonicalBoundary.SharedVEXSurfaceAuthorityClaim &&
		!model.CanonicalBoundary.DashboardSurfaceAuthorityClaim &&
		!model.CanonicalBoundary.RemediationSurfaceAuthorityClaim &&
		!model.CanonicalBoundary.PRProposalSurfaceAuthorityClaim &&
		!model.CanonicalBoundary.PropagationSurfaceAuthorityClaim &&
		!model.CanonicalBoundary.OfficialOSSAuthorityClaim &&
		!model.CanonicalBoundary.SharedSignalSilentOverride &&
		!model.CanonicalBoundary.FinalClosureApprovesDeployment &&
		!model.CanonicalBoundary.FinalClosureApprovesProductionUse &&
		ossTrustNetworkValEPassAllowedClaim(model.Point9PassReason) {
		return OSSTrustNetworkValEFinalPassRuleStateActive
	}
	if reasonState == ossTrustNetworkValEPoint9PassReasonStateBlocked {
		return OSSTrustNetworkValEFinalPassRuleStateBlocked
	}
	if model.DependencyState == OSSTrustNetworkValEDependencyStateBlocked ||
		model.IntegratedClosureState == OSSTrustNetworkValEIntegratedClosureStateBlocked ||
		model.CanonicalBoundaryState == OSSTrustNetworkValECanonicalBoundaryStateBlocked ||
		model.EvidenceQualityState == OSSTrustNetworkValEEvidenceQualityStateBlocked ||
		model.NoOverclaimState == OSSTrustNetworkValENoOverclaimStateBlocked ||
		ossTrustNetworkValEContainsForbiddenClaim(model.Point9PassReason) {
		return OSSTrustNetworkValEFinalPassRuleStateBlocked
	}
	if model.DependencyState == OSSTrustNetworkValEDependencyStateIncomplete ||
		model.IntegratedClosureState == OSSTrustNetworkValEIntegratedClosureStateIncomplete ||
		model.CanonicalBoundaryState == OSSTrustNetworkValECanonicalBoundaryStateIncomplete ||
		model.EvidenceQualityState == OSSTrustNetworkValEEvidenceQualityStateIncomplete ||
		model.NoOverclaimState == OSSTrustNetworkValENoOverclaimStateIncomplete {
		return OSSTrustNetworkValEFinalPassRuleStateIncomplete
	}
	if model.DependencyState == OSSTrustNetworkValEDependencyStateUnknown ||
		model.IntegratedClosureState == OSSTrustNetworkValEIntegratedClosureStateUnknown ||
		model.CanonicalBoundaryState == OSSTrustNetworkValECanonicalBoundaryStateUnknown ||
		model.EvidenceQualityState == OSSTrustNetworkValEEvidenceQualityStateUnknown ||
		model.NoOverclaimState == OSSTrustNetworkValENoOverclaimStateUnknown {
		return OSSTrustNetworkValEFinalPassRuleStateUnknown
	}
	if reasonState == ossTrustNetworkValEPoint9PassReasonStateUnknown {
		return OSSTrustNetworkValEFinalPassRuleStateUnknown
	}
	return OSSTrustNetworkValEFinalPassRuleStatePartial
}

func ossTrustNetworkValECanPromotePoint9PassReason(model OSSTrustNetworkValEIntegratedClosure) bool {
	if ossTrustNetworkValEPoint9PassReasonState(model.Point9PassReason) == ossTrustNetworkValEPoint9PassReasonStateUnknown &&
		!ossTrustNetworkValEExactSafePoint9PassDiagnostic(model.Point9PassReason) {
		return false
	}
	return model.DependencyState == OSSTrustNetworkValEDependencyStateActive &&
		model.IntegratedClosureState == OSSTrustNetworkValEIntegratedClosureStateActive &&
		model.CanonicalBoundaryState == OSSTrustNetworkValECanonicalBoundaryStateActive &&
		model.EvidenceQualityState == OSSTrustNetworkValEEvidenceQualityStateActive &&
		model.NoOverclaimState == OSSTrustNetworkValENoOverclaimStateActive
}

func EvaluateOSSTrustNetworkValEClosureState(model OSSTrustNetworkValEIntegratedClosure) string {
	switch EvaluateOSSTrustNetworkValEFinalPassRuleState(model) {
	case OSSTrustNetworkValEFinalPassRuleStateActive:
		return OSSTrustNetworkValEClosureStateActive
	case OSSTrustNetworkValEFinalPassRuleStateBlocked:
		return OSSTrustNetworkValEClosureStateBlocked
	case OSSTrustNetworkValEFinalPassRuleStateIncomplete:
		return OSSTrustNetworkValEClosureStateIncomplete
	case OSSTrustNetworkValEFinalPassRuleStateUnknown:
		return OSSTrustNetworkValEClosureStateUnknown
	default:
		return OSSTrustNetworkValEClosureStatePartial
	}
}

func EvaluateOSSTrustNetworkValEPoint9State(model OSSTrustNetworkValEIntegratedClosure) string {
	if EvaluateOSSTrustNetworkValEFinalPassRuleState(model) == OSSTrustNetworkValEFinalPassRuleStateActive {
		return OSSTrustNetworkPoint9StatePass
	}
	return OSSTrustNetworkPoint9StateNotComplete
}

func EvaluateOSSTrustNetworkValEState(model OSSTrustNetworkValEIntegratedClosure) string {
	passRuleState := EvaluateOSSTrustNetworkValEFinalPassRuleState(model)
	if passRuleState == OSSTrustNetworkValEFinalPassRuleStateActive && EvaluateOSSTrustNetworkValEPoint9State(model) == OSSTrustNetworkPoint9StatePass {
		return OSSTrustNetworkValEStatePass
	}
	switch passRuleState {
	case OSSTrustNetworkValEFinalPassRuleStateBlocked:
		return OSSTrustNetworkValEStateBlocked
	case OSSTrustNetworkValEFinalPassRuleStateIncomplete:
		return OSSTrustNetworkValEStateIncomplete
	case OSSTrustNetworkValEFinalPassRuleStateUnknown:
		return OSSTrustNetworkValEStateUnknown
	default:
		return OSSTrustNetworkValEStatePartial
	}
}

func ossTrustNetworkValEBlockingReasons(model OSSTrustNetworkValEIntegratedClosure) []string {
	reasons := []string{}
	if model.DependencyState != OSSTrustNetworkValEDependencyStateActive {
		reasons = append(reasons, "Val D dependency is not fully exact and active.")
	}
	if model.Val0SourceState != OSSTrustNetworkValESourceStateActive {
		reasons = append(reasons, "Val 0 source foundation is not fully exact and active.")
	}
	if model.ValASourceState != OSSTrustNetworkValESourceStateActive {
		reasons = append(reasons, "Val A release trust source is not fully exact and active.")
	}
	if model.ValBSourceState != OSSTrustNetworkValESourceStateActive {
		reasons = append(reasons, "Val B reviewed intelligence source is not fully exact and active.")
	}
	if model.ValCSourceState != OSSTrustNetworkValESourceStateActive {
		reasons = append(reasons, "Val C remediation and visibility source is not fully exact and active.")
	}
	if model.ValDSourceState != OSSTrustNetworkValESourceStateActive {
		reasons = append(reasons, "Val D final readiness source is not fully exact and active.")
	}
	if model.IntegratedClosureState != OSSTrustNetworkValEIntegratedClosureStateActive {
		reasons = append(reasons, "Integrated OSTN closure gate is not fully active.")
	}
	if model.CanonicalBoundaryState != OSSTrustNetworkValECanonicalBoundaryStateActive {
		reasons = append(reasons, "Canonical boundary gate is not fully active.")
	}
	if model.EvidenceQualityState != OSSTrustNetworkValEEvidenceQualityStateActive {
		reasons = append(reasons, "Val E proof refs or evidence quality are not exact and fresh.")
	}
	if model.NoOverclaimState != OSSTrustNetworkValENoOverclaimStateActive {
		reasons = append(reasons, "No-overclaim gate is not fully active.")
	}
	if model.FinalPassRuleState != OSSTrustNetworkValEFinalPassRuleStateActive {
		reasons = append(reasons, "Final Point 9 pass rule remains fail-closed until every Val E gate is active.")
	}
	return ossTrustNetworkValECollectText(reasons)
}

func ComputeOSSTrustNetworkValEClosure(model OSSTrustNetworkValEIntegratedClosure) OSSTrustNetworkValEIntegratedClosure {
	claimedValDSourceState := model.ValDSourceState

	model.Val0SourceState = EvaluateOSSTrustNetworkValEVal0SourceState(model.Val0Source)
	model.ValASourceState = EvaluateOSSTrustNetworkValEValASourceState(model.ValASource)
	model.ValBSourceState = EvaluateOSSTrustNetworkValEValBSourceState(model.ValBSource)
	model.ValCSourceState = EvaluateOSSTrustNetworkValEValCSourceState(model.ValCSource)
	model.ValDSourceState = EvaluateOSSTrustNetworkValEValDSourceState(model.ValDSource)
	if claimedValDSourceState != "" && claimedValDSourceState != model.ValDSourceState {
		model.ValDSourceState = OSSTrustNetworkValESourceStateBlocked
	}
	model.DependencyState = EvaluateOSSTrustNetworkValEDependencyState(model)
	model.IntegratedClosureState = EvaluateOSSTrustNetworkValEIntegratedClosureState(model)
	model.CanonicalBoundaryState = EvaluateOSSTrustNetworkValECanonicalBoundaryState(model)
	model.EvidenceQualityState = EvaluateOSSTrustNetworkValEEvidenceQualityState(model.EvidenceQuality)
	model.NoOverclaimState = EvaluateOSSTrustNetworkValENoOverclaimState(model)
	if ossTrustNetworkValECanPromotePoint9PassReason(model) {
		model.Point9PassReason = OSSTrustNetworkValEPoint9PassReasonAllowed
	} else {
		model.Point9PassReason = OSSTrustNetworkValEPoint9PassReasonBlocked
	}
	model.FinalPassRuleState = EvaluateOSSTrustNetworkValEFinalPassRuleState(model)
	model.ClosureState = EvaluateOSSTrustNetworkValEClosureState(model)
	model.Point9State = EvaluateOSSTrustNetworkValEPoint9State(model)
	model.Point9PassAllowed = model.Point9State == OSSTrustNetworkPoint9StatePass && model.FinalPassRuleState == OSSTrustNetworkValEFinalPassRuleStateActive
	model.CurrentState = EvaluateOSSTrustNetworkValEState(model)
	model.BlockingReasons = ossTrustNetworkValEBlockingReasons(model)

	model.IntegratedClosure.CurrentState = model.IntegratedClosureState
	model.CanonicalBoundary.CurrentState = model.CanonicalBoundaryState
	model.EvidenceQuality.CurrentState = model.EvidenceQualityState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	return model
}
