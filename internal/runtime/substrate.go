package runtime

import (
	"strings"
	"time"
)

const (
	SubstrateTruthSchemaVersion = "2.runtime_substrate_truth.v1"
	ProfileMatchSchemaVersion   = "2.runtime_profile_match.v1"

	SignalPathTelemetry = "telemetry"
	SignalPathDecision  = "decision"
	SignalPathAction    = "action"

	TrustBoundaryAppLayer            = "app_layer"
	TrustBoundaryNodeLayer           = "node_layer"
	TrustBoundaryKernelRuntimeLayer  = "kernel_runtime_layer"
	TrustBoundaryAttestationProvider = "attestation_provider_layer"
	TrustBoundaryControlPlane        = "control_plane_layer"

	SubstrateClassStandard     = "standard"
	SubstrateClassHardened     = "hardened"
	SubstrateClassConfidential = "confidential"

	AttestationStateVerified = "verified"
	AttestationStateDegraded = "degraded"
	AttestationStateMissing  = "missing"
	AttestationStateMismatch = "mismatch"
	AttestationStateRevoked  = "revoked"

	SubstrateTruthStateBound    = "runtime_truth_bound"
	SubstrateTruthStateDegraded = "runtime_truth_degraded"
	SubstrateTruthStateMismatch = "runtime_truth_mismatch"

	ProfileMatchStateMatch               = "match"
	ProfileMatchStateDegradedAcceptable  = "degraded_acceptable"
	ProfileMatchStateMismatchDeny        = "mismatch_deny"
	ProfileMatchStateQuarantineCandidate = "mismatch_quarantine_candidate"

	CredentialReleaseWithheld = "withheld"
	CredentialReleaseReleased = "released"
)

type WorkloadIdentity struct {
	ClusterID     string `json:"cluster_id,omitempty"`
	Namespace     string `json:"namespace,omitempty"`
	WorkloadKind  string `json:"workload_kind,omitempty"`
	Workload      string `json:"workload,omitempty"`
	PodUID        string `json:"pod_uid,omitempty"`
	ImageDigest   string `json:"image_digest,omitempty"`
	PolicySubject string `json:"policy_subject_ref,omitempty"`
}

type ProcessIdentity struct {
	ProcessName string `json:"process_name,omitempty"`
	ProcessPath string `json:"process_path,omitempty"`
	PID         int    `json:"pid,omitempty"`
	CgroupID    string `json:"cgroup_id,omitempty"`
	NamespaceID string `json:"namespace_id,omitempty"`
	LineageRef  string `json:"lineage_ref,omitempty"`
}

type NodeIdentity struct {
	NodeID         string `json:"node_id,omitempty"`
	SubstrateClass string `json:"substrate_class,omitempty"`
	TrustBoundary  string `json:"trust_boundary,omitempty"`
	AttestationRef string `json:"attestation_ref,omitempty"`
}

type AttestationBinding struct {
	Provider               string    `json:"provider,omitempty"`
	QuoteType              string    `json:"quote_type,omitempty"`
	Measurement            string    `json:"measurement,omitempty"`
	LifecycleState         string    `json:"lifecycle_state,omitempty"`
	ObservedState          string    `json:"observed_state,omitempty"`
	CredentialReleaseState string    `json:"credential_release_state,omitempty"`
	VerifiedAt             time.Time `json:"verified_at,omitempty"`
}

type ExecutionProfile struct {
	ProfileID                string   `json:"profile_id"`
	DisplayName              string   `json:"display_name,omitempty"`
	AllowedProviders         []string `json:"allowed_providers,omitempty"`
	AllowedSubstrateClasses  []string `json:"allowed_substrate_classes,omitempty"`
	RequiredMeasurements     []string `json:"required_measurements,omitempty"`
	RequireAttestation       bool     `json:"require_attestation"`
	RequireCredentialRelease bool     `json:"require_credential_release"`
	AllowDegraded            bool     `json:"allow_degraded"`
	MinimumConfidence        int      `json:"minimum_confidence,omitempty"`
	TrustBoundary            string   `json:"trust_boundary,omitempty"`
}

