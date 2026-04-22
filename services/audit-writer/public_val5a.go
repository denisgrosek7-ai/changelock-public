package main

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	publicHandoffSpecSchema            = "5a.public_handoff_spec.v1"
	publicProofVerificationSpecSchema  = "5a.public_proof_verification_spec.v1"
	publicValidationCertificateSchema  = "5a.public_validation_certificate_spec.v1"
	publicFederationExchangeSpecSchema = "5a.public_federation_exchange_spec.v1"
	publicVerifierProfilesSchema       = "5a.public_verifier_profiles.v1"
	publicOfflineGuideSchema           = "5a.public_offline_verification_guide.v1"
	publicExplainabilitySchema         = "5a.public_explainability_boundaries.v1"
	publicHandoffSampleSchema          = "5a.public_handoff_sample.v1"
	publicProofSampleSchema            = "5a.public_proof_verification_sample.v1"
	publicValidationSampleSchema       = "5a.public_validation_certificate_sample.v1"
	publicFederationSampleSchema       = "5a.public_federation_exchange_sample.v1"
	publicConformancePackSchema        = "5a.public_conformance_pack.v1"
	publicSchemaIndexSchema            = "5a.public_schema_index.v1"
	publicSchemaExportSchema           = "5a.public_schema_export.v1"
	publicVerifierReferencePackSchema  = "5a.public_verifier_reference_pack.v1"
)

type publicSpecCompatibility struct {
	CurrentVersion        string   `json:"current_version"`
	StabilityStatus       string   `json:"stability_status"`
	StatusClasses         []string `json:"status_classes,omitempty"`
	BackwardCompatibility []string `json:"backward_compatibility,omitempty"`
	BreakingChangePolicy  []string `json:"breaking_change_policy,omitempty"`
}

type publicSpecField struct {
	Field       string   `json:"field"`
	Required    bool     `json:"required"`
	Meaning     string   `json:"meaning"`
	Example     string   `json:"example,omitempty"`
	Limitations []string `json:"limitations,omitempty"`
}

type publicFailureSemantic struct {
	State               string `json:"state"`
	Meaning             string `json:"meaning"`
	VerifierExpectation string `json:"verifier_expectation"`
	LocalPolicyMeaning  string `json:"local_policy_meaning"`
}

type publicVerificationStep struct {
	Step             int      `json:"step"`
	Action           string   `json:"action"`
	ExpectedEvidence []string `json:"expected_evidence,omitempty"`
	FailureStates    []string `json:"failure_states,omitempty"`
}

type publicHandoffSpecResponse struct {
	SchemaVersion               string                   `json:"schema_version"`
	FormatID                    string                   `json:"format_id"`
	Compatibility               publicSpecCompatibility  `json:"compatibility"`
	ManifestFields              []publicSpecField        `json:"manifest_fields,omitempty"`
	ArtifactListSemantics       []string                 `json:"artifact_list_semantics,omitempty"`
	EvidenceReferenceSemantics  []string                 `json:"evidence_reference_semantics,omitempty"`
	SignatureSealingSemantics   []string                 `json:"signature_sealing_semantics,omitempty"`
	VerificationNarrativeFields []string                 `json:"verification_narrative_fields,omitempty"`
	FailureStates               []publicFailureSemantic  `json:"failure_states,omitempty"`
	QualityGateSemantics        []string                 `json:"quality_gate_semantics,omitempty"`
	ArchiveIntegrityFields      []string                 `json:"archive_integrity_fields,omitempty"`
	OfflineVerificationSteps    []publicVerificationStep `json:"offline_verification_steps,omitempty"`
	SampleBundlePaths           []string                 `json:"sample_bundle_paths,omitempty"`
	Limitations                 []string                 `json:"limitations,omitempty"`
}

type publicProofVerificationSpecResponse struct {
	SchemaVersion           string                   `json:"schema_version"`
	FormatID                string                   `json:"format_id"`
	Compatibility           publicSpecCompatibility  `json:"compatibility"`
	EnvelopeFields          []publicSpecField        `json:"envelope_fields,omitempty"`
	TrustAnchorExpectations []string                 `json:"trust_anchor_expectations,omitempty"`
	LocalVerificationSteps  []publicVerificationStep `json:"local_verification_steps,omitempty"`
	FreshnessSemantics      []string                 `json:"freshness_semantics,omitempty"`
	AdmissibilityModel      []string                 `json:"admissibility_model,omitempty"`
	RejectionReasons        []publicFailureSemantic  `json:"rejection_reasons,omitempty"`
	OfflineVerificationPath []string                 `json:"offline_verification_path,omitempty"`
	LocalPolicyOverride     []string                 `json:"local_policy_override,omitempty"`
	Limitations             []string                 `json:"limitations,omitempty"`
}

type publicValidationCertificateSpecResponse struct {
	SchemaVersion             string                  `json:"schema_version"`
	FormatID                  string                  `json:"format_id"`
	Compatibility             publicSpecCompatibility `json:"compatibility"`
	CertificateFields         []publicSpecField       `json:"certificate_fields,omitempty"`
	ExecutionProfileSemantics []string                `json:"execution_profile_semantics,omitempty"`
	PassFailSemantics         []string                `json:"pass_fail_semantics,omitempty"`
	LimitationSemantics       []string                `json:"limitation_semantics,omitempty"`
	SealReadySemantics        []string                `json:"seal_ready_semantics,omitempty"`
	CompatibilityRunSemantics []string                `json:"compatibility_run_semantics,omitempty"`
	FlakyRegressionIndicators []string                `json:"flaky_regression_indicators,omitempty"`
	AuthoritativeVsAdvisory   []string                `json:"authoritative_vs_advisory,omitempty"`
	FailureStates             []publicFailureSemantic `json:"failure_states,omitempty"`
	Limitations               []string                `json:"limitations,omitempty"`
}

type publicFederationExchangeSpecResponse struct {
	SchemaVersion              string                  `json:"schema_version"`
	FormatID                   string                  `json:"format_id"`
	Compatibility              publicSpecCompatibility `json:"compatibility"`
	EnvelopeFields             []publicSpecField       `json:"envelope_fields,omitempty"`
	FreshnessSemantics         []string                `json:"freshness_semantics,omitempty"`
	DisclosureProfiles         []string                `json:"disclosure_profiles,omitempty"`
	CompatibilitySignaling     []string                `json:"compatibility_signaling,omitempty"`
	StalePeerSemantics         []string                `json:"stale_peer_semantics,omitempty"`
	DivergenceDistrustModel    []string                `json:"divergence_distrust_model,omitempty"`
	LocalVerificationNarrative []string                `json:"local_verification_narrative,omitempty"`
	FailureStates              []publicFailureSemantic `json:"failure_states,omitempty"`
	NoGlobalAuthority          bool                    `json:"no_global_authority"`
	Limitations                []string                `json:"limitations,omitempty"`
}

