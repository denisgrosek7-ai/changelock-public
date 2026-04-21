package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/evidence"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
	"github.com/denisgrosek/changelock/internal/policy"
	"github.com/denisgrosek/changelock/internal/signingidentity"
	"github.com/denisgrosek/changelock/internal/verify"
)

var allowedOIDCIssuers = []string{
	"https://token.actions.githubusercontent.com",
}

var artifactVerifier verify.ArtifactVerifier = newArtifactVerifier()
var auditWriter = audit.NewDefaultWriter()
var exceptionValidator audit.ExceptionValidator = newExceptionValidator()
var vulnerabilityEvaluator vulnerabilityNetEvaluator = newVulnerabilityNetEvaluator()
var signerIdentityEnforcer signerIdentityEvaluator = newSignerIdentityEvaluator()

type admissionReview struct {
	APIVersion string             `json:"apiVersion,omitempty"`
	Kind       string             `json:"kind,omitempty"`
	Request    *admissionRequest  `json:"request,omitempty"`
	Response   *admissionResponse `json:"response,omitempty"`
}

type admissionRequest struct {
	UID       string          `json:"uid"`
	Namespace string          `json:"namespace,omitempty"`
	Kind      objectReference `json:"kind,omitempty"`
	Object    pod             `json:"object"`
}

type objectReference struct {
	Kind string `json:"kind,omitempty"`
}

type admissionResponse struct {
	UID     string          `json:"uid"`
	Allowed bool            `json:"allowed"`
	Status  *statusEnvelope `json:"status,omitempty"`
}

type statusEnvelope struct {
	Message string `json:"message,omitempty"`
}

type pod struct {
	Metadata objectMeta `json:"metadata"`
	Spec     podSpec    `json:"spec"`
}

type objectMeta struct {
	Name        string            `json:"name,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type podSpec struct {
	Containers      []container         `json:"containers,omitempty"`
	InitContainers  []container         `json:"initContainers,omitempty"`
	SecurityContext *podSecurityContext `json:"securityContext,omitempty"`
	HostNetwork     bool                `json:"hostNetwork,omitempty"`
	HostPID         bool                `json:"hostPID,omitempty"`
	HostIPC         bool                `json:"hostIPC,omitempty"`
}

type container struct {
	Name            string           `json:"name,omitempty"`
	Image           string           `json:"image"`
	SecurityContext *securityContext `json:"securityContext,omitempty"`
}

type podSecurityContext struct {
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty"`
}

type securityContext struct {
	RunAsNonRoot             *bool         `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem   *bool         `json:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation *bool         `json:"allowPrivilegeEscalation,omitempty"`
	Capabilities             *capabilities `json:"capabilities,omitempty"`
}

type capabilities struct {
	Add []string `json:"add,omitempty"`
}

func main() {
	if err := validateExceptionValidatorConfig(); err != nil {
		log.Fatal(err)
	}
	if err := validateVulnerabilityNetEvaluatorConfig(); err != nil {
		log.Fatal(err)
	}
	if err := validateSignerIdentityEvaluatorConfig(); err != nil {
		log.Fatal(err)
	}
	addr := ":" + envOrDefault("PORT", "8092")
	log.Printf("deploy-gate listening on %s", addr)
	certFile := os.Getenv("TLS_CERT_FILE")
	keyFile := os.Getenv("TLS_KEY_FILE")
	server := &http.Server{
		Addr:              addr,
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	if certFile != "" && keyFile != "" {
		log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
	}
	log.Fatal(server.ListenAndServe())
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/admission/review", admissionReviewHandler)
	return metrics.InstrumentHTTP("deploy-gate", mux)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	httpjson.Write(w, http.StatusOK, map[string]string{"status": "ok"})
}

func admissionReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var review admissionReview
	if err := decodeAdmissionReview(r, &review); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if review.Request == nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "admission review request is required"})
		return
	}

	response := evaluateAdmission(*review.Request)
	httpjson.Write(w, http.StatusOK, admissionReview{
		APIVersion: "admission.k8s.io/v1",
		Kind:       "AdmissionReview",
		Response:   &response,
	})
}

