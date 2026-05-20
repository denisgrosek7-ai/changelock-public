package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

var (
	point13ValCActiveFoundationBaselineJSON []byte
	point13ValCActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint13ValCFoundation(model Point13ValCFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint13ValCFoundation(payload []byte) Point13ValCFoundation {
	var clone Point13ValCFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func uncachedActivePoint13ValCFoundation() Point13ValCFoundation {
	return ComputePoint13ValCFoundation(Point13ValCFoundationModel())
}

func activePoint13ValCFoundation() Point13ValCFoundation {
	point13ValCActiveFoundationBaselineOnce.Do(func() {
		point13ValCActiveFoundationBaselineJSON = mustMarshalPoint13ValCFoundation(uncachedActivePoint13ValCFoundation())
	})
	return clonePoint13ValCFoundation(point13ValCActiveFoundationBaselineJSON)
}

func point13ValCRecomputeExportManifestHash(model *Point13ValCFoundation) {
	model.CustomerEvidenceExportPackage.ExportManifestHash = point13ValCComputedExportManifestHash(model.CustomerEvidenceExportPackage)
}

func point13ValCRecomputeChecklistBindingHash(model *Point13ValCFoundation) {
	model.OperationalHandoffChecklist.ChecklistBindingHash = point13ValCComputedChecklistBindingHash(model.OperationalHandoffChecklist)
}

func TestPoint13ValCFoundationFixtureIsolation(t *testing.T) {
	t.Run("raw production path still computes", func(t *testing.T) {
		model := uncachedActivePoint13ValCFoundation()
		payload := string(mustMarshalPoint13ValCFoundation(model))
		if model.CurrentState != Point13ValCStateActive {
			t.Fatalf("expected raw production path to compute active foundation, got %#v", model)
		}
		if strings.Contains(payload, point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected no point_13_pass token in active ValC payload, got %s", payload)
		}
	})

	t.Run("cached fixture mutation does not contaminate next clone", func(t *testing.T) {
		mutated := activePoint13ValCFoundation()
		mutated.Dependency.ValBCurrentState = Point13ValBStateBlocked
		mutated.CustomerEvidenceExportPackage.ExportedEvidenceRefs = []string{"artifact_cross-tenant_candidate_001"}
		mutated.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs = []string{"evidence_hash_cross-tenant_candidate_001"}
		point13ValCRecomputeExportManifestHash(&mutated)

		fresh := activePoint13ValCFoundation()
		if fresh.Dependency.ValBCurrentState != Point13ValBStateActive {
			t.Fatalf("expected ValB dependency state to remain active on fresh clone, got %#v", fresh.Dependency)
		}
		if !point12Val0ExactStringSetMatch(
			fresh.CustomerEvidenceExportPackage.ExportedEvidenceRefs,
			fresh.Dependency.ValB.Dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs,
		) {
			t.Fatalf("expected fresh export refs to remain canonical, got %#v", fresh.CustomerEvidenceExportPackage)
		}
	})
}

func TestPoint13ValCDependencyState(t *testing.T) {
	t.Run("valid valb dependency active", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if model.DependencyState != Point13ValCStateActive || model.CurrentState != Point13ValCStateActive {
			t.Fatalf("expected active dependency and foundation, got %#v", model)
		}
	})

	testCases := []struct {
		name          string
		mutate        func(*Point13ValCFoundation)
		expectedState string
	}{
		{
			name: "missing valb dependency blocks",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.ValBCurrentState = ""
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "valb blocked blocks valc",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.ValBCurrentState = Point13ValBStateBlocked
				model.Dependency.ValB.CurrentState = Point13ValBStateBlocked
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "stale valb review required summary blocks valc",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.ValBCurrentState = Point13ValBStateReviewRequired
				model.Dependency.ValB.CurrentState = Point13ValBStateReviewRequired
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "stale valb incomplete summary blocks valc",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.ValBCurrentState = Point13ValBStateIncomplete
				model.Dependency.ValB.CurrentState = Point13ValBStateIncomplete
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "valb point13 pass appearance blocks",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.ValBPoint13PassSeen = true
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "local valc readiness cannot override valb failure",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.ValBCurrentState = Point13ValBStateBlocked
				model.Dependency.ValB.CurrentState = Point13ValBStateBlocked
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "stale inherited point12 review requirement through valb blocks",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.InheritedPoint12CurrentState = Point12ValEStateReviewRequired
				model.Dependency.ValB.Dependency.InheritedPoint12CurrentState = Point12ValEStateReviewRequired
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "inherited point12 binding mismatch through valb blocks",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.InheritedPoint12CurrentState = Point12ValEStateReviewRequired
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "padded nested valb state blocks raw-exact dependency binding",
			mutate: func(model *Point13ValCFoundation) {
				model.Dependency.ValB.CurrentState = " " + Point13ValBStateActive + " "
			},
			expectedState: Point13ValCStateBlocked,
		},
		{
			name: "tab newline inherited point12 reviewer blocks raw-exact dependency binding",
			mutate: func(model *Point13ValCFoundation) {
				retagged := model.Dependency.InheritedPoint12ReviewerResult + "\n"
				model.Dependency.InheritedPoint12ReviewerResult = retagged
				model.Dependency.ValB.Dependency.InheritedPoint12ReviewerResult = retagged
			},
			expectedState: Point13ValCStateBlocked,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model)
			model = ComputePoint13ValCFoundation(model)
			if model.DependencyState != tc.expectedState || model.CurrentState != tc.expectedState {
				t.Fatalf("expected dependency/current state %q, got %#v", tc.expectedState, model)
			}
		})
	}

	t.Run("padded valb point identity blocks exact dependency identity", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.Dependency.ValBPointID = " " + point13Val0PointID + " "
		state, reasons := point13ValCDependencyStateAndReasons(model.Dependency)
		if state != Point13ValCStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValCFoundation(model)
		if model.DependencyState != Point13ValCStateBlocked || model.CurrentState != Point13ValCStateBlocked {
			t.Fatalf("expected padded ValB point identity to block foundation, got %#v", model)
		}
	})

	t.Run("padded valb wave identity blocks exact dependency identity", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.Dependency.ValBWaveID = " " + point13ValBWaveID + " "
		state, reasons := point13ValCDependencyStateAndReasons(model.Dependency)
		if state != Point13ValCStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValCFoundation(model)
		if model.DependencyState != Point13ValCStateBlocked || model.CurrentState != Point13ValCStateBlocked {
			t.Fatalf("expected padded ValB wave identity to block foundation, got %#v", model)
		}
	})

	t.Run("padded nested valb state reports exact binding mismatch", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.Dependency.ValB.CurrentState = " " + Point13ValBStateActive + " "
		state, reasons := point13ValCDependencyStateAndReasons(model.Dependency)
		if state != Point13ValCStateBlocked || !point13Val0StringSliceContains(reasons, "valb_recomputed_snapshot_mismatch") {
			t.Fatalf("expected exact ValB recomputed snapshot mismatch, got state %q reasons %#v", state, reasons)
		}
	})

	t.Run("retagged inherited point12 reviewer reports exact identity invalid", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		retagged := model.Dependency.InheritedPoint12ReviewerResult + "\n"
		model.Dependency.InheritedPoint12ReviewerResult = retagged
		model.Dependency.ValB.Dependency.InheritedPoint12ReviewerResult = retagged
		state, reasons := point13ValCDependencyStateAndReasons(model.Dependency)
		if state != Point13ValCStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
	})

	t.Run("stale embedded valb vala val0 point12 profile mutation blocks recompute", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.Dependency.ValB.Dependency.ValA.Dependency.Val0.Dependency.Point12.Dependency.Val0.Manifest.ProfileContext.CurrentProfileHash = ""
		state, reasons := point13ValCDependencyStateAndReasons(model.Dependency)
		if state != Point13ValCStateBlocked || !point13Val0StringSliceContains(reasons, "valb_recomputed_snapshot_mismatch") {
			t.Fatalf("expected exact ValB recomputed snapshot mismatch, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValCFoundation(model)
		if model.DependencyState != Point13ValCStateBlocked || model.CurrentState != Point13ValCStateBlocked {
			t.Fatalf("expected stale embedded ValB profile mutation to block foundation, got %#v", model)
		}
	})
}

