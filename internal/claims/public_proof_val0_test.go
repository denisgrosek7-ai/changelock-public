package claims

import (
	"testing"

	"github.com/denisgrosek/changelock/internal/signing"
)

func TestMeasuredPublicProofVal0StateRequiresAllFoundationSlices(t *testing.T) {
	got := EvaluateMeasuredPublicProofVal0State(
		MeasuredPublicProofVal0ClaimRegistryStateActive,
		MeasuredPublicProofVal0RedactionTierStateActive,
		MeasuredPublicProofVal0SigningAuthorityStateActive,
		MeasuredPublicProofVal0CompatibilityStatePartial,
	)
	if got != MeasuredPublicProofVal0StateSubstantial {
		t.Fatalf("expected substantial val0 state, got %q", got)
	}
}

func TestMeasuredPublicProofVal0ClaimRegistryRequiresLifecycleAndBoundaryRules(t *testing.T) {
	model := MeasuredPublicProofVal0ClaimRegistryModel()
	model.LifecycleStates = nil
	if got := EvaluateMeasuredPublicProofVal0ClaimRegistryState(model); got != MeasuredPublicProofVal0ClaimRegistryStateIncomplete {
		t.Fatalf("expected incomplete claim registry without lifecycle states, got %q", got)
	}
}

func TestMeasuredPublicProofVal0SigningAuthorityRequiresTrustRootFields(t *testing.T) {
	model := MeasuredPublicProofVal0SigningAuthority(signing.ProviderDescriptor{
		ProviderMode:         signing.ModeSoftware,
		TrustBoundary:        signing.TrustBoundaryApplicationLocal,
		ActiveLifecycleState: signing.KeyStateActive,
	})
	model.TrustRoots[0].KeyVersion = ""
	if got := EvaluateMeasuredPublicProofVal0SigningAuthorityState(model); got != MeasuredPublicProofVal0SigningAuthorityStatePartial {
		t.Fatalf("expected partial signing authority without key version, got %q", got)
	}
}

func TestMeasuredPublicProofVal0FoundationIsActive(t *testing.T) {
	registry := MeasuredPublicProofVal0ClaimRegistryModel()
	if got := EvaluateMeasuredPublicProofVal0ClaimRegistryState(registry); got != MeasuredPublicProofVal0ClaimRegistryStateActive {
		t.Fatalf("expected active claim registry state, got %q", got)
	}

	redaction := MeasuredPublicProofVal0RedactionTiers()
	if got := EvaluateMeasuredPublicProofVal0RedactionTierState(redaction); got != MeasuredPublicProofVal0RedactionTierStateActive {
		t.Fatalf("expected active redaction tiers, got %q", got)
	}

	signingAuthority := MeasuredPublicProofVal0SigningAuthority(signing.ProviderDescriptor{
		ProviderMode:                   signing.ModeSoftware,
		TrustBoundary:                  signing.TrustBoundaryApplicationLocal,
		ActiveLifecycleState:           signing.KeyStateActive,
		SupportsRotation:               true,
		SupportsVerifyOnlyRetirement:   true,
		SupportsRevocation:             true,
		SupportsHistoricalVerification: true,
		KeyClasses:                     []string{signing.KeyClassSealing, signing.KeyClassVerificationRoot},
	})
	if got := EvaluateMeasuredPublicProofVal0SigningAuthorityState(signingAuthority); got != MeasuredPublicProofVal0SigningAuthorityStateActive {
		t.Fatalf("expected active signing authority state, got %q", got)
	}

	compatibility := MeasuredPublicProofVal0CompatibilityBaseline()
	if got := EvaluateMeasuredPublicProofVal0CompatibilityState(compatibility); got != MeasuredPublicProofVal0CompatibilityStateActive {
		t.Fatalf("expected active compatibility state, got %q", got)
	}

	if got := EvaluateMeasuredPublicProofVal0State(
		registry.CurrentState,
		EvaluateMeasuredPublicProofVal0RedactionTierState(redaction),
		signingAuthority.CurrentState,
		compatibility.CurrentState,
	); got != MeasuredPublicProofVal0StateActive {
		t.Fatalf("expected active val0 state, got %q", got)
	}
}
