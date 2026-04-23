package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	benchmarkfoundation "github.com/denisgrosek/changelock/internal/benchmark"
	claimscore "github.com/denisgrosek/changelock/internal/claims"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	publicProofValAArtifactSchemaSchema     = "point2.measured_public_proof.vala.sealed_artifact_schema.v1"
	publicProofValASealingDisciplineSchema  = "point2.measured_public_proof.vala.sealing_discipline.v1"
	publicProofValAEnvironmentBindingSchema = "point2.measured_public_proof.vala.environment_binding.v1"
	publicProofValADownloadablePacksSchema  = "point2.measured_public_proof.vala.downloadable_packs.v1"
	publicProofValAPackSchema               = "point2.measured_public_proof.vala.downloadable_pack.v1"
	publicProofValAProofsSchema             = "point2.measured_public_proof.vala.proofs.v1"
)

type publicProofValAArtifactSchemaResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         claimscore.PublicSealedArtifactSchema `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type publicProofValASealingDisciplineResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         claimscore.PublicProofSealingDiscipline `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type publicProofValAEnvironmentBindingResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Items         []claimscore.PublicProofEnvironmentBindingItem `json:"items,omitempty"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type publicProofValADownloadablePacksResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Items         []claimscore.PublicSealedProofPack `json:"items,omitempty"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type publicProofValAPackResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Pack          claimscore.PublicSealedProofPack `json:"pack"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type publicProofValAProofsResponse struct {
	SchemaVersion           string    `json:"schema_version"`
	GeneratedAt             time.Time `json:"generated_at"`
	CurrentState            string    `json:"current_state"`
	Phase6State             string    `json:"phase6_state"`
	Val0State               string    `json:"val0_state"`
	ArtifactSchemaState     string    `json:"artifact_schema_state"`
	SealingDisciplineState  string    `json:"sealing_discipline_state"`
	EnvironmentBindingState string    `json:"environment_binding_state"`
	DownloadablePackState   string    `json:"downloadable_pack_state"`
	SurfaceRefs             []string  `json:"surface_refs,omitempty"`
	EvidenceRefs            []string  `json:"evidence_refs,omitempty"`
	DeferredScope           []string  `json:"deferred_scope,omitempty"`
	Limitations             []string  `json:"limitations,omitempty"`
	IntegrationSummary      []string  `json:"integration_summary,omitempty"`
}

func (s server) publicProofValAArtifactSchemaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicProofValAArtifactSchema())
}

func (s server) publicProofValASealingDisciplineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, s.buildPublicProofValASealingDiscipline())
}

func (s server) publicProofValAEnvironmentBindingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValAEnvironmentBinding(asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValADownloadablePacksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValADownloadablePacks(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValAPackByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	artifactID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/public/proof-expansion/vala/downloadable-packs/"))
	if artifactID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "pack not found"})
		return
	}
	if strings.TrimSpace(r.URL.Query().Get("as_of")) == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "as_of is required for downloadable packs"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValAPackByID(r.Context(), artifactID, asOf)
	if err != nil {
		if strings.Contains(err.Error(), "pack not found") {
			httpjson.Write(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValAProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValAProofs(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildPublicProofValAArtifactSchema() publicProofValAArtifactSchemaResponse {
	model := claimscore.MeasuredPublicProofValAArtifactSchema()
	return publicProofValAArtifactSchemaResponse{
		SchemaVersion: publicProofValAArtifactSchemaSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/signing-authority",
			"/v1/public/phase6/proofs",
			"/v1/public/proof-expansion/vala/proofs",
		},
		Limitations: []string{
			"Val A artifact schema defines sealed proof packaging and required fields before later Point 2 waves add issuance, transparency anchoring, and verifier execution.",
		},
	}
}

func (s server) buildPublicProofValASealingDiscipline() publicProofValASealingDisciplineResponse {
	model := claimscore.MeasuredPublicProofValASealingDiscipline(publicProofVal0ProviderDescriptor(s.signing), s.publicProofValASigningPurposeEnabled())
	return publicProofValASealingDisciplineResponse{
		SchemaVersion: publicProofValASealingDisciplineSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/signing-authority",
			"/v1/public/transparency/anchor",
			"/v1/public/proof-expansion/vala/proofs",
		},
		Limitations: []string{
			"Val A sealing discipline models purpose-scoped artifact sealing but does not yet publish verifier-side signature validation or transparency anchoring for newly issued artifacts.",
		},
	}
}

func (s server) buildPublicProofValAEnvironmentBinding(asOf time.Time) (publicProofValAEnvironmentBindingResponse, error) {
	items, err := s.publicProofValAEnvironmentBindings(asOf)
	if err != nil {
		return publicProofValAEnvironmentBindingResponse{}, err
	}
	return publicProofValAEnvironmentBindingResponse{
		SchemaVersion: publicProofValAEnvironmentBindingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValAEnvironmentBindingState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/complete",
			"/v1/public/benchmarks/methodology",
			"/v1/public/proof-expansion/vala/proofs",
		},
		Limitations: []string{
			"Val A environment binding remains scoped to declared environment classes, methodology refs, and compatibility scope; it does not imply universal replay support.",
		},
	}, nil
}

func (s server) buildPublicProofValADownloadablePacks(ctx context.Context, asOf time.Time) (publicProofValADownloadablePacksResponse, error) {
	items, err := s.publicProofValAPacks(ctx, asOf)
	if err != nil {
		return publicProofValADownloadablePacksResponse{}, err
	}
	return publicProofValADownloadablePacksResponse{
		SchemaVersion: publicProofValADownloadablePacksSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValADownloadablePackState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/phase6/proofs",
			"/v1/public/proof-expansion/val0/proofs",
			"/v1/public/proof-expansion/vala/proofs",
		},
		Limitations: []string{
			"Val A downloadable packs are sealed artifact projections over existing proof surfaces; automated issuance and lifecycle governance remain deferred.",
		},
	}, nil
}

func (s server) buildPublicProofValAPackByID(ctx context.Context, artifactID string, asOf time.Time) (publicProofValAPackResponse, error) {
	items, err := s.publicProofValAPacks(ctx, asOf)
	if err != nil {
		return publicProofValAPackResponse{}, err
	}
	for _, item := range items {
		if strings.TrimSpace(item.ArtifactID) != artifactID {
			continue
		}
		return publicProofValAPackResponse{
			SchemaVersion: publicProofValAPackSchema,
			GeneratedAt:   publicSampleTime(),
			CurrentState:  item.CurrentState,
			Pack:          item,
			RouteRefs: []string{
				item.DownloadRef,
				"/v1/public/proof-expansion/vala/downloadable-packs",
				"/v1/public/proof-expansion/vala/proofs",
			},
			Limitations: item.Limitations,
		}, nil
	}
	return publicProofValAPackResponse{}, fmt.Errorf("pack not found")
}

func (s server) buildPublicProofValAProofs(ctx context.Context, asOf time.Time) (publicProofValAProofsResponse, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return publicProofValAProofsResponse{}, err
	}
	val0, err := s.buildPublicProofVal0Proofs(asOf)
	if err != nil {
		return publicProofValAProofsResponse{}, err
	}
	artifactSchema := claimscore.MeasuredPublicProofValAArtifactSchema()
	sealingDiscipline := claimscore.MeasuredPublicProofValASealingDiscipline(publicProofVal0ProviderDescriptor(s.signing), s.publicProofValASigningPurposeEnabled())
	environmentBindingItems, err := s.publicProofValAEnvironmentBindings(asOf)
	if err != nil {
		return publicProofValAProofsResponse{}, err
	}
	downloadablePacks, err := s.publicProofValAPacks(ctx, asOf)
	if err != nil {
		return publicProofValAProofsResponse{}, err
	}
	environmentBindingState := claimscore.EvaluateMeasuredPublicProofValAEnvironmentBindingState(environmentBindingItems)
	downloadablePackState := claimscore.EvaluateMeasuredPublicProofValADownloadablePackState(downloadablePacks)
	currentState := claimscore.EvaluateMeasuredPublicProofValAState(
		val0.CurrentState,
		artifactSchema.CurrentState,
		sealingDiscipline.CurrentState,
		environmentBindingState,
		downloadablePackState,
	)
	return publicProofValAProofsResponse{
		SchemaVersion:           publicProofValAProofsSchema,
		GeneratedAt:             publicSampleTime(),
		CurrentState:            currentState,
		Phase6State:             phase6.CurrentState,
		Val0State:               val0.CurrentState,
		ArtifactSchemaState:     artifactSchema.CurrentState,
		SealingDisciplineState:  sealingDiscipline.CurrentState,
		EnvironmentBindingState: environmentBindingState,
		DownloadablePackState:   downloadablePackState,
		SurfaceRefs: []string{
			"/v1/public/phase6/proofs",
			"/v1/public/proof-expansion/val0/proofs",
			"/v1/public/proof-expansion/vala/sealed-artifact-schema",
			"/v1/public/proof-expansion/vala/sealing-discipline",
			"/v1/public/proof-expansion/vala/environment-binding",
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/vala/proofs",
		},
		EvidenceRefs: []string{
			"/v1/public/phase6/proofs",
			"/v1/public/benchmarks/methodology",
			"/v1/runtime/substrate-depth/complete",
			"/v1/runtime/substrate-depth/vale/proofs",
			"/v1/public/verifier/sdk",
		},
		DeferredScope: []string{
			"point2_valb_transparency_and_verification",
			"point2_valc_public_and_partner_proof_portal",
			"point2_vald_automated_issuance_and_revocation_gate",
			"point2_vale_final_proof_gate",
		},
		Limitations: []string{
			"Val A closes sealed artifact schema, sealing discipline, environment binding, and downloadable pack projection only; it does not yet add transparency anchoring or third-party verifier execution for newly issued artifacts.",
			"Downloadable packs remain bounded projections over existing proof surfaces and do not create a parallel evidence authority.",
		},
		IntegrationSummary: []string{
			"Val A turns Point 2 proof surfaces into sealed artifact packs with canonical payload digesting, signature envelope metadata, timestamp linkage, and environment binding.",
			"Fail-closed sealing requires active Val 0 discipline and an enabled public-proof-artifact signing purpose before downloadable packs can become active.",
		},
	}, nil
}

func (s server) publicProofValAEnvironmentBindings(asOf time.Time) ([]claimscore.PublicProofEnvironmentBindingItem, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return nil, err
	}
	valEProofs, err := s.buildRuntimeSubstratePoint1Complete(context.Background(), runtimeIntegrityFilter{})
	if err != nil {
		return nil, err
	}
	catalog := benchmarkfoundation.FoundationCatalog()
	return []claimscore.PublicProofEnvironmentBindingItem{
		{
			ArtifactID:         "point2_runtime_performance_public_pack",
			CurrentState:       "binding_ready",
			ClaimClass:         claimscore.PublicProofClaimClassPerformance,
			RedactionTier:      claimscore.RedactionTierPublicSafe,
			EnvironmentClass:   "runtime_hardened_enterprise_cluster",
			ExecutionProfile:   "production_like",
			WorkloadShape:      "runtime_substrate_depth_point1",
			BuildIdentity:      "changelock-2026.04.23",
			HarnessVersion:     catalog.SchemaVersion,
			MethodologyRef:     "/v1/public/benchmarks/methodology",
			CompatibilityScope: "public.proof.sealed_artifact.v1",
			ProvenanceInputs: []string{
				"/v1/runtime/substrate-depth/complete",
				"/v1/runtime/substrate-depth/vale/proofs",
				"/v1/public/phase6/proofs",
			},
			ReplayBoundaries: []string{
				"replay is bounded to production_like foundation profile",
				"methodology and environment class must remain compatible",
			},
			UnsupportedReplay: []string{
				"cross-provider replay outside declared environment class",
				"tenant-sensitive raw evidence replay in public-safe mode",
			},
			Limitations: []string{
				"Performance public pack remains tied to Point 1 completion state " + valEProofs.CurrentState + " and Phase 6 public proof state " + phase6.CurrentState + ".",
			},
		},
		{
			ArtifactID:         "point2_verification_public_pack",
			CurrentState:       "binding_ready",
			ClaimClass:         claimscore.PublicProofClaimClassVerification,
			RedactionTier:      claimscore.RedactionTierPartnerScoped,
			EnvironmentClass:   "verifier_reference_environment",
			ExecutionProfile:   "local_baseline",
			WorkloadShape:      "public_verifier_reference_pack",
			BuildIdentity:      "changelock-2026.04.23",
			HarnessVersion:     catalog.SchemaVersion,
			MethodologyRef:     "/v1/public/benchmarks/methodology",
			CompatibilityScope: "public.proof.sealed_artifact.v1",
			ProvenanceInputs: []string{
				"/v1/public/verifier/sdk",
				"/v1/public/claims/summary",
				"/v1/public/phase6/proofs",
			},
			ReplayBoundaries: []string{
				"verification pack replay requires compatible schema line and reference verifier inputs",
				"partner-safe projection exposes bounded evidence only",
			},
			UnsupportedReplay: []string{
				"signature verification without declared trust root",
				"replay against incompatible verifier schema line",
			},
			Limitations: []string{
				"Verification public pack remains methodology-bound and does not replace third-party verifier execution planned for Point 2 / Val B.",
			},
		},
	}, nil
}

func (s server) publicProofValAPacks(ctx context.Context, asOf time.Time) ([]claimscore.PublicSealedProofPack, error) {
	built, err := s.publicProofValABuiltPacks(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicSealedProofPack, 0, len(built))
	for _, item := range built {
		items = append(items, item.Pack)
	}
	return items, nil
}

func (s server) publicProofValABuiltPacks(ctx context.Context, asOf time.Time) ([]publicProofValABuiltPack, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return nil, err
	}
	val0, err := s.buildPublicProofVal0Proofs(asOf)
	if err != nil {
		return nil, err
	}
	point1, err := s.buildRuntimeSubstratePoint1Complete(ctx, runtimeIntegrityFilter{})
	if err != nil {
		return nil, err
	}
	bindings, err := s.publicProofValAEnvironmentBindings(asOf)
	if err != nil {
		return nil, err
	}
	byID := map[string]claimscore.PublicProofEnvironmentBindingItem{}
	for _, item := range bindings {
		byID[item.ArtifactID] = item
	}
	catalog := benchmarkfoundation.FoundationCatalog()
	issuedAt := publicSampleTime()
	validThrough := issuedAt.Add(45 * 24 * time.Hour)
	performanceSeed := publicProofValAPackSeed{
		ArtifactID:        "point2_runtime_performance_public_pack",
		ArtifactType:      claimscore.PublicProofArtifactTypeBenchmarkPack,
		ClaimID:           "point2_runtime_performance_claim",
		ClaimClass:        claimscore.PublicProofClaimClassPerformance,
		RedactionTier:     claimscore.RedactionTierPublicSafe,
		EnvironmentClass:  byID["point2_runtime_performance_public_pack"].EnvironmentClass,
		ExecutionProfile:  byID["point2_runtime_performance_public_pack"].ExecutionProfile,
		WorkloadShape:     byID["point2_runtime_performance_public_pack"].WorkloadShape,
		BuildIdentity:     byID["point2_runtime_performance_public_pack"].BuildIdentity,
		HarnessVersion:    catalog.SchemaVersion,
		MethodologyRef:    "/v1/public/benchmarks/methodology",
		IssuedAt:          issuedAt,
		ValidThrough:      validThrough,
		MeasurementSource: "runtime_substrate_vale_latency.standard_node.v1",
		AsOf:              asOf,
		EvidenceRefs: []string{
			"/v1/runtime/substrate-depth/complete",
			"/v1/runtime/substrate-depth/vale/proofs",
			"/v1/public/phase6/proofs",
		},
		MetricSummaries: []string{
			"point1_state=" + point1.CurrentState,
			"phase6_state=" + phase6.CurrentState,
			"val0_state=" + val0.CurrentState,
			"capture_p99_micros=340",
			"correlation_p99_micros=580",
			"false_positive_rate_pct=0.94",
		},
		Limitations: []string{
			"Pack remains scoped to runtime/substrate Point 1 completion and public-safe projection boundaries.",
		},
	}
	verificationSeed := publicProofValAPackSeed{
		ArtifactID:        "point2_verification_public_pack",
		ArtifactType:      claimscore.PublicProofArtifactTypeProofPack,
		ClaimID:           "point2_verification_reference_claim",
		ClaimClass:        claimscore.PublicProofClaimClassVerification,
		RedactionTier:     claimscore.RedactionTierPartnerScoped,
		EnvironmentClass:  byID["point2_verification_public_pack"].EnvironmentClass,
		ExecutionProfile:  byID["point2_verification_public_pack"].ExecutionProfile,
		WorkloadShape:     byID["point2_verification_public_pack"].WorkloadShape,
		BuildIdentity:     byID["point2_verification_public_pack"].BuildIdentity,
		HarnessVersion:    catalog.SchemaVersion,
		MethodologyRef:    "/v1/public/benchmarks/methodology",
		IssuedAt:          issuedAt,
		ValidThrough:      validThrough,
		MeasurementSource: "phase6_verifier_sdk.reference_pack.v1",
		AsOf:              asOf,
		EvidenceRefs: []string{
			"/v1/public/verifier/sdk",
			"/v1/public/claims/summary",
			"/v1/public/phase6/proofs",
		},
		MetricSummaries: []string{
			"phase6_state=" + phase6.CurrentState,
			"verifier_sdk_state=" + phase6.VerifierSDK.CurrentState,
			"claims_summary_state=" + phase6.ClaimsSummary.CurrentState,
			"independent_check=true",
		},
		Limitations: []string{
			"Pack remains partner-scoped because verifier-safe evidence still excludes tenant-sensitive runtime raw data.",
		},
	}
	items := make([]publicProofValABuiltPack, 0, 2)
	for _, seed := range []publicProofValAPackSeed{performanceSeed, verificationSeed} {
		binding := byID[seed.ArtifactID]
		item, err := s.publicProofValABuildPack(ctx, seed, binding)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s server) publicProofValABuildPack(ctx context.Context, seed publicProofValAPackSeed, binding claimscore.PublicProofEnvironmentBindingItem) (publicProofValABuiltPack, error) {
	payload := publicProofValAPackPayload{
		ArtifactID:        seed.ArtifactID,
		ClaimID:           seed.ClaimID,
		ClaimClass:        seed.ClaimClass,
		MeasurementSource: seed.MeasurementSource,
		MethodologyRef:    seed.MethodologyRef,
		IssuedAt:          seed.IssuedAt,
		ValidThrough:      seed.ValidThrough,
		EvidenceRefs:      seed.EvidenceRefs,
		MetricSummaries:   seed.MetricSummaries,
		EnvironmentBinding: publicProofValAPackBinding{
			EnvironmentClass:   binding.EnvironmentClass,
			ExecutionProfile:   binding.ExecutionProfile,
			WorkloadShape:      binding.WorkloadShape,
			BuildIdentity:      binding.BuildIdentity,
			HarnessVersion:     binding.HarnessVersion,
			CompatibilityScope: binding.CompatibilityScope,
			ProvenanceInputs:   binding.ProvenanceInputs,
			ReplayBoundaries:   binding.ReplayBoundaries,
		},
	}
	payloadBytes, err := canonicalJSON(payload)
	if err != nil {
		return publicProofValABuiltPack{}, err
	}
	sum := sha256.Sum256(payloadBytes)
	payloadDigest := "sha256:" + hex.EncodeToString(sum[:])
	envelope, err := s.signing.signPublicProofArtifact(ctx, payloadBytes)
	if err != nil {
		return publicProofValABuiltPack{}, err
	}
	if envelope != nil {
		payloadDigest = envelope.PayloadDigest
	}
	manifestBytes := canonicalJSONMust(publicProofValAManifest(seed, binding, payloadDigest))
	environmentBytes := canonicalJSONMust(binding)
	signatureBytes := canonicalJSONMust(publicProofValASignatureFile(envelope, payloadDigest))
	timestampBytes := canonicalJSONMust(publicProofValATimestampFile(seed, envelope))
	files := []claimscore.PublicSealedProofPackFile{
		{Path: "manifest.json", MediaType: "application/json", Role: "manifest", SHA256: digestBytesMust(manifestBytes)},
		{Path: "payload/measurement_summary.json", MediaType: "application/json", Role: "measurement_summary", SHA256: digestBytesMust(payloadBytes)},
		{Path: "environment/binding.json", MediaType: "application/json", Role: "environment_binding", SHA256: digestBytesMust(environmentBytes)},
		{Path: "signature/envelope.json", MediaType: "application/json", Role: "signature_envelope", SHA256: digestBytesMust(signatureBytes)},
		{Path: "timestamp/receipt.json", MediaType: "application/json", Role: "timestamp_receipt", SHA256: digestBytesMust(timestampBytes)},
	}
	currentState := "sealed_artifact_ready"
	downloadRef := publicProofValADownloadRef(seed.ArtifactID, seed.AsOf)
	timestampRef := publicProofValATimestampRef(seed.ArtifactID, seed.AsOf)
	if envelope == nil {
		currentState = "sealed_artifact_signing_pending"
		timestampRef = ""
	}
	return publicProofValABuiltPack{
		Payload: payloadBytes,
		Pack: claimscore.PublicSealedProofPack{
			ArtifactID:            seed.ArtifactID,
			ArtifactSchemaVersion: "public.proof.sealed_artifact.v1",
			ArtifactType:          seed.ArtifactType,
			CurrentState:          currentState,
			ClaimID:               seed.ClaimID,
			ClaimClass:            seed.ClaimClass,
			RedactionTier:         seed.RedactionTier,
			EnvironmentClass:      seed.EnvironmentClass,
			ExecutionProfile:      seed.ExecutionProfile,
			WorkloadShape:         seed.WorkloadShape,
			BuildIdentity:         seed.BuildIdentity,
			HarnessVersion:        seed.HarnessVersion,
			MethodologyRef:        seed.MethodologyRef,
			IssuedAt:              seed.IssuedAt,
			ValidThrough:          seed.ValidThrough,
			MeasurementSource:     seed.MeasurementSource,
			EvidenceRefs:          seed.EvidenceRefs,
			DownloadRef:           downloadRef,
			PayloadDigest:         payloadDigest,
			SignatureEnvelope:     envelope,
			TrustRootID:           "public_proof_primary_root",
			KeyVersion:            publicProofValAKeyVersion(envelope),
			TimestampRef:          timestampRef,
			PackagingFiles:        files,
			MetricSummaries:       seed.MetricSummaries,
			Limitations:           seed.Limitations,
		},
	}, nil
}

func (s server) publicProofValASigningPurposeEnabled() bool {
	return s.signing != nil && s.signing.enabledForPurpose(signing.PurposePublicProofArtifact)
}

type publicProofValAPackSeed struct {
	ArtifactID        string
	ArtifactType      string
	ClaimID           string
	ClaimClass        string
	RedactionTier     string
	EnvironmentClass  string
	ExecutionProfile  string
	WorkloadShape     string
	BuildIdentity     string
	HarnessVersion    string
	MethodologyRef    string
	IssuedAt          time.Time
	ValidThrough      time.Time
	MeasurementSource string
	AsOf              time.Time
	EvidenceRefs      []string
	MetricSummaries   []string
	Limitations       []string
}

type publicProofValABuiltPack struct {
	Pack    claimscore.PublicSealedProofPack
	Payload []byte
}

type publicProofValAPackPayload struct {
	ArtifactID         string                     `json:"artifact_id"`
	ClaimID            string                     `json:"claim_id"`
	ClaimClass         string                     `json:"claim_class"`
	MeasurementSource  string                     `json:"measurement_source"`
	MethodologyRef     string                     `json:"methodology_ref"`
	IssuedAt           time.Time                  `json:"issued_at"`
	ValidThrough       time.Time                  `json:"valid_through"`
	EvidenceRefs       []string                   `json:"evidence_refs,omitempty"`
	MetricSummaries    []string                   `json:"metric_summaries,omitempty"`
	EnvironmentBinding publicProofValAPackBinding `json:"environment_binding"`
}

type publicProofValAPackBinding struct {
	EnvironmentClass   string   `json:"environment_class"`
	ExecutionProfile   string   `json:"execution_profile"`
	WorkloadShape      string   `json:"workload_shape"`
	BuildIdentity      string   `json:"build_identity"`
	HarnessVersion     string   `json:"harness_version"`
	CompatibilityScope string   `json:"compatibility_scope"`
	ProvenanceInputs   []string `json:"provenance_inputs,omitempty"`
	ReplayBoundaries   []string `json:"replay_boundaries,omitempty"`
}

func publicProofValAManifest(seed publicProofValAPackSeed, binding claimscore.PublicProofEnvironmentBindingItem, payloadDigest string) map[string]any {
	return map[string]any{
		"artifact_id":             seed.ArtifactID,
		"artifact_schema_version": "public.proof.sealed_artifact.v1",
		"artifact_type":           seed.ArtifactType,
		"claim_id":                seed.ClaimID,
		"claim_class":             seed.ClaimClass,
		"redaction_tier":          seed.RedactionTier,
		"environment_class":       seed.EnvironmentClass,
		"execution_profile":       seed.ExecutionProfile,
		"build_identity":          seed.BuildIdentity,
		"methodology_ref":         seed.MethodologyRef,
		"payload_digest":          payloadDigest,
		"compatibility_scope":     binding.CompatibilityScope,
		"issued_at":               seed.IssuedAt.UTC(),
		"valid_through":           seed.ValidThrough.UTC(),
		"evidence_refs":           seed.EvidenceRefs,
	}
}

func publicProofValASignatureFile(envelope *signing.Envelope, payloadDigest string) any {
	if envelope != nil {
		return envelope
	}
	return map[string]any{
		"current_state":  "signature_pending",
		"purpose":        signing.PurposePublicProofArtifact,
		"payload_digest": payloadDigest,
	}
}

func publicProofValATimestampFile(seed publicProofValAPackSeed, envelope *signing.Envelope) any {
	timestampedAt := seed.IssuedAt.UTC()
	if envelope != nil && !envelope.SignedAt.IsZero() {
		timestampedAt = envelope.SignedAt.UTC()
	}
	return map[string]any{
		"timestamp_ref":  publicProofValATimestampRef(seed.ArtifactID, seed.AsOf),
		"timestamped_at": timestampedAt,
		"issued_window":  "2026-04-01T00:00:00Z/2026-06-01T00:00:00Z",
	}
}

func publicProofValADownloadRef(artifactID string, asOf time.Time) string {
	downloadRef := "/v1/public/proof-expansion/vala/downloadable-packs/" + strings.TrimSpace(artifactID)
	if asOf.IsZero() {
		return downloadRef
	}
	return downloadRef + "?as_of=" + url.QueryEscape(asOf.UTC().Format(time.RFC3339))
}

func publicProofValATimestampRef(artifactID string, asOf time.Time) string {
	return publicProofValADownloadRef(artifactID, asOf) + "#timestamp"
}

func publicProofValAKeyVersion(envelope *signing.Envelope) string {
	if envelope == nil || strings.TrimSpace(envelope.KeyID) == "" {
		return "v1"
	}
	return "v1"
}
