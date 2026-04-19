package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/evidence"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/signingidentity"
)

type signingIdentityPoliciesResponse struct {
	Policies []signingidentity.Policy `json:"policies"`
}

type signingIdentityPolicyResponse struct {
	Status string                 `json:"status"`
	Policy signingidentity.Policy `json:"policy"`
}

type signingIdentityObservationsResponse struct {
	Items []signingidentity.Observation `json:"items"`
}

type signingIdentityFindingsResponse struct {
	Items []signingidentity.Finding `json:"items"`
}

type signingIdentityStatusResponse struct {
	Status signingidentity.StatusSummary `json:"status"`
}

type signingIdentityEvaluateRequest struct {
	Issuer            string     `json:"issuer"`
	SignerIdentity    string     `json:"signer_identity"`
	Subject           string     `json:"subject,omitempty"`
	Repository        string     `json:"repository,omitempty"`
	Workflow          string     `json:"workflow,omitempty"`
	Ref               string     `json:"ref,omitempty"`
	TenantID          string     `json:"tenant_id,omitempty"`
	ClusterID         string     `json:"cluster_id,omitempty"`
	Environment       string     `json:"environment,omitempty"`
	TransparencyState string     `json:"transparency_state,omitempty"`
	EvidenceAt        *time.Time `json:"evidence_at,omitempty"`
}

type signingIdentityDistrustRequest struct {
	DistrustedAfter *time.Time `json:"distrusted_after,omitempty"`
	Reason          string     `json:"reason,omitempty"`
}

func loadSigningIdentityConfig() (signingidentity.Config, error) {
	return signingidentity.ParseConfig(os.Getenv)
}

func (s server) signingIdentityObservationsHandler(w http.ResponseWriter, r *http.Request) {
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
	config, err := loadSigningIdentityConfig()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	policies, err := s.signingIdentityPolicies(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	observations, err := s.signingIdentityObservations(ctx, config, policies, parseObservationFilter(r))
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, signingIdentityObservationsResponse{Items: observations})
}

func (s server) signingIdentityPoliciesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		policies, err := s.signingIdentityPolicies(ctx)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		filtered := filterPoliciesForPrincipal(principal, policies)
		httpjson.Write(w, http.StatusOK, signingIdentityPoliciesResponse{Policies: filtered})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var request signingidentity.CreatePolicyRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request, err := applyPrincipalTenantToSigningPolicyRequest(principal, request)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		policy, err := signingidentity.NewPolicy(request, principal.Subject, time.Now().UTC())
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		currentPolicies, err := s.signingIdentityPolicies(ctx)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		for _, existing := range currentPolicies {
			if existing.ID == policy.ID {
				httpjson.Write(w, http.StatusConflict, map[string]string{"error": "signing identity policy already exists"})
				return
			}
		}
		if err := s.storeSigningIdentityPolicyEvent(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeSigningIdentityPolicyRecorded, policy, "signing identity policy recorded"); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, signingIdentityPolicyResponse{Status: "created", Policy: policy})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) signingIdentityPolicyByIDHandler(w http.ResponseWriter, r *http.Request) {
	policyID, action, err := signingIdentityPolicyPath(r.URL.Path)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	switch {
	case r.Method == http.MethodPost && action == "distrust":
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var request signingIdentityDistrustRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		policies, err := s.signingIdentityPolicies(ctx)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		policy, ok := findPolicyForPrincipal(principal, policies, policyID)
		if !ok {
			httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "signing identity policy not found"})
			return
		}
		distrustedAt := time.Now().UTC()
		if request.DistrustedAfter != nil && !request.DistrustedAfter.IsZero() {
			distrustedAt = request.DistrustedAfter.UTC()
		}
		policy.DistrustedAfter = &distrustedAt
		policy.DistrustReason = strings.TrimSpace(request.Reason)
		policy.UpdatedAt = time.Now().UTC()
		policy.UpdatedBy = principal.Subject
		if err := s.storeSigningIdentityPolicyEvent(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeSigningIdentityPolicyDistrusted, policy, firstNonEmpty(policy.DistrustReason, "signing identity policy distrusted")); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, signingIdentityPolicyResponse{Status: "distrusted", Policy: policy})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) signingIdentityStatusHandler(w http.ResponseWriter, r *http.Request) {
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
	config, err := loadSigningIdentityConfig()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	policies, err := s.signingIdentityPolicies(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	observations, err := s.signingIdentityObservations(ctx, config, policies, parseObservationFilter(r))
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	findings, err := s.signingIdentityFindings(ctx, config, policies, observations)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, signingIdentityStatusResponse{Status: signingIdentityStatus(config, policies, observations, findings)})
}

