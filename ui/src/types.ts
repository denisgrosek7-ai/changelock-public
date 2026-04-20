export type Decision = "ALLOW" | "DENY" | "ERROR";
export type ExceptionStatus = "PENDING" | "APPROVED" | "REJECTED" | "REVOKED" | "EXPIRED";
export type VulnerabilityDecisionValue = "NOT_AFFECTED" | "ACCEPTED_RISK" | "FIX_REQUIRED" | "UNDER_INVESTIGATION";
export type VulnerabilityStatus = "OPEN" | "RESOLVED" | "SUPPRESSED";
export type VEXStatus = "not_affected" | "affected" | "fixed" | "under_investigation";
export type TabKey = "overview" | "events" | "denies" | "runtime" | "analytics" | "topology" | "forensics" | "exceptions" | "inventory" | "vulnerabilities" | "signing" | "scorecard" | "guidance";

export interface AuditHealth {
  status: string;
  backend?: string;
  error?: string;
}

export interface AuthStatus {
  authenticated: boolean;
  auth_mode: string;
  subject?: string;
  role?: string;
  token_id?: string;
  identity_type?: string;
  email?: string;
  tenant_id?: string;
  global_scope?: boolean;
}

export interface SyncStatus {
  sync_mode?: string;
  mode: string;
  cluster_id?: string;
  hub_url?: string;
  fail_mode?: string;
  health: string;
  current_revision?: string;
  revision_etag?: string;
  last_successful_sync_at?: string;
  last_attempt_at?: string;
  last_error?: string;
  cache_present: boolean;
  stale_after_seconds?: number;
  summary?: string;
}

export interface RuntimeDriftFinding {
  id: string;
  tenant_id?: string;
  cluster_id?: string;
  namespace: string;
  workload_kind: string;
  workload: string;
  service_account?: string;
  drift_result: string;
  drift_classes?: string[];
  drift_severity?: string;
  remediation_mode?: string;
  remediation_attempt?: number;
  remediable: boolean;
  status: string;
  quarantine_reason?: string;
  desired_state_verification_state?: string;
  detected_at: string;
  last_updated_at: string;
  last_event_type: string;
  reasons?: string[];
  evidence?: Record<string, unknown>;
}

export interface RuntimeDriftFindingsResponse {
  items: RuntimeDriftFinding[];
}

export interface RuntimeDriftStatus {
  total_findings: number;
  active_findings: number;
  quarantined: number;
  failed: number;
  remediated: number;
  detected: number;
  counts_by_severity: Record<string, number>;
  counts_by_status: Record<string, number>;
  last_detected_at?: string;
  last_updated_at?: string;
}

export interface RuntimeActiveState {
  id: string;
  tenant_id?: string;
  cluster_id?: string;
  namespace: string;
  workload_kind: string;
  workload: string;
  service_account?: string;
  observed_digest?: string;
  approved_digest?: string;
  observed_config_hash?: string;
  expected_config_hash?: string;
  drift_result?: string;
  drift_classes?: string[];
  drift_severity?: string;
  reconciliation_status: string;
  remediation_mode?: string;
  remediation_attempt?: number;
  remediable: boolean;
  quarantine_reason?: string;
  quarantine_type?: string;
  protected_target?: boolean;
  protected_reason?: string;
  desired_state_source_ref?: string;
  desired_state_approval_id?: string;
  desired_state_verification_state?: string;
  last_error?: string;
  last_reconciled_at: string;
  reasons?: string[];
  evidence?: Record<string, unknown>;
}

export interface RuntimeActiveStatesResponse {
  items: RuntimeActiveState[];
}

export interface RuntimeClosedLoopStatus {
  total_targets: number;
  in_sync: number;
  drift_detected: number;
  remediated: number;
  failed: number;
  quarantined: number;
  protected_blocked: number;
  counts_by_status: Record<string, number>;
  counts_by_quarantine_type: Record<string, number>;
  last_reconciled_at?: string;
}

export interface SigningIdentityPolicy {
  id: string;
  name?: string;
  provider_type: string;
  issuer?: string;
  signer_identity?: string;
  subject?: string;
  repository?: string;
  workflow?: string;
  ref?: string;
  tenant_id?: string;
  cluster_id?: string;
  environment?: string;
  enabled: boolean;
  distrusted_after?: string;
  distrust_reason?: string;
  created_at: string;
  updated_at: string;
  created_by?: string;
  updated_by?: string;
}

