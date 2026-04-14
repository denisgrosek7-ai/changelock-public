package audit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrInvalidEvent  = errors.New("invalid audit event")
	ErrInvalidFilter = errors.New("invalid audit filter")
)

type Store interface {
	Ingest(ctx context.Context, event Event) (StoredEvent, error)
	ListEvents(ctx context.Context, filter EventFilter) ([]StoredEvent, error)
	Summary(ctx context.Context, filter EventFilter) (Summary, error)
	Ping(ctx context.Context) error
	Close()
}

type StoredEvent struct {
	ID         int64     `json:"id"`
	ReceivedAt time.Time `json:"received_at"`
	Event
	RawEvent json.RawMessage `json:"raw_event,omitempty"`
}

type EventFilter struct {
	Decision    string
	EventType   string
	Component   string
	Repo        string
	Environment string
	TenantID    string
	Limit       int
}

type Summary struct {
	TotalEvents            int64            `json:"total_events"`
	TotalAllow             int64            `json:"total_allow"`
	TotalDeny              int64            `json:"total_deny"`
	TotalError             int64            `json:"total_error"`
	CountsByEventType      map[string]int64 `json:"counts_by_event_type"`
	TopDenyReasons         []ReasonCount    `json:"top_deny_reasons"`
	RecentRuntimeDriftDeny int64            `json:"recent_runtime_drift_deny"`
}

type ReasonCount struct {
	Reason string `json:"reason"`
	Count  int64  `json:"count"`
}

func NormalizeEvent(event Event, now func() time.Time) Event {
	if event.RequestID == "" {
		event.RequestID = NewRequestID()
	}
	if event.Timestamp.IsZero() {
		if now == nil {
			now = time.Now
		}
		event.Timestamp = now().UTC()
	}
	if event.TenantID == "" {
		event.TenantID = TenantFromNamespace(event.Namespace)
	}
	if event.Environment == "" {
		event.Environment = EnvironmentFromNamespace(event.Namespace)
	}
	if event.Digest == "" {
		event.Digest = DigestFromImage(event.Image)
	}
	return EnsureDecisionHash(event)
}

func ValidateEvent(event Event) error {
	if strings.TrimSpace(event.Component) == "" {
		return fmt.Errorf("%w: component is required", ErrInvalidEvent)
	}
	if strings.TrimSpace(event.EventType) == "" {
		return fmt.Errorf("%w: event_type is required", ErrInvalidEvent)
	}
	switch event.Decision {
	case DecisionAllow, DecisionDeny, DecisionError:
	default:
		return fmt.Errorf("%w: decision must be one of %q, %q, %q", ErrInvalidEvent, DecisionAllow, DecisionDeny, DecisionError)
	}
	return nil
}

func NormalizeFilter(filter EventFilter) (EventFilter, error) {
	filter.Decision = strings.TrimSpace(filter.Decision)
	filter.EventType = strings.TrimSpace(filter.EventType)
	filter.Component = strings.TrimSpace(filter.Component)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.TenantID = strings.TrimSpace(filter.TenantID)

	switch filter.Decision {
	case "", DecisionAllow, DecisionDeny, DecisionError:
	default:
		return filter, fmt.Errorf("%w: unsupported decision %q", ErrInvalidFilter, filter.Decision)
	}

	if filter.Limit <= 0 {
		filter.Limit = 50
	}
	if filter.Limit > 500 {
		filter.Limit = 500
	}

	return filter, nil
}
