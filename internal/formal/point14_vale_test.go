package formal

import "testing"

func point14ValEValidClosureEvaluatorModel() Point14ValEClosureEvaluator {
	model := point14ValEClosureEvaluatorModel()
	model.DependencyState = Point14ValEStatePassConfirmed
	model.ValidationClosureState = Point14ValEStatePassConfirmed
	model.DisputeClosureState = Point14ValEStatePassConfirmed
	model.CorrectionPublicationClosureState = Point14ValEStatePassConfirmed
	model.TimelineProjectionClosureState = Point14ValEStatePassConfirmed
	model.AuthorityBoundaryState = Point14ValEStatePassConfirmed
	model.TenantPrivacyState = Point14ValEStatePassConfirmed
	model.TimestampIntegrityState = Point14ValEStatePassConfirmed
	model.AgentAdvisoryState = Point14ValEStatePassConfirmed
	model.NoOverclaimState = Point14ValEStatePassConfirmed
	model.CLBFinalState = Point14ValEStatePassConfirmed
	model.FinalPassAllowed = true
	return model
}

func point14ValEValidPassClosureManifestModel() Point14PassClosureManifest {
	model := point14ValEPassClosureManifestModel()
	model.DependencyGateResult = Point14ValEStatePassConfirmed
	model.ClosureEvaluatorResult = Point14ValEStatePassConfirmed
	model.ProjectionBoundaryResult = Point14ValEStatePassConfirmed
	model.NoExternalAuthorityResult = Point14ValEStatePassConfirmed
	model.NoOverclaimGrepResult = Point14ValEStatePassConfirmed
	model.TenantPrivacyResult = Point14ValEStatePassConfirmed
	model.TimestampIntegrityResult = Point14ValEStatePassConfirmed
	model.AIAgentBoundaryResult = Point14ValEStatePassConfirmed
	model.CLBResult = Point14ValEStatePassConfirmed
	model.Point14PassAllowed = true
	model.Point14PassToken = point14Val0BlockedPassToken
	return model
}

