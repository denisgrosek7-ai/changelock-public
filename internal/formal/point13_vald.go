package formal

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

const (
	Point13ValDStateActive         = "point13_vald_customer_auditor_timeline_active"
	Point13ValDStateBlocked        = "point13_vald_customer_auditor_timeline_blocked"
	Point13ValDStateReviewRequired = "point13_vald_customer_auditor_timeline_review_required"
	Point13ValDStateIncomplete     = "point13_vald_customer_auditor_timeline_incomplete"
)

const (
	point13ValDWaveID                       = "val_d"
	point13ValDPreviousWaveID               = point13ValCWaveID
	point13ValDProjectionDisclaimerBaseline = "projection_only not_canonical_truth point13_vald_customer_auditor_timeline"

	point13ValDTimelineKindExportCreated              = "export_created"
	point13ValDTimelineKindRedactionApplied           = "redaction_applied"
	point13ValDTimelineKindHandoffStarted             = "handoff_started"
	point13ValDTimelineKindHandoffAcknowledged        = "handoff_acknowledged"
	point13ValDTimelineKindCustomerAcceptanceRecorded = "customer_acceptance_recorded"
	point13ValDTimelineKindSupportOffboardingPrepared = "support_offboarding_prepared"
	point13ValDTimelineKindAILineageIncluded          = "ai_lineage_included"

	point13ValDQueryKindTimelineSummary       = "timeline_summary"
	point13ValDQueryKindHandoffStatus         = "handoff_status"
	point13ValDQueryKindExportPackageRead     = "export_package_read"
	point13ValDQueryKindAcceptanceLimitations = "acceptance_limitations"
	point13ValDQueryKindSupportOffboarding    = "support_offboarding"
	point13ValDQueryKindAILineage             = "ai_lineage"

	point13ValDExplanationKindOperationalTimeline = "operational_timeline"
	point13ValDExplanationKindHandoffLimits       = "handoff_limits"
	point13ValDExplanationKindExportSummary       = "export_summary"
	point13ValDExplanationKindAcceptanceSummary   = "acceptance_summary"
	point13ValDExplanationKindAILineageSummary    = "ai_lineage_summary"

	point13ValDAudienceCustomer = "customer"
	point13ValDAudienceAuditor  = "auditor"

	point13ValDTimeSourceServerUTC             = "server_utc"
	point13ValDTimeSourceApprovedCustomerClock = "approved_customer_time_source"
)

type Point13ValDDependencySnapshot struct {
	ValCCurrentState                       string                `json:"valc_current_state"`
	ValCDependencyState                    string                `json:"valc_dependency_state"`
	ValCCustomerEvidenceExportPackageState string                `json:"valc_customer_evidence_export_package_state"`
	ValCRedactionSafeDisclosureState       string                `json:"valc_redaction_safe_disclosure_state"`
	ValCOperationalHandoffChecklistState   string                `json:"valc_operational_handoff_checklist_state"`
	ValCCustomerAcceptanceTraceState       string                `json:"valc_customer_acceptance_trace_state"`
	ValCSupportOffboardingHandoffState     string                `json:"valc_support_offboarding_handoff_state"`
	ValCAIEvidenceExportLineageState       string                `json:"valc_ai_evidence_export_lineage_state"`
	ValCNoOverclaimState                   string                `json:"valc_no_overclaim_state"`
	ValCPointID                            string                `json:"valc_point_id"`
	ValCWaveID                             string                `json:"valc_wave_id"`
	ValCDependencyComputedFromUpstream     bool                  `json:"valc_dependency_computed_from_upstream"`
	ValCPoint13PassSeen                    bool                  `json:"valc_point13_pass_seen"`
	InheritedValBCurrentState              string                `json:"inherited_valb_current_state"`
	InheritedValACurrentState              string                `json:"inherited_vala_current_state"`
	InheritedVal0CurrentState              string                `json:"inherited_val0_current_state"`
	InheritedPoint12CurrentState           string                `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState        string                `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState       string                `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult         string                `json:"inherited_point12_reviewer_result"`
	InheritedTenantScope                   string                `json:"inherited_tenant_scope"`
	InheritedAIModelOrRuleVersionRef       string                `json:"inherited_ai_model_or_rule_version_ref"`
	InheritedAIPermissionManifestHash      string                `json:"inherited_ai_permission_manifest_hash"`
	SnapshotFromComputedOutput             bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                    []string              `json:"review_prerequisites,omitempty"`
	ValC                                   Point13ValCFoundation `json:"valc"`
}

type Point13ValDTimelineEvent struct {
	EventID             string `json:"event_id"`
	EventKind           string `json:"event_kind"`
	SourceRef           string `json:"source_ref"`
	AuditEventRef       string `json:"audit_event_ref"`
	CanonicalOccurredAt string `json:"canonical_occurred_at"`
	TimeSource          string `json:"time_source"`
	ClientReportedAt    string `json:"client_reported_at,omitempty"`
	SourceMetadataRef   string `json:"source_metadata_ref"`
	ReadOnly            bool   `json:"read_only"`
}

type Point13ValDCustomerAuditorOperationalTimeline struct {
	TimelineID                  string                     `json:"timeline_id"`
	TenantScope                 string                     `json:"tenant_scope"`
	ExportPackageRef            string                     `json:"export_package_ref"`
	HandoffChecklistRef         string                     `json:"handoff_checklist_ref"`
	CustomerAcceptanceRef       string                     `json:"customer_acceptance_ref"`
	SupportOffboardingRef       string                     `json:"support_offboarding_ref"`
	AITimelineSummaryRef        string                     `json:"ai_timeline_summary_ref"`
	TimelineEntries             []Point13ValDTimelineEvent `json:"timeline_entries,omitempty"`
	TimelineHash                string                     `json:"timeline_hash"`
	TimelineReadOnly            bool                       `json:"timeline_read_only"`
	TimelineCannotMutateState   bool                       `json:"timeline_cannot_mutate_state"`
	RedactionLimitationsVisible bool                       `json:"redaction_limitations_visible"`
}

type Point13ValDHandoffTraceQueryProjection struct {
	QueryProjectionID   string   `json:"query_projection_id"`
	TenantScope         string   `json:"tenant_scope"`
	TimelineRef         string   `json:"timeline_ref"`
	ExportPackageRef    string   `json:"export_package_ref"`
	HandoffChecklistRef string   `json:"handoff_checklist_ref"`
	QueryKind           string   `json:"query_kind"`
	FilterRefs          []string `json:"filter_refs,omitempty"`
	AuditEventRef       string   `json:"audit_event_ref"`
	ReadOnly            bool     `json:"read_only"`
	MutationRequested   bool     `json:"mutation_requested"`
	WriteRequested      bool     `json:"write_requested"`
	ProjectionHash      string   `json:"projection_hash"`
}

type Point13ValDExportPackageReadProjection struct {
	ReadProjectionID       string   `json:"read_projection_id"`
	TenantScope            string   `json:"tenant_scope"`
	ExportPackageRef       string   `json:"export_package_ref"`
	RedactionBoundaryRef   string   `json:"redaction_boundary_ref"`
	ExportedEvidenceRefs   []string `json:"exported_evidence_refs,omitempty"`
	ExportedEvidenceHashes []string `json:"exported_evidence_hashes,omitempty"`
	ExportManifestHash     string   `json:"export_manifest_hash"`
	RetentionClassRef      string   `json:"retention_class_ref"`
	AuditEventRef          string   `json:"audit_event_ref"`
	ProjectionHash         string   `json:"projection_hash"`
	ReadOnly               bool     `json:"read_only"`
	CannotOverwriteHashes  bool     `json:"cannot_overwrite_hashes"`
	LimitationsVisible     bool     `json:"limitations_visible"`
	VisibleLimitations     []string `json:"visible_limitations,omitempty"`
}

type Point13ValDAuditorAnnotation struct {
	AnnotationID       string `json:"annotation_id"`
	AnnotatorRef       string `json:"annotator_ref"`
	Text               string `json:"text"`
	AuditEventRef      string `json:"audit_event_ref"`
	AnnotationOnly     bool   `json:"annotation_only"`
	ApprovesProduction bool   `json:"approves_production"`
	ChangesExportState bool   `json:"changes_export_state"`
}

type Point13ValDCustomerAuditorExplanationProjection struct {
	ExplanationProjectionID            string                         `json:"explanation_projection_id"`
	TenantScope                        string                         `json:"tenant_scope"`
	TimelineRef                        string                         `json:"timeline_ref"`
	QueryProjectionRef                 string                         `json:"query_projection_ref"`
	ExportReadProjectionRef            string                         `json:"export_read_projection_ref"`
	ExplanationKind                    string                         `json:"explanation_kind"`
	ExplanationText                    string                         `json:"explanation_text"`
	VisibleLimitations                 []string                       `json:"visible_limitations,omitempty"`
	AuditorAnnotations                 []Point13ValDAuditorAnnotation `json:"auditor_annotations,omitempty"`
	AuditEventRef                      string                         `json:"audit_event_ref"`
	AdvisoryOnly                       bool                           `json:"advisory_only"`
	ExplanationCannotStrengthenClaims  bool                           `json:"explanation_cannot_strengthen_claims"`
	ExplanationCannotApproveProduction bool                           `json:"explanation_cannot_approve_production"`
	ExplanationCannotCreatePass        bool                           `json:"explanation_cannot_create_pass"`
	ProjectionHash                     string                         `json:"projection_hash"`
}

type Point13ValDTimelineAccessBoundary struct {
	AccessBoundaryID           string `json:"access_boundary_id"`
	TenantScope                string `json:"tenant_scope"`
	AudienceType               string `json:"audience_type"`
	AudienceRef                string `json:"audience_ref"`
	CustomerOwnerRef           string `json:"customer_owner_ref"`
	AuditorOwnerRef            string `json:"auditor_owner_ref"`
	TimelineRef                string `json:"timeline_ref"`
	QueryProjectionRef         string `json:"query_projection_ref"`
	ExportReadProjectionRef    string `json:"export_read_projection_ref"`
	AccessAuditEventRef        string `json:"access_audit_event_ref"`
	ReadOnly                   bool   `json:"read_only"`
	MutationRequested          bool   `json:"mutation_requested"`
	CrossTenantAccessRequested bool   `json:"cross_tenant_access_requested"`
}

type Point13ValDAITimelineLineageProjection struct {
	AIProjectionID                string   `json:"ai_projection_id"`
	TenantScope                   string   `json:"tenant_scope"`
	TimelineRef                   string   `json:"timeline_ref"`
	AIExportSummaryRef            string   `json:"ai_export_summary_ref"`
	AIOutputType                  string   `json:"ai_output_type"`
	EvidenceCandidateRef          string   `json:"evidence_candidate_ref"`
	InputEvidenceRefs             []string `json:"input_evidence_refs,omitempty"`
	InputEvidenceHashRefs         []string `json:"input_evidence_hash_refs,omitempty"`
	ModelOrRuleVersionRef         string   `json:"model_or_rule_version_ref"`
	PermissionManifestHash        string   `json:"permission_manifest_hash"`
	AuditEventRef                 string   `json:"audit_event_ref"`
	ReadOnly                      bool     `json:"read_only"`
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
	CanStrengthenTimelineClaim    bool     `json:"can_strengthen_timeline_claim"`
	CanSatisfyAcceptanceByItself  bool     `json:"can_satisfy_acceptance_by_itself"`
}

type Point13ValDNoOverclaimProjectionWording struct {
	ObservedTimelineTexts                []string `json:"observed_timeline_texts,omitempty"`
	ObservedQueryTexts                   []string `json:"observed_query_texts,omitempty"`
	ObservedReadProjectionTexts          []string `json:"observed_read_projection_texts,omitempty"`
	ObservedExplanationTexts             []string `json:"observed_explanation_texts,omitempty"`
	ObservedSupportOffboardingTexts      []string `json:"observed_support_offboarding_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point13ValDFoundation struct {
	CurrentState                              string                                          `json:"current_state"`
	BlockingReasons                           []string                                        `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                       []string                                        `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer                      string                                          `json:"projection_disclaimer"`
	DependencyState                           string                                          `json:"dependency_state"`
	CustomerAuditorOperationalTimelineState   string                                          `json:"customer_auditor_operational_timeline_state"`
	HandoffTraceQueryProjectionState          string                                          `json:"handoff_trace_query_projection_state"`
	ExportPackageReadProjectionState          string                                          `json:"export_package_read_projection_state"`
	CustomerAuditorExplanationProjectionState string                                          `json:"customer_auditor_explanation_projection_state"`
	TimelineAccessBoundaryState               string                                          `json:"timeline_access_boundary_state"`
	AITimelineLineageProjectionState          string                                          `json:"ai_timeline_lineage_projection_state"`
	NoOverclaimState                          string                                          `json:"no_overclaim_state"`
	Dependency                                Point13ValDDependencySnapshot                   `json:"dependency"`
	CustomerAuditorOperationalTimeline        Point13ValDCustomerAuditorOperationalTimeline   `json:"customer_auditor_operational_timeline"`
	HandoffTraceQueryProjection               Point13ValDHandoffTraceQueryProjection          `json:"handoff_trace_query_projection"`
	ExportPackageReadProjection               Point13ValDExportPackageReadProjection          `json:"export_package_read_projection"`
	CustomerAuditorExplanationProjection      Point13ValDCustomerAuditorExplanationProjection `json:"customer_auditor_explanation_projection"`
	TimelineAccessBoundary                    Point13ValDTimelineAccessBoundary               `json:"timeline_access_boundary"`
	AITimelineLineageProjection               Point13ValDAITimelineLineageProjection          `json:"ai_timeline_lineage_projection"`
	NoOverclaimProjectionWording              Point13ValDNoOverclaimProjectionWording         `json:"no_overclaim_projection_wording"`
}

