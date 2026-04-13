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
	"github.com/denisgrosek/changelock/internal/verify"
)

var artifactVerifier verify.ArtifactVerifier = verify.NewCosignVerifier(envOrDefault("CHANGELOCK_COSIGN_BIN", "cosign"))
var auditWriter = audit.NewDefaultWriter()

func main() {
	addr := ":" + envOrDefault("PORT", "8091")
	log.Printf("attestation-verifier listening on %s", addr)
	log.Fatal((&http.Server{
		Addr:              addr,
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       2 * time.Minute,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       60 * time.Second,
	}).ListenAndServe())
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/verify/artifact", artifactHandler)
	return metrics.InstrumentHTTP("attestation-verifier", mux)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	httpjson.Write(w, http.StatusOK, map[string]string{"status": "ok"})
}

func artifactHandler(w http.ResponseWriter, r *http.Request) {
	requestID := requestIDFromHeader(r)

	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request verify.ArtifactVerificationRequest
	if err := httpjson.Decode(r, &request); err != nil {
		writeAuditEvent(r.Context(), audit.Event{
			RequestID: requestID,
			Component: "attestation-verifier",
			EventType: audit.EventTypeArtifactVerificationResult,
			Decision:  audit.DecisionDeny,
			Reasons:   []string{"invalid verification request: " + err.Error()},
		})
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Minute)
	defer cancel()

	result, err := artifactVerifier.VerifyArtifact(ctx, request)
	writeAuditEvent(ctx, buildVerificationAuditEvent(requestID, request, &result, err))
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, result)
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

func buildVerificationAuditEvent(requestID string, request verify.ArtifactVerificationRequest, result *verify.ArtifactVerification, err error) audit.Event {
	decision := audit.DecisionAllow
	reasons := []string{}
	if err != nil {
		decision = audit.DecisionError
		reasons = append(reasons, err.Error())
	} else if result == nil || !result.SignatureValid || !result.AttestationValid || len(result.Reasons) > 0 {
		decision = audit.DecisionDeny
	}
	if result != nil {
		reasons = append(reasons, result.Reasons...)
	}

	summary, evidence := audit.FromArtifactVerification(result)

	return audit.Event{
		RequestID:       requestID,
		Component:       "attestation-verifier",
		EventType:       audit.EventTypeArtifactVerificationResult,
		Repo:            firstNonEmpty(resultValue(result, func(v *verify.ArtifactVerification) string { return v.VerifiedRepo }), request.ExpectedRepository),
		Branch:          audit.BranchFromRef(firstNonEmpty(resultValue(result, func(v *verify.ArtifactVerification) string { return v.VerifiedRef }), request.ExpectedRef)),
		Image:           request.Image,
		Digest:          firstNonEmpty(resultValue(result, func(v *verify.ArtifactVerification) string { return v.VerifiedDigest }), audit.DigestFromImage(request.Image)),
		Decision:        decision,
		Reasons:         reasons,
		VerifierSummary: summary,
		Evidence:        evidence,
	}
}

func writeAuditEvent(ctx context.Context, event audit.Event) {
	metrics.IncDecision("attestation-verifier", event.Decision, event.EventType)
	if event.Decision == audit.DecisionAllow {
		metrics.IncArtifactVerificationSuccess("attestation-verifier")
	} else {
		metrics.IncArtifactVerificationFailure("attestation-verifier")
	}
	if err := auditWriter.Write(ctx, event); err != nil {
		log.Printf("attestation-verifier audit write failed: %v", err)
	}
}

func resultValue(result *verify.ArtifactVerification, getter func(*verify.ArtifactVerification) string) string {
	if result == nil {
		return ""
	}
	return getter(result)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
