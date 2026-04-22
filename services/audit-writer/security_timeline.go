package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/compliance"
	"github.com/denisgrosek/changelock/internal/connectors"
	"github.com/denisgrosek/changelock/internal/handoff"
	"github.com/denisgrosek/changelock/internal/httpjson"
	supplychaincore "github.com/denisgrosek/changelock/internal/supplychain"
	vulnerabilitycore "github.com/denisgrosek/changelock/internal/vulnerability"
	"github.com/denisgrosek/changelock/internal/workflow"
)

type securityTimelineResponse struct {
	SchemaVersion     string                  `json:"schema_version"`
	GeneratedAt       time.Time               `json:"generated_at"`
	CountsBySource    map[string]int          `json:"counts_by_source"`
	CountsByLifecycle map[string]int          `json:"counts_by_lifecycle"`
	CountsBySeverity  map[string]int          `json:"counts_by_severity"`
	Entries           []securityTimelineEntry `json:"entries"`
	Limitations       []string                `json:"limitations,omitempty"`
}

type securityTimelineEntry struct {
	SchemaVersion               string    `json:"schema_version"`
	EntryID                     string    `json:"entry_id"`
	Timestamp                   time.Time `json:"timestamp"`
	SubjectRef                  string    `json:"subject_ref"`
	SubjectType                 string    `json:"subject_type"`
	SubjectLabel                string    `json:"subject_label"`
	SourceSubsystem             string    `json:"source_subsystem"`
	LifecyclePhase              string    `json:"lifecycle_phase"`
	EventType                   string    `json:"event_type"`
	Severity                    string    `json:"severity"`
	Importance                  string    `json:"importance"`
	Outcome                     string    `json:"outcome"`
	Title                       string    `json:"title"`
	Summary                     string    `json:"summary"`
	EvidenceRefs                []string  `json:"evidence_refs,omitempty"`
	IncidentRef                 string    `json:"incident_ref,omitempty"`
	RecommendationRef           string    `json:"recommendation_ref,omitempty"`
	NextAction                  string    `json:"next_action,omitempty"`
	DrilldownTab                string    `json:"drilldown_tab,omitempty"`
	DrilldownLabel              string    `json:"drilldown_label,omitempty"`
	DrilldownTargetKind         string    `json:"drilldown_target_kind,omitempty"`
	DrilldownTargetRef          string    `json:"drilldown_target_ref,omitempty"`
	DrilldownTargetSecondaryRef string    `json:"drilldown_target_secondary_ref,omitempty"`
	ResourceURI                 string    `json:"resource_uri,omitempty"`
	PersonaHints                []string  `json:"persona_hints,omitempty"`
	Limitations                 []string  `json:"limitations,omitempty"`
}

type securityTimelineIncidentIndex struct {
	byID       map[string]*investigationIncident
	byScope    map[string]*investigationIncident
	byRepo     map[string]*investigationIncident
	byWorkload map[string]*investigationIncident
}

func (s server) securityTimelineHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	lifecyclePhase := normalizeCommandCenterLifecyclePhase(r.URL.Query().Get("lifecycle_phase"))
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildSecurityTimeline(ctx, filter, lifecyclePhase)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildSecurityTimeline(ctx context.Context, filter audit.EventFilter, lifecyclePhase string) (securityTimelineResponse, error) {
	limit := filter.Limit
	if limit <= 0 {
		limit = 25
	}
	eventFilter := securityTimelineEventFilter(filter, limit)
	contextFilter := securityTimelineContextFilter(filter, limit)

	events, err := s.store.ListEvents(ctx, eventFilter)
	if err != nil {
		return securityTimelineResponse{}, err
	}
	incidents, err := s.listIncidents(ctx, incidentFilter{event: contextFilter})
	if err != nil {
		return securityTimelineResponse{}, err
	}
	recommendations, err := s.listRecommendations(ctx, recommendationFilter{
		event: contextFilter,
		Limit: maxInt(limit, 12),
	})
	if err != nil {
		return securityTimelineResponse{}, err
	}

	incidentIndex := buildSecurityTimelineIncidentIndex(incidents)
	allEntries := make([]securityTimelineEntry, 0, len(events))
	for _, event := range events {
		entry := buildSecurityTimelineEntry(event, incidentIndex, recommendations)
		if strings.TrimSpace(entry.EntryID) == "" {
			continue
		}
		allEntries = append(allEntries, entry)
	}
	countsBySource := map[string]int{}
	countsByLifecycle := map[string]int{}
	countsBySeverity := map[string]int{}
	for _, entry := range allEntries {
		countsBySource[entry.SourceSubsystem]++
		countsByLifecycle[entry.LifecyclePhase]++
		countsBySeverity[entry.Severity]++
	}

	entries := make([]securityTimelineEntry, 0, len(allEntries))
	for _, entry := range allEntries {
		if lifecyclePhase != "" && entry.LifecyclePhase != lifecyclePhase {
			continue
		}
		entries = append(entries, entry)
	}
	sort.Slice(entries, func(i, j int) bool {
		if !entries[i].Timestamp.Equal(entries[j].Timestamp) {
			return entries[i].Timestamp.After(entries[j].Timestamp)
		}
		if entries[i].Severity != entries[j].Severity {
			return entries[i].Severity < entries[j].Severity
		}
		return entries[i].EntryID < entries[j].EntryID
	})
	if len(entries) > limit {
		entries = entries[:limit]
	}

	generatedAt := time.Unix(0, 0).UTC()
	if len(entries) > 0 {
		generatedAt = entries[0].Timestamp.UTC()
	}
	return securityTimelineResponse{
		SchemaVersion:     securityTimelineSchemaVersion,
		GeneratedAt:       generatedAt,
		CountsBySource:    countsBySource,
		CountsByLifecycle: countsByLifecycle,
		CountsBySeverity:  countsBySeverity,
		Entries:           entries,
		Limitations: []string{
			"Unified security timeline aggregates existing evidence-backed audit, incident, recommendation, runtime, validation, handoff, and federation signals; it does not introduce a new canonical truth store.",
			"Lifecycle-phase filtering is a bounded projection over canonical events and never replaces the underlying event stream.",
			"Drill-down routing stays bounded to existing operator surfaces and preserves approval or verification semantics from the source subsystem.",
		},
	}, nil
}

func securityTimelineEventFilter(filter audit.EventFilter, limit int) audit.EventFilter {
	normalized := filter
	normalized.Limit = maxInt(limit*6, 120)
	return normalized
}

func securityTimelineContextFilter(filter audit.EventFilter, limit int) audit.EventFilter {
	return audit.EventFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Since:       filter.Since,
		Until:       filter.Until,
		Limit:       maxInt(limit*8, 160),
	}
}

