package main

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	recommendationComponent = "recommendation-manager"

	recommendationEventAcknowledged    = "recommendation_acknowledged"
	recommendationEventAccepted        = "recommendation_accepted"
	recommendationEventRejected        = "recommendation_rejected"
	recommendationEventExecuted        = "recommendation_executed"
	recommendationEventVerified        = "recommendation_verified"
	recommendationEventAssigned        = "recommendation_assigned"
	recommendationEventCommented       = "recommendation_commented"
	recommendationEventApprovalRequest = "recommendation_approval_requested"

	recommendationStatusShown              = "shown"
	recommendationStatusAcknowledged       = "acknowledged"
	recommendationStatusAccepted           = "accepted"
	recommendationStatusRejected           = "rejected"
	recommendationStatusExecuted           = "executed"
	recommendationStatusExpired            = "expired"
	recommendationStatusSuperseded         = "superseded"
	recommendationStatusVerifiedSuccessful = "verified_successful"
	recommendationStatusExecutedNoEffect   = "executed_no_effect"
	recommendationStatusPartiallyEffective = "partially_effective"
	recommendationStatusRegressed          = "regressed"

	recommendationApprovalAutoSafe    = "auto_safe"
	recommendationApprovalHumanReview = "approval_required"
)

type recommendationListResponse struct {
	SchemaVersion   string           `json:"schema_version"`
	Recommendations []recommendation `json:"recommendations"`
}

type recommendationActionsResponse struct {
	SchemaVersion string                         `json:"schema_version"`
	Templates     []recommendationActionTemplate `json:"templates"`
}

type recommendation struct {
	SchemaVersion       string                       `json:"schema_version"`
	RecommendationID    string                       `json:"recommendation_id"`
	SourceType          string                       `json:"source_type"`
	SourceRef           string                       `json:"source_ref"`
	SubjectType         string                       `json:"subject_type"`
	SubjectRef          string                       `json:"subject_ref"`
	Team                string                       `json:"team,omitempty"`
	Service             string                       `json:"service,omitempty"`
	Repo                string                       `json:"repo,omitempty"`
	Environment         string                       `json:"environment,omitempty"`
	RecommendationType  string                       `json:"recommendation_type"`
	Title               string                       `json:"title"`
	Description         string                       `json:"description"`
	RecommendedAction   string                       `json:"recommended_action"`
	Rationale           string                       `json:"rationale"`
	EvidenceRefs        []string                     `json:"evidence_refs"`
	ReadbackRefs        []advisoryReadbackRef        `json:"readback_refs"`
	RelatedIncidentRefs []string                     `json:"related_incident_refs"`
	PriorityBand        string                       `json:"priority_band"`
	ImpactScore         int                          `json:"impact_score"`
	EffortScore         int                          `json:"effort_score"`
	ConfidenceScore     int                          `json:"confidence_score"`
	ApprovalMode        string                       `json:"approval_mode"`
	Status              string                       `json:"status"`
	CreatedAt           time.Time                    `json:"created_at"`
	ExpiresAt           *time.Time                   `json:"expires_at,omitempty"`
	SupersededBy        string                       `json:"superseded_by,omitempty"`
	VerificationPlan    []string                     `json:"verification_plan"`
	FeedbackSummary     string                       `json:"feedback_summary,omitempty"`
	ActionTemplate      recommendationActionTemplate `json:"action_template"`
	Owner               string                       `json:"owner,omitempty"`
	Comments            []recommendationComment      `json:"comments,omitempty"`
	History             []recommendationHistoryEntry `json:"history,omitempty"`
	Outcome             recommendationOutcome        `json:"outcome"`
	AdvisoryOnly        bool                         `json:"advisory_only"`
	Limitations         []string                     `json:"limitations"`
}

type recommendationActionTemplate struct {
	TemplateID         string   `json:"template_id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	RecommendationType string   `json:"recommendation_type"`
	ApprovalMode       string   `json:"approval_mode"`
	RequiredInputs     []string `json:"required_inputs"`
	AllowedAudiences   []string `json:"allowed_audiences"`
	Idempotent         bool     `json:"idempotent"`
	CancelSemantics    string   `json:"cancel_semantics"`
}

type recommendationOutcome struct {
	Status     string     `json:"status"`
	Summary    string     `json:"summary,omitempty"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
}

type recommendationComment struct {
	ID        string     `json:"id"`
	Comment   string     `json:"comment"`
	Actor     string     `json:"actor,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type recommendationHistoryEntry struct {
	ID        string     `json:"id"`
	EventType string     `json:"event_type"`
	Title     string     `json:"title"`
	Summary   string     `json:"summary"`
	Actor     string     `json:"actor,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type recommendationFilter struct {
	event              audit.EventFilter
	IncidentIDs        []string
	PackageIncidentIDs []string
	SourceType         string
	SubjectType        string
	RecommendationType string
	Team               string
	Service            string
	Status             string
	Limit              int
}

type recommendationActionRequest struct {
	RecommendationID string `json:"recommendation_id,omitempty"`
	TemplateID       string `json:"template_id,omitempty"`
	Summary          string `json:"summary,omitempty"`
}

type recommendationRejectRequest struct {
	Reason string `json:"reason"`
}

type recommendationAssignRequest struct {
	Owner  string `json:"owner"`
	Reason string `json:"reason,omitempty"`
}

type recommendationCommentRequest struct {
	Comment string `json:"comment"`
}

type recommendationApprovalRequest struct {
	Summary string `json:"summary,omitempty"`
}

type recommendationSyntheticState struct {
	actionSummary string
}

var recommendationTemplateCatalog = []recommendationActionTemplate{
	{
		TemplateID:         "create_ticket",
		Title:              "Create ticket",
		Description:        "Prepare a tracked remediation ticket with evidence and ownership routing.",
		RecommendationType: "workflow",
		ApprovalMode:       recommendationApprovalAutoSafe,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         true,
		CancelSemantics:    "Ticket draft can be superseded by a newer recommendation or closed manually.",
	},
	{
		TemplateID:         "open_sandbox",
		Title:              "Open sandbox",
		Description:        "Prepare a bounded validation sandbox for investigating the weakest control path safely.",
		RecommendationType: "investigation",
		ApprovalMode:       recommendationApprovalAutoSafe,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         true,
		CancelSemantics:    "Sandbox validation can be cancelled if the recommendation is superseded or verified no longer needed.",
	},
	{
		TemplateID:         "generate_remediation_draft",
		Title:              "Generate remediation draft",
		Description:        "Draft a bounded hardening plan from the linked evidence, replay, and defense-gap context.",
		RecommendationType: "remediation",
		ApprovalMode:       recommendationApprovalAutoSafe,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         true,
		CancelSemantics:    "Drafts can be regenerated after new evidence or trends change the recommended fix path.",
	},
	{
		TemplateID:         "draft_vex",
		Title:              "Draft VEX",
		Description:        "Prepare a VEX-style triage draft for evidence-backed vulnerability clarification.",
		RecommendationType: "documentation",
		ApprovalMode:       recommendationApprovalAutoSafe,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         true,
		CancelSemantics:    "Draft can be discarded without changing canonical truth.",
	},
	{
		TemplateID:         "request_exception",
		Title:              "Request exception",
		Description:        "Open a bounded exception request when remediation cannot safely complete in the current window.",
		RecommendationType: "governance",
		ApprovalMode:       recommendationApprovalHumanReview,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         false,
		CancelSemantics:    "Approval is required before any exception request is executed.",
	},
	{
		TemplateID:         "notify_owner",
		Title:              "Notify owner",
		Description:        "Route the recommendation to the current owner or responder team.",
		RecommendationType: "workflow",
		ApprovalMode:       recommendationApprovalAutoSafe,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         true,
		CancelSemantics:    "Notifications can be repeated when ownership changes.",
	},
	{
		TemplateID:         "create_security_review",
		Title:              "Create security review",
		Description:        "Escalate repeated weakness patterns into a formal security review path.",
		RecommendationType: "governance",
		ApprovalMode:       recommendationApprovalHumanReview,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         false,
		CancelSemantics:    "Review request requires human approval before execution.",
	},
	{
		TemplateID:         "compare_artifact_versions",
		Title:              "Compare artifact versions",
		Description:        "Prepare an artifact-to-artifact delta investigation package.",
		RecommendationType: "investigation",
		ApprovalMode:       recommendationApprovalAutoSafe,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         true,
		CancelSemantics:    "Comparison drafts can be rerun as new versions arrive.",
	},
	{
		TemplateID:         "archive_stale_exception",
		Title:              "Archive stale exception",
		Description:        "Queue a stale exception for review and retirement.",
		RecommendationType: "governance",
		ApprovalMode:       recommendationApprovalHumanReview,
		RequiredInputs:     []string{"recommendation_id"},
		AllowedAudiences:   []string{incidentAudienceInternal},
		Idempotent:         false,
		CancelSemantics:    "Review request must be approved before retirement is executed.",
	},
}

var errRecommendationApprovalRequired = errors.New("recommendation approval required before execution")

func (s server) recommendationsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseRecommendationFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	recommendations, err := s.listRecommendations(ctx, filter)
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, recommendationListResponse{
		SchemaVersion:   recommendationListSchemaVersion,
		Recommendations: recommendations,
	})
}

func (s server) recommendationActionsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	path := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/recommendation-actions"))
	if path == "" || path == "/" {
		if r.Method != http.MethodGet {
			httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		httpjson.Write(w, http.StatusOK, recommendationActionsResponse{
			SchemaVersion: recommendationTemplatesSchemaVersion,
			Templates:     recommendationTemplateCatalog,
		})
		return
	}
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	templateID := strings.TrimSpace(strings.TrimPrefix(path, "/"))
	if templateID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "recommendation action not found"})
		return
	}
	template, ok := recommendationTemplateByID(templateID)
	if !ok {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "recommendation action not found"})
		return
	}
	var request recommendationActionRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	recommendationID := strings.TrimSpace(request.RecommendationID)
	if recommendationID == "" {
		recommendationID = strings.TrimSpace(r.URL.Query().Get("recommendation_id"))
	}
	if recommendationID == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "recommendation_id is required"})
		return
	}
	filter, err := parseRecommendationFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	recommendation, err := s.getRecommendationByID(ctx, recommendationID, filter)
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	updated, err := s.executeRecommendation(ctx, principal, recommendation, template, "")
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) recommendationByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/recommendations/"))
	if path == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "recommendation not found"})
		return
	}
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "recommendation not found"})
		return
	}
	recommendationID := strings.TrimSpace(parts[0])
	action := ""
	if len(parts) == 2 {
		action = strings.TrimSpace(parts[1])
	}
	if recommendationID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "recommendation not found"})
		return
	}

	if r.Method == http.MethodGet && action == "" {
		s.getRecommendationHandler(w, r, recommendationID)
		return
	}

	switch action {
	case "acknowledge":
		if r.Method == http.MethodPost {
			s.acknowledgeRecommendationHandler(w, r, recommendationID)
			return
		}
	case "accept":
		if r.Method == http.MethodPost {
			s.acceptRecommendationHandler(w, r, recommendationID)
			return
		}
	case "reject":
		if r.Method == http.MethodPost {
			s.rejectRecommendationHandler(w, r, recommendationID)
			return
		}
	case "execute":
		if r.Method == http.MethodPost {
			s.executeRecommendationHandler(w, r, recommendationID)
			return
		}
	case "verify":
		if r.Method == http.MethodPost {
			s.verifyRecommendationHandler(w, r, recommendationID)
			return
		}
	case "assign":
		if r.Method == http.MethodPost {
			s.assignRecommendationHandler(w, r, recommendationID)
			return
		}
	case "comment":
		if r.Method == http.MethodPost {
			s.commentRecommendationHandler(w, r, recommendationID)
			return
		}
	case "approval-request":
		if r.Method == http.MethodPost {
			s.approvalRequestRecommendationHandler(w, r, recommendationID)
			return
		}
	}

	httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func (s server) getRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRecommendationFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	recommendation, err := s.getRecommendationByID(ctx, recommendationID, filter)
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, recommendation)
}

func (s server) acknowledgeRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	s.mutateRecommendationStatusHandler(w, r, recommendationID, recommendationEventAcknowledged, recommendationStatusAcknowledged, nil)
}

func (s server) acceptRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	s.mutateRecommendationStatusHandler(w, r, recommendationID, recommendationEventAccepted, recommendationStatusAccepted, nil)
}

