package formal

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	Point13ValEStateActive         = "point13_vale_final_operationalization_closure_active"
	Point13ValEStateBlocked        = "point13_vale_final_operationalization_closure_blocked"
	Point13ValEStateReviewRequired = "point13_vale_final_operationalization_closure_review_required"
	Point13ValEStateIncomplete     = "point13_vale_final_operationalization_closure_incomplete"
	Point13ValEStatePassConfirmed  = "point13_vale_final_operationalization_closure_pass_confirmed"
)

const (
	point13ValEWaveID                       = "val_e"
	point13ValEPreviousWaveID               = point13ValDWaveID
	point13ValEScope                        = "final_operationalization_closure_gate"
	point13ValEProjectionDisclaimerBaseline = "projection_only not_canonical_truth point13_vale_final_operationalization_closure_gate"
	point13ValEPoint13PassToken             = "point_13_pass"
)

type Point13ValEDependencySnapshot struct {
	ValDCurrentState                              string                `json:"vald_current_state"`
	ValDDependencyState                           string                `json:"vald_dependency_state"`
	ValDCustomerAuditorOperationalTimelineState   string                `json:"vald_customer_auditor_operational_timeline_state"`
	ValDHandoffTraceQueryProjectionState          string                `json:"vald_handoff_trace_query_projection_state"`
	ValDExportPackageReadProjectionState          string                `json:"vald_export_package_read_projection_state"`
	ValDCustomerAuditorExplanationProjectionState string                `json:"vald_customer_auditor_explanation_projection_state"`
	ValDTimelineAccessBoundaryState               string                `json:"vald_timeline_access_boundary_state"`
	ValDAITimelineLineageProjectionState          string                `json:"vald_ai_timeline_lineage_projection_state"`
	ValDNoOverclaimState                          string                `json:"vald_no_overclaim_state"`
	ValDPointID                                   string                `json:"vald_point_id"`
	ValDWaveID                                    string                `json:"vald_wave_id"`
	ValDDependencyComputedFromUpstream            bool                  `json:"vald_dependency_computed_from_upstream"`
	ValDPoint13PassSeen                           bool                  `json:"vald_point13_pass_seen"`
	InheritedValCCurrentState                     string                `json:"inherited_valc_current_state"`
	InheritedValBCurrentState                     string                `json:"inherited_valb_current_state"`
	InheritedValACurrentState                     string                `json:"inherited_vala_current_state"`
	InheritedVal0CurrentState                     string                `json:"inherited_val0_current_state"`
	InheritedPoint12CurrentState                  string                `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState               string                `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState              string                `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult                string                `json:"inherited_point12_reviewer_result"`
	InheritedTenantScope                          string                `json:"inherited_tenant_scope"`
	InheritedAIModelOrRuleVersionRef              string                `json:"inherited_ai_model_or_rule_version_ref"`
	InheritedAIPermissionManifestHash             string                `json:"inherited_ai_permission_manifest_hash"`
	SnapshotFromComputedOutput                    bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                           []string              `json:"review_prerequisites,omitempty"`
	ValD                                          Point13ValDFoundation `json:"vald"`
}

type Point13ValEClosureEvaluator struct {
	ClosureEvaluatorID       string   `json:"closure_evaluator_id"`
	TenantScope              string   `json:"tenant_scope"`
	DependencyGateResult     string   `json:"dependency_gate_result"`
	NoOverclaimResult        string   `json:"no_overclaim_result"`
	AuthorityBoundaryResult  string   `json:"authority_boundary_result"`
	TimestampIntegrityResult string   `json:"timestamp_integrity_result"`
	TenantIsolationResult    string   `json:"tenant_isolation_result"`
	EvidenceIntegrityResult  string   `json:"evidence_integrity_result"`
	ReadOnlyProjectionOnly   bool     `json:"read_only_projection_only"`
	NoMutationPathsDetected  bool     `json:"no_mutation_paths_detected"`
	NoPrematurePoint13Pass   bool     `json:"no_premature_point13_pass"`
	CommandsRun              []string `json:"commands_run,omitempty"`
	TestsRun                 []string `json:"tests_run,omitempty"`
	GrepsRun                 []string `json:"greps_run,omitempty"`
	NegativeFixturesRun      []string `json:"negative_fixtures_run,omitempty"`
	ReviewerResult           string   `json:"reviewer_result"`
	CurrentState             string   `json:"current_state"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type Point13ValENoOverclaimFinalCheck struct {
	ObservedExportTexts                  []string `json:"observed_export_texts,omitempty"`
	ObservedTimelineTexts                []string `json:"observed_timeline_texts,omitempty"`
	ObservedQueryTexts                   []string `json:"observed_query_texts,omitempty"`
	ObservedExplanationTexts             []string `json:"observed_explanation_texts,omitempty"`
	ObservedAcceptanceTexts              []string `json:"observed_acceptance_texts,omitempty"`
	ObservedSupportOffboardingTexts      []string `json:"observed_support_offboarding_texts,omitempty"`
	ObservedAITexts                      []string `json:"observed_ai_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point13ValEAuthorityBoundaryCheck struct {
	CheckID                                  string `json:"check_id"`
	ValCExportReadOnly                       bool   `json:"valc_export_read_only"`
	ValCExportOperationalOnly                bool   `json:"valc_export_operational_only"`
	ValCExportCannotCreatePass               bool   `json:"valc_export_cannot_create_pass"`
	ValCExportCannotApproveProduction        bool   `json:"valc_export_cannot_approve_production"`
	ValCExportCannotCertify                  bool   `json:"valc_export_cannot_certify"`
	ValCExportCannotMutateCanonicalEvidence  bool   `json:"valc_export_cannot_mutate_canonical_evidence"`
	ValCHandoffOperationalOnly               bool   `json:"valc_handoff_operational_only"`
	ValCHandoffCannotApproveProduction       bool   `json:"valc_handoff_cannot_approve_production"`
	ValCHandoffCannotAuthorizeDeployment     bool   `json:"valc_handoff_cannot_authorize_deployment"`
	ValCHandoffCannotCreatePass              bool   `json:"valc_handoff_cannot_create_pass"`
	ValCAcceptanceNotProductionApproval      bool   `json:"valc_acceptance_not_production_approval"`
	ValCAcceptanceNotComplianceAttestation   bool   `json:"valc_acceptance_not_compliance_attestation"`
	ValCAcceptanceCannotCreatePass           bool   `json:"valc_acceptance_cannot_create_pass"`
	ValCSupportCandidateOnly                 bool   `json:"valc_support_candidate_only"`
	ValCSupportCannotMutateCanonicalEvidence bool   `json:"valc_support_cannot_mutate_canonical_evidence"`
	ValCSupportCannotOverrideCoreDecision    bool   `json:"valc_support_cannot_override_core_decision"`
	ValCSupportCannotApproveProduction       bool   `json:"valc_support_cannot_approve_production"`
	ValDTimelineReadOnly                     bool   `json:"vald_timeline_read_only"`
	ValDTimelineCannotMutateState            bool   `json:"vald_timeline_cannot_mutate_state"`
	ValDQueryReadOnly                        bool   `json:"vald_query_read_only"`
	ValDQueryMutationRequested               bool   `json:"vald_query_mutation_requested"`
	ValDQueryWriteRequested                  bool   `json:"vald_query_write_requested"`
	ValDExportReadOnly                       bool   `json:"vald_export_read_only"`
	ValDExportCannotOverwriteHashes          bool   `json:"vald_export_cannot_overwrite_hashes"`
	ValDExplanationAdvisoryOnly              bool   `json:"vald_explanation_advisory_only"`
	ValDExplanationCannotStrengthenClaims    bool   `json:"vald_explanation_cannot_strengthen_claims"`
	ValDExplanationCannotApproveProduction   bool   `json:"vald_explanation_cannot_approve_production"`
	ValDExplanationCannotCreatePass          bool   `json:"vald_explanation_cannot_create_pass"`
	ValDAccessReadOnly                       bool   `json:"vald_access_read_only"`
	ValDAccessMutationRequested              bool   `json:"vald_access_mutation_requested"`
	ValDAIReadOnly                           bool   `json:"vald_ai_read_only"`
	ValDAIAdvisoryOnly                       bool   `json:"vald_ai_advisory_only"`
	ValDAIEvidenceCandidateOnly              bool   `json:"vald_ai_evidence_candidate_only"`
	ValDAIPassAllowed                        bool   `json:"vald_ai_pass_allowed"`
	ValDAIApprovalGranted                    bool   `json:"vald_ai_approval_granted"`
	ValDAIDeploymentAuthorized               bool   `json:"vald_ai_deployment_authorized"`
	ValDAIProductionReadinessClaimed         bool   `json:"vald_ai_production_readiness_claimed"`
	ValDAIProductionMutationAllowed          bool   `json:"vald_ai_production_mutation_allowed"`
	ValDAICanonicalMutationAllowed           bool   `json:"vald_ai_canonical_mutation_allowed"`
	ValDAIExternalAPIAllowed                 bool   `json:"vald_ai_external_api_allowed"`
	CurrentState                             string `json:"current_state"`
	ProjectionDisclaimer                     string `json:"projection_disclaimer"`
}

type Point13ValETimestampIntegrityCheck struct {
	CheckID                string   `json:"check_id"`
	TenantScope            string   `json:"tenant_scope"`
	TimelineRef            string   `json:"timeline_ref"`
	EventRefs              []string `json:"event_refs,omitempty"`
	AuditEventRefs         []string `json:"audit_event_refs,omitempty"`
	CanonicalOccurredAts   []string `json:"canonical_occurred_ats,omitempty"`
	TimeSources            []string `json:"time_sources,omitempty"`
	SourceMetadataRefs     []string `json:"source_metadata_refs,omitempty"`
	ClientReportedAts      []string `json:"client_reported_ats,omitempty"`
	VerifiedAt             string   `json:"verified_at"`
	ClientTimeMetadataOnly bool     `json:"client_time_metadata_only"`
	CurrentState           string   `json:"current_state"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
}

type Point13ValETwitterIsolationCheck struct {
	CheckID                 string   `json:"check_id"`
	TenantScope             string   `json:"tenant_scope"`
	ExportPackageRef        string   `json:"export_package_ref"`
	TimelineRef             string   `json:"timeline_ref"`
	QueryProjectionRef      string   `json:"query_projection_ref"`
	ExportReadProjectionRef string   `json:"export_read_projection_ref"`
	AccessBoundaryRef       string   `json:"access_boundary_ref"`
	AudienceType            string   `json:"audience_type"`
	AudienceRef             string   `json:"audience_ref"`
	CustomerOwnerRef        string   `json:"customer_owner_ref"`
	AuditorOwnerRef         string   `json:"auditor_owner_ref"`
	ExportedEvidenceRefs    []string `json:"exported_evidence_refs,omitempty"`
	TimelineSourceRefs      []string `json:"timeline_source_refs,omitempty"`
	QueryFilterRefs         []string `json:"query_filter_refs,omitempty"`
	CurrentState            string   `json:"current_state"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
}

type Point13ValEEvidenceIntegrityCheck struct {
	CheckID                  string   `json:"check_id"`
	TenantScope              string   `json:"tenant_scope"`
	ExportPackageRef         string   `json:"export_package_ref"`
	ReadProjectionRef        string   `json:"read_projection_ref"`
	AIProjectionRef          string   `json:"ai_projection_ref"`
	ExportedEvidenceRefs     []string `json:"exported_evidence_refs,omitempty"`
	ExportedEvidenceHashes   []string `json:"exported_evidence_hashes,omitempty"`
	ExportManifestHash       string   `json:"export_manifest_hash"`
	AIEvidenceCandidateRef   string   `json:"ai_evidence_candidate_ref"`
	AIInputEvidenceRefs      []string `json:"ai_input_evidence_refs,omitempty"`
	AIInputEvidenceHashes    []string `json:"ai_input_evidence_hashes,omitempty"`
	LineageComplete          bool     `json:"lineage_complete"`
	NoRecomputeDriftDetected bool     `json:"no_recompute_drift_detected"`
	CurrentState             string   `json:"current_state"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type Point13PassClosureManifest struct {
	CurrentState                    string   `json:"current_state"`
	ClosureManifestID               string   `json:"closure_manifest_id"`
	PointID                         string   `json:"point_id"`
	WaveID                          string   `json:"wave_id"`
	Scope                           string   `json:"scope"`
	DependencyGateResult            string   `json:"dependency_gate_result"`
	ClosureEvaluatorResult          string   `json:"closure_evaluator_result"`
	NoOverclaimResult               string   `json:"no_overclaim_result"`
	AuthorityBoundaryResult         string   `json:"authority_boundary_result"`
	TimestampIntegrityResult        string   `json:"timestamp_integrity_result"`
	TenantIsolationResult           string   `json:"tenant_isolation_result"`
	EvidenceIntegrityResult         string   `json:"evidence_integrity_result"`
	ValDCurrentState                string   `json:"vald_current_state"`
	ValDTimelineState               string   `json:"vald_timeline_state"`
	ValDQueryProjectionState        string   `json:"vald_query_projection_state"`
	ValDExportReadProjectionState   string   `json:"vald_export_read_projection_state"`
	ValDExplanationProjectionState  string   `json:"vald_explanation_projection_state"`
	ValDAccessBoundaryState         string   `json:"vald_access_boundary_state"`
	ValDAIProjectionState           string   `json:"vald_ai_projection_state"`
	ValDNoOverclaimState            string   `json:"vald_no_overclaim_state"`
	ValCExportPackageRef            string   `json:"valc_export_package_ref"`
	ValCHandoffChecklistRef         string   `json:"valc_handoff_checklist_ref"`
	ValCAcceptanceTraceRef          string   `json:"valc_acceptance_trace_ref"`
	ValCSupportOffboardingPacketRef string   `json:"valc_support_offboarding_packet_ref"`
	ValCAIExportSummaryRef          string   `json:"valc_ai_export_summary_ref"`
	ValDTimelineRef                 string   `json:"vald_timeline_ref"`
	ValDQueryProjectionRef          string   `json:"vald_query_projection_ref"`
	ValDExportReadProjectionRef     string   `json:"vald_export_read_projection_ref"`
	ValDExplanationProjectionRef    string   `json:"vald_explanation_projection_ref"`
	ValDAccessBoundaryRef           string   `json:"vald_access_boundary_ref"`
	ValDAIProjectionRef             string   `json:"vald_ai_projection_ref"`
	TenantScope                     string   `json:"tenant_scope"`
	ExportedEvidenceRefs            []string `json:"exported_evidence_refs,omitempty"`
	ExportedEvidenceHashes          []string `json:"exported_evidence_hashes,omitempty"`
	ExportManifestHash              string   `json:"export_manifest_hash"`
	RetentionClassRef               string   `json:"retention_class_ref"`
	AIModelOrRuleVersionRef         string   `json:"ai_model_or_rule_version_ref"`
	AIPermissionManifestHash        string   `json:"ai_permission_manifest_hash"`
	CommandsRun                     []string `json:"commands_run,omitempty"`
	TestsRun                        []string `json:"tests_run,omitempty"`
	GrepsRun                        []string `json:"greps_run,omitempty"`
	NegativeFixturesRun             []string `json:"negative_fixtures_run,omitempty"`
	ReviewerResult                  string   `json:"reviewer_result"`
	GeneratedAt                     string   `json:"generated_at"`
	Point13PassAllowed              bool     `json:"point13_pass_allowed"`
	Point13PassToken                string   `json:"point13_pass_token"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
}

type Point13ValEFoundation struct {
	CurrentState                 string                             `json:"current_state"`
	BlockingReasons              []string                           `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites          []string                           `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer         string                             `json:"projection_disclaimer"`
	DependencyState              string                             `json:"dependency_state"`
	ClosureEvaluatorState        string                             `json:"closure_evaluator_state"`
	NoOverclaimFinalCheckState   string                             `json:"no_overclaim_final_check_state"`
	AuthorityBoundaryCheckState  string                             `json:"authority_boundary_check_state"`
	TimestampIntegrityCheckState string                             `json:"timestamp_integrity_check_state"`
	TenantIsolationCheckState    string                             `json:"tenant_isolation_check_state"`
	EvidenceIntegrityCheckState  string                             `json:"evidence_integrity_check_state"`
	PassClosureManifestState     string                             `json:"pass_closure_manifest_state"`
	Point13PassAllowed           bool                               `json:"point13_pass_allowed"`
	Point13PassToken             string                             `json:"point13_pass_token"`
	Dependency                   Point13ValEDependencySnapshot      `json:"dependency"`
	ClosureEvaluator             Point13ValEClosureEvaluator        `json:"closure_evaluator"`
	NoOverclaimFinalCheck        Point13ValENoOverclaimFinalCheck   `json:"no_overclaim_final_check"`
	AuthorityBoundaryCheck       Point13ValEAuthorityBoundaryCheck  `json:"authority_boundary_check"`
	TimestampIntegrityCheck      Point13ValETimestampIntegrityCheck `json:"timestamp_integrity_check"`
	TenantIsolationCheck         Point13ValETwitterIsolationCheck   `json:"tenant_isolation_check"`
	EvidenceIntegrityCheck       Point13ValEEvidenceIntegrityCheck  `json:"evidence_integrity_check"`
	PassClosureManifest          Point13PassClosureManifest         `json:"pass_closure_manifest"`
}

func point13ValEStates() []string {
	return []string{
		Point13ValEStateActive,
		Point13ValEStateBlocked,
		Point13ValEStateReviewRequired,
		Point13ValEStateIncomplete,
		Point13ValEStatePassConfirmed,
	}
}

func point13ValEStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point13ValEStates(), value)
}

func point13ValEClosureEvaluatorRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "closure_evaluator_point13_vale_")
}

func point13ValENoOverclaimCheckRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "no_overclaim_check_point13_vale_")
}

func point13ValEAuthorityBoundaryCheckRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "authority_boundary_check_point13_vale_")
}

func point13ValETimestampIntegrityCheckRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "timestamp_integrity_check_point13_vale_")
}

func point13ValETwitterIsolationCheckRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "tenant_isolation_check_point13_vale_")
}

func point13ValEEvidenceIntegrityCheckRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "evidence_integrity_check_point13_vale_")
}

func point13ValEClosureManifestRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "closure_manifest_point13_vale_")
}

func point13ValECommandRunRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "command_run_point13_vale_")
}

func point13ValETestRunRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "test_run_point13_vale_")
}

func point13ValEGrepRunRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "grep_run_point13_vale_")
}

func point13ValENegativeFixtureRunRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "negative_fixture_point13_vale_")
}

func point13ValEAllowedSafeWording() []string {
	return []string{
		"final operationalization closure gate",
		"pass closure manifest",
		"customer evidence export package",
		"operational handoff checklist",
		"customer acceptance trace",
		"support offboarding handoff",
		"customer auditor operational timeline",
		"handoff trace query projection",
		"export package read projection",
		"customer auditor explanation projection",
		"advisory ai evidence candidate",
		"evidence support for customer/auditor review",
	}
}

func point13ValEValDPayloadContainsPointPass(valD Point13ValDFoundation) bool {
	payload, err := json.Marshal(valD)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point13ValEPoint13PassToken)
}

func point13ValETimelineEventRefs(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.EventID)
	}
	return values
}

func point13ValETimelineAuditRefs(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.AuditEventRef)
	}
	return values
}

func point13ValETimelineOccurredAts(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.CanonicalOccurredAt)
	}
	return values
}

func point13ValETimelineTimeSources(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.TimeSource)
	}
	return values
}

func point13ValETimelineSourceMetadataRefs(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.SourceMetadataRef)
	}
	return values
}

func point13ValETimelineClientReportedAts(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.ClientReportedAt)
	}
	return values
}

func point13ValETimelineSourceRefs(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.SourceRef)
	}
	return values
}

func point13ValETwitterIsolationRefListValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	for _, value := range values {
		if !point11Val0IdentityValueValid(strings.TrimSpace(value)) {
			return false
		}
	}
	return true
}

func point13ValETwitterValueListValid(values []string, validator func(string) bool) bool {
	if len(values) == 0 {
		return false
	}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return false
		}
		if validator != nil && !validator(trimmed) {
			return false
		}
	}
	return true
}

func point13ValETwitterUniqueTexts(values ...[]string) []string {
	result := []string{}
	seen := map[string]struct{}{}
	for _, group := range values {
		for _, value := range group {
			trimmed := strings.TrimSpace(value)
			if trimmed == "" {
				continue
			}
			key := strings.ToLower(trimmed)
			if _, exists := seen[key]; exists {
				continue
			}
			seen[key] = struct{}{}
			result = append(result, trimmed)
		}
	}
	return result
}

func point13ValETwitterClientMetadataOnly(values []string) bool {
	for _, value := range values {
		if value == "" {
			continue
		}
		if !point11Val0ValidTimestamp(value) {
			return false
		}
	}
	return true
}

func point13ValETwitterCompare(left, right string) (int, bool) {
	leftTime, err := time.Parse(time.RFC3339, strings.TrimSpace(left))
	if err != nil {
		return 0, false
	}
	rightTime, err := time.Parse(time.RFC3339, strings.TrimSpace(right))
	if err != nil {
		return 0, false
	}
	switch {
	case leftTime.Before(rightTime):
		return -1, true
	case leftTime.After(rightTime):
		return 1, true
	default:
		return 0, true
	}
}

func point13ValETwitterFutureDated(entries []Point13ValDTimelineEvent, verifiedAt string) bool {
	for _, entry := range entries {
		comparison, ok := point13ValETwitterCompare(entry.CanonicalOccurredAt, verifiedAt)
		if !ok || comparison > 0 {
			return true
		}
	}
	return false
}

func point13ValEAggregateDerivedStates(states ...string) string {
	hasReviewRequired := false
	hasIncomplete := false
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point13ValEStateBlocked:
			return Point13ValEStateBlocked
		case Point13ValEStateReviewRequired:
			hasReviewRequired = true
		case Point13ValEStateIncomplete:
			hasIncomplete = true
		case Point13ValEStateActive:
		default:
			return Point13ValEStateBlocked
		}
	}
	if hasReviewRequired {
		return Point13ValEStateReviewRequired
	}
	if hasIncomplete {
		return Point13ValEStateIncomplete
	}
	return Point13ValEStateActive
}

func point13ValETwitterIsolationCrossTenant(values ...[]string) bool {
	for _, set := range values {
		if point13ValAContainsCrossTenantArtifact(set) {
			return true
		}
	}
	return false
}

func point13ValEDependencySnapshotFromUpstream(valD Point13ValDFoundation) Point13ValEDependencySnapshot {
	return Point13ValEDependencySnapshot{
		ValDCurrentState:                              valD.CurrentState,
		ValDDependencyState:                           valD.DependencyState,
		ValDCustomerAuditorOperationalTimelineState:   valD.CustomerAuditorOperationalTimelineState,
		ValDHandoffTraceQueryProjectionState:          valD.HandoffTraceQueryProjectionState,
		ValDExportPackageReadProjectionState:          valD.ExportPackageReadProjectionState,
		ValDCustomerAuditorExplanationProjectionState: valD.CustomerAuditorExplanationProjectionState,
		ValDTimelineAccessBoundaryState:               valD.TimelineAccessBoundaryState,
		ValDAITimelineLineageProjectionState:          valD.AITimelineLineageProjectionState,
		ValDNoOverclaimState:                          valD.NoOverclaimState,
		ValDPointID:                                   point13Val0PointID,
		ValDWaveID:                                    point13ValDWaveID,
		ValDDependencyComputedFromUpstream:            valD.Dependency.SnapshotFromComputedOutput,
		ValDPoint13PassSeen:                           point13ValEValDPayloadContainsPointPass(valD),
		InheritedValCCurrentState:                     valD.Dependency.ValCCurrentState,
		InheritedValBCurrentState:                     valD.Dependency.InheritedValBCurrentState,
		InheritedValACurrentState:                     valD.Dependency.InheritedValACurrentState,
		InheritedVal0CurrentState:                     valD.Dependency.InheritedVal0CurrentState,
		InheritedPoint12CurrentState:                  valD.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:               valD.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:              valD.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint12ReviewerResult:                valD.Dependency.InheritedPoint12ReviewerResult,
		InheritedTenantScope:                          valD.Dependency.InheritedTenantScope,
		InheritedAIModelOrRuleVersionRef:              valD.Dependency.InheritedAIModelOrRuleVersionRef,
		InheritedAIPermissionManifestHash:             valD.Dependency.InheritedAIPermissionManifestHash,
		SnapshotFromComputedOutput:                    true,
		ReviewPrerequisites:                           append([]string{}, valD.ReviewPrerequisites...),
		ValD:                                          valD,
	}
}

func point13ValEDependencySnapshotModel() Point13ValEDependencySnapshot {
	return point13ValEDependencySnapshotFromUpstream(ComputePoint13ValDFoundation(Point13ValDFoundationModel()))
}

func point13ValEDependencyStateAndReasons(model Point13ValEDependencySnapshot) (string, []string) {
	reasons := []string{}
	if !model.SnapshotFromComputedOutput || !model.ValDDependencyComputedFromUpstream {
		reasons = append(reasons, "vald_dependency_not_computed_from_upstream")
	}
	if !point13ValDStateValid(model.ValDCurrentState) ||
		!point13ValDStateValid(model.ValDDependencyState) ||
		!point13ValDStateValid(model.ValDCustomerAuditorOperationalTimelineState) ||
		!point13ValDStateValid(model.ValDHandoffTraceQueryProjectionState) ||
		!point13ValDStateValid(model.ValDExportPackageReadProjectionState) ||
		!point13ValDStateValid(model.ValDCustomerAuditorExplanationProjectionState) ||
		!point13ValDStateValid(model.ValDTimelineAccessBoundaryState) ||
		!point13ValDStateValid(model.ValDAITimelineLineageProjectionState) ||
		!point13ValDStateValid(model.ValDNoOverclaimState) ||
		strings.TrimSpace(model.ValDPointID) != point13Val0PointID ||
		strings.TrimSpace(model.ValDWaveID) != point13ValDWaveID ||
		!point13ValCStateValid(model.InheritedValCCurrentState) ||
		!point13ValBStateValid(model.InheritedValBCurrentState) ||
		!point13ValAStateValid(model.InheritedValACurrentState) ||
		!point13Val0StateValid(model.InheritedVal0CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) ||
		!point12Val0VersionRefValid(model.InheritedAIModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.InheritedAIPermissionManifestHash) {
		reasons = append(reasons, "dependency_snapshot_identity_invalid")
	}
	if model.ValDPoint13PassSeen {
		reasons = append(reasons, "vald_point13_pass_seen_before_vale")
	}
	if strings.TrimSpace(model.ValDCurrentState) != strings.TrimSpace(model.ValD.CurrentState) ||
		strings.TrimSpace(model.ValDDependencyState) != strings.TrimSpace(model.ValD.DependencyState) ||
		strings.TrimSpace(model.ValDCustomerAuditorOperationalTimelineState) != strings.TrimSpace(model.ValD.CustomerAuditorOperationalTimelineState) ||
		strings.TrimSpace(model.ValDHandoffTraceQueryProjectionState) != strings.TrimSpace(model.ValD.HandoffTraceQueryProjectionState) ||
		strings.TrimSpace(model.ValDExportPackageReadProjectionState) != strings.TrimSpace(model.ValD.ExportPackageReadProjectionState) ||
		strings.TrimSpace(model.ValDCustomerAuditorExplanationProjectionState) != strings.TrimSpace(model.ValD.CustomerAuditorExplanationProjectionState) ||
		strings.TrimSpace(model.ValDTimelineAccessBoundaryState) != strings.TrimSpace(model.ValD.TimelineAccessBoundaryState) ||
		strings.TrimSpace(model.ValDAITimelineLineageProjectionState) != strings.TrimSpace(model.ValD.AITimelineLineageProjectionState) ||
		strings.TrimSpace(model.ValDNoOverclaimState) != strings.TrimSpace(model.ValD.NoOverclaimState) ||
		model.ValDDependencyComputedFromUpstream != model.ValD.Dependency.SnapshotFromComputedOutput ||
		strings.TrimSpace(model.InheritedValCCurrentState) != strings.TrimSpace(model.ValD.Dependency.ValCCurrentState) ||
		strings.TrimSpace(model.InheritedValBCurrentState) != strings.TrimSpace(model.ValD.Dependency.InheritedValBCurrentState) ||
		strings.TrimSpace(model.InheritedValACurrentState) != strings.TrimSpace(model.ValD.Dependency.InheritedValACurrentState) ||
		strings.TrimSpace(model.InheritedVal0CurrentState) != strings.TrimSpace(model.ValD.Dependency.InheritedVal0CurrentState) ||
		strings.TrimSpace(model.InheritedPoint12CurrentState) != strings.TrimSpace(model.ValD.Dependency.InheritedPoint12CurrentState) ||
		strings.TrimSpace(model.InheritedPoint12DependencyState) != strings.TrimSpace(model.ValD.Dependency.InheritedPoint12DependencyState) ||
		strings.TrimSpace(model.InheritedPoint12PassClosureState) != strings.TrimSpace(model.ValD.Dependency.InheritedPoint12PassClosureState) ||
		strings.TrimSpace(model.InheritedPoint12ReviewerResult) != strings.TrimSpace(model.ValD.Dependency.InheritedPoint12ReviewerResult) ||
		strings.TrimSpace(model.InheritedTenantScope) != strings.TrimSpace(model.ValD.Dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.InheritedAIModelOrRuleVersionRef) != strings.TrimSpace(model.ValD.Dependency.InheritedAIModelOrRuleVersionRef) ||
		strings.TrimSpace(model.InheritedAIPermissionManifestHash) != strings.TrimSpace(model.ValD.Dependency.InheritedAIPermissionManifestHash) {
		reasons = append(reasons, "dependency_snapshot_binding_mismatch")
	}
	for _, state := range []string{
		model.ValDCurrentState,
		model.ValDDependencyState,
		model.ValDCustomerAuditorOperationalTimelineState,
		model.ValDHandoffTraceQueryProjectionState,
		model.ValDExportPackageReadProjectionState,
		model.ValDCustomerAuditorExplanationProjectionState,
		model.ValDTimelineAccessBoundaryState,
		model.ValDAITimelineLineageProjectionState,
		model.ValDNoOverclaimState,
		model.InheritedValCCurrentState,
		model.InheritedValBCurrentState,
		model.InheritedValACurrentState,
		model.InheritedVal0CurrentState,
	} {
		if strings.TrimSpace(state) != Point13ValDStateActive &&
			strings.TrimSpace(state) != Point13ValCStateActive &&
			strings.TrimSpace(state) != Point13ValBStateActive &&
			strings.TrimSpace(state) != Point13ValAStateActive &&
			strings.TrimSpace(state) != Point13Val0StateActive {
			reasons = append(reasons, "upstream_point13_state_not_active")
			break
		}
	}
	if strings.TrimSpace(model.InheritedPoint12CurrentState) != Point12ValEStatePassConfirmed ||
		strings.TrimSpace(model.InheritedPoint12DependencyState) != Point12ValEStateActive ||
		strings.TrimSpace(model.InheritedPoint12PassClosureState) != Point12ValEStateActive ||
		strings.TrimSpace(model.InheritedPoint12ReviewerResult) != point12ValEReviewerResultPassConfirmed {
		reasons = append(reasons, "inherited_point12_closure_not_confirmed")
	}
	if len(reasons) > 0 {
		return Point13ValEStateBlocked, reasons
	}
	return Point13ValEStateActive, nil
}

func point13ValENoOverclaimFinalCheckModel(dependency Point13ValEDependencySnapshot) Point13ValENoOverclaimFinalCheck {
	return Point13ValENoOverclaimFinalCheck{
		ObservedExportTexts:                  point13ValETwitterUniqueTexts(dependency.ValD.Dependency.ValC.NoOverclaimExportWording.ObservedCustomerExportTexts),
		ObservedTimelineTexts:                point13ValETwitterUniqueTexts(dependency.ValD.NoOverclaimProjectionWording.ObservedTimelineTexts),
		ObservedQueryTexts:                   point13ValETwitterUniqueTexts(dependency.ValD.NoOverclaimProjectionWording.ObservedQueryTexts),
		ObservedExplanationTexts:             point13ValETwitterUniqueTexts(dependency.ValD.NoOverclaimProjectionWording.ObservedExplanationTexts, []string{dependency.ValD.CustomerAuditorExplanationProjection.ExplanationText}),
		ObservedAcceptanceTexts:              point13ValETwitterUniqueTexts(dependency.ValD.Dependency.ValC.NoOverclaimExportWording.ObservedAcceptanceTexts),
		ObservedSupportOffboardingTexts:      point13ValETwitterUniqueTexts(dependency.ValD.Dependency.ValC.NoOverclaimExportWording.ObservedSupportOffboardingTexts, dependency.ValD.NoOverclaimProjectionWording.ObservedSupportOffboardingTexts),
		ObservedAITexts:                      point13ValETwitterUniqueTexts([]string{"advisory ai evidence candidate"}),
		InternalDiagnosticTexts:              point13ValETwitterUniqueTexts(dependency.ValD.Dependency.ValC.NoOverclaimExportWording.InternalDiagnosticTexts, dependency.ValD.NoOverclaimProjectionWording.InternalDiagnosticTexts),
		InternalDiagnosticsClassifiedBlocked: dependency.ValD.Dependency.ValC.NoOverclaimExportWording.InternalDiagnosticsClassifiedBlocked && dependency.ValD.NoOverclaimProjectionWording.InternalDiagnosticsClassifiedBlocked,
		AllowedSafeWording:                   point13ValEAllowedSafeWording(),
		BlockedWording:                       point13Val0ForbiddenClaims(),
		ProjectionDisclaimer:                 point13ValEProjectionDisclaimerBaseline,
	}
}

func EvaluatePoint13ValENoOverclaimFinalCheckState(model Point13ValENoOverclaimFinalCheck) string {
	if !point13ValCTextListValid(model.ObservedExportTexts) ||
		!point13ValCTextListValid(model.ObservedTimelineTexts) ||
		!point13ValCTextListValid(model.ObservedQueryTexts) ||
		!point13ValCTextListValid(model.ObservedExplanationTexts) ||
		!point13ValCTextListValid(model.ObservedAcceptanceTexts) ||
		!point13ValCTextListValid(model.ObservedSupportOffboardingTexts) ||
		!point13ValCTextListValid(model.ObservedAITexts) ||
		!point13ValCTextListValid(model.InternalDiagnosticTexts) ||
		!point13ValCTextListValid(model.AllowedSafeWording) ||
		!point13ValCTextListValid(model.BlockedWording) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValEStateBlocked
	}
	if !point12Val0ExactStringSetMatch(model.AllowedSafeWording, point13ValEAllowedSafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point13Val0ForbiddenClaims()) ||
		strings.TrimSpace(model.ProjectionDisclaimer) != point13ValEProjectionDisclaimerBaseline {
		return Point13ValEStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.ObservedExportTexts, " "),
		strings.Join(model.ObservedTimelineTexts, " "),
		strings.Join(model.ObservedQueryTexts, " "),
		strings.Join(model.ObservedExplanationTexts, " "),
		strings.Join(model.ObservedAcceptanceTexts, " "),
		strings.Join(model.ObservedSupportOffboardingTexts, " "),
		strings.Join(model.ObservedAITexts, " "),
	) {
		return Point13ValEStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(strings.Join(model.InternalDiagnosticTexts, " ")) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point13ValEStateBlocked
	}
	return Point13ValEStateActive
}

func point13ValEAuthorityBoundaryCheckModel(dependency Point13ValEDependencySnapshot) Point13ValEAuthorityBoundaryCheck {
	valC := dependency.ValD.Dependency.ValC
	valD := dependency.ValD
	return Point13ValEAuthorityBoundaryCheck{
		CheckID:                                  "authority_boundary_check_point13_vale_001",
		ValCExportReadOnly:                       valC.CustomerEvidenceExportPackage.ExportIsReadOnly,
		ValCExportOperationalOnly:                valC.CustomerEvidenceExportPackage.ExportIsOperationalEvidenceOnly,
		ValCExportCannotCreatePass:               valC.CustomerEvidenceExportPackage.ExportCannotCreatePass,
		ValCExportCannotApproveProduction:        valC.CustomerEvidenceExportPackage.ExportCannotApproveProduction,
		ValCExportCannotCertify:                  valC.CustomerEvidenceExportPackage.ExportCannotCertify,
		ValCExportCannotMutateCanonicalEvidence:  valC.CustomerEvidenceExportPackage.ExportCannotMutateCanonicalEvidence,
		ValCHandoffOperationalOnly:               valC.OperationalHandoffChecklist.HandoffIsOperationalOnly,
		ValCHandoffCannotApproveProduction:       valC.OperationalHandoffChecklist.HandoffCannotApproveProduction,
		ValCHandoffCannotAuthorizeDeployment:     valC.OperationalHandoffChecklist.HandoffCannotAuthorizeDeployment,
		ValCHandoffCannotCreatePass:              valC.OperationalHandoffChecklist.HandoffCannotCreatePass,
		ValCAcceptanceNotProductionApproval:      valC.CustomerAcceptanceTrace.AcceptanceIsNotProductionApproval,
		ValCAcceptanceNotComplianceAttestation:   valC.CustomerAcceptanceTrace.AcceptanceIsNotComplianceAttest,
		ValCAcceptanceCannotCreatePass:           valC.CustomerAcceptanceTrace.AcceptanceCannotCreatePass,
		ValCSupportCandidateOnly:                 valC.SupportOffboardingHandoffPacket.SupportMaterialCandidateOnly,
		ValCSupportCannotMutateCanonicalEvidence: valC.SupportOffboardingHandoffPacket.SupportOffboardingCannotMutateCanonical,
		ValCSupportCannotOverrideCoreDecision:    valC.SupportOffboardingHandoffPacket.SupportOffboardingCannotOverrideDecision,
		ValCSupportCannotApproveProduction:       valC.SupportOffboardingHandoffPacket.SupportOffboardingCannotApproveProduction,
		ValDTimelineReadOnly:                     valD.CustomerAuditorOperationalTimeline.TimelineReadOnly,
		ValDTimelineCannotMutateState:            valD.CustomerAuditorOperationalTimeline.TimelineCannotMutateState,
		ValDQueryReadOnly:                        valD.HandoffTraceQueryProjection.ReadOnly,
		ValDQueryMutationRequested:               valD.HandoffTraceQueryProjection.MutationRequested,
		ValDQueryWriteRequested:                  valD.HandoffTraceQueryProjection.WriteRequested,
		ValDExportReadOnly:                       valD.ExportPackageReadProjection.ReadOnly,
		ValDExportCannotOverwriteHashes:          valD.ExportPackageReadProjection.CannotOverwriteHashes,
		ValDExplanationAdvisoryOnly:              valD.CustomerAuditorExplanationProjection.AdvisoryOnly,
		ValDExplanationCannotStrengthenClaims:    valD.CustomerAuditorExplanationProjection.ExplanationCannotStrengthenClaims,
		ValDExplanationCannotApproveProduction:   valD.CustomerAuditorExplanationProjection.ExplanationCannotApproveProduction,
		ValDExplanationCannotCreatePass:          valD.CustomerAuditorExplanationProjection.ExplanationCannotCreatePass,
		ValDAccessReadOnly:                       valD.TimelineAccessBoundary.ReadOnly,
		ValDAccessMutationRequested:              valD.TimelineAccessBoundary.MutationRequested,
		ValDAIReadOnly:                           valD.AITimelineLineageProjection.ReadOnly,
		ValDAIAdvisoryOnly:                       valD.AITimelineLineageProjection.AdvisoryOnly,
		ValDAIEvidenceCandidateOnly:              valD.AITimelineLineageProjection.EvidenceCandidateOnly,
		ValDAIPassAllowed:                        valD.AITimelineLineageProjection.PassAllowed,
		ValDAIApprovalGranted:                    valD.AITimelineLineageProjection.ApprovalGranted,
		ValDAIDeploymentAuthorized:               valD.AITimelineLineageProjection.DeploymentAuthorized,
		ValDAIProductionReadinessClaimed:         valD.AITimelineLineageProjection.ProductionReadinessClaimed,
		ValDAIProductionMutationAllowed:          valD.AITimelineLineageProjection.ProductionMutationAllowed,
		ValDAICanonicalMutationAllowed:           valD.AITimelineLineageProjection.CanonicalMutationAllowed,
		ValDAIExternalAPIAllowed:                 valD.AITimelineLineageProjection.ExternalAPIAllowed,
		CurrentState:                             Point13ValEStateActive,
		ProjectionDisclaimer:                     point13ValEProjectionDisclaimerBaseline,
	}
}

func EvaluatePoint13ValEAuthorityBoundaryCheckState(model Point13ValEAuthorityBoundaryCheck, dependency Point13ValEDependencySnapshot) string {
	valC := dependency.ValD.Dependency.ValC
	valD := dependency.ValD
	if !point13ValEAuthorityBoundaryCheckRefValid(model.CheckID) ||
		!point13ValEStateValid(model.CurrentState) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValEStateBlocked
	}
	if model.ValCExportReadOnly != valC.CustomerEvidenceExportPackage.ExportIsReadOnly ||
		model.ValCExportOperationalOnly != valC.CustomerEvidenceExportPackage.ExportIsOperationalEvidenceOnly ||
		model.ValCExportCannotCreatePass != valC.CustomerEvidenceExportPackage.ExportCannotCreatePass ||
		model.ValCExportCannotApproveProduction != valC.CustomerEvidenceExportPackage.ExportCannotApproveProduction ||
		model.ValCExportCannotCertify != valC.CustomerEvidenceExportPackage.ExportCannotCertify ||
		model.ValCExportCannotMutateCanonicalEvidence != valC.CustomerEvidenceExportPackage.ExportCannotMutateCanonicalEvidence ||
		model.ValCHandoffOperationalOnly != valC.OperationalHandoffChecklist.HandoffIsOperationalOnly ||
		model.ValCHandoffCannotApproveProduction != valC.OperationalHandoffChecklist.HandoffCannotApproveProduction ||
		model.ValCHandoffCannotAuthorizeDeployment != valC.OperationalHandoffChecklist.HandoffCannotAuthorizeDeployment ||
		model.ValCHandoffCannotCreatePass != valC.OperationalHandoffChecklist.HandoffCannotCreatePass ||
		model.ValCAcceptanceNotProductionApproval != valC.CustomerAcceptanceTrace.AcceptanceIsNotProductionApproval ||
		model.ValCAcceptanceNotComplianceAttestation != valC.CustomerAcceptanceTrace.AcceptanceIsNotComplianceAttest ||
		model.ValCAcceptanceCannotCreatePass != valC.CustomerAcceptanceTrace.AcceptanceCannotCreatePass ||
		model.ValCSupportCandidateOnly != valC.SupportOffboardingHandoffPacket.SupportMaterialCandidateOnly ||
		model.ValCSupportCannotMutateCanonicalEvidence != valC.SupportOffboardingHandoffPacket.SupportOffboardingCannotMutateCanonical ||
		model.ValCSupportCannotOverrideCoreDecision != valC.SupportOffboardingHandoffPacket.SupportOffboardingCannotOverrideDecision ||
		model.ValCSupportCannotApproveProduction != valC.SupportOffboardingHandoffPacket.SupportOffboardingCannotApproveProduction ||
		model.ValDTimelineReadOnly != valD.CustomerAuditorOperationalTimeline.TimelineReadOnly ||
		model.ValDTimelineCannotMutateState != valD.CustomerAuditorOperationalTimeline.TimelineCannotMutateState ||
		model.ValDQueryReadOnly != valD.HandoffTraceQueryProjection.ReadOnly ||
		model.ValDQueryMutationRequested != valD.HandoffTraceQueryProjection.MutationRequested ||
		model.ValDQueryWriteRequested != valD.HandoffTraceQueryProjection.WriteRequested ||
		model.ValDExportReadOnly != valD.ExportPackageReadProjection.ReadOnly ||
		model.ValDExportCannotOverwriteHashes != valD.ExportPackageReadProjection.CannotOverwriteHashes ||
		model.ValDExplanationAdvisoryOnly != valD.CustomerAuditorExplanationProjection.AdvisoryOnly ||
		model.ValDExplanationCannotStrengthenClaims != valD.CustomerAuditorExplanationProjection.ExplanationCannotStrengthenClaims ||
		model.ValDExplanationCannotApproveProduction != valD.CustomerAuditorExplanationProjection.ExplanationCannotApproveProduction ||
		model.ValDExplanationCannotCreatePass != valD.CustomerAuditorExplanationProjection.ExplanationCannotCreatePass ||
		model.ValDAccessReadOnly != valD.TimelineAccessBoundary.ReadOnly ||
		model.ValDAccessMutationRequested != valD.TimelineAccessBoundary.MutationRequested ||
		model.ValDAIReadOnly != valD.AITimelineLineageProjection.ReadOnly ||
		model.ValDAIAdvisoryOnly != valD.AITimelineLineageProjection.AdvisoryOnly ||
		model.ValDAIEvidenceCandidateOnly != valD.AITimelineLineageProjection.EvidenceCandidateOnly ||
		model.ValDAIPassAllowed != valD.AITimelineLineageProjection.PassAllowed ||
		model.ValDAIApprovalGranted != valD.AITimelineLineageProjection.ApprovalGranted ||
		model.ValDAIDeploymentAuthorized != valD.AITimelineLineageProjection.DeploymentAuthorized ||
		model.ValDAIProductionReadinessClaimed != valD.AITimelineLineageProjection.ProductionReadinessClaimed ||
		model.ValDAIProductionMutationAllowed != valD.AITimelineLineageProjection.ProductionMutationAllowed ||
		model.ValDAICanonicalMutationAllowed != valD.AITimelineLineageProjection.CanonicalMutationAllowed ||
		model.ValDAIExternalAPIAllowed != valD.AITimelineLineageProjection.ExternalAPIAllowed ||
		strings.TrimSpace(model.ProjectionDisclaimer) != point13ValEProjectionDisclaimerBaseline {
		return Point13ValEStateBlocked
	}
	if !model.ValCExportReadOnly ||
		!model.ValCExportOperationalOnly ||
		!model.ValCExportCannotCreatePass ||
		!model.ValCExportCannotApproveProduction ||
		!model.ValCExportCannotCertify ||
		!model.ValCExportCannotMutateCanonicalEvidence ||
		!model.ValCHandoffOperationalOnly ||
		!model.ValCHandoffCannotApproveProduction ||
		!model.ValCHandoffCannotAuthorizeDeployment ||
		!model.ValCHandoffCannotCreatePass ||
		!model.ValCAcceptanceNotProductionApproval ||
		!model.ValCAcceptanceNotComplianceAttestation ||
		!model.ValCAcceptanceCannotCreatePass ||
		!model.ValCSupportCandidateOnly ||
		!model.ValCSupportCannotMutateCanonicalEvidence ||
		!model.ValCSupportCannotOverrideCoreDecision ||
		!model.ValCSupportCannotApproveProduction ||
		!model.ValDTimelineReadOnly ||
		!model.ValDTimelineCannotMutateState ||
		!model.ValDQueryReadOnly ||
		model.ValDQueryMutationRequested ||
		model.ValDQueryWriteRequested ||
		!model.ValDExportReadOnly ||
		!model.ValDExportCannotOverwriteHashes ||
		!model.ValDExplanationAdvisoryOnly ||
		!model.ValDExplanationCannotStrengthenClaims ||
		!model.ValDExplanationCannotApproveProduction ||
		!model.ValDExplanationCannotCreatePass ||
		!model.ValDAccessReadOnly ||
		model.ValDAccessMutationRequested ||
		!model.ValDAIReadOnly ||
		!model.ValDAIAdvisoryOnly ||
		!model.ValDAIEvidenceCandidateOnly ||
		model.ValDAIPassAllowed ||
		model.ValDAIApprovalGranted ||
		model.ValDAIDeploymentAuthorized ||
		model.ValDAIProductionReadinessClaimed ||
		model.ValDAIProductionMutationAllowed ||
		model.ValDAICanonicalMutationAllowed ||
		model.ValDAIExternalAPIAllowed {
		return Point13ValEStateBlocked
	}
	return Point13ValEStateActive
}

func point13ValETimestampIntegrityCheckModel(dependency Point13ValEDependencySnapshot) Point13ValETimestampIntegrityCheck {
	entries := dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries
	return Point13ValETimestampIntegrityCheck{
		CheckID:                "timestamp_integrity_check_point13_vale_001",
		TenantScope:            dependency.InheritedTenantScope,
		TimelineRef:            dependency.ValD.CustomerAuditorOperationalTimeline.TimelineID,
		EventRefs:              point13ValETimelineEventRefs(entries),
		AuditEventRefs:         point13ValETimelineAuditRefs(entries),
		CanonicalOccurredAts:   point13ValETimelineOccurredAts(entries),
		TimeSources:            point13ValETimelineTimeSources(entries),
		SourceMetadataRefs:     point13ValETimelineSourceMetadataRefs(entries),
		ClientReportedAts:      point13ValETimelineClientReportedAts(entries),
		VerifiedAt:             "2026-05-05T07:00:00Z",
		ClientTimeMetadataOnly: true,
		CurrentState:           Point13ValEStateActive,
		ProjectionDisclaimer:   point13ValEProjectionDisclaimerBaseline,
	}
}

func EvaluatePoint13ValETimestampIntegrityCheckState(model Point13ValETimestampIntegrityCheck, dependency Point13ValEDependencySnapshot) string {
	entries := dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries
	if !point13ValETimestampIntegrityCheckRefValid(model.CheckID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValDTimelineRefValid(model.TimelineRef) ||
		!point13ValETwitterIsolationRefListValid(model.EventRefs) ||
		!point12Val0StringListValid(model.AuditEventRefs, point12Val0AuditRefValid) ||
		!point13ValETwitterValueListValid(model.CanonicalOccurredAts, point11Val0ValidTimestamp) ||
		!point13ValETwitterValueListValid(model.TimeSources, point13ValDCanonicalTimeSourceValid) ||
		!point13ValETwitterValueListValid(model.SourceMetadataRefs, point13ValDSourceMetadataRefValid) ||
		!point11Val0ValidTimestamp(model.VerifiedAt) ||
		!model.ClientTimeMetadataOnly ||
		!point13ValEStateValid(model.CurrentState) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValEStateBlocked
	}
	if !point13ValETwitterClientMetadataOnly(model.ClientReportedAts) {
		return Point13ValEStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.TimelineRef) != strings.TrimSpace(dependency.ValD.CustomerAuditorOperationalTimeline.TimelineID) ||
		!point12Val0ExactStringSetMatch(model.EventRefs, point13ValETimelineEventRefs(entries)) ||
		!point12Val0ExactStringSetMatch(model.AuditEventRefs, point13ValETimelineAuditRefs(entries)) ||
		!point12Val0ExactStringSetMatch(model.CanonicalOccurredAts, point13ValETimelineOccurredAts(entries)) ||
		!point12Val0ExactStringSetMatch(model.TimeSources, point13ValETimelineTimeSources(entries)) ||
		!point12Val0ExactStringSetMatch(model.SourceMetadataRefs, point13ValETimelineSourceMetadataRefs(entries)) ||
		!point12Val0ExactStringSetMatch(model.ClientReportedAts, point13ValETimelineClientReportedAts(entries)) ||
		strings.TrimSpace(model.ProjectionDisclaimer) != point13ValEProjectionDisclaimerBaseline {
		return Point13ValEStateBlocked
	}
	if !point13ValDTimelineEventsOrdered(entries) || point13ValDTimelineAcceptanceBackdated(entries) || point13ValETwitterFutureDated(entries, model.VerifiedAt) {
		return Point13ValEStateBlocked
	}
	return Point13ValEStateActive
}

func point13ValETwitterIsolationCheckModel(dependency Point13ValEDependencySnapshot) Point13ValETwitterIsolationCheck {
	return Point13ValETwitterIsolationCheck{
		CheckID:                 "tenant_isolation_check_point13_vale_001",
		TenantScope:             dependency.InheritedTenantScope,
		ExportPackageRef:        dependency.ValD.Dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID,
		TimelineRef:             dependency.ValD.CustomerAuditorOperationalTimeline.TimelineID,
		QueryProjectionRef:      dependency.ValD.HandoffTraceQueryProjection.QueryProjectionID,
		ExportReadProjectionRef: dependency.ValD.ExportPackageReadProjection.ReadProjectionID,
		AccessBoundaryRef:       dependency.ValD.TimelineAccessBoundary.AccessBoundaryID,
		AudienceType:            dependency.ValD.TimelineAccessBoundary.AudienceType,
		AudienceRef:             dependency.ValD.TimelineAccessBoundary.AudienceRef,
		CustomerOwnerRef:        dependency.ValD.TimelineAccessBoundary.CustomerOwnerRef,
		AuditorOwnerRef:         dependency.ValD.TimelineAccessBoundary.AuditorOwnerRef,
		ExportedEvidenceRefs:    append([]string{}, dependency.ValD.ExportPackageReadProjection.ExportedEvidenceRefs...),
		TimelineSourceRefs:      point13ValETimelineSourceRefs(dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries),
		QueryFilterRefs:         append([]string{}, dependency.ValD.HandoffTraceQueryProjection.FilterRefs...),
		CurrentState:            Point13ValEStateActive,
		ProjectionDisclaimer:    point13ValEProjectionDisclaimerBaseline,
	}
}

func EvaluatePoint13ValETwitterIsolationCheckState(model Point13ValETwitterIsolationCheck, dependency Point13ValEDependencySnapshot) string {
	if !point13ValETwitterIsolationCheckRefValid(model.CheckID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValDTimelineRefValid(model.TimelineRef) ||
		!point13ValDQueryProjectionRefValid(model.QueryProjectionRef) ||
		!point13ValDExportReadProjectionRefValid(model.ExportReadProjectionRef) ||
		!point13ValDTimelineAccessBoundaryRefValid(model.AccessBoundaryRef) ||
		!point13ValDAudienceTypeValid(model.AudienceType) ||
		!point13ValAOwnerRefValid(model.AudienceRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point13ValAOwnerRefValid(model.AuditorOwnerRef) ||
		!point13ValBEvidenceRefsValid(model.ExportedEvidenceRefs) ||
		!point13ValETwitterIsolationRefListValid(model.TimelineSourceRefs) ||
		!point13ValDQueryFilterRefsValid(model.QueryFilterRefs) ||
		!point13ValEStateValid(model.CurrentState) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValEStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.ExportPackageRef) != strings.TrimSpace(dependency.ValD.Dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID) ||
		strings.TrimSpace(model.TimelineRef) != strings.TrimSpace(dependency.ValD.CustomerAuditorOperationalTimeline.TimelineID) ||
		strings.TrimSpace(model.QueryProjectionRef) != strings.TrimSpace(dependency.ValD.HandoffTraceQueryProjection.QueryProjectionID) ||
		strings.TrimSpace(model.ExportReadProjectionRef) != strings.TrimSpace(dependency.ValD.ExportPackageReadProjection.ReadProjectionID) ||
		strings.TrimSpace(model.AccessBoundaryRef) != strings.TrimSpace(dependency.ValD.TimelineAccessBoundary.AccessBoundaryID) ||
		strings.TrimSpace(model.AudienceType) != strings.TrimSpace(dependency.ValD.TimelineAccessBoundary.AudienceType) ||
		strings.TrimSpace(model.AudienceRef) != strings.TrimSpace(dependency.ValD.TimelineAccessBoundary.AudienceRef) ||
		strings.TrimSpace(model.CustomerOwnerRef) != strings.TrimSpace(dependency.ValD.TimelineAccessBoundary.CustomerOwnerRef) ||
		strings.TrimSpace(model.AuditorOwnerRef) != strings.TrimSpace(dependency.ValD.TimelineAccessBoundary.AuditorOwnerRef) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceRefs, dependency.ValD.ExportPackageReadProjection.ExportedEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.TimelineSourceRefs, point13ValETimelineSourceRefs(dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries)) ||
		!point12Val0ExactStringSetMatch(model.QueryFilterRefs, dependency.ValD.HandoffTraceQueryProjection.FilterRefs) ||
		strings.TrimSpace(model.ProjectionDisclaimer) != point13ValEProjectionDisclaimerBaseline {
		return Point13ValEStateBlocked
	}
	if point13ValETwitterIsolationCrossTenant(model.ExportedEvidenceRefs, model.TimelineSourceRefs, model.QueryFilterRefs) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValD.TimelineAccessBoundary.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValD.ExportPackageReadProjection.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValD.HandoffTraceQueryProjection.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValD.CustomerAuditorOperationalTimeline.TenantScope) {
		return Point13ValEStateBlocked
	}
	return Point13ValEStateActive
}

func point13ValEEvidenceIntegrityCheckModel(dependency Point13ValEDependencySnapshot) Point13ValEEvidenceIntegrityCheck {
	valC := dependency.ValD.Dependency.ValC
	valD := dependency.ValD
	return Point13ValEEvidenceIntegrityCheck{
		CheckID:                  "evidence_integrity_check_point13_vale_001",
		TenantScope:              dependency.InheritedTenantScope,
		ExportPackageRef:         valC.CustomerEvidenceExportPackage.ExportPackageID,
		ReadProjectionRef:        valD.ExportPackageReadProjection.ReadProjectionID,
		AIProjectionRef:          valD.AITimelineLineageProjection.AIProjectionID,
		ExportedEvidenceRefs:     append([]string{}, valD.ExportPackageReadProjection.ExportedEvidenceRefs...),
		ExportedEvidenceHashes:   append([]string{}, valD.ExportPackageReadProjection.ExportedEvidenceHashes...),
		ExportManifestHash:       valD.ExportPackageReadProjection.ExportManifestHash,
		AIEvidenceCandidateRef:   valD.AITimelineLineageProjection.EvidenceCandidateRef,
		AIInputEvidenceRefs:      append([]string{}, valD.AITimelineLineageProjection.InputEvidenceRefs...),
		AIInputEvidenceHashes:    append([]string{}, valD.AITimelineLineageProjection.InputEvidenceHashRefs...),
		LineageComplete:          true,
		NoRecomputeDriftDetected: true,
		CurrentState:             Point13ValEStateActive,
		ProjectionDisclaimer:     point13ValEProjectionDisclaimerBaseline,
	}
}

func EvaluatePoint13ValEEvidenceIntegrityCheckState(model Point13ValEEvidenceIntegrityCheck, dependency Point13ValEDependencySnapshot) string {
	valC := dependency.ValD.Dependency.ValC
	valD := dependency.ValD
	if !point13ValEEvidenceIntegrityCheckRefValid(model.CheckID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValDExportReadProjectionRefValid(model.ReadProjectionRef) ||
		!point13ValDAITimelineProjectionRefValid(model.AIProjectionRef) ||
		!point13ValBEvidenceRefsValid(model.ExportedEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.ExportedEvidenceHashes) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.ExportedEvidenceRefs, model.ExportedEvidenceHashes) ||
		!point12Val0HashValid(model.ExportManifestHash) ||
		!point13ValAAIEvidenceCandidateRefValid(model.AIEvidenceCandidateRef) ||
		!point13ValBEvidenceRefsValid(model.AIInputEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.AIInputEvidenceHashes) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.AIInputEvidenceRefs, model.AIInputEvidenceHashes) ||
		!model.NoRecomputeDriftDetected ||
		!point13ValEStateValid(model.CurrentState) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValEStateBlocked
	}
	if !model.LineageComplete {
		return Point13ValEStateIncomplete
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.ExportPackageRef) != strings.TrimSpace(valC.CustomerEvidenceExportPackage.ExportPackageID) ||
		strings.TrimSpace(model.ReadProjectionRef) != strings.TrimSpace(valD.ExportPackageReadProjection.ReadProjectionID) ||
		strings.TrimSpace(model.AIProjectionRef) != strings.TrimSpace(valD.AITimelineLineageProjection.AIProjectionID) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceRefs, valC.CustomerEvidenceExportPackage.ExportedEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceHashes, valC.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceRefs, valD.ExportPackageReadProjection.ExportedEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceHashes, valD.ExportPackageReadProjection.ExportedEvidenceHashes) ||
		strings.TrimSpace(model.ExportManifestHash) != strings.TrimSpace(valC.CustomerEvidenceExportPackage.ExportManifestHash) ||
		strings.TrimSpace(model.ExportManifestHash) != strings.TrimSpace(valD.ExportPackageReadProjection.ExportManifestHash) ||
		strings.TrimSpace(model.AIEvidenceCandidateRef) != strings.TrimSpace(valD.AITimelineLineageProjection.EvidenceCandidateRef) ||
		!point12Val0ExactStringSetMatch(model.AIInputEvidenceRefs, valD.AITimelineLineageProjection.InputEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.AIInputEvidenceHashes, valD.AITimelineLineageProjection.InputEvidenceHashRefs) ||
		strings.TrimSpace(model.ProjectionDisclaimer) != point13ValEProjectionDisclaimerBaseline {
		return Point13ValEStateBlocked
	}
	return Point13ValEStateActive
}

func point13ValEClosureEvaluatorModel(dependency Point13ValEDependencySnapshot) Point13ValEClosureEvaluator {
	return Point13ValEClosureEvaluator{
		ClosureEvaluatorID:       "closure_evaluator_point13_vale_001",
		TenantScope:              dependency.InheritedTenantScope,
		DependencyGateResult:     Point13ValEStateActive,
		NoOverclaimResult:        Point13ValEStateActive,
		AuthorityBoundaryResult:  Point13ValEStateActive,
		TimestampIntegrityResult: Point13ValEStateActive,
		TenantIsolationResult:    Point13ValEStateActive,
		EvidenceIntegrityResult:  Point13ValEStateActive,
		ReadOnlyProjectionOnly:   true,
		NoMutationPathsDetected:  true,
		NoPrematurePoint13Pass:   true,
		CommandsRun: []string{
			"command_run_point13_vale_gofmt_001",
			"command_run_point13_vale_go_test_formal_001",
			"command_run_point13_vale_go_test_all_001",
		},
		TestsRun: []string{
			"test_run_point13_vale_internal_formal_001",
			"test_run_point13_vale_point13_regressions_001",
			"test_run_point13_vale_go_test_all_001",
		},
		GrepsRun: []string{
			"grep_run_point13_vale_pass_token_001",
			"grep_run_point13_vale_ai_authority_001",
			"grep_run_point13_vale_forbidden_wording_001",
			"grep_run_point13_vale_mutation_flags_001",
		},
		NegativeFixturesRun: []string{
			"negative_fixture_point13_vale_dependency_gate_001",
			"negative_fixture_point13_vale_authority_boundary_001",
			"negative_fixture_point13_vale_tenant_isolation_001",
		},
		ReviewerResult:       point12ValEReviewerResultPassConfirmed,
		CurrentState:         Point13ValEStateActive,
		ProjectionDisclaimer: point13ValEProjectionDisclaimerBaseline,
	}
}

func EvaluatePoint13ValEClosureEvaluatorState(model Point13ValEClosureEvaluator, foundation Point13ValEFoundation) string {
	if !point13ValEClosureEvaluatorRefValid(model.ClosureEvaluatorID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValEStateValid(model.DependencyGateResult) ||
		!point13ValEStateValid(model.NoOverclaimResult) ||
		!point13ValEStateValid(model.AuthorityBoundaryResult) ||
		!point13ValEStateValid(model.TimestampIntegrityResult) ||
		!point13ValEStateValid(model.TenantIsolationResult) ||
		!point13ValEStateValid(model.EvidenceIntegrityResult) ||
		!point12Val0StringListValid(model.CommandsRun, point13ValECommandRunRefValid) ||
		!point12Val0StringListValid(model.TestsRun, point13ValETestRunRefValid) ||
		!point12Val0StringListValid(model.GrepsRun, point13ValEGrepRunRefValid) ||
		!point12Val0StringListValid(model.NegativeFixturesRun, point13ValENegativeFixtureRunRefValid) ||
		!point12ValEReviewerResultValid(model.ReviewerResult) ||
		!point13ValEStateValid(model.CurrentState) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValEStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(foundation.Dependency.InheritedTenantScope) ||
		strings.TrimSpace(model.DependencyGateResult) != strings.TrimSpace(foundation.DependencyState) ||
		strings.TrimSpace(model.NoOverclaimResult) != strings.TrimSpace(foundation.NoOverclaimFinalCheckState) ||
		strings.TrimSpace(model.AuthorityBoundaryResult) != strings.TrimSpace(foundation.AuthorityBoundaryCheckState) ||
		strings.TrimSpace(model.TimestampIntegrityResult) != strings.TrimSpace(foundation.TimestampIntegrityCheckState) ||
		strings.TrimSpace(model.TenantIsolationResult) != strings.TrimSpace(foundation.TenantIsolationCheckState) ||
		strings.TrimSpace(model.EvidenceIntegrityResult) != strings.TrimSpace(foundation.EvidenceIntegrityCheckState) ||
		strings.TrimSpace(model.ProjectionDisclaimer) != point13ValEProjectionDisclaimerBaseline {
		return Point13ValEStateBlocked
	}
	componentState := point13ValEAggregateDerivedStates(
		model.DependencyGateResult,
		model.NoOverclaimResult,
		model.AuthorityBoundaryResult,
		model.TimestampIntegrityResult,
		model.TenantIsolationResult,
		model.EvidenceIntegrityResult,
	)
	switch componentState {
	case Point13ValEStateBlocked:
		return Point13ValEStateBlocked
	case Point13ValEStateReviewRequired:
		return Point13ValEStateReviewRequired
	case Point13ValEStateIncomplete:
		return Point13ValEStateIncomplete
	case Point13ValEStateActive:
	default:
		return Point13ValEStateBlocked
	}
	if !model.ReadOnlyProjectionOnly || !model.NoMutationPathsDetected || !model.NoPrematurePoint13Pass {
		return Point13ValEStateBlocked
	}
	switch strings.TrimSpace(model.ReviewerResult) {
	case point12ValEReviewerResultPassConfirmed:
		return Point13ValEStateActive
	case point12ValEReviewerResultReviewRequired:
		return Point13ValEStateReviewRequired
	default:
		return Point13ValEStateBlocked
	}
}

func point13ValEPassClosureManifestModel(dependency Point13ValEDependencySnapshot) Point13PassClosureManifest {
	valC := dependency.ValD.Dependency.ValC
	valD := dependency.ValD
	return Point13PassClosureManifest{
		CurrentState:                    Point13ValEStateActive,
		ClosureManifestID:               "closure_manifest_point13_vale_001",
		PointID:                         point13Val0PointID,
		WaveID:                          point13ValEWaveID,
		Scope:                           point13ValEScope,
		DependencyGateResult:            Point13ValEStateActive,
		ClosureEvaluatorResult:          Point13ValEStateActive,
		NoOverclaimResult:               Point13ValEStateActive,
		AuthorityBoundaryResult:         Point13ValEStateActive,
		TimestampIntegrityResult:        Point13ValEStateActive,
		TenantIsolationResult:           Point13ValEStateActive,
		EvidenceIntegrityResult:         Point13ValEStateActive,
		ValDCurrentState:                dependency.ValDCurrentState,
		ValDTimelineState:               dependency.ValDCustomerAuditorOperationalTimelineState,
		ValDQueryProjectionState:        dependency.ValDHandoffTraceQueryProjectionState,
		ValDExportReadProjectionState:   dependency.ValDExportPackageReadProjectionState,
		ValDExplanationProjectionState:  dependency.ValDCustomerAuditorExplanationProjectionState,
		ValDAccessBoundaryState:         dependency.ValDTimelineAccessBoundaryState,
		ValDAIProjectionState:           dependency.ValDAITimelineLineageProjectionState,
		ValDNoOverclaimState:            dependency.ValDNoOverclaimState,
		ValCExportPackageRef:            valC.CustomerEvidenceExportPackage.ExportPackageID,
		ValCHandoffChecklistRef:         valC.OperationalHandoffChecklist.HandoffChecklistID,
		ValCAcceptanceTraceRef:          valC.CustomerAcceptanceTrace.AcceptanceTraceID,
		ValCSupportOffboardingPacketRef: valC.SupportOffboardingHandoffPacket.SupportOffboardingPacketID,
		ValCAIExportSummaryRef:          valC.AIEvidenceExportLineageSummary.AIExportSummaryID,
		ValDTimelineRef:                 valD.CustomerAuditorOperationalTimeline.TimelineID,
		ValDQueryProjectionRef:          valD.HandoffTraceQueryProjection.QueryProjectionID,
		ValDExportReadProjectionRef:     valD.ExportPackageReadProjection.ReadProjectionID,
		ValDExplanationProjectionRef:    valD.CustomerAuditorExplanationProjection.ExplanationProjectionID,
		ValDAccessBoundaryRef:           valD.TimelineAccessBoundary.AccessBoundaryID,
		ValDAIProjectionRef:             valD.AITimelineLineageProjection.AIProjectionID,
		TenantScope:                     dependency.InheritedTenantScope,
		ExportedEvidenceRefs:            append([]string{}, valC.CustomerEvidenceExportPackage.ExportedEvidenceRefs...),
		ExportedEvidenceHashes:          append([]string{}, valC.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs...),
		ExportManifestHash:              valC.CustomerEvidenceExportPackage.ExportManifestHash,
		RetentionClassRef:               valC.CustomerEvidenceExportPackage.RetentionClassRef,
		AIModelOrRuleVersionRef:         dependency.InheritedAIModelOrRuleVersionRef,
		AIPermissionManifestHash:        dependency.InheritedAIPermissionManifestHash,
		CommandsRun: []string{
			"command_run_point13_vale_gofmt_001",
			"command_run_point13_vale_go_test_formal_001",
			"command_run_point13_vale_go_test_all_001",
		},
		TestsRun: []string{
			"test_run_point13_vale_internal_formal_001",
			"test_run_point13_vale_point13_regressions_001",
			"test_run_point13_vale_go_test_all_001",
		},
		GrepsRun: []string{
			"grep_run_point13_vale_pass_token_001",
			"grep_run_point13_vale_ai_authority_001",
			"grep_run_point13_vale_forbidden_wording_001",
			"grep_run_point13_vale_mutation_flags_001",
		},
		NegativeFixturesRun: []string{
			"negative_fixture_point13_vale_dependency_gate_001",
			"negative_fixture_point13_vale_authority_boundary_001",
			"negative_fixture_point13_vale_timestamp_integrity_001",
		},
		ReviewerResult:       point12ValEReviewerResultPassConfirmed,
		GeneratedAt:          "2026-05-05T07:05:00Z",
		Point13PassAllowed:   true,
		Point13PassToken:     point13ValEPoint13PassToken,
		ProjectionDisclaimer: point13ValEProjectionDisclaimerBaseline,
	}
}

func point13ValEPassCandidate(foundation Point13ValEFoundation) bool {
	return foundation.DependencyState == Point13ValEStateActive &&
		foundation.ClosureEvaluatorState == Point13ValEStateActive &&
		foundation.NoOverclaimFinalCheckState == Point13ValEStateActive &&
		foundation.AuthorityBoundaryCheckState == Point13ValEStateActive &&
		foundation.TimestampIntegrityCheckState == Point13ValEStateActive &&
		foundation.TenantIsolationCheckState == Point13ValEStateActive &&
		foundation.EvidenceIntegrityCheckState == Point13ValEStateActive
}

func point13ValEPassClosureManifestStateAndReasons(model Point13PassClosureManifest, foundation Point13ValEFoundation, expectedPassAllowed bool) (string, []string) {
	reasons := []string{}
	valC := foundation.Dependency.ValD.Dependency.ValC
	valD := foundation.Dependency.ValD
	if !point13ValEClosureManifestRefValid(model.ClosureManifestID) ||
		strings.TrimSpace(model.PointID) != point13Val0PointID ||
		strings.TrimSpace(model.WaveID) != point13ValEWaveID ||
		strings.TrimSpace(model.Scope) != point13ValEScope ||
		!point13ValEStateValid(model.CurrentState) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "pass_closure_manifest_identity_invalid")
	}
	if !point13ValEStateValid(model.DependencyGateResult) ||
		!point13ValEStateValid(model.ClosureEvaluatorResult) ||
		!point13ValEStateValid(model.NoOverclaimResult) ||
		!point13ValEStateValid(model.AuthorityBoundaryResult) ||
		!point13ValEStateValid(model.TimestampIntegrityResult) ||
		!point13ValEStateValid(model.TenantIsolationResult) ||
		!point13ValEStateValid(model.EvidenceIntegrityResult) ||
		strings.TrimSpace(model.ValDCurrentState) != Point13ValDStateActive ||
		strings.TrimSpace(model.ValDTimelineState) != Point13ValDStateActive ||
		strings.TrimSpace(model.ValDQueryProjectionState) != Point13ValDStateActive ||
		strings.TrimSpace(model.ValDExportReadProjectionState) != Point13ValDStateActive ||
		strings.TrimSpace(model.ValDExplanationProjectionState) != Point13ValDStateActive ||
		strings.TrimSpace(model.ValDAccessBoundaryState) != Point13ValDStateActive ||
		strings.TrimSpace(model.ValDAIProjectionState) != Point13ValDStateActive ||
		strings.TrimSpace(model.ValDNoOverclaimState) != Point13ValDStateActive {
		reasons = append(reasons, "pass_closure_manifest_result_state_invalid")
	}
	if !point13ValCExportPackageRefValid(model.ValCExportPackageRef) ||
		!point13ValCHandoffChecklistRefValid(model.ValCHandoffChecklistRef) ||
		!point13ValCAcceptanceTraceRefValid(model.ValCAcceptanceTraceRef) ||
		!point13ValCSupportOffboardingPacketRefValid(model.ValCSupportOffboardingPacketRef) ||
		!point13ValCAIExportSummaryRefValid(model.ValCAIExportSummaryRef) ||
		!point13ValDTimelineRefValid(model.ValDTimelineRef) ||
		!point13ValDQueryProjectionRefValid(model.ValDQueryProjectionRef) ||
		!point13ValDExportReadProjectionRefValid(model.ValDExportReadProjectionRef) ||
		!point13ValDExplanationProjectionRefValid(model.ValDExplanationProjectionRef) ||
		!point13ValDTimelineAccessBoundaryRefValid(model.ValDAccessBoundaryRef) ||
		!point13ValDAITimelineProjectionRefValid(model.ValDAIProjectionRef) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point13ValBEvidenceRefsValid(model.ExportedEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.ExportedEvidenceHashes) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.ExportedEvidenceRefs, model.ExportedEvidenceHashes) ||
		!point12Val0HashValid(model.ExportManifestHash) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12Val0VersionRefValid(model.AIModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.AIPermissionManifestHash) ||
		!point12Val0StringListValid(model.CommandsRun, point13ValECommandRunRefValid) ||
		!point12Val0StringListValid(model.TestsRun, point13ValETestRunRefValid) ||
		!point12Val0StringListValid(model.GrepsRun, point13ValEGrepRunRefValid) ||
		!point12Val0StringListValid(model.NegativeFixturesRun, point13ValENegativeFixtureRunRefValid) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) {
		reasons = append(reasons, "pass_closure_manifest_required_fields_invalid")
	}
	if strings.TrimSpace(model.ValCExportPackageRef) != strings.TrimSpace(valC.CustomerEvidenceExportPackage.ExportPackageID) ||
		strings.TrimSpace(model.ValCHandoffChecklistRef) != strings.TrimSpace(valC.OperationalHandoffChecklist.HandoffChecklistID) ||
		strings.TrimSpace(model.ValCAcceptanceTraceRef) != strings.TrimSpace(valC.CustomerAcceptanceTrace.AcceptanceTraceID) ||
		strings.TrimSpace(model.ValCSupportOffboardingPacketRef) != strings.TrimSpace(valC.SupportOffboardingHandoffPacket.SupportOffboardingPacketID) ||
		strings.TrimSpace(model.ValCAIExportSummaryRef) != strings.TrimSpace(valC.AIEvidenceExportLineageSummary.AIExportSummaryID) ||
		strings.TrimSpace(model.ValDTimelineRef) != strings.TrimSpace(valD.CustomerAuditorOperationalTimeline.TimelineID) ||
		strings.TrimSpace(model.ValDQueryProjectionRef) != strings.TrimSpace(valD.HandoffTraceQueryProjection.QueryProjectionID) ||
		strings.TrimSpace(model.ValDExportReadProjectionRef) != strings.TrimSpace(valD.ExportPackageReadProjection.ReadProjectionID) ||
		strings.TrimSpace(model.ValDExplanationProjectionRef) != strings.TrimSpace(valD.CustomerAuditorExplanationProjection.ExplanationProjectionID) ||
		strings.TrimSpace(model.ValDAccessBoundaryRef) != strings.TrimSpace(valD.TimelineAccessBoundary.AccessBoundaryID) ||
		strings.TrimSpace(model.ValDAIProjectionRef) != strings.TrimSpace(valD.AITimelineLineageProjection.AIProjectionID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(foundation.Dependency.InheritedTenantScope) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceRefs, valC.CustomerEvidenceExportPackage.ExportedEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceHashes, valC.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs) ||
		strings.TrimSpace(model.ExportManifestHash) != strings.TrimSpace(valC.CustomerEvidenceExportPackage.ExportManifestHash) ||
		strings.TrimSpace(model.RetentionClassRef) != strings.TrimSpace(valC.CustomerEvidenceExportPackage.RetentionClassRef) ||
		strings.TrimSpace(model.AIModelOrRuleVersionRef) != strings.TrimSpace(foundation.Dependency.InheritedAIModelOrRuleVersionRef) ||
		strings.TrimSpace(model.AIPermissionManifestHash) != strings.TrimSpace(foundation.Dependency.InheritedAIPermissionManifestHash) {
		reasons = append(reasons, "pass_closure_manifest_binding_mismatch")
	}
	if !point12ValEReviewerResultValid(model.ReviewerResult) {
		reasons = append(reasons, "pass_closure_manifest_reviewer_result_invalid")
	}
	componentState := point13ValEAggregateDerivedStates(
		model.DependencyGateResult,
		model.ClosureEvaluatorResult,
		model.NoOverclaimResult,
		model.AuthorityBoundaryResult,
		model.TimestampIntegrityResult,
		model.TenantIsolationResult,
		model.EvidenceIntegrityResult,
	)
	reviewerState := Point13ValEStateActive
	switch strings.TrimSpace(model.ReviewerResult) {
	case point12ValEReviewerResultPassConfirmed:
	case point12ValEReviewerResultReviewRequired:
		reviewerState = Point13ValEStateReviewRequired
	default:
		reasons = append(reasons, "pass_closure_manifest_reviewer_not_final")
	}
	if strings.TrimSpace(model.Point13PassToken) != "" && strings.TrimSpace(model.Point13PassToken) != point13ValEPoint13PassToken {
		reasons = append(reasons, "pass_closure_manifest_token_invalid")
	}
	aggregateState := point13ValEAggregateDerivedStates(componentState, reviewerState)
	if aggregateState != Point13ValEStateActive {
		if model.Point13PassAllowed || strings.TrimSpace(model.Point13PassToken) != "" {
			reasons = append(reasons, "pass_closure_manifest_token_present_before_final_happy_path")
		}
		if len(reasons) > 0 {
			return Point13ValEStateBlocked, reasons
		}
		switch aggregateState {
		case Point13ValEStateReviewRequired:
			return Point13ValEStateReviewRequired, []string{"pass_closure_manifest_waiting_for_reviewable_closure"}
		case Point13ValEStateIncomplete:
			return Point13ValEStateIncomplete, []string{"pass_closure_manifest_waiting_for_complete_closure"}
		default:
			return Point13ValEStateBlocked, []string{"pass_closure_manifest_foundation_gates_not_active"}
		}
	}
	if !expectedPassAllowed || !point13ValEPassCandidate(foundation) || !model.Point13PassAllowed || strings.TrimSpace(model.Point13PassToken) != point13ValEPoint13PassToken {
		reasons = append(reasons, "pass_closure_manifest_pass_not_fully_authorized")
	}
	if len(reasons) > 0 {
		return Point13ValEStateBlocked, reasons
	}
	return Point13ValEStateActive, nil
}

func point13ValEComponentStates(model Point13ValEFoundation) []string {
	return []string{
		model.DependencyState,
		model.ClosureEvaluatorState,
		model.NoOverclaimFinalCheckState,
		model.AuthorityBoundaryCheckState,
		model.TimestampIntegrityCheckState,
		model.TenantIsolationCheckState,
		model.EvidenceIntegrityCheckState,
		model.PassClosureManifestState,
	}
}

func EvaluatePoint13ValEState(model Point13ValEFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValEStateBlocked
	}
	hasReview := false
	hasIncomplete := false
	for _, state := range point13ValEComponentStates(model) {
		switch strings.TrimSpace(state) {
		case Point13ValEStateBlocked:
			return Point13ValEStateBlocked
		case Point13ValEStateReviewRequired:
			hasReview = true
		case Point13ValEStateIncomplete:
			hasIncomplete = true
		case Point13ValEStateActive:
		default:
			return Point13ValEStateBlocked
		}
	}
	if hasReview {
		return Point13ValEStateReviewRequired
	}
	if hasIncomplete {
		return Point13ValEStateIncomplete
	}
	return Point13ValEStateActive
}

func point13ValECurrentState(model Point13ValEFoundation) string {
	state := EvaluatePoint13ValEState(model)
	if state != Point13ValEStateActive {
		return state
	}
	if model.Point13PassAllowed &&
		strings.TrimSpace(model.Point13PassToken) == point13ValEPoint13PassToken &&
		strings.TrimSpace(model.PassClosureManifest.ReviewerResult) == point12ValEReviewerResultPassConfirmed {
		return Point13ValEStatePassConfirmed
	}
	return Point13ValEStateActive
}

func point13ValEBlockingReasons(model Point13ValEFoundation) []string {
	reasons := []string{}
	componentStates := map[string]string{
		"dependency":            model.DependencyState,
		"closure_evaluator":     model.ClosureEvaluatorState,
		"no_overclaim":          model.NoOverclaimFinalCheckState,
		"authority_boundary":    model.AuthorityBoundaryCheckState,
		"timestamp_integrity":   model.TimestampIntegrityCheckState,
		"tenant_isolation":      model.TenantIsolationCheckState,
		"evidence_integrity":    model.EvidenceIntegrityCheckState,
		"pass_closure_manifest": model.PassClosureManifestState,
	}
	for name, state := range componentStates {
		switch strings.TrimSpace(state) {
		case Point13ValEStateBlocked, Point13ValEStateIncomplete:
			reasons = append(reasons, name+":"+state)
		}
	}
	if model.Point13PassAllowed && strings.TrimSpace(model.Point13PassToken) != point13ValEPoint13PassToken {
		reasons = append(reasons, "point13_pass_token_invalid")
	}
	return reasons
}

func point13ValEFoundationModelFromUpstream(valD Point13ValDFoundation) Point13ValEFoundation {
	dependency := point13ValEDependencySnapshotFromUpstream(valD)
	noOverclaim := point13ValENoOverclaimFinalCheckModel(dependency)
	authority := point13ValEAuthorityBoundaryCheckModel(dependency)
	timestamp := point13ValETimestampIntegrityCheckModel(dependency)
	tenant := point13ValETwitterIsolationCheckModel(dependency)
	evidence := point13ValEEvidenceIntegrityCheckModel(dependency)
	return Point13ValEFoundation{
		CurrentState:                 Point13ValEStatePassConfirmed,
		ProjectionDisclaimer:         point13ValEProjectionDisclaimerBaseline,
		DependencyState:              Point13ValEStateActive,
		ClosureEvaluatorState:        Point13ValEStateActive,
		NoOverclaimFinalCheckState:   Point13ValEStateActive,
		AuthorityBoundaryCheckState:  Point13ValEStateActive,
		TimestampIntegrityCheckState: Point13ValEStateActive,
		TenantIsolationCheckState:    Point13ValEStateActive,
		EvidenceIntegrityCheckState:  Point13ValEStateActive,
		PassClosureManifestState:     Point13ValEStateActive,
		Point13PassAllowed:           true,
		Point13PassToken:             point13ValEPoint13PassToken,
		Dependency:                   dependency,
		ClosureEvaluator:             point13ValEClosureEvaluatorModel(dependency),
		NoOverclaimFinalCheck:        noOverclaim,
		AuthorityBoundaryCheck:       authority,
		TimestampIntegrityCheck:      timestamp,
		TenantIsolationCheck:         tenant,
		EvidenceIntegrityCheck:       evidence,
		PassClosureManifest:          point13ValEPassClosureManifestModel(dependency),
	}
}

func Point13ValEFoundationModel() Point13ValEFoundation {
	valD := ComputePoint13ValDFoundation(Point13ValDFoundationModel())
	return point13ValEFoundationModelFromUpstream(valD)
}

func ComputePoint13ValEFoundation(model Point13ValEFoundation) Point13ValEFoundation {
	dependencyState, dependencyReasons := point13ValEDependencyStateAndReasons(model.Dependency)
	model.DependencyState = dependencyState
	noOverclaimState := EvaluatePoint13ValENoOverclaimFinalCheckState(model.NoOverclaimFinalCheck)
	model.NoOverclaimFinalCheckState = noOverclaimState
	authorityState := EvaluatePoint13ValEAuthorityBoundaryCheckState(model.AuthorityBoundaryCheck, model.Dependency)
	model.AuthorityBoundaryCheckState = authorityState
	timestampState := EvaluatePoint13ValETimestampIntegrityCheckState(model.TimestampIntegrityCheck, model.Dependency)
	model.TimestampIntegrityCheckState = timestampState
	tenantState := EvaluatePoint13ValETwitterIsolationCheckState(model.TenantIsolationCheck, model.Dependency)
	model.TenantIsolationCheckState = tenantState
	evidenceState := EvaluatePoint13ValEEvidenceIntegrityCheckState(model.EvidenceIntegrityCheck, model.Dependency)
	model.EvidenceIntegrityCheckState = evidenceState
	model.ClosureEvaluator.TenantScope = model.Dependency.InheritedTenantScope
	model.ClosureEvaluator.DependencyGateResult = model.DependencyState
	model.ClosureEvaluator.NoOverclaimResult = model.NoOverclaimFinalCheckState
	model.ClosureEvaluator.AuthorityBoundaryResult = model.AuthorityBoundaryCheckState
	model.ClosureEvaluator.TimestampIntegrityResult = model.TimestampIntegrityCheckState
	model.ClosureEvaluator.TenantIsolationResult = model.TenantIsolationCheckState
	model.ClosureEvaluator.EvidenceIntegrityResult = model.EvidenceIntegrityCheckState
	closureState := EvaluatePoint13ValEClosureEvaluatorState(model.ClosureEvaluator, model)
	model.ClosureEvaluatorState = closureState
	passCandidate := point13ValEPassCandidate(model)
	model.PassClosureManifest.DependencyGateResult = model.DependencyState
	model.PassClosureManifest.ClosureEvaluatorResult = model.ClosureEvaluatorState
	model.PassClosureManifest.NoOverclaimResult = model.NoOverclaimFinalCheckState
	model.PassClosureManifest.AuthorityBoundaryResult = model.AuthorityBoundaryCheckState
	model.PassClosureManifest.TimestampIntegrityResult = model.TimestampIntegrityCheckState
	model.PassClosureManifest.TenantIsolationResult = model.TenantIsolationCheckState
	model.PassClosureManifest.EvidenceIntegrityResult = model.EvidenceIntegrityCheckState
	if !passCandidate {
		model.PassClosureManifest.Point13PassAllowed = false
		model.PassClosureManifest.Point13PassToken = ""
	}
	manifestState, manifestReasons := point13ValEPassClosureManifestStateAndReasons(model.PassClosureManifest, model, passCandidate)
	model.PassClosureManifestState = manifestState
	model.PassClosureManifest.CurrentState = manifestState
	model.Point13PassAllowed = passCandidate &&
		manifestState == Point13ValEStateActive &&
		closureState == Point13ValEStateActive &&
		strings.TrimSpace(model.PassClosureManifest.ReviewerResult) == point12ValEReviewerResultPassConfirmed &&
		model.PassClosureManifest.Point13PassAllowed &&
		strings.TrimSpace(model.PassClosureManifest.Point13PassToken) == point13ValEPoint13PassToken
	if model.Point13PassAllowed {
		model.Point13PassToken = point13ValEPoint13PassToken
	} else {
		model.Point13PassToken = ""
	}
	model.CurrentState = point13ValECurrentState(model)
	model.BlockingReasons = point13ValEBlockingReasons(model)
	model.ReviewPrerequisites = nil
	if model.DependencyState == Point13ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, dependencyReasons...)
	}
	if model.PassClosureManifestState == Point13ValEStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, manifestReasons...)
	}
	return model
}
