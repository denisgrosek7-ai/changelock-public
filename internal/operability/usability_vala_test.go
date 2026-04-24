package operability

import "testing"

func TestProductionUsabilityValAStateRequiresActiveVal0(t *testing.T) {
	got := EvaluateProductionUsabilityValAState(
		ProductionUsabilityVal0StateSubstantial,
		ProductionUsabilityValAConfigFactoryStateActive,
		ProductionUsabilityValABootstrapValidationStateActive,
		ProductionUsabilityValAPolicySchemaStateActive,
		ProductionUsabilityValAEffectiveConfigStateActive,
		ProductionUsabilityValARejectionLayerStateActive,
		ProductionUsabilityValADryRunStateActive,
		ProductionUsabilityValAExplainStateActive,
		ProductionUsabilityValARecoveryGuidanceStateActive,
		ProductionUsabilityValAFirstRunStateActive,
		ProductionUsabilityValAUpgradePreviewStateActive,
	)
	if got != ProductionUsabilityValAStateIncomplete {
		t.Fatalf("expected incomplete Val A state without active Val 0, got %q", got)
	}
}

func TestProductionUsabilityValAUnknownConfigFieldsRejectByDefault(t *testing.T) {
	model := ProductionUsabilityValAConfigFactory()
	model.UnknownFieldsDetected = []string{"mystery_field"}
	if got := EvaluateProductionUsabilityValAConfigFactoryState(model); got == ProductionUsabilityValAConfigFactoryStateActive {
		t.Fatalf("expected non-active config factory when unknown fields are rejected by default, got %q", got)
	}
}

func TestProductionUsabilityValAUnknownConfigFieldsWarnOnlyWhenExplicitlyAllowed(t *testing.T) {
	model := ProductionUsabilityValAConfigFactory()
	model.UnknownFieldsDetected = []string{"extra_preview_toggle"}
	model.CurrentUnknownFieldPolicy = ProductionUsabilityUnknownFieldWarn
	model.ExplicitUnknownFieldAllowance = true
	if got := EvaluateProductionUsabilityValAConfigFactoryState(model); got != ProductionUsabilityValAConfigFactoryStateActive {
		t.Fatalf("expected active config factory when unknown fields are explicitly allowed with warning, got %q", got)
	}
}

func TestProductionUsabilityValAUnsupportedSchemaVersionFailsClosed(t *testing.T) {
	model := ProductionUsabilityValAConfigFactory()
	model.SchemaVersion = "point4.production_usability.vala.config.v999"
	if got := EvaluateProductionUsabilityValAConfigFactoryState(model); got == ProductionUsabilityValAConfigFactoryStateActive {
		t.Fatalf("expected non-active config factory for unsupported schema version, got %q", got)
	}
}

func TestProductionUsabilityValADeprecatedSchemaRequiresWarningsAndMetadata(t *testing.T) {
	model := ProductionUsabilityValAConfigFactory()
	model.SchemaVersion = "point4.production_usability.vala.config.v0"
	model.CurrentCompatibility = ProductionUsabilityCompatibilityDeprecated
	model.CompatibilityWarnings = []string{"deprecated_schema_requires_upgrade_planning"}
	if got := EvaluateProductionUsabilityValAConfigFactoryState(model); got != ProductionUsabilityValAConfigFactoryStateActive {
		t.Fatalf("expected active config factory for deprecated schema with warnings, got %q", got)
	}
}

func TestProductionUsabilityValAMigrationWarningsDoNotEqualCompletedMigration(t *testing.T) {
	model := ProductionUsabilityValAConfigFactory()
	model.CurrentCompatibility = ProductionUsabilityCompatibilityMigrationRequired
	model.MigrationWarnings = []string{"rename_timeout_field"}
	model.MigrationCompleted = true
	if got := EvaluateProductionUsabilityValAConfigFactoryState(model); got == ProductionUsabilityValAConfigFactoryStateActive {
		t.Fatalf("expected non-active config factory when migration warnings are treated as completed migration, got %q", got)
	}
}

