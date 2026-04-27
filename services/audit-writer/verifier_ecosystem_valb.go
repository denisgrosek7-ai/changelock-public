package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	verifierEcosystemValBCompatibilityMatrixSchema  = "point7.verifier_ecosystem.valb.compatibility_matrix.v1"
	verifierEcosystemValBSchemaProofSchema          = "point7.verifier_ecosystem.valb.schema_proof_compatibility.v1"
	verifierEcosystemValBMixedVersionSchema         = "point7.verifier_ecosystem.valb.mixed_version_diagnostics.v1"
	verifierEcosystemValBDiagnosticPrecedenceSchema = "point7.verifier_ecosystem.valb.diagnostic_precedence.v1"
	verifierEcosystemValBFixtureDescriptorsSchema   = "point7.verifier_ecosystem.valb.fixture_descriptors.v1"
	verifierEcosystemValBConformanceCasesSchema     = "point7.verifier_ecosystem.valb.conformance_cases.v1"
	verifierEcosystemValBConformanceSuiteSchema     = "point7.verifier_ecosystem.valb.conformance_suite.v1"
	verifierEcosystemValBOutputClassesSchema        = "point7.verifier_ecosystem.valb.output_classes.v1"
	verifierEcosystemValBProofsSchema               = "point7.verifier_ecosystem.valb.proofs.v1"
)

type verifierEcosystemValBModelResponse struct {
	SchemaVersion string    `json:"schema_version"`
	GeneratedAt   time.Time `json:"generated_at"`
	CurrentState  string    `json:"current_state"`
	Model         any       `json:"model"`
	RouteRefs     []string  `json:"route_refs,omitempty"`
	Limitations   []string  `json:"limitations,omitempty"`
}

type verifierEcosystemValBProofsResponse struct {
	SchemaVersion                 string    `json:"schema_version"`
	GeneratedAt                   time.Time `json:"generated_at"`
	CurrentState                  string    `json:"current_state"`
	Point5State                   string    `json:"point_5_state"`
	Point5DependencyState         string    `json:"point_5_dependency_state"`
	Point6State                   string    `json:"point_6_state"`
	Point6ClosureState            string    `json:"point_6_closure_state"`
	Point6ClosurePrerequisite     string    `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariant        string    `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState       string    `json:"point_6_proof_surface_state"`
	Point6PassRuleState           string    `json:"point_6_pass_rule_state"`
	Point6PassAllowed             bool      `json:"point_6_pass_allowed"`
	Val0CurrentState              string    `json:"val_0_current_state"`
	Val0State                     string    `json:"val_0_state"`
	ValACurrentState              string    `json:"val_a_current_state"`
	ValAState                     string    `json:"val_a_state"`
	ValBState                     string    `json:"val_b_state"`
	Point7State                   string    `json:"point_7_state"`
	CompatibilityMatrixState      string    `json:"compatibility_matrix_state"`
	SchemaProofCompatibilityState string    `json:"schema_proof_compatibility_state"`
	MixedVersionDiagnosticState   string    `json:"mixed_version_diagnostic_state"`
	DiagnosticPrecedenceState     string    `json:"diagnostic_precedence_state"`
	FixtureDescriptorState        string    `json:"fixture_descriptor_state"`
	ConformanceCaseState          string    `json:"conformance_case_state"`
	ConformanceSuiteState         string    `json:"conformance_suite_state"`
	OutputClassState              string    `json:"output_class_state"`
	CompatibilityState            string    `json:"compatibility_state"`
	DerivedDiagnosticClass        string    `json:"derived_diagnostic_class"`
	DerivedOutputClass            string    `json:"derived_output_class"`
	WhyPoint7NotPass              []string  `json:"why_point_7_not_pass,omitempty"`
	SurfaceRefs                   []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                  []string  `json:"evidence_refs,omitempty"`
	Limitations                   []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer          string    `json:"projection_disclaimer"`
	IntegrationSummary            []string  `json:"integration_summary,omitempty"`
}

