package operability

import "strings"

const (
	ReferenceArchitectureValCScenarioPackStateActive     = "reference_architecture_valc_scenario_pack_active"
	ReferenceArchitectureValCScenarioPackStatePartial    = "reference_architecture_valc_scenario_pack_partial"
	ReferenceArchitectureValCScenarioPackStateIncomplete = "reference_architecture_valc_scenario_pack_incomplete"
	ReferenceArchitectureValCScenarioPackStateBlocked    = "reference_architecture_valc_scenario_pack_blocked"
	ReferenceArchitectureValCScenarioPackStateUnknown    = "reference_architecture_valc_scenario_pack_unknown"

	ReferenceArchitectureValCFailureTaxonomyStateActive     = "reference_architecture_valc_failure_taxonomy_active"
	ReferenceArchitectureValCFailureTaxonomyStatePartial    = "reference_architecture_valc_failure_taxonomy_partial"
	ReferenceArchitectureValCFailureTaxonomyStateIncomplete = "reference_architecture_valc_failure_taxonomy_incomplete"
	ReferenceArchitectureValCFailureTaxonomyStateBlocked    = "reference_architecture_valc_failure_taxonomy_blocked"
	ReferenceArchitectureValCFailureTaxonomyStateUnknown    = "reference_architecture_valc_failure_taxonomy_unknown"

	ReferenceArchitectureValCScenarioDescriptorStateActive     = "reference_architecture_valc_scenario_descriptor_active"
	ReferenceArchitectureValCScenarioDescriptorStatePartial    = "reference_architecture_valc_scenario_descriptor_partial"
	ReferenceArchitectureValCScenarioDescriptorStateIncomplete = "reference_architecture_valc_scenario_descriptor_incomplete"
	ReferenceArchitectureValCScenarioDescriptorStateBlocked    = "reference_architecture_valc_scenario_descriptor_blocked"
	ReferenceArchitectureValCScenarioDescriptorStateUnknown    = "reference_architecture_valc_scenario_descriptor_unknown"

	ReferenceArchitectureValCDegradedModeStateActive     = "reference_architecture_valc_degraded_mode_active"
	ReferenceArchitectureValCDegradedModeStatePartial    = "reference_architecture_valc_degraded_mode_partial"
	ReferenceArchitectureValCDegradedModeStateIncomplete = "reference_architecture_valc_degraded_mode_incomplete"
	ReferenceArchitectureValCDegradedModeStateBlocked    = "reference_architecture_valc_degraded_mode_blocked"
	ReferenceArchitectureValCDegradedModeStateUnknown    = "reference_architecture_valc_degraded_mode_unknown"

	ReferenceArchitectureValCRecoveryExpectationStateActive     = "reference_architecture_valc_recovery_expectation_active"
	ReferenceArchitectureValCRecoveryExpectationStatePartial    = "reference_architecture_valc_recovery_expectation_partial"
	ReferenceArchitectureValCRecoveryExpectationStateIncomplete = "reference_architecture_valc_recovery_expectation_incomplete"
	ReferenceArchitectureValCRecoveryExpectationStateBlocked    = "reference_architecture_valc_recovery_expectation_blocked"
	ReferenceArchitectureValCRecoveryExpectationStateUnknown    = "reference_architecture_valc_recovery_expectation_unknown"

	ReferenceArchitectureValCScalingScenarioStateActive     = "reference_architecture_valc_scaling_scenario_active"
	ReferenceArchitectureValCScalingScenarioStatePartial    = "reference_architecture_valc_scaling_scenario_partial"
	ReferenceArchitectureValCScalingScenarioStateIncomplete = "reference_architecture_valc_scaling_scenario_incomplete"
	ReferenceArchitectureValCScalingScenarioStateBlocked    = "reference_architecture_valc_scaling_scenario_blocked"
	ReferenceArchitectureValCScalingScenarioStateUnknown    = "reference_architecture_valc_scaling_scenario_unknown"

	ReferenceArchitectureValCTrustPathStateActive     = "reference_architecture_valc_trust_path_active"
	ReferenceArchitectureValCTrustPathStatePartial    = "reference_architecture_valc_trust_path_partial"
	ReferenceArchitectureValCTrustPathStateIncomplete = "reference_architecture_valc_trust_path_incomplete"
	ReferenceArchitectureValCTrustPathStateBlocked    = "reference_architecture_valc_trust_path_blocked"
	ReferenceArchitectureValCTrustPathStateUnknown    = "reference_architecture_valc_trust_path_unknown"

	ReferenceArchitectureValCAuditPathStateActive     = "reference_architecture_valc_audit_path_active"
	ReferenceArchitectureValCAuditPathStatePartial    = "reference_architecture_valc_audit_path_partial"
	ReferenceArchitectureValCAuditPathStateIncomplete = "reference_architecture_valc_audit_path_incomplete"
	ReferenceArchitectureValCAuditPathStateBlocked    = "reference_architecture_valc_audit_path_blocked"
	ReferenceArchitectureValCAuditPathStateUnknown    = "reference_architecture_valc_audit_path_unknown"

	ReferenceArchitectureValCControlPlaneStateActive     = "reference_architecture_valc_control_plane_active"
	ReferenceArchitectureValCControlPlaneStatePartial    = "reference_architecture_valc_control_plane_partial"
	ReferenceArchitectureValCControlPlaneStateIncomplete = "reference_architecture_valc_control_plane_incomplete"
	ReferenceArchitectureValCControlPlaneStateBlocked    = "reference_architecture_valc_control_plane_blocked"
	ReferenceArchitectureValCControlPlaneStateUnknown    = "reference_architecture_valc_control_plane_unknown"

	ReferenceArchitectureValCStateActive     = "reference_architecture_valc_active"
	ReferenceArchitectureValCStatePartial    = "reference_architecture_valc_partial"
	ReferenceArchitectureValCStateIncomplete = "reference_architecture_valc_incomplete"
	ReferenceArchitectureValCStateBlocked    = "reference_architecture_valc_blocked"
	ReferenceArchitectureValCStateUnknown    = "reference_architecture_valc_unknown"

	ReferenceArchitectureValCFailureTrustAnchorUnavailable = "trust_anchor_unavailable"
	ReferenceArchitectureValCFailureSigningPathDegraded    = "signing_path_degraded"
	ReferenceArchitectureValCFailureAuditWriterLatency     = "audit_writer_latency"
	ReferenceArchitectureValCFailureAuditWriterUnavailable = "audit_writer_unavailable"
	ReferenceArchitectureValCFailureEvidenceStorage        = "evidence_storage_degraded"
	ReferenceArchitectureValCFailurePolicyDistribution     = "policy_distribution_delayed"
	ReferenceArchitectureValCFailureConnectorDegraded      = "connector_degraded"
	ReferenceArchitectureValCFailurePartialStorageFailure  = "partial_storage_failure"
	ReferenceArchitectureValCFailureVerifierUnavailable    = "verifier_access_unavailable"
	ReferenceArchitectureValCFailureAirGapTransferDelayed  = "air_gap_transfer_delayed"
	ReferenceArchitectureValCFailureControlPlaneOverload   = "control_plane_overload"
	ReferenceArchitectureValCFailureDataPlaneDegraded      = "data_plane_degraded"
	ReferenceArchitectureValCFailureRecoveryUnavailable    = "recovery_path_unavailable"
	ReferenceArchitectureValCFailureBackupRestoreUnverf    = "backup_restore_unverified"
	ReferenceArchitectureValCFailureDependencyTimeout      = "dependency_timeout"
	ReferenceArchitectureValCFailureUnknown                = "unknown"

	ReferenceArchitectureValCScenarioReady        = "scenario_ready"
	ReferenceArchitectureValCBoundedDegraded      = "bounded_degraded"
	ReferenceArchitectureValCFailClosed           = "fail_closed"
	ReferenceArchitectureValCScenarioUnsupported  = "unsupported"
	ReferenceArchitectureValCScenarioStale        = "stale"
	ReferenceArchitectureValCScenarioUnknownState = "unknown"

	ReferenceArchitectureValCScalingControlPlaneCapacity  = "control_plane_capacity"
	ReferenceArchitectureValCScalingAuditWriteCapacity    = "audit_write_path_capacity"
	ReferenceArchitectureValCScalingEvidenceCapacity      = "evidence_storage_capacity"
	ReferenceArchitectureValCScalingVerificationRead      = "verification_read_path_capacity"
	ReferenceArchitectureValCScalingPolicyDistribution    = "policy_distribution_capacity"
	ReferenceArchitectureValCScalingConnectorBackpressure = "connector_backpressure"
	ReferenceArchitectureValCScalingRateLimit             = "rate_limit_behavior"
	ReferenceArchitectureValCScalingNoisyNeighbor         = "noisy_neighbor_boundary"
	ReferenceArchitectureValCScalingQueueBacklog          = "queue_backlog_behavior"
	ReferenceArchitectureValCScalingUnknown               = "unknown"
)

type ReferenceArchitectureResilienceScenarioPack struct {
	CurrentState            string                                   `json:"current_state"`
	ScenarioPackID          string                                   `json:"scenario_pack_id"`
	Version                 string                                   `json:"version"`
	BlueprintFamily         string                                   `json:"blueprint_family"`
	BlueprintID             string                                   `json:"blueprint_id"`
	PackRef                 string                                   `json:"pack_ref"`
	LifecycleState          string                                   `json:"lifecycle_state"`
	CompatibilityState      string                                   `json:"compatibility_state"`
	Owner                   string                                   `json:"owner"`
	TargetEnvironmentRef    string                                   `json:"target_environment_ref"`
	ResilienceScope         string                                   `json:"resilience_scope"`
	ScenarioRefs            []string                                 `json:"scenario_refs,omitempty"`
	ExpectedBehaviorRefs    []string                                 `json:"expected_behavior_refs,omitempty"`
	RecoveryExpectationRefs []string                                 `json:"recovery_expectation_refs,omitempty"`
	EvidenceRefs            []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	SupportBoundaryRef      string                                   `json:"support_boundary_ref"`
	Caveats                 []string                                 `json:"caveats,omitempty"`
	ProjectionDisclaimer    string                                   `json:"projection_disclaimer"`
	CreatedAt               string                                   `json:"created_at"`
	UpdatedAt               string                                   `json:"updated_at"`
	GuaranteedResilience    bool                                     `json:"guaranteed_resilience_claim_present"`
	CertifiedRecovery       bool                                     `json:"certified_recovery_claim_present"`
	ClaimsPoint6Pass        bool                                     `json:"claims_point_6_pass"`
}