func buildSecurityTimelineIncidentIndex(incidents []investigationIncident) securityTimelineIncidentIndex {
	index := securityTimelineIncidentIndex{
		byID:       map[string]*investigationIncident{},
		byScope:    map[string]*investigationIncident{},
		byRepo:     map[string]*investigationIncident{},
		byWorkload: map[string]*investigationIncident{},
	}
	for _, incident := range incidents {
		item := incident
		if id := strings.TrimSpace(item.ID); id != "" {
			index.byID[id] = &item
		}
		if scope := strings.TrimSpace(item.ScopeRef); scope != "" {
			index.byScope[scope] = &item
		}
		if repo := strings.TrimSpace(item.Repository); repo != "" {
			index.byRepo[repo] = &item
		}
		for _, workload := range item.AffectedWorkloads {
			if trimmed := strings.TrimSpace(workload); trimmed != "" {
				index.byWorkload[trimmed] = &item
			}
		}
	}
	return index
}

func buildSecurityTimelineEntry(event audit.StoredEvent, incidents securityTimelineIncidentIndex, recommendations []recommendation) securityTimelineEntry {
	source := securityTimelineSource(event)
	incident := matchSecurityTimelineIncident(event, incidents)
	recommendation := matchSecurityTimelineRecommendation(event, incident, source, recommendations)
	handoffRecord := parseSecurityTimelineHandoff(event.Handoff)
	federationEvent := parseSecurityTimelineFederation(event.Federation)
	validationRecord := parseSecurityTimelineValidation(event.ValidationHarness)
	hardeningPayload := parseSecurityTimelineHardening(event.RuntimeIntegrity)
	runtimePayload := parseSecurityTimelineRuntime(event.RuntimeIntegrity)
	intelligencePayload := parsePhase3IntelligencePayload(event.Intelligence)
	enterprisePayload := parsePhase4EnterprisePayload(event.Enterprise)

	subjectRef, subjectType, subjectLabel := securityTimelineSubject(event, incident, handoffRecord, federationEvent, validationRecord, hardeningPayload, intelligencePayload, enterprisePayload)
	title, summary := securityTimelineNarrative(event, incident, source, subjectLabel, handoffRecord, federationEvent, validationRecord, hardeningPayload, runtimePayload, intelligencePayload, enterprisePayload)
	drilldownTab, drilldownLabel := securityTimelineDrilldown(source)
	drilldownKind, drilldownRef, drilldownSecondaryRef, resourceURI := securityTimelineTarget(source, event, incident, recommendation, handoffRecord, federationEvent, validationRecord, hardeningPayload, intelligencePayload, enterprisePayload)
	severity := securityTimelineSeverity(event, incident, validationRecord, source, intelligencePayload, enterprisePayload)

	entry := securityTimelineEntry{
		SchemaVersion:               securityTimelineEntrySchemaVersion,
		EntryID:                     fmt.Sprintf("tle-%d", event.ID),
		Timestamp:                   eventTimestamp(event).UTC(),
		SubjectRef:                  subjectRef,
		SubjectType:                 subjectType,
		SubjectLabel:                subjectLabel,
		SourceSubsystem:             source,
		LifecyclePhase:              commandCenterLifecyclePhase(source),
		EventType:                   strings.TrimSpace(event.EventType),
		Severity:                    severity,
		Importance:                  securityTimelineImportance(severity),
		Outcome:                     securityTimelineOutcome(event, validationRecord),
		Title:                       title,
		Summary:                     summary,
		IncidentRef:                 firstNonEmpty(strings.TrimSpace(event.IncidentID), incidentRef(incident)),
		DrilldownTab:                drilldownTab,
		DrilldownLabel:              drilldownLabel,
		DrilldownTargetKind:         drilldownKind,
		DrilldownTargetRef:          drilldownRef,
		DrilldownTargetSecondaryRef: drilldownSecondaryRef,
		ResourceURI:                 resourceURI,
		PersonaHints:                securityTimelinePersonaHints(source, severity),
		Limitations:                 securityTimelineEventLimitations(source, event, validationRecord, intelligencePayload, enterprisePayload),
	}
	if recommendation != nil {
		entry.RecommendationRef = recommendation.RecommendationID
		entry.NextAction = recommendation.RecommendedAction
	} else {
		entry.NextAction = securityTimelineDefaultNextAction(source)
	}
	entry.EvidenceRefs = securityTimelineEvidenceRefs(event, incident, recommendation, intelligencePayload, enterprisePayload)
	return entry
}

func securityTimelineSource(event audit.StoredEvent) string {
	switch {
	case strings.TrimSpace(event.Component) == phase3IntelligenceComponent:
		return "intelligence"
	case strings.TrimSpace(event.Component) == phase4EnterpriseComponent:
		switch event.EventType {
		case audit.EventTypeEnterprisePartnerTrustRecorded:
			return "partner"
		case audit.EventTypeEnterpriseComplianceMappingRecorded, audit.EventTypeEnterprisePolicyDriftRecorded, audit.EventTypeEnterpriseExecutiveReportRecorded:
			return "governance"
		default:
			return "workflow"
		}
	case strings.HasPrefix(event.EventType, "incident_"):
		return "incident"
	case strings.HasPrefix(event.EventType, "recommendation_"):
		return "recommendation"
	case event.EventType == audit.EventTypeHandoffSealed || event.EventType == audit.EventTypeHandoffCosigned:
		return "handoff"
	case strings.HasPrefix(event.EventType, "federation_"):
		return "federation"
	case event.EventType == audit.EventTypeValidationHarnessRunRecorded:
		return "validation"
	case strings.HasPrefix(event.EventType, "hardening_"):
		return "hardening"
	case strings.HasPrefix(event.EventType, "runtime_") || strings.HasPrefix(event.EventType, "drift_"):
		return "runtime"
	case event.EventType == audit.EventTypeDeployGateDecision || event.EventType == audit.EventTypePolicyDecision || event.EventType == audit.EventTypeArtifactVerificationResult:
		return "deploy"
	default:
		return firstNonEmpty(strings.TrimSpace(event.Component), "events")
	}
}

func matchSecurityTimelineIncident(event audit.StoredEvent, incidents securityTimelineIncidentIndex) *investigationIncident {
	if incidentID := strings.TrimSpace(event.IncidentID); incidentID != "" {
		if incident, ok := incidents.byID[incidentID]; ok {
			return incident
		}
	}
	if workload := strings.TrimSpace(event.Workload); workload != "" {
		if incident, ok := incidents.byScope[workload]; ok {
			return incident
		}
		if incident, ok := incidents.byWorkload[workload]; ok {
			return incident
		}
	}
	if repo := strings.TrimSpace(event.Repo); repo != "" {
		if incident, ok := incidents.byRepo[repo]; ok {
			return incident
		}
	}
	return nil
}

