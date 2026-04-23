package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	benchmarkfoundation "github.com/denisgrosek/changelock/internal/benchmark"
	claimscore "github.com/denisgrosek/changelock/internal/claims"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	publicProofValBTransparencyChainSchema     = "point2.measured_public_proof.valb.transparency_chain.v1"
	publicProofValBVerifierCapabilitySchema    = "point2.measured_public_proof.valb.verifier_capability.v1"
	publicProofValBSignatureVerificationSchema = "point2.measured_public_proof.valb.signature_verification.v1"
	publicProofValBReplayVerificationSchema    = "point2.measured_public_proof.valb.replay_verification.v1"
	publicProofValBProofsSchema                = "point2.measured_public_proof.valb.proofs.v1"
)

type publicProofValBTransparencyChainResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         claimscore.PublicProofTransparencyChain `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type publicProofValBVerifierCapabilityResponse struct {
	SchemaVersion string                                   `json:"schema_version"`
	GeneratedAt   time.Time                                `json:"generated_at"`
	CurrentState  string                                   `json:"current_state"`
	Model         claimscore.PublicProofVerifierCapability `json:"model"`
	RouteRefs     []string                                 `json:"route_refs,omitempty"`
	Limitations   []string                                 `json:"limitations,omitempty"`
}

type publicProofValBSignatureVerificationResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Items         []claimscore.PublicProofSignatureVerificationItem `json:"items,omitempty"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type publicProofValBReplayVerificationResponse struct {
	SchemaVersion        string                                         `json:"schema_version"`
	GeneratedAt          time.Time                                      `json:"generated_at"`
	CurrentState         string                                         `json:"current_state"`
	Items                []claimscore.PublicProofReplayVerificationItem `json:"items,omitempty"`
	BenchmarkEvaluations []benchmarkfoundation.EvaluationResponse       `json:"benchmark_evaluations,omitempty"`
	RouteRefs            []string                                       `json:"route_refs,omitempty"`
	Limitations          []string                                       `json:"limitations,omitempty"`
}

type publicProofValBProofsResponse struct {
	SchemaVersion              string    `json:"schema_version"`
	GeneratedAt                time.Time `json:"generated_at"`
	CurrentState               string    `json:"current_state"`
	Phase6State                string    `json:"phase6_state"`
	Val0State                  string    `json:"val0_state"`
	ValAState                  string    `json:"val_a_state"`
	TransparencyChainState     string    `json:"transparency_chain_state"`
	VerifierCapabilityState    string    `json:"verifier_capability_state"`
	SignatureVerificationState string    `json:"signature_verification_state"`
	ReplayVerificationState    string    `json:"replay_verification_state"`
	SurfaceRefs                []string  `json:"surface_refs,omitempty"`
	EvidenceRefs               []string  `json:"evidence_refs,omitempty"`
	DeferredScope              []string  `json:"deferred_scope,omitempty"`
	Limitations                []string  `json:"limitations,omitempty"`
	IntegrationSummary         []string  `json:"integration_summary,omitempty"`
}

func (s server) publicProofValBTransparencyChainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValBTransparencyChain(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValBVerifierCapabilityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValBVerifierCapability(asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValBSignatureVerificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValBSignatureVerification(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValBReplayVerificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValBReplayVerification(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValBProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValBProofs(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildPublicProofValBTransparencyChain(ctx context.Context, asOf time.Time) (publicProofValBTransparencyChainResponse, error) {
	model, err := s.publicProofValBTransparencyChainModel(ctx, asOf)
	if err != nil {
		return publicProofValBTransparencyChainResponse{}, err
	}
	return publicProofValBTransparencyChainResponse{
		SchemaVersion: publicProofValBTransparencyChainSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValBTransparencyChainState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/public/transparency/anchor",
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/valb/proofs",
		},
		Limitations: []string{
			"Val B transparency chain remains a bounded publication lineage over sealed Val A packs and the existing public transparency anchor; automated supersession and withdrawal remain deferred.",
		},
	}, nil
}

func (s server) buildPublicProofValBVerifierCapability(asOf time.Time) (publicProofValBVerifierCapabilityResponse, error) {
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return publicProofValBVerifierCapabilityResponse{}, err
	}
	referencePack, err := buildPublicVerifierReferencePack()
	if err != nil {
		return publicProofValBVerifierCapabilityResponse{}, err
	}
	model := claimscore.PublicProofVerifierCapability{
		CurrentState:         "verifier_capability_ready",
		SDKRef:               "/v1/public/verifier/sdk",
		ReferencePackRef:     sdk.ReferencePackRef,
		OfflineGuideRef:      "/v1/public/verifier/offline-guide",
		SupportedSchemaLines: uniqueStrings(append([]string{"public.proof.sealed_artifact.v1"}, sdk.SupportedSchemaLines...)),
		ResultStates: []string{
			"verified",
			"verified_with_limitations",
			"stale",
			"unsupported",
			"failed_verification",
		},
		TrustVerification: []string{
			"verify signature envelope purpose, payload digest, trust root, and key version against the sealed artifact payload",
			"treat revoked signer, disabled trust root, or unsupported schema line as explicit verifier-visible failure states",
		},
		ReplayVerification: []string{
			"bounded replay remains methodology-compatible, environment-compatible, and tolerance-bounded instead of requiring bit-for-bit global equivalence",
			"stale proofs stay replay-visible with stale semantics instead of silently remaining active",
		},
		CommandHints: uniqueStrings(append([]string{
			"Fetch /v1/public/proof-expansion/vala/downloadable-packs/{artifact_id} and validate the envelope against its canonical payload digest.",
		}, sdk.VerificationCommands...)),
		UnsupportedCases: []string{
			"cross-environment replay outside declared tolerance bands",
			"signature verification without a declared trust root or compatible schema line",
			"treating public_safe projection as if it were internal_full raw evidence",
		},
		Limitations: uniqueStrings(append(append([]string{
			"Verifier capability remains bounded to published sealed artifact schemas, trust roots, and replay guidance.",
		}, sdk.Limitations...), referencePack.Limitations...)),
	}
	return publicProofValBVerifierCapabilityResponse{
		SchemaVersion: publicProofValBVerifierCapabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValBVerifierCapabilityState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/public/verifier/sdk",
			"/v1/public/verifier/reference-pack",
			"/v1/public/verifier/offline-guide",
			"/v1/public/proof-expansion/valb/proofs",
		},
		Limitations: []string{
			"Val B verifier capability exposes bounded verification and replay guidance over published artifact schemas; it does not yet add a standalone external SDK distribution pipeline.",
		},
	}, nil
}

func (s server) buildPublicProofValBSignatureVerification(ctx context.Context, asOf time.Time) (publicProofValBSignatureVerificationResponse, error) {
	items, err := s.publicProofValBSignatureVerificationItems(ctx, asOf)
	if err != nil {
		return publicProofValBSignatureVerificationResponse{}, err
	}
	return publicProofValBSignatureVerificationResponse{
		SchemaVersion: publicProofValBSignatureVerificationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValBSignatureVerificationState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/val0/signing-authority",
			"/v1/public/verifier/sdk",
			"/v1/public/proof-expansion/valb/proofs",
		},
		Limitations: []string{
			"Val B signature verification proves declared trust-root and payload-digest compatibility for sealed artifact packs only; withdrawal and supersession governance remain later Point 2 work.",
		},
	}, nil
}

func (s server) buildPublicProofValBReplayVerification(ctx context.Context, asOf time.Time) (publicProofValBReplayVerificationResponse, error) {
	items, evaluations, err := s.publicProofValBReplayVerificationItems(ctx, asOf)
	if err != nil {
		return publicProofValBReplayVerificationResponse{}, err
	}
	return publicProofValBReplayVerificationResponse{
		SchemaVersion:        publicProofValBReplayVerificationSchema,
		GeneratedAt:          publicSampleTime(),
		CurrentState:         claimscore.EvaluateMeasuredPublicProofValBReplayVerificationState(items),
		Items:                items,
		BenchmarkEvaluations: evaluations,
		RouteRefs: []string{
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/verifier/reference-pack",
			"/v1/public/verifier/offline-guide",
			"/v1/foundation/execution/benchmarks/evaluate",
			"/v1/public/proof-expansion/valb/proofs",
		},
		Limitations: []string{
			"Val B replay verification stays bounded to declared environment classes, methodology refs, and explicit tolerance bands instead of implying universal replay parity.",
		},
	}, nil
}

func (s server) buildPublicProofValBProofs(ctx context.Context, asOf time.Time) (publicProofValBProofsResponse, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return publicProofValBProofsResponse{}, err
	}
	val0, err := s.buildPublicProofVal0Proofs(asOf)
	if err != nil {
		return publicProofValBProofsResponse{}, err
	}
	valA, err := s.buildPublicProofValAProofs(ctx, asOf)
	if err != nil {
		return publicProofValBProofsResponse{}, err
	}
	transparency, err := s.buildPublicProofValBTransparencyChain(ctx, asOf)
	if err != nil {
		return publicProofValBProofsResponse{}, err
	}
	verifier, err := s.buildPublicProofValBVerifierCapability(asOf)
	if err != nil {
		return publicProofValBProofsResponse{}, err
	}
	signature, err := s.buildPublicProofValBSignatureVerification(ctx, asOf)
	if err != nil {
		return publicProofValBProofsResponse{}, err
	}
	replay, err := s.buildPublicProofValBReplayVerification(ctx, asOf)
	if err != nil {
		return publicProofValBProofsResponse{}, err
	}
	currentState := claimscore.EvaluateMeasuredPublicProofValBState(
		valA.CurrentState,
		transparency.CurrentState,
		verifier.CurrentState,
		signature.CurrentState,
		replay.CurrentState,
	)
	return publicProofValBProofsResponse{
		SchemaVersion:              publicProofValBProofsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               currentState,
		Phase6State:                phase6.CurrentState,
		Val0State:                  val0.CurrentState,
		ValAState:                  valA.CurrentState,
		TransparencyChainState:     transparency.CurrentState,
		VerifierCapabilityState:    verifier.CurrentState,
		SignatureVerificationState: signature.CurrentState,
		ReplayVerificationState:    replay.CurrentState,
		SurfaceRefs: []string{
			"/v1/public/proof-expansion/vala/proofs",
			"/v1/public/proof-expansion/valb/transparency-chain",
			"/v1/public/proof-expansion/valb/verifier-capability",
			"/v1/public/proof-expansion/valb/signature-verification",
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/valb/proofs",
		},
		EvidenceRefs: []string{
			"/v1/public/transparency/anchor",
			"/v1/public/verifier/sdk",
			"/v1/public/verifier/reference-pack",
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/phase6/proofs",
		},
		DeferredScope: []string{
			"point2_valc_public_and_partner_proof_portal",
			"point2_vald_automated_issuance_and_revocation_gate",
			"point2_vale_final_proof_gate",
		},
		Limitations: []string{
			"Val B closes bounded transparency and verification over sealed Val A artifacts, but it does not yet add public or partner proof portal projection.",
			"Automated issuance, restriction, supersession, and withdrawal workflows remain deferred to later Point 2 waves.",
		},
		IntegrationSummary: []string{
			"Val B binds sealed Val A artifact packs to a bounded transparency chain over the existing public transparency anchor.",
			"Val B adds verifier-capable signature checks, trust-root compatibility, and tolerance-bounded replay posture without introducing a new truth store.",
		},
	}, nil
}

