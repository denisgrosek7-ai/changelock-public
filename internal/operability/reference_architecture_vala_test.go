package operability

import "testing"

func activeReferenceArchitectureValAPrereqs() (string, string, string, string) {
	return IntelligenceCalibrationPoint5StatePass,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitecturePoint6StateNotComplete
}

func TestReferenceArchitectureValARegistryContainsAllSixFamilies(t *testing.T) {
	registry := ReferenceArchitectureValAFamilyRegistry()
	if len(registry.Profiles) != 6 {
		t.Fatalf("expected 6 Val A family profiles, got %d", len(registry.Profiles))
	}
	if !containsExactTrimmedStringSet(registry.SupportedFamilies, referenceArchitectureVal0Families()...) {
		t.Fatalf("expected supported families to match Val 0 taxonomy, got %#v", registry.SupportedFamilies)
	}
}

func TestReferenceArchitectureValAUnknownFamilyLookupFailsClosed(t *testing.T) {
	if _, ok := LookupReferenceArchitectureValAFamilyProfile("unknown_family"); ok {
		t.Fatalf("expected unknown family lookup to fail closed")
	}
}

func TestReferenceArchitectureValAFamilyProfilesValidateRequiredFields(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	for _, profile := range ReferenceArchitectureValAFamilyRegistry().Profiles {
		if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got != ReferenceArchitectureValAFamilyProfileStateActive {
			t.Fatalf("expected active family profile state for %s, got %q", profile.Family, got)
		}
		if profile.LifecycleState == "" || profile.CompatibilityState == "" || len(profile.RequiredCapabilities) == 0 || len(profile.RequiredEvidenceTypes) == 0 || profile.SupportBoundaryRef == "" || profile.ProjectionDisclaimer == "" {
			t.Fatalf("expected required fields for %s to be populated, got %#v", profile.Family, profile)
		}
		if len(profile.DegradedConditions) == 0 || len(profile.UnsupportedConditions) == 0 {
			t.Fatalf("expected degraded and unsupported conditions for %s, got %#v", profile.Family, profile)
		}
	}
}

func TestReferenceArchitectureValARegistryStateIsActiveWithValidProfiles(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	if got := EvaluateReferenceArchitectureValAFamilyRegistryState(point5State, val0CurrentState, val0State, point6State, ReferenceArchitectureValAFamilyRegistry()); got != ReferenceArchitectureValAFamilyRegistryStateActive {
		t.Fatalf("expected active family registry state, got %q", got)
	}
}

func TestReferenceArchitectureValARegistryRejectsDuplicateFamilyEntries(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	registry := ReferenceArchitectureValAFamilyRegistry()
	registry.Profiles[1].Family = registry.Profiles[0].Family
	if got := EvaluateReferenceArchitectureValAFamilyRegistryState(point5State, val0CurrentState, val0State, point6State, registry); got == ReferenceArchitectureValAFamilyRegistryStateActive {
		t.Fatalf("expected non-active registry state for duplicate family entry, got %q", got)
	}
}

func TestReferenceArchitectureValARegistryRejectsDuplicateBlueprintIDs(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	registry := ReferenceArchitectureValAFamilyRegistry()
	registry.Profiles[1].BlueprintID = registry.Profiles[0].BlueprintID
	if got := EvaluateReferenceArchitectureValAFamilyRegistryState(point5State, val0CurrentState, val0State, point6State, registry); got == ReferenceArchitectureValAFamilyRegistryStateActive {
		t.Fatalf("expected non-active registry state for duplicate blueprint_id, got %q", got)
	}
}

func TestReferenceArchitectureValAEnterpriseDefaultRequiresBalancedProductionCapabilities(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	profile, ok := LookupReferenceArchitectureValAFamilyProfile(ReferenceArchitectureFamilyEnterpriseDefault)
	if !ok {
		t.Fatalf("expected enterprise_default profile to exist")
	}
	if !containsAllTrimmedStrings(profile.RequiredCapabilities,
		ReferenceArchitectureCapabilitySigning,
		ReferenceArchitectureCapabilityAuditWriter,
		ReferenceArchitectureCapabilityEvidenceStorage,
		ReferenceArchitectureCapabilityPolicyDist,
		ReferenceArchitectureCapabilityRecovery,
	) {
		t.Fatalf("expected enterprise_default required capabilities to be present, got %#v", profile.RequiredCapabilities)
	}
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got != ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected active enterprise_default profile, got %q", got)
	}
}

