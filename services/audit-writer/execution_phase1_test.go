package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	benchmarkfoundation "github.com/denisgrosek/changelock/internal/benchmark"
	"github.com/denisgrosek/changelock/internal/signing"
)

func TestExecutionFoundationSummaryHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected execution foundation summary 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase1ExecutionFoundationResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode execution foundation summary: %v", err)
	}
	if response.SchemaVersion != phase1ExecutionFoundationSchema || len(response.Gates) < 4 {
		t.Fatalf("expected phase 1 foundation summary with gates, got %#v", response)
	}
}

func TestExecutionFoundationContractsHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/contracts", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected execution foundation contracts 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase1ExecutionContractsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode execution foundation contracts: %v", err)
	}
	if response.SchemaVersion != phase1ExecutionContractsSchema || response.CanonicalEventSchema != audit.ExecutionEventSchemaVersion {
		t.Fatalf("expected execution contract response with canonical event schema, got %#v", response)
	}
	if len(response.EnvelopeFields) < 8 || len(response.DegradedModes) < 4 {
		t.Fatalf("expected canonical envelope fields and degraded modes, got %#v", response)
	}
}

func TestExecutionFoundationAsyncHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/async", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected execution foundation async 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase1ExecutionAsyncResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode execution foundation async: %v", err)
	}
	if response.SchemaVersion != phase1ExecutionAsyncSchema || response.EventEnvelope.SchemaVersion != audit.ExecutionEventSchemaVersion {
		t.Fatalf("expected async foundation response with canonical event envelope, got %#v", response)
	}
	if len(response.SynchronousPath) == 0 || len(response.MigratedAsyncPaths) == 0 || len(response.TargetAsyncPaths) == 0 || len(response.FailureSemantics) < 5 {
		t.Fatalf("expected sync/async split and failure semantics, got %#v", response)
	}
}

func TestExecutionFoundationBenchmarksHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/benchmarks", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected execution foundation benchmarks 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase1ExecutionBenchmarksResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode execution foundation benchmarks: %v", err)
	}
	if response.SchemaVersion != phase1ExecutionBenchmarksSchema || len(response.Profiles) < 3 || len(response.CriticalPaths) < 4 {
		t.Fatalf("expected benchmark foundation profiles and measured paths, got %#v", response)
	}
}

func TestExecutionFoundationBenchmarkHarnessHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/benchmarks/harness", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected benchmark harness 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response benchmarkfoundation.FoundationHarness
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode benchmark harness: %v", err)
	}
	if response.SchemaVersion != benchmarkfoundation.FoundationHarnessSchemaVersion || len(response.Families) < 4 {
		t.Fatalf("expected benchmark harness families, got %#v", response)
	}
}

