package formal

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

const (
	Point13ValCStateActive         = "point13_valc_customer_export_handoff_active"
	Point13ValCStateBlocked        = "point13_valc_customer_export_handoff_blocked"
	Point13ValCStateReviewRequired = "point13_valc_customer_export_handoff_review_required"
	Point13ValCStateIncomplete     = "point13_valc_customer_export_handoff_incomplete"
)

const (
	point13ValCWaveID                       = "val_c"
	point13ValCPreviousWaveID               = point13ValBWaveID
	point13ValCProjectionDisclaimerBaseline = "projection_only not_canonical_truth point13_valc_customer_export_handoff"
)

type Point13ValCDependencySnapshot struct {
	ValBCurrentState                      string                `json:"valb_current_state"`
	ValBDependencyState                   string                `json:"valb_dependency_state"`
	ValBPilotEvidenceOperationLedgerState string                `json:"valb_pilot_evidence_operation_ledger_state"`
	ValBCustomerReviewTraceState          string                `json:"valb_customer_review_trace_state"`
	ValBSupportActionTraceState           string                `json:"valb_support_action_trace_state"`
	ValBPilotExitEvidencePacketState      string                `json:"valb_pilot_exit_evidence_packet_state"`
	ValBAIEvidenceOperationTraceState     string                `json:"valb_ai_evidence_operation_trace_state"`
	ValBNoOverclaimState                  string                `json:"valb_no_overclaim_state"`
	ValBPointID                           string                `json:"valb_point_id"`
	ValBWaveID                            string                `json:"valb_wave_id"`
	ValBDependencyComputedFromUpstream    bool                  `json:"valb_dependency_computed_from_upstream"`
	ValBPoint13PassSeen                   bool                  `json:"valb_point13_pass_seen"`
	InheritedValACurrentState             string                `json:"inherited_vala_current_state"`
	InheritedVal0CurrentState             string                `json:"inherited_val0_current_state"`
	InheritedPoint12CurrentState          string                `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState       string                `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState      string                `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult        string                `json:"inherited_point12_reviewer_result"`
	InheritedTenantScope                  string                `json:"inherited_tenant_scope"`
	InheritedAIModelOrRuleVersionRef      string                `json:"inherited_ai_model_or_rule_version_ref"`
	InheritedAIPermissionManifestHash     string                `json:"inherited_ai_permission_manifest_hash"`
	SnapshotFromComputedOutput            bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                   []string              `json:"review_prerequisites,omitempty"`
	ValB                                  Point13ValBFoundation `json:"valb"`
}

type Point13ValCCustomerEvidenceExportPackage struct {
	ExportPackageID                     string   `json:"export_package_id"`
	TenantScope                         string   `json:"tenant_scope"`
	OperationLedgerRef                  string   `json:"operation_ledger_ref"`
	CustomerReviewTraceRef              string   `json:"customer_review_trace_ref"`
	SupportTraceRef                     string   `json:"support_trace_ref"`
	ExitEvidencePacketRef               string   `json:"exit_evidence_packet_ref"`
	ExportedEvidenceRefs                []string `json:"exported_evidence_refs,omitempty"`
	ExportedEvidenceHashRefs            []string `json:"exported_evidence_hash_refs,omitempty"`
	ExportManifestHash                  string   `json:"export_manifest_hash"`
	RetentionClassRef                   string   `json:"retention_class_ref"`
	ExportOwnerRef                      string   `json:"export_owner_ref"`
	CustomerOwnerRef                    string   `json:"customer_owner_ref"`
	AuditEventRef                       string   `json:"audit_event_ref"`
	PublicPrivateClassification         string   `json:"public_private_classification"`
	ExportIsReadOnly                    bool     `json:"export_is_read_only"`
	ExportIsOperationalEvidenceOnly     bool     `json:"export_is_operational_evidence_only"`
	ExportCannotCreatePass              bool     `json:"export_cannot_create_pass"`
	ExportCannotApproveProduction       bool     `json:"export_cannot_approve_production"`
	ExportCannotCertify                 bool     `json:"export_cannot_certify"`
	ExportCannotMutateCanonicalEvidence bool     `json:"export_cannot_mutate_canonical_evidence"`
}

type Point13ValCRedactionSafeDisclosureBoundary struct {
	RedactionBoundaryID                     string   `json:"redaction_boundary_id"`
	TenantScope                             string   `json:"tenant_scope"`
	ExportPackageID                         string   `json:"export_package_id"`
	RedactionManifestRef                    string   `json:"redaction_manifest_ref"`
	RedactedFields                          []string `json:"redacted_fields,omitempty"`
	RedactionReasons                        []string `json:"redaction_reasons,omitempty"`
	RedactionApproverRef                    string   `json:"redaction_approver_ref"`
	RedactionAuditEventRef                  string   `json:"redaction_audit_event_ref"`
	RedactionAffectsDecision                bool     `json:"redaction_affects_decision"`
	RedactionAffectsReplay                  bool     `json:"redaction_affects_replay"`
	MinimumSafeStatement                    string   `json:"minimum_safe_statement"`
	DisallowedCustomerClaims                []string `json:"disallowed_customer_claims,omitempty"`
	SurvivingCustomerClaims                 []string `json:"surviving_customer_claims,omitempty"`
	DecisiveEvidenceRemoved                 bool     `json:"decisive_evidence_removed"`
	RedactionCannotStrengthenClaim          bool     `json:"redaction_cannot_strengthen_claim"`
	RedactionCannotHideDecisiveMissingProof bool     `json:"redaction_cannot_hide_decisive_missing_evidence"`
}

type Point13ValCOperationalHandoffChecklist struct {
	HandoffChecklistID               string   `json:"handoff_checklist_id"`
	TenantScope                      string   `json:"tenant_scope"`
	HandoffOwnerRef                  string   `json:"handoff_owner_ref"`
	CustomerOwnerRef                 string   `json:"customer_owner_ref"`
	SupportOwnerRef                  string   `json:"support_owner_ref"`
	ExportPackageRef                 string   `json:"export_package_ref"`
	ExitPacketRef                    string   `json:"exit_packet_ref"`
	ChecklistItems                   []string `json:"checklist_items,omitempty"`
	RequiredAckRefs                  []string `json:"required_ack_refs,omitempty"`
	AuditEventRefs                   []string `json:"audit_event_refs,omitempty"`
	ChecklistBindingHash             string   `json:"checklist_binding_hash"`
	HandoffIsOperationalOnly         bool     `json:"handoff_is_operational_only"`
	HandoffCannotApproveProduction   bool     `json:"handoff_cannot_approve_production"`
	HandoffCannotAuthorizeDeployment bool     `json:"handoff_cannot_authorize_deployment"`
	HandoffCannotCreatePass          bool     `json:"handoff_cannot_create_pass"`
}

type Point13ValCCustomerAcceptanceTrace struct {
	AcceptanceTraceID                 string   `json:"acceptance_trace_id"`
	TenantScope                       string   `json:"tenant_scope"`
	CustomerOwnerRef                  string   `json:"customer_owner_ref"`
	ExportPackageRef                  string   `json:"export_package_ref"`
	HandoffChecklistRef               string   `json:"handoff_checklist_ref"`
	AcceptedLimitations               []string `json:"accepted_limitations,omitempty"`
	RejectedItems                     []string `json:"rejected_items,omitempty"`
	CustomerQuestions                 []string `json:"customer_questions,omitempty"`
	ResponseRefs                      []string `json:"response_refs,omitempty"`
	AuditEventRefs                    []string `json:"audit_event_refs,omitempty"`
	AcceptanceIsNotProductionApproval bool     `json:"acceptance_is_not_production_approval"`
	AcceptanceIsNotComplianceAttest   bool     `json:"acceptance_is_not_compliance_attestation"`
	AcceptanceCannotCreatePass        bool     `json:"acceptance_cannot_create_pass"`
}

type Point13ValCSupportOffboardingHandoffPacket struct {
	SupportOffboardingPacketID                string   `json:"support_offboarding_packet_id"`
	TenantScope                               string   `json:"tenant_scope"`
	SupportOwnerRef                           string   `json:"support_owner_ref"`
	CustomerOwnerRef                          string   `json:"customer_owner_ref"`
	RetentionOwnerRef                         string   `json:"retention_owner_ref"`
	ExportPackageRef                          string   `json:"export_package_ref"`
	SupportTraceRef                           string   `json:"support_trace_ref"`
	OffboardingPlanRef                        string   `json:"offboarding_plan_ref"`
	DisposalPathRef                           string   `json:"disposal_path_ref"`
	RetentionClassRefs                        []string `json:"retention_class_refs,omitempty"`
	AuditEventRefs                            []string `json:"audit_event_refs,omitempty"`
	SupportMaterialCandidateOnly              bool     `json:"support_material_candidate_only"`
	SupportOffboardingCannotMutateCanonical   bool     `json:"support_offboarding_cannot_mutate_canonical_evidence"`
	SupportOffboardingCannotOverrideDecision  bool     `json:"support_offboarding_cannot_override_core_decision"`
	SupportOffboardingCannotApproveProduction bool     `json:"support_offboarding_cannot_approve_production"`
	IndefiniteRetentionRequested              bool     `json:"indefinite_retention_requested"`
	RetentionGovernanceEventRef               string   `json:"retention_governance_event_ref"`
}

type Point13ValCAIEvidenceExportLineageSummary struct {
	AIExportSummaryID                        string   `json:"ai_export_summary_id"`
	TenantScope                              string   `json:"tenant_scope"`
	AITraceRef                               string   `json:"ai_trace_ref"`
	AIOutputType                             string   `json:"ai_output_type"`
	EvidenceCandidateRef                     string   `json:"evidence_candidate_ref"`
	InputEvidenceRefs                        []string `json:"input_evidence_refs,omitempty"`
	InputEvidenceHashRefs                    []string `json:"input_evidence_hash_refs,omitempty"`
	ModelOrRuleVersionRef                    string   `json:"model_or_rule_version_ref"`
	PermissionManifestHash                   string   `json:"permission_manifest_hash"`
	AuditEventRef                            string   `json:"audit_event_ref"`
	IncludedInExport                         bool     `json:"included_in_export"`
	AdvisoryOnly                             bool     `json:"advisory_only"`
	EvidenceCandidateOnly                    bool     `json:"evidence_candidate_only"`
	PassAllowed                              bool     `json:"pass_allowed"`
	ApprovalGranted                          bool     `json:"approval_granted"`
	DeploymentAuthorized                     bool     `json:"deployment_authorized"`
	ProductionReadinessClaimed               bool     `json:"production_readiness_claimed"`
	ProductionMutationAllowed                bool     `json:"production_mutation_allowed"`
	CanonicalMutationAllowed                 bool     `json:"canonical_mutation_allowed"`
	ExternalAPIAllowed                       bool     `json:"external_api_allowed"`
	ExternalAPIGovernanceEventRef            string   `json:"external_api_governance_event_ref"`
	AISummaryCannotStrengthenExportClaim     bool     `json:"ai_summary_cannot_strengthen_export_claim"`
	AISummaryCannotSatisfyAcceptanceByItself bool     `json:"ai_summary_cannot_satisfy_customer_acceptance_by_itself"`
}

type Point13ValCNoOverclaimExportWording struct {
	ObservedCustomerExportTexts          []string `json:"observed_customer_export_texts,omitempty"`
	ObservedHandoffTexts                 []string `json:"observed_handoff_texts,omitempty"`
	ObservedAcceptanceTexts              []string `json:"observed_acceptance_texts,omitempty"`
	ObservedSupportOffboardingTexts      []string `json:"observed_support_offboarding_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point13ValCFoundation struct {
	CurrentState                       string                                     `json:"current_state"`
	BlockingReasons                    []string                                   `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                []string                                   `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer               string                                     `json:"projection_disclaimer"`
	DependencyState                    string                                     `json:"dependency_state"`
	CustomerEvidenceExportPackageState string                                     `json:"customer_evidence_export_package_state"`
	RedactionSafeDisclosureState       string                                     `json:"redaction_safe_disclosure_state"`
	OperationalHandoffChecklistState   string                                     `json:"operational_handoff_checklist_state"`
	CustomerAcceptanceTraceState       string                                     `json:"customer_acceptance_trace_state"`
	SupportOffboardingHandoffState     string                                     `json:"support_offboarding_handoff_state"`
	AIEvidenceExportLineageState       string                                     `json:"ai_evidence_export_lineage_state"`
	NoOverclaimState                   string                                     `json:"no_overclaim_state"`
	Dependency                         Point13ValCDependencySnapshot              `json:"dependency"`
	CustomerEvidenceExportPackage      Point13ValCCustomerEvidenceExportPackage   `json:"customer_evidence_export_package"`
	RedactionSafeDisclosureBoundary    Point13ValCRedactionSafeDisclosureBoundary `json:"redaction_safe_disclosure_boundary"`
	OperationalHandoffChecklist        Point13ValCOperationalHandoffChecklist     `json:"operational_handoff_checklist"`
	CustomerAcceptanceTrace            Point13ValCCustomerAcceptanceTrace         `json:"customer_acceptance_trace"`
	SupportOffboardingHandoffPacket    Point13ValCSupportOffboardingHandoffPacket `json:"support_offboarding_handoff_packet"`
	AIEvidenceExportLineageSummary     Point13ValCAIEvidenceExportLineageSummary  `json:"ai_evidence_export_lineage_summary"`
	NoOverclaimExportWording           Point13ValCNoOverclaimExportWording        `json:"no_overclaim_export_wording"`
}