type SubstrateTruthRecord struct {
	SchemaVersion         string             `json:"schema_version"`
	SubjectRef            string             `json:"subject_ref"`
	Workload              WorkloadIdentity   `json:"workload"`
	Process               ProcessIdentity    `json:"process"`
	Node                  NodeIdentity       `json:"node"`
	Attestation           AttestationBinding `json:"attestation"`
	TelemetryPath         string             `json:"telemetry_path,omitempty"`
	DecisionPath          string             `json:"decision_path,omitempty"`
	ActionPath            string             `json:"action_path,omitempty"`
	ConfidenceScore       int                `json:"confidence_score"`
	SeverityScore         int                `json:"severity_score"`
	ActionabilityScore    int                `json:"actionability_score"`
	FalsePositiveBudget   string             `json:"false_positive_budget,omitempty"`
	ResponseLatencyBudget string             `json:"response_latency_budget,omitempty"`
	CurrentState          string             `json:"current_state"`
	Reasons               []string           `json:"reasons,omitempty"`
	ObservedAt            time.Time          `json:"observed_at"`
}

type ProfileMatch struct {
	SchemaVersion            string    `json:"schema_version"`
	ProfileID                string    `json:"profile_id"`
	SubjectRef               string    `json:"subject_ref"`
	CurrentState             string    `json:"current_state"`
	Allowed                  bool      `json:"allowed"`
	CredentialReleaseAllowed bool      `json:"credential_release_allowed"`
	ConfidenceScore          int       `json:"confidence_score"`
	Reasons                  []string  `json:"reasons,omitempty"`
	EvaluatedAt              time.Time `json:"evaluated_at"`
}

func DefaultExecutionProfiles() []ExecutionProfile {
	return []ExecutionProfile{
		{
			ProfileID:                "confidential-strict",
			DisplayName:              "Confidential Strict",
			AllowedProviders:         []string{"sgx", "tdx", "sev"},
			AllowedSubstrateClasses:  []string{SubstrateClassConfidential},
			RequireAttestation:       true,
			RequireCredentialRelease: true,
			AllowDegraded:            false,
			MinimumConfidence:        80,
			TrustBoundary:            TrustBoundaryAttestationProvider,
		},
		{
			ProfileID:               "hardened-node",
			DisplayName:             "Hardened Node",
			AllowedProviders:        []string{"sgx", "tdx", "sev"},
			AllowedSubstrateClasses: []string{SubstrateClassHardened, SubstrateClassConfidential},
			RequireAttestation:      true,
			AllowDegraded:           true,
			MinimumConfidence:       60,
			TrustBoundary:           TrustBoundaryNodeLayer,
		},
		{
			ProfileID:                "crypto-hardening",
			DisplayName:              "Crypto Hardening",
			AllowedProviders:         []string{"sgx", "tdx", "sev"},
			AllowedSubstrateClasses:  []string{SubstrateClassHardened, SubstrateClassConfidential},
			RequireAttestation:       true,
			RequireCredentialRelease: true,
			AllowDegraded:            false,
			MinimumConfidence:        70,
			TrustBoundary:            TrustBoundaryControlPlane,
		},
	}
}

func ExecutionProfileByID(id string) (ExecutionProfile, bool) {
	id = strings.TrimSpace(id)
	for _, profile := range DefaultExecutionProfiles() {
		if profile.ProfileID == id {
			return profile, true
		}
	}
	return ExecutionProfile{}, false
}

func NormalizeSubstrateTruthRecord(record SubstrateTruthRecord, now func() time.Time) SubstrateTruthRecord {
	if now == nil {
		now = time.Now
	}
	if strings.TrimSpace(record.SchemaVersion) == "" {
		record.SchemaVersion = SubstrateTruthSchemaVersion
	}
	record.SubjectRef = strings.TrimSpace(record.SubjectRef)
	record.Workload.ClusterID = strings.TrimSpace(record.Workload.ClusterID)
	record.Workload.Namespace = strings.TrimSpace(record.Workload.Namespace)
	record.Workload.WorkloadKind = strings.TrimSpace(record.Workload.WorkloadKind)
	record.Workload.Workload = strings.TrimSpace(record.Workload.Workload)
	record.Workload.PodUID = strings.TrimSpace(record.Workload.PodUID)
	record.Workload.ImageDigest = strings.TrimSpace(record.Workload.ImageDigest)
	record.Workload.PolicySubject = strings.TrimSpace(record.Workload.PolicySubject)
	record.Process.ProcessName = strings.TrimSpace(record.Process.ProcessName)
	record.Process.ProcessPath = strings.TrimSpace(record.Process.ProcessPath)
	record.Process.CgroupID = strings.TrimSpace(record.Process.CgroupID)
	record.Process.NamespaceID = strings.TrimSpace(record.Process.NamespaceID)
	record.Process.LineageRef = strings.TrimSpace(record.Process.LineageRef)
	record.Node.NodeID = strings.TrimSpace(record.Node.NodeID)
	record.Node.SubstrateClass = normalizeSubstrateClass(record.Node.SubstrateClass)
	record.Node.TrustBoundary = normalizeTrustBoundary(record.Node.TrustBoundary)
	record.Node.AttestationRef = strings.TrimSpace(record.Node.AttestationRef)
	record.Attestation.Provider = normalizeProvider(record.Attestation.Provider)
	record.Attestation.QuoteType = strings.TrimSpace(record.Attestation.QuoteType)
	record.Attestation.Measurement = strings.TrimSpace(record.Attestation.Measurement)
	record.Attestation.LifecycleState = strings.TrimSpace(record.Attestation.LifecycleState)
	record.Attestation.ObservedState = normalizeAttestationState(record.Attestation.ObservedState)
	record.Attestation.CredentialReleaseState = normalizeCredentialReleaseState(record.Attestation.CredentialReleaseState)
	if record.ObservedAt.IsZero() {
		record.ObservedAt = now().UTC()
	}
	if strings.TrimSpace(record.TelemetryPath) == "" {
		record.TelemetryPath = SignalPathTelemetry
	}
	if strings.TrimSpace(record.DecisionPath) == "" {
		record.DecisionPath = SignalPathDecision
	}
	if strings.TrimSpace(record.ActionPath) == "" {
		record.ActionPath = SignalPathAction
	}
	record.ConfidenceScore, record.SeverityScore, record.ActionabilityScore, record.CurrentState, record.Reasons = evaluateSubstrateTruth(record)
	if strings.TrimSpace(record.FalsePositiveBudget) == "" {
		record.FalsePositiveBudget = "bounded_fp_budget"
	}
	if strings.TrimSpace(record.ResponseLatencyBudget) == "" {
		record.ResponseLatencyBudget = "bounded_response_latency"
	}
	return record
}

