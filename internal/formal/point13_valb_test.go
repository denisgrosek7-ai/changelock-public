package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

var (
	point13ValBActiveFoundationBaselineJSON []byte
	point13ValBActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint13ValBFoundation(model Point13ValBFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint13ValBFoundation(payload []byte) Point13ValBFoundation {
	var clone Point13ValBFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func uncachedActivePoint13ValBFoundation() Point13ValBFoundation {
	return ComputePoint13ValBFoundation(Point13ValBFoundationModel())
}

func activePoint13ValBFoundation() Point13ValBFoundation {
	point13ValBActiveFoundationBaselineOnce.Do(func() {
		point13ValBActiveFoundationBaselineJSON = mustMarshalPoint13ValBFoundation(uncachedActivePoint13ValBFoundation())
	})
	return clonePoint13ValBFoundation(point13ValBActiveFoundationBaselineJSON)
}

func point13ValBHashRefsForEvidenceRefs(refs []string) []string {
	hashes := make([]string, 0, len(refs))
	for _, ref := range refs {
		hashes = append(hashes, "evidence_hash_"+point13ValBTokenForEvidenceRef(ref))
	}
	return hashes
}

func point13ValBRecomputeLedgerBindingHash(model *Point13ValBFoundation) {
	model.PilotEvidenceOperationLedger.LedgerBindingHash = point13ValBComputedLedgerBindingHash(model.PilotEvidenceOperationLedger)
}

func point13ValBConfiguredLedgerEntry(operationType string, dependency Point13ValBDependencySnapshot) Point13ValBOperationLedgerEntry {
	entry := Point13ValBOperationLedgerEntry{
		EntryID:          "ledger_entry_point13_valb_test_001",
		OperationType:    operationType,
		OwnerRef:         point13ValBExpectedLedgerOwner(operationType, dependency),
		EvidenceRefs:     append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs...),
		EvidenceHashRefs: append([]string{}, dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactHashRefs...),
		AuditEventRef:    "audit_point13_valb_ledger_test_001",
		CandidateOnly:    true,
	}
	switch operationType {
	case point13ValBLedgerOperationCustomerArtifactReceived:
		entry.CustodyRef = dependency.ValA.CustomerIntakeEvidenceGovernance.CustodyRef
		entry.SourceRef = point13ValAEvidenceSourceCustomerUpload
	case point13ValBLedgerOperationEvidenceCandidateRegister:
		entry.SourceRef = point13ValAEvidenceSourceAuditExport
	case point13ValBLedgerOperationCustodyVerified:
		entry.CustodyRef = dependency.ValA.CustomerIntakeEvidenceGovernance.CustodyRef
	case point13ValBLedgerOperationSandboxResultRecorded:
		entry.SourceRef = point13ValAEvidenceSourceSandboxResult
	case point13ValBLedgerOperationSupportActionRecorded:
		entry.SourceRef = point13ValAEvidenceSourceSupportAttachment
	case point13ValBLedgerOperationCustomerReviewRecorded, point13ValBLedgerOperationExitEvidencePacket:
		entry.SourceRef = point13ValAEvidenceSourceAuditExport
	}
	return entry
}

func TestPoint13ValBFoundationFixtureIsolation(t *testing.T) {
	t.Run("raw production path still computes", func(t *testing.T) {
		model := uncachedActivePoint13ValBFoundation()
		payload := string(mustMarshalPoint13ValBFoundation(model))
		if model.CurrentState != Point13ValBStateActive {
			t.Fatalf("expected raw production path to compute active foundation, got %#v", model)
		}
		if strings.Contains(payload, point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected no point_13_pass token in active ValB payload, got %s", payload)
		}
	})

	t.Run("cached fixture mutation does not contaminate next clone", func(t *testing.T) {
		mutated := activePoint13ValBFoundation()
		mutated.Dependency.ValACurrentState = Point13ValAStateBlocked
		mutated.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceRefs = []string{"artifact_cross-tenant_candidate_001"}
		mutated.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceHashRefs = []string{"evidence_hash_cross-tenant_candidate_001"}
		point13ValBRecomputeLedgerBindingHash(&mutated)

		fresh := activePoint13ValBFoundation()
		if fresh.Dependency.ValACurrentState != Point13ValAStateActive {
			t.Fatalf("expected ValA dependency state to remain active on fresh clone, got %#v", fresh.Dependency)
		}
		if !point12Val0ExactStringSetMatch(
			fresh.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceRefs,
			fresh.Dependency.ValA.CustomerIntakeEvidenceGovernance.CustomerArtifactRefs,
		) {
			t.Fatalf("expected fresh ledger refs to remain canonical, got %#v", fresh.PilotEvidenceOperationLedger)
		}
	})
}

func TestPoint13ValBDependencyState(t *testing.T) {
	t.Run("valid vala dependency active", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		if model.DependencyState != Point13ValBStateActive || model.CurrentState != Point13ValBStateActive {
			t.Fatalf("expected active dependency and foundation, got %#v", model)
		}
	})

	testCases := []struct {
		name          string
		mutate        func(*Point13ValBFoundation)
		expectedState string
	}{
		{
			name: "missing vala dependency blocks",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.ValACurrentState = ""
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "vala blocked blocks valb",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.ValACurrentState = Point13ValAStateBlocked
				model.Dependency.ValA.CurrentState = Point13ValAStateBlocked
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "stale vala review required summary blocks valb",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.ValACurrentState = Point13ValAStateReviewRequired
				model.Dependency.ValA.CurrentState = Point13ValAStateReviewRequired
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "stale vala incomplete summary blocks valb",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.ValACurrentState = Point13ValAStateIncomplete
				model.Dependency.ValA.CurrentState = Point13ValAStateIncomplete
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "vala point13 pass appearance blocks",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.ValAPoint13PassSeen = true
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "local valb readiness cannot override vala failure",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.ValACurrentState = Point13ValAStateBlocked
				model.Dependency.ValA.CurrentState = Point13ValAStateBlocked
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "stale inherited point12 review requirement through vala blocks",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.InheritedPoint12CurrentState = Point12ValEStateReviewRequired
				model.Dependency.ValA.Dependency.Point12CurrentState = Point12ValEStateReviewRequired
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "inherited point12 binding mismatch through vala blocks",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.InheritedPoint12CurrentState = Point12ValEStateReviewRequired
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "padded nested vala state blocks raw-exact dependency binding",
			mutate: func(model *Point13ValBFoundation) {
				model.Dependency.ValA.CurrentState = " " + Point13ValAStateActive + " "
			},
			expectedState: Point13ValBStateBlocked,
		},
		{
			name: "tab newline inherited tenant scope blocks raw-exact dependency binding",
			mutate: func(model *Point13ValBFoundation) {
				retagged := model.Dependency.InheritedTenantScope + "\n"
				model.Dependency.InheritedTenantScope = retagged
				model.Dependency.ValA.Dependency.Point12TenantScope = retagged
			},
			expectedState: Point13ValBStateBlocked,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			tc.mutate(&model)
			model = ComputePoint13ValBFoundation(model)
			if model.DependencyState != tc.expectedState || model.CurrentState != tc.expectedState {
				t.Fatalf("expected dependency/current state %q, got %#v", tc.expectedState, model)
			}
		})
	}

	t.Run("padded vala point identity blocks exact dependency identity", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.Dependency.ValAPointID = " " + point13Val0PointID + " "
		state, reasons := point13ValBDependencyStateAndReasons(model.Dependency)
		if state != Point13ValBStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValBFoundation(model)
		if model.DependencyState != Point13ValBStateBlocked || model.CurrentState != Point13ValBStateBlocked {
			t.Fatalf("expected padded ValA point identity to block foundation, got %#v", model)
		}
	})

	t.Run("padded vala wave identity blocks exact dependency identity", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.Dependency.ValAWaveID = " " + point13ValAWaveID + " "
		state, reasons := point13ValBDependencyStateAndReasons(model.Dependency)
		if state != Point13ValBStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValBFoundation(model)
		if model.DependencyState != Point13ValBStateBlocked || model.CurrentState != Point13ValBStateBlocked {
			t.Fatalf("expected padded ValA wave identity to block foundation, got %#v", model)
		}
	})

	t.Run("padded nested vala state reports exact binding mismatch", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.Dependency.ValA.CurrentState = " " + Point13ValAStateActive + " "
		state, reasons := point13ValBDependencyStateAndReasons(model.Dependency)
		if state != Point13ValBStateBlocked || !point13Val0StringSliceContains(reasons, "vala_recomputed_snapshot_mismatch") {
			t.Fatalf("expected exact ValA recomputed snapshot mismatch, got state %q reasons %#v", state, reasons)
		}
	})

	t.Run("retagged inherited tenant scope reports exact identity invalid", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		retagged := model.Dependency.InheritedTenantScope + "\n"
		model.Dependency.InheritedTenantScope = retagged
		model.Dependency.ValA.Dependency.Point12TenantScope = retagged
		state, reasons := point13ValBDependencyStateAndReasons(model.Dependency)
		if state != Point13ValBStateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected exact dependency identity invalid reason, got state %q reasons %#v", state, reasons)
		}
	})

	t.Run("stale embedded vala val0 point12 profile mutation blocks recompute", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.Dependency.ValA.Dependency.Val0.Dependency.Point12.Dependency.Val0.Manifest.ProfileContext.CurrentProfileHash = ""
		state, reasons := point13ValBDependencyStateAndReasons(model.Dependency)
		if state != Point13ValBStateBlocked || !point13Val0StringSliceContains(reasons, "vala_recomputed_snapshot_mismatch") {
			t.Fatalf("expected exact ValA recomputed snapshot mismatch, got state %q reasons %#v", state, reasons)
		}
		model = ComputePoint13ValBFoundation(model)
		if model.DependencyState != Point13ValBStateBlocked || model.CurrentState != Point13ValBStateBlocked {
			t.Fatalf("expected stale embedded ValA profile mutation to block foundation, got %#v", model)
		}
	})
}