func point13ValCStates() []string {
	return []string{
		Point13ValCStateActive,
		Point13ValCStateBlocked,
		Point13ValCStateReviewRequired,
		Point13ValCStateIncomplete,
	}
}

func point13ValCStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValCStates(), value)
}

func point13ValCExportPackageRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "export_package_")
}

func point13ValCRedactionBoundaryRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "redaction_boundary_")
}

func point13ValCHandoffChecklistRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "handoff_checklist_")
}

func point13ValCAcceptanceTraceRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "acceptance_trace_")
}

func point13ValCSupportOffboardingPacketRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "support_offboarding_packet_")
}

func point13ValCAIExportSummaryRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "ai_export_summary_")
}

func point13ValCAckRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "ack_")
}

func point13ValCResponseRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "response_")
}

func point13ValCOffboardingPlanRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "offboarding_plan_")
}

func point13ValCAllowedSafeWording() []string {
	return []string{
		"customer evidence export package",
		"operational handoff checklist",
		"customer acceptance trace",
		"support offboarding handoff",
		"advisory ai evidence candidate",
		"evidence support for customer/auditor review",
	}
}

func point13ValCTextListValid(values []string) bool {
	seen := map[string]struct{}{}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return false
		}
		key := strings.ToLower(trimmed)
		if _, exists := seen[key]; exists {
			return false
		}
		seen[key] = struct{}{}
	}
	return true
}

