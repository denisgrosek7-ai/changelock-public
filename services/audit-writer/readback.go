package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	readbackEnvelopeSchemaVersion   = "9b.v1"
	readbackProjectionSchemaVersion = "9b.readback.v1"

	readbackResourceDefenseGap   = "defense-gap"
	readbackResourcePolicyReplay = "policy-replay"
	readbackResourceSystemicWeak = "systemic-weakness"
	readbackProjectionInternal   = incidentAudienceInternal
)

type advisoryReadbackRef struct {
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id"`
	ResourceURI  string `json:"resource_uri"`
	EvidenceHash string `json:"evidence_hash"`
}

type readbackScope struct {
	TenantID     string `json:"tenant_id,omitempty"`
	ClusterID    string `json:"cluster_id,omitempty"`
	Environment  string `json:"environment,omitempty"`
	Repository   string `json:"repository,omitempty"`
	State        string `json:"state,omitempty"`
	Severity     string `json:"severity,omitempty"`
	ScorecardRef string `json:"scorecard_ref,omitempty"`
}

type readbackDescriptor struct {
	ResourceType string        `json:"resource_type"`
	SubjectType  string        `json:"subject_type"`
	SubjectRef   string        `json:"subject_ref"`
	Scope        readbackScope `json:"scope"`
}

type decisionVerdictContext struct {
	Summary         string   `json:"summary,omitempty"`
	CurrentOutcome  string   `json:"current_outcome,omitempty"`
	ProposedOutcome string   `json:"proposed_outcome,omitempty"`
	Delta           string   `json:"delta,omitempty"`
	PatternKey      string   `json:"pattern_key,omitempty"`
	GapTypes        []string `json:"gap_types,omitempty"`
}

type decisionSnapshotRefs struct {
	PolicySnapshotRef   string   `json:"policy_snapshot_ref"`
	EvaluatorInputHash  string   `json:"evaluator_input_hash"`
	EvaluatorOutputHash string   `json:"evaluator_output_hash"`
	EvidenceRefs        []string `json:"evidence_refs"`
}

type decisionEvidenceEnvelope struct {
	SchemaVersion           string                 `json:"schema_version"`
	ResourceType            string                 `json:"resource_type"`
	ResourceID              string                 `json:"resource_id"`
	EvidenceHash            string                 `json:"evidence_hash"`
	GeneratedAt             time.Time              `json:"generated_at"`
	SubjectType             string                 `json:"subject_type"`
	SubjectRef              string                 `json:"subject_ref"`
	VerdictContext          decisionVerdictContext `json:"verdict_context"`
	SnapshotRefs            decisionSnapshotRefs   `json:"snapshot_refs"`
	RedactionProfileVersion string                 `json:"redaction_profile_version"`
	ProjectionSchemaVersion string                 `json:"projection_schema_version"`
	AdvisoryOnly            bool                   `json:"advisory_only"`
	Limitations             []string               `json:"limitations"`
}

type readbackResponse struct {
	ResourceType       string                   `json:"resource_type"`
	ResourceID         string                   `json:"resource_id"`
	PermanentURI       string                   `json:"permanent_uri"`
	ProjectionAudience string                   `json:"projection_audience"`
	AdvisoryOnly       bool                     `json:"advisory_only"`
	PayloadSummary     string                   `json:"payload_summary"`
	EvidenceEnvelope   decisionEvidenceEnvelope `json:"evidence_envelope"`
	Payload            any                      `json:"payload"`
	TopologyContext    *readbackTopologyContext `json:"topology_context,omitempty"`
	Limitations        []string                 `json:"limitations"`
}

type readbackTopologyNodeSummary struct {
	NodeID           string `json:"node_id,omitempty"`
	Service          string `json:"service,omitempty"`
	Namespace        string `json:"namespace,omitempty"`
	Cluster          string `json:"cluster,omitempty"`
	Environment      string `json:"environment,omitempty"`
	PublicExposure   bool   `json:"public_exposure"`
	SensitivityClass string `json:"sensitivity_class,omitempty"`
}

type readbackTopologyContext struct {
	AdvisoryOnly         bool                         `json:"advisory_only"`
	SubjectType          string                       `json:"subject_type"`
	SubjectRef           string                       `json:"subject_ref"`
	PrimaryAffectedNode  *readbackTopologyNodeSummary `json:"primary_affected_node,omitempty"`
	BlastRadiusScore     int                          `json:"blast_radius_score"`
	CriticalReachCount   int                          `json:"critical_reach_count"`
	TopRiskPathSummaries []string                     `json:"top_risk_path_summaries,omitempty"`
	Limitations          []string                     `json:"limitations,omitempty"`
}

type readbackGrantRequest struct {
	ResourceType     string `json:"resource_type"`
	ResourceID       string `json:"resource_id"`
	Audience         string `json:"audience"`
	ExpiresInMinutes int    `json:"expires_in_minutes"`
	Purpose          string `json:"purpose,omitempty"`
}

