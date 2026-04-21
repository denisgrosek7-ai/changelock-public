package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	identityFabricSchemaVersion        = "4a.identity_fabric.v1"
	itsmLifecycleSchemaVersion         = "4a.itsm_lifecycle.v1"
	itsmLifecycleFlowsSchemaVersion    = "4a.itsm_lifecycle_flows.v1"
	siemSyncSchemaVersion              = "4a.siem_sync.v1"
	siemSyncEvaluationSchemaVersion    = "4a.siem_sync_evaluation.v1"
	incidentCollaborationSchemaVersion = "4a.incident_collaboration.v1"
	integrationSafetySchemaVersion     = "4a.integration_safety.v1"
	integrationSafetyHealthSchema      = "4a.integration_safety_health.v1"
)

type identityFabricRoleMapping struct {
	BusinessRole   string   `json:"business_role"`
	ChangelockRole string   `json:"changelock_role"`
	BindingValues  []string `json:"binding_values,omitempty"`
	Subjects       []string `json:"subjects,omitempty"`
}

type identityFabricApproverClass struct {
	ClassID             string   `json:"class_id"`
	DisplayName         string   `json:"display_name"`
	AllowedRoles        []string `json:"allowed_roles"`
	DelegationSemantics string   `json:"delegation_semantics"`
	Capabilities        []string `json:"capabilities,omitempty"`
}

type identityFabricResponse struct {
	SchemaVersion               string                        `json:"schema_version"`
	CurrentActor                authInfoResponse              `json:"current_actor"`
	AuthModel                   auth.Description              `json:"auth_model"`
	TenantToBusinessRoleMapping []identityFabricRoleMapping   `json:"tenant_to_business_role_mapping,omitempty"`
	ApproverClasses             []identityFabricApproverClass `json:"approver_classes,omitempty"`
	PrivilegedActionAttribution []string                      `json:"privileged_action_attribution,omitempty"`
	PolicyBindingSemantics      []string                      `json:"policy_binding_semantics,omitempty"`
	BreakGlassActorTreatment    []string                      `json:"break_glass_actor_treatment,omitempty"`
	AuditLineage                []string                      `json:"audit_lineage,omitempty"`
	Limitations                 []string                      `json:"limitations,omitempty"`
}

type itsmSystemContract struct {
	System                string   `json:"system"`
	CurrentState          string   `json:"current_state"`
	WriteMode             string   `json:"write_mode"`
	SupportedTicketTypes  []string `json:"supported_ticket_types,omitempty"`
	StatusSyncDiscipline  []string `json:"status_sync_discipline,omitempty"`
	OperatorOverrideModel []string `json:"operator_override_model,omitempty"`
}

type itsmLifecyclePhase struct {
	ArtifactType      string   `json:"artifact_type"`
	InternalStates    []string `json:"internal_states,omitempty"`
	ExternalLifecycle []string `json:"external_lifecycle,omitempty"`
	ClosureConditions []string `json:"closure_conditions,omitempty"`
}

type itsmLifecycleScopeSummary struct {
	IncidentCount                       int `json:"incident_count"`
	ResolvedIncidentCount               int `json:"resolved_incident_count"`
	RecommendationCount                 int `json:"recommendation_count"`
	ApprovalRequiredRecommendationCount int `json:"approval_required_recommendation_count"`
	VerifiedRecommendationCount         int `json:"verified_recommendation_count"`
	PendingExceptionCount               int `json:"pending_exception_count"`
}

type itsmLifecycleResponse struct {
	SchemaVersion          string                    `json:"schema_version"`
	Systems                []itsmSystemContract      `json:"systems,omitempty"`
	LifecyclePhases        []itsmLifecyclePhase      `json:"lifecycle_phases,omitempty"`
	TicketClasses          []itsmTicketClass         `json:"ticket_classes,omitempty"`
	StateSyncRules         []itsmStateSyncRule       `json:"state_sync_rules,omitempty"`
	ScopeSummary           itsmLifecycleScopeSummary `json:"scope_summary"`
	EvidenceLinkage        []string                  `json:"evidence_linkage,omitempty"`
	ReassignmentEscalation []string                  `json:"reassignment_escalation,omitempty"`
	ClosureDiscipline      []string                  `json:"closure_discipline,omitempty"`
	OperatorOverrides      []string                  `json:"operator_overrides,omitempty"`
	Limitations            []string                  `json:"limitations,omitempty"`
}

type itsmTicketClass struct {
	TicketClass      string   `json:"ticket_class"`
	SourceSurface    string   `json:"source_surface"`
	ExternalType     string   `json:"external_type"`
	ApprovalRequired bool     `json:"approval_required"`
	ClosureBoundBy   []string `json:"closure_bound_by,omitempty"`
}

type itsmStateSyncRule struct {
	ArtifactType     string `json:"artifact_type"`
	InternalState    string `json:"internal_state"`
	ExternalState    string `json:"external_state"`
	SyncAction       string `json:"sync_action"`
	VerificationGate string `json:"verification_gate,omitempty"`
}

type itsmLifecycleFlowItem struct {
	FlowID               string   `json:"flow_id"`
	IncidentRef          string   `json:"incident_ref"`
	TicketClass          string   `json:"ticket_class"`
	CurrentState         string   `json:"current_state"`
	ExternalState        string   `json:"external_state"`
	ApprovalRequired     bool     `json:"approval_required"`
	Owner                string   `json:"owner,omitempty"`
	EscalationPath       []string `json:"escalation_path,omitempty"`
	ClosureReady         bool     `json:"closure_ready"`
	ClosureBlockers      []string `json:"closure_blockers,omitempty"`
	LinkedEvidenceRefs   []string `json:"linked_evidence_refs,omitempty"`
	LinkedResourceRefs   []string `json:"linked_resource_refs,omitempty"`
	OperatorOverrideRefs []string `json:"operator_override_refs,omitempty"`
}

type itsmLifecycleFlowsResponse struct {
	SchemaVersion string                  `json:"schema_version"`
	Items         []itsmLifecycleFlowItem `json:"items,omitempty"`
	Limitations   []string                `json:"limitations,omitempty"`
}

type siemOutboundContract struct {
	Endpoint           string   `json:"endpoint"`
	SchemaFields       []string `json:"schema_fields,omitempty"`
	CorrelationIDField string   `json:"correlation_id_field"`
	DecisionDiscipline string   `json:"decision_discipline"`
}

type siemInboundTrustClass struct {
	SourceTrust       string   `json:"source_trust"`
	MaxActionability  string   `json:"max_actionability"`
	LocalPolicyGate   string   `json:"local_policy_gate"`
	SafetyConstraints []string `json:"safety_constraints,omitempty"`
}

type siemSeverityNormalization struct {
	ExternalSeverity   string `json:"external_severity"`
	NormalizedSeverity string `json:"normalized_severity"`
}

