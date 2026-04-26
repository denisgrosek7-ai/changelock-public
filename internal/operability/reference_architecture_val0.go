package operability

import (
	"strings"
	"time"
)

const (
	ReferenceArchitectureVal0BlueprintDisciplineStateActive     = "reference_architecture_val0_blueprint_discipline_active"
	ReferenceArchitectureVal0BlueprintDisciplineStatePartial    = "reference_architecture_val0_blueprint_discipline_partial"
	ReferenceArchitectureVal0BlueprintDisciplineStateIncomplete = "reference_architecture_val0_blueprint_discipline_incomplete"

	ReferenceArchitectureVal0TaxonomyStateActive     = "reference_architecture_val0_taxonomy_active"
	ReferenceArchitectureVal0TaxonomyStatePartial    = "reference_architecture_val0_taxonomy_partial"
	ReferenceArchitectureVal0TaxonomyStateIncomplete = "reference_architecture_val0_taxonomy_incomplete"

	ReferenceArchitectureVal0EnvironmentFitStateActive     = "reference_architecture_val0_environment_fit_active"
	ReferenceArchitectureVal0EnvironmentFitStatePartial    = "reference_architecture_val0_environment_fit_partial"
	ReferenceArchitectureVal0EnvironmentFitStateIncomplete = "reference_architecture_val0_environment_fit_incomplete"

	ReferenceArchitectureVal0EvidenceDisciplineStateActive     = "reference_architecture_val0_evidence_discipline_active"
	ReferenceArchitectureVal0EvidenceDisciplineStatePartial    = "reference_architecture_val0_evidence_discipline_partial"
	ReferenceArchitectureVal0EvidenceDisciplineStateIncomplete = "reference_architecture_val0_evidence_discipline_incomplete"

	ReferenceArchitectureVal0CompatibilityBaselineStateActive     = "reference_architecture_val0_compatibility_baseline_active"
	ReferenceArchitectureVal0CompatibilityBaselineStatePartial    = "reference_architecture_val0_compatibility_baseline_partial"
	ReferenceArchitectureVal0CompatibilityBaselineStateIncomplete = "reference_architecture_val0_compatibility_baseline_incomplete"

	ReferenceArchitectureVal0StateIncomplete  = "reference_architecture_val0_incomplete"
	ReferenceArchitectureVal0StateSubstantial = "reference_architecture_val0_substantially_ready"
	ReferenceArchitectureVal0StateActive      = "reference_architecture_val0_active"

	ReferenceArchitecturePoint6StatePass        = "reference_architecture_point_6_pass"
	ReferenceArchitecturePoint6StateNotComplete = "reference_architecture_point_6_not_complete"

	ReferenceArchitectureFamilyEnterpriseDefault     = "enterprise_default"
	ReferenceArchitectureFamilyHighAssurance         = "high_assurance"
	ReferenceArchitectureFamilyRegulatedPrivacyFirst = "regulated_privacy_first"
	ReferenceArchitectureFamilySovereignAirGapped    = "sovereign_air_gapped"
	ReferenceArchitectureFamilyPerformanceSensitive  = "performance_sensitive"
	ReferenceArchitectureFamilyPartnerMSPSuitable    = "partner_msp_suitable"

	ReferenceArchitectureLifecycleActive      = "active"
	ReferenceArchitectureLifecycleDeprecated  = "deprecated"
	ReferenceArchitectureLifecycleSuperseded  = "superseded"
	ReferenceArchitectureLifecycleUnsupported = "unsupported"
	ReferenceArchitectureLifecycleUnknown     = "unknown"

	ReferenceArchitectureCompatibilityCompatible            = "compatible"
	ReferenceArchitectureCompatibilityCompatibleWithWarning = "compatible_with_warnings"
	ReferenceArchitectureCompatibilityDeprecated            = "deprecated"
	ReferenceArchitectureCompatibilitySuperseded            = "superseded"
	ReferenceArchitectureCompatibilityUnsupported           = "unsupported"
	ReferenceArchitectureCompatibilityUnknown               = "unknown"

	ReferenceArchitectureConformanceMatched          = "matched"
	ReferenceArchitectureConformancePartiallyMatched = "partially_matched"
	ReferenceArchitectureConformanceDegraded         = "degraded"
	ReferenceArchitectureConformanceUnsupported      = "unsupported"
	ReferenceArchitectureConformanceDrifted          = "drifted"
	ReferenceArchitectureConformanceSupersededRef    = "superseded_reference"
	ReferenceArchitectureConformanceUnknown          = "unknown"

	ReferenceArchitectureTopologySingleRegion      = "single_region_cluster"
	ReferenceArchitectureTopologyMultiRegion       = "multi_region_service_mesh"
	ReferenceArchitectureTopologyAirGappedCell     = "air_gapped_cell"
	ReferenceArchitectureTopologyPartnerIsolated   = "partner_managed_isolated_stack"
	ReferenceArchitectureTrustCentralized          = "centralized_signing"
	ReferenceArchitectureTrustHSMBacked            = "hsm_backed_signing"
	ReferenceArchitectureTrustOfflineIntermediates = "offline_root_with_online_intermediates"
	ReferenceArchitectureTrustAirGappedOfflineRoot = "air_gapped_offline_root"
	ReferenceArchitectureAuditCentralized          = "centralized_audit_writer"
	ReferenceArchitectureAuditRegional             = "regional_audit_writer"
	ReferenceArchitectureAuditDeferredAirGap       = "deferred_air_gap_transfer"
	ReferenceArchitectureConnectivityConnected     = "connected"
	ReferenceArchitectureConnectivityRestricted    = "restricted_egress"
	ReferenceArchitectureConnectivityAirGapped     = "air_gapped"
	ReferenceArchitectureResidencyRegional         = "regional"
	ReferenceArchitectureResidencyMultiRegion      = "multi_region"
	ReferenceArchitectureResidencySovereignLocal   = "sovereign_local"
	ReferenceArchitectureOperatorCustomer          = "customer_operated"
	ReferenceArchitectureOperatorPartner           = "partner_operated"
	ReferenceArchitectureOperatorShared            = "shared_responsibility"
	ReferenceArchitectureAccessDirectVerifier      = "direct_verifier_access"
	ReferenceArchitectureAccessBrokeredPartner     = "brokered_partner_access"
	ReferenceArchitectureAccessOfflineEvidence     = "offline_evidence_exchange"

	ReferenceArchitectureCapabilitySigning         = "signing_capability"
	ReferenceArchitectureCapabilityAuditWriter     = "audit_writer_capability"
	ReferenceArchitectureCapabilityEvidenceStorage = "evidence_storage_capability"
	ReferenceArchitectureCapabilityPolicyDist      = "policy_distribution_capability"
	ReferenceArchitectureCapabilityRecovery        = "recovery_capability"
	ReferenceArchitectureCapabilityVerifierAccess  = "verifier_access_capability"
	ReferenceArchitectureCapabilityAirGapTransfer  = "air_gap_transfer_capability"

	ReferenceArchitectureEvidenceDeploymentObservation = "deployment_observation"
	ReferenceArchitectureEvidenceCapabilityAttestation = "capability_attestation"
	ReferenceArchitectureEvidenceAuditSnapshot         = "audit_snapshot"
	ReferenceArchitectureEvidenceCompatibilityReport   = "compatibility_report"
	ReferenceArchitectureEvidenceSupportBoundary       = "support_boundary_record"

	ReferenceArchitectureDegradedCapabilityGap     = "capability_gap"
	ReferenceArchitectureDegradedAuditPathReduced  = "audit_path_reduced"
	ReferenceArchitectureDegradedSupportLimited    = "support_boundary_limited"
	ReferenceArchitectureDegradedCompatibilityWarn = "compatibility_warning"
	ReferenceArchitectureDegradedEvidenceCaveated  = "evidence_caveated"
	ReferenceArchitectureDegradedTopologyVariance  = "topology_variance"

	ReferenceArchitectureUnsupportedTrustMismatch   = "trust_anchor_incompatible"
	ReferenceArchitectureUnsupportedAirGapTransfer  = "air_gap_transfer_missing"
	ReferenceArchitectureUnsupportedConnectivity    = "unsupported_connectivity_mode"
	ReferenceArchitectureUnsupportedEvidenceMissing = "evidence_missing"
	ReferenceArchitectureUnsupportedLifecycle       = "lifecycle_unsupported"
	ReferenceArchitectureUnsupportedUnknownEnv      = "unknown_environment"
)