type publicVerifierProfile struct {
	ProfileID            string   `json:"profile_id"`
	DisplayName          string   `json:"display_name"`
	SupportedSpecs       []string `json:"supported_specs,omitempty"`
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	EvidenceExpectations []string `json:"evidence_expectations,omitempty"`
	ConformanceMeaning   []string `json:"conformance_meaning,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type publicVerifierProfilesResponse struct {
	SchemaVersion       string                  `json:"schema_version"`
	Profiles            []publicVerifierProfile `json:"profiles,omitempty"`
	ConformanceLevels   []string                `json:"conformance_levels,omitempty"`
	CompatibilityPolicy []string                `json:"compatibility_policy,omitempty"`
	Limitations         []string                `json:"limitations,omitempty"`
}

type publicOfflineGuideResponse struct {
	SchemaVersion      string                   `json:"schema_version"`
	GuideVersion       string                   `json:"guide_version"`
	RequiredInputs     []string                 `json:"required_inputs,omitempty"`
	VerificationSteps  []publicVerificationStep `json:"verification_steps,omitempty"`
	FailureHandling    []string                 `json:"failure_handling,omitempty"`
	ConformanceTargets []string                 `json:"conformance_targets,omitempty"`
	ExampleArtifacts   []string                 `json:"example_artifacts,omitempty"`
	Limitations        []string                 `json:"limitations,omitempty"`
}

type publicExplainabilityResponse struct {
	SchemaVersion      string   `json:"schema_version"`
	PubliclyProves     []string `json:"publicly_proves,omitempty"`
	PubliclyInterprets []string `json:"publicly_interprets,omitempty"`
	LocalPolicyOwned   []string `json:"local_policy_owned,omitempty"`
	NotPubliclyClaimed []string `json:"not_publicly_claimed,omitempty"`
	Limitations        []string `json:"limitations,omitempty"`
}

type publicProofVerificationSample struct {
	CaseID         string                 `json:"case_id"`
	Proof          federatedProofResponse `json:"proof"`
	Verification   verificationResult     `json:"verification"`
	Decision       federatedTrustDecision `json:"decision"`
	Interpretation []string               `json:"interpretation,omitempty"`
}

type publicFederationExchangeSample struct {
	CaseID         string                 `json:"case_id"`
	Peer           federationPeer         `json:"peer"`
	Proof          federatedProofResponse `json:"proof"`
	Decision       federatedTrustDecision `json:"decision"`
	PolicyState    policyFederationState  `json:"policy_state"`
	Interpretation []string               `json:"interpretation,omitempty"`
}

type publicHandoffSampleResponse struct {
	SchemaVersion           string              `json:"schema_version"`
	SampleID                string              `json:"sample_id"`
	FormatID                string              `json:"format_id"`
	Sample                  handoffSealResponse `json:"sample"`
	ExpectedVerifierOutcome verificationResult  `json:"expected_verifier_outcome"`
	ConformanceNotes        []string            `json:"conformance_notes,omitempty"`
	Limitations             []string            `json:"limitations,omitempty"`
}

type publicProofVerificationSampleResponse struct {
	SchemaVersion    string                        `json:"schema_version"`
	SampleID         string                        `json:"sample_id"`
	FormatID         string                        `json:"format_id"`
	AcceptedExample  publicProofVerificationSample `json:"accepted_example"`
	RejectedExample  publicProofVerificationSample `json:"rejected_example"`
	ConformanceNotes []string                      `json:"conformance_notes,omitempty"`
	Limitations      []string                      `json:"limitations,omitempty"`
}

type publicValidationCertificateSampleResponse struct {
	SchemaVersion        string                `json:"schema_version"`
	SampleID             string                `json:"sample_id"`
	FormatID             string                `json:"format_id"`
	Sample               validationCertificate `json:"sample"`
	VerifierExpectations []string              `json:"verifier_expectations,omitempty"`
	Limitations          []string              `json:"limitations,omitempty"`
}

type publicFederationExchangeSampleResponse struct {
	SchemaVersion    string                         `json:"schema_version"`
	SampleID         string                         `json:"sample_id"`
	FormatID         string                         `json:"format_id"`
	ReadyExample     publicFederationExchangeSample `json:"ready_example"`
	StaleExample     publicFederationExchangeSample `json:"stale_example"`
	DivergedExample  publicFederationExchangeSample `json:"diverged_example"`
	ConformanceNotes []string                       `json:"conformance_notes,omitempty"`
	Limitations      []string                       `json:"limitations,omitempty"`
}

type publicConformanceAssertion struct {
	AssertionID   string   `json:"assertion_id"`
	ProfileID     string   `json:"profile_id"`
	SampleRef     string   `json:"sample_ref"`
	ExpectedState string   `json:"expected_state"`
	Checks        []string `json:"checks,omitempty"`
}

type publicConformancePackResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	PackVersion   string                       `json:"pack_version"`
	SampleRefs    []string                     `json:"sample_refs,omitempty"`
	Assertions    []publicConformanceAssertion `json:"assertions,omitempty"`
	Profiles      []string                     `json:"profiles,omitempty"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

type publicSchemaIndexItem struct {
	SchemaID            string   `json:"schema_id"`
	DisplayName         string   `json:"display_name"`
	FormatID            string   `json:"format_id"`
	ExportURI           string   `json:"export_uri"`
	SampleURI           string   `json:"sample_uri,omitempty"`
	StabilityStatus     string   `json:"stability_status"`
	ConformanceProfiles []string `json:"conformance_profiles,omitempty"`
}

type publicSchemaIndexResponse struct {
	SchemaVersion string                  `json:"schema_version"`
	Schemas       []publicSchemaIndexItem `json:"schemas,omitempty"`
	Limitations   []string                `json:"limitations,omitempty"`
}

type publicSchemaFieldDefinition struct {
	Path          string   `json:"path"`
	Type          string   `json:"type"`
	Required      bool     `json:"required"`
	Repeated      bool     `json:"repeated,omitempty"`
	SemanticClass string   `json:"semantic_class"`
	EnumValues    []string `json:"enum_values,omitempty"`
	Notes         []string `json:"notes,omitempty"`
}

type publicSchemaExportResponse struct {
	SchemaVersion    string                        `json:"schema_version"`
	ExportID         string                        `json:"export_id"`
	DisplayName      string                        `json:"display_name"`
	FormatID         string                        `json:"format_id"`
	Compatibility    publicSpecCompatibility       `json:"compatibility"`
	RequiredFields   []string                      `json:"required_fields,omitempty"`
	FieldDefinitions []publicSchemaFieldDefinition `json:"field_definitions,omitempty"`
	FailureStates    []string                      `json:"failure_states,omitempty"`
	SampleRef        string                        `json:"sample_ref,omitempty"`
	ConformanceRefs  []string                      `json:"conformance_refs,omitempty"`
	Limitations      []string                      `json:"limitations,omitempty"`
}

type publicVerifierReplayInput struct {
	InputID         string   `json:"input_id"`
	ProfileID       string   `json:"profile_id"`
	SampleRef       string   `json:"sample_ref"`
	InputType       string   `json:"input_type"`
	MediaType       string   `json:"media_type"`
	PayloadEncoding string   `json:"payload_encoding"`
	Payload         string   `json:"payload"`
	ExpectedState   string   `json:"expected_state"`
	ExpectedChecks  []string `json:"expected_checks,omitempty"`
	Notes           []string `json:"notes,omitempty"`
}

type publicVerifierReferencePackResponse struct {
	SchemaVersion string                      `json:"schema_version"`
	PackID        string                      `json:"pack_id"`
	Profiles      []string                    `json:"profiles,omitempty"`
	ReplayInputs  []publicVerifierReplayInput `json:"replay_inputs,omitempty"`
	UsageNotes    []string                    `json:"usage_notes,omitempty"`
	Limitations   []string                    `json:"limitations,omitempty"`
}

func (s server) publicHandoffSpecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicHandoffSpec())
}

func (s server) publicProofVerificationSpecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicProofVerificationSpec())
}

func (s server) publicValidationCertificateSpecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicValidationCertificateSpec())
}

func (s server) publicFederationExchangeSpecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicFederationExchangeSpec())
}

func (s server) publicVerifierProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicVerifierProfiles())
}

func (s server) publicOfflineGuideHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicOfflineGuide())
}

func (s server) publicExplainabilityBoundariesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicExplainabilityBoundaries())
}

func (s server) publicHandoffSampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicHandoffSample())
}

func (s server) publicProofVerificationSampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicProofVerificationSample())
}

func (s server) publicValidationCertificateSampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicValidationCertificateSample())
}

func (s server) publicFederationExchangeSampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicFederationExchangeSample())
}

func (s server) publicConformancePackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicConformancePack())
}

func (s server) publicSchemaIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicSchemaIndex())
}

func (s server) publicSchemaExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	exportID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/public/schemas/"))
	response, ok := buildPublicSchemaExport(exportID)
	if !ok {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "public schema not found"})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicVerifierReferencePackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	response, err := buildPublicVerifierReferencePack()
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildPublicHandoffSpec() publicHandoffSpecResponse {
	return publicHandoffSpecResponse{
		SchemaVersion: publicHandoffSpecSchema,
		FormatID:      federationProofTypeHandoff,
		Compatibility: buildPublicSpecCompatibility(publicHandoffSpecSchema),
		ManifestFields: []publicSpecField{
			{Field: "package_id", Required: true, Meaning: "Stable bundle identity used across manifest, archive, and verification surfaces.", Example: "HOF-SAMPLE"},
			{Field: "artifacts[].path", Required: true, Meaning: "Archive-relative artifact path covered by manifest hashing and verifier replay.", Example: "manifest.json"},
			{Field: "artifacts[].sha256", Required: true, Meaning: "Expected artifact digest that offline verifiers recompute locally.", Example: "sha256:..."},
			{Field: "readback_refs", Required: false, Meaning: "Bounded references to readback or evidence surfaces included for interpretation, not for remote authority.", Example: "/v1/handoff/HOF-SAMPLE/manifest"},
			{Field: "verification_uri", Required: false, Meaning: "Convenience URI for local product-native verification; offline verification remains authoritative.", Example: "/v1/handoff/HOF-SAMPLE/verification"},
		},
		ArtifactListSemantics: []string{
			"Artifact ordering is not authoritative; verifiers compare paths and digests rather than preserving source archive order.",
			"Every artifact entry declares verifier-relevant path and digest semantics; advisory-only payloads must remain explicitly marked in the bundle.",
		},
		EvidenceReferenceSemantics: []string{
			"Evidence references are pointers for explanation and drill-down, not implicit proof of remote retrievability.",
			"Public handoff interpretation must tolerate sealed bundles that include evidence refs but omit internal source systems.",
		},
		SignatureSealingSemantics: []string{
			"Manifest signatures, detached timestamp, and detached transparency records are verified independently against included public keys.",
			"Seal-ready bundles can be partially signed; verifier output must distinguish valid, partial, and invalid states.",
		},
		VerificationNarrativeFields: []string{
			"package_id",
			"manifest_hash",
			"verification.overall_status",
			"verification.signer_identities",
			"bundle.signature_count",
			"bundle.timestamp_status",
			"bundle.transparency_status",
		},
		FailureStates: []publicFailureSemantic{
			{State: handoffVerificationValid, Meaning: "Manifest, artifacts, signatures, timestamp, and transparency checks are locally valid.", VerifierExpectation: "Verifier can accept cryptographic integrity as locally verified.", LocalPolicyMeaning: "Local policy may still reject on scope, freshness, or disclosure grounds outside bundle integrity."},
			{State: handoffVerificationPartial, Meaning: "Some cryptographic or quality gates remain incomplete but enough structure exists for bounded interpretation.", VerifierExpectation: "Verifier must report which gates are incomplete and avoid presenting full acceptance.", LocalPolicyMeaning: "Requires stricter review or additional signatures before stronger reuse."},
			{State: handoffVerificationInvalid, Meaning: "Manifest or supporting cryptographic material failed local verification.", VerifierExpectation: "Verifier rejects integrity claims and records failure reasons.", LocalPolicyMeaning: "Bundle cannot be reused as accepted proof."},
		},
		QualityGateSemantics: []string{
			"Quality gate semantics remain public and bounded: deterministic assembly, signature verification, timestamp/transparency edges, and archive integrity must be independently visible.",
			"Failure-state semantics are stable and public even when internal assembly implementation evolves.",
		},
		ArchiveIntegrityFields: []string{
			"verify/manifest.sha256",
			"verify/public_keys.json",
			"signatures/signatures.json",
			"timestamp/manifest.timestamp.json",
			"transparency/manifest.transparency.json",
		},
		OfflineVerificationSteps: buildPublicHandoffOfflineSteps(),
		SampleBundlePaths: []string{
			"manifest.json",
			"verify/manifest.sha256",
			"verify/public_keys.json",
			"signatures/signatures.json",
			"timestamp/manifest.timestamp.json",
			"transparency/manifest.transparency.json",
		},
		Limitations: []string{
			"This public spec documents stable verifier-facing semantics, not every internal bundle assembly detail.",
			"Public handoff verification does not claim revocation, remote service health, or perpetual signer validity beyond the evidence present in the bundle.",
		},
	}
}