func evaluateAdmission(request admissionRequest) admissionResponse {
	requestID := request.UID
	if requestID == "" {
		requestID = audit.NewRequestID()
	}
	if request.Kind.Kind != "" && request.Kind.Kind != "Pod" {
		return admissionResponse{UID: requestID, Allowed: true}
	}

	tenant := resolveTenant(request.Namespace, request.Object.Metadata.Annotations)
	bundle, err := policy.LoadBundle(policy.DefaultPoliciesDir(), tenant)
	if err != nil {
		response := deny(requestID, "policy bundle unavailable: "+err.Error())
		writeAuditEvent(context.Background(), audit.Event{
			RequestID:   requestID,
			Component:   "deploy-gate",
			EventType:   audit.EventTypeDeployGateDecision,
			TenantID:    tenant,
			Environment: audit.EnvironmentFromNamespace(request.Namespace),
			Namespace:   request.Namespace,
			Workload:    request.Object.Metadata.Name,
			Decision:    audit.DecisionDeny,
			Reasons:     []string{response.Status.Message},
		})
		return response
	}

	annotations := request.Object.Metadata.Annotations
	repository := resolveRepository(annotations)
	branch := audit.BranchFromRef(annotations["changelock.io/workflow-ref"])
	environment := firstNonEmpty(annotations["changelock.io/environment"], audit.EnvironmentFromNamespace(request.Namespace))
	actor := request.Object.Metadata.Annotations["changelock.io/actor"]
	containers := append([]container{}, append(request.Object.Spec.InitContainers, request.Object.Spec.Containers...)...)
	primaryImage := selectPrimaryImage(containers, "")
	primaryDigest := audit.DigestFromImage(primaryImage)
	if response, handled := maybeBypassAdmission(context.Background(), requestID, bundle, request, tenant, repository, branch, environment, actor, primaryImage, primaryDigest); handled {
		return response
	}

	reasons := []string{}
	reasons = append(reasons, evaluateRuntimePolicy(bundle.Runtime.Spec, request.Object)...)
	trustedExecution, trustedReasons := evaluateTrustedExecutionProfile(request, selectPrimaryImage(containers, primaryImage))
	reasons = append(reasons, trustedReasons...)
	_ = trustedExecution
	var primaryVerification *verify.ArtifactVerification
	for _, workloadContainer := range append([]container{}, append(request.Object.Spec.InitContainers, request.Object.Spec.Containers...)...) {
		verificationRequest, requestErr := buildVerificationRequest(bundle, annotations, workloadContainer.Image)
		if requestErr != nil {
			reasons = append(reasons, workloadContainer.Name+": "+requestErr.Error())
			continue
		}
		verification, verifyErr := artifactVerifier.VerifyArtifact(context.Background(), verificationRequest)
		if verifyErr != nil {
			reasons = append(reasons, workloadContainer.Name+": artifact verifier error: "+verifyErr.Error())
			continue
		}
		if primaryVerification == nil {
			verificationCopy := verification
			primaryVerification = &verificationCopy
			primaryImage = workloadContainer.Image
		}

		artifactRequest := buildArtifactRequest(tenant, annotations, workloadContainer.Image, verification)
		artifactDecision := policy.EvaluateArtifact(bundle, artifactRequest)
		artifactDecision = policy.WithIdentity(bundle, artifactDecision, policy.DecisionIdentityInput{
			RequestID:   requestID,
			ImageDigest: firstNonEmpty(verification.VerifiedDigest, audit.DigestFromImage(workloadContainer.Image)),
			Component:   "deploy-gate",
			Repo:        firstNonEmpty(verification.VerifiedRepo, repository),
			Environment: environment,
		})
		identityDecision := evaluateSignerIdentity(context.Background(), tenant, environment, verification)
		if identityDecision != nil && signerIdentityShouldBlock(identityDecision, signerIdentityEnforcer) {
			artifactDecision.Reasons = append(artifactDecision.Reasons, "signer identity authorization failed: "+firstNonEmpty(identityDecision.ReasonDetail, identityDecision.ReasonCode))
			artifactDecision.Decision = audit.DecisionDeny
		}
		for _, reason := range artifactDecision.Reasons {
			reasons = append(reasons, workloadContainer.Name+": "+reason)
		}
		if vulnerabilityEvaluator != nil && vulnerabilityEvaluator.Enabled() {
			threshold := ""
			if verification.Evidence.SupplyChain != nil {
				threshold = strings.ToUpper(strings.TrimSpace(verification.Evidence.SupplyChain.VulnerabilityScanSeverityThreshold))
			}
			if threshold != "" {
				netResponse, netErr := vulnerabilityEvaluator.NetVulnerabilities(
					context.Background(),
					tenant,
					environment,
					firstNonEmpty(verification.VerifiedDigest, audit.DigestFromImage(workloadContainer.Image)),
					threshold,
				)
				if netErr != nil {
					reasons = append(reasons, workloadContainer.Name+": vex-aware vulnerability evaluation failed: "+netErr.Error())
				} else if netResponse.ThresholdBreached {
					reasons = append(reasons, workloadContainer.Name+": net actionable vulnerabilities remain at or above "+threshold+" after VEX evaluation")
				}
			}
		}
		summary, evidence := audit.FromArtifactVerification(&verification)
		if evidence != nil && identityDecision != nil {
			evidence.SigningIdentity = buildSigningIdentityAuditEvidence(*identityDecision)
		}
		writeAuditEvent(context.Background(), audit.Event{
			RequestID:        requestID,
			Component:        "deploy-gate",
			EventType:        audit.EventTypePolicyDecision,
			Actor:            actor,
			TenantID:         tenant,
			Repo:             firstNonEmpty(verification.VerifiedRepo, repository),
			Branch:           firstNonEmpty(audit.BranchFromRef(verification.VerifiedRef), branch),
			Environment:      environment,
			Namespace:        request.Namespace,
			Workload:         request.Object.Metadata.Name,
			Image:            workloadContainer.Image,
			Digest:           firstNonEmpty(verification.VerifiedDigest, audit.DigestFromImage(workloadContainer.Image)),
			Decision:         artifactDecision.Decision,
			Reasons:          artifactDecision.Reasons,
			VerifierSummary:  summary,
			Evidence:         evidence,
			PolicyVersion:    bundle.Artifact.Metadata.Name,
			PolicyBundleID:   artifactDecision.PolicyBundleID,
			PolicyBundleHash: artifactDecision.PolicyBundleHash,
			DecisionHash:     artifactDecision.DecisionHash,
		})
	}

	decision := audit.DecisionAllow
	if len(reasons) > 0 {
		decision = audit.DecisionDeny
	}
	finalDecision := policy.WithIdentity(bundle, policy.Decision{
		Decision: decision,
		Reasons:  reasons,
	}, policy.DecisionIdentityInput{
		RequestID:   requestID,
		ImageDigest: firstNonEmpty(resultDigest(primaryVerification), audit.DigestFromImage(selectPrimaryImage(containers, primaryImage))),
		Component:   "deploy-gate",
		Repo:        firstNonEmpty(resultRepo(primaryVerification), repository),
		Environment: environment,
	})
	summary, evidence := audit.FromArtifactVerification(primaryVerification)
	writeAuditEvent(context.Background(), audit.Event{
		RequestID:        requestID,
		Component:        "deploy-gate",
		EventType:        audit.EventTypeDeployGateDecision,
		Actor:            actor,
		TenantID:         tenant,
		Repo:             firstNonEmpty(resultRepo(primaryVerification), repository),
		Branch:           firstNonEmpty(audit.BranchFromRef(resultRef(primaryVerification)), branch),
		Environment:      environment,
		Namespace:        request.Namespace,
		Workload:         request.Object.Metadata.Name,
		Image:            selectPrimaryImage(containers, primaryImage),
		Digest:           firstNonEmpty(resultDigest(primaryVerification), audit.DigestFromImage(selectPrimaryImage(containers, primaryImage))),
		Decision:         finalDecision.Decision,
		Reasons:          finalDecision.Reasons,
		VerifierSummary:  summary,
		Evidence:         evidence,
		PolicyVersion:    bundle.Runtime.Metadata.Name,
		PolicyBundleID:   finalDecision.PolicyBundleID,
		PolicyBundleHash: finalDecision.PolicyBundleHash,
		DecisionHash:     finalDecision.DecisionHash,
	})

	if len(reasons) > 0 {
		return deny(requestID, strings.Join(reasons, "; "))
	}

	return admissionResponse{
		UID:     requestID,
		Allowed: true,
	}
}