type ReferenceArchitectureEnvironmentProfile struct {
	DeploymentTopology          string `json:"deployment_topology"`
	TrustAnchorMode             string `json:"trust_anchor_mode"`
	AuditPathMode               string `json:"audit_path_mode"`
	ConnectivityMode            string `json:"connectivity_mode"`
	DataResidencyMode           string `json:"data_residency_mode"`
	OperatorControlModel        string `json:"operator_control_model"`
	VerifierOrPartnerAccessMode string `json:"verifier_or_partner_access_mode"`
}

type ReferenceArchitectureEvidenceReference struct {
	EvidenceID     string   `json:"evidence_id"`
	EvidenceType   string   `json:"evidence_type"`
	Source         string   `json:"source"`
	Timestamp      string   `json:"timestamp"`
	FreshnessState string   `json:"freshness_state"`
	Scope          string   `json:"scope"`
	Caveats        []string `json:"caveats,omitempty"`
}

type ReferenceArchitectureBlueprintContract struct {
	CurrentState                   string                                   `json:"current_state"`
	BlueprintID                    string                                   `json:"blueprint_id"`
	Version                        string                                   `json:"version"`
	Family                         string                                   `json:"family"`
	LifecycleState                 string                                   `json:"lifecycle_state"`
	Owner                          string                                   `json:"owner"`
	CreatedAt                      string                                   `json:"created_at"`
	UpdatedAt                      string                                   `json:"updated_at"`
	SupportedFamilies              []string                                 `json:"supported_families,omitempty"`
	SupportedLifecycleStates       []string                                 `json:"supported_lifecycle_states,omitempty"`
	SupportedCompatibilityStates   []string                                 `json:"supported_compatibility_states,omitempty"`
	SupportedConformanceStates     []string                                 `json:"supported_conformance_states,omitempty"`
	TargetEnvironment              ReferenceArchitectureEnvironmentProfile  `json:"target_environment"`
	ObservedEnvironment            ReferenceArchitectureEnvironmentProfile  `json:"observed_environment"`
	RequiredCapabilities           []string                                 `json:"required_capabilities,omitempty"`
	OptionalCapabilities           []string                                 `json:"optional_capabilities,omitempty"`
	ObservedCapabilities           []string                                 `json:"observed_capabilities,omitempty"`
	InfrastructureAssumptions      []string                                 `json:"infrastructure_assumptions,omitempty"`
	NetworkAssumptions             []string                                 `json:"network_assumptions,omitempty"`
	IdentityAccessAssumptions      []string                                 `json:"identity_access_assumptions,omitempty"`
	TrustCustodyAssumptions        []string                                 `json:"trust_custody_assumptions,omitempty"`
	StorageAssumptions             []string                                 `json:"storage_assumptions,omitempty"`
	OperationalAssumptions         []string                                 `json:"operational_assumptions,omitempty"`
	SupportAssumptions             []string                                 `json:"support_assumptions,omitempty"`
	SupportedDegradedConditions    []string                                 `json:"supported_degraded_conditions,omitempty"`
	SupportedUnsupportedConditions []string                                 `json:"supported_unsupported_conditions,omitempty"`
	TriggeredDegradedConditions    []string                                 `json:"triggered_degraded_conditions,omitempty"`
	TriggeredUnsupportedConditions []string                                 `json:"triggered_unsupported_conditions,omitempty"`
	MissingCapabilities            []string                                 `json:"missing_capabilities,omitempty"`
	CaveatedItems                  []string                                 `json:"caveated_items,omitempty"`
	DegradedReasons                []string                                 `json:"degraded_reasons,omitempty"`
	UnsupportedReasons             []string                                 `json:"unsupported_reasons,omitempty"`
	CompatibilityState             string                                   `json:"compatibility_state"`
	ConformanceState               string                                   `json:"conformance_state"`
	SupportBoundaryRef             string                                   `json:"support_boundary_ref"`
	EvidenceRefs                   []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	RedactionKeepsCaveats          bool                                     `json:"redaction_keeps_caveats"`
	CertifiedLanguagePresent       bool                                     `json:"certified_architecture_language_present"`
	ProjectionDisclaimer           string                                   `json:"projection_disclaimer"`
}

