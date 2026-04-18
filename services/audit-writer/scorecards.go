package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

type trustAuditConfig struct {
	PublicationMode        string
	HardeningReviewEnabled bool
	StaleExceptionDays     int
	EventLimit             int
	SeverityThreshold      string
}

type trustScopeRequest struct {
	TenantID          string `json:"tenant_id,omitempty"`
	ClusterID         string `json:"cluster_id,omitempty"`
	Environment       string `json:"environment,omitempty"`
	Repo              string `json:"repo,omitempty"`
	Format            string `json:"format,omitempty"`
	IncludePublicView bool   `json:"include_public_view,omitempty"`
}

type trustBadgesResponse struct {
	Items           []audit.TrustBadge `json:"items"`
	PublicationMode string             `json:"publication_mode"`
}

type scorecardFindingsResponse struct {
	Items []audit.AuditFinding `json:"items"`
}

func loadTrustAuditConfigFromEnv() (trustAuditConfig, error) {
	publicationMode := strings.ToLower(strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_TRUST_PUBLICATION_MODE"), audit.TrustPublicationDisabled)))
	switch publicationMode {
	case audit.TrustPublicationDisabled, audit.TrustPublicationPreview, audit.TrustPublicationExport:
	default:
		return trustAuditConfig{}, errors.New("invalid CHANGELOCK_TRUST_PUBLICATION_MODE")
	}

	hardeningReviewEnabled, err := strconv.ParseBool(firstNonEmpty(strings.TrimSpace(os.Getenv("CHANGELOCK_HARDENING_REVIEW_ENABLED")), "true"))
	if err != nil {
		return trustAuditConfig{}, errors.New("invalid CHANGELOCK_HARDENING_REVIEW_ENABLED")
	}

	staleExceptionDays := parseIntOrDefault(strings.TrimSpace(os.Getenv("CHANGELOCK_HARDENING_REVIEW_STALE_EXCEPTION_DAYS")), 14)
	if staleExceptionDays <= 0 {
		return trustAuditConfig{}, errors.New("invalid CHANGELOCK_HARDENING_REVIEW_STALE_EXCEPTION_DAYS")
	}

	eventLimit := parseIntOrDefault(strings.TrimSpace(os.Getenv("CHANGELOCK_SCORECARD_EVENT_LIMIT")), 250)
	if eventLimit <= 0 || eventLimit > 500 {
		return trustAuditConfig{}, errors.New("invalid CHANGELOCK_SCORECARD_EVENT_LIMIT")
	}

	severityThreshold := strings.ToUpper(strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_SCORECARD_SEVERITY_THRESHOLD"), "HIGH")))
	switch severityThreshold {
	case "", "UNKNOWN", "LOW", "MEDIUM", "HIGH", "CRITICAL":
	default:
		return trustAuditConfig{}, errors.New("invalid CHANGELOCK_SCORECARD_SEVERITY_THRESHOLD")
	}

	return trustAuditConfig{
		PublicationMode:        publicationMode,
		HardeningReviewEnabled: hardeningReviewEnabled,
		StaleExceptionDays:     staleExceptionDays,
		EventLimit:             eventLimit,
		SeverityThreshold:      severityThreshold,
	}, nil
}

func (s server) scorecardHandler(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	scope := parseTrustScopeRequestFromQuery(r)
	input, err := s.collectTrustScorecardInput(ctx, scope, cfg)
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, audit.ComputeTrustScorecard(input))
}

func (s server) scorecardFindingsHandler(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	scope := parseTrustScopeRequestFromQuery(r)
	input, err := s.collectTrustScorecardInput(ctx, scope, cfg)
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	card := audit.ComputeTrustScorecard(input)
	findings := []audit.AuditFinding{}
	if cfg.HardeningReviewEnabled {
		findings = audit.BuildHardeningReview(input, card)
	}
	httpjson.Write(w, http.StatusOK, scorecardFindingsResponse{Items: findings})
}

func (s server) trustBadgesHandler(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	scope := parseTrustScopeRequestFromQuery(r)
	input, err := s.collectTrustScorecardInput(ctx, scope, cfg)
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	card := audit.ComputeTrustScorecard(input)
	httpjson.Write(w, http.StatusOK, trustBadgesResponse{
		Items:           audit.BuildTrustBadges(card, input),
		PublicationMode: cfg.PublicationMode,
	})
}

func (s server) publishedTrustViewHandler(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if cfg.PublicationMode == audit.TrustPublicationDisabled {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "published trust view is disabled"})
		return
	}
	scope := parseTrustScopeRequestFromQuery(r)
	input, err := s.collectTrustScorecardInput(ctx, scope, cfg)
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	card := audit.ComputeTrustScorecard(input)
	badges := audit.BuildTrustBadges(card, input)
	mappings := audit.BuildStandardsMapping(card)
	publicView := audit.BuildPublishedTrustView(card, badges, mappings)
	if publicView == nil {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "published trust view is disabled"})
		return
	}
	httpjson.Write(w, http.StatusOK, publicView)
}

