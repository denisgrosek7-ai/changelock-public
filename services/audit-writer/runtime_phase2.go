package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	attestationruntime "github.com/denisgrosek/changelock/internal/attestation"
	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	runtimePhase2Component                = "runtime-phase2-manager"
	runtimePhase2SubstrateTruthSchema     = "2.runtime_substrate_truth_list.v1"
	runtimePhase2TrustedProfilesSchema    = "2.runtime_trusted_execution_profiles.v1"
	runtimePhase2AttestationListSchema    = "2.runtime_attestation_verification_list.v1"
	runtimePhase2ResponseSimulationSchema = "2.runtime_response_simulation.v1"
	runtimePhase2RollbackDrillSchema      = "2.runtime_rollback_drill.v1"
	runtimePhase2ProofsSchema             = "2.runtime_phase2_proofs.v1"
	runtimePhase2EventPayloadSchema       = "2.runtime_phase2_event_payload.v1"
)

type runtimePhase2EventPayload struct {
	SchemaVersion      string                                 `json:"schema_version"`
	SubstrateTruth     *runtimesubstrate.SubstrateTruthRecord `json:"substrate_truth,omitempty"`
	ProfileMatch       *runtimesubstrate.ProfileMatch         `json:"profile_match,omitempty"`
	Attestation        *attestationruntime.VerificationResult `json:"attestation,omitempty"`
	ResponseSimulation *phase2ResponseSimulationRecord        `json:"response_simulation,omitempty"`
	RollbackDrill      *phase2RollbackDrillRecord             `json:"rollback_drill,omitempty"`
}

type phase2SubstrateTruthListResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	CurrentState  string                                  `json:"current_state"`
	Items         []runtimesubstrate.SubstrateTruthRecord `json:"items"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type phase2SubstrateTruthEvaluateRequest struct {
	Truth     runtimesubstrate.SubstrateTruthRecord `json:"truth"`
	ProfileID string                                `json:"profile_id,omitempty"`
}

type phase2SubstrateTruthEvaluateResponse struct {
	Status string                                `json:"status"`
	Truth  runtimesubstrate.SubstrateTruthRecord `json:"truth"`
	Match  *runtimesubstrate.ProfileMatch        `json:"match,omitempty"`
}

type phase2TrustedExecutionProfilesResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	CurrentState  string                               `json:"current_state"`
	Profiles      []runtimesubstrate.ExecutionProfile  `json:"profiles"`
	Adapters      []attestationruntime.ProviderAdapter `json:"adapters"`
	Limitations   []string                             `json:"limitations,omitempty"`
}

type phase2AttestationListResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	CurrentState  string                                  `json:"current_state"`
	Items         []attestationruntime.VerificationResult `json:"items"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type phase2AttestationVerifyResponse struct {
	Status       string                                `json:"status"`
	Verification attestationruntime.VerificationResult `json:"verification"`
}

type phase2ResponseSimulationRequest struct {
	FindingID  string `json:"finding_id,omitempty"`
	SubjectRef string `json:"subject_ref,omitempty"`
	Action     string `json:"action,omitempty"`
	Summary    string `json:"summary,omitempty"`
}

type phase2ResponseSimulationRecord struct {
	SchemaVersion         string                     `json:"schema_version"`
	SimulationID          string                     `json:"simulation_id"`
	SubjectRef            string                     `json:"subject_ref"`
	RequestedAction       string                     `json:"requested_action"`
	ProgressiveQuarantine []string                   `json:"progressive_quarantine,omitempty"`
	EvidenceLockState     string                     `json:"evidence_lock_state"`
	Decision              runtimeEnforcementDecision `json:"decision"`
	CurrentState          string                     `json:"current_state"`
	SimulatedAt           time.Time                  `json:"simulated_at"`
	Limitations           []string                   `json:"limitations,omitempty"`
}

type phase2ResponseSimulationResponse struct {
	Status     string                         `json:"status"`
	Simulation phase2ResponseSimulationRecord `json:"simulation"`
}