func (s server) rejectRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	principal, recommendation, _, ctx, cancel, ok := s.authorizeRecommendationMutation(w, r, recommendationID)
	if !ok {
		return
	}
	defer cancel()
	var request recommendationRejectRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	reason := strings.TrimSpace(request.Reason)
	if reason == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "reason is required"})
		return
	}
	updated, err := s.recordRecommendationMutation(ctx, principal, recommendation, recommendationEventRejected, func(event *audit.Event) {
		event.RecommendationStatus = recommendationStatusRejected
		event.RecommendationFeedbackSummary = reason
	})
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) executeRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	principal, recommendation, _, ctx, cancel, ok := s.authorizeRecommendationMutation(w, r, recommendationID)
	if !ok {
		return
	}
	defer cancel()
	var request recommendationActionRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	template := recommendation.ActionTemplate
	if strings.TrimSpace(request.TemplateID) != "" {
		var found bool
		template, found = recommendationTemplateByID(request.TemplateID)
		if !found {
			httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "recommendation action not found"})
			return
		}
	}
	updated, err := s.executeRecommendation(ctx, principal, recommendation, template, strings.TrimSpace(request.Summary))
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) verifyRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	principal, recommendation, filter, ctx, cancel, ok := s.authorizeRecommendationMutation(w, r, recommendationID)
	if !ok {
		return
	}
	defer cancel()
	resultStatus, summary, err := s.verifyRecommendation(ctx, recommendation, filter)
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	updated, err := s.recordRecommendationMutation(ctx, principal, recommendation, recommendationEventVerified, func(event *audit.Event) {
		event.RecommendationStatus = resultStatus
		event.RecommendationVerificationResult = resultStatus
		event.RecommendationFeedbackSummary = summary
	})
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) assignRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	principal, recommendation, _, ctx, cancel, ok := s.authorizeRecommendationMutation(w, r, recommendationID)
	if !ok {
		return
	}
	defer cancel()
	var request recommendationAssignRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	owner := strings.TrimSpace(request.Owner)
	if owner == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "owner is required"})
		return
	}
	updated, err := s.recordRecommendationMutation(ctx, principal, recommendation, recommendationEventAssigned, func(event *audit.Event) {
		event.RecommendationOwner = owner
		event.RecommendationFeedbackSummary = strings.TrimSpace(request.Reason)
	})
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) commentRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	principal, recommendation, _, ctx, cancel, ok := s.authorizeRecommendationMutation(w, r, recommendationID)
	if !ok {
		return
	}
	defer cancel()
	var request recommendationCommentRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	comment := strings.TrimSpace(request.Comment)
	if comment == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "comment is required"})
		return
	}
	updated, err := s.recordRecommendationMutation(ctx, principal, recommendation, recommendationEventCommented, func(event *audit.Event) {
		event.RecommendationComment = comment
	})
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) approvalRequestRecommendationHandler(w http.ResponseWriter, r *http.Request, recommendationID string) {
	principal, recommendation, _, ctx, cancel, ok := s.authorizeRecommendationMutation(w, r, recommendationID)
	if !ok {
		return
	}
	defer cancel()
	var request recommendationApprovalRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	updated, err := s.recordRecommendationMutation(ctx, principal, recommendation, recommendationEventApprovalRequest, func(event *audit.Event) {
		event.RecommendationFeedbackSummary = strings.TrimSpace(firstNonEmpty(request.Summary, "Approval requested for a sensitive workflow action."))
	})
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) mutateRecommendationStatusHandler(w http.ResponseWriter, r *http.Request, recommendationID string, eventType string, status string, mutate func(*audit.Event)) {
	principal, recommendation, _, ctx, cancel, ok := s.authorizeRecommendationMutation(w, r, recommendationID)
	if !ok {
		return
	}
	defer cancel()
	updated, err := s.recordRecommendationMutation(ctx, principal, recommendation, eventType, func(event *audit.Event) {
		event.RecommendationStatus = status
		if mutate != nil {
			mutate(event)
		}
	})
	if err != nil {
		writeRecommendationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) authorizeRecommendationMutation(w http.ResponseWriter, r *http.Request, recommendationID string) (auth.Principal, recommendation, recommendationFilter, context.Context, context.CancelFunc, bool) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return auth.Principal{}, recommendation{}, recommendationFilter{}, nil, nil, false
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return auth.Principal{}, recommendation{}, recommendationFilter{}, nil, nil, false
	}
	filter, err := parseRecommendationFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return auth.Principal{}, recommendation{}, recommendationFilter{}, nil, nil, false
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	currentRecommendation, err := s.getRecommendationByID(ctx, recommendationID, filter)
	if err != nil {
		cancel()
		writeRecommendationError(w, err)
		return auth.Principal{}, recommendation{}, recommendationFilter{}, nil, nil, false
	}
	return principal, currentRecommendation, filter, ctx, cancel, true
}

func parseRecommendationFilter(r *http.Request) (recommendationFilter, error) {
	base, err := parseFilter(r)
	if err != nil {
		return recommendationFilter{}, err
	}
	requestedLimit := base.Limit
	base.Decision = ""
	base.Component = ""
	base.EventType = ""
	if base.Limit <= 0 {
		base.Limit = 25
	}
	if base.Limit < 100 {
		base.Limit = 500
	}
	query := r.URL.Query()
	filter := recommendationFilter{
		event:              base,
		IncidentIDs:        uniqueStrings(query["incident_id"]),
		PackageIncidentIDs: uniqueStrings(query["package_incident_id"]),
		SourceType:         strings.TrimSpace(query.Get("source_type")),
		SubjectType:        strings.TrimSpace(query.Get("subject_type")),
		RecommendationType: strings.TrimSpace(query.Get("recommendation_type")),
		Team:               strings.TrimSpace(query.Get("team")),
		Service:            strings.TrimSpace(query.Get("service")),
		Status:             strings.TrimSpace(query.Get("status")),
		Limit:              minInt(requestedLimit, 50),
	}
	if filter.Limit <= 0 {
		filter.Limit = 12
	}
	return filter, nil
}

func (s server) listRecommendations(ctx context.Context, filter recommendationFilter) ([]recommendation, error) {
	incidents, err := s.listIncidents(ctx, incidentFilter{event: filter.event})
	if err != nil {
		return nil, err
	}
	candidates, err := s.buildRecommendationCandidates(ctx, incidents, filter)
	if err != nil {
		return nil, err
	}
	events, err := s.listRecommendationMutationEvents(ctx, filter)
	if err != nil {
		return nil, err
	}
	recommendations := applyRecommendationMutations(candidates, events)
	recommendations = filterRecommendations(recommendations, filter)
	sortRecommendations(recommendations)
	if len(recommendations) > filter.Limit {
		recommendations = recommendations[:filter.Limit]
	}
	for i := range recommendations {
		recommendations[i].SchemaVersion = recommendationSchemaVersion
	}
	return recommendations, nil
}

func (s server) getRecommendationByID(ctx context.Context, recommendationID string, filter recommendationFilter) (recommendation, error) {
	recommendations, err := s.listRecommendations(ctx, recommendationFilter{
		event:              filter.event,
		IncidentIDs:        filter.IncidentIDs,
		PackageIncidentIDs: filter.PackageIncidentIDs,
		SourceType:         filter.SourceType,
		SubjectType:        filter.SubjectType,
		RecommendationType: filter.RecommendationType,
		Team:               filter.Team,
		Service:            filter.Service,
		Status:             filter.Status,
		Limit:              200,
	})
	if err != nil {
		return recommendation{}, err
	}
	for _, item := range recommendations {
		if item.RecommendationID == recommendationID {
			return item, nil
		}
	}
	return recommendation{}, errIncidentNotFound
}

func (s server) buildRecommendationCandidates(ctx context.Context, incidents []investigationIncident, filter recommendationFilter) ([]recommendation, error) {
	selectedIncidentIDs := filter.IncidentIDs
	packageIncidentIDs := filter.PackageIncidentIDs
	incidentScope := selectIncidentsByID(incidents, selectedIncidentIDs)
	packageScope := selectIncidentsByID(incidents, packageIncidentIDs)

	candidates := make([]recommendation, 0, 12)
	if len(packageIncidentIDs) > 0 {
		candidates = append(candidates, buildPackageRecommendations(packageScope, filter)...)
		return candidates, nil
	}
	if len(selectedIncidentIDs) > 0 {
		candidates = append(candidates, buildIncidentRecommendations(incidentScope, incidents, filter)...)
		return candidates, nil
	}

	topIncidents := incidents[:minInt(len(incidents), 5)]
	candidates = append(candidates, buildIncidentRecommendations(topIncidents, incidents, filter)...)
	candidates = append(candidates, buildSystemicRecommendations(incidents, filter)...)
	if topologyRecommendations, err := s.buildTopologyRecommendations(ctx, incidents, filter); err == nil {
		candidates = append(candidates, topologyRecommendations...)
	}
	if forensicRecommendations, err := s.buildForensicsRecommendations(ctx, incidents, filter); err == nil {
		candidates = append(candidates, forensicRecommendations...)
	}
	if runtimeRecommendations, err := s.buildRuntimeRecommendations(ctx, incidents, filter); err == nil {
		candidates = append(candidates, runtimeRecommendations...)
	}
	if hardeningRecommendations, err := s.buildHardeningRecommendations(ctx, incidents, filter); err == nil {
		candidates = append(candidates, hardeningRecommendations...)
	}
	if federationRecommendations, err := s.buildFederationRecommendations(ctx, incidents, filter); err == nil {
		candidates = append(candidates, federationRecommendations...)
	}
	if validationRecommendations, err := s.buildValidationRecommendations(ctx, incidents, filter); err == nil {
		candidates = append(candidates, validationRecommendations...)
	}

	anomalyContext, err := s.buildRecommendationAnomalyContext(ctx, filter)
	if err == nil {
		candidates = applyAnomalyContextToRecommendations(candidates, anomalyContext.Items)
		if len(candidates) == 0 && len(anomalyContext.Items) > 0 {
			candidates = append(candidates, buildAnomalyRecommendations(anomalyContext, incidents, filter)...)
		}
	}
	return candidates, nil
}

func buildIncidentRecommendations(selected []investigationIncident, all []investigationIncident, filter recommendationFilter) []recommendation {
	if len(selected) == 0 {
		return nil
	}
	recommendations := make([]recommendation, 0, len(selected))
	for _, incident := range selected {
		defenseAssessment := attachDefenseGapReadback(buildIncidentDefenseGapAssessment(incident, all), filter.toIncidentFilter())
		replayAssessment := attachPolicyReplayReadback(buildIncidentPolicyReplayAssessment(incident, all), filter.toIncidentFilter())
		topGap := firstDefenseGap(defenseAssessment.DefenseGaps)
		template := recommendationTemplateCatalog[2]
		recommendationType := "remediation"
		if topGap.GapType == "containment_gap" || incident.Status == "active" && incident.State == incidentStateReopened {
			template = recommendationTemplateCatalog[1]
			recommendationType = "investigation"
		}
		readbackRefs := []advisoryReadbackRef{}
		if defenseAssessment.Readback.ResourceID != "" {
			readbackRefs = append(readbackRefs, defenseAssessment.Readback)
		}
		if replayAssessment.Readback.ResourceID != "" {
			readbackRefs = append(readbackRefs, replayAssessment.Readback)
		}
		title := fmt.Sprintf("Reduce %s pressure for %s", strings.ReplaceAll(topGap.GapType, "_", " "), incident.ID)
		if topGap.Title != "" {
			title = topGap.Title
		}
		description := firstNonEmpty(incident.Summary, incident.LikelyCause, incident.RecommendedAction)
		recommendedAction := firstString(append(append(append(append([]string{}, topGap.RecommendedActions.Hardening...), topGap.RecommendedActions.Containment...), topGap.RecommendedActions.GovernanceFix...), incident.RecommendedAction))
		recommendations = append(recommendations, recommendation{
			RecommendationID:    recommendationID("incident", incident.ID, template.TemplateID),
			SourceType:          "incident",
			SourceRef:           incident.ID,
			SubjectType:         "incident",
			SubjectRef:          incident.ID,
			Team:                incident.TenantID,
			Service:             firstString(append(append([]string{incident.ScopeRef}, incident.AffectedWorkloads...), incident.AffectedRepos...)),
			Repo:                incident.Repository,
			Environment:         incident.Environment,
			RecommendationType:  recommendationType,
			Title:               title,
			Description:         description,
			RecommendedAction:   recommendedAction,
			Rationale:           strings.TrimSpace(fmt.Sprintf("%s %s", topGap.WhyItMatters, firstNonEmpty(replayAssessment.ReplayResults[0].Delta, ""))),
			EvidenceRefs:        limitStrings(uniqueStrings(append(append([]string{}, topGap.EvidenceRefs...), incident.EvidenceRefs...)), 10),
			ReadbackRefs:        readbackRefs,
			RelatedIncidentRefs: []string{incident.ID},
			PriorityBand:        recommendationPriorityBand(incident.Priority),
			ImpactScore:         recommendationImpactScore(incident.Severity, len(incident.AffectedWorkloads), len(incident.AffectedEnvironments)),
			EffortScore:         recommendationEffortScore(template.TemplateID),
			ConfidenceScore:     recommendationConfidenceScore(topGap.Confidence),
			ApprovalMode:        template.ApprovalMode,
			Status:              recommendationStatusShown,
			CreatedAt:           recommendationCreatedAt(incident.UpdatedAt, incident.LastActivityAt, incident.OpenedAt),
			ExpiresAt:           recommendationExpiry(incident.UpdatedAt),
			VerificationPlan: []string{
				"Confirm the next deployment or runtime review no longer reproduces the same deny, replay delta, or defense-gap path.",
				"Verify the linked incident either resolves cleanly or moves out of active pressure without widening exception scope.",
			},
			ActionTemplate: template,
			Outcome: recommendationOutcome{
				Status: recommendationStatusShown,
			},
			AdvisoryOnly: true,
			Limitations: []string{
				"Recommendation workflow is an operator overlay and does not mutate canonical incident, lifecycle, evidence, or report truth.",
				"Recommended action remains evidence-backed and advisory until a human accepts or executes the workflow step.",
			},
		})
	}
	return recommendations
}

