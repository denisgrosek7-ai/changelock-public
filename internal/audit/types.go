package audit

import (
	"strings"
	"time"

	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
	"github.com/denisgrosek/changelock/internal/verify"
)

const (
	EventTypeArtifactVerificationResult = "artifact_verification_result"
	EventTypeDeployGateDecision         = "deploy_gate_decision"
	EventTypePolicyDecision             = "policy_decision"
	EventTypeRuntimeDriftResult         = "runtime_drift_result"

	DecisionAllow = "ALLOW"
	DecisionDeny  = "DENY"
	DecisionError = "ERROR"
)

type Event struct {
	RequestID                string           `json:"request_id"`
	Timestamp                time.Time        `json:"timestamp"`
	Component                string           `json:"component"`
	EventType                string           `json:"event_type"`
	Actor                    string           `json:"actor,omitempty"`
	ClusterID                string           `json:"cluster_id,omitempty"`
	TenantID                 string           `json:"tenant_id,omitempty"`
	Repo                     string           `json:"repo,omitempty"`
	Branch                   string           `json:"branch,omitempty"`
	Environment              string           `json:"environment,omitempty"`
	Namespace                string           `json:"namespace,omitempty"`
	Workload                 string           `json:"workload,omitempty"`
	Image                    string           `json:"image,omitempty"`
	Digest                   string           `json:"digest,omitempty"`
	CVEID                    string           `json:"cve_id,omitempty"`
	Decision                 string           `json:"decision"`
	Reasons                  []string         `json:"reasons,omitempty"`
	DriftResult              string           `json:"drift_result,omitempty"`
	DriftClasses             []string         `json:"drift_classes,omitempty"`
	VerifierSummary          *VerifierSummary `json:"verifier_summary,omitempty"`
	PolicyVersion            string           `json:"policy_version,omitempty"`
	PolicyBundleID           string           `json:"policy_bundle_id,omitempty"`
	PolicyBundleHash         string           `json:"policy_bundle_hash,omitempty"`
	DecisionHash             string           `json:"decision_hash,omitempty"`
	IsException              bool             `json:"is_exception,omitempty"`
	ExceptionID              string           `json:"exception_id,omitempty"`
	ExceptionType            string           `json:"exception_type,omitempty"`
	ExceptionStatus          string           `json:"exception_status,omitempty"`
	ExceptionReason          string           `json:"exception_reason,omitempty"`
	ExceptionTicketID        string           `json:"exception_ticket_id,omitempty"`
	ExceptionRequestedBy     string           `json:"exception_requested_by,omitempty"`
	ExceptionRequestedAt     *time.Time       `json:"exception_requested_at,omitempty"`
	ExceptionApprovedBy      string           `json:"exception_approved_by,omitempty"`
	ExceptionApprovedAt      *time.Time       `json:"exception_approved_at,omitempty"`
	ExceptionRejectedBy      string           `json:"exception_rejected_by,omitempty"`
	ExceptionRejectedAt      *time.Time       `json:"exception_rejected_at,omitempty"`
	ExceptionRejectionReason string           `json:"exception_rejection_reason,omitempty"`
	ExceptionExpiresAt       *time.Time       `json:"exception_expires_at,omitempty"`
	Evidence                 *Evidence        `json:"evidence,omitempty"`
}

type VerifierSummary struct {
	SignatureValid   bool `json:"signature_valid"`
	AttestationValid bool `json:"attestation_valid"`
}

type Evidence struct {
	Artifact *ArtifactEvidence `json:"artifact,omitempty"`
	Runtime  *RuntimeEvidence  `json:"runtime,omitempty"`
}

