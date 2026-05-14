package operability

import (
	"strings"
	"testing"
)

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

func TestOSSTrustNetworkValEProjectionDisclaimerExactBoundedBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*OSSTrustNetworkValEIntegratedClosure)
		eval   func(OSSTrustNetworkValEIntegratedClosure) string
		want   string
	}{
		{
			name: "dependency leading whitespace blocks live dependency gate",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.ProjectionDisclaimer = " " + ossTrustNetworkValEProjectionDisclaimer()
			},
			eval: EvaluateOSSTrustNetworkValEDependencyState,
			want: OSSTrustNetworkValEDependencyStateUnknown,
		},
		{
			name: "dependency uppercase retagging blocks live dependency gate",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.ProjectionDisclaimer = strings.ToUpper(ossTrustNetworkValEProjectionDisclaimer())
			},
			eval: EvaluateOSSTrustNetworkValEDependencyState,
			want: OSSTrustNetworkValEDependencyStateUnknown,
		},
		{
			name: "integrated closure aggregate snapshot blocks live gate",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.IntegratedClosure.ProjectionDisclaimer = ossTrustNetworkValEProjectionDisclaimer() + " aggregate_dependency_snapshot"
			},
			eval: EvaluateOSSTrustNetworkValEIntegratedClosureState,
			want: OSSTrustNetworkValEIntegratedClosureStateUnknown,
		},
		{
			name: "no-overclaim aggregate snapshot blocks live gate",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.NoOverclaim.ProjectionDisclaimer = ossTrustNetworkValEProjectionDisclaimer() + " aggregate_dependency_snapshot"
			},
			eval: EvaluateOSSTrustNetworkValENoOverclaimState,
			want: OSSTrustNetworkValENoOverclaimStateUnknown,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		tc.mutate(&model)
		if got := tc.eval(model); got != tc.want {
			t.Fatalf("%s: expected %s, got %s with %#v", tc.name, tc.want, got, model)
		}
	}
}

func TestOSSTrustNetworkValESourceProjectionDisclaimerExactBoundedBlockers(t *testing.T) {
	testCases := []struct {
		name string
		eval func() string
		want string
	}{
		{
			name: "val0 source leading whitespace disclaimer blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.Val0Source.ProjectionDisclaimer = " " + ossTrustNetworkVal0ProjectionDisclaimer()
				return EvaluateOSSTrustNetworkValEVal0SourceState(model.Val0Source)
			},
			want: OSSTrustNetworkValESourceStateUnknown,
		},
		{
			name: "vala source uppercase disclaimer blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValASource.ProjectionDisclaimer = strings.ToUpper(ossTrustNetworkValAProjectionDisclaimer())
				return EvaluateOSSTrustNetworkValEValASourceState(model.ValASource)
			},
			want: OSSTrustNetworkValESourceStateUnknown,
		},
		{
			name: "valb source trailing whitespace disclaimer blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValBSource.ProjectionDisclaimer = ossTrustNetworkValBProjectionDisclaimer() + " "
				return EvaluateOSSTrustNetworkValEValBSourceState(model.ValBSource)
			},
			want: OSSTrustNetworkValESourceStateUnknown,
		},
		{
			name: "valc source uppercase disclaimer blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValCSource.ProjectionDisclaimer = strings.ToUpper(ossTrustNetworkValCProjectionDisclaimer())
				return EvaluateOSSTrustNetworkValEValCSourceState(model.ValCSource)
			},
			want: OSSTrustNetworkValESourceStateUnknown,
		},
		{
			name: "vald source leading whitespace disclaimer blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValDSource.ProjectionDisclaimer = " " + ossTrustNetworkValDProjectionDisclaimer()
				return EvaluateOSSTrustNetworkValEValDSourceState(model.ValDSource)
			},
			want: OSSTrustNetworkValESourceStateUnknown,
		},
	}

	for _, tc := range testCases {
		if got := tc.eval(); got != tc.want {
			t.Fatalf("%s: expected %s, got %s", tc.name, tc.want, got)
		}
	}
}

