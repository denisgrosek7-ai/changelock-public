package audit

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type MemoryStore struct {
	mu                  sync.Mutex
	records             []StoredEvent
	exceptions          map[string]PolicyException
	approvalLogs        []ApprovalLog
	sbomDocuments       map[string]SBOMDocument
	sbomComponents      []SBOMComponent
	scanRuns            []VulnerabilityScanRun
	findings            map[string]VulnerabilityFinding
	decisions           map[int64]VulnerabilityDecision
	nextID              int64
	nextExceptionID     int64
	nextApprovalLogID   int64
	nextSBOMDocumentID  int64
	nextSBOMComponentID int64
	nextScanRunID       int64
	nextFindingID       int64
	nextDecisionID      int64
	now                 func() time.Time
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		exceptions:          map[string]PolicyException{},
		approvalLogs:        []ApprovalLog{},
		sbomDocuments:       map[string]SBOMDocument{},
		sbomComponents:      []SBOMComponent{},
		scanRuns:            []VulnerabilityScanRun{},
		findings:            map[string]VulnerabilityFinding{},
		decisions:           map[int64]VulnerabilityDecision{},
		nextID:              1,
		nextExceptionID:     1,
		nextApprovalLogID:   1,
		nextSBOMDocumentID:  1,
		nextSBOMComponentID: 1,
		nextScanRunID:       1,
		nextFindingID:       1,
		nextDecisionID:      1,
		now:                 time.Now,
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

	now := s.now().UTC()
	exception := PolicyException{
		ID:            s.nextExceptionID,
		ExceptionID:   request.ExceptionID,
		ExceptionType: request.ExceptionType,
		Status:        ExceptionStatusApproved,
		TenantID:      request.TenantID,
		Environment:   request.Environment,
		Namespace:     request.Namespace,
		Repo:          request.Repo,
		ImageDigest:   request.ImageDigest,
		CVEID:         request.CVEID,
		Reason:        request.Reason,
		TicketID:      request.TicketID,
		RequestedBy:   request.ApprovedBy,
		RequestedAt:   timePointer(now),
		ApprovedBy:    request.ApprovedBy,
		ApprovedAt:    timePointer(now),
		CreatedAt:     now,
		ExpiresAt:     request.ExpiresAt.UTC(),
		Active:        true,
		LastUpdatedAt: timePointer(now),
		Metadata:      normalizeMetadata(request.Metadata),
	}
	s.nextExceptionID++
	s.exceptions[exception.ExceptionID] = exception
	s.appendApprovalLogLocked(exception.ExceptionID, ApprovalActionApproved, exception.ApprovedBy, "", request.Reason, request.Metadata)

	return clonePolicyException(exception), nil
}

func (s *MemoryStore) RequestException(_ context.Context, request ExceptionCreateRequest, requestedBy string, requesterRole string) (PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	request, err := NormalizeExceptionCreateRequest(request, s.now)
	if err != nil {
		return PolicyException{}, err
	}
	if _, exists := s.exceptions[request.ExceptionID]; exists {
		return PolicyException{}, fmt.Errorf("%w: exception_id %q already exists", ErrInvalidException, request.ExceptionID)
	}

	now := s.now().UTC()
	exception := PolicyException{
		ID:            s.nextExceptionID,
		ExceptionID:   request.ExceptionID,
		ExceptionType: request.ExceptionType,
		Status:        ExceptionStatusPending,
		TenantID:      request.TenantID,
		Environment:   request.Environment,
		Namespace:     request.Namespace,
		Repo:          request.Repo,
		ImageDigest:   request.ImageDigest,
		CVEID:         request.CVEID,
		Reason:        request.Reason,
		TicketID:      request.TicketID,
		RequestedBy:   strings.TrimSpace(requestedBy),
		RequestedAt:   timePointer(now),
		CreatedAt:     now,
		ExpiresAt:     request.ExpiresAt.UTC(),
		Active:        false,
		LastUpdatedAt: timePointer(now),
		Metadata:      normalizeMetadata(request.Metadata),
	}
	s.nextExceptionID++
	s.exceptions[exception.ExceptionID] = exception
	s.appendApprovalLogLocked(exception.ExceptionID, ApprovalActionRequested, requestedBy, requesterRole, request.Reason, request.Metadata)

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
		exceptions = append(exceptions, exception.WithEffectiveStatus(now))
	}

	sort.Slice(exceptions, func(i, j int) bool {
		return exceptions[i].CreatedAt.After(exceptions[j].CreatedAt)
	})
	if len(exceptions) > filter.Limit {
		exceptions = exceptions[:filter.Limit]
	}

	return exceptions, nil
}