func TestExecutionFoundationBenchmarkEvaluateHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/benchmarks/evaluate?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"profile_id":"production_like",
		"observations":[
			{
				"family_id":"deploy_gate_admission",
				"profile_id":"production_like",
				"metric_class":"user_facing_latency",
				"metric_name":"p95_latency_ms",
				"unit":"ms",
				"baseline_value":100,
				"observed_value":125
			}
		],
		"override":{"reason":"known CI runner noise","approved_by":"secops"}
	}`))
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected benchmark evaluate 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response benchmarkfoundation.EvaluationResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode benchmark evaluation: %v", err)
	}
	if response.CurrentState != "passed_with_override" || len(response.Results) != 1 {
		t.Fatalf("expected benchmark override response, got %#v", response)
	}

	events, err := fixture.store.ListEvents(req.Context(), audit.EventFilter{EventType: audit.EventTypeExecutionBenchmarkGateEvaluated, Component: "audit-writer", Limit: 20})
	if err != nil {
		t.Fatalf("ListEvents() error = %v", err)
	}
	if len(events) == 0 || len(events[0].ExecutionFoundation) == 0 {
		t.Fatalf("expected persisted benchmark gate evaluation event, got %#v", events)
	}
}

func TestExecutionFoundationTrustHandler(t *testing.T) {
	t.Run("disabled signer", func(t *testing.T) {
		handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

		req := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/trust", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected execution foundation trust 200, got %d: %s", rec.Code, rec.Body.String())
		}

		var response phase1ExecutionTrustResponse
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("decode execution foundation trust: %v", err)
		}
		if response.SchemaVersion != phase1ExecutionTrustSchema || response.Provider.ProviderMode != signing.ModeDisabled {
			t.Fatalf("expected disabled provider descriptor, got %#v", response)
		}
	})

	t.Run("software signer", func(t *testing.T) {
		t.Setenv("CHANGELOCK_SIGNER_MODE", signing.ModeSoftware)
		t.Setenv("CHANGELOCK_SIGNER_SOFTWARE_SECRET", "phase1-secret")
		t.Setenv("CHANGELOCK_SIGNER_KEY_ID", "phase1-software")
		signingRuntime, err := loadSigningRuntimeFromEnv()
		if err != nil {
			t.Fatalf("loadSigningRuntimeFromEnv() error = %v", err)
		}
		handler := newHandlerWithDeps(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t), nil, newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}), signingRuntime)

		req := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/trust", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected execution foundation trust 200, got %d: %s", rec.Code, rec.Body.String())
		}

		var response phase1ExecutionTrustResponse
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("decode execution foundation trust: %v", err)
		}
		if response.Provider.ProviderMode != signing.ModeSoftware || response.CurrentState != "provider_abstraction_ready_local_signer_active_rotation_drill_ready" {
			t.Fatalf("expected software provider descriptor, got %#v", response)
		}
		if len(response.ProviderCapabilityMatrix) < 3 || len(response.LifecycleStates) < 6 {
			t.Fatalf("expected capability matrix and lifecycle states, got %#v", response)
		}
	})
}

func TestExecutionFoundationAsyncTaskLifecycleHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	createReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/async/tasks?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"task_type":"connector_dispatch",
		"queue_class":"connector",
		"backpressure_tier":"bounded",
		"idempotency_key":"dispatch-acme-1",
		"payload_hash":"sha256:task-payload",
		"max_attempts":4
	}`))
	createReq.Header.Set("Authorization", "Bearer operator-demo-token")
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("expected async task create 201, got %d: %s", createRec.Code, createRec.Body.String())
	}

	var created phase1AsyncTaskMutationResponse
	if err := json.NewDecoder(createRec.Body).Decode(&created); err != nil {
		t.Fatalf("decode created task: %v", err)
	}
	if created.Task.TaskID == "" || created.Task.CurrentState != audit.ExecutionTaskStateQueued {
		t.Fatalf("expected queued task, got %#v", created)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/async/tasks?tenant_id=acme&environment=prod&limit=20", nil)
	listReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	listRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected async task list 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	var listed phase1AsyncTaskListResponse
	if err := json.NewDecoder(listRec.Body).Decode(&listed); err != nil {
		t.Fatalf("decode task list: %v", err)
	}
	if len(listed.Tasks) == 0 || listed.Tasks[0].TaskID != created.Task.TaskID {
		t.Fatalf("expected created task in list, got %#v", listed)
	}

	statusReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/async/tasks/"+created.Task.TaskID+"/status?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"current_state":"failed_retryable",
		"failure_class":"transient_external",
		"failure_reason":"connector timeout",
		"increment_attempt":true
	}`))
	statusReq.Header.Set("Authorization", "Bearer operator-demo-token")
	statusReq.Header.Set("Content-Type", "application/json")
	statusRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(statusRec, statusReq)
	if statusRec.Code != http.StatusOK {
		t.Fatalf("expected async task status update 200, got %d: %s", statusRec.Code, statusRec.Body.String())
	}

	taskEvents, err := fixture.store.ListEvents(context.Background(), audit.EventFilter{
		Component: "audit-writer",
		EventType: audit.EventTypeExecutionTaskRecorded,
		TenantID:  "acme",
		Limit:     20,
	})
	if err != nil {
		t.Fatalf("ListEvents(task events) error = %v", err)
	}
	transitionIdempotencyKeys := map[string]struct{}{}
	logicalTaskKeys := map[string]struct{}{}
	for _, item := range taskEvents {
		task, err := audit.UnmarshalExecutionTaskRecord(item.Event)
		if err != nil || task.TaskID != created.Task.TaskID {
			continue
		}
		transitionIdempotencyKeys[item.Event.IdempotencyKey] = struct{}{}
		logicalTaskKeys[task.IdempotencyKey] = struct{}{}
	}
	if len(logicalTaskKeys) != 1 {
		t.Fatalf("expected one logical task idempotency key across lifecycle transitions, got %#v", logicalTaskKeys)
	}
	if len(transitionIdempotencyKeys) < 2 {
		t.Fatalf("expected distinct event-level idempotency keys across task state transitions, got %#v", transitionIdempotencyKeys)
	}

	replayReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/async/tasks/"+created.Task.TaskID+"/replay?tenant_id=acme&environment=prod", nil)
	replayReq.Header.Set("Authorization", "Bearer operator-demo-token")
	replayRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(replayRec, replayReq)
	if replayRec.Code != http.StatusCreated {
		t.Fatalf("expected async task replay 201, got %d: %s", replayRec.Code, replayRec.Body.String())
	}

	var replayed phase1AsyncTaskMutationResponse
	if err := json.NewDecoder(replayRec.Body).Decode(&replayed); err != nil {
		t.Fatalf("decode replayed task: %v", err)
	}
	if replayed.Task.ReplayOfTaskID != created.Task.TaskID || replayed.Task.CurrentState != audit.ExecutionTaskStateReplayQueued {
		t.Fatalf("expected replay lineage, got %#v", replayed)
	}
}

func TestExecutionFoundationAsyncTaskIdempotencyLookupEscapesLastFiftyTaskWindow(t *testing.T) {
	fixture := forensicsTestFixture(t)

	createReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/async/tasks?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"task_type":"connector_dispatch",
		"queue_class":"connector",
		"backpressure_tier":"bounded",
		"idempotency_key":"dispatch-acme-stable",
		"payload_hash":"sha256:task-payload",
		"max_attempts":4
	}`))
	createReq.Header.Set("Authorization", "Bearer operator-demo-token")
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("expected async task create 201, got %d: %s", createRec.Code, createRec.Body.String())
	}

	var created phase1AsyncTaskMutationResponse
	if err := json.NewDecoder(createRec.Body).Decode(&created); err != nil {
		t.Fatalf("decode created task: %v", err)
	}

	for i := 0; i < 75; i++ {
		task := audit.NormalizeExecutionTaskRecord(audit.ExecutionTaskRecord{
			TaskType:         "connector_dispatch",
			CurrentState:     audit.ExecutionTaskStateQueued,
			SourceComponent:  "audit-writer",
			TenantID:         "acme",
			Environment:      "prod",
			QueueClass:       "connector",
			BackpressureTier: "bounded",
			TraceID:          fmt.Sprintf("trace-seed-%03d", i),
			CorrelationID:    fmt.Sprintf("correlation-seed-%03d", i),
			DecisionID:       fmt.Sprintf("decision-seed-%03d", i),
			IdempotencyKey:   fmt.Sprintf("dispatch-acme-seed-%03d", i),
			PayloadHash:      fmt.Sprintf("sha256:seed-%03d", i),
			MaxAttempts:      4,
		}, time.Now)
		mustIngestExecutionTaskRecord(t, fixture.store, fmt.Sprintf("seed-%03d", i), task)
	}

	repeatReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/async/tasks?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"task_type":"connector_dispatch",
		"queue_class":"connector",
		"backpressure_tier":"bounded",
		"idempotency_key":"dispatch-acme-stable",
		"payload_hash":"sha256:task-payload",
		"max_attempts":4
	}`))
	repeatReq.Header.Set("Authorization", "Bearer operator-demo-token")
	repeatReq.Header.Set("Content-Type", "application/json")
	repeatRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(repeatRec, repeatReq)
	if repeatRec.Code != http.StatusOK {
		t.Fatalf("expected existing task 200, got %d: %s", repeatRec.Code, repeatRec.Body.String())
	}

	var repeated phase1AsyncTaskMutationResponse
	if err := json.NewDecoder(repeatRec.Body).Decode(&repeated); err != nil {
		t.Fatalf("decode repeated task: %v", err)
	}
	if repeated.Status != "existing" || repeated.Task.TaskID != created.Task.TaskID {
		t.Fatalf("expected canonical existing task, got %#v", repeated)
	}

	events, err := fixture.store.ListEvents(context.Background(), audit.EventFilter{
		Component:   "audit-writer",
		EventType:   audit.EventTypeExecutionTaskRecorded,
		TenantID:    "acme",
		Environment: "prod",
		Limit:       5000,
	})
	if err != nil {
		t.Fatalf("ListEvents(task events) error = %v", err)
	}
	logicalTaskIDs := map[string]struct{}{}
	for _, item := range events {
		task, err := audit.UnmarshalExecutionTaskRecord(item.Event)
		if err != nil {
			continue
		}
		if task.TaskType == "connector_dispatch" && task.IdempotencyKey == "dispatch-acme-stable" {
			logicalTaskIDs[task.TaskID] = struct{}{}
		}
	}
	if len(logicalTaskIDs) != 1 {
		t.Fatalf("expected one canonical task for idempotency key, got %#v", logicalTaskIDs)
	}
}

func TestExecutionFoundationIngestMigratesSyncForwardToAsyncTask(t *testing.T) {
	store := audit.NewMemoryStore()
	var (
		mu       sync.Mutex
		received []audit.Event
	)
	hub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var event audit.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err == nil {
			mu.Lock()
			received = append(received, event)
			mu.Unlock()
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"status":"stored"}`))
	}))
	defer hub.Close()

	handler := newHandlerWithRuntimesAndSigning(
		store,
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{
			Mode:         audit.SyncModeSpoke,
			ClusterID:    "cluster-a",
			HubURL:       hub.URL,
			Token:        "service-internal-demo-token",
			PollInterval: time.Minute,
			FailMode:     audit.SyncFailModeLastKnownGood,
			CacheDir:     t.TempDir(),
		}),
		nil,
	)

	req := httptest.NewRequest(http.MethodPost, "/v1/ingest", bytes.NewBufferString(`{
		"component":"runtime-agent",
		"event_type":"runtime_observation_recorded",
		"decision":"ALLOW",
		"namespace":"acme-prod",
		"workload_kind":"deployment",
		"workload":"payments-api",
		"reasons":["phase1 async migration test"]
	}`))
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected ingest 201, got %d: %s", rec.Code, rec.Body.String())
	}

	waitForPhase1Condition(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		return len(received) == 1
	})

	tasks, err := store.ListEvents(context.Background(), audit.EventFilter{
		Component: "audit-writer",
		EventType: audit.EventTypeExecutionTaskRecorded,
		TenantID:  "acme",
		Limit:     20,
	})
	if err != nil {
		t.Fatalf("ListEvents(tasks) error = %v", err)
	}
	if len(tasks) == 0 {
		t.Fatalf("expected async task events, got %#v", tasks)
	}
	foundCompleted := false
	for _, item := range tasks {
		task, err := audit.UnmarshalExecutionTaskRecord(item.Event)
		if err != nil {
			continue
		}
		if task.TaskType == phase1ExecutionTaskSyncForwardEvent && task.CurrentState == audit.ExecutionTaskStateCompleted {
			foundCompleted = true
			break
		}
	}
	if !foundCompleted {
		t.Fatalf("expected completed sync-forward task, got %#v", tasks)
	}

	traces, err := store.ListEvents(context.Background(), audit.EventFilter{
		Component: "audit-writer",
		EventType: audit.EventTypeExecutionTraceRecorded,
		Limit:     20,
	})
	if err != nil {
		t.Fatalf("ListEvents(traces) error = %v", err)
	}
	if len(traces) < 2 {
		t.Fatalf("expected trace evidence for ingest and async forward, got %#v", traces)
	}
	operations := map[string]bool{}
	for _, item := range traces {
		trace, err := audit.UnmarshalExecutionTraceRecord(item.Event)
		if err != nil {
			continue
		}
		operations[trace.Operation] = true
	}
	if !operations["audit_ingest"] || !operations[phase1ExecutionTaskSyncForwardEvent] {
		t.Fatalf("expected audit_ingest and sync forward trace operations, got %#v", operations)
	}
}

