package formal

import (
	"sync"
	"testing"
)

var (
	point15ValCFoundationOnce sync.Once
	point15ValCFoundationBase Point15ValCEnforcementBoundaryFoundation
)

func point15ValCCloneStrings(values []string) []string {
	return append([]string(nil), values...)
}

func clonePoint15ValCFoundation(model Point15ValCEnforcementBoundaryFoundation) Point15ValCEnforcementBoundaryFoundation {
	model.BlockingReasons = point15ValCCloneStrings(model.BlockingReasons)
	model.ReviewPrerequisites = point15ValCCloneStrings(model.ReviewPrerequisites)
	model.Dependency.ReviewPrerequisites = point15ValCCloneStrings(model.Dependency.ReviewPrerequisites)
	model.Dependency.Point15ValB = clonePoint15ValBFoundation(model.Dependency.Point15ValB)
	model.NoOverclaimGuard.ObservedTexts = point15ValCCloneStrings(model.NoOverclaimGuard.ObservedTexts)
	model.NoOverclaimGuard.InternalDiagnosticTexts = point15ValCCloneStrings(model.NoOverclaimGuard.InternalDiagnosticTexts)
	model.NoOverclaimGuard.AllowedSafeWording = point15ValCCloneStrings(model.NoOverclaimGuard.AllowedSafeWording)
	model.NoOverclaimGuard.BlockedWording = point15ValCCloneStrings(model.NoOverclaimGuard.BlockedWording)
	return model
}

func point15ValCValidFoundationModel() Point15ValCEnforcementBoundaryFoundation {
	point15ValCFoundationOnce.Do(func() {
		point15ValCFoundationBase = Point15ValCFoundationModel()
	})
	return clonePoint15ValCFoundation(point15ValCFoundationBase)
}

func point15ValCValidDependencyModel() Point15ValCDependencySnapshot {
	return point15ValCValidFoundationModel().Dependency
}

func point15ValCValidEnforcementActionModel() Point15ValCEnforcementAction {
	return point15ValCValidFoundationModel().EnforcementAction
}

func point15ValCValidLifecycleModel() Point15ValCEvidenceLifecycleBoundary {
	return point15ValCValidFoundationModel().EvidenceLifecycle
}

func point15ValCValidRevocationModel() Point15ValCRevocationBoundary {
	return point15ValCValidFoundationModel().RevocationBoundary
}

func point15ValCValidExpiryModel() Point15ValCExpiryBoundary {
	return point15ValCValidFoundationModel().ExpiryBoundary
}

func point15ValCValidSupersessionModel() Point15ValCSupersessionBoundary {
	return point15ValCValidFoundationModel().SupersessionBoundary
}

func point15ValCValidReplayHistoryModel() Point15ValCReplayProofHistoryBoundary {
	return point15ValCValidFoundationModel().ReplayProofHistory
}

func point15ValCValidTimestampModel() Point15ValCTimestampDiscipline {
	return point15ValCValidFoundationModel().TimestampDiscipline
}

func point15ValCValidTenantModel() Point15ValCTenantBoundary {
	return point15ValCValidFoundationModel().TenantBoundary
}

func point15ValCValidAuthorityModel() Point15ValCAuthorityBoundary {
	return point15ValCValidFoundationModel().AuthorityBoundary
}

func point15ValCValidNoOverclaimModel() Point15ValCNoOverclaimGuard {
	return point15ValCValidFoundationModel().NoOverclaimGuard
}

func point15ValCRevokedFoundationModel() Point15ValCEnforcementBoundaryFoundation {
	model := point15ValCValidFoundationModel()
	model.EnforcementAction.EnforcementAction = point15ValCActionBlocked
	model.EnforcementAction.EnforcementReason = point15ValCReasonRevoked
	model.EnforcementAction.TargetState = Point15Val0StateBlocked
	model.EnforcementAction.DowngradeOutcome = point15Val0DowngradeBlocked
	model.EnforcementAction.RetainsActiveClosure = false
	model.EvidenceLifecycle.LifecycleStatus = point15ValCLifecycleRevoked
	model.EvidenceLifecycle.LifecycleReason = point15ValCReasonRevoked
	model.RevocationBoundary.RevocationPresent = true
	model.RevocationBoundary.RevocationSourceRef = "revocation_source_point15_valc_001"
	model.RevocationBoundary.RevocationReceivedAt = "2026-05-07T09:00:00Z"
	model.RevocationBoundary.RevocationValidatedAt = "2026-05-07T09:05:00Z"
	model.RevocationBoundary.RevocationTimeSource = point14Val0TimeSourceServerUTC
	model.TimestampDiscipline.RevocationPresent = true
	model.TimestampDiscipline.ReceivedAt = model.RevocationBoundary.RevocationReceivedAt
	model.TimestampDiscipline.ValidatedAt = model.RevocationBoundary.RevocationValidatedAt
	model.TimestampDiscipline.EnforcedAt = "2026-05-07T09:06:00Z"
	model.TimestampDiscipline.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
	return model
}

