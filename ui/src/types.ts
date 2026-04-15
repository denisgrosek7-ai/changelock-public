export type Decision = "ALLOW" | "DENY" | "ERROR";
export type TabKey = "overview" | "events" | "denies" | "runtime";

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
  exception_reason?: string;
  exception_ticket_id?: string;
  exception_approved_by?: string;
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
