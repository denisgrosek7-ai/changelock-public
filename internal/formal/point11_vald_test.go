package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

func mustMarshalPoint11ValDDependencyBundle(model Point11ValDDependencyBundle) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint11ValDDependencyBundle(payload []byte) Point11ValDDependencyBundle {
	var clone Point11ValDDependencyBundle
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func mustMarshalPoint11ValDFoundation(model Point11ValDFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint11ValDFoundation(payload []byte) Point11ValDFoundation {
	var clone Point11ValDFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

var (
	point11ValDActiveDependencyBundleBaselineJSON []byte
	point11ValDActiveDependencyBundleBaselineOnce sync.Once
	point11ValDActiveFoundationBaselineJSON       []byte
	point11ValDActiveFoundationBaselineOnce       sync.Once
)

func point11ValDActiveVal0DependencySnapshot() Point11ValDVal0DependencySnapshot {
	val0 := activePoint11Val0Foundation()
	return SnapshotPoint11ValDVal0DependencyFromComputed(val0, Point11ValDVal0ReviewContext{})
}

func point11ValDActiveValADependencySnapshot() Point11ValDValADependencySnapshot {
	valA := activePoint11ValAFoundation()
	return SnapshotPoint11ValDValADependencyFromComputed(valA, Point11ValDValAReviewContext{})
}

func point11ValDActiveValBDependencySnapshot() Point11ValDValBDependencySnapshot {
	valB := activePoint11ValBFoundation()
	return SnapshotPoint11ValDValBDependencyFromComputed(valB, Point11ValDValBReviewContext{})
}

func point11ValDActiveValCDependencySnapshot() Point11ValDValCDependencySnapshot {
	valC := activePoint11ValCFoundation()
	return SnapshotPoint11ValDValCDependencyFromComputed(valC, Point11ValDValCReviewContext{})
}

func uncachedPoint11ValDActiveDependencyBundle() Point11ValDDependencyBundle {
	return Point11ValDDependencyBundle{
		Val0: point11ValDActiveVal0DependencySnapshot(),
		ValA: point11ValDActiveValADependencySnapshot(),
		ValB: point11ValDActiveValBDependencySnapshot(),
		ValC: point11ValDActiveValCDependencySnapshot(),
	}
}

func point11ValDActiveDependencyBundle() Point11ValDDependencyBundle {
	point11ValDActiveDependencyBundleBaselineOnce.Do(func() {
		point11ValDActiveDependencyBundleBaselineJSON = mustMarshalPoint11ValDDependencyBundle(uncachedPoint11ValDActiveDependencyBundle())
	})
	return clonePoint11ValDDependencyBundle(point11ValDActiveDependencyBundleBaselineJSON)
}

func uncachedActivePoint11ValDFoundation() Point11ValDFoundation {
	model := Point11ValDFoundationModel()
	model.Val0Dependency = point11ValDActiveVal0DependencySnapshot()
	model.ValADependency = point11ValDActiveValADependencySnapshot()
	model.ValBDependency = point11ValDActiveValBDependencySnapshot()
	model.ValCDependency = point11ValDActiveValCDependencySnapshot()
	return ComputePoint11ValDFoundation(model)
}

func activePoint11ValDFoundation() Point11ValDFoundation {
	point11ValDActiveFoundationBaselineOnce.Do(func() {
		point11ValDActiveFoundationBaselineJSON = mustMarshalPoint11ValDFoundation(uncachedActivePoint11ValDFoundation())
	})
	return clonePoint11ValDFoundation(point11ValDActiveFoundationBaselineJSON)
}

func TestPoint11ValDDependencyState(t *testing.T) {
	t.Run("happy path val 0 a b c dependencies active", func(t *testing.T) {
		bundle := point11ValDActiveDependencyBundle()
		if got := EvaluatePoint11ValDDependencyState(bundle); got != Point11ValDDependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", bundle)
		}
	})

	disclaimerTests := []struct {
		name   string
		mutate func(*Point11ValDDependencyBundle)
		want   string
	}{
		{
			name: "copied val 0 aggregate projection disclaimer propagates exactly",
			mutate: func(model *Point11ValDDependencyBundle) {
				val0 := activePoint11Val0Foundation()
				val0.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_val0_disclaimer"
				val0.PolicyContract.ProjectionDisclaimer = "projection_only not_canonical_truth component_policy_contract"
				model.Val0 = SnapshotPoint11ValDVal0DependencyFromComputed(val0, Point11ValDVal0ReviewContext{})
				if model.Val0.ProjectionDisclaimer != val0.ProjectionDisclaimer {
					t.Fatalf("expected copied val0 disclaimer, got snapshot=%q val0=%q", model.Val0.ProjectionDisclaimer, val0.ProjectionDisclaimer)
				}
			},
			want: Point11ValDDependencyStateActive,
		},
		{
			name: "copied val a aggregate projection disclaimer propagates exactly",
			mutate: func(model *Point11ValDDependencyBundle) {
				valA := activePoint11ValAFoundation()
				valA.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_vala_disclaimer"
				valA.Registry.ProjectionDisclaimer = "projection_only not_canonical_truth component_registry_disclaimer"
				model.ValA = SnapshotPoint11ValDValADependencyFromComputed(valA, Point11ValDValAReviewContext{})
				if model.ValA.ProjectionDisclaimer != valA.ProjectionDisclaimer {
					t.Fatalf("expected copied vala disclaimer, got snapshot=%q vala=%q", model.ValA.ProjectionDisclaimer, valA.ProjectionDisclaimer)
				}
			},
			want: Point11ValDDependencyStateActive,
		},
		{
			name: "copied val b aggregate projection disclaimer propagates exactly",
			mutate: func(model *Point11ValDDependencyBundle) {
				valB := activePoint11ValBFoundation()
				valB.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_valb_disclaimer"
				valB.ClaimTypeDefinition.ProjectionDisclaimer = "projection_only not_canonical_truth component_claim_type"
				model.ValB = SnapshotPoint11ValDValBDependencyFromComputed(valB, Point11ValDValBReviewContext{})
				if model.ValB.ProjectionDisclaimer != valB.ProjectionDisclaimer {
					t.Fatalf("expected copied valb disclaimer, got snapshot=%q valb=%q", model.ValB.ProjectionDisclaimer, valB.ProjectionDisclaimer)
				}
			},
			want: Point11ValDDependencyStateActive,
		},
		{
			name: "copied val c aggregate projection disclaimer propagates exactly",
			mutate: func(model *Point11ValDDependencyBundle) {
				valC := activePoint11ValCFoundation()
				valC.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_valc_disclaimer"
				valC.EnforcementInput.ProjectionDisclaimer = "projection_only not_canonical_truth component_enforcement_input"
				model.ValC = SnapshotPoint11ValDValCDependencyFromComputed(valC, Point11ValDValCReviewContext{})
				if model.ValC.ProjectionDisclaimer != valC.ProjectionDisclaimer {
					t.Fatalf("expected copied valc disclaimer, got snapshot=%q valc=%q", model.ValC.ProjectionDisclaimer, valC.ProjectionDisclaimer)
				}
			},
			want: Point11ValDDependencyStateActive,
		},
		{
			name:   "malformed val 0 disclaimer blocks",
			mutate: func(model *Point11ValDDependencyBundle) { model.Val0.ProjectionDisclaimer = "canonical_truth" },
			want:   Point11ValDDependencyStateBlocked,
		},
		{
			name:   "malformed val a disclaimer blocks",
			mutate: func(model *Point11ValDDependencyBundle) { model.ValA.ProjectionDisclaimer = "canonical_truth" },
			want:   Point11ValDDependencyStateBlocked,
		},
		{
			name:   "malformed val b disclaimer blocks",
			mutate: func(model *Point11ValDDependencyBundle) { model.ValB.ProjectionDisclaimer = "canonical_truth" },
			want:   Point11ValDDependencyStateBlocked,
		},
		{
			name:   "malformed val c disclaimer blocks",
			mutate: func(model *Point11ValDDependencyBundle) { model.ValC.ProjectionDisclaimer = "canonical_truth" },
			want:   Point11ValDDependencyStateBlocked,
		},
		{
			name: "blocked val 0 component blocks",
			mutate: func(model *Point11ValDDependencyBundle) {
				model.Val0.PolicyContractState = Point11Val0PolicyContractStateBlocked
			},
			want: Point11ValDDependencyStateBlocked,
		},
		{
			name:   "blocked val a component blocks",
			mutate: func(model *Point11ValDDependencyBundle) { model.ValA.RegistryState = Point11ValARegistryStateBlocked },
			want:   Point11ValDDependencyStateBlocked,
		},
		{
			name: "blocked val b component blocks",
			mutate: func(model *Point11ValDDependencyBundle) {
				model.ValB.VerificationState = Point11ValBVerificationStateBlocked
			},
			want: Point11ValDDependencyStateBlocked,
		},
		{
			name:   "blocked val c component blocks",
			mutate: func(model *Point11ValDDependencyBundle) { model.ValC.DashboardState = Point11ValCDashboardStateBlocked },
			want:   Point11ValDDependencyStateBlocked,
		},
		{
			name: "review required upstream blocks final point 11 pass",
			mutate: func(model *Point11ValDDependencyBundle) {
				valC := activePoint11ValCFoundation()
				valC.CurrentState = Point11ValCStateReviewRequired
				valC.DependencyState = Point11ValCDependencyStateReviewRequired
				valC.ReviewPrerequisites = []string{"external_review_prerequisite_point11_valc"}
				model.ValC = SnapshotPoint11ValDValCDependencyFromComputed(valC, Point11ValDValCReviewContext{
					ReviewPrerequisites: []string{"external_review_prerequisite_point11_valc"},
				})
			},
			want: Point11ValDDependencyStateReviewRequired,
		},
	}

	for _, testCase := range disclaimerTests {
		t.Run(testCase.name, func(t *testing.T) {
			model := point11ValDActiveDependencyBundle()
			testCase.mutate(&model)
			if got := EvaluatePoint11ValDDependencyState(model); got != testCase.want {
				t.Fatalf("expected dependency state %q, got %#v", testCase.want, model)
			}
		})
	}
}

func TestPoint11ValDIntegratedInvariantState(t *testing.T) {
	t.Run("valid integrated invariant active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.IntegratedInvariantState != Point11ValDIntegratedInvariantStateActive {
			t.Fatalf("expected active integrated invariant state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValDIntegratedGovernanceInvariantReview)
	}{
		{name: "policy authority inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.PolicyAuthorityConsistencyState = point11ValDCheckStateBlocked
		}},
		{name: "claim authority inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.ClaimAuthorityConsistencyState = point11ValDCheckStateBlocked
		}},
		{name: "publication boundary inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.PublicationBoundaryState = point11ValDCheckStateBlocked
		}},
		{name: "no overclaim inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.NoOverclaimState = point11ValDCheckStateBlocked
		}},
		{name: "clean room ip inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.CleanRoomIPState = point11ValDCheckStateBlocked
		}},
		{name: "projection boundary inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.ProjectionBoundaryState = point11ValDCheckStateBlocked
		}},
		{name: "exception emergency inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.ExceptionEmergencyConsistencyState = point11ValDCheckStateBlocked
		}},
		{name: "abac enforcement inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.ABACEnforcementConsistencyState = point11ValDCheckStateBlocked
		}},
		{name: "dashboard projection inconsistency blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.DashboardProjectionState = point11ValDCheckStateBlocked
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *Point11ValDIntegratedGovernanceInvariantReview) {
			model.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model.IntegratedInvariantReview)
			model = ComputePoint11ValDFoundation(model)
			if model.IntegratedInvariantState != Point11ValDIntegratedInvariantStateBlocked {
				t.Fatalf("expected blocked integrated invariant state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDQualityMapState(t *testing.T) {
	t.Run("valid quality map active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.QualityMapState != Point11ValDQualityMapStateActive {
			t.Fatalf("expected active quality map, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValDEvidenceGovernanceQualityMap)
	}{
		{name: "missing required evidence ref blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) { model.EvidenceRefs = nil }},
		{name: "evidence unknown blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) { model.EvidenceRefs = []string{"unknown"} }},
		{name: "evidence revoked blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) { model.EvidenceRefs = []string{"revoked"} }},
		{name: "evidence invalid blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) { model.EvidenceRefs = []string{"invalid"} }},
		{name: "duplicate identity blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) {
			model.DuplicateState = point11ValDDuplicateStateBlocked
		}},
		{name: "conflicting identity blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) {
			model.ConflictState = point11ValDConflictStateBlocked
		}},
		{name: "unrelated evidence blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) {
			model.UnrelatedState = point11ValDUnrelatedStateBlocked
		}},
		{name: "cross tenant reuse blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) {
			model.TenantScopeState = point11ValDTenantScopeStateBlocked
		}},
		{name: "global unscoped wildcard scope blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) { model.PolicyRefs = []string{"policy_unknown"} }},
		{name: "stale freshness blocks where freshness required", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) {
			model.FreshnessState = point11ValDFreshnessStateStale
		}},
		{name: "revoked policy as active use authority blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) {
			model.RevocationState = point11ValDRevocationStateBlocked
		}},
		{name: "superseded claim as active use authority blocks", mutate: func(model *Point11ValDEvidenceGovernanceQualityMap) {
			model.SupersessionState = point11ValDSupersessionStateBlocked
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model.QualityMap)
			model = ComputePoint11ValDFoundation(model)
			if model.QualityMapState != Point11ValDQualityMapStateBlocked {
				t.Fatalf("expected blocked quality map state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDPublicationReviewState(t *testing.T) {
	t.Run("valid modeled only publication review active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.PublicationReviewState != Point11ValDPublicationReviewStateActive {
			t.Fatalf("expected active publication review, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValDPublicationBoundaryFinalReview)
	}{
		{name: "creates publication side effects blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.CreatesPublicationSideEffects = true }},
		{name: "creates customer facing material blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.CreatesCustomerFacingMaterial = true }},
		{name: "authority claim blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.CreatesAuthorityClaim = true }},
		{name: "certification claim blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.CreatesCertificationClaim = true }},
		{name: "regulatory claim blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.CreatesRegulatoryClaim = true }},
		{name: "compliance guarantee blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.CreatesComplianceGuarantee = true }},
		{name: "public customer export buyer partner demo sales docs portal without clean room ip blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.CleanRoomIPReviewRefs = nil }},
		{name: "public customer export buyer partner demo sales docs portal without governance event blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) { model.GovernanceEventRefs = nil }},
		{name: "agent output public customer claim without governance event blocks", mutate: func(model *Point11ValDPublicationBoundaryFinalReview) {
			model.AgentOutputSurfaces = []string{point11ValDPublicationSurfaceAgentOutput}
			model.CustomerVisibleSurfaces = []string{point11Val0PublicationSurfacePortal}
			model.GovernanceEventRefs = nil
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model.PublicationReview)
			model = ComputePoint11ValDFoundation(model)
			if model.PublicationReviewState != Point11ValDPublicationReviewStateBlocked {
				t.Fatalf("expected blocked publication review state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDNoOverclaimReviewState(t *testing.T) {
	t.Run("valid safe wording active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.NoOverclaimReviewState != Point11ValDNoOverclaimReviewStateActive {
			t.Fatalf("expected active no overclaim review, got %#v", model)
		}
	})

	for _, phrase := range []string{
		"certified",
		"regulator-approved",
		"compliance guaranteed",
		"production approved",
		"AI-approved",
		"AI decision",
		"continuous compliance attestation",
		"deployment approved",
		"public badge",
		"official authority",
		"global truth",
		"source of truth",
		"canonical truth",
		"supreme authority",
		"impossible to violate without detection",
	} {
		t.Run("phrase "+phrase+" blocks", func(t *testing.T) {
			model := activePoint11ValDFoundation()
			model.NoOverclaimReview.ObservedDashboardText = []string{phrase}
			model = ComputePoint11ValDFoundation(model)
			if model.NoOverclaimReviewState != Point11ValDNoOverclaimReviewStateBlocked {
				t.Fatalf("expected forbidden phrase to block, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDCleanRoomIPReviewState(t *testing.T) {
	t.Run("valid clean room ip review active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.CleanRoomIPReviewState != Point11ValDCleanRoomIPReviewStateActive {
			t.Fatalf("expected active clean room ip review, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValDCleanRoomIPFinalReview)
	}{
		{name: "copied competitor material blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) { model.CopiedCompetitorMaterialDetected = true }},
		{name: "copied ui blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) { model.CopiedUIDetected = true }},
		{name: "copied workflow blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) { model.CopiedWorkflowDetected = true }},
		{name: "copied private docs blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) { model.CopiedPrivateDocsDetected = true }},
		{name: "legal clearance claimed without external review blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) {
			model.LegalClearanceClaimed = true
			model.ExternalLegalReviewRef = ""
		}},
		{name: "fto claimed without external review blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) {
			model.FTOClaimed = true
			model.ExternalFTOReviewRef = ""
		}},
		{name: "third party ref without license ip review blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) {
			model.LicenseReviewRefs = nil
			model.IPReviewRefs = nil
		}},
		{name: "public buyer partner customer claim without clean room ip review blocks", mutate: func(model *Point11ValDCleanRoomIPFinalReview) {
			model.IPReviewRefs = nil
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model.CleanRoomIPReview)
			model = ComputePoint11ValDFoundation(model)
			if model.CleanRoomIPReviewState != Point11ValDCleanRoomIPReviewStateBlocked {
				t.Fatalf("expected blocked clean room ip review state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDCLBLedgerState(t *testing.T) {
	t.Run("valid empty clb0 1 2 ledger active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.CLBClosureState != Point11ValDCLBClosureStateActive {
			t.Fatalf("expected active clb closure state, got %#v", model)
		}
	})

	t.Run("clb3 advisory does not block active closure", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		model.CLBLedger.CLB3Findings = []string{"finding_point11_vald_advisory_extra_001"}
		model = ComputePoint11ValDFoundation(model)
		if model.CLBClosureState != Point11ValDCLBClosureStateActive {
			t.Fatalf("expected clb3 advisory not to block, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValDCLBClosureLedger)
	}{
		{name: "open clb0 blocks", mutate: func(model *Point11ValDCLBClosureLedger) {
			model.CLB0Findings = []string{"finding_point11_vald_clb0_001"}
		}},
		{name: "open clb1 blocks", mutate: func(model *Point11ValDCLBClosureLedger) {
			model.CLB1Findings = []string{"finding_point11_vald_clb1_001"}
		}},
		{name: "open clb2 blocks", mutate: func(model *Point11ValDCLBClosureLedger) {
			model.CLB2Findings = []string{"finding_point11_vald_clb2_001"}
		}},
		{name: "expired accepted risk blocks", mutate: func(model *Point11ValDCLBClosureLedger) {
			model.AcceptedRisks = []Point11ValDAcceptedRisk{{
				RiskID:      "accepted_risk_point11_vald_001",
				EvidenceRef: "evidence:point11-vald-risk-001",
				Scope:       "tenant_scope_alpha",
				OwnerRef:    "actor_point11_vald_owner",
				Expiry:      "2000-01-01T00:00:00Z",
			}}
		}},
		{name: "accepted risk without evidence scope owner expiry blocks", mutate: func(model *Point11ValDCLBClosureLedger) {
			model.AcceptedRisks = []Point11ValDAcceptedRisk{{RiskID: "accepted_risk_point11_vald_002"}}
		}},
		{name: "reviewer result pass but not pass confirmed blocks", mutate: func(model *Point11ValDCLBClosureLedger) { model.ReviewerResult = point11ValDReviewerResultPass }},
		{name: "pass confirmed with open clb1 blocks", mutate: func(model *Point11ValDCLBClosureLedger) {
			model.ReviewerResult = point11ValDReviewerResultPassConfirmed
			model.CLB1Findings = []string{"finding_point11_vald_clb1_open_002"}
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model.CLBLedger)
			model = ComputePoint11ValDFoundation(model)
			if model.CLBClosureState != Point11ValDCLBClosureStateBlocked {
				t.Fatalf("expected blocked clb closure state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDPassClosureManifestState(t *testing.T) {
	t.Run("valid pass closure manifest active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.PassClosureManifestState != Point11ValDPassClosureManifestStateActive {
			t.Fatalf("expected active pass closure manifest, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValDPassClosureManifest)
	}{
		{name: "wrong point id blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.PointID = "point_10" }},
		{name: "wrong wave id blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.WaveID = "val_c" }},
		{name: "missing dependency gate result blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.DependencyGateResult = "" }},
		{name: "missing tests run blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.TestsRun = nil }},
		{name: "missing greps run blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.GrepsRun = nil }},
		{name: "missing negative fixtures blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.NegativeFixturesRun = nil }},
		{name: "reviewer result pass but not pass confirmed blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.ReviewerResult = point11ValDReviewerResultPass }},
		{name: "point11 pass allowed false blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.Point11PassAllowed = false }},
		{name: "wrong point11 pass token blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.Point11PassToken = "point_11_fail" }},
		{name: "malformed projection disclaimer blocks", mutate: func(model *Point11ValDPassClosureManifest) { model.ProjectionDisclaimer = "canonical_truth" }},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model.PassClosureManifest)
			model = ComputePoint11ValDFoundation(model)
			if model.PassClosureManifestState != Point11ValDPassClosureManifestStateBlocked {
				t.Fatalf("expected blocked pass closure manifest state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDFinalPassGateState(t *testing.T) {
	t.Run("valid final pass gate emits point 11 pass", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.FinalPassGateState != Point11ValDFinalPassGateStateActive {
			t.Fatalf("expected active final pass gate, got %#v", model)
		}
		if model.Point11PassToken != point11ValDPoint11PassToken {
			t.Fatalf("expected point11 pass token, got %#v", model)
		}
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal foundation: %v", err)
		}
		if !strings.Contains(string(body), point11ValDPoint11PassToken) {
			t.Fatalf("expected final closure output to contain point11 pass token, got %s", body)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValDFoundation)
	}{
		{name: "blocked dependency gate prevents point11 pass", mutate: func(model *Point11ValDFoundation) {
			model.ValBDependency.VerificationState = Point11ValBVerificationStateBlocked
		}},
		{name: "review required dependency gate prevents point11 pass", mutate: func(model *Point11ValDFoundation) {
			valC := activePoint11ValCFoundation()
			valC.CurrentState = Point11ValCStateReviewRequired
			valC.DependencyState = Point11ValCDependencyStateReviewRequired
			valC.ReviewPrerequisites = []string{"external_review_prerequisite_point11_valc"}
			model.ValCDependency = SnapshotPoint11ValDValCDependencyFromComputed(valC, Point11ValDValCReviewContext{
				ReviewPrerequisites: []string{"external_review_prerequisite_point11_valc"},
			})
		}},
		{name: "blocked invariant prevents point11 pass", mutate: func(model *Point11ValDFoundation) {
			model.IntegratedInvariantReview.PolicyAuthorityConsistencyState = point11ValDCheckStateBlocked
		}},
		{name: "blocked quality map prevents point11 pass", mutate: func(model *Point11ValDFoundation) { model.QualityMap.DuplicateState = point11ValDDuplicateStateBlocked }},
		{name: "blocked publication boundary prevents point11 pass", mutate: func(model *Point11ValDFoundation) { model.PublicationReview.CreatesPublicationSideEffects = true }},
		{name: "blocked no overclaim prevents point11 pass", mutate: func(model *Point11ValDFoundation) {
			model.NoOverclaimReview.ObservedClaims = []string{"production approved"}
		}},
		{name: "blocked clean room ip prevents point11 pass", mutate: func(model *Point11ValDFoundation) { model.CleanRoomIPReview.CopiedCompetitorMaterialDetected = true }},
		{name: "blocked clb closure prevents point11 pass", mutate: func(model *Point11ValDFoundation) {
			model.CLBLedger.CLB1Findings = []string{"finding_point11_vald_clb1_003"}
		}},
		{name: "blocked manifest prevents point11 pass", mutate: func(model *Point11ValDFoundation) { model.PassClosureManifest.Point11PassAllowed = false }},
		{name: "point11 pass outside val d final closure blocks", mutate: func(model *Point11ValDFoundation) { model.FinalPassGate.Point11PassObservedOutsideFinalClosure = true }},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model)
			model = ComputePoint11ValDFoundation(model)
			if model.FinalPassGateState != Point11ValDFinalPassGateStateBlocked {
				t.Fatalf("expected blocked final pass gate, got %#v", model)
			}
			if model.Point11PassToken != "" {
				t.Fatalf("expected no point11 pass token, got %#v", model)
			}
		})
	}
}

func TestPoint11ValDAggregateState(t *testing.T) {
	t.Run("aggregate active only when all final gates active", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.CurrentState != Point11ValDStateActive {
			t.Fatalf("expected active aggregate state, got %#v", model)
		}
	})

	t.Run("aggregate emits point11 pass only in final happy path", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.Point11PassToken != point11ValDPoint11PassToken {
			t.Fatalf("expected point11 pass token only in final happy path, got %#v", model)
		}
	})

	t.Run("aggregate blocked when any local component blocked", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		model.PublicationReview.CreatesPublicationSideEffects = true
		model = ComputePoint11ValDFoundation(model)
		if model.CurrentState != Point11ValDStateBlocked {
			t.Fatalf("expected blocked aggregate state, got %#v", model)
		}
	})

	t.Run("aggregate review required cannot emit point11 pass", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		valC := activePoint11ValCFoundation()
		valC.CurrentState = Point11ValCStateReviewRequired
		valC.DependencyState = Point11ValCDependencyStateReviewRequired
		valC.ReviewPrerequisites = []string{"external_review_prerequisite_point11_valc"}
		model.ValCDependency = SnapshotPoint11ValDValCDependencyFromComputed(valC, Point11ValDValCReviewContext{
			ReviewPrerequisites: []string{"external_review_prerequisite_point11_valc"},
		})
		model = ComputePoint11ValDFoundation(model)
		if model.CurrentState != Point11ValDStateReviewRequired {
			t.Fatalf("expected review required aggregate state, got %#v", model)
		}
		if model.Point11PassToken != "" {
			t.Fatalf("expected no point11 pass token under review required, got %#v", model)
		}
	})

	t.Run("blocked overrides review required", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		valC := activePoint11ValCFoundation()
		valC.CurrentState = Point11ValCStateReviewRequired
		valC.DependencyState = Point11ValCDependencyStateReviewRequired
		valC.ReviewPrerequisites = []string{"external_review_prerequisite_point11_valc"}
		model.ValCDependency = SnapshotPoint11ValDValCDependencyFromComputed(valC, Point11ValDValCReviewContext{
			ReviewPrerequisites: []string{"external_review_prerequisite_point11_valc"},
		})
		model.CLBLedger.CLB2Findings = []string{"finding_point11_vald_clb2_open_001"}
		model = ComputePoint11ValDFoundation(model)
		if model.CurrentState != Point11ValDStateBlocked {
			t.Fatalf("expected blocked to override review required, got %#v", model)
		}
	})

	t.Run("diagnostics include component blocking reason", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		model.QualityMap.DuplicateState = point11ValDDuplicateStateBlocked
		model = ComputePoint11ValDFoundation(model)
		if !point11Val0ContainsTrimmed(model.Diagnostics.BlockingReasons, "quality_map_blocked") {
			t.Fatalf("expected diagnostics to include quality map blocker, got %#v", model)
		}
	})

	t.Run("aggregate creates no legal regulatory certification authority", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.CreatesAuthorityClaims {
			t.Fatalf("expected no authority claims, got %#v", model)
		}
	})

	t.Run("aggregate creates no publication side effects", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.CreatesPublicationSideEffects {
			t.Fatalf("expected no publication side effects, got %#v", model)
		}
	})

	t.Run("aggregate creates no signing anchoring external api production mutation side effects", func(t *testing.T) {
		model := activePoint11ValDFoundation()
		if model.CreatesSigningSideEffects || model.CreatesAnchoringSideEffects || model.CreatesExternalAPISideEffects || model.CreatesProductionSideEffects {
			t.Fatalf("expected no external side effects, got %#v", model)
		}
	})
}

func TestPoint11ValDSemanticAntiGreenRefs(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11ValDFoundation)
	}{
		{name: "manifest unknown blocks even with valid prefix", mutate: func(model *Point11ValDFoundation) {
			model.PassClosureManifest.ManifestID = "manifest_unknown"
		}},
		{name: "closure revoked blocks even with valid prefix", mutate: func(model *Point11ValDFoundation) {
			model.CLBLedger.ClosureLedgerID = "closure_revoked"
		}},
		{name: "quality invalid blocks even with valid prefix", mutate: func(model *Point11ValDFoundation) {
			model.QualityMap.QualityMapID = "quality_map_invalid"
		}},
		{name: "publication expired blocks even with valid prefix", mutate: func(model *Point11ValDFoundation) {
			model.PublicationReview.PublicationReviewID = "publication_review_expired"
		}},
		{name: "clean room placeholder blocks even with valid prefix", mutate: func(model *Point11ValDFoundation) {
			model.CleanRoomIPReview.CleanRoomReviewID = "clean_room_review_placeholder"
		}},
		{name: "final gate blocked blocks even with valid prefix", mutate: func(model *Point11ValDFoundation) {
			model.FinalPassGate.FinalGateID = "final_gate_blocked"
		}},
		{name: "revoked invalid marker blocks in every critical val d ref family tested", mutate: func(model *Point11ValDFoundation) {
			model.IntegratedInvariantReview.InvariantReviewID = "revoked/invalid marker"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValDFoundation()
			testCase.mutate(&model)
			model = ComputePoint11ValDFoundation(model)
			if model.CurrentState != Point11ValDStateBlocked {
				t.Fatalf("expected semantic anti-green case to block, got %#v", model)
			}
		})
	}
}
