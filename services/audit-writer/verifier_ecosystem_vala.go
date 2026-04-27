package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
	verifiercore "github.com/denisgrosek/changelock/internal/verifier"
)

const (
	verifierEcosystemValAInputModelSchema         = "point7.verifier_ecosystem.vala.input_model.v1"
	verifierEcosystemValAVerifierEngineSchema     = "point7.verifier_ecosystem.vala.verifier_engine.v1"
	verifierEcosystemValAVerificationReportSchema = "point7.verifier_ecosystem.vala.verification_report.v1"
	verifierEcosystemValADiagnosticsMappingSchema = "point7.verifier_ecosystem.vala.diagnostics_mapping.v1"
	verifierEcosystemValACommandContractSchema    = "point7.verifier_ecosystem.vala.command_contract.v1"
	verifierEcosystemValASDKEntrypointSchema      = "point7.verifier_ecosystem.vala.sdk_entrypoint.v1"
	verifierEcosystemValAProofsSchema             = "point7.verifier_ecosystem.vala.proofs.v1"
)

type verifierEcosystemValAModelResponse struct {
	SchemaVersion string    `json:"schema_version"`
	GeneratedAt   time.Time `json:"generated_at"`
	CurrentState  string    `json:"current_state"`
	Model         any       `json:"model"`
	RouteRefs     []string  `json:"route_refs,omitempty"`
	Limitations   []string  `json:"limitations,omitempty"`
}

type verifierEcosystemValAProofsResponse struct {
	SchemaVersion               string    `json:"schema_version"`
	GeneratedAt                 time.Time `json:"generated_at"`
	CurrentState                string    `json:"current_state"`
	Point5State                 string    `json:"point_5_state"`
	Point5DependencyState       string    `json:"point_5_dependency_state"`
	Point6State                 string    `json:"point_6_state"`
	Point6ClosureState          string    `json:"point_6_closure_state"`
	Point6ClosurePrerequisite   string    `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariant      string    `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState     string    `json:"point_6_proof_surface_state"`
	Point6PassRuleState         string    `json:"point_6_pass_rule_state"`
	Point6PassAllowed           bool      `json:"point_6_pass_allowed"`
	Val0CurrentState            string    `json:"val_0_current_state"`
	Val0State                   string    `json:"val_0_state"`
	ValAState                   string    `json:"val_a_state"`
	Point7State                 string    `json:"point_7_state"`
	InputModelState             string    `json:"input_model_state"`
	VerifierEngineState         string    `json:"verifier_engine_state"`
	VerificationResultState     string    `json:"verification_result_state"`
	DiagnosticsMappingState     string    `json:"diagnostics_mapping_state"`
	CommandContractState        string    `json:"command_contract_state"`
	SDKEntrypointState          string    `json:"sdk_entrypoint_state"`
	VerificationOverallResult   string    `json:"verification_overall_result"`
	VerificationDiagnosticClass string    `json:"verification_diagnostic_class"`
	WhyPoint7NotPass            []string  `json:"why_point_7_not_pass,omitempty"`
	SurfaceRefs                 []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                []string  `json:"evidence_refs,omitempty"`
	Limitations                 []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer        string    `json:"projection_disclaimer"`
	IntegrationSummary          []string  `json:"integration_summary,omitempty"`
}

func verifierEcosystemValAAllSurfaceRefs() []string {
	return operability.VerifierEcosystemValAProofSurfaceRefs()
}

func verifierEcosystemValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth reference_verifier_tooling advisory_projection"
}

func verifierEcosystemValAEvidenceRefs() []string {
	refs := []string{"point6_integrated_closure", "point7_verifier_discipline_foundation", "point7_reference_verifier_tooling"}
	for _, evidence := range operability.VerifierEcosystemValAVerifierEvidence() {
		if evidence.EvidenceID != "" {
			refs = append(refs, evidence.EvidenceID)
		}
	}
	return refs
}

