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

func TestScanAllowsNoDrift(t *testing.T) {
	auditPath := filepath.Join(t.TempDir(), "audit.jsonl")
	previousWriter := auditWriter
	previousReader := stateReader
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	stateReader = fakeStateReader{}
	defer func() {
		auditWriter = previousWriter
		stateReader = previousReader
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
	previousReader := stateReader
	auditWriter = audit.NewWriter(audit.NewFileSink(auditPath))
	stateReader = fakeStateReader{}
	defer func() {
		auditWriter = previousWriter
		stateReader = previousReader
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
		Workload:           "runtime-agent",
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
		Namespace:        "acme-prod",
		Workload:         "runtime-agent",
		ActualConfigHash: "cfg-123",
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