func TestPoint13ValCStateAggregation(t *testing.T) {
	reviewCases := []struct {
		name   string
		mutate func(*Point13ValCFoundation)
	}{
		{name: "customer evidence export package review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackageState = Point13ValCStateReviewRequired
		}},
		{name: "redaction safe disclosure review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureState = Point13ValCStateReviewRequired
		}},
		{name: "operational handoff checklist review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklistState = Point13ValCStateReviewRequired
		}},
		{name: "customer acceptance trace review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.CustomerAcceptanceTraceState = Point13ValCStateReviewRequired
		}},
		{name: "support offboarding handoff review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.SupportOffboardingHandoffState = Point13ValCStateReviewRequired
		}},
		{name: "ai evidence export lineage review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.AIEvidenceExportLineageState = Point13ValCStateReviewRequired
		}},
		{name: "no overclaim review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.NoOverclaimState = Point13ValCStateReviewRequired
		}},
		{name: "dependency review required prevents active", mutate: func(model *Point13ValCFoundation) {
			model.DependencyState = Point13ValCStateReviewRequired
		}},
	}

	for _, tc := range reviewCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model)
			if got := EvaluatePoint13ValCState(model); got != Point13ValCStateReviewRequired {
				t.Fatalf("expected review_required aggregation, got %q for %#v", got, model)
			}
		})
	}

	t.Run("any component blocked returns blocked", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.SupportOffboardingHandoffState = Point13ValCStateBlocked
		model.NoOverclaimState = Point13ValCStateReviewRequired
		if got := EvaluatePoint13ValCState(model); got != Point13ValCStateBlocked {
			t.Fatalf("expected blocked aggregation, got %q for %#v", got, model)
		}
	})

	t.Run("incomplete returned only when no blocked or review required exists", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.CustomerAcceptanceTraceState = Point13ValCStateIncomplete
		if got := EvaluatePoint13ValCState(model); got != Point13ValCStateIncomplete {
			t.Fatalf("expected incomplete aggregation, got %q for %#v", got, model)
		}
	})

	t.Run("active only when all components active", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if got := EvaluatePoint13ValCState(model); got != Point13ValCStateActive {
			t.Fatalf("expected active aggregation, got %q for %#v", got, model)
		}
	})
}

