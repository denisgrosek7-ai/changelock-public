package operability

import "strings"

const (
	ReferenceArchitectureValBPackStateActive     = "reference_architecture_valb_pack_active"
	ReferenceArchitectureValBPackStatePartial    = "reference_architecture_valb_pack_partial"
	ReferenceArchitectureValBPackStateIncomplete = "reference_architecture_valb_pack_incomplete"
	ReferenceArchitectureValBPackStateBlocked    = "reference_architecture_valb_pack_blocked"
	ReferenceArchitectureValBPackStateUnknown    = "reference_architecture_valb_pack_unknown"

	ReferenceArchitectureValBManifestStateActive     = "reference_architecture_valb_manifest_active"
	ReferenceArchitectureValBManifestStatePartial    = "reference_architecture_valb_manifest_partial"
	ReferenceArchitectureValBManifestStateIncomplete = "reference_architecture_valb_manifest_incomplete"
	ReferenceArchitectureValBManifestStateBlocked    = "reference_architecture_valb_manifest_blocked"
	ReferenceArchitectureValBManifestStateUnknown    = "reference_architecture_valb_manifest_unknown"

	ReferenceArchitectureValBBundleStateActive     = "reference_architecture_valb_bundle_active"
	ReferenceArchitectureValBBundleStatePartial    = "reference_architecture_valb_bundle_partial"
	ReferenceArchitectureValBBundleStateIncomplete = "reference_architecture_valb_bundle_incomplete"
	ReferenceArchitectureValBBundleStateBlocked    = "reference_architecture_valb_bundle_blocked"
	ReferenceArchitectureValBBundleStateUnknown    = "reference_architecture_valb_bundle_unknown"

	ReferenceArchitectureValBReadinessStateActive     = "reference_architecture_valb_readiness_active"
	ReferenceArchitectureValBReadinessStatePartial    = "reference_architecture_valb_readiness_partial"
	ReferenceArchitectureValBReadinessStateIncomplete = "reference_architecture_valb_readiness_incomplete"
	ReferenceArchitectureValBReadinessStateBlocked    = "reference_architecture_valb_readiness_blocked"
	ReferenceArchitectureValBReadinessStateUnknown    = "reference_architecture_valb_readiness_unknown"

	ReferenceArchitectureValBHookStateActive     = "reference_architecture_valb_hook_active"
	ReferenceArchitectureValBHookStatePartial    = "reference_architecture_valb_hook_partial"
	ReferenceArchitectureValBHookStateIncomplete = "reference_architecture_valb_hook_incomplete"
	ReferenceArchitectureValBHookStateBlocked    = "reference_architecture_valb_hook_blocked"
	ReferenceArchitectureValBHookStateUnknown    = "reference_architecture_valb_hook_unknown"

	ReferenceArchitectureValBDeviationStateActive     = "reference_architecture_valb_deviation_active"
	ReferenceArchitectureValBDeviationStatePartial    = "reference_architecture_valb_deviation_partial"
	ReferenceArchitectureValBDeviationStateIncomplete = "reference_architecture_valb_deviation_incomplete"
	ReferenceArchitectureValBDeviationStateBlocked    = "reference_architecture_valb_deviation_blocked"
	ReferenceArchitectureValBDeviationStateUnknown    = "reference_architecture_valb_deviation_unknown"

	ReferenceArchitectureValBConformanceKitStateActive     = "reference_architecture_valb_conformance_kit_active"
	ReferenceArchitectureValBConformanceKitStatePartial    = "reference_architecture_valb_conformance_kit_partial"
	ReferenceArchitectureValBConformanceKitStateIncomplete = "reference_architecture_valb_conformance_kit_incomplete"
	ReferenceArchitectureValBConformanceKitStateBlocked    = "reference_architecture_valb_conformance_kit_blocked"
	ReferenceArchitectureValBConformanceKitStateUnknown    = "reference_architecture_valb_conformance_kit_unknown"

	ReferenceArchitectureValBStateActive     = "reference_architecture_valb_active"
	ReferenceArchitectureValBStatePartial    = "reference_architecture_valb_partial"
	ReferenceArchitectureValBStateIncomplete = "reference_architecture_valb_incomplete"
	ReferenceArchitectureValBStateBlocked    = "reference_architecture_valb_blocked"
	ReferenceArchitectureValBStateUnknown    = "reference_architecture_valb_unknown"

	ReferenceArchitectureValBArtifactProfileBundle   = "profile_bundle"
	ReferenceArchitectureValBArtifactConfigBundle    = "config_bundle"
	ReferenceArchitectureValBArtifactPolicyBundle    = "policy_bundle"
	ReferenceArchitectureValBArtifactReadinessBundle = "readiness_bundle"
	ReferenceArchitectureValBArtifactValidationPack  = "validation_pack"
	ReferenceArchitectureValBArtifactConformanceKit  = "conformance_kit"
	ReferenceArchitectureValBArtifactSupportDoc      = "support_boundary_doc"
	ReferenceArchitectureValBArtifactUpgradeNote     = "upgrade_note"
	ReferenceArchitectureValBArtifactCaveatNote      = "caveat_note"

	ReferenceArchitectureValBArtifactRequired = "required"
	ReferenceArchitectureValBArtifactOptional = "optional"

	ReferenceArchitectureValBReadinessCapabilityFit      = "capability_fit"
	ReferenceArchitectureValBReadinessEnvironmentFit     = "environment_fit"
	ReferenceArchitectureValBReadinessTrustAnchorFit     = "trust_anchor_fit"
	ReferenceArchitectureValBReadinessAuditPathFit       = "audit_path_fit"
	ReferenceArchitectureValBReadinessEvidenceStorageFit = "evidence_storage_fit"
	ReferenceArchitectureValBReadinessPolicyDistFit      = "policy_distribution_fit"
	ReferenceArchitectureValBReadinessRecoveryFit        = "recovery_fit"
	ReferenceArchitectureValBReadinessSupportBoundaryFit = "support_boundary_fit"
	ReferenceArchitectureValBReadinessCompatibilityFit   = "compatibility_fit"
	ReferenceArchitectureValBReadinessFreshnessFit       = "freshness_fit"

	ReferenceArchitectureValBReadinessReady          = "ready"
	ReferenceArchitectureValBReadinessPartiallyReady = "partially_ready"
	ReferenceArchitectureValBReadinessDegraded       = "degraded"
	ReferenceArchitectureValBReadinessUnsupported    = "unsupported"
	ReferenceArchitectureValBReadinessStale          = "stale"
	ReferenceArchitectureValBReadinessUnknown        = "unknown"

	ReferenceArchitectureValBHookConfigValidation            = "config_validation"
	ReferenceArchitectureValBHookReadinessValidation         = "readiness_validation"
	ReferenceArchitectureValBHookPolicyValidation            = "policy_validation"
	ReferenceArchitectureValBHookTrustAnchorValidation       = "trust_anchor_validation"
	ReferenceArchitectureValBHookAuditPathValidation         = "audit_path_validation"
	ReferenceArchitectureValBHookEvidenceFreshnessValidation = "evidence_freshness_validation"
	ReferenceArchitectureValBHookCompatibilityValidation     = "compatibility_validation"
	ReferenceArchitectureValBHookConformanceValidation       = "conformance_validation"

	ReferenceArchitectureValBHookFailureFailClosed = "fail_closed"
	ReferenceArchitectureValBHookFailureDegrade    = "degrade_only"
	ReferenceArchitectureValBHookFailureDrift      = "drift_only"

	ReferenceArchitectureValBDeviationMissingRequiredArtifact    = "missing_required_artifact"
	ReferenceArchitectureValBDeviationUnsupportedArtifactType    = "unsupported_artifact_type"
	ReferenceArchitectureValBDeviationMissingRequiredCapability  = "missing_required_capability"
	ReferenceArchitectureValBDeviationEnvironmentMismatch        = "environment_mismatch"
	ReferenceArchitectureValBDeviationTrustAnchorMismatch        = "trust_anchor_mismatch"
	ReferenceArchitectureValBDeviationAuditPathMismatch          = "audit_path_mismatch"
	ReferenceArchitectureValBDeviationStaleEvidence              = "stale_evidence"
	ReferenceArchitectureValBDeviationMalformedEvidenceTimestamp = "malformed_evidence_timestamp"
	ReferenceArchitectureValBDeviationUnsupportedCompatibility   = "unsupported_compatibility"
	ReferenceArchitectureValBDeviationDeprecatedOrSupersededPack = "deprecated_or_superseded_pack"
	ReferenceArchitectureValBDeviationMissingSupportBoundary     = "missing_support_boundary"
	ReferenceArchitectureValBDeviationProjectionMissing          = "projection_disclaimer_missing"
	ReferenceArchitectureValBDeviationOverclaimLanguageDetected  = "overclaim_language_detected"

	ReferenceArchitectureValBSeverityCritical = "critical"
	ReferenceArchitectureValBSeverityHigh     = "high"
	ReferenceArchitectureValBSeverityMedium   = "medium"
	ReferenceArchitectureValBSeverityLow      = "low"
)

type ReferenceArchitectureBlueprintPack struct {
	CurrentState              string                                   `json:"current_state"`
	PackID                    string                                   `json:"pack_id"`
	Version                   string                                   `json:"version"`
	BlueprintFamily           string                                   `json:"blueprint_family"`
	BlueprintID               string                                   `json:"blueprint_id"`
	LifecycleState            string                                   `json:"lifecycle_state"`
	CompatibilityState        string                                   `json:"compatibility_state"`
	Owner                     string                                   `json:"owner"`
	TargetEnvironmentRef      string                                   `json:"target_environment_ref"`
	ProfileRef                string                                   `json:"profile_ref"`
	ConfigBundleRef           string                                   `json:"config_bundle_ref"`
	PolicyBundleRef           string                                   `json:"policy_bundle_ref"`
	ReadinessBundleRef        string                                   `json:"readiness_bundle_ref"`
	ValidationPackRef         string                                   `json:"validation_pack_ref"`
	ConformanceKitRef         string                                   `json:"conformance_kit_ref"`
	SupportBoundaryRef        string                                   `json:"support_boundary_ref"`
	EvidenceRefs              []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	Caveats                   []string                                 `json:"caveats,omitempty"`
	ProjectionDisclaimer      string                                   `json:"projection_disclaimer"`
	CreatedAt                 string                                   `json:"created_at"`
	UpdatedAt                 string                                   `json:"updated_at"`
	CertifiedLanguagePresent  bool                                     `json:"certified_language_present"`
	GuaranteedDeploymentClaim bool                                     `json:"guaranteed_deployment_claim_present"`
	ClaimsPoint6Pass          bool                                     `json:"claims_point_6_pass"`
}

