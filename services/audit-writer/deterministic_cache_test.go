package main

import (
	"reflect"
	"sync"
	"testing"
)

func captureVerifierEcosystemValDSharedStates() verifierEcosystemValDSharedStatesSnapshot {
	dependency, correctness, tooling, schemaCompatibility, diagnosticsConformance, trustKeyRotation, negativeDiagnostics, redaction, publisherArtifact, noOverclaim, correctnessState, toolingState, schemaCompatibilityState, diagnosticsConformanceState, trustKeyRotationState, negativeDiagnosticsState, redactionState, publisherArtifactState, noOverclaimState, valDState := buildVerifierEcosystemValDSharedStates()
	return verifierEcosystemValDSharedStatesSnapshot{
		Dependency:                  dependency,
		Correctness:                 correctness,
		Tooling:                     tooling,
		SchemaCompatibility:         schemaCompatibility,
		DiagnosticsConformance:      diagnosticsConformance,
		TrustKeyRotation:            trustKeyRotation,
		NegativeDiagnostics:         negativeDiagnostics,
		Redaction:                   redaction,
		PublisherArtifact:           publisherArtifact,
		NoOverclaim:                 noOverclaim,
		CorrectnessState:            correctnessState,
		ToolingState:                toolingState,
		SchemaCompatibilityState:    schemaCompatibilityState,
		DiagnosticsConformanceState: diagnosticsConformanceState,
		TrustKeyRotationState:       trustKeyRotationState,
		NegativeDiagnosticsState:    negativeDiagnosticsState,
		RedactionState:              redactionState,
		PublisherArtifactState:      publisherArtifactState,
		NoOverclaimState:            noOverclaimState,
		ValDState:                   valDState,
	}
}

func TestDeterministicAuditWriterCachesAreIsolated(t *testing.T) {
	t.Run("developer_valb_model", func(t *testing.T) {
		first := buildDeveloperEcosystemValBModel()
		if len(first.ProofSurfaceRefs) == 0 || len(first.ValECompatibility.SurfaceRefs) == 0 {
			t.Fatalf("expected canonical proof refs to exist, got %#v", first)
		}
		first.ProofSurfaceRefs[0] = "mutated-proof-ref"
		first.ValECompatibility.SurfaceRefs[0] = "mutated-vale-ref"

		second := buildDeveloperEcosystemValBModel()
		if second.ProofSurfaceRefs[0] == "mutated-proof-ref" || second.ValECompatibility.SurfaceRefs[0] == "mutated-vale-ref" {
			t.Fatalf("expected cached ValB model to remain isolated across calls, got %#v", second)
		}
	})

	t.Run("developer_valc_model", func(t *testing.T) {
		first := buildDeveloperEcosystemValCModel()
		if len(first.CapabilityDeclaration.DeclaredCapabilities) == 0 || len(first.ValBCompatibility.SurfaceRefs) == 0 {
			t.Fatalf("expected canonical capability and proof refs to exist, got %#v", first)
		}
		first.CapabilityDeclaration.DeclaredCapabilities[0] = "mutated-capability"
		first.ValBCompatibility.SurfaceRefs[0] = "mutated-surface-ref"

		second := buildDeveloperEcosystemValCModel()
		if second.CapabilityDeclaration.DeclaredCapabilities[0] == "mutated-capability" || second.ValBCompatibility.SurfaceRefs[0] == "mutated-surface-ref" {
			t.Fatalf("expected cached ValC model to remain isolated across calls, got %#v", second)
		}
	})

	t.Run("verifier_vald_shared_states", func(t *testing.T) {
		first := captureVerifierEcosystemValDSharedStates()
		if len(first.Correctness.SourceValStates) == 0 || len(first.DiagnosticsConformance.ObservedDiagnostics) == 0 {
			t.Fatalf("expected canonical ValD shared-state slices to exist, got %#v", first)
		}
		first.Correctness.SourceValStates[0] = "mutated-source-state"
		first.DiagnosticsConformance.ObservedDiagnostics[0] = "mutated-diagnostic"

		second := captureVerifierEcosystemValDSharedStates()
		if second.Correctness.SourceValStates[0] == "mutated-source-state" || second.DiagnosticsConformance.ObservedDiagnostics[0] == "mutated-diagnostic" {
			t.Fatalf("expected cached ValD shared states to remain isolated across calls, got %#v", second)
		}
	})

	t.Run("verifier_vale_closure_model", func(t *testing.T) {
		first := buildVerifierEcosystemValEClosureModel()
		if len(first.ProofSurfaceRefs) == 0 || len(first.ValB.SurfaceRefs) == 0 {
			t.Fatalf("expected canonical ValE closure refs to exist, got %#v", first)
		}
		first.ProofSurfaceRefs[0] = "mutated-proof-surface"
		first.ValB.SurfaceRefs[0] = "mutated-valb-surface"

		second := buildVerifierEcosystemValEClosureModel()
		if second.ProofSurfaceRefs[0] == "mutated-proof-surface" || second.ValB.SurfaceRefs[0] == "mutated-valb-surface" {
			t.Fatalf("expected cached ValE closure model to remain isolated across calls, got %#v", second)
		}
	})
}

func TestDeterministicAuditWriterCachesRemainConcurrentDeterministic(t *testing.T) {
	baselineB := buildDeveloperEcosystemValBModel()
	baselineD := captureVerifierEcosystemValDSharedStates()
	baselineE := buildVerifierEcosystemValEClosureModel()

	var wg sync.WaitGroup
	errs := make(chan string, 24)
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := buildDeveloperEcosystemValBModel(); !reflect.DeepEqual(got, baselineB) {
				errs <- "developer_valb_model"
			}
			if got := captureVerifierEcosystemValDSharedStates(); !reflect.DeepEqual(got, baselineD) {
				errs <- "verifier_vald_shared_states"
			}
			if got := buildVerifierEcosystemValEClosureModel(); !reflect.DeepEqual(got, baselineE) {
				errs <- "verifier_vale_closure_model"
			}
		}()
	}
	wg.Wait()
	close(errs)
	for name := range errs {
		t.Fatalf("expected concurrent cached builder %s to remain deterministic", name)
	}
}
