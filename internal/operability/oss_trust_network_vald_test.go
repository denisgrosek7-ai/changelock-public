package operability

import "testing"

func activeOSSTrustNetworkValDModel() OSSTrustNetworkValDCore {
	model := OSSTrustNetworkValDCoreModel()
	return ComputeOSSTrustNetworkValDCore(model)
}

func activeOSSTrustNetworkValDLimitations() []string {
	return []string{
		"Val D defines the final OSTN readiness gate only and does not implement integrated closure, Val E, or Točka 10.",
		"Readiness, visibility, remediation, and propagation outputs remain bounded advisory projections and do not become canonical truth, approval authority, or hidden mutation paths.",
	}
}

func TestOSSTrustNetworkValDHappyPathAndPoint9NotComplete(t *testing.T) {
	model := activeOSSTrustNetworkValDModel()
	if model.CurrentState != OSSTrustNetworkValDStateActive {
		t.Fatalf("expected active OSTN Val D state, got %#v", model)
	}
	if model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected point 9 to remain not complete, got %#v", model)
	}
	if got := EvaluateOSSTrustNetworkValDProofsState(model, activeOSSTrustNetworkValDLimitations()); got != OSSTrustNetworkValDStateActive {
		t.Fatalf("expected active Val D proofs state, got %q", got)
	}
}