func buildVerifierEcosystemValADependencySnapshot() operability.VerifierEcosystemValADependencySnapshot {
	valE := buildReferenceArchitectureValEProofs()
	val0 := buildVerifierEcosystemVal0Proofs()
	return operability.VerifierEcosystemValADependencySnapshot{
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
		Point7State:                    val0.Point7State,
	}
}

func verifierEcosystemValASDKRequest(input operability.VerifierEcosystemValAReferenceVerifierInput, reportFormat string) verifiercore.ReferenceVerifierRequest {
	return verifiercore.ReferenceVerifierRequest{
		RequestID:                       input.VerificationRequestID,
		VerifierContractRef:             input.VerifierContractRef,
		ProofEnvelopeRef:                input.ProofEnvelopeRef,
		ArtifactRef:                     input.ArtifactRef,
		ArtifactDigest:                  input.ArtifactDigest,
		ArtifactDigestAlgorithm:         input.ArtifactDigestAlgorithm,
		SignatureRef:                    input.SignatureRef,
		IssuerRef:                       input.IssuerRef,
		TrustRootRef:                    input.TrustRootRef,
		SchemaVersion:                   input.SchemaVersion,
		ProofType:                       input.ProofType,
		RequestedScope:                  input.RequestedScope,
		VerificationTime:                input.VerificationTime,
		ExpectedOutputBoundary:          input.ExpectedOutputBoundary,
		CompatibilityPolicyRef:          input.CompatibilityPolicyRef,
		RevocationMaterialRef:           input.RevocationMaterialRef,
		SupersessionMaterialRef:         input.SupersessionMaterialRef,
		EvidenceRefs:                    input.EvidenceRefs,
		DigestVerificationState:         input.DigestVerificationState,
		SignatureVerificationState:      input.SignatureVerificationState,
		SchemaVerificationState:         input.SchemaVerificationState,
		ScopeVerificationState:          input.ScopeVerificationState,
		FreshnessVerificationState:      input.FreshnessVerificationState,
		TrustRootVerificationState:      input.TrustRootVerificationState,
		IssuerVerificationState:         input.IssuerVerificationState,
		CompatibilityEvaluationState:    input.CompatibilityEvaluationState,
		RevocationEvaluationState:       input.RevocationEvaluationState,
		SupersessionEvaluationState:     input.SupersessionEvaluationState,
		LineageVerificationState:        input.LineageVerificationState,
		OutputBoundaryVerificationState: input.OutputBoundaryVerificationState,
		ReportFormat:                    reportFormat,
		StrictFailClosed:                input.StrictFailClosed,
		TruthOutsideScopeClaim:          input.TruthOutsideScopeClaim,
		ClaimsActualCryptoVerification:  input.ClaimsActualCryptoVerification,
		Caveats:                         input.Caveats,
		ProjectionDisclaimer:            input.ProjectionDisclaimer,
	}
}

func verifierEcosystemValAResultFromSDK(result verifiercore.ReferenceVerificationResult) operability.VerifierEcosystemValAVerificationResult {
	return operability.VerifierEcosystemValAVerificationResult{
		CurrentState:           "verifier_ecosystem_vala_result_ready",
		VerificationResultID:   result.ResultID,
		RequestID:              result.RequestID,
		VerifierVersion:        result.VerifierVersion,
		ProofType:              result.ProofType,
		SchemaVersion:          result.SchemaVersion,
		Scope:                  result.Scope,
		OutputBoundary:         result.OutputBoundary,
		OverallResult:          result.OverallResult,
		DiagnosticClass:        result.DiagnosticClass,
		DigestResult:           result.DigestResult,
		SignatureResult:        result.SignatureResult,
		SchemaResult:           result.SchemaResult,
		ScopeResult:            result.ScopeResult,
		FreshnessResult:        result.FreshnessResult,
		TrustRootResult:        result.TrustRootResult,
		IssuerResult:           result.IssuerResult,
		CompatibilityResult:    result.CompatibilityResult,
		RevocationResult:       result.RevocationResult,
		SupersessionResult:     result.SupersessionResult,
		LineageResult:          result.LineageResult,
		OutputBoundaryResult:   result.OutputBoundaryResult,
		EvidenceRefs:           result.EvidenceRefs,
		Caveats:                result.Caveats,
		Limitations:            result.Limitations,
		ProjectionDisclaimer:   result.ProjectionDisclaimer,
		VerifiedAt:             result.VerifiedAt,
		TruthOutsideScopeClaim: result.TruthOutsideScopeClaim,
	}
}