type ArtifactEvidence struct {
	SignerIdentity                 string                `json:"signer_identity,omitempty"`
	Issuer                         string                `json:"issuer,omitempty"`
	Subject                        string                `json:"subject,omitempty"`
	Repository                     string                `json:"repository,omitempty"`
	Workflow                       string                `json:"workflow,omitempty"`
	Ref                            string                `json:"ref,omitempty"`
	CommitSHA                      string                `json:"commit_sha,omitempty"`
	Digest                         string                `json:"digest,omitempty"`
	MatchedIdentity                string                `json:"matched_identity,omitempty"`
	AttestationPredicate           string                `json:"attestation_predicate_type,omitempty"`
	AttestationSubjectName         string                `json:"attestation_subject_name,omitempty"`
	AttestationSubjectDigest       string                `json:"attestation_subject_digest,omitempty"`
	SBOMFormat                     string                `json:"sbom_format,omitempty"`
	SBOMDigestRef                  string                `json:"sbom_digest_ref,omitempty"`
	SBOMHash                       string                `json:"sbom_hash,omitempty"`
	SBOMArtifactRef                string                `json:"sbom_artifact_ref,omitempty"`
	VulnerabilityScanStatus        string                `json:"vulnerability_scan_status,omitempty"`
	VulnerabilityScanTool          string                `json:"vulnerability_scan_tool,omitempty"`
	VulnerabilitySeverityThreshold string                `json:"vulnerability_scan_severity_threshold,omitempty"`
	VulnerabilitySummary           *VulnerabilitySummary `json:"vulnerability_summary,omitempty"`
	VulnerabilityReportRef         string                `json:"vulnerability_report_ref,omitempty"`
}

type VulnerabilitySummary struct {
	Critical int `json:"critical,omitempty"`
	High     int `json:"high,omitempty"`
	Medium   int `json:"medium,omitempty"`
	Low      int `json:"low,omitempty"`
	Unknown  int `json:"unknown,omitempty"`
	Total    int `json:"total,omitempty"`
}

type RuntimeEvidence struct {
	ApprovedDigest            string                           `json:"approved_digest,omitempty"`
	RunningDigest             string                           `json:"running_digest,omitempty"`
	ExpectedConfigHash        string                           `json:"expected_config_hash,omitempty"`
	ActualConfigHash          string                           `json:"actual_config_hash,omitempty"`
	MissingContainers         []string                         `json:"missing_containers,omitempty"`
	UnexpectedContainers      []string                         `json:"unexpected_containers,omitempty"`
	ImageMismatches           []RuntimeImageMismatch           `json:"image_mismatches,omitempty"`
	SecurityContextMismatches []RuntimeSecurityContextMismatch `json:"security_context_mismatches,omitempty"`
}

type RuntimeImageMismatch struct {
	Container      string `json:"container,omitempty"`
	ApprovedImage  string `json:"approved_image,omitempty"`
	RunningImage   string `json:"running_image,omitempty"`
	ApprovedDigest string `json:"approved_digest,omitempty"`
	RunningDigest  string `json:"running_digest,omitempty"`
}

type RuntimeSecurityContextMismatch struct {
	Container string `json:"container,omitempty"`
	Field     string `json:"field,omitempty"`
	Expected  bool   `json:"expected"`
	Actual    bool   `json:"actual"`
}

func FromArtifactVerification(result *verify.ArtifactVerification) (*VerifierSummary, *Evidence) {
	if result == nil {
		return nil, nil
	}

	summary := &VerifierSummary{
		SignatureValid:   result.SignatureValid,
		AttestationValid: result.AttestationValid,
	}

	artifact := &ArtifactEvidence{
		SignerIdentity:           result.VerifiedIdentity,
		Issuer:                   result.VerifiedIssuer,
		Subject:                  result.VerifiedSubject,
		Repository:               result.VerifiedRepo,
		Workflow:                 result.VerifiedWorkflow,
		Ref:                      result.VerifiedRef,
		CommitSHA:                result.VerifiedCommitSHA,
		Digest:                   result.VerifiedDigest,
		MatchedIdentity:          result.Evidence.MatchedIdentity,
		AttestationPredicate:     result.Evidence.AttestationPredicateType,
		AttestationSubjectName:   result.Evidence.AttestationSubjectName,
		AttestationSubjectDigest: result.Evidence.AttestationSubjectDigest,
	}
	if result.Evidence.SupplyChain != nil {
		artifact.VulnerabilityScanStatus = result.Evidence.SupplyChain.VulnerabilityScanStatus
		artifact.VulnerabilityScanTool = result.Evidence.SupplyChain.VulnerabilityScanTool
		artifact.VulnerabilitySeverityThreshold = result.Evidence.SupplyChain.VulnerabilityScanSeverityThreshold
		artifact.VulnerabilityReportRef = result.Evidence.SupplyChain.VulnerabilityReportRef
		if result.Evidence.SupplyChain.VulnerabilitySummary != nil {
			artifact.VulnerabilitySummary = &VulnerabilitySummary{
				Critical: result.Evidence.SupplyChain.VulnerabilitySummary.Critical,
				High:     result.Evidence.SupplyChain.VulnerabilitySummary.High,
				Medium:   result.Evidence.SupplyChain.VulnerabilitySummary.Medium,
				Low:      result.Evidence.SupplyChain.VulnerabilitySummary.Low,
				Unknown:  result.Evidence.SupplyChain.VulnerabilitySummary.Unknown,
				Total:    result.Evidence.SupplyChain.VulnerabilitySummary.Total,
			}
		}
		if result.Evidence.SupplyChain.MatchesDigest(result.VerifiedDigest) {
			artifact.SBOMFormat = result.Evidence.SupplyChain.SBOMFormat
			artifact.SBOMDigestRef = result.Evidence.SupplyChain.SBOMDigestRef
			artifact.SBOMHash = result.Evidence.SupplyChain.SBOMHash
			artifact.SBOMArtifactRef = result.Evidence.SupplyChain.SBOMArtifactRef
		}
	}

	evidence := &Evidence{Artifact: artifact}
	if isEmptyArtifactEvidence(*artifact) {
		evidence.Artifact = nil
	}
	if isEmptyEvidence(*evidence) {
		evidence = nil
	}

	return summary, evidence
}