func TestPoint13ValCCustomerEvidenceExportPackageState(t *testing.T) {
	t.Run("valid export package active", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if model.CustomerEvidenceExportPackageState != Point13ValCStateActive {
			t.Fatalf("expected active export package, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValCFoundation)
	}{
		{name: "missing operation ledger ref blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.OperationLedgerRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "missing customer review trace ref blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.CustomerReviewTraceRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "missing support trace ref blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.SupportTraceRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "missing exit evidence packet ref blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExitEvidencePacketRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "missing export owner blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportOwnerRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "missing customer owner blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.CustomerOwnerRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "missing audit event blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.AuditEventRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "missing retention class blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.RetentionClassRef = ""
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "evidence hash mismatch blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs[0] = "evidence_hash_point13_valc_other_candidate_001"
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "cross tenant evidence blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportedEvidenceRefs = []string{"artifact_cross_tenant_candidate_001"}
			model.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs = []string{"evidence_hash_cross_tenant_candidate_001"}
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "cross-tenant evidence blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportedEvidenceRefs = []string{"artifact_cross-tenant_candidate_001"}
			model.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs = []string{"evidence_hash_cross-tenant_candidate_001"}
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "cross tenant spaced evidence blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportedEvidenceRefs = []string{"artifact_cross tenant_candidate_001"}
			model.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs = []string{"evidence_hash_cross tenant_candidate_001"}
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "export read only false blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportIsReadOnly = false
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "export operational only false blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportIsOperationalEvidenceOnly = false
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "export can create pass blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportCannotCreatePass = false
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "export can approve production blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportCannotApproveProduction = false
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "export can certify blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportCannotCertify = false
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "export can mutate canonical evidence blocks", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportCannotMutateCanonicalEvidence = false
			point13ValCRecomputeExportManifestHash(model)
		}},
		{name: "padded export manifest hash blocks raw-exact binding", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportManifestHash = " " + model.CustomerEvidenceExportPackage.ExportManifestHash + " "
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model)
			model = ComputePoint13ValCFoundation(model)
			if model.CustomerEvidenceExportPackageState != Point13ValCStateBlocked {
				t.Fatalf("expected export package mutation to block, got %#v", model)
			}
		})
	}

	t.Run("recomputed export hash after evidence hash tenant authority mutation still blocks", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.CustomerEvidenceExportPackage.TenantScope = "tenant_scope_point13_valc_wrong_001"
		model.CustomerEvidenceExportPackage.ExportedEvidenceRefs = []string{"artifact_cross-tenant_candidate_001"}
		model.CustomerEvidenceExportPackage.ExportedEvidenceHashRefs = []string{"evidence_hash_cross-tenant_candidate_001"}
		model.CustomerEvidenceExportPackage.ExportCannotApproveProduction = false
		point13ValCRecomputeExportManifestHash(&model)
		model = ComputePoint13ValCFoundation(model)
		if model.CurrentState == Point13ValCStateActive {
			t.Fatalf("expected recomputed export hash not to hide drift, got %#v", model)
		}
	})
}

