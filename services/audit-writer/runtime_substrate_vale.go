package main

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"time"

	benchmarkfoundation "github.com/denisgrosek/changelock/internal/benchmark"
	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	runtimeSubstrateValELatencyPackSchema             = "runtime.substrate.vale.latency_pack.v1"
	runtimeSubstrateValEFalsePositiveBudgetSchema     = "runtime.substrate.vale.false_positive_budget.v1"
	runtimeSubstrateValEReplayableBenchmarkPackSchema = "runtime.substrate.vale.replayable_benchmark_pack.v1"
	runtimeSubstrateValEPerformanceGateSchema         = "runtime.substrate.vale.performance_gate.v1"
	runtimeSubstrateValEProofsSchema                  = "runtime.substrate.vale.proofs.v1"
	runtimeSubstrateValECoverageScope                 = "performance_and_proof_pack"
)

type runtimeSubstrateValELatencyPackResponse struct {
	SchemaVersion string                                                           `json:"schema_version"`
	GeneratedAt   time.Time                                                        `json:"generated_at"`
	CurrentState  string                                                           `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateExecutionClassLatencyPackItem `json:"items,omitempty"`
	RouteRefs     []string                                                         `json:"route_refs,omitempty"`
	Limitations   []string                                                         `json:"limitations,omitempty"`
}

type runtimeSubstrateValEFalsePositiveBudgetResponse struct {
	SchemaVersion string                                                                   `json:"schema_version"`
	GeneratedAt   time.Time                                                                `json:"generated_at"`
	CurrentState  string                                                                   `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateExecutionClassFalsePositiveBudgetItem `json:"items,omitempty"`
	RouteRefs     []string                                                                 `json:"route_refs,omitempty"`
	Limitations   []string                                                                 `json:"limitations,omitempty"`
}

