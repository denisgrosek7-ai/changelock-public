package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	intelligencecore "github.com/denisgrosek/changelock/internal/intelligence"
	supplychaincore "github.com/denisgrosek/changelock/internal/supplychain"
	vulnerabilitycore "github.com/denisgrosek/changelock/internal/vulnerability"
)

const (
	phase3IntelligenceComponent         = "phase3-intelligence-manager"
	phase3IntelligencePayloadSchema     = "3.intelligence_event_payload.v1"
	phase3VulnerabilityListSchema       = "3.intelligence_vulnerability_list.v1"
	phase3SupplyChainListSchema         = "3.intelligence_supply_chain_list.v1"
	phase3StrategicQuerySchema          = "3.intelligence_query_list.v1"
	phase3StrategicAssessmentListSchema = "3.intelligence_strategic_list.v1"
	phase3ProofsSchema                  = "3.intelligence_phase3_proofs.v1"
	phase3ProofStateIncomplete          = "phase3_core_incomplete"
	phase3ProofStateActive              = "phase3_core_slice_active"
	phase3VulnerabilityStateActive      = "vulnerability_relevance_active"
	phase3SupplyChainStateActive        = "supply_chain_pattern_active"
	phase3StrategicStateActive          = "strategic_support_active"
	phase3QueryStateActive              = "grounded_query_active"
)

var errPhase3SubjectRequired = errors.New("subject_ref is required for bounded same-subject Phase 3 strategic operations")

type phase3IntelligencePayload struct {
	SchemaVersion string                                `json:"schema_version"`
	Vulnerability *vulnerabilitycore.RelevanceVerdict   `json:"vulnerability,omitempty"`
	SupplyChain   *supplychaincore.PatternVerdict       `json:"supply_chain,omitempty"`
	Strategic     *intelligencecore.StrategicAssessment `json:"strategic,omitempty"`
	Query         *intelligencecore.QueryResponse       `json:"query,omitempty"`
}

type phase3IntelligenceFilter struct {
	ClusterID       string
	TenantID        string
	Environment     string
	Repo            string
	SubjectRef      string
	VulnerabilityID string
	PackageName     string
	Limit           int
}

type phase3VulnerabilityEvaluateRequest struct {
	Input vulnerabilitycore.Input `json:"input"`
}

type phase3VulnerabilityListResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	CurrentState  string                               `json:"current_state"`
	Items         []vulnerabilitycore.RelevanceVerdict `json:"items"`
	Limitations   []string                             `json:"limitations,omitempty"`
}

type phase3VulnerabilityEvaluateResponse struct {
	Status  string                             `json:"status"`
	Verdict vulnerabilitycore.RelevanceVerdict `json:"verdict"`
}

type phase3SupplyChainEvaluateRequest struct {
	Input supplychaincore.Input `json:"input"`
}

type phase3SupplyChainListResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	CurrentState  string                           `json:"current_state"`
	Items         []supplychaincore.PatternVerdict `json:"items"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type phase3SupplyChainEvaluateResponse struct {
	Status  string                         `json:"status"`
	Pattern supplychaincore.PatternVerdict `json:"pattern"`
}

type phase3StrategicSimulateRequest struct {
	SubjectRef          string   `json:"subject_ref,omitempty"`
	CandidateAction     string   `json:"candidate_action,omitempty"`
	DelayDays           int      `json:"delay_days,omitempty"`
	EffortBand          string   `json:"effort_band,omitempty"`
	BlastRadiusScore    int      `json:"blast_radius_score,omitempty"`
	VulnerabilityID     string   `json:"vulnerability_id,omitempty"`
	PackageName         string   `json:"package_name,omitempty"`
	ObservedFacts       []string `json:"observed_facts,omitempty"`
	InferredConclusions []string `json:"inferred_conclusions,omitempty"`
	EvidenceRefs        []string `json:"evidence_refs,omitempty"`
}

type phase3StrategicSimulateResponse struct {
	Status     string                               `json:"status"`
	Assessment intelligencecore.StrategicAssessment `json:"assessment"`
}

type phase3StrategicQueryRequest struct {
	Query           string `json:"query"`
	SubjectRef      string `json:"subject_ref,omitempty"`
	VulnerabilityID string `json:"vulnerability_id,omitempty"`
	PackageName     string `json:"package_name,omitempty"`
}

type phase3StrategicQueryResponse struct {
	Status   string                         `json:"status"`
	Response intelligencecore.QueryResponse `json:"response"`
}

type phase3StrategicAssessmentListResponse struct {
	SchemaVersion string                                 `json:"schema_version"`
	CurrentState  string                                 `json:"current_state"`
	Items         []intelligencecore.StrategicAssessment `json:"items"`
	Limitations   []string                               `json:"limitations,omitempty"`
}

type phase3StrategicQueryListResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	CurrentState  string                           `json:"current_state"`
	Items         []intelligencecore.QueryResponse `json:"items"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type phase3ProofsResponse struct {
	SchemaVersion          string                                 `json:"schema_version"`
	CurrentState           string                                 `json:"current_state"`
	VulnerabilityArtifacts []vulnerabilitycore.RelevanceVerdict   `json:"vulnerability_artifacts,omitempty"`
	SupplyChainArtifacts   []supplychaincore.PatternVerdict       `json:"supply_chain_artifacts,omitempty"`
	StrategicArtifacts     []intelligencecore.StrategicAssessment `json:"strategic_artifacts,omitempty"`
	QueryArtifacts         []intelligencecore.QueryResponse       `json:"query_artifacts,omitempty"`
	Limitations            []string                               `json:"limitations,omitempty"`
}

func (s server) intelligenceVulnerabilityRelevanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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
		filter, err := parsePhase3IntelligenceFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase3VulnerabilityRelevance(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase3VulnerabilityListResponse{
			SchemaVersion: phase3VulnerabilityListSchema,
			CurrentState:  map[bool]string{true: phase3VulnerabilityStateActive, false: "vulnerability_relevance_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Vulnerability relevance remains advisory and evidence-backed; package presence alone does not imply active priority.",
			},
		})
	case http.MethodPost:
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
		filter, err := parsePhase3IntelligenceFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase3VulnerabilityEvaluateRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Input.SubjectRef = firstNonEmpty(normalizePhase3SubjectRef(request.Input.SubjectRef), filter.SubjectRef)
		request.Input.VulnerabilityID = strings.ToUpper(strings.TrimSpace(request.Input.VulnerabilityID))
		request.Input.PackageName = strings.TrimSpace(request.Input.PackageName)
		verdict := vulnerabilitycore.Evaluate(request.Input, time.Now)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase3Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeIntelligenceVulnerabilityRecorded, filter, &verdict, nil, nil, nil); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, phase3VulnerabilityEvaluateResponse{
			Status:  "recorded",
			Verdict: verdict,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) intelligenceSupplyChainPatternsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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
		filter, err := parsePhase3IntelligenceFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase3SupplyChainPatterns(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase3SupplyChainListResponse{
			SchemaVersion: phase3SupplyChainListSchema,
			CurrentState:  map[bool]string{true: phase3SupplyChainStateActive, false: "supply_chain_patterns_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Supply-chain pattern signals are bounded heuristics with reason codes and weighted federation inputs; they do not automatically mutate trust or policy.",
			},
		})
	case http.MethodPost:
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
		filter, err := parsePhase3IntelligenceFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase3SupplyChainEvaluateRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Input.SubjectRef = firstNonEmpty(normalizePhase3SubjectRef(request.Input.SubjectRef), filter.SubjectRef)
		request.Input.PackageName = firstNonEmpty(strings.TrimSpace(request.Input.PackageName), filter.PackageName)
		pattern := supplychaincore.EvaluatePattern(request.Input, time.Now)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase3Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeIntelligenceSupplyChainRecorded, filter, nil, &pattern, nil, nil); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, phase3SupplyChainEvaluateResponse{
			Status:  "recorded",
			Pattern: pattern,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) intelligenceStrategicSimulationHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parsePhase3IntelligenceFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	var request phase3StrategicSimulateRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter.SubjectRef = firstNonEmpty(normalizePhase3SubjectRef(request.SubjectRef), filter.SubjectRef)
	if strings.TrimSpace(filter.SubjectRef) == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": errPhase3SubjectRequired.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	input, err := s.resolvePhase3StrategicInput(ctx, filter, request)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, errPhase3SubjectRequired) {
			status = http.StatusBadRequest
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	assessment := intelligencecore.Assess(input, time.Now)
	if err := s.persistPhase3Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeIntelligenceStrategicRecorded, filter, nil, nil, &assessment, nil); err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, phase3StrategicSimulateResponse{
		Status:     "recorded",
		Assessment: assessment,
	})
}

