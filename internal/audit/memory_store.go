package audit

import (
	"context"
	"encoding/json"
	"slices"
	"sort"
	"sync"
	"time"
)

type MemoryStore struct {
	mu      sync.Mutex
	records []StoredEvent
	nextID  int64
	now     func() time.Time
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		nextID: 1,
		now:    time.Now,
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
