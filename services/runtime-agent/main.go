package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
	"github.com/denisgrosek/changelock/internal/signingidentity"
)

var auditWriter = audit.NewDefaultWriter()
var runtimeControl = newAgentRuntime()

type agentRuntime struct {
	config         runtimestate.SelfHealingConfig
	identityConfig signingidentity.Config
	stateReader    runtimestate.StateReader
	remediator     runtimestate.RemediationClient
	tracker        *runtimestate.Tracker
	controlPlane   runtimeControlPlane
}

type scanRequest struct {
	ScanID   string                              `json:"scan_id,omitempty"`
	Approved runtimestate.ApprovedWorkloadState  `json:"approved"`
	Observed *runtimestate.ObservedWorkloadState `json:"observed,omitempty"`
}

type scanResponse struct {
	ScanID      string                           `json:"scan_id"`
	Result      runtimestate.ComparisonResult    `json:"result"`
	Remediation *runtimestate.RemediationOutcome `json:"remediation,omitempty"`
}

func main() {
	addr := ":" + envOrDefault("PORT", "8093")
	if needsKubernetesClient(runtimeControl.config.Mode) {
		if _, ok := runtimeControl.remediator.(*runtimestate.KubernetesClient); !ok {
			log.Fatal("runtime-agent self-healing mutation modes require Kubernetes API access")
		}
	}
	runtimeControl.startClosedLoop(context.Background())
	log.Printf("runtime-agent listening on %s", addr)
	log.Fatal((&http.Server{
		Addr:              addr,
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}).ListenAndServe())
}

func newAgentRuntime() agentRuntime {
	config, err := runtimestate.LoadSelfHealingConfig()
	if err != nil {
		panic(err)
	}
	identityConfig, err := signingidentity.ParseConfig(os.Getenv)
	if err != nil {
		panic(err)
	}

	reader, remediator := newRuntimeClients()
	return agentRuntime{
		config:         config,
		identityConfig: identityConfig,
		stateReader:    reader,
		remediator:     remediator,
		tracker:        runtimestate.NewTracker(),
		controlPlane:   newControlPlaneClient(),
	}
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/scan", scanHandler)
	return metrics.InstrumentHTTP("runtime-agent", mux)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	httpjson.Write(w, http.StatusOK, map[string]string{"status": "ok"})
}

func scanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request scanRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if request.Approved.Namespace == "" || request.Approved.Workload == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "approved namespace and workload are required"})
		return
	}
	request.Approved.WorkloadKind = firstNonEmpty(runtimestate.NormalizeWorkloadKind(request.Approved.WorkloadKind), "Deployment")
	if request.Approved.DesiredStateVerificationState == "" {
		request.Approved.DesiredStateVerificationState = runtimestate.VerificationStateDisabled
	}

	scanID := firstNonEmpty(request.ScanID, r.Header.Get("X-Request-Id"), audit.NewRequestID())
	writeAuditEvent(r.Context(), buildDesiredStateAuditEvent(scanID, request.Approved))

	observed, err := resolveObservedState(r.Context(), request)
	if err != nil {
		event := audit.Event{
			RequestID:    scanID,
			Component:    "runtime-agent",
			EventType:    audit.EventTypeRuntimeDriftResult,
			TenantID:     firstNonEmpty(request.Approved.TenantID, audit.TenantFromNamespace(request.Approved.Namespace)),
			ClusterID:    request.Approved.ClusterID,
			Environment:  audit.EnvironmentFromNamespace(request.Approved.Namespace),
			Namespace:    request.Approved.Namespace,
			WorkloadKind: request.Approved.WorkloadKind,
			Workload:     request.Approved.Workload,
			Decision:     audit.DecisionError,
			Reasons:      []string{"runtime state unavailable: " + err.Error()},
		}
		writeAuditEvent(r.Context(), event)
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	result := runtimestate.Compare(request.Approved, observed)
	result.ScanID = scanID
	mode, reason := runtimestate.SelectRemediationMode(runtimeControl.config, request.Approved, result)
	result.SelectedRemediationMode = mode

	var remediation *runtimestate.RemediationOutcome
	if result.HasDrift() {
		writeAuditEvent(r.Context(), buildDriftDetectedEvent(scanID, request.Approved, result))
		outcome := applyRemediation(r.Context(), scanID, request.Approved, observed, result, mode, reason)
		remediation = &outcome
	} else {
		runtimeControl.tracker.ClearQuarantine(runtimestate.RemediationKey(result.ClusterID, request.Approved))
	}

	writeAuditEvent(r.Context(), buildAuditEvent(scanID, request.Approved, result))
	writeAuditEvent(r.Context(), buildActiveStateEvent(scanID, request.Approved, observed, result, remediation, false, "", ""))
	httpjson.Write(w, http.StatusOK, scanResponse{
		ScanID:      scanID,
		Result:      result,
		Remediation: remediation,
	})
}