type siemSyncResponse struct {
	SchemaVersion           string                      `json:"schema_version"`
	CurrentState            string                      `json:"current_state"`
	OutboundContract        siemOutboundContract        `json:"outbound_contract"`
	InboundEvaluateEndpoint string                      `json:"inbound_evaluate_endpoint"`
	SupportedSignalTypes    []string                    `json:"supported_signal_types,omitempty"`
	SourceTrustClasses      []siemInboundTrustClass     `json:"source_trust_classes,omitempty"`
	SeverityNormalization   []siemSeverityNormalization `json:"severity_normalization,omitempty"`
	ActionMappingMatrix     []siemActionMappingRule     `json:"action_mapping_matrix,omitempty"`
	SourceLabeling          []string                    `json:"source_labeling,omitempty"`
	Limitations             []string                    `json:"limitations,omitempty"`
}

type siemActionMappingRule struct {
	SourceTrust          string   `json:"source_trust"`
	SignalType           string   `json:"signal_type"`
	SeverityBand         string   `json:"severity_band"`
	Actionability        string   `json:"actionability"`
	MappedRecommendation string   `json:"mapped_recommendation"`
	ApprovalMode         string   `json:"approval_mode"`
	SafetyLimitRef       string   `json:"safety_limit_ref"`
	Notes                []string `json:"notes,omitempty"`
}

type siemSignalEvaluationRequest struct {
	SchemaVersion string   `json:"schema_version,omitempty"`
	SourceSystem  string   `json:"source_system"`
	SourceTrust   string   `json:"source_trust"`
	SignalType    string   `json:"signal_type"`
	Severity      string   `json:"severity"`
	SubjectType   string   `json:"subject_type,omitempty"`
	SubjectRef    string   `json:"subject_ref,omitempty"`
	CorrelationID string   `json:"correlation_id,omitempty"`
	HintedAction  string   `json:"hinted_action,omitempty"`
	EvidenceRefs  []string `json:"evidence_refs,omitempty"`
}

type siemSignalEvaluationResponse struct {
	SchemaVersion           string   `json:"schema_version"`
	SourceSystem            string   `json:"source_system"`
	SourceTrust             string   `json:"source_trust"`
	SourceTrustLabel        string   `json:"source_trust_label"`
	SignalType              string   `json:"signal_type"`
	CorrelationID           string   `json:"correlation_id,omitempty"`
	NormalizedSeverity      string   `json:"normalized_severity"`
	SeverityLabel           string   `json:"severity_label"`
	ActionabilityState      string   `json:"actionability_state"`
	LocalPolicyGate         string   `json:"local_policy_gate"`
	MappedRecommendation    string   `json:"mapped_recommendation"`
	ApprovalMode            string   `json:"approval_mode"`
	MappedWorkflowState     string   `json:"mapped_workflow_state"`
	SafetyLimitRef          string   `json:"safety_limit_ref"`
	ResponseHintDisposition string   `json:"response_hint_disposition"`
	Explanation             []string `json:"explanation,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type incidentCollaborationRecommendationSummary struct {
	RecommendationID string `json:"recommendation_id"`
	Title            string `json:"title"`
	Status           string `json:"status"`
	ApprovalMode     string `json:"approval_mode"`
	TemplateID       string `json:"template_id"`
}

type incidentCollaborationExportVariant struct {
	Audience       string `json:"audience"`
	URI            string `json:"uri"`
	DisclosureMode string `json:"disclosure_mode"`
}

type incidentCollaborationApprovalVisibility struct {
	SourceType      string `json:"source_type"`
	SourceRef       string `json:"source_ref"`
	ApprovalMode    string `json:"approval_mode"`
	CurrentState    string `json:"current_state"`
	ActorAttributed bool   `json:"actor_attributed"`
}

type incidentCollaborationProgress struct {
	TotalRecommendations   int  `json:"total_recommendations"`
	ApprovalPending        int  `json:"approval_pending"`
	Executed               int  `json:"executed"`
	Verified               int  `json:"verified"`
	OpenOrWatchingIncident bool `json:"open_or_watching_incident"`
}

type incidentCollaborationResponse struct {
	SchemaVersion                string                                       `json:"schema_version"`
	IncidentRef                  string                                       `json:"incident_ref"`
	SharedContext                investigationIncident                        `json:"shared_context"`
	SharedContextModel           []string                                     `json:"shared_context_model,omitempty"`
	LinkedEvidenceRefs           []string                                     `json:"linked_evidence_refs,omitempty"`
	ReadbackRefs                 []string                                     `json:"readback_refs,omitempty"`
	HandoffRefs                  []string                                     `json:"handoff_refs,omitempty"`
	ValidationRefs               []string                                     `json:"validation_refs,omitempty"`
	Recommendations              []incidentCollaborationRecommendationSummary `json:"recommendations,omitempty"`
	ApprovalVisibility           []incidentCollaborationApprovalVisibility    `json:"approval_visibility,omitempty"`
	RemediationProgress          incidentCollaborationProgress                `json:"remediation_progress"`
	ClosureBlockers              []string                                     `json:"closure_blockers,omitempty"`
	VerificationState            string                                       `json:"verification_state"`
	VerificationAfterRemediation incidentCollaborationVerification            `json:"verification_after_remediation"`
	ExportVariants               []incidentCollaborationExportVariant         `json:"export_variants,omitempty"`
	AudienceExportDiscipline     []string                                     `json:"audience_export_discipline,omitempty"`
	Limitations                  []string                                     `json:"limitations,omitempty"`
}

type incidentCollaborationVerification struct {
	CurrentState string   `json:"current_state"`
	NextGate     string   `json:"next_gate"`
	Blockers     []string `json:"blockers,omitempty"`
}

type integrationConnectorSafety struct {
	ConnectorID          string   `json:"connector_id"`
	CurrentState         string   `json:"current_state"`
	SourceClass          string   `json:"source_class"`
	WritePermissions     string   `json:"write_permissions"`
	ReplaySafeBehavior   []string `json:"replay_safe_behavior,omitempty"`
	RateLimitSemantics   []string `json:"rate_limit_semantics,omitempty"`
	DegradedModeBehavior []string `json:"degraded_mode_behavior,omitempty"`
	HealthSummary        string   `json:"health_summary"`
	Auditability         string   `json:"auditability"`
}

type integrationSafetyResponse struct {
	SchemaVersion          string                       `json:"schema_version"`
	NoNewTruthLayer        bool                         `json:"no_new_truth_layer"`
	TrustedInboundSources  []string                     `json:"trusted_inbound_sources,omitempty"`
	AdvisoryInboundSources []string                     `json:"advisory_inbound_sources,omitempty"`
	BoundedOutboundSinks   []string                     `json:"bounded_outbound_sinks,omitempty"`
	Connectors             []integrationConnectorSafety `json:"connectors,omitempty"`
	Limitations            []string                     `json:"limitations,omitempty"`
}

type integrationConnectorHealth struct {
	ConnectorID            string   `json:"connector_id"`
	HealthState            string   `json:"health_state"`
	SourceClass            string   `json:"source_class"`
	VisibilitySurfaces     []string `json:"visibility_surfaces,omitempty"`
	ReplaySafe             bool     `json:"replay_safe"`
	BoundedWritePermission string   `json:"bounded_write_permission"`
	DegradedBehavior       string   `json:"degraded_behavior"`
}

type integrationSafetyHealthResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	Connectors    []integrationConnectorHealth `json:"connectors,omitempty"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

func (s server) identityFabricHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, buildIdentityFabricResponse(principal, s.authConfig.Describe()))
}

