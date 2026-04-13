package audit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPSinkWriteSuccess(t *testing.T) {
	var captured Event
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/ingest" {
			t.Fatalf("unexpected path %s", r.URL.Path)
		}
		if r.Header.Get("X-Request-Id") != "req-123" {
			t.Fatalf("expected request id header, got %q", r.Header.Get("X-Request-Id"))
		}
		if err := json.NewDecoder(r.Body).Decode(&captured); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	sink := NewHTTPSink(server.URL, time.Second)
	err := sink.Write(context.Background(), Event{
		RequestID: "req-123",
		Component: "deploy-gate",
		EventType: EventTypeDeployGateDecision,
		Decision:  DecisionAllow,
	})
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if captured.Component != "deploy-gate" || captured.Decision != DecisionAllow {
		t.Fatalf("unexpected captured event %#v", captured)
	}
}

func TestHTTPSinkWriteReturnsRemoteFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "unavailable", http.StatusServiceUnavailable)
	}))
	defer server.Close()

	sink := NewHTTPSink(server.URL, time.Second)
	err := sink.Write(context.Background(), Event{
		Component: "deploy-gate",
		EventType: EventTypeDeployGateDecision,
		Decision:  DecisionDeny,
	})
	if err == nil {
		t.Fatalf("expected remote failure")
	}
}
