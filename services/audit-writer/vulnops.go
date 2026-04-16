package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	internalvulnops "github.com/denisgrosek/changelock/internal/vulnops"
)

type sbomComponentsResponse struct {
	Components []audit.SBOMComponent `json:"components"`
}

type vulnerabilitiesResponse struct {
	Findings []audit.VulnerabilityFinding `json:"findings"`
}

type vulnerabilityDecisionsResponse struct {
	Decisions []audit.VulnerabilityDecision `json:"decisions"`
}

type vulnerabilityDecisionActionResponse struct {
	Status   string                      `json:"status"`
	Decision audit.VulnerabilityDecision `json:"decision"`
}

type vulnOpsRuntime struct {
	config  internalvulnops.Config
	scanner internalvulnops.Scanner
}

func loadVulnOpsRuntimeFromEnv() (*vulnOpsRuntime, error) {
	config, err := internalvulnops.ConfigFromEnv()
	if err != nil {
		return nil, err
	}
	scanner, err := internalvulnops.NewScanner(config)
	if err != nil {
		return nil, err
	}
	return &vulnOpsRuntime{
		config:  config,
		scanner: scanner,
	}, nil
}

func (v *vulnOpsRuntime) start(ctx context.Context, store audit.Store) {
	if v == nil || !v.config.Enabled || v.scanner == nil {
		return
	}
	go func() {
		ticker := time.NewTicker(v.config.ScanInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				scanCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
				if _, err := v.rescan(scanCtx, store, "vulnops-worker", audit.VulnerabilityScanModePeriodic, audit.NewRequestID(), audit.VulnerabilityRescanRequest{}); err != nil {
					log.Printf("vulnops periodic rescan failed: %v", err)
				}
				cancel()
			}
		}
	}()
}

func (v *vulnOpsRuntime) sbomIngestEnabled() bool {
	if v == nil {
		return true
	}
	return v.config.SBOMIngestEnabled
}

func (v *vulnOpsRuntime) rescansEnabled() bool {
	return v != nil && v.config.Enabled && v.scanner != nil
}

func (v *vulnOpsRuntime) rescan(
	ctx context.Context,
	store audit.Store,
	actor string,
	scanMode string,
	requestID string,
	request audit.VulnerabilityRescanRequest,
) (audit.VulnerabilityRescanResponse, error) {
	if !v.rescansEnabled() {
		return audit.VulnerabilityRescanResponse{}, errors.New("vulnerability rescanning is disabled")
	}

	targets, err := v.resolveTargets(ctx, store, request)
	if err != nil {
		return audit.VulnerabilityRescanResponse{}, err
	}
	if len(targets) == 0 {
		return audit.VulnerabilityRescanResponse{
			Status:         "no_targets",
			ScannedDigests: []string{},
			ScanRuns:       0,
		}, nil
	}

	response := audit.VulnerabilityRescanResponse{
		Status:         "completed",
		ScannedDigests: make([]string, 0, len(targets)),
		ScanRuns:       0,
	}
	failures := []string{}
	for _, target := range targets {
		response.ScannedDigests = append(response.ScannedDigests, target.ImageDigest)

		result, scanErr := v.scanner.ScanDigest(ctx, target)
		if scanErr != nil {
			now := time.Now().UTC()
			failedRun, recordErr := store.RecordVulnerabilityScan(ctx, audit.VulnerabilityScanRequest{
				ImageDigest: target.ImageDigest,
				ImageRef:    target.ImageRef,
				Scanner:     v.config.Scanner,
				ScanMode:    scanMode,
				StartedAt:   now,
				CompletedAt: &now,
				Status:      audit.VulnerabilityScanStatusFailed,
				Summary:     errorSummaryJSON(scanErr),
				SourceRef:   firstNonEmpty(target.ImageRef, target.ImageDigest),
			})
			if recordErr == nil {
				response.ScanRuns++
				v.writeScanAuditEvent(ctx, store, requestID, actor, target, failedRun, scanErr)
			}
			failures = append(failures, scanErr.Error())
			continue
		}

		persisted, recordErr := store.RecordVulnerabilityScan(ctx, audit.VulnerabilityScanRequest{
			ImageDigest: result.ImageDigest,
			ImageRef:    result.ImageRef,
			Scanner:     result.Scanner,
			ScanMode:    scanMode,
			StartedAt:   result.StartedAt,
			CompletedAt: &result.CompletedAt,
			Status:      result.Status,
			Summary:     result.Summary,
			SourceRef:   result.SourceRef,
			Findings:    result.Findings,
		})
		if recordErr != nil {
			failures = append(failures, recordErr.Error())
			continue
		}

		response.ScanRuns++
		v.writeScanAuditEvent(ctx, store, requestID, actor, target, persisted, nil)
		if persisted.HadPriorSuccessfulRun && len(persisted.NewFindings) > 0 {
			for _, finding := range persisted.NewFindings {
				v.writeDriftAuditEvent(ctx, store, requestID, actor, target, finding)
			}
		}
	}

	switch {
	case len(failures) > 0 && response.ScanRuns == 0:
		return response, errors.New(strings.Join(failures, "; "))
	case len(failures) > 0:
		response.Status = "completed_with_errors"
	default:
		response.Status = "completed"
	}
	return response, nil
}