type ReferenceArchitectureResilienceScenarioPackRegistry struct {
	CurrentState         string                                        `json:"current_state"`
	RegistryID           string                                        `json:"registry_id"`
	Version              string                                        `json:"version"`
	SupportedFamilies    []string                                      `json:"supported_families,omitempty"`
	ScenarioPacks        []ReferenceArchitectureResilienceScenarioPack `json:"scenario_packs,omitempty"`
	ProjectionDisclaimer string                                        `json:"projection_disclaimer"`
}

type ReferenceArchitectureFailureModeTaxonomy struct {
	CurrentState         string   `json:"current_state"`
	TaxonomyID           string   `json:"taxonomy_id"`
	SupportedCategories  []string `json:"supported_categories,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureScenarioDescriptor struct {
	ScenarioID                 string   `json:"scenario_id"`
	Category                   string   `json:"category"`
	BlueprintFamily            string   `json:"blueprint_family"`
	AffectedScope              string   `json:"affected_scope"`
	TriggerCondition           string   `json:"trigger_condition"`
	ExpectedState              string   `json:"expected_state"`
	ExpectedFailClosedBehavior string   `json:"expected_fail_closed_behavior"`
	ExpectedDegradedBehavior   string   `json:"expected_degraded_behavior"`
	ExpectedRecoveryBehavior   string   `json:"expected_recovery_behavior"`
	RequiredEvidenceTypes      []string `json:"required_evidence_types,omitempty"`
	FreshnessRequirement       string   `json:"freshness_requirement"`
	Severity                   string   `json:"severity"`
	BlocksMatched              bool     `json:"blocks_matched"`
	AdvisoryOnly               bool     `json:"advisory_only"`
	Caveats                    []string `json:"caveats,omitempty"`
}

type ReferenceArchitectureScenarioDescriptorPack struct {
	CurrentState         string                                    `json:"current_state"`
	PackID               string                                    `json:"pack_id"`
	BlueprintFamily      string                                    `json:"blueprint_family"`
	SupportedCategories  []string                                  `json:"supported_categories,omitempty"`
	SupportedStates      []string                                  `json:"supported_states,omitempty"`
	SupportedSeverities  []string                                  `json:"supported_severities,omitempty"`
	Scenarios            []ReferenceArchitectureScenarioDescriptor `json:"scenarios,omitempty"`
	ProjectionDisclaimer string                                    `json:"projection_disclaimer"`
}

type ReferenceArchitectureScenarioDescriptorCollection struct {
	CurrentState         string                                        `json:"current_state"`
	CollectionID         string                                        `json:"collection_id"`
	SupportedCategories  []string                                      `json:"supported_categories,omitempty"`
	SupportedStates      []string                                      `json:"supported_states,omitempty"`
	SupportedSeverities  []string                                      `json:"supported_severities,omitempty"`
	Packs                []ReferenceArchitectureScenarioDescriptorPack `json:"packs,omitempty"`
	ProjectionDisclaimer string                                        `json:"projection_disclaimer"`
}

type ReferenceArchitectureDegradedModeBehavior struct {
	DegradedModeID         string   `json:"degraded_mode_id"`
	ScenarioRef            string   `json:"scenario_ref"`
	AllowedOperations      []string `json:"allowed_operations,omitempty"`
	BlockedOperations      []string `json:"blocked_operations,omitempty"`
	RequiredOperatorAction string   `json:"required_operator_action"`
	EvidenceRequired       []string `json:"evidence_required,omitempty"`
	FreshnessExpectation   string   `json:"freshness_expectation"`
	RecoveryRequired       bool     `json:"recovery_required"`
	SupportBoundaryRef     string   `json:"support_boundary_ref"`
	Caveats                []string `json:"caveats,omitempty"`
	RedactionKeepsCaveats  bool     `json:"redaction_keeps_caveats"`
	UnsupportedBehavior    bool     `json:"unsupported_behavior"`
}

type ReferenceArchitectureDegradedModePack struct {
	CurrentState         string                                      `json:"current_state"`
	PackID               string                                      `json:"pack_id"`
	BlueprintFamily      string                                      `json:"blueprint_family"`
	Modes                []ReferenceArchitectureDegradedModeBehavior `json:"modes,omitempty"`
	ProjectionDisclaimer string                                      `json:"projection_disclaimer"`
}

type ReferenceArchitectureDegradedModeCollection struct {
	CurrentState         string                                  `json:"current_state"`
	CollectionID         string                                  `json:"collection_id"`
	Packs                []ReferenceArchitectureDegradedModePack `json:"packs,omitempty"`
	ProjectionDisclaimer string                                  `json:"projection_disclaimer"`
}

type ReferenceArchitectureRecoveryExpectation struct {
	RecoveryID                string   `json:"recovery_id"`
	ScenarioRef               string   `json:"scenario_ref"`
	ExpectedRecoveryPath      string   `json:"expected_recovery_path"`
	RequiredEvidenceTypes     []string `json:"required_evidence_types,omitempty"`
	OperatorActionRequired    string   `json:"operator_action_required"`
	RollbackOrRestoreBoundary string   `json:"rollback_or_restore_boundary"`
	VerificationRequired      bool     `json:"verification_required"`
	Timestamp                 string   `json:"timestamp"`
	FreshnessState            string   `json:"freshness_state"`
	SupportedEnvironmentScope string   `json:"supported_environment_scope"`
	UnsupportedConditions     []string `json:"unsupported_conditions,omitempty"`
	Caveats                   []string `json:"caveats,omitempty"`
	CertifiedRecoveryClaim    bool     `json:"certified_recovery_claim_present"`
}

type ReferenceArchitectureRecoveryExpectationPack struct {
	CurrentState         string                                     `json:"current_state"`
	PackID               string                                     `json:"pack_id"`
	BlueprintFamily      string                                     `json:"blueprint_family"`
	Expectations         []ReferenceArchitectureRecoveryExpectation `json:"expectations,omitempty"`
	ProjectionDisclaimer string                                     `json:"projection_disclaimer"`
}

type ReferenceArchitectureRecoveryExpectationCollection struct {
	CurrentState         string                                         `json:"current_state"`
	CollectionID         string                                         `json:"collection_id"`
	Packs                []ReferenceArchitectureRecoveryExpectationPack `json:"packs,omitempty"`
	ProjectionDisclaimer string                                         `json:"projection_disclaimer"`
}

type ReferenceArchitectureScalingScenarioDescriptor struct {
	ScalingScenarioID        string   `json:"scaling_scenario_id"`
	Category                 string   `json:"category"`
	BlueprintFamily          string   `json:"blueprint_family"`
	ExpectedCapacityBoundary string   `json:"expected_capacity_boundary"`
	DegradationThreshold     int      `json:"degradation_threshold"`
	FailClosedThreshold      int      `json:"fail_closed_threshold"`
	RequiredEvidenceTypes    []string `json:"required_evidence_types,omitempty"`
	Timestamp                string   `json:"timestamp"`
	FreshnessState           string   `json:"freshness_state"`
	Caveats                  []string `json:"caveats,omitempty"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
	PerformanceGuarantee     bool     `json:"performance_guarantee_claim_present"`
}

type ReferenceArchitectureScalingScenarioPack struct {
	CurrentState         string                                           `json:"current_state"`
	PackID               string                                           `json:"pack_id"`
	BlueprintFamily      string                                           `json:"blueprint_family"`
	SupportedCategories  []string                                         `json:"supported_categories,omitempty"`
	Scenarios            []ReferenceArchitectureScalingScenarioDescriptor `json:"scenarios,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type ReferenceArchitectureScalingScenarioCollection struct {
	CurrentState         string                                     `json:"current_state"`
	CollectionID         string                                     `json:"collection_id"`
	SupportedCategories  []string                                   `json:"supported_categories,omitempty"`
	Packs                []ReferenceArchitectureScalingScenarioPack `json:"packs,omitempty"`
	ProjectionDisclaimer string                                     `json:"projection_disclaimer"`
}

type ReferenceArchitectureTrustPathContinuityCheck struct {
	CurrentState                        string                                   `json:"current_state"`
	CheckID                             string                                   `json:"check_id"`
	BlueprintFamily                     string                                   `json:"blueprint_family"`
	TrustAnchorAvailabilityExpected     bool                                     `json:"trust_anchor_availability_expected"`
	StrictCustodyBoundaryExpected       bool                                     `json:"strict_custody_boundary_expected"`
	VerifierReplayVisibilityExpected    bool                                     `json:"verifier_replay_visibility_expected"`
	AirGappedLocalTrustExpected         bool                                     `json:"air_gapped_local_trust_expected"`
	RequiresLiveExternalTrustDependency bool                                     `json:"requires_live_external_trust_dependency"`
	EvidenceRefs                        []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer                string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureTrustPathCollection struct {
	CurrentState         string                                          `json:"current_state"`
	CollectionID         string                                          `json:"collection_id"`
	Checks               []ReferenceArchitectureTrustPathContinuityCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                          `json:"projection_disclaimer"`
}

type ReferenceArchitectureAuditPathDegradationCheck struct {
	CurrentState                    string                                   `json:"current_state"`
	CheckID                         string                                   `json:"check_id"`
	BlueprintFamily                 string                                   `json:"blueprint_family"`
	AuditWriterAvailabilityExpected bool                                     `json:"audit_writer_availability_expected"`
	AuditLatencyDegradedBehavior    string                                   `json:"audit_latency_degraded_behavior"`
	EvidenceCustodyPath             string                                   `json:"evidence_custody_path"`
	PartialFailureHandling          string                                   `json:"partial_failure_handling"`
	OperatorVisibilityRequired      bool                                     `json:"operator_visibility_required"`
	RecoveryExpectationRef          string                                   `json:"recovery_expectation_ref"`
	CanonicalFailureStatePreserved  bool                                     `json:"canonical_failure_state_preserved"`
	EvidenceRefs                    []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer            string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureAuditPathCollection struct {
	CurrentState         string                                           `json:"current_state"`
	CollectionID         string                                           `json:"collection_id"`
	Checks               []ReferenceArchitectureAuditPathDegradationCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type ReferenceArchitectureControlPlaneSafetyCheck struct {
	CurrentState               string   `json:"current_state"`
	CheckID                    string   `json:"check_id"`
	BlueprintFamily            string   `json:"blueprint_family"`
	OverloadBehavior           string   `json:"overload_behavior"`
	BackpressureSemantics      string   `json:"backpressure_semantics"`
	RateLimitBehavior          string   `json:"rate_limit_behavior"`
	DependencyTimeoutBehavior  string   `json:"dependency_timeout_behavior"`
	FailClosedBehavior         string   `json:"fail_closed_behavior"`
	OperatorActionRequired     string   `json:"operator_action_required"`
	EvidenceOutputRequirements []string `json:"evidence_output_requirements,omitempty"`
	AutomaticApproval          bool     `json:"automatic_approval"`
	AutomaticMutation          bool     `json:"automatic_mutation"`
	Caveats                    []string `json:"caveats,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureControlPlaneSafetyCollection struct {
	CurrentState         string                                         `json:"current_state"`
	CollectionID         string                                         `json:"collection_id"`
	Checks               []ReferenceArchitectureControlPlaneSafetyCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                         `json:"projection_disclaimer"`
}

