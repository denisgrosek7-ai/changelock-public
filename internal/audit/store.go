package audit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/signing"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

var (
	ErrInvalidEvent  = errors.New("invalid audit event")
	ErrInvalidFilter = errors.New("invalid audit filter")
)

type Store interface {
	Ingest(ctx context.Context, event Event) (StoredEvent, error)
	ListEvents(ctx context.Context, filter EventFilter) ([]StoredEvent, error)
	Summary(ctx context.Context, filter EventFilter) (Summary, error)
	IngestSBOM(ctx context.Context, request SBOMIngestRequest) (SBOMIngestResult, error)
	GetSBOMImage(ctx context.Context, imageDigest string, limit int) (SBOMImageResponse, error)
	SearchSBOMComponents(ctx context.Context, filter SBOMComponentSearchFilter) ([]SBOMComponent, error)
	RecordVulnerabilityScan(ctx context.Context, request VulnerabilityScanRequest) (VulnerabilityScanIngestResult, error)
	ListActiveVulnerabilities(ctx context.Context, filter VulnerabilityActiveFilter) ([]VulnerabilityFinding, error)
	VulnerabilityBlastRadius(ctx context.Context, filter VulnerabilityBlastRadiusFilter) (VulnerabilityBlastRadiusResponse, error)
	VulnerabilityTimeline(ctx context.Context, filter VulnerabilityTimelineFilter) (VulnerabilityTimelineResponse, error)
	ListVulnerabilityDecisions(ctx context.Context, filter VulnerabilityDecisionFilter) ([]VulnerabilityDecision, error)
	GetVulnerabilityDecision(ctx context.Context, decisionID int64) (VulnerabilityDecision, error)
	CreateVulnerabilityDecision(ctx context.Context, request VulnerabilityDecisionCreateRequest, decidedBy string) (VulnerabilityDecision, error)
	DeactivateVulnerabilityDecision(ctx context.Context, decisionID int64) (VulnerabilityDecision, error)
	ListVEXStatements(ctx context.Context, filter internalvex.Filter) ([]internalvex.Statement, error)
	GetVEXStatement(ctx context.Context, statementID int64) (internalvex.Statement, error)
	CreateVEXStatement(ctx context.Context, request internalvex.CreateRequest, createdBy string) (internalvex.Statement, error)
	RevokeVEXStatement(ctx context.Context, statementID int64, revokedBy string) (internalvex.Statement, error)
	ListActiveDigests(ctx context.Context, windowDays int, limit int) ([]ActiveDigestRef, error)
	LookupDigestScopes(ctx context.Context, imageDigest string, limit int) ([]ActiveWorkloadRef, error)
	CreateException(ctx context.Context, request ExceptionCreateRequest) (PolicyException, error)
	RequestException(ctx context.Context, request ExceptionCreateRequest, requestedBy string, requesterRole string) (PolicyException, error)
	GetException(ctx context.Context, exceptionID string) (PolicyException, error)
	ListExceptions(ctx context.Context, filter ExceptionFilter) ([]PolicyException, error)
	ApproveException(ctx context.Context, exceptionID string, approvedBy string, approverRole string) (PolicyException, error)
	RejectException(ctx context.Context, exceptionID string, reason string, rejectedBy string, rejectorRole string) (PolicyException, error)
	RevokeException(ctx context.Context, exceptionID string) (PolicyException, error)
	SetExceptionSignature(ctx context.Context, exceptionID string, envelope *signing.Envelope) (PolicyException, error)
	ValidateException(ctx context.Context, request ExceptionValidationRequest) (ExceptionValidationResult, error)
	ExceptionReport(ctx context.Context, filter ExceptionFilter) (ExceptionReport, error)
	ListApprovalLogs(ctx context.Context, exceptionID string, limit int) ([]ApprovalLog, error)
	Trends(ctx context.Context, filter TrendsFilter) (TrendsResponse, error)
	TopViolators(ctx context.Context, filter TopViolatorsFilter) (TopViolatorsResponse, error)
	DriftStats(ctx context.Context, filter DriftStatsFilter) (DriftStatsResponse, error)
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
	ClusterID   string
	Repo        string
	Environment string
	TenantID    string
	Since       *time.Time
	Until       *time.Time
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
	event = EnsureDecisionHash(event)
	return EnsureExecutionEnvelope(event)
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
	if strings.TrimSpace(event.SchemaVersion) == "" {
		return fmt.Errorf("%w: schema_version is required", ErrInvalidEvent)
	}
	if strings.TrimSpace(event.EventID) == "" {
		return fmt.Errorf("%w: event_id is required", ErrInvalidEvent)
	}
	if strings.TrimSpace(event.TraceID) == "" {
		return fmt.Errorf("%w: trace_id is required", ErrInvalidEvent)
	}
	if strings.TrimSpace(event.CorrelationID) == "" {
		return fmt.Errorf("%w: correlation_id is required", ErrInvalidEvent)
	}
	if strings.TrimSpace(event.DecisionID) == "" {
		return fmt.Errorf("%w: decision_id is required", ErrInvalidEvent)
	}
	if strings.TrimSpace(event.IdempotencyKey) == "" {
		return fmt.Errorf("%w: idempotency_key is required", ErrInvalidEvent)
	}
	if strings.TrimSpace(event.PayloadHash) == "" {
		return fmt.Errorf("%w: payload_hash is required", ErrInvalidEvent)
	}
	return nil
}

func NormalizeFilter(filter EventFilter) (EventFilter, error) {
	filter.Decision = strings.TrimSpace(filter.Decision)
	filter.EventType = strings.TrimSpace(filter.EventType)
	filter.Component = strings.TrimSpace(filter.Component)
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
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
	if filter.Limit > 5000 {
		filter.Limit = 5000
	}

	if filter.Since != nil {
		value := filter.Since.UTC()
		filter.Since = &value
	}
	if filter.Until != nil {
		value := filter.Until.UTC()
		filter.Until = &value
	}
	if filter.Since != nil && filter.Until != nil && filter.Since.After(*filter.Until) {
		return filter, fmt.Errorf("%w: since must be before until", ErrInvalidFilter)
	}

	return filter, nil
}
