package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

var (
	point13ValAActiveFoundationBaselineJSON []byte
	point13ValAActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint13ValAFoundation(model Point13ValAFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint13ValAFoundation(payload []byte) Point13ValAFoundation {
	var clone Point13ValAFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func uncachedActivePoint13ValAFoundation() Point13ValAFoundation {
	return ComputePoint13ValAFoundation(Point13ValAFoundationModel())
}

func activePoint13ValAFoundation() Point13ValAFoundation {
	point13ValAActiveFoundationBaselineOnce.Do(func() {
		point13ValAActiveFoundationBaselineJSON = mustMarshalPoint13ValAFoundation(uncachedActivePoint13ValAFoundation())
	})
	return clonePoint13ValAFoundation(point13ValAActiveFoundationBaselineJSON)
}

func TestPoint13ValAFoundationFixtureIsolation(t *testing.T) {
	t.Run("raw production path still computes", func(t *testing.T) {
		model := uncachedActivePoint13ValAFoundation()
		if model.CurrentState != Point13ValAStateActive {
			t.Fatalf("expected raw production path to compute active foundation, got %#v", model)
		}
	})

	t.Run("cached fixture mutation does not contaminate next clone", func(t *testing.T) {
		mutated := activePoint13ValAFoundation()
		mutated.Dependency.Val0CurrentState = Point13Val0StateBlocked
		mutated.PilotExecutionContract.PilotOwnerRef = ""
		mutated.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs = []string{"artifact_point13_vala_other_candidate_001"}

		fresh := activePoint13ValAFoundation()
		if fresh.Dependency.Val0CurrentState != Point13Val0StateActive {
			t.Fatalf("expected dependency to remain active on fresh clone, got %#v", fresh.Dependency)
		}
		if fresh.PilotExecutionContract.PilotOwnerRef == "" {
			t.Fatalf("expected pilot owner ref on fresh clone, got %#v", fresh.PilotExecutionContract)
		}
		if !point12Val0ExactStringSetMatch(fresh.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs, []string{
			"artifact_point13_vala_customer_candidate_001",
			"artifact_point13_vala_customer_candidate_002",
		}) {
			t.Fatalf("expected intake refs to remain canonical on fresh clone, got %#v", fresh.CustomerIntakeEvidenceGovernance)
		}
	})
}

func TestPoint13ValADependencyState(t *testing.T) {
	t.Run("valid val0 dependency active", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		if model.DependencyState != Point13ValAStateActive || model.CurrentState != Point13ValAStateActive {
			t.Fatalf("expected active dependency and foundation, got %#v", model)
		}
	})

	t.Run("missing val0 dependency blocks", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0CurrentState = ""
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected missing Val0 dependency to block, got %#v", model)
		}
	})

	t.Run("val0 blocked blocks vala", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0CurrentState = Point13Val0StateBlocked
		model.Dependency.Val0.CurrentState = Point13Val0StateBlocked
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected blocked Val0 to block ValA, got %#v", model)
		}
	})

	t.Run("stale val0 review required summary blocks vala", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0CurrentState = Point13Val0StateReviewRequired
		model.Dependency.Val0.CurrentState = Point13Val0StateReviewRequired
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected stale review required Val0 summary to block ValA, got %#v", model)
		}
	})

	t.Run("stale val0 incomplete summary blocks vala", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0CurrentState = Point13Val0StateIncomplete
		model.Dependency.Val0.CurrentState = Point13Val0StateIncomplete
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected stale incomplete Val0 summary to block ValA, got %#v", model)
		}
	})

	t.Run("val0 point13 pass appearance blocks", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0Point13PassSeen = true
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected point_13_pass appearance in Val0 dependency to block, got %#v", model)
		}
	})

	t.Run("local vala readiness cannot override val0 failure", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0CurrentState = Point13Val0StateBlocked
		model.Dependency.Val0.CurrentState = Point13Val0StateBlocked
		model = ComputePoint13ValAFoundation(model)
		if model.CurrentState == Point13ValAStateActive {
			t.Fatalf("expected local ValA readiness not to override Val0 failure, got %#v", model)
		}
	})

	t.Run("stale point12 inherited review requirement through val0 blocks", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Point12CurrentState = Point12ValEStateReviewRequired
		model.Dependency.Val0.Dependency.Point12CurrentState = Point12ValEStateReviewRequired
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected stale inherited Point12 review requirement to block, got %#v", model)
		}
	})

	t.Run("padded nested val0 state blocks recomputed dependency binding", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0.CurrentState = " " + Point13Val0StateActive + " "
		state, reasons := point13ValADependencyStateAndReasons(model.Dependency)
		if state != Point13ValAStateBlocked || !point13Val0StringSliceContains(reasons, "val0_recomputed_snapshot_mismatch") {
			t.Fatalf("expected recomputed dependency snapshot mismatch, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected padded nested Val0 state to block foundation, got %#v", model)
		}
	})

	t.Run("padded val0 point identity blocks exact dependency identity", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0PointID = " " + point13Val0PointID + " "
		state, reasons := point13ValADependencyStateAndReasons(model.Dependency)
		if state != Point13ValAStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected padded Val0 point identity to block foundation, got %#v", model)
		}
	})

	t.Run("padded val0 wave identity blocks exact dependency identity", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0WaveID = " " + point13Val0WaveID + " "
		state, reasons := point13ValADependencyStateAndReasons(model.Dependency)
		if state != Point13ValAStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected padded Val0 wave identity to block foundation, got %#v", model)
		}
	})

	t.Run("padded inherited point12 pass token blocks raw-exact dependency binding", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		padded := " " + point12ValEPoint12PassToken + " "
		model.Dependency.Point12PassToken = padded
		model.Dependency.Val0.Dependency.Point12PassToken = padded
		state, reasons := point13ValADependencyStateAndReasons(model.Dependency)
		if state != Point13ValAStateBlocked || !point13Val0StringSliceContains(reasons, "point12_inherited_not_pass_confirmed") {
			t.Fatalf("expected exact point12 pass-token mismatch, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected padded inherited Point12 pass token to block foundation, got %#v", model)
		}
	})

	t.Run("stale embedded val0 point12 profile mutation blocks recompute", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.Dependency.Val0.Dependency.Point12.Dependency.Val0.Manifest.ProfileContext.CurrentProfileHash = ""
		state, reasons := point13ValADependencyStateAndReasons(model.Dependency)
		if state != Point13ValAStateBlocked || !point13Val0StringSliceContains(reasons, "val0_recomputed_snapshot_mismatch") {
			t.Fatalf("expected exact Val0 recomputed snapshot mismatch, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValAFoundation(model)
		if model.DependencyState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected stale embedded Val0 profile mutation to block foundation, got %#v", model)
		}
	})
}

