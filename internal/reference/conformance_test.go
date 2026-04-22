package reference

import "testing"

func TestEvaluateKeepsUnsupportedOptionalCheckVisibleWithoutBlockingActive(t *testing.T) {
	result := Evaluate(Input{
		ArchitectureID: "runtime-hardened-enterprise-cluster",
		Checks: []Check{
			{CheckID: "transparency_anchor", Required: true, State: "active"},
			{CheckID: "formal_certification", Required: false, State: "unsupported"},
		},
	})

	if result.CurrentState != StateActive {
		t.Fatalf("expected active conformance with visible unsupported boundary, got %#v", result)
	}
	if len(result.UnsupportedChecks) != 1 || result.UnsupportedChecks[0] != "formal_certification" {
		t.Fatalf("expected unsupported optional check to stay visible, got %#v", result)
	}
}

func TestEvaluateReturnsPartialForOptionalDegradedCheck(t *testing.T) {
	result := Evaluate(Input{
		ArchitectureID: "runtime-hardened-enterprise-cluster",
		Checks: []Check{
			{CheckID: "transparency_anchor", Required: true, State: "active"},
			{CheckID: "benchmark_packs", Required: false, State: "degraded"},
		},
	})

	if result.CurrentState != StatePartial {
		t.Fatalf("expected partial conformance, got %#v", result)
	}
	if len(result.DegradedChecks) != 1 || result.DegradedChecks[0] != "benchmark_packs" {
		t.Fatalf("expected degraded optional check to stay visible, got %#v", result)
	}
}

func TestEvaluateReturnsIncompleteForRequiredDeviation(t *testing.T) {
	result := Evaluate(Input{
		ArchitectureID: "runtime-hardened-enterprise-cluster",
		Checks: []Check{
			{CheckID: "transparency_anchor", Required: true, State: "active"},
			{CheckID: "verifier_sdk", Required: true, State: "missing"},
		},
	})

	if result.CurrentState != StateIncomplete {
		t.Fatalf("expected incomplete conformance, got %#v", result)
	}
}