func (s server) itsmLifecycleHandler(w http.ResponseWriter, r *http.Request) {
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
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildITSMLifecycleResponse(ctx, r)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) itsmLifecycleFlowsHandler(w http.ResponseWriter, r *http.Request) {
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
	incidentID := strings.TrimSpace(r.URL.Query().Get("incident_id"))
	if incidentID == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "incident_id is required"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildITSMLifecycleFlowsResponse(ctx, r, incidentID)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) siemSyncHandler(w http.ResponseWriter, r *http.Request) {
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
	switch r.Method {
	case http.MethodGet:
		httpjson.Write(w, http.StatusOK, buildSIEMSyncResponse())
	case http.MethodPost:
		var request siemSignalEvaluationRequest
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "invalid siem evaluation request"})
			return
		}
		httpjson.Write(w, http.StatusOK, evaluateSIEMSignal(request))
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) incidentCollaborationHandler(w http.ResponseWriter, r *http.Request) {
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
	incidentID := strings.TrimSpace(r.URL.Query().Get("incident_id"))
	if incidentID == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "incident_id is required"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildIncidentCollaborationResponse(ctx, r, incidentID)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) integrationSafetyHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, s.buildIntegrationSafetyResponse())
}

func (s server) integrationSafetyHealthHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, s.buildIntegrationSafetyHealthResponse())
}

func buildIdentityFabricResponse(principal auth.Principal, model auth.Description) identityFabricResponse {
	roleBindingDetails := model.RoleBindingDetails
	if len(roleBindingDetails) == 0 {
		roleBindingDetails = make([]auth.RoleBindingDescription, 0, len(model.RoleBindings))
		for role, bindings := range model.RoleBindings {
			roleBindingDetails = append(roleBindingDetails, auth.RoleBindingDescription{
				ChangelockRole: role,
				BindingValues:  cloneStrings(bindings),
			})
		}
	}
	roleMappings := make([]identityFabricRoleMapping, 0, len(roleBindingDetails))
	for _, item := range roleBindingDetails {
		roleMappings = append(roleMappings, identityFabricRoleMapping{
			BusinessRole:   businessRoleForChangelockRole(item.ChangelockRole),
			ChangelockRole: item.ChangelockRole,
			BindingValues:  item.BindingValues,
			Subjects:       item.Subjects,
		})
	}
	sort.Slice(roleMappings, func(i, j int) bool {
		if roleMappings[i].BusinessRole == roleMappings[j].BusinessRole {
			return roleMappings[i].ChangelockRole < roleMappings[j].ChangelockRole
		}
		return roleMappings[i].BusinessRole < roleMappings[j].BusinessRole
	})

	return identityFabricResponse{
		SchemaVersion:               identityFabricSchemaVersion,
		CurrentActor:                authInfoFromPrincipal(principal),
		AuthModel:                   model,
		TenantToBusinessRoleMapping: roleMappings,
		ApproverClasses: []identityFabricApproverClass{
			{
				ClassID:             "incident_operator",
				DisplayName:         "Incident operator",
				AllowedRoles:        []string{auth.RoleOperator, auth.RoleSecurityAdmin},
				DelegationSemantics: "Operators can request review, assignment, and remediation routing, but privileged closures stay bounded by verification and approval contracts.",
				Capabilities:        []string{"request_review", "assign_owner", "route_remediation"},
			},
			{
				ClassID:             "security_governance_approver",
				DisplayName:         "Security governance approver",
				AllowedRoles:        []string{auth.RoleSecurityAdmin},
				DelegationSemantics: "Approvals remain actor-attributed; break-glass and exception semantics still require explicit approver lineage in audit.",
				Capabilities:        []string{"approve_exception", "approve_recovery", "approve_governance_workflow"},
			},
		},
		PrivilegedActionAttribution: []string{
			"Recommendation approval requests, exception approvals, runtime recovery, and incident resolution keep actor identity lineage in canonical audit.",
			"Privileged actions remain attributable to authenticated human or service principals; no local role mapping bypass is introduced by this surface.",
		},
		PolicyBindingSemantics: []string{
			"Business-role bindings map enterprise identity claims or configured subjects into ChangeLock roles before approval and remediation workflows execute.",
			"Tenant-scoped actors remain pinned to their tenant unless global security-admin behavior is explicitly allowed by the configured auth model.",
		},
		BreakGlassActorTreatment: []string{
			"Break-glass remains a bounded exception path with explicit requested_by and approved_by lineage.",
			"Break-glass treatment does not bypass actor attribution, tenant scope, or later verification and cleanup review.",
		},
		AuditLineage: []string{
			"/v1/auth/me",
			"/v1/exceptions",
			"/v1/recommendations",
			"/v1/incidents",
		},
		Limitations: append([]string{
			"Identity fabric integration summarizes configured role and actor semantics; it does not claim upstream workflow delegation from an external IdP directory without explicit connector evidence.",
		}, model.Limitations...),
	}
}