func TestProductionUsabilityValAInvalidConfigBlocksFailFastBootstrap(t *testing.T) {
	model := ProductionUsabilityValABootstrapValidation()
	model.ConfigValidationResult = ProductionUsabilityValidationInvalid
	model.BootstrapDisposition = ProductionUsabilityBootstrapAllowed
	model.BlockingReasonCodes = []string{"missing_required_field"}
	if got := EvaluateProductionUsabilityValABootstrapValidationState(model); got == ProductionUsabilityValABootstrapValidationStateActive {
		t.Fatalf("expected non-active bootstrap validation when invalid config is still allowed, got %q", got)
	}
}

func TestProductionUsabilityValADegradedBootstrapRequiresExplicitBoundaries(t *testing.T) {
	model := ProductionUsabilityValABootstrapValidation()
	model.ConfigValidationResult = ProductionUsabilityValidationDegraded
	model.BootstrapDisposition = ProductionUsabilityBootstrapDegradedAllowed
	model.DegradedAllowedExplicit = true
	model.DegradedBoundaries = nil
	if got := EvaluateProductionUsabilityValABootstrapValidationState(model); got == ProductionUsabilityValABootstrapValidationStateActive {
		t.Fatalf("expected non-active bootstrap validation without degraded boundaries, got %q", got)
	}
}

func TestProductionUsabilityValAEffectiveConfigSeparatesDefaultsFromUserFields(t *testing.T) {
	model := ProductionUsabilityValAEffectiveConfigInspection()
	model.UserProvidedFields = append(model.UserProvidedFields, "audit_store")
	if got := EvaluateProductionUsabilityValAEffectiveConfigState(model); got == ProductionUsabilityValAEffectiveConfigStateActive {
		t.Fatalf("expected non-active effective config when defaults overlap user-provided fields, got %q", got)
	}
}

func TestProductionUsabilityValAEffectiveConfigRedactsSecrets(t *testing.T) {
	model := ProductionUsabilityValAEffectiveConfigInspection()
	model.SecretsExposed = true
	if got := EvaluateProductionUsabilityValAEffectiveConfigState(model); got == ProductionUsabilityValAEffectiveConfigStateActive {
		t.Fatalf("expected non-active effective config when secrets are exposed, got %q", got)
	}
}

func TestProductionUsabilityValAEffectiveConfigRemainsProjectionOnly(t *testing.T) {
	model := ProductionUsabilityValAEffectiveConfigInspection()
	model.SourceProjectionDisclaimer = "effective_config_is_canonical_truth"
	if got := EvaluateProductionUsabilityValAEffectiveConfigState(model); got == ProductionUsabilityValAEffectiveConfigStateActive {
		t.Fatalf("expected non-active effective config without projection-only disclaimer, got %q", got)
	}
}

func TestProductionUsabilityValAHumanReadableRejectionIncludesRequiredFields(t *testing.T) {
	model := ProductionUsabilityValAHumanReadableRejectionLayer()
	model.RequiredFields = []string{"reason_code", "severity"}
	if got := EvaluateProductionUsabilityValARejectionLayerState(model); got == ProductionUsabilityValARejectionLayerStateActive {
		t.Fatalf("expected non-active rejection layer without required explainability fields, got %q", got)
	}
}

func TestProductionUsabilityValAHumanReadableRejectionDoesNotLeakSecretsInRestrictedScopes(t *testing.T) {
	model := ProductionUsabilityValAHumanReadableRejectionLayer()
	model.SecretsRedacted = false
	if got := EvaluateProductionUsabilityValARejectionLayerState(model); got == ProductionUsabilityValARejectionLayerStateActive {
		t.Fatalf("expected non-active rejection layer when secrets are not redacted, got %q", got)
	}
}

