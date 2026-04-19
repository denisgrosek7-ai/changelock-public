package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

type vexStatementsResponse struct {
	Statements []internalvex.Statement `json:"statements"`
}

type vexStatementResponse struct {
	Status    string                `json:"status"`
	Statement internalvex.Statement `json:"statement"`
}

type vexImportResponse struct {
	Status string                 `json:"status"`
	Result internalvex.ImportResult `json:"result"`
}

func loadVEXConfigFromEnv() (internalvex.Config, error) {
	return internalvex.ParseConfig(os.Getenv)
}

func importVEXDirectory(ctx context.Context, store audit.Store, config internalvex.Config, actor string) (int, error) {
	if strings.TrimSpace(config.ImportDir) == "" {
		return 0, nil
	}
	entries, err := os.ReadDir(config.ImportDir)
	if err != nil {
		return 0, err
	}
	imported := 0
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".json") {
			continue
		}
		path := filepath.Join(config.ImportDir, entry.Name())
		payload, err := os.ReadFile(path)
		if err != nil {
			return imported, err
		}
		createRequests, _, err := internalvex.ParseIngestRequest(internalvex.IngestRequest{
			SourceRef: path,
			Payload:   payload,
		})
		if err != nil {
			return imported, err
		}
		for _, request := range createRequests {
			if _, err := store.CreateVEXStatement(ctx, request, actor); err != nil {
				return imported, err
			}
			imported++
		}
	}
	return imported, nil
}

func (s server) vexStatementsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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
		filter, err := parseVEXFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		statements, err := s.store.ListVEXStatements(ctx, filter)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrInvalidFilter) || errors.Is(err, internalvex.ErrInvalidFilter) {
				status = http.StatusBadRequest
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, vexStatementsResponse{Statements: statements})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var request internalvex.CreateRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request, err := applyPrincipalTenantToVEXRequest(principal, request)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		statement, err := s.store.CreateVEXStatement(ctx, request, principal.Subject)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, internalvex.ErrInvalidStatement) || errors.Is(err, audit.ErrInvalidException) {
				status = http.StatusBadRequest
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		s.writeVEXAuditEvent(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeVEXStatementRecorded, "vex statement recorded", statement)
		httpjson.Write(w, http.StatusCreated, vexStatementResponse{Status: "created", Statement: statement})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) vexIngestHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request internalvex.IngestRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request.Scope = applyPrincipalTenantToVEXScope(principal, request.Scope)
	statements, format, err := internalvex.ParseIngestRequest(request)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	result := internalvex.ImportResult{
		SourceFormat: format,
		Imported:     0,
		Statements:   []internalvex.Statement{},
	}
	for _, statementRequest := range statements {
		statement, err := s.store.CreateVEXStatement(ctx, statementRequest, principal.Subject)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		result.Imported++
		result.Statements = append(result.Statements, statement)
		s.writeVEXAuditEvent(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeVEXStatementRecorded, "vex statement ingested", statement)
	}
	httpjson.Write(w, http.StatusCreated, vexImportResponse{Status: "imported", Result: result})
}

