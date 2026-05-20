package formal

import "strings"

const (
	Point13Val0StateActive         = "point13_val0_customer_operationalization_active"
	Point13Val0StateBlocked        = "point13_val0_customer_operationalization_blocked"
	Point13Val0StateReviewRequired = "point13_val0_customer_operationalization_review_required"
	Point13Val0StateIncomplete     = "point13_val0_customer_operationalization_incomplete"
)

const (
	point13Val0PointID                      = "point_13"
	point13Val0WaveID                       = "val_0"
	point13Val0ProjectionDisclaimerBaseline = "projection_only not_canonical_truth point13_val0_customer_operationalization_foundation"
	point13Val0BlockedPoint13PassToken      = "point_13_pass"

	point13Val0CustomerArtifactCandidateOnly = "candidate_support_material"
	point13Val0SupportSeverityCritical       = "sev1_critical"
	point13Val0SupportSeverityHigh           = "sev2_high"
	point13Val0SupportSeverityModerate       = "sev3_moderate"
	point13Val0SupportSeverityLow            = "sev4_low"
)

type Point13Val0DependencySnapshot struct {
	Point12CurrentState             string                `json:"point12_current_state"`
	Point12DependencyState          string                `json:"point12_dependency_state"`
	Point12PassClosureManifestState string                `json:"point12_pass_closure_manifest_state"`
	Point12ReviewerResult           string                `json:"point12_reviewer_result"`
	Point12PassAllowed              bool                  `json:"point12_pass_allowed"`
	Point12PassToken                string                `json:"point12_pass_token"`
	Point12ClosureManifestRef       string                `json:"point12_closure_manifest_ref"`
	Point12DependencySnapshotRef    string                `json:"point12_dependency_snapshot_ref"`
	Point12ProofPackID              string                `json:"point12_proof_pack_id"`
	Point12TenantScope              string                `json:"point12_tenant_scope"`
	AIGovernanceBackfillVerified    bool                  `json:"ai_governance_backfill_verified"`
	AIGovernanceBackfillMerged      bool                  `json:"ai_governance_backfill_merged"`
	GitHubCIVerified                bool                  `json:"github_ci_verified"`
	BackfillPullRequestRef          string                `json:"backfill_pull_request_ref"`
	BackfillHeadCommitRef           string                `json:"backfill_head_commit_ref"`
	BackfillMergeCommitRef          string                `json:"backfill_merge_commit_ref"`
	BackfillTestWorkflowRunRef      string                `json:"backfill_test_workflow_run_ref"`
	BackfillLintWorkflowRunRef      string                `json:"backfill_lint_workflow_run_ref"`
	SnapshotFromComputedOutput      bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites             []string              `json:"review_prerequisites,omitempty"`
	Point12                         Point12ValEFoundation `json:"point12"`
}

type Point13Val0PilotReadinessFoundation struct {
	PilotEntryCriteria                        []string `json:"pilot_entry_criteria,omitempty"`
	PilotExitCriteria                         []string `json:"pilot_exit_criteria,omitempty"`
	PilotOwnerRef                             string   `json:"pilot_owner_ref"`
	PilotTenantScope                          string   `json:"pilot_tenant_scope"`
	FirstRepoScope                            string   `json:"first_repo_scope"`
	EvidenceHandlingBoundary                  string   `json:"evidence_handling_boundary"`
	SuccessMetrics                            []string `json:"success_metrics,omitempty"`
	PilotSuccessDoesNotMeanProductionApproval bool     `json:"pilot_success_does_not_mean_production_approval"`
	ProductionApprovalImplied                 bool     `json:"production_approval_implied"`
	DeploymentApprovalImplied                 bool     `json:"deployment_approval_implied"`
	ComplianceGuaranteeImplied                bool     `json:"compliance_guarantee_implied"`
}

type Point13Val0CustomerOnboardingBoundary struct {
	OnboardingChecklistRef                   string `json:"onboarding_checklist_ref"`
	FirstRepoIntakeBoundary                  string `json:"first_repo_intake_boundary"`
	TenantScopeRequired                      string `json:"tenant_scope_required"`
	CustomerArtifactClassification           string `json:"customer_artifact_classification"`
	CustomerUploadIsCandidateOnly            bool   `json:"customer_upload_is_candidate_only"`
	CanonicalEvidenceRequiresGovernanceEvent bool   `json:"canonical_evidence_requires_governance_event"`
	SupportMaterialMutatesCanonicalEvidence  bool   `json:"support_material_mutates_canonical_evidence"`
	CustomerArtifactPromotedToCanonical      bool   `json:"customer_artifact_promoted_to_canonical"`
	CanonicalGovernanceEventRef              string `json:"canonical_governance_event_ref"`
	EvidenceIdentityRef                      string `json:"evidence_identity_ref"`
	PolicyVersionRef                         string `json:"policy_version_ref"`
	EngineVersionRef                         string `json:"engine_version_ref"`
	SchemaVersionRef                         string `json:"schema_version_ref"`
}

