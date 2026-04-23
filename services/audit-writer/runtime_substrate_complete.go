package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

const runtimeSubstratePoint1CompleteSchema = "runtime.substrate.point1.complete.v1"

type runtimeSubstratePoint1CompleteResponse struct {
	SchemaVersion      string    `json:"schema_version"`
	GeneratedAt        time.Time `json:"generated_at"`
	CurrentState       string    `json:"current_state"`
	Point1State        string    `json:"point_1_state"`
	ValAState          string    `json:"val_a_state"`
	ValBState          string    `json:"val_b_state"`
	ValCState          string    `json:"val_c_state"`
	ValDState          string    `json:"val_d_state"`
	ValEState          string    `json:"val_e_state"`
	IntegrationSummary []string  `json:"integration_summary,omitempty"`
	SurfaceRefs        []string  `json:"surface_refs,omitempty"`
	EvidenceRefs       []string  `json:"evidence_refs,omitempty"`
	DocumentationRefs  []string  `json:"documentation_refs,omitempty"`
	DeferredScope      []string  `json:"deferred_scope,omitempty"`
	Limitations        []string  `json:"limitations,omitempty"`
}

func (s server) runtimeSubstratePoint1CompleteHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildRuntimeSubstratePoint1Complete(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildRuntimeSubstratePoint1Complete(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstratePoint1CompleteResponse, error) {
	valA, err := s.buildRuntimeSubstrateValAProofs(ctx, filter)
	if err != nil {
		return runtimeSubstratePoint1CompleteResponse{}, err
	}
	valB, err := s.buildRuntimeSubstrateValBProofs(ctx, filter)
	if err != nil {
		return runtimeSubstratePoint1CompleteResponse{}, err
	}
	valC, err := s.buildRuntimeSubstrateValCProofs(ctx, filter)
	if err != nil {
		return runtimeSubstratePoint1CompleteResponse{}, err
	}
	valD, err := s.buildRuntimeSubstrateValDProofs(ctx, filter)
	if err != nil {
		return runtimeSubstratePoint1CompleteResponse{}, err
	}
	valE, err := s.buildRuntimeSubstrateValEProofs(ctx, filter)
	if err != nil {
		return runtimeSubstratePoint1CompleteResponse{}, err
	}

	documentationRefs := runtimesubstrate.RuntimeSubstratePoint1DocumentationRefs()
	surfaceRefs := uniqueStrings(append(runtimesubstrate.RuntimeSubstratePoint1SurfaceRefs(), "/v1/runtime/substrate-depth/complete"))
	evidenceRefs := uniqueStrings(runtimesubstrate.RuntimeSubstratePoint1EvidenceRefs())
	deferredScope := runtimesubstrate.RuntimeSubstratePoint1DeferredScope()
	integrationSummary := runtimeSubstratePoint1IntegrationSummary()
	limitations := runtimeSubstratePoint1Limitations()
	point1State := runtimesubstrate.EvaluateRuntimeSubstratePoint1State(
		valA.CurrentState,
		valB.CurrentState,
		valC.CurrentState,
		valD.CurrentState,
		valE.CurrentState,
		documentationRefs,
		surfaceRefs,
		evidenceRefs,
		limitations,
		deferredScope,
		strings.Join(integrationSummary, " "),
	)

	return runtimeSubstratePoint1CompleteResponse{
		SchemaVersion:      runtimeSubstratePoint1CompleteSchema,
		GeneratedAt:        publicSampleTime(),
		CurrentState:       point1State,
		Point1State:        point1State,
		ValAState:          valA.CurrentState,
		ValBState:          valB.CurrentState,
		ValCState:          valC.CurrentState,
		ValDState:          valD.CurrentState,
		ValEState:          valE.CurrentState,
		IntegrationSummary: integrationSummary,
		SurfaceRefs:        surfaceRefs,
		EvidenceRefs:       evidenceRefs,
		DocumentationRefs:  documentationRefs,
		DeferredScope:      deferredScope,
		Limitations:        limitations,
	}, nil
}

func runtimeSubstratePoint1IntegrationSummary() []string {
	return []string{
		"Val A established bounded substrate observability with explicit observed, partial, stale, and unsupported state handling.",
		"Val B added evidence-traceable process-image and provenance correlation with explicit drift classes.",
		"Val C added bounded enforcement taxonomy, policy-hook mapping, and canonical decision audit semantics.",
		"Val D closed execution-class proof completeness by combining support matrix, signal coverage, enforcement availability, and measured overhead visibility.",
		"Val E converted performance into a gate with measured latency packs, false-positive budgets, replayable benchmark packs, and passed performance gates.",
		"Point 1 integrated closure binds those five accepted slices into one fail-closed runtime/substrate completion summary without adding new runtime mechanics.",
	}
}

func runtimeSubstratePoint1Limitations() []string {
	return []string{
		"Point 1 integrated closure is a summary and lock layer over Val A through Val E; it does not add new runtime capture, provenance, enforcement, execution-class, or benchmark mechanics.",
		"Point 1 pass remains bounded to internal runtime/substrate proof completeness and does not automatically become a public benchmark or public proof publication layer.",
	}
}
