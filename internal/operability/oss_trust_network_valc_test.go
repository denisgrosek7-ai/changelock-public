package operability

import "testing"

func copyOSSTrustNetworkValCEvidence() []ReferenceArchitectureEvidenceReference {
	evidence := ossTrustNetworkValCEvidence()
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

func activeOSSTrustNetworkValCModel() OSSTrustNetworkValCCore {
	model := OSSTrustNetworkValCCoreModel()
	return ComputeOSSTrustNetworkValCCore(model)
}

func activeOSSTrustNetworkValCLimitations() []string {
	return []string{
		"Val C defines bounded remediation and ecosystem visibility only and does not implement final OSTN gates, integrated closure, or Točka 10.",
		"Visibility, remediation suggestions, proposal descriptors, and local overrides remain advisory and bounded by explicit no-hidden-mutation and no-overclaim discipline.",
	}
}

func TestOSSTrustNetworkValCHappyPathAndPoint9NotComplete(t *testing.T) {
	model := activeOSSTrustNetworkValCModel()
	if model.CurrentState != OSSTrustNetworkValCStateActive {
		t.Fatalf("expected active OSTN Val C state, got %#v", model)
	}
	if model.Point9State != OSSTrustNetworkPoint9StateNotComplete {
		t.Fatalf("expected point 9 to remain not complete, got %#v", model)
	}
	if got := EvaluateOSSTrustNetworkValCProofsState(model, activeOSSTrustNetworkValCLimitations()); got != OSSTrustNetworkValCStateActive {
		t.Fatalf("expected active Val C proofs state, got %q", got)
	}
	if ossTrustNetworkValCContainsForbiddenClaim("not canonical truth") {
		t.Fatalf("expected exact safe negative wording to remain allowed")
	}
}

func TestOSSTrustNetworkValCDependencyAndVisibilityBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValCCore)
		wantState string
	}{
		{
			name: "missing valb dependency blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.Dependency.ValBCurrentState = OSSTrustNetworkValBStateBlocked
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "route presence alone cannot satisfy valb dependency",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.Dependency = OSSTrustNetworkValCDependencySnapshot{
					ValBCurrentState:                OSSTrustNetworkValBStateActive,
					ValBPoint9State:                 OSSTrustNetworkPoint9StateNotComplete,
					ValBDependencyState:             OSSTrustNetworkValBDependencyStateActive,
					ValBCandidateSignalIntakeState:  OSSTrustNetworkValBCandidateSignalIntakeStateActive,
					ValBReviewWorkflowState:         OSSTrustNetworkValBReviewWorkflowStatePartial,
					ValBSharedVEXTriageState:        OSSTrustNetworkValBSharedVEXTriageStateActive,
					ValBSourceWeightingState:        OSSTrustNetworkValBSourceWeightingStateActive,
					ValBLocalApplicabilityState:     OSSTrustNetworkValBLocalApplicabilityStateActive,
					ValBPropagationExchangeState:    OSSTrustNetworkValBPropagationExchangeStateActive,
					ValBSupersessionRevocationState: OSSTrustNetworkValBSupersessionRevocationStateActive,
					ValBReviewerAuditabilityState:   OSSTrustNetworkValBReviewerAuditabilityStateActive,
					ValBNoOverclaimState:            OSSTrustNetworkValBNoOverclaimStateActive,
					ValBProofSurfaceRefs:            append([]string{}, OSSTrustNetworkValBProofSurfaceRefs()...),
					ValBEvidenceRefs:                append([]string{}, OSSTrustNetworkValBProofEvidenceRefs()...),
					ValBProjectionDisclaimer:        ossTrustNetworkValBProjectionDisclaimer(),
				}
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "valb malformed evidence identity blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.Dependency.ValBEvidenceRefs = append(model.Dependency.ValBEvidenceRefs, "evidence:ostn-valb-extra-001")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "valb non active subgate blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.Dependency.ValBNoOverclaimState = OSSTrustNetworkValBNoOverclaimStateBlocked
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "non active valb current state blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.Dependency.ValBCurrentState = OSSTrustNetworkValBStatePartial
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name:      "oss trust visibility visible with evidence freshness review local applicability caveats passes",
			mutate:    func(model *OSSTrustNetworkValCCore) {},
			wantState: OSSTrustNetworkValCStateActive,
		},
		{
			name: "oss trust visibility limited is partial",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.VisibilityState = OSSTrustNetworkValCVisibilityLimited
			},
			wantState: OSSTrustNetworkValCStatePartial,
		},
		{
			name: "oss trust visibility hidden blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.VisibilityState = OSSTrustNetworkValCVisibilityHidden
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "oss trust visibility unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.VisibilityState = OSSTrustNetworkValCVisibilityUnsupported
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "oss trust visibility stale blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.VisibilityState = OSSTrustNetworkValCVisibilityStale
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "oss trust visibility unknown blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.VisibilityState = OSSTrustNetworkValCVisibilityUnknown
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "oss trust visibility malformed blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.VisibilityState = "visible-ish"
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "oss trust visibility malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name: "visibility cannot imply package safety certification approval or global truth",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.PackageSafetyClaim = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValCModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValCCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValCPackageExportAndRemediationBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValCCore)
		wantState string
	}{
		{
			name:      "package status reviewed signal available with reviewed intelligence evidence local applicability passes",
			mutate:    func(model *OSSTrustNetworkValCCore) {},
			wantState: OSSTrustNetworkValCStateActive,
		},
		{
			name: "package status candidate signal available is partial and not reviewed",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusCandidateSignalAvailable
				model.PackageTrustStatus.DisplayedAsReviewed = false
				model.EcosystemConsistency.PackageStatusClass = OSSTrustNetworkValCPackageStatusCandidateSignalAvailable
				model.EcosystemConsistency.DisplayedAsReviewed = false
				model.EcosystemConsistency.ReviewedExchangePresentedActive = false
			},
			wantState: OSSTrustNetworkValCStatePartial,
		},
		{
			name: "package status local review needed is partial",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusLocalReviewNeeded
				model.PackageTrustStatus.DisplayedAsReviewed = false
				model.EcosystemConsistency.PackageStatusClass = OSSTrustNetworkValCPackageStatusLocalReviewNeeded
				model.EcosystemConsistency.DisplayedAsReviewed = false
				model.EcosystemConsistency.ReviewedExchangePresentedActive = false
			},
			wantState: OSSTrustNetworkValCStatePartial,
		},
		{
			name: "package status local review needed cannot display as reviewed",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusLocalReviewNeeded
				model.PackageTrustStatus.DisplayedAsReviewed = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "package status superseded without replacement blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusSupersededSignal
				model.PackageTrustStatus.ReplacementRef = ""
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "package status revoked blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusRevokedSignal
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "package status unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusUnsupportedSignal
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "package status unknown blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusUnknownSignal
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "package trust malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name: "package status cannot become score badge or generic safety",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.BadgeScoreClaim = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name:      "export class exact set passes",
			mutate:    func(model *OSSTrustNetworkValCCore) {},
			wantState: OSSTrustNetworkValCStateActive,
		},
		{
			name: "unknown export class blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.ExportBoundary.ExportClass = "partner-public-ish"
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "export boundary malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.ExportBoundary.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name: "public summary cannot expose canonical internals or imply certification",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.ExportBoundary.ExportClass = OSSTrustNetworkValCExportClassPublicSummaryView
				model.ExportBoundary.CanonicalInternalExposure = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "redaction cannot convert candidate rejected revoked or unknown into reviewed active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.ExportBoundary.CandidatePromotedToReviewed = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name:      "remediation suggestion with evidence compatibility risk local applicability caveats passes",
			mutate:    func(model *OSSTrustNetworkValCCore) {},
			wantState: OSSTrustNetworkValCStateActive,
		},
		{
			name: "remediation suggestion missing compatibility note blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.CompatibilityNote = ""
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation suggestion missing risk note blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.RiskNote = ""
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation suggestion missing local applicability note blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.LocalApplicabilityNote = ""
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "no action suggestion without rationale blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.SuggestionClass = OSSTrustNetworkValCSuggestionClassNoAction
				model.RemediationSuggestion.Rationale = ""
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "unsupported suggestion is non active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.SuggestionClass = OSSTrustNetworkValCSuggestionClassUnsupported
			},
			wantState: OSSTrustNetworkValCStatePartial,
		},
		{
			name: "remediation suggestion cannot mutate dependencies or override policy",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.DependencyMutationAttempt = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation suggestion malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name: "remediation suggestion package identity mismatch blocks val c active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.PackageOrProjectIdentity = "github.com/example/other-project"
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation suggestion affected release mismatch blocks val c active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.AffectedReleaseOrVersion = "refs/tags/v9.9.9"
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation target release may differ without blocking",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSuggestion.TargetReleaseOrVersion = "refs/tags/v9.9.9"
			},
			wantState: OSSTrustNetworkValCStateActive,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValCModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValCCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValCProposalOverrideSafetyConsistencyAndOverclaimBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*OSSTrustNetworkValCCore)
		wantState string
	}{
		{
			name:      "pr proposal proposal ready with reviewer requirement no automerge no hidden mutation passes",
			mutate:    func(model *OSSTrustNetworkValCCore) {},
			wantState: OSSTrustNetworkValCStateActive,
		},
		{
			name: "pr proposal needs review is partial",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PRProposal.ProposalState = OSSTrustNetworkValCProposalStateNeedsReview
			},
			wantState: OSSTrustNetworkValCStatePartial,
		},
		{
			name: "pr proposal unsupported blocked unknown malformed blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PRProposal.ProposalState = OSSTrustNetworkValCProposalStateUnsupported
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "pr proposal with auto merge or hidden mutation blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PRProposal.AutoMerge = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "pr proposal with branch write network action or dependency mutation blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PRProposal.BranchWrite = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "pr proposal malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PRProposal.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name:      "local override no override passes",
			mutate:    func(model *OSSTrustNetworkValCCore) {},
			wantState: OSSTrustNetworkValCStateActive,
		},
		{
			name: "local override override present with evidence rationale scope owner caveats passes",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.OverrideState = OSSTrustNetworkValCOverrideStateOverridePresent
			},
			wantState: OSSTrustNetworkValCStateActive,
		},
		{
			name: "local override override requires review is partial",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.OverrideState = OSSTrustNetworkValCOverrideStateOverrideRequiresReview
				model.EcosystemConsistency.DisplayedAsApplicable = false
			},
			wantState: OSSTrustNetworkValCStatePartial,
		},
		{
			name: "local override override rejected blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.OverrideState = OSSTrustNetworkValCOverrideStateOverrideRejected
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "local override unsupported blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.OverrideState = OSSTrustNetworkValCOverrideStateUnsupported
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "local override unknown blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.OverrideState = OSSTrustNetworkValCOverrideStateUnknown
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "local override cannot rewrite canonical evidence or silently suppress network intelligence",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.SilentlySuppressNetworkIntelligence = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "local override malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name: "local override requires review blocks applicable display",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.LocalOverride.OverrideState = OSSTrustNetworkValCOverrideStateOverrideRequiresReview
				model.EcosystemConsistency.DisplayedAsApplicable = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation safety missing test validation note blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSafety.TestValidationNote = ""
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation safety missing rollback note blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSafety.RollbackNote = ""
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "high risk suggestion without reviewer requirement blocks",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSafety.RiskClass = OSSTrustNetworkValCRiskClassHigh
				model.RemediationSafety.ReviewerRequired = false
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "remediation safety malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.RemediationSafety.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name: "ecosystem consistency rejects reviewed display when valb review is not reviewed accepted",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.ReviewState = OSSTrustNetworkValBReviewStateCandidate
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "trust visibility package identity mismatch blocks val c active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.PackageOrProjectIdentity = "github.com/example/other-project"
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "trust visibility release version mismatch blocks val c active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.TrustVisibility.ReleaseOrVersionRef = "refs/tags/v9.9.9"
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "package trust status class mismatch against ecosystem consistency blocks val c active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.StatusClass = OSSTrustNetworkValCPackageStatusCandidateSignalAvailable
				model.PackageTrustStatus.DisplayedAsReviewed = false
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "package trust status displayed as reviewed mismatch blocks val c active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.PackageTrustStatus.DisplayedAsReviewed = false
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects candidate displayed as reviewed",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.PackageStatusClass = OSSTrustNetworkValCPackageStatusCandidateSignalAvailable
				model.EcosystemConsistency.DisplayedAsReviewed = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects revoked superseded displayed as active reviewed exchange",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.PropagationState = OSSTrustNetworkValBPropagationStateRevoked
				model.EcosystemConsistency.ReviewedExchangePresentedActive = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects revoked package status displayed as active reviewed exchange",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.PackageStatusClass = OSSTrustNetworkValCPackageStatusRevokedSignal
				model.EcosystemConsistency.ReviewedExchangePresentedActive = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects superseded package status displayed as active reviewed exchange",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.PackageStatusClass = OSSTrustNetworkValCPackageStatusSupersededSignal
				model.EcosystemConsistency.ReviewedExchangePresentedActive = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects unsupported package status with reviewed exchange",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.PackageStatusClass = OSSTrustNetworkValCPackageStatusUnsupportedSignal
				model.EcosystemConsistency.ReviewedExchangePresentedActive = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects unknown package status with reviewed exchange",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.PackageStatusClass = OSSTrustNetworkValCPackageStatusUnknownSignal
				model.EcosystemConsistency.ReviewedExchangePresentedActive = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects local applicability unknown displayed as applicable",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.LocalApplicabilityState = OSSTrustNetworkValBLocalApplicabilityStatusUnknown
				model.EcosystemConsistency.DisplayedAsApplicable = true
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency rejects incompatible package release identity",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.SuggestionPackageOrProjectIdentity = "github.com/example/other-project"
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "ecosystem consistency malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
		{
			name: "stale evidence freshness blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.FreshnessState = IntelligenceCalibrationFreshnessStale
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "unknown evidence freshness blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.EcosystemConsistency.FreshnessState = IntelligenceCalibrationFreshnessUnknown
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "no overclaim denylist blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "point_9_pass")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "officially safe package blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "officially safe package")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "compliance guaranteed blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "compliance guaranteed")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "deployment approved blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "deployment approved")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "legal certification blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "legal certification")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "universal trust score blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "universal trust score")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "reviewed means safe blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "reviewed means safe")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "community truth blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "community truth")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "network truth blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ObservedClaims = append(model.NoOverclaim.ObservedClaims, "network truth")
			},
			wantState: OSSTrustNetworkValCStateBlocked,
		},
		{
			name: "no overclaim malformed projection disclaimer blocks active",
			mutate: func(model *OSSTrustNetworkValCCore) {
				model.NoOverclaim.ProjectionDisclaimer = "projection_only but not val c"
			},
			wantState: OSSTrustNetworkValCStateUnknown,
		},
	}

	for _, tc := range testCases {
		model := activeOSSTrustNetworkValCModel()
		tc.mutate(&model)
		model = ComputeOSSTrustNetworkValCCore(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestOSSTrustNetworkValCProofEvidenceQualityValidation(t *testing.T) {
	testCases := []struct {
		name     string
		evidence func() []ReferenceArchitectureEvidenceReference
		refs     func() []string
		want     bool
	}{
		{
			name: "exact evidence model passes",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				return copyOSSTrustNetworkValCEvidence()
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValCProofEvidenceRefs()...)
			},
			want: true,
		},
		{
			name: "mismatched evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValCEvidence()
				evidence[0].EvidenceID, evidence[1].EvidenceID = evidence[1].EvidenceID, evidence[0].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValCProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "duplicate evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValCEvidence()
				evidence[0].EvidenceID = evidence[1].EvidenceID
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValCProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "unknown evidence id fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValCEvidence()
				evidence[0].EvidenceID = "evidence:ostn-valc-unknown-001"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValCProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong scope fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValCEvidence()
				evidence[0].Scope = "wrong_scope"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValCProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong evidence type fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValCEvidence()
				evidence[0].EvidenceType = "wrong_type"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValCProofEvidenceRefs()...)
			},
			want: false,
		},
		{
			name: "wrong source fails",
			evidence: func() []ReferenceArchitectureEvidenceReference {
				evidence := copyOSSTrustNetworkValCEvidence()
				evidence[0].Source = "wrong/source"
				return evidence
			},
			refs: func() []string {
				return append([]string{}, OSSTrustNetworkValCProofEvidenceRefs()...)
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := OSSTrustNetworkValCProofEvidenceQualityValid(tc.evidence(), tc.refs()); got != tc.want {
			t.Fatalf("%s: expected %t, got %t", tc.name, tc.want, got)
		}
	}
}