func resolveObservedState(ctx context.Context, request scanRequest) (runtimestate.ObservedWorkloadState, error) {
	if request.Observed != nil {
		return *request.Observed, nil
	}
	return runtimeControl.stateReader.ReadObservedWorkload(ctx, runtimestate.WorkloadTarget{
		ClusterID: request.Approved.ClusterID,
		Namespace: request.Approved.Namespace,
		Kind:      request.Approved.WorkloadKind,
		Workload:  request.Approved.Workload,
	})
}

func applyRemediation(ctx context.Context, scanID string, desired runtimestate.ApprovedWorkloadState, observed runtimestate.ObservedWorkloadState, result runtimestate.ComparisonResult, mode runtimestate.RemediationMode, reason string) runtimestate.RemediationOutcome {
	key := runtimestate.RemediationKey(result.ClusterID, desired)
	attemptCount, quarantined, quarantineReason := runtimeControl.tracker.Current(key, runtimeControl.config.Window)
	if quarantined {
		outcome := runtimestate.RemediationOutcome{
			Status:           runtimestate.FindingStatusQuarantined,
			Mode:             mode,
			AttemptCount:     attemptCount,
			Quarantined:      true,
			QuarantineReason: quarantineReason,
			Message:          "workload is already quarantined from repeated drift",
		}
		writeAuditEvent(ctx, buildQuarantineEvent(scanID, desired, result, outcome, "drift"))
		return outcome
	}

	switch mode {
	case runtimestate.RemediationModeDisabled, runtimestate.RemediationModeAlertOnly:
		return runtimestate.RemediationOutcome{
			Status:       runtimestate.FindingStatusDetected,
			Mode:         mode,
			AttemptCount: attemptCount,
			Message:      firstNonEmpty(reason, "drift recorded; no automated remediation configured"),
		}
	case runtimestate.RemediationModeQuarantine:
		runtimeControl.tracker.Quarantine(key, firstNonEmpty(reason, "quarantine mode configured"))
		outcome := runtimestate.RemediationOutcome{
			Status:           runtimestate.FindingStatusQuarantined,
			Mode:             mode,
			AttemptCount:     attemptCount,
			Quarantined:      true,
			QuarantineReason: firstNonEmpty(reason, "quarantine mode configured"),
			Message:          "drift quarantined for operator review",
		}
		writeAuditEvent(ctx, buildQuarantineEvent(scanID, desired, result, outcome, "drift"))
		return outcome
	case runtimestate.RemediationModeRestartApprovedState:
		if !restartSafe(desired, observed) {
			return runtimestate.RemediationOutcome{
				Status:       runtimestate.FindingStatusFailed,
				Mode:         mode,
				AttemptCount: attemptCount,
				Message:      "restart-to-approved-state requires the controller spec to already match the approved desired state",
			}
		}
	}

	attemptCount = runtimeControl.tracker.RecordAttempt(key, runtimeControl.config.Window)
	if attemptCount > runtimeControl.config.MaxAttempts {
		runtimeControl.tracker.Quarantine(key, "drift remediation threshold exceeded")
		outcome := runtimestate.RemediationOutcome{
			Status:           runtimestate.FindingStatusQuarantined,
			Mode:             mode,
			AttemptCount:     attemptCount,
			Quarantined:      true,
			QuarantineReason: "drift remediation threshold exceeded",
			Message:          "auto-remediation stopped after repeated drift attempts",
		}
		writeAuditEvent(ctx, buildQuarantineEvent(scanID, desired, result, outcome, "repeat-drift"))
		return outcome
	}

	writeAuditEvent(ctx, buildRemediationLifecycleEvent(audit.EventTypeDriftRemediationStarted, scanID, desired, result, attemptCount, "starting remediation"))
	var err error
	switch mode {
	case runtimestate.RemediationModePatchApprovedState:
		err = runtimeControl.remediator.PatchApprovedState(ctx, desired)
	case runtimestate.RemediationModeRestartApprovedState:
		err = runtimeControl.remediator.RestartToApprovedState(ctx, desired)
	}
	if err != nil {
		outcome := runtimestate.RemediationOutcome{
			Status:       runtimestate.FindingStatusFailed,
			Mode:         mode,
			AttemptCount: attemptCount,
			Message:      err.Error(),
		}
		writeAuditEvent(ctx, buildRemediationLifecycleEvent(audit.EventTypeDriftRemediationFailed, scanID, desired, result, attemptCount, err.Error()))
		return outcome
	}

	outcome := runtimestate.RemediationOutcome{
		Status:       runtimestate.FindingStatusRemediated,
		Mode:         mode,
		AttemptCount: attemptCount,
		Message:      "approved state restored",
	}
	writeAuditEvent(ctx, buildRemediationLifecycleEvent(audit.EventTypeDriftRemediationSucceeded, scanID, desired, result, attemptCount, outcome.Message))
	return outcome
}