func verifierEcosystemValBAllSurfaceRefs() []string {
	return operability.VerifierEcosystemValBProofSurfaceRefs()
}

func verifierEcosystemValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth compatibility_diagnostics_conformance advisory_projection"
}

func verifierEcosystemValBEvidenceRefs() []string {
	return operability.VerifierEcosystemValBProofEvidenceRefs()
}

func buildVerifierEcosystemValBDependencySnapshot() operability.VerifierEcosystemValBDependencySnapshot {
	valE := buildReferenceArchitectureValEProofs()
	val0 := buildVerifierEcosystemVal0Proofs()
	valA := buildVerifierEcosystemValAProofs()
	return operability.VerifierEcosystemValBDependencySnapshot{
		Point5State:                    valE.Point5State,
		Point5DependencyState:          valE.Point5DependencyState,
		Point6State:                    valE.Point6State,
		Point6ClosureState:             valE.ValEState,
		Point6ClosurePrerequisiteState: valE.ClosurePrerequisiteState,
		Point6ClosureInvariantState:    valE.ClosureInvariantState,
		Point6ProofSurfaceState:        valE.ProofSurfaceState,
		Point6PassRuleState:            valE.PassRuleState,
		Point6PassAllowed:              valE.Point6PassAllowed,
		Val0CurrentState:               val0.CurrentState,
		Val0State:                      val0.Val0State,
		ValACurrentState:               valA.CurrentState,
		ValAState:                      valA.ValAState,
		Point7State:                    valA.Point7State,
	}
}

func buildVerifierEcosystemValBSharedStates() (
	operability.VerifierEcosystemValBDependencySnapshot,
	operability.VerifierEcosystemValBCompatibilityMatrix,
	operability.VerifierEcosystemValBSchemaProofCompatibility,
	operability.VerifierEcosystemValBMixedVersionDiagnosticsCatalog,
	operability.VerifierEcosystemValBDiagnosticPrecedence,
	operability.VerifierEcosystemValBFixtureCatalog,
	operability.VerifierEcosystemValBConformanceCaseCatalog,
	operability.VerifierEcosystemValBConformanceSuite,
	operability.VerifierEcosystemValBOutputClassCatalog,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
) {
	dependency := buildVerifierEcosystemValBDependencySnapshot()
	_, _, _, valAResult, _, _, _, _, _, _, _, _, _, _ := buildVerifierEcosystemValASharedStates()
	compatibilityMatrix := operability.VerifierEcosystemValBCompatibilityMatrixModel()
	schemaProofCompatibility := operability.VerifierEcosystemValBSchemaProofCompatibilityModel()
	schemaProofCompatibility.SchemaVersion = valAResult.SchemaVersion
	schemaProofCompatibility.ProofType = valAResult.ProofType
	schemaProofCompatibility.VerifierVersion = valAResult.VerifierVersion
	mixedVersion := operability.VerifierEcosystemValBMixedVersionDiagnosticsCatalogModel()
	diagnosticPrecedence := operability.VerifierEcosystemValBDiagnosticPrecedenceModel()
	diagnosticPrecedence.ObservedDiagnostics = []string{valAResult.DiagnosticClass}
	diagnosticPrecedence.DerivedDiagnosticClass = operability.DeriveVerifierEcosystemValBDiagnostic(diagnosticPrecedence.ObservedDiagnostics, diagnosticPrecedence.Caveats)
	fixtures := operability.VerifierEcosystemValBFixtureCatalogModel()
	conformanceCases := operability.VerifierEcosystemValBConformanceCaseCatalogModel()
	conformanceSuite := operability.VerifierEcosystemValBConformanceSuiteModel()
	outputClasses := operability.VerifierEcosystemValBOutputClassCatalogModel()

	compatibilityMatrixState := operability.EvaluateVerifierEcosystemValBCompatibilityMatrixState(compatibilityMatrix)
	schemaProofCompatibilityState := operability.EvaluateVerifierEcosystemValBSchemaProofCompatibilityState(schemaProofCompatibility, compatibilityMatrix)
	mixedVersionState := operability.EvaluateVerifierEcosystemValBMixedVersionDiagnosticsState(mixedVersion)
	diagnosticPrecedenceState := operability.EvaluateVerifierEcosystemValBDiagnosticPrecedenceState(diagnosticPrecedence)
	fixtureDescriptorState := operability.EvaluateVerifierEcosystemValBFixtureDescriptorState(fixtures)
	outputClassState := operability.EvaluateVerifierEcosystemValBOutputClassState(outputClasses)
	conformanceCaseState := operability.EvaluateVerifierEcosystemValBConformanceCaseState(conformanceCases, fixtures, outputClasses)
	conformanceSuiteState := operability.EvaluateVerifierEcosystemValBConformanceSuiteState(conformanceSuite, conformanceCases, fixtures, outputClasses)
	valBState := operability.EvaluateVerifierEcosystemValBState(
		dependency,
		compatibilityMatrixState,
		schemaProofCompatibilityState,
		mixedVersionState,
		diagnosticPrecedenceState,
		fixtureDescriptorState,
		conformanceCaseState,
		conformanceSuiteState,
		outputClassState,
	)
	return dependency, compatibilityMatrix, schemaProofCompatibility, mixedVersion, diagnosticPrecedence, fixtures, conformanceCases, conformanceSuite, outputClasses, compatibilityMatrixState, schemaProofCompatibilityState, mixedVersionState, diagnosticPrecedenceState, fixtureDescriptorState, conformanceCaseState, conformanceSuiteState, outputClassState, valBState
}

