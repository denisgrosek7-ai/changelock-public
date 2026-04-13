package verify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type attestationEnvelope struct {
	Payload string
}

type inTotoStatement struct {
	PredicateType string          `json:"predicateType"`
	Subject       []inTotoSubject `json:"subject"`
}

type inTotoSubject struct {
	Name   string            `json:"name"`
	Digest map[string]string `json:"digest"`
}

func (v *CosignVerifier) verifyAttestation(ctx context.Context, request ArtifactVerificationRequest, matchedIdentity string) (ArtifactVerification, error) {
	result := ArtifactVerification{
		VerifiedDigest: digestFromImage(request.Image),
	}

	identities := []string{}
	if matchedIdentity != "" {
		identities = append(identities, matchedIdentity)
	} else {
		identities = append(identities, request.AllowedSignerIdentities...)
	}
	if len(identities) == 0 {
		result.Reasons = append(result.Reasons, "no allowed signer identities configured for attestation verification")
		return result, nil
	}
	issuers := request.AllowedOIDCIssuers
	if len(issuers) == 0 {
		result.Reasons = append(result.Reasons, "no allowed OIDC issuers configured for attestation verification")
		return result, nil
	}

	predicateType := firstNonEmpty(request.PredicateType, DefaultPredicateType)
	attemptReasons := []string{}
	for _, identity := range identities {
		repository, workflowFile, ref := parseWorkflowIdentity(identity)
		expectedRepository := firstNonEmpty(request.ExpectedRepository, repository)
		expectedRef := firstNonEmpty(request.ExpectedRef, ref)

		for _, issuer := range issuers {
			args := []string{
				"verify-attestation",
				"--output", "json",
				"--type", predicateType,
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
				attemptReasons = append(attemptReasons, formatCommandFailure("attestation verification", identity, issuer, output, err))
				continue
			}

			statements, parseErr := parseAttestationStatements(output.Stdout)
			if parseErr != nil {
				return result, fmt.Errorf("parse cosign verify-attestation output: %w", parseErr)
			}

			subjectName, subjectDigest, subjectErr := matchStatementSubject(statements, request.Image)
			if subjectErr != nil {
				attemptReasons = append(attemptReasons, fmt.Sprintf("attestation verification failed for identity %q and issuer %q: %s", identity, issuer, subjectErr.Error()))
				continue
			}

			result.AttestationValid = true
			result.VerifiedIdentity = identity
			result.VerifiedIssuer = issuer
			result.VerifiedRepo = expectedRepository
			result.VerifiedWorkflow = workflowFile
			result.VerifiedRef = expectedRef
			result.VerifiedCommitSHA = request.ExpectedCommitSHA
			result.VerifiedDigest = subjectDigest
			result.VerifiedSubject = repoSubject(expectedRepository)
			result.Evidence.MatchedIdentity = identity
			result.Evidence.AttestationCount = len(statements)
			result.Evidence.AttestationPredicateType = firstNonEmpty(statements[0].PredicateType, predicateType)
			result.Evidence.AttestationSubjectName = subjectName
			result.Evidence.AttestationSubjectDigest = subjectDigest
			return result, nil
		}
	}

	if len(attemptReasons) == 0 {
		attemptReasons = append(attemptReasons, "attestation verification failed without diagnostic output")
	}
	result.Reasons = append(result.Reasons, attemptReasons...)
	return result, nil
}

func parseAttestationStatements(data []byte) ([]inTotoStatement, error) {
	items, err := parseJSONObjects(data)
	if err != nil {
		return nil, err
	}

	statements := make([]inTotoStatement, 0, len(items))
	for _, item := range items {
		payload, payloadErr := extractPayload(item)
		if payloadErr != nil {
			return nil, payloadErr
		}
		decodedPayload, decodeErr := decodePayload(payload)
		if decodeErr != nil {
			return nil, decodeErr
		}

		var statement inTotoStatement
		if err := json.Unmarshal(decodedPayload, &statement); err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}

	return statements, nil
}

func extractPayload(item json.RawMessage) (string, error) {
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(item, &envelope); err != nil {
		return "", err
	}

	for _, key := range []string{"payload", "Payload"} {
		if raw, ok := envelope[key]; ok {
			var payload string
			if err := json.Unmarshal(raw, &payload); err != nil {
				return "", err
			}
			return payload, nil
		}
	}

	return "", fmt.Errorf("attestation payload missing from cosign output")
}

func decodePayload(payload string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err == nil {
		return decoded, nil
	}
	return base64.RawStdEncoding.DecodeString(payload)
}

func matchStatementSubject(statements []inTotoStatement, image string) (string, string, error) {
	imageRepository := normalizeImageRepository(image)
	imageDigest := digestFromImage(image)
	for _, statement := range statements {
		for _, subject := range statement.Subject {
			subjectDigest := normalizeDigest(subject.Digest["sha256"])
			subjectRepository := normalizeImageRepository(subject.Name)
			if subjectDigest == imageDigest && subjectRepository == imageRepository {
				return subject.Name, subjectDigest, nil
			}
		}
	}

	return "", "", fmt.Errorf("attestation subjects do not match image repository %q and digest %q", imageRepository, imageDigest)
}

func parseWorkflowIdentity(identity string) (repository, workflowFile, ref string) {
	const prefix = "https://github.com/"
	if !strings.HasPrefix(identity, prefix) {
		return "", "", ""
	}

	withoutPrefix := strings.TrimPrefix(identity, prefix)
	parts := strings.SplitN(withoutPrefix, "/.github/workflows/", 2)
	if len(parts) != 2 {
		return "", "", ""
	}

	repository = strings.Trim(parts[0], "/")
	workflowAndRef := parts[1]
	refParts := strings.SplitN(workflowAndRef, "@", 2)
	workflowFile = ".github/workflows/" + refParts[0]
	if len(refParts) == 2 {
		ref = refParts[1]
	}
	return repository, workflowFile, ref
}

func normalizeImageRepository(image string) string {
	image = strings.TrimSpace(image)
	if image == "" {
		return ""
	}

	if parts := strings.SplitN(image, "@", 2); len(parts) == 2 {
		image = parts[0]
	}

	lastSlash := strings.LastIndex(image, "/")
	lastColon := strings.LastIndex(image, ":")
	if lastColon > lastSlash {
		image = image[:lastColon]
	}

	return image
}

func digestFromImage(image string) string {
	parts := strings.SplitN(strings.TrimSpace(image), "@", 2)
	if len(parts) != 2 {
		return ""
	}
	return normalizeDigest(parts[1])
}

func normalizeDigest(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if strings.HasPrefix(value, "sha256:") {
		return value
	}
	return "sha256:" + value
}