func (s *MemoryStore) GetException(_ context.Context, exceptionID string) (PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	exception, ok := s.exceptions[strings.TrimSpace(exceptionID)]
	if !ok {
		return PolicyException{}, ErrExceptionNotFound
	}
	return exception.WithEffectiveStatus(s.now().UTC()), nil
}

func (s *MemoryStore) ReplaceApprovedExceptions(_ context.Context, exceptions []SyncedException) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.exceptions = map[string]PolicyException{}
	s.approvalLogs = nil
	s.nextExceptionID = 1
	s.nextApprovalLogID = 1

	now := s.now().UTC()
	for _, synced := range exceptions {
		exception := synced.ToPolicyException(now, s.nextExceptionID)
		s.exceptions[exception.ExceptionID] = exception
		s.nextExceptionID++
	}

	return nil
}

func (s *MemoryStore) ApproveException(_ context.Context, exceptionID string, approvedBy string, approverRole string) (PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	exception, ok := s.exceptions[strings.TrimSpace(exceptionID)]
	if !ok {
		return PolicyException{}, ErrExceptionNotFound
	}
	now := s.now().UTC()
	if status := exception.EffectiveStatus(now); status != ExceptionStatusPending {
		return PolicyException{}, fmt.Errorf("%w: only pending exceptions can be approved", ErrInvalidException)
	}

	exception.Status = ExceptionStatusApproved
	exception.Active = true
	exception.ApprovedBy = strings.TrimSpace(approvedBy)
	exception.ApprovedAt = timePointer(now)
	exception.LastUpdatedAt = timePointer(now)
	s.exceptions[exception.ExceptionID] = exception
	s.appendApprovalLogLocked(exception.ExceptionID, ApprovalActionApproved, approvedBy, approverRole, exception.Reason, nil)

	return exception.WithEffectiveStatus(now), nil
}

func (s *MemoryStore) RejectException(_ context.Context, exceptionID string, reason string, rejectedBy string, rejectorRole string) (PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reason = strings.TrimSpace(reason)
	if reason == "" {
		return PolicyException{}, fmt.Errorf("%w: rejection reason is required", ErrInvalidException)
	}
	exception, ok := s.exceptions[strings.TrimSpace(exceptionID)]
	if !ok {
		return PolicyException{}, ErrExceptionNotFound
	}
	now := s.now().UTC()
	if status := exception.EffectiveStatus(now); status != ExceptionStatusPending {
		return PolicyException{}, fmt.Errorf("%w: only pending exceptions can be rejected", ErrInvalidException)
	}

	exception.Status = ExceptionStatusRejected
	exception.Active = false
	exception.RejectedBy = strings.TrimSpace(rejectedBy)
	exception.RejectedAt = timePointer(now)
	exception.RejectionReason = reason
	exception.LastUpdatedAt = timePointer(now)
	s.exceptions[exception.ExceptionID] = exception
	s.appendApprovalLogLocked(exception.ExceptionID, ApprovalActionRejected, rejectedBy, rejectorRole, reason, nil)

	return exception.WithEffectiveStatus(now), nil
}