type Point13Val0SupportEscalationBoundary struct {
	EscalationOwnerRef                     string   `json:"escalation_owner_ref"`
	Severity                               string   `json:"severity"`
	SupportAccessScope                     string   `json:"support_access_scope"`
	TenantScope                            string   `json:"tenant_scope"`
	EvidenceRefs                           []string `json:"evidence_refs,omitempty"`
	AuditEventRef                          string   `json:"audit_event_ref"`
	SupportCannotBypassEvidenceSpine       bool     `json:"support_cannot_bypass_evidence_spine"`
	SupportCannotOverrideCoreDecision      bool     `json:"support_cannot_override_core_decision"`
	SupportCannotApproveProductionMutation bool     `json:"support_cannot_approve_production_mutation"`
	EvidenceSpineBypassAttempted           bool     `json:"evidence_spine_bypass_attempted"`
	CoreDecisionOverrideAttempted          bool     `json:"core_decision_override_attempted"`
	ProductionMutationApprovalAttempted    bool     `json:"production_mutation_approval_attempted"`
}

type Point13Val0OffboardingRetentionBoundary struct {
	RetentionOwnerRef                           string `json:"retention_owner_ref"`
	DisposalPathRef                             string `json:"disposal_path_ref"`
	TenantScope                                 string `json:"tenant_scope"`
	CustomerArtifactDisposalBoundary            string `json:"customer_artifact_disposal_boundary"`
	EvidenceRetentionClassRef                   string `json:"evidence_retention_class_ref"`
	SupportArtifactRetentionClassRef            string `json:"support_artifact_retention_class_ref"`
	NoIndefiniteRetentionWithoutGovernanceEvent bool   `json:"no_indefinite_retention_without_governance_event"`
	IndefiniteRetentionRequested                bool   `json:"indefinite_retention_requested"`
	RetentionGovernanceEventRef                 string `json:"retention_governance_event_ref"`
	PilotArtifactPromotedToCanonical            bool   `json:"pilot_artifact_promoted_to_canonical"`
	SupportArtifactPromotedToCanonical          bool   `json:"support_artifact_promoted_to_canonical"`
	CanonicalizationGovernanceEventRef          string `json:"canonicalization_governance_event_ref"`
}

type Point13Val0NoOverclaimCustomerWording struct {
	ObservedCustomerFacingTexts          []string `json:"observed_customer_facing_texts,omitempty"`
	ObservedExportFacingTexts            []string `json:"observed_export_facing_texts,omitempty"`
	ObservedSupportFacingTexts           []string `json:"observed_support_facing_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedCustomerFacingWording         []string `json:"allowed_customer_facing_wording,omitempty"`
	BlockedCustomerFacingWording         []string `json:"blocked_customer_facing_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point13Val0AIEvidenceCandidatePilotBoundary struct {
	AIOutputType                       string `json:"ai_output_type"`
	EvidenceCandidateOnly              bool   `json:"evidence_candidate_only"`
	AdvisoryOnly                       bool   `json:"advisory_only"`
	PassAllowed                        bool   `json:"pass_allowed"`
	ApprovalGranted                    bool   `json:"approval_granted"`
	ApprovalRequestCreatesApproval     bool   `json:"approval_request_creates_approval"`
	DeploymentAuthorized               bool   `json:"deployment_authorized"`
	ProductionReadinessClaimed         bool   `json:"production_readiness_claimed"`
	ProductionMutationAllowed          bool   `json:"production_mutation_allowed"`
	CanonicalMutationAllowed           bool   `json:"canonical_mutation_allowed"`
	HumanApprovalRequired              bool   `json:"human_approval_required"`
	ProductionImpactingActionRequested bool   `json:"production_impacting_action_requested"`
	HumanApprovalRef                   string `json:"human_approval_ref"`
	TenantScope                        string `json:"tenant_scope"`
	ReasonRef                          string `json:"reason_ref"`
	ExpiryWindowRef                    string `json:"expiry_window_ref"`
	SandboxResultRef                   string `json:"sandbox_result_ref"`
	RollbackPlanRef                    string `json:"rollback_plan_ref"`
	AuditEventRef                      string `json:"audit_event_ref"`
	PostActionVerificationPlanRef      string `json:"post_action_verification_plan_ref"`
}

type Point13Val0Foundation struct {
	CurrentState                     string                                      `json:"current_state"`
	BlockingReasons                  []string                                    `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites              []string                                    `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer             string                                      `json:"projection_disclaimer"`
	DependencyState                  string                                      `json:"dependency_state"`
	PilotReadinessState              string                                      `json:"pilot_readiness_state"`
	CustomerOnboardingState          string                                      `json:"customer_onboarding_state"`
	SupportEscalationState           string                                      `json:"support_escalation_state"`
	OffboardingRetentionState        string                                      `json:"offboarding_retention_state"`
	NoOverclaimState                 string                                      `json:"no_overclaim_state"`
	AIPilotBoundaryState             string                                      `json:"ai_pilot_boundary_state"`
	Dependency                       Point13Val0DependencySnapshot               `json:"dependency"`
	PilotReadiness                   Point13Val0PilotReadinessFoundation         `json:"pilot_readiness"`
	CustomerOnboardingBoundary       Point13Val0CustomerOnboardingBoundary       `json:"customer_onboarding_boundary"`
	SupportEscalationBoundary        Point13Val0SupportEscalationBoundary        `json:"support_escalation_boundary"`
	OffboardingRetentionBoundary     Point13Val0OffboardingRetentionBoundary     `json:"offboarding_retention_boundary"`
	NoOverclaimCustomerWording       Point13Val0NoOverclaimCustomerWording       `json:"no_overclaim_customer_wording"`
	AIEvidenceCandidatePilotBoundary Point13Val0AIEvidenceCandidatePilotBoundary `json:"ai_evidence_candidate_pilot_boundary"`
}

