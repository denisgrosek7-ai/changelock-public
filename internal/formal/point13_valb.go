package formal

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

const (
	Point13ValBStateActive         = "point13_valb_pilot_evidence_operations_active"
	Point13ValBStateBlocked        = "point13_valb_pilot_evidence_operations_blocked"
	Point13ValBStateReviewRequired = "point13_valb_pilot_evidence_operations_review_required"
	Point13ValBStateIncomplete     = "point13_valb_pilot_evidence_operations_incomplete"
)

const (
	point13ValBWaveID                       = "val_b"
	point13ValBPreviousWaveID               = point13ValAWaveID
	point13ValBProjectionDisclaimerBaseline = "projection_only not_canonical_truth point13_valb_pilot_evidence_operations"

	point13ValBLedgerOperationCustomerArtifactReceived  = "customer_artifact_received"
	point13ValBLedgerOperationEvidenceCandidateRegister = "evidence_candidate_registered"
	point13ValBLedgerOperationCustodyVerified           = "custody_verified"
	point13ValBLedgerOperationSandboxResultRecorded     = "sandbox_result_recorded"
	point13ValBLedgerOperationSupportActionRecorded     = "support_action_recorded"
	point13ValBLedgerOperationCustomerReviewRecorded    = "customer_review_recorded"
	point13ValBLedgerOperationExitEvidencePacket        = "exit_evidence_packet_prepared"

	point13ValBSupportActionTriage                 = "triage"
	point13ValBSupportActionEvidenceRequest        = "evidence_request"
	point13ValBSupportActionEvidenceClassification = "evidence_classification"
	point13ValBSupportActionSandboxFollowup        = "sandbox_followup"
	point13ValBSupportActionCustomerQuestion       = "customer_question_answered"
	point13ValBSupportActionOffboardingAssist      = "offboarding_assist"
)

type Point13ValBDependencySnapshot struct {
	ValACurrentState                          string                `json:"vala_current_state"`
	ValADependencyState                       string                `json:"vala_dependency_state"`
	ValAPilotExecutionContractState           string                `json:"vala_pilot_execution_contract_state"`
	ValACustomerIntakeEvidenceGovernanceState string                `json:"vala_customer_intake_evidence_governance_state"`
	ValAPilotRunPhaseBoundaryState            string                `json:"vala_pilot_run_phase_boundary_state"`
	ValASupportResponsibilityMatrixState      string                `json:"vala_support_responsibility_matrix_state"`
	ValAPilotExitReviewGateState              string                `json:"vala_pilot_exit_review_gate_state"`
	ValAAIAssistedPilotExecutionBoundaryState string                `json:"vala_ai_assisted_pilot_execution_boundary_state"`
	ValANoOverclaimState                      string                `json:"vala_no_overclaim_state"`
	ValAPointID                               string                `json:"vala_point_id"`
	ValAWaveID                                string                `json:"vala_wave_id"`
	ValADependencyComputedFromUpstream        bool                  `json:"vala_dependency_computed_from_upstream"`
	ValAPoint13PassSeen                       bool                  `json:"vala_point13_pass_seen"`
	InheritedVal0CurrentState                 string                `json:"inherited_val0_current_state"`
	InheritedPoint12CurrentState              string                `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState           string                `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureManifestState  string                `json:"inherited_point12_pass_closure_manifest_state"`
	InheritedPoint12ReviewerResult            string                `json:"inherited_point12_reviewer_result"`
	InheritedTenantScope                      string                `json:"inherited_tenant_scope"`
	InheritedAIModelOrRuleVersionRef          string                `json:"inherited_ai_model_or_rule_version_ref"`
	InheritedAIPermissionManifestHash         string                `json:"inherited_ai_permission_manifest_hash"`
	SnapshotFromComputedOutput                bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                       []string              `json:"review_prerequisites,omitempty"`
	ValA                                      Point13ValAFoundation `json:"vala"`
}

type Point13ValBOperationLedgerEntry struct {
	EntryID                   string   `json:"entry_id"`
	OperationType             string   `json:"operation_type"`
	OwnerRef                  string   `json:"owner_ref"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs          []string `json:"evidence_hash_refs,omitempty"`
	AuditEventRef             string   `json:"audit_event_ref"`
	CustodyRef                string   `json:"custody_ref"`
	SourceRef                 string   `json:"source_ref"`
	CandidateOnly             bool     `json:"candidate_only"`
	CanonicalMutationAllowed  bool     `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed bool     `json:"production_mutation_allowed"`
	PassAllowed               bool     `json:"pass_allowed"`
}

type Point13ValBPilotEvidenceOperationLedger struct {
	LedgerID          string                            `json:"ledger_id"`
	TenantScope       string                            `json:"tenant_scope"`
	PilotScopeRef     string                            `json:"pilot_scope_ref"`
	OperationEntries  []Point13ValBOperationLedgerEntry `json:"operation_entries,omitempty"`
	LedgerBindingHash string                            `json:"ledger_binding_hash"`
}

type Point13ValBCustomerReviewTrace struct {
	ReviewTraceID                         string   `json:"review_trace_id"`
	TenantScope                           string   `json:"tenant_scope"`
	CustomerReviewRefs                    []string `json:"customer_review_refs,omitempty"`
	ReviewedEvidenceRefs                  []string `json:"reviewed_evidence_refs,omitempty"`
	ReviewedEvidenceHashRefs              []string `json:"reviewed_evidence_hash_refs,omitempty"`
	CustomerOwnerRef                      string   `json:"customer_owner_ref"`
	InternalOwnerRef                      string   `json:"internal_owner_ref"`
	SupportOwnerRef                       string   `json:"support_owner_ref"`
	AuditEventRefs                        []string `json:"audit_event_refs,omitempty"`
	FinalCustomerStatement                string   `json:"final_customer_statement"`
	CustomerReviewIsOperationalOnly       bool     `json:"customer_review_is_operational_only"`
	CustomerReviewCannotApproveProduction bool     `json:"customer_review_cannot_approve_production"`
	CustomerReviewCannotCreatePass        bool     `json:"customer_review_cannot_create_pass"`
}