func (s server) vexStatementByIDHandler(w http.ResponseWriter, r *http.Request) {
	statementID, action, err := vexPath(r.URL.Path)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	switch {
	case r.Method == http.MethodGet && action == "":
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		statement, err := s.store.GetVEXStatement(ctx, statementID)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrExceptionNotFound) {
				status = http.StatusNotFound
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		if err := ensureVEXTenantAccess(ctx, s.store, principal, statement); err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, statement)
	case r.Method == http.MethodPost && action == "revoke":
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		statement, err := s.store.GetVEXStatement(ctx, statementID)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrExceptionNotFound) {
				status = http.StatusNotFound
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		if err := ensureVEXTenantAccess(ctx, s.store, principal, statement); err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		statement, err = s.store.RevokeVEXStatement(ctx, statementID, principal.Subject)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		s.writeVEXAuditEvent(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeVEXStatementRevoked, "vex statement revoked", statement)
		httpjson.Write(w, http.StatusOK, vexStatementResponse{Status: "revoked", Statement: statement})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) vexStatusHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseVEXFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter.Limit = 1000
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	statements, err := s.store.ListVEXStatements(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	now := time.Now().UTC()
	summary := internalvex.StatusSummary{
		CountsByStatus: map[string]int{},
		AppliedFilters: map[string]string{
			"tenant_id":        filter.TenantID,
			"cluster_id":       filter.ClusterID,
			"source_format":    filter.SourceFormat,
			"vulnerability_id": filter.VulnerabilityID,
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
	httpjson.Write(w, http.StatusOK, summary)
}

func (s server) vulnerabilityNetHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseVulnerabilityActiveFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	severityThreshold := strings.ToUpper(strings.TrimSpace(r.URL.Query().Get("severity_threshold")))
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.computeVulnerabilityNetResponse(ctx, filter, severityThreshold)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) computeVulnerabilityNetResponse(ctx context.Context, filter audit.VulnerabilityActiveFilter, severityThreshold string) (audit.VulnerabilityNetResponse, error) {
	rawFindings, err := s.store.ListActiveVulnerabilities(ctx, audit.VulnerabilityActiveFilter{
		Severity:          filter.Severity,
		CVEID:             filter.CVEID,
		ImageDigest:       filter.ImageDigest,
		ComponentName:     filter.ComponentName,
		TenantID:          filter.TenantID,
		Environment:       filter.Environment,
		Limit:             filter.Limit,
		IncludeSuppressed: true,
	})
	if err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	actionable, err := s.store.ListActiveVulnerabilities(ctx, filter)
	if err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	response := audit.VulnerabilityNetResponse{
		RawCount:        len(rawFindings),
		ActionableCount: len(actionable),
		Findings:        actionable,
		AppliedFilters: map[string]string{
			"severity":       filter.Severity,
			"cve_id":         filter.CVEID,
			"image_digest":   filter.ImageDigest,
			"component_name": filter.ComponentName,
			"tenant_id":      filter.TenantID,
			"environment":    filter.Environment,
		},
	}
	for _, finding := range rawFindings {
		if finding.VEX == nil {
			continue
		}
		switch finding.VEX.Status {
		case internalvex.StatusNotAffected:
			response.ResolvedByVEXCount++
		case internalvex.StatusUnderInvestigation:
			response.UnderInvestigationCount++
		}
	}
	if severityThreshold != "" {
		response.SeverityThreshold = severityThreshold
		for _, finding := range actionable {
			if vulnerabilitySeverityRank(finding.Severity) >= vulnerabilitySeverityRank(severityThreshold) {
				response.ThresholdBreached = true
				break
			}
		}
	}
	return response, nil
}

func parseVEXFilter(r *http.Request) (internalvex.Filter, error) {
	activeParam := strings.TrimSpace(r.URL.Query().Get("active"))
	var active *bool
	if activeParam != "" {
		value, err := strconv.ParseBool(activeParam)
		if err != nil {
			return internalvex.Filter{}, err
		}
		active = &value
	}
	return internalvex.NormalizeFilter(internalvex.Filter{
		VulnerabilityID: r.URL.Query().Get("vulnerability_id"),
		ImageDigest:     r.URL.Query().Get("image_digest"),
		PackageName:     r.URL.Query().Get("package_name"),
		PURL:            r.URL.Query().Get("purl"),
		Repo:            r.URL.Query().Get("repo"),
		Workload:        r.URL.Query().Get("workload"),
		TenantID:        r.URL.Query().Get("tenant_id"),
		ClusterID:       r.URL.Query().Get("cluster_id"),
		Environment:     r.URL.Query().Get("environment"),
		Namespace:       r.URL.Query().Get("namespace"),
		SourceFormat:    r.URL.Query().Get("source_format"),
		Status:          r.URL.Query().Get("status"),
		SourceRef:       r.URL.Query().Get("source_ref"),
		Active:          active,
		Limit:           parseIntOrDefault(r.URL.Query().Get("limit"), 100),
	})
}

func applyPrincipalTenantToVEXRequest(principal auth.Principal, request internalvex.CreateRequest) (internalvex.CreateRequest, error) {
	request.Scope = applyPrincipalTenantToVEXScope(principal, request.Scope)
	if principal.TenantID != "" && !principal.GlobalScope && request.Scope.TenantID != principal.TenantID {
		return request, auth.ErrInsufficientPermissions
	}
	return request, nil
}

func applyPrincipalTenantToVEXScope(principal auth.Principal, scope internalvex.Scope) internalvex.Scope {
	if principal.TenantID != "" && !principal.GlobalScope {
		scope.TenantID = principal.TenantID
	}
	return scope
}

func ensureVEXTenantAccess(ctx context.Context, store audit.Store, principal auth.Principal, statement internalvex.Statement) error {
	if principal.GlobalScope || principal.TenantID == "" {
		return nil
	}
	if statement.Scope.TenantID != "" {
		if statement.Scope.TenantID != principal.TenantID {
			return auth.ErrInsufficientPermissions
		}
		return nil
	}
	if statement.Scope.ImageDigest != "" {
		return ensureDigestTenantAccess(ctx, store, principal, statement.Scope.ImageDigest)
	}
	return auth.ErrInsufficientPermissions
}

func vexPath(path string) (int64, string, error) {
	trimmed := strings.TrimPrefix(strings.TrimSpace(path), "/v1/vex/")
	parts := strings.Split(strings.Trim(trimmed, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return 0, "", fmt.Errorf("vex statement id is required")
	}
	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil || id <= 0 {
		return 0, "", fmt.Errorf("invalid vex statement id")
	}
	if len(parts) == 1 {
		return id, "", nil
	}
	return id, parts[1], nil
}

func (s server) writeVEXAuditEvent(ctx context.Context, requestID string, actor string, eventType string, reason string, statement internalvex.Statement) {
	_, _ = s.store.Ingest(ctx, audit.Event{
		RequestID:   requestID,
		Component:   "audit-writer",
		EventType:   eventType,
		Actor:       actor,
		TenantID:    statement.Scope.TenantID,
		ClusterID:   statement.Scope.ClusterID,
		Repo:        statement.Scope.Repo,
		Environment: statement.Scope.Environment,
		Namespace:   statement.Scope.Namespace,
		Workload:    statement.Scope.Workload,
		Digest:      statement.Scope.ImageDigest,
		CVEID:       statement.VulnerabilityID,
		Decision:    audit.DecisionAllow,
		Reasons:     []string{reason},
	})
}

func vulnerabilitySeverityRank(value string) int {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "CRITICAL":
		return 5
	case "HIGH":
		return 4
	case "MEDIUM":
		return 3
	case "LOW":
		return 2
	case "UNKNOWN":
		return 1
	default:
		return 0
	}
}
