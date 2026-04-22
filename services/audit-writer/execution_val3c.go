package main

import (
	"context"
	"net/http"
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimectrl "github.com/denisgrosek/changelock/internal/runtime"
	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	executionAmbientReadinessSchema      = "3c.ambient_readiness.v1"
	executionConfidentialReadinessSchema = "3c.confidential_readiness.v1"
	executionComplianceReadinessSchema   = "3c.compliance_crypto_readiness.v1"
)

type executionAmbientCapability struct {
	CapabilityID           string   `json:"capability_id"`
	DisplayName            string   `json:"display_name"`
	CurrentState           string   `json:"current_state"`
	SidecarRequirement     string   `json:"sidecar_requirement"`
	EvidenceSurfaces       []string `json:"evidence_surfaces,omitempty"`
	OperatorExplainability string   `json:"operator_explainability"`
	Limitations            []string `json:"limitations,omitempty"`
}

type executionAmbientOverheadComparison struct {
	CurrentImplementation string   `json:"current_implementation"`
	AmbientCandidate      string   `json:"ambient_candidate"`
	SidecarBaseline       string   `json:"sidecar_baseline"`
	Notes                 []string `json:"notes,omitempty"`
}

type executionAmbientReadinessResponse struct {
	SchemaVersion             string                             `json:"schema_version"`
	CurrentState              string                             `json:"current_state"`
	SupportedWorkloadKinds    []string                           `json:"supported_workload_kinds,omitempty"`
	NodeLevelInteractionModel []string                           `json:"node_level_interaction_model,omitempty"`
	CapabilityMatrix          []executionAmbientCapability       `json:"capability_matrix,omitempty"`
	OverheadComparison        executionAmbientOverheadComparison `json:"overhead_comparison"`
	ClosedLoopSummary         audit.RuntimeClosedLoopStatus      `json:"closed_loop_summary"`
	PostureSummary            runtimePostureLinkageSummary       `json:"posture_summary"`
	BlastRadiusCompatibility  []string                           `json:"blast_radius_compatibility,omitempty"`
	FallbackSemantics         []string                           `json:"fallback_semantics,omitempty"`
	Limitations               []string                           `json:"limitations,omitempty"`
}

type executionConfidentialSummary struct {
	TotalSubjects        int `json:"total_subjects"`
	EvidenceBackedClaims int `json:"evidence_backed_claims"`
	MetadataOnlySubjects int `json:"metadata_only_subjects"`
}

