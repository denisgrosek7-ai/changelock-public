package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

func mustMarshalPoint13ValEFoundation(t *testing.T, model Point13ValEFoundation) string {
	t.Helper()
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal point13 vale foundation: %v", err)
	}
	return string(payload)
}

func clonePoint13ValEFoundation(t *testing.T, model Point13ValEFoundation) Point13ValEFoundation {
	t.Helper()
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal point13 vale clone: %v", err)
	}
	var cloned Point13ValEFoundation
	if err := json.Unmarshal(payload, &cloned); err != nil {
		t.Fatalf("unmarshal point13 vale clone: %v", err)
	}
	return cloned
}

var (
	point13ValEActiveFoundationBaseline     Point13ValEFoundation
	point13ValEActiveFoundationBaselineOnce sync.Once
)

func uncachedActivePoint13ValEFoundation() Point13ValEFoundation {
	return ComputePoint13ValEFoundation(Point13ValEFoundationModel())
}

func activePoint13ValEFoundation(t *testing.T) Point13ValEFoundation {
	t.Helper()
	point13ValEActiveFoundationBaselineOnce.Do(func() {
		point13ValEActiveFoundationBaseline = uncachedActivePoint13ValEFoundation()
	})
	return clonePoint13ValEFoundation(t, point13ValEActiveFoundationBaseline)
}

func TestPoint13ValEFoundationFixtureIsolation(t *testing.T) {
	t.Run("raw production path computes final pass confirmed closure", func(t *testing.T) {
		model := uncachedActivePoint13ValEFoundation()
		if model.CurrentState != Point13ValEStatePassConfirmed {
			t.Fatalf("expected pass confirmed raw ValE foundation, got %#v", model)
		}
		payload := mustMarshalPoint13ValEFoundation(t, model)
		if !strings.Contains(payload, point13ValEPoint13PassToken) {
			t.Fatalf("expected point_13_pass token in final ValE payload, got %s", payload)
		}
	})

	t.Run("cached fixture mutation does not contaminate next clone", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.PassClosureManifest.ClosureManifestID = "closure_manifest_point13_vale_mutated_001"
		next := activePoint13ValEFoundation(t)
		if next.PassClosureManifest.ClosureManifestID != "closure_manifest_point13_vale_001" {
			t.Fatalf("expected cached baseline isolation, got %#v", next)
		}
	})
}

func TestPoint13ValEDependencyState(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point13ValEFoundation)
	}{
		{name: "missing vald dependency blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.SnapshotFromComputedOutput = false
		}},
		{name: "vald blocked blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValDCurrentState = Point13ValDStateBlocked
			model.Dependency.ValD.CurrentState = Point13ValDStateBlocked
		}},
		{name: "valc blocked blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.InheritedValCCurrentState = Point13ValCStateBlocked
			model.Dependency.ValD.Dependency.ValCCurrentState = Point13ValCStateBlocked
		}},
		{name: "valb blocked blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.InheritedValBCurrentState = Point13ValBStateBlocked
			model.Dependency.ValD.Dependency.InheritedValBCurrentState = Point13ValBStateBlocked
		}},
		{name: "vala blocked blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.InheritedValACurrentState = Point13ValAStateBlocked
			model.Dependency.ValD.Dependency.InheritedValACurrentState = Point13ValAStateBlocked
		}},
		{name: "val0 blocked blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.InheritedVal0CurrentState = Point13Val0StateBlocked
			model.Dependency.ValD.Dependency.InheritedVal0CurrentState = Point13Val0StateBlocked
		}},
		{name: "point13 pass before vale blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValDPoint13PassSeen = true
		}},
		{name: "inherited point12 closure mismatch blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.InheritedPoint12ReviewerResult = point12ValEReviewerResultReviewRequired
			model.Dependency.ValD.Dependency.InheritedPoint12ReviewerResult = point12ValEReviewerResultReviewRequired
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValEFoundation(t)
			tc.mutate(&model)
			model = ComputePoint13ValEFoundation(model)
			if model.DependencyState != Point13ValEStateBlocked {
				t.Fatalf("expected blocked dependency state, got %#v", model)
			}
		})
	}
}

func TestPoint13ValEStateAggregation(t *testing.T) {
	t.Run("any component blocked returns blocked", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.EvidenceIntegrityCheckState = Point13ValEStateBlocked
		if state := EvaluatePoint13ValEState(model); state != Point13ValEStateBlocked {
			t.Fatalf("expected blocked aggregate, got %q", state)
		}
	})

	t.Run("any review required and no blocked returns review required", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.ClosureEvaluatorState = Point13ValEStateReviewRequired
		if state := EvaluatePoint13ValEState(model); state != Point13ValEStateReviewRequired {
			t.Fatalf("expected review_required aggregate, got %q", state)
		}
	})

	t.Run("incomplete returned only when no blocked or review required exists", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.EvidenceIntegrityCheckState = Point13ValEStateIncomplete
		if state := EvaluatePoint13ValEState(model); state != Point13ValEStateIncomplete {
			t.Fatalf("expected incomplete aggregate, got %q", state)
		}
	})

	t.Run("active only when all components active", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if state := EvaluatePoint13ValEState(model); state != Point13ValEStateActive {
			t.Fatalf("expected active aggregate, got %q", state)
		}
	})
}