func (s server) auditReportsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request trustScopeRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request, err := applyPrincipalTenantToTrustScopeRequest(principal, request)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	input, err := s.collectTrustScorecardInput(ctx, request, cfg)
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	report := s.buildAuditReport(principal, input, cfg, request.Format, request.IncludePublicView)
	if strings.EqualFold(request.Format, "html") {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, audit.RenderAuditReportHTML(report))
		return
	}
	httpjson.Write(w, http.StatusOK, report)
}

func (s server) auditExportsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request trustScopeRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request, err := applyPrincipalTenantToTrustScopeRequest(principal, request)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if request.IncludePublicView && cfg.PublicationMode == audit.TrustPublicationDisabled {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": "sanitized public trust export is disabled"})
		return
	}
	input, err := s.collectTrustScorecardInput(ctx, request, cfg)
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	report := s.buildAuditReport(principal, input, cfg, firstNonEmpty(request.Format, "json"), request.IncludePublicView)
	exportBundle := audit.BuildAuditExportBundle(report, audit.BuildTrustEvidenceExport(input))
	httpjson.Write(w, http.StatusOK, exportBundle)
}

func (s server) buildAuditReport(principal auth.Principal, input audit.TrustScorecardInput, cfg trustAuditConfig, format string, includePublicView bool) audit.AuditReport {
	card := audit.ComputeTrustScorecard(input)
	badges := audit.BuildTrustBadges(card, input)
	mappings := audit.BuildStandardsMapping(card)
	findings := []audit.AuditFinding{}
	if cfg.HardeningReviewEnabled {
		findings = audit.BuildHardeningReview(input, card)
	}
	var publicView *audit.PublishedTrustView
	if includePublicView && cfg.PublicationMode != audit.TrustPublicationDisabled {
		publicView = audit.BuildPublishedTrustView(card, badges, mappings)
	}
	return audit.BuildAuditReport(card, findings, badges, mappings, publicView, format, principal.Subject)
}

