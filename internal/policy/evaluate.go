package policy

import (
	"fmt"
	"path"
	"strings"

	"github.com/denisgrosek/changelock/internal/verify"
)

type Decision struct {
	Decision string   `json:"decision"`
	Reasons  []string `json:"reasons"`
}

type ChangeEvaluationRequest struct {
	Tenant             string   `json:"tenant"`
	Repository         string   `json:"repository"`
	Branch             string   `json:"branch"`
	PullRequest        bool     `json:"pullRequest"`
	SignedCommits      bool     `json:"signedCommits"`
	Approvals          int      `json:"approvals"`
	SecurityApprovals  int      `json:"securityApprovals"`
	CodeOwnersApproved bool     `json:"codeOwnersApproved"`
	ForcePush          bool     `json:"forcePush"`
	ChangedFiles       []string `json:"changedFiles"`
}

type ArtifactEvaluationRequest struct {
	Tenant         string                       `json:"tenant"`
	Repository     string                       `json:"repository"`
	Image          string                       `json:"image"`
	Registry       string                       `json:"registry"`
	DigestPinned   bool                         `json:"digestPinned"`
	HasProvenance  bool                         `json:"hasProvenance"`
	HasSignature   bool                         `json:"hasSignature"`
	SignerIdentity string                       `json:"signerIdentity"`
	WorkflowFile   string                       `json:"workflowFile"`
	Subject        string                       `json:"subject"`
	Verification   *verify.ArtifactVerification `json:"verification,omitempty"`
}

func EvaluateChange(bundle *Bundle, request ChangeEvaluationRequest) Decision {
	reasons := []string{}

	if !bundle.RepositoryAllowed(request.Repository) {
		reasons = append(reasons, fmt.Sprintf("repository %q is not allowed for tenant %q", request.Repository, bundle.Tenant.Metadata.Name))
	}
	if !matchesAny(request.Branch, bundle.Change.Spec.AllowedBranches) {
		reasons = append(reasons, fmt.Sprintf("branch %q is not allowed", request.Branch))
	}
	if bundle.Change.Spec.RequirePullRequest && !request.PullRequest {
		reasons = append(reasons, "pull request is required")
	}
	if bundle.Change.Spec.RequireSignedCommits && !request.SignedCommits {
		reasons = append(reasons, "signed commits are required")
	}
	if request.Approvals < bundle.Change.Spec.MinimumApprovals {
		reasons = append(reasons, fmt.Sprintf("minimum approvals not met: need %d", bundle.Change.Spec.MinimumApprovals))
	}
	if request.ForcePush && bundle.Change.Spec.BlockForcePushOnProtectedBranches && matchesAny(request.Branch, bundle.Change.Spec.AllowedBranches) {
		reasons = append(reasons, "force push is blocked on protected branches")
	}

	if hasCriticalPathChange(request.ChangedFiles, bundle.AllCriticalPathPatterns()) {
		if request.SecurityApprovals < bundle.Change.Spec.CriticalPathRules.MinimumSecurityApprovals {
			reasons = append(reasons, fmt.Sprintf("security approvals not met for critical path change: need %d", bundle.Change.Spec.CriticalPathRules.MinimumSecurityApprovals))
		}
		if bundle.Change.Spec.CriticalPathRules.RequireCodeOwnersApproval && !request.CodeOwnersApproved {
			reasons = append(reasons, "CODEOWNERS approval is required for critical path change")
		}
	}

	return decisionFromReasons(reasons)
}