func buildPublicProofVerificationSpec() publicProofVerificationSpecResponse {
	return publicProofVerificationSpecResponse{
		SchemaVersion: publicProofVerificationSpecSchema,
		FormatID:      "trust_proof_envelope",
		Compatibility: buildPublicSpecCompatibility(publicProofVerificationSpecSchema),
		EnvelopeFields: []publicSpecField{
			{Field: "peer_id", Required: true, Meaning: "Registered trust peer identity evaluated by the local verifier.", Example: "partner-west"},
			{Field: "proof_type", Required: true, Meaning: "Public proof format identifier understood by the verifier profile.", Example: federationProofTypeHandoff},
			{Field: "manifest_hash", Required: true, Meaning: "Digest that binds the proof envelope to the sealed bundle or referenced artifact set.", Example: "sha256:..."},
			{Field: "scope", Required: true, Meaning: "Requested trust scope used for local admissibility and audience checks.", Example: "{\"tenant_id\":\"acme\",\"environment\":\"prod\",\"audience\":\"auditor_safe\"}"},
			{Field: "freshness", Required: true, Meaning: "Issued-at and valid-until semantics interpreted relative to local freshness policy.", Example: "{\"issued_at\":\"2026-04-21T08:00:00Z\",\"valid_until\":\"2026-04-21T09:00:00Z\"}"},
		},
		TrustAnchorExpectations: []string{
			"Verification expects locally registered trust anchors or peer public keys; remote proof material does not register itself.",
			"Verifier profiles must document whether they accept direct public-key trust, anchored peer registration, or auditor-curated trust roots.",
		},
		LocalVerificationSteps: []publicVerificationStep{
			{Step: 1, Action: "Verify manifest hash and sealed bundle cryptographic material locally.", ExpectedEvidence: []string{"manifest.json", "verify/public_keys.json", "signatures/signatures.json"}, FailureStates: []string{handoffVerificationInvalid}},
			{Step: 2, Action: "Check peer identity, trust anchors, disclosure mode, and accepted audiences against local policy.", ExpectedEvidence: []string{"peer registration metadata", "local admissibility policy"}, FailureStates: []string{federationDecisionRejectedUntrustedPeer, federationDecisionRejectedPolicyConflict}},
			{Step: 3, Action: "Evaluate scope and freshness before accepting the proof for any downstream recommendation or reuse.", ExpectedEvidence: []string{"proof scope", "proof freshness"}, FailureStates: []string{federationDecisionRejectedScopeMismatch, federationDecisionRejectedStale}},
		},
		FreshnessSemantics: []string{
			"Freshness is bounded by local window expectations and can reject otherwise cryptographically valid proofs.",
			"Public proof verification must report stale state explicitly rather than silently degrading to acceptance.",
		},
		AdmissibilityModel: []string{
			"Local admissibility is authoritative and can reject a remote proof after cryptographic verification succeeds.",
			"Admissibility covers scope, disclosure profile, partner trust state, and local override posture.",
		},
		RejectionReasons: []publicFailureSemantic{
			{State: federationDecisionRejectedUnverifiable, Meaning: "Bundle or envelope failed local cryptographic verification.", VerifierExpectation: "Report cryptographic failure and stop acceptance flow.", LocalPolicyMeaning: "No proof reuse allowed."},
			{State: federationDecisionRejectedScopeMismatch, Meaning: "Requested scope or audience did not match local expectations.", VerifierExpectation: "Report mismatch in requested scope fields.", LocalPolicyMeaning: "Proof stays locally rejected even if integrity is valid."},
			{State: federationDecisionRejectedPolicyConflict, Meaning: "Local policy or disclosure constraints conflict with remote proof usage.", VerifierExpectation: "Explain which bounded policy gate blocked acceptance.", LocalPolicyMeaning: "Requires policy adjustment or narrower disclosure."},
			{State: federationDecisionRejectedStale, Meaning: "Proof or peer freshness exceeded local admissibility window.", VerifierExpectation: "Flag stale state with explicit freshness context.", LocalPolicyMeaning: "Revalidation or renewed proof is required."},
			{State: federationDecisionRejectedUntrustedPeer, Meaning: "Peer registration, trust anchors, or capability claims are not locally trusted.", VerifierExpectation: "Explain local trust-anchor failure or missing peer registration.", LocalPolicyMeaning: "Remote peer is not accepted for this local verifier."},
		},
		OfflineVerificationPath: []string{
			"Offline verification first replays sealed handoff verification, then applies local trust-anchor and admissibility checks without requiring remote connectivity.",
			"Public verifier implementations may be minimal, but must still surface cryptographic failures, freshness state, and local-policy rejections distinctly.",
		},
		LocalPolicyOverride: []string{
			"Remote proof validity never bypasses local policy overrides, local distrust, or local disclosure exclusions.",
			"Acceptance and rejection narratives must distinguish proof validity from local admissibility decisions.",
		},
		Limitations: []string{
			"This public proof verification spec defines a verifier-facing envelope and decision model; it does not establish a new global trust authority or shared control plane.",
		},
	}
}

func buildPublicValidationCertificateSpec() publicValidationCertificateSpecResponse {
	return publicValidationCertificateSpecResponse{
		SchemaVersion: publicValidationCertificateSchema,
		FormatID:      "validation_certificate",
		Compatibility: buildPublicSpecCompatibility(publicValidationCertificateSchema),
		CertificateFields: []publicSpecField{
			{Field: "certificate_id", Required: true, Meaning: "Stable validation certificate identity for cross-reference and sealing.", Example: "VALCERT-SAMPLE"},
			{Field: "run_id", Required: true, Meaning: "Execution run identity for replay, troubleshooting, and scenario grouping.", Example: "VALRUN-SAMPLE"},
			{Field: "scenario_set[]", Required: true, Meaning: "Scenario identities included in the validation certificate.", Example: validationScenarioSafeRelease},
			{Field: "environment_summary", Required: true, Meaning: "Execution profile and bounded deployment context used during the run.", Example: "{\"mode\":\"policy_dry_run\",\"environment\":\"prod\"}"},
			{Field: "overall_status", Required: true, Meaning: "Rollup status derived from individual verdicts.", Example: validationStatusPass},
			{Field: "seal_ready", Required: true, Meaning: "Indicates whether outputs are suitable for sealed export or handoff inclusion.", Example: "true"},
		},
		ExecutionProfileSemantics: []string{
			"Execution profile records mode, namespace, environment tag, and isolation class so third parties can interpret whether the run was production-like, compatibility-only, or simulation-derived.",
			"Scenario identity and version remain explicit because validation meaning changes with scenario revisions.",
		},
		PassFailSemantics: []string{
			validationStatusPass + " means expected control behavior matched observed outcomes within bounded limitations.",
			validationStatusPartial + " means some but not all expected controls or thresholds were met.",
			validationStatusFail + " means at least one required control behavior or threshold was not met.",
			validationStatusFlaky + " means recent runs did not converge on one stable outcome and the scenario requires caution.",
			validationStatusUnknown + " means evidence was insufficient for authoritative interpretation.",
		},
		LimitationSemantics: []string{
			"Limitations are first-class output and must remain visible in public certificates.",
			"Simulation-derived or compatibility runs must remain clearly marked and never presented as production mutation proof.",
		},
		SealReadySemantics: []string{
			"seal_ready indicates export readiness into sealed handoff or proof exchange workflows, not an unconditional production approval.",
		},
		CompatibilityRunSemantics: []string{
			"Compatibility validation is simulation-derived and predicts behavior under changed platform or policy assumptions.",
			"Compatibility outputs remain informative but are not retroactive proof of historical production behavior.",
		},
		FlakyRegressionIndicators: []string{
			"Scenario-level flaky indicators remain public because verifier consumers need to distinguish stable from unstable evidence.",
			"Regression interpretation requires both current verdict and recent-history convergence context.",
		},
		AuthoritativeVsAdvisory: []string{
			"Validation certificates are authoritative about the bounded run they describe, not about all future runtime states.",
			"Guidance derived from validation remains advisory when external deployment or runtime posture changes are outside the run scope.",
		},
		FailureStates: []publicFailureSemantic{
			{State: validationStatusFail, Meaning: "One or more required scenario controls did not meet expectations.", VerifierExpectation: "Treat certificate as failing for the bounded run.", LocalPolicyMeaning: "Policy may block promotion, require remediation, or demand rerun."},
			{State: validationStatusFlaky, Meaning: "Recent scenario history did not converge to one stable outcome.", VerifierExpectation: "Show instability rather than collapsing to pass or fail.", LocalPolicyMeaning: "Requires caution, rerun, or tighter controls before relying on the result."},
			{State: validationStatusUnknown, Meaning: "Evidence was insufficient for authoritative interpretation.", VerifierExpectation: "Do not overstate confidence.", LocalPolicyMeaning: "Local policy may reject or downgrade this certificate."},
		},
		Limitations: []string{
			"Public validation certificates remain bounded by scenario set, execution profile, and limitations; they do not act as perpetual deployment guarantees.",
		},
	}
}

