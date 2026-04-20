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
	"github.com/denisgrosek/changelock/internal/httpjson"
)

type securityTimelineResponse struct {
	SchemaVersion    string                  `json:"schema_version"`
	GeneratedAt      time.Time               `json:"generated_at"`
	CountsBySource   map[string]int          `json:"counts_by_source"`
	CountsBySeverity map[string]int          `json:"counts_by_severity"`
	Entries          []securityTimelineEntry `json:"entries"`
	Limitations      []string                `json:"limitations,omitempty"`
}

type securityTimelineEntry struct {
	SchemaVersion               string    `json:"schema_version"`
	EntryID                     string    `json:"entry_id"`
	Timestamp                   time.Time `json:"timestamp"`
	SubjectRef                  string    `json:"subject_ref"`
	SubjectType                 string    `json:"subject_type"`
	SubjectLabel                string    `json:"subject_label"`
	SourceSubsystem             string    `json:"source_subsystem"`
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
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildSecurityTimeline(ctx, filter)
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

func (s server) buildSecurityTimeline(ctx context.Context, filter audit.EventFilter) (securityTimelineResponse, error) {
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
	entries := make([]securityTimelineEntry, 0, len(events))
	for _, event := range events {
		entry := buildSecurityTimelineEntry(event, incidentIndex, recommendations)
		if strings.TrimSpace(entry.EntryID) == "" {
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
	countsBySource := map[string]int{}
	countsBySeverity := map[string]int{}
	for _, entry := range entries {
		countsBySource[entry.SourceSubsystem]++
		countsBySeverity[entry.Severity]++
	}

	return securityTimelineResponse{
		SchemaVersion:    securityTimelineSchemaVersion,
		GeneratedAt:      generatedAt,
		CountsBySource:   countsBySource,
		CountsBySeverity: countsBySeverity,
		Entries:          entries,
		Limitations: []string{
			"Unified security timeline aggregates existing evidence-backed audit, incident, recommendation, runtime, validation, handoff, and federation signals; it does not introduce a new canonical truth store.",
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

	subjectRef, subjectType, subjectLabel := securityTimelineSubject(event, incident, handoffRecord, federationEvent, validationRecord, hardeningPayload)
	title, summary := securityTimelineNarrative(event, incident, source, subjectLabel, handoffRecord, federationEvent, validationRecord, hardeningPayload, runtimePayload)
	drilldownTab, drilldownLabel := securityTimelineDrilldown(source)
	drilldownKind, drilldownRef, drilldownSecondaryRef, resourceURI := securityTimelineTarget(source, event, incident, recommendation, handoffRecord, federationEvent, validationRecord, hardeningPayload)
	severity := securityTimelineSeverity(event, incident, validationRecord)

	entry := securityTimelineEntry{
		SchemaVersion:               securityTimelineEntrySchemaVersion,
		EntryID:                     fmt.Sprintf("tle-%d", event.ID),
		Timestamp:                   eventTimestamp(event).UTC(),
		SubjectRef:                  subjectRef,
		SubjectType:                 subjectType,
		SubjectLabel:                subjectLabel,
		SourceSubsystem:             source,
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
		Limitations:                 securityTimelineEventLimitations(source, event, validationRecord),
	}
	if recommendation != nil {
		entry.RecommendationRef = recommendation.RecommendationID
		entry.NextAction = recommendation.RecommendedAction
	} else {
		entry.NextAction = securityTimelineDefaultNextAction(source)
	}
	entry.EvidenceRefs = securityTimelineEvidenceRefs(event, incident, recommendation)
	return entry
}

func securityTimelineSource(event audit.StoredEvent) string {
	switch {
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

func securityTimelineSubject(event audit.StoredEvent, incident *investigationIncident, handoffRecord *handoffStoredRecord, federationEvent *federationProofEvent, validationRecord *validationHarnessStoredRecord, hardeningPayload *hardeningEventPayload) (string, string, string) {
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

func securityTimelineNarrative(event audit.StoredEvent, incident *investigationIncident, source string, subjectLabel string, handoffRecord *handoffStoredRecord, federationEvent *federationProofEvent, validationRecord *validationHarnessStoredRecord, hardeningPayload *hardeningEventPayload, runtimePayload *runtimeIntegrityEventPayload) (string, string) {
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
	case "federation":
		return "federation", "Open Federation"
	case "handoff", "incident", "recommendation", "deploy":
		return "events", "Open Investigations"
	case "topology":
		return "topology", "Open Topology"
	case "forensics":
		return "forensics", "Open Forensics"
	default:
		return "events", "Open Investigations"
	}
}

func securityTimelineTarget(source string, event audit.StoredEvent, incident *investigationIncident, recommendation *recommendation, handoffRecord *handoffStoredRecord, federationEvent *federationProofEvent, validationRecord *validationHarnessStoredRecord, hardeningPayload *hardeningEventPayload) (string, string, string, string) {
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

func securityTimelineSeverity(event audit.StoredEvent, incident *investigationIncident, validationRecord *validationHarnessStoredRecord) string {
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
	case "runtime", "hardening", "forensics", "topology":
		hints = append(hints, "platform_operator")
	case "handoff", "validation", "federation":
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

func securityTimelineEventLimitations(source string, event audit.StoredEvent, validationRecord *validationHarnessStoredRecord) []string {
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
	if source == "runtime" && event.EventType == audit.EventTypeRuntimeSBOMVerificationRecorded {
		limitations = append(limitations, "Runtime loaded-state verification is evidence-backed and bounded; unverifiable states remain explicit rather than implied safe.")
	}
	if validationRecord != nil && validationRecord.Bundle.SimulationDerived {
		limitations = append(limitations, "Compatibility-oriented validation output is simulation-derived and must not be interpreted as historical production fact.")
	}
	return uniqueStrings(limitations)
}

func securityTimelineEvidenceRefs(event audit.StoredEvent, incident *investigationIncident, recommendation *recommendation) []string {
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
	return uniqueStrings(refs)
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