func point13Val0States() []string {
	return []string{
		Point13Val0StateActive,
		Point13Val0StateBlocked,
		Point13Val0StateReviewRequired,
		Point13Val0StateIncomplete,
	}
}

func point13Val0StateValid(value string) bool {
	return point13Val0RawExactOneOf(value, point13Val0States())
}

func point13Val0OperationalRefValid(value string, prefixes ...string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, prefixes)
}

func point13Val0OwnerRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "pilot_owner_", "support_owner_", "retention_owner_", "owner_")
}

func point13Val0ChecklistRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "checklist_", "onboarding_checklist_")
}

func point13Val0BoundaryRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "boundary_", "intake_boundary_", "evidence_boundary_", "disposal_boundary_")
}

func point13Val0PullRequestRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "pull_request_")
}

func point13Val0CommitRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "commit_", "merge_commit_")
}

func point13Val0WorkflowRunRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "workflow_run_")
}

func point13Val0ReasonRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "reason_", "change_reason_")
}

func point13Val0ExpiryWindowRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "expiry_window_")
}

func point13Val0VersionContextRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{
		"policy_version_",
		"engine_version_",
		"schema_version_",
	})
}

func point13Val0RollbackPlanRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "rollback_plan_")
}

func point13Val0VerificationPlanRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "verification_plan_", "post_action_verification_plan_")
}

func point13Val0ApprovalRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "approval_event_", "human_approval_")
}

func point13Val0CustomerArtifactClassifications() []string {
	return []string{point13Val0CustomerArtifactCandidateOnly}
}

func point13Val0SupportSeverities() []string {
	return []string{
		point13Val0SupportSeverityCritical,
		point13Val0SupportSeverityHigh,
		point13Val0SupportSeverityModerate,
		point13Val0SupportSeverityLow,
	}
}

func point13Val0RawExactOneOf(value string, values []string) bool {
	if !formalRawExactNonEmpty(value) {
		return false
	}
	for _, candidate := range values {
		if value == candidate {
			return true
		}
	}
	return false
}

func point13Val0RawOperationalRefValid(value string, prefixes ...string) bool {
	return formalRawExactTokenValid(value, func(candidate string) bool {
		return point13Val0OperationalRefValid(candidate, prefixes...)
	})
}

func point13Val0RawOwnerRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0OwnerRefValid)
}

func point13Val0RawChecklistRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0ChecklistRefValid)
}

func point13Val0RawBoundaryRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0BoundaryRefValid)
}

func point13Val0RawPullRequestRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0PullRequestRefValid)
}

func point13Val0RawCommitRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0CommitRefValid)
}

func point13Val0RawWorkflowRunRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0WorkflowRunRefValid)
}

func point13Val0RawReasonRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0ReasonRefValid)
}

func point13Val0RawExpiryWindowRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0ExpiryWindowRefValid)
}

func point13Val0RawVersionContextRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0VersionContextRefValid)
}

func point13Val0RawRollbackPlanRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0RollbackPlanRefValid)
}

func point13Val0RawVerificationPlanRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0VerificationPlanRefValid)
}

func point13Val0RawApprovalRefValid(value string) bool {
	return formalRawExactTokenValid(value, point13Val0ApprovalRefValid)
}

func point13Val0RawScopeValid(value string) bool {
	return formalRawExactTokenValid(value, point11Val0ScopeValid)
}

func point13Val0RawIdentityValueValid(value string) bool {
	return formalRawExactTokenValid(value, point11Val0IdentityValueValid)
}

func point13Val0AllowedCustomerWording() []string {
	return []string{
		"pilot readiness support",
		"evidence candidate",
		"advisory recommendation",
		"evidence support for customer/auditor review",
		"operational onboarding boundary",
	}
}

func point13Val0ForbiddenClaims() []string {
	return append([]string{
		"certified",
		"guaranteed secure",
		"regulator-approved",
		"compliance guaranteed",
		"production approved",
		"deployment approved",
		"public badge",
		"global truth",
		"official authority",
		"ai certified",
		"ai-approved",
		"ai approved",
		"ai decision",
		"autonomous remediation",
		"continuous compliance attestation",
		"ai legal proof",
		"legal proof",
		"financial guarantee",
		"premium reduction guarantee",
	}, inheritedDeploymentReadinessOverclaimClaims()...)
}

func point13Val0ContainsForbiddenClaim(values ...string) bool {
	return point13Val0ContainsForbiddenClaimWithAllowed(point13Val0AllowedCustomerWording(), values...)
}

func point13Val0ContainsForbiddenClaimWithAllowed(allowedValues []string, values ...string) bool {
	return point13Val0ContainsForbiddenClaimWithAllowedAndNormalizer(allowedValues, formalNoOverclaimNormalizePublicText, values...)
}

func point13Val0ContainsForbiddenInternalClaim(values ...string) bool {
	return point13Val0ContainsForbiddenClaimWithAllowedAndNormalizer(point13Val0AllowedCustomerWording(), formalNoOverclaimNormalizeText, values...)
}

