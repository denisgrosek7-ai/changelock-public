package main

import (
	"context"
	"os"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/policy"
)

func newExceptionValidator() audit.ExceptionValidator {
	baseURL := firstNonEmpty(
		os.Getenv("CHANGELOCK_EXCEPTIONS_URL"),
		os.Getenv("AUDIT_WRITER_URL"),
		os.Getenv("CHANGELOCK_AUDIT_WRITER_URL"),
	)
	return audit.NewHTTPExceptionClient(baseURL, 2*time.Second, os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN"))
}

func maybeBypassChange(ctx context.Context, requestID string, bundle *policy.Bundle, request policy.ChangeEvaluationRequest) (policy.Decision, bool) {
	if request.Exception == nil || !request.Exception.Requested() {
		return policy.Decision{}, false
	}

	input := policy.DecisionIdentityInput{
		RequestID:   requestID,
		Component:   "policy-engine",
		Repo:        request.Repository,
		Environment: request.Environment,
	}
	validationRequest := audit.ExceptionValidationRequest{
		ExceptionID:   request.Exception.ExceptionID,
		ExceptionType: request.Exception.ExceptionType,
		TenantID:      request.Tenant,
		Environment:   request.Environment,
		Namespace:     request.Namespace,
		Repo:          request.Repository,
		CVEID:         request.Exception.CVEID,
	}

	return validateAndMaybeBypass(ctx, requestID, bundle, request.Tenant, request.Branch, request.Repository, request.Environment, request.Namespace, "", request.Exception, input, bundle.Change.Metadata.Name, validationRequest)
}

func maybeBypassArtifact(ctx context.Context, requestID string, bundle *policy.Bundle, request policy.ArtifactEvaluationRequest, digest string) (policy.Decision, bool) {
	if request.Exception == nil || !request.Exception.Requested() {
		return policy.Decision{}, false
	}

	input := policy.DecisionIdentityInput{
		RequestID:   requestID,
		ImageDigest: digest,
		Component:   "policy-engine",
		Repo:        request.Repository,
		Environment: request.Environment,
	}
	validationRequest := audit.ExceptionValidationRequest{
		ExceptionID:   request.Exception.ExceptionID,
		ExceptionType: request.Exception.ExceptionType,
		TenantID:      request.Tenant,
		Environment:   request.Environment,
		Namespace:     request.Namespace,
		Repo:          request.Repository,
		ImageDigest:   digest,
		CVEID:         request.Exception.CVEID,
	}

	return validateAndMaybeBypass(ctx, requestID, bundle, request.Tenant, "", request.Repository, request.Environment, request.Namespace, request.Image, request.Exception, input, bundle.Artifact.Metadata.Name, validationRequest)
}

func validateAndMaybeBypass(
	ctx context.Context,
	requestID string,
	bundle *policy.Bundle,
	tenant string,
	branch string,
	repo string,
	environment string,
	namespace string,
	image string,
	intent *policy.ExceptionIntent,
	identityInput policy.DecisionIdentityInput,
	policyVersion string,
	validationRequest audit.ExceptionValidationRequest,
) (policy.Decision, bool) {
	if intent == nil || !intent.Requested() {
		return policy.Decision{}, false
	}

	if !intent.BreakGlass {
		decision := exceptionFailureDecision(bundle, intent, identityInput, "break_glass must be true when exception intent is present")
		writePolicyExceptionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent, decision, audit.EventTypeExceptionValidationFailed)
		writePolicyDecisionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent.CVEID, policyVersion, decision)
		return decision, true
	}
	if exceptionValidator == nil {
		decision := exceptionFailureDecision(bundle, intent, identityInput, "exception validation is unavailable")
		writePolicyExceptionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent, decision, audit.EventTypeExceptionValidationFailed)
		writePolicyDecisionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent.CVEID, policyVersion, decision)
		return decision, true
	}

	result, err := exceptionValidator.Validate(ctx, validationRequest)
	if err != nil {
		decision := exceptionFailureDecision(bundle, intent, identityInput, "exception validation failed: "+err.Error())
		writePolicyExceptionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent, decision, audit.EventTypeExceptionValidationFailed)
		writePolicyDecisionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent.CVEID, policyVersion, decision)
		return decision, true
	}
	if !result.Valid || result.Exception == nil {
		decision := exceptionFailureDecision(bundle, intent, identityInput, firstNonEmpty(result.Reason, "exception validation failed"))
		writePolicyExceptionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent, decision, audit.EventTypeExceptionValidationFailed)
		writePolicyDecisionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent.CVEID, policyVersion, decision)
		return decision, true
	}

	decision := exceptionSuccessDecision(bundle, identityInput, *result.Exception)
	writePolicyExceptionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent, decision, audit.EventTypeExceptionUsed)
	writePolicyDecisionAudit(ctx, requestID, tenant, branch, repo, environment, namespace, image, identityInput.ImageDigest, intent.CVEID, policyVersion, decision)
	return decision, true
}

