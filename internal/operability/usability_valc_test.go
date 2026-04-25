package operability

import "testing"

func TestProductionUsabilityValCStateRequiresActiveVal0(t *testing.T) {
	got := EvaluateProductionUsabilityValCState(
		ProductionUsabilityVal0StateSubstantial,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCReadinessStateActive,
		ProductionUsabilityValCGuidedReadinessStateActive,
		ProductionUsabilityValCSupportBundleStateActive,
		ProductionUsabilityValCDiagnosticsStateActive,
		ProductionUsabilityValCHealthSnapshotStateActive,
		ProductionUsabilityValCRecoveryPlaybookStateActive,
		ProductionUsabilityValCUpgradeAdvisoryStateActive,
		ProductionUsabilityValCPermissionSupportStateActive,
		ProductionUsabilityValCExportSafetyStateActive,
	)
	if got != ProductionUsabilityValCStateIncomplete {
		t.Fatalf("expected incomplete Val C state without active Val 0, got %q", got)
	}
}

func TestProductionUsabilityValCStateRequiresActiveValA(t *testing.T) {
	got := EvaluateProductionUsabilityValCState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateSubstantial,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCReadinessStateActive,
		ProductionUsabilityValCGuidedReadinessStateActive,
		ProductionUsabilityValCSupportBundleStateActive,
		ProductionUsabilityValCDiagnosticsStateActive,
		ProductionUsabilityValCHealthSnapshotStateActive,
		ProductionUsabilityValCRecoveryPlaybookStateActive,
		ProductionUsabilityValCUpgradeAdvisoryStateActive,
		ProductionUsabilityValCPermissionSupportStateActive,
		ProductionUsabilityValCExportSafetyStateActive,
	)
	if got != ProductionUsabilityValCStateIncomplete {
		t.Fatalf("expected incomplete Val C state without active Val A, got %q", got)
	}
}

func TestProductionUsabilityValCStateRequiresActiveValB(t *testing.T) {
	got := EvaluateProductionUsabilityValCState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateSubstantial,
		ProductionUsabilityValCReadinessStateActive,
		ProductionUsabilityValCGuidedReadinessStateActive,
		ProductionUsabilityValCSupportBundleStateActive,
		ProductionUsabilityValCDiagnosticsStateActive,
		ProductionUsabilityValCHealthSnapshotStateActive,
		ProductionUsabilityValCRecoveryPlaybookStateActive,
		ProductionUsabilityValCUpgradeAdvisoryStateActive,
		ProductionUsabilityValCPermissionSupportStateActive,
		ProductionUsabilityValCExportSafetyStateActive,
	)
	if got != ProductionUsabilityValCStateIncomplete {
		t.Fatalf("expected incomplete Val C state without active Val B, got %q", got)
	}
}

func TestProductionUsabilityValCReadinessPassBlockedWhenBlockingChecksFail(t *testing.T) {
	model := ProductionUsabilityValCReadinessChecks()
	model.Items[0].Status = ProductionUsabilityReadinessFail
	if got := EvaluateProductionUsabilityValCReadinessState(model); got == ProductionUsabilityValCReadinessStateActive {
		t.Fatalf("expected non-active readiness when blocking check fails, got %q", got)
	}
}

func TestProductionUsabilityValCReadinessWarningIsVisible(t *testing.T) {
	model := ProductionUsabilityValCReadinessChecks()
	hasWarning := false
	for _, item := range model.Items {
		if item.Status == ProductionUsabilityReadinessWarning {
			hasWarning = true
		}
	}
	if !hasWarning {
		t.Fatalf("expected sample readiness model to expose a visible warning check")
	}
	if got := EvaluateProductionUsabilityValCReadinessState(model); got != ProductionUsabilityValCReadinessStateActive {
		t.Fatalf("expected active readiness state with visible non-blocking warning, got %q", got)
	}
}

func TestProductionUsabilityValCDegradedReadinessCannotReportFullPass(t *testing.T) {
	model := ProductionUsabilityValCReadinessChecks()
	model.Items[1].Status = ProductionUsabilityReadinessDegraded
	if got := EvaluateProductionUsabilityValCReadinessState(model); got == ProductionUsabilityValCReadinessStateActive {
		t.Fatalf("expected non-active readiness when degraded status is present, got %q", got)
	}
}

