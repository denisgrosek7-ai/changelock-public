package audit

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
)

type MemoryStore struct {
	mu              sync.Mutex
	records         []StoredEvent
	exceptions      map[string]PolicyException
	nextID          int64
	nextExceptionID int64
	now             func() time.Time
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		exceptions:      map[string]PolicyException{},
		nextID:          1,
		nextExceptionID: 1,
		now:             time.Now,
	}
}

func (s *MemoryStore) Ping(context.Context) error { return nil }
func (s *MemoryStore) Close()                     {}

func (s *MemoryStore) Ingest(_ context.Context, event Event) (StoredEvent, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	event = NormalizeEvent(event, s.now)
	if err := ValidateEvent(event); err != nil {
		return StoredEvent{}, err
	}

	raw, err := json.Marshal(event)
	if err != nil {
		return StoredEvent{}, err
	}

	record := StoredEvent{
		ID:         s.nextID,
		ReceivedAt: s.now().UTC(),
		Event:      event,
		RawEvent:   append(json.RawMessage(nil), raw...),
	}
	s.nextID++
	s.records = append(s.records, record)
	return record, nil
}

func (s *MemoryStore) ListEvents(_ context.Context, filter EventFilter) ([]StoredEvent, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeFilter(filter)
	if err != nil {
		return nil, err
	}

	records := make([]StoredEvent, 0, len(s.records))
	for _, record := range s.records {
		if !matchesFilter(record.Event, filter, true) {
			continue
		}
		records = append(records, cloneStoredEvent(record))
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ReceivedAt.After(records[j].ReceivedAt)
	})
	if len(records) > filter.Limit {
		records = records[:filter.Limit]
	}
	return records, nil
}

func (s *MemoryStore) Summary(_ context.Context, filter EventFilter) (Summary, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeFilter(filter)
	if err != nil {
		return Summary{}, err
	}

	summary := Summary{
		CountsByEventType: map[string]int64{},
		TopDenyReasons:    []ReasonCount{},
	}

	denyCounts := map[string]int64{}
	now := s.now().UTC()
	for _, record := range s.records {
		if !matchesFilter(record.Event, filter, true) {
			continue
		}
		summary.TotalEvents++
		switch record.Decision {
		case DecisionAllow:
			summary.TotalAllow++
		case DecisionDeny:
			summary.TotalDeny++
		case DecisionError:
			summary.TotalError++
		}
		summary.CountsByEventType[record.EventType]++

		if matchesFilter(record.Event, filterWithoutDecision(filter), false) && record.Decision == DecisionDeny {
			for _, reason := range record.Reasons {
				denyCounts[reason]++
			}
		}
		if matchesFilter(record.Event, filterWithoutDecision(filter), false) &&
			record.EventType == EventTypeRuntimeDriftResult &&
			record.Decision == DecisionDeny &&
			record.ReceivedAt.After(now.Add(-24*time.Hour)) {
			summary.RecentRuntimeDriftDeny++
		}
	}

	reasons := make([]ReasonCount, 0, len(denyCounts))
	for reason, count := range denyCounts {
		reasons = append(reasons, ReasonCount{Reason: reason, Count: count})
	}
	sort.Slice(reasons, func(i, j int) bool {
		if reasons[i].Count == reasons[j].Count {
			return reasons[i].Reason < reasons[j].Reason
		}
		return reasons[i].Count > reasons[j].Count
	})
	if len(reasons) > 5 {
		reasons = reasons[:5]
	}
	summary.TopDenyReasons = reasons

	return summary, nil
}

func matchesFilter(event Event, filter EventFilter, includeDecision bool) bool {
	if includeDecision && filter.Decision != "" && event.Decision != filter.Decision {
		return false
	}
	if filter.EventType != "" && event.EventType != filter.EventType {
		return false
	}
	if filter.Component != "" && event.Component != filter.Component {
		return false
	}
	if filter.Repo != "" && event.Repo != filter.Repo {
		return false
	}
	if filter.Environment != "" && event.Environment != filter.Environment {
		return false
	}
	if filter.TenantID != "" && event.TenantID != filter.TenantID {
		return false
	}
	return true
}

func cloneStoredEvent(record StoredEvent) StoredEvent {
	record.Reasons = slices.Clone(record.Reasons)
	record.DriftClasses = slices.Clone(record.DriftClasses)
	if record.RawEvent != nil {
		record.RawEvent = slices.Clone(record.RawEvent)
	}
	return record
}

func (s *MemoryStore) CreateException(_ context.Context, request ExceptionCreateRequest) (PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	request, err := NormalizeExceptionCreateRequest(request, s.now)
	if err != nil {
		return PolicyException{}, err
	}
	if _, exists := s.exceptions[request.ExceptionID]; exists {
		return PolicyException{}, fmt.Errorf("%w: exception_id %q already exists", ErrInvalidException, request.ExceptionID)
	}

	exception := PolicyException{
		ID:            s.nextExceptionID,
		ExceptionID:   request.ExceptionID,
		ExceptionType: request.ExceptionType,
		TenantID:      request.TenantID,
		Environment:   request.Environment,
		Namespace:     request.Namespace,
		Repo:          request.Repo,
		ImageDigest:   request.ImageDigest,
		CVEID:         request.CVEID,
		Reason:        request.Reason,
		TicketID:      request.TicketID,
		ApprovedBy:    request.ApprovedBy,
		CreatedAt:     s.now().UTC(),
		ExpiresAt:     request.ExpiresAt.UTC(),
		Active:        true,
		Metadata:      normalizeMetadata(request.Metadata),
	}
	s.nextExceptionID++
	s.exceptions[exception.ExceptionID] = exception

	return clonePolicyException(exception), nil
}

