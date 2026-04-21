package audit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	ExecutionAsyncTaskSchemaVersion     = "1.execution_async_task.v1"
	ExecutionBenchmarkGateSchemaVersion = "1.execution_benchmark_gate.v1"
	ExecutionTraceRecordSchemaVersion   = "1.execution_trace_record.v1"
	ExecutionRotationDrillSchemaVersion = "1.execution_rotation_drill.v1"

	ExecutionTaskStateQueued          = "queued"
	ExecutionTaskStateRunning         = "running"
	ExecutionTaskStateCompleted       = "completed"
	ExecutionTaskStateFailedRetryable = "failed_retryable"
	ExecutionTaskStateFailedTerminal  = "failed_terminal"
	ExecutionTaskStateDeadLettered    = "dead_lettered"
	ExecutionTaskStateReplayQueued    = "replay_queued"

	ExecutionTaskFailureRetryable         = "retryable"
	ExecutionTaskFailureTransientExternal = "transient_external"
	ExecutionTaskFailureBusinessRule      = "permanent_business_rule_failure"
	ExecutionTaskFailureSchema            = "schema_failure"
	ExecutionTaskFailurePoisonPayload     = "poison_payload"
	ExecutionTaskFailureTrustProvider     = "trust_provider_failure"
)

type ExecutionTaskRecord struct {
	SchemaVersion    string    `json:"schema_version"`
	TaskID           string    `json:"task_id"`
	TaskType         string    `json:"task_type"`
	CurrentState     string    `json:"current_state"`
	SourceComponent  string    `json:"source_component"`
	SourceEventID    string    `json:"source_event_id,omitempty"`
	TenantID         string    `json:"tenant_id,omitempty"`
	Environment      string    `json:"environment,omitempty"`
	QueueClass       string    `json:"queue_class,omitempty"`
	BackpressureTier string    `json:"backpressure_tier,omitempty"`
	TraceID          string    `json:"trace_id,omitempty"`
	CorrelationID    string    `json:"correlation_id,omitempty"`
	DecisionID       string    `json:"decision_id,omitempty"`
	CausalParent     string    `json:"causal_parent,omitempty"`
	IdempotencyKey   string    `json:"idempotency_key,omitempty"`
	PayloadHash      string    `json:"payload_hash,omitempty"`
	TrustContextRef  string    `json:"trust_context_ref,omitempty"`
	FailureClass     string    `json:"failure_class,omitempty"`
	FailureReason    string    `json:"failure_reason,omitempty"`
	ReplayOfTaskID   string    `json:"replay_of_task_id,omitempty"`
	Attempts         int       `json:"attempts,omitempty"`
	MaxAttempts      int       `json:"max_attempts,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Notes            []string  `json:"notes,omitempty"`
}

type ExecutionBenchmarkObservation struct {
	FamilyID      string  `json:"family_id"`
	ProfileID     string  `json:"profile_id"`
	MetricClass   string  `json:"metric_class"`
	MetricName    string  `json:"metric_name"`
	Unit          string  `json:"unit"`
	BaselineValue float64 `json:"baseline_value"`
	ObservedValue float64 `json:"observed_value"`
}

type ExecutionBenchmarkEvaluationRecord struct {
	SchemaVersion  string                         `json:"schema_version"`
	EvaluationID   string                         `json:"evaluation_id"`
	ProfileID      string                         `json:"profile_id"`
	CurrentState   string                         `json:"current_state"`
	OverrideReason string                         `json:"override_reason,omitempty"`
	ApprovedBy     string                         `json:"approved_by,omitempty"`
	Results        []ExecutionBenchmarkGateResult `json:"results,omitempty"`
	ObservedAt     time.Time                      `json:"observed_at"`
}

type ExecutionTraceRecord struct {
	SchemaVersion string            `json:"schema_version"`
	TraceID       string            `json:"trace_id"`
	SpanID        string            `json:"span_id"`
	ParentSpanID  string            `json:"parent_span_id,omitempty"`
	Component     string            `json:"component"`
	Operation     string            `json:"operation"`
	TenantID      string            `json:"tenant_id,omitempty"`
	Environment   string            `json:"environment,omitempty"`
	SubjectRef    string            `json:"subject_ref,omitempty"`
	EventID       string            `json:"event_id,omitempty"`
	DecisionID    string            `json:"decision_id,omitempty"`
	CorrelationID string            `json:"correlation_id,omitempty"`
	Status        string            `json:"status"`
	StartedAt     time.Time         `json:"started_at"`
	EndedAt       time.Time         `json:"ended_at"`
	DurationMs    int64             `json:"duration_ms"`
	Attributes    map[string]string `json:"attributes,omitempty"`
	Notes         []string          `json:"notes,omitempty"`
}

type ExecutionRotationDrillRuntime struct {
	ProviderMode string `json:"provider_mode"`
	KeyID        string `json:"key_id,omitempty"`
}

type ExecutionRotationDrillVerification struct {
	State          string `json:"state"`
	Reason         string `json:"reason,omitempty"`
	VerificationBy string `json:"verification_by,omitempty"`
	LifecycleState string `json:"lifecycle_state,omitempty"`
}

type ExecutionRotationDrillRecord struct {
	SchemaVersion        string                             `json:"schema_version"`
	DrillID              string                             `json:"drill_id"`
	Purpose              string                             `json:"purpose"`
	TenantID             string                             `json:"tenant_id,omitempty"`
	Environment          string                             `json:"environment,omitempty"`
	CurrentRuntime       ExecutionRotationDrillRuntime      `json:"current_runtime"`
	NextRuntime          ExecutionRotationDrillRuntime      `json:"next_runtime"`
	CurrentSignedAt      time.Time                          `json:"current_signed_at"`
	NextSignedAt         time.Time                          `json:"next_signed_at"`
	CurrentVerification  ExecutionRotationDrillVerification `json:"current_verification"`
	NextVerification     ExecutionRotationDrillVerification `json:"next_verification"`
	RevokedVerification  ExecutionRotationDrillVerification `json:"revoked_verification"`
	CurrentState         string                             `json:"current_state"`
	ObservedAt           time.Time                          `json:"observed_at"`
	TrustBoundary        string                             `json:"trust_boundary,omitempty"`
	HistoricalVerifyPath string                             `json:"historical_verify_path,omitempty"`
	Notes                []string                           `json:"notes,omitempty"`
}

type ExecutionBenchmarkGateResult struct {
	FamilyID      string  `json:"family_id"`
	MetricClass   string  `json:"metric_class"`
	MetricName    string  `json:"metric_name"`
	Status        string  `json:"status"`
	ObservedValue float64 `json:"observed_value"`
	BaselineValue float64 `json:"baseline_value"`
	DeltaPercent  float64 `json:"delta_percent"`
	ThresholdPct  float64 `json:"threshold_pct"`
	Summary       string  `json:"summary"`
}

func NormalizeExecutionTaskRecord(task ExecutionTaskRecord, now func() time.Time) ExecutionTaskRecord {
	if now == nil {
		now = time.Now
	}
	if strings.TrimSpace(task.SchemaVersion) == "" {
		task.SchemaVersion = ExecutionAsyncTaskSchemaVersion
	}
	if strings.TrimSpace(task.TaskID) == "" {
		task.TaskID = "task-" + NewRequestID()
	}
	if strings.TrimSpace(task.CurrentState) == "" {
		task.CurrentState = ExecutionTaskStateQueued
	}
	if task.MaxAttempts <= 0 {
		task.MaxAttempts = 3
	}
	if task.CreatedAt.IsZero() {
		task.CreatedAt = now().UTC()
	}
	if task.UpdatedAt.IsZero() {
		task.UpdatedAt = task.CreatedAt
	}
	task.TaskType = strings.TrimSpace(task.TaskType)
	task.SourceComponent = strings.TrimSpace(task.SourceComponent)
	task.SourceEventID = strings.TrimSpace(task.SourceEventID)
	task.TenantID = strings.TrimSpace(task.TenantID)
	task.Environment = strings.TrimSpace(task.Environment)
	task.QueueClass = strings.TrimSpace(task.QueueClass)
	task.BackpressureTier = strings.TrimSpace(task.BackpressureTier)
	task.TraceID = strings.TrimSpace(task.TraceID)
	task.CorrelationID = strings.TrimSpace(task.CorrelationID)
	task.DecisionID = strings.TrimSpace(task.DecisionID)
	task.CausalParent = strings.TrimSpace(task.CausalParent)
	task.IdempotencyKey = strings.TrimSpace(task.IdempotencyKey)
	task.PayloadHash = strings.TrimSpace(task.PayloadHash)
	task.TrustContextRef = strings.TrimSpace(task.TrustContextRef)
	task.FailureClass = strings.TrimSpace(task.FailureClass)
	task.FailureReason = strings.TrimSpace(task.FailureReason)
	task.ReplayOfTaskID = strings.TrimSpace(task.ReplayOfTaskID)
	return task
}

func ValidateExecutionTaskRecord(task ExecutionTaskRecord) error {
	if strings.TrimSpace(task.TaskType) == "" {
		return errors.New("task_type is required")
	}
	if strings.TrimSpace(task.SourceComponent) == "" {
		return errors.New("source_component is required")
	}
	if strings.TrimSpace(task.TraceID) == "" {
		return errors.New("trace_id is required")
	}
	if strings.TrimSpace(task.CorrelationID) == "" {
		return errors.New("correlation_id is required")
	}
	if strings.TrimSpace(task.DecisionID) == "" {
		return errors.New("decision_id is required")
	}
	switch strings.TrimSpace(task.CurrentState) {
	case ExecutionTaskStateQueued,
		ExecutionTaskStateRunning,
		ExecutionTaskStateCompleted,
		ExecutionTaskStateFailedRetryable,
		ExecutionTaskStateFailedTerminal,
		ExecutionTaskStateDeadLettered,
		ExecutionTaskStateReplayQueued:
	default:
		return fmt.Errorf("unsupported current_state %q", task.CurrentState)
	}
	if task.Attempts < 0 {
		return errors.New("attempts cannot be negative")
	}
	if task.MaxAttempts <= 0 {
		return errors.New("max_attempts must be positive")
	}
	return nil
}

func MarshalExecutionTaskRecord(task ExecutionTaskRecord) (json.RawMessage, error) {
	payload, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(payload), nil
}

func UnmarshalExecutionTaskRecord(event Event) (ExecutionTaskRecord, error) {
	var task ExecutionTaskRecord
	if len(event.ExecutionFoundation) == 0 || string(event.ExecutionFoundation) == "null" {
		return task, errors.New("execution_foundation payload is missing")
	}
	if err := json.Unmarshal(event.ExecutionFoundation, &task); err != nil {
		return task, err
	}
	return task, nil
}

func NormalizeExecutionTraceRecord(record ExecutionTraceRecord, now func() time.Time) ExecutionTraceRecord {
	if now == nil {
		now = time.Now
	}
	if strings.TrimSpace(record.SchemaVersion) == "" {
		record.SchemaVersion = ExecutionTraceRecordSchemaVersion
	}
	if strings.TrimSpace(record.SpanID) == "" {
		record.SpanID = "span-" + NewRequestID()
	}
	if record.StartedAt.IsZero() {
		record.StartedAt = now().UTC()
	}
	if record.EndedAt.IsZero() {
		record.EndedAt = record.StartedAt
	}
	if record.DurationMs == 0 {
		record.DurationMs = record.EndedAt.Sub(record.StartedAt).Milliseconds()
	}
	record.TraceID = strings.TrimSpace(record.TraceID)
	record.ParentSpanID = strings.TrimSpace(record.ParentSpanID)
	record.Component = strings.TrimSpace(record.Component)
	record.Operation = strings.TrimSpace(record.Operation)
	record.TenantID = strings.TrimSpace(record.TenantID)
	record.Environment = strings.TrimSpace(record.Environment)
	record.SubjectRef = strings.TrimSpace(record.SubjectRef)
	record.EventID = strings.TrimSpace(record.EventID)
	record.DecisionID = strings.TrimSpace(record.DecisionID)
	record.CorrelationID = strings.TrimSpace(record.CorrelationID)
	record.Status = strings.TrimSpace(record.Status)
	return record
}

func ValidateExecutionTraceRecord(record ExecutionTraceRecord) error {
	if strings.TrimSpace(record.TraceID) == "" {
		return errors.New("trace_id is required")
	}
	if strings.TrimSpace(record.SpanID) == "" {
		return errors.New("span_id is required")
	}
	if strings.TrimSpace(record.Component) == "" {
		return errors.New("component is required")
	}
	if strings.TrimSpace(record.Operation) == "" {
		return errors.New("operation is required")
	}
	if strings.TrimSpace(record.Status) == "" {
		return errors.New("status is required")
	}
	if record.EndedAt.Before(record.StartedAt) {
		return errors.New("ended_at cannot be before started_at")
	}
	return nil
}

func MarshalExecutionTraceRecord(record ExecutionTraceRecord) (json.RawMessage, error) {
	payload, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(payload), nil
}

func UnmarshalExecutionTraceRecord(event Event) (ExecutionTraceRecord, error) {
	var record ExecutionTraceRecord
	if len(event.ExecutionFoundation) == 0 || string(event.ExecutionFoundation) == "null" {
		return record, errors.New("execution_foundation payload is missing")
	}
	if err := json.Unmarshal(event.ExecutionFoundation, &record); err != nil {
		return record, err
	}
	return record, nil
}

func NormalizeExecutionRotationDrillRecord(record ExecutionRotationDrillRecord, now func() time.Time) ExecutionRotationDrillRecord {
	if now == nil {
		now = time.Now
	}
	if strings.TrimSpace(record.SchemaVersion) == "" {
		record.SchemaVersion = ExecutionRotationDrillSchemaVersion
	}
	if strings.TrimSpace(record.DrillID) == "" {
		record.DrillID = "rotation-drill-" + NewRequestID()
	}
	if record.ObservedAt.IsZero() {
		record.ObservedAt = now().UTC()
	}
	record.Purpose = strings.TrimSpace(record.Purpose)
	record.TenantID = strings.TrimSpace(record.TenantID)
	record.Environment = strings.TrimSpace(record.Environment)
	record.CurrentState = strings.TrimSpace(record.CurrentState)
	record.TrustBoundary = strings.TrimSpace(record.TrustBoundary)
	record.HistoricalVerifyPath = strings.TrimSpace(record.HistoricalVerifyPath)
	return record
}

func ValidateExecutionRotationDrillRecord(record ExecutionRotationDrillRecord) error {
	if strings.TrimSpace(record.DrillID) == "" {
		return errors.New("drill_id is required")
	}
	if strings.TrimSpace(record.Purpose) == "" {
		return errors.New("purpose is required")
	}
	if strings.TrimSpace(record.CurrentRuntime.ProviderMode) == "" {
		return errors.New("current_runtime.provider_mode is required")
	}
	if strings.TrimSpace(record.NextRuntime.ProviderMode) == "" {
		return errors.New("next_runtime.provider_mode is required")
	}
	if strings.TrimSpace(record.CurrentState) == "" {
		return errors.New("current_state is required")
	}
	return nil
}

func MarshalExecutionRotationDrillRecord(record ExecutionRotationDrillRecord) (json.RawMessage, error) {
	payload, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(payload), nil
}

func UnmarshalExecutionRotationDrillRecord(event Event) (ExecutionRotationDrillRecord, error) {
	var record ExecutionRotationDrillRecord
	if len(event.ExecutionFoundation) == 0 || string(event.ExecutionFoundation) == "null" {
		return record, errors.New("execution_foundation payload is missing")
	}
	if err := json.Unmarshal(event.ExecutionFoundation, &record); err != nil {
		return record, err
	}
	return record, nil
}