func TestPoint13ValCRedactionSafeDisclosureState(t *testing.T) {
	t.Run("valid redaction boundary active", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if model.RedactionSafeDisclosureState != Point13ValCStateActive {
			t.Fatalf("expected active redaction boundary, got %#v", model)
		}
	})

	testCases := []struct {
		name          string
		mutate        func(*Point13ValCFoundation)
		expectedState string
	}{
		{name: "missing redaction manifest blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionManifestRef = ""
		}, expectedState: Point13ValCStateBlocked},
		{name: "missing approver blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionApproverRef = ""
		}, expectedState: Point13ValCStateBlocked},
		{name: "padded approver ref blocks raw exact", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionApproverRef = " " + model.RedactionSafeDisclosureBoundary.RedactionApproverRef
		}, expectedState: Point13ValCStateBlocked},
		{name: "missing audit event blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionAuditEventRef = ""
		}, expectedState: Point13ValCStateBlocked},
		{name: "decisive evidence removed blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.DecisiveEvidenceRemoved = true
		}, expectedState: Point13ValCStateBlocked},
		{name: "redaction affects decision requires review", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionAffectsDecision = true
		}, expectedState: Point13ValCStateReviewRequired},
		{name: "redaction affects replay requires review", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionAffectsReplay = true
		}, expectedState: Point13ValCStateReviewRequired},
		{name: "redaction strengthens claim blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionCannotStrengthenClaim = false
		}, expectedState: Point13ValCStateBlocked},
		{name: "redaction hides decisive missing evidence blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.RedactionCannotHideDecisiveMissingProof = false
		}, expectedState: Point13ValCStateBlocked},
		{name: "surviving customer claims overlapping disallowed claims blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.SurvivingCustomerClaims = []string{"production approved"}
		}, expectedState: Point13ValCStateBlocked},
		{name: "normalized surviving customer claim variant overlapping disallowed claim blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.DisallowedCustomerClaims = []string{"customer-ready"}
			model.RedactionSafeDisclosureBoundary.SurvivingCustomerClaims = []string{"customer ready"}
		}, expectedState: Point13ValCStateBlocked},
		{name: "normalized duplicate denylist variants do not block safe disclosure", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.DisallowedCustomerClaims = []string{"ai-approved", "ai approved"}
		}, expectedState: Point13ValCStateActive},
		{name: "forbidden minimum safe statement blocks", mutate: func(model *Point13ValCFoundation) {
			model.RedactionSafeDisclosureBoundary.MinimumSafeStatement = "production approved"
		}, expectedState: Point13ValCStateBlocked},
		{name: "redaction cannot convert incomplete export to active", mutate: func(model *Point13ValCFoundation) {
			model.CustomerEvidenceExportPackage.ExportIsReadOnly = false
			point13ValCRecomputeExportManifestHash(model)
		}, expectedState: Point13ValCStateBlocked},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model)
			model = ComputePoint13ValCFoundation(model)
			if model.RedactionSafeDisclosureState != tc.expectedState {
				t.Fatalf("expected redaction state %q, got %#v", tc.expectedState, model)
			}
		})
	}
}

