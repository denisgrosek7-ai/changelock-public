package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

func mustMarshalPoint13ValDFoundation(t *testing.T, model Point13ValDFoundation) string {
	t.Helper()
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal point13 vald foundation: %v", err)
	}
	return string(payload)
}

func clonePoint13ValDFoundation(t *testing.T, model Point13ValDFoundation) Point13ValDFoundation {
	t.Helper()
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal point13 vald clone: %v", err)
	}
	var cloned Point13ValDFoundation
	if err := json.Unmarshal(payload, &cloned); err != nil {
		t.Fatalf("unmarshal point13 vald clone: %v", err)
	}
	return cloned
}

var (
	point13ValDActiveFoundationBaseline     Point13ValDFoundation
	point13ValDActiveFoundationBaselineOnce sync.Once
)

func uncachedActivePoint13ValDFoundation() Point13ValDFoundation {
	return ComputePoint13ValDFoundation(Point13ValDFoundationModel())
}

func activePoint13ValDFoundation(t *testing.T) Point13ValDFoundation {
	t.Helper()
	point13ValDActiveFoundationBaselineOnce.Do(func() {
		point13ValDActiveFoundationBaseline = uncachedActivePoint13ValDFoundation()
	})
	return clonePoint13ValDFoundation(t, point13ValDActiveFoundationBaseline)
}

func point13ValDRecomputeTimelineHash(model *Point13ValDFoundation) {
	model.CustomerAuditorOperationalTimeline.TimelineHash = point13ValDComputedTimelineHash(model.CustomerAuditorOperationalTimeline)
}

func point13ValDRecomputeQueryHash(model *Point13ValDFoundation) {
	model.HandoffTraceQueryProjection.ProjectionHash = point13ValDComputedQueryHash(model.HandoffTraceQueryProjection)
}

func point13ValDRecomputeExportReadHash(model *Point13ValDFoundation) {
	model.ExportPackageReadProjection.ProjectionHash = point13ValDComputedExportReadHash(model.ExportPackageReadProjection)
}

func point13ValDRecomputeExplanationHash(model *Point13ValDFoundation) {
	model.CustomerAuditorExplanationProjection.ProjectionHash = point13ValDComputedExplanationHash(model.CustomerAuditorExplanationProjection)
}

func TestPoint13ValDFoundationFixtureIsolation(t *testing.T) {
	t.Run("raw production path still computes", func(t *testing.T) {
		model := uncachedActivePoint13ValDFoundation()
		if model.CurrentState != Point13ValDStateActive {
			t.Fatalf("expected active raw ValD foundation, got %#v", model)
		}
		payload := mustMarshalPoint13ValDFoundation(t, model)
		if strings.Contains(payload, point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected no point_13_pass token in active ValD payload, got %s", payload)
		}
	})

	t.Run("cached fixture mutation does not contaminate next clone", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.CustomerAuditorOperationalTimeline.TimelineID = "timeline_point13_vald_mutated_001"
		model.CustomerAuditorOperationalTimeline.TimelineEntries[0].SourceRef = "export_package_point13_vald_mutated_001"
		next := activePoint13ValDFoundation(t)
		if next.CustomerAuditorOperationalTimeline.TimelineID != "timeline_point13_vald_001" {
			t.Fatalf("expected cached baseline isolation, got %#v", next)
		}
	})
}