type ReferenceArchitectureBlueprintPackRegistry struct {
	CurrentState         string                               `json:"current_state"`
	RegistryID           string                               `json:"registry_id"`
	Version              string                               `json:"version"`
	SupportedFamilies    []string                             `json:"supported_families,omitempty"`
	Packs                []ReferenceArchitectureBlueprintPack `json:"packs,omitempty"`
	ProjectionDisclaimer string                               `json:"projection_disclaimer"`
}

type ReferenceArchitectureArtifactEntry struct {
	ArtifactID         string   `json:"artifact_id"`
	ArtifactType       string   `json:"artifact_type"`
	Version            string   `json:"version"`
	Scope              string   `json:"scope"`
	SourceRef          string   `json:"source_ref"`
	IntegrityRef       string   `json:"integrity_ref"`
	Timestamp          string   `json:"timestamp"`
	FreshnessState     string   `json:"freshness_state"`
	RequirementLevel   string   `json:"required_or_optional"`
	CompatibilityState string   `json:"compatibility_state"`
	Caveats            []string `json:"caveats,omitempty"`
}

type ReferenceArchitectureArtifactManifest struct {
	CurrentState         string                               `json:"current_state"`
	ManifestID           string                               `json:"manifest_id"`
	PackID               string                               `json:"pack_id"`
	BlueprintFamily      string                               `json:"blueprint_family"`
	SupportedTypes       []string                             `json:"supported_artifact_types,omitempty"`
	Artifacts            []ReferenceArchitectureArtifactEntry `json:"artifacts,omitempty"`
	ProjectionDisclaimer string                               `json:"projection_disclaimer"`
}

type ReferenceArchitectureArtifactManifestCollection struct {
	CurrentState         string                                  `json:"current_state"`
	CollectionID         string                                  `json:"collection_id"`
	SupportedTypes       []string                                `json:"supported_artifact_types,omitempty"`
	Manifests            []ReferenceArchitectureArtifactManifest `json:"manifests,omitempty"`
	ProjectionDisclaimer string                                  `json:"projection_disclaimer"`
}

type ReferenceArchitectureConfigProfilePolicyBundle struct {
	CurrentState                string   `json:"current_state"`
	BundleID                    string   `json:"bundle_id"`
	PackID                      string   `json:"pack_id"`
	BlueprintFamily             string   `json:"blueprint_family"`
	RequiredConfigKeys          []string `json:"required_config_keys,omitempty"`
	RequiredConfigGroups        []string `json:"required_config_groups,omitempty"`
	ProfileAssumptions          []string `json:"profile_assumptions,omitempty"`
	PolicyBaselineRefs          []string `json:"policy_baseline_refs,omitempty"`
	EnvironmentPrerequisites    []string `json:"environment_prerequisites,omitempty"`
	ForbiddenUnsupportedOptions []string `json:"forbidden_unsupported_options,omitempty"`
	CompatibilityCaveats        []string `json:"compatibility_caveats,omitempty"`
	EvidenceRequirements        []string `json:"evidence_requirements,omitempty"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureBundleCollection struct {
	CurrentState         string                                           `json:"current_state"`
	CollectionID         string                                           `json:"collection_id"`
	Bundles              []ReferenceArchitectureConfigProfilePolicyBundle `json:"bundles,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type ReferenceArchitectureReadinessCheck struct {
	CheckID               string   `json:"check_id"`
	Category              string   `json:"category"`
	State                 string   `json:"state"`
	Timestamp             string   `json:"timestamp"`
	FreshnessState        string   `json:"freshness_state"`
	EvidenceRefs          []string `json:"evidence_refs,omitempty"`
	Caveats               []string `json:"caveats,omitempty"`
	RedactionKeepsCaveats bool     `json:"redaction_keeps_caveats"`
}

type ReferenceArchitectureReadinessBundle struct {
	CurrentState         string                                `json:"current_state"`
	BundleID             string                                `json:"bundle_id"`
	PackID               string                                `json:"pack_id"`
	BlueprintFamily      string                                `json:"blueprint_family"`
	SupportedCategories  []string                              `json:"supported_categories,omitempty"`
	SupportedStates      []string                              `json:"supported_states,omitempty"`
	Checks               []ReferenceArchitectureReadinessCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                `json:"projection_disclaimer"`
}

type ReferenceArchitectureReadinessCollection struct {
	CurrentState         string                                 `json:"current_state"`
	CollectionID         string                                 `json:"collection_id"`
	SupportedCategories  []string                               `json:"supported_categories,omitempty"`
	SupportedStates      []string                               `json:"supported_states,omitempty"`
	Bundles              []ReferenceArchitectureReadinessBundle `json:"bundles,omitempty"`
	ProjectionDisclaimer string                                 `json:"projection_disclaimer"`
}