func (s *MemoryStore) RevokeException(_ context.Context, exceptionID string) (PolicyException, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	exceptionID = strings.TrimSpace(exceptionID)
	exception, ok := s.exceptions[exceptionID]
	if !ok {
		return PolicyException{}, ErrExceptionNotFound
	}
	now := s.now().UTC()
	exception.Active = false
	exception.Status = ExceptionStatusRevoked
	exception.LastUpdatedAt = timePointer(now)
	s.exceptions[exceptionID] = exception
	s.appendApprovalLogLocked(exception.ExceptionID, ApprovalActionRevoked, exception.ApprovedBy, "", "exception revoked", nil)
	return exception.WithEffectiveStatus(now), nil
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
		s.appendApprovalLogLocked(request.ExceptionID, ApprovalActionValidationFailed, "", "", "exception not found", nil)
		return ExceptionValidationResult{Valid: false, Reason: "exception not found"}, nil
	}

	now := s.now().UTC()
	matched, reason := exception.Matches(request, now)
	if !matched {
		s.appendApprovalLogLocked(exception.ExceptionID, ApprovalActionValidationFailed, "", "", reason, nil)
		return ExceptionValidationResult{Valid: false, Reason: reason}, nil
	}

	s.appendApprovalLogLocked(exception.ExceptionID, ApprovalActionUsed, "", "", "exception used", nil)
	exceptionCopy := exception.WithEffectiveStatus(now)
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
		Pending:        []PolicyException{},
		Rejected:       []PolicyException{},
		Revoked:        []PolicyException{},
		Expired:        []PolicyException{},
		RecentUsed:     []StoredEvent{},
		RecentInactive: []PolicyException{},
		StatusCounts:   map[string]int64{},
	}

	for _, exception := range s.exceptions {
		if !matchesExceptionScope(exception, filter) {
			continue
		}
		exception = exception.WithEffectiveStatus(now)
		report.StatusCounts[exception.Status]++

		switch exception.Status {
		case ExceptionStatusApproved:
			report.Active = append(report.Active, exception)
		case ExceptionStatusPending:
			report.Pending = append(report.Pending, exception)
		case ExceptionStatusRejected:
			report.Rejected = append(report.Rejected, exception)
			report.RecentInactive = append(report.RecentInactive, exception)
		case ExceptionStatusRevoked:
			report.Revoked = append(report.Revoked, exception)
			report.RecentInactive = append(report.RecentInactive, exception)
		case ExceptionStatusExpired:
			report.Expired = append(report.Expired, exception)
			report.RecentInactive = append(report.RecentInactive, exception)
		}
	}

	sortExceptions := func(exceptions []PolicyException) {
		sort.Slice(exceptions, func(i, j int) bool {
			return exceptions[i].CreatedAt.After(exceptions[j].CreatedAt)
		})
	}
	sortExceptions(report.Active)
	sortExceptions(report.Pending)
	sortExceptions(report.Rejected)
	sortExceptions(report.Revoked)
	sortExceptions(report.Expired)
	sortExceptions(report.RecentInactive)

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

	trimExceptions := func(exceptions []PolicyException) []PolicyException {
		if len(exceptions) > filter.Limit {
			return exceptions[:filter.Limit]
		}
		return exceptions
	}
	report.Active = trimExceptions(report.Active)
	report.Pending = trimExceptions(report.Pending)
	report.Rejected = trimExceptions(report.Rejected)
	report.Revoked = trimExceptions(report.Revoked)
	report.Expired = trimExceptions(report.Expired)
	report.RecentInactive = trimExceptions(report.RecentInactive)
	if len(report.RecentUsed) > filter.Limit {
		report.RecentUsed = report.RecentUsed[:filter.Limit]
	}

	return report, nil
}

func (s *MemoryStore) ListApprovalLogs(_ context.Context, exceptionID string, limit int) ([]ApprovalLog, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	exceptionID = strings.TrimSpace(exceptionID)
	if limit <= 0 {
		limit = 50
	}
	if limit > 500 {
		limit = 500
	}

	logs := make([]ApprovalLog, 0, len(s.approvalLogs))
	for _, log := range s.approvalLogs {
		if exceptionID != "" && log.ExceptionID != exceptionID {
			continue
		}
		logs = append(logs, cloneApprovalLog(log))
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].CreatedAt.After(logs[j].CreatedAt)
	})
	if len(logs) > limit {
		logs = logs[:limit]
	}
	return logs, nil
}