func buildPublicFederationExchangeSpec() publicFederationExchangeSpecResponse {
	return publicFederationExchangeSpecResponse{
		SchemaVersion: publicFederationExchangeSpecSchema,
		FormatID:      "federation_peer_proof_exchange",
		Compatibility: buildPublicSpecCompatibility(publicFederationExchangeSpecSchema),
		EnvelopeFields: []publicSpecField{
			{Field: "peer_id", Required: true, Meaning: "Remote peer identity used for local acceptance and distrust handling.", Example: "supplier-west"},
			{Field: "redaction_profile", Required: true, Meaning: "Disclosure-minimized profile applied to the exported proof.", Example: "sealed_proof_only"},
			{Field: "freshness", Required: true, Meaning: "Issued-at and validity window semantics for cross-org proof exchange.", Example: "{\"freshness_minutes\":60}"},
			{Field: "status", Required: true, Meaning: "Bounded proof exchange state, distinct from local decision outcome.", Example: federationProofStatusReady},
			{Field: "limitations", Required: false, Meaning: "Peer-specific or exchange-specific limitations that must remain visible to consumers.", Example: "peer freshness window is shorter than default"},
		},
		FreshnessSemantics: []string{
			"Cross-org proof freshness is explicit and locally evaluated; stale peers or stale proofs remain visible and do not silently downgrade to accepted state.",
			"Freshness semantics are part of the public format because partner verifiers need repeatable expiry behavior.",
		},
		DisclosureProfiles: []string{
			"sealed_proof_only",
			"auditor_safe",
			"customer_safe",
		},
		CompatibilitySignaling: []string{
			"Peer proof exchange must surface compatibility state, accepted audiences, and proof-type support without requiring direct access to remote internal systems.",
			"Public exchange format supports bounded compatibility signaling; it does not claim universal cross-version interoperability.",
		},
		StalePeerSemantics: []string{
			"Stale peers remain locally visible and can be rejected even if earlier proofs were cryptographically valid.",
			"Peer freshness is a first-class part of the public exchange narrative.",
		},
		DivergenceDistrustModel: []string{
			"Local divergence, local override, and local distrust remain authoritative and publicly documented in verifier guidance.",
			"Remote proof exchange does not override local federation policy root or local exception posture.",
		},
		LocalVerificationNarrative: []string{
			"Federation proof exchange remains a local verification process combining sealed handoff verification with peer registration, freshness, scope, and disclosure checks.",
			"Public exchange consumers must distinguish ready proof state from accepted local trust decision state.",
		},
		FailureStates: []publicFailureSemantic{
			{State: federationProofStatusRejected, Meaning: "Proof exchange item was rejected by local verification or policy.", VerifierExpectation: "Show whether cryptographic, freshness, scope, or trust-anchor issues caused rejection.", LocalPolicyMeaning: "No local reuse or acceptance."},
			{State: federationPeerStatusStale, Meaning: "Peer freshness is stale under local policy.", VerifierExpectation: "Treat stale peer state as visible failure context, not hidden metadata.", LocalPolicyMeaning: "May block acceptance or force narrower use."},
			{State: federationSyncStatusDiverged, Meaning: "Local and remote policy states diverged.", VerifierExpectation: "Explain divergence rather than flattening it into generic incompatibility.", LocalPolicyMeaning: "Local policy remains authoritative until resolved."},
		},
		NoGlobalAuthority: true,
		Limitations: []string{
			"This public federation exchange spec documents bounded exchange semantics and local verification expectations; it does not establish a global shared authority or mandatory shared transparency service.",
		},
	}
}

func buildPublicVerifierProfiles() publicVerifierProfilesResponse {
	return publicVerifierProfilesResponse{
		SchemaVersion: publicVerifierProfilesSchema,
		Profiles: []publicVerifierProfile{
			{
				ProfileID:      "minimal_verifier",
				DisplayName:    "Minimal verifier",
				SupportedSpecs: []string{publicHandoffSpecSchema, publicProofVerificationSpecSchema},
				RequiredCapabilities: []string{
					"bundle hash verification",
					"signature verification",
					"freshness evaluation",
					"explicit accept/reject reporting",
				},
				EvidenceExpectations: []string{
					"sealed bundle or proof envelope",
					"local trust anchors",
				},
				ConformanceMeaning: []string{
					"Can independently verify bundle integrity and render bounded proof acceptance or rejection.",
				},
				Limitations: []string{
					"May not interpret validation certificates or partner-specific governance overlays.",
				},
			},
			{
				ProfileID:      "full_verifier",
				DisplayName:    "Full verifier",
				SupportedSpecs: []string{publicHandoffSpecSchema, publicProofVerificationSpecSchema, publicValidationCertificateSchema, publicFederationExchangeSpecSchema},
				RequiredCapabilities: []string{
					"minimal verifier capabilities",
					"validation certificate interpretation",
					"federation divergence handling",
					"conformance reporting",
				},
				EvidenceExpectations: []string{
					"sealed bundle or proof envelope",
					"validation certificate",
					"peer metadata and local policy",
				},
				ConformanceMeaning: []string{
					"Can verify, interpret, and report all public 5A verification surfaces with explicit limitation handling.",
				},
			},
			{
				ProfileID:      "auditor",
				DisplayName:    "Auditor verifier",
				SupportedSpecs: []string{publicHandoffSpecSchema, publicValidationCertificateSchema, publicFederationExchangeSpecSchema},
				RequiredCapabilities: []string{
					"offline bundle verification",
					"limitation reporting",
					"historical evidence interpretation",
				},
				EvidenceExpectations: []string{
					"sealed handoff archive",
					"validation certificate",
					"documented local admissibility notes",
				},
				ConformanceMeaning: []string{
					"Can independently assess archived evidence without becoming the system of record for local policy.",
				},
			},
			{
				ProfileID:      "partner_verifier",
				DisplayName:    "Partner verifier",
				SupportedSpecs: []string{publicProofVerificationSpecSchema, publicFederationExchangeSpecSchema},
				RequiredCapabilities: []string{
					"partner proof acceptance",
					"disclosure-profile handling",
					"local override and distrust support",
				},
				EvidenceExpectations: []string{
					"proof envelope",
					"peer trust anchor registration",
					"local policy profile",
				},
				ConformanceMeaning: []string{
					"Can participate in bounded partner proof exchange without inheriting remote authority.",
				},
			},
		},
		ConformanceLevels: []string{
			"minimal",
			"full",
			"auditor",
			"partner",
		},
		CompatibilityPolicy: []string{
			"Conformance is defined per public schema_version and profile, not as a perpetual blanket claim.",
			"Additive optional fields may extend a public profile without invalidating earlier compatible implementations.",
			"Breaking semantic changes require a new public schema_version and updated conformance target.",
		},
		Limitations: []string{
			"Conformance profiles define verifier capability expectations; they are not trust marks, badges, or legal certifications by themselves.",
		},
	}
}

func buildPublicOfflineGuide() publicOfflineGuideResponse {
	return publicOfflineGuideResponse{
		SchemaVersion: publicOfflineGuideSchema,
		GuideVersion:  "5a.offline_guide.v1",
		RequiredInputs: []string{
			"sealed handoff bundle or public proof envelope",
			"local trust anchors or peer registration data",
			"local admissibility and disclosure policy",
			"public spec version and conformance target",
		},
		VerificationSteps: append(
			buildPublicHandoffOfflineSteps(),
			publicVerificationStep{
				Step:             6,
				Action:           "Evaluate local admissibility, freshness, and disclosure boundaries before accepting the proof.",
				ExpectedEvidence: []string{"local policy", "peer freshness", "requested scope"},
				FailureStates:    []string{federationDecisionRejectedPolicyConflict, federationDecisionRejectedScopeMismatch, federationDecisionRejectedStale},
			},
		),
		FailureHandling: []string{
			"Report cryptographic verification failure separately from local-policy rejection.",
			"Do not collapse stale, unverifiable, and policy-conflict outcomes into one generic failure class.",
			"Preserve limitation and uncertainty labels in offline reports.",
		},
		ConformanceTargets: []string{
			"minimal_verifier",
			"full_verifier",
			"auditor",
			"partner_verifier",
		},
		ExampleArtifacts: []string{
			"manifest.json",
			"verify/public_keys.json",
			"validation certificate JSON",
			"federation proof envelope JSON",
		},
		Limitations: []string{
			"Offline verification guide documents repeatable local verification steps; it does not replace local policy ownership or remote operational due diligence.",
		},
	}
}

func buildPublicExplainabilityBoundaries() publicExplainabilityResponse {
	return publicExplainabilityResponse{
		SchemaVersion: publicExplainabilitySchema,
		PubliclyProves: []string{
			"bundle integrity and manifest consistency",
			"signature, timestamp, and transparency verification outcomes",
			"validation run identity, scenario set, and bounded run result",
			"proof freshness and explicit peer stale/divergence state when present",
		},
		PubliclyInterprets: []string{
			"what each failure state means for a verifier consumer",
			"which public fields are stable, optional, or limitation-bearing",
			"whether a proof is cryptographically valid, stale, partial, or rejected",
		},
		LocalPolicyOwned: []string{
			"final admissibility and reuse decisions",
			"disclosure scope and audience acceptance",
			"partner distrust, override, or exception posture",
			"promotion, remediation, and runtime action decisions driven by local governance",
		},
		NotPubliclyClaimed: []string{
			"global authority over remote systems",
			"universal proof acceptance across organizations",
			"perpetual validity of certificates or trust marks",
			"black-box scores without drill-down and limitation context",
		},
		Limitations: []string{
			"Public explainability boundaries document what outside verifiers can infer from public formats; they do not expose internal tenant data or replace local governance context.",
		},
	}
}

func buildPublicHandoffSample() publicHandoffSampleResponse {
	sample := publicSampleHandoff()
	return publicHandoffSampleResponse{
		SchemaVersion:           publicHandoffSampleSchema,
		SampleID:                "handoff_bundle_sample",
		FormatID:                federationProofTypeHandoff,
		Sample:                  sample,
		ExpectedVerifierOutcome: sample.Verification,
		ConformanceNotes: []string{
			"Minimal verifiers must reproduce the sample verification outcome from the bundle and public keys alone.",
			"Full and auditor verifier profiles should preserve seal status, signer identities, and limitation labels.",
		},
		Limitations: []string{
			"Sample payload is a public reference example, not a claim about any live tenant, package, or signer.",
		},
	}
}

func buildPublicProofVerificationSample() publicProofVerificationSampleResponse {
	accepted := publicAcceptedProofSample()
	rejected := publicRejectedProofSample()
	return publicProofVerificationSampleResponse{
		SchemaVersion:   publicProofSampleSchema,
		SampleID:        "proof_verification_cases",
		FormatID:        "trust_proof_envelope",
		AcceptedExample: accepted,
		RejectedExample: rejected,
		ConformanceNotes: []string{
			"Partner and full verifier profiles must distinguish cryptographic validity from local admissibility.",
			"Rejected examples must preserve explicit rejection reasons instead of collapsing them into a generic invalid state.",
		},
		Limitations: []string{
			"These public proof samples exercise bounded local-policy outcomes; they are reference cases rather than a complete interoperability matrix.",
		},
	}
}

