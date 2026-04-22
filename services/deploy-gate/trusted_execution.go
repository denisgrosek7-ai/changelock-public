package main

import (
	"fmt"
	"strings"
	"time"

	attestationruntime "github.com/denisgrosek/changelock/internal/attestation"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

type trustedExecutionEvaluation struct {
	Profile     runtimesubstrate.ExecutionProfile
	Truth       runtimesubstrate.SubstrateTruthRecord
	Attestation attestationruntime.VerificationResult
	Match       runtimesubstrate.ProfileMatch
}

func evaluateTrustedExecutionProfile(request admissionRequest, primaryImage string) (*trustedExecutionEvaluation, []string) {
	annotations := request.Object.Metadata.Annotations
	profileID := strings.TrimSpace(annotations["changelock.io/trusted-execution-profile"])
	if profileID == "" {
		return nil, nil
	}

	profile, ok := runtimesubstrate.ExecutionProfileByID(profileID)
	if !ok {
		return nil, []string{"trusted execution profile is unknown: " + profileID}
	}
	if measurement := strings.TrimSpace(annotations["changelock.io/trusted-measurement"]); measurement != "" && len(profile.RequiredMeasurements) == 0 {
		profile.RequiredMeasurements = []string{measurement}
	}
	validUntil, err := parseOptionalTime(annotations["changelock.io/attestation-valid-until"])
	if err != nil {
		return nil, []string{
			fmt.Sprintf("trusted execution profile %s rejected workload: attestation expiry annotation is invalid", profile.ProfileID),
		}
	}

	verifier := attestationruntime.NewVerifier()
	verification := verifier.Verify(attestationruntime.VerificationRequest{
		SubjectRef:               trustedExecutionSubjectRef(request),
		TenantID:                 resolveTenant(request.Namespace, annotations),
		Environment:              firstNonEmpty(annotations["changelock.io/environment"], auditEnvironment(request.Namespace)),
		Provider:                 strings.TrimSpace(annotations["changelock.io/attestation-provider"]),
		QuoteType:                defaultQuoteType(strings.TrimSpace(annotations["changelock.io/attestation-provider"]), annotations["changelock.io/attestation-quote-type"]),
		Measurement:              strings.TrimSpace(annotations["changelock.io/attestation-measurement"]),
		LifecycleState:           firstNonEmpty(annotations["changelock.io/attestation-lifecycle"], "active"),
		NodeID:                   strings.TrimSpace(annotations["changelock.io/node-id"]),
		SubstrateClass:           strings.TrimSpace(annotations["changelock.io/node-substrate-class"]),
		TrustedMeasurements:      profile.RequiredMeasurements,
		RequireCredentialRelease: profile.RequireCredentialRelease,
		ValidUntil:               validUntil,
	})

	truth := runtimesubstrate.NormalizeSubstrateTruthRecord(runtimesubstrate.SubstrateTruthRecord{
		SubjectRef: trustedExecutionSubjectRef(request),
		Workload: runtimesubstrate.WorkloadIdentity{
			ClusterID:     strings.TrimSpace(annotations["changelock.io/cluster-id"]),
			Namespace:     request.Namespace,
			WorkloadKind:  firstNonEmpty(request.Kind.Kind, "Pod"),
			Workload:      request.Object.Metadata.Name,
			ImageDigest:   auditDigest(primaryImage),
			PolicySubject: firstNonEmpty(annotations["changelock.io/subject"], annotations["changelock.io/repository"]),
		},
		Node: runtimesubstrate.NodeIdentity{
			NodeID:         strings.TrimSpace(annotations["changelock.io/node-id"]),
			SubstrateClass: strings.TrimSpace(annotations["changelock.io/node-substrate-class"]),
			TrustBoundary:  firstNonEmpty(annotations["changelock.io/node-trust-boundary"], runtimesubstrate.TrustBoundaryAttestationProvider),
			AttestationRef: strings.TrimSpace(annotations["changelock.io/node-attestation-ref"]),
		},
		Attestation: runtimesubstrate.AttestationBinding{
			Provider:               verification.Provider,
			QuoteType:              verification.QuoteType,
			Measurement:            verification.Measurement,
			LifecycleState:         verification.LifecycleState,
			ObservedState:          verification.CurrentState,
			CredentialReleaseState: credentialReleaseState(verification.CredentialReleaseAllowed),
			VerifiedAt:             verification.VerifiedAt,
		},
		ObservedAt: time.Now().UTC(),
	}, time.Now)
	match := runtimesubstrate.MatchExecutionProfile(profile, truth)

	evaluation := &trustedExecutionEvaluation{
		Profile:     profile,
		Truth:       truth,
		Attestation: verification,
		Match:       match,
	}
	if match.Allowed {
		return evaluation, nil
	}
	return evaluation, []string{
		fmt.Sprintf("trusted execution profile %s rejected workload: %s", profile.ProfileID, strings.Join(match.Reasons, ", ")),
	}
}

func trustedExecutionSubjectRef(request admissionRequest) string {
	clusterID := strings.TrimSpace(request.Object.Metadata.Annotations["changelock.io/cluster-id"])
	kind := firstNonEmpty(request.Kind.Kind, "Pod")
	return strings.Join([]string{
		firstNonEmpty(clusterID, "local"),
		request.Namespace,
		kind,
		request.Object.Metadata.Name,
	}, "/")
}

func auditEnvironment(namespace string) string {
	return strings.TrimSpace(strings.TrimPrefix(namespace, resolveTenant(namespace, map[string]string{})+"-"))
}

func defaultQuoteType(provider string, explicit string) string {
	if explicit = strings.TrimSpace(explicit); explicit != "" {
		return explicit
	}
	switch strings.ToLower(strings.TrimSpace(provider)) {
	case "sgx":
		return "sgx_quote"
	case "tdx":
		return "tdx_quote"
	case "sev":
		return "snp_report"
	default:
		return ""
	}
}

func auditDigest(image string) string {
	if idx := strings.Index(strings.TrimSpace(image), "@"); idx >= 0 && idx+1 < len(strings.TrimSpace(image)) {
		return strings.TrimSpace(image)[idx+1:]
	}
	return ""
}

func credentialReleaseState(allowed bool) string {
	if allowed {
		return runtimesubstrate.CredentialReleaseReleased
	}
	return runtimesubstrate.CredentialReleaseWithheld
}

func parseOptionalTime(raw string) (time.Time, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return time.Time{}, nil
	}
	if parsed, err := time.Parse(time.RFC3339, raw); err == nil {
		return parsed.UTC(), nil
	}
	return time.Time{}, fmt.Errorf("invalid RFC3339 timestamp")
}