func TestPoint13ValAFoundationAggregationState(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point13ValAFoundation)
	}{
		{name: "pilot run phase boundary review required prevents active", mutate: func(model *Point13ValAFoundation) {
			model.PilotRunPhaseBoundaryState = Point13ValAStateReviewRequired
		}},
		{name: "pilot execution contract review required prevents active", mutate: func(model *Point13ValAFoundation) {
			model.PilotExecutionContractState = Point13ValAStateReviewRequired
		}},
		{name: "customer intake evidence governance review required prevents active", mutate: func(model *Point13ValAFoundation) {
			model.CustomerIntakeEvidenceGovernanceState = Point13ValAStateReviewRequired
		}},
		{name: "support responsibility matrix review required prevents active", mutate: func(model *Point13ValAFoundation) {
			model.SupportResponsibilityMatrixState = Point13ValAStateReviewRequired
		}},
		{name: "pilot exit review gate review required prevents active", mutate: func(model *Point13ValAFoundation) {
			model.PilotExitReviewGateState = Point13ValAStateReviewRequired
		}},
		{name: "ai assisted pilot execution boundary review required prevents active", mutate: func(model *Point13ValAFoundation) {
			model.AIAssistedPilotExecutionBoundaryState = Point13ValAStateReviewRequired
		}},
		{name: "no overclaim review required prevents active", mutate: func(model *Point13ValAFoundation) {
			model.NoOverclaimState = Point13ValAStateReviewRequired
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model)
			if got := EvaluatePoint13ValAState(model); got != Point13ValAStateReviewRequired {
				t.Fatalf("expected review_required aggregation, got %q for %#v", got, model)
			}
		})
	}

	t.Run("blocked outranks review required", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.PilotExecutionContractState = Point13ValAStateBlocked
		model.PilotRunPhaseBoundaryState = Point13ValAStateReviewRequired
		if got := EvaluatePoint13ValAState(model); got != Point13ValAStateBlocked {
			t.Fatalf("expected blocked to outrank review_required, got %q for %#v", got, model)
		}
	})

	t.Run("incomplete returned only when no blocked or review required exists", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.SupportResponsibilityMatrixState = Point13ValAStateIncomplete
		if got := EvaluatePoint13ValAState(model); got != Point13ValAStateIncomplete {
			t.Fatalf("expected incomplete aggregation, got %q for %#v", got, model)
		}
	})

	t.Run("active returned only when all component states active", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		if got := EvaluatePoint13ValAState(model); got != Point13ValAStateActive {
			t.Fatalf("expected active aggregation, got %q for %#v", got, model)
		}
	})
}