func buildPackageRecommendations(selected []investigationIncident, filter recommendationFilter) []recommendation {
	if len(selected) == 0 {
		return nil
	}
	pkg := buildIncidentPackage(selected, incidentIDs(selected), filter.toIncidentFilter(), incidentAudienceInternal)
	scopeReplay := attachPolicyReplayReadback(buildScopePolicyReplayAssessment(selected), filter.toIncidentFilter())
	weaknessScope := attachSystemicWeaknessReadback(buildSystemicWeaknessResponse(selected, pkg.SelectionSummary), filter.toIncidentFilter())
	readbackRefs := []advisoryReadbackRef{}
	if scopeReplay.Readback.ResourceID != "" {
		readbackRefs = append(readbackRefs, scopeReplay.Readback)
	}
	if len(weaknessScope.Weaknesses) > 0 && weaknessScope.Weaknesses[0].Readback.ResourceID != "" {
		readbackRefs = append(readbackRefs, weaknessScope.Weaknesses[0].Readback)
	}
	evidenceRefs := packageEvidenceRefs(pkg.PackageIntel)
	template := recommendationTemplateCatalog[0]
	return []recommendation{{
		RecommendationID:    recommendationID("package", packageRecommendationSourceRef(pkg.IncidentRefs), template.TemplateID),
		SourceType:          "package",
		SourceRef:           packageRecommendationSourceRef(pkg.IncidentRefs),
		SubjectType:         "package",
		SubjectRef:          packageRecommendationSourceRef(pkg.IncidentRefs),
		Team:                firstPackageTenant(selected),
		Service:             pkg.SelectionSummary,
		Repo:                firstPackageRepo(selected),
		Environment:         firstPackageEnvironment(selected),
		RecommendationType:  "workflow",
		Title:               "Open a package-level hardening work item",
		Description:         pkg.PackageSummary,
		RecommendedAction:   firstString(append(append(append([]string{}, pkg.PackageIntel.RecommendedActions.ImmediateContainment...), pkg.PackageIntel.RecommendedActions.NearTermHardening...), pkg.PackageIntel.RecommendedActions.GovernanceFix...)),
		Rationale:           strings.TrimSpace(fmt.Sprintf("%s %s", pkg.PackageIntel.DefenseGapSummary.Rationale, pkg.PackageIntel.PolicyReplaySummary.ShadowModeImpact)),
		EvidenceRefs:        evidenceRefs,
		ReadbackRefs:        readbackRefs,
		RelatedIncidentRefs: pkg.IncidentRefs,
		PriorityBand:        packageRecommendationPriority(pkg.PackageIntel),
		ImpactScore:         minInt(95, 45+pkg.IncidentCount*8),
		EffortScore:         recommendationEffortScore(template.TemplateID),
		ConfidenceScore:     packageRecommendationConfidence(pkg.PackageIntel),
		ApprovalMode:        template.ApprovalMode,
		Status:              recommendationStatusShown,
		CreatedAt:           recommendationCreatedAt(timePointer(pkg.GeneratedAt), nil, nil),
		ExpiresAt:           recommendationExpiry(timePointer(pkg.GeneratedAt)),
		VerificationPlan: []string{
			"Verify that the related incident set shrinks or resolves after the package-level remediation path is executed.",
			"Confirm package defense-gap pressure and replay delta both move down in the same scoped bundle.",
		},
		ActionTemplate: template,
		Outcome:        recommendationOutcome{Status: recommendationStatusShown},
		AdvisoryOnly:   true,
		Limitations: []string{
			"Package recommendation remains query-derived from the included incident IDs and current package intelligence summary.",
			"Package workflow state does not replace the canonical case history or export lineage for any included incident.",
		},
	}}
}

func buildSystemicRecommendations(incidents []investigationIncident, filter recommendationFilter) []recommendation {
	response := attachSystemicWeaknessReadback(buildSystemicWeaknessResponse(incidents, "current filtered scope"), filter.toIncidentFilter())
	if len(response.Weaknesses) == 0 {
		return nil
	}
	recommendations := make([]recommendation, 0, minInt(len(response.Weaknesses), 3))
	for _, weakness := range response.Weaknesses[:minInt(len(response.Weaknesses), 3)] {
		template := recommendationTemplateCatalog[0]
		if len(weakness.ProcessFragility) > 0 {
			template = recommendationTemplateCatalog[6]
		}
		recommendations = append(recommendations, recommendation{
			RecommendationID:    recommendationID("systemic", weakness.PatternKey, template.TemplateID),
			SourceType:          "systemic_weakness",
			SourceRef:           weakness.PatternKey,
			SubjectType:         "cluster",
			SubjectRef:          weakness.PatternKey,
			Team:                firstIncidentTenant(incidents),
			Service:             "current filtered scope",
			Repo:                firstIncidentRepo(incidents),
			Environment:         firstIncidentEnvironment(incidents),
			RecommendationType:  "governance",
			Title:               weakness.Title,
			Description:         weakness.Summary,
			RecommendedAction:   weakness.ExecutiveRecommendation,
			Rationale:           weakness.RootCauseHypothesis,
			EvidenceRefs:        weakness.EvidenceRefs,
			ReadbackRefs:        []advisoryReadbackRef{weakness.Readback},
			RelatedIncidentRefs: weakness.RelatedIncidentRefs,
			PriorityBand:        recommendationPriorityBand(weakness.Priority),
			ImpactScore:         systemicImpactScore(weakness),
			EffortScore:         recommendationEffortScore(template.TemplateID),
			ConfidenceScore:     78,
			ApprovalMode:        template.ApprovalMode,
			Status:              recommendationStatusShown,
			CreatedAt:           recommendationCreatedAt(timePointer(response.GeneratedAt), nil, nil),
			ExpiresAt:           recommendationExpiry(timePointer(response.GeneratedAt)),
			VerificationPlan: []string{
				"Track whether the same weakness pattern disappears or narrows in the next systemic weakness refresh.",
				"Confirm the linked incidents stop clustering under the same root-cause hypothesis before closing the workflow.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations:    append([]string{}, weakness.Limitations...),
		})
	}
	return recommendations
}

func buildAnomalyRecommendations(response audit.AnalyticsAnomaliesResponse, incidents []investigationIncident, filter recommendationFilter) []recommendation {
	if len(response.Items) == 0 {
		return nil
	}
	scopeReplay := attachPolicyReplayReadback(buildScopePolicyReplayAssessment(incidents), filter.toIncidentFilter())
	readbackRefs := []advisoryReadbackRef{}
	if scopeReplay.Readback.ResourceID != "" {
		readbackRefs = append(readbackRefs, scopeReplay.Readback)
	}
	recommendations := make([]recommendation, 0, minInt(len(response.Items), 2))
	for _, item := range response.Items[:minInt(len(response.Items), 2)] {
		template := recommendationTemplateCatalog[1]
		recommendations = append(recommendations, recommendation{
			RecommendationID:   recommendationID("anomaly", item.Type+":"+item.Segment, template.TemplateID),
			SourceType:         "anomaly",
			SourceRef:          item.Type + ":" + item.Segment,
			SubjectType:        "segment",
			SubjectRef:         item.Segment,
			Team:               firstIncidentTenant(incidents),
			Service:            item.Segment,
			Repo:               filter.event.Repo,
			Environment:        filter.event.Environment,
			RecommendationType: "investigation",
			Title:              item.Title,
			Description:        item.Reason,
			RecommendedAction:  item.RecommendedNextStep,
			Rationale:          fmt.Sprintf("%s %s", item.Baseline, item.Deviation),
			EvidenceRefs:       item.EvidenceRefs,
			ReadbackRefs:       readbackRefs,
			PriorityBand:       recommendationPriorityBand(item.Severity),
			ImpactScore:        anomalyImpactScore(item.Severity),
			EffortScore:        recommendationEffortScore(template.TemplateID),
			ConfidenceScore:    72,
			ApprovalMode:       template.ApprovalMode,
			Status:             recommendationStatusShown,
			CreatedAt:          recommendationCreatedAt(timePointer(response.Comparison.CurrentEnd), nil, nil),
			ExpiresAt:          recommendationExpiry(timePointer(response.Comparison.CurrentEnd)),
			VerificationPlan: []string{
				"Re-run the same analytics window and confirm the anomaly no longer crosses the explainable threshold.",
				"Check that the underlying friction, exception, signer, or drift pressure moves back toward baseline.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations:    append([]string{}, item.Limitations...),
		})
	}
	return recommendations
}

func (s server) buildTopologyRecommendations(ctx context.Context, incidents []investigationIncident, filter recommendationFilter) ([]recommendation, error) {
	topologyFilter, err := recommendationTopologyFilter(filter)
	if err != nil {
		return nil, err
	}
	delta, err := s.buildTopologyDeltaResponse(ctx, topologyFilter)
	if err != nil {
		return nil, err
	}
	recommendations := make([]recommendation, 0, minInt(len(delta.Items), 2))
	for _, item := range delta.Items {
		if item.Delta <= 0 && len(item.DriftSignals) == 0 {
			continue
		}
		serviceFilter := topologyFilter
		serviceFilter.Service = item.Service
		blast, err := s.buildTopologyBlastRadiusForService(ctx, serviceFilter)
		if err != nil {
			continue
		}
		related := incidentsForTopologyService(incidents, item.Service)
		template := topologyRecommendationTemplate(item, blast)
		recommendations = append(recommendations, recommendation{
			RecommendationID:    recommendationID("topology", item.Service, template.TemplateID),
			SourceType:          "topology_signal",
			SourceRef:           item.Service,
			SubjectType:         "service",
			SubjectRef:          item.Service,
			Team:                firstIncidentTenant(related),
			Service:             item.Service,
			Repo:                firstIncidentRepo(related),
			Environment:         firstIncidentEnvironment(related),
			RecommendationType:  template.RecommendationType,
			Title:               fmt.Sprintf("Reduce expanding blast radius for %s", item.Service),
			Description:         firstNonEmpty(firstString(item.DriftSignals), fmt.Sprintf("Blast radius expanded for %s in the current topology window.", item.Service)),
			RecommendedAction:   topologyRecommendationAction(blast, item),
			Rationale:           topologyRecommendationRationale(item, blast),
			EvidenceRefs:        topologyRecommendationEvidenceRefs(blast, related),
			ReadbackRefs:        topologyRecommendationReadbackRefs(related, incidents, filter),
			RelatedIncidentRefs: incidentIDs(related),
			PriorityBand:        topologyRecommendationPriority(item, blast),
			ImpactScore:         minInt(100, maxInt(48+blast.BlastRadiusScore/2+item.CriticalReachDelta*8, 40)),
			EffortScore:         recommendationEffortScore(template.TemplateID),
			ConfidenceScore:     topologyRecommendationConfidence(blast, item),
			ApprovalMode:        template.ApprovalMode,
			Status:              recommendationStatusShown,
			CreatedAt:           recommendationCreatedAt(timePointer(delta.Comparison.CurrentEnd), nil, nil),
			ExpiresAt:           recommendationExpiry(timePointer(delta.Comparison.CurrentEnd)),
			VerificationPlan: []string{
				"Re-run service-graph blast radius for the same service and confirm the effective blast-radius score drops.",
				"Confirm topology delta no longer shows blast-radius expansion, public exposure widening, or new critical downstream reach for this service.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations: uniqueStrings(append([]string{
				"Topology recommendation remains a workflow overlay derived from the advisory 9e service graph and does not mutate canonical incident or runtime truth.",
			}, append(delta.Limitations, blast.Limitations...)...)),
		})
		if len(recommendations) == 2 {
			break
		}
	}
	return recommendations, nil
}

