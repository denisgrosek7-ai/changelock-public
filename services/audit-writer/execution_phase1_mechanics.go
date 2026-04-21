package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	benchmarkfoundation "github.com/denisgrosek/changelock/internal/benchmark"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
)

type phase1BenchmarkHarnessResponse = benchmarkfoundation.FoundationHarness
type phase1BenchmarkEvaluationRequest = benchmarkfoundation.EvaluationRequest
type phase1BenchmarkEvaluationResponse = benchmarkfoundation.EvaluationResponse

type phase1AsyncTaskListResponse struct {
	SchemaVersion string                      `json:"schema_version"`
	CurrentState  string                      `json:"current_state"`
	Tasks         []audit.ExecutionTaskRecord `json:"tasks,omitempty"`
	Limitations   []string                    `json:"limitations,omitempty"`
}

type phase1AsyncTaskMutationResponse struct {
	Status string                    `json:"status"`
	Task   audit.ExecutionTaskRecord `json:"task"`
}

type phase1AsyncTaskCreateRequest struct {
	TaskType         string   `json:"task_type"`
	SourceEventID    string   `json:"source_event_id,omitempty"`
	QueueClass       string   `json:"queue_class,omitempty"`
	BackpressureTier string   `json:"backpressure_tier,omitempty"`
	TraceID          string   `json:"trace_id,omitempty"`
	CorrelationID    string   `json:"correlation_id,omitempty"`
	DecisionID       string   `json:"decision_id,omitempty"`
	IdempotencyKey   string   `json:"idempotency_key,omitempty"`
	PayloadHash      string   `json:"payload_hash,omitempty"`
	TrustContextRef  string   `json:"trust_context_ref,omitempty"`
	CausalParent     string   `json:"causal_parent,omitempty"`
	MaxAttempts      int      `json:"max_attempts,omitempty"`
	Notes            []string `json:"notes,omitempty"`
}

type phase1AsyncTaskStatusRequest struct {
	CurrentState     string `json:"current_state"`
	FailureClass     string `json:"failure_class,omitempty"`
	FailureReason    string `json:"failure_reason,omitempty"`
	IncrementAttempt bool   `json:"increment_attempt,omitempty"`
	Note             string `json:"note,omitempty"`
}

func (s server) executionFoundationBenchmarkHarnessHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	_ = principal
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, benchmarkfoundation.FoundationCatalog())
}

