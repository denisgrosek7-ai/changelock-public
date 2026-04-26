package operability

import "strings"

const (
	ReferenceArchitectureValAFamilyProfileStateActive     = "reference_architecture_vala_family_profile_active"
	ReferenceArchitectureValAFamilyProfileStatePartial    = "reference_architecture_vala_family_profile_partial"
	ReferenceArchitectureValAFamilyProfileStateIncomplete = "reference_architecture_vala_family_profile_incomplete"
	ReferenceArchitectureValAFamilyProfileStateBlocked    = "reference_architecture_vala_family_profile_blocked"
	ReferenceArchitectureValAFamilyProfileStateUnknown    = "reference_architecture_vala_family_profile_unknown"

	ReferenceArchitectureValAFamilyRegistryStateActive     = "reference_architecture_vala_family_registry_active"
	ReferenceArchitectureValAFamilyRegistryStatePartial    = "reference_architecture_vala_family_registry_partial"
	ReferenceArchitectureValAFamilyRegistryStateIncomplete = "reference_architecture_vala_family_registry_incomplete"
	ReferenceArchitectureValAFamilyRegistryStateBlocked    = "reference_architecture_vala_family_registry_blocked"
	ReferenceArchitectureValAFamilyRegistryStateUnknown    = "reference_architecture_vala_family_registry_unknown"

	ReferenceArchitectureValAStateActive     = "reference_architecture_vala_active"
	ReferenceArchitectureValAStatePartial    = "reference_architecture_vala_partial"
	ReferenceArchitectureValAStateIncomplete = "reference_architecture_vala_incomplete"
	ReferenceArchitectureValAStateBlocked    = "reference_architecture_vala_blocked"
	ReferenceArchitectureValAStateUnknown    = "reference_architecture_vala_unknown"
)

type ReferenceArchitectureBlueprintFamilyProfile struct {
	CurrentState                         string                                  `json:"current_state"`
	BlueprintID                          string                                  `json:"blueprint_id"`
	Version                              string                                  `json:"version"`
	Family                               string                                  `json:"family"`
	Title                                string                                  `json:"title"`
	LifecycleState                       string                                  `json:"lifecycle_state"`
	CompatibilityState                   string                                  `json:"compatibility_state"`
	Owner                                string                                  `json:"owner"`
	TargetEnvironment                    ReferenceArchitectureEnvironmentProfile `json:"target_environment"`
	InfrastructureAssumptions            []string                                `json:"infrastructure_assumptions,omitempty"`
	NetworkAssumptions                   []string                                `json:"network_assumptions,omitempty"`
	IdentityAccessAssumptions            []string                                `json:"identity_access_assumptions,omitempty"`
	TrustCustodyAssumptions              []string                                `json:"trust_custody_assumptions,omitempty"`
	StorageAssumptions                   []string                                `json:"storage_assumptions,omitempty"`
	OperationalAssumptions               []string                                `json:"operational_assumptions,omitempty"`
	SupportAssumptions                   []string                                `json:"support_assumptions,omitempty"`
	RequiredCapabilities                 []string                                `json:"required_capabilities,omitempty"`
	OptionalCapabilities                 []string                                `json:"optional_capabilities,omitempty"`
	DegradedConditions                   []string                                `json:"degraded_conditions,omitempty"`
	UnsupportedConditions                []string                                `json:"unsupported_conditions,omitempty"`
	RequiredEvidenceTypes                []string                                `json:"required_evidence_types,omitempty"`
	SupportBoundaryRef                   string                                  `json:"support_boundary_ref"`
	Caveats                              []string                                `json:"caveats,omitempty"`
	ProjectionDisclaimer                 string                                  `json:"projection_disclaimer"`
	CreatedAt                            string                                  `json:"created_at"`
	UpdatedAt                            string                                  `json:"updated_at"`
	OperatorSupportBoundaryRequired      bool                                    `json:"operator_support_boundary_required"`
	StrongerTrustAnchorMode              bool                                    `json:"stronger_trust_anchor_mode"`
	StricterAuditCustody                 bool                                    `json:"stricter_audit_custody"`
	StrongerEvidenceStorage              bool                                    `json:"stronger_evidence_storage"`
	StrongerRecoveryExpectations         bool                                    `json:"stronger_recovery_expectations"`
	TighterOperatorControl               bool                                    `json:"tighter_operator_control"`
	DataResidencyDisciplineRequired      bool                                    `json:"data_residency_discipline_required"`
	RedactionExportBoundaryRequired      bool                                    `json:"redaction_export_boundary_required"`
	EvidenceCustodyRequired              bool                                    `json:"evidence_custody_required"`
	OfflineTransferBoundaryRequired      bool                                    `json:"offline_transfer_boundary_required"`
	LocalTrustAnchorAssumptionRequired   bool                                    `json:"local_trust_anchor_assumption_required"`
	LocalOperatorControlRequired         bool                                    `json:"local_operator_control_required"`
	LiveExternalDependencyAllowedOffline bool                                    `json:"live_external_dependency_allowed_offline"`
	PerformanceEnvelopeRequired          bool                                    `json:"performance_envelope_required"`
	AuditWritePathDisciplineRequired     bool                                    `json:"audit_write_path_discipline_required"`
	ControlPlaneCapacityRequired         bool                                    `json:"control_plane_capacity_required"`
	CustomerAuthorityBoundaryRequired    bool                                    `json:"customer_authority_boundary_required"`
	NoPartnerShadowTruthRule             bool                                    `json:"no_partner_shadow_truth_rule"`
	PartnerVisibilityRestrictionRequired bool                                    `json:"partner_visibility_restriction_required"`
	PartnerCanonicalTruthOverrideAllowed bool                                    `json:"partner_canonical_truth_override_allowed"`
	RequiresAllWorkloadsInEnclaves       bool                                    `json:"requires_all_workloads_in_enclaves"`
	CertifiedLanguagePresent             bool                                    `json:"certified_language_present"`
	GuaranteedSecurityClaimPresent       bool                                    `json:"guaranteed_security_claim_present"`
	AbsoluteSecurityClaimPresent         bool                                    `json:"absolute_security_claim_present"`
	ClaimsPoint6Pass                     bool                                    `json:"claims_point_6_pass"`
}