type phase2RollbackDrillRequest struct {
	FindingID               string `json:"finding_id,omitempty"`
	SubjectRef              string `json:"subject_ref,omitempty"`
	TargetRef               string `json:"target_ref,omitempty"`
	TargetDigest            string `json:"target_digest,omitempty"`
	GitOpsSystem            string `json:"gitops_system,omitempty"`
	TargetVerificationState string `json:"target_verification_state,omitempty"`
	EvidenceLockState       string `json:"evidence_lock_state,omitempty"`
	Summary                 string `json:"summary,omitempty"`
}

type phase2RollbackDrillRecord struct {
	SchemaVersion           string                     `json:"schema_version"`
	DrillID                 string                     `json:"drill_id"`
	SubjectRef              string                     `json:"subject_ref"`
	TargetRef               string                     `json:"target_ref,omitempty"`
	TargetDigest            string                     `json:"target_digest,omitempty"`
	GitOpsSystem            string                     `json:"gitops_system,omitempty"`
	TargetVerificationState string                     `json:"target_verification_state"`
	EvidenceLockState       string                     `json:"evidence_lock_state"`
	Decision                runtimeEnforcementDecision `json:"decision"`
	CurrentState            string                     `json:"current_state"`
	Reasons                 []string                   `json:"reasons,omitempty"`
	SimulatedAt             time.Time                  `json:"simulated_at"`
	Limitations             []string                   `json:"limitations,omitempty"`
}

type phase2RollbackDrillResponse struct {
	Status string                    `json:"status"`
	Drill  phase2RollbackDrillRecord `json:"drill"`
}

type phase2RuntimeProofsResponse struct {
	SchemaVersion        string                                  `json:"schema_version"`
	CurrentState         string                                  `json:"current_state"`
	SubstrateArtifacts   []runtimesubstrate.SubstrateTruthRecord `json:"substrate_artifacts,omitempty"`
	AttestationArtifacts []attestationruntime.VerificationResult `json:"attestation_artifacts,omitempty"`
	ResponseSimulations  []phase2ResponseSimulationRecord        `json:"response_simulations,omitempty"`
	RollbackDrills       []phase2RollbackDrillRecord             `json:"rollback_drills,omitempty"`
	Limitations          []string                                `json:"limitations,omitempty"`
}

func (s server) runtimeSubstrateTruthHandler(w http.ResponseWriter, r *http.Request) {
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
		filter, err := parseRuntimeIntegrityFilter(r)
		if err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase2SubstrateTruth(ctx, filter)
		if err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, phase2SubstrateTruthListResponse{
			SchemaVersion: runtimePhase2SubstrateTruthSchema,
			CurrentState:  map[bool]string{true: "substrate_truth_evidence_active", false: "substrate_truth_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Substrate truth records are bounded execution evidence snapshots tied to canonical audit events; they do not create a parallel runtime database.",
			},
		})
	case http.MethodPost:
		var request phase2SubstrateTruthEvaluateRequest
		if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Truth.SubjectRef = normalizePhase2SubjectRef(request.Truth.SubjectRef)
		request.Truth = runtimesubstrate.NormalizeSubstrateTruthRecord(request.Truth, time.Now)
		if request.Truth.SubjectRef == "" {
			request.Truth.SubjectRef = runtimeSubjectRef(
				request.Truth.Workload.ClusterID,
				request.Truth.Workload.Namespace,
				request.Truth.Workload.WorkloadKind,
				request.Truth.Workload.Workload,
			)
		}
		var match *runtimesubstrate.ProfileMatch
		if profile, ok := runtimesubstrate.ExecutionProfileByID(strings.TrimSpace(request.ProfileID)); ok {
			profileMatch := runtimesubstrate.MatchExecutionProfile(profile, request.Truth)
			match = &profileMatch
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase2Event(ctx, principal.Subject, audit.EventTypeRuntimeSubstrateTruthRecorded, request.Truth, match, nil, nil, nil); err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		httpjson.Write(w, http.StatusCreated, phase2SubstrateTruthEvaluateResponse{
			Status: "recorded",
			Truth:  request.Truth,
			Match:  match,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) runtimeTrustedExecutionProfilesHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, phase2TrustedExecutionProfilesResponse{
		SchemaVersion: runtimePhase2TrustedProfilesSchema,
		CurrentState:  "trusted_execution_profiles_active",
		Profiles:      runtimesubstrate.DefaultExecutionProfiles(),
		Adapters:      attestationruntime.Catalog(),
		Limitations: []string{
			"Trusted execution profiles are bounded contract profiles for substrate-aware admission and runtime validation. They are not a claim of universal confidential provider coverage.",
		},
	})
}

