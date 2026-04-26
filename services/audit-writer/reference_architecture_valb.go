package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	referenceArchitectureValBPackRegistrySchema = "point6.reference_architecture.valb.pack_registry.v1"
	referenceArchitectureValBBundlesSchema      = "point6.reference_architecture.valb.bundles.v1"
	referenceArchitectureValBManifestSchema     = "point6.reference_architecture.valb.artifact_manifests.v1"
	referenceArchitectureValBReadinessSchema    = "point6.reference_architecture.valb.readiness_checks.v1"
	referenceArchitectureValBHookSchema         = "point6.reference_architecture.valb.validation_hooks.v1"
	referenceArchitectureValBConformanceSchema  = "point6.reference_architecture.valb.conformance_kit.v1"
	referenceArchitectureValBDeviationSchema    = "point6.reference_architecture.valb.deviations.v1"
	referenceArchitectureValBProofsSchema       = "point6.reference_architecture.valb.proofs.v1"
)

type referenceArchitectureValBFamilyStatus struct {
	Family                 string `json:"family"`
	PackID                 string `json:"pack_id"`
	PackState              string `json:"pack_state"`
	ManifestState          string `json:"artifact_manifest_state"`
	BundleState            string `json:"bundle_state"`
	ReadinessState         string `json:"readiness_state"`
	HookState              string `json:"validation_hook_state"`
	ConformanceState       string `json:"conformance_kit_state"`
	DeviationState         string `json:"deviation_state"`
	RequiredArtifactCount  int    `json:"required_artifact_count"`
	RequiredReadinessCount int    `json:"required_readiness_check_count"`
	RequiredHookCount      int    `json:"required_hook_count"`
	BlockingDeviationCount int    `json:"blocking_deviation_count"`
}

type referenceArchitectureValBPackRegistryResponse struct {
	SchemaVersion string                                                 `json:"schema_version"`
	GeneratedAt   time.Time                                              `json:"generated_at"`
	CurrentState  string                                                 `json:"current_state"`
	Model         operability.ReferenceArchitectureBlueprintPackRegistry `json:"model"`
	FamilyStates  []referenceArchitectureValBFamilyStatus                `json:"family_states,omitempty"`
	RouteRefs     []string                                               `json:"route_refs,omitempty"`
	Limitations   []string                                               `json:"limitations,omitempty"`
}

type referenceArchitectureValBCollectionResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	FamilyStates  []referenceArchitectureValBFamilyStatus `json:"family_states,omitempty"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
	Model         any                                     `json:"model"`
}

type referenceArchitectureValBProofsResponse struct {
	SchemaVersion         string                                  `json:"schema_version"`
	GeneratedAt           time.Time                               `json:"generated_at"`
	CurrentState          string                                  `json:"current_state"`
	Point5DependencyState string                                  `json:"point_5_dependency_state"`
	Point5State           string                                  `json:"point_5_state"`
	Val0DependencyState   string                                  `json:"val_0_dependency_state"`
	Val0State             string                                  `json:"val_0_state"`
	ValADependencyState   string                                  `json:"val_a_dependency_state"`
	ValAState             string                                  `json:"val_a_state"`
	ValBState             string                                  `json:"val_b_state"`
	Point6State           string                                  `json:"point_6_state"`
	PackRegistryState     string                                  `json:"pack_registry_state"`
	BundleState           string                                  `json:"bundle_state"`
	ArtifactManifestState string                                  `json:"artifact_manifest_state"`
	ReadinessState        string                                  `json:"readiness_state"`
	ValidationHookState   string                                  `json:"validation_hook_state"`
	ConformanceKitState   string                                  `json:"conformance_kit_state"`
	DeviationState        string                                  `json:"deviation_classifier_state"`
	SupportedFamilies     []string                                `json:"supported_blueprint_families,omitempty"`
	FamilyStates          []referenceArchitectureValBFamilyStatus `json:"family_states,omitempty"`
	WhyPoint6NotPass      []string                                `json:"why_point_6_not_pass,omitempty"`
	SurfaceRefs           []string                                `json:"surface_refs,omitempty"`
	EvidenceRefs          []string                                `json:"evidence_refs,omitempty"`
	Limitations           []string                                `json:"limitations,omitempty"`
	ProjectionDisclaimer  string                                  `json:"projection_disclaimer"`
	IntegrationSummary    []string                                `json:"integration_summary,omitempty"`
}

func referenceArchitectureValBAllSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/pack-registry",
		"/v1/reference-architecture/valb/bundles",
		"/v1/reference-architecture/valb/artifact-manifests",
		"/v1/reference-architecture/valb/readiness-checks",
		"/v1/reference-architecture/valb/validation-hooks",
		"/v1/reference-architecture/valb/conformance-kit",
		"/v1/reference-architecture/valb/deviations",
		"/v1/reference-architecture/valb/proofs",
	}
}

func referenceArchitectureValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_blueprint_as_code_validation"
}

func referenceArchitectureValBEvidenceRefs(
	registry operability.ReferenceArchitectureBlueprintPackRegistry,
	kits operability.ReferenceArchitectureConformanceKitCollection,
) []string {
	refs := []string{"point5_integrated_closure", "point6_val0_proofs", "point6_vala_proofs", registry.RegistryID, kits.CollectionID}
	for _, pack := range registry.Packs {
		if pack.PackID != "" {
			refs = append(refs, pack.PackID)
		}
		for _, evidence := range pack.EvidenceRefs {
			if evidence.EvidenceID != "" {
				refs = append(refs, evidence.EvidenceID)
			}
		}
	}
	return refs
}

func referenceArchitectureValBCountBlockingDeviations(report operability.ReferenceArchitectureDeviationReport) int {
	count := 0
	for _, deviation := range report.Deviations {
		if deviation.BlocksMatched {
			count++
		}
	}
	return count
}

func buildReferenceArchitectureValBFamilyStatuses(
	registry operability.ReferenceArchitectureBlueprintPackRegistry,
	manifests operability.ReferenceArchitectureArtifactManifestCollection,
	bundles operability.ReferenceArchitectureBundleCollection,
	readiness operability.ReferenceArchitectureReadinessCollection,
	hooks operability.ReferenceArchitectureValidationHookCollection,
	deviations operability.ReferenceArchitectureDeviationCollection,
	kits operability.ReferenceArchitectureConformanceKitCollection,
) []referenceArchitectureValBFamilyStatus {
	manifestsByFamily := map[string]operability.ReferenceArchitectureArtifactManifest{}
	for _, manifest := range manifests.Manifests {
		manifestsByFamily[manifest.BlueprintFamily] = manifest
	}
	bundlesByFamily := map[string]operability.ReferenceArchitectureConfigProfilePolicyBundle{}
	for _, bundle := range bundles.Bundles {
		bundlesByFamily[bundle.BlueprintFamily] = bundle
	}
	readinessByFamily := map[string]operability.ReferenceArchitectureReadinessBundle{}
	for _, bundle := range readiness.Bundles {
		readinessByFamily[bundle.BlueprintFamily] = bundle
	}
	hooksByFamily := map[string]operability.ReferenceArchitectureValidationHookPack{}
	for _, hookPack := range hooks.HookPacks {
		hooksByFamily[hookPack.BlueprintFamily] = hookPack
	}
	deviationsByFamily := map[string]operability.ReferenceArchitectureDeviationReport{}
	for _, report := range deviations.Reports {
		deviationsByFamily[report.BlueprintFamily] = report
	}
	kitsByFamily := map[string]operability.ReferenceArchitectureConformanceKit{}
	for _, kit := range kits.Kits {
		kitsByFamily[kit.BlueprintFamily] = kit
	}

	statuses := make([]referenceArchitectureValBFamilyStatus, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		manifest := manifestsByFamily[pack.BlueprintFamily]
		bundle := bundlesByFamily[pack.BlueprintFamily]
		readinessBundle := readinessByFamily[pack.BlueprintFamily]
		hookPack := hooksByFamily[pack.BlueprintFamily]
		report := deviationsByFamily[pack.BlueprintFamily]
		kit := kitsByFamily[pack.BlueprintFamily]
		packState := operability.EvaluateReferenceArchitectureValBPackState(pack)
		manifestState := operability.EvaluateReferenceArchitectureValBArtifactManifestState(manifest)
		bundleState := operability.EvaluateReferenceArchitectureValBBundleState(bundle)
		readinessState := operability.EvaluateReferenceArchitectureValBReadinessBundleState(readinessBundle)
		hookState := operability.EvaluateReferenceArchitectureValBHookPackState(hookPack)
		deviationState := operability.EvaluateReferenceArchitectureValBDeviationReportState(report)
		conformanceState := operability.EvaluateReferenceArchitectureValBConformanceKitState(
			packState,
			manifestState,
			bundleState,
			readinessState,
			hookState,
			deviationState,
			kit,
			pack,
			report,
		)
		requiredArtifacts := 0
		for _, artifact := range manifest.Artifacts {
			if artifact.RequirementLevel == operability.ReferenceArchitectureValBArtifactRequired {
				requiredArtifacts++
			}
		}
		statuses = append(statuses, referenceArchitectureValBFamilyStatus{
			Family:                 pack.BlueprintFamily,
			PackID:                 pack.PackID,
			PackState:              packState,
			ManifestState:          manifestState,
			BundleState:            bundleState,
			ReadinessState:         readinessState,
			HookState:              hookState,
			ConformanceState:       conformanceState,
			DeviationState:         deviationState,
			RequiredArtifactCount:  requiredArtifacts,
			RequiredReadinessCount: len(readinessBundle.Checks),
			RequiredHookCount:      len(hookPack.Hooks),
			BlockingDeviationCount: referenceArchitectureValBCountBlockingDeviations(report),
		})
	}
	return statuses
}

func (s server) referenceArchitectureValBPackRegistryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBPackRegistry())
}

func (s server) referenceArchitectureValBBundlesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBBundles())
}

func (s server) referenceArchitectureValBArtifactManifestHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBArtifactManifests())
}

func (s server) referenceArchitectureValBReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBReadiness())
}

func (s server) referenceArchitectureValBValidationHooksHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBValidationHooks())
}

func (s server) referenceArchitectureValBConformanceKitHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBConformanceKit())
}

func (s server) referenceArchitectureValBDeviationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBDeviations())
}

func (s server) referenceArchitectureValBProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValBProofs())
}

func buildReferenceArchitectureValBPackRegistry() referenceArchitectureValBPackRegistryResponse {
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()
	return referenceArchitectureValBPackRegistryResponse{
		SchemaVersion: referenceArchitectureValBPackRegistrySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValBPackRegistryState(registry),
		Model:         registry,
		FamilyStates:  buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		RouteRefs: []string{
			"/v1/reference-architecture/vala/proofs",
			"/v1/reference-architecture/valb/artifact-manifests",
			"/v1/reference-architecture/valb/proofs",
		},
		Limitations: []string{
			"Val B defines bounded blueprint-as-code delivery pack contracts only and does not provision infrastructure.",
			"Pack registry remains advisory and does not approve deployment or mutate canonical truth.",
		},
	}
}

func buildReferenceArchitectureValBBundles() referenceArchitectureValBCollectionResponse {
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()
	return referenceArchitectureValBCollectionResponse{
		SchemaVersion: referenceArchitectureValBBundlesSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValBBundleCollectionState(bundles),
		Model:         bundles,
		FamilyStates:  buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/pack-registry",
			"/v1/reference-architecture/valb/proofs",
		},
		Limitations: []string{
			"Bundle contracts describe expected config, profile, policy, and evidence requirements only.",
			"Val B does not generate Terraform, Helm, or production deployment recipes.",
		},
	}
}

func buildReferenceArchitectureValBArtifactManifests() referenceArchitectureValBCollectionResponse {
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()
	return referenceArchitectureValBCollectionResponse{
		SchemaVersion: referenceArchitectureValBManifestSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValBArtifactManifestCollectionState(manifests),
		Model:         manifests,
		FamilyStates:  buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/pack-registry",
			"/v1/reference-architecture/valb/proofs",
		},
		Limitations: []string{
			"Artifact manifests remain bounded delivery metadata with integrity placeholders and freshness checks.",
			"Missing required artifacts fail closed and optional artifacts do not create deployment authority.",
		},
	}
}

func buildReferenceArchitectureValBReadiness() referenceArchitectureValBCollectionResponse {
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()
	return referenceArchitectureValBCollectionResponse{
		SchemaVersion: referenceArchitectureValBReadinessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValBReadinessCollectionState(readiness),
		Model:         readiness,
		FamilyStates:  buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/bundles",
			"/v1/reference-architecture/valb/proofs",
		},
		Limitations: []string{
			"Readiness checks are bounded pre-flight indicators and not production deployment approval.",
			"Stale or malformed readiness evidence cannot return ready.",
		},
	}
}

func buildReferenceArchitectureValBValidationHooks() referenceArchitectureValBCollectionResponse {
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()
	return referenceArchitectureValBCollectionResponse{
		SchemaVersion: referenceArchitectureValBHookSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValBValidationHookCollectionState(hooks),
		Model:         hooks,
		FamilyStates:  buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/readiness-checks",
			"/v1/reference-architecture/valb/proofs",
		},
		Limitations: []string{
			"Validation hooks are descriptors only; Val B does not execute resilience, chaos, or scale scenarios.",
			"Unknown or incomplete validation hook descriptors fail closed.",
		},
	}
}

func buildReferenceArchitectureValBConformanceKit() referenceArchitectureValBCollectionResponse {
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()
	return referenceArchitectureValBCollectionResponse{
		SchemaVersion: referenceArchitectureValBConformanceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readiness, hooks, deviations),
		Model:         kits,
		FamilyStates:  buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/pack-registry",
			"/v1/reference-architecture/valb/deviations",
			"/v1/reference-architecture/valb/proofs",
		},
		Limitations: []string{
			"Conformance kit output remains evidence-linked advisory projection and not production deployment approval.",
			"Matched requires fresh evidence, required artifacts, readiness checks, hook descriptors, and absence of blocking deviations.",
		},
	}
}

func buildReferenceArchitectureValBDeviations() referenceArchitectureValBCollectionResponse {
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()
	return referenceArchitectureValBCollectionResponse{
		SchemaVersion: referenceArchitectureValBDeviationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValBDeviationCollectionState(deviations),
		Model:         deviations,
		FamilyStates:  buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/conformance-kit",
			"/v1/reference-architecture/valb/proofs",
		},
		Limitations: []string{
			"Deviation reports classify conformance drift and blocking conditions with evidence linkage.",
			"Deviation classification remains fail-closed and does not mutate canonical state.",
		},
	}
}

func buildReferenceArchitectureValBProofs() referenceArchitectureValBProofsResponse {
	val0 := buildReferenceArchitectureVal0Proofs()
	valA := buildReferenceArchitectureValAProofs()
	registry := operability.ReferenceArchitectureValBPackRegistry()
	manifests := operability.ReferenceArchitectureValBArtifactManifestCollection()
	bundles := operability.ReferenceArchitectureValBBundleCollection()
	readiness := operability.ReferenceArchitectureValBReadinessCollection()
	hooks := operability.ReferenceArchitectureValBValidationHookCollection()
	deviations := operability.ReferenceArchitectureValBDeviationCollection()
	kits := operability.ReferenceArchitectureValBConformanceKitCollection()

	packState := operability.EvaluateReferenceArchitectureValBPackRegistryState(registry)
	manifestState := operability.EvaluateReferenceArchitectureValBArtifactManifestCollectionState(manifests)
	bundleState := operability.EvaluateReferenceArchitectureValBBundleCollectionState(bundles)
	readinessState := operability.EvaluateReferenceArchitectureValBReadinessCollectionState(readiness)
	hookState := operability.EvaluateReferenceArchitectureValBValidationHookCollectionState(hooks)
	deviationState := operability.EvaluateReferenceArchitectureValBDeviationCollectionState(deviations)
	conformanceState := operability.EvaluateReferenceArchitectureValBConformanceKitCollectionState(kits, registry, manifests, bundles, readiness, hooks, deviations)

	valBState := operability.EvaluateReferenceArchitectureValBState(
		val0.Point5State,
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		val0.Point6State,
		packState,
		manifestState,
		bundleState,
		readinessState,
		hookState,
		conformanceState,
		deviationState,
	)

	surfaceRefs := referenceArchitectureValBAllSurfaceRefs()
	evidenceRefs := referenceArchitectureValBEvidenceRefs(registry, kits)
	limitations := []string{
		"Val B defines bounded blueprint-as-code packs, manifests, readiness contracts, validation hook descriptors, conformance kits, and deviation classifiers only.",
		"Val B does not provision infrastructure, execute resilience or load scenarios, or create final reference architecture closure.",
		"Točka 6 remains not complete until later Val C through Val E layers are implemented.",
	}
	whyPoint6NotPass := []string{
		"Val B adds blueprint-as-code and validation contracts only and does not implement resilience hardening, final gate, or integrated closure.",
		"point_6_pass remains reserved for Val E integrated closure.",
	}
	currentState := operability.EvaluateReferenceArchitectureValBProofsState(
		valBState,
		operability.ReferenceArchitecturePoint6StateNotComplete,
		registry.SupportedFamilies,
		surfaceRefs,
		evidenceRefs,
		limitations,
		referenceArchitectureValBProjectionDisclaimer(),
	)

	return referenceArchitectureValBProofsResponse{
		SchemaVersion:         referenceArchitectureValBProofsSchema,
		GeneratedAt:           publicSampleTime(),
		CurrentState:          currentState,
		Point5DependencyState: val0.Point5DependencyState,
		Point5State:           val0.Point5State,
		Val0DependencyState:   val0.CurrentState,
		Val0State:             val0.Val0State,
		ValADependencyState:   valA.CurrentState,
		ValAState:             valA.ValAState,
		ValBState:             valBState,
		Point6State:           operability.ReferenceArchitecturePoint6StateNotComplete,
		PackRegistryState:     packState,
		BundleState:           bundleState,
		ArtifactManifestState: manifestState,
		ReadinessState:        readinessState,
		ValidationHookState:   hookState,
		ConformanceKitState:   conformanceState,
		DeviationState:        deviationState,
		SupportedFamilies:     registry.SupportedFamilies,
		FamilyStates:          buildReferenceArchitectureValBFamilyStatuses(registry, manifests, bundles, readiness, hooks, deviations, kits),
		WhyPoint6NotPass:      whyPoint6NotPass,
		SurfaceRefs:           surfaceRefs,
		EvidenceRefs:          evidenceRefs,
		Limitations:           limitations,
		ProjectionDisclaimer:  referenceArchitectureValBProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val B packages Val A families into bounded blueprint-as-code delivery packs, manifests, readiness checks, validation hook descriptors, conformance kits, and deviation reports.",
			"Blueprint-as-code remains advisory and evidence-linked; it does not provision infrastructure or approve production deployment.",
			"Točka 6 remains not complete until later vals add resilience hardening, operational visibility, and integrated closure.",
		},
	}
}