func TestProductionUsabilityValCUnsupportedAndNotRunReadinessAreNotPass(t *testing.T) {
	model := ProductionUsabilityValCReadinessChecks()
	model.Items[1].Status = ProductionUsabilityReadinessUnsupported
	if got := EvaluateProductionUsabilityValCReadinessState(model); got == ProductionUsabilityValCReadinessStateActive {
		t.Fatalf("expected non-active readiness when unsupported status is present, got %q", got)
	}

	model = ProductionUsabilityValCReadinessChecks()
	model.Items[1].Status = ProductionUsabilityReadinessNotRun
	if got := EvaluateProductionUsabilityValCReadinessState(model); got == ProductionUsabilityValCReadinessStateActive {
		t.Fatalf("expected non-active readiness when not_run status is present, got %q", got)
	}
}

func TestProductionUsabilityValCGuidedReadinessBlocksGoLiveWhenBlockingStepsMissing(t *testing.T) {
	model := ProductionUsabilityValCGuidedReadiness()
	model.BlockingSteps = []string{"support_bundle_path_verified"}
	model.GoLiveAllowed = true
	if got := EvaluateProductionUsabilityValCGuidedReadinessState(model); got == ProductionUsabilityValCGuidedReadinessStateActive {
		t.Fatalf("expected non-active guided readiness when blocking steps remain, got %q", got)
	}
}

func TestProductionUsabilityValCSampleConfigDoesNotAutoEnableProductionMode(t *testing.T) {
	model := ProductionUsabilityValCGuidedReadiness()
	model.SampleConfigDetected = true
	model.AutoProductionEnablement = true
	if got := EvaluateProductionUsabilityValCGuidedReadinessState(model); got == ProductionUsabilityValCGuidedReadinessStateActive {
		t.Fatalf("expected non-active guided readiness when sample config auto-enables production, got %q", got)
	}
}

func TestProductionUsabilityValCFakeDemoEvidenceBlocksGoLiveReadiness(t *testing.T) {
	model := ProductionUsabilityValCGuidedReadiness()
	model.FakeDemoEvidenceDetected = true
	if got := EvaluateProductionUsabilityValCGuidedReadinessState(model); got == ProductionUsabilityValCGuidedReadinessStateActive {
		t.Fatalf("expected non-active guided readiness with fake demo evidence, got %q", got)
	}
}

func TestProductionUsabilityValCSupportBundleWithoutManifestFailsClosed(t *testing.T) {
	model := ProductionUsabilityValCSupportBundleQualityGate()
	model.ManifestPresent = false
	if got := EvaluateProductionUsabilityValCSupportBundleState(model); got == ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected non-active support bundle state without manifest, got %q", got)
	}
}

func TestProductionUsabilityValCSupportBundleWithRawSecretsTokensOrEnvFailsClosed(t *testing.T) {
	model := ProductionUsabilityValCSupportBundleQualityGate()
	model.RawSecretDetected = true
	if got := EvaluateProductionUsabilityValCSupportBundleState(model); got == ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected non-active support bundle state with raw secret, got %q", got)
	}

	model = ProductionUsabilityValCSupportBundleQualityGate()
	model.RawTokenDetected = true
	if got := EvaluateProductionUsabilityValCSupportBundleState(model); got == ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected non-active support bundle state with raw token, got %q", got)
	}

	model = ProductionUsabilityValCSupportBundleQualityGate()
	model.UnfilteredEnvDetected = true
	if got := EvaluateProductionUsabilityValCSupportBundleState(model); got == ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected non-active support bundle state with unfiltered env, got %q", got)
	}
}

func TestProductionUsabilityValCSupportBundleCannotClaimLocalCacheAsCanonicalTruth(t *testing.T) {
	model := ProductionUsabilityValCSupportBundleQualityGate()
	model.CacheClaimsCanonicalTruth = true
	if got := EvaluateProductionUsabilityValCSupportBundleState(model); got == ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected non-active support bundle state when cache claims canonical truth, got %q", got)
	}
}

