package operability

import "testing"

func activeReferenceArchitectureValBPrereqs() (string, string, string, string, string, string) {
	return IntelligenceCalibrationPoint5StatePass,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitectureVal0StateActive,
		ReferenceArchitectureValAStateActive,
		ReferenceArchitectureValAStateActive,
		ReferenceArchitecturePoint6StateNotComplete
}

func activeReferenceArchitectureValBComponents() (
	ReferenceArchitectureBlueprintPackRegistry,
	ReferenceArchitectureArtifactManifestCollection,
	ReferenceArchitectureBundleCollection,
	ReferenceArchitectureReadinessCollection,
	ReferenceArchitectureValidationHookCollection,
	ReferenceArchitectureDeviationCollection,
	ReferenceArchitectureConformanceKitCollection,
) {
	return ReferenceArchitectureValBPackRegistry(),
		ReferenceArchitectureValBArtifactManifestCollection(),
		ReferenceArchitectureValBBundleCollection(),
		ReferenceArchitectureValBReadinessCollection(),
		ReferenceArchitectureValBValidationHookCollection(),
		ReferenceArchitectureValBDeviationCollection(),
		ReferenceArchitectureValBConformanceKitCollection()
}

func TestReferenceArchitectureValBDependencyGates(t *testing.T) {
	point5State, val0CurrentState, val0State, valACurrentState, valAState, point6State := activeReferenceArchitectureValBPrereqs()
	registry, manifests, bundles, readiness, hooks, deviations, kits := activeReferenceArchitectureValBComponents()
	packState := EvaluateReferenceArchitectureValBPackRegistryState(registry)
	manifestState := EvaluateReferenceArchitectureValBArtifactManifestCollectionState(manifests)
	bundleState := EvaluateReferenceArchitectureValBBundleCollectionState(bundles)
	readinessState := EvaluateReferenceArchitectureValBReadinessCollectionState(readiness)
	hookState := EvaluateReferenceArchitectureValBValidationHookCollectionState(hooks)
	deviationState := EvaluateReferenceArchitectureValBDeviationCollectionState(deviations)
	conformanceState := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readiness, hooks, deviations)
	if got := EvaluateReferenceArchitectureValBState(point5State, val0CurrentState, val0State, valACurrentState, valAState, point6State, packState, manifestState, bundleState, readinessState, hookState, conformanceState, deviationState); got != ReferenceArchitectureValBStateActive {
		t.Fatalf("expected active Val B state with valid dependencies and components, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValBState(point5State, ReferenceArchitectureVal0StateIncomplete, val0State, valACurrentState, valAState, point6State, packState, manifestState, bundleState, readinessState, hookState, conformanceState, deviationState); got != ReferenceArchitectureValBStateBlocked {
		t.Fatalf("expected blocked Val B state when Val 0 dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValBState(point5State, val0CurrentState, val0State, ReferenceArchitectureValAStatePartial, valAState, point6State, packState, manifestState, bundleState, readinessState, hookState, conformanceState, deviationState); got != ReferenceArchitectureValBStateBlocked {
		t.Fatalf("expected blocked Val B state when Val A dependency is missing, got %q", got)
	}
	if got := EvaluateReferenceArchitectureValBState(point5State, val0CurrentState, val0State, valACurrentState, valAState, ReferenceArchitecturePoint6StatePass, packState, manifestState, bundleState, readinessState, hookState, conformanceState, deviationState); got != ReferenceArchitectureValBStateBlocked {
		t.Fatalf("expected blocked Val B state when point 6 is not not_complete, got %q", got)
	}
}

func TestReferenceArchitectureValBPackValidation(t *testing.T) {
	registry := ReferenceArchitectureValBPackRegistry()
	pack := registry.Packs[0]
	if got := EvaluateReferenceArchitectureValBPackState(pack); got != ReferenceArchitectureValBPackStateActive {
		t.Fatalf("expected active pack state for valid pack, got %q", got)
	}
	pack = registry.Packs[0]
	pack.PackID = ""
	if got := EvaluateReferenceArchitectureValBPackState(pack); got == ReferenceArchitectureValBPackStateActive {
		t.Fatalf("expected non-active pack state for missing pack_id, got %q", got)
	}
	pack = registry.Packs[0]
	pack.BlueprintFamily = "enterprise-defualt"
	if got := EvaluateReferenceArchitectureValBPackState(pack); got == ReferenceArchitectureValBPackStateActive {
		t.Fatalf("expected non-active pack state for unknown family, got %q", got)
	}
	pack = registry.Packs[0]
	pack.LifecycleState = "acitve"
	if got := EvaluateReferenceArchitectureValBPackState(pack); got == ReferenceArchitectureValBPackStateActive {
		t.Fatalf("expected non-active pack state for typo lifecycle, got %q", got)
	}
	pack = registry.Packs[0]
	pack.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
	if got := EvaluateReferenceArchitectureValBPackState(pack); got == ReferenceArchitectureValBPackStateActive {
		t.Fatalf("expected non-active pack state for unsupported compatibility, got %q", got)
	}
	pack = registry.Packs[0]
	pack.LifecycleState = ReferenceArchitectureLifecycleDeprecated
	if got := EvaluateReferenceArchitectureValBPackState(pack); got == ReferenceArchitectureValBPackStateActive {
		t.Fatalf("expected non-active pack state for deprecated lifecycle, got %q", got)
	}
	pack = registry.Packs[0]
	pack.ProjectionDisclaimer = ""
	if got := EvaluateReferenceArchitectureValBPackState(pack); got == ReferenceArchitectureValBPackStateActive {
		t.Fatalf("expected non-active pack state without projection disclaimer, got %q", got)
	}
}

func TestReferenceArchitectureValBArtifactManifestValidation(t *testing.T) {
	collection := ReferenceArchitectureValBArtifactManifestCollection()
	manifest := collection.Manifests[0]
	if got := EvaluateReferenceArchitectureValBArtifactManifestState(manifest); got != ReferenceArchitectureValBManifestStateActive {
		t.Fatalf("expected active artifact manifest state, got %q", got)
	}
	collection = ReferenceArchitectureValBArtifactManifestCollection()
	manifest = collection.Manifests[0]
	manifest.Artifacts = manifest.Artifacts[1:]
	if got := EvaluateReferenceArchitectureValBArtifactManifestState(manifest); got == ReferenceArchitectureValBManifestStateActive {
		t.Fatalf("expected non-active manifest state when a required artifact is missing, got %q", got)
	}
	collection = ReferenceArchitectureValBArtifactManifestCollection()
	manifest = collection.Manifests[0]
	manifest.Artifacts = append(manifest.Artifacts[1:], manifest.Artifacts[0])
	manifest.Artifacts[0].ArtifactType = ReferenceArchitectureValBArtifactProfileBundle
	if got := EvaluateReferenceArchitectureValBArtifactManifestState(manifest); got == ReferenceArchitectureValBManifestStateActive {
		t.Fatalf("expected non-active manifest state when duplicate artifacts compensate for missing required type, got %q", got)
	}
	collection = ReferenceArchitectureValBArtifactManifestCollection()
	manifest = collection.Manifests[0]
	manifest.Artifacts[0].ArtifactType = "artifact_bundle"
	if got := EvaluateReferenceArchitectureValBArtifactManifestState(manifest); got == ReferenceArchitectureValBManifestStateActive {
		t.Fatalf("expected non-active manifest state for unknown artifact type, got %q", got)
	}
	collection = ReferenceArchitectureValBArtifactManifestCollection()
	manifest = collection.Manifests[0]
	manifest.Artifacts[0].Timestamp = "2026/04/26"
	if got := EvaluateReferenceArchitectureValBArtifactManifestState(manifest); got == ReferenceArchitectureValBManifestStateActive {
		t.Fatalf("expected non-active manifest state for malformed timestamp, got %q", got)
	}
	collection = ReferenceArchitectureValBArtifactManifestCollection()
	manifest = collection.Manifests[0]
	manifest.Artifacts[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValBArtifactManifestState(manifest); got == ReferenceArchitectureValBManifestStateActive {
		t.Fatalf("expected non-active manifest state for stale artifact freshness, got %q", got)
	}
	collection = ReferenceArchitectureValBArtifactManifestCollection()
	manifest = collection.Manifests[0]
	filtered := make([]ReferenceArchitectureArtifactEntry, 0, len(manifest.Artifacts))
	for _, artifact := range manifest.Artifacts {
		if artifact.RequirementLevel == ReferenceArchitectureValBArtifactOptional {
			continue
		}
		filtered = append(filtered, artifact)
	}
	manifest.Artifacts = filtered
	if got := EvaluateReferenceArchitectureValBArtifactManifestState(manifest); got != ReferenceArchitectureValBManifestStateActive {
		t.Fatalf("expected optional artifacts to be removable without blocking active state, got %q", got)
	}
}

func TestReferenceArchitectureValBReadinessValidation(t *testing.T) {
	collection := ReferenceArchitectureValBReadinessCollection()
	bundle := collection.Bundles[0]
	if got := EvaluateReferenceArchitectureValBReadinessBundleState(bundle); got != ReferenceArchitectureValBReadinessStateActive {
		t.Fatalf("expected active readiness bundle state, got %q", got)
	}
	bundle = collection.Bundles[0]
	bundle.Checks = bundle.Checks[1:]
	if got := EvaluateReferenceArchitectureValBReadinessBundleState(bundle); got == ReferenceArchitectureValBReadinessStateActive {
		t.Fatalf("expected non-active readiness state for missing required check, got %q", got)
	}
	bundle = collection.Bundles[0]
	bundle.Checks[0].State = "rdy"
	if got := EvaluateReferenceArchitectureValBReadinessBundleState(bundle); got == ReferenceArchitectureValBReadinessStateActive {
		t.Fatalf("expected non-active readiness state for unknown readiness state, got %q", got)
	}
	bundle = collection.Bundles[0]
	bundle.Checks[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValBReadinessBundleState(bundle); got == ReferenceArchitectureValBReadinessStateActive {
		t.Fatalf("expected non-active readiness state for stale evidence, got %q", got)
	}
	bundle = collection.Bundles[0]
	bundle.Checks[0].Category = "capacity_fit"
	if got := EvaluateReferenceArchitectureValBReadinessBundleState(bundle); got == ReferenceArchitectureValBReadinessStateActive {
		t.Fatalf("expected non-active readiness state for unsupported readiness category, got %q", got)
	}
	bundle = collection.Bundles[0]
	bundle.Checks[0].RedactionKeepsCaveats = false
	if got := EvaluateReferenceArchitectureValBReadinessBundleState(bundle); got == ReferenceArchitectureValBReadinessStateActive {
		t.Fatalf("expected non-active readiness state when redaction can hide caveats, got %q", got)
	}
}

func TestReferenceArchitectureValBValidationHookDescriptors(t *testing.T) {
	collection := ReferenceArchitectureValBValidationHookCollection()
	hookPack := collection.HookPacks[0]
	if got := EvaluateReferenceArchitectureValBHookPackState(hookPack); got != ReferenceArchitectureValBHookStateActive {
		t.Fatalf("expected active validation hook state, got %q", got)
	}
	hookPack = collection.HookPacks[0]
	hookPack.Hooks = hookPack.Hooks[1:]
	if got := EvaluateReferenceArchitectureValBHookPackState(hookPack); got == ReferenceArchitectureValBHookStateActive {
		t.Fatalf("expected non-active hook state for missing required hook, got %q", got)
	}
	hookPack = collection.HookPacks[0]
	hookPack.Hooks[0].Category = "schema_validation"
	if got := EvaluateReferenceArchitectureValBHookPackState(hookPack); got == ReferenceArchitectureValBHookStateActive {
		t.Fatalf("expected non-active hook state for unknown category, got %q", got)
	}
	hookPack = collection.HookPacks[0]
	hookPack.Hooks = append(hookPack.Hooks[1:], hookPack.Hooks[0])
	hookPack.Hooks[0].Category = ReferenceArchitectureValBHookConfigValidation
	if got := EvaluateReferenceArchitectureValBHookPackState(hookPack); got == ReferenceArchitectureValBHookStateActive {
		t.Fatalf("expected non-active hook state when duplicate hooks compensate for missing required category, got %q", got)
	}
	hookPack = collection.HookPacks[0]
	hookPack.Hooks[0].ExpectedInputRefs = nil
	if got := EvaluateReferenceArchitectureValBHookPackState(hookPack); got == ReferenceArchitectureValBHookStateActive {
		t.Fatalf("expected non-active hook state with missing expected refs, got %q", got)
	}
	hookPack = collection.HookPacks[0]
	hookPack.HookPackRef = "validation-pack/" + hookPack.BlueprintFamily
	if got := EvaluateReferenceArchitectureValBHookPackState(hookPack); got == ReferenceArchitectureValBHookStateActive {
		t.Fatalf("expected non-active hook state when hook pack ref does not use hook-pack identity, got %q", got)
	}
}

func TestReferenceArchitectureValBConformanceKitEvaluation(t *testing.T) {
	registry, manifests, bundles, readiness, hooks, deviations, kits := activeReferenceArchitectureValBComponents()
	pack := registry.Packs[0]
	manifest := manifests.Manifests[0]
	bundle := bundles.Bundles[0]
	readinessBundle := readiness.Bundles[0]
	hookPack := hooks.HookPacks[0]
	report := deviations.Reports[0]
	kit := kits.Kits[0]
	packState := EvaluateReferenceArchitectureValBPackState(pack)
	manifestState := EvaluateReferenceArchitectureValBArtifactManifestState(manifest)
	bundleState := EvaluateReferenceArchitectureValBBundleState(bundle)
	readinessState := EvaluateReferenceArchitectureValBReadinessBundleState(readinessBundle)
	hookState := EvaluateReferenceArchitectureValBHookPackState(hookPack)
	deviationState := EvaluateReferenceArchitectureValBDeviationReportState(report)
	if got := EvaluateReferenceArchitectureValBConformanceKitState(packState, manifestState, bundleState, readinessState, hookState, deviationState, kit, pack, report); got != ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected matched conformance kit state, got %q", got)
	}
	manifest = manifests.Manifests[0]
	manifest.Artifacts = manifest.Artifacts[1:]
	if got := EvaluateReferenceArchitectureValBConformanceKitState(packState, EvaluateReferenceArchitectureValBArtifactManifestState(manifest), bundleState, readinessState, hookState, deviationState, kit, pack, report); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected missing required artifact to block matched conformance state, got %q", got)
	}
	kit = kits.Kits[0]
	kit.EvidenceRefs[0].FreshnessState = IntelligenceCalibrationFreshnessStale
	if got := EvaluateReferenceArchitectureValBConformanceKitState(packState, manifestState, bundleState, readinessState, hookState, deviationState, kit, pack, report); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected stale evidence to block matched conformance state, got %q", got)
	}
	pack = registry.Packs[0]
	pack.CompatibilityState = ReferenceArchitectureCompatibilityUnsupported
	if got := EvaluateReferenceArchitectureValBConformanceKitState(EvaluateReferenceArchitectureValBPackState(pack), manifestState, bundleState, readinessState, hookState, deviationState, kits.Kits[0], pack, report); got != ReferenceArchitectureConformanceUnsupported {
		t.Fatalf("expected unsupported compatibility to return unsupported conformance state, got %q", got)
	}
	report = deviations.Reports[0]
	report.Deviations = []ReferenceArchitectureDeviation{
		{DeviationID: "dev-overclaim", Category: ReferenceArchitectureValBDeviationOverclaimLanguageDetected, Severity: ReferenceArchitectureValBSeverityHigh, AffectedScope: "pack", EvidenceRef: "pack-evidence", Explanation: "overclaim blocks matched", BlocksMatched: true},
	}
	if got := EvaluateReferenceArchitectureValBConformanceKitState(packState, manifestState, bundleState, readinessState, hookState, EvaluateReferenceArchitectureValBDeviationReportState(report), kits.Kits[0], registry.Packs[0], report); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected blocking deviation to prevent matched conformance state, got %q", got)
	}
	kit = kits.Kits[0]
	kit.ConformanceState = "matchd"
	if got := EvaluateReferenceArchitectureValBConformanceKitState(packState, manifestState, bundleState, readinessState, hookState, deviationState, kit, registry.Packs[0], deviations.Reports[0]); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected unknown conformance state to fail closed, got %q", got)
	}
	kit = kits.Kits[0]
	kit.ValidationHookPackRef = "validation-pack/" + kit.BlueprintFamily
	if got := EvaluateReferenceArchitectureValBConformanceKitState(packState, manifestState, bundleState, readinessState, hookState, deviationState, kit, registry.Packs[0], deviations.Reports[0]); got == ReferenceArchitectureConformanceMatched {
		t.Fatalf("expected invalid hook pack ref to block matched conformance state, got %q", got)
	}
}