func point13ValCChecklistItemsValid(values []string) bool {
	return point12Val0StringListValid(values, point11Val0IdentityValueValid)
}

func point13ValCAckRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValCAckRefValid)
}

func point13ValCResponseRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point13ValCResponseRefValid)
}

func point13ValCResponseCoverageValid(questions, responses []string) bool {
	if len(questions) == 0 {
		return len(responses) == 0
	}
	return len(questions) == len(responses) && point13ValCResponseRefsValid(responses)
}

func point13ValCRetentionClassRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point12Val0RetentionClassRefValid)
}

func point13ValCClaimsOverlap(left, right []string) bool {
	rightSeen := map[string]struct{}{}
	for _, value := range right {
		rightSeen[strings.ToLower(strings.TrimSpace(value))] = struct{}{}
	}
	for _, value := range left {
		if _, exists := rightSeen[strings.ToLower(strings.TrimSpace(value))]; exists {
			return true
		}
	}
	return false
}

func point13ValCComputedBindingHash(parts ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(parts, "\n")))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func point13ValCComputedExportManifestHash(model Point13ValCCustomerEvidenceExportPackage) string {
	return point13ValCComputedBindingHash(
		strings.TrimSpace(model.ExportPackageID),
		strings.TrimSpace(model.TenantScope),
		strings.TrimSpace(model.OperationLedgerRef),
		strings.TrimSpace(model.CustomerReviewTraceRef),
		strings.TrimSpace(model.SupportTraceRef),
		strings.TrimSpace(model.ExitEvidencePacketRef),
		strings.Join(model.ExportedEvidenceRefs, ","),
		strings.Join(model.ExportedEvidenceHashRefs, ","),
		strings.TrimSpace(model.RetentionClassRef),
		strings.TrimSpace(model.ExportOwnerRef),
		strings.TrimSpace(model.CustomerOwnerRef),
		strings.TrimSpace(model.AuditEventRef),
		strings.TrimSpace(model.PublicPrivateClassification),
		strconv.FormatBool(model.ExportIsReadOnly),
		strconv.FormatBool(model.ExportIsOperationalEvidenceOnly),
		strconv.FormatBool(model.ExportCannotCreatePass),
		strconv.FormatBool(model.ExportCannotApproveProduction),
		strconv.FormatBool(model.ExportCannotCertify),
		strconv.FormatBool(model.ExportCannotMutateCanonicalEvidence),
	)
}

func point13ValCComputedChecklistBindingHash(model Point13ValCOperationalHandoffChecklist) string {
	return point13ValCComputedBindingHash(
		strings.TrimSpace(model.HandoffChecklistID),
		strings.TrimSpace(model.TenantScope),
		strings.TrimSpace(model.HandoffOwnerRef),
		strings.TrimSpace(model.CustomerOwnerRef),
		strings.TrimSpace(model.SupportOwnerRef),
		strings.TrimSpace(model.ExportPackageRef),
		strings.TrimSpace(model.ExitPacketRef),
		strings.Join(model.ChecklistItems, ","),
		strings.Join(model.RequiredAckRefs, ","),
		strings.Join(model.AuditEventRefs, ","),
		strconv.FormatBool(model.HandoffIsOperationalOnly),
		strconv.FormatBool(model.HandoffCannotApproveProduction),
		strconv.FormatBool(model.HandoffCannotAuthorizeDeployment),
		strconv.FormatBool(model.HandoffCannotCreatePass),
	)
}

func point13ValCDefaultChecklistItems() []string {
	return []string{
		"handoff_scope_confirmed",
		"retention_class_disclosed",
		"limitations_acknowledged",
	}
}

func point13ValCDefaultRequiredAckRefs() []string {
	return []string{"ack_point13_valc_001"}
}

func point13ValCDefaultHandoffAuditRefs() []string {
	return []string{
		"audit_point13_valc_handoff_001",
		"audit_point13_valc_handoff_002",
	}
}

func point13ValCValBPayloadContainsPointPass(valB Point13ValBFoundation) bool {
	payload, err := json.Marshal(valB)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point13Val0BlockedPoint13PassToken)
}

func point13ValCDependencySnapshotFromUpstream(valB Point13ValBFoundation) Point13ValCDependencySnapshot {
	return Point13ValCDependencySnapshot{
		ValBCurrentState:                      valB.CurrentState,
		ValBDependencyState:                   valB.DependencyState,
		ValBPilotEvidenceOperationLedgerState: valB.PilotEvidenceOperationLedgerState,
		ValBCustomerReviewTraceState:          valB.CustomerReviewTraceState,
		ValBSupportActionTraceState:           valB.SupportActionTraceState,
		ValBPilotExitEvidencePacketState:      valB.PilotExitEvidencePacketState,
		ValBAIEvidenceOperationTraceState:     valB.AIEvidenceOperationTraceState,
		ValBNoOverclaimState:                  valB.NoOverclaimState,
		ValBPointID:                           point13Val0PointID,
		ValBWaveID:                            point13ValBWaveID,
		ValBDependencyComputedFromUpstream:    valB.Dependency.SnapshotFromComputedOutput,
		ValBPoint13PassSeen:                   point13ValCValBPayloadContainsPointPass(valB),
		InheritedValACurrentState:             valB.Dependency.ValACurrentState,
		InheritedVal0CurrentState:             valB.Dependency.InheritedVal0CurrentState,
		InheritedPoint12CurrentState:          valB.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:       valB.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:      valB.Dependency.InheritedPoint12PassClosureManifestState,
		InheritedPoint12ReviewerResult:        valB.Dependency.InheritedPoint12ReviewerResult,
		InheritedTenantScope:                  valB.Dependency.InheritedTenantScope,
		InheritedAIModelOrRuleVersionRef:      valB.Dependency.InheritedAIModelOrRuleVersionRef,
		InheritedAIPermissionManifestHash:     valB.Dependency.InheritedAIPermissionManifestHash,
		SnapshotFromComputedOutput:            true,
		ReviewPrerequisites:                   append([]string{}, valB.ReviewPrerequisites...),
		ValB:                                  valB,
	}
}