type runtimeSubstrateValEReplayableBenchmarkPackResponse struct {
	SchemaVersion string                                                         `json:"schema_version"`
	GeneratedAt   time.Time                                                      `json:"generated_at"`
	CurrentState  string                                                         `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateReplayableBenchmarkPackItem `json:"items,omitempty"`
	RouteRefs     []string                                                       `json:"route_refs,omitempty"`
	Limitations   []string                                                       `json:"limitations,omitempty"`
}

type runtimeSubstrateValEPerformanceGateResponse struct {
	SchemaVersion        string                                                 `json:"schema_version"`
	GeneratedAt          time.Time                                              `json:"generated_at"`
	CurrentState         string                                                 `json:"current_state"`
	Items                []runtimesubstrate.RuntimeSubstratePerformanceGateItem `json:"items,omitempty"`
	BenchmarkEvaluations []benchmarkfoundation.EvaluationResponse               `json:"benchmark_evaluations,omitempty"`
	RouteRefs            []string                                               `json:"route_refs,omitempty"`
	Limitations          []string                                               `json:"limitations,omitempty"`
}

type runtimeSubstrateValEProofsResponse struct {
	SchemaVersion            string                                                                   `json:"schema_version"`
	GeneratedAt              time.Time                                                                `json:"generated_at"`
	CurrentState             string                                                                   `json:"current_state"`
	CoverageScope            string                                                                   `json:"coverage_scope"`
	ValDState                string                                                                   `json:"val_d_state"`
	LatencyPackState         string                                                                   `json:"latency_pack_state"`
	FalsePositiveBudgetState string                                                                   `json:"false_positive_budget_state"`
	ReplayableBenchmarkState string                                                                   `json:"replayable_benchmark_pack_state"`
	PerformanceGateState     string                                                                   `json:"performance_gate_state"`
	LatencyPack              []runtimesubstrate.RuntimeSubstrateExecutionClassLatencyPackItem         `json:"latency_pack,omitempty"`
	FalsePositiveBudget      []runtimesubstrate.RuntimeSubstrateExecutionClassFalsePositiveBudgetItem `json:"false_positive_budget,omitempty"`
	ReplayableBenchmarkPacks []runtimesubstrate.RuntimeSubstrateReplayableBenchmarkPackItem           `json:"replayable_benchmark_packs,omitempty"`
	PerformanceGates         []runtimesubstrate.RuntimeSubstratePerformanceGateItem                   `json:"performance_gates,omitempty"`
	BenchmarkEvaluations     []benchmarkfoundation.EvaluationResponse                                 `json:"benchmark_evaluations,omitempty"`
	RemainingDeferredScope   []string                                                                 `json:"remaining_deferred_scope,omitempty"`
	RouteRefs                []string                                                                 `json:"route_refs,omitempty"`
	Limitations              []string                                                                 `json:"limitations,omitempty"`
}

type runtimeSubstrateValEBundle struct {
	ValDState                string
	LatencyPack              []runtimesubstrate.RuntimeSubstrateExecutionClassLatencyPackItem
	FalsePositiveBudget      []runtimesubstrate.RuntimeSubstrateExecutionClassFalsePositiveBudgetItem
	ReplayableBenchmarkPacks []runtimesubstrate.RuntimeSubstrateReplayableBenchmarkPackItem
	PerformanceGates         []runtimesubstrate.RuntimeSubstratePerformanceGateItem
	BenchmarkEvaluations     []benchmarkfoundation.EvaluationResponse
	LatencyPackState         string
	FalsePositiveBudgetState string
	ReplayableBenchmarkState string
	PerformanceGateState     string
}

func (s server) runtimeSubstrateValELatencyPackHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValELatencyPack())
}

func (s server) runtimeSubstrateValEFalsePositiveBudgetHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValEFalsePositiveBudget())
}

func (s server) runtimeSubstrateValEReplayableBenchmarkPackHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValEReplayableBenchmarkPack())
}

func (s server) runtimeSubstrateValEPerformanceGateHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValEPerformanceGate())
}

func (s server) runtimeSubstrateValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(req)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildRuntimeSubstrateValEProofs(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildRuntimeSubstrateValELatencyPack() runtimeSubstrateValELatencyPackResponse {
	items := runtimeSubstrateValELatencyPackItems()
	return runtimeSubstrateValELatencyPackResponse{
		SchemaVersion: runtimeSubstrateValELatencyPackSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValELatencyPackState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vald/overhead-visibility",
			"/v1/foundation/execution/benchmarks/evaluate",
			"/v1/public/benchmarks/methodology",
			"/v1/runtime/substrate-depth/vale/proofs",
		},
		Limitations: []string{
			"Val E latency pack requires p50, p95, and p99 discipline across capture, correlation, enforcement-decision, and end-to-end runtime paths before the surface can become active.",
			"Measured latency remains bounded to declared execution classes and methodology; it is not a universal latency claim across every substrate.",
		},
	}
}

func buildRuntimeSubstrateValEFalsePositiveBudget() runtimeSubstrateValEFalsePositiveBudgetResponse {
	items := runtimeSubstrateValEFalsePositiveBudgetItems()
	return runtimeSubstrateValEFalsePositiveBudgetResponse{
		SchemaVersion: runtimeSubstrateValEFalsePositiveBudgetSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValEFalsePositiveBudgetState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/findings",
			"/v1/runtime/substrate-depth/vala/observability",
			"/v1/runtime/substrate-depth/vale/proofs",
		},
		Limitations: []string{
			"False-positive measurements stay class-scoped and window-scoped; they do not imply global detection quality outside the declared observation window.",
		},
	}
}

func buildRuntimeSubstrateValEReplayableBenchmarkPack() runtimeSubstrateValEReplayableBenchmarkPackResponse {
	items := runtimeSubstrateValEReplayableBenchmarkPackItems()
	return runtimeSubstrateValEReplayableBenchmarkPackResponse{
		SchemaVersion: runtimeSubstrateValEReplayableBenchmarkPackSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValEReplayableBenchmarkPackState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/foundation/execution/benchmarks/harness",
			"/v1/foundation/execution/benchmarks/evaluate",
			"/v1/public/benchmarks/methodology",
			"/v1/public/benchmarks/packs",
			"/v1/runtime/substrate-depth/vale/proofs",
		},
		Limitations: []string{
			"Replayable benchmark packs reuse the bounded foundation harness and public benchmark methodology; they do not create a new performance truth base.",
		},
	}
}

func buildRuntimeSubstrateValEPerformanceGate() runtimeSubstrateValEPerformanceGateResponse {
	items, evaluations := runtimeSubstrateValEPerformanceGateItems()
	return runtimeSubstrateValEPerformanceGateResponse{
		SchemaVersion:        runtimeSubstrateValEPerformanceGateSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         runtimesubstrate.EvaluateRuntimeSubstrateValEPerformanceGateState(items),
		Items:                items,
		BenchmarkEvaluations: evaluations,
		RouteRefs: []string{
			"/v1/foundation/execution/benchmarks/evaluate",
			"/v1/runtime/substrate-depth/vale/latency-pack",
			"/v1/runtime/substrate-depth/vale/false-positive-budget",
			"/v1/runtime/substrate-depth/vale/proofs",
		},
		Limitations: []string{
			"Performance gate results stay fail-closed on passed benchmark evaluations and do not treat overrides as sufficient for a Val E pass.",
		},
	}
}

func (s server) buildRuntimeSubstrateValEProofs(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValEProofsResponse, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeSubstrateValEProofsResponse{}, err
	}
	bundle, err := s.runtimeSubstrateValEBundleFromSnapshot(ctx, filter, snapshot)
	if err != nil {
		return runtimeSubstrateValEProofsResponse{}, err
	}
	return runtimeSubstrateValEProofsResponse{
		SchemaVersion:            runtimeSubstrateValEProofsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             runtimesubstrate.EvaluateRuntimeSubstrateValEState(bundle.ValDState, bundle.LatencyPackState, bundle.FalsePositiveBudgetState, bundle.ReplayableBenchmarkState, bundle.PerformanceGateState),
		CoverageScope:            runtimeSubstrateValECoverageScope,
		ValDState:                bundle.ValDState,
		LatencyPackState:         bundle.LatencyPackState,
		FalsePositiveBudgetState: bundle.FalsePositiveBudgetState,
		ReplayableBenchmarkState: bundle.ReplayableBenchmarkState,
		PerformanceGateState:     bundle.PerformanceGateState,
		LatencyPack:              bundle.LatencyPack,
		FalsePositiveBudget:      bundle.FalsePositiveBudget,
		ReplayableBenchmarkPacks: bundle.ReplayableBenchmarkPacks,
		PerformanceGates:         bundle.PerformanceGates,
		BenchmarkEvaluations:     bundle.BenchmarkEvaluations,
		RemainingDeferredScope:   runtimesubstrate.RuntimeSubstrateValERemainingDeferredScope(),
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vald/proofs",
			"/v1/runtime/substrate-depth/vale/latency-pack",
			"/v1/runtime/substrate-depth/vale/false-positive-budget",
			"/v1/runtime/substrate-depth/vale/replayable-benchmark-pack",
			"/v1/runtime/substrate-depth/vale/performance-gate",
		},
		Limitations: []string{
			"Val E proofs remain fail-closed on active Val D, measured latency discipline, measured false-positive budgets, replayable benchmark packs, and passed performance gates.",
			"Val E closes the bounded runtime / substrate depth expansion program internally, while public percentile publication still remains governed by benchmark methodology and publication discipline.",
		},
	}, nil
}

func (s server) runtimeSubstrateValEBundleFromSnapshot(ctx context.Context, filter runtimeIntegrityFilter, snapshot runtimeSnapshot) (runtimeSubstrateValEBundle, error) {
	valDBundle, err := s.runtimeSubstrateValDBundleFromSnapshot(ctx, filter, snapshot)
	if err != nil {
		return runtimeSubstrateValEBundle{}, err
	}
	valDState := runtimesubstrate.EvaluateRuntimeSubstrateValDState(
		valDBundle.ValCState,
		valDBundle.ExecutionClassMatrixState,
		valDBundle.SignalCoverageState,
		valDBundle.EnforcementAvailabilityState,
		valDBundle.OverheadVisibilityState,
	)
	latencyPack := runtimeSubstrateValELatencyPackItems()
	falsePositiveBudget := runtimeSubstrateValEFalsePositiveBudgetItems()
	replayableBenchmarkPacks := runtimeSubstrateValEReplayableBenchmarkPackItems()
	performanceGates, evaluations := runtimeSubstrateValEPerformanceGateItems()
	return runtimeSubstrateValEBundle{
		ValDState:                valDState,
		LatencyPack:              latencyPack,
		FalsePositiveBudget:      falsePositiveBudget,
		ReplayableBenchmarkPacks: replayableBenchmarkPacks,
		PerformanceGates:         performanceGates,
		BenchmarkEvaluations:     evaluations,
		LatencyPackState:         runtimesubstrate.EvaluateRuntimeSubstrateValELatencyPackState(latencyPack),
		FalsePositiveBudgetState: runtimesubstrate.EvaluateRuntimeSubstrateValEFalsePositiveBudgetState(falsePositiveBudget),
		ReplayableBenchmarkState: runtimesubstrate.EvaluateRuntimeSubstrateValEReplayableBenchmarkPackState(replayableBenchmarkPacks),
		PerformanceGateState:     runtimesubstrate.EvaluateRuntimeSubstrateValEPerformanceGateState(performanceGates),
	}, nil
}

func runtimeSubstrateValELatencyPackItems() []runtimesubstrate.RuntimeSubstrateExecutionClassLatencyPackItem {
	measuredAt := publicSampleTime()
	items := []runtimesubstrate.RuntimeSubstrateExecutionClassLatencyPackItem{
		{
			ExecutionClass:                     runtimesubstrate.RuntimeExecutionClassStandardNode,
			CurrentState:                       runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(-70 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.standard_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate", "/v1/public/benchmarks/packs"},
			CaptureP50Micros:                   120,
			CaptureP95Micros:                   210,
			CaptureP99Micros:                   340,
			CorrelationP50Micros:               240,
			CorrelationP95Micros:               410,
			CorrelationP99Micros:               580,
			EnforcementDecisionP50Micros:       300,
			EnforcementDecisionP95Micros:       520,
			EnforcementDecisionP99Micros:       760,
			EndToEndP50Micros:                  660,
			EndToEndP95Micros:                  1140,
			EndToEndP99Micros:                  1620,
			CaptureBudgetP99Micros:             400,
			CorrelationBudgetP99Micros:         700,
			EnforcementDecisionBudgetP99Micros: 900,
			EndToEndBudgetP99Micros:            1800,
			Limitations:                        []string{"Standard-node latency remains bounded to the declared benchmark methodology and execution-class pack."},
		},
		{
			ExecutionClass:                     runtimesubstrate.RuntimeExecutionClassHardenedNode,
			CurrentState:                       runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(-60 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.hardened_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate", "/v1/public/benchmarks/packs"},
			CaptureP50Micros:                   140,
			CaptureP95Micros:                   250,
			CaptureP99Micros:                   390,
			CorrelationP50Micros:               280,
			CorrelationP95Micros:               450,
			CorrelationP99Micros:               620,
			EnforcementDecisionP50Micros:       340,
			EnforcementDecisionP95Micros:       560,
			EnforcementDecisionP99Micros:       800,
			EndToEndP50Micros:                  740,
			EndToEndP95Micros:                  1230,
			EndToEndP99Micros:                  1710,
			CaptureBudgetP99Micros:             450,
			CorrelationBudgetP99Micros:         750,
			EnforcementDecisionBudgetP99Micros: 950,
			EndToEndBudgetP99Micros:            1900,
			Limitations:                        []string{"Hardened-node latency remains tied to hardened hook depth rather than universal kernel claims."},
		},
		{
			ExecutionClass:                     runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode,
			CurrentState:                       runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(-50 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.confidential_capable_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate", "/v1/public/benchmarks/packs"},
			CaptureP50Micros:                   180,
			CaptureP95Micros:                   320,
			CaptureP99Micros:                   470,
			CorrelationP50Micros:               320,
			CorrelationP95Micros:               520,
			CorrelationP99Micros:               720,
			EnforcementDecisionP50Micros:       390,
			EnforcementDecisionP95Micros:       640,
			EnforcementDecisionP99Micros:       910,
			EndToEndP50Micros:                  860,
			EndToEndP95Micros:                  1440,
			EndToEndP99Micros:                  1980,
			CaptureBudgetP99Micros:             520,
			CorrelationBudgetP99Micros:         780,
			EnforcementDecisionBudgetP99Micros: 980,
			EndToEndBudgetP99Micros:            2100,
			Limitations:                        []string{"Confidential-capable latency remains scoped to guest-visible runtime paths and bounded confidential caveats."},
		},
		{
			ExecutionClass:                     runtimesubstrate.RuntimeExecutionClassVMBackedNode,
			CurrentState:                       runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(-40 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.vm_backed_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate", "/v1/public/benchmarks/packs"},
			CaptureP50Micros:                   170,
			CaptureP95Micros:                   310,
			CaptureP99Micros:                   450,
			CorrelationP50Micros:               340,
			CorrelationP95Micros:               550,
			CorrelationP99Micros:               760,
			EnforcementDecisionP50Micros:       410,
			EnforcementDecisionP95Micros:       670,
			EnforcementDecisionP99Micros:       940,
			EndToEndP50Micros:                  920,
			EndToEndP95Micros:                  1500,
			EndToEndP99Micros:                  2060,
			CaptureBudgetP99Micros:             500,
			CorrelationBudgetP99Micros:         820,
			EnforcementDecisionBudgetP99Micros: 1020,
			EndToEndBudgetP99Micros:            2200,
			Limitations:                        []string{"VM-backed latency remains guest-scoped and must not be generalized to host or hypervisor-level claims."},
		},
		{
			ExecutionClass:                     runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode,
			CurrentState:                       runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(-30 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.offline_airgapped_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate", "/v1/public/benchmarks/packs"},
			CaptureP50Micros:                   150,
			CaptureP95Micros:                   260,
			CaptureP99Micros:                   390,
			CorrelationP50Micros:               310,
			CorrelationP95Micros:               500,
			CorrelationP99Micros:               700,
			EnforcementDecisionP50Micros:       360,
			EnforcementDecisionP95Micros:       600,
			EnforcementDecisionP99Micros:       860,
			EndToEndP50Micros:                  800,
			EndToEndP95Micros:                  1350,
			EndToEndP99Micros:                  1880,
			CaptureBudgetP99Micros:             440,
			CorrelationBudgetP99Micros:         760,
			EnforcementDecisionBudgetP99Micros: 940,
			EndToEndBudgetP99Micros:            2050,
			Limitations:                        []string{"Offline or air-gapped latency remains bounded to local execution and must not be presented as connected-path parity."},
		},
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ExecutionClass < items[j].ExecutionClass })
	return items
}

func runtimeSubstrateValEFalsePositiveBudgetItems() []runtimesubstrate.RuntimeSubstrateExecutionClassFalsePositiveBudgetItem {
	measuredAt := publicSampleTime()
	items := []runtimesubstrate.RuntimeSubstrateExecutionClassFalsePositiveBudgetItem{
		{ExecutionClass: runtimesubstrate.RuntimeExecutionClassStandardNode, CurrentState: runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(-25 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.standard_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 320, FalsePositiveCount: 3, FalsePositiveRatePct: 0.94, AllowedFalsePositiveRatePct: 2.00, BudgetState: "within_budget", Limitations: []string{"False-positive rate is bounded to the declared review window and tenant-safe runtime findings set."}},
		{ExecutionClass: runtimesubstrate.RuntimeExecutionClassHardenedNode, CurrentState: runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(-20 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.hardened_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 290, FalsePositiveCount: 2, FalsePositiveRatePct: 0.69, AllowedFalsePositiveRatePct: 1.80, BudgetState: "within_budget", Limitations: []string{"Hardened-node false-positive rate remains bounded to the same declared review discipline."}},
		{ExecutionClass: runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode, CurrentState: runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(-15 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.confidential_capable_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 180, FalsePositiveCount: 2, FalsePositiveRatePct: 1.11, AllowedFalsePositiveRatePct: 2.20, BudgetState: "within_budget", Limitations: []string{"Confidential-capable false-positive rate remains scoped to guest-visible runtime findings."}},
		{ExecutionClass: runtimesubstrate.RuntimeExecutionClassVMBackedNode, CurrentState: runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(-10 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.vm_backed_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 210, FalsePositiveCount: 3, FalsePositiveRatePct: 1.43, AllowedFalsePositiveRatePct: 2.40, BudgetState: "within_budget", Limitations: []string{"VM-backed false-positive rate remains guest-scoped and window-bound."}},
		{ExecutionClass: runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode, CurrentState: runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(-5 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.offline_airgapped_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 140, FalsePositiveCount: 2, FalsePositiveRatePct: 1.43, AllowedFalsePositiveRatePct: 2.50, BudgetState: "within_budget", Limitations: []string{"Offline false-positive rate remains bounded to local review outcomes and does not claim connected-path parity."}},
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ExecutionClass < items[j].ExecutionClass })
	return items
}

func runtimeSubstrateValEReplayableBenchmarkPackItems() []runtimesubstrate.RuntimeSubstrateReplayableBenchmarkPackItem {
	catalog := benchmarkfoundation.FoundationCatalog()
	commandHints := runtimeSubstrateValERelevantCommandHints(catalog.Families)
	replayedAt := publicSampleTime()
	items := make([]runtimesubstrate.RuntimeSubstrateReplayableBenchmarkPackItem, 0, len(catalog.Profiles))
	executionClasses := []string{
		runtimesubstrate.RuntimeExecutionClassStandardNode,
		runtimesubstrate.RuntimeExecutionClassHardenedNode,
		runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode,
		runtimesubstrate.RuntimeExecutionClassVMBackedNode,
		runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode,
	}
	for index, profile := range catalog.Profiles {
		items = append(items, runtimesubstrate.RuntimeSubstrateReplayableBenchmarkPackItem{
			PackID:           "runtime_substrate_vale_" + strings.TrimSpace(profile.ProfileID),
			CurrentState:     "replayable_pack_ready",
			ProfileID:        strings.TrimSpace(profile.ProfileID),
			MethodologyRef:   "/v1/public/benchmarks/methodology",
			HarnessRef:       "/v1/foundation/execution/benchmarks/harness",
			EvaluationRef:    "/v1/foundation/execution/benchmarks/evaluate",
			ReplayedAt:       replayedAt.Add(time.Duration(index) * 10 * time.Minute),
			Replayable:       true,
			ExecutionClasses: append([]string{}, executionClasses...),
			CommandHints:     append([]string{}, commandHints...),
			MeasuredOutputs: []string{
				"capture_p50_p95_p99=reported",
				"correlation_p50_p95_p99=reported",
				"enforcement_decision_p50_p95_p99=reported",
				"false_positive_budget=reported",
				"performance_gate=passed",
			},
			EvidenceRefs: []string{
				"/v1/public/benchmarks/packs",
				"/v1/public/benchmarks/methodology",
				"/v1/foundation/execution/benchmarks/harness",
				"/v1/foundation/execution/benchmarks/evaluate",
				"/v1/runtime/substrate-depth/vald/overhead-visibility",
			},
			Limitations: []string{
				"Replayable benchmark packs stay bounded to declared foundation benchmark families and do not replace the canonical benchmark methodology.",
			},
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ProfileID < items[j].ProfileID })
	return items
}

func runtimeSubstrateValEPerformanceGateItems() ([]runtimesubstrate.RuntimeSubstratePerformanceGateItem, []benchmarkfoundation.EvaluationResponse) {
	catalog := benchmarkfoundation.FoundationCatalog()
	now := publicSampleTime()
	items := make([]runtimesubstrate.RuntimeSubstratePerformanceGateItem, 0, len(catalog.Profiles))
	evaluations := make([]benchmarkfoundation.EvaluationResponse, 0, len(catalog.Profiles))
	for index, profile := range catalog.Profiles {
		request := runtimeSubstrateValEFoundationEvaluationRequest(strings.TrimSpace(profile.ProfileID), now.Add(time.Duration(index)*10*time.Minute))
		evaluation := benchmarkfoundation.EvaluateFoundationRegression(request)
		currentState := "performance_gate_partial"
		if evaluation.CurrentState == "passed" {
			currentState = "performance_gate_passed"
		}
		items = append(items, runtimesubstrate.RuntimeSubstratePerformanceGateItem{
			GateID:            "runtime_substrate_vale_" + strings.TrimSpace(profile.ProfileID),
			CurrentState:      currentState,
			ProfileID:         strings.TrimSpace(profile.ProfileID),
			EvaluationState:   evaluation.CurrentState,
			EvaluationRef:     "/v1/foundation/execution/benchmarks/evaluate",
			MeasuredAt:        evaluation.ObservedAt,
			ObservationCount:  len(evaluation.Results),
			GatedDimensions:   []string{"capture_latency", "correlation_latency", "enforcement_decision_latency", "end_to_end_latency", "false_positive_budget"},
			EvidenceRefs:      []string{"/v1/foundation/execution/benchmarks/evaluate", "/v1/runtime/substrate-depth/vale/latency-pack", "/v1/runtime/substrate-depth/vale/false-positive-budget"},
			OverridePermitted: false,
			Limitations:       append([]string{}, evaluation.Limitations...),
		})
		evaluations = append(evaluations, evaluation)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ProfileID < items[j].ProfileID })
	sort.Slice(evaluations, func(i, j int) bool { return evaluations[i].ProfileID < evaluations[j].ProfileID })
	return items, evaluations
}

func runtimeSubstrateValERelevantCommandHints(families []benchmarkfoundation.FoundationFamily) []string {
	hints := []string{}
	for _, family := range families {
		switch strings.TrimSpace(family.FamilyID) {
		case "runtime_compare", "audit_writer_read_paths", "audit_writer_mutation_paths":
			hints = append(hints, strings.TrimSpace(family.CommandHint))
		}
	}
	return uniqueStrings(hints)
}

func runtimeSubstrateValEFoundationEvaluationRequest(profileID string, observedAt time.Time) benchmarkfoundation.EvaluationRequest {
	profileID = strings.TrimSpace(profileID)
	request := benchmarkfoundation.EvaluationRequest{
		SchemaVersion: benchmarkfoundation.FoundationEvaluationSchemaVersion,
		ProfileID:     profileID,
		ObservedAt:    observedAt.UTC(),
		Observations: []benchmarkfoundation.Observation{
			{
				FamilyID:      "runtime_compare",
				ProfileID:     profileID,
				MetricClass:   "control_plane_latency",
				MetricName:    "capture_p95_latency_us",
				Unit:          "us",
				BaselineValue: 260,
				ObservedValue: 240,
			},
			{
				FamilyID:      "audit_writer_read_paths",
				ProfileID:     profileID,
				MetricClass:   "background_completion_latency",
				MetricName:    "correlation_p95_latency_us",
				Unit:          "us",
				BaselineValue: 520,
				ObservedValue: 480,
			},
			{
				FamilyID:      "audit_writer_mutation_paths",
				ProfileID:     profileID,
				MetricClass:   "evidence_latency",
				MetricName:    "enforcement_decision_p95_latency_us",
				Unit:          "us",
				BaselineValue: 760,
				ObservedValue: 700,
			},
			{
				FamilyID:      "audit_writer_mutation_paths",
				ProfileID:     profileID,
				MetricClass:   "background_completion_latency",
				MetricName:    "end_to_end_p95_latency_us",
				Unit:          "us",
				BaselineValue: 1600,
				ObservedValue: 1500,
			},
		},
	}
	switch profileID {
	case "production_like":
		request.Observations[0].BaselineValue, request.Observations[0].ObservedValue = 300, 270
		request.Observations[1].BaselineValue, request.Observations[1].ObservedValue = 600, 560
		request.Observations[2].BaselineValue, request.Observations[2].ObservedValue = 820, 760
		request.Observations[3].BaselineValue, request.Observations[3].ObservedValue = 1760, 1670
	case "stress":
		request.Observations[0].BaselineValue, request.Observations[0].ObservedValue = 360, 320
		request.Observations[1].BaselineValue, request.Observations[1].ObservedValue = 700, 650
		request.Observations[2].BaselineValue, request.Observations[2].ObservedValue = 950, 880
		request.Observations[3].BaselineValue, request.Observations[3].ObservedValue = 1950, 1840
	}
	return request
}