func buildDesiredStateAuditEvent(scanID string, desired runtimestate.ApprovedWorkloadState) audit.Event {
	evidence := &audit.Evidence{
		Runtime: &audit.RuntimeEvidence{
			ClusterID:              desired.ClusterID,
			WorkloadKind:           desired.WorkloadKind,
			ServiceAccountExpected: desired.ServiceAccountName,
			ApprovedLabels:         cloneDesiredLabels(desired.Labels),
			ApprovedDigest:         firstApprovedDigest(desired.Containers),
			ExpectedConfigHash:     desired.ExpectedConfigHash,
			ApprovedContainers:     approvedRuntimeContainers(desired.Containers),
		},
	}
	return audit.Event{
		RequestID:                scanID,
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeDesiredStateRecorded,
		TenantID:                 firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(desired.Namespace)),
		ClusterID:                desired.ClusterID,
		Environment:              audit.EnvironmentFromNamespace(desired.Namespace),
		Namespace:                desired.Namespace,
		WorkloadKind:             desired.WorkloadKind,
		Workload:                 desired.Workload,
		ServiceAccount:           desired.ServiceAccountName,
		DesiredStateSourceRef:    desired.SourceRef,
		DesiredStateApprovalID:   desired.ApprovalCorrelationID,
		DesiredStateVerification: string(desired.DesiredStateVerificationState),
		Decision:                 audit.DecisionAllow,
		Evidence:                 evidence,
	}
}

func buildDriftDetectedEvent(scanID string, desired runtimestate.ApprovedWorkloadState, result runtimestate.ComparisonResult) audit.Event {
	return audit.Event{
		RequestID:                scanID,
		Component:                "runtime-agent",
		EventType:                audit.EventTypeDriftDetected,
		TenantID:                 firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(result.Namespace)),
		ClusterID:                result.ClusterID,
		Environment:              audit.EnvironmentFromNamespace(result.Namespace),
		Namespace:                result.Namespace,
		WorkloadKind:             result.WorkloadKind,
		Workload:                 result.Workload,
		ServiceAccount:           firstNonEmpty(result.ServiceAccountObserved, result.ServiceAccountExpected),
		Image:                    result.Image,
		Digest:                   firstNonEmpty(result.RunningDigest, result.ApprovedDigest),
		Decision:                 audit.DecisionDeny,
		Reasons:                  append([]string(nil), result.Reasons...),
		DriftResult:              result.Result,
		DriftClasses:             append([]string(nil), result.Classes...),
		DriftSeverity:            string(result.Severity),
		Remediable:               result.Remediable,
		RemediationMode:          string(result.SelectedRemediationMode),
		DesiredStateSourceRef:    desired.SourceRef,
		DesiredStateApprovalID:   desired.ApprovalCorrelationID,
		DesiredStateVerification: string(result.DesiredStateVerificationState),
		Evidence:                 audit.FromRuntimeComparison(&result),
	}
}

