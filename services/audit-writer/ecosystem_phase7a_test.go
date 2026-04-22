package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPhase7DeveloperWorkbenchHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/developer/workbench", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected workbench 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7DeveloperWorkbenchResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode workbench: %v", err)
	}
	if response.CurrentState != phase7DeveloperWorkbenchStateActive {
		t.Fatalf("expected active developer workbench, got %#v", response)
	}
	if !containsString(response.OutputSemantics, "uncertainty") || !containsString(response.OutputSemantics, "observed_fact") {
		t.Fatalf("expected bounded semantics in workbench, got %#v", response.OutputSemantics)
	}
	if !hasDeveloperCommand(response.Commands, "local_check") || !hasDeveloperCommand(response.Commands, "post_run_guidance") {
		t.Fatalf("expected local and guidance commands, got %#v", response.Commands)
	}
	if !containsString(response.RouteRefs, "/v1/ecosystem/phase7/developer/context") {
		t.Fatalf("expected developer context route ref, got %#v", response.RouteRefs)
	}
}

func TestPhase7DeveloperContextHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/developer/context", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected context 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7DeveloperContextResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode context: %v", err)
	}
	if response.CurrentState != phase7DeveloperContextStateActive {
		t.Fatalf("expected active context pack, got %#v", response)
	}
	if !hasDeveloperContextClass(response.Items, "observed_fact") || !hasDeveloperContextClass(response.Items, "derived_relevance") || !hasDeveloperContextClass(response.Items, "recommendation") {
		t.Fatalf("expected observed, derived, and recommendation classes, got %#v", response.Items)
	}
	if !developerContextHasRoute(response.Items, "vex_relevance_context", "/v1/intelligence/vulnerability-relevance") {
		t.Fatalf("expected VEX context route refs, got %#v", response.Items)
	}
}

func TestPhase7DeveloperPreCommitHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/developer/pre-commit", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected pre-commit 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7DeveloperPreCommitResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode pre-commit: %v", err)
	}
	if response.CurrentState != phase7DeveloperPreCommitStateActive {
		t.Fatalf("expected active pre-commit profile, got %#v", response)
	}
	if response.DisablePath != "CHGLOCK_PRECOMMIT_DISABLE=1" {
		t.Fatalf("expected explicit disable path, got %#v", response)
	}
	if response.LatencyBudget != 1500 {
		t.Fatalf("expected 1500ms latency budget, got %#v", response)
	}
	if !containsString(response.FailSafeBehaviors, "Pre-commit stays non-mutating and cannot auto-open PRs or apply policy changes.") {
		t.Fatalf("expected bounded non-mutating fail-safe behavior, got %#v", response.FailSafeBehaviors)
	}
	if !hasDeveloperCommand(response.Commands, "pre_commit_check") || !hasDeveloperCommand(response.Commands, "pre_commit_preview") {
		t.Fatalf("expected pre-commit commands, got %#v", response.Commands)
	}
}

func hasDeveloperCommand(items []phase7DeveloperCommand, commandID string) bool {
	for _, item := range items {
		if item.CommandID == commandID {
			return true
		}
	}
	return false
}

func hasDeveloperContextClass(items []phase7DeveloperContextItem, class string) bool {
	for _, item := range items {
		if item.SignalClass == class {
			return true
		}
	}
	return false
}

func developerContextHasRoute(items []phase7DeveloperContextItem, contextID, route string) bool {
	for _, item := range items {
		if item.ContextID != contextID {
			continue
		}
		for _, ref := range item.RouteRefs {
			if ref == route {
				return true
			}
		}
	}
	return false
}
