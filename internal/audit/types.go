package audit

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/evidence"
	runtimestate "github.com/denisgrosek/changelock/internal/runtime"
	"github.com/denisgrosek/changelock/internal/verify"
)

const (
	EventTypeArtifactVerificationResult      = "artifact_verification_result"
	EventTypeDeployGateDecision              = "deploy_gate_decision"
	EventTypePolicyDecision                  = "policy_decision"
	EventTypeRuntimeDriftResult              = "runtime_drift_result"
	EventTypeRuntimeDesiredStateRecorded     = "runtime_desired_state_recorded"
	EventTypeRuntimeActiveStateObserved      = "runtime_active_state_observed"
	EventTypeDriftDetected                   = "drift_detected"
	EventTypeDriftRemediationStarted         = "drift_remediation_started"
	EventTypeDriftRemediationSucceeded       = "drift_remediation_succeeded"
	EventTypeDriftRemediationFailed          = "drift_remediation_failed"
	EventTypeDriftQuarantined                = "drift_quarantined"
	EventTypeSigningIdentityPolicyRecorded   = "signing_identity_policy_recorded"
	EventTypeSigningIdentityPolicyDistrusted = "signing_identity_policy_distrusted"
	EventTypeSigningIdentityFinding          = "signing_identity_finding"
	EventTypeHandoffSealed                   = "handoff_sealed"
	EventTypeHandoffCosigned                 = "handoff_cosigned"

	DecisionAllow = "ALLOW"
	DecisionDeny  = "DENY"
	DecisionError = "ERROR"
)