func MatchExecutionProfile(profile ExecutionProfile, truth SubstrateTruthRecord) ProfileMatch {
	truth = NormalizeSubstrateTruthRecord(truth, nil)
	profile.ProfileID = strings.TrimSpace(profile.ProfileID)
	profile.DisplayName = strings.TrimSpace(profile.DisplayName)
	profile.TrustBoundary = normalizeTrustBoundary(profile.TrustBoundary)
	if profile.MinimumConfidence <= 0 {
		profile.MinimumConfidence = 60
	}
	reasons := []string{}
	mismatchDeny := false
	quarantineCandidate := false

	if profile.RequireAttestation && truth.Attestation.ObservedState == AttestationStateMissing {
		reasons = append(reasons, "required_attestation_missing")
		mismatchDeny = true
	}
	if truth.Attestation.ObservedState == AttestationStateRevoked || truth.Attestation.ObservedState == AttestationStateMismatch {
		reasons = append(reasons, "attestation_state_not_trusted")
		mismatchDeny = true
	}
	if len(profile.AllowedProviders) > 0 && !containsString(profile.AllowedProviders, truth.Attestation.Provider) {
		reasons = append(reasons, "provider_not_allowed")
		mismatchDeny = true
	}
	if len(profile.AllowedSubstrateClasses) > 0 && !containsString(profile.AllowedSubstrateClasses, truth.Node.SubstrateClass) {
		reasons = append(reasons, "substrate_class_not_allowed")
		mismatchDeny = true
	}
	if len(profile.RequiredMeasurements) > 0 && !containsString(profile.RequiredMeasurements, truth.Attestation.Measurement) {
		reasons = append(reasons, "measurement_not_trusted")
		mismatchDeny = true
	}
	if profile.TrustBoundary != "" && truth.Node.TrustBoundary != "" && profile.TrustBoundary != truth.Node.TrustBoundary {
		reasons = append(reasons, "trust_boundary_mismatch")
		quarantineCandidate = true
	}
	if truth.ConfidenceScore < profile.MinimumConfidence {
		reasons = append(reasons, "confidence_below_profile_threshold")
		quarantineCandidate = true
	}
	credentialReleaseAllowed := truth.Attestation.CredentialReleaseState == CredentialReleaseReleased && truth.Attestation.ObservedState == AttestationStateVerified
	if profile.RequireCredentialRelease && !credentialReleaseAllowed {
		reasons = append(reasons, "credential_release_not_allowed")
		if truth.Attestation.ObservedState == AttestationStateVerified {
			quarantineCandidate = true
		} else {
			mismatchDeny = true
		}
	}
	if truth.CurrentState == SubstrateTruthStateDegraded && !profile.AllowDegraded {
		reasons = append(reasons, "degraded_truth_not_allowed")
		quarantineCandidate = true
	}

	currentState := ProfileMatchStateMatch
	allowed := true
	switch {
	case mismatchDeny:
		currentState = ProfileMatchStateMismatchDeny
		allowed = false
	case quarantineCandidate:
		if profile.AllowDegraded && truth.CurrentState == SubstrateTruthStateDegraded {
			currentState = ProfileMatchStateDegradedAcceptable
			allowed = true
		} else {
			currentState = ProfileMatchStateQuarantineCandidate
			allowed = false
		}
	case truth.CurrentState == SubstrateTruthStateDegraded && profile.AllowDegraded:
		currentState = ProfileMatchStateDegradedAcceptable
	default:
		currentState = ProfileMatchStateMatch
	}

	return ProfileMatch{
		SchemaVersion:            ProfileMatchSchemaVersion,
		ProfileID:                profile.ProfileID,
		SubjectRef:               truth.SubjectRef,
		CurrentState:             currentState,
		Allowed:                  allowed,
		CredentialReleaseAllowed: credentialReleaseAllowed,
		ConfidenceScore:          truth.ConfidenceScore,
		Reasons:                  uniqueStrings(reasons),
		EvaluatedAt:              time.Now().UTC(),
	}
}