func TestReferenceArchitectureValAHighAssuranceIsStricterThanEnterpriseDefault(t *testing.T) {
	enterprise, ok := LookupReferenceArchitectureValAFamilyProfile(ReferenceArchitectureFamilyEnterpriseDefault)
	if !ok {
		t.Fatalf("expected enterprise_default profile to exist")
	}
	highAssurance, ok := LookupReferenceArchitectureValAFamilyProfile(ReferenceArchitectureFamilyHighAssurance)
	if !ok {
		t.Fatalf("expected high_assurance profile to exist")
	}
	if !highAssurance.StrongerTrustAnchorMode || !highAssurance.StricterAuditCustody || !highAssurance.StrongerRecoveryExpectations {
		t.Fatalf("expected high_assurance stricter posture flags, got %#v", highAssurance)
	}
	if enterprise.TargetEnvironment.TrustAnchorMode == highAssurance.TargetEnvironment.TrustAnchorMode {
		t.Fatalf("expected high_assurance trust anchor mode to be stricter than enterprise_default")
	}
	if highAssurance.RequiresAllWorkloadsInEnclaves {
		t.Fatalf("did not expect high_assurance to require all workloads in enclaves")
	}
}

func TestReferenceArchitectureValARegulatedPrivacyFirstRequiresResidencyRedactionExportAndCustody(t *testing.T) {
	profile, ok := LookupReferenceArchitectureValAFamilyProfile(ReferenceArchitectureFamilyRegulatedPrivacyFirst)
	if !ok {
		t.Fatalf("expected regulated_privacy_first profile to exist")
	}
	if !profile.DataResidencyDisciplineRequired || !profile.RedactionExportBoundaryRequired || !profile.EvidenceCustodyRequired {
		t.Fatalf("expected regulated_privacy_first privacy-first boundary flags, got %#v", profile)
	}
}

func TestReferenceArchitectureValASovereignAirGappedRequiresOfflineLocalBoundaries(t *testing.T) {
	profile, ok := LookupReferenceArchitectureValAFamilyProfile(ReferenceArchitectureFamilySovereignAirGapped)
	if !ok {
		t.Fatalf("expected sovereign_air_gapped profile to exist")
	}
	if profile.TargetEnvironment.ConnectivityMode != ReferenceArchitectureConnectivityAirGapped || !profile.OfflineTransferBoundaryRequired || !profile.LocalTrustAnchorAssumptionRequired || !profile.LocalOperatorControlRequired {
		t.Fatalf("expected sovereign_air_gapped offline and local boundary requirements, got %#v", profile)
	}
	if profile.LiveExternalDependencyAllowedOffline {
		t.Fatalf("did not expect sovereign_air_gapped profile to allow live external dependency while offline")
	}
}

func TestReferenceArchitectureValAPerformanceSensitiveRequiresPerformanceAssumptions(t *testing.T) {
	profile, ok := LookupReferenceArchitectureValAFamilyProfile(ReferenceArchitectureFamilyPerformanceSensitive)
	if !ok {
		t.Fatalf("expected performance_sensitive profile to exist")
	}
	if !profile.PerformanceEnvelopeRequired || !profile.AuditWritePathDisciplineRequired || !profile.ControlPlaneCapacityRequired {
		t.Fatalf("expected performance_sensitive profile to require performance and control-plane assumptions, got %#v", profile)
	}
}

func TestReferenceArchitectureValAPartnerMSPSuitableRequiresAccessAndAuthorityBoundaries(t *testing.T) {
	profile, ok := LookupReferenceArchitectureValAFamilyProfile(ReferenceArchitectureFamilyPartnerMSPSuitable)
	if !ok {
		t.Fatalf("expected partner_msp_suitable profile to exist")
	}
	if !profile.CustomerAuthorityBoundaryRequired || !profile.NoPartnerShadowTruthRule || !profile.PartnerVisibilityRestrictionRequired {
		t.Fatalf("expected partner_msp_suitable authority and visibility boundaries, got %#v", profile)
	}
	if profile.PartnerCanonicalTruthOverrideAllowed {
		t.Fatalf("did not expect partner_msp_suitable to allow partner canonical truth override")
	}
}

func TestReferenceArchitectureValAMissingVal0PrerequisiteBlocksActiveState(t *testing.T) {
	registry := ReferenceArchitectureValAFamilyRegistry()
	if got := EvaluateReferenceArchitectureValAFamilyRegistryState(IntelligenceCalibrationPoint5StatePass, ReferenceArchitectureVal0StateIncomplete, ReferenceArchitectureVal0StateActive, ReferenceArchitecturePoint6StateNotComplete, registry); got != ReferenceArchitectureValAFamilyRegistryStateBlocked {
		t.Fatalf("expected blocked registry state without active Val 0 prerequisite, got %q", got)
	}
}