func buildPublicValidationCertificateSample() publicValidationCertificateSampleResponse {
	sample := publicSampleValidationCertificate()
	return publicValidationCertificateSampleResponse{
		SchemaVersion: publicValidationSampleSchema,
		SampleID:      "validation_certificate_sample",
		FormatID:      "validation_certificate",
		Sample:        sample,
		VerifierExpectations: []string{
			"Verifier should preserve scenario identity, scenario version, execution profile, and limitations when rendering the certificate.",
			"Verifier should distinguish pass, flaky, and advisory semantics without converting the certificate into a perpetual guarantee.",
		},
		Limitations: []string{
			"Sample validation certificate represents one bounded run and one bounded execution profile.",
		},
	}
}

func buildPublicFederationExchangeSample() publicFederationExchangeSampleResponse {
	ready := publicReadyFederationSample()
	stale := publicStaleFederationSample()
	diverged := publicDivergedFederationSample()
	return publicFederationExchangeSampleResponse{
		SchemaVersion:   publicFederationSampleSchema,
		SampleID:        "federation_exchange_cases",
		FormatID:        "federation_peer_proof_exchange",
		ReadyExample:    ready,
		StaleExample:    stale,
		DivergedExample: diverged,
		ConformanceNotes: []string{
			"Federation-capable verifier profiles must keep stale and diverged cases distinct from cryptographic invalidity.",
			"Ready exchange state does not imply global trust authority; local decision and policy state remain explicit.",
		},
		Limitations: []string{
			"Public federation samples show bounded peer exchange semantics only and omit tenant-internal evidence stores.",
		},
	}
}

func buildPublicConformancePack() publicConformancePackResponse {
	return publicConformancePackResponse{
		SchemaVersion: publicConformancePackSchema,
		PackVersion:   "5a.conformance_pack.v1",
		SampleRefs: []string{
			"/v1/public/samples/handoff",
			"/v1/public/samples/proof-verification",
			"/v1/public/samples/validation-certificate",
			"/v1/public/samples/federation-proof-exchange",
		},
		Assertions: []publicConformanceAssertion{
			{
				AssertionID:   "minimal-handoff-valid",
				ProfileID:     "minimal_verifier",
				SampleRef:     "/v1/public/samples/handoff",
				ExpectedState: handoffVerificationValid,
				Checks:        []string{"manifest_valid", "artifact_hashes_valid", "signatures_valid", "timestamp_valid", "transparency_valid"},
			},
			{
				AssertionID:   "partner-proof-accept",
				ProfileID:     "partner_verifier",
				SampleRef:     "/v1/public/samples/proof-verification",
				ExpectedState: federationDecisionAccepted,
				Checks:        []string{"local verification preserved", "decision reason retained", "freshness not stale"},
			},
			{
				AssertionID:   "partner-proof-reject-stale",
				ProfileID:     "partner_verifier",
				SampleRef:     "/v1/public/samples/proof-verification",
				ExpectedState: federationDecisionRejectedStale,
				Checks:        []string{"stale flagged explicitly", "cryptographic validity not confused with acceptance"},
			},
			{
				AssertionID:   "validation-certificate-pass",
				ProfileID:     "full_verifier",
				SampleRef:     "/v1/public/samples/validation-certificate",
				ExpectedState: validationStatusPass,
				Checks:        []string{"scenario_set preserved", "seal_ready preserved", "limitations preserved"},
			},
			{
				AssertionID:   "federation-divergence-visible",
				ProfileID:     "full_verifier",
				SampleRef:     "/v1/public/samples/federation-proof-exchange",
				ExpectedState: federationSyncStatusDiverged,
				Checks:        []string{"policy divergence visible", "local overrides visible", "no global authority implied"},
			},
		},
		Profiles: []string{"minimal_verifier", "full_verifier", "auditor", "partner_verifier"},
		Limitations: []string{
			"Conformance pack is a bounded public reference test pack; it does not certify full implementation parity beyond the published assertions.",
		},
	}
}

func buildPublicSchemaIndex() publicSchemaIndexResponse {
	return publicSchemaIndexResponse{
		SchemaVersion: publicSchemaIndexSchema,
		Schemas: []publicSchemaIndexItem{
			{SchemaID: "handoff", DisplayName: "Public sealed handoff schema", FormatID: federationProofTypeHandoff, ExportURI: "/v1/public/schemas/handoff", SampleURI: "/v1/public/samples/handoff", StabilityStatus: "stable", ConformanceProfiles: []string{"minimal_verifier", "full_verifier", "auditor"}},
			{SchemaID: "proof-verification", DisplayName: "Public proof verification schema", FormatID: "trust_proof_envelope", ExportURI: "/v1/public/schemas/proof-verification", SampleURI: "/v1/public/samples/proof-verification", StabilityStatus: "stable", ConformanceProfiles: []string{"minimal_verifier", "full_verifier", "partner_verifier"}},
			{SchemaID: "validation-certificate", DisplayName: "Public validation certificate schema", FormatID: "validation_certificate", ExportURI: "/v1/public/schemas/validation-certificate", SampleURI: "/v1/public/samples/validation-certificate", StabilityStatus: "stable", ConformanceProfiles: []string{"full_verifier", "auditor"}},
			{SchemaID: "federation-proof-exchange", DisplayName: "Public federation proof exchange schema", FormatID: "federation_peer_proof_exchange", ExportURI: "/v1/public/schemas/federation-proof-exchange", SampleURI: "/v1/public/samples/federation-proof-exchange", StabilityStatus: "stable", ConformanceProfiles: []string{"full_verifier", "partner_verifier"}},
		},
		Limitations: []string{
			"Schema index lists stable public verifier-facing formats only; internal event or storage formats remain out of scope.",
		},
	}
}

func buildPublicSchemaExport(exportID string) (publicSchemaExportResponse, bool) {
	switch strings.TrimSpace(exportID) {
	case "handoff":
		return publicSchemaExportResponse{
			SchemaVersion:  publicSchemaExportSchema,
			ExportID:       "handoff",
			DisplayName:    "Public sealed handoff schema export",
			FormatID:       federationProofTypeHandoff,
			Compatibility:  buildPublicSpecCompatibility(publicHandoffSpecSchema),
			RequiredFields: []string{"package_id", "package_type", "schema_version", "created_at", "scope", "redaction_profile", "artifacts", "evidence_refs", "root_hash"},
			FieldDefinitions: []publicSchemaFieldDefinition{
				{Path: "package_id", Type: "string", Required: true, SemanticClass: "identity"},
				{Path: "scope.audience", Type: "string", Required: true, SemanticClass: "disclosure_control", EnumValues: []string{incidentAudienceAuditorSafe, incidentAudienceCustomerSafe, incidentAudienceInternal}},
				{Path: "artifacts[].path", Type: "string", Required: true, Repeated: true, SemanticClass: "content_locator"},
				{Path: "artifacts[].sha256", Type: "string", Required: true, Repeated: true, SemanticClass: "integrity_digest"},
				{Path: "readback_refs[].resource_uri", Type: "string", Required: false, Repeated: true, SemanticClass: "bounded_drilldown"},
				{Path: "forensic_refs[].advisory_only", Type: "bool", Required: true, Repeated: true, SemanticClass: "advisory_flag"},
				{Path: "limitations[]", Type: "string", Required: false, Repeated: true, SemanticClass: "limitation"},
			},
			FailureStates:   []string{handoffVerificationValid, handoffVerificationPartial, handoffVerificationInvalid},
			SampleRef:       "/v1/public/samples/handoff",
			ConformanceRefs: []string{"/v1/public/conformance-pack", "/v1/public/verifier/reference-pack"},
			Limitations: []string{
				"Schema export is machine-readable field guidance, not a promise that every deployment exposes all optional readback or forensic refs.",
			},
		}, true
	case "proof-verification":
		return publicSchemaExportResponse{
			SchemaVersion:  publicSchemaExportSchema,
			ExportID:       "proof-verification",
			DisplayName:    "Public proof verification schema export",
			FormatID:       "trust_proof_envelope",
			Compatibility:  buildPublicSpecCompatibility(publicProofVerificationSpecSchema),
			RequiredFields: []string{"peer_id", "proof_type", "manifest_hash", "scope", "freshness"},
			FieldDefinitions: []publicSchemaFieldDefinition{
				{Path: "peer_id", Type: "string", Required: true, SemanticClass: "peer_identity"},
				{Path: "proof_type", Type: "string", Required: true, SemanticClass: "format_id", EnumValues: []string{federationProofTypeHandoff}},
				{Path: "scope.audience", Type: "string", Required: true, SemanticClass: "admissibility_scope", EnumValues: []string{incidentAudienceAuditorSafe, incidentAudienceCustomerSafe, incidentAudienceInternal}},
				{Path: "freshness.stale", Type: "bool", Required: true, SemanticClass: "freshness_state"},
				{Path: "status", Type: "string", Required: true, SemanticClass: "exchange_state", EnumValues: []string{federationProofStatusReady, federationProofStatusAccepted, federationProofStatusRejected}},
				{Path: "limitations[]", Type: "string", Required: false, Repeated: true, SemanticClass: "limitation"},
			},
			FailureStates:   []string{federationDecisionRejectedUnverifiable, federationDecisionRejectedScopeMismatch, federationDecisionRejectedPolicyConflict, federationDecisionRejectedStale, federationDecisionRejectedUntrustedPeer},
			SampleRef:       "/v1/public/samples/proof-verification",
			ConformanceRefs: []string{"/v1/public/conformance-pack", "/v1/public/verifier/reference-pack"},
			Limitations: []string{
				"Export covers verifier-facing envelope semantics only; local policy data remains implementation-owned.",
			},
		}, true
	case "validation-certificate":
		return publicSchemaExportResponse{
			SchemaVersion:  publicSchemaExportSchema,
			ExportID:       "validation-certificate",
			DisplayName:    "Public validation certificate schema export",
			FormatID:       "validation_certificate",
			Compatibility:  buildPublicSpecCompatibility(publicValidationCertificateSchema),
			RequiredFields: []string{"run_id", "certificate_id", "scope", "scenario_set", "issued_at", "overall_status", "scenario_results", "timing_summary", "environment_summary", "simulation_derived", "seal_ready"},
			FieldDefinitions: []publicSchemaFieldDefinition{
				{Path: "scenario_set[]", Type: "string", Required: true, Repeated: true, SemanticClass: "scenario_identity"},
				{Path: "scenario_results[].status", Type: "string", Required: true, Repeated: true, SemanticClass: "verdict_status", EnumValues: []string{validationStatusPass, validationStatusPartial, validationStatusFail, validationStatusFlaky, validationStatusUnknown}},
				{Path: "environment_summary.mode", Type: "string", Required: true, SemanticClass: "execution_profile", EnumValues: []string{validationModePolicyDryRun, validationModeRegression, validationModeCompatibility, validationModeControlledChaos}},
				{Path: "simulation_derived", Type: "bool", Required: true, SemanticClass: "simulation_flag"},
				{Path: "seal_ready", Type: "bool", Required: true, SemanticClass: "export_readiness"},
				{Path: "limitations[]", Type: "string", Required: false, Repeated: true, SemanticClass: "limitation"},
			},
			FailureStates:   []string{validationStatusFail, validationStatusFlaky, validationStatusUnknown},
			SampleRef:       "/v1/public/samples/validation-certificate",
			ConformanceRefs: []string{"/v1/public/conformance-pack", "/v1/public/verifier/reference-pack"},
			Limitations: []string{
				"Validation schema export preserves bounded run semantics and does not define organization-specific promotion policy.",
			},
		}, true
	case "federation-proof-exchange":
		return publicSchemaExportResponse{
			SchemaVersion:  publicSchemaExportSchema,
			ExportID:       "federation-proof-exchange",
			DisplayName:    "Public federation proof exchange schema export",
			FormatID:       "federation_peer_proof_exchange",
			Compatibility:  buildPublicSpecCompatibility(publicFederationExchangeSpecSchema),
			RequiredFields: []string{"peer_id", "redaction_profile", "freshness", "status"},
			FieldDefinitions: []publicSchemaFieldDefinition{
				{Path: "peer_id", Type: "string", Required: true, SemanticClass: "peer_identity"},
				{Path: "redaction_profile", Type: "string", Required: true, SemanticClass: "disclosure_profile", EnumValues: []string{"sealed_proof_only", "auditor_safe", "customer_safe"}},
				{Path: "freshness.stale", Type: "bool", Required: true, SemanticClass: "freshness_state"},
				{Path: "status", Type: "string", Required: true, SemanticClass: "exchange_status", EnumValues: []string{federationProofStatusReady, federationProofStatusAccepted, federationProofStatusRejected}},
				{Path: "policy_state.sync_status", Type: "string", Required: false, SemanticClass: "policy_sync_state", EnumValues: []string{federationSyncStatusLocalOnly, federationSyncStatusSynced, federationSyncStatusSyncedWithOverrides, federationSyncStatusDiverged, federationSyncStatusStale}},
				{Path: "limitations[]", Type: "string", Required: false, Repeated: true, SemanticClass: "limitation"},
			},
			FailureStates:   []string{federationProofStatusRejected, federationPeerStatusStale, federationSyncStatusDiverged},
			SampleRef:       "/v1/public/samples/federation-proof-exchange",
			ConformanceRefs: []string{"/v1/public/conformance-pack", "/v1/public/verifier/reference-pack"},
			Limitations: []string{
				"Federation schema export documents bounded exchange semantics and does not define a global consortium governance model.",
			},
		}, true
	default:
		return publicSchemaExportResponse{}, false
	}
}