func TestPoint13ValBStateAggregation(t *testing.T) {
	reviewCases := []struct {
		name   string
		mutate func(*Point13ValBFoundation)
	}{
		{name: "pilot evidence operation ledger review required prevents active", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedgerState = Point13ValBStateReviewRequired
		}},
		{name: "customer review trace review required prevents active", mutate: func(model *Point13ValBFoundation) {
			model.CustomerReviewTraceState = Point13ValBStateReviewRequired
		}},
		{name: "support action trace review required prevents active", mutate: func(model *Point13ValBFoundation) {
			model.SupportActionTraceState = Point13ValBStateReviewRequired
		}},
		{name: "pilot exit evidence packet review required prevents active", mutate: func(model *Point13ValBFoundation) {
			model.PilotExitEvidencePacketState = Point13ValBStateReviewRequired
		}},
		{name: "ai evidence operation trace review required prevents active", mutate: func(model *Point13ValBFoundation) {
			model.AIEvidenceOperationTraceState = Point13ValBStateReviewRequired
		}},
		{name: "no overclaim review required prevents active", mutate: func(model *Point13ValBFoundation) {
			model.NoOverclaimState = Point13ValBStateReviewRequired
		}},
		{name: "dependency review required prevents active", mutate: func(model *Point13ValBFoundation) {
			model.DependencyState = Point13ValBStateReviewRequired
		}},
	}

	for _, tc := range reviewCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			tc.mutate(&model)
			if got := EvaluatePoint13ValBState(model); got != Point13ValBStateReviewRequired {
				t.Fatalf("expected review_required aggregation, got %q for %#v", got, model)
			}
		})
	}

	t.Run("any component blocked returns blocked", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.SupportActionTraceState = Point13ValBStateBlocked
		model.NoOverclaimState = Point13ValBStateReviewRequired
		if got := EvaluatePoint13ValBState(model); got != Point13ValBStateBlocked {
			t.Fatalf("expected blocked aggregation, got %q for %#v", got, model)
		}
	})

	t.Run("incomplete returned only when no blocked or review required exists", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.PilotExitEvidencePacketState = Point13ValBStateIncomplete
		if got := EvaluatePoint13ValBState(model); got != Point13ValBStateIncomplete {
			t.Fatalf("expected incomplete aggregation, got %q for %#v", got, model)
		}
	})

	t.Run("active only when all components active", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		if got := EvaluatePoint13ValBState(model); got != Point13ValBStateActive {
			t.Fatalf("expected active aggregation, got %q for %#v", got, model)
		}
	})
}