func point13ValDStates() []string {
	return []string{
		Point13ValDStateActive,
		Point13ValDStateBlocked,
		Point13ValDStateReviewRequired,
		Point13ValDStateIncomplete,
	}
}

func point13ValDStateValid(value string) bool {
	return point13Val0RawExactOneOf(value, point13ValDStates())
}

func point13ValDTimelineRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "timeline_point13_vald_", "customer_auditor_timeline_")
}

func point13ValDTimelineEventRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "timeline_event_")
}

func point13ValDQueryProjectionRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "handoff_query_projection_", "query_projection_")
}

func point13ValDExportReadProjectionRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "export_read_projection_")
}

func point13ValDExplanationProjectionRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "explanation_projection_")
}

func point13ValDTimelineAccessBoundaryRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "timeline_access_boundary_")
}

func point13ValDAITimelineProjectionRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "ai_timeline_projection_")
}

func point13ValDAuditorAnnotationRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "auditor_annotation_")
}

func point13ValDSourceMetadataRefValid(value string) bool {
	return point13Val0OperationalRefValid(value, "source_metadata_", "time_source_metadata_")
}

func point13ValDTimelineEventKinds() []string {
	return []string{
		point13ValDTimelineKindExportCreated,
		point13ValDTimelineKindRedactionApplied,
		point13ValDTimelineKindHandoffStarted,
		point13ValDTimelineKindHandoffAcknowledged,
		point13ValDTimelineKindCustomerAcceptanceRecorded,
		point13ValDTimelineKindSupportOffboardingPrepared,
		point13ValDTimelineKindAILineageIncluded,
	}
}

func point13ValDTimelineEventKindValid(value string) bool {
	return point13Val0RawExactOneOf(value, point13ValDTimelineEventKinds())
}

func point13ValDQueryKinds() []string {
	return []string{
		point13ValDQueryKindTimelineSummary,
		point13ValDQueryKindHandoffStatus,
		point13ValDQueryKindExportPackageRead,
		point13ValDQueryKindAcceptanceLimitations,
		point13ValDQueryKindSupportOffboarding,
		point13ValDQueryKindAILineage,
	}
}

func point13ValDQueryKindValid(value string) bool {
	return point13Val0RawExactOneOf(value, point13ValDQueryKinds())
}

func point13ValDExplanationKinds() []string {
	return []string{
		point13ValDExplanationKindOperationalTimeline,
		point13ValDExplanationKindHandoffLimits,
		point13ValDExplanationKindExportSummary,
		point13ValDExplanationKindAcceptanceSummary,
		point13ValDExplanationKindAILineageSummary,
	}
}

func point13ValDExplanationKindValid(value string) bool {
	return point13Val0RawExactOneOf(value, point13ValDExplanationKinds())
}

func point13ValDAudienceTypes() []string {
	return []string{
		point13ValDAudienceCustomer,
		point13ValDAudienceAuditor,
	}
}

func point13ValDAudienceTypeValid(value string) bool {
	return point13Val0RawExactOneOf(value, point13ValDAudienceTypes())
}

func point13ValDCanonicalTimeSources() []string {
	return []string{
		point13ValDTimeSourceServerUTC,
		point13ValDTimeSourceApprovedCustomerClock,
	}
}

func point13ValDCanonicalTimeSourceValid(value string) bool {
	return point13Val0RawExactOneOf(value, point13ValDCanonicalTimeSources())
}

func point13ValDRawStringInSet(values []string, target string) bool {
	if !formalRawExactNonEmpty(target) {
		return false
	}
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func point13ValDTimelineEventListValid(values []Point13ValDTimelineEvent) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point13ValDTimelineEventRefValid(value.EventID) || !point13ValDTimelineEventKindValid(value.EventKind) {
			return false
		}
		key := value.EventID
		if _, exists := seen[key]; exists {
			return false
		}
		seen[key] = struct{}{}
	}
	return true
}

func point13ValDQueryFilterRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point11Val0IdentityValueValid)
}