func referenceArchitectureVal0Families() []string {
	return []string{
		ReferenceArchitectureFamilyEnterpriseDefault,
		ReferenceArchitectureFamilyHighAssurance,
		ReferenceArchitectureFamilyRegulatedPrivacyFirst,
		ReferenceArchitectureFamilySovereignAirGapped,
		ReferenceArchitectureFamilyPerformanceSensitive,
		ReferenceArchitectureFamilyPartnerMSPSuitable,
	}
}

func referenceArchitectureVal0LifecycleStates() []string {
	return []string{
		ReferenceArchitectureLifecycleActive,
		ReferenceArchitectureLifecycleDeprecated,
		ReferenceArchitectureLifecycleSuperseded,
		ReferenceArchitectureLifecycleUnsupported,
		ReferenceArchitectureLifecycleUnknown,
	}
}

func referenceArchitectureVal0CompatibilityStates() []string {
	return []string{
		ReferenceArchitectureCompatibilityCompatible,
		ReferenceArchitectureCompatibilityCompatibleWithWarning,
		ReferenceArchitectureCompatibilityDeprecated,
		ReferenceArchitectureCompatibilitySuperseded,
		ReferenceArchitectureCompatibilityUnsupported,
		ReferenceArchitectureCompatibilityUnknown,
	}
}

func referenceArchitectureVal0ConformanceStates() []string {
	return []string{
		ReferenceArchitectureConformanceMatched,
		ReferenceArchitectureConformancePartiallyMatched,
		ReferenceArchitectureConformanceDegraded,
		ReferenceArchitectureConformanceUnsupported,
		ReferenceArchitectureConformanceDrifted,
		ReferenceArchitectureConformanceSupersededRef,
		ReferenceArchitectureConformanceUnknown,
	}
}

func referenceArchitectureVal0SupportedCapabilities() []string {
	return []string{
		ReferenceArchitectureCapabilitySigning,
		ReferenceArchitectureCapabilityAuditWriter,
		ReferenceArchitectureCapabilityEvidenceStorage,
		ReferenceArchitectureCapabilityPolicyDist,
		ReferenceArchitectureCapabilityRecovery,
		ReferenceArchitectureCapabilityVerifierAccess,
		ReferenceArchitectureCapabilityAirGapTransfer,
	}
}

func referenceArchitectureVal0SupportedEvidenceTypes() []string {
	return []string{
		ReferenceArchitectureEvidenceDeploymentObservation,
		ReferenceArchitectureEvidenceCapabilityAttestation,
		ReferenceArchitectureEvidenceAuditSnapshot,
		ReferenceArchitectureEvidenceCompatibilityReport,
		ReferenceArchitectureEvidenceSupportBoundary,
	}
}

func referenceArchitectureVal0SupportedTopologies() []string {
	return []string{
		ReferenceArchitectureTopologySingleRegion,
		ReferenceArchitectureTopologyMultiRegion,
		ReferenceArchitectureTopologyAirGappedCell,
		ReferenceArchitectureTopologyPartnerIsolated,
	}
}

func referenceArchitectureVal0SupportedTrustAnchors() []string {
	return []string{
		ReferenceArchitectureTrustCentralized,
		ReferenceArchitectureTrustHSMBacked,
		ReferenceArchitectureTrustOfflineIntermediates,
		ReferenceArchitectureTrustAirGappedOfflineRoot,
	}
}

func referenceArchitectureVal0SupportedAuditPaths() []string {
	return []string{
		ReferenceArchitectureAuditCentralized,
		ReferenceArchitectureAuditRegional,
		ReferenceArchitectureAuditDeferredAirGap,
	}
}

func referenceArchitectureVal0SupportedConnectivityModes() []string {
	return []string{
		ReferenceArchitectureConnectivityConnected,
		ReferenceArchitectureConnectivityRestricted,
		ReferenceArchitectureConnectivityAirGapped,
	}
}

func referenceArchitectureVal0SupportedResidencyModes() []string {
	return []string{
		ReferenceArchitectureResidencyRegional,
		ReferenceArchitectureResidencyMultiRegion,
		ReferenceArchitectureResidencySovereignLocal,
	}
}

func referenceArchitectureVal0SupportedOperatorModels() []string {
	return []string{
		ReferenceArchitectureOperatorCustomer,
		ReferenceArchitectureOperatorPartner,
		ReferenceArchitectureOperatorShared,
	}
}

func referenceArchitectureVal0SupportedAccessModes() []string {
	return []string{
		ReferenceArchitectureAccessDirectVerifier,
		ReferenceArchitectureAccessBrokeredPartner,
		ReferenceArchitectureAccessOfflineEvidence,
	}
}