func (s server) executionFoundationBenchmarkEvaluateHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request phase1BenchmarkEvaluationRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	requestID := requestIDFromHeader(r)
	response := benchmarkfoundation.EvaluateFoundationRegression(request)

	record := audit.ExecutionBenchmarkEvaluationRecord{
		SchemaVersion:  audit.ExecutionBenchmarkGateSchemaVersion,
		EvaluationID:   "bench-gate-" + audit.NewRequestID(),
		ProfileID:      strings.TrimSpace(response.ProfileID),
		CurrentState:   response.CurrentState,
		OverrideReason: strings.TrimSpace(response.OverrideReason),
		ApprovedBy:     strings.TrimSpace(principal.Subject),
		ObservedAt:     response.ObservedAt,
		Results:        convertBenchmarkResults(response.Results),
	}
	payload, err := json.Marshal(record)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:           requestID,
		Component:           "audit-writer",
		EventType:           audit.EventTypeExecutionBenchmarkGateEvaluated,
		Actor:               principal.Subject,
		TenantID:            strings.TrimSpace(r.URL.Query().Get("tenant_id")),
		Environment:         strings.TrimSpace(r.URL.Query().Get("environment")),
		Decision:            benchmarkDecision(response.CurrentState),
		Reasons:             []string{"phase 1 benchmark gate evaluated", firstNonEmpty(strings.TrimSpace(response.OverrideReason), response.CurrentState)},
		ExecutionFoundation: payload,
	})
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	_ = s.persistExecutionTrace(ctx, requestID, principal.Subject, audit.NormalizeExecutionTraceRecord(audit.ExecutionTraceRecord{
		TraceID:       "benchmark:" + record.EvaluationID,
		Component:     "audit-writer",
		Operation:     "benchmark_gate_evaluate",
		TenantID:      strings.TrimSpace(r.URL.Query().Get("tenant_id")),
		Environment:   strings.TrimSpace(r.URL.Query().Get("environment")),
		DecisionID:    record.EvaluationID,
		CorrelationID: firstNonEmpty(strings.TrimSpace(requestID), record.EvaluationID),
		Status:        map[bool]string{true: "completed", false: "failed"}[response.CurrentState == "passed" || response.CurrentState == "passed_with_override"],
		StartedAt:     record.ObservedAt,
		EndedAt:       record.ObservedAt,
		Attributes: map[string]string{
			"profile_id":      record.ProfileID,
			"current_state":   record.CurrentState,
			"override_reason": record.OverrideReason,
		},
	}, time.Now))
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) executionFoundationAsyncTasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		r, err := applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		tasks, err := s.listExecutionTasks(ctx, r)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase1AsyncTaskListResponse{
			SchemaVersion: audit.ExecutionAsyncTaskSchemaVersion,
			CurrentState:  "audit_backed_durable_task_baseline",
			Tasks:         tasks,
			Limitations: []string{
				"Phase 1 async tasks are durably persisted through canonical audit events. Broader workflow migration and dedicated queue infrastructure remain future work.",
			},
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		r, err := applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		var request phase1AsyncTaskCreateRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		requestID := requestIDFromHeader(r)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()

		existing, found, err := s.findExecutionTaskByIdempotency(ctx, r, request.TaskType, request.IdempotencyKey)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		if found {
			httpjson.Write(w, http.StatusOK, phase1AsyncTaskMutationResponse{Status: "existing", Task: existing})
			return
		}

		task := audit.NormalizeExecutionTaskRecord(audit.ExecutionTaskRecord{
			TaskType:         request.TaskType,
			CurrentState:     audit.ExecutionTaskStateQueued,
			SourceComponent:  "audit-writer",
			SourceEventID:    request.SourceEventID,
			TenantID:         strings.TrimSpace(r.URL.Query().Get("tenant_id")),
			Environment:      strings.TrimSpace(r.URL.Query().Get("environment")),
			QueueClass:       request.QueueClass,
			BackpressureTier: request.BackpressureTier,
			TraceID:          firstNonEmpty(strings.TrimSpace(request.TraceID), requestID),
			CorrelationID:    firstNonEmpty(strings.TrimSpace(request.CorrelationID), strings.TrimSpace(request.TraceID), requestID),
			DecisionID:       firstNonEmpty(strings.TrimSpace(request.DecisionID), strings.TrimSpace(request.CorrelationID), requestID),
			CausalParent:     request.CausalParent,
			IdempotencyKey:   request.IdempotencyKey,
			PayloadHash:      request.PayloadHash,
			TrustContextRef:  request.TrustContextRef,
			MaxAttempts:      request.MaxAttempts,
			Notes:            request.Notes,
		}, time.Now)
		if err := audit.ValidateExecutionTaskRecord(task); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err := s.persistExecutionTask(ctx, requestID, principal.Subject, task); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		metrics.IncExecutionAsyncTask("audit-writer", task.TaskType, task.CurrentState)
		s.dispatchExecutionTask(task, principal.Subject)
		httpjson.Write(w, http.StatusCreated, phase1AsyncTaskMutationResponse{Status: "queued", Task: task})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) executionFoundationAsyncTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/v1/foundation/execution/async/tasks/")
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "task action not found"})
		return
	}
	taskID := strings.TrimSpace(parts[0])
	action := strings.TrimSpace(parts[1])

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	requestID := requestIDFromHeader(r)
	task, found, err := s.findExecutionTaskByID(ctx, r, taskID)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if !found {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "task not found"})
		return
	}

	switch {
	case action == "status" && r.Method == http.MethodPost:
		var request phase1AsyncTaskStatusRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		task.CurrentState = strings.TrimSpace(request.CurrentState)
		task.FailureClass = strings.TrimSpace(request.FailureClass)
		task.FailureReason = strings.TrimSpace(request.FailureReason)
		task.UpdatedAt = time.Now().UTC()
		if request.IncrementAttempt {
			task.Attempts++
		}
		if strings.TrimSpace(request.Note) != "" {
			task.Notes = append(task.Notes, strings.TrimSpace(request.Note))
		}
		if err := audit.ValidateExecutionTaskRecord(task); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err := s.persistExecutionTask(ctx, requestID, principal.Subject, task); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		metrics.IncExecutionAsyncTask("audit-writer", task.TaskType, task.CurrentState)
		httpjson.Write(w, http.StatusOK, phase1AsyncTaskMutationResponse{Status: "updated", Task: task})
	case action == "replay" && r.Method == http.MethodPost:
		replayTask := audit.NormalizeExecutionTaskRecord(audit.ExecutionTaskRecord{
			TaskType:         task.TaskType,
			CurrentState:     audit.ExecutionTaskStateReplayQueued,
			SourceComponent:  task.SourceComponent,
			SourceEventID:    task.SourceEventID,
			TenantID:         task.TenantID,
			Environment:      task.Environment,
			QueueClass:       task.QueueClass,
			BackpressureTier: task.BackpressureTier,
			TraceID:          task.TraceID,
			CorrelationID:    task.CorrelationID,
			DecisionID:       task.DecisionID,
			CausalParent:     task.TaskID,
			IdempotencyKey:   task.IdempotencyKey + ":replay",
			PayloadHash:      task.PayloadHash,
			TrustContextRef:  task.TrustContextRef,
			MaxAttempts:      task.MaxAttempts,
			ReplayOfTaskID:   task.TaskID,
			Notes:            append([]string{}, task.Notes...),
		}, time.Now)
		replayTask.Notes = append(replayTask.Notes, "Replay queued from prior task lineage.")
		if err := s.persistExecutionTask(ctx, requestID, principal.Subject, replayTask); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		metrics.IncExecutionAsyncReplay("audit-writer", replayTask.TaskType, "queued")
		s.dispatchExecutionTask(replayTask, principal.Subject)
		httpjson.Write(w, http.StatusCreated, phase1AsyncTaskMutationResponse{Status: "replay_queued", Task: replayTask})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) listExecutionTasks(ctx context.Context, r *http.Request) ([]audit.ExecutionTaskRecord, error) {
	filter, err := parseFilter(r)
	if err != nil {
		return nil, err
	}
	filter.Component = "audit-writer"
	filter.EventType = audit.EventTypeExecutionTaskRecorded
	if filter.Limit <= 0 {
		filter.Limit = 500
	}
	events, err := s.store.ListEvents(ctx, filter)
	if err != nil {
		return nil, err
	}
	latest := map[string]audit.ExecutionTaskRecord{}
	for _, event := range events {
		task, err := audit.UnmarshalExecutionTaskRecord(event.Event)
		if err != nil {
			continue
		}
		if _, exists := latest[task.TaskID]; !exists {
			latest[task.TaskID] = task
		}
	}
	items := make([]audit.ExecutionTaskRecord, 0, len(latest))
	for _, task := range latest {
		items = append(items, task)
	}
	sortExecutionTasks(items)
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, nil
}