func TestProductionUsabilityValARedactedExplainOutputRemainsHonestAboutHiddenEvidence(t *testing.T) {
	model := ProductionUsabilityValAPermissionAwareExplainOutputs()
	for idx := range model.Items {
		if model.Items[idx].VisibilityScope == ProductionUsabilityVisibilityPartner {
			model.Items[idx].HiddenEvidenceMarker = ""
			break
		}
	}
	if got := EvaluateProductionUsabilityValAExplainState(model); got == ProductionUsabilityValAExplainStateActive {
		t.Fatalf("expected non-active explain output when hidden evidence marker is missing, got %q", got)
	}
}

func TestProductionUsabilityValAPolicySchemaUnsupportedFailsClosed(t *testing.T) {
	model := ProductionUsabilityValAPolicySchemaDiscipline()
	model.CurrentCompatibility = ProductionUsabilityCompatibilityUnsupported
	if got := EvaluateProductionUsabilityValAPolicySchemaState(model); got == ProductionUsabilityValAPolicySchemaStateActive {
		t.Fatalf("expected non-active policy schema discipline for unsupported policy schema, got %q", got)
	}
}

func TestProductionUsabilityValADryRunAndAuditOnlyDoNotMutateCanonicalState(t *testing.T) {
	model := ProductionUsabilityValAPolicyDryRunAuditFlow()
	model.MutatesCanonicalState = true
	if got := EvaluateProductionUsabilityValADryRunState(model); got == ProductionUsabilityValADryRunStateActive {
		t.Fatalf("expected non-active dry-run flow when canonical state mutates, got %q", got)
	}
}

func TestProductionUsabilityValADryRunSuccessDoesNotEqualProductionActivation(t *testing.T) {
	model := ProductionUsabilityValAPolicyDryRunAuditFlow()
	model.DryRunSuccessImpliesActivate = true
	if got := EvaluateProductionUsabilityValADryRunState(model); got == ProductionUsabilityValADryRunStateActive {
		t.Fatalf("expected non-active dry-run flow when dry-run success implies activation, got %q", got)
	}
}

func TestProductionUsabilityValARecoveryGuidanceDoesNotSuggestUnsafeRetryOrPolicyBypass(t *testing.T) {
	model := ProductionUsabilityValARecoveryGuidance()
	model.Items[0].UnsafeRetrySuggested = true
	if got := EvaluateProductionUsabilityValARecoveryGuidanceState(model); got == ProductionUsabilityValARecoveryGuidanceStateActive {
		t.Fatalf("expected non-active recovery guidance when unsafe retry is suggested, got %q", got)
	}

	model = ProductionUsabilityValARecoveryGuidance()
	model.Items[0].PolicyBypassSuggested = true
	if got := EvaluateProductionUsabilityValARecoveryGuidanceState(model); got == ProductionUsabilityValARecoveryGuidanceStateActive {
		t.Fatalf("expected non-active recovery guidance when policy bypass is suggested, got %q", got)
	}
}

func TestProductionUsabilityValAFirstRunSampleConfigDoesNotAutoEnableProduction(t *testing.T) {
	model := ProductionUsabilityValAFirstRunSafeBootstrap()
	model.AutoEnablesProduction = true
	if got := EvaluateProductionUsabilityValAFirstRunState(model); got == ProductionUsabilityValAFirstRunStateActive {
		t.Fatalf("expected non-active first-run bootstrap when sample config auto-enables production, got %q", got)
	}
}

func TestProductionUsabilityValAUpgradeImpactPreviewFailsClosedOnUnknownTargetSchema(t *testing.T) {
	model := ProductionUsabilityValAUpgradeImpactPreview()
	model.TargetSchemaVersion = "point4.production_usability.vala.config.v999"
	if got := EvaluateProductionUsabilityValAUpgradePreviewState(model); got == ProductionUsabilityValAUpgradePreviewStateActive {
		t.Fatalf("expected non-active upgrade preview for unknown target schema, got %q", got)
	}
}