func evaluateSubstrateTruth(record SubstrateTruthRecord) (confidence, severity, actionability int, currentState string, reasons []string) {
	confidence = 70
	severity = 10
	actionability = 40

	if record.Workload.ImageDigest != "" {
		confidence += 10
	}
	if record.Node.NodeID == "" {
		confidence -= 10
		reasons = append(reasons, "node_identity_missing")
	}
	if record.Node.SubstrateClass == SubstrateClassConfidential {
		confidence += 10
	}
	switch record.Attestation.ObservedState {
	case AttestationStateVerified:
		confidence += 15
		actionability += 20
	case AttestationStateDegraded:
		confidence -= 10
		severity += 20
		actionability += 10
		reasons = append(reasons, "attestation_degraded")
	case AttestationStateMissing:
		confidence -= 20
		severity += 30
		actionability += 10
		reasons = append(reasons, "attestation_missing")
	case AttestationStateMismatch:
		confidence -= 25
		severity += 40
		actionability += 25
		reasons = append(reasons, "attestation_mismatch")
	case AttestationStateRevoked:
		confidence -= 35
		severity += 50
		actionability += 35
		reasons = append(reasons, "attestation_revoked")
	}
	if record.Process.CgroupID == "" && record.Process.ProcessName == "" {
		confidence -= 5
		reasons = append(reasons, "process_binding_weak")
	}
	if record.Node.TrustBoundary == TrustBoundaryAttestationProvider || record.Node.TrustBoundary == TrustBoundaryKernelRuntimeLayer {
		confidence += 5
	}

	confidence = clamp(confidence, 0, 100)
	severity = clamp(severity, 0, 100)
	actionability = clamp(actionability, 0, 100)

	switch {
	case record.Attestation.ObservedState == AttestationStateMismatch || record.Attestation.ObservedState == AttestationStateRevoked || confidence < 40:
		currentState = SubstrateTruthStateMismatch
	case record.Attestation.ObservedState == AttestationStateDegraded || record.Attestation.ObservedState == AttestationStateMissing || confidence < 75:
		currentState = SubstrateTruthStateDegraded
	default:
		currentState = SubstrateTruthStateBound
	}
	if len(reasons) == 0 {
		reasons = []string{"workload_process_node_identity_bound"}
	}
	return confidence, severity, actionability, currentState, uniqueStrings(reasons)
}

func normalizeProvider(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func normalizeSubstrateClass(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case SubstrateClassHardened, SubstrateClassConfidential:
		return value
	default:
		return SubstrateClassStandard
	}
}

func normalizeTrustBoundary(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case TrustBoundaryAppLayer, TrustBoundaryNodeLayer, TrustBoundaryKernelRuntimeLayer, TrustBoundaryAttestationProvider, TrustBoundaryControlPlane:
		return value
	default:
		return TrustBoundaryAppLayer
	}
}

func normalizeAttestationState(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case AttestationStateVerified, AttestationStateDegraded, AttestationStateMismatch, AttestationStateRevoked:
		return value
	default:
		return AttestationStateMissing
	}
}

func normalizeCredentialReleaseState(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case CredentialReleaseReleased:
		return CredentialReleaseReleased
	default:
		return CredentialReleaseWithheld
	}
}

func clamp(value, minValue, maxValue int) int {
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}

func containsString(values []string, needle string) bool {
	needle = strings.TrimSpace(needle)
	for _, value := range values {
		if strings.TrimSpace(value) == needle {
			return true
		}
	}
	return false
}

func uniqueStrings(values []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	return out
}