func referenceArchitectureVal0SupportedDegradedConditions() []string {
	return []string{
		ReferenceArchitectureDegradedCapabilityGap,
		ReferenceArchitectureDegradedAuditPathReduced,
		ReferenceArchitectureDegradedSupportLimited,
		ReferenceArchitectureDegradedCompatibilityWarn,
		ReferenceArchitectureDegradedEvidenceCaveated,
		ReferenceArchitectureDegradedTopologyVariance,
	}
}

func referenceArchitectureVal0SupportedUnsupportedConditions() []string {
	return []string{
		ReferenceArchitectureUnsupportedTrustMismatch,
		ReferenceArchitectureUnsupportedAirGapTransfer,
		ReferenceArchitectureUnsupportedConnectivity,
		ReferenceArchitectureUnsupportedEvidenceMissing,
		ReferenceArchitectureUnsupportedLifecycle,
		ReferenceArchitectureUnsupportedUnknownEnv,
	}
}

func referenceArchitectureVal0CriticalCapabilities() []string {
	return []string{
		ReferenceArchitectureCapabilitySigning,
		ReferenceArchitectureCapabilityAuditWriter,
		ReferenceArchitectureCapabilityEvidenceStorage,
		ReferenceArchitectureCapabilityRecovery,
	}
}

func ReferenceArchitectureVal0BlueprintContract() ReferenceArchitectureBlueprintContract {
	return ReferenceArchitectureBlueprintContract{
		CurrentState:                 "reference_architecture_val0_blueprint_ready",
		BlueprintID:                  "reference-blueprint-val0-001",
		Version:                      "0.1.0",
		Family:                       ReferenceArchitectureFamilyEnterpriseDefault,
		LifecycleState:               ReferenceArchitectureLifecycleActive,
		Owner:                        "reference_architecture_program",
		CreatedAt:                    "2026-04-26T07:00:00Z",
		UpdatedAt:                    "2026-04-26T08:00:00Z",
		SupportedFamilies:            referenceArchitectureVal0Families(),
		SupportedLifecycleStates:     referenceArchitectureVal0LifecycleStates(),
		SupportedCompatibilityStates: referenceArchitectureVal0CompatibilityStates(),
		SupportedConformanceStates:   referenceArchitectureVal0ConformanceStates(),
		TargetEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologyMultiRegion,
			TrustAnchorMode:             ReferenceArchitectureTrustHSMBacked,
			AuditPathMode:               ReferenceArchitectureAuditRegional,
			ConnectivityMode:            ReferenceArchitectureConnectivityRestricted,
			DataResidencyMode:           ReferenceArchitectureResidencyRegional,
			OperatorControlModel:        ReferenceArchitectureOperatorCustomer,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessBrokeredPartner,
		},
		ObservedEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologyMultiRegion,
			TrustAnchorMode:             ReferenceArchitectureTrustHSMBacked,
			AuditPathMode:               ReferenceArchitectureAuditRegional,
			ConnectivityMode:            ReferenceArchitectureConnectivityRestricted,
			DataResidencyMode:           ReferenceArchitectureResidencyRegional,
			OperatorControlModel:        ReferenceArchitectureOperatorCustomer,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessBrokeredPartner,
		},
		RequiredCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
			ReferenceArchitectureCapabilityVerifierAccess,
		},
		OptionalCapabilities: []string{
			ReferenceArchitectureCapabilityAirGapTransfer,
		},
		ObservedCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
			ReferenceArchitectureCapabilityVerifierAccess,
		},
		InfrastructureAssumptions:      []string{"regional kubernetes substrate with signer isolation and durable control-plane quorum"},
		NetworkAssumptions:             []string{"restricted egress is enforced and audit path replication can reach the declared regional audit writer"},
		IdentityAccessAssumptions:      []string{"operator identities are federated and scoped by tenant and environment boundary"},
		TrustCustodyAssumptions:        []string{"hsm backed trust anchors remain under declared customer custody with governed intermediate rotation"},
		StorageAssumptions:             []string{"evidence storage remains durable, append-oriented, and recoverable within declared retention window"},
		OperationalAssumptions:         []string{"sre runbooks and recovery drills remain maintained for signer, audit writer, and policy distribution dependencies"},
		SupportAssumptions:             []string{"declared support boundary remains staffed and measured for the target topology and verifier access model"},
		SupportedDegradedConditions:    referenceArchitectureVal0SupportedDegradedConditions(),
		SupportedUnsupportedConditions: referenceArchitectureVal0SupportedUnsupportedConditions(),
		CompatibilityState:             ReferenceArchitectureCompatibilityCompatible,
		ConformanceState:               ReferenceArchitectureConformanceMatched,
		SupportBoundaryRef:             "support-boundary:enterprise-default",
		EvidenceRefs: []ReferenceArchitectureEvidenceReference{
			{
				EvidenceID:     "evidence:blueprint-topology-001",
				EvidenceType:   ReferenceArchitectureEvidenceDeploymentObservation,
				Source:         "inventory/topology",
				Timestamp:      "2026-04-26T08:00:00Z",
				FreshnessState: IntelligenceCalibrationFreshnessFresh,
				Scope:          "environment_fit",
				Caveats:        []string{"bounded to declared prod topology and target operating envelope"},
			},
			{
				EvidenceID:     "evidence:blueprint-capabilities-001",
				EvidenceType:   ReferenceArchitectureEvidenceCapabilityAttestation,
				Source:         "control-plane/capabilities",
				Timestamp:      "2026-04-26T08:05:00Z",
				FreshnessState: IntelligenceCalibrationFreshnessFresh,
				Scope:          "capability_conformance",
				Caveats:        []string{"bounded to capabilities declared in the current environment snapshot"},
			},
			{
				EvidenceID:     "evidence:blueprint-audit-001",
				EvidenceType:   ReferenceArchitectureEvidenceAuditSnapshot,
				Source:         "audit-writer/regional",
				Timestamp:      "2026-04-26T08:07:00Z",
				FreshnessState: IntelligenceCalibrationFreshnessFresh,
				Scope:          "audit_path",
				Caveats:        []string{"bounded to the declared audit path mode and current regional replication posture"},
			},
			{
				EvidenceID:     "evidence:blueprint-compatibility-001",
				EvidenceType:   ReferenceArchitectureEvidenceCompatibilityReport,
				Source:         "compatibility/reports",
				Timestamp:      "2026-04-26T08:10:00Z",
				FreshnessState: IntelligenceCalibrationFreshnessFresh,
				Scope:          "compatibility_baseline",
				Caveats:        []string{"bounded to current supported component set and declared lifecycle posture"},
			},
			{
				EvidenceID:     "evidence:blueprint-support-001",
				EvidenceType:   ReferenceArchitectureEvidenceSupportBoundary,
				Source:         "support/catalog",
				Timestamp:      "2026-04-26T08:11:00Z",
				FreshnessState: IntelligenceCalibrationFreshnessFresh,
				Scope:          "support_boundary",
				Caveats:        []string{"bounded to named support boundary and declared operator control model"},
			},
		},
		RedactionKeepsCaveats:    true,
		CertifiedLanguagePresent: false,
		ProjectionDisclaimer:     "projection_only not_canonical_truth validated_reference_blueprint measured_conformance",
	}
}

