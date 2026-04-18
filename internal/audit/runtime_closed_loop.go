package audit

import (
	"sort"
	"strings"
	"time"
)

type RuntimeActiveStateFilter struct {
	ClusterID             string
	TenantID              string
	Namespace             string
	WorkloadKind          string
	Workload              string
	ReconciliationStatus  string
	QuarantineType        string
	Limit                 int
}

type RuntimeActiveStateView struct {
	ID                       string           `json:"id"`
	TenantID                 string           `json:"tenant_id,omitempty"`
	ClusterID                string           `json:"cluster_id,omitempty"`
	Namespace                string           `json:"namespace"`
	WorkloadKind             string           `json:"workload_kind"`
	Workload                 string           `json:"workload"`
	ServiceAccount           string           `json:"service_account,omitempty"`
	ObservedDigest           string           `json:"observed_digest,omitempty"`
	ApprovedDigest           string           `json:"approved_digest,omitempty"`
	ObservedConfigHash       string           `json:"observed_config_hash,omitempty"`
	ExpectedConfigHash       string           `json:"expected_config_hash,omitempty"`
	DriftResult              string           `json:"drift_result,omitempty"`
	DriftClasses             []string         `json:"drift_classes,omitempty"`
	DriftSeverity            string           `json:"drift_severity,omitempty"`
	ReconciliationStatus     string           `json:"reconciliation_status"`
	RemediationMode          string           `json:"remediation_mode,omitempty"`
	RemediationAttempt       int              `json:"remediation_attempt,omitempty"`
	Remediable               bool             `json:"remediable"`
	QuarantineReason         string           `json:"quarantine_reason,omitempty"`
	QuarantineType           string           `json:"quarantine_type,omitempty"`
	ProtectedTarget          bool             `json:"protected_target,omitempty"`
	ProtectedReason          string           `json:"protected_reason,omitempty"`
	DesiredStateSourceRef    string           `json:"desired_state_source_ref,omitempty"`
	DesiredStateApprovalID   string           `json:"desired_state_approval_id,omitempty"`
	DesiredStateVerification string           `json:"desired_state_verification_state,omitempty"`
	LastError                string           `json:"last_error,omitempty"`
	LastReconciledAt         time.Time        `json:"last_reconciled_at"`
	Reasons                  []string         `json:"reasons,omitempty"`
	Evidence                 *RuntimeEvidence `json:"evidence,omitempty"`
}

type RuntimeClosedLoopStatus struct {
	TotalTargets          int64            `json:"total_targets"`
	InSync                int64            `json:"in_sync"`
	DriftDetected         int64            `json:"drift_detected"`
	Remediated            int64            `json:"remediated"`
	Failed                int64            `json:"failed"`
	Quarantined           int64            `json:"quarantined"`
	ProtectedBlocked      int64            `json:"protected_blocked"`
	CountsByStatus        map[string]int64 `json:"counts_by_status"`
	CountsByQuarantine    map[string]int64 `json:"counts_by_quarantine_type"`
	LastReconciledAt      *time.Time       `json:"last_reconciled_at,omitempty"`
}