func (s server) signingIdentityFindingsHandler(w http.ResponseWriter, r *http.Request) {
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
	config, err := loadSigningIdentityConfig()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	policies, err := s.signingIdentityPolicies(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	observations, err := s.signingIdentityObservations(ctx, config, policies, parseObservationFilter(r))
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	findings, err := s.signingIdentityFindings(ctx, config, policies, observations)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, signingIdentityFindingsResponse{Items: findings})
}

func (s server) signingIdentityObservationByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	observationID := strings.TrimPrefix(strings.TrimSpace(r.URL.Path), "/v1/signing-identities/")
	observationID = strings.Trim(observationID, "/")
	if observationID == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "signing identity observation id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	config, err := loadSigningIdentityConfig()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	policies, err := s.signingIdentityPolicies(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	observations, err := s.signingIdentityObservations(ctx, config, policies, parseObservationFilter(r))
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	for _, observation := range observations {
		if observation.ID == observationID {
			httpjson.Write(w, http.StatusOK, observation)
			return
		}
	}
	httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "signing identity observation not found"})
}

func (s server) signingIdentityEvaluateHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleService, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request signingIdentityEvaluateRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	config, err := loadSigningIdentityConfig()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	policies, err := s.signingIdentityPolicies(ctx)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	filtered := filterPoliciesForPrincipal(principal, policies)
	decision := signingidentity.Evaluate(config, filtered, signingidentity.DecisionInput{
		Issuer:            request.Issuer,
		SignerIdentity:    request.SignerIdentity,
		Subject:           request.Subject,
		Repository:        request.Repository,
		Workflow:          request.Workflow,
		Ref:               request.Ref,
		TenantID:          request.TenantID,
		ClusterID:         request.ClusterID,
		Environment:       request.Environment,
		TransparencyState: request.TransparencyState,
		EvidenceAt:        request.EvidenceAt,
	})
	httpjson.Write(w, http.StatusOK, decision)
}

func (s server) signingIdentityPolicies(ctx context.Context) ([]signingidentity.Policy, error) {
	records, err := s.store.ListEvents(ctx, audit.EventFilter{EventType: audit.EventTypeSigningIdentityPolicyRecorded, Limit: 500})
	if err != nil {
		return nil, err
	}
	distrustRecords, err := s.store.ListEvents(ctx, audit.EventFilter{EventType: audit.EventTypeSigningIdentityPolicyDistrusted, Limit: 500})
	if err != nil {
		return nil, err
	}
	records = append(records, distrustRecords...)
	sort.Slice(records, func(i, j int) bool {
		return records[i].ReceivedAt.Before(records[j].ReceivedAt)
	})

	policies := map[string]signingidentity.Policy{}
	for _, record := range records {
		identityEvidence := record.Evidence
		if identityEvidence == nil || identityEvidence.SigningIdentity == nil {
			continue
		}
		policy, ok := policyFromEvent(record.Event)
		if !ok {
			continue
		}
		switch record.EventType {
		case audit.EventTypeSigningIdentityPolicyRecorded:
			policies[policy.ID] = policy
		case audit.EventTypeSigningIdentityPolicyDistrusted:
			existing := policies[policy.ID]
			if existing.ID == "" {
				existing = policy
			}
			existing.DistrustedAfter = policy.DistrustedAfter
			existing.DistrustReason = policy.DistrustReason
			existing.UpdatedAt = policy.UpdatedAt
			existing.UpdatedBy = policy.UpdatedBy
			policies[existing.ID] = existing
		}
	}

	items := make([]signingidentity.Policy, 0, len(policies))
	for _, policy := range policies {
		items = append(items, policy)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].UpdatedAt.Equal(items[j].UpdatedAt) {
			return items[i].ID < items[j].ID
		}
		return items[i].UpdatedAt.After(items[j].UpdatedAt)
	})
	return items, nil
}

