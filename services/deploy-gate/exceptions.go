package main

import (
	"context"
	"os"
	"strings"
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
	return audit.NewHTTPExceptionClient(baseURL, 2*time.Second)
}

func maybeBypassAdmission(
	ctx context.Context,
	requestID string,
	bundle *policy.Bundle,
	request admissionRequest,
	tenant string,
	repo string,
	branch string,
	environment string,
	actor string,
	image string,
	digest string,
) (admissionResponse, bool) {
	intent := exceptionIntentFromAnnotations(request.Object.Metadata.Annotations)
	if intent == nil || !intent.Requested() {
		return admissionResponse{}, false
	}

	identityInput := policy.DecisionIdentityInput{
		RequestID:   requestID,
		ImageDigest: digest,
		Component:   "deploy-gate",
		Repo:        repo,
		Environment: environment,
	}
	validationRequest := audit.ExceptionValidationRequest{
		ExceptionID:   intent.ExceptionID,
		ExceptionType: intent.ExceptionType,
		TenantID:      tenant,
		Environment:   environment,
		Namespace:     request.Namespace,
		Repo:          repo,
		ImageDigest:   digest,
		CVEID:         intent.CVEID,
	}

	if !intent.BreakGlass {
		decision := exceptionFailureDecision(bundle, intent, identityInput, "break-glass annotation must be true when exception metadata is present")
		writeDeployGateExceptionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, decision, audit.EventTypeExceptionValidationFailed)
		writeDeployGateDecisionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, bundle.Runtime.Metadata.Name, decision)
		return deny(requestID, decision.Reasons[0]), true
	}
	if exceptionValidator == nil {
		decision := exceptionFailureDecision(bundle, intent, identityInput, "exception validation is unavailable")
		writeDeployGateExceptionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, decision, audit.EventTypeExceptionValidationFailed)
		writeDeployGateDecisionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, bundle.Runtime.Metadata.Name, decision)
		return deny(requestID, decision.Reasons[0]), true
	}

	result, err := exceptionValidator.Validate(ctx, validationRequest)
	if err != nil {
		decision := exceptionFailureDecision(bundle, intent, identityInput, "exception validation failed: "+err.Error())
		writeDeployGateExceptionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, decision, audit.EventTypeExceptionValidationFailed)
		writeDeployGateDecisionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, bundle.Runtime.Metadata.Name, decision)
		return deny(requestID, decision.Reasons[0]), true
	}
	if !result.Valid || result.Exception == nil {
		decision := exceptionFailureDecision(bundle, intent, identityInput, firstNonEmpty(result.Reason, "exception validation failed"))
		writeDeployGateExceptionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, decision, audit.EventTypeExceptionValidationFailed)
		writeDeployGateDecisionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, bundle.Runtime.Metadata.Name, decision)
		return deny(requestID, decision.Reasons[0]), true
	}

	decision := exceptionSuccessDecision(bundle, identityInput, *result.Exception)
	writeDeployGateExceptionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, decision, audit.EventTypeExceptionUsed)
	writeDeployGateDecisionAudit(ctx, requestID, request, tenant, repo, branch, environment, actor, image, digest, bundle.Runtime.Metadata.Name, decision)
	return admissionResponse{UID: requestID, Allowed: true}, true
}

func exceptionIntentFromAnnotations(annotations map[string]string) *policy.ExceptionIntent {
	intent := &policy.ExceptionIntent{
		BreakGlass:  strings.EqualFold(strings.TrimSpace(annotations["changelock.io/break-glass"]), "true"),
		ExceptionID: strings.TrimSpace(annotations["changelock.io/exception-id"]),
		Reason:      strings.TrimSpace(annotations["changelock.io/reason"]),
		TicketID:    strings.TrimSpace(annotations["changelock.io/ticket-id"]),
	}
	if !intent.Requested() {
		return nil
	}
	return intent
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

func writeDeployGateExceptionAudit(ctx context.Context, requestID string, request admissionRequest, tenant, repo, branch, environment, actor, image, digest string, decision policy.Decision, eventType string) {
	writeAuditEvent(ctx, audit.Event{
		RequestID:           requestID,
		Component:           "deploy-gate",
		EventType:           eventType,
		Actor:               actor,
		TenantID:            tenant,
		Repo:                repo,
		Branch:              branch,
		Environment:         environment,
		Namespace:           request.Namespace,
		Workload:            request.Object.Metadata.Name,
		Image:               image,
		Digest:              digest,
		Decision:            decision.Decision,
		Reasons:             decision.Reasons,
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

func writeDeployGateDecisionAudit(ctx context.Context, requestID string, request admissionRequest, tenant, repo, branch, environment, actor, image, digest, policyVersion string, decision policy.Decision) {
	writeAuditEvent(ctx, audit.Event{
		RequestID:           requestID,
		Component:           "deploy-gate",
		EventType:           audit.EventTypePolicyDecision,
		Actor:               actor,
		TenantID:            tenant,
		Repo:                repo,
		Branch:              branch,
		Environment:         environment,
		Namespace:           request.Namespace,
		Workload:            request.Object.Metadata.Name,
		Image:               image,
		Digest:              digest,
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
	writeAuditEvent(ctx, audit.Event{
		RequestID:           requestID,
		Component:           "deploy-gate",
		EventType:           audit.EventTypeDeployGateDecision,
		Actor:               actor,
		TenantID:            tenant,
		Repo:                repo,
		Branch:              branch,
		Environment:         environment,
		Namespace:           request.Namespace,
		Workload:            request.Object.Metadata.Name,
		Image:               image,
		Digest:              digest,
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