export interface SigningIdentityObservation {
  id: string;
  provider_type?: string;
  issuer?: string;
  signer_identity?: string;
  subject?: string;
  repository?: string;
  workflow?: string;
  ref?: string;
  commit_sha?: string;
  image_digest?: string;
  tenant_id?: string;
  cluster_id?: string;
  environment?: string;
  first_seen_at?: string;
  last_seen_at?: string;
  event_count: number;
  artifact_count: number;
  verification_state?: string;
  authorized: string;
  matched_policy_id?: string;
  distrusted_after?: string;
  reason_code?: string;
  reason_detail?: string;
}

export interface SigningIdentityFinding {
  id: string;
  type: string;
  severity: string;
  repository?: string;
  workflow?: string;
  ref?: string;
  policy_id?: string;
  observation_id?: string;
  reason: string;
  detected_at?: string;
  advisory: boolean;
}

export interface SigningIdentityStatus {
  enforcement_mode: string;
  require_rekor: boolean;
  total_policies: number;
  enabled_policies: number;
  observed_identities: number;
  authorized: number;
  unauthorized: number;
  unknown: number;
  findings: number;
  workflow_drift_findings: number;
  counts_by_reason_code: Record<string, number>;
}

export interface SigningIdentityObservationsResponse {
  items: SigningIdentityObservation[];
}

export interface SigningIdentityFindingsResponse {
  items: SigningIdentityFinding[];
}

export interface SigningIdentityPoliciesResponse {
  policies: SigningIdentityPolicy[];
}

export interface SigningIdentityStatusResponse {
  status: SigningIdentityStatus;
}

export interface ReasonCount {
  reason: string;
  count: number;
}

export interface Summary {
  total_events: number;
  total_allow: number;
  total_deny: number;
  total_error: number;
  counts_by_event_type: Record<string, number>;
  top_deny_reasons: ReasonCount[];
  recent_runtime_drift_deny: number;
}

export interface VerifierSummary {
  signature_valid: boolean;
  attestation_valid: boolean;
}

export interface StoredEvent {
  id: number;
  received_at: string;
  request_id?: string;
  timestamp?: string;
  component: string;
  event_type: string;
  actor?: string;
  cluster_id?: string;
  tenant_id?: string;
  repo?: string;
  branch?: string;
  environment?: string;
  namespace?: string;
  workload?: string;
  image?: string;
  digest?: string;
  cve_id?: string;
  decision: Decision;
  reasons?: string[];
  drift_result?: string;
  drift_classes?: string[];
  verifier_summary?: VerifierSummary;
  policy_version?: string;
  policy_bundle_id?: string;
  policy_bundle_hash?: string;
  decision_hash?: string;
  is_exception?: boolean;
  exception_id?: string;
  exception_type?: string;
  exception_status?: ExceptionStatus;
  exception_reason?: string;
  exception_ticket_id?: string;
  exception_requested_by?: string;
  exception_requested_at?: string;
  exception_approved_by?: string;
  exception_approved_at?: string;
  exception_rejected_by?: string;
  exception_rejected_at?: string;
  exception_rejection_reason?: string;
  exception_expires_at?: string;
  evidence?: Record<string, unknown>;
  raw_event?: Record<string, unknown>;
}

export interface EventsResponse {
  events: StoredEvent[];
}

export interface EventFilters {
  decision: string;
  component: string;
  repo: string;
  environment: string;
  tenant_id: string;
  limit: string;
}

export interface PolicyException {
  id: number;
  exception_id: string;
  exception_type: string;
  status: ExceptionStatus;
  tenant_id?: string;
  environment?: string;
  namespace?: string;
  repo?: string;
  image_digest?: string;
  cve_id?: string;
  reason: string;
  ticket_id: string;
  requested_by?: string;
  requested_at?: string;
  approved_by?: string;
  approved_at?: string;
  rejected_by?: string;
  rejected_at?: string;
  rejection_reason?: string;
  created_at: string;
  expires_at: string;
  active: boolean;
  last_updated_at?: string;
  metadata?: Record<string, unknown>;
}