func buildPublicVerifierReferencePack() (publicVerifierReferencePackResponse, error) {
	record := publicSampleHandoffRecord()
	bundleBytes, err := buildHandoffBundle(record)
	if err != nil {
		return publicVerifierReferencePackResponse{}, err
	}
	acceptedProofBytes, err := canonicalJSON(publicAcceptedProofSample())
	if err != nil {
		return publicVerifierReferencePackResponse{}, err
	}
	rejectedProofBytes, err := canonicalJSON(publicRejectedProofSample())
	if err != nil {
		return publicVerifierReferencePackResponse{}, err
	}
	validationBytes, err := canonicalJSON(publicSampleValidationCertificate())
	if err != nil {
		return publicVerifierReferencePackResponse{}, err
	}
	federationBytes, err := canonicalJSON(publicDivergedFederationSample())
	if err != nil {
		return publicVerifierReferencePackResponse{}, err
	}
	return publicVerifierReferencePackResponse{
		SchemaVersion: publicVerifierReferencePackSchema,
		PackID:        "5a.reference_verifier_pack.v1",
		Profiles:      []string{"minimal_verifier", "full_verifier", "auditor", "partner_verifier"},
		ReplayInputs: []publicVerifierReplayInput{
			{
				InputID:         "handoff_bundle_valid",
				ProfileID:       "minimal_verifier",
				SampleRef:       "/v1/public/samples/handoff",
				InputType:       "handoff_bundle",
				MediaType:       "application/vnd.changelock.safepkg+zip",
				PayloadEncoding: "base64",
				Payload:         base64.StdEncoding.EncodeToString(bundleBytes),
				ExpectedState:   handoffVerificationValid,
				ExpectedChecks:  []string{"manifest_valid", "artifact_hashes_valid", "signatures_valid", "timestamp_valid", "transparency_valid"},
				Notes:           []string{"Replay input is a real sealed bundle built from the same sample record semantics published in the public sample surfaces."},
			},
			{
				InputID:         "proof_accepted_json",
				ProfileID:       "partner_verifier",
				SampleRef:       "/v1/public/samples/proof-verification",
				InputType:       "proof_verification_case",
				MediaType:       "application/json",
				PayloadEncoding: "utf8_json",
				Payload:         string(acceptedProofBytes),
				ExpectedState:   federationDecisionAccepted,
				ExpectedChecks:  []string{"cryptographic validity preserved", "local decision preserved", "freshness_stale=false"},
			},
			{
				InputID:         "proof_rejected_stale_json",
				ProfileID:       "partner_verifier",
				SampleRef:       "/v1/public/samples/proof-verification",
				InputType:       "proof_verification_case",
				MediaType:       "application/json",
				PayloadEncoding: "utf8_json",
				Payload:         string(rejectedProofBytes),
				ExpectedState:   federationDecisionRejectedStale,
				ExpectedChecks:  []string{"stale rejection preserved", "cryptographic validity distinguished from admissibility"},
			},
			{
				InputID:         "validation_certificate_json",
				ProfileID:       "full_verifier",
				SampleRef:       "/v1/public/samples/validation-certificate",
				InputType:       "validation_certificate",
				MediaType:       "application/json",
				PayloadEncoding: "utf8_json",
				Payload:         string(validationBytes),
				ExpectedState:   validationStatusPass,
				ExpectedChecks:  []string{"scenario_set preserved", "seal_ready preserved", "limitations preserved"},
			},
			{
				InputID:         "federation_diverged_json",
				ProfileID:       "full_verifier",
				SampleRef:       "/v1/public/samples/federation-proof-exchange",
				InputType:       "federation_exchange_case",
				MediaType:       "application/json",
				PayloadEncoding: "utf8_json",
				Payload:         string(federationBytes),
				ExpectedState:   federationSyncStatusDiverged,
				ExpectedChecks:  []string{"policy divergence visible", "local overrides visible", "no global authority implied"},
			},
		},
		UsageNotes: []string{
			"Replay inputs are bounded public reference samples meant for verifier conformance testing, not for live trust decisions.",
			"Verifiers should preserve expected state and limitation semantics without introducing stronger claims than the public samples provide.",
		},
		Limitations: []string{
			"Reference pack covers a minimal bounded set of public verifier cases and does not replace broader interoperability or security review.",
		},
	}, nil
}

func buildPublicSpecCompatibility(version string) publicSpecCompatibility {
	return publicSpecCompatibility{
		CurrentVersion:  version,
		StabilityStatus: "stable",
		StatusClasses: []string{
			"stable",
			"experimental",
			"advisory_only",
		},
		BackwardCompatibility: []string{
			"Additive optional fields are allowed within the same stable schema line.",
			"Existing required fields cannot change meaning without a new schema_version.",
		},
		BreakingChangePolicy: []string{
			"Breaking field removal, incompatible type change, or semantic redefinition requires a new public schema_version.",
			"Experimental fields must be explicitly labeled and cannot silently become required.",
			"Public failure and rejection states remain stable within one schema version line.",
		},
	}
}

func buildPublicHandoffOfflineSteps() []publicVerificationStep {
	record := handoffStoredRecord{
		PackageID:    "HOF-SAMPLE",
		ManifestHash: "sha256:sample-manifest-hash",
	}
	lines := strings.Split(buildVerificationInstructions(record), "\n")
	steps := make([]publicVerificationStep, 0, 5)
	stepIndex := 1
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || !strings.Contains(line, ". ") {
			continue
		}
		step := publicVerificationStep{
			Step:   stepIndex,
			Action: line,
		}
		switch stepIndex {
		case 1:
			step.ExpectedEvidence = []string{"manifest.json", "manifest.artifacts[].sha256"}
			step.FailureStates = []string{handoffVerificationInvalid}
		case 2:
			step.ExpectedEvidence = []string{"verify/manifest.sha256"}
			step.FailureStates = []string{handoffVerificationInvalid}
		case 3:
			step.ExpectedEvidence = []string{"signatures/signatures.json", "verify/public_keys.json"}
			step.FailureStates = []string{handoffVerificationPartial, handoffVerificationInvalid}
		case 4:
			step.ExpectedEvidence = []string{"timestamp/manifest.timestamp.json", "verify/public_keys.json"}
			step.FailureStates = []string{handoffVerificationPartial, handoffVerificationInvalid}
		case 5:
			step.ExpectedEvidence = []string{"transparency/manifest.transparency.json", "verify/public_keys.json"}
			step.FailureStates = []string{handoffVerificationPartial, handoffVerificationInvalid}
		}
		steps = append(steps, step)
		stepIndex++
	}
	return steps
}