func (s server) buildForensicsRecommendations(ctx context.Context, incidents []investigationIncident, filter recommendationFilter) ([]recommendation, error) {
	forensicFilter, err := recommendationForensicsFilter(filter, incidents)
	if err != nil {
		return nil, err
	}
	delta, err := s.buildForensicsDeltaResponse(ctx, forensicFilter)
	if err != nil {
		return nil, err
	}
	historicalFilter := withForensicsTimestamp(forensicFilter, delta.Comparison.T1)
	replay, err := s.buildForensicsReplay(ctx, historicalFilter, forensicsReplayModernFullStack)
	if err != nil {
		return nil, err
	}
	if replay.VerdictDelta == "no_change" && len(delta.TopologyDelta) == 0 && len(delta.IdentityDelta.Modified) == 0 && len(delta.VulnerabilityDelta.Added) == 0 {
		return nil, nil
	}

	related := incidents
	subjectType := "scope"
	subjectRef := forensicSubjectSummary(forensicFilter)
	if forensicFilter.IncidentID != "" {
		subjectType = "incident"
		subjectRef = forensicFilter.IncidentID
		related = selectIncidentsByID(incidents, []string{forensicFilter.IncidentID})
	} else if forensicFilter.Service != "" {
		subjectType = "service"
		subjectRef = forensicFilter.Service
		related = incidentsForTopologyService(incidents, forensicFilter.Service)
	}

	template := recommendationTemplateCatalog[1]
	if replay.VerdictDelta != "no_change" {
		template = recommendationTemplateCatalog[2]
	}
	if len(delta.TopologyDelta) > 0 || len(delta.IdentityDelta.Modified) > 0 {
		template = recommendationTemplateCatalog[6]
	}

	title := fmt.Sprintf("Review historical control drift for %s", subjectRef)
	description := fmt.Sprintf(
		"Historical state at %s reconstructed as %s, while modern replay evaluates as %s.",
		delta.Comparison.T1.Format(time.RFC3339),
		strings.ToLower(replay.HistoricalVerdict),
		strings.ToLower(replay.ReplayVerdict),
	)
	recommendedAction := firstNonEmpty(
		firstString(replay.Explanations),
		"Replay the historical state in a bounded sandbox and route follow-up remediation before widening any exceptions.",
	)
	rationale := forensicRecommendationRationale(replay, delta)
	evidenceRefs := uniqueStrings(append(append([]string{}, replay.EvidenceRefs...), delta.EvidenceRefs...))
	priorityBand := "TODAY"
	if replay.VerdictDelta != "no_change" || len(delta.VulnerabilityDelta.Added) > 0 {
		priorityBand = "NOW"
	}
	confidence := 72
	if len(delta.TopologyDelta) > 0 {
		confidence += 6
	}
	if len(replay.ReadbackRefs) > 0 {
		confidence += 4
	}

	return []recommendation{{
		RecommendationID:    recommendationID("forensic", subjectRef, template.TemplateID),
		SourceType:          "forensic_signal",
		SourceRef:           forensicRecommendationSourceRef(subjectRef, delta.Comparison.T1),
		SubjectType:         subjectType,
		SubjectRef:          subjectRef,
		Team:                firstIncidentTenant(related),
		Service:             firstNonEmpty(forensicFilter.Service, subjectRef),
		Repo:                firstIncidentRepo(related),
		Environment:         firstNonEmpty(filter.event.Environment, firstIncidentEnvironment(related)),
		RecommendationType:  template.RecommendationType,
		Title:               title,
		Description:         description,
		RecommendedAction:   recommendedAction,
		Rationale:           rationale,
		EvidenceRefs:        evidenceRefs,
		ReadbackRefs:        replay.ReadbackRefs,
		RelatedIncidentRefs: incidentIDs(related),
		PriorityBand:        priorityBand,
		ImpactScore:         minInt(96, 52+len(delta.VulnerabilityDelta.Added)*8+len(delta.TopologyDelta)*10),
		EffortScore:         recommendationEffortScore(template.TemplateID),
		ConfidenceScore:     minInt(90, confidence),
		ApprovalMode:        template.ApprovalMode,
		Status:              recommendationStatusShown,
		CreatedAt:           recommendationCreatedAt(timePointer(delta.Comparison.T2), nil, nil),
		ExpiresAt:           recommendationExpiry(timePointer(delta.Comparison.T2)),
		VerificationPlan: []string{
			fmt.Sprintf("Re-run forensics replay for %s and confirm the replay verdict no longer diverges from the historical verdict at %s.", subjectRef, delta.Comparison.T1.Format(time.RFC3339)),
			"Confirm historical delta pressure narrows: fewer later disclosures, identity drifts, or topology changes should remain in the same compare window.",
		},
		ActionTemplate: template,
		Outcome:        recommendationOutcome{Status: recommendationStatusShown},
		AdvisoryOnly:   true,
		Limitations: uniqueStrings(append(append([]string{
			"Forensic recommendation remains an overlay derived from reconstructed historical state and counterfactual replay; it does not mutate canonical incident or evidence truth.",
		}, replay.Limitations...), delta.Limitations...)),
	}}, nil
}

func (s server) buildRuntimeRecommendations(ctx context.Context, incidents []investigationIncident, filter recommendationFilter) ([]recommendation, error) {
	runtimeFilter := recommendationRuntimeFilter(filter)
	findings, _, err := s.buildRuntimeFindings(ctx, runtimeFilter)
	if err != nil {
		return nil, err
	}
	if len(findings) == 0 {
		return nil, nil
	}
	workloads, _, err := s.buildRuntimeWorkloads(ctx, runtimeFilter)
	if err != nil {
		return nil, err
	}
	workloadBySubject := map[string]runtimeWorkloadView{}
	for _, item := range workloads {
		workloadBySubject[item.SubjectRef] = item
	}
	recommendations := make([]recommendation, 0, minInt(len(findings), 3))
	for _, finding := range findings {
		if finding.Status == runtimeFindingStatusRemediated {
			continue
		}
		workload := workloadBySubject[finding.SubjectRef]
		decision, err := s.evaluateRuntimeEnforcement(ctx, runtimeFilter, runtimeActionRequest{
			FindingID:  finding.FindingID,
			SubjectRef: finding.SubjectRef,
		}, "")
		if err != nil {
			continue
		}
		related := incidentsForRuntimeSubject(incidents, finding.SubjectRef)
		template := runtimeRecommendationTemplate(finding, decision)
		workloadName := runtimeRecommendationSubjectName(finding.SubjectRef)
		title := fmt.Sprintf("Contain %s on %s", strings.ReplaceAll(finding.FindingType, "_", " "), workloadName)
		if finding.Severity == "critical" {
			title = fmt.Sprintf("Review critical runtime drift on %s", workloadName)
		}
		description := finding.Summary
		if workload.State.CurrentSandboxClass != "" {
			description = fmt.Sprintf("%s Current sandbox class is %s.", finding.Summary, workload.State.CurrentSandboxClass)
		}
		rationale := finding.Summary
		if decision.TopologyContext != nil && decision.TopologyContext.BlastRadiusScore > 0 {
			rationale = fmt.Sprintf(
				"%s Containment affects service %s with blast radius %d and %d critical downstream reach.",
				finding.Summary,
				firstNonEmpty(decision.TopologyContext.PrimaryService, workloadName),
				decision.TopologyContext.BlastRadiusScore,
				decision.TopologyContext.CriticalReachCount,
			)
		}
		recommendations = append(recommendations, recommendation{
			RecommendationID:    recommendationID("runtime", finding.SubjectRef, finding.FindingType),
			SourceType:          "runtime_signal",
			SourceRef:           finding.FindingID,
			SubjectType:         "workload",
			SubjectRef:          finding.SubjectRef,
			Team:                firstNonEmpty(firstIncidentTenant(related), filter.event.TenantID),
			Service:             workloadName,
			Repo:                firstNonEmpty(firstString(workload.Profile.ProfileSource), firstIncidentRepo(related), filter.event.Repo),
			Environment:         firstNonEmpty(workload.Environment, firstIncidentEnvironment(related), filter.event.Environment),
			RecommendationType:  template.RecommendationType,
			Title:               title,
			Description:         description,
			RecommendedAction:   decision.Action,
			Rationale:           rationale,
			EvidenceRefs:        uniqueStrings(append(append([]string{}, finding.EvidenceRefs...), workload.State.EvidenceRefs...)),
			ReadbackRefs:        finding.ReadbackRefs,
			RelatedIncidentRefs: incidentIDs(related),
			PriorityBand:        runtimeRecommendationPriority(finding),
			ImpactScore:         runtimeRecommendationImpact(finding, decision),
			EffortScore:         recommendationEffortScore(template.TemplateID),
			ConfidenceScore:     runtimeRecommendationConfidence(finding),
			ApprovalMode:        decision.ApprovalMode,
			Status:              recommendationStatusShown,
			CreatedAt:           recommendationCreatedAt(timePointer(workload.State.LastVerifiedAt), nil, nil),
			ExpiresAt:           recommendationExpiry(timePointer(workload.State.LastVerifiedAt)),
			VerificationPlan: []string{
				fmt.Sprintf("Re-run runtime integrity for %s and confirm %s is no longer active for the same subject.", workloadName, finding.FindingType),
				"Confirm the workload sandbox class, SBOM verification, and topology-aware containment posture all move back toward the expected state without reopening drift.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations: uniqueStrings(append([]string{
				"Runtime recommendation remains an overlay derived from canonical runtime observations, profile evaluation, and policy-gated containment logic; it does not itself execute containment.",
			}, append(finding.Limitations, decision.Limitations...)...)),
		})
		if len(recommendations) == 3 {
			break
		}
	}
	return recommendations, nil
}

func (s server) buildHardeningRecommendations(ctx context.Context, incidents []investigationIncident, filter recommendationFilter) ([]recommendation, error) {
	runtimeFilter := recommendationRuntimeFilter(filter)
	executions, _, err := s.listHardeningExecutions(ctx, runtimeFilter)
	if err != nil {
		return nil, err
	}
	if len(executions) == 0 {
		return nil, nil
	}
	posture, _, err := s.buildDefensePostureStates(ctx, runtimeFilter)
	if err != nil {
		return nil, err
	}
	postureBySubject := map[string]defensePostureState{}
	for _, item := range posture {
		postureBySubject[item.SubjectRef] = item
	}
	findings, _, err := s.buildRuntimeFindings(ctx, runtimeFilter)
	if err != nil {
		return nil, err
	}
	findingByID := map[string]runtimeIntegrityFinding{}
	for _, item := range findings {
		findingByID[item.FindingID] = item
	}
	recommendations := make([]recommendation, 0, minInt(len(executions), 3))
	for _, execution := range executions {
		if execution.ExecutionResult == "rollback_applied" || execution.ExecutionResult == "trusted_recovery_completed" {
			continue
		}
		postureState := postureBySubject[execution.SubjectRef]
		finding, hasFinding := findingByID[execution.TriggerRef]
		template := recommendationTemplateCatalog[2]
		if strings.Contains(execution.ExecutionResult, "pending") {
			template = recommendationTemplateCatalog[6]
		} else if containsHardeningActionType(execution.ActionsApplied, hardeningActionRequestForensics) {
			template = recommendationTemplateCatalog[1]
		}
		workloadName := runtimeRecommendationSubjectName(execution.SubjectRef)
		title := fmt.Sprintf("Convert temporary hardening on %s into a durable remediation plan", workloadName)
		if strings.Contains(execution.ExecutionResult, "pending") {
			title = fmt.Sprintf("Review pending runtime hardening approval for %s", workloadName)
		}
		rationale := "Temporary autonomous hardening reduced immediate runtime pressure, but permanent remediation still belongs in a bounded workflow."
		if hasFinding {
			rationale = fmt.Sprintf("%s Temporary hardening left %s in %s mode, which now needs a durable follow-up in workflow rather than indefinite autonomous posture.", finding.Summary, workloadName, postureState.CurrentMode)
		}
		priority := "TODAY"
		if hasFinding && finding.Severity == "critical" {
			priority = "NOW"
		}
		recommendations = append(recommendations, recommendation{
			RecommendationID:    hardeningRecommendationID(execution.SubjectRef, execution.ExecutionID),
			SourceType:          "hardening_signal",
			SourceRef:           execution.ExecutionID,
			SubjectType:         "workload",
			SubjectRef:          execution.SubjectRef,
			Team:                firstIncidentTenant(incidentsForRuntimeSubject(incidents, execution.SubjectRef)),
			Service:             workloadName,
			Repo:                firstIncidentRepo(incidentsForRuntimeSubject(incidents, execution.SubjectRef)),
			Environment:         firstIncidentEnvironment(incidentsForRuntimeSubject(incidents, execution.SubjectRef)),
			RecommendationType:  template.RecommendationType,
			Title:               title,
			Description:         firstNonEmpty(postureState.TriggerSummary, execution.ExecutionResult),
			RecommendedAction:   firstNonEmpty(firstHardeningActionType(execution.ActionsApplied), "generate_remediation_draft"),
			Rationale:           rationale,
			EvidenceRefs:        uniqueStrings(append(append([]string{}, execution.ForensicRefs...), execution.IncidentRefs...)),
			ReadbackRefs:        nil,
			RelatedIncidentRefs: uniqueStrings(append([]string{}, execution.IncidentRefs...)),
			PriorityBand:        priority,
			ImpactScore:         minInt(100, 58+len(execution.ActionsApplied)*8+len(execution.IncidentRefs)*5),
			EffortScore:         recommendationEffortScore(template.TemplateID),
			ConfidenceScore:     74,
			ApprovalMode:        template.ApprovalMode,
			Status:              recommendationStatusShown,
			CreatedAt:           execution.ExecutedAt,
			ExpiresAt:           execution.ExpiresAt,
			VerificationPlan: []string{
				fmt.Sprintf("Confirm %s no longer re-triggers the same runtime hardening path for %s in the current scope.", firstNonEmpty(execution.TriggerRef, "the linked finding"), workloadName),
				"Verify temporary restrictions can be removed without reopening runtime drift, then close the hardening follow-up as permanent remediation.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations: []string{
				"Hardening recommendation is an overlay derived from bounded 9k temporary response records; it does not treat temporary autonomous containment as permanent remediation truth.",
			},
		})
		if len(recommendations) == 3 {
			break
		}
	}
	return recommendations, nil
}

