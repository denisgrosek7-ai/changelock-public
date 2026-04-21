package main

import (
	"net/http"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	phase1ExecutionFoundationSchema = "1.execution_foundation_summary.v1"
	phase1ExecutionContractsSchema  = "1.execution_foundation_contracts.v1"
	phase1ExecutionBenchmarksSchema = "1.execution_foundation_benchmarks.v1"
	phase1ExecutionAsyncSchema      = "1.execution_foundation_async.v1"
	phase1ExecutionTrustSchema      = "1.execution_foundation_trust.v1"
)

type phase1FoundationGate struct {
	GateID       string   `json:"gate_id"`
	Status       string   `json:"status"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
}

type phase1ExecutionFoundationResponse struct {
	SchemaVersion      string                 `json:"schema_version"`
	PhaseID            string                 `json:"phase_id"`
	CurrentState       string                 `json:"current_state"`
	CanonicalTruthLine []string               `json:"canonical_truth_line,omitempty"`
	Gates              []phase1FoundationGate `json:"gates,omitempty"`
	Limitations        []string               `json:"limitations,omitempty"`
}

type phase1DegradedMode struct {
	Component           string `json:"component"`
	FailureClass        string `json:"failure_class"`
	ExpectedBehavior    string `json:"expected_behavior"`
	OperatorExpectation string `json:"operator_expectation"`
}

type phase1ExecutionContractsResponse struct {
	SchemaVersion         string               `json:"schema_version"`
	CanonicalEventSchema  string               `json:"canonical_event_schema"`
	EnvelopeFields        []string             `json:"envelope_fields,omitempty"`
	CorrelationDiscipline []string             `json:"correlation_discipline,omitempty"`
	SchemaDiscipline      []string             `json:"schema_discipline,omitempty"`
	DegradedModes         []phase1DegradedMode `json:"degraded_modes,omitempty"`
	ExceptionGovernance   []string             `json:"exception_governance,omitempty"`
	ObservabilitySecurity []string             `json:"observability_security,omitempty"`
	Limitations           []string             `json:"limitations,omitempty"`
}

type phase1BenchmarkProfile struct {
	ProfileID       string   `json:"profile_id"`
	DisplayName     string   `json:"display_name"`
	Characteristics []string `json:"characteristics,omitempty"`
}

type phase1MeasuredPath struct {
	PathID        string   `json:"path_id"`
	PathType      string   `json:"path_type"`
	CurrentStatus string   `json:"current_status"`
	EvidenceRefs  []string `json:"evidence_refs,omitempty"`
}

type phase1ExecutionBenchmarksResponse struct {
	SchemaVersion        string                   `json:"schema_version"`
	CurrentState         string                   `json:"current_state"`
	MetricTaxonomy       []string                 `json:"metric_taxonomy,omitempty"`
	Profiles             []phase1BenchmarkProfile `json:"profiles,omitempty"`
	CriticalPaths        []phase1MeasuredPath     `json:"critical_paths,omitempty"`
	RegressionDiscipline []string                 `json:"regression_discipline,omitempty"`
	Limitations          []string                 `json:"limitations,omitempty"`
}

type phase1AsyncEventEnvelope struct {
	SchemaVersion  string   `json:"schema_version"`
	RequiredFields []string `json:"required_fields,omitempty"`
}

type phase1ExecutionAsyncResponse struct {
	SchemaVersion      string                   `json:"schema_version"`
	CurrentState       string                   `json:"current_state"`
	SynchronousPath    []string                 `json:"synchronous_path,omitempty"`
	MigratedAsyncPaths []string                 `json:"migrated_async_paths,omitempty"`
	TargetAsyncPaths   []string                 `json:"target_async_paths,omitempty"`
	EventEnvelope      phase1AsyncEventEnvelope `json:"event_envelope"`
	WorkerDiscipline   []string                 `json:"worker_discipline,omitempty"`
	FailureSemantics   []string                 `json:"failure_semantics,omitempty"`
	BackpressureRules  []string                 `json:"backpressure_rules,omitempty"`
	ReplayGovernance   []string                 `json:"replay_governance,omitempty"`
	ConnectorIsolation []string                 `json:"connector_isolation,omitempty"`
	Limitations        []string                 `json:"limitations,omitempty"`
}

type phase1ProviderCapability struct {
	ProviderMode    string   `json:"provider_mode"`
	CapabilityClass string   `json:"capability_class"`
	CurrentSupport  string   `json:"current_support"`
	Notes           []string `json:"notes,omitempty"`
}

type phase1ExecutionTrustResponse struct {
	SchemaVersion            string                     `json:"schema_version"`
	CurrentState             string                     `json:"current_state"`
	Provider                 signing.ProviderDescriptor `json:"provider"`
	LifecycleStates          []string                   `json:"lifecycle_states,omitempty"`
	KeyClasses               []string                   `json:"key_classes,omitempty"`
	RotationModel            []string                   `json:"rotation_model,omitempty"`
	VerificationContinuity   []string                   `json:"verification_continuity,omitempty"`
	EnvironmentBoundaries    []string                   `json:"environment_boundaries,omitempty"`
	ProviderCapabilityMatrix []phase1ProviderCapability `json:"provider_capability_matrix,omitempty"`
	Limitations              []string                   `json:"limitations,omitempty"`
}

func (s server) executionFoundationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, s.buildExecutionFoundationResponse())
}

func (s server) executionFoundationContractsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildExecutionFoundationContractsResponse())
}

func (s server) executionFoundationBenchmarksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildExecutionFoundationBenchmarksResponse())
}

func (s server) executionFoundationAsyncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildExecutionFoundationAsyncResponse())
}

func (s server) executionFoundationTrustHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, s.buildExecutionFoundationTrustResponse())
}

func (s server) buildExecutionFoundationResponse() phase1ExecutionFoundationResponse {
	trustStatus := "provider_abstraction_ready_signer_not_activated"
	if provider := s.phase1TrustProvider(); provider.ProviderMode == signing.ModeVaultTransit {
		trustStatus = "provider_abstraction_ready_external_managed_active"
	} else if provider.ProviderMode == signing.ModeSoftware {
		trustStatus = "provider_abstraction_ready_local_signer_active"
	}

	return phase1ExecutionFoundationResponse{
		SchemaVersion: phase1ExecutionFoundationSchema,
		PhaseID:       "phase_1_execution_foundation",
		CurrentState:  "phase1_bounded_closure_ready",
		CanonicalTruthLine: []string{
			"Canonical evidence remains authoritative across verify, admit, run, respond, preserve, and exchange paths.",
			"Execution foundation contracts are additive and must not introduce a new shadow truth layer for recommendations, analytics, or AI surfaces.",
		},
		Gates: []phase1FoundationGate{
			{
				GateID:       "g0_contracts",
				Status:       "implemented_foundation_slice",
				Summary:      "Canonical execution envelope fields, correlation discipline, schema versioning rules, and degraded-mode semantics are now explicit in code and API contracts.",
				EvidenceRefs: []string{"/v1/foundation/execution/contracts"},
			},
			{
				GateID:       "g1_measurement",
				Status:       "baseline_measured_regression_gated_and_proof_visible",
				Summary:      "Prometheus metrics, benchmark baselines, explicit harness catalog, regression evaluation logic, and benchmark proof readback now exist in code.",
				EvidenceRefs: []string{"/metrics", "/v1/foundation/execution/benchmarks", "/v1/foundation/execution/benchmarks/harness", "/v1/foundation/execution/benchmarks/evaluate", "/v1/foundation/execution/proofs", "/v1/public/benchmarks/methodology", "/v1/public/benchmarks/set"},
			},
			{
				GateID:       "g2_async_foundation",
				Status:       "audit_backed_async_migration_active_for_sync_forward",
				Summary:      "Critical-path split, canonical event envelope, replay semantics, and connector isolation rules are explicit; the visible sync forward path now runs through the bounded async task path instead of the ingest request path.",
				EvidenceRefs: []string{"/v1/foundation/execution/async", "/v1/foundation/execution/async/tasks", "/v1/foundation/execution/traces", "/v1/foundation/execution/proofs"},
			},
			{
				GateID:       "g3_trust_material",
				Status:       trustStatus,
				Summary:      "Signer provider abstraction exists today with software and vault-transit implementations; lifecycle, rotation, historical verification continuity, and provider-backed rotation drill evidence now exist.",
				EvidenceRefs: []string{"/v1/foundation/execution/trust", "/v1/foundation/execution/trust/rotation-drill", "/v1/foundation/execution/trust/rotation-drills"},
			},
			{
				GateID:       "g4_operational_proof",
				Status:       "proof_pack_surface_active",
				Summary:      "Tracing, async migration, benchmark rerun evidence, and rotation drill evidence are readable through one bounded operational proof surface.",
				EvidenceRefs: []string{"/v1/foundation/execution/proofs", "/v1/foundation/execution/traces"},
			},
		},
		Limitations: []string{
			"Phase 1 is closed in a bounded foundation sense: canonical contracts, benchmark gates, durable async migration for the visible sync-forward path, and signer rotation drill evidence exist in code.",
			"This does not claim a universal queue substrate, universal provider matrix, or a full third-party tracing platform.",
		},
	}
}

func buildExecutionFoundationContractsResponse() phase1ExecutionContractsResponse {
	return phase1ExecutionContractsResponse{
		SchemaVersion:        phase1ExecutionContractsSchema,
		CanonicalEventSchema: audit.ExecutionEventSchemaVersion,
		EnvelopeFields:       audit.ExecutionEnvelopeFieldSet(),
		CorrelationDiscipline: []string{
			"Every canonical event must preserve request-linked correlation through trace_id, correlation_id, decision_id, event_id, and idempotency_key.",
			"Decision, incident, and recommendation overlays may enrich the graph but must not replace the canonical execution identifiers already attached to the event envelope.",
		},
		SchemaDiscipline: []string{
			"Schema changes on canonical execution envelopes require explicit versioning rather than silent field drift.",
			"Backward-compatible additions are preferred; removals or semantic changes require migration and deprecation guidance.",
			"Recommendation, analytics, and AI layers remain downstream readers of canonical evidence rather than replacement truth stores.",
		},
		DegradedModes: []phase1DegradedMode{
			{
				Component:           "audit_store",
				FailureClass:        "persistence_unavailable",
				ExpectedBehavior:    "Readiness must degrade explicitly; canonical ingest may fail closed rather than silently dropping evidence.",
				OperatorExpectation: "Investigate store health before trusting new control-plane state transitions as durably anchored.",
			},
			{
				Component:           "signer_provider",
				FailureClass:        "trust_provider_unavailable",
				ExpectedBehavior:    "Signing and verify-on-read semantics remain explicit; no implicit fallback from externally managed trust to weaker local trust is allowed.",
				OperatorExpectation: "Treat provider outages as visible trust degradation requiring review or bounded fail-safe posture.",
			},
			{
				Component:           "connector",
				FailureClass:        "external_delivery_failure",
				ExpectedBehavior:    "Connector failures stay advisory and isolated from canonical truth writes; retries and replay remain visible.",
				OperatorExpectation: "Observe connector health separately from control-plane allow/deny truth.",
			},
			{
				Component:           "observability",
				FailureClass:        "metrics_or_tracing_unavailable",
				ExpectedBehavior:    "Execution paths remain bounded and deterministic; observability loss does not authorize hidden behavior changes.",
				OperatorExpectation: "Treat missing telemetry as a reduced-visibility state, not as evidence of system health.",
			},
		},
		ExceptionGovernance: []string{
			"Benchmark overrides, replay overrides, signer changes, and trust-provider exceptions must leave an audit trail rather than silently mutating system posture.",
			"Emergency behavior must remain visible as exception governance, not baked into the normal success path.",
		},
		ObservabilitySecurity: []string{
			"Tracing and telemetry must redact secrets and sensitive payload fragments before publication.",
			"Sampling and retention settings must stay bounded so observability does not become a new leakage path or storage denial vector.",
		},
		Limitations: []string{
			"The contract surface is implemented, but not every component has been migrated yet to use the expanded execution envelope as its primary external-facing response shape.",
		},
	}
}

func buildExecutionFoundationBenchmarksResponse() phase1ExecutionBenchmarksResponse {
	return phase1ExecutionBenchmarksResponse{
		SchemaVersion: phase1ExecutionBenchmarksSchema,
		CurrentState:  "baseline_measured_and_publication_disciplined",
		MetricTaxonomy: []string{
			"user_facing_latency",
			"control_plane_latency",
			"evidence_latency",
			"background_completion_latency",
			"resource_cost_envelope",
			"failure_and_recovery_latency",
		},
		Profiles: []phase1BenchmarkProfile{
			{
				ProfileID:   "local_baseline",
				DisplayName: "Local Baseline",
				Characteristics: []string{
					"fast feedback for development and regression detection",
					"low-to-medium concurrency with deterministic fixture inputs",
				},
			},
			{
				ProfileID:   "production_like",
				DisplayName: "Production-like",
				Characteristics: []string{
					"representative control-plane concurrency and evidence write patterns",
					"used to compare admission, runtime, audit, sealing, and validation envelopes under realistic pressure",
				},
			},
			{
				ProfileID:   "stress",
				DisplayName: "Stress",
				Characteristics: []string{
					"burst deployment, audit spike, reconnect, retry, and high-fanout dispatch scenarios",
					"used to measure queue lag, replay cost, and recovery posture rather than average-case UX only",
				},
			},
		},
		CriticalPaths: []phase1MeasuredPath{
			{
				PathID:        "deploy_gate_admission",
				PathType:      "synchronous_control_path",
				CurrentStatus: "baseline_measured",
				EvidenceRefs:  []string{"/v1/public/benchmarks/set", "/metrics"},
			},
			{
				PathID:        "runtime_signal_and_dispatch",
				PathType:      "runtime_control_and_background_split",
				CurrentStatus: "baseline_and_boundaries_present",
				EvidenceRefs:  []string{"/v1/runtime/boundaries", "/v1/runtime/response-policy"},
			},
			{
				PathID:        "handoff_and_validation",
				PathType:      "background_heavy_path",
				CurrentStatus: "baseline_measured_high_cost_path",
				EvidenceRefs:  []string{"/v1/public/benchmarks/set", "/v1/public/specs/handoff", "/v1/public/specs/validation-certificate"},
			},
			{
				PathID:        "forensics_reconstruction",
				PathType:      "read_heavy_replay_path",
				CurrentStatus: "baseline_measured",
				EvidenceRefs:  []string{"/v1/forensics/state", "/v1/forensics/replay"},
			},
		},
		RegressionDiscipline: []string{
			"Regression review must compare build-to-build baselines, not single best-case numbers.",
			"p50, p95, and p99 remain required for critical decision paths; averages alone are not sufficient.",
			"Benchmark overrides must remain exception-governed rather than silently bypassing the gate.",
			"Benchmark harness catalog and evaluation logic are now explicit in the foundation surface so CI gates can stay profile-aware rather than ad hoc.",
		},
		Limitations: []string{
			"Current baseline data is stronger than methodology-only, and regression evaluation logic now exists in code, but full production-profile automation and substrate-specific benchmark publication packs are still future work.",
		},
	}
}

func buildExecutionFoundationAsyncResponse() phase1ExecutionAsyncResponse {
	return phase1ExecutionAsyncResponse{
		SchemaVersion: phase1ExecutionAsyncSchema,
		CurrentState:  "critical_path_split_with_sync_forward_migrated",
		SynchronousPath: []string{
			"admission decision and minimal evidence anchoring",
			"minimal verification required for allow/deny truth",
			"canonical event normalization and stable identity assignment",
		},
		MigratedAsyncPaths: []string{
			"sync runtime forward and remote delivery",
			"audit-backed bounded async tasks for replayable background execution",
		},
		TargetAsyncPaths: []string{
			"connector dispatch and remote delivery",
			"heavy report generation and export assembly",
			"partner handoff and proof exchange delivery",
			"expensive enrichment and recompute-style workflows",
		},
		EventEnvelope: phase1AsyncEventEnvelope{
			SchemaVersion:  audit.ExecutionEventSchemaVersion,
			RequiredFields: audit.ExecutionEnvelopeFieldSet(),
		},
		WorkerDiscipline: []string{
			"background workers must use bounded concurrency and timeout discipline",
			"graceful shutdown must preserve replay-safe task semantics",
			"async tasks must be observable by task type and failure class",
		},
		FailureSemantics: []string{
			"retryable",
			"transient_external",
			"permanent_business_rule_failure",
			"schema_failure",
			"poison_payload",
			"trust_provider_failure",
		},
		BackpressureRules: []string{
			"queue depth must remain visible before workers saturate",
			"connector isolation and rate limits must prevent external systems from dragging the canonical control path into failure",
			"degraded mode must remain explicit when backlog exceeds the bounded envelope",
		},
		ReplayGovernance: []string{
			"replay must preserve original versus replayed intent through distinct IDs and audit trail",
			"replay may recompute background work, but it must not create hidden duplicate mutations against canonical truth",
		},
		ConnectorIsolation: []string{
			"external connectors remain outside the critical decision path",
			"delivery failures stay visible without rewriting the canonical allow/deny record",
		},
		Limitations: []string{
			"Phase 1 now persists bounded async tasks durably through canonical audit events and uses them for the visible sync-forward migration path. The target async paths above remain broader migration targets, not a claim that every one of them is already moved.",
		},
	}
}

func (s server) buildExecutionFoundationTrustResponse() phase1ExecutionTrustResponse {
	provider := s.phase1TrustProvider()
	currentState := "provider_abstraction_ready_signer_not_activated"
	switch provider.ProviderMode {
	case signing.ModeSoftware:
		currentState = "provider_abstraction_ready_local_signer_active_rotation_drill_ready"
	case signing.ModeVaultTransit:
		currentState = "provider_abstraction_ready_external_managed_active_rotation_drill_ready"
	}

	return phase1ExecutionTrustResponse{
		SchemaVersion: phase1ExecutionTrustSchema,
		CurrentState:  currentState,
		Provider:      provider,
		LifecycleStates: []string{
			signing.KeyStateProvisioned,
			signing.KeyStateActive,
			signing.KeyStateRotatePending,
			signing.KeyStateRetiredVerifyOnly,
			signing.KeyStateRevoked,
			signing.KeyStateDestroyed,
		},
		KeyClasses: provider.KeyClasses,
		RotationModel: []string{
			"rotation must support dual-trust windows rather than an instantaneous cutover",
			"retired keys may stay verify-only for historical evidence continuity",
			"revocation must remain visible in audit and trust-provider semantics rather than silently deleting history",
		},
		VerificationContinuity: []string{
			"historical verification must remain possible across rotated signer sets",
			"verification context must keep signer identity, provider identity, and lifecycle state visible",
			"trust-set verification can preserve verify-only retired members while still rejecting revoked members",
		},
		EnvironmentBoundaries: []string{
			"development trust remains lower-assurance and explicitly separated from production trust material",
			"production trust should use externally managed custody rather than application-local secrets",
		},
		ProviderCapabilityMatrix: []phase1ProviderCapability{
			{
				ProviderMode:    signing.ModeSoftware,
				CapabilityClass: "development_or_lower_trust",
				CurrentSupport:  "implemented",
				Notes: []string{
					"Supports sign/verify and rotation semantics in code.",
					"Not a substitute for externally managed production trust custody.",
				},
			},
			{
				ProviderMode:    signing.ModeVaultTransit,
				CapabilityClass: "production_managed_trust",
				CurrentSupport:  "implemented",
				Notes: []string{
					"Supports externally managed signing and verification through Vault transit.",
					"Wider cloud KMS and HSM appliance coverage is still future work.",
				},
			},
			{
				ProviderMode:    "external_kms_hsm_future",
				CapabilityClass: "production_managed_trust",
				CurrentSupport:  "planned_not_implemented",
				Notes: []string{
					"Provider-specific adapters and deeper capability matrices belong to later foundation hardening work.",
				},
			},
		},
		Limitations: []string{
			"Provider abstraction and lifecycle semantics are now explicit. Broader provider coverage, automated rotation orchestration, and non-Vault enterprise custody integrations remain future work.",
		},
	}
}

func (s server) phase1TrustProvider() signing.ProviderDescriptor {
	if s.signing == nil || s.signing.runtime == nil {
		return (*signing.Runtime)(nil).DescribeProvider()
	}
	return s.signing.runtime.DescribeProvider()
}