func TestExecutionFoundationTrustRotationDrillHandlers(t *testing.T) {
	handler := newHandlerWithDeps(
		audit.NewMemoryStore(),
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "secret-a"),
	)

	drillReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/trust/rotation-drill?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"purpose":"sync-snapshots",
		"next_signer_mode":"software",
		"next_key_id":"rotated-key-b",
		"software_secret":"secret-b"
	}`))
	drillReq.Header.Set("Authorization", "Bearer operator-demo-token")
	drillReq.Header.Set("Content-Type", "application/json")
	drillRec := httptest.NewRecorder()
	handler.ServeHTTP(drillRec, drillReq)
	if drillRec.Code != http.StatusOK {
		t.Fatalf("expected rotation drill 200, got %d: %s", drillRec.Code, drillRec.Body.String())
	}

	var drill phase1RotationDrillMutationResponse
	if err := json.NewDecoder(drillRec.Body).Decode(&drill); err != nil {
		t.Fatalf("decode rotation drill: %v", err)
	}
	if drill.Drill.CurrentState != "passed" || drill.Drill.CurrentVerification.State != signing.StateVerified || drill.Drill.NextVerification.State != signing.StateVerified {
		t.Fatalf("expected passed rotation drill, got %#v", drill)
	}
	if drill.Drill.RevokedVerification.State == signing.StateVerified {
		t.Fatalf("expected revoked signer verification to fail, got %#v", drill.Drill)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/trust/rotation-drills", nil)
	listReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected rotation drill list 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	var listed phase1RotationDrillListResponse
	if err := json.NewDecoder(listRec.Body).Decode(&listed); err != nil {
		t.Fatalf("decode listed rotation drills: %v", err)
	}
	if len(listed.Items) == 0 || listed.Items[0].DrillID == "" {
		t.Fatalf("expected persisted rotation drill, got %#v", listed)
	}
}

func TestExecutionFoundationProofsHandler(t *testing.T) {
	store := audit.NewMemoryStore()
	var (
		mu       sync.Mutex
		received []audit.Event
	)
	hub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var event audit.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err == nil {
			mu.Lock()
			received = append(received, event)
			mu.Unlock()
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer hub.Close()

	handler := newHandlerWithDeps(
		store,
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{
			Mode:         audit.SyncModeSpoke,
			ClusterID:    "cluster-a",
			HubURL:       hub.URL,
			Token:        "service-internal-demo-token",
			PollInterval: time.Minute,
			FailMode:     audit.SyncFailModeLastKnownGood,
			CacheDir:     t.TempDir(),
		}),
		newTestSoftwareSigningRuntime(t, "secret-a"),
	)

	benchReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/benchmarks/evaluate?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"profile_id":"production_like",
		"observations":[{"family_id":"deploy_gate_admission","profile_id":"production_like","metric_class":"user_facing_latency","metric_name":"p95_latency_ms","unit":"ms","baseline_value":100,"observed_value":100}]
	}`))
	benchReq.Header.Set("Authorization", "Bearer operator-demo-token")
	benchReq.Header.Set("Content-Type", "application/json")
	benchRec := httptest.NewRecorder()
	handler.ServeHTTP(benchRec, benchReq)
	if benchRec.Code != http.StatusOK {
		t.Fatalf("expected benchmark evaluation 200, got %d: %s", benchRec.Code, benchRec.Body.String())
	}

	drillReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/trust/rotation-drill?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"purpose":"sync-snapshots",
		"next_signer_mode":"software",
		"next_key_id":"rotated-key-b",
		"software_secret":"secret-b"
	}`))
	drillReq.Header.Set("Authorization", "Bearer operator-demo-token")
	drillReq.Header.Set("Content-Type", "application/json")
	drillRec := httptest.NewRecorder()
	handler.ServeHTTP(drillRec, drillReq)
	if drillRec.Code != http.StatusOK {
		t.Fatalf("expected rotation drill 200, got %d: %s", drillRec.Code, drillRec.Body.String())
	}

	ingestReq := httptest.NewRequest(http.MethodPost, "/v1/ingest", bytes.NewBufferString(`{
		"component":"runtime-agent",
		"event_type":"runtime_observation_recorded",
		"decision":"ALLOW",
		"namespace":"acme-prod",
		"workload":"checkout"
	}`))
	ingestReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	ingestReq.Header.Set("Content-Type", "application/json")
	ingestRec := httptest.NewRecorder()
	handler.ServeHTTP(ingestRec, ingestReq)
	if ingestRec.Code != http.StatusCreated {
		t.Fatalf("expected ingest 201, got %d: %s", ingestRec.Code, ingestRec.Body.String())
	}

	waitForPhase1Condition(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		return len(received) == 1
	})

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/proofs?tenant_id=acme&environment=prod", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}

	var proofs phase1ExecutionProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != "phase1_operational_proof_active" {
		t.Fatalf("expected active operational proof state, got %#v", proofs)
	}
	if len(proofs.LatestBenchmarks) == 0 || len(proofs.LatestRotationDrills) == 0 || proofs.TraceSummary.TraceCount == 0 {
		t.Fatalf("expected benchmark, drill, and trace proof evidence, got %#v", proofs)
	}
	if len(proofs.BenchmarkArtifacts) == 0 || len(proofs.TraceArtifacts) == 0 || len(proofs.TaskArtifacts) == 0 {
		t.Fatalf("expected proof artifacts, got %#v", proofs)
	}
}

func TestExecutionFoundationProofsRequireBenchmarkEvidence(t *testing.T) {
	store := audit.NewMemoryStore()
	var (
		mu       sync.Mutex
		received []audit.Event
	)
	hub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var event audit.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err == nil {
			mu.Lock()
			received = append(received, event)
			mu.Unlock()
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer hub.Close()

	handler := newHandlerWithDeps(
		store,
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{
			Mode:         audit.SyncModeSpoke,
			ClusterID:    "cluster-a",
			HubURL:       hub.URL,
			Token:        "service-internal-demo-token",
			PollInterval: time.Minute,
			FailMode:     audit.SyncFailModeLastKnownGood,
			CacheDir:     t.TempDir(),
		}),
		newTestSoftwareSigningRuntime(t, "secret-a"),
	)

	drillReq := httptest.NewRequest(http.MethodPost, "/v1/foundation/execution/trust/rotation-drill?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
		"purpose":"sync-snapshots",
		"next_signer_mode":"software",
		"next_key_id":"rotated-key-b",
		"software_secret":"secret-b"
	}`))
	drillReq.Header.Set("Authorization", "Bearer operator-demo-token")
	drillReq.Header.Set("Content-Type", "application/json")
	drillRec := httptest.NewRecorder()
	handler.ServeHTTP(drillRec, drillReq)
	if drillRec.Code != http.StatusOK {
		t.Fatalf("expected rotation drill 200, got %d: %s", drillRec.Code, drillRec.Body.String())
	}

	ingestReq := httptest.NewRequest(http.MethodPost, "/v1/ingest", bytes.NewBufferString(`{
		"component":"runtime-agent",
		"event_type":"runtime_observation_recorded",
		"decision":"ALLOW",
		"namespace":"acme-prod",
		"workload":"checkout"
	}`))
	ingestReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	ingestReq.Header.Set("Content-Type", "application/json")
	ingestRec := httptest.NewRecorder()
	handler.ServeHTTP(ingestRec, ingestReq)
	if ingestRec.Code != http.StatusCreated {
		t.Fatalf("expected ingest 201, got %d: %s", ingestRec.Code, ingestRec.Body.String())
	}

	waitForPhase1Condition(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		return len(received) == 1
	})

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/foundation/execution/proofs?tenant_id=acme&environment=prod", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}

	var proofs phase1ExecutionProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != "phase1_closure_incomplete" {
		t.Fatalf("expected incomplete proof state without benchmark evidence, got %#v", proofs)
	}
}

