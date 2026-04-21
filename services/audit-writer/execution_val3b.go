package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	executionCoverageSchemaVersion   = "3b.execution_coverage.v1"
	executionCoverageMatrixSchema    = "3b.execution_coverage_matrix.v1"
	executionVMLineageSchemaVersion  = "3b.vm_lineage.v1"
	executionEphemeralSchemaVersion  = "3b.ephemeral_execution.v1"
	executionSubstrateVMWorkload     = "vm_workload"
	executionSubstrateRuntimeJob     = "runtime_job"
	executionSubstrateValidationExec = "validation_execution"
)

type executionCoverageScope struct {
	ClusterID    string `json:"cluster_id,omitempty"`
	TenantID     string `json:"tenant_id,omitempty"`
	Environment  string `json:"environment,omitempty"`
	Repo         string `json:"repo,omitempty"`
	Namespace    string `json:"namespace,omitempty"`
	WorkloadKind string `json:"workload_kind,omitempty"`
	Workload     string `json:"workload,omitempty"`
}

type executionHandoffCoverageItem struct {
	PackageID              string    `json:"package_id"`
	PackageType            string    `json:"package_type,omitempty"`
	CreatedAt              time.Time `json:"created_at"`
	ScopeSummary           string    `json:"scope_summary,omitempty"`
	VerificationStatus     string    `json:"verification_status,omitempty"`
	VerificationURI        string    `json:"verification_uri,omitempty"`
	OfflineVerifierPresent bool      `json:"offline_verifier_present"`
}

type executionHandoffCoverage struct {
	OfflineVerificationSupported bool                           `json:"offline_verification_supported"`
	Items                        []executionHandoffCoverageItem `json:"items,omitempty"`
	Limitations                  []string                       `json:"limitations,omitempty"`
}

type executionFederationCoverage struct {
	OfflineModeSupported bool                     `json:"offline_mode_supported"`
	SyncStatus           string                   `json:"sync_status,omitempty"`
	TrustHealth          string                   `json:"trust_health,omitempty"`
	DelayedSyncRequired  bool                     `json:"delayed_sync_required"`
	LocalTrustAnchors    []federationAnchorRecord `json:"local_trust_anchors,omitempty"`
	StalePeers           []string                 `json:"stale_peers,omitempty"`
	PolicyDivergence     []string                 `json:"policy_divergence,omitempty"`
	Limitations          []string                 `json:"limitations,omitempty"`
}

type executionValidationCoverage struct {
	OfflineModeSupported  bool     `json:"offline_mode_supported"`
	ScenarioCatalogSize   int      `json:"scenario_catalog_size"`
	RecentExecutionRefs   []string `json:"recent_execution_refs,omitempty"`
	RecentCertificateRefs []string `json:"recent_certificate_refs,omitempty"`
	Limitations           []string `json:"limitations,omitempty"`
}

type executionCoverageCapability struct {
	CapabilityID     string   `json:"capability_id"`
	DisplayName      string   `json:"display_name"`
	Supported        bool     `json:"supported"`
	CurrentState     string   `json:"current_state"`
	EvidenceSurfaces []string `json:"evidence_surfaces,omitempty"`
	ContractSummary  string   `json:"contract_summary"`
	Limitations      []string `json:"limitations,omitempty"`
}