type ReferenceArchitectureValidationHookDescriptor struct {
	HookID                string   `json:"hook_id"`
	Category              string   `json:"category"`
	ExpectedInputRefs     []string `json:"expected_input_refs,omitempty"`
	ExpectedOutputRefs    []string `json:"expected_output_refs,omitempty"`
	RequiredEvidenceTypes []string `json:"required_evidence_types,omitempty"`
	FailureSemantics      string   `json:"failure_semantics"`
	TimeoutExpectation    string   `json:"timeout_or_freshness_expectation"`
	ProjectionDisclaimer  string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureValidationHookPack struct {
	CurrentState         string                                          `json:"current_state"`
	HookPackRef          string                                          `json:"hook_pack_ref"`
	PackID               string                                          `json:"pack_id"`
	BlueprintFamily      string                                          `json:"blueprint_family"`
	SupportedCategories  []string                                        `json:"supported_categories,omitempty"`
	Hooks                []ReferenceArchitectureValidationHookDescriptor `json:"hooks,omitempty"`
	ProjectionDisclaimer string                                          `json:"projection_disclaimer"`
}

type ReferenceArchitectureValidationHookCollection struct {
	CurrentState         string                                    `json:"current_state"`
	CollectionID         string                                    `json:"collection_id"`
	SupportedCategories  []string                                  `json:"supported_categories,omitempty"`
	HookPacks            []ReferenceArchitectureValidationHookPack `json:"hook_packs,omitempty"`
	ProjectionDisclaimer string                                    `json:"projection_disclaimer"`
}

type ReferenceArchitectureDeviation struct {
	DeviationID   string `json:"deviation_id"`
	Category      string `json:"category"`
	Severity      string `json:"severity"`
	AffectedScope string `json:"affected_scope"`
	EvidenceRef   string `json:"evidence_ref"`
	Explanation   string `json:"explanation"`
	BlocksMatched bool   `json:"blocks_matched"`
	AdvisoryOnly  bool   `json:"advisory_only"`
}

type ReferenceArchitectureDeviationReport struct {
	CurrentState         string                           `json:"current_state"`
	ReportID             string                           `json:"report_id"`
	PackID               string                           `json:"pack_id"`
	BlueprintFamily      string                           `json:"blueprint_family"`
	SupportedCategories  []string                         `json:"supported_categories,omitempty"`
	SupportedSeverities  []string                         `json:"supported_severities,omitempty"`
	Deviations           []ReferenceArchitectureDeviation `json:"deviations,omitempty"`
	ProjectionDisclaimer string                           `json:"projection_disclaimer"`
}

type ReferenceArchitectureDeviationCollection struct {
	CurrentState         string                                 `json:"current_state"`
	CollectionID         string                                 `json:"collection_id"`
	SupportedCategories  []string                               `json:"supported_categories,omitempty"`
	SupportedSeverities  []string                               `json:"supported_severities,omitempty"`
	Reports              []ReferenceArchitectureDeviationReport `json:"reports,omitempty"`
	ProjectionDisclaimer string                                 `json:"projection_disclaimer"`
}

type ReferenceArchitectureConformanceKit struct {
	CurrentState          string                                   `json:"current_state"`
	KitID                 string                                   `json:"kit_id"`
	PackID                string                                   `json:"pack_id"`
	BlueprintFamily       string                                   `json:"blueprint_family"`
	ProfileRef            string                                   `json:"profile_ref"`
	ManifestRef           string                                   `json:"manifest_ref"`
	BundleRef             string                                   `json:"bundle_ref"`
	ReadinessBundleRef    string                                   `json:"readiness_bundle_ref"`
	ValidationHookPackRef string                                   `json:"validation_hook_pack_ref"`
	DeviationReportRef    string                                   `json:"deviation_report_ref"`
	EvidenceRefs          []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	ConformanceState      string                                   `json:"conformance_state"`
	Caveats               []string                                 `json:"caveats,omitempty"`
	DegradedReasons       []string                                 `json:"degraded_reasons,omitempty"`
	UnsupportedReasons    []string                                 `json:"unsupported_reasons,omitempty"`
	ProjectionDisclaimer  string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureConformanceKitCollection struct {
	CurrentState               string                                `json:"current_state"`
	CollectionID               string                                `json:"collection_id"`
	SupportedConformanceStates []string                              `json:"supported_conformance_states,omitempty"`
	Kits                       []ReferenceArchitectureConformanceKit `json:"kits,omitempty"`
	ProjectionDisclaimer       string                                `json:"projection_disclaimer"`
}

func referenceArchitectureValBArtifactTypes() []string {
	return []string{
		ReferenceArchitectureValBArtifactProfileBundle,
		ReferenceArchitectureValBArtifactConfigBundle,
		ReferenceArchitectureValBArtifactPolicyBundle,
		ReferenceArchitectureValBArtifactReadinessBundle,
		ReferenceArchitectureValBArtifactValidationPack,
		ReferenceArchitectureValBArtifactConformanceKit,
		ReferenceArchitectureValBArtifactSupportDoc,
		ReferenceArchitectureValBArtifactUpgradeNote,
		ReferenceArchitectureValBArtifactCaveatNote,
	}
}

func referenceArchitectureValBRequiredArtifactTypes() []string {
	return []string{
		ReferenceArchitectureValBArtifactProfileBundle,
		ReferenceArchitectureValBArtifactConfigBundle,
		ReferenceArchitectureValBArtifactPolicyBundle,
		ReferenceArchitectureValBArtifactReadinessBundle,
		ReferenceArchitectureValBArtifactValidationPack,
		ReferenceArchitectureValBArtifactConformanceKit,
		ReferenceArchitectureValBArtifactSupportDoc,
	}
}

func referenceArchitectureValBArtifactRequirementLevels() []string {
	return []string{
		ReferenceArchitectureValBArtifactRequired,
		ReferenceArchitectureValBArtifactOptional,
	}
}

func referenceArchitectureValBReadinessCategories() []string {
	return []string{
		ReferenceArchitectureValBReadinessCapabilityFit,
		ReferenceArchitectureValBReadinessEnvironmentFit,
		ReferenceArchitectureValBReadinessTrustAnchorFit,
		ReferenceArchitectureValBReadinessAuditPathFit,
		ReferenceArchitectureValBReadinessEvidenceStorageFit,
		ReferenceArchitectureValBReadinessPolicyDistFit,
		ReferenceArchitectureValBReadinessRecoveryFit,
		ReferenceArchitectureValBReadinessSupportBoundaryFit,
		ReferenceArchitectureValBReadinessCompatibilityFit,
		ReferenceArchitectureValBReadinessFreshnessFit,
	}
}

func referenceArchitectureValBReadinessStates() []string {
	return []string{
		ReferenceArchitectureValBReadinessReady,
		ReferenceArchitectureValBReadinessPartiallyReady,
		ReferenceArchitectureValBReadinessDegraded,
		ReferenceArchitectureValBReadinessUnsupported,
		ReferenceArchitectureValBReadinessStale,
		ReferenceArchitectureValBReadinessUnknown,
	}
}

func referenceArchitectureValBHookCategories() []string {
	return []string{
		ReferenceArchitectureValBHookConfigValidation,
		ReferenceArchitectureValBHookReadinessValidation,
		ReferenceArchitectureValBHookPolicyValidation,
		ReferenceArchitectureValBHookTrustAnchorValidation,
		ReferenceArchitectureValBHookAuditPathValidation,
		ReferenceArchitectureValBHookEvidenceFreshnessValidation,
		ReferenceArchitectureValBHookCompatibilityValidation,
		ReferenceArchitectureValBHookConformanceValidation,
	}
}

func referenceArchitectureValBHookFailureSemantics() []string {
	return []string{
		ReferenceArchitectureValBHookFailureFailClosed,
		ReferenceArchitectureValBHookFailureDegrade,
		ReferenceArchitectureValBHookFailureDrift,
	}
}

func referenceArchitectureValBDeviationCategories() []string {
	return []string{
		ReferenceArchitectureValBDeviationMissingRequiredArtifact,
		ReferenceArchitectureValBDeviationUnsupportedArtifactType,
		ReferenceArchitectureValBDeviationMissingRequiredCapability,
		ReferenceArchitectureValBDeviationEnvironmentMismatch,
		ReferenceArchitectureValBDeviationTrustAnchorMismatch,
		ReferenceArchitectureValBDeviationAuditPathMismatch,
		ReferenceArchitectureValBDeviationStaleEvidence,
		ReferenceArchitectureValBDeviationMalformedEvidenceTimestamp,
		ReferenceArchitectureValBDeviationUnsupportedCompatibility,
		ReferenceArchitectureValBDeviationDeprecatedOrSupersededPack,
		ReferenceArchitectureValBDeviationMissingSupportBoundary,
		ReferenceArchitectureValBDeviationProjectionMissing,
		ReferenceArchitectureValBDeviationOverclaimLanguageDetected,
	}
}

func referenceArchitectureValBDeviationSeverities() []string {
	return []string{
		ReferenceArchitectureValBSeverityCritical,
		ReferenceArchitectureValBSeverityHigh,
		ReferenceArchitectureValBSeverityMedium,
		ReferenceArchitectureValBSeverityLow,
	}
}

func referenceArchitectureValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_blueprint_as_code_validation"
}

func referenceArchitectureValBHookPackRefForPack(pack ReferenceArchitectureBlueprintPack) string {
	return "hook-pack/" + strings.TrimSpace(pack.BlueprintFamily)
}

func referenceArchitectureValBPackCaveat(profile ReferenceArchitectureBlueprintFamilyProfile) string {
	switch strings.TrimSpace(profile.Family) {
	case ReferenceArchitectureFamilyHighAssurance:
		return "bounded high-assurance delivery pack with stricter trust and recovery assumptions"
	case ReferenceArchitectureFamilyRegulatedPrivacyFirst:
		return "bounded privacy-first delivery pack with evidence export and residency caveats"
	case ReferenceArchitectureFamilySovereignAirGapped:
		return "bounded sovereign air-gapped delivery pack without offline tooling execution"
	case ReferenceArchitectureFamilyPerformanceSensitive:
		return "bounded performance-sensitive delivery pack without runtime performance guarantees"
	case ReferenceArchitectureFamilyPartnerMSPSuitable:
		return "bounded partner or MSP delivery pack without partner canonical truth authority"
	default:
		return "bounded enterprise delivery pack with measured conformance only"
	}
}

func referenceArchitectureValBPackEvidenceRefs(profile ReferenceArchitectureBlueprintFamilyProfile) []ReferenceArchitectureEvidenceReference {
	refs := append([]ReferenceArchitectureEvidenceReference{}, referenceArchitectureValAProfileEvidenceRefs(profile)...)
	refs = append(refs, ReferenceArchitectureEvidenceReference{
		EvidenceID:     "evidence:" + strings.ReplaceAll(profile.Family, "_", "-") + "-valb-pack",
		EvidenceType:   ReferenceArchitectureEvidenceCompatibilityReport,
		Source:         "reference-architecture/valb/pack-registry",
		Timestamp:      profile.UpdatedAt,
		FreshnessState: IntelligenceCalibrationFreshnessFresh,
		Scope:          "blueprint_pack/" + profile.Family,
		Caveats:        []string{referenceArchitectureValBPackCaveat(profile)},
	})
	return refs
}

func referenceArchitectureValBPackFromProfile(profile ReferenceArchitectureBlueprintFamilyProfile) ReferenceArchitectureBlueprintPack {
	return ReferenceArchitectureBlueprintPack{
		CurrentState:         "reference_architecture_valb_pack_ready",
		PackID:               "blueprint-pack-" + strings.ReplaceAll(profile.Family, "_", "-") + "-001",
		Version:              "1.0.0",
		BlueprintFamily:      profile.Family,
		BlueprintID:          profile.BlueprintID,
		LifecycleState:       profile.LifecycleState,
		CompatibilityState:   profile.CompatibilityState,
		Owner:                "reference_architecture_program",
		TargetEnvironmentRef: "environment-ref/" + profile.Family,
		ProfileRef:           profile.BlueprintID,
		ConfigBundleRef:      "config-bundle/" + profile.Family,
		PolicyBundleRef:      "policy-bundle/" + profile.Family,
		ReadinessBundleRef:   "readiness-bundle/" + profile.Family,
		ValidationPackRef:    "validation-pack/" + profile.Family,
		ConformanceKitRef:    "conformance-kit/" + profile.Family,
		SupportBoundaryRef:   profile.SupportBoundaryRef,
		EvidenceRefs:         referenceArchitectureValBPackEvidenceRefs(profile),
		Caveats:              []string{referenceArchitectureValBPackCaveat(profile)},
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
		CreatedAt:            profile.CreatedAt,
		UpdatedAt:            profile.UpdatedAt,
	}
}

func ReferenceArchitectureValBPackRegistry() ReferenceArchitectureBlueprintPackRegistry {
	packs := make([]ReferenceArchitectureBlueprintPack, 0, len(ReferenceArchitectureValAFamilyProfiles()))
	for _, profile := range ReferenceArchitectureValAFamilyProfiles() {
		packs = append(packs, referenceArchitectureValBPackFromProfile(profile))
	}
	return ReferenceArchitectureBlueprintPackRegistry{
		CurrentState:         "reference_architecture_valb_pack_registry_ready",
		RegistryID:           "reference-architecture-valb-pack-registry",
		Version:              "1.0.0",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Packs:                packs,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func referenceArchitectureValBArtifactManifestForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureArtifactManifest {
	familyScope := strings.TrimSpace(pack.BlueprintFamily)
	artifacts := []ReferenceArchitectureArtifactEntry{
		{ArtifactID: pack.PackID + ":profile", ArtifactType: ReferenceArchitectureValBArtifactProfileBundle, Version: pack.Version, Scope: familyScope, SourceRef: pack.ProfileRef, IntegrityRef: "integrity/" + familyScope + "/profile", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactRequired, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":config", ArtifactType: ReferenceArchitectureValBArtifactConfigBundle, Version: pack.Version, Scope: familyScope, SourceRef: pack.ConfigBundleRef, IntegrityRef: "integrity/" + familyScope + "/config", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactRequired, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":policy", ArtifactType: ReferenceArchitectureValBArtifactPolicyBundle, Version: pack.Version, Scope: familyScope, SourceRef: pack.PolicyBundleRef, IntegrityRef: "integrity/" + familyScope + "/policy", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactRequired, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":readiness", ArtifactType: ReferenceArchitectureValBArtifactReadinessBundle, Version: pack.Version, Scope: familyScope, SourceRef: pack.ReadinessBundleRef, IntegrityRef: "integrity/" + familyScope + "/readiness", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactRequired, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":validation", ArtifactType: ReferenceArchitectureValBArtifactValidationPack, Version: pack.Version, Scope: familyScope, SourceRef: pack.ValidationPackRef, IntegrityRef: "integrity/" + familyScope + "/validation", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactRequired, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":conformance", ArtifactType: ReferenceArchitectureValBArtifactConformanceKit, Version: pack.Version, Scope: familyScope, SourceRef: pack.ConformanceKitRef, IntegrityRef: "integrity/" + familyScope + "/conformance", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactRequired, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":support", ArtifactType: ReferenceArchitectureValBArtifactSupportDoc, Version: pack.Version, Scope: familyScope, SourceRef: pack.SupportBoundaryRef, IntegrityRef: "integrity/" + familyScope + "/support", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactRequired, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":upgrade", ArtifactType: ReferenceArchitectureValBArtifactUpgradeNote, Version: pack.Version, Scope: familyScope, SourceRef: "upgrade-note/" + familyScope, IntegrityRef: "integrity/" + familyScope + "/upgrade", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactOptional, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
		{ArtifactID: pack.PackID + ":caveat", ArtifactType: ReferenceArchitectureValBArtifactCaveatNote, Version: pack.Version, Scope: familyScope, SourceRef: "caveat-note/" + familyScope, IntegrityRef: "integrity/" + familyScope + "/caveat", Timestamp: pack.UpdatedAt, FreshnessState: IntelligenceCalibrationFreshnessFresh, RequirementLevel: ReferenceArchitectureValBArtifactOptional, CompatibilityState: pack.CompatibilityState, Caveats: append([]string{}, pack.Caveats...)},
	}
	return ReferenceArchitectureArtifactManifest{
		CurrentState:         "reference_architecture_valb_manifest_ready",
		ManifestID:           "artifact-manifest/" + familyScope,
		PackID:               pack.PackID,
		BlueprintFamily:      familyScope,
		SupportedTypes:       referenceArchitectureValBArtifactTypes(),
		Artifacts:            artifacts,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValBArtifactManifestCollection() ReferenceArchitectureArtifactManifestCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	manifests := make([]ReferenceArchitectureArtifactManifest, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		manifests = append(manifests, referenceArchitectureValBArtifactManifestForPack(pack))
	}
	return ReferenceArchitectureArtifactManifestCollection{
		CurrentState:         "reference_architecture_valb_manifest_collection_ready",
		CollectionID:         "reference-architecture-valb-artifact-manifests",
		SupportedTypes:       referenceArchitectureValBArtifactTypes(),
		Manifests:            manifests,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func referenceArchitectureValBBundleForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureConfigProfilePolicyBundle {
	requiredGroups := []string{"identity", "signing", "audit", "evidence", "policy", "recovery"}
	forbiddenOptions := []string{"disable_audit_writer", "disable_evidence_storage", "skip_support_boundary"}
	profileAssumptions := []string{"family profile remains bounded by declared target environment and support boundary"}
	environmentPrereqs := []string{"target environment reference is declared", "profile reference is declared"}
	evidenceReqs := []string{ReferenceArchitectureEvidenceDeploymentObservation, ReferenceArchitectureEvidenceCapabilityAttestation, ReferenceArchitectureEvidenceCompatibilityReport}
	if pack.BlueprintFamily == ReferenceArchitectureFamilySovereignAirGapped {
		requiredGroups = append(requiredGroups, "air_gap_transfer")
		environmentPrereqs = append(environmentPrereqs, "offline transfer boundary is declared")
		forbiddenOptions = append(forbiddenOptions, "require_live_external_dependency")
	}
	if pack.BlueprintFamily == ReferenceArchitectureFamilyRegulatedPrivacyFirst {
		requiredGroups = append(requiredGroups, "residency", "redaction")
		forbiddenOptions = append(forbiddenOptions, "unbounded_evidence_export")
	}
	if pack.BlueprintFamily == ReferenceArchitectureFamilyPerformanceSensitive {
		requiredGroups = append(requiredGroups, "capacity", "performance_envelope")
	}
	if pack.BlueprintFamily == ReferenceArchitectureFamilyPartnerMSPSuitable {
		requiredGroups = append(requiredGroups, "customer_authority", "partner_visibility")
		forbiddenOptions = append(forbiddenOptions, "partner_canonical_override")
	}
	return ReferenceArchitectureConfigProfilePolicyBundle{
		CurrentState:                "reference_architecture_valb_bundle_ready",
		BundleID:                    "bundle/" + pack.BlueprintFamily,
		PackID:                      pack.PackID,
		BlueprintFamily:             pack.BlueprintFamily,
		RequiredConfigKeys:          []string{"tenant_id", "environment", "support_boundary_ref"},
		RequiredConfigGroups:        requiredGroups,
		ProfileAssumptions:          profileAssumptions,
		PolicyBaselineRefs:          []string{"policy-baseline/" + pack.BlueprintFamily},
		EnvironmentPrerequisites:    environmentPrereqs,
		ForbiddenUnsupportedOptions: forbiddenOptions,
		CompatibilityCaveats:        append([]string{}, pack.Caveats...),
		EvidenceRequirements:        evidenceReqs,
		ProjectionDisclaimer:        referenceArchitectureValBProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValBBundleCollection() ReferenceArchitectureBundleCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	bundles := make([]ReferenceArchitectureConfigProfilePolicyBundle, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		bundles = append(bundles, referenceArchitectureValBBundleForPack(pack))
	}
	return ReferenceArchitectureBundleCollection{
		CurrentState:         "reference_architecture_valb_bundle_collection_ready",
		CollectionID:         "reference-architecture-valb-bundles",
		Bundles:              bundles,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func referenceArchitectureValBReadinessBundleForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureReadinessBundle {
	checks := make([]ReferenceArchitectureReadinessCheck, 0, len(referenceArchitectureValBReadinessCategories()))
	for _, category := range referenceArchitectureValBReadinessCategories() {
		checks = append(checks, ReferenceArchitectureReadinessCheck{
			CheckID:               pack.PackID + ":" + category,
			Category:              category,
			State:                 ReferenceArchitectureValBReadinessReady,
			Timestamp:             pack.UpdatedAt,
			FreshnessState:        IntelligenceCalibrationFreshnessFresh,
			EvidenceRefs:          []string{"evidence/" + pack.BlueprintFamily + "/" + category},
			Caveats:               append([]string{}, pack.Caveats...),
			RedactionKeepsCaveats: true,
		})
	}
	return ReferenceArchitectureReadinessBundle{
		CurrentState:         "reference_architecture_valb_readiness_ready",
		BundleID:             "readiness/" + pack.BlueprintFamily,
		PackID:               pack.PackID,
		BlueprintFamily:      pack.BlueprintFamily,
		SupportedCategories:  referenceArchitectureValBReadinessCategories(),
		SupportedStates:      referenceArchitectureValBReadinessStates(),
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValBReadinessCollection() ReferenceArchitectureReadinessCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	bundles := make([]ReferenceArchitectureReadinessBundle, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		bundles = append(bundles, referenceArchitectureValBReadinessBundleForPack(pack))
	}
	return ReferenceArchitectureReadinessCollection{
		CurrentState:         "reference_architecture_valb_readiness_collection_ready",
		CollectionID:         "reference-architecture-valb-readiness",
		SupportedCategories:  referenceArchitectureValBReadinessCategories(),
		SupportedStates:      referenceArchitectureValBReadinessStates(),
		Bundles:              bundles,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func referenceArchitectureValBHookPackForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureValidationHookPack {
	hooks := make([]ReferenceArchitectureValidationHookDescriptor, 0, len(referenceArchitectureValBHookCategories()))
	for _, category := range referenceArchitectureValBHookCategories() {
		hooks = append(hooks, ReferenceArchitectureValidationHookDescriptor{
			HookID:                pack.PackID + ":" + category,
			Category:              category,
			ExpectedInputRefs:     []string{pack.ProfileRef, pack.ConfigBundleRef, pack.PolicyBundleRef},
			ExpectedOutputRefs:    []string{"result/" + pack.BlueprintFamily + "/" + category},
			RequiredEvidenceTypes: []string{ReferenceArchitectureEvidenceCapabilityAttestation, ReferenceArchitectureEvidenceCompatibilityReport},
			FailureSemantics:      ReferenceArchitectureValBHookFailureFailClosed,
			TimeoutExpectation:    "fresh_within_rfc3339_bounded_window",
			ProjectionDisclaimer:  referenceArchitectureValBProjectionDisclaimer(),
		})
	}
	return ReferenceArchitectureValidationHookPack{
		CurrentState:         "reference_architecture_valb_hook_pack_ready",
		HookPackRef:          referenceArchitectureValBHookPackRefForPack(pack),
		PackID:               pack.PackID,
		BlueprintFamily:      pack.BlueprintFamily,
		SupportedCategories:  referenceArchitectureValBHookCategories(),
		Hooks:                hooks,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValBValidationHookCollection() ReferenceArchitectureValidationHookCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	packs := make([]ReferenceArchitectureValidationHookPack, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		packs = append(packs, referenceArchitectureValBHookPackForPack(pack))
	}
	return ReferenceArchitectureValidationHookCollection{
		CurrentState:         "reference_architecture_valb_hook_collection_ready",
		CollectionID:         "reference-architecture-valb-validation-hooks",
		SupportedCategories:  referenceArchitectureValBHookCategories(),
		HookPacks:            packs,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func referenceArchitectureValBDeviationReportForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureDeviationReport {
	return ReferenceArchitectureDeviationReport{
		CurrentState:         "reference_architecture_valb_deviation_report_ready",
		ReportID:             "deviations/" + pack.BlueprintFamily,
		PackID:               pack.PackID,
		BlueprintFamily:      pack.BlueprintFamily,
		SupportedCategories:  referenceArchitectureValBDeviationCategories(),
		SupportedSeverities:  referenceArchitectureValBDeviationSeverities(),
		Deviations:           []ReferenceArchitectureDeviation{},
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValBDeviationCollection() ReferenceArchitectureDeviationCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	reports := make([]ReferenceArchitectureDeviationReport, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		reports = append(reports, referenceArchitectureValBDeviationReportForPack(pack))
	}
	return ReferenceArchitectureDeviationCollection{
		CurrentState:         "reference_architecture_valb_deviation_collection_ready",
		CollectionID:         "reference-architecture-valb-deviations",
		SupportedCategories:  referenceArchitectureValBDeviationCategories(),
		SupportedSeverities:  referenceArchitectureValBDeviationSeverities(),
		Reports:              reports,
		ProjectionDisclaimer: referenceArchitectureValBProjectionDisclaimer(),
	}
}

func referenceArchitectureValBConformanceKitForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureConformanceKit {
	return ReferenceArchitectureConformanceKit{
		CurrentState:          "reference_architecture_valb_conformance_kit_ready",
		KitID:                 "conformance-kit/" + pack.BlueprintFamily,
		PackID:                pack.PackID,
		BlueprintFamily:       pack.BlueprintFamily,
		ProfileRef:            pack.ProfileRef,
		ManifestRef:           "artifact-manifest/" + pack.BlueprintFamily,
		BundleRef:             "bundle/" + pack.BlueprintFamily,
		ReadinessBundleRef:    "readiness/" + pack.BlueprintFamily,
		ValidationHookPackRef: referenceArchitectureValBHookPackRefForPack(pack),
		DeviationReportRef:    "deviations/" + pack.BlueprintFamily,
		EvidenceRefs:          append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		ConformanceState:      ReferenceArchitectureConformanceMatched,
		ProjectionDisclaimer:  referenceArchitectureValBProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValBConformanceKitCollection() ReferenceArchitectureConformanceKitCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	kits := make([]ReferenceArchitectureConformanceKit, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		kits = append(kits, referenceArchitectureValBConformanceKitForPack(pack))
	}
	return ReferenceArchitectureConformanceKitCollection{
		CurrentState:               "reference_architecture_valb_conformance_collection_ready",
		CollectionID:               "reference-architecture-valb-conformance-kits",
		SupportedConformanceStates: referenceArchitectureVal0ConformanceStates(),
		Kits:                       kits,
		ProjectionDisclaimer:       referenceArchitectureValBProjectionDisclaimer(),
	}
}

func referenceArchitectureValBHasProjectionDisclaimer(value string) bool {
	return strings.Contains(strings.TrimSpace(value), "projection_only") &&
		strings.Contains(strings.TrimSpace(value), "not_canonical_truth")
}

func referenceArchitectureValBEvidenceValid(refs []ReferenceArchitectureEvidenceReference) (allFresh bool, stale bool, ok bool) {
	if len(refs) == 0 {
		return false, false, false
	}
	allFresh = true
	for _, evidence := range refs {
		if strings.TrimSpace(evidence.EvidenceID) == "" ||
			strings.TrimSpace(evidence.Source) == "" ||
			strings.TrimSpace(evidence.Scope) == "" ||
			len(evidence.Caveats) == 0 ||
			!containsTrimmedString(referenceArchitectureVal0SupportedEvidenceTypes(), evidence.EvidenceType) ||
			!containsTrimmedString([]string{
				IntelligenceCalibrationFreshnessFresh,
				IntelligenceCalibrationFreshnessStale,
				IntelligenceCalibrationFreshnessUnsupported,
				IntelligenceCalibrationFreshnessUnknown,
			}, evidence.FreshnessState) {
			return false, false, false
		}
		if _, ok := referenceArchitectureVal0ParseTimestamp(evidence.Timestamp); !ok {
			return false, false, false
		}
		if strings.TrimSpace(evidence.FreshnessState) != IntelligenceCalibrationFreshnessFresh {
			allFresh = false
			if strings.TrimSpace(evidence.FreshnessState) == IntelligenceCalibrationFreshnessStale {
				stale = true
			}
		}
	}
	return allFresh, stale, true
}

func referenceArchitectureValBRequiredRefsPresent(values ...string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}

func EvaluateReferenceArchitectureValBPackState(pack ReferenceArchitectureBlueprintPack) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		pack.PackID,
		pack.Version,
		pack.BlueprintFamily,
		pack.BlueprintID,
		pack.LifecycleState,
		pack.CompatibilityState,
		pack.Owner,
		pack.TargetEnvironmentRef,
		pack.ProfileRef,
		pack.ConfigBundleRef,
		pack.PolicyBundleRef,
		pack.ReadinessBundleRef,
		pack.ValidationPackRef,
		pack.ConformanceKitRef,
		pack.SupportBoundaryRef,
		pack.CreatedAt,
		pack.UpdatedAt,
		pack.ProjectionDisclaimer,
	) {
		return ReferenceArchitectureValBPackStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), pack.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0LifecycleStates(), pack.LifecycleState) ||
		!containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), pack.CompatibilityState) ||
		!referenceArchitectureValBHasProjectionDisclaimer(pack.ProjectionDisclaimer) {
		return ReferenceArchitectureValBPackStatePartial
	}
	if pack.CertifiedLanguagePresent || pack.GuaranteedDeploymentClaim || pack.ClaimsPoint6Pass {
		return ReferenceArchitectureValBPackStateBlocked
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(pack.CreatedAt); !ok {
		return ReferenceArchitectureValBPackStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(pack.UpdatedAt); !ok {
		return ReferenceArchitectureValBPackStatePartial
	}
	profile, found := LookupReferenceArchitectureValAFamilyProfile(pack.BlueprintFamily)
	if !found || strings.TrimSpace(profile.BlueprintID) != strings.TrimSpace(pack.BlueprintID) || strings.TrimSpace(profile.SupportBoundaryRef) != strings.TrimSpace(pack.SupportBoundaryRef) {
		return ReferenceArchitectureValBPackStatePartial
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(pack.EvidenceRefs)
	if !ok {
		return ReferenceArchitectureValBPackStatePartial
	}
	if !allFresh || stale {
		return ReferenceArchitectureValBPackStatePartial
	}
	if pack.LifecycleState != ReferenceArchitectureLifecycleActive || pack.CompatibilityState != ReferenceArchitectureCompatibilityCompatible {
		return ReferenceArchitectureValBPackStatePartial
	}
	return ReferenceArchitectureValBPackStateActive
}

func EvaluateReferenceArchitectureValBPackRegistryState(registry ReferenceArchitectureBlueprintPackRegistry) string {
	if strings.TrimSpace(registry.RegistryID) == "" || strings.TrimSpace(registry.Version) == "" || strings.TrimSpace(registry.ProjectionDisclaimer) == "" || len(registry.Packs) == 0 {
		return ReferenceArchitectureValBPackStateIncomplete
	}
	if !containsExactTrimmedStringSet(registry.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValBHasProjectionDisclaimer(registry.ProjectionDisclaimer) ||
		len(registry.Packs) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValBPackStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenPacks := map[string]struct{}{}
	for _, pack := range registry.Packs {
		family := strings.TrimSpace(pack.BlueprintFamily)
		packID := strings.TrimSpace(pack.PackID)
		if family == "" || packID == "" {
			return ReferenceArchitectureValBPackStateIncomplete
		}
		if _, ok := seenFamilies[family]; ok {
			return ReferenceArchitectureValBPackStatePartial
		}
		if _, ok := seenPacks[packID]; ok {
			return ReferenceArchitectureValBPackStatePartial
		}
		seenFamilies[family] = struct{}{}
		seenPacks[packID] = struct{}{}
		if EvaluateReferenceArchitectureValBPackState(pack) != ReferenceArchitectureValBPackStateActive {
			return ReferenceArchitectureValBPackStatePartial
		}
	}
	return ReferenceArchitectureValBPackStateActive
}

func EvaluateReferenceArchitectureValBArtifactManifestState(manifest ReferenceArchitectureArtifactManifest) string {
	if !referenceArchitectureValBRequiredRefsPresent(manifest.ManifestID, manifest.PackID, manifest.BlueprintFamily, manifest.ProjectionDisclaimer) || len(manifest.Artifacts) == 0 {
		return ReferenceArchitectureValBManifestStateIncomplete
	}
	if !containsExactTrimmedStringSet(manifest.SupportedTypes, referenceArchitectureValBArtifactTypes()...) || !referenceArchitectureValBHasProjectionDisclaimer(manifest.ProjectionDisclaimer) {
		return ReferenceArchitectureValBManifestStatePartial
	}
	requiredPresent := map[string]struct{}{}
	seenIDs := map[string]struct{}{}
	for _, artifact := range manifest.Artifacts {
		if !referenceArchitectureValBRequiredRefsPresent(artifact.ArtifactID, artifact.ArtifactType, artifact.Version, artifact.Scope, artifact.SourceRef, artifact.Timestamp, artifact.RequirementLevel, artifact.CompatibilityState) {
			return ReferenceArchitectureValBManifestStateIncomplete
		}
		if _, ok := seenIDs[strings.TrimSpace(artifact.ArtifactID)]; ok {
			return ReferenceArchitectureValBManifestStatePartial
		}
		seenIDs[strings.TrimSpace(artifact.ArtifactID)] = struct{}{}
		if !containsTrimmedString(referenceArchitectureValBArtifactTypes(), artifact.ArtifactType) ||
			!containsTrimmedString(referenceArchitectureValBArtifactRequirementLevels(), artifact.RequirementLevel) ||
			!containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), artifact.CompatibilityState) ||
			!containsTrimmedString([]string{IntelligenceCalibrationFreshnessFresh, IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessUnsupported, IntelligenceCalibrationFreshnessUnknown}, artifact.FreshnessState) {
			return ReferenceArchitectureValBManifestStatePartial
		}
		if _, ok := referenceArchitectureVal0ParseTimestamp(artifact.Timestamp); !ok {
			return ReferenceArchitectureValBManifestStatePartial
		}
		if strings.TrimSpace(artifact.FreshnessState) != IntelligenceCalibrationFreshnessFresh {
			return ReferenceArchitectureValBManifestStatePartial
		}
		if strings.TrimSpace(artifact.CompatibilityState) == ReferenceArchitectureCompatibilityUnsupported {
			return ReferenceArchitectureValBManifestStatePartial
		}
		if artifact.RequirementLevel == ReferenceArchitectureValBArtifactRequired {
			requiredPresent[strings.TrimSpace(artifact.ArtifactType)] = struct{}{}
		}
	}
	if len(requiredPresent) != len(referenceArchitectureValBRequiredArtifactTypes()) {
		return ReferenceArchitectureValBManifestStatePartial
	}
	for _, artifactType := range referenceArchitectureValBRequiredArtifactTypes() {
		if _, ok := requiredPresent[strings.TrimSpace(artifactType)]; !ok {
			return ReferenceArchitectureValBManifestStatePartial
		}
	}
	return ReferenceArchitectureValBManifestStateActive
}

func EvaluateReferenceArchitectureValBArtifactManifestCollectionState(collection ReferenceArchitectureArtifactManifestCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Manifests) == 0 {
		return ReferenceArchitectureValBManifestStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedTypes, referenceArchitectureValBArtifactTypes()...) || !referenceArchitectureValBHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.Manifests) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValBManifestStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenPacks := map[string]struct{}{}
	for _, manifest := range collection.Manifests {
		if _, ok := seenFamilies[strings.TrimSpace(manifest.BlueprintFamily)]; ok {
			return ReferenceArchitectureValBManifestStatePartial
		}
		if _, ok := seenPacks[strings.TrimSpace(manifest.PackID)]; ok {
			return ReferenceArchitectureValBManifestStatePartial
		}
		seenFamilies[strings.TrimSpace(manifest.BlueprintFamily)] = struct{}{}
		seenPacks[strings.TrimSpace(manifest.PackID)] = struct{}{}
		if EvaluateReferenceArchitectureValBArtifactManifestState(manifest) != ReferenceArchitectureValBManifestStateActive {
			return ReferenceArchitectureValBManifestStatePartial
		}
	}
	return ReferenceArchitectureValBManifestStateActive
}

func EvaluateReferenceArchitectureValBBundleState(bundle ReferenceArchitectureConfigProfilePolicyBundle) string {
	if !referenceArchitectureValBRequiredRefsPresent(bundle.BundleID, bundle.PackID, bundle.BlueprintFamily, bundle.ProjectionDisclaimer) {
		return ReferenceArchitectureValBBundleStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), bundle.BlueprintFamily) || !referenceArchitectureValBHasProjectionDisclaimer(bundle.ProjectionDisclaimer) {
		return ReferenceArchitectureValBBundleStatePartial
	}
	if len(bundle.RequiredConfigKeys) == 0 ||
		len(bundle.RequiredConfigGroups) == 0 ||
		len(bundle.ProfileAssumptions) == 0 ||
		len(bundle.PolicyBaselineRefs) == 0 ||
		len(bundle.EnvironmentPrerequisites) == 0 ||
		len(bundle.ForbiddenUnsupportedOptions) == 0 ||
		len(bundle.EvidenceRequirements) == 0 {
		return ReferenceArchitectureValBBundleStatePartial
	}
	if !containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), bundle.EvidenceRequirements...) {
		return ReferenceArchitectureValBBundleStatePartial
	}
	return ReferenceArchitectureValBBundleStateActive
}

func EvaluateReferenceArchitectureValBBundleCollectionState(collection ReferenceArchitectureBundleCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Bundles) == 0 {
		return ReferenceArchitectureValBBundleStateIncomplete
	}
	if !referenceArchitectureValBHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.Bundles) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValBBundleStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenPacks := map[string]struct{}{}
	for _, bundle := range collection.Bundles {
		if _, ok := seenFamilies[strings.TrimSpace(bundle.BlueprintFamily)]; ok {
			return ReferenceArchitectureValBBundleStatePartial
		}
		if _, ok := seenPacks[strings.TrimSpace(bundle.PackID)]; ok {
			return ReferenceArchitectureValBBundleStatePartial
		}
		seenFamilies[strings.TrimSpace(bundle.BlueprintFamily)] = struct{}{}
		seenPacks[strings.TrimSpace(bundle.PackID)] = struct{}{}
		if EvaluateReferenceArchitectureValBBundleState(bundle) != ReferenceArchitectureValBBundleStateActive {
			return ReferenceArchitectureValBBundleStatePartial
		}
	}
	return ReferenceArchitectureValBBundleStateActive
}

func EvaluateReferenceArchitectureValBReadinessBundleState(bundle ReferenceArchitectureReadinessBundle) string {
	if !referenceArchitectureValBRequiredRefsPresent(bundle.BundleID, bundle.PackID, bundle.BlueprintFamily, bundle.ProjectionDisclaimer) || len(bundle.Checks) == 0 {
		return ReferenceArchitectureValBReadinessStateIncomplete
	}
	if !containsExactTrimmedStringSet(bundle.SupportedCategories, referenceArchitectureValBReadinessCategories()...) ||
		!containsExactTrimmedStringSet(bundle.SupportedStates, referenceArchitectureValBReadinessStates()...) ||
		!referenceArchitectureValBHasProjectionDisclaimer(bundle.ProjectionDisclaimer) {
		return ReferenceArchitectureValBReadinessStatePartial
	}
	seenCategories := map[string]struct{}{}
	for _, check := range bundle.Checks {
		if !referenceArchitectureValBRequiredRefsPresent(check.CheckID, check.Category, check.State, check.Timestamp) || len(check.EvidenceRefs) == 0 {
			return ReferenceArchitectureValBReadinessStateIncomplete
		}
		if _, ok := seenCategories[strings.TrimSpace(check.Category)]; ok {
			return ReferenceArchitectureValBReadinessStatePartial
		}
		seenCategories[strings.TrimSpace(check.Category)] = struct{}{}
		if !containsTrimmedString(referenceArchitectureValBReadinessCategories(), check.Category) ||
			!containsTrimmedString(referenceArchitectureValBReadinessStates(), check.State) ||
			!containsTrimmedString([]string{IntelligenceCalibrationFreshnessFresh, IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessUnsupported, IntelligenceCalibrationFreshnessUnknown}, check.FreshnessState) {
			return ReferenceArchitectureValBReadinessStatePartial
		}
		if _, ok := referenceArchitectureVal0ParseTimestamp(check.Timestamp); !ok {
			return ReferenceArchitectureValBReadinessStatePartial
		}
		if !check.RedactionKeepsCaveats || strings.TrimSpace(check.FreshnessState) != IntelligenceCalibrationFreshnessFresh || strings.TrimSpace(check.State) != ReferenceArchitectureValBReadinessReady {
			return ReferenceArchitectureValBReadinessStatePartial
		}
		if (check.State == ReferenceArchitectureValBReadinessDegraded || check.State == ReferenceArchitectureValBReadinessUnsupported) && len(check.Caveats) == 0 {
			return ReferenceArchitectureValBReadinessStatePartial
		}
	}
	if len(seenCategories) != len(referenceArchitectureValBReadinessCategories()) {
		return ReferenceArchitectureValBReadinessStatePartial
	}
	return ReferenceArchitectureValBReadinessStateActive
}

func EvaluateReferenceArchitectureValBReadinessCollectionState(collection ReferenceArchitectureReadinessCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Bundles) == 0 {
		return ReferenceArchitectureValBReadinessStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedCategories, referenceArchitectureValBReadinessCategories()...) ||
		!containsExactTrimmedStringSet(collection.SupportedStates, referenceArchitectureValBReadinessStates()...) ||
		!referenceArchitectureValBHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Bundles) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValBReadinessStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenPacks := map[string]struct{}{}
	for _, bundle := range collection.Bundles {
		if _, ok := seenFamilies[strings.TrimSpace(bundle.BlueprintFamily)]; ok {
			return ReferenceArchitectureValBReadinessStatePartial
		}
		if _, ok := seenPacks[strings.TrimSpace(bundle.PackID)]; ok {
			return ReferenceArchitectureValBReadinessStatePartial
		}
		seenFamilies[strings.TrimSpace(bundle.BlueprintFamily)] = struct{}{}
		seenPacks[strings.TrimSpace(bundle.PackID)] = struct{}{}
		if EvaluateReferenceArchitectureValBReadinessBundleState(bundle) != ReferenceArchitectureValBReadinessStateActive {
			return ReferenceArchitectureValBReadinessStatePartial
		}
	}
	return ReferenceArchitectureValBReadinessStateActive
}

func EvaluateReferenceArchitectureValBHookPackState(pack ReferenceArchitectureValidationHookPack) string {
	if !referenceArchitectureValBRequiredRefsPresent(pack.HookPackRef, pack.PackID, pack.BlueprintFamily, pack.ProjectionDisclaimer) || len(pack.Hooks) == 0 {
		return ReferenceArchitectureValBHookStateIncomplete
	}
	if !containsExactTrimmedStringSet(pack.SupportedCategories, referenceArchitectureValBHookCategories()...) || !referenceArchitectureValBHasProjectionDisclaimer(pack.ProjectionDisclaimer) {
		return ReferenceArchitectureValBHookStatePartial
	}
	if strings.TrimSpace(pack.HookPackRef) != referenceArchitectureValBHookPackRefForPack(ReferenceArchitectureBlueprintPack{BlueprintFamily: pack.BlueprintFamily}) {
		return ReferenceArchitectureValBHookStatePartial
	}
	seenCategories := map[string]struct{}{}
	for _, hook := range pack.Hooks {
		if !referenceArchitectureValBRequiredRefsPresent(hook.HookID, hook.Category, hook.FailureSemantics, hook.TimeoutExpectation, hook.ProjectionDisclaimer) ||
			len(hook.ExpectedInputRefs) == 0 ||
			len(hook.ExpectedOutputRefs) == 0 ||
			len(hook.RequiredEvidenceTypes) == 0 {
			return ReferenceArchitectureValBHookStateIncomplete
		}
		if _, ok := seenCategories[strings.TrimSpace(hook.Category)]; ok {
			return ReferenceArchitectureValBHookStatePartial
		}
		seenCategories[strings.TrimSpace(hook.Category)] = struct{}{}
		if !containsTrimmedString(referenceArchitectureValBHookCategories(), hook.Category) ||
			!containsTrimmedString(referenceArchitectureValBHookFailureSemantics(), hook.FailureSemantics) ||
			!containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), hook.RequiredEvidenceTypes...) ||
			!referenceArchitectureValBHasProjectionDisclaimer(hook.ProjectionDisclaimer) {
			return ReferenceArchitectureValBHookStatePartial
		}
	}
	if len(seenCategories) != len(referenceArchitectureValBHookCategories()) {
		return ReferenceArchitectureValBHookStatePartial
	}
	return ReferenceArchitectureValBHookStateActive
}

func EvaluateReferenceArchitectureValBValidationHookCollectionState(collection ReferenceArchitectureValidationHookCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.HookPacks) == 0 {
		return ReferenceArchitectureValBHookStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedCategories, referenceArchitectureValBHookCategories()...) || !referenceArchitectureValBHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.HookPacks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValBHookStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenPacks := map[string]struct{}{}
	for _, pack := range collection.HookPacks {
		if _, ok := seenFamilies[strings.TrimSpace(pack.BlueprintFamily)]; ok {
			return ReferenceArchitectureValBHookStatePartial
		}
		if _, ok := seenPacks[strings.TrimSpace(pack.PackID)]; ok {
			return ReferenceArchitectureValBHookStatePartial
		}
		seenFamilies[strings.TrimSpace(pack.BlueprintFamily)] = struct{}{}
		seenPacks[strings.TrimSpace(pack.PackID)] = struct{}{}
		if EvaluateReferenceArchitectureValBHookPackState(pack) != ReferenceArchitectureValBHookStateActive {
			return ReferenceArchitectureValBHookStatePartial
		}
	}
	return ReferenceArchitectureValBHookStateActive
}

func EvaluateReferenceArchitectureValBDeviationReportState(report ReferenceArchitectureDeviationReport) string {
	if !referenceArchitectureValBRequiredRefsPresent(report.ReportID, report.PackID, report.BlueprintFamily, report.ProjectionDisclaimer) {
		return ReferenceArchitectureValBDeviationStateIncomplete
	}
	if !containsExactTrimmedStringSet(report.SupportedCategories, referenceArchitectureValBDeviationCategories()...) ||
		!containsExactTrimmedStringSet(report.SupportedSeverities, referenceArchitectureValBDeviationSeverities()...) ||
		!referenceArchitectureValBHasProjectionDisclaimer(report.ProjectionDisclaimer) {
		return ReferenceArchitectureValBDeviationStatePartial
	}
	seen := map[string]struct{}{}
	for _, deviation := range report.Deviations {
		if !referenceArchitectureValBRequiredRefsPresent(deviation.DeviationID, deviation.Category, deviation.Severity, deviation.AffectedScope, deviation.Explanation) {
			return ReferenceArchitectureValBDeviationStateIncomplete
		}
		if _, ok := seen[strings.TrimSpace(deviation.DeviationID)]; ok {
			return ReferenceArchitectureValBDeviationStatePartial
		}
		seen[strings.TrimSpace(deviation.DeviationID)] = struct{}{}
		if !containsTrimmedString(referenceArchitectureValBDeviationCategories(), deviation.Category) ||
			!containsTrimmedString(referenceArchitectureValBDeviationSeverities(), deviation.Severity) {
			return ReferenceArchitectureValBDeviationStatePartial
		}
	}
	return ReferenceArchitectureValBDeviationStateActive
}

func EvaluateReferenceArchitectureValBDeviationCollectionState(collection ReferenceArchitectureDeviationCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Reports) == 0 {
		return ReferenceArchitectureValBDeviationStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedCategories, referenceArchitectureValBDeviationCategories()...) ||
		!containsExactTrimmedStringSet(collection.SupportedSeverities, referenceArchitectureValBDeviationSeverities()...) ||
		!referenceArchitectureValBHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Reports) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValBDeviationStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenPacks := map[string]struct{}{}
	for _, report := range collection.Reports {
		if _, ok := seenFamilies[strings.TrimSpace(report.BlueprintFamily)]; ok {
			return ReferenceArchitectureValBDeviationStatePartial
		}
		if _, ok := seenPacks[strings.TrimSpace(report.PackID)]; ok {
			return ReferenceArchitectureValBDeviationStatePartial
		}
		seenFamilies[strings.TrimSpace(report.BlueprintFamily)] = struct{}{}
		seenPacks[strings.TrimSpace(report.PackID)] = struct{}{}
		if EvaluateReferenceArchitectureValBDeviationReportState(report) != ReferenceArchitectureValBDeviationStateActive {
			return ReferenceArchitectureValBDeviationStatePartial
		}
	}
	return ReferenceArchitectureValBDeviationStateActive
}

func referenceArchitectureValBDeviationBlocksMatched(report ReferenceArchitectureDeviationReport) bool {
	for _, deviation := range report.Deviations {
		if deviation.BlocksMatched {
			return true
		}
	}
	return false
}

func referenceArchitectureValBFindBlockingDeviation(report ReferenceArchitectureDeviationReport, category string) bool {
	for _, deviation := range report.Deviations {
		if strings.TrimSpace(deviation.Category) == strings.TrimSpace(category) && deviation.BlocksMatched {
			return true
		}
	}
	return false
}

func EvaluateReferenceArchitectureValBConformanceKitState(
	packState, manifestState, bundleState, readinessState, hookState, deviationState string,
	kit ReferenceArchitectureConformanceKit,
	pack ReferenceArchitectureBlueprintPack,
	report ReferenceArchitectureDeviationReport,
) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		kit.KitID,
		kit.PackID,
		kit.BlueprintFamily,
		kit.ProfileRef,
		kit.ManifestRef,
		kit.BundleRef,
		kit.ReadinessBundleRef,
		kit.ValidationHookPackRef,
		kit.DeviationReportRef,
		kit.ConformanceState,
		kit.ProjectionDisclaimer,
	) {
		return ReferenceArchitectureValBConformanceKitStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), kit.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0ConformanceStates(), kit.ConformanceState) ||
		!referenceArchitectureValBHasProjectionDisclaimer(kit.ProjectionDisclaimer) {
		return ReferenceArchitectureValBConformanceKitStateUnknown
	}
	if strings.TrimSpace(kit.ValidationHookPackRef) != referenceArchitectureValBHookPackRefForPack(pack) {
		return ReferenceArchitectureConformancePartiallyMatched
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(kit.EvidenceRefs)
	if !ok {
		return ReferenceArchitectureValBConformanceKitStateUnknown
	}
	if pack.CompatibilityState == ReferenceArchitectureCompatibilityUnsupported || pack.LifecycleState == ReferenceArchitectureLifecycleUnsupported {
		return ReferenceArchitectureConformanceUnsupported
	}
	if pack.CompatibilityState == ReferenceArchitectureCompatibilitySuperseded ||
		pack.CompatibilityState == ReferenceArchitectureCompatibilityDeprecated ||
		pack.LifecycleState == ReferenceArchitectureLifecycleSuperseded ||
		pack.LifecycleState == ReferenceArchitectureLifecycleDeprecated ||
		stale || !allFresh {
		return ReferenceArchitectureConformanceDrifted
	}
	if packState != ReferenceArchitectureValBPackStateActive ||
		manifestState != ReferenceArchitectureValBManifestStateActive ||
		bundleState != ReferenceArchitectureValBBundleStateActive ||
		readinessState != ReferenceArchitectureValBReadinessStateActive ||
		hookState != ReferenceArchitectureValBHookStateActive ||
		deviationState != ReferenceArchitectureValBDeviationStateActive {
		return ReferenceArchitectureConformancePartiallyMatched
	}
	if referenceArchitectureValBDeviationBlocksMatched(report) {
		if referenceArchitectureValBFindBlockingDeviation(report, ReferenceArchitectureValBDeviationUnsupportedCompatibility) {
			return ReferenceArchitectureConformanceUnsupported
		}
		return ReferenceArchitectureConformanceDegraded
	}
	if len(kit.UnsupportedReasons) > 0 {
		return ReferenceArchitectureConformanceUnsupported
	}
	if len(kit.DegradedReasons) > 0 {
		return ReferenceArchitectureConformanceDegraded
	}
	if len(kit.Caveats) > 0 {
		return ReferenceArchitectureConformancePartiallyMatched
	}
	return ReferenceArchitectureConformanceMatched
}