type readbackGrantResponse struct {
	GrantID      string    `json:"grant_id"`
	ShareURL     string    `json:"share_url"`
	ResourceType string    `json:"resource_type"`
	ResourceID   string    `json:"resource_id"`
	Audience     string    `json:"audience"`
	ExpiresAt    time.Time `json:"expires_at"`
	Purpose      string    `json:"purpose,omitempty"`
}

type signedReadbackGrant struct {
	Version              string    `json:"version"`
	ResourceType         string    `json:"resource_type"`
	ResourceID           string    `json:"resource_id"`
	Audience             string    `json:"audience"`
	ExpectedEvidenceHash string    `json:"expected_evidence_hash"`
	ExpiresAt            time.Time `json:"expires_at"`
	Purpose              string    `json:"purpose,omitempty"`
}

func readbackScopeFromFilter(filter incidentFilter) readbackScope {
	return readbackScope{
		TenantID:     strings.TrimSpace(filter.event.TenantID),
		ClusterID:    strings.TrimSpace(filter.event.ClusterID),
		Environment:  strings.TrimSpace(filter.event.Environment),
		Repository:   strings.TrimSpace(filter.event.Repo),
		State:        strings.TrimSpace(filter.State),
		Severity:     strings.TrimSpace(filter.Severity),
		ScorecardRef: strings.TrimSpace(filter.ScorecardRef),
	}
}

func (scope readbackScope) toIncidentFilter() incidentFilter {
	return incidentFilter{
		event: audit.EventFilter{
			TenantID:    scope.TenantID,
			ClusterID:   scope.ClusterID,
			Environment: scope.Environment,
			Repo:        scope.Repository,
		},
		State:        scope.State,
		Severity:     scope.Severity,
		ScorecardRef: scope.ScorecardRef,
	}
}

func canonicalBytes(value any) []byte {
	data, err := json.Marshal(value)
	if err != nil {
		return []byte("{}")
	}
	return data
}