func point13ValCDependencySnapshotModel() Point13ValCDependencySnapshot {
	return point13ValCDependencySnapshotFromUpstream(ComputePoint13ValBFoundation(Point13ValBFoundationModel()))
}

func point13ValCDependencyStateAndReasons(model Point13ValCDependencySnapshot) (string, []string) {
	reviewReasons := []string{}
	blockedReasons := []string{}
	incompleteReasons := []string{}

	if !model.SnapshotFromComputedOutput || !model.ValBDependencyComputedFromUpstream {
		blockedReasons = append(blockedReasons, "valb_dependency_not_computed_from_upstream")
	}
	if !point13ValBStateValid(model.ValBCurrentState) ||
		!point13ValBStateValid(model.ValBDependencyState) ||
		!point13ValBStateValid(model.ValBPilotEvidenceOperationLedgerState) ||
		!point13ValBStateValid(model.ValBCustomerReviewTraceState) ||
		!point13ValBStateValid(model.ValBSupportActionTraceState) ||
		!point13ValBStateValid(model.ValBPilotExitEvidencePacketState) ||
		!point13ValBStateValid(model.ValBAIEvidenceOperationTraceState) ||
		!point13ValBStateValid(model.ValBNoOverclaimState) ||
		strings.TrimSpace(model.ValBPointID) != point13Val0PointID ||
		strings.TrimSpace(model.ValBWaveID) != point13ValBWaveID ||
		!point13ValAStateValid(model.InheritedValACurrentState) ||
		!point13Val0StateValid(model.InheritedVal0CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) ||
		!point12Val0VersionRefValid(model.InheritedAIModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.InheritedAIPermissionManifestHash) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_identity_invalid")
	}
	if strings.TrimSpace(model.ValBCurrentState) != strings.TrimSpace(model.ValB.CurrentState) ||
		strings.TrimSpace(model.ValBDependencyState) != strings.TrimSpace(model.ValB.DependencyState) ||
		strings.TrimSpace(model.ValBPilotEvidenceOperationLedgerState) != strings.TrimSpace(model.ValB.PilotEvidenceOperationLedgerState) ||
		strings.TrimSpace(model.ValBCustomerReviewTraceState) != strings.TrimSpace(model.ValB.CustomerReviewTraceState) ||
		strings.TrimSpace(model.ValBSupportActionTraceState) != strings.TrimSpace(model.ValB.SupportActionTraceState) ||
		strings.TrimSpace(model.ValBPilotExitEvidencePacketState) != strings.TrimSpace(model.ValB.PilotExitEvidencePacketState) ||
		strings.TrimSpace(model.ValBAIEvidenceOperationTraceState) != strings.TrimSpace(model.ValB.AIEvidenceOperationTraceState) ||
		strings.TrimSpace(model.ValBNoOverclaimState) != strings.TrimSpace(model.ValB.NoOverclaimState) ||
		model.ValBDependencyComputedFromUpstream != model.ValB.Dependency.SnapshotFromComputedOutput ||
		strings.TrimSpace(model.InheritedValACurrentState) != strings.TrimSpace(model.ValB.Dependency.ValACurrentState) ||
		strings.TrimSpace(model.InheritedVal0CurrentState) != strings.TrimSpace(model.ValB.Dependency.InheritedVal0CurrentState) ||
		strings.TrimSpace(model.InheritedPoint12CurrentState) != strings.TrimSpace(model.ValB.Dependency.InheritedPoint12CurrentState) ||
		strings.TrimSpace(model.InheritedPoint12DependencyState) != strings.TrimSpace(model.ValB.Dependency.InheritedPoint12DependencyState) ||
		strings.TrimSpace(model.InheritedPoint12PassClosureState) != strings.TrimSpace(model.ValB.Dependency.InheritedPoint12PassClosureManifestState) ||
		strings.TrimSpace(model.InheritedPoint12ReviewerResult) != strings.TrimSpace(model.ValB.Dependency.InheritedPoint12ReviewerResult) ||
		strings.TrimSpace(model.InheritedTenantScope) != strings.TrimSpace(model.ValB.Dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.InheritedAIModelOrRuleVersionRef) != strings.TrimSpace(model.ValB.Dependency.InheritedAIModelOrRuleVersionRef) ||
		strings.TrimSpace(model.InheritedAIPermissionManifestHash) != strings.TrimSpace(model.ValB.Dependency.InheritedAIPermissionManifestHash) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_binding_mismatch")
	}
	if model.ValBPoint13PassSeen {
		blockedReasons = append(blockedReasons, "valb_point13_pass_seen")
	}
	for _, state := range []string{
		model.ValBCurrentState,
		model.ValBDependencyState,
		model.ValBPilotEvidenceOperationLedgerState,
		model.ValBCustomerReviewTraceState,
		model.ValBSupportActionTraceState,
		model.ValBPilotExitEvidencePacketState,
		model.ValBAIEvidenceOperationTraceState,
		model.ValBNoOverclaimState,
	} {
		switch strings.TrimSpace(state) {
		case Point13ValBStateBlocked:
			blockedReasons = append(blockedReasons, "valb_component_blocked")
		case Point13ValBStateReviewRequired:
			reviewReasons = append(reviewReasons, "valb_component_review_required")
		case Point13ValBStateIncomplete:
			incompleteReasons = append(incompleteReasons, "valb_component_incomplete")
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
			strings.TrimSpace(model.InheritedPoint12PassClosureState) != Point12ValEStateActive ||
			strings.TrimSpace(model.InheritedPoint12ReviewerResult) != point12ValEReviewerResultPassConfirmed) {
		blockedReasons = append(blockedReasons, "point12_inherited_not_pass_confirmed")
	}
	if len(blockedReasons) > 0 {
		return Point13ValCStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point13ValCStateReviewRequired, reviewReasons
	}
	if len(incompleteReasons) > 0 {
		return Point13ValCStateIncomplete, incompleteReasons
	}
	return Point13ValCStateActive, nil
}