func TestPoint13ValAPilotExecutionContractState(t *testing.T) {
	t.Run("valid pilot execution contract becomes active without point13 pass", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		payload := string(mustMarshalPoint13ValAFoundation(model))
		if model.PilotExecutionContractState != Point13ValAStateActive || strings.Contains(payload, point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected active ValA contract without point_13_pass, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValAPilotExecutionContract)
	}{
		{name: "missing pilot owner blocks", mutate: func(model *Point13ValAPilotExecutionContract) { model.PilotOwnerRef = "" }},
		{name: "missing customer owner blocks", mutate: func(model *Point13ValAPilotExecutionContract) { model.CustomerOwnerRef = "" }},
		{name: "missing tenant scope blocks", mutate: func(model *Point13ValAPilotExecutionContract) { model.TenantScope = "" }},
		{name: "missing entry criteria blocks", mutate: func(model *Point13ValAPilotExecutionContract) { model.EntryCriteriaRefs = nil }},
		{name: "missing exit criteria blocks", mutate: func(model *Point13ValAPilotExecutionContract) { model.ExitCriteriaRefs = nil }},
		{name: "production approval excluded must be true", mutate: func(model *Point13ValAPilotExecutionContract) { model.ProductionApprovalExcluded = false }},
		{name: "deployment approval excluded must be true", mutate: func(model *Point13ValAPilotExecutionContract) { model.DeploymentApprovalExcluded = false }},
		{name: "compliance guarantee excluded must be true", mutate: func(model *Point13ValAPilotExecutionContract) { model.ComplianceGuaranteeExcluded = false }},
		{name: "customer success cannot mean production approval", mutate: func(model *Point13ValAPilotExecutionContract) { model.CustomerSuccessNotProductionApproval = false }},
		{name: "contract cannot create pass", mutate: func(model *Point13ValAPilotExecutionContract) { model.ContractCannotCreatePass = false }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model.PilotExecutionContract)
			model = ComputePoint13ValAFoundation(model)
			if model.PilotExecutionContractState != Point13ValAStateBlocked {
				t.Fatalf("expected pilot execution contract mutation to block, got %#v", model)
			}
		})
	}
}