func sha256Prefixed(value any) string {
	sum := sha256.Sum256(canonicalBytes(value))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func encodeReadbackDescriptor(descriptor readbackDescriptor) string {
	return base64.RawURLEncoding.EncodeToString(canonicalBytes(descriptor))
}

func decodeReadbackDescriptor(encoded string) (readbackDescriptor, error) {
	data, err := base64.RawURLEncoding.DecodeString(strings.TrimSpace(encoded))
	if err != nil {
		return readbackDescriptor{}, errors.New("invalid readback resource id")
	}
	var descriptor readbackDescriptor
	if err := json.Unmarshal(data, &descriptor); err != nil {
		return readbackDescriptor{}, errors.New("invalid readback resource descriptor")
	}
	if descriptor.ResourceType == "" || descriptor.SubjectType == "" || descriptor.SubjectRef == "" {
		return readbackDescriptor{}, errors.New("invalid readback resource descriptor")
	}
	return descriptor, nil
}

func buildAdvisoryReadbackRef(resourceType string, descriptor readbackDescriptor, envelope decisionEvidenceEnvelope) advisoryReadbackRef {
	resourceID := encodeReadbackDescriptor(descriptor)
	return advisoryReadbackRef{
		ResourceType: resourceType,
		ResourceID:   resourceID,
		ResourceURI:  fmt.Sprintf("/r/%s/%s", resourceType, resourceID),
		EvidenceHash: envelope.EvidenceHash,
	}
}

func attachDefenseGapReadback(assessment defenseGapAssessment, filter incidentFilter) defenseGapAssessment {
	descriptor := readbackDescriptor{
		ResourceType: readbackResourceDefenseGap,
		SubjectType:  assessment.SubjectType,
		SubjectRef:   assessment.SubjectRef,
		Scope:        readbackScopeFromFilter(filter),
	}
	envelope := buildDefenseGapEnvelope(assessment, descriptor)
	assessment.Readback = buildAdvisoryReadbackRef(readbackResourceDefenseGap, descriptor, envelope)
	return assessment
}

func attachPolicyReplayReadback(assessment policyReplayAssessment, filter incidentFilter) policyReplayAssessment {
	descriptor := readbackDescriptor{
		ResourceType: readbackResourcePolicyReplay,
		SubjectType:  assessment.SubjectType,
		SubjectRef:   assessment.SubjectRef,
		Scope:        readbackScopeFromFilter(filter),
	}
	envelope := buildPolicyReplayEnvelope(assessment, descriptor)
	assessment.Readback = buildAdvisoryReadbackRef(readbackResourcePolicyReplay, descriptor, envelope)
	return assessment
}

func attachSystemicWeaknessReadback(response systemicWeaknessResponse, filter incidentFilter) systemicWeaknessResponse {
	for index, weakness := range response.Weaknesses {
		descriptor := readbackDescriptor{
			ResourceType: readbackResourceSystemicWeak,
			SubjectType:  "cluster",
			SubjectRef:   weakness.PatternKey,
			Scope:        readbackScopeFromFilter(filter),
		}
		envelope := buildSystemicWeaknessEnvelope(weakness, descriptor, response.GeneratedAt)
		response.Weaknesses[index].Readback = buildAdvisoryReadbackRef(readbackResourceSystemicWeak, descriptor, envelope)
	}
	return response
}

func buildDefenseGapEnvelope(assessment defenseGapAssessment, descriptor readbackDescriptor) decisionEvidenceEnvelope {
	gapTypes := make([]string, 0, len(assessment.DefenseGaps))
	evidenceRefs := make([]string, 0, len(assessment.DefenseGaps)*4)
	for _, finding := range assessment.DefenseGaps {
		gapTypes = append(gapTypes, finding.GapType)
		evidenceRefs = append(evidenceRefs, finding.EvidenceRefs...)
	}
	envelope := decisionEvidenceEnvelope{
		SchemaVersion: readbackEnvelopeSchemaVersion,
		ResourceType:  descriptor.ResourceType,
		ResourceID:    encodeReadbackDescriptor(descriptor),
		GeneratedAt:   assessment.GeneratedAt,
		SubjectType:   descriptor.SubjectType,
		SubjectRef:    descriptor.SubjectRef,
		VerdictContext: decisionVerdictContext{
			Summary:    assessment.SystemicPattern.Summary,
			PatternKey: assessment.SystemicPattern.PatternKey,
			GapTypes:   uniqueStrings(gapTypes),
		},
		SnapshotRefs: decisionSnapshotRefs{
			PolicySnapshotRef: readbackPolicySnapshotRef(descriptor),
			EvaluatorInputHash: sha256Prefixed(struct {
				Descriptor   readbackDescriptor `json:"descriptor"`
				EvidenceRefs []string           `json:"evidence_refs"`
			}{Descriptor: descriptor, EvidenceRefs: uniqueStrings(evidenceRefs)}),
			EvaluatorOutputHash: sha256Prefixed(struct {
				SubjectRef  string              `json:"subject_ref"`
				DefenseGaps []defenseGapFinding `json:"defense_gaps"`
				Pattern     defenseGapPattern   `json:"systemic_pattern"`
			}{SubjectRef: assessment.SubjectRef, DefenseGaps: assessment.DefenseGaps, Pattern: assessment.SystemicPattern}),
			EvidenceRefs: limitStrings(uniqueStrings(evidenceRefs), 12),
		},
		RedactionProfileVersion: "8q.v2",
		ProjectionSchemaVersion: readbackProjectionSchemaVersion,
		AdvisoryOnly:            true,
		Limitations: append([]string{
			"Readback is reconstructed from canonical advisory inputs already present in ChangeLock and does not create a new truth model.",
		}, assessment.Limitations...),
	}
	envelope.EvidenceHash = computeEnvelopeHash(envelope)
	return envelope
}

func buildPolicyReplayEnvelope(assessment policyReplayAssessment, descriptor readbackDescriptor) decisionEvidenceEnvelope {
	evidenceRefs := make([]string, 0, len(assessment.ReplayResults)*4)
	currentOutcome := "No replay result available."
	proposedOutcome := ""
	delta := ""
	if len(assessment.ReplayResults) > 0 {
		currentOutcome = assessment.ReplayResults[0].CurrentOutcome
		proposedOutcome = assessment.ReplayResults[0].ProposedOutcome
		delta = assessment.ReplayResults[0].Delta
	}
	for _, result := range assessment.ReplayResults {
		evidenceRefs = append(evidenceRefs, result.SupportingEvidenceRefs...)
	}
	for _, gap := range assessment.CoverageGaps {
		evidenceRefs = append(evidenceRefs, gap.EvidenceRefs...)
	}
	envelope := decisionEvidenceEnvelope{
		SchemaVersion: readbackEnvelopeSchemaVersion,
		ResourceType:  descriptor.ResourceType,
		ResourceID:    encodeReadbackDescriptor(descriptor),
		GeneratedAt:   assessment.GeneratedAt,
		SubjectType:   descriptor.SubjectType,
		SubjectRef:    descriptor.SubjectRef,
		VerdictContext: decisionVerdictContext{
			Summary:         fmt.Sprintf("%d replay result(s) and %d coverage gap(s) are attached to this advisory readback.", len(assessment.ReplayResults), len(assessment.CoverageGaps)),
			CurrentOutcome:  currentOutcome,
			ProposedOutcome: proposedOutcome,
			Delta:           delta,
		},
		SnapshotRefs: decisionSnapshotRefs{
			PolicySnapshotRef: readbackPolicySnapshotRef(descriptor),
			EvaluatorInputHash: sha256Prefixed(struct {
				Descriptor   readbackDescriptor `json:"descriptor"`
				EvidenceRefs []string           `json:"evidence_refs"`
				Subject      string             `json:"subject"`
			}{Descriptor: descriptor, EvidenceRefs: uniqueStrings(evidenceRefs), Subject: assessment.SubjectRef}),
			EvaluatorOutputHash: sha256Prefixed(struct {
				SubjectRef    string               `json:"subject_ref"`
				ReplayResults []policyReplayResult `json:"replay_results"`
				CoverageGaps  []coverageGapFinding `json:"coverage_gaps"`
				BlastRadius   replayBlastRadius    `json:"blast_radius"`
			}{SubjectRef: assessment.SubjectRef, ReplayResults: assessment.ReplayResults, CoverageGaps: assessment.CoverageGaps, BlastRadius: assessment.BlastRadius}),
			EvidenceRefs: limitStrings(uniqueStrings(evidenceRefs), 12),
		},
		RedactionProfileVersion: "8q.v2",
		ProjectionSchemaVersion: readbackProjectionSchemaVersion,
		AdvisoryOnly:            true,
		Limitations: append([]string{
			"Readback captures historical replay context from canonical incidents and evidence refs already present in the system; it is not an enforcement mutation path.",
		}, assessment.Limitations...),
	}
	envelope.EvidenceHash = computeEnvelopeHash(envelope)
	return envelope
}

func buildSystemicWeaknessEnvelope(weakness systemicWeakness, descriptor readbackDescriptor, generatedAt time.Time) decisionEvidenceEnvelope {
	envelope := decisionEvidenceEnvelope{
		SchemaVersion: readbackEnvelopeSchemaVersion,
		ResourceType:  descriptor.ResourceType,
		ResourceID:    encodeReadbackDescriptor(descriptor),
		GeneratedAt:   generatedAt.UTC(),
		SubjectType:   descriptor.SubjectType,
		SubjectRef:    descriptor.SubjectRef,
		VerdictContext: decisionVerdictContext{
			Summary:    weakness.Summary,
			PatternKey: weakness.PatternKey,
		},
		SnapshotRefs: decisionSnapshotRefs{
			PolicySnapshotRef: readbackPolicySnapshotRef(descriptor),
			EvaluatorInputHash: sha256Prefixed(struct {
				Descriptor          readbackDescriptor `json:"descriptor"`
				RelatedIncidentRefs []string           `json:"related_incident_refs"`
			}{Descriptor: descriptor, RelatedIncidentRefs: weakness.RelatedIncidentRefs}),
			EvaluatorOutputHash: sha256Prefixed(struct {
				PatternKey          string   `json:"pattern_key"`
				Summary             string   `json:"summary"`
				RelatedIncidentRefs []string `json:"related_incident_refs"`
				EvidenceRefs        []string `json:"evidence_refs"`
			}{PatternKey: weakness.PatternKey, Summary: weakness.Summary, RelatedIncidentRefs: weakness.RelatedIncidentRefs, EvidenceRefs: weakness.EvidenceRefs}),
			EvidenceRefs: limitStrings(uniqueStrings(append([]string{}, weakness.EvidenceRefs...)), 12),
		},
		RedactionProfileVersion: "8q.v2",
		ProjectionSchemaVersion: readbackProjectionSchemaVersion,
		AdvisoryOnly:            true,
		Limitations: append([]string{
			"Systemic weakness readback remains an evidence-backed root-cause hypothesis, not canonical incident truth.",
		}, weakness.Limitations...),
	}
	envelope.EvidenceHash = computeEnvelopeHash(envelope)
	return envelope
}

func computeEnvelopeHash(envelope decisionEvidenceEnvelope) string {
	copyEnvelope := envelope
	copyEnvelope.EvidenceHash = ""
	return sha256Prefixed(copyEnvelope)
}

func readbackPolicySnapshotRef(descriptor readbackDescriptor) string {
	parts := []string{"policy-snapshot", descriptor.ResourceType, descriptor.SubjectType, descriptor.SubjectRef}
	if descriptor.Scope.ScorecardRef != "" {
		parts = append(parts, descriptor.Scope.ScorecardRef)
	}
	if descriptor.Scope.Environment != "" {
		parts = append(parts, descriptor.Scope.Environment)
	}
	return strings.Join(parts, ":")
}

func readbackGrantSecret() string {
	if secret := strings.TrimSpace(os.Getenv("CHANGELOCK_READBACK_GRANT_SECRET")); secret != "" {
		return secret
	}
	return strings.TrimSpace(os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN"))
}

func (s server) readbackGrantSecret() string {
	if secret := strings.TrimSpace(s.readbackGrantSecretValue); secret != "" {
		return secret
	}
	if secret := strings.TrimSpace(s.internalToken); secret != "" {
		return secret
	}
	return readbackGrantSecret()
}

func signReadbackGrant(grant signedReadbackGrant, secret string) (string, error) {
	payloadBytes := canonicalBytes(grant)
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(payload))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return payload + "." + signature, nil
}

func parseAndVerifyReadbackGrant(token string, secret string) (signedReadbackGrant, error) {
	parts := strings.Split(strings.TrimSpace(token), ".")
	if len(parts) != 2 {
		return signedReadbackGrant{}, errors.New("invalid readback grant")
	}
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(parts[0]))
	expected := mac.Sum(nil)
	actual, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil || !hmac.Equal(actual, expected) {
		return signedReadbackGrant{}, errors.New("invalid readback grant")
	}
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return signedReadbackGrant{}, errors.New("invalid readback grant")
	}
	var grant signedReadbackGrant
	if err := json.Unmarshal(payloadBytes, &grant); err != nil {
		return signedReadbackGrant{}, errors.New("invalid readback grant")
	}
	return grant, nil
}

