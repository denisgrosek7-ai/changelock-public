package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
)

type fakeStateReader struct {
	state runtimestate.ObservedWorkloadState
	err   error
}

func (f fakeStateReader) ReadObservedWorkload(_ context.Context, _ runtimestate.WorkloadTarget) (runtimestate.ObservedWorkloadState, error) {
	return f.state, f.err
}

type fakeRemediator struct {
	patchCalls      int
	restartCalls    int
	quarantineCalls int
	err             error
}

func (f *fakeRemediator) PatchApprovedState(_ context.Context, _ runtimestate.ApprovedWorkloadState) error {
	f.patchCalls++
	return f.err
}

func (f *fakeRemediator) RestartToApprovedState(_ context.Context, _ runtimestate.ApprovedWorkloadState) error {
	f.restartCalls++
	return f.err
}

func (f *fakeRemediator) ApplyQuarantineOverlay(_ context.Context, _ runtimestate.ApprovedWorkloadState, _ runtimestate.ObservedWorkloadState) error {
	f.quarantineCalls++
	return f.err
}

func TestScanAllowsNoDrift(t *testing.T) {
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousWriter := auditWriter
	previousRuntime := runtimeControl
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	runtimeControl = agentRuntime{
		config:      runtimestate.SelfHealingConfig{Mode: runtimestate.RemediationModeDisabled},
		stateReader: fakeStateReader{},
		remediator:  runtimestate.NoopRemediationClient{},
		tracker:     runtimestate.NewTracker(),
	}
	defer func() {
		auditWriter = previousWriter
		runtimeControl = previousRuntime
	}()

	payload := scanRequest{
		ScanID:   "scan-allow",
		Approved: approvedState(),
		Observed: observedState(),
	}

	response := executeScanRequest(t, payload)
	if response.Result.Result != string(runtimestate.DriftClassNoDrift) {
		t.Fatalf("expected no_drift, got %#v", response.Result)
	}

	events := readAuditEvents(t, auditPath)
	event := findDecisionEvent(events, audit.EventTypeRuntimeDriftResult, audit.DecisionAllow)
	if event == nil || event.DriftResult != string(runtimestate.DriftClassNoDrift) {
		t.Fatalf("expected ALLOW runtime drift event, got %#v", events)
	}
}

func TestScanDeniesDriftedWorkload(t *testing.T) {
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousWriter := auditWriter
	previousRuntime := runtimeControl
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	runtimeControl = agentRuntime{
		config:      runtimestate.SelfHealingConfig{Mode: runtimestate.RemediationModeAlertOnly},
		stateReader: fakeStateReader{},
		remediator:  runtimestate.NoopRemediationClient{},
		tracker:     runtimestate.NewTracker(),
	}
	defer func() {
		auditWriter = previousWriter
		runtimeControl = previousRuntime
	}()

	observed := observedState()
	observed.Containers[0].RunningDigest = "sha256:mutated"
	observed.Containers[0].Runtime.Privileged = true

	response := executeScanRequest(t, scanRequest{
		ScanID:   "scan-deny",
		Approved: approvedState(),
		Observed: observed,
	})
	if response.Result.Result != string(runtimestate.DriftClassMultiple) {
		t.Fatalf("expected multiple_drift, got %#v", response.Result)
	}

	events := readAuditEvents(t, auditPath)
	event := findDecisionEvent(events, audit.EventTypeRuntimeDriftResult, audit.DecisionDeny)
	if event == nil || len(event.Reasons) == 0 || event.Evidence == nil || event.Evidence.Runtime == nil {
		t.Fatalf("expected explainable DENY runtime drift event, got %#v", events)
	}
	if response.Remediation == nil || response.Remediation.Mode != runtimestate.RemediationModeAlertOnly {
		t.Fatalf("expected alert-only remediation outcome, got %#v", response.Remediation)
	}
}