func point13Val0ContainsForbiddenClaimWithAllowedAndNormalizer(allowedValues []string, normalize func(string) string, values ...string) bool {
	allowed := map[string]struct{}{}
	for _, value := range allowedValues {
		allowed[normalize(value)] = struct{}{}
	}
	crossNormalizedParts := make([]string, 0, len(values))
	crossPartAllowed := make([]bool, 0, len(values))
	for _, value := range values {
		normalized := normalize(value)
		if normalized == "" {
			continue
		}
		_, isAllowed := allowed[normalized]
		crossNormalizedParts = append(crossNormalizedParts, normalized)
		crossPartAllowed = append(crossPartAllowed, isAllowed)
		if isAllowed {
			continue
		}
		for _, forbidden := range point13Val0ForbiddenClaims() {
			if formalNoOverclaimContainsForbidden(normalized, normalize(forbidden)) {
				return true
			}
		}
	}
	for _, forbidden := range point13Val0ForbiddenClaims() {
		if point13Val0ForbiddenPhraseAcrossValues(crossNormalizedParts, crossPartAllowed, normalize(forbidden)) {
			return true
		}
	}
	return false
}

func point13Val0ForbiddenPhraseAcrossValues(values []string, allowed []bool, phrase string) bool {
	if len(values) != len(allowed) {
		return false
	}
	if formalNoOverclaimForbiddenCompactAcrossValues(values, allowed, phrase) {
		return true
	}
	phraseTokens := strings.Fields(phrase)
	if len(phraseTokens) < 2 {
		return false
	}
	matched := 0
	distinctBuckets := 0
	lastBucket := -1
	matchedIncludesNonAllowed := false
	for bucketIndex, value := range values {
		bucketTokens := strings.Fields(value)
		if len(bucketTokens) == 0 {
			continue
		}
		for _, token := range bucketTokens {
			if token != phraseTokens[matched] {
				continue
			}
			if bucketIndex != lastBucket {
				distinctBuckets++
				lastBucket = bucketIndex
			}
			if !allowed[bucketIndex] {
				matchedIncludesNonAllowed = true
			}
			matched++
			if matched == len(phraseTokens) {
				return distinctBuckets > 1 && matchedIncludesNonAllowed
			}
		}
	}
	return false
}

func point13Val0StringListValid(values []string) bool {
	return point12Val0StringListValid(values, point11Val0IdentityValueValid)
}

func point13Val0DependencySnapshotFromUpstream(point12 Point12ValEFoundation) Point13Val0DependencySnapshot {
	return Point13Val0DependencySnapshot{
		Point12CurrentState:             point12.CurrentState,
		Point12DependencyState:          point12.DependencyState,
		Point12PassClosureManifestState: point12.PassClosureManifestState,
		Point12ReviewerResult:           point12.PassClosureManifest.ReviewerResult,
		Point12PassAllowed:              point12.Point12PassAllowed,
		Point12PassToken:                point12.Point12PassToken,
		Point12ClosureManifestRef:       point12.PassClosureManifest.ClosureManifestID,
		Point12DependencySnapshotRef:    point12.Dependency.SnapshotRef,
		Point12ProofPackID:              point12.PassClosureManifest.ProofPackID,
		Point12TenantScope:              point12.PassClosureManifest.TenantScope,
		AIGovernanceBackfillVerified:    true,
		AIGovernanceBackfillMerged:      true,
		GitHubCIVerified:                true,
		BackfillPullRequestRef:          "pull_request_ai_governance_backfill_126",
		BackfillHeadCommitRef:           "commit_15abca9f19645418acd1413e4e05ce320810a151",
		BackfillMergeCommitRef:          "merge_commit_a80d735f2c121216ff7b7ed8d12bb55279c1e94a",
		BackfillTestWorkflowRunRef:      "workflow_run_test_25334743024",
		BackfillLintWorkflowRunRef:      "workflow_run_lint_25334743188",
		SnapshotFromComputedOutput:      true,
		Point12:                         point12,
	}
}

func point13Val0DependencySnapshotModel() Point13Val0DependencySnapshot {
	return point13Val0DependencySnapshotFromUpstream(ComputePoint12ValEFoundation(Point12ValEFoundationModel()))
}