func matchSecurityTimelineRecommendation(event audit.StoredEvent, incident *investigationIncident, source string, recommendations []recommendation) *recommendation {
	recommendationID := strings.TrimSpace(event.RecommendationID)
	for i := range recommendations {
		if recommendationID != "" && recommendations[i].RecommendationID == recommendationID {
			return &recommendations[i]
		}
	}
	if incident != nil {
		for i := range recommendations {
			if containsString(recommendations[i].RelatedIncidentRefs, incident.ID) {
				return &recommendations[i]
			}
		}
	}
	workload := strings.TrimSpace(event.Workload)
	if workload != "" {
		for i := range recommendations {
			if strings.TrimSpace(recommendations[i].SubjectRef) == workload {
				return &recommendations[i]
			}
		}
	}
	if source == "intelligence" || source == "workflow" || source == "partner" || source == "governance" {
		return nil
	}
	repo := strings.TrimSpace(event.Repo)
	if repo != "" {
		for i := range recommendations {
			if strings.TrimSpace(recommendations[i].Repo) == repo {
				return &recommendations[i]
			}
		}
	}
	switch source {
	case "runtime":
		for i := range recommendations {
			if recommendations[i].SourceType == "runtime_signal" {
				return &recommendations[i]
			}
		}
	case "hardening":
		for i := range recommendations {
			if recommendations[i].SourceType == "hardening_signal" {
				return &recommendations[i]
			}
		}
	case "validation":
		for i := range recommendations {
			if recommendations[i].SourceType == "validation_signal" {
				return &recommendations[i]
			}
		}
	case "federation":
		for i := range recommendations {
			if recommendations[i].SourceType == "federation_signal" {
				return &recommendations[i]
			}
		}
	}
	return nil
}

func securityTimelineSubject(event audit.StoredEvent, incident *investigationIncident, handoffRecord *handoffStoredRecord, federationEvent *federationProofEvent, validationRecord *validationHarnessStoredRecord, hardeningPayload *hardeningEventPayload, intelligencePayload phase3IntelligencePayload, enterprisePayload phase4EnterprisePayload) (string, string, string) {
	if intelligencePayload.Vulnerability != nil {
		ref := firstNonEmpty(strings.TrimSpace(intelligencePayload.Vulnerability.SubjectRef), strings.TrimSpace(intelligencePayload.Vulnerability.VerdictID), strings.TrimSpace(intelligencePayload.Vulnerability.VulnerabilityID))
		label := firstNonEmpty(strings.TrimSpace(intelligencePayload.Vulnerability.VulnerabilityID), strings.TrimSpace(intelligencePayload.Vulnerability.PackageName), ref)
		return ref, "vulnerability", label
	}
	if intelligencePayload.SupplyChain != nil {
		ref := firstNonEmpty(strings.TrimSpace(intelligencePayload.SupplyChain.SubjectRef), strings.TrimSpace(intelligencePayload.SupplyChain.PatternID), strings.TrimSpace(intelligencePayload.SupplyChain.PackageName))
		label := firstNonEmpty(strings.TrimSpace(intelligencePayload.SupplyChain.PackageName), ref)
		return ref, "package", label
	}
	if intelligencePayload.Strategic != nil {
		ref := firstNonEmpty(strings.TrimSpace(intelligencePayload.Strategic.SubjectRef), strings.TrimSpace(intelligencePayload.Strategic.AssessmentID))
		return ref, "strategic_assessment", ref
	}
	if intelligencePayload.Query != nil {
		ref := firstNonEmpty(strings.TrimSpace(intelligencePayload.Query.Scope.SubjectRef), strings.TrimSpace(intelligencePayload.Query.QueryID))
		label := firstNonEmpty(strings.TrimSpace(intelligencePayload.Query.Scope.SubjectRef), strings.TrimSpace(intelligencePayload.Query.Query))
		return ref, "grounded_query", label
	}
	if enterprisePayload.Workflow != nil {
		ref := firstNonEmpty(strings.TrimSpace(enterprisePayload.Workflow.SubjectRef), strings.TrimSpace(enterprisePayload.Workflow.WorkflowID))
		label := firstNonEmpty(strings.TrimSpace(enterprisePayload.Workflow.WorkflowID), ref)
		return ref, "workflow", label
	}
	if enterprisePayload.Reconciliation != nil {
		ref := firstNonEmpty(strings.TrimSpace(enterprisePayload.Reconciliation.SubjectRef), strings.TrimSpace(enterprisePayload.Reconciliation.WorkflowID), strings.TrimSpace(enterprisePayload.Reconciliation.ConnectorRef))
		label := firstNonEmpty(strings.TrimSpace(enterprisePayload.Reconciliation.WorkflowID), strings.TrimSpace(enterprisePayload.Reconciliation.ConnectorRef), ref)
		return ref, "workflow_reconciliation", label
	}
	if enterprisePayload.PartnerIntake != nil {
		ref := firstNonEmpty(strings.TrimSpace(enterprisePayload.PartnerIntake.PartnerID), strings.TrimSpace(enterprisePayload.PartnerIntake.HandoffRef))
		label := firstNonEmpty(strings.TrimSpace(enterprisePayload.PartnerIntake.Organization), strings.TrimSpace(enterprisePayload.PartnerIntake.PartnerID), ref)
		return ref, "partner", label
	}
	if enterprisePayload.Compliance != nil {
		ref := firstNonEmpty(strings.TrimSpace(enterprisePayload.Compliance.SubjectRef), strings.TrimSpace(enterprisePayload.Compliance.ControlID))
		label := firstNonEmpty(strings.TrimSpace(enterprisePayload.Compliance.ControlID), strings.TrimSpace(enterprisePayload.Compliance.ControlFamily), ref)
		return ref, "control_mapping", label
	}
	if enterprisePayload.PolicyDrift != nil {
		ref := firstNonEmpty(strings.TrimSpace(enterprisePayload.PolicyDrift.SubjectRef), strings.TrimSpace(enterprisePayload.PolicyDrift.ExceptionID), strings.TrimSpace(enterprisePayload.PolicyDrift.Actor))
		label := firstNonEmpty(strings.TrimSpace(enterprisePayload.PolicyDrift.ExceptionID), strings.TrimSpace(enterprisePayload.PolicyDrift.SubjectRef), ref)
		return ref, "policy_drift", label
	}
	if enterprisePayload.Executive != nil {
		ref := firstNonEmpty(strings.TrimSpace(enterprisePayload.Executive.ScopeRef), "executive_report")
		label := firstNonEmpty(strings.TrimSpace(enterprisePayload.Executive.ScopeRef), "executive report")
		return ref, "executive_report", label
	}
	if handoffRecord != nil && strings.TrimSpace(handoffRecord.PackageID) != "" {
		return handoffRecord.PackageID, "handoff_package", handoffRecord.PackageID
	}
	if federationEvent != nil && strings.TrimSpace(federationEvent.Decision.SubjectRef) != "" {
		ref := strings.TrimSpace(federationEvent.Decision.SubjectRef)
		return ref, "federated_subject", ref
	}
	if validationRecord != nil {
		if strings.TrimSpace(validationRecord.Bundle.Scope) != "" {
			ref := strings.TrimSpace(validationRecord.Bundle.Scope)
			return ref, "validation_scope", ref
		}
		if strings.TrimSpace(validationRecord.Run.RunID) != "" {
			ref := strings.TrimSpace(validationRecord.Run.RunID)
			return ref, "validation_run", ref
		}
	}
	if hardeningPayload != nil && hardeningPayload.Execution != nil && strings.TrimSpace(hardeningPayload.Execution.SubjectRef) != "" {
		ref := strings.TrimSpace(hardeningPayload.Execution.SubjectRef)
		return ref, "workload", ref
	}
	if incident != nil {
		label := firstNonEmpty(strings.TrimSpace(incident.Title), strings.TrimSpace(incident.ScopeRef), incident.ID)
		return incident.ID, "incident", label
	}
	if workload := strings.TrimSpace(event.Workload); workload != "" {
		return workload, "workload", workload
	}
	if repo := strings.TrimSpace(event.Repo); repo != "" {
		return repo, "repo", repo
	}
	if digest := strings.TrimSpace(event.Digest); digest != "" {
		return digest, "artifact", digest
	}
	if requestID := strings.TrimSpace(event.RequestID); requestID != "" {
		return requestID, "request", requestID
	}
	ref := fmt.Sprintf("event-%d", event.ID)
	return ref, "event", ref
}