func point13ValDAuditorAnnotationsValid(values []Point13ValDAuditorAnnotation) bool {
	seen := map[string]struct{}{}
	for _, value := range values {
		if !formalRawExactTokenValid(value.AnnotationID, point13ValDAuditorAnnotationRefValid) ||
			!point13Val0RawOwnerRefValid(value.AnnotatorRef) ||
			strings.TrimSpace(value.Text) == "" ||
			!point12Val0AuditRefValid(value.AuditEventRef) {
			return false
		}
		key := value.AnnotationID
		if _, exists := seen[key]; exists {
			return false
		}
		seen[key] = struct{}{}
	}
	return true
}

func point13ValDComputedBindingHash(parts ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(parts, "\n")))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func point13ValDComputedTimelineHash(model Point13ValDCustomerAuditorOperationalTimeline) string {
	parts := []string{
		model.TimelineID,
		model.TenantScope,
		model.ExportPackageRef,
		model.HandoffChecklistRef,
		model.CustomerAcceptanceRef,
		model.SupportOffboardingRef,
		model.AITimelineSummaryRef,
		strconv.FormatBool(model.TimelineReadOnly),
		strconv.FormatBool(model.TimelineCannotMutateState),
		strconv.FormatBool(model.RedactionLimitationsVisible),
	}
	for _, entry := range model.TimelineEntries {
		parts = append(parts,
			entry.EventID,
			entry.EventKind,
			entry.SourceRef,
			entry.AuditEventRef,
			entry.CanonicalOccurredAt,
			entry.TimeSource,
			entry.ClientReportedAt,
			entry.SourceMetadataRef,
			strconv.FormatBool(entry.ReadOnly),
		)
	}
	return point13ValDComputedBindingHash(parts...)
}

func point13ValDComputedQueryHash(model Point13ValDHandoffTraceQueryProjection) string {
	return point13ValDComputedBindingHash(
		model.QueryProjectionID,
		model.TenantScope,
		model.TimelineRef,
		model.ExportPackageRef,
		model.HandoffChecklistRef,
		model.QueryKind,
		strings.Join(model.FilterRefs, ","),
		model.AuditEventRef,
		strconv.FormatBool(model.ReadOnly),
		strconv.FormatBool(model.MutationRequested),
		strconv.FormatBool(model.WriteRequested),
	)
}

func point13ValDComputedExportReadHash(model Point13ValDExportPackageReadProjection) string {
	return point13ValDComputedBindingHash(
		model.ReadProjectionID,
		model.TenantScope,
		model.ExportPackageRef,
		model.RedactionBoundaryRef,
		strings.Join(model.ExportedEvidenceRefs, ","),
		strings.Join(model.ExportedEvidenceHashes, ","),
		model.ExportManifestHash,
		model.RetentionClassRef,
		model.AuditEventRef,
		strconv.FormatBool(model.ReadOnly),
		strconv.FormatBool(model.CannotOverwriteHashes),
		strconv.FormatBool(model.LimitationsVisible),
		strings.Join(model.VisibleLimitations, ","),
	)
}

func point13ValDComputedExplanationHash(model Point13ValDCustomerAuditorExplanationProjection) string {
	parts := []string{
		model.ExplanationProjectionID,
		model.TenantScope,
		model.TimelineRef,
		model.QueryProjectionRef,
		model.ExportReadProjectionRef,
		model.ExplanationKind,
		model.ExplanationText,
		strings.Join(model.VisibleLimitations, ","),
		model.AuditEventRef,
		strconv.FormatBool(model.AdvisoryOnly),
		strconv.FormatBool(model.ExplanationCannotStrengthenClaims),
		strconv.FormatBool(model.ExplanationCannotApproveProduction),
		strconv.FormatBool(model.ExplanationCannotCreatePass),
	}
	for _, annotation := range model.AuditorAnnotations {
		parts = append(parts,
			annotation.AnnotationID,
			annotation.AnnotatorRef,
			annotation.Text,
			annotation.AuditEventRef,
			strconv.FormatBool(annotation.AnnotationOnly),
			strconv.FormatBool(annotation.ApprovesProduction),
			strconv.FormatBool(annotation.ChangesExportState),
		)
	}
	return point13ValDComputedBindingHash(parts...)
}

func point13ValDAllowedSafeWording() []string {
	return []string{
		"customer auditor operational timeline",
		"handoff trace query projection",
		"export package read projection",
		"customer auditor explanation projection",
		"support offboarding handoff",
		"advisory ai evidence candidate",
		"evidence support for customer/auditor review",
	}
}

func point13ValDExpectedVisibleLimitations(valC Point13ValCFoundation) []string {
	limitations := append([]string{}, valC.CustomerAcceptanceTrace.AcceptedLimitations...)
	if len(valC.RedactionSafeDisclosureBoundary.RedactedFields) > 0 {
		limitations = append(limitations, "redaction_applied_customer_safe_disclosure")
	}
	return limitations
}

func point13ValDValCPayloadContainsPointPass(valC Point13ValCFoundation) bool {
	payload, err := json.Marshal(valC)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point13Val0BlockedPoint13PassToken)
}

func point13ValDDependencySnapshotFromUpstream(valC Point13ValCFoundation) Point13ValDDependencySnapshot {
	return Point13ValDDependencySnapshot{
		ValCCurrentState:                       valC.CurrentState,
		ValCDependencyState:                    valC.DependencyState,
		ValCCustomerEvidenceExportPackageState: valC.CustomerEvidenceExportPackageState,
		ValCRedactionSafeDisclosureState:       valC.RedactionSafeDisclosureState,
		ValCOperationalHandoffChecklistState:   valC.OperationalHandoffChecklistState,
		ValCCustomerAcceptanceTraceState:       valC.CustomerAcceptanceTraceState,
		ValCSupportOffboardingHandoffState:     valC.SupportOffboardingHandoffState,
		ValCAIEvidenceExportLineageState:       valC.AIEvidenceExportLineageState,
		ValCNoOverclaimState:                   valC.NoOverclaimState,
		ValCPointID:                            point13Val0PointID,
		ValCWaveID:                             point13ValCWaveID,
		ValCDependencyComputedFromUpstream:     valC.Dependency.SnapshotFromComputedOutput,
		ValCPoint13PassSeen:                    point13ValDValCPayloadContainsPointPass(valC),
		InheritedValBCurrentState:              valC.Dependency.ValBCurrentState,
		InheritedValACurrentState:              valC.Dependency.InheritedValACurrentState,
		InheritedVal0CurrentState:              valC.Dependency.InheritedVal0CurrentState,
		InheritedPoint12CurrentState:           valC.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:        valC.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:       valC.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint12ReviewerResult:         valC.Dependency.InheritedPoint12ReviewerResult,
		InheritedTenantScope:                   valC.Dependency.InheritedTenantScope,
		InheritedAIModelOrRuleVersionRef:       valC.Dependency.InheritedAIModelOrRuleVersionRef,
		InheritedAIPermissionManifestHash:      valC.Dependency.InheritedAIPermissionManifestHash,
		SnapshotFromComputedOutput:             true,
		ReviewPrerequisites:                    append([]string{}, valC.ReviewPrerequisites...),
		ValC:                                   valC,
	}
}

func point13ValDDependencySnapshotModel() Point13ValDDependencySnapshot {
	return point13ValDDependencySnapshotFromUpstream(ComputePoint13ValCFoundation(Point13ValCFoundationModel()))
}