func TestPoint13ValBPilotEvidenceOperationLedgerState(t *testing.T) {
	t.Run("valid ledger active", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		if model.PilotEvidenceOperationLedgerState != Point13ValBStateActive {
			t.Fatalf("expected active ledger, got %#v", model)
		}
	})

	t.Run("all allowed operation taxonomy values validate", func(t *testing.T) {
		for _, operationType := range point13ValBLedgerOperationTypes() {
			model := activePoint13ValBFoundation()
			model.PilotEvidenceOperationLedger.OperationEntries = []Point13ValBOperationLedgerEntry{
				point13ValBConfiguredLedgerEntry(operationType, model.Dependency),
			}
			point13ValBRecomputeLedgerBindingHash(&model)
			model = ComputePoint13ValBFoundation(model)
			if model.PilotEvidenceOperationLedgerState != Point13ValBStateActive {
				t.Fatalf("expected operation type %q to remain active, got %#v", operationType, model)
			}
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValBFoundation)
	}{
		{name: "unsupported operation taxonomy blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].OperationType = "customer_artifact_promoted"
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "missing entry owner blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].OwnerRef = ""
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "missing evidence refs blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceRefs = nil
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "missing evidence hash refs blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceHashRefs = nil
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "missing audit event blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].AuditEventRef = ""
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "artifact hash mismatch blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceHashRefs[0] = "evidence_hash_point13_valb_other_candidate_001"
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "cross tenant evidence blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceRefs = []string{"artifact_cross_tenant_candidate_001"}
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceHashRefs = []string{"evidence_hash_cross_tenant_candidate_001"}
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "cross-tenant evidence blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceRefs = []string{"artifact_cross-tenant_candidate_001"}
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceHashRefs = []string{"evidence_hash_cross-tenant_candidate_001"}
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "cross tenant spaced evidence blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceRefs = []string{"artifact_cross tenant_candidate_001"}
			model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceHashRefs = []string{"evidence_hash_cross tenant_candidate_001"}
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "canonical mutation allowed blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].CanonicalMutationAllowed = true
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "production mutation allowed blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].ProductionMutationAllowed = true
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "pass allowed blocks", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.OperationEntries[0].PassAllowed = true
			point13ValBRecomputeLedgerBindingHash(model)
		}},
		{name: "padded ledger binding hash blocks raw-exact manifest binding", mutate: func(model *Point13ValBFoundation) {
			model.PilotEvidenceOperationLedger.LedgerBindingHash = " " + model.PilotEvidenceOperationLedger.LedgerBindingHash + " "
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			tc.mutate(&model)
			model = ComputePoint13ValBFoundation(model)
			if model.PilotEvidenceOperationLedgerState != Point13ValBStateBlocked {
				t.Fatalf("expected ledger mutation to block, got %#v", model)
			}
		})
	}

	t.Run("recomputed ledger hash after evidence hash tenant authority mutation still blocks", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.PilotEvidenceOperationLedger.TenantScope = "tenant_scope_point13_valb_wrong_001"
		model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceRefs = []string{"artifact_cross-tenant_candidate_001"}
		model.PilotEvidenceOperationLedger.OperationEntries[0].EvidenceHashRefs = []string{"evidence_hash_cross-tenant_candidate_001"}
		model.PilotEvidenceOperationLedger.OperationEntries[0].ProductionMutationAllowed = true
		point13ValBRecomputeLedgerBindingHash(&model)
		model = ComputePoint13ValBFoundation(model)
		if model.CurrentState == Point13ValBStateActive {
			t.Fatalf("expected recomputed ledger hash not to hide drift, got %#v", model)
		}
	})
}

