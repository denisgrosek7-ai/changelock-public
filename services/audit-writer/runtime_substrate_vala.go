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

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	runtimeSubstrateDepthEntryGateSchema    = "runtime.substrate.entry_gate.v1"
	runtimeSubstrateValAEventSchemaVersion  = "runtime.substrate.vala.event_schema.v1"
	runtimeSubstrateValASupportMatrixSchema = "runtime.substrate.vala.support_matrix.v1"
	runtimeSubstrateValAObservabilitySchema = "runtime.substrate.vala.observability.v1"
	runtimeSubstrateValAProofsSchema        = "runtime.substrate.vala.proofs.v1"
	runtimeSubstrateValACoverageScope       = "substrate_observability_baseline"
	runtimeSubstrateValAStatusRecorded      = "recorded"
)

type runtimeSubstrateEntryGateResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	EntryGate     runtimesubstrate.RuntimeSubstrateEntryGate `json:"entry_gate"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type runtimeSubstrateEventSchemaResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	CurrentState  string                                       `json:"current_state"`
	EventSchema   runtimesubstrate.RuntimeSubstrateEventSchema `json:"event_schema"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type runtimeSubstrateSupportMatrixResponse struct {
	SchemaVersion       string                                                   `json:"schema_version"`
	GeneratedAt         time.Time                                                `json:"generated_at"`
	CurrentState        string                                                   `json:"current_state"`
	ImplementationState string                                                   `json:"implementation_state"`
	ExecutionClasses    []runtimesubstrate.RuntimeSubstrateExecutionClassSupport `json:"execution_classes,omitempty"`
	RouteRefs           []string                                                 `json:"route_refs,omitempty"`
	Limitations         []string                                                 `json:"limitations,omitempty"`
}

type runtimeSubstrateObservabilityResponse struct {
	SchemaVersion               string                                           `json:"schema_version"`
	GeneratedAt                 time.Time                                        `json:"generated_at"`
	CurrentState                string                                           `json:"current_state"`
	ImplementationState         string                                           `json:"implementation_state"`
	RecordKind                  string                                           `json:"record_kind"`
	NotRuntimeTruthSurface      bool                                             `json:"not_runtime_truth_surface"`
	ExecAttributionState        string                                           `json:"exec_attribution_state"`
	ProcessLineageState         string                                           `json:"process_lineage_state"`
	FileAttributionState        string                                           `json:"file_attribution_state"`
	NetworkAttributionState     string                                           `json:"network_attribution_state"`
	WorkloadNodeEnrichmentState string                                           `json:"workload_node_enrichment_state"`
	Items                       []runtimesubstrate.RuntimeSubstrateObservedEvent `json:"items,omitempty"`
	RouteRefs                   []string                                         `json:"route_refs,omitempty"`
	Limitations                 []string                                         `json:"limitations,omitempty"`
}

type runtimeSubstrateValAProofsResponse struct {
	SchemaVersion          string                                           `json:"schema_version"`
	GeneratedAt            time.Time                                        `json:"generated_at"`
	CurrentState           string                                           `json:"current_state"`
	ImplementationState    string                                           `json:"implementation_state"`
	CoverageScope          string                                           `json:"coverage_scope"`
	EntryGateState         string                                           `json:"entry_gate_state"`
	EventSchemaState       string                                           `json:"event_schema_state"`
	SupportMatrixState     string                                           `json:"support_matrix_state"`
	ObservabilityState     string                                           `json:"observability_state"`
	ObservedEvents         []runtimesubstrate.RuntimeSubstrateObservedEvent `json:"observed_events,omitempty"`
	RemainingDeferredScope []string                                         `json:"remaining_deferred_scope,omitempty"`
	RouteRefs              []string                                         `json:"route_refs,omitempty"`
	Limitations            []string                                         `json:"limitations,omitempty"`
}

type runtimeSubstrateValAObservationWriteRequest struct {
	Event runtimesubstrate.RuntimeSubstrateObservedEvent `json:"event"`
}

type runtimeSubstrateValAObservationWriteResponse struct {
	Status string                                         `json:"status"`
	Event  runtimesubstrate.RuntimeSubstrateObservedEvent `json:"event"`
}

