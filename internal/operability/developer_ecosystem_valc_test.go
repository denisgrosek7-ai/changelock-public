package operability

import "testing"

func activeDeveloperEcosystemValCModel() DeveloperEcosystemValCIntegration {
	model := DeveloperEcosystemValCIntegrationModel()
	model.ValECompatibility = DeveloperEcosystemValCValECompatibilityGateModel()
	model.ValBCompatibility = DeveloperEcosystemValCValBCompatibilityGateModel()
	model.Dependency = DeveloperEcosystemValCDependencySnapshot{
		ValBCurrentState:          DeveloperEcosystemValBStateActive,
		ValBPoint8State:           DeveloperEcosystemPoint8StateNotComplete,
		ValECompatibilityState:    DeveloperEcosystemValBValECompatibilityStateActive,
		DependencyState:           DeveloperEcosystemValBDependencyStateActive,
		RepoConfigSchemaState:     DeveloperEcosystemValBRepoConfigSchemaStateActive,
		RepoConfigValidationState: DeveloperEcosystemValBRepoConfigValidationStateActive,
		PolicyPreviewState:        DeveloperEcosystemValBPolicyPreviewStateActive,
		LocalCIContinuityState:    DeveloperEcosystemValBLocalCIContinuityStateActive,
		APISDKSurfaceState:        DeveloperEcosystemValBAPISDKSurfaceStateActive,
		ExamplesTemplatesState:    DeveloperEcosystemValBExamplesTemplatesStateActive,
		APIVersioningState:        DeveloperEcosystemValBAPIVersioningStateActive,
		NoOverclaimState:          DeveloperEcosystemValBNoOverclaimStateActive,
		ValBProofSurfaceRefs:      DeveloperEcosystemValBProofSurfaceRefs(),
		ValBEvidenceRefs:          DeveloperEcosystemValBProofEvidenceRefs(),
		ValBProjectionDisclaimer:  developerEcosystemValBProjectionDisclaimer(),
	}
	return ComputeDeveloperEcosystemValCIntegration(model)
}

func activeDeveloperEcosystemValCLimitations() []string {
	return []string{
		"Val C implements plugin and extensibility contracts only and does not implement a plugin runtime, marketplace, external registry, remote installation, production SDK runtime, or Točka 9 work.",
		"Točka 8 remains not complete because later developer ecosystem waves are still required before any integrated closure can exist.",
		"Plugin manifests, diagnostics, custom checks, compatibility descriptors, and samples remain advisory only and cannot approve deployment, certify trust, or create canonical evidence.",
	}
}

func TestDeveloperEcosystemValCHappyPathAndPoint8NotComplete(t *testing.T) {
	model := activeDeveloperEcosystemValCModel()
	if model.CurrentState != DeveloperEcosystemValCStateActive {
		t.Fatalf("expected active developer Val C state, got %#v", model)
	}
	if model.PluginPerformance.PluginExecutionBudgetRef != DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		t.Fatalf("expected canonical Val 0 budget ref, got %#v", model.PluginPerformance)
	}
	if model.SandboxIsolation.DisciplineID != DeveloperEcosystemValCSandboxIsolationDisciplineID ||
		model.SandboxIsolation.Version != DeveloperEcosystemValCSandboxIsolationVersion {
		t.Fatalf("expected canonical sandbox identity metadata, got %#v", model.SandboxIsolation)
	}
	if model.ExtensionCompatibility.PluginAPIVersionIdentity != DeveloperEcosystemValCPluginAPIVersionIdentity ||
		model.ExtensionCompatibility.CompatibilityWindow != DeveloperEcosystemValCPluginAPICompatibilityWindow {
		t.Fatalf("expected canonical plugin API identity/window, got %#v", model.ExtensionCompatibility)
	}
	if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
		t.Fatalf("expected point 8 to remain not complete in Val C, got %#v", model)
	}
	if got := EvaluateDeveloperEcosystemValCProofsState(model, activeDeveloperEcosystemValCLimitations()); got != DeveloperEcosystemValCStateActive {
		t.Fatalf("expected active developer Val C proofs state, got %q", got)
	}
}