type Point13ValBSupportActionTrace struct {
	SupportTraceID                    string   `json:"support_trace_id"`
	TenantScope                       string   `json:"tenant_scope"`
	SupportActionRefs                 []string `json:"support_action_refs,omitempty"`
	SupportActionTypes                []string `json:"support_action_types,omitempty"`
	SupportOwnerRef                   string   `json:"support_owner_ref"`
	CustomerOwnerRef                  string   `json:"customer_owner_ref"`
	EvidenceRefs                      []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs                  []string `json:"evidence_hash_refs,omitempty"`
	AuditEventRefs                    []string `json:"audit_event_refs,omitempty"`
	SupportCanViewCandidateArtifacts  bool     `json:"support_can_view_candidate_artifacts"`
	SupportCanMutateCanonicalEvidence bool     `json:"support_can_mutate_canonical_evidence"`
	SupportCanOverrideCoreDecision    bool     `json:"support_can_override_core_decision"`
	SupportCanApproveProduction       bool     `json:"support_can_approve_production"`
}

type Point13ValBPilotExitEvidencePacket struct {
	PacketID                 string   `json:"packet_id"`
	TenantScope              string   `json:"tenant_scope"`
	OperationalReadinessOnly bool     `json:"operational_readiness_only"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs         []string `json:"evidence_hash_refs,omitempty"`
	CustomerReviewTraceRef   string   `json:"customer_review_trace_ref"`
	SupportTraceRef          string   `json:"support_trace_ref"`
	OperationLedgerRef       string   `json:"operation_ledger_ref"`
	UnresolvedBlockers       []string `json:"unresolved_blockers,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
	SafeCustomerStatement    string   `json:"safe_customer_statement"`
	NoProductionApproval     bool     `json:"no_production_approval"`
	NoDeploymentApproval     bool     `json:"no_deployment_approval"`
	NoComplianceGuarantee    bool     `json:"no_compliance_guarantee"`
	NoCertification          bool     `json:"no_certification"`
	NoFinancialGuarantee     bool     `json:"no_financial_guarantee"`
	NoPoint13Pass            bool     `json:"no_point13_pass"`
}

type Point13ValBAIEvidenceOperationTrace struct {
	AITraceID                              string   `json:"ai_trace_id"`
	TenantScope                            string   `json:"tenant_scope"`
	AIOutputType                           string   `json:"ai_output_type"`
	EvidenceCandidateRef                   string   `json:"evidence_candidate_ref"`
	InputEvidenceRefs                      []string `json:"input_evidence_refs,omitempty"`
	InputEvidenceHashRefs                  []string `json:"input_evidence_hash_refs,omitempty"`
	ModelOrRuleVersionRef                  string   `json:"model_or_rule_version_ref"`
	PermissionManifestHash                 string   `json:"permission_manifest_hash"`
	SandboxResultRef                       string   `json:"sandbox_result_ref"`
	ApprovalRequestRef                     string   `json:"approval_request_ref"`
	ReviewerRef                            string   `json:"reviewer_ref"`
	AuditEventRef                          string   `json:"audit_event_ref"`
	AdvisoryOnly                           bool     `json:"advisory_only"`
	EvidenceCandidateOnly                  bool     `json:"evidence_candidate_only"`
	PassAllowed                            bool     `json:"pass_allowed"`
	ApprovalGranted                        bool     `json:"approval_granted"`
	DeploymentAuthorized                   bool     `json:"deployment_authorized"`
	ProductionReadinessClaimed             bool     `json:"production_readiness_claimed"`
	ProductionMutationAllowed              bool     `json:"production_mutation_allowed"`
	CanonicalMutationAllowed               bool     `json:"canonical_mutation_allowed"`
	ExternalAPIAllowed                     bool     `json:"external_api_allowed"`
	ExternalAPIGovernanceEventRef          string   `json:"external_api_governance_event_ref"`
	AITraceCannotSatisfyExitPacketByItself bool     `json:"ai_trace_cannot_satisfy_exit_packet_by_itself"`
}