func TestPoint13ValBCustomerReviewTraceState(t *testing.T) {
	t.Run("valid customer review trace active", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		if model.CustomerReviewTraceState != Point13ValBStateActive {
			t.Fatalf("expected active customer review trace, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValBCustomerReviewTrace)
	}{
		{name: "missing customer owner blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.CustomerOwnerRef = "" }},
		{name: "missing internal owner blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.InternalOwnerRef = "" }},
		{name: "missing support owner blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.SupportOwnerRef = "" }},
		{name: "unexpected internal owner identity blocks", mutate: func(model *Point13ValBCustomerReviewTrace) {
			model.InternalOwnerRef = "owner_point13_valb_unexpected_internal_001"
		}},
		{name: "unexpected support owner identity blocks", mutate: func(model *Point13ValBCustomerReviewTrace) {
			model.SupportOwnerRef = "owner_point13_valb_unexpected_support_001"
		}},
		{name: "missing audit event blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.AuditEventRefs = nil }},
		{name: "missing evidence refs blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.ReviewedEvidenceRefs = nil }},
		{name: "missing evidence hash refs blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.ReviewedEvidenceHashRefs = nil }},
		{name: "production approval wording blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.FinalCustomerStatement = "production approved" }},
		{name: "deployment approval wording blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.FinalCustomerStatement = "deployment approved" }},
		{name: "compliance claim wording blocks", mutate: func(model *Point13ValBCustomerReviewTrace) { model.FinalCustomerStatement = "compliance guaranteed" }},
		{name: "review evidence hash drift blocks", mutate: func(model *Point13ValBCustomerReviewTrace) {
			model.ReviewedEvidenceHashRefs[0] = "evidence_hash_point13_valb_other_candidate_001"
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			tc.mutate(&model.CustomerReviewTrace)
			model = ComputePoint13ValBFoundation(model)
			if model.CustomerReviewTraceState != Point13ValBStateBlocked {
				t.Fatalf("expected customer review trace mutation to block, got %#v", model)
			}
		})
	}

	t.Run("safe operational only statement passes", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.CustomerReviewTrace.FinalCustomerStatement = "customer review trace supports operational readiness review"
		model = ComputePoint13ValBFoundation(model)
		if model.CustomerReviewTraceState != Point13ValBStateActive {
			t.Fatalf("expected safe operational-only wording to remain active, got %#v", model)
		}
	})
}

func TestPoint13ValBSupportActionTraceState(t *testing.T) {
	t.Run("valid support trace active", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		if model.SupportActionTraceState != Point13ValBStateActive {
			t.Fatalf("expected active support trace, got %#v", model)
		}
	})

	t.Run("all allowed support action taxonomy values validate", func(t *testing.T) {
		for _, actionType := range point13ValBSupportActionTypes() {
			model := activePoint13ValBFoundation()
			model.SupportActionTrace.SupportActionRefs = []string{"support_action_point13_valb_taxonomy_001"}
			model.SupportActionTrace.SupportActionTypes = []string{actionType}
			model.SupportActionTrace.AuditEventRefs = []string{"audit_point13_valb_support_taxonomy_001"}
			model = ComputePoint13ValBFoundation(model)
			if model.SupportActionTraceState != Point13ValBStateActive {
				t.Fatalf("expected support action type %q to remain active, got %#v", actionType, model)
			}
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValBSupportActionTrace)
	}{
		{name: "unsupported support action blocks", mutate: func(model *Point13ValBSupportActionTrace) {
			model.SupportActionTypes[0] = "production_approval"
		}},
		{name: "missing support owner blocks", mutate: func(model *Point13ValBSupportActionTrace) { model.SupportOwnerRef = "" }},
		{name: "missing audit event blocks", mutate: func(model *Point13ValBSupportActionTrace) { model.AuditEventRefs = nil }},
		{name: "support canonical mutation blocks", mutate: func(model *Point13ValBSupportActionTrace) { model.SupportCanMutateCanonicalEvidence = true }},
		{name: "support core decision override blocks", mutate: func(model *Point13ValBSupportActionTrace) { model.SupportCanOverrideCoreDecision = true }},
		{name: "support production approval blocks", mutate: func(model *Point13ValBSupportActionTrace) { model.SupportCanApproveProduction = true }},
		{name: "audit event mutation blocks", mutate: func(model *Point13ValBSupportActionTrace) { model.AuditEventRefs[0] = "bad_ref" }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			tc.mutate(&model.SupportActionTrace)
			model = ComputePoint13ValBFoundation(model)
			if model.SupportActionTraceState != Point13ValBStateBlocked {
				t.Fatalf("expected support trace mutation to block, got %#v", model)
			}
		})
	}
}

func TestPoint13ValBPilotExitEvidencePacketState(t *testing.T) {
	t.Run("valid exit evidence packet active as operational readiness only", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		payload := string(mustMarshalPoint13ValBFoundation(model))
		if model.PilotExitEvidencePacketState != Point13ValBStateActive {
			t.Fatalf("expected active exit evidence packet, got %#v", model)
		}
		if strings.Contains(payload, point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected no point_13_pass token in active ValB payload, got %s", payload)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13ValBPilotExitEvidencePacket, *Point13ValBAIEvidenceOperationTrace)
	}{
		{name: "unresolved blockers block", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.UnresolvedBlockers = []string{"open_blocker_point13_valb_001"}
		}},
		{name: "missing customer review trace ref blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.CustomerReviewTraceRef = ""
		}},
		{name: "missing support trace ref blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.SupportTraceRef = ""
		}},
		{name: "missing operation ledger ref blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.OperationLedgerRef = ""
		}},
		{name: "evidence hash mismatch blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.EvidenceHashRefs[0] = "evidence_hash_point13_valb_other_candidate_001"
		}},
		{name: "forbidden wording blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.SafeCustomerStatement = "production approved"
		}},
		{name: "production approval flag blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.NoProductionApproval = false
		}},
		{name: "deployment approval flag blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.NoDeploymentApproval = false
		}},
		{name: "compliance guarantee flag blocks", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.NoComplianceGuarantee = false
		}},
		{name: "point13 pass cannot appear", mutate: func(model *Point13ValBPilotExitEvidencePacket, _ *Point13ValBAIEvidenceOperationTrace) {
			model.NoPoint13Pass = false
			model.SafeCustomerStatement = point13Val0BlockedPoint13PassToken
		}},
		{name: "ai trace cannot satisfy exit packet by itself", mutate: func(model *Point13ValBPilotExitEvidencePacket, ai *Point13ValBAIEvidenceOperationTrace) {
			model.EvidenceRefs = []string{ai.EvidenceCandidateRef}
			model.EvidenceHashRefs = []string{"evidence_hash_" + point13ValBTokenForEvidenceRef(ai.EvidenceCandidateRef)}
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			tc.mutate(&model.PilotExitEvidencePacket, &model.AIEvidenceOperationTrace)
			model = ComputePoint13ValBFoundation(model)
			if model.PilotExitEvidencePacketState != Point13ValBStateBlocked {
				t.Fatalf("expected exit evidence packet mutation to block, got %#v", model)
			}
		})
	}
}

