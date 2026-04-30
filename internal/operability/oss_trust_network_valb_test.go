package operability

import "testing"

func copyOSSTrustNetworkValBEvidence() []ReferenceArchitectureEvidenceReference {
	evidence := ossTrustNetworkValBEvidence()
	cloned := make([]ReferenceArchitectureEvidenceReference, 0, len(evidence))
	for _, item := range evidence {
		cloned = append(cloned, ReferenceArchitectureEvidenceReference{
			EvidenceID:     item.EvidenceID,
			EvidenceType:   item.EvidenceType,
			Source:         item.Source,
			Timestamp:      item.Timestamp,
			FreshnessState: item.FreshnessState,
			Scope:          item.Scope,
			Caveats:        append([]string{}, item.Caveats...),
		})
	}
	return cloned
}

func activeOSSTrustNetworkValBModel() OSSTrustNetworkValBCore {
	model := OSSTrustNetworkValBCoreModel()
	return ComputeOSSTrustNetworkValBCore(model)
}

func activeOSSTrustNetworkValBLimitations() []string {
	return []string{
		"Val B defines bounded shared reviewed intelligence only and does not implement dashboards, remediation workflows, final closure, or Točka 10.",
		"Shared reviewed intelligence remains advisory, locally bounded, and cannot become canonical truth, certification, approval, or final Point 9 closure.",
	}
}

func TestOSSTrustNetworkValBHappyPathAndPoint9NotComplete(t *testing.T) {
	model := activeOSSTrustNetworkValBModel()
	if model.CurrentState != OSSTrustNetworkValBStateActive {
		t.Fatalf("expected active OSTN Val B state, got %#v", model)
	}
	if model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected point 9 to remain not complete, got %#v", model)
	}
	if got := EvaluateOSSTrustNetworkValBProofsState(model, activeOSSTrustNetworkValBLimitations()); got != OSSTrustNetworkValBStateActive {
		t.Fatalf("expected active Val B proofs state, got %q", got)
	}
	if ossTrustNetworkValBContainsForbiddenClaim("not canonical truth") {
		t.Fatalf("expected exact safe negative wording to remain allowed")
	}
}

