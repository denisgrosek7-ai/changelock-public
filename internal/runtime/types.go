package runtime

type DriftClass string

const (
	DriftClassNoDrift         DriftClass = "no_drift"
	DriftClassImageDigest     DriftClass = "image_digest_drift"
	DriftClassSecurityContext DriftClass = "security_context_drift"
	DriftClassServiceAccount  DriftClass = "service_account_drift"
	DriftClassWorkloadSpec    DriftClass = "workload_spec_drift"
	DriftClassUnknown         DriftClass = "unknown_runtime_drift"
	DriftClassMultiple        DriftClass = "multiple_drift"
)

type DriftSeverity string

const (
	DriftSeverityLow      DriftSeverity = "low"
	DriftSeverityMedium   DriftSeverity = "medium"
	DriftSeverityHigh     DriftSeverity = "high"
	DriftSeverityCritical DriftSeverity = "critical"
)

type RemediationMode string

const (
	RemediationModeDisabled             RemediationMode = "disabled"
	RemediationModeAlertOnly            RemediationMode = "alert-only"
	RemediationModeQuarantine           RemediationMode = "quarantine"
	RemediationModePatchApprovedState   RemediationMode = "patch-approved-state"
	RemediationModeRestartApprovedState RemediationMode = "restart-to-approved-state"
)

type VerificationState string

const (
	VerificationStateDisabled   VerificationState = "disabled"
	VerificationStateUnverified VerificationState = "unverified"
	VerificationStateVerified   VerificationState = "verified"
	VerificationStateFailed     VerificationState = "failed"
)

type ApprovedWorkloadState struct {
	TenantID                      string                   `json:"tenant_id,omitempty" yaml:"tenantId,omitempty"`
	ClusterID                     string                   `json:"cluster_id,omitempty" yaml:"clusterId,omitempty"`
	Namespace                     string                   `json:"namespace" yaml:"namespace"`
	WorkloadKind                  string                   `json:"workload_kind,omitempty" yaml:"workloadKind,omitempty"`
	Workload                      string                   `json:"workload" yaml:"workload"`
	ServiceAccountName            string                   `json:"service_account_name,omitempty" yaml:"serviceAccountName,omitempty"`
	ExpectedConfigHash            string                   `json:"expected_config_hash,omitempty" yaml:"expectedConfigHash,omitempty"`
	ApprovalCorrelationID         string                   `json:"approval_correlation_id,omitempty" yaml:"approvalCorrelationId,omitempty"`
	SourceRef                     string                   `json:"source_ref,omitempty" yaml:"sourceRef,omitempty"`
	DesiredStateVerificationState VerificationState        `json:"desired_state_verification_state,omitempty" yaml:"desiredStateVerificationState,omitempty"`
	Containers                    []ApprovedContainerState `json:"containers" yaml:"containers"`
	Labels                        map[string]string        `json:"labels,omitempty" yaml:"labels,omitempty"`
}

type ApprovedContainerState struct {
	Name           string              `json:"name" yaml:"name"`
	Image          string              `json:"image,omitempty" yaml:"image,omitempty"`
	ApprovedDigest string              `json:"approved_digest" yaml:"approvedDigest"`
	Runtime        SecurityConstraints `json:"runtime" yaml:"runtime"`
}

type ObservedWorkloadState struct {
	ClusterID          string                   `json:"cluster_id,omitempty" yaml:"clusterId,omitempty"`
	Namespace          string                   `json:"namespace" yaml:"namespace"`
	WorkloadKind       string                   `json:"workload_kind,omitempty" yaml:"workloadKind,omitempty"`
	Workload           string                   `json:"workload" yaml:"workload"`
	ServiceAccountName string                   `json:"service_account_name,omitempty" yaml:"serviceAccountName,omitempty"`
	ActualConfigHash   string                   `json:"actual_config_hash,omitempty" yaml:"actualConfigHash,omitempty"`
	Containers         []ObservedContainerState `json:"containers" yaml:"containers"`
	PodLabels          map[string]string        `json:"pod_labels,omitempty" yaml:"podLabels,omitempty"`
}