func EvaluatePoint13ValCCustomerEvidenceExportPackageState(model Point13ValCCustomerEvidenceExportPackage, dependency Point13ValCDependencySnapshot) string {
	if !point13ValCExportPackageRefValid(model.ExportPackageID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValBLedgerRefValid(model.OperationLedgerRef) ||
		!point13ValBReviewTraceRefValid(model.CustomerReviewTraceRef) ||
		!point13ValBSupportTraceRefValid(model.SupportTraceRef) ||
		!point13ValBPacketRefValid(model.ExitEvidencePacketRef) ||
		!point13ValBEvidenceRefsValid(model.ExportedEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.ExportedEvidenceHashRefs) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.ExportedEvidenceRefs, model.ExportedEvidenceHashRefs) ||
		strings.TrimSpace(model.ExportManifestHash) != point13ValCComputedExportManifestHash(model) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point13ValAOwnerRefValid(model.ExportOwnerRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		!point12ValCPublicPrivateClassificationValid(model.PublicPrivateClassification) ||
		!model.ExportIsReadOnly ||
		!model.ExportIsOperationalEvidenceOnly ||
		!model.ExportCannotCreatePass ||
		!model.ExportCannotApproveProduction ||
		!model.ExportCannotCertify ||
		!model.ExportCannotMutateCanonicalEvidence {
		return Point13ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.OperationLedgerRef) != strings.TrimSpace(dependency.ValB.PilotEvidenceOperationLedger.LedgerID) ||
		strings.TrimSpace(model.CustomerReviewTraceRef) != strings.TrimSpace(dependency.ValB.CustomerReviewTrace.ReviewTraceID) ||
		strings.TrimSpace(model.SupportTraceRef) != strings.TrimSpace(dependency.ValB.SupportActionTrace.SupportTraceID) ||
		strings.TrimSpace(model.ExitEvidencePacketRef) != strings.TrimSpace(dependency.ValB.PilotExitEvidencePacket.PacketID) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceRefs, dependency.ValB.Dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceHashRefs, dependency.ValB.Dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs) ||
		strings.TrimSpace(model.ExportOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.PilotExecutionContract.PilotOwnerRef) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef) ||
		strings.TrimSpace(model.RetentionClassRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.EvidenceRetentionClassRef) ||
		point13ValAContainsCrossTenantArtifact(model.ExportedEvidenceRefs) {
		return Point13ValCStateBlocked
	}
	return Point13ValCStateActive
}

