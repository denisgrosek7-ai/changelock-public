package runtime

import "testing"

func TestRuntimeSubstrateValBCorrelationModelIsActive(t *testing.T) {
	model := RuntimeSubstrateValBCorrelationModel()
	if got := EvaluateRuntimeSubstrateValBCorrelationModelState(model); got != RuntimeSubstrateValBCorrelationModelStateActive {
		t.Fatalf("expected active correlation model, got %q", got)
	}
}

func TestRuntimeSubstrateValBStateRequiresActiveValA(t *testing.T) {
	if got := EvaluateRuntimeSubstrateValBState(
		RuntimeSubstrateValAStateContractDefined,
		RuntimeSubstrateValBCorrelationModelStateActive,
		RuntimeSubstrateValBProcessImageStateActive,
		RuntimeSubstrateValBProvenanceStateActive,
		RuntimeSubstrateValBDriftCatalogStateActive,
	); got != RuntimeSubstrateValBStateIncomplete {
		t.Fatalf("expected incomplete state without active Val A, got %q", got)
	}
}

func TestRuntimeSubstrateValBProcessAndProvenanceStates(t *testing.T) {
	processItems := []RuntimeSubstrateProcessImageLinkage{
		{
			SubjectRef:   "cluster-a|acme-prod|Deployment|api",
			EventID:      "exec-api-1",
			Process:      ProcessIdentity{ProcessName: "api", ProcessPath: "/app/api", BinaryDigest: "sha256:bin"},
			Workload:     WorkloadIdentity{Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "api", ImageDigest: "sha256:image"},
			CurrentState: RuntimeSubstrateCorrelationStateSupported,
			DriftClass:   RuntimeSubstrateDriftExpected,
			Reasons:      []string{"binary_digest_matches_provenance_digest"},
		},
		{
			SubjectRef:        "cluster-a|acme-prod|Deployment|batch",
			EventID:           "exec-batch-1",
			Process:           ProcessIdentity{ProcessName: "batch"},
			Workload:          WorkloadIdentity{Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "batch"},
			CurrentState:      RuntimeSubstrateCorrelationStateUnsupported,
			DriftClass:        RuntimeSubstrateDriftLowRisk,
			UnsupportedFields: []string{"process.binary_digest"},
			Reasons:           []string{"direct_digest_match_not_supportable"},
		},
	}
	if got := EvaluateRuntimeSubstrateValBProcessImageState(processItems); got != RuntimeSubstrateValBProcessImageStateActive {
		t.Fatalf("expected active process-image state, got %q", got)
	}

	provenanceItems := []RuntimeSubstrateProvenanceLinkage{
		{
			SubjectRef:          "cluster-a|acme-prod|Deployment|api",
			CurrentState:        RuntimeSubstrateCorrelationStatePartial,
			DriftClass:          RuntimeSubstrateDriftLowRisk,
			Reasons:             []string{"signed_provenance_present_without_direct_digest_match"},
			SignerIdentities:    []string{"signer"},
			UnsupportedFields:   []string{"artifact.attestation_subject_digest"},
			WorkloadImageDigest: "sha256:image",
		},
	}
	if got := EvaluateRuntimeSubstrateValBProvenanceState(provenanceItems); got != RuntimeSubstrateValBProvenanceStateActive {
		t.Fatalf("expected active provenance state, got %q", got)
	}

	driftItems := []RuntimeSubstrateDriftRecord{
		{
			SubjectRef:   "cluster-a|acme-prod|Deployment|api",
			SourceKind:   "process_image",
			SourceRef:    "exec-api-1",
			CurrentState: RuntimeSubstrateCorrelationStateSupported,
			DriftClass:   RuntimeSubstrateDriftExpected,
			Summary:      "binary digest matches provenance digest",
			Reasons:      []string{"binary_digest_matches_provenance_digest"},
		},
	}
	if got := EvaluateRuntimeSubstrateValBDriftCatalogState(driftItems); got != RuntimeSubstrateValBDriftCatalogStateActive {
		t.Fatalf("expected active drift catalog state, got %q", got)
	}
}
