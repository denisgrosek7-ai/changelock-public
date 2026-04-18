package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/guidance"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

type aiInsightsResponse struct {
	Summary  guidance.Summary `json:"summary"`
	TopItems []guidance.Item  `json:"top_items"`
}

type aiVEXDraftRequest struct {
	TenantID      string `json:"tenant_id,omitempty"`
	Environment   string `json:"environment,omitempty"`
	Repo          string `json:"repo,omitempty"`
	ImageDigest   string `json:"image_digest,omitempty"`
	CVEID         string `json:"cve_id,omitempty"`
	ComponentName string `json:"component_name,omitempty"`
}

type aiVEXDraftResponse struct {
	Item  guidance.Item                 `json:"item"`
	Draft *guidance.VEXDraftSuggestion  `json:"draft,omitempty"`
}

type aiBreakGlassGuidanceRequest struct {
	ExceptionID string `json:"exception_id"`
	TenantID    string `json:"tenant_id,omitempty"`
}

type aiBreakGlassGuidanceResponse struct {
	Item     guidance.Item                  `json:"item"`
	Guidance *guidance.BreakGlassGuidance   `json:"guidance,omitempty"`
}

func loadAIGuidanceConfigFromEnv() (guidance.Config, error) {
	return guidance.ParseConfig(os.Getenv)
}

func (s server) aiGuidanceHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildAIGuidanceResponse(r.Context(), parseTrustScopeRequestFromQuery(r))
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) aiGuidanceByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	id := strings.TrimPrefix(strings.TrimSpace(r.URL.Path), "/v1/ai/guidance/")
	if id == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "guidance item not found"})
		return
	}
	response, err := s.buildAIGuidanceResponse(r.Context(), parseTrustScopeRequestFromQuery(r))
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	for _, item := range response.Items {
		if item.ID == id {
			httpjson.Write(w, http.StatusOK, item)
			return
		}
	}
	httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "guidance item not found"})
}

func (s server) aiInsightsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildAIGuidanceResponse(r.Context(), parseTrustScopeRequestFromQuery(r))
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	limit := 3
	if len(response.Items) < limit {
		limit = len(response.Items)
	}
	httpjson.Write(w, http.StatusOK, aiInsightsResponse{
		Summary:  response.Summary,
		TopItems: append([]guidance.Item(nil), response.Items[:limit]...),
	})
}

func (s server) aiVEXDraftsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request aiVEXDraftRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	tenantID, err := coerceTenantScope(principal, request.TenantID)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	request.TenantID = tenantID

	cfg, err := loadAIGuidanceConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	netResponse, err := s.computeVulnerabilityNetResponse(ctx, audit.VulnerabilityActiveFilter{
		CVEID:         strings.TrimSpace(request.CVEID),
		ImageDigest:   strings.TrimSpace(request.ImageDigest),
		ComponentName: strings.TrimSpace(request.ComponentName),
		TenantID:      strings.TrimSpace(request.TenantID),
		Environment:   strings.TrimSpace(request.Environment),
		Limit:         max(cfg.MaxItems, 50),
	}, "HIGH")
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	fact := guidance.InputFact{
		ID:              "vex-draft:" + firstNonEmpty(request.ImageDigest, request.CVEID, request.ComponentName),
		Category:        guidance.CategoryVulnerability,
		SourceComponent: "audit-writer",
		Severity:        vulnerabilitySeverity(netResponse.Findings),
		Summary:         "Contextual VEX draft suggestion",
		Detail:          "Derived from the current net actionable vulnerability posture.",
		RelatedReasonCodes: []string{
			audit.ScorecardReasonVulnActionable,
		},
		FindingRefs:   []string{"report:/v1/vulnerabilities/net"},
		EvidenceRefs:  []string{"report:/v1/vulnerabilities/net"},
		DocsRefs:      []string{"docs/vex-exploitability-ops.md", "docs/deeper-ai-guidance.md"},
		TenantID:      request.TenantID,
		Environment:   strings.TrimSpace(request.Environment),
		Repository:    strings.TrimSpace(request.Repo),
		Deterministic: true,
		Blocking:      netResponse.ActionableCount > 0,
		Metadata: map[string]string{
			"actionable_count":      itoa(netResponse.ActionableCount),
			"resolved_by_vex_count": itoa(netResponse.ResolvedByVEXCount),
		},
	}
	response := guidance.Build(guidance.Scope{
		ScopeType:   "repository",
		ScopeRef:    firstNonEmpty(strings.TrimSpace(request.Repo), strings.TrimSpace(request.ImageDigest), strings.TrimSpace(request.CVEID), "vex-draft"),
		TenantID:    request.TenantID,
		Environment: strings.TrimSpace(request.Environment),
		Repository:  strings.TrimSpace(request.Repo),
	}, []guidance.InputFact{fact}, cfg, timeNowUTC())
	if len(response.Items) == 0 || response.Items[0].VEXDraft == nil {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": "no VEX draft suggestion is available for the current scope"})
		return
	}
	httpjson.Write(w, http.StatusOK, aiVEXDraftResponse{
		Item:  response.Items[0],
		Draft: response.Items[0].VEXDraft,
	})
}

func (s server) aiBreakGlassGuidanceHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request aiBreakGlassGuidanceRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	tenantID, err := coerceTenantScope(principal, request.TenantID)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	request.TenantID = tenantID

	cfg, err := loadAIGuidanceConfigFromEnv()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	exception, err := s.store.GetException(ctx, strings.TrimSpace(request.ExceptionID))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrExceptionNotFound) {
			status = http.StatusNotFound
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	if err := ensureExceptionTenantAccess(principal, exception); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if exception.ExceptionType != audit.ExceptionTypeBreakGlass {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": "guidance is only available for BREAK_GLASS exceptions"})
		return
	}

	fact := guidance.InputFact{
		ID:              "break-glass:" + exception.ExceptionID,
		Category:        guidance.CategoryBreakGlass,
		SourceComponent: "audit-writer",
		Severity:        "medium",
		Summary:         "Break-glass guidance request",
		Detail:          firstNonEmpty(exception.Reason, "break-glass exception requires explicit cleanup review"),
		RelatedReasonCodes: []string{
			guidance.ReasonGuidanceBreakGlassActive,
		},
		FindingRefs:   []string{exception.ExceptionID},
		EvidenceRefs:  []string{"report:/v1/reports/exceptions"},
		DocsRefs:      []string{"docs/auth-rbac.md", "docs/deeper-ai-guidance.md"},
		ScopeType:     "repository",
		ScopeRef:      firstNonEmpty(exception.Repo, exception.ExceptionID),
		TenantID:      exception.TenantID,
		Environment:   exception.Environment,
		Repository:    exception.Repo,
		Deterministic: true,
		Metadata: map[string]string{
			"exception_type":         exception.ExceptionType,
			"active_exception_count": "1",
		},
	}
	response := guidance.Build(guidance.Scope{
		ScopeType:   "repository",
		ScopeRef:    firstNonEmpty(exception.Repo, exception.ExceptionID),
		TenantID:    exception.TenantID,
		Environment: exception.Environment,
		Repository:  exception.Repo,
	}, []guidance.InputFact{fact}, cfg, timeNowUTC())
	if len(response.Items) == 0 || response.Items[0].BreakGlassGuidance == nil {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": "no break-glass guidance is available for the current exception"})
		return
	}
	httpjson.Write(w, http.StatusOK, aiBreakGlassGuidanceResponse{
		Item:     response.Items[0],
		Guidance: response.Items[0].BreakGlassGuidance,
	})
}

func (s server) buildAIGuidanceResponse(ctx context.Context, scope trustScopeRequest) (guidance.Response, error) {
	cfg, err := loadAIGuidanceConfigFromEnv()
	if err != nil {
		return guidance.Response{}, err
	}
	scorecardCfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		return guidance.Response{}, err
	}
	input, err := s.collectTrustScorecardInput(ctx, scope, scorecardCfg)
	if err != nil {
		return guidance.Response{}, err
	}
	scorecard := audit.ComputeTrustScorecard(input)
	findings := audit.BuildHardeningReview(input, scorecard)
	facts := s.guidanceFactsFromTrustInput(input, scorecard, findings)
	return guidance.Build(guidance.Scope{
		ScopeType:   scorecard.ScopeType,
		ScopeRef:    scorecard.ScopeRef,
		TenantID:    scorecard.TenantID,
		ClusterID:   scorecard.ClusterID,
		Environment: scorecard.Environment,
		Repository:  scorecard.Repo,
	}, facts, cfg, timeNowUTC()), nil
}