func TestReferenceArchitectureValAUnknownFamilyEnumFailsClosed(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	profile := referenceArchitectureValAEnterpriseDefaultProfile()
	profile.Family = "enterprise-defualt"
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state for unknown family enum, got %q", got)
	}
}

func TestReferenceArchitectureValATypoLifecycleAndCompatibilityFailClosed(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	profile := referenceArchitectureValAEnterpriseDefaultProfile()
	profile.LifecycleState = "acitve"
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state for typo lifecycle, got %q", got)
	}

	profile = referenceArchitectureValAEnterpriseDefaultProfile()
	profile.CompatibilityState = "compatibel"
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state for typo compatibility, got %q", got)
	}
}

func TestReferenceArchitectureValAMissingRequiredCapabilityBlocksActiveState(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	profile := referenceArchitectureValAEnterpriseDefaultProfile()
	profile.RequiredCapabilities = []string{
		ReferenceArchitectureCapabilitySigning,
		ReferenceArchitectureCapabilityAuditWriter,
	}
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state with missing required capabilities, got %q", got)
	}
}

func TestReferenceArchitectureValAMissingDegradedUnsupportedOrEvidenceBlocksActiveState(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	profile := referenceArchitectureValAEnterpriseDefaultProfile()
	profile.DegradedConditions = nil
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state without degraded conditions, got %q", got)
	}

	profile = referenceArchitectureValAEnterpriseDefaultProfile()
	profile.UnsupportedConditions = nil
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state without unsupported conditions, got %q", got)
	}

	profile = referenceArchitectureValAEnterpriseDefaultProfile()
	profile.RequiredEvidenceTypes = nil
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state without evidence requirements, got %q", got)
	}
}

func TestReferenceArchitectureValAMissingProjectionDisclaimerBlocksActiveState(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	profile := referenceArchitectureValAEnterpriseDefaultProfile()
	profile.ProjectionDisclaimer = ""
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state without projection disclaimer, got %q", got)
	}
}

func TestReferenceArchitectureValACertifiedGuaranteedAbsoluteLanguageBlocksActiveState(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	profile := referenceArchitectureValAEnterpriseDefaultProfile()
	profile.CertifiedLanguagePresent = true
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state with certified language, got %q", got)
	}

	profile = referenceArchitectureValAEnterpriseDefaultProfile()
	profile.GuaranteedSecurityClaimPresent = true
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state with guaranteed language, got %q", got)
	}

	profile = referenceArchitectureValAEnterpriseDefaultProfile()
	profile.AbsoluteSecurityClaimPresent = true
	if got := EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile); got == ReferenceArchitectureValAFamilyProfileStateActive {
		t.Fatalf("expected non-active family profile state with absolute security language, got %q", got)
	}
}

func TestReferenceArchitectureValANoOverclaimAndPoint6NotCompleteDiscipline(t *testing.T) {
	point5State, val0CurrentState, val0State, point6State := activeReferenceArchitectureValAPrereqs()
	registryState := EvaluateReferenceArchitectureValAFamilyRegistryState(point5State, val0CurrentState, val0State, point6State, ReferenceArchitectureValAFamilyRegistry())
	valAState := EvaluateReferenceArchitectureValAState(point5State, val0CurrentState, val0State, point6State, registryState)
	if valAState != ReferenceArchitectureValAStateActive {
		t.Fatalf("expected active Val A state with valid registry, got %q", valAState)
	}
	if got := EvaluateReferenceArchitectureValAProofsState(
		valAState,
		ReferenceArchitecturePoint6StatePass,
		referenceArchitectureVal0Families(),
		[]string{
			"/v1/reference-architecture/val0/proofs",
			"/v1/reference-architecture/vala/family-registry",
			"/v1/reference-architecture/vala/family-profiles",
			"/v1/reference-architecture/vala/proofs",
		},
		[]string{
			"point5_integrated_closure",
			"point6_val0_proofs",
			"registry",
			"p1",
			"p2",
			"p3",
			"p4",
			"p5",
			"p6",
		},
		[]string{"Val A keeps Točka 6 not complete."},
		"projection_only not_canonical_truth bounded_reference_architecture_family_profiles",
	); got == ReferenceArchitectureValAStateActive {
		t.Fatalf("expected non-active proofs state when point_6_pass is claimed in Val A, got %q", got)
	}
}