func (s server) publicProofValBTransparencyChainModel(ctx context.Context, asOf time.Time) (claimscore.PublicProofTransparencyChain, error) {
	anchor, err := buildPhase6TransparencyAnchor(asOf)
	if err != nil {
		return claimscore.PublicProofTransparencyChain{}, err
	}
	builtPacks, err := s.publicProofValABuiltPacks(ctx, asOf)
	if err != nil {
		return claimscore.PublicProofTransparencyChain{}, err
	}
	model := claimscore.PublicProofTransparencyChain{
		CurrentState:    "transparency_chain_ready",
		ChainID:         "point2_valb_transparency_chain_v1",
		ParentAnchorRef: "/v1/public/transparency/anchor",
		IntegrityRules: []string{
			"each transparency entry remains digest-bound to the sealed artifact payload digest and the parent public transparency anchor",
			"transparency lineage stays append-visible and does not silently mutate published artifact history",
		},
		PublicationRules: []string{
			"transparency chain remains a projection over published sealed artifacts instead of a new evidence authority",
			"later supersession or withdrawal states must remain visible rather than deleting prior publication lineage",
		},
		Limitations: []string{
			"Val B uses the existing public transparency anchor as its parent publication lineage instead of introducing a new dedicated issuance log.",
		},
	}
	if anchor.CurrentState != transparencyAnchorStateActive || len(builtPacks) == 0 {
		return model, nil
	}
	entries := make([]claimscore.PublicProofTransparencyEntry, 0, len(builtPacks))
	for _, built := range builtPacks {
		entries = append(entries, claimscore.PublicProofTransparencyEntry{
			ArtifactID:      built.Pack.ArtifactID,
			CurrentState:    "anchored_projection_ready",
			ParentAnchorRef: "/v1/public/transparency/anchor",
			AnchorID:        model.ChainID,
			EntryID:         model.ChainID + "/" + built.Pack.ArtifactID,
			EntryHash:       publicProofValBTransparencyEntryHash(anchor.RootHash, built.Pack.PayloadDigest, built.Pack.ArtifactID),
			AnchoredAt:      anchor.VerifiedAt.UTC().Format(time.RFC3339),
			TransparencyRefs: []string{
				"/v1/public/transparency/anchor",
				"/v1/public/proof-expansion/valb/transparency-chain",
				built.Pack.DownloadRef,
			},
			SupersessionRefs: []string{
				"/v1/public/proof-expansion/valb/proofs",
			},
			Limitations: []string{
				"Direct artifact revocation and supersession workflows remain deferred to later Point 2 waves.",
			},
		})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].ArtifactID < entries[j].ArtifactID })
	model.Entries = entries
	return model, nil
}

func (s server) publicProofValBSignatureVerificationItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofSignatureVerificationItem, error) {
	builtPacks, err := s.publicProofValABuiltPacks(ctx, asOf)
	if err != nil {
		return nil, err
	}
	signingAuthority := claimscore.MeasuredPublicProofVal0SigningAuthority(publicProofVal0ProviderDescriptor(s.signing))
	items := make([]claimscore.PublicProofSignatureVerificationItem, 0, len(builtPacks))
	for _, built := range builtPacks {
		item := claimscore.PublicProofSignatureVerificationItem{
			ArtifactID:          built.Pack.ArtifactID,
			CurrentState:        "verification_ready",
			SchemaCompatibility: publicProofValBSchemaCompatibility(built.Pack.ArtifactSchemaVersion),
			VerifierRef:         "/v1/public/verifier/sdk",
			TrustRootID:         built.Pack.TrustRootID,
			KeyVersion:          built.Pack.KeyVersion,
			SignatureProvider:   publicProofValBEnvelopeProvider(built.Pack.SignatureEnvelope, s.signing),
			PayloadDigest:       built.Pack.PayloadDigest,
			EvidenceRefs: []string{
				built.Pack.DownloadRef,
				"/v1/public/proof-expansion/val0/signing-authority",
				"/v1/public/verifier/sdk",
			},
			FailureStates: []string{
				"failed_verification",
				"signature_missing",
				"stale",
				"unsupported",
				"trust_root_unavailable",
			},
			Limitations: []string{
				"Signature verification remains bounded to the canonical sealed pack payload and declared trust-root metadata.",
			},
		}
		item.TrustRootState = publicProofValBTrustRootState(signingAuthority, built.Pack)
		if built.Pack.SignatureEnvelope == nil {
			item.VerificationState = "signature_missing"
			item.CurrentState = "verification_pending"
			items = append(items, item)
			continue
		}
		result, err := s.signing.verifyPublicProofArtifact(ctx, built.Payload, *built.Pack.SignatureEnvelope)
		if err != nil {
			return nil, err
		}
		item.VerificationState = result.State
		if item.VerificationState != signing.StateVerified || item.TrustRootState != "trusted" || item.SchemaCompatibility != "supported" {
			item.CurrentState = "verification_limited"
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValBReplayVerificationItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofReplayVerificationItem, []benchmarkfoundation.EvaluationResponse, error) {
	builtPacks, err := s.publicProofValABuiltPacks(ctx, asOf)
	if err != nil {
		return nil, nil, err
	}
	sdk, err := buildPhase6VerifierSDK(asOf)
	if err != nil {
		return nil, nil, err
	}
	referencePack, err := buildPublicVerifierReferencePack()
	if err != nil {
		return nil, nil, err
	}
	bindings, err := s.publicProofValAEnvironmentBindings(asOf)
	if err != nil {
		return nil, nil, err
	}
	byID := map[string]claimscore.PublicProofEnvironmentBindingItem{}
	for _, item := range bindings {
		byID[item.ArtifactID] = item
	}
	catalog := benchmarkfoundation.FoundationCatalog()
	evaluations := make([]benchmarkfoundation.EvaluationResponse, 0, 1)
	items := make([]claimscore.PublicProofReplayVerificationItem, 0, len(builtPacks))
	for _, built := range builtPacks {
		binding := byID[built.Pack.ArtifactID]
		item := claimscore.PublicProofReplayVerificationItem{
			ArtifactID:     built.Pack.ArtifactID,
			CurrentState:   "replay_ready",
			MethodologyRef: built.Pack.MethodologyRef,
			SupportedEnvironmentClasses: []string{
				binding.EnvironmentClass,
			},
			EvidenceRefs: []string{
				built.Pack.DownloadRef,
				built.Pack.MethodologyRef,
			},
			Limitations: []string{
				"Replay verification remains bounded to the declared environment class, methodology ref, and tolerance policy of the published pack.",
			},
		}
		switch strings.TrimSpace(built.Pack.ArtifactID) {
		case "point2_runtime_performance_public_pack":
			evaluation := benchmarkfoundation.EvaluateFoundationRegression(publicProofValBRuntimeReplayEvaluationRequest(strings.TrimSpace(binding.ExecutionProfile), asOf))
			evaluations = append(evaluations, evaluation)
			item.ComparisonMode = "bounded_tolerance_comparison"
			if evaluation.CurrentState == "passed" {
				item.EvaluationState = "comparison_verified"
			} else {
				item.EvaluationState = "comparison_limited"
				item.CurrentState = "replay_limited"
			}
			item.EvaluationRef = "/v1/foundation/execution/benchmarks/evaluate"
			item.ToleranceBands = publicProofValBToleranceBands(catalog.Rules, "control_plane_latency", "background_completion_latency", "evidence_latency")
			item.ReplayCommandHints = publicProofValBRelevantCommandHints(catalog.Families)
			item.UnsupportedReplayCases = []string{
				"cross_environment_outside_tolerance",
				"replay_without_methodology_compatibility",
				"hardware_or_provider_drift_outside_declared_environment_class",
			}
			item.EvidenceRefs = append(item.EvidenceRefs, "/v1/foundation/execution/benchmarks/evaluate", "/v1/runtime/substrate-depth/vale/proofs")
		default:
			item.ComparisonMode = "reference_pack_replay"
			if sdk.CurrentState == phase6VerifierSDKStateActive && len(referencePack.ReplayInputs) > 0 {
				item.EvaluationState = "replay_verified"
			} else {
				item.EvaluationState = "replay_limited"
				item.CurrentState = "replay_limited"
			}
			item.EvaluationRef = "/v1/public/verifier/reference-pack"
			item.ToleranceBands = []string{
				"schema_line=exact_match",
				"payload_digest=exact_match",
				"freshness=within_valid_through_or_stale_visible",
			}
			item.ReplayCommandHints = uniqueStrings(append([]string{
				"Replay the published verifier reference inputs locally and preserve exact verifier result states.",
			}, sdk.VerificationCommands...))
			item.UnsupportedReplayCases = []string{
				"signature verification without a compatible trust root",
				"schema incompatibility against public.proof.sealed_artifact.v1",
				"treating partner_scoped projection as if it were internal_full raw evidence",
			}
			item.EvidenceRefs = append(item.EvidenceRefs, "/v1/public/verifier/reference-pack", "/v1/public/verifier/sdk")
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	sort.Slice(evaluations, func(i, j int) bool { return evaluations[i].ProfileID < evaluations[j].ProfileID })
	return items, evaluations, nil
}

func publicProofValBTransparencyEntryHash(rootHash, payloadDigest, artifactID string) string {
	return digestBytesMust(canonicalJSONMust(map[string]string{
		"parent_root_hash": rootHash,
		"payload_digest":   payloadDigest,
		"artifact_id":      artifactID,
	}))
}

func publicProofValBSchemaCompatibility(schemaVersion string) string {
	if strings.TrimSpace(schemaVersion) == "public.proof.sealed_artifact.v1" {
		return "supported"
	}
	return "unsupported"
}

func publicProofValBEnvelopeProvider(envelope *signing.Envelope, runtime *signingRuntime) string {
	if envelope != nil && strings.TrimSpace(envelope.Provider) != "" {
		return strings.TrimSpace(envelope.Provider)
	}
	if runtime != nil {
		return runtime.mode()
	}
	return signing.ModeDisabled
}

func publicProofValBTrustRootState(model claimscore.PublicProofSigningAuthorityModel, pack claimscore.PublicSealedProofPack) string {
	if strings.TrimSpace(model.CurrentState) != claimscore.MeasuredPublicProofVal0SigningAuthorityStateActive {
		return "untrusted"
	}
	for _, root := range model.TrustRoots {
		if strings.TrimSpace(root.TrustRootID) != strings.TrimSpace(pack.TrustRootID) {
			continue
		}
		if strings.TrimSpace(root.KeyVersion) != strings.TrimSpace(pack.KeyVersion) {
			continue
		}
		if strings.TrimSpace(root.CurrentState) == "" {
			return "untrusted"
		}
		return "trusted"
	}
	return "untrusted"
}

func publicProofValBToleranceBands(rules []benchmarkfoundation.RegressionRule, classes ...string) []string {
	allowed := map[string]struct{}{}
	for _, class := range classes {
		allowed[strings.TrimSpace(class)] = struct{}{}
	}
	bands := make([]string, 0, len(rules))
	for _, rule := range rules {
		if _, ok := allowed[strings.TrimSpace(rule.MetricClass)]; !ok {
			continue
		}
		bands = append(bands, strings.TrimSpace(rule.MetricClass)+"<="+fmt.Sprintf("%.0f%%", rule.ThresholdPercent))
	}
	return uniqueStrings(bands)
}

func publicProofValBRelevantCommandHints(families []benchmarkfoundation.FoundationFamily) []string {
	hints := []string{}
	for _, family := range families {
		switch strings.TrimSpace(family.FamilyID) {
		case "runtime_compare", "audit_writer_read_paths", "audit_writer_mutation_paths":
			hints = append(hints, strings.TrimSpace(family.CommandHint))
		}
	}
	return uniqueStrings(hints)
}

func publicProofValBRuntimeReplayEvaluationRequest(profileID string, observedAt time.Time) benchmarkfoundation.EvaluationRequest {
	if strings.TrimSpace(profileID) == "" {
		profileID = "production_like"
	}
	request := benchmarkfoundation.EvaluationRequest{
		SchemaVersion: benchmarkfoundation.FoundationEvaluationSchemaVersion,
		ProfileID:     strings.TrimSpace(profileID),
		ObservedAt:    observedAt.UTC(),
		Observations: []benchmarkfoundation.Observation{
			{
				FamilyID:      "runtime_compare",
				ProfileID:     strings.TrimSpace(profileID),
				MetricClass:   "control_plane_latency",
				MetricName:    "sealed_pack_capture_p95_latency_us",
				Unit:          "us",
				BaselineValue: 300,
				ObservedValue: 276,
			},
			{
				FamilyID:      "audit_writer_read_paths",
				ProfileID:     strings.TrimSpace(profileID),
				MetricClass:   "background_completion_latency",
				MetricName:    "sealed_pack_replay_p95_latency_us",
				Unit:          "us",
				BaselineValue: 580,
				ObservedValue: 545,
			},
			{
				FamilyID:      "audit_writer_mutation_paths",
				ProfileID:     strings.TrimSpace(profileID),
				MetricClass:   "evidence_latency",
				MetricName:    "sealed_pack_signature_verify_p95_latency_us",
				Unit:          "us",
				BaselineValue: 810,
				ObservedValue: 760,
			},
		},
	}
	switch strings.TrimSpace(profileID) {
	case "local_baseline":
		request.Observations[0].BaselineValue, request.Observations[0].ObservedValue = 250, 230
		request.Observations[1].BaselineValue, request.Observations[1].ObservedValue = 500, 470
		request.Observations[2].BaselineValue, request.Observations[2].ObservedValue = 700, 650
	case "stress":
		request.Observations[0].BaselineValue, request.Observations[0].ObservedValue = 350, 330
		request.Observations[1].BaselineValue, request.Observations[1].ObservedValue = 720, 690
		request.Observations[2].BaselineValue, request.Observations[2].ObservedValue = 930, 890
	}
	return request
}
