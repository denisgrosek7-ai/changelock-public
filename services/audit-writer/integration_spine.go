package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	integrationSystemJira       = "jira"
	integrationSystemServiceNow = "servicenow"
)

type integrationIdentityResponse struct {
	SchemaVersion string              `json:"schema_version"`
	CurrentActor  auth.Principal      `json:"current_actor"`
	AuthModel     auth.Description    `json:"auth_model"`
	RoleMapping   map[string][]string `json:"role_mapping,omitempty"`
	Limitations   []string            `json:"limitations,omitempty"`
}

type integrationTicketSystem struct {
	System            string            `json:"system"`
	StatusMapping     map[string]string `json:"status_mapping"`
	PriorityMapping   map[string]string `json:"priority_mapping"`
	RequiredFields    []string          `json:"required_fields"`
	ApprovalAware     bool              `json:"approval_aware"`
	DeeplinkSupported bool              `json:"deeplink_supported"`
}

type integrationTicketCatalogResponse struct {
	SchemaVersion string                    `json:"schema_version"`
	Systems       []integrationTicketSystem `json:"systems"`
	Limitations   []string                  `json:"limitations,omitempty"`
}

type integrationTicketPrepareRequest struct {
	System           string `json:"system"`
	IncidentID       string `json:"incident_id,omitempty"`
	RecommendationID string `json:"recommendation_id,omitempty"`
	Summary          string `json:"summary,omitempty"`
}

type integrationTicketDraftResponse struct {
	SchemaVersion     string            `json:"schema_version"`
	System            string            `json:"system"`
	DraftID           string            `json:"draft_id"`
	Title             string            `json:"title"`
	Description       string            `json:"description"`
	Status            string            `json:"status"`
	Priority          string            `json:"priority"`
	IncidentRef       string            `json:"incident_ref,omitempty"`
	RecommendationRef string            `json:"recommendation_ref,omitempty"`
	ApprovalRequired  bool              `json:"approval_required"`
	EvidenceRefs      []string          `json:"evidence_refs,omitempty"`
	ValidationRefs    []string          `json:"validation_refs,omitempty"`
	HandoffRefs       []string          `json:"handoff_refs,omitempty"`
	DeepLinks         map[string]string `json:"deep_links,omitempty"`
	Payload           map[string]any    `json:"payload"`
	Limitations       []string          `json:"limitations,omitempty"`
}

type integrationSIEMEvent struct {
	EventID           int64     `json:"event_id"`
	EmittedAt         time.Time `json:"emitted_at"`
	EventType         string    `json:"event_type"`
	SourceComponent   string    `json:"source_component"`
	Severity          string    `json:"severity"`
	Decision          string    `json:"decision"`
	CorrelationID     string    `json:"correlation_id,omitempty"`
	SubjectRef        string    `json:"subject_ref,omitempty"`
	SubjectType       string    `json:"subject_type,omitempty"`
	IncidentRef       string    `json:"incident_ref,omitempty"`
	RecommendationRef string    `json:"recommendation_ref,omitempty"`
	EvidenceRefs      []string  `json:"evidence_refs,omitempty"`
}

type integrationSIEMExportResponse struct {
	SchemaVersion string                 `json:"schema_version"`
	ExportedAt    time.Time              `json:"exported_at"`
	Items         []integrationSIEMEvent `json:"items"`
	Limitations   []string               `json:"limitations,omitempty"`
}

type integrationEvidenceExportRequest struct {
	IncidentID       string `json:"incident_id,omitempty"`
	RecommendationID string `json:"recommendation_id,omitempty"`
	PackageID        string `json:"package_id,omitempty"`
	ValidationRunID  string `json:"validation_run_id,omitempty"`
	Audience         string `json:"audience,omitempty"`
}