func (s server) runtimeAttestationVerifyHandler(w http.ResponseWriter, r *http.Request) {
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
	var request attestationruntime.VerificationRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request.SubjectRef = normalizePhase2SubjectRef(request.SubjectRef)
	verifier := attestationruntime.NewVerifier()
	result := verifier.Verify(request)
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	if err := s.persistPhase2Event(ctx, principal.Subject, audit.EventTypeRuntimeAttestationVerified, runtimesubstrate.SubstrateTruthRecord{}, nil, &result, nil, nil); err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, phase2AttestationVerifyResponse{
		Status:       "verified",
		Verification: result,
	})
}

func (s server) runtimeAttestationVerificationsHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, err := s.listPhase2Attestations(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, phase2AttestationListResponse{
		SchemaVersion: runtimePhase2AttestationListSchema,
		CurrentState:  map[bool]string{true: "attestation_verifications_active", false: "attestation_verifications_empty"}[len(items) > 0],
		Items:         items,
		Limitations: []string{
			"Attestation verification results remain bounded to structured verifier outputs and explicit measurement trust lists already supplied to the verifier.",
		},
	})
}

func (s server) runtimeResponseSimulationHandler(w http.ResponseWriter, r *http.Request) {
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
	var request phase2ResponseSimulationRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	decision, err := s.evaluateRuntimeEnforcement(ctx, filter, runtimeActionRequest{
		FindingID:  strings.TrimSpace(request.FindingID),
		SubjectRef: normalizePhase2SubjectRef(request.SubjectRef),
		Summary:    strings.TrimSpace(request.Summary),
	}, strings.TrimSpace(request.Action))
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	record := phase2ResponseSimulationRecord{
		SchemaVersion:         runtimePhase2ResponseSimulationSchema,
		SimulationID:          recommendationID("runtime-simulation", decision.SubjectRef, decision.Action),
		SubjectRef:            decision.SubjectRef,
		RequestedAction:       decision.Action,
		ProgressiveQuarantine: phase2ProgressiveQuarantine(decision.Action, decision.ForensicFirst),
		EvidenceLockState:     map[bool]string{true: "evidence_locked_before_destructive_action", false: "bounded_non_destructive_response"}[decision.ForensicFirst],
		Decision:              decision,
		CurrentState:          "simulation_recorded",
		SimulatedAt:           time.Now().UTC(),
		Limitations: []string{
			"Response simulation uses the current bounded runtime enforcement engine and explainability contracts; it does not claim universal kernel-level remediation.",
		},
	}
	if err := s.persistPhase2Event(ctx, principal.Subject, audit.EventTypeRuntimeResponseSimulated, runtimesubstrate.SubstrateTruthRecord{}, nil, nil, &record, nil); err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, phase2ResponseSimulationResponse{
		Status:     "recorded",
		Simulation: record,
	})
}

