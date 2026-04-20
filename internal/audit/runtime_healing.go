package audit

import (
	"sort"
	"strings"
	"time"
)

type RuntimeDesiredStateFilter struct {
	ClusterID    string
	TenantID     string
	Namespace    string
	WorkloadKind string
	Workload     string
	Limit        int
}

type RuntimeDesiredStateView struct {
	ID                       string                     `json:"id"`
	TenantID                 string                     `json:"tenant_id,omitempty"`
	ClusterID                string                     `json:"cluster_id,omitempty"`
	Namespace                string                     `json:"namespace"`
	WorkloadKind             string                     `json:"workload_kind"`
	Workload                 string                     `json:"workload"`
	ServiceAccount           string                     `json:"service_account,omitempty"`
	Labels                   map[string]string          `json:"labels,omitempty"`
	Containers               []RuntimeApprovedContainer `json:"containers,omitempty"`
	ApprovedDigest           string                     `json:"approved_digest,omitempty"`
	ExpectedConfigHash       string                     `json:"expected_config_hash,omitempty"`
	DesiredStateSourceRef    string                     `json:"desired_state_source_ref,omitempty"`
	DesiredStateApprovalID   string                     `json:"desired_state_approval_id,omitempty"`
	DesiredStateVerification string                     `json:"desired_state_verification_state,omitempty"`
	LastApprovedAt           time.Time                  `json:"last_approved_at"`
}

type RuntimeDriftFilter struct {
	ClusterID    string
	TenantID     string
	Namespace    string
	WorkloadKind string
	Workload     string
	Severity     string
	Status       string
	Limit        int
}

type RuntimeDriftFinding struct {
	ID                       string           `json:"id"`
	TenantID                 string           `json:"tenant_id,omitempty"`
	ClusterID                string           `json:"cluster_id,omitempty"`
	Namespace                string           `json:"namespace"`
	WorkloadKind             string           `json:"workload_kind"`
	Workload                 string           `json:"workload"`
	ServiceAccount           string           `json:"service_account,omitempty"`
	DriftResult              string           `json:"drift_result"`
	DriftClasses             []string         `json:"drift_classes,omitempty"`
	DriftSeverity            string           `json:"drift_severity,omitempty"`
	RemediationMode          string           `json:"remediation_mode,omitempty"`
	RemediationAttempt       int              `json:"remediation_attempt,omitempty"`
	Remediable               bool             `json:"remediable"`
	Status                   string           `json:"status"`
	QuarantineReason         string           `json:"quarantine_reason,omitempty"`
	DesiredStateVerification string           `json:"desired_state_verification_state,omitempty"`
	DetectedAt               time.Time        `json:"detected_at"`
	LastUpdatedAt            time.Time        `json:"last_updated_at"`
	LastEventType            string           `json:"last_event_type"`
	Reasons                  []string         `json:"reasons,omitempty"`
	Evidence                 *RuntimeEvidence `json:"evidence,omitempty"`
}

type RuntimeDriftStatus struct {
	TotalFindings    int64            `json:"total_findings"`
	ActiveFindings   int64            `json:"active_findings"`
	Quarantined      int64            `json:"quarantined"`
	Failed           int64            `json:"failed"`
	Remediated       int64            `json:"remediated"`
	Detected         int64            `json:"detected"`
	CountsBySeverity map[string]int64 `json:"counts_by_severity"`
	CountsByStatus   map[string]int64 `json:"counts_by_status"`
	LastDetectedAt   *time.Time       `json:"last_detected_at,omitempty"`
	LastUpdatedAt    *time.Time       `json:"last_updated_at,omitempty"`
}