func (s server) buildValidationRecommendations(ctx context.Context, incidents []investigationIncident, filter recommendationFilter) ([]recommendation, error) {
	validationFilter := validationHarnessFilter{
		ClusterID:   filter.event.ClusterID,
		TenantID:    filter.event.TenantID,
		Environment: filter.event.Environment,
		Repo:        filter.event.Repo,
		Service:     filter.Service,
		Limit:       maxInt(filter.Limit, 6),
		event: audit.EventFilter{
			ClusterID:   filter.event.ClusterID,
			TenantID:    filter.event.TenantID,
			Environment: filter.event.Environment,
			Repo:        filter.event.Repo,
			Limit:       maxInt(filter.Limit*20, 500),
		},
	}
	score, err := s.buildValidationHarnessScore(ctx, validationFilter)
	if err != nil {
		return nil, err
	}
	scenarioTimes := map[string]*time.Time{}
	if runs, _, err := s.listStrictValidationRuns(ctx, validationFilter); err == nil {
		for _, run := range runs {
			issuedAt := run.Certificate.IssuedAt.UTC()
			for _, verdict := range run.Verdicts {
				if _, exists := scenarioTimes[verdict.ScenarioID]; exists {
					continue
				}
				timestamp := issuedAt
				scenarioTimes[verdict.ScenarioID] = &timestamp
			}
		}
	}
	if score.FailedScenarios == 0 && score.PartialScenarios == 0 {
		return nil, nil
	}
	recommendations := []recommendation{}
	for _, result := range score.Results {
		if result.Status == validationStatusPass {
			continue
		}
		template := recommendationTemplateCatalog[2]
		if result.Status == validationStatusFail {
			template = recommendationTemplateCatalog[6]
		}
		if strings.Contains(result.ScenarioID, "containment") {
			template = recommendationTemplateCatalog[1]
		}
		title := fmt.Sprintf("Close validation gap for %s", strings.ReplaceAll(result.ScenarioID, "_", " "))
		if result.Status == validationStatusFail {
			title = fmt.Sprintf("Repair failed validation scenario: %s", strings.ReplaceAll(result.ScenarioID, "_", " "))
		}
		recommendations = append(recommendations, recommendation{
			RecommendationID:   recommendationID("validation", result.ScenarioID, template.TemplateID),
			SourceType:         "validation_signal",
			SourceRef:          result.ScenarioID,
			SubjectType:        "validation_harness",
			SubjectRef:         firstNonEmpty(validationFilter.Service, validationScopeSummary(validationFilter)),
			Team:               firstIncidentTenant(incidents),
			Service:            firstNonEmpty(validationFilter.Service, filter.Service),
			Repo:               firstNonEmpty(filter.event.Repo, firstIncidentRepo(incidents)),
			Environment:        firstNonEmpty(filter.event.Environment, firstIncidentEnvironment(incidents)),
			RecommendationType: template.RecommendationType,
			Title:              title,
			Description:        result.Summary,
			RecommendedAction:  firstNonEmpty(firstString(result.TriggeredControls), template.TemplateID),
			Rationale: fmt.Sprintf(
				"Validation harness currently reports %s for scenario %s. The gap should be closed before treating the control path as verified.",
				result.Status,
				result.ScenarioID,
			),
			EvidenceRefs:        uniqueStrings(result.EvidenceRefs),
			ReadbackRefs:        append([]advisoryReadbackRef(nil), result.ReadbackRefs...),
			RelatedIncidentRefs: incidentIDs(incidents),
			PriorityBand:        mapValidationRecommendationPriority(result.Status),
			ImpactScore:         minInt(95, 50+len(result.EvidenceRefs)*4+score.FailedScenarios*8),
			EffortScore:         recommendationEffortScore(template.TemplateID),
			ConfidenceScore:     mapValidationRecommendationConfidence(result.Status, score.ConfidenceLevel),
			ApprovalMode:        template.ApprovalMode,
			Status:              recommendationStatusShown,
			CreatedAt:           recommendationCreatedAt(scenarioTimes[result.ScenarioID]),
			ExpiresAt:           recommendationExpiry(scenarioTimes[result.ScenarioID]),
			VerificationPlan: []string{
				fmt.Sprintf("Re-run validation harness scenario %s and confirm it moves out of %s into pass.", result.ScenarioID, result.Status),
				"Confirm the corresponding runtime, topology, forensics, or policy evidence path no longer appears as a validation gap in the current scope.",
			},
			ActionTemplate: template,
			Outcome:        recommendationOutcome{Status: recommendationStatusShown},
			AdvisoryOnly:   true,
			Limitations: uniqueStrings(append([]string{
				"Validation recommendation is an overlay derived from 9j dry-run results and does not mutate canonical incident, evidence, or runtime truth.",
			}, append(score.Limitations, result.Limitations...)...)),
		})
		if len(recommendations) == 2 {
			break
		}
	}
	return recommendations, nil
}

func recommendationRuntimeFilter(filter recommendationFilter) runtimeIntegrityFilter {
	limit := maxInt(filter.Limit, 25)
	workload := strings.TrimSpace(filter.Service)
	return runtimeIntegrityFilter{
		ClusterID:   filter.event.ClusterID,
		TenantID:    filter.event.TenantID,
		Environment: filter.event.Environment,
		Repo:        filter.event.Repo,
		Workload:    workload,
		Limit:       limit,
		event: audit.EventFilter{
			ClusterID:   filter.event.ClusterID,
			TenantID:    filter.event.TenantID,
			Environment: filter.event.Environment,
			Repo:        filter.event.Repo,
			Limit:       maxInt(limit*8, 500),
		},
	}
}

func incidentsForRuntimeSubject(incidents []investigationIncident, subjectRef string) []investigationIncident {
	_, namespace, _, workload, err := parseRuntimeSubjectRef(subjectRef)
	if err != nil {
		return nil
	}
	matches := make([]investigationIncident, 0, len(incidents))
	for _, incident := range incidents {
		if strings.EqualFold(strings.TrimSpace(incident.ScopeRef), workload) ||
			containsString(incident.AffectedWorkloads, workload) ||
			containsString(incident.AffectedNamespaces, namespace) {
			matches = append(matches, incident)
		}
	}
	return matches
}

func runtimeRecommendationTemplate(finding runtimeIntegrityFinding, decision runtimeEnforcementDecision) recommendationActionTemplate {
	switch {
	case decision.ApprovalMode == recommendationApprovalHumanReview:
		return recommendationTemplateCatalog[6]
	case decision.Action == runtimeActionCaptureForensics:
		return recommendationTemplateCatalog[1]
	case decision.Action == runtimeActionAlert:
		return recommendationTemplateCatalog[5]
	default:
		return recommendationTemplateCatalog[2]
	}
}

func runtimeRecommendationPriority(finding runtimeIntegrityFinding) string {
	switch finding.Severity {
	case "critical":
		return "NOW"
	case "high":
		return "TODAY"
	case "medium":
		return "THIS_WEEK"
	default:
		return "BACKLOG"
	}
}

func runtimeRecommendationImpact(finding runtimeIntegrityFinding, decision runtimeEnforcementDecision) int {
	score := 40 + runtimeSeverityRank(finding.Severity)*12
	if decision.TopologyContext != nil {
		score += minInt(24, decision.TopologyContext.BlastRadiusScore/4+decision.TopologyContext.CriticalReachCount*6)
	}
	return minInt(100, score)
}

func runtimeRecommendationConfidence(finding runtimeIntegrityFinding) int {
	switch strings.TrimSpace(finding.Confidence) {
	case runtimeConfidenceHigh:
		return 84
	case runtimeConfidenceLow:
		return 58
	default:
		return 72
	}
}

func containsHardeningActionType(actions []hardeningAction, actionType string) bool {
	for _, action := range actions {
		if action.ActionType == actionType {
			return true
		}
	}
	return false
}

func firstHardeningActionType(actions []hardeningAction) string {
	for _, action := range actions {
		if strings.TrimSpace(action.ActionType) != "" {
			return action.ActionType
		}
	}
	return ""
}

func mapValidationRecommendationPriority(status string) string {
	switch status {
	case validationStatusFail:
		return "NOW"
	case validationStatusPartial:
		return "TODAY"
	default:
		return "THIS_WEEK"
	}
}

func mapValidationRecommendationConfidence(status, confidenceLevel string) int {
	score := 68
	switch status {
	case validationStatusFail:
		score += 10
	case validationStatusPass:
		score -= 8
	}
	switch confidenceLevel {
	case "high":
		score += 8
	case "low":
		score -= 6
	}
	return minInt(90, maxInt(score, 45))
}

func runtimeRecommendationSubjectName(subjectRef string) string {
	_, _, _, workload, err := parseRuntimeSubjectRef(subjectRef)
	if err != nil {
		return subjectRef
	}
	return workload
}

func recommendationForensicsFilter(filter recommendationFilter, incidents []investigationIncident) (forensicsFilter, error) {
	analyticsFilter, err := audit.NormalizeAnalyticsFilter(audit.AnalyticsFilter{
		Window:      "28d",
		CompareTo:   "previous_window",
		GroupBy:     "service",
		ClusterID:   filter.event.ClusterID,
		TenantID:    filter.event.TenantID,
		Environment: filter.event.Environment,
		Repo:        filter.event.Repo,
		Service:     filter.Service,
		Team:        filter.Team,
	})
	if err != nil {
		return forensicsFilter{}, err
	}
	incidentID := firstString(filter.IncidentIDs)
	service := strings.TrimSpace(filter.Service)
	if service == "" {
		service = firstIncidentServiceRef(incidents)
	}
	return forensicsFilter{
		event: audit.EventFilter{
			ClusterID:   analyticsFilter.ClusterID,
			TenantID:    analyticsFilter.TenantID,
			Environment: analyticsFilter.Environment,
			Repo:        analyticsFilter.Repo,
			Limit:       forensicsHistoryLimit,
		},
		analytics:  analyticsFilter,
		Timestamp:  time.Now().UTC(),
		IncidentID: incidentID,
		Service:    service,
		Limit:      20,
	}, nil
}

func firstIncidentServiceRef(incidents []investigationIncident) string {
	for _, incident := range incidents {
		if service := firstNonEmpty(incident.ScopeRef, firstString(incident.AffectedWorkloads)); strings.TrimSpace(service) != "" {
			return service
		}
	}
	return ""
}

func forensicRecommendationSourceRef(subjectRef string, timestamp time.Time) string {
	subjectRef = strings.ReplaceAll(strings.TrimSpace(subjectRef), "@", "_")
	return fmt.Sprintf("%s@%d", firstNonEmpty(subjectRef, "forensic-scope"), timestamp.UTC().Unix())
}

func parseForensicRecommendationSourceRef(sourceRef string, fallback time.Time) time.Time {
	sourceRef = strings.TrimSpace(sourceRef)
	parts := strings.Split(sourceRef, "@")
	if len(parts) == 2 {
		if unix := parseIntOrDefault(parts[1], 0); unix > 0 {
			return time.Unix(int64(unix), 0).UTC()
		}
	}
	return fallback.UTC()
}

func forensicRecommendationRationale(replay forensicReplayResponse, delta timeDeltaResult) string {
	parts := []string{
		fmt.Sprintf("Historical verdict %s vs replay verdict %s.", strings.ToLower(replay.HistoricalVerdict), strings.ToLower(replay.ReplayVerdict)),
	}
	if len(replay.Explanations) > 0 {
		parts = append(parts, firstString(replay.Explanations))
	}
	if len(delta.VulnerabilityDelta.Added) > 0 {
		parts = append(parts, fmt.Sprintf("%d later-disclosed vulnerability signal(s) change the historical known-state.", len(delta.VulnerabilityDelta.Added)))
	}
	if len(delta.IdentityDelta.Modified) > 0 {
		parts = append(parts, fmt.Sprintf("%d identity drift signal(s) were added between the compared forensic windows.", len(delta.IdentityDelta.Modified)))
	}
	if len(delta.TopologyDelta) > 0 {
		parts = append(parts, "Topology blast-radius drift is also present in the same forensic comparison window.")
	}
	return strings.TrimSpace(strings.Join(uniqueStrings(parts), " "))
}

func recommendationTopologyFilter(filter recommendationFilter) (topologyFilter, error) {
	analyticsFilter, err := audit.NormalizeAnalyticsFilter(audit.AnalyticsFilter{
		Window:      "28d",
		CompareTo:   "previous_window",
		GroupBy:     "service",
		ClusterID:   filter.event.ClusterID,
		TenantID:    filter.event.TenantID,
		Environment: filter.event.Environment,
		Repo:        filter.event.Repo,
		Service:     filter.Service,
		Team:        filter.Team,
	})
	if err != nil {
		return topologyFilter{}, err
	}
	return topologyFilterFromAnalyticsFilter(analyticsFilter), nil
}

func incidentsForTopologyService(incidents []investigationIncident, service string) []investigationIncident {
	if strings.TrimSpace(service) == "" {
		return nil
	}
	matches := make([]investigationIncident, 0, len(incidents))
	for _, incident := range incidents {
		if incidentMatchesTopologyService(incident, service) {
			matches = append(matches, incident)
		}
	}
	return matches
}

func incidentMatchesTopologyService(incident investigationIncident, service string) bool {
	service = strings.ToLower(strings.TrimSpace(service))
	if service == "" {
		return false
	}
	if strings.ToLower(strings.TrimSpace(incident.ScopeRef)) == service {
		return true
	}
	for _, workload := range incident.AffectedWorkloads {
		if strings.ToLower(strings.TrimSpace(workload)) == service {
			return true
		}
	}
	repository := strings.ToLower(strings.TrimSpace(incident.Repository))
	if repository != "" && (repository == service || strings.Contains(repository, service)) {
		return true
	}
	return false
}

func topologyRecommendationTemplate(item topologyDeltaItem, blast topologyBlastRadiusResponse) recommendationActionTemplate {
	if containsString(item.DriftSignals, "public exposure widened") || item.CriticalReachDelta > 0 || blast.BlastRadiusScore >= 85 {
		return recommendationTemplateCatalog[6]
	}
	if containsString(item.DriftSignals, "new connectivity path detected") || blast.BlastRadiusScore >= 60 {
		return recommendationTemplateCatalog[1]
	}
	return recommendationTemplateCatalog[0]
}

func topologyRecommendationPriority(item topologyDeltaItem, blast topologyBlastRadiusResponse) string {
	switch {
	case containsString(item.DriftSignals, "public exposure widened") || item.CriticalReachDelta > 0 || blast.BlastRadiusScore >= 85:
		return "NOW"
	case item.Delta > 0 || blast.BlastRadiusScore >= 60:
		return "TODAY"
	default:
		return "THIS_WEEK"
	}
}