func securityTimelineNarrative(event audit.StoredEvent, incident *investigationIncident, source string, subjectLabel string, handoffRecord *handoffStoredRecord, federationEvent *federationProofEvent, validationRecord *validationHarnessStoredRecord, hardeningPayload *hardeningEventPayload, runtimePayload *runtimeIntegrityEventPayload, intelligencePayload phase3IntelligencePayload, enterprisePayload phase4EnterprisePayload) (string, string) {
	switch source {
	case "intelligence":
		switch {
		case intelligencePayload.Vulnerability != nil:
			return fmt.Sprintf("Vulnerability relevance updated for %s", subjectLabel), firstNonEmpty(firstString(intelligencePayload.Vulnerability.Explanation.Derived), securityTimelineReasonSummary(event, incident, "Bounded vulnerability relevance was recalculated from reachability and exploitability evidence."))
		case intelligencePayload.SupplyChain != nil:
			return fmt.Sprintf("Supply-chain pattern updated for %s", subjectLabel), firstNonEmpty(firstString(intelligencePayload.SupplyChain.Explanation.Derived), securityTimelineReasonSummary(event, incident, "Bounded supply-chain anomaly and trust-drift evaluation recorded a new pattern verdict."))
		case intelligencePayload.Strategic != nil:
			return fmt.Sprintf("Strategic advisory updated for %s", subjectLabel), firstNonEmpty(firstString(intelligencePayload.Strategic.RecommendedActions), securityTimelineReasonSummary(event, incident, "Strategic advisory assessment recorded a new ranked next action."))
		case intelligencePayload.Query != nil:
			return fmt.Sprintf("Grounded query answered for %s", subjectLabel), firstNonEmpty(firstString(intelligencePayload.Query.RecommendedActions), securityTimelineReasonSummary(event, incident, "A bounded retrieval-grounded security query response was recorded."))
		}
	case "workflow":
		switch {
		case enterprisePayload.Workflow != nil:
			return fmt.Sprintf("Workflow state updated for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Enterprise workflow lifecycle state changed under validation and approval discipline.")
		case enterprisePayload.Reconciliation != nil:
			return fmt.Sprintf("Connector reconciliation updated for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "External workflow projection was reconciled against canonical technical truth.")
		}
	case "partner":
		if enterprisePayload.PartnerIntake != nil {
			return fmt.Sprintf("Partner trust updated for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Partner intake lifecycle recorded a new verification, freshness, or compatibility state.")
		}
	case "governance":
		switch {
		case enterprisePayload.Compliance != nil:
			return fmt.Sprintf("Compliance mapping updated for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Control coverage and evidence-vault posture were recorded for the current subject.")
		case enterprisePayload.PolicyDrift != nil:
			return fmt.Sprintf("Policy drift recorded for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Policy posture changed and the identity-linked governance trail was updated.")
		case enterprisePayload.Executive != nil:
			return fmt.Sprintf("Executive report updated for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Executive governance posture was recalculated from workflow, partner, compliance, and drift artifacts.")
		}
	}
	switch event.EventType {
	case audit.EventTypeDeployGateDecision:
		if event.Decision == audit.DecisionDeny {
			return fmt.Sprintf("Deploy gate denied %s", subjectLabel), securityTimelineReasonSummary(event, incident, "The deploy gate rejected this release under the current trust policy.")
		}
		return fmt.Sprintf("Deploy gate verified %s", subjectLabel), securityTimelineReasonSummary(event, incident, "The deploy gate recorded an allow path for the current subject.")
	case audit.EventTypePolicyDecision:
		return fmt.Sprintf("Policy decision recorded for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Policy evaluation completed for the current control-plane subject.")
	case audit.EventTypeArtifactVerificationResult:
		return fmt.Sprintf("Artifact verification updated %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Artifact trust verification recorded a new signed-evidence result.")
	case audit.EventTypeRuntimeObservationRecorded:
		return fmt.Sprintf("Runtime observation captured for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Runtime observation was correlated with the trusted workload identity.")
	case audit.EventTypeRuntimeSBOMVerificationRecorded:
		if runtimePayload != nil && runtimePayload.SBOMVerification != nil && strings.TrimSpace(runtimePayload.SBOMVerification.Status) != "" {
			return fmt.Sprintf("Runtime SBOM verification %s for %s", strings.ReplaceAll(runtimePayload.SBOMVerification.Status, "_", " "), subjectLabel), securityTimelineReasonSummary(event, incident, "Runtime loaded-state verification completed against the expected SBOM-backed artifact model.")
		}
		return fmt.Sprintf("Runtime SBOM verification updated %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Runtime loaded-state verification completed against the expected SBOM-backed artifact model.")
	case audit.EventTypeRuntimeEnforcementEvaluated:
		return fmt.Sprintf("Runtime enforcement evaluated for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Runtime enforcement policy evaluated the current drift signal and produced a bounded response decision.")
	case audit.EventTypeRuntimeNetworkIsolationApplied:
		return fmt.Sprintf("Network isolation applied to %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Bounded network isolation was applied to reduce blast radius while preserving workflow traceability.")
	case audit.EventTypeRuntimeForensicSnapshotRequested:
		return fmt.Sprintf("Forensic snapshot requested for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Forensic-first runtime handling requested a snapshot before stronger containment or recovery.")
	case audit.EventTypeHardeningActionApplied:
		if hardeningPayload != nil && hardeningPayload.Execution != nil {
			return fmt.Sprintf("Hardening action applied to %s", subjectLabel), firstNonEmpty(strings.TrimSpace(hardeningPayload.Execution.ExecutionResult), securityTimelineReasonSummary(event, incident, "Bounded runtime hardening was applied for the current workload."))
		}
		return fmt.Sprintf("Hardening action applied to %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Bounded runtime hardening was applied for the current workload.")
	case audit.EventTypeHardeningRollbackApplied:
		return fmt.Sprintf("Hardening rollback executed for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Temporary runtime restrictions were rolled back under the bounded hardening policy.")
	case audit.EventTypeHardeningRecoveryCompleted:
		return fmt.Sprintf("Trusted recovery completed for %s", subjectLabel), securityTimelineReasonSummary(event, incident, "Trusted recovery completed after runtime verification and bounded containment review.")
	case audit.EventTypeValidationHarnessRunRecorded:
		if validationRecord != nil && validationRecord.Bundle.Certificate.CertificateID != "" {
			return fmt.Sprintf("Validation run %s", strings.ReplaceAll(validationRecord.Bundle.Certificate.OverallStatus, "_", " ")), fmt.Sprintf("%d strict scenarios completed in %s within %s.", len(validationRecord.Bundle.Executions), validationRecord.Bundle.Mode, validationRecord.Bundle.Scope)
		}
		if validationRecord != nil && validationRecord.Run.RunID != "" {
			return fmt.Sprintf("Validation harness %s", strings.ReplaceAll(validationRecord.Run.OverallStatus, "_", " ")), fmt.Sprintf("%d scenarios were evaluated in %s.", len(validationRecord.Run.Results), validationRecord.Run.Mode)
		}
		return "Validation harness recorded a run", securityTimelineReasonSummary(event, incident, "Controlled validation output was captured for the current scope.")
	case audit.EventTypeHandoffSealed, audit.EventTypeHandoffCosigned:
		if handoffRecord != nil {
			return fmt.Sprintf("Sealed handoff prepared for %s", subjectLabel), fmt.Sprintf("Audience %s · package %s · seal status %s.", handoffRecord.Manifest.Scope.Audience, handoffRecord.PackageID, handoffRecord.Bundle.SealStatus)
		}
		return "Sealed handoff prepared", securityTimelineReasonSummary(event, incident, "A sealed handoff package was generated for the current evidence scope.")
	case audit.EventTypeFederationProofRequested:
		if federationEvent != nil {
			return fmt.Sprintf("Federation proof requested for %s", subjectLabel), fmt.Sprintf("Peer %s requested %s proof exchange.", federationEvent.Request.RequestingPeer, federationEvent.Response.ProofType)
		}
		return "Federation proof requested", securityTimelineReasonSummary(event, incident, "A remote proof exchange request was issued under the bounded federation trust model.")
	case audit.EventTypeFederationProofVerified:
		if federationEvent != nil {
			return fmt.Sprintf("Federation proof %s for %s", strings.ReplaceAll(federationEvent.Decision.Decision, "_", " "), subjectLabel), fmt.Sprintf("Peer %s · %s.", federationEvent.Decision.PeerID, firstNonEmpty(firstString(federationEvent.Decision.Reasons), "remote proof was evaluated against local trust policy"))
		}
		return "Federation proof verified", securityTimelineReasonSummary(event, incident, "Federated proof evaluation completed under the local trust policy.")
	default:
		if strings.HasPrefix(event.EventType, "incident_") {
			return fmt.Sprintf("%s for %s", formatIncidentEventType(event.EventType), subjectLabel), securityTimelineReasonSummary(event, incident, "Incident lifecycle state changed for the current evidence-backed investigation.")
		}
		if strings.HasPrefix(event.EventType, "recommendation_") {
			return fmt.Sprintf("%s for %s", formatIncidentEventType(event.EventType), subjectLabel), securityTimelineReasonSummary(event, incident, "Recommendation workflow state changed for the current evidence-backed remediation path.")
		}
		return fmt.Sprintf("%s updated %s", strings.ReplaceAll(source, "_", " "), subjectLabel), securityTimelineReasonSummary(event, incident, "A new security-relevant control-plane or runtime signal was recorded.")
	}
}

func securityTimelineReasonSummary(event audit.StoredEvent, incident *investigationIncident, fallback string) string {
	if incident != nil && strings.TrimSpace(incident.StatusNarrative) != "" {
		return strings.TrimSpace(incident.StatusNarrative)
	}
	if summary := strings.TrimSpace(event.IncidentSummary); summary != "" {
		return summary
	}
	if resolution := strings.TrimSpace(event.IncidentResolutionSummary); resolution != "" {
		return resolution
	}
	if len(event.Reasons) > 0 {
		return strings.Join(uniqueStrings(event.Reasons), "; ")
	}
	return fallback
}

func securityTimelineDrilldown(source string) (string, string) {
	switch source {
	case "runtime", "hardening":
		return "runtime", "Open Runtime"
	case "validation":
		return "validation", "Open Validation"
	case "federation", "partner":
		return "federation", "Open Federation"
	case "intelligence":
		return "guidance", "Open Guidance"
	case "governance":
		return "analytics", "Open Governance"
	case "handoff", "incident", "recommendation", "deploy":
		return "events", "Open Investigations"
	case "workflow":
		return "exceptions", "Open Workflow"
	case "topology":
		return "topology", "Open Topology"
	case "forensics":
		return "forensics", "Open Forensics"
	default:
		return "events", "Open Investigations"
	}
}

func securityTimelineTarget(source string, event audit.StoredEvent, incident *investigationIncident, recommendation *recommendation, handoffRecord *handoffStoredRecord, federationEvent *federationProofEvent, validationRecord *validationHarnessStoredRecord, hardeningPayload *hardeningEventPayload, intelligencePayload phase3IntelligencePayload, enterprisePayload phase4EnterprisePayload) (string, string, string, string) {
	if recommendation != nil && strings.TrimSpace(recommendation.RecommendationID) != "" {
		return "recommendation", recommendation.RecommendationID, firstNonEmpty(recommendation.RelatedIncidentRefs...), ""
	}
	if incident != nil && strings.TrimSpace(incident.ID) != "" {
		return "incident", incident.ID, "", ""
	}
	if incidentID := strings.TrimSpace(event.IncidentID); incidentID != "" {
		return "incident", incidentID, "", ""
	}
	switch source {
	case "intelligence":
		switch {
		case intelligencePayload.Vulnerability != nil:
			return "vulnerability_relevance", strings.TrimSpace(intelligencePayload.Vulnerability.VerdictID), strings.TrimSpace(intelligencePayload.Vulnerability.VulnerabilityID), fmt.Sprintf("/v1/intelligence/vulnerability-relevance?subject_ref=%s&vulnerability_id=%s", intelligencePayload.Vulnerability.SubjectRef, intelligencePayload.Vulnerability.VulnerabilityID)
		case intelligencePayload.SupplyChain != nil:
			return "supply_chain_pattern", strings.TrimSpace(intelligencePayload.SupplyChain.PatternID), strings.TrimSpace(intelligencePayload.SupplyChain.PackageName), fmt.Sprintf("/v1/intelligence/supply-chain/patterns?subject_ref=%s&package_name=%s", intelligencePayload.SupplyChain.SubjectRef, intelligencePayload.SupplyChain.PackageName)
		case intelligencePayload.Strategic != nil:
			return "strategic_assessment", strings.TrimSpace(intelligencePayload.Strategic.AssessmentID), strings.TrimSpace(intelligencePayload.Strategic.SubjectRef), fmt.Sprintf("/v1/intelligence/strategic/query?subject_ref=%s", intelligencePayload.Strategic.SubjectRef)
		case intelligencePayload.Query != nil:
			return "grounded_query", strings.TrimSpace(intelligencePayload.Query.QueryID), strings.TrimSpace(intelligencePayload.Query.Scope.SubjectRef), fmt.Sprintf("/v1/intelligence/strategic/query?subject_ref=%s&vulnerability_id=%s&package_name=%s", intelligencePayload.Query.Scope.SubjectRef, intelligencePayload.Query.Scope.VulnerabilityID, intelligencePayload.Query.Scope.PackageName)
		}
	case "workflow":
		switch {
		case enterprisePayload.Workflow != nil:
			return "workflow_record", strings.TrimSpace(enterprisePayload.Workflow.WorkflowID), strings.TrimSpace(enterprisePayload.Workflow.SubjectRef), fmt.Sprintf("/v1/enterprise/workflow/lifecycle?subject_ref=%s&workflow_id=%s", enterprisePayload.Workflow.SubjectRef, enterprisePayload.Workflow.WorkflowID)
		case enterprisePayload.Reconciliation != nil:
			return "workflow_reconciliation", strings.TrimSpace(enterprisePayload.Reconciliation.WorkflowID), strings.TrimSpace(enterprisePayload.Reconciliation.ConnectorRef), fmt.Sprintf("/v1/enterprise/workflow/connectors/reconcile?subject_ref=%s&workflow_id=%s", enterprisePayload.Reconciliation.SubjectRef, enterprisePayload.Reconciliation.WorkflowID)
		}
	case "partner":
		if enterprisePayload.PartnerIntake != nil {
			return "partner_trust", strings.TrimSpace(enterprisePayload.PartnerIntake.PartnerID), strings.TrimSpace(enterprisePayload.PartnerIntake.HandoffRef), fmt.Sprintf("/v1/enterprise/partner-trust/dashboard?partner_id=%s", enterprisePayload.PartnerIntake.PartnerID)
		}
	case "governance":
		switch {
		case enterprisePayload.Compliance != nil:
			return "compliance_mapping", strings.TrimSpace(enterprisePayload.Compliance.ControlID), strings.TrimSpace(enterprisePayload.Compliance.SubjectRef), fmt.Sprintf("/v1/enterprise/governance/compliance-mapping?subject_ref=%s", enterprisePayload.Compliance.SubjectRef)
		case enterprisePayload.PolicyDrift != nil:
			return "policy_drift", firstNonEmpty(strings.TrimSpace(enterprisePayload.PolicyDrift.ExceptionID), strings.TrimSpace(enterprisePayload.PolicyDrift.SubjectRef)), strings.TrimSpace(enterprisePayload.PolicyDrift.SubjectRef), fmt.Sprintf("/v1/enterprise/governance/policy-drift?subject_ref=%s", enterprisePayload.PolicyDrift.SubjectRef)
		case enterprisePayload.Executive != nil:
			return "executive_report", firstNonEmpty(strings.TrimSpace(enterprisePayload.Executive.ScopeRef), "executive_report"), "", fmt.Sprintf("/v1/enterprise/governance/executive-report?scope_ref=%s", enterprisePayload.Executive.ScopeRef)
		}
	case "runtime":
		if subjectRef := firstNonEmpty(strings.TrimSpace(event.RecommendationSubjectRef), strings.TrimSpace(event.Workload), strings.TrimSpace(event.Namespace)); subjectRef != "" {
			if findingID := strings.TrimSpace(event.RecommendationSourceRef); findingID != "" {
				return "runtime_finding", findingID, subjectRef, ""
			}
			return "runtime_subject", subjectRef, "", ""
		}
	case "hardening":
		if hardeningPayload != nil && hardeningPayload.Execution != nil && strings.TrimSpace(hardeningPayload.Execution.ExecutionID) != "" {
			return "hardening_execution", strings.TrimSpace(hardeningPayload.Execution.ExecutionID), strings.TrimSpace(hardeningPayload.Execution.SubjectRef), ""
		}
		if hardeningPayload != nil && hardeningPayload.Posture != nil && strings.TrimSpace(hardeningPayload.Posture.SubjectRef) != "" {
			return "runtime_subject", strings.TrimSpace(hardeningPayload.Posture.SubjectRef), "", ""
		}
	case "validation":
		if validationRecord != nil && strings.TrimSpace(validationRecord.Bundle.RunID) != "" {
			return "validation_run", strings.TrimSpace(validationRecord.Bundle.RunID), "", ""
		}
		if validationRecord != nil && strings.TrimSpace(validationRecord.Run.RunID) != "" {
			return "validation_run", strings.TrimSpace(validationRecord.Run.RunID), "", ""
		}
	case "federation":
		if federationEvent != nil && strings.TrimSpace(federationEvent.Request.RequestingPeer) != "" {
			return "federation_peer", strings.TrimSpace(federationEvent.Request.RequestingPeer), "", ""
		}
	case "handoff":
		if handoffRecord != nil && strings.TrimSpace(handoffRecord.PackageID) != "" {
			return "handoff_package", strings.TrimSpace(handoffRecord.PackageID), "", fmt.Sprintf("/v1/handoff/%s", strings.TrimSpace(handoffRecord.PackageID))
		}
	}
	if strings.TrimSpace(event.RecommendationID) != "" {
		return "recommendation", strings.TrimSpace(event.RecommendationID), strings.TrimSpace(event.IncidentID), ""
	}
	return "", "", "", ""
}

func securityTimelineSeverity(event audit.StoredEvent, incident *investigationIncident, validationRecord *validationHarnessStoredRecord, source string, intelligencePayload phase3IntelligencePayload, enterprisePayload phase4EnterprisePayload) string {
	if incident != nil && strings.TrimSpace(incident.Severity) != "" {
		return incident.Severity
	}
	if severity := strings.TrimSpace(event.IncidentSeverity); severity != "" {
		return strings.ToLower(severity)
	}
	if severity := strings.TrimSpace(event.DriftSeverity); severity != "" {
		return strings.ToLower(severity)
	}
	if validationRecord != nil && validationRecord.Bundle.Certificate.OverallStatus == validationStatusFail {
		return "high"
	}
	switch source {
	case "intelligence":
		switch {
		case intelligencePayload.Vulnerability != nil:
			if intelligencePayload.Vulnerability.CurrentState == vulnerabilitycore.RelevanceActivePriority || intelligencePayload.Vulnerability.CurrentState == vulnerabilitycore.RelevanceReachableExternally {
				return "high"
			}
			if intelligencePayload.Vulnerability.CurrentState == vulnerabilitycore.RelevanceReachableLowExploit {
				return "medium"
			}
		case intelligencePayload.SupplyChain != nil:
			if intelligencePayload.SupplyChain.CurrentState == supplychaincore.PatternStateCrossClusterConcern || intelligencePayload.SupplyChain.CurrentState == supplychaincore.PatternStateTyposquat {
				return "high"
			}
			if intelligencePayload.SupplyChain.CurrentState != supplychaincore.PatternStateStableTrusted {
				return "medium"
			}
		case intelligencePayload.Strategic != nil:
			if intelligencePayload.Strategic.Recommendation.PriorityBand == "critical" || intelligencePayload.Strategic.Recommendation.PriorityBand == "high" {
				return "high"
			}
			return "medium"
		case intelligencePayload.Query != nil:
			return "low"
		}
	case "workflow":
		if enterprisePayload.Workflow != nil {
			if enterprisePayload.Workflow.CanonicalState == workflow.StateRejected || enterprisePayload.Workflow.CurrentState == workflow.StateUnderValidation {
				return "high"
			}
			if enterprisePayload.Workflow.ExceptionActive {
				return "medium"
			}
		}
		if enterprisePayload.Reconciliation != nil {
			if enterprisePayload.Reconciliation.CurrentState == connectors.StateExternalClosurePendingValidation || enterprisePayload.Reconciliation.CurrentState == connectors.StateReopenedForValidation {
				return "high"
			}
			if enterprisePayload.Reconciliation.CurrentState == connectors.StateConnectorDegraded || enterprisePayload.Reconciliation.CurrentState == connectors.StateAwaitingExternalReconciliation {
				return "medium"
			}
		}
	case "partner":
		if enterprisePayload.PartnerIntake != nil {
			if enterprisePayload.PartnerIntake.CurrentState == handoff.IntakeStateRejected || enterprisePayload.PartnerIntake.CurrentState == handoff.IntakeStateExpired {
				return "high"
			}
			if enterprisePayload.PartnerIntake.CurrentState != handoff.IntakeStateAccepted {
				return "medium"
			}
		}
	case "governance":
		if enterprisePayload.Compliance != nil {
			if enterprisePayload.Compliance.CoverageState == compliance.CoverageMissing {
				return "high"
			}
			if enterprisePayload.Compliance.CoverageState != compliance.CoverageFull {
				return "medium"
			}
		}
		if enterprisePayload.PolicyDrift != nil {
			if enterprisePayload.PolicyDrift.CurrentState == compliance.DriftStateSoftened {
				return "high"
			}
			return "medium"
		}
		if enterprisePayload.Executive != nil && enterprisePayload.Executive.CurrentState == "executive_governance_attention_required" {
			return "high"
		}
	}
	switch event.Decision {
	case audit.DecisionError:
		return "critical"
	case audit.DecisionDeny:
		return "high"
	default:
		if strings.HasPrefix(event.EventType, "hardening_") || strings.HasPrefix(event.EventType, "federation_") {
			return "medium"
		}
		return "low"
	}
}

func securityTimelineImportance(severity string) string {
	switch strings.ToLower(strings.TrimSpace(severity)) {
	case "critical", "high":
		return "now"
	case "medium":
		return "today"
	default:
		return "watch"
	}
}

func securityTimelineOutcome(event audit.StoredEvent, validationRecord *validationHarnessStoredRecord) string {
	if validationRecord != nil && strings.TrimSpace(validationRecord.Bundle.Certificate.OverallStatus) != "" {
		return validationRecord.Bundle.Certificate.OverallStatus
	}
	switch event.Decision {
	case audit.DecisionAllow:
		return "allow"
	case audit.DecisionDeny:
		return "deny"
	case audit.DecisionError:
		return "error"
	default:
		return "signal"
	}
}

func securityTimelinePersonaHints(source, severity string) []string {
	hints := []string{"security_engineer"}
	switch source {
	case "deploy", "recommendation":
		hints = append(hints, "developer")
	case "runtime", "hardening", "forensics", "topology", "workflow":
		hints = append(hints, "platform_operator")
	case "handoff", "validation", "federation", "partner":
		hints = append(hints, "auditor")
	case "intelligence":
		hints = append(hints, "developer")
	case "governance":
		hints = append(hints, "auditor")
	case "incident":
		hints = append(hints, "platform_operator")
	}
	if severity == "critical" || severity == "high" || source == "incident" {
		hints = append(hints, "executive")
	}
	return uniqueStrings(hints)
}

func securityTimelineDefaultNextAction(source string) string {
	switch source {
	case "runtime", "hardening":
		return "Review runtime findings, containment posture, and recovery gates before widening or clearing action."
	case "validation":
		return "Compare expected versus observed validation behavior and route any fail or flaky path into remediation."
	case "federation":
		return "Re-check peer trust, freshness, and disclosure scope before reusing remote proof."
	case "handoff":
		return "Verify the sealed package, audience, and included evidence scope before sharing it onward."
	case "deploy":
		return "Inspect the deny reason, incident linkage, and recommended remediation before retrying promotion."
	default:
		return "Open the linked operator surface and confirm the evidence-backed next step before acting."
	}
}

func securityTimelineEventLimitations(source string, event audit.StoredEvent, validationRecord *validationHarnessStoredRecord, intelligencePayload phase3IntelligencePayload, enterprisePayload phase4EnterprisePayload) []string {
	limitations := []string{}
	if source == "validation" {
		limitations = append(limitations, "Validation events remain bounded dry-run or simulated control checks and do not become production incident truth.")
	}
	if source == "handoff" {
		limitations = append(limitations, "Sealed handoff events describe package generation and verification state, not blanket assurance over all later downstream handling.")
	}
	if source == "federation" {
		limitations = append(limitations, "Federation events describe local proof evaluation outcomes and never import remote canonical truth wholesale.")
	}
	if source == "intelligence" {
		limitations = append(limitations, "Intelligence events remain advisory and explanation-backed; they do not overwrite canonical runtime or workflow truth.")
	}
	if source == "workflow" || source == "partner" || source == "governance" {
		limitations = append(limitations, "Enterprise events remain workflow and governance projections anchored to canonical audit evidence rather than a separate truth layer.")
	}
	if source == "runtime" && event.EventType == audit.EventTypeRuntimeSBOMVerificationRecorded {
		limitations = append(limitations, "Runtime loaded-state verification is evidence-backed and bounded; unverifiable states remain explicit rather than implied safe.")
	}
	if validationRecord != nil && validationRecord.Bundle.SimulationDerived {
		limitations = append(limitations, "Compatibility-oriented validation output is simulation-derived and must not be interpreted as historical production fact.")
	}
	if intelligencePayload.Query != nil {
		limitations = append(limitations, "Grounded query entries reflect retrieval-bounded answers and preserve uncertainty as first-class output.")
	}
	if enterprisePayload.PartnerIntake != nil {
		limitations = append(limitations, "Partner entries remain redaction-aware and do not expose internal-only sensitive investigation context.")
	}
	return uniqueStrings(limitations)
}

func securityTimelineEvidenceRefs(event audit.StoredEvent, incident *investigationIncident, recommendation *recommendation, intelligencePayload phase3IntelligencePayload, enterprisePayload phase4EnterprisePayload) []string {
	refs := append([]string{}, event.IncidentEvidenceRefs...)
	refs = append(refs, event.RecommendationEvidenceRefs...)
	refs = append(refs, event.IncidentResolutionRefs...)
	if incident != nil {
		refs = append(refs, incident.EvidenceRefs...)
	}
	if recommendation != nil {
		refs = append(refs, recommendation.EvidenceRefs...)
		for _, ref := range recommendation.ReadbackRefs {
			if uri := strings.TrimSpace(ref.ResourceURI); uri != "" {
				refs = append(refs, uri)
			}
		}
	}
	if event.Evidence != nil && event.Evidence.Artifact != nil {
		refs = append(refs,
			strings.TrimSpace(event.Evidence.Artifact.SBOMDigestRef),
			strings.TrimSpace(event.Evidence.Artifact.SBOMArtifactRef),
			strings.TrimSpace(event.Evidence.Artifact.VulnerabilityReportRef),
		)
	}
	if intelligencePayload.Vulnerability != nil {
		refs = append(refs, intelligencePayload.Vulnerability.EvidenceRefs...)
		if intelligencePayload.Vulnerability.VEXCandidate != nil {
			refs = append(refs, intelligencePayload.Vulnerability.VEXCandidate.EvidenceRefs...)
		}
	}
	if intelligencePayload.SupplyChain != nil {
		refs = append(refs, intelligencePayload.SupplyChain.EvidenceRefs...)
	}
	if intelligencePayload.Strategic != nil {
		refs = append(refs, intelligencePayload.Strategic.EvidenceRefs...)
		refs = append(refs, intelligencePayload.Strategic.Recommendation.EvidenceRefs...)
	}
	if intelligencePayload.Query != nil {
		refs = append(refs, intelligencePayload.Query.EvidenceRefs...)
	}
	if enterprisePayload.Workflow != nil {
		refs = append(refs, enterprisePayload.Workflow.EvidenceRefs...)
	}
	if enterprisePayload.Reconciliation != nil {
		refs = append(refs, enterprisePayload.Reconciliation.EvidenceRefs...)
	}
	if enterprisePayload.PartnerIntake != nil {
		refs = append(refs, enterprisePayload.PartnerIntake.EvidenceRefs...)
		refs = append(refs, enterprisePayload.PartnerIntake.Dashboard.PartnerVisibleEvidence...)
	}
	if enterprisePayload.Compliance != nil {
		refs = append(refs, enterprisePayload.Compliance.EvidenceRefs...)
		refs = append(refs, enterprisePayload.Compliance.TechnicalEventRefs...)
		refs = append(refs, enterprisePayload.Compliance.EvidenceVault.EvidenceRefs...)
	}
	if enterprisePayload.PolicyDrift != nil {
		refs = append(refs, enterprisePayload.PolicyDrift.EvidenceRefs...)
	}
	if enterprisePayload.Executive != nil {
		refs = append(refs, enterprisePayload.Executive.EvidenceTraceRefs...)
	}
	return uniqueStrings(refs)
}

func commandCenterLifecyclePhase(source string) string {
	switch source {
	case "deploy":
		return "build_verify"
	case "runtime", "hardening", "forensics", "topology":
		return "runtime"
	case "validation":
		return "validation"
	case "intelligence":
		return "intelligence"
	case "incident", "recommendation", "workflow":
		return "workflow"
	case "handoff", "federation", "partner":
		return "partner"
	case "governance":
		return "governance"
	default:
		return "workflow"
	}
}

func normalizeCommandCenterLifecyclePhase(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "build_verify", "runtime", "validation", "intelligence", "workflow", "partner", "governance":
		return strings.ToLower(strings.TrimSpace(value))
	default:
		return ""
	}
}