type ReferenceArchitectureBlueprintFamilyRegistry struct {
	CurrentState         string                                        `json:"current_state"`
	RegistryID           string                                        `json:"registry_id"`
	Version              string                                        `json:"version"`
	SupportedFamilies    []string                                      `json:"supported_families,omitempty"`
	Profiles             []ReferenceArchitectureBlueprintFamilyProfile `json:"profiles,omitempty"`
	ProjectionDisclaimer string                                        `json:"projection_disclaimer"`
}

func referenceArchitectureValAProfileProjectionDisclaimer() string {
	return "projection_only not_canonical_truth validated_reference_blueprint_profile advisory_only"
}

func referenceArchitectureValARegistryProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_reference_architecture_family_registry"
}

func referenceArchitectureValAEnterpriseDefaultProfile() ReferenceArchitectureBlueprintFamilyProfile {
	return ReferenceArchitectureBlueprintFamilyProfile{
		CurrentState:       "reference_architecture_vala_family_profile_ready",
		BlueprintID:        "reference-blueprint-vala-enterprise-default-001",
		Version:            "1.0.0",
		Family:             ReferenceArchitectureFamilyEnterpriseDefault,
		Title:              "Enterprise Default",
		LifecycleState:     ReferenceArchitectureLifecycleActive,
		CompatibilityState: ReferenceArchitectureCompatibilityCompatible,
		Owner:              "reference_architecture_program",
		TargetEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologyMultiRegion,
			TrustAnchorMode:             ReferenceArchitectureTrustHSMBacked,
			AuditPathMode:               ReferenceArchitectureAuditRegional,
			ConnectivityMode:            ReferenceArchitectureConnectivityRestricted,
			DataResidencyMode:           ReferenceArchitectureResidencyRegional,
			OperatorControlModel:        ReferenceArchitectureOperatorCustomer,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessDirectVerifier,
		},
		InfrastructureAssumptions: []string{
			"regional enterprise substrate supports signer isolation, resilient policy distribution, and durable service boundaries",
		},
		NetworkAssumptions: []string{
			"standard enterprise restricted egress and internal control-plane connectivity remain available",
		},
		IdentityAccessAssumptions: []string{
			"enterprise identities are governed with scoped operator and verifier access",
		},
		TrustCustodyAssumptions: []string{
			"trust anchors remain under governed enterprise custody with standard rotation controls",
		},
		StorageAssumptions: []string{
			"evidence storage remains durable, append-oriented, and available for bounded recovery flows",
		},
		OperationalAssumptions: []string{
			"standard production runbooks, HA readiness, and external integration boundaries remain maintained",
		},
		SupportAssumptions: []string{
			"named operator support boundary exists for enterprise production operation",
		},
		RequiredCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
		},
		OptionalCapabilities: []string{
			ReferenceArchitectureCapabilityVerifierAccess,
		},
		DegradedConditions: []string{
			ReferenceArchitectureDegradedCapabilityGap,
			ReferenceArchitectureDegradedAuditPathReduced,
			ReferenceArchitectureDegradedSupportLimited,
			ReferenceArchitectureDegradedCompatibilityWarn,
			ReferenceArchitectureDegradedTopologyVariance,
		},
		UnsupportedConditions: []string{
			ReferenceArchitectureUnsupportedTrustMismatch,
			ReferenceArchitectureUnsupportedConnectivity,
			ReferenceArchitectureUnsupportedEvidenceMissing,
			ReferenceArchitectureUnsupportedLifecycle,
		},
		RequiredEvidenceTypes: []string{
			ReferenceArchitectureEvidenceDeploymentObservation,
			ReferenceArchitectureEvidenceCapabilityAttestation,
			ReferenceArchitectureEvidenceAuditSnapshot,
			ReferenceArchitectureEvidenceCompatibilityReport,
			ReferenceArchitectureEvidenceSupportBoundary,
		},
		SupportBoundaryRef:              "support-boundary:enterprise-default",
		Caveats:                         []string{"bounded balanced production profile with measured conformance only"},
		ProjectionDisclaimer:            referenceArchitectureValAProfileProjectionDisclaimer(),
		CreatedAt:                       "2026-04-26T11:00:00Z",
		UpdatedAt:                       "2026-04-26T11:05:00Z",
		OperatorSupportBoundaryRequired: true,
	}
}

func referenceArchitectureValAHighAssuranceProfile() ReferenceArchitectureBlueprintFamilyProfile {
	return ReferenceArchitectureBlueprintFamilyProfile{
		CurrentState:       "reference_architecture_vala_family_profile_ready",
		BlueprintID:        "reference-blueprint-vala-high-assurance-001",
		Version:            "1.0.0",
		Family:             ReferenceArchitectureFamilyHighAssurance,
		Title:              "High Assurance",
		LifecycleState:     ReferenceArchitectureLifecycleActive,
		CompatibilityState: ReferenceArchitectureCompatibilityCompatible,
		Owner:              "reference_architecture_program",
		TargetEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologyMultiRegion,
			TrustAnchorMode:             ReferenceArchitectureTrustOfflineIntermediates,
			AuditPathMode:               ReferenceArchitectureAuditRegional,
			ConnectivityMode:            ReferenceArchitectureConnectivityRestricted,
			DataResidencyMode:           ReferenceArchitectureResidencyRegional,
			OperatorControlModel:        ReferenceArchitectureOperatorCustomer,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessDirectVerifier,
		},
		InfrastructureAssumptions: []string{
			"critical production substrate preserves stronger signer isolation and control-plane fault boundaries",
		},
		NetworkAssumptions: []string{
			"restricted network posture preserves stronger audit and policy transport discipline",
		},
		IdentityAccessAssumptions: []string{
			"higher-assurance operator identities are tightly scoped and periodically reviewed",
		},
		TrustCustodyAssumptions: []string{
			"stronger trust anchor custody is maintained through offline-root or equivalent governed intermediate discipline",
		},
		StorageAssumptions: []string{
			"evidence storage uses stronger custody and tamper-resistance expectations than enterprise default",
		},
		OperationalAssumptions: []string{
			"recovery drills and operator change discipline are tighter than enterprise default for selected critical workloads",
		},
		SupportAssumptions: []string{
			"high-assurance operator support boundary remains staffed for tighter review and recovery expectations",
		},
		RequiredCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
		},
		OptionalCapabilities: []string{
			ReferenceArchitectureCapabilityVerifierAccess,
		},
		DegradedConditions: []string{
			ReferenceArchitectureDegradedCapabilityGap,
			ReferenceArchitectureDegradedAuditPathReduced,
			ReferenceArchitectureDegradedSupportLimited,
			ReferenceArchitectureDegradedEvidenceCaveated,
		},
		UnsupportedConditions: []string{
			ReferenceArchitectureUnsupportedTrustMismatch,
			ReferenceArchitectureUnsupportedEvidenceMissing,
			ReferenceArchitectureUnsupportedLifecycle,
		},
		RequiredEvidenceTypes: []string{
			ReferenceArchitectureEvidenceDeploymentObservation,
			ReferenceArchitectureEvidenceCapabilityAttestation,
			ReferenceArchitectureEvidenceAuditSnapshot,
			ReferenceArchitectureEvidenceCompatibilityReport,
			ReferenceArchitectureEvidenceSupportBoundary,
		},
		SupportBoundaryRef:              "support-boundary:high-assurance",
		Caveats:                         []string{"bounded high-assurance profile for selected workloads; it does not claim absolute security"},
		ProjectionDisclaimer:            referenceArchitectureValAProfileProjectionDisclaimer(),
		CreatedAt:                       "2026-04-26T11:00:00Z",
		UpdatedAt:                       "2026-04-26T11:05:00Z",
		OperatorSupportBoundaryRequired: true,
		StrongerTrustAnchorMode:         true,
		StricterAuditCustody:            true,
		StrongerEvidenceStorage:         true,
		StrongerRecoveryExpectations:    true,
		TighterOperatorControl:          true,
	}
}