func TestDeveloperEcosystemValCValEAndValBCompatibilityGates(t *testing.T) {
	model := activeDeveloperEcosystemValCModel()
	if model.ValECompatibilityState != DeveloperEcosystemValCValECompatibilityStateActive {
		t.Fatalf("expected active Val E compatibility gate, got %#v", model)
	}
	if model.ValBCompatibilityState != DeveloperEcosystemValCValBCompatibilityStateActive {
		t.Fatalf("expected active Val B compatibility gate, got %#v", model)
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValCIntegration)
	}{
		{name: "vale point7 pass overclaim blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ValECompatibility.Point7PassReason = "point_7_pass production approved"
		}},
		{name: "vale no overclaim blocked", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ValECompatibility.NoOverclaimState = VerifierEcosystemValENoOverclaimStateBlocked
		}},
		{name: "valb repo schema compatibility weakened", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ValBCompatibility.RepoConfigCompatibilityBehavior = "permissive"
		}},
		{name: "valb api version identity weakened", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ValBCompatibility.APIVersionIdentity = "developer_api_sdk_surface.v2"
		}},
		{name: "valb api compatibility window weakened", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ValBCompatibility.APICompatibilityWindow = "multi_major_unbounded_window"
		}},
		{name: "valb dependency partial blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.Dependency.ValBCurrentState = DeveloperEcosystemValBStatePartial
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValCModel()
		tc.mutate(&mutated)
		mutated = ComputeDeveloperEcosystemValCIntegration(mutated)
		if mutated.CurrentState == DeveloperEcosystemValCStateActive {
			t.Fatalf("expected %s to prevent active state, got %#v", tc.name, mutated)
		}
		if mutated.Point8State != DeveloperEcosystemPoint8StateNotComplete {
			t.Fatalf("expected %s to keep point 8 not complete, got %#v", tc.name, mutated)
		}
	}
}