type ObservedContainerState struct {
	Name          string          `json:"name" yaml:"name"`
	Image         string          `json:"image,omitempty" yaml:"image,omitempty"`
	RunningDigest string          `json:"running_digest" yaml:"runningDigest"`
	Runtime       SecurityPosture `json:"runtime" yaml:"runtime"`
}

type SecurityConstraints struct {
	RunAsNonRoot             bool `json:"run_as_non_root" yaml:"runAsNonRoot"`
	ReadOnlyRootFilesystem   bool `json:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem"`
	AllowPrivilegeEscalation bool `json:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation"`
	DropAllCapabilities      bool `json:"drop_all_capabilities" yaml:"dropAllCapabilities"`
	SeccompRuntimeDefault    bool `json:"seccomp_runtime_default" yaml:"seccompRuntimeDefault"`
	DenyPrivileged           bool `json:"deny_privileged" yaml:"denyPrivileged"`
}

type SecurityPosture struct {
	RunAsNonRoot             bool `json:"run_as_non_root" yaml:"runAsNonRoot"`
	ReadOnlyRootFilesystem   bool `json:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem"`
	AllowPrivilegeEscalation bool `json:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation"`
	DropAllCapabilities      bool `json:"drop_all_capabilities" yaml:"dropAllCapabilities"`
	SeccompRuntimeDefault    bool `json:"seccomp_runtime_default" yaml:"seccompRuntimeDefault"`
	Privileged               bool `json:"privileged" yaml:"privileged"`
}

type ComparisonResult struct {
	ScanID                        string            `json:"scan_id,omitempty"`
	ClusterID                     string            `json:"cluster_id,omitempty"`
	Namespace                     string            `json:"namespace"`
	WorkloadKind                  string            `json:"workload_kind,omitempty"`
	Workload                      string            `json:"workload"`
	Image                         string            `json:"image,omitempty"`
	ApprovedDigest                string            `json:"approved_digest,omitempty"`
	RunningDigest                 string            `json:"running_digest,omitempty"`
	ServiceAccountExpected        string            `json:"service_account_expected,omitempty"`
	ServiceAccountObserved        string            `json:"service_account_observed,omitempty"`
	Result                        string            `json:"drift_result"`
	Classes                       []string          `json:"drift_classes,omitempty"`
	Severity                      DriftSeverity     `json:"severity,omitempty"`
	Remediable                    bool              `json:"remediable"`
	SelectedRemediationMode       RemediationMode   `json:"selected_remediation_mode,omitempty"`
	DesiredStateVerificationState VerificationState `json:"desired_state_verification_state,omitempty"`
	Reasons                       []string          `json:"reasons,omitempty"`
	Evidence                      *DriftEvidence    `json:"evidence,omitempty"`
}

func (r ComparisonResult) HasDrift() bool {
	return r.Result != string(DriftClassNoDrift)
}

type DriftEvidence struct {
	ImageMismatches           []ImageMismatch           `json:"image_mismatches,omitempty"`
	ConfigExpectation         string                    `json:"expected_config_hash,omitempty"`
	ConfigObserved            string                    `json:"actual_config_hash,omitempty"`
	ServiceAccountExpected    string                    `json:"service_account_expected,omitempty"`
	ServiceAccountObserved    string                    `json:"service_account_observed,omitempty"`
	SecurityContextMismatches []SecurityContextMismatch `json:"security_context_mismatches,omitempty"`
	MissingContainers         []string                  `json:"missing_containers,omitempty"`
	UnexpectedContainers      []string                  `json:"unexpected_containers,omitempty"`
}

type ImageMismatch struct {
	Container      string `json:"container,omitempty"`
	ApprovedImage  string `json:"approved_image,omitempty"`
	RunningImage   string `json:"running_image,omitempty"`
	ApprovedDigest string `json:"approved_digest,omitempty"`
	RunningDigest  string `json:"running_digest,omitempty"`
}

type SecurityContextMismatch struct {
	Container string `json:"container,omitempty"`
	Field     string `json:"field,omitempty"`
	Expected  bool   `json:"expected"`
	Actual    bool   `json:"actual"`
}