func (s server) intelligenceStrategicQueryHandler(w http.ResponseWriter, r *http.Request) {
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
		filter, err := parsePhase3IntelligenceFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase3Queries(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase3StrategicQueryListResponse{
			SchemaVersion: phase3StrategicQuerySchema,
			CurrentState:  map[bool]string{true: phase3QueryStateActive, false: "grounded_query_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Natural-language security answers are retrieval-grounded and advisory-only; they do not approve or mutate policy, runtime, or VEX state.",
			},
		})
	case http.MethodPost:
		filter, err := parsePhase3IntelligenceFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase3StrategicQueryRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Query = strings.TrimSpace(request.Query)
		if request.Query == "" {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "query is required"})
			return
		}
		filter.SubjectRef = firstNonEmpty(normalizePhase3SubjectRef(request.SubjectRef), filter.SubjectRef)
		filter.VulnerabilityID = firstNonEmpty(strings.ToUpper(strings.TrimSpace(request.VulnerabilityID)), filter.VulnerabilityID)
		filter.PackageName = firstNonEmpty(strings.TrimSpace(request.PackageName), filter.PackageName)
		if strings.TrimSpace(filter.SubjectRef) == "" {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": errPhase3SubjectRequired.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		response, err := s.buildPhase3GroundedQuery(ctx, filter, request)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, errPhase3SubjectRequired) {
				status = http.StatusBadRequest
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		if err := s.persistPhase3Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeIntelligenceQueryRecorded, filter, nil, nil, nil, &response); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase3StrategicQueryResponse{
			Status:   "recorded",
			Response: response,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) intelligencePhase3ProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parsePhase3IntelligenceFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	vulnerabilities, err := s.listPhase3VulnerabilityRelevance(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	patterns, err := s.listPhase3SupplyChainPatterns(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	strategic, err := s.listPhase3StrategicAssessments(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	queries, err := s.listPhase3Queries(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	currentState := phase3ProofStateIncomplete
	if len(vulnerabilities) > 0 && hasPhase3VEXCandidate(vulnerabilities) && len(patterns) > 0 && len(strategic) > 0 && len(queries) > 0 {
		currentState = phase3ProofStateActive
	}
	httpjson.Write(w, http.StatusOK, phase3ProofsResponse{
		SchemaVersion:          phase3ProofsSchema,
		CurrentState:           currentState,
		VulnerabilityArtifacts: takePhase3VulnerabilityArtifacts(vulnerabilities, 5),
		SupplyChainArtifacts:   takePhase3SupplyChainArtifacts(patterns, 5),
		StrategicArtifacts:     takePhase3StrategicArtifacts(strategic, 5),
		QueryArtifacts:         takePhase3QueryArtifacts(queries, 5),
		Limitations: []string{
			"Phase 3 proofs expose bounded intelligence artifacts for relevance, supply-chain patterns, strategic assessments, and grounded advisory queries.",
		},
	})
}

func parsePhase3IntelligenceFilter(r *http.Request) (phase3IntelligenceFilter, error) {
	limit := runtimeLimit(r)
	if limit <= 0 {
		limit = 100
	}
	if limit > 500 {
		limit = 500
	}
	filter := phase3IntelligenceFilter{
		TenantID:        strings.TrimSpace(r.URL.Query().Get("tenant_id")),
		Environment:     strings.TrimSpace(r.URL.Query().Get("environment")),
		Repo:            strings.TrimSpace(r.URL.Query().Get("repo")),
		SubjectRef:      normalizePhase3SubjectRef(strings.TrimSpace(r.URL.Query().Get("subject_ref"))),
		VulnerabilityID: strings.ToUpper(strings.TrimSpace(firstNonEmpty(r.URL.Query().Get("vulnerability_id"), r.URL.Query().Get("cve_id")))),
		PackageName:     strings.TrimSpace(r.URL.Query().Get("package_name")),
		Limit:           limit,
	}
	if filter.SubjectRef != "" {
		clusterID, _, _, _, err := parseRuntimeSubjectRef(filter.SubjectRef)
		if err != nil {
			return filter, audit.ErrInvalidFilter
		}
		filter.ClusterID = clusterID
	}
	return filter, nil
}

func (s server) resolvePhase3StrategicInput(ctx context.Context, filter phase3IntelligenceFilter, request phase3StrategicSimulateRequest) (intelligencecore.StrategicAssessmentInput, error) {
	filter.SubjectRef = firstNonEmpty(normalizePhase3SubjectRef(request.SubjectRef), filter.SubjectRef)
	filter.VulnerabilityID = firstNonEmpty(strings.ToUpper(strings.TrimSpace(request.VulnerabilityID)), filter.VulnerabilityID)
	filter.PackageName = firstNonEmpty(strings.TrimSpace(request.PackageName), filter.PackageName)
	if strings.TrimSpace(filter.SubjectRef) == "" {
		return intelligencecore.StrategicAssessmentInput{}, errPhase3SubjectRequired
	}

	vulnerabilities, err := s.listPhase3VulnerabilityRelevance(ctx, filter)
	if err != nil {
		return intelligencecore.StrategicAssessmentInput{}, err
	}
	patterns, err := s.listPhase3SupplyChainPatterns(ctx, filter)
	if err != nil {
		return intelligencecore.StrategicAssessmentInput{}, err
	}

	subjectRef := filter.SubjectRef
	relevanceScore := 50
	patternTrustScore := 60
	blastRadius := request.BlastRadiusScore
	observed := append([]string{}, request.ObservedFacts...)
	inferred := append([]string{}, request.InferredConclusions...)
	evidenceRefs := append([]string{}, request.EvidenceRefs...)

	if len(vulnerabilities) > 0 {
		latest := vulnerabilities[0]
		subjectRef = firstNonEmpty(subjectRef, latest.SubjectRef)
		relevanceScore = latest.RelevanceScore
		observed = append(observed, latest.Explanation.Observed...)
		inferred = append(inferred, latest.Explanation.Derived...)
		evidenceRefs = append(evidenceRefs, latest.EvidenceRefs...)
		if latest.VEXCandidate != nil {
			inferred = append(inferred, "bounded VEX candidate status is "+latest.VEXCandidate.CandidateStatus)
			evidenceRefs = append(evidenceRefs, latest.VEXCandidate.EvidenceRefs...)
		}
		if blastRadius <= 0 {
			switch latest.CurrentState {
			case vulnerabilitycore.RelevanceActivePriority:
				blastRadius = 80
			case vulnerabilitycore.RelevanceReachableExternally:
				blastRadius = 70
			default:
				blastRadius = 50
			}
		}
	}
	if len(patterns) > 0 {
		latest := patterns[0]
		subjectRef = firstNonEmpty(subjectRef, latest.SubjectRef)
		patternTrustScore = latest.TrustScore
		observed = append(observed, latest.Explanation.Observed...)
		inferred = append(inferred, latest.Explanation.Derived...)
		evidenceRefs = append(evidenceRefs, latest.EvidenceRefs...)
		if blastRadius <= 0 && latest.CurrentState == supplychaincore.PatternStateCrossClusterConcern {
			blastRadius = 75
		}
	}
	if blastRadius <= 0 {
		blastRadius = 45
	}

	return intelligencecore.StrategicAssessmentInput{
		SubjectRef:          subjectRef,
		CandidateAction:     strings.TrimSpace(request.CandidateAction),
		RelevanceScore:      relevanceScore,
		PatternTrustScore:   patternTrustScore,
		BlastRadiusScore:    blastRadius,
		DelayDays:           request.DelayDays,
		EffortBand:          strings.TrimSpace(request.EffortBand),
		ObservedFacts:       uniqueStrings(observed),
		InferredConclusions: uniqueStrings(inferred),
		EvidenceRefs:        uniqueStrings(evidenceRefs),
	}, nil
}

func (s server) buildPhase3GroundedQuery(ctx context.Context, filter phase3IntelligenceFilter, request phase3StrategicQueryRequest) (intelligencecore.QueryResponse, error) {
	filter.SubjectRef = firstNonEmpty(normalizePhase3SubjectRef(request.SubjectRef), filter.SubjectRef)
	filter.VulnerabilityID = firstNonEmpty(strings.ToUpper(strings.TrimSpace(request.VulnerabilityID)), filter.VulnerabilityID)
	filter.PackageName = firstNonEmpty(strings.TrimSpace(request.PackageName), filter.PackageName)
	if strings.TrimSpace(filter.SubjectRef) == "" {
		return intelligencecore.QueryResponse{}, errPhase3SubjectRequired
	}

	vulnerabilities, err := s.listPhase3VulnerabilityRelevance(ctx, filter)
	if err != nil {
		return intelligencecore.QueryResponse{}, err
	}
	patterns, err := s.listPhase3SupplyChainPatterns(ctx, filter)
	if err != nil {
		return intelligencecore.QueryResponse{}, err
	}
	strategic, err := s.listPhase3StrategicAssessments(ctx, filter)
	if err != nil {
		return intelligencecore.QueryResponse{}, err
	}

	observed := []string{}
	inferred := []string{}
	recommended := []string{}
	uncertainty := []string{
		"Natural-language output is limited to currently persisted Phase 3 intelligence artifacts in the filtered scope.",
	}
	evidenceRefs := []string{}

	if len(vulnerabilities) > 0 {
		latest := vulnerabilities[0]
		observed = append(observed, latest.Explanation.Observed...)
		inferred = append(inferred, latest.Explanation.Derived...)
		recommended = append(recommended, latest.Explanation.Recommended...)
		evidenceRefs = append(evidenceRefs, latest.EvidenceRefs...)
		if latest.VEXCandidate != nil {
			inferred = append(inferred, "bounded VEX candidate status is "+latest.VEXCandidate.CandidateStatus)
			evidenceRefs = append(evidenceRefs, latest.VEXCandidate.EvidenceRefs...)
		}
	}
	if len(patterns) > 0 {
		latest := patterns[0]
		observed = append(observed, latest.Explanation.Observed...)
		inferred = append(inferred, latest.Explanation.Derived...)
		evidenceRefs = append(evidenceRefs, latest.EvidenceRefs...)
	}
	if len(strategic) > 0 {
		latest := strategic[0]
		observed = append(observed, latest.ObservedFacts...)
		inferred = append(inferred, latest.InferredConclusions...)
		recommended = append(recommended, latest.RecommendedActions...)
		evidenceRefs = append(evidenceRefs, latest.EvidenceRefs...)
		uncertainty = append(uncertainty, latest.Uncertainty...)
	}
	if len(observed) == 0 && len(inferred) == 0 && len(recommended) == 0 {
		observed = append(observed, "no matching intelligence artifact was found in the current bounded evidence set")
	}

	return intelligencecore.BuildGroundedQuery(
		request.Query,
		intelligencecore.QueryScope{
			SubjectRef:      filter.SubjectRef,
			VulnerabilityID: filter.VulnerabilityID,
			PackageName:     filter.PackageName,
			TenantID:        filter.TenantID,
			Environment:     filter.Environment,
			Repo:            filter.Repo,
		},
		uniqueStrings(observed),
		uniqueStrings(inferred),
		uniqueStrings(recommended),
		uniqueStrings(uncertainty),
		uniqueStrings(evidenceRefs),
		time.Now,
	), nil
}

func (s server) persistPhase3Event(ctx context.Context, requestID, actor, eventType string, filter phase3IntelligenceFilter, vulnerability *vulnerabilitycore.RelevanceVerdict, pattern *supplychaincore.PatternVerdict, strategic *intelligencecore.StrategicAssessment, query *intelligencecore.QueryResponse) error {
	payload, err := canonicalJSON(phase3IntelligencePayload{
		SchemaVersion: phase3IntelligencePayloadSchema,
		Vulnerability: vulnerability,
		SupplyChain:   pattern,
		Strategic:     strategic,
		Query:         query,
	})
	if err != nil {
		return err
	}

	subjectRef := ""
	digest := ""
	reasons := []string{"phase3_advisory_intelligence_recorded"}
	if vulnerability != nil {
		subjectRef = firstNonEmpty(subjectRef, vulnerability.SubjectRef)
		digest = firstNonEmpty(digest, vulnerability.ImageDigest)
		reasons = append(reasons, vulnerability.CurrentState, vulnerability.VulnerabilityID)
		if vulnerability.VEXCandidate != nil {
			reasons = append(reasons, vulnerability.VEXCandidate.CandidateStatus)
		}
	}
	if pattern != nil {
		subjectRef = firstNonEmpty(subjectRef, pattern.SubjectRef)
		reasons = append(reasons, pattern.CurrentState)
		reasons = append(reasons, pattern.ReasonCodes...)
	}
	if strategic != nil {
		subjectRef = firstNonEmpty(subjectRef, strategic.SubjectRef)
		reasons = append(reasons, strategic.CurrentState)
		reasons = append(reasons, strategic.RecommendedActions...)
	}
	if query != nil {
		subjectRef = firstNonEmpty(subjectRef, query.Scope.SubjectRef)
		reasons = append(reasons, query.CurrentState, query.RetrievalMode)
		reasons = append(reasons, firstNonEmpty(query.Scope.VulnerabilityID, query.Scope.PackageName))
	}

	clusterID := filter.ClusterID
	namespace := ""
	workloadKind := ""
	workload := ""
	if parsedCluster, parsedNamespace, parsedWorkloadKind, parsedWorkload, err := parseRuntimeSubjectRef(subjectRef); err == nil {
		clusterID = firstNonEmpty(clusterID, parsedCluster)
		namespace = parsedNamespace
		workloadKind = parsedWorkloadKind
		workload = parsedWorkload
	}

	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:    firstNonEmpty(strings.TrimSpace(requestID), audit.NewRequestID()),
		Component:    phase3IntelligenceComponent,
		EventType:    eventType,
		Actor:        strings.TrimSpace(actor),
		ClusterID:    clusterID,
		TenantID:     firstNonEmpty(filter.TenantID, audit.TenantFromNamespace(namespace)),
		Environment:  firstNonEmpty(filter.Environment, audit.EnvironmentFromNamespace(namespace)),
		Namespace:    namespace,
		WorkloadKind: workloadKind,
		Workload:     workload,
		Repo:         filter.Repo,
		Digest:       digest,
		Decision:     audit.DecisionAllow,
		Reasons:      uniqueStrings(reasons),
		Intelligence: payload,
	})
	return err
}

func (s server) listPhase3VulnerabilityRelevance(ctx context.Context, filter phase3IntelligenceFilter) ([]vulnerabilitycore.RelevanceVerdict, error) {
	events, err := s.listPhase3Events(ctx, filter, audit.EventTypeIntelligenceVulnerabilityRecorded)
	if err != nil {
		return nil, err
	}
	items := []vulnerabilitycore.RelevanceVerdict{}
	for _, item := range events {
		payload := parsePhase3IntelligencePayload(item.Intelligence)
		if payload.Vulnerability == nil {
			continue
		}
		if !matchesPhase3Subject(filter.SubjectRef, payload.Vulnerability.SubjectRef) {
			continue
		}
		if filter.VulnerabilityID != "" && payload.Vulnerability.VulnerabilityID != filter.VulnerabilityID {
			continue
		}
		if filter.PackageName != "" && !strings.EqualFold(payload.Vulnerability.PackageName, filter.PackageName) {
			continue
		}
		items = append(items, *payload.Vulnerability)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].EvaluatedAt.After(items[j].EvaluatedAt) })
	return limitPhase3Items(items, filter.Limit), nil
}