func DeriveRuntimeActiveStates(events []StoredEvent, filter RuntimeActiveStateFilter) []RuntimeActiveStateView {
	normalizeRuntimeActiveStateFilter(&filter)
	sort.Slice(events, func(i, j int) bool {
		return events[i].ReceivedAt.After(events[j].ReceivedAt)
	})

	latest := map[string]RuntimeActiveStateView{}
	for _, record := range events {
		if record.EventType != EventTypeRuntimeActiveStateObserved {
			continue
		}
		item := RuntimeActiveStateView{
			ID:                       runtimeTargetID(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload),
			TenantID:                 record.TenantID,
			ClusterID:                record.ClusterID,
			Namespace:                record.Namespace,
			WorkloadKind:             normalizeRuntimeKind(record.WorkloadKind),
			Workload:                 record.Workload,
			ServiceAccount:           record.ServiceAccount,
			ObservedDigest:           record.Digest,
			DriftResult:              record.DriftResult,
			DriftClasses:             append([]string(nil), record.DriftClasses...),
			DriftSeverity:            record.DriftSeverity,
			ReconciliationStatus:     record.ReconciliationStatus,
			RemediationMode:          record.RemediationMode,
			RemediationAttempt:       record.RemediationAttempt,
			Remediable:               record.Remediable,
			QuarantineReason:         record.QuarantineReason,
			QuarantineType:           record.QuarantineType,
			ProtectedTarget:          record.ProtectedTarget,
			ProtectedReason:          record.ProtectedReason,
			DesiredStateSourceRef:    record.DesiredStateSourceRef,
			DesiredStateApprovalID:   record.DesiredStateApprovalID,
			DesiredStateVerification: record.DesiredStateVerification,
			LastReconciledAt:         record.Timestamp,
			Reasons:                  append([]string(nil), record.Reasons...),
		}
		if len(item.Reasons) > 0 {
			item.LastError = item.Reasons[0]
		}
		if record.Evidence != nil && record.Evidence.Runtime != nil {
			runtimeEvidence := *record.Evidence.Runtime
			item.Evidence = &runtimeEvidence
			item.ApprovedDigest = runtimeEvidence.ApprovedDigest
			item.ExpectedConfigHash = runtimeEvidence.ExpectedConfigHash
			item.ObservedConfigHash = runtimeEvidence.ActualConfigHash
			if item.ServiceAccount == "" {
				item.ServiceAccount = FirstNonEmpty(runtimeEvidence.ServiceAccountObserved, runtimeEvidence.ServiceAccountExpected)
			}
			if item.ObservedDigest == "" {
				item.ObservedDigest = runtimeEvidence.RunningDigest
			}
		}
		if item.ID == "" {
			continue
		}
		if !matchesRuntimeTarget(item.TenantID, item.ClusterID, item.Namespace, item.WorkloadKind, item.Workload, filter.TenantID, filter.ClusterID, filter.Namespace, filter.WorkloadKind, filter.Workload) {
			continue
		}
		if filter.ReconciliationStatus != "" && item.ReconciliationStatus != filter.ReconciliationStatus {
			continue
		}
		if filter.QuarantineType != "" && item.QuarantineType != filter.QuarantineType {
			continue
		}
		if _, ok := latest[item.ID]; ok {
			continue
		}
		latest[item.ID] = item
	}

	items := make([]RuntimeActiveStateView, 0, len(latest))
	for _, item := range latest {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].LastReconciledAt.After(items[j].LastReconciledAt) })
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items
}

func DeriveRuntimeClosedLoopStatus(items []RuntimeActiveStateView) RuntimeClosedLoopStatus {
	status := RuntimeClosedLoopStatus{
		CountsByStatus:     map[string]int64{},
		CountsByQuarantine: map[string]int64{},
	}
	for _, item := range items {
		status.TotalTargets++
		status.CountsByStatus[item.ReconciliationStatus]++
		switch item.ReconciliationStatus {
		case "in_sync":
			status.InSync++
		case "drift_detected":
			status.DriftDetected++
		case "remediated":
			status.Remediated++
		case "failed":
			status.Failed++
		case "quarantined":
			status.Quarantined++
		}
		if item.ProtectedTarget {
			status.ProtectedBlocked++
		}
		if item.QuarantineType != "" {
			status.CountsByQuarantine[item.QuarantineType]++
		}
		if !item.LastReconciledAt.IsZero() && (status.LastReconciledAt == nil || item.LastReconciledAt.After(*status.LastReconciledAt)) {
			copy := item.LastReconciledAt
			status.LastReconciledAt = &copy
		}
	}
	return status
}

func normalizeRuntimeActiveStateFilter(filter *RuntimeActiveStateFilter) {
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Namespace = strings.TrimSpace(filter.Namespace)
	filter.WorkloadKind = normalizeRuntimeKind(filter.WorkloadKind)
	filter.Workload = strings.TrimSpace(filter.Workload)
	filter.ReconciliationStatus = strings.TrimSpace(filter.ReconciliationStatus)
	filter.QuarantineType = strings.TrimSpace(filter.QuarantineType)
	if filter.Limit <= 0 {
		filter.Limit = 100
	}
}