func TestDeveloperEcosystemValCPluginContractsBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValCIntegration)
	}{
		{name: "manifest unknown schema version fails closed", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginManifest.ManifestSchemaVersion = "changelock_plugin_manifest.v2_unknown"
		}},
		{name: "manifest missing required fields fails closed", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginManifest.PluginIdentity = ""
		}},
		{name: "unsupported extension point blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginManifest.RequestedExtensionPoints = append(model.PluginManifest.RequestedExtensionPoints, "unsupported_extension_point")
		}},
		{name: "privileged capability blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.CapabilityDeclaration.PrivilegedCapabilities = []string{"deployment_approval"}
		}},
		{name: "duplicate capability blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.CapabilityDeclaration.DeclaredCapabilities = append(model.CapabilityDeclaration.DeclaredCapabilities, "advisory_signal")
		}},
		{name: "revoked lifecycle blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginLifecycle.Revoked = true
		}},
		{name: "disabled lifecycle blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginLifecycle.Disabled = true
		}},
		{name: "unsupported lifecycle blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginLifecycle.Unsupported = true
		}},
		{name: "missing isolation declaration blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.ExecutionIsolationExpectation = ""
		}},
		{name: "hidden network access blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.HiddenNetworkAccess = true
		}},
		{name: "hidden file access blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.HiddenFileSystemAccess = true
		}},
		{name: "hidden secret access blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.HiddenSecretAccess = true
		}},
		{name: "hidden outbound mutation blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.HiddenOutboundMutationPath = true
		}},
		{name: "custom checks cannot approve deployment", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.BoundedCustomChecks.ApprovesDeployment = true
		}},
		{name: "custom checks cannot override policy", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.BoundedCustomChecks.OverridesEnterprisePolicy = true
		}},
		{name: "custom checks unsupported class fails closed", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.BoundedCustomChecks.OutputClass = "production_authorization"
		}},
		{name: "diagnostics cannot hide failures", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginDiagnostics.FailureReasonsVisible = false
		}},
		{name: "diagnostics cannot hide production unknowns", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginDiagnostics.ProductionOnlyUnknownVisible = false
		}},
		{name: "diagnostics cannot turn advisory into pass", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginDiagnostics.AdvisoryAsPass = true
		}},
		{name: "plugin performance old dangling budget ref blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginPerformance.PluginExecutionBudgetRef = "developer-performance-budget"
		}},
		{name: "plugin performance unknown budget ref blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginPerformance.PluginExecutionBudgetRef = "developer-ecosystem-performance-budget-next"
		}},
		{name: "silent timeout blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginPerformance.SilentTimeout = true
		}},
		{name: "silent bypass blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginPerformance.SilentBypass = true
		}},
		{name: "hidden failure suppression blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.PluginPerformance.HiddenFailureSuppression = true
		}},
		{name: "sample plugin cannot imply certification", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SamplePluginDescriptors.CertificationClaim = true
		}},
		{name: "sample plugin cannot imply production readiness", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SamplePluginDescriptors.ProductionReadinessClaim = true
		}},
		{name: "unknown plugin api version fails closed", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ExtensionCompatibility.UnknownVersionDetected = true
		}},
		{name: "unsupported plugin api version fails closed", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ExtensionCompatibility.UnsupportedVersionDetected = true
		}},
		{name: "revoked plugin api version fails closed", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ExtensionCompatibility.RevokedVersionDetected = true
		}},
		{name: "point8 pass claim blocks", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.NoOverclaim.Point8PassClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValCModel()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemValCIntegration(model)
		if model.CurrentState == DeveloperEcosystemValCStateActive {
			t.Fatalf("expected %s to prevent active state, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemValCPartialAndExactValidationStates(t *testing.T) {
	model := activeDeveloperEcosystemValCModel()
	model.PluginLifecycle.LifecycleState = DeveloperEcosystemPluginLifecycleDeprecated
	model = ComputeDeveloperEcosystemValCIntegration(model)
	if model.PluginLifecycleState != DeveloperEcosystemValCPluginLifecycleStatePartial {
		t.Fatalf("expected deprecated plugin lifecycle to become partial, got %#v", model.PluginLifecycle)
	}

	model = activeDeveloperEcosystemValCModel()
	model.PluginLifecycle.LifecycleState = DeveloperEcosystemPluginLifecycleDeprecated
	model.PluginLifecycle.DeprecatedCompatibilityVisible = false
	model = ComputeDeveloperEcosystemValCIntegration(model)
	if model.PluginLifecycleState != DeveloperEcosystemValCPluginLifecycleStateBlocked {
		t.Fatalf("expected deprecated plugin lifecycle without compatibility to block, got %#v", model.PluginLifecycle)
	}

	model = activeDeveloperEcosystemValCModel()
	model.PluginPerformance.PluginExecutionBudgetRef = DeveloperEcosystemVal0PerformanceBudgetDisciplineID
	if got := EvaluateDeveloperEcosystemValCPluginPerformanceState(model.PluginPerformance); got != DeveloperEcosystemValCPluginPerformanceStateActive {
		t.Fatalf("expected canonical plugin performance budget ref to be active, got %q", got)
	}
}

func TestDeveloperEcosystemValCSandboxIsolationIdentityValidation(t *testing.T) {
	model := DeveloperEcosystemValCSandboxIsolationExpectationModel()
	if got := EvaluateDeveloperEcosystemValCSandboxIsolationState(model); got != DeveloperEcosystemValCSandboxIsolationStateActive {
		t.Fatalf("expected canonical sandbox contract to be active, got %q", got)
	}

	testCases := []struct {
		name     string
		mutate   func(*DeveloperEcosystemValCSandboxIsolationExpectation)
		expected string
	}{
		{name: "blank discipline id incomplete", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.DisciplineID = ""
		}, expected: DeveloperEcosystemValCSandboxIsolationStateIncomplete},
		{name: "whitespace discipline id incomplete", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.DisciplineID = "   "
		}, expected: DeveloperEcosystemValCSandboxIsolationStateIncomplete},
		{name: "blank version incomplete", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.Version = ""
		}, expected: DeveloperEcosystemValCSandboxIsolationStateIncomplete},
		{name: "whitespace version incomplete", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.Version = "  "
		}, expected: DeveloperEcosystemValCSandboxIsolationStateIncomplete},
		{name: "unknown discipline id unknown", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.DisciplineID = "developer-ecosystem-plugin-sandbox-next"
		}, expected: DeveloperEcosystemValCSandboxIsolationStateUnknown},
		{name: "unknown version unknown", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.Version = "2026.05"
		}, expected: DeveloperEcosystemValCSandboxIsolationStateUnknown},
		{name: "hidden network access blocked", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.HiddenNetworkAccess = true
		}, expected: DeveloperEcosystemValCSandboxIsolationStateBlocked},
		{name: "hidden file access blocked", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.HiddenFileSystemAccess = true
		}, expected: DeveloperEcosystemValCSandboxIsolationStateBlocked},
		{name: "hidden secret access blocked", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.HiddenSecretAccess = true
		}, expected: DeveloperEcosystemValCSandboxIsolationStateBlocked},
		{name: "hidden outbound mutation blocked", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.HiddenOutboundMutationPath = true
		}, expected: DeveloperEcosystemValCSandboxIsolationStateBlocked},
		{name: "missing isolation declaration blocked", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.ExecutionIsolationExpectation = ""
		}, expected: DeveloperEcosystemValCSandboxIsolationStateBlocked},
		{name: "sandbox bypass blocked", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.SandboxBypassClaim = true
		}, expected: DeveloperEcosystemValCSandboxIsolationStateBlocked},
		{name: "production safety certification blocked", mutate: func(model *DeveloperEcosystemValCSandboxIsolationExpectation) {
			model.ProductionSafetyCertificationClaim = true
		}, expected: DeveloperEcosystemValCSandboxIsolationStateBlocked},
	}

	for _, tc := range testCases {
		mutated := DeveloperEcosystemValCSandboxIsolationExpectationModel()
		tc.mutate(&mutated)
		if got := EvaluateDeveloperEcosystemValCSandboxIsolationState(mutated); got != tc.expected {
			t.Fatalf("expected %s to return %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestDeveloperEcosystemValCSandboxIsolationPreventsOverallActiveState(t *testing.T) {
	testCases := []struct {
		name            string
		mutate          func(*DeveloperEcosystemValCIntegration)
		expectedSandbox string
	}{
		{name: "blank discipline id", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.DisciplineID = ""
		}, expectedSandbox: DeveloperEcosystemValCSandboxIsolationStateIncomplete},
		{name: "blank version", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.Version = ""
		}, expectedSandbox: DeveloperEcosystemValCSandboxIsolationStateIncomplete},
		{name: "unknown discipline id", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.DisciplineID = "developer-ecosystem-plugin-sandbox-next"
		}, expectedSandbox: DeveloperEcosystemValCSandboxIsolationStateUnknown},
		{name: "production safety certification claim", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.SandboxIsolation.ProductionSafetyCertificationClaim = true
		}, expectedSandbox: DeveloperEcosystemValCSandboxIsolationStateBlocked},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValCModel()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemValCIntegration(model)
		if model.SandboxIsolationState != tc.expectedSandbox {
			t.Fatalf("expected %s to produce sandbox state %q, got %#v", tc.name, tc.expectedSandbox, model.SandboxIsolation)
		}
		if model.CurrentState == DeveloperEcosystemValCStateActive {
			t.Fatalf("expected %s to prevent overall active state, got %#v", tc.name, model)
		}
		if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
			t.Fatalf("expected %s to keep point 8 not complete, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemValCProofSurfaceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemValCModel()
	if got := EvaluateDeveloperEcosystemValCProofsState(model, activeDeveloperEcosystemValCLimitations()); got != DeveloperEcosystemValCStateActive {
		t.Fatalf("expected exact Val C proof surface set to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValCIntegration)
	}{
		{name: "missing vale closure fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/closure")
		}},
		{name: "missing valb proofs fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/valb/proofs")
		}},
		{name: "missing valc status fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/valc/status")
		}},
		{name: "duplicate proof ref fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/valc/proofs")
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/developer-ecosystem/val0/proofs")
		}},
		{name: "unknown extra proof ref fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/developer-ecosystem/valc/extra")
		}},
		{name: "whitespace proof ref fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.ProofSurfaceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValCModel()
		tc.mutate(&mutated)
		if got := EvaluateDeveloperEcosystemValCProofsState(mutated, activeDeveloperEcosystemValCLimitations()); got == DeveloperEcosystemValCStateActive {
			t.Fatalf("expected %s to prevent active proofs state, got %q", tc.name, got)
		}
	}
}

func TestDeveloperEcosystemValCEvidenceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemValCModel()
	if !DeveloperEcosystemValCProofEvidenceQualityValid(developerEcosystemValCEvidence(), model.EvidenceRefs) {
		t.Fatalf("expected exact Val C evidence set to be valid")
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValCIntegration)
	}{
		{name: "missing plugin manifest evidence fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:developer-plugin-manifest-001")
		}},
		{name: "missing valb evidence fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "point8_repo_sdk_integration")
		}},
		{name: "duplicate evidence ref fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:point8-valc-governance-001")
			model.EvidenceRefs = append(model.EvidenceRefs, "point7_vale_compatibility_gate")
		}},
		{name: "unknown extra evidence ref fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:developer-plugin-unknown-999")
		}},
		{name: "whitespace evidence ref fails", mutate: func(model *DeveloperEcosystemValCIntegration) {
			model.EvidenceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValCModel()
		tc.mutate(&mutated)
		if DeveloperEcosystemValCProofEvidenceQualityValid(developerEcosystemValCEvidence(), mutated.EvidenceRefs) {
			t.Fatalf("expected %s to invalidate evidence set", tc.name)
		}
	}
}