func (s server) runtimeRollbackDrillHandler(w http.ResponseWriter, r *http.Request) {
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
	var request phase2RollbackDrillRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	decision, err := s.evaluateRuntimeEnforcement(ctx, filter, runtimeActionRequest{
		FindingID:  strings.TrimSpace(request.FindingID),
		SubjectRef: normalizePhase2SubjectRef(request.SubjectRef),
		Summary:    strings.TrimSpace(request.Summary),
	}, runtimeActionRestartTrusted)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	targetVerificationState := firstNonEmpty(strings.TrimSpace(request.TargetVerificationState), "unverified")
	evidenceLockState := firstNonEmpty(strings.TrimSpace(request.EvidenceLockState), "captured")
	currentState := "rollback_drill_failed"
	reasons := []string{}
	if targetVerificationState != "verified" {
		reasons = append(reasons, "rollback_target_not_verified")
	}
	if evidenceLockState != "captured" && evidenceLockState != "sealed" {
		reasons = append(reasons, "evidence_not_locked_before_rollback")
	}
	if decision.Action != runtimeActionRestartTrusted || !decision.RollbackRequired {
		reasons = append(reasons, "rollback_action_not_required")
	}
	if len(reasons) == 0 {
		currentState = "rollback_drill_passed"
		reasons = []string{"trusted_target_verified", "evidence_locked_before_rollback", "bounded_gitops_recovery_ready"}
	}
	record := phase2RollbackDrillRecord{
		SchemaVersion:           runtimePhase2RollbackDrillSchema,
		DrillID:                 recommendationID("runtime-rollback-drill", decision.SubjectRef, firstNonEmpty(request.TargetDigest, request.TargetRef)),
		SubjectRef:              decision.SubjectRef,
		TargetRef:               strings.TrimSpace(request.TargetRef),
		TargetDigest:            strings.TrimSpace(request.TargetDigest),
		GitOpsSystem:            firstNonEmpty(strings.TrimSpace(request.GitOpsSystem), "bounded_gitops"),
		TargetVerificationState: targetVerificationState,
		EvidenceLockState:       evidenceLockState,
		Decision:                decision,
		CurrentState:            currentState,
		Reasons:                 reasons,
		SimulatedAt:             time.Now().UTC(),
		Limitations: []string{
			"Rollback drills validate bounded response semantics, trusted target verification, and evidence-lock ordering. They are not a claim of universal GitOps ownership across every delivery system.",
		},
	}
	if err := s.persistPhase2Event(ctx, principal.Subject, audit.EventTypeRuntimeRollbackDrillRecorded, runtimesubstrate.SubstrateTruthRecord{}, nil, nil, nil, &record); err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, phase2RollbackDrillResponse{
		Status: "recorded",
		Drill:  record,
	})
}