func TestPoint13ValBAIEvidenceOperationTraceState(t *testing.T) {
	t.Run("all allowed ai output types remain advisory when authority flags false", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValBFoundation()
			model.AIEvidenceOperationTrace.AIOutputType = outputType
			model = ComputePoint13ValBFoundation(model)
			if model.AIEvidenceOperationTraceState != Point13ValBStateActive {
				t.Fatalf("expected allowed AI output type %q to remain active advisory candidate, got %#v", outputType, model)
			}
		}
	})

	t.Run("all allowed ai output types block on deployment authorized", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValBFoundation()
			model.AIEvidenceOperationTrace.AIOutputType = outputType
			model.AIEvidenceOperationTrace.DeploymentAuthorized = true
			model = ComputePoint13ValBFoundation(model)
			if model.AIEvidenceOperationTraceState != Point13ValBStateBlocked {
				t.Fatalf("expected deployment authority on %q to block, got %#v", outputType, model)
			}
		}
	})

	t.Run("all allowed ai output types block on production readiness claimed", func(t *testing.T) {
		for _, outputType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13ValBFoundation()
			model.AIEvidenceOperationTrace.AIOutputType = outputType
			model.AIEvidenceOperationTrace.ProductionReadinessClaimed = true
			model = ComputePoint13ValBFoundation(model)
			if model.AIEvidenceOperationTraceState != Point13ValBStateBlocked {
				t.Fatalf("expected production readiness claim on %q to block, got %#v", outputType, model)
			}
		}
	})

	t.Run("all blocked ai taxonomy values rejected", func(t *testing.T) {
		for _, outputType := range point12Val0BlockedAIEvidenceCandidateTypes() {
			model := activePoint13ValBFoundation()
			model.AIEvidenceOperationTrace.AIOutputType = outputType
			model = ComputePoint13ValBFoundation(model)
			if model.AIEvidenceOperationTraceState != Point13ValBStateBlocked {
				t.Fatalf("expected blocked AI output type %q to block, got %#v", outputType, model)
			}
		}
	})

	testCases := []struct {
		name          string
		mutate        func(*Point13ValBAIEvidenceOperationTrace)
		expectedState string
	}{
		{name: "ai approval request is not approval", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.AIOutputType = "AI_APPROVAL_REQUEST"
			model.ApprovalGranted = true
		}, expectedState: Point13ValBStateBlocked},
		{name: "ai patch proposal is not deployment", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.AIOutputType = "AI_PATCH_PROPOSAL"
			model.DeploymentAuthorized = true
		}, expectedState: Point13ValBStateBlocked},
		{name: "ai sandbox result is not production readiness", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.AIOutputType = "AI_SANDBOX_RESULT"
			model.ProductionReadinessClaimed = true
		}, expectedState: Point13ValBStateBlocked},
		{name: "external api allowed without governance event blocks", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.ExternalAPIAllowed = true
		}, expectedState: Point13ValBStateBlocked},
		{name: "external api allowed with governance event requires review", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.ExternalAPIAllowed = true
			model.ExternalAPIGovernanceEventRef = "governance_event_point13_valb_001"
		}, expectedState: Point13ValBStateReviewRequired},
		{name: "permission manifest hash mutation blocks", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.PermissionManifestHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}, expectedState: Point13ValBStateBlocked},
		{name: "input evidence hash refs mutation blocks", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.InputEvidenceHashRefs[0] = "evidence_hash_point13_valb_other_candidate_001"
		}, expectedState: Point13ValBStateBlocked},
		{name: "tenant scope mutation blocks", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.TenantScope = "tenant_scope_point13_valb_wrong_001"
		}, expectedState: Point13ValBStateBlocked},
		{name: "model or rule version ref mutation blocks", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.ModelOrRuleVersionRef = "model_version_point13_valb_wrong_001"
		}, expectedState: Point13ValBStateBlocked},
		{name: "audit event ref mutation blocks", mutate: func(model *Point13ValBAIEvidenceOperationTrace) {
			model.AuditEventRef = "audit_point13_valb_wrong_001"
		}, expectedState: Point13ValBStateBlocked},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			tc.mutate(&model.AIEvidenceOperationTrace)
			model = ComputePoint13ValBFoundation(model)
			if model.AIEvidenceOperationTraceState != tc.expectedState {
				t.Fatalf("expected AI trace state %q, got %#v", tc.expectedState, model)
			}
		})
	}
}