func (s server) signingIdentityObservations(ctx context.Context, config signingidentity.Config, policies []signingidentity.Policy, filter map[string]string) ([]signingidentity.Observation, error) {
	records, err := s.store.ListEvents(ctx, audit.EventFilter{EventType: audit.EventTypePolicyDecision, Limit: 500})
	if err != nil {
		return nil, err
	}
	verificationRecords, err := s.store.ListEvents(ctx, audit.EventFilter{EventType: audit.EventTypeArtifactVerificationResult, Limit: 500})
	if err != nil {
		return nil, err
	}
	records = append(records, verificationRecords...)
	sort.Slice(records, func(i, j int) bool {
		return records[i].ReceivedAt.Before(records[j].ReceivedAt)
	})

	observations := map[string]*signingidentity.Observation{}
	artifactDigests := map[string]map[string]struct{}{}
	for _, record := range records {
		if !recordMatchesTenantFilter(record.Event, filter) {
			continue
		}
		artifact := record.Evidence
		if artifact == nil || artifact.Artifact == nil {
			continue
		}
		input := signingidentity.DecisionInput{
			Issuer:            firstNonEmpty(artifact.Artifact.Issuer, record.Event.Evidence.Artifact.Issuer),
			SignerIdentity:    artifact.Artifact.SignerIdentity,
			Subject:           artifact.Artifact.Subject,
			Repository:        firstNonEmpty(artifact.Artifact.Repository, record.Repo),
			Workflow:          artifact.Artifact.Workflow,
			Ref:               firstNonEmpty(artifact.Artifact.Ref, record.Branch),
			TenantID:          record.TenantID,
			ClusterID:         record.ClusterID,
			Environment:       record.Environment,
			TransparencyState: artifact.VerificationState,
			EvidenceAt:        evidenceTimestamp(artifact.Bundle, record.Timestamp),
		}
		digest := firstNonEmpty(artifact.Artifact.Digest, record.Digest)
		if digestFilter := strings.TrimSpace(filter["image_digest"]); digestFilter != "" && digest != digestFilter {
			continue
		}
		if repoFilter := strings.TrimSpace(filter["repo"]); repoFilter != "" && input.Repository != repoFilter {
			continue
		}
		if workflowFilter := strings.TrimSpace(filter["workflow"]); workflowFilter != "" && input.Workflow != workflowFilter {
			continue
		}
		observationID := signingidentity.ObservationID(input, digest)
		decision := signingidentity.Evaluate(config, policies, input)
		observation := observations[observationID]
		if observation == nil {
			observation = &signingidentity.Observation{
				ID:                observationID,
				ProviderType:      providerTypeFromIdentity(input.SignerIdentity),
				Issuer:            input.Issuer,
				SignerIdentity:    input.SignerIdentity,
				Subject:           input.Subject,
				Repository:        input.Repository,
				Workflow:          input.Workflow,
				Ref:               input.Ref,
				CommitSHA:         artifact.Artifact.CommitSHA,
				ImageDigest:       digest,
				TenantID:          record.TenantID,
				ClusterID:         record.ClusterID,
				Environment:       record.Environment,
				VerificationState: artifact.VerificationState,
				Authorized:        decision.Authorized,
				MatchedPolicyID:   decision.MatchedPolicyID,
				DistrustedAfter:   decision.DistrustedAfter,
				ReasonCode:        decision.ReasonCode,
				ReasonDetail:      decision.ReasonDetail,
			}
			observations[observationID] = observation
			artifactDigests[observationID] = map[string]struct{}{}
		}
		if observation.FirstSeenAt == nil || record.ReceivedAt.Before(*observation.FirstSeenAt) {
			timestamp := record.ReceivedAt.UTC()
			observation.FirstSeenAt = &timestamp
		}
		if observation.LastSeenAt == nil || record.ReceivedAt.After(*observation.LastSeenAt) {
			timestamp := record.ReceivedAt.UTC()
			observation.LastSeenAt = &timestamp
		}
		observation.EventCount++
		if digest != "" {
			artifactDigests[observationID][digest] = struct{}{}
			observation.ImageDigest = digest
		}
	}

	items := make([]signingidentity.Observation, 0, len(observations))
	for id, observation := range observations {
		observation.ArtifactCount = len(artifactDigests[id])
		items = append(items, *observation)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].LastSeenAt == nil || items[j].LastSeenAt == nil {
			return items[i].ID < items[j].ID
		}
		return items[i].LastSeenAt.After(*items[j].LastSeenAt)
	})
	limit := 100
	if raw := strings.TrimSpace(filter["limit"]); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 && parsed <= 500 {
			limit = parsed
		}
	}
	if len(items) > limit {
		items = items[:limit]
	}
	return items, nil
}