func referenceArchitectureVal0ParseTimestamp(value string) (time.Time, bool) {
	parsed, err := time.Parse(time.RFC3339, strings.TrimSpace(value))
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}

func referenceArchitectureVal0EnvironmentSupported(profile ReferenceArchitectureEnvironmentProfile) bool {
	return containsTrimmedString(referenceArchitectureVal0SupportedTopologies(), profile.DeploymentTopology) &&
		containsTrimmedString(referenceArchitectureVal0SupportedTrustAnchors(), profile.TrustAnchorMode) &&
		containsTrimmedString(referenceArchitectureVal0SupportedAuditPaths(), profile.AuditPathMode) &&
		containsTrimmedString(referenceArchitectureVal0SupportedConnectivityModes(), profile.ConnectivityMode) &&
		containsTrimmedString(referenceArchitectureVal0SupportedResidencyModes(), profile.DataResidencyMode) &&
		containsTrimmedString(referenceArchitectureVal0SupportedOperatorModels(), profile.OperatorControlModel) &&
		containsTrimmedString(referenceArchitectureVal0SupportedAccessModes(), profile.VerifierOrPartnerAccessMode)
}

func referenceArchitectureVal0AllAssumptionsPresent(model ReferenceArchitectureBlueprintContract) bool {
	return len(model.InfrastructureAssumptions) > 0 &&
		len(model.NetworkAssumptions) > 0 &&
		len(model.IdentityAccessAssumptions) > 0 &&
		len(model.TrustCustodyAssumptions) > 0 &&
		len(model.StorageAssumptions) > 0 &&
		len(model.OperationalAssumptions) > 0 &&
		len(model.SupportAssumptions) > 0
}

func referenceArchitectureVal0MissingRequiredCapabilities(model ReferenceArchitectureBlueprintContract) []string {
	seen := make(map[string]struct{}, len(model.ObservedCapabilities))
	for _, capability := range model.ObservedCapabilities {
		if containsTrimmedString(referenceArchitectureVal0SupportedCapabilities(), capability) {
			seen[strings.TrimSpace(capability)] = struct{}{}
		}
	}
	missing := make([]string, 0, len(model.RequiredCapabilities))
	for _, capability := range model.RequiredCapabilities {
		trimmed := strings.TrimSpace(capability)
		if trimmed == "" || !containsTrimmedString(referenceArchitectureVal0SupportedCapabilities(), trimmed) {
			missing = append(missing, trimmed)
			continue
		}
		if _, ok := seen[trimmed]; !ok {
			missing = append(missing, trimmed)
		}
	}
	for _, declared := range model.MissingCapabilities {
		trimmed := strings.TrimSpace(declared)
		if trimmed != "" && !containsTrimmedString(missing, trimmed) {
			missing = append(missing, trimmed)
		}
	}
	return missing
}

func referenceArchitectureVal0HasCriticalMissingCapabilities(missing []string) bool {
	for _, capability := range missing {
		if containsTrimmedString(referenceArchitectureVal0CriticalCapabilities(), capability) {
			return true
		}
	}
	return false
}

func referenceArchitectureVal0EvidenceValid(model ReferenceArchitectureBlueprintContract) (allFresh bool, stale bool, ok bool) {
	if len(model.EvidenceRefs) == 0 {
		return false, false, false
	}
	allFresh = true
	for _, evidence := range model.EvidenceRefs {
		if strings.TrimSpace(evidence.EvidenceID) == "" ||
			strings.TrimSpace(evidence.Source) == "" ||
			strings.TrimSpace(evidence.Scope) == "" ||
			len(evidence.Caveats) == 0 ||
			!containsTrimmedString(referenceArchitectureVal0SupportedEvidenceTypes(), evidence.EvidenceType) ||
			!containsTrimmedString([]string{
				IntelligenceCalibrationFreshnessFresh,
				IntelligenceCalibrationFreshnessStale,
				IntelligenceCalibrationFreshnessExpired,
				IntelligenceCalibrationFreshnessUnknown,
				IntelligenceCalibrationFreshnessUnsupported,
			}, evidence.FreshnessState) {
			return false, false, false
		}
		if _, parsed := referenceArchitectureVal0ParseTimestamp(evidence.Timestamp); !parsed {
			return false, false, false
		}
		switch strings.TrimSpace(evidence.FreshnessState) {
		case IntelligenceCalibrationFreshnessFresh:
		case IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessExpired:
			allFresh = false
			stale = true
		default:
			return false, false, false
		}
	}
	return allFresh, stale, true
}