func verifierEcosystemValBDerivedOutputClass(catalog operability.VerifierEcosystemValBOutputClassCatalog, overallResult, diagnosticClass string) string {
	for _, mapping := range catalog.Mappings {
		if mapping.OverallResult == overallResult && mapping.DiagnosticClass == diagnosticClass {
			return mapping.OutputClass
		}
	}
	return operability.VerifierEcosystemValBOutputClassUnknown
}

func (s server) verifierEcosystemValBCompatibilityMatrixHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBCompatibilityMatrix())
}

func (s server) verifierEcosystemValBSchemaProofCompatibilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBSchemaProofCompatibility())
}

func (s server) verifierEcosystemValBMixedVersionDiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBMixedVersionDiagnostics())
}

func (s server) verifierEcosystemValBDiagnosticPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBDiagnosticPrecedence())
}

func (s server) verifierEcosystemValBFixtureDescriptorsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBFixtureDescriptors())
}

func (s server) verifierEcosystemValBConformanceCasesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBConformanceCases())
}

func (s server) verifierEcosystemValBConformanceSuiteHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBConformanceSuite())
}

func (s server) verifierEcosystemValBOutputClassesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBOutputClasses())
}

func (s server) verifierEcosystemValBProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValBProofs())
}

func buildVerifierEcosystemValBCompatibilityMatrix() verifierEcosystemValBModelResponse {
	_, compatibilityMatrix, _, _, _, _, _, _, _, compatibilityMatrixState, _, _, _, _, _, _, _, _ := buildVerifierEcosystemValBSharedStates()
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBCompatibilityMatrixSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  compatibilityMatrixState,
		Model:         compatibilityMatrix,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/schema-proof-compatibility",
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Compatibility matrix remains advisory, version-bound, and does not certify a verifier.",
			"Deprecated, superseded, warning-bearing, and unsupported compatibility states remain explicit and fail closed.",
		},
	}
}

func buildVerifierEcosystemValBSchemaProofCompatibility() verifierEcosystemValBModelResponse {
	_, compatibilityMatrix, schemaProofCompatibility, _, _, _, _, _, _, _, schemaProofCompatibilityState, _, _, _, _, _, _, _ := buildVerifierEcosystemValBSharedStates()
	_ = compatibilityMatrix
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBSchemaProofSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  schemaProofCompatibilityState,
		Model:         schemaProofCompatibility,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/mixed-version-diagnostics",
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Schema and proof-type compatibility remains bounded to declared matrix entries and explicit migration or supersession guidance.",
			"Unsupported combinations remain non-verified and fail closed.",
		},
	}
}