func publicSampleTime() time.Time {
	return time.Date(2026, time.April, 21, 10, 0, 0, 0, time.UTC)
}

func publicSampleHandoff() handoffSealResponse {
	record := publicSampleHandoffRecord()
	verification := server{}.verifyStoredHandoff(record)
	return handoffSealResponse{
		PackageID:       record.PackageID,
		Manifest:        record.Manifest,
		Session:         record.Session,
		Bundle:          record.Bundle,
		Verification:    verification,
		DownloadURI:     "/v1/public/samples/handoff",
		VerificationURI: "/v1/public/samples/handoff",
	}
}

func publicSampleHandoffRecord() handoffStoredRecord {
	ts := publicSampleTime()
	seed := "public-sample-seed"
	files := sortHandoffFiles([]handoffBundleFile{
		{Path: "report/report.html", MediaType: "text/html; charset=utf-8", Role: "human_report", Content: "<html><body><h1>Public sample handoff</h1></body></html>"},
		{Path: "evidence/package.json", MediaType: "application/json", Role: "machine_package", Content: `{"incident_refs":["INC-SAMPLE-001"],"summary":"public sample incident package"}`},
		{Path: "evidence/readback_refs.json", MediaType: "application/json", Role: "readback_lineage", Content: `{"readback_refs":[{"resource_type":"incident","resource_id":"INC-SAMPLE-001","evidence_hash":"sha256:incident-sample","resource_uri":"sample://incident/INC-SAMPLE-001"}]}`, AdvisoryOnly: true},
		{Path: "evidence/validation_harness.json", MediaType: "application/json", Role: "validation_harness", Content: string(canonicalJSONMust(publicSampleValidationCertificate())), AdvisoryOnly: true},
	})
	artifacts := make([]sealedManifestArtifact, 0, len(files))
	for _, file := range files {
		artifacts = append(artifacts, sealedManifestArtifact{
			Path:         file.Path,
			MediaType:    file.MediaType,
			SHA256:       digestString(file.Content),
			Role:         file.Role,
			AdvisoryOnly: file.AdvisoryOnly,
		})
	}
	manifest := sealedManifest{
		PackageID:         "HOF-SAMPLE-001",
		PackageType:       handoffPackageTypeIncidentPackage,
		SchemaVersion:     handoffManifestSchemaVersion,
		CreatedAt:         ts,
		GeneratorIdentity: "changelock:public-sample",
		Scope: sealedManifestScope{
			Audience:         incidentAudienceAuditorSafe,
			SelectionMode:    "bounded_sample",
			SelectionSummary: "single incident sample package",
			IncidentCount:    1,
			IncidentRefs:     []string{"INC-SAMPLE-001"},
			TenantID:         "public-sample",
			Environment:      "prod",
			Repo:             "public/sample-app",
		},
		RedactionProfile: sealedManifestRedaction{
			Audience:       incidentAudienceAuditorSafe,
			ProfileVersion: handoffRedactionProfileVersion,
			Summary: []string{
				"sample bundle excludes tenant-internal source record bodies",
				"auditor-safe disclosure preserves verification lineage and bounded readback refs",
			},
		},
		Artifacts:    artifacts,
		EvidenceRefs: []string{"INC-SAMPLE-001", "VALCERT-SAMPLE-001"},
		ReadbackRefs: []sealedManifestReadbackRef{
			{ResourceType: "incident", ResourceID: "INC-SAMPLE-001", EvidenceHash: "sha256:incident-sample", ResourceURI: "sample://incident/INC-SAMPLE-001"},
		},
		ForensicRefs: []sealedManifestForensicRef{
			{ContextURI: "sample://forensics/HOF-SAMPLE-001", ContextType: forensicsModeHistoricalReconstruction, Timestamp: ts.Format(time.RFC3339), AdvisoryOnly: true},
		},
		Limitations: []string{
			"Sample bundle is bounded reference data and does not imply live operational freshness.",
		},
	}
	manifest.RootHash = digestBytesMust(canonicalJSONMust(manifestRootPreimage(manifest)))
	manifestHash := digestBytesMust(canonicalJSONMust(manifest))
	signatures := []signatureRecord{
		signHandoffRecord(manifestHash, seed, handoffSignerRoleSystem),
		signHandoffRecord(manifestHash, seed, handoffSignerRoleAuditor),
	}
	timestampRecord := buildTimestampRecord(manifestHash, seed, ts)
	transparencyRecord := buildTransparencyRecord(manifestHash, seed, ts)
	return handoffStoredRecord{
		PackageID:    manifest.PackageID,
		PackageType:  manifest.PackageType,
		Manifest:     manifest,
		ManifestHash: manifestHash,
		Session: handoffSessionRecord{
			SessionID:      "HOFSESSION-SAMPLE-001",
			PackageID:      manifest.PackageID,
			PackageType:    manifest.PackageType,
			ScopeSummary:   manifest.Scope.SelectionSummary,
			InitiatedBy:    "public-sample",
			InitiatedAt:    ts,
			SignMode:       "public_sample_key",
			CoSignMode:     handoffCoSignRequired,
			Status:         handoffSealStatusFullySealed,
			FinalBundleRef: "/v1/public/samples/handoff",
			ManifestHash:   manifestHash,
		},
		Bundle: sealedBundleMetadata{
			PackageID:              manifest.PackageID,
			BundlePath:             "public-samples/handoff-sample.safepkg",
			ManifestHash:           manifestHash,
			SealStatus:             handoffSealStatusFullySealed,
			SignatureCount:         len(signatures),
			TimestampStatus:        handoffVerificationValid,
			TransparencyStatus:     handoffVerificationValid,
			VerificationURI:        "/v1/public/samples/handoff",
			OfflineVerifierPresent: true,
		},
		Files:           files,
		Signatures:      signatures,
		Timestamp:       timestampRecord,
		Transparency:    transparencyRecord,
		DownloadURI:     "/v1/public/samples/handoff",
		VerificationURI: "/v1/public/samples/handoff",
	}
}

func publicAcceptedProofSample() publicProofVerificationSample {
	ts := publicSampleTime()
	return publicProofVerificationSample{
		CaseID: "accepted_partner_proof",
		Proof: federatedProofResponse{
			RequestID:         "REQ-SAMPLE-001",
			RespondingPeer:    "partner-west",
			ProofType:         federationProofTypeHandoff,
			SealedManifestRef: "/v1/public/samples/handoff",
			ManifestHash:      "sha256:manifest-hash-sample",
			SignatureRefs:     []string{"signatures/system", "signatures/auditor"},
			TimestampRef:      "timestamp/manifest.timestamp.json",
			TransparencyRef:   "transparency/manifest.transparency.json",
			Scope:             federationScope{TenantID: "public-sample", Environment: "prod", Audience: incidentAudienceAuditorSafe, TrustDomain: "partner.public-sample.invalid"},
			RedactionProfile:  incidentAudienceAuditorSafe,
			Freshness: federationProofFreshness{
				IssuedAt:         ts,
				ValidUntil:       ts.Add(45 * time.Minute),
				FreshnessMinutes: 45,
				Stale:            false,
			},
			ReadbackRefs: []sealedManifestReadbackRef{
				{ResourceType: "handoff_manifest", ResourceID: "HOF-SAMPLE-001", EvidenceHash: "sha256:manifest-sample", ResourceURI: "/v1/public/samples/handoff"},
			},
			ForensicRefs: []sealedManifestForensicRef{
				{ContextType: forensicsModeHistoricalReconstruction, Timestamp: ts.Format(time.RFC3339), AdvisoryOnly: true},
			},
			Status: federationProofStatusAccepted,
			Limitations: []string{
				"Accepted example remains bounded by local admissibility and freshness policy.",
			},
		},
		Verification: verificationResult{
			PackageID:           "HOF-SAMPLE-001",
			ManifestValid:       true,
			ArtifactHashesValid: true,
			SignaturesValid:     true,
			TimestampValid:      true,
			TransparencyValid:   true,
			SignerIdentities:    []string{"handoff:system:abcd1234"},
			RedactionProfile:    incidentAudienceAuditorSafe,
			OverallStatus:       handoffVerificationValid,
		},
		Decision: federatedTrustDecision{
			Decision:           federationDecisionAccepted,
			DecisionID:         "FEDDEC-SAMPLE-ACCEPT",
			SubjectRef:         "public/sample-app",
			PeerID:             "partner-west",
			Reasons:            []string{"bundle verified locally", "scope matches local policy", "peer freshness is current"},
			LocalPolicyVersion: "public-sample-policy.1",
			ManifestHash:       "sha256:manifest-hash-sample",
			VerifiedAt:         ts.Add(2 * time.Minute),
		},
		Interpretation: []string{
			"Accepted sample shows one case where cryptographic validity and local admissibility both succeed.",
		},
	}
}