export interface ExceptionsResponse {
  exceptions: PolicyException[];
}

export interface ExceptionActionResponse {
  status: string;
  exception: PolicyException;
}

export interface ExceptionReport {
  active: PolicyException[];
  pending?: PolicyException[];
  rejected?: PolicyException[];
  revoked?: PolicyException[];
  expired?: PolicyException[];
  recent_used: StoredEvent[];
  recent_inactive: PolicyException[];
  status_counts?: Record<string, number>;
}

export interface ExceptionRequestInput {
  exception_id: string;
  exception_type: string;
  tenant_id?: string;
  environment?: string;
  namespace?: string;
  repo?: string;
  image_digest?: string;
  cve_id?: string;
  reason: string;
  ticket_id: string;
  ttl_hours?: number;
}

export interface TrendBucket {
  timestamp: string;
  allow_count: number;
  deny_count: number;
  error_count: number;
}

export interface TrendsResponse {
  buckets: TrendBucket[];
  totals: Record<string, number>;
  applied_filters: Record<string, string>;
  metric_trends?: AnalyticsMetricTrend[];
  comparison?: AnalyticsComparisonContext;
  limitations?: string[];
}

export interface AnalyticsComparisonContext {
  window: string;
  compare_to: string;
  group_by: string;
  current_start: string;
  current_end: string;
  baseline_start: string;
  baseline_end: string;
  applied_filters: Record<string, string>;
}

export interface AnalyticsMetricDefinition {
  key: string;
  label: string;
  metric_class: string;
  description: string;
  formula: string;
  grain: string;
  default_window: string;
  segments?: string[];
  exclusions?: string[];
  owner: string;
  interpretation: string;
}

export interface AnalyticsSegmentDelta {
  segment_key: string;
  segment_label: string;
  current_value: number;
  baseline_value: number;
  delta_value: number;
  delta_percent: number;
  direction: string;
}

export interface AnalyticsMetricTrend {
  definition: AnalyticsMetricDefinition;
  current_value: number;
  baseline_value: number;
  delta_value: number;
  delta_percent: number;
  direction: string;
  velocity: string;
  summary: string;
  segment_highlights?: AnalyticsSegmentDelta[];
  limitations?: string[];
}

export interface AnalyticsDeltaResponse {
  definition: AnalyticsMetricDefinition;
  comparison: AnalyticsComparisonContext;
  segments: AnalyticsSegmentDelta[];
  summary: string;
  limitations?: string[];
}

export interface AnalyticsAnomaly {
  type: string;
  title: string;
  metric_key: string;
  reason: string;
  baseline: string;
  deviation: string;
  segment: string;
  severity: string;
  recommended_next_step: string;
  evidence_refs?: string[];
  limitations?: string[];
}

export interface AnalyticsAnomaliesResponse {
  comparison: AnalyticsComparisonContext;
  items: AnalyticsAnomaly[];
  limitations?: string[];
}

export interface AnalyticsScorecardCard {
  definition: AnalyticsMetricDefinition;
  status: string;
  current_value: number;
  baseline_value: number;
  delta_value: number;
  delta_percent: number;
  direction: string;
  summary: string;
}

export interface AnalyticsScorecardsResponse {
  comparison: AnalyticsComparisonContext;
  cards: AnalyticsScorecardCard[];
  limitations?: string[];
}

export interface AnalyticsSegmentCatalogItem {
  group: string;
  values: string[];
}

export interface AnalyticsSegmentsResponse {
  comparison: AnalyticsComparisonContext;
  items: AnalyticsSegmentCatalogItem[];
  limitations?: string[];
}

export interface TopologyNode {
  node_id: string;
  service: string;
  workload?: string;
  namespace?: string;
  cluster?: string;
  environment?: string;
  team?: string;
  repo?: string;
  artifact_digest?: string;
  public_exposure: boolean;
  sensitivity_class: string;
  node_risk_score: number;
  blast_radius_score: number;
  critical_reach_count: number;
  public_entry_flag: boolean;
  sensitive_asset_reach_flag: boolean;
  propagation_class: string;
  trust_boundary_crossings: number;
  last_seen: string;
  evidence_refs?: string[];
}