func point13ValDDependencyStateAndReasons(model Point13ValDDependencySnapshot) (string, []string) {
	reviewReasons := []string{}
	blockedReasons := []string{}
	incompleteReasons := []string{}
	recomputedValC := ComputePoint13ValCFoundation(model.ValC)

	if !model.SnapshotFromComputedOutput || !model.ValCDependencyComputedFromUpstream {
		blockedReasons = append(blockedReasons, "valc_dependency_not_computed_from_upstream")
	}
	if !point13Val0RawExactOneOf(model.ValCCurrentState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCDependencyState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCCustomerEvidenceExportPackageState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCRedactionSafeDisclosureState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCOperationalHandoffChecklistState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCCustomerAcceptanceTraceState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCSupportOffboardingHandoffState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCAIEvidenceExportLineageState, point13ValCStates()) ||
		!point13Val0RawExactOneOf(model.ValCNoOverclaimState, point13ValCStates()) ||
		model.ValCPointID != point13Val0PointID ||
		model.ValCWaveID != point13ValCWaveID ||
		!point13Val0RawExactOneOf(model.InheritedValBCurrentState, point13ValBStates()) ||
		!point13Val0RawExactOneOf(model.InheritedValACurrentState, point13ValAStates()) ||
		!point13Val0RawExactOneOf(model.InheritedVal0CurrentState, point13Val0States()) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point13Val0RawExactOneOf(model.InheritedPoint12ReviewerResult, []string{point12ValEReviewerResultPassConfirmed, point12ValEReviewerResultPass, point12ValEReviewerResultReviewRequired}) ||
		!point13Val0RawScopeValid(model.InheritedTenantScope) ||
		!point12Val0VersionRefValid(model.InheritedAIModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.InheritedAIPermissionManifestHash) {
		blockedReasons = append(blockedReasons, "dependency_snapshot_identity_invalid")
	}
	if model.ValCPoint13PassSeen {
		blockedReasons = append(blockedReasons, "valc_point13_pass_seen")
	}
	if model.ValC.CurrentState != recomputedValC.CurrentState ||
		model.ValC.DependencyState != recomputedValC.DependencyState ||
		model.ValC.CustomerEvidenceExportPackageState != recomputedValC.CustomerEvidenceExportPackageState ||
		model.ValC.RedactionSafeDisclosureState != recomputedValC.RedactionSafeDisclosureState ||
		model.ValC.OperationalHandoffChecklistState != recomputedValC.OperationalHandoffChecklistState ||
		model.ValC.CustomerAcceptanceTraceState != recomputedValC.CustomerAcceptanceTraceState ||
		model.ValC.SupportOffboardingHandoffState != recomputedValC.SupportOffboardingHandoffState ||
		model.ValC.AIEvidenceExportLineageState != recomputedValC.AIEvidenceExportLineageState ||
		model.ValC.NoOverclaimState != recomputedValC.NoOverclaimState ||
		model.ValC.Dependency.SnapshotFromComputedOutput != recomputedValC.Dependency.SnapshotFromComputedOutput ||
		model.ValC.Dependency.ValBCurrentState != recomputedValC.Dependency.ValBCurrentState ||
		model.ValC.Dependency.InheritedValACurrentState != recomputedValC.Dependency.InheritedValACurrentState ||
		model.ValC.Dependency.InheritedVal0CurrentState != recomputedValC.Dependency.InheritedVal0CurrentState ||
		model.ValC.Dependency.InheritedPoint12CurrentState != recomputedValC.Dependency.InheritedPoint12CurrentState ||
		model.ValC.Dependency.InheritedPoint12DependencyState != recomputedValC.Dependency.InheritedPoint12DependencyState ||
		model.ValC.Dependency.InheritedPoint12PassClosureState != recomputedValC.Dependency.InheritedPoint12PassClosureState ||
		model.ValC.Dependency.InheritedPoint12ReviewerResult != recomputedValC.Dependency.InheritedPoint12ReviewerResult ||
		model.ValC.Dependency.InheritedTenantScope != recomputedValC.Dependency.InheritedTenantScope ||
		model.ValC.Dependency.InheritedAIModelOrRuleVersionRef != recomputedValC.Dependency.InheritedAIModelOrRuleVersionRef ||
		model.ValC.Dependency.InheritedAIPermissionManifestHash != recomputedValC.Dependency.InheritedAIPermissionManifestHash {
		blockedReasons = append(blockedReasons, "valc_recomputed_snapshot_mismatch")
	}
	if model.ValCCurrentState != recomputedValC.CurrentState ||
		model.ValCDependencyState != recomputedValC.DependencyState ||
		model.ValCCustomerEvidenceExportPackageState != recomputedValC.CustomerEvidenceExportPackageState ||
		model.ValCRedactionSafeDisclosureState != recomputedValC.RedactionSafeDisclosureState ||
		model.ValCOperationalHandoffChecklistState != recomputedValC.OperationalHandoffChecklistState ||
		model.ValCCustomerAcceptanceTraceState != recomputedValC.CustomerAcceptanceTraceState ||
		model.ValCSupportOffboardingHandoffState != recomputedValC.SupportOffboardingHandoffState ||
		model.ValCAIEvidenceExportLineageState != recomputedValC.AIEvidenceExportLineageState ||
		model.ValCNoOverclaimState != recomputedValC.NoOverclaimState ||
		model.ValCDependencyComputedFromUpstream != recomputedValC.Dependency.SnapshotFromComputedOutput ||
		model.InheritedValBCurrentState != recomputedValC.Dependency.ValBCurrentState ||
		model.InheritedValACurrentState != recomputedValC.Dependency.InheritedValACurrentState ||
		model.InheritedVal0CurrentState != recomputedValC.Dependency.InheritedVal0CurrentState ||
		model.InheritedPoint12CurrentState != recomputedValC.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != recomputedValC.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != recomputedValC.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != recomputedValC.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedTenantScope != recomputedValC.Dependency.InheritedTenantScope ||
		model.InheritedAIModelOrRuleVersionRef != recomputedValC.Dependency.InheritedAIModelOrRuleVersionRef ||
		model.InheritedAIPermissionManifestHash != recomputedValC.Dependency.InheritedAIPermissionManifestHash {
		blockedReasons = append(blockedReasons, "dependency_snapshot_binding_mismatch")
	}
	for _, state := range []string{
		model.ValCCurrentState,
		model.ValCDependencyState,
		model.ValCCustomerEvidenceExportPackageState,
		model.ValCRedactionSafeDisclosureState,
		model.ValCOperationalHandoffChecklistState,
		model.ValCCustomerAcceptanceTraceState,
		model.ValCSupportOffboardingHandoffState,
		model.ValCAIEvidenceExportLineageState,
		model.ValCNoOverclaimState,
	} {
		switch state {
		case Point13ValCStateBlocked:
			blockedReasons = append(blockedReasons, "valc_component_blocked")
		case Point13ValCStateReviewRequired:
			reviewReasons = append(reviewReasons, "valc_component_review_required")
		case Point13ValCStateIncomplete:
			incompleteReasons = append(incompleteReasons, "valc_component_incomplete")
		}
	}
	for _, state := range []string{
		model.InheritedValBCurrentState,
		model.InheritedValACurrentState,
		model.InheritedVal0CurrentState,
	} {
		switch state {
		case Point13ValBStateBlocked, Point13ValAStateBlocked, Point13Val0StateBlocked:
			blockedReasons = append(blockedReasons, "inherited_val_component_blocked")
		case Point13ValBStateReviewRequired, Point13ValAStateReviewRequired, Point13Val0StateReviewRequired:
			reviewReasons = append(reviewReasons, "inherited_val_component_review_required")
		case Point13ValBStateIncomplete, Point13ValAStateIncomplete, Point13Val0StateIncomplete:
			incompleteReasons = append(incompleteReasons, "inherited_val_component_incomplete")
		}
	}
	switch model.InheritedPoint12CurrentState {
	case Point12ValEStateBlocked:
		blockedReasons = append(blockedReasons, "point12_inherited_blocked")
	case Point12ValEStateReviewRequired:
		reviewReasons = append(reviewReasons, "point12_inherited_review_required")
	case Point12ValEStateIncomplete:
		incompleteReasons = append(incompleteReasons, "point12_inherited_incomplete")
	}
	if model.InheritedPoint12CurrentState == Point12ValEStatePassConfirmed &&
		(model.InheritedPoint12DependencyState != Point12ValEStateActive ||
			model.InheritedPoint12PassClosureState != Point12ValEStateActive ||
			model.InheritedPoint12ReviewerResult != point12ValEReviewerResultPassConfirmed) {
		blockedReasons = append(blockedReasons, "point12_inherited_not_pass_confirmed")
	}
	if len(blockedReasons) > 0 {
		return Point13ValDStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point13ValDStateReviewRequired, reviewReasons
	}
	if len(incompleteReasons) > 0 {
		return Point13ValDStateIncomplete, incompleteReasons
	}
	return Point13ValDStateActive, nil
}