func TestOSSTrustNetworkValESourceCanonicalRefsRequireRawExactBinding(t *testing.T) {
	testCases := []struct {
		name string
		eval func() string
	}{
		{
			name: "val0 source leading whitespace evidence ref blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.Val0Source.EvidenceRefs[0] = " " + OSSTrustNetworkVal0ProofEvidenceRefs()[0]
				return EvaluateOSSTrustNetworkValEVal0SourceState(model.Val0Source)
			},
		},
		{
			name: "vala source trailing whitespace proof ref blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValASource.ProofSurfaceRefs[0] = OSSTrustNetworkValAProofSurfaceRefs()[0] + " "
				return EvaluateOSSTrustNetworkValEValASourceState(model.ValASource)
			},
		},
		{
			name: "valb source tab padded evidence ref blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValBSource.EvidenceRefs[0] = "\t" + OSSTrustNetworkValBProofEvidenceRefs()[0]
				return EvaluateOSSTrustNetworkValEValBSourceState(model.ValBSource)
			},
		},
		{
			name: "valc source newline padded proof ref blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValCSource.ProofSurfaceRefs[0] = OSSTrustNetworkValCProofSurfaceRefs()[0] + "\n"
				return EvaluateOSSTrustNetworkValEValCSourceState(model.ValCSource)
			},
		},
		{
			name: "vald source leading whitespace evidence ref blocks",
			eval: func() string {
				model := activeOSSTrustNetworkValEModel()
				model.ValDSource.EvidenceRefs[0] = " " + OSSTrustNetworkValDProofEvidenceRefs()[0]
				return EvaluateOSSTrustNetworkValEValDSourceState(model.ValDSource)
			},
		},
	}

	for _, tc := range testCases {
		if got := tc.eval(); got != OSSTrustNetworkValESourceStateBlocked {
			t.Fatalf("%s: expected blocked source state, got %s", tc.name, got)
		}
	}
}

func TestOSSTrustNetworkValECanonicalRefsRequireRawExactBinding(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*OSSTrustNetworkValEIntegratedClosure)
		eval   func(OSSTrustNetworkValEIntegratedClosure) string
		want   string
	}{
		{
			name: "dependency padded vald proof ref blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.ValDSource.ProofSurfaceRefs[0] = " " + OSSTrustNetworkValDProofSurfaceRefs()[0]
			},
			eval: EvaluateOSSTrustNetworkValEDependencyState,
			want: OSSTrustNetworkValEDependencyStateBlocked,
		},
		{
			name: "integrated closure padded evidence ref blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.IntegratedClosure.EvidenceRefs[0] = OSSTrustNetworkValEProofEvidenceRefs()[6] + "\n"
			},
			eval: EvaluateOSSTrustNetworkValEIntegratedClosureState,
			want: OSSTrustNetworkValEIntegratedClosureStateBlocked,
		},
		{
			name: "integrated closure padded gate id blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.IntegratedClosure.GateID = " " + OSSTrustNetworkValEIntegratedClosureGateModel().GateID
			},
			eval: EvaluateOSSTrustNetworkValEIntegratedClosureState,
			want: OSSTrustNetworkValEIntegratedClosureStateBlocked,
		},
		{
			name: "integrated closure padded version blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.IntegratedClosure.Version = OSSTrustNetworkValEIntegratedClosureGateModel().Version + "\n"
			},
			eval: EvaluateOSSTrustNetworkValEIntegratedClosureState,
			want: OSSTrustNetworkValEIntegratedClosureStateBlocked,
		},
		{
			name: "canonical boundary padded evidence ref blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.CanonicalBoundary.EvidenceRefs[0] = " " + "evidence:ostn-vale-canonical-boundary-001"
			},
			eval: EvaluateOSSTrustNetworkValECanonicalBoundaryState,
			want: OSSTrustNetworkValECanonicalBoundaryStateBlocked,
		},
		{
			name: "canonical boundary padded boundary id blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.CanonicalBoundary.BoundaryID = "\t" + OSSTrustNetworkValECanonicalBoundaryGateModel().BoundaryID + "\n"
			},
			eval: EvaluateOSSTrustNetworkValECanonicalBoundaryState,
			want: OSSTrustNetworkValECanonicalBoundaryStateBlocked,
		},
		{
			name: "evidence quality padded proof ref blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.EvidenceQuality.ProofSurfaceRefs[0] = "\t" + OSSTrustNetworkValEProofSurfaceRefs()[0]
			},
			eval: func(model OSSTrustNetworkValEIntegratedClosure) string {
				return EvaluateOSSTrustNetworkValEEvidenceQualityState(model.EvidenceQuality)
			},
			want: OSSTrustNetworkValEEvidenceQualityStateBlocked,
		},
		{
			name: "evidence quality padded dependency evidence ref blocks",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.EvidenceQuality.DependencyEvidenceRefs[0] = OSSTrustNetworkValDProofEvidenceRefs()[0] + " "
			},
			eval: func(model OSSTrustNetworkValEIntegratedClosure) string {
				return EvaluateOSSTrustNetworkValEEvidenceQualityState(model.EvidenceQuality)
			},
			want: OSSTrustNetworkValEEvidenceQualityStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		tc.mutate(&model)
		if got := tc.eval(model); got != tc.want {
			t.Fatalf("%s: expected %s, got %s with %#v", tc.name, tc.want, got, model)
		}
	}
}