type runtimeSubstrateValAObservationPayload struct {
	SchemaVersion         string                            `json:"schema_version"`
	EventID               string                            `json:"event_id,omitempty"`
	EventFamily           string                            `json:"event_family"`
	CurrentState          string                            `json:"current_state"`
	Process               runtimesubstrate.ProcessIdentity  `json:"process"`
	ParentPID             int                               `json:"parent_pid,omitempty"`
	ThreadID              int                               `json:"thread_id,omitempty"`
	ContainerRuntime      string                            `json:"container_runtime,omitempty"`
	NamespaceMode         string                            `json:"namespace_mode,omitempty"`
	Workload              runtimesubstrate.WorkloadIdentity `json:"workload"`
	Node                  runtimesubstrate.NodeIdentity     `json:"node"`
	AttributionConfidence string                            `json:"attribution_confidence"`
	FreshnessState        string                            `json:"freshness_state"`
	UnsupportedFields     []string                          `json:"unsupported_fields,omitempty"`
	Reasons               []string                          `json:"reasons,omitempty"`
}

func (s server) runtimeSubstrateDepthEntryGateHandler(w http.ResponseWriter, r *http.Request) {
	r, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateEntryGate())
}

func (s server) runtimeSubstrateValAEventSchemaHandler(w http.ResponseWriter, r *http.Request) {
	r, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValAEventSchema())
}

func (s server) runtimeSubstrateValASupportMatrixHandler(w http.ResponseWriter, r *http.Request) {
	r, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValASupportMatrix())
}

func (s server) runtimeSubstrateValAObservabilityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		req, ok := s.authorizeRuntimeSubstrateRead(w, r)
		if !ok {
			return
		}
		filter, err := parseRuntimeIntegrityFilter(req)
		if err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
		defer cancel()
		response, err := s.buildRuntimeSubstrateValAObservability(ctx, filter)
		if err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, response)
	case http.MethodPost:
		principal, req, ok := s.authorizeRuntimeSubstrateWrite(w, r)
		if !ok {
			return
		}
		filter, err := parseRuntimeIntegrityFilter(req)
		if err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		var request runtimeSubstrateValAObservationWriteRequest
		if err := httpjson.Decode(req, &request); err != nil && !errors.Is(err, io.EOF) {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		event, err := normalizeRuntimeSubstrateValAObservation(filter, request.Event)
		if err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistRuntimeSubstrateValAObservation(ctx, principal.Subject, filter, event); err != nil {
			writeRuntimeIntegrityError(w, err)
			return
		}
		httpjson.Write(w, http.StatusCreated, runtimeSubstrateValAObservationWriteResponse{
			Status: runtimeSubstrateValAStatusRecorded,
			Event:  event,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) runtimeSubstrateValAProofsHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if r.Method != http.MethodGet {
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
	response, err := s.buildRuntimeSubstrateValAProofs(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) authorizeRuntimeSubstrateRead(w http.ResponseWriter, r *http.Request) (*http.Request, bool) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return nil, false
	}
	req, err := applyPrincipalTenantToRequest(principal, authorizedRequest)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return nil, false
	}
	return req, true
}

func (s server) authorizeRuntimeSubstrateWrite(w http.ResponseWriter, r *http.Request) (auth.Principal, *http.Request, bool) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return auth.Principal{}, nil, false
	}
	req, err := applyPrincipalTenantToRequest(principal, authorizedRequest)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return auth.Principal{}, nil, false
	}
	return principal, req, true
}

func buildRuntimeSubstrateEntryGate() runtimeSubstrateEntryGateResponse {
	entryGate := runtimesubstrate.RuntimeSubstrateEntryGateRuntimeBaseline()
	return runtimeSubstrateEntryGateResponse{
		SchemaVersion: runtimeSubstrateDepthEntryGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  entryGate.CurrentState,
		EntryGate:     entryGate,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/event-schema",
			"/v1/runtime/substrate-depth/vala/observability",
			"/v1/runtime/substrate-depth/vala/proofs",
		},
		Limitations: []string{
			"Entry gate now reflects a bounded Val A runtime baseline only; provenance correlation, enforcement taxonomy, and benchmark waves remain deferred.",
		},
	}
}