func grantID(token string) string {
	sum := sha256.Sum256([]byte(token))
	return "GRANT-" + strings.ToUpper(hex.EncodeToString(sum[:]))[:10]
}

func (s server) readbackGrantHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	var req readbackGrantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	audience, err := parseIncidentExportAudience(req.Audience)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if req.ExpiresInMinutes <= 0 {
		req.ExpiresInMinutes = 60
	}
	if req.ExpiresInMinutes > 24*60 {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "expires_in_minutes must be <= 1440"})
		return
	}
	descriptor, err := decodeReadbackDescriptor(req.ResourceID)
	if err != nil {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	if descriptor.ResourceType != req.ResourceType {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "readback resource not found"})
		return
	}
	if err := ensurePrincipalReadbackDescriptorScope(principal, descriptor); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	readback, err := s.materializeReadback(ctx, req.ResourceType, req.ResourceID, audience)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	secret := s.readbackGrantSecret()
	if secret == "" {
		httpjson.Write(w, http.StatusServiceUnavailable, map[string]string{"error": "readback share grants are not configured"})
		return
	}
	grant := signedReadbackGrant{
		Version:              readbackEnvelopeSchemaVersion,
		ResourceType:         req.ResourceType,
		ResourceID:           req.ResourceID,
		Audience:             audience,
		ExpectedEvidenceHash: readback.EvidenceEnvelope.EvidenceHash,
		ExpiresAt:            time.Now().UTC().Add(time.Duration(req.ExpiresInMinutes) * time.Minute),
		Purpose:              strings.TrimSpace(req.Purpose),
	}
	token, err := signReadbackGrant(grant, secret)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": "failed to sign readback grant"})
		return
	}
	httpjson.Write(w, http.StatusCreated, readbackGrantResponse{
		GrantID:      grantID(token),
		ShareURL:     "/s/" + token,
		ResourceType: req.ResourceType,
		ResourceID:   req.ResourceID,
		Audience:     audience,
		ExpiresAt:    grant.ExpiresAt,
		Purpose:      grant.Purpose,
	})
}