func (s server) buildITSMLifecycleResponse(ctx context.Context, r *http.Request) (itsmLifecycleResponse, error) {
	incidentFilter, err := parseIncidentFilter(r)
	if err != nil {
		return itsmLifecycleResponse{}, err
	}
	recommendationFilter, err := parseRecommendationFilter(r)
	if err != nil {
		return itsmLifecycleResponse{}, err
	}
	exceptionFilter, err := parseExceptionFilter(r)
	if err != nil {
		return itsmLifecycleResponse{}, err
	}

	incidents, err := s.listIncidents(ctx, incidentFilter)
	if err != nil {
		return itsmLifecycleResponse{}, err
	}
	recommendations, err := s.listRecommendations(ctx, recommendationFilter)
	if err != nil {
		return itsmLifecycleResponse{}, err
	}
	exceptionFilter.Status = audit.ExceptionStatusPending
	pendingExceptions, err := s.store.ListExceptions(ctx, exceptionFilter)
	if err != nil {
		return itsmLifecycleResponse{}, err
	}

	summary := itsmLifecycleScopeSummary{
		IncidentCount:         len(incidents),
		RecommendationCount:   len(recommendations),
		PendingExceptionCount: len(pendingExceptions),
	}
	for _, incident := range incidents {
		if incident.State == incidentStateResolved {
			summary.ResolvedIncidentCount++
		}
	}
	for _, recommendation := range recommendations {
		if recommendation.ApprovalMode == recommendationApprovalHumanReview {
			summary.ApprovalRequiredRecommendationCount++
		}
		if strings.HasPrefix(recommendation.Status, "verified_") {
			summary.VerifiedRecommendationCount++
		}
	}

	return itsmLifecycleResponse{
		SchemaVersion: itsmLifecycleSchemaVersion,
		Systems: []itsmSystemContract{
			{
				System:               "jira",
				CurrentState:         "draft_contract_ready",
				WriteMode:            "draft_before_write_only",
				SupportedTicketTypes: []string{"incident", "remediation", "approval"},
				StatusSyncDiscipline: []string{"create", "update", "resolve_or_close"},
				OperatorOverrideModel: []string{
					"Recommendation and incident workflows can be reassigned, commented, reopened, or superseded without mutating external truth automatically.",
				},
			},
			{
				System:               "servicenow",
				CurrentState:         "draft_contract_ready",
				WriteMode:            "draft_before_write_only",
				SupportedTicketTypes: []string{"incident", "remediation", "approval"},
				StatusSyncDiscipline: []string{"create", "update", "resolve_or_close"},
				OperatorOverrideModel: []string{
					"Closure remains tied to verification state, not merely to an external status transition.",
				},
			},
		},
		LifecyclePhases: []itsmLifecyclePhase{
			{
				ArtifactType:      "incident",
				InternalStates:    []string{incidentStateOpen, incidentStateAcknowledged, incidentStateWatching, incidentStateResolved, incidentStateReopened},
				ExternalLifecycle: []string{"create", "update", "resolve", "reopen"},
				ClosureConditions: []string{"incident state resolved", "resolution refs attached", "verification state no longer contradicts closure"},
			},
			{
				ArtifactType:      "remediation",
				InternalStates:    []string{recommendationStatusShown, recommendationStatusAccepted, recommendationStatusExecuted, recommendationStatusVerifiedSuccessful, recommendationStatusExecutedNoEffect, recommendationStatusPartiallyEffective},
				ExternalLifecycle: []string{"create", "update", "resolve"},
				ClosureConditions: []string{"recommendation verified or superseded", "linked incident pressure reduced or closed"},
			},
			{
				ArtifactType:      "approval",
				InternalStates:    []string{audit.ExceptionStatusPending, audit.ExceptionStatusApproved, audit.ExceptionStatusRejected, audit.ExceptionStatusRevoked},
				ExternalLifecycle: []string{"create", "update", "close"},
				ClosureConditions: []string{"approval explicitly approved or rejected", "revocation keeps audit trail rather than deleting the workflow"},
			},
		},
		TicketClasses: []itsmTicketClass{
			{
				TicketClass:      "incident",
				SourceSurface:    "/v1/incidents",
				ExternalType:     "incident",
				ApprovalRequired: false,
				ClosureBoundBy:   []string{"incident resolution", "verification state"},
			},
			{
				TicketClass:      "remediation",
				SourceSurface:    "/v1/recommendations",
				ExternalType:     "task",
				ApprovalRequired: false,
				ClosureBoundBy:   []string{"recommendation verify or supersede"},
			},
			{
				TicketClass:      "approval",
				SourceSurface:    "/v1/recommendations approval-request and /v1/exceptions",
				ExternalType:     "change_review",
				ApprovalRequired: true,
				ClosureBoundBy:   []string{"explicit approve or reject", "revocation remains visible"},
			},
		},
		StateSyncRules: []itsmStateSyncRule{
			{ArtifactType: "incident", InternalState: incidentStateOpen, ExternalState: "open", SyncAction: "create_or_update", VerificationGate: "none"},
			{ArtifactType: "incident", InternalState: incidentStateResolved, ExternalState: "resolved", SyncAction: "resolve", VerificationGate: "resolution evidence remains linked"},
			{ArtifactType: "remediation", InternalState: recommendationStatusShown, ExternalState: "open", SyncAction: "create"},
			{ArtifactType: "remediation", InternalState: recommendationStatusExecuted, ExternalState: "in_progress", SyncAction: "update", VerificationGate: "verification pending"},
			{ArtifactType: "remediation", InternalState: recommendationStatusVerifiedSuccessful, ExternalState: "resolved", SyncAction: "resolve", VerificationGate: "verified successful"},
			{ArtifactType: "approval", InternalState: audit.ExceptionStatusPending, ExternalState: "awaiting_approval", SyncAction: "create"},
			{ArtifactType: "approval", InternalState: audit.ExceptionStatusApproved, ExternalState: "approved", SyncAction: "update", VerificationGate: "approval lineage retained"},
			{ArtifactType: "approval", InternalState: audit.ExceptionStatusRejected, ExternalState: "rejected", SyncAction: "close", VerificationGate: "rejection reason retained"},
		},
		ScopeSummary: summary,
		EvidenceLinkage: []string{
			"/v1/incidents/{incident_id}",
			"/v1/incidents/{incident_id}/export",
			"/v1/recommendations/{recommendation_id}",
			"/v1/exceptions/{exception_id}",
			"/v1/handoff/{package_id}",
			"/v1/validation/executions/{run_id}",
		},
		ReassignmentEscalation: []string{
			"Incident assignment and recommendation assignment remain separate from closure and verification so ownership can change without losing lineage.",
			"Approval request flow stays explicit through recommendation approval-request and exception approval events.",
		},
		ClosureDiscipline: []string{
			"Incident closure is bounded by resolution and verification posture, not just by an external ticket close event.",
			"Recommendation closure remains evidence-backed through verify or supersede semantics.",
			"Approval closure keeps explicit approved/rejected/revoked lineage instead of overwriting history.",
		},
		OperatorOverrides: []string{
			"Incident reopen, reassign, and note flows remain operator-visible and do not silently overwrite historical states.",
			"Recommendation reject, assign, approval-request, execute, and verify remain separate workflow mutations with canonical audit history.",
		},
		Limitations: []string{
			"Current ITSM lifecycle surface documents a bounded workflow contract and draft discipline; it does not claim live outbound Jira or ServiceNow mutation in this slice.",
		},
	}, nil
}

