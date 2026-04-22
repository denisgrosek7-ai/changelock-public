package connectors

import (
	"strings"
	"time"
)

const (
	ReconciliationSchemaVersion = "4.enterprise_connector_reconciliation.v1"

	HealthHealthy  = "healthy"
	HealthDegraded = "degraded"
	HealthFailing  = "failing"

	StateSynced                           = "synced"
	StateAwaitingExternalReconciliation   = "awaiting_external_reconciliation"
	StateExternalClosurePendingValidation = "external_closure_pending_validation"
	StateConnectorDegraded                = "connector_degraded_core_preserved"
	StateReopenedForValidation            = "reopened_for_validation"
)

type ConnectorHealth struct {
	CurrentState  string `json:"current_state,omitempty"`
	LastError     string `json:"last_error,omitempty"`
	RetryCount    int    `json:"retry_count,omitempty"`
	RateLimited   bool   `json:"rate_limited,omitempty"`
	AsyncIsolated bool   `json:"async_isolated,omitempty"`
	Replayable    bool   `json:"replayable,omitempty"`
	HealthVisible bool   `json:"health_visible,omitempty"`
}

type ReconciliationInput struct {
	WorkflowID      string          `json:"workflow_id,omitempty"`
	SubjectRef      string          `json:"subject_ref,omitempty"`
	ConnectorSystem string          `json:"connector_system,omitempty"`
	ConnectorRef    string          `json:"connector_ref,omitempty"`
	ObjectType      string          `json:"object_type,omitempty"`
	InternalState   string          `json:"internal_state,omitempty"`
	ExternalState   string          `json:"external_state,omitempty"`
	ValidationState string          `json:"validation_state,omitempty"`
	ApprovalState   string          `json:"approval_state,omitempty"`
	Health          ConnectorHealth `json:"health"`
	EvidenceRefs    []string        `json:"evidence_refs,omitempty"`
	ReasonCodes     []string        `json:"reason_codes,omitempty"`
}

type ReconciliationRecord struct {
	SchemaVersion      string          `json:"schema_version"`
	WorkflowID         string          `json:"workflow_id,omitempty"`
	SubjectRef         string          `json:"subject_ref,omitempty"`
	ConnectorSystem    string          `json:"connector_system,omitempty"`
	ConnectorRef       string          `json:"connector_ref,omitempty"`
	ObjectType         string          `json:"object_type,omitempty"`
	CurrentState       string          `json:"current_state"`
	CanonicalState     string          `json:"canonical_state"`
	InternalState      string          `json:"internal_state,omitempty"`
	ExternalState      string          `json:"external_state,omitempty"`
	ValidationState    string          `json:"validation_state,omitempty"`
	ApprovalState      string          `json:"approval_state,omitempty"`
	Health             ConnectorHealth `json:"health"`
	SafeToAutoClose    bool            `json:"safe_to_auto_close"`
	ConflictResolution string          `json:"conflict_resolution,omitempty"`
	EvidenceRefs       []string        `json:"evidence_refs,omitempty"`
	ReasonCodes        []string        `json:"reason_codes,omitempty"`
	ObservedAt         time.Time       `json:"observed_at"`
	Limitations        []string        `json:"limitations,omitempty"`
}

func Reconcile(input ReconciliationInput, now func() time.Time) ReconciliationRecord {
	if now == nil {
		now = time.Now
	}
	input.WorkflowID = strings.TrimSpace(input.WorkflowID)
	input.SubjectRef = strings.TrimSpace(input.SubjectRef)
	input.ConnectorSystem = strings.TrimSpace(input.ConnectorSystem)
	input.ConnectorRef = strings.TrimSpace(input.ConnectorRef)
	input.ObjectType = strings.TrimSpace(input.ObjectType)
	input.InternalState = strings.ToLower(strings.TrimSpace(input.InternalState))
	input.ExternalState = strings.ToLower(strings.TrimSpace(input.ExternalState))
	input.ValidationState = strings.ToLower(strings.TrimSpace(input.ValidationState))
	input.ApprovalState = strings.ToLower(strings.TrimSpace(input.ApprovalState))
	input.Health = normalizeHealth(input.Health)

	reasons := append([]string{}, input.ReasonCodes...)
	currentState := StateSynced
	canonicalState := input.InternalState
	safeToAutoClose := false
	conflictResolution := "internal_state_canonical"

	switch {
	case input.Health.CurrentState == HealthFailing:
		currentState = StateConnectorDegraded
		reasons = append(reasons, "connector_failure_does_not_block_core")
	case externalClosed(input.ExternalState) && input.ValidationState != "verified":
		currentState = StateExternalClosurePendingValidation
		canonicalState = "under_validation"
		reasons = append(reasons, "external_closure_requires_validation")
	case externalClosed(input.ExternalState) && input.ValidationState == "verified":
		currentState = StateSynced
		safeToAutoClose = true
		reasons = append(reasons, "validated_closure_reconciled")
	case input.InternalState == "resolved" && !externalClosed(input.ExternalState):
		currentState = StateAwaitingExternalReconciliation
		reasons = append(reasons, "external_projection_stale")
	case input.InternalState == "reopened" && externalClosed(input.ExternalState):
		currentState = StateReopenedForValidation
		reasons = append(reasons, "local_reopen_overrides_external_closure")
	}
	if input.Health.CurrentState == HealthDegraded {
		reasons = append(reasons, "connector_health_degraded")
	}

	return ReconciliationRecord{
		SchemaVersion:      ReconciliationSchemaVersion,
		WorkflowID:         input.WorkflowID,
		SubjectRef:         input.SubjectRef,
		ConnectorSystem:    input.ConnectorSystem,
		ConnectorRef:       input.ConnectorRef,
		ObjectType:         input.ObjectType,
		CurrentState:       currentState,
		CanonicalState:     canonicalState,
		InternalState:      input.InternalState,
		ExternalState:      input.ExternalState,
		ValidationState:    input.ValidationState,
		ApprovalState:      input.ApprovalState,
		Health:             input.Health,
		SafeToAutoClose:    safeToAutoClose,
		ConflictResolution: conflictResolution,
		EvidenceRefs:       uniqueStrings(input.EvidenceRefs),
		ReasonCodes:        uniqueStrings(reasons),
		ObservedAt:         now().UTC(),
		Limitations: []string{
			"Connector state is a workflow projection and reconciliation layer; it does not replace canonical technical truth stored in audit evidence.",
			"External closure is not accepted as resolution without matching validation evidence where validation is required.",
		},
	}
}

func normalizeHealth(health ConnectorHealth) ConnectorHealth {
	health.CurrentState = strings.ToLower(strings.TrimSpace(health.CurrentState))
	switch health.CurrentState {
	case HealthHealthy, HealthDegraded, HealthFailing:
	default:
		health.CurrentState = HealthHealthy
	}
	health.LastError = strings.TrimSpace(health.LastError)
	if !health.AsyncIsolated {
		health.AsyncIsolated = true
	}
	if !health.Replayable {
		health.Replayable = true
	}
	if !health.HealthVisible {
		health.HealthVisible = true
	}
	return health
}

func externalClosed(value string) bool {
	switch value {
	case "resolved", "closed", "done", "validated":
		return true
	default:
		return false
	}
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		items = append(items, value)
	}
	if len(items) == 0 {
		return nil
	}
	return items
}