func EvaluatePoint13ValCRedactionSafeDisclosureState(model Point13ValCRedactionSafeDisclosureBoundary, dependency Point13ValCDependencySnapshot, exportState string, exportPackage Point13ValCCustomerEvidenceExportPackage) string {
	if !point13ValCRedactionBoundaryRefValid(model.RedactionBoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValCExportPackageRefValid(model.ExportPackageID) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestRef) ||
		!point12Val0RedactionFieldValuesValid(model.RedactedFields) ||
		!point13ValAStringListValid(model.RedactionReasons) ||
		!point11Val0IdentityValueValid(strings.TrimSpace(model.RedactionApproverRef)) ||
		!point12Val0AuditRefValid(model.RedactionAuditEventRef) ||
		strings.TrimSpace(model.MinimumSafeStatement) == "" ||
		!point12Val0OptionalClaimTextListValid(model.DisallowedCustomerClaims) ||
		!point12Val0OptionalClaimTextListValid(model.SurvivingCustomerClaims) ||
		!model.RedactionCannotStrengthenClaim ||
		!model.RedactionCannotHideDecisiveMissingProof {
		return Point13ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.ExportPackageID) != strings.TrimSpace(exportPackage.ExportPackageID) {
		return Point13ValCStateBlocked
	}
	if strings.TrimSpace(exportState) != Point13ValCStateActive {
		return Point13ValCStateBlocked
	}
	if model.DecisiveEvidenceRemoved || point13ValCClaimsOverlap(model.SurvivingCustomerClaims, model.DisallowedCustomerClaims) {
		return Point13ValCStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(model.MinimumSafeStatement, strings.Join(model.SurvivingCustomerClaims, " ")) {
		return Point13ValCStateBlocked
	}
	if model.RedactionAffectsDecision || model.RedactionAffectsReplay {
		return Point13ValCStateReviewRequired
	}
	return Point13ValCStateActive
}

func EvaluatePoint13ValCOperationalHandoffChecklistState(model Point13ValCOperationalHandoffChecklist, dependency Point13ValCDependencySnapshot, exportPackage Point13ValCCustomerEvidenceExportPackage) string {
	if !point13ValCHandoffChecklistRefValid(model.HandoffChecklistID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValAOwnerRefValid(model.HandoffOwnerRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point13ValAOwnerRefValid(model.SupportOwnerRef) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValBPacketRefValid(model.ExitPacketRef) ||
		!point13ValCChecklistItemsValid(model.ChecklistItems) ||
		len(model.ChecklistItems) == 0 ||
		!point13ValCAckRefsValid(model.RequiredAckRefs) ||
		len(model.RequiredAckRefs) == 0 ||
		!point13ValAAuditRefsValid(model.AuditEventRefs) ||
		len(model.AuditEventRefs) == 0 ||
		strings.TrimSpace(model.ChecklistBindingHash) != point13ValCComputedChecklistBindingHash(model) ||
		!model.HandoffIsOperationalOnly ||
		!model.HandoffCannotApproveProduction ||
		!model.HandoffCannotAuthorizeDeployment ||
		!model.HandoffCannotCreatePass {
		return Point13ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.HandoffOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.PilotExecutionContract.PilotOwnerRef) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef) ||
		strings.TrimSpace(model.SupportOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef) ||
		strings.TrimSpace(model.ExportPackageRef) != strings.TrimSpace(exportPackage.ExportPackageID) ||
		strings.TrimSpace(model.ExitPacketRef) != strings.TrimSpace(dependency.ValB.PilotExitEvidencePacket.PacketID) ||
		!point12Val0ExactStringSetMatch(model.ChecklistItems, point13ValCDefaultChecklistItems()) ||
		!point12Val0ExactStringSetMatch(model.RequiredAckRefs, point13ValCDefaultRequiredAckRefs()) ||
		!point12Val0ExactStringSetMatch(model.AuditEventRefs, point13ValCDefaultHandoffAuditRefs()) {
		return Point13ValCStateBlocked
	}
	return Point13ValCStateActive
}

func EvaluatePoint13ValCCustomerAcceptanceTraceState(model Point13ValCCustomerAcceptanceTrace, dependency Point13ValCDependencySnapshot, exportPackage Point13ValCCustomerEvidenceExportPackage, handoff Point13ValCOperationalHandoffChecklist) string {
	if !point13ValCAcceptanceTraceRefValid(model.AcceptanceTraceID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValCHandoffChecklistRefValid(model.HandoffChecklistRef) ||
		!point13ValCTextListValid(model.AcceptedLimitations) ||
		!point13ValCTextListValid(model.RejectedItems) ||
		!point13ValCTextListValid(model.CustomerQuestions) ||
		!point13ValCResponseCoverageValid(model.CustomerQuestions, model.ResponseRefs) ||
		!point13ValAAuditRefsValid(model.AuditEventRefs) ||
		len(model.AuditEventRefs) == 0 ||
		!model.AcceptanceIsNotProductionApproval ||
		!model.AcceptanceIsNotComplianceAttest ||
		!model.AcceptanceCannotCreatePass {
		return Point13ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef) ||
		strings.TrimSpace(model.ExportPackageRef) != strings.TrimSpace(exportPackage.ExportPackageID) ||
		strings.TrimSpace(model.HandoffChecklistRef) != strings.TrimSpace(handoff.HandoffChecklistID) {
		return Point13ValCStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.AcceptedLimitations, " "),
		strings.Join(model.RejectedItems, " "),
		strings.Join(model.CustomerQuestions, " "),
	) {
		return Point13ValCStateBlocked
	}
	if len(model.RejectedItems) > 0 {
		return Point13ValCStateReviewRequired
	}
	return Point13ValCStateActive
}

func EvaluatePoint13ValCSupportOffboardingHandoffState(model Point13ValCSupportOffboardingHandoffPacket, dependency Point13ValCDependencySnapshot, exportPackage Point13ValCCustomerEvidenceExportPackage) string {
	if !point13ValCSupportOffboardingPacketRefValid(model.SupportOffboardingPacketID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValAOwnerRefValid(model.SupportOwnerRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point13ValAOwnerRefValid(model.RetentionOwnerRef) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValBSupportTraceRefValid(model.SupportTraceRef) ||
		!point13ValCOffboardingPlanRefValid(model.OffboardingPlanRef) ||
		!point13Val0OperationalRefValid(model.DisposalPathRef, "disposal_path_") ||
		!point13ValCRetentionClassRefsValid(model.RetentionClassRefs) ||
		len(model.RetentionClassRefs) == 0 ||
		!point13ValAAuditRefsValid(model.AuditEventRefs) ||
		len(model.AuditEventRefs) == 0 ||
		!model.SupportMaterialCandidateOnly ||
		!model.SupportOffboardingCannotMutateCanonical ||
		!model.SupportOffboardingCannotOverrideDecision ||
		!model.SupportOffboardingCannotApproveProduction {
		return Point13ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.SupportOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef) ||
		strings.TrimSpace(model.RetentionOwnerRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.RetentionOwnerRef) ||
		strings.TrimSpace(model.ExportPackageRef) != strings.TrimSpace(exportPackage.ExportPackageID) ||
		strings.TrimSpace(model.SupportTraceRef) != strings.TrimSpace(dependency.ValB.SupportActionTrace.SupportTraceID) ||
		strings.TrimSpace(model.DisposalPathRef) != strings.TrimSpace(dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.DisposalPathRef) ||
		!point12Val0ExactStringSetMatch(model.RetentionClassRefs, []string{
			dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.EvidenceRetentionClassRef,
			dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.SupportArtifactRetentionClassRef,
		}) {
		return Point13ValCStateBlocked
	}
	if model.IndefiniteRetentionRequested && !point12Val0GovernanceEventRefValid(model.RetentionGovernanceEventRef) {
		return Point13ValCStateBlocked
	}
	return Point13ValCStateActive
}

func EvaluatePoint13ValCAIEvidenceExportLineageState(model Point13ValCAIEvidenceExportLineageSummary, dependency Point13ValCDependencySnapshot) string {
	if point11Val0ContainsTrimmed(point12Val0BlockedAIEvidenceCandidateTypes(), model.AIOutputType) {
		return Point13ValCStateBlocked
	}
	aiTrace := dependency.ValB.AIEvidenceOperationTrace
	if !point13ValCAIExportSummaryRefValid(model.AIExportSummaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValBAITraceRefValid(model.AITraceRef) ||
		!point12Val0AIEvidenceCandidateTypeValid(model.AIOutputType) ||
		!point13ValAAIEvidenceCandidateRefValid(model.EvidenceCandidateRef) ||
		!point13ValBEvidenceRefsValid(model.InputEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.InputEvidenceHashRefs) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.InputEvidenceRefs, model.InputEvidenceHashRefs) ||
		!point12Val0VersionRefValid(model.ModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.PermissionManifestHash) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		!model.IncludedInExport ||
		!model.AdvisoryOnly ||
		!model.EvidenceCandidateOnly ||
		!model.AISummaryCannotStrengthenExportClaim ||
		!model.AISummaryCannotSatisfyAcceptanceByItself ||
		model.PassAllowed ||
		model.ApprovalGranted ||
		model.DeploymentAuthorized ||
		model.ProductionReadinessClaimed ||
		model.ProductionMutationAllowed ||
		model.CanonicalMutationAllowed {
		return Point13ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.AITraceRef) != strings.TrimSpace(aiTrace.AITraceID) ||
		strings.TrimSpace(model.EvidenceCandidateRef) != strings.TrimSpace(aiTrace.EvidenceCandidateRef) ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceRefs, aiTrace.InputEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceHashRefs, aiTrace.InputEvidenceHashRefs) ||
		strings.TrimSpace(model.ModelOrRuleVersionRef) != strings.TrimSpace(aiTrace.ModelOrRuleVersionRef) ||
		strings.TrimSpace(model.PermissionManifestHash) != strings.TrimSpace(aiTrace.PermissionManifestHash) ||
		strings.TrimSpace(model.AuditEventRef) != strings.TrimSpace(aiTrace.AuditEventRef) ||
		point13ValAContainsCrossTenantArtifact(model.InputEvidenceRefs) {
		return Point13ValCStateBlocked
	}
	if model.ExternalAPIAllowed {
		if !point12Val0GovernanceEventRefValid(model.ExternalAPIGovernanceEventRef) {
			return Point13ValCStateBlocked
		}
		return Point13ValCStateReviewRequired
	}
	return Point13ValCStateActive
}

func EvaluatePoint13ValCNoOverclaimState(model Point13ValCNoOverclaimExportWording) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point13ValCTextListValid(model.AllowedSafeWording) ||
		!point13ValCTextListValid(model.BlockedWording) {
		return Point13ValCStateBlocked
	}
	if !point12Val0ExactStringSetMatch(model.AllowedSafeWording, point13ValCAllowedSafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point13Val0ForbiddenClaims()) ||
		strings.TrimSpace(model.ProjectionDisclaimer) != point13ValCProjectionDisclaimerBaseline {
		return Point13ValCStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.ObservedCustomerExportTexts, " "),
		strings.Join(model.ObservedHandoffTexts, " "),
		strings.Join(model.ObservedAcceptanceTexts, " "),
		strings.Join(model.ObservedSupportOffboardingTexts, " "),
	) {
		return Point13ValCStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(strings.Join(model.InternalDiagnosticTexts, " ")) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point13ValCStateBlocked
	}
	return Point13ValCStateActive
}

type point13ValCComponentState struct {
	Name  string
	State string
}

func point13ValCComponentStates(model Point13ValCFoundation) []point13ValCComponentState {
	return []point13ValCComponentState{
		{Name: "dependency", State: model.DependencyState},
		{Name: "customer_evidence_export_package", State: model.CustomerEvidenceExportPackageState},
		{Name: "redaction_safe_disclosure", State: model.RedactionSafeDisclosureState},
		{Name: "operational_handoff_checklist", State: model.OperationalHandoffChecklistState},
		{Name: "customer_acceptance_trace", State: model.CustomerAcceptanceTraceState},
		{Name: "support_offboarding_handoff", State: model.SupportOffboardingHandoffState},
		{Name: "ai_evidence_export_lineage", State: model.AIEvidenceExportLineageState},
		{Name: "no_overclaim", State: model.NoOverclaimState},
	}
}

func point13ValCBlockingReasons(model Point13ValCFoundation) []string {
	reasons := []string{}
	for _, component := range point13ValCComponentStates(model) {
		if component.State == Point13ValCStateBlocked || component.State == Point13ValCStateIncomplete {
			reasons = append(reasons, component.Name+":"+component.State)
		}
	}
	return reasons
}

func EvaluatePoint13ValCState(model Point13ValCFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValCStateBlocked
	}
	hasReviewRequired := false
	hasIncomplete := false
	for _, component := range point13ValCComponentStates(model) {
		switch component.State {
		case Point13ValCStateBlocked:
			return Point13ValCStateBlocked
		case Point13ValCStateReviewRequired:
			hasReviewRequired = true
		case Point13ValCStateIncomplete:
			hasIncomplete = true
		}
	}
	if hasReviewRequired {
		return Point13ValCStateReviewRequired
	}
	if hasIncomplete {
		return Point13ValCStateIncomplete
	}
	return Point13ValCStateActive
}