func (s server) findExecutionTaskByID(ctx context.Context, r *http.Request, taskID string) (audit.ExecutionTaskRecord, bool, error) {
	tasks, err := s.listExecutionTasks(ctx, r)
	if err != nil {
		return audit.ExecutionTaskRecord{}, false, err
	}
	for _, task := range tasks {
		if task.TaskID == strings.TrimSpace(taskID) {
			return task, true, nil
		}
	}
	return audit.ExecutionTaskRecord{}, false, nil
}

func (s server) findExecutionTaskByIdempotency(ctx context.Context, r *http.Request, taskType, idempotencyKey string) (audit.ExecutionTaskRecord, bool, error) {
	if strings.TrimSpace(idempotencyKey) == "" {
		return audit.ExecutionTaskRecord{}, false, nil
	}
	tasks, err := s.listExecutionTasks(ctx, r)
	if err != nil {
		return audit.ExecutionTaskRecord{}, false, err
	}
	for _, task := range tasks {
		if task.TaskType == strings.TrimSpace(taskType) && task.IdempotencyKey == strings.TrimSpace(idempotencyKey) {
			return task, true, nil
		}
	}
	return audit.ExecutionTaskRecord{}, false, nil
}

func (s server) persistExecutionTask(ctx context.Context, requestID, actor string, task audit.ExecutionTaskRecord) error {
	payload, err := audit.MarshalExecutionTaskRecord(task)
	if err != nil {
		return err
	}
	eventIdempotencyKey, eventPayloadHash := executionTaskTransitionEventIdentity(task)
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:           firstNonEmpty(strings.TrimSpace(requestID), audit.NewRequestID()),
		Component:           "audit-writer",
		EventType:           audit.EventTypeExecutionTaskRecorded,
		Actor:               strings.TrimSpace(actor),
		TraceID:             task.TraceID,
		CorrelationID:       task.CorrelationID,
		DecisionID:          task.DecisionID,
		CausalParent:        task.CausalParent,
		IdempotencyKey:      eventIdempotencyKey,
		PayloadHash:         eventPayloadHash,
		TenantID:            task.TenantID,
		Environment:         task.Environment,
		Decision:            audit.DecisionAllow,
		Reasons:             []string{"phase 1 async task recorded", task.CurrentState, task.TaskType},
		ExecutionFoundation: payload,
	})
	return err
}

func convertBenchmarkResults(items []benchmarkfoundation.EvaluationResult) []audit.ExecutionBenchmarkGateResult {
	results := make([]audit.ExecutionBenchmarkGateResult, 0, len(items))
	for _, item := range items {
		results = append(results, audit.ExecutionBenchmarkGateResult{
			FamilyID:      item.FamilyID,
			MetricClass:   item.MetricClass,
			MetricName:    item.MetricName,
			Status:        item.Status,
			ObservedValue: item.ObservedValue,
			BaselineValue: item.BaselineValue,
			DeltaPercent:  item.DeltaPercent,
			ThresholdPct:  item.ThresholdPct,
			Summary:       item.Summary,
		})
	}
	return results
}

func benchmarkDecision(state string) string {
	switch strings.TrimSpace(state) {
	case "passed", "passed_with_override":
		return audit.DecisionAllow
	default:
		return audit.DecisionDeny
	}
}

func sortExecutionTasks(items []audit.ExecutionTaskRecord) {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[j].UpdatedAt.After(items[i].UpdatedAt) {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
}
