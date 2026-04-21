package main

import (
	"context"
	"net/http"

	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	runtimePostureLinkageSchemaVersion    = "3a.runtime_posture_linkage.v1"
	runtimeBoundaryDisciplineSchema       = "3a.runtime_boundary_discipline.v1"
	runtimeBoundaryPhasePreExecution      = "pre_execution_guidance"
	runtimeBoundaryPhaseNearExecution     = "near_execution_observation"
	runtimeBoundaryPhasePostExecution     = "post_execution_containment"
	runtimeBoundaryCoverageProcessExec    = "process_exec_and_identity_lineage"
	runtimeBoundaryCoverageEgressTopology = "egress_and_topology"
	runtimeBoundaryCoverageFilesystemSBOM = "filesystem_library_and_sbom"
	runtimeBoundaryCoverageMemoryFileless = "in_memory_and_fileless_anomaly"
)

type runtimePostureLinkageSchedulingSemantics struct {
	Decision          string   `json:"decision"`
	Meaning           string   `json:"meaning"`
	ApprovalMode      string   `json:"approval_mode"`
	TriggeredWhen     []string `json:"triggered_when,omitempty"`
	RecommendedAction string   `json:"recommended_action,omitempty"`
	Limitations       []string `json:"limitations,omitempty"`
}

type runtimePostureLinkageMismatchSemantics struct {
	Code                string `json:"code"`
	Meaning             string `json:"meaning"`
	SchedulingImpact    string `json:"scheduling_impact"`
	EvidenceExpectation string `json:"evidence_expectation"`
}

type runtimePostureLinkageSemantics struct {
	ModuleReadinessContract []string                                   `json:"module_readiness_contract,omitempty"`
	ExpectedActualContract  []string                                   `json:"expected_actual_contract,omitempty"`
	SchedulingDecisionModel []runtimePostureLinkageSchedulingSemantics `json:"scheduling_decision_model,omitempty"`
	MismatchModel           []runtimePostureLinkageMismatchSemantics   `json:"mismatch_model,omitempty"`
	Limitations             []string                                   `json:"limitations,omitempty"`
}

type runtimePostureLinkageSummary struct {
	TotalSubjects       int            `json:"total_subjects"`
	RuntimeModuleReady  int            `json:"runtime_module_ready"`
	SchedulingDecisions map[string]int `json:"scheduling_decisions,omitempty"`
	MismatchCounts      map[string]int `json:"mismatch_counts,omitempty"`
	ApprovalModeCounts  map[string]int `json:"approval_mode_counts,omitempty"`
}

type runtimePostureLinkageResponse struct {
	SchemaVersion string                         `json:"schema_version"`
	Semantics     runtimePostureLinkageSemantics `json:"semantics"`
	Summary       runtimePostureLinkageSummary   `json:"summary"`
	Items         []runtimePostureState          `json:"items"`
	Limitations   []string                       `json:"limitations,omitempty"`
}

type runtimeBoundarySignalPath struct {
	CurrentPathModel        string   `json:"current_path_model"`
	KernelAdjacentReadiness string   `json:"kernel_adjacent_readiness"`
	TimingSemantics         []string `json:"timing_semantics,omitempty"`
	UnsupportedClaims       []string `json:"unsupported_claims,omitempty"`
}

type runtimeBoundaryEnforcementPhase struct {
	Phase              string   `json:"phase"`
	Applicability      string   `json:"applicability"`
	SupportedRulePacks []string `json:"supported_rule_packs,omitempty"`
	SupportedActions   []string `json:"supported_actions,omitempty"`
	Limitations        []string `json:"limitations,omitempty"`
}

type runtimeBoundaryCoverage struct {
	BoundaryID            string   `json:"boundary_id"`
	DisplayName           string   `json:"display_name"`
	CoverageState         string   `json:"coverage_state"`
	EvidenceModel         []string `json:"evidence_model,omitempty"`
	SupportedFindingTypes []string `json:"supported_finding_types,omitempty"`
	SupportedRulePacks    []string `json:"supported_rule_packs,omitempty"`
	SupportedActions      []string `json:"supported_actions,omitempty"`
	Limitations           []string `json:"limitations,omitempty"`
}

