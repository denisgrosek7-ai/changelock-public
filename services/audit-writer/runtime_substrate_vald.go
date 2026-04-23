package main

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	runtimeSubstrateValDExecutionClassMatrixSchema    = "runtime.substrate.vald.execution_class_matrix.v1"
	runtimeSubstrateValDSignalCoverageSchema          = "runtime.substrate.vald.signal_coverage.v1"
	runtimeSubstrateValDEnforcementAvailabilitySchema = "runtime.substrate.vald.enforcement_availability.v1"
	runtimeSubstrateValDOverheadVisibilitySchema      = "runtime.substrate.vald.overhead_visibility.v1"
	runtimeSubstrateValDProofsSchema                  = "runtime.substrate.vald.proofs.v1"
	runtimeSubstrateValDCoverageScope                 = "execution_class_matrix_depth"
)

type runtimeSubstrateValDExecutionClassMatrixResponse struct {
	SchemaVersion string                                                      `json:"schema_version"`
	GeneratedAt   time.Time                                                   `json:"generated_at"`
	CurrentState  string                                                      `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateExecutionClassMatrixItem `json:"items,omitempty"`
	RouteRefs     []string                                                    `json:"route_refs,omitempty"`
	Limitations   []string                                                    `json:"limitations,omitempty"`
}

type runtimeSubstrateValDSignalCoverageResponse struct {
	SchemaVersion string                                                              `json:"schema_version"`
	GeneratedAt   time.Time                                                           `json:"generated_at"`
	CurrentState  string                                                              `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateExecutionClassSignalCoverageItem `json:"items,omitempty"`
	RouteRefs     []string                                                            `json:"route_refs,omitempty"`
	Limitations   []string                                                            `json:"limitations,omitempty"`
}

type runtimeSubstrateValDEnforcementAvailabilityResponse struct {
	SchemaVersion string                                                                       `json:"schema_version"`
	GeneratedAt   time.Time                                                                    `json:"generated_at"`
	CurrentState  string                                                                       `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem `json:"items,omitempty"`
	RouteRefs     []string                                                                     `json:"route_refs,omitempty"`
	Limitations   []string                                                                     `json:"limitations,omitempty"`
}

type runtimeSubstrateValDOverheadVisibilityResponse struct {
	SchemaVersion string                                                                  `json:"schema_version"`
	GeneratedAt   time.Time                                                               `json:"generated_at"`
	CurrentState  string                                                                  `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateExecutionClassOverheadVisibilityItem `json:"items,omitempty"`
	RouteRefs     []string                                                                `json:"route_refs,omitempty"`
	Limitations   []string                                                                `json:"limitations,omitempty"`
}

type runtimeSubstrateValDProofsResponse struct {
	SchemaVersion                string                                                                       `json:"schema_version"`
	GeneratedAt                  time.Time                                                                    `json:"generated_at"`
	CurrentState                 string                                                                       `json:"current_state"`
	CoverageScope                string                                                                       `json:"coverage_scope"`
	ValCState                    string                                                                       `json:"val_c_state"`
	ExecutionClassMatrixState    string                                                                       `json:"execution_class_matrix_state"`
	SignalCoverageState          string                                                                       `json:"signal_coverage_state"`
	EnforcementAvailabilityState string                                                                       `json:"enforcement_availability_state"`
	OverheadVisibilityState      string                                                                       `json:"overhead_visibility_state"`
	ExecutionClassMatrix         []runtimesubstrate.RuntimeSubstrateExecutionClassMatrixItem                  `json:"execution_class_matrix,omitempty"`
	SignalCoverage               []runtimesubstrate.RuntimeSubstrateExecutionClassSignalCoverageItem          `json:"signal_coverage,omitempty"`
	EnforcementAvailability      []runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem `json:"enforcement_availability,omitempty"`
	OverheadVisibility           []runtimesubstrate.RuntimeSubstrateExecutionClassOverheadVisibilityItem      `json:"overhead_visibility,omitempty"`
	RemainingDeferredScope       []string                                                                     `json:"remaining_deferred_scope,omitempty"`
	RouteRefs                    []string                                                                     `json:"route_refs,omitempty"`
	Limitations                  []string                                                                     `json:"limitations,omitempty"`
}