func TestProductionUsabilityValCExcludedSupportBundleSectionsRequireReasons(t *testing.T) {
	model := ProductionUsabilityValCSupportBundleQualityGate()
	model.ExclusionReasons = []string{"only_one_reason"}
	if got := EvaluateProductionUsabilityValCSupportBundleState(model); got == ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected non-active support bundle state when exclusion reasons do not match sections, got %q", got)
	}
}

func TestProductionUsabilityValCRedactedEvidenceRemainsRepresentedAsRedactedMetadata(t *testing.T) {
	model := ProductionUsabilityValCSupportBundleQualityGate()
	model.RedactedEvidenceRepresented = false
	if got := EvaluateProductionUsabilityValCSupportBundleState(model); got == ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected non-active support bundle state when redacted evidence is not represented, got %q", got)
	}
}

func TestProductionUsabilityValCDiagnosticsSafeToShareRequiresRedactionAndSecretScan(t *testing.T) {
	model := ProductionUsabilityValCDiagnosticsHardening()
	model.SensitiveFieldsRedacted = false
	if got := EvaluateProductionUsabilityValCDiagnosticsState(model); got == ProductionUsabilityValCDiagnosticsStateActive {
		t.Fatalf("expected non-active diagnostics without redaction, got %q", got)
	}

	model = ProductionUsabilityValCDiagnosticsHardening()
	model.SecretScanStatus = ProductionUsabilitySecretScanFailed
	if got := EvaluateProductionUsabilityValCDiagnosticsState(model); got == ProductionUsabilityValCDiagnosticsStateActive {
		t.Fatalf("expected non-active diagnostics with failed secret scan, got %q", got)
	}
}

func TestProductionUsabilityValCDiagnosticsExposeStalePartialUnsupportedSectionsExplicitly(t *testing.T) {
	model := ProductionUsabilityValCDiagnosticsHardening()
	model.UnsupportedSectionsExplicit = false
	if got := EvaluateProductionUsabilityValCDiagnosticsState(model); got == ProductionUsabilityValCDiagnosticsStateActive {
		t.Fatalf("expected non-active diagnostics when unsupported sections are not explicit, got %q", got)
	}

	model = ProductionUsabilityValCDiagnosticsHardening()
	model.StaleSectionsExplicit = false
	if got := EvaluateProductionUsabilityValCDiagnosticsState(model); got == ProductionUsabilityValCDiagnosticsStateActive {
		t.Fatalf("expected non-active diagnostics when stale sections are not explicit, got %q", got)
	}

	model = ProductionUsabilityValCDiagnosticsHardening()
	model.PartialSectionsExplicit = false
	if got := EvaluateProductionUsabilityValCDiagnosticsState(model); got == ProductionUsabilityValCDiagnosticsStateActive {
		t.Fatalf("expected non-active diagnostics when partial sections are not explicit, got %q", got)
	}
}

func TestProductionUsabilityValCHealthSnapshotHealthyIsBlockedByDegradedOrUnhealthyComponents(t *testing.T) {
	model := ProductionUsabilityValCHealthSnapshot()
	model.ComponentStates["api"] = ProductionUsabilityHealthDegraded
	model.DegradedComponents = []string{"api"}
	if got := EvaluateProductionUsabilityValCHealthSnapshotState(model); got == ProductionUsabilityValCHealthSnapshotStateActive {
		t.Fatalf("expected non-active health snapshot when healthy snapshot contains degraded component, got %q", got)
	}

	model = ProductionUsabilityValCHealthSnapshot()
	model.ComponentStates["db"] = ProductionUsabilityHealthUnhealthy
	model.DegradedComponents = []string{"db"}
	if got := EvaluateProductionUsabilityValCHealthSnapshotState(model); got == ProductionUsabilityValCHealthSnapshotStateActive {
		t.Fatalf("expected non-active health snapshot when healthy snapshot contains unhealthy component, got %q", got)
	}
}

