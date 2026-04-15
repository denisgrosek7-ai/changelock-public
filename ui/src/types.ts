export type Decision = "ALLOW" | "DENY" | "ERROR";
export type ExceptionStatus = "PENDING" | "APPROVED" | "REJECTED" | "REVOKED" | "EXPIRED";
export type TabKey = "overview" | "events" | "denies" | "runtime" | "analytics" | "exceptions";

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