type executionConfidentialEvidenceContract struct {
	RequiredEvidence     []string `json:"required_evidence,omitempty"`
	PolicyHints          []string `json:"policy_hints,omitempty"`
	ClaimEvaluationModel []string `json:"claim_evaluation_model,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type executionConfidentialReadinessResponse struct {
	SchemaVersion           string                                `json:"schema_version"`
	CurrentState            string                                `json:"current_state"`
	WorkloadMarkingContract []string                              `json:"workload_marking_contract,omitempty"`
	AttestationLinkage      []string                              `json:"attestation_linkage,omitempty"`
	EvidenceContract        executionConfidentialEvidenceContract `json:"evidence_contract"`
	ScopeSummary            executionConfidentialSummary          `json:"scope_summary"`
	FallbackSemantics       []string                              `json:"fallback_semantics,omitempty"`
	Limitations             []string                              `json:"limitations,omitempty"`
}

type executionCryptoModuleBoundary struct {
	BoundaryID       string   `json:"boundary_id"`
	DisplayName      string   `json:"display_name"`
	Purpose          string   `json:"purpose"`
	ProviderMode     string   `json:"provider_mode"`
	Enabled          bool     `json:"enabled"`
	VerifyOnRead     bool     `json:"verify_on_read"`
	Algorithm        string   `json:"algorithm,omitempty"`
	KeyID            string   `json:"key_id,omitempty"`
	EvidenceSurfaces []string `json:"evidence_surfaces,omitempty"`
	Limitations      []string `json:"limitations,omitempty"`
}

type executionComplianceModeMetadata struct {
	PublicationMode        string `json:"publication_mode"`
	HardeningReviewEnabled bool   `json:"hardening_review_enabled"`
	SignerMode             string `json:"signer_mode"`
	SignerVerifyOnRead     bool   `json:"signer_verify_on_read"`
	StandardsMappingCount  int    `json:"standards_mapping_count"`
}

type executionFIPSReadiness struct {
	State           string   `json:"state"`
	Applicable      bool     `json:"applicable"`
	CurrentProvider string   `json:"current_provider"`
	Summary         string   `json:"summary"`
	Limitations     []string `json:"limitations,omitempty"`
}

type executionComplianceReadinessResponse struct {
	SchemaVersion          string                          `json:"schema_version"`
	ComplianceModeMetadata executionComplianceModeMetadata `json:"compliance_mode_metadata"`
	CryptoModuleBoundaries []executionCryptoModuleBoundary `json:"crypto_module_boundaries,omitempty"`
	StandardsMappings      []audit.StandardsMapping        `json:"standards_mappings,omitempty"`
	FIPSReadiness          executionFIPSReadiness          `json:"fips_readiness"`
	OperatorGuidance       []string                        `json:"operator_guidance,omitempty"`
	FormalClaimsExcluded   []string                        `json:"formal_claims_excluded,omitempty"`
	Limitations            []string                        `json:"limitations,omitempty"`
}

func (s server) executionAmbientReadinessHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseExecutionCoverageFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildExecutionAmbientReadiness(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) executionConfidentialReadinessHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseExecutionCoverageFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildExecutionConfidentialReadiness(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) executionComplianceReadinessHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := applyPrincipalTenantToTrustScopeRequest(principal, parseTrustScopeRequestFromQuery(r))
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildExecutionComplianceReadiness(ctx, scope)
	if err != nil {
		s.writeScorecardError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildExecutionAmbientReadiness(ctx context.Context, filter runtimeIntegrityFilter) (executionAmbientReadinessResponse, error) {
	runtimeCfg, err := runtimectrl.LoadSelfHealingConfig()
	if err != nil {
		return executionAmbientReadinessResponse{}, err
	}
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		Component:   "runtime-agent",
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       maxInt(filter.Limit*20, 250),
	})
	if err != nil {
		return executionAmbientReadinessResponse{}, err
	}
	activeStates := audit.DeriveRuntimeActiveStates(events, audit.RuntimeActiveStateFilter{
		ClusterID:    filter.ClusterID,
		TenantID:     filter.TenantID,
		Namespace:    filter.Namespace,
		WorkloadKind: filter.WorkloadKind,
		Workload:     filter.Workload,
		Limit:        maxInt(filter.Limit*4, 250),
	})
	status := audit.DeriveRuntimeClosedLoopStatus(activeStates)
	postureItems, _, err := s.buildRuntimePostureStates(ctx, filter)
	if err != nil {
		return executionAmbientReadinessResponse{}, err
	}
	supportedKinds := sortedAllowedKinds(runtimeCfg.AllowedKinds)
	currentState := "readiness_only"
	switch {
	case runtimeCfg.Mode != runtimectrl.RemediationModeDisabled && runtimeCfg.QuarantineNetworkPolicy:
		currentState = "sidecarless_policy_overlay_candidate"
	case runtimeCfg.Mode != runtimectrl.RemediationModeDisabled:
		currentState = "controller_only_candidate"
	}

	return executionAmbientReadinessResponse{
		SchemaVersion:          executionAmbientReadinessSchema,
		CurrentState:           currentState,
		SupportedWorkloadKinds: supportedKinds,
		NodeLevelInteractionModel: []string{
			"Ambient readiness is currently bounded to runtime-agent reconciliation, workload posture linkage, and optional NetworkPolicy quarantine overlay.",
			"Current sidecarless candidate path uses control-plane and namespace policy interactions instead of per-workload sidecar injection.",
			"Protected namespaces and workloads remain explicit guardrails against broad autonomous mutation.",
		},
		CapabilityMatrix: []executionAmbientCapability{
			{
				CapabilityID:           "workload_posture_guidance",
				DisplayName:            "Workload posture and scheduling guidance",
				CurrentState:           "ready",
				SidecarRequirement:     "not_required",
				EvidenceSurfaces:       []string{"/v1/runtime/posture", "/v1/runtime/posture-linkage", "/v1/runtime/boundaries"},
				OperatorExplainability: "Scheduling and mismatch posture remains explainable through rule packs, posture linkage, and boundary contracts.",
			},
			{
				CapabilityID:           "controller_reconciliation",
				DisplayName:            "Controller-style sidecarless reconciliation",
				CurrentState:           mapAmbientState(runtimeCfg.Mode != runtimectrl.RemediationModeDisabled, "ready", "disabled"),
				SidecarRequirement:     "not_required",
				EvidenceSurfaces:       []string{"/v1/runtime/active-state", "/v1/runtime/closed-loop/status"},
				OperatorExplainability: "Closed-loop status and active-state evidence expose reconciliation, quarantine, and protected-target behavior without sidecars.",
				Limitations: []string{
					"Automatic reconciliation remains intentionally bounded to Deployment, DaemonSet, and StatefulSet.",
				},
			},
			{
				CapabilityID:           "network_policy_overlay",
				DisplayName:            "NetworkPolicy-based sidecarless containment overlay",
				CurrentState:           mapAmbientState(runtimeCfg.QuarantineNetworkPolicy, "candidate_enabled", "candidate_disabled"),
				SidecarRequirement:     "not_required",
				EvidenceSurfaces:       []string{"/v1/runtime/quarantine", "/v1/runtime/closed-loop/status"},
				OperatorExplainability: "Containment remains visible through quarantine type, reason, and active-state evidence rather than hidden dataplane mutation.",
				Limitations: []string{
					"Overlay enforcement depends on cluster CNI NetworkPolicy support and is bounded to containment intent plus evidence when CNI enforcement is missing.",
					"ChangeLock does not synthesize broad mesh or firewall policy in this slice.",
				},
			},
		},
		OverheadComparison: executionAmbientOverheadComparison{
			CurrentImplementation: "runtime-agent plus control-plane evidence surfaces; no per-workload sidecar injection is required today.",
			AmbientCandidate:      "Bounded sidecarless readiness uses controller reconciliation and optional NetworkPolicy overlay as the nearest ambient path.",
			SidecarBaseline:       "No sidecar baseline is implemented in the current workspace, so comparison remains structural rather than benchmark-backed.",
			Notes: []string{
				"Use runtime-agent sizing starting points and control-plane latency budget discipline as overhead ceilings, not as benchmark guarantees.",
			},
		},
		ClosedLoopSummary: status,
		PostureSummary:    summarizeRuntimePostureLinkage(postureItems),
		BlastRadiusCompatibility: []string{
			"Workload-only containment remains the preferred sidecarless blast-radius boundary.",
			"Namespace-level effects depend on verified NetworkPolicy behavior and operator review.",
			"Cluster-wide mesh, firewall, or ambient dataplane synthesis is outside the current slice.",
		},
		FallbackSemantics: []string{
			"If NetworkPolicy overlay is disabled or unsupported, ambient readiness falls back to controller-only evidence and advisory containment intent.",
			"If desired-state trust is missing, remediation downgrades to quarantine or alert-only based on closed-loop fail mode instead of silently mutating.",
			"If protected targets match, automatic sidecarless mutation remains blocked and the system stays in detect-and-explain mode.",
		},
		Limitations: []string{
			"Ambient readiness is a bounded readiness path, not a claim that a meshless dataplane enforcement layer is fully implemented.",
			"Current overhead comparison is structural and policy-based; it is not benchmark-backed sidecar-vs-ambient latency proof.",
		},
	}, nil
}

func (s server) buildExecutionConfidentialReadiness(ctx context.Context, filter runtimeIntegrityFilter) (executionConfidentialReadinessResponse, error) {
	postureItems, _, err := s.buildRuntimePostureStates(ctx, filter)
	if err != nil {
		return executionConfidentialReadinessResponse{}, err
	}
	summary := summarizeConfidentialReadiness(postureItems)
	currentState := "metadata_only"
	switch {
	case summary.TotalSubjects > 0 && summary.EvidenceBackedClaims == summary.TotalSubjects:
		currentState = "evidence_backed_candidate"
	case summary.EvidenceBackedClaims > 0:
		currentState = "mixed_scope"
	}
	return executionConfidentialReadinessResponse{
		SchemaVersion: executionConfidentialReadinessSchema,
		CurrentState:  currentState,
		WorkloadMarkingContract: []string{
			"Confidential execution readiness remains metadata-only in this slice and should be expressed through explicit workload policy hints rather than implicit scheduler assumptions.",
			"Recommended hint categories are: confidential_execution_requested, confidential_execution_required, and confidential_execution_fallback_allowed.",
			"Absent confidential substrate evidence, these hints stay advisory and must not be reinterpreted as enclave-backed guarantees.",
		},
		AttestationLinkage: []string{
			"Confidential claims require explicit substrate or enclave attestation evidence that can be linked into workload trust posture.",
			"Runtime posture linkage remains the source for expected-versus-actual trust state; confidential claims must appear there as evidence-backed inputs rather than free-form labels.",
			"When confidential substrate evidence is missing, no enclave-backed claim is emitted and the workload remains on standard verified posture only.",
		},
		EvidenceContract: executionConfidentialEvidenceContract{
			RequiredEvidence: []string{
				"workload-scoped substrate or enclave attestation proof",
				"evidence reference that binds the confidential claim to the workload execution substrate",
				"fallback marker when confidential substrate is unavailable",
			},
			PolicyHints: []string{
				"confidential_execution_requested",
				"confidential_execution_required",
				"confidential_execution_fallback_allowed",
			},
			ClaimEvaluationModel: []string{
				"metadata-only hints without substrate evidence remain advisory",
				"evidence-backed confidential claims must stay linked to runtime posture and attestation lineage",
				"fallback semantics are mandatory when confidential substrate is unavailable or unverifiable",
			},
			Limitations: []string{
				"This slice does not infer enclave, TEE, SEV, TDX, or other confidential substrate guarantees unless explicit evidence is already present in scope.",
			},
		},
		ScopeSummary: summary,
		FallbackSemantics: []string{
			"If confidential substrate is unavailable, the workload must remain on standard verified execution posture and record fallback status explicitly.",
			"If attestation linkage is missing, confidential readiness falls back to metadata-only and no evidence-backed confidential claim is emitted.",
			"Recovery, hardening, and scheduling decisions continue to use the standard runtime posture contract until substrate evidence becomes explicit.",
		},
		Limitations: []string{
			"Confidential readiness is readiness metadata plus evidence rules only; it is not an implemented enclave scheduling or remote-attestation control plane.",
		},
	}, nil
}

func (s server) buildExecutionComplianceReadiness(ctx context.Context, scope trustScopeRequest) (executionComplianceReadinessResponse, error) {
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		return executionComplianceReadinessResponse{}, err
	}
	input, err := s.collectTrustScorecardInput(ctx, scope, cfg)
	if err != nil {
		return executionComplianceReadinessResponse{}, err
	}
	card := audit.ComputeTrustScorecard(input)
	mappings := audit.BuildStandardsMapping(card)
	cryptoBoundaries := executionCryptoModuleBoundaries(s.signing)
	signerMode := signing.ModeDisabled
	verifyOnRead := false
	if s.signing != nil && s.signing.runtime != nil {
		signerMode = s.signing.runtime.Config.Mode
		verifyOnRead = s.signing.runtime.Config.VerifyOnRead
	}
	return executionComplianceReadinessResponse{
		SchemaVersion: executionComplianceReadinessSchema,
		ComplianceModeMetadata: executionComplianceModeMetadata{
			PublicationMode:        cfg.PublicationMode,
			HardeningReviewEnabled: cfg.HardeningReviewEnabled,
			SignerMode:             signerMode,
			SignerVerifyOnRead:     verifyOnRead,
			StandardsMappingCount:  len(mappings),
		},
		CryptoModuleBoundaries: cryptoBoundaries,
		StandardsMappings:      mappings,
		FIPSReadiness:          executionFIPSReadinessFromSigningRuntime(s.signing),
		OperatorGuidance: []string{
			"Treat standards mappings and crypto readiness as operator guidance and deployment validation inputs, not as certification outcomes.",
			"If signer mode is provider-backed, validate provider connectivity, verify-on-read behavior, and go-live smoke checks before relying on regulated deployment posture.",
			"Keep private-key custody outside ChangeLock when stricter crypto-hardening is required; software mode is development-friendly but not a provider-backed regulated posture.",
		},
		FormalClaimsExcluded: []string{
			"formal FIPS 140-2 or FIPS 140-3 certification claim",
			"formal third-party compliance certification",
			"implicit regulated-environment suitability without deployment-specific validation",
		},
		Limitations: []string{
			"Compliance and crypto-hardening readiness remain bounded to measured standards mappings, signer runtime configuration, and provider model metadata already present in ChangeLock.",
			"Readiness path is not a substitute for formal certification, external audit, or provider-specific compliance attestations.",
		},
	}, nil
}

func executionCryptoModuleBoundaries(runtime *signingRuntime) []executionCryptoModuleBoundary {
	providerMode := signing.ModeDisabled
	algorithm := ""
	keyID := ""
	verifyOnRead := false
	if runtime != nil && runtime.runtime != nil {
		providerMode = runtime.runtime.Config.Mode
		algorithm = runtime.runtime.Config.Algorithm
		keyID = runtime.runtime.Config.KeyID
		verifyOnRead = runtime.runtime.Config.VerifyOnRead
	}
	items := []executionCryptoModuleBoundary{
		{
			BoundaryID:       "exceptions_evidence_signing",
			DisplayName:      "Exceptions evidence signing",
			Purpose:          signing.PurposeExceptions,
			ProviderMode:     providerMode,
			Enabled:          runtime != nil && runtime.enabledForPurpose(signing.PurposeExceptions),
			VerifyOnRead:     runtime != nil && runtime.verifyOnRead(signing.PurposeExceptions),
			Algorithm:        algorithm,
			KeyID:            keyID,
			EvidenceSurfaces: []string{"approved exception evidence", "exception verification state on read"},
			Limitations: []string{
				"Boundary is limited to ChangeLock-controlled evidence signatures; it does not replace image signing, Fulcio, Rekor, or external PKI.",
			},
		},
		{
			BoundaryID:       "sync_snapshot_signing",
			DisplayName:      "Sync snapshot signing",
			Purpose:          signing.PurposeSyncSnapshots,
			ProviderMode:     providerMode,
			Enabled:          runtime != nil && runtime.enabledForPurpose(signing.PurposeSyncSnapshots),
			VerifyOnRead:     runtime != nil && runtime.verifyOnRead(signing.PurposeSyncSnapshots),
			Algorithm:        algorithm,
			KeyID:            keyID,
			EvidenceSurfaces: []string{"exception sync snapshots", "sync snapshot verification state on read"},
			Limitations: []string{
				"Boundary is limited to ChangeLock-controlled sync evidence; it is not a generic document-signing module for arbitrary user payloads.",
			},
		},
		{
			BoundaryID:       "control_plane_verification_mode",
			DisplayName:      "Control-plane verification mode",
			Purpose:          "verify_on_read",
			ProviderMode:     providerMode,
			Enabled:          verifyOnRead,
			VerifyOnRead:     verifyOnRead,
			Algorithm:        algorithm,
			KeyID:            keyID,
			EvidenceSurfaces: []string{"exception verification state", "sync snapshot verification state"},
			Limitations: []string{
				"Verify-on-read expresses bounded verification behavior for ChangeLock-controlled evidence only.",
			},
		},
	}
	sort.Slice(items, func(i, j int) bool { return items[i].BoundaryID < items[j].BoundaryID })
	return items
}

func executionFIPSReadinessFromSigningRuntime(runtime *signingRuntime) executionFIPSReadiness {
	if runtime == nil || runtime.runtime == nil {
		return executionFIPSReadiness{
			State:           "disabled",
			Applicable:      false,
			CurrentProvider: signing.ModeDisabled,
			Summary:         "No provider-backed signing runtime is active, so stricter crypto-hardening readiness remains disabled.",
			Limitations: []string{
				"No formal FIPS claim is made when signing is disabled.",
			},
		}
	}
	switch runtime.runtime.Config.Mode {
	case signing.ModeVaultTransit:
		return executionFIPSReadiness{
			State:           "provider_backed_candidate",
			Applicable:      true,
			CurrentProvider: signing.ModeVaultTransit,
			Summary:         "Vault transit mode keeps key custody outside ChangeLock and is the current provider-backed candidate path for stricter crypto-hardening readiness.",
			Limitations: []string{
				"Provider-backed mode improves crypto custody posture but is not itself a FIPS certification claim.",
				"Direct PKCS#11 or broader cloud-KMS-specific compliance integrations remain out of scope in the current workspace.",
			},
		}
	case signing.ModeSoftware:
		return executionFIPSReadiness{
			State:           "software_only_not_ready",
			Applicable:      true,
			CurrentProvider: signing.ModeSoftware,
			Summary:         "Software signer mode is useful for local and bounded environments, but it is not the provider-backed path for stricter crypto-hardening readiness.",
			Limitations: []string{
				"Software mode should not be described as FIPS-ready or provider-backed.",
			},
		}
	default:
		return executionFIPSReadiness{
			State:           "disabled",
			Applicable:      false,
			CurrentProvider: signing.ModeDisabled,
			Summary:         "No active signer provider mode is configured for a stricter crypto-hardening path.",
			Limitations: []string{
				"Disabled mode does not support FIPS-readiness mapping beyond explicit not-enabled status.",
			},
		}
	}
}

func summarizeConfidentialReadiness(items []runtimePostureState) executionConfidentialSummary {
	summary := executionConfidentialSummary{TotalSubjects: len(items)}
	for _, item := range items {
		if hasConfidentialEvidence(item) {
			summary.EvidenceBackedClaims++
		} else {
			summary.MetadataOnlySubjects++
		}
	}
	return summary
}

func hasConfidentialEvidence(item runtimePostureState) bool {
	for _, value := range item.ExpectedTrustState.AttestationInputs {
		if isStructuredConfidentialEvidenceMarker(value) {
			return true
		}
	}
	for _, value := range item.ActualTrustState.AttestationInputs {
		if isStructuredConfidentialEvidenceMarker(value) {
			return true
		}
	}
	return false
}

func isStructuredConfidentialEvidenceMarker(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "confidential_attestation",
		"confidential_substrate",
		"enclave_attestation",
		"tee_attestation",
		"tdx_attestation",
		"sev_attestation",
		"sgx_attestation":
		return true
	default:
		return false
	}
}

func sortedAllowedKinds(values map[string]struct{}) []string {
	items := make([]string, 0, len(values))
	for key := range values {
		items = append(items, key)
	}
	sort.Strings(items)
	return items
}

func mapAmbientState(enabled bool, whenEnabled, whenDisabled string) string {
	if enabled {
		return whenEnabled
	}
	return whenDisabled
}
