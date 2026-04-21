package runtime

import "testing"

func TestNormalizeSubstrateTruthRecordVerified(t *testing.T) {
	record := NormalizeSubstrateTruthRecord(SubstrateTruthRecord{
		SubjectRef: "cluster-a/prod/Deployment/api",
		Workload: WorkloadIdentity{
			ClusterID:    "cluster-a",
			Namespace:    "acme-prod",
			WorkloadKind: "Deployment",
			Workload:     "api",
			ImageDigest:  "sha256:abc",
		},
		Process: ProcessIdentity{
			ProcessName: "api",
			CgroupID:    "cg-1",
		},
		Node: NodeIdentity{
			NodeID:         "node-a",
			SubstrateClass: SubstrateClassConfidential,
			TrustBoundary:  TrustBoundaryAttestationProvider,
		},
		Attestation: AttestationBinding{
			Provider:               "sgx",
			QuoteType:              "sgx_quote",
			Measurement:            "m-1",
			LifecycleState:         "active",
			ObservedState:          AttestationStateVerified,
			CredentialReleaseState: CredentialReleaseReleased,
		},
	}, nil)

	if record.CurrentState != SubstrateTruthStateBound {
		t.Fatalf("expected bound truth, got %#v", record)
	}
	if record.ConfidenceScore < 80 {
		t.Fatalf("expected strong confidence, got %#v", record)
	}
}

func TestMatchExecutionProfileMismatchDeny(t *testing.T) {
	profile, ok := ExecutionProfileByID("confidential-strict")
	if !ok {
		t.Fatal("expected profile")
	}
	profile.RequiredMeasurements = []string{"m-expected"}
	truth := NormalizeSubstrateTruthRecord(SubstrateTruthRecord{
		SubjectRef: "cluster-a/acme-prod/Deployment/api",
		Workload: WorkloadIdentity{
			ClusterID:    "cluster-a",
			Namespace:    "acme-prod",
			WorkloadKind: "Deployment",
			Workload:     "api",
			ImageDigest:  "sha256:abc",
		},
		Node: NodeIdentity{
			NodeID:         "node-a",
			SubstrateClass: SubstrateClassConfidential,
			TrustBoundary:  TrustBoundaryAttestationProvider,
		},
		Attestation: AttestationBinding{
			Provider:               "sgx",
			Measurement:            "m-observed",
			ObservedState:          AttestationStateVerified,
			CredentialReleaseState: CredentialReleaseReleased,
		},
	}, nil)

	match := MatchExecutionProfile(profile, truth)
	if match.Allowed || match.CurrentState != ProfileMatchStateMismatchDeny {
		t.Fatalf("expected mismatch deny, got %#v", match)
	}
}

func TestMatchExecutionProfileDegradedAcceptable(t *testing.T) {
	profile, ok := ExecutionProfileByID("hardened-node")
	if !ok {
		t.Fatal("expected profile")
	}
	truth := NormalizeSubstrateTruthRecord(SubstrateTruthRecord{
		SubjectRef: "cluster-a/acme-prod/Deployment/api",
		Workload: WorkloadIdentity{
			ClusterID:    "cluster-a",
			Namespace:    "acme-prod",
			WorkloadKind: "Deployment",
			Workload:     "api",
			ImageDigest:  "sha256:abc",
		},
		Node: NodeIdentity{
			NodeID:         "node-a",
			SubstrateClass: SubstrateClassHardened,
			TrustBoundary:  TrustBoundaryNodeLayer,
		},
		Attestation: AttestationBinding{
			Provider:      "tdx",
			ObservedState: AttestationStateDegraded,
		},
	}, nil)

	match := MatchExecutionProfile(profile, truth)
	if !match.Allowed || match.CurrentState != ProfileMatchStateDegradedAcceptable {
		t.Fatalf("expected degraded acceptable, got %#v", match)
	}
}