func buildRemediationLifecycleEvent(eventType string, scanID string, desired runtimestate.ApprovedWorkloadState, result runtimestate.ComparisonResult, attempt int, message string) audit.Event {
	decision := audit.DecisionAllow
	if eventType == audit.EventTypeDriftRemediationFailed {
		decision = audit.DecisionError
	}
	return audit.Event{
		RequestID:                scanID,
		Component:                "runtime-agent",
		EventType:                eventType,
		TenantID:                 firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(result.Namespace)),
		ClusterID:                result.ClusterID,
		Environment:              audit.EnvironmentFromNamespace(result.Namespace),
		Namespace:                result.Namespace,
		WorkloadKind:             result.WorkloadKind,
		Workload:                 result.Workload,
		ServiceAccount:           firstNonEmpty(result.ServiceAccountObserved, result.ServiceAccountExpected),
		Image:                    result.Image,
		Digest:                   firstNonEmpty(result.RunningDigest, result.ApprovedDigest),
		Decision:                 decision,
		Reasons:                  append([]string(nil), message),
		DriftResult:              result.Result,
		DriftClasses:             append([]string(nil), result.Classes...),
		DriftSeverity:            string(result.Severity),
		Remediable:               result.Remediable,
		RemediationMode:          string(result.SelectedRemediationMode),
		RemediationAttempt:       attempt,
		DesiredStateSourceRef:    desired.SourceRef,
		DesiredStateApprovalID:   desired.ApprovalCorrelationID,
		DesiredStateVerification: string(result.DesiredStateVerificationState),
		Evidence:                 audit.FromRuntimeComparison(&result),
	}
}

func buildQuarantineEvent(scanID string, desired runtimestate.ApprovedWorkloadState, result runtimestate.ComparisonResult, outcome runtimestate.RemediationOutcome, quarantineType string) audit.Event {
	event := buildRemediationLifecycleEvent(audit.EventTypeDriftQuarantined, scanID, desired, result, outcome.AttemptCount, firstNonEmpty(outcome.QuarantineReason, outcome.Message))
	event.Decision = audit.DecisionDeny
	event.QuarantineReason = outcome.QuarantineReason
	event.QuarantineType = quarantineType
	return event
}

func buildAuditEvent(scanID string, desired runtimestate.ApprovedWorkloadState, result runtimestate.ComparisonResult) audit.Event {
	decision := audit.DecisionAllow
	if result.HasDrift() {
		decision = audit.DecisionDeny
	}

	return audit.Event{
		RequestID:                scanID,
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeDriftResult,
		TenantID:                 firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(result.Namespace)),
		ClusterID:                result.ClusterID,
		Environment:              audit.EnvironmentFromNamespace(result.Namespace),
		Namespace:                result.Namespace,
		WorkloadKind:             result.WorkloadKind,
		Workload:                 result.Workload,
		ServiceAccount:           firstNonEmpty(result.ServiceAccountObserved, result.ServiceAccountExpected),
		Image:                    result.Image,
		Digest:                   firstNonEmpty(result.RunningDigest, result.ApprovedDigest),
		Decision:                 decision,
		Reasons:                  result.Reasons,
		DriftResult:              result.Result,
		DriftClasses:             append([]string(nil), result.Classes...),
		DriftSeverity:            string(result.Severity),
		Remediable:               result.Remediable,
		RemediationMode:          string(result.SelectedRemediationMode),
		DesiredStateSourceRef:    desired.SourceRef,
		DesiredStateApprovalID:   desired.ApprovalCorrelationID,
		DesiredStateVerification: string(result.DesiredStateVerificationState),
		Evidence:                 audit.FromRuntimeComparison(&result),
	}
}