func point13ValDTimelineModel(dependency Point13ValDDependencySnapshot) Point13ValDCustomerAuditorOperationalTimeline {
	valC := dependency.ValC
	model := Point13ValDCustomerAuditorOperationalTimeline{
		TimelineID:                  "timeline_point13_vald_001",
		TenantScope:                 dependency.InheritedTenantScope,
		ExportPackageRef:            valC.CustomerEvidenceExportPackage.ExportPackageID,
		HandoffChecklistRef:         valC.OperationalHandoffChecklist.HandoffChecklistID,
		CustomerAcceptanceRef:       valC.CustomerAcceptanceTrace.AcceptanceTraceID,
		SupportOffboardingRef:       valC.SupportOffboardingHandoffPacket.SupportOffboardingPacketID,
		AITimelineSummaryRef:        "ai_timeline_projection_point13_vald_001",
		TimelineReadOnly:            true,
		TimelineCannotMutateState:   true,
		RedactionLimitationsVisible: true,
		TimelineEntries: []Point13ValDTimelineEvent{
			{EventID: "timeline_event_point13_vald_001", EventKind: point13ValDTimelineKindExportCreated, SourceRef: valC.CustomerEvidenceExportPackage.ExportPackageID, AuditEventRef: valC.CustomerEvidenceExportPackage.AuditEventRef, CanonicalOccurredAt: "2026-05-05T06:00:00Z", TimeSource: point13ValDTimeSourceServerUTC, SourceMetadataRef: "source_metadata_point13_vald_export_001", ReadOnly: true},
			{EventID: "timeline_event_point13_vald_002", EventKind: point13ValDTimelineKindRedactionApplied, SourceRef: valC.RedactionSafeDisclosureBoundary.RedactionBoundaryID, AuditEventRef: valC.RedactionSafeDisclosureBoundary.RedactionAuditEventRef, CanonicalOccurredAt: "2026-05-05T06:01:00Z", TimeSource: point13ValDTimeSourceServerUTC, SourceMetadataRef: "source_metadata_point13_vald_redaction_001", ReadOnly: true},
			{EventID: "timeline_event_point13_vald_003", EventKind: point13ValDTimelineKindHandoffStarted, SourceRef: valC.OperationalHandoffChecklist.HandoffChecklistID, AuditEventRef: valC.OperationalHandoffChecklist.AuditEventRefs[0], CanonicalOccurredAt: "2026-05-05T06:05:00Z", TimeSource: point13ValDTimeSourceServerUTC, SourceMetadataRef: "source_metadata_point13_vald_handoff_start_001", ReadOnly: true},
			{EventID: "timeline_event_point13_vald_004", EventKind: point13ValDTimelineKindHandoffAcknowledged, SourceRef: valC.OperationalHandoffChecklist.HandoffChecklistID, AuditEventRef: valC.OperationalHandoffChecklist.AuditEventRefs[1], CanonicalOccurredAt: "2026-05-05T06:06:00Z", TimeSource: point13ValDTimeSourceServerUTC, SourceMetadataRef: "source_metadata_point13_vald_handoff_ack_001", ReadOnly: true},
			{EventID: "timeline_event_point13_vald_005", EventKind: point13ValDTimelineKindCustomerAcceptanceRecorded, SourceRef: valC.CustomerAcceptanceTrace.AcceptanceTraceID, AuditEventRef: valC.CustomerAcceptanceTrace.AuditEventRefs[0], CanonicalOccurredAt: "2026-05-05T06:10:00Z", TimeSource: point13ValDTimeSourceServerUTC, SourceMetadataRef: "source_metadata_point13_vald_acceptance_001", ReadOnly: true},
			{EventID: "timeline_event_point13_vald_006", EventKind: point13ValDTimelineKindSupportOffboardingPrepared, SourceRef: valC.SupportOffboardingHandoffPacket.SupportOffboardingPacketID, AuditEventRef: valC.SupportOffboardingHandoffPacket.AuditEventRefs[0], CanonicalOccurredAt: "2026-05-05T06:12:00Z", TimeSource: point13ValDTimeSourceServerUTC, SourceMetadataRef: "source_metadata_point13_vald_offboarding_001", ReadOnly: true},
			{EventID: "timeline_event_point13_vald_007", EventKind: point13ValDTimelineKindAILineageIncluded, SourceRef: valC.AIEvidenceExportLineageSummary.AIExportSummaryID, AuditEventRef: valC.AIEvidenceExportLineageSummary.AuditEventRef, CanonicalOccurredAt: "2026-05-05T06:13:00Z", TimeSource: point13ValDTimeSourceServerUTC, SourceMetadataRef: "source_metadata_point13_vald_ai_001", ReadOnly: true},
		},
	}
	model.TimelineHash = point13ValDComputedTimelineHash(model)
	return model
}

func point13ValDQueryProjectionModel(timeline Point13ValDCustomerAuditorOperationalTimeline, dependency Point13ValDDependencySnapshot) Point13ValDHandoffTraceQueryProjection {
	model := Point13ValDHandoffTraceQueryProjection{
		QueryProjectionID:   "handoff_query_projection_point13_vald_001",
		TenantScope:         dependency.InheritedTenantScope,
		TimelineRef:         timeline.TimelineID,
		ExportPackageRef:    dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID,
		HandoffChecklistRef: dependency.ValC.OperationalHandoffChecklist.HandoffChecklistID,
		QueryKind:           point13ValDQueryKindTimelineSummary,
		FilterRefs: []string{
			dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID,
			dependency.ValC.CustomerAcceptanceTrace.AcceptanceTraceID,
		},
		AuditEventRef:     "audit_point13_vald_query_001",
		ReadOnly:          true,
		MutationRequested: false,
		WriteRequested:    false,
	}
	model.ProjectionHash = point13ValDComputedQueryHash(model)
	return model
}

func point13ValDExportReadProjectionModel(dependency Point13ValDDependencySnapshot) Point13ValDExportPackageReadProjection {
	valC := dependency.ValC
	model := Point13ValDExportPackageReadProjection{
		ReadProjectionID:       "export_read_projection_point13_vald_001",
		TenantScope:            dependency.InheritedTenantScope,
		ExportPackageRef:       valC.CustomerEvidenceExportPackage.ExportPackageID,
		RedactionBoundaryRef:   valC.RedactionSafeDisclosureBoundary.RedactionBoundaryID,
		ExportedEvidenceRefs:   append([]string{}, valC.CustomerEvidenceExportPackage.ExportedEvidenceRefs...),
		ExportedEvidenceHashes: append([]string{}, valC.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs...),
		ExportManifestHash:     valC.CustomerEvidenceExportPackage.ExportManifestHash,
		RetentionClassRef:      valC.CustomerEvidenceExportPackage.RetentionClassRef,
		AuditEventRef:          "audit_point13_vald_export_read_001",
		ReadOnly:               true,
		CannotOverwriteHashes:  true,
		LimitationsVisible:     true,
		VisibleLimitations:     point13ValDExpectedVisibleLimitations(valC),
	}
	model.ProjectionHash = point13ValDComputedExportReadHash(model)
	return model
}

func point13ValDExplanationProjectionModel(timeline Point13ValDCustomerAuditorOperationalTimeline, query Point13ValDHandoffTraceQueryProjection, exportRead Point13ValDExportPackageReadProjection, dependency Point13ValDDependencySnapshot) Point13ValDCustomerAuditorExplanationProjection {
	model := Point13ValDCustomerAuditorExplanationProjection{
		ExplanationProjectionID: "explanation_projection_point13_vald_001",
		TenantScope:             dependency.InheritedTenantScope,
		TimelineRef:             timeline.TimelineID,
		QueryProjectionRef:      query.QueryProjectionID,
		ExportReadProjectionRef: exportRead.ReadProjectionID,
		ExplanationKind:         point13ValDExplanationKindOperationalTimeline,
		ExplanationText:         "customer auditor explanation projection provides bounded operational timeline and handoff limits",
		VisibleLimitations:      append([]string{}, exportRead.VisibleLimitations...),
		AuditorAnnotations: []Point13ValDAuditorAnnotation{
			{
				AnnotationID:       "auditor_annotation_point13_vald_001",
				AnnotatorRef:       "owner_point13_vald_auditor_001",
				Text:               "annotation only for auditor review context",
				AuditEventRef:      "audit_point13_vald_annotation_001",
				AnnotationOnly:     true,
				ApprovesProduction: false,
				ChangesExportState: false,
			},
		},
		AuditEventRef:                      "audit_point13_vald_explanation_001",
		AdvisoryOnly:                       true,
		ExplanationCannotStrengthenClaims:  true,
		ExplanationCannotApproveProduction: true,
		ExplanationCannotCreatePass:        true,
	}
	model.ProjectionHash = point13ValDComputedExplanationHash(model)
	return model
}

func point13ValDAccessBoundaryModel(timeline Point13ValDCustomerAuditorOperationalTimeline, query Point13ValDHandoffTraceQueryProjection, exportRead Point13ValDExportPackageReadProjection, dependency Point13ValDDependencySnapshot) Point13ValDTimelineAccessBoundary {
	return Point13ValDTimelineAccessBoundary{
		AccessBoundaryID:           "timeline_access_boundary_point13_vald_001",
		TenantScope:                dependency.InheritedTenantScope,
		AudienceType:               point13ValDAudienceCustomer,
		AudienceRef:                dependency.ValC.CustomerEvidenceExportPackage.CustomerOwnerRef,
		CustomerOwnerRef:           dependency.ValC.CustomerEvidenceExportPackage.CustomerOwnerRef,
		AuditorOwnerRef:            "owner_point13_vald_auditor_001",
		TimelineRef:                timeline.TimelineID,
		QueryProjectionRef:         query.QueryProjectionID,
		ExportReadProjectionRef:    exportRead.ReadProjectionID,
		AccessAuditEventRef:        "audit_point13_vald_access_001",
		ReadOnly:                   true,
		MutationRequested:          false,
		CrossTenantAccessRequested: false,
	}
}