func Point13ValCFoundationModel() Point13ValCFoundation {
	disclaimer := point13ValCProjectionDisclaimerBaseline
	dependency := point13ValCDependencySnapshotModel()
	tenantScope := dependency.InheritedTenantScope

	exportPackage := Point13ValCCustomerEvidenceExportPackage{
		ExportPackageID:                     "export_package_point13_valc_001",
		TenantScope:                         tenantScope,
		OperationLedgerRef:                  dependency.ValB.PilotEvidenceOperationLedger.LedgerID,
		CustomerReviewTraceRef:              dependency.ValB.CustomerReviewTrace.ReviewTraceID,
		SupportTraceRef:                     dependency.ValB.SupportActionTrace.SupportTraceID,
		ExitEvidencePacketRef:               dependency.ValB.PilotExitEvidencePacket.PacketID,
		ExportedEvidenceRefs:                append([]string{}, dependency.ValB.Dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs...),
		ExportedEvidenceHashRefs:            append([]string{}, dependency.ValB.Dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs...),
		RetentionClassRef:                   dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.EvidenceRetentionClassRef,
		ExportOwnerRef:                      dependency.ValB.Dependency.ValA.PilotExecutionContract.PilotOwnerRef,
		CustomerOwnerRef:                    dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
		AuditEventRef:                       "audit_point13_valc_export_001",
		PublicPrivateClassification:         point12ValCClassificationCustomerRedacted,
		ExportIsReadOnly:                    true,
		ExportIsOperationalEvidenceOnly:     true,
		ExportCannotCreatePass:              true,
		ExportCannotApproveProduction:       true,
		ExportCannotCertify:                 true,
		ExportCannotMutateCanonicalEvidence: true,
	}
	exportPackage.ExportManifestHash = point13ValCComputedExportManifestHash(exportPackage)

	redaction := Point13ValCRedactionSafeDisclosureBoundary{
		RedactionBoundaryID:                     "redaction_boundary_point13_valc_001",
		TenantScope:                             tenantScope,
		ExportPackageID:                         exportPackage.ExportPackageID,
		RedactionManifestRef:                    "redaction_manifest_point13_valc_001",
		RedactedFields:                          []string{"customer_contact_email"},
		RedactionReasons:                        []string{"tenant_private_contact_redaction"},
		RedactionApproverRef:                    "redaction_approver_point13_valc_001",
		RedactionAuditEventRef:                  "audit_point13_valc_redaction_001",
		MinimumSafeStatement:                    "customer evidence export package prepared for operational handoff",
		DisallowedCustomerClaims:                point13Val0ForbiddenClaims(),
		SurvivingCustomerClaims:                 []string{"customer evidence export package", "operational handoff checklist"},
		RedactionCannotStrengthenClaim:          true,
		RedactionCannotHideDecisiveMissingProof: true,
	}

	handoff := Point13ValCOperationalHandoffChecklist{
		HandoffChecklistID:               "handoff_checklist_point13_valc_001",
		TenantScope:                      tenantScope,
		HandoffOwnerRef:                  dependency.ValB.Dependency.ValA.PilotExecutionContract.PilotOwnerRef,
		CustomerOwnerRef:                 dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
		SupportOwnerRef:                  dependency.ValB.Dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef,
		ExportPackageRef:                 exportPackage.ExportPackageID,
		ExitPacketRef:                    dependency.ValB.PilotExitEvidencePacket.PacketID,
		ChecklistItems:                   point13ValCDefaultChecklistItems(),
		RequiredAckRefs:                  point13ValCDefaultRequiredAckRefs(),
		AuditEventRefs:                   point13ValCDefaultHandoffAuditRefs(),
		HandoffIsOperationalOnly:         true,
		HandoffCannotApproveProduction:   true,
		HandoffCannotAuthorizeDeployment: true,
		HandoffCannotCreatePass:          true,
	}
	handoff.ChecklistBindingHash = point13ValCComputedChecklistBindingHash(handoff)

	acceptance := Point13ValCCustomerAcceptanceTrace{
		AcceptanceTraceID:                 "acceptance_trace_point13_valc_001",
		TenantScope:                       tenantScope,
		CustomerOwnerRef:                  dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
		ExportPackageRef:                  exportPackage.ExportPackageID,
		HandoffChecklistRef:               handoff.HandoffChecklistID,
		AcceptedLimitations:               []string{"operational evidence only"},
		CustomerQuestions:                 []string{"customer export package review question"},
		ResponseRefs:                      []string{"response_point13_valc_001"},
		AuditEventRefs:                    []string{"audit_point13_valc_acceptance_001"},
		AcceptanceIsNotProductionApproval: true,
		AcceptanceIsNotComplianceAttest:   true,
		AcceptanceCannotCreatePass:        true,
	}

	supportOffboarding := Point13ValCSupportOffboardingHandoffPacket{
		SupportOffboardingPacketID: "support_offboarding_packet_point13_valc_001",
		TenantScope:                tenantScope,
		SupportOwnerRef:            dependency.ValB.Dependency.ValA.SupportResponsibilityMatrix.EscalationOwnerRef,
		CustomerOwnerRef:           dependency.ValB.Dependency.ValA.PilotExecutionContract.CustomerOwnerRef,
		RetentionOwnerRef:          dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.RetentionOwnerRef,
		ExportPackageRef:           exportPackage.ExportPackageID,
		SupportTraceRef:            dependency.ValB.SupportActionTrace.SupportTraceID,
		OffboardingPlanRef:         "offboarding_plan_point13_valc_001",
		DisposalPathRef:            dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.DisposalPathRef,
		RetentionClassRefs: []string{
			dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.EvidenceRetentionClassRef,
			dependency.ValB.Dependency.ValA.Dependency.Val0.OffboardingRetentionBoundary.SupportArtifactRetentionClassRef,
		},
		AuditEventRefs:                            []string{"audit_point13_valc_offboarding_001", "audit_point13_valc_offboarding_002"},
		SupportMaterialCandidateOnly:              true,
		SupportOffboardingCannotMutateCanonical:   true,
		SupportOffboardingCannotOverrideDecision:  true,
		SupportOffboardingCannotApproveProduction: true,
	}

	aiSummary := Point13ValCAIEvidenceExportLineageSummary{
		AIExportSummaryID:                        "ai_export_summary_point13_valc_001",
		TenantScope:                              tenantScope,
		AITraceRef:                               dependency.ValB.AIEvidenceOperationTrace.AITraceID,
		AIOutputType:                             dependency.ValB.AIEvidenceOperationTrace.AIOutputType,
		EvidenceCandidateRef:                     dependency.ValB.AIEvidenceOperationTrace.EvidenceCandidateRef,
		InputEvidenceRefs:                        append([]string{}, dependency.ValB.AIEvidenceOperationTrace.InputEvidenceRefs...),
		InputEvidenceHashRefs:                    append([]string{}, dependency.ValB.AIEvidenceOperationTrace.InputEvidenceHashRefs...),
		ModelOrRuleVersionRef:                    dependency.ValB.AIEvidenceOperationTrace.ModelOrRuleVersionRef,
		PermissionManifestHash:                   dependency.ValB.AIEvidenceOperationTrace.PermissionManifestHash,
		AuditEventRef:                            dependency.ValB.AIEvidenceOperationTrace.AuditEventRef,
		IncludedInExport:                         true,
		AdvisoryOnly:                             true,
		EvidenceCandidateOnly:                    true,
		AISummaryCannotStrengthenExportClaim:     true,
		AISummaryCannotSatisfyAcceptanceByItself: true,
	}

	return Point13ValCFoundation{
		CurrentState:                       Point13ValCStateActive,
		ProjectionDisclaimer:               disclaimer,
		DependencyState:                    Point13ValCStateActive,
		CustomerEvidenceExportPackageState: Point13ValCStateActive,
		RedactionSafeDisclosureState:       Point13ValCStateActive,
		OperationalHandoffChecklistState:   Point13ValCStateActive,
		CustomerAcceptanceTraceState:       Point13ValCStateActive,
		SupportOffboardingHandoffState:     Point13ValCStateActive,
		AIEvidenceExportLineageState:       Point13ValCStateActive,
		NoOverclaimState:                   Point13ValCStateActive,
		Dependency:                         dependency,
		CustomerEvidenceExportPackage:      exportPackage,
		RedactionSafeDisclosureBoundary:    redaction,
		OperationalHandoffChecklist:        handoff,
		CustomerAcceptanceTrace:            acceptance,
		SupportOffboardingHandoffPacket:    supportOffboarding,
		AIEvidenceExportLineageSummary:     aiSummary,
		NoOverclaimExportWording: Point13ValCNoOverclaimExportWording{
			ObservedCustomerExportTexts:          []string{"customer evidence export package"},
			ObservedHandoffTexts:                 []string{"operational handoff checklist"},
			ObservedAcceptanceTexts:              []string{"customer acceptance trace"},
			ObservedSupportOffboardingTexts:      []string{"support offboarding handoff"},
			InternalDiagnosticTexts:              []string{"blocked wording remains denylisted internally"},
			InternalDiagnosticsClassifiedBlocked: true,
			AllowedSafeWording:                   point13ValCAllowedSafeWording(),
			BlockedWording:                       point13Val0ForbiddenClaims(),
			ProjectionDisclaimer:                 disclaimer,
		},
	}
}