func TestPoint13ValDDependencyState(t *testing.T) {
	testCases := []struct {
		name          string
		mutate        func(*Point13ValDDependencySnapshot)
		expectedState string
	}{
		{name: "valid valc dependency active", mutate: func(*Point13ValDDependencySnapshot) {}, expectedState: Point13ValDStateActive},
		{name: "missing valc dependency blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.SnapshotFromComputedOutput = false
		}, expectedState: Point13ValDStateBlocked},
		{name: "valc blocked blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.ValCCurrentState = Point13ValCStateBlocked
			model.ValC.CurrentState = Point13ValCStateBlocked
		}, expectedState: Point13ValDStateBlocked},
		{name: "valc review required prevents active", mutate: func(model *Point13ValDDependencySnapshot) {
			model.ValCCurrentState = Point13ValCStateReviewRequired
			model.ValC.CurrentState = Point13ValCStateReviewRequired
		}, expectedState: Point13ValDStateReviewRequired},
		{name: "valc incomplete prevents active", mutate: func(model *Point13ValDDependencySnapshot) {
			model.ValCCurrentState = Point13ValCStateIncomplete
			model.ValC.CurrentState = Point13ValCStateIncomplete
		}, expectedState: Point13ValDStateIncomplete},
		{name: "point13 pass before vale blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.ValCPoint13PassSeen = true
		}, expectedState: Point13ValDStateBlocked},
		{name: "local vald readiness cannot override valc failure", mutate: func(model *Point13ValDDependencySnapshot) {
			model.ValCCurrentState = Point13ValCStateBlocked
			model.ValCDependencyState = Point13ValCStateActive
			model.ValCCustomerEvidenceExportPackageState = Point13ValCStateActive
		}, expectedState: Point13ValDStateBlocked},
		{name: "inherited point12 mismatch through valc blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.InheritedPoint12CurrentState = Point12ValEStateActive
		}, expectedState: Point13ValDStateBlocked},
		{name: "inherited point12 dependency reviewer closure mismatch blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.InheritedPoint12DependencyState = Point12ValEStateBlocked
			model.InheritedPoint12ReviewerResult = point12ValEReviewerResultReviewRequired
		}, expectedState: Point13ValDStateBlocked},
		{name: "inherited valb state mismatch blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.InheritedValBCurrentState = Point13ValBStateBlocked
		}, expectedState: Point13ValDStateBlocked},
		{name: "inherited vala state mismatch blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.InheritedValACurrentState = Point13ValAStateBlocked
		}, expectedState: Point13ValDStateBlocked},
		{name: "inherited val0 state mismatch blocks", mutate: func(model *Point13ValDDependencySnapshot) {
			model.InheritedVal0CurrentState = Point13Val0StateBlocked
		}, expectedState: Point13ValDStateBlocked},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValDFoundation(t)
			tc.mutate(&model.Dependency)
			model = ComputePoint13ValDFoundation(model)
			if model.DependencyState != tc.expectedState {
				t.Fatalf("expected dependency state %q, got %#v", tc.expectedState, model)
			}
		})
	}
}

func TestPoint13ValDStateAggregation(t *testing.T) {
	t.Run("any component blocked returns blocked", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.ExportPackageReadProjectionState = Point13ValDStateBlocked
		if state := EvaluatePoint13ValDState(model); state != Point13ValDStateBlocked {
			t.Fatalf("expected blocked aggregate, got %q", state)
		}
	})

	t.Run("any review required and no blocked returns review required", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.CustomerAuditorOperationalTimelineState = Point13ValDStateReviewRequired
		if state := EvaluatePoint13ValDState(model); state != Point13ValDStateReviewRequired {
			t.Fatalf("expected review_required aggregate, got %q", state)
		}
	})

	t.Run("incomplete returned only when no blocked or review required exists", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.CustomerAuditorOperationalTimelineState = Point13ValDStateIncomplete
		if state := EvaluatePoint13ValDState(model); state != Point13ValDStateIncomplete {
			t.Fatalf("expected incomplete aggregate, got %q", state)
		}
	})

	t.Run("active only when all components active", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		if state := EvaluatePoint13ValDState(model); state != Point13ValDStateActive {
			t.Fatalf("expected active aggregate, got %q", state)
		}
	})
}