type runtimeSubstrateValDBundle struct {
	ValCState                    string
	ExecutionClassMatrix         []runtimesubstrate.RuntimeSubstrateExecutionClassMatrixItem
	SignalCoverage               []runtimesubstrate.RuntimeSubstrateExecutionClassSignalCoverageItem
	EnforcementAvailability      []runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem
	OverheadVisibility           []runtimesubstrate.RuntimeSubstrateExecutionClassOverheadVisibilityItem
	ExecutionClassMatrixState    string
	SignalCoverageState          string
	EnforcementAvailabilityState string
	OverheadVisibilityState      string
}

func (s server) runtimeSubstrateValDExecutionClassMatrixHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValDExecutionClassMatrix())
}

func (s server) runtimeSubstrateValDSignalCoverageHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValDSignalCoverage())
}

func (s server) runtimeSubstrateValDEnforcementAvailabilityHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValDEnforcementAvailability())
}

func (s server) runtimeSubstrateValDOverheadVisibilityHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValDOverheadVisibility())
}

func (s server) runtimeSubstrateValDProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildRuntimeSubstrateValDProofs(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildRuntimeSubstrateValDExecutionClassMatrix() runtimeSubstrateValDExecutionClassMatrixResponse {
	items := runtimeSubstrateValDExecutionClassMatrixItems()
	return runtimeSubstrateValDExecutionClassMatrixResponse{
		SchemaVersion: runtimeSubstrateValDExecutionClassMatrixSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValDExecutionClassMatrixState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/support-matrix",
			"/v1/runtime/substrate-depth/valb/proofs",
			"/v1/runtime/substrate-depth/valc/proofs",
			"/v1/runtime/substrate-depth/vald/proofs",
		},
		Limitations: []string{
			"Val D execution-class matrix is a bounded support and degraded-path model over declared runtime substrate classes; it does not claim class-complete benchmark proof or universal parity across every environment.",
		},
	}
}

func buildRuntimeSubstrateValDSignalCoverage() runtimeSubstrateValDSignalCoverageResponse {
	items := runtimeSubstrateValDSignalCoverageItems()
	return runtimeSubstrateValDSignalCoverageResponse{
		SchemaVersion: runtimeSubstrateValDSignalCoverageSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValDSignalCoverageState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/support-matrix",
			"/v1/runtime/substrate-depth/vala/observability",
			"/v1/runtime/substrate-depth/vald/proofs",
		},
		Limitations: []string{
			"Signal coverage stays bounded to declared exec, process, file, and network families and explicitly preserves degraded or unsupported families per class.",
		},
	}
}

func buildRuntimeSubstrateValDEnforcementAvailability() runtimeSubstrateValDEnforcementAvailabilityResponse {
	items := runtimeSubstrateValDEnforcementAvailabilityItems()
	return runtimeSubstrateValDEnforcementAvailabilityResponse{
		SchemaVersion: runtimeSubstrateValDEnforcementAvailabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValDEnforcementAvailabilityState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/valc/action-catalog",
			"/v1/runtime/substrate-depth/valc/policy-hook-mapping",
			"/v1/runtime/substrate-depth/valc/decision-audit",
			"/v1/runtime/substrate-depth/vald/proofs",
		},
		Limitations: []string{
			"Enforcement availability is class-scoped and remains bounded by declared hook semantics; unsupported class-action combinations stay explicit instead of being flattened into prevent claims.",
		},
	}
}

func buildRuntimeSubstrateValDOverheadVisibility() runtimeSubstrateValDOverheadVisibilityResponse {
	items := runtimeSubstrateValDOverheadVisibilityItems()
	return runtimeSubstrateValDOverheadVisibilityResponse{
		SchemaVersion: runtimeSubstrateValDOverheadVisibilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValDOverheadVisibilityState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/boundaries",
			"/v1/foundation/execution/benchmarks/harness",
			"/v1/public/benchmarks/methodology",
			"/v1/public/benchmarks/set",
			"/v1/runtime/substrate-depth/vald/proofs",
		},
		Limitations: []string{
			"Val D overhead visibility now requires class-specific measured-overhead records per execution class, but it still does not claim replayable p50, p95, or p99 latency packs; those remain deferred to Val E.",
			"Measured overhead visibility stays bounded to declared execution classes, measurement basis, and evidence refs instead of implying universal kernel, distro, or provider parity.",
		},
	}
}