func (s *MemoryStore) ListExceptions(_ context.Context, filter ExceptionFilter) ([]PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeExceptionFilter(filter)
	if err != nil {
		return nil, err
	}

	exceptions := make([]PolicyException, 0, len(s.exceptions))
	now := s.now().UTC()
	for _, exception := range s.exceptions {
		if !matchesExceptionFilter(exception, filter, now) {
			continue
		}
		exceptions = append(exceptions, clonePolicyException(exception))
	}

	sort.Slice(exceptions, func(i, j int) bool {
		return exceptions[i].CreatedAt.After(exceptions[j].CreatedAt)
	})
	if len(exceptions) > filter.Limit {
		exceptions = exceptions[:filter.Limit]
	}

	return exceptions, nil
}

func (s *MemoryStore) RevokeException(_ context.Context, exceptionID string) (PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	exceptionID = strings.TrimSpace(exceptionID)
	exception, ok := s.exceptions[exceptionID]
	if !ok {
		return PolicyException{}, ErrExceptionNotFound
	}
	exception.Active = false
	s.exceptions[exceptionID] = exception
	return clonePolicyException(exception), nil
}

func (s *MemoryStore) ValidateException(_ context.Context, request ExceptionValidationRequest) (ExceptionValidationResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	request, err := NormalizeExceptionValidationRequest(request)
	if err != nil {
		return ExceptionValidationResult{}, err
	}

	exception, ok := s.exceptions[request.ExceptionID]
	if !ok {
		return ExceptionValidationResult{Valid: false, Reason: "exception not found"}, nil
	}

	matched, reason := exception.Matches(request, s.now().UTC())
	if !matched {
		return ExceptionValidationResult{Valid: false, Reason: reason}, nil
	}

	exceptionCopy := clonePolicyException(exception)
	return ExceptionValidationResult{
		Valid:     true,
		Exception: &exceptionCopy,
	}, nil
}

func (s *MemoryStore) ExceptionReport(_ context.Context, filter ExceptionFilter) (ExceptionReport, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeExceptionFilter(filter)
	if err != nil {
		return ExceptionReport{}, err
	}

	now := s.now().UTC()
	report := ExceptionReport{
		Active:         []PolicyException{},
		RecentUsed:     []StoredEvent{},
		RecentInactive: []PolicyException{},
	}

	for _, exception := range s.exceptions {
		if !matchesExceptionScope(exception, filter) {
			continue
		}
		if exception.IsCurrentlyActive(now) {
			report.Active = append(report.Active, clonePolicyException(exception))
			continue
		}
		report.RecentInactive = append(report.RecentInactive, clonePolicyException(exception))
	}

	sort.Slice(report.Active, func(i, j int) bool {
		return report.Active[i].CreatedAt.After(report.Active[j].CreatedAt)
	})
	sort.Slice(report.RecentInactive, func(i, j int) bool {
		return report.RecentInactive[i].CreatedAt.After(report.RecentInactive[j].CreatedAt)
	})

	for _, record := range s.records {
		if record.EventType != EventTypeExceptionUsed {
			continue
		}
		if !matchesExceptionEvent(record.Event, filter) {
			continue
		}
		report.RecentUsed = append(report.RecentUsed, cloneStoredEvent(record))
	}
	sort.Slice(report.RecentUsed, func(i, j int) bool {
		return report.RecentUsed[i].ReceivedAt.After(report.RecentUsed[j].ReceivedAt)
	})

	if len(report.Active) > filter.Limit {
		report.Active = report.Active[:filter.Limit]
	}
	if len(report.RecentInactive) > filter.Limit {
		report.RecentInactive = report.RecentInactive[:filter.Limit]
	}
	if len(report.RecentUsed) > filter.Limit {
		report.RecentUsed = report.RecentUsed[:filter.Limit]
	}

	return report, nil
}

func matchesExceptionFilter(exception PolicyException, filter ExceptionFilter, now time.Time) bool {
	if filter.Active != nil {
		currentlyActive := exception.IsCurrentlyActive(now)
		if *filter.Active != currentlyActive {
			return false
		}
	}
	return matchesExceptionScope(exception, filter)
}

func matchesExceptionScope(exception PolicyException, filter ExceptionFilter) bool {
	if filter.ExceptionType != "" && exception.ExceptionType != filter.ExceptionType {
		return false
	}
	if filter.TenantID != "" && exception.TenantID != filter.TenantID {
		return false
	}
	if filter.Environment != "" && exception.Environment != filter.Environment {
		return false
	}
	if filter.Namespace != "" && exception.Namespace != filter.Namespace {
		return false
	}
	if filter.Repo != "" && exception.Repo != filter.Repo {
		return false
	}
	if filter.ImageDigest != "" && exception.ImageDigest != filter.ImageDigest {
		return false
	}
	if filter.CVEID != "" && exception.CVEID != filter.CVEID {
		return false
	}
	return true
}

func matchesExceptionEvent(event Event, filter ExceptionFilter) bool {
	if !matchesExceptionScope(PolicyException{
		ExceptionType: event.ExceptionType,
		TenantID:      event.TenantID,
		Environment:   event.Environment,
		Namespace:     event.Namespace,
		Repo:          event.Repo,
		ImageDigest:   event.Digest,
	}, filter) {
		return false
	}
	return true
}