func TestPoint13ValDCustomerAuditorOperationalTimelineState(t *testing.T) {
	t.Run("valid read only timeline active", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		if model.CustomerAuditorOperationalTimelineState != Point13ValDStateActive {
			t.Fatalf("expected active timeline, got %#v", model)
		}
	})

	testCases := []struct {
		name          string
		mutate        func(*Point13ValDFoundation)
		expectedState string
	}{
		{name: "timeline mutation flag blocks", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.TimelineCannotMutateState = false
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateBlocked},
		{name: "missing source ref blocks", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.TimelineEntries[0].SourceRef = ""
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateBlocked},
		{name: "missing audit ref blocks", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.TimelineEntries[0].AuditEventRef = ""
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateBlocked},
		{name: "missing timestamp blocks", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.TimelineEntries[0].CanonicalOccurredAt = ""
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateBlocked},
		{name: "missing source metadata blocks", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.TimelineEntries[0].SourceMetadataRef = ""
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateBlocked},
		{name: "local client time cannot create canonical timeline event", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.TimelineEntries[0].TimeSource = "client_local"
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateBlocked},
		{name: "backdated acceptance before export handoff requires review", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.TimelineEntries[4].CanonicalOccurredAt = "2026-05-05T05:59:00Z"
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateReviewRequired},
		{name: "redacted limitations remain visible", mutate: func(model *Point13ValDFoundation) {
			model.CustomerAuditorOperationalTimeline.RedactionLimitationsVisible = false
			point13ValDRecomputeTimelineHash(model)
		}, expectedState: Point13ValDStateBlocked},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValDFoundation(t)
			tc.mutate(&model)
			model = ComputePoint13ValDFoundation(model)
			if model.CustomerAuditorOperationalTimelineState != tc.expectedState {
				t.Fatalf("expected timeline state %q, got %#v", tc.expectedState, model)
			}
		})
	}
}

func TestPoint13ValDHandoffTraceQueryProjectionState(t *testing.T) {
	t.Run("valid query projection active", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		if model.HandoffTraceQueryProjectionState != Point13ValDStateActive {
			t.Fatalf("expected active query projection, got %#v", model)
		}
	})

	for _, tc := range []struct {
		name   string
		mutate func(*Point13ValDFoundation)
	}{
		{name: "query mutation write flag blocks", mutate: func(model *Point13ValDFoundation) {
			model.HandoffTraceQueryProjection.MutationRequested = true
			point13ValDRecomputeQueryHash(model)
		}},
		{name: "query write flag blocks", mutate: func(model *Point13ValDFoundation) {
			model.HandoffTraceQueryProjection.WriteRequested = true
			point13ValDRecomputeQueryHash(model)
		}},
		{name: "unexpected valid filter ref blocks", mutate: func(model *Point13ValDFoundation) {
			model.HandoffTraceQueryProjection.FilterRefs = append(model.HandoffTraceQueryProjection.FilterRefs, "export_package_point13_vald_unrelated_001")
			point13ValDRecomputeQueryHash(model)
		}},
		{name: "cross tenant filter ref blocks", mutate: func(model *Point13ValDFoundation) {
			model.HandoffTraceQueryProjection.FilterRefs = append(model.HandoffTraceQueryProjection.FilterRefs, "artifact_cross-tenant_candidate_001")
			point13ValDRecomputeQueryHash(model)
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValDFoundation(t)
			tc.mutate(&model)
			model = ComputePoint13ValDFoundation(model)
			if model.HandoffTraceQueryProjectionState != Point13ValDStateBlocked {
				t.Fatalf("expected blocked query projection, got %#v", model)
			}
		})
	}
}

func TestPoint13ValDExportPackageReadProjectionState(t *testing.T) {
	t.Run("valid export read projection active", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		if model.ExportPackageReadProjectionState != Point13ValDStateActive {
			t.Fatalf("expected active export read projection, got %#v", model)
		}
	})

	t.Run("export read projection cannot recompute hash drift", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.ExportPackageReadProjection.ExportManifestHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		point13ValDRecomputeExportReadHash(&model)
		model = ComputePoint13ValDFoundation(model)
		if model.ExportPackageReadProjectionState != Point13ValDStateBlocked {
			t.Fatalf("expected export read projection hash drift to block, got %#v", model)
		}
	})
}

