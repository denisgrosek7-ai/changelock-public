package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	phase1ExecutionTracesSchema         = "1.execution_foundation_traces.v1"
	phase1ExecutionProofsSchema         = "1.execution_foundation_proofs.v1"
	phase1ExecutionRotationDrillsSchema = "1.execution_foundation_rotation_drills.v1"

	phase1ExecutionTaskSyncForwardEvent = "sync_runtime_forward_event"
)

type phase1TraceListResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	CurrentState  string                       `json:"current_state"`
	Items         []audit.ExecutionTraceRecord `json:"items,omitempty"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

type phase1RotationDrillRequest struct {
	Purpose          string   `json:"purpose,omitempty"`
	NextSignerMode   string   `json:"next_signer_mode,omitempty"`
	NextKeyID        string   `json:"next_key_id,omitempty"`
	SoftwareSecret   string   `json:"software_secret,omitempty"`
	VaultAddr        string   `json:"vault_addr,omitempty"`
	VaultToken       string   `json:"vault_token,omitempty"`
	VaultTransitPath string   `json:"vault_transit_path,omitempty"`
	VaultTransitKey  string   `json:"vault_transit_key,omitempty"`
	Notes            []string `json:"notes,omitempty"`
}

type phase1RotationDrillMutationResponse struct {
	Status string                             `json:"status"`
	Drill  audit.ExecutionRotationDrillRecord `json:"drill"`
}

type phase1RotationDrillListResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	CurrentState  string                               `json:"current_state"`
	Items         []audit.ExecutionRotationDrillRecord `json:"items,omitempty"`
	Limitations   []string                             `json:"limitations,omitempty"`
}

type phase1ProofBenchmarkSummary struct {
	ProfileID    string `json:"profile_id"`
	CurrentState string `json:"current_state"`
	ObservedAt   string `json:"observed_at,omitempty"`
}

type phase1ProofAsyncSummary struct {
	CurrentState          string         `json:"current_state"`
	TaskCountsByState     map[string]int `json:"task_counts_by_state,omitempty"`
	MigratedCriticalPaths []string       `json:"migrated_critical_paths,omitempty"`
	FailureClasses        map[string]int `json:"failure_classes,omitempty"`
}

type phase1ProofTraceSummary struct {
	CurrentState          string         `json:"current_state"`
	TraceCount            int            `json:"trace_count"`
	OperationCounts       map[string]int `json:"operation_counts,omitempty"`
	FailedOperationCounts map[string]int `json:"failed_operation_counts,omitempty"`
}

type phase1ExecutionProofsResponse struct {
	SchemaVersion        string                                     `json:"schema_version"`
	CurrentState         string                                     `json:"current_state"`
	TraceSummary         phase1ProofTraceSummary                    `json:"trace_summary"`
	AsyncSummary         phase1ProofAsyncSummary                    `json:"async_summary"`
	LatestBenchmarks     []phase1ProofBenchmarkSummary              `json:"latest_benchmarks,omitempty"`
	BenchmarkArtifacts   []audit.ExecutionBenchmarkEvaluationRecord `json:"benchmark_artifacts,omitempty"`
	TraceArtifacts       []audit.ExecutionTraceRecord               `json:"trace_artifacts,omitempty"`
	TaskArtifacts        []audit.ExecutionTaskRecord                `json:"task_artifacts,omitempty"`
	LatestRotationDrills []audit.ExecutionRotationDrillRecord       `json:"latest_rotation_drills,omitempty"`
	Limitations          []string                                   `json:"limitations,omitempty"`
}

func (s server) executionFoundationTracesHandler(w http.ResponseWriter, r *http.Request) {
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
	items, err := s.listExecutionTraces(ctx, r)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	currentState := "trace_evidence_empty"
	if len(items) > 0 {
		currentState = "end_to_end_trace_evidence_active"
	}
	httpjson.Write(w, http.StatusOK, phase1TraceListResponse{
		SchemaVersion: phase1ExecutionTracesSchema,
		CurrentState:  currentState,
		Items:         items,
		Limitations: []string{
			"Phase 1 tracing is evidence-backed and correlation-safe, but it is a bounded execution trace layer rather than a full distributed tracing vendor platform.",
		},
	})
}

func (s server) executionFoundationTrustRotationDrillHandler(w http.ResponseWriter, r *http.Request) {
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
	var request phase1RotationDrillRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	requestID := requestIDFromHeader(r)
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.runExecutionRotationDrill(ctx, request, strings.TrimSpace(r.URL.Query().Get("tenant_id")), strings.TrimSpace(r.URL.Query().Get("environment")))
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := s.persistExecutionRotationDrill(ctx, requestID, principal.Subject, record); err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, phase1RotationDrillMutationResponse{
		Status: "recorded",
		Drill:  record,
	})
}

func (s server) executionFoundationTrustRotationDrillsHandler(w http.ResponseWriter, r *http.Request) {
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
	items, err := s.listExecutionRotationDrills(ctx, r)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	currentState := "rotation_drill_surface_empty"
	if len(items) > 0 {
		currentState = "provider_backed_rotation_drill_active"
	}
	httpjson.Write(w, http.StatusOK, phase1RotationDrillListResponse{
		SchemaVersion: phase1ExecutionRotationDrillsSchema,
		CurrentState:  currentState,
		Items:         items,
		Limitations: []string{
			"Rotation drills prove cutover and historical verification continuity for configured signer providers, but they do not claim a universal external KMS/HSM matrix beyond the providers currently implemented in code.",
		},
	})
}

func (s server) executionFoundationProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	traces, err := s.listExecutionTraces(ctx, r)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	tasks, err := s.listExecutionTasks(ctx, r)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	benchmarks, err := s.listExecutionBenchmarkEvaluations(ctx, r)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	drills, err := s.listExecutionRotationDrills(ctx, r)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, buildExecutionProofsResponse(traces, tasks, benchmarks, drills))
}

func buildExecutionProofsResponse(traces []audit.ExecutionTraceRecord, tasks []audit.ExecutionTaskRecord, benchmarks []audit.ExecutionBenchmarkEvaluationRecord, drills []audit.ExecutionRotationDrillRecord) phase1ExecutionProofsResponse {
	traceOps := map[string]int{}
	traceFailed := map[string]int{}
	for _, trace := range traces {
		traceOps[trace.Operation]++
		if trace.Status != "completed" {
			traceFailed[trace.Operation]++
		}
	}
	taskCounts := map[string]int{}
	failureClasses := map[string]int{}
	for _, task := range tasks {
		taskCounts[task.CurrentState]++
		if task.FailureClass != "" {
			failureClasses[task.FailureClass]++
		}
	}
	latestBenchmarks := make([]phase1ProofBenchmarkSummary, 0, minInt(len(benchmarks), 5))
	hasPassingBenchmark := false
	for _, item := range benchmarks {
		if item.CurrentState == "passed" || item.CurrentState == "passed_with_override" {
			hasPassingBenchmark = true
		}
		if len(latestBenchmarks) == 5 {
			break
		}
		latestBenchmarks = append(latestBenchmarks, phase1ProofBenchmarkSummary{
			ProfileID:    item.ProfileID,
			CurrentState: item.CurrentState,
			ObservedAt:   item.ObservedAt.Format(time.RFC3339),
		})
	}
	if len(drills) > 5 {
		drills = drills[:5]
	}
	traceArtifacts := traces
	if len(traceArtifacts) > 5 {
		traceArtifacts = traceArtifacts[:5]
	}
	taskArtifacts := tasks
	if len(taskArtifacts) > 5 {
		taskArtifacts = taskArtifacts[:5]
	}
	benchmarkArtifacts := benchmarks
	if len(benchmarkArtifacts) > 5 {
		benchmarkArtifacts = benchmarkArtifacts[:5]
	}
	currentState := "phase1_closure_incomplete"
	if len(traces) > 0 && taskCounts[audit.ExecutionTaskStateCompleted] > 0 && len(drills) > 0 && hasPassingBenchmark {
		currentState = "phase1_operational_proof_active"
	}
	return phase1ExecutionProofsResponse{
		SchemaVersion: phase1ExecutionProofsSchema,
		CurrentState:  currentState,
		TraceSummary: phase1ProofTraceSummary{
			CurrentState:          "trace_records_present",
			TraceCount:            len(traces),
			OperationCounts:       traceOps,
			FailedOperationCounts: traceFailed,
		},
		AsyncSummary: phase1ProofAsyncSummary{
			CurrentState:          "critical_path_migration_evidence_present",
			TaskCountsByState:     taskCounts,
			MigratedCriticalPaths: []string{"sync_runtime_forward_event"},
			FailureClasses:        failureClasses,
		},
		LatestBenchmarks:     latestBenchmarks,
		BenchmarkArtifacts:   benchmarkArtifacts,
		TraceArtifacts:       traceArtifacts,
		TaskArtifacts:        taskArtifacts,
		LatestRotationDrills: drills,
		Limitations: []string{
			"Operational proofs expose bounded evidence artifacts for tracing, async migration, benchmark regression, and signer rotation drill outputs.",
		},
	}
}

func (s server) dispatchExecutionTask(task audit.ExecutionTaskRecord, actor string) {
	if !phase1TaskAutoExecutable(task.TaskType) {
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout)
		defer cancel()
		if err := s.executeExecutionTask(ctx, audit.NewRequestID(), actor, task); err != nil {
			// The failure state is persisted inside executeExecutionTask; do not surface a second mutation path here.
		}
	}()
}

func (s server) enqueueSyncForwardTask(ctx context.Context, requestID, actor string, event audit.Event) (audit.ExecutionTaskRecord, error) {
	if s.syncRuntime == nil || s.syncRuntime.config.Mode != audit.SyncModeSpoke {
		return audit.ExecutionTaskRecord{}, nil
	}
	task := audit.NormalizeExecutionTaskRecord(audit.ExecutionTaskRecord{
		TaskType:         phase1ExecutionTaskSyncForwardEvent,
		CurrentState:     audit.ExecutionTaskStateQueued,
		SourceComponent:  "audit-writer",
		SourceEventID:    event.EventID,
		TenantID:         event.TenantID,
		Environment:      event.Environment,
		QueueClass:       "connector",
		BackpressureTier: "bounded",
		TraceID:          event.TraceID,
		CorrelationID:    event.CorrelationID,
		DecisionID:       event.DecisionID,
		CausalParent:     event.EventID,
		IdempotencyKey:   event.IdempotencyKey + ":sync-forward",
		PayloadHash:      event.PayloadHash,
		TrustContextRef:  "sync_runtime:" + s.syncRuntime.config.Mode,
		MaxAttempts:      3,
		Notes: []string{
			"Critical path migrated from synchronous sync forward to bounded async task.",
			"Source event remains canonical evidence; task only carries orchestration lineage.",
		},
	}, time.Now)
	if err := s.persistExecutionTask(ctx, requestID, actor, task); err != nil {
		return audit.ExecutionTaskRecord{}, err
	}
	metrics.IncExecutionAsyncTask("audit-writer", task.TaskType, task.CurrentState)
	s.dispatchExecutionTask(task, actor)
	return task, nil
}

func phase1TaskAutoExecutable(taskType string) bool {
	switch strings.TrimSpace(taskType) {
	case phase1ExecutionTaskSyncForwardEvent:
		return true
	default:
		return false
	}
}

func (s server) executeExecutionTask(ctx context.Context, requestID, actor string, task audit.ExecutionTaskRecord) error {
	switch strings.TrimSpace(task.TaskType) {
	case phase1ExecutionTaskSyncForwardEvent:
		return s.executeSyncForwardTask(ctx, requestID, actor, task)
	default:
		return nil
	}
}

func (s server) executeSyncForwardTask(ctx context.Context, requestID, actor string, task audit.ExecutionTaskRecord) error {
	startedAt := time.Now().UTC()
	running := task
	running.CurrentState = audit.ExecutionTaskStateRunning
	running.Attempts++
	running.UpdatedAt = startedAt
	running.FailureClass = ""
	running.FailureReason = ""
	if err := s.persistExecutionTask(ctx, requestID, actor, running); err != nil {
		return err
	}
	metrics.IncExecutionAsyncTask("audit-writer", running.TaskType, running.CurrentState)

	sourceEvent, err := s.findStoredEventByEventID(ctx, running)
	if err != nil {
		return s.finishExecutionTask(ctx, requestID, actor, running, "failed", classifyExecutionTaskFailure(err), err.Error(), audit.ExecutionTaskStateFailedTerminal, startedAt, nil)
	}
	if s.syncRuntime == nil || s.syncRuntime.config.Mode != audit.SyncModeSpoke {
		return s.finishExecutionTask(ctx, requestID, actor, running, "failed", audit.ExecutionTaskFailureBusinessRule, "sync runtime is not in spoke mode", audit.ExecutionTaskStateFailedTerminal, startedAt, nil)
	}
	err = s.syncRuntime.forwardEvent(ctx, sourceEvent.Event)
	if err != nil {
		state := audit.ExecutionTaskStateFailedRetryable
		if running.Attempts >= running.MaxAttempts {
			state = audit.ExecutionTaskStateDeadLettered
		}
		return s.finishExecutionTask(ctx, requestID, actor, running, "failed", classifyExecutionTaskFailure(err), err.Error(), state, startedAt, map[string]string{
			"source_event_id": sourceEvent.EventID,
		})
	}
	return s.finishExecutionTask(ctx, requestID, actor, running, "completed", "", "", audit.ExecutionTaskStateCompleted, startedAt, map[string]string{
		"source_event_id":   sourceEvent.EventID,
		"forwarded_cluster": sourceEvent.ClusterID,
	})
}

func (s server) finishExecutionTask(ctx context.Context, requestID, actor string, task audit.ExecutionTaskRecord, traceStatus, failureClass, failureReason, nextState string, startedAt time.Time, attributes map[string]string) error {
	updated := task
	updated.CurrentState = nextState
	updated.UpdatedAt = time.Now().UTC()
	updated.FailureClass = strings.TrimSpace(failureClass)
	updated.FailureReason = strings.TrimSpace(failureReason)
	if nextState == audit.ExecutionTaskStateCompleted {
		updated.Notes = append(updated.Notes, "Execution completed through bounded async worker.")
	}
	if err := s.persistExecutionTask(ctx, requestID, actor, updated); err != nil {
		return err
	}
	metrics.IncExecutionAsyncTask("audit-writer", updated.TaskType, updated.CurrentState)
	trace := audit.NormalizeExecutionTraceRecord(audit.ExecutionTraceRecord{
		TraceID:       updated.TraceID,
		ParentSpanID:  updated.CausalParent,
		Component:     "audit-writer",
		Operation:     updated.TaskType,
		TenantID:      updated.TenantID,
		Environment:   updated.Environment,
		EventID:       updated.SourceEventID,
		DecisionID:    updated.DecisionID,
		CorrelationID: updated.CorrelationID,
		Status:        traceStatus,
		StartedAt:     startedAt,
		EndedAt:       updated.UpdatedAt,
		Attributes:    attributes,
		Notes:         append([]string{}, updated.Notes...),
	}, time.Now)
	if traceAttributes := trace.Attributes; traceAttributes == nil {
		trace.Attributes = map[string]string{}
	}
	trace.Attributes["task_state"] = updated.CurrentState
	if updated.FailureClass != "" {
		trace.Attributes["failure_class"] = updated.FailureClass
	}
	return s.persistExecutionTrace(ctx, requestID, actor, trace)
}

func classifyExecutionTaskFailure(err error) string {
	switch {
	case err == nil:
		return ""
	case errors.Is(err, context.DeadlineExceeded):
		return audit.ExecutionTaskFailureTransientExternal
	default:
		return audit.ExecutionTaskFailureRetryable
	}
}

func (s server) findStoredEventByEventID(ctx context.Context, task audit.ExecutionTaskRecord) (audit.StoredEvent, error) {
	filter := audit.EventFilter{
		TenantID:    task.TenantID,
		Environment: task.Environment,
		Limit:       5000,
	}
	events, err := s.store.ListEvents(ctx, filter)
	if err != nil {
		return audit.StoredEvent{}, err
	}
	for _, item := range events {
		if item.EventID == task.SourceEventID {
			return item, nil
		}
	}
	return audit.StoredEvent{}, errors.New("source event for async task not found")
}

func (s server) persistExecutionTrace(ctx context.Context, requestID, actor string, record audit.ExecutionTraceRecord) error {
	payload, err := audit.MarshalExecutionTraceRecord(record)
	if err != nil {
		return err
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:           firstNonEmpty(strings.TrimSpace(requestID), audit.NewRequestID()),
		Component:           record.Component,
		EventType:           audit.EventTypeExecutionTraceRecorded,
		Actor:               strings.TrimSpace(actor),
		TraceID:             record.TraceID,
		CorrelationID:       record.CorrelationID,
		DecisionID:          record.DecisionID,
		CausalParent:        record.ParentSpanID,
		IdempotencyKey:      record.SpanID,
		PayloadHash:         audit.CanonicalDecisionID(audit.Event{RequestID: record.TraceID, Component: record.Component, EventType: record.Operation, Decision: audit.DecisionAllow}),
		TenantID:            record.TenantID,
		Environment:         record.Environment,
		Decision:            benchmarkDecision(traceDecisionState(record.Status)),
		Reasons:             []string{"phase 1 execution trace recorded", record.Operation, record.Status},
		ExecutionFoundation: payload,
	})
	return err
}

func executionTaskTransitionEventIdentity(task audit.ExecutionTaskRecord) (string, string) {
	attempts := maxInt(task.Attempts, 0)
	eventID := strings.Join([]string{
		task.TaskID,
		task.CurrentState,
		strconv.Itoa(attempts),
		firstNonEmpty(task.FailureClass, "none"),
		firstNonEmpty(task.ReplayOfTaskID, "root"),
	}, ":")
	payloadHash := strings.Join([]string{
		task.TaskID,
		task.TaskType,
		task.CurrentState,
		strconv.Itoa(attempts),
		firstNonEmpty(task.SourceEventID, "none"),
		firstNonEmpty(task.FailureClass, "none"),
		firstNonEmpty(task.FailureReason, "none"),
		firstNonEmpty(task.ReplayOfTaskID, "root"),
	}, ":")
	return eventID, payloadHash
}

func traceDecisionState(status string) string {
	if strings.TrimSpace(status) == "completed" {
		return "passed"
	}
	return "failed"
}

func (s server) listExecutionTraces(ctx context.Context, r *http.Request) ([]audit.ExecutionTraceRecord, error) {
	filter, err := parseFilter(r)
	if err != nil {
		return nil, err
	}
	filter.Component = "audit-writer"
	filter.EventType = audit.EventTypeExecutionTraceRecorded
	if filter.Limit <= 0 {
		filter.Limit = 500
	}
	traceID := strings.TrimSpace(r.URL.Query().Get("trace_id"))
	events, err := s.store.ListEvents(ctx, filter)
	if err != nil {
		return nil, err
	}
	items := make([]audit.ExecutionTraceRecord, 0, len(events))
	for _, event := range events {
		record, err := audit.UnmarshalExecutionTraceRecord(event.Event)
		if err != nil {
			continue
		}
		if traceID != "" && record.TraceID != traceID {
			continue
		}
		items = append(items, record)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].StartedAt.Equal(items[j].StartedAt) {
			return items[i].SpanID < items[j].SpanID
		}
		return items[i].StartedAt.After(items[j].StartedAt)
	})
	return items, nil
}

func (s server) runExecutionRotationDrill(ctx context.Context, request phase1RotationDrillRequest, tenantID, environment string) (audit.ExecutionRotationDrillRecord, error) {
	if s.signing == nil || s.signing.runtime == nil || !s.signing.runtime.Enabled() {
		return audit.ExecutionRotationDrillRecord{}, errors.New("current signing runtime is not enabled")
	}
	purpose := firstNonEmpty(strings.TrimSpace(request.Purpose), signing.PurposeSyncSnapshots)
	if !s.signing.runtime.SupportsPurpose(purpose) {
		return audit.ExecutionRotationDrillRecord{}, errors.New("current signing runtime does not support the requested purpose")
	}
	nextRuntime, err := buildRotationDrillRuntime(request, purpose)
	if err != nil {
		return audit.ExecutionRotationDrillRecord{}, err
	}
	payload := []byte("phase1-rotation-drill:" + purpose + ":" + time.Now().UTC().Format(time.RFC3339Nano))
	currentEnvelope, err := s.signing.runtime.Sign(ctx, purpose, payload)
	if err != nil {
		return audit.ExecutionRotationDrillRecord{}, err
	}
	nextEnvelope, err := nextRuntime.Sign(ctx, purpose, payload)
	if err != nil {
		return audit.ExecutionRotationDrillRecord{}, err
	}
	trustSet := signing.TrustSet{
		Members: []signing.TrustSetMember{
			{MemberID: "current-retired", Runtime: s.signing.runtime, LifecycleState: signing.KeyStateRetiredVerifyOnly},
			{MemberID: "next-active", Runtime: nextRuntime, LifecycleState: signing.KeyStateActive},
		},
	}
	currentResult, currentPath, err := trustSet.Verify(ctx, purpose, payload, currentEnvelope)
	if err != nil {
		return audit.ExecutionRotationDrillRecord{}, err
	}
	nextResult, nextPath, err := trustSet.Verify(ctx, purpose, payload, nextEnvelope)
	if err != nil {
		return audit.ExecutionRotationDrillRecord{}, err
	}
	revokedSet := signing.TrustSet{
		Members: []signing.TrustSetMember{
			{MemberID: "current-revoked", Runtime: s.signing.runtime, LifecycleState: signing.KeyStateRevoked},
			{MemberID: "next-active", Runtime: nextRuntime, LifecycleState: signing.KeyStateActive},
		},
	}
	revokedResult, revokedPath, err := revokedSet.Verify(ctx, purpose, payload, currentEnvelope)
	if err != nil {
		return audit.ExecutionRotationDrillRecord{}, err
	}
	currentState := "passed"
	if currentResult.State != signing.StateVerified || nextResult.State != signing.StateVerified || revokedResult.State == signing.StateVerified {
		currentState = "failed"
	}
	record := audit.NormalizeExecutionRotationDrillRecord(audit.ExecutionRotationDrillRecord{
		Purpose:     purpose,
		TenantID:    tenantID,
		Environment: environment,
		CurrentRuntime: audit.ExecutionRotationDrillRuntime{
			ProviderMode: s.signing.runtime.Config.Mode,
			KeyID:        s.signing.runtime.Config.KeyID,
		},
		NextRuntime: audit.ExecutionRotationDrillRuntime{
			ProviderMode: nextRuntime.Config.Mode,
			KeyID:        nextRuntime.Config.KeyID,
		},
		CurrentSignedAt: currentEnvelope.SignedAt,
		NextSignedAt:    nextEnvelope.SignedAt,
		CurrentVerification: audit.ExecutionRotationDrillVerification{
			State:          currentResult.State,
			Reason:         currentResult.Reason,
			VerificationBy: currentPath.MemberID,
			LifecycleState: currentPath.LifecycleState,
		},
		NextVerification: audit.ExecutionRotationDrillVerification{
			State:          nextResult.State,
			Reason:         nextResult.Reason,
			VerificationBy: nextPath.MemberID,
			LifecycleState: nextPath.LifecycleState,
		},
		RevokedVerification: audit.ExecutionRotationDrillVerification{
			State:          revokedResult.State,
			Reason:         revokedResult.Reason,
			VerificationBy: revokedPath.MemberID,
			LifecycleState: revokedPath.LifecycleState,
		},
		CurrentState:         currentState,
		TrustBoundary:        s.signing.runtime.DescribeProvider().TrustBoundary,
		HistoricalVerifyPath: currentPath.MemberID,
		Notes: append([]string{
			"Current signer is retired to verify-only for historical continuity.",
			"Next signer becomes active for new signatures.",
			"Revoked signer must fail historical verification after revocation takes effect.",
		}, request.Notes...),
	}, time.Now)
	return record, audit.ValidateExecutionRotationDrillRecord(record)
}

func buildRotationDrillRuntime(request phase1RotationDrillRequest, purpose string) (*signing.Runtime, error) {
	mode := strings.TrimSpace(request.NextSignerMode)
	if mode == "" {
		return nil, errors.New("next_signer_mode is required")
	}
	config := signing.Config{
		Mode:             mode,
		Purposes:         map[string]struct{}{purpose: {}},
		KeyID:            strings.TrimSpace(request.NextKeyID),
		VerifyOnRead:     true,
		SoftwareSecret:   strings.TrimSpace(request.SoftwareSecret),
		VaultAddr:        strings.TrimSpace(request.VaultAddr),
		VaultToken:       strings.TrimSpace(request.VaultToken),
		VaultTransitPath: strings.TrimSpace(request.VaultTransitPath),
		VaultTransitKey:  strings.TrimSpace(request.VaultTransitKey),
	}
	return signing.NewRuntime(config, signing.ProviderOptions{})
}

func (s server) persistExecutionRotationDrill(ctx context.Context, requestID, actor string, record audit.ExecutionRotationDrillRecord) error {
	payload, err := audit.MarshalExecutionRotationDrillRecord(record)
	if err != nil {
		return err
	}
	trace := audit.NormalizeExecutionTraceRecord(audit.ExecutionTraceRecord{
		TraceID:       "rotation:" + record.DrillID,
		Component:     "audit-writer",
		Operation:     "trust_rotation_drill",
		TenantID:      record.TenantID,
		Environment:   record.Environment,
		DecisionID:    record.DrillID,
		CorrelationID: record.DrillID,
		Status:        map[bool]string{true: "completed", false: "failed"}[record.CurrentState == "passed"],
		StartedAt:     record.ObservedAt,
		EndedAt:       record.ObservedAt,
		Attributes: map[string]string{
			"purpose":       record.Purpose,
			"current_mode":  record.CurrentRuntime.ProviderMode,
			"next_mode":     record.NextRuntime.ProviderMode,
			"current_state": record.CurrentState,
		},
	}, time.Now)
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:           firstNonEmpty(strings.TrimSpace(requestID), audit.NewRequestID()),
		Component:           "audit-writer",
		EventType:           audit.EventTypeExecutionTrustRotationDrill,
		Actor:               strings.TrimSpace(actor),
		TraceID:             trace.TraceID,
		CorrelationID:       trace.CorrelationID,
		DecisionID:          trace.DecisionID,
		IdempotencyKey:      record.DrillID,
		PayloadHash:         audit.CanonicalDecisionID(audit.Event{RequestID: record.DrillID, Component: "audit-writer", EventType: audit.EventTypeExecutionTrustRotationDrill, Decision: benchmarkDecision(record.CurrentState)}),
		TenantID:            record.TenantID,
		Environment:         record.Environment,
		Decision:            benchmarkDecision(record.CurrentState),
		Reasons:             []string{"phase 1 signer rotation drill", record.CurrentState, record.Purpose},
		ExecutionFoundation: payload,
	})
	if err != nil {
		return err
	}
	return s.persistExecutionTrace(ctx, requestID, actor, trace)
}

func (s server) listExecutionRotationDrills(ctx context.Context, r *http.Request) ([]audit.ExecutionRotationDrillRecord, error) {
	filter, err := parseFilter(r)
	if err != nil {
		return nil, err
	}
	filter.Component = "audit-writer"
	filter.EventType = audit.EventTypeExecutionTrustRotationDrill
	if filter.Limit <= 0 {
		filter.Limit = 100
	}
	events, err := s.store.ListEvents(ctx, filter)
	if err != nil {
		return nil, err
	}
	items := make([]audit.ExecutionRotationDrillRecord, 0, len(events))
	for _, event := range events {
		record, err := audit.UnmarshalExecutionRotationDrillRecord(event.Event)
		if err != nil {
			continue
		}
		items = append(items, record)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ObservedAt.After(items[j].ObservedAt)
	})
	return items, nil
}

func (s server) listExecutionBenchmarkEvaluations(ctx context.Context, r *http.Request) ([]audit.ExecutionBenchmarkEvaluationRecord, error) {
	filter, err := parseFilter(r)
	if err != nil {
		return nil, err
	}
	filter.Component = "audit-writer"
	filter.EventType = audit.EventTypeExecutionBenchmarkGateEvaluated
	if filter.Limit <= 0 {
		filter.Limit = 100
	}
	events, err := s.store.ListEvents(ctx, filter)
	if err != nil {
		return nil, err
	}
	items := make([]audit.ExecutionBenchmarkEvaluationRecord, 0, len(events))
	for _, event := range events {
		var record audit.ExecutionBenchmarkEvaluationRecord
		if len(event.ExecutionFoundation) == 0 || string(event.ExecutionFoundation) == "null" {
			continue
		}
		if err := json.Unmarshal(event.ExecutionFoundation, &record); err != nil {
			continue
		}
		items = append(items, record)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ObservedAt.After(items[j].ObservedAt)
	})
	return items, nil
}