func mustIngestExecutionTaskRecord(t *testing.T, store audit.Store, requestID string, task audit.ExecutionTaskRecord) {
	t.Helper()

	payload, err := audit.MarshalExecutionTaskRecord(task)
	if err != nil {
		t.Fatalf("MarshalExecutionTaskRecord() error = %v", err)
	}
	if _, err := store.Ingest(context.Background(), audit.Event{
		RequestID:           requestID,
		Component:           "audit-writer",
		EventType:           audit.EventTypeExecutionTaskRecorded,
		Actor:               "seed",
		TraceID:             task.TraceID,
		CorrelationID:       task.CorrelationID,
		DecisionID:          task.DecisionID,
		CausalParent:        task.CausalParent,
		IdempotencyKey:      task.TaskType + ":" + task.TaskID + ":" + task.CurrentState,
		PayloadHash:         task.PayloadHash,
		TenantID:            task.TenantID,
		Environment:         task.Environment,
		Decision:            audit.DecisionAllow,
		Reasons:             []string{"phase 1 async task recorded", task.CurrentState, task.TaskType},
		ExecutionFoundation: payload,
	}); err != nil {
		t.Fatalf("Ingest(task event) error = %v", err)
	}
}

func waitForPhase1Condition(t *testing.T, condition func() bool) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if condition() {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatal("timed out waiting for phase1 async condition")
}