func TestProductionUsabilityValAUpgradeImpactRollbackPerspectiveStaysBounded(t *testing.T) {
	model := ProductionUsabilityValAUpgradeImpactPreview()
	model.RollbackPerspective = "full_runtime_orchestrated_rollback"
	if got := EvaluateProductionUsabilityValAUpgradePreviewState(model); got == ProductionUsabilityValAUpgradePreviewStateActive {
		t.Fatalf("expected non-active upgrade preview when rollback perspective is not bounded, got %q", got)
	}
}

func TestProductionUsabilityValAProofsPassOnlyAsCoreWhilePoint4RemainsNotComplete(t *testing.T) {
	got := EvaluateProductionUsabilityValAProofsState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAConfigFactoryStateActive,
		ProductionUsabilityValABootstrapValidationStateActive,
		ProductionUsabilityValAPolicySchemaStateActive,
		ProductionUsabilityValAEffectiveConfigStateActive,
		ProductionUsabilityValARejectionLayerStateActive,
		ProductionUsabilityValADryRunStateActive,
		ProductionUsabilityValAExplainStateActive,
		ProductionUsabilityValARecoveryGuidanceStateActive,
		ProductionUsabilityValAFirstRunStateActive,
		ProductionUsabilityValAUpgradePreviewStateActive,
		[]string{
			"/v1/production/usability-operability-recovery/val0/proofs",
			"/v1/production/usability-operability-recovery/vala/config-factory",
			"/v1/production/usability-operability-recovery/vala/bootstrap-validation",
			"/v1/production/usability-operability-recovery/vala/policy-schema",
			"/v1/production/usability-operability-recovery/vala/effective-config",
			"/v1/production/usability-operability-recovery/vala/rejections",
			"/v1/production/usability-operability-recovery/vala/policy-dry-run",
			"/v1/production/usability-operability-recovery/vala/explain",
			"/v1/production/usability-operability-recovery/vala/recovery-guidance",
			"/v1/production/usability-operability-recovery/vala/first-run-bootstrap",
			"/v1/production/usability-operability-recovery/vala/upgrade-impact-preview",
			"/v1/production/usability-operability-recovery/vala/proofs",
		},
		[]string{"val0_proofs", "config_factory_core", "bootstrap_core", "policy_schema_core", "effective_config", "rejections", "dry_run"},
		[]string{"Val A only proves Config & Explainability Core."},
		[]string{"Točka 4 full PASS remains closed until later waves land."},
	)
	if got != ProductionUsabilityValAStateActive {
		t.Fatalf("expected active Val A proofs state for complete core, got %q", got)
	}
	if ProductionUsabilityPoint4StateNotComplete == ProductionUsabilityValAStateActive {
		t.Fatalf("point 4 state must remain distinct from active Val A core")
	}
}

func TestProductionUsabilityValAMissingRequiredComponentKeepsValANonActive(t *testing.T) {
	got := EvaluateProductionUsabilityValAState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAConfigFactoryStateActive,
		ProductionUsabilityValABootstrapValidationStateActive,
		ProductionUsabilityValAPolicySchemaStateActive,
		ProductionUsabilityValAEffectiveConfigStateActive,
		ProductionUsabilityValARejectionLayerStateActive,
		ProductionUsabilityValADryRunStateActive,
		ProductionUsabilityValAExplainStateActive,
		ProductionUsabilityValARecoveryGuidanceStateActive,
		ProductionUsabilityValAFirstRunStateIncomplete,
		ProductionUsabilityValAUpgradePreviewStateActive,
	)
	if got == ProductionUsabilityValAStateActive {
		t.Fatalf("expected non-active Val A state with missing required component, got %q", got)
	}
}