func TestProductionUsabilityValCStaleHealthIsNotReportedAsFresh(t *testing.T) {
	model := ProductionUsabilityValCHealthSnapshot()
	model.StaleComponents = []string{"search"}
	if got := EvaluateProductionUsabilityValCHealthSnapshotState(model); got == ProductionUsabilityValCHealthSnapshotStateActive {
		t.Fatalf("expected non-active health snapshot when stale components are reported as fresh, got %q", got)
	}
}

func TestProductionUsabilityValCUnsupportedHealthComponentsAreNotHidden(t *testing.T) {
	model := ProductionUsabilityValCHealthSnapshot()
	model.ComponentStates["public_export"] = ProductionUsabilityHealthUnsupported
	if got := EvaluateProductionUsabilityValCHealthSnapshotState(model); got == ProductionUsabilityValCHealthSnapshotStateActive {
		t.Fatalf("expected non-active health snapshot when unsupported components are hidden, got %q", got)
	}
}

func TestProductionUsabilityValCRecoveryPlaybookDoesNotRecommendPolicyBypass(t *testing.T) {
	model := ProductionUsabilityValCRecoveryPlaybooks()
	model.Items[0].PolicyBypassSuggested = true
	if got := EvaluateProductionUsabilityValCRecoveryPlaybookState(model); got == ProductionUsabilityValCRecoveryPlaybookStateActive {
		t.Fatalf("expected non-active recovery playbook state when policy bypass is suggested, got %q", got)
	}
}

func TestProductionUsabilityValCRecoveryPlaybookDoesNotRecommendUnsafeRetryForRetryUnsafeOperations(t *testing.T) {
	model := ProductionUsabilityValCRecoveryPlaybooks()
	model.Items[0].UnsafeRetrySuggested = true
	if got := EvaluateProductionUsabilityValCRecoveryPlaybookState(model); got == ProductionUsabilityValCRecoveryPlaybookStateActive {
		t.Fatalf("expected non-active recovery playbook state when unsafe retry is suggested, got %q", got)
	}
}

func TestProductionUsabilityValCRecoveryDistinguishesSafeFromUnsafeSteps(t *testing.T) {
	model := ProductionUsabilityValCRecoveryPlaybooks()
	model.Items[0].UnsafeSteps = append(model.Items[0].UnsafeSteps, model.Items[0].SafeSteps[0])
	if got := EvaluateProductionUsabilityValCRecoveryPlaybookState(model); got == ProductionUsabilityValCRecoveryPlaybookStateActive {
		t.Fatalf("expected non-active recovery playbook state when safe and unsafe steps overlap, got %q", got)
	}
}

func TestProductionUsabilityValCUpgradeRollbackAdvisoryDoesNotMutateState(t *testing.T) {
	model := ProductionUsabilityValCUpgradeRollbackAdvisory()
	model.MutatesState = true
	if got := EvaluateProductionUsabilityValCUpgradeAdvisoryState(model); got == ProductionUsabilityValCUpgradeAdvisoryStateActive {
		t.Fatalf("expected non-active advisory state when it mutates state, got %q", got)
	}
}

func TestProductionUsabilityValCUnknownTargetVersionFailsClosed(t *testing.T) {
	model := ProductionUsabilityValCUpgradeRollbackAdvisory()
	model.TargetVersion = "2027.01.0"
	if got := EvaluateProductionUsabilityValCUpgradeAdvisoryState(model); got == ProductionUsabilityValCUpgradeAdvisoryStateActive {
		t.Fatalf("expected non-active advisory state for unknown target version, got %q", got)
	}
}

func TestProductionUsabilityValCRollbackAvailabilityIsBoundedAndScopeLimited(t *testing.T) {
	model := ProductionUsabilityValCUpgradeRollbackAdvisory()
	model.RollbackScope = ""
	if got := EvaluateProductionUsabilityValCUpgradeAdvisoryState(model); got == ProductionUsabilityValCUpgradeAdvisoryStateActive {
		t.Fatalf("expected non-active advisory state without rollback scope, got %q", got)
	}
}