func TestPoint13ValACustomerIntakeEvidenceGovernanceState(t *testing.T) {
	t.Run("valid customer intake evidence remains candidate only", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		if model.CustomerIntakeEvidenceGovernanceState != Point13ValAStateActive ||
			!model.CustomerIntakeEvidenceGovernance.EvidenceCandidateOnly {
			t.Fatalf("expected active candidate-only intake governance, got %#v", model)
		}
	})

	t.Run("valid same tenant artifact still passes", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs = []string{"artifact_same_tenant_candidate_001"}
		model.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs = []string{"evidence_hash_same_tenant_candidate_001"}
		model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		model = ComputePoint13ValAFoundation(model)
		if model.CustomerIntakeEvidenceGovernanceState != Point13ValAStateActive {
			t.Fatalf("expected valid same-tenant artifact to remain active, got %#v", model)
		}
	})

	t.Run("artifact hash mismatch blocks", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs[0] = "evidence_hash_point13_vala_other_candidate_001"
		model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		model = ComputePoint13ValAFoundation(model)
		if model.CustomerIntakeEvidenceGovernanceState != Point13ValAStateBlocked {
			t.Fatalf("expected artifact/hash mismatch to block, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValACustomerIntakeEvidenceGovernance)
	}{
		{name: "missing custody ref blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) { model.CustodyRef = "" }},
		{name: "missing source owner blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) { model.SourceOwnerRef = "" }},
		{name: "missing consent authority ref blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) { model.ConsentOrAuthorityRef = "" }},
		{name: "promotion to canonical without governance event blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactPromotedToCanonical = true
		}},
		{name: "customer upload cannot mutate canonical spine", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerUploadCannotMutateCanonicalSpine = false
		}},
		{name: "support attachment cannot mutate canonical spine", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.SupportAttachmentCannotMutateCanonicalSpine = false
		}},
		{name: "cross tenant artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_cross_tenant_candidate_001"
		}},
		{name: "cross-tenant artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_cross-tenant_candidate_001"
		}},
		{name: "other tenant artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_other_tenant_candidate_001"
		}},
		{name: "other-tenant artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_other-tenant_candidate_001"
		}},
		{name: "all tenants artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_all_tenants_candidate_001"
		}},
		{name: "all-tenants artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_all-tenants_candidate_001"
		}},
		{name: "global artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_global_candidate_001"
		}},
		{name: "wildcard artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_wildcard_candidate_001"
		}},
		{name: "unscoped artifact blocks", mutate: func(model *Point13ValACustomerIntakeEvidenceGovernance) {
			model.CustomerArtifactRefs[0] = "artifact_unscoped_candidate_001"
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model.CustomerIntakeEvidenceGovernance)
			model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
			model = ComputePoint13ValAFoundation(model)
			if model.CustomerIntakeEvidenceGovernanceState != Point13ValAStateBlocked {
				t.Fatalf("expected customer intake governance mutation to block, got %#v", model)
			}
		})
	}

	t.Run("artifact cross tenant helper detects space-separated variant", func(t *testing.T) {
		if !point13ValAContainsCrossTenantArtifact([]string{"artifact_cross tenant_candidate_001"}) {
			t.Fatalf("expected cross-tenant helper to detect space-separated variant")
		}
	})

	t.Run("recomputing local intake hash after artifact hash source tenant mutation cannot hide drift", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.CustomerIntakeEvidenceGovernance.TenantScope = "tenant_scope_point13_vala_wrong_001"
		model.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs = []string{"artifact_point13_vala_customer_candidate_999"}
		model.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs = []string{"evidence_hash_point13_vala_customer_candidate_999"}
		model.CustomerIntakeEvidenceGovernance.SourceOwnerRef = "customer_owner_point13_vala_999"
		model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		model = ComputePoint13ValAFoundation(model)
		if model.CurrentState == Point13ValAStateActive {
			t.Fatalf("expected recomputed intake hash not to hide drift, got %#v", model)
		}
	})

	t.Run("recomputing local intake hash after cross-tenant artifact mutation cannot hide drift", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs = []string{"artifact_cross-tenant_candidate_001"}
		model.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs = []string{"evidence_hash_cross-tenant_candidate_001"}
		model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		model = ComputePoint13ValAFoundation(model)
		if model.CurrentState == Point13ValAStateActive {
			t.Fatalf("expected recomputed intake hash not to hide cross-tenant drift, got %#v", model)
		}
	})
}

func TestPoint13ValAPilotRunPhaseBoundaryState(t *testing.T) {
	t.Run("all allowed phase taxonomy values validate", func(t *testing.T) {
		for _, phaseType := range point13ValAPhaseTaxonomy() {
			model := activePoint13ValAFoundation()
			model.PilotRunPhaseBoundary.Phases[0].PhaseType = phaseType
			model = ComputePoint13ValAFoundation(model)
			if model.PilotRunPhaseBoundaryState != Point13ValAStateActive {
				t.Fatalf("expected phase taxonomy %q to remain active, got %#v", phaseType, model)
			}
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValAPilotRunPhase)
	}{
		{name: "invalid phase taxonomy blocks", mutate: func(model *Point13ValAPilotRunPhase) { model.PhaseType = "production_rollout" }},
		{name: "missing phase owner blocks", mutate: func(model *Point13ValAPilotRunPhase) { model.PhaseOwnerRef = "" }},
		{name: "missing evidence refs blocks", mutate: func(model *Point13ValAPilotRunPhase) { model.EvidenceRefs = nil }},
		{name: "missing audit event blocks", mutate: func(model *Point13ValAPilotRunPhase) { model.AuditEventRef = "" }},
		{name: "phase cannot approve production", mutate: func(model *Point13ValAPilotRunPhase) { model.ProductionApprovalAttempted = true }},
		{name: "phase cannot deploy", mutate: func(model *Point13ValAPilotRunPhase) { model.DeploymentAttempted = true }},
		{name: "phase cannot create pass", mutate: func(model *Point13ValAPilotRunPhase) { model.PassAttempted = true }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model.PilotRunPhaseBoundary.Phases[0])
			model = ComputePoint13ValAFoundation(model)
			if model.PilotRunPhaseBoundaryState == Point13ValAStateActive {
				t.Fatalf("expected pilot run phase mutation to block, got %#v", model)
			}
		})
	}
}

