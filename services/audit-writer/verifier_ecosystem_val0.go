package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	verifierEcosystemVal0ContractSchema       = "point7.verifier_ecosystem.val0.contract.v1"
	verifierEcosystemVal0EnvelopeSchema       = "point7.verifier_ecosystem.val0.proof_envelope.v1"
	verifierEcosystemVal0ScopeSchema          = "point7.verifier_ecosystem.val0.verification_scope.v1"
	verifierEcosystemVal0CompatibilitySchema  = "point7.verifier_ecosystem.val0.schema_compatibility.v1"
	verifierEcosystemVal0TrustIssuerSchema    = "point7.verifier_ecosystem.val0.trust_root_issuer.v1"
	verifierEcosystemVal0DiagnosticsSchema    = "point7.verifier_ecosystem.val0.diagnostics.v1"
	verifierEcosystemVal0OutputBoundarySchema = "point7.verifier_ecosystem.val0.output_boundaries.v1"
	verifierEcosystemVal0ProofsSchema         = "point7.verifier_ecosystem.val0.proofs.v1"
)

type verifierEcosystemVal0ModelResponse struct {
	SchemaVersion string    `json:"schema_version"`
	GeneratedAt   time.Time `json:"generated_at"`
	CurrentState  string    `json:"current_state"`
	Model         any       `json:"model"`
	RouteRefs     []string  `json:"route_refs,omitempty"`
	Limitations   []string  `json:"limitations,omitempty"`
}

type verifierEcosystemVal0ProofsResponse struct {
	SchemaVersion              string    `json:"schema_version"`
	GeneratedAt                time.Time `json:"generated_at"`
	CurrentState               string    `json:"current_state"`
	Point5State                string    `json:"point_5_state"`
	Point5DependencyState      string    `json:"point_5_dependency_state"`
	Point6State                string    `json:"point_6_state"`
	Point6ClosureState         string    `json:"point_6_closure_state"`
	Point6ClosurePrerequisites string    `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariants    string    `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState    string    `json:"point_6_proof_surface_state"`
	Point6PassRuleState        string    `json:"point_6_pass_rule_state"`
	Point6PassAllowed          bool      `json:"point_6_pass_allowed"`
	Val0State                  string    `json:"val_0_state"`
	Point7State                string    `json:"point_7_state"`
	VerifierContractState      string    `json:"verifier_contract_state"`
	ProofEnvelopeState         string    `json:"proof_envelope_state"`
	VerificationScopeState     string    `json:"verification_scope_state"`
	SchemaCompatibilityState   string    `json:"schema_compatibility_state"`
	TrustRootIssuerState       string    `json:"trust_root_issuer_state"`
	DiagnosticsState           string    `json:"diagnostics_state"`
	OutputBoundaryState        string    `json:"output_boundary_state"`
	SupportedProfiles          []string  `json:"supported_profiles,omitempty"`
	SupportedModes             []string  `json:"supported_modes,omitempty"`
	SupportedScopeClasses      []string  `json:"supported_scope_classes,omitempty"`
	WhyPoint7NotPass           []string  `json:"why_point_7_not_pass,omitempty"`
	SurfaceRefs                []string  `json:"surface_refs,omitempty"`
	EvidenceRefs               []string  `json:"evidence_refs,omitempty"`
	Limitations                []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer       string    `json:"projection_disclaimer"`
	IntegrationSummary         []string  `json:"integration_summary,omitempty"`
}

func verifierEcosystemVal0AllSurfaceRefs() []string {
	return operability.VerifierEcosystemVal0ProofSurfaceRefs()
}

func verifierEcosystemVal0ProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_verifier_discipline_foundation advisory_projection"
}

func verifierEcosystemVal0EvidenceRefs() []string {
	refs := []string{"point6_integrated_closure", "verifier_discipline_foundation"}
	for _, evidence := range operability.VerifierEcosystemVal0VerifierEvidence() {
		if evidence.EvidenceID != "" {
			refs = append(refs, evidence.EvidenceID)
		}
	}
	return refs
}