func TestProductionUsabilityValCPreviewAuditOnlyAdvisoryDoesNotEqualApproval(t *testing.T) {
	model := ProductionUsabilityValCUpgradeRollbackAdvisory()
	model.ApprovalImplied = true
	if got := EvaluateProductionUsabilityValCUpgradeAdvisoryState(model); got == ProductionUsabilityValCUpgradeAdvisoryStateActive {
		t.Fatalf("expected non-active advisory state when preview implies approval, got %q", got)
	}
}

func TestProductionUsabilityValCPermissionAwareSupportFlowsDoNotMutateCanonicalState(t *testing.T) {
	model := ProductionUsabilityValCPermissionSupportFlows()
	model.Items[0].MutatesCanonicalState = true
	if got := EvaluateProductionUsabilityValCPermissionSupportState(model); got == ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected non-active permission support state when support flow mutates canonical state, got %q", got)
	}
}

func TestProductionUsabilityValCPartnerPublicSupportFlowsDoNotExposeRawInternalEvidence(t *testing.T) {
	model := ProductionUsabilityValCPermissionSupportFlows()
	for idx := range model.Items {
		if model.Items[idx].VisibilityScope == ProductionUsabilityVisibilityPartner {
			model.Items[idx].EvidenceVisibility = ProductionUsabilityEvidenceFull
		}
	}
	if got := EvaluateProductionUsabilityValCPermissionSupportState(model); got == ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected non-active permission support state when partner flow exposes full evidence, got %q", got)
	}
}

func TestProductionUsabilityValCPermissionSupportFailsClosedWhenPartnerScopeIsMissing(t *testing.T) {
	model := ProductionUsabilityValCPermissionSupportFlows()
	filtered := model.Items[:0]
	for _, item := range model.Items {
		if item.VisibilityScope == ProductionUsabilityVisibilityPartner {
			continue
		}
		filtered = append(filtered, item)
	}
	model.Items = filtered
	if got := EvaluateProductionUsabilityValCPermissionSupportState(model); got == ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected non-active permission support state when partner scope is missing, got %q", got)
	}
}

func TestProductionUsabilityValCPermissionSupportFailsClosedWhenPublicSafeScopeIsMissing(t *testing.T) {
	model := ProductionUsabilityValCPermissionSupportFlows()
	filtered := model.Items[:0]
	for _, item := range model.Items {
		if item.VisibilityScope == ProductionUsabilityVisibilityPublicSafe {
			continue
		}
		filtered = append(filtered, item)
	}
	model.Items = filtered
	if got := EvaluateProductionUsabilityValCPermissionSupportState(model); got == ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected non-active permission support state when public_safe scope is missing, got %q", got)
	}
}

func TestProductionUsabilityValCPermissionSupportFailsClosedWhenAnyRequiredScopeIsMissing(t *testing.T) {
	model := ProductionUsabilityValCPermissionSupportFlows()
	model.Items = model.Items[1:]
	if got := EvaluateProductionUsabilityValCPermissionSupportState(model); got == ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected non-active permission support state when a required scope is missing, got %q", got)
	}
}

func TestProductionUsabilityValCPermissionSupportPassesWhenAllRequiredScopesArePresent(t *testing.T) {
	model := ProductionUsabilityValCPermissionSupportFlows()
	if got := EvaluateProductionUsabilityValCPermissionSupportState(model); got != ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected active permission support state when all required scopes are present, got %q", got)
	}
}

func TestProductionUsabilityValCHiddenSupportSectionsAreRepresented(t *testing.T) {
	model := ProductionUsabilityValCPermissionSupportFlows()
	for idx := range model.Items {
		if model.Items[idx].VisibilityScope == ProductionUsabilityVisibilityPublicSafe {
			model.Items[idx].HiddenSections = nil
		}
	}
	if got := EvaluateProductionUsabilityValCPermissionSupportState(model); got == ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected non-active permission support state when hidden sections are omitted, got %q", got)
	}
}

func TestProductionUsabilityValCExportFailsClosedOnRawSecretDetection(t *testing.T) {
	model := ProductionUsabilityValCRedactionSafeExport()
	model.RawSecretDetected = true
	if got := EvaluateProductionUsabilityValCExportSafetyState(model); got == ProductionUsabilityValCExportSafetyStateActive {
		t.Fatalf("expected non-active export safety state on raw secret detection, got %q", got)
	}
}