func publicRejectedProofSample() publicProofVerificationSample {
	ts := publicSampleTime()
	return publicProofVerificationSample{
		CaseID: "rejected_stale_partner_proof",
		Proof: federatedProofResponse{
			RequestID:         "REQ-SAMPLE-002",
			RespondingPeer:    "partner-stale",
			ProofType:         federationProofTypeHandoff,
			SealedManifestRef: "/v1/public/samples/handoff",
			ManifestHash:      "sha256:manifest-hash-sample",
			SignatureRefs:     []string{"signatures/system"},
			TimestampRef:      "timestamp/manifest.timestamp.json",
			TransparencyRef:   "transparency/manifest.transparency.json",
			Scope:             federationScope{TenantID: "public-sample", Environment: "prod", Audience: incidentAudienceAuditorSafe, TrustDomain: "stale.public-sample.invalid"},
			RedactionProfile:  incidentAudienceAuditorSafe,
			Freshness: federationProofFreshness{
				IssuedAt:         ts.Add(-3 * time.Hour),
				ValidUntil:       ts.Add(-2 * time.Hour),
				FreshnessMinutes: 60,
				Stale:            true,
			},
			Status: federationProofStatusRejected,
			Limitations: []string{
				"Rejected sample stays cryptographically interpretable but no longer fresh enough for local acceptance.",
			},
		},
		Verification: verificationResult{
			PackageID:           "HOF-SAMPLE-001",
			ManifestValid:       true,
			ArtifactHashesValid: true,
			SignaturesValid:     true,
			TimestampValid:      true,
			TransparencyValid:   true,
			SignerIdentities:    []string{"handoff:system:abcd1234"},
			RedactionProfile:    incidentAudienceAuditorSafe,
			OverallStatus:       handoffVerificationValid,
		},
		Decision: federatedTrustDecision{
			Decision:           federationDecisionRejectedStale,
			DecisionID:         "FEDDEC-SAMPLE-STALE",
			SubjectRef:         "public/sample-app",
			PeerID:             "partner-stale",
			Reasons:            []string{"proof freshness window expired under local policy"},
			LocalPolicyVersion: "public-sample-policy.1",
			ManifestHash:       "sha256:manifest-hash-sample",
			VerifiedAt:         ts,
			Limitations: []string{
				"Cryptographic verification alone does not imply acceptance.",
			},
		},
		Interpretation: []string{
			"Rejected sample preserves valid cryptographic state while still rejecting on freshness grounds.",
		},
	}
}

func publicSampleValidationCertificate() validationCertificate {
	ts := publicSampleTime()
	return validationCertificate{
		RunID:         "VALRUN-SAMPLE-001",
		CertificateID: "VALCERT-SAMPLE-001",
		Scope:         "tenant=public-sample environment=prod repo=public/sample-app service=sample-gateway",
		ScenarioSet:   []string{validationScenarioSafeRelease, validationScenarioLatencyBudget},
		IssuedAt:      ts,
		OverallStatus: validationStatusPass,
		ScenarioResults: []validationVerdict{
			{
				RunID:       "VALRUN-SAMPLE-001",
				VerdictID:   "VERDICT-SAMPLE-001",
				ExecutionID: "VALEXEC-SAMPLE-001",
				ScenarioID:  validationScenarioSafeRelease,
				Status:      validationStatusPass,
				ExpectedOutcome: validationExpectedOutcome{
					Verdict:            validationStatusPass,
					LatencyThresholdMS: 250,
				},
				ObservedOutcome: validationObservedOutcome{
					Verdict:   validationStatusPass,
					LatencyMS: 180,
					Summary:   "Bounded release gate checks matched expected outcome.",
				},
			},
			{
				RunID:       "VALRUN-SAMPLE-001",
				VerdictID:   "VERDICT-SAMPLE-002",
				ExecutionID: "VALEXEC-SAMPLE-002",
				ScenarioID:  validationScenarioLatencyBudget,
				Status:      validationStatusPass,
				ExpectedOutcome: validationExpectedOutcome{
					Verdict:            validationStatusPass,
					LatencyThresholdMS: 600,
				},
				ObservedOutcome: validationObservedOutcome{
					Verdict:   validationStatusPass,
					LatencyMS: 240,
					Summary:   "Control-plane latency stayed within the declared bounded threshold.",
				},
			},
		},
		TimingSummary: validationTimingSummary{
			AverageLatencyMS: 210,
			MaxLatencyMS:     240,
		},
		EnvironmentSummary: validationEnvironmentSummary{
			Environment:    "prod",
			Namespace:      "public-sample-prod",
			Mode:           validationModeRegression,
			EnvironmentTag: "runtime_hardened",
			IsolationClass: "shadow",
			ScopeSummary:   "bounded runtime-hardened enterprise cluster",
			ClusterID:      "public-sample-eu-1",
			TenantID:       "public-sample",
			Repo:           "public/sample-app",
			Service:        "sample-gateway",
		},
		EvidenceRefs: []string{
			"sample://validation/VALRUN-SAMPLE-001",
			"/v1/public/samples/handoff",
		},
		SimulationDerived: false,
		SealReady:         true,
		Limitations: []string{
			"Certificate covers one bounded regression run and must be interpreted together with execution profile and limitations.",
		},
	}
}

func publicReadyFederationSample() publicFederationExchangeSample {
	ts := publicSampleTime()
	return publicFederationExchangeSample{
		CaseID: "ready_exchange",
		Peer: federationPeer{
			PeerID:            "partner-west",
			Organization:      "Partner West",
			Region:            "eu-west",
			Cluster:           "partner-eu-1",
			TrustDomain:       "partner.public-sample.invalid",
			Endpoint:          "https://partner.public-sample.invalid",
			PublicKeys:        []string{"c2FtcGxlLXB1YmxpYy1rZXk="},
			Capabilities:      []string{"sealed_handoff", "policy_sync"},
			PolicyRole:        federationPolicyRoleSupplier,
			Status:            federationPeerStatusActive,
			LastSeen:          ts,
			AcceptedAudiences: []string{incidentAudienceAuditorSafe},
			DisclosureMode:    "sealed_proof_only",
			TrustState: federationPeerTrustState{
				IdentityVerified:        true,
				TrustAnchorFingerprints: []string{"anchor:sample:west"},
				ChannelMode:             "sealed_proof_only",
				FreshnessWindowMinutes:  60,
			},
		},
		Proof:       publicAcceptedProofSample().Proof,
		Decision:    publicAcceptedProofSample().Decision,
		PolicyState: policyFederationState{SyncStatus: federationSyncStatusSynced, EffectivePolicyRoot: "sha256:policy-root-sample", RemotePolicyVersion: "public-sample-policy.1"},
		Interpretation: []string{
			"Ready exchange sample shows a peer in active freshness state with synchronized policy metadata.",
		},
	}
}

func publicStaleFederationSample() publicFederationExchangeSample {
	ts := publicSampleTime()
	return publicFederationExchangeSample{
		CaseID: "stale_exchange",
		Peer: federationPeer{
			PeerID:            "partner-stale",
			Organization:      "Partner Stale",
			Region:            "us-east",
			Cluster:           "partner-us-1",
			TrustDomain:       "stale.public-sample.invalid",
			Endpoint:          "https://stale.public-sample.invalid",
			PublicKeys:        []string{"c3RhbGUtcHVibGljLWtleQ=="},
			Capabilities:      []string{"sealed_handoff"},
			PolicyRole:        federationPolicyRoleSupplier,
			Status:            federationPeerStatusStale,
			LastSeen:          ts.Add(-8 * time.Hour),
			AcceptedAudiences: []string{incidentAudienceAuditorSafe},
			DisclosureMode:    "sealed_proof_only",
			TrustState: federationPeerTrustState{
				IdentityVerified:        true,
				TrustAnchorFingerprints: []string{"anchor:sample:stale"},
				ChannelMode:             "sealed_proof_only",
				FreshnessWindowMinutes:  60,
			},
		},
		Proof:       publicRejectedProofSample().Proof,
		Decision:    publicRejectedProofSample().Decision,
		PolicyState: policyFederationState{SyncStatus: federationSyncStatusStale, EffectivePolicyRoot: "sha256:policy-root-sample"},
		Interpretation: []string{
			"Stale exchange sample preserves peer freshness failure as a first-class rejection reason.",
		},
	}
}

func publicDivergedFederationSample() publicFederationExchangeSample {
	ts := publicSampleTime()
	return publicFederationExchangeSample{
		CaseID: "diverged_exchange",
		Peer: federationPeer{
			PeerID:            "partner-diverged",
			Organization:      "Partner Diverged",
			Region:            "eu-central",
			Cluster:           "partner-eu-2",
			TrustDomain:       "diverged.public-sample.invalid",
			Endpoint:          "https://diverged.public-sample.invalid",
			PublicKeys:        []string{"ZGl2ZXJnZWQtcHVibGljLWtleQ=="},
			Capabilities:      []string{"sealed_handoff", "policy_sync"},
			PolicyRole:        federationPolicyRoleLeader,
			Status:            federationPeerStatusActive,
			LastSeen:          ts,
			AcceptedAudiences: []string{incidentAudienceAuditorSafe},
			DisclosureMode:    "sealed_proof_only",
			TrustState: federationPeerTrustState{
				IdentityVerified:        true,
				TrustAnchorFingerprints: []string{"anchor:sample:diverged"},
				ChannelMode:             "sealed_proof_only",
				FreshnessWindowMinutes:  60,
			},
		},
		Proof: publicAcceptedProofSample().Proof,
		Decision: federatedTrustDecision{
			Decision:            federationDecisionAcceptedWithOverrides,
			DecisionID:          "FEDDEC-SAMPLE-DIVERGED",
			SubjectRef:          "public/sample-app",
			PeerID:              "partner-diverged",
			Reasons:             []string{"bundle verified locally", "local override retained due to policy divergence"},
			LocalPolicyVersion:  "public-sample-policy.2",
			RemotePolicyVersion: "public-sample-policy.1",
			ManifestHash:        "sha256:manifest-hash-sample",
			VerifiedAt:          ts,
		},
		PolicyState: policyFederationState{
			LeaderPeer:          "partner-diverged",
			GlobalPolicyRoot:    "sha256:remote-root-sample",
			LocalPolicyRoot:     "sha256:local-root-sample",
			EffectivePolicyRoot: "sha256:local-root-sample",
			SyncStatus:          federationSyncStatusDiverged,
			LocalOverrides:      []string{"runtime response remains local-policy gated"},
			DivergenceReasons:   []string{"local approval guard differs from remote policy root"},
			RemotePolicyVersion: "public-sample-policy.1",
		},
		Interpretation: []string{
			"Diverged exchange sample shows that verified remote proof can still remain locally bounded by overrides and local effective policy.",
		},
	}
}