func (s server) runtimePhase2ProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	substrate, err := s.listPhase2SubstrateTruth(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	attestations, err := s.listPhase2Attestations(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	simulations, err := s.listPhase2ResponseSimulations(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	drills, err := s.listPhase2RollbackDrills(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	currentState := "phase2_core_incomplete"
	if len(substrate) > 0 && hasVerifiedAttestation(attestations) && len(simulations) > 0 && hasPassingRollbackDrill(drills) {
		currentState = "phase2_core_slice_active"
	}
	httpjson.Write(w, http.StatusOK, phase2RuntimeProofsResponse{
		SchemaVersion:        runtimePhase2ProofsSchema,
		CurrentState:         currentState,
		SubstrateArtifacts:   takeSubstrateArtifacts(substrate, 5),
		AttestationArtifacts: takeAttestationArtifacts(attestations, 5),
		ResponseSimulations:  takeResponseSimulationArtifacts(simulations, 5),
		RollbackDrills:       takeRollbackDrillArtifacts(drills, 5),
		Limitations: []string{
			"Phase 2 proofs expose bounded artifacts for substrate truth, attestation verification, response simulation, and rollback drill evidence.",
		},
	})
}

func (s server) persistPhase2Event(ctx context.Context, actor, eventType string, truth runtimesubstrate.SubstrateTruthRecord, match *runtimesubstrate.ProfileMatch, attestation *attestationruntime.VerificationResult, simulation *phase2ResponseSimulationRecord, drill *phase2RollbackDrillRecord) error {
	payload, err := canonicalJSON(runtimePhase2EventPayload{
		SchemaVersion:      runtimePhase2EventPayloadSchema,
		SubstrateTruth:     optionalSubstrateTruth(truth),
		ProfileMatch:       match,
		Attestation:        attestation,
		ResponseSimulation: simulation,
		RollbackDrill:      drill,
	})
	if err != nil {
		return err
	}
	var (
		clusterID    string
		namespace    string
		workloadKind string
		workload     string
		digest       string
		subjectRef   string
		decision     = audit.DecisionAllow
		reasons      []string
	)
	if match != nil {
		subjectRef = match.SubjectRef
		if !match.Allowed {
			decision = audit.DecisionDeny
		}
		reasons = append(reasons, match.CurrentState)
		reasons = append(reasons, match.Reasons...)
	}
	if truth.SubjectRef != "" {
		subjectRef = truth.SubjectRef
		digest = truth.Workload.ImageDigest
	}
	if attestation != nil {
		subjectRef = firstNonEmpty(subjectRef, attestation.SubjectRef)
		reasons = append(reasons, attestation.CurrentState)
		reasons = append(reasons, attestation.Reasons...)
		if attestation.CurrentState == attestationruntime.VerdictMismatch || attestation.CurrentState == attestationruntime.VerdictRejected {
			decision = audit.DecisionDeny
		}
	}
	if simulation != nil {
		subjectRef = firstNonEmpty(subjectRef, simulation.SubjectRef)
		reasons = append(reasons, simulation.RequestedAction, simulation.CurrentState)
	}
	if drill != nil {
		subjectRef = firstNonEmpty(subjectRef, drill.SubjectRef)
		reasons = append(reasons, drill.CurrentState)
		reasons = append(reasons, drill.Reasons...)
		if drill.CurrentState != "rollback_drill_passed" {
			decision = audit.DecisionDeny
		}
	}
	if subjectRef != "" {
		clusterID, namespace, workloadKind, workload, _ = parseRuntimeSubjectRef(subjectRef)
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:        audit.NewRequestID(),
		Component:        runtimePhase2Component,
		EventType:        eventType,
		Actor:            strings.TrimSpace(actor),
		ClusterID:        clusterID,
		TenantID:         audit.TenantFromNamespace(namespace),
		Environment:      audit.EnvironmentFromNamespace(namespace),
		Namespace:        namespace,
		WorkloadKind:     workloadKind,
		Workload:         workload,
		Digest:           digest,
		Decision:         decision,
		Reasons:          uniqueStrings(reasons),
		RuntimeIntegrity: payload,
	})
	return err
}

func (s server) listPhase2SubstrateTruth(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimesubstrate.SubstrateTruthRecord, error) {
	events, err := s.listPhase2Events(ctx, filter, audit.EventTypeRuntimeSubstrateTruthRecorded)
	if err != nil {
		return nil, err
	}
	items := []runtimesubstrate.SubstrateTruthRecord{}
	for _, item := range events {
		payload := parseRuntimePhase2Payload(item.RuntimeIntegrity)
		if payload.SubstrateTruth != nil {
			items = append(items, *payload.SubstrateTruth)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ObservedAt.After(items[j].ObservedAt) })
	return items, nil
}

func (s server) listPhase2Attestations(ctx context.Context, filter runtimeIntegrityFilter) ([]attestationruntime.VerificationResult, error) {
	events, err := s.listPhase2Events(ctx, filter, audit.EventTypeRuntimeAttestationVerified)
	if err != nil {
		return nil, err
	}
	items := []attestationruntime.VerificationResult{}
	for _, item := range events {
		payload := parseRuntimePhase2Payload(item.RuntimeIntegrity)
		if payload.Attestation != nil {
			items = append(items, *payload.Attestation)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].VerifiedAt.After(items[j].VerifiedAt) })
	return items, nil
}

func (s server) listPhase2ResponseSimulations(ctx context.Context, filter runtimeIntegrityFilter) ([]phase2ResponseSimulationRecord, error) {
	events, err := s.listPhase2Events(ctx, filter, audit.EventTypeRuntimeResponseSimulated)
	if err != nil {
		return nil, err
	}
	items := []phase2ResponseSimulationRecord{}
	for _, item := range events {
		payload := parseRuntimePhase2Payload(item.RuntimeIntegrity)
		if payload.ResponseSimulation != nil {
			items = append(items, *payload.ResponseSimulation)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].SimulatedAt.After(items[j].SimulatedAt) })
	return items, nil
}