func buildVerifierEcosystemVal0DependencySnapshot() operability.VerifierEcosystemVal0DependencySnapshot {
	valE := buildReferenceArchitectureValEProofs()
	return operability.VerifierEcosystemVal0DependencySnapshot{
		Point5State:                    valE.Point5State,
		Point5DependencyState:          valE.Point5DependencyState,
		Point6State:                    valE.Point6State,
		Point6ClosureState:             valE.ValEState,
		Point6ClosurePrerequisiteState: valE.ClosurePrerequisiteState,
		Point6ClosureInvariantState:    valE.ClosureInvariantState,
		Point6ProofSurfaceState:        valE.ProofSurfaceState,
		Point6PassRuleState:            valE.PassRuleState,
		Point6PassAllowed:              valE.Point6PassAllowed,
	}
}

func buildVerifierEcosystemVal0SharedStates() (
	operability.VerifierEcosystemVal0DependencySnapshot,
	operability.VerifierEcosystemVal0VerifierContract,
	operability.VerifierEcosystemVal0ProofEnvelope,
	operability.VerifierEcosystemVal0VerificationScopeCatalog,
	operability.VerifierEcosystemVal0SchemaCompatibilityBaseline,
	operability.VerifierEcosystemVal0TrustIssuerDiscipline,
	operability.VerifierEcosystemVal0DiagnosticsModel,
	operability.VerifierEcosystemVal0OutputBoundaryCollection,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
	string,
) {
	dependency := buildVerifierEcosystemVal0DependencySnapshot()
	contract := operability.VerifierEcosystemVal0VerifierContractModel()
	envelope := operability.VerifierEcosystemVal0ProofEnvelopeModel()
	scopeCatalog := operability.VerifierEcosystemVal0VerificationScopeCatalogModel()
	compatibility := operability.VerifierEcosystemVal0SchemaCompatibilityBaselineModel()
	trust := operability.VerifierEcosystemVal0TrustIssuerDisciplineModel()
	diagnostics := operability.VerifierEcosystemVal0DiagnosticsCatalogModel()
	outputBoundaries := operability.VerifierEcosystemVal0OutputBoundaryCollectionModel()

	contractState := operability.EvaluateVerifierEcosystemVal0VerifierContractState(contract)
	envelopeState := operability.EvaluateVerifierEcosystemVal0ProofEnvelopeState(envelope)
	scopeState := operability.EvaluateVerifierEcosystemVal0VerificationScopeState(scopeCatalog)
	compatibilityState := operability.EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(compatibility)
	trustState := operability.EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(trust)
	diagnosticsState := operability.EvaluateVerifierEcosystemVal0DiagnosticsState(diagnostics)
	outputBoundaryState := operability.EvaluateVerifierEcosystemVal0OutputBoundaryState(outputBoundaries)
	val0State := operability.EvaluateVerifierEcosystemVal0State(
		dependency,
		contractState,
		envelopeState,
		scopeState,
		compatibilityState,
		trustState,
		diagnosticsState,
		outputBoundaryState,
	)
	return dependency, contract, envelope, scopeCatalog, compatibility, trust, diagnostics, outputBoundaries, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, val0State
}

func (s server) verifierEcosystemVal0ContractHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0Contract())
}

func (s server) verifierEcosystemVal0ProofEnvelopeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0ProofEnvelope())
}

func (s server) verifierEcosystemVal0VerificationScopeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0VerificationScope())
}

func (s server) verifierEcosystemVal0SchemaCompatibilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0SchemaCompatibility())
}

func (s server) verifierEcosystemVal0TrustIssuerHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0TrustIssuer())
}

func (s server) verifierEcosystemVal0DiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0Diagnostics())
}

func (s server) verifierEcosystemVal0OutputBoundariesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0OutputBoundaries())
}

func (s server) verifierEcosystemVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildVerifierEcosystemVal0Proofs())
}

func buildVerifierEcosystemVal0Contract() verifierEcosystemVal0ModelResponse {
	_, contract, _, _, _, _, _, _, contractState, _, _, _, _, _, _, _ := buildVerifierEcosystemVal0SharedStates()
	return verifierEcosystemVal0ModelResponse{
		SchemaVersion: verifierEcosystemVal0ContractSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  contractState,
		Model:         contract,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/val0/proof-envelope",
			"/v1/verifier-ecosystem/val0/proofs",
		},
		Limitations: []string{
			"Val 0 defines verifier contract discipline only and does not implement standalone CLI, SDK bindings, or public verifier hub execution.",
			"Verifier contracts remain advisory projections over the canonical evidence spine and do not approve deployment or create canonical truth.",
		},
	}
}