func incidentRef(incident *investigationIncident) string {
	if incident == nil {
		return ""
	}
	return strings.TrimSpace(incident.ID)
}

func parseSecurityTimelineHandoff(raw json.RawMessage) *handoffStoredRecord {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var record handoffStoredRecord
	if err := json.Unmarshal(raw, &record); err != nil {
		return nil
	}
	return &record
}

func parseSecurityTimelineFederation(raw json.RawMessage) *federationProofEvent {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var payload federationProofEvent
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil
	}
	return &payload
}

func parseSecurityTimelineValidation(raw json.RawMessage) *validationHarnessStoredRecord {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var payload validationHarnessStoredRecord
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil
	}
	return &payload
}

func parseSecurityTimelineHardening(raw json.RawMessage) *hardeningEventPayload {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var payload hardeningEventPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil
	}
	if payload.Execution == nil && payload.Trigger == nil && payload.Posture == nil {
		return nil
	}
	return &payload
}

func parseSecurityTimelineRuntime(raw json.RawMessage) *runtimeIntegrityEventPayload {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var payload runtimeIntegrityEventPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil
	}
	if payload.Observation == nil && payload.SBOMVerification == nil && strings.TrimSpace(payload.ForensicContext) == "" {
		return nil
	}
	return &payload
}