export interface TopologyEdge {
  source: string;
  target: string;
  edge_type: string;
  connectivity_class: string;
  evidence_source: string;
  confidence: string;
  last_seen?: string;
  environment_scope?: string;
  evidence_refs?: string[];
}

export interface TopologyGraphView {
  nodes: TopologyNode[];
  edges: TopologyEdge[];
}

export interface TopologyGraphSummary {
  declared_nodes: number;
  declared_edges: number;
  observed_nodes: number;
  observed_edges: number;
  effective_nodes: number;
  effective_edges: number;
  public_entry_nodes: number;
  critical_nodes: number;
  high_blast_radius: number;
}

export interface TopologyGraphResponse {
  declared_graph: TopologyGraphView;
  observed_graph: TopologyGraphView;
  effective_graph: TopologyGraphView;
  summary: TopologyGraphSummary;
  applied_filters: Record<string, string>;
  limitations?: string[];
}

export interface TopologyServicesResponse {
  items: TopologyNode[];
  applied_filters: Record<string, string>;
  limitations?: string[];
}

export interface TopologyHeatmapResponse {
  items: TopologyNode[];
  applied_filters: Record<string, string>;
  limitations?: string[];
}

export interface TopologyRiskPath {
  nodes: string[];
  edge_types: string[];
  summary: string;
}

export interface TopologyContainmentOption {
  option_id: string;
  title: string;
  summary: string;
  restriction_plan: string[];
  closed_edge_types: string[];
  estimated_score_reduction: number;
  approval_mode: string;
  evidence_refs?: string[];
}

export interface TopologyBlastRadiusResponse {
  subject_ref: string;
  subject_type: string;
  affected_nodes: TopologyNode[];
  primary_affected_node?: TopologyNode;
  reachable_nodes: TopologyNode[];
  critical_reach_count: number;
  blast_radius_score: number;
  trust_boundary_crossings: number;
  declared_edge_count: number;
  observed_edge_count: number;
  top_risk_paths: TopologyRiskPath[];
  containment_options: TopologyContainmentOption[];
  evidence_refs?: string[];
  limitations?: string[];
}

export interface TopologyDeltaItem {
  node_id: string;
  service: string;
  current_blast_radius_score: number;
  baseline_blast_radius_score: number;
  delta: number;
  edge_additions: number;
  critical_reach_delta: number;
  drift_signals?: string[];
}

export interface TopologyDeltaResponse {
  comparison: AnalyticsComparisonContext;
  items: TopologyDeltaItem[];
  limitations?: string[];
}

export interface TopViolator {
  key: string;
  deny_count: number;
  top_reasons: ReasonCount[];
}

export interface TopViolatorsResponse {
  items: TopViolator[];
  applied_filters: Record<string, string>;
}

export interface DriftWorkloadCount {
  workload: string;
  namespace?: string;
  tenant_id?: string;
  environment?: string;
  count: number;
}

export interface DriftStatsResponse {
  total_runtime_drift_denies: number;
  counts_by_drift_class: Record<string, number>;
  top_drifted_workloads: DriftWorkloadCount[];
  mean_time_to_resolve_seconds?: number | null;
  applied_filters: Record<string, string>;
}

export interface SBOMDocument {
  id: number;
  image_digest: string;
  image_ref?: string;
  sbom_format: string;
  source_ref?: string;
  sbom_hash?: string;
  created_at: string;
}

export interface SBOMComponent {
  id: number;
  image_digest: string;
  component_name: string;
  component_version?: string;
  component_type?: string;
  license?: string;
  purl?: string;
  metadata?: Record<string, unknown>;
  created_at: string;
}

export interface SBOMImageResponse {
  document: SBOMDocument;
  component_count: number;
  components: SBOMComponent[];
}

export interface SBOMComponentsResponse {
  components: SBOMComponent[];
}

export interface ActiveWorkloadRef {
  tenant_id?: string;
  environment?: string;
  namespace?: string;
  workload?: string;
  repo?: string;
  image?: string;
  digest?: string;
}

export interface VEXScope {
  image_digest?: string;
  package_name?: string;
  purl?: string;
  repo?: string;
  workload?: string;
  tenant_id?: string;
  cluster_id?: string;
  environment?: string;
  namespace?: string;
}