func referenceArchitectureValARegulatedPrivacyFirstProfile() ReferenceArchitectureBlueprintFamilyProfile {
	return ReferenceArchitectureBlueprintFamilyProfile{
		CurrentState:       "reference_architecture_vala_family_profile_ready",
		BlueprintID:        "reference-blueprint-vala-regulated-privacy-first-001",
		Version:            "1.0.0",
		Family:             ReferenceArchitectureFamilyRegulatedPrivacyFirst,
		Title:              "Regulated Privacy First",
		LifecycleState:     ReferenceArchitectureLifecycleActive,
		CompatibilityState: ReferenceArchitectureCompatibilityCompatible,
		Owner:              "reference_architecture_program",
		TargetEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologySingleRegion,
			TrustAnchorMode:             ReferenceArchitectureTrustHSMBacked,
			AuditPathMode:               ReferenceArchitectureAuditRegional,
			ConnectivityMode:            ReferenceArchitectureConnectivityRestricted,
			DataResidencyMode:           ReferenceArchitectureResidencySovereignLocal,
			OperatorControlModel:        ReferenceArchitectureOperatorCustomer,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessBrokeredPartner,
		},
		InfrastructureAssumptions: []string{
			"regulated deployment substrate preserves local evidence handling and bounded export controls",
		},
		NetworkAssumptions: []string{
			"restricted sharing and controlled egress boundaries are enforced around audit and evidence paths",
		},
		IdentityAccessAssumptions: []string{
			"privacy-sensitive operator and auditor access is scoped through reviewed release boundaries",
		},
		TrustCustodyAssumptions: []string{
			"evidence custody and trust anchor handling remain bounded by regulated data-handling controls",
		},
		StorageAssumptions: []string{
			"evidence and audit storage preserve regional residency and reviewed retention boundaries",
		},
		OperationalAssumptions: []string{
			"redaction, export review, and freshness discipline remain explicit for external evidence sharing",
		},
		SupportAssumptions: []string{
			"regulated support boundary remains defined for privacy-first operations and controlled evidence exchange",
		},
		RequiredCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
		},
		OptionalCapabilities: []string{
			ReferenceArchitectureCapabilityVerifierAccess,
		},
		DegradedConditions: []string{
			ReferenceArchitectureDegradedCapabilityGap,
			ReferenceArchitectureDegradedSupportLimited,
			ReferenceArchitectureDegradedEvidenceCaveated,
			ReferenceArchitectureDegradedCompatibilityWarn,
		},
		UnsupportedConditions: []string{
			ReferenceArchitectureUnsupportedConnectivity,
			ReferenceArchitectureUnsupportedEvidenceMissing,
			ReferenceArchitectureUnsupportedLifecycle,
		},
		RequiredEvidenceTypes: []string{
			ReferenceArchitectureEvidenceDeploymentObservation,
			ReferenceArchitectureEvidenceCapabilityAttestation,
			ReferenceArchitectureEvidenceAuditSnapshot,
			ReferenceArchitectureEvidenceCompatibilityReport,
			ReferenceArchitectureEvidenceSupportBoundary,
		},
		SupportBoundaryRef:              "support-boundary:regulated-privacy-first",
		Caveats:                         []string{"bounded privacy-first profile with evidence export and redaction discipline; it is not legal certification"},
		ProjectionDisclaimer:            referenceArchitectureValAProfileProjectionDisclaimer(),
		CreatedAt:                       "2026-04-26T11:00:00Z",
		UpdatedAt:                       "2026-04-26T11:05:00Z",
		OperatorSupportBoundaryRequired: true,
		DataResidencyDisciplineRequired: true,
		RedactionExportBoundaryRequired: true,
		EvidenceCustodyRequired:         true,
	}
}

