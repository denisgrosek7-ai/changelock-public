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

func point13ValEStringSliceContains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func assertPoint13ValEReason(t *testing.T, reasons []string, want string) {
	t.Helper()
	for _, reason := range reasons {
		if reason == want {
			return
		}
	}
	t.Fatalf("expected exact reason %q, got %#v", want, reasons)
}

func assertPoint13ValENoPass(t *testing.T, model Point13ValEFoundation) {
	t.Helper()
	if model.Point13PassAllowed || model.Point13PassToken != "" {
		t.Fatalf("expected point13 pass to be cleared, got %#v", model)
	}
	if model.PassClosureManifest.Point13PassAllowed || model.PassClosureManifest.Point13PassToken != "" {
		t.Fatalf("expected nested point13 pass manifest token to be cleared, got %#v", model.PassClosureManifest)
	}
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
		{name: "whitespace retagged vald point id blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValDPointID = " " + point13Val0PointID + " "
		}},
		{name: "tab newline retagged inherited point12 reviewer blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.InheritedPoint12ReviewerResult = "\t" + point12ValEReviewerResultPassConfirmed + "\n"
			model.Dependency.ValD.Dependency.InheritedPoint12ReviewerResult = "\t" + point12ValEReviewerResultPassConfirmed + "\n"
		}},
		{name: "padded inherited tenant scope blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.InheritedTenantScope = " " + model.Dependency.InheritedTenantScope + " "
			model.Dependency.ValD.Dependency.InheritedTenantScope = model.Dependency.InheritedTenantScope
		}},
		{name: "stale nested vald export limitations state blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValD.ExportPackageReadProjection.LimitationsVisible = false
			point13ValDRecomputeExportReadHash(&model.Dependency.ValD)
		}},
		{name: "stale nested vald valc valb vala val0 point12 profile mutation blocks", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValD.Dependency.ValC.Dependency.ValB.Dependency.ValA.Dependency.Val0.Dependency.Point12.Dependency.Val0.Manifest.ProfileContext.CurrentProfileHash = ""
		}},
		{name: "nested val0 allowed wording laundering blocks final dependency", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValD.Dependency.ValC.Dependency.ValB.Dependency.ValA.Dependency.Val0.NoOverclaimCustomerWording.AllowedCustomerFacingWording = []string{"production approved"}
		}},
		{name: "nested vala allowed wording laundering blocks final dependency", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValD.Dependency.ValC.Dependency.ValB.Dependency.ValA.NoOverclaimCustomerWording.AllowedCustomerFacingWording = []string{"deployment approved"}
		}},
		{name: "nested valb allowed wording laundering blocks final dependency", mutate: func(model *Point13ValEFoundation) {
			model.Dependency.ValD.Dependency.ValC.Dependency.ValB.NoOverclaimTrace.AllowedSafeWording = []string{"public badge"}
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
			assertPoint13ValENoPass(t, model)
		})
	}
}

