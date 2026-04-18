package runtime

import "testing"

func TestCompareNoDrift(t *testing.T) {
	result := Compare(sampleApprovedState(), sampleObservedState())

	if result.Result != string(DriftClassNoDrift) {
		t.Fatalf("expected no_drift, got %q", result.Result)
	}
	if result.HasDrift() {
		t.Fatalf("expected HasDrift() to be false")
	}
	if len(result.Reasons) != 0 {
		t.Fatalf("expected no reasons, got %#v", result.Reasons)
	}
}

func TestCompareDetectsImageDrift(t *testing.T) {
	observed := sampleObservedState()
	observed.Containers[0].RunningDigest = "sha256:mutated"

	result := Compare(sampleApprovedState(), observed)

	if result.Result != string(DriftClassImageDigest) {
		t.Fatalf("expected image_digest_drift, got %q", result.Result)
	}
	if len(result.Classes) != 1 || result.Classes[0] != string(DriftClassImageDigest) {
		t.Fatalf("unexpected classes %#v", result.Classes)
	}
	if result.Severity != DriftSeverityHigh {
		t.Fatalf("expected high severity, got %q", result.Severity)
	}
	if result.Evidence == nil || len(result.Evidence.ImageMismatches) != 1 {
		t.Fatalf("expected image mismatch evidence, got %#v", result.Evidence)
	}
}

func TestCompareDetectsConfigDrift(t *testing.T) {
	observed := sampleObservedState()
	observed.ActualConfigHash = "cfg-mutated"

	result := Compare(sampleApprovedState(), observed)

	if result.Result != string(DriftClassWorkloadSpec) {
		t.Fatalf("expected workload_spec_drift, got %q", result.Result)
	}
	if result.Evidence == nil || result.Evidence.ConfigObserved != "cfg-mutated" {
		t.Fatalf("expected config mismatch evidence, got %#v", result.Evidence)
	}
}

func TestCompareDetectsSecurityContextDrift(t *testing.T) {
	observed := sampleObservedState()
	observed.Containers[0].Runtime.AllowPrivilegeEscalation = true

	result := Compare(sampleApprovedState(), observed)

	if result.Result != string(DriftClassSecurityContext) {
		t.Fatalf("expected security_context_drift, got %q", result.Result)
	}
	if result.Evidence == nil || len(result.Evidence.SecurityContextMismatches) != 1 {
		t.Fatalf("expected security mismatch evidence, got %#v", result.Evidence)
	}
	if result.Severity != DriftSeverityCritical {
		t.Fatalf("expected critical severity, got %q", result.Severity)
	}
}

func TestCompareDetectsServiceAccountDrift(t *testing.T) {
	observed := sampleObservedState()
	observed.ServiceAccountName = "mutated-sa"

	result := Compare(sampleApprovedState(), observed)

	if result.Result != string(DriftClassServiceAccount) {
		t.Fatalf("expected service_account_drift, got %q", result.Result)
	}
	if result.Evidence == nil || result.Evidence.ServiceAccountObserved != "mutated-sa" {
		t.Fatalf("expected service account evidence, got %#v", result.Evidence)
	}
}

func TestCompareDetectsMultipleDriftClasses(t *testing.T) {
	observed := sampleObservedState()
	observed.Containers[0].RunningDigest = "sha256:mutated"
	observed.ActualConfigHash = "cfg-mutated"
	observed.Containers[0].Runtime.Privileged = true

	result := Compare(sampleApprovedState(), observed)

	if result.Result != string(DriftClassMultiple) {
		t.Fatalf("expected multiple_drift, got %q", result.Result)
	}
	expectedClasses := []string{string(DriftClassImageDigest), string(DriftClassSecurityContext), string(DriftClassWorkloadSpec)}
	if len(result.Classes) != len(expectedClasses) {
		t.Fatalf("expected classes %#v, got %#v", expectedClasses, result.Classes)
	}
	for idx, class := range expectedClasses {
		if result.Classes[idx] != class {
			t.Fatalf("expected class %q at index %d, got %#v", class, idx, result.Classes)
		}
	}
	if len(result.Reasons) < 3 {
		t.Fatalf("expected combined reasons, got %#v", result.Reasons)
	}
}

func sampleApprovedState() ApprovedWorkloadState {
	return ApprovedWorkloadState{
		Namespace:          "acme-prod",
		WorkloadKind:       "Deployment",
		Workload:           "booking-api",
		ServiceAccountName: "booking-api",
		ExpectedConfigHash: "cfg-123",
		Containers: []ApprovedContainerState{
			{
				Name:           "app",
				Image:          "ghcr.io/my-org/booking-api@sha256:abc123",
				ApprovedDigest: "sha256:abc123",
				Runtime: SecurityConstraints{
					RunAsNonRoot:             true,
					ReadOnlyRootFilesystem:   true,
					AllowPrivilegeEscalation: false,
					DropAllCapabilities:      true,
					SeccompRuntimeDefault:    true,
					DenyPrivileged:           true,
				},
			},
		},
	}
}

func sampleObservedState() ObservedWorkloadState {
	return ObservedWorkloadState{
		Namespace:        "acme-prod",
		WorkloadKind:     "Deployment",
		Workload:         "booking-api",
		ServiceAccountName: "booking-api",
		ActualConfigHash: "cfg-123",
		Containers: []ObservedContainerState{
			{
				Name:          "app",
				Image:         "ghcr.io/my-org/booking-api@sha256:abc123",
				RunningDigest: "sha256:abc123",
				Runtime: SecurityPosture{
					RunAsNonRoot:             true,
					ReadOnlyRootFilesystem:   true,
					AllowPrivilegeEscalation: false,
					DropAllCapabilities:      true,
					SeccompRuntimeDefault:    true,
					Privileged:               false,
				},
			},
		},
	}
}