func buildRuntimeSubstrateValAEventSchema() runtimeSubstrateEventSchemaResponse {
	schema := runtimesubstrate.RuntimeSubstrateValAEventSchema()
	return runtimeSubstrateEventSchemaResponse{
		SchemaVersion: runtimeSubstrateValAEventSchemaVersion,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  schema.CurrentState,
		EventSchema:   schema,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/entry-gate",
			"/v1/runtime/substrate-depth/vala/observability",
			"/v1/runtime/substrate-depth/vala/proofs",
		},
		Limitations: []string{
			"Val A event schema remains bounded to runtime observability and does not yet bind runtime events to signed image provenance or attestation-backed binary truth.",
		},
	}
}

func buildRuntimeSubstrateValASupportMatrix() runtimeSubstrateSupportMatrixResponse {
	matrix := runtimesubstrate.RuntimeSubstrateValASupportMatrix()
	return runtimeSubstrateSupportMatrixResponse{
		SchemaVersion:       runtimeSubstrateValASupportMatrixSchema,
		GeneratedAt:         publicSampleTime(),
		CurrentState:        runtimesubstrate.EvaluateRuntimeSubstrateValASupportMatrixState(matrix),
		ImplementationState: runtimesubstrate.RuntimeSubstrateImplementationStateRuntimePath,
		ExecutionClasses:    matrix,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/event-schema",
			"/v1/runtime/substrate-depth/vala/observability",
			"/v1/runtime/substrate-depth/vala/proofs",
		},
		Limitations: []string{
			"Support matrix now backs a minimal runtime path, but execution-class coverage still remains bounded to declared hook assumptions and degraded behaviors in Val A.",
		},
	}
}

func (s server) buildRuntimeSubstrateValAObservability(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateObservabilityResponse, error) {
	items, err := s.listRuntimeSubstrateValAObservedEvents(ctx, filter)
	if err != nil {
		return runtimeSubstrateObservabilityResponse{}, err
	}
	return runtimeSubstrateObservabilityResponse{
		SchemaVersion:               runtimeSubstrateValAObservabilitySchema,
		GeneratedAt:                 publicSampleTime(),
		CurrentState:                runtimesubstrate.EvaluateRuntimeSubstrateValAObservabilityState(items),
		ImplementationState:         runtimesubstrate.RuntimeSubstrateImplementationStateRuntimePath,
		RecordKind:                  runtimesubstrate.RuntimeSubstrateRecordKindObserved,
		NotRuntimeTruthSurface:      false,
		ExecAttributionState:        runtimeSubstrateFamilyCoverageState(items, runtimesubstrate.RuntimeSubstrateEventFamilyExecLifecycle, "exec_identity_records_active", "exec_identity_records_missing"),
		ProcessLineageState:         runtimeSubstrateFamilyCoverageState(items, runtimesubstrate.RuntimeSubstrateEventFamilyProcessLineage, "process_lineage_records_active", "process_lineage_records_missing"),
		FileAttributionState:        runtimeSubstrateFamilyCoverageState(items, runtimesubstrate.RuntimeSubstrateEventFamilyFileActivity, "file_activity_attribution_active", "file_activity_attribution_missing"),
		NetworkAttributionState:     runtimeSubstrateFamilyCoverageState(items, runtimesubstrate.RuntimeSubstrateEventFamilyNetworkActivity, "network_activity_attribution_active", "network_activity_attribution_missing"),
		WorkloadNodeEnrichmentState: runtimeSubstrateWorkloadNodeCoverageState(items),
		Items:                       items,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/event-schema",
			"/v1/runtime/substrate-depth/vala/support-matrix",
			"/v1/runtime/substrate-depth/vala/proofs",
		},
		Limitations: append([]string{
			"Observed events are loaded from canonical audit-backed runtime observations and do not create a second runtime database outside canonical evidence lineage.",
			"Attribution confidence remains capture and correlation quality only and cannot be reinterpreted as a deny or kill decision signal in Val A.",
		}, runtimeSubstrateEmptyObservabilityLimitation(items)...),
	}, nil
}