func TestOSSTrustNetworkValBDependencyAndCandidateIntakeBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValBCore)
		wantState string
	}{
		{
			name: "missing vala dependency blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.Dependency.ValACurrentState = OSSTrustNetworkValAStateBlocked
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "route presence alone cannot satisfy vala dependency",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.Dependency = OSSTrustNetworkValBDependencySnapshot{
					ValACurrentState:            OSSTrustNetworkValAStateActive,
					ValAPoint9State:             OSSTrustNetworkPoint9StateNotComplete,
					ValADependencyState:         OSSTrustNetworkValADependencyStateActive,
					ValAReleaseTrustIntakeState: OSSTrustNetworkValAReleaseTrustIntakeStatePartial,
					ValASigningSignalState:      OSSTrustNetworkValASigningSignalStateActive,
					ValAMaintainerState:         OSSTrustNetworkValAMaintainerAttestationStateActive,
					ValAProvenanceState:         OSSTrustNetworkValAProvenanceMaterialStateActive,
					ValARegistryDescriptorState: OSSTrustNetworkValARegistryDescriptorStateActive,
					ValARegistryMetadataState:   OSSTrustNetworkValARegistryMetadataStateActive,
					ValATypoWarningState:        OSSTrustNetworkValATypoSquattingWarningStateActive,
					ValADriftSignalState:        OSSTrustNetworkValADriftSignalStateActive,
					ValANoOverclaimState:        OSSTrustNetworkValANoOverclaimStateActive,
					ValAProofSurfaceRefs:        append([]string{}, OSSTrustNetworkValAProofSurfaceRefs()...),
					ValAEvidenceRefs:            append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...),
					ValAProjectionDisclaimer:    ossTrustNetworkValAProjectionDisclaimer(),
				}
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "vala malformed evidence identity blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.Dependency.ValAEvidenceRefs = append(model.Dependency.ValAEvidenceRefs, "evidence:ostn-vala-extra-001")
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "vala non active subgate blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.Dependency.ValANoOverclaimState = OSSTrustNetworkValANoOverclaimStateBlocked
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name:      "candidate intake normalized with evidence freshness scope caveats passes",
			mutate:    func(model *OSSTrustNetworkValBCore) {},
			wantState: OSSTrustNetworkValBStateActive,
		},
		{
			name: "candidate intake received is partial",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.CandidateSignalIntake.IntakeState = OSSTrustNetworkValBCandidateIntakeStateReceived
			},
			wantState: OSSTrustNetworkValBStatePartial,
		},
		{
			name: "candidate intake unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.CandidateSignalIntake.IntakeState = OSSTrustNetworkValBCandidateIntakeStateUnsupported
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "candidate intake stale blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.CandidateSignalIntake.IntakeState = OSSTrustNetworkValBCandidateIntakeStateStale
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "candidate intake malformed blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.CandidateSignalIntake.IntakeState = OSSTrustNetworkValBCandidateIntakeStateMalformed
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "candidate intake unknown blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.CandidateSignalIntake.IntakeState = OSSTrustNetworkValBCandidateIntakeStateUnknown
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "candidate intake cannot create reviewed trust",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.CandidateSignalIntake.CreatesReviewedTrust = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "candidate intake cannot create global blocklist",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.CandidateSignalIntake.GlobalBlocklistClaim = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValBModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValBCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValBReviewAndSharedVEXBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValBCore)
		wantState string
	}{
		{
			name:      "review workflow reviewed accepted with evidence rationale passes",
			mutate:    func(model *OSSTrustNetworkValBCore) {},
			wantState: OSSTrustNetworkValBStateActive,
		},
		{
			name: "review candidate is partial",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewWorkflow.ReviewState = OSSTrustNetworkValBReviewStateCandidate
				model.ReviewWorkflow.ReviewerDecisionState = OSSTrustNetworkValBReviewerDecisionStateNone
			},
			wantState: OSSTrustNetworkValBStatePartial,
		},
		{
			name: "review in review is partial",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewWorkflow.ReviewState = OSSTrustNetworkValBReviewStateInReview
				model.ReviewWorkflow.ReviewerDecisionState = OSSTrustNetworkValBReviewerDecisionStateNeedsMoreEvidence
			},
			wantState: OSSTrustNetworkValBStatePartial,
		},
		{
			name: "rejected cannot propagate as usable trust",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewWorkflow.ReviewState = OSSTrustNetworkValBReviewStateRejected
				model.ReviewWorkflow.ReviewerDecisionState = OSSTrustNetworkValBReviewerDecisionStateRejected
				model.ReviewWorkflow.ReviewerRationale = "rejected with bounded rationale"
				model.ReviewWorkflow.RejectedPropagatedUsable = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "superseded without replacement blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewWorkflow.ReviewState = OSSTrustNetworkValBReviewStateSuperseded
				model.ReviewWorkflow.ReviewerDecisionState = OSSTrustNetworkValBReviewerDecisionStateSuperseded
				model.ReviewWorkflow.ReviewerRationale = "superseded without replacement"
				model.ReviewWorkflow.ReplacementRef = ""
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "revoked fails closed",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewWorkflow.ReviewState = OSSTrustNetworkValBReviewStateRevoked
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "malformed review state blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewWorkflow.ReviewState = "reviewed-ish"
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "malformed reviewer decision state blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewWorkflow.ReviewerDecisionState = "accepted-ish"
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name:      "shared vex reviewed with evidence local applicability caveats passes",
			mutate:    func(model *OSSTrustNetworkValBCore) {},
			wantState: OSSTrustNetworkValBStateActive,
		},
		{
			name: "shared vex candidate is partial",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkValBSharedVEXStateCandidate
			},
			wantState: OSSTrustNetworkValBStatePartial,
		},
		{
			name: "shared vex rejected blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkValBSharedVEXStateRejected
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "shared vex revoked blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkValBSharedVEXStateRevoked
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "shared vex unknown blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkValBSharedVEXStateUnknown
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "shared vex unsupported blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkValBSharedVEXStateUnsupported
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "shared vex superseded without replacement blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkValBSharedVEXStateSuperseded
				model.SharedVEXTriage.SupersedesRef = ""
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "shared vex cannot override local enterprise applicability",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SharedVEXTriage.OverridesLocalEnterpriseApplicability = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValBModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValBCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValBSourceApplicabilityPropagationAndOverclaimBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValBCore)
		wantState string
	}{
		{
			name:      "source weighting exact class set passes",
			mutate:    func(model *OSSTrustNetworkValBCore) {},
			wantState: OSSTrustNetworkValBStateActive,
		},
		{
			name: "duplicate source class fails closed",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.SupportedSourceClasses = append(
					removeTrimmedString(model.SourceWeighting.SupportedSourceClasses, OSSTrustNetworkValBCandidateSourceClassAutomatedHeuristic),
					OSSTrustNetworkValBCandidateSourceClassCommunity,
				)
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "duplicate source weight class fails closed",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.SupportedSourceWeightClasses = append(
					removeTrimmedString(model.SourceWeighting.SupportedSourceWeightClasses, OSSTrustNetworkValBSourceWeightClassBounded),
					OSSTrustNetworkValBSourceWeightClassMedium,
				)
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "unknown source class fails closed",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.SourceClass = "social"
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "unknown source weight class fails closed",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.SourceWeightClass = "very_high"
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "source weighting malformed review state blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.ReviewState = "reviewed-ish"
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "source weighting malformed projection disclaimer blocks val b active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.ProjectionDisclaimer = "not_canonical_truth but not valb projection"
			},
			wantState: OSSTrustNetworkValBStateUnknown,
		},
		{
			name: "automated heuristic alone cannot become high confidence reviewed signal",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.SourceClass = OSSTrustNetworkValBCandidateSourceClassAutomatedHeuristic
				model.SourceWeighting.SourceWeightClass = OSSTrustNetworkValBSourceWeightClassHigh
				model.SourceWeighting.AutomatedHeuristicStandalone = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "community input alone cannot become reviewed without review workflow",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SourceWeighting.SourceClass = OSSTrustNetworkValBCandidateSourceClassCommunity
				model.SourceWeighting.CommunityInputWithoutReview = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "local applicability applicable requires local evidence and scope",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.LocalApplicability.LocalEvidenceLinked = false
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "local applicability malformed projection disclaimer blocks val b active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.LocalApplicability.ProjectionDisclaimer = "projection_only but not valb"
			},
			wantState: OSSTrustNetworkValBStateUnknown,
		},
		{
			name: "local applicability not applicable requires rationale evidence",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.LocalApplicability.ApplicabilityState = OSSTrustNetworkValBLocalApplicabilityStatusNotApplicable
				model.LocalApplicability.Rationale = ""
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "local applicability needs local review is partial",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.LocalApplicability.ApplicabilityState = OSSTrustNetworkValBLocalApplicabilityStatusNeedsLocalReview
			},
			wantState: OSSTrustNetworkValBStatePartial,
		},
		{
			name: "local applicability unknown blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.LocalApplicability.ApplicabilityState = OSSTrustNetworkValBLocalApplicabilityStatusUnknown
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "local applicability unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.LocalApplicability.ApplicabilityState = OSSTrustNetworkValBLocalApplicabilityStatusUnsupported
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name:      "reviewed exchange with reviewed state evidence source weighting local applicability caveats freshness passes",
			mutate:    func(model *OSSTrustNetworkValBCore) {},
			wantState: OSSTrustNetworkValBStateActive,
		},
		{
			name: "candidate exchange is partial",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateCandidateExchange
				model.PropagationExchange.ReviewState = OSSTrustNetworkValBReviewStateCandidate
				model.PropagationExchange.PresentedAsReviewed = false
			},
			wantState: OSSTrustNetworkValBStatePartial,
		},
		{
			name: "propagation candidate exchange with malformed review state blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateCandidateExchange
				model.PropagationExchange.ReviewState = "candidate-ish"
				model.PropagationExchange.PresentedAsReviewed = false
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "propagation reviewed exchange with malformed local applicability state blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateReviewedExchange
				model.PropagationExchange.ReviewState = OSSTrustNetworkValBReviewStateReviewed
				model.PropagationExchange.LocalApplicabilityState = "applicable-ish"
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "candidate exchange cannot be displayed as reviewed",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateCandidateExchange
				model.PropagationExchange.ReviewState = OSSTrustNetworkValBReviewStateCandidate
				model.PropagationExchange.PresentedAsReviewed = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "rejected propagation blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateRejected
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "revoked propagation blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateRevoked
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "unsupported propagation blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateUnsupported
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "unknown propagation blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateUnknown
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "superseded propagation without replacement blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.PropagationState = OSSTrustNetworkValBPropagationStateSuperseded
				model.PropagationExchange.ReplacementRef = ""
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "automatic global spread blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.AutomaticGlobalSpread = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "propagation malformed projection disclaimer blocks val b active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.ProjectionDisclaimer = "projection_only not_canonical_truth but not valb"
			},
			wantState: OSSTrustNetworkValBStateUnknown,
		},
		{
			name: "network signal overriding local enterprise policy blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.PropagationExchange.EnterpriseOverride = true
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "revocation without reason evidence blocks",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SupersessionRevocation.LifecycleState = OSSTrustNetworkValBLifecycleStateRevoked
				model.SupersessionRevocation.RevocationReason = ""
				model.SupersessionRevocation.RevocationTimestamp = ""
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "reviewer decision missing rationale blocks reviewed state",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewerAuditability.Rationale = ""
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
		{
			name: "supersession revocation malformed projection disclaimer blocks val b active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.SupersessionRevocation.ProjectionDisclaimer = "projection_only not_canonical_truth but not valb"
			},
			wantState: OSSTrustNetworkValBStateUnknown,
		},
		{
			name: "reviewer auditability malformed projection disclaimer blocks val b active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.ReviewerAuditability.ProjectionDisclaimer = "projection_only not_canonical_truth but not valb"
			},
			wantState: OSSTrustNetworkValBStateUnknown,
		},
		{
			name: "no overclaim denylist blocks active",
			mutate: func(model *OSSTrustNetworkValBCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "point_9_pass")
			},
			wantState: OSSTrustNetworkValBStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValBModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValBCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValBProofEvidenceQualityValidation(t *testing.T) {
	testCases := []struct {
		name     string
		evidence func() []ReferenceArchitectureEvidenceReference
		refs     func() []string
		want     bool
	}{
		{
			name: "exact evidence model passes",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return copyOSSTrustNetworkValBEvidence()
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: true,
		},
		{
			name: "mismatched evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValBEvidence()
				evidence[0].EvidenceID, evidence[1].EvidenceID = evidence[1].EvidenceID, evidence[0].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "duplicate evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValBEvidence()
				evidence[0].EvidenceID = evidence[1].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unknown evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValBEvidence()
				evidence[0].EvidenceID = "evidence:ostn-valb-unknown-001"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "whitespace evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValBEvidence()
				evidence[0].EvidenceID = " "
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong scope fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValBEvidence()
				evidence[0].Scope = "wrong_scope"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong evidence type fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValBEvidence()
				evidence[0].EvidenceType = "wrong_type"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong source fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValBEvidence()
				evidence[0].Source = "wrong/source"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "fresh but unrelated evidence payload fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := make([]ReferenceArchitectureEvidenceReference, 0, len(OSSTrustNetworkValBProofEvidenceRefs()))
				for idx := range OSSTrustNetworkValBProofEvidenceRefs() {
					evidence = append(evidence, ReferenceArchitectureEvidenceReference{
						EvidenceID:     "evidence:unrelated-payload-" + string(rune('a'+idx)),
						EvidenceType:   "unrelated_type",
						Source:         "unrelated/source",
						Timestamp:      "2026-04-29T11:30:00Z",
						FreshnessState: IntelligenceCalibrationFreshnessFresh,
						Scope:          "unrelated_scope",
						Caveats:        []string{"fresh but unrelated"},
					})
				}
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...)
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := OSSTrustNetworkValBProofEvidenceQualityValid(tc.evidence(), tc.refs()); got != tc.want {
			t.Fatalf("%s: expected %t, got %t", tc.name, tc.want, got)
		}
	}
}