func buildVerifierEcosystemValASharedStates() (
	operability.VerifierEcosystemValADependencySnapshot,
	operability.VerifierEcosystemValAReferenceVerifierInput,
	operability.VerifierEcosystemValAReferenceVerifierEngine,
	operability.VerifierEcosystemValAVerificationResult,
	operability.VerifierEcosystemValADiagnosticsMapping,
	operability.VerifierEcosystemValACommandContract,
	operability.VerifierEcosystemValASDKEntrypoint,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
) {
	dependency := buildVerifierEcosystemValADependencySnapshot()
	input := operability.VerifierEcosystemValAReferenceVerifierInputModel()
	command := operability.VerifierEcosystemValACommandContractModel()
	sdk := operability.VerifierEcosystemValASDKEntrypointModel()
	_, _, _, _, _, _, _, _, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, _ := buildVerifierEcosystemVal0SharedStates()
	engine := operability.VerifierEcosystemValAReferenceVerifierEngineModel(
		contractState,
		envelopeState,
		scopeState,
		compatibilityState,
		trustState,
		diagnosticsState,
		outputBoundaryState,
	)
	result := verifierEcosystemValAResultFromSDK(verifiercore.VerifyReferenceVerifierRequest(
		verifierEcosystemValASDKRequest(input, command.ReportFormat),
	))
	diagnosticsMapping := operability.VerifierEcosystemValADiagnosticsMappingModel(result)

	inputState := operability.EvaluateVerifierEcosystemValAReferenceVerifierInputState(input)
	engineState := operability.EvaluateVerifierEcosystemValAReferenceVerifierEngineState(engine)
	resultState := operability.EvaluateVerifierEcosystemValAVerificationResultState(result)
	diagnosticsMappingState := operability.EvaluateVerifierEcosystemValADiagnosticsMappingState(diagnosticsMapping)
	commandState := operability.EvaluateVerifierEcosystemValACommandContractState(command)
	sdkState := operability.EvaluateVerifierEcosystemValASDKEntrypointState(sdk)
	valAState := operability.EvaluateVerifierEcosystemValAState(
		dependency,
		inputState,
		engineState,
		resultState,
		diagnosticsMappingState,
		commandState,
		sdkState,
	)
	return dependency, input, engine, result, diagnosticsMapping, command, sdk, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, valAState
}

func (s server) verifierEcosystemValAInputModelHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValAInputModel())
}

func (s server) verifierEcosystemValAVerifierEngineHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValAVerifierEngine())
}

func (s server) verifierEcosystemValAVerificationReportHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValAVerificationReport())
}

func (s server) verifierEcosystemValADiagnosticsMappingHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValADiagnosticsMapping())
}

func (s server) verifierEcosystemValACommandContractHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValACommandContract())
}

func (s server) verifierEcosystemValASDKEntrypointHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValASDKEntrypoint())
}

func (s server) verifierEcosystemValAProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemValAProofs())
}

