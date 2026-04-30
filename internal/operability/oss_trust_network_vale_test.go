package operability

import "testing"

func activeOSSTrustNetworkValEModel() OSSTrustNetworkValEIntegratedClosure {
	return ComputeOSSTrustNetworkValEClosure(OSSTrustNetworkValEIntegratedClosureModel())
}

func TestOSSTrustNetworkValEHappyPathPassAndPoint9PassAllowed(t *testing.T) {
	model := activeOSSTrustNetworkValEModel()
	if model.CurrentState != OSSTrustNetworkValEStatePass {
		t.Fatalf("expected Val E pass state, got %#v", model)
	}
	if model.Point9State != OSSTrustNetworkPoint9StatePass || !model.Point9PassAllowed {
		t.Fatalf("expected point_9_pass only in Val E, got %#v", model)
	}
	if model.Point9PassReason != OSSTrustNetworkValEPoint9PassReasonAllowed {
		t.Fatalf("expected canonical allowed pass reason, got %#v", model)
	}
	if model.DependencyState != OSSTrustNetworkValEDependencyStateActive ||
		model.IntegratedClosureState != OSSTrustNetworkValEIntegratedClosureStateActive ||
		model.CanonicalBoundaryState != OSSTrustNetworkValECanonicalBoundaryStateActive ||
		model.EvidenceQualityState != OSSTrustNetworkValEEvidenceQualityStateActive ||
		model.NoOverclaimState != OSSTrustNetworkValENoOverclaimStateActive ||
		model.FinalPassRuleState != OSSTrustNetworkValEFinalPassRuleStateActive ||
		model.ClosureState != OSSTrustNetworkValEClosureStateActive {
		t.Fatalf("expected active Val E closure gates, got %#v", model)
	}
	if model.Val0Source.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValASource.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValBSource.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValCSource.Point9State != OSSTrustNetworkPoint9StateNotComplete ||
		model.ValDSource.Point9State != OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected Val 0-D to remain prerequisites, got %#v", model)
	}
}

func TestOSSTrustNetworkValEDependencyAndSourceBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*OSSTrustNetworkValEIntegratedClosure)
	}{
		{name: "missing vald dependency blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.DependencyState = OSSTrustNetworkValDDependencyStatePartial
		}},
		{name: "route presence alone cannot satisfy vald dependency", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.CurrentState = ""
		}},
		{name: "vald malformed evidence identity blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.EvidenceRefs[0] = "evidence:ostn-vald-unknown-001"
		}},
		{name: "vald non active subgate blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.NoOverclaimState = OSSTrustNetworkValDNoOverclaimStatePartial
		}},
		{name: "non active vald current state blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.CurrentState = OSSTrustNetworkValDStatePartial
		}},
		{name: "val0 evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.Val0Source.EvidenceRefs[0] = "evidence:ostn-val0-unknown-001"
		}},
		{name: "vala evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValASource.EvidenceRefs[0] = "evidence:ostn-vala-unknown-001"
		}},
		{name: "valb evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValBSource.EvidenceRefs[0] = "evidence:ostn-valb-unknown-001"
		}},
		{name: "valc evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValCSource.EvidenceRefs[0] = "evidence:ostn-valc-unknown-001"
		}},
		{name: "vald proof refs mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.ProofSurfaceRefs = append(model.ValDSource.ProofSurfaceRefs, "/v1/oss-trust-network/vald/other")
		}},
		{name: "vald source state partial blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = OSSTrustNetworkValESourceStatePartial
		}},
		{name: "vald source state unknown blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = OSSTrustNetworkValESourceStateUnknown
		}},
		{name: "vald source state incomplete blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = OSSTrustNetworkValESourceStateIncomplete
		}},
		{name: "vald source state malformed blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = "active-ish"
		}},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValEClosure(model)
		if model.CurrentState == OSSTrustNetworkValEStatePass || model.Point9PassAllowed {
			t.Fatalf("%s: expected pass to remain blocked, got %#v", tc.name, model)
		}
		if model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
			t.Fatalf("%s: expected point 9 to remain not complete, got %#v", tc.name, model)
		}
	}
}