func EvaluateReferenceArchitectureValBConformanceKitCollectionState(
	kits ReferenceArchitectureConformanceKitCollection,
	registry ReferenceArchitectureBlueprintPackRegistry,
	manifestCollection ReferenceArchitectureArtifactManifestCollection,
	bundleCollection ReferenceArchitectureBundleCollection,
	readinessCollection ReferenceArchitectureReadinessCollection,
	hookCollection ReferenceArchitectureValidationHookCollection,
	deviationCollection ReferenceArchitectureDeviationCollection,
) string {
	if strings.TrimSpace(kits.CollectionID) == "" || strings.TrimSpace(kits.ProjectionDisclaimer) == "" || len(kits.Kits) == 0 {
		return ReferenceArchitectureValBConformanceKitStateIncomplete
	}
	if !containsExactTrimmedStringSet(kits.SupportedConformanceStates, referenceArchitectureVal0ConformanceStates()...) ||
		!referenceArchitectureValBHasProjectionDisclaimer(kits.ProjectionDisclaimer) ||
		len(kits.Kits) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValBConformanceKitStatePartial
	}
	registryState := EvaluateReferenceArchitectureValBPackRegistryState(registry)
	manifestState := EvaluateReferenceArchitectureValBArtifactManifestCollectionState(manifestCollection)
	bundleState := EvaluateReferenceArchitectureValBBundleCollectionState(bundleCollection)
	readinessState := EvaluateReferenceArchitectureValBReadinessCollectionState(readinessCollection)
	hookState := EvaluateReferenceArchitectureValBValidationHookCollectionState(hookCollection)
	deviationState := EvaluateReferenceArchitectureValBDeviationCollectionState(deviationCollection)
	dependencyStates := []string{registryState, manifestState, bundleState, readinessState, hookState, deviationState}
	for _, state := range dependencyStates {
		switch strings.TrimSpace(state) {
		case ReferenceArchitectureValBPackStateActive,
			ReferenceArchitectureValBManifestStateActive,
			ReferenceArchitectureValBBundleStateActive,
			ReferenceArchitectureValBReadinessStateActive,
			ReferenceArchitectureValBHookStateActive,
			ReferenceArchitectureValBDeviationStateActive:
		case ReferenceArchitectureValBPackStateIncomplete,
			ReferenceArchitectureValBManifestStateIncomplete,
			ReferenceArchitectureValBBundleStateIncomplete,
			ReferenceArchitectureValBReadinessStateIncomplete,
			ReferenceArchitectureValBHookStateIncomplete,
			ReferenceArchitectureValBDeviationStateIncomplete:
			return ReferenceArchitectureValBConformanceKitStateIncomplete
		case ReferenceArchitectureValBPackStateBlocked,
			ReferenceArchitectureValBManifestStateBlocked,
			ReferenceArchitectureValBBundleStateBlocked,
			ReferenceArchitectureValBReadinessStateBlocked,
			ReferenceArchitectureValBHookStateBlocked,
			ReferenceArchitectureValBDeviationStateBlocked:
			return ReferenceArchitectureValBConformanceKitStateBlocked
		case ReferenceArchitectureValBPackStateUnknown,
			ReferenceArchitectureValBManifestStateUnknown,
			ReferenceArchitectureValBBundleStateUnknown,
			ReferenceArchitectureValBReadinessStateUnknown,
			ReferenceArchitectureValBHookStateUnknown,
			ReferenceArchitectureValBDeviationStateUnknown:
			return ReferenceArchitectureValBConformanceKitStateUnknown
		default:
			return ReferenceArchitectureValBConformanceKitStatePartial
		}
	}
	packsByFamily := map[string]ReferenceArchitectureBlueprintPack{}
	for _, pack := range registry.Packs {
		packsByFamily[strings.TrimSpace(pack.BlueprintFamily)] = pack
	}
	manifestsByFamily := map[string]ReferenceArchitectureArtifactManifest{}
	for _, manifest := range manifestCollection.Manifests {
		manifestsByFamily[strings.TrimSpace(manifest.BlueprintFamily)] = manifest
	}
	bundlesByFamily := map[string]ReferenceArchitectureConfigProfilePolicyBundle{}
	for _, bundle := range bundleCollection.Bundles {
		bundlesByFamily[strings.TrimSpace(bundle.BlueprintFamily)] = bundle
	}
	readinessByFamily := map[string]ReferenceArchitectureReadinessBundle{}
	for _, bundle := range readinessCollection.Bundles {
		readinessByFamily[strings.TrimSpace(bundle.BlueprintFamily)] = bundle
	}
	hooksByRef := map[string]ReferenceArchitectureValidationHookPack{}
	for _, hookPack := range hookCollection.HookPacks {
		hooksByRef[strings.TrimSpace(hookPack.HookPackRef)] = hookPack
	}
	deviationsByFamily := map[string]ReferenceArchitectureDeviationReport{}
	for _, report := range deviationCollection.Reports {
		deviationsByFamily[strings.TrimSpace(report.BlueprintFamily)] = report
	}
	seenFamilies := map[string]struct{}{}
	for _, kit := range kits.Kits {
		family := strings.TrimSpace(kit.BlueprintFamily)
		if _, ok := seenFamilies[family]; ok {
			return ReferenceArchitectureValBConformanceKitStatePartial
		}
		seenFamilies[family] = struct{}{}
		pack := packsByFamily[family]
		manifest := manifestsByFamily[family]
		bundle := bundlesByFamily[family]
		readiness := readinessByFamily[family]
		hookPack, ok := hooksByRef[strings.TrimSpace(kit.ValidationHookPackRef)]
		if !ok {
			return ReferenceArchitectureValBConformanceKitStatePartial
		}
		if strings.TrimSpace(hookPack.BlueprintFamily) != family || strings.TrimSpace(hookPack.PackID) != strings.TrimSpace(kit.PackID) {
			return ReferenceArchitectureValBConformanceKitStatePartial
		}
		report := deviationsByFamily[family]
		conformanceState := EvaluateReferenceArchitectureValBConformanceKitState(
			EvaluateReferenceArchitectureValBPackState(pack),
			EvaluateReferenceArchitectureValBArtifactManifestState(manifest),
			EvaluateReferenceArchitectureValBBundleState(bundle),
			EvaluateReferenceArchitectureValBReadinessBundleState(readiness),
			EvaluateReferenceArchitectureValBHookPackState(hookPack),
			EvaluateReferenceArchitectureValBDeviationReportState(report),
			kit,
			pack,
			report,
		)
		if conformanceState != ReferenceArchitectureConformanceMatched {
			return ReferenceArchitectureValBConformanceKitStatePartial
		}
	}
	return ReferenceArchitectureValBConformanceKitStateActive
}