func DeriveRuntimeDesiredStates(events []StoredEvent, filter RuntimeDesiredStateFilter) []RuntimeDesiredStateView {
	normalizeRuntimeDesiredStateFilter(&filter)
	latest := map[string]RuntimeDesiredStateView{}

	sort.Slice(events, func(i, j int) bool {
		return events[i].ReceivedAt.After(events[j].ReceivedAt)
	})
	for _, record := range events {
		if record.EventType != EventTypeRuntimeDesiredStateRecorded {
			continue
		}
		view := RuntimeDesiredStateView{
			ID:                       runtimeTargetID(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload),
			TenantID:                 record.TenantID,
			ClusterID:                record.ClusterID,
			Namespace:                record.Namespace,
			WorkloadKind:             normalizeRuntimeKind(record.WorkloadKind),
			Workload:                 record.Workload,
			ServiceAccount:           record.ServiceAccount,
			DesiredStateSourceRef:    record.DesiredStateSourceRef,
			DesiredStateApprovalID:   record.DesiredStateApprovalID,
			DesiredStateVerification: record.DesiredStateVerification,
			LastApprovedAt:           record.Timestamp,
		}
		if record.Evidence != nil && record.Evidence.Runtime != nil {
			view.ApprovedDigest = record.Evidence.Runtime.ApprovedDigest
			view.ExpectedConfigHash = record.Evidence.Runtime.ExpectedConfigHash
			view.Labels = cloneStringMap(record.Evidence.Runtime.ApprovedLabels)
			view.Containers = append([]RuntimeApprovedContainer(nil), record.Evidence.Runtime.ApprovedContainers...)
			if view.ServiceAccount == "" {
				view.ServiceAccount = record.Evidence.Runtime.ServiceAccountExpected
			}
		}
		if !matchesRuntimeTarget(view.TenantID, view.ClusterID, view.Namespace, view.WorkloadKind, view.Workload, filter.TenantID, filter.ClusterID, filter.Namespace, filter.WorkloadKind, filter.Workload) {
			continue
		}
		if _, exists := latest[view.ID]; exists {
			continue
		}
		latest[view.ID] = view
	}

	items := make([]RuntimeDesiredStateView, 0, len(latest))
	for _, item := range latest {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].LastApprovedAt.After(items[j].LastApprovedAt) })
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items
}

func DeriveRuntimeDriftFindings(events []StoredEvent, filter RuntimeDriftFilter) []RuntimeDriftFinding {
	normalizeRuntimeDriftFilter(&filter)
	sort.Slice(events, func(i, j int) bool {
		return events[i].ReceivedAt.Before(events[j].ReceivedAt)
	})

	findings := map[string]RuntimeDriftFinding{}
	for _, record := range events {
		if !isRuntimeLifecycleEvent(record.EventType) {
			continue
		}
		id := runtimeTargetID(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload)
		if id == "" {
			continue
		}
		current, exists := findings[id]
		if !exists {
			current = RuntimeDriftFinding{
				ID:           id,
				TenantID:     record.TenantID,
				ClusterID:    record.ClusterID,
				Namespace:    record.Namespace,
				WorkloadKind: normalizeRuntimeKind(record.WorkloadKind),
				Workload:     record.Workload,
				DetectedAt:   record.Timestamp,
			}
		}
		if !record.Timestamp.IsZero() {
			current.LastUpdatedAt = record.Timestamp
		}
		if current.DetectedAt.IsZero() || (!record.Timestamp.IsZero() && record.Timestamp.Before(current.DetectedAt)) {
			current.DetectedAt = record.Timestamp
		}
		current.ServiceAccount = FirstNonEmpty(record.ServiceAccount, current.ServiceAccount)
		current.DriftResult = FirstNonEmpty(record.DriftResult, current.DriftResult)
		if len(record.DriftClasses) > 0 {
			current.DriftClasses = append([]string(nil), record.DriftClasses...)
		}
		current.DriftSeverity = FirstNonEmpty(record.DriftSeverity, current.DriftSeverity)
		current.RemediationMode = FirstNonEmpty(record.RemediationMode, current.RemediationMode)
		if record.RemediationAttempt > current.RemediationAttempt {
			current.RemediationAttempt = record.RemediationAttempt
		}
		current.Remediable = current.Remediable || record.Remediable
		current.DesiredStateVerification = FirstNonEmpty(record.DesiredStateVerification, current.DesiredStateVerification)
		if len(record.Reasons) > 0 {
			current.Reasons = append([]string(nil), record.Reasons...)
		}
		if record.Evidence != nil && record.Evidence.Runtime != nil {
			runtimeEvidence := *record.Evidence.Runtime
			current.Evidence = &runtimeEvidence
		}
		current.LastEventType = record.EventType
		switch record.EventType {
		case EventTypeDriftDetected, EventTypeRuntimeDriftResult:
			current.Status = "detected"
		case EventTypeDriftRemediationSucceeded:
			current.Status = "remediated"
		case EventTypeDriftRemediationFailed:
			current.Status = "failed"
		case EventTypeDriftQuarantined:
			current.Status = "quarantined"
			current.QuarantineReason = FirstNonEmpty(record.QuarantineReason, current.QuarantineReason)
		case EventTypeDriftRemediationStarted:
			if current.Status == "" {
				current.Status = "detected"
			}
		}
		findings[id] = current
	}

	items := make([]RuntimeDriftFinding, 0, len(findings))
	for _, item := range findings {
		if !matchesRuntimeTarget(item.TenantID, item.ClusterID, item.Namespace, item.WorkloadKind, item.Workload, filter.TenantID, filter.ClusterID, filter.Namespace, filter.WorkloadKind, filter.Workload) {
			continue
		}
		if filter.Severity != "" && item.DriftSeverity != filter.Severity {
			continue
		}
		if filter.Status != "" && item.Status != filter.Status {
			continue
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].LastUpdatedAt.After(items[j].LastUpdatedAt) })
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items
}