func (v *vulnOpsRuntime) resolveTargets(ctx context.Context, store audit.Store, request audit.VulnerabilityRescanRequest) ([]audit.ActiveDigestRef, error) {
	request.ImageDigest = strings.TrimSpace(request.ImageDigest)
	request.ImageRef = strings.TrimSpace(request.ImageRef)
	if request.ImageRef != "" && request.ImageDigest == "" {
		return nil, fmt.Errorf("%w: image_digest is required when image_ref is supplied", audit.ErrInvalidException)
	}
	if request.ImageDigest != "" {
		targets, err := store.ListActiveDigests(ctx, 30, 500)
		if err != nil {
			return nil, err
		}
		for _, target := range targets {
			if target.ImageDigest == request.ImageDigest {
				if request.ImageRef != "" {
					target.ImageRef = request.ImageRef
				}
				return []audit.ActiveDigestRef{target}, nil
			}
		}
		return []audit.ActiveDigestRef{{
			ImageDigest: request.ImageDigest,
			ImageRef:    request.ImageRef,
		}}, nil
	}
	return store.ListActiveDigests(ctx, 30, 500)
}

func (v *vulnOpsRuntime) writeScanAuditEvent(
	ctx context.Context,
	store audit.Store,
	requestID string,
	actor string,
	target audit.ActiveDigestRef,
	result audit.VulnerabilityScanIngestResult,
	scanErr error,
) {
	scope := primaryScope(target)
	decision := audit.DecisionAllow
	reasons := []string{fmt.Sprintf("%s scan completed with %d active findings", result.Run.Scanner, len(result.Findings))}
	if scanErr != nil {
		decision = audit.DecisionError
		reasons = []string{scanErr.Error()}
	}
	_, _ = store.Ingest(ctx, audit.Event{
		RequestID:   requestID,
		Component:   "audit-writer",
		EventType:   audit.EventTypeVulnerabilityScanResult,
		Actor:       actor,
		TenantID:    scope.TenantID,
		Repo:        firstNonEmpty(scope.Repo, target.Repo),
		Environment: scope.Environment,
		Namespace:   scope.Namespace,
		Workload:    scope.Workload,
		Image:       firstNonEmpty(target.ImageRef, scope.Image),
		Digest:      target.ImageDigest,
		Decision:    decision,
		Reasons:     reasons,
	})
}

func (v *vulnOpsRuntime) writeDriftAuditEvent(
	ctx context.Context,
	store audit.Store,
	requestID string,
	actor string,
	target audit.ActiveDigestRef,
	finding audit.VulnerabilityFinding,
) {
	scope := primaryScope(target)
	reason := fmt.Sprintf("new %s vulnerability %s in %s", strings.ToLower(firstNonEmpty(finding.Severity, "unknown")), finding.CVEID, firstNonEmpty(finding.PackageName, finding.PURL, "unknown component"))
	_, _ = store.Ingest(ctx, audit.Event{
		RequestID:   requestID,
		Component:   "audit-writer",
		EventType:   audit.EventTypeVulnerabilityDriftDetected,
		Actor:       actor,
		TenantID:    scope.TenantID,
		Repo:        firstNonEmpty(scope.Repo, target.Repo),
		Environment: scope.Environment,
		Namespace:   scope.Namespace,
		Workload:    scope.Workload,
		Image:       firstNonEmpty(target.ImageRef, scope.Image),
		Digest:      target.ImageDigest,
		CVEID:       finding.CVEID,
		Decision:    audit.DecisionDeny,
		Reasons:     []string{reason},
	})
}

func (s server) sbomIngestHandler(w http.ResponseWriter, r *http.Request) {
	_, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin, auth.RoleService)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	if s.vulnOps != nil && !s.vulnOps.sbomIngestEnabled() {
		httpjson.Write(w, http.StatusServiceUnavailable, map[string]string{"error": "sbom ingestion is disabled"})
		return
	}

	var request audit.SBOMIngestRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	result, err := s.store.IngestSBOM(ctx, request)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusCreated, result)
}