func topologyRecommendationConfidence(blast topologyBlastRadiusResponse, item topologyDeltaItem) int {
	confidence := 70
	if blast.ObservedEdgeCount > 0 {
		confidence += 8
	}
	if len(item.DriftSignals) > 0 {
		confidence += 6
	}
	return minInt(90, confidence)
}

func topologyRecommendationAction(blast topologyBlastRadiusResponse, item topologyDeltaItem) string {
	if len(blast.ContainmentOptions) > 0 {
		return firstNonEmpty(blast.ContainmentOptions[0].Title, blast.ContainmentOptions[0].Summary)
	}
	if containsString(item.DriftSignals, "public exposure widened") {
		return "Open a security review and simulate a minimal ingress restriction plan for this service."
	}
	return "Open a bounded topology investigation and route a remediation work item to the current owner."
}

func topologyRecommendationRationale(item topologyDeltaItem, blast topologyBlastRadiusResponse) string {
	riskPath := firstString(topologyRiskPathSummaries(blast.TopRiskPaths, 1))
	if riskPath == "" {
		riskPath = "No single topological path summary was available beyond the effective graph score."
	}
	return strings.TrimSpace(fmt.Sprintf(
		"Blast radius expanded from %d to %d for %s. Critical reach changed by %d. %s",
		item.BaselineBlastRadiusScore,
		item.CurrentBlastRadiusScore,
		item.Service,
		item.CriticalReachDelta,
		riskPath,
	))
}

func topologyRecommendationEvidenceRefs(blast topologyBlastRadiusResponse, related []investigationIncident) []string {
	evidenceRefs := append([]string{}, blast.EvidenceRefs...)
	for _, incident := range related {
		evidenceRefs = append(evidenceRefs, incident.EvidenceRefs...)
	}
	return limitStrings(uniqueStrings(evidenceRefs), 12)
}

func topologyRecommendationReadbackRefs(related []investigationIncident, all []investigationIncident, filter recommendationFilter) []advisoryReadbackRef {
	refs := []advisoryReadbackRef{}
	for _, incident := range related[:minInt(len(related), 2)] {
		defenseAssessment := attachDefenseGapReadback(buildIncidentDefenseGapAssessment(incident, all), filter.toIncidentFilter())
		replayAssessment := attachPolicyReplayReadback(buildIncidentPolicyReplayAssessment(incident, all), filter.toIncidentFilter())
		if defenseAssessment.Readback.ResourceID != "" {
			refs = append(refs, defenseAssessment.Readback)
		}
		if replayAssessment.Readback.ResourceID != "" {
			refs = append(refs, replayAssessment.Readback)
		}
	}
	return uniqueAdvisoryReadbackRefs(refs)
}

func uniqueAdvisoryReadbackRefs(refs []advisoryReadbackRef) []advisoryReadbackRef {
	seen := map[string]struct{}{}
	unique := make([]advisoryReadbackRef, 0, len(refs))
	for _, ref := range refs {
		key := ref.ResourceType + ":" + ref.ResourceID
		if ref.ResourceID == "" {
			continue
		}
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		unique = append(unique, ref)
	}
	return unique
}

func topologyRiskPathSummaries(paths []topologyRiskPath, limit int) []string {
	summaries := make([]string, 0, minInt(len(paths), limit))
	for _, path := range paths[:minInt(len(paths), limit)] {
		if summary := strings.TrimSpace(path.Summary); summary != "" {
			summaries = append(summaries, summary)
		}
	}
	return uniqueStrings(summaries)
}

func (s server) buildRecommendationAnomalyContext(ctx context.Context, filter recommendationFilter) (audit.AnalyticsAnomaliesResponse, error) {
	analyticsFilter := audit.AnalyticsFilter{
		Window:      "28d",
		CompareTo:   "previous_window",
		GroupBy:     "service",
		ClusterID:   filter.event.ClusterID,
		TenantID:    filter.event.TenantID,
		Environment: filter.event.Environment,
		Repo:        filter.event.Repo,
		Service:     filter.Service,
		Team:        filter.Team,
	}
	analyticsFilter, err := audit.NormalizeAnalyticsFilter(analyticsFilter)
	if err != nil {
		return audit.AnalyticsAnomaliesResponse{}, err
	}
	return s.buildAnalyticsAnomaliesResponse(ctx, analyticsFilter)
}

func applyAnomalyContextToRecommendations(recommendations []recommendation, anomalies []audit.AnalyticsAnomaly) []recommendation {
	if len(recommendations) == 0 || len(anomalies) == 0 {
		return recommendations
	}
	for i := range recommendations {
		for _, anomaly := range anomalies {
			if recommendationMatchesAnomaly(recommendations[i], anomaly) {
				recommendations[i].PriorityBand = bumpPriorityBand(recommendations[i].PriorityBand)
				recommendations[i].ImpactScore = minInt(100, recommendations[i].ImpactScore+8)
				recommendations[i].Rationale = strings.TrimSpace(fmt.Sprintf("%s Recent anomaly context: %s — %s.", recommendations[i].Rationale, anomaly.Title, anomaly.Deviation))
				recommendations[i].Limitations = append(recommendations[i].Limitations, "Priority was elevated by an explainable anomaly signal from the same canonical scope.")
				break
			}
		}
	}
	return recommendations
}

func recommendationMatchesAnomaly(item recommendation, anomaly audit.AnalyticsAnomaly) bool {
	segment := strings.ToLower(strings.TrimSpace(anomaly.Segment))
	for _, candidate := range []string{item.Service, item.Environment, item.Team, item.Repo} {
		if strings.ToLower(strings.TrimSpace(candidate)) == segment && segment != "" {
			return true
		}
	}
	return false
}

func (s server) listRecommendationMutationEvents(ctx context.Context, filter recommendationFilter) ([]audit.StoredEvent, error) {
	eventFilter := filter.event
	eventFilter.Component = recommendationComponent
	eventFilter.Limit = 1000
	events, err := s.store.ListEvents(ctx, eventFilter)
	if err != nil {
		return nil, err
	}
	mutations := make([]audit.StoredEvent, 0, len(events))
	for _, event := range events {
		if isRecommendationMutationEvent(event) {
			mutations = append(mutations, event)
		}
	}
	sort.Slice(mutations, func(i, j int) bool {
		return eventTimestamp(mutations[i]).Before(eventTimestamp(mutations[j]))
	})
	return mutations, nil
}

func applyRecommendationMutations(candidates []recommendation, events []audit.StoredEvent) []recommendation {
	index := map[string]*recommendation{}
	for i := range candidates {
		index[candidates[i].RecommendationID] = &candidates[i]
	}
	for _, event := range events {
		recommendationID := strings.TrimSpace(event.RecommendationID)
		if recommendationID == "" {
			continue
		}
		target, ok := index[recommendationID]
		if !ok {
			synthetic := recommendationFromMutationEvent(event)
			index[recommendationID] = &synthetic
			candidates = append(candidates, synthetic)
			target = &candidates[len(candidates)-1]
		}
		applyRecommendationMutation(target, event)
	}
	return candidates
}

func recommendationFromMutationEvent(event audit.StoredEvent) recommendation {
	timestamp := eventTimestamp(event)
	template, ok := recommendationTemplateByID(event.RecommendationTemplateID)
	if !ok {
		template = recommendationTemplateCatalog[0]
	}
	readbackRefs := make([]advisoryReadbackRef, 0, len(event.RecommendationReadbackRefs))
	for _, value := range event.RecommendationReadbackRefs {
		parts := strings.Split(strings.TrimSpace(value), "|")
		if len(parts) != 4 {
			continue
		}
		readbackRefs = append(readbackRefs, advisoryReadbackRef{
			ResourceType: parts[0],
			ResourceID:   parts[1],
			ResourceURI:  parts[2],
			EvidenceHash: parts[3],
		})
	}
	return recommendation{
		RecommendationID:    strings.TrimSpace(event.RecommendationID),
		SourceType:          strings.TrimSpace(event.RecommendationSourceType),
		SourceRef:           strings.TrimSpace(event.RecommendationSourceRef),
		SubjectType:         strings.TrimSpace(event.RecommendationSubjectType),
		SubjectRef:          strings.TrimSpace(event.RecommendationSubjectRef),
		Team:                strings.TrimSpace(event.TenantID),
		Service:             strings.TrimSpace(event.Workload),
		Repo:                strings.TrimSpace(event.Repo),
		Environment:         strings.TrimSpace(event.Environment),
		RecommendationType:  strings.TrimSpace(event.RecommendationType),
		Title:               strings.TrimSpace(event.RecommendationTitle),
		Description:         strings.TrimSpace(event.RecommendationDescription),
		RecommendedAction:   strings.TrimSpace(event.RecommendationAction),
		Rationale:           strings.TrimSpace(event.RecommendationRationale),
		EvidenceRefs:        cloneStrings(event.RecommendationEvidenceRefs),
		ReadbackRefs:        readbackRefs,
		RelatedIncidentRefs: cloneStrings(event.RecommendationRelatedIncidentRefs),
		PriorityBand:        strings.TrimSpace(event.RecommendationPriorityBand),
		ImpactScore:         event.RecommendationImpactScore,
		EffortScore:         event.RecommendationEffortScore,
		ConfidenceScore:     event.RecommendationConfidenceScore,
		ApprovalMode:        firstNonEmpty(strings.TrimSpace(event.RecommendationApprovalMode), template.ApprovalMode),
		Status:              firstNonEmpty(strings.TrimSpace(event.RecommendationStatus), recommendationStatusShown),
		CreatedAt:           timestamp,
		ExpiresAt:           event.RecommendationExpiresAt,
		SupersededBy:        strings.TrimSpace(event.RecommendationSupersededBy),
		VerificationPlan:    cloneStrings(event.RecommendationVerificationPlan),
		FeedbackSummary:     strings.TrimSpace(event.RecommendationFeedbackSummary),
		ActionTemplate:      template,
		Owner:               strings.TrimSpace(event.RecommendationOwner),
		Outcome: recommendationOutcome{
			Status: strings.TrimSpace(firstNonEmpty(event.RecommendationVerificationResult, event.RecommendationStatus, recommendationStatusShown)),
		},
		AdvisoryOnly: true,
		Limitations: []string{
			"The original recommendation context is no longer active in the current scope, so this overlay entry is replayed from stored workflow mutations only.",
		},
	}
}

func applyRecommendationMutation(item *recommendation, event audit.StoredEvent) {
	if item == nil {
		return
	}
	timestamp := eventTimestamp(event)
	if item.CreatedAt.IsZero() || timestamp.Before(item.CreatedAt) {
		item.CreatedAt = timestamp
	}
	if strings.TrimSpace(event.RecommendationOwner) != "" {
		item.Owner = strings.TrimSpace(event.RecommendationOwner)
	}
	if event.RecommendationExpiresAt != nil {
		item.ExpiresAt = event.RecommendationExpiresAt
	}
	if strings.TrimSpace(event.RecommendationSupersededBy) != "" {
		item.SupersededBy = strings.TrimSpace(event.RecommendationSupersededBy)
	}
	if strings.TrimSpace(event.RecommendationFeedbackSummary) != "" {
		item.FeedbackSummary = strings.TrimSpace(event.RecommendationFeedbackSummary)
		item.Outcome.Summary = item.FeedbackSummary
	}
	if strings.TrimSpace(event.RecommendationStatus) != "" {
		item.Status = strings.TrimSpace(event.RecommendationStatus)
	}
	if strings.TrimSpace(event.RecommendationVerificationResult) != "" {
		item.Outcome.Status = strings.TrimSpace(event.RecommendationVerificationResult)
		item.Outcome.VerifiedAt = timePointer(timestamp)
	}
	if strings.TrimSpace(event.RecommendationTemplateID) != "" {
		if template, ok := recommendationTemplateByID(event.RecommendationTemplateID); ok {
			item.ActionTemplate = template
			item.ApprovalMode = template.ApprovalMode
		}
	}
	if strings.TrimSpace(event.RecommendationComment) != "" {
		item.Comments = append(item.Comments, recommendationComment{
			ID:        fmt.Sprintf("comment-%d", timestamp.UnixNano()),
			Comment:   strings.TrimSpace(event.RecommendationComment),
			Actor:     strings.TrimSpace(event.Actor),
			Timestamp: timePointer(timestamp),
		})
	}
	item.History = append(item.History, recommendationHistoryEntry{
		ID:        fmt.Sprintf("%s-%d", event.EventType, timestamp.UnixNano()),
		EventType: event.EventType,
		Title:     recommendationMutationTitle(event),
		Summary:   recommendationMutationSummary(event),
		Actor:     strings.TrimSpace(event.Actor),
		Timestamp: timePointer(timestamp),
	})
}