type Event struct {
	RequestID                         string           `json:"request_id"`
	Timestamp                         time.Time        `json:"timestamp"`
	Component                         string           `json:"component"`
	EventType                         string           `json:"event_type"`
	Actor                             string           `json:"actor,omitempty"`
	ClusterID                         string           `json:"cluster_id,omitempty"`
	TenantID                          string           `json:"tenant_id,omitempty"`
	Repo                              string           `json:"repo,omitempty"`
	Branch                            string           `json:"branch,omitempty"`
	Environment                       string           `json:"environment,omitempty"`
	Namespace                         string           `json:"namespace,omitempty"`
	WorkloadKind                      string           `json:"workload_kind,omitempty"`
	Workload                          string           `json:"workload,omitempty"`
	ServiceAccount                    string           `json:"service_account,omitempty"`
	Image                             string           `json:"image,omitempty"`
	Digest                            string           `json:"digest,omitempty"`
	CVEID                             string           `json:"cve_id,omitempty"`
	Decision                          string           `json:"decision"`
	Reasons                           []string         `json:"reasons,omitempty"`
	DriftResult                       string           `json:"drift_result,omitempty"`
	DriftClasses                      []string         `json:"drift_classes,omitempty"`
	DriftSeverity                     string           `json:"drift_severity,omitempty"`
	ReconciliationStatus              string           `json:"reconciliation_status,omitempty"`
	RemediationMode                   string           `json:"remediation_mode,omitempty"`
	RemediationAttempt                int              `json:"remediation_attempt,omitempty"`
	Remediable                        bool             `json:"remediable,omitempty"`
	QuarantineReason                  string           `json:"quarantine_reason,omitempty"`
	QuarantineType                    string           `json:"quarantine_type,omitempty"`
	ProtectedTarget                   bool             `json:"protected_target,omitempty"`
	ProtectedReason                   string           `json:"protected_reason,omitempty"`
	DesiredStateSourceRef             string           `json:"desired_state_source_ref,omitempty"`
	DesiredStateApprovalID            string           `json:"desired_state_approval_id,omitempty"`
	DesiredStateVerification          string           `json:"desired_state_verification_state,omitempty"`
	VerifierSummary                   *VerifierSummary `json:"verifier_summary,omitempty"`
	PolicyVersion                     string           `json:"policy_version,omitempty"`
	PolicyBundleID                    string           `json:"policy_bundle_id,omitempty"`
	PolicyBundleHash                  string           `json:"policy_bundle_hash,omitempty"`
	DecisionHash                      string           `json:"decision_hash,omitempty"`
	IsException                       bool             `json:"is_exception,omitempty"`
	ExceptionID                       string           `json:"exception_id,omitempty"`
	ExceptionType                     string           `json:"exception_type,omitempty"`
	ExceptionStatus                   string           `json:"exception_status,omitempty"`
	ExceptionReason                   string           `json:"exception_reason,omitempty"`
	ExceptionTicketID                 string           `json:"exception_ticket_id,omitempty"`
	ExceptionRequestedBy              string           `json:"exception_requested_by,omitempty"`
	ExceptionRequestedAt              *time.Time       `json:"exception_requested_at,omitempty"`
	ExceptionApprovedBy               string           `json:"exception_approved_by,omitempty"`
	ExceptionApprovedAt               *time.Time       `json:"exception_approved_at,omitempty"`
	ExceptionRejectedBy               string           `json:"exception_rejected_by,omitempty"`
	ExceptionRejectedAt               *time.Time       `json:"exception_rejected_at,omitempty"`
	ExceptionRejectionReason          string           `json:"exception_rejection_reason,omitempty"`
	ExceptionExpiresAt                *time.Time       `json:"exception_expires_at,omitempty"`
	IncidentID                        string           `json:"incident_id,omitempty"`
	IncidentIdentityKey               string           `json:"incident_identity_key,omitempty"`
	IncidentTitle                     string           `json:"incident_title,omitempty"`
	IncidentSummary                   string           `json:"incident_summary,omitempty"`
	IncidentCategory                  string           `json:"incident_category,omitempty"`
	IncidentSeverity                  string           `json:"incident_severity,omitempty"`
	IncidentPriority                  string           `json:"incident_priority,omitempty"`
	IncidentState                     string           `json:"incident_state,omitempty"`
	IncidentScopeType                 string           `json:"incident_scope_type,omitempty"`
	IncidentScopeRef                  string           `json:"incident_scope_ref,omitempty"`
	IncidentOwner                     string           `json:"incident_owner,omitempty"`
	IncidentAssignmentReason          string           `json:"incident_assignment_reason,omitempty"`
	IncidentResolutionType            string           `json:"incident_resolution_type,omitempty"`
	IncidentResolutionSummary         string           `json:"incident_resolution_summary,omitempty"`
	IncidentResolutionDetails         string           `json:"incident_resolution_details,omitempty"`
	IncidentFollowUpRequired          bool             `json:"incident_follow_up_required,omitempty"`
	IncidentResolutionRefs            []string         `json:"incident_resolution_refs,omitempty"`
	IncidentFindingRefs               []string         `json:"incident_finding_refs,omitempty"`
	IncidentGuidanceRefs              []string         `json:"incident_guidance_refs,omitempty"`
	IncidentScorecardRefs             []string         `json:"incident_scorecard_refs,omitempty"`
	IncidentEvidenceRefs              []string         `json:"incident_evidence_refs,omitempty"`
	IncidentReasonCodes               []string         `json:"incident_reason_codes,omitempty"`
	IncidentLabels                    []string         `json:"incident_labels,omitempty"`
	IncidentNote                      string           `json:"incident_note,omitempty"`
	RecommendationID                  string           `json:"recommendation_id,omitempty"`
	RecommendationSourceType          string           `json:"recommendation_source_type,omitempty"`
	RecommendationSourceRef           string           `json:"recommendation_source_ref,omitempty"`
	RecommendationSubjectType         string           `json:"recommendation_subject_type,omitempty"`
	RecommendationSubjectRef          string           `json:"recommendation_subject_ref,omitempty"`
	RecommendationType                string           `json:"recommendation_type,omitempty"`
	RecommendationTitle               string           `json:"recommendation_title,omitempty"`
	RecommendationDescription         string           `json:"recommendation_description,omitempty"`
	RecommendationAction              string           `json:"recommendation_action,omitempty"`
	RecommendationRationale           string           `json:"recommendation_rationale,omitempty"`
	RecommendationEvidenceRefs        []string         `json:"recommendation_evidence_refs,omitempty"`
	RecommendationReadbackRefs        []string         `json:"recommendation_readback_refs,omitempty"`
	RecommendationRelatedIncidentRefs []string         `json:"recommendation_related_incident_refs,omitempty"`
	RecommendationPriorityBand        string           `json:"recommendation_priority_band,omitempty"`
	RecommendationImpactScore         int              `json:"recommendation_impact_score,omitempty"`
	RecommendationEffortScore         int              `json:"recommendation_effort_score,omitempty"`
	RecommendationConfidenceScore     int              `json:"recommendation_confidence_score,omitempty"`
	RecommendationApprovalMode        string           `json:"recommendation_approval_mode,omitempty"`
	RecommendationStatus              string           `json:"recommendation_status,omitempty"`
	RecommendationTemplateID          string           `json:"recommendation_template_id,omitempty"`
	RecommendationVerificationPlan    []string         `json:"recommendation_verification_plan,omitempty"`
	RecommendationFeedbackSummary     string           `json:"recommendation_feedback_summary,omitempty"`
	RecommendationOwner               string           `json:"recommendation_owner,omitempty"`
	RecommendationComment             string           `json:"recommendation_comment,omitempty"`
	RecommendationSupersededBy        string           `json:"recommendation_superseded_by,omitempty"`
	RecommendationExpiresAt           *time.Time       `json:"recommendation_expires_at,omitempty"`
	RecommendationVerificationResult  string           `json:"recommendation_verification_result,omitempty"`
	Evidence                          *Evidence        `json:"evidence,omitempty"`
	Handoff                           json.RawMessage  `json:"handoff,omitempty"`
}