func (s server) guidanceFactsFromTrustInput(input audit.TrustScorecardInput, scorecard audit.TrustScorecard, findings []audit.AuditFinding) []guidance.InputFact {
	facts := make([]guidance.InputFact, 0, len(findings)+len(scorecard.Metrics)+4)
	for _, finding := range findings {
		facts = append(facts, guidance.InputFact{
			ID:                 finding.ID,
			Category:           guidanceCategoryFromAudit(finding.Category),
			SourceComponent:    "audit-writer",
			RelatedReasonCodes: []string{finding.ReasonCode},
			FindingRefs:        []string{finding.ID},
			EvidenceRefs:       append([]string(nil), finding.EvidenceRefs...),
			DocsRefs:           guidanceDocsForReason(finding.ReasonCode),
			ScopeType:          scorecard.ScopeType,
			ScopeRef:           scorecard.ScopeRef,
			TenantID:           input.TenantID,
			ClusterID:          input.ClusterID,
			Environment:        input.Environment,
			Repository:         input.Repo,
			Severity:           finding.Severity,
			Summary:            finding.Category,
			Detail:             finding.ReasonDetail,
			Deterministic:      true,
			Metadata: map[string]string{
				"actionable_count":      itoa(input.VulnerabilityNet.ActionableCount),
				"resolved_by_vex_count": itoa(input.VulnerabilityNet.ResolvedByVEXCount),
				"quarantined_count":     itoa64(input.RuntimeStatus.Quarantined),
				"stale_exception_count": itoa(scorecard.StaleExceptionCount),
			},
		})
	}
	for _, metric := range scorecard.Metrics {
		if metric.Status == audit.TrustMetricStatusVerified && metric.ID != audit.ScorecardMetricVulnerability {
			continue
		}
		facts = append(facts, guidance.InputFact{
			ID:                 "metric:" + metric.ID,
			Category:           guidanceCategoryFromMetric(metric.ID),
			SourceComponent:    "scorecard",
			RelatedReasonCodes: []string{metric.ReasonCode},
			FindingRefs:        []string{metric.ID},
			EvidenceRefs:       append([]string(nil), metric.EvidenceRefs...),
			DocsRefs:           guidanceDocsForReason(metric.ReasonCode),
			ScopeType:          scorecard.ScopeType,
			ScopeRef:           scorecard.ScopeRef,
			TenantID:           input.TenantID,
			ClusterID:          input.ClusterID,
			Environment:        input.Environment,
			Repository:         input.Repo,
			Severity:           metricSeverity(metric.Status),
			Summary:            metric.Name,
			Detail:             metric.ReasonDetail,
			Deterministic:      true,
			Metadata: map[string]string{
				"actionable_count":      itoa(input.VulnerabilityNet.ActionableCount),
				"resolved_by_vex_count": itoa(input.VulnerabilityNet.ResolvedByVEXCount),
				"quarantined_count":     itoa64(input.RuntimeStatus.Quarantined),
				"stale_exception_count": itoa(scorecard.StaleExceptionCount),
			},
		})
	}
	breakGlassCount := 0
	for _, exception := range append(append([]audit.PolicyException{}, input.ExceptionReport.Active...), input.ExceptionReport.Pending...) {
		if exception.ExceptionType == audit.ExceptionTypeBreakGlass {
			breakGlassCount++
		}
	}
	if breakGlassCount > 0 {
		facts = append(facts, guidance.InputFact{
			ID:              "break-glass:" + scorecard.ScopeRef,
			Category:        guidance.CategoryBreakGlass,
			SourceComponent: "audit-writer",
			RelatedReasonCodes: []string{
				guidance.ReasonGuidanceBreakGlassActive,
			},
			FindingRefs:   []string{"report:/v1/reports/exceptions"},
			EvidenceRefs:  []string{"report:/v1/reports/exceptions"},
			DocsRefs:      []string{"docs/auth-rbac.md", "docs/deeper-ai-guidance.md"},
			ScopeType:     scorecard.ScopeType,
			ScopeRef:      scorecard.ScopeRef,
			TenantID:      input.TenantID,
			ClusterID:     input.ClusterID,
			Environment:   input.Environment,
			Repository:    input.Repo,
			Severity:      "medium",
			Summary:       "Active break-glass posture",
			Detail:        "At least one active or pending break-glass exception is present in this scope.",
			Deterministic: true,
			Metadata: map[string]string{
				"exception_type":         audit.ExceptionTypeBreakGlass,
				"active_exception_count": itoa(breakGlassCount),
			},
		})
	}
	return facts
}