func referenceArchitectureValASovereignAirGappedProfile() ReferenceArchitectureBlueprintFamilyProfile {
	return ReferenceArchitectureBlueprintFamilyProfile{
		CurrentState:       "reference_architecture_vala_family_profile_ready",
		BlueprintID:        "reference-blueprint-vala-sovereign-air-gapped-001",
		Version:            "1.0.0",
		Family:             ReferenceArchitectureFamilySovereignAirGapped,
		Title:              "Sovereign Air Gapped",
		LifecycleState:     ReferenceArchitectureLifecycleActive,
		CompatibilityState: ReferenceArchitectureCompatibilityCompatible,
		Owner:              "reference_architecture_program",
		TargetEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologyAirGappedCell,
			TrustAnchorMode:             ReferenceArchitectureTrustAirGappedOfflineRoot,
			AuditPathMode:               ReferenceArchitectureAuditDeferredAirGap,
			ConnectivityMode:            ReferenceArchitectureConnectivityAirGapped,
			DataResidencyMode:           ReferenceArchitectureResidencySovereignLocal,
			OperatorControlModel:        ReferenceArchitectureOperatorCustomer,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessOfflineEvidence,
		},
		InfrastructureAssumptions: []string{
			"sovereign air-gapped substrate keeps local control-plane, signer, and evidence services self-contained",
		},
		NetworkAssumptions: []string{
			"restricted or offline connectivity is expected and external dependencies remain bounded out of the reference profile",
		},
		IdentityAccessAssumptions: []string{
			"local operator identities remain under sovereign control with explicit offline evidence exchange discipline",
		},
		TrustCustodyAssumptions: []string{
			"local trust anchors remain under sovereign custody with offline root handling and reviewed transfer boundaries",
		},
		StorageAssumptions: []string{
			"local evidence custody remains durable and export occurs only through bounded offline transfer paths",
		},
		OperationalAssumptions: []string{
			"offline publication and transfer remain manually governed outside this Val A profile",
		},
		SupportAssumptions: []string{
			"local operator support boundary is explicit because live partner or verifier access is restricted",
		},
		RequiredCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
			ReferenceArchitectureCapabilityAirGapTransfer,
		},
		OptionalCapabilities: []string{},
		DegradedConditions: []string{
			ReferenceArchitectureDegradedCapabilityGap,
			ReferenceArchitectureDegradedAuditPathReduced,
			ReferenceArchitectureDegradedSupportLimited,
			ReferenceArchitectureDegradedEvidenceCaveated,
		},
		UnsupportedConditions: []string{
			ReferenceArchitectureUnsupportedConnectivity,
			ReferenceArchitectureUnsupportedAirGapTransfer,
			ReferenceArchitectureUnsupportedTrustMismatch,
			ReferenceArchitectureUnsupportedEvidenceMissing,
			ReferenceArchitectureUnsupportedLifecycle,
		},
		RequiredEvidenceTypes: []string{
			ReferenceArchitectureEvidenceDeploymentObservation,
			ReferenceArchitectureEvidenceCapabilityAttestation,
			ReferenceArchitectureEvidenceAuditSnapshot,
			ReferenceArchitectureEvidenceCompatibilityReport,
			ReferenceArchitectureEvidenceSupportBoundary,
		},
		SupportBoundaryRef:                   "support-boundary:sovereign-air-gapped",
		Caveats:                              []string{"bounded sovereign or air-gapped profile; offline transfer tooling and publication remain out of scope for Val A"},
		ProjectionDisclaimer:                 referenceArchitectureValAProfileProjectionDisclaimer(),
		CreatedAt:                            "2026-04-26T11:00:00Z",
		UpdatedAt:                            "2026-04-26T11:05:00Z",
		OperatorSupportBoundaryRequired:      true,
		OfflineTransferBoundaryRequired:      true,
		LocalTrustAnchorAssumptionRequired:   true,
		LocalOperatorControlRequired:         true,
		LiveExternalDependencyAllowedOffline: false,
	}
}

func referenceArchitectureValAPerformanceSensitiveProfile() ReferenceArchitectureBlueprintFamilyProfile {
	return ReferenceArchitectureBlueprintFamilyProfile{
		CurrentState:       "reference_architecture_vala_family_profile_ready",
		BlueprintID:        "reference-blueprint-vala-performance-sensitive-001",
		Version:            "1.0.0",
		Family:             ReferenceArchitectureFamilyPerformanceSensitive,
		Title:              "Performance Sensitive",
		LifecycleState:     ReferenceArchitectureLifecycleActive,
		CompatibilityState: ReferenceArchitectureCompatibilityCompatible,
		Owner:              "reference_architecture_program",
		TargetEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologyMultiRegion,
			TrustAnchorMode:             ReferenceArchitectureTrustHSMBacked,
			AuditPathMode:               ReferenceArchitectureAuditRegional,
			ConnectivityMode:            ReferenceArchitectureConnectivityConnected,
			DataResidencyMode:           ReferenceArchitectureResidencyRegional,
			OperatorControlModel:        ReferenceArchitectureOperatorCustomer,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessDirectVerifier,
		},
		InfrastructureAssumptions: []string{
			"performance-sensitive substrate preserves control-plane capacity and bounded latency under declared operating envelope",
		},
		NetworkAssumptions: []string{
			"audit and evidence write paths remain provisioned so latency-sensitive workloads do not silently drop conformance visibility",
		},
		IdentityAccessAssumptions: []string{
			"operator and verifier access remains bounded so control-plane load can be forecast and reviewed",
		},
		TrustCustodyAssumptions: []string{
			"trust and signing flows remain stable under the declared performance envelope",
		},
		StorageAssumptions: []string{
			"evidence storage and audit buffering remain provisioned to tolerate declared throughput and backlog windows",
		},
		OperationalAssumptions: []string{
			"performance assumptions, audit write path expectations, and recovery windows remain explicit and reviewed",
		},
		SupportAssumptions: []string{
			"performance-sensitive support boundary remains explicit for overload and delayed-path degradation handling",
		},
		RequiredCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
		},
		OptionalCapabilities: []string{
			ReferenceArchitectureCapabilityVerifierAccess,
		},
		DegradedConditions: []string{
			ReferenceArchitectureDegradedCapabilityGap,
			ReferenceArchitectureDegradedAuditPathReduced,
			ReferenceArchitectureDegradedTopologyVariance,
			ReferenceArchitectureDegradedSupportLimited,
		},
		UnsupportedConditions: []string{
			ReferenceArchitectureUnsupportedEvidenceMissing,
			ReferenceArchitectureUnsupportedLifecycle,
			ReferenceArchitectureUnsupportedUnknownEnv,
		},
		RequiredEvidenceTypes: []string{
			ReferenceArchitectureEvidenceDeploymentObservation,
			ReferenceArchitectureEvidenceCapabilityAttestation,
			ReferenceArchitectureEvidenceAuditSnapshot,
			ReferenceArchitectureEvidenceCompatibilityReport,
			ReferenceArchitectureEvidenceSupportBoundary,
		},
		SupportBoundaryRef:               "support-boundary:performance-sensitive",
		Caveats:                          []string{"bounded performance-sensitive profile with explicit degradation windows; it does not guarantee latency or throughput"},
		ProjectionDisclaimer:             referenceArchitectureValAProfileProjectionDisclaimer(),
		CreatedAt:                        "2026-04-26T11:00:00Z",
		UpdatedAt:                        "2026-04-26T11:05:00Z",
		OperatorSupportBoundaryRequired:  true,
		PerformanceEnvelopeRequired:      true,
		AuditWritePathDisciplineRequired: true,
		ControlPlaneCapacityRequired:     true,
	}
}