func point13Val0DependencyStateAndReasons(model Point13Val0DependencySnapshot) (string, []string) {
	reviewReasons := []string{}
	blockedReasons := []string{}
	recomputedPoint12 := ComputePoint12ValEFoundation(model.Point12)

	if !model.SnapshotFromComputedOutput ||
		!point13Val0RawPullRequestRefValid(model.BackfillPullRequestRef) ||
		!point13Val0RawCommitRefValid(model.BackfillHeadCommitRef) ||
		!point13Val0RawCommitRefValid(model.BackfillMergeCommitRef) ||
		!point13Val0RawWorkflowRunRefValid(model.BackfillTestWorkflowRunRef) ||
		!point13Val0RawWorkflowRunRefValid(model.BackfillLintWorkflowRunRef) {
		reviewReasons = append(reviewReasons, "dependency_verification_refs_missing")
	}
	if !model.AIGovernanceBackfillVerified || !model.AIGovernanceBackfillMerged || !model.GitHubCIVerified {
		reviewReasons = append(reviewReasons, "ai_governance_backfill_unverified")
	}
	if !formalRawExactValid(model.Point12CurrentState, point12ValEStateValid) ||
		!formalRawExactValid(model.Point12DependencyState, point12ValEStateValid) ||
		!formalRawExactValid(model.Point12PassClosureManifestState, point12ValEStateValid) ||
		!formalRawExactValid(model.Point12ReviewerResult, point12ValEReviewerResultValid) ||
		!formalRawExactTokenValid(model.Point12ClosureManifestRef, point12ValEClosureManifestRefValid) ||
		!formalRawExactTokenValid(model.Point12DependencySnapshotRef, point12ValEDependencySnapshotRefValid) ||
		!formalRawExactTokenValid(model.Point12ProofPackID, point12Val0ProofPackRefValid) ||
		!point13Val0RawScopeValid(model.Point12TenantScope) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_identity_invalid")
	}
	if model.Point12.CurrentState != recomputedPoint12.CurrentState ||
		model.Point12.DependencyState != recomputedPoint12.DependencyState ||
		model.Point12.PassClosureManifestState != recomputedPoint12.PassClosureManifestState ||
		model.Point12.PassClosureManifest.ReviewerResult != recomputedPoint12.PassClosureManifest.ReviewerResult ||
		model.Point12.Point12PassAllowed != recomputedPoint12.Point12PassAllowed ||
		model.Point12.Point12PassToken != recomputedPoint12.Point12PassToken ||
		model.Point12.PassClosureManifest.ClosureManifestID != recomputedPoint12.PassClosureManifest.ClosureManifestID ||
		model.Point12.Dependency.SnapshotRef != recomputedPoint12.Dependency.SnapshotRef ||
		model.Point12.PassClosureManifest.ProofPackID != recomputedPoint12.PassClosureManifest.ProofPackID ||
		model.Point12.PassClosureManifest.TenantScope != recomputedPoint12.PassClosureManifest.TenantScope {
		blockedReasons = append(blockedReasons, "point12_recomputed_snapshot_mismatch")
	}
	if model.Point12CurrentState != recomputedPoint12.CurrentState ||
		model.Point12DependencyState != recomputedPoint12.DependencyState ||
		model.Point12PassClosureManifestState != recomputedPoint12.PassClosureManifestState ||
		model.Point12ReviewerResult != recomputedPoint12.PassClosureManifest.ReviewerResult ||
		model.Point12PassAllowed != recomputedPoint12.Point12PassAllowed ||
		model.Point12PassToken != recomputedPoint12.Point12PassToken ||
		model.Point12ClosureManifestRef != recomputedPoint12.PassClosureManifest.ClosureManifestID ||
		model.Point12DependencySnapshotRef != recomputedPoint12.Dependency.SnapshotRef ||
		model.Point12ProofPackID != recomputedPoint12.PassClosureManifest.ProofPackID ||
		model.Point12TenantScope != recomputedPoint12.PassClosureManifest.TenantScope {
		blockedReasons = append(blockedReasons, "dependency_snapshot_binding_mismatch")
	}
	if model.Point12CurrentState == Point12ValEStateReviewRequired ||
		model.Point12DependencyState == Point12ValEStateReviewRequired ||
		model.Point12PassClosureManifestState == Point12ValEStateReviewRequired ||
		model.Point12ReviewerResult == point12ValEReviewerResultReviewRequired {
		reviewReasons = append(reviewReasons, "point12_review_required")
	}
	if model.Point12CurrentState != Point12ValEStatePassConfirmed {
		blockedReasons = append(blockedReasons, "point12_not_pass_confirmed")
	}
	if model.Point12DependencyState != Point12ValEStateActive {
		blockedReasons = append(blockedReasons, "point12_dependency_not_active")
	}
	if model.Point12PassClosureManifestState != Point12ValEStateActive {
		blockedReasons = append(blockedReasons, "point12_pass_closure_manifest_not_active")
	}
	if model.Point12ReviewerResult != point12ValEReviewerResultPassConfirmed {
		blockedReasons = append(blockedReasons, "point12_reviewer_not_pass_confirmed")
	}
	if !model.Point12PassAllowed || model.Point12PassToken != point12ValEPoint12PassToken {
		blockedReasons = append(blockedReasons, "point12_pass_evidence_missing")
	}
	if len(blockedReasons) > 0 {
		return Point13Val0StateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point13Val0StateReviewRequired, reviewReasons
	}
	return Point13Val0StateActive, nil
}

func EvaluatePoint13Val0PilotReadinessState(model Point13Val0PilotReadinessFoundation, dependency Point13Val0DependencySnapshot) string {
	if !point13Val0StringListValid(model.PilotEntryCriteria) ||
		!point13Val0StringListValid(model.PilotExitCriteria) ||
		!point13Val0RawOwnerRefValid(model.PilotOwnerRef) ||
		!point13Val0RawScopeValid(model.PilotTenantScope) ||
		!point13Val0RawIdentityValueValid(model.FirstRepoScope) ||
		!point13Val0RawBoundaryRefValid(model.EvidenceHandlingBoundary) ||
		!point13Val0StringListValid(model.SuccessMetrics) ||
		!model.PilotSuccessDoesNotMeanProductionApproval ||
		model.ProductionApprovalImplied ||
		model.DeploymentApprovalImplied ||
		model.ComplianceGuaranteeImplied {
		return Point13Val0StateBlocked
	}
	if model.PilotTenantScope != dependency.Point12TenantScope {
		return Point13Val0StateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.SuccessMetrics, " "),
		strings.Join(model.PilotEntryCriteria, " "),
		strings.Join(model.PilotExitCriteria, " "),
	) {
		return Point13Val0StateBlocked
	}
	return Point13Val0StateActive
}