type runtimeBoundaryOverheadCeiling struct {
	MeasurementStatus      string   `json:"measurement_status"`
	ControlPlaneBudget     string   `json:"control_plane_budget"`
	StartingPointRefs      []string `json:"starting_point_refs,omitempty"`
	ResourceStartingPoints []string `json:"resource_starting_points,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type runtimeBoundaryDisciplineResponse struct {
	SchemaVersion      string                            `json:"schema_version"`
	SignalPath         runtimeBoundarySignalPath         `json:"signal_path"`
	EnforcementPhases  []runtimeBoundaryEnforcementPhase `json:"enforcement_phases,omitempty"`
	CoverageBoundaries []runtimeBoundaryCoverage         `json:"coverage_boundaries,omitempty"`
	OverheadCeiling    runtimeBoundaryOverheadCeiling    `json:"overhead_ceiling"`
	Limitations        []string                          `json:"limitations,omitempty"`
}

func (s server) runtimePostureLinkageHandler(w http.ResponseWriter, r *http.Request) {
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
	items, limitations, err := s.buildRuntimePostureStates(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, runtimePostureLinkageResponse{
		SchemaVersion: runtimePostureLinkageSchemaVersion,
		Semantics:     runtimePostureLinkageSemanticsCatalog(),
		Summary:       summarizeRuntimePostureLinkage(items),
		Items:         items,
		Limitations: uniqueStrings(append([]string{
			"Posture linkage is bounded to workload-scoped desired-state, active-state, artifact, SBOM, and runtime evidence already present in ChangeLock.",
			"Scheduling linkage remains evidence-backed and does not claim node-level attestation or substrate guarantees when that evidence is not in scope.",
		}, limitations...)),
	})
}

func (s server) runtimeBoundaryDisciplineHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, runtimeBoundaryDisciplineCatalog())
}

func runtimePostureLinkageSemanticsCatalog() runtimePostureLinkageSemantics {
	return runtimePostureLinkageSemantics{
		ModuleReadinessContract: []string{
			"`runtime_module_ready` is true only when both desired-state and active-state runtime views are present for the workload.",
			"Readiness signals remain additive and bounded to canonical runtime evidence already stored by ChangeLock.",
			"Missing attestation, signer, SBOM, or critical-finding signals degrade scheduling guidance instead of being silently ignored.",
		},
		ExpectedActualContract: []string{
			"`expected_trust_state` describes approved or expected trust posture from desired-state verification, signer expectations, SBOM state, and assigned sandbox class.",
			"`actual_trust_state` reflects the observed runtime posture derived from current evidence and enforcement state.",
			"Mismatch evaluation is attestation-aware and explicitly surfaces when verified desired state lacks runtime attestation provenance.",
		},
		SchedulingDecisionModel: []runtimePostureLinkageSchedulingSemantics{
			{
				Decision:          runtimeSchedulingAllowStandard,
				Meaning:           "Expected and actual trust posture remain within bounded standard scheduling posture.",
				ApprovalMode:      recommendationApprovalAutoSafe,
				TriggeredWhen:     []string{"no critical mismatches", "runtime module ready", "no elevated sandbox restriction"},
				RecommendedAction: runtimeActionAlert,
			},
			{
				Decision:          runtimeSchedulingRestricted,
				Meaning:           "Workload remains schedulable only in a restricted posture because trust inputs, SBOM state, or signer evidence are degraded.",
				ApprovalMode:      recommendationApprovalHumanReview,
				TriggeredWhen:     []string{"missing attestation or signer evidence", "SBOM drift", "restricted or hardened sandbox assignment"},
				RecommendedAction: runtimeActionCaptureForensics,
			},
			{
				Decision:          runtimeSchedulingIsolatedReview,
				Meaning:           "Workload requires isolated review posture because critical runtime findings or severe trust mismatches remain active.",
				ApprovalMode:      recommendationApprovalHumanReview,
				TriggeredWhen:     []string{"critical runtime findings present", "isolated review sandbox assigned", "identity or attestation mismatch remains unresolved"},
				RecommendedAction: runtimeActionApplyNetworkIsolation,
				Limitations: []string{
					"Isolated review is a bounded scheduling and response posture. It is not a claim of universal pre-execution blocking.",
				},
			},
		},
		MismatchModel: []runtimePostureLinkageMismatchSemantics{
			{
				Code:                runtimeMismatchDesiredState,
				Meaning:             "Desired state is missing or not currently verified.",
				SchedulingImpact:    runtimeSchedulingRestricted,
				EvidenceExpectation: "Expected desired-state verification state should be visible in runtime desired-state evidence.",
			},
			{
				Code:                runtimeMismatchAttestation,
				Meaning:             "Verified desired state lacks matching runtime attestation provenance in scope.",
				SchedulingImpact:    runtimeSchedulingRestricted,
				EvidenceExpectation: "Attestation provenance must remain linked to the workload trust inputs when desired state is verified.",
			},
			{
				Code:                runtimeMismatchSigner,
				Meaning:             "Expected signer-backed artifact evidence is missing from the current runtime scope.",
				SchedulingImpact:    runtimeSchedulingRestricted,
				EvidenceExpectation: "Signed-artifact evidence should remain attached when signer expectations exist.",
			},
			{
				Code:                runtimeMismatchIdentity,
				Meaning:             "Observed runtime identity or attestation-linked trust path diverges from expected posture.",
				SchedulingImpact:    runtimeSchedulingIsolatedReview,
				EvidenceExpectation: "Identity drift or attestation mismatch findings must remain explainable through linked runtime evidence.",
			},
			{
				Code:                runtimeMismatchSBOM,
				Meaning:             "Observed runtime digest or library state diverges from the approved SBOM-linked artifact posture.",
				SchedulingImpact:    runtimeSchedulingRestricted,
				EvidenceExpectation: "SBOM verification or unexpected artifact refs should explain the drift.",
			},
			{
				Code:                runtimeMismatchCriticalFindings,
				Meaning:             "Critical runtime findings remain active in the current scope.",
				SchedulingImpact:    runtimeSchedulingIsolatedReview,
				EvidenceExpectation: "Critical findings must remain linked to runtime evidence, forensics, and explainable response posture.",
			},
		},
		Limitations: []string{
			"Posture linkage is workload-scoped and does not infer node, hypervisor, enclave, or confidential-computing guarantees without explicit substrate evidence.",
		},
	}
}

func summarizeRuntimePostureLinkage(items []runtimePostureState) runtimePostureLinkageSummary {
	summary := runtimePostureLinkageSummary{
		TotalSubjects:       len(items),
		SchedulingDecisions: map[string]int{},
		MismatchCounts:      map[string]int{},
		ApprovalModeCounts:  map[string]int{},
	}
	for _, item := range items {
		if item.RuntimeModuleReady {
			summary.RuntimeModuleReady++
		}
		if item.SchedulingGuidance.Decision != "" {
			summary.SchedulingDecisions[item.SchedulingGuidance.Decision]++
		}
		if item.SchedulingGuidance.ApprovalMode != "" {
			summary.ApprovalModeCounts[item.SchedulingGuidance.ApprovalMode]++
		}
		for _, mismatch := range item.Mismatches {
			if mismatch.Code != "" {
				summary.MismatchCounts[mismatch.Code]++
			}
		}
	}
	return summary
}

func runtimeBoundaryDisciplineCatalog() runtimeBoundaryDisciplineResponse {
	return runtimeBoundaryDisciplineResponse{
		SchemaVersion: runtimeBoundaryDisciplineSchema,
		SignalPath: runtimeBoundarySignalPath{
			CurrentPathModel:        "low_latency_runtime_observation_with_bounded_near_execution_response",
			KernelAdjacentReadiness: "kernel_adjacent_ready_but_not_universally_claimed",
			TimingSemantics: []string{
				"Current runtime protection is evidence-backed and operates on runtime observation, trust linkage, and bounded response decisions already in scope.",
				"Near-execution containment is only claimed where explicit runtime evidence arrives early enough to drive bounded response or scheduling restriction.",
				"Pre-execution claims remain limited to guidance and posture restriction; ChangeLock does not claim universal pre-CPU blocking semantics.",
			},
			UnsupportedClaims: []string{
				"zero-latency monitoring",
				"attack blocked before it executes in the processor in all cases",
				"universal fileless protection",
				"full live memory scanning across every substrate",
			},
		},
		EnforcementPhases: []runtimeBoundaryEnforcementPhase{
			{
				Phase:              runtimeBoundaryPhasePreExecution,
				Applicability:      "Bounded scheduling and posture guidance before broader containment or recovery decisions are taken.",
				SupportedRulePacks: []string{"runtime_identity_and_attestation", "sbom_runtime_integrity"},
				SupportedActions:   []string{runtimeActionAlert, runtimeActionCaptureForensics},
				Limitations: []string{
					"Pre-execution guidance is evidence-backed posture restriction, not universal process blocking.",
				},
			},
			{
				Phase:              runtimeBoundaryPhaseNearExecution,
				Applicability:      "Runtime observation arrives close enough to execution, egress, or mutation time to drive bounded recommendation or containment semantics.",
				SupportedRulePacks: []string{"binary_execution_integrity", "runtime_identity_and_attestation", "outbound_and_topology_expansion", "filesystem_and_memory_execution"},
				SupportedActions:   []string{runtimeActionCaptureForensics, runtimeActionRecommendQuarantine, runtimeActionApplyNetworkIsolation},
				Limitations: []string{
					"Near-execution handling depends on the presence and timing of explicit runtime evidence rather than a blanket kernel-wide interception claim.",
				},
			},
			{
				Phase:              runtimeBoundaryPhasePostExecution,
				Applicability:      "Bounded containment, hardening, rollback, and trusted-recovery decisions after evidence-backed runtime findings are established.",
				SupportedRulePacks: []string{"binary_execution_integrity", "runtime_identity_and_attestation", "sbom_runtime_integrity", "outbound_and_topology_expansion", "privilege_and_profile_drift", "filesystem_and_memory_execution"},
				SupportedActions:   []string{runtimeActionCaptureForensics, runtimeActionRecommendQuarantine, runtimeActionApplyNetworkIsolation, runtimeActionRestartTrusted},
				Limitations: []string{
					"Post-execution response remains bounded by approval gates, forensic-first policy, rollback discipline, and workload-scoped blast-radius limits.",
				},
			},
		},
		CoverageBoundaries: []runtimeBoundaryCoverage{
			{
				BoundaryID:            runtimeBoundaryCoverageProcessExec,
				DisplayName:           "Process Execution and Identity Lineage",
				CoverageState:         "bounded_near_execution_or_observation",
				EvidenceModel:         []string{"runtime observation", "desired-state digest", "artifact verification evidence", "forensic context"},
				SupportedFindingTypes: []string{runtimeFindingUnknownBinaryExec, runtimeFindingUnsignedBinaryExec, runtimeFindingIdentityDrift, runtimeFindingContainerIDMismatch, runtimeFindingAttestationMismatch},
				SupportedRulePacks:    []string{"binary_execution_integrity", "runtime_identity_and_attestation"},
				SupportedActions:      []string{runtimeActionCaptureForensics, runtimeActionRecommendQuarantine, runtimeActionApplyNetworkIsolation},
				Limitations: []string{
					"Coverage is bounded to evidence-backed execution and identity lineage; it does not claim universal pre-CPU prevention.",
				},
			},
			{
				BoundaryID:            runtimeBoundaryCoverageEgressTopology,
				DisplayName:           "Egress and Topology Expansion",
				CoverageState:         "bounded_runtime_observation",
				EvidenceModel:         []string{"runtime network observation", "topology blast-radius context", "incident linkage"},
				SupportedFindingTypes: []string{runtimeFindingOutboundDrift, runtimeFindingTopologyExpansion},
				SupportedRulePacks:    []string{"outbound_and_topology_expansion"},
				SupportedActions:      []string{runtimeActionRecommendQuarantine, runtimeActionApplyNetworkIsolation},
				Limitations: []string{
					"Egress coverage is bounded to observed runtime network drift and topology context already in scope; it is not a universal packet-inspection claim.",
				},
			},
			{
				BoundaryID:            runtimeBoundaryCoverageFilesystemSBOM,
				DisplayName:           "Filesystem, Library, and SBOM Integrity",
				CoverageState:         "bounded_runtime_integrity",
				EvidenceModel:         []string{"runtime SBOM verification", "observed library refs", "filesystem mutation observations", "approved and observed digests"},
				SupportedFindingTypes: []string{runtimeFindingUnexpectedLibrary, runtimeFindingSBOMMismatch, runtimeFindingFilesystemMutation},
				SupportedRulePacks:    []string{"sbom_runtime_integrity", "filesystem_and_memory_execution"},
				SupportedActions:      []string{runtimeActionCaptureForensics, runtimeActionRestartTrusted},
				Limitations: []string{
					"Filesystem and library coverage is bounded to the digest, library, and mutation signals already captured by runtime evidence.",
				},
			},
			{
				BoundaryID:            runtimeBoundaryCoverageMemoryFileless,
				DisplayName:           "In-Memory and Fileless Anomaly Handling",
				CoverageState:         "bounded_anomaly_signal",
				EvidenceModel:         []string{"runtime observation", "executable-memory anomaly signal", "forensic context"},
				SupportedFindingTypes: []string{runtimeFindingMemoryExecAnomaly},
				SupportedRulePacks:    []string{"filesystem_and_memory_execution"},
				SupportedActions:      []string{runtimeActionCaptureForensics, runtimeActionRecommendQuarantine},
				Limitations: []string{
					"In-memory and fileless handling is limited to bounded anomaly and execution signals; it does not claim universal real-time memory scanning or full memory acquisition.",
				},
			},
		},
		OverheadCeiling: runtimeBoundaryOverheadCeiling{
			MeasurementStatus:  "starting_points_only_not_benchmark_guarantee",
			ControlPlaneBudget: "Bounded runtime aggregation and validation should stay within the existing control-plane latency budget discipline used by strict validation.",
			StartingPointRefs: []string{
				"docs/operations/sizing.md",
				"validation scenario: control_plane_latency_budget",
			},
			ResourceStartingPoints: []string{
				"audit-writer request 250m / 512Mi, limit 1 CPU / 1Gi",
				"runtime-agent request 100m / 192Mi, limit 500m / 384Mi",
			},
			Limitations: []string{
				"These values are practical starting points, not benchmark-backed guarantees.",
				"Boundary discipline explicitly excludes universal always-on memory scanning and cluster-wide kernel interception claims from the current overhead model.",
			},
		},
		Limitations: []string{
			"Runtime boundary discipline documents bounded enforcement and coverage semantics for the current runtime slice; it is not a claim of technical monopoly, universal pre-execution blocking, or substrate-wide invisibility.",
		},
	}
}