func referenceArchitectureValAPartnerMSPSuitableProfile() ReferenceArchitectureBlueprintFamilyProfile {
	return ReferenceArchitectureBlueprintFamilyProfile{
		CurrentState:       "reference_architecture_vala_family_profile_ready",
		BlueprintID:        "reference-blueprint-vala-partner-msp-suitable-001",
		Version:            "1.0.0",
		Family:             ReferenceArchitectureFamilyPartnerMSPSuitable,
		Title:              "Partner MSP Suitable",
		LifecycleState:     ReferenceArchitectureLifecycleActive,
		CompatibilityState: ReferenceArchitectureCompatibilityCompatible,
		Owner:              "reference_architecture_program",
		TargetEnvironment: ReferenceArchitectureEnvironmentProfile{
			DeploymentTopology:          ReferenceArchitectureTopologyPartnerIsolated,
			TrustAnchorMode:             ReferenceArchitectureTrustHSMBacked,
			AuditPathMode:               ReferenceArchitectureAuditRegional,
			ConnectivityMode:            ReferenceArchitectureConnectivityRestricted,
			DataResidencyMode:           ReferenceArchitectureResidencyRegional,
			OperatorControlModel:        ReferenceArchitectureOperatorPartner,
			VerifierOrPartnerAccessMode: ReferenceArchitectureAccessBrokeredPartner,
		},
		InfrastructureAssumptions: []string{
			"partner-operated substrate preserves customer-controlled audit and evidence boundaries",
		},
		NetworkAssumptions: []string{
			"partner visibility remains scoped through reviewed connectivity and evidence-release boundaries",
		},
		IdentityAccessAssumptions: []string{
			"partner, verifier, and customer identities remain separately scoped with customer authority preserved",
		},
		TrustCustodyAssumptions: []string{
			"customer trust and custody boundaries remain explicit even when partner assists operations",
		},
		StorageAssumptions: []string{
			"evidence visibility is restricted and does not create partner shadow truth",
		},
		OperationalAssumptions: []string{
			"customer authority boundary and support escalation path remain explicit for partner-managed operations",
		},
		SupportAssumptions: []string{
			"partner or MSP support boundary remains defined without becoming approval or canonical truth authority",
		},
		RequiredCapabilities: []string{
			ReferenceArchitectureCapabilitySigning,
			ReferenceArchitectureCapabilityAuditWriter,
			ReferenceArchitectureCapabilityEvidenceStorage,
			ReferenceArchitectureCapabilityPolicyDist,
			ReferenceArchitectureCapabilityRecovery,
			ReferenceArchitectureCapabilityVerifierAccess,
		},
		OptionalCapabilities: []string{},
		DegradedConditions: []string{
			ReferenceArchitectureDegradedCapabilityGap,
			ReferenceArchitectureDegradedSupportLimited,
			ReferenceArchitectureDegradedEvidenceCaveated,
			ReferenceArchitectureDegradedCompatibilityWarn,
		},
		UnsupportedConditions: []string{
			ReferenceArchitectureUnsupportedTrustMismatch,
			ReferenceArchitectureUnsupportedEvidenceMissing,
			ReferenceArchitectureUnsupportedLifecycle,
		},
		RequiredEvidenceTypes: []string{
			ReferenceArchitectureEvidenceDeploymentObservation,
			ReferenceArchitectureEvidenceCapabilityAttestation,
			ReferenceArchitectureEvidenceAuditSnapshot,
			ReferenceArchitectureEvidenceCompatibilityReport,
			ReferenceArchitectureEvidenceSupportBoundary,
		},
		SupportBoundaryRef:                   "support-boundary:partner-msp-suitable",
		Caveats:                              []string{"bounded partner or MSP suitability profile; partner visibility does not become canonical truth or approval authority"},
		ProjectionDisclaimer:                 referenceArchitectureValAProfileProjectionDisclaimer(),
		CreatedAt:                            "2026-04-26T11:00:00Z",
		UpdatedAt:                            "2026-04-26T11:05:00Z",
		OperatorSupportBoundaryRequired:      true,
		CustomerAuthorityBoundaryRequired:    true,
		NoPartnerShadowTruthRule:             true,
		PartnerVisibilityRestrictionRequired: true,
		PartnerCanonicalTruthOverrideAllowed: false,
	}
}

func ReferenceArchitectureValAFamilyProfiles() []ReferenceArchitectureBlueprintFamilyProfile {
	return []ReferenceArchitectureBlueprintFamilyProfile{
		referenceArchitectureValAEnterpriseDefaultProfile(),
		referenceArchitectureValAHighAssuranceProfile(),
		referenceArchitectureValARegulatedPrivacyFirstProfile(),
		referenceArchitectureValASovereignAirGappedProfile(),
		referenceArchitectureValAPerformanceSensitiveProfile(),
		referenceArchitectureValAPartnerMSPSuitableProfile(),
	}
}

func ReferenceArchitectureValAFamilyRegistry() ReferenceArchitectureBlueprintFamilyRegistry {
	return ReferenceArchitectureBlueprintFamilyRegistry{
		CurrentState:         "reference_architecture_vala_family_registry_ready",
		RegistryID:           "reference-architecture-vala-family-registry",
		Version:              "1.0.0",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Profiles:             ReferenceArchitectureValAFamilyProfiles(),
		ProjectionDisclaimer: referenceArchitectureValARegistryProjectionDisclaimer(),
	}
}