func evaluateSignerIdentity(ctx context.Context, tenant string, environment string, verification verify.ArtifactVerification) *signingidentity.Decision {
	if signerIdentityEnforcer == nil || !signerIdentityEnforcer.Enabled() {
		return nil
	}
	request := signingIdentityEvaluateRequest{
		Issuer:            verification.VerifiedIssuer,
		SignerIdentity:    verification.VerifiedIdentity,
		Subject:           verification.VerifiedSubject,
		Repository:        verification.VerifiedRepo,
		Workflow:          verification.VerifiedWorkflow,
		Ref:               verification.VerifiedRef,
		TenantID:          tenant,
		Environment:       environment,
		TransparencyState: verification.Evidence.TransparencyLogState,
	}
	if verification.Evidence.Bundle != nil {
		request.EvidenceAt = firstNonNilTime(verification.Evidence.Bundle.IntegratedTime, verification.Evidence.Bundle.SignedAt)
	}
	decision, err := signerIdentityEnforcer.Evaluate(ctx, request)
	if err != nil {
		failed := signingidentity.Decision{
			Authorized:        signingidentity.AuthorizationUnknown,
			EnforcementMode:   signerIdentityEnforcer.Mode(),
			ReasonCode:        signingidentity.ReasonUnknown,
			ReasonDetail:      "signing identity evaluation failed: " + err.Error(),
			Deny:              signerIdentityEnforcer.Mode() == signingidentity.EnforcementEnforce,
			TransparencyState: verification.Evidence.TransparencyLogState,
		}
		return &failed
	}
	return &decision
}