func TestPoint13ValASupportResponsibilityMatrixState(t *testing.T) {
	t.Run("valid support responsibility matrix active", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		if model.SupportResponsibilityMatrixState != Point13ValAStateActive {
			t.Fatalf("expected active support responsibility matrix, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValACustomerSupportResponsibilityMatrix)
	}{
		{name: "support canonical evidence mutation blocks", mutate: func(model *Point13ValACustomerSupportResponsibilityMatrix) {
			model.SupportCanMutateCanonicalEvidence = true
		}},
		{name: "support core decision override blocks", mutate: func(model *Point13ValACustomerSupportResponsibilityMatrix) {
			model.SupportCanOverrideCoreDecision = true
		}},
		{name: "support production approval blocks", mutate: func(model *Point13ValACustomerSupportResponsibilityMatrix) { model.SupportCanApproveProduction = true }},
		{name: "missing support audit event blocks", mutate: func(model *Point13ValACustomerSupportResponsibilityMatrix) { model.AuditEventRefs = nil }},
		{name: "audit event ref mutation blocks", mutate: func(model *Point13ValACustomerSupportResponsibilityMatrix) { model.AuditEventRefs[0] = "bad_ref" }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model.SupportResponsibilityMatrix)
			model = ComputePoint13ValAFoundation(model)
			if model.SupportResponsibilityMatrixState != Point13ValAStateBlocked {
				t.Fatalf("expected support responsibility mutation to block, got %#v", model)
			}
		})
	}
}

func TestPoint13ValAPilotExitReviewGateState(t *testing.T) {
	t.Run("valid exit review gate active as operational readiness only", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		if model.PilotExitReviewGateState != Point13ValAStateActive {
			t.Fatalf("expected active pilot exit review gate, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValAPilotExitReviewGate, *Point13ValAAIAssistedPilotExecutionBoundary)
	}{
		{name: "unresolved blockers block", mutate: func(model *Point13ValAPilotExitReviewGate, _ *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.UnresolvedBlockers = []string{"open_blocker_point13_vala_001"}
		}},
		{name: "production approval requested blocks", mutate: func(model *Point13ValAPilotExitReviewGate, _ *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.ProductionApprovalRequested = true
		}},
		{name: "deployment approval requested blocks", mutate: func(model *Point13ValAPilotExitReviewGate, _ *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.DeploymentApprovalRequested = true
		}},
		{name: "compliance claim requested blocks", mutate: func(model *Point13ValAPilotExitReviewGate, _ *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.ComplianceClaimRequested = true
		}},
		{name: "forbidden final customer statement blocks", mutate: func(model *Point13ValAPilotExitReviewGate, _ *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.FinalCustomerStatement = "production approved"
		}},
		{name: "ai output cannot satisfy exit gate by itself", mutate: func(model *Point13ValAPilotExitReviewGate, ai *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.EvidenceRefs = []string{ai.EvidenceCandidateRef}
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model.PilotExitReviewGate, &model.AIAssistedPilotExecutionBoundary)
			model = ComputePoint13ValAFoundation(model)
			if model.PilotExitReviewGateState != Point13ValAStateBlocked {
				t.Fatalf("expected pilot exit review mutation to block, got %#v", model)
			}
		})
	}

	t.Run("exit success cannot create pass or point13 pass", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		payload := string(mustMarshalPoint13ValAFoundation(model))
		if model.PilotExitReviewGateState != Point13ValAStateActive || model.CurrentState != Point13ValAStateActive {
			t.Fatalf("expected operational-readiness-only active ValA exit state, got %#v", model)
		}
		if strings.Contains(payload, point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected no point_13_pass token in ValA payload, got %s", payload)
		}
	})
}