func point15ValCExpiredFoundationModel() Point15ValCEnforcementBoundaryFoundation {
	model := point15ValCValidFoundationModel()
	model.EnforcementAction.EnforcementAction = point15ValCActionBlocked
	model.EnforcementAction.EnforcementReason = point15ValCReasonExpired
	model.EnforcementAction.TargetState = Point15Val0StateBlocked
	model.EnforcementAction.DowngradeOutcome = point15Val0DowngradeBlocked
	model.EnforcementAction.RetainsActiveClosure = false
	model.EvidenceLifecycle.LifecycleStatus = point15ValCLifecycleExpired
	model.EvidenceLifecycle.LifecycleReason = point15ValCReasonExpired
	model.ExpiryBoundary.ExpiresAt = "2026-05-07T09:00:00Z"
	model.ExpiryBoundary.EvaluatedAt = "2026-05-07T09:05:00Z"
	model.ExpiryBoundary.ExpiryEnforced = true
	model.TimestampDiscipline.ExpiryEnforced = true
	model.TimestampDiscipline.EnforcedAt = "2026-05-07T09:06:00Z"
	model.TimestampDiscipline.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
	return model
}

func point15ValCSupersededFoundationModel(withLineage bool) Point15ValCEnforcementBoundaryFoundation {
	model := point15ValCValidFoundationModel()
	model.EvidenceLifecycle.LifecycleStatus = point15ValCLifecycleSuperseded
	model.EvidenceLifecycle.PreviousEvidenceRef = "evidence_point15_valc_old"
	model.EvidenceLifecycle.ReplacementEvidenceRef = "evidence_point15_valc_new"
	model.SupersessionBoundary.SupersessionPresent = true
	model.SupersessionBoundary.OldEvidenceRef = "evidence_point15_valc_old"
	model.SupersessionBoundary.NewEvidenceRef = "evidence_point15_valc_new"
	model.SupersessionBoundary.PriorHash = "hash_point15_valc_old"
	model.SupersessionBoundary.ReplacementHash = "hash_point15_valc_new"
	if withLineage {
		model.EnforcementAction.EnforcementAction = point15ValCActionPreserveReview
		model.EnforcementAction.EnforcementReason = point15ValCReasonSupersededWithLineage
		model.EnforcementAction.TargetState = Point15Val0StateReviewRequired
		model.EnforcementAction.DowngradeOutcome = point15Val0DowngradeReview
		model.EnforcementAction.RetainsActiveClosure = false
		model.EvidenceLifecycle.LifecycleReason = point15ValCReasonSupersededWithLineage
		model.EvidenceLifecycle.LineageRef = "lineage_point15_valc_001"
		model.SupersessionBoundary.LineageRef = "lineage_point15_valc_001"
		model.TimestampDiscipline.SupersessionPresent = true
		model.TimestampDiscipline.EnforcedAt = "2026-05-07T09:06:00Z"
		model.TimestampDiscipline.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
		return model
	}
	model.EnforcementAction.EnforcementAction = point15ValCActionPreserveBlock
	model.EnforcementAction.EnforcementReason = point15ValCReasonSupersededNoLineage
	model.EnforcementAction.TargetState = Point15Val0StateBlocked
	model.EnforcementAction.DowngradeOutcome = point15Val0DowngradeBlocked
	model.EnforcementAction.RetainsActiveClosure = false
	model.EvidenceLifecycle.LifecycleReason = point15ValCReasonSupersededNoLineage
	model.TimestampDiscipline.SupersessionPresent = true
	model.TimestampDiscipline.EnforcedAt = "2026-05-07T09:06:00Z"
	model.TimestampDiscipline.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
	return model
}