func TestProductionUsabilityValCExportFailsClosedWhenPolicyDisallowsExport(t *testing.T) {
	model := ProductionUsabilityValCRedactionSafeExport()
	model.PolicyAllowsExport = false
	if got := EvaluateProductionUsabilityValCExportSafetyState(model); got == ProductionUsabilityValCExportSafetyStateActive {
		t.Fatalf("expected non-active export safety state when policy disallows export, got %q", got)
	}
}

func TestProductionUsabilityValCPublicPartnerSafeExportDoesNotIncludeRawInternalEvidence(t *testing.T) {
	model := ProductionUsabilityValCRedactionSafeExport()
	model.PublicSafe = true
	model.EvidenceHandling = ProductionUsabilityEvidenceFull
	if got := EvaluateProductionUsabilityValCExportSafetyState(model); got == ProductionUsabilityValCExportSafetyStateActive {
		t.Fatalf("expected non-active export safety state when public-safe export includes full evidence, got %q", got)
	}
}

func TestProductionUsabilityValCAuditorSafeDoesNotImplyPublicSafe(t *testing.T) {
	model := ProductionUsabilityValCRedactionSafeExport()
	model.PublicSafe = true
	if got := EvaluateProductionUsabilityValCExportSafetyState(model); got == ProductionUsabilityValCExportSafetyStateActive {
		t.Fatalf("expected non-active export safety state when auditor-safe export is treated as public-safe, got %q", got)
	}
}

func TestProductionUsabilityValCProofsCanBecomeActiveOnlyAsSupportabilityWhilePoint4RemainsNotComplete(t *testing.T) {
	got := EvaluateProductionUsabilityValCProofsState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCReadinessStateActive,
		ProductionUsabilityValCGuidedReadinessStateActive,
		ProductionUsabilityValCSupportBundleStateActive,
		ProductionUsabilityValCDiagnosticsStateActive,
		ProductionUsabilityValCHealthSnapshotStateActive,
		ProductionUsabilityValCRecoveryPlaybookStateActive,
		ProductionUsabilityValCUpgradeAdvisoryStateActive,
		ProductionUsabilityValCPermissionSupportStateActive,
		ProductionUsabilityValCExportSafetyStateActive,
		[]string{
			"/v1/production/usability-operability-recovery/val0/proofs",
			"/v1/production/usability-operability-recovery/vala/proofs",
			"/v1/production/usability-operability-recovery/valb/proofs",
			"/v1/production/usability-operability-recovery/valc/readiness",
			"/v1/production/usability-operability-recovery/valc/guided-readiness",
			"/v1/production/usability-operability-recovery/valc/support-bundle",
			"/v1/production/usability-operability-recovery/valc/diagnostics",
			"/v1/production/usability-operability-recovery/valc/health-snapshot",
			"/v1/production/usability-operability-recovery/valc/recovery-playbooks",
			"/v1/production/usability-operability-recovery/valc/upgrade-rollback-advisory",
			"/v1/production/usability-operability-recovery/valc/permission-support-flows",
			"/v1/production/usability-operability-recovery/valc/redaction-export-safety",
			"/v1/production/usability-operability-recovery/valc/proofs",
		},
		[]string{"val0_proofs", "vala_proofs", "valb_proofs", "readiness", "guided", "bundle", "diagnostics", "health", "recovery", "advisory"},
		[]string{"Val C proves bounded supportability only."},
		[]string{"Point 4 full PASS remains deferred."},
	)
	if got != ProductionUsabilityValCStateActive {
		t.Fatalf("expected active Val C proofs state for complete supportability slice, got %q", got)
	}
	if ProductionUsabilityPoint4StateNotComplete == ProductionUsabilityValCStateActive {
		t.Fatalf("point 4 state must remain distinct from active Val C state")
	}
}

