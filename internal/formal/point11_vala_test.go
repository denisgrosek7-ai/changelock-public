package formal

import (
	"encoding/json"
	"strings"
	"testing"
)

func point11ValAActiveDependencySnapshot() Point11ValADependencySnapshot {
	val0 := activePoint11Val0Foundation()
	return SnapshotPoint11ValADependencyFromComputedVal0(val0, Point11ValAVal0ReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
}

func activePoint11ValAFoundation() Point11ValAFoundation {
	model := Point11ValAFoundationModel()
	model.Dependency = point11ValAActiveDependencySnapshot()
	return ComputePoint11ValAFoundation(model)
}

func reviewRequiredPoint11ValADependencySnapshot() Point11ValADependencySnapshot {
	val0 := activePoint11Val0Foundation()
	val0.CurrentState = Point11Val0StateReviewRequired
	val0.DependencyState = Point11Val0DependencyStateReviewRequired
	val0.ReviewPrerequisites = []string{"val0_repo_visibility_review_prerequisite"}
	return SnapshotPoint11ValADependencyFromComputedVal0(val0, Point11ValAVal0ReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
}

func point11ValAFoundationWithValidSupersededTransition() Point11ValAFoundation {
	model := activePoint11ValAFoundation()
	model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
	model.Lifecycle.ToState = point11ValAPolicyLifecycleSuperseded
	model.Lifecycle.GovernanceEventRef = "governance_event_point11_vala_supersession"
	model.Lifecycle.Reason = "policy_supersession_rollout"
	model.Lifecycle.ApprovalEvidenceRefs = []string{"evidence:point11-vala-supersession-001"}
	model.Lifecycle.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
	model.Lifecycle.SuccessorPolicyRef = "policy_successor_2026_05_02"
	model.Graph.SuccessorPolicyRef = model.Lifecycle.SuccessorPolicyRef
	model.Graph.SupersessionReason = "policy_supersession_rollout"
	model.Graph.CompatibilityVersion = "point11_vala_compat_v1"
	model.Graph.CompatibilityReviewRef = model.Lifecycle.CompatibilityReviewRef
	model.Graph.SuccessorLifecycleState = point11ValAPolicyLifecycleActive
	model.Graph.LineagePath = []string{model.Graph.SourcePolicyRef, model.Graph.SuccessorPolicyRef}
	return ComputePoint11ValAFoundation(model)
}

func point11ValAFoundationWithValidRevokedTransition() Point11ValAFoundation {
	model := activePoint11ValAFoundation()
	model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
	model.Lifecycle.ToState = point11ValAPolicyLifecycleRevoked
	model.Lifecycle.GovernanceEventRef = "governance_event_point11_vala_revocation"
	model.Lifecycle.ApproverRef = "governance_final_approver"
	model.Lifecycle.Reason = "policy_revocation_required"
	model.Lifecycle.ApprovalEvidenceRefs = []string{"evidence:point11-vala-revocation-001"}
	model.Graph.RevokedByRef = model.Lifecycle.GovernanceEventRef
	model.Graph.RevocationReason = "policy_revocation_required"
	return ComputePoint11ValAFoundation(model)
}

func TestPoint11ValADependencyState(t *testing.T) {
	t.Run("happy path val0 dependency active", func(t *testing.T) {
		snapshot := point11ValAActiveDependencySnapshot()
		if got := EvaluatePoint11ValADependencyState(snapshot); got != Point11ValADependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", snapshot)
		}
	})

	t.Run("copies aggregate val0 projection disclaimer exactly", func(t *testing.T) {
		val0 := activePoint11Val0Foundation()
		val0.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_val0_disclaimer"
		val0.PolicyContract.ProjectionDisclaimer = "projection_only not_canonical_truth component_policy_contract_disclaimer"
		snapshot := SnapshotPoint11ValADependencyFromComputedVal0(val0, Point11ValAVal0ReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
		})
		if snapshot.ProjectionDisclaimer != val0.ProjectionDisclaimer {
			t.Fatalf("expected aggregate val0 projection disclaimer, got snapshot=%q val0=%q", snapshot.ProjectionDisclaimer, val0.ProjectionDisclaimer)
		}
		if got := EvaluatePoint11ValADependencyState(snapshot); got != Point11ValADependencyStateActive {
			t.Fatalf("expected active dependency with propagated disclaimer, got %#v", snapshot)
		}
	})

	t.Run("malformed aggregate val0 projection disclaimer blocks even if component disclaimer looks valid", func(t *testing.T) {
		val0 := activePoint11Val0Foundation()
		val0.ProjectionDisclaimer = "canonical_truth"
		val0.PolicyContract.ProjectionDisclaimer = "projection_only not_canonical_truth component_policy_contract_disclaimer"
		snapshot := SnapshotPoint11ValADependencyFromComputedVal0(val0, Point11ValAVal0ReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
		})
		if snapshot.ProjectionDisclaimer != val0.ProjectionDisclaimer {
			t.Fatalf("expected malformed aggregate disclaimer to propagate without fallback, got snapshot=%q val0=%q", snapshot.ProjectionDisclaimer, val0.ProjectionDisclaimer)
		}
		if got := EvaluatePoint11ValADependencyState(snapshot); got != Point11ValADependencyStateBlocked {
			t.Fatalf("expected malformed aggregate disclaimer to block dependency, got %#v", snapshot)
		}
	})

	t.Run("valid aggregate disclaimer wins over differing component disclaimer", func(t *testing.T) {
		val0 := activePoint11Val0Foundation()
		val0.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_val0_disclaimer"
		val0.PolicyContract.ProjectionDisclaimer = "canonical_truth"
		snapshot := SnapshotPoint11ValADependencyFromComputedVal0(val0, Point11ValAVal0ReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
		})
		if snapshot.ProjectionDisclaimer != val0.ProjectionDisclaimer {
			t.Fatalf("expected snapshot to use aggregate disclaimer instead of component fallback, got snapshot=%q val0=%q", snapshot.ProjectionDisclaimer, val0.ProjectionDisclaimer)
		}
		if snapshot.ProjectionDisclaimer == val0.PolicyContract.ProjectionDisclaimer {
			t.Fatalf("expected snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", snapshot.ProjectionDisclaimer, val0.PolicyContract.ProjectionDisclaimer)
		}
		if got := EvaluatePoint11ValADependencyState(snapshot); got != Point11ValADependencyStateActive {
			t.Fatalf("expected active dependency when aggregate disclaimer is valid, got %#v", snapshot)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValADependencySnapshot)
		want   string
	}{
		{name: "malformed val0 projection disclaimer blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.ProjectionDisclaimer = "canonical_truth"
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 policy contract blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0PolicyContractState = Point11Val0PolicyContractStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 claim governance blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0ClaimGovernanceState = Point11Val0ClaimGovernanceStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 authority matrix blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0AuthorityMatrixState = Point11Val0AuthorityMatrixStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 exception governance blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0ExceptionGovernanceState = Point11Val0ExceptionGovernanceStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 abac blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0ABACState = Point11Val0ABACStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 decision binding blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0DecisionBindingState = Point11Val0DecisionBindingStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 no overclaim blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0NoOverclaimState = Point11Val0NoOverclaimStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "blocked val0 cross domain compatibility blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0CrossDomainCompatibilityState = Point11Val0CrossDomainCompatibilityStateBlocked
		}, want: Point11ValADependencyStateBlocked},
		{name: "val0 point11 pass emission marker blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0Point11PassEmitted = true
		}, want: Point11ValADependencyStateBlocked},
		{name: "val0 authority marker blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0CreatesAuthorityClaims = true
		}, want: Point11ValADependencyStateBlocked},
		{name: "val0 publication side effect marker blocks", mutate: func(model *Point11ValADependencySnapshot) {
			model.Val0CreatesPublicationSideEffects = true
		}, want: Point11ValADependencyStateBlocked},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := point11ValAActiveDependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint11ValADependencyState(model); got != testCase.want {
				t.Fatalf("expected dependency state %q, got %q for %#v", testCase.want, got, model)
			}
		})
	}
}