export interface VEXMatch {
  id: number;
  source_format: string;
  source_ref?: string;
  vulnerability_id: string;
  status: VEXStatus;
  justification?: string;
  action_statement?: string;
  impact_statement?: string;
  fixed_version?: string;
  created_by?: string;
  updated_by?: string;
  expires_at?: string;
  created_at: string;
  updated_at: string;
}

export interface VEXStatement {
  id: number;
  statement_key?: string;
  source_format: string;
  source_ref?: string;
  vulnerability_id: string;
  scope: VEXScope;
  status: VEXStatus;
  justification?: string;
  action_statement?: string;
  impact_statement?: string;
  fixed_version?: string;
  created_by?: string;
  updated_by?: string;
  expires_at?: string;
  revoked_at?: string;
  revoked_by?: string;
  active: boolean;
  metadata?: Record<string, unknown>;
  created_at: string;
  updated_at: string;
}

export interface VEXStatementsResponse {
  statements: VEXStatement[];
}

export interface VEXStatementActionResponse {
  status: string;
  statement: VEXStatement;
}

export interface VEXStatusSummary {
  active_count: number;
  expiring_count: number;
  revoked_count: number;
  counts_by_status: Record<string, number>;
  applied_filters?: Record<string, string>;
}

export interface VEXCreateInput {
  source_format?: "api";
  source_ref?: string;
  vulnerability_id: string;
  scope: VEXScope;
  status: VEXStatus;
  justification?: string;
  action_statement?: string;
  impact_statement?: string;
  fixed_version?: string;
  expires_at?: string;
  metadata?: Record<string, unknown>;
}

export interface VulnerabilityDecision {
  id: number;
  image_digest: string;
  cve_id: string;
  decision: VulnerabilityDecisionValue;
  justification: string;
  decided_by: string;
  expires_at?: string;
  active: boolean;
  metadata?: Record<string, unknown>;
  created_at: string;
  updated_at: string;
}

export interface VulnerabilityFinding {
  id: number;
  image_digest: string;
  image_ref?: string;
  scan_run_id: number;
  cve_id: string;
  severity?: string;
  package_name?: string;
  package_version?: string;
  fixed_version?: string;
  purl?: string;
  status: VulnerabilityStatus;
  title?: string;
  description?: string;
  source?: string;
  metadata?: Record<string, unknown>;
  first_seen_at: string;
  last_seen_at: string;
  vex?: VEXMatch;
  decision?: VulnerabilityDecision;
}

export interface VulnerabilitiesResponse {
  findings: VulnerabilityFinding[];
}

export interface VulnerabilityNetResponse {
  raw_count: number;
  resolved_by_vex_count: number;
  actionable_count: number;
  under_investigation_count: number;
  severity_threshold?: string;
  threshold_breached: boolean;
  findings: VulnerabilityFinding[];
  applied_filters: Record<string, string>;
}

export interface VulnerabilityBlastRadiusItem {
  image_digest: string;
  image_ref?: string;
  findings: VulnerabilityFinding[];
  workloads: ActiveWorkloadRef[];
}

export interface VulnerabilityBlastRadiusResponse {
  items: VulnerabilityBlastRadiusItem[];
  applied_filters: Record<string, string>;
}

export interface VulnerabilityTimelineEntry {
  image_digest: string;
  cve_id: string;
  package_name?: string;
  package_version?: string;
  severity?: string;
  status: VulnerabilityStatus;
  first_seen_at: string;
  last_seen_at: string;
  decision?: VulnerabilityDecision;
}

export interface VulnerabilityTimelineResponse {
  items: VulnerabilityTimelineEntry[];
  applied_filters: Record<string, string>;
}

export interface VulnerabilityDecisionsResponse {
  decisions: VulnerabilityDecision[];
}

export interface VulnerabilityDecisionInput {
  image_digest: string;
  cve_id: string;
  decision: VulnerabilityDecisionValue;
  justification: string;
  ttl_hours?: number;
}

export interface VulnerabilityRescanResponse {
  status: string;
  scanned_digests: string[];
  scan_runs: number;
}

export interface TrustScoreMetric {
  id: string;
  name: string;
  weight: number;
  score: number;
  status: string;
  reason_code: string;
  reason_detail?: string;
  evidence_refs?: string[];
  advisory_only: boolean;
  public_publishable: boolean;
  mapping_refs?: string[];
}

