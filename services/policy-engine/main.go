package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
	"github.com/denisgrosek/changelock/internal/policy"
	"github.com/denisgrosek/changelock/internal/verify"
)

var auditWriter = audit.NewDefaultWriter()
var exceptionValidator audit.ExceptionValidator = newExceptionValidator()

func main() {
	if err := validateExceptionValidatorConfig(); err != nil {
		log.Fatal(err)
	}
	addr := ":" + envOrDefault("PORT", "8090")
	log.Printf("policy-engine listening on %s", addr)
	log.Fatal((&http.Server{
		Addr:              addr,
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}).ListenAndServe())
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/evaluate/change", evaluateChangeHandler)
	mux.HandleFunc("/evaluate/artifact", evaluateArtifactHandler)
	return metrics.InstrumentHTTP("policy-engine", mux)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	httpjson.Write(w, http.StatusOK, map[string]string{"status": "ok"})
}

func evaluateChangeHandler(w http.ResponseWriter, r *http.Request) {
	requestID := requestIDFromHeader(r)

	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request policy.ChangeEvaluationRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	bundle, err := policy.LoadBundle(policy.DefaultPoliciesDir(), request.Tenant)
	if err != nil {
		decision := policy.Decision{
			Decision: "DENY",
			Reasons:  []string{"policy bundle unavailable: " + err.Error()},
		}
		decision = policy.WithIdentity(nil, decision, policy.DecisionIdentityInput{
			RequestID: requestID,
			Component: "policy-engine",
			Repo:      request.Repository,
		})
		writeAuditEvent(r.Context(), audit.Event{
			RequestID:     requestID,
			Component:     "policy-engine",
			EventType:     audit.EventTypePolicyDecision,
			TenantID:      request.Tenant,
			Repo:          request.Repository,
			Branch:        request.Branch,
			Decision:      decision.Decision,
			Reasons:       decision.Reasons,
			PolicyVersion: "bundle-load-error",
			DecisionHash:  decision.DecisionHash,
		})
		httpjson.Write(w, http.StatusOK, decision)
		return
	}

	if decision, handled := maybeBypassChange(r.Context(), requestID, bundle, request); handled {
		httpjson.Write(w, http.StatusOK, decision)
		return
	}

	decision := policy.EvaluateChange(bundle, request)
	decision = policy.WithIdentity(bundle, decision, policy.DecisionIdentityInput{
		RequestID:   requestID,
		Component:   "policy-engine",
		Repo:        request.Repository,
		Environment: request.Environment,
	})
	writeAuditEvent(r.Context(), audit.Event{
		RequestID:                requestID,
		Component:                "policy-engine",
		EventType:                audit.EventTypePolicyDecision,
		TenantID:                 request.Tenant,
		Repo:                     request.Repository,
		Branch:                   request.Branch,
		Decision:                 decision.Decision,
		Reasons:                  decision.Reasons,
		PolicyVersion:            bundle.Change.Metadata.Name,
		PolicyBundleID:           decision.PolicyBundleID,
		PolicyBundleHash:         decision.PolicyBundleHash,
		DecisionHash:             decision.DecisionHash,
		IsException:              decision.IsException,
		ExceptionID:              decision.ExceptionID,
		ExceptionType:            decision.ExceptionType,
		ExceptionStatus:          decision.ExceptionStatus,
		ExceptionReason:          decision.ExceptionReason,
		ExceptionTicketID:        decision.ExceptionTicketID,
		ExceptionRequestedBy:     decision.ExceptionRequestedBy,
		ExceptionRequestedAt:     decision.ExceptionRequestedAt,
		ExceptionApprovedBy:      decision.ExceptionApprovedBy,
		ExceptionApprovedAt:      decision.ExceptionApprovedAt,
		ExceptionRejectedBy:      decision.ExceptionRejectedBy,
		ExceptionRejectedAt:      decision.ExceptionRejectedAt,
		ExceptionRejectionReason: decision.ExceptionRejectionReason,
		ExceptionExpiresAt:       decision.ExceptionExpiresAt,
	})
	httpjson.Write(w, http.StatusOK, decision)
}