func TestPoint13ValCOperationalHandoffChecklistState(t *testing.T) {
	t.Run("valid handoff checklist active", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if model.OperationalHandoffChecklistState != Point13ValCStateActive {
			t.Fatalf("expected active handoff checklist, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValCFoundation)
	}{
		{name: "missing handoff owner blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.HandoffOwnerRef = ""
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "missing customer owner blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.CustomerOwnerRef = ""
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "missing support owner blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.SupportOwnerRef = ""
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "missing checklist items blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.ChecklistItems = nil
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "missing required acknowledgement refs blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.RequiredAckRefs = nil
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "missing audit events blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.AuditEventRefs = nil
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "production approval flag blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.HandoffCannotApproveProduction = false
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "deployment authorization flag blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.HandoffCannotAuthorizeDeployment = false
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "pass creation flag blocks", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.HandoffCannotCreatePass = false
			point13ValCRecomputeChecklistBindingHash(model)
		}},
		{name: "padded checklist binding hash blocks raw-exact binding", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.ChecklistBindingHash = " " + model.OperationalHandoffChecklist.ChecklistBindingHash + " "
		}},
		{name: "checklist mutation blocks when recomputed", mutate: func(model *Point13ValCFoundation) {
			model.OperationalHandoffChecklist.ChecklistItems = []string{
				"handoff_scope_confirmed",
				"retention_class_disclosed",
				"handoff_item_mutated",
			}
			point13ValCRecomputeChecklistBindingHash(model)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model)
			model = ComputePoint13ValCFoundation(model)
			if model.OperationalHandoffChecklistState != Point13ValCStateBlocked {
				t.Fatalf("expected handoff checklist mutation to block, got %#v", model)
			}
		})
	}
}

func TestPoint13ValCCustomerAcceptanceTraceState(t *testing.T) {
	t.Run("valid acceptance trace active", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if model.CustomerAcceptanceTraceState != Point13ValCStateActive {
			t.Fatalf("expected active acceptance trace, got %#v", model)
		}
	})

	testCases := []struct {
		name          string
		mutate        func(*Point13ValCCustomerAcceptanceTrace)
		expectedState string
	}{
		{name: "missing customer owner blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.CustomerOwnerRef = ""
		}, expectedState: Point13ValCStateBlocked},
		{name: "missing export package ref blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.ExportPackageRef = ""
		}, expectedState: Point13ValCStateBlocked},
		{name: "missing handoff checklist ref blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.HandoffChecklistRef = ""
		}, expectedState: Point13ValCStateBlocked},
		{name: "missing audit event blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.AuditEventRefs = nil
		}, expectedState: Point13ValCStateBlocked},
		{name: "multiple customer questions without matching response coverage blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.CustomerQuestions = []string{
				"customer export package review question",
				"customer handoff retention question",
			}
			model.ResponseRefs = []string{"response_point13_valc_only_one_001"}
		}, expectedState: Point13ValCStateBlocked},
		{name: "unresolved rejected items require review", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.RejectedItems = []string{"rejected_operational_item"}
		}, expectedState: Point13ValCStateReviewRequired},
		{name: "acceptance as production approval blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.AcceptanceIsNotProductionApproval = false
		}, expectedState: Point13ValCStateBlocked},
		{name: "acceptance as compliance attestation blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.AcceptanceIsNotComplianceAttest = false
		}, expectedState: Point13ValCStateBlocked},
		{name: "acceptance creates pass blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.AcceptanceCannotCreatePass = false
		}, expectedState: Point13ValCStateBlocked},
		{name: "forbidden acceptance wording blocks", mutate: func(model *Point13ValCCustomerAcceptanceTrace) {
			model.AcceptedLimitations = []string{"production approved"}
		}, expectedState: Point13ValCStateBlocked},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model.CustomerAcceptanceTrace)
			model = ComputePoint13ValCFoundation(model)
			if model.CustomerAcceptanceTraceState != tc.expectedState {
				t.Fatalf("expected acceptance trace state %q, got %#v", tc.expectedState, model)
			}
		})
	}

	t.Run("safe acceptance limitations wording passes", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.CustomerAcceptanceTrace.AcceptedLimitations = []string{"operational handoff checklist acknowledged"}
		model = ComputePoint13ValCFoundation(model)
		if model.CustomerAcceptanceTraceState != Point13ValCStateActive {
			t.Fatalf("expected safe acceptance wording to remain active, got %#v", model)
		}
	})
}