func point13ValDAITimelineProjectionModel(timeline Point13ValDCustomerAuditorOperationalTimeline, dependency Point13ValDDependencySnapshot) Point13ValDAITimelineLineageProjection {
	ai := dependency.ValC.AIEvidenceExportLineageSummary
	return Point13ValDAITimelineLineageProjection{
		AIProjectionID:                "ai_timeline_projection_point13_vald_001",
		TenantScope:                   dependency.InheritedTenantScope,
		TimelineRef:                   timeline.TimelineID,
		AIExportSummaryRef:            ai.AIExportSummaryID,
		AIOutputType:                  ai.AIOutputType,
		EvidenceCandidateRef:          ai.EvidenceCandidateRef,
		InputEvidenceRefs:             append([]string{}, ai.InputEvidenceRefs...),
		InputEvidenceHashRefs:         append([]string{}, ai.InputEvidenceHashRefs...),
		ModelOrRuleVersionRef:         ai.ModelOrRuleVersionRef,
		PermissionManifestHash:        ai.PermissionManifestHash,
		AuditEventRef:                 ai.AuditEventRef,
		ReadOnly:                      true,
		AdvisoryOnly:                  true,
		EvidenceCandidateOnly:         true,
		PassAllowed:                   false,
		ApprovalGranted:               false,
		DeploymentAuthorized:          false,
		ProductionReadinessClaimed:    false,
		ProductionMutationAllowed:     false,
		CanonicalMutationAllowed:      false,
		ExternalAPIAllowed:            false,
		ExternalAPIGovernanceEventRef: "",
		CanStrengthenTimelineClaim:    false,
		CanSatisfyAcceptanceByItself:  false,
	}
}

func point13ValDNoOverclaimModel() Point13ValDNoOverclaimProjectionWording {
	return Point13ValDNoOverclaimProjectionWording{
		ObservedTimelineTexts: []string{
			"customer auditor operational timeline",
		},
		ObservedQueryTexts: []string{
			"handoff trace query projection",
		},
		ObservedReadProjectionTexts: []string{
			"export package read projection",
		},
		ObservedExplanationTexts: []string{
			"customer auditor explanation projection",
		},
		ObservedSupportOffboardingTexts: []string{
			"support offboarding handoff",
		},
		InternalDiagnosticTexts: []string{
			"blocked wording remains denylisted internally",
		},
		InternalDiagnosticsClassifiedBlocked: true,
		AllowedSafeWording:                   point13ValDAllowedSafeWording(),
		BlockedWording:                       point13Val0ForbiddenClaims(),
		ProjectionDisclaimer:                 point13ValDProjectionDisclaimerBaseline,
	}
}

func point13ValDTimelineEventKindsFromEntries(entries []Point13ValDTimelineEvent) []string {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.EventKind)
	}
	return values
}

func point13ValDTimelineEventsOrdered(entries []Point13ValDTimelineEvent) bool {
	for i := 1; i < len(entries); i++ {
		comparison, ok := point13ValDTimelineTimestampCompare(entries[i-1].CanonicalOccurredAt, entries[i].CanonicalOccurredAt)
		if !ok || comparison > 0 {
			return false
		}
	}
	return true
}

func point13ValDTimelineTimestampCompare(left, right string) (int, bool) {
	if !point12Val0RawTimestampValid(left) || !point12Val0RawTimestampValid(right) {
		return 0, false
	}
	leftTime, leftOK := point14Val0ParsedTime(left)
	rightTime, rightOK := point14Val0ParsedTime(right)
	if !leftOK || !rightOK {
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

func point13ValDTimelineAcceptanceBackdated(entries []Point13ValDTimelineEvent) bool {
	exportAt := ""
	handoffAckAt := ""
	acceptanceAt := ""
	for _, entry := range entries {
		switch entry.EventKind {
		case point13ValDTimelineKindExportCreated:
			exportAt = entry.CanonicalOccurredAt
		case point13ValDTimelineKindHandoffAcknowledged:
			handoffAckAt = entry.CanonicalOccurredAt
		case point13ValDTimelineKindCustomerAcceptanceRecorded:
			acceptanceAt = entry.CanonicalOccurredAt
		}
	}
	if acceptanceAt == "" {
		return false
	}
	if exportAt != "" {
		comparison, ok := point13ValDTimelineTimestampCompare(acceptanceAt, exportAt)
		if !ok || comparison < 0 {
			return true
		}
	}
	if handoffAckAt != "" {
		comparison, ok := point13ValDTimelineTimestampCompare(acceptanceAt, handoffAckAt)
		if !ok || comparison < 0 {
			return true
		}
	}
	return false
}

func point13ValDAllowedQueryFilterRefs(timeline Point13ValDCustomerAuditorOperationalTimeline, dependency Point13ValDDependencySnapshot) []string {
	return []string{
		timeline.TimelineID,
		dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID,
		dependency.ValC.OperationalHandoffChecklist.HandoffChecklistID,
		dependency.ValC.CustomerAcceptanceTrace.AcceptanceTraceID,
		dependency.ValC.SupportOffboardingHandoffPacket.SupportOffboardingPacketID,
		dependency.ValC.AIEvidenceExportLineageSummary.AIExportSummaryID,
	}
}

func point13ValDQueryFilterRefsWithinScope(values []string, timeline Point13ValDCustomerAuditorOperationalTimeline, dependency Point13ValDDependencySnapshot) bool {
	allowed := point13ValDAllowedQueryFilterRefs(timeline, dependency)
	for _, value := range values {
		if !point13ValDRawStringInSet(allowed, value) {
			return false
		}
	}
	return true
}

func point13ValDTimelineEventSourceMatches(entry Point13ValDTimelineEvent, dependency Point13ValDDependencySnapshot) bool {
	switch entry.EventKind {
	case point13ValDTimelineKindExportCreated:
		return entry.SourceRef == dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID &&
			entry.AuditEventRef == dependency.ValC.CustomerEvidenceExportPackage.AuditEventRef
	case point13ValDTimelineKindRedactionApplied:
		return entry.SourceRef == dependency.ValC.RedactionSafeDisclosureBoundary.RedactionBoundaryID &&
			entry.AuditEventRef == dependency.ValC.RedactionSafeDisclosureBoundary.RedactionAuditEventRef
	case point13ValDTimelineKindHandoffStarted, point13ValDTimelineKindHandoffAcknowledged:
		return entry.SourceRef == dependency.ValC.OperationalHandoffChecklist.HandoffChecklistID &&
			point13ValDRawStringInSet(dependency.ValC.OperationalHandoffChecklist.AuditEventRefs, entry.AuditEventRef)
	case point13ValDTimelineKindCustomerAcceptanceRecorded:
		return entry.SourceRef == dependency.ValC.CustomerAcceptanceTrace.AcceptanceTraceID &&
			point13ValDRawStringInSet(dependency.ValC.CustomerAcceptanceTrace.AuditEventRefs, entry.AuditEventRef)
	case point13ValDTimelineKindSupportOffboardingPrepared:
		return entry.SourceRef == dependency.ValC.SupportOffboardingHandoffPacket.SupportOffboardingPacketID &&
			point13ValDRawStringInSet(dependency.ValC.SupportOffboardingHandoffPacket.AuditEventRefs, entry.AuditEventRef)
	case point13ValDTimelineKindAILineageIncluded:
		return entry.SourceRef == dependency.ValC.AIEvidenceExportLineageSummary.AIExportSummaryID &&
			entry.AuditEventRef == dependency.ValC.AIEvidenceExportLineageSummary.AuditEventRef
	default:
		return false
	}
}

func EvaluatePoint13ValDCustomerAuditorOperationalTimelineState(model Point13ValDCustomerAuditorOperationalTimeline, dependency Point13ValDDependencySnapshot) string {
	if !point13ValDTimelineRefValid(model.TimelineID) ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValCHandoffChecklistRefValid(model.HandoffChecklistRef) ||
		!point13ValCAcceptanceTraceRefValid(model.CustomerAcceptanceRef) ||
		!point13ValCSupportOffboardingPacketRefValid(model.SupportOffboardingRef) ||
		!point13ValDAITimelineProjectionRefValid(model.AITimelineSummaryRef) ||
		!point13ValDTimelineEventListValid(model.TimelineEntries) ||
		model.TimelineHash != point13ValDComputedTimelineHash(model) ||
		!model.TimelineReadOnly ||
		!model.TimelineCannotMutateState ||
		!model.RedactionLimitationsVisible {
		return Point13ValDStateBlocked
	}
	if model.TenantScope != dependency.InheritedTenantScope ||
		model.ExportPackageRef != dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID ||
		model.HandoffChecklistRef != dependency.ValC.OperationalHandoffChecklist.HandoffChecklistID ||
		model.CustomerAcceptanceRef != dependency.ValC.CustomerAcceptanceTrace.AcceptanceTraceID ||
		model.SupportOffboardingRef != dependency.ValC.SupportOffboardingHandoffPacket.SupportOffboardingPacketID ||
		!point12Val0ExactStringSetMatch(point13ValDTimelineEventKindsFromEntries(model.TimelineEntries), point13ValDTimelineEventKinds()) {
		return Point13ValDStateBlocked
	}
	for _, entry := range model.TimelineEntries {
		if !point13Val0RawIdentityValueValid(entry.SourceRef) ||
			!point12Val0AuditRefValid(entry.AuditEventRef) ||
			!point12Val0RawTimestampValid(entry.CanonicalOccurredAt) ||
			!point13ValDCanonicalTimeSourceValid(entry.TimeSource) ||
			(entry.ClientReportedAt != "" && !point12Val0RawTimestampValid(entry.ClientReportedAt)) ||
			!point13ValDSourceMetadataRefValid(entry.SourceMetadataRef) ||
			!entry.ReadOnly {
			return Point13ValDStateBlocked
		}
		if !point13ValDTimelineEventSourceMatches(entry, dependency) {
			return Point13ValDStateBlocked
		}
	}
	if !point13ValDTimelineEventsOrdered(model.TimelineEntries) || point13ValDTimelineAcceptanceBackdated(model.TimelineEntries) {
		return Point13ValDStateReviewRequired
	}
	return Point13ValDStateActive
}

func EvaluatePoint13ValDHandoffTraceQueryProjectionState(model Point13ValDHandoffTraceQueryProjection, dependency Point13ValDDependencySnapshot, timeline Point13ValDCustomerAuditorOperationalTimeline, timelineState string) string {
	if !point13ValDQueryProjectionRefValid(model.QueryProjectionID) ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point13ValDTimelineRefValid(model.TimelineRef) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValCHandoffChecklistRefValid(model.HandoffChecklistRef) ||
		!point13ValDQueryFilterRefsValid(model.FilterRefs) ||
		!point13ValDQueryKindValid(model.QueryKind) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		model.ProjectionHash != point13ValDComputedQueryHash(model) ||
		!model.ReadOnly ||
		model.MutationRequested ||
		model.WriteRequested {
		return Point13ValDStateBlocked
	}
	if timelineState != Point13ValDStateActive ||
		model.TenantScope != dependency.InheritedTenantScope ||
		model.TimelineRef != timeline.TimelineID ||
		model.ExportPackageRef != dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID ||
		model.HandoffChecklistRef != dependency.ValC.OperationalHandoffChecklist.HandoffChecklistID ||
		point13ValAContainsCrossTenantArtifact(model.FilterRefs) ||
		!point13ValDQueryFilterRefsWithinScope(model.FilterRefs, timeline, dependency) {
		return Point13ValDStateBlocked
	}
	return Point13ValDStateActive
}

func EvaluatePoint13ValDExportPackageReadProjectionState(model Point13ValDExportPackageReadProjection, dependency Point13ValDDependencySnapshot) string {
	if !point13ValDExportReadProjectionRefValid(model.ReadProjectionID) ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point13ValCExportPackageRefValid(model.ExportPackageRef) ||
		!point13ValCRedactionBoundaryRefValid(model.RedactionBoundaryRef) ||
		!point13ValBEvidenceRefsValid(model.ExportedEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.ExportedEvidenceHashes) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.ExportedEvidenceRefs, model.ExportedEvidenceHashes) ||
		!point12Val0HashValid(model.ExportManifestHash) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		model.ProjectionHash != point13ValDComputedExportReadHash(model) ||
		!model.ReadOnly ||
		!model.CannotOverwriteHashes ||
		!model.LimitationsVisible ||
		!point13ValCTextListValid(model.VisibleLimitations) {
		return Point13ValDStateBlocked
	}
	if model.TenantScope != dependency.InheritedTenantScope ||
		model.ExportPackageRef != dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID ||
		model.RedactionBoundaryRef != dependency.ValC.RedactionSafeDisclosureBoundary.RedactionBoundaryID ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceRefs, dependency.ValC.CustomerEvidenceExportPackage.ExportedEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.ExportedEvidenceHashes, dependency.ValC.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs) ||
		model.ExportManifestHash != dependency.ValC.CustomerEvidenceExportPackage.ExportManifestHash ||
		model.RetentionClassRef != dependency.ValC.CustomerEvidenceExportPackage.RetentionClassRef ||
		!point12Val0ExactStringSetMatch(model.VisibleLimitations, point13ValDExpectedVisibleLimitations(dependency.ValC)) {
		return Point13ValDStateBlocked
	}
	return Point13ValDStateActive
}