func (s server) sbomImageHandler(w http.ResponseWriter, r *http.Request) {
	_, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	imageDigest, err := imageDigestFromSBOMPath(r.URL.Path)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	limit := parseIntOrDefault(r.URL.Query().Get("limit"), 100)

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	result, err := s.store.GetSBOMImage(ctx, imageDigest, limit)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, audit.ErrInvalidFilter):
			status = http.StatusBadRequest
		case errors.Is(err, audit.ErrExceptionNotFound):
			status = http.StatusNotFound
		case errors.Is(err, context.DeadlineExceeded):
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, result)
}

func (s server) sbomComponentsSearchHandler(w http.ResponseWriter, r *http.Request) {
	_, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter := audit.SBOMComponentSearchFilter{
		ComponentName: r.URL.Query().Get("component_name"),
		PURL:          r.URL.Query().Get("purl"),
		ImageDigest:   r.URL.Query().Get("image_digest"),
		Limit:         parseIntOrDefault(r.URL.Query().Get("limit"), 50),
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	components, err := s.store.SearchSBOMComponents(ctx, filter)
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
	httpjson.Write(w, http.StatusOK, sbomComponentsResponse{Components: components})
}

func (s server) activeVulnerabilitiesHandler(w http.ResponseWriter, r *http.Request) {
	_, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseVulnerabilityActiveFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	findings, err := s.store.ListActiveVulnerabilities(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, vulnerabilitiesResponse{Findings: findings})
}

func (s server) vulnerabilityBlastRadiusHandler(w http.ResponseWriter, r *http.Request) {
	_, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseVulnerabilityBlastRadiusFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	result, err := s.store.VulnerabilityBlastRadius(ctx, filter)
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
	httpjson.Write(w, http.StatusOK, result)
}

func (s server) vulnerabilityTimelineHandler(w http.ResponseWriter, r *http.Request) {
	_, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseVulnerabilityTimelineFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	result, err := s.store.VulnerabilityTimeline(ctx, filter)
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
	httpjson.Write(w, http.StatusOK, result)
}

func (s server) vulnerabilityDecisionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		filter, err := parseVulnerabilityDecisionFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		decisions, err := s.store.ListVulnerabilityDecisions(ctx, filter)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, context.DeadlineExceeded) {
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, vulnerabilityDecisionsResponse{Decisions: decisions})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var request audit.VulnerabilityDecisionCreateRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		decision, err := s.store.CreateVulnerabilityDecision(ctx, request, principal.Subject)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrInvalidException) {
				status = http.StatusBadRequest
			} else if errors.Is(err, context.DeadlineExceeded) {
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		s.writeVulnerabilityDecisionAuditEvent(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeVulnerabilityDecisionRecorded, "vulnerability decision recorded", decision)
		httpjson.Write(w, http.StatusCreated, vulnerabilityDecisionActionResponse{Status: "created", Decision: decision})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) vulnerabilityDecisionByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest

	decisionID, action, err := vulnerabilityDecisionPath(r.URL.Path)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if action != "deactivate" {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	decision, err := s.store.DeactivateVulnerabilityDecision(ctx, decisionID)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, audit.ErrExceptionNotFound):
			status = http.StatusNotFound
		case errors.Is(err, context.DeadlineExceeded):
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	s.writeVulnerabilityDecisionAuditEvent(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeVulnerabilityDecisionDeactivated, "vulnerability decision deactivated", decision)
	httpjson.Write(w, http.StatusOK, vulnerabilityDecisionActionResponse{Status: "deactivated", Decision: decision})
}

func (s server) vulnerabilityRescanHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin, auth.RoleService)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	if s.vulnOps == nil || !s.vulnOps.rescansEnabled() {
		httpjson.Write(w, http.StatusServiceUnavailable, map[string]string{"error": "vulnerability rescanning is disabled"})
		return
	}

	var request audit.VulnerabilityRescanRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Minute)
	defer cancel()

	response, err := s.vulnOps.rescan(ctx, s.store, principal.Subject, audit.VulnerabilityScanModeManual, requestIDFromHeader(r), request)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func parseVulnerabilityActiveFilter(r *http.Request) (audit.VulnerabilityActiveFilter, error) {
	query := r.URL.Query()
	filter := audit.VulnerabilityActiveFilter{
		Severity:      query.Get("severity"),
		CVEID:         query.Get("cve_id"),
		ImageDigest:   query.Get("image_digest"),
		ComponentName: query.Get("component_name"),
		TenantID:      query.Get("tenant_id"),
		Environment:   query.Get("environment"),
		Limit:         parseIntOrDefault(query.Get("limit"), 50),
	}
	if raw := strings.TrimSpace(query.Get("include_suppressed")); raw != "" {
		value, err := strconv.ParseBool(raw)
		if err != nil {
			return audit.VulnerabilityActiveFilter{}, errors.New("include_suppressed must be a boolean")
		}
		filter.IncludeSuppressed = value
	}
	return audit.NormalizeVulnerabilityActiveFilter(filter), nil
}