func EvaluatePoint13Val0CustomerOnboardingBoundaryState(model Point13Val0CustomerOnboardingBoundary, dependency Point13Val0DependencySnapshot, pilot Point13Val0PilotReadinessFoundation) string {
	if !point13Val0RawChecklistRefValid(model.OnboardingChecklistRef) ||
		!point13Val0RawBoundaryRefValid(model.FirstRepoIntakeBoundary) ||
		!point13Val0RawScopeValid(model.TenantScopeRequired) ||
		!point13Val0RawExactOneOf(model.CustomerArtifactClassification, point13Val0CustomerArtifactClassifications()) ||
		!model.CustomerUploadIsCandidateOnly ||
		!model.CanonicalEvidenceRequiresGovernanceEvent ||
		model.SupportMaterialMutatesCanonicalEvidence ||
		!formalRawExactTokenValid(model.EvidenceIdentityRef, point12Val0ArtifactRefValid) ||
		!point13Val0RawVersionContextRefValid(model.PolicyVersionRef) ||
		!point13Val0RawVersionContextRefValid(model.EngineVersionRef) ||
		!point13Val0RawVersionContextRefValid(model.SchemaVersionRef) {
		return Point13Val0StateBlocked
	}
	if model.TenantScopeRequired != dependency.Point12TenantScope {
		return Point13Val0StateBlocked
	}
	if model.FirstRepoIntakeBoundary != pilot.EvidenceHandlingBoundary ||
		model.EvidenceIdentityRef != dependency.Point12.PassClosureManifest.ArtifactRef {
		return Point13Val0StateBlocked
	}
	if model.CustomerArtifactPromotedToCanonical && !point12Val0GovernanceEventRefValid(model.CanonicalGovernanceEventRef) {
		return Point13Val0StateBlocked
	}
	return Point13Val0StateActive
}

func EvaluatePoint13Val0SupportEscalationBoundaryState(model Point13Val0SupportEscalationBoundary, dependency Point13Val0DependencySnapshot) string {
	if !point13Val0RawOwnerRefValid(model.EscalationOwnerRef) ||
		!point13Val0RawExactOneOf(model.Severity, point13Val0SupportSeverities()) ||
		!point13Val0RawIdentityValueValid(model.SupportAccessScope) ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point12Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		!model.SupportCannotBypassEvidenceSpine ||
		!model.SupportCannotOverrideCoreDecision ||
		!model.SupportCannotApproveProductionMutation ||
		model.EvidenceSpineBypassAttempted ||
		model.CoreDecisionOverrideAttempted ||
		model.ProductionMutationApprovalAttempted {
		return Point13Val0StateBlocked
	}
	if model.TenantScope != dependency.Point12TenantScope ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, dependency.Point12.Dependency.ValD.ProofChain.EvidenceRefs) {
		return Point13Val0StateBlocked
	}
	return Point13Val0StateActive
}

func EvaluatePoint13Val0OffboardingRetentionBoundaryState(model Point13Val0OffboardingRetentionBoundary, dependency Point13Val0DependencySnapshot) string {
	if !point13Val0RawOwnerRefValid(model.RetentionOwnerRef) ||
		!point13Val0RawOperationalRefValid(model.DisposalPathRef, "disposal_path_") ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point13Val0RawBoundaryRefValid(model.CustomerArtifactDisposalBoundary) ||
		!formalRawExactTokenValid(model.EvidenceRetentionClassRef, point12Val0RetentionClassRefValid) ||
		!formalRawExactTokenValid(model.SupportArtifactRetentionClassRef, point12Val0RetentionClassRefValid) ||
		!model.NoIndefiniteRetentionWithoutGovernanceEvent {
		return Point13Val0StateBlocked
	}
	if model.TenantScope != dependency.Point12TenantScope {
		return Point13Val0StateBlocked
	}
	if model.IndefiniteRetentionRequested && !point12Val0GovernanceEventRefValid(model.RetentionGovernanceEventRef) {
		return Point13Val0StateBlocked
	}
	if (model.PilotArtifactPromotedToCanonical || model.SupportArtifactPromotedToCanonical) &&
		!point12Val0GovernanceEventRefValid(model.CanonicalizationGovernanceEventRef) {
		return Point13Val0StateBlocked
	}
	return Point13Val0StateActive
}

func EvaluatePoint13Val0NoOverclaimCustomerWordingState(model Point13Val0NoOverclaimCustomerWording) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point13Val0StringListValid(model.AllowedCustomerFacingWording) ||
		!point13Val0StringListValid(model.BlockedCustomerFacingWording) {
		return Point13Val0StateBlocked
	}
	if !point12Val0ExactStringSetMatch(model.AllowedCustomerFacingWording, point13Val0AllowedCustomerWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedCustomerFacingWording, point13Val0ForbiddenClaims()) {
		return Point13Val0StateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.ObservedCustomerFacingTexts, " "),
		strings.Join(model.ObservedExportFacingTexts, " "),
		strings.Join(model.ObservedSupportFacingTexts, " "),
	) {
		return Point13Val0StateBlocked
	}
	if point13Val0ContainsForbiddenInternalClaim(strings.Join(model.InternalDiagnosticTexts, " ")) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point13Val0StateBlocked
	}
	return Point13Val0StateActive
}