func (s server) readbackShareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/s/"))
	if token == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "share link not found"})
		return
	}
	secret := s.readbackGrantSecret()
	if secret == "" {
		httpjson.Write(w, http.StatusServiceUnavailable, map[string]string{"error": "readback share grants are not configured"})
		return
	}
	grant, err := parseAndVerifyReadbackGrant(token, secret)
	if err != nil {
		httpjson.Write(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}
	if time.Now().UTC().After(grant.ExpiresAt) {
		httpjson.Write(w, http.StatusUnauthorized, map[string]string{"error": "readback share grant expired"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	readback, err := s.materializeReadback(ctx, grant.ResourceType, grant.ResourceID, grant.Audience)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(readback.EvidenceEnvelope.EvidenceHash), []byte(grant.ExpectedEvidenceHash)) != 1 {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": "readback evidence envelope no longer matches the signed share grant"})
		return
	}
	httpjson.Write(w, http.StatusOK, readback)
}

func (s server) readbackDefenseGapHandler(w http.ResponseWriter, r *http.Request) {
	s.readbackHandler(w, r, readbackResourceDefenseGap)
}

func (s server) readbackPolicyReplayHandler(w http.ResponseWriter, r *http.Request) {
	s.readbackHandler(w, r, readbackResourcePolicyReplay)
}

func (s server) readbackSystemicWeaknessHandler(w http.ResponseWriter, r *http.Request) {
	s.readbackHandler(w, r, readbackResourceSystemicWeak)
}

func (s server) readbackHandler(w http.ResponseWriter, r *http.Request, resourceType string) {
	if strings.HasSuffix(strings.TrimSpace(r.URL.Path), "/forensic-context") {
		s.readbackForensicContextHandler(w, r, resourceType)
		return
	}
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	resourceID := readbackResourceIDFromPath(r.URL.Path, resourceType)
	if resourceID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "readback resource not found"})
		return
	}
	descriptor, err := decodeReadbackDescriptor(resourceID)
	if err != nil {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	if descriptor.ResourceType != resourceType {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "readback resource not found"})
		return
	}
	if err := ensurePrincipalReadbackDescriptorScope(principal, descriptor); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	readback, err := s.materializeReadback(ctx, resourceType, resourceID, readbackProjectionInternal)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	if err := ensurePrincipalReadbackScope(principal, readback.EvidenceEnvelope); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, readback)
}

