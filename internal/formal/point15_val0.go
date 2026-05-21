package formal

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point15Val0StateActive         = "point15_val0_freshness_discipline_active"
	Point15Val0StateBlocked        = "point15_val0_freshness_discipline_blocked"
	Point15Val0StateReviewRequired = "point15_val0_freshness_discipline_review_required"
	Point15Val0StateIncomplete     = "point15_val0_freshness_discipline_incomplete"
)

const (
	point15Val0PointID             = "point_15"
	point15Val0WaveID              = "val_0"
	point15Val0Scope               = "freshness_discipline_foundation_and_downgrade_taxonomy"
	point15Val0FreshnessDisclaimer = "bounded_freshness_discipline no_continuous_assurance_guarantee point15_val0"
	point15Val0BlockedPassToken    = "point_15_pass"

	point15Val0FreshnessFresh       = "fresh"
	point15Val0FreshnessStale       = "stale"
	point15Val0FreshnessExpired     = "expired"
	point15Val0FreshnessRevoked     = "revoked"
	point15Val0FreshnessSuperseded  = "superseded"
	point15Val0FreshnessDrifted     = "drifted"
	point15Val0FreshnessMissing     = "missing"
	point15Val0FreshnessUnsupported = "unsupported"
	point15Val0FreshnessTampered    = "tampered"

	point15Val0DowngradeRetainActive = "retain_active_only_if_fresh_and_bound"
	point15Val0DowngradeReview       = "downgrade_to_review_required"
	point15Val0DowngradeBlocked      = "downgrade_to_blocked"
	point15Val0DowngradeIncomplete   = "downgrade_to_incomplete"
)

type Point15Val0DependencySnapshot struct {
	Point14ValECurrentState              string                `json:"point14_vale_current_state"`
	Point14ValEDependencyState           string                `json:"point14_vale_dependency_state"`
	Point14ValEClosureEvaluatorState     string                `json:"point14_vale_closure_evaluator_state"`
	Point14ValEPassClosureManifestState  string                `json:"point14_vale_pass_closure_manifest_state"`
	Point14PassAllowed                   bool                  `json:"point14_pass_allowed"`
	Point14PassToken                     string                `json:"point14_pass_token"`
	Point14PassManifestPointID           string                `json:"point14_pass_manifest_point_id"`
	Point14PassManifestWaveID            string                `json:"point14_pass_manifest_wave_id"`
	Point14PassManifestClosureToken      string                `json:"point14_pass_manifest_closure_token"`
	Point14ValEComputedFromUpstream      bool                  `json:"point14_vale_computed_from_upstream"`
	Point14ValEMerged                    bool                  `json:"point14_vale_merged"`
	Point14ValECIGreen                   bool                  `json:"point14_vale_ci_green"`
	Point14ValEReviewedOnMain            bool                  `json:"point14_vale_reviewed_on_main"`
	Point15PassSeen                      bool                  `json:"point15_pass_seen"`
	InheritedPoint13ValECurrentState     string                `json:"inherited_point13_vale_current_state"`
	InheritedPoint13ValEPassClosureState string                `json:"inherited_point13_vale_pass_closure_state"`
	InheritedPoint12CurrentState         string                `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState      string                `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState     string                `json:"inherited_point12_pass_closure_state"`
	InheritedPoint11CurrentState         string                `json:"inherited_point11_current_state"`
	InheritedPoint11PublicationState     string                `json:"inherited_point11_publication_state"`
	InheritedPoint11NoOverclaimState     string                `json:"inherited_point11_no_overclaim_state"`
	InheritedPoint11FinalPassGateState   string                `json:"inherited_point11_final_pass_gate_state"`
	InheritedPoint10CurrentState         string                `json:"inherited_point10_current_state"`
	InheritedPoint10NoOverclaimState     string                `json:"inherited_point10_no_overclaim_state"`
	InheritedPoint10ProjectionState      string                `json:"inherited_point10_projection_state"`
	InheritedPoint10PassRuleState        string                `json:"inherited_point10_pass_rule_state"`
	InheritedTenantScope                 string                `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput           bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                  []string              `json:"review_prerequisites,omitempty"`
	Point14ValE                          Point14ValEFoundation `json:"point14_vale"`
}

type Point15Val0FreshnessDisciplineFoundation struct {
	CurrentState             string                               `json:"current_state"`
	BlockingReasons          []string                             `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites      []string                             `json:"review_prerequisites,omitempty"`
	FreshnessDisclaimer      string                               `json:"freshness_disclaimer"`
	DependencyState          string                               `json:"dependency_state"`
	FreshnessTaxonomyState   string                               `json:"freshness_taxonomy_state"`
	DowngradeTaxonomyState   string                               `json:"downgrade_taxonomy_state"`
	EvidenceContextState     string                               `json:"evidence_context_state"`
	TenantBoundaryState      string                               `json:"tenant_boundary_state"`
	TimestampDisciplineState string                               `json:"timestamp_discipline_state"`
	AuthorityBoundaryState   string                               `json:"authority_boundary_state"`
	NoOverclaimState         string                               `json:"no_overclaim_state"`
	Dependency               Point15Val0DependencySnapshot        `json:"dependency"`
	FreshnessTaxonomy        Point15Val0EvidenceFreshnessTaxonomy `json:"freshness_taxonomy"`
	DowngradeTaxonomy        Point15Val0DowngradeTaxonomy         `json:"downgrade_taxonomy"`
	EvidenceContext          Point15Val0FreshnessEvidenceContext  `json:"evidence_context"`
	TimestampDiscipline      Point15Val0TimestampDiscipline       `json:"timestamp_discipline"`
	AuthorityBoundary        Point15Val0AuthorityBoundary         `json:"authority_boundary"`
	NoOverclaimGuard         Point15Val0NoOverclaimGuard          `json:"no_overclaim_guard"`
}

type Point15Val0EvidenceFreshnessTaxonomy struct {
	TaxonomyID               string   `json:"taxonomy_id"`
	FreshnessStatus          string   `json:"freshness_status"`
	MappedState              string   `json:"mapped_state"`
	MappedDowngradeOutcome   string   `json:"mapped_downgrade_outcome"`
	AllowedFreshnessStates   []string `json:"allowed_freshness_states,omitempty"`
	AllowedDowngradeOutcomes []string `json:"allowed_downgrade_outcomes,omitempty"`
}

type Point15Val0DowngradeTaxonomy struct {
	TaxonomyID                    string `json:"taxonomy_id"`
	FreshnessStatus               string `json:"freshness_status"`
	DowngradeOutcome              string `json:"downgrade_outcome"`
	SupersessionLineageRef        string `json:"supersession_lineage_ref"`
	DriftIsDecisive               bool   `json:"drift_is_decisive"`
	MissingFreshnessProofDecisive bool   `json:"missing_freshness_proof_decisive"`
	FreshnessProofPresent         bool   `json:"freshness_proof_present"`
	RetainsPass                   bool   `json:"retains_pass"`
	RetainsActiveClosure          bool   `json:"retains_active_closure"`
}

type Point15Val0FreshnessEvidenceContext struct {
	ContextID                      string `json:"context_id"`
	EvidenceID                     string `json:"evidence_id"`
	TenantScope                    string `json:"tenant_scope"`
	ReferencedTenantScope          string `json:"referenced_tenant_scope"`
	SourceID                       string `json:"source_id"`
	EvidenceHash                   string `json:"evidence_hash"`
	PolicyVersion                  string `json:"policy_version"`
	EngineVersion                  string `json:"engine_version"`
	SchemaVersion                  string `json:"schema_version"`
	ObservedAt                     string `json:"observed_at"`
	ObservedTimeSource             string `json:"observed_time_source"`
	EvaluatedAt                    string `json:"evaluated_at"`
	EvaluatedTimeSource            string `json:"evaluated_time_source"`
	ValidatedAt                    string `json:"validated_at"`
	ValidatedTimeSource            string `json:"validated_time_source"`
	FreshnessStatus                string `json:"freshness_status"`
	DowngradeOutcome               string `json:"downgrade_outcome"`
	FreshnessProofRequired         bool   `json:"freshness_proof_required"`
	FreshnessProofRef              string `json:"freshness_proof_ref"`
	IdentityInferredFromNameOrPath bool   `json:"identity_inferred_from_name_or_path"`
}

type Point15Val0TimestampDiscipline struct {
	DisciplineID                string `json:"discipline_id"`
	TenantScope                 string `json:"tenant_scope"`
	FreshnessStatus             string `json:"freshness_status"`
	ObservedAt                  string `json:"observed_at"`
	ObservedTimeSource          string `json:"observed_time_source"`
	EvaluatedAt                 string `json:"evaluated_at"`
	EvaluatedTimeSource         string `json:"evaluated_time_source"`
	ValidatedAt                 string `json:"validated_at"`
	ValidatedTimeSource         string `json:"validated_time_source"`
	ReviewerApprovedAt          string `json:"reviewer_approved_at"`
	ReviewerApprovedTimeSource  string `json:"reviewer_approved_time_source"`
	SourceEventAt               string `json:"source_event_at"`
	SourceEventTimeSource       string `json:"source_event_time_source"`
	ReferenceNow                string `json:"reference_now"`
	ReferenceNowTimeSource      string `json:"reference_now_time_source"`
	ClientLocalCreatesCanonical bool   `json:"client_local_creates_canonical"`
	SourceEventCreatesCanonical bool   `json:"source_event_creates_canonical"`
}