type Point13ValBNoOverclaimTrace struct {
	ObservedCustomerTexts                []string `json:"observed_customer_texts,omitempty"`
	ObservedSupportTexts                 []string `json:"observed_support_texts,omitempty"`
	ObservedExitPacketTexts              []string `json:"observed_exit_packet_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point13ValBFoundation struct {
	CurrentState                      string                                  `json:"current_state"`
	BlockingReasons                   []string                                `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites               []string                                `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer              string                                  `json:"projection_disclaimer"`
	DependencyState                   string                                  `json:"dependency_state"`
	PilotEvidenceOperationLedgerState string                                  `json:"pilot_evidence_operation_ledger_state"`
	CustomerReviewTraceState          string                                  `json:"customer_review_trace_state"`
	SupportActionTraceState           string                                  `json:"support_action_trace_state"`
	PilotExitEvidencePacketState      string                                  `json:"pilot_exit_evidence_packet_state"`
	AIEvidenceOperationTraceState     string                                  `json:"ai_evidence_operation_trace_state"`
	NoOverclaimState                  string                                  `json:"no_overclaim_state"`
	Dependency                        Point13ValBDependencySnapshot           `json:"dependency"`
	PilotEvidenceOperationLedger      Point13ValBPilotEvidenceOperationLedger `json:"pilot_evidence_operation_ledger"`
	CustomerReviewTrace               Point13ValBCustomerReviewTrace          `json:"customer_review_trace"`
	SupportActionTrace                Point13ValBSupportActionTrace           `json:"support_action_trace"`
	PilotExitEvidencePacket           Point13ValBPilotExitEvidencePacket      `json:"pilot_exit_evidence_packet"`
	AIEvidenceOperationTrace          Point13ValBAIEvidenceOperationTrace     `json:"ai_evidence_operation_trace"`
	NoOverclaimTrace                  Point13ValBNoOverclaimTrace             `json:"no_overclaim_trace"`
}

func point13ValBStates() []string {
	return []string{
		Point13ValBStateActive,
		Point13ValBStateBlocked,
		Point13ValBStateReviewRequired,
		Point13ValBStateIncomplete,
	}
}

func point13ValBStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValBStates(), value)
}

func point13ValBLedgerRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "pilot_evidence_ledger_", "ledger_")
}

func point13ValBLedgerEntryRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "ledger_entry_")
}

func point13ValBSupportTraceRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "support_trace_")
}

func point13ValBSupportActionRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "support_action_")
}

func point13ValBReviewTraceRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "review_trace_")
}

func point13ValBPacketRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "exit_packet_")
}

func point13ValBAITraceRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "ai_trace_")
}

func point13ValBSourceRefValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValAAllowedEvidenceSources(), value) ||
		point13Val0OperationalRefValid(value, "source_")
}

func point13ValBLedgerOperationTypes() []string {
	return []string{
		point13ValBLedgerOperationCustomerArtifactReceived,
		point13ValBLedgerOperationEvidenceCandidateRegister,
		point13ValBLedgerOperationCustodyVerified,
		point13ValBLedgerOperationSandboxResultRecorded,
		point13ValBLedgerOperationSupportActionRecorded,
		point13ValBLedgerOperationCustomerReviewRecorded,
		point13ValBLedgerOperationExitEvidencePacket,
	}
}

func point13ValBLedgerOperationTypeValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValBLedgerOperationTypes(), value)
}

func point13ValBSupportActionTypes() []string {
	return []string{
		point13ValBSupportActionTriage,
		point13ValBSupportActionEvidenceRequest,
		point13ValBSupportActionEvidenceClassification,
		point13ValBSupportActionSandboxFollowup,
		point13ValBSupportActionCustomerQuestion,
		point13ValBSupportActionOffboardingAssist,
	}
}

func point13ValBSupportActionTypeValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValBSupportActionTypes(), value)
}

func point13ValBAllowedSafeWording() []string {
	return []string{
		"pilot evidence operations ledger",
		"customer review trace",
		"support action trace",
		"operational readiness packet",
		"advisory ai evidence candidate",
		"evidence support for customer/auditor review",
	}
}

func point13ValBStringListValid(values []string) bool {
	return point13ValAStringListValid(values)
}

func point13ValBEvidenceRefsValid(values []string) bool {
	return point13ValAOperationalEvidenceRefsValid(values)
}

func point13ValBEvidenceHashRefsValid(values []string) bool {
	return point13ValACustomerArtifactHashRefsValid(values)
}

func point13ValBReviewTraceRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValACustomerReviewRefValid)
}

func point13ValBSupportActionRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValBSupportActionRefValid)
}

func point13ValBSupportActionTypesValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValBSupportActionTypeValid)
}

func point13ValBTokenForEvidenceRef(value string) string {
	trimmed := strings.TrimSpace(value)
	for _, prefix := range []string{"artifact_", "evidence_", "evidence_candidate_"} {
		trimmed = strings.TrimPrefix(trimmed, prefix)
	}
	return trimmed
}

func point13ValBEvidenceHashRefsMatchEvidenceRefs(evidenceRefs, evidenceHashRefs []string) bool {
	if len(evidenceRefs) == 0 || len(evidenceRefs) != len(evidenceHashRefs) {
		return false
	}
	for i := range evidenceRefs {
		refToken := point13ValBTokenForEvidenceRef(evidenceRefs[i])
		hashToken := strings.TrimSpace(strings.TrimPrefix(evidenceHashRefs[i], "evidence_hash_"))
		if refToken == "" || hashToken == "" || refToken != hashToken {
			return false
		}
	}
	return true
}

func point13ValBComputedBindingHash(parts ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(parts, "\n")))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func point13ValBComputedLedgerBindingHash(model Point13ValBPilotEvidenceOperationLedger) string {
	parts := []string{
		strings.TrimSpace(model.LedgerID),
		strings.TrimSpace(model.TenantScope),
		strings.TrimSpace(model.PilotScopeRef),
	}
	for _, entry := range model.OperationEntries {
		parts = append(parts,
			strings.TrimSpace(entry.EntryID),
			strings.TrimSpace(entry.OperationType),
			strings.TrimSpace(entry.OwnerRef),
			strings.Join(entry.EvidenceRefs, ","),
			strings.Join(entry.EvidenceHashRefs, ","),
			strings.TrimSpace(entry.AuditEventRef),
			strings.TrimSpace(entry.CustodyRef),
			strings.TrimSpace(entry.SourceRef),
			strconv.FormatBool(entry.CandidateOnly),
			strconv.FormatBool(entry.CanonicalMutationAllowed),
			strconv.FormatBool(entry.ProductionMutationAllowed),
			strconv.FormatBool(entry.PassAllowed),
		)
	}
	return point13ValBComputedBindingHash(parts...)
}

func point13ValBValAPayloadContainsPointPass(valA Point13ValAFoundation) bool {
	payload, err := json.Marshal(valA)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point13Val0BlockedPoint13PassToken)
}

func point13ValBDependencySnapshotFromUpstream(valA Point13ValAFoundation) Point13ValBDependencySnapshot {
	return Point13ValBDependencySnapshot{
		ValACurrentState:                          valA.CurrentState,
		ValADependencyState:                       valA.DependencyState,
		ValAPilotExecutionContractState:           valA.PilotExecutionContractState,
		ValACustomerIntakeEvidenceGovernanceState: valA.CustomerIntakeEvidenceGovernanceState,
		ValAPilotRunPhaseBoundaryState:            valA.PilotRunPhaseBoundaryState,
		ValASupportResponsibilityMatrixState:      valA.SupportResponsibilityMatrixState,
		ValAPilotExitReviewGateState:              valA.PilotExitReviewGateState,
		ValAAIAssistedPilotExecutionBoundaryState: valA.AIAssistedPilotExecutionBoundaryState,
		ValANoOverclaimState:                      valA.NoOverclaimState,
		ValAPointID:                               point13Val0PointID,
		ValAWaveID:                                point13ValAWaveID,
		ValADependencyComputedFromUpstream:        valA.Dependency.SnapshotFromComputedOutput,
		ValAPoint13PassSeen:                       point13ValBValAPayloadContainsPointPass(valA),
		InheritedVal0CurrentState:                 valA.Dependency.Val0CurrentState,
		InheritedPoint12CurrentState:              valA.Dependency.Point12CurrentState,
		InheritedPoint12DependencyState:           valA.Dependency.Point12DependencyState,
		InheritedPoint12PassClosureManifestState:  valA.Dependency.Point12PassClosureManifestState,
		InheritedPoint12ReviewerResult:            valA.Dependency.Point12ReviewerResult,
		InheritedTenantScope:                      valA.Dependency.Point12TenantScope,
		InheritedAIModelOrRuleVersionRef:          valA.Dependency.InheritedAIModelOrRuleVersionRef,
		InheritedAIPermissionManifestHash:         valA.Dependency.InheritedAIPermissionManifestHash,
		SnapshotFromComputedOutput:                true,
		ReviewPrerequisites:                       append([]string{}, valA.ReviewPrerequisites...),
		ValA:                                      valA,
	}
}

func point13ValBDependencySnapshotModel() Point13ValBDependencySnapshot {
	return point13ValBDependencySnapshotFromUpstream(ComputePoint13ValAFoundation(Point13ValAFoundationModel()))
}

func point13ValBDependencyStateAndReasons(model Point13ValBDependencySnapshot) (string, []string) {
	reviewReasons := []string{}
	blockedReasons := []string{}
	incompleteReasons := []string{}

	if !model.SnapshotFromComputedOutput || !model.ValADependencyComputedFromUpstream {
		blockedReasons = append(blockedReasons, "vala_dependency_not_computed_from_upstream")
	}
	if !point13ValAStateValid(model.ValACurrentState) ||
		!point13ValAStateValid(model.ValADependencyState) ||
		!point13ValAStateValid(model.ValAPilotExecutionContractState) ||
		!point13ValAStateValid(model.ValACustomerIntakeEvidenceGovernanceState) ||
		!point13ValAStateValid(model.ValAPilotRunPhaseBoundaryState) ||
		!point13ValAStateValid(model.ValASupportResponsibilityMatrixState) ||
		!point13ValAStateValid(model.ValAPilotExitReviewGateState) ||
		!point13ValAStateValid(model.ValAAIAssistedPilotExecutionBoundaryState) ||
		!point13ValAStateValid(model.ValANoOverclaimState) ||
		strings.TrimSpace(model.ValAPointID) != point13Val0PointID ||
		strings.TrimSpace(model.ValAWaveID) != point13ValAWaveID ||
		!point13Val0StateValid(model.InheritedVal0CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureManifestState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) ||
		!point12Val0VersionRefValid(model.InheritedAIModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.InheritedAIPermissionManifestHash) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_identity_invalid")
	}
	if strings.TrimSpace(model.ValACurrentState) != strings.TrimSpace(model.ValA.CurrentState) ||
		strings.TrimSpace(model.ValADependencyState) != strings.TrimSpace(model.ValA.DependencyState) ||
		strings.TrimSpace(model.ValAPilotExecutionContractState) != strings.TrimSpace(model.ValA.PilotExecutionContractState) ||
		strings.TrimSpace(model.ValACustomerIntakeEvidenceGovernanceState) != strings.TrimSpace(model.ValA.CustomerIntakeEvidenceGovernanceState) ||
		strings.TrimSpace(model.ValAPilotRunPhaseBoundaryState) != strings.TrimSpace(model.ValA.PilotRunPhaseBoundaryState) ||
		strings.TrimSpace(model.ValASupportResponsibilityMatrixState) != strings.TrimSpace(model.ValA.SupportResponsibilityMatrixState) ||
		strings.TrimSpace(model.ValAPilotExitReviewGateState) != strings.TrimSpace(model.ValA.PilotExitReviewGateState) ||
		strings.TrimSpace(model.ValAAIAssistedPilotExecutionBoundaryState) != strings.TrimSpace(model.ValA.AIAssistedPilotExecutionBoundaryState) ||
		strings.TrimSpace(model.ValANoOverclaimState) != strings.TrimSpace(model.ValA.NoOverclaimState) ||
		model.ValADependencyComputedFromUpstream != model.ValA.Dependency.SnapshotFromComputedOutput ||
		strings.TrimSpace(model.InheritedVal0CurrentState) != strings.TrimSpace(model.ValA.Dependency.Val0CurrentState) ||
		strings.TrimSpace(model.InheritedPoint12CurrentState) != strings.TrimSpace(model.ValA.Dependency.Point12CurrentState) ||
		strings.TrimSpace(model.InheritedPoint12DependencyState) != strings.TrimSpace(model.ValA.Dependency.Point12DependencyState) ||
		strings.TrimSpace(model.InheritedPoint12PassClosureManifestState) != strings.TrimSpace(model.ValA.Dependency.Point12PassClosureManifestState) ||
		strings.TrimSpace(model.InheritedPoint12ReviewerResult) != strings.TrimSpace(model.ValA.Dependency.Point12ReviewerResult) ||
		strings.TrimSpace(model.InheritedTenantScope) != strings.TrimSpace(model.ValA.Dependency.Point12TenantScope) ||
		strings.TrimSpace(model.InheritedAIModelOrRuleVersionRef) != strings.TrimSpace(model.ValA.Dependency.InheritedAIModelOrRuleVersionRef) ||
		strings.TrimSpace(model.InheritedAIPermissionManifestHash) != strings.TrimSpace(model.ValA.Dependency.InheritedAIPermissionManifestHash) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_binding_mismatch")
	}
	if model.ValAPoint13PassSeen {
		blockedReasons = append(blockedReasons, "vala_point13_pass_seen")
	}
	for _, state := range []string{
		model.ValACurrentState,
		model.ValADependencyState,
		model.ValAPilotExecutionContractState,
		model.ValACustomerIntakeEvidenceGovernanceState,
		model.ValAPilotRunPhaseBoundaryState,
		model.ValASupportResponsibilityMatrixState,
		model.ValAPilotExitReviewGateState,
		model.ValAAIAssistedPilotExecutionBoundaryState,
		model.ValANoOverclaimState,
	} {
		switch strings.TrimSpace(state) {
		case Point13ValAStateBlocked:
			blockedReasons = append(blockedReasons, "vala_component_blocked")
		case Point13ValAStateReviewRequired:
			reviewReasons = append(reviewReasons, "vala_component_review_required")
		case Point13ValAStateIncomplete:
			incompleteReasons = append(incompleteReasons, "vala_component_incomplete")
		}
	}
	switch strings.TrimSpace(model.InheritedPoint12CurrentState) {
	case Point12ValEStateBlocked:
		blockedReasons = append(blockedReasons, "point12_inherited_blocked")
	case Point12ValEStateReviewRequired:
		reviewReasons = append(reviewReasons, "point12_inherited_review_required")
	case Point12ValEStateIncomplete:
		incompleteReasons = append(incompleteReasons, "point12_inherited_incomplete")
	}
	if strings.TrimSpace(model.InheritedPoint12CurrentState) == Point12ValEStatePassConfirmed &&
		(strings.TrimSpace(model.InheritedPoint12DependencyState) != Point12ValEStateActive ||
			strings.TrimSpace(model.InheritedPoint12PassClosureManifestState) != Point12ValEStateActive ||
			strings.TrimSpace(model.InheritedPoint12ReviewerResult) != point12ValEReviewerResultPassConfirmed) {
		blockedReasons = append(blockedReasons, "point12_inherited_not_pass_confirmed")
	}
	if len(blockedReasons) > 0 {
		return Point13ValBStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point13ValBStateReviewRequired, reviewReasons
	}
	if len(incompleteReasons) > 0 {
		return Point13ValBStateIncomplete, incompleteReasons
	}
	return Point13ValBStateActive, nil
}

func point13ValBExpectedLedgerOwner(operationType string, dependency Point13ValBDependencySnapshot) string {
	switch strings.TrimSpace(operationType) {
	case point13ValBLedgerOperationCustomerArtifactReceived, point13ValBLedgerOperationCustomerReviewRecorded:
		return dependency.ValA.PilotExecutionContract.CustomerOwnerRef
	case point13ValBLedgerOperationSupportActionRecorded:
		return dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef
	default:
		return dependency.ValA.PilotExecutionContract.PilotOwnerRef
	}
}

func EvaluatePoint13ValBPilotEvidenceOperationLedgerState(model Point13ValBPilotEvidenceOperationLedger, dependency Point13ValBDependencySnapshot) string {
	if !point13ValBLedgerRefValid(model.LedgerID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValAPilotScopeRefValid(model.PilotScopeRef) ||
		len(model.OperationEntries) == 0 ||
		strings.TrimSpace(model.LedgerBindingHash) != point13ValBComputedLedgerBindingHash(model) {
		return Point13ValBStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.PilotScopeRef) != strings.TrimSpace(dependency.ValA.PilotExecutionContract.PilotScopeRef) {
		return Point13ValBStateBlocked
	}
	for _, entry := range model.OperationEntries {
		if !point13ValBLedgerEntryRefValid(entry.EntryID) ||
			!point13ValBLedgerOperationTypeValid(entry.OperationType) ||
			!point13ValAOwnerRefValid(entry.OwnerRef) ||
			!point13ValBEvidenceRefsValid(entry.EvidenceRefs) ||
			!point13ValBEvidenceHashRefsValid(entry.EvidenceHashRefs) ||
			!point13ValBEvidenceHashRefsMatchEvidenceRefs(entry.EvidenceRefs, entry.EvidenceHashRefs) ||
			!point12Val0AuditRefValid(entry.AuditEventRef) ||
			!entry.CandidateOnly ||
			entry.CanonicalMutationAllowed ||
			entry.ProductionMutationAllowed ||
			entry.PassAllowed {
			return Point13ValBStateBlocked
		}
		if strings.TrimSpace(entry.OwnerRef) != strings.TrimSpace(point13ValBExpectedLedgerOwner(entry.OperationType, dependency)) ||
			!point12Val0ExactStringSetMatch(entry.EvidenceRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs) ||
			!point12Val0ExactStringSetMatch(entry.EvidenceHashRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs) ||
			point13ValAContainsCrossTenantArtifact(entry.EvidenceRefs) {
			return Point13ValBStateBlocked
		}
		switch strings.TrimSpace(entry.OperationType) {
		case point13ValBLedgerOperationCustomerArtifactReceived, point13ValBLedgerOperationCustodyVerified:
			if !point13ValACustodyRefValid(entry.CustodyRef) {
				return Point13ValBStateBlocked
			}
		}
		switch strings.TrimSpace(entry.OperationType) {
		case point13ValBLedgerOperationCustomerArtifactReceived,
			point13ValBLedgerOperationEvidenceCandidateRegister,
			point13ValBLedgerOperationSandboxResultRecorded,
			point13ValBLedgerOperationSupportActionRecorded,
			point13ValBLedgerOperationCustomerReviewRecorded,
			point13ValBLedgerOperationExitEvidencePacket:
			if !point13ValBSourceRefValid(entry.SourceRef) {
				return Point13ValBStateBlocked
			}
		}
	}
	return Point13ValBStateActive
}

func EvaluatePoint13ValBCustomerReviewTraceState(model Point13ValBCustomerReviewTrace, dependency Point13ValBDependencySnapshot) string {
	if !point13ValBReviewTraceRefValid(model.ReviewTraceID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValBReviewTraceRefsValid(model.CustomerReviewRefs) ||
		!point13ValBEvidenceRefsValid(model.ReviewedEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.ReviewedEvidenceHashRefs) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.ReviewedEvidenceRefs, model.ReviewedEvidenceHashRefs) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point13ValAOwnerRefValid(model.InternalOwnerRef) ||
		!point13ValAOwnerRefValid(model.SupportOwnerRef) ||
		!point13ValAAuditRefsValid(model.AuditEventRefs) ||
		strings.TrimSpace(model.FinalCustomerStatement) == "" ||
		!model.CustomerReviewIsOperationalOnly ||
		!model.CustomerReviewCannotApproveProduction ||
		!model.CustomerReviewCannotCreatePass {
		return Point13ValBStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(dependency.ValA.PilotExecutionContract.CustomerOwnerRef) ||
		strings.TrimSpace(model.InternalOwnerRef) != strings.TrimSpace(dependency.ValA.PilotExecutionContract.PilotOwnerRef) ||
		strings.TrimSpace(model.SupportOwnerRef) != strings.TrimSpace(dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef) ||
		!point12Val0ExactStringSetMatch(model.ReviewedEvidenceRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs) ||
		!point12Val0ExactStringSetMatch(model.ReviewedEvidenceHashRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs) ||
		point13ValAContainsCrossTenantArtifact(model.ReviewedEvidenceRefs) {
		return Point13ValBStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(model.FinalCustomerStatement) {
		return Point13ValBStateBlocked
	}
	return Point13ValBStateActive
}

func EvaluatePoint13ValBSupportActionTraceState(model Point13ValBSupportActionTrace, dependency Point13ValBDependencySnapshot) string {
	if !point13ValBSupportTraceRefValid(model.SupportTraceID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValBSupportActionRefsValid(model.SupportActionRefs) ||
		!point13ValBSupportActionTypesValid(model.SupportActionTypes) ||
		len(model.SupportActionRefs) != len(model.SupportActionTypes) ||
		!point13ValAOwnerRefValid(model.SupportOwnerRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point13ValBEvidenceRefsValid(model.EvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.EvidenceHashRefs) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.EvidenceRefs, model.EvidenceHashRefs) ||
		!point13ValAAuditRefsValid(model.AuditEventRefs) ||
		!model.SupportCanViewCandidateArtifacts ||
		model.SupportCanMutateCanonicalEvidence ||
		model.SupportCanOverrideCoreDecision ||
		model.SupportCanApproveProduction {
		return Point13ValBStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.SupportOwnerRef) != strings.TrimSpace(dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(dependency.ValA.PilotExecutionContract.CustomerOwnerRef) ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceHashRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs) ||
		point13ValAContainsCrossTenantArtifact(model.EvidenceRefs) {
		return Point13ValBStateBlocked
	}
	return Point13ValBStateActive
}

func EvaluatePoint13ValBPilotExitEvidencePacketState(model Point13ValBPilotExitEvidencePacket, dependency Point13ValBDependencySnapshot, ledger Point13ValBPilotEvidenceOperationLedger, reviewTrace Point13ValBCustomerReviewTrace, supportTrace Point13ValBSupportActionTrace, aiTrace Point13ValBAIEvidenceOperationTrace) string {
	if !point13ValBPacketRefValid(model.PacketID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!model.OperationalReadinessOnly ||
		!point13ValBEvidenceRefsValid(model.EvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.EvidenceHashRefs) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.EvidenceRefs, model.EvidenceHashRefs) ||
		!point13ValBReviewTraceRefValid(model.CustomerReviewTraceRef) ||
		!point13ValBSupportTraceRefValid(model.SupportTraceRef) ||
		!point13ValBLedgerRefValid(model.OperationLedgerRef) ||
		!point13ValBStringListValid(model.Limitations) ||
		strings.TrimSpace(model.SafeCustomerStatement) == "" ||
		!model.NoProductionApproval ||
		!model.NoDeploymentApproval ||
		!model.NoComplianceGuarantee ||
		!model.NoCertification ||
		!model.NoFinancialGuarantee ||
		!model.NoPoint13Pass {
		return Point13ValBStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceHashRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs) ||
		strings.TrimSpace(model.CustomerReviewTraceRef) != strings.TrimSpace(reviewTrace.ReviewTraceID) ||
		strings.TrimSpace(model.SupportTraceRef) != strings.TrimSpace(supportTrace.SupportTraceID) ||
		strings.TrimSpace(model.OperationLedgerRef) != strings.TrimSpace(ledger.LedgerID) ||
		point13ValAContainsCrossTenantArtifact(model.EvidenceRefs) {
		return Point13ValBStateBlocked
	}
	if len(model.UnresolvedBlockers) > 0 ||
		point13Val0ContainsForbiddenClaim(model.SafeCustomerStatement, strings.Join(model.Limitations, " ")) ||
		strings.Contains(model.SafeCustomerStatement, point13Val0BlockedPoint13PassToken) {
		return Point13ValBStateBlocked
	}
	if len(model.EvidenceRefs) == 1 && strings.TrimSpace(model.EvidenceRefs[0]) == strings.TrimSpace(aiTrace.EvidenceCandidateRef) {
		return Point13ValBStateBlocked
	}
	return Point13ValBStateActive
}

func EvaluatePoint13ValBAIEvidenceOperationTraceState(model Point13ValBAIEvidenceOperationTrace, dependency Point13ValBDependencySnapshot, supportTrace Point13ValBSupportActionTrace) string {
	if point11Val0ContainsTrimmed(point12Val0BlockedAIEvidenceCandidateTypes(), model.AIOutputType) {
		return Point13ValBStateBlocked
	}
	expectedSupportAuditRef := ""
	if len(supportTrace.AuditEventRefs) > 0 {
		expectedSupportAuditRef = strings.TrimSpace(supportTrace.AuditEventRefs[0])
	}
	if !point13ValBAITraceRefValid(model.AITraceID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0AIEvidenceCandidateTypeValid(model.AIOutputType) ||
		!point13ValAAIEvidenceCandidateRefValid(model.EvidenceCandidateRef) ||
		!point13ValBEvidenceRefsValid(model.InputEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.InputEvidenceHashRefs) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.InputEvidenceRefs, model.InputEvidenceHashRefs) ||
		!point12Val0VersionRefValid(model.ModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.PermissionManifestHash) ||
		(model.SandboxResultRef != "" && !point13Val0OperationalRefValid(model.SandboxResultRef, "sandbox_result_")) ||
		(model.ApprovalRequestRef != "" && !point13Val0OperationalRefValid(model.ApprovalRequestRef, "approval_request_")) ||
		(model.ReviewerRef != "" && !point13ValAReviewerRefValid(model.ReviewerRef)) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		!model.AdvisoryOnly ||
		!model.EvidenceCandidateOnly ||
		!model.AITraceCannotSatisfyExitPacketByItself ||
		model.PassAllowed ||
		model.ApprovalGranted ||
		model.DeploymentAuthorized ||
		model.ProductionReadinessClaimed ||
		model.ProductionMutationAllowed ||
		model.CanonicalMutationAllowed ||
		expectedSupportAuditRef == "" {
		return Point13ValBStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs) ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceHashRefs, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs) ||
		strings.TrimSpace(model.ModelOrRuleVersionRef) != strings.TrimSpace(dependency.InheritedAIModelOrRuleVersionRef) ||
		strings.TrimSpace(model.PermissionManifestHash) != strings.TrimSpace(dependency.InheritedAIPermissionManifestHash) ||
		strings.TrimSpace(model.AuditEventRef) != expectedSupportAuditRef ||
		point13ValAContainsCrossTenantArtifact(model.InputEvidenceRefs) {
		return Point13ValBStateBlocked
	}
	if model.ExternalAPIAllowed {
		if !point12Val0GovernanceEventRefValid(model.ExternalAPIGovernanceEventRef) {
			return Point13ValBStateBlocked
		}
		return Point13ValBStateReviewRequired
	}
	return Point13ValBStateActive
}

func EvaluatePoint13ValBNoOverclaimTraceState(model Point13ValBNoOverclaimTrace) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point13ValBStringListValid(model.AllowedSafeWording) ||
		!point13ValBStringListValid(model.BlockedWording) {
		return Point13ValBStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.ObservedCustomerTexts, " "),
		strings.Join(model.ObservedSupportTexts, " "),
		strings.Join(model.ObservedExitPacketTexts, " "),
	) {
		return Point13ValBStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(strings.Join(model.InternalDiagnosticTexts, " ")) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point13ValBStateBlocked
	}
	return Point13ValBStateActive
}

type point13ValBComponentState struct {
	Name  string
	State string
}

func point13ValBComponentStates(model Point13ValBFoundation) []point13ValBComponentState {
	return []point13ValBComponentState{
		{Name: "dependency", State: model.DependencyState},
		{Name: "pilot_evidence_operation_ledger", State: model.PilotEvidenceOperationLedgerState},
		{Name: "customer_review_trace", State: model.CustomerReviewTraceState},
		{Name: "support_action_trace", State: model.SupportActionTraceState},
		{Name: "pilot_exit_evidence_packet", State: model.PilotExitEvidencePacketState},
		{Name: "ai_evidence_operation_trace", State: model.AIEvidenceOperationTraceState},
		{Name: "no_overclaim", State: model.NoOverclaimState},
	}
}

func point13ValBBlockingReasons(model Point13ValBFoundation) []string {
	reasons := []string{}
	for _, component := range point13ValBComponentStates(model) {
		if component.State == Point13ValBStateBlocked || component.State == Point13ValBStateIncomplete {
			reasons = append(reasons, component.Name+":"+component.State)
		}
	}
	return reasons
}

func EvaluatePoint13ValBState(model Point13ValBFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValBStateBlocked
	}
	hasReviewRequired := false
	hasIncomplete := false
	for _, component := range point13ValBComponentStates(model) {
		switch component.State {
		case Point13ValBStateBlocked:
			return Point13ValBStateBlocked
		case Point13ValBStateReviewRequired:
			hasReviewRequired = true
		case Point13ValBStateIncomplete:
			hasIncomplete = true
		}
	}
	if hasReviewRequired {
		return Point13ValBStateReviewRequired
	}
	if hasIncomplete {
		return Point13ValBStateIncomplete
	}
	return Point13ValBStateActive
}

func point13ValBDefaultLedgerEntries(dependency Point13ValBDependencySnapshot) []Point13ValBOperationLedgerEntry {
	evidenceRefs := append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs...)
	evidenceHashRefs := append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs...)
	return []Point13ValBOperationLedgerEntry{
		{
			EntryID:          "ledger_entry_point13_valb_001",
			OperationType:    point13ValBLedgerOperationCustomerArtifactReceived,
			OwnerRef:         dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
			EvidenceRefs:     evidenceRefs,
			EvidenceHashRefs: evidenceHashRefs,
			AuditEventRef:    "audit_point13_valb_ledger_001",
			CustodyRef:       dependency.ValA.CustomerIntakeEvidenceGovernance.CustodyRef,
			SourceRef:        point13ValAEvidenceSourceCustomerUpload,
			CandidateOnly:    true,
		},
		{
			EntryID:          "ledger_entry_point13_valb_002",
			OperationType:    point13ValBLedgerOperationEvidenceCandidateRegister,
			OwnerRef:         dependency.ValA.PilotExecutionContract.PilotOwnerRef,
			EvidenceRefs:     evidenceRefs,
			EvidenceHashRefs: evidenceHashRefs,
			AuditEventRef:    "audit_point13_valb_ledger_002",
			SourceRef:        point13ValAEvidenceSourceAuditExport,
			CandidateOnly:    true,
		},
		{
			EntryID:          "ledger_entry_point13_valb_003",
			OperationType:    point13ValBLedgerOperationCustodyVerified,
			OwnerRef:         dependency.ValA.PilotExecutionContract.PilotOwnerRef,
			EvidenceRefs:     evidenceRefs,
			EvidenceHashRefs: evidenceHashRefs,
			AuditEventRef:    "audit_point13_valb_ledger_003",
			CustodyRef:       dependency.ValA.CustomerIntakeEvidenceGovernance.CustodyRef,
			CandidateOnly:    true,
		},
		{
			EntryID:          "ledger_entry_point13_valb_004",
			OperationType:    point13ValBLedgerOperationSandboxResultRecorded,
			OwnerRef:         dependency.ValA.PilotExecutionContract.PilotOwnerRef,
			EvidenceRefs:     evidenceRefs,
			EvidenceHashRefs: evidenceHashRefs,
			AuditEventRef:    "audit_point13_valb_ledger_004",
			SourceRef:        point13ValAEvidenceSourceSandboxResult,
			CandidateOnly:    true,
		},
		{
			EntryID:          "ledger_entry_point13_valb_005",
			OperationType:    point13ValBLedgerOperationSupportActionRecorded,
			OwnerRef:         dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef,
			EvidenceRefs:     evidenceRefs,
			EvidenceHashRefs: evidenceHashRefs,
			AuditEventRef:    "audit_point13_valb_ledger_005",
			SourceRef:        point13ValAEvidenceSourceSupportAttachment,
			CandidateOnly:    true,
		},
		{
			EntryID:          "ledger_entry_point13_valb_006",
			OperationType:    point13ValBLedgerOperationCustomerReviewRecorded,
			OwnerRef:         dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
			EvidenceRefs:     evidenceRefs,
			EvidenceHashRefs: evidenceHashRefs,
			AuditEventRef:    "audit_point13_valb_ledger_006",
			SourceRef:        point13ValAEvidenceSourceAuditExport,
			CandidateOnly:    true,
		},
		{
			EntryID:          "ledger_entry_point13_valb_007",
			OperationType:    point13ValBLedgerOperationExitEvidencePacket,
			OwnerRef:         dependency.ValA.PilotExecutionContract.PilotOwnerRef,
			EvidenceRefs:     evidenceRefs,
			EvidenceHashRefs: evidenceHashRefs,
			AuditEventRef:    "audit_point13_valb_ledger_007",
			SourceRef:        point13ValAEvidenceSourceAuditExport,
			CandidateOnly:    true,
		},
	}
}

func Point13ValBFoundationModel() Point13ValBFoundation {
	disclaimer := point13ValBProjectionDisclaimerBaseline
	dependency := point13ValBDependencySnapshotModel()
	tenantScope := dependency.InheritedTenantScope
	ledger := Point13ValBPilotEvidenceOperationLedger{
		LedgerID:         "pilot_evidence_ledger_point13_valb_001",
		TenantScope:      tenantScope,
		PilotScopeRef:    dependency.ValA.PilotExecutionContract.PilotScopeRef,
		OperationEntries: point13ValBDefaultLedgerEntries(dependency),
	}
	ledger.LedgerBindingHash = point13ValBComputedLedgerBindingHash(ledger)
	customerReview := Point13ValBCustomerReviewTrace{
		ReviewTraceID:                         "review_trace_point13_valb_001",
		TenantScope:                           tenantScope,
		CustomerReviewRefs:                    []string{"customer_review_point13_valb_001"},
		ReviewedEvidenceRefs:                  append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs...),
		ReviewedEvidenceHashRefs:              append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs...),
		CustomerOwnerRef:                      dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
		InternalOwnerRef:                      dependency.ValA.PilotExecutionContract.PilotOwnerRef,
		SupportOwnerRef:                       dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef,
		AuditEventRefs:                        []string{"audit_point13_valb_customer_review_001", "audit_point13_valb_customer_review_002"},
		FinalCustomerStatement:                "operational readiness packet prepared for customer review",
		CustomerReviewIsOperationalOnly:       true,
		CustomerReviewCannotApproveProduction: true,
		CustomerReviewCannotCreatePass:        true,
	}
	supportTrace := Point13ValBSupportActionTrace{
		SupportTraceID:                   "support_trace_point13_valb_001",
		TenantScope:                      tenantScope,
		SupportActionRefs:                []string{"support_action_point13_valb_001", "support_action_point13_valb_002", "support_action_point13_valb_003", "support_action_point13_valb_004", "support_action_point13_valb_005", "support_action_point13_valb_006"},
		SupportActionTypes:               point13ValBSupportActionTypes(),
		SupportOwnerRef:                  dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef,
		CustomerOwnerRef:                 dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
		EvidenceRefs:                     append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs...),
		EvidenceHashRefs:                 append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs...),
		AuditEventRefs:                   []string{"audit_point13_valb_support_001", "audit_point13_valb_support_002"},
		SupportCanViewCandidateArtifacts: true,
	}
	aiTrace := Point13ValBAIEvidenceOperationTrace{
		AITraceID:                              "ai_trace_point13_valb_001",
		TenantScope:                            tenantScope,
		AIOutputType:                           "AI_RECOMMENDATION",
		EvidenceCandidateRef:                   "evidence_candidate_point13_valb_001",
		InputEvidenceRefs:                      append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs...),
		InputEvidenceHashRefs:                  append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs...),
		ModelOrRuleVersionRef:                  dependency.InheritedAIModelOrRuleVersionRef,
		PermissionManifestHash:                 dependency.InheritedAIPermissionManifestHash,
		ReviewerRef:                            "reviewer_point13_valb_internal_001",
		AuditEventRef:                          supportTrace.AuditEventRefs[0],
		AdvisoryOnly:                           true,
		EvidenceCandidateOnly:                  true,
		AITraceCannotSatisfyExitPacketByItself: true,
	}
	exitPacket := Point13ValBPilotExitEvidencePacket{
		PacketID:                 "exit_packet_point13_valb_001",
		TenantScope:              tenantScope,
		OperationalReadinessOnly: true,
		EvidenceRefs:             append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs...),
		EvidenceHashRefs:         append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs...),
		CustomerReviewTraceRef:   customerReview.ReviewTraceID,
		SupportTraceRef:          supportTrace.SupportTraceID,
		OperationLedgerRef:       ledger.LedgerID,
		Limitations:              []string{"operational readiness packet only"},
		SafeCustomerStatement:    "operational readiness packet prepared",
		NoProductionApproval:     true,
		NoDeploymentApproval:     true,
		NoComplianceGuarantee:    true,
		NoCertification:          true,
		NoFinancialGuarantee:     true,
		NoPoint13Pass:            true,
	}
	return Point13ValBFoundation{
		CurrentState:                      Point13ValBStateActive,
		ProjectionDisclaimer:              disclaimer,
		DependencyState:                   Point13ValBStateActive,
		PilotEvidenceOperationLedgerState: Point13ValBStateActive,
		CustomerReviewTraceState:          Point13ValBStateActive,
		SupportActionTraceState:           Point13ValBStateActive,
		PilotExitEvidencePacketState:      Point13ValBStateActive,
		AIEvidenceOperationTraceState:     Point13ValBStateActive,
		NoOverclaimState:                  Point13ValBStateActive,
		Dependency:                        dependency,
		PilotEvidenceOperationLedger:      ledger,
		CustomerReviewTrace:               customerReview,
		SupportActionTrace:                supportTrace,
		PilotExitEvidencePacket:           exitPacket,
		AIEvidenceOperationTrace:          aiTrace,
		NoOverclaimTrace: Point13ValBNoOverclaimTrace{
			ObservedCustomerTexts:                []string{"pilot evidence operations ledger", "customer review trace"},
			ObservedSupportTexts:                 []string{"support action trace"},
			ObservedExitPacketTexts:              []string{"operational readiness packet"},
			InternalDiagnosticTexts:              []string{"blocked wording remains denylisted internally"},
			InternalDiagnosticsClassifiedBlocked: true,
			AllowedSafeWording:                   point13ValBAllowedSafeWording(),
			BlockedWording:                       point13Val0ForbiddenClaims(),
			ProjectionDisclaimer:                 disclaimer,
		},
	}
}

func ComputePoint13ValBFoundation(model Point13ValBFoundation) Point13ValBFoundation {
	dependencyState, dependencyReasons := point13ValBDependencyStateAndReasons(model.Dependency)
	model.DependencyState = dependencyState
	model.PilotEvidenceOperationLedgerState = EvaluatePoint13ValBPilotEvidenceOperationLedgerState(model.PilotEvidenceOperationLedger, model.Dependency)
	model.CustomerReviewTraceState = EvaluatePoint13ValBCustomerReviewTraceState(model.CustomerReviewTrace, model.Dependency)
	model.SupportActionTraceState = EvaluatePoint13ValBSupportActionTraceState(model.SupportActionTrace, model.Dependency)
	model.AIEvidenceOperationTraceState = EvaluatePoint13ValBAIEvidenceOperationTraceState(model.AIEvidenceOperationTrace, model.Dependency, model.SupportActionTrace)
	model.PilotExitEvidencePacketState = EvaluatePoint13ValBPilotExitEvidencePacketState(model.PilotExitEvidencePacket, model.Dependency, model.PilotEvidenceOperationLedger, model.CustomerReviewTrace, model.SupportActionTrace, model.AIEvidenceOperationTrace)
	model.NoOverclaimState = EvaluatePoint13ValBNoOverclaimTraceState(model.NoOverclaimTrace)
	model.CurrentState = EvaluatePoint13ValBState(model)
	model.BlockingReasons = point13ValBBlockingReasons(model)
	model.ReviewPrerequisites = nil
	if model.DependencyState == Point13ValBStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, dependencyReasons...)
	}
	if model.AIEvidenceOperationTraceState == Point13ValBStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "ai_evidence_operation_trace_requires_governance_review")
	}
	return model
}