func referenceArchitectureVal0EnvironmentMismatch(model ReferenceArchitectureBlueprintContract) (trustMismatch bool, auditMismatch bool, connectivityMismatch bool, supportMissing bool, topologyMismatch bool) {
	trustMismatch = strings.TrimSpace(model.TargetEnvironment.TrustAnchorMode) != strings.TrimSpace(model.ObservedEnvironment.TrustAnchorMode)
	auditMismatch = strings.TrimSpace(model.TargetEnvironment.AuditPathMode) != strings.TrimSpace(model.ObservedEnvironment.AuditPathMode)
	connectivityMismatch = strings.TrimSpace(model.TargetEnvironment.ConnectivityMode) != strings.TrimSpace(model.ObservedEnvironment.ConnectivityMode)
	supportMissing = strings.TrimSpace(model.SupportBoundaryRef) == ""
	topologyMismatch = strings.TrimSpace(model.TargetEnvironment.DeploymentTopology) != strings.TrimSpace(model.ObservedEnvironment.DeploymentTopology)
	return trustMismatch, auditMismatch, connectivityMismatch, supportMissing, topologyMismatch
}

func referenceArchitectureVal0EnvironmentUnsupported(model ReferenceArchitectureBlueprintContract, missingCapabilities []string) bool {
	if containsTrimmedString(model.TriggeredUnsupportedConditions, ReferenceArchitectureUnsupportedTrustMismatch) ||
		containsTrimmedString(model.TriggeredUnsupportedConditions, ReferenceArchitectureUnsupportedConnectivity) ||
		containsTrimmedString(model.TriggeredUnsupportedConditions, ReferenceArchitectureUnsupportedAirGapTransfer) ||
		len(model.UnsupportedReasons) > 0 {
		return true
	}
	if referenceArchitectureVal0HasCriticalMissingCapabilities(missingCapabilities) {
		return true
	}
	if strings.TrimSpace(model.TargetEnvironment.ConnectivityMode) == ReferenceArchitectureConnectivityAirGapped &&
		(strings.TrimSpace(model.ObservedEnvironment.ConnectivityMode) != ReferenceArchitectureConnectivityAirGapped ||
			!containsTrimmedString(model.ObservedCapabilities, ReferenceArchitectureCapabilityAirGapTransfer)) {
		return true
	}
	return false
}

func EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model ReferenceArchitectureBlueprintContract) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.BlueprintID) == "" ||
		strings.TrimSpace(model.Version) == "" ||
		strings.TrimSpace(model.Family) == "" ||
		strings.TrimSpace(model.LifecycleState) == "" ||
		strings.TrimSpace(model.Owner) == "" ||
		strings.TrimSpace(model.CreatedAt) == "" ||
		strings.TrimSpace(model.UpdatedAt) == "" ||
		strings.TrimSpace(model.CompatibilityState) == "" ||
		strings.TrimSpace(model.ConformanceState) == "" ||
		strings.TrimSpace(model.ProjectionDisclaimer) == "" ||
		len(model.RequiredCapabilities) == 0 {
		return ReferenceArchitectureVal0BlueprintDisciplineStateIncomplete
	}
	if !referenceArchitectureVal0AllAssumptionsPresent(model) ||
		!containsExactTrimmedStringSet(model.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(model.SupportedLifecycleStates, referenceArchitectureVal0LifecycleStates()...) ||
		!containsExactTrimmedStringSet(model.SupportedCompatibilityStates, referenceArchitectureVal0CompatibilityStates()...) ||
		!containsExactTrimmedStringSet(model.SupportedConformanceStates, referenceArchitectureVal0ConformanceStates()...) ||
		!containsExactTrimmedStringSet(model.SupportedDegradedConditions, referenceArchitectureVal0SupportedDegradedConditions()...) ||
		!containsExactTrimmedStringSet(model.SupportedUnsupportedConditions, referenceArchitectureVal0SupportedUnsupportedConditions()...) ||
		!containsTrimmedString(referenceArchitectureVal0Families(), model.Family) ||
		!containsTrimmedString(referenceArchitectureVal0LifecycleStates(), model.LifecycleState) ||
		!containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), model.CompatibilityState) ||
		!containsTrimmedString(referenceArchitectureVal0ConformanceStates(), model.ConformanceState) ||
		!referenceArchitectureVal0EnvironmentSupported(model.TargetEnvironment) ||
		!referenceArchitectureVal0EnvironmentSupported(model.ObservedEnvironment) ||
		!containsAllTrimmedStrings(referenceArchitectureVal0SupportedCapabilities(), model.RequiredCapabilities...) ||
		!containsAllTrimmedStrings(referenceArchitectureVal0SupportedCapabilities(), model.OptionalCapabilities...) ||
		model.CertifiedLanguagePresent ||
		!model.RedactionKeepsCaveats ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ReferenceArchitectureVal0BlueprintDisciplineStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(model.CreatedAt); !ok {
		return ReferenceArchitectureVal0BlueprintDisciplineStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(model.UpdatedAt); !ok {
		return ReferenceArchitectureVal0BlueprintDisciplineStatePartial
	}
	return ReferenceArchitectureVal0BlueprintDisciplineStateActive
}