type executionDelayedSyncSemantics struct {
	Required               bool     `json:"required"`
	SyncStatus             string   `json:"sync_status"`
	PolicyPropagationModel string   `json:"policy_propagation_model"`
	SafeReadSemantics      []string `json:"safe_read_semantics,omitempty"`
	BlockedClaims          []string `json:"blocked_claims,omitempty"`
	RecoverySignals        []string `json:"recovery_signals,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type executionDegradedMode struct {
	Active           bool     `json:"active"`
	Causes           []string `json:"causes,omitempty"`
	EvidenceModel    string   `json:"evidence_model"`
	EvidenceSurfaces []string `json:"evidence_surfaces,omitempty"`
	SummarySemantics string   `json:"summary_semantics"`
}

type executionCoverageResponse struct {
	SchemaVersion                   string                        `json:"schema_version"`
	Scope                           executionCoverageScope        `json:"scope"`
	DisconnectedHandoffVerification executionHandoffCoverage      `json:"disconnected_handoff_verification"`
	OfflineFederation               executionFederationCoverage   `json:"offline_federation"`
	OfflineValidation               executionValidationCoverage   `json:"offline_validation"`
	CapabilityMatrix                []executionCoverageCapability `json:"capability_matrix,omitempty"`
	DelayedSyncSemantics            executionDelayedSyncSemantics `json:"delayed_sync_semantics"`
	DegradedMode                    executionDegradedMode         `json:"degraded_mode"`
	Limitations                     []string                      `json:"limitations,omitempty"`
}

type executionCoverageMatrixResponse struct {
	SchemaVersion        string                        `json:"schema_version"`
	Scope                executionCoverageScope        `json:"scope"`
	CapabilityMatrix     []executionCoverageCapability `json:"capability_matrix,omitempty"`
	DelayedSyncSemantics executionDelayedSyncSemantics `json:"delayed_sync_semantics"`
	DegradedMode         executionDegradedMode         `json:"degraded_mode"`
	Limitations          []string                      `json:"limitations,omitempty"`
}

type executionVMPolicyEvidenceParity struct {
	PolicyParity   []string `json:"policy_parity,omitempty"`
	EvidenceParity []string `json:"evidence_parity,omitempty"`
	PostureParity  []string `json:"posture_parity,omitempty"`
	Limitations    []string `json:"limitations,omitempty"`
}

type executionVMLineageItem struct {
	SubjectRef               string                `json:"subject_ref"`
	Cluster                  string                `json:"cluster,omitempty"`
	Environment              string                `json:"environment,omitempty"`
	Namespace                string                `json:"namespace,omitempty"`
	WorkloadKind             string                `json:"workload_kind"`
	Workload                 string                `json:"workload"`
	ServiceAccount           string                `json:"service_account,omitempty"`
	ImageDigest              string                `json:"image_digest,omitempty"`
	DesiredStateSourceRef    string                `json:"desired_state_source_ref,omitempty"`
	DesiredStateApprovalID   string                `json:"desired_state_approval_id,omitempty"`
	DesiredStateVerification string                `json:"desired_state_verification_state,omitempty"`
	ApprovedDigest           string                `json:"approved_digest,omitempty"`
	RunningDigest            string                `json:"running_digest,omitempty"`
	TrustInputs              []string              `json:"trust_inputs,omitempty"`
	ExpectedSigners          []string              `json:"expected_signers,omitempty"`
	RuntimeState             runtimeIntegrityState `json:"runtime_state"`
	RuntimePosture           runtimePostureState   `json:"runtime_posture"`
	ValidationRefs           []string              `json:"validation_refs,omitempty"`
	ReadbackRefs             []advisoryReadbackRef `json:"readback_refs,omitempty"`
	EvidenceRefs             []string              `json:"evidence_refs,omitempty"`
	Limitations              []string              `json:"limitations,omitempty"`
}

type executionVMLineageResponse struct {
	SchemaVersion        string                          `json:"schema_version"`
	PolicyEvidenceParity executionVMPolicyEvidenceParity `json:"policy_evidence_parity"`
	Items                []executionVMLineageItem        `json:"items"`
	Limitations          []string                        `json:"limitations,omitempty"`
}

type executionEphemeralItem struct {
	ExecutionRef      string                `json:"execution_ref"`
	SubstrateType     string                `json:"substrate_type"`
	SubjectRef        string                `json:"subject_ref,omitempty"`
	WorkloadKind      string                `json:"workload_kind,omitempty"`
	Workload          string                `json:"workload,omitempty"`
	Environment       string                `json:"environment,omitempty"`
	Namespace         string                `json:"namespace,omitempty"`
	Mode              string                `json:"mode,omitempty"`
	IsolationClass    string                `json:"isolation_class,omitempty"`
	Status            string                `json:"status"`
	SnapshotSemantics string                `json:"snapshot_semantics"`
	RetentionModel    string                `json:"retention_model"`
	CertificateRef    string                `json:"certificate_ref,omitempty"`
	StartedAt         time.Time             `json:"started_at"`
	CompletedAt       time.Time             `json:"completed_at"`
	EvidenceRefs      []string              `json:"evidence_refs,omitempty"`
	ReadbackRefs      []advisoryReadbackRef `json:"readback_refs,omitempty"`
	Limitations       []string              `json:"limitations,omitempty"`
}

type executionEphemeralRetentionContract struct {
	SnapshotSemantics   []string `json:"snapshot_semantics,omitempty"`
	RetentionSemantics  []string `json:"retention_semantics,omitempty"`
	SummarySemantics    string   `json:"summary_semantics"`
	CorrelationBehavior string   `json:"correlation_behavior"`
	Limitations         []string `json:"limitations,omitempty"`
}

type executionEphemeralResponse struct {
	SchemaVersion     string                              `json:"schema_version"`
	RetentionContract executionEphemeralRetentionContract `json:"retention_contract"`
	Items             []executionEphemeralItem            `json:"items"`
	Limitations       []string                            `json:"limitations,omitempty"`
}

func (s server) executionCoverageHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseExecutionCoverageFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildExecutionCoverageSummary(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) executionCoverageMatrixHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseExecutionCoverageFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildExecutionCoverageSummary(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, executionCoverageMatrixResponse{
		SchemaVersion:        executionCoverageMatrixSchema,
		Scope:                response.Scope,
		CapabilityMatrix:     response.CapabilityMatrix,
		DelayedSyncSemantics: response.DelayedSyncSemantics,
		DegradedMode:         response.DegradedMode,
		Limitations: uniqueStrings(append([]string{
			"Execution coverage matrix isolates offline, delayed-sync, and degraded-mode semantics without introducing a new execution truth layer.",
		}, response.Limitations...)),
	})
}

func (s server) executionVMLineageHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseExecutionCoverageFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildExecutionVMLineage(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, executionVMLineageResponse{
		SchemaVersion:        executionVMLineageSchemaVersion,
		PolicyEvidenceParity: executionVMPolicyEvidenceParityContract(),
		Items:                items,
		Limitations:          limitations,
	})
}

func (s server) executionEphemeralHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseExecutionCoverageFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildExecutionEphemeralCoverage(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, executionEphemeralResponse{
		SchemaVersion:     executionEphemeralSchemaVersion,
		RetentionContract: executionEphemeralRetentionContractModel(),
		Items:             items,
		Limitations:       limitations,
	})
}

func (s server) buildExecutionCoverageSummary(ctx context.Context, filter runtimeIntegrityFilter) (executionCoverageResponse, error) {
	handoffs, err := s.listStoredHandoffs(ctx, filter)
	if err != nil {
		return executionCoverageResponse{}, err
	}
	federationView, err := s.buildFederationGlobalView(ctx)
	if err != nil {
		return executionCoverageResponse{}, err
	}
	validationRuns, _, err := s.listStrictValidationRuns(ctx, validationFilterFromRuntimeFilter(filter, filter.Workload))
	if err != nil {
		return executionCoverageResponse{}, err
	}

	handoffItems := make([]executionHandoffCoverageItem, 0, len(handoffs))
	offlineVerifySupported := false
	for _, record := range handoffs {
		if record.Bundle.OfflineVerifierPresent {
			offlineVerifySupported = true
		}
		handoffItems = append(handoffItems, executionHandoffCoverageItem{
			PackageID:              record.PackageID,
			PackageType:            record.PackageType,
			CreatedAt:              record.Manifest.CreatedAt,
			ScopeSummary:           record.Session.ScopeSummary,
			VerificationStatus:     record.Verification.OverallStatus,
			VerificationURI:        record.VerificationURI,
			OfflineVerifierPresent: record.Bundle.OfflineVerifierPresent,
		})
	}
	if len(handoffItems) > filter.Limit {
		handoffItems = handoffItems[:filter.Limit]
	}

	localAnchors := []federationAnchorRecord{}
	for _, item := range federationView.Anchors {
		if item.PeerID == federationLocalPeerID {
			localAnchors = append(localAnchors, item)
		}
	}

	recentExecutionRefs := []string{}
	recentCertificateRefs := []string{}
	for _, run := range validationRuns {
		for _, execution := range run.Executions {
			recentExecutionRefs = append(recentExecutionRefs, "/v1/validation/executions/"+execution.ExecutionID)
		}
		if run.Certificate.CertificateID != "" {
			recentCertificateRefs = append(recentCertificateRefs, "/v1/validation/certificates/"+run.Certificate.CertificateID)
		}
		if len(recentExecutionRefs) >= filter.Limit && len(recentCertificateRefs) >= filter.Limit {
			break
		}
	}
	recentExecutionRefs = uniqueStrings(recentExecutionRefs)
	recentCertificateRefs = uniqueStrings(recentCertificateRefs)
	if len(recentExecutionRefs) > filter.Limit {
		recentExecutionRefs = recentExecutionRefs[:filter.Limit]
	}
	if len(recentCertificateRefs) > filter.Limit {
		recentCertificateRefs = recentCertificateRefs[:filter.Limit]
	}

	delayedSyncRequired := federationView.PolicyState.SyncStatus != federationSyncStatusSynced &&
		federationView.PolicyState.SyncStatus != federationSyncStatusSyncedWithOverrides
	degradedCauses := []string{}
	if len(federationView.StalePeers) > 0 {
		degradedCauses = append(degradedCauses, "stale_federation_peers_present")
	}
	if federationView.PolicyState.SyncStatus == federationSyncStatusDiverged {
		degradedCauses = append(degradedCauses, "policy_sync_diverged")
	}
	if federationView.TrustHealth != "healthy" && len(degradedCauses) == 0 {
		degradedCauses = append(degradedCauses, "federation_trust_health_degraded")
	}

	return executionCoverageResponse{
		SchemaVersion: executionCoverageSchemaVersion,
		Scope: executionCoverageScope{
			ClusterID:    filter.ClusterID,
			TenantID:     filter.TenantID,
			Environment:  filter.Environment,
			Repo:         filter.Repo,
			Namespace:    filter.Namespace,
			WorkloadKind: filter.WorkloadKind,
			Workload:     filter.Workload,
		},
		DisconnectedHandoffVerification: executionHandoffCoverage{
			OfflineVerificationSupported: offlineVerifySupported,
			Items:                        handoffItems,
			Limitations: []string{
				"Disconnected handoff verification derives from sealed bundle metadata and verification lineage already persisted in ChangeLock; it does not claim live reachability to a remote exchange partner.",
			},
		},
		OfflineFederation: executionFederationCoverage{
			OfflineModeSupported: len(localAnchors) > 0,
			SyncStatus:           federationView.PolicyState.SyncStatus,
			TrustHealth:          federationView.TrustHealth,
			DelayedSyncRequired:  delayedSyncRequired,
			LocalTrustAnchors:    localAnchors,
			StalePeers:           append([]string(nil), federationView.StalePeers...),
			PolicyDivergence:     append([]string(nil), federationView.PolicyDivergence...),
			Limitations: uniqueStrings(append([]string{
				"Offline federation remains bounded by local trust anchors, last verified peer proofs, and local policy admissibility; remote proof never bypasses local trust policy.",
			}, federationView.Limitations...)),
		},
		OfflineValidation: executionValidationCoverage{
			OfflineModeSupported:  len(validationScenarioCatalog()) > 0,
			ScenarioCatalogSize:   len(validationScenarioCatalog()),
			RecentExecutionRefs:   recentExecutionRefs,
			RecentCertificateRefs: recentCertificateRefs,
			Limitations: []string{
				"Offline validation coverage describes locally executable strict validation harness semantics and persisted certificate lineage; it does not imply remote substrate emulation beyond the evidence already in scope.",
			},
		},
		CapabilityMatrix: buildExecutionCoverageCapabilityMatrix(
			offlineVerifySupported,
			len(localAnchors) > 0,
			federationView.PolicyState.SyncStatus,
			delayedSyncRequired,
			len(validationScenarioCatalog()) > 0,
			len(recentExecutionRefs) > 0 || len(recentCertificateRefs) > 0,
			len(degradedCauses) > 0,
		),
		DelayedSyncSemantics: buildExecutionDelayedSyncSemantics(federationView.PolicyState.SyncStatus, delayedSyncRequired),
		DegradedMode: executionDegradedMode{
			Active:           len(degradedCauses) > 0,
			Causes:           degradedCauses,
			EvidenceModel:    "Degraded mode is evidence-backed and surfaces local anchors, stale peers, delayed policy sync, and persisted validation or handoff artifacts instead of hiding limitation state.",
			EvidenceSurfaces: []string{"federation.local_trust_anchors", "federation.stale_peers", "federation.policy_state.sync_status", "handoff.verification_uri", "validation.execution_refs", "validation.certificate_refs"},
			SummarySemantics: "Degraded execution posture preserves a bounded operator summary over local trust anchors, delayed sync state, stale-peer evidence, and persisted validation or handoff references; it does not silently pretend to be fully connected.",
		},
		Limitations: []string{
			"Execution coverage 3B expands disconnected, VM, and ephemeral visibility by composing existing handoff, federation, validation, and runtime evidence; it does not introduce a new execution truth store.",
		},
	}, nil
}

func (s server) buildExecutionVMLineage(ctx context.Context, filter runtimeIntegrityFilter) ([]executionVMLineageItem, []string, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	items := []executionVMLineageItem{}
	for _, subject := range snapshot.sortedSubjects() {
		if !isVMWorkloadKind(subject.WorkloadKind) {
			continue
		}
		hydrated, err := s.hydrateExecutionSubjectStates(ctx, filter, subject)
		if err != nil {
			return nil, nil, err
		}
		profile, err := s.profileFromSubject(ctx, filter, hydrated)
		if err != nil {
			return nil, nil, err
		}
		findings, err := s.findingsForSubject(ctx, filter, hydrated, profile)
		if err != nil {
			return nil, nil, err
		}
		sbom := s.subjectSBOMVerification(hydrated)
		sandbox := s.buildRuntimeSandboxDecision(hydrated, findings, sbom)
		state := s.buildRuntimeIntegrityState(hydrated, findings, sandbox, sbom)
		posture := buildRuntimePostureState(hydrated, findings, sbom, sandbox, state)
		readbackRefs, err := s.runtimeReadbackRefs(ctx, filter, hydrated)
		if err != nil {
			return nil, nil, err
		}
		validationRefs, err := s.executionValidationRefsForWorkload(ctx, filter, hydrated.Workload)
		if err != nil {
			return nil, nil, err
		}
		items = append(items, executionVMLineageItem{
			SubjectRef:               hydrated.SubjectRef,
			Cluster:                  hydrated.Cluster,
			Environment:              hydrated.Environment,
			Namespace:                hydrated.Namespace,
			WorkloadKind:             hydrated.WorkloadKind,
			Workload:                 hydrated.Workload,
			ServiceAccount:           hydrated.ServiceAccount,
			ImageDigest:              hydrated.ImageDigest,
			DesiredStateSourceRef:    firstNonEmpty(runtimeDesiredStateSource(hydrated), ""),
			DesiredStateApprovalID:   firstNonEmpty(runtimeDesiredStateApproval(hydrated), ""),
			DesiredStateVerification: runtimeFindingDesiredStateVerification(hydrated),
			ApprovedDigest:           runtimeApprovedDigest(hydrated),
			RunningDigest:            runtimeObservedDigest(hydrated),
			TrustInputs:              uniqueStrings(mapKeys(hydrated.TrustInputs)),
			ExpectedSigners:          uniqueStrings(hydrated.ExpectedSigners),
			RuntimeState:             state,
			RuntimePosture:           posture,
			ValidationRefs:           validationRefs,
			ReadbackRefs:             readbackRefs,
			EvidenceRefs:             uniqueStrings(mapKeys(hydrated.EvidenceRefs)),
			Limitations: []string{
				"VM lineage is workload-scoped and preserves policy, evidence, runtime posture, and validation linkage parity where the corresponding VM evidence exists in ChangeLock.",
			},
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].RuntimeState.LastVerifiedAt.Equal(items[j].RuntimeState.LastVerifiedAt) {
			return items[i].SubjectRef < items[j].SubjectRef
		}
		return items[i].RuntimeState.LastVerifiedAt.After(items[j].RuntimeState.LastVerifiedAt)
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, []string{
		"VM lineage is derived from the same evidence-native runtime model used for container workloads and does not collapse VM support into a separate truth layer.",
		"Validation linkage is advisory and only appears when matching strict validation evidence exists for the VM workload scope.",
	}, nil
}

func (s server) buildExecutionEphemeralCoverage(ctx context.Context, filter runtimeIntegrityFilter) ([]executionEphemeralItem, []string, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	items := []executionEphemeralItem{}
	for _, subject := range snapshot.sortedSubjects() {
		if !isEphemeralWorkloadKind(subject.WorkloadKind) {
			continue
		}
		hydrated, err := s.hydrateExecutionSubjectStates(ctx, filter, subject)
		if err != nil {
			return nil, nil, err
		}
		readbackRefs, err := s.runtimeReadbackRefs(ctx, filter, hydrated)
		if err != nil {
			return nil, nil, err
		}
		startedAt, completedAt := ephemeralRuntimeWindow(hydrated)
		items = append(items, executionEphemeralItem{
			ExecutionRef:      hydrated.SubjectRef,
			SubstrateType:     executionSubstrateRuntimeJob,
			SubjectRef:        hydrated.SubjectRef,
			WorkloadKind:      hydrated.WorkloadKind,
			Workload:          hydrated.Workload,
			Environment:       hydrated.Environment,
			Namespace:         hydrated.Namespace,
			Status:            ephemeralRuntimeStatus(hydrated),
			SnapshotSemantics: "Bounded runtime lineage for a short-lived workload is assembled from desired-state, active-state, and linked runtime evidence snapshots rather than long-lived process residency.",
			RetentionModel:    "Short-lived runtime workloads retain summary evidence, subject lineage, and linked readback references even when the process itself no longer exists.",
			StartedAt:         startedAt,
			CompletedAt:       completedAt,
			EvidenceRefs:      uniqueStrings(mapKeys(hydrated.EvidenceRefs)),
			ReadbackRefs:      readbackRefs,
			Limitations: []string{
				"Ephemeral runtime coverage is bounded to the evidence captured during the observed workload window and does not imply continuous process residency after the job exits.",
			},
		})
	}

	runs, _, err := s.listStrictValidationRuns(ctx, validationFilterFromRuntimeFilter(filter, filter.Workload))
	if err != nil {
		return nil, nil, err
	}
	for _, run := range runs {
		verdictByExecution := map[string]validationVerdict{}
		for _, verdict := range run.Verdicts {
			verdictByExecution[verdict.ExecutionID] = verdict
		}
		for _, execution := range run.Executions {
			verdict := verdictByExecution[execution.ExecutionID]
			items = append(items, executionEphemeralItem{
				ExecutionRef:      "/v1/validation/executions/" + execution.ExecutionID,
				SubstrateType:     executionSubstrateValidationExec,
				Environment:       execution.Environment,
				Namespace:         execution.Namespace,
				Mode:              execution.Mode,
				IsolationClass:    execution.IsolationClass,
				Status:            execution.Status,
				SnapshotSemantics: "Strict validation executions remain isolated and are retained as execution, verdict, and certificate snapshots instead of becoming production runtime truth.",
				RetentionModel:    "Validation executions retain bounded execution evidence, verdict lineage, and a seal-ready certificate reference for short-lived harness runs.",
				CertificateRef:    "/v1/validation/certificates/" + run.Certificate.CertificateID,
				StartedAt:         execution.StartedAt,
				CompletedAt:       execution.CompletedAt,
				EvidenceRefs: uniqueStrings(append(
					append([]string{}, execution.EvidenceRefs...),
					append(verdict.EvidenceRefs, run.Certificate.EvidenceRefs...)...,
				)),
				ReadbackRefs: verdict.ObservedOutcome.ReadbackRefs,
				Limitations: uniqueStrings(append([]string{
					"Validation execution evidence is isolated by design and never claims production mutation for short-lived test runs.",
				}, execution.Limitations...)),
			})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].CompletedAt.Equal(items[j].CompletedAt) {
			return items[i].ExecutionRef < items[j].ExecutionRef
		}
		return items[i].CompletedAt.After(items[j].CompletedAt)
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, []string{
		"Ephemeral execution coverage combines short-lived runtime workload lineage with strict validation execution evidence so evidence does not disappear when the execution unit is no longer resident.",
		"Retention is summary-first and bounded; the API preserves lineage and certificate references rather than claiming full process replay for every ephemeral execution.",
	}, nil
}

func buildExecutionCoverageCapabilityMatrix(offlineVerifySupported, offlineFederationSupported bool, syncStatus string, delayedSyncRequired, offlineValidationSupported, validationEvidencePresent, degradedActive bool) []executionCoverageCapability {
	federationState := "metadata_only"
	if offlineFederationSupported {
		federationState = "ready"
	}
	if degradedActive {
		federationState = "degraded"
	}

	validationState := "unsupported"
	switch {
	case offlineValidationSupported && validationEvidencePresent:
		validationState = "ready"
	case offlineValidationSupported:
		validationState = "catalog_only"
	}

	items := []executionCoverageCapability{
		{
			CapabilityID:     "disconnected_handoff_verification",
			DisplayName:      "Disconnected handoff verification",
			Supported:        offlineVerifySupported,
			CurrentState:     mapBoolState(offlineVerifySupported, "ready", "metadata_only"),
			EvidenceSurfaces: []string{"handoff.bundle.offline_verifier_present", "handoff.verification_uri", "handoff.manifest.scope"},
			ContractSummary:  "Sealed handoff bundles can be verified from persisted bundle metadata and verification lineage without claiming live remote reachability.",
			Limitations: []string{
				"Disconnected handoff verification proves bundle verification posture, not real-time partner availability.",
			},
		},
		{
			CapabilityID:     "offline_federation_local_trust",
			DisplayName:      "Offline federation local trust",
			Supported:        offlineFederationSupported,
			CurrentState:     federationState,
			EvidenceSurfaces: []string{"federation.local_trust_anchors", "federation.policy_state.sync_status", "federation.policy_divergence", "federation.stale_peers"},
			ContractSummary:  "Air-gapped and disconnected federation remains bounded by local trust anchors, last verified peer proofs, and local admissibility rules.",
			Limitations: []string{
				"Remote proof never bypasses local policy, and stale or diverged peers remain visible instead of being hidden behind synthetic health.",
			},
		},
		{
			CapabilityID:     "delayed_sync_policy_propagation",
			DisplayName:      "Delayed sync policy propagation",
			Supported:        offlineFederationSupported,
			CurrentState:     mapBoolState(delayedSyncRequired, "required", "not_required"),
			EvidenceSurfaces: []string{"federation.policy_state.sync_status", "federation.stale_peers", "federation.policy_divergence"},
			ContractSummary:  "Federated policy propagation is bounded and safety-gated; delayed sync preserves local admissibility while exposing sync lag and divergence.",
			Limitations: []string{
				"Delayed sync never upgrades a stale peer into an implicitly current source of truth.",
				"Current sync state is bounded to " + firstNonEmpty(syncStatus, federationSyncStatusLocalOnly) + ".",
			},
		},
		{
			CapabilityID:     "offline_validation_execution",
			DisplayName:      "Offline validation execution",
			Supported:        offlineValidationSupported,
			CurrentState:     validationState,
			EvidenceSurfaces: []string{"validation.scenario_catalog", "validation.execution_refs", "validation.certificate_refs"},
			ContractSummary:  "Strict validation remains locally executable with persisted execution and certificate lineage even when the environment is disconnected.",
			Limitations: []string{
				"Offline validation coverage is bounded to locally available scenarios and persisted outputs; it does not claim arbitrary remote substrate emulation.",
			},
		},
		{
			CapabilityID:     "degraded_mode_evidence_summary",
			DisplayName:      "Degraded-mode evidence summary",
			Supported:        true,
			CurrentState:     mapBoolState(degradedActive, "active", "standby"),
			EvidenceSurfaces: []string{"federation.local_trust_anchors", "federation.stale_peers", "handoff.verification_uri", "validation.execution_refs", "validation.certificate_refs"},
			ContractSummary:  "When execution posture is degraded, ChangeLock surfaces a bounded evidence summary instead of inventing a connected-state view.",
			Limitations: []string{
				"Degraded-mode summary preserves limitation state and evidence references; it is not a substitute for restored remote freshness.",
			},
		},
	}
	return items
}

func buildExecutionDelayedSyncSemantics(syncStatus string, delayedSyncRequired bool) executionDelayedSyncSemantics {
	return executionDelayedSyncSemantics{
		Required:               delayedSyncRequired,
		SyncStatus:             firstNonEmpty(syncStatus, federationSyncStatusLocalOnly),
		PolicyPropagationModel: "Federated policy propagation is bounded by local trust anchors, last verified peer proofs, and explicit safety gates. Delayed sync keeps local admissibility authoritative until freshness and divergence checks recover.",
		SafeReadSemantics: []string{
			"Local trust anchors remain authoritative for disconnected decision-making.",
			"Last verified peer freshness, stale state, and divergence remain visible to operators.",
			"Sealed handoff verification and persisted validation certificates remain usable while sync is delayed.",
		},
		BlockedClaims: []string{
			"Do not claim current upstream peer health when sync is stale or diverged.",
			"Do not claim inherited global policy freshness when local divergence or local-only status is active.",
			"Do not suppress degraded-mode evidence just because local fallback remains available.",
		},
		RecoverySignals: []string{
			"policy sync status returns to synced or synced_with_overrides",
			"stale peers present fresh proofs and circuit state closes",
			"local policy divergence is resolved or remains explicitly acknowledged via overrides",
		},
		Limitations: []string{
			"Delayed sync semantics bound what ChangeLock can safely assert while disconnected; they are not an eventually-consistent guarantee of instant federation convergence.",
		},
	}
}

func executionVMPolicyEvidenceParityContract() executionVMPolicyEvidenceParity {
	return executionVMPolicyEvidenceParity{
		PolicyParity: []string{
			"VM workloads reuse desired-state approval, desired-state verification, signer expectation, and attestation-aware runtime posture semantics from the workload trust model.",
			"Runtime scheduling guidance and mismatch detection stay aligned with the same 3A posture contract used for non-VM workloads.",
			"Validation linkage remains advisory and uses the same execution and certificate reference surfaces when matching strict validation evidence exists.",
		},
		EvidenceParity: []string{
			"VM lineage preserves desired_state_source_ref, desired_state_approval_id, desired_state_verification_state, approved_digest, running_digest, evidence_refs, and readback_refs as first-class evidence surfaces.",
			"Attestation inputs, expected signers, SBOM verification, and runtime posture are derived from the same evidence-native runtime snapshot model as other workloads.",
		},
		PostureParity: []string{
			"Expected and actual trust state remain explicit through runtime_posture.expected_trust_state, runtime_posture.actual_trust_state, and mismatch evidence.",
		},
		Limitations: []string{
			"VM parity is bounded to the evidence present in ChangeLock; hypervisor, node, or confidential substrate guarantees are not inferred when that substrate evidence is absent.",
			"VM support is not a separate truth layer and does not collapse lineage into a mere workload label swap.",
		},
	}
}

func executionEphemeralRetentionContractModel() executionEphemeralRetentionContract {
	return executionEphemeralRetentionContract{
		SnapshotSemantics: []string{
			"Short-lived runtime workloads are represented through bounded desired-state, active-state, and linked runtime evidence snapshots.",
			"Strict validation runs are preserved as isolated execution, verdict, and certificate snapshots rather than converted into production truth.",
		},
		RetentionSemantics: []string{
			"Summary evidence, subject lineage, readback references, and linked certificates remain queryable after the execution unit exits.",
			"Retention is bounded and summary-first; ChangeLock does not claim full process replay for every ephemeral execution.",
		},
		SummarySemantics:    "When a short-lived execution disappears, ChangeLock preserves a bounded summary through execution_ref or subject_ref, completion window, evidence_refs, readback_refs, and certificate linkage.",
		CorrelationBehavior: "Ephemeral workloads correlate by workload identity and subject_ref while resident, then by execution_ref, certificate_ref, and preserved evidence lineage after completion; correlation is not dependent on long-lived pod or PID residency.",
		Limitations: []string{
			"Ephemeral retention preserves bounded lineage and summary evidence, not universal raw process state after exit.",
		},
	}
}

func mapBoolState(value bool, whenTrue, whenFalse string) string {
	if value {
		return whenTrue
	}
	return whenFalse
}

func (s server) executionValidationRefsForWorkload(ctx context.Context, filter runtimeIntegrityFilter, workload string) ([]string, error) {
	if strings.TrimSpace(workload) == "" {
		return nil, nil
	}
	runs, _, err := s.listStrictValidationRuns(ctx, validationFilterFromRuntimeFilter(filter, workload))
	if err != nil {
		return nil, err
	}
	refs := []string{}
	for _, run := range runs {
		if run.Certificate.CertificateID != "" {
			refs = append(refs, "/v1/validation/certificates/"+run.Certificate.CertificateID)
		}
		for _, execution := range run.Executions {
			refs = append(refs, "/v1/validation/executions/"+execution.ExecutionID)
		}
	}
	refs = uniqueStrings(refs)
	if len(refs) > filter.Limit {
		refs = refs[:filter.Limit]
	}
	return refs, nil
}

func validationFilterFromRuntimeFilter(filter runtimeIntegrityFilter, service string) validationHarnessFilter {
	return validationHarnessFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Service:     strings.TrimSpace(service),
		Limit:       minInt(maxInt(filter.Limit, 10), validationHarnessLimit),
		event: audit.EventFilter{
			ClusterID:   filter.ClusterID,
			TenantID:    filter.TenantID,
			Environment: filter.Environment,
			Repo:        filter.Repo,
			Limit:       maxInt(filter.Limit*20, 500),
		},
	}
}

func parseExecutionCoverageFilter(r *http.Request) (runtimeIntegrityFilter, error) {
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		return filter, err
	}
	if strings.TrimSpace(r.URL.Query().Get("workload_kind")) == "" && strings.TrimSpace(r.URL.Query().Get("subject_ref")) == "" {
		filter.WorkloadKind = ""
	}
	return filter, nil
}

func (s server) listStoredHandoffs(ctx context.Context, filter runtimeIntegrityFilter) ([]handoffStoredRecord, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		Component:   handoffComponent,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       maxInt(filter.Limit*20, 200),
	})
	if err != nil {
		return nil, err
	}
	latest := map[string]handoffStoredRecord{}
	latestAt := map[string]time.Time{}
	for _, event := range events {
		if event.EventType != audit.EventTypeHandoffSealed && event.EventType != audit.EventTypeHandoffCosigned {
			continue
		}
		if len(event.Handoff) == 0 || string(event.Handoff) == "null" {
			continue
		}
		var record handoffStoredRecord
		if err := json.Unmarshal(event.Handoff, &record); err != nil {
			continue
		}
		if filter.TenantID != "" && strings.TrimSpace(record.Manifest.Scope.TenantID) != filter.TenantID {
			continue
		}
		if filter.Environment != "" && strings.TrimSpace(record.Manifest.Scope.Environment) != filter.Environment {
			continue
		}
		if filter.Repo != "" && strings.TrimSpace(record.Manifest.Scope.Repo) != filter.Repo {
			continue
		}
		if currentAt, ok := latestAt[record.PackageID]; !ok || event.ReceivedAt.After(currentAt) {
			latest[record.PackageID] = record
			latestAt[record.PackageID] = event.ReceivedAt
		}
	}
	items := make([]handoffStoredRecord, 0, len(latest))
	for _, item := range latest {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Manifest.CreatedAt.Equal(items[j].Manifest.CreatedAt) {
			return items[i].PackageID < items[j].PackageID
		}
		return items[i].Manifest.CreatedAt.After(items[j].Manifest.CreatedAt)
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, nil
}

func isVMWorkloadKind(kind string) bool {
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case "vm", "virtualmachine", "virtualmachineinstance", "kubevirtvirtualmachine", "kubevirtvirtualmachineinstance":
		return true
	default:
		return false
	}
}

func isEphemeralWorkloadKind(kind string) bool {
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case "job", "cronjob", "batchjob":
		return true
	default:
		return false
	}
}

func runtimeDesiredStateSource(subject *runtimeSnapshotSubject) string {
	if subject.ActiveState != nil && strings.TrimSpace(subject.ActiveState.DesiredStateSourceRef) != "" {
		return strings.TrimSpace(subject.ActiveState.DesiredStateSourceRef)
	}
	if subject.DesiredState != nil {
		return strings.TrimSpace(subject.DesiredState.DesiredStateSourceRef)
	}
	return ""
}

func runtimeDesiredStateApproval(subject *runtimeSnapshotSubject) string {
	if subject.ActiveState != nil && strings.TrimSpace(subject.ActiveState.DesiredStateApprovalID) != "" {
		return strings.TrimSpace(subject.ActiveState.DesiredStateApprovalID)
	}
	if subject.DesiredState != nil {
		return strings.TrimSpace(subject.DesiredState.DesiredStateApprovalID)
	}
	return ""
}

func runtimeApprovedDigest(subject *runtimeSnapshotSubject) string {
	if subject.DesiredState != nil && strings.TrimSpace(subject.DesiredState.ApprovedDigest) != "" {
		return strings.TrimSpace(subject.DesiredState.ApprovedDigest)
	}
	if subject.ActiveState != nil {
		return strings.TrimSpace(subject.ActiveState.ApprovedDigest)
	}
	return ""
}

func runtimeObservedDigest(subject *runtimeSnapshotSubject) string {
	if subject.ActiveState != nil {
		return strings.TrimSpace(subject.ActiveState.ObservedDigest)
	}
	return ""
}

func ephemeralRuntimeWindow(subject *runtimeSnapshotSubject) (time.Time, time.Time) {
	startedAt := time.Time{}
	completedAt := time.Time{}
	if subject.DesiredState != nil && !subject.DesiredState.LastApprovedAt.IsZero() {
		startedAt = subject.DesiredState.LastApprovedAt
	}
	if subject.ActiveState != nil && !subject.ActiveState.LastReconciledAt.IsZero() {
		completedAt = subject.ActiveState.LastReconciledAt
		if startedAt.IsZero() || completedAt.Before(startedAt) {
			startedAt = completedAt
		}
	}
	if startedAt.IsZero() {
		startedAt = runtimeLastVerified(subject)
	}
	if completedAt.IsZero() {
		completedAt = runtimeLastVerified(subject)
	}
	return startedAt, completedAt
}

func ephemeralRuntimeStatus(subject *runtimeSnapshotSubject) string {
	if subject.ActiveState != nil && strings.TrimSpace(subject.ActiveState.ReconciliationStatus) != "" {
		return strings.TrimSpace(subject.ActiveState.ReconciliationStatus)
	}
	if subject.LegacyDrift != nil && strings.TrimSpace(subject.LegacyDrift.Status) != "" {
		return strings.TrimSpace(subject.LegacyDrift.Status)
	}
	return "observed"
}

func (s server) hydrateExecutionSubjectStates(ctx context.Context, filter runtimeIntegrityFilter, subject *runtimeSnapshotSubject) (*runtimeSnapshotSubject, error) {
	hydrated := cloneRuntimeSnapshotSubject(subject)
	if hydrated.DesiredState != nil && hydrated.ActiveState != nil {
		return hydrated, nil
	}
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		ClusterID:   firstNonEmpty(filter.ClusterID, hydrated.Cluster),
		TenantID:    firstNonEmpty(filter.TenantID, hydrated.TenantID),
		Environment: firstNonEmpty(filter.Environment, hydrated.Environment),
		Repo:        firstNonEmpty(filter.Repo, hydrated.Repo),
		Limit:       maxInt(filter.Limit*20, 200),
	})
	if err != nil {
		return nil, err
	}
	if hydrated.DesiredState == nil {
		desired := audit.DeriveRuntimeDesiredStates(events, audit.RuntimeDesiredStateFilter{
			ClusterID:    hydrated.Cluster,
			TenantID:     hydrated.TenantID,
			Namespace:    hydrated.Namespace,
			WorkloadKind: hydrated.WorkloadKind,
			Workload:     hydrated.Workload,
			Limit:        1,
		})
		if len(desired) > 0 {
			copyDesired := desired[0]
			hydrated.DesiredState = &copyDesired
			syncExecutionSubjectFromDesiredState(hydrated, copyDesired)
		}
	}
	if hydrated.ActiveState == nil {
		active := audit.DeriveRuntimeActiveStates(events, audit.RuntimeActiveStateFilter{
			ClusterID:            hydrated.Cluster,
			TenantID:             hydrated.TenantID,
			Namespace:            hydrated.Namespace,
			WorkloadKind:         hydrated.WorkloadKind,
			Workload:             hydrated.Workload,
			ReconciliationStatus: "",
			QuarantineType:       "",
			Limit:                1,
		})
		if len(active) > 0 {
			copyActive := active[0]
			hydrated.ActiveState = &copyActive
			syncExecutionSubjectFromActiveState(hydrated, copyActive)
		}
	}
	return hydrated, nil
}

func cloneRuntimeSnapshotSubject(subject *runtimeSnapshotSubject) *runtimeSnapshotSubject {
	if subject == nil {
		return &runtimeSnapshotSubject{}
	}
	clone := *subject
	if subject.DesiredState != nil {
		copyDesired := *subject.DesiredState
		clone.DesiredState = &copyDesired
	}
	if subject.ActiveState != nil {
		copyActive := *subject.ActiveState
		clone.ActiveState = &copyActive
	}
	if subject.LegacyDrift != nil {
		copyDrift := *subject.LegacyDrift
		clone.LegacyDrift = &copyDrift
	}
	if subject.SBOMVerification != nil {
		copySBOM := *subject.SBOMVerification
		clone.SBOMVerification = &copySBOM
	}
	clone.ExpectedSigners = append([]string(nil), subject.ExpectedSigners...)
	clone.Observations = append([]runtimeObservation(nil), subject.Observations...)
	clone.ProfileHints = append([]runtimeProfileHint(nil), subject.ProfileHints...)
	clone.Enforcements = append([]runtimeEnforcementDecision(nil), subject.Enforcements...)
	clone.EvidenceRefs = map[string]struct{}{}
	for key := range subject.EvidenceRefs {
		clone.EvidenceRefs[key] = struct{}{}
	}
	clone.TrustInputs = map[string]struct{}{}
	for key := range subject.TrustInputs {
		clone.TrustInputs[key] = struct{}{}
	}
	return &clone
}

func syncExecutionSubjectFromDesiredState(subject *runtimeSnapshotSubject, desired audit.RuntimeDesiredStateView) {
	subject.Cluster = firstNonEmpty(subject.Cluster, desired.ClusterID, "local")
	subject.TenantID = firstNonEmpty(subject.TenantID, desired.TenantID)
	subject.Namespace = firstNonEmpty(subject.Namespace, desired.Namespace)
	subject.WorkloadKind = firstNonEmpty(subject.WorkloadKind, desired.WorkloadKind)
	subject.Workload = firstNonEmpty(subject.Workload, desired.Workload)
	subject.ServiceAccount = firstNonEmpty(subject.ServiceAccount, desired.ServiceAccount)
	subject.ImageDigest = firstNonEmpty(subject.ImageDigest, desired.ApprovedDigest)
	if subject.TrustInputs == nil {
		subject.TrustInputs = map[string]struct{}{}
	}
	subject.TrustInputs["desired_state_present"] = struct{}{}
	if subject.EvidenceRefs == nil {
		subject.EvidenceRefs = map[string]struct{}{}
	}
	if strings.TrimSpace(desired.ApprovedDigest) != "" {
		subject.EvidenceRefs[strings.TrimSpace(desired.ApprovedDigest)] = struct{}{}
	}
}

func syncExecutionSubjectFromActiveState(subject *runtimeSnapshotSubject, active audit.RuntimeActiveStateView) {
	subject.Cluster = firstNonEmpty(subject.Cluster, active.ClusterID, "local")
	subject.TenantID = firstNonEmpty(subject.TenantID, active.TenantID)
	subject.Namespace = firstNonEmpty(subject.Namespace, active.Namespace)
	subject.WorkloadKind = firstNonEmpty(subject.WorkloadKind, active.WorkloadKind)
	subject.Workload = firstNonEmpty(subject.Workload, active.Workload)
	subject.ServiceAccount = firstNonEmpty(subject.ServiceAccount, active.ServiceAccount)
	subject.ImageDigest = firstNonEmpty(subject.ImageDigest, active.ObservedDigest, active.ApprovedDigest)
	if subject.EvidenceRefs == nil {
		subject.EvidenceRefs = map[string]struct{}{}
	}
	for _, ref := range compactStrings(active.ObservedDigest, active.ApprovedDigest, active.ExpectedConfigHash, active.ObservedConfigHash) {
		subject.EvidenceRefs[ref] = struct{}{}
	}
}