func signerIdentityShouldBlock(decision *signingidentity.Decision, evaluator signerIdentityEvaluator) bool {
	if decision == nil || evaluator == nil {
		return false
	}
	return evaluator.Mode() == signingidentity.EnforcementEnforce && decision.Deny
}

func buildSigningIdentityAuditEvidence(decision signingidentity.Decision) *audit.SigningIdentityEvidence {
	return &audit.SigningIdentityEvidence{
		PolicyID:             decision.MatchedPolicyID,
		PolicyName:           decision.MatchedPolicyName,
		EnforcementMode:      decision.EnforcementMode,
		Authorized:           decision.Authorized,
		ReasonCode:           decision.ReasonCode,
		ReasonDetail:         decision.ReasonDetail,
		DistrustedAfter:      decision.DistrustedAfter,
		TransparencyRequired: decision.TransparencyRequired,
		TransparencyState:    decision.TransparencyState,
	}
}

func firstNonNilTime(values ...*time.Time) *time.Time {
	for _, value := range values {
		if value != nil && !value.IsZero() {
			timestamp := value.UTC()
			return &timestamp
		}
	}
	return nil
}

func buildArtifactRequest(tenant string, annotations map[string]string, image string, verification verify.ArtifactVerification) policy.ArtifactEvaluationRequest {
	repository := annotations["changelock.io/repository"]
	subject := annotations["changelock.io/subject"]
	if repository == "" && strings.HasPrefix(subject, "repo:") {
		repository = strings.TrimPrefix(subject, "repo:")
	}
	if subject == "" && repository != "" {
		subject = "repo:" + repository
	}

	return policy.ArtifactEvaluationRequest{
		Tenant:         tenant,
		Repository:     repository,
		Image:          image,
		DigestPinned:   strings.Contains(image, "@sha256:"),
		HasProvenance:  verification.AttestationValid,
		HasSignature:   verification.SignatureValid,
		SignerIdentity: verification.VerifiedIdentity,
		WorkflowFile:   verification.VerifiedWorkflow,
		Subject:        subject,
		Verification:   &verification,
	}
}

func buildVerificationRequest(bundle *policy.Bundle, annotations map[string]string, image string) (verify.ArtifactVerificationRequest, error) {
	repository := annotations["changelock.io/repository"]
	subject := annotations["changelock.io/subject"]
	if repository == "" && strings.HasPrefix(subject, "repo:") {
		repository = strings.TrimPrefix(subject, "repo:")
	}

	evidenceBundle, err := parseEvidenceBundleAnnotation(annotations)
	if err != nil {
		return verify.ArtifactVerificationRequest{}, err
	}

	return verify.ArtifactVerificationRequest{
		Image:                   image,
		ExpectedRepository:      repository,
		ExpectedRef:             annotations["changelock.io/workflow-ref"],
		ExpectedCommitSHA:       annotations["changelock.io/workflow-sha"],
		AllowedSignerIdentities: bundle.Artifact.Spec.AllowedSignerIdentities,
		AllowedOIDCIssuers:      allowedOIDCIssuers,
		EvidenceBundle:          evidenceBundle,
	}, nil
}