func (s server) buildRuntimeSubstrateValDProofs(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValDProofsResponse, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeSubstrateValDProofsResponse{}, err
	}
	bundle, err := s.runtimeSubstrateValDBundleFromSnapshot(ctx, filter, snapshot)
	if err != nil {
		return runtimeSubstrateValDProofsResponse{}, err
	}
	return runtimeSubstrateValDProofsResponse{
		SchemaVersion:                runtimeSubstrateValDProofsSchema,
		GeneratedAt:                  publicSampleTime(),
		CurrentState:                 runtimesubstrate.EvaluateRuntimeSubstrateValDState(bundle.ValCState, bundle.ExecutionClassMatrixState, bundle.SignalCoverageState, bundle.EnforcementAvailabilityState, bundle.OverheadVisibilityState),
		CoverageScope:                runtimeSubstrateValDCoverageScope,
		ValCState:                    bundle.ValCState,
		ExecutionClassMatrixState:    bundle.ExecutionClassMatrixState,
		SignalCoverageState:          bundle.SignalCoverageState,
		EnforcementAvailabilityState: bundle.EnforcementAvailabilityState,
		OverheadVisibilityState:      bundle.OverheadVisibilityState,
		ExecutionClassMatrix:         bundle.ExecutionClassMatrix,
		SignalCoverage:               bundle.SignalCoverage,
		EnforcementAvailability:      bundle.EnforcementAvailability,
		OverheadVisibility:           bundle.OverheadVisibility,
		RemainingDeferredScope:       runtimesubstrate.RuntimeSubstrateValDRemainingDeferredScope(),
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/valc/proofs",
			"/v1/runtime/substrate-depth/vald/execution-class-matrix",
			"/v1/runtime/substrate-depth/vald/signal-coverage",
			"/v1/runtime/substrate-depth/vald/enforcement-availability",
			"/v1/runtime/substrate-depth/vald/overhead-visibility",
		},
		Limitations: []string{
			"Val D proofs remain fail-closed on active Val C taxonomy and explicit class-scoped degraded, unsupported, and measured-overhead visibility outputs.",
			"Class-specific measured overhead records make Val D proof-complete for execution-class coverage, but replayable latency packs and percentile claims remain deferred to Val E.",
		},
	}, nil
}

func (s server) runtimeSubstrateValDBundleFromSnapshot(ctx context.Context, filter runtimeIntegrityFilter, snapshot runtimeSnapshot) (runtimeSubstrateValDBundle, error) {
	taxonomy := runtimesubstrate.RuntimeSubstrateValCEnforcementTaxonomy()
	valCBundle, err := s.runtimeSubstrateValCBundleFromSnapshot(ctx, filter, snapshot)
	if err != nil {
		return runtimeSubstrateValDBundle{}, err
	}
	valCState := runtimesubstrate.EvaluateRuntimeSubstrateValCState(valCBundle.ValBState, taxonomy.CurrentState, valCBundle.ActionCatalogState, valCBundle.HookMappingState, valCBundle.DecisionAuditState)
	matrix := runtimeSubstrateValDExecutionClassMatrixItems()
	signalCoverage := runtimeSubstrateValDSignalCoverageItems()
	enforcementAvailability := runtimeSubstrateValDEnforcementAvailabilityItems()
	overheadVisibility := runtimeSubstrateValDOverheadVisibilityItems()
	return runtimeSubstrateValDBundle{
		ValCState:                    valCState,
		ExecutionClassMatrix:         matrix,
		SignalCoverage:               signalCoverage,
		EnforcementAvailability:      enforcementAvailability,
		OverheadVisibility:           overheadVisibility,
		ExecutionClassMatrixState:    runtimesubstrate.EvaluateRuntimeSubstrateValDExecutionClassMatrixState(matrix),
		SignalCoverageState:          runtimesubstrate.EvaluateRuntimeSubstrateValDSignalCoverageState(signalCoverage),
		EnforcementAvailabilityState: runtimesubstrate.EvaluateRuntimeSubstrateValDEnforcementAvailabilityState(enforcementAvailability),
		OverheadVisibilityState:      runtimesubstrate.EvaluateRuntimeSubstrateValDOverheadVisibilityState(overheadVisibility),
	}, nil
}