func writeAuditEvent(ctx context.Context, event audit.Event) {
	metrics.IncDecision("runtime-agent", event.Decision, event.EventType)
	if event.DriftResult == string(runtimestate.DriftClassNoDrift) {
		metrics.IncRuntimeNoDrift("runtime-agent")
	} else if event.DriftResult != "" {
		metrics.IncRuntimeDrift("runtime-agent", event.DriftResult)
	}
	if err := auditWriter.Write(ctx, event); err != nil {
		log.Printf("runtime-agent audit write failed: %v", err)
	}
}

func newRuntimeClients() (runtimestate.StateReader, runtimestate.RemediationClient) {
	if path := strings.TrimSpace(os.Getenv("CHANGELOCK_RUNTIME_FIXTURE")); path != "" {
		reader, err := runtimestate.NewFixtureReader(path)
		if err != nil {
			log.Printf("runtime-agent fixture reader unavailable: %v", err)
			return runtimestate.NoopReader{}, runtimestate.NoopRemediationClient{}
		}
		return reader, runtimestate.NoopRemediationClient{}
	}

	kubeClient, err := runtimestate.NewKubernetesClientFromInCluster()
	if err != nil {
		return runtimestate.NoopReader{}, runtimestate.NoopRemediationClient{}
	}
	return kubeClient, kubeClient
}

func needsKubernetesClient(mode runtimestate.RemediationMode) bool {
	return mode == runtimestate.RemediationModePatchApprovedState || mode == runtimestate.RemediationModeRestartApprovedState
}

func restartSafe(desired runtimestate.ApprovedWorkloadState, observed runtimestate.ObservedWorkloadState) bool {
	if desired.ServiceAccountName != observed.ServiceAccountName {
		return false
	}
	if desired.ExpectedConfigHash != "" && desired.ExpectedConfigHash != observed.ActualConfigHash {
		return false
	}
	observedByName := map[string]runtimestate.ObservedContainerState{}
	for _, container := range observed.Containers {
		observedByName[container.Name] = container
	}
	for _, container := range desired.Containers {
		observedContainer, ok := observedByName[container.Name]
		if !ok {
			return false
		}
		if container.Image != "" && container.Image != observedContainer.Image {
			return false
		}
	}
	return true
}

func firstApprovedDigest(containers []runtimestate.ApprovedContainerState) string {
	for _, container := range containers {
		if container.ApprovedDigest != "" {
			return container.ApprovedDigest
		}
	}
	return ""
}

func approvedRuntimeContainers(containers []runtimestate.ApprovedContainerState) []audit.RuntimeApprovedContainer {
	if len(containers) == 0 {
		return nil
	}
	items := make([]audit.RuntimeApprovedContainer, 0, len(containers))
	for _, container := range containers {
		items = append(items, audit.RuntimeApprovedContainer{
			Name:           container.Name,
			Image:          container.Image,
			ApprovedDigest: container.ApprovedDigest,
			Runtime: audit.RuntimeSecurityConstraints{
				RunAsNonRoot:             container.Runtime.RunAsNonRoot,
				ReadOnlyRootFilesystem:   container.Runtime.ReadOnlyRootFilesystem,
				AllowPrivilegeEscalation: container.Runtime.AllowPrivilegeEscalation,
				DropAllCapabilities:      container.Runtime.DropAllCapabilities,
				SeccompRuntimeDefault:    container.Runtime.SeccompRuntimeDefault,
				DenyPrivileged:           container.Runtime.DenyPrivileged,
			},
		})
	}
	return items
}

func cloneDesiredLabels(labels map[string]string) map[string]string {
	if len(labels) == 0 {
		return nil
	}
	cloned := make(map[string]string, len(labels))
	for key, value := range labels {
		cloned[key] = value
	}
	return cloned
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