func (s server) buildITSMLifecycleFlowsResponse(ctx context.Context, r *http.Request, incidentID string) (itsmLifecycleFlowsResponse, error) {
	incidentFilter, err := parseIncidentFilter(r)
	if err != nil {
		return itsmLifecycleFlowsResponse{}, err
	}
	incident, err := s.getIncidentByID(ctx, incidentID, incidentFilter)
	if err != nil {
		return itsmLifecycleFlowsResponse{}, err
	}
	recommendationFilter, err := parseRecommendationFilter(r)
	if err != nil {
		return itsmLifecycleFlowsResponse{}, err
	}
	recommendationFilter.IncidentIDs = []string{incidentID}
	recommendationFilter.Limit = maxInt(recommendationFilter.Limit, 20)
	recommendations, err := s.listRecommendations(ctx, recommendationFilter)
	if err != nil {
		return itsmLifecycleFlowsResponse{}, err
	}

	items := []itsmLifecycleFlowItem{
		{
			FlowID:               incident.ID + ":incident",
			IncidentRef:          incident.ID,
			TicketClass:          "incident",
			CurrentState:         incident.State,
			ExternalState:        externalStateForIncident(incident.State),
			ApprovalRequired:     false,
			Owner:                incident.Owner,
			EscalationPath:       []string{"/v1/incidents/" + incident.ID + "/assign", "/v1/incidents/" + incident.ID + "/state"},
			ClosureReady:         incident.State == incidentStateResolved,
			ClosureBlockers:      incidentFlowClosureBlockers(incident),
			LinkedEvidenceRefs:   uniqueStrings(incident.EvidenceRefs),
			LinkedResourceRefs:   []string{"/v1/incidents/" + incident.ID, "/v1/incidents/" + incident.ID + "/export"},
			OperatorOverrideRefs: []string{"/v1/incidents/" + incident.ID + "/reopen", "/v1/incidents/" + incident.ID + "/notes"},
		},
	}
	for _, item := range recommendations {
		items = append(items, itsmLifecycleFlowItem{
			FlowID:               item.RecommendationID + ":remediation",
			IncidentRef:          incident.ID,
			TicketClass:          "remediation",
			CurrentState:         item.Status,
			ExternalState:        externalStateForRecommendation(item.Status),
			ApprovalRequired:     item.ApprovalMode == recommendationApprovalHumanReview,
			Owner:                item.Owner,
			EscalationPath:       []string{"/v1/recommendations/" + item.RecommendationID + "/assign", "/v1/recommendations/" + item.RecommendationID + "/comment"},
			ClosureReady:         recommendationClosureReady(item),
			ClosureBlockers:      recommendationClosureBlockers(item),
			LinkedEvidenceRefs:   uniqueStrings(item.EvidenceRefs),
			LinkedResourceRefs:   append([]string{"/v1/recommendations/" + item.RecommendationID}, readbackURIs(item.ReadbackRefs)...),
			OperatorOverrideRefs: []string{"/v1/recommendations/" + item.RecommendationID + "/reject", "/v1/recommendations/" + item.RecommendationID + "/verify"},
		})
		if item.ApprovalMode == recommendationApprovalHumanReview {
			items = append(items, itsmLifecycleFlowItem{
				FlowID:               item.RecommendationID + ":approval",
				IncidentRef:          incident.ID,
				TicketClass:          "approval",
				CurrentState:         approvalStateForRecommendation(item),
				ExternalState:        externalStateForApprovalFlow(item),
				ApprovalRequired:     true,
				Owner:                item.Owner,
				EscalationPath:       []string{"/v1/recommendations/" + item.RecommendationID + "/approval-request"},
				ClosureReady:         recommendationClosureReady(item),
				ClosureBlockers:      approvalClosureBlockers(item),
				LinkedEvidenceRefs:   uniqueStrings(item.EvidenceRefs),
				LinkedResourceRefs:   []string{"/v1/recommendations/" + item.RecommendationID},
				OperatorOverrideRefs: []string{"/v1/recommendations/" + item.RecommendationID + "/accept", "/v1/recommendations/" + item.RecommendationID + "/reject"},
			})
		}
	}

	return itsmLifecycleFlowsResponse{
		SchemaVersion: itsmLifecycleFlowsSchemaVersion,
		Items:         items,
		Limitations: []string{
			"Lifecycle flow projection is derived from canonical incident and recommendation state in the current scope; it does not mutate or confirm an external ITSM record.",
		},
	}, nil
}

func buildSIEMSyncResponse() siemSyncResponse {
	return siemSyncResponse{
		SchemaVersion:           siemSyncSchemaVersion,
		CurrentState:            "bounded_policy_gate",
		InboundEvaluateEndpoint: "/v1/integrations/siem-sync/evaluate",
		OutboundContract: siemOutboundContract{
			Endpoint:           "/v1/reports/events",
			SchemaFields:       []string{"event_type", "decision", "severity", "correlation_id", "subject_ref", "subject_type", "incident_ref", "recommendation_ref", "evidence_refs"},
			CorrelationIDField: "incident_ref",
			DecisionDiscipline: "Outbound export reflects canonical local audit decisions and does not promote external sink state into new truth.",
		},
		SupportedSignalTypes: []string{"threat_intel", "enrichment", "response_hint"},
		SourceTrustClasses: []siemInboundTrustClass{
			{
				SourceTrust:      "trusted",
				MaxActionability: "bounded_recommendation",
				LocalPolicyGate:  "trusted_signal_requires_local_policy_mapping",
				SafetyConstraints: []string{
					"Trusted inbound signals can influence recommendation priority or review routing, but cannot execute runtime actions directly.",
				},
			},
			{
				SourceTrust:      "advisory",
				MaxActionability: "review_required",
				LocalPolicyGate:  "advisory_signal_requires_human_review",
				SafetyConstraints: []string{
					"Advisory sources remain enrichment only until a human or local policy flow accepts the signal.",
				},
			},
			{
				SourceTrust:      "quarantined",
				MaxActionability: "advisory_only",
				LocalPolicyGate:  "quarantined_source_never_drives_response",
				SafetyConstraints: []string{
					"Quarantined or replay-suspect sources can be retained for evidence comparison, but not for response hints.",
				},
			},
		},
		SeverityNormalization: []siemSeverityNormalization{
			{ExternalSeverity: "informational", NormalizedSeverity: "low"},
			{ExternalSeverity: "low", NormalizedSeverity: "low"},
			{ExternalSeverity: "medium", NormalizedSeverity: "medium"},
			{ExternalSeverity: "high", NormalizedSeverity: "high"},
			{ExternalSeverity: "critical", NormalizedSeverity: "critical"},
		},
		ActionMappingMatrix: []siemActionMappingRule{
			{
				SourceTrust:          "trusted",
				SignalType:           "response_hint",
				SeverityBand:         "critical",
				Actionability:        "review_required",
				MappedRecommendation: "create_security_review",
				ApprovalMode:         recommendationApprovalHumanReview,
				SafetyLimitRef:       "siem_sync_safety.v1:security_review_only",
				Notes:                []string{"Critical trusted hints can open review, but never direct runtime enforcement."},
			},
			{
				SourceTrust:          "trusted",
				SignalType:           "threat_intel",
				SeverityBand:         "high",
				Actionability:        "bounded_recommendation",
				MappedRecommendation: "open_sandbox",
				ApprovalMode:         recommendationApprovalAutoSafe,
				SafetyLimitRef:       "siem_sync_safety.v1:validation_only",
			},
			{
				SourceTrust:          "advisory",
				SignalType:           "enrichment",
				SeverityBand:         "medium",
				Actionability:        "review_required",
				MappedRecommendation: "open_sandbox",
				ApprovalMode:         recommendationApprovalHumanReview,
				SafetyLimitRef:       "siem_sync_safety.v1:human_review_required",
			},
			{
				SourceTrust:          "quarantined",
				SignalType:           "response_hint",
				SeverityBand:         "high",
				Actionability:        "advisory_only",
				MappedRecommendation: "notify_owner",
				ApprovalMode:         recommendationApprovalAutoSafe,
				SafetyLimitRef:       "siem_sync_safety.v1:quarantined_source",
			},
		},
		SourceLabeling: []string{
			"Every inbound signal keeps source_system, source_trust, correlation_id, and normalized severity visible in the local recommendation narrative.",
			"External severity is normalized before any recommendation mapping so SOAR/SIEM sinks cannot inject opaque urgency semantics.",
		},
		Limitations: []string{
			"Inbound SIEM/SOAR sync is bounded to evaluation and recommendation mapping in this slice; it does not create a new external truth layer or direct runtime control path.",
		},
	}
}

