package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	federationResilienceRateLimitWindow    = time.Hour
	federationResilienceMaxRequestsPerHour = 6
	federationResilienceRejectBurstOpen    = 3
	validationMaxScenarioSelection         = 8
	validationReadinessMaxListedRuns       = 20
	validationReadinessQuotaCPU            = "500m"
	validationReadinessQuotaMemory         = "512Mi"
	validationReadinessQuotaEphemeral      = "256Mi"
	validationReadinessMaxExecutionMinutes = 15
	handoffSignerBackendContractVersion    = "handoff.signer_backend.v1"
	handoffSignerBackendDeterministicSeed  = "deterministic_seed"
	handoffSignerBackendKMSAdapter         = "kms_adapter"
	handoffSignerBackendHSMAdapter         = "hsm_adapter"
)

type handoffQualityGateCheck struct {
	GateID       string   `json:"gate_id"`
	Status       string   `json:"status"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
	Limitations  []string `json:"limitations,omitempty"`
}

type handoffQualityGatesResponse struct {
	SchemaVersion          string                    `json:"schema_version"`
	GeneratedAt            time.Time                 `json:"generated_at"`
	SignerBackendContract  string                    `json:"signer_backend_contract"`
	ActiveSignerBackend    string                    `json:"active_signer_backend"`
	SupportedBackendModels []string                  `json:"supported_backend_models"`
	OfflineVerifySupported bool                      `json:"offline_verify_supported"`
	ArchiveIntegrityPolicy []string                  `json:"archive_integrity_policy"`
	LongTermIntegrityModel []string                  `json:"long_term_integrity_model"`
	Gates                  []handoffQualityGateCheck `json:"gates"`
	Limitations            []string                  `json:"limitations,omitempty"`
}

type federationPeerResilienceState struct {
	PeerID                  string   `json:"peer_id"`
	Status                  string   `json:"status"`
	FreshnessStatus         string   `json:"freshness_status"`
	RateLimitStatus         string   `json:"rate_limit_status"`
	CircuitState            string   `json:"circuit_state"`
	DisclosureMode          string   `json:"disclosure_mode"`
	CompatibilityState      string   `json:"compatibility_state"`
	RecentRequests          int      `json:"recent_requests"`
	RecentRejectedDecisions int      `json:"recent_rejected_decisions"`
	ReasonSummary           string   `json:"reason_summary"`
	Limitations             []string `json:"limitations,omitempty"`
}

type federationResilienceResponse struct {
	SchemaVersion        string                          `json:"schema_version"`
	GeneratedAt          time.Time                       `json:"generated_at"`
	RateLimitWindowMin   int                             `json:"rate_limit_window_minutes"`
	MaxRequestsPerWindow int                             `json:"max_requests_per_window"`
	OpenCircuitPeers     []string                        `json:"open_circuit_peers,omitempty"`
	Peers                []federationPeerResilienceState `json:"peers"`
	Limitations          []string                        `json:"limitations,omitempty"`
}

type validationResourceQuota struct {
	MaxScenariosPerRun int    `json:"max_scenarios_per_run"`
	MaxListedRuns      int    `json:"max_listed_runs"`
	CPU                string `json:"cpu"`
	Memory             string `json:"memory"`
	EphemeralStorage   string `json:"ephemeral_storage"`
	MaxExecutionMin    int    `json:"max_execution_minutes"`
}

type validationReadinessResponse struct {
	SchemaVersion        string                  `json:"schema_version"`
	GeneratedAt          time.Time               `json:"generated_at"`
	IsolationModel       []string                `json:"isolation_model"`
	ResourceQuota        validationResourceQuota `json:"resource_quota"`
	RegressionGate       string                  `json:"regression_gate_status"`
	FlakyScenarios       []string                `json:"flaky_scenarios,omitempty"`
	CertificateStability string                  `json:"certificate_stability"`
	SealReadyOutputs     bool                    `json:"seal_ready_outputs"`
	Limitations          []string                `json:"limitations,omitempty"`
}

type selfAuditEvent struct {
	SchemaVersion string    `json:"schema_version"`
	EventID       int64     `json:"event_id"`
	Timestamp     time.Time `json:"timestamp"`
	Category      string    `json:"category"`
	Severity      string    `json:"severity"`
	Action        string    `json:"action"`
	Actor         string    `json:"actor,omitempty"`
	Component     string    `json:"component"`
	SubjectRef    string    `json:"subject_ref,omitempty"`
	ResourceURI   string    `json:"resource_uri,omitempty"`
	EvidenceRefs  []string  `json:"evidence_refs,omitempty"`
	Limitations   []string  `json:"limitations,omitempty"`
}

type selfAuditSummaryResponse struct {
	SchemaVersion         string           `json:"schema_version"`
	GeneratedAt           time.Time        `json:"generated_at"`
	CountsByCategory      map[string]int   `json:"counts_by_category"`
	LatestCriticalActions []selfAuditEvent `json:"latest_critical_actions"`
	Limitations           []string         `json:"limitations,omitempty"`
}

type selfAuditEventsResponse struct {
	SchemaVersion string           `json:"schema_version"`
	Events        []selfAuditEvent `json:"events"`
	Limitations   []string         `json:"limitations,omitempty"`
}

func (s server) handoffQualityGatesHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildHandoffQualityGates(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) federationResilienceHandler(w http.ResponseWriter, r *http.Request) {
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
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildFederationResilience(ctx)
	if err != nil {
		writeFederationError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) validationReadinessHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildValidationReadiness(ctx, filter)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) selfAuditSummaryHandler(w http.ResponseWriter, r *http.Request) {
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
	items, err := s.listSelfAuditEvents(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	counts := map[string]int{}
	for _, item := range items {
		counts[item.Category]++
	}
	httpjson.Write(w, http.StatusOK, selfAuditSummaryResponse{
		SchemaVersion:         selfAuditSummarySchemaVersion,
		GeneratedAt:           selfAuditGeneratedAt(items),
		CountsByCategory:      counts,
		LatestCriticalActions: limitSelfAuditEvents(items, 8),
		Limitations: []string{
			"Self-audit is reconstructed from canonical audit events for ChangeLock-triggered policy, signing, federation, validation, and operator-critical changes; out-of-band infrastructure drift remains outside this bounded loop.",
		},
	})
}

func (s server) selfAuditEventsHandler(w http.ResponseWriter, r *http.Request) {
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
	items, err := s.listSelfAuditEvents(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, selfAuditEventsResponse{
		SchemaVersion: selfAuditEventsSchemaVersion,
		Events:        items,
		Limitations: []string{
			"Self-audit events are evidence-backed summaries over existing ChangeLock audit records and remain distinct from incident or runtime truth.",
		},
	})
}

func (s server) buildHandoffQualityGates(ctx context.Context, filter audit.EventFilter) (handoffQualityGatesResponse, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Component:   handoffComponent,
		Limit:       maxInt(filter.Limit, 200),
	})
	if err != nil {
		return handoffQualityGatesResponse{}, err
	}
	records := make([]handoffStoredRecord, 0, len(events))
	for _, event := range events {
		record := parseSecurityTimelineHandoff(event.Handoff)
		if record != nil {
			records = append(records, *record)
		}
	}
	sort.Slice(records, func(i, j int) bool {
		if !records[i].Session.InitiatedAt.Equal(records[j].Session.InitiatedAt) {
			return records[i].Session.InitiatedAt.After(records[j].Session.InitiatedAt)
		}
		return records[i].PackageID < records[j].PackageID
	})
	generatedAt := time.Unix(0, 0).UTC()
	if len(records) > 0 {
		generatedAt = records[0].Session.InitiatedAt.UTC()
	}
	manifestConflict := false
	verificationIssues := 0
	edgeStateIssues := 0
	archiveComplete := 0
	seenManifestHashes := map[string]string{}
	latestEvidence := []string{}
	for _, record := range records {
		if previous, ok := seenManifestHashes[record.PackageID]; ok && previous != record.ManifestHash {
			manifestConflict = true
		}
		seenManifestHashes[record.PackageID] = record.ManifestHash
		if record.Verification.OverallStatus != handoffVerificationValid {
			verificationIssues++
		}
		if !record.Verification.TimestampValid || !record.Verification.TransparencyValid {
			edgeStateIssues++
		}
		if record.Bundle.BundlePath != "" && record.ManifestHash != "" && len(record.Signatures) > 0 {
			archiveComplete++
		}
		if len(latestEvidence) == 0 {
			latestEvidence = limitStrings(record.Manifest.EvidenceRefs, 8)
		}
	}
	gates := []handoffQualityGateCheck{
		{
			GateID:       "deterministic_assembly",
			Status:       ternaryStatus(!manifestConflict, "pass", "degraded"),
			Summary:      ternarySummary(!manifestConflict, "Stored handoff packages preserve stable manifest hashes per package identity.", "Observed package identity reused with conflicting manifest hash; deterministic sealing contract requires investigation."),
			EvidenceRefs: latestEvidence,
		},
		{
			GateID:       "offline_verify",
			Status:       "pass",
			Summary:      "Offline verification remains available through canonical bundle verification and stored verification surfaces.",
			EvidenceRefs: latestEvidence,
		},
		{
			GateID:       "timestamp_transparency_edges",
			Status:       ternaryStatus(edgeStateIssues == 0, "pass", "watch"),
			Summary:      ternarySummary(edgeStateIssues == 0, "Timestamp and transparency edge-state handling stayed within the expected verification model.", fmt.Sprintf("%d package(s) surfaced timestamp or transparency edge conditions that remain explainable but require operator review.", edgeStateIssues)),
			EvidenceRefs: latestEvidence,
		},
		{
			GateID:       "archive_integrity",
			Status:       ternaryStatus(len(records) == 0 || archiveComplete == len(records), "pass", "watch"),
			Summary:      ternarySummary(len(records) == 0 || archiveComplete == len(records), "Stored bundles carry manifest, signature, and bundle references suitable for bounded archive integrity workflows.", "Some stored bundles are incomplete for full archive integrity review and should be re-sealed or re-verified before long-term retention."),
			EvidenceRefs: latestEvidence,
		},
		{
			GateID:       "verification_failures",
			Status:       ternaryStatus(verificationIssues == 0, "pass", "watch"),
			Summary:      ternarySummary(verificationIssues == 0, "Recent handoff packages verified cleanly.", fmt.Sprintf("%d stored handoff package(s) require follow-up because verification was partial or invalid.", verificationIssues)),
			EvidenceRefs: latestEvidence,
		},
	}
	return handoffQualityGatesResponse{
		SchemaVersion:         handoffQualitySchemaVersion,
		GeneratedAt:           generatedAt,
		SignerBackendContract: handoffSignerBackendContractVersion,
		ActiveSignerBackend:   handoffActiveSignerBackend(),
		SupportedBackendModels: []string{
			handoffSignerBackendDeterministicSeed,
			handoffSignerBackendKMSAdapter,
			handoffSignerBackendHSMAdapter,
		},
		OfflineVerifySupported: true,
		ArchiveIntegrityPolicy: []string{
			"Retain canonical manifest hash, bundle metadata, verification record, and evidence refs together.",
			"Treat partial or invalid verification states as archive blockers until re-verified or re-sealed.",
			"Preserve timestamp and transparency metadata for long-term freshness reasoning.",
		},
		LongTermIntegrityModel: []string{
			"Current deployment uses deterministic seed signing unless a stronger signer backend is configured.",
			"Offline verify remains the baseline integrity backstop for archive recall.",
			"Long-term retention should plan for re-timestamp or re-attest workflows rather than assuming perpetual external freshness proof.",
		},
		Gates: gates,
		Limitations: []string{
			"Handoff quality gates summarize production-hardening posture over stored bundle truth and existing verification records; they do not claim external TSA or transparency availability at read time.",
		},
	}, nil
}

func handoffActiveSignerBackend() string {
	if strings.TrimSpace(os.Getenv("CHANGELOCK_HANDOFF_SIGNING_SEED")) != "" {
		return handoffSignerBackendDeterministicSeed
	}
	return "unconfigured"
}

func (s server) buildFederationResilience(ctx context.Context) (federationResilienceResponse, error) {
	view, err := s.buildFederationGlobalView(ctx)
	if err != nil {
		return federationResilienceResponse{}, err
	}
	items := make([]federationPeerResilienceState, 0, len(view.Peers))
	openPeers := []string{}
	for _, peer := range view.Peers {
		state := federationPeerResilienceStateFromView(view, peer)
		items = append(items, state)
		if state.CircuitState == "open" {
			openPeers = append(openPeers, peer.PeerID)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].CircuitState != items[j].CircuitState {
			return items[i].CircuitState > items[j].CircuitState
		}
		if items[i].RateLimitStatus != items[j].RateLimitStatus {
			return items[i].RateLimitStatus > items[j].RateLimitStatus
		}
		return items[i].PeerID < items[j].PeerID
	})
	return federationResilienceResponse{
		SchemaVersion:        federationResilienceSchemaVersion,
		GeneratedAt:          time.Now().UTC(),
		RateLimitWindowMin:   int(federationResilienceRateLimitWindow / time.Minute),
		MaxRequestsPerWindow: federationResilienceMaxRequestsPerHour,
		OpenCircuitPeers:     openPeers,
		Peers:                items,
		Limitations: []string{
			"Federation resilience state remains bounded by local peer registry, proof history, and local policy state; remote proof acceptance never bypasses local admissibility policy.",
		},
	}, nil
}

func federationPeerResilienceStateFromView(view federationGlobalView, peer federationPeer) federationPeerResilienceState {
	now := time.Now().UTC()
	cutoff := now.Add(-federationResilienceRateLimitWindow)
	recentRequests := 0
	recentRejected := 0
	for _, item := range view.ProofHistory {
		if item.PeerID != peer.PeerID {
			continue
		}
		if federationHistorySortTime(item).Before(cutoff) {
			continue
		}
		recentRequests++
		if strings.HasPrefix(item.Decision, "rejected") || strings.EqualFold(item.Status, "rejected") {
			recentRejected++
		}
	}
	freshnessStatus := "fresh"
	if containsString(view.StalePeers, peer.PeerID) || strings.Contains(strings.ToLower(peer.Status), "stale") {
		freshnessStatus = "stale"
	} else if recentRequests >= federationResilienceMaxRequestsPerHour-2 {
		freshnessStatus = "watch"
	}
	rateLimitStatus := "ready"
	if recentRequests >= federationResilienceMaxRequestsPerHour {
		rateLimitStatus = "limited"
	}
	circuitState := "closed"
	reasons := []string{}
	if freshnessStatus == "stale" {
		circuitState = "open"
		reasons = append(reasons, "peer freshness is stale")
	}
	if recentRejected >= federationResilienceRejectBurstOpen {
		circuitState = "open"
		reasons = append(reasons, "recent rejected proof burst exceeded threshold")
	}
	if peer.PolicyRole == federationPolicyRoleLeader && view.PolicyState.SyncStatus == federationSyncStatusDiverged {
		circuitState = "open"
		reasons = append(reasons, "leader policy state is divergent")
	}
	compatibilityState := "compatible"
	if !containsString(peer.Capabilities, "sealed_handoff") {
		compatibilityState = "bounded"
		reasons = append(reasons, "peer does not advertise sealed_handoff capability")
	}
	if strings.TrimSpace(peer.DisclosureMode) == "" {
		reasons = append(reasons, "peer disclosure mode is not explicitly declared")
	}
	limitations := []string{}
	if strings.Contains(strings.ToLower(peer.DisclosureMode), "summary") {
		limitations = append(limitations, "Peer is configured for disclosure-minimized exchange and may provide only summary-safe proof context.")
	}
	return federationPeerResilienceState{
		PeerID:                  peer.PeerID,
		Status:                  peer.Status,
		FreshnessStatus:         freshnessStatus,
		RateLimitStatus:         rateLimitStatus,
		CircuitState:            circuitState,
		DisclosureMode:          firstNonEmpty(peer.DisclosureMode, "standard"),
		CompatibilityState:      compatibilityState,
		RecentRequests:          recentRequests,
		RecentRejectedDecisions: recentRejected,
		ReasonSummary:           firstNonEmpty(firstString(reasons), "peer is within bounded federation request, freshness, and policy limits"),
		Limitations:             limitations,
	}
}

func (s server) enforceFederationResilience(ctx context.Context, peerID string) error {
	view, err := s.buildFederationResilience(ctx)
	if err != nil {
		return err
	}
	for _, peer := range view.Peers {
		if peer.PeerID != peerID {
			continue
		}
		if peer.CircuitState == "open" {
			return fmt.Errorf("%w: %s", errFederationCircuitOpen, peer.ReasonSummary)
		}
		if peer.RateLimitStatus == "limited" {
			return fmt.Errorf("%w: %s", errFederationRateLimited, peer.ReasonSummary)
		}
		return nil
	}
	return errFederationPeerNotFound
}

func (s server) buildValidationReadiness(ctx context.Context, filter validationHarnessFilter) (validationReadinessResponse, error) {
	filter.Limit = maxInt(filter.Limit, validationReadinessMaxListedRuns)
	runs, limitations, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return validationReadinessResponse{}, err
	}
	flaky := map[string]struct{}{}
	regressionGate := "unverified"
	certificateStable := true
	sealReady := true
	for _, run := range runs {
		if run.Mode == validationModeRegression && regressionGate == "unverified" {
			switch run.Certificate.OverallStatus {
			case validationStatusPass:
				regressionGate = "passing"
			case validationStatusPartial, validationStatusFlaky:
				regressionGate = "review_required"
			default:
				regressionGate = "failing"
			}
		}
		if !run.Certificate.SealReady {
			sealReady = false
		}
		if run.Certificate.OverallStatus == validationStatusFlaky || run.Certificate.OverallStatus == validationStatusFail {
			certificateStable = false
		}
		for _, verdict := range run.Verdicts {
			if verdict.Status == validationStatusFlaky {
				flaky[verdict.ScenarioID] = struct{}{}
				certificateStable = false
			}
		}
	}
	flakyScenarios := make([]string, 0, len(flaky))
	for scenarioID := range flaky {
		flakyScenarios = append(flakyScenarios, scenarioID)
	}
	sort.Strings(flakyScenarios)
	return validationReadinessResponse{
		SchemaVersion: validationReadinessSchemaVersion,
		GeneratedAt:   time.Now().UTC(),
		IsolationModel: []string{
			"Validation runs execute in bounded shadow, twin, or compatibility-lab semantics and never claim production mutation.",
			"Scenario registry, cleanup plan, blast-radius limit, and rollback metadata remain first-class parts of execution evidence.",
			"Resource quotas bound scenario fan-out so validation cannot silently become an unbounded attack playground.",
		},
		ResourceQuota: validationResourceQuota{
			MaxScenariosPerRun: validationMaxScenarioSelection,
			MaxListedRuns:      validationReadinessMaxListedRuns,
			CPU:                validationReadinessQuotaCPU,
			Memory:             validationReadinessQuotaMemory,
			EphemeralStorage:   validationReadinessQuotaEphemeral,
			MaxExecutionMin:    validationReadinessMaxExecutionMinutes,
		},
		RegressionGate:       regressionGate,
		FlakyScenarios:       flakyScenarios,
		CertificateStability: ternarySummary(certificateStable, "stable", "watch"),
		SealReadyOutputs:     sealReady,
		Limitations: uniqueStrings(append([]string{
			"Validation readiness is derived from declarative scenario registry and audit-backed strict harness runs; it does not claim destructive production execution.",
		}, limitations...)),
	}, nil
}

func (s server) listSelfAuditEvents(ctx context.Context, filter audit.EventFilter) ([]selfAuditEvent, error) {
	events, err := s.store.ListEvents(ctx, securityTimelineEventFilter(filter, maxInt(filter.Limit, 120)))
	if err != nil {
		return nil, err
	}
	items := make([]selfAuditEvent, 0, len(events))
	for _, event := range events {
		category, severity, action, resourceURI := classifySelfAuditEvent(event)
		if category == "" {
			continue
		}
		items = append(items, selfAuditEvent{
			SchemaVersion: selfAuditEventSchemaVersion,
			EventID:       event.ID,
			Timestamp:     eventTimestamp(event).UTC(),
			Category:      category,
			Severity:      severity,
			Action:        action,
			Actor:         strings.TrimSpace(event.Actor),
			Component:     event.Component,
			SubjectRef:    firstNonEmpty(strings.TrimSpace(event.IncidentScopeRef), strings.TrimSpace(event.Workload), strings.TrimSpace(event.RecommendationSubjectRef), strings.TrimSpace(event.Repo)),
			ResourceURI:   resourceURI,
			EvidenceRefs:  limitStrings(securityTimelineEvidenceRefs(event, nil, nil, phase3IntelligencePayload{}, phase4EnterprisePayload{}), 8),
			Limitations: []string{
				"Self-audit event is a derived summary over canonical audit truth rather than a separate mutable control-plane state.",
			},
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if !items[i].Timestamp.Equal(items[j].Timestamp) {
			return items[i].Timestamp.After(items[j].Timestamp)
		}
		return items[i].EventID > items[j].EventID
	})
	return items, nil
}

func classifySelfAuditEvent(event audit.StoredEvent) (category, severity, action, resourceURI string) {
	switch {
	case strings.HasPrefix(event.EventType, "handoff_"):
		record := parseSecurityTimelineHandoff(event.Handoff)
		if record != nil {
			resourceURI = fmt.Sprintf("/v1/handoff/%s", record.PackageID)
		}
		return "signing_change", "high", strings.ReplaceAll(event.EventType, "_", " "), resourceURI
	case strings.HasPrefix(event.EventType, "federation_"):
		peer := parseSecurityTimelineFederation(event.Federation)
		if peer != nil {
			resourceURI = fmt.Sprintf("/v1/federation/peers/%s", firstNonEmpty(peer.Request.RequestingPeer, peer.Decision.PeerID))
		}
		return "federation_change", "high", strings.ReplaceAll(event.EventType, "_", " "), resourceURI
	case event.EventType == audit.EventTypeValidationHarnessRunRecorded:
		run := parseSelfAuditValidationRun(event.ValidationHarness)
		if run != nil {
			resourceURI = fmt.Sprintf("/v1/validation/executions/%s", run.RunID)
		}
		return "validation_change", "medium", "validation harness run recorded", resourceURI
	case strings.HasPrefix(event.EventType, "hardening_"):
		return "operator_critical_action", "high", strings.ReplaceAll(event.EventType, "_", " "), "/v1/hardening/posture"
	case strings.HasPrefix(event.EventType, "runtime_"):
		return "operator_critical_action", "medium", strings.ReplaceAll(event.EventType, "_", " "), "/v1/runtime/integrity"
	case strings.Contains(event.EventType, "policy") || event.Component == "policy-engine" || event.EventType == audit.EventTypeDeployGateDecision:
		return "policy_change", "medium", strings.ReplaceAll(event.EventType, "_", " "), "/v1/reports/events"
	case strings.HasPrefix(event.EventType, "signing_identity_"):
		return "trust_anchor_change", "high", strings.ReplaceAll(event.EventType, "_", " "), "/v1/signing-identities/status"
	default:
		return "", "", "", ""
	}
}

func parseSelfAuditValidationRun(raw []byte) *validationExecutionRun {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var stored validationHarnessStoredRecord
	if err := json.Unmarshal(raw, &stored); err != nil {
		return nil
	}
	if stored.Bundle.RunID != "" {
		return &stored.Bundle
	}
	return nil
}

func selfAuditGeneratedAt(items []selfAuditEvent) time.Time {
	if len(items) == 0 {
		return time.Unix(0, 0).UTC()
	}
	return items[0].Timestamp.UTC()
}

func limitSelfAuditEvents(items []selfAuditEvent, limit int) []selfAuditEvent {
	if limit <= 0 || len(items) <= limit {
		return items
	}
	return items[:limit]
}

func ternaryStatus(condition bool, okStatus, degradedStatus string) string {
	if condition {
		return okStatus
	}
	return degradedStatus
}

func ternarySummary(condition bool, okValue, degradedValue string) string {
	if condition {
		return okValue
	}
	return degradedValue
}
