package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	referenceArchitectureValDVisibilitySchema    = "point6.reference_architecture.vald.operational_visibility.v1"
	referenceArchitectureValDAlignmentSchema     = "point6.reference_architecture.vald.alignment_summary.v1"
	referenceArchitectureValDDeviationSchema     = "point6.reference_architecture.vald.deviation_alerts.v1"
	referenceArchitectureValDSupportSchema       = "point6.reference_architecture.vald.support_boundaries.v1"
	referenceArchitectureValDMigrationSchema     = "point6.reference_architecture.vald.migration_upgrade.v1"
	referenceArchitectureValDTopologySchema      = "point6.reference_architecture.vald.topology_gate.v1"
	referenceArchitectureValDSecuritySchema      = "point6.reference_architecture.vald.security_boundary_gate.v1"
	referenceArchitectureValDOperabilitySchema   = "point6.reference_architecture.vald.operability_gate.v1"
	referenceArchitectureValDCompatibilitySchema = "point6.reference_architecture.vald.compatibility_gate.v1"
	referenceArchitectureValDFinalGateSchema     = "point6.reference_architecture.vald.final_gate.v1"
	referenceArchitectureValDProofsSchema        = "point6.reference_architecture.vald.proofs.v1"
)

type referenceArchitectureValDFamilyStatus struct {
	Family                    string `json:"family"`
	VisibilityReportID        string `json:"visibility_report_id"`
	VisibilityState           string `json:"operational_visibility_state"`
	AlignmentSummaryState     string `json:"alignment_summary_state"`
	DeviationAlertState       string `json:"deviation_alert_state"`
	SupportBoundaryState      string `json:"support_boundary_state"`
	MigrationUpgradeState     string `json:"migration_upgrade_state"`
	TopologyGateState         string `json:"topology_gate_state"`
	SecurityBoundaryGateState string `json:"security_boundary_gate_state"`
	OperabilityGateState      string `json:"operability_gate_state"`
	CompatibilityGateState    string `json:"compatibility_gate_state"`
	FinalGateState            string `json:"final_gate_state"`
	BlockingAlertCount        int    `json:"blocking_alert_count"`
	SupportGapCount           int    `json:"support_gap_count"`
}

type referenceArchitectureValDVisibilityResponse struct {
	SchemaVersion string                                                           `json:"schema_version"`
	GeneratedAt   time.Time                                                        `json:"generated_at"`
	CurrentState  string                                                           `json:"current_state"`
	Model         operability.ReferenceArchitectureOperationalVisibilityCollection `json:"model"`
	FamilyStates  []referenceArchitectureValDFamilyStatus                          `json:"family_states,omitempty"`
	RouteRefs     []string                                                         `json:"route_refs,omitempty"`
	Limitations   []string                                                         `json:"limitations,omitempty"`
}

type referenceArchitectureValDCollectionResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	FamilyStates  []referenceArchitectureValDFamilyStatus `json:"family_states,omitempty"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
	Model         any                                     `json:"model"`
}

type referenceArchitectureValDProofsResponse struct {
	SchemaVersion              string                                  `json:"schema_version"`
	GeneratedAt                time.Time                               `json:"generated_at"`
	CurrentState               string                                  `json:"current_state"`
	Point5DependencyState      string                                  `json:"point_5_dependency_state"`
	Point5State                string                                  `json:"point_5_state"`
	Val0DependencyState        string                                  `json:"val_0_dependency_state"`
	Val0State                  string                                  `json:"val_0_state"`
	ValADependencyState        string                                  `json:"val_a_dependency_state"`
	ValAState                  string                                  `json:"val_a_state"`
	ValBDependencyState        string                                  `json:"val_b_dependency_state"`
	ValBState                  string                                  `json:"val_b_state"`
	ValCDependencyState        string                                  `json:"val_c_dependency_state"`
	ValCState                  string                                  `json:"val_c_state"`
	ValDState                  string                                  `json:"val_d_state"`
	Point6State                string                                  `json:"point_6_state"`
	OperationalVisibilityState string                                  `json:"operational_visibility_state"`
	AlignmentSummaryState      string                                  `json:"alignment_summary_state"`
	DeviationAlertState        string                                  `json:"deviation_alert_state"`
	SupportBoundaryState       string                                  `json:"support_boundary_state"`
	MigrationUpgradeState      string                                  `json:"migration_upgrade_state"`
	TopologyGateState          string                                  `json:"topology_gate_state"`
	SecurityBoundaryGateState  string                                  `json:"security_boundary_gate_state"`
	OperabilityGateState       string                                  `json:"operability_gate_state"`
	CompatibilityGateState     string                                  `json:"compatibility_gate_state"`
	FinalGateState             string                                  `json:"final_gate_state"`
	SupportedFamilies          []string                                `json:"supported_blueprint_families,omitempty"`
	FamilyStates               []referenceArchitectureValDFamilyStatus `json:"family_states,omitempty"`
	WhyPoint6NotPass           []string                                `json:"why_point_6_not_pass,omitempty"`
	SurfaceRefs                []string                                `json:"surface_refs,omitempty"`
	EvidenceRefs               []string                                `json:"evidence_refs,omitempty"`
	Limitations                []string                                `json:"limitations,omitempty"`
	ProjectionDisclaimer       string                                  `json:"projection_disclaimer"`
	IntegrationSummary         []string                                `json:"integration_summary,omitempty"`
}

func referenceArchitectureValDAllSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/proofs",
		"/v1/reference-architecture/vald/operational-visibility",
		"/v1/reference-architecture/vald/alignment-summary",
		"/v1/reference-architecture/vald/deviation-alerts",
		"/v1/reference-architecture/vald/support-boundaries",
		"/v1/reference-architecture/vald/migration-upgrade",
		"/v1/reference-architecture/vald/topology-gate",
		"/v1/reference-architecture/vald/security-boundary-gate",
		"/v1/reference-architecture/vald/operability-gate",
		"/v1/reference-architecture/vald/compatibility-gate",
		"/v1/reference-architecture/vald/final-gate",
		"/v1/reference-architecture/vald/proofs",
	}
}

func referenceArchitectureValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_operational_visibility_final_reference_gate"
}

func referenceArchitectureValDAppendUnique(values []string, value string) []string {
	if value == "" {
		return values
	}
	for _, existing := range values {
		if existing == value {
			return values
		}
	}
	return append(values, value)
}

func referenceArchitectureValDCountBlockingAlerts(report operability.ReferenceArchitectureDeviationAlertReport) int {
	count := 0
	for _, alert := range report.Alerts {
		if alert.BlocksAlignment {
			count++
		}
	}
	return count
}

func buildReferenceArchitectureValDFamilyStatuses(
	visibility operability.ReferenceArchitectureOperationalVisibilityCollection,
	alignment operability.ReferenceArchitectureBlueprintAlignmentCollection,
	alerts operability.ReferenceArchitectureDeviationAlertCollection,
	support operability.ReferenceArchitectureSupportBoundaryCollection,
	migration operability.ReferenceArchitectureMigrationUpgradeCollection,
	topology operability.ReferenceArchitectureTopologyGateCollection,
	security operability.ReferenceArchitectureSecurityBoundaryCollection,
	operabilityCollection operability.ReferenceArchitectureOperabilityGateCollection,
	compatibility operability.ReferenceArchitectureCompatibilityGateCollection,
	finalGate operability.ReferenceArchitectureFinalGateCollection,
) []referenceArchitectureValDFamilyStatus {
	visibilityByFamily := map[string]operability.ReferenceArchitectureOperationalVisibilityReport{}
	for _, report := range visibility.Reports {
		visibilityByFamily[report.BlueprintFamily] = report
	}
	alignmentByFamily := map[string]operability.ReferenceArchitectureBlueprintAlignmentSummary{}
	for _, summary := range alignment.Summaries {
		alignmentByFamily[summary.BlueprintFamily] = summary
	}
	alertsByFamily := map[string]operability.ReferenceArchitectureDeviationAlertReport{}
	for _, report := range alerts.Reports {
		alertsByFamily[report.BlueprintFamily] = report
	}
	supportByFamily := map[string]operability.ReferenceArchitectureSupportBoundaryView{}
	for _, view := range support.Views {
		supportByFamily[view.BlueprintFamily] = view
	}
	migrationByFamily := map[string]operability.ReferenceArchitectureMigrationUpgradeVisibility{}
	for _, view := range migration.Views {
		migrationByFamily[view.BlueprintFamily] = view
	}
	topologyByFamily := map[string]operability.ReferenceArchitectureTopologyGateCheck{}
	for _, check := range topology.Checks {
		topologyByFamily[check.BlueprintFamily] = check
	}
	securityByFamily := map[string]operability.ReferenceArchitectureSecurityBoundaryGateCheck{}
	for _, check := range security.Checks {
		securityByFamily[check.BlueprintFamily] = check
	}
	operabilityByFamily := map[string]operability.ReferenceArchitectureOperabilityGateCheck{}
	for _, check := range operabilityCollection.Checks {
		operabilityByFamily[check.BlueprintFamily] = check
	}
	compatibilityByFamily := map[string]operability.ReferenceArchitectureCompatibilityGateCheck{}
	for _, check := range compatibility.Checks {
		compatibilityByFamily[check.BlueprintFamily] = check
	}
	finalByFamily := map[string]operability.ReferenceArchitectureFinalGateReport{}
	for _, report := range finalGate.Reports {
		finalByFamily[report.BlueprintFamily] = report
	}

	statuses := make([]referenceArchitectureValDFamilyStatus, 0, len(visibility.SupportedFamilies))
	for _, family := range visibility.SupportedFamilies {
		visibilityReport := visibilityByFamily[family]
		alignmentSummary := alignmentByFamily[family]
		alertReport := alertsByFamily[family]
		supportView := supportByFamily[family]
		migrationView := migrationByFamily[family]
		topologyCheck := topologyByFamily[family]
		securityCheck := securityByFamily[family]
		operabilityCheck := operabilityByFamily[family]
		compatibilityCheck := compatibilityByFamily[family]
		finalReport := finalByFamily[family]

		statuses = append(statuses, referenceArchitectureValDFamilyStatus{
			Family:                    family,
			VisibilityReportID:        visibilityReport.VisibilityReportID,
			VisibilityState:           operability.EvaluateReferenceArchitectureValDOperationalVisibilityReportState(visibilityReport),
			AlignmentSummaryState:     operability.EvaluateReferenceArchitectureValDAlignmentSummaryState(alignmentSummary),
			DeviationAlertState:       operability.EvaluateReferenceArchitectureValDDeviationAlertReportState(alertReport),
			SupportBoundaryState:      operability.EvaluateReferenceArchitectureValDSupportBoundaryViewState(supportView),
			MigrationUpgradeState:     operability.EvaluateReferenceArchitectureValDMigrationVisibilityState(migrationView),
			TopologyGateState:         operability.EvaluateReferenceArchitectureValDTopologyGateState(topologyCheck),
			SecurityBoundaryGateState: operability.EvaluateReferenceArchitectureValDSecurityBoundaryGateState(securityCheck),
			OperabilityGateState:      operability.EvaluateReferenceArchitectureValDOperabilityGateState(operabilityCheck),
			CompatibilityGateState:    operability.EvaluateReferenceArchitectureValDCompatibilityGateState(compatibilityCheck),
			FinalGateState:            operability.EvaluateReferenceArchitectureValDFinalGateReportState(finalReport),
			BlockingAlertCount:        referenceArchitectureValDCountBlockingAlerts(alertReport),
			SupportGapCount:           len(supportView.UnsupportedConditions),
		})
	}
	return statuses
}

func referenceArchitectureValDEvidenceRefs(
	visibility operability.ReferenceArchitectureOperationalVisibilityCollection,
	migration operability.ReferenceArchitectureMigrationUpgradeCollection,
	topology operability.ReferenceArchitectureTopologyGateCollection,
	security operability.ReferenceArchitectureSecurityBoundaryCollection,
	operabilityCollection operability.ReferenceArchitectureOperabilityGateCollection,
	compatibility operability.ReferenceArchitectureCompatibilityGateCollection,
	finalGate operability.ReferenceArchitectureFinalGateCollection,
) []string {
	refs := []string{
		"point5_integrated_closure",
		"point6_val0_proofs",
		"point6_vala_proofs",
		"point6_valb_proofs",
		"point6_valc_proofs",
		visibility.CollectionID,
		migration.CollectionID,
		topology.CollectionID,
		security.CollectionID,
		operabilityCollection.CollectionID,
		compatibility.CollectionID,
		finalGate.CollectionID,
	}
	for _, report := range visibility.Reports {
		refs = referenceArchitectureValDAppendUnique(refs, report.VisibilityReportID)
		for _, evidence := range report.EvidenceRefs {
			refs = referenceArchitectureValDAppendUnique(refs, evidence.EvidenceID)
		}
	}
	for _, view := range migration.Views {
		refs = referenceArchitectureValDAppendUnique(refs, view.VisibilityID)
		for _, evidence := range view.EvidenceRefs {
			refs = referenceArchitectureValDAppendUnique(refs, evidence.EvidenceID)
		}
	}
	for _, check := range topology.Checks {
		refs = referenceArchitectureValDAppendUnique(refs, check.CheckID)
		for _, evidence := range check.EvidenceRefs {
			refs = referenceArchitectureValDAppendUnique(refs, evidence.EvidenceID)
		}
	}
	for _, check := range security.Checks {
		refs = referenceArchitectureValDAppendUnique(refs, check.CheckID)
		for _, evidence := range check.EvidenceRefs {
			refs = referenceArchitectureValDAppendUnique(refs, evidence.EvidenceID)
		}
	}
	for _, check := range operabilityCollection.Checks {
		refs = referenceArchitectureValDAppendUnique(refs, check.CheckID)
		for _, evidence := range check.EvidenceRefs {
			refs = referenceArchitectureValDAppendUnique(refs, evidence.EvidenceID)
		}
	}
	for _, check := range compatibility.Checks {
		refs = referenceArchitectureValDAppendUnique(refs, check.CheckID)
		for _, evidence := range check.EvidenceRefs {
			refs = referenceArchitectureValDAppendUnique(refs, evidence.EvidenceID)
		}
	}
	for _, report := range finalGate.Reports {
		refs = referenceArchitectureValDAppendUnique(refs, report.GateID)
	}
	return refs
}

func (s server) referenceArchitectureValDOperationalVisibilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDOperationalVisibility())
}

func (s server) referenceArchitectureValDAlignmentSummaryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDAlignmentSummary())
}

func (s server) referenceArchitectureValDDeviationAlertsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDDeviationAlerts())
}

func (s server) referenceArchitectureValDSupportBoundariesHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDSupportBoundaries())
}

func (s server) referenceArchitectureValDMigrationUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDMigrationUpgrade())
}

func (s server) referenceArchitectureValDTopologyGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDTopologyGate())
}

func (s server) referenceArchitectureValDSecurityBoundaryGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDSecurityBoundaryGate())
}

func (s server) referenceArchitectureValDOperabilityGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDOperabilityGate())
}

func (s server) referenceArchitectureValDCompatibilityGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDCompatibilityGate())
}

func (s server) referenceArchitectureValDFinalGateHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDFinalGate())
}

func (s server) referenceArchitectureValDProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildReferenceArchitectureValDProofs())
}

func buildReferenceArchitectureValDSharedCollections() (
	operability.ReferenceArchitectureOperationalVisibilityCollection,
	operability.ReferenceArchitectureBlueprintAlignmentCollection,
	operability.ReferenceArchitectureDeviationAlertCollection,
	operability.ReferenceArchitectureSupportBoundaryCollection,
	operability.ReferenceArchitectureMigrationUpgradeCollection,
	operability.ReferenceArchitectureTopologyGateCollection,
	operability.ReferenceArchitectureSecurityBoundaryCollection,
	operability.ReferenceArchitectureOperabilityGateCollection,
	operability.ReferenceArchitectureCompatibilityGateCollection,
	operability.ReferenceArchitectureFinalGateCollection,
) {
	visibility := operability.ReferenceArchitectureValDOperationalVisibilityCollection()
	alignment := operability.ReferenceArchitectureValDAlignmentSummaryCollection()
	alerts := operability.ReferenceArchitectureValDDeviationAlertCollection()
	support := operability.ReferenceArchitectureValDSupportBoundaryCollection()
	migration := operability.ReferenceArchitectureValDMigrationUpgradeCollection()
	topology := operability.ReferenceArchitectureValDTopologyGateCollection()
	security := operability.ReferenceArchitectureValDSecurityBoundaryCollection()
	operabilityCollection := operability.ReferenceArchitectureValDOperabilityGateCollection()
	compatibility := operability.ReferenceArchitectureValDCompatibilityGateCollection()
	finalGate := operability.ReferenceArchitectureValDFinalGateCollectionFromComponents(
		visibility,
		alignment,
		alerts,
		support,
		migration,
		topology,
		security,
		operabilityCollection,
		compatibility,
	)
	return visibility,
		alignment,
		alerts,
		support,
		migration,
		topology,
		security,
		operabilityCollection,
		compatibility,
		finalGate
}

func buildReferenceArchitectureValDOperationalVisibility() referenceArchitectureValDVisibilityResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDVisibilityResponse{
		SchemaVersion: referenceArchitectureValDVisibilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDOperationalVisibilityCollectionState(visibility),
		Model:         visibility,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/valc/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Operational visibility remains a dashboard-ready advisory projection over the canonical evidence spine.",
			"Val D surfaces do not approve deployment, mutate canonical truth, or close Točka 6.",
		},
	}
}

func buildReferenceArchitectureValDAlignmentSummary() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDAlignmentSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDAlignmentSummaryCollectionState(alignment),
		Model:         alignment,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/val0/proofs",
			"/v1/reference-architecture/vala/proofs",
			"/v1/reference-architecture/valb/proofs",
			"/v1/reference-architecture/valc/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Alignment summary remains evidence-linked and fail-closed; degraded, stale, blocked, and unsupported states remain visible.",
			"Summary language is bounded and does not imply certification or guaranteed security.",
		},
	}
}

func buildReferenceArchitectureValDDeviationAlerts() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDDeviationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDDeviationAlertCollectionState(alerts),
		Model:         alerts,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/proofs",
			"/v1/reference-architecture/valc/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Deviation alerts remain advisory and do not suppress, mutate, or approve anything.",
			"Blocking alert evidence must remain visible and fresh to preserve fail-closed semantics.",
		},
	}
}

func buildReferenceArchitectureValDSupportBoundaries() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDSupportSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDSupportBoundaryCollectionState(support),
		Model:         support,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/vala/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Support boundary views preserve unsupported and degraded conditions and do not create canonical authority for partners or MSPs.",
			"Support boundaries remain bounded visibility artifacts, not approval or mutation surfaces.",
		},
	}
}

func buildReferenceArchitectureValDMigrationUpgrade() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDMigrationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDMigrationUpgradeCollectionState(migration),
		Model:         migration,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Migration and upgrade visibility does not execute migration, rollback, or upgrade workflows.",
			"Deprecated and superseded visibility remains explicit and cannot silently become clean matched state.",
		},
	}
}

func buildReferenceArchitectureValDTopologyGate() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDTopologySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDTopologyGateCollectionState(topology),
		Model:         topology,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/val0/proofs",
			"/v1/reference-architecture/valc/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Topology gate output is evidence-linked and fail-closed; unsupported or mismatched topology cannot be redacted into matched.",
			"Offline and sovereign topology compatibility remains bounded and advisory.",
		},
	}
}

func buildReferenceArchitectureValDSecurityBoundaryGate() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDSecuritySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDSecurityBoundaryCollectionState(security),
		Model:         security,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/vala/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Security boundary gate preserves no-shadow-truth, no-approval-authority, and no-mutation-authority semantics.",
			"Boundary views remain advisory projections and do not become policy or approval authority.",
		},
	}
}

func buildReferenceArchitectureValDOperabilityGate() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDOperabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDOperabilityGateCollectionState(operabilityCollection),
		Model:         operabilityCollection,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/valb/proofs",
			"/v1/reference-architecture/valc/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Operability gate reuses readiness and resilience evidence; it does not execute remediation or approve production use.",
			"Degraded state remains visible and operator action guidance is mandatory for clean alignment.",
		},
	}
}

func buildReferenceArchitectureValDCompatibilityGate() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDCompatibilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDCompatibilityGateCollectionState(compatibility),
		Model:         compatibility,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/val0/proofs",
			"/v1/reference-architecture/valb/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Compatibility gate remains bounded and does not claim universal support or deployment approval.",
			"Deprecated, superseded, stale, and unsupported compatibility paths remain explicit non-final states.",
		},
	}
}

func buildReferenceArchitectureValDFinalGate() referenceArchitectureValDCollectionResponse {
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()
	return buildReferenceArchitectureValDFinalGateResponse(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate)
}

func buildReferenceArchitectureValDFinalGateResponse(
	visibility operability.ReferenceArchitectureOperationalVisibilityCollection,
	alignment operability.ReferenceArchitectureBlueprintAlignmentCollection,
	alerts operability.ReferenceArchitectureDeviationAlertCollection,
	support operability.ReferenceArchitectureSupportBoundaryCollection,
	migration operability.ReferenceArchitectureMigrationUpgradeCollection,
	topology operability.ReferenceArchitectureTopologyGateCollection,
	security operability.ReferenceArchitectureSecurityBoundaryCollection,
	operabilityCollection operability.ReferenceArchitectureOperabilityGateCollection,
	compatibility operability.ReferenceArchitectureCompatibilityGateCollection,
	finalGate operability.ReferenceArchitectureFinalGateCollection,
) referenceArchitectureValDCollectionResponse {
	return referenceArchitectureValDCollectionResponse{
		SchemaVersion: referenceArchitectureValDFinalGateSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateReferenceArchitectureValDFinalGateCollectionState(finalGate),
		Model:         finalGate,
		FamilyStates:  buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		RouteRefs: []string{
			"/v1/reference-architecture/val0/proofs",
			"/v1/reference-architecture/vala/proofs",
			"/v1/reference-architecture/valb/proofs",
			"/v1/reference-architecture/valc/proofs",
			"/v1/reference-architecture/vald/proofs",
		},
		Limitations: []string{
			"Val D final reference architecture gate is point-specific only and does not close Točka 6.",
			"Val E integrated closure remains the only place where point_6_pass can be decided.",
		},
	}
}

func buildReferenceArchitectureValDProofs() referenceArchitectureValDProofsResponse {
	val0 := buildReferenceArchitectureVal0Proofs()
	valA := buildReferenceArchitectureValAProofs()
	valB := buildReferenceArchitectureValBProofs()
	valC := buildReferenceArchitectureValCProofs()
	visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate := buildReferenceArchitectureValDSharedCollections()

	visibilityState := operability.EvaluateReferenceArchitectureValDOperationalVisibilityCollectionState(visibility)
	alignmentState := operability.EvaluateReferenceArchitectureValDAlignmentSummaryCollectionState(alignment)
	alertState := operability.EvaluateReferenceArchitectureValDDeviationAlertCollectionState(alerts)
	supportState := operability.EvaluateReferenceArchitectureValDSupportBoundaryCollectionState(support)
	migrationState := operability.EvaluateReferenceArchitectureValDMigrationUpgradeCollectionState(migration)
	topologyState := operability.EvaluateReferenceArchitectureValDTopologyGateCollectionState(topology)
	securityState := operability.EvaluateReferenceArchitectureValDSecurityBoundaryCollectionState(security)
	operabilityState := operability.EvaluateReferenceArchitectureValDOperabilityGateCollectionState(operabilityCollection)
	compatibilityState := operability.EvaluateReferenceArchitectureValDCompatibilityGateCollectionState(compatibility)
	finalGateState := operability.EvaluateReferenceArchitectureValDFinalGateCollectionState(finalGate)

	valDState := operability.EvaluateReferenceArchitectureValDState(
		val0.Point5State,
		val0.Point5DependencyState,
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		valB.CurrentState,
		valB.ValBState,
		valC.CurrentState,
		valC.ValCState,
		val0.Point6State,
		visibilityState,
		alignmentState,
		alertState,
		supportState,
		migrationState,
		topologyState,
		securityState,
		operabilityState,
		compatibilityState,
		finalGateState,
	)

	surfaceRefs := referenceArchitectureValDAllSurfaceRefs()
	evidenceRefs := referenceArchitectureValDEvidenceRefs(visibility, migration, topology, security, operabilityCollection, compatibility, finalGate)
	limitations := []string{
		"Val D provides dashboard-ready operational visibility and a point-specific final reference architecture gate only.",
		"Val D does not approve deployment, execute migration or rollback, or close Točka 6.",
		"Deviation alerts remain advisory and cannot suppress, mutate, or approve anything.",
		"Točka 6 remains not_complete until integrated closure in Val E.",
	}
	currentState := operability.EvaluateReferenceArchitectureValDProofsState(
		valDState,
		val0.Point5DependencyState,
		val0.Point6State,
		visibility.SupportedFamilies,
		surfaceRefs,
		evidenceRefs,
		limitations,
		referenceArchitectureValDProjectionDisclaimer(),
	)

	return referenceArchitectureValDProofsResponse{
		SchemaVersion:              referenceArchitectureValDProofsSchema,
		GeneratedAt:                publicSampleTime(),
		CurrentState:               currentState,
		Point5DependencyState:      val0.Point5DependencyState,
		Point5State:                val0.Point5State,
		Val0DependencyState:        val0.CurrentState,
		Val0State:                  val0.Val0State,
		ValADependencyState:        valA.CurrentState,
		ValAState:                  valA.ValAState,
		ValBDependencyState:        valB.CurrentState,
		ValBState:                  valB.ValBState,
		ValCDependencyState:        valC.CurrentState,
		ValCState:                  valC.ValCState,
		ValDState:                  valDState,
		Point6State:                operability.ReferenceArchitecturePoint6StateNotComplete,
		OperationalVisibilityState: visibilityState,
		AlignmentSummaryState:      alignmentState,
		DeviationAlertState:        alertState,
		SupportBoundaryState:       supportState,
		MigrationUpgradeState:      migrationState,
		TopologyGateState:          topologyState,
		SecurityBoundaryGateState:  securityState,
		OperabilityGateState:       operabilityState,
		CompatibilityGateState:     compatibilityState,
		FinalGateState:             finalGateState,
		SupportedFamilies:          visibility.SupportedFamilies,
		FamilyStates:               buildReferenceArchitectureValDFamilyStatuses(visibility, alignment, alerts, support, migration, topology, security, operabilityCollection, compatibility, finalGate),
		WhyPoint6NotPass: []string{
			"Val D provides operational visibility and a point-specific final reference architecture gate only.",
			"Val E integrated closure remains required before any final point_6_pass decision.",
		},
		SurfaceRefs:          surfaceRefs,
		EvidenceRefs:         evidenceRefs,
		Limitations:          limitations,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val D binds Val 0 through Val C into evidence-linked operational visibility, support boundary views, migration visibility, and final point-specific gate reporting.",
			"These outputs remain advisory projections and do not become deployment approval authority, canonical truth, or integrated closure.",
		},
	}
}