func buildVerifierEcosystemValBMixedVersionDiagnostics() verifierEcosystemValBModelResponse {
	_, _, _, mixedVersion, _, _, _, _, _, _, _, mixedVersionState, _, _, _, _, _, _ := buildVerifierEcosystemValBSharedStates()
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBMixedVersionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  mixedVersionState,
		Model:         mixedVersion,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/diagnostic-precedence",
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Mixed-version diagnostics remain bounded to explicit declared cases and do not claim universal compatibility.",
		},
	}
}

func buildVerifierEcosystemValBDiagnosticPrecedence() verifierEcosystemValBModelResponse {
	_, _, _, _, diagnosticPrecedence, _, _, _, _, _, _, _, diagnosticPrecedenceState, _, _, _, _, _ := buildVerifierEcosystemValBSharedStates()
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBDiagnosticPrecedenceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  diagnosticPrecedenceState,
		Model:         diagnosticPrecedence,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/output-classes",
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Diagnostic precedence remains deterministic, fail-closed, and explicit about warning-bearing versus hard-failure outcomes.",
		},
	}
}

func buildVerifierEcosystemValBFixtureDescriptors() verifierEcosystemValBModelResponse {
	_, _, _, _, _, fixtures, _, _, _, _, _, _, _, fixtureDescriptorState, _, _, _, _ := buildVerifierEcosystemValBSharedStates()
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBFixtureDescriptorsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  fixtureDescriptorState,
		Model:         fixtures,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/conformance-cases",
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Fixture descriptors are bounded conformance descriptors only and do not create fake production evidence or a public verifier hub.",
		},
	}
}

func buildVerifierEcosystemValBConformanceCases() verifierEcosystemValBModelResponse {
	_, _, _, _, _, _, conformanceCases, _, _, _, _, _, _, _, conformanceCaseState, _, _, _ := buildVerifierEcosystemValBSharedStates()
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBConformanceCasesSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  conformanceCaseState,
		Model:         conformanceCases,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/conformance-suite",
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Conformance cases prove behavior against declared fixtures only and do not certify a verifier or approve deployments.",
		},
	}
}

func buildVerifierEcosystemValBConformanceSuite() verifierEcosystemValBModelResponse {
	_, _, _, _, _, _, _, conformanceSuite, _, _, _, _, _, _, _, conformanceSuiteState, _, _ := buildVerifierEcosystemValBSharedStates()
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBConformanceSuiteSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  conformanceSuiteState,
		Model:         conformanceSuite,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/output-classes",
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Conformance suite remains advisory and bounded to declared cases; it does not create verifier certification or integrity ratings.",
		},
	}
}

func buildVerifierEcosystemValBOutputClasses() verifierEcosystemValBModelResponse {
	_, _, _, _, _, _, _, _, outputClasses, _, _, _, _, _, _, _, outputClassState, _ := buildVerifierEcosystemValBSharedStates()
	return verifierEcosystemValBModelResponse{
		SchemaVersion: verifierEcosystemValBOutputClassesSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  outputClassState,
		Model:         outputClasses,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/valb/proofs",
		},
		Limitations: []string{
			"Verification output classes remain bounded advisory report classes and do not convert non-verified outcomes into verified results.",
		},
	}
}

