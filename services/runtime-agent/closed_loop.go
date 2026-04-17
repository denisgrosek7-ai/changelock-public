package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
	"github.com/denisgrosek/changelock/internal/signingidentity"
)

type controlPlaneClient struct {
	baseURL string
	token   string
	client  *http.Client
}

type runtimeControlPlane interface {
	enabled() bool
	desiredStates(ctx context.Context) ([]audit.RuntimeDesiredStateView, error)
	activeStates(ctx context.Context) ([]audit.RuntimeActiveStateView, error)
	netVulnerabilities(ctx context.Context, tenantID, environment, imageDigest, severityThreshold string) (audit.VulnerabilityNetResponse, error)
	signingIdentities(ctx context.Context, tenantID, clusterID, environment, imageDigest string) ([]signingidentity.Observation, error)
}

type runtimeDesiredStatesResponse struct {
	Items []audit.RuntimeDesiredStateView `json:"items"`
}

type runtimeActiveStateResponse struct {
	Items []audit.RuntimeActiveStateView `json:"items"`
}

type signingIdentityObservationResponse struct {
	Items []signingidentity.Observation `json:"items"`
}

func newControlPlaneClient() *controlPlaneClient {
	baseURL := strings.TrimRight(strings.TrimSpace(firstNonEmpty(
		os.Getenv("AUDIT_WRITER_URL"),
		os.Getenv("CHANGELOCK_AUDIT_WRITER_URL"),
	)), "/")
	if baseURL == "" {
		return nil
	}
	return &controlPlaneClient{
		baseURL: baseURL,
		token:   strings.TrimSpace(os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")),
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *controlPlaneClient) enabled() bool {
	return c != nil && c.baseURL != ""
}

func (c *controlPlaneClient) desiredStates(ctx context.Context) ([]audit.RuntimeDesiredStateView, error) {
	var response runtimeDesiredStatesResponse
	if err := c.getJSON(ctx, "/v1/runtime/desired-state?limit=500", &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *controlPlaneClient) activeStates(ctx context.Context) ([]audit.RuntimeActiveStateView, error) {
	var response runtimeActiveStateResponse
	if err := c.getJSON(ctx, "/v1/runtime/active-state?limit=500", &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *controlPlaneClient) netVulnerabilities(ctx context.Context, tenantID, environment, imageDigest, severityThreshold string) (audit.VulnerabilityNetResponse, error) {
	query := url.Values{}
	query.Set("image_digest", strings.TrimSpace(imageDigest))
	query.Set("limit", "100")
	if strings.TrimSpace(tenantID) != "" {
		query.Set("tenant_id", strings.TrimSpace(tenantID))
	}
	if strings.TrimSpace(environment) != "" {
		query.Set("environment", strings.TrimSpace(environment))
	}
	if strings.TrimSpace(severityThreshold) != "" {
		query.Set("severity_threshold", strings.TrimSpace(severityThreshold))
	}
	var result audit.VulnerabilityNetResponse
	if err := c.getJSON(ctx, "/v1/vulnerabilities/net?"+query.Encode(), &result); err != nil {
		return audit.VulnerabilityNetResponse{}, err
	}
	return result, nil
}

func (c *controlPlaneClient) signingIdentities(ctx context.Context, tenantID, clusterID, environment, imageDigest string) ([]signingidentity.Observation, error) {
	query := url.Values{}
	query.Set("image_digest", strings.TrimSpace(imageDigest))
	query.Set("limit", "25")
	if strings.TrimSpace(tenantID) != "" {
		query.Set("tenant_id", strings.TrimSpace(tenantID))
	}
	if strings.TrimSpace(clusterID) != "" {
		query.Set("cluster_id", strings.TrimSpace(clusterID))
	}
	if strings.TrimSpace(environment) != "" {
		query.Set("environment", strings.TrimSpace(environment))
	}
	var response signingIdentityObservationResponse
	if err := c.getJSON(ctx, "/v1/signing-identities?"+query.Encode(), &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *controlPlaneClient) getJSON(ctx context.Context, path string, out any) error {
	if !c.enabled() {
		return fmt.Errorf("control-plane client is disabled")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("control-plane request %s returned %s", path, resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (a agentRuntime) startClosedLoop(ctx context.Context) {
	if a.config.Mode == runtimestate.RemediationModeDisabled || a.config.ReconcileInterval <= 0 || a.controlPlane == nil || !a.controlPlane.enabled() {
		return
	}

	go func() {
		a.runClosedLoopPass(ctx)
		ticker := time.NewTicker(a.config.ReconcileInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				a.runClosedLoopPass(ctx)
			}
		}
	}()
}

func (a agentRuntime) runClosedLoopPass(ctx context.Context) {
	desiredStates, err := a.controlPlane.desiredStates(ctx)
	if err != nil {
		log.Printf("runtime-agent closed-loop desired-state load failed: %v", err)
		return
	}
	if len(desiredStates) == 0 {
		return
	}
	activeStates, err := a.controlPlane.activeStates(ctx)
	if err != nil {
		log.Printf("runtime-agent closed-loop active-state load failed: %v", err)
		return
	}
	byID := map[string]audit.RuntimeActiveStateView{}
	for _, item := range activeStates {
		byID[item.ID] = item
	}

	for _, desiredView := range desiredStates {
		a.reconcileDesiredState(ctx, desiredView, byID[desiredView.ID])
	}
}

func (a agentRuntime) reconcileDesiredState(ctx context.Context, desiredView audit.RuntimeDesiredStateView, previous audit.RuntimeActiveStateView) {
	desired := approvedStateFromView(desiredView)
	if desired.Namespace == "" || desired.Workload == "" {
		return
	}

	scanID := audit.NewRequestID()
	observed, err := a.stateReader.ReadObservedWorkload(ctx, runtimestate.WorkloadTarget{
		ClusterID: desired.ClusterID,
		Namespace: desired.Namespace,
		Kind:      desired.WorkloadKind,
		Workload:  desired.Workload,
	})
	if err != nil {
		event := audit.Event{
			RequestID:                scanID,
			Component:                "runtime-agent",
			EventType:                audit.EventTypeRuntimeActiveStateObserved,
			TenantID:                 firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(desired.Namespace)),
			ClusterID:                desired.ClusterID,
			Environment:              audit.EnvironmentFromNamespace(desired.Namespace),
			Namespace:                desired.Namespace,
			WorkloadKind:             desired.WorkloadKind,
			Workload:                 desired.Workload,
			ServiceAccount:           desired.ServiceAccountName,
			Decision:                 audit.DecisionError,
			ReconciliationStatus:     string(runtimestate.ReconciliationStatusFailed),
			DesiredStateSourceRef:    desired.SourceRef,
			DesiredStateApprovalID:   desired.ApprovalCorrelationID,
			DesiredStateVerification: string(desired.DesiredStateVerificationState),
			Reasons:                  []string{"runtime state unavailable: " + err.Error()},
		}
		writeAuditEvent(ctx, event)
		return
	}

	result := runtimestate.Compare(desired, observed)
	result.ScanID = scanID
	result.SelectedRemediationMode, _ = runtimestate.SelectRemediationMode(a.config, desired, result)

	protected, protectedReason := runtimestate.ProtectedTarget(a.config, desired)
	quarantineType := ""
	var remediation *runtimestate.RemediationOutcome

	if signerOutcome := a.runtimeSignerIdentityQuarantine(ctx, desired, result, protected, protectedReason); signerOutcome != nil {
		if signerOutcome.Status == runtimestate.FindingStatusQuarantined {
			outcome := a.quarantineOutcome(ctx, scanID, desired, observed, result, previous.RemediationAttempt, signerOutcome.QuarantineReason, "signer-identity")
			remediation = &outcome
		} else {
			remediation = signerOutcome
		}
		quarantineType = "signer-identity"
	} else if vexOutcome := a.runtimeVEXQuarantine(ctx, desired, result, protected, protectedReason); vexOutcome != nil {
		if vexOutcome.Status == runtimestate.FindingStatusQuarantined {
			outcome := a.quarantineOutcome(ctx, scanID, desired, observed, result, previous.RemediationAttempt, vexOutcome.QuarantineReason, "vex")
			remediation = &outcome
		} else {
			remediation = vexOutcome
		}
		quarantineType = "vex"
	} else if result.HasDrift() {
		writeAuditEvent(ctx, buildDriftDetectedEvent(scanID, desired, result))
		outcome := a.applyClosedLoopRemediation(ctx, scanID, desired, observed, result, previous, protected, protectedReason)
		remediation = &outcome
		if outcome.Status == runtimestate.FindingStatusQuarantined {
			quarantineType = firstNonEmpty(quarantineTypeForOutcome(outcome, result), "drift")
		}
	}

	writeAuditEvent(ctx, buildAuditEvent(scanID, desired, result))
	writeAuditEvent(ctx, buildActiveStateEvent(scanID, desired, observed, result, remediation, protected, protectedReason, quarantineType))
}

func (a agentRuntime) applyClosedLoopRemediation(ctx context.Context, scanID string, desired runtimestate.ApprovedWorkloadState, observed runtimestate.ObservedWorkloadState, result runtimestate.ComparisonResult, previous audit.RuntimeActiveStateView, protected bool, protectedReason string) runtimestate.RemediationOutcome {
	if previous.ReconciliationStatus == string(runtimestate.ReconciliationStatusQuarantined) {
		return runtimestate.RemediationOutcome{
			Status:           runtimestate.FindingStatusQuarantined,
			Mode:             result.SelectedRemediationMode,
			AttemptCount:     previous.RemediationAttempt,
			Quarantined:      true,
			QuarantineReason: firstNonEmpty(previous.QuarantineReason, "workload remains quarantined"),
			Message:          "workload remains quarantined until operator release",
		}
	}
	if protected && result.SelectedRemediationMode != runtimestate.RemediationModeDisabled && result.SelectedRemediationMode != runtimestate.RemediationModeAlertOnly {
		return runtimestate.RemediationOutcome{
			Status:       runtimestate.FindingStatusDetected,
			Mode:         runtimestate.RemediationModeAlertOnly,
			AttemptCount: previous.RemediationAttempt,
			Message:      firstNonEmpty(protectedReason, "protected target blocked automated action"),
		}
	}

	switch result.SelectedRemediationMode {
	case runtimestate.RemediationModeDisabled, runtimestate.RemediationModeAlertOnly:
		return runtimestate.RemediationOutcome{
			Status:       runtimestate.FindingStatusDetected,
			Mode:         result.SelectedRemediationMode,
			AttemptCount: previous.RemediationAttempt,
			Message:      "drift recorded; no automated remediation configured",
		}
	case runtimestate.RemediationModeQuarantine:
		return a.quarantineOutcome(ctx, scanID, desired, observed, result, previous.RemediationAttempt, "drift requires operator review", "drift")
	case runtimestate.RemediationModeRestartApprovedState:
		if !restartSafe(desired, observed) {
			return runtimestate.RemediationOutcome{
				Status:       runtimestate.FindingStatusFailed,
				Mode:         result.SelectedRemediationMode,
				AttemptCount: previous.RemediationAttempt,
				Message:      "restart-to-approved-state requires the controller spec to already match the approved state",
			}
		}
	}

	attemptCount := previous.RemediationAttempt + 1
	if attemptCount > a.config.MaxAttempts {
		return a.quarantineOutcome(ctx, scanID, desired, observed, result, attemptCount, "drift remediation threshold exceeded", "repeat-drift")
	}

	writeAuditEvent(ctx, buildRemediationLifecycleEvent(audit.EventTypeDriftRemediationStarted, scanID, desired, result, attemptCount, "starting remediation"))
	var err error
	switch result.SelectedRemediationMode {
	case runtimestate.RemediationModePatchApprovedState:
		err = a.remediator.PatchApprovedState(ctx, desired)
	case runtimestate.RemediationModeRestartApprovedState:
		err = a.remediator.RestartToApprovedState(ctx, desired)
	}
	if err != nil {
		outcome := runtimestate.RemediationOutcome{
			Status:       runtimestate.FindingStatusFailed,
			Mode:         result.SelectedRemediationMode,
			AttemptCount: attemptCount,
			Message:      err.Error(),
		}
		writeAuditEvent(ctx, buildRemediationLifecycleEvent(audit.EventTypeDriftRemediationFailed, scanID, desired, result, attemptCount, err.Error()))
		return outcome
	}

	outcome := runtimestate.RemediationOutcome{
		Status:       runtimestate.FindingStatusRemediated,
		Mode:         result.SelectedRemediationMode,
		AttemptCount: attemptCount,
		Message:      "approved state restored",
	}
	writeAuditEvent(ctx, buildRemediationLifecycleEvent(audit.EventTypeDriftRemediationSucceeded, scanID, desired, result, attemptCount, outcome.Message))
	return outcome
}

func (a agentRuntime) quarantineOutcome(ctx context.Context, scanID string, desired runtimestate.ApprovedWorkloadState, observed runtimestate.ObservedWorkloadState, result runtimestate.ComparisonResult, attemptCount int, reason string, quarantineType string) runtimestate.RemediationOutcome {
	message := firstNonEmpty(reason, "workload quarantined for operator review")
	if a.config.QuarantineNetworkPolicy {
		if err := a.remediator.ApplyQuarantineOverlay(ctx, desired, observed); err != nil {
			writeAuditEvent(ctx, buildRemediationLifecycleEvent(audit.EventTypeDriftRemediationFailed, scanID, desired, result, attemptCount, err.Error()))
			return runtimestate.RemediationOutcome{
				Status:       runtimestate.FindingStatusFailed,
				Mode:         runtimestate.RemediationModeQuarantine,
				AttemptCount: attemptCount,
				Message:      err.Error(),
			}
		}
	}
	outcome := runtimestate.RemediationOutcome{
		Status:           runtimestate.FindingStatusQuarantined,
		Mode:             runtimestate.RemediationModeQuarantine,
		AttemptCount:     attemptCount,
		Quarantined:      true,
		QuarantineReason: message,
		Message:          message,
	}
	writeAuditEvent(ctx, buildQuarantineEvent(scanID, desired, result, outcome, quarantineType))
	return outcome
}

func (a agentRuntime) runtimeVEXQuarantine(ctx context.Context, desired runtimestate.ApprovedWorkloadState, result runtimestate.ComparisonResult, protected bool, protectedReason string) *runtimestate.RemediationOutcome {
	if !a.config.RuntimeVEXQuarantine || a.controlPlane == nil || !a.controlPlane.enabled() {
		return nil
	}
	imageDigest := firstNonEmpty(result.RunningDigest, result.ApprovedDigest)
	if imageDigest == "" {
		return nil
	}
	netResult, err := a.controlPlane.netVulnerabilities(ctx, firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(desired.Namespace)), audit.EnvironmentFromNamespace(desired.Namespace), imageDigest, a.config.RuntimeVEXSeverity)
	if err != nil {
		log.Printf("runtime-agent VEX quarantine lookup failed for %s/%s: %v", desired.Namespace, desired.Workload, err)
		return nil
	}
	if a.config.RuntimeVEXRequireNet && !netResult.ThresholdBreached {
		return nil
	}
	if netResult.ActionableCount == 0 {
		return nil
	}
	if protected {
		return &runtimestate.RemediationOutcome{
			Status:       runtimestate.FindingStatusFailed,
			Mode:         runtimestate.RemediationModeAlertOnly,
			Message:      firstNonEmpty(protectedReason, "protected target blocked VEX-driven quarantine"),
			AttemptCount: 0,
		}
	}
	outcome := runtimestate.RemediationOutcome{
		Status:           runtimestate.FindingStatusQuarantined,
		Mode:             runtimestate.RemediationModeQuarantine,
		Quarantined:      true,
		QuarantineReason: fmt.Sprintf("net actionable %s vulnerability requires containment", strings.ToLower(a.config.RuntimeVEXSeverity)),
		Message:          fmt.Sprintf("%d vulnerabilities remain net actionable after VEX merge", netResult.ActionableCount),
	}
	return &outcome
}

func (a agentRuntime) runtimeSignerIdentityQuarantine(ctx context.Context, desired runtimestate.ApprovedWorkloadState, result runtimestate.ComparisonResult, protected bool, protectedReason string) *runtimestate.RemediationOutcome {
	if !a.identityConfig.QuarantineOnDrift || a.identityConfig.Enforcement == signingidentity.EnforcementDisabled || a.controlPlane == nil || !a.controlPlane.enabled() {
		return nil
	}
	imageDigest := firstNonEmpty(result.RunningDigest, result.ApprovedDigest)
	if imageDigest == "" {
		return nil
	}
	observations, err := a.controlPlane.signingIdentities(ctx, firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(desired.Namespace)), desired.ClusterID, audit.EnvironmentFromNamespace(desired.Namespace), imageDigest)
	if err != nil {
		log.Printf("runtime-agent signer identity lookup failed for %s/%s: %v", desired.Namespace, desired.Workload, err)
		return nil
	}
	for _, observation := range observations {
		if observation.Authorized != signingidentity.AuthorizationUnauthorized {
			continue
		}
		if protected {
			return &runtimestate.RemediationOutcome{
				Status:       runtimestate.FindingStatusFailed,
				Mode:         runtimestate.RemediationModeAlertOnly,
				Message:      firstNonEmpty(protectedReason, "protected target blocked signer-identity quarantine"),
				AttemptCount: 0,
			}
		}
		reason := firstNonEmpty(observation.ReasonDetail, "artifact signer identity is unauthorized")
		outcome := runtimestate.RemediationOutcome{
			Status:           runtimestate.FindingStatusQuarantined,
			Mode:             runtimestate.RemediationModeQuarantine,
			Quarantined:      true,
			QuarantineReason: reason,
			Message:          reason,
		}
		return &outcome
	}
	return nil
}

func approvedStateFromView(view audit.RuntimeDesiredStateView) runtimestate.ApprovedWorkloadState {
	desired := runtimestate.ApprovedWorkloadState{
		TenantID:                      view.TenantID,
		ClusterID:                     view.ClusterID,
		Namespace:                     view.Namespace,
		WorkloadKind:                  view.WorkloadKind,
		Workload:                      view.Workload,
		ServiceAccountName:            view.ServiceAccount,
		ExpectedConfigHash:            view.ExpectedConfigHash,
		ApprovalCorrelationID:         view.DesiredStateApprovalID,
		SourceRef:                     view.DesiredStateSourceRef,
		DesiredStateVerificationState: runtimestate.VerificationState(view.DesiredStateVerification),
		Labels:                        cloneLabels(view.Labels),
	}
	for _, container := range view.Containers {
		desired.Containers = append(desired.Containers, runtimestate.ApprovedContainerState{
			Name:           container.Name,
			Image:          container.Image,
			ApprovedDigest: container.ApprovedDigest,
			Runtime: runtimestate.SecurityConstraints{
				RunAsNonRoot:             container.Runtime.RunAsNonRoot,
				ReadOnlyRootFilesystem:   container.Runtime.ReadOnlyRootFilesystem,
				AllowPrivilegeEscalation: container.Runtime.AllowPrivilegeEscalation,
				DropAllCapabilities:      container.Runtime.DropAllCapabilities,
				SeccompRuntimeDefault:    container.Runtime.SeccompRuntimeDefault,
				DenyPrivileged:           container.Runtime.DenyPrivileged,
			},
		})
	}
	return desired
}

func buildActiveStateEvent(scanID string, desired runtimestate.ApprovedWorkloadState, observed runtimestate.ObservedWorkloadState, result runtimestate.ComparisonResult, outcome *runtimestate.RemediationOutcome, protected bool, protectedReason string, quarantineType string) audit.Event {
	reconciliationStatus := string(runtimestate.ReconciliationStatusFromOutcome(result.HasDrift(), outcome))
	reasons := append([]string(nil), result.Reasons...)
	decision := audit.DecisionAllow
	if reconciliationStatus == string(runtimestate.ReconciliationStatusFailed) {
		decision = audit.DecisionError
	} else if reconciliationStatus != string(runtimestate.ReconciliationStatusInSync) {
		decision = audit.DecisionDeny
	}
	if outcome != nil && outcome.Message != "" && len(reasons) == 0 {
		reasons = append(reasons, outcome.Message)
	}
	evidence := audit.FromRuntimeComparison(&result)
	if evidence == nil {
		evidence = &audit.Evidence{Runtime: &audit.RuntimeEvidence{}}
	}
	if evidence.Runtime == nil {
		evidence.Runtime = &audit.RuntimeEvidence{}
	}
	evidence.Runtime.ClusterID = desired.ClusterID
	evidence.Runtime.WorkloadKind = desired.WorkloadKind
	evidence.Runtime.ServiceAccountExpected = desired.ServiceAccountName
	evidence.Runtime.ServiceAccountObserved = observed.ServiceAccountName
	evidence.Runtime.ApprovedLabels = cloneLabels(desired.Labels)
	evidence.Runtime.ApprovedContainers = approvedRuntimeContainers(desired.Containers)

	attemptCount := 0
	remediationMode := string(result.SelectedRemediationMode)
	quarantineReason := ""
	if outcome != nil {
		attemptCount = outcome.AttemptCount
		if outcome.Mode != "" {
			remediationMode = string(outcome.Mode)
		}
		quarantineReason = outcome.QuarantineReason
	}

	return audit.Event{
		RequestID:                scanID,
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeActiveStateObserved,
		TenantID:                 firstNonEmpty(desired.TenantID, audit.TenantFromNamespace(desired.Namespace)),
		ClusterID:                desired.ClusterID,
		Environment:              audit.EnvironmentFromNamespace(desired.Namespace),
		Namespace:                desired.Namespace,
		WorkloadKind:             desired.WorkloadKind,
		Workload:                 desired.Workload,
		ServiceAccount:           firstNonEmpty(observed.ServiceAccountName, desired.ServiceAccountName),
		Image:                    result.Image,
		Digest:                   firstNonEmpty(result.RunningDigest, result.ApprovedDigest),
		Decision:                 decision,
		Reasons:                  reasons,
		DriftResult:              result.Result,
		DriftClasses:             append([]string(nil), result.Classes...),
		DriftSeverity:            string(result.Severity),
		ReconciliationStatus:     reconciliationStatus,
		RemediationMode:          remediationMode,
		RemediationAttempt:       attemptCount,
		Remediable:               result.Remediable,
		QuarantineReason:         quarantineReason,
		QuarantineType:           quarantineType,
		ProtectedTarget:          protected,
		ProtectedReason:          protectedReason,
		DesiredStateSourceRef:    desired.SourceRef,
		DesiredStateApprovalID:   desired.ApprovalCorrelationID,
		DesiredStateVerification: string(desired.DesiredStateVerificationState),
		Evidence:                 evidence,
	}
}

func cloneLabels(values map[string]string) map[string]string {
	if len(values) == 0 {
		return nil
	}
	cloned := make(map[string]string, len(values))
	for key, value := range values {
		cloned[key] = value
	}
	return cloned
}

func quarantineTypeForOutcome(outcome runtimestate.RemediationOutcome, result runtimestate.ComparisonResult) string {
	if outcome.QuarantineReason != "" && strings.Contains(strings.ToLower(outcome.QuarantineReason), "threshold") {
		return "repeat-drift"
	}
	if result.DesiredStateVerificationState != runtimestate.VerificationStateVerified {
		return "trust"
	}
	return "drift"
}
