package attestation

import (
	"strings"
	"time"
)

const (
	SchemaVersion = "2.remote_attestation_verifier.v1"

	VerdictVerified = "verified"
	VerdictDegraded = "degraded"
	VerdictMismatch = "mismatch"
	VerdictRejected = "rejected"
)

type ProviderAdapter struct {
	ProviderID                string   `json:"provider_id"`
	CapabilityClass           string   `json:"capability_class"`
	TrustBoundary             string   `json:"trust_boundary"`
	SupportedQuoteTypes       []string `json:"supported_quote_types,omitempty"`
	SupportedSubstrateClasses []string `json:"supported_substrate_classes,omitempty"`
	Limitations               []string `json:"limitations,omitempty"`
}

type VerificationRequest struct {
	SubjectRef               string    `json:"subject_ref,omitempty"`
	TenantID                 string    `json:"tenant_id,omitempty"`
	Environment              string    `json:"environment,omitempty"`
	Provider                 string    `json:"provider"`
	QuoteType                string    `json:"quote_type,omitempty"`
	Measurement              string    `json:"measurement,omitempty"`
	LifecycleState           string    `json:"lifecycle_state,omitempty"`
	NodeID                   string    `json:"node_id,omitempty"`
	SubstrateClass           string    `json:"substrate_class,omitempty"`
	TrustedMeasurements      []string  `json:"trusted_measurements,omitempty"`
	RequireCredentialRelease bool      `json:"require_credential_release,omitempty"`
	ValidUntil               time.Time `json:"valid_until,omitempty"`
}

type VerificationResult struct {
	SchemaVersion            string    `json:"schema_version"`
	SubjectRef               string    `json:"subject_ref,omitempty"`
	Provider                 string    `json:"provider"`
	QuoteType                string    `json:"quote_type,omitempty"`
	Measurement              string    `json:"measurement,omitempty"`
	LifecycleState           string    `json:"lifecycle_state,omitempty"`
	SubstrateClass           string    `json:"substrate_class,omitempty"`
	TrustBoundary            string    `json:"trust_boundary,omitempty"`
	CurrentState             string    `json:"current_state"`
	CredentialReleaseAllowed bool      `json:"credential_release_allowed"`
	Reasons                  []string  `json:"reasons,omitempty"`
	VerifiedAt               time.Time `json:"verified_at"`
}

type Verifier struct {
	adapters map[string]ProviderAdapter
	now      func() time.Time
}

func NewVerifier() Verifier {
	adapters := map[string]ProviderAdapter{}
	for _, item := range Catalog() {
		adapters[item.ProviderID] = item
	}
	return Verifier{
		adapters: adapters,
		now:      time.Now,
	}
}

func Catalog() []ProviderAdapter {
	return []ProviderAdapter{
		{
			ProviderID:                "sgx",
			CapabilityClass:           "enclave_attestation",
			TrustBoundary:             "attestation_provider_layer",
			SupportedQuoteTypes:       []string{"sgx_quote"},
			SupportedSubstrateClasses: []string{"confidential"},
			Limitations: []string{
				"Bounded verifier logic validates structured quote metadata and trusted measurements; it is not a substitute for full vendor SDK coverage.",
			},
		},
		{
			ProviderID:                "tdx",
			CapabilityClass:           "confidential_vm_attestation",
			TrustBoundary:             "attestation_provider_layer",
			SupportedQuoteTypes:       []string{"tdx_quote"},
			SupportedSubstrateClasses: []string{"confidential", "hardened"},
		},
		{
			ProviderID:                "sev",
			CapabilityClass:           "confidential_vm_attestation",
			TrustBoundary:             "attestation_provider_layer",
			SupportedQuoteTypes:       []string{"sev_quote", "snp_report"},
			SupportedSubstrateClasses: []string{"confidential", "hardened"},
		},
	}
}

func (v Verifier) Verify(request VerificationRequest) VerificationResult {
	if v.now == nil {
		v.now = time.Now
	}
	request.Provider = strings.ToLower(strings.TrimSpace(request.Provider))
	request.QuoteType = strings.ToLower(strings.TrimSpace(request.QuoteType))
	request.Measurement = strings.TrimSpace(request.Measurement)
	request.LifecycleState = strings.TrimSpace(request.LifecycleState)
	request.SubstrateClass = strings.ToLower(strings.TrimSpace(request.SubstrateClass))

	result := VerificationResult{
		SchemaVersion:  SchemaVersion,
		SubjectRef:     strings.TrimSpace(request.SubjectRef),
		Provider:       request.Provider,
		QuoteType:      request.QuoteType,
		Measurement:    request.Measurement,
		LifecycleState: request.LifecycleState,
		SubstrateClass: request.SubstrateClass,
		VerifiedAt:     v.now().UTC(),
		CurrentState:   VerdictRejected,
	}

	adapter, ok := v.adapters[request.Provider]
	if !ok {
		result.Reasons = []string{"provider_not_supported"}
		return result
	}
	result.TrustBoundary = adapter.TrustBoundary
	if request.QuoteType == "" || !contains(adapter.SupportedQuoteTypes, request.QuoteType) {
		result.Reasons = append(result.Reasons, "quote_type_not_supported")
		return result
	}
	if request.SubstrateClass != "" && !contains(adapter.SupportedSubstrateClasses, request.SubstrateClass) {
		result.CurrentState = VerdictMismatch
		result.Reasons = append(result.Reasons, "substrate_class_not_supported")
		return result
	}
	switch strings.ToLower(request.LifecycleState) {
	case "revoked", "destroyed":
		result.CurrentState = VerdictMismatch
		result.Reasons = append(result.Reasons, "lifecycle_not_trusted")
		return result
	case "rotate_pending", "provisioned", "":
		result.CurrentState = VerdictDegraded
		result.Reasons = append(result.Reasons, "lifecycle_not_fully_active")
	default:
		result.CurrentState = VerdictVerified
	}
	if !request.ValidUntil.IsZero() && request.ValidUntil.Before(result.VerifiedAt) {
		result.CurrentState = VerdictMismatch
		result.Reasons = append(result.Reasons, "attestation_expired")
		return result
	}
	if len(request.TrustedMeasurements) > 0 && !contains(request.TrustedMeasurements, request.Measurement) {
		result.CurrentState = VerdictMismatch
		result.Reasons = append(result.Reasons, "measurement_not_trusted")
		return result
	}
	if request.Measurement == "" {
		if request.RequireCredentialRelease {
			result.CurrentState = VerdictRejected
			result.Reasons = append(result.Reasons, "measurement_missing_for_credential_release")
			return result
		}
		result.CurrentState = VerdictDegraded
		result.Reasons = append(result.Reasons, "measurement_missing")
	}
	result.CredentialReleaseAllowed = result.CurrentState == VerdictVerified
	if len(result.Reasons) == 0 {
		result.Reasons = []string{"attestation_verified"}
	}
	return result
}

func contains(values []string, needle string) bool {
	needle = strings.TrimSpace(needle)
	for _, value := range values {
		if strings.TrimSpace(value) == needle {
			return true
		}
	}
	return false
}