func (s server) buildRuntimeSubstrateValAProofs(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValAProofsResponse, error) {
	entryGate := runtimesubstrate.RuntimeSubstrateEntryGateRuntimeBaseline()
	schema := runtimesubstrate.RuntimeSubstrateValAEventSchema()
	matrix := runtimesubstrate.RuntimeSubstrateValASupportMatrix()
	events, err := s.listRuntimeSubstrateValAObservedEvents(ctx, filter)
	if err != nil {
		return runtimeSubstrateValAProofsResponse{}, err
	}
	supportMatrixState := runtimesubstrate.EvaluateRuntimeSubstrateValASupportMatrixState(matrix)
	observabilityState := runtimesubstrate.EvaluateRuntimeSubstrateValAObservabilityState(events)
	return runtimeSubstrateValAProofsResponse{
		SchemaVersion:          runtimeSubstrateValAProofsSchema,
		GeneratedAt:            publicSampleTime(),
		CurrentState:           runtimesubstrate.EvaluateRuntimeSubstrateValAState(entryGate.CurrentState, schema.CurrentState, supportMatrixState, observabilityState),
		ImplementationState:    runtimesubstrate.RuntimeSubstrateImplementationStateRuntimePath,
		CoverageScope:          runtimeSubstrateValACoverageScope,
		EntryGateState:         entryGate.CurrentState,
		EventSchemaState:       schema.CurrentState,
		SupportMatrixState:     supportMatrixState,
		ObservabilityState:     observabilityState,
		ObservedEvents:         events,
		RemainingDeferredScope: runtimesubstrate.RuntimeSubstrateRemainingDeferredScope(),
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/entry-gate",
			"/v1/runtime/substrate-depth/vala/event-schema",
			"/v1/runtime/substrate-depth/vala/support-matrix",
			"/v1/runtime/substrate-depth/vala/observability",
		},
		Limitations: []string{
			"Val A proofs require real canonical runtime observations across all mandatory families and states before the slice becomes active.",
			"Proofs do not claim later-wave provenance correlation, enforcement taxonomy, or benchmark-backed latency budgets.",
		},
	}, nil
}

func normalizeRuntimeSubstrateValAObservation(filter runtimeIntegrityFilter, event runtimesubstrate.RuntimeSubstrateObservedEvent) (runtimesubstrate.RuntimeSubstrateObservedEvent, error) {
	event = runtimesubstrate.NormalizeRuntimeSubstrateObservedEvent(event, time.Now)
	event.Workload.ClusterID = firstNonEmpty(strings.TrimSpace(event.Workload.ClusterID), strings.TrimSpace(filter.ClusterID), "local")
	event.Workload.Namespace = firstNonEmpty(strings.TrimSpace(event.Workload.Namespace), strings.TrimSpace(filter.Namespace))
	event.Workload.WorkloadKind = firstNonEmpty(strings.TrimSpace(event.Workload.WorkloadKind), strings.TrimSpace(filter.WorkloadKind))
	event.Workload.Workload = firstNonEmpty(strings.TrimSpace(event.Workload.Workload), strings.TrimSpace(filter.Workload))
	subjectRef := runtimeSubjectRef(event.Workload.ClusterID, event.Workload.Namespace, event.Workload.WorkloadKind, event.Workload.Workload)
	scope, err := phase2PersistScopeFromFilter(filter, subjectRef)
	if err != nil {
		return runtimesubstrate.RuntimeSubstrateObservedEvent{}, err
	}
	event.Workload.ClusterID = firstNonEmpty(strings.TrimSpace(event.Workload.ClusterID), scope.ClusterID)
	event.Workload.Namespace = firstNonEmpty(strings.TrimSpace(event.Workload.Namespace), scope.Namespace)
	event.Workload.WorkloadKind = firstNonEmpty(strings.TrimSpace(event.Workload.WorkloadKind), scope.WorkloadKind)
	event.Workload.Workload = firstNonEmpty(strings.TrimSpace(event.Workload.Workload), scope.Workload)
	if strings.TrimSpace(event.Workload.PolicySubject) == "" {
		event.Workload.PolicySubject = runtimeSubjectRef(event.Workload.ClusterID, event.Workload.Namespace, event.Workload.WorkloadKind, event.Workload.Workload)
	}
	if err := runtimesubstrate.ValidateRuntimeSubstrateObservedEvent(event); err != nil {
		return runtimesubstrate.RuntimeSubstrateObservedEvent{}, err
	}
	return event, nil
}