func EvaluateReferenceArchitectureVal0TaxonomyState(model ReferenceArchitectureBlueprintContract) string {
	if len(model.SupportedFamilies) == 0 || len(model.SupportedLifecycleStates) == 0 || len(model.SupportedCompatibilityStates) == 0 || len(model.SupportedConformanceStates) == 0 {
		return ReferenceArchitectureVal0TaxonomyStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(model.SupportedLifecycleStates, referenceArchitectureVal0LifecycleStates()...) ||
		!containsExactTrimmedStringSet(model.SupportedCompatibilityStates, referenceArchitectureVal0CompatibilityStates()...) ||
		!containsExactTrimmedStringSet(model.SupportedConformanceStates, referenceArchitectureVal0ConformanceStates()...) {
		return ReferenceArchitectureVal0TaxonomyStatePartial
	}
	return ReferenceArchitectureVal0TaxonomyStateActive
}

func EvaluateReferenceArchitectureVal0EnvironmentFitState(model ReferenceArchitectureBlueprintContract) string {
	if strings.TrimSpace(model.TargetEnvironment.DeploymentTopology) == "" ||
		strings.TrimSpace(model.TargetEnvironment.TrustAnchorMode) == "" ||
		strings.TrimSpace(model.TargetEnvironment.AuditPathMode) == "" ||
		strings.TrimSpace(model.TargetEnvironment.ConnectivityMode) == "" ||
		strings.TrimSpace(model.ObservedEnvironment.DeploymentTopology) == "" ||
		strings.TrimSpace(model.ObservedEnvironment.TrustAnchorMode) == "" ||
		strings.TrimSpace(model.ObservedEnvironment.AuditPathMode) == "" ||
		strings.TrimSpace(model.ObservedEnvironment.ConnectivityMode) == "" {
		return ReferenceArchitectureVal0EnvironmentFitStateIncomplete
	}
	if !referenceArchitectureVal0EnvironmentSupported(model.TargetEnvironment) || !referenceArchitectureVal0EnvironmentSupported(model.ObservedEnvironment) || len(referenceArchitectureVal0MissingRequiredCapabilities(model)) > 0 {
		return ReferenceArchitectureVal0EnvironmentFitStatePartial
	}
	trustMismatch, auditMismatch, connectivityMismatch, supportMissing, topologyMismatch := referenceArchitectureVal0EnvironmentMismatch(model)
	if trustMismatch || auditMismatch || connectivityMismatch || supportMissing || topologyMismatch {
		return ReferenceArchitectureVal0EnvironmentFitStatePartial
	}
	if len(model.TriggeredDegradedConditions) > 0 || len(model.TriggeredUnsupportedConditions) > 0 || len(model.DegradedReasons) > 0 || len(model.UnsupportedReasons) > 0 {
		return ReferenceArchitectureVal0EnvironmentFitStatePartial
	}
	return ReferenceArchitectureVal0EnvironmentFitStateActive
}

func EvaluateReferenceArchitectureVal0EvidenceDisciplineState(model ReferenceArchitectureBlueprintContract) string {
	if len(model.EvidenceRefs) == 0 {
		return ReferenceArchitectureVal0EvidenceDisciplineStateIncomplete
	}
	allFresh, stale, ok := referenceArchitectureVal0EvidenceValid(model)
	if !ok {
		return ReferenceArchitectureVal0EvidenceDisciplineStatePartial
	}
	if !allFresh || stale {
		return ReferenceArchitectureVal0EvidenceDisciplineStatePartial
	}
	return ReferenceArchitectureVal0EvidenceDisciplineStateActive
}

func EvaluateReferenceArchitectureVal0CompatibilityBaselineState(model ReferenceArchitectureBlueprintContract) string {
	if strings.TrimSpace(model.LifecycleState) == "" || strings.TrimSpace(model.CompatibilityState) == "" {
		return ReferenceArchitectureVal0CompatibilityBaselineStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0LifecycleStates(), model.LifecycleState) || !containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), model.CompatibilityState) {
		return ReferenceArchitectureVal0CompatibilityBaselineStatePartial
	}
	if model.LifecycleState != ReferenceArchitectureLifecycleActive || model.CompatibilityState != ReferenceArchitectureCompatibilityCompatible {
		return ReferenceArchitectureVal0CompatibilityBaselineStatePartial
	}
	return ReferenceArchitectureVal0CompatibilityBaselineStateActive
}

func referenceArchitectureVal0LifecycleConformanceOverride(lifecycleState string) (string, bool) {
	switch strings.TrimSpace(lifecycleState) {
	case ReferenceArchitectureLifecycleSuperseded:
		return ReferenceArchitectureConformanceSupersededRef, true
	case ReferenceArchitectureLifecycleUnsupported:
		return ReferenceArchitectureConformanceUnsupported, true
	case ReferenceArchitectureLifecycleUnknown:
		return ReferenceArchitectureConformanceUnknown, true
	}
	return "", false
}

func referenceArchitectureVal0CompatibilityConformanceOverride(compatibilityState string) (string, bool) {
	switch strings.TrimSpace(compatibilityState) {
	case ReferenceArchitectureCompatibilitySuperseded:
		return ReferenceArchitectureConformanceSupersededRef, true
	case ReferenceArchitectureCompatibilityUnsupported:
		return ReferenceArchitectureConformanceUnsupported, true
	case ReferenceArchitectureCompatibilityUnknown:
		return ReferenceArchitectureConformanceUnknown, true
	}
	return "", false
}

