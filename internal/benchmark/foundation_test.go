package benchmark

import "testing"

func TestEvaluateFoundationRegressionDetectsLatencyRegression(t *testing.T) {
	response := EvaluateFoundationRegression(EvaluationRequest{
		ProfileID: "production_like",
		Observations: []Observation{{
			FamilyID:      "deploy_gate_admission",
			ProfileID:     "production_like",
			MetricClass:   "user_facing_latency",
			MetricName:    "p95_latency_ms",
			Unit:          "ms",
			BaselineValue: 100,
			ObservedValue: 121,
		}},
	})

	if response.CurrentState != "failed" || len(response.Results) != 1 || response.Results[0].Status != "regression" {
		t.Fatalf("expected latency regression failure, got %#v", response)
	}
}

func TestEvaluateFoundationRegressionSupportsOverride(t *testing.T) {
	response := EvaluateFoundationRegression(EvaluationRequest{
		ProfileID: "production_like",
		Observations: []Observation{{
			FamilyID:      "audit_ingest",
			ProfileID:     "production_like",
			MetricClass:   "throughput",
			MetricName:    "events_per_second",
			Unit:          "eps",
			BaselineValue: 1000,
			ObservedValue: 700,
		}},
		Override: &Override{Reason: "known noisy CI runner"},
	})

	if response.CurrentState != "passed_with_override" || response.OverrideReason == "" {
		t.Fatalf("expected override-applied response, got %#v", response)
	}
}