func evaluateSIEMSignal(request siemSignalEvaluationRequest) siemSignalEvaluationResponse {
	sourceSystem := firstNonEmpty(strings.TrimSpace(request.SourceSystem), "unknown")
	sourceTrust := normalizeSIEMSourceTrust(request.SourceTrust)
	signalType := normalizeSIEMSignalType(request.SignalType)
	severity := normalizeSIEMSeverity(request.Severity)
	hintedAction := strings.ToLower(strings.TrimSpace(request.HintedAction))

	response := siemSignalEvaluationResponse{
		SchemaVersion:           siemSyncEvaluationSchemaVersion,
		SourceSystem:            sourceSystem,
		SourceTrust:             sourceTrust,
		SourceTrustLabel:        sourceTrustLabel(sourceTrust),
		SignalType:              signalType,
		CorrelationID:           strings.TrimSpace(request.CorrelationID),
		NormalizedSeverity:      severity,
		SeverityLabel:           strings.ToUpper(severity),
		ActionabilityState:      "advisory_only",
		LocalPolicyGate:         "external_signal_cannot_bypass_local_policy",
		MappedRecommendation:    "notify_owner",
		ApprovalMode:            recommendationApprovalAutoSafe,
		MappedWorkflowState:     "draft_only",
		SafetyLimitRef:          "siem_sync_safety.v1:advisory_only",
		ResponseHintDisposition: "recommendation_only",
		Explanation: []string{
			"External signal is normalized before policy mapping so source-specific semantics do not bypass local controls.",
			"Result remains a recommendation or review path; no runtime enforcement is executed from this endpoint.",
		},
		Limitations: []string{
			"Evaluation is bounded to explainable recommendation mapping and does not persist, replay, or execute the suggested action.",
		},
	}

	switch sourceTrust {
	case "quarantined":
		response.ActionabilityState = "advisory_only"
		response.LocalPolicyGate = "quarantined_source_never_drives_response"
		response.MappedRecommendation = "notify_owner"
		response.SafetyLimitRef = "siem_sync_safety.v1:quarantined_source"
		response.Explanation = append(response.Explanation, "Quarantined source trust forces advisory-only treatment.")
		return response
	case "advisory":
		response.ActionabilityState = "review_required"
		response.LocalPolicyGate = "advisory_signal_requires_human_review"
		response.MappedRecommendation = "open_sandbox"
		response.ApprovalMode = recommendationApprovalHumanReview
		response.MappedWorkflowState = "pending_review"
		response.SafetyLimitRef = "siem_sync_safety.v1:human_review_required"
		response.Explanation = append(response.Explanation, "Advisory signal can open a bounded review path, but still requires local approval before governance-style execution.")
		return response
	}

	switch {
	case severity == "critical" || strings.Contains(hintedAction, "review"):
		response.ActionabilityState = "review_required"
		response.LocalPolicyGate = "trusted_signal_requires_security_review"
		response.MappedRecommendation = "create_security_review"
		response.ApprovalMode = recommendationApprovalHumanReview
		response.MappedWorkflowState = "pending_security_review"
		response.SafetyLimitRef = "siem_sync_safety.v1:security_review_only"
	case signalType == "response_hint" || strings.Contains(hintedAction, "ticket"):
		response.ActionabilityState = "bounded_recommendation"
		response.LocalPolicyGate = "trusted_signal_mapped_to_ticket_workflow"
		response.MappedRecommendation = "create_ticket"
		response.MappedWorkflowState = "ticket_draft_ready"
		response.SafetyLimitRef = "siem_sync_safety.v1:workflow_only"
	case signalType == "threat_intel":
		response.ActionabilityState = "bounded_recommendation"
		response.LocalPolicyGate = "trusted_signal_mapped_to_validation"
		response.MappedRecommendation = "open_sandbox"
		response.MappedWorkflowState = "validation_draft_ready"
		response.SafetyLimitRef = "siem_sync_safety.v1:validation_only"
	default:
		response.ActionabilityState = "bounded_recommendation"
		response.LocalPolicyGate = "trusted_signal_mapped_to_owner_notification"
		response.MappedRecommendation = "notify_owner"
		response.MappedWorkflowState = "owner_notification_ready"
		response.SafetyLimitRef = "siem_sync_safety.v1:workflow_only"
	}
	return response
}

func (s server) buildIncidentCollaborationResponse(ctx context.Context, r *http.Request, incidentID string) (incidentCollaborationResponse, error) {
	incidentFilter, err := parseIncidentFilter(r)
	if err != nil {
		return incidentCollaborationResponse{}, err
	}
	incident, err := s.getIncidentByID(ctx, incidentID, incidentFilter)
	if err != nil {
		return incidentCollaborationResponse{}, err
	}
	recommendationFilter, err := parseRecommendationFilter(r)
	if err != nil {
		return incidentCollaborationResponse{}, err
	}
	recommendationFilter.IncidentIDs = []string{incidentID}
	recommendationFilter.Limit = maxInt(recommendationFilter.Limit, 12)
	recommendations, err := s.listRecommendations(ctx, recommendationFilter)
	if err != nil {
		return incidentCollaborationResponse{}, err
	}

	recommendationSummaries := make([]incidentCollaborationRecommendationSummary, 0, len(recommendations))
	approvalVisibility := make([]incidentCollaborationApprovalVisibility, 0, len(recommendations))
	readbackRefs := []string{}
	handoffRefs := []string{}
	validationRefs := []string{}
	progress := incidentCollaborationProgress{TotalRecommendations: len(recommendations)}
	for _, item := range recommendations {
		recommendationSummaries = append(recommendationSummaries, incidentCollaborationRecommendationSummary{
			RecommendationID: item.RecommendationID,
			Title:            item.Title,
			Status:           item.Status,
			ApprovalMode:     item.ApprovalMode,
			TemplateID:       item.ActionTemplate.TemplateID,
		})
		approvalVisibility = append(approvalVisibility, incidentCollaborationApprovalVisibility{
			SourceType:      item.SourceType,
			SourceRef:       item.SourceRef,
			ApprovalMode:    item.ApprovalMode,
			CurrentState:    item.Status,
			ActorAttributed: true,
		})
		if item.ApprovalMode == recommendationApprovalHumanReview && item.Status != recommendationStatusAccepted && item.Status != recommendationStatusExecuted && !strings.HasPrefix(item.Status, "verified_") {
			progress.ApprovalPending++
		}
		if item.Status == recommendationStatusExecuted {
			progress.Executed++
		}
		if strings.HasPrefix(item.Status, "verified_") {
			progress.Verified++
		}
		for _, ref := range item.ReadbackRefs {
			if strings.TrimSpace(ref.ResourceURI) != "" {
				readbackRefs = append(readbackRefs, ref.ResourceURI)
			}
			resourceType := strings.ToLower(strings.TrimSpace(ref.ResourceType))
			switch {
			case strings.Contains(resourceType, "handoff") || strings.Contains(strings.ToLower(ref.ResourceURI), "/handoff/"):
				handoffRefs = append(handoffRefs, ref.ResourceURI)
			case strings.Contains(resourceType, "validation") || strings.Contains(strings.ToLower(ref.ResourceURI), "/validation/"):
				validationRefs = append(validationRefs, ref.ResourceURI)
			}
		}
	}
	progress.OpenOrWatchingIncident = incident.State != incidentStateResolved
	closureBlockers := incidentCollaborationClosureBlockers(incident, progress)

	return incidentCollaborationResponse{
		SchemaVersion: incidentCollaborationSchemaVersion,
		IncidentRef:   incident.ID,
		SharedContext: incident,
		SharedContextModel: []string{
			"Shared context bundles the canonical incident, linked recommendation workflow, and audience-aware export paths without creating a separate mutable incident store.",
			"Collaboration lifecycle stays tied to evidence, approval visibility, and post-remediation verification state.",
		},
		LinkedEvidenceRefs:  uniqueStrings(incident.EvidenceRefs),
		ReadbackRefs:        uniqueStrings(readbackRefs),
		HandoffRefs:         uniqueStrings(handoffRefs),
		ValidationRefs:      uniqueStrings(validationRefs),
		Recommendations:     recommendationSummaries,
		ApprovalVisibility:  approvalVisibility,
		RemediationProgress: progress,
		ClosureBlockers:     closureBlockers,
		VerificationState:   incidentCollaborationVerificationState(incident, progress),
		VerificationAfterRemediation: incidentCollaborationVerification{
			CurrentState: incidentCollaborationVerificationState(incident, progress),
			NextGate:     incidentCollaborationNextGate(incident, progress),
			Blockers:     closureBlockers,
		},
		ExportVariants: []incidentCollaborationExportVariant{
			{Audience: incidentAudienceInternal, URI: "/v1/incidents/" + incident.ID + "/export?audience=" + incidentAudienceInternal, DisclosureMode: "full_internal"},
			{Audience: incidentAudienceAuditorSafe, URI: "/v1/incidents/" + incident.ID + "/export?audience=" + incidentAudienceAuditorSafe, DisclosureMode: "audit_safe"},
			{Audience: incidentAudienceCustomerSafe, URI: "/v1/incidents/" + incident.ID + "/export?audience=" + incidentAudienceCustomerSafe, DisclosureMode: "customer_safe"},
		},
		AudienceExportDiscipline: []string{
			"Internal exports keep richer operational context, while auditor-safe and customer-safe variants remain disclosure-bounded.",
			"Audience-aware export never drops verification lineage; it only narrows narrative and sensitive detail exposure.",
		},
		Limitations: []string{
			"Incident collaboration stitches together canonical incident, recommendation, and readback surfaces already present in ChangeLock; it does not create a separate mutable collaboration truth store.",
		},
	}, nil
}