func referenceArchitectureValCFailureModes() []string {
	return []string{
		ReferenceArchitectureValCFailureTrustAnchorUnavailable,
		ReferenceArchitectureValCFailureSigningPathDegraded,
		ReferenceArchitectureValCFailureAuditWriterLatency,
		ReferenceArchitectureValCFailureAuditWriterUnavailable,
		ReferenceArchitectureValCFailureEvidenceStorage,
		ReferenceArchitectureValCFailurePolicyDistribution,
		ReferenceArchitectureValCFailureConnectorDegraded,
		ReferenceArchitectureValCFailurePartialStorageFailure,
		ReferenceArchitectureValCFailureVerifierUnavailable,
		ReferenceArchitectureValCFailureAirGapTransferDelayed,
		ReferenceArchitectureValCFailureControlPlaneOverload,
		ReferenceArchitectureValCFailureDataPlaneDegraded,
		ReferenceArchitectureValCFailureRecoveryUnavailable,
		ReferenceArchitectureValCFailureBackupRestoreUnverf,
		ReferenceArchitectureValCFailureDependencyTimeout,
		ReferenceArchitectureValCFailureUnknown,
	}
}

func referenceArchitectureValCRequiredFailureModes() []string {
	return []string{
		ReferenceArchitectureValCFailureTrustAnchorUnavailable,
		ReferenceArchitectureValCFailureSigningPathDegraded,
		ReferenceArchitectureValCFailureAuditWriterLatency,
		ReferenceArchitectureValCFailureAuditWriterUnavailable,
		ReferenceArchitectureValCFailureEvidenceStorage,
		ReferenceArchitectureValCFailurePolicyDistribution,
		ReferenceArchitectureValCFailureConnectorDegraded,
		ReferenceArchitectureValCFailurePartialStorageFailure,
		ReferenceArchitectureValCFailureVerifierUnavailable,
		ReferenceArchitectureValCFailureAirGapTransferDelayed,
		ReferenceArchitectureValCFailureControlPlaneOverload,
		ReferenceArchitectureValCFailureDataPlaneDegraded,
		ReferenceArchitectureValCFailureRecoveryUnavailable,
		ReferenceArchitectureValCFailureBackupRestoreUnverf,
		ReferenceArchitectureValCFailureDependencyTimeout,
	}
}

func referenceArchitectureValCScenarioStates() []string {
	return []string{
		ReferenceArchitectureValCScenarioReady,
		ReferenceArchitectureValCBoundedDegraded,
		ReferenceArchitectureValCFailClosed,
		ReferenceArchitectureValCScenarioUnsupported,
		ReferenceArchitectureValCScenarioStale,
		ReferenceArchitectureValCScenarioUnknownState,
	}
}

func referenceArchitectureValCScalingCategories() []string {
	return []string{
		ReferenceArchitectureValCScalingControlPlaneCapacity,
		ReferenceArchitectureValCScalingAuditWriteCapacity,
		ReferenceArchitectureValCScalingEvidenceCapacity,
		ReferenceArchitectureValCScalingVerificationRead,
		ReferenceArchitectureValCScalingPolicyDistribution,
		ReferenceArchitectureValCScalingConnectorBackpressure,
		ReferenceArchitectureValCScalingRateLimit,
		ReferenceArchitectureValCScalingNoisyNeighbor,
		ReferenceArchitectureValCScalingQueueBacklog,
		ReferenceArchitectureValCScalingUnknown,
	}
}

func referenceArchitectureValCRequiredScalingCategories() []string {
	return []string{
		ReferenceArchitectureValCScalingControlPlaneCapacity,
		ReferenceArchitectureValCScalingAuditWriteCapacity,
		ReferenceArchitectureValCScalingEvidenceCapacity,
		ReferenceArchitectureValCScalingVerificationRead,
		ReferenceArchitectureValCScalingPolicyDistribution,
		ReferenceArchitectureValCScalingConnectorBackpressure,
		ReferenceArchitectureValCScalingRateLimit,
		ReferenceArchitectureValCScalingNoisyNeighbor,
		ReferenceArchitectureValCScalingQueueBacklog,
	}
}

func referenceArchitectureValCScenarioSeverities() []string {
	return []string{
		ReferenceArchitectureValBSeverityCritical,
		ReferenceArchitectureValBSeverityHigh,
		ReferenceArchitectureValBSeverityMedium,
		ReferenceArchitectureValBSeverityLow,
	}
}

func referenceArchitectureValCProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_resilience_scaling_hardening"
}

func referenceArchitectureValCHasProjectionDisclaimer(value string) bool {
	return strings.Contains(strings.TrimSpace(value), "projection_only") &&
		strings.Contains(strings.TrimSpace(value), "not_canonical_truth")
}

func referenceArchitectureValCScenarioExpectedState(category string) string {
	switch strings.TrimSpace(category) {
	case ReferenceArchitectureValCFailureTrustAnchorUnavailable,
		ReferenceArchitectureValCFailureAuditWriterUnavailable,
		ReferenceArchitectureValCFailureDependencyTimeout:
		return ReferenceArchitectureValCFailClosed
	case ReferenceArchitectureValCFailureRecoveryUnavailable,
		ReferenceArchitectureValCFailureBackupRestoreUnverf:
		return ReferenceArchitectureValCScenarioUnsupported
	default:
		return ReferenceArchitectureValCBoundedDegraded
	}
}

func referenceArchitectureValCScenarioSeverity(category string) string {
	switch strings.TrimSpace(category) {
	case ReferenceArchitectureValCFailureTrustAnchorUnavailable,
		ReferenceArchitectureValCFailureAuditWriterUnavailable,
		ReferenceArchitectureValCFailureRecoveryUnavailable,
		ReferenceArchitectureValCFailureBackupRestoreUnverf:
		return ReferenceArchitectureValBSeverityCritical
	case ReferenceArchitectureValCFailureControlPlaneOverload,
		ReferenceArchitectureValCFailureDependencyTimeout,
		ReferenceArchitectureValCFailureSigningPathDegraded:
		return ReferenceArchitectureValBSeverityHigh
	default:
		return ReferenceArchitectureValBSeverityMedium
	}
}

func referenceArchitectureValCScenarioBlocksMatched(category string) bool {
	switch strings.TrimSpace(category) {
	case ReferenceArchitectureValCFailureTrustAnchorUnavailable,
		ReferenceArchitectureValCFailureAuditWriterUnavailable,
		ReferenceArchitectureValCFailureRecoveryUnavailable,
		ReferenceArchitectureValCFailureBackupRestoreUnverf,
		ReferenceArchitectureValCFailureDependencyTimeout:
		return true
	default:
		return false
	}
}

func referenceArchitectureValCScenarioEvidenceTypes(category string) []string {
	evidenceTypes := []string{
		ReferenceArchitectureEvidenceDeploymentObservation,
		ReferenceArchitectureEvidenceCapabilityAttestation,
		ReferenceArchitectureEvidenceCompatibilityReport,
	}
	switch strings.TrimSpace(category) {
	case ReferenceArchitectureValCFailureAuditWriterLatency,
		ReferenceArchitectureValCFailureAuditWriterUnavailable:
		evidenceTypes = append(evidenceTypes, ReferenceArchitectureEvidenceAuditSnapshot)
	case ReferenceArchitectureValCFailureTrustAnchorUnavailable:
		evidenceTypes = append(evidenceTypes, ReferenceArchitectureEvidenceSupportBoundary)
	}
	return evidenceTypes
}

func referenceArchitectureValCScenarioAffectedScope(category string) string {
	switch strings.TrimSpace(category) {
	case ReferenceArchitectureValCFailureTrustAnchorUnavailable,
		ReferenceArchitectureValCFailureSigningPathDegraded:
		return "trust_path"
	case ReferenceArchitectureValCFailureAuditWriterLatency,
		ReferenceArchitectureValCFailureAuditWriterUnavailable,
		ReferenceArchitectureValCFailureEvidenceStorage:
		return "audit_path"
	case ReferenceArchitectureValCFailureControlPlaneOverload,
		ReferenceArchitectureValCFailureDependencyTimeout:
		return "control_plane"
	default:
		return "platform_boundary"
	}
}

func referenceArchitectureValCScenarioPackEvidenceRefs(pack ReferenceArchitectureBlueprintPack) []ReferenceArchitectureEvidenceReference {
	refs := append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...)
	refs = append(refs, ReferenceArchitectureEvidenceReference{
		EvidenceID:     "resilience/" + pack.BlueprintFamily + "/continuity",
		EvidenceType:   ReferenceArchitectureEvidenceAuditSnapshot,
		Source:         "resilience-scenario-pack",
		Timestamp:      pack.UpdatedAt,
		FreshnessState: IntelligenceCalibrationFreshnessFresh,
		Scope:          pack.BlueprintFamily,
		Caveats:        []string{"bounded resilience scenario evidence only"},
	})
	return refs
}

func referenceArchitectureValCScenarioRef(family, category string) string {
	return "scenario/" + strings.TrimSpace(family) + "/" + strings.TrimSpace(category)
}

func referenceArchitectureValCExpectedBehaviorRef(family, category string) string {
	return "expected-behavior/" + strings.TrimSpace(family) + "/" + strings.TrimSpace(category)
}