func TestPoint13ValAAIAssistedPilotExecutionBoundaryState(t *testing.T) {
	t.Run("all allowed ai output types remain advisory when authority flags false", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValAFoundation()
			model.AIAssistedPilotExecutionBoundary.AIOutputType = outputType
			model = ComputePoint13ValAFoundation(model)
			if model.AIAssistedPilotExecutionBoundaryState != Point13ValAStateActive {
				t.Fatalf("expected allowed AI output type %q to remain active advisory candidate, got %#v", outputType, model)
			}
		}
	})

	t.Run("all allowed ai output types block on deployment authorized", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValAFoundation()
			model.AIAssistedPilotExecutionBoundary.AIOutputType = outputType
			model.AIAssistedPilotExecutionBoundary.DeploymentAuthorized = true
			model = ComputePoint13ValAFoundation(model)
			if model.AIAssistedPilotExecutionBoundaryState != Point13ValAStateBlocked {
				t.Fatalf("expected deployment authority on %q to block, got %#v", outputType, model)
			}
		}
	})

	t.Run("all allowed ai output types block on production readiness claimed", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValAFoundation()
			model.AIAssistedPilotExecutionBoundary.AIOutputType = outputType
			model.AIAssistedPilotExecutionBoundary.ProductionReadinessClaimed = true
			model = ComputePoint13ValAFoundation(model)
			if model.AIAssistedPilotExecutionBoundaryState != Point13ValAStateBlocked {
				t.Fatalf("expected production readiness claim on %q to block, got %#v", outputType, model)
			}
		}
	})

	t.Run("all blocked ai taxonomy values rejected", func(t *testing.T) {
		for _, outputType := range point12Val0BlockedAIEvidenceCandidateTypes() {
			model := activePoint13ValAFoundation()
			model.AIAssistedPilotExecutionBoundary.AIOutputType = outputType
			model = ComputePoint13ValAFoundation(model)
			if model.AIAssistedPilotExecutionBoundaryState != Point13ValAStateBlocked {
				t.Fatalf("expected blocked AI output type %q to block, got %#v", outputType, model)
			}
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValAAIAssistedPilotExecutionBoundary)
	}{
		{name: "ai approval request is not approval", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.AIOutputType = "AI_APPROVAL_REQUEST"
			model.ApprovalGranted = true
		}},
		{name: "ai patch proposal is not deployment", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.AIOutputType = "AI_PATCH_PROPOSAL"
			model.DeploymentAuthorized = true
		}},
		{name: "ai sandbox result is not production readiness", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.AIOutputType = "AI_SANDBOX_RESULT"
			model.ProductionReadinessClaimed = true
		}},
		{name: "ai output cannot promote customer artifact to canonical evidence", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.AIOutputType = "AI_FINDING"
			model.CanonicalMutationAllowed = true
		}},
		{name: "external api allowed without governance event blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.ExternalAPIAllowed = true
		}},
		{name: "permission manifest hash mutation blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.PermissionManifestHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}},
		{name: "input evidence hash refs mutation blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.InputEvidenceHashRefs[0] = "evidence_hash_point13_vala_customer_candidate_999"
		}},
		{name: "tenant scope mutation blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.TenantScope = "tenant_scope_point13_vala_wrong_001"
		}},
		{name: "model or rule version ref mutation blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.ModelOrRuleVersionRef = "model_version_point13_vala_wrong_001"
		}},
		{name: "audit event ref mutation blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.AuditEventRef = "audit_point13_vala_wrong_001"
		}},
		{name: "ai finding deployment authority blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.AIOutputType = "AI_FINDING"
			model.DeploymentAuthorized = true
		}},
		{name: "ai recommendation production readiness claim blocks", mutate: func(model *Point13ValAAIAssistedPilotExecutionBoundary) {
			model.AIOutputType = "AI_RECOMMENDATION"
			model.ProductionReadinessClaimed = true
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model.AIAssistedPilotExecutionBoundary)
			model = ComputePoint13ValAFoundation(model)
			if model.AIAssistedPilotExecutionBoundaryState == Point13ValAStateActive {
				t.Fatalf("expected AI-assisted pilot execution mutation to block/review, got %#v", model)
			}
		})
	}
}