func guidanceCategoryFromAudit(category string) string {
	switch strings.TrimSpace(strings.ToLower(category)) {
	case "vulnerabilities":
		return guidance.CategoryVulnerability
	case "signing":
		return guidance.CategorySigning
	case "runtime":
		return guidance.CategoryRuntime
	case "exceptions":
		return guidance.CategoryException
	case "policy":
		return guidance.CategoryPolicy
	case "evidence":
		return guidance.CategoryArtifact
	default:
		return guidance.CategoryContext
	}
}

func guidanceCategoryFromMetric(metricID string) string {
	switch metricID {
	case audit.ScorecardMetricVulnerability:
		return guidance.CategoryVulnerability
	case audit.ScorecardMetricSignerGovernance:
		return guidance.CategorySigning
	case audit.ScorecardMetricRuntimeHardening:
		return guidance.CategoryRuntime
	case audit.ScorecardMetricExceptionHygiene:
		return guidance.CategoryException
	case audit.ScorecardMetricArtifactIntegrity:
		return guidance.CategoryArtifact
	case audit.ScorecardMetricPolicyEvidence:
		return guidance.CategoryPolicy
	default:
		return guidance.CategoryContext
	}
}

func guidanceDocsForReason(reason string) []string {
	switch {
	case strings.Contains(reason, "vulnerability"), strings.Contains(reason, "vex"):
		return []string{"docs/vex-exploitability-ops.md", "docs/vulnerability-ops.md", "docs/deeper-ai-guidance.md"}
	case strings.Contains(reason, "signer"):
		return []string{"docs/signing-identity-monitoring.md", "docs/deeper-ai-guidance.md"}
	case strings.Contains(reason, "runtime"), strings.Contains(reason, "quarantine"):
		return []string{"docs/runtime-closed-loop-hardening.md", "docs/deeper-ai-guidance.md"}
	case strings.Contains(reason, "exception"), strings.Contains(reason, "break_glass"):
		return []string{"docs/auth-rbac.md", "docs/deeper-ai-guidance.md"}
	case strings.Contains(reason, "artifact"), strings.Contains(reason, "transparency"):
		return []string{"docs/immutable-evidence-transparency-log.md", "docs/audit-evidence.md", "docs/deeper-ai-guidance.md"}
	case strings.Contains(reason, "policy"), strings.Contains(reason, "manifest"):
		return []string{"docs/shift-left-integration.md", "docs/developer-preflight-cli.md", "docs/deeper-ai-guidance.md"}
	default:
		return []string{"docs/deeper-ai-guidance.md"}
	}
}

func metricSeverity(status string) string {
	switch status {
	case audit.TrustMetricStatusGap:
		return "high"
	case audit.TrustMetricStatusPartial:
		return "medium"
	default:
		return "low"
	}
}

func vulnerabilitySeverity(findings []audit.VulnerabilityFinding) string {
	result := "medium"
	for _, finding := range findings {
		switch strings.ToUpper(strings.TrimSpace(finding.Severity)) {
		case "CRITICAL":
			return "critical"
		case "HIGH":
			result = "high"
		}
	}
	return result
}

func itoa(value int) string {
	return strconv.Itoa(value)
}

func itoa64(value int64) string {
	return strconv.FormatInt(value, 10)
}

func timeNowUTC() time.Time {
	return time.Now().UTC()
}