func (s server) recordRecommendationMutation(ctx context.Context, principal auth.Principal, item recommendation, eventType string, mutate func(*audit.Event)) (recommendation, error) {
	event := audit.Event{
		Component:                         recommendationComponent,
		EventType:                         eventType,
		Decision:                          audit.DecisionAllow,
		Actor:                             incidentActor(principal),
		ClusterID:                         "",
		TenantID:                          item.Team,
		Repo:                              item.Repo,
		Environment:                       item.Environment,
		Workload:                          item.Service,
		RecommendationID:                  item.RecommendationID,
		RecommendationSourceType:          item.SourceType,
		RecommendationSourceRef:           item.SourceRef,
		RecommendationSubjectType:         item.SubjectType,
		RecommendationSubjectRef:          item.SubjectRef,
		RecommendationType:                item.RecommendationType,
		RecommendationTitle:               item.Title,
		RecommendationDescription:         item.Description,
		RecommendationAction:              item.RecommendedAction,
		RecommendationRationale:           item.Rationale,
		RecommendationEvidenceRefs:        item.EvidenceRefs,
		RecommendationReadbackRefs:        serializeReadbackRefs(item.ReadbackRefs),
		RecommendationRelatedIncidentRefs: item.RelatedIncidentRefs,
		RecommendationPriorityBand:        item.PriorityBand,
		RecommendationImpactScore:         item.ImpactScore,
		RecommendationEffortScore:         item.EffortScore,
		RecommendationConfidenceScore:     item.ConfidenceScore,
		RecommendationApprovalMode:        item.ApprovalMode,
		RecommendationStatus:              item.Status,
		RecommendationTemplateID:          item.ActionTemplate.TemplateID,
		RecommendationVerificationPlan:    item.VerificationPlan,
		RecommendationFeedbackSummary:     item.FeedbackSummary,
		RecommendationOwner:               item.Owner,
		RecommendationSupersededBy:        item.SupersededBy,
		RecommendationExpiresAt:           item.ExpiresAt,
	}
	if mutate != nil {
		mutate(&event)
	}
	if _, err := s.store.Ingest(ctx, event); err != nil {
		return recommendation{}, err
	}
	return s.getRecommendationByID(ctx, item.RecommendationID, recommendationFilter{
		event: audit.EventFilter{
			TenantID:    item.Team,
			Environment: item.Environment,
			Repo:        item.Repo,
			Limit:       500,
		},
		Limit: 200,
	})
}

func (s server) executeRecommendation(ctx context.Context, principal auth.Principal, item recommendation, template recommendationActionTemplate, summary string) (recommendation, error) {
	if template.ApprovalMode == recommendationApprovalHumanReview && item.Status != recommendationStatusAccepted {
		return recommendation{}, errRecommendationApprovalRequired
	}
	feedback := recommendationExecutionSummary(item, template, summary)
	return s.recordRecommendationMutation(ctx, principal, item, recommendationEventExecuted, func(event *audit.Event) {
		event.RecommendationStatus = recommendationStatusExecuted
		event.RecommendationTemplateID = template.TemplateID
		event.RecommendationFeedbackSummary = feedback
	})
}

func (s server) verifyRecommendation(ctx context.Context, item recommendation, filter recommendationFilter) (string, string, error) {
	incidents, err := s.listIncidents(ctx, incidentFilter{event: filter.event})
	if err != nil {
		return "", "", err
	}
	switch item.SourceType {
	case "incident":
		for _, incident := range incidents {
			if incident.ID != item.SubjectRef {
				continue
			}
			switch incident.State {
			case incidentStateResolved:
				return recommendationStatusVerifiedSuccessful, "The linked incident is resolved and no longer carries active pressure in the current scope.", nil
			case incidentStateWatching, incidentStateAcknowledged:
				return recommendationStatusPartiallyEffective, "The linked incident is still present but pressure has narrowed into a watched or acknowledged state.", nil
			case incidentStateReopened:
				return recommendationStatusRegressed, "The linked incident reopened, so the executed recommendation did not hold the line.", nil
			default:
				return recommendationStatusExecutedNoEffect, "The linked incident remains active in the current scope, so the executed step has not reduced the main signal yet.", nil
			}
		}
		return recommendationStatusVerifiedSuccessful, "The linked incident is no longer present in the current scope.", nil
	case "package":
		related := selectIncidentsByID(incidents, item.RelatedIncidentRefs)
		if len(related) == 0 {
			return recommendationStatusVerifiedSuccessful, "All package-linked incidents dropped out of the current scope after the workflow action.", nil
		}
		resolved := 0
		for _, incident := range related {
			if incident.State == incidentStateResolved {
				resolved++
			}
		}
		if resolved == len(related) {
			return recommendationStatusVerifiedSuccessful, "Every incident linked to the package recommendation is now resolved.", nil
		}
		if resolved > 0 {
			return recommendationStatusPartiallyEffective, fmt.Sprintf("%d of %d package-linked incidents are resolved, but the package still carries residual pressure.", resolved, len(related)), nil
		}
		return recommendationStatusExecutedNoEffect, "The package still carries the same linked incident set in the current scope.", nil
	case "systemic_weakness":
		response := buildSystemicWeaknessResponse(incidents, "current filtered scope")
		for _, weakness := range response.Weaknesses {
			if weakness.PatternKey != item.SourceRef {
				continue
			}
			if len(weakness.RelatedIncidentRefs) < len(item.RelatedIncidentRefs) {
				return recommendationStatusPartiallyEffective, "The same weakness pattern is still present, but it now covers fewer incidents than before.", nil
			}
			return recommendationStatusExecutedNoEffect, "The same systemic weakness pattern is still present at the same scale.", nil
		}
		return recommendationStatusVerifiedSuccessful, "The previously linked systemic weakness pattern is no longer present in the current scope.", nil
	case "anomaly":
		response, err := s.buildRecommendationAnomalyContext(ctx, filter)
		if err != nil {
			return "", "", err
		}
		for _, anomaly := range response.Items {
			if item.SourceRef == anomaly.Type+":"+anomaly.Segment {
				if anomaly.Severity == "high" {
					return recommendationStatusRegressed, "The same anomaly is still firing at high severity in the current comparison window.", nil
				}
				return recommendationStatusExecutedNoEffect, "The same anomaly is still present in the current comparison window.", nil
			}
		}
		return recommendationStatusVerifiedSuccessful, "The anomaly signal is no longer present in the current comparison window.", nil
	case "topology_signal":
		topologyFilter, err := recommendationTopologyFilter(filter)
		if err != nil {
			return "", "", err
		}
		topologyFilter.Service = firstNonEmpty(item.SubjectRef, item.Service)
		blast, err := s.buildTopologyBlastRadiusForService(ctx, topologyFilter)
		if err != nil {
			return "", "", err
		}
		delta, err := s.buildTopologyDeltaResponse(ctx, topologyFilter)
		if err != nil {
			return "", "", err
		}
		for _, deltaItem := range delta.Items {
			if deltaItem.Service != topologyFilter.Service {
				continue
			}
			switch {
			case containsString(deltaItem.DriftSignals, "public exposure widened") || deltaItem.CriticalReachDelta > 0:
				return recommendationStatusRegressed, "The same service still shows widening public or critical downstream reach in topology delta.", nil
			case deltaItem.Delta > 0 || blast.BlastRadiusScore >= 60:
				return recommendationStatusExecutedNoEffect, "The same service still carries elevated blast-radius pressure in the current topology window.", nil
			case blast.BlastRadiusScore > 0:
				return recommendationStatusPartiallyEffective, "Blast-radius pressure narrowed, but the service still carries residual topological reach that should be tracked.", nil
			}
		}
		if blast.BlastRadiusScore == 0 && len(blast.ReachableNodes) == 0 {
			return recommendationStatusVerifiedSuccessful, "The topology signal no longer maps to an active blast-radius path in the current scope.", nil
		}
		return recommendationStatusPartiallyEffective, "The strongest drift signal cleared, but the service still retains some topology-mapped reach.", nil
	case "forensic_signal":
		forensicFilter, err := recommendationForensicsFilter(filter, incidents)
		if err != nil {
			return "", "", err
		}
		forensicFilter.Timestamp = parseForensicRecommendationSourceRef(item.SourceRef, item.CreatedAt)
		forensicFilter.Service = firstNonEmpty(item.Service, forensicFilter.Service)
		if item.SubjectType == "incident" {
			forensicFilter.IncidentID = item.SubjectRef
		}
		delta, err := s.buildForensicsDeltaResponse(ctx, forensicFilter)
		if err != nil {
			return "", "", err
		}
		replay, err := s.buildForensicsReplay(ctx, withForensicsTimestamp(forensicFilter, delta.Comparison.T1), forensicsReplayModernFullStack)
		if err != nil {
			return "", "", err
		}
		switch {
		case replay.VerdictDelta == "no_change" && len(delta.VulnerabilityDelta.Added) == 0 && len(delta.IdentityDelta.Modified) == 0 && len(delta.TopologyDelta) == 0:
			return recommendationStatusVerifiedSuccessful, "Historical replay no longer diverges from the current control stack and the compared forensic windows no longer show residual drift.", nil
		case replay.VerdictDelta == "no_change":
			return recommendationStatusPartiallyEffective, "The strict replay delta cleared, but residual forensic drift still exists in the compared windows.", nil
		default:
			return recommendationStatusExecutedNoEffect, "Historical replay still diverges under the current control stack, so the forensic review remains active.", nil
		}
	case "runtime_signal":
		runtimeFilter := recommendationRuntimeFilter(filter)
		runtimeFilter.SubjectRef = item.SubjectRef
		findings, _, err := s.buildRuntimeFindings(ctx, runtimeFilter)
		if err != nil {
			return "", "", err
		}
		states, _, err := s.buildRuntimeIntegrityStates(ctx, runtimeFilter)
		if err != nil {
			return "", "", err
		}
		for _, finding := range findings {
			if finding.FindingID != item.SourceRef {
				continue
			}
			switch finding.Severity {
			case "critical":
				return recommendationStatusRegressed, "The same critical runtime finding is still active for the workload, so the executed workflow did not hold containment.", nil
			case "high":
				return recommendationStatusExecutedNoEffect, "The same high-severity runtime finding is still active for the workload.", nil
			default:
				return recommendationStatusPartiallyEffective, "The same runtime finding is still present, but it has narrowed below the highest severity band.", nil
			}
		}
		for _, state := range states {
			if state.SubjectRef == item.SubjectRef && len(state.ActiveFindings) > 0 {
				return recommendationStatusPartiallyEffective, "The original runtime signal cleared, but the workload still carries residual runtime findings in the current integrity state.", nil
			}
		}
		return recommendationStatusVerifiedSuccessful, "The runtime drift signal is no longer active for the workload and the current integrity state no longer carries the same finding.", nil
	case "hardening_signal":
		runtimeFilter := recommendationRuntimeFilter(filter)
		runtimeFilter.SubjectRef = item.SubjectRef
		executions, _, err := s.listHardeningExecutions(ctx, runtimeFilter)
		if err != nil {
			return "", "", err
		}
		posture, _, err := s.buildDefensePostureStates(ctx, runtimeFilter)
		if err != nil {
			return "", "", err
		}
		findings, _, err := s.buildRuntimeFindings(ctx, runtimeFilter)
		if err != nil {
			return "", "", err
		}
		for _, finding := range findings {
			if finding.SubjectRef == item.SubjectRef && runtimeSeverityRank(finding.Severity) >= runtimeSeverityRank("high") {
				return recommendationStatusRegressed, "The workload still carries a high-severity runtime signal, so temporary hardening has not stabilized the subject yet.", nil
			}
		}
		for _, execution := range executions {
			if execution.ExecutionID != item.SourceRef {
				continue
			}
			if execution.ExecutionResult == "rollback_applied" || execution.ExecutionResult == "trusted_recovery_completed" {
				return recommendationStatusVerifiedSuccessful, "Temporary hardening has already been cleared through rollback or trusted recovery in the current scope.", nil
			}
			for _, state := range posture {
				if state.SubjectRef != item.SubjectRef {
					continue
				}
				if len(state.ActiveRestrictions) > 0 || state.CurrentMode == hardeningModePendingApproval {
					return recommendationStatusPartiallyEffective, "Temporary hardening still remains active or pending approval, so the workflow should stay open until the durable remediation closes the loop.", nil
				}
			}
			return recommendationStatusPartiallyEffective, "The original hardening record remains visible, but active restrictions are no longer the dominant posture in the current scope.", nil
		}
		return recommendationStatusVerifiedSuccessful, "The referenced hardening action is no longer active in the current scope.", nil
	case "federation_signal":
		view, err := s.buildFederationGlobalView(ctx)
		if err != nil {
			return "", "", err
		}
		signalType, signalRef := parseFederationRecommendationSourceRef(item.SourceRef)
		switch signalType {
		case "peer":
			for _, peerID := range view.StalePeers {
				if peerID == signalRef {
					return recommendationStatusExecutedNoEffect, "The federation peer is still stale and has not re-established a fresh trusted state.", nil
				}
			}
			return recommendationStatusVerifiedSuccessful, "The federation peer is no longer stale and currently passes the local trust health view.", nil
		case "policy":
			if len(view.PolicyDivergence) == 0 {
				return recommendationStatusVerifiedSuccessful, "Federation policy divergence no longer blocks the local effective policy state.", nil
			}
			return recommendationStatusExecutedNoEffect, "Federation policy divergence remains unresolved in the current local-overrides view.", nil
		default:
			for _, historyItem := range view.ProofHistory {
				if historyItem.PeerID == signalRef && strings.HasPrefix(historyItem.Decision, "rejected") {
					return recommendationStatusExecutedNoEffect, "The same federated proof path is still being rejected by the local trust decision engine.", nil
				}
			}
			return recommendationStatusVerifiedSuccessful, "The previously rejected or stale federated proof path is no longer present in the current trust history.", nil
		}
	case "validation_signal":
		score, err := s.buildValidationHarnessScore(ctx, validationHarnessFilter{
			ClusterID:   filter.event.ClusterID,
			TenantID:    filter.event.TenantID,
			Environment: filter.event.Environment,
			Repo:        filter.event.Repo,
			Service:     item.Service,
			Limit:       maxInt(filter.Limit, 6),
			event: audit.EventFilter{
				ClusterID:   filter.event.ClusterID,
				TenantID:    filter.event.TenantID,
				Environment: filter.event.Environment,
				Repo:        filter.event.Repo,
				Limit:       maxInt(filter.Limit*20, 500),
			},
		})
		if err != nil {
			return "", "", err
		}
		for _, result := range score.Results {
			if result.ScenarioID != item.SourceRef {
				continue
			}
			switch result.Status {
			case validationStatusPass:
				return recommendationStatusVerifiedSuccessful, "The same validation harness scenario now passes in the current scope.", nil
			case validationStatusPartial:
				return recommendationStatusPartiallyEffective, "The validation harness scenario improved but still remains only partially verified.", nil
			default:
				return recommendationStatusExecutedNoEffect, "The validation harness still reports the same scenario as a gap in the current scope.", nil
			}
		}
		return recommendationStatusVerifiedSuccessful, "The validation harness no longer reports the same scenario as a current gap.", nil
	default:
		return recommendationStatusExecutedNoEffect, "The recommendation source does not yet expose a richer verification rule, so the workflow remains advisory-only.", nil
	}
}