func EvaluatePoint13Val0AIPilotBoundaryState(model Point13Val0AIEvidenceCandidatePilotBoundary, dependency Point13Val0DependencySnapshot) string {
	if point13Val0RawExactOneOf(model.AIOutputType, point12Val0BlockedAIEvidenceCandidateTypes()) {
		return Point13Val0StateBlocked
	}
	if !formalRawExactValid(model.AIOutputType, point12Val0AIEvidenceCandidateTypeValid) ||
		!model.EvidenceCandidateOnly ||
		!model.AdvisoryOnly ||
		model.PassAllowed ||
		model.ApprovalGranted ||
		model.ApprovalRequestCreatesApproval ||
		model.ProductionMutationAllowed ||
		model.CanonicalMutationAllowed ||
		!model.HumanApprovalRequired ||
		!point13Val0RawScopeValid(model.TenantScope) {
		return Point13Val0StateBlocked
	}
	if model.TenantScope != dependency.Point12TenantScope {
		return Point13Val0StateBlocked
	}
	if model.DeploymentAuthorized || model.ProductionReadinessClaimed {
		return Point13Val0StateBlocked
	}
	if model.ProductionImpactingActionRequested {
		if !point13Val0RawApprovalRefValid(model.HumanApprovalRef) ||
			!point13Val0RawReasonRefValid(model.ReasonRef) ||
			!point13Val0RawExpiryWindowRefValid(model.ExpiryWindowRef) ||
			!point13Val0RawOperationalRefValid(model.SandboxResultRef, "sandbox_result_") ||
			!point13Val0RawRollbackPlanRefValid(model.RollbackPlanRef) ||
			!formalRawExactTokenValid(model.AuditEventRef, point12Val0AuditRefValid) ||
			!point13Val0RawVerificationPlanRefValid(model.PostActionVerificationPlanRef) {
			return Point13Val0StateBlocked
		}
	}
	return Point13Val0StateActive
}

func point13Val0BlockingReasons(model Point13Val0Foundation) []string {
	reasons := []string{}
	componentStates := map[string]string{
		"dependency":            model.DependencyState,
		"pilot_readiness":       model.PilotReadinessState,
		"customer_onboarding":   model.CustomerOnboardingState,
		"support_escalation":    model.SupportEscalationState,
		"offboarding_retention": model.OffboardingRetentionState,
		"no_overclaim":          model.NoOverclaimState,
		"ai_pilot_boundary":     model.AIPilotBoundaryState,
	}
	for name, state := range componentStates {
		if !formalRawExactValid(state, point13Val0StateValid) {
			reasons = append(reasons, name+":invalid_state")
			continue
		}
		if state == Point13Val0StateBlocked || state == Point13Val0StateIncomplete {
			reasons = append(reasons, name+":"+state)
		}
	}
	return reasons
}

func EvaluatePoint13Val0State(model Point13Val0Foundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13Val0StateBlocked
	}
	hasReview := false
	hasIncomplete := false
	for _, state := range []string{
		model.DependencyState,
		model.PilotReadinessState,
		model.CustomerOnboardingState,
		model.SupportEscalationState,
		model.OffboardingRetentionState,
		model.NoOverclaimState,
		model.AIPilotBoundaryState,
	} {
		if !formalRawExactValid(state, point13Val0StateValid) {
			return Point13Val0StateBlocked
		}
		switch state {
		case Point13Val0StateBlocked:
			return Point13Val0StateBlocked
		case Point13Val0StateReviewRequired:
			hasReview = true
		case Point13Val0StateIncomplete:
			hasIncomplete = true
		}
	}
	if hasReview {
		return Point13Val0StateReviewRequired
	}
	if hasIncomplete {
		return Point13Val0StateIncomplete
	}
	return Point13Val0StateActive
}