type VerifierSummary struct {
	SignatureValid   bool `json:"signature_valid"`
	AttestationValid bool `json:"attestation_valid"`
}

type Evidence struct {
	Artifact           *ArtifactEvidence        `json:"artifact,omitempty"`
	Runtime            *RuntimeEvidence         `json:"runtime,omitempty"`
	Bundle             *evidence.Bundle         `json:"bundle,omitempty"`
	VerificationState  string                   `json:"verification_state,omitempty"`
	VerificationReason string                   `json:"verification_reason,omitempty"`
	SigningIdentity    *SigningIdentityEvidence `json:"signing_identity,omitempty"`
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
	ClusterID                 string                           `json:"cluster_id,omitempty"`
	WorkloadKind              string                           `json:"workload_kind,omitempty"`
	ServiceAccountExpected    string                           `json:"service_account_expected,omitempty"`
	ServiceAccountObserved    string                           `json:"service_account_observed,omitempty"`
	ApprovedLabels            map[string]string                `json:"approved_labels,omitempty"`
	ApprovedDigest            string                           `json:"approved_digest,omitempty"`
	RunningDigest             string                           `json:"running_digest,omitempty"`
	ExpectedConfigHash        string                           `json:"expected_config_hash,omitempty"`
	ActualConfigHash          string                           `json:"actual_config_hash,omitempty"`
	ApprovedContainers        []RuntimeApprovedContainer       `json:"approved_containers,omitempty"`
	MissingContainers         []string                         `json:"missing_containers,omitempty"`
	UnexpectedContainers      []string                         `json:"unexpected_containers,omitempty"`
	ImageMismatches           []RuntimeImageMismatch           `json:"image_mismatches,omitempty"`
	SecurityContextMismatches []RuntimeSecurityContextMismatch `json:"security_context_mismatches,omitempty"`
}

type RuntimeApprovedContainer struct {
	Name           string                     `json:"name"`
	Image          string                     `json:"image,omitempty"`
	ApprovedDigest string                     `json:"approved_digest,omitempty"`
	Runtime        RuntimeSecurityConstraints `json:"runtime"`
}