func referenceArchitectureValCRecoveryRef(family, category string) string {
	return "recovery/" + strings.TrimSpace(family) + "/" + strings.TrimSpace(category)
}

func referenceArchitectureValCScenarioPackForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureResilienceScenarioPack {
	scenarioRefs := make([]string, 0, len(referenceArchitectureValCRequiredFailureModes()))
	expectedBehaviorRefs := make([]string, 0, len(referenceArchitectureValCRequiredFailureModes()))
	recoveryRefs := make([]string, 0, len(referenceArchitectureValCRequiredFailureModes()))
	for _, category := range referenceArchitectureValCRequiredFailureModes() {
		scenarioRefs = append(scenarioRefs, referenceArchitectureValCScenarioRef(pack.BlueprintFamily, category))
		expectedBehaviorRefs = append(expectedBehaviorRefs, referenceArchitectureValCExpectedBehaviorRef(pack.BlueprintFamily, category))
		recoveryRefs = append(recoveryRefs, referenceArchitectureValCRecoveryRef(pack.BlueprintFamily, category))
	}
	return ReferenceArchitectureResilienceScenarioPack{
		CurrentState:            "reference_architecture_valc_scenario_pack_ready",
		ScenarioPackID:          "resilience-pack/" + pack.BlueprintFamily,
		Version:                 pack.Version,
		BlueprintFamily:         pack.BlueprintFamily,
		BlueprintID:             pack.BlueprintID,
		PackRef:                 pack.PackID,
		LifecycleState:          pack.LifecycleState,
		CompatibilityState:      pack.CompatibilityState,
		Owner:                   pack.Owner,
		TargetEnvironmentRef:    pack.TargetEnvironmentRef,
		ResilienceScope:         "bounded_resilience_and_scaling_contracts",
		ScenarioRefs:            scenarioRefs,
		ExpectedBehaviorRefs:    expectedBehaviorRefs,
		RecoveryExpectationRefs: recoveryRefs,
		EvidenceRefs:            referenceArchitectureValCScenarioPackEvidenceRefs(pack),
		SupportBoundaryRef:      pack.SupportBoundaryRef,
		Caveats:                 []string{"scenario pack remains advisory projection and does not execute chaos or recovery"},
		ProjectionDisclaimer:    referenceArchitectureValCProjectionDisclaimer(),
		CreatedAt:               pack.CreatedAt,
		UpdatedAt:               pack.UpdatedAt,
	}
}