func parseVulnerabilityBlastRadiusFilter(r *http.Request) (audit.VulnerabilityBlastRadiusFilter, error) {
	return audit.NormalizeVulnerabilityBlastRadiusFilter(audit.VulnerabilityBlastRadiusFilter{
		CVEID:         r.URL.Query().Get("cve_id"),
		ComponentName: r.URL.Query().Get("component_name"),
		PURL:          r.URL.Query().Get("purl"),
		Limit:         parseIntOrDefault(r.URL.Query().Get("limit"), 50),
	})
}

func parseVulnerabilityTimelineFilter(r *http.Request) (audit.VulnerabilityTimelineFilter, error) {
	return audit.NormalizeVulnerabilityTimelineFilter(audit.VulnerabilityTimelineFilter{
		ImageDigest: r.URL.Query().Get("image_digest"),
		CVEID:       r.URL.Query().Get("cve_id"),
		WindowDays:  parseIntOrDefault(r.URL.Query().Get("window_days"), 30),
	})
}

func parseVulnerabilityDecisionFilter(r *http.Request) (audit.VulnerabilityDecisionFilter, error) {
	filter := audit.VulnerabilityDecisionFilter{
		ImageDigest: r.URL.Query().Get("image_digest"),
		CVEID:       r.URL.Query().Get("cve_id"),
		Limit:       parseIntOrDefault(r.URL.Query().Get("limit"), 50),
	}
	if raw := strings.TrimSpace(r.URL.Query().Get("active")); raw != "" {
		value, err := strconv.ParseBool(raw)
		if err != nil {
			return audit.VulnerabilityDecisionFilter{}, errors.New("active must be a boolean")
		}
		filter.Active = &value
	}
	return audit.NormalizeVulnerabilityDecisionFilter(filter), nil
}

func imageDigestFromSBOMPath(path string) (string, error) {
	raw := strings.TrimPrefix(path, "/v1/sbom/images/")
	raw = strings.TrimSpace(strings.Trim(raw, "/"))
	if raw == "" {
		return "", errors.New("image_digest path segment is required")
	}
	value, err := url.PathUnescape(raw)
	if err != nil {
		return "", errors.New("invalid image_digest path segment")
	}
	return strings.TrimSpace(value), nil
}

func vulnerabilityDecisionPath(path string) (int64, string, error) {
	raw := strings.TrimPrefix(path, "/v1/vulnerabilities/decisions/")
	raw = strings.TrimSpace(strings.Trim(raw, "/"))
	if raw == "" {
		return 0, "", errors.New("decision id path segment is required")
	}
	parts := strings.Split(raw, "/")
	if len(parts) != 2 {
		return 0, "", errors.New("invalid vulnerability decision path")
	}
	value, err := url.PathUnescape(parts[0])
	if err != nil {
		return 0, "", errors.New("invalid vulnerability decision id")
	}
	id, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64)
	if err != nil || id <= 0 {
		return 0, "", errors.New("invalid vulnerability decision id")
	}
	return id, strings.TrimSpace(parts[1]), nil
}

func parseIntOrDefault(raw string, fallback int) int {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return fallback
	}
	value, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return value
}

func primaryScope(target audit.ActiveDigestRef) audit.ActiveWorkloadRef {
	if len(target.Scopes) == 0 {
		return audit.ActiveWorkloadRef{}
	}
	return target.Scopes[0]
}

func errorSummaryJSON(err error) json.RawMessage {
	encoded, _ := json.Marshal(map[string]string{"error": err.Error()})
	return encoded
}

func (s server) writeVulnerabilityDecisionAuditEvent(ctx context.Context, requestID, actor, eventType, reason string, decision audit.VulnerabilityDecision) {
	_, _ = s.store.Ingest(ctx, audit.Event{
		RequestID: requestID,
		Component: "audit-writer",
		EventType: eventType,
		Actor:     actor,
		Digest:    decision.ImageDigest,
		CVEID:     decision.CVEID,
		Decision:  audit.DecisionAllow,
		Reasons: []string{
			decision.Decision,
			reason,
			decision.Justification,
		},
	})
}