func (s server) listPhase3SupplyChainPatterns(ctx context.Context, filter phase3IntelligenceFilter) ([]supplychaincore.PatternVerdict, error) {
	events, err := s.listPhase3Events(ctx, filter, audit.EventTypeIntelligenceSupplyChainRecorded)
	if err != nil {
		return nil, err
	}
	items := []supplychaincore.PatternVerdict{}
	for _, item := range events {
		payload := parsePhase3IntelligencePayload(item.Intelligence)
		if payload.SupplyChain == nil {
			continue
		}
		if !matchesPhase3Subject(filter.SubjectRef, payload.SupplyChain.SubjectRef) {
			continue
		}
		if filter.PackageName != "" && !strings.EqualFold(payload.SupplyChain.PackageName, filter.PackageName) {
			continue
		}
		items = append(items, *payload.SupplyChain)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ObservedAt.After(items[j].ObservedAt) })
	return limitPhase3Items(items, filter.Limit), nil
}

func (s server) listPhase3StrategicAssessments(ctx context.Context, filter phase3IntelligenceFilter) ([]intelligencecore.StrategicAssessment, error) {
	events, err := s.listPhase3Events(ctx, filter, audit.EventTypeIntelligenceStrategicRecorded)
	if err != nil {
		return nil, err
	}
	items := []intelligencecore.StrategicAssessment{}
	for _, item := range events {
		payload := parsePhase3IntelligencePayload(item.Intelligence)
		if payload.Strategic == nil {
			continue
		}
		if !matchesPhase3Subject(filter.SubjectRef, payload.Strategic.SubjectRef) {
			continue
		}
		items = append(items, *payload.Strategic)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].EvaluatedAt.After(items[j].EvaluatedAt) })
	return limitPhase3Items(items, filter.Limit), nil
}