func exceptionSuccessDecision(bundle *policy.Bundle, input policy.DecisionIdentityInput, exception audit.PolicyException) policy.Decision {
	decision := policy.Decision{
		Decision:            audit.DecisionAllow,
		Reasons:             []string{"exception " + exception.ExceptionID + " approved temporary bypass"},
		IsException:         true,
		ExceptionID:         exception.ExceptionID,
		ExceptionType:       exception.ExceptionType,
		ExceptionReason:     exception.Reason,
		ExceptionTicketID:   exception.TicketID,
		ExceptionApprovedBy: exception.ApprovedBy,
		ExceptionExpiresAt:  &exception.ExpiresAt,
	}
	return policy.WithIdentity(bundle, decision, input)
}

func exceptionFailureDecision(bundle *policy.Bundle, intent *policy.ExceptionIntent, input policy.DecisionIdentityInput, reason string) policy.Decision {
	decision := policy.Decision{
		Decision:          audit.DecisionDeny,
		Reasons:           []string{reason},
		ExceptionID:       firstNonEmpty(intent.ExceptionID, ""),
		ExceptionType:     firstNonEmpty(intent.ExceptionType, ""),
		ExceptionReason:   firstNonEmpty(intent.Reason, ""),
		ExceptionTicketID: firstNonEmpty(intent.TicketID, ""),
	}
	return policy.WithIdentity(bundle, decision, input)
}

func writePolicyExceptionAudit(ctx context.Context, requestID, tenant, branch, repo, environment, namespace, image, digest string, intent *policy.ExceptionIntent, decision policy.Decision, eventType string) {
	writeAuditEvent(ctx, audit.Event{
		RequestID:           requestID,
		Component:           "policy-engine",
		EventType:           eventType,
		TenantID:            tenant,
		Repo:                repo,
		Branch:              branch,
		Environment:         environment,
		Namespace:           namespace,
		Image:               image,
		Digest:              digest,
		CVEID:               intent.CVEID,
		Decision:            decision.Decision,
		Reasons:             decision.Reasons,
		PolicyBundleID:      decision.PolicyBundleID,
		PolicyBundleHash:    decision.PolicyBundleHash,
		DecisionHash:        decision.DecisionHash,
		IsException:         decision.IsException,
		ExceptionID:         firstNonEmpty(decision.ExceptionID, intent.ExceptionID),
		ExceptionType:       firstNonEmpty(decision.ExceptionType, intent.ExceptionType),
		ExceptionReason:     firstNonEmpty(decision.ExceptionReason, intent.Reason),
		ExceptionTicketID:   firstNonEmpty(decision.ExceptionTicketID, intent.TicketID),
		ExceptionApprovedBy: decision.ExceptionApprovedBy,
		ExceptionExpiresAt:  decision.ExceptionExpiresAt,
	})
}

func writePolicyDecisionAudit(ctx context.Context, requestID, tenant, branch, repo, environment, namespace, image, digest, cveID, policyVersion string, decision policy.Decision) {
	writeAuditEvent(ctx, audit.Event{
		RequestID:           requestID,
		Component:           "policy-engine",
		EventType:           audit.EventTypePolicyDecision,
		TenantID:            tenant,
		Repo:                repo,
		Branch:              branch,
		Environment:         environment,
		Namespace:           namespace,
		Image:               image,
		Digest:              digest,
		CVEID:               cveID,
		Decision:            decision.Decision,
		Reasons:             decision.Reasons,
		PolicyVersion:       policyVersion,
		PolicyBundleID:      decision.PolicyBundleID,
		PolicyBundleHash:    decision.PolicyBundleHash,
		DecisionHash:        decision.DecisionHash,
		IsException:         decision.IsException,
		ExceptionID:         decision.ExceptionID,
		ExceptionType:       decision.ExceptionType,
		ExceptionReason:     decision.ExceptionReason,
		ExceptionTicketID:   decision.ExceptionTicketID,
		ExceptionApprovedBy: decision.ExceptionApprovedBy,
		ExceptionExpiresAt:  decision.ExceptionExpiresAt,
	})
}