export interface TrustScorecard {
  id: string;
  scope_type: string;
  scope_ref: string;
  tenant_id?: string;
  cluster_id?: string;
  environment?: string;
  repo?: string;
  calculated_at: string;
  overall_grade: string;
  overall_score: number;
  signing_coverage: number;
  transparency_coverage: number;
  sbom_or_provenance_coverage: number;
  actionable_vulnerability_count: number;
  stale_exception_count: number;
  publication_mode: string;
  metrics: TrustScoreMetric[];
  notes?: string[];
}

export interface TrustBadge {
  id: string;
  label: string;
  state: string;
  summary: string;
  public_publishable: boolean;
  svg?: string;
}

export interface AuditFinding {
  id: string;
  category: string;
  severity: string;
  status: string;
  reason_code: string;
  reason_detail?: string;
  scope_ref?: string;
  evidence_refs?: string[];
  advisory_only: boolean;
  public_publishable: boolean;
  detected_at: string;
}

export interface StandardsMapping {
  standard: string;
  control: string;
  status: string;
  summary: string;
  evidence_refs?: string[];
}

export interface PublishedTrustView {
  generated_at: string;
  scope_type: string;
  scope_ref: string;
  overall_grade: string;
  overall_score: number;
  badges: TrustBadge[];
  metrics: TrustScoreMetric[];
  mapping: StandardsMapping[];
  notes?: string[];
}

export interface AuditReport {
  id: string;
  generated_at: string;
  scope_type: string;
  scope_ref: string;
  scorecard: TrustScorecard;
  findings: AuditFinding[];
  badges: TrustBadge[];
  standards_mapping: StandardsMapping[];
  public_view?: PublishedTrustView;
  limitations?: string[];
  format?: string;
  generated_by?: string;
}

export interface GuidanceGrouping {
  key: string;
  label: string;
  category: string;
  finding_count: number;
  priority: string;
  contextual_risk_score: number;
  heuristic: boolean;
}

export interface GuidanceVEXDraftSuggestion {
  id: string;
  candidate_status: string;
  justification: string;
  impact_statement: string;
  missing_evidence?: string[];
  confidence: string;
  confidence_basis?: string;
  advisory_only: boolean;
  requires_human_review: boolean;
  docs_refs?: string[];
}

export interface BreakGlassGuidance {
  scope_explanation: string;
  narrower_alternative?: string;
  cleanup_reminders?: string[];
  proposed_containment_steps?: string[];
  confidence: string;
  confidence_basis?: string;
  advisory_only: boolean;
  requires_human_review: boolean;
  docs_refs?: string[];
}

export interface GuidanceItem {
  id: string;
  category: string;
  source_component?: string;
  grouping: GuidanceGrouping;
  related_reason_codes?: string[];
  finding_refs?: string[];
  evidence_refs?: string[];
  docs_refs?: string[];
  scope_type?: string;
  scope_ref?: string;
  tenant_id?: string;
  cluster_id?: string;
  environment?: string;
  repository?: string;
  severity?: string;
  priority: string;
  confidence: string;
  confidence_basis?: string;
  explanation?: string;
  recommendation_summary?: string;
  recommendation_steps?: string[];
  safer_alternative?: string;
  impact_summary?: string;
  data_limitations?: string[];
  advisory_only: boolean;
  requires_human_review: boolean;
  generated_at: string;
  generated_by: string;
  template_version?: string;
  heuristic: boolean;
  vex_draft?: GuidanceVEXDraftSuggestion;
  break_glass_guidance?: BreakGlassGuidance;
}

export interface GuidanceSummary {
  total_items: number;
  counts_by_category?: Record<string, number>;
  counts_by_priority?: Record<string, number>;
  guidance_mode: string;
  ai_enabled: boolean;
  deterministic_only: boolean;
  limitations?: string[];
}

export interface GuidanceResponse {
  generated_at: string;
  scope_type?: string;
  scope_ref?: string;
  tenant_id?: string;
  cluster_id?: string;
  environment?: string;
  repository?: string;
  items: GuidanceItem[];
  summary: GuidanceSummary;
}

export interface AIInsightsResponse {
  summary: GuidanceSummary;
  top_items: GuidanceItem[];
}