func (s server) listPhase3Queries(ctx context.Context, filter phase3IntelligenceFilter) ([]intelligencecore.QueryResponse, error) {
	events, err := s.listPhase3Events(ctx, filter, audit.EventTypeIntelligenceQueryRecorded)
	if err != nil {
		return nil, err
	}
	items := []intelligencecore.QueryResponse{}
	for _, item := range events {
		payload := parsePhase3IntelligencePayload(item.Intelligence)
		if payload.Query == nil {
			continue
		}
		if !matchesPhase3QueryScope(filter, payload.Query.Scope) {
			continue
		}
		items = append(items, *payload.Query)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].AnsweredAt.After(items[j].AnsweredAt) })
	return limitPhase3Items(items, filter.Limit), nil
}

func (s server) listPhase3Events(ctx context.Context, filter phase3IntelligenceFilter, eventType string) ([]audit.StoredEvent, error) {
	return s.store.ListEvents(ctx, audit.EventFilter{
		EventType:   eventType,
		Component:   phase3IntelligenceComponent,
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       max(filter.Limit, 200),
	})
}

func parsePhase3IntelligencePayload(value json.RawMessage) phase3IntelligencePayload {
	if len(value) == 0 || string(value) == "null" {
		return phase3IntelligencePayload{}
	}
	var payload phase3IntelligencePayload
	if err := json.Unmarshal(value, &payload); err != nil {
		return phase3IntelligencePayload{}
	}
	return payload
}