func EvaluatePoint13ValDCustomerAuditorExplanationProjectionState(model Point13ValDCustomerAuditorExplanationProjection, dependency Point13ValDDependencySnapshot, timeline Point13ValDCustomerAuditorOperationalTimeline, query Point13ValDHandoffTraceQueryProjection, exportRead Point13ValDExportPackageReadProjection, exportReadState string) string {
	if !point13ValDExplanationProjectionRefValid(model.ExplanationProjectionID) ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point13ValDTimelineRefValid(model.TimelineRef) ||
		!point13ValDQueryProjectionRefValid(model.QueryProjectionRef) ||
		!point13ValDExportReadProjectionRefValid(model.ExportReadProjectionRef) ||
		!point13ValDExplanationKindValid(model.ExplanationKind) ||
		strings.TrimSpace(model.ExplanationText) == "" ||
		!point13ValCTextListValid(model.VisibleLimitations) ||
		!point13ValDAuditorAnnotationsValid(model.AuditorAnnotations) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		model.ProjectionHash != point13ValDComputedExplanationHash(model) ||
		!model.AdvisoryOnly ||
		!model.ExplanationCannotStrengthenClaims ||
		!model.ExplanationCannotApproveProduction ||
		!model.ExplanationCannotCreatePass {
		return Point13ValDStateBlocked
	}
	if exportReadState != Point13ValDStateActive ||
		model.TenantScope != dependency.InheritedTenantScope ||
		model.TimelineRef != timeline.TimelineID ||
		model.QueryProjectionRef != query.QueryProjectionID ||
		model.ExportReadProjectionRef != exportRead.ReadProjectionID ||
		!point12Val0ExactStringSetMatch(model.VisibleLimitations, exportRead.VisibleLimitations) {
		return Point13ValDStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(model.ExplanationText, strings.Join(model.VisibleLimitations, " ")) {
		return Point13ValDStateBlocked
	}
	for _, annotation := range model.AuditorAnnotations {
		if !annotation.AnnotationOnly || annotation.ApprovesProduction || annotation.ChangesExportState || point13Val0ContainsForbiddenClaim(annotation.Text) {
			return Point13ValDStateBlocked
		}
	}
	return Point13ValDStateActive
}

func EvaluatePoint13ValDTimelineAccessBoundaryState(model Point13ValDTimelineAccessBoundary, dependency Point13ValDDependencySnapshot, timeline Point13ValDCustomerAuditorOperationalTimeline, query Point13ValDHandoffTraceQueryProjection, exportRead Point13ValDExportPackageReadProjection) string {
	if !point13ValDTimelineAccessBoundaryRefValid(model.AccessBoundaryID) ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point13ValDAudienceTypeValid(model.AudienceType) ||
		!point13ValAOwnerRefValid(model.AudienceRef) ||
		!point13ValAOwnerRefValid(model.CustomerOwnerRef) ||
		!point13ValAOwnerRefValid(model.AuditorOwnerRef) ||
		!point13ValDTimelineRefValid(model.TimelineRef) ||
		!point13ValDQueryProjectionRefValid(model.QueryProjectionRef) ||
		!point13ValDExportReadProjectionRefValid(model.ExportReadProjectionRef) ||
		!point12Val0AuditRefValid(model.AccessAuditEventRef) ||
		!model.ReadOnly ||
		model.MutationRequested ||
		model.CrossTenantAccessRequested {
		return Point13ValDStateBlocked
	}
	if model.TenantScope != dependency.InheritedTenantScope ||
		model.CustomerOwnerRef != dependency.ValC.CustomerEvidenceExportPackage.CustomerOwnerRef ||
		model.TimelineRef != timeline.TimelineID ||
		model.QueryProjectionRef != query.QueryProjectionID ||
		model.ExportReadProjectionRef != exportRead.ReadProjectionID {
		return Point13ValDStateBlocked
	}
	switch model.AudienceType {
	case point13ValDAudienceCustomer:
		if model.AudienceRef != model.CustomerOwnerRef {
			return Point13ValDStateBlocked
		}
	case point13ValDAudienceAuditor:
		if model.AudienceRef != model.AuditorOwnerRef {
			return Point13ValDStateBlocked
		}
	default:
		return Point13ValDStateBlocked
	}
	return Point13ValDStateActive
}

func EvaluatePoint13ValDAITimelineLineageProjectionState(model Point13ValDAITimelineLineageProjection, dependency Point13ValDDependencySnapshot, timeline Point13ValDCustomerAuditorOperationalTimeline, timelineState string) string {
	if point13ValDRawStringInSet(point12Val0BlockedAIEvidenceCandidateTypes(), model.AIOutputType) {
		return Point13ValDStateBlocked
	}
	ai := dependency.ValC.AIEvidenceExportLineageSummary
	if !point13ValDAITimelineProjectionRefValid(model.AIProjectionID) ||
		!point13Val0RawScopeValid(model.TenantScope) ||
		!point13ValDTimelineRefValid(model.TimelineRef) ||
		!point13ValCAIExportSummaryRefValid(model.AIExportSummaryRef) ||
		!point12Val0AIEvidenceCandidateTypeValid(model.AIOutputType) ||
		!point13ValAAIEvidenceCandidateRefValid(model.EvidenceCandidateRef) ||
		!point13ValBEvidenceRefsValid(model.InputEvidenceRefs) ||
		!point13ValBEvidenceHashRefsValid(model.InputEvidenceHashRefs) ||
		!point13ValBEvidenceHashRefsMatchEvidenceRefs(model.InputEvidenceRefs, model.InputEvidenceHashRefs) ||
		!point12Val0VersionRefValid(model.ModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(model.PermissionManifestHash) ||
		!point12Val0AuditRefValid(model.AuditEventRef) ||
		!model.ReadOnly ||
		!model.AdvisoryOnly ||
		!model.EvidenceCandidateOnly ||
		model.PassAllowed ||
		model.ApprovalGranted ||
		model.DeploymentAuthorized ||
		model.ProductionReadinessClaimed ||
		model.ProductionMutationAllowed ||
		model.CanonicalMutationAllowed ||
		model.CanStrengthenTimelineClaim ||
		model.CanSatisfyAcceptanceByItself {
		return Point13ValDStateBlocked
	}
	if timelineState != Point13ValDStateActive ||
		model.TenantScope != dependency.InheritedTenantScope ||
		model.TimelineRef != timeline.TimelineID ||
		model.AIExportSummaryRef != ai.AIExportSummaryID ||
		model.EvidenceCandidateRef != ai.EvidenceCandidateRef ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceRefs, ai.InputEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.InputEvidenceHashRefs, ai.InputEvidenceHashRefs) ||
		model.ModelOrRuleVersionRef != ai.ModelOrRuleVersionRef ||
		model.PermissionManifestHash != ai.PermissionManifestHash ||
		model.AuditEventRef != ai.AuditEventRef ||
		point13ValAContainsCrossTenantArtifact(model.InputEvidenceRefs) {
		return Point13ValDStateBlocked
	}
	if model.ExternalAPIAllowed {
		if !point12Val0GovernanceEventRefValid(model.ExternalAPIGovernanceEventRef) {
			return Point13ValDStateBlocked
		}
		return Point13ValDStateReviewRequired
	}
	return Point13ValDStateActive
}

func EvaluatePoint13ValDNoOverclaimState(model Point13ValDNoOverclaimProjectionWording) string {
	if !point13ValCTextListValid(model.ObservedTimelineTexts) ||
		!point13ValCTextListValid(model.ObservedQueryTexts) ||
		!point13ValCTextListValid(model.ObservedReadProjectionTexts) ||
		!point13ValCTextListValid(model.ObservedExplanationTexts) ||
		!point13ValCTextListValid(model.ObservedSupportOffboardingTexts) ||
		!point13ValCTextListValid(model.InternalDiagnosticTexts) ||
		!point13ValCTextListValid(model.AllowedSafeWording) ||
		!point13ValCTextListValid(model.BlockedWording) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValDStateBlocked
	}
	if !point12Val0ExactStringSetMatch(model.AllowedSafeWording, point13ValDAllowedSafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point13Val0ForbiddenClaims()) ||
		model.ProjectionDisclaimer != point13ValDProjectionDisclaimerBaseline {
		return Point13ValDStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(
		strings.Join(model.ObservedTimelineTexts, " "),
		strings.Join(model.ObservedQueryTexts, " "),
		strings.Join(model.ObservedReadProjectionTexts, " "),
		strings.Join(model.ObservedExplanationTexts, " "),
		strings.Join(model.ObservedSupportOffboardingTexts, " "),
	) {
		return Point13ValDStateBlocked
	}
	if point13Val0ContainsForbiddenClaim(strings.Join(model.InternalDiagnosticTexts, " ")) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point13ValDStateBlocked
	}
	return Point13ValDStateActive
}