func (s server) collectTrustScorecardInput(ctx context.Context, scope trustScopeRequest, cfg trustAuditConfig) (audit.TrustScorecardInput, error) {
	scope = normalizeTrustScopeRequest(scope)
	artifactEvents, err := s.store.ListEvents(ctx, audit.EventFilter{
		EventType:   audit.EventTypeArtifactVerificationResult,
		ClusterID:   scope.ClusterID,
		Repo:        scope.Repo,
		Environment: scope.Environment,
		TenantID:    scope.TenantID,
		Limit:       cfg.EventLimit,
	})
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	policyEvents, err := s.store.ListEvents(ctx, audit.EventFilter{
		EventType:   audit.EventTypePolicyDecision,
		ClusterID:   scope.ClusterID,
		Repo:        scope.Repo,
		Environment: scope.Environment,
		TenantID:    scope.TenantID,
		Limit:       cfg.EventLimit,
	})
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	deployEvents, err := s.store.ListEvents(ctx, audit.EventFilter{
		EventType:   audit.EventTypeDeployGateDecision,
		ClusterID:   scope.ClusterID,
		Repo:        scope.Repo,
		Environment: scope.Environment,
		TenantID:    scope.TenantID,
		Limit:       cfg.EventLimit,
	})
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	policyEvents = append(policyEvents, deployEvents...)

	summary, err := s.store.Summary(ctx, audit.EventFilter{
		ClusterID:   scope.ClusterID,
		Repo:        scope.Repo,
		Environment: scope.Environment,
		TenantID:    scope.TenantID,
		Limit:       cfg.EventLimit,
	})
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	exceptionReport, err := s.store.ExceptionReport(ctx, audit.ExceptionFilter{
		TenantID:    scope.TenantID,
		Environment: scope.Environment,
		Repo:        scope.Repo,
		Limit:       cfg.EventLimit,
	})
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	signingCfg, err := loadSigningIdentityConfig()
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	signingPolicies, err := s.signingIdentityPolicies(ctx)
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	signingFilter := map[string]string{
		"tenant_id":   scope.TenantID,
		"cluster_id":  scope.ClusterID,
		"environment": scope.Environment,
		"repo":        scope.Repo,
		"limit":       strconv.Itoa(cfg.EventLimit),
	}
	observations, err := s.signingIdentityObservations(ctx, signingCfg, signingPolicies, signingFilter)
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	findings, err := s.signingIdentityFindings(ctx, signingCfg, signingPolicies, observations)
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	signingStatus := signingIdentityStatus(signingCfg, signingPolicies, observations, findings)

	runtimeEvents, err := s.store.ListEvents(ctx, audit.EventFilter{
		Component: "runtime-agent",
		ClusterID: scope.ClusterID,
		TenantID:  scope.TenantID,
		Limit:     max(cfg.EventLimit, 250),
	})
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	runtimeEvents = filterEventsForTrustScope(runtimeEvents, scope)
	activeStates := audit.DeriveRuntimeActiveStates(runtimeEvents, audit.RuntimeActiveStateFilter{
		ClusterID: scope.ClusterID,
		TenantID:  scope.TenantID,
		Limit:     max(cfg.EventLimit, 250),
	})
	runtimeStatus := audit.DeriveRuntimeClosedLoopStatus(activeStates)

	vexFilter := internalvex.Filter{
		TenantID:    scope.TenantID,
		ClusterID:   scope.ClusterID,
		Environment: scope.Environment,
		Repo:        scope.Repo,
		Active:      trustBoolPointer(true),
		Limit:       cfg.EventLimit,
	}
	vexStatements, err := s.store.ListVEXStatements(ctx, vexFilter)
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}
	vexStatus := summarizeVEXStatements(vexStatements, vexFilter)
	vulnNet, err := s.computeVulnerabilityNetResponse(ctx, audit.VulnerabilityActiveFilter{
		TenantID:    scope.TenantID,
		Environment: scope.Environment,
		Limit:       cfg.EventLimit,
	}, cfg.SeverityThreshold)
	if err != nil {
		return audit.TrustScorecardInput{}, err
	}

	return audit.TrustScorecardInput{
		ScopeType:                  trustScopeType(scope),
		ScopeRef:                   trustScopeRef(scope),
		TenantID:                   scope.TenantID,
		ClusterID:                  scope.ClusterID,
		Environment:                scope.Environment,
		Repo:                       scope.Repo,
		CalculatedAt:               time.Now().UTC(),
		ArtifactVerificationEvents: filterEventsForTrustScope(artifactEvents, scope),
		PolicyDecisionEvents:       filterEventsForTrustScope(policyEvents, scope),
		Summary:                    summary,
		VulnerabilityNet:           vulnNet,
		VEXStatus:                  vexStatus,
		SigningIdentityStatus:      signingStatus,
		SigningIdentityFindings:    findings,
		RuntimeStatus:              runtimeStatus,
		RuntimeActiveStates:        activeStates,
		ExceptionReport:            exceptionReport,
		PublicationMode:            cfg.PublicationMode,
		StaleExceptionDays:         cfg.StaleExceptionDays,
	}, nil
}