func TestPoint13ValDCustomerAuditorExplanationProjectionState(t *testing.T) {
	t.Run("valid explanation projection active", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		if model.CustomerAuditorExplanationProjectionState != Point13ValDStateActive {
			t.Fatalf("expected active explanation projection, got %#v", model)
		}
	})

	t.Run("customer auditor explanation cannot strengthen claims", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.CustomerAuditorExplanationProjection.ExplanationCannotStrengthenClaims = false
		point13ValDRecomputeExplanationHash(&model)
		model = ComputePoint13ValDFoundation(model)
		if model.CustomerAuditorExplanationProjectionState != Point13ValDStateBlocked {
			t.Fatalf("expected strengthened explanation to block, got %#v", model)
		}
	})

	t.Run("auditor annotation cannot approve production", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.CustomerAuditorExplanationProjection.AuditorAnnotations[0].ApprovesProduction = true
		point13ValDRecomputeExplanationHash(&model)
		model = ComputePoint13ValDFoundation(model)
		if model.CustomerAuditorExplanationProjectionState != Point13ValDStateBlocked {
			t.Fatalf("expected approving auditor annotation to block, got %#v", model)
		}
	})
}

func TestPoint13ValDTimelineAccessBoundaryState(t *testing.T) {
	t.Run("valid timeline access boundary active", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		if model.TimelineAccessBoundaryState != Point13ValDStateActive {
			t.Fatalf("expected active access boundary, got %#v", model)
		}
	})

	t.Run("cross tenant timeline access blocks", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.TimelineAccessBoundary.TenantScope = "tenant_scope_point13_vald_other"
		model = ComputePoint13ValDFoundation(model)
		if model.TimelineAccessBoundaryState != Point13ValDStateBlocked {
			t.Fatalf("expected cross-tenant access to block, got %#v", model)
		}
	})
}

func TestPoint13ValDAITimelineLineageProjectionState(t *testing.T) {
	t.Run("allowed ai output types remain advisory", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValDFoundation(t)
			model.AITimelineLineageProjection.AIOutputType = outputType
			model = ComputePoint13ValDFoundation(model)
			if model.AITimelineLineageProjectionState != Point13ValDStateActive {
				t.Fatalf("expected allowed AI output type %q to remain active advisory, got %#v", outputType, model)
			}
		}
	})

	t.Run("ai production deployment readiness authority flags block", func(t *testing.T) {
		for _, mutate := range []func(*Point13ValDAITimelineLineageProjection){
			func(model *Point13ValDAITimelineLineageProjection) { model.DeploymentAuthorized = true },
			func(model *Point13ValDAITimelineLineageProjection) { model.ProductionReadinessClaimed = true },
			func(model *Point13ValDAITimelineLineageProjection) { model.ApprovalGranted = true },
		} {
			model := activePoint13ValDFoundation(t)
			mutate(&model.AITimelineLineageProjection)
			model = ComputePoint13ValDFoundation(model)
			if model.AITimelineLineageProjectionState != Point13ValDStateBlocked {
				t.Fatalf("expected AI authority flag to block, got %#v", model)
			}
		}
	})
}

func TestPoint13ValDNoOverclaimState(t *testing.T) {
	t.Run("safe bounded wording allowed", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		if model.NoOverclaimState != Point13ValDStateActive {
			t.Fatalf("expected active no-overclaim state, got %#v", model)
		}
	})

	t.Run("forbidden wording blocks", func(t *testing.T) {
		model := activePoint13ValDFoundation(t)
		model.NoOverclaimProjectionWording.ObservedExplanationTexts = []string{"production approved"}
		model = ComputePoint13ValDFoundation(model)
		if model.NoOverclaimState != Point13ValDStateBlocked {
			t.Fatalf("expected forbidden wording to block, got %#v", model)
		}
	})
}