func (s server) buildIntegrationSafetyResponse() integrationSafetyResponse {
	syncStatus := audit.SyncStatus{
		Mode:     audit.SyncModeDisabled,
		SyncMode: audit.SyncModeDisabled,
		Health:   audit.SyncHealthDisabled,
		Summary:  "No cross-cluster connector is enabled.",
	}
	if s.syncRuntime != nil {
		syncStatus = s.syncRuntime.statusSnapshot()
	}

	connectors := []integrationConnectorSafety{
		{
			ConnectorID:      "identity_fabric",
			CurrentState:     firstNonEmpty(s.authConfig.Mode, auth.ModeDisabled),
			SourceClass:      "trusted_inbound",
			WritePermissions: "none",
			ReplaySafeBehavior: []string{
				"Identity mapping evaluates each request independently and does not replay external directory mutations into local workflow state.",
			},
			RateLimitSemantics: []string{
				"Bearer auth evaluation is request-bounded; upstream JWKS caching is local and fail-closed.",
			},
			DegradedModeBehavior: []string{
				"Authentication failures fail closed instead of granting a degraded allow path.",
			},
			HealthSummary: "Auth decisions remain local-policy gated.",
			Auditability:  "Actor attribution stays visible through auth, incident, recommendation, and exception events.",
		},
		{
			ConnectorID:      "itsm_workflow_drafts",
			CurrentState:     "draft_only",
			SourceClass:      "bounded_outbound",
			WritePermissions: "draft_generation_only",
			ReplaySafeBehavior: []string{
				"Draft regeneration is idempotent and does not mutate canonical incident or recommendation truth.",
			},
			RateLimitSemantics: []string{
				"External ticket writes are intentionally absent in this slice; repeated draft generation remains bounded to local reads.",
			},
			DegradedModeBehavior: []string{
				"Connector absence leaves local workflow intact because lifecycle semantics remain visible through native endpoints.",
			},
			HealthSummary: "No live ITSM write path is enabled.",
			Auditability:  "All workflow mutations stay in canonical audit before any future external sync.",
		},
		{
			ConnectorID:      "siem_sync",
			CurrentState:     "policy_gated",
			SourceClass:      "advisory_inbound",
			WritePermissions: "recommendation_mapping_only",
			ReplaySafeBehavior: []string{
				"Inbound signal evaluation is stateless and explainable; it does not persist or execute external hints.",
			},
			RateLimitSemantics: []string{
				"Outbound export is read-only via /v1/reports/events; inbound evaluation is bounded to explicit request handling.",
			},
			DegradedModeBehavior: []string{
				"Loss of external SIEM/SOAR context removes enrichment only; local evidence and policy decisions remain authoritative.",
			},
			HealthSummary: "External signals stay advisory until local policy mapping accepts them.",
			Auditability:  "Source trust and correlation semantics remain explicit in the evaluation response.",
		},
		{
			ConnectorID:      "cluster_sync",
			CurrentState:     firstNonEmpty(syncStatus.Health, audit.SyncHealthDisabled),
			SourceClass:      "trusted_inbound",
			WritePermissions: "bounded_exception_snapshot_only",
			ReplaySafeBehavior: []string{
				"Snapshot revisions and signatures prevent silent replay or overwrite of exception state.",
			},
			RateLimitSemantics: []string{
				"Sync runtime uses poll-interval and cache semantics rather than unbounded push mutation.",
			},
			DegradedModeBehavior: []string{
				firstNonEmpty(syncStatus.Summary, "Degraded sync status is surfaced explicitly instead of silently continuing."),
			},
			HealthSummary: firstNonEmpty(syncStatus.Summary, "Cross-cluster sync disabled."),
			Auditability:  "Sync state, verification, and last-known-good posture remain visible through /v1/sync/status and local audit.",
		},
	}

	return integrationSafetyResponse{
		SchemaVersion:          integrationSafetySchemaVersion,
		NoNewTruthLayer:        true,
		TrustedInboundSources:  []string{"identity_fabric", "cluster_sync_signed_snapshots"},
		AdvisoryInboundSources: []string{"siem_sync", "soar_response_hints"},
		BoundedOutboundSinks:   []string{"itsm_workflow_drafts", "events_export"},
		Connectors:             connectors,
		Limitations: []string{
			"Enterprise integrations in this slice remain bounded coordination surfaces and do not become a shadow control plane over runtime, governance, or evidence truth.",
		},
	}
}