func TestProductionUsabilityValAFoundationIsActive(t *testing.T) {
	configFactory := ProductionUsabilityValAConfigFactory()
	if got := EvaluateProductionUsabilityValAConfigFactoryState(configFactory); got != ProductionUsabilityValAConfigFactoryStateActive {
		t.Fatalf("expected active config factory state, got %q", got)
	}

	bootstrap := ProductionUsabilityValABootstrapValidation()
	if got := EvaluateProductionUsabilityValABootstrapValidationState(bootstrap); got != ProductionUsabilityValABootstrapValidationStateActive {
		t.Fatalf("expected active bootstrap validation state, got %q", got)
	}

	policySchema := ProductionUsabilityValAPolicySchemaDiscipline()
	if got := EvaluateProductionUsabilityValAPolicySchemaState(policySchema); got != ProductionUsabilityValAPolicySchemaStateActive {
		t.Fatalf("expected active policy schema state, got %q", got)
	}

	effectiveConfig := ProductionUsabilityValAEffectiveConfigInspection()
	if got := EvaluateProductionUsabilityValAEffectiveConfigState(effectiveConfig); got != ProductionUsabilityValAEffectiveConfigStateActive {
		t.Fatalf("expected active effective config state, got %q", got)
	}

	rejectionLayer := ProductionUsabilityValAHumanReadableRejectionLayer()
	if got := EvaluateProductionUsabilityValARejectionLayerState(rejectionLayer); got != ProductionUsabilityValARejectionLayerStateActive {
		t.Fatalf("expected active rejection layer state, got %q", got)
	}

	dryRun := ProductionUsabilityValAPolicyDryRunAuditFlow()
	if got := EvaluateProductionUsabilityValADryRunState(dryRun); got != ProductionUsabilityValADryRunStateActive {
		t.Fatalf("expected active dry-run state, got %q", got)
	}

	explain := ProductionUsabilityValAPermissionAwareExplainOutputs()
	if got := EvaluateProductionUsabilityValAExplainState(explain); got != ProductionUsabilityValAExplainStateActive {
		t.Fatalf("expected active explain state, got %q", got)
	}

	recovery := ProductionUsabilityValARecoveryGuidance()
	if got := EvaluateProductionUsabilityValARecoveryGuidanceState(recovery); got != ProductionUsabilityValARecoveryGuidanceStateActive {
		t.Fatalf("expected active recovery guidance state, got %q", got)
	}

	firstRun := ProductionUsabilityValAFirstRunSafeBootstrap()
	if got := EvaluateProductionUsabilityValAFirstRunState(firstRun); got != ProductionUsabilityValAFirstRunStateActive {
		t.Fatalf("expected active first-run state, got %q", got)
	}

	upgradePreview := ProductionUsabilityValAUpgradeImpactPreview()
	if got := EvaluateProductionUsabilityValAUpgradePreviewState(upgradePreview); got != ProductionUsabilityValAUpgradePreviewStateActive {
		t.Fatalf("expected active upgrade preview state, got %q", got)
	}

	if got := EvaluateProductionUsabilityValAState(
		ProductionUsabilityVal0StateActive,
		EvaluateProductionUsabilityValAConfigFactoryState(configFactory),
		EvaluateProductionUsabilityValABootstrapValidationState(bootstrap),
		EvaluateProductionUsabilityValAPolicySchemaState(policySchema),
		EvaluateProductionUsabilityValAEffectiveConfigState(effectiveConfig),
		EvaluateProductionUsabilityValARejectionLayerState(rejectionLayer),
		EvaluateProductionUsabilityValADryRunState(dryRun),
		EvaluateProductionUsabilityValAExplainState(explain),
		EvaluateProductionUsabilityValARecoveryGuidanceState(recovery),
		EvaluateProductionUsabilityValAFirstRunState(firstRun),
		EvaluateProductionUsabilityValAUpgradePreviewState(upgradePreview),
	); got != ProductionUsabilityValAStateActive {
		t.Fatalf("expected active Val A state, got %q", got)
	}
}