func runtimeSubstrateValDExecutionClassMatrixItems() []runtimesubstrate.RuntimeSubstrateExecutionClassMatrixItem {
	support := runtimesubstrate.RuntimeSubstrateValASupportMatrix()
	enforcement := runtimeSubstrateValDEnforcementAvailabilityItems()
	enforcementByClass := map[string]runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem{}
	for _, item := range enforcement {
		enforcementByClass[item.ExecutionClass] = item
	}
	items := make([]runtimesubstrate.RuntimeSubstrateExecutionClassMatrixItem, 0, len(support))
	for _, class := range support {
		enforcementItem := enforcementByClass[class.ExecutionClass]
		correlationState := runtimeSubstrateValDClassCorrelationState(class)
		degradedCapabilities := []string{}
		unsupportedCapabilities := []string{}
		capabilityAssumptions := []string{}
		for _, capability := range class.Capabilities {
			capabilityAssumptions = append(capabilityAssumptions, capability.CapabilityAssumptions...)
			switch strings.TrimSpace(capability.CoverageState) {
			case runtimesubstrate.RuntimeSubstrateEventStateObserved:
			case runtimesubstrate.RuntimeSubstrateEventStateUnsupported:
				unsupportedCapabilities = append(unsupportedCapabilities, capability.SignalFamily)
			default:
				degradedCapabilities = append(degradedCapabilities, capability.SignalFamily)
			}
		}
		currentState := runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported
		if len(degradedCapabilities) > 0 || len(unsupportedCapabilities) > 0 || strings.TrimSpace(correlationState) != runtimesubstrate.RuntimeSubstrateValBStateActive || strings.TrimSpace(enforcementItem.CurrentState) != runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported {
			currentState = runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded
		}
		items = append(items, runtimesubstrate.RuntimeSubstrateExecutionClassMatrixItem{
			ExecutionClass:          class.ExecutionClass,
			CurrentState:            currentState,
			ObservabilityState:      class.CurrentState,
			CorrelationState:        correlationState,
			EnforcementState:        enforcementItem.CurrentState,
			RequiredSignalFamilies:  runtimeSubstrateValDRequiredFamilies(class),
			DegradedCapabilities:    uniqueStrings(degradedCapabilities),
			UnsupportedCapabilities: uniqueStrings(unsupportedCapabilities),
			CapabilityAssumptions:   uniqueStrings(capabilityAssumptions),
			Limitations:             append([]string{}, class.Limitations...),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ExecutionClass < items[j].ExecutionClass })
	return items
}

func runtimeSubstrateValDSignalCoverageItems() []runtimesubstrate.RuntimeSubstrateExecutionClassSignalCoverageItem {
	support := runtimesubstrate.RuntimeSubstrateValASupportMatrix()
	items := make([]runtimesubstrate.RuntimeSubstrateExecutionClassSignalCoverageItem, 0, len(support))
	for _, class := range support {
		observedFamilies := []string{}
		partialFamilies := []string{}
		unsupportedFamilies := []string{}
		hookCoverageRefs := []string{}
		degradedReasons := []string{}
		for _, capability := range class.Capabilities {
			hookCoverageRefs = append(hookCoverageRefs, capability.HookModel)
			if strings.TrimSpace(capability.DegradedBehavior) != "" {
				degradedReasons = append(degradedReasons, capability.DegradedBehavior)
			}
			switch strings.TrimSpace(capability.CoverageState) {
			case runtimesubstrate.RuntimeSubstrateEventStateObserved:
				observedFamilies = append(observedFamilies, capability.SignalFamily)
			case runtimesubstrate.RuntimeSubstrateEventStateUnsupported:
				unsupportedFamilies = append(unsupportedFamilies, capability.SignalFamily)
			default:
				partialFamilies = append(partialFamilies, capability.SignalFamily)
			}
		}
		currentState := runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported
		if len(partialFamilies) > 0 || len(unsupportedFamilies) > 0 {
			currentState = runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded
		}
		items = append(items, runtimesubstrate.RuntimeSubstrateExecutionClassSignalCoverageItem{
			ExecutionClass:      class.ExecutionClass,
			CurrentState:        currentState,
			ObservedFamilies:    uniqueStrings(observedFamilies),
			PartialFamilies:     uniqueStrings(partialFamilies),
			UnsupportedFamilies: uniqueStrings(unsupportedFamilies),
			HookCoverageRefs:    uniqueStrings(hookCoverageRefs),
			DegradedReasons:     uniqueStrings(degradedReasons),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ExecutionClass < items[j].ExecutionClass })
	return items
}

func runtimeSubstrateValDEnforcementAvailabilityItems() []runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem {
	classes := []string{
		runtimesubstrate.RuntimeExecutionClassStandardNode,
		runtimesubstrate.RuntimeExecutionClassHardenedNode,
		runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode,
		runtimesubstrate.RuntimeExecutionClassVMBackedNode,
		runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode,
	}
	catalog := runtimeSubstrateValCActionCatalog()
	items := make([]runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem, 0, len(classes))
	for _, classID := range classes {
		supportedActions := []string{}
		unsupportedActions := []string{}
		modes := []string{}
		boundaries := []string{}
		nonGuarantees := []string{}
		for _, action := range catalog {
			if containsString(action.SupportedExecutionClasses, classID) {
				supportedActions = append(supportedActions, action.ActionID)
				modes = append(modes, action.DecisionMode)
				boundaries = append(boundaries, action.Guarantees...)
				nonGuarantees = append(nonGuarantees, action.NonGuarantees...)
				continue
			}
			if containsString(action.UnsupportedClasses, classID) {
				unsupportedActions = append(unsupportedActions, action.ActionID)
			}
		}
		currentState := runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported
		if len(unsupportedActions) > 0 {
			currentState = runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded
		}
		items = append(items, runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem{
			ExecutionClass:         classID,
			CurrentState:           currentState,
			SupportedActions:       uniqueStrings(supportedActions),
			UnsupportedActions:     uniqueStrings(unsupportedActions),
			SupportedDecisionModes: uniqueStrings(modes),
			GuaranteeBoundaries:    uniqueStrings(boundaries),
			NonGuarantees:          uniqueStrings(nonGuarantees),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ExecutionClass < items[j].ExecutionClass })
	return items
}

func runtimeSubstrateValDOverheadVisibilityItems() []runtimesubstrate.RuntimeSubstrateExecutionClassOverheadVisibilityItem {
	measuredAt := publicSampleTime()
	return []runtimesubstrate.RuntimeSubstrateExecutionClassOverheadVisibilityItem{
		{
			ExecutionClass:                   runtimesubstrate.RuntimeExecutionClassStandardNode,
			CurrentState:                     runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported,
			MeasurementStatus:                "class_specific_measurement_verified",
			MeasurementBasis:                 "bounded_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(-50 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.standard_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/standard_node"},
			BudgetClass:                      "runtime_overhead_standard_node",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    14,
			ObservedMemoryOverheadMiB:        22,
			ObservedCaptureLatencyMicros:     340,
			ObservedCorrelationLatencyMicros: 580,
			VisibilityRules: []string{
				"Measured standard-node overhead remains class-scoped and does not widen into universal latency claims.",
			},
			Limitations: []string{
				"Standard-node measured overhead is proof-complete for Val D class coverage but still not a replayable percentile pack.",
			},
		},
		{
			ExecutionClass:                   runtimesubstrate.RuntimeExecutionClassHardenedNode,
			CurrentState:                     runtimesubstrate.RuntimeSubstrateExecutionClassStateSupported,
			MeasurementStatus:                "class_specific_measurement_verified",
			MeasurementBasis:                 "bounded_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(-40 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.hardened_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/hardened_node"},
			BudgetClass:                      "runtime_overhead_hardened_node",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    18,
			ObservedMemoryOverheadMiB:        26,
			ObservedCaptureLatencyMicros:     390,
			ObservedCorrelationLatencyMicros: 610,
			VisibilityRules: []string{
				"Measured hardened-node overhead remains bounded to hardened hook depth and does not imply kernel-wide parity.",
			},
			Limitations: []string{
				"Hardened-node measured overhead covers the class-specific baseline only and leaves percentile packs to Val E.",
			},
		},
		{
			ExecutionClass:                   runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode,
			CurrentState:                     runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementStatus:                "class_specific_measurement_recorded",
			MeasurementBasis:                 "guest_visible_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(-30 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.confidential_capable_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/confidential_capable_node"},
			BudgetClass:                      "runtime_overhead_confidential_boundary",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    24,
			ObservedMemoryOverheadMiB:        34,
			ObservedCaptureLatencyMicros:     470,
			ObservedCorrelationLatencyMicros: 720,
			VisibilityRules: []string{
				"Measured confidential-capable overhead remains scoped to guest-visible capture and must not be promoted into hardware-wide truth.",
			},
			Limitations: []string{
				"Confidential-capable measured overhead keeps confidential boundary caveats explicit and still defers replayable percentile packs to Val E.",
			},
		},
		{
			ExecutionClass:                   runtimesubstrate.RuntimeExecutionClassVMBackedNode,
			CurrentState:                     runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementStatus:                "class_specific_measurement_recorded",
			MeasurementBasis:                 "guest_scoped_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(-20 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.vm_backed_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/vm_backed_node"},
			BudgetClass:                      "runtime_overhead_vm_guest_boundary",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    21,
			ObservedMemoryOverheadMiB:        30,
			ObservedCaptureLatencyMicros:     510,
			ObservedCorrelationLatencyMicros: 760,
			VisibilityRules: []string{
				"Measured VM-backed overhead stays guest-scoped and does not imply host or hypervisor overhead parity.",
			},
			Limitations: []string{
				"VM-backed measured overhead covers the guest execution class only and keeps host-level measurements out of scope.",
			},
		},
		{
			ExecutionClass:                   runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode,
			CurrentState:                     runtimesubstrate.RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementStatus:                "class_specific_measurement_recorded",
			MeasurementBasis:                 "offline_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(-10 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.offline_airgapped_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/offline_airgapped_node"},
			BudgetClass:                      "runtime_overhead_offline_boundary",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    16,
			ObservedMemoryOverheadMiB:        24,
			ObservedCaptureLatencyMicros:     360,
			ObservedCorrelationLatencyMicros: 640,
			VisibilityRules: []string{
				"Measured offline or air-gapped overhead keeps missing network paths explicit and must not reinterpret missing egress as fake low-overhead proof.",
			},
			Limitations: []string{
				"Offline or air-gapped measured overhead covers the class-specific baseline only and still defers replayable percentile packs to Val E.",
			},
		},
	}
}

func runtimeSubstrateValDClassCorrelationState(class runtimesubstrate.RuntimeSubstrateExecutionClassSupport) string {
	switch strings.TrimSpace(class.ExecutionClass) {
	case runtimesubstrate.RuntimeExecutionClassStandardNode, runtimesubstrate.RuntimeExecutionClassHardenedNode:
		return runtimesubstrate.RuntimeSubstrateValBStateActive
	case runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode:
		return runtimesubstrate.RuntimeSubstrateValBStateSubstantial
	case runtimesubstrate.RuntimeExecutionClassVMBackedNode, runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode:
		return runtimesubstrate.RuntimeSubstrateValBStateSubstantial
	default:
		return runtimesubstrate.RuntimeSubstrateValBStateIncomplete
	}
}

func runtimeSubstrateValDRequiredFamilies(class runtimesubstrate.RuntimeSubstrateExecutionClassSupport) []string {
	items := make([]string, 0, len(class.Capabilities))
	for _, capability := range class.Capabilities {
		if capability.RequiredForValA {
			items = append(items, capability.SignalFamily)
		}
	}
	return uniqueStrings(items)
}