func TestPoint13ValEDependencyStateRawExactReasons(t *testing.T) {
	testCases := []struct {
		name       string
		mutate     func(*Point13ValEDependencySnapshot)
		wantReason string
	}{
		{
			name: "whitespace retagged vald point id",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.ValDPointID = " " + point13Val0PointID + " "
			},
			wantReason: "dependency_snapshot_identity_invalid",
		},
		{
			name: "tab newline retagged inherited point12 reviewer",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.InheritedPoint12ReviewerResult = "\t" + point12ValEReviewerResultPassConfirmed + "\n"
				model.ValD.Dependency.InheritedPoint12ReviewerResult = model.InheritedPoint12ReviewerResult
			},
			wantReason: "dependency_snapshot_identity_invalid",
		},
		{
			name: "sibling inherited state raw mismatch",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.InheritedValBCurrentState = " " + model.InheritedValBCurrentState + " "
			},
			wantReason: "dependency_snapshot_binding_mismatch",
		},
		{
			name: "stale nested vald export read recomputation mismatch",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.ValD.ExportPackageReadProjection.LimitationsVisible = false
				point13ValDRecomputeExportReadHash(&model.ValD)
			},
			wantReason: "dependency_snapshot_vald_recomputed_state_mismatch",
		},
		{
			name: "stale nested point12 profile mutation recomputation mismatch",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.ValD.Dependency.ValC.Dependency.ValB.Dependency.ValA.Dependency.Val0.Dependency.Point12.Dependency.Val0.Manifest.ProfileContext.CurrentProfileHash = ""
			},
			wantReason: "dependency_snapshot_vald_recomputed_state_mismatch",
		},
		{
			name: "nested val0 allowed wording laundering recomputation mismatch",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.ValD.Dependency.ValC.Dependency.ValB.Dependency.ValA.Dependency.Val0.NoOverclaimCustomerWording.AllowedCustomerFacingWording = []string{"production approved"}
			},
			wantReason: "dependency_snapshot_vald_recomputed_state_mismatch",
		},
		{
			name: "nested vala allowed wording laundering recomputation mismatch",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.ValD.Dependency.ValC.Dependency.ValB.Dependency.ValA.NoOverclaimCustomerWording.AllowedCustomerFacingWording = []string{"deployment approved"}
			},
			wantReason: "dependency_snapshot_vald_recomputed_state_mismatch",
		},
		{
			name: "nested valb allowed wording laundering recomputation mismatch",
			mutate: func(model *Point13ValEDependencySnapshot) {
				model.ValD.Dependency.ValC.Dependency.ValB.NoOverclaimTrace.AllowedSafeWording = []string{"public badge"}
			},
			wantReason: "dependency_snapshot_vald_recomputed_state_mismatch",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValEFoundation(t)
			tc.mutate(&model.Dependency)
			state, reasons := point13ValEDependencyStateAndReasons(model.Dependency)
			if state != Point13ValEStateBlocked {
				t.Fatalf("expected blocked dependency state, got %q with reasons %#v", state, reasons)
			}
			if !point13ValEStringSliceContains(reasons, tc.wantReason) {
				t.Fatalf("expected reason %q, got %#v", tc.wantReason, reasons)
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

	t.Run("derived aggregate rejects padded active state", func(t *testing.T) {
		state := point13ValEAggregateDerivedStates(" "+Point13ValEStateActive+" ", Point13ValEStateActive)
		if state != Point13ValEStateBlocked {
			t.Fatalf("expected padded derived aggregate state to block, got %q", state)
		}
	})

	t.Run("derived aggregate rejects tab newline active state", func(t *testing.T) {
		state := point13ValEAggregateDerivedStates(Point13ValEStateActive, "\t"+Point13ValEStateActive+"\n")
		if state != Point13ValEStateBlocked {
			t.Fatalf("expected tab/newline derived aggregate state to block, got %q", state)
		}
	})

	t.Run("derived aggregate rejects empty input", func(t *testing.T) {
		if state := point13ValEAggregateDerivedStates(); state != Point13ValEStateBlocked {
			t.Fatalf("expected empty derived aggregate to block, got %q", state)
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

	t.Run("padded current state blocks authority boundary and final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.AuthorityBoundaryCheck.CurrentState = " " + model.AuthorityBoundaryCheck.CurrentState + " "
		model = ComputePoint13ValEFoundation(model)
		if model.AuthorityBoundaryCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected padded authority current state to block, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "authority_boundary:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
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
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("unhyphenated regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"regulator approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected unhyphenated regulator approval wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("long filler production approval wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"production is now fully globally approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected long filler production approval wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("unicode dash regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"regulator\u2011approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected unicode dash regulator approval wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("confusable split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"pr\u0254duction"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected confusable split production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("greek upsilon production wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"prod\u03c5ction approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected greek upsilon production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("small cap u production wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"prod\U00001d1cction approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected small-cap u production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin upsilon production wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"prod\u028action approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin upsilon production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("greek nu approved wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"production appro\u03bded"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected greek nu production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("greek delta approved wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"production approve\u03b4"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected greek delta production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("small cap t official authority wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"official au\U00001d1bhority"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected small-cap t official authority wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin alpha global truth wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"glob\u0251l truth"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin alpha global truth wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin iota official authority wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"off\u0269cial authority"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin iota official authority wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("dental click global truth wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"g\u01c0obal truth"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected dental-click global truth wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("armenian oh official authority wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"\u0585fficial authority"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected armenian-oh official authority wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("greek eta production wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"productio\u03b7 approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected greek eta production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin eng production wording blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExplanationTexts = []string{"productio\u014b approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin eng production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("right leg u split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"prod\uab4e"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"ction approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected right-leg u split production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin upsilon split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"prod\u028a"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"ction approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin upsilon split production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("greek nu split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"production"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"appro\u03bded"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected greek nu split production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("greek delta split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"production"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"approve\u03b4"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected greek delta split production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("small cap t split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"official au"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"\U00001d1bhority"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected small-cap t split official authority wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin alpha split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"glob\u0251l"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"truth"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin alpha split global truth wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin iota split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"off\u0269cial"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"authority"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin iota split official authority wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("dental click split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"g\u01c0obal"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"truth"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected dental-click split global truth wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("armenian oh split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"\u0585fficial"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"authority"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected armenian-oh split official authority wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("armenian vo split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"productio\u0578"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected armenian vo split production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("latin n with long right leg split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"productio\u019e"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected latin n with long right leg split production approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("split forbidden wording across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"deployment"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected split deployment approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("split regulator approval across final surfaces blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.NoOverclaimFinalCheck.ObservedExportTexts = []string{"regulator"}
		model.NoOverclaimFinalCheck.ObservedTimelineTexts = []string{"approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected split regulator approved wording to block final state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})
}

func TestPoint13ValETimestampIntegrityCheckState(t *testing.T) {
	t.Run("valid timestamp integrity active", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		if model.TimestampIntegrityCheckState != Point13ValEStateActive {
			t.Fatalf("expected active timestamp integrity state, got %#v", model)
		}
	})

	t.Run("empty optional client reported ats preserve raw sequence match", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		for _, value := range model.TimestampIntegrityCheck.ClientReportedAts {
			if value != "" {
				t.Fatalf("expected empty optional client timestamps in canonical fixture, got %#v", model.TimestampIntegrityCheck.ClientReportedAts)
			}
		}
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateActive {
			t.Fatalf("expected empty optional client timestamps to preserve timestamp integrity, got %#v", model)
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

	t.Run("future dated timezone offset blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries[6].CanonicalOccurredAt = "2026-05-05T06:30:00-02:00"
		model.TimestampIntegrityCheck.CanonicalOccurredAts[6] = "2026-05-05T06:30:00-02:00"
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected future dated offset event to block, got %#v", model)
		}
	})

	t.Run("padded canonical timestamp blocks even when mirrored", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries[0].CanonicalOccurredAt += " "
		model.TimestampIntegrityCheck.CanonicalOccurredAts[0] += " "
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected padded canonical timestamp retag to block, got %#v", model)
		}
	})

	t.Run("tab newline verified timestamp blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.TimestampIntegrityCheck.VerifiedAt = "\t" + model.TimestampIntegrityCheck.VerifiedAt + "\n"
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected tab newline verified timestamp retag to block, got %#v", model)
		}
	})

	t.Run("client reported timestamp retag blocks even when mirrored", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.CustomerAuditorOperationalTimeline.TimelineEntries[0].ClientReportedAt = "2026-05-05T06:00:00Z\n"
		model.TimestampIntegrityCheck.ClientReportedAts[0] = "2026-05-05T06:00:00Z\n"
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected client reported timestamp retag to block, got %#v", model)
		}
	})

	t.Run("padded inherited tenant scope blocks timestamp integrity", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.TimestampIntegrityCheck.TenantScope = " " + model.TimestampIntegrityCheck.TenantScope + " "
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected raw tenant retag to block timestamp integrity, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "timestamp_integrity:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("tab newline current state blocks timestamp integrity and final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.TimestampIntegrityCheck.CurrentState = "\t" + model.TimestampIntegrityCheck.CurrentState + "\n"
		model = ComputePoint13ValEFoundation(model)
		if model.TimestampIntegrityCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected raw current state retag to block timestamp integrity, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "timestamp_integrity:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
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

	t.Run("tab newline export package ref blocks tenant isolation", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.TenantIsolationCheck.ExportPackageRef = "\t" + model.TenantIsolationCheck.ExportPackageRef + "\n"
		model = ComputePoint13ValEFoundation(model)
		if model.TenantIsolationCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected raw export package retag to block tenant isolation, got %#v", model)
		}
	})

	t.Run("tab newline inherited tenant scope blocks tenant isolation and final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.TenantIsolationCheck.TenantScope = "\t" + model.TenantIsolationCheck.TenantScope + "\n"
		model = ComputePoint13ValEFoundation(model)
		if model.TenantIsolationCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected raw tenant retag to block tenant isolation, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "tenant_isolation:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("padded current state blocks tenant isolation and final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.TenantIsolationCheck.CurrentState = " " + model.TenantIsolationCheck.CurrentState + " "
		model = ComputePoint13ValEFoundation(model)
		if model.TenantIsolationCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected raw current state retag to block tenant isolation, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "tenant_isolation:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
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

	t.Run("padded export manifest hash blocks evidence integrity", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.EvidenceIntegrityCheck.ExportManifestHash = " " + model.EvidenceIntegrityCheck.ExportManifestHash + " "
		model = ComputePoint13ValEFoundation(model)
		if model.EvidenceIntegrityCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected raw manifest hash retag to block evidence integrity, got %#v", model)
		}
	})

	t.Run("padded inherited tenant scope blocks evidence integrity and final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.EvidenceIntegrityCheck.TenantScope = " " + model.EvidenceIntegrityCheck.TenantScope + " "
		model = ComputePoint13ValEFoundation(model)
		if model.EvidenceIntegrityCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected raw tenant retag to block evidence integrity, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "evidence_integrity:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("tab newline current state blocks evidence integrity and final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.EvidenceIntegrityCheck.CurrentState = "\t" + model.EvidenceIntegrityCheck.CurrentState + "\n"
		model = ComputePoint13ValEFoundation(model)
		if model.EvidenceIntegrityCheckState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected raw current state retag to block evidence integrity, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "evidence_integrity:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
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

	t.Run("incomplete lineage propagates incomplete closure evaluator and foundation state", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.EvidenceIntegrityCheck.LineageComplete = false
		model = ComputePoint13ValEFoundation(model)
		if model.ClosureEvaluatorState != Point13ValEStateIncomplete {
			t.Fatalf("expected incomplete closure evaluator, got %#v", model)
		}
		if model.PassClosureManifestState != Point13ValEStateIncomplete {
			t.Fatalf("expected incomplete pass closure manifest, got %#v", model)
		}
		if model.CurrentState != Point13ValEStateIncomplete {
			t.Fatalf("expected incomplete final state, got %#v", model)
		}
	})

	t.Run("padded current state blocks closure evaluator and final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.ClosureEvaluator.CurrentState = " " + model.ClosureEvaluator.CurrentState + " "
		model = ComputePoint13ValEFoundation(model)
		if model.ClosureEvaluatorState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected padded closure evaluator current state to block, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "closure_evaluator:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("tab newline inherited tenant scope blocks closure evaluator", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.ClosureEvaluator.TenantScope = "\t" + model.ClosureEvaluator.TenantScope + "\n"
		if got := EvaluatePoint13ValEClosureEvaluatorState(model.ClosureEvaluator, model); got != Point13ValEStateBlocked {
			t.Fatalf("expected raw tenant retag to block closure evaluator, got %s", got)
		}
	})

	t.Run("retagged inherited tenant scope propagates to closure evaluator and clears final pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		retagged := "\t" + model.Dependency.InheritedTenantScope + "\n"
		model.Dependency.InheritedTenantScope = retagged
		model.Dependency.ValD.Dependency.InheritedTenantScope = retagged
		model = ComputePoint13ValEFoundation(model)
		if model.ClosureEvaluatorState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected raw inherited tenant retag to block aggregate closure evaluator, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "closure_evaluator:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})
}

func TestPoint13ValEPassClosureManifestRawExactBinding(t *testing.T) {
	tests := []struct {
		name       string
		mutate     func(*Point13PassClosureManifest)
		wantReason string
	}{
		{name: "padded point id blocks", mutate: func(model *Point13PassClosureManifest) {
			model.PointID = " " + model.PointID + " "
		}, wantReason: "pass_closure_manifest_identity_invalid"},
		{name: "tab newline wave id blocks", mutate: func(model *Point13PassClosureManifest) {
			model.WaveID = "\t" + model.WaveID + "\n"
		}, wantReason: "pass_closure_manifest_identity_invalid"},
		{name: "padded final point13 token blocks", mutate: func(model *Point13PassClosureManifest) {
			model.Point13PassToken = point13ValEPoint13PassToken + " "
		}, wantReason: "pass_closure_manifest_token_invalid"},
		{name: "padded generated timestamp blocks raw exact closure manifest", mutate: func(model *Point13PassClosureManifest) {
			model.GeneratedAt = " " + model.GeneratedAt + " "
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "non UTC offset generated timestamp blocks raw exact closure manifest", mutate: func(model *Point13PassClosureManifest) {
			model.GeneratedAt = "2026-05-05T08:05:00+01:00"
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "padded manifest binding ref blocks", mutate: func(model *Point13PassClosureManifest) {
			model.ValDTimelineRef = " " + model.ValDTimelineRef + " "
		}, wantReason: "pass_closure_manifest_binding_mismatch"},
		{name: "padded current state blocks raw exact pass closure manifest", mutate: func(model *Point13PassClosureManifest) {
			model.CurrentState = " " + model.CurrentState + " "
		}, wantReason: "pass_closure_manifest_identity_invalid"},
		{name: "fake prefix-shaped gate refs cannot replace canonical run evidence", mutate: func(model *Point13PassClosureManifest) {
			model.CommandsRun = []string{"command_run_point13_vale_fake_001", "command_run_point13_vale_fake_002", "command_run_point13_vale_fake_003"}
			model.TestsRun = []string{"test_run_point13_vale_fake_001", "test_run_point13_vale_fake_002", "test_run_point13_vale_fake_003"}
			model.GrepsRun = []string{"grep_run_point13_vale_fake_001", "grep_run_point13_vale_fake_002", "grep_run_point13_vale_fake_003", "grep_run_point13_vale_fake_004"}
			model.NegativeFixturesRun = []string{"negative_fixture_point13_vale_fake_001", "negative_fixture_point13_vale_fake_002", "negative_fixture_point13_vale_fake_003"}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "duplicate command run cannot satisfy missing canonical command run", mutate: func(model *Point13PassClosureManifest) {
			expected := point13ValECommandsRun()
			model.CommandsRun = []string{expected[0], expected[0], expected[2]}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "duplicate test run cannot satisfy missing canonical test run", mutate: func(model *Point13PassClosureManifest) {
			expected := point13ValETestsRun()
			model.TestsRun = []string{expected[0], expected[0], expected[2]}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "duplicate grep run cannot satisfy missing canonical grep run", mutate: func(model *Point13PassClosureManifest) {
			expected := point13ValEGrepsRun()
			model.GrepsRun = []string{expected[0], expected[0], expected[2], expected[3]}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "duplicate negative fixture cannot satisfy missing canonical negative fixture", mutate: func(model *Point13PassClosureManifest) {
			expected := point13ValENegativeFixturesRun()
			model.NegativeFixturesRun = []string{expected[0], expected[0], expected[2]}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "zero width fake command run blocks raw exact gate refs", mutate: func(model *Point13PassClosureManifest) {
			model.CommandsRun = []string{"command_run_point13_vale_\u200dfake_001", "command_run_point13_vale_go_test_formal_001", "command_run_point13_vale_go_test_all_001"}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "zero width fake test run blocks raw exact gate refs", mutate: func(model *Point13PassClosureManifest) {
			model.TestsRun = []string{"test_run_point13_vale_\u200dfake_001", "test_run_point13_vale_point13_regressions_001", "test_run_point13_vale_go_test_all_001"}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "zero width fake grep run blocks raw exact gate refs", mutate: func(model *Point13PassClosureManifest) {
			model.GrepsRun = []string{"grep_run_point13_vale_\u200dfake_001", "grep_run_point13_vale_ai_authority_001", "grep_run_point13_vale_forbidden_wording_001", "grep_run_point13_vale_mutation_flags_001"}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "zero width fake negative fixture blocks raw exact gate refs", mutate: func(model *Point13PassClosureManifest) {
			model.NegativeFixturesRun = []string{"negative_fixture_point13_vale_\u200dfake_001", "negative_fixture_point13_vale_authority_boundary_001", "negative_fixture_point13_vale_timestamp_integrity_001"}
		}, wantReason: "pass_closure_manifest_required_fields_invalid"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValEFoundation(t)
			tc.mutate(&model.PassClosureManifest)
			got, reasons := point13ValEPassClosureManifestStateAndReasons(model.PassClosureManifest, model, true)
			if got != Point13ValEStateBlocked {
				t.Fatalf("expected blocked pass closure manifest, got %s reasons=%v model=%#v", got, reasons, model.PassClosureManifest)
			}
			found := false
			for _, reason := range reasons {
				if reason == tc.wantReason {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("expected exact reason %q, got %#v", tc.wantReason, reasons)
			}
		})
	}

	t.Run("polluted dependency and manifest ValD timeline ref still blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		polluted := " " + model.Dependency.ValD.CustomerAuditorOperationalTimeline.TimelineID + " "
		model.Dependency.ValD.CustomerAuditorOperationalTimeline.TimelineID = polluted
		model.PassClosureManifest.ValDTimelineRef = polluted
		got, reasons := point13ValEPassClosureManifestStateAndReasons(model.PassClosureManifest, model, true)
		if got != Point13ValEStateBlocked {
			t.Fatalf("expected polluted dependency-bound ValD timeline ref to block, got %s reasons=%v", got, reasons)
		}
		assertPoint13ValEReason(t, reasons, "pass_closure_manifest_required_fields_invalid")
	})

	t.Run("polluted dependency and manifest ValC export ref still blocks", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		polluted := model.Dependency.ValD.Dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID + "\n"
		model.Dependency.ValD.Dependency.ValC.CustomerEvidenceExportPackage.ExportPackageID = polluted
		model.PassClosureManifest.ValCExportPackageRef = polluted
		got, reasons := point13ValEPassClosureManifestStateAndReasons(model.PassClosureManifest, model, true)
		if got != Point13ValEStateBlocked {
			t.Fatalf("expected polluted dependency-bound ValC export ref to block, got %s reasons=%v", got, reasons)
		}
		assertPoint13ValEReason(t, reasons, "pass_closure_manifest_required_fields_invalid")
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
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("blocked manifest clears nested point13 pass token", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.PassClosureManifest.GeneratedAt = " " + model.PassClosureManifest.GeneratedAt + " "
		model = ComputePoint13ValEFoundation(model)
		if model.PassClosureManifestState != Point13ValEStateBlocked || model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected blocked manifest and current state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "pass_closure_manifest:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("stale dependency ValD no-overclaim text clears point13 pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.NoOverclaimProjectionWording.ObservedTimelineTexts = []string{"production approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected dependency-derived no-overclaim block, got %#v", model)
		}
		if model.ClosureEvaluator.NoOverclaimResult != Point13ValEStateBlocked ||
			model.PassClosureManifest.NoOverclaimResult != Point13ValEStateBlocked {
			t.Fatalf("expected blocked no-overclaim propagated to closure and manifest, got %#v", model)
		}
		if model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected blocked current state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("stale dependency split no-overclaim text clears point13 pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.NoOverclaimProjectionWording.ObservedTimelineTexts = []string{"deployment"}
		model.Dependency.ValD.NoOverclaimProjectionWording.ObservedQueryTexts = []string{"approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected dependency-derived split no-overclaim block, got %#v", model)
		}
		if model.ClosureEvaluator.NoOverclaimResult != Point13ValEStateBlocked ||
			model.PassClosureManifest.NoOverclaimResult != Point13ValEStateBlocked {
			t.Fatalf("expected split blocked no-overclaim propagated to closure and manifest, got %#v", model)
		}
		if model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected split blocked current state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("stale dependency ValC no-overclaim text clears point13 pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.Dependency.ValC.NoOverclaimExportWording.ObservedCustomerExportTexts = []string{"deployment approved"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected dependency-derived no-overclaim block, got %#v", model)
		}
		if model.ClosureEvaluator.NoOverclaimResult != Point13ValEStateBlocked ||
			model.PassClosureManifest.NoOverclaimResult != Point13ValEStateBlocked {
			t.Fatalf("expected blocked no-overclaim propagated to closure and manifest, got %#v", model)
		}
		if model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected blocked current state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})

	t.Run("inherited point10 readiness overclaim clears point13 pass", func(t *testing.T) {
		model := activePoint13ValEFoundation(t)
		model.Dependency.ValD.Dependency.ValC.NoOverclaimExportWording.ObservedCustomerExportTexts = []string{"marketplace production ready"}
		model = ComputePoint13ValEFoundation(model)
		if model.NoOverclaimFinalCheckState != Point13ValEStateBlocked {
			t.Fatalf("expected inherited readiness overclaim block, got %#v", model)
		}
		if model.ClosureEvaluator.NoOverclaimResult != Point13ValEStateBlocked ||
			model.PassClosureManifest.NoOverclaimResult != Point13ValEStateBlocked {
			t.Fatalf("expected blocked no-overclaim propagated to closure and manifest, got %#v", model)
		}
		if model.CurrentState != Point13ValEStateBlocked {
			t.Fatalf("expected blocked current state, got %#v", model)
		}
		assertPoint13ValEReason(t, model.BlockingReasons, "no_overclaim:"+Point13ValEStateBlocked)
		assertPoint13ValENoPass(t, model)
	})
}
