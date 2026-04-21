package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"

	"github.com/denisgrosek/changelock/internal/identity"
)

const (
	ExecutionEventSchemaVersion = "1.execution_event.v1"
)

var executionEnvelopeFields = []string{
	"schema_version",
	"event_id",
	"trace_id",
	"correlation_id",
	"decision_id",
	"causal_parent",
	"idempotency_key",
	"payload_hash",
}

func EnsureExecutionEnvelope(event Event) Event {
	if strings.TrimSpace(event.SchemaVersion) == "" {
		event.SchemaVersion = ExecutionEventSchemaVersion
	}

	if strings.TrimSpace(event.CorrelationID) == "" {
		event.CorrelationID = firstNonEmptyEnvelope(
			event.IncidentIdentityKey,
			event.IncidentID,
			event.RecommendationID,
			event.RequestID,
		)
	}
	if strings.TrimSpace(event.TraceID) == "" {
		event.TraceID = firstNonEmptyEnvelope(event.CorrelationID, event.RequestID)
	}
	if strings.TrimSpace(event.DecisionID) == "" {
		event.DecisionID = firstNonEmptyEnvelope(event.DecisionHash, event.RequestID)
	}
	if strings.TrimSpace(event.EventID) == "" {
		event.EventID = computeExecutionEventHash(
			"event_id",
			event.RequestID,
			event.Component,
			event.EventType,
			event.Decision,
			event.DecisionHash,
			event.Digest,
			event.Workload,
			event.Namespace,
			event.Repo,
			event.Environment,
			event.Timestamp.UTC().Format(timeLayoutSecondPrecision),
		)
	}
	if strings.TrimSpace(event.IdempotencyKey) == "" {
		event.IdempotencyKey = computeExecutionEventHash(
			"idempotency",
			event.EventID,
			event.Component,
			event.EventType,
			event.CorrelationID,
			event.CausalParent,
		)
	}
	if strings.TrimSpace(event.PayloadHash) == "" {
		event.PayloadHash = computeExecutionEventHash(
			"payload",
			event.Component,
			event.EventType,
			event.Decision,
			event.RequestID,
			event.Repo,
			event.Environment,
			event.Namespace,
			event.Workload,
			event.Digest,
			event.PolicyBundleHash,
			event.ExceptionID,
			event.IncidentID,
			event.RecommendationID,
		)
	}

	return event
}

func ExecutionEnvelopeFieldSet() []string {
	return append([]string(nil), executionEnvelopeFields...)
}

func computeExecutionEventHash(kind string, values ...string) string {
	hasher := sha256.New()
	io.WriteString(hasher, strings.TrimSpace(kind))
	hasher.Write([]byte{0})
	for _, value := range values {
		io.WriteString(hasher, strings.TrimSpace(value))
		hasher.Write([]byte{0})
	}
	return "sha256:" + hex.EncodeToString(hasher.Sum(nil))
}

func firstNonEmptyEnvelope(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

const timeLayoutSecondPrecision = "2006-01-02T15:04:05Z07:00"

func CanonicalDecisionID(event Event) string {
	return identity.DecisionHash(identity.DecisionInput{
		PolicyBundleHash: event.PolicyBundleHash,
		ImageDigest:      event.Digest,
		RequestID:        event.RequestID,
		Decision:         event.Decision,
		Component:        event.Component,
		Repo:             event.Repo,
		Environment:      event.Environment,
	})
}