func (s server) persistRuntimeSubstrateValAObservation(ctx context.Context, actor string, filter runtimeIntegrityFilter, event runtimesubstrate.RuntimeSubstrateObservedEvent) error {
	scope, err := phase2PersistScopeFromFilter(filter, event.Workload.PolicySubject)
	if err != nil {
		return err
	}
	payload, err := canonicalJSON(runtimeIntegrityEventPayload{
		Observation: &runtimeObservationPayload{
			Node:         event.Node.NodeID,
			Pod:          event.Workload.PodUID,
			EventType:    event.EventFamily,
			EventPayload: runtimeSubstrateValAEventPayload(event),
			Confidence:   runtimeSubstrateConfidenceToRuntimeConfidence(event.AttributionConfidence),
		},
	})
	if err != nil {
		return err
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:        audit.NewRequestID(),
		Component:        runtimeIntegrityComponent,
		EventType:        audit.EventTypeRuntimeObservationRecorded,
		Actor:            strings.TrimSpace(actor),
		ClusterID:        scope.ClusterID,
		TenantID:         scope.TenantID,
		Environment:      scope.Environment,
		Namespace:        scope.Namespace,
		WorkloadKind:     scope.WorkloadKind,
		Workload:         scope.Workload,
		Digest:           event.Workload.ImageDigest,
		Decision:         audit.DecisionAllow,
		Reasons:          uniqueStrings(append([]string{event.EventFamily, event.CurrentState}, event.Reasons...)),
		RuntimeIntegrity: payload,
		Timestamp:        event.ObservedAt.UTC(),
	})
	return err
}