func TestOSSTrustNetworkValEDependencyAndSourceBlockers(t *testing.T) {
	testCases := []struct {
		name           string
		mutate         func(*OSSTrustNetworkValEIntegratedClosure)
		wantCurrent    string
		wantVal0Source string
		wantValASource string
		wantValBSource string
		wantValCSource string
		wantValDSource string
		wantDependency string
	}{
		{name: "missing vald dependency blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.DependencyState = OSSTrustNetworkValDDependencyStatePartial
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "route presence alone cannot satisfy vald dependency", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.CurrentState = ""
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "vald malformed evidence identity blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.EvidenceRefs[0] = "evidence:ostn-vald-unknown-001"
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "vald non active subgate blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.NoOverclaimState = OSSTrustNetworkValDNoOverclaimStatePartial
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "non active vald current state blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.CurrentState = OSSTrustNetworkValDStatePartial
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "val0 evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.Val0Source.EvidenceRefs[0] = "evidence:ostn-val0-unknown-001"
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantVal0Source: OSSTrustNetworkValESourceStateBlocked},
		{name: "vala evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValASource.EvidenceRefs[0] = "evidence:ostn-vala-unknown-001"
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValASource: OSSTrustNetworkValESourceStateBlocked},
		{name: "valb evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValBSource.EvidenceRefs[0] = "evidence:ostn-valb-unknown-001"
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValBSource: OSSTrustNetworkValESourceStateBlocked},
		{name: "valc evidence mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValCSource.EvidenceRefs[0] = "evidence:ostn-valc-unknown-001"
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValCSource: OSSTrustNetworkValESourceStateBlocked},
		{name: "vald proof refs mismatch blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.ProofSurfaceRefs = append(model.ValDSource.ProofSurfaceRefs, "/v1/oss-trust-network/vald/other")
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "vald source state partial blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = OSSTrustNetworkValESourceStatePartial
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "vald source state unknown blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = OSSTrustNetworkValESourceStateUnknown
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "vald source state incomplete blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = OSSTrustNetworkValESourceStateIncomplete
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "vald source state malformed blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = "active-ish"
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
		{name: "whitespace retagged claimed vald source state blocks pass", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSourceState = " " + OSSTrustNetworkValESourceStateActive + " "
		}, wantCurrent: OSSTrustNetworkValEStateBlocked, wantValDSource: OSSTrustNetworkValESourceStateBlocked, wantDependency: OSSTrustNetworkValEDependencyStateBlocked},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValEClosure(model)
		if model.CurrentState != tc.wantCurrent || model.Point9PassAllowed || model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
			t.Fatalf("%s: expected current=%q not-complete/no-pass, got %#v", tc.name, tc.wantCurrent, model)
		}
		if tc.wantVal0Source != "" && model.Val0SourceState != tc.wantVal0Source {
			t.Fatalf("%s: expected val0 source state %q, got %#v", tc.name, tc.wantVal0Source, model)
		}
		if tc.wantValASource != "" && model.ValASourceState != tc.wantValASource {
			t.Fatalf("%s: expected vala source state %q, got %#v", tc.name, tc.wantValASource, model)
		}
		if tc.wantValBSource != "" && model.ValBSourceState != tc.wantValBSource {
			t.Fatalf("%s: expected valb source state %q, got %#v", tc.name, tc.wantValBSource, model)
		}
		if tc.wantValCSource != "" && model.ValCSourceState != tc.wantValCSource {
			t.Fatalf("%s: expected valc source state %q, got %#v", tc.name, tc.wantValCSource, model)
		}
		if tc.wantValDSource != "" && model.ValDSourceState != tc.wantValDSource {
			t.Fatalf("%s: expected vald source state %q, got %#v", tc.name, tc.wantValDSource, model)
		}
		if tc.wantDependency != "" && model.DependencyState != tc.wantDependency {
			t.Fatalf("%s: expected dependency state %q, got %#v", tc.name, tc.wantDependency, model)
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
			wantPass:   OSSTrustNetworkValEFinalPassRuleStateBlocked,
		},
		{name: "exact allowed reason passes", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, wantNoOver: OSSTrustNetworkValENoOverclaimStateActive, wantPass: OSSTrustNetworkValEFinalPassRuleStateActive},
		{name: "blocked ineligible reason blocks", reason: "point_9_pass production approved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "blocked ineligible production approval reason blocks", reason: "point_9_pass production approval", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "blocked ineligible deployment approval reason blocks", reason: "point_9_pass deployment approval", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "spaced point 9 pass production approved blocks", reason: "point 9 pass production approved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "confusable production approved blocks", reason: "prοductiοn apprοved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "open-o confusable production approved blocks", reason: "prɔduction apprɔved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "extended allowed reason with certified blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed + " certified", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "extended allowed reason with production approved blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed + " production approved", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "extended allowed reason with global truth blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed + " global truth", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "split production approved across reason and observed claim blocks", reason: "production", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "approved")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "split point9 pass global truth across reason and observed claim blocks", reason: "point_9_pass global", mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "truth")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "substring bypass containing point_9_pass plus overclaim blocks", reason: "safe point_9_pass maybe later with public badge", wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "confusable public badge blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "public bᴀdge")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "confusable certified package blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "certifiᴄed package")
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
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
		{name: "whitespace retagged val0 source no-overclaim blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.Val0Source.NoOverclaimState = " " + OSSTrustNetworkVal0NoOverclaimStateActive + " "
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
		{name: "newline retagged vald source no-overclaim blocks", reason: OSSTrustNetworkValEPoint9PassReasonAllowed, mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
			model.ValDSource.NoOverclaimState = OSSTrustNetworkValDNoOverclaimStateActive + "\n"
		}, wantNoOver: OSSTrustNetworkValENoOverclaimStateBlocked, wantPass: OSSTrustNetworkValEFinalPassRuleStateBlocked},
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
		if tc.wantPass == OSSTrustNetworkValEFinalPassRuleStateActive {
			if model.CurrentState != OSSTrustNetworkValEStatePass || !model.Point9PassAllowed {
				t.Fatalf("%s: expected exact pass current state and point9 pass allowed, got %#v", tc.name, model)
			}
			continue
		}
		if model.CurrentState != OSSTrustNetworkValEStateBlocked || model.Point9PassAllowed {
			t.Fatalf("%s: expected exact blocked current state and no point9 pass, got %#v", tc.name, model)
		}
	}
}