func TestReferenceArchitectureValBConformanceKitCollectionRequiresActiveDependencies(t *testing.T) {
	registry, manifests, bundles, readiness, hooks, deviations, kits := activeReferenceArchitectureValBComponents()
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readiness, hooks, deviations); got != ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected active conformance kit collection state for valid fixtures, got %q", got)
	}

	manifestPartial := ReferenceArchitectureValBArtifactManifestCollection()
	manifestPartial.Manifests[0].Artifacts = manifestPartial.Manifests[0].Artifacts[1:]
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifestPartial, bundles, readiness, hooks, deviations); got == ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected non-active conformance kit collection when artifact manifest collection is partial, got %q", got)
	}

	bundlePartial := ReferenceArchitectureValBBundleCollection()
	bundlePartial.Bundles[0].EvidenceRequirements = nil
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundlePartial, readiness, hooks, deviations); got == ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected non-active conformance kit collection when bundle collection is partial, got %q", got)
	}

	readinessPartial := ReferenceArchitectureValBReadinessCollection()
	readinessPartial.Bundles[0].Checks[0].State = ReferenceArchitectureValBReadinessUnknown
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readinessPartial, hooks, deviations); got == ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected non-active conformance kit collection when readiness collection is partial, got %q", got)
	}

	hookPartial := ReferenceArchitectureValBValidationHookCollection()
	hookPartial.HookPacks[0].Hooks[0].Category = "schema_validation"
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readiness, hookPartial, deviations); got == ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected non-active conformance kit collection when hook collection is partial, got %q", got)
	}

	deviationPartial := ReferenceArchitectureValBDeviationCollection()
	deviationPartial.Reports[0].Deviations = []ReferenceArchitectureDeviation{
		{DeviationID: "dev-unknown", Category: "unknown_category", Severity: ReferenceArchitectureValBSeverityMedium, AffectedScope: "pack", Explanation: "unsupported"},
	}
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readiness, hooks, deviationPartial); got == ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected non-active conformance kit collection when deviation collection is partial, got %q", got)
	}

	manifestDuplicate := ReferenceArchitectureValBArtifactManifestCollection()
	manifestDuplicate.Manifests = append(manifestDuplicate.Manifests, manifestDuplicate.Manifests[0])
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifestDuplicate, bundles, readiness, hooks, deviations); got == ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected duplicate dependency records to remain visible and block active state, got %q", got)
	}

	bundleExtra := ReferenceArchitectureValBBundleCollection()
	extraBundle := bundleExtra.Bundles[0]
	extraBundle.BundleID = "bundle/extra"
	extraBundle.PackID = "pack/extra"
	extraBundle.BlueprintFamily = "extra_family"
	bundleExtra.Bundles = append(bundleExtra.Bundles, extraBundle)
	if got := EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundleExtra, readiness, hooks, deviations); got == ReferenceArchitectureValBConformanceKitStateActive {
		t.Fatalf("expected extra dependency records to remain visible and block active state, got %q", got)
	}
}