func TestPoint13ValCSupportOffboardingHandoffState(t *testing.T) {
	t.Run("valid support offboarding packet active", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if model.SupportOffboardingHandoffState != Point13ValCStateActive {
			t.Fatalf("expected active support offboarding packet, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValCSupportOffboardingHandoffPacket)
	}{
		{name: "missing support owner blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) { model.SupportOwnerRef = "" }},
		{name: "missing customer owner blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) { model.CustomerOwnerRef = "" }},
		{name: "missing retention owner blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) { model.RetentionOwnerRef = "" }},
		{name: "missing offboarding plan blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) { model.OffboardingPlanRef = "" }},
		{name: "missing disposal path blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) { model.DisposalPathRef = "" }},
		{name: "missing retention class blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) { model.RetentionClassRefs = nil }},
		{name: "missing audit event blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) { model.AuditEventRefs = nil }},
		{name: "canonical mutation flag blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) {
			model.SupportOffboardingCannotMutateCanonical = false
		}},
		{name: "core decision override flag blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) {
			model.SupportOffboardingCannotOverrideDecision = false
		}},
		{name: "production approval flag blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) {
			model.SupportOffboardingCannotApproveProduction = false
		}},
		{name: "indefinite retention without governance event blocks", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) {
			model.IndefiniteRetentionRequested = true
			model.RetentionGovernanceEventRef = ""
		}},
		{name: "unexpected retention class refs block", mutate: func(model *Point13ValCSupportOffboardingHandoffPacket) {
			model.RetentionClassRefs = []string{
				"retention_class_point13_valc_unexpected_evidence_001",
				"retention_class_point13_valc_unexpected_support_001",
			}
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model.SupportOffboardingHandoffPacket)
			model = ComputePoint13ValCFoundation(model)
			if model.SupportOffboardingHandoffState != Point13ValCStateBlocked {
				t.Fatalf("expected support offboarding mutation to block, got %#v", model)
			}
		})
	}
}

func TestPoint13ValCAIEvidenceExportLineageState(t *testing.T) {
	t.Run("all allowed ai output types remain advisory when authority flags false", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValCFoundation()
			model.AIEvidenceExportLineageSummary.AIOutputType = outputType
			model = ComputePoint13ValCFoundation(model)
			if model.AIEvidenceExportLineageState != Point13ValCStateActive {
				t.Fatalf("expected allowed AI output type %q to remain active advisory candidate, got %#v", outputType, model)
			}
		}
	})

	t.Run("all allowed ai output types block on deployment authorized", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValCFoundation()
			model.AIEvidenceExportLineageSummary.AIOutputType = outputType
			model.AIEvidenceExportLineageSummary.DeploymentAuthorized = true
			model = ComputePoint13ValCFoundation(model)
			if model.AIEvidenceExportLineageState != Point13ValCStateBlocked {
				t.Fatalf("expected deployment authority on %q to block, got %#v", outputType, model)
			}
		}
	})

	t.Run("all allowed ai output types block on production readiness claimed", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValCFoundation()
			model.AIEvidenceExportLineageSummary.AIOutputType = outputType
			model.AIEvidenceExportLineageSummary.ProductionReadinessClaimed = true
			model = ComputePoint13ValCFoundation(model)
			if model.AIEvidenceExportLineageState != Point13ValCStateBlocked {
				t.Fatalf("expected production readiness claim on %q to block, got %#v", outputType, model)
			}
		}
	})

	t.Run("all blocked ai taxonomy values rejected", func(t *testing.T) {
		for _, outputType := range point12Val0BlockedAIEvidenceCandidateTypes() {
			model := activePoint13ValCFoundation()
			model.AIEvidenceExportLineageSummary.AIOutputType = outputType
			model = ComputePoint13ValCFoundation(model)
			if model.AIEvidenceExportLineageState != Point13ValCStateBlocked {
				t.Fatalf("expected blocked AI output type %q to block, got %#v", outputType, model)
			}
		}
	})

	testCases := []struct {
		name          string
		mutate        func(*Point13ValCAIEvidenceExportLineageSummary)
		expectedState string
	}{
		{name: "ai approval request is not approval", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.AIOutputType = "AI_APPROVAL_REQUEST"
			model.ApprovalGranted = true
		}, expectedState: Point13ValCStateBlocked},
		{name: "ai patch proposal is not deployment", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.AIOutputType = "AI_PATCH_PROPOSAL"
			model.DeploymentAuthorized = true
		}, expectedState: Point13ValCStateBlocked},
		{name: "ai sandbox result is not production readiness", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.AIOutputType = "AI_SANDBOX_RESULT"
			model.ProductionReadinessClaimed = true
		}, expectedState: Point13ValCStateBlocked},
		{name: "ai summary cannot strengthen export claim", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.AISummaryCannotStrengthenExportClaim = false
		}, expectedState: Point13ValCStateBlocked},
		{name: "ai summary cannot satisfy customer acceptance by itself", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.AISummaryCannotSatisfyAcceptanceByItself = false
		}, expectedState: Point13ValCStateBlocked},
		{name: "external api allowed without governance event blocks", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.ExternalAPIAllowed = true
		}, expectedState: Point13ValCStateBlocked},
		{name: "external api allowed with governance event requires review", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.ExternalAPIAllowed = true
			model.ExternalAPIGovernanceEventRef = "governance_event_point13_valc_001"
		}, expectedState: Point13ValCStateReviewRequired},
		{name: "permission manifest hash mutation blocks", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.PermissionManifestHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}, expectedState: Point13ValCStateBlocked},
		{name: "input evidence hash refs mutation blocks", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.InputEvidenceHashRefs[0] = "evidence_hash_point13_valc_other_candidate_001"
		}, expectedState: Point13ValCStateBlocked},
		{name: "tenant scope mutation blocks", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.TenantScope = "tenant_scope_point13_valc_wrong_001"
		}, expectedState: Point13ValCStateBlocked},
		{name: "model or rule version ref mutation blocks", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.ModelOrRuleVersionRef = "model_version_point13_valc_wrong_001"
		}, expectedState: Point13ValCStateBlocked},
		{name: "audit event ref mutation blocks", mutate: func(model *Point13ValCAIEvidenceExportLineageSummary) {
			model.AuditEventRef = "audit_point13_valc_wrong_001"
		}, expectedState: Point13ValCStateBlocked},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			tc.mutate(&model.AIEvidenceExportLineageSummary)
			model = ComputePoint13ValCFoundation(model)
			if model.AIEvidenceExportLineageState != tc.expectedState {
				t.Fatalf("expected AI export lineage state %q, got %#v", tc.expectedState, model)
			}
		})
	}
}

