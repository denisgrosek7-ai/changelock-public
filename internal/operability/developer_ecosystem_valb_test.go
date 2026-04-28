package operability

import "testing"

func activeDeveloperEcosystemValBModel() DeveloperEcosystemValBIntegration {
	model := DeveloperEcosystemValBIntegrationModel()
	model.ValECompatibility.ValECurrentState = VerifierEcosystemValEStatePass
	model.ValECompatibility.Point7State = VerifierEcosystemPoint7StatePass
	model.ValECompatibility.PassRuleState = VerifierEcosystemValEPassRuleStateActive
	model.ValECompatibility.NoOverclaimState = VerifierEcosystemValENoOverclaimStateActive
	model.ValECompatibility.ProofSurfaceState = VerifierEcosystemValEProofSurfaceStateActive
	model.ValECompatibility.EvidenceQualityState = VerifierEcosystemValEEvidenceQualityStateActive
	model.ValECompatibility.Point7PassAllowed = true
	model.ValECompatibility.Point7PassReason = VerifierEcosystemValEPoint7PassReasonAllowed
	model.ValECompatibility.SurfaceRefs = VerifierEcosystemValEProofSurfaceRefs()
	model.ValECompatibility.EvidenceRefs = VerifierEcosystemValEProofEvidenceRefs()
	model.ValECompatibility.ProjectionDisclaimer = verifierEcosystemValEProjectionDisclaimer()
	model.Dependency = DeveloperEcosystemValBDependencySnapshot{
		ValACurrentState:         DeveloperEcosystemValAStateActive,
		ValAPoint8State:          DeveloperEcosystemPoint8StateNotComplete,
		ValADependencyState:      DeveloperEcosystemValADependencyStateActive,
		IDEBaselineState:         DeveloperEcosystemValAIDEBaselineStateActive,
		TrustFeedbackState:       DeveloperEcosystemValATrustFeedbackStateActive,
		CAVIVEXContextState:      DeveloperEcosystemValACAVIVEXStateActive,
		LocalAdvisoryState:       DeveloperEcosystemValALocalAdvisoryStateActive,
		ValidationHarnessState:   DeveloperEcosystemValAValidationHarnessStateActive,
		MockVerificationState:    DeveloperEcosystemValAMockVerificationStateActive,
		InspectExplainState:      DeveloperEcosystemValAInspectExplainStateActive,
		DegradedModeState:        DeveloperEcosystemValADegradedModeStateActive,
		NoOverclaimState:         DeveloperEcosystemValANoOverclaimStateActive,
		ValAProofSurfaceRefs:     DeveloperEcosystemValAProofSurfaceRefs(),
		ValAEvidenceRefs:         DeveloperEcosystemValAProofEvidenceRefs(),
		ValAProjectionDisclaimer: developerEcosystemValAProjectionDisclaimer(),
	}
	return ComputeDeveloperEcosystemValBIntegration(model)
}

func activeDeveloperEcosystemValBLimitations() []string {
	return []string{
		"Val B implements repo and SDK integration contracts only and does not implement a production SDK runtime, repo parser/runtime, plugin runtime, marketplace publishing, or Točka 9 work.",
		"Točka 8 remains not complete because later developer ecosystem waves are still required before any integrated closure can exist.",
		"Repo config, policy preview, local-to-CI continuity, SDK/API, and examples/templates remain advisory only and cannot approve deployment, certify trust, or create canonical evidence.",
	}
}

func TestDeveloperEcosystemValBHappyPathAndPoint8NotComplete(t *testing.T) {
	model := activeDeveloperEcosystemValBModel()
	if model.CurrentState != DeveloperEcosystemValBStateActive {
		t.Fatalf("expected active developer Val B state, got %#v", model)
	}
	if model.RepoConfigSchema.CompatibilityBehavior != DeveloperEcosystemValBRepoConfigCompatibilityBehaviorBounded {
		t.Fatalf("expected canonical repo schema compatibility behavior, got %#v", model.RepoConfigSchema)
	}
	if model.APIVersioning.VersionIdentity != DeveloperEcosystemValBAPIVersionIdentity ||
		model.APIVersioning.CompatibilityWindow != DeveloperEcosystemValBAPICompatibilityWindow {
		t.Fatalf("expected canonical API version identity/window, got %#v", model.APIVersioning)
	}
	if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
		t.Fatalf("expected point 8 to remain not complete in Val B, got %#v", model)
	}
	if got := EvaluateDeveloperEcosystemValBProofsState(model, activeDeveloperEcosystemValBLimitations()); got != DeveloperEcosystemValBStateActive {
		t.Fatalf("expected active developer Val B proofs state, got %q", got)
	}
}