func evaluateRuntimePolicy(runtimePolicy policy.RuntimePolicySpec, workload pod) []string {
	reasons := []string{}
	if !runtimePolicy.AllowHostNetwork && workload.Spec.HostNetwork {
		reasons = append(reasons, "hostNetwork is not allowed")
	}
	if !runtimePolicy.AllowHostPID && workload.Spec.HostPID {
		reasons = append(reasons, "hostPID is not allowed")
	}
	if !runtimePolicy.AllowHostIPC && workload.Spec.HostIPC {
		reasons = append(reasons, "hostIPC is not allowed")
	}

	containers := append([]container{}, append(workload.Spec.InitContainers, workload.Spec.Containers...)...)
	for _, workloadContainer := range containers {
		if runtimePolicy.BlockLatestTag && strings.HasSuffix(workloadContainer.Image, ":latest") {
			reasons = append(reasons, workloadContainer.Name+": latest tag is blocked")
		}
		if runtimePolicy.RequireReadOnlyRootFilesystem {
			if workloadContainer.SecurityContext == nil || workloadContainer.SecurityContext.ReadOnlyRootFilesystem == nil || !*workloadContainer.SecurityContext.ReadOnlyRootFilesystem {
				reasons = append(reasons, workloadContainer.Name+": readOnlyRootFilesystem must be true")
			}
		}
		if !runtimePolicy.AllowPrivilegeEscalation {
			if workloadContainer.SecurityContext == nil || workloadContainer.SecurityContext.AllowPrivilegeEscalation == nil || *workloadContainer.SecurityContext.AllowPrivilegeEscalation {
				reasons = append(reasons, workloadContainer.Name+": allowPrivilegeEscalation must be false")
			}
		}
		if runtimePolicy.RequireNonRoot && !containerRunsAsNonRoot(workload.Spec.SecurityContext, workloadContainer.SecurityContext) {
			reasons = append(reasons, workloadContainer.Name+": runAsNonRoot must be true")
		}
		if len(runtimePolicy.MaxContainerCapabilities) == 0 && workloadContainer.SecurityContext != nil && workloadContainer.SecurityContext.Capabilities != nil && len(workloadContainer.SecurityContext.Capabilities.Add) > 0 {
			reasons = append(reasons, workloadContainer.Name+": additional Linux capabilities are not allowed")
		}
	}

	return reasons
}

func containerRunsAsNonRoot(podSecurity *podSecurityContext, containerSecurity *securityContext) bool {
	if containerSecurity != nil && containerSecurity.RunAsNonRoot != nil {
		return *containerSecurity.RunAsNonRoot
	}
	if podSecurity != nil && podSecurity.RunAsNonRoot != nil {
		return *podSecurity.RunAsNonRoot
	}
	return false
}

func resolveTenant(namespace string, annotations map[string]string) string {
	if tenant := annotations["changelock.io/tenant"]; tenant != "" {
		return tenant
	}
	if idx := strings.Index(namespace, "-"); idx > 0 {
		return namespace[:idx]
	}
	return "acme"
}

func resolveRepository(annotations map[string]string) string {
	repository := annotations["changelock.io/repository"]
	if repository == "" {
		subject := annotations["changelock.io/subject"]
		if strings.HasPrefix(subject, "repo:") {
			return strings.TrimPrefix(subject, "repo:")
		}
	}
	return repository
}

func resultRepo(result *verify.ArtifactVerification) string {
	if result == nil {
		return ""
	}
	return result.VerifiedRepo
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

func selectPrimaryImage(containers []container, current string) string {
	if current != "" {
		return current
	}
	if len(containers) == 1 {
		return containers[0].Image
	}
	return ""
}

func writeAuditEvent(ctx context.Context, event audit.Event) {
	metrics.IncDecision("deploy-gate", event.Decision, event.EventType)
	if err := auditWriter.Write(ctx, event); err != nil {
		log.Printf("deploy-gate audit write failed: %v", err)
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func deny(uid, message string) admissionResponse {
	return admissionResponse{
		UID:     uid,
		Allowed: false,
		Status: &statusEnvelope{
			Message: message,
		},
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func decodeAdmissionReview(r *http.Request, dst *admissionReview) error {
	if r.Body == nil {
		return io.EOF
	}
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	return decoder.Decode(dst)
}

func newArtifactVerifier() verify.ArtifactVerifier {
	evidenceConfig, err := evidence.LoadConfigFromEnv(os.Getenv)
	if err != nil {
		panic(err)
	}
	if fixturePath := os.Getenv("CHANGELOCK_VERIFIER_FIXTURE"); fixturePath != "" {
		verifier, err := verify.NewFixtureVerifier(fixturePath)
		if err != nil {
			log.Printf("deploy-gate fixture verifier unavailable, falling back to cosign: %v", err)
		} else {
			return verifier
		}
	}
	return verify.NewCosignVerifierWithEvidence(envOrDefault("CHANGELOCK_COSIGN_BIN", "cosign"), evidenceConfig)
}

func parseEvidenceBundleAnnotation(annotations map[string]string) (*evidence.Bundle, error) {
	raw := strings.TrimSpace(annotations["changelock.io/evidence-bundle"])
	if raw == "" {
		return nil, nil
	}

	var bundle evidence.Bundle
	if err := json.Unmarshal([]byte(raw), &bundle); err != nil {
		return nil, fmt.Errorf("invalid changelock.io/evidence-bundle annotation: %w", err)
	}
	return evidence.CloneBundle(&bundle), nil
}