func FromRuntimeComparison(result *runtimestate.ComparisonResult) *Evidence {
	if result == nil || result.Evidence == nil {
		return nil
	}

	runtimeEvidence := &RuntimeEvidence{
		ApprovedDigest:       result.ApprovedDigest,
		RunningDigest:        result.RunningDigest,
		ExpectedConfigHash:   result.Evidence.ConfigExpectation,
		ActualConfigHash:     result.Evidence.ConfigObserved,
		MissingContainers:    append([]string(nil), result.Evidence.MissingContainers...),
		UnexpectedContainers: append([]string(nil), result.Evidence.UnexpectedContainers...),
	}

	for _, mismatch := range result.Evidence.ImageMismatches {
		runtimeEvidence.ImageMismatches = append(runtimeEvidence.ImageMismatches, RuntimeImageMismatch{
			Container:      mismatch.Container,
			ApprovedImage:  mismatch.ApprovedImage,
			RunningImage:   mismatch.RunningImage,
			ApprovedDigest: mismatch.ApprovedDigest,
			RunningDigest:  mismatch.RunningDigest,
		})
	}
	for _, mismatch := range result.Evidence.SecurityContextMismatches {
		runtimeEvidence.SecurityContextMismatches = append(runtimeEvidence.SecurityContextMismatches, RuntimeSecurityContextMismatch{
			Container: mismatch.Container,
			Field:     mismatch.Field,
			Expected:  mismatch.Expected,
			Actual:    mismatch.Actual,
		})
	}

	evidence := &Evidence{Runtime: runtimeEvidence}
	if isEmptyRuntimeEvidence(*runtimeEvidence) {
		evidence.Runtime = nil
	}
	if isEmptyEvidence(*evidence) {
		return nil
	}
	return evidence
}

func DigestFromImage(image string) string {
	parts := strings.SplitN(strings.TrimSpace(image), "@", 2)
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func BranchFromRef(ref string) string {
	trimmed := strings.TrimSpace(ref)
	trimmed = strings.TrimPrefix(trimmed, "refs/heads/")
	trimmed = strings.TrimPrefix(trimmed, "refs/tags/")
	return trimmed
}

func EnvironmentFromNamespace(namespace string) string {
	namespace = strings.TrimSpace(namespace)
	if namespace == "" {
		return ""
	}
	parts := strings.SplitN(namespace, "-", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return namespace
}

func TenantFromNamespace(namespace string) string {
	namespace = strings.TrimSpace(namespace)
	if namespace == "" {
		return ""
	}
	parts := strings.SplitN(namespace, "-", 2)
	return parts[0]
}

func FirstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func isEmptyEvidence(evidence Evidence) bool {
	return evidence == Evidence{}
}

func isEmptyArtifactEvidence(evidence ArtifactEvidence) bool {
	return evidence == ArtifactEvidence{}
}

func isEmptyRuntimeEvidence(evidence RuntimeEvidence) bool {
	return evidence.ApprovedDigest == "" &&
		evidence.RunningDigest == "" &&
		evidence.ExpectedConfigHash == "" &&
		evidence.ActualConfigHash == "" &&
		len(evidence.MissingContainers) == 0 &&
		len(evidence.UnexpectedContainers) == 0 &&
		len(evidence.ImageMismatches) == 0 &&
		len(evidence.SecurityContextMismatches) == 0
}