func TestPoint13ValBNoOverclaimTraceState(t *testing.T) {
	for _, phrase := range point13Val0ForbiddenClaims() {
		t.Run("forbidden wording blocks "+phrase, func(t *testing.T) {
			model := activePoint13ValBFoundation()
			model.NoOverclaimTrace.ObservedCustomerTexts = []string{phrase}
			model = ComputePoint13ValBFoundation(model)
			if model.NoOverclaimState != Point13ValBStateBlocked {
				t.Fatalf("expected forbidden phrase %q to block, got %#v", phrase, model)
			}
		})
	}

	t.Run("safe wording remains allowed", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.NoOverclaimTrace.ObservedCustomerTexts = append([]string{}, point13ValBAllowedSafeWording()...)
		model.NoOverclaimTrace.ObservedSupportTexts = []string{"support action trace"}
		model.NoOverclaimTrace.ObservedExitPacketTexts = []string{"operational readiness packet"}
		model = ComputePoint13ValBFoundation(model)
		if model.NoOverclaimState != Point13ValBStateActive {
			t.Fatalf("expected safe wording to remain allowed, got %#v", model)
		}
	})

	t.Run("forbidden wording cannot be laundered through allowed list", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.NoOverclaimTrace.AllowedSafeWording = []string{"public badge"}
		model = ComputePoint13ValBFoundation(model)
		if model.NoOverclaimState != Point13ValBStateBlocked || model.CurrentState != Point13ValBStateBlocked {
			t.Fatalf("expected forbidden allowed wording list mutation to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13ValBStateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("forbidden wording allowed only in internal blocked diagnostics", func(t *testing.T) {
		model := activePoint13ValBFoundation()
		model.NoOverclaimTrace.InternalDiagnosticTexts = []string{"production approved"}
		model.NoOverclaimTrace.InternalDiagnosticsClassifiedBlocked = true
		model = ComputePoint13ValBFoundation(model)
		if model.NoOverclaimState != Point13ValBStateActive {
			t.Fatalf("expected blocked internal diagnostics classification to preserve active state, got %#v", model)
		}
	})
}