func (s server) buildIntegrationSafetyHealthResponse() integrationSafetyHealthResponse {
	syncStatus := audit.SyncStatus{
		Mode:     audit.SyncModeDisabled,
		SyncMode: audit.SyncModeDisabled,
		Health:   audit.SyncHealthDisabled,
		Summary:  "No cross-cluster connector is enabled.",
	}
	if s.syncRuntime != nil {
		syncStatus = s.syncRuntime.statusSnapshot()
	}
	return integrationSafetyHealthResponse{
		SchemaVersion: integrationSafetyHealthSchema,
		Connectors: []integrationConnectorHealth{
			{
				ConnectorID:            "identity_fabric",
				HealthState:            "policy_gated",
				SourceClass:            "trusted_inbound",
				VisibilitySurfaces:     []string{"/v1/auth/me", "/v1/integrations/identity-fabric"},
				ReplaySafe:             true,
				BoundedWritePermission: "none",
				DegradedBehavior:       "authentication failures fail closed",
			},
			{
				ConnectorID:            "itsm_workflow_drafts",
				HealthState:            "local_only",
				SourceClass:            "bounded_outbound",
				VisibilitySurfaces:     []string{"/v1/integrations/itsm-lifecycle", "/v1/integrations/itsm-lifecycle/flows"},
				ReplaySafe:             true,
				BoundedWritePermission: "draft_generation_only",
				DegradedBehavior:       "loss of remote system leaves local workflow intact because no live write path is enabled",
			},
			{
				ConnectorID:            "siem_sync",
				HealthState:            "policy_gated",
				SourceClass:            "advisory_inbound",
				VisibilitySurfaces:     []string{"/v1/integrations/siem-sync", "/v1/integrations/siem-sync/evaluate"},
				ReplaySafe:             true,
				BoundedWritePermission: "recommendation_mapping_only",
				DegradedBehavior:       "loss of external enrichment removes hints only; local policy remains authoritative",
			},
			{
				ConnectorID:            "cluster_sync",
				HealthState:            firstNonEmpty(syncStatus.Health, audit.SyncHealthDisabled),
				SourceClass:            "trusted_inbound",
				VisibilitySurfaces:     []string{"/v1/sync/status", "/v1/integrations/safety"},
				ReplaySafe:             true,
				BoundedWritePermission: "approved_exception_snapshot_only",
				DegradedBehavior:       firstNonEmpty(syncStatus.Summary, "no cross-cluster connector is enabled"),
			},
		},
		Limitations: []string{
			"Connector health surface summarizes bounded integration posture and degraded behavior; it does not infer uptime or SLA for external systems that are not actively connected in this slice.",
		},
	}
}

func businessRoleForChangelockRole(role string) string {
	switch strings.TrimSpace(role) {
	case auth.RoleViewer:
		return "trust_reader"
	case auth.RoleOperator:
		return "incident_operator"
	case auth.RoleSecurityAdmin:
		return "security_governance_admin"
	case auth.RoleService:
		return "service_automation"
	default:
		return "unmapped"
	}
}

func authInfoFromPrincipal(principal auth.Principal) authInfoResponse {
	return authInfoResponse{
		Authenticated: principal.Authenticated,
		AuthMode:      principal.AuthMode,
		Subject:       principal.Subject,
		Role:          principal.Role,
		TokenID:       principal.TokenID,
		IdentityType:  principal.IdentityType,
		Email:         principal.Email,
		TenantID:      principal.TenantID,
		GlobalScope:   principal.GlobalScope,
	}
}

func normalizeSIEMSourceTrust(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "trusted":
		return "trusted"
	case "quarantined":
		return "quarantined"
	default:
		return "advisory"
	}
}

func normalizeSIEMSignalType(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "threat_intel":
		return "threat_intel"
	case "response_hint":
		return "response_hint"
	default:
		return "enrichment"
	}
}

func normalizeSIEMSeverity(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "critical":
		return "critical"
	case "high":
		return "high"
	case "medium":
		return "medium"
	default:
		return "low"
	}
}

func sourceTrustLabel(sourceTrust string) string {
	switch sourceTrust {
	case "trusted":
		return "TRUSTED"
	case "quarantined":
		return "QUARANTINED"
	default:
		return "ADVISORY"
	}
}

func incidentCollaborationVerificationState(incident investigationIncident, progress incidentCollaborationProgress) string {
	switch {
	case incident.State == incidentStateResolved && progress.Verified > 0:
		return "verified_closure_ready"
	case progress.ApprovalPending > 0:
		return "approval_pending"
	case progress.Executed > 0:
		return "verification_pending"
	default:
		return "active_investigation"
	}
}

func incidentCollaborationNextGate(incident investigationIncident, progress incidentCollaborationProgress) string {
	switch {
	case progress.ApprovalPending > 0:
		return "approval_resolution_required"
	case progress.Executed > 0 && progress.Verified == 0:
		return "post_remediation_verification_required"
	case incident.State != incidentStateResolved:
		return "incident_resolution_required"
	default:
		return "collaboration_ready_for_closure_review"
	}
}

func incidentCollaborationClosureBlockers(incident investigationIncident, progress incidentCollaborationProgress) []string {
	blockers := []string{}
	if incident.State != incidentStateResolved {
		blockers = append(blockers, "incident_not_resolved")
	}
	if progress.ApprovalPending > 0 {
		blockers = append(blockers, "approval_pending")
	}
	if progress.Executed > 0 && progress.Verified == 0 {
		blockers = append(blockers, "verification_pending")
	}
	return blockers
}

func externalStateForIncident(state string) string {
	switch state {
	case incidentStateResolved:
		return "resolved"
	case incidentStateAcknowledged, incidentStateWatching:
		return "in_progress"
	case incidentStateReopened:
		return "reopened"
	default:
		return "open"
	}
}

func externalStateForRecommendation(state string) string {
	switch {
	case strings.HasPrefix(state, "verified_"):
		return "resolved"
	case state == recommendationStatusExecuted:
		return "in_progress"
	case state == recommendationStatusAccepted:
		return "approved"
	default:
		return "open"
	}
}

func approvalStateForRecommendation(item recommendation) string {
	switch item.Status {
	case recommendationStatusAccepted, recommendationStatusExecuted:
		return "approved_for_execution"
	case recommendationStatusRejected:
		return "rejected"
	default:
		return "awaiting_approval"
	}
}

func externalStateForApprovalFlow(item recommendation) string {
	switch approvalStateForRecommendation(item) {
	case "approved_for_execution":
		return "approved"
	case "rejected":
		return "rejected"
	default:
		return "awaiting_approval"
	}
}

func recommendationClosureReady(item recommendation) bool {
	return item.Status == recommendationStatusVerifiedSuccessful || item.Status == recommendationStatusSuperseded
}

func recommendationClosureBlockers(item recommendation) []string {
	blockers := []string{}
	if item.ApprovalMode == recommendationApprovalHumanReview && item.Status != recommendationStatusAccepted && item.Status != recommendationStatusExecuted && !strings.HasPrefix(item.Status, "verified_") {
		blockers = append(blockers, "approval_pending")
	}
	if item.Status == recommendationStatusExecuted && !strings.HasPrefix(item.Status, "verified_") {
		blockers = append(blockers, "verification_pending")
	}
	if item.Status == recommendationStatusShown {
		blockers = append(blockers, "workflow_not_started")
	}
	return blockers
}

func approvalClosureBlockers(item recommendation) []string {
	if recommendationClosureReady(item) {
		return nil
	}
	if item.Status == recommendationStatusRejected {
		return nil
	}
	return []string{"approval_decision_not_final"}
}

func incidentFlowClosureBlockers(incident investigationIncident) []string {
	if incident.State == incidentStateResolved {
		return nil
	}
	return []string{"incident_not_resolved"}
}

func readbackURIs(items []advisoryReadbackRef) []string {
	uris := make([]string, 0, len(items))
	for _, item := range items {
		if strings.TrimSpace(item.ResourceURI) == "" {
			continue
		}
		uris = append(uris, item.ResourceURI)
	}
	return uniqueStrings(uris)
}