func TestPoint14ValEDependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValEDependencySnapshot)
		want   string
	}{
		{
			name:   "canonical inherited boundary snapshot stays pass confirmed",
			mutate: func(model *Point14ValEDependencySnapshot) {},
			want:   Point14ValEStatePassConfirmed,
		},
		{
			name: "missing point14 vald blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValDCurrentState = ""
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "mutated vald allowed no overclaim wording blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.NoOverclaimTimelineWording.AllowedSafeWording = append(
					model.Point14ValD.NoOverclaimTimelineWording.AllowedSafeWording,
					"timeline proves truth",
				)
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "mutated valc allowed no overclaim wording blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValC.NoOverclaimPublicationWording.AllowedSafeWording = append(
					model.Point14ValD.Dependency.Point14ValC.NoOverclaimPublicationWording.AllowedSafeWording,
					"publication proves safety",
				)
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "mutated valb allowed no overclaim wording blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValB.NoOverclaimDisputeWording.AllowedSafeWording = append(
					model.Point14ValD.Dependency.Point14ValB.NoOverclaimDisputeWording.AllowedSafeWording,
					"dispute resolved by ai",
				)
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "mutated vala allowed no overclaim wording blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValA.NoOverclaimValidationWording.AllowedSafeWording = append(
					model.Point14ValD.Dependency.Point14ValA.NoOverclaimValidationWording.AllowedSafeWording,
					"production approved",
				)
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "mutated val0 allowed no overclaim wording blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValA.Dependency.Point14Val0.NoOverclaimEcosystemWording.AllowedSafeWording = append(
					model.Point14ValD.Dependency.Point14ValA.Dependency.Point14Val0.NoOverclaimEcosystemWording.AllowedSafeWording,
					"production approved",
				)
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 current state in vald dependency blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.InheritedPoint10CurrentState += " "
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 no overclaim state in vald dependency blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.InheritedPoint10NoOverclaimState += " "
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 projection state in vald dependency blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.InheritedPoint10ProjectionState += " "
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 pass rule state in vald dependency blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.InheritedPoint10PassRuleState += " "
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "tab newline retagged embedded vald current state blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.CurrentState = "\t" + model.Point14ValD.CurrentState + "\n"
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "whitespace retagged nested point11 current state in vald dependency blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.InheritedPoint11CurrentState += " "
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "whitespace retagged point11 current state blocks raw exact",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.InheritedPoint11CurrentState += " "
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "nested vald point11 current state mismatch blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.InheritedPoint11CurrentState = Point11ValDStateReviewRequired
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "nested vald embedded point11 current state mismatch blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point11.CurrentState = Point11ValDStateReviewRequired
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "nested vald embedded point11 pass manifest current state mismatch blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point11.PassClosureManifest.CurrentState = Point11ValDPassClosureManifestStateBlocked
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "deep nested point11 final pass gate current state mismatch blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValC.Dependency.Point11.FinalPassGate.CurrentState = Point11ValDFinalPassGateStateBlocked
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "point11 current state review required blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.InheritedPoint11CurrentState = Point11ValDStateReviewRequired
				model.Point14ValD.Dependency.InheritedPoint11CurrentState = Point11ValDStateReviewRequired
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "whitespace retagged nested point11 final pass gate in vald dependency blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.InheritedPoint11FinalPassGateState += " "
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "point14 vald blocked blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValDCurrentState = Point14ValDStateBlocked
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "point14 vald review required blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValDCurrentState = Point14ValDStateReviewRequired
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "point14 vald incomplete blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValDCurrentState = Point14ValDStateIncomplete
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "missing point14 valc blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.InheritedPoint14ValCCurrentState = ""
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "missing point14 valb blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.InheritedPoint14ValBCurrentState = ""
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "missing point14 vala blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.InheritedPoint14ValACurrentState = ""
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "missing point14 val0 blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.InheritedPoint14Val0CurrentState = ""
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "nested vald embedded point11 dependency state mismatch blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point11.DependencyState = Point11ValDDependencyStateBlocked
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "deep nested valc embedded point11 val0 dependency mismatch blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValC.Dependency.Point11.Val0Dependency.CurrentState = Point11Val0StateBlocked
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "missing point13 vale blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.InheritedPoint13ValECurrentState = ""
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "premature point14 pass outside vale closure path blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14PassSeen = true
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "local vale readiness cannot override missing vald closure",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValDCurrentState = ""
				model.Point14ValDTimelineProjectionState = Point14ValDStateActive
				model.Point14ValDDisputeTimelineState = Point14ValDStateActive
				model.InheritedPoint14ValCCurrentState = Point14ValCStateActive
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "direct embedded valb state drift blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValB.CurrentState = Point14ValBStateReviewRequired
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "direct embedded vala state drift blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValA.CurrentState = Point14ValAStateReviewRequired
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "nested valc embedded vala state drift blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValA.CurrentState = Point14ValAStateReviewRequired
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "embedded valb production approval authority blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValB.AgentDisputeRecommendationBoundary.ProductionApproved = true
			},
			want: Point14ValEStateBlocked,
		},
		{
			name: "synchronized embedded valb production approval authority blocks",
			mutate: func(model *Point14ValEDependencySnapshot) {
				model.Point14ValD.Dependency.Point14ValB.AgentDisputeRecommendationBoundary.ProductionApproved = true
				model.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.AgentDisputeRecommendationBoundary.ProductionApproved = true
			},
			want: Point14ValEStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValEDependencySnapshotModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValEDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValEClosureEvaluatorState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValEClosureEvaluator)
		want   string
	}{
		{"valid full closure produces pass confirmed", func(model *Point14ValEClosureEvaluator) {}, Point14ValEStatePassConfirmed},
		{"any dependency blocked returns blocked", func(model *Point14ValEClosureEvaluator) { model.DependencyState = Point14ValEStateBlocked }, Point14ValEStateBlocked},
		{"any validation closure blocked returns blocked", func(model *Point14ValEClosureEvaluator) { model.ValidationClosureState = Point14ValEStateBlocked }, Point14ValEStateBlocked},
		{"review required prevents pass confirmed", func(model *Point14ValEClosureEvaluator) { model.DisputeClosureState = Point14ValEStateReviewRequired }, Point14ValEStateReviewRequired},
		{"incomplete prevents pass confirmed", func(model *Point14ValEClosureEvaluator) { model.DisputeClosureState = Point14ValEStateIncomplete }, Point14ValEStateIncomplete},
		{"whitespace retagged pass confirmed state blocks", func(model *Point14ValEClosureEvaluator) { model.NoOverclaimState = Point14ValEStatePassConfirmed + " " }, Point14ValEStateBlocked},
		{"tab newline retagged pass confirmed state blocks", func(model *Point14ValEClosureEvaluator) {
			model.CLBFinalState = "\t" + Point14ValEStatePassConfirmed + "\n"
		}, Point14ValEStateBlocked},
		{"no mutation paths detected false blocks", func(model *Point14ValEClosureEvaluator) { model.NoMutationPathsDetected = false }, Point14ValEStateBlocked},
		{"no external authority detected false blocks", func(model *Point14ValEClosureEvaluator) { model.NoExternalAuthorityDetected = false }, Point14ValEStateBlocked},
		{"no premature point14 pass false blocks", func(model *Point14ValEClosureEvaluator) { model.NoPrematurePoint14Pass = false }, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValEValidClosureEvaluatorModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValEClosureEvaluatorState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14PassClosureManifestState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14PassClosureManifest)
		want   string
	}{
		{"valid manifest pass confirmed", func(model *Point14PassClosureManifest) {}, Point14ValEStatePassConfirmed},
		{"missing point id blocks", func(model *Point14PassClosureManifest) { model.PointID = "" }, Point14ValEStateBlocked},
		{"wrong point id blocks", func(model *Point14PassClosureManifest) { model.PointID = "point_15" }, Point14ValEStateBlocked},
		{"missing wave id blocks", func(model *Point14PassClosureManifest) { model.WaveID = "" }, Point14ValEStateBlocked},
		{"wrong closure token blocks", func(model *Point14PassClosureManifest) { model.ClosureToken = "point_14_prepass" }, Point14ValEStateBlocked},
		{"whitespace retagged generated at blocks", func(model *Point14PassClosureManifest) { model.GeneratedAt = " " + model.GeneratedAt + " " }, Point14ValEStateBlocked},
		{"different canonical generated at blocks", func(model *Point14PassClosureManifest) { model.GeneratedAt = "2026-05-07T10:00:00Z" }, Point14ValEStateBlocked},
		{"missing dependency gate result blocks", func(model *Point14PassClosureManifest) { model.DependencyGateResult = "" }, Point14ValEStateBlocked},
		{"missing tests run blocks", func(model *Point14PassClosureManifest) { model.TestsRun = nil }, Point14ValEStateBlocked},
		{"missing negative fixtures run blocks", func(model *Point14PassClosureManifest) { model.NegativeFixturesRun = nil }, Point14ValEStateBlocked},
		{"reviewer result not pass confirmed blocks", func(model *Point14PassClosureManifest) {
			model.ReviewerResult = point12ValEReviewerResultReviewRequired
		}, Point14ValEStateBlocked},
		{"review required dependency returns review required", func(model *Point14PassClosureManifest) {
			model.DependencyGateResult = Point14ValEStateReviewRequired
			model.Point14PassAllowed = false
			model.Point14PassToken = ""
		}, Point14ValEStateReviewRequired},
		{"incomplete dependency returns incomplete", func(model *Point14PassClosureManifest) {
			model.DependencyGateResult = Point14ValEStateIncomplete
			model.Point14PassAllowed = false
			model.Point14PassToken = ""
		}, Point14ValEStateIncomplete},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValEValidPassClosureManifestModel()
			tc.mutate(&model)
			if got := EvaluatePoint14PassClosureManifestState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValEExternalSignalValidationClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValEExternalSignalValidationClosureCheck)
		want   string
	}{
		{"external signal candidate remains bounded active", func(model *Point14ValEExternalSignalValidationClosureCheck) {}, Point14ValEStatePassConfirmed},
		{"whitespace retagged vala current state blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) {
			model.ValACurrentState = " " + model.ValACurrentState + " "
		}, Point14ValEStateBlocked},
		{"tab newline retagged validation result state blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) {
			model.ValidationResultState = "\t" + model.ValidationResultState + "\n"
		}, Point14ValEStateBlocked},
		{"external signal canonical authority blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) { model.CanonicalAuthorityGranted = true }, Point14ValEStateBlocked},
		{"external signal pass blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) { model.PassEmitted = true }, Point14ValEStateBlocked},
		{"source identity authority grant blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) {
			model.SourceIdentityAuthorityGranted = true
		}, Point14ValEStateBlocked},
		{"duplicate active evidence path blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) { model.DuplicateActiveEvidencePath = true }, Point14ValEStateBlocked},
		{"unrelated signal active path blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) { model.UnrelatedSignalActivePath = true }, Point14ValEStateBlocked},
		{"cross tenant signal active path blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) { model.CrossTenantSignalActivePath = true }, Point14ValEStateBlocked},
		{"canonical mutation path blocks", func(model *Point14ValEExternalSignalValidationClosureCheck) { model.CanonicalMutationDetected = true }, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValEExternalSignalValidationClosureCheckModel(point14ValEDependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14ValEExternalSignalValidationClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValEConflictDisputeClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValEConflictDisputeClosureCheck)
		want   string
	}{
		{"bounded unresolved dispute review required prevents pass", func(model *Point14ValEConflictDisputeClosureCheck) { model.UnresolvedDisputeReviewRequired = true }, Point14ValEStateReviewRequired},
		{"whitespace retagged valb current state blocks", func(model *Point14ValEConflictDisputeClosureCheck) {
			model.ValBCurrentState = " " + model.ValBCurrentState + " "
		}, Point14ValEStateBlocked},
		{"tab newline retagged dispute triage result state blocks", func(model *Point14ValEConflictDisputeClosureCheck) {
			model.DisputeTriageResultState = "\t" + model.DisputeTriageResultState + "\n"
		}, Point14ValEStateBlocked},
		{"evidence required prevents pass unless properly closed by governance", func(model *Point14ValEConflictDisputeClosureCheck) { model.EvidenceRequiredUnclosed = true }, Point14ValEStateIncomplete},
		{"crowd public vendor scanner auditor partner agent authority resolution blocks", func(model *Point14ValEConflictDisputeClosureCheck) { model.ExternalAuthorityResolutionDetected = true }, Point14ValEStateBlocked},
		{"dispute auto resolution blocks", func(model *Point14ValEConflictDisputeClosureCheck) { model.DisputeAutoResolved = true }, Point14ValEStateBlocked},
		{"conflict resolving to pass blocks", func(model *Point14ValEConflictDisputeClosureCheck) { model.ConflictResolvedToPass = true }, Point14ValEStateBlocked},
		{"missing governance escalation for unresolved conflict blocks", func(model *Point14ValEConflictDisputeClosureCheck) { model.GovernanceEscalationMissing = true }, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValEConflictDisputeClosureCheckModel(point14ValEDependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14ValEConflictDisputeClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValECorrectionPublicationClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValECorrectionPublicationClosureCheck)
		want   string
	}{
		{"bounded correction notice passes only with governance audit limitations", func(model *Point14ValECorrectionPublicationClosureCheck) {}, Point14ValEStatePassConfirmed},
		{"whitespace retagged valc current state blocks", func(model *Point14ValECorrectionPublicationClosureCheck) {
			model.ValCCurrentState = " " + model.ValCCurrentState + " "
		}, Point14ValEStateBlocked},
		{"tab newline retagged publication approval state blocks", func(model *Point14ValECorrectionPublicationClosureCheck) {
			model.PublicationApprovalState = "\t" + model.PublicationApprovalState + "\n"
		}, Point14ValEStateBlocked},
		{"correction auto published blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.CorrectionAutoPublished = true }, Point14ValEStateBlocked},
		{"revocation auto executed blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.RevocationAutoExecuted = true }, Point14ValEStateBlocked},
		{"supersession silent replacement blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.SupersessionSilentReplacement = true }, Point14ValEStateBlocked},
		{"publication public badge blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.PublicNoticeBecameBadge = true }, Point14ValEStateBlocked},
		{"publication production approved blocks", func(model *Point14ValECorrectionPublicationClosureCheck) {
			model.PublicationApprovalBecameProduction = true
		}, Point14ValEStateBlocked},
		{"publication certified blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.PublicationCertified = true }, Point14ValEStateBlocked},
		{"redaction hiding decisive evidence blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.RedactionHidesDecisiveEvidence = true }, Point14ValEStateBlocked},
		{"limitation omitted blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.LimitationsOmitted = true }, Point14ValEStateBlocked},
		{"publication strengthens claim blocks", func(model *Point14ValECorrectionPublicationClosureCheck) { model.PublicationStrengthensClaims = true }, Point14ValEStateBlocked},
		{"bounded proof missing becomes incomplete", func(model *Point14ValECorrectionPublicationClosureCheck) { model.BoundedCorrectionNoticeProven = false }, Point14ValEStateIncomplete},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValECorrectionPublicationClosureCheckModel(point14ValEDependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14ValECorrectionPublicationClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValETimelineProjectionClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValETimelineProjectionClosureCheck)
		want   string
	}{
		{"read only timeline active", func(model *Point14ValETimelineProjectionClosureCheck) {}, Point14ValEStatePassConfirmed},
		{"whitespace retagged vald current state blocks", func(model *Point14ValETimelineProjectionClosureCheck) {
			model.ValDCurrentState = " " + model.ValDCurrentState + " "
		}, Point14ValEStateBlocked},
		{"tab newline retagged query projection state blocks", func(model *Point14ValETimelineProjectionClosureCheck) {
			model.QueryProjectionState = "\t" + model.QueryProjectionState + "\n"
		}, Point14ValEStateBlocked},
		{"timeline mutation blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.TimelineMutationDetected = true }, Point14ValEStateBlocked},
		{"query mutation blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.QueryMutationDetected = true }, Point14ValEStateBlocked},
		{"timeline dispute resolution blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.TimelineResolvesDisputes = true }, Point14ValEStateBlocked},
		{"timeline correction publication blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.TimelinePublishesCorrections = true }, Point14ValEStateBlocked},
		{"timeline revocation execution blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.TimelineExecutesRevocation = true }, Point14ValEStateBlocked},
		{"query hiding decisive missing evidence blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.QueryHidesDecisiveMissingEvidence = true }, Point14ValEStateBlocked},
		{"timeline authority grant blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.TimelineCreatesAuthority = true }, Point14ValEStateBlocked},
		{"timeline strengthening claim blocks", func(model *Point14ValETimelineProjectionClosureCheck) { model.TimelineStrengthensClaims = true }, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValETimelineProjectionClosureCheckModel(point14ValEDependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14ValETimelineProjectionClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValETenantPrivacyClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValETenantPrivacyClosureCheck)
		want   string
	}{
		{"tenant privacy happy path passes", func(model *Point14ValETenantPrivacyClosureCheck) {}, Point14ValEStatePassConfirmed},
		{"cross tenant timeline query publication blocks clb0", func(model *Point14ValETenantPrivacyClosureCheck) { model.CrossTenantDetected = true }, Point14ValEStateBlocked},
		{"tenant private data exposure blocks clb0", func(model *Point14ValETenantPrivacyClosureCheck) { model.TenantPrivateDataExposed = true }, Point14ValEStateBlocked},
		{"public private classification missing blocks", func(model *Point14ValETenantPrivacyClosureCheck) { model.PublicPrivateClassificationPresent = false }, Point14ValEStateBlocked},
		{"public notice leaking tenant private data blocks", func(model *Point14ValETenantPrivacyClosureCheck) { model.PublicNoticeLeaksTenantPrivateData = true }, Point14ValEStateBlocked},
		{"tenant scope substitution cannot pass final privacy closure", func(model *Point14ValETenantPrivacyClosureCheck) {
			model.TenantScope = "tenant_point14_other"
		}, Point14ValEStateBlocked},
		{"redaction limitation missing blocks", func(model *Point14ValETenantPrivacyClosureCheck) {
			model.RequiredRedactionLimitationRefsPresent = false
		}, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dependency := point14ValEDependencySnapshotModel()
			model := point14ValETenantPrivacyClosureCheckModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValETenantPrivacyClosureCheckState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValETimestampIntegrityClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValETimestampIntegrityClosureCheck)
		want   string
	}{
		{"server utc canonical times pass", func(model *Point14ValETimestampIntegrityClosureCheck) {}, Point14ValEStatePassConfirmed},
		{"approved customer controlled time source passes", func(model *Point14ValETimestampIntegrityClosureCheck) {
			model.EventTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.GeneratedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.ApprovalTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.SourceEventTimeSource = point14Val0TimeSourceApprovedCustomerTime
		}, Point14ValEStatePassConfirmed},
		{"client local time as canonical blocks", func(model *Point14ValETimestampIntegrityClosureCheck) { model.ClientLocalCreatesCanonical = true }, Point14ValEStateBlocked},
		{"source event at authority blocks", func(model *Point14ValETimestampIntegrityClosureCheck) { model.SourceEventCreatesAuthority = true }, Point14ValEStateBlocked},
		{"future dated active event blocks", func(model *Point14ValETimestampIntegrityClosureCheck) { model.FutureDatedActiveEvent = true }, Point14ValEStateBlocked},
		{"backdated correction publication approval blocks", func(model *Point14ValETimestampIntegrityClosureCheck) { model.BackdatedApproval = true }, Point14ValEStateBlocked},
		{"impossible ordering blocks", func(model *Point14ValETimestampIntegrityClosureCheck) { model.ImpossibleOrdering = true }, Point14ValEStateBlocked},
		{"timeline ordering cannot upgrade validity", func(model *Point14ValETimestampIntegrityClosureCheck) { model.TimelineOrderingUpgradesValidity = true }, Point14ValEStateBlocked},
		{"tenant scope mismatch blocks current context substitution", func(model *Point14ValETimestampIntegrityClosureCheck) {
			model.TenantScope = "tenant_scope_point12_beta"
		}, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dependency := point14ValEDependencySnapshotModel()
			model := point14ValETimestampIntegrityClosureCheckModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValETimestampIntegrityClosureCheckState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValEAgentAdvisoryClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValEAgentAdvisoryClosureCheck)
		want   string
	}{
		{"agent advisory input allowed", func(model *Point14ValEAgentAdvisoryClosureCheck) {}, Point14ValEStatePassConfirmed},
		{"agent resolves dispute blocks", func(model *Point14ValEAgentAdvisoryClosureCheck) { model.AgentResolvesDispute = true }, Point14ValEStateBlocked},
		{"agent publishes correction blocks", func(model *Point14ValEAgentAdvisoryClosureCheck) { model.AgentPublishesCorrection = true }, Point14ValEStateBlocked},
		{"agent revokes claim blocks", func(model *Point14ValEAgentAdvisoryClosureCheck) { model.AgentRevokesClaim = true }, Point14ValEStateBlocked},
		{"agent satisfies governance alone blocks", func(model *Point14ValEAgentAdvisoryClosureCheck) { model.AgentSatisfiesGovernanceAlone = true }, Point14ValEStateBlocked},
		{"ai agent authority flags block", func(model *Point14ValEAgentAdvisoryClosureCheck) { model.AgentAuthorityFlags = true }, Point14ValEStateBlocked},
		{"ai agent pass blocks", func(model *Point14ValEAgentAdvisoryClosureCheck) { model.AgentPassAllowed = true }, Point14ValEStateBlocked},
		{"ai agent public badge blocks", func(model *Point14ValEAgentAdvisoryClosureCheck) { model.AgentPublicBadgeAllowed = true }, Point14ValEStateBlocked},
		{"tenant scope mismatch blocks current context substitution", func(model *Point14ValEAgentAdvisoryClosureCheck) {
			model.TenantScope = "tenant_scope_point12_beta"
		}, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dependency := point14ValEDependencySnapshotModel()
			model := point14ValEAgentAdvisoryClosureCheckModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValEAgentAdvisoryClosureCheckState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValEAuthorityBoundaryClosureCheckState(t *testing.T) {
	t.Run("zero-width observed authority mark blocks", func(t *testing.T) {
		model := point14ValEAuthorityBoundaryClosureCheckModel(point14ValEDependencySnapshotModel())
		model.ObservedAuthorityMarks = []string{"external_pass\u200b"}
		if got := EvaluatePoint14ValEAuthorityBoundaryClosureCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("long-s observed authority mark blocks", func(t *testing.T) {
		model := point14ValEAuthorityBoundaryClosureCheckModel(point14ValEDependencySnapshotModel())
		model.ObservedAuthorityMarks = []string{"external_paſs"}
		if got := EvaluatePoint14ValEAuthorityBoundaryClosureCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("split observed authority mark blocks", func(t *testing.T) {
		model := point14ValEAuthorityBoundaryClosureCheckModel(point14ValEDependencySnapshotModel())
		model.ObservedAuthorityMarks = []string{"public", "badge", "authority"}
		if got := EvaluatePoint14ValEAuthorityBoundaryClosureCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
}

func TestPoint14ValENoOverclaimFinalCheckState(t *testing.T) {
	t.Run("forbidden final closure wording blocks", func(t *testing.T) {
		model := point14ValENoOverclaimFinalCheckModel(point14ValEDependencySnapshotModel())
		model.ObservedTexts = []string{"timeline proves truth"}
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("zero-width forbidden final closure wording blocks", func(t *testing.T) {
		model := point14ValENoOverclaimFinalCheckModel(point14ValEDependencySnapshotModel())
		model.ObservedTexts = []string{"timeline proves tr\u200buth"}
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("split forbidden final closure wording blocks across observed corpus", func(t *testing.T) {
		model := point14ValENoOverclaimFinalCheckModel(point14ValEDependencySnapshotModel())
		model.ObservedTexts = []string{"timeline proves", "truth"}
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("long-s final pass wording blocks", func(t *testing.T) {
		model := point14ValENoOverclaimFinalCheckModel(point14ValEDependencySnapshotModel())
		model.ObservedTexts = []string{"external PAſS"}
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("inherited ValA omitted observed category blocks", func(t *testing.T) {
		dependency := point14ValEDependencySnapshotModel()
		dependency.Point14ValD.Dependency.Point14ValA.NoOverclaimValidationWording.ObservedSourceIdentityTexts = []string{"production approved"}
		model := point14ValENoOverclaimFinalCheckModel(dependency)
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("inherited ValB-only forbidden wording blocks", func(t *testing.T) {
		dependency := point14ValEDependencySnapshotModel()
		dependency.Point14ValD.Dependency.Point14ValB.NoOverclaimDisputeWording.ObservedDisputeTexts = []string{"dispute resolved by AI"}
		model := point14ValENoOverclaimFinalCheckModel(dependency)
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("inherited ValC-only forbidden wording blocks", func(t *testing.T) {
		dependency := point14ValEDependencySnapshotModel()
		dependency.Point14ValD.Dependency.Point14ValC.NoOverclaimPublicationWording.ObservedCorrectionTexts = []string{"correction certified"}
		model := point14ValENoOverclaimFinalCheckModel(dependency)
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("inherited ValD omitted observed category blocks", func(t *testing.T) {
		dependency := point14ValEDependencySnapshotModel()
		dependency.Point14ValD.NoOverclaimTimelineWording.ObservedAccessTexts = []string{"public badge"}
		model := point14ValENoOverclaimFinalCheckModel(dependency)
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("nested ValC copy ValB forbidden wording blocks", func(t *testing.T) {
		dependency := point14ValEDependencySnapshotModel()
		dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.NoOverclaimDisputeWording.ObservedDisputeTexts = []string{"dispute resolved by AI"}
		model := point14ValENoOverclaimFinalCheckModel(dependency)
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("nested ValB copy ValA forbidden wording blocks", func(t *testing.T) {
		dependency := point14ValEDependencySnapshotModel()
		dependency.Point14ValD.Dependency.Point14ValB.Dependency.Point14ValA.NoOverclaimValidationWording.ObservedSourceIdentityTexts = []string{"production approved"}
		model := point14ValENoOverclaimFinalCheckModel(dependency)
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("safe bounded wording passes", func(t *testing.T) {
		model := point14ValENoOverclaimFinalCheckModel(point14ValEDependencySnapshotModel())
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStatePassConfirmed {
			t.Fatalf("expected pass_confirmed, got %s", got)
		}
	})

	t.Run("forbidden wording allowed only in internal blocked diagnostic context", func(t *testing.T) {
		model := point14ValENoOverclaimFinalCheckModel(point14ValEDependencySnapshotModel())
		model.InternalDiagnosticTexts = []string{"query approved"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStatePassConfirmed {
			t.Fatalf("expected pass_confirmed, got %s", got)
		}
	})

	t.Run("duplicate safe internal diagnostics do not false positive", func(t *testing.T) {
		dependency := point14ValEDependencySnapshotModel()
		dependency.Point14ValD.NoOverclaimTimelineWording.InternalDiagnosticTexts = []string{"bounded timeline projection"}
		dependency.Point14ValD.Dependency.Point14ValC.NoOverclaimPublicationWording.InternalDiagnosticTexts = []string{"bounded timeline projection"}
		model := point14ValENoOverclaimFinalCheckModel(dependency)
		if len(model.InternalDiagnosticTexts) != 1 {
			t.Fatalf("expected duplicate internal diagnostic to be deduped, got %#v", model.InternalDiagnosticTexts)
		}
		if got := EvaluatePoint14ValENoOverclaimFinalCheckState(model); got != Point14ValEStatePassConfirmed {
			t.Fatalf("expected pass_confirmed, got %s", got)
		}
	})
}

func TestPoint14ValECLBFinalCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValECLBFinalCheck)
		want   string
	}{
		{"clb final happy path passes", func(model *Point14ValECLBFinalCheck) {}, Point14ValEStatePassConfirmed},
		{"clb0 present blocks", func(model *Point14ValECLBFinalCheck) { model.CLB0Present = true }, Point14ValEStateBlocked},
		{"clb1 present blocks", func(model *Point14ValECLBFinalCheck) { model.CLB1Present = true }, Point14ValEStateBlocked},
		{"clb2 present blocks", func(model *Point14ValECLBFinalCheck) { model.CLB2Present = true }, Point14ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValECLBFinalCheckModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValECLBFinalCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValEFoundationHappyPath(t *testing.T) {
	model := ComputePoint14ValEFoundation(Point14ValEFoundationModel())
	if model.CurrentState != Point14ValEStatePassConfirmed {
		t.Fatalf("expected pass_confirmed, got %#v", model)
	}
	if !model.Point14PassAllowed || model.Point14PassToken != point14Val0BlockedPassToken {
		t.Fatalf("expected final point_14_pass happy path, got %#v", model)
	}
	if model.PassClosureManifestState != Point14ValEStatePassConfirmed || model.ClosureEvaluatorState != Point14ValEStatePassConfirmed {
		t.Fatalf("expected manifest and closure evaluator pass_confirmed, got %#v", model)
	}
}

func TestPoint14ValEFoundationStaleSafeNoOverclaimFinalCheckBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.NoOverclaimFinalCheck.ObservedTexts = append(model.NoOverclaimFinalCheck.ObservedTexts, "validated deployment baseline")
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.NoOverclaimFinalCheckState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "no_overclaim") {
		t.Fatalf("expected stale safe no-overclaim final check mismatch to block, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after stale no-overclaim final check mismatch, got %#v", model)
	}
}

func TestPoint14ValEFoundationTimestampTenantMismatchBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.TimestampIntegrityClosureCheck.TenantScope = "tenant_scope_point12_beta"
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.TimestampIntegrityClosureState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "timestamp_integrity") {
		t.Fatalf("expected timestamp tenant mismatch to block final pass, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after timestamp tenant mismatch, got %#v", model)
	}
}

func TestPoint14ValEFoundationAgentTenantMismatchBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.AgentAdvisoryClosureCheck.TenantScope = "tenant_scope_point12_beta"
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.AgentAdvisoryClosureState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "agent_advisory") {
		t.Fatalf("expected agent tenant mismatch to block final pass, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after agent tenant mismatch, got %#v", model)
	}
}

func TestPoint14ValEFoundationNestedInheritedNoOverclaimBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.NoOverclaimDisputeWording.ObservedDisputeTexts = []string{"dispute resolved by AI"}
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.NoOverclaimFinalCheckState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "no_overclaim") {
		t.Fatalf("expected nested inherited no-overclaim block, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after nested inherited overclaim, got %#v", model)
	}
}

func TestPoint14ValEFoundationNestedValBValANoOverclaimBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.Dependency.Point14ValD.Dependency.Point14ValB.Dependency.Point14ValA.NoOverclaimValidationWording.ObservedSourceIdentityTexts = []string{"production approved"}
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.NoOverclaimFinalCheckState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "no_overclaim") {
		t.Fatalf("expected nested ValB->ValA no-overclaim block, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after nested ValA overclaim, got %#v", model)
	}
}

func TestPoint14ValEFoundationNestedVal0NoOverclaimBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.Dependency.Point14ValD.Dependency.Point14ValA.Dependency.Point14Val0.NoOverclaimEcosystemWording.ObservedAgentTexts = []string{"public badge"}
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.DependencyState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "dependency") {
		t.Fatalf("expected nested Val0 no-overclaim dependency block, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after nested Val0 overclaim, got %#v", model)
	}
}

func TestPoint14ValEFoundationNestedVal0AllowedNoOverclaimLedgerBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.Dependency.Point14ValD.Dependency.Point14ValA.Dependency.Point14Val0.NoOverclaimEcosystemWording.AllowedSafeWording = append(
		model.Dependency.Point14ValD.Dependency.Point14ValA.Dependency.Point14Val0.NoOverclaimEcosystemWording.AllowedSafeWording,
		"production approved",
	)
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.DependencyState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "dependency") {
		t.Fatalf("expected nested Val0 allowed no-overclaim ledger mutation to block, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after nested Val0 allowed wording mutation, got %#v", model)
	}
}

func TestPoint14ValEFoundationTenantSubstitutionBlocksFinalPass(t *testing.T) {
	model := Point14ValEFoundationModel()
	model.TenantPrivacyClosureCheck.TenantScope = "tenant_point14_other"
	model = ComputePoint14ValEFoundation(model)
	if model.CurrentState != Point14ValEStateBlocked ||
		model.TenantPrivacyClosureState != Point14ValEStateBlocked ||
		!point12Val0StringSliceContains(model.BlockingReasons, "tenant_privacy") {
		t.Fatalf("expected tenant privacy closure blocked, got %#v", model)
	}
	if model.Point14PassAllowed || model.Point14PassToken != "" {
		t.Fatalf("expected no final point_14_pass after tenant substitution, got %#v", model)
	}
}