type integrationEvidenceExportItem struct {
	ItemType     string   `json:"item_type"`
	Reference    string   `json:"reference"`
	URI          string   `json:"uri"`
	Sealed       bool     `json:"sealed"`
	AdvisoryOnly bool     `json:"advisory_only"`
	Description  string   `json:"description"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
}

type integrationEvidenceExportResponse struct {
	SchemaVersion string                          `json:"schema_version"`
	Scope         string                          `json:"scope"`
	Items         []integrationEvidenceExportItem `json:"items"`
	Limitations   []string                        `json:"limitations,omitempty"`
}

func (s server) integrationIdentityHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if _, err := applyPrincipalTenantToRequest(principal, r); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	description := s.authConfig.Describe()
	httpjson.Write(w, http.StatusOK, integrationIdentityResponse{
		SchemaVersion: integrationIdentitySchemaVersion,
		CurrentActor:  principal,
		AuthModel:     description,
		RoleMapping:   description.RoleBindings,
		Limitations: []string{
			"Identity integration introspection describes the configured auth and role-mapping model; it does not prove that the upstream IdP is currently healthy.",
		},
	})
}

func (s server) integrationTicketCatalogHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, integrationTicketCatalogResponse{
		SchemaVersion: integrationTicketCatalogSchemaVersion,
		Systems: []integrationTicketSystem{
			{
				System:            integrationSystemJira,
				StatusMapping:     map[string]string{"active": "To Do", "watching": "In Progress", "resolved": "Done"},
				PriorityMapping:   map[string]string{"critical": "Highest", "high": "High", "medium": "Medium", "low": "Low"},
				RequiredFields:    []string{"summary", "description", "labels", "priority"},
				ApprovalAware:     true,
				DeeplinkSupported: true,
			},
			{
				System:            integrationSystemServiceNow,
				StatusMapping:     map[string]string{"active": "new", "watching": "work_in_progress", "resolved": "closed_complete"},
				PriorityMapping:   map[string]string{"critical": "1", "high": "2", "medium": "3", "low": "4"},
				RequiredFields:    []string{"short_description", "description", "severity", "correlation_id"},
				ApprovalAware:     true,
				DeeplinkSupported: true,
			},
		},
		Limitations: []string{
			"Ticket catalog defines the bounded outbound payload contract only; ChangeLock does not perform direct remote ticket writes in this baseline slice.",
		},
	})
}

func (s server) integrationTicketPrepareHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
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
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request integrationTicketPrepareRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildTicketDraft(ctx, r, request)
	if err != nil {
		writeIntegrationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) integrationSIEMExportHandler(w http.ResponseWriter, r *http.Request) {
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
	items, err := s.store.ListEvents(ctx, securityTimelineEventFilter(filter, maxInt(filter.Limit, 50)))
	if err != nil {
		writeIntegrationError(w, err)
		return
	}
	response := integrationSIEMExportResponse{
		SchemaVersion: integrationSIEMExportSchemaVersion,
		ExportedAt:    time.Unix(0, 0).UTC(),
		Items:         make([]integrationSIEMEvent, 0, len(items)),
		Limitations: []string{
			"SIEM export is a stable normalized event feed over existing audit truth; downstream indexing, alerting, and retention remain the responsibility of the receiving system.",
		},
	}
	for _, event := range items {
		intelligencePayload := parsePhase3IntelligencePayload(event.Intelligence)
		enterprisePayload := parsePhase4EnterprisePayload(event.Enterprise)
		validationRecord := parseSecurityTimelineValidation(event.ValidationHarness)
		response.Items = append(response.Items, integrationSIEMEvent{
			EventID:           event.ID,
			EmittedAt:         eventTimestamp(event).UTC(),
			EventType:         event.EventType,
			SourceComponent:   event.Component,
			Severity:          securityTimelineSeverity(event, nil, validationRecord, securityTimelineSource(event), intelligencePayload, enterprisePayload),
			Decision:          event.Decision,
			CorrelationID:     firstNonEmpty(strings.TrimSpace(event.RequestID), strings.TrimSpace(event.IncidentID), strings.TrimSpace(event.RecommendationID)),
			SubjectRef:        firstNonEmpty(strings.TrimSpace(event.IncidentScopeRef), strings.TrimSpace(event.Workload), strings.TrimSpace(event.RecommendationSubjectRef), strings.TrimSpace(event.Repo)),
			SubjectType:       integrationSIEMSubjectType(event),
			IncidentRef:       strings.TrimSpace(event.IncidentID),
			RecommendationRef: strings.TrimSpace(event.RecommendationID),
			EvidenceRefs:      limitStrings(securityTimelineEvidenceRefs(event, nil, nil, intelligencePayload, enterprisePayload), 8),
		})
	}
	if len(response.Items) > 0 {
		sort.Slice(response.Items, func(i, j int) bool {
			if !response.Items[i].EmittedAt.Equal(response.Items[j].EmittedAt) {
				return response.Items[i].EmittedAt.After(response.Items[j].EmittedAt)
			}
			return response.Items[i].EventID > response.Items[j].EventID
		})
		response.ExportedAt = response.Items[0].EmittedAt
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) integrationEvidenceExportHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request integrationEvidenceExportRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildEvidenceExport(ctx, r, request)
	if err != nil {
		writeIntegrationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildTicketDraft(ctx context.Context, r *http.Request, request integrationTicketPrepareRequest) (integrationTicketDraftResponse, error) {
	system := strings.ToLower(strings.TrimSpace(request.System))
	if system != integrationSystemJira && system != integrationSystemServiceNow {
		return integrationTicketDraftResponse{}, fmt.Errorf("%w: unsupported integration system", audit.ErrInvalidFilter)
	}
	baseFilter, err := parseFilter(r)
	if err != nil {
		return integrationTicketDraftResponse{}, err
	}
	incidentFilter := incidentFilter{event: securityTimelineContextFilter(baseFilter, 50)}
	recommendationFilter := recommendationFilter{event: securityTimelineContextFilter(baseFilter, 50), Limit: 200}

	var incident *investigationIncident
	if id := strings.TrimSpace(request.IncidentID); id != "" {
		value, err := s.getIncidentByID(ctx, id, incidentFilter)
		if err != nil {
			return integrationTicketDraftResponse{}, err
		}
		incident = &value
	}
	var item *recommendation
	if id := strings.TrimSpace(request.RecommendationID); id != "" {
		value, err := s.getRecommendationByID(ctx, id, recommendationFilter)
		if err != nil {
			return integrationTicketDraftResponse{}, err
		}
		item = &value
		if incident == nil && len(value.RelatedIncidentRefs) > 0 {
			if linked, err := s.getIncidentByID(ctx, value.RelatedIncidentRefs[0], incidentFilter); err == nil {
				incident = &linked
			}
		}
	}
	if incident == nil && item == nil {
		return integrationTicketDraftResponse{}, fmt.Errorf("%w: incident_id or recommendation_id is required", audit.ErrInvalidFilter)
	}

	title := firstNonEmpty(strings.TrimSpace(request.Summary), recommendationTitleForTicket(item), incidentTitleForTicket(incident))
	description := buildTicketDescription(incident, item)
	priority := ticketPriority(incident, item)
	status := ticketStatus(incident)
	approvalRequired := item != nil && item.ApprovalMode == recommendationApprovalHumanReview
	evidenceRefs := uniqueStrings(append(ticketEvidenceRefs(incident), ticketRecommendationEvidenceRefs(item)...))
	validationRefs, _ := s.ticketValidationRefs(ctx, baseFilter, incident, item)
	handoffRefs, _ := s.ticketHandoffRefs(ctx, baseFilter, incident, item)
	tenantID := firstNonEmpty(baseFilter.TenantID, incidentTenant(incident), recommendationTenant(item))
	environment := firstNonEmpty(baseFilter.Environment, incidentEnvironment(incident), recommendationEnvironment(item))
	repo := firstNonEmpty(baseFilter.Repo, incidentRepo(incident), recommendationRepo(item))
	deepLinks := map[string]string{}
	if incident != nil {
		deepLinks["incident"] = fmt.Sprintf("/v1/incidents/%s%s", incident.ID, integrationQuerySuffix(map[string]string{
			"tenant_id":   tenantID,
			"environment": environment,
			"repo":        repo,
		}))
		deepLinks["incident_export"] = fmt.Sprintf("/v1/incidents/%s/export%s", incident.ID, integrationQuerySuffix(map[string]string{
			"tenant_id":   tenantID,
			"environment": environment,
			"repo":        repo,
			"audience":    incidentAudienceInternal,
		}))
	}
	if item != nil {
		deepLinks["recommendation"] = fmt.Sprintf("/v1/recommendations/%s%s", item.RecommendationID, integrationQuerySuffix(map[string]string{
			"tenant_id":   tenantID,
			"environment": environment,
			"repo":        repo,
		}))
	}
	if len(validationRefs) > 0 {
		deepLinks["validation"] = validationRefs[0]
	}
	if len(handoffRefs) > 0 {
		deepLinks["handoff"] = handoffRefs[0]
	}

	payload := buildTicketPayload(system, title, description, priority, status, approvalRequired, deepLinks, incident, item)
	return integrationTicketDraftResponse{
		SchemaVersion:     integrationTicketDraftSchemaVersion,
		System:            system,
		DraftID:           recommendationID("integration-ticket", firstNonEmpty(request.IncidentID, request.RecommendationID, title), system),
		Title:             title,
		Description:       description,
		Status:            status,
		Priority:          priority,
		IncidentRef:       incidentIDValue(incident),
		RecommendationRef: recommendationIDValue(item),
		ApprovalRequired:  approvalRequired,
		EvidenceRefs:      limitStrings(evidenceRefs, 12),
		ValidationRefs:    limitStrings(validationRefs, 4),
		HandoffRefs:       limitStrings(handoffRefs, 4),
		DeepLinks:         deepLinks,
		Payload:           payload,
		Limitations: []string{
			"Ticket preparation produces a stable outbound draft only; remote Jira or ServiceNow writes stay outside this baseline slice.",
		},
	}, nil
}

func (s server) buildEvidenceExport(ctx context.Context, r *http.Request, request integrationEvidenceExportRequest) (integrationEvidenceExportResponse, error) {
	baseFilter, err := parseFilter(r)
	if err != nil {
		return integrationEvidenceExportResponse{}, err
	}
	incidentFilter := incidentFilter{event: securityTimelineContextFilter(baseFilter, 50)}
	recommendationFilter := recommendationFilter{event: securityTimelineContextFilter(baseFilter, 50), Limit: 200}
	response := integrationEvidenceExportResponse{
		SchemaVersion: integrationEvidenceSchemaVersion,
		Items:         []integrationEvidenceExportItem{},
		Limitations: []string{
			"Evidence export lists bounded API references and sealed/unsealed semantics; it does not itself freeze downstream evidence state.",
		},
	}
	switch {
	case strings.TrimSpace(request.IncidentID) != "":
		incident, err := s.getIncidentByID(ctx, strings.TrimSpace(request.IncidentID), incidentFilter)
		if err != nil {
			return integrationEvidenceExportResponse{}, err
		}
		response.Scope = "incident:" + incident.ID
		response.Items = append(response.Items,
			integrationEvidenceExportItem{ItemType: "incident", Reference: incident.ID, URI: fmt.Sprintf("/v1/incidents/%s%s", incident.ID, integrationQuerySuffix(map[string]string{"tenant_id": incidentTenant(&incident), "environment": incidentEnvironment(&incident), "repo": incidentRepo(&incident)})), Sealed: false, AdvisoryOnly: false, Description: "Incident detail and lifecycle state.", EvidenceRefs: limitStrings(incident.EvidenceRefs, 8)},
			integrationEvidenceExportItem{ItemType: "incident_export", Reference: incident.ID, URI: fmt.Sprintf("/v1/incidents/%s/export%s", incident.ID, integrationQuerySuffix(map[string]string{"tenant_id": incidentTenant(&incident), "environment": incidentEnvironment(&incident), "repo": incidentRepo(&incident), "audience": firstNonEmpty(strings.TrimSpace(request.Audience), incidentAudienceInternal)})), Sealed: false, AdvisoryOnly: false, Description: "Audience-bounded incident report.", EvidenceRefs: limitStrings(incident.EvidenceRefs, 8)},
			integrationEvidenceExportItem{ItemType: "defense_gap", Reference: incident.ID, URI: fmt.Sprintf("/v1/incidents/%s/defense-gaps%s", incident.ID, integrationQuerySuffix(map[string]string{"tenant_id": incidentTenant(&incident), "environment": incidentEnvironment(&incident), "repo": incidentRepo(&incident)})), Sealed: false, AdvisoryOnly: true, Description: "Defense gap readback surface.", EvidenceRefs: limitStrings(incident.EvidenceRefs, 8)},
			integrationEvidenceExportItem{ItemType: "policy_replay", Reference: incident.ID, URI: fmt.Sprintf("/v1/incidents/%s/policy-replay%s", incident.ID, integrationQuerySuffix(map[string]string{"tenant_id": incidentTenant(&incident), "environment": incidentEnvironment(&incident), "repo": incidentRepo(&incident)})), Sealed: false, AdvisoryOnly: true, Description: "Policy replay and coverage context.", EvidenceRefs: limitStrings(incident.EvidenceRefs, 8)},
			integrationEvidenceExportItem{ItemType: "forensic_state", Reference: incident.ID, URI: fmt.Sprintf("/v1/incidents/%s/forensic-state%s", incident.ID, integrationQuerySuffix(map[string]string{"tenant_id": incidentTenant(&incident), "environment": incidentEnvironment(&incident), "repo": incidentRepo(&incident)})), Sealed: false, AdvisoryOnly: true, Description: "Historical forensic reconstruction anchor.", EvidenceRefs: limitStrings(incident.EvidenceRefs, 8)},
		)
	case strings.TrimSpace(request.RecommendationID) != "":
		item, err := s.getRecommendationByID(ctx, strings.TrimSpace(request.RecommendationID), recommendationFilter)
		if err != nil {
			return integrationEvidenceExportResponse{}, err
		}
		queryTenant := firstNonEmpty(baseFilter.TenantID, recommendationTenant(&item))
		queryEnvironment := firstNonEmpty(baseFilter.Environment, recommendationEnvironment(&item))
		queryRepo := firstNonEmpty(baseFilter.Repo, recommendationRepo(&item))
		response.Scope = "recommendation:" + item.RecommendationID
		response.Items = append(response.Items,
			integrationEvidenceExportItem{ItemType: "recommendation", Reference: item.RecommendationID, URI: fmt.Sprintf("/v1/recommendations/%s%s", item.RecommendationID, integrationQuerySuffix(map[string]string{"tenant_id": queryTenant, "environment": queryEnvironment, "repo": queryRepo})), Sealed: false, AdvisoryOnly: item.AdvisoryOnly, Description: "Recommendation detail and execution history.", EvidenceRefs: limitStrings(item.EvidenceRefs, 8)},
		)
		for _, ref := range item.ReadbackRefs {
			response.Items = append(response.Items, integrationEvidenceExportItem{
				ItemType:     "readback",
				Reference:    ref.ResourceID,
				URI:          ref.ResourceURI,
				Sealed:       false,
				AdvisoryOnly: true,
				Description:  "Advisory readback lineage reference.",
				EvidenceRefs: limitStrings([]string{ref.EvidenceHash}, 1),
			})
		}
		if incidentID := firstNonEmpty(item.RelatedIncidentRefs...); incidentID != "" {
			response.Items = append(response.Items, integrationEvidenceExportItem{
				ItemType:     "linked_incident",
				Reference:    incidentID,
				URI:          fmt.Sprintf("/v1/incidents/%s%s", incidentID, integrationQuerySuffix(map[string]string{"tenant_id": queryTenant, "environment": queryEnvironment, "repo": queryRepo})),
				Sealed:       false,
				AdvisoryOnly: false,
				Description:  "Canonical incident linked from recommendation workflow.",
				EvidenceRefs: limitStrings(item.EvidenceRefs, 8),
			})
		}
	case strings.TrimSpace(request.PackageID) != "":
		record, err := s.getStoredHandoffRecord(ctx, strings.TrimSpace(request.PackageID))
		if err != nil {
			return integrationEvidenceExportResponse{}, err
		}
		response.Scope = "handoff:" + record.PackageID
		response.Items = append(response.Items,
			integrationEvidenceExportItem{ItemType: "handoff", Reference: record.PackageID, URI: fmt.Sprintf("/v1/handoff/%s", record.PackageID), Sealed: true, AdvisoryOnly: false, Description: "Stored sealed handoff package metadata.", EvidenceRefs: limitStrings(record.Manifest.EvidenceRefs, 8)},
			integrationEvidenceExportItem{ItemType: "handoff_manifest", Reference: record.PackageID, URI: fmt.Sprintf("/v1/handoff/%s/manifest", record.PackageID), Sealed: true, AdvisoryOnly: false, Description: "Canonical sealed manifest.", EvidenceRefs: limitStrings(record.Manifest.EvidenceRefs, 8)},
			integrationEvidenceExportItem{ItemType: "handoff_verification", Reference: record.PackageID, URI: fmt.Sprintf("/v1/handoff/%s/verification", record.PackageID), Sealed: true, AdvisoryOnly: false, Description: "Verification and offline check surface.", EvidenceRefs: limitStrings(record.Manifest.EvidenceRefs, 8)},
		)
	case strings.TrimSpace(request.ValidationRunID) != "":
		runs, _, err := s.listStrictValidationRuns(ctx, validationHarnessFilter{
			event:       securityTimelineContextFilter(baseFilter, 50),
			ClusterID:   baseFilter.ClusterID,
			TenantID:    baseFilter.TenantID,
			Environment: baseFilter.Environment,
			Repo:        baseFilter.Repo,
			Limit:       50,
		})
		if err != nil {
			return integrationEvidenceExportResponse{}, err
		}
		for _, run := range runs {
			if run.RunID != strings.TrimSpace(request.ValidationRunID) {
				continue
			}
			response.Scope = "validation:" + run.RunID
			response.Items = append(response.Items,
				integrationEvidenceExportItem{ItemType: "validation_certificate", Reference: run.Certificate.CertificateID, URI: fmt.Sprintf("/v1/validation/certificates/%s", run.Certificate.CertificateID), Sealed: false, AdvisoryOnly: run.SimulationDerived, Description: "Validation certificate with scenario verdict matrix.", EvidenceRefs: limitStrings(run.Certificate.EvidenceRefs, 8)},
			)
			for _, verdict := range run.Verdicts {
				response.Items = append(response.Items, integrationEvidenceExportItem{
					ItemType:     "validation_verdict",
					Reference:    verdict.VerdictID,
					URI:          fmt.Sprintf("/v1/validation/verdicts/%s", verdict.VerdictID),
					Sealed:       false,
					AdvisoryOnly: run.SimulationDerived,
					Description:  "Per-scenario validation verdict.",
					EvidenceRefs: limitStrings(verdict.EvidenceRefs, 6),
				})
			}
			sort.Slice(response.Items, func(i, j int) bool { return response.Items[i].Reference < response.Items[j].Reference })
			return response, nil
		}
		return integrationEvidenceExportResponse{}, errIncidentNotFound
	default:
		return integrationEvidenceExportResponse{}, fmt.Errorf("%w: one export scope is required", audit.ErrInvalidFilter)
	}
	sort.Slice(response.Items, func(i, j int) bool {
		if response.Items[i].ItemType != response.Items[j].ItemType {
			return response.Items[i].ItemType < response.Items[j].ItemType
		}
		return response.Items[i].Reference < response.Items[j].Reference
	})
	return response, nil
}

func buildTicketDescription(incident *investigationIncident, item *recommendation) string {
	parts := []string{}
	if incident != nil {
		parts = append(parts,
			fmt.Sprintf("Incident %s: %s", incident.ID, firstNonEmpty(incident.CaseSummary, incident.Summary, incident.LikelyCause)),
			fmt.Sprintf("Severity %s · state %s · scope %s", incident.Severity, incident.State, incident.ScopeRef),
		)
	}
	if item != nil {
		parts = append(parts,
			fmt.Sprintf("Recommendation %s: %s", item.RecommendationID, item.RecommendedAction),
			fmt.Sprintf("Verification path: %s", strings.Join(item.VerificationPlan, " | ")),
		)
	}
	return strings.Join(parts, "\n\n")
}

func buildTicketPayload(system, title, description, priority, status string, approvalRequired bool, deeplinks map[string]string, incident *investigationIncident, item *recommendation) map[string]any {
	correlationID := firstNonEmpty(incidentIDValue(incident), recommendationIDValue(item), title)
	switch system {
	case integrationSystemServiceNow:
		return map[string]any{
			"short_description": title,
			"description":       description,
			"severity":          priority,
			"state":             status,
			"correlation_id":    correlationID,
			"approval_required": approvalRequired,
			"deep_links":        deeplinks,
		}
	default:
		return map[string]any{
			"fields": map[string]any{
				"summary":        title,
				"description":    description,
				"priority":       map[string]string{"name": priority},
				"labels":         []string{"changelock", strings.ToLower(strings.TrimSpace(priority))},
				"status":         status,
				"correlation_id": correlationID,
			},
			"approval_required": approvalRequired,
			"deep_links":        deeplinks,
		}
	}
}

func ticketPriority(incident *investigationIncident, item *recommendation) string {
	if incident != nil {
		return strings.ToLower(strings.TrimSpace(incident.Severity))
	}
	if item != nil {
		switch recommendationPriorityBand(item.PriorityBand) {
		case "NOW":
			return "high"
		case "TODAY":
			return "medium"
		default:
			return "low"
		}
	}
	return "low"
}

func ticketStatus(incident *investigationIncident) string {
	if incident == nil {
		return "active"
	}
	return strings.ToLower(strings.TrimSpace(incident.State))
}

func ticketEvidenceRefs(incident *investigationIncident) []string {
	if incident == nil {
		return nil
	}
	return incident.EvidenceRefs
}

func ticketRecommendationEvidenceRefs(item *recommendation) []string {
	if item == nil {
		return nil
	}
	return item.EvidenceRefs
}

func incidentTitleForTicket(incident *investigationIncident) string {
	if incident == nil {
		return ""
	}
	return fmt.Sprintf("[%s] %s", strings.ToUpper(firstNonEmpty(incident.Severity, incident.Priority, "WATCH")), incident.Title)
}

func recommendationTitleForTicket(item *recommendation) string {
	if item == nil {
		return ""
	}
	return fmt.Sprintf("[%s] %s", recommendationPriorityBand(item.PriorityBand), item.Title)
}

func recommendationIDValue(item *recommendation) string {
	if item == nil {
		return ""
	}
	return item.RecommendationID
}

func incidentIDValue(incident *investigationIncident) string {
	if incident == nil {
		return ""
	}
	return incident.ID
}

func (s server) ticketValidationRefs(ctx context.Context, baseFilter audit.EventFilter, incident *investigationIncident, item *recommendation) ([]string, error) {
	filter := validationHarnessFilter{
		event: audit.EventFilter{
			TenantID:    firstNonEmpty(baseFilter.TenantID, incidentTenant(incident), recommendationTenant(item)),
			Environment: firstNonEmpty(baseFilter.Environment, incidentEnvironment(incident), recommendationEnvironment(item)),
			Repo:        strings.TrimSpace(baseFilter.Repo),
			Limit:       50,
		},
		TenantID:    firstNonEmpty(baseFilter.TenantID, incidentTenant(incident), recommendationTenant(item)),
		Environment: firstNonEmpty(baseFilter.Environment, incidentEnvironment(incident), recommendationEnvironment(item)),
		Repo:        strings.TrimSpace(baseFilter.Repo),
		Limit:       6,
	}
	runs, _, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return nil, err
	}
	refs := []string{}
	for _, run := range runs {
		refs = append(refs, fmt.Sprintf("/v1/validation/certificates/%s", run.Certificate.CertificateID))
		if len(refs) >= 3 {
			break
		}
	}
	return uniqueStrings(refs), nil
}

func (s server) ticketHandoffRefs(ctx context.Context, baseFilter audit.EventFilter, incident *investigationIncident, item *recommendation) ([]string, error) {
	if incident == nil && item == nil {
		return nil, nil
	}
	targetIncident := firstNonEmpty(incidentIDValue(incident), firstNonEmpty(item.RelatedIncidentRefs...))
	if targetIncident == "" {
		return nil, nil
	}
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		ClusterID:   baseFilter.ClusterID,
		TenantID:    firstNonEmpty(baseFilter.TenantID, incidentTenant(incident), recommendationTenant(item)),
		Environment: firstNonEmpty(baseFilter.Environment, incidentEnvironment(incident), recommendationEnvironment(item)),
		Repo:        strings.TrimSpace(baseFilter.Repo),
		Component:   handoffComponent,
		Limit:       100,
	})
	if err != nil {
		return nil, err
	}
	refs := []string{}
	for _, event := range events {
		record := parseSecurityTimelineHandoff(event.Handoff)
		if record == nil {
			continue
		}
		if containsString(record.Manifest.Scope.IncidentRefs, targetIncident) {
			refs = append(refs, fmt.Sprintf("/v1/handoff/%s", record.PackageID))
		}
	}
	return uniqueStrings(refs), nil
}

func integrationSIEMSubjectType(event audit.StoredEvent) string {
	switch {
	case strings.TrimSpace(event.IncidentScopeRef) != "":
		return "incident_scope"
	case strings.TrimSpace(event.RecommendationSubjectRef) != "":
		return strings.TrimSpace(event.RecommendationSubjectType)
	case strings.TrimSpace(event.Workload) != "":
		return "workload"
	case strings.TrimSpace(event.Repo) != "":
		return "repo"
	default:
		return "unknown"
	}
}

func integrationQuerySuffix(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	values := url.Values{}
	for key, value := range params {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			values.Set(key, trimmed)
		}
	}
	if len(values) == 0 {
		return ""
	}
	return "?" + values.Encode()
}

func incidentTenant(incident *investigationIncident) string {
	if incident == nil {
		return ""
	}
	return strings.TrimSpace(incident.TenantID)
}

func incidentEnvironment(incident *investigationIncident) string {
	if incident == nil {
		return ""
	}
	return strings.TrimSpace(incident.Environment)
}

func incidentRepo(incident *investigationIncident) string {
	if incident == nil {
		return ""
	}
	return strings.TrimSpace(incident.Repository)
}

func incidentService(incident *investigationIncident) string {
	if incident == nil {
		return ""
	}
	return firstNonEmpty(incident.ScopeRef, firstNonEmpty(incident.AffectedWorkloads...))
}

func recommendationTenant(item *recommendation) string {
	if item == nil {
		return ""
	}
	return strings.TrimSpace(item.Team)
}

func recommendationEnvironment(item *recommendation) string {
	if item == nil {
		return ""
	}
	return strings.TrimSpace(item.Environment)
}

func recommendationRepo(item *recommendation) string {
	if item == nil {
		return ""
	}
	return strings.TrimSpace(item.Repo)
}

func recommendationService(item *recommendation) string {
	if item == nil {
		return ""
	}
	return firstNonEmpty(item.Service, item.SubjectRef)
}

func writeIntegrationError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, audit.ErrInvalidFilter):
		status = http.StatusBadRequest
	case errors.Is(err, errIncidentNotFound), errors.Is(err, errHandoffNotFound):
		status = http.StatusNotFound
	case errors.Is(err, context.DeadlineExceeded):
		status = http.StatusGatewayTimeout
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}