func buildVerifierEcosystemValBProofs() verifierEcosystemValBProofsResponse {
	dependency, compatibilityMatrix, schemaProofCompatibility, mixedVersion, diagnosticPrecedence, fixtures, conformanceCases, conformanceSuite, outputClasses, compatibilityMatrixState, schemaProofCompatibilityState, mixedVersionState, diagnosticPrecedenceState, fixtureDescriptorState, conformanceCaseState, conformanceSuiteState, outputClassState, valBState := buildVerifierEcosystemValBSharedStates()
	_, _, _, valAResult, _, _, _, _, _, _, _, _, _, _ := buildVerifierEcosystemValASharedStates()
	_ = compatibilityMatrix
	_ = mixedVersion
	_ = fixtures
	_ = conformanceCases
	_ = conformanceSuite
	point7State := operability.EvaluateVerifierEcosystemValBPoint7State(valBState)
	derivedOutputClass := verifierEcosystemValBDerivedOutputClass(outputClasses, valAResult.OverallResult, diagnosticPrecedence.DerivedDiagnosticClass)
	limitations := []string{
		"Val B implements compatibility, diagnostics, and conformance only and does not implement public verifier hub, partner or auditor portal, third-party publisher profile, or final Točka 7 closure.",
		"Conformance suites remain advisory and do not certify a verifier or produce marketplace ratings.",
		"Unsupported, stale, revoked, superseded, malformed, and scope-breaking artifacts remain explicitly non-verified.",
	}
	currentState := operability.EvaluateVerifierEcosystemValBProofsState(
		valBState,
		point7State,
		dependency.Val0CurrentState,
		dependency.ValACurrentState,
		verifierEcosystemValBAllSurfaceRefs(),
		verifierEcosystemValBEvidenceRefs(),
		limitations,
		[]string{
			"Val B cannot return point_7_pass and remains compatibility, diagnostics, and conformance only.",
			"Val C through Val E remain required before Točka 7 integrated closure can exist.",
			"Verifier outputs remain advisory and bounded over the canonical execution, audit, and evidence spine.",
		},
		verifierEcosystemValBProjectionDisclaimer(),
	)
	return verifierEcosystemValBProofsResponse{
		SchemaVersion:                 verifierEcosystemValBProofsSchema,
		GeneratedAt:                   publicSampleTime(),
		CurrentState:                  currentState,
		Point5State:                   dependency.Point5State,
		Point5DependencyState:         dependency.Point5DependencyState,
		Point6State:                   dependency.Point6State,
		Point6ClosureState:            dependency.Point6ClosureState,
		Point6ClosurePrerequisite:     dependency.Point6ClosurePrerequisiteState,
		Point6ClosureInvariant:        dependency.Point6ClosureInvariantState,
		Point6ProofSurfaceState:       dependency.Point6ProofSurfaceState,
		Point6PassRuleState:           dependency.Point6PassRuleState,
		Point6PassAllowed:             dependency.Point6PassAllowed,
		Val0CurrentState:              dependency.Val0CurrentState,
		Val0State:                     dependency.Val0State,
		ValACurrentState:              dependency.ValACurrentState,
		ValAState:                     dependency.ValAState,
		ValBState:                     valBState,
		Point7State:                   point7State,
		CompatibilityMatrixState:      compatibilityMatrixState,
		SchemaProofCompatibilityState: schemaProofCompatibilityState,
		MixedVersionDiagnosticState:   mixedVersionState,
		DiagnosticPrecedenceState:     diagnosticPrecedenceState,
		FixtureDescriptorState:        fixtureDescriptorState,
		ConformanceCaseState:          conformanceCaseState,
		ConformanceSuiteState:         conformanceSuiteState,
		OutputClassState:              outputClassState,
		CompatibilityState:            schemaProofCompatibility.CompatibilityState,
		DerivedDiagnosticClass:        diagnosticPrecedence.DerivedDiagnosticClass,
		DerivedOutputClass:            derivedOutputClass,
		WhyPoint7NotPass: []string{
			"Val B expands compatibility, diagnostics, and conformance only and cannot return point_7_pass.",
			"Točka 7 remains not complete until Val E integrated closure.",
			"Verifier quality outputs remain advisory and do not create certification, approval, or canonical authority.",
		},
		SurfaceRefs:          verifierEcosystemValBAllSurfaceRefs(),
		EvidenceRefs:         verifierEcosystemValBEvidenceRefs(),
		Limitations:          limitations,
		ProjectionDisclaimer: verifierEcosystemValBProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val B makes verifier behavior deterministic and explainable across compatibility, diagnostics, fixtures, conformance cases, and output classes.",
			"Val B builds on actual Točka 6 closure plus active Točka 7 Val 0 and Val A proof surfaces while keeping point_7_state not complete.",
		},
	}
}