func buildVerifierEcosystemVal0ProofEnvelope() verifierEcosystemVal0ModelResponse {
	_, _, envelope, _, _, _, _, _, _, envelopeState, _, _, _, _, _, _ := buildVerifierEcosystemVal0SharedStates()
	return verifierEcosystemVal0ModelResponse{
		SchemaVersion: verifierEcosystemVal0EnvelopeSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  envelopeState,
		Model:         envelope,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/val0/schema-compatibility",
			"/v1/verifier-ecosystem/val0/proofs",
		},
		Limitations: []string{
			"Proof envelope verification remains bounded to digest, signature, issuer, trust-root, scope, freshness, lineage, compatibility, revocation, and supersession metadata.",
			"Envelope validation does not claim semantic truth beyond the declared proof scope.",
		},
	}
}

func buildVerifierEcosystemVal0VerificationScope() verifierEcosystemVal0ModelResponse {
	_, _, _, scopeCatalog, _, _, _, _, _, _, scopeState, _, _, _, _, _ := buildVerifierEcosystemVal0SharedStates()
	return verifierEcosystemVal0ModelResponse{
		SchemaVersion: verifierEcosystemVal0ScopeSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  scopeState,
		Model:         scopeCatalog,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/val0/output-boundaries",
			"/v1/verifier-ecosystem/val0/proofs",
		},
		Limitations: []string{
			"Scope classes remain bounded and preserve different output constraints for public, partner, auditor, internal, and restricted offline verification.",
			"Internal diagnostic scope must not be reused as public-safe output.",
		},
	}
}

func buildVerifierEcosystemVal0SchemaCompatibility() verifierEcosystemVal0ModelResponse {
	_, _, _, _, compatibility, _, _, _, _, _, _, compatibilityState, _, _, _, _ := buildVerifierEcosystemVal0SharedStates()
	return verifierEcosystemVal0ModelResponse{
		SchemaVersion: verifierEcosystemVal0CompatibilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  compatibilityState,
		Model:         compatibility,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/val0/trust-root-issuer",
			"/v1/verifier-ecosystem/val0/proofs",
		},
		Limitations: []string{
			"Unsupported, unknown, deprecated, and superseded verifier compatibility results remain explicit and fail closed.",
			"Compatibility diagnostics remain schema-governed and do not imply universal verifier support.",
		},
	}
}

func buildVerifierEcosystemVal0TrustIssuer() verifierEcosystemVal0ModelResponse {
	_, _, _, _, _, trust, _, _, _, _, _, _, trustState, _, _, _ := buildVerifierEcosystemVal0SharedStates()
	return verifierEcosystemVal0ModelResponse{
		SchemaVersion: verifierEcosystemVal0TrustIssuerSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  trustState,
		Model:         trust,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/val0/diagnostics",
			"/v1/verifier-ecosystem/val0/proofs",
		},
		Limitations: []string{
			"Trust-root and issuer discovery remain bounded, versioned, and scoped; no global all-instance key directory is created.",
			"Revoked, expired, unsupported, and unknown trust-root states fail closed.",
		},
	}
}

func buildVerifierEcosystemVal0Diagnostics() verifierEcosystemVal0ModelResponse {
	_, _, _, _, _, _, diagnostics, _, _, _, _, _, _, diagnosticsState, _, _ := buildVerifierEcosystemVal0SharedStates()
	return verifierEcosystemVal0ModelResponse{
		SchemaVersion: verifierEcosystemVal0DiagnosticsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  diagnosticsState,
		Model:         diagnostics,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/val0/output-boundaries",
			"/v1/verifier-ecosystem/val0/proofs",
		},
		Limitations: []string{
			"Verifier diagnostics preserve failure reasons, freshness, scope, and trust-material posture instead of suppressing invalid state.",
			"Diagnostics remain bounded to verifier scope and do not create canonical truth or certification.",
		},
	}
}