func TestReferenceArchitectureValBValidationHookPackRefUsesHookCollectionIdentity(t *testing.T) {
	registry := ReferenceArchitectureValBPackRegistry()
	hooks := ReferenceArchitectureValBValidationHookCollection()
	kits := ReferenceArchitectureValBConformanceKitCollection()

	pack := registry.Packs[0]
	hookPack := hooks.HookPacks[0]
	kit := kits.Kits[0]
	expectedRef := referenceArchitectureValBHookPackRefForPack(pack)

	if hookPack.HookPackRef != expectedRef {
		t.Fatalf("expected hook pack ref %q, got %q", expectedRef, hookPack.HookPackRef)
	}
	if kit.ValidationHookPackRef != expectedRef {
		t.Fatalf("expected conformance kit hook pack ref %q, got %q", expectedRef, kit.ValidationHookPackRef)
	}

	found := false
	for _, candidate := range hooks.HookPacks {
		if candidate.HookPackRef == kit.ValidationHookPackRef {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected ValidationHookPackRef %q to dereference against the validation hook collection", kit.ValidationHookPackRef)
	}
}

func TestReferenceArchitectureValBDeviationClassifierValidation(t *testing.T) {
	report := ReferenceArchitectureValBDeviationCollection().Reports[0]
	for _, category := range referenceArchitectureValBDeviationCategories() {
		candidate := report
		candidate.Deviations = []ReferenceArchitectureDeviation{
			{
				DeviationID:   "deviation-" + category,
				Category:      category,
				Severity:      ReferenceArchitectureValBSeverityMedium,
				AffectedScope: "pack",
				EvidenceRef:   "evidence-ref",
				Explanation:   "supported deviation category",
				BlocksMatched: category == ReferenceArchitectureValBDeviationMissingSupportBoundary || category == ReferenceArchitectureValBDeviationOverclaimLanguageDetected,
				AdvisoryOnly:  category != ReferenceArchitectureValBDeviationMissingSupportBoundary,
			},
		}
		if got := EvaluateReferenceArchitectureValBDeviationReportState(candidate); got != ReferenceArchitectureValBDeviationStateActive {
			t.Fatalf("expected supported deviation category %s to be accepted, got %q", category, got)
		}
	}
	report = ReferenceArchitectureValBDeviationCollection().Reports[0]
	report.Deviations = []ReferenceArchitectureDeviation{
		{DeviationID: "dev-unknown", Category: "unknown_category", Severity: ReferenceArchitectureValBSeverityMedium, AffectedScope: "pack", Explanation: "unsupported"},
	}
	if got := EvaluateReferenceArchitectureValBDeviationReportState(report); got == ReferenceArchitectureValBDeviationStateActive {
		t.Fatalf("expected unknown deviation category to fail closed, got %q", got)
	}
	report = ReferenceArchitectureValBDeviationCollection().Reports[0]
	report.Deviations = []ReferenceArchitectureDeviation{
		{DeviationID: "dev-unknown-severity", Category: ReferenceArchitectureValBDeviationOverclaimLanguageDetected, Severity: "severe", AffectedScope: "pack", Explanation: "unsupported"},
	}
	if got := EvaluateReferenceArchitectureValBDeviationReportState(report); got == ReferenceArchitectureValBDeviationStateActive {
		t.Fatalf("expected unknown severity to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValBNoOverclaimAndPoint6NotComplete(t *testing.T) {
	point5State, val0CurrentState, val0State, valACurrentState, valAState, point6State := activeReferenceArchitectureValBPrereqs()
	registry, manifests, bundles, readiness, hooks, deviations, kits := activeReferenceArchitectureValBComponents()
	valBState := EvaluateReferenceArchitectureValBState(
		point5State,
		val0CurrentState,
		val0State,
		valACurrentState,
		valAState,
		point6State,
		EvaluateReferenceArchitectureValBPackRegistryState(registry),
		EvaluateReferenceArchitectureValBArtifactManifestCollectionState(manifests),
		EvaluateReferenceArchitectureValBBundleCollectionState(bundles),
		EvaluateReferenceArchitectureValBReadinessCollectionState(readiness),
		EvaluateReferenceArchitectureValBValidationHookCollectionState(hooks),
		EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readiness, hooks, deviations),
		EvaluateReferenceArchitectureValBDeviationCollectionState(deviations),
	)
	if valBState != ReferenceArchitectureValBStateActive {
		t.Fatalf("expected active Val B state with valid fixtures, got %q", valBState)
	}
	if got := EvaluateReferenceArchitectureValBProofsState(
		valBState,
		ReferenceArchitecturePoint6StatePass,
		referenceArchitectureVal0Families(),
		referenceArchitectureValBProofSurfaceRefs(),
		[]string{"point5_integrated_closure", "point6_val0_proofs", "point6_vala_proofs", "registry", "kit", "p1", "p2", "p3", "p4", "p5"},
		[]string{"Val B keeps point 6 not complete."},
		referenceArchitectureValBProjectionDisclaimer(),
	); got == ReferenceArchitectureValBStateActive {
		t.Fatalf("expected non-active proofs state when point 6 pass is claimed in Val B, got %q", got)
	}
}