func TestOSSTrustNetworkValDDependencySignalAndReleaseBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValDCore)
		wantState string
	}{
		{
			name: "missing valc dependency blocks active",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.Dependency.ValCCurrentState = OSSTrustNetworkValCStateBlocked
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "route presence alone cannot satisfy valc dependency",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.Dependency.ValCTrustVisibilityState = OSSTrustNetworkValCTrustVisibilityStatePartial
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "valc malformed evidence identity blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.Dependency.ValCEvidenceRefs = append(model.Dependency.ValCEvidenceRefs, "evidence:ostn-valc-extra-001")
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "valc non active subgate blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.Dependency.ValCNoOverclaimState = OSSTrustNetworkValCNoOverclaimStateBlocked
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "non active valc current state blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.Dependency.ValCCurrentState = OSSTrustNetworkValCStatePartial
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "candidate signal displayed as reviewed blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.SignalLifecycleState = OSSTrustNetworkValDSignalLifecycleCandidate
				model.SignalCorrectness.CandidateDisplayedAsReviewed = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "reviewed signal without accepted decision blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.ReviewerDecisionState = OSSTrustNetworkValBReviewerDecisionStateRejected
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "rejected signal as usable trust blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.SignalLifecycleState = OSSTrustNetworkValDSignalLifecycleRejected
				model.SignalCorrectness.RejectedUsableTrust = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "revoked signal as usable trust blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.SignalLifecycleState = OSSTrustNetworkValDSignalLifecycleRevoked
				model.SignalCorrectness.RevokedUsableTrust = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "superseded signal without replacement blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.SignalLifecycleState = OSSTrustNetworkValDSignalLifecycleSuperseded
				model.SignalCorrectness.ReplacementRef = ""
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "unknown signal state blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.SignalLifecycleState = OSSTrustNetworkValDSignalLifecycleUnknown
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "unsupported signal state blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.SignalLifecycleState = OSSTrustNetworkValDSignalLifecycleUnsupported
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "malformed signal state blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.SignalCorrectness.SignalLifecycleState = "reviewed-ish"
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "signing non verified blocks release foundation gate",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.SigningVerificationState = "present"
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "signing unscoped blocks release foundation gate",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.SigningScoped = false
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "maintainer attestation without key linkage blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.MaintainerKeyLinked = false
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "maintainer attestation without rotation handling blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.MaintainerRotationHandled = false
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "provenance unverified blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.ProvenanceVerificationState = "present_unverified"
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "provenance stale scope blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.ProvenanceReleaseArtifactScoped = false
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "registry metadata alone creating reviewed trust blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.RegistryMetadataCreatesReviewedTrust = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "registry descriptor live fetch behavior blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.LiveRegistryFetch = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "typo warning auto global block blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.TypoWarningAutoGlobalBlock = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "typo warning canonical truth claim blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.TypoWarningCanonicalTruthClaim = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "drift signal overriding local enterprise blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReleaseFoundation.DriftOverridesLocalEnterprise = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValDModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValDCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValDReviewedPropagationAndRemediationBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValDCore)
		wantState string
	}{
		{
			name: "review workflow not reviewed accepted blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.ReviewState = OSSTrustNetworkValBReviewStateCandidate
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "shared vex candidate displayed as reviewed blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.SharedVEXState = OSSTrustNetworkValBSharedVEXStateCandidate
				model.ReviewedIntelligence.SharedVEXDisplayedAsReviewed = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "shared vex rejected blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.SharedVEXState = OSSTrustNetworkValBSharedVEXStateRejected
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "shared vex revoked blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.SharedVEXState = OSSTrustNetworkValBSharedVEXStateRevoked
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "shared vex unknown blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.SharedVEXState = OSSTrustNetworkValBSharedVEXStateUnknown
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "source weighting score claim blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.UniversalTrustScoreClaim = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "local applicability unknown displayed as applicable blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.LocalApplicabilityStatus = OSSTrustNetworkValBLocalApplicabilityStatusUnknown
				model.ReviewedIntelligence.DisplayedAsApplicable = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "reviewer auditability missing rationale blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.ReviewerRationale = ""
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "reviewer auditability missing evidence blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.ReviewedIntelligence.ReviewerEvidenceLinked = false
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "reviewed exchange without similarity context gating blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.PropagationSafety.SimilarityContextGating = false
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "candidate exchange displayed as reviewed blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.PropagationSafety.PropagationState = OSSTrustNetworkValBPropagationStateCandidateExchange
				model.PropagationSafety.CandidateDisplayedAsReviewed = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "automatic global spread blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.PropagationSafety.AutomaticGlobalSpread = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "global blocklist blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.PropagationSafety.GlobalBlocklistClaim = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "enterprise override blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.PropagationSafety.EnterpriseOverride = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "remediation suggestion missing compatibility note blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.CompatibilityNote = ""
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "remediation suggestion missing risk note blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.RiskNote = ""
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "remediation suggestion missing rollback note blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.RollbackNote = ""
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "remediation suggestion missing test validation note blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.TestValidationNote = ""
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "high risk suggestion without reviewer requirement blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.RiskClass = OSSTrustNetworkValCRiskClassHigh
				model.RemediationPRSafety.ReviewerRequired = false
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "no action suggestion hiding risk blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.SuggestionClass = OSSTrustNetworkValCSuggestionClassNoAction
				model.RemediationPRSafety.NoActionHidesRisk = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "pr proposal with auto merge blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.ProposalAutoMerge = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "pr proposal with branch write blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.ProposalBranchWrite = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "pr proposal with network action blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.ProposalNetworkAction = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "pr proposal with dependency mutation blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.ProposalDependencyMutation = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "pr proposal with pr creation blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.ProposalPRCreation = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "hidden mutation path blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.RemediationPRSafety.HiddenMutationPath = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValDModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValDCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValDEcosystemEvidenceAndOverclaimBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValDCore)
		wantState string
	}{
		{
			name: "public summary exposing canonical internals blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EcosystemVisibilityConsistency.ExportClass = OSSTrustNetworkValCExportClassPublicSummaryView
				model.EcosystemVisibilityConsistency.CanonicalInternalExposure = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "redaction converting candidate to reviewed blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EcosystemVisibilityConsistency.CandidatePromotedToReviewed = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "local override rewriting canonical evidence blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EcosystemVisibilityConsistency.RewriteCanonicalEvidence = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "shared signal silently overriding local enterprise decision blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EcosystemVisibilityConsistency.SharedSignalOverridesLocalDecision = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "vald evidence ref duplicate blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceRefs[0] = model.EvidenceRefs[1]
				model.EvidenceQuality.EvidenceRefs[0] = model.EvidenceQuality.EvidenceRefs[1]
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "vald evidence ref unknown blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceQuality.EvidenceRefs = append(model.EvidenceQuality.EvidenceRefs, "evidence:ostn-vald-extra-001")
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "vald evidence ref whitespace blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceQuality.EvidenceRefs[0] = " "
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "vald evidence object wrong source blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceQuality.Evidence[0].Source = "oss-trust-network/vald/other"
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "fresh but unrelated evidence payload blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceQuality.Evidence[0].Scope = "unrelated_scope"
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "stale evidence blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceQuality.Evidence[0].FreshnessState = IntelligenceCalibrationFreshnessStale
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "unknown evidence blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceQuality.Evidence[0].FreshnessState = IntelligenceCalibrationFreshnessUnknown
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "unsupported evidence blocks",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.EvidenceQuality.Evidence[0].FreshnessState = IntelligenceCalibrationFreshnessUnsupported
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "no overclaim denylist blocks active",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "point_9_pass")
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "public badge claim blocks active",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.NoOverclaim.PublicBadgeClaim = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
		{
			name: "official oss authority claim blocks active",
			mutate: func(model *OSSTrustNetworkValDCore) {
				model.NoOverclaim.OfficialOSSAuthorityClaim = true
			},
			wantState: OSSTrustNetworkValDStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValDModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValDCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValDProofEvidenceQualityValidation(t *testing.T) {
	testCases := []struct {
		name     string
		evidence func() []ReferenceArchitectureEvidenceReference
		refs     func() []string
		want     bool
	}{
		{
			name: "exact evidence model passes",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return ossTrustNetworkValDCopyEvidence(ossTrustNetworkValDEvidence())
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValDProofEvidenceRefs()...)
			},
			want: true,
		},
		{
			name: "duplicate evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValDCopyEvidence(ossTrustNetworkValDEvidence())
				evidence[0].EvidenceID = evidence[1].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValDProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unexpected evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValDCopyEvidence(ossTrustNetworkValDEvidence())
				evidence[0].EvidenceID = "evidence:ostn-vald-unknown-001"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValDProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong scope fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValDCopyEvidence(ossTrustNetworkValDEvidence())
				evidence[0].Scope = "other_scope"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValDProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "stale freshness fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValDCopyEvidence(ossTrustNetworkValDEvidence())
				evidence[0].FreshnessState = IntelligenceCalibrationFreshnessStale
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValDProofEvidenceRefs()...)
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := OSSTrustNetworkValDProofEvidenceQualityValid(tc.evidence(), tc.refs()); got != tc.want {
			t.Fatalf("%s: expected %t, got %t", tc.name, tc.want, got)
		}
	}
}
