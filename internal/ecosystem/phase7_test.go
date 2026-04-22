package ecosystem

import "testing"

func TestEvaluateFoundationStateRequiresAllCrossCuttingBaselines(t *testing.T) {
	coverage := ContractsCoverage()
	if got := EvaluateFoundationState(coverage); got != FoundationStateActive {
		t.Fatalf("expected active foundation baseline, got %q", got)
	}

	coverage.SignalContracts = 0
	if got := EvaluateFoundationState(coverage); got != FoundationStateIncomplete {
		t.Fatalf("expected incomplete foundation without signal contracts, got %q", got)
	}
}

func TestEntryGateBaselineLeavesCanonicalWorkspaceRuntimeInjected(t *testing.T) {
	entry := EntryGateBaseline()
	if entry.CanonicalWorkspace != "" {
		t.Fatalf("expected semantic baseline to leave canonical workspace empty, got %#v", entry)
	}
}

func TestEvaluatePhase7StateRequiresEntryGateFoundationAndAllPresenceSurfaces(t *testing.T) {
	if got := EvaluatePhase7State(EntryGateStateIncomplete, FoundationStateActive, DeveloperPresenceStateActive, OSSPresenceStateActive, DistributionPresenceStateActive); got != Phase7StateIncomplete {
		t.Fatalf("expected incomplete without entry gate, got %q", got)
	}
	if got := EvaluatePhase7State(EntryGateStateReady, FoundationStateActive, OSSPresenceStateActive, OSSPresenceStateActive, DistributionPresenceStateActive); got != Phase7StateIncomplete {
		t.Fatalf("expected incomplete with invalid developer slot state, got %q", got)
	}
	if got := EvaluatePhase7State(EntryGateStateReady, FoundationStateActive, DeveloperPresenceStateActive, OSSPresenceStatePartial, DistributionPresenceStateActive); got != Phase7StateSubstantial {
		t.Fatalf("expected substantial with partial oss presence, got %q", got)
	}
	if got := EvaluatePhase7State(EntryGateStateReady, FoundationStateActive, DeveloperPresenceStateActive, OSSPresenceStateActive, DistributionPresenceStateActive); got != Phase7StateActive {
		t.Fatalf("expected active with all core surfaces, got %q", got)
	}
}

func TestObservationAndClaimSignalsRemainSeparated(t *testing.T) {
	var observation SignalContract
	var reviewedClaim SignalContract
	for _, item := range SignalContractsForGroup("oss") {
		switch item.SignalID {
		case "oss.registry_provenance_observation":
			observation = item
		case "oss.reviewed_trust_claim":
			reviewedClaim = item
		}
	}
	if observation.SignalID == "" || reviewedClaim.SignalID == "" {
		t.Fatalf("expected both observation and reviewed claim contracts, got %#v %#v", observation, reviewedClaim)
	}
	if observation.Status != StatusReviewRequired || observation.MutationCapability != MutationNever {
		t.Fatalf("expected observation to stay candidate/review-required and non-mutating, got %#v", observation)
	}
	if reviewedClaim.Status != StatusVerifierBacked || reviewedClaim.SignalClass != SignalClassVerifierBacked {
		t.Fatalf("expected reviewed claim to be verifier-backed, got %#v", reviewedClaim)
	}
}

func TestFoundationStateFailsClosedWhenRequiredCoreSurfaceCoverageIsMissing(t *testing.T) {
	coverage := buildSurfaceCoverageMap()
	current := coverage["distribution.partner_export"]
	current.hasDataBoundary = false
	coverage["distribution.partner_export"] = current
	if got := evaluateFoundationStateForCoverageMap(coverage); got != FoundationStateIncomplete {
		t.Fatalf("expected incomplete foundation when required export boundary is missing, got %q", got)
	}
}

func TestOSSPresenceIgnoresDeferredExpandedRemediationSurface(t *testing.T) {
	if !isDeferredExpandedSurface("oss.remediation_pr") {
		t.Fatal("expected remediation PR surface to stay deferred from core activation")
	}
	for _, surfaceID := range phase7CoreSurfacesByGroup["oss"] {
		if surfaceID == "oss.remediation_pr" {
			t.Fatalf("expected remediation PR surface to stay out of core OSS surfaces, got %#v", phase7CoreSurfacesByGroup["oss"])
		}
	}
	if got := EvaluateOSSPresenceState(); got != OSSPresenceStateActive {
		t.Fatalf("expected active OSS presence from observation and claim pipelines only, got %q", got)
	}
}

func TestContractsCoverageCountsOnlyCorePassSurfaces(t *testing.T) {
	coverage := ContractsCoverage()
	if coverage.SignalContracts != 9 {
		t.Fatalf("expected 9 core signal contracts without deferred remediation PR, got %#v", coverage)
	}
	if coverage.AuthoritySurfaces != 8 {
		t.Fatalf("expected 8 core authority surfaces without deferred remediation PR, got %#v", coverage)
	}
	if coverage.AbuseControls != 5 {
		t.Fatalf("expected 5 core abuse controls without deferred remediation PR, got %#v", coverage)
	}
}

func TestPartnerExportBoundaryNeverBecomesPublic(t *testing.T) {
	boundaries := DataBoundariesForGroup("distribution")
	for _, boundary := range boundaries {
		if boundary.SurfaceID != "distribution.partner_export" {
			continue
		}
		if len(boundary.PublicExportable) != 0 {
			t.Fatalf("expected partner export to avoid public exportable fields, got %#v", boundary)
		}
		if len(boundary.PartnerVisible) == 0 || !boundary.RedactedByDefault {
			t.Fatalf("expected partner export to stay scoped and redacted, got %#v", boundary)
		}
		return
	}
	t.Fatal("expected partner export boundary contract")
}