func hasPhase3VEXCandidate(items []vulnerabilitycore.RelevanceVerdict) bool {
	for _, item := range items {
		if item.VEXCandidate != nil && strings.TrimSpace(item.VEXCandidate.CandidateStatus) != "" {
			return true
		}
	}
	return false
}

func takePhase3VulnerabilityArtifacts(items []vulnerabilitycore.RelevanceVerdict, limit int) []vulnerabilitycore.RelevanceVerdict {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase3SupplyChainArtifacts(items []supplychaincore.PatternVerdict, limit int) []supplychaincore.PatternVerdict {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase3StrategicArtifacts(items []intelligencecore.StrategicAssessment, limit int) []intelligencecore.StrategicAssessment {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase3QueryArtifacts(items []intelligencecore.QueryResponse, limit int) []intelligencecore.QueryResponse {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func matchesPhase3Subject(filterSubject, itemSubject string) bool {
	filterSubject = normalizePhase3SubjectRef(filterSubject)
	itemSubject = normalizePhase3SubjectRef(itemSubject)
	return filterSubject == "" || filterSubject == itemSubject
}

func matchesPhase3QueryScope(filter phase3IntelligenceFilter, scope intelligencecore.QueryScope) bool {
	if !matchesPhase3Subject(filter.SubjectRef, scope.SubjectRef) {
		return false
	}
	if filter.VulnerabilityID != "" && !strings.EqualFold(strings.TrimSpace(scope.VulnerabilityID), filter.VulnerabilityID) {
		return false
	}
	if filter.PackageName != "" && !strings.EqualFold(strings.TrimSpace(scope.PackageName), filter.PackageName) {
		return false
	}
	if filter.TenantID != "" && strings.TrimSpace(scope.TenantID) != "" && !strings.EqualFold(strings.TrimSpace(scope.TenantID), filter.TenantID) {
		return false
	}
	if filter.Environment != "" && strings.TrimSpace(scope.Environment) != "" && !strings.EqualFold(strings.TrimSpace(scope.Environment), filter.Environment) {
		return false
	}
	if filter.Repo != "" && strings.TrimSpace(scope.Repo) != "" && !strings.EqualFold(strings.TrimSpace(scope.Repo), filter.Repo) {
		return false
	}
	return true
}

func limitPhase3Items[T any](items []T, limit int) []T {
	if limit > 0 && len(items) > limit {
		return items[:limit]
	}
	return items
}

func normalizePhase3SubjectRef(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || strings.Contains(value, "|") {
		return value
	}
	parts := strings.Split(value, "/")
	if len(parts) != 4 {
		return value
	}
	return runtimeSubjectRef(parts[0], parts[1], parts[2], parts[3])
}
