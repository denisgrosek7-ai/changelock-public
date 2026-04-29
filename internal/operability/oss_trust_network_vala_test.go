package operability

import "testing"

func copyOSSTrustNetworkValAEvidence() []ReferenceArchitectureEvidenceReference {
	evidence := ossTrustNetworkValAEvidence()
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

func activeOSSTrustNetworkValAModel() OSSTrustNetworkValACore {
	model := OSSTrustNetworkValACoreModel()
	return ComputeOSSTrustNetworkValACore(model)
}

func activeOSSTrustNetworkValALimitations() []string {
	return []string{
		"Val A defines bounded release trust and registry core only and does not implement shared reviewed intelligence workflows, dashboards, final closure, or Točka 10.",
		"Registry descriptors remain descriptor-only in Val A and do not perform live network trust fetches or create canonical truth.",
	}
}

func TestOSSTrustNetworkValAHappyPathAndPoint9NotComplete(t *testing.T) {
	model := activeOSSTrustNetworkValAModel()
	if model.CurrentState != OSSTrustNetworkValAStateActive {
		t.Fatalf("expected active OSTN Val A state, got %#v", model)
	}
	if model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected point 9 to remain not complete, got %#v", model)
	}
	if got := EvaluateOSSTrustNetworkValAProofsState(model, activeOSSTrustNetworkValALimitations()); got != OSSTrustNetworkValAStateActive {
		t.Fatalf("expected active Val A proofs state, got %q", got)
	}
	if model.SigningSignalState != OSSTrustNetworkValASigningSignalStateActive ||
		model.MaintainerAttestationState != OSSTrustNetworkValAMaintainerAttestationStateActive ||
		model.ProvenanceMaterialState != OSSTrustNetworkValAProvenanceMaterialStateActive ||
		model.RegistryDescriptorState != OSSTrustNetworkValARegistryDescriptorStateActive ||
		model.RegistryMetadataState != OSSTrustNetworkValARegistryMetadataStateActive {
		t.Fatalf("expected active release trust component states, got %#v", model)
	}
}