func filterRecommendations(values []recommendation, filter recommendationFilter) []recommendation {
	if len(values) == 0 {
		return values
	}
	filtered := make([]recommendation, 0, len(values))
	for _, item := range values {
		if filter.SourceType != "" && item.SourceType != filter.SourceType {
			continue
		}
		if filter.SubjectType != "" && item.SubjectType != filter.SubjectType {
			continue
		}
		if filter.RecommendationType != "" && item.RecommendationType != filter.RecommendationType {
			continue
		}
		if filter.Team != "" && item.Team != filter.Team {
			continue
		}
		if filter.Service != "" && item.Service != filter.Service && item.Repo != filter.Service {
			continue
		}
		if filter.Status != "" && item.Status != filter.Status && item.Outcome.Status != filter.Status {
			continue
		}
		filtered = append(filtered, item)
	}
	return filtered
}

func sortRecommendations(values []recommendation) {
	sort.Slice(values, func(i, j int) bool {
		if priorityRank(values[i].PriorityBand) != priorityRank(values[j].PriorityBand) {
			return priorityRank(values[i].PriorityBand) > priorityRank(values[j].PriorityBand)
		}
		if values[i].ImpactScore != values[j].ImpactScore {
			return values[i].ImpactScore > values[j].ImpactScore
		}
		if values[i].ConfidenceScore != values[j].ConfidenceScore {
			return values[i].ConfidenceScore > values[j].ConfidenceScore
		}
		return values[i].RecommendationID < values[j].RecommendationID
	})
}

func recommendationTemplateByID(templateID string) (recommendationActionTemplate, bool) {
	for _, template := range recommendationTemplateCatalog {
		if template.TemplateID == strings.TrimSpace(templateID) {
			return template, true
		}
	}
	return recommendationActionTemplate{}, false
}

func firstDefenseGap(values []defenseGapFinding) defenseGapFinding {
	if len(values) == 0 {
		return buildDefenseGapFinding("policy_coverage", "limited", nil, nil, "Current scoped incident")
	}
	return values[0]
}

func selectIncidentsByID(values []investigationIncident, ids []string) []investigationIncident {
	if len(ids) == 0 {
		return values
	}
	selected := make([]investigationIncident, 0, len(ids))
	index := map[string]investigationIncident{}
	for _, incident := range values {
		index[incident.ID] = incident
	}
	for _, id := range ids {
		if incident, ok := index[id]; ok {
			selected = append(selected, incident)
		}
	}
	return selected
}

func firstIncidentTenant(values []investigationIncident) string {
	for _, item := range values {
		if strings.TrimSpace(item.TenantID) != "" {
			return strings.TrimSpace(item.TenantID)
		}
	}
	return ""
}

func firstIncidentRepo(values []investigationIncident) string {
	for _, item := range values {
		if strings.TrimSpace(item.Repository) != "" {
			return strings.TrimSpace(item.Repository)
		}
	}
	return ""
}

func firstIncidentEnvironment(values []investigationIncident) string {
	for _, item := range values {
		if strings.TrimSpace(item.Environment) != "" {
			return strings.TrimSpace(item.Environment)
		}
	}
	return ""
}

func firstPackageTenant(values []investigationIncident) string {
	return firstIncidentTenant(values)
}

func firstPackageRepo(values []investigationIncident) string {
	return firstIncidentRepo(values)
}

func firstPackageEnvironment(values []investigationIncident) string {
	return firstIncidentEnvironment(values)
}

func packageEvidenceRefs(intelligence packageIntelligence) []string {
	values := make([]string, 0, 16)
	for _, finding := range intelligence.DefenseGapSummary.TopFindings {
		values = append(values, finding.EvidenceRefs...)
	}
	for _, gap := range intelligence.PolicyReplaySummary.TopCoverageGaps {
		values = append(values, gap.EvidenceRefs...)
	}
	for _, pattern := range intelligence.SystemicWeakness.TopPatterns {
		values = append(values, pattern.EvidenceRefs...)
	}
	return limitStrings(uniqueStrings(values), 12)
}

func packageRecommendationSourceRef(incidentRefs []string) string {
	sum := sha1.Sum([]byte(strings.Join(uniqueStrings(incidentRefs), "\x1f")))
	return "PKG-" + strings.ToUpper(fmt.Sprintf("%x", sum[:]))[:10]
}

func recommendationID(sourceType, sourceRef, templateID string) string {
	sum := sha1.Sum([]byte(strings.Join([]string{sourceType, sourceRef, templateID}, "\x1f")))
	return "REC-" + strings.ToUpper(fmt.Sprintf("%x", sum[:]))[:12]
}

func recommendationPriorityBand(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "critical", "now":
		return "NOW"
	case "high", "today":
		return "TODAY"
	case "medium", "this_week", "watch":
		return "THIS_WEEK"
	default:
		return "BACKLOG"
	}
}

func bumpPriorityBand(value string) string {
	switch recommendationPriorityBand(value) {
	case "BACKLOG":
		return "THIS_WEEK"
	case "THIS_WEEK":
		return "TODAY"
	default:
		return "NOW"
	}
}

func priorityRank(value string) int {
	switch recommendationPriorityBand(value) {
	case "NOW":
		return 4
	case "TODAY":
		return 3
	case "THIS_WEEK":
		return 2
	default:
		return 1
	}
}

func recommendationImpactScore(severity string, workloads int, environments int) int {
	score := 40
	switch strings.ToLower(strings.TrimSpace(severity)) {
	case "critical":
		score = 92
	case "high":
		score = 82
	case "medium":
		score = 64
	default:
		score = 45
	}
	score += minInt(10, workloads*3)
	score += minInt(8, environments*2)
	return minInt(score, 100)
}

func systemicImpactScore(value systemicWeakness) int {
	score := 60 + minInt(20, len(value.RelatedIncidentRefs)*4)
	if strings.ToLower(strings.TrimSpace(value.Priority)) == "critical" {
		score += 10
	}
	return minInt(score, 100)
}

func anomalyImpactScore(severity string) int {
	switch strings.ToLower(strings.TrimSpace(severity)) {
	case "high":
		return 85
	case "medium":
		return 68
	default:
		return 55
	}
}

func recommendationEffortScore(templateID string) int {
	switch templateID {
	case "create_ticket", "notify_owner":
		return 20
	case "open_sandbox", "compare_artifact_versions":
		return 35
	case "generate_remediation_draft", "draft_vex":
		return 45
	case "request_exception", "create_security_review", "archive_stale_exception":
		return 60
	default:
		return 40
	}
}

func recommendationConfidenceScore(value string) int {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "high":
		return 90
	case "medium":
		return 72
	default:
		return 55
	}
}

func packageRecommendationPriority(intelligence packageIntelligence) string {
	if len(intelligence.DefenseGapSummary.TopGapTypes) == 0 {
		return "THIS_WEEK"
	}
	if intelligence.DefenseGapSummary.ConfidenceMix["high"] > 0 || intelligence.PolicyReplaySummary.Delta.ImpactedCases >= 3 {
		return "NOW"
	}
	return "TODAY"
}

func packageRecommendationConfidence(intelligence packageIntelligence) int {
	if intelligence.DefenseGapSummary.ConfidenceMix["high"] > 0 {
		return 86
	}
	if intelligence.DefenseGapSummary.ConfidenceMix["medium"] > 0 {
		return 72
	}
	return 58
}

func recommendationCreatedAt(candidates ...*time.Time) time.Time {
	for _, candidate := range candidates {
		if candidate != nil && !candidate.IsZero() {
			return candidate.UTC()
		}
	}
	return time.Unix(0, 0).UTC()
}

func recommendationExpiry(candidate *time.Time) *time.Time {
	if candidate == nil || candidate.IsZero() {
		return nil
	}
	value := candidate.UTC().Add(7 * 24 * time.Hour)
	return &value
}

func serializeReadbackRefs(values []advisoryReadbackRef) []string {
	serialized := make([]string, 0, len(values))
	for _, value := range values {
		if strings.TrimSpace(value.ResourceID) == "" {
			continue
		}
		serialized = append(serialized, strings.Join([]string{
			value.ResourceType,
			value.ResourceID,
			value.ResourceURI,
			value.EvidenceHash,
		}, "|"))
	}
	return serialized
}

func recommendationMutationTitle(event audit.StoredEvent) string {
	switch event.EventType {
	case recommendationEventAcknowledged:
		return "Recommendation acknowledged"
	case recommendationEventAccepted:
		return "Recommendation accepted"
	case recommendationEventRejected:
		return "Recommendation rejected"
	case recommendationEventExecuted:
		return "Action executed"
	case recommendationEventVerified:
		return "Outcome verified"
	case recommendationEventAssigned:
		return "Owner assigned"
	case recommendationEventCommented:
		return "Comment added"
	case recommendationEventApprovalRequest:
		return "Approval requested"
	default:
		return "Recommendation updated"
	}
}

func recommendationMutationSummary(event audit.StoredEvent) string {
	switch event.EventType {
	case recommendationEventRejected:
		return firstNonEmpty(strings.TrimSpace(event.RecommendationFeedbackSummary), "The recommendation was rejected.")
	case recommendationEventExecuted:
		return firstNonEmpty(strings.TrimSpace(event.RecommendationFeedbackSummary), "A workflow action was executed for this recommendation.")
	case recommendationEventVerified:
		return firstNonEmpty(strings.TrimSpace(event.RecommendationFeedbackSummary), "Recommendation impact was verified against current canonical signals.")
	case recommendationEventAssigned:
		if strings.TrimSpace(event.RecommendationOwner) != "" {
			return fmt.Sprintf("Assigned to %s.", strings.TrimSpace(event.RecommendationOwner))
		}
	case recommendationEventCommented:
		return firstNonEmpty(strings.TrimSpace(event.RecommendationComment), "A workflow comment was recorded.")
	case recommendationEventApprovalRequest:
		return firstNonEmpty(strings.TrimSpace(event.RecommendationFeedbackSummary), "Approval was requested for a sensitive workflow action.")
	}
	if strings.TrimSpace(event.RecommendationFeedbackSummary) != "" {
		return strings.TrimSpace(event.RecommendationFeedbackSummary)
	}
	return recommendationMutationTitle(event)
}

func recommendationExecutionSummary(item recommendation, template recommendationActionTemplate, summary string) string {
	if strings.TrimSpace(summary) != "" {
		return strings.TrimSpace(summary)
	}
	switch template.TemplateID {
	case "create_ticket":
		return fmt.Sprintf("Prepared a ticket-ready remediation brief for %s with %d evidence refs and %d linked incidents.", item.SubjectRef, len(item.EvidenceRefs), len(item.RelatedIncidentRefs))
	case "open_sandbox":
		return fmt.Sprintf("Prepared a bounded sandbox workflow for %s so the weakest control path can be reproduced safely before production changes.", item.SubjectRef)
	case "generate_remediation_draft":
		return fmt.Sprintf("Generated a remediation draft for %s using the linked evidence, replay, and recommendation rationale.", item.SubjectRef)
	case "draft_vex":
		return fmt.Sprintf("Prepared a VEX-style triage draft for %s from the current evidence bundle.", item.SubjectRef)
	case "notify_owner":
		return fmt.Sprintf("Prepared an owner notification for %s so the recommendation can be routed without mutating canonical truth.", item.SubjectRef)
	default:
		return fmt.Sprintf("Executed %s for %s.", template.Title, item.SubjectRef)
	}
}

func isRecommendationMutationEvent(event audit.StoredEvent) bool {
	if strings.TrimSpace(event.Component) != recommendationComponent {
		return false
	}
	switch event.EventType {
	case recommendationEventAcknowledged, recommendationEventAccepted, recommendationEventRejected, recommendationEventExecuted, recommendationEventVerified, recommendationEventAssigned, recommendationEventCommented, recommendationEventApprovalRequest:
		return true
	default:
		return false
	}
}

func writeRecommendationError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if errors.Is(err, errIncidentNotFound) {
		status = http.StatusNotFound
	} else if errors.Is(err, errRecommendationApprovalRequired) {
		status = http.StatusConflict
	} else if errors.Is(err, audit.ErrInvalidFilter) {
		status = http.StatusBadRequest
	} else if errors.Is(err, context.DeadlineExceeded) {
		status = http.StatusGatewayTimeout
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}

func (f recommendationFilter) toIncidentFilter() incidentFilter {
	return incidentFilter{event: f.event}
}
