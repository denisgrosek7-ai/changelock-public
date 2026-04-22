package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPhase7OSSConnectorsHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/oss/connectors", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected oss connectors 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7OSSConnectorsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode oss connectors: %v", err)
	}
	if response.CurrentState != phase7OSSConnectorsStateActive {
		t.Fatalf("expected active oss connectors response, got %#v", response)
	}
	if len(response.Connectors) != 3 {
		t.Fatalf("expected three bounded registry connectors, got %#v", response.Connectors)
	}
	for _, item := range response.Connectors {
		if item.ObservationState != "candidate_only_until_review" {
			t.Fatalf("expected connector observation state to stay candidate-only, got %#v", item)
		}
		if !containsString(item.RouteRefs, "/v1/ecosystem/phase7/oss/observations") {
			t.Fatalf("expected connector route ref to observations, got %#v", item.RouteRefs)
		}
	}
	if len(response.CompatibilityRefs) == 0 || len(response.AbuseControlRefs) == 0 {
		t.Fatalf("expected compatibility and abuse refs, got %#v", response)
	}
}

func TestPhase7OSSObservationsHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/oss/observations", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected oss observations 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7OSSObservationsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode oss observations: %v", err)
	}
	if response.CurrentState != phase7OSSObservationsStateActive {
		t.Fatalf("expected active observation feed, got %#v", response)
	}
	if response.CommunityInputState != "community_candidate_only_review_required" {
		t.Fatalf("expected community input to remain candidate-only, got %#v", response)
	}
	if len(response.Observations) != 3 {
		t.Fatalf("expected three observation examples, got %#v", response.Observations)
	}
	if !hasOSSObservationState(response.Observations, "blocked_candidate") {
		t.Fatalf("expected blocked candidate example, got %#v", response.Observations)
	}
	if hasOSSObservationState(response.Observations, "reviewed") {
		t.Fatalf("did not expect reviewed publication in observation feed, got %#v", response.Observations)
	}
}

func TestPhase7OSSReviewFlowAndReviewedSignalsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	reviewReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/oss/review-flow", nil)
	reviewRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(reviewRec, reviewReq)
	if reviewRec.Code != http.StatusOK {
		t.Fatalf("expected oss review flow 200, got %d: %s", reviewRec.Code, reviewRec.Body.String())
	}

	var review phase7OSSReviewFlowResponse
	if err := json.NewDecoder(reviewRec.Body).Decode(&review); err != nil {
		t.Fatalf("decode review flow: %v", err)
	}
	if review.CurrentState != phase7OSSReviewFlowStateActive {
		t.Fatalf("expected active review flow, got %#v", review)
	}
	if !containsString(review.ReviewStates, "reviewed") || !containsString(review.ReviewStates, "revoked") {
		t.Fatalf("expected reviewed and revoked lifecycle states, got %#v", review.ReviewStates)
	}
	if !containsString(review.ExpandedScopeDeferred, "automated_pr_discipline") {
		t.Fatalf("expected automated PR to stay deferred, got %#v", review.ExpandedScopeDeferred)
	}

	signalsReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/oss/reviewed-signals", nil)
	signalsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(signalsRec, signalsReq)
	if signalsRec.Code != http.StatusOK {
		t.Fatalf("expected reviewed signals 200, got %d: %s", signalsRec.Code, signalsRec.Body.String())
	}

	var signals phase7OSSReviewedSignalsResponse
	if err := json.NewDecoder(signalsRec.Body).Decode(&signals); err != nil {
		t.Fatalf("decode reviewed signals: %v", err)
	}
	if signals.CurrentState != phase7OSSReviewedSignalsStateActive {
		t.Fatalf("expected active reviewed signals pack, got %#v", signals)
	}
	if len(signals.Records) != 3 {
		t.Fatalf("expected three reviewed-signal lifecycle records, got %#v", signals.Records)
	}
	if !hasReviewedSignalState(signals.Records, "reviewed") || !hasReviewedSignalState(signals.Records, "superseded") || !hasReviewedSignalState(signals.Records, "revoked") {
		t.Fatalf("expected reviewed, superseded, and revoked records, got %#v", signals.Records)
	}
	if !containsString(signals.RouteRefs, "/v1/public/proof-portal") {
		t.Fatalf("expected proof portal route ref, got %#v", signals.RouteRefs)
	}
}

func hasOSSObservationState(items []phase7OSSObservationItem, state string) bool {
	for _, item := range items {
		if item.ReviewState == state {
			return true
		}
	}
	return false
}

func hasReviewedSignalState(items []phase7OSSReviewedSignalRecord, state string) bool {
	for _, item := range items {
		if item.ReviewState == state {
			return true
		}
	}
	return false
}