func buildVerifierEcosystemVal0OutputBoundaries() verifierEcosystemVal0ModelResponse {
	_, _, _, _, _, _, _, outputBoundaries, _, _, _, _, _, _, outputBoundaryState, _ := buildVerifierEcosystemVal0SharedStates()
	return verifierEcosystemVal0ModelResponse{
		SchemaVersion: verifierEcosystemVal0OutputBoundarySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  outputBoundaryState,
		Model:         outputBoundaries,
		RouteRefs: []string{
			"/v1/verifier-ecosystem/val0/contract",
			"/v1/verifier-ecosystem/val0/proofs",
		},
		Limitations: []string{
			"Output boundaries preserve public, partner, auditor, internal, and restricted offline differences without converting invalid artifacts into verified outputs.",
			"Public-safe outputs remain redacted and internal diagnostic outputs remain non-public.",
		},
	}
}

func buildVerifierEcosystemVal0Proofs() verifierEcosystemVal0ProofsResponse {
	dependency, contract, _, scopeCatalog, _, _, _, _, contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState, val0State := buildVerifierEcosystemVal0SharedStates()
	point7State := operability.EvaluateVerifierEcosystemPoint7State(val0State)
	limitations := []string{
		"Val 0 defines verifier discipline foundation only and does not implement standalone CLI execution, SDK bindings, public verifier hub, or third-party publisher onboarding.",
		"Točka 7 remains not complete because verifier tooling, ecosystem distribution, operational surfaces, and integrated closure remain for Val A through Val E.",
		"Verifier outputs remain advisory projections over the canonical evidence spine and must not become approval, certification, suppression, or mutation authority.",
	}
	currentState := operability.EvaluateVerifierEcosystemVal0ProofsState(
		val0State,
		point7State,
		contract.SupportedVerifierProfiles,
		contract.SupportedVerifierModes,
		verifierEcosystemVal0AllSurfaceRefs(),
		verifierEcosystemVal0EvidenceRefs(),
		limitations,
		verifierEcosystemVal0ProjectionDisclaimer(),
	)
	return verifierEcosystemVal0ProofsResponse{
		SchemaVersion:              verifierEcosystemVal0ProofsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               currentState,
		Point5State:                dependency.Point5State,
		Point5DependencyState:      dependency.Point5DependencyState,
		Point6State:                dependency.Point6State,
		Point6ClosureState:         dependency.Point6ClosureState,
		Point6ClosurePrerequisites: dependency.Point6ClosurePrerequisiteState,
		Point6ClosureInvariants:    dependency.Point6ClosureInvariantState,
		Point6ProofSurfaceState:    dependency.Point6ProofSurfaceState,
		Point6PassRuleState:        dependency.Point6PassRuleState,
		Point6PassAllowed:          dependency.Point6PassAllowed,
		Val0State:                  val0State,
		Point7State:                point7State,
		VerifierContractState:      contractState,
		ProofEnvelopeState:         envelopeState,
		VerificationScopeState:     scopeState,
		SchemaCompatibilityState:   compatibilityState,
		TrustRootIssuerState:       trustState,
		DiagnosticsState:           diagnosticsState,
		OutputBoundaryState:        outputBoundaryState,
		SupportedProfiles:          contract.SupportedVerifierProfiles,
		SupportedModes:             contract.SupportedVerifierModes,
		SupportedScopeClasses:      scopeCatalog.SupportedScopeClasses,
		WhyPoint7NotPass: []string{
			"Val 0 defines verifier discipline foundation only and cannot return point_7_pass.",
			"Val A through Val E remain required before Točka 7 integrated closure can exist.",
			"Verifier outputs remain bounded advisory projections and do not create canonical truth or deployment approval authority.",
		},
		SurfaceRefs:          verifierEcosystemVal0AllSurfaceRefs(),
		EvidenceRefs:         verifierEcosystemVal0EvidenceRefs(),
		Limitations:          limitations,
		ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val 0 establishes verifier contract, proof envelope, scope, compatibility, trust-root, diagnostics, and output-boundary discipline before later Točka 7 tooling waves.",
			"Verifier discipline remains fail-closed, trust-root-aware, scope-bounded, and advisory-only over the canonical execution, audit, and evidence spine.",
		},
	}
}