func evaluateArtifactHandler(w http.ResponseWriter, r *http.Request) {
	requestID := requestIDFromHeader(r)

	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request policy.ArtifactEvaluationRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	bundle, err := policy.LoadBundle(policy.DefaultPoliciesDir(), request.Tenant)
	if err != nil {
		decision := policy.Decision{
			Decision: "DENY",
			Reasons:  []string{"policy bundle unavailable: " + err.Error()},
		}
		decision = policy.WithIdentity(nil, decision, policy.DecisionIdentityInput{
			RequestID:   requestID,
			ImageDigest: audit.DigestFromImage(request.Image),
			Component:   "policy-engine",
			Repo:        request.Repository,
		})
		summary, evidence := audit.FromArtifactVerification(request.Verification)
		writeAuditEvent(r.Context(), audit.Event{
			RequestID:       requestID,
			Component:       "policy-engine",
			EventType:       audit.EventTypePolicyDecision,
			TenantID:        request.Tenant,
			Repo:            request.Repository,
			Image:           request.Image,
			Digest:          audit.DigestFromImage(request.Image),
			Decision:        decision.Decision,
			Reasons:         decision.Reasons,
			VerifierSummary: summary,
			Evidence:        evidence,
			PolicyVersion:   "bundle-load-error",
			DecisionHash:    decision.DecisionHash,
		})
		httpjson.Write(w, http.StatusOK, decision)
		return
	}

	digest := firstNonEmpty(resultDigest(request.Verification), audit.DigestFromImage(request.Image))
	if decision, handled := maybeBypassArtifact(r.Context(), requestID, bundle, request, digest); handled {
		httpjson.Write(w, http.StatusOK, decision)
		return
	}

	decision := policy.EvaluateArtifact(bundle, request)
	decision = policy.WithIdentity(bundle, decision, policy.DecisionIdentityInput{
		RequestID:   requestID,
		ImageDigest: digest,
		Component:   "policy-engine",
		Repo:        request.Repository,
		Environment: request.Environment,
	})
	summary, evidence := audit.FromArtifactVerification(request.Verification)
	writeAuditEvent(r.Context(), audit.Event{
		RequestID:                requestID,
		Component:                "policy-engine",
		EventType:                audit.EventTypePolicyDecision,
		TenantID:                 request.Tenant,
		Repo:                     request.Repository,
		Branch:                   audit.BranchFromRef(firstNonEmpty(resultRef(request.Verification), "")),
		Image:                    request.Image,
		Digest:                   digest,
		Decision:                 decision.Decision,
		Reasons:                  decision.Reasons,
		VerifierSummary:          summary,
		Evidence:                 evidence,
		PolicyVersion:            bundle.Artifact.Metadata.Name,
		PolicyBundleID:           decision.PolicyBundleID,
		PolicyBundleHash:         decision.PolicyBundleHash,
		DecisionHash:             decision.DecisionHash,
		IsException:              decision.IsException,
		ExceptionID:              decision.ExceptionID,
		ExceptionType:            decision.ExceptionType,
		ExceptionStatus:          decision.ExceptionStatus,
		ExceptionReason:          decision.ExceptionReason,
		ExceptionTicketID:        decision.ExceptionTicketID,
		ExceptionRequestedBy:     decision.ExceptionRequestedBy,
		ExceptionRequestedAt:     decision.ExceptionRequestedAt,
		ExceptionApprovedBy:      decision.ExceptionApprovedBy,
		ExceptionApprovedAt:      decision.ExceptionApprovedAt,
		ExceptionRejectedBy:      decision.ExceptionRejectedBy,
		ExceptionRejectedAt:      decision.ExceptionRejectedAt,
		ExceptionRejectionReason: decision.ExceptionRejectionReason,
		ExceptionExpiresAt:       decision.ExceptionExpiresAt,
	})
	httpjson.Write(w, http.StatusOK, decision)
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func requestIDFromHeader(r *http.Request) string {
	if requestID := r.Header.Get("X-Request-Id"); requestID != "" {
		return requestID
	}
	return audit.NewRequestID()
}

func writeAuditEvent(ctx context.Context, event audit.Event) {
	metrics.IncDecision("policy-engine", event.Decision, event.EventType)
	if err := auditWriter.Write(ctx, event); err != nil {
		log.Printf("policy-engine audit write failed: %v", err)
	}
}

func resultRef(result *verify.ArtifactVerification) string {
	if result == nil {
		return ""
	}
	return result.VerifiedRef
}

func resultDigest(result *verify.ArtifactVerification) string {
	if result == nil {
		return ""
	}
	return result.VerifiedDigest
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