func buildVerifierEcosystemValAInputModel() verifierEcosystemValAModelResponse {
	_, input, _, _, _, _, _, inputState, _, _, _, _, _, _ := buildVerifierEcosystemValASharedStates()
	return verifierEcosystemValAModelResponse{
		SchemaVersion: verifierEcosystemValAInputModelSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  inputState,
		Model:         input,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vala/verifier-engine",
			"/v1/verifier-ecosystem/vala/proofs",
		},
		Limitations: []string{
			"Reference verifier input remains bounded to declared proof envelope, schema, scope, trust-root material, freshness, compatibility, revocation, supersession, and output-boundary descriptors.",
			"Input descriptors do not imply access to the canonical evidence spine.",
		},
	}
}

func buildVerifierEcosystemValAVerifierEngine() verifierEcosystemValAModelResponse {
	_, _, engine, _, _, _, _, _, engineState, _, _, _, _, _ := buildVerifierEcosystemValASharedStates()
	return verifierEcosystemValAModelResponse{
		SchemaVersion: verifierEcosystemValAVerifierEngineSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  engineState,
		Model:         engine,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vala/verification-report",
			"/v1/verifier-ecosystem/vala/proofs",
		},
		Limitations: []string{
			"Reference verifier engine remains deterministic, fail-closed, and bounded to explicit modeled verification semantics unless repository crypto primitives are actually invoked.",
			"Engine output remains advisory and cannot approve deployment or create canonical truth.",
		},
	}
}

func buildVerifierEcosystemValAVerificationReport() verifierEcosystemValAModelResponse {
	_, _, _, result, _, _, _, _, _, resultState, _, _, _, _ := buildVerifierEcosystemValASharedStates()
	return verifierEcosystemValAModelResponse{
		SchemaVersion: verifierEcosystemValAVerificationReportSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  resultState,
		Model:         result,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vala/diagnostics-mapping",
			"/v1/verifier-ecosystem/vala/proofs",
		},
		Limitations: []string{
			"Verification reports remain bounded by proof envelope, schema, scope, freshness, trust-root, issuer, compatibility, revocation, supersession, lineage, and output-boundary semantics.",
			"Reports do not claim truth outside the declared verifier scope.",
		},
	}
}

func buildVerifierEcosystemValADiagnosticsMapping() verifierEcosystemValAModelResponse {
	_, _, _, _, diagnosticsMapping, _, _, _, _, _, diagnosticsMappingState, _, _, _ := buildVerifierEcosystemValASharedStates()
	return verifierEcosystemValAModelResponse{
		SchemaVersion: verifierEcosystemValADiagnosticsMappingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  diagnosticsMappingState,
		Model:         diagnosticsMapping,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vala/verification-report",
			"/v1/verifier-ecosystem/vala/proofs",
		},
		Limitations: []string{
			"Diagnostic mapping stays deterministic and preserves failure precedence instead of suppressing invalid, stale, revoked, unsupported, or incomplete conditions.",
		},
	}
}

func buildVerifierEcosystemValACommandContract() verifierEcosystemValAModelResponse {
	_, _, _, _, _, command, _, _, _, _, _, commandState, _, _ := buildVerifierEcosystemValASharedStates()
	return verifierEcosystemValAModelResponse{
		SchemaVersion: verifierEcosystemValACommandContractSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  commandState,
		Model:         command,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vala/sdk-entrypoint",
			"/v1/verifier-ecosystem/vala/proofs",
		},
		Limitations: []string{
			"Val A defines a narrow CLI-oriented command contract only and does not introduce a full standalone verifier CLI subsystem.",
			"Command contract remains non-mutating, non-approving, non-suppressing, and non-publishing.",
		},
	}
}

func buildVerifierEcosystemValASDKEntrypoint() verifierEcosystemValAModelResponse {
	_, _, _, _, _, _, sdk, _, _, _, _, _, sdkState, _ := buildVerifierEcosystemValASharedStates()
	return verifierEcosystemValAModelResponse{
		SchemaVersion: verifierEcosystemValASDKEntrypointSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  sdkState,
		Model:         sdk,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/vala/verification-report",
			"/v1/verifier-ecosystem/vala/proofs",
		},
		Limitations: []string{
			"SDK entrypoint remains deterministic, local, and advisory; it does not mutate canonical state or require the main ChangeLock instance.",
		},
	}
}