func LookupReferenceArchitectureValAFamilyProfile(family string) (ReferenceArchitectureBlueprintFamilyProfile, bool) {
	for _, profile := range ReferenceArchitectureValAFamilyRegistry().Profiles {
		if strings.TrimSpace(profile.Family) == strings.TrimSpace(family) {
			return profile, true
		}
	}
	return ReferenceArchitectureBlueprintFamilyProfile{}, false
}

func referenceArchitectureValAObservedCapabilities(profile ReferenceArchitectureBlueprintFamilyProfile) []string {
	observed := make([]string, 0, len(profile.RequiredCapabilities)+len(profile.OptionalCapabilities))
	seen := make(map[string]struct{}, len(profile.RequiredCapabilities)+len(profile.OptionalCapabilities))
	for _, capability := range append(append([]string{}, profile.RequiredCapabilities...), profile.OptionalCapabilities...) {
		trimmed := strings.TrimSpace(capability)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		observed = append(observed, trimmed)
	}
	return observed
}

func referenceArchitectureValAProfileEvidenceRefs(profile ReferenceArchitectureBlueprintFamilyProfile) []ReferenceArchitectureEvidenceReference {
	refs := make([]ReferenceArchitectureEvidenceReference, 0, len(profile.RequiredEvidenceTypes))
	caveats := append([]string{}, profile.Caveats...)
	if len(caveats) == 0 {
		caveats = []string{"bounded evidence requirement metadata for validated reference blueprint profile"}
	}
	for idx, evidenceType := range profile.RequiredEvidenceTypes {
		refs = append(refs, ReferenceArchitectureEvidenceReference{
			EvidenceID:     "evidence:" + strings.ReplaceAll(profile.Family, "_", "-") + "-required-" + strings.TrimSpace(evidenceType),
			EvidenceType:   evidenceType,
			Source:         "reference-profile/" + profile.Family,
			Timestamp:      profile.UpdatedAt,
			FreshnessState: IntelligenceCalibrationFreshnessFresh,
			Scope:          "required_evidence/" + profile.Family + "/" + strings.TrimSpace(evidenceType),
			Caveats:        append([]string{}, caveats...),
		})
		if idx == len(profile.RequiredEvidenceTypes)-1 {
			continue
		}
	}
	return refs
}

func referenceArchitectureValAProfileBlueprintContract(profile ReferenceArchitectureBlueprintFamilyProfile) ReferenceArchitectureBlueprintContract {
	return ReferenceArchitectureBlueprintContract{
		CurrentState:                   profile.CurrentState,
		BlueprintID:                    profile.BlueprintID,
		Version:                        profile.Version,
		Family:                         profile.Family,
		LifecycleState:                 profile.LifecycleState,
		Owner:                          profile.Owner,
		CreatedAt:                      profile.CreatedAt,
		UpdatedAt:                      profile.UpdatedAt,
		SupportedFamilies:              referenceArchitectureVal0Families(),
		SupportedLifecycleStates:       referenceArchitectureVal0LifecycleStates(),
		SupportedCompatibilityStates:   referenceArchitectureVal0CompatibilityStates(),
		SupportedConformanceStates:     referenceArchitectureVal0ConformanceStates(),
		TargetEnvironment:              profile.TargetEnvironment,
		ObservedEnvironment:            profile.TargetEnvironment,
		RequiredCapabilities:           append([]string{}, profile.RequiredCapabilities...),
		OptionalCapabilities:           append([]string{}, profile.OptionalCapabilities...),
		ObservedCapabilities:           referenceArchitectureValAObservedCapabilities(profile),
		InfrastructureAssumptions:      append([]string{}, profile.InfrastructureAssumptions...),
		NetworkAssumptions:             append([]string{}, profile.NetworkAssumptions...),
		IdentityAccessAssumptions:      append([]string{}, profile.IdentityAccessAssumptions...),
		TrustCustodyAssumptions:        append([]string{}, profile.TrustCustodyAssumptions...),
		StorageAssumptions:             append([]string{}, profile.StorageAssumptions...),
		OperationalAssumptions:         append([]string{}, profile.OperationalAssumptions...),
		SupportAssumptions:             append([]string{}, profile.SupportAssumptions...),
		SupportedDegradedConditions:    referenceArchitectureVal0SupportedDegradedConditions(),
		SupportedUnsupportedConditions: referenceArchitectureVal0SupportedUnsupportedConditions(),
		CompatibilityState:             profile.CompatibilityState,
		ConformanceState:               ReferenceArchitectureConformanceMatched,
		SupportBoundaryRef:             profile.SupportBoundaryRef,
		EvidenceRefs:                   referenceArchitectureValAProfileEvidenceRefs(profile),
		RedactionKeepsCaveats:          true,
		CertifiedLanguagePresent:       profile.CertifiedLanguagePresent,
		ProjectionDisclaimer:           profile.ProjectionDisclaimer,
	}
}

func referenceArchitectureValARequiresVal0Active(point5State, val0CurrentState, val0State, point6State string) bool {
	return strings.TrimSpace(point5State) == IntelligenceCalibrationPoint5StatePass &&
		strings.TrimSpace(val0CurrentState) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(val0State) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(point6State) == ReferenceArchitecturePoint6StateNotComplete
}

func referenceArchitectureValAHasClaimViolation(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	return profile.CertifiedLanguagePresent ||
		profile.GuaranteedSecurityClaimPresent ||
		profile.AbsoluteSecurityClaimPresent ||
		profile.ClaimsPoint6Pass ||
		profile.RequiresAllWorkloadsInEnclaves
}

func referenceArchitectureValAFamilyProfileHasRequiredFields(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	return strings.TrimSpace(profile.BlueprintID) != "" &&
		strings.TrimSpace(profile.Version) != "" &&
		strings.TrimSpace(profile.Family) != "" &&
		strings.TrimSpace(profile.Title) != "" &&
		strings.TrimSpace(profile.LifecycleState) != "" &&
		strings.TrimSpace(profile.CompatibilityState) != "" &&
		strings.TrimSpace(profile.Owner) != "" &&
		strings.TrimSpace(profile.SupportBoundaryRef) != "" &&
		strings.TrimSpace(profile.ProjectionDisclaimer) != "" &&
		strings.TrimSpace(profile.CreatedAt) != "" &&
		strings.TrimSpace(profile.UpdatedAt) != "" &&
		len(profile.RequiredCapabilities) > 0 &&
		len(profile.RequiredEvidenceTypes) > 0 &&
		len(profile.DegradedConditions) > 0 &&
		len(profile.UnsupportedConditions) > 0 &&
		referenceArchitectureVal0AllAssumptionsPresent(referenceArchitectureValAProfileBlueprintContract(profile))
}