func referenceArchitectureValBRequiresPriorStates(point5State, val0CurrentState, val0State, valACurrentState, valAState, point6State string) bool {
	return strings.TrimSpace(point5State) == IntelligenceCalibrationPoint5StatePass &&
		strings.TrimSpace(val0CurrentState) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(val0State) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(valACurrentState) == ReferenceArchitectureValAStateActive &&
		strings.TrimSpace(valAState) == ReferenceArchitectureValAStateActive &&
		strings.TrimSpace(point6State) == ReferenceArchitecturePoint6StateNotComplete
}

func EvaluateReferenceArchitectureValBState(
	point5State, val0CurrentState, val0State, valACurrentState, valAState, point6State,
	packState, manifestState, bundleState, readinessState, hookState, conformanceKitState, deviationState string,
) string {
	if !referenceArchitectureValBRequiresPriorStates(point5State, val0CurrentState, val0State, valACurrentState, valAState, point6State) {
		return ReferenceArchitectureValBStateBlocked
	}
	states := []string{packState, manifestState, bundleState, readinessState, hookState, conformanceKitState, deviationState}
	for _, state := range states {
		if strings.TrimSpace(state) == "" {
			return ReferenceArchitectureValBStateIncomplete
		}
	}
	if packState == ReferenceArchitectureValBPackStateActive &&
		manifestState == ReferenceArchitectureValBManifestStateActive &&
		bundleState == ReferenceArchitectureValBBundleStateActive &&
		readinessState == ReferenceArchitectureValBReadinessStateActive &&
		hookState == ReferenceArchitectureValBHookStateActive &&
		conformanceKitState == ReferenceArchitectureValBConformanceKitStateActive &&
		deviationState == ReferenceArchitectureValBDeviationStateActive {
		return ReferenceArchitectureValBStateActive
	}
	if packState == ReferenceArchitectureValBPackStateIncomplete ||
		manifestState == ReferenceArchitectureValBManifestStateIncomplete ||
		bundleState == ReferenceArchitectureValBBundleStateIncomplete ||
		readinessState == ReferenceArchitectureValBReadinessStateIncomplete ||
		hookState == ReferenceArchitectureValBHookStateIncomplete ||
		conformanceKitState == ReferenceArchitectureValBConformanceKitStateIncomplete ||
		deviationState == ReferenceArchitectureValBDeviationStateIncomplete {
		return ReferenceArchitectureValBStateIncomplete
	}
	return ReferenceArchitectureValBStatePartial
}

func referenceArchitectureValBProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/pack-registry",
		"/v1/reference-architecture/valb/bundles",
		"/v1/reference-architecture/valb/artifact-manifests",
		"/v1/reference-architecture/valb/readiness-checks",
		"/v1/reference-architecture/valb/validation-hooks",
		"/v1/reference-architecture/valb/conformance-kit",
		"/v1/reference-architecture/valb/deviations",
		"/v1/reference-architecture/valb/proofs",
	}
}

func EvaluateReferenceArchitectureValBProofsState(valBState, point6State string, supportedFamilies, surfaceRefs, evidenceRefs, limitations []string, projectionDisclaimer string) string {
	baseState := strings.TrimSpace(valBState)
	if !containsExactTrimmedStringSet(supportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(surfaceRefs, referenceArchitectureValBProofSurfaceRefs()...) ||
		len(evidenceRefs) < 10 ||
		len(limitations) == 0 ||
		!referenceArchitectureValBHasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == ReferenceArchitectureValBStateActive {
			return ReferenceArchitectureValBStatePartial
		}
		return baseState
	}
	if baseState == ReferenceArchitectureValBStateActive && strings.TrimSpace(point6State) != ReferenceArchitecturePoint6StateNotComplete {
		return ReferenceArchitectureValBStatePartial
	}
	return baseState
}