type RuntimeSecurityConstraints struct {
	RunAsNonRoot             bool `json:"run_as_non_root"`
	ReadOnlyRootFilesystem   bool `json:"read_only_root_filesystem"`
	AllowPrivilegeEscalation bool `json:"allow_privilege_escalation"`
	DropAllCapabilities      bool `json:"drop_all_capabilities"`
	SeccompRuntimeDefault    bool `json:"seccomp_runtime_default"`
	DenyPrivileged           bool `json:"deny_privileged"`
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

type SigningIdentityEvidence struct {
	PolicyID              string     `json:"policy_id,omitempty"`
	PolicyName            string     `json:"policy_name,omitempty"`
	ProviderType          string     `json:"provider_type,omitempty"`
	Issuer                string     `json:"issuer,omitempty"`
	SignerIdentity        string     `json:"signer_identity,omitempty"`
	Subject               string     `json:"subject,omitempty"`
	Repository            string     `json:"repository,omitempty"`
	Workflow              string     `json:"workflow,omitempty"`
	Ref                   string     `json:"ref,omitempty"`
	EnforcementMode       string     `json:"enforcement_mode,omitempty"`
	Authorized            string     `json:"authorized,omitempty"`
	ReasonCode            string     `json:"reason_code,omitempty"`
	ReasonDetail          string     `json:"reason_detail,omitempty"`
	DistrustedAfter       *time.Time `json:"distrusted_after,omitempty"`
	WorkflowDriftDetected bool       `json:"workflow_drift_detected,omitempty"`
	TransparencyRequired  bool       `json:"transparency_required,omitempty"`
	TransparencyState     string     `json:"transparency_state,omitempty"`
	TransparencyReason    string     `json:"transparency_reason,omitempty"`
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
	evidence.Bundle = evidencepkgClone(result.Evidence.Bundle)
	evidence.VerificationState = strings.TrimSpace(result.Evidence.TransparencyLogState)
	evidence.VerificationReason = strings.TrimSpace(result.Evidence.TransparencyLogReason)
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
		ClusterID:              result.ClusterID,
		WorkloadKind:           result.WorkloadKind,
		ServiceAccountExpected: result.Evidence.ServiceAccountExpected,
		ServiceAccountObserved: result.Evidence.ServiceAccountObserved,
		ApprovedDigest:         result.ApprovedDigest,
		RunningDigest:          result.RunningDigest,
		ExpectedConfigHash:     result.Evidence.ConfigExpectation,
		ActualConfigHash:       result.Evidence.ConfigObserved,
		MissingContainers:      append([]string(nil), result.Evidence.MissingContainers...),
		UnexpectedContainers:   append([]string(nil), result.Evidence.UnexpectedContainers...),
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
	return evidence.Artifact == nil &&
		evidence.Runtime == nil &&
		evidence.Bundle == nil &&
		evidence.SigningIdentity == nil &&
		strings.TrimSpace(evidence.VerificationState) == "" &&
		strings.TrimSpace(evidence.VerificationReason) == ""
}

func isEmptyArtifactEvidence(evidence ArtifactEvidence) bool {
	return evidence == ArtifactEvidence{}
}

func isEmptyRuntimeEvidence(evidence RuntimeEvidence) bool {
	return evidence.ClusterID == "" &&
		evidence.WorkloadKind == "" &&
		evidence.ServiceAccountExpected == "" &&
		evidence.ServiceAccountObserved == "" &&
		len(evidence.ApprovedLabels) == 0 &&
		evidence.ApprovedDigest == "" &&
		evidence.RunningDigest == "" &&
		evidence.ExpectedConfigHash == "" &&
		evidence.ActualConfigHash == "" &&
		len(evidence.ApprovedContainers) == 0 &&
		len(evidence.MissingContainers) == 0 &&
		len(evidence.UnexpectedContainers) == 0 &&
		len(evidence.ImageMismatches) == 0 &&
		len(evidence.SecurityContextMismatches) == 0
}

func evidencepkgClone(bundle *evidence.Bundle) *evidence.Bundle {
	return evidence.CloneBundle(bundle)
}