func TestPoint15ValCDependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValCDependencySnapshot)
		want   string
	}{
		{"active when valb clean", func(*Point15ValCDependencySnapshot) {}, Point15ValCStateActive},
		{"blocks when valb missing", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBCurrentState = ""
		}, Point15ValCStateBlocked},
		{"blocks when valb blocked", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBCurrentState = Point15ValBStateBlocked
		}, Point15ValCStateBlocked},
		{"blocks when valb review required", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBCurrentState = Point15ValBStateReviewRequired
		}, Point15ValCStateBlocked},
		{"blocks when valb incomplete", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBCurrentState = Point15ValBStateIncomplete
		}, Point15ValCStateBlocked},
		{"blocks when valb not merged", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBMerged = false
		}, Point15ValCStateBlocked},
		{"blocks when valb ci not green", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBCIGreen = false
		}, Point15ValCStateBlocked},
		{"blocks when valb not reviewed on main", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBReviewedOnMain = false
		}, Point15ValCStateBlocked},
		{"blocks on point15 pass token", func(model *Point15ValCDependencySnapshot) {
			model.Point15PassSeen = true
		}, Point15ValCStateBlocked},
		{"blocks when computed provenance mismatches upstream", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBComputedFromUpstream = false
		}, Point15ValCStateBlocked},
		{"blocks padded valb current state raw exact", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValBCurrentState = " " + Point15ValBStateActive + " "
			model.Point15ValB.CurrentState = model.Point15ValBCurrentState
		}, Point15ValCStateBlocked},
		{"blocks tab newline inherited vala current state raw exact", func(model *Point15ValCDependencySnapshot) {
			model.InheritedPoint15ValACurrentState = "\t" + Point15ValAStateActive + "\n"
			model.Point15ValB.Dependency.Point15ValA.CurrentState = model.InheritedPoint15ValACurrentState
		}, Point15ValCStateBlocked},
		{"blocks padded inherited point14 pass state raw exact", func(model *Point15ValCDependencySnapshot) {
			model.InheritedPoint14ValECurrentState = " " + Point14ValEStatePassConfirmed + " "
		}, Point15ValCStateBlocked},
		{"blocks padded inherited tenant scope raw exact", func(model *Point15ValCDependencySnapshot) {
			model.InheritedTenantScope = " " + model.InheritedTenantScope + " "
			model.Point15ValB.Dependency.InheritedTenantScope = model.InheritedTenantScope
		}, Point15ValCStateBlocked},
		{"blocks stale embedded valb no-overclaim blocked ledger mutation", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValB.NoOverclaimGuard.BlockedWording = append(model.Point15ValB.NoOverclaimGuard.BlockedWording, "validated revalidation schedule")
		}, Point15ValCStateBlocked},
		{"blocks stale embedded valb split no-overclaim wording", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValB.NoOverclaimGuard.ObservedTexts = []string{"continuous assurance", "guaranteed"}
		}, Point15ValCStateBlocked},
		{"blocks stale embedded valb confusable no-overclaim wording", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValB.NoOverclaimGuard.ObservedTexts = []string{"production appro\u03bded"}
		}, Point15ValCStateBlocked},
		{"blocks stale embedded vala no-overclaim disclaimer mutation", func(model *Point15ValCDependencySnapshot) {
			model.Point15ValB.Dependency.Point15ValA.NoOverclaimGuard.TriggerDisclaimer = " " + point15ValATriggerDisclaimer + " "
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValCValidDependencyModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValCDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCEnforcementActionState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCEnforcementAction
		want  string
	}{
		{"no action active", func() Point15ValCEnforcementAction {
			return point15ValCValidEnforcementActionModel()
		}, Point15ValCStateActive},
		{"expired maps to blocked", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementAction = point15ValCActionBlocked
			model.EnforcementReason = point15ValCReasonExpired
			model.TargetState = Point15Val0StateBlocked
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			return model
		}, Point15ValCStateBlocked},
		{"revoked maps to blocked", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementAction = point15ValCActionBlocked
			model.EnforcementReason = point15ValCReasonRevoked
			model.TargetState = Point15Val0StateBlocked
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			return model
		}, Point15ValCStateBlocked},
		{"stale maps to review", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementAction = point15ValCActionReview
			model.EnforcementReason = point15ValCReasonStale
			model.TargetState = Point15Val0StateReviewRequired
			model.DowngradeOutcome = point15Val0DowngradeReview
			model.RetainsActiveClosure = false
			return model
		}, Point15ValCStateReviewRequired},
		{"missing freshness proof non decisive incomplete", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementAction = point15ValCActionIncomplete
			model.EnforcementReason = point15ValCReasonMissing
			model.TargetState = Point15Val0StateIncomplete
			model.DowngradeOutcome = point15Val0DowngradeIncomplete
			model.RetainsActiveClosure = false
			return model
		}, Point15ValCStateIncomplete},
		{"missing freshness proof decisive blocks", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementAction = point15ValCActionBlocked
			model.EnforcementReason = point15ValCReasonMissing
			model.ReasonDecisive = true
			model.TargetState = Point15Val0StateBlocked
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			return model
		}, Point15ValCStateBlocked},
		{"tampered maps to quarantine blocked", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementAction = point15ValCActionQuarantine
			model.EnforcementReason = point15ValCReasonTampered
			model.TargetState = Point15Val0StateBlocked
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			return model
		}, Point15ValCStateBlocked},
		{"pass preserving alias blocks", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementAction = point15ValCActionReview
			model.EnforcementReason = point15ValCReasonStale
			model.TargetState = Point15Val0StateReviewRequired
			model.DowngradeOutcome = point15Val0DowngradeReview
			model.RetainsPass = true
			model.RetainsActiveClosure = false
			return model
		}, Point15ValCStateBlocked},
		{"no action rejects whitespace-only reason raw exact", func() Point15ValCEnforcementAction {
			model := point15ValCValidEnforcementActionModel()
			model.EnforcementReason = " "
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCEnforcementActionState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCEvidenceLifecycleBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCEvidenceLifecycleBoundary
		want  string
	}{
		{"active bound active", func() Point15ValCEvidenceLifecycleBoundary {
			return point15ValCValidLifecycleModel()
		}, Point15ValCStateActive},
		{"expired evidence blocked", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.LifecycleStatus = point15ValCLifecycleExpired
			model.LifecycleReason = point15ValCReasonExpired
			return model
		}, Point15ValCStateBlocked},
		{"revoked evidence blocked", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.LifecycleStatus = point15ValCLifecycleRevoked
			model.LifecycleReason = point15ValCReasonRevoked
			return model
		}, Point15ValCStateBlocked},
		{"tampered evidence blocked", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.LifecycleStatus = point15ValCLifecycleTampered
			model.LifecycleReason = point15ValCReasonTampered
			return model
		}, Point15ValCStateBlocked},
		{"superseded without lineage blocks", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.LifecycleStatus = point15ValCLifecycleSuperseded
			model.LifecycleReason = point15ValCReasonSupersededNoLineage
			model.PreviousEvidenceRef = "evidence_point15_valc_old"
			model.ReplacementEvidenceRef = "evidence_point15_valc_new"
			return model
		}, Point15ValCStateBlocked},
		{"superseded with lineage review", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.LifecycleStatus = point15ValCLifecycleSuperseded
			model.LifecycleReason = point15ValCReasonSupersededWithLineage
			model.PreviousEvidenceRef = "evidence_point15_valc_old"
			model.ReplacementEvidenceRef = "evidence_point15_valc_new"
			model.LineageRef = "lineage_point15_valc_001"
			return model
		}, Point15ValCStateReviewRequired},
		{"canonical mutation attempted blocks", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.CanonicalMutationAttempted = true
			return model
		}, Point15ValCStateBlocked},
		{"history must be preserved", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.HistoryPreserved = false
			return model
		}, Point15ValCStateBlocked},
		{"active lifecycle rejects whitespace-only previous ref raw exact", func() Point15ValCEvidenceLifecycleBoundary {
			model := point15ValCValidLifecycleModel()
			model.PreviousEvidenceRef = "\t"
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCEvidenceLifecycleBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCRevocationBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCRevocationBoundary
		want  string
	}{
		{"no revocation active", func() Point15ValCRevocationBoundary {
			return point15ValCValidRevocationModel()
		}, Point15ValCStateActive},
		{"revocation input blocks active closure", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.RevocationPresent = true
			model.RevocationSourceRef = "revocation_source_point15_valc_001"
			model.RevocationReceivedAt = "2026-05-07T09:00:00Z"
			model.RevocationValidatedAt = "2026-05-07T09:05:00Z"
			model.RevocationTimeSource = point14Val0TimeSourceServerUTC
			return model
		}, Point15ValCStateBlocked},
		{"revocation source cannot become canonical authority", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.SourceAuthorityGranted = true
			return model
		}, Point15ValCStateBlocked},
		{"auto revoked blocks", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.AutoRevoked = true
			return model
		}, Point15ValCStateBlocked},
		{"auto published blocks", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.AutoPublished = true
			return model
		}, Point15ValCStateBlocked},
		{"revocation without validated at incomplete", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.RevocationPresent = true
			model.RevocationSourceRef = "revocation_source_point15_valc_001"
			model.RevocationReceivedAt = "2026-05-07T09:00:00Z"
			model.RevocationTimeSource = point14Val0TimeSourceServerUTC
			return model
		}, Point15ValCStateIncomplete},
		{"client local revocation time blocks", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.RevocationPresent = true
			model.RevocationSourceRef = "revocation_source_point15_valc_001"
			model.RevocationReceivedAt = "2026-05-07T09:00:00Z"
			model.RevocationValidatedAt = "2026-05-07T09:05:00Z"
			model.RevocationTimeSource = point14Val0TimeSourceClientLocal
			return model
		}, Point15ValCStateBlocked},
		{"revocation history must remain visible", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.HistoryPreserved = false
			return model
		}, Point15ValCStateBlocked},
		{"no revocation rejects whitespace-only source ref raw exact", func() Point15ValCRevocationBoundary {
			model := point15ValCValidRevocationModel()
			model.RevocationSourceRef = " "
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCRevocationBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCExpiryBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCExpiryBoundary
		want  string
	}{
		{"future expiry active", func() Point15ValCExpiryBoundary {
			return point15ValCValidExpiryModel()
		}, Point15ValCStateActive},
		{"expires before evaluated blocks", func() Point15ValCExpiryBoundary {
			model := point15ValCValidExpiryModel()
			model.ExpiresAt = "2026-05-07T09:00:00Z"
			model.EvaluatedAt = "2026-05-07T09:05:00Z"
			return model
		}, Point15ValCStateBlocked},
		{"missing expires at incomplete", func() Point15ValCExpiryBoundary {
			model := point15ValCValidExpiryModel()
			model.ExpiresAt = ""
			return model
		}, Point15ValCStateIncomplete},
		{"client local expiry time blocks", func() Point15ValCExpiryBoundary {
			model := point15ValCValidExpiryModel()
			model.ExpiryTimeSource = point14Val0TimeSourceClientLocal
			return model
		}, Point15ValCStateBlocked},
		{"expiry cannot hide history", func() Point15ValCExpiryBoundary {
			model := point15ValCValidExpiryModel()
			model.ExpiryHistoryPreserved = false
			return model
		}, Point15ValCStateBlocked},
		{"whitespace expiry timestamp blocks as malformed raw value", func() Point15ValCExpiryBoundary {
			model := point15ValCValidExpiryModel()
			model.ExpiresAt = " "
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCExpiryBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCSupersessionBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCSupersessionBoundary
		want  string
	}{
		{"not present active", func() Point15ValCSupersessionBoundary {
			return point15ValCValidSupersessionModel()
		}, Point15ValCStateActive},
		{"missing lineage blocks", func() Point15ValCSupersessionBoundary {
			model := point15ValCValidSupersessionModel()
			model.SupersessionPresent = true
			model.OldEvidenceRef = "evidence_point15_valc_old"
			model.NewEvidenceRef = "evidence_point15_valc_new"
			model.PriorHash = "hash_point15_valc_old"
			model.ReplacementHash = "hash_point15_valc_new"
			return model
		}, Point15ValCStateBlocked},
		{"silent replacement blocks", func() Point15ValCSupersessionBoundary {
			model := point15ValCValidSupersessionModel()
			model.SilentReplacementDetected = true
			return model
		}, Point15ValCStateBlocked},
		{"valid lineage review required", func() Point15ValCSupersessionBoundary {
			model := point15ValCValidSupersessionModel()
			model.SupersessionPresent = true
			model.OldEvidenceRef = "evidence_point15_valc_old"
			model.NewEvidenceRef = "evidence_point15_valc_new"
			model.PriorHash = "hash_point15_valc_old"
			model.ReplacementHash = "hash_point15_valc_new"
			model.LineageRef = "lineage_point15_valc_001"
			return model
		}, Point15ValCStateReviewRequired},
		{"old and new ref must differ", func() Point15ValCSupersessionBoundary {
			model := point15ValCValidSupersessionModel()
			model.SupersessionPresent = true
			model.OldEvidenceRef = "evidence_point15_valc_old"
			model.NewEvidenceRef = "evidence_point15_valc_old"
			model.PriorHash = "hash_point15_valc_old"
			model.ReplacementHash = "hash_point15_valc_new"
			model.LineageRef = "lineage_point15_valc_001"
			return model
		}, Point15ValCStateBlocked},
		{"supersession cannot auto publish or approve", func() Point15ValCSupersessionBoundary {
			model := point15ValCValidSupersessionModel()
			model.AutoPublished = true
			return model
		}, Point15ValCStateBlocked},
		{"not present rejects whitespace-only old evidence ref raw exact", func() Point15ValCSupersessionBoundary {
			model := point15ValCValidSupersessionModel()
			model.OldEvidenceRef = "\n"
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCSupersessionBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCReplayProofHistoryBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCReplayProofHistoryBoundary
		want  string
	}{
		{"history boundary active", func() Point15ValCReplayProofHistoryBoundary {
			return point15ValCValidReplayHistoryModel()
		}, Point15ValCStateActive},
		{"replay proof refs required", func() Point15ValCReplayProofHistoryBoundary {
			model := point15ValCValidReplayHistoryModel()
			model.ReplayRef = ""
			return model
		}, Point15ValCStateBlocked},
		{"decisive evidence must remain visible", func() Point15ValCReplayProofHistoryBoundary {
			model := point15ValCValidReplayHistoryModel()
			model.DecisiveEvidenceVisible = false
			return model
		}, Point15ValCStateBlocked},
		{"prior state visibility required", func() Point15ValCReplayProofHistoryBoundary {
			model := point15ValCValidReplayHistoryModel()
			model.PriorStateVisible = false
			return model
		}, Point15ValCStateBlocked},
		{"projection wording cannot strengthen claims", func() Point15ValCReplayProofHistoryBoundary {
			model := point15ValCValidReplayHistoryModel()
			model.ProjectionStrengthensClaims = true
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCReplayProofHistoryBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCTimestampDisciplineState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCTimestampDiscipline
		want  string
	}{
		{"trusted timing active", func() Point15ValCTimestampDiscipline {
			return point15ValCValidTimestampModel()
		}, Point15ValCStateActive},
		{"source event alone cannot enforce revocation", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.RevocationPresent = true
			model.SourceEventAt = "2026-05-07T09:00:00Z"
			model.SourceEventTimeSource = point14Val0TimeSourceServerUTC
			model.ReceivedAt = ""
			model.ValidatedAt = ""
			model.EnforcedAt = ""
			return model
		}, Point15ValCStateIncomplete},
		{"future enforced at blocks", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.RevocationPresent = true
			model.EnforcedAt = "2026-05-07T09:20:00Z"
			model.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
			return model
		}, Point15ValCStateBlocked},
		{"enforced before validated blocks", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.RevocationPresent = true
			model.EnforcedAt = "2026-05-07T09:04:00Z"
			model.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
			return model
		}, Point15ValCStateBlocked},
		{"validated before received review required", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.RevocationPresent = true
			model.EnforcedAt = "2026-05-07T09:06:00Z"
			model.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
			model.ValidatedAt = "2026-05-07T08:59:00Z"
			return model
		}, Point15ValCStateReviewRequired},
		{"backdated enforcement review required", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.RevocationPresent = true
			model.EnforcedAt = "2026-05-07T09:06:00Z"
			model.EnforcedAtTimeSource = point14Val0TimeSourceServerUTC
			model.EvaluatedAt = "2026-05-07T09:07:00Z"
			return model
		}, Point15ValCStateReviewRequired},
		{"untrusted time source blocks", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.RevocationPresent = true
			model.EnforcedAt = "2026-05-07T09:06:00Z"
			model.EnforcedAtTimeSource = point14Val0TimeSourceClientLocal
			return model
		}, Point15ValCStateBlocked},
		{"source event parse failure blocks", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.SourceEventAt = "not-a-time"
			model.SourceEventTimeSource = point14Val0TimeSourceServerUTC
			return model
		}, Point15ValCStateBlocked},
		{"source event untrusted time source blocks", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.SourceEventAt = "2026-05-07T09:00:00Z"
			model.SourceEventTimeSource = point14Val0TimeSourceClientLocal
			return model
		}, Point15ValCStateBlocked},
		{"whitespace source event timestamp blocks raw exact", func() Point15ValCTimestampDiscipline {
			model := point15ValCValidTimestampModel()
			model.SourceEventAt = " "
			model.SourceEventTimeSource = point14Val0TimeSourceServerUTC
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCTimestampDisciplineState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCTenantBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCTenantBoundary
		want  string
	}{
		{"matching tenant scopes active", func() Point15ValCTenantBoundary {
			return point15ValCValidTenantModel()
		}, Point15ValCStateActive},
		{"cross tenant enforcement blocks", func() Point15ValCTenantBoundary {
			model := point15ValCValidTenantModel()
			model.CrossTenantDetected = true
			return model
		}, Point15ValCStateBlocked},
		{"tenant mismatch blocks", func() Point15ValCTenantBoundary {
			model := point15ValCValidTenantModel()
			model.ReferencedTenantScope = "tenant_point15_valc_other"
			return model
		}, Point15ValCStateBlocked},
		{"missing tenant scope incomplete", func() Point15ValCTenantBoundary {
			model := point15ValCValidTenantModel()
			model.TenantScope = ""
			return model
		}, Point15ValCStateIncomplete},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCTenantBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCAuthorityBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValCAuthorityBoundary)
		want   string
	}{
		{"authority boundary active", func(*Point15ValCAuthorityBoundary) {}, Point15ValCStateActive},
		{"scheduler cannot enforce revocation or expiry", func(model *Point15ValCAuthorityBoundary) {
			model.SchedulerEnforcesBoundary = true
		}, Point15ValCStateBlocked},
		{"connector cannot restore active closure", func(model *Point15ValCAuthorityBoundary) {
			model.ConnectorRestoresActiveClosure = true
		}, Point15ValCStateBlocked},
		{"dashboard cannot suppress enforcement", func(model *Point15ValCAuthorityBoundary) {
			model.DashboardSuppressesEnforcement = true
		}, Point15ValCStateBlocked},
		{"portal projection cannot mutate enforcement", func(model *Point15ValCAuthorityBoundary) {
			model.PortalProjectionMutatesEnforcement = true
		}, Point15ValCStateBlocked},
		{"customer projection cannot mutate enforcement", func(model *Point15ValCAuthorityBoundary) {
			model.CustomerProjectionMutatesEnforcement = true
		}, Point15ValCStateBlocked},
		{"auditor projection cannot mutate enforcement", func(model *Point15ValCAuthorityBoundary) {
			model.AuditorProjectionMutatesEnforcement = true
		}, Point15ValCStateBlocked},
		{"agent recommendation cannot satisfy enforcement", func(model *Point15ValCAuthorityBoundary) {
			model.AgentSatisfiesEnforcement = true
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValCValidAuthorityModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValCAuthorityBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCNoOverclaimGuardState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValCNoOverclaimGuard
		want  string
	}{
		{"safe bounded wording passes", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.ObservedTexts = []string{"downgrade trigger detected", "enforcement preserves proof history"}
			return model
		}, Point15ValCStateActive},
		{"forbidden wording blocks", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.ObservedTexts = []string{"continuous assurance guaranteed"}
			return model
		}, Point15ValCStateBlocked},
		{"split forbidden wording blocks", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.ObservedTexts = []string{"continuous assurance", "guaranteed"}
			return model
		}, Point15ValCStateBlocked},
		{"confusable forbidden wording blocks", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.ObservedTexts = []string{"production appro\u03bded"}
			return model
		}, Point15ValCStateBlocked},
		{"classified internal blocked diagnostics allowed", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.InternalDiagnosticsClassifiedBlocked = true
			model.InternalDiagnosticTexts = []string{"guaranteed secure"}
			return model
		}, Point15ValCStateActive},
		{"unclassified internal blocked diagnostics block", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.InternalDiagnosticTexts = []string{"public badge"}
			return model
		}, Point15ValCStateBlocked},
		{"mutated allowed safe wording blocks", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.AllowedSafeWording = []string{"downgrade trigger detected"}
			return model
		}, Point15ValCStateBlocked},
		{"mutated blocked wording blocks", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.BlockedWording = []string{"continuous assurance guaranteed"}
			return model
		}, Point15ValCStateBlocked},
		{"padded enforcement disclaimer blocks raw exact", func() Point15ValCNoOverclaimGuard {
			model := point15ValCValidNoOverclaimModel()
			model.EnforcementDisclaimer = " " + point15ValCEnforcementDisclaimer + " "
			return model
		}, Point15ValCStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValCNoOverclaimGuardState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValCEnforcementBoundaryFoundationState(t *testing.T) {
	t.Run("happy path active", func(t *testing.T) {
		model := point15ValCValidFoundationModel()
		got := ComputePoint15ValCEnforcementBoundaryFoundation(model)
		if got.CurrentState != Point15ValCStateActive ||
			got.DependencyState != Point15ValCStateActive ||
			got.EnforcementActionState != Point15ValCStateActive ||
			got.EvidenceLifecycleState != Point15ValCStateActive ||
			got.RevocationBoundaryState != Point15ValCStateActive ||
			got.ExpiryBoundaryState != Point15ValCStateActive ||
			got.SupersessionState != Point15ValCStateActive ||
			got.ReplayProofHistoryState != Point15ValCStateActive ||
			got.TimestampDisciplineState != Point15ValCStateActive ||
			got.AuthorityBoundaryState != Point15ValCStateActive ||
			got.TenantBoundaryState != Point15ValCStateActive ||
			got.NoOverclaimState != Point15ValCStateActive {
			t.Fatalf("expected full point15 valc foundation active, got %#v", got)
		}
	})

	tests := []struct {
		name   string
		mutate func(*Point15ValCEnforcementBoundaryFoundation)
		assert func(*testing.T, Point15ValCEnforcementBoundaryFoundation)
	}{
		{
			name: "expired evidence cannot remain active",
			mutate: func(model *Point15ValCEnforcementBoundaryFoundation) {
				*model = point15ValCExpiredFoundationModel()
			},
			assert: func(t *testing.T, got Point15ValCEnforcementBoundaryFoundation) {
				if got.CurrentState != Point15ValCStateBlocked || got.EnforcementActionState != Point15ValCStateBlocked || got.ExpiryBoundaryState != Point15ValCStateBlocked {
					t.Fatalf("expected expired enforcement to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "revoked evidence cannot remain active",
			mutate: func(model *Point15ValCEnforcementBoundaryFoundation) {
				*model = point15ValCRevokedFoundationModel()
			},
			assert: func(t *testing.T, got Point15ValCEnforcementBoundaryFoundation) {
				if got.CurrentState != Point15ValCStateBlocked || got.EnforcementActionState != Point15ValCStateBlocked || got.RevocationBoundaryState != Point15ValCStateBlocked {
					t.Fatalf("expected revoked enforcement to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "superseded without lineage blocks",
			mutate: func(model *Point15ValCEnforcementBoundaryFoundation) {
				*model = point15ValCSupersededFoundationModel(false)
			},
			assert: func(t *testing.T, got Point15ValCEnforcementBoundaryFoundation) {
				if got.CurrentState != Point15ValCStateBlocked || got.SupersessionState != Point15ValCStateBlocked {
					t.Fatalf("expected superseded without lineage to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "superseded with lineage requires review",
			mutate: func(model *Point15ValCEnforcementBoundaryFoundation) {
				*model = point15ValCSupersededFoundationModel(true)
			},
			assert: func(t *testing.T, got Point15ValCEnforcementBoundaryFoundation) {
				if got.CurrentState != Point15ValCStateReviewRequired || got.SupersessionState != Point15ValCStateReviewRequired || got.EnforcementActionState != Point15ValCStateReviewRequired {
					t.Fatalf("expected superseded with lineage review, got %#v", got)
				}
			},
		},
		{
			name: "replay history cannot hide decisive evidence",
			mutate: func(model *Point15ValCEnforcementBoundaryFoundation) {
				model.ReplayProofHistory.DecisiveEvidenceVisible = false
			},
			assert: func(t *testing.T, got Point15ValCEnforcementBoundaryFoundation) {
				if got.CurrentState != Point15ValCStateBlocked || got.ReplayProofHistoryState != Point15ValCStateBlocked {
					t.Fatalf("expected hidden decisive evidence to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "cross tenant enforcement blocks",
			mutate: func(model *Point15ValCEnforcementBoundaryFoundation) {
				model.TenantBoundary.ReferencedTenantScope = "tenant_point15_valc_other"
			},
			assert: func(t *testing.T, got Point15ValCEnforcementBoundaryFoundation) {
				if got.CurrentState != Point15ValCStateBlocked || got.TenantBoundaryState != Point15ValCStateBlocked {
					t.Fatalf("expected cross-tenant enforcement to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "source refs must bind exactly to upstream valb and vala",
			mutate: func(model *Point15ValCEnforcementBoundaryFoundation) {
				model.EnforcementAction.SourceValBTriggerRef = "point15_valb_binding_other"
			},
			assert: func(t *testing.T, got Point15ValCEnforcementBoundaryFoundation) {
				if got.CurrentState != Point15ValCStateBlocked || got.EnforcementActionState != Point15ValCStateBlocked {
					t.Fatalf("expected source-ref mismatch to block foundation, got %#v", got)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValCValidFoundationModel()
			tc.mutate(&model)
			got := ComputePoint15ValCEnforcementBoundaryFoundation(model)
			tc.assert(t, got)
		})
	}

	t.Run("stale embedded valb no-overclaim mutation records exact dependency reason", func(t *testing.T) {
		model := point15ValCValidFoundationModel()
		model.Dependency.Point15ValB.NoOverclaimGuard.BlockedWording = append(model.Dependency.Point15ValB.NoOverclaimGuard.BlockedWording, "validated revalidation schedule")
		computed := ComputePoint15ValCEnforcementBoundaryFoundation(model)
		if computed.CurrentState != Point15ValCStateBlocked {
			t.Fatalf("expected stale embedded ValB no-overclaim mutation to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})

	t.Run("stale embedded valb split no-overclaim wording records exact dependency reason", func(t *testing.T) {
		model := point15ValCValidFoundationModel()
		model.Dependency.Point15ValB.NoOverclaimGuard.ObservedTexts = []string{"continuous assurance", "guaranteed"}
		computed := ComputePoint15ValCEnforcementBoundaryFoundation(model)
		if computed.CurrentState != Point15ValCStateBlocked {
			t.Fatalf("expected stale embedded ValB split no-overclaim wording to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})

	t.Run("stale embedded vala no-overclaim mutation records exact dependency reason", func(t *testing.T) {
		model := point15ValCValidFoundationModel()
		model.Dependency.Point15ValB.Dependency.Point15ValA.NoOverclaimGuard.TriggerDisclaimer = " " + point15ValATriggerDisclaimer + " "
		computed := ComputePoint15ValCEnforcementBoundaryFoundation(model)
		if computed.CurrentState != Point15ValCStateBlocked {
			t.Fatalf("expected stale embedded ValA no-overclaim mutation to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})
}

func TestPoint10ThroughPoint15ValCCurrentSweep(t *testing.T) {
	computed := ComputePoint15ValCEnforcementBoundaryFoundation(point15ValCValidFoundationModel())
	if computed.DependencyState != Point15ValCStateActive {
		t.Fatalf("expected dependency active, got %s", computed.DependencyState)
	}
	if computed.CurrentState != Point15ValCStateActive {
		t.Fatalf("expected current state active, got %s", computed.CurrentState)
	}
	if computed.Dependency.Point15PassSeen {
		t.Fatal("expected no point_15_pass in point15 val c sweep")
	}
}

func TestPoint15ValCCachedHelperIsolation(t *testing.T) {
	model := point15ValCValidFoundationModel()
	originalAllowed := model.NoOverclaimGuard.AllowedSafeWording[0]
	model.NoOverclaimGuard.AllowedSafeWording[0] = "mutated"

	fresh := point15ValCValidFoundationModel()
	if fresh.NoOverclaimGuard.AllowedSafeWording[0] != originalAllowed {
		t.Fatalf("expected cached point15 valc helper to return isolated copy, got %#v", fresh.NoOverclaimGuard.AllowedSafeWording)
	}
}

func TestPoint15ValCCachedHelperNestedDependencyIsolation(t *testing.T) {
	model := point15ValCValidFoundationModel()
	originalAllowed := model.Dependency.Point15ValB.NoOverclaimGuard.AllowedSafeWording[0]
	model.Dependency.Point15ValB.NoOverclaimGuard.AllowedSafeWording[0] = "mutated"

	fresh := point15ValCValidFoundationModel()
	if fresh.Dependency.Point15ValB.NoOverclaimGuard.AllowedSafeWording[0] != originalAllowed {
		t.Fatalf("expected cached point15 valc nested dependency helper to return isolated copy, got %#v", fresh.Dependency.Point15ValB.NoOverclaimGuard.AllowedSafeWording)
	}
}

func TestPoint15ValCAggregateRawExact(t *testing.T) {
	tests := []struct {
		name   string
		states []string
		want   string
	}{
		{"happy path active", []string{Point15ValCStateActive, Point15ValCStateActive}, Point15ValCStateActive},
		{"direct exploit padded active blocks", []string{Point15ValCStateActive, " " + Point15ValCStateActive + " "}, Point15ValCStateBlocked},
		{"hard invalid tab newline active blocks", []string{Point15ValCStateActive, "\t" + Point15ValCStateActive + "\n"}, Point15ValCStateBlocked},
		{"sibling review path preserved", []string{Point15ValCStateActive, Point15ValCStateReviewRequired}, Point15ValCStateReviewRequired},
		{"blocked path preserved", []string{Point15ValCStateActive, Point15ValCStateBlocked}, Point15ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := point15ValCAggregate(tc.states...); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}