func (s server) signingIdentityFindings(ctx context.Context, _ signingidentity.Config, policies []signingidentity.Policy, observations []signingidentity.Observation) ([]signingidentity.Finding, error) {
	findings := make([]signingidentity.Finding, 0)
	for _, observation := range observations {
		switch observation.ReasonCode {
		case signingidentity.ReasonAuthorized:
			continue
		case signingidentity.ReasonDistrustedAfterCutoff:
			findings = append(findings, signingidentity.Finding{
				ID:            observation.ID + ":distrusted",
				Type:          signingidentity.FindingDistrustedIdentity,
				Severity:      "high",
				Repository:    observation.Repository,
				Workflow:      observation.Workflow,
				Ref:           observation.Ref,
				ObservationID: observation.ID,
				Reason:        observation.ReasonDetail,
				DetectedAt:    observation.LastSeenAt,
				Advisory:      false,
			})
		case signingidentity.ReasonTransparencyUnverified, signingidentity.ReasonRekorRequired:
			findings = append(findings, signingidentity.Finding{
				ID:            observation.ID + ":transparency",
				Type:          signingidentity.FindingTransparencyMissing,
				Severity:      "medium",
				Repository:    observation.Repository,
				Workflow:      observation.Workflow,
				Ref:           observation.Ref,
				ObservationID: observation.ID,
				Reason:        observation.ReasonDetail,
				DetectedAt:    observation.LastSeenAt,
				Advisory:      false,
			})
		default:
			findings = append(findings, signingidentity.Finding{
				ID:            observation.ID + ":unauthorized",
				Type:          signingidentity.FindingUnauthorizedIdentity,
				Severity:      "high",
				Repository:    observation.Repository,
				Workflow:      observation.Workflow,
				Ref:           observation.Ref,
				ObservationID: observation.ID,
				Reason:        observation.ReasonDetail,
				DetectedAt:    observation.LastSeenAt,
				Advisory:      false,
			})
		}
	}

	workflowDocs, err := signingidentity.ScanWorkflowDocuments(policyWorkflowsDir())
	if err != nil {
		return nil, err
	}
	findings = append(findings, signingidentity.BuildWorkflowFindings(workflowDocs, policies, time.Now().UTC())...)
	sort.Slice(findings, func(i, j int) bool {
		if findings[i].DetectedAt == nil || findings[j].DetectedAt == nil {
			return findings[i].ID < findings[j].ID
		}
		return findings[i].DetectedAt.After(*findings[j].DetectedAt)
	})
	return findings, nil
}