func TestOSSTrustNetworkValEPoint9PassReasonDisciplineAndNoOverclaim(t *testing.T) {
	testCases := []struct {
		name       string
		reason     string
		mutate     func(*OSSTrustNetworkValEIntegratedClosure)
		wantNoOver string
		wantPass   string
	}{
		{
			name:   "canonical blocked reason preserves fail closed without overclaim",
			reason: OSSTrustNetworkValEPoint9PassReasonBlocked,
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.ValASource.CurrentState = OSSTrustNetworkValAStatePartial
			},
			wantNoOver: OSSTrustNetworkValENoOverclaimStateActive,
			wantPass:   OSSTrustNetworkValEFinalPassRuleStatePartial,
		},
		{name: "exact allowed reason passes", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, wantNoOver: OSSTrustNetworkValENoOverclaimStateActive, wantPass: OSSTrustNetworkValEFinalPassRuleStateActive},
		{name: "blocked ineligible reason blocks", reason: "point_9_pass production approved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "spaced point 9 pass production approved blocks", reason: "point 9 pass production approved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "extended allowed reason with certified blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed + " certified", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "extended allowed reason with production approved blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed + " production approved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "extended allowed reason with global truth blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed + " global truth", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "substring bypass containing point_9_pass plus overclaim blocks", reason: "safe point_9_pass maybe later with public badge", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "changelock verified blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "changelock verified")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "regulator approved blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "regulator approved")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "regulator hyphen approved blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "regulator-approved")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "compliance guaranteed blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "compliance guaranteed")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "legal certification blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "legal certification")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "patent cleared blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "patent cleared")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "fto cleared blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "fto cleared")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "fto hyphen cleared blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "fto-cleared")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "de facto standard blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "de facto standard")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "de facto hyphen standard blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "de-facto standard")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "universal trust score blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "universal trust score")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "integrity score blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "integrity score")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "score over 90 blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "score > 90")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "reviewed means safe blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "reviewed means safe")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "community truth blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "community truth")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "network truth blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "network truth")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "crowd sourced truth layer blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "crowd sourced truth layer")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "crowd sourced hyphen truth layer blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "crowd-sourced truth layer")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "auto remediated blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "auto remediated")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "auto hyphen remediated blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "auto-remediated")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "auto merged blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "auto merged")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "auto hyphen merged blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "auto-merged")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "production autopatch blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "production autopatch")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "safe not canonical truth stays allowed", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = []string{"not canonical truth"}
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateActive, wantPass: OSSTrustNetworkValEFinalPassRuleStateActive},
		{name: "safe not formal certification stays allowed", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = []string{"not formal certification"}
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateActive, wantPass: OSSTrustNetworkValEFinalPassRuleStateActive},
		{name: "safe not production approval stays allowed", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = []string{"not production approval"}
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateActive, wantPass: OSSTrustNetworkValEFinalPassRuleStateActive},
		{name: "safe not deployment approval stays allowed", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = []string{"not deployment approval"}
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateActive, wantPass: OSSTrustNetworkValEFinalPassRuleStateActive},
		{name: "safe not official oss authority stays allowed", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = []string{"not official oss authority"}
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateActive, wantPass: OSSTrustNetworkValEFinalPassRuleStateActive},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		model.Point9PassReason = tc.reason
		if tc.mutate != nil {
			tc.mutate(&model)
		}
		model = ComputeOSSTrustNetworkValEClosure(model)
		if model.NoOverclaimState != tc.wantNoOver || model.FinalPassRuleState != tc.wantPass {
			t.Fatalf("%s: got no-overclaim=%q pass=%q model=%#v", tc.name, model.NoOverclaimState, model.FinalPassRuleState, model)
		}
	}
}

func TestOSSTrustNetworkValECanonicalBoundaryBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*OSSTrustNetworkValEIntegratedClosure)
	}{
		{name: "missing canonical boundary blocks", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.IntegratedClosure.CanonicalExecutionAuditEvidenceSourceOfTruth = false
		}},
		{name: "registry surface becoming authority blocks", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.CanonicalBoundary.RegistrySurfaceAuthorityClaim = true
		}},
		{name: "local enterprise override not evidence linked blocks", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValCSource.LocalOverride.EvidenceRefs = nil
		}},
		{name: "local enterprise override hidden blocks", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.CanonicalBoundary.LocalOverrideVisible = false
		}},
		{name: "shared signal silently overriding local enterprise decision blocks", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.CanonicalBoundary.SharedSignalSilentOverride = true
		}},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValEClosure(model)
		if model.CurrentState == OSSTrustNetworkValEStatePass || model.Point9PassAllowed {
			t.Fatalf("%s: expected pass to remain blocked, got %#v", tc.name, model)
		}
		if model.CanonicalBoundaryState != OSSTrustNetworkValECanonicalBoundaryStateBlocked &&
			model.IntegratedClosureState != OSSTrustNetworkValEIntegratedClosureStateBlocked {
			t.Fatalf("%s: expected canonical boundary or integrated closure to block, got %#v", tc.name, model)
		}
	}
}

func TestOSSTrustNetworkValEProofEvidenceQualityValidation(t *testing.T) {
	testCases := []struct {
		name     string
		evidence func() []ReferenceArchitectureEvidenceReference
		refs     func() []string
		want     bool
	}{
		{
			name: "exact evidence model passes",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: true,
		},
		{
			name: "duplicate evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].EvidenceID = evidence[1].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unknown evidence ref fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
			},
			refs: func() []string {
				refs := append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
				refs[0] = "evidence:ostn-vale-unknown-001"
				return refs
			},
			want: false,
		},
		{
			name: "whitespace evidence ref fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
			},
			refs: func() []string {
				refs := append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
				refs[0] = " "
				return refs
			},
			want: false,
		},
		{
			name: "wrong evidence type fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].EvidenceType = "other_type"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong source fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].Source = "oss-trust-network/vale/unrelated"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong scope fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].Scope = "other_scope"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "fresh but unrelated evidence payload fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].EvidenceID = "evidence:ostn-vale-unrelated-001"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "stale evidence fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].FreshnessState = IntelligenceCalibrationFreshnessStale
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "expired evidence fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].FreshnessState = IntelligenceCalibrationFreshnessExpired
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unsupported evidence fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].FreshnessState = IntelligenceCalibrationFreshnessUnsupported
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unknown evidence freshness fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].FreshnessState = IntelligenceCalibrationFreshnessUnknown
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := OSSTrustNetworkValEProofEvidenceQualityValid(tc.evidence(), tc.refs()); got != tc.want {
			t.Fatalf("%s: expected %t, got %t", tc.name, tc.want, got)
		}
	}
}