func (s server) listPhase2RollbackDrills(ctx context.Context, filter runtimeIntegrityFilter) ([]phase2RollbackDrillRecord, error) {
	events, err := s.listPhase2Events(ctx, filter, audit.EventTypeRuntimeRollbackDrillRecorded)
	if err != nil {
		return nil, err
	}
	items := []phase2RollbackDrillRecord{}
	for _, item := range events {
		payload := parseRuntimePhase2Payload(item.RuntimeIntegrity)
		if payload.RollbackDrill != nil {
			items = append(items, *payload.RollbackDrill)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].SimulatedAt.After(items[j].SimulatedAt) })
	return items, nil
}

func (s server) listPhase2Events(ctx context.Context, filter runtimeIntegrityFilter, eventType string) ([]audit.StoredEvent, error) {
	return s.store.ListEvents(ctx, audit.EventFilter{
		EventType:   eventType,
		Component:   runtimePhase2Component,
		ClusterID:   filter.ClusterID,
		Environment: filter.Environment,
		TenantID:    filter.TenantID,
		Repo:        filter.Repo,
		Limit:       max(filter.Limit, 100),
	})
}

func parseRuntimePhase2Payload(value json.RawMessage) runtimePhase2EventPayload {
	if len(value) == 0 || string(value) == "null" {
		return runtimePhase2EventPayload{}
	}
	var payload runtimePhase2EventPayload
	if err := json.Unmarshal(value, &payload); err != nil {
		return runtimePhase2EventPayload{}
	}
	return payload
}

func phase2ProgressiveQuarantine(action string, forensicFirst bool) []string {
	steps := []string{"observe", "traffic_restriction"}
	if forensicFirst {
		steps = append(steps, "forensic_hold", "sealed_snapshot")
	}
	switch strings.TrimSpace(action) {
	case runtimeActionApplyNetworkIsolation:
		steps = append(steps, "network_isolation", "capability_reduction")
	case runtimeActionRestartTrusted:
		steps = append(steps, "network_isolation", "bounded_restart", "verified_rollback_target")
	default:
		steps = append(steps, "bounded_quarantine")
	}
	return uniqueStrings(steps)
}

func optionalSubstrateTruth(truth runtimesubstrate.SubstrateTruthRecord) *runtimesubstrate.SubstrateTruthRecord {
	if strings.TrimSpace(truth.SubjectRef) == "" {
		return nil
	}
	return &truth
}

func hasVerifiedAttestation(items []attestationruntime.VerificationResult) bool {
	for _, item := range items {
		if item.CurrentState == attestationruntime.VerdictVerified {
			return true
		}
	}
	return false
}

func hasPassingRollbackDrill(items []phase2RollbackDrillRecord) bool {
	for _, item := range items {
		if item.CurrentState == "rollback_drill_passed" {
			return true
		}
	}
	return false
}

func takeSubstrateArtifacts(items []runtimesubstrate.SubstrateTruthRecord, limit int) []runtimesubstrate.SubstrateTruthRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takeAttestationArtifacts(items []attestationruntime.VerificationResult, limit int) []attestationruntime.VerificationResult {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takeResponseSimulationArtifacts(items []phase2ResponseSimulationRecord, limit int) []phase2ResponseSimulationRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takeRollbackDrillArtifacts(items []phase2RollbackDrillRecord, limit int) []phase2RollbackDrillRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func normalizePhase2SubjectRef(value string) string {
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