func TestPoint13ValCNoOverclaimState(t *testing.T) {
	t.Run("safe wording remains allowed", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		if model.NoOverclaimState != Point13ValCStateActive {
			t.Fatalf("expected active no-overclaim state, got %#v", model)
		}
	})

	for _, phrase := range point13Val0ForbiddenClaims() {
		t.Run("forbidden wording blocks "+phrase, func(t *testing.T) {
			model := activePoint13ValCFoundation()
			model.NoOverclaimExportWording.ObservedCustomerExportTexts = []string{phrase}
			model = ComputePoint13ValCFoundation(model)
			if model.NoOverclaimState != Point13ValCStateBlocked {
				t.Fatalf("expected forbidden phrase %q to block, got %#v", phrase, model)
			}
		})
	}

	t.Run("forbidden wording may appear in internal diagnostics when classified blocked", func(t *testing.T) {
		model := activePoint13ValCFoundation()
		model.NoOverclaimExportWording.InternalDiagnosticTexts = []string{"production approved"}
		model.NoOverclaimExportWording.InternalDiagnosticsClassifiedBlocked = true
		model = ComputePoint13ValCFoundation(model)
		if model.NoOverclaimState != Point13ValCStateActive {
			t.Fatalf("expected internal blocked diagnostics to remain active, got %#v", model)
		}
	})
}
