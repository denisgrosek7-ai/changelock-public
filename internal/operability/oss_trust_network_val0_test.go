package operability

import "testing"

func copyOSSTrustNetworkVal0Evidence() []ReferenceArchitectureEvidenceReference {
	evidence := ossTrustNetworkVal0Evidence()
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

func activeOSSTrustNetworkVal0Model() OSSTrustNetworkVal0Foundation {
	model := OSSTrustNetworkVal0FoundationModel()
	return ComputeOSSTrustNetworkVal0Foundation(model)
}

func activeOSSTrustNetworkVal0Limitations() []string {
	return []string{
		"Val 0 defines only the OSS trust discipline foundation and does not implement registry connectors, signing integrations, shared intelligence workflow, dashboards, or later-wave closure.",
		"OSS trust network outputs remain advisory or projection-only and cannot approve deployment, certify packages, create enterprise authority, or mutate canonical evidence.",
	}
}

func TestOSSTrustNetworkVal0HappyPathAndPoint9NotComplete(t *testing.T) {
	model := activeOSSTrustNetworkVal0Model()
	if model.CurrentState != OSSTrustNetworkVal0StateActive {
		t.Fatalf("expected active OSTN Val 0 state, got %#v", model)
	}
	if model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected point 9 to remain not complete, got %#v", model)
	}
	if got := EvaluateOSSTrustNetworkVal0ProofsState(model, activeOSSTrustNetworkVal0Limitations()); got != OSSTrustNetworkVal0StateActive {
		t.Fatalf("expected active proofs state, got %q", got)
	}
	if ossTrustNetworkVal0ContainsForbiddenClaim("not a global truth layer") {
		t.Fatalf("expected exact safe negative wording to remain allowed")
	}
}

