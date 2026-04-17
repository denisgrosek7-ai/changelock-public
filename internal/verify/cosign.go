package verify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/evidence"
)

type commandOutput struct {
	Stdout []byte
	Stderr []byte
}

type commandRunner interface {
	Run(ctx context.Context, name string, args ...string) (commandOutput, error)
}

type execRunner struct{}

func (execRunner) Run(ctx context.Context, name string, args ...string) (commandOutput, error) {
	cmd := exec.CommandContext(ctx, name, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return commandOutput{
		Stdout: stdout.Bytes(),
		Stderr: stderr.Bytes(),
	}, err
}

type CosignVerifier struct {
	binary         string
	runner         commandRunner
	evidenceConfig evidence.Config
	tlogClient     evidence.TLogClient
}

func NewCosignVerifier(binary string) *CosignVerifier {
	if strings.TrimSpace(binary) == "" {
		binary = "cosign"
	}
	return &CosignVerifier{
		binary:     binary,
		runner:     execRunner{},
		tlogClient: evidence.NewHTTPTLogClient(5 * time.Second),
	}
}

func NewCosignVerifierWithEvidence(binary string, cfg evidence.Config) *CosignVerifier {
	verifier := NewCosignVerifier(binary)
	verifier.evidenceConfig = cfg
	return verifier
}

func (v *CosignVerifier) VerifyArtifact(ctx context.Context, request ArtifactVerificationRequest) (ArtifactVerification, error) {
	verification := ArtifactVerification{
		VerifiedDigest: digestFromImage(request.Image),
	}
	mergeSupplyChainEvidence(&verification.Evidence.SupplyChain, request.SupplyChain)

	if verification.VerifiedDigest == "" {
		verification.Reasons = append(verification.Reasons, "image must be digest-pinned for verification")
		return verification, nil
	}

	signatureResult, signatureErr := v.verifySignature(ctx, request)
	if signatureErr != nil {
		return verification, signatureErr
	}
	mergeVerification(&verification, signatureResult)

	attestationResult, attestationErr := v.verifyAttestation(ctx, request, verification.VerifiedIdentity)
	if attestationErr != nil {
		return verification, attestationErr
	}
	mergeVerification(&verification, attestationResult)

	state, reason, verifiedAt := evidence.EvaluateBundle(ctx, v.evidenceConfig, request.EvidenceBundle, verification.VerifiedDigest, v.tlogClient, time.Now)
	verification.Evidence.TransparencyLogState = state
	verification.Evidence.TransparencyLogReason = reason
	verification.Evidence.Bundle = evidence.ApplyVerificationResult(request.EvidenceBundle, state, reason, verifiedAt)
	if shouldEnforceTransparencyEvidence(v.evidenceConfig) && state != evidence.StateVerified {
		verification.Reasons = append(verification.Reasons, "transparency log evidence verification failed: "+reason)
	}

	return verification, nil
}

func (v *CosignVerifier) verifySignature(ctx context.Context, request ArtifactVerificationRequest) (ArtifactVerification, error) {
	result := ArtifactVerification{
		VerifiedDigest: digestFromImage(request.Image),
	}

	identities := request.AllowedSignerIdentities
	if len(identities) == 0 {
		result.Reasons = append(result.Reasons, "no allowed signer identities configured for signature verification")
		return result, nil
	}
	issuers := request.AllowedOIDCIssuers
	if len(issuers) == 0 {
		result.Reasons = append(result.Reasons, "no allowed OIDC issuers configured for signature verification")
		return result, nil
	}

	attemptReasons := []string{}
	for _, identity := range identities {
		repository, _, ref := parseWorkflowIdentity(identity)
		expectedRepository := firstNonEmpty(request.ExpectedRepository, repository)
		expectedRef := firstNonEmpty(request.ExpectedRef, ref)

		for _, issuer := range issuers {
			args := []string{
				"verify",
				"--output", "json",
				"--certificate-identity", identity,
				"--certificate-oidc-issuer", issuer,
			}
			if expectedRepository != "" {
				args = append(args, "--certificate-github-workflow-repository", expectedRepository)
			}
			if expectedRef != "" {
				args = append(args, "--certificate-github-workflow-ref", expectedRef)
			}
			if request.ExpectedCommitSHA != "" {
				args = append(args, "--certificate-github-workflow-sha", request.ExpectedCommitSHA)
			}
			args = append(args, request.Image)

			output, err := v.runner.Run(ctx, v.binary, args...)
			if err != nil {
				if errors.Is(err, exec.ErrNotFound) {
					return result, fmt.Errorf("cosign binary %q not found: %w", v.binary, err)
				}
				attemptReasons = append(attemptReasons, formatCommandFailure("signature verification", identity, issuer, output, err))
				continue
			}

			claims, parseErr := parseJSONObjects(output.Stdout)
			if parseErr != nil {
				return result, fmt.Errorf("parse cosign verify output: %w", parseErr)
			}

			_, workflowFile, workflowRef := parseWorkflowIdentity(identity)
			result.SignatureValid = true
			result.VerifiedIdentity = identity
			result.VerifiedIssuer = issuer
			result.VerifiedRepo = expectedRepository
			result.VerifiedWorkflow = workflowFile
			result.VerifiedRef = firstNonEmpty(expectedRef, workflowRef)
			result.VerifiedCommitSHA = request.ExpectedCommitSHA
			result.VerifiedSubject = repoSubject(expectedRepository)
			result.Evidence.MatchedIdentity = identity
			result.Evidence.SignatureClaimsCount = len(claims)
			return result, nil
		}
	}

	if len(attemptReasons) == 0 {
		attemptReasons = append(attemptReasons, "signature verification failed without diagnostic output")
	}
	result.Reasons = append(result.Reasons, attemptReasons...)
	return result, nil
}

func mergeVerification(dst *ArtifactVerification, src ArtifactVerification) {
	dst.SignatureValid = dst.SignatureValid || src.SignatureValid
	dst.AttestationValid = dst.AttestationValid || src.AttestationValid
	dst.VerifiedIdentity = firstNonEmpty(dst.VerifiedIdentity, src.VerifiedIdentity)
	dst.VerifiedIssuer = firstNonEmpty(dst.VerifiedIssuer, src.VerifiedIssuer)
	dst.VerifiedSubject = firstNonEmpty(dst.VerifiedSubject, src.VerifiedSubject)
	dst.VerifiedRepo = firstNonEmpty(dst.VerifiedRepo, src.VerifiedRepo)
	dst.VerifiedWorkflow = firstNonEmpty(dst.VerifiedWorkflow, src.VerifiedWorkflow)
	dst.VerifiedRef = firstNonEmpty(dst.VerifiedRef, src.VerifiedRef)
	dst.VerifiedCommitSHA = firstNonEmpty(dst.VerifiedCommitSHA, src.VerifiedCommitSHA)
	dst.VerifiedDigest = firstNonEmpty(dst.VerifiedDigest, src.VerifiedDigest)

	if dst.Evidence.MatchedIdentity == "" {
		dst.Evidence.MatchedIdentity = src.Evidence.MatchedIdentity
	}
	if dst.Evidence.SignatureClaimsCount == 0 {
		dst.Evidence.SignatureClaimsCount = src.Evidence.SignatureClaimsCount
	}
	if dst.Evidence.AttestationCount == 0 {
		dst.Evidence.AttestationCount = src.Evidence.AttestationCount
	}
	if dst.Evidence.AttestationPredicateType == "" {
		dst.Evidence.AttestationPredicateType = src.Evidence.AttestationPredicateType
	}
	if dst.Evidence.AttestationSubjectName == "" {
		dst.Evidence.AttestationSubjectName = src.Evidence.AttestationSubjectName
	}
	if dst.Evidence.AttestationSubjectDigest == "" {
		dst.Evidence.AttestationSubjectDigest = src.Evidence.AttestationSubjectDigest
	}
	mergeSupplyChainEvidence(&dst.Evidence.SupplyChain, src.Evidence.SupplyChain)
	if dst.Evidence.Bundle == nil && src.Evidence.Bundle != nil {
		dst.Evidence.Bundle = evidence.CloneBundle(src.Evidence.Bundle)
	}
	if dst.Evidence.TransparencyLogState == "" {
		dst.Evidence.TransparencyLogState = src.Evidence.TransparencyLogState
	}
	if dst.Evidence.TransparencyLogReason == "" {
		dst.Evidence.TransparencyLogReason = src.Evidence.TransparencyLogReason
	}

	dst.Reasons = append(dst.Reasons, src.Reasons...)
}

func shouldEnforceTransparencyEvidence(cfg evidence.Config) bool {
	return cfg.Mode == evidence.ModeRekorRequired && cfg.VerifyOnDeploy
}

func mergeSupplyChainEvidence(dst **SupplyChainEvidence, src *SupplyChainEvidence) {
	if src == nil {
		return
	}
	if *dst == nil {
		*dst = cloneSupplyChainEvidence(src)
		return
	}

	target := *dst
	if target.SBOMFormat == "" {
		target.SBOMFormat = src.SBOMFormat
	}
	if target.SBOMDigestRef == "" {
		target.SBOMDigestRef = src.SBOMDigestRef
	}
	if target.SBOMHash == "" {
		target.SBOMHash = src.SBOMHash
	}
	if target.SBOMArtifactRef == "" {
		target.SBOMArtifactRef = src.SBOMArtifactRef
	}
	if target.VulnerabilityScanStatus == "" {
		target.VulnerabilityScanStatus = src.VulnerabilityScanStatus
	}
	if target.VulnerabilityScanTool == "" {
		target.VulnerabilityScanTool = src.VulnerabilityScanTool
	}
	if target.VulnerabilityScanSeverityThreshold == "" {
		target.VulnerabilityScanSeverityThreshold = src.VulnerabilityScanSeverityThreshold
	}
	if target.VulnerabilitySummary == nil && src.VulnerabilitySummary != nil {
		target.VulnerabilitySummary = cloneVulnerabilitySummary(src.VulnerabilitySummary)
	}
	if target.VulnerabilityReportRef == "" {
		target.VulnerabilityReportRef = src.VulnerabilityReportRef
	}
}

func cloneSupplyChainEvidence(src *SupplyChainEvidence) *SupplyChainEvidence {
	if src == nil {
		return nil
	}
	clone := *src
	if src.VulnerabilitySummary != nil {
		clone.VulnerabilitySummary = cloneVulnerabilitySummary(src.VulnerabilitySummary)
	}
	return &clone
}

func cloneVulnerabilitySummary(src *VulnerabilitySummary) *VulnerabilitySummary {
	if src == nil {
		return nil
	}
	clone := *src
	return &clone
}

func parseJSONObjects(data []byte) ([]json.RawMessage, error) {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return nil, errors.New("empty JSON output")
	}

	if json.Valid(trimmed) && len(trimmed) > 0 && trimmed[0] == '{' {
		return []json.RawMessage{append([]byte(nil), trimmed...)}, nil
	}

	var array []json.RawMessage
	if err := json.Unmarshal(trimmed, &array); err == nil {
		return array, nil
	}

	decoder := json.NewDecoder(bytes.NewReader(trimmed))
	items := []json.RawMessage{}
	for decoder.More() {
		var raw json.RawMessage
		if err := decoder.Decode(&raw); err != nil {
			return nil, err
		}
		items = append(items, raw)
	}
	if len(items) == 0 {
		return nil, errors.New("no JSON objects found")
	}
	return items, nil
}

func formatCommandFailure(operation, identity, issuer string, output commandOutput, err error) string {
	detail := strings.TrimSpace(string(bytes.TrimSpace(output.Stderr)))
	if detail == "" {
		detail = strings.TrimSpace(string(bytes.TrimSpace(output.Stdout)))
	}
	if detail == "" {
		detail = err.Error()
	}
	return fmt.Sprintf("%s failed for identity %q and issuer %q: %s", operation, identity, issuer, detail)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func repoSubject(repository string) string {
	repository = strings.TrimSpace(repository)
	if repository == "" {
		return ""
	}
	return "repo:" + repository
}