func ReferenceArchitectureValCScenarioPackRegistry() ReferenceArchitectureResilienceScenarioPackRegistry {
	registry := ReferenceArchitectureValBPackRegistry()
	packs := make([]ReferenceArchitectureResilienceScenarioPack, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		packs = append(packs, referenceArchitectureValCScenarioPackForPack(pack))
	}
	return ReferenceArchitectureResilienceScenarioPackRegistry{
		CurrentState:         "reference_architecture_valc_scenario_pack_registry_ready",
		RegistryID:           "reference-architecture-valc-resilience-packs",
		Version:              "1.0.0",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		ScenarioPacks:        packs,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCFailureModeTaxonomy() ReferenceArchitectureFailureModeTaxonomy {
	return ReferenceArchitectureFailureModeTaxonomy{
		CurrentState:         "reference_architecture_valc_failure_taxonomy_ready",
		TaxonomyID:           "reference-architecture-valc-failure-taxonomy",
		SupportedCategories:  referenceArchitectureValCFailureModes(),
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCScenarioDescriptorForPack(pack ReferenceArchitectureResilienceScenarioPack, category string) ReferenceArchitectureScenarioDescriptor {
	return ReferenceArchitectureScenarioDescriptor{
		ScenarioID:                 referenceArchitectureValCScenarioRef(pack.BlueprintFamily, category),
		Category:                   category,
		BlueprintFamily:            pack.BlueprintFamily,
		AffectedScope:              referenceArchitectureValCScenarioAffectedScope(category),
		TriggerCondition:           "bounded scenario condition for " + category,
		ExpectedState:              referenceArchitectureValCScenarioExpectedState(category),
		ExpectedFailClosedBehavior: "preserve canonical failure semantics and deny optimistic matched state",
		ExpectedDegradedBehavior:   "surface bounded degraded mode with operator action and evidence retention",
		ExpectedRecoveryBehavior:   "bind recovery expectation contract before returning normal service assumptions",
		RequiredEvidenceTypes:      referenceArchitectureValCScenarioEvidenceTypes(category),
		FreshnessRequirement:       "fresh_rfc3339_evidence_required",
		Severity:                   referenceArchitectureValCScenarioSeverity(category),
		BlocksMatched:              referenceArchitectureValCScenarioBlocksMatched(category),
		AdvisoryOnly:               category == ReferenceArchitectureValCFailurePolicyDistribution || category == ReferenceArchitectureValCFailureConnectorDegraded,
		Caveats:                    append([]string{}, pack.Caveats...),
	}
}

func referenceArchitectureValCScenarioDescriptorPackForPack(pack ReferenceArchitectureResilienceScenarioPack) ReferenceArchitectureScenarioDescriptorPack {
	scenarios := make([]ReferenceArchitectureScenarioDescriptor, 0, len(referenceArchitectureValCRequiredFailureModes()))
	for _, category := range referenceArchitectureValCRequiredFailureModes() {
		scenarios = append(scenarios, referenceArchitectureValCScenarioDescriptorForPack(pack, category))
	}
	return ReferenceArchitectureScenarioDescriptorPack{
		CurrentState:         "reference_architecture_valc_scenario_descriptor_pack_ready",
		PackID:               "scenario-descriptors/" + pack.BlueprintFamily,
		BlueprintFamily:      pack.BlueprintFamily,
		SupportedCategories:  referenceArchitectureValCFailureModes(),
		SupportedStates:      referenceArchitectureValCScenarioStates(),
		SupportedSeverities:  referenceArchitectureValCScenarioSeverities(),
		Scenarios:            scenarios,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCScenarioDescriptorCollection() ReferenceArchitectureScenarioDescriptorCollection {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	packs := make([]ReferenceArchitectureScenarioDescriptorPack, 0, len(registry.ScenarioPacks))
	for _, pack := range registry.ScenarioPacks {
		packs = append(packs, referenceArchitectureValCScenarioDescriptorPackForPack(pack))
	}
	return ReferenceArchitectureScenarioDescriptorCollection{
		CurrentState:         "reference_architecture_valc_scenario_descriptor_collection_ready",
		CollectionID:         "reference-architecture-valc-scenario-descriptors",
		SupportedCategories:  referenceArchitectureValCFailureModes(),
		SupportedStates:      referenceArchitectureValCScenarioStates(),
		SupportedSeverities:  referenceArchitectureValCScenarioSeverities(),
		Packs:                packs,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCDegradedModeForScenario(pack ReferenceArchitectureResilienceScenarioPack, scenario ReferenceArchitectureScenarioDescriptor) ReferenceArchitectureDegradedModeBehavior {
	return ReferenceArchitectureDegradedModeBehavior{
		DegradedModeID:         "degraded-mode/" + pack.BlueprintFamily + "/" + scenario.Category,
		ScenarioRef:            scenario.ScenarioID,
		AllowedOperations:      []string{"inspect_status", "collect_evidence", "execute_bounded_recovery_review"},
		BlockedOperations:      []string{"approve_deployment", "suppress_canonical_failure", "relax_support_boundary"},
		RequiredOperatorAction: "follow bounded runbook for " + scenario.Category,
		EvidenceRequired:       append([]string{}, scenario.RequiredEvidenceTypes...),
		FreshnessExpectation:   scenario.FreshnessRequirement,
		RecoveryRequired:       true,
		SupportBoundaryRef:     pack.SupportBoundaryRef,
		Caveats:                append([]string{}, pack.Caveats...),
		RedactionKeepsCaveats:  true,
		UnsupportedBehavior:    false,
	}
}

func referenceArchitectureValCDegradedModePackForDescriptors(pack ReferenceArchitectureResilienceScenarioPack, descriptorPack ReferenceArchitectureScenarioDescriptorPack) ReferenceArchitectureDegradedModePack {
	modes := make([]ReferenceArchitectureDegradedModeBehavior, 0, len(descriptorPack.Scenarios))
	for _, scenario := range descriptorPack.Scenarios {
		modes = append(modes, referenceArchitectureValCDegradedModeForScenario(pack, scenario))
	}
	return ReferenceArchitectureDegradedModePack{
		CurrentState:         "reference_architecture_valc_degraded_mode_pack_ready",
		PackID:               "degraded-modes/" + pack.BlueprintFamily,
		BlueprintFamily:      pack.BlueprintFamily,
		Modes:                modes,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCDegradedModeCollection() ReferenceArchitectureDegradedModeCollection {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	descriptorCollection := ReferenceArchitectureValCScenarioDescriptorCollection()
	packs := make([]ReferenceArchitectureDegradedModePack, 0, len(registry.ScenarioPacks))
	for idx, pack := range registry.ScenarioPacks {
		packs = append(packs, referenceArchitectureValCDegradedModePackForDescriptors(pack, descriptorCollection.Packs[idx]))
	}
	return ReferenceArchitectureDegradedModeCollection{
		CurrentState:         "reference_architecture_valc_degraded_mode_collection_ready",
		CollectionID:         "reference-architecture-valc-degraded-modes",
		Packs:                packs,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCRecoveryExpectationForScenario(pack ReferenceArchitectureResilienceScenarioPack, scenario ReferenceArchitectureScenarioDescriptor) ReferenceArchitectureRecoveryExpectation {
	return ReferenceArchitectureRecoveryExpectation{
		RecoveryID:                referenceArchitectureValCRecoveryRef(pack.BlueprintFamily, scenario.Category),
		ScenarioRef:               scenario.ScenarioID,
		ExpectedRecoveryPath:      "recover/" + pack.BlueprintFamily + "/" + scenario.Category,
		RequiredEvidenceTypes:     append([]string{}, scenario.RequiredEvidenceTypes...),
		OperatorActionRequired:    "operator review and recovery verification for " + scenario.Category,
		RollbackOrRestoreBoundary: pack.SupportBoundaryRef,
		VerificationRequired:      true,
		Timestamp:                 pack.UpdatedAt,
		FreshnessState:            IntelligenceCalibrationFreshnessFresh,
		SupportedEnvironmentScope: pack.TargetEnvironmentRef,
		UnsupportedConditions:     []string{"missing recovery path evidence", "missing rollback boundary"},
		Caveats:                   append([]string{}, pack.Caveats...),
	}
}

func referenceArchitectureValCRecoveryExpectationPackForDescriptors(pack ReferenceArchitectureResilienceScenarioPack, descriptorPack ReferenceArchitectureScenarioDescriptorPack) ReferenceArchitectureRecoveryExpectationPack {
	expectations := make([]ReferenceArchitectureRecoveryExpectation, 0, len(descriptorPack.Scenarios))
	for _, scenario := range descriptorPack.Scenarios {
		expectations = append(expectations, referenceArchitectureValCRecoveryExpectationForScenario(pack, scenario))
	}
	return ReferenceArchitectureRecoveryExpectationPack{
		CurrentState:         "reference_architecture_valc_recovery_expectation_pack_ready",
		PackID:               "recovery-expectations/" + pack.BlueprintFamily,
		BlueprintFamily:      pack.BlueprintFamily,
		Expectations:         expectations,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCRecoveryExpectationCollection() ReferenceArchitectureRecoveryExpectationCollection {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	descriptorCollection := ReferenceArchitectureValCScenarioDescriptorCollection()
	packs := make([]ReferenceArchitectureRecoveryExpectationPack, 0, len(registry.ScenarioPacks))
	for idx, pack := range registry.ScenarioPacks {
		packs = append(packs, referenceArchitectureValCRecoveryExpectationPackForDescriptors(pack, descriptorCollection.Packs[idx]))
	}
	return ReferenceArchitectureRecoveryExpectationCollection{
		CurrentState:         "reference_architecture_valc_recovery_expectation_collection_ready",
		CollectionID:         "reference-architecture-valc-recovery-expectations",
		Packs:                packs,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCScalingThresholds(family string) (int, int) {
	switch strings.TrimSpace(family) {
	case ReferenceArchitectureFamilyPerformanceSensitive:
		return 150, 220
	case ReferenceArchitectureFamilyHighAssurance:
		return 90, 130
	default:
		return 100, 150
	}
}

func referenceArchitectureValCScalingScenarioForPack(pack ReferenceArchitectureResilienceScenarioPack, category string) ReferenceArchitectureScalingScenarioDescriptor {
	degradeThreshold, failClosedThreshold := referenceArchitectureValCScalingThresholds(pack.BlueprintFamily)
	return ReferenceArchitectureScalingScenarioDescriptor{
		ScalingScenarioID:        "scaling/" + pack.BlueprintFamily + "/" + category,
		Category:                 category,
		BlueprintFamily:          pack.BlueprintFamily,
		ExpectedCapacityBoundary: "bounded_capacity_" + category,
		DegradationThreshold:     degradeThreshold,
		FailClosedThreshold:      failClosedThreshold,
		RequiredEvidenceTypes: []string{
			ReferenceArchitectureEvidenceCapabilityAttestation,
			ReferenceArchitectureEvidenceCompatibilityReport,
		},
		Timestamp:            pack.UpdatedAt,
		FreshnessState:       IntelligenceCalibrationFreshnessFresh,
		Caveats:              append([]string{}, pack.Caveats...),
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCScalingScenarioPackForPack(pack ReferenceArchitectureResilienceScenarioPack) ReferenceArchitectureScalingScenarioPack {
	scenarios := make([]ReferenceArchitectureScalingScenarioDescriptor, 0, len(referenceArchitectureValCRequiredScalingCategories()))
	for _, category := range referenceArchitectureValCRequiredScalingCategories() {
		scenarios = append(scenarios, referenceArchitectureValCScalingScenarioForPack(pack, category))
	}
	return ReferenceArchitectureScalingScenarioPack{
		CurrentState:         "reference_architecture_valc_scaling_scenario_pack_ready",
		PackID:               "scaling-scenarios/" + pack.BlueprintFamily,
		BlueprintFamily:      pack.BlueprintFamily,
		SupportedCategories:  referenceArchitectureValCScalingCategories(),
		Scenarios:            scenarios,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCScalingScenarioCollection() ReferenceArchitectureScalingScenarioCollection {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	packs := make([]ReferenceArchitectureScalingScenarioPack, 0, len(registry.ScenarioPacks))
	for _, pack := range registry.ScenarioPacks {
		packs = append(packs, referenceArchitectureValCScalingScenarioPackForPack(pack))
	}
	return ReferenceArchitectureScalingScenarioCollection{
		CurrentState:         "reference_architecture_valc_scaling_scenario_collection_ready",
		CollectionID:         "reference-architecture-valc-scaling-scenarios",
		SupportedCategories:  referenceArchitectureValCScalingCategories(),
		Packs:                packs,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCTrustPathCheckForProfile(profile ReferenceArchitectureBlueprintFamilyProfile, pack ReferenceArchitectureResilienceScenarioPack) ReferenceArchitectureTrustPathContinuityCheck {
	return ReferenceArchitectureTrustPathContinuityCheck{
		CurrentState:                        "reference_architecture_valc_trust_path_ready",
		CheckID:                             "trust-path/" + profile.Family,
		BlueprintFamily:                     profile.Family,
		TrustAnchorAvailabilityExpected:     true,
		StrictCustodyBoundaryExpected:       profile.StrongerTrustAnchorMode,
		VerifierReplayVisibilityExpected:    true,
		AirGappedLocalTrustExpected:         profile.LocalTrustAnchorAssumptionRequired,
		RequiresLiveExternalTrustDependency: false,
		EvidenceRefs:                        append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		ProjectionDisclaimer:                referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCTrustPathCollection() ReferenceArchitectureTrustPathCollection {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	checks := make([]ReferenceArchitectureTrustPathContinuityCheck, 0, len(registry.ScenarioPacks))
	for _, pack := range registry.ScenarioPacks {
		profile, _ := LookupReferenceArchitectureValAFamilyProfile(pack.BlueprintFamily)
		checks = append(checks, referenceArchitectureValCTrustPathCheckForProfile(profile, pack))
	}
	return ReferenceArchitectureTrustPathCollection{
		CurrentState:         "reference_architecture_valc_trust_path_collection_ready",
		CollectionID:         "reference-architecture-valc-trust-path",
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCAuditPathCheckForPack(pack ReferenceArchitectureResilienceScenarioPack) ReferenceArchitectureAuditPathDegradationCheck {
	return ReferenceArchitectureAuditPathDegradationCheck{
		CurrentState:                    "reference_architecture_valc_audit_path_ready",
		CheckID:                         "audit-path/" + pack.BlueprintFamily,
		BlueprintFamily:                 pack.BlueprintFamily,
		AuditWriterAvailabilityExpected: true,
		AuditLatencyDegradedBehavior:    "bounded latency surfaces degraded mode and preserves evidence capture expectations",
		EvidenceCustodyPath:             "evidence-custody/" + pack.BlueprintFamily,
		PartialFailureHandling:          "retain canonical failure state and require operator visibility on partial audit failures",
		OperatorVisibilityRequired:      true,
		RecoveryExpectationRef:          referenceArchitectureValCRecoveryRef(pack.BlueprintFamily, ReferenceArchitectureValCFailureAuditWriterUnavailable),
		CanonicalFailureStatePreserved:  true,
		EvidenceRefs:                    append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		ProjectionDisclaimer:            referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCAuditPathCollection() ReferenceArchitectureAuditPathCollection {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	checks := make([]ReferenceArchitectureAuditPathDegradationCheck, 0, len(registry.ScenarioPacks))
	for _, pack := range registry.ScenarioPacks {
		checks = append(checks, referenceArchitectureValCAuditPathCheckForPack(pack))
	}
	return ReferenceArchitectureAuditPathCollection{
		CurrentState:         "reference_architecture_valc_audit_path_collection_ready",
		CollectionID:         "reference-architecture-valc-audit-path",
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func referenceArchitectureValCControlPlaneCheckForProfile(profile ReferenceArchitectureBlueprintFamilyProfile, pack ReferenceArchitectureResilienceScenarioPack) ReferenceArchitectureControlPlaneSafetyCheck {
	overloadBehavior := "bounded_degraded_with_backpressure"
	if profile.PerformanceEnvelopeRequired {
		overloadBehavior = "bounded_degraded_with_priority_backpressure"
	}
	return ReferenceArchitectureControlPlaneSafetyCheck{
		CurrentState:              "reference_architecture_valc_control_plane_ready",
		CheckID:                   "control-plane/" + pack.BlueprintFamily,
		BlueprintFamily:           pack.BlueprintFamily,
		OverloadBehavior:          overloadBehavior,
		BackpressureSemantics:     "bounded_backpressure_preserves_fail_closed_state",
		RateLimitBehavior:         "rate_limit_behaviour_requires_explicit_operator_visibility",
		DependencyTimeoutBehavior: ReferenceArchitectureValCFailClosed,
		FailClosedBehavior:        "retain canonical denial and evidence emission under timeout or overload",
		OperatorActionRequired:    "operator reviews backlog and dependency health before resuming normal assumptions",
		EvidenceOutputRequirements: []string{
			ReferenceArchitectureEvidenceDeploymentObservation,
			ReferenceArchitectureEvidenceAuditSnapshot,
		},
		Caveats:              append([]string{}, pack.Caveats...),
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValCControlPlaneSafetyCollection() ReferenceArchitectureControlPlaneSafetyCollection {
	registry := ReferenceArchitectureValCScenarioPackRegistry()
	checks := make([]ReferenceArchitectureControlPlaneSafetyCheck, 0, len(registry.ScenarioPacks))
	for _, pack := range registry.ScenarioPacks {
		profile, _ := LookupReferenceArchitectureValAFamilyProfile(pack.BlueprintFamily)
		checks = append(checks, referenceArchitectureValCControlPlaneCheckForProfile(profile, pack))
	}
	return ReferenceArchitectureControlPlaneSafetyCollection{
		CurrentState:         "reference_architecture_valc_control_plane_collection_ready",
		CollectionID:         "reference-architecture-valc-control-plane",
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValCProjectionDisclaimer(),
	}
}

func EvaluateReferenceArchitectureValCScenarioPackState(pack ReferenceArchitectureResilienceScenarioPack) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		pack.ScenarioPackID,
		pack.Version,
		pack.BlueprintFamily,
		pack.BlueprintID,
		pack.PackRef,
		pack.LifecycleState,
		pack.CompatibilityState,
		pack.Owner,
		pack.TargetEnvironmentRef,
		pack.ResilienceScope,
		pack.SupportBoundaryRef,
		pack.ProjectionDisclaimer,
		pack.CreatedAt,
		pack.UpdatedAt,
	) {
		return ReferenceArchitectureValCScenarioPackStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), pack.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0LifecycleStates(), pack.LifecycleState) ||
		!containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), pack.CompatibilityState) ||
		!referenceArchitectureValCHasProjectionDisclaimer(pack.ProjectionDisclaimer) {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	if pack.GuaranteedResilience || pack.CertifiedRecovery || pack.ClaimsPoint6Pass {
		return ReferenceArchitectureValCScenarioPackStateBlocked
	}
	if len(pack.ScenarioRefs) != len(referenceArchitectureValCRequiredFailureModes()) ||
		len(pack.ExpectedBehaviorRefs) != len(pack.ScenarioRefs) ||
		len(pack.RecoveryExpectationRefs) != len(pack.ScenarioRefs) {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(pack.CreatedAt); !ok {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(pack.UpdatedAt); !ok {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	if pack.LifecycleState != ReferenceArchitectureLifecycleActive || pack.CompatibilityState != ReferenceArchitectureCompatibilityCompatible {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	valBRegistry := ReferenceArchitectureValBPackRegistry()
	found := false
	for _, candidate := range valBRegistry.Packs {
		if strings.TrimSpace(candidate.PackID) == strings.TrimSpace(pack.PackRef) &&
			strings.TrimSpace(candidate.BlueprintFamily) == strings.TrimSpace(pack.BlueprintFamily) &&
			strings.TrimSpace(candidate.BlueprintID) == strings.TrimSpace(pack.BlueprintID) &&
			strings.TrimSpace(candidate.SupportBoundaryRef) == strings.TrimSpace(pack.SupportBoundaryRef) {
			found = true
			break
		}
	}
	if !found {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(pack.EvidenceRefs)
	if !ok || !allFresh || stale {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	return ReferenceArchitectureValCScenarioPackStateActive
}

func EvaluateReferenceArchitectureValCScenarioPackRegistryState(registry ReferenceArchitectureResilienceScenarioPackRegistry) string {
	if strings.TrimSpace(registry.RegistryID) == "" || strings.TrimSpace(registry.Version) == "" || strings.TrimSpace(registry.ProjectionDisclaimer) == "" || len(registry.ScenarioPacks) == 0 {
		return ReferenceArchitectureValCScenarioPackStateIncomplete
	}
	if !containsExactTrimmedStringSet(registry.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValCHasProjectionDisclaimer(registry.ProjectionDisclaimer) ||
		len(registry.ScenarioPacks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCScenarioPackStatePartial
	}
	seenFamilies := map[string]struct{}{}
	seenPacks := map[string]struct{}{}
	for _, pack := range registry.ScenarioPacks {
		if _, ok := seenFamilies[strings.TrimSpace(pack.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCScenarioPackStatePartial
		}
		if _, ok := seenPacks[strings.TrimSpace(pack.ScenarioPackID)]; ok {
			return ReferenceArchitectureValCScenarioPackStatePartial
		}
		seenFamilies[strings.TrimSpace(pack.BlueprintFamily)] = struct{}{}
		seenPacks[strings.TrimSpace(pack.ScenarioPackID)] = struct{}{}
		if EvaluateReferenceArchitectureValCScenarioPackState(pack) != ReferenceArchitectureValCScenarioPackStateActive {
			return ReferenceArchitectureValCScenarioPackStatePartial
		}
	}
	return ReferenceArchitectureValCScenarioPackStateActive
}

func EvaluateReferenceArchitectureValCFailureModeTaxonomyState(taxonomy ReferenceArchitectureFailureModeTaxonomy) string {
	if strings.TrimSpace(taxonomy.TaxonomyID) == "" || strings.TrimSpace(taxonomy.ProjectionDisclaimer) == "" {
		return ReferenceArchitectureValCFailureTaxonomyStateIncomplete
	}
	if !containsExactTrimmedStringSet(taxonomy.SupportedCategories, referenceArchitectureValCFailureModes()...) ||
		!referenceArchitectureValCHasProjectionDisclaimer(taxonomy.ProjectionDisclaimer) {
		return ReferenceArchitectureValCFailureTaxonomyStatePartial
	}
	return ReferenceArchitectureValCFailureTaxonomyStateActive
}

func EvaluateReferenceArchitectureValCScenarioDescriptorPackState(pack ReferenceArchitectureScenarioDescriptorPack) string {
	if !referenceArchitectureValBRequiredRefsPresent(pack.PackID, pack.BlueprintFamily, pack.ProjectionDisclaimer) || len(pack.Scenarios) == 0 {
		return ReferenceArchitectureValCScenarioDescriptorStateIncomplete
	}
	if !containsExactTrimmedStringSet(pack.SupportedCategories, referenceArchitectureValCFailureModes()...) ||
		!containsExactTrimmedStringSet(pack.SupportedStates, referenceArchitectureValCScenarioStates()...) ||
		!containsExactTrimmedStringSet(pack.SupportedSeverities, referenceArchitectureValCScenarioSeverities()...) ||
		!referenceArchitectureValCHasProjectionDisclaimer(pack.ProjectionDisclaimer) {
		return ReferenceArchitectureValCScenarioDescriptorStatePartial
	}
	seenIDs := map[string]struct{}{}
	seenCategories := map[string]struct{}{}
	for _, scenario := range pack.Scenarios {
		if !referenceArchitectureValBRequiredRefsPresent(
			scenario.ScenarioID,
			scenario.Category,
			scenario.BlueprintFamily,
			scenario.AffectedScope,
			scenario.TriggerCondition,
			scenario.ExpectedState,
			scenario.ExpectedFailClosedBehavior,
			scenario.ExpectedDegradedBehavior,
			scenario.ExpectedRecoveryBehavior,
			scenario.FreshnessRequirement,
			scenario.Severity,
		) || len(scenario.RequiredEvidenceTypes) == 0 {
			return ReferenceArchitectureValCScenarioDescriptorStateIncomplete
		}
		if _, ok := seenIDs[strings.TrimSpace(scenario.ScenarioID)]; ok {
			return ReferenceArchitectureValCScenarioDescriptorStatePartial
		}
		if _, ok := seenCategories[strings.TrimSpace(scenario.Category)]; ok {
			return ReferenceArchitectureValCScenarioDescriptorStatePartial
		}
		seenIDs[strings.TrimSpace(scenario.ScenarioID)] = struct{}{}
		seenCategories[strings.TrimSpace(scenario.Category)] = struct{}{}
		if strings.TrimSpace(scenario.BlueprintFamily) != strings.TrimSpace(pack.BlueprintFamily) ||
			!containsTrimmedString(referenceArchitectureValCRequiredFailureModes(), scenario.Category) ||
			!containsTrimmedString(referenceArchitectureValCScenarioStates(), scenario.ExpectedState) ||
			strings.TrimSpace(scenario.ExpectedState) == ReferenceArchitectureValCScenarioUnknownState ||
			!containsTrimmedString(referenceArchitectureValCScenarioSeverities(), scenario.Severity) ||
			!containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), scenario.RequiredEvidenceTypes...) {
			return ReferenceArchitectureValCScenarioDescriptorStatePartial
		}
		if scenario.BlocksMatched && strings.TrimSpace(scenario.ExpectedState) == ReferenceArchitectureValCScenarioReady {
			return ReferenceArchitectureValCScenarioDescriptorStatePartial
		}
	}
	if len(seenCategories) != len(referenceArchitectureValCRequiredFailureModes()) {
		return ReferenceArchitectureValCScenarioDescriptorStatePartial
	}
	return ReferenceArchitectureValCScenarioDescriptorStateActive
}

func EvaluateReferenceArchitectureValCScenarioDescriptorCollectionState(collection ReferenceArchitectureScenarioDescriptorCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Packs) == 0 {
		return ReferenceArchitectureValCScenarioDescriptorStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedCategories, referenceArchitectureValCFailureModes()...) ||
		!containsExactTrimmedStringSet(collection.SupportedStates, referenceArchitectureValCScenarioStates()...) ||
		!containsExactTrimmedStringSet(collection.SupportedSeverities, referenceArchitectureValCScenarioSeverities()...) ||
		!referenceArchitectureValCHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Packs) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCScenarioDescriptorStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, pack := range collection.Packs {
		if _, ok := seenFamilies[strings.TrimSpace(pack.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCScenarioDescriptorStatePartial
		}
		seenFamilies[strings.TrimSpace(pack.BlueprintFamily)] = struct{}{}
		if EvaluateReferenceArchitectureValCScenarioDescriptorPackState(pack) != ReferenceArchitectureValCScenarioDescriptorStateActive {
			return ReferenceArchitectureValCScenarioDescriptorStatePartial
		}
	}
	return ReferenceArchitectureValCScenarioDescriptorStateActive
}

func EvaluateReferenceArchitectureValCDegradedModePackState(pack ReferenceArchitectureDegradedModePack) string {
	if !referenceArchitectureValBRequiredRefsPresent(pack.PackID, pack.BlueprintFamily, pack.ProjectionDisclaimer) || len(pack.Modes) == 0 {
		return ReferenceArchitectureValCDegradedModeStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(pack.ProjectionDisclaimer) {
		return ReferenceArchitectureValCDegradedModeStatePartial
	}
	seenScenarioRefs := map[string]struct{}{}
	for _, mode := range pack.Modes {
		if !referenceArchitectureValBRequiredRefsPresent(mode.DegradedModeID, mode.ScenarioRef, mode.RequiredOperatorAction, mode.FreshnessExpectation, mode.SupportBoundaryRef) ||
			len(mode.AllowedOperations) == 0 ||
			len(mode.BlockedOperations) == 0 ||
			len(mode.EvidenceRequired) == 0 {
			return ReferenceArchitectureValCDegradedModeStateIncomplete
		}
		if _, ok := seenScenarioRefs[strings.TrimSpace(mode.ScenarioRef)]; ok {
			return ReferenceArchitectureValCDegradedModeStatePartial
		}
		seenScenarioRefs[strings.TrimSpace(mode.ScenarioRef)] = struct{}{}
		if mode.UnsupportedBehavior || !mode.RedactionKeepsCaveats || !containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), mode.EvidenceRequired...) {
			return ReferenceArchitectureValCDegradedModeStatePartial
		}
	}
	if len(seenScenarioRefs) != len(referenceArchitectureValCRequiredFailureModes()) {
		return ReferenceArchitectureValCDegradedModeStatePartial
	}
	return ReferenceArchitectureValCDegradedModeStateActive
}

func EvaluateReferenceArchitectureValCDegradedModeCollectionState(collection ReferenceArchitectureDegradedModeCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Packs) == 0 {
		return ReferenceArchitectureValCDegradedModeStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.Packs) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCDegradedModeStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, pack := range collection.Packs {
		if _, ok := seenFamilies[strings.TrimSpace(pack.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCDegradedModeStatePartial
		}
		seenFamilies[strings.TrimSpace(pack.BlueprintFamily)] = struct{}{}
		if EvaluateReferenceArchitectureValCDegradedModePackState(pack) != ReferenceArchitectureValCDegradedModeStateActive {
			return ReferenceArchitectureValCDegradedModeStatePartial
		}
	}
	return ReferenceArchitectureValCDegradedModeStateActive
}

func EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack ReferenceArchitectureRecoveryExpectationPack) string {
	if !referenceArchitectureValBRequiredRefsPresent(pack.PackID, pack.BlueprintFamily, pack.ProjectionDisclaimer) || len(pack.Expectations) == 0 {
		return ReferenceArchitectureValCRecoveryExpectationStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(pack.ProjectionDisclaimer) {
		return ReferenceArchitectureValCRecoveryExpectationStatePartial
	}
	seenScenarioRefs := map[string]struct{}{}
	for _, expectation := range pack.Expectations {
		if !referenceArchitectureValBRequiredRefsPresent(expectation.RecoveryID, expectation.ScenarioRef, expectation.ExpectedRecoveryPath, expectation.OperatorActionRequired, expectation.RollbackOrRestoreBoundary, expectation.Timestamp, expectation.SupportedEnvironmentScope) ||
			len(expectation.RequiredEvidenceTypes) == 0 ||
			len(expectation.UnsupportedConditions) == 0 {
			return ReferenceArchitectureValCRecoveryExpectationStateIncomplete
		}
		if _, ok := seenScenarioRefs[strings.TrimSpace(expectation.ScenarioRef)]; ok {
			return ReferenceArchitectureValCRecoveryExpectationStatePartial
		}
		seenScenarioRefs[strings.TrimSpace(expectation.ScenarioRef)] = struct{}{}
		if !containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), expectation.RequiredEvidenceTypes...) ||
			expectation.CertifiedRecoveryClaim {
			return ReferenceArchitectureValCRecoveryExpectationStatePartial
		}
		if _, ok := referenceArchitectureVal0ParseTimestamp(expectation.Timestamp); !ok {
			return ReferenceArchitectureValCRecoveryExpectationStatePartial
		}
		if strings.TrimSpace(expectation.FreshnessState) != IntelligenceCalibrationFreshnessFresh {
			return ReferenceArchitectureValCRecoveryExpectationStatePartial
		}
	}
	if len(seenScenarioRefs) != len(referenceArchitectureValCRequiredFailureModes()) {
		return ReferenceArchitectureValCRecoveryExpectationStatePartial
	}
	return ReferenceArchitectureValCRecoveryExpectationStateActive
}

func EvaluateReferenceArchitectureValCRecoveryExpectationCollectionState(collection ReferenceArchitectureRecoveryExpectationCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Packs) == 0 {
		return ReferenceArchitectureValCRecoveryExpectationStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.Packs) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCRecoveryExpectationStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, pack := range collection.Packs {
		if _, ok := seenFamilies[strings.TrimSpace(pack.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCRecoveryExpectationStatePartial
		}
		seenFamilies[strings.TrimSpace(pack.BlueprintFamily)] = struct{}{}
		if EvaluateReferenceArchitectureValCRecoveryExpectationPackState(pack) != ReferenceArchitectureValCRecoveryExpectationStateActive {
			return ReferenceArchitectureValCRecoveryExpectationStatePartial
		}
	}
	return ReferenceArchitectureValCRecoveryExpectationStateActive
}

func EvaluateReferenceArchitectureValCScalingScenarioPackState(pack ReferenceArchitectureScalingScenarioPack) string {
	if !referenceArchitectureValBRequiredRefsPresent(pack.PackID, pack.BlueprintFamily, pack.ProjectionDisclaimer) || len(pack.Scenarios) == 0 {
		return ReferenceArchitectureValCScalingScenarioStateIncomplete
	}
	if !containsExactTrimmedStringSet(pack.SupportedCategories, referenceArchitectureValCScalingCategories()...) || !referenceArchitectureValCHasProjectionDisclaimer(pack.ProjectionDisclaimer) {
		return ReferenceArchitectureValCScalingScenarioStatePartial
	}
	seenCategories := map[string]struct{}{}
	for _, scenario := range pack.Scenarios {
		if !referenceArchitectureValBRequiredRefsPresent(scenario.ScalingScenarioID, scenario.Category, scenario.BlueprintFamily, scenario.ExpectedCapacityBoundary, scenario.Timestamp, scenario.ProjectionDisclaimer) ||
			len(scenario.RequiredEvidenceTypes) == 0 {
			return ReferenceArchitectureValCScalingScenarioStateIncomplete
		}
		if _, ok := seenCategories[strings.TrimSpace(scenario.Category)]; ok {
			return ReferenceArchitectureValCScalingScenarioStatePartial
		}
		seenCategories[strings.TrimSpace(scenario.Category)] = struct{}{}
		if strings.TrimSpace(scenario.BlueprintFamily) != strings.TrimSpace(pack.BlueprintFamily) ||
			!containsTrimmedString(referenceArchitectureValCRequiredScalingCategories(), scenario.Category) ||
			!containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), scenario.RequiredEvidenceTypes...) ||
			!referenceArchitectureValCHasProjectionDisclaimer(scenario.ProjectionDisclaimer) ||
			scenario.PerformanceGuarantee {
			return ReferenceArchitectureValCScalingScenarioStatePartial
		}
		if scenario.DegradationThreshold <= 0 || scenario.FailClosedThreshold <= 0 || scenario.DegradationThreshold >= scenario.FailClosedThreshold {
			return ReferenceArchitectureValCScalingScenarioStatePartial
		}
		if _, ok := referenceArchitectureVal0ParseTimestamp(scenario.Timestamp); !ok {
			return ReferenceArchitectureValCScalingScenarioStatePartial
		}
		if strings.TrimSpace(scenario.FreshnessState) != IntelligenceCalibrationFreshnessFresh {
			return ReferenceArchitectureValCScalingScenarioStatePartial
		}
	}
	if len(seenCategories) != len(referenceArchitectureValCRequiredScalingCategories()) {
		return ReferenceArchitectureValCScalingScenarioStatePartial
	}
	return ReferenceArchitectureValCScalingScenarioStateActive
}

func EvaluateReferenceArchitectureValCScalingScenarioCollectionState(collection ReferenceArchitectureScalingScenarioCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Packs) == 0 {
		return ReferenceArchitectureValCScalingScenarioStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedCategories, referenceArchitectureValCScalingCategories()...) ||
		!referenceArchitectureValCHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Packs) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCScalingScenarioStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, pack := range collection.Packs {
		if _, ok := seenFamilies[strings.TrimSpace(pack.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCScalingScenarioStatePartial
		}
		seenFamilies[strings.TrimSpace(pack.BlueprintFamily)] = struct{}{}
		if EvaluateReferenceArchitectureValCScalingScenarioPackState(pack) != ReferenceArchitectureValCScalingScenarioStateActive {
			return ReferenceArchitectureValCScalingScenarioStatePartial
		}
	}
	return ReferenceArchitectureValCScalingScenarioStateActive
}

func EvaluateReferenceArchitectureValCTrustPathCheckState(check ReferenceArchitectureTrustPathContinuityCheck) string {
	if !referenceArchitectureValBRequiredRefsPresent(check.CheckID, check.BlueprintFamily, check.ProjectionDisclaimer) || len(check.EvidenceRefs) == 0 {
		return ReferenceArchitectureValCTrustPathStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(check.ProjectionDisclaimer) || !containsTrimmedString(referenceArchitectureVal0Families(), check.BlueprintFamily) {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	if !check.TrustAnchorAvailabilityExpected {
		return ReferenceArchitectureValCTrustPathStateBlocked
	}
	profile, ok := LookupReferenceArchitectureValAFamilyProfile(check.BlueprintFamily)
	if !ok {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	allFresh, stale, evidenceOK := referenceArchitectureValBEvidenceValid(check.EvidenceRefs)
	if !evidenceOK || !allFresh || stale {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	if profile.StrongerTrustAnchorMode && !check.StrictCustodyBoundaryExpected {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	if strings.TrimSpace(check.BlueprintFamily) == ReferenceArchitectureFamilySovereignAirGapped && check.RequiresLiveExternalTrustDependency {
		return ReferenceArchitectureValCTrustPathStateBlocked
	}
	if profile.LocalTrustAnchorAssumptionRequired && !check.AirGappedLocalTrustExpected {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	if !check.VerifierReplayVisibilityExpected {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	return ReferenceArchitectureValCTrustPathStateActive
}

func EvaluateReferenceArchitectureValCTrustPathCollectionState(collection ReferenceArchitectureTrustPathCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Checks) == 0 {
		return ReferenceArchitectureValCTrustPathStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.Checks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	checksByFamily := map[string]ReferenceArchitectureTrustPathContinuityCheck{}
	for _, check := range collection.Checks {
		if _, ok := checksByFamily[strings.TrimSpace(check.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCTrustPathStatePartial
		}
		checksByFamily[strings.TrimSpace(check.BlueprintFamily)] = check
		if EvaluateReferenceArchitectureValCTrustPathCheckState(check) != ReferenceArchitectureValCTrustPathStateActive {
			return ReferenceArchitectureValCTrustPathStatePartial
		}
	}
	enterprise := checksByFamily[ReferenceArchitectureFamilyEnterpriseDefault]
	highAssurance := checksByFamily[ReferenceArchitectureFamilyHighAssurance]
	if !highAssurance.StrictCustodyBoundaryExpected || enterprise.StrictCustodyBoundaryExpected {
		return ReferenceArchitectureValCTrustPathStatePartial
	}
	return ReferenceArchitectureValCTrustPathStateActive
}

func EvaluateReferenceArchitectureValCAuditPathCheckState(check ReferenceArchitectureAuditPathDegradationCheck) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		check.CheckID,
		check.BlueprintFamily,
		check.AuditLatencyDegradedBehavior,
		check.EvidenceCustodyPath,
		check.PartialFailureHandling,
		check.RecoveryExpectationRef,
		check.ProjectionDisclaimer,
	) || len(check.EvidenceRefs) == 0 {
		return ReferenceArchitectureValCAuditPathStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(check.ProjectionDisclaimer) || !containsTrimmedString(referenceArchitectureVal0Families(), check.BlueprintFamily) {
		return ReferenceArchitectureValCAuditPathStatePartial
	}
	if !check.AuditWriterAvailabilityExpected || !check.OperatorVisibilityRequired || !check.CanonicalFailureStatePreserved {
		return ReferenceArchitectureValCAuditPathStatePartial
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(check.EvidenceRefs)
	if !ok || !allFresh || stale {
		return ReferenceArchitectureValCAuditPathStatePartial
	}
	return ReferenceArchitectureValCAuditPathStateActive
}

func EvaluateReferenceArchitectureValCAuditPathCollectionState(collection ReferenceArchitectureAuditPathCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Checks) == 0 {
		return ReferenceArchitectureValCAuditPathStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.Checks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCAuditPathStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, check := range collection.Checks {
		if _, ok := seenFamilies[strings.TrimSpace(check.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCAuditPathStatePartial
		}
		seenFamilies[strings.TrimSpace(check.BlueprintFamily)] = struct{}{}
		if EvaluateReferenceArchitectureValCAuditPathCheckState(check) != ReferenceArchitectureValCAuditPathStateActive {
			return ReferenceArchitectureValCAuditPathStatePartial
		}
	}
	return ReferenceArchitectureValCAuditPathStateActive
}

func EvaluateReferenceArchitectureValCControlPlaneCheckState(check ReferenceArchitectureControlPlaneSafetyCheck) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		check.CheckID,
		check.BlueprintFamily,
		check.OverloadBehavior,
		check.BackpressureSemantics,
		check.RateLimitBehavior,
		check.DependencyTimeoutBehavior,
		check.FailClosedBehavior,
		check.OperatorActionRequired,
		check.ProjectionDisclaimer,
	) || len(check.EvidenceOutputRequirements) == 0 {
		return ReferenceArchitectureValCControlPlaneStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(check.ProjectionDisclaimer) || !containsTrimmedString(referenceArchitectureVal0Families(), check.BlueprintFamily) {
		return ReferenceArchitectureValCControlPlaneStatePartial
	}
	if !containsAllTrimmedStrings(referenceArchitectureVal0SupportedEvidenceTypes(), check.EvidenceOutputRequirements...) {
		return ReferenceArchitectureValCControlPlaneStatePartial
	}
	if !containsTrimmedString([]string{ReferenceArchitectureValCBoundedDegraded, ReferenceArchitectureValCFailClosed, ReferenceArchitectureValCScenarioUnsupported}, check.DependencyTimeoutBehavior) {
		return ReferenceArchitectureValCControlPlaneStatePartial
	}
	if check.AutomaticApproval || check.AutomaticMutation {
		return ReferenceArchitectureValCControlPlaneStateBlocked
	}
	return ReferenceArchitectureValCControlPlaneStateActive
}

func EvaluateReferenceArchitectureValCControlPlaneCollectionState(collection ReferenceArchitectureControlPlaneSafetyCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Checks) == 0 {
		return ReferenceArchitectureValCControlPlaneStateIncomplete
	}
	if !referenceArchitectureValCHasProjectionDisclaimer(collection.ProjectionDisclaimer) || len(collection.Checks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValCControlPlaneStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, check := range collection.Checks {
		if _, ok := seenFamilies[strings.TrimSpace(check.BlueprintFamily)]; ok {
			return ReferenceArchitectureValCControlPlaneStatePartial
		}
		seenFamilies[strings.TrimSpace(check.BlueprintFamily)] = struct{}{}
		if EvaluateReferenceArchitectureValCControlPlaneCheckState(check) != ReferenceArchitectureValCControlPlaneStateActive {
			return ReferenceArchitectureValCControlPlaneStatePartial
		}
	}
	return ReferenceArchitectureValCControlPlaneStateActive
}

func referenceArchitectureValCRequiresPriorStates(point5State, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, point6State string) bool {
	return strings.TrimSpace(point5State) == IntelligenceCalibrationPoint5StatePass &&
		strings.TrimSpace(val0CurrentState) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(val0State) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(valACurrentState) == ReferenceArchitectureValAStateActive &&
		strings.TrimSpace(valAState) == ReferenceArchitectureValAStateActive &&
		strings.TrimSpace(valBCurrentState) == ReferenceArchitectureValBStateActive &&
		strings.TrimSpace(valBState) == ReferenceArchitectureValBStateActive &&
		strings.TrimSpace(point6State) == ReferenceArchitecturePoint6StateNotComplete
}

func EvaluateReferenceArchitectureValCState(
	point5State, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, point6State,
	scenarioPackState, failureTaxonomyState, scenarioDescriptorState, degradedModeState, recoveryExpectationState,
	scalingScenarioState, trustPathState, auditPathState, controlPlaneState string,
) string {
	if !referenceArchitectureValCRequiresPriorStates(point5State, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, point6State) {
		return ReferenceArchitectureValCStateBlocked
	}
	componentStates := []string{
		scenarioPackState,
		failureTaxonomyState,
		scenarioDescriptorState,
		degradedModeState,
		recoveryExpectationState,
		scalingScenarioState,
		trustPathState,
		auditPathState,
		controlPlaneState,
	}
	for _, state := range componentStates {
		if strings.TrimSpace(state) == "" {
			return ReferenceArchitectureValCStateIncomplete
		}
	}
	if scenarioPackState == ReferenceArchitectureValCScenarioPackStateActive &&
		failureTaxonomyState == ReferenceArchitectureValCFailureTaxonomyStateActive &&
		scenarioDescriptorState == ReferenceArchitectureValCScenarioDescriptorStateActive &&
		degradedModeState == ReferenceArchitectureValCDegradedModeStateActive &&
		recoveryExpectationState == ReferenceArchitectureValCRecoveryExpectationStateActive &&
		scalingScenarioState == ReferenceArchitectureValCScalingScenarioStateActive &&
		trustPathState == ReferenceArchitectureValCTrustPathStateActive &&
		auditPathState == ReferenceArchitectureValCAuditPathStateActive &&
		controlPlaneState == ReferenceArchitectureValCControlPlaneStateActive {
		return ReferenceArchitectureValCStateActive
	}
	if scenarioPackState == ReferenceArchitectureValCScenarioPackStateIncomplete ||
		failureTaxonomyState == ReferenceArchitectureValCFailureTaxonomyStateIncomplete ||
		scenarioDescriptorState == ReferenceArchitectureValCScenarioDescriptorStateIncomplete ||
		degradedModeState == ReferenceArchitectureValCDegradedModeStateIncomplete ||
		recoveryExpectationState == ReferenceArchitectureValCRecoveryExpectationStateIncomplete ||
		scalingScenarioState == ReferenceArchitectureValCScalingScenarioStateIncomplete ||
		trustPathState == ReferenceArchitectureValCTrustPathStateIncomplete ||
		auditPathState == ReferenceArchitectureValCAuditPathStateIncomplete ||
		controlPlaneState == ReferenceArchitectureValCControlPlaneStateIncomplete {
		return ReferenceArchitectureValCStateIncomplete
	}
	if scenarioPackState == ReferenceArchitectureValCScenarioPackStateBlocked ||
		failureTaxonomyState == ReferenceArchitectureValCFailureTaxonomyStateBlocked ||
		scenarioDescriptorState == ReferenceArchitectureValCScenarioDescriptorStateBlocked ||
		degradedModeState == ReferenceArchitectureValCDegradedModeStateBlocked ||
		recoveryExpectationState == ReferenceArchitectureValCRecoveryExpectationStateBlocked ||
		scalingScenarioState == ReferenceArchitectureValCScalingScenarioStateBlocked ||
		trustPathState == ReferenceArchitectureValCTrustPathStateBlocked ||
		auditPathState == ReferenceArchitectureValCAuditPathStateBlocked ||
		controlPlaneState == ReferenceArchitectureValCControlPlaneStateBlocked {
		return ReferenceArchitectureValCStateBlocked
	}
	return ReferenceArchitectureValCStatePartial
}

func referenceArchitectureValCProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/scenario-packs",
		"/v1/reference-architecture/valc/failure-taxonomy",
		"/v1/reference-architecture/valc/scenario-descriptors",
		"/v1/reference-architecture/valc/degraded-modes",
		"/v1/reference-architecture/valc/recovery-expectations",
		"/v1/reference-architecture/valc/scaling-scenarios",
		"/v1/reference-architecture/valc/trust-path",
		"/v1/reference-architecture/valc/audit-path",
		"/v1/reference-architecture/valc/control-plane-safety",
		"/v1/reference-architecture/valc/proofs",
	}
}

func EvaluateReferenceArchitectureValCProofsState(valCState, point6State string, supportedFamilies, surfaceRefs, evidenceRefs, limitations []string, projectionDisclaimer string) string {
	baseState := strings.TrimSpace(valCState)
	if !containsExactTrimmedStringSet(supportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(surfaceRefs, referenceArchitectureValCProofSurfaceRefs()...) ||
		len(evidenceRefs) < 12 ||
		len(limitations) == 0 ||
		!referenceArchitectureValCHasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == ReferenceArchitectureValCStateActive {
			return ReferenceArchitectureValCStatePartial
		}
		return baseState
	}
	if baseState == ReferenceArchitectureValCStateActive && strings.TrimSpace(point6State) != ReferenceArchitecturePoint6StateNotComplete {
		return ReferenceArchitectureValCStatePartial
	}
	return baseState
}