func ensurePrincipalReadbackDescriptorScope(principal auth.Principal, descriptor readbackDescriptor) error {
	if !tenantScoped(principal) {
		return nil
	}
	return ensureTenantMatch(principal, descriptor.Scope.TenantID)
}

func ensurePrincipalReadbackScope(principal auth.Principal, envelope decisionEvidenceEnvelope) error {
	_ = principal
	_ = envelope
	// Descriptor-based tenant scope validation happens before materialization.
	return nil
}

func readbackResourceIDFromPath(path string, resourceType string) string {
	prefixes := []string{
		"/v1/readback/" + resourceType + "/",
		"/r/" + resourceType + "/",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(path, prefix) {
			return strings.TrimSpace(strings.TrimPrefix(path, prefix))
		}
	}
	return ""
}

func readbackNotFound(message string) error {
	message = strings.TrimSpace(message)
	if message == "" {
		message = "readback resource not found"
	}
	return fmt.Errorf("%w: %s", errIncidentNotFound, message)
}

func (s server) materializeReadback(ctx context.Context, resourceType string, resourceID string, audience string) (readbackResponse, error) {
	descriptor, err := decodeReadbackDescriptor(resourceID)
	if err != nil {
		return readbackResponse{}, readbackNotFound(err.Error())
	}
	if descriptor.ResourceType != resourceType {
		return readbackResponse{}, readbackNotFound("")
	}
	filter := descriptor.Scope.toIncidentFilter()
	topologyContext := s.materializeReadbackTopologyContext(ctx, descriptor, filter, audience)
	switch resourceType {
	case readbackResourceDefenseGap:
		payload, envelope, err := s.materializeDefenseGapReadback(ctx, descriptor, filter, audience)
		if err != nil {
			return readbackResponse{}, err
		}
		return buildReadbackResponse(resourceType, resourceID, audience, payload, envelope, fmt.Sprintf("Defense-gap readback for %s.", descriptor.SubjectRef), topologyContext), nil
	case readbackResourcePolicyReplay:
		payload, envelope, err := s.materializePolicyReplayReadback(ctx, descriptor, filter, audience)
		if err != nil {
			return readbackResponse{}, err
		}
		return buildReadbackResponse(resourceType, resourceID, audience, payload, envelope, fmt.Sprintf("Policy replay readback for %s.", descriptor.SubjectRef), topologyContext), nil
	case readbackResourceSystemicWeak:
		payload, envelope, err := s.materializeSystemicWeaknessReadback(ctx, descriptor, filter, audience)
		if err != nil {
			return readbackResponse{}, err
		}
		return buildReadbackResponse(resourceType, resourceID, audience, payload, envelope, fmt.Sprintf("Systemic weakness readback for %s.", descriptor.SubjectRef), topologyContext), nil
	default:
		return readbackResponse{}, readbackNotFound("")
	}
}

func buildReadbackResponse(resourceType string, resourceID string, audience string, payload any, envelope decisionEvidenceEnvelope, payloadSummary string, topologyContext *readbackTopologyContext) readbackResponse {
	return readbackResponse{
		ResourceType:       resourceType,
		ResourceID:         resourceID,
		PermanentURI:       fmt.Sprintf("/r/%s/%s", resourceType, resourceID),
		ProjectionAudience: audience,
		AdvisoryOnly:       true,
		PayloadSummary:     payloadSummary,
		EvidenceEnvelope:   envelope,
		Payload:            payload,
		TopologyContext:    topologyContext,
		Limitations: append([]string{
			"Readback is derived from the canonical advisory payload and its frozen evidence envelope hash, not from a separate report or archive truth layer.",
		}, envelope.Limitations...),
	}
}

func topologyFilterFromReadbackScope(scope readbackScope) (topologyFilter, error) {
	analyticsFilter, err := audit.NormalizeAnalyticsFilter(audit.AnalyticsFilter{
		Window:      "28d",
		CompareTo:   "previous_window",
		GroupBy:     "service",
		ClusterID:   scope.ClusterID,
		TenantID:    scope.TenantID,
		Environment: scope.Environment,
		Repo:        scope.Repository,
	})
	if err != nil {
		return topologyFilter{}, err
	}
	return topologyFilter{
		analytics: analyticsFilter,
		event: audit.EventFilter{
			ClusterID:   analyticsFilter.ClusterID,
			TenantID:    analyticsFilter.TenantID,
			Environment: analyticsFilter.Environment,
			Repo:        analyticsFilter.Repo,
			Limit:       topologyHistoryLimit,
		},
		Limit: 10,
	}, nil
}