func buildVerifierEcosystemValAProofs() verifierEcosystemValAProofsResponse {
	dependency, _, _, result, _, _, _, inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState, valAState := buildVerifierEcosystemValASharedStates()
	point7State := operability.EvaluateVerifierEcosystemValAPoint7State(valAState)
	val0 := buildVerifierEcosystemVal0Proofs()
	limitations := []string{
		"Val A implements reference verifier tooling only and does not implement public verifier hub, partner or auditor portal, third-party publisher profile, or final Točka 7 closure.",
		"If repository cryptographic primitives are not actually invoked, digest and signature handling remains modeled verification semantics and must not be read as claimed full cryptographic verification.",
		"Verifier outputs remain advisory projections over the canonical execution, audit, and evidence spine and do not become approval, certification, publication, suppression, or mutation authority.",
	}
	evidenceRefs := referenceArchitectureValEUniqueRefs(val0.EvidenceRefs, verifierEcosystemValAEvidenceRefs(), result.EvidenceRefs)
	currentState := operability.EvaluateVerifierEcosystemValAProofsState(
		valAState,
		point7State,
		dependency.Val0CurrentState,
		verifierEcosystemValAAllSurfaceRefs(),
		evidenceRefs,
		limitations,
		[]string{
			"Val A cannot return point_7_pass and remains bounded verifier tooling only.",
			"Val B through Val E remain required before Točka 7 integrated closure can exist.",
			"Verifier tooling remains advisory and scope-bounded over the canonical execution, audit, and evidence spine.",
		},
		verifierEcosystemValAProjectionDisclaimer(),
	)
	return verifierEcosystemValAProofsResponse{
		SchemaVersion:               verifierEcosystemValAProofsSchema,
		GeneratedAt:                 publicSampleTime(),
		CurrentState:                currentState,
		Point5State:                 dependency.Point5State,
		Point5DependencyState:       dependency.Point5DependencyState,
		Point6State:                 dependency.Point6State,
		Point6ClosureState:          dependency.Point6ClosureState,
		Point6ClosurePrerequisite:   dependency.Point6ClosurePrerequisiteState,
		Point6ClosureInvariant:      dependency.Point6ClosureInvariantState,
		Point6ProofSurfaceState:     dependency.Point6ProofSurfaceState,
		Point6PassRuleState:         dependency.Point6PassRuleState,
		Point6PassAllowed:           dependency.Point6PassAllowed,
		Val0CurrentState:            dependency.Val0CurrentState,
		Val0State:                   dependency.Val0State,
		ValAState:                   valAState,
		Point7State:                 point7State,
		InputModelState:             inputState,
		VerifierEngineState:         engineState,
		VerificationResultState:     resultState,
		DiagnosticsMappingState:     diagnosticsMappingState,
		CommandContractState:        commandState,
		SDKEntrypointState:          sdkState,
		VerificationOverallResult:   result.OverallResult,
		VerificationDiagnosticClass: result.DiagnosticClass,
		WhyPoint7NotPass: []string{
			"Val A provides bounded reference verifier tooling only and cannot return point_7_pass.",
			"Točka 7 remains not complete until Val E integrated closure.",
			"Verifier tooling remains advisory and does not create canonical truth or deployment approval.",
		},
		SurfaceRefs:          verifierEcosystemValAAllSurfaceRefs(),
		EvidenceRefs:         evidenceRefs,
		Limitations:          limitations,
		ProjectionDisclaimer: verifierEcosystemValAProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val A operationalizes Val 0 verifier discipline through a deterministic reference verifier engine, CLI-oriented command contract surface, and primary SDK entrypoint.",
			"Reference verifier tooling remains bounded by proof envelope, schema, scope, trust-root material, freshness, compatibility, revocation, supersession, lineage, and output-boundary discipline.",
		},
	}
}