func TestPoint11ValARegistryState(t *testing.T) {
	t.Run("valid signed anchored policy registry entry can become active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		if model.RegistryState != Point11ValARegistryStateActive {
			t.Fatalf("expected active registry state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValASignedPolicyRegistry)
	}{
		{name: "missing registry id blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.RegistryID = "" }},
		{name: "missing policy pack id blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.PolicyPackID = "" }},
		{name: "missing policy id blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.PolicyID = "" }},
		{name: "global policy scope blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.PolicyScope = "global_policy_scope" }},
		{name: "missing signature ref blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SignatureRef = "" }},
		{name: "signature unknown blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SignatureRef = "signature_unknown" }},
		{name: "signature revoked blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SignatureRef = "signature_revoked" }},
		{name: "signature invalid blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SignatureRef = "signature_invalid" }},
		{name: "missing signing key ref blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SigningKeyRef = "" }},
		{name: "signing key unknown blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SigningKeyRef = "signing_key_unknown" }},
		{name: "signing key revoked blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SigningKeyRef = "signing_key_revoked" }},
		{name: "signing key invalid blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SigningKeyRef = "signing_key_invalid" }},
		{name: "unsupported signing algorithm blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.SigningAlgorithm = "rsa4096" }},
		{name: "missing anchor ref blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.AnchorRef = "" }},
		{name: "anchor unknown blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.AnchorRef = "anchor_unknown" }},
		{name: "anchor revoked blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.AnchorRef = "anchor_revoked" }},
		{name: "anchor invalid blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.AnchorRef = "anchor_invalid" }},
		{name: "revoked lifecycle blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.LifecycleState = point11ValAPolicyLifecycleRevoked }},
		{name: "expired lifecycle blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.LifecycleState = point11ValAPolicyLifecycleExpired }},
		{name: "superseded lifecycle without valid successor and compatibility context blocks", mutate: func(model *Point11ValASignedPolicyRegistry) {
			model.LifecycleState = point11ValAPolicyLifecycleSuperseded
			model.SupersededBy = ""
			model.CompatibilityVersion = ""
		}},
		{name: "expired effective until blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.EffectiveUntil = "2000-01-01T00:00:00Z" }},
		{name: "missing approval evidence refs blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.ApprovalEvidenceRefs = nil }},
		{name: "missing governance event ref blocks", mutate: func(model *Point11ValASignedPolicyRegistry) { model.GovernanceEventRef = "" }},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValAFoundation()
			testCase.mutate(&model.Registry)
			model = ComputePoint11ValAFoundation(model)
			if model.RegistryState != Point11ValARegistryStateBlocked {
				t.Fatalf("expected blocked registry state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValASignatureState(t *testing.T) {
	t.Run("valid signature envelope active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		if model.SignatureState != Point11ValASignatureStateActive {
			t.Fatalf("expected active signature state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValAPolicySignatureEnvelope)
	}{
		{name: "missing signature ref blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignatureRef = "" }},
		{name: "signature ref unknown blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignatureRef = "signature_unknown" }},
		{name: "signature ref revoked blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignatureRef = "signature_revoked" }},
		{name: "signature ref invalid blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignatureRef = "signature_invalid" }},
		{name: "signing key ref unknown blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SigningKeyRef = "signing_key_unknown" }},
		{name: "signed subject hash missing blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignedSubjectHash = "" }},
		{name: "signed subject hash malformed blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignedSubjectHash = "sha256:xyz" }},
		{name: "expired signature blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.ExpiresAt = "2000-01-01T00:00:00Z" }},
		{name: "revoked signature state blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignatureState = "revoked_signature_state" }},
		{name: "unknown signature state blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignatureState = "unknown_signature_state" }},
		{name: "unsupported signature state blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.SignatureState = "unsupported_signature_state" }},
		{name: "verification result not active blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.VerificationResult = "verification_pending" }},
		{name: "missing verification evidence refs blocks", mutate: func(model *Point11ValAPolicySignatureEnvelope) { model.VerificationEvidenceRefs = nil }},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValAFoundation()
			testCase.mutate(&model.Signature)
			model = ComputePoint11ValAFoundation(model)
			if model.SignatureState != Point11ValASignatureStateBlocked {
				t.Fatalf("expected blocked signature state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValAAnchorState(t *testing.T) {
	t.Run("valid anchor envelope active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		if model.AnchorState != Point11ValAAnchorStateActive {
			t.Fatalf("expected active anchor state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValAPolicyAnchorEnvelope)
	}{
		{name: "missing anchor ref blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorRef = "" }},
		{name: "anchor ref unknown blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorRef = "anchor_unknown" }},
		{name: "anchor ref revoked blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorRef = "anchor_revoked" }},
		{name: "unsupported anchor type blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorType = "external_blockchain" }},
		{name: "anchored subject hash missing blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchoredSubjectHash = "" }},
		{name: "anchored subject hash malformed blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchoredSubjectHash = "sha256:xyz" }},
		{name: "invalid anchor timestamp blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorTimestamp = "not-a-timestamp" }},
		{name: "revoked anchor state blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorState = "revoked_anchor_state" }},
		{name: "unknown anchor state blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorState = "unknown_anchor_state" }},
		{name: "anchor verification result not active blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorVerificationResult = "verification_pending" }},
		{name: "missing anchor evidence refs blocks", mutate: func(model *Point11ValAPolicyAnchorEnvelope) { model.AnchorEvidenceRefs = nil }},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValAFoundation()
			testCase.mutate(&model.Anchor)
			model = ComputePoint11ValAFoundation(model)
			if model.AnchorState != Point11ValAAnchorStateBlocked {
				t.Fatalf("expected blocked anchor state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValALifecycleState(t *testing.T) {
	t.Run("valid review required to approved transition is transition active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleReviewRequired
		model.Lifecycle.ToState = point11ValAPolicyLifecycleApproved
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateActive {
			t.Fatalf("expected active transition state, got %#v", model)
		}
		if model.PolicyUseState != Point11ValAPolicyUseStateNotYetActive {
			t.Fatalf("expected policy use not yet active, got %#v", model)
		}
	})

	t.Run("review required to approved without approver blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleReviewRequired
		model.Lifecycle.ToState = point11ValAPolicyLifecycleApproved
		model.Lifecycle.ApproverRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("review required to approved without governance event ref blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleReviewRequired
		model.Lifecycle.ToState = point11ValAPolicyLifecycleApproved
		model.Lifecycle.GovernanceEventRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("approved to active with active signature and anchor is transition active and policy use active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateActive || model.PolicyUseState != Point11ValAPolicyUseStateActive {
			t.Fatalf("expected active transition and policy use, got %#v", model)
		}
	})

	t.Run("approved to active with blocked signature blocks transition or active use", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.SignatureEnvelopeState = Point11ValASignatureStateBlocked
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked || model.PolicyUseState != Point11ValAPolicyUseStateBlocked {
			t.Fatalf("expected blocked transition and policy use, got %#v", model)
		}
	})

	t.Run("approved to active with blocked anchor blocks transition or active use", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.AnchorEnvelopeState = Point11ValAAnchorStateBlocked
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked || model.PolicyUseState != Point11ValAPolicyUseStateBlocked {
			t.Fatalf("expected blocked transition and policy use, got %#v", model)
		}
	})

	t.Run("valid active to deprecated is transition active but policy use is not active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleDeprecated
		model.Lifecycle.GovernanceEventRef = "governance_event_point11_vala_deprecation"
		model.Lifecycle.Reason = "policy_deprecation_reason"
		model.Lifecycle.ApprovalEvidenceRefs = []string{"evidence:point11-vala-deprecation-001"}
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateActive {
			t.Fatalf("expected active deprecated transition, got %#v", model)
		}
		if model.PolicyUseState != Point11ValAPolicyUseStateHistoricalOnly {
			t.Fatalf("expected historical only policy use, got %#v", model)
		}
	})

	t.Run("active to deprecated without governance event blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleDeprecated
		model.Lifecycle.GovernanceEventRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("valid active to superseded is transition active but old policy use is historical only", func(t *testing.T) {
		model := point11ValAFoundationWithValidSupersededTransition()
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateActive {
			t.Fatalf("expected active superseded transition, got %#v", model)
		}
		if model.PolicyUseState != Point11ValAPolicyUseStateHistoricalOnly {
			t.Fatalf("expected historical only policy use, got %#v", model)
		}
	})

	t.Run("active to superseded without successor policy ref blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleSuperseded
		model.Lifecycle.SuccessorPolicyRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("active to superseded without compatibility review blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleSuperseded
		model.Lifecycle.SuccessorPolicyRef = "policy_successor_2026_05_02"
		model.Lifecycle.CompatibilityReviewRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("active to superseded without governance event blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleSuperseded
		model.Lifecycle.SuccessorPolicyRef = "policy_successor_2026_05_02"
		model.Lifecycle.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
		model.Lifecycle.GovernanceEventRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("valid active to revoked is transition active but policy use blocked", func(t *testing.T) {
		model := point11ValAFoundationWithValidRevokedTransition()
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateActive {
			t.Fatalf("expected active revoked transition, got %#v", model)
		}
		if model.PolicyUseState != Point11ValAPolicyUseStateBlocked {
			t.Fatalf("expected blocked policy use, got %#v", model)
		}
	})

	t.Run("active to revoked without governance event blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleRevoked
		model.Lifecycle.GovernanceEventRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("active to revoked without reason blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleRevoked
		model.Lifecycle.Reason = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("active to revoked without evidence or audit blocks transition", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		model.Lifecycle.ToState = point11ValAPolicyLifecycleRevoked
		model.Lifecycle.ApprovalEvidenceRefs = nil
		model.Lifecycle.AuditID = ""
		model = ComputePoint11ValAFoundation(model)
		if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateBlocked {
			t.Fatalf("expected blocked transition state, got %#v", model)
		}
	})

	t.Run("diagnostics distinguish valid revoked transition from malformed revoked transition", func(t *testing.T) {
		valid := point11ValAFoundationWithValidRevokedTransition()
		if valid.LifecycleEvaluation.Reason != "active_to_revoked_transition_valid" ||
			!point11Val0ContainsTrimmed(valid.LifecycleEvaluation.Diagnostics, "policy_use_blocked_by_revocation") {
			t.Fatalf("expected valid revoked transition diagnostics, got %#v", valid.LifecycleEvaluation)
		}

		malformed := activePoint11ValAFoundation()
		malformed.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		malformed.Lifecycle.ToState = point11ValAPolicyLifecycleRevoked
		malformed.Lifecycle.GovernanceEventRef = ""
		malformed = ComputePoint11ValAFoundation(malformed)
		if malformed.LifecycleEvaluation.Reason != "active_to_revoked_transition_missing_revocation_context" ||
			!point11Val0ContainsTrimmed(malformed.LifecycleEvaluation.Diagnostics, "missing_approver_or_governance_event_or_reason_or_evidence") {
			t.Fatalf("expected malformed revoked transition diagnostics, got %#v", malformed.LifecycleEvaluation)
		}
	})

	t.Run("diagnostics distinguish valid superseded transition from malformed superseded transition", func(t *testing.T) {
		valid := point11ValAFoundationWithValidSupersededTransition()
		if valid.LifecycleEvaluation.Reason != "active_to_superseded_transition_valid" ||
			!point11Val0ContainsTrimmed(valid.LifecycleEvaluation.Diagnostics, "policy_use_historical_only_due_to_supersession") {
			t.Fatalf("expected valid superseded transition diagnostics, got %#v", valid.LifecycleEvaluation)
		}

		malformed := activePoint11ValAFoundation()
		malformed.Lifecycle.FromState = point11ValAPolicyLifecycleActive
		malformed.Lifecycle.ToState = point11ValAPolicyLifecycleSuperseded
		malformed.Lifecycle.SuccessorPolicyRef = ""
		malformed = ComputePoint11ValAFoundation(malformed)
		if malformed.LifecycleEvaluation.Reason != "active_to_superseded_transition_missing_supersession_context" ||
			!point11Val0ContainsTrimmed(malformed.LifecycleEvaluation.Diagnostics, "missing_successor_or_compatibility_review_or_governance_event_or_reason_or_evidence") {
			t.Fatalf("expected malformed superseded transition diagnostics, got %#v", malformed.LifecycleEvaluation)
		}
	})
}

func TestPoint11ValAGraphState(t *testing.T) {
	t.Run("valid supersession graph active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Graph.SuccessorPolicyRef = "policy_successor_2026_05_02"
		model.Graph.SupersessionReason = "policy_successor_compatibility_rollout"
		model.Graph.CompatibilityVersion = "point11_vala_compat_v1"
		model.Graph.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
		model.Graph.SuccessorLifecycleState = point11ValAPolicyLifecycleActive
		model.Graph.LineagePath = []string{model.Graph.SourcePolicyRef, model.Graph.SuccessorPolicyRef}
		model = ComputePoint11ValAFoundation(model)
		if model.GraphState != Point11ValAGraphStateActive {
			t.Fatalf("expected active graph state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValAPolicySupersessionRevocationGraph)
	}{
		{name: "self supersession blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = model.SourcePolicyRef
			model.SupersessionReason = "self_supersession_invalid"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
		}},
		{name: "cycle blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_successor_2026_05_02"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
			model.LineagePath = []string{model.SourcePolicyRef, model.SuccessorPolicyRef, model.SourcePolicyRef}
		}},
		{name: "successor policy unknown blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_unknown"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
		}},
		{name: "successor policy revoked blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_revoked"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
		}},
		{name: "successor policy invalid blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_invalid"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
		}},
		{name: "revoked successor blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_successor_2026_05_02"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
			model.SuccessorLifecycleState = point11ValAPolicyLifecycleRevoked
		}},
		{name: "expired successor blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_successor_2026_05_02"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
			model.SuccessorLifecycleState = point11ValAPolicyLifecycleExpired
		}},
		{name: "missing compatibility version blocks when supersession exists", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_successor_2026_05_02"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = ""
			model.CompatibilityReviewRef = "compatibility_review_point11_vala_successor"
		}},
		{name: "missing compatibility review ref blocks when supersession exists", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.SuccessorPolicyRef = "policy_successor_2026_05_02"
			model.SupersessionReason = "policy_successor_compatibility_rollout"
			model.CompatibilityVersion = "point11_vala_compat_v1"
			model.CompatibilityReviewRef = ""
		}},
		{name: "revocation without revoked by ref blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.RevocationReason = "policy_revocation_without_ref"
		}},
		{name: "revoked by policy unknown blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.RevokedByRef = "policy_unknown"
			model.RevocationReason = "policy_revocation_reason"
		}},
		{name: "revoked by governance event unknown blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.RevokedByRef = "governance_event_unknown"
			model.RevocationReason = "policy_revocation_reason"
		}},
		{name: "revocation without reason blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.RevokedByRef = "governance_event_point11_vala_revocation"
		}},
		{name: "missing evidence refs blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.EvidenceRefs = nil
		}},
		{name: "missing audit id blocks", mutate: func(model *Point11ValAPolicySupersessionRevocationGraph) {
			model.AuditID = ""
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValAFoundation()
			testCase.mutate(&model.Graph)
			model = ComputePoint11ValAFoundation(model)
			if model.GraphState != Point11ValAGraphStateBlocked {
				t.Fatalf("expected blocked graph state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValAAggregateState(t *testing.T) {
	t.Run("aggregate active only when all components active", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		if model.CurrentState != Point11ValAStateActive {
			t.Fatalf("expected active aggregate state, got %#v", model)
		}
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal val a foundation: %v", err)
		}
		if strings.Contains(string(body), point11Val0PassToken()) {
			t.Fatalf("expected no point11 pass token in val a output, got %s", body)
		}
	})

	t.Run("dependency review required yields aggregate review required if no local blockers", func(t *testing.T) {
		model := Point11ValAFoundationModel()
		model.Dependency = reviewRequiredPoint11ValADependencySnapshot()
		model = ComputePoint11ValAFoundation(model)
		if model.DependencyState != Point11ValADependencyStateReviewRequired {
			t.Fatalf("expected dependency review required state, got %#v", model)
		}
		if model.CurrentState != Point11ValAStateReviewRequired {
			t.Fatalf("expected aggregate review required state, got %#v", model)
		}
	})

	t.Run("any local component blocked yields aggregate blocked", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Signature.SignatureRef = "signature_invalid"
		model = ComputePoint11ValAFoundation(model)
		if model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected blocked aggregate state, got %#v", model)
		}
	})

	t.Run("diagnostics include component blocking reason", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.Signature.SignatureRef = "signature_invalid"
		model = ComputePoint11ValAFoundation(model)
		if !point11Val0ContainsTrimmed(model.Diagnostics.BlockingReasons, "policy_signature_blocked") {
			t.Fatalf("expected diagnostics to include signature blocking reason, got %#v", model.Diagnostics)
		}
	})

	t.Run("aggregate authority marker blocks", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.CreatesLegalRegulatoryCertificationClaim = true
		model = ComputePoint11ValAFoundation(model)
		if model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected authority marker to block aggregate state, got %#v", model)
		}
	})

	t.Run("aggregate publication side effects marker blocks", func(t *testing.T) {
		model := activePoint11ValAFoundation()
		model.CreatesPublicationSideEffects = true
		model = ComputePoint11ValAFoundation(model)
		if model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected publication side effect marker to block aggregate state, got %#v", model)
		}
	})

	t.Run("aggregate does not treat revoked policy as active use eligible", func(t *testing.T) {
		model := point11ValAFoundationWithValidRevokedTransition()
		if model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected revoked policy to block aggregate active use, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.Diagnostics.BlockingReasons, "policy_use_not_active") {
			t.Fatalf("expected policy use blocking reason, got %#v", model.Diagnostics)
		}
	})

	t.Run("aggregate does not treat superseded policy as active use eligible", func(t *testing.T) {
		model := point11ValAFoundationWithValidSupersededTransition()
		if model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected superseded policy to block aggregate active use, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.Diagnostics.BlockingReasons, "policy_use_historical_only") {
			t.Fatalf("expected historical only policy use reason, got %#v", model.Diagnostics)
		}
	})

	t.Run("valid supersession transition with malformed graph blocks aggregate", func(t *testing.T) {
		model := point11ValAFoundationWithValidSupersededTransition()
		model.Graph.SuccessorPolicyRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.GraphState != Point11ValAGraphStateBlocked || model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected malformed supersession graph to block aggregate, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.LifecycleEvaluation.Diagnostics, "graph_missing_or_mismatched_successor_policy_ref") {
			t.Fatalf("expected lifecycle diagnostics to capture supersession graph mismatch, got %#v", model.LifecycleEvaluation)
		}
	})

	t.Run("superseded lifecycle requires graph compatibility review ref", func(t *testing.T) {
		model := point11ValAFoundationWithValidSupersededTransition()
		model.Graph.CompatibilityReviewRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.GraphState != Point11ValAGraphStateBlocked || model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected missing supersession graph compatibility review to block aggregate, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.LifecycleEvaluation.Diagnostics, "graph_missing_or_mismatched_compatibility_review_ref") {
			t.Fatalf("expected lifecycle diagnostics to capture supersession compatibility review mismatch, got %#v", model.LifecycleEvaluation)
		}
	})

	t.Run("valid revocation transition with malformed graph blocks aggregate", func(t *testing.T) {
		model := point11ValAFoundationWithValidRevokedTransition()
		model.Graph.RevokedByRef = ""
		model = ComputePoint11ValAFoundation(model)
		if model.GraphState != Point11ValAGraphStateBlocked || model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected malformed revocation graph to block aggregate, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.LifecycleEvaluation.Diagnostics, "graph_missing_or_invalid_revoked_by_ref") {
			t.Fatalf("expected lifecycle diagnostics to capture revocation graph mismatch, got %#v", model.LifecycleEvaluation)
		}
	})

	t.Run("revoked lifecycle requires graph revocation reason", func(t *testing.T) {
		model := point11ValAFoundationWithValidRevokedTransition()
		model.Graph.RevocationReason = ""
		model = ComputePoint11ValAFoundation(model)
		if model.GraphState != Point11ValAGraphStateBlocked || model.CurrentState != Point11ValAStateBlocked {
			t.Fatalf("expected missing revocation graph reason to block aggregate, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.LifecycleEvaluation.Diagnostics, "graph_missing_revocation_reason") {
			t.Fatalf("expected lifecycle diagnostics to capture revocation reason mismatch, got %#v", model.LifecycleEvaluation)
		}
	})
}

func TestPoint11ValASemanticAntiGreenRefs(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11ValAFoundation)
	}{
		{name: "policy unknown blocks even though it has a valid prefix", mutate: func(model *Point11ValAFoundation) {
			model.Registry.PolicyID = "policy_unknown"
		}},
		{name: "policy revoked blocks even though it has a valid prefix", mutate: func(model *Point11ValAFoundation) {
			model.Registry.PolicyID = "policy_revoked"
		}},
		{name: "signature invalid blocks even though it has a valid prefix", mutate: func(model *Point11ValAFoundation) {
			model.Signature.SignatureRef = "signature_invalid"
		}},
		{name: "anchor expired blocks even though it has a valid prefix", mutate: func(model *Point11ValAFoundation) {
			model.Anchor.AnchorRef = "anchor_expired"
		}},
		{name: "governance event placeholder blocks even though it has a valid prefix", mutate: func(model *Point11ValAFoundation) {
			model.Registry.GovernanceEventRef = "governance_event_placeholder"
		}},
		{name: "revoked invalid marker blocks in security critical ref field", mutate: func(model *Point11ValAFoundation) {
			model.Registry.SignatureRef = "revoked/invalid marker"
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValAFoundation()
			testCase.mutate(&model)
			model = ComputePoint11ValAFoundation(model)
			if model.CurrentState != Point11ValAStateBlocked {
				t.Fatalf("expected blocked aggregate state, got %#v", model)
			}
		})
	}
}