func TestProductionUsabilityValCMissingRequiredComponentKeepsValCInactive(t *testing.T) {
	got := EvaluateProductionUsabilityValCState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCReadinessStateActive,
		ProductionUsabilityValCGuidedReadinessStateActive,
		ProductionUsabilityValCSupportBundleStateIncomplete,
		ProductionUsabilityValCDiagnosticsStateActive,
		ProductionUsabilityValCHealthSnapshotStateActive,
		ProductionUsabilityValCRecoveryPlaybookStateActive,
		ProductionUsabilityValCUpgradeAdvisoryStateActive,
		ProductionUsabilityValCPermissionSupportStateActive,
		ProductionUsabilityValCExportSafetyStateActive,
	)
	if got == ProductionUsabilityValCStateActive {
		t.Fatalf("expected non-active Val C state with missing required component, got %q", got)
	}
}

func TestProductionUsabilityValCFoundationIsActive(t *testing.T) {
	readiness := ProductionUsabilityValCReadinessChecks()
	if got := EvaluateProductionUsabilityValCReadinessState(readiness); got != ProductionUsabilityValCReadinessStateActive {
		t.Fatalf("expected active readiness state, got %q", got)
	}

	guided := ProductionUsabilityValCGuidedReadiness()
	if got := EvaluateProductionUsabilityValCGuidedReadinessState(guided); got != ProductionUsabilityValCGuidedReadinessStateActive {
		t.Fatalf("expected active guided readiness state, got %q", got)
	}

	supportBundle := ProductionUsabilityValCSupportBundleQualityGate()
	if got := EvaluateProductionUsabilityValCSupportBundleState(supportBundle); got != ProductionUsabilityValCSupportBundleStateActive {
		t.Fatalf("expected active support bundle state, got %q", got)
	}

	diagnostics := ProductionUsabilityValCDiagnosticsHardening()
	if got := EvaluateProductionUsabilityValCDiagnosticsState(diagnostics); got != ProductionUsabilityValCDiagnosticsStateActive {
		t.Fatalf("expected active diagnostics state, got %q", got)
	}

	health := ProductionUsabilityValCHealthSnapshot()
	if got := EvaluateProductionUsabilityValCHealthSnapshotState(health); got != ProductionUsabilityValCHealthSnapshotStateActive {
		t.Fatalf("expected active health snapshot state, got %q", got)
	}

	recovery := ProductionUsabilityValCRecoveryPlaybooks()
	if got := EvaluateProductionUsabilityValCRecoveryPlaybookState(recovery); got != ProductionUsabilityValCRecoveryPlaybookStateActive {
		t.Fatalf("expected active recovery playbook state, got %q", got)
	}

	advisory := ProductionUsabilityValCUpgradeRollbackAdvisory()
	if got := EvaluateProductionUsabilityValCUpgradeAdvisoryState(advisory); got != ProductionUsabilityValCUpgradeAdvisoryStateActive {
		t.Fatalf("expected active advisory state, got %q", got)
	}

	permission := ProductionUsabilityValCPermissionSupportFlows()
	if got := EvaluateProductionUsabilityValCPermissionSupportState(permission); got != ProductionUsabilityValCPermissionSupportStateActive {
		t.Fatalf("expected active permission support state, got %q", got)
	}

	export := ProductionUsabilityValCRedactionSafeExport()
	if got := EvaluateProductionUsabilityValCExportSafetyState(export); got != ProductionUsabilityValCExportSafetyStateActive {
		t.Fatalf("expected active export safety state, got %q", got)
	}

	if got := EvaluateProductionUsabilityValCState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		EvaluateProductionUsabilityValCReadinessState(readiness),
		EvaluateProductionUsabilityValCGuidedReadinessState(guided),
		EvaluateProductionUsabilityValCSupportBundleState(supportBundle),
		EvaluateProductionUsabilityValCDiagnosticsState(diagnostics),
		EvaluateProductionUsabilityValCHealthSnapshotState(health),
		EvaluateProductionUsabilityValCRecoveryPlaybookState(recovery),
		EvaluateProductionUsabilityValCUpgradeAdvisoryState(advisory),
		EvaluateProductionUsabilityValCPermissionSupportState(permission),
		EvaluateProductionUsabilityValCExportSafetyState(export),
	); got != ProductionUsabilityValCStateActive {
		t.Fatalf("expected active overall Val C state, got %q", got)
	}
}