func TestOSSTrustNetworkVal0DependencyAndSignalContractBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkVal0Foundation)
		wantState string
	}{
		{
			name: "missing tocka 8 vale dependency blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Dependency.CurrentState = DeveloperEcosystemValEStateBlocked
			},
			wantState: OSSTrustNetworkVal0StatePartial,
		},
		{
			name: "route presence alone cannot satisfy dependency gate",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Dependency = OSSTrustNetworkVal0DependencySnapshot{
					CurrentState:         DeveloperEcosystemValEStatePass,
					Point8State:          DeveloperEcosystemPoint8StateNotComplete,
					Point8PassAllowed:    false,
					Point8PassReason:     DeveloperEcosystemValEPoint8PassReasonBlocked,
					ClosureState:         DeveloperEcosystemValEClosureStateActive,
					NoOverclaimState:     DeveloperEcosystemValENoOverclaimStateActive,
					ProofSurfaceRefs:     append([]string{}, DeveloperEcosystemValEProofSurfaceRefs()...),
					EvidenceRefs:         append([]string{}, DeveloperEcosystemValEProofEvidenceRefs()...),
					ProjectionDisclaimer: developerEcosystemValEProjectionDisclaimer(),
				}
			},
			wantState: OSSTrustNetworkVal0StatePartial,
		},
		{
			name: "point8 pass allowed with blocked reason blocks dependency",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Dependency.Point8PassAllowed = true
				model.Dependency.Point8PassReason = DeveloperEcosystemValEPoint8PassReasonBlocked
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "point8 pass allowed with extended reason blocks dependency",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Dependency.Point8PassAllowed = true
				model.Dependency.Point8PassReason = "point_8_pass allowed certified"
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "point8 pass allowed with empty reason blocks dependency",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Dependency.Point8PassAllowed = true
				model.Dependency.Point8PassReason = " "
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "unknown review state fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.ReviewState = "reviewed-ish"
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "duplicate review state fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.SupportedReviewStates = append(
					removeTrimmedString(model.SignalContract.SupportedReviewStates, OSSTrustNetworkReviewStateRevoked),
					OSSTrustNetworkReviewStateReviewed,
				)
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "duplicate signal evidence ref fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.EvidenceRefs = []string{
					"evidence:ostn-val0-signal-contract-001",
					"evidence:ostn-val0-signal-contract-001",
				}
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "candidate signal cannot be emitted as reviewed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.ReviewState = OSSTrustNetworkReviewStateCandidate
				model.SignalContract.PresentedAsReviewed = true
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "shared vex candidate cannot become reviewed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkReviewStateCandidate
				model.SharedVEXTriage.PresentedAsReviewed = true
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "rejected signal cannot be propagated as usable trust",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkReviewStateRejected
				model.SharedVEXTriage.RejectedPropagatedUsable = true
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "revoked signal fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkReviewStateRevoked
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "superseded signal without replacement fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SharedVEXTriage.ReviewState = OSSTrustNetworkReviewStateSuperseded
				model.SharedVEXTriage.ReplacementRef = ""
			},
			wantState: OSSTrustNetworkVal0StateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkVal0Model()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkVal0Foundation(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkVal0EvidenceFreshnessAndBoundaryBlockers(t *testing.T) {
	testCases := []struct {
		name            string
		mutate          func(*OSSTrustNetworkVal0Foundation)
		expectNotActive bool
		wantCurrent     string
	}{
		{
			name: "duplicate foundation evidence ref fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.EvidenceRefs = append(
					removeTrimmedString(model.EvidenceRefs, "evidence:ostn-val0-point9-governance-001"),
					"evidence:ostn-val0-dependency-001",
				)
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "whitespace foundation evidence ref fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.EvidenceRefs[0] = " "
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "unknown foundation evidence ref fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.EvidenceRefs = append(model.EvidenceRefs, "evidence:ostn-val0-extra-001")
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "unknown proof ref fails closed",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/oss-trust-network/val0/debug")
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "maintainer attestation without linkage blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.MaintainerIdentity.KeyToMaintainerLinkage = false
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "key rotation absence blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.MaintainerIdentity.KeyRotationHandled = false
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "registry freshness unknown blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.RegistryFreshness.FreshnessState = IntelligenceCalibrationFreshnessUnknown
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "signal contract unknown freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.FreshnessState = IntelligenceCalibrationFreshnessUnknown
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "signal contract stale freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.FreshnessState = IntelligenceCalibrationFreshnessStale
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "signal contract expired freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.FreshnessState = IntelligenceCalibrationFreshnessExpired
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "signal contract unsupported freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.FreshnessState = IntelligenceCalibrationFreshnessUnsupported
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "signal contract malformed freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.SignalContract.FreshnessState = "fresh-ish"
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "unsupported registry state is explicit and non-active",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.RegistryFreshness.FreshnessState = IntelligenceCalibrationFreshnessUnsupported
				model.RegistryFreshness.UnsupportedRegistryStateExplicit = true
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StatePartial,
		},
		{
			name: "registry metadata alone cannot create reviewed trust",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.RegistryFreshness.RegistryMetadataAloneReviewedTrust = true
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "propagation unknown freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Propagation.FreshnessState = IntelligenceCalibrationFreshnessUnknown
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "propagation stale freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Propagation.FreshnessState = IntelligenceCalibrationFreshnessStale
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
		{
			name: "propagation malformed freshness blocks",
			mutate: func(model *OSSTrustNetworkVal0Foundation) {
				model.Propagation.FreshnessState = "fresh-ish"
			},
			expectNotActive: true,
			wantCurrent:     OSSTrustNetworkVal0StateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkVal0Model()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkVal0Foundation(model)
		if model.CurrentState != tc.wantCurrent {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantCurrent, model)
		}
		if tc.expectNotActive && model.CurrentState == OSSTrustNetworkVal0StateActive {
			t.Fatalf("%s: expected non-active state, got %#v", tc.name, model)
		}
	}
}

func TestOSSTrustNetworkVal0ProofEvidenceQualityValidation(t *testing.T) {
	testCases := []struct {
		name     string
		evidence func() []ReferenceArchitectureEvidenceReference
		refs     func() []string
		want     bool
	}{
		{
			name: "exact evidence model passes",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return copyOSSTrustNetworkVal0Evidence()
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: true,
		},
		{
			name: "mismatched evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkVal0Evidence()
				evidence[0].EvidenceID, evidence[1].EvidenceID = evidence[1].EvidenceID, evidence[0].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "duplicate evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkVal0Evidence()
				evidence[0].EvidenceID = evidence[1].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unknown evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkVal0Evidence()
				evidence[0].EvidenceID = "evidence:ostn-val0-unknown-001"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "whitespace evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkVal0Evidence()
				evidence[0].EvidenceID = " "
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong scope fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkVal0Evidence()
				evidence[0].Scope = "wrong_scope"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong evidence type fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkVal0Evidence()
				evidence[0].EvidenceType = "wrong_type"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong source fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkVal0Evidence()
				evidence[0].Source = "wrong/source"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "fresh but unrelated evidence payload fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := make([]ReferenceArchitectureEvidenceReference, 0, len(OSSTrustNetworkVal0ProofEvidenceRefs()))
				for idx := range OSSTrustNetworkVal0ProofEvidenceRefs() {
					evidence = append(evidence, ReferenceArchitectureEvidenceReference{
						EvidenceID:     "evidence:unrelated-payload-" + string(rune('a'+idx)),
						EvidenceType:   "unrelated_type",
						Source:         "unrelated/source",
						Timestamp:      "2026-04-29T09:00:00Z",
						FreshnessState: IntelligenceCalibrationFreshnessFresh,
						Scope:          "unrelated_scope",
						Caveats:        []string{"fresh but unrelated"},
					})
				}
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...)
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := OSSTrustNetworkVal0ProofEvidenceQualityValid(tc.evidence(), tc.refs()); got != tc.want {
			t.Fatalf("%s: expected %t, got %t", tc.name, tc.want, got)
		}
	}
}

func TestOSSTrustNetworkVal0TrustPropagationAndOverclaimBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*OSSTrustNetworkVal0Foundation)
	}{
		{name: "generic changelock verified badge blocks", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.TrustMarking.GenericVerifiedBadgeClaim = true
		}},
		{name: "integrity score blocks", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.TrustMarking.IntegrityScoreClaim = true
		}},
		{name: "score greater than ninety blocks", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.TrustMarking.ScoreGreaterThanNinetyClaim = true
		}},
		{name: "universal trust score blocks", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.TrustMarking.UniversalTrustScoreClaim = true
		}},
		{name: "automatic global propagation blocks", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.Propagation.AutomaticGlobalSpread = true
		}},
		{name: "community candidate cannot become canonical truth", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.Propagation.CommunityCandidateCanonicalTruth = true
		}},
		{name: "community report cannot become final truth", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.Propagation.CommunityReportAsFinalTruth = true
		}},
		{name: "local enterprise applicability cannot be silently overridden", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.LocalApplicability.CommunitySignalRewritesEnterprise = true
		}},
		{name: "legal certification claim blocks", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.NoOverclaim.LegalCertification = true
		}},
		{name: "forbidden pass claim blocks", mutate: func(model *OSSTrustNetworkVal0Foundation) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "point_9_pass")
		}},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkVal0Model()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkVal0Foundation(model)
		if model.CurrentState == OSSTrustNetworkVal0StateActive {
			t.Fatalf("%s: expected fail-closed non-active state, got %#v", tc.name, model)
		}
	}
}
