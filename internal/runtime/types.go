package runtime

type DriftClass string

const (
	DriftClassNoDrift         DriftClass = "no_drift"
	DriftClassImage           DriftClass = "image_drift"
	DriftClassConfig          DriftClass = "config_drift"
	DriftClassSecurityContext DriftClass = "security_context_drift"
	DriftClassMultiple        DriftClass = "multiple_drift"
)

type ApprovedWorkloadState struct {
	Namespace          string                   `json:"namespace" yaml:"namespace"`
	Workload           string                   `json:"workload" yaml:"workload"`
	ExpectedConfigHash string                   `json:"expected_config_hash,omitempty" yaml:"expectedConfigHash,omitempty"`
	Containers         []ApprovedContainerState `json:"containers" yaml:"containers"`
	Labels             map[string]string        `json:"labels,omitempty" yaml:"labels,omitempty"`
}

type ApprovedContainerState struct {
	Name           string              `json:"name" yaml:"name"`
	Image          string              `json:"image,omitempty" yaml:"image,omitempty"`
	ApprovedDigest string              `json:"approved_digest" yaml:"approvedDigest"`
	Runtime        SecurityConstraints `json:"runtime" yaml:"runtime"`
}

type ObservedWorkloadState struct {
	Namespace        string                   `json:"namespace" yaml:"namespace"`
	Workload         string                   `json:"workload" yaml:"workload"`
	ActualConfigHash string                   `json:"actual_config_hash,omitempty" yaml:"actualConfigHash,omitempty"`
	Containers       []ObservedContainerState `json:"containers" yaml:"containers"`
	PodLabels        map[string]string        `json:"pod_labels,omitempty" yaml:"podLabels,omitempty"`
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
	ScanID         string         `json:"scan_id,omitempty"`
	Namespace      string         `json:"namespace"`
	Workload       string         `json:"workload"`
	Image          string         `json:"image,omitempty"`
	ApprovedDigest string         `json:"approved_digest,omitempty"`
	RunningDigest  string         `json:"running_digest,omitempty"`
	Result         string         `json:"drift_result"`
	Classes        []string       `json:"drift_classes,omitempty"`
	Reasons        []string       `json:"reasons,omitempty"`
	Evidence       *DriftEvidence `json:"evidence,omitempty"`
}

func (r ComparisonResult) HasDrift() bool {
	return r.Result != string(DriftClassNoDrift)
}

type DriftEvidence struct {
	ImageMismatches           []ImageMismatch           `json:"image_mismatches,omitempty"`
	ConfigExpectation         string                    `json:"expected_config_hash,omitempty"`
	ConfigObserved            string                    `json:"actual_config_hash,omitempty"`
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