func (s server) materializeReadbackTopologyContext(ctx context.Context, descriptor readbackDescriptor, filter incidentFilter, audience string) *readbackTopologyContext {
	topologyFilter, err := topologyFilterFromReadbackScope(descriptor.Scope)
	if err != nil {
		return &readbackTopologyContext{
			AdvisoryOnly: true,
			SubjectType:  descriptor.SubjectType,
			SubjectRef:   descriptor.SubjectRef,
			Limitations: []string{
				"Topology context could not be normalized for this readback scope and remains advisory-only when present.",
			},
		}
	}

	var response topologyBlastRadiusResponse
	switch descriptor.SubjectType {
	case "incident":
		incident, err := s.getIncidentByID(ctx, descriptor.SubjectRef, filter)
		if err != nil {
			return nil
		}
		response, err = s.buildIncidentBlastRadiusResponse(ctx, topologyFilter, incident)
		if err != nil {
			return nil
		}
	case "metric":
		metricFilter := filter
		metricFilter.ScorecardRef = descriptor.SubjectRef
		incidents, err := s.listIncidents(ctx, metricFilter)
		if err != nil {
			return nil
		}
		response, err = s.buildMetricBlastRadiusResponse(ctx, topologyFilter, descriptor.SubjectRef, incidents)
		if err != nil {
			return nil
		}
	case "scope":
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			return nil
		}
		response, err = s.buildScopedBlastRadiusResponse(ctx, topologyFilter, "scope", descriptor.SubjectRef, incidents)
		if err != nil {
			return nil
		}
	case "cluster":
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			return nil
		}
		weaknesses := attachSystemicWeaknessReadback(buildSystemicWeaknessResponse(incidents, "readback scope"), filter)
		for _, weakness := range weaknesses.Weaknesses {
			if weakness.PatternKey != descriptor.SubjectRef {
				continue
			}
			related := selectIncidentsByID(incidents, weakness.RelatedIncidentRefs)
			response, err = s.buildScopedBlastRadiusResponse(ctx, topologyFilter, "systemic_weakness", descriptor.SubjectRef, related)
			if err != nil {
				return nil
			}
			break
		}
	default:
		return nil
	}
	return buildReadbackTopologyContext(response, audience)
}

func buildReadbackTopologyContext(response topologyBlastRadiusResponse, audience string) *readbackTopologyContext {
	if response.SubjectRef == "" && response.PrimaryAffectedNode == nil && response.BlastRadiusScore == 0 && len(response.Limitations) == 0 {
		return nil
	}
	limitations := append([]string{}, response.Limitations...)
	limitations = append(limitations, "Topology context is a derived advisory snapshot at readback time; it does not replace canonical incident, evidence, or report truth.")
	return &readbackTopologyContext{
		AdvisoryOnly:         true,
		SubjectType:          response.SubjectType,
		SubjectRef:           response.SubjectRef,
		PrimaryAffectedNode:  projectReadbackTopologyNode(response.PrimaryAffectedNode, audience),
		BlastRadiusScore:     response.BlastRadiusScore,
		CriticalReachCount:   response.CriticalReachCount,
		TopRiskPathSummaries: summarizeReadbackTopologyPaths(response.TopRiskPaths, 3),
		Limitations:          uniqueStrings(limitations),
	}
}

func projectReadbackTopologyNode(node *topologyNode, audience string) *readbackTopologyNodeSummary {
	if node == nil {
		return nil
	}
	projected := &readbackTopologyNodeSummary{
		Service:          node.Service,
		Environment:      node.Environment,
		PublicExposure:   node.PublicExposure,
		SensitivityClass: node.SensitivityClass,
	}
	if audience == incidentAudienceInternal {
		projected.NodeID = node.NodeID
		projected.Namespace = node.Namespace
		projected.Cluster = node.Cluster
	} else if audience == incidentAudienceAuditorSafe {
		projected.Namespace = node.Namespace
	}
	return projected
}

func summarizeReadbackTopologyPaths(paths []topologyRiskPath, limit int) []string {
	summaries := make([]string, 0, minInt(len(paths), limit))
	for _, path := range paths[:minInt(len(paths), limit)] {
		if summary := strings.TrimSpace(path.Summary); summary != "" {
			summaries = append(summaries, summary)
		}
	}
	return uniqueStrings(summaries)
}

func (s server) materializeDefenseGapReadback(ctx context.Context, descriptor readbackDescriptor, filter incidentFilter, audience string) (defenseGapAssessment, decisionEvidenceEnvelope, error) {
	var payload defenseGapAssessment
	switch descriptor.SubjectType {
	case "incident":
		incident, err := s.getIncidentByID(ctx, descriptor.SubjectRef, filter)
		if err != nil {
			return payload, decisionEvidenceEnvelope{}, err
		}
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			return payload, decisionEvidenceEnvelope{}, err
		}
		payload = attachDefenseGapReadback(buildIncidentDefenseGapAssessment(incident, incidents), filter)
	case "metric":
		filter.ScorecardRef = descriptor.SubjectRef
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			return payload, decisionEvidenceEnvelope{}, err
		}
		payload = attachDefenseGapReadback(buildMetricDefenseGapAssessment(descriptor.SubjectRef, incidents), filter)
	default:
		return payload, decisionEvidenceEnvelope{}, readbackNotFound("")
	}
	payload = projectDefenseGapAssessment(payload, audience)
	envelope := buildDefenseGapEnvelope(payload, descriptor)
	payload.Readback = buildAdvisoryReadbackRef(readbackResourceDefenseGap, descriptor, envelope)
	return payload, envelope, nil
}