func (s *MemoryStore) Trends(_ context.Context, filter TrendsFilter) (TrendsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeTrendsFilter(filter)
	if err != nil {
		return TrendsResponse{}, err
	}

	windowStart := s.now().UTC().AddDate(0, 0, -filter.WindowDays)
	bucketMap := map[time.Time]*TrendBucket{}
	totals := map[string]int64{
		"allow": 0,
		"deny":  0,
		"error": 0,
	}

	for _, record := range s.records {
		if record.ReceivedAt.Before(windowStart) || !matchesAnalyticsEvent(record.Event, filter.ClusterID, filter.TenantID, filter.Environment, filter.Repo, filter.EventType) {
			continue
		}

		timestamp := truncateTrendTime(record.ReceivedAt.UTC(), filter.Granularity)
		bucket := bucketMap[timestamp]
		if bucket == nil {
			bucket = &TrendBucket{Timestamp: timestamp}
			bucketMap[timestamp] = bucket
		}

		switch record.Decision {
		case DecisionAllow:
			bucket.AllowCount++
			totals["allow"]++
		case DecisionDeny:
			bucket.DenyCount++
			totals["deny"]++
		case DecisionError:
			bucket.ErrorCount++
			totals["error"]++
		}
	}

	buckets := make([]TrendBucket, 0, len(bucketMap))
	for _, bucket := range bucketMap {
		buckets = append(buckets, *bucket)
	}
	sort.Slice(buckets, func(i, j int) bool { return buckets[i].Timestamp.Before(buckets[j].Timestamp) })

	return TrendsResponse{
		Buckets: buckets,
		Totals:  totals,
		AppliedFilters: map[string]string{
			"window_days": fmt.Sprint(filter.WindowDays),
			"granularity": filter.Granularity,
			"cluster_id":  filter.ClusterID,
			"tenant_id":   filter.TenantID,
			"environment": filter.Environment,
			"repo":        filter.Repo,
			"event_type":  filter.EventType,
		},
	}, nil
}

func (s *MemoryStore) TopViolators(_ context.Context, filter TopViolatorsFilter) (TopViolatorsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeTopViolatorsFilter(filter)
	if err != nil {
		return TopViolatorsResponse{}, err
	}

	windowStart := s.now().UTC().AddDate(0, 0, -filter.WindowDays)
	type aggregate struct {
		denyCount   int64
		reasonCount map[string]int64
	}
	aggregates := map[string]*aggregate{}

	for _, record := range s.records {
		if record.ReceivedAt.Before(windowStart) || record.Decision != DecisionDeny {
			continue
		}
		if !matchesAnalyticsEvent(record.Event, filter.ClusterID, filter.TenantID, filter.Environment, filter.Repo, "") {
			continue
		}

		key := violatorKey(record.Event, filter.Dimension)
		item := aggregates[key]
		if item == nil {
			item = &aggregate{reasonCount: map[string]int64{}}
			aggregates[key] = item
		}
		item.denyCount++
		for _, reason := range record.Reasons {
			item.reasonCount[reason]++
		}
	}

	items := make([]TopViolator, 0, len(aggregates))
	for key, aggregate := range aggregates {
		reasons := make([]ReasonCount, 0, len(aggregate.reasonCount))
		for reason, count := range aggregate.reasonCount {
			reasons = append(reasons, ReasonCount{Reason: reason, Count: count})
		}
		sort.Slice(reasons, func(i, j int) bool {
			if reasons[i].Count == reasons[j].Count {
				return reasons[i].Reason < reasons[j].Reason
			}
			return reasons[i].Count > reasons[j].Count
		})
		if len(reasons) > 3 {
			reasons = reasons[:3]
		}
		items = append(items, TopViolator{
			Key:        key,
			DenyCount:  aggregate.denyCount,
			TopReasons: reasons,
		})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].DenyCount == items[j].DenyCount {
			return items[i].Key < items[j].Key
		}
		return items[i].DenyCount > items[j].DenyCount
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}

	return TopViolatorsResponse{
		Items: items,
		AppliedFilters: map[string]string{
			"window_days": filterWindowString(filter.WindowDays),
			"dimension":   filter.Dimension,
			"cluster_id":  filter.ClusterID,
			"tenant_id":   filter.TenantID,
			"environment": filter.Environment,
			"repo":        filter.Repo,
		},
	}, nil
}