func point13ValDComponentStates(model Point13ValDFoundation) []string {
	return []string{
		model.DependencyState,
		model.CustomerAuditorOperationalTimelineState,
		model.HandoffTraceQueryProjectionState,
		model.ExportPackageReadProjectionState,
		model.CustomerAuditorExplanationProjectionState,
		model.TimelineAccessBoundaryState,
		model.AITimelineLineageProjectionState,
		model.NoOverclaimState,
	}
}

func point13ValDBlockingReasons(model Point13ValDFoundation) []string {
	reasons := []string{}
	for _, state := range point13ValDComponentStates(model) {
		if state == Point13ValDStateBlocked {
			reasons = append(reasons, "component_blocked")
			break
		}
	}
	return reasons
}

func EvaluatePoint13ValDState(model Point13ValDFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		return Point13ValDStateBlocked
	}
	hasReview := false
	hasIncomplete := false
	for _, state := range point13ValDComponentStates(model) {
		switch state {
		case Point13ValDStateBlocked:
			return Point13ValDStateBlocked
		case Point13ValDStateReviewRequired:
			hasReview = true
		case Point13ValDStateIncomplete:
			hasIncomplete = true
		case Point13ValDStateActive:
		default:
			return Point13ValDStateBlocked
		}
	}
	if hasReview {
		return Point13ValDStateReviewRequired
	}
	if hasIncomplete {
		return Point13ValDStateIncomplete
	}
	return Point13ValDStateActive
}

func Point13ValDFoundationModel() Point13ValDFoundation {
	dependency := point13ValDDependencySnapshotModel()
	timeline := point13ValDTimelineModel(dependency)
	query := point13ValDQueryProjectionModel(timeline, dependency)
	exportRead := point13ValDExportReadProjectionModel(dependency)
	explanation := point13ValDExplanationProjectionModel(timeline, query, exportRead, dependency)
	access := point13ValDAccessBoundaryModel(timeline, query, exportRead, dependency)
	aiProjection := point13ValDAITimelineProjectionModel(timeline, dependency)
	noOverclaim := point13ValDNoOverclaimModel()
	return Point13ValDFoundation{
		ProjectionDisclaimer:                 point13ValDProjectionDisclaimerBaseline,
		Dependency:                           dependency,
		CustomerAuditorOperationalTimeline:   timeline,
		HandoffTraceQueryProjection:          query,
		ExportPackageReadProjection:          exportRead,
		CustomerAuditorExplanationProjection: explanation,
		TimelineAccessBoundary:               access,
		AITimelineLineageProjection:          aiProjection,
		NoOverclaimProjectionWording:         noOverclaim,
	}
}

func ComputePoint13ValDFoundation(model Point13ValDFoundation) Point13ValDFoundation {
	dependencyState, dependencyReasons := point13ValDDependencyStateAndReasons(model.Dependency)
	model.DependencyState = dependencyState
	timelineState := EvaluatePoint13ValDCustomerAuditorOperationalTimelineState(model.CustomerAuditorOperationalTimeline, model.Dependency)
	model.CustomerAuditorOperationalTimelineState = timelineState
	queryState := EvaluatePoint13ValDHandoffTraceQueryProjectionState(model.HandoffTraceQueryProjection, model.Dependency, model.CustomerAuditorOperationalTimeline, timelineState)
	model.HandoffTraceQueryProjectionState = queryState
	exportReadState := EvaluatePoint13ValDExportPackageReadProjectionState(model.ExportPackageReadProjection, model.Dependency)
	model.ExportPackageReadProjectionState = exportReadState
	explanationState := EvaluatePoint13ValDCustomerAuditorExplanationProjectionState(model.CustomerAuditorExplanationProjection, model.Dependency, model.CustomerAuditorOperationalTimeline, model.HandoffTraceQueryProjection, model.ExportPackageReadProjection, exportReadState)
	model.CustomerAuditorExplanationProjectionState = explanationState
	accessState := EvaluatePoint13ValDTimelineAccessBoundaryState(model.TimelineAccessBoundary, model.Dependency, model.CustomerAuditorOperationalTimeline, model.HandoffTraceQueryProjection, model.ExportPackageReadProjection)
	model.TimelineAccessBoundaryState = accessState
	aiState := EvaluatePoint13ValDAITimelineLineageProjectionState(model.AITimelineLineageProjection, model.Dependency, model.CustomerAuditorOperationalTimeline, timelineState)
	model.AITimelineLineageProjectionState = aiState
	model.NoOverclaimState = EvaluatePoint13ValDNoOverclaimState(model.NoOverclaimProjectionWording)
	model.BlockingReasons = append([]string{}, dependencyReasons...)
	model.BlockingReasons = append(model.BlockingReasons, point13ValDBlockingReasons(model)...)
	model.CurrentState = EvaluatePoint13ValDState(model)
	return model
}