type Point15Val0AuthorityBoundary struct {
	BoundaryID                         string `json:"boundary_id"`
	TenantScope                        string `json:"tenant_scope"`
	ExternalSourceInputOnly            bool   `json:"external_source_input_only"`
	AgentRecommendationAdvisoryOnly    bool   `json:"agent_recommendation_advisory_only"`
	SchedulerPassAllowed               bool   `json:"scheduler_pass_allowed"`
	FreshnessAuthorityAllowed          bool   `json:"freshness_authority_allowed"`
	DashboardFreshnessAllowed          bool   `json:"dashboard_freshness_allowed"`
	AgentFreshnessAllowed              bool   `json:"agent_freshness_allowed"`
	ConnectorFreshnessAuthorityAllowed bool   `json:"connector_freshness_authority_allowed"`
	CustomerProjectionMutatesFreshness bool   `json:"customer_projection_mutates_freshness"`
	AuditorProjectionMutatesFreshness  bool   `json:"auditor_projection_mutates_freshness"`
	PortalProjectionMutatesFreshness   bool   `json:"portal_projection_mutates_freshness"`
	CanonicalMutationAllowed           bool   `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed          bool   `json:"production_mutation_allowed"`
	PassAllowed                        bool   `json:"pass_allowed"`
}

type Point15Val0NoOverclaimGuard struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	FreshnessDisclaimer                  string   `json:"freshness_disclaimer"`
}

func point15Val0States() []string {
	return []string{
		Point15Val0StateActive,
		Point15Val0StateBlocked,
		Point15Val0StateReviewRequired,
		Point15Val0StateIncomplete,
	}
}

func point15Val0StateValid(value string) bool {
	return point14Val0ExactValueValid(value, point15Val0States())
}

func point15Val0FreshnessStatuses() []string {
	return []string{
		point15Val0FreshnessFresh,
		point15Val0FreshnessStale,
		point15Val0FreshnessExpired,
		point15Val0FreshnessRevoked,
		point15Val0FreshnessSuperseded,
		point15Val0FreshnessDrifted,
		point15Val0FreshnessMissing,
		point15Val0FreshnessUnsupported,
		point15Val0FreshnessTampered,
	}
}

func point15Val0DowngradeOutcomes() []string {
	return []string{
		point15Val0DowngradeRetainActive,
		point15Val0DowngradeReview,
		point15Val0DowngradeBlocked,
		point15Val0DowngradeIncomplete,
	}
}

func point15Val0FreshnessStatusValid(value string) bool {
	return point14Val0ExactValueValid(value, point15Val0FreshnessStatuses())
}

func point15Val0DowngradeOutcomeValid(value string) bool {
	return point14Val0ExactValueValid(value, point15Val0DowngradeOutcomes())
}

func point15Val0ForbiddenWording() []string {
	return []string{
		"certified",
		"guaranteed secure",
		"continuous assurance guaranteed",
		"regulator-approved",
		"compliance guaranteed",
		"production approved",
		"deployment approved",
		"public badge",
		"global truth",
		"official authority",
		"legal proof",
		"financial guarantee",
	}
}

func point15Val0SafeWording() []string {
	return []string{
		"freshness status bounded by evidence context",
		"stale evidence requires review",
		"expired evidence blocks active closure",
		"external sources remain evidence input only",
		"point 15 val 0 verifies freshness discipline only",
	}
}

func point15Val0ObservedTextContainsForbiddenWording(text string) bool {
	return point14Val0ContainsForbiddenWordingFor(text, point15Val0SafeWording(), point15Val0ForbiddenWording())
}

func point15Val0ObservedListContainsForbiddenWording(values []string) bool {
	return point14Val0ListContainsForbiddenWordingFor(values, point15Val0SafeWording(), point15Val0ForbiddenWording())
}

func point15Val0DependencyRefValid(value string) bool {
	return point14Val0RefValid(value, "point15_val0_", "freshness_", "downgrade_", "authority_", "timestamp_")
}

func point15Val0EvidenceIDValid(value string) bool {
	return point14Val0RefValid(value, "evidence_")
}

func point15Val0SourceIDValid(value string) bool {
	return point14Val0RefValid(value, "source_")
}

func point15Val0FreshnessProofRefValid(value string) bool {
	return point14Val0RefValid(value, "freshness_proof_")
}

func point15Val0LineageRefValid(value string) bool {
	return point14Val0RefValid(value, "lineage_", "supersession_lineage_")
}

func point15Val0ValEPayloadContainsPoint15Pass(valE Point14ValEFoundation) bool {
	payload, err := json.Marshal(valE)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point15Val0BlockedPassToken)
}

func point15Val0ExpectedMappedState(status string) string {
	switch status {
	case point15Val0FreshnessFresh:
		return Point15Val0StateActive
	case point15Val0FreshnessStale, point15Val0FreshnessSuperseded, point15Val0FreshnessDrifted:
		return Point15Val0StateReviewRequired
	case point15Val0FreshnessMissing:
		return Point15Val0StateIncomplete
	case point15Val0FreshnessExpired, point15Val0FreshnessRevoked, point15Val0FreshnessUnsupported, point15Val0FreshnessTampered:
		return Point15Val0StateBlocked
	default:
		return Point15Val0StateBlocked
	}
}

func point15Val0ExpectedMappedDowngrade(status string) string {
	switch status {
	case point15Val0FreshnessFresh:
		return point15Val0DowngradeRetainActive
	case point15Val0FreshnessStale, point15Val0FreshnessSuperseded, point15Val0FreshnessDrifted:
		return point15Val0DowngradeReview
	case point15Val0FreshnessMissing:
		return point15Val0DowngradeIncomplete
	case point15Val0FreshnessExpired, point15Val0FreshnessRevoked, point15Val0FreshnessUnsupported, point15Val0FreshnessTampered:
		return point15Val0DowngradeBlocked
	default:
		return ""
	}
}

func point15Val0ExpectedDowngradeOutcome(model Point15Val0DowngradeTaxonomy) string {
	switch model.FreshnessStatus {
	case point15Val0FreshnessFresh:
		return point15Val0DowngradeRetainActive
	case point15Val0FreshnessStale:
		return point15Val0DowngradeReview
	case point15Val0FreshnessExpired, point15Val0FreshnessRevoked, point15Val0FreshnessUnsupported, point15Val0FreshnessTampered:
		return point15Val0DowngradeBlocked
	case point15Val0FreshnessSuperseded:
		if !point15Val0LineageRefValid(model.SupersessionLineageRef) {
			return point15Val0DowngradeBlocked
		}
		return point15Val0DowngradeReview
	case point15Val0FreshnessDrifted:
		if model.DriftIsDecisive {
			return point15Val0DowngradeBlocked
		}
		return point15Val0DowngradeReview
	case point15Val0FreshnessMissing:
		if model.MissingFreshnessProofDecisive {
			return point15Val0DowngradeBlocked
		}
		return point15Val0DowngradeIncomplete
	default:
		return ""
	}
}

func point15Val0ExpectedDowngradeState(model Point15Val0DowngradeTaxonomy) string {
	switch point15Val0ExpectedDowngradeOutcome(model) {
	case point15Val0DowngradeRetainActive:
		return Point15Val0StateActive
	case point15Val0DowngradeReview:
		return Point15Val0StateReviewRequired
	case point15Val0DowngradeIncomplete:
		return Point15Val0StateIncomplete
	case point15Val0DowngradeBlocked:
		return Point15Val0StateBlocked
	default:
		return Point15Val0StateBlocked
	}
}

func point15Val0FreshnessStatusAllowsContextOutcome(status, outcome string) bool {
	switch status {
	case point15Val0FreshnessFresh:
		return outcome == point15Val0DowngradeRetainActive
	case point15Val0FreshnessStale:
		return outcome == point15Val0DowngradeReview
	case point15Val0FreshnessExpired, point15Val0FreshnessRevoked, point15Val0FreshnessUnsupported, point15Val0FreshnessTampered:
		return outcome == point15Val0DowngradeBlocked
	case point15Val0FreshnessSuperseded, point15Val0FreshnessDrifted:
		return outcome == point15Val0DowngradeReview || outcome == point15Val0DowngradeBlocked
	case point15Val0FreshnessMissing:
		return outcome == point15Val0DowngradeIncomplete || outcome == point15Val0DowngradeBlocked
	default:
		return false
	}
}

func point15Val0CommandsRun() []string {
	return []string{
		"git diff --check",
		"gofmt on changed Go files",
		"go test ./internal/formal -run 'Test.*Point15Val0.*|Test.*Point15.*Val0.*' -v",
		"go test ./internal/formal -run 'Test.*Point14ValE.*|Test.*Point14.*ValE.*' -v",
		"go test ./internal/formal -run 'TestPoint10ThroughPoint14CurrentSweep' -v",
		"go test ./internal/formal -run 'Test.*Point14.*' -v",
		"go test ./internal/formal -run 'Test.*Point13.*' -v",
		"go test ./internal/formal -run 'Test.*Point12.*|Test.*Replay.*|Test.*ProofPack.*|Test.*Binding.*|Test.*Mutation.*' -v",
		"go test ./internal/formal -run 'Test.*Point11.*|Test.*Claim.*|Test.*NoOverclaim.*|Test.*Governance.*' -v",
		"go test ./internal/formal -run 'Test.*AI.*|Test.*Agent.*|Test.*Lineage.*|Test.*Provenance.*' -v",
		"go test -timeout 20m ./...",
	}
}

func point15Val0GrepScans() []string {
	return []string{
		"point_15_pass scan",
		"forbidden wording scan",
		"ai authority scan",
		"mutation authority scan",
		"external api scan",
		"skip todo fixme scan",
	}
}

func point15Val0DependencySnapshotFromUpstream(valE Point14ValEFoundation) Point15Val0DependencySnapshot {
	return Point15Val0DependencySnapshot{
		Point14ValECurrentState:              valE.CurrentState,
		Point14ValEDependencyState:           valE.DependencyState,
		Point14ValEClosureEvaluatorState:     valE.ClosureEvaluatorState,
		Point14ValEPassClosureManifestState:  valE.PassClosureManifestState,
		Point14PassAllowed:                   valE.Point14PassAllowed,
		Point14PassToken:                     valE.Point14PassToken,
		Point14PassManifestPointID:           valE.PassClosureManifest.PointID,
		Point14PassManifestWaveID:            valE.PassClosureManifest.WaveID,
		Point14PassManifestClosureToken:      valE.PassClosureManifest.ClosureToken,
		Point14ValEComputedFromUpstream:      valE.Dependency.SnapshotFromComputedOutput,
		Point14ValEMerged:                    true,
		Point14ValECIGreen:                   true,
		Point14ValEReviewedOnMain:            true,
		Point15PassSeen:                      point15Val0ValEPayloadContainsPoint15Pass(valE),
		InheritedPoint13ValECurrentState:     valE.Dependency.InheritedPoint13ValECurrentState,
		InheritedPoint13ValEPassClosureState: valE.Dependency.InheritedPoint13ValEPassClosureState,
		InheritedPoint12CurrentState:         valE.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:      valE.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:     valE.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint11CurrentState:         valE.Dependency.InheritedPoint11CurrentState,
		InheritedPoint11PublicationState:     valE.Dependency.InheritedPoint11PublicationState,
		InheritedPoint11NoOverclaimState:     valE.Dependency.InheritedPoint11NoOverclaimState,
		InheritedPoint11FinalPassGateState:   valE.Dependency.InheritedPoint11FinalPassGateState,
		InheritedPoint10CurrentState:         valE.Dependency.InheritedPoint10CurrentState,
		InheritedPoint10NoOverclaimState:     valE.Dependency.InheritedPoint10NoOverclaimState,
		InheritedPoint10ProjectionState:      valE.Dependency.InheritedPoint10ProjectionState,
		InheritedPoint10PassRuleState:        valE.Dependency.InheritedPoint10PassRuleState,
		InheritedTenantScope:                 valE.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:           true,
		ReviewPrerequisites:                  append([]string{}, valE.ReviewPrerequisites...),
		Point14ValE:                          valE,
	}
}

func point15Val0DependencySnapshotModel() Point15Val0DependencySnapshot {
	valE := ComputePoint14ValEFoundation(Point14ValEFoundationModel())
	return point15Val0DependencySnapshotFromUpstream(valE)
}

func EvaluatePoint15Val0DependencyState(model Point15Val0DependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValEComputedFromUpstream ||
		!model.Point14ValEMerged ||
		!model.Point14ValECIGreen ||
		!model.Point14ValEReviewedOnMain ||
		model.Point15PassSeen ||
		!point14ValEStateValid(model.Point14ValECurrentState) ||
		!point14ValEStateValid(model.Point14ValEDependencyState) ||
		!point14ValEStateValid(model.Point14ValEClosureEvaluatorState) ||
		!point14ValEStateValid(model.Point14ValEPassClosureManifestState) ||
		!point13ValEStateValid(model.InheritedPoint13ValECurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValEPassClosureState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		model.InheritedPoint11CurrentState == "" ||
		model.InheritedPoint11PublicationState == "" ||
		model.InheritedPoint11NoOverclaimState == "" ||
		model.InheritedPoint11FinalPassGateState == "" ||
		model.InheritedPoint10CurrentState == "" ||
		model.InheritedPoint10NoOverclaimState == "" ||
		model.InheritedPoint10ProjectionState == "" ||
		model.InheritedPoint10PassRuleState == "" ||
		!point14ValEFoundationComputedPassConfirmed(model.Point14ValE) ||
		!point14ValDFoundationEmbeddedSnapshotCopiesExact(model.Point14ValE.Dependency.Point14ValD) ||
		!point14Val0Point11FoundationActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point14Val0.Dependency.Point11) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point15Val0StateBlocked
	}
	if model.Point14ValECurrentState != model.Point14ValE.CurrentState ||
		model.Point14ValEDependencyState != model.Point14ValE.DependencyState ||
		model.Point14ValEClosureEvaluatorState != model.Point14ValE.ClosureEvaluatorState ||
		model.Point14ValEPassClosureManifestState != model.Point14ValE.PassClosureManifestState ||
		model.Point14ValEComputedFromUpstream != model.Point14ValE.Dependency.SnapshotFromComputedOutput ||
		model.Point14PassAllowed != model.Point14ValE.Point14PassAllowed ||
		model.Point14PassToken != model.Point14ValE.Point14PassToken ||
		model.Point14PassManifestPointID != model.Point14ValE.PassClosureManifest.PointID ||
		model.Point14PassManifestWaveID != model.Point14ValE.PassClosureManifest.WaveID ||
		model.Point14PassManifestClosureToken != model.Point14ValE.PassClosureManifest.ClosureToken ||
		model.InheritedPoint13ValECurrentState != model.Point14ValE.Dependency.InheritedPoint13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14ValE.Dependency.InheritedPoint13ValEPassClosureState ||
		model.InheritedPoint12CurrentState != model.Point14ValE.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14ValE.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14ValE.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint11CurrentState != model.Point14ValE.Dependency.InheritedPoint11CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValE.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValE.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValE.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint11CurrentState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14ValE.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14ValE.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14ValE.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14ValE.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14ValE.Dependency.InheritedTenantScope {
		return Point15Val0StateBlocked
	}
	if model.Point14ValECurrentState != Point14ValEStatePassConfirmed ||
		model.Point14ValEDependencyState != Point14ValEStatePassConfirmed ||
		model.Point14ValEClosureEvaluatorState != Point14ValEStatePassConfirmed ||
		model.Point14ValEPassClosureManifestState != Point14ValEStatePassConfirmed ||
		!model.Point14PassAllowed ||
		model.Point14PassToken != point14Val0BlockedPassToken ||
		model.Point14PassManifestPointID != point14Val0PointID ||
		model.Point14PassManifestWaveID != point14ValEWaveID ||
		model.Point14PassManifestClosureToken != point14Val0BlockedPassToken ||
		model.InheritedPoint13ValECurrentState != Point13ValEStatePassConfirmed ||
		model.InheritedPoint13ValEPassClosureState != Point13ValEStateActive ||
		model.InheritedPoint12CurrentState != Point12ValEStatePassConfirmed ||
		model.InheritedPoint12DependencyState != Point12ValEStateActive ||
		model.InheritedPoint12PassClosureState != Point12ValEStateActive ||
		model.InheritedPoint11CurrentState != Point11ValDStateActive ||
		model.InheritedPoint11PublicationState != Point11ValDPublicationReviewStateActive ||
		model.InheritedPoint11NoOverclaimState != Point11ValDNoOverclaimReviewStateActive ||
		model.InheritedPoint11FinalPassGateState != Point11ValDFinalPassGateStateActive ||
		model.InheritedPoint10CurrentState != operability.DeploymentMultiTenantPoint10StatePass ||
		model.InheritedPoint10NoOverclaimState != operability.DeploymentMultiTenantValENoOverclaimStateActive ||
		model.InheritedPoint10ProjectionState != operability.DeploymentMultiTenantValEProjectionBoundaryStateActive ||
		model.InheritedPoint10PassRuleState != operability.DeploymentMultiTenantValEPoint10PassRuleStateActive {
		return Point15Val0StateBlocked
	}
	return Point15Val0StateActive
}

func point15Val0Point11FoundationSnapshotActive(model Point11ValDFoundation) bool {
	dependencyState := EvaluatePoint11ValDDependencyState(Point11ValDDependencyBundle{
		Val0: model.Val0Dependency,
		ValA: model.ValADependency,
		ValB: model.ValBDependency,
		ValC: model.ValCDependency,
	})
	integratedInvariantState := EvaluatePoint11ValDIntegratedInvariantState(model.IntegratedInvariantReview)
	qualityMapState := EvaluatePoint11ValDQualityMapState(model.QualityMap)
	publicationReviewState := EvaluatePoint11ValDPublicationReviewState(model.PublicationReview)
	noOverclaimReviewState := EvaluatePoint11ValDNoOverclaimReviewState(model.NoOverclaimReview)
	cleanRoomIPReviewState := EvaluatePoint11ValDCleanRoomIPReviewState(model.CleanRoomIPReview)
	clbClosureState := EvaluatePoint11ValDCLBClosureState(model.CLBLedger)
	passClosureManifestState := EvaluatePoint11ValDPassClosureManifestState(model.PassClosureManifest, model)
	finalPassGateState := EvaluatePoint11ValDFinalPassGateState(model.FinalPassGate, model)

	evaluated := model
	evaluated.DependencyState = dependencyState
	evaluated.IntegratedInvariantState = integratedInvariantState
	evaluated.QualityMapState = qualityMapState
	evaluated.PublicationReviewState = publicationReviewState
	evaluated.NoOverclaimReviewState = noOverclaimReviewState
	evaluated.CleanRoomIPReviewState = cleanRoomIPReviewState
	evaluated.CLBClosureState = clbClosureState
	evaluated.PassClosureManifestState = passClosureManifestState
	evaluated.PassClosureManifest.CurrentState = passClosureManifestState
	evaluated.FinalPassGateState = finalPassGateState
	evaluated.FinalPassGate.CurrentState = finalPassGateState

	return model.CurrentState == Point11ValDStateActive &&
		EvaluatePoint11ValDFoundationState(evaluated) == Point11ValDStateActive &&
		model.DependencyState == dependencyState &&
		model.IntegratedInvariantState == integratedInvariantState &&
		model.QualityMapState == qualityMapState &&
		model.PublicationReviewState == publicationReviewState &&
		model.NoOverclaimReviewState == noOverclaimReviewState &&
		model.CleanRoomIPReviewState == cleanRoomIPReviewState &&
		model.CLBClosureState == clbClosureState &&
		model.PassClosureManifestState == passClosureManifestState &&
		model.FinalPassGateState == finalPassGateState &&
		dependencyState == Point11ValDDependencyStateActive &&
		integratedInvariantState == Point11ValDIntegratedInvariantStateActive &&
		qualityMapState == Point11ValDQualityMapStateActive &&
		publicationReviewState == Point11ValDPublicationReviewStateActive &&
		noOverclaimReviewState == Point11ValDNoOverclaimReviewStateActive &&
		cleanRoomIPReviewState == Point11ValDCleanRoomIPReviewStateActive &&
		clbClosureState == Point11ValDCLBClosureStateActive &&
		passClosureManifestState == Point11ValDPassClosureManifestStateActive &&
		finalPassGateState == Point11ValDFinalPassGateStateActive &&
		model.Point11PassToken == point11ValDPoint11PassToken &&
		model.PassClosureManifest.CurrentState == passClosureManifestState &&
		model.PassClosureManifest.Point11PassAllowed &&
		model.PassClosureManifest.Point11PassToken == point11ValDPoint11PassToken &&
		model.FinalPassGate.CurrentState == finalPassGateState &&
		model.FinalPassGate.Point11PassAllowed &&
		model.FinalPassGate.Point11PassEmitted &&
		model.FinalPassGate.Point11PassToken == point11ValDPoint11PassToken
}

func point15Val0Point14ValDDependencySnapshotActive(model Point14ValDDependencySnapshot) bool {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValCComputedFromUpstream ||
		!model.Point14ValCMerged ||
		!model.Point14ValCCIGreen ||
		!model.Point14ValCReviewedOnMain ||
		model.Point14PassSeen ||
		model.Point14ValCPointID != point14Val0PointID ||
		model.Point14ValCWaveID != point14ValCWaveID ||
		!point14ValCStateValid(model.Point14ValCCurrentState) ||
		!point14ValCStateValid(model.Point14ValCDependencyState) ||
		!point14ValCStateValid(model.Point14ValCCorrectionNoticeState) ||
		!point14ValCStateValid(model.Point14ValCRevocationRequestState) ||
		!point14ValCStateValid(model.Point14ValCSupersessionRecordState) ||
		!point14ValCStateValid(model.Point14ValCPublicationApprovalState) ||
		!point14ValCStateValid(model.Point14ValCVisibilityBoundaryState) ||
		!point14ValCStateValid(model.Point14ValCTenantPrivacyState) ||
		!point14ValCStateValid(model.Point14ValCRedactionLimitationState) ||
		!point14ValCStateValid(model.Point14ValCGovernanceTraceState) ||
		!point14ValCStateValid(model.Point14ValCAgentPublicationBoundaryState) ||
		!point14ValCStateValid(model.Point14ValCNoExternalAuthorityState) ||
		!point14ValCStateValid(model.Point14ValCNoOverclaimState) ||
		!point14ValBStateValid(model.InheritedPoint14ValBCurrentState) ||
		!point14ValAStateValid(model.InheritedPoint14ValACurrentState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0CurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValECurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValEPassClosureState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		model.InheritedPoint11CurrentState == "" ||
		model.InheritedPoint11PublicationState == "" ||
		model.InheritedPoint11NoOverclaimState == "" ||
		model.InheritedPoint11FinalPassGateState == "" ||
		model.InheritedPoint10CurrentState == "" ||
		model.InheritedPoint10NoOverclaimState == "" ||
		model.InheritedPoint10ProjectionState == "" ||
		model.InheritedPoint10PassRuleState == "" ||
		!point14ValDDependencyEmbeddedSnapshotCopiesExact(model) ||
		!point14ValCDependencyChainComputedActive(model.Point14ValC) ||
		!point14ValBDependencyChainComputedActive(model.Point14ValB) ||
		!point14ValAFoundationComputedActive(model.Point14ValA) ||
		!point14Val0FoundationComputedActive(model.Point14Val0) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValC.Dependency.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValC.Dependency.Point14ValB.Dependency.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point14Val0.Dependency.Point11) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return false
	}
	if model.Point14ValCCurrentState != model.Point14ValC.CurrentState ||
		model.Point14ValCDependencyState != model.Point14ValC.DependencyState ||
		model.Point14ValCCorrectionNoticeState != model.Point14ValC.CorrectionNoticeState ||
		model.Point14ValCRevocationRequestState != model.Point14ValC.RevocationRequestState ||
		model.Point14ValCSupersessionRecordState != model.Point14ValC.SupersessionRecordState ||
		model.Point14ValCPublicationApprovalState != model.Point14ValC.PublicationApprovalState ||
		model.Point14ValCVisibilityBoundaryState != model.Point14ValC.VisibilityBoundaryState ||
		model.Point14ValCTenantPrivacyState != model.Point14ValC.TenantPrivacyState ||
		model.Point14ValCRedactionLimitationState != model.Point14ValC.RedactionLimitationState ||
		model.Point14ValCGovernanceTraceState != model.Point14ValC.GovernanceTraceState ||
		model.Point14ValCAgentPublicationBoundaryState != model.Point14ValC.AgentPublicationBoundaryState ||
		model.Point14ValCNoExternalAuthorityState != model.Point14ValC.NoExternalAuthorityState ||
		model.Point14ValCNoOverclaimState != model.Point14ValC.NoOverclaimState ||
		model.Point14ValCComputedFromUpstream != model.Point14ValC.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint14ValBCurrentState != model.Point14ValC.Dependency.Point14ValBCurrentState ||
		model.InheritedPoint14ValACurrentState != model.Point14ValC.Dependency.InheritedPoint14ValACurrentState ||
		model.InheritedPoint14Val0CurrentState != model.Point14ValC.Dependency.InheritedPoint14Val0CurrentState ||
		model.InheritedPoint13ValECurrentState != model.Point14ValC.Dependency.InheritedPoint13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14ValC.Dependency.InheritedPoint13ValEPassClosureState ||
		model.InheritedPoint13ValEPassAllowed != model.Point14ValC.Dependency.InheritedPoint13ValEPassAllowed ||
		model.InheritedPoint13ValEPassToken != model.Point14ValC.Dependency.InheritedPoint13ValEPassToken ||
		model.InheritedPoint12CurrentState != model.Point14ValC.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14ValC.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14ValC.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != model.Point14ValC.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedPoint11CurrentState != model.Point14ValC.Dependency.InheritedPoint11CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValC.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValC.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValC.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint11CurrentState != model.Point14ValC.Dependency.Point11.CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValC.Dependency.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValC.Dependency.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValC.Dependency.Point11.FinalPassGateState ||
		model.InheritedPoint11CurrentState != model.Point11.CurrentState ||
		model.InheritedPoint11PublicationState != model.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point11.FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14ValC.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14ValC.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14ValC.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14ValC.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14ValC.Dependency.InheritedTenantScope {
		return false
	}
	return model.Point14ValCCurrentState == Point14ValCStateActive &&
		model.Point14ValCDependencyState == Point14ValCStateActive &&
		model.Point14ValCCorrectionNoticeState == Point14ValCStateActive &&
		model.Point14ValCRevocationRequestState == Point14ValCStateActive &&
		model.Point14ValCSupersessionRecordState == Point14ValCStateActive &&
		model.Point14ValCPublicationApprovalState == Point14ValCStateActive &&
		model.Point14ValCVisibilityBoundaryState == Point14ValCStateActive &&
		model.Point14ValCTenantPrivacyState == Point14ValCStateActive &&
		model.Point14ValCRedactionLimitationState == Point14ValCStateActive &&
		model.Point14ValCGovernanceTraceState == Point14ValCStateActive &&
		model.Point14ValCAgentPublicationBoundaryState == Point14ValCStateActive &&
		model.Point14ValCNoExternalAuthorityState == Point14ValCStateActive &&
		model.Point14ValCNoOverclaimState == Point14ValCStateActive &&
		model.InheritedPoint14ValBCurrentState == Point14ValBStateActive &&
		model.InheritedPoint14ValACurrentState == Point14ValAStateActive &&
		model.InheritedPoint14Val0CurrentState == Point14Val0StateActive &&
		model.InheritedPoint13ValECurrentState == Point13ValEStatePassConfirmed &&
		model.InheritedPoint13ValEPassClosureState == Point13ValEStateActive &&
		model.InheritedPoint13ValEPassAllowed &&
		model.InheritedPoint13ValEPassToken == point13ValEPoint13PassToken &&
		model.InheritedPoint12CurrentState == Point12ValEStatePassConfirmed &&
		model.InheritedPoint12DependencyState == Point12ValEStateActive &&
		model.InheritedPoint12PassClosureState == Point12ValEStateActive &&
		model.InheritedPoint12ReviewerResult == point12ValEReviewerResultPassConfirmed &&
		model.InheritedPoint11CurrentState == Point11ValDStateActive &&
		model.InheritedPoint11PublicationState == Point11ValDPublicationReviewStateActive &&
		model.InheritedPoint11NoOverclaimState == Point11ValDNoOverclaimReviewStateActive &&
		model.InheritedPoint11FinalPassGateState == Point11ValDFinalPassGateStateActive &&
		model.InheritedPoint10CurrentState == operability.DeploymentMultiTenantPoint10StatePass &&
		model.InheritedPoint10NoOverclaimState == operability.DeploymentMultiTenantValENoOverclaimStateActive &&
		model.InheritedPoint10ProjectionState == operability.DeploymentMultiTenantValEProjectionBoundaryStateActive &&
		model.InheritedPoint10PassRuleState == operability.DeploymentMultiTenantValEPoint10PassRuleStateActive
}

func point15Val0Point14ValDChainSnapshotActive(model Point14ValDFoundation) bool {
	timelineProjectionState := EvaluatePoint14ValDTimelineProjectionState(model.TimelineProjection)
	signalTimelineEntryState := EvaluatePoint14ValDSignalTimelineEntryState(model.SignalTimelineEntry)
	disputeTimelineState := EvaluatePoint14ValDDisputeTimelineProjectionState(model.DisputeTimelineProjection)
	correctionReadProjectionState := EvaluatePoint14ValDCorrectionReadProjectionState(model.CorrectionReadProjection)
	governanceTraceProjectionState := EvaluatePoint14ValDGovernanceTraceProjectionState(model.GovernanceTraceProjection)
	queryProjectionState := EvaluatePoint14ValDQueryProjectionState(model.QueryProjection)
	accessBoundaryState := EvaluatePoint14ValDAccessBoundaryState(model.AccessBoundary, model.Dependency)
	tenantPrivacyTimelineState := EvaluatePoint14ValDTenantPrivacyTimelineProjectionGuardState(model.TenantPrivacyTimelineProjectionGuard, model.Dependency)
	agentTimelineProjectionState := EvaluatePoint14ValDAgentTimelineProjectionState(model.AgentTimelineProjection, model.Dependency)
	timestampIntegrityState := EvaluatePoint14ValDTimestampIntegrityGuardState(model.TimestampIntegrityGuard)
	noMutationProjectionGuardState := EvaluatePoint14ValDNoMutationProjectionGuardState(model.NoMutationProjectionGuard)
	noOverclaimTimelineWordingState := EvaluatePoint14ValDNoOverclaimTimelineWordingState(model.NoOverclaimTimelineWording)
	dependencySnapshotActive := point15Val0Point14ValDDependencySnapshotActive(model.Dependency)

	return model.CurrentState == Point14ValDStateActive &&
		model.DependencyState == Point14ValDStateActive &&
		model.TimelineProjectionState == timelineProjectionState &&
		model.SignalTimelineEntryState == signalTimelineEntryState &&
		model.DisputeTimelineState == disputeTimelineState &&
		model.CorrectionReadProjectionState == correctionReadProjectionState &&
		model.GovernanceTraceProjectionState == governanceTraceProjectionState &&
		model.QueryProjectionState == queryProjectionState &&
		model.AccessBoundaryState == accessBoundaryState &&
		model.TenantPrivacyTimelineState == tenantPrivacyTimelineState &&
		model.AgentTimelineProjectionState == agentTimelineProjectionState &&
		model.TimestampIntegrityState == timestampIntegrityState &&
		model.NoMutationProjectionGuardState == noMutationProjectionGuardState &&
		model.NoOverclaimTimelineWordingState == noOverclaimTimelineWordingState &&
		timelineProjectionState == Point14ValDStateActive &&
		signalTimelineEntryState == Point14ValDStateActive &&
		disputeTimelineState == Point14ValDStateActive &&
		correctionReadProjectionState == Point14ValDStateActive &&
		governanceTraceProjectionState == Point14ValDStateActive &&
		queryProjectionState == Point14ValDStateActive &&
		accessBoundaryState == Point14ValDStateActive &&
		tenantPrivacyTimelineState == Point14ValDStateActive &&
		agentTimelineProjectionState == Point14ValDStateActive &&
		timestampIntegrityState == Point14ValDStateActive &&
		noMutationProjectionGuardState == Point14ValDStateActive &&
		noOverclaimTimelineWordingState == Point14ValDStateActive &&
		dependencySnapshotActive &&
		point14ValDFoundationEmbeddedSnapshotCopiesExact(model) &&
		EvaluatePoint14ValCNoOverclaimPublicationWordingState(model.Dependency.Point14ValC.NoOverclaimPublicationWording) == Point14ValCStateActive &&
		EvaluatePoint14ValBNoOverclaimDisputeWordingState(model.Dependency.Point14ValB.NoOverclaimDisputeWording) == Point14ValBStateActive &&
		EvaluatePoint14ValANoOverclaimValidationWordingState(model.Dependency.Point14ValA.NoOverclaimValidationWording) == Point14ValAStateActive
}

func point15Val0Point14ValEDependencySnapshotPassConfirmed(model Point14ValEDependencySnapshot) bool {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValDComputedFromUpstream ||
		!model.Point14ValDMerged ||
		!model.Point14ValDCIGreen ||
		!model.Point14ValDReviewedOnMain ||
		model.Point14PassSeen ||
		model.Point14ValDPointID != point14Val0PointID ||
		model.Point14ValDWaveID != point14ValDWaveID ||
		!point14ValDStateValid(model.Point14ValDCurrentState) ||
		!point14ValDStateValid(model.Point14ValDDependencyState) ||
		!point14ValDStateValid(model.Point14ValDTimelineProjectionState) ||
		!point14ValDStateValid(model.Point14ValDSignalTimelineEntryState) ||
		!point14ValDStateValid(model.Point14ValDDisputeTimelineState) ||
		!point14ValDStateValid(model.Point14ValDCorrectionReadProjectionState) ||
		!point14ValDStateValid(model.Point14ValDGovernanceTraceProjectionState) ||
		!point14ValDStateValid(model.Point14ValDQueryProjectionState) ||
		!point14ValDStateValid(model.Point14ValDAccessBoundaryState) ||
		!point14ValDStateValid(model.Point14ValDTenantPrivacyTimelineState) ||
		!point14ValDStateValid(model.Point14ValDAgentTimelineProjectionState) ||
		!point14ValDStateValid(model.Point14ValDTimestampIntegrityState) ||
		!point14ValDStateValid(model.Point14ValDNoMutationProjectionGuardState) ||
		!point14ValDStateValid(model.Point14ValDNoOverclaimTimelineWordingState) ||
		!point14ValCStateValid(model.InheritedPoint14ValCCurrentState) ||
		!point14ValBStateValid(model.InheritedPoint14ValBCurrentState) ||
		!point14ValAStateValid(model.InheritedPoint14ValACurrentState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0CurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValECurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValEPassClosureState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		model.InheritedPoint11CurrentState == "" ||
		model.InheritedPoint11PublicationState == "" ||
		model.InheritedPoint11NoOverclaimState == "" ||
		model.InheritedPoint11FinalPassGateState == "" ||
		model.InheritedPoint10CurrentState == "" ||
		model.InheritedPoint10NoOverclaimState == "" ||
		model.InheritedPoint10ProjectionState == "" ||
		model.InheritedPoint10PassRuleState == "" ||
		!point11Val0ScopeValid(model.InheritedTenantScope) ||
		!point15Val0Point14ValDChainSnapshotActive(model.Point14ValD) {
		return false
	}
	if model.Point14ValDCurrentState != model.Point14ValD.CurrentState ||
		model.Point14ValDDependencyState != model.Point14ValD.DependencyState ||
		model.Point14ValDTimelineProjectionState != model.Point14ValD.TimelineProjectionState ||
		model.Point14ValDSignalTimelineEntryState != model.Point14ValD.SignalTimelineEntryState ||
		model.Point14ValDDisputeTimelineState != model.Point14ValD.DisputeTimelineState ||
		model.Point14ValDCorrectionReadProjectionState != model.Point14ValD.CorrectionReadProjectionState ||
		model.Point14ValDGovernanceTraceProjectionState != model.Point14ValD.GovernanceTraceProjectionState ||
		model.Point14ValDQueryProjectionState != model.Point14ValD.QueryProjectionState ||
		model.Point14ValDAccessBoundaryState != model.Point14ValD.AccessBoundaryState ||
		model.Point14ValDTenantPrivacyTimelineState != model.Point14ValD.TenantPrivacyTimelineState ||
		model.Point14ValDAgentTimelineProjectionState != model.Point14ValD.AgentTimelineProjectionState ||
		model.Point14ValDTimestampIntegrityState != model.Point14ValD.TimestampIntegrityState ||
		model.Point14ValDNoMutationProjectionGuardState != model.Point14ValD.NoMutationProjectionGuardState ||
		model.Point14ValDNoOverclaimTimelineWordingState != model.Point14ValD.NoOverclaimTimelineWordingState ||
		model.Point14ValDComputedFromUpstream != model.Point14ValD.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint14ValCCurrentState != model.Point14ValD.Dependency.Point14ValCCurrentState ||
		model.InheritedPoint14ValBCurrentState != model.Point14ValD.Dependency.InheritedPoint14ValBCurrentState ||
		model.InheritedPoint14ValACurrentState != model.Point14ValD.Dependency.InheritedPoint14ValACurrentState ||
		model.InheritedPoint14Val0CurrentState != model.Point14ValD.Dependency.InheritedPoint14Val0CurrentState ||
		model.InheritedPoint13ValECurrentState != model.Point14ValD.Dependency.InheritedPoint13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14ValD.Dependency.InheritedPoint13ValEPassClosureState ||
		model.InheritedPoint13ValEPassAllowed != model.Point14ValD.Dependency.InheritedPoint13ValEPassAllowed ||
		model.InheritedPoint13ValEPassToken != model.Point14ValD.Dependency.InheritedPoint13ValEPassToken ||
		model.InheritedPoint12CurrentState != model.Point14ValD.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14ValD.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14ValD.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != model.Point14ValD.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedPoint11CurrentState != model.Point14ValD.Dependency.InheritedPoint11CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValD.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValD.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValD.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint11CurrentState != model.Point14ValD.Dependency.Point11.CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValD.Dependency.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValD.Dependency.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValD.Dependency.Point11.FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14ValD.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14ValD.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14ValD.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14ValD.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14ValD.Dependency.InheritedTenantScope {
		return false
	}
	return model.Point14ValDCurrentState == Point14ValDStateActive &&
		model.Point14ValDDependencyState == Point14ValDStateActive &&
		model.Point14ValDTimelineProjectionState == Point14ValDStateActive &&
		model.Point14ValDSignalTimelineEntryState == Point14ValDStateActive &&
		model.Point14ValDDisputeTimelineState == Point14ValDStateActive &&
		model.Point14ValDCorrectionReadProjectionState == Point14ValDStateActive &&
		model.Point14ValDGovernanceTraceProjectionState == Point14ValDStateActive &&
		model.Point14ValDQueryProjectionState == Point14ValDStateActive &&
		model.Point14ValDAccessBoundaryState == Point14ValDStateActive &&
		model.Point14ValDTenantPrivacyTimelineState == Point14ValDStateActive &&
		model.Point14ValDAgentTimelineProjectionState == Point14ValDStateActive &&
		model.Point14ValDTimestampIntegrityState == Point14ValDStateActive &&
		model.Point14ValDNoMutationProjectionGuardState == Point14ValDStateActive &&
		model.Point14ValDNoOverclaimTimelineWordingState == Point14ValDStateActive &&
		model.InheritedPoint14ValCCurrentState == Point14ValCStateActive &&
		model.InheritedPoint14ValBCurrentState == Point14ValBStateActive &&
		model.InheritedPoint14ValACurrentState == Point14ValAStateActive &&
		model.InheritedPoint14Val0CurrentState == Point14Val0StateActive &&
		model.InheritedPoint13ValECurrentState == Point13ValEStatePassConfirmed &&
		model.InheritedPoint13ValEPassClosureState == Point13ValEStateActive &&
		model.InheritedPoint13ValEPassAllowed &&
		model.InheritedPoint13ValEPassToken == point13ValEPoint13PassToken &&
		model.InheritedPoint12CurrentState == Point12ValEStatePassConfirmed &&
		model.InheritedPoint12DependencyState == Point12ValEStateActive &&
		model.InheritedPoint12PassClosureState == Point12ValEStateActive &&
		model.InheritedPoint12ReviewerResult == point12ValEReviewerResultPassConfirmed &&
		model.InheritedPoint11CurrentState == Point11ValDStateActive &&
		model.InheritedPoint11PublicationState == Point11ValDPublicationReviewStateActive &&
		model.InheritedPoint11NoOverclaimState == Point11ValDNoOverclaimReviewStateActive &&
		model.InheritedPoint11FinalPassGateState == Point11ValDFinalPassGateStateActive &&
		model.InheritedPoint10CurrentState == operability.DeploymentMultiTenantPoint10StatePass &&
		model.InheritedPoint10NoOverclaimState == operability.DeploymentMultiTenantValENoOverclaimStateActive &&
		model.InheritedPoint10ProjectionState == operability.DeploymentMultiTenantValEProjectionBoundaryStateActive &&
		model.InheritedPoint10PassRuleState == operability.DeploymentMultiTenantValEPoint10PassRuleStateActive
}

func point15Val0Point14ValEFoundationSnapshotPassConfirmed(model Point14ValEFoundation) bool {
	dependencyState := Point14ValEStateBlocked
	if point15Val0Point14ValEDependencySnapshotPassConfirmed(model.Dependency) {
		dependencyState = Point14ValEStatePassConfirmed
	}
	externalSignalState := EvaluatePoint14ValEExternalSignalValidationClosureCheckState(model.ExternalSignalValidationClosureCheck)
	conflictDisputeState := EvaluatePoint14ValEConflictDisputeClosureCheckState(model.ConflictDisputeClosureCheck)
	correctionPublicationState := EvaluatePoint14ValECorrectionPublicationClosureCheckState(model.CorrectionPublicationClosureCheck)
	timelineProjectionState := EvaluatePoint14ValETimelineProjectionClosureCheckState(model.TimelineProjectionClosureCheck)
	authorityBoundaryState := EvaluatePoint14ValEAuthorityBoundaryClosureCheckState(model.AuthorityBoundaryClosureCheck)
	tenantPrivacyState := EvaluatePoint14ValETenantPrivacyClosureCheckState(model.TenantPrivacyClosureCheck, model.Dependency)
	timestampIntegrityState := EvaluatePoint14ValETimestampIntegrityClosureCheckState(model.TimestampIntegrityClosureCheck, model.Dependency)
	agentAdvisoryState := EvaluatePoint14ValEAgentAdvisoryClosureCheckState(model.AgentAdvisoryClosureCheck, model.Dependency)
	derivedNoOverclaim := point14ValENoOverclaimFinalCheckModel(model.Dependency)
	noOverclaimState := Point14ValEStateBlocked
	if reflect.DeepEqual(model.NoOverclaimFinalCheck, derivedNoOverclaim) {
		noOverclaimState = EvaluatePoint14ValENoOverclaimFinalCheckState(model.NoOverclaimFinalCheck)
	}
	clbState := EvaluatePoint14ValECLBFinalCheckState(model.CLBFinalCheck)

	passCandidate := point14ValEFoundationState(
		dependencyState,
		externalSignalState,
		conflictDisputeState,
		correctionPublicationState,
		timelineProjectionState,
		authorityBoundaryState,
		tenantPrivacyState,
		timestampIntegrityState,
		agentAdvisoryState,
		noOverclaimState,
		clbState,
	) == Point14ValEStatePassConfirmed &&
		model.ClosureEvaluator.ReadOnlyProjectionConfirmed &&
		model.ClosureEvaluator.NoMutationPathsDetected &&
		model.ClosureEvaluator.NoExternalAuthorityDetected &&
		model.ClosureEvaluator.NoPrematurePoint14Pass &&
		model.ClosureEvaluator.ReviewerResult == point12ValEReviewerResultPassConfirmed

	closureEvaluator := model.ClosureEvaluator
	closureEvaluator.DependencyState = dependencyState
	closureEvaluator.ValidationClosureState = externalSignalState
	closureEvaluator.DisputeClosureState = conflictDisputeState
	closureEvaluator.CorrectionPublicationClosureState = correctionPublicationState
	closureEvaluator.TimelineProjectionClosureState = timelineProjectionState
	closureEvaluator.AuthorityBoundaryState = authorityBoundaryState
	closureEvaluator.TenantPrivacyState = tenantPrivacyState
	closureEvaluator.TimestampIntegrityState = timestampIntegrityState
	closureEvaluator.AgentAdvisoryState = agentAdvisoryState
	closureEvaluator.NoOverclaimState = noOverclaimState
	closureEvaluator.CLBFinalState = clbState
	closureEvaluator.FinalPassAllowed = passCandidate
	closureEvaluatorState := EvaluatePoint14ValEClosureEvaluatorState(closureEvaluator)
	closureEvaluator.CurrentState = closureEvaluatorState

	passManifest := model.PassClosureManifest
	passManifest.DependencyGateResult = dependencyState
	passManifest.ClosureEvaluatorResult = closureEvaluatorState
	passManifest.ProjectionBoundaryResult = timelineProjectionState
	passManifest.NoExternalAuthorityResult = authorityBoundaryState
	passManifest.NoOverclaimGrepResult = noOverclaimState
	passManifest.TenantPrivacyResult = tenantPrivacyState
	passManifest.TimestampIntegrityResult = timestampIntegrityState
	passManifest.AIAgentBoundaryResult = agentAdvisoryState
	passManifest.CLBResult = clbState
	passManifest.Point14PassAllowed = passCandidate && closureEvaluatorState == Point14ValEStatePassConfirmed
	if passManifest.Point14PassAllowed {
		passManifest.Point14PassToken = point14Val0BlockedPassToken
	} else {
		passManifest.Point14PassToken = ""
	}
	passManifestState := EvaluatePoint14PassClosureManifestState(passManifest)
	passManifest.CurrentState = passManifestState

	currentState := point14ValEFoundationState(
		dependencyState,
		externalSignalState,
		conflictDisputeState,
		correctionPublicationState,
		timelineProjectionState,
		authorityBoundaryState,
		tenantPrivacyState,
		timestampIntegrityState,
		agentAdvisoryState,
		noOverclaimState,
		clbState,
		closureEvaluatorState,
		passManifestState,
	)
	point14PassAllowed := currentState == Point14ValEStatePassConfirmed &&
		closureEvaluatorState == Point14ValEStatePassConfirmed &&
		passManifestState == Point14ValEStatePassConfirmed &&
		passManifest.Point14PassAllowed
	point14PassToken := ""
	if point14PassAllowed {
		point14PassToken = point14Val0BlockedPassToken
	}

	return currentState == Point14ValEStatePassConfirmed &&
		model.CurrentState == currentState &&
		model.DependencyState == dependencyState &&
		model.ExternalSignalValidationClosureState == externalSignalState &&
		model.ConflictDisputeClosureState == conflictDisputeState &&
		model.CorrectionPublicationClosureState == correctionPublicationState &&
		model.TimelineProjectionClosureState == timelineProjectionState &&
		model.AuthorityBoundaryClosureState == authorityBoundaryState &&
		model.TenantPrivacyClosureState == tenantPrivacyState &&
		model.TimestampIntegrityClosureState == timestampIntegrityState &&
		model.AgentAdvisoryClosureState == agentAdvisoryState &&
		model.NoOverclaimFinalCheckState == noOverclaimState &&
		model.CLBFinalCheckState == clbState &&
		model.ClosureEvaluatorState == closureEvaluatorState &&
		model.ClosureEvaluator.CurrentState == closureEvaluatorState &&
		reflect.DeepEqual(model.ClosureEvaluator, closureEvaluator) &&
		model.PassClosureManifestState == passManifestState &&
		model.PassClosureManifest.CurrentState == passManifestState &&
		reflect.DeepEqual(model.PassClosureManifest, passManifest) &&
		model.Point14PassAllowed == point14PassAllowed &&
		model.Point14PassToken == point14PassToken &&
		point14PassAllowed &&
		point14PassToken == point14Val0BlockedPassToken
}

func point15Val0EmbeddedDependencySnapshotActive(model Point15Val0DependencySnapshot) bool {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValEComputedFromUpstream ||
		!model.Point14ValEMerged ||
		!model.Point14ValECIGreen ||
		!model.Point14ValEReviewedOnMain ||
		model.Point15PassSeen ||
		!point14ValEStateValid(model.Point14ValECurrentState) ||
		!point14ValEStateValid(model.Point14ValEDependencyState) ||
		!point14ValEStateValid(model.Point14ValEClosureEvaluatorState) ||
		!point14ValEStateValid(model.Point14ValEPassClosureManifestState) ||
		!point13ValEStateValid(model.InheritedPoint13ValECurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValEPassClosureState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		model.InheritedPoint11CurrentState == "" ||
		model.InheritedPoint11PublicationState == "" ||
		model.InheritedPoint11NoOverclaimState == "" ||
		model.InheritedPoint11FinalPassGateState == "" ||
		model.InheritedPoint10CurrentState == "" ||
		model.InheritedPoint10NoOverclaimState == "" ||
		model.InheritedPoint10ProjectionState == "" ||
		model.InheritedPoint10PassRuleState == "" ||
		!point15Val0Point14ValEFoundationSnapshotPassConfirmed(model.Point14ValE) ||
		!point14ValDFoundationEmbeddedSnapshotCopiesExact(model.Point14ValE.Dependency.Point14ValD) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point11) ||
		!point15Val0Point11FoundationSnapshotActive(model.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point14Val0.Dependency.Point11) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return false
	}
	if model.Point14ValECurrentState != model.Point14ValE.CurrentState ||
		model.Point14ValEDependencyState != model.Point14ValE.DependencyState ||
		model.Point14ValEClosureEvaluatorState != model.Point14ValE.ClosureEvaluatorState ||
		model.Point14ValEPassClosureManifestState != model.Point14ValE.PassClosureManifestState ||
		model.Point14ValEComputedFromUpstream != model.Point14ValE.Dependency.SnapshotFromComputedOutput ||
		model.Point14PassAllowed != model.Point14ValE.Point14PassAllowed ||
		model.Point14PassToken != model.Point14ValE.Point14PassToken ||
		model.Point14PassManifestPointID != model.Point14ValE.PassClosureManifest.PointID ||
		model.Point14PassManifestWaveID != model.Point14ValE.PassClosureManifest.WaveID ||
		model.Point14PassManifestClosureToken != model.Point14ValE.PassClosureManifest.ClosureToken ||
		model.InheritedPoint13ValECurrentState != model.Point14ValE.Dependency.InheritedPoint13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14ValE.Dependency.InheritedPoint13ValEPassClosureState ||
		model.InheritedPoint12CurrentState != model.Point14ValE.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14ValE.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14ValE.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint11CurrentState != model.Point14ValE.Dependency.InheritedPoint11CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValE.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValE.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValE.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint11CurrentState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValE.Dependency.Point14ValD.Dependency.Point11.FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14ValE.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14ValE.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14ValE.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14ValE.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14ValE.Dependency.InheritedTenantScope {
		return false
	}
	return model.Point14ValECurrentState == Point14ValEStatePassConfirmed &&
		model.Point14ValEDependencyState == Point14ValEStatePassConfirmed &&
		model.Point14ValEClosureEvaluatorState == Point14ValEStatePassConfirmed &&
		model.Point14ValEPassClosureManifestState == Point14ValEStatePassConfirmed &&
		model.Point14PassAllowed &&
		model.Point14PassToken == point14Val0BlockedPassToken &&
		model.Point14PassManifestPointID == point14Val0PointID &&
		model.Point14PassManifestWaveID == point14ValEWaveID &&
		model.Point14PassManifestClosureToken == point14Val0BlockedPassToken &&
		model.InheritedPoint13ValECurrentState == Point13ValEStatePassConfirmed &&
		model.InheritedPoint13ValEPassClosureState == Point13ValEStateActive &&
		model.InheritedPoint12CurrentState == Point12ValEStatePassConfirmed &&
		model.InheritedPoint12DependencyState == Point12ValEStateActive &&
		model.InheritedPoint12PassClosureState == Point12ValEStateActive &&
		model.InheritedPoint11CurrentState == Point11ValDStateActive &&
		model.InheritedPoint11PublicationState == Point11ValDPublicationReviewStateActive &&
		model.InheritedPoint11NoOverclaimState == Point11ValDNoOverclaimReviewStateActive &&
		model.InheritedPoint11FinalPassGateState == Point11ValDFinalPassGateStateActive &&
		model.InheritedPoint10CurrentState == operability.DeploymentMultiTenantPoint10StatePass &&
		model.InheritedPoint10NoOverclaimState == operability.DeploymentMultiTenantValENoOverclaimStateActive &&
		model.InheritedPoint10ProjectionState == operability.DeploymentMultiTenantValEProjectionBoundaryStateActive &&
		model.InheritedPoint10PassRuleState == operability.DeploymentMultiTenantValEPoint10PassRuleStateActive
}

func point15Val0FreshnessTaxonomyModel() Point15Val0EvidenceFreshnessTaxonomy {
	return Point15Val0EvidenceFreshnessTaxonomy{
		TaxonomyID:               "point15_val0_freshness_taxonomy_001",
		FreshnessStatus:          point15Val0FreshnessFresh,
		MappedState:              Point15Val0StateActive,
		MappedDowngradeOutcome:   point15Val0DowngradeRetainActive,
		AllowedFreshnessStates:   point15Val0FreshnessStatuses(),
		AllowedDowngradeOutcomes: point15Val0DowngradeOutcomes(),
	}
}

func EvaluatePoint15Val0EvidenceFreshnessTaxonomyState(model Point15Val0EvidenceFreshnessTaxonomy) string {
	if !point15Val0DependencyRefValid(model.TaxonomyID) ||
		!point15Val0FreshnessStatusValid(model.FreshnessStatus) ||
		!point15Val0StateValid(model.MappedState) ||
		!point15Val0DowngradeOutcomeValid(model.MappedDowngradeOutcome) ||
		!point12Val0ExactStringSetMatch(model.AllowedFreshnessStates, point15Val0FreshnessStatuses()) ||
		!point12Val0ExactStringSetMatch(model.AllowedDowngradeOutcomes, point15Val0DowngradeOutcomes()) {
		return Point15Val0StateBlocked
	}
	if model.MappedState != point15Val0ExpectedMappedState(model.FreshnessStatus) ||
		model.MappedDowngradeOutcome != point15Val0ExpectedMappedDowngrade(model.FreshnessStatus) {
		return Point15Val0StateBlocked
	}
	return model.MappedState
}

func point15Val0DowngradeTaxonomyModel() Point15Val0DowngradeTaxonomy {
	return Point15Val0DowngradeTaxonomy{
		TaxonomyID:            "point15_val0_downgrade_taxonomy_001",
		FreshnessStatus:       point15Val0FreshnessFresh,
		DowngradeOutcome:      point15Val0DowngradeRetainActive,
		FreshnessProofPresent: true,
	}
}

func EvaluatePoint15Val0DowngradeTaxonomyState(model Point15Val0DowngradeTaxonomy) string {
	if !point15Val0DependencyRefValid(model.TaxonomyID) ||
		!point15Val0FreshnessStatusValid(model.FreshnessStatus) ||
		!point15Val0DowngradeOutcomeValid(model.DowngradeOutcome) {
		return Point15Val0StateBlocked
	}
	expectedOutcome := point15Val0ExpectedDowngradeOutcome(model)
	if expectedOutcome == "" {
		return Point15Val0StateBlocked
	}
	if model.FreshnessStatus == point15Val0FreshnessMissing {
		if model.FreshnessProofPresent {
			return Point15Val0StateBlocked
		}
	} else if !model.FreshnessProofPresent && expectedOutcome == point15Val0DowngradeRetainActive {
		return Point15Val0StateIncomplete
	}
	if model.FreshnessStatus == point15Val0FreshnessSuperseded && !point15Val0LineageRefValid(model.SupersessionLineageRef) {
		if model.DowngradeOutcome != point15Val0DowngradeBlocked {
			return Point15Val0StateBlocked
		}
		return Point15Val0StateBlocked
	}
	if model.RetainsPass || model.RetainsActiveClosure {
		if model.FreshnessStatus != point15Val0FreshnessFresh {
			return Point15Val0StateBlocked
		}
		if model.DowngradeOutcome != point15Val0DowngradeRetainActive {
			return Point15Val0StateBlocked
		}
	}
	if model.DowngradeOutcome != expectedOutcome {
		return Point15Val0StateBlocked
	}
	return point15Val0ExpectedDowngradeState(model)
}

func point15Val0FreshnessEvidenceContextModel(dependency Point15Val0DependencySnapshot) Point15Val0FreshnessEvidenceContext {
	return Point15Val0FreshnessEvidenceContext{
		ContextID:              "freshness_context_point15_val0_001",
		EvidenceID:             "evidence_point15_val0_001",
		TenantScope:            dependency.InheritedTenantScope,
		ReferencedTenantScope:  dependency.InheritedTenantScope,
		SourceID:               "source_point15_val0_001",
		EvidenceHash:           "hash_point15_val0_001",
		PolicyVersion:          "policy_version_point15_val0_v1",
		EngineVersion:          "engine_version_point15_val0_v1",
		SchemaVersion:          "schema_version_point15_val0_v1",
		ObservedAt:             "2026-05-06T18:15:00Z",
		ObservedTimeSource:     point14Val0TimeSourceServerUTC,
		EvaluatedAt:            "2026-05-06T18:20:00Z",
		EvaluatedTimeSource:    point14Val0TimeSourceServerUTC,
		ValidatedAt:            "2026-05-06T18:25:00Z",
		ValidatedTimeSource:    point14Val0TimeSourceServerUTC,
		FreshnessStatus:        point15Val0FreshnessFresh,
		DowngradeOutcome:       point15Val0DowngradeRetainActive,
		FreshnessProofRequired: true,
		FreshnessProofRef:      "freshness_proof_point15_val0_001",
	}
}

func EvaluatePoint15Val0FreshnessEvidenceContextState(model Point15Val0FreshnessEvidenceContext) string {
	if !point15Val0DependencyRefValid(model.ContextID) ||
		!point15Val0FreshnessStatusValid(model.FreshnessStatus) ||
		!point15Val0DowngradeOutcomeValid(model.DowngradeOutcome) ||
		(model.ReferencedTenantScope != "" && !point11Val0ScopeValid(model.ReferencedTenantScope)) ||
		(model.ObservedTimeSource != "" && !point14Val0TimeSourceValid(model.ObservedTimeSource)) ||
		(model.EvaluatedTimeSource != "" && !point14Val0TimeSourceValid(model.EvaluatedTimeSource)) ||
		(model.ValidatedTimeSource != "" && !point14Val0TimeSourceValid(model.ValidatedTimeSource)) {
		return Point15Val0StateBlocked
	}
	if model.IdentityInferredFromNameOrPath {
		return Point15Val0StateBlocked
	}
	if model.TenantScope == "" ||
		model.EvidenceID == "" ||
		model.SourceID == "" ||
		model.EvidenceHash == "" ||
		model.PolicyVersion == "" ||
		model.EngineVersion == "" ||
		model.SchemaVersion == "" ||
		model.ObservedAt == "" ||
		model.ObservedTimeSource == "" ||
		model.EvaluatedAt == "" ||
		model.EvaluatedTimeSource == "" ||
		model.ValidatedAt == "" ||
		model.ValidatedTimeSource == "" {
		return Point15Val0StateIncomplete
	}
	if model.ReferencedTenantScope != "" && model.ReferencedTenantScope != model.TenantScope {
		return Point15Val0StateBlocked
	}
	if !point11Val0ScopeValid(model.TenantScope) ||
		!point15Val0EvidenceIDValid(model.EvidenceID) ||
		!point15Val0SourceIDValid(model.SourceID) ||
		!point14Val0HashRefValid(model.EvidenceHash) ||
		!point12Val0VersionIdentityValid(model.PolicyVersion) ||
		!point12Val0VersionIdentityValid(model.EngineVersion) ||
		!point12Val0VersionIdentityValid(model.SchemaVersion) ||
		!point14Val0ParsedTimeOk(model.ObservedAt) ||
		!point14Val0ParsedTimeOk(model.EvaluatedAt) ||
		!point14Val0ParsedTimeOk(model.ValidatedAt) {
		return Point15Val0StateBlocked
	}
	if model.FreshnessProofRequired && !point15Val0FreshnessProofRefValid(model.FreshnessProofRef) {
		return Point15Val0StateIncomplete
	}
	if !point15Val0FreshnessStatusAllowsContextOutcome(model.FreshnessStatus, model.DowngradeOutcome) {
		return Point15Val0StateBlocked
	}
	switch model.FreshnessStatus {
	case point15Val0FreshnessFresh:
		return Point15Val0StateActive
	case point15Val0FreshnessStale, point15Val0FreshnessSuperseded, point15Val0FreshnessDrifted:
		if model.DowngradeOutcome == point15Val0DowngradeBlocked {
			return Point15Val0StateBlocked
		}
		return Point15Val0StateReviewRequired
	case point15Val0FreshnessMissing:
		if model.DowngradeOutcome == point15Val0DowngradeBlocked {
			return Point15Val0StateBlocked
		}
		return Point15Val0StateIncomplete
	default:
		return Point15Val0StateBlocked
	}
}

func EvaluatePoint15Val0TenantBoundaryState(model Point15Val0FreshnessEvidenceContext) string {
	if model.TenantScope == "" {
		return Point15Val0StateIncomplete
	}
	if !point11Val0ScopeValid(model.TenantScope) {
		return Point15Val0StateBlocked
	}
	if model.ReferencedTenantScope != "" && model.ReferencedTenantScope != model.TenantScope {
		return Point15Val0StateBlocked
	}
	return Point15Val0StateActive
}

func point15Val0TimestampDisciplineModel(dependency Point15Val0DependencySnapshot) Point15Val0TimestampDiscipline {
	return Point15Val0TimestampDiscipline{
		DisciplineID:               "timestamp_discipline_point15_val0_001",
		TenantScope:                dependency.InheritedTenantScope,
		FreshnessStatus:            point15Val0FreshnessFresh,
		ObservedAt:                 "2026-05-06T18:15:00Z",
		ObservedTimeSource:         point14Val0TimeSourceServerUTC,
		EvaluatedAt:                "2026-05-06T18:20:00Z",
		EvaluatedTimeSource:        point14Val0TimeSourceServerUTC,
		ValidatedAt:                "2026-05-06T18:25:00Z",
		ValidatedTimeSource:        point14Val0TimeSourceServerUTC,
		ReviewerApprovedAt:         "2026-05-06T18:27:00Z",
		ReviewerApprovedTimeSource: point14Val0TimeSourceServerUTC,
		SourceEventAt:              "2026-05-06T18:10:00Z",
		SourceEventTimeSource:      point14Val0TimeSourceApprovedCustomerTime,
		ReferenceNow:               "2026-05-06T18:30:00Z",
		ReferenceNowTimeSource:     point14Val0TimeSourceServerUTC,
	}
}

func EvaluatePoint15Val0TimestampDisciplineState(model Point15Val0TimestampDiscipline) string {
	if !point15Val0DependencyRefValid(model.DisciplineID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point15Val0FreshnessStatusValid(model.FreshnessStatus) ||
		!point14Val0ParsedTimeOk(model.ObservedAt) ||
		!point14Val0ParsedTimeOk(model.ReferenceNow) ||
		!point14Val0CanonicalTimeSourceValid(model.ObservedTimeSource) ||
		!point14Val0CanonicalTimeSourceValid(model.ReferenceNowTimeSource) {
		return Point15Val0StateBlocked
	}
	if model.EvaluatedAt == "" ||
		model.EvaluatedTimeSource == "" ||
		model.ValidatedAt == "" ||
		model.ValidatedTimeSource == "" {
		return Point15Val0StateIncomplete
	}
	if !point14Val0ParsedTimeOk(model.EvaluatedAt) ||
		!point14Val0ParsedTimeOk(model.ValidatedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.EvaluatedTimeSource) ||
		!point14Val0CanonicalTimeSourceValid(model.ValidatedTimeSource) {
		return Point15Val0StateBlocked
	}
	if model.ReviewerApprovedAt != "" {
		if !point14Val0ParsedTimeOk(model.ReviewerApprovedAt) || !point14Val0CanonicalTimeSourceValid(model.ReviewerApprovedTimeSource) {
			return Point15Val0StateBlocked
		}
	}
	if model.SourceEventAt != "" {
		if !point14Val0ParsedTimeOk(model.SourceEventAt) || !point14Val0TimeSourceValid(model.SourceEventTimeSource) {
			return Point15Val0StateBlocked
		}
	}
	if model.ClientLocalCreatesCanonical || model.SourceEventCreatesCanonical {
		return Point15Val0StateBlocked
	}
	observedAt, _ := point14Val0ParsedTime(model.ObservedAt)
	evaluatedAt, _ := point14Val0ParsedTime(model.EvaluatedAt)
	validatedAt, _ := point14Val0ParsedTime(model.ValidatedAt)
	referenceNow, _ := point14Val0ParsedTime(model.ReferenceNow)
	if observedAt.After(evaluatedAt) || evaluatedAt.After(validatedAt) {
		return Point15Val0StateBlocked
	}
	if observedAt.After(referenceNow) || evaluatedAt.After(referenceNow) || validatedAt.After(referenceNow) {
		return Point15Val0StateBlocked
	}
	if model.ReviewerApprovedAt != "" {
		reviewerApprovedAt, _ := point14Val0ParsedTime(model.ReviewerApprovedAt)
		if reviewerApprovedAt.Before(validatedAt) {
			return Point15Val0StateReviewRequired
		}
		if reviewerApprovedAt.After(referenceNow) {
			return Point15Val0StateBlocked
		}
	}
	return Point15Val0StateActive
}

func point15Val0AuthorityBoundaryModel(dependency Point15Val0DependencySnapshot) Point15Val0AuthorityBoundary {
	return Point15Val0AuthorityBoundary{
		BoundaryID:                      "authority_boundary_point15_val0_001",
		TenantScope:                     dependency.InheritedTenantScope,
		ExternalSourceInputOnly:         true,
		AgentRecommendationAdvisoryOnly: true,
	}
}

func EvaluatePoint15Val0AuthorityBoundaryState(model Point15Val0AuthorityBoundary) string {
	if !point15Val0DependencyRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!model.ExternalSourceInputOnly ||
		!model.AgentRecommendationAdvisoryOnly {
		return Point15Val0StateBlocked
	}
	if model.SchedulerPassAllowed ||
		model.FreshnessAuthorityAllowed ||
		model.DashboardFreshnessAllowed ||
		model.AgentFreshnessAllowed ||
		model.ConnectorFreshnessAuthorityAllowed ||
		model.CustomerProjectionMutatesFreshness ||
		model.AuditorProjectionMutatesFreshness ||
		model.PortalProjectionMutatesFreshness ||
		model.CanonicalMutationAllowed ||
		model.ProductionMutationAllowed ||
		model.PassAllowed {
		return Point15Val0StateBlocked
	}
	return Point15Val0StateActive
}

func point15Val0NoOverclaimGuardModel() Point15Val0NoOverclaimGuard {
	return Point15Val0NoOverclaimGuard{
		ObservedTexts: []string{
			"freshness status bounded by evidence context",
			"point 15 val 0 verifies freshness discipline only",
		},
		AllowedSafeWording:  point15Val0SafeWording(),
		BlockedWording:      point15Val0ForbiddenWording(),
		FreshnessDisclaimer: point15Val0FreshnessDisclaimer,
	}
}

func EvaluatePoint15Val0NoOverclaimGuardState(model Point15Val0NoOverclaimGuard) string {
	if model.FreshnessDisclaimer != point15Val0FreshnessDisclaimer ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point15Val0SafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point15Val0ForbiddenWording()) {
		return Point15Val0StateBlocked
	}
	if point15Val0ObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point15Val0StateBlocked
	}
	if point15Val0ObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point15Val0StateBlocked
	}
	return Point15Val0StateActive
}

func point15Val0FoundationModelFromUpstream(valE Point14ValEFoundation) Point15Val0FreshnessDisciplineFoundation {
	dependency := point15Val0DependencySnapshotFromUpstream(valE)
	return Point15Val0FreshnessDisciplineFoundation{
		FreshnessDisclaimer: point15Val0FreshnessDisclaimer,
		Dependency:          dependency,
		FreshnessTaxonomy:   point15Val0FreshnessTaxonomyModel(),
		DowngradeTaxonomy:   point15Val0DowngradeTaxonomyModel(),
		EvidenceContext:     point15Val0FreshnessEvidenceContextModel(dependency),
		TimestampDiscipline: point15Val0TimestampDisciplineModel(dependency),
		AuthorityBoundary:   point15Val0AuthorityBoundaryModel(dependency),
		NoOverclaimGuard:    point15Val0NoOverclaimGuardModel(),
	}
}

func Point15Val0FoundationModel() Point15Val0FreshnessDisciplineFoundation {
	valE := ComputePoint14ValEFoundation(Point14ValEFoundationModel())
	return point15Val0FoundationModelFromUpstream(valE)
}

func point15Val0Aggregate(states ...string) string {
	for _, state := range states {
		if !point15Val0StateValid(state) {
			return Point15Val0StateBlocked
		}
		if state == Point15Val0StateBlocked {
			return Point15Val0StateBlocked
		}
	}
	for _, state := range states {
		if state == Point15Val0StateReviewRequired {
			return Point15Val0StateReviewRequired
		}
	}
	for _, state := range states {
		if state == Point15Val0StateIncomplete {
			return Point15Val0StateIncomplete
		}
	}
	return Point15Val0StateActive
}

func point15Val0BlockingReasons(model Point15Val0FreshnessDisciplineFoundation) []string {
	componentStates := map[string]string{
		"dependency":           model.DependencyState,
		"freshness_taxonomy":   model.FreshnessTaxonomyState,
		"downgrade_taxonomy":   model.DowngradeTaxonomyState,
		"evidence_context":     model.EvidenceContextState,
		"tenant_boundary":      model.TenantBoundaryState,
		"timestamp_discipline": model.TimestampDisciplineState,
		"authority_boundary":   model.AuthorityBoundaryState,
		"no_overclaim":         model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if state == Point15Val0StateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point15Val0ReviewPrerequisites(model Point15Val0FreshnessDisciplineFoundation) []string {
	componentStates := map[string]string{
		"freshness_taxonomy":   model.FreshnessTaxonomyState,
		"downgrade_taxonomy":   model.DowngradeTaxonomyState,
		"evidence_context":     model.EvidenceContextState,
		"tenant_boundary":      model.TenantBoundaryState,
		"timestamp_discipline": model.TimestampDisciplineState,
		"authority_boundary":   model.AuthorityBoundaryState,
		"no_overclaim":         model.NoOverclaimState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if state == Point15Val0StateReviewRequired || state == Point15Val0StateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint15Val0FreshnessDisciplineFoundation(model Point15Val0FreshnessDisciplineFoundation) Point15Val0FreshnessDisciplineFoundation {
	model.DependencyState = EvaluatePoint15Val0DependencyState(model.Dependency)
	model.FreshnessTaxonomyState = EvaluatePoint15Val0EvidenceFreshnessTaxonomyState(model.FreshnessTaxonomy)
	model.DowngradeTaxonomyState = EvaluatePoint15Val0DowngradeTaxonomyState(model.DowngradeTaxonomy)
	model.EvidenceContextState = EvaluatePoint15Val0FreshnessEvidenceContextState(model.EvidenceContext)
	model.TenantBoundaryState = EvaluatePoint15Val0TenantBoundaryState(model.EvidenceContext)
	model.TimestampDisciplineState = EvaluatePoint15Val0TimestampDisciplineState(model.TimestampDiscipline)
	model.AuthorityBoundaryState = EvaluatePoint15Val0AuthorityBoundaryState(model.AuthorityBoundary)
	model.NoOverclaimState = EvaluatePoint15Val0NoOverclaimGuardState(model.NoOverclaimGuard)
	if model.FreshnessDisclaimer != point15Val0FreshnessDisclaimer {
		model.NoOverclaimState = Point15Val0StateBlocked
	}
	expectedTenantScope := model.Dependency.InheritedTenantScope
	if model.EvidenceContext.FreshnessStatus != model.FreshnessTaxonomy.FreshnessStatus {
		model.EvidenceContextState = Point15Val0StateBlocked
	}
	if model.DowngradeTaxonomy.FreshnessStatus != model.FreshnessTaxonomy.FreshnessStatus {
		model.DowngradeTaxonomyState = Point15Val0StateBlocked
	}
	if model.TimestampDiscipline.FreshnessStatus != model.FreshnessTaxonomy.FreshnessStatus {
		model.TimestampDisciplineState = Point15Val0StateBlocked
	}
	if model.EvidenceContext.DowngradeOutcome != model.DowngradeTaxonomy.DowngradeOutcome {
		model.EvidenceContextState = Point15Val0StateBlocked
	}
	if expectedTenantScope == "" ||
		model.EvidenceContext.TenantScope != expectedTenantScope ||
		(model.EvidenceContext.ReferencedTenantScope != "" && model.EvidenceContext.ReferencedTenantScope != expectedTenantScope) {
		model.EvidenceContextState = Point15Val0StateBlocked
		model.TenantBoundaryState = Point15Val0StateBlocked
	}
	if expectedTenantScope == "" || model.TimestampDiscipline.TenantScope != expectedTenantScope {
		model.TimestampDisciplineState = Point15Val0StateBlocked
	}
	if expectedTenantScope == "" || model.AuthorityBoundary.TenantScope != expectedTenantScope {
		model.AuthorityBoundaryState = Point15Val0StateBlocked
	}
	model.CurrentState = point15Val0Aggregate(
		model.DependencyState,
		model.FreshnessTaxonomyState,
		model.DowngradeTaxonomyState,
		model.EvidenceContextState,
		model.TenantBoundaryState,
		model.TimestampDisciplineState,
		model.AuthorityBoundaryState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point15Val0BlockingReasons(model)
	model.ReviewPrerequisites = point15Val0ReviewPrerequisites(model)
	return model
}