func signingIdentityStatus(config signingidentity.Config, policies []signingidentity.Policy, observations []signingidentity.Observation, findings []signingidentity.Finding) signingidentity.StatusSummary {
	summary := signingidentity.StatusSummary{
		EnforcementMode:    config.Enforcement,
		RequireRekor:       config.RequireRekor,
		TotalPolicies:      len(policies),
		ObservedIdentities: len(observations),
		Findings:           len(findings),
		CountsByReasonCode: map[string]int{},
	}
	for _, policy := range policies {
		if policy.Enabled {
			summary.EnabledPolicies++
		}
	}
	for _, observation := range observations {
		switch observation.Authorized {
		case signingidentity.AuthorizationAuthorized:
			summary.Authorized++
		case signingidentity.AuthorizationUnauthorized:
			summary.Unauthorized++
		default:
			summary.Unknown++
		}
		if observation.ReasonCode != "" {
			summary.CountsByReasonCode[observation.ReasonCode]++
		}
	}
	for _, finding := range findings {
		if finding.Advisory {
			summary.WorkflowDriftFindings++
		}
	}
	return summary
}

func (s server) storeSigningIdentityPolicyEvent(ctx context.Context, requestID string, actor string, eventType string, policy signingidentity.Policy, reason string) error {
	event := audit.Event{
		RequestID:   requestID,
		Component:   "audit-writer",
		EventType:   eventType,
		Actor:       actor,
		TenantID:    policy.TenantID,
		ClusterID:   policy.ClusterID,
		Repo:        policy.Repository,
		Environment: policy.Environment,
		Decision:    audit.DecisionAllow,
		Reasons:     []string{reason},
		Evidence: &audit.Evidence{
			SigningIdentity: &audit.SigningIdentityEvidence{
				PolicyID:        policy.ID,
				PolicyName:      policy.Name,
				ProviderType:    policy.ProviderType,
				Issuer:          policy.Issuer,
				SignerIdentity:  policy.SignerIdentity,
				Subject:         policy.Subject,
				Repository:      policy.Repository,
				Workflow:        policy.Workflow,
				Ref:             policy.Ref,
				DistrustedAfter: policy.DistrustedAfter,
				ReasonDetail:    policy.DistrustReason,
			},
		},
	}
	_, err := s.store.Ingest(ctx, event)
	return err
}

func policyFromEvent(event audit.Event) (signingidentity.Policy, bool) {
	if event.Evidence == nil || event.Evidence.SigningIdentity == nil {
		return signingidentity.Policy{}, false
	}
	identity := event.Evidence.SigningIdentity
	policy := signingidentity.Policy{
		ID:              strings.TrimSpace(identity.PolicyID),
		Name:            strings.TrimSpace(identity.PolicyName),
		ProviderType:    strings.TrimSpace(identity.ProviderType),
		Issuer:          strings.TrimSpace(identity.Issuer),
		SignerIdentity:  strings.TrimSpace(identity.SignerIdentity),
		Subject:         strings.TrimSpace(identity.Subject),
		Repository:      strings.TrimSpace(identity.Repository),
		Workflow:        strings.TrimSpace(identity.Workflow),
		Ref:             strings.TrimSpace(identity.Ref),
		TenantID:        strings.TrimSpace(event.TenantID),
		ClusterID:       strings.TrimSpace(event.ClusterID),
		Environment:     strings.TrimSpace(event.Environment),
		Enabled:         true,
		DistrustedAfter: identity.DistrustedAfter,
		DistrustReason:  strings.TrimSpace(identity.ReasonDetail),
		CreatedAt:       event.Timestamp.UTC(),
		UpdatedAt:       event.Timestamp.UTC(),
		CreatedBy:       event.Actor,
		UpdatedBy:       event.Actor,
	}
	if policy.ID == "" {
		policy.ID = signingidentity.PolicyID(policy)
	}
	return policy, true
}