func EvaluateArtifact(bundle *Bundle, request ArtifactEvaluationRequest) Decision {
	reasons := []string{}
	effectiveRepository := request.Repository
	effectiveDigestPinned := request.DigestPinned
	effectiveHasProvenance := request.HasProvenance
	effectiveHasSignature := request.HasSignature
	effectiveSignerIdentity := request.SignerIdentity
	effectiveWorkflowFile := request.WorkflowFile
	effectiveSubject := request.Subject

	if request.Verification != nil {
		effectiveRepository = firstNonEmpty(request.Verification.VerifiedRepo, effectiveRepository)
		effectiveDigestPinned = effectiveDigestPinned || request.Verification.VerifiedDigest != ""
		effectiveHasProvenance = request.Verification.AttestationValid
		effectiveHasSignature = request.Verification.SignatureValid
		effectiveSignerIdentity = firstNonEmpty(request.Verification.VerifiedIdentity, effectiveSignerIdentity)
		effectiveWorkflowFile = firstNonEmpty(request.Verification.VerifiedWorkflow, effectiveWorkflowFile)
		effectiveSubject = firstNonEmpty(request.Verification.VerifiedSubject, effectiveSubject)
		reasons = append(reasons, request.Verification.Reasons...)
	}

	if !bundle.RepositoryAllowed(effectiveRepository) {
		reasons = append(reasons, fmt.Sprintf("repository %q is not allowed for tenant %q", effectiveRepository, bundle.Tenant.Metadata.Name))
	}

	image := strings.TrimSpace(request.Image)
	if !hasAllowedRegistry(image, request.Registry, bundle.Artifact.Spec.AllowedRegistries) {
		reasons = append(reasons, fmt.Sprintf("image %q is not from an allowed registry", image))
	}

	digestPinned := effectiveDigestPinned || strings.Contains(image, "@sha256:")
	if bundle.Artifact.Spec.RequireDigestPinning && !digestPinned {
		reasons = append(reasons, "image must be digest-pinned")
	}

	if bundle.Artifact.Spec.RequireProvenance && !effectiveHasProvenance {
		reasons = append(reasons, "provenance attestation verification failed")
	}
	if bundle.Artifact.Spec.RequireSignature && !effectiveHasSignature {
		reasons = append(reasons, "artifact signature verification failed")
	}
	if !containsExact(effectiveWorkflowFile, bundle.AllowedWorkflowFiles(effectiveRepository)) {
		reasons = append(reasons, fmt.Sprintf("workflow file %q is not allowed", effectiveWorkflowFile))
	}
	if !containsExact(effectiveSubject, bundle.Artifact.Spec.AllowedSubjects) {
		reasons = append(reasons, fmt.Sprintf("subject %q is not allowed", effectiveSubject))
	}
	if !containsExact(effectiveSignerIdentity, bundle.Artifact.Spec.AllowedSignerIdentities) {
		reasons = append(reasons, fmt.Sprintf("signer identity %q is not allowed", effectiveSignerIdentity))
	}

	return decisionFromReasons(reasons)
}

func decisionFromReasons(reasons []string) Decision {
	if len(reasons) == 0 {
		return Decision{
			Decision: "ALLOW",
			Reasons:  []string{},
		}
	}
	return Decision{
		Decision: "DENY",
		Reasons:  reasons,
	}
}

func hasCriticalPathChange(files []string, patterns []string) bool {
	for _, file := range files {
		if matchesAny(file, patterns) {
			return true
		}
	}
	return false
}

func matchesAny(value string, patterns []string) bool {
	for _, pattern := range patterns {
		if matchesPattern(value, pattern) {
			return true
		}
	}
	return false
}

func matchesPattern(value, pattern string) bool {
	switch {
	case pattern == value:
		return true
	case strings.Contains(pattern, "**"):
		prefix := strings.TrimSuffix(pattern, "**")
		return strings.HasPrefix(value, prefix)
	case strings.HasSuffix(pattern, "/*"):
		prefix := strings.TrimSuffix(pattern, "*")
		return strings.HasPrefix(value, prefix)
	case strings.Contains(pattern, "*"):
		matched, err := path.Match(pattern, value)
		return err == nil && matched
	default:
		return value == pattern
	}
}

func hasAllowedRegistry(image, explicitRegistry string, allowed []string) bool {
	if explicitRegistry != "" && containsPrefix(explicitRegistry, allowed) {
		return true
	}
	return containsPrefix(image, allowed)
}

func containsExact(value string, allowed []string) bool {
	for _, candidate := range allowed {
		if value == candidate {
			return true
		}
	}
	return false
}

func containsPrefix(value string, allowed []string) bool {
	for _, candidate := range allowed {
		if strings.HasPrefix(value, candidate) {
			return true
		}
	}
	return false
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}
