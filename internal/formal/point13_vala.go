package formal

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
)

const (
	Point13ValAStateActive         = "point13_vala_pilot_execution_contract_active"
	Point13ValAStateBlocked        = "point13_vala_pilot_execution_contract_blocked"
	Point13ValAStateReviewRequired = "point13_vala_pilot_execution_contract_review_required"
	Point13ValAStateIncomplete     = "point13_vala_pilot_execution_contract_incomplete"
)

const (
	point13ValAWaveID                       = "val_a"
	point13ValAPreviousWaveID               = point13Val0WaveID
	point13ValAProjectionDisclaimerBaseline = "projection_only not_canonical_truth point13_vala_pilot_execution_contract"

	point13ValAArtifactClassSupportAttachment = "support_attachment_candidate"
	point13ValAArtifactClassSandboxCandidate  = "sandbox_evidence_candidate"

	point13ValAEvidenceSourceCustomerUpload            = "customer_upload"
	point13ValAEvidenceSourceSupportAttachment         = "support_attachment"
	point13ValAEvidenceSourceSandboxResult             = "sandbox_result"
	point13ValAEvidenceSourceAuditExport               = "audit_export"
	point13ValADisallowedSourceCanonicalization        = "customer_triggered_canonicalization"
	point13ValADisallowedSourceSupportCanonicalization = "support_attachment_canonical_mutation"
	point13ValADisallowedSourceProductionMutation      = "production_mutation"
	point13ValADisallowedSourceConnectorMutation       = "connector_mutation"

	point13ValAPhaseIntake            = "intake"
	point13ValAPhaseBaselineCapture   = "baseline_capture"
	point13ValAPhaseEvidenceReview    = "evidence_review"
	point13ValAPhaseSandboxValidation = "sandbox_validation"
	point13ValAPhaseCustomerReview    = "customer_review"
	point13ValAPhaseExitReview        = "exit_review"
)

type Point13ValADependencySnapshot struct {
	Val0CurrentState                   string                `json:"val0_current_state"`
	Val0DependencyState                string                `json:"val0_dependency_state"`
	Val0PilotReadinessState            string                `json:"val0_pilot_readiness_state"`
	Val0CustomerOnboardingState        string                `json:"val0_customer_onboarding_state"`
	Val0SupportEscalationState         string                `json:"val0_support_escalation_state"`
	Val0OffboardingRetentionState      string                `json:"val0_offboarding_retention_state"`
	Val0NoOverclaimState               string                `json:"val0_no_overclaim_state"`
	Val0AIPilotBoundaryState           string                `json:"val0_ai_pilot_boundary_state"`
	Val0PointID                        string                `json:"val0_point_id"`
	Val0WaveID                         string                `json:"val0_wave_id"`
	Val0DependencyComputedFromUpstream bool                  `json:"val0_dependency_computed_from_upstream"`
	Val0Point13PassSeen                bool                  `json:"val0_point13_pass_seen"`
	Point12CurrentState                string                `json:"point12_current_state"`
	Point12DependencyState             string                `json:"point12_dependency_state"`
	Point12PassClosureManifestState    string                `json:"point12_pass_closure_manifest_state"`
	Point12ReviewerResult              string                `json:"point12_reviewer_result"`
	Point12PassAllowed                 bool                  `json:"point12_pass_allowed"`
	Point12PassToken                   string                `json:"point12_pass_token"`
	Point12TenantScope                 string                `json:"point12_tenant_scope"`
	InheritedAIModelOrRuleVersionRef   string                `json:"inherited_ai_model_or_rule_version_ref"`
	InheritedAIPermissionManifestHash  string                `json:"inherited_ai_permission_manifest_hash"`
	SnapshotFromComputedOutput         bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                []string              `json:"review_prerequisites,omitempty"`
	Val0                               Point13Val0Foundation `json:"val0"`
}

type Point13ValAPilotExecutionContract struct {
	ContractID                           string   `json:"contract_id"`
	PilotOwnerRef                        string   `json:"pilot_owner_ref"`
	CustomerOwnerRef                     string   `json:"customer_owner_ref"`
	TenantScope                          string   `json:"tenant_scope"`
	PilotScopeRef                        string   `json:"pilot_scope_ref"`
	FirstRepoScope                       string   `json:"first_repo_scope"`
	AllowedArtifactClasses               []string `json:"allowed_artifact_classes,omitempty"`
	AllowedEvidenceSources               []string `json:"allowed_evidence_sources,omitempty"`
	DisallowedEvidenceSources            []string `json:"disallowed_evidence_sources,omitempty"`
	EntryCriteriaRefs                    []string `json:"entry_criteria_refs,omitempty"`
	ExitCriteriaRefs                     []string `json:"exit_criteria_refs,omitempty"`
	SuccessMetricsOperationalOnly        []string `json:"success_metrics_operational_only,omitempty"`
	ProductionApprovalExcluded           bool     `json:"production_approval_excluded"`
	DeploymentApprovalExcluded           bool     `json:"deployment_approval_excluded"`
	ComplianceGuaranteeExcluded          bool     `json:"compliance_guarantee_excluded"`
	CustomerSuccessNotProductionApproval bool     `json:"customer_success_not_production_approval"`
	ContractCannotCreatePass             bool     `json:"contract_cannot_create_pass"`
}

type Point13ValACustomerIntakeEvidenceGovernance struct {
	IntakeID                                    string   `json:"intake_id"`
	TenantScope                                 string   `json:"tenant_scope"`
	CustomerArtifactRefs                        []string `json:"customer_artifact_refs,omitempty"`
	CustomerArtifactHashRefs                    []string `json:"customer_artifact_hash_refs,omitempty"`
	CustomerArtifactClassification              string   `json:"customer_artifact_classification"`
	CustodyRef                                  string   `json:"custody_ref"`
	SourceOwnerRef                              string   `json:"source_owner_ref"`
	ConsentOrAuthorityRef                       string   `json:"consent_or_authority_ref"`
	CanonicalizationGovernanceEventRef          string   `json:"canonicalization_governance_event_ref"`
	EvidenceCandidateOnly                       bool     `json:"evidence_candidate_only"`
	CanonicalEvidenceRequiresGovernanceEvent    bool     `json:"canonical_evidence_requires_governance_event"`
	CustomerUploadCannotMutateCanonicalSpine    bool     `json:"customer_upload_cannot_mutate_canonical_spine"`
	SupportAttachmentCannotMutateCanonicalSpine bool     `json:"support_attachment_cannot_mutate_canonical_spine"`
	CustomerArtifactPromotedToCanonical         bool     `json:"customer_artifact_promoted_to_canonical"`
	SupportAttachmentPromotedToCanonical        bool     `json:"support_attachment_promoted_to_canonical"`
	IntakeBindingHash                           string   `json:"intake_binding_hash"`
}