func filterPoliciesForPrincipal(principal auth.Principal, policies []signingidentity.Policy) []signingidentity.Policy {
	if principal.GlobalScope || principal.TenantID == "" {
		return policies
	}
	filtered := make([]signingidentity.Policy, 0, len(policies))
	for _, policy := range policies {
		if policy.TenantID == "" || policy.TenantID == principal.TenantID {
			filtered = append(filtered, policy)
		}
	}
	return filtered
}

func findPolicyForPrincipal(principal auth.Principal, policies []signingidentity.Policy, policyID string) (signingidentity.Policy, bool) {
	for _, policy := range filterPoliciesForPrincipal(principal, policies) {
		if policy.ID == policyID {
			return policy, true
		}
	}
	return signingidentity.Policy{}, false
}

func applyPrincipalTenantToSigningPolicyRequest(principal auth.Principal, request signingidentity.CreatePolicyRequest) (signingidentity.CreatePolicyRequest, error) {
	if principal.GlobalScope || principal.TenantID == "" {
		return request, nil
	}
	if strings.TrimSpace(request.TenantID) != "" && strings.TrimSpace(request.TenantID) != principal.TenantID {
		return request, auth.ErrInsufficientPermissions
	}
	request.TenantID = principal.TenantID
	return request, nil
}

func parseObservationFilter(r *http.Request) map[string]string {
	query := r.URL.Query()
	return map[string]string{
		"tenant_id":    strings.TrimSpace(query.Get("tenant_id")),
		"cluster_id":   strings.TrimSpace(query.Get("cluster_id")),
		"environment":  strings.TrimSpace(query.Get("environment")),
		"repo":         strings.TrimSpace(query.Get("repo")),
		"workflow":     strings.TrimSpace(query.Get("workflow")),
		"image_digest": strings.TrimSpace(query.Get("image_digest")),
		"limit":        firstNonEmpty(query.Get("limit"), "100"),
	}
}

func recordMatchesTenantFilter(event audit.Event, filter map[string]string) bool {
	if tenantID := strings.TrimSpace(filter["tenant_id"]); tenantID != "" && event.TenantID != tenantID {
		return false
	}
	if clusterID := strings.TrimSpace(filter["cluster_id"]); clusterID != "" && event.ClusterID != clusterID {
		return false
	}
	if environment := strings.TrimSpace(filter["environment"]); environment != "" && event.Environment != environment {
		return false
	}
	return true
}

func evidenceTimestamp(bundle *evidence.Bundle, fallback time.Time) *time.Time {
	if bundle != nil {
		if bundle.IntegratedTime != nil && !bundle.IntegratedTime.IsZero() {
			timestamp := bundle.IntegratedTime.UTC()
			return &timestamp
		}
		if bundle.SignedAt != nil && !bundle.SignedAt.IsZero() {
			timestamp := bundle.SignedAt.UTC()
			return &timestamp
		}
	}
	if fallback.IsZero() {
		return nil
	}
	timestamp := fallback.UTC()
	return &timestamp
}

func providerTypeFromIdentity(identity string) string {
	if strings.Contains(identity, "github.com/") {
		return signingidentity.ProviderGitHubOIDC
	}
	return signingidentity.ProviderGenericOIDC
}

func signingIdentityPolicyPath(path string) (string, string, error) {
	trimmed := strings.TrimPrefix(strings.TrimSpace(path), "/v1/signing-identities/policies/")
	if trimmed == "" || trimmed == path {
		return "", "", errors.New("signing identity policy id is required")
	}
	parts := strings.Split(strings.Trim(trimmed, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return "", "", errors.New("signing identity policy id is required")
	}
	if len(parts) == 1 {
		return parts[0], "", nil
	}
	return parts[0], parts[1], nil
}

func policyWorkflowsDir() string {
	config, err := loadSigningIdentityConfig()
	if err != nil {
		return ".github/workflows"
	}
	return config.WorkflowsDir
}
