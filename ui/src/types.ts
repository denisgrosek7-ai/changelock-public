export type Decision = "ALLOW" | "DENY" | "ERROR";
export type ExceptionStatus = "PENDING" | "APPROVED" | "REJECTED" | "REVOKED" | "EXPIRED";
export type VulnerabilityDecisionValue = "NOT_AFFECTED" | "ACCEPTED_RISK" | "FIX_REQUIRED" | "UNDER_INVESTIGATION";
export type VulnerabilityStatus = "OPEN" | "RESOLVED" | "SUPPRESSED";
export type VEXStatus = "not_affected" | "affected" | "fixed" | "under_investigation";
export type TabKey = "overview" | "events" | "denies" | "runtime" | "analytics" | "exceptions" | "inventory" | "vulnerabilities";

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