type Point13ValAPilotRunPhase struct {
	PhaseID                     string   `json:"phase_id"`
	PhaseType                   string   `json:"phase_type"`
	PhaseOwnerRef               string   `json:"phase_owner_ref"`
	TenantScope                 string   `json:"tenant_scope"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	AuditEventRef               string   `json:"audit_event_ref"`
	PhaseState                  string   `json:"phase_state"`
	ProductionApprovalAttempted bool     `json:"production_approval_attempted"`
	DeploymentAttempted         bool     `json:"deployment_attempted"`
	PassAttempted               bool     `json:"pass_attempted"`
}

type Point13ValAPilotRunPhaseBoundary struct {
	Phases []Point13ValAPilotRunPhase `json:"phases,omitempty"`
}

type Point13ValACustomerSupportResponsibilityMatrix struct {
	MatrixID                          string   `json:"matrix_id"`
	EscalationOwnerRef                string   `json:"escalation_owner_ref"`
	CustomerOwnerRef                  string   `json:"customer_owner_ref"`
	SupportAccessScope                string   `json:"support_access_scope"`
	TenantScope                       string   `json:"tenant_scope"`
	EvidenceRefs                      []string `json:"evidence_refs,omitempty"`
	AuditEventRefs                    []string `json:"audit_event_refs,omitempty"`
	SupportCanViewCandidateArtifacts  bool     `json:"support_can_view_candidate_artifacts"`
	SupportCanMutateCanonicalEvidence bool     `json:"support_can_mutate_canonical_evidence"`
	SupportCanOverrideCoreDecision    bool     `json:"support_can_override_core_decision"`
	SupportCanApproveProduction       bool     `json:"support_can_approve_production"`
	SupportActionRequiresAuditEvent   bool     `json:"support_action_requires_audit_event"`
}

type Point13ValAPilotExitReviewGate struct {
	ExitReviewID                             string   `json:"exit_review_id"`
	TenantScope                              string   `json:"tenant_scope"`
	ExitCriteriaRefs                         []string `json:"exit_criteria_refs,omitempty"`
	EvidenceRefs                             []string `json:"evidence_refs,omitempty"`
	UnresolvedBlockers                       []string `json:"unresolved_blockers,omitempty"`
	CustomerReviewRef                        string   `json:"customer_review_ref"`
	InternalReviewRef                        string   `json:"internal_review_ref"`
	SupportReviewRef                         string   `json:"support_review_ref"`
	OffboardingOrRetentionPlanRef            string   `json:"offboarding_or_retention_plan_ref"`
	ProductionApprovalRequested              bool     `json:"production_approval_requested"`
	DeploymentApprovalRequested              bool     `json:"deployment_approval_requested"`
	ComplianceClaimRequested                 bool     `json:"compliance_claim_requested"`
	FinalCustomerStatement                   string   `json:"final_customer_statement"`
	PilotExitCanOnlyMarkOperationalReadiness bool     `json:"pilot_exit_can_only_mark_operational_readiness"`
}

type Point13ValAAIAssistedPilotExecutionBoundary struct {
	AIBoundaryID                  string   `json:"ai_boundary_id"`
	TenantScope                   string   `json:"tenant_scope"`
	AIOutputType                  string   `json:"ai_output_type"`
	EvidenceCandidateRef          string   `json:"evidence_candidate_ref"`
	InputEvidenceRefs             []string `json:"input_evidence_refs,omitempty"`
	InputEvidenceHashRefs         []string `json:"input_evidence_hash_refs,omitempty"`
	ModelOrRuleVersionRef         string   `json:"model_or_rule_version_ref"`
	PermissionManifestHash        string   `json:"permission_manifest_hash"`
	SandboxResultRef              string   `json:"sandbox_result_ref"`
	ApprovalRequestRef            string   `json:"approval_request_ref"`
	ReviewerRef                   string   `json:"reviewer_ref"`
	AuditEventRef                 string   `json:"audit_event_ref"`
	AdvisoryOnly                  bool     `json:"advisory_only"`
	EvidenceCandidateOnly         bool     `json:"evidence_candidate_only"`
	PassAllowed                   bool     `json:"pass_allowed"`
	ApprovalGranted               bool     `json:"approval_granted"`
	DeploymentAuthorized          bool     `json:"deployment_authorized"`
	ProductionReadinessClaimed    bool     `json:"production_readiness_claimed"`
	ProductionMutationAllowed     bool     `json:"production_mutation_allowed"`
	CanonicalMutationAllowed      bool     `json:"canonical_mutation_allowed"`
	ExternalAPIAllowed            bool     `json:"external_api_allowed"`
	ExternalAPIGovernanceEventRef string   `json:"external_api_governance_event_ref"`
}

type Point13ValANoOverclaimCustomerWording struct {
	ObservedCustomerFacingTexts          []string `json:"observed_customer_facing_texts,omitempty"`
	ObservedExportFacingTexts            []string `json:"observed_export_facing_texts,omitempty"`
	ObservedSupportFacingTexts           []string `json:"observed_support_facing_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedCustomerFacingWording         []string `json:"allowed_customer_facing_wording,omitempty"`
	BlockedCustomerFacingWording         []string `json:"blocked_customer_facing_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point13ValAFoundation struct {
	CurrentState                          string                                         `json:"current_state"`
	BlockingReasons                       []string                                       `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                   []string                                       `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer                  string                                         `json:"projection_disclaimer"`
	DependencyState                       string                                         `json:"dependency_state"`
	PilotExecutionContractState           string                                         `json:"pilot_execution_contract_state"`
	CustomerIntakeEvidenceGovernanceState string                                         `json:"customer_intake_evidence_governance_state"`
	PilotRunPhaseBoundaryState            string                                         `json:"pilot_run_phase_boundary_state"`
	SupportResponsibilityMatrixState      string                                         `json:"support_responsibility_matrix_state"`
	PilotExitReviewGateState              string                                         `json:"pilot_exit_review_gate_state"`
	AIAssistedPilotExecutionBoundaryState string                                         `json:"ai_assisted_pilot_execution_boundary_state"`
	NoOverclaimState                      string                                         `json:"no_overclaim_state"`
	Dependency                            Point13ValADependencySnapshot                  `json:"dependency"`
	PilotExecutionContract                Point13ValAPilotExecutionContract              `json:"pilot_execution_contract"`
	CustomerIntakeEvidenceGovernance      Point13ValACustomerIntakeEvidenceGovernance    `json:"customer_intake_evidence_governance"`
	PilotRunPhaseBoundary                 Point13ValAPilotRunPhaseBoundary               `json:"pilot_run_phase_boundary"`
	SupportResponsibilityMatrix           Point13ValACustomerSupportResponsibilityMatrix `json:"support_responsibility_matrix"`
	PilotExitReviewGate                   Point13ValAPilotExitReviewGate                 `json:"pilot_exit_review_gate"`
	AIAssistedPilotExecutionBoundary      Point13ValAAIAssistedPilotExecutionBoundary    `json:"ai_assisted_pilot_execution_boundary"`
	NoOverclaimCustomerWording            Point13ValANoOverclaimCustomerWording          `json:"no_overclaim_customer_wording"`
}

func point13ValAStates() []string {
	return []string{
		Point13ValAStateActive,
		Point13ValAStateBlocked,
		Point13ValAStateReviewRequired,
		Point13ValAStateIncomplete,
	}
}

func point13ValAStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValAStates(), value)
}

func point13ValAOwnerRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "pilot_owner_", "customer_owner_", "support_owner_", "retention_owner_", "owner_")
}

func point13ValAContractRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "pilot_execution_contract_", "contract_")
}

func point13ValAIntakeRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "customer_intake_", "intake_")
}

func point13ValACustodyRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "custody_")
}

func point13ValAConsentOrAuthorityRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "consent_", "authority_")
}

func point13ValAPilotScopeRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "pilot_scope_")
}

func point13ValAEntryCriteriaRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "entry_criteria_")
}

func point13ValAExitCriteriaRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "exit_criteria_")
}

func point13ValAPhaseRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "phase_")
}

func point13ValAMatrixRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "support_matrix_", "matrix_")
}

func point13ValAExitReviewRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "exit_review_")
}

func point13ValACustomerReviewRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "customer_review_")
}

func point13ValAInternalReviewRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "internal_review_")
}

func point13ValASupportReviewRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "support_review_")
}

func point13ValAAIBoundaryRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "ai_boundary_")
}

func point13ValAAIEvidenceCandidateRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "evidence_candidate_")
}

func point13ValAReviewerRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "reviewer_", "owner_")
}

func point13ValAAllowedArtifactClasses() []string {
	return []string{
		point13Val0CustomerArtifactCandidateOnly,
		point13ValAArtifactClassSupportAttachment,
		point13ValAArtifactClassSandboxCandidate,
	}
}

func point13ValAAllowedEvidenceSources() []string {
	return []string{
		point13ValAEvidenceSourceCustomerUpload,
		point13ValAEvidenceSourceSupportAttachment,
		point13ValAEvidenceSourceSandboxResult,
		point13ValAEvidenceSourceAuditExport,
	}
}

func point13ValADisallowedEvidenceSources() []string {
	return []string{
		point13ValADisallowedSourceCanonicalization,
		point13ValADisallowedSourceSupportCanonicalization,
		point13ValADisallowedSourceProductionMutation,
		point13ValADisallowedSourceConnectorMutation,
	}
}

func point13ValAPhaseTaxonomy() []string {
	return []string{
		point13ValAPhaseIntake,
		point13ValAPhaseBaselineCapture,
		point13ValAPhaseEvidenceReview,
		point13ValAPhaseSandboxValidation,
		point13ValAPhaseCustomerReview,
		point13ValAPhaseExitReview,
	}
}

func point13ValAAllowedCustomerWording() []string {
	return []string{
		"pilot execution contract",
		"customer intake evidence candidate",
		"advisory ai evidence candidate",
		"operational readiness review",
		"evidence support for customer/auditor review",
	}
}

func point13ValAStringListValid(values []string) bool {
	return point12Val0StringListValid(values, point11Val0IdentityValueValid)
}

func point13ValAArtifactClassListValid(values []string) bool {
	return point12Val0StringListValid(values, func(value string) bool {
		return point11Val0ContainsTrimmed(point13ValAAllowedArtifactClasses(), value)
	})
}

func point13ValAEvidenceSourceListValid(values []string, allowed []string) bool {
	return point12Val0StringListValid(values, func(value string) bool {
		return point11Val0ContainsTrimmed(allowed, value)
	})
}

func point13ValAEntryCriteriaRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValAEntryCriteriaRefValid)
}

func point13ValAExitCriteriaRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValAExitCriteriaRefValid)
}

func point13ValACustomerArtifactRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point12Val0ArtifactRefValid)
}

func point13ValACustomerArtifactHashRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point12Val0EvidenceHashRefValid)
}

func point13ValAOperationalEvidenceRefValid(value string) bool {
	return point12Val0ArtifactRefValid(value) ||
		point12Val0EvidenceRefValid(value) ||
		point13ValAAIEvidenceCandidateRefValid(value)
}

func point13ValAOperationalEvidenceRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValAOperationalEvidenceRefValid)
}

func point13ValAAuditRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point12Val0AuditRefValid)
}

func point13ValAPhaseTypeValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValAPhaseTaxonomy(), value)
}

func point13ValAArtifactHashRefsMatchArtifacts(artifacts, hashes []string) bool {
	if len(artifacts) == 0 || len(artifacts) != len(hashes) {
		return false
	}
	for i := range artifacts {
		artifactToken := strings.TrimSpace(strings.TrimPrefix(artifacts[i], "artifact_"))
		hashToken := strings.TrimSpace(strings.TrimPrefix(hashes[i], "evidence_hash_"))
		if artifactToken == "" || hashToken == "" || artifactToken != hashToken {
			return false
		}
	}
	return true
}

func point13ValAContainsCrossTenantArtifact(values []string) bool {
	replacer := strings.NewReplacer("-", "_", " ", "_", ".", "_", "/", "_")
	for _, value := range values {
		normalized := point11Val0NormalizeText(value)
		canonical := replacer.Replace(normalized)
		for strings.Contains(canonical, "__") {
			canonical = strings.ReplaceAll(canonical, "__", "_")
		}
		for _, blocked := range []string{"cross_tenant", "other_tenant", "all_tenants", "global", "wildcard", "unscoped"} {
			if strings.Contains(normalized, blocked) || strings.Contains(canonical, blocked) {
				return true
			}
		}
	}
	return false
}

func point13ValAComputedBindingHash(parts ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(parts, "\n")))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func point13ValAComputedIntakeBindingHash(model Point13ValACustomerIntakeEvidenceGovernance) string {
	return point13ValAComputedBindingHash(
		strings.TrimSpace(model.IntakeID),
		strings.TrimSpace(model.TenantScope),
		strings.Join(model.CustomerArtifactRefs, ","),
		strings.Join(model.CustomerArtifactHashRefs, ","),
		strings.TrimSpace(model.CustomerArtifactClassification),
		strings.TrimSpace(model.CustodyRef),
		strings.TrimSpace(model.SourceOwnerRef),
		strings.TrimSpace(model.ConsentOrAuthorityRef),
	)
}

func point13ValAVal0PayloadContainsPointPass(val0 Point13Val0Foundation) bool {
	payload, err := json.Marshal(val0)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point13Val0BlockedPoint13PassToken)
}

func point13ValAUpstreamAgentLineage(point12 Point12ValEFoundation) Point12Val0AgentLineageRecord {
	if len(point12.Dependency.Val0.ProvenanceProfile.AgentLineages) == 0 {
		return Point12Val0AgentLineageRecord{}
	}
	return point12.Dependency.Val0.ProvenanceProfile.AgentLineages[0]
}

func point13ValADependencySnapshotFromUpstream(val0 Point13Val0Foundation) Point13ValADependencySnapshot {
	upstreamLineage := point13ValAUpstreamAgentLineage(val0.Dependency.Point12)
	return Point13ValADependencySnapshot{
		Val0CurrentState:                   val0.CurrentState,
		Val0DependencyState:                val0.DependencyState,
		Val0PilotReadinessState:            val0.PilotReadinessState,
		Val0CustomerOnboardingState:        val0.CustomerOnboardingState,
		Val0SupportEscalationState:         val0.SupportEscalationState,
		Val0OffboardingRetentionState:      val0.OffboardingRetentionState,
		Val0NoOverclaimState:               val0.NoOverclaimState,
		Val0AIPilotBoundaryState:           val0.AIPilotBoundaryState,
		Val0PointID:                        point13Val0PointID,
		Val0WaveID:                         point13Val0WaveID,
		Val0DependencyComputedFromUpstream: val0.Dependency.SnapshotFromComputedOutput,
		Val0Point13PassSeen:                point13ValAVal0PayloadContainsPointPass(val0),
		Point12CurrentState:                val0.Dependency.Point12CurrentState,
		Point12DependencyState:             val0.Dependency.Point12DependencyState,
		Point12PassClosureManifestState:    val0.Dependency.Point12PassClosureManifestState,
		Point12ReviewerResult:              val0.Dependency.Point12ReviewerResult,
		Point12PassAllowed:                 val0.Dependency.Point12PassAllowed,
		Point12PassToken:                   val0.Dependency.Point12PassToken,
		Point12TenantScope:                 val0.Dependency.Point12TenantScope,
		InheritedAIModelOrRuleVersionRef:   upstreamLineage.ModelOrRuleVersionRef,
		InheritedAIPermissionManifestHash:  upstreamLineage.PermissionManifestHash,
		SnapshotFromComputedOutput:         true,
		ReviewPrerequisites:                append([]string{}, val0.ReviewPrerequisites...),
		Val0:                               val0,
	}
}

func point13ValADependencySnapshotModel() Point13ValADependencySnapshot {
	return point13ValADependencySnapshotFromUpstream(ComputePoint13Val0Foundation(Point13Val0FoundationModel()))
}

func point13ValADependencyStateAndReasons(model Point13ValADependencySnapshot) (string, []string) {
	reviewReasons := []string{}
	blockedReasons := []string{}
	incompleteReasons := []string{}

	if !model.SnapshotFromComputedOutput || !model.Val0DependencyComputedFromUpstream {
		blockedReasons = append(blockedReasons, "val0_dependency_not_computed_from_upstream")
	}
	if !point13Val0StateValid(model.Val0CurrentState) ||
		!point13Val0StateValid(model.Val0DependencyState) ||
		!point13Val0StateValid(model.Val0PilotReadinessState) ||
		!point13Val0StateValid(model.Val0CustomerOnboardingState) ||
		!point13Val0StateValid(model.Val0SupportEscalationState) ||
		!point13Val0StateValid(model.Val0OffboardingRetentionState) ||
		!point13Val0StateValid(model.Val0NoOverclaimState) ||
		!point13Val0StateValid(model.Val0AIPilotBoundaryState) ||
		strings.TrimSpace(model.Val0PointID) != point13Val0PointID ||
		strings.TrimSpace(model.Val0WaveID) != point13Val0WaveID ||
		!point12ValEStateValid(model.Point12CurrentState) ||
		!point12ValEStateValid(model.Point12DependencyState) ||
		!point12ValEStateValid(model.Point12PassClosureManifestState) ||
		!point12ValEReviewerResultValid(model.Point12ReviewerResult) ||
		!point11Val0ScopeValid(model.Point12TenantScope) ||
		!point12Val0VersionRefValid(model.InheritedAIModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.InheritedAIPermissionManifestHash) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_identity_invalid")
	}
	if strings.TrimSpace(model.Val0CurrentState) != strings.TrimSpace(model.Val0.CurrentState) ||
		strings.TrimSpace(model.Val0DependencyState) != strings.TrimSpace(model.Val0.DependencyState) ||
		strings.TrimSpace(model.Val0PilotReadinessState) != strings.TrimSpace(model.Val0.PilotReadinessState) ||
		strings.TrimSpace(model.Val0CustomerOnboardingState) != strings.TrimSpace(model.Val0.CustomerOnboardingState) ||
		strings.TrimSpace(model.Val0SupportEscalationState) != strings.TrimSpace(model.Val0.SupportEscalationState) ||
		strings.TrimSpace(model.Val0OffboardingRetentionState) != strings.TrimSpace(model.Val0.OffboardingRetentionState) ||
		strings.TrimSpace(model.Val0NoOverclaimState) != strings.TrimSpace(model.Val0.NoOverclaimState) ||
		strings.TrimSpace(model.Val0AIPilotBoundaryState) != strings.TrimSpace(model.Val0.AIPilotBoundaryState) ||
		model.Val0DependencyComputedFromUpstream != model.Val0.Dependency.SnapshotFromComputedOutput ||
		strings.TrimSpace(model.Point12CurrentState) != strings.TrimSpace(model.Val0.Dependency.Point12CurrentState) ||
		strings.TrimSpace(model.Point12DependencyState) != strings.TrimSpace(model.Val0.Dependency.Point12DependencyState) ||
		strings.TrimSpace(model.Point12PassClosureManifestState) != strings.TrimSpace(model.Val0.Dependency.Point12PassClosureManifestState) ||
		strings.TrimSpace(model.Point12ReviewerResult) != strings.TrimSpace(model.Val0.Dependency.Point12ReviewerResult) ||
		model.Point12PassAllowed != model.Val0.Dependency.Point12PassAllowed ||
		strings.TrimSpace(model.Point12PassToken) != strings.TrimSpace(model.Val0.Dependency.Point12PassToken) ||
		strings.TrimSpace(model.Point12TenantScope) != strings.TrimSpace(model.Val0.Dependency.Point12TenantScope) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_binding_mismatch")
	}
	if model.Val0Point13PassSeen {
		blockedReasons = append(blockedReasons, "val0_point13_pass_seen")
	}
	switch strings.TrimSpace(model.Val0CurrentState) {
	case Point13Val0StateBlocked:
		blockedReasons = append(blockedReasons, "val0_blocked")
	case Point13Val0StateReviewRequired:
		reviewReasons = append(reviewReasons, "val0_review_required")
	case Point13Val0StateIncomplete:
		incompleteReasons = append(incompleteReasons, "val0_incomplete")
	}
	for _, state := range []string{
		model.Val0DependencyState,
		model.Val0PilotReadinessState,
		model.Val0CustomerOnboardingState,
		model.Val0SupportEscalationState,
		model.Val0OffboardingRetentionState,
		model.Val0NoOverclaimState,
		model.Val0AIPilotBoundaryState,
	} {
		switch strings.TrimSpace(state) {
		case Point13Val0StateBlocked:
			blockedReasons = append(blockedReasons, "val0_component_blocked")
		case Point13Val0StateReviewRequired:
			reviewReasons = append(reviewReasons, "val0_component_review_required")
		case Point13Val0StateIncomplete:
			incompleteReasons = append(incompleteReasons, "val0_component_incomplete")
		}
	}
	switch strings.TrimSpace(model.Point12CurrentState) {
	case Point12ValEStateBlocked:
		blockedReasons = append(blockedReasons, "point12_inherited_blocked")
	case Point12ValEStateReviewRequired:
		reviewReasons = append(reviewReasons, "point12_inherited_review_required")
	case Point12ValEStateIncomplete:
		incompleteReasons = append(incompleteReasons, "point12_inherited_incomplete")
	}
	if strings.TrimSpace(model.Point12CurrentState) == Point12ValEStatePassConfirmed &&
		(strings.TrimSpace(model.Point12DependencyState) != Point12ValEStateActive ||
			strings.TrimSpace(model.Point12PassClosureManifestState) != Point12ValEStateActive ||
			strings.TrimSpace(model.Point12ReviewerResult) != point12ValEReviewerResultPassConfirmed ||
			!model.Point12PassAllowed ||
			strings.TrimSpace(model.Point12PassToken) != point12ValEPoint12PassToken) {
		blockedReasons = append(blockedReasons, "point12_inherited_not_pass_confirmed")
	}
	if len(blockedReasons) > 0 {
		return Point13ValAStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point13ValAStateReviewRequired, reviewReasons
	}
	if len(incompleteReasons) > 0 {
		return Point13ValAStateIncomplete, incompleteReasons
	}
	return Point13ValAStateActive, nil
}

func EvaluatePoint13ValAPilotExecutionContractState(model Point13ValAPilotExecutionContract, dependency Point13ValADependencySnapshot) string {
	if !point13ValAContractRefValid(model.ContractID) ||
		!point13ValAOwnerRefValid(model.PilotOwnerRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValAPilotScopeRefValid(model.PilotScopeRef) ||
		!point11Val0IdentityValueValid(model.FirstRepoScope) ||
		!point13ValAArtifactClassListValid(model.AllowedArtifactClasses) ||
		!point13ValAEvidenceSourceListValid(model.AllowedEvidenceSources, point13ValAAllowedEvidenceSources()) ||
		!point13ValAEvidenceSourceListValid(model.DisallowedEvidenceSources, point13ValADisallowedEvidenceSources()) ||
		!point13ValAEntryCriteriaRefsValid(model.EntryCriteriaRefs) ||
		!point13ValAExitCriteriaRefsValid(model.ExitCriteriaRefs) ||
		!point13ValAStringListValid(model.SuccessMetricsOperationalOnly) ||
		!model.ProductionApprovalExcluded ||
		!model.DeploymentApprovalExcluded ||
		!model.ComplianceGuaranteeExcluded ||
		!model.CustomerSuccessNotProductionApproval ||
		!model.ContractCannotCreatePass {
		return Point13ValAStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.Point12TenantScope) ||
		strings.TrimSpace(model.PilotOwnerRef) != strings.TrimSpace(dependency.Val0.PilotReadiness.PilotOwnerRef) ||
		strings.TrimSpace(model.FirstRepoScope) != strings.TrimSpace(dependency.Val0.PilotReadiness.FirstRepoScope) {
		return Point13ValAStateBlocked
	}
	for _, disallowed := range model.DisallowedEvidenceSources {
		if point11Val0ContainsTrimmed(model.AllowedEvidenceSources, disallowed) {
			return Point13ValAStateBlocked
		}
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.SuccessMetricsOperationalOnly, " "),
		strings.Join(model.EntryCriteriaRefs, " "),
		strings.Join(model.ExitCriteriaRefs, " "),
	) {
		return Point13ValAStateBlocked
	}
	return Point13ValAStateActive
}

func EvaluatePoint13ValACustomerIntakeEvidenceGovernanceState(model Point13ValACustomerIntakeEvidenceGovernance, dependency Point13ValADependencySnapshot, contract Point13ValAPilotExecutionContract) string {
	if !point13ValAIntakeRefValid(model.IntakeID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValACustomerArtifactRefsValid(model.CustomerArtifactRefs) ||
		!point13ValACustomerArtifactHashRefsValid(model.CustomerArtifactHashRefs) ||
		!point13ValAArtifactHashRefsMatchArtifacts(model.CustomerArtifactRefs, model.CustomerArtifactHashRefs) ||
		!point11Val0ContainsTrimmed(contract.AllowedArtifactClasses, model.CustomerArtifactClassification) ||
		!point13ValACustodyRefValid(model.CustodyRef) ||
		!point13ValAOwnerRefValid(model.SourceOwnerRef) ||
		!point13ValAConsentOrAuthorityRefValid(model.ConsentOrAuthorityRef) ||
		!model.EvidenceCandidateOnly ||
		!model.CanonicalEvidenceRequiresGovernanceEvent ||
		!model.CustomerUploadCannotMutateCanonicalSpine ||
		!model.SupportAttachmentCannotMutateCanonicalSpine ||
		strings.TrimSpace(model.IntakeBindingHash) != point13ValAComputedIntakeBindingHash(model) {
		return Point13ValAStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(contract.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.Point12TenantScope) ||
		strings.TrimSpace(model.SourceOwnerRef) != strings.TrimSpace(contract.CustomerOwnerRef) {
		return Point13ValAStateBlocked
	}
	if point13ValAContainsCrossTenantArtifact(model.CustomerArtifactRefs) {
		return Point13ValAStateBlocked
	}
	if model.CustomerArtifactPromotedToCanonical || model.SupportAttachmentPromotedToCanonical {
		if !point12Val0GovernanceEventRefValid(model.CanonicalizationGovernanceEventRef) {
			return Point13ValAStateBlocked
		}
		return Point13ValAStateBlocked
	}
	return Point13ValAStateActive
}

func EvaluatePoint13ValAPilotRunPhaseBoundaryState(model Point13ValAPilotRunPhaseBoundary, dependency Point13ValADependencySnapshot, contract Point13ValAPilotExecutionContract, intake Point13ValACustomerIntakeEvidenceGovernance) string {
	if len(model.Phases) == 0 {
		return Point13ValAStateBlocked
	}
	for _, phase := range model.Phases {
		if !point13ValAPhaseRefValid(phase.PhaseID) ||
			!point13ValAPhaseTypeValid(phase.PhaseType) ||
			!point13ValAOwnerRefValid(phase.PhaseOwnerRef) ||
			!point11Val0ScopeValid(phase.TenantScope) ||
			!point13ValAOperationalEvidenceRefsValid(phase.EvidenceRefs) ||
			!point12Val0AuditRefValid(phase.AuditEventRef) ||
			!point13ValAStateValid(phase.PhaseState) ||
			phase.ProductionApprovalAttempted ||
			phase.DeploymentAttempted ||
			phase.PassAttempted {
			return Point13ValAStateBlocked
		}
		if strings.TrimSpace(phase.TenantScope) != strings.TrimSpace(contract.TenantScope) ||
			strings.TrimSpace(phase.TenantScope) != strings.TrimSpace(dependency.Point12TenantScope) ||
			!point12Val0ExactStringSetMatch(phase.EvidenceRefs, intake.CustomerArtifactRefs) {
			return Point13ValAStateBlocked
		}
		if strings.TrimSpace(phase.PhaseState) != Point13ValAStateActive {
			if strings.TrimSpace(phase.PhaseState) == Point13ValAStateReviewRequired {
				return Point13ValAStateReviewRequired
			}
			if strings.TrimSpace(phase.PhaseState) == Point13ValAStateIncomplete {
				return Point13ValAStateIncomplete
			}
			return Point13ValAStateBlocked
		}
	}
	return Point13ValAStateActive
}

func EvaluatePoint13ValASupportResponsibilityMatrixState(model Point13ValACustomerSupportResponsibilityMatrix, dependency Point13ValADependencySnapshot, contract Point13ValAPilotExecutionContract, intake Point13ValACustomerIntakeEvidenceGovernance) string {
	if !point13ValAMatrixRefValid(model.MatrixID) ||
		!point13ValAOwnerRefValid(model.EscalationOwnerRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point11Val0IdentityValueValid(model.SupportAccessScope) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValAOperationalEvidenceRefsValid(model.EvidenceRefs) ||
		!point13ValAAuditRefsValid(model.AuditEventRefs) ||
		!model.SupportCanViewCandidateArtifacts ||
		model.SupportCanMutateCanonicalEvidence ||
		model.SupportCanOverrideCoreDecision ||
		model.SupportCanApproveProduction ||
		!model.SupportActionRequiresAuditEvent {
		return Point13ValAStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(contract.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.Point12TenantScope) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(contract.CustomerOwnerRef) ||
		strings.TrimSpace(model.EscalationOwnerRef) != strings.TrimSpace(dependency.Val0.SupportEscalationBoundary.EscalationOwnerRef) ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, intake.CustomerArtifactRefs) {
		return Point13ValAStateBlocked
	}
	return Point13ValAStateActive
}

func EvaluatePoint13ValAPilotExitReviewGateState(model Point13ValAPilotExitReviewGate, dependency Point13ValADependencySnapshot, contract Point13ValAPilotExecutionContract, intake Point13ValACustomerIntakeEvidenceGovernance, aiBoundary Point13ValAAIAssistedPilotExecutionBoundary) string {
	if !point13ValAExitReviewRefValid(model.ExitReviewID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValAExitCriteriaRefsValid(model.ExitCriteriaRefs) ||
		!point13ValAOperationalEvidenceRefsValid(model.EvidenceRefs) ||
		!point13ValACustomerReviewRefValid(model.CustomerReviewRef) ||
		!point13ValAInternalReviewRefValid(model.InternalReviewRef) ||
		!point13ValASupportReviewRefValid(model.SupportReviewRef) ||
		!point13Val0OperationalRefValid(model.OffboardingOrRetentionPlanRef, "disposal_path_") ||
		!model.PilotExitCanOnlyMarkOperationalReadiness ||
		strings.TrimSpace(model.FinalCustomerStatement) == "" {
		return Point13ValAStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(contract.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.Point12TenantScope) ||
		!point12Val0ExactStringSetMatch(model.ExitCriteriaRefs, contract.ExitCriteriaRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, intake.CustomerArtifactRefs) ||
		strings.TrimSpace(model.OffboardingOrRetentionPlanRef) != strings.TrimSpace(dependency.Val0.OffboardingRetentionBoundary.DisposalPathRef) {
		return Point13ValAStateBlocked
	}
	if len(model.UnresolvedBlockers) > 0 ||
		model.ProductionApprovalRequested ||
		model.DeploymentApprovalRequested ||
		model.ComplianceClaimRequested ||
		point13Val0ContainsForbiddenClaim(model.FinalCustomerStatement) {
		return Point13ValAStateBlocked
	}
	if len(model.EvidenceRefs) == 1 && strings.TrimSpace(model.EvidenceRefs[0]) == strings.TrimSpace(aiBoundary.EvidenceCandidateRef) {
		return Point13ValAStateBlocked
	}
	return Point13ValAStateActive
}

func EvaluatePoint13ValAAIAssistedPilotExecutionBoundaryState(model Point13ValAAIAssistedPilotExecutionBoundary, dependency Point13ValADependencySnapshot, contract Point13ValAPilotExecutionContract, intake Point13ValACustomerIntakeEvidenceGovernance, support Point13ValACustomerSupportResponsibilityMatrix) string {
	if point11Val0ContainsTrimmed(point12Val0BlockedAIEvidenceCandidateTypes(), model.AIOutputType) {
		return Point13ValAStateBlocked
	}
	expectedSupportAuditRef := ""
	if len(support.AuditEventRefs) > 0 {
		expectedSupportAuditRef = strings.TrimSpace(support.AuditEventRefs[0])
	}
	if !point13ValAAIBoundaryRefValid(model.AIBoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0AIEvidenceCandidateTypeValid(model.AIOutputType) ||
		!point13ValAAIEvidenceCandidateRefValid(model.EvidenceCandidateRef) ||
		!point13ValAOperationalEvidenceRefsValid(model.InputEvidenceRefs) ||
		!point13ValACustomerArtifactHashRefsValid(model.InputEvidenceHashRefs) ||
		!point13ValAArtifactHashRefsMatchArtifacts(model.InputEvidenceRefs, model.InputEvidenceHashRefs) ||
		!point12Val0VersionRefValid(model.ModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.PermissionManifestHash) ||
		(model.SandboxResultRef != "" && !point13Val0OperationalRefValid(model.SandboxResultRef, "sandbox_result_")) ||
		(model.ApprovalRequestRef != "" && !point13Val0OperationalRefValid(model.ApprovalRequestRef, "approval_request_")) ||
		(model.ReviewerRef != "" && !point13ValAReviewerRefValid(model.ReviewerRef)) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		!model.AdvisoryOnly ||
		!model.EvidenceCandidateOnly ||
		model.PassAllowed ||
		model.ApprovalGranted ||
		model.DeploymentAuthorized ||
		model.ProductionReadinessClaimed ||
		model.ProductionMutationAllowed ||
		model.CanonicalMutationAllowed ||
		expectedSupportAuditRef == "" {
		return Point13ValAStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(contract.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.Point12TenantScope) ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceRefs, intake.CustomerArtifactRefs) ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceHashRefs, intake.CustomerArtifactHashRefs) ||
		strings.TrimSpace(model.ModelOrRuleVersionRef) != strings.TrimSpace(dependency.InheritedAIModelOrRuleVersionRef) ||
		strings.TrimSpace(model.PermissionManifestHash) != strings.TrimSpace(dependency.InheritedAIPermissionManifestHash) ||
		strings.TrimSpace(model.AuditEventRef) != expectedSupportAuditRef {
		return Point13ValAStateBlocked
	}
	if model.ExternalAPIAllowed {
		if !point12Val0GovernanceEventRefValid(model.ExternalAPIGovernanceEventRef) {
			return Point13ValAStateBlocked
		}
		return Point13ValAStateReviewRequired
	}
	return Point13ValAStateActive
}

func EvaluatePoint13ValANoOverclaimCustomerWordingState(model Point13ValANoOverclaimCustomerWording) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point13ValAStringListValid(model.AllowedCustomerFacingWording) ||
		!point13ValAStringListValid(model.BlockedCustomerFacingWording) {
		return Point13ValAStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.ObservedCustomerFacingTexts, " "),
		strings.Join(model.ObservedExportFacingTexts, " "),
		strings.Join(model.ObservedSupportFacingTexts, " "),
	) {
		return Point13ValAStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(strings.Join(model.InternalDiagnosticTexts, " ")) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point13ValAStateBlocked
	}
	return Point13ValAStateActive
}

type point13ValAComponentState struct {
	Name  string
	State string
}

func point13ValAComponentStates(model Point13ValAFoundation) []point13ValAComponentState {
	return []point13ValAComponentState{
		{Name: "dependency", State: model.DependencyState},
		{Name: "pilot_execution_contract", State: model.PilotExecutionContractState},
		{Name: "customer_intake_evidence_governance", State: model.CustomerIntakeEvidenceGovernanceState},
		{Name: "pilot_run_phase_boundary", State: model.PilotRunPhaseBoundaryState},
		{Name: "support_responsibility_matrix", State: model.SupportResponsibilityMatrixState},
		{Name: "pilot_exit_review_gate", State: model.PilotExitReviewGateState},
		{Name: "ai_assisted_pilot_execution", State: model.AIAssistedPilotExecutionBoundaryState},
		{Name: "no_overclaim", State: model.NoOverclaimState},
	}
}

func point13ValABlockingReasons(model Point13ValAFoundation) []string {
	reasons := []string{}
	for _, component := range point13ValAComponentStates(model) {
		if component.State == Point13ValAStateBlocked || component.State == Point13ValAStateIncomplete {
			reasons = append(reasons, component.Name+":"+component.State)
		}
	}
	return reasons
}

func EvaluatePoint13ValAState(model Point13ValAFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValAStateBlocked
	}

	hasReviewRequired := false
	hasIncomplete := false
	for _, component := range point13ValAComponentStates(model) {
		switch component.State {
		case Point13ValAStateBlocked:
			return Point13ValAStateBlocked
		case Point13ValAStateReviewRequired:
			hasReviewRequired = true
		case Point13ValAStateIncomplete:
			hasIncomplete = true
		}
	}

	if hasReviewRequired {
		return Point13ValAStateReviewRequired
	}
	if hasIncomplete {
		return Point13ValAStateIncomplete
	}
	return Point13ValAStateActive
}

func point13ValADefaultPilotRunPhases(contract Point13ValAPilotExecutionContract, intake Point13ValACustomerIntakeEvidenceGovernance) []Point13ValAPilotRunPhase {
	return []Point13ValAPilotRunPhase{
		{
			PhaseID:       "phase_point13_vala_intake_001",
			PhaseType:     point13ValAPhaseIntake,
			PhaseOwnerRef: contract.CustomerOwnerRef,
			TenantScope:   contract.TenantScope,
			EvidenceRefs:  append([]string{}, intake.CustomerArtifactRefs...),
			AuditEventRef: "audit_point13_vala_phase_intake_001",
			PhaseState:    Point13ValAStateActive,
		},
		{
			PhaseID:       "phase_point13_vala_baseline_001",
			PhaseType:     point13ValAPhaseBaselineCapture,
			PhaseOwnerRef: contract.PilotOwnerRef,
			TenantScope:   contract.TenantScope,
			EvidenceRefs:  append([]string{}, intake.CustomerArtifactRefs...),
			AuditEventRef: "audit_point13_vala_phase_baseline_001",
			PhaseState:    Point13ValAStateActive,
		},
		{
			PhaseID:       "phase_point13_vala_evidence_review_001",
			PhaseType:     point13ValAPhaseEvidenceReview,
			PhaseOwnerRef: contract.PilotOwnerRef,
			TenantScope:   contract.TenantScope,
			EvidenceRefs:  append([]string{}, intake.CustomerArtifactRefs...),
			AuditEventRef: "audit_point13_vala_phase_evidence_review_001",
			PhaseState:    Point13ValAStateActive,
		},
		{
			PhaseID:       "phase_point13_vala_sandbox_validation_001",
			PhaseType:     point13ValAPhaseSandboxValidation,
			PhaseOwnerRef: contract.PilotOwnerRef,
			TenantScope:   contract.TenantScope,
			EvidenceRefs:  append([]string{}, intake.CustomerArtifactRefs...),
			AuditEventRef: "audit_point13_vala_phase_sandbox_001",
			PhaseState:    Point13ValAStateActive,
		},
		{
			PhaseID:       "phase_point13_vala_customer_review_001",
			PhaseType:     point13ValAPhaseCustomerReview,
			PhaseOwnerRef: contract.CustomerOwnerRef,
			TenantScope:   contract.TenantScope,
			EvidenceRefs:  append([]string{}, intake.CustomerArtifactRefs...),
			AuditEventRef: "audit_point13_vala_phase_customer_review_001",
			PhaseState:    Point13ValAStateActive,
		},
		{
			PhaseID:       "phase_point13_vala_exit_review_001",
			PhaseType:     point13ValAPhaseExitReview,
			PhaseOwnerRef: contract.PilotOwnerRef,
			TenantScope:   contract.TenantScope,
			EvidenceRefs:  append([]string{}, intake.CustomerArtifactRefs...),
			AuditEventRef: "audit_point13_vala_phase_exit_review_001",
			PhaseState:    Point13ValAStateActive,
		},
	}
}

func Point13ValAFoundationModel() Point13ValAFoundation {
	disclaimer := point13ValAProjectionDisclaimerBaseline
	dependency := point13ValADependencySnapshotModel()
	tenantScope := dependency.Point12TenantScope

	contract := Point13ValAPilotExecutionContract{
		ContractID:                           "pilot_execution_contract_point13_vala_001",
		PilotOwnerRef:                        dependency.Val0.PilotReadiness.PilotOwnerRef,
		CustomerOwnerRef:                     "customer_owner_point13_vala_001",
		TenantScope:                          tenantScope,
		PilotScopeRef:                        "pilot_scope_point13_vala_001",
		FirstRepoScope:                       dependency.Val0.PilotReadiness.FirstRepoScope,
		AllowedArtifactClasses:               []string{point13Val0CustomerArtifactCandidateOnly, point13ValAArtifactClassSupportAttachment},
		AllowedEvidenceSources:               []string{point13ValAEvidenceSourceCustomerUpload, point13ValAEvidenceSourceSupportAttachment, point13ValAEvidenceSourceSandboxResult, point13ValAEvidenceSourceAuditExport},
		DisallowedEvidenceSources:            point13ValADisallowedEvidenceSources(),
		EntryCriteriaRefs:                    []string{"entry_criteria_point13_vala_001", "entry_criteria_point13_vala_002"},
		ExitCriteriaRefs:                     []string{"exit_criteria_point13_vala_001", "exit_criteria_point13_vala_002"},
		SuccessMetricsOperationalOnly:        []string{"pilot execution contract", "operational readiness review"},
		ProductionApprovalExcluded:           true,
		DeploymentApprovalExcluded:           true,
		ComplianceGuaranteeExcluded:          true,
		CustomerSuccessNotProductionApproval: true,
		ContractCannotCreatePass:             true,
	}

	intake := Point13ValACustomerIntakeEvidenceGovernance{
		IntakeID:                                    "customer_intake_point13_vala_001",
		TenantScope:                                 tenantScope,
		CustomerArtifactRefs:                        []string{"artifact_point13_vala_customer_candidate_001", "artifact_point13_vala_customer_candidate_002"},
		CustomerArtifactHashRefs:                    []string{"evidence_hash_point13_vala_customer_candidate_001", "evidence_hash_point13_vala_customer_candidate_002"},
		CustomerArtifactClassification:              point13Val0CustomerArtifactCandidateOnly,
		CustodyRef:                                  "custody_point13_vala_001",
		SourceOwnerRef:                              contract.CustomerOwnerRef,
		ConsentOrAuthorityRef:                       "consent_point13_vala_customer_artifact_001",
		EvidenceCandidateOnly:                       true,
		CanonicalEvidenceRequiresGovernanceEvent:    true,
		CustomerUploadCannotMutateCanonicalSpine:    true,
		SupportAttachmentCannotMutateCanonicalSpine: true,
	}
	intake.IntakeBindingHash = point13ValAComputedIntakeBindingHash(intake)

	support := Point13ValACustomerSupportResponsibilityMatrix{
		MatrixID:                         "support_matrix_point13_vala_001",
		EscalationOwnerRef:               dependency.Val0.SupportEscalationBoundary.EscalationOwnerRef,
		CustomerOwnerRef:                 contract.CustomerOwnerRef,
		SupportAccessScope:               "tenant_scoped_candidate_artifact_view",
		TenantScope:                      tenantScope,
		EvidenceRefs:                     append([]string{}, intake.CustomerArtifactRefs...),
		AuditEventRefs:                   []string{"audit_point13_vala_support_matrix_001", "audit_point13_vala_support_matrix_002"},
		SupportCanViewCandidateArtifacts: true,
		SupportActionRequiresAuditEvent:  true,
	}

	return Point13ValAFoundation{
		CurrentState:                          Point13ValAStateActive,
		ProjectionDisclaimer:                  disclaimer,
		DependencyState:                       Point13ValAStateActive,
		PilotExecutionContractState:           Point13ValAStateActive,
		CustomerIntakeEvidenceGovernanceState: Point13ValAStateActive,
		PilotRunPhaseBoundaryState:            Point13ValAStateActive,
		SupportResponsibilityMatrixState:      Point13ValAStateActive,
		PilotExitReviewGateState:              Point13ValAStateActive,
		AIAssistedPilotExecutionBoundaryState: Point13ValAStateActive,
		NoOverclaimState:                      Point13ValAStateActive,
		Dependency:                            dependency,
		PilotExecutionContract:                contract,
		CustomerIntakeEvidenceGovernance:      intake,
		PilotRunPhaseBoundary: Point13ValAPilotRunPhaseBoundary{
			Phases: point13ValADefaultPilotRunPhases(contract, intake),
		},
		SupportResponsibilityMatrix: support,
		PilotExitReviewGate: Point13ValAPilotExitReviewGate{
			ExitReviewID:                             "exit_review_point13_vala_001",
			TenantScope:                              tenantScope,
			ExitCriteriaRefs:                         append([]string{}, contract.ExitCriteriaRefs...),
			EvidenceRefs:                             append([]string{}, intake.CustomerArtifactRefs...),
			CustomerReviewRef:                        "customer_review_point13_vala_001",
			InternalReviewRef:                        "internal_review_point13_vala_001",
			SupportReviewRef:                         "support_review_point13_vala_001",
			OffboardingOrRetentionPlanRef:            dependency.Val0.OffboardingRetentionBoundary.DisposalPathRef,
			FinalCustomerStatement:                   "operational readiness review complete",
			PilotExitCanOnlyMarkOperationalReadiness: true,
		},
		AIAssistedPilotExecutionBoundary: Point13ValAAIAssistedPilotExecutionBoundary{
			AIBoundaryID:           "ai_boundary_point13_vala_001",
			TenantScope:            tenantScope,
			AIOutputType:           "AI_RECOMMENDATION",
			EvidenceCandidateRef:   "evidence_candidate_point13_vala_001",
			InputEvidenceRefs:      append([]string{}, intake.CustomerArtifactRefs...),
			InputEvidenceHashRefs:  append([]string{}, intake.CustomerArtifactHashRefs...),
			ModelOrRuleVersionRef:  dependency.InheritedAIModelOrRuleVersionRef,
			PermissionManifestHash: dependency.InheritedAIPermissionManifestHash,
			ReviewerRef:            "reviewer_point13_vala_internal_001",
			AuditEventRef:          support.AuditEventRefs[0],
			AdvisoryOnly:           true,
			EvidenceCandidateOnly:  true,
		},
		NoOverclaimCustomerWording: Point13ValANoOverclaimCustomerWording{
			ObservedCustomerFacingTexts:          []string{"pilot execution contract", "customer intake evidence candidate"},
			ObservedExportFacingTexts:            []string{"evidence support for customer/auditor review"},
			ObservedSupportFacingTexts:           []string{"operational readiness review"},
			InternalDiagnosticTexts:              []string{"blocked wording remains denylisted internally"},
			InternalDiagnosticsClassifiedBlocked: true,
			AllowedCustomerFacingWording:         point13ValAAllowedCustomerWording(),
			BlockedCustomerFacingWording:         point13Val0ForbiddenClaims(),
			ProjectionDisclaimer:                 disclaimer,
		},
	}
}

func ComputePoint13ValAFoundation(model Point13ValAFoundation) Point13ValAFoundation {
	dependencyState, dependencyReasons := point13ValADependencyStateAndReasons(model.Dependency)
	model.DependencyState = dependencyState
	model.PilotExecutionContractState = EvaluatePoint13ValAPilotExecutionContractState(model.PilotExecutionContract, model.Dependency)
	model.CustomerIntakeEvidenceGovernanceState = EvaluatePoint13ValACustomerIntakeEvidenceGovernanceState(model.CustomerIntakeEvidenceGovernance, model.Dependency, model.PilotExecutionContract)
	model.PilotRunPhaseBoundaryState = EvaluatePoint13ValAPilotRunPhaseBoundaryState(model.PilotRunPhaseBoundary, model.Dependency, model.PilotExecutionContract, model.CustomerIntakeEvidenceGovernance)
	model.SupportResponsibilityMatrixState = EvaluatePoint13ValASupportResponsibilityMatrixState(model.SupportResponsibilityMatrix, model.Dependency, model.PilotExecutionContract, model.CustomerIntakeEvidenceGovernance)
	model.PilotExitReviewGateState = EvaluatePoint13ValAPilotExitReviewGateState(model.PilotExitReviewGate, model.Dependency, model.PilotExecutionContract, model.CustomerIntakeEvidenceGovernance, model.AIAssistedPilotExecutionBoundary)
	model.AIAssistedPilotExecutionBoundaryState = EvaluatePoint13ValAAIAssistedPilotExecutionBoundaryState(model.AIAssistedPilotExecutionBoundary, model.Dependency, model.PilotExecutionContract, model.CustomerIntakeEvidenceGovernance, model.SupportResponsibilityMatrix)
	model.NoOverclaimState = EvaluatePoint13ValANoOverclaimCustomerWordingState(model.NoOverclaimCustomerWording)
	model.CurrentState = EvaluatePoint13ValAState(model)
	model.BlockingReasons = point13ValABlockingReasons(model)
	model.ReviewPrerequisites = nil
	if model.DependencyState == Point13ValAStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, dependencyReasons...)
	}
	if model.AIAssistedPilotExecutionBoundaryState == Point13ValAStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "ai_assisted_boundary_requires_governance_review")
	}
	return model
}