func TestOSSTrustNetworkValEPoint9PassReasonRequiresRawExactBinding(t *testing.T) {
	testCases := []struct {
		name       string
		reason     string
		wantPass   bool
		wantReason string
	}{
		{
			name:       "allowed reason padded with spaces blocks",
			reason:     " " + OSSTrustNetworkValEPoint9PassReasonAllowed + " ",
			wantPass:   false,
			wantReason: OSSTrustNetworkValEPoint9PassReasonBlocked,
		},
		{
			name:       "allowed reason padded with tab newline blocks",
			reason:     "\t" + OSSTrustNetworkValEPoint9PassReasonAllowed + "\n",
			wantPass:   false,
			wantReason: OSSTrustNetworkValEPoint9PassReasonBlocked,
		},
		{
			name:       "allowed reason uppercased blocks",
			reason:     strings.ToUpper(OSSTrustNetworkValEPoint9PassReasonAllowed),
			wantPass:   false,
			wantReason: OSSTrustNetworkValEPoint9PassReasonBlocked,
		},
		{
			name:       "blocked reason padded with spaces stays blocked canonically",
			reason:     " " + OSSTrustNetworkValEPoint9PassReasonBlocked + " ",
			wantPass:   false,
			wantReason: OSSTrustNetworkValEPoint9PassReasonBlocked,
		},
		{
			name:       "exact safe diagnostic still promotes only through active closure",
			reason:     OSSTrustNetworkValEPoint9PassSafeDiagnosticValDCannotReturn,
			wantPass:   true,
			wantReason: OSSTrustNetworkValEPoint9PassReasonBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		model.Point9PassReason = tc.reason
		model = ComputeOSSTrustNetworkValEClosure(model)
		if tc.wantPass {
			if model.FinalPassRuleState != OSSTrustNetworkValEFinalPassRuleStateActive ||
				model.CurrentState != OSSTrustNetworkValEStatePass ||
				!model.Point9PassAllowed ||
				model.Point9State != OSSTrustNetworkPoint9StatePass {
				t.Fatalf("%s: expected exact pass current state and point9 pass, got %#v", tc.name, model)
			}
			if model.Point9PassReason != OSSTrustNetworkValEPoint9PassReasonAllowed {
				t.Fatalf("%s: expected canonical allowed pass reason, got %#v", tc.name, model)
			}
			continue
		}
		if model.FinalPassRuleState != OSSTrustNetworkValEFinalPassRuleStateBlocked ||
			model.CurrentState != OSSTrustNetworkValEStateBlocked ||
			model.Point9PassAllowed ||
			model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
			t.Fatalf("%s: expected exact blocked current state and no point9 pass, got %#v", tc.name, model)
		}
		if model.Point9PassReason != tc.wantReason {
			t.Fatalf("%s: expected canonical blocked pass reason %q, got %#v", tc.name, tc.wantReason, model)
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
		if model.CurrentState != OSSTrustNetworkValEStateBlocked || model.Point9PassAllowed {
			t.Fatalf("%s: expected exact blocked current state and no pass, got %#v", tc.name, model)
		}
		switch tc.name {
		case "missing canonical boundary blocks":
			if model.IntegratedClosureState != OSSTrustNetworkValEIntegratedClosureStateBlocked {
				t.Fatalf("%s: expected integrated closure blocked, got %#v", tc.name, model)
			}
		default:
			if model.CanonicalBoundaryState != OSSTrustNetworkValECanonicalBoundaryStateBlocked {
				t.Fatalf("%s: expected canonical boundary blocked, got %#v", tc.name, model)
			}
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
			name: "whitespace padded evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].EvidenceID = " " + evidence[0].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "whitespace padded evidence type fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].EvidenceType = evidence[0].EvidenceType + " "
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "tab padded source fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].Source = "\t" + evidence[0].Source
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValEProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "newline padded scope fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := ossTrustNetworkValECopyEvidence(ossTrustNetworkValEEvidence())
				evidence[0].Scope = evidence[0].Scope + "\n"
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

func TestOSSTrustNetworkValEEvidenceQualityMetadataRequiresRawExactBinding(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*OSSTrustNetworkValEIntegratedClosure)
	}{
		{
			name: "whitespace padded evidence id blocks evidence quality",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.EvidenceQuality.Evidence[0].EvidenceID = " " + model.EvidenceQuality.Evidence[0].EvidenceID
			},
		},
		{
			name: "whitespace padded evidence type blocks evidence quality",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.EvidenceQuality.Evidence[0].EvidenceType = model.EvidenceQuality.Evidence[0].EvidenceType + " "
			},
		},
		{
			name: "tab padded source blocks evidence quality",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.EvidenceQuality.Evidence[0].Source = "\t" + model.EvidenceQuality.Evidence[0].Source
			},
		},
		{
			name: "newline padded scope blocks evidence quality",
			mutate: func(model *OSSTrustNetworkValEIntegratedClosure) {
				model.EvidenceQuality.Evidence[0].Scope = model.EvidenceQuality.Evidence[0].Scope + "\n"
			},
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValEModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValEClosure(model)
		if model.EvidenceQualityState != OSSTrustNetworkValEEvidenceQualityStateBlocked || model.CurrentState != OSSTrustNetworkValEStateBlocked || model.Point9PassAllowed {
			t.Fatalf("%s: expected blocked evidence quality and blocked top-level state with no point9 pass, got %#v", tc.name, model)
		}
	}
}