func summarizeVEXStatements(statements []internalvex.Statement, filter internalvex.Filter) internalvex.StatusSummary {
	now := time.Now().UTC()
	summary := internalvex.StatusSummary{
		CountsByStatus: map[string]int{},
		AppliedFilters: map[string]string{
			"tenant_id":   filter.TenantID,
			"cluster_id":  filter.ClusterID,
			"environment": filter.Environment,
			"repo":        filter.Repo,
		},
	}
	for _, statement := range statements {
		if statement.RevokedAt != nil || !statement.Active {
			summary.RevokedCount++
			continue
		}
		if statement.ExpiresAt != nil && statement.ExpiresAt.After(now) && statement.ExpiresAt.Before(now.Add(72*time.Hour)) {
			summary.ExpiringCount++
		}
		if statement.ExpiresAt == nil || statement.ExpiresAt.After(now) {
			summary.ActiveCount++
		}
		summary.CountsByStatus[statement.Status]++
	}
	return summary
}

func filterEventsForTrustScope(events []audit.StoredEvent, scope trustScopeRequest) []audit.StoredEvent {
	filtered := make([]audit.StoredEvent, 0, len(events))
	for _, event := range events {
		if scope.TenantID != "" && event.TenantID != scope.TenantID {
			continue
		}
		if scope.ClusterID != "" && event.ClusterID != scope.ClusterID {
			continue
		}
		if scope.Environment != "" && event.Environment != scope.Environment {
			continue
		}
		if scope.Repo != "" && event.Repo != scope.Repo {
			continue
		}
		filtered = append(filtered, event)
	}
	return filtered
}

func parseTrustScopeRequestFromQuery(r *http.Request) trustScopeRequest {
	query := r.URL.Query()
	return normalizeTrustScopeRequest(trustScopeRequest{
		TenantID:    strings.TrimSpace(query.Get("tenant_id")),
		ClusterID:   strings.TrimSpace(query.Get("cluster_id")),
		Environment: strings.TrimSpace(query.Get("environment")),
		Repo:        strings.TrimSpace(query.Get("repo")),
		Format:      strings.TrimSpace(query.Get("format")),
	})
}

func normalizeTrustScopeRequest(request trustScopeRequest) trustScopeRequest {
	request.TenantID = strings.TrimSpace(request.TenantID)
	request.ClusterID = strings.TrimSpace(request.ClusterID)
	request.Environment = strings.TrimSpace(request.Environment)
	request.Repo = strings.TrimSpace(request.Repo)
	request.Format = strings.ToLower(strings.TrimSpace(request.Format))
	return request
}

func applyPrincipalTenantToTrustScopeRequest(principal auth.Principal, request trustScopeRequest) (trustScopeRequest, error) {
	tenantID, err := coerceTenantScope(principal, request.TenantID)
	if err != nil {
		return trustScopeRequest{}, err
	}
	request.TenantID = tenantID
	return normalizeTrustScopeRequest(request), nil
}

func trustScopeType(scope trustScopeRequest) string {
	switch {
	case scope.Repo != "":
		return "repository"
	case scope.ClusterID != "":
		return "cluster"
	case scope.TenantID != "":
		return "tenant"
	default:
		return "global"
	}
}

func trustScopeRef(scope trustScopeRequest) string {
	switch trustScopeType(scope) {
	case "repository":
		if scope.Environment != "" {
			return scope.Repo + "@" + scope.Environment
		}
		return scope.Repo
	case "cluster":
		if scope.Environment != "" {
			return scope.ClusterID + "@" + scope.Environment
		}
		return scope.ClusterID
	case "tenant":
		if scope.Environment != "" {
			return scope.TenantID + "@" + scope.Environment
		}
		return scope.TenantID
	default:
		if scope.Environment != "" {
			return "global@" + scope.Environment
		}
		return "global"
	}
}

func (s server) writeScorecardError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, audit.ErrInvalidFilter), errors.Is(err, audit.ErrInvalidException), errors.Is(err, internalvex.ErrInvalidFilter):
		status = http.StatusBadRequest
	case errors.Is(err, context.DeadlineExceeded):
		status = http.StatusGatewayTimeout
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}

func trustBoolPointer(value bool) *bool {
	return &value
}