export interface HistoricalVulnerabilityFinding {
  cve_id: string;
  image_digest?: string;
  severity?: string;
  status?: string;
  known_at_t: boolean;
  first_seen_at?: string;
  last_seen_at?: string;
  evidence_refs?: string[];
}

export interface HistoricalVEXState {
  statement_id: number;
  vulnerability_id: string;
  status: string;
  justification?: string;
  created_at: string;
  revoked_at?: string;
  source_ref?: string;
}

export interface ForensicsPolicyContext {
  policy_bundle_hash?: string;
  active_rules: string[];
  rule_versions: string[];
}

export interface ForensicsInventoryContext {
  running_subjects: string[];
  artifact_digests: string[];
  sbom_refs: string[];
}

export interface ForensicsVulnerabilityContext {
  known_findings: HistoricalVulnerabilityFinding[];
  unknown_later_disclosed_refs: string[];
  vex_state?: HistoricalVEXState[];
}

export interface ForensicsIdentityContext {
  signers: string[];
  trust_roots: string[];
  identity_drift_flags: string[];
}

export interface ForensicsExceptionContext {
  active_exceptions: string[];
  break_glass_state: boolean;
}

export interface ForensicsIncidentSummary {
  incident_id: string;
  state: string;
  severity: string;
  scope_ref?: string;
}

export interface ForensicsIncidentContext {
  relevant_incidents: ForensicsIncidentSummary[];
}

export interface ForensicsTopologyContext {
  advisory_only: boolean;
  primary_service?: string;
  blast_radius_score: number;
  critical_reach_count: number;
  top_risk_paths?: string[];
  heatmap?: TopologyNode[];
  limitations?: string[];
}

export interface PointInTimeState {
  mode: string;
  timestamp: string;
  tenant_id?: string;
  environment?: string;
  subject_summary?: string;
  policy_context: ForensicsPolicyContext;
  inventory_context: ForensicsInventoryContext;
  vulnerability_context: ForensicsVulnerabilityContext;
  identity_context: ForensicsIdentityContext;
  exception_context: ForensicsExceptionContext;
  incident_context: ForensicsIncidentContext;
  topology_context?: ForensicsTopologyContext;
  evidence_refs?: string[];
  limitations?: string[];
}

export interface TimeDeltaSet {
  added?: string[];
  removed?: string[];
  modified?: string[];
}

export interface TimeDeltaResult {
  mode: string;
  comparison: {
    t1: string;
    t2: string;
    source: string;
    analytics_comparison?: AnalyticsComparisonContext;
  };
  policy_delta: TimeDeltaSet;
  inventory_delta: TimeDeltaSet;
  vulnerability_delta: TimeDeltaSet;
  identity_delta: TimeDeltaSet;
  exception_delta: TimeDeltaSet;
  incident_delta: TimeDeltaSet;
  topology_delta?: TopologyDeltaItem[];
  evidence_refs?: string[];
  limitations?: string[];
}

export interface VEXFlashbackResponse {
  mode: string;
  timestamp: string;
  image_digest?: string;
  cve_id?: string;
  historical_vulnerability_state: HistoricalVulnerabilityFinding[];
  disclosed_after_t_refs?: string[];
  vex_flashback: HistoricalVEXState[];
  historical_decision_basis: string;
  evidence_refs?: string[];
  limitations?: string[];
}

export interface ForensicTimelineMarker {
  marker_id: string;
  timestamp: string;
  marker_type: string;
  title: string;
  severity: string;
  subject_ref?: string;
  evidence_refs?: string[];
}

export interface ForensicTimelineResponse {
  mode: string;
  comparison: {
    t1: string;
    t2: string;
    source: string;
    analytics_comparison?: AnalyticsComparisonContext;
  };
  markers: ForensicTimelineMarker[];
  limitations?: string[];
}

export interface ForensicReplayResponse {
  mode: string;
  counterfactual: boolean;
  replay_mode: string;
  historical_timestamp: string;
  historical_verdict: string;
  replay_verdict: string;
  verdict_delta: string;
  policy_delta_applied?: string[];
  vulnerability_delta_applied?: string[];
  identity_delta_applied?: string[];
  explanations?: string[];
  evidence_refs?: string[];
  limitations?: string[];
}