func DeriveRuntimeDriftStatus(findings []RuntimeDriftFinding) RuntimeDriftStatus {
	status := RuntimeDriftStatus{
		CountsBySeverity: map[string]int64{},
		CountsByStatus:   map[string]int64{},
	}
	for _, finding := range findings {
		status.TotalFindings++
		status.CountsByStatus[finding.Status]++
		if finding.DriftSeverity != "" {
			status.CountsBySeverity[finding.DriftSeverity]++
		}
		switch finding.Status {
		case "detected":
			status.Detected++
			status.ActiveFindings++
		case "failed":
			status.Failed++
			status.ActiveFindings++
		case "quarantined":
			status.Quarantined++
			status.ActiveFindings++
		case "remediated":
			status.Remediated++
		}
		if !finding.DetectedAt.IsZero() && (status.LastDetectedAt == nil || finding.DetectedAt.After(*status.LastDetectedAt)) {
			copy := finding.DetectedAt
			status.LastDetectedAt = &copy
		}
		if !finding.LastUpdatedAt.IsZero() && (status.LastUpdatedAt == nil || finding.LastUpdatedAt.After(*status.LastUpdatedAt)) {
			copy := finding.LastUpdatedAt
			status.LastUpdatedAt = &copy
		}
	}
	return status
}

func normalizeRuntimeDesiredStateFilter(filter *RuntimeDesiredStateFilter) {
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Namespace = strings.TrimSpace(filter.Namespace)
	filter.WorkloadKind = normalizeRuntimeKind(filter.WorkloadKind)
	filter.Workload = strings.TrimSpace(filter.Workload)
	if filter.Limit <= 0 {
		filter.Limit = 50
	}
}

func cloneStringMap(values map[string]string) map[string]string {
	if len(values) == 0 {
		return nil
	}
	cloned := make(map[string]string, len(values))
	for key, value := range values {
		cloned[key] = value
	}
	return cloned
}

func normalizeRuntimeDriftFilter(filter *RuntimeDriftFilter) {
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Namespace = strings.TrimSpace(filter.Namespace)
	filter.WorkloadKind = normalizeRuntimeKind(filter.WorkloadKind)
	filter.Workload = strings.TrimSpace(filter.Workload)
	filter.Severity = strings.TrimSpace(filter.Severity)
	filter.Status = strings.TrimSpace(filter.Status)
	if filter.Limit <= 0 {
		filter.Limit = 100
	}
}

func isRuntimeLifecycleEvent(eventType string) bool {
	switch eventType {
	case EventTypeRuntimeDriftResult, EventTypeDriftDetected, EventTypeDriftRemediationStarted, EventTypeDriftRemediationSucceeded, EventTypeDriftRemediationFailed, EventTypeDriftQuarantined:
		return true
	default:
		return false
	}
}

func runtimeTargetID(clusterID, namespace, kind, workload string) string {
	namespace = strings.TrimSpace(namespace)
	workload = strings.TrimSpace(workload)
	if namespace == "" || workload == "" {
		return ""
	}
	return FirstNonEmpty(strings.TrimSpace(clusterID), "local") + "|" + namespace + "|" + normalizeRuntimeKind(kind) + "|" + workload
}

func normalizeRuntimeKind(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "deployment":
		return "Deployment"
	case "daemonset":
		return "DaemonSet"
	case "statefulset":
		return "StatefulSet"
	default:
		return strings.TrimSpace(value)
	}
}

func matchesRuntimeTarget(eventTenant, eventCluster, eventNamespace, eventKind, eventWorkload, filterTenant, filterCluster, filterNamespace, filterKind, filterWorkload string) bool {
	if filterTenant != "" && eventTenant != filterTenant {
		return false
	}
	if filterCluster != "" && eventCluster != filterCluster {
		return false
	}
	if filterNamespace != "" && eventNamespace != filterNamespace {
		return false
	}
	if filterKind != "" && normalizeRuntimeKind(eventKind) != filterKind {
		return false
	}
	if filterWorkload != "" && eventWorkload != filterWorkload {
		return false
	}
	return true
}