func TestDeveloperEcosystemValBRepoSchemaCompatibilityBehaviorValidation(t *testing.T) {
	model := activeDeveloperEcosystemValBModel()
	if got := EvaluateDeveloperEcosystemValBRepoConfigSchemaState(model.RepoConfigSchema); got != DeveloperEcosystemValBRepoConfigSchemaStateActive {
		t.Fatalf("expected default repo schema compatibility behavior to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "empty behavior incomplete", value: "", expected: DeveloperEcosystemValBRepoConfigSchemaStateIncomplete},
		{name: "mismatched behavior unknown", value: "compatibility_relaxed", expected: DeveloperEcosystemValBRepoConfigSchemaStateUnknown},
		{name: "permissive behavior blocked", value: "permissive", expected: DeveloperEcosystemValBRepoConfigSchemaStateBlocked},
		{name: "unsupported behavior blocked", value: "unsupported", expected: DeveloperEcosystemValBRepoConfigSchemaStateBlocked},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValBModel()
		mutated.RepoConfigSchema.CompatibilityBehavior = tc.value
		if got := EvaluateDeveloperEcosystemValBRepoConfigSchemaState(mutated.RepoConfigSchema); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}

	mutated := activeDeveloperEcosystemValBModel()
	mutated.RepoConfigSchema.EnterprisePolicyOverrideAttempt = true
	if got := EvaluateDeveloperEcosystemValBRepoConfigSchemaState(mutated.RepoConfigSchema); got != DeveloperEcosystemValBRepoConfigSchemaStateBlocked {
		t.Fatalf("expected enterprise policy override to remain blocked, got %q", got)
	}
}

func TestDeveloperEcosystemValBValECompatibilityAndDependencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValBIntegration)
	}{
		{name: "vale point7 pass reason overclaim blocks", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ValECompatibility.Point7PassReason = "point_7_pass production approved"
		}},
		{name: "vale no overclaim blocked", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ValECompatibility.NoOverclaimState = VerifierEcosystemValENoOverclaimStateBlocked
		}},
		{name: "vala dependency partial blocks", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.Dependency.ValACurrentState = DeveloperEcosystemValAStatePartial
		}},
		{name: "repo config unsupported schema version fails closed", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.RepoConfigSchema.SchemaVersion = "changelock_repo_config.v2_unknown"
		}},
		{name: "repo config cannot override enterprise policy", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.RepoConfigSchema.EnterprisePolicyOverrideAttempt = true
		}},
		{name: "repo config cannot disable canonical evidence rules", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.RepoConfigSchema.DisableCanonicalEvidenceRules = true
		}},
		{name: "repo config cannot grant approval authority", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.RepoConfigSchema.GrantApprovalAuthority = true
		}},
		{name: "malformed config non active", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.RepoConfigValidation.MalformedConfigDetected = true
		}},
		{name: "deprecated config not silently active", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.RepoConfigValidation.DeprecatedFieldDetected = true
		}},
		{name: "policy preview cannot approve deployment", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.PolicyPreview.ApprovesDeployment = true
		}},
		{name: "policy preview cannot hide production unknowns", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.PolicyPreview.ProductionOnlyUnknownsVisible = false
		}},
		{name: "local to ci continuity rejects missing descriptors", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.LocalCIContinuity.MissingDescriptors = true
		}},
		{name: "local pass like output cannot become ci pass", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.LocalCIContinuity.LocalPassBecomesCIPass = true
		}},
		{name: "sdk contract cannot expose hidden mutation path", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.DeveloperAPISDK.HiddenMutationPath = true
		}},
		{name: "sdk contract cannot approve deployment", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.DeveloperAPISDK.ApprovesDeployment = true
		}},
		{name: "unsupported sdk version non active", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.DeveloperAPISDK.UnsupportedVersionActive = true
		}},
		{name: "templates cannot imply compliance certification", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ExamplesTemplates.ComplianceCertificationClaim = true
		}},
		{name: "stale templates not silently active", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ExamplesTemplates.StaleTemplateReferenceDetected = true
		}},
		{name: "unknown api version fails closed", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.APIVersioning.UnknownVersionDetected = true
		}},
		{name: "unsupported api version fails closed", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.APIVersioning.UnsupportedVersionDetected = true
		}},
		{name: "no overclaim point8 pass blocks", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.NoOverclaim.Point8PassClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValBModel()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemValBIntegration(model)
		if model.CurrentState == DeveloperEcosystemValBStateActive {
			t.Fatalf("expected %s to prevent active state, got %#v", tc.name, model)
		}
		if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
			t.Fatalf("expected %s to keep point 8 not complete, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemValBValECompatibilityGate(t *testing.T) {
	model := activeDeveloperEcosystemValBModel()
	if model.ValECompatibilityState != DeveloperEcosystemValBValECompatibilityStateActive {
		t.Fatalf("expected active Val E compatibility gate, got %#v", model)
	}

	mutated := activeDeveloperEcosystemValBModel()
	mutated.ValECompatibility.Point7PassReason = VerifierEcosystemValEPoint7PassReasonBlocked
	mutated = ComputeDeveloperEcosystemValBIntegration(mutated)
	if mutated.ValECompatibilityState == DeveloperEcosystemValBValECompatibilityStateActive {
		t.Fatalf("expected canonical blocked Val E reason to keep compatibility gate non-active, got %#v", mutated)
	}

	mutated = activeDeveloperEcosystemValBModel()
	mutated.ValECompatibility.Point7PassReason = VerifierEcosystemValEPoint7PassReasonAllowed + " certified"
	mutated = ComputeDeveloperEcosystemValBIntegration(mutated)
	if mutated.ValECompatibilityState != DeveloperEcosystemValBValECompatibilityStateBlocked {
		t.Fatalf("expected extended allowed reason to block compatibility gate, got %#v", mutated)
	}
}

func TestDeveloperEcosystemValBAPIVersionIdentityAndWindowValidation(t *testing.T) {
	model := activeDeveloperEcosystemValBModel()
	if got := EvaluateDeveloperEcosystemValBAPIVersioningState(model.APIVersioning); got != DeveloperEcosystemValBAPIVersioningStateActive {
		t.Fatalf("expected default API versioning contract to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*DeveloperEcosystemValBAPIVersioningCompatibility)
		expected string
	}{
		{name: "empty version identity incomplete", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.VersionIdentity = ""
		}, expected: DeveloperEcosystemValBAPIVersioningStateIncomplete},
		{name: "mismatched version identity unknown", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.VersionIdentity = "developer_api_sdk_surface.v2"
		}, expected: DeveloperEcosystemValBAPIVersioningStateUnknown},
		{name: "empty compatibility window incomplete", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.CompatibilityWindow = ""
		}, expected: DeveloperEcosystemValBAPIVersioningStateIncomplete},
		{name: "mismatched compatibility window unknown", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.CompatibilityWindow = "multi_major_unbounded_window"
		}, expected: DeveloperEcosystemValBAPIVersioningStateUnknown},
		{name: "unsupported version detected blocked", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.UnsupportedVersionDetected = true
		}, expected: DeveloperEcosystemValBAPIVersioningStateBlocked},
		{name: "unknown version detected unknown", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.UnknownVersionDetected = true
		}, expected: DeveloperEcosystemValBAPIVersioningStateUnknown},
		{name: "deprecated version detected partial", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.DeprecatedVersionDetected = true
		}, expected: DeveloperEcosystemValBAPIVersioningStatePartial},
		{name: "migration hint cannot rescue unsupported", mutate: func(model *DeveloperEcosystemValBAPIVersioningCompatibility) {
			model.UnsupportedVersionDetected = true
			model.MigrationHint = "upgrade later but still active"
		}, expected: DeveloperEcosystemValBAPIVersioningStateBlocked},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValBModel()
		tc.mutate(&mutated.APIVersioning)
		if got := EvaluateDeveloperEcosystemValBAPIVersioningState(mutated.APIVersioning); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestDeveloperEcosystemValBProofSurfaceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemValBModel()
	if got := EvaluateDeveloperEcosystemValBProofsState(model, activeDeveloperEcosystemValBLimitations()); got != DeveloperEcosystemValBStateActive {
		t.Fatalf("expected exact Val B proof surface set to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValBIntegration)
	}{
		{name: "missing vale closure fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/closure")
		}},
		{name: "missing vala proofs fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/vala/proofs")
		}},
		{name: "missing valb status fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/valb/status")
		}},
		{name: "duplicate proof ref fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/valb/proofs")
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/developer-ecosystem/val0/proofs")
		}},
		{name: "unknown extra proof ref fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/developer-ecosystem/valb/extra")
		}},
		{name: "whitespace proof ref fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.ProofSurfaceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValBModel()
		tc.mutate(&mutated)
		if got := EvaluateDeveloperEcosystemValBProofsState(mutated, activeDeveloperEcosystemValBLimitations()); got == DeveloperEcosystemValBStateActive {
			t.Fatalf("expected %s to fail exact proof validation, got %q", tc.name, got)
		}
	}
}

func TestDeveloperEcosystemValBEvidenceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemValBModel()
	if !DeveloperEcosystemValBProofEvidenceQualityValid(developerEcosystemValBEvidence(), model.EvidenceRefs) {
		t.Fatalf("expected exact developer Val B evidence refs to be valid")
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValBIntegration)
	}{
		{name: "missing vale compatibility evidence fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point7_vale_compatibility_gate")
		}},
		{name: "missing api sdk evidence fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:developer-api-sdk-surface-001")
		}},
		{name: "duplicate evidence ref fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:point8-valb-governance-001")
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:developer-api-sdk-surface-001")
		}},
		{name: "unknown extra evidence ref fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:developer-valb-extra-001")
		}},
		{name: "whitespace evidence ref fails", mutate: func(model *DeveloperEcosystemValBIntegration) {
			model.EvidenceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValBModel()
		tc.mutate(&mutated)
		if DeveloperEcosystemValBProofEvidenceQualityValid(developerEcosystemValBEvidence(), mutated.EvidenceRefs) {
			t.Fatalf("expected %s to fail exact evidence validation", tc.name)
		}
	}
}