func ComputePoint13ValCFoundation(model Point13ValCFoundation) Point13ValCFoundation {
	dependencyState, dependencyReasons := point13ValCDependencyStateAndReasons(model.Dependency)
	model.DependencyState = dependencyState
	model.CustomerEvidenceExportPackageState = EvaluatePoint13ValCCustomerEvidenceExportPackageState(model.CustomerEvidenceExportPackage, model.Dependency)
	model.RedactionSafeDisclosureState = EvaluatePoint13ValCRedactionSafeDisclosureState(model.RedactionSafeDisclosureBoundary, model.Dependency, model.CustomerEvidenceExportPackageState, model.CustomerEvidenceExportPackage)
	model.OperationalHandoffChecklistState = EvaluatePoint13ValCOperationalHandoffChecklistState(model.OperationalHandoffChecklist, model.Dependency, model.CustomerEvidenceExportPackage)
	model.CustomerAcceptanceTraceState = EvaluatePoint13ValCCustomerAcceptanceTraceState(model.CustomerAcceptanceTrace, model.Dependency, model.CustomerEvidenceExportPackage, model.OperationalHandoffChecklist)
	model.SupportOffboardingHandoffState = EvaluatePoint13ValCSupportOffboardingHandoffState(model.SupportOffboardingHandoffPacket, model.Dependency, model.CustomerEvidenceExportPackage)
	model.AIEvidenceExportLineageState = EvaluatePoint13ValCAIEvidenceExportLineageState(model.AIEvidenceExportLineageSummary, model.Dependency)
	model.NoOverclaimState = EvaluatePoint13ValCNoOverclaimState(model.NoOverclaimExportWording)
	model.CurrentState = EvaluatePoint13ValCState(model)
	model.BlockingReasons = point13ValCBlockingReasons(model)
	model.ReviewPrerequisites = nil
	if model.DependencyState == Point13ValCStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, dependencyReasons...)
	}
	if model.RedactionSafeDisclosureState == Point13ValCStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "redaction_safe_disclosure_requires_review")
	}
	if model.CustomerAcceptanceTraceState == Point13ValCStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "customer_acceptance_trace_requires_review")
	}
	if model.AIEvidenceExportLineageState == Point13ValCStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "ai_evidence_export_lineage_requires_governance_review")
	}
	return model
}