func EvaluateReferenceArchitectureVal0ReferenceConformanceState(model ReferenceArchitectureBlueprintContract) string {
	if EvaluateReferenceArchitectureVal0BlueprintDisciplineState(model) != ReferenceArchitectureVal0BlueprintDisciplineStateActive {
		return ReferenceArchitectureConformanceUnknown
	}
	if strings.TrimSpace(model.ConformanceState) == "" || !containsTrimmedString(referenceArchitectureVal0ConformanceStates(), model.ConformanceState) {
		return ReferenceArchitectureConformanceUnknown
	}
	if model.CertifiedLanguagePresent || !model.RedactionKeepsCaveats {
		return ReferenceArchitectureConformanceUnknown
	}

	allFresh, staleEvidence, evidenceOK := referenceArchitectureVal0EvidenceValid(model)
	if !evidenceOK {
		return ReferenceArchitectureConformanceUnknown
	}

	// Lifecycle and compatibility overrides are resolved before any degraded,
	// partial, or matched fallback so unsupported or superseded references
	// cannot silently fall through to a clean matched state.
	if override, ok := referenceArchitectureVal0LifecycleConformanceOverride(model.LifecycleState); ok {
		return override
	}
	if override, ok := referenceArchitectureVal0CompatibilityConformanceOverride(model.CompatibilityState); ok {
		return override
	}

	missingCapabilities := referenceArchitectureVal0MissingRequiredCapabilities(model)
	trustMismatch, auditMismatch, connectivityMismatch, supportMissing, topologyMismatch := referenceArchitectureVal0EnvironmentMismatch(model)

	if referenceArchitectureVal0EnvironmentUnsupported(model, missingCapabilities) {
		return ReferenceArchitectureConformanceUnsupported
	}
	if staleEvidence {
		return ReferenceArchitectureConformanceDrifted
	}
	if len(missingCapabilities) > 0 || trustMismatch || auditMismatch || connectivityMismatch || supportMissing || topologyMismatch || len(model.TriggeredDegradedConditions) > 0 || len(model.DegradedReasons) > 0 {
		return ReferenceArchitectureConformanceDegraded
	}
	if strings.TrimSpace(model.LifecycleState) == ReferenceArchitectureLifecycleDeprecated ||
		strings.TrimSpace(model.CompatibilityState) == ReferenceArchitectureCompatibilityCompatibleWithWarning ||
		strings.TrimSpace(model.CompatibilityState) == ReferenceArchitectureCompatibilityDeprecated ||
		len(model.CaveatedItems) > 0 ||
		!allFresh {
		return ReferenceArchitectureConformancePartiallyMatched
	}
	return ReferenceArchitectureConformanceMatched
}

func EvaluateReferenceArchitectureVal0State(point5State, blueprintDisciplineState, taxonomyState, environmentFitState, evidenceDisciplineState, compatibilityBaselineState, conformanceState string) string {
	if strings.TrimSpace(point5State) == "" ||
		strings.TrimSpace(blueprintDisciplineState) == "" ||
		strings.TrimSpace(taxonomyState) == "" ||
		strings.TrimSpace(environmentFitState) == "" ||
		strings.TrimSpace(evidenceDisciplineState) == "" ||
		strings.TrimSpace(compatibilityBaselineState) == "" ||
		strings.TrimSpace(conformanceState) == "" {
		return ReferenceArchitectureVal0StateIncomplete
	}
	if strings.TrimSpace(point5State) != IntelligenceCalibrationPoint5StatePass {
		return ReferenceArchitectureVal0StateIncomplete
	}
	if blueprintDisciplineState == ReferenceArchitectureVal0BlueprintDisciplineStateActive &&
		taxonomyState == ReferenceArchitectureVal0TaxonomyStateActive &&
		environmentFitState == ReferenceArchitectureVal0EnvironmentFitStateActive &&
		evidenceDisciplineState == ReferenceArchitectureVal0EvidenceDisciplineStateActive &&
		compatibilityBaselineState == ReferenceArchitectureVal0CompatibilityBaselineStateActive &&
		conformanceState == ReferenceArchitectureConformanceMatched {
		return ReferenceArchitectureVal0StateActive
	}
	if blueprintDisciplineState != ReferenceArchitectureVal0BlueprintDisciplineStateIncomplete &&
		taxonomyState != ReferenceArchitectureVal0TaxonomyStateIncomplete &&
		environmentFitState != ReferenceArchitectureVal0EnvironmentFitStateIncomplete &&
		evidenceDisciplineState != ReferenceArchitectureVal0EvidenceDisciplineStateIncomplete &&
		compatibilityBaselineState != ReferenceArchitectureVal0CompatibilityBaselineStateIncomplete &&
		conformanceState != ReferenceArchitectureConformanceUnknown {
		return ReferenceArchitectureVal0StateSubstantial
	}
	return ReferenceArchitectureVal0StateIncomplete
}

func EvaluateReferenceArchitectureVal0ProofsState(point5State, val0State, blueprintDisciplineState, taxonomyState, environmentFitState, evidenceDisciplineState, compatibilityBaselineState, conformanceState, point6State string, supportedFamilies, supportedConformanceStates, supportedCompatibilityStates, surfaceRefs, evidenceRefs, limitations []string, projectionDisclaimer string) string {
	baseState := EvaluateReferenceArchitectureVal0State(point5State, blueprintDisciplineState, taxonomyState, environmentFitState, evidenceDisciplineState, compatibilityBaselineState, conformanceState)
	if !containsExactTrimmedStringSet(supportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(supportedConformanceStates, referenceArchitectureVal0ConformanceStates()...) ||
		!containsExactTrimmedStringSet(supportedCompatibilityStates, referenceArchitectureVal0CompatibilityStates()...) ||
		len(surfaceRefs) < 5 ||
		len(evidenceRefs) < 6 ||
		len(limitations) == 0 ||
		!strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == ReferenceArchitectureVal0StateActive {
			return ReferenceArchitectureVal0StateSubstantial
		}
		return baseState
	}
	if strings.TrimSpace(val0State) == ReferenceArchitectureVal0StateActive && strings.TrimSpace(point6State) != ReferenceArchitecturePoint6StateNotComplete {
		return ReferenceArchitectureVal0StateSubstantial
	}
	return baseState
}