func (s server) materializePolicyReplayReadback(ctx context.Context, descriptor readbackDescriptor, filter incidentFilter, audience string) (policyReplayAssessment, decisionEvidenceEnvelope, error) {
	var payload policyReplayAssessment
	switch descriptor.SubjectType {
	case "incident":
		incident, err := s.getIncidentByID(ctx, descriptor.SubjectRef, filter)
		if err != nil {
			return payload, decisionEvidenceEnvelope{}, err
		}
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			return payload, decisionEvidenceEnvelope{}, err
		}
		payload = attachPolicyReplayReadback(buildIncidentPolicyReplayAssessment(incident, incidents), filter)
	case "metric":
		filter.ScorecardRef = descriptor.SubjectRef
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			return payload, decisionEvidenceEnvelope{}, err
		}
		payload = attachPolicyReplayReadback(buildMetricPolicyReplayAssessment(descriptor.SubjectRef, incidents), filter)
	case "scope":
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			return payload, decisionEvidenceEnvelope{}, err
		}
		payload = attachPolicyReplayReadback(buildScopePolicyReplayAssessment(incidents), filter)
	default:
		return payload, decisionEvidenceEnvelope{}, readbackNotFound("")
	}
	payload = projectPolicyReplayAssessment(payload, audience)
	envelope := buildPolicyReplayEnvelope(payload, descriptor)
	payload.Readback = buildAdvisoryReadbackRef(readbackResourcePolicyReplay, descriptor, envelope)
	return payload, envelope, nil
}

func (s server) materializeSystemicWeaknessReadback(ctx context.Context, descriptor readbackDescriptor, filter incidentFilter, audience string) (systemicWeakness, decisionEvidenceEnvelope, error) {
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		return systemicWeakness{}, decisionEvidenceEnvelope{}, err
	}
	response := attachSystemicWeaknessReadback(buildSystemicWeaknessResponse(incidents, "readback scope"), filter)
	for _, weakness := range response.Weaknesses {
		if weakness.PatternKey == descriptor.SubjectRef {
			projected := projectSystemicWeakness(weakness, audience)
			envelope := buildSystemicWeaknessEnvelope(projected, descriptor, response.GeneratedAt)
			projected.Readback = buildAdvisoryReadbackRef(readbackResourceSystemicWeak, descriptor, envelope)
			return projected, envelope, nil
		}
	}
	return systemicWeakness{}, decisionEvidenceEnvelope{}, readbackNotFound("")
}

func projectDefenseGapAssessment(payload defenseGapAssessment, audience string) defenseGapAssessment {
	if audience == incidentAudienceInternal {
		return payload
	}
	payload.Limitations = append(payload.Limitations, "This projection is audience-aware and may omit evidence refs or related incident refs from the internal view.")
	for index := range payload.DefenseGaps {
		payload.DefenseGaps[index].EvidenceRefs = nil
		if audience == incidentAudienceCustomerSafe {
			payload.DefenseGaps[index].RelatedIncidentRefs = nil
		}
	}
	if audience == incidentAudienceCustomerSafe {
		payload.SystemicPattern.RelatedIncidentRefs = nil
	}
	return payload
}

func projectPolicyReplayAssessment(payload policyReplayAssessment, audience string) policyReplayAssessment {
	if audience == incidentAudienceInternal {
		return payload
	}
	payload.Limitations = append(payload.Limitations, "This projection is audience-aware and may omit supporting evidence refs or scope markers from the internal view.")
	for index := range payload.ReplayResults {
		payload.ReplayResults[index].SupportingEvidenceRefs = nil
	}
	for index := range payload.CoverageGaps {
		payload.CoverageGaps[index].EvidenceRefs = nil
		if audience == incidentAudienceCustomerSafe {
			payload.CoverageGaps[index].RelatedIncidentRefs = nil
		}
	}
	if audience == incidentAudienceCustomerSafe {
		payload.BlastRadius.TopScopes = nil
	}
	return payload
}

func projectSystemicWeakness(payload systemicWeakness, audience string) systemicWeakness {
	if audience == incidentAudienceInternal {
		return payload
	}
	payload.Limitations = append(payload.Limitations, "This projection is audience-aware and may omit supporting evidence refs or related incident refs from the internal view.")
	payload.EvidenceRefs = nil
	if audience == incidentAudienceCustomerSafe {
		payload.RelatedIncidentRefs = nil
	}
	return payload
}