func TestPoint13ValANoOverclaimState(t *testing.T) {
	for _, phrase := range point13Val0ForbiddenClaims() {
		t.Run("forbidden wording blocks "+phrase, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{phrase}
			model = ComputePoint13ValAFoundation(model)
			if model.NoOverclaimState != Point13ValAStateBlocked {
				t.Fatalf("expected forbidden phrase %q to block, got %#v", phrase, model)
			}
		})
	}

	t.Run("safe wording remains allowed", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = append([]string{}, point13ValAAllowedCustomerWording()...)
		model = ComputePoint13ValAFoundation(model)
		if model.NoOverclaimState != Point13ValAStateActive {
			t.Fatalf("expected safe wording to remain allowed, got %#v", model)
		}
	})

	t.Run("forbidden wording cannot be laundered through allowed list", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.NoOverclaimCustomerWording.AllowedCustomerFacingWording = []string{"deployment approved"}
		model = ComputePoint13ValAFoundation(model)
		if model.NoOverclaimState != Point13ValAStateBlocked || model.CurrentState != Point13ValAStateBlocked {
			t.Fatalf("expected forbidden allowed wording list mutation to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13ValAStateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("forbidden wording may appear only in classified internal diagnostics", func(t *testing.T) {
		model := activePoint13ValAFoundation()
		model.NoOverclaimCustomerWording.InternalDiagnosticTexts = []string{"blocked phrase: production approved"}
		model.NoOverclaimCustomerWording.InternalDiagnosticsClassifiedBlocked = true
		model = ComputePoint13ValAFoundation(model)
		if model.NoOverclaimState != Point13ValAStateActive {
			t.Fatalf("expected classified internal diagnostics to remain active, got %#v", model)
		}
	})
}

func TestPoint13ValAMutationClosure(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point13ValAFoundation)
	}{
		{name: "mutate tenant scope blocks", mutate: func(model *Point13ValAFoundation) {
			model.PilotExecutionContract.TenantScope = "tenant_scope_point13_vala_wrong_001"
		}},
		{name: "mutate evidence refs blocks", mutate: func(model *Point13ValAFoundation) {
			model.PilotExitReviewGate.EvidenceRefs = []string{"artifact_point13_vala_customer_candidate_999"}
		}},
		{name: "mutate evidence hash refs blocks", mutate: func(model *Point13ValAFoundation) {
			model.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs[0] = "evidence_hash_point13_vala_customer_candidate_999"
			model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		}},
		{name: "mutate custody ref blocks", mutate: func(model *Point13ValAFoundation) {
			model.CustomerIntakeEvidenceGovernance.CustodyRef = "bad_ref"
			model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		}},
		{name: "mutate governance event refs blocks", mutate: func(model *Point13ValAFoundation) {
			model.CustomerIntakeEvidenceGovernance.CustomerArtifactPromotedToCanonical = true
			model.CustomerIntakeEvidenceGovernance.CanonicalizationGovernanceEventRef = "bad_ref"
			model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		}},
		{name: "mutate support audit event ref blocks", mutate: func(model *Point13ValAFoundation) {
			model.SupportResponsibilityMatrix.AuditEventRefs[0] = "bad_ref"
		}},
		{name: "mutate exit review ref blocks", mutate: func(model *Point13ValAFoundation) {
			model.PilotExitReviewGate.CustomerReviewRef = ""
		}},
		{name: "recomputing local intake hash after dependency provenance mutation does not hide drift", mutate: func(model *Point13ValAFoundation) {
			model.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs = []string{"artifact_point13_vala_customer_candidate_999"}
			model.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs = []string{"evidence_hash_point13_vala_customer_candidate_999"}
			model.CustomerIntakeEvidenceGovernance.IntakeBindingHash = point13ValAComputedIntakeBindingHash(model.CustomerIntakeEvidenceGovernance)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValAFoundation()
			tc.mutate(&model)
			model = ComputePoint13ValAFoundation(model)
			if model.CurrentState == Point13ValAStateActive {
				t.Fatalf("expected ValA mutation closure to fail closed, got %#v", model)
			}
		})
	}
}