func (s server) listRuntimeSubstrateValAObservedEvents(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimesubstrate.RuntimeSubstrateObservedEvent, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, err
	}
	items := []runtimesubstrate.RuntimeSubstrateObservedEvent{}
	for _, subject := range snapshot.sortedSubjects() {
		for _, observation := range subject.Observations {
			if event, ok := runtimeSubstrateObservedEventFromObservation(observation, subject); ok {
				items = append(items, event)
			}
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ObservedAt.After(items[j].ObservedAt)
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, nil
}

func runtimeSubstrateObservedEventFromObservation(observation runtimeObservation, subject *runtimeSnapshotSubject) (runtimesubstrate.RuntimeSubstrateObservedEvent, bool) {
	payload, ok := parseRuntimeSubstrateValAObservationPayload(observation.EventPayload)
	if !ok {
		return runtimesubstrate.RuntimeSubstrateObservedEvent{}, false
	}
	event := runtimesubstrate.RuntimeSubstrateObservedEvent{
		SchemaVersion:         payload.SchemaVersion,
		EventID:               firstNonEmpty(strings.TrimSpace(payload.EventID), strings.TrimSpace(observation.ObservationID)),
		EventFamily:           payload.EventFamily,
		CurrentState:          payload.CurrentState,
		Process:               payload.Process,
		ParentPID:             payload.ParentPID,
		ThreadID:              payload.ThreadID,
		ContainerRuntime:      payload.ContainerRuntime,
		NamespaceMode:         payload.NamespaceMode,
		Workload:              payload.Workload,
		Node:                  payload.Node,
		AttributionConfidence: payload.AttributionConfidence,
		FreshnessState:        payload.FreshnessState,
		UnsupportedFields:     payload.UnsupportedFields,
		EvidenceRefs:          uniqueStrings(observation.EvidenceRefs),
		Reasons:               uniqueStrings(payload.Reasons),
		ObservedAt:            observation.Timestamp,
	}
	if subject != nil {
		event.Workload.ClusterID = firstNonEmpty(strings.TrimSpace(event.Workload.ClusterID), strings.TrimSpace(subject.Cluster))
		event.Workload.Namespace = firstNonEmpty(strings.TrimSpace(event.Workload.Namespace), strings.TrimSpace(subject.Namespace))
		event.Workload.WorkloadKind = firstNonEmpty(strings.TrimSpace(event.Workload.WorkloadKind), strings.TrimSpace(subject.WorkloadKind))
		event.Workload.Workload = firstNonEmpty(strings.TrimSpace(event.Workload.Workload), strings.TrimSpace(subject.Workload))
		event.Workload.ImageDigest = firstNonEmpty(strings.TrimSpace(event.Workload.ImageDigest), strings.TrimSpace(subject.ImageDigest), strings.TrimSpace(observation.ImageDigest))
	}
	if event.Workload.PodUID == "" {
		event.Workload.PodUID = strings.TrimSpace(observation.Pod)
	}
	if event.Node.NodeID == "" {
		event.Node.NodeID = strings.TrimSpace(observation.Node)
	}
	return runtimesubstrate.NormalizeRuntimeSubstrateObservedEvent(event, nil), true
}

func runtimeSubstrateValAEventPayload(event runtimesubstrate.RuntimeSubstrateObservedEvent) map[string]any {
	return map[string]any{
		"schema_version":         event.SchemaVersion,
		"event_id":               event.EventID,
		"event_family":           event.EventFamily,
		"current_state":          event.CurrentState,
		"process":                event.Process,
		"parent_pid":             event.ParentPID,
		"thread_id":              event.ThreadID,
		"container_runtime":      event.ContainerRuntime,
		"namespace_mode":         event.NamespaceMode,
		"workload":               event.Workload,
		"node":                   event.Node,
		"attribution_confidence": event.AttributionConfidence,
		"freshness_state":        event.FreshnessState,
		"unsupported_fields":     event.UnsupportedFields,
		"reasons":                event.Reasons,
	}
}

func parseRuntimeSubstrateValAObservationPayload(value map[string]any) (runtimeSubstrateValAObservationPayload, bool) {
	if len(value) == 0 {
		return runtimeSubstrateValAObservationPayload{}, false
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return runtimeSubstrateValAObservationPayload{}, false
	}
	var payload runtimeSubstrateValAObservationPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		return runtimeSubstrateValAObservationPayload{}, false
	}
	if strings.TrimSpace(payload.SchemaVersion) != runtimesubstrate.RuntimeSubstrateValAEventRecordSchema {
		return runtimeSubstrateValAObservationPayload{}, false
	}
	return payload, true
}

func runtimeSubstrateConfidenceToRuntimeConfidence(value string) string {
	switch strings.TrimSpace(value) {
	case runtimesubstrate.RuntimeSubstrateConfidenceHighFidelity:
		return runtimeConfidenceHigh
	case runtimesubstrate.RuntimeSubstrateConfidenceUnsupportedSignal:
		return runtimeConfidenceLow
	default:
		return runtimeConfidenceMedium
	}
}

func runtimeSubstrateFamilyCoverageState(events []runtimesubstrate.RuntimeSubstrateObservedEvent, family, activeState, missingState string) string {
	for _, event := range events {
		if strings.TrimSpace(event.EventFamily) == family {
			return activeState
		}
	}
	return missingState
}

func runtimeSubstrateWorkloadNodeCoverageState(events []runtimesubstrate.RuntimeSubstrateObservedEvent) string {
	for _, event := range events {
		if event.CurrentState == runtimesubstrate.RuntimeSubstrateEventStateUnsupported {
			continue
		}
		if strings.TrimSpace(event.Workload.Namespace) != "" && strings.TrimSpace(event.Workload.Workload) != "" && strings.TrimSpace(event.Node.NodeID) != "" {
			return "workload_and_node_enrichment_active"
		}
	}
	return "workload_and_node_enrichment_missing"
}

func runtimeSubstrateEmptyObservabilityLimitation(items []runtimesubstrate.RuntimeSubstrateObservedEvent) []string {
	if len(items) == 0 {
		return []string{
			"No runtime substrate observations have been recorded for the current filter yet, so observability remains incomplete until canonical observations are ingested.",
		}
	}
	return nil
}
