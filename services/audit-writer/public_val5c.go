package main

import (
	"net/http"

	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	publicBenchmarkMethodologySchema           = "5c.public_benchmark_methodology.v1"
	publicBenchmarkSetSchema                   = "5c.public_benchmark_set.v1"
	publicAnalyticsPublicationDisciplineSchema = "5c.public_trust_analytics_publication_discipline.v1"
	publicCaseStudyPacksSchema                 = "5c.public_case_study_packs.v1"
)

type publicBenchmarkMethodologyResponse struct {
	SchemaVersion          string   `json:"schema_version"`
	MethodologyID          string   `json:"methodology_id"`
	InputSizeDiscipline    []string `json:"input_size_discipline,omitempty"`
	EnvironmentAssumptions []string `json:"environment_assumptions,omitempty"`
	WorkloadProfiles       []string `json:"workload_profiles,omitempty"`
	SubstrateContext       []string `json:"substrate_context,omitempty"`
	RepeatabilityRules     []string `json:"repeatability_rules,omitempty"`
	VariabilityDisclosure  []string `json:"variability_disclosure,omitempty"`
	FailureCaseHandling    []string `json:"failure_case_handling,omitempty"`
	NotMeasured            []string `json:"not_measured,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type publicBenchmarkDefinition struct {
	BenchmarkID            string   `json:"benchmark_id"`
	DisplayName            string   `json:"display_name"`
	MetricClass            string   `json:"metric_class"`
	MeasurementScope       []string `json:"measurement_scope,omitempty"`
	InputShape             []string `json:"input_shape,omitempty"`
	EnvironmentAssumptions []string `json:"environment_assumptions,omitempty"`
	PublicationStatus      string   `json:"publication_status"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	NotClaimed             []string `json:"not_claimed,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type publicBenchmarkSetResponse struct {
	SchemaVersion  string                      `json:"schema_version"`
	MethodologyRef string                      `json:"methodology_ref"`
	Benchmarks     []publicBenchmarkDefinition `json:"benchmarks,omitempty"`
	Limitations    []string                    `json:"limitations,omitempty"`
}

type publicAnalyticsPublicationDisciplineResponse struct {
	SchemaVersion          string   `json:"schema_version"`
	PublicationMode        string   `json:"publication_mode"`
	AnonymizationRules     []string `json:"anonymization_rules,omitempty"`
	AggregationThresholds  []string `json:"aggregation_thresholds,omitempty"`
	PublicationReviewGate  []string `json:"publication_review_gate,omitempty"`
	ConfidenceUncertainty  []string `json:"confidence_and_uncertainty,omitempty"`
	DataFreshnessLabeling  []string `json:"data_freshness_labeling,omitempty"`
	ComparabilityLimits    []string `json:"comparability_limits,omitempty"`
	DoNotPublishConditions []string `json:"do_not_publish_conditions,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type publicCaseStudyPack struct {
	PackID               string   `json:"pack_id"`
	DisplayName          string   `json:"display_name"`
	ScenarioType         string   `json:"scenario_type"`
	EvidenceClass        string   `json:"evidence_class"`
	ArchitectureRefs     []string `json:"architecture_refs,omitempty"`
	ScenarioSummary      []string `json:"scenario_summary,omitempty"`
	BeforeState          []string `json:"before_state,omitempty"`
	AfterState           []string `json:"after_state,omitempty"`
	MeasuredOutputs      []string `json:"measured_outputs,omitempty"`
	ReproducibilityNotes []string `json:"reproducibility_notes,omitempty"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type publicCaseStudyPacksResponse struct {
	SchemaVersion  string                `json:"schema_version"`
	MethodologyRef string                `json:"methodology_ref"`
	Packs          []publicCaseStudyPack `json:"packs,omitempty"`
	Limitations    []string              `json:"limitations,omitempty"`
}

func (s server) publicBenchmarkMethodologyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicBenchmarkMethodology())
}

func (s server) publicBenchmarkSetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicBenchmarkSet())
}

func (s server) publicAnalyticsPublicationDisciplineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicAnalyticsPublicationDiscipline())
}

func (s server) publicCaseStudyPacksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPublicCaseStudyPacks())
}

func buildPublicBenchmarkMethodology() publicBenchmarkMethodologyResponse {
	return publicBenchmarkMethodologyResponse{
		SchemaVersion: publicBenchmarkMethodologySchema,
		MethodologyID: "public_benchmark_methodology_v1",
		InputSizeDiscipline: []string{
			"Every published benchmark must declare input shape, event volume, artifact count, or workload duration before any result is interpreted.",
			"Input-size changes require a new benchmark note or a separate published series rather than silent comparison to previous runs.",
		},
		EnvironmentAssumptions: []string{
			"Benchmark publication must state hardware class, execution substrate, storage assumptions, and whether the run is connected, hybrid, or offline.",
			"Runtime benchmarks must state whether numbers are measured on kernel-adjacent signal paths or projected from sizing guidance only.",
		},
		WorkloadProfiles: []string{
			"deploy-gate admission path",
			"audit ingest and evidence retention path",
			"handoff seal and offline verification path",
			"federation proof verification path",
			"validation execution path",
			"runtime detection and bounded response path",
		},
		SubstrateContext: []string{
			"Cluster, VM, or air-gapped substrate differences must be labeled explicitly.",
			"Sidecarless, ambient, or confidential readiness claims remain separate from measured benchmark claims unless directly measured.",
		},
		RepeatabilityRules: []string{
			"Each public benchmark must document warm-up, iteration count, and whether median, p95, or throughput averages are reported.",
			"Variability across runs must be disclosed rather than replaced by a single best-case number.",
		},
		VariabilityDisclosure: []string{
			"Benchmark publication must disclose noise sources, sample count, and meaningful variance bands.",
			"When measurement is still a starting point rather than a published benchmark, the status must say so explicitly.",
		},
		FailureCaseHandling: []string{
			"Failed or partial runs must be recorded as failed, excluded with reason, or published separately as failure-case data.",
			"Degraded-mode behavior must be described distinctly from normal connected-path measurements.",
		},
		NotMeasured: []string{
			"No public benchmark in this methodology claims universal security gain percentages, universal runtime prevention rates, or blanket latency guarantees across all substrates.",
			"Trust quality, governance maturity, and partner admissibility remain evidence-backed judgments rather than benchmark-only outputs.",
		},
		Limitations: []string{
			"This methodology defines publication discipline; it does not by itself prove any specific performance or security outcome until measured results are attached to a benchmark pack.",
		},
	}
}

func buildPublicBenchmarkSet() publicBenchmarkSetResponse {
	return publicBenchmarkSetResponse{
		SchemaVersion:  publicBenchmarkSetSchema,
		MethodologyRef: "/v1/public/benchmarks/methodology",
		Benchmarks: []publicBenchmarkDefinition{
			{
				BenchmarkID:            "deploy_gate_latency",
				DisplayName:            "Deploy-gate latency",
				MetricClass:            "performance",
				MeasurementScope:       []string{"time from evaluated change request to bounded allow/deny response", "includes policy and evidence lookup path only"},
				InputShape:             []string{"declare manifest size", "declare policy count", "declare validation dependency presence"},
				EnvironmentAssumptions: []string{"connected control plane", "stable policy store"},
				PublicationStatus:      "methodology_defined_pending_public_measurement",
				EvidenceRefs:           []string{"/v1/public/benchmark-methodology", "/v1/runtime/boundaries"},
				NotClaimed:             []string{"No blanket per-cluster latency guarantee is claimed until measurement packs are published."},
			},
			{
				BenchmarkID:            "audit_ingest_throughput",
				DisplayName:            "Audit ingest throughput",
				MetricClass:            "performance",
				MeasurementScope:       []string{"event ingest throughput into canonical audit stream", "bounded by stored evidence and persistence path"},
				InputShape:             []string{"declare event size", "declare ingest burst shape", "declare retention mode"},
				EnvironmentAssumptions: []string{"canonical event store available", "no hidden event dropping"},
				PublicationStatus:      "methodology_defined_pending_public_measurement",
				EvidenceRefs:           []string{"/v1/audit/exports", "/v1/public/benchmarks/methodology"},
				NotClaimed:             []string{"No claim of infinite scaling or substrate-independent throughput."},
			},
			{
				BenchmarkID:            "handoff_seal_and_verify",
				DisplayName:            "Handoff seal and verify",
				MetricClass:            "performance_and_security",
				MeasurementScope:       []string{"sealed bundle assembly", "offline verification replay time", "quality gate interpretation path"},
				InputShape:             []string{"declare artifact count", "declare signature and transparency record count"},
				EnvironmentAssumptions: []string{"sealed bundle contains complete verifier inputs"},
				PublicationStatus:      "sample_and_reference_pack_ready",
				EvidenceRefs:           []string{"/v1/public/specs/handoff", "/v1/public/verifier/reference-pack", "/v1/public/samples/handoff"},
				NotClaimed:             []string{"No perpetual signer validity or permanent timestamp service availability is claimed."},
			},
			{
				BenchmarkID:            "federation_proof_verification",
				DisplayName:            "Federation proof verification",
				MetricClass:            "security_and_interoperability",
				MeasurementScope:       []string{"peer proof freshness check", "disclosure-profile interpretation", "local accept/reject narrative"},
				InputShape:             []string{"declare proof freshness", "declare peer state", "declare disclosure profile"},
				EnvironmentAssumptions: []string{"local policy remains authoritative"},
				PublicationStatus:      "sample_and_spec_ready",
				EvidenceRefs:           []string{"/v1/public/specs/federation-proof-exchange", "/v1/public/samples/federation-proof-exchange"},
				NotClaimed:             []string{"No shared global authority or remote bypass of local admissibility is claimed."},
			},
			{
				BenchmarkID:            "validation_execution",
				DisplayName:            "Validation execution",
				MetricClass:            "performance_and_correctness",
				MeasurementScope:       []string{"time and bounded resource envelope for validation execution", "certificate emission readiness"},
				InputShape:             []string{"declare scenario count", "declare execution profile", "declare quota model"},
				EnvironmentAssumptions: []string{"validation harness quotas remain explicit"},
				PublicationStatus:      "methodology_defined_pending_public_measurement",
				EvidenceRefs:           []string{"/v1/public/specs/validation-certificate", "/v1/validation/executions"},
				NotClaimed:             []string{"No absolute guarantee of scenario completeness or universal production equivalence is claimed."},
			},
			{
				BenchmarkID:            "runtime_overhead",
				DisplayName:            "Runtime overhead",
				MetricClass:            "performance",
				MeasurementScope:       []string{"bounded control-plane and runtime-agent overhead in the declared substrate", "overhead starting points versus measured publication status"},
				InputShape:             []string{"declare runtime profile", "declare enforcement mode", "declare workload criticality"},
				EnvironmentAssumptions: []string{"runtime boundary semantics are documented before public measurement"},
				PublicationStatus:      "starting_points_only_not_public_claim",
				EvidenceRefs:           []string{"/v1/runtime/boundaries", "/v1/execution/ambient-readiness"},
				NotClaimed:             []string{"No universal low-latency or under-one-percent overhead claim is made from this surface alone."},
			},
			{
				BenchmarkID:            "runtime_response_latency",
				DisplayName:            "Runtime response latency",
				MetricClass:            "security_and_performance",
				MeasurementScope:       []string{"time from bounded signal detection to selected response decision", "forensic-first and approval gate semantics remain visible"},
				InputShape:             []string{"declare signal type", "declare confidence threshold", "declare response mode"},
				EnvironmentAssumptions: []string{"response policy and least-invasive-first ordering are explicit"},
				PublicationStatus:      "methodology_defined_pending_public_measurement",
				EvidenceRefs:           []string{"/v1/runtime/response-policy", "/v1/runtime/rule-packs"},
				NotClaimed:             []string{"No promise that every attack is blocked before execution or before CPU dispatch."},
			},
			{
				BenchmarkID:            "degraded_mode_behavior",
				DisplayName:            "Degraded-mode behavior",
				MetricClass:            "resilience",
				MeasurementScope:       []string{"offline, delayed-sync, stale-peer, or connector-degraded publication behavior", "visibility of limitations during degraded operation"},
				InputShape:             []string{"declare degraded scenario", "declare retained evidence class", "declare delayed-sync semantics"},
				EnvironmentAssumptions: []string{"degraded mode remains explicit and policy-bounded"},
				PublicationStatus:      "spec_and_contract_ready",
				EvidenceRefs:           []string{"/v1/execution/coverage/matrix", "/v1/integrations/safety", "/v1/b2b/consortium-readiness"},
				NotClaimed:             []string{"No hidden convergence or silent fallback to healthy-state semantics is claimed."},
			},
		},
		Limitations: []string{
			"The public benchmark set defines publication targets and status classes. Only entries with measured publication packs should be used for comparative market claims.",
		},
	}
}

func buildPublicAnalyticsPublicationDiscipline() publicAnalyticsPublicationDisciplineResponse {
	return publicAnalyticsPublicationDisciplineResponse{
		SchemaVersion:   publicAnalyticsPublicationDisciplineSchema,
		PublicationMode: "aggregated_and_anonymized_only",
		AnonymizationRules: []string{
			"Published analytics must remove customer identifiers, partner-specific secrets, and topology details that would reveal internal estates.",
			"Publication must aggregate at cohort or sector level rather than exposing single-customer posture unless the customer explicitly owns that publication.",
		},
		AggregationThresholds: []string{
			"Sector or cohort views must not publish if the cohort is too small to prevent easy re-identification.",
			"Rare incident classes or unique deployment profiles must be withheld or generalized when aggregation would disclose a single organization.",
		},
		PublicationReviewGate: []string{
			"Every public analytics release requires methodology review, disclosure review, and confirmation that confidence and freshness labels remain attached.",
			"Comparative publication requires an explicit note on what is not measured or not comparable across sectors.",
		},
		ConfidenceUncertainty: []string{
			"Every published trend, score, or cohort summary must include uncertainty and limitation labels.",
			"Black-box trust indices without evidence drill-down or uncertainty signaling are out of scope.",
		},
		DataFreshnessLabeling: []string{
			"Publication must state collection window and freshness date.",
			"Delayed-sync, stale, or partially missing data must remain labeled rather than smoothed away.",
		},
		ComparabilityLimits: []string{
			"Cross-sector comparisons must disclose differences in deployment topology, runtime profile, governance expectations, and offline constraints.",
			"Partner and internal trust posture views are not automatically comparable to public benchmark sets.",
		},
		DoNotPublishConditions: []string{
			"Do not publish analytics that would leak customer-specific incident narratives, low-sample outliers, or identifiable partner distrust states.",
			"Do not publish trust scores as standalone market claims without the accompanying methodology and limitation narrative.",
		},
		Limitations: []string{
			"This discipline governs public aggregate analytics publication. It does not authorize publication of customer-specific or partner-specific evidence without separate approval.",
		},
	}
}

func buildPublicCaseStudyPacks() publicCaseStudyPacksResponse {
	return publicCaseStudyPacksResponse{
		SchemaVersion:  publicCaseStudyPacksSchema,
		MethodologyRef: "/v1/public/benchmarks/methodology",
		Packs: []publicCaseStudyPack{
			{
				PackID:               "offline_handoff_verification",
				DisplayName:          "Offline sealed handoff verification",
				ScenarioType:         "handoff_and_offline_verification",
				EvidenceClass:        "replayable_reference_pack",
				ArchitectureRefs:     []string{"air-gapped-regulated"},
				ScenarioSummary:      []string{"A disconnected operator receives a sealed handoff bundle and replays verification locally.", "The narrative focuses on verifier inputs, failure states, and archive integrity interpretation."},
				BeforeState:          []string{"No local interpretation of the incoming bundle exists.", "Connectivity to the originating environment is unavailable or intentionally blocked."},
				AfterState:           []string{"Local verifier reaches an explainable verification result.", "Archive integrity, signature state, and quality gate outcomes are preserved for audit use."},
				MeasuredOutputs:      []string{"bundle_replayable=true", "offline_verification_path=available", "quality_gate_narrative=present"},
				ReproducibilityNotes: []string{"Use /v1/public/verifier/reference-pack as the canonical replay input.", "Treat failure states as verifier outcomes rather than as hidden internal errors."},
				EvidenceRefs:         []string{"/v1/public/verifier/reference-pack", "/v1/public/samples/handoff", "/v1/public/specs/handoff"},
				Limitations:          []string{"This pack demonstrates verifier flow discipline, not live remote signer health or current upstream service state."},
			},
			{
				PackID:               "supplier_proof_acceptance",
				DisplayName:          "Supplier proof acceptance with local policy gate",
				ScenarioType:         "b2b_proof_exchange",
				EvidenceClass:        "public_sample_and_policy_narrative",
				ArchitectureRefs:     []string{"supplier-federation"},
				ScenarioSummary:      []string{"A supplier submits a sealed proof and the receiving organization performs local admissibility and freshness checks.", "Accepted and rejected paths are both published to keep rejection semantics visible."},
				BeforeState:          []string{"Partner onboarding exists but admissibility still depends on local verification.", "Disclosure profile and local policy override semantics are explicit."},
				AfterState:           []string{"Accepted and rejected proof outcomes remain explainable.", "Local policy continues to own the final decision."},
				MeasuredOutputs:      []string{"accepted_path=sampled", "rejected_path=sampled", "local_policy_override=explicit"},
				ReproducibilityNotes: []string{"Replay both accepted and rejected proof verification samples.", "Treat freshness, stale peer, and divergence fields as first-class verifier inputs."},
				EvidenceRefs:         []string{"/v1/public/samples/proof-verification", "/v1/public/samples/federation-proof-exchange", "/v1/public/specs/proof-verification"},
				Limitations:          []string{"This pack does not imply automatic partner trust or eliminate bilateral governance requirements."},
			},
			{
				PackID:               "runtime_hardening_tuning",
				DisplayName:          "Runtime hardening tuning before public claims",
				ScenarioType:         "runtime_and_publication_discipline",
				EvidenceClass:        "contract_linked_summary",
				ArchitectureRefs:     []string{"runtime-hardened-enterprise-cluster", "regulated-saas"},
				ScenarioSummary:      []string{"A deployment moves from conservative runtime posture to stronger bounded hardening only after response policy and runtime boundaries are explicit.", "The pack shows how public narrative stays tied to measured or documented boundaries."},
				BeforeState:          []string{"Runtime findings exist but public performance/security claims are not publication-ready.", "Operator still tunes confidence thresholds and least-invasive-first ordering."},
				AfterState:           []string{"Runtime response policy is explicit.", "Public benchmark methodology and boundary semantics can frame later publication without over-claiming."},
				MeasuredOutputs:      []string{"response_policy=explicit", "boundary_doc=present", "publication_status=methodology_only_until_measured"},
				ReproducibilityNotes: []string{"Use runtime rule packs, response policy, and runtime boundaries as the source contract set.", "Do not convert starting-point overhead data into market claims without measured benchmark packs."},
				EvidenceRefs:         []string{"/v1/runtime/response-policy", "/v1/runtime/boundaries", "/v1/public/benchmarks/set"},
				Limitations:          []string{"This pack is a bounded publication-discipline example, not a measured public performance proof."},
			},
		},
		Limitations: []string{
			"Case-study packs are bounded evidence narratives or replayable examples. Synthetic or contract-linked examples must remain labeled as such and must not be presented as unnamed customer proof.",
		},
	}
}