func referenceArchitectureValAFamilyProfileSupportsVal0Enums(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	if !containsTrimmedString(referenceArchitectureVal0Families(), profile.Family) ||
		!containsTrimmedString(referenceArchitectureVal0LifecycleStates(), profile.LifecycleState) ||
		!containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), profile.CompatibilityState) ||
		!referenceArchitectureVal0EnvironmentSupported(profile.TargetEnvironment) ||
		!containsAllTrimmedStrings(referenceArchitectureVal0SupportedCapabilities(), profile.RequiredCapabilities...) ||
		!containsAllTrimmedStrings(referenceArchitectureVal0SupportedCapabilities(), profile.OptionalCapabilities...) ||
		!containsAllTrimmedStrings(referenceArchitectureVal0SupportedDegradedConditions(), profile.DegradedConditions...) ||
		!containsAllTrimmedStrings(referenceArchitectureVal0SupportedUnsupportedConditions(), profile.UnsupportedConditions...) ||
		!containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), profile.RequiredEvidenceTypes...) {
		return false
	}
	return true
}

func referenceArchitectureValAEnterpriseDefaultSatisfied(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	required := []string{
		ReferenceArchitectureCapabilitySigning,
		ReferenceArchitectureCapabilityAuditWriter,
		ReferenceArchitectureCapabilityEvidenceStorage,
		ReferenceArchitectureCapabilityPolicyDist,
		ReferenceArchitectureCapabilityRecovery,
	}
	return containsAllTrimmedStrings(profile.RequiredCapabilities, required...) &&
		profile.OperatorSupportBoundaryRequired
}

func referenceArchitectureValAHighAssuranceSatisfied(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	return profile.StrongerTrustAnchorMode &&
		profile.StricterAuditCustody &&
		profile.StrongerEvidenceStorage &&
		profile.StrongerRecoveryExpectations &&
		profile.TighterOperatorControl &&
		strings.TrimSpace(profile.TargetEnvironment.TrustAnchorMode) != strings.TrimSpace(referenceArchitectureValAEnterpriseDefaultProfile().TargetEnvironment.TrustAnchorMode) &&
		!profile.RequiresAllWorkloadsInEnclaves
}

func referenceArchitectureValARegulatedPrivacyFirstSatisfied(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	return profile.DataResidencyDisciplineRequired &&
		profile.RedactionExportBoundaryRequired &&
		profile.EvidenceCustodyRequired &&
		strings.TrimSpace(profile.TargetEnvironment.DataResidencyMode) != ""
}

func referenceArchitectureValASovereignAirGappedSatisfied(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	return profile.OfflineTransferBoundaryRequired &&
		profile.LocalTrustAnchorAssumptionRequired &&
		profile.LocalOperatorControlRequired &&
		!profile.LiveExternalDependencyAllowedOffline &&
		strings.TrimSpace(profile.TargetEnvironment.ConnectivityMode) == ReferenceArchitectureConnectivityAirGapped &&
		containsTrimmedString(profile.RequiredCapabilities, ReferenceArchitectureCapabilityAirGapTransfer)
}

func referenceArchitectureValAPerformanceSensitiveSatisfied(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	return profile.PerformanceEnvelopeRequired &&
		profile.AuditWritePathDisciplineRequired &&
		profile.ControlPlaneCapacityRequired
}

func referenceArchitectureValAPartnerMSPSuitableSatisfied(profile ReferenceArchitectureBlueprintFamilyProfile) bool {
	return profile.CustomerAuthorityBoundaryRequired &&
		profile.NoPartnerShadowTruthRule &&
		profile.PartnerVisibilityRestrictionRequired &&
		!profile.PartnerCanonicalTruthOverrideAllowed &&
		strings.TrimSpace(profile.TargetEnvironment.VerifierOrPartnerAccessMode) == ReferenceArchitectureAccessBrokeredPartner
}

func EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State string, profile ReferenceArchitectureBlueprintFamilyProfile) string {
	if !referenceArchitectureValARequiresVal0Active(point5State, val0CurrentState, val0State, point6State) {
		return ReferenceArchitectureValAFamilyProfileStateBlocked
	}
	if !referenceArchitectureValAFamilyProfileHasRequiredFields(profile) {
		return ReferenceArchitectureValAFamilyProfileStateIncomplete
	}
	if !referenceArchitectureValAFamilyProfileSupportsVal0Enums(profile) ||
		!strings.Contains(strings.TrimSpace(profile.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(profile.ProjectionDisclaimer), "not_canonical_truth") {
		return ReferenceArchitectureValAFamilyProfileStatePartial
	}
	if referenceArchitectureValAHasClaimViolation(profile) {
		return ReferenceArchitectureValAFamilyProfileStateBlocked
	}
	if strings.TrimSpace(profile.TargetEnvironment.OperatorControlModel) == ReferenceArchitectureOperatorPartner && !profile.OperatorSupportBoundaryRequired {
		return ReferenceArchitectureValAFamilyProfileStatePartial
	}
	contract := referenceArchitectureValAProfileBlueprintContract(profile)
	if EvaluateReferenceArchitectureVal0BlueprintDisciplineState(contract) != ReferenceArchitectureVal0BlueprintDisciplineStateActive ||
		EvaluateReferenceArchitectureVal0TaxonomyState(contract) != ReferenceArchitectureVal0TaxonomyStateActive ||
		EvaluateReferenceArchitectureVal0EnvironmentFitState(contract) != ReferenceArchitectureVal0EnvironmentFitStateActive ||
		EvaluateReferenceArchitectureVal0EvidenceDisciplineState(contract) != ReferenceArchitectureVal0EvidenceDisciplineStateActive ||
		EvaluateReferenceArchitectureVal0CompatibilityBaselineState(contract) != ReferenceArchitectureVal0CompatibilityBaselineStateActive ||
		EvaluateReferenceArchitectureVal0ReferenceConformanceState(contract) != ReferenceArchitectureConformanceMatched {
		return ReferenceArchitectureValAFamilyProfileStatePartial
	}
	switch strings.TrimSpace(profile.Family) {
	case ReferenceArchitectureFamilyEnterpriseDefault:
		if !referenceArchitectureValAEnterpriseDefaultSatisfied(profile) {
			return ReferenceArchitectureValAFamilyProfileStatePartial
		}
	case ReferenceArchitectureFamilyHighAssurance:
		if !referenceArchitectureValAEnterpriseDefaultSatisfied(profile) || !referenceArchitectureValAHighAssuranceSatisfied(profile) {
			return ReferenceArchitectureValAFamilyProfileStatePartial
		}
	case ReferenceArchitectureFamilyRegulatedPrivacyFirst:
		if !referenceArchitectureValAEnterpriseDefaultSatisfied(profile) || !referenceArchitectureValARegulatedPrivacyFirstSatisfied(profile) {
			return ReferenceArchitectureValAFamilyProfileStatePartial
		}
	case ReferenceArchitectureFamilySovereignAirGapped:
		if !referenceArchitectureValAEnterpriseDefaultSatisfied(profile) || !referenceArchitectureValASovereignAirGappedSatisfied(profile) {
			return ReferenceArchitectureValAFamilyProfileStatePartial
		}
	case ReferenceArchitectureFamilyPerformanceSensitive:
		if !referenceArchitectureValAEnterpriseDefaultSatisfied(profile) || !referenceArchitectureValAPerformanceSensitiveSatisfied(profile) {
			return ReferenceArchitectureValAFamilyProfileStatePartial
		}
	case ReferenceArchitectureFamilyPartnerMSPSuitable:
		if !referenceArchitectureValAEnterpriseDefaultSatisfied(profile) || !referenceArchitectureValAPartnerMSPSuitableSatisfied(profile) {
			return ReferenceArchitectureValAFamilyProfileStatePartial
		}
	default:
		return ReferenceArchitectureValAFamilyProfileStateUnknown
	}
	return ReferenceArchitectureValAFamilyProfileStateActive
}

func EvaluateReferenceArchitectureValAFamilyRegistryState(point5State, val0CurrentState, val0State, point6State string, registry ReferenceArchitectureBlueprintFamilyRegistry) string {
	if !referenceArchitectureValARequiresVal0Active(point5State, val0CurrentState, val0State, point6State) {
		return ReferenceArchitectureValAFamilyRegistryStateBlocked
	}
	if strings.TrimSpace(registry.RegistryID) == "" || strings.TrimSpace(registry.Version) == "" || len(registry.Profiles) == 0 || strings.TrimSpace(registry.ProjectionDisclaimer) == "" {
		return ReferenceArchitectureValAFamilyRegistryStateIncomplete
	}
	if !containsExactTrimmedStringSet(registry.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		len(registry.Profiles) != len(referenceArchitectureVal0Families()) ||
		!strings.Contains(strings.TrimSpace(registry.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(registry.ProjectionDisclaimer), "not_canonical_truth") {
		return ReferenceArchitectureValAFamilyRegistryStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenBlueprints := map[string]struct{}{}
	for _, profile := range registry.Profiles {
		family := strings.TrimSpace(profile.Family)
		blueprintID := strings.TrimSpace(profile.BlueprintID)
		if family == "" || blueprintID == "" {
			return ReferenceArchitectureValAFamilyRegistryStateIncomplete
		}
		if _, ok := seenFamilies[family]; ok {
			return ReferenceArchitectureValAFamilyRegistryStatePartial
		}
		if _, ok := seenBlueprints[blueprintID]; ok {
			return ReferenceArchitectureValAFamilyRegistryStatePartial
		}
		seenFamilies[family] = struct{}{}
		seenBlueprints[blueprintID] = struct{}{}
		if EvaluateReferenceArchitectureValAFamilyProfileState(point5State, val0CurrentState, val0State, point6State, profile) != ReferenceArchitectureValAFamilyProfileStateActive {
			return ReferenceArchitectureValAFamilyRegistryStatePartial
		}
	}
	return ReferenceArchitectureValAFamilyRegistryStateActive
}

func EvaluateReferenceArchitectureValAState(point5State, val0CurrentState, val0State, point6State, registryState string) string {
	if !referenceArchitectureValARequiresVal0Active(point5State, val0CurrentState, val0State, point6State) {
		return ReferenceArchitectureValAStateBlocked
	}
	switch strings.TrimSpace(registryState) {
	case ReferenceArchitectureValAFamilyRegistryStateActive:
		return ReferenceArchitectureValAStateActive
	case ReferenceArchitectureValAFamilyRegistryStateBlocked:
		return ReferenceArchitectureValAStateBlocked
	case ReferenceArchitectureValAFamilyRegistryStateIncomplete:
		return ReferenceArchitectureValAStateIncomplete
	case ReferenceArchitectureValAFamilyRegistryStatePartial:
		return ReferenceArchitectureValAStatePartial
	default:
		return ReferenceArchitectureValAStateUnknown
	}
}

func referenceArchitectureValAProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/family-registry",
		"/v1/reference-architecture/vala/family-profiles",
		"/v1/reference-architecture/vala/proofs",
	}
}

func EvaluateReferenceArchitectureValAProofsState(valAState, point6State string, supportedFamilies, surfaceRefs, evidenceRefs, limitations []string, projectionDisclaimer string) string {
	baseState := strings.TrimSpace(valAState)
	if !containsExactTrimmedStringSet(supportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(surfaceRefs, referenceArchitectureValAProofSurfaceRefs()...) ||
		len(evidenceRefs) < 8 ||
		len(limitations) == 0 ||
		!strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == ReferenceArchitectureValAStateActive {
			return ReferenceArchitectureValAStatePartial
		}
		return baseState
	}
	if baseState == ReferenceArchitectureValAStateActive && strings.TrimSpace(point6State) != ReferenceArchitecturePoint6StateNotComplete {
		return ReferenceArchitectureValAStatePartial
	}
	return baseState
}