func TestPoint13ValEAuthorityBoundaryCheckState(t *testing.T) {
	t.Run("valid authority boundary active", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.AuthorityBoundaryCheckState != Point13ValEStateActive {
			t.Fatalf("expected active authority boundary check, got %#v", model)
		}
	})

	t.Run("projection layer attempts mutation blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.HandoffTraceQueryProjection.MutationRequested = true
		model.AuthorityBoundaryCheck.ValDQueryMutationRequested = true
		model = ComputePoint13ValEFoundation(model)
		if model.AuthorityBoundaryCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected mutation attempt to block, got %#v", model)
		}
	})

	t.Run("ai authority flags block", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.AITimelineLineageProjection.DeploymentAuthorized = true
		model.AuthorityBoundaryCheck.ValDAIDeploymentAuthorized = true
		model = ComputePoint13ValEFoundation(model)
		if model.AuthorityBoundaryCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected ai authority flag to block, got %#v", model)
		}
	})
}

func TestPoint13ValENoOverclaimFinalCheckState(t *testing.T) {
	t.Run("safe bounded wording allowed", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.NoOverclaimFinalCheckState != Point13ValEStateActive {
			t.Fatalf("expected active no-overclaim state, got %#v", model)
		}
	})

	t.Run("forbidden wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"production approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected forbidden wording to block, got %#v", model)
		}
	})
}

func TestPoint13ValETimestampIntegrityCheckState(t *testing.T) {
	t.Run("valid timestamp integrity active", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.TimestampIntegrityCheckState != Point13ValEStateActive {
			t.Fatalf("expected active timestamp integrity state, got %#v", model)
		}
	})

	t.Run("local client time cannot create canonical timeline event", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries[0].TimeSource = "client_local"
		model.TimestampIntegrityCheck.TimeSources[0] = "client_local"
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected client canonical time to block, got %#v", model)
		}
	})

	t.Run("backdated events block", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries[4].CanonicalOccurredAt = "2026-05-05T05:59:00Z"
		model.TimestampIntegrityCheck.CanonicalOccurredAts[4] = "2026-05-05T05:59:00Z"
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected backdated event ordering to block, got %#v", model)
		}
	})
}

func TestPoint13ValETwitterIsolationCheckState(t *testing.T) {
	t.Run("valid tenant isolation active", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.TenantIsolationCheckState != Point13ValEStateActive {
			t.Fatalf("expected active tenant isolation state, got %#v", model)
		}
	})

	t.Run("cross tenant timeline access blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.HandoffTraceQueryProjection.FilterRefs = append(model.Dependency.ValD.HandoffTraceQueryProjection.FilterRefs, "artifact_cross-tenant_candidate_001")
		model.TenantIsolationCheck.QueryFilterRefs = append(model.TenantIsolationCheck.QueryFilterRefs, "artifact_cross-tenant_candidate_001")
		model = ComputePoint13ValEFoundation(model)
		if model.TenantIsolationCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected cross-tenant filter ref to block, got %#v", model)
		}
	})
}

func TestPoint13ValEEvidenceIntegrityCheckState(t *testing.T) {
	t.Run("valid evidence integrity active", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.EvidenceIntegrityCheckState != Point13ValEStateActive {
			t.Fatalf("expected active evidence integrity state, got %#v", model)
		}
	})

	t.Run("export read projection cannot recompute hash drift", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.ExportPackageReadProjection.ExportedEvidenceHashes[0] = "evidence_hash_ref_point13_vale_drift_001"
		model.EvidenceIntegrityCheck.ExportedEvidenceHashes[0] = "evidence_hash_ref_point13_vale_drift_001"
		model = ComputePoint13ValEFoundation(model)
		if model.EvidenceIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected recompute drift to block, got %#v", model)
		}
	})

	t.Run("missing lineage incomplete", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.EvidenceIntegrityCheck.LineageComplete = false
		model = ComputePoint13ValEFoundation(model)
		if model.EvidenceIntegrityCheckState != Point13ValEStateIncomplete {
			t.Fatalf("expected missing lineage to be incomplete, got %#v", model)
		}
	})
}

func TestPoint13ValEClosureEvaluatorState(t *testing.T) {
	t.Run("valid closure evaluator active", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.ClosureEvaluatorState != Point13ValEStateActive {
			t.Fatalf("expected active closure evaluator, got %#v", model)
		}
	})

	t.Run("review required reviewer prevents final active evaluator", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.ClosureEvaluator.ReviewerResult = point12ValEReviewerResultReviewRequired
		model = ComputePoint13ValEFoundation(model)
		if model.ClosureEvaluatorState != Point13ValEStateReviewRequired {
			t.Fatalf("expected review required closure evaluator, got %#v", model)
		}
	})
}

func TestPoint13ValEFinalPassGate(t *testing.T) {
	t.Run("pass only when all checks pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.CurrentState != Point13ValEStatePassConfirmed || !model.Point13PassAllowed || model.Point13PassToken != point13ValEPoint13PassToken {
			t.Fatalf("expected pass confirmed final state, got %#v", model)
		}
	})

	t.Run("blocked component clears point13 pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"production approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected blocked current state, got %#v", model)
		}
		if model.Point13PassAllowed || model.Point13PassToken != "" {
			t.Fatalf("expected point13 pass to be cleared, got %#v", model)
		}
	})
}