func TestScanPatchApprovedStateRemediatesWithMockedClient(t *testing.T) {
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousWriter := auditWriter
	previousRuntime := runtimeControl
	remediator := &fakeRemediator{}
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	runtimeControl = agentRuntime{
		config: runtimestate.SelfHealingConfig{
			Mode:         runtimestate.RemediationModePatchApprovedState,
			MaxAttempts:  3,
			Window:       time.Minute,
			AllowedKinds: map[string]struct{}{"Deployment": {}},
		},
		stateReader: fakeStateReader{},
		remediator:  remediator,
		tracker:     runtimestate.NewTracker(),
	}
	defer func() {
		auditWriter = previousWriter
		runtimeControl = previousRuntime
	}()

	observed := observedState()
	observed.Containers[0].RunningDigest = "sha256:mutated"

	response := executeScanRequest(t, scanRequest{
		ScanID:   "scan-remediate",
		Approved: approvedState(),
		Observed: observed,
	})

	if remediator.patchCalls != 1 {
		t.Fatalf("expected one patch remediation call, got %d", remediator.patchCalls)
	}
	if response.Remediation == nil || response.Remediation.Status != runtimestate.FindingStatusRemediated {
		t.Fatalf("expected remediated outcome, got %#v", response.Remediation)
	}
	events := readAuditEvents(t, auditPath)
	if findDecisionEvent(events, audit.EventTypeDriftRemediationSucceeded, audit.DecisionAllow) == nil {
		t.Fatalf("expected remediation succeeded audit event, got %#v", events)
	}
}

func TestScanQuarantinesWhenSignedDesiredStateIsRequired(t *testing.T) {
	previousRuntime := runtimeControl
	runtimeControl = agentRuntime{
		config: runtimestate.SelfHealingConfig{
			Mode:                      runtimestate.RemediationModePatchApprovedState,
			MaxAttempts:               3,
			Window:                    time.Minute,
			AllowedKinds:              map[string]struct{}{"Deployment": {}},
			RequireSignedDesiredState: true,
		},
		stateReader: fakeStateReader{},
		remediator:  &fakeRemediator{},
		tracker:     runtimestate.NewTracker(),
	}
	defer func() {
		runtimeControl = previousRuntime
	}()

	approved := approvedState()
	approved.DesiredStateVerificationState = runtimestate.VerificationStateUnverified
	observed := observedState()
	observed.Containers[0].RunningDigest = "sha256:mutated"

	response := executeScanRequest(t, scanRequest{
		ScanID:   "scan-quarantine",
		Approved: approved,
		Observed: observed,
	})

	if response.Remediation == nil || response.Remediation.Status != runtimestate.FindingStatusQuarantined {
		t.Fatalf("expected quarantined outcome, got %#v", response.Remediation)
	}
}

func executeScanRequest(t *testing.T, payload scanRequest) scanResponse {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/scan", bytes.NewReader(body))
	recorder := httptest.NewRecorder()

	newHandler().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("unexpected status code %d: %s", recorder.Code, recorder.Body.String())
	}

	var response scanResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	return response
}

func approvedState() runtimestate.ApprovedWorkloadState {
	return runtimestate.ApprovedWorkloadState{
		Namespace:          "acme-prod",
		WorkloadKind:       "Deployment",
		Workload:           "runtime-agent",
		ServiceAccountName: "runtime-agent",
		ExpectedConfigHash: "cfg-123",
		Containers: []runtimestate.ApprovedContainerState{
			{
				Name:           "agent",
				Image:          "ghcr.io/my-org/runtime-agent@sha256:abc123",
				ApprovedDigest: "sha256:abc123",
				Runtime: runtimestate.SecurityConstraints{
					RunAsNonRoot:             true,
					ReadOnlyRootFilesystem:   true,
					AllowPrivilegeEscalation: false,
					DropAllCapabilities:      true,
					SeccompRuntimeDefault:    true,
					DenyPrivileged:           true,
				},
			},
		},
	}
}

func observedState() *runtimestate.ObservedWorkloadState {
	return &runtimestate.ObservedWorkloadState{
		Namespace:          "acme-prod",
		WorkloadKind:       "Deployment",
		Workload:           "runtime-agent",
		ServiceAccountName: "runtime-agent",
		ActualConfigHash:   "cfg-123",
		Containers: []runtimestate.ObservedContainerState{
			{
				Name:          "agent",
				Image:         "ghcr.io/my-org/runtime-agent@sha256:abc123",
				RunningDigest: "sha256:abc123",
				Runtime: runtimestate.SecurityPosture{
					RunAsNonRoot:             true,
					ReadOnlyRootFilesystem:   true,
					AllowPrivilegeEscalation: false,
					DropAllCapabilities:      true,
					SeccompRuntimeDefault:    true,
					Privileged:               false,
				},
			},
		},
	}
}

func readAuditEvents(t *testing.T, path string) []audit.Event {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	events := make([]audit.Event, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var event audit.Event
		if err := json.Unmarshal(line, &event); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		events = append(events, event)
	}

	return events
}

func findDecisionEvent(events []audit.Event, eventType, decision string) *audit.Event {
	for _, event := range events {
		if event.EventType == eventType && event.Decision == decision {
			eventCopy := event
			return &eventCopy
		}
	}
	return nil
}