func (s *MemoryStore) DriftStats(_ context.Context, filter DriftStatsFilter) (DriftStatsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeDriftStatsFilter(filter)
	if err != nil {
		return DriftStatsResponse{}, err
	}

	windowStart := s.now().UTC().AddDate(0, 0, -filter.WindowDays)
	countsByClass := map[string]int64{}
	workloadCounts := map[string]*DriftWorkloadCount{}
	scopeRecords := []driftScopeRecord{}
	totalRuntimeDriftDenies := int64(0)

	for _, record := range s.records {
		if record.EventType != EventTypeRuntimeDriftResult || record.ReceivedAt.Before(windowStart) {
			continue
		}
		if !matchesAnalyticsEvent(record.Event, filter.ClusterID, filter.TenantID, filter.Environment, filter.Repo, EventTypeRuntimeDriftResult) {
			continue
		}
		if filter.Namespace != "" && record.Namespace != filter.Namespace {
			continue
		}
		if filter.Workload != "" && record.Workload != filter.Workload {
			continue
		}

		scopeRecords = append(scopeRecords, driftScopeRecord{scopeKey: driftScopeKey(record.Event), record: record})
		if record.Decision != DecisionDeny {
			continue
		}

		totalRuntimeDriftDenies++
		driftClasses := record.DriftClasses
		if len(driftClasses) == 0 && strings.TrimSpace(record.DriftResult) != "" {
			driftClasses = []string{record.DriftResult}
		}
		for _, class := range driftClasses {
			countsByClass[class]++
		}

		key := driftScopeKey(record.Event)
		workload := workloadCounts[key]
		if workload == nil {
			workload = &DriftWorkloadCount{
				Workload:    record.Workload,
				Namespace:   record.Namespace,
				TenantID:    record.TenantID,
				Environment: record.Environment,
			}
			workloadCounts[key] = workload
		}
		workload.Count++
	}

	workloads := make([]DriftWorkloadCount, 0, len(workloadCounts))
	for _, workload := range workloadCounts {
		workloads = append(workloads, *workload)
	}
	sort.Slice(workloads, func(i, j int) bool {
		if workloads[i].Count == workloads[j].Count {
			return workloads[i].Workload < workloads[j].Workload
		}
		return workloads[i].Count > workloads[j].Count
	})
	if len(workloads) > 5 {
		workloads = workloads[:5]
	}

	mttr := computeApproximateMTTR(scopeRecords)
	return DriftStatsResponse{
		TotalRuntimeDriftDenies:  totalRuntimeDriftDenies,
		CountsByDriftClass:       countsByClass,
		TopDriftedWorkloads:      workloads,
		MeanTimeToResolveSeconds: mttr,
		AppliedFilters: map[string]string{
			"window_days": filterWindowString(filter.WindowDays),
			"cluster_id":  filter.ClusterID,
			"tenant_id":   filter.TenantID,
			"environment": filter.Environment,
			"repo":        filter.Repo,
			"namespace":   filter.Namespace,
			"workload":    filter.Workload,
		},
	}, nil
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
	if filter.ClusterID != "" && event.ClusterID != filter.ClusterID {
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
	record.Reasons = slicesCloneStrings(record.Reasons)
	record.DriftClasses = slicesCloneStrings(record.DriftClasses)
	if record.RawEvent != nil {
		record.RawEvent = append(record.RawEvent[:0:0], record.RawEvent...)
	}
	return record
}

func matchesExceptionFilter(exception PolicyException, filter ExceptionFilter, now time.Time) bool {
	if filter.Active != nil {
		currentlyActive := exception.IsCurrentlyActive(now)
		if *filter.Active != currentlyActive {
			return false
		}
	}
	if filter.Status != "" && exception.EffectiveStatus(now) != filter.Status {
		return false
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
	return matchesExceptionScope(PolicyException{
		ExceptionType: event.ExceptionType,
		TenantID:      event.TenantID,
		Environment:   event.Environment,
		Namespace:     event.Namespace,
		Repo:          event.Repo,
		ImageDigest:   event.Digest,
		CVEID:         event.CVEID,
	}, filter)
}

func (s *MemoryStore) appendApprovalLogLocked(exceptionID, action, actor, actorRole, reason string, metadata json.RawMessage) {
	log := NormalizeApprovalLog(ApprovalLog{
		ID:          s.nextApprovalLogID,
		ExceptionID: strings.TrimSpace(exceptionID),
		Action:      action,
		Actor:       actor,
		ActorRole:   actorRole,
		Reason:      reason,
		CreatedAt:   s.now().UTC(),
		Metadata:    metadata,
	})
	s.nextApprovalLogID++
	s.approvalLogs = append(s.approvalLogs, log)
}

func timePointer(value time.Time) *time.Time {
	value = value.UTC()
	return &value
}

func matchesAnalyticsEvent(event Event, clusterID, tenantID, environment, repo, eventType string) bool {
	if clusterID != "" && event.ClusterID != clusterID {
		return false
	}
	if tenantID != "" && event.TenantID != tenantID {
		return false
	}
	if environment != "" && event.Environment != environment {
		return false
	}
	if repo != "" && event.Repo != repo {
		return false
	}
	if eventType != "" && event.EventType != eventType {
		return false
	}
	return true
}

func truncateTrendTime(value time.Time, granularity string) time.Time {
	value = value.UTC()
	if granularity == "hour" {
		return value.Truncate(time.Hour)
	}
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC)
}

func violatorKey(event Event, dimension string) string {
	switch dimension {
	case "tenant":
		return firstAnalyticsKey(event.TenantID)
	case "environment":
		return firstAnalyticsKey(event.Environment)
	default:
		return firstAnalyticsKey(event.Repo)
	}
}

func driftScopeKey(event Event) string {
	return strings.Join([]string{
		firstAnalyticsKey(event.TenantID),
		firstAnalyticsKey(event.Environment),
		firstAnalyticsKey(event.Namespace),
		firstAnalyticsKey(event.Workload),
	}, "|")
}

func computeApproximateMTTR(records []driftScopeRecord) *int64 {
	if len(records) == 0 {
		return nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].record.ReceivedAt.Before(records[j].record.ReceivedAt)
	})

	lastDeny := map[string]time.Time{}
	var total int64
	var resolved int64

	for _, item := range records {
		record := item.record
		if record.Decision == DecisionDeny {
			lastDeny[item.scopeKey] = record.ReceivedAt
			continue
		}
		if record.Decision != DecisionAllow && strings.TrimSpace(record.DriftResult) != "no_drift" {
			continue
		}
		start, ok := lastDeny[item.scopeKey]
		if !ok || !record.ReceivedAt.After(start) {
			continue
		}
		total += int64(record.ReceivedAt.Sub(start).Seconds())
		resolved++
		delete(lastDeny, item.scopeKey)
	}

	if resolved == 0 {
		return nil
	}
	mean := total / resolved
	return &mean
}

func firstAnalyticsKey(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "unknown"
	}
	return value
}

func filterWindowString(days int) string {
	return fmt.Sprint(days)
}

func slicesCloneStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	cloned := make([]string, len(values))
	copy(cloned, values)
	return cloned
}