func TestOSSTrustNetworkValADependencyAndReleaseTrustBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValACore)
		wantState string
	}{
		{
			name: "missing val0 dependency blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0CurrentState = OSSTrustNetworkVal0StateBlocked
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 current state partial blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0CurrentState = OSSTrustNetworkVal0StatePartial
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 current state unknown blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0CurrentState = OSSTrustNetworkVal0StateUnknown
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 current state incomplete blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0CurrentState = OSSTrustNetworkVal0StateIncomplete
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "route presence alone cannot satisfy dependency",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency = OSSTrustNetworkValADependencySnapshot{
					Val0CurrentState:         OSSTrustNetworkVal0StateActive,
					Val0Point9State:          OSSTrustNetworkPoint9StateNotComplete,
					Val0DependencyState:      OSSTrustNetworkVal0DependencyStateActive,
					Val0NoOverclaimState:     OSSTrustNetworkVal0NoOverclaimStateActive,
					Val0Point8State:          DeveloperEcosystemPoint8StatePass,
					Val0Point8PassAllowed:    false,
					Val0Point8PassReason:     DeveloperEcosystemValEPoint8PassReasonBlocked,
					Val0Point8ClosureState:   DeveloperEcosystemValEClosureStateActive,
					Val0ProofSurfaceRefs:     append([]string{}, OSSTrustNetworkVal0ProofSurfaceRefs()...),
					Val0EvidenceRefs:         append([]string{}, OSSTrustNetworkVal0ProofEvidenceRefs()...),
					Val0ProjectionDisclaimer: ossTrustNetworkVal0ProjectionDisclaimer(),
				}
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 dependency state partial blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0DependencyState = OSSTrustNetworkVal0DependencyStatePartial
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 no overclaim blocked blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0NoOverclaimState = OSSTrustNetworkVal0NoOverclaimStateBlocked
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 malformed evidence identity blocks",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0EvidenceRefs = append(model.Dependency.Val0EvidenceRefs, "evidence:ostn-val0-extra-001")
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 non canonical dependency pass reason blocks",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0Point8PassReason = "point_8_pass allowed certified"
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 blocked dependency pass reason blocks",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0Point8PassReason = DeveloperEcosystemValEPoint8PassReasonBlocked
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 empty dependency pass reason blocks",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0Point8PassReason = " "
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "val0 point8 pass allowed false blocks",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.Dependency.Val0Point8PassAllowed = false
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "missing release identity blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ReleaseTrustIntake.ReleaseRef = " "
			},
			wantState: OSSTrustNetworkValAStateIncomplete,
		},
		{
			name: "missing artifact identity blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ReleaseTrustIntake.ArtifactIdentity = ""
			},
			wantState: OSSTrustNetworkValAStateIncomplete,
		},
		{
			name: "release trust intake unknown freshness blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ReleaseTrustIntake.FreshnessState = IntelligenceCalibrationFreshnessUnknown
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "release trust intake stale freshness blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ReleaseTrustIntake.FreshnessState = IntelligenceCalibrationFreshnessStale
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "release trust intake expired freshness blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ReleaseTrustIntake.FreshnessState = IntelligenceCalibrationFreshnessExpired
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValAModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValACore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValASigningMaintainerAndProvenanceBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValACore)
		wantState string
	}{
		{
			name:      "signing verified with exact evidence passes",
			mutate:    func(model *OSSTrustNetworkValACore) {},
			wantState: OSSTrustNetworkValAStateActive,
		},
		{
			name: "signing present without verification is non active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.SigningSignal.SigningState = OSSTrustNetworkValASigningStatePresent
				model.SigningSignal.VerifiedEvidenceLinked = false
			},
			wantState: OSSTrustNetworkValAStatePartial,
		},
		{
			name: "signing missing blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.SigningSignal.SigningState = OSSTrustNetworkValASigningStateMissing
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "signing mismatch blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.SigningSignal.SigningState = OSSTrustNetworkValASigningStateMismatch
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "signing revoked blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.SigningSignal.SigningState = OSSTrustNetworkValASigningStateRevoked
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "signing unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.SigningSignal.SigningState = OSSTrustNetworkValASigningStateUnsupported
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "signing unknown blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.SigningSignal.SigningState = OSSTrustNetworkValASigningStateUnknown
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "signing malformed blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.SigningSignal.SigningState = "verified-ish"
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name:      "maintainer attested with lifecycle discipline passes",
			mutate:    func(model *OSSTrustNetworkValACore) {},
			wantState: OSSTrustNetworkValAStateActive,
		},
		{
			name: "maintainer attestation without key linkage blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.MaintainerAttestation.KeyLinkageVisible = false
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "maintainer attestation revoked blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.MaintainerAttestation.AttestationState = OSSTrustNetworkValAAttestationStateRevoked
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "maintainer attestation stale blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.MaintainerAttestation.AttestationState = OSSTrustNetworkValAAttestationStateStale
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "maintainer attestation unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.MaintainerAttestation.AttestationState = OSSTrustNetworkValAAttestationStateUnsupported
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "maintainer attestation unknown blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.MaintainerAttestation.AttestationState = OSSTrustNetworkValAAttestationStateUnknown
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "delegated maintainer attestation without reviewed delegation blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.MaintainerAttestation.AttestationState = OSSTrustNetworkValAAttestationStateDelegated
				model.MaintainerAttestation.DelegatedSigningReviewed = false
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name:      "provenance verified with exact evidence passes",
			mutate:    func(model *OSSTrustNetworkValACore) {},
			wantState: OSSTrustNetworkValAStateActive,
		},
		{
			name: "provenance present unverified is non active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ProvenanceMaterial.ProvenanceState = OSSTrustNetworkValAProvenanceStatePresentUnverified
			},
			wantState: OSSTrustNetworkValAStatePartial,
		},
		{
			name: "provenance missing blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ProvenanceMaterial.ProvenanceState = OSSTrustNetworkValAProvenanceStateMissing
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "provenance mismatch blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ProvenanceMaterial.ProvenanceState = OSSTrustNetworkValAProvenanceStateMismatch
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "provenance stale blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ProvenanceMaterial.ProvenanceState = OSSTrustNetworkValAProvenanceStateStale
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "provenance unknown blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ProvenanceMaterial.ProvenanceState = OSSTrustNetworkValAProvenanceStateUnknown
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "provenance unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ProvenanceMaterial.ProvenanceState = OSSTrustNetworkValAProvenanceStateUnsupported
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "provenance malformed blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.ProvenanceMaterial.ProvenanceState = "verified-ish"
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValAModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValACore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValARegistryTypoDriftAndOverclaimBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValACore)
		wantState string
	}{
		{
			name:      "registry descriptor exact supported set passes",
			mutate:    func(model *OSSTrustNetworkValACore) {},
			wantState: OSSTrustNetworkValAStateActive,
		},
		{
			name: "duplicate registry descriptor fails closed",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.RegistryDescriptor.SupportedRegistryDescriptors = append(
					removeTrimmedString(model.RegistryDescriptor.SupportedRegistryDescriptors, OSSTrustNetworkValARegistryDescriptorGenericOSSRegistry),
					OSSTrustNetworkValARegistryDescriptorNPM,
				)
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "unknown registry descriptor fails closed",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.RegistryDescriptor.RequestedRegistryDescriptor = "crates_io"
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "unsupported registry descriptor is explicit and non active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.RegistryDescriptor.RequestedRegistryDescriptor = OSSTrustNetworkValARegistryDescriptorUnsupported
				model.RegistryDescriptor.UnsupportedRegistryExplicit = true
			},
			wantState: OSSTrustNetworkValAStatePartial,
		},
		{
			name: "registry metadata alone cannot create reviewed trust",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.RegistryMetadata.RegistryMetadataCreatesReviewedTrust = true
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "missing normalized registry metadata blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.RegistryMetadata.PackageName = ""
			},
			wantState: OSSTrustNetworkValAStateIncomplete,
		},
		{
			name: "stale normalized metadata freshness blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.RegistryMetadata.MetadataFreshness = IntelligenceCalibrationFreshnessStale
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "unknown normalized metadata freshness blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.RegistryMetadata.MetadataFreshness = IntelligenceCalibrationFreshnessUnknown
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name:      "typo squatting early warning remains candidate and bounded",
			mutate:    func(model *OSSTrustNetworkValACore) {},
			wantState: OSSTrustNetworkValAStateActive,
		},
		{
			name: "typo squatting warning cannot automatically block globally",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.TypoSquattingWarning.AutomaticGlobalBlock = true
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "typo squatting warning cannot become canonical truth",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.TypoSquattingWarning.CanonicalTruthClaim = true
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "drift signal must be evidence linked source weighted scoped and caveated",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.DriftSignal.EvidenceRefs = nil
			},
			wantState: OSSTrustNetworkValAStateIncomplete,
		},
		{
			name: "drift signal cannot override local enterprise applicability",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.DriftSignal.OverridesLocalEnterprise = true
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
		{
			name: "no overclaim denylist blocks active",
			mutate: func(model *OSSTrustNetworkValACore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "point_9_pass")
			},
			wantState: OSSTrustNetworkValAStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValAModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValACore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValAProofEvidenceQualityValidation(t *testing.T) {
	testCases := []struct {
		name     string
		evidence func() []ReferenceArchitectureEvidenceReference
		refs     func() []string
		want     bool
	}{
		{
			name: "exact evidence model passes",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return copyOSSTrustNetworkValAEvidence()
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: true,
		},
		{
			name: "mismatched evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValAEvidence()
				evidence[0].EvidenceID, evidence[1].EvidenceID = evidence[1].EvidenceID, evidence[0].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "duplicate evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValAEvidence()
				evidence[0].EvidenceID = evidence[1].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unknown evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValAEvidence()
				evidence[0].EvidenceID = "evidence:ostn-vala-unknown-001"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "whitespace evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValAEvidence()
				evidence[0].EvidenceID = " "
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong scope fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValAEvidence()
				evidence[0].Scope = "wrong_scope"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong evidence type fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValAEvidence()
				evidence[0].EvidenceType = "wrong_type"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong source fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValAEvidence()
				evidence[0].Source = "wrong/source"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "fresh but unrelated evidence payload fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := make([]ReferenceArchitectureEvidenceReference, 0, len(OSSTrustNetworkValAProofEvidenceRefs()))
				for idx := range OSSTrustNetworkValAProofEvidenceRefs() {
					evidence = append(evidence, ReferenceArchitectureEvidenceReference{
						EvidenceID:     "evidence:unrelated-vala-" + string(rune('a'+idx)),
						EvidenceType:   "unrelated_type",
						Source:         "unrelated/source",
						Timestamp:      "2026-04-29T11:00:00Z",
						FreshnessState: IntelligenceCalibrationFreshnessFresh,
						Scope:          "unrelated_scope",
						Caveats:        []string{"fresh but unrelated"},
					})
				}
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValAProofEvidenceRefs()...)
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := OSSTrustNetworkValAProofEvidenceQualityValid(tc.evidence(), tc.refs()); got != tc.want {
			t.Fatalf("%s: expected %t, got %t", tc.name, tc.want, got)
		}
	}
}
