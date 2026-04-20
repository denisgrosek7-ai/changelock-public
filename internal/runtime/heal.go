package runtime

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type FindingStatus string

const (
	FindingStatusDetected    FindingStatus = "detected"
	FindingStatusRemediated  FindingStatus = "remediated"
	FindingStatusFailed      FindingStatus = "failed"
	FindingStatusQuarantined FindingStatus = "quarantined"
)

type ReconciliationStatus string

const (
	ReconciliationStatusInSync      ReconciliationStatus = "in_sync"
	ReconciliationStatusDrift       ReconciliationStatus = "drift_detected"
	ReconciliationStatusRemediating ReconciliationStatus = "remediating"
	ReconciliationStatusRemediated  ReconciliationStatus = "remediated"
	ReconciliationStatusFailed      ReconciliationStatus = "failed"
	ReconciliationStatusQuarantined ReconciliationStatus = "quarantined"
)

type RemediationOutcome struct {
	Status           FindingStatus   `json:"status"`
	Mode             RemediationMode `json:"mode"`
	AttemptCount     int             `json:"attempt_count"`
	Quarantined      bool            `json:"quarantined"`
	QuarantineReason string          `json:"quarantine_reason,omitempty"`
	Message          string          `json:"message,omitempty"`
}

type trackerState struct {
	attempts         []time.Time
	quarantined      bool
	quarantineReason string
}

type Tracker struct {
	mu     sync.Mutex
	states map[string]*trackerState
	now    func() time.Time
}

func NewTracker() *Tracker {
	return &Tracker{
		states: map[string]*trackerState{},
		now:    time.Now,
	}
}

func (t *Tracker) Current(key string, window time.Duration) (attemptCount int, quarantined bool, quarantineReason string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	state := t.states[key]
	if state == nil {
		return 0, false, ""
	}
	state.attempts = filterRecentAttempts(state.attempts, t.now(), window)
	return len(state.attempts), state.quarantined, state.quarantineReason
}

func (t *Tracker) RecordAttempt(key string, window time.Duration) int {
	t.mu.Lock()
	defer t.mu.Unlock()

	state := t.ensure(key)
	state.attempts = append(filterRecentAttempts(state.attempts, t.now(), window), t.now().UTC())
	return len(state.attempts)
}

func (t *Tracker) Quarantine(key, reason string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	state := t.ensure(key)
	state.quarantined = true
	state.quarantineReason = strings.TrimSpace(reason)
}

func (t *Tracker) ClearQuarantine(key string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if state := t.states[key]; state != nil {
		state.quarantined = false
		state.quarantineReason = ""
	}
}

func (t *Tracker) ensure(key string) *trackerState {
	state := t.states[key]
	if state == nil {
		state = &trackerState{}
		t.states[key] = state
	}
	return state
}

func filterRecentAttempts(values []time.Time, now time.Time, window time.Duration) []time.Time {
	if len(values) == 0 || window <= 0 {
		return values
	}
	threshold := now.Add(-window)
	filtered := values[:0]
	for _, value := range values {
		if value.After(threshold) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func SelectRemediationMode(config SelfHealingConfig, desired ApprovedWorkloadState, result ComparisonResult) (RemediationMode, string) {
	if config.Mode == RemediationModeDisabled {
		return RemediationModeDisabled, "self-healing disabled"
	}
	if config.Mode == RemediationModeAlertOnly {
		return RemediationModeAlertOnly, "alert-only mode configured"
	}
	if config.CriticalOnly && result.Severity != DriftSeverityCritical {
		return RemediationModeAlertOnly, "critical-only remediation is enabled"
	}
	if _, ok := config.AllowedKinds[normalizeWorkloadKind(desired.WorkloadKind)]; !ok {
		return RemediationModeAlertOnly, "workload kind is not allowed for remediation"
	}
	if (config.RequireSignedDesiredState || config.VerifyDesiredStateOnLoop) && desired.DesiredStateVerificationState != VerificationStateVerified {
		if config.FailMode == RemediationModeAlertOnly {
			return RemediationModeAlertOnly, "desired state verification is required before remediation"
		}
		return RemediationModeQuarantine, "desired state verification is required before remediation"
	}
	if !result.Remediable {
		return RemediationModeAlertOnly, "drift is not remediable with the configured safe strategies"
	}
	return config.Mode, ""
}

func RemediationKey(clusterID string, desired ApprovedWorkloadState) string {
	return fmt.Sprintf("%s/%s/%s/%s", strings.TrimSpace(clusterID), strings.TrimSpace(desired.Namespace), normalizeWorkloadKind(desired.WorkloadKind), strings.TrimSpace(desired.Workload))
}

func normalizeWorkloadKind(value string) string {
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

func NormalizeWorkloadKind(value string) string {
	return normalizeWorkloadKind(value)
}

func ReconciliationStatusFromOutcome(hasDrift bool, outcome *RemediationOutcome) ReconciliationStatus {
	if outcome == nil {
		if !hasDrift {
			return ReconciliationStatusInSync
		}
		return ReconciliationStatusDrift
	}
	switch outcome.Status {
	case FindingStatusRemediated:
		return ReconciliationStatusRemediated
	case FindingStatusFailed:
		return ReconciliationStatusFailed
	case FindingStatusQuarantined:
		return ReconciliationStatusQuarantined
	default:
		if !hasDrift {
			return ReconciliationStatusInSync
		}
		return ReconciliationStatusDrift
	}
}

func ProtectedTarget(config SelfHealingConfig, desired ApprovedWorkloadState) (bool, string) {
	namespace := strings.TrimSpace(desired.Namespace)
	if namespace != "" {
		if _, ok := config.ProtectedNamespaces[namespace]; ok {
			return true, "protected namespace"
		}
	}
	workloadKey := strings.TrimSpace(desired.Namespace) + "/" + normalizeWorkloadKind(desired.WorkloadKind) + "/" + strings.TrimSpace(desired.Workload)
	if _, ok := config.ProtectedWorkloads[workloadKey]; ok {
		return true, "protected workload"
	}
	return false, ""
}