func Point13Val0FoundationModel() Point13Val0Foundation {
	disclaimer := point13Val0ProjectionDisclaimerBaseline
	dependency := point13Val0DependencySnapshotModel()
	tenantScope := dependency.Point12TenantScope
	artifactRef := dependency.Point12.PassClosureManifest.ArtifactRef
	return Point13Val0Foundation{
		CurrentState:              Point13Val0StateActive,
		ProjectionDisclaimer:      disclaimer,
		DependencyState:           Point13Val0StateActive,
		PilotReadinessState:       Point13Val0StateActive,
		CustomerOnboardingState:   Point13Val0StateActive,
		SupportEscalationState:    Point13Val0StateActive,
		OffboardingRetentionState: Point13Val0StateActive,
		NoOverclaimState:          Point13Val0StateActive,
		AIPilotBoundaryState:      Point13Val0StateActive,
		Dependency:                dependency,
		PilotReadiness: Point13Val0PilotReadinessFoundation{
			PilotEntryCriteria:       []string{"point12_vale_pass_confirmed", "tenant_scope_declared", "onboarding_boundary_declared"},
			PilotExitCriteria:        []string{"operational_readiness_review_complete", "support_escalation_boundary_reviewed"},
			PilotOwnerRef:            "pilot_owner_point13_val0_001",
			PilotTenantScope:         tenantScope,
			FirstRepoScope:           "repo_scope_point13_first_repo_001",
			EvidenceHandlingBoundary: "boundary_point13_first_repo_intake_001",
			SuccessMetrics:           []string{"pilot_readiness_support", "operational_onboarding_boundary"},
			PilotSuccessDoesNotMeanProductionApproval: true,
		},
		CustomerOnboardingBoundary: Point13Val0CustomerOnboardingBoundary{
			OnboardingChecklistRef:                   "onboarding_checklist_point13_val0_001",
			FirstRepoIntakeBoundary:                  "boundary_point13_first_repo_intake_001",
			TenantScopeRequired:                      tenantScope,
			CustomerArtifactClassification:           point13Val0CustomerArtifactCandidateOnly,
			CustomerUploadIsCandidateOnly:            true,
			CanonicalEvidenceRequiresGovernanceEvent: true,
			EvidenceIdentityRef:                      artifactRef,
			PolicyVersionRef:                         dependency.Point12.Dependency.Val0.Manifest.PolicyVersion,
			EngineVersionRef:                         dependency.Point12.Dependency.Val0.Manifest.EngineVersion,
			SchemaVersionRef:                         dependency.Point12.Dependency.Val0.Manifest.SchemaVersion,
		},
		SupportEscalationBoundary: Point13Val0SupportEscalationBoundary{
			EscalationOwnerRef:                     "support_owner_point13_val0_001",
			Severity:                               point13Val0SupportSeverityHigh,
			SupportAccessScope:                     "tenant_scoped_read_only",
			TenantScope:                            tenantScope,
			EvidenceRefs:                           append([]string{}, dependency.Point12.Dependency.ValD.ProofChain.EvidenceRefs...),
			AuditEventRef:                          "audit_point13_val0_support_001",
			SupportCannotBypassEvidenceSpine:       true,
			SupportCannotOverrideCoreDecision:      true,
			SupportCannotApproveProductionMutation: true,
		},
		OffboardingRetentionBoundary: Point13Val0OffboardingRetentionBoundary{
			RetentionOwnerRef:                           "retention_owner_point13_val0_001",
			DisposalPathRef:                             "disposal_path_point13_val0_001",
			TenantScope:                                 tenantScope,
			CustomerArtifactDisposalBoundary:            "disposal_boundary_point13_customer_artifacts_001",
			EvidenceRetentionClassRef:                   "retention_class_point13_evidence_candidate",
			SupportArtifactRetentionClassRef:            "retention_class_point13_support_artifact",
			NoIndefiniteRetentionWithoutGovernanceEvent: true,
		},
		NoOverclaimCustomerWording: Point13Val0NoOverclaimCustomerWording{
			ObservedCustomerFacingTexts:          []string{"pilot readiness support", "evidence candidate"},
			ObservedExportFacingTexts:            []string{"evidence support for customer/auditor review"},
			ObservedSupportFacingTexts:           []string{"operational onboarding boundary"},
			InternalDiagnosticTexts:              []string{"blocked wording remains denylisted internally"},
			InternalDiagnosticsClassifiedBlocked: true,
			AllowedCustomerFacingWording:         point13Val0AllowedCustomerWording(),
			BlockedCustomerFacingWording:         point13Val0ForbiddenClaims(),
			ProjectionDisclaimer:                 disclaimer,
		},
		AIEvidenceCandidatePilotBoundary: Point13Val0AIEvidenceCandidatePilotBoundary{
			AIOutputType:          "AI_RECOMMENDATION",
			EvidenceCandidateOnly: true,
			AdvisoryOnly:          true,
			HumanApprovalRequired: true,
			TenantScope:           tenantScope,
		},
	}
}

func ComputePoint13Val0Foundation(model Point13Val0Foundation) Point13Val0Foundation {
	dependencyState, dependencyReasons := point13Val0DependencyStateAndReasons(model.Dependency)
	model.DependencyState = dependencyState
	model.PilotReadinessState = EvaluatePoint13Val0PilotReadinessState(model.PilotReadiness, model.Dependency)
	model.CustomerOnboardingState = EvaluatePoint13Val0CustomerOnboardingBoundaryState(model.CustomerOnboardingBoundary, model.Dependency, model.PilotReadiness)
	model.SupportEscalationState = EvaluatePoint13Val0SupportEscalationBoundaryState(model.SupportEscalationBoundary, model.Dependency)
	model.OffboardingRetentionState = EvaluatePoint13Val0OffboardingRetentionBoundaryState(model.OffboardingRetentionBoundary, model.Dependency)
	model.NoOverclaimState = EvaluatePoint13Val0NoOverclaimCustomerWordingState(model.NoOverclaimCustomerWording)
	model.AIPilotBoundaryState = EvaluatePoint13Val0AIPilotBoundaryState(model.AIEvidenceCandidatePilotBoundary, model.Dependency)
	model.CurrentState = EvaluatePoint13Val0State(model)
	model.BlockingReasons = point13Val0BlockingReasons(model)
	model.ReviewPrerequisites = nil
	if model.DependencyState == Point13Val0StateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, dependencyReasons...)
	}
	return model
}
