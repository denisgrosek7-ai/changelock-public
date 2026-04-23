package runtime

import "strings"

const (
	RuntimeSubstratePoint1StateIncomplete  = "runtime_substrate_point_1_incomplete"
	RuntimeSubstratePoint1StateSubstantial = "runtime_substrate_point_1_substantially_ready"
	RuntimeSubstratePoint1StateActive      = "runtime_substrate_point_1_active"
)

func RuntimeSubstratePoint1DocumentationRefs() []string {
	return []string{
		"docs/runtime-substrate-depth-vala-core.md",
		"docs/runtime-substrate-depth-valb-core.md",
		"docs/runtime-substrate-depth-valc-core.md",
		"docs/runtime-substrate-depth-vald-core.md",
		"docs/runtime-substrate-depth-vale-core.md",
		"docs/runtime-substrate-depth-complete.md",
	}
}

func RuntimeSubstratePoint1SurfaceRefs() []string {
	return []string{
		"/v1/runtime/substrate-depth/vala/proofs",
		"/v1/runtime/substrate-depth/valb/proofs",
		"/v1/runtime/substrate-depth/valc/proofs",
		"/v1/runtime/substrate-depth/vald/proofs",
		"/v1/runtime/substrate-depth/vale/proofs",
		"/v1/runtime/substrate-depth/complete",
	}
}

func RuntimeSubstratePoint1EvidenceRefs() []string {
	return []string{
		"/v1/runtime/substrate-depth/vala/proofs",
		"/v1/runtime/substrate-depth/valb/proofs",
		"/v1/runtime/substrate-depth/valc/proofs",
		"/v1/runtime/substrate-depth/vald/proofs",
		"/v1/runtime/substrate-depth/vale/proofs",
		"/v1/foundation/execution/benchmarks/evaluate",
	}
}

func RuntimeSubstratePoint1DeferredScope() []string {
	return []string{
		"point_2_public_benchmark_publication_and_percentile_packaging",
		"point_2_customer_safe_external_proof_and_narrative_publication",
		"point_2_broader_public_claim_governance_beyond_internal_runtime_proof_closure",
	}
}

func EvaluateRuntimeSubstratePoint1State(valAState, valBState, valCState, valDState, valEState string, documentationRefs, surfaceRefs, evidenceRefs, limitations, deferredScope []string, integrationSummary string) string {
	if !runtimeSubstratePoint1MetadataPresent(documentationRefs, surfaceRefs, evidenceRefs, limitations, deferredScope, integrationSummary) {
		if runtimeSubstratePoint1AllValsActive(valAState, valBState, valCState, valDState, valEState) {
			return RuntimeSubstratePoint1StateSubstantial
		}
		return RuntimeSubstratePoint1StateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(valAState),
		strings.TrimSpace(valBState),
		strings.TrimSpace(valCState),
		strings.TrimSpace(valDState),
		strings.TrimSpace(valEState),
	} {
		switch state {
		case RuntimeSubstrateValAStateActive,
			RuntimeSubstrateValBStateActive,
			RuntimeSubstrateValCStateActive,
			RuntimeSubstrateValDStateActive,
			RuntimeSubstrateValEStateActive:
		case RuntimeSubstrateValAStateSubstantial,
			RuntimeSubstrateValAStateContractDefined,
			RuntimeSubstrateValBStateSubstantial,
			RuntimeSubstrateValCStateSubstantial,
			RuntimeSubstrateValDStateSubstantial,
			RuntimeSubstrateValEStateSubstantial:
			hasPartial = true
		default:
			return RuntimeSubstratePoint1StateIncomplete
		}
	}
	if hasPartial {
		return RuntimeSubstratePoint1StateSubstantial
	}
	return RuntimeSubstratePoint1StateActive
}

func runtimeSubstratePoint1MetadataPresent(documentationRefs, surfaceRefs, evidenceRefs, limitations, deferredScope []string, integrationSummary string) bool {
	return len(documentationRefs) >= 6 &&
		len(surfaceRefs) >= 6 &&
		len(evidenceRefs) >= 6 &&
		len(limitations) > 0 &&
		len(deferredScope) > 0 &&
		strings.TrimSpace(integrationSummary) != ""
}

func runtimeSubstratePoint1AllValsActive(valAState, valBState, valCState, valDState, valEState string) bool {
	return strings.TrimSpace(valAState) == RuntimeSubstrateValAStateActive &&
		strings.TrimSpace(valBState) == RuntimeSubstrateValBStateActive &&
		strings.TrimSpace(valCState) == RuntimeSubstrateValCStateActive &&
		strings.TrimSpace(valDState) == RuntimeSubstrateValDStateActive &&
		strings.TrimSpace(valEState) == RuntimeSubstrateValEStateActive
}
