import type { StoredEvent } from "./types";

export type IncidentSeverity = "critical" | "high" | "medium" | "low";
export type IncidentStatus = "active" | "contained" | "watch";
export type IncidentLifecycleState = "open" | "acknowledged" | "watching" | "resolved" | "reopened";
export type IncidentPriority = "critical" | "high" | "medium" | "low";
export type IncidentImpactTone = "critical" | "warning" | "allow" | "muted";
export type IncidentTimelineOutcome = "deny" | "allow" | "error" | "signal";
export type IncidentReportAudience = "internal" | "auditor_safe" | "customer_safe";

export type IncidentTimelineEntry = {
  id: string;
  kind?: string;
  timestamp?: string;
  title: string;
  summary: string;
  eventType: string;
  outcome: IncidentTimelineOutcome;
  requestID?: string;
  actor?: string;
};

export type InvestigationImpact = {
  id: string;
  title: string;
  detail: string;
  tone: IncidentImpactTone;
};

export type IncidentEvidencePack = {
  requestIDs: string[];
  digests: string[];
  bundles: string[];
  exceptions: string[];
  vulnerabilities: string[];
};

export type IncidentAssignment = {
  owner?: string;
  at?: string;
  by?: string;
  reason?: string;
};

export type IncidentResolution = {
  type?: string;
  summary?: string;
  details?: string;
  refs: string[];
  by?: string;
  at?: string;
  followUpRequired?: boolean;
};

export type IncidentMetricLink = {
  metricKey: string;
  metricLabel: string;
  linkReason: string;
  supportingRefs: string[];
  impactWeight: number;
};

export type IncidentExportEventRef = {
  eventID: number;
  requestID?: string;
  timestamp: string;
  component: string;
  eventType: string;
  decision: string;
  decisionHash?: string;
};

export type IncidentExport = {
  generatedAt: string;
  audience: IncidentReportAudience;
  redacted: boolean;
  redactionSummary: string[];
  incidentID: string;
  identityKey?: string;
  title: string;
  summary: string;
  state: IncidentLifecycleState;
  severity: IncidentSeverity;
  priority: IncidentPriority;
  owner?: string;
  openedAt?: string;
  updatedAt?: string;
  resolvedAt?: string;
  scopeType?: string;
  scopeRef?: string;
  tenantID?: string;
  clusterID?: string;
  environment?: string;
  repository?: string;
  governanceImpacts: InvestigationImpact[];
  reasonCodes: string[];
  findingRefs: string[];
  guidanceRefs: string[];
  scorecardRefs: string[];
  metricLinks: IncidentMetricLink[];
  evidenceRefs: string[];
  evidencePack: IncidentEvidencePack;
  history: IncidentHistoryEntry[];
  resolution: IncidentResolution;
  notes: IncidentNote[];
  newActivityDetected: boolean;
  relatedEventRefs: IncidentExportEventRef[];
  limitations: string[];
};

export type IncidentPackageAggregate = {
  byState: Record<string, number>;
  bySeverity: Record<string, number>;
  byCategory: Record<string, number>;
};

export type IncidentPackageItem = {
  incidentID: string;
  title: string;
  summary: string;
  state: IncidentLifecycleState;
  severity: IncidentSeverity;
  priority: IncidentPriority;
  category: string;
  scopeLabel?: string;
  openedAt?: string;
  updatedAt?: string;
  resolvedAt?: string;
};

export type IncidentPackage = {
  generatedAt: string;
  audience: IncidentReportAudience;
  redacted: boolean;
  redactionSummary: string[];
  selectionMode: "explicit" | "query_derived";
  selectionSummary: string;
  packageSummary: string;
  incidentCount: number;
  incidentRefs: string[];
  aggregate: IncidentPackageAggregate;
  incidents: IncidentPackageItem[];
  packageIntelligence: PackageIntelligence;
  limitations: string[];
};

export type PackageDefenseGapSummary = {
  topGapTypes: string[];
  confidenceMix: Record<string, number>;
  topFindings: DefenseGapFinding[];
  rationale: string;
  limitations: string[];
};

export type PackageReplayOutcomeCurrent = {
  blockingOrSurfacing: number;
  monitoringOnly: number;
  resolvedOrReviewed: number;
};

export type PackageReplayOutcomeProposed = {
  earlierDenials: number;
  evidenceHolds: number;
  earlierContainment: number;
  narrowerExceptions: number;
};

export type PackageReplayDelta = {
  additionalRejections: number;
  earlierContainmentPaths: number;
  impactedCases: number;
};

export type PackagePolicyReplaySummary = {
  currentOutcome: PackageReplayOutcomeCurrent;
  proposedOutcome: PackageReplayOutcomeProposed;
  delta: PackageReplayDelta;
  blastRadius: ReplayBlastRadius;
  topCoverageGaps: CoverageGapFinding[];
  shadowModeImpact: string;
  limitations: string[];
};

export type PackageSystemicPattern = {
  patternKey: string;
  title: string;
  priority: "critical" | "high" | "medium" | "low";
  relatedIncidentRefs: string[];
  evidenceRefs: string[];
};

export type PackageSystemicWeaknessSummary = {
  topPatterns: PackageSystemicPattern[];
  rootCauseHypothesis: string;
  processFragility: boolean;
  supplyChainBlindSpots: boolean;
  executiveRecommendation: string;
  limitations: string[];
};

export type PackageRecommendedActions = {
  whyThisMattersNow: string;
  immediateContainment: string[];
  nearTermHardening: string[];
  governanceFix: string[];
};

export type PackageIntelligence = {
  advisoryOnly: boolean;
  generatedAt: string;
  defenseGapSummary: PackageDefenseGapSummary;
  policyReplaySummary: PackagePolicyReplaySummary;
  systemicWeaknessSummary: PackageSystemicWeaknessSummary;
  recommendedActions: PackageRecommendedActions;
};

export type AdvisoryReadbackRef = {
  resourceType: "defense-gap" | "policy-replay" | "systemic-weakness";
  resourceID: string;
  resourceURI: string;
  evidenceHash: string;
};

export type DecisionVerdictContext = {
  summary?: string;
  currentOutcome?: string;
  proposedOutcome?: string;
  delta?: string;
  patternKey?: string;
  gapTypes: string[];
};

export type DecisionSnapshotRefs = {
  policySnapshotRef: string;
  evaluatorInputHash: string;
  evaluatorOutputHash: string;
  evidenceRefs: string[];
};

export type DecisionEvidenceEnvelope = {
  schemaVersion: string;
  resourceType: string;
  resourceID: string;
  evidenceHash: string;
  generatedAt: string;
  subjectType: string;
  subjectRef: string;
  verdictContext: DecisionVerdictContext;
  snapshotRefs: DecisionSnapshotRefs;
  redactionProfileVersion: string;
  projectionSchemaVersion: string;
  advisoryOnly: boolean;
  limitations: string[];
};

export type AdvisoryReadbackResponse<TPayload = unknown> = {
  resourceType: string;
  resourceID: string;
  permanentURI: string;
  projectionAudience: IncidentReportAudience;
  advisoryOnly: boolean;
  payloadSummary: string;
  evidenceEnvelope: DecisionEvidenceEnvelope;
  payload: TPayload;
  limitations: string[];
};

export type AdvisoryShareGrant = {
  grantID: string;
  shareURL: string;
  resourceType: string;
  resourceID: string;
  audience: IncidentReportAudience;
  expiresAt: string;
  purpose?: string;
};

export type DefenseGapRecommendations = {
  containment: string[];
  hardening: string[];
  governanceFix: string[];
};

export type DefenseGapFinding = {
  gapType: string;
  title: string;
  confidence: "high" | "medium" | "limited";
  whyItMatters: string;
  evidenceRefs: string[];
  relatedIncidentRefs: string[];
  recommendedActions: DefenseGapRecommendations;
};

export type DefenseGapPattern = {
  present: boolean;
  patternKey?: string;
  summary: string;
  relatedIncidentRefs: string[];
};

export type DefenseGapAssessment = {
  assessmentID: string;
  subjectType: "incident" | "metric";
  subjectRef: string;
  generatedAt: string;
  advisoryOnly: boolean;
  defenseGaps: DefenseGapFinding[];
  systemicPattern: DefenseGapPattern;
  readback: AdvisoryReadbackRef;
  limitations: string[];
};

export type CoverageGapFinding = {
  gapType: string;
  title: string;
  summary: string;
  confidence: "high" | "medium" | "limited";
  evidenceRefs: string[];
  relatedIncidentRefs: string[];
  recommendedAction: string;
};

export type PolicyReplayResult = {
  caseRef: string;
  title: string;
  currentOutcome: string;
  proposedOutcome: string;
  delta: string;
  supportingEvidenceRefs: string[];
  confidence: "high" | "medium" | "limited";
  limitations: string[];
};

export type ReplayBlastRadius = {
  incidentCount: number;
  repoCount: number;
  environmentCount: number;
  workloadCount: number;
  topScopes: string[];
};

export type PolicyReplayAssessment = {
  assessmentID: string;
  subjectType: "incident" | "metric" | "scope";
  subjectRef: string;
  generatedAt: string;
  advisoryOnly: boolean;
  shadowMode: boolean;
  replayResults: PolicyReplayResult[];
  coverageGaps: CoverageGapFinding[];
  blastRadius: ReplayBlastRadius;
  readback: AdvisoryReadbackRef;
  limitations: string[];
};

export type SystemicWeakness = {
  patternKey: string;
  title: string;
  priority: "critical" | "high" | "medium" | "low";
  summary: string;
  processFragility: string[];
  supplyChainBlindSpots: string[];
  rootCauseHypothesis: string;
  executiveRecommendation: string;
  relatedIncidentRefs: string[];
  evidenceRefs: string[];
  readback: AdvisoryReadbackRef;
  limitations: string[];
};

export type SystemicWeaknessResponse = {
  generatedAt: string;
  advisoryOnly: boolean;
  scopeSummary: string;
  weaknesses: SystemicWeakness[];
  limitations: string[];
};

export type ExecutiveDefenseSummary = {
  topRisks: string[];
  topImprovements: string[];
  trendChange: string;
  whatMattersNow: string;
};

export type ExecutiveStrategicGap = {
  id: string;
  title: string;
  summary: string;
  investmentTarget: string;
  confidence: "high" | "medium" | "limited";
  relatedIncidentRefs: string[];
  evidenceRefs: string[];
};

export type ExecutiveRiskTrend = {
  key: string;
  label: string;
  direction: "improving" | "watch" | "worsening";
  value: string;
  summary: string;
  evidenceRefs: string[];
};

export type ExecutiveShieldHealthComponent = {
  key: string;
  label: string;
  score: number;
  summary: string;
  evidenceRefs: string[];
};

export type ExecutiveShieldHealth = {
  score: number;
  band: "strong" | "watch" | "at_risk";
  summary: string;
  components: ExecutiveShieldHealthComponent[];
};

export type ExecutiveImpactEstimate = {
  key: string;
  label: string;
  value: string;
  confidence: "high" | "medium" | "limited";
  summary: string;
  assumptions: string[];
};

export type ExecutiveBusinessImpact = {
  summary: string;
  estimates: ExecutiveImpactEstimate[];
};

export type ExecutiveBoardPackage = {
  headline: string;
  narrative: string;
  investmentPriorities: string[];
  nextQuarterPriorities: string[];
  packageSummary: string;
};

export type ExecutiveDefenseReport = {
  generatedAt: string;
  audience: IncidentReportAudience;
  redacted: boolean;
  redactionSummary: string[];
  advisoryOnly: boolean;
  selectionMode: "query_derived" | "explicit";
  scopeSummary: string;
  incidentCount: number;
  incidentRefs: string[];
  executiveSummary: ExecutiveDefenseSummary;
  strategicGaps: ExecutiveStrategicGap[];
  riskReductionTrends: ExecutiveRiskTrend[];
  shieldHealth: ExecutiveShieldHealth;
  businessImpact: ExecutiveBusinessImpact;
  boardPackage: ExecutiveBoardPackage;
  limitations: string[];
};

export type MetricIncidentDrilldown = {
  metricKey: string;
  metricLabel: string;
  incidents: InvestigationIncident[];
  limitations: string[];
};

export type IncidentNote = {
  id: string;
  note: string;
  actor?: string;
  timestamp?: string;
};

export type IncidentHistoryEntry = {
  id: string;
  kind: string;
  timestamp?: string;
  actor?: string;
  summary: string;
  state?: string;
  owner?: string;
  note?: string;
};

export type IncidentLifecycleOverlay = {
  state: IncidentLifecycleState;
  owner?: string;
  assignment: IncidentAssignment;
  resolution: IncidentResolution;
  resolutionSummary?: string;
  notes: IncidentNote[];
  history: IncidentHistoryEntry[];
  lastOperatorUpdateAt?: string;
  newActivityDetected: boolean;
};

export type InvestigationIncident = {
  id: string;
  identityKey?: string;
  categoryKey?: string;
  title: string;
  category: string;
  severity: IncidentSeverity;
  priority: IncidentPriority;
  state: IncidentLifecycleState;
  status: IncidentStatus;
  scopeType?: string;
  scopeRef?: string;
  tenantID?: string;
  clusterID?: string;
  environment?: string;
  repository?: string;
  summary: string;
  caseSummary: string;
  statusNarrative: string;
  likelyCause: string;
  recommendedAction: string;
  remediationChecklist: string[];
  eventCount: number;
  denyCount: number;
  allowCount: number;
  errorCount: number;
  openedAt?: string;
  updatedAt?: string;
  lastActivityAt?: string;
  lastOperatorUpdateAt?: string;
  resolvedAt?: string;
  owner?: string;
  assignment: IncidentAssignment;
  resolution: IncidentResolution;
  lifecycle?: IncidentLifecycleOverlay;
  resolutionSummary?: string;
  newActivityDetected: boolean;
  notes: IncidentNote[];
  history: IncidentHistoryEntry[];
  firstSeenAt?: string;
  lastSeenAt?: string;
  primaryReason: string;
  reasonCodes: string[];
  relatedReasons: string[];
  findingRefs: string[];
  guidanceRefs: string[];
  scorecardRefs: string[];
  metricLinks: IncidentMetricLink[];
  affectedRepos: string[];
  affectedEnvironments: string[];
  affectedTenants: string[];
  affectedNamespaces: string[];
  affectedWorkloads: string[];
  affectedImages: string[];
  affectedComponents: string[];
  evidenceRefs: string[];
  evidencePack: IncidentEvidencePack;
  governanceImpacts: InvestigationImpact[];
  labels: string[];
  timeline: IncidentTimelineEntry[];
  events: StoredEvent[];
};

type IncidentAccumulator = {
  id: string;
  title: string;
  category: string;
  likelyCause: string;
  recommendedAction: string;
  primaryReason: string;
  events: StoredEvent[];
  relatedReasons: Set<string>;
  affectedRepos: Set<string>;
  affectedEnvironments: Set<string>;
  affectedTenants: Set<string>;
  affectedNamespaces: Set<string>;
  affectedWorkloads: Set<string>;
  affectedImages: Set<string>;
  affectedComponents: Set<string>;
  evidenceRefs: Set<string>;
  requestIDs: Set<string>;
  digests: Set<string>;
  bundles: Set<string>;
  exceptions: Set<string>;
  vulnerabilities: Set<string>;
  denyCount: number;
  allowCount: number;
  errorCount: number;
  firstSeenAt?: string;
  lastSeenAt?: string;
};

function toTimestamp(value?: string) {
  if (!value) {
    return 0;
  }
  const parsed = new Date(value).getTime();
  return Number.isFinite(parsed) ? parsed : 0;
}

function eventTimestamp(event: StoredEvent) {
  return event.timestamp || event.received_at;
}

function slug(value: string) {
  return value.trim().toLowerCase().replace(/[^a-z0-9]+/g, "-").replace(/(^-|-$)/g, "") || "generic";
}

function listFromSet(values: Set<string>) {
  return Array.from(values).sort((left, right) => left.localeCompare(right));
}

function compact(values: Array<string | undefined>) {
  return values.map((value) => String(value || "").trim()).filter(Boolean);
}

function classifyIncident(event: StoredEvent) {
  const primaryReason = event.reasons?.[0] || event.drift_result || event.event_type;
  const normalizedReasons = (event.reasons || []).map((reason) => reason.trim().toLowerCase());
  const normalizedPrimary = primaryReason.trim().toLowerCase();
  const driftResult = String(event.drift_result || "").trim().toLowerCase();

  if (normalizedReasons.some((reason) => reason.includes("workflow mismatch")) || normalizedPrimary.includes("workflow mismatch")) {
    return {
      id: "workflow-trust-drift",
      title: "Workflow trust drift",
      category: "signing and workflow governance",
      likelyCause: "A signing-capable workflow or trusted workflow ref changed, but signer policy still expects the previous identity or workflow path.",
      recommendedAction: "Compare the latest workflow ref and signer policy, then update the trusted workflow rule instead of widening exceptions.",
      primaryReason,
    };
  }

  if (driftResult.includes("image") || normalizedReasons.some((reason) => reason.includes("image drift")) || normalizedPrimary.includes("image drift")) {
    return {
      id: "runtime-image-drift",
      title: "Runtime image drift",
      category: "runtime reconciliation",
      likelyCause: "A deployed workload diverged from the last approved image digest or the parent controller spec no longer matches the desired state.",
      recommendedAction: "Inspect the affected workload, compare it to the approved digest, and reconcile the parent controller before restarting pods.",
      primaryReason,
    };
  }

  if (normalizedReasons.some((reason) => reason.includes("signature verification failed")) || normalizedPrimary.includes("signature verification failed")) {
    return {
      id: "artifact-signature-failure",
      title: "Artifact signature failure",
      category: "artifact trust",
      likelyCause: "The artifact verification path did not accept the signature, evidence, or signer identity for the submitted digest.",
      recommendedAction: "Review the failing digest, signer identity, and transparency evidence before re-running the deployment.",
      primaryReason,
    };
  }

  if (normalizedReasons.some((reason) => reason.includes("digest-pinned")) || normalizedPrimary.includes("digest-pinned")) {
    return {
      id: "digest-pinning-gap",
      title: "Digest pinning hygiene gap",
      category: "artifact integrity",
      likelyCause: "The deployment still references mutable tags or a trust flow that requires digest pinning but did not receive it.",
      recommendedAction: "Pin the image to an immutable digest and keep the same digest through verification, policy, and deployment.",
      primaryReason,
    };
  }

  if (event.exception_id || event.event_type.startsWith("exception_")) {
    return {
      id: "exception-governance",
      title: "Exception governance pressure",
      category: "approval governance",
      likelyCause: "The current scope is leaning on active, pending, or recently used exceptions that still need bounded review.",
      recommendedAction: "Review active exception scope, confirm the evidence trail, and revoke or narrow any entry that is no longer justified.",
      primaryReason,
    };
  }

  if (event.decision === "ERROR") {
    return {
      id: "control-plane-error",
      title: "Control-plane execution error",
      category: "platform reliability",
      likelyCause: "At least one audit or decision path failed before returning a clean allow or deny result.",
      recommendedAction: "Open the related evidence payload and verify backend, verifier, and policy-engine health before retrying affected operations.",
      primaryReason,
    };
  }

  if (event.event_type === "runtime_drift_result") {
    return {
      id: "runtime-drift",
      title: "Runtime hardening drift",
      category: "runtime reconciliation",
      likelyCause: "Workload state changed after deployment and drift controls are now reporting it back into the audit path.",
      recommendedAction: "Review the drift class, identify the owning workload, and decide whether remediation or quarantine is appropriate.",
      primaryReason,
    };
  }

  if (event.event_type === "deploy_gate_decision" || event.event_type === "policy_decision") {
    return {
      id: `policy-enforcement-${slug(primaryReason)}`,
      title: "Policy enforcement regression",
      category: "policy governance",
      likelyCause: "A policy or admission condition is repeatedly rejecting the same class of change in the current scope.",
      recommendedAction: "Review the repeated deny reason, compare it to the latest manifest or workflow change, and fix the source rather than suppressing the signal.",
      primaryReason,
    };
  }

  return {
    id: `investigation-${slug(primaryReason)}`,
    title: "Investigation cluster",
    category: "general control-plane signal",
    likelyCause: "Several related audit signals share the same reason pattern and should be triaged together.",
    recommendedAction: "Inspect the linked events and evidence payloads to confirm whether this is one root cause or several separate issues.",
    primaryReason,
  };
}

function severityForIncident(accumulator: IncidentAccumulator): IncidentSeverity {
  if (accumulator.errorCount > 0 || accumulator.denyCount >= 25) {
    return "critical";
  }
  if (accumulator.denyCount >= 10 || accumulator.events.length >= 16) {
    return "high";
  }
  if (accumulator.denyCount >= 4 || accumulator.events.length >= 6) {
    return "medium";
  }
  return "low";
}

function statusForIncident(accumulator: IncidentAccumulator): IncidentStatus {
  if (accumulator.errorCount > 0 || accumulator.denyCount > 0) {
    return "active";
  }
  if (accumulator.allowCount > 0 && accumulator.events.length > 0) {
    return "watch";
  }
  return "contained";
}

function lifecycleStateForIncident(accumulator: IncidentAccumulator): IncidentLifecycleState {
  return statusForIncident(accumulator) === "contained" ? "watching" : "open";
}

function priorityForIncident(accumulator: IncidentAccumulator): IncidentPriority {
  const severity = severityForIncident(accumulator);
  const status = statusForIncident(accumulator);
  if (severity === "critical" || (severity === "high" && status === "active")) {
    return "critical";
  }
  if (severity === "high") {
    return "high";
  }
  if (severity === "medium" || status === "watch") {
    return "medium";
  }
  return "low";
}

function incidentSummary(accumulator: IncidentAccumulator) {
  const parts = [
    `${accumulator.events.length} related events`,
    `${accumulator.denyCount} deny`,
    `${accumulator.errorCount} error`,
  ];
  return `${parts.join(" · ")}. Dominant reason: ${accumulator.primaryReason}.`;
}

function caseSummary(accumulator: IncidentAccumulator) {
  const repoCount = accumulator.affectedRepos.size;
  const envCount = accumulator.affectedEnvironments.size;
  const workloadCount = accumulator.affectedWorkloads.size;
  return `This case groups ${accumulator.events.length} related signals across ${repoCount || 0} repos, ${envCount || 0} environments, and ${workloadCount || 0} workloads around the same dominant reason pattern.`;
}

function statusNarrative(accumulator: IncidentAccumulator, status: IncidentStatus) {
  if (status === "active") {
    return `This incident is still active because deny or error paths are still present in the current event set. Treat it as an open operator issue until the repeated decision path stops.`;
  }
  if (status === "watch") {
    return `The hard failure path has eased, but related signals are still appearing. Keep it under watch until a clean cycle confirms the change is stable.`;
  }
  return `No current deny or error path is visible in the grouped events. Keep the evidence trail attached, but this case is currently contained.`;
}

function addScope(target: Set<string>, value?: string) {
  const normalized = String(value || "").trim();
  if (normalized) {
    target.add(normalized);
  }
}

function formatEventType(value: string) {
  return value
    .split("_")
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(" ");
}

function eventTitle(event: StoredEvent) {
  switch (event.event_type) {
    case "deploy_gate_decision":
      return "Deploy gate decision";
    case "policy_decision":
      return "Policy decision";
    case "runtime_drift_result":
      return "Runtime drift signal";
    case "artifact_verification_result":
      return "Artifact verification result";
    case "exception_created":
      return "Exception requested";
    case "exception_rejected":
      return "Exception rejected";
    case "exception_revoked":
      return "Exception revoked";
    case "exception_used":
      return "Exception consumed";
    default:
      return formatEventType(event.event_type);
  }
}

function eventOutcome(event: StoredEvent): IncidentTimelineOutcome {
  if (event.decision === "DENY") {
    return "deny";
  }
  if (event.decision === "ALLOW") {
    return "allow";
  }
  if (event.decision === "ERROR") {
    return "error";
  }
  return "signal";
}

function eventTimelineSummary(event: StoredEvent) {
  const parts = [
    event.reasons?.[0] || event.drift_result || event.component,
    event.repo,
    event.environment,
    event.workload || event.namespace,
  ].filter(Boolean);
  return parts.join(" · ");
}

function buildTimeline(events: StoredEvent[]): IncidentTimelineEntry[] {
  return [...events]
    .sort((left, right) => toTimestamp(eventTimestamp(left)) - toTimestamp(eventTimestamp(right)))
    .slice(-8)
    .map((event) => ({
      id: `${event.id}`,
      kind: "finding_attached",
      timestamp: eventTimestamp(event),
      title: eventTitle(event),
      summary: eventTimelineSummary(event),
      eventType: event.event_type,
      outcome: eventOutcome(event),
      requestID: event.request_id,
      actor: event.actor,
    }));
}

function buildChecklist(categoryKey: string): string[] {
  if (categoryKey.includes("workflow")) {
    return [
      "Compare the latest signing-capable workflow ref with the trusted workflow path expected by signer policy.",
      "Review the newest request IDs and confirm the signer identity still belongs to the intended repo and ref scope.",
      "Update the workflow trust rule or signer policy instead of widening runtime or deployment exceptions.",
    ];
  }
  if (categoryKey.includes("runtime")) {
    return [
      "Inspect the affected workload and parent controller to confirm the approved image digest is still the desired state.",
      "Reconcile the parent spec before restarting pods so remediation does not loop on the same runtime change.",
      "Use the linked drift evidence to decide whether the case needs containment, reconciliation, or a documented exception review.",
    ];
  }
  if (categoryKey.includes("artifact")) {
    return [
      "Check the failing digest, signer identity, and transparency evidence before retrying the deployment.",
      "Confirm the artifact was signed by an authorized identity and that the trust bundle still matches the environment policy.",
      "Only retry once signature, attestation, and evidence verification all line up with the same immutable artifact digest.",
    ];
  }
  if (categoryKey.includes("digest")) {
    return [
      "Replace mutable image tags with immutable digests in the manifest or release input.",
      "Verify the same digest is used consistently across policy evaluation, verification, and deployment.",
      "Re-run preflight and deploy-gate after pinning so the trust path uses one artifact identity end to end.",
    ];
  }
  if (categoryKey.includes("exception")) {
    return [
      "Review active and recently used exceptions to confirm the current scope is still justified and bounded.",
      "Link the incident review to the recorded exception evidence before approving, extending, or revoking anything.",
      "Prefer narrowing or expiring the exception instead of letting it become a standing bypass.",
    ];
  }
  if (categoryKey.includes("control")) {
    return [
      "Open the newest failing request IDs and verify backend health before retrying affected operations.",
      "Confirm verifier, policy, and evidence services returned coherent results instead of partial failures.",
      "Escalate the platform reliability path if repeated control-plane errors continue to mask allow-or-deny decisions.",
    ];
  }
  return [
    "Review the repeated reason pattern and confirm whether the grouped events really share one root cause.",
    "Use the attached evidence refs and request IDs to validate the highest-severity path first.",
    "Fix the source change or trust mismatch before considering any exception or scope widening.",
  ];
}

function buildGovernanceImpacts(accumulator: IncidentAccumulator): InvestigationImpact[] {
  const impacts: InvestigationImpact[] = [];

  if (accumulator.denyCount > 0) {
    impacts.push({
      id: "policy-enforcement",
      title: "Policy enforcement impact",
      detail: `${accumulator.denyCount} deny decisions are part of this case, so deploy-time governance is already actively blocking it.`,
      tone: accumulator.denyCount >= 10 ? "critical" : "warning",
    });
  }

  if (accumulator.events.some((event) => event.verifier_summary && (!event.verifier_summary.signature_valid || !event.verifier_summary.attestation_valid))) {
    impacts.push({
      id: "evidence-verification",
      title: "Artifact and evidence trust impact",
      detail: "Verifier output shows signature or attestation validation pressure inside this incident. Keep evidence review attached to any remediation.",
      tone: "warning",
    });
  }

  if (accumulator.exceptions.size > 0 || accumulator.events.some((event) => event.event_type.startsWith("exception_"))) {
    impacts.push({
      id: "exception-governance",
      title: "Exception governance impact",
      detail: "Exception activity is already part of this case, so any approval or revocation should be reviewed against the same evidence trail.",
      tone: "warning",
    });
  }

  if (accumulator.events.some((event) => event.event_type === "runtime_drift_result" || String(event.drift_result || "").trim() !== "")) {
    impacts.push({
      id: "runtime-hardening",
      title: "Runtime hardening impact",
      detail: "Runtime drift signals are part of this case, which means the incident can affect closed-loop reconciliation or containment decisions.",
      tone: "warning",
    });
  }

  if (accumulator.vulnerabilities.size > 0) {
    impacts.push({
      id: "vulnerability-posture",
      title: "Vulnerability and VEX impact",
      detail: "Linked CVE references mean vulnerability triage or VEX guidance may be part of the review path for this incident.",
      tone: "muted",
    });
  }

  if (accumulator.errorCount > 0) {
    impacts.push({
      id: "platform-reliability",
      title: "Control-plane reliability impact",
      detail: `${accumulator.errorCount} error events are attached to the same case, so reliability review should stay coupled to the incident narrative.`,
      tone: "critical",
    });
  }

  return impacts.slice(0, 5);
}

export function buildIncidents(events: StoredEvent[]): InvestigationIncident[] {
  const grouped = new Map<string, IncidentAccumulator>();

  for (const event of events) {
    const classification = classifyIncident(event);
    const timestamp = eventTimestamp(event);
    const existing = grouped.get(classification.id) || {
      id: classification.id,
      title: classification.title,
      category: classification.category,
      likelyCause: classification.likelyCause,
      recommendedAction: classification.recommendedAction,
      primaryReason: classification.primaryReason,
      events: [],
      relatedReasons: new Set<string>(),
      affectedRepos: new Set<string>(),
      affectedEnvironments: new Set<string>(),
      affectedTenants: new Set<string>(),
      affectedNamespaces: new Set<string>(),
      affectedWorkloads: new Set<string>(),
      affectedImages: new Set<string>(),
      affectedComponents: new Set<string>(),
      evidenceRefs: new Set<string>(),
      requestIDs: new Set<string>(),
      digests: new Set<string>(),
      bundles: new Set<string>(),
      exceptions: new Set<string>(),
      vulnerabilities: new Set<string>(),
      denyCount: 0,
      allowCount: 0,
      errorCount: 0,
      firstSeenAt: timestamp,
      lastSeenAt: timestamp,
    };

    existing.events.push(event);
    compact(event.reasons || []).forEach((reason) => existing.relatedReasons.add(reason));
    addScope(existing.affectedRepos, event.repo);
    addScope(existing.affectedEnvironments, event.environment);
    addScope(existing.affectedTenants, event.tenant_id);
    addScope(existing.affectedNamespaces, event.namespace);
    addScope(existing.affectedWorkloads, event.workload || event.namespace);
    addScope(existing.affectedImages, event.image || event.digest);
    addScope(existing.affectedComponents, event.component);
    compact([event.request_id, event.digest, event.policy_bundle_id, event.exception_id, event.cve_id]).forEach((ref) => existing.evidenceRefs.add(ref));
    compact([event.request_id]).forEach((ref) => existing.requestIDs.add(ref));
    compact([event.digest]).forEach((ref) => existing.digests.add(ref));
    compact([event.policy_bundle_id]).forEach((ref) => existing.bundles.add(ref));
    compact([event.exception_id]).forEach((ref) => existing.exceptions.add(ref));
    compact([event.cve_id]).forEach((ref) => existing.vulnerabilities.add(ref));

    if (event.decision === "DENY") existing.denyCount += 1;
    if (event.decision === "ALLOW") existing.allowCount += 1;
    if (event.decision === "ERROR") existing.errorCount += 1;

    if (!existing.firstSeenAt || toTimestamp(timestamp) < toTimestamp(existing.firstSeenAt)) {
      existing.firstSeenAt = timestamp;
    }
    if (!existing.lastSeenAt || toTimestamp(timestamp) > toTimestamp(existing.lastSeenAt)) {
      existing.lastSeenAt = timestamp;
    }

    grouped.set(classification.id, existing);
  }

  return Array.from(grouped.values())
    .map((incident) => ({
      id: incident.id,
      identityKey: incident.id,
      categoryKey: incident.id,
      title: incident.title,
      category: incident.category,
      severity: severityForIncident(incident),
      priority: priorityForIncident(incident),
      state: lifecycleStateForIncident(incident),
      status: statusForIncident(incident),
      scopeType: incident.affectedRepos.size > 0 ? "repository" : incident.affectedWorkloads.size > 0 ? "workload" : incident.affectedEnvironments.size > 0 ? "environment" : "component",
      scopeRef:
        listFromSet(incident.affectedRepos)[0] ||
        listFromSet(incident.affectedWorkloads)[0] ||
        listFromSet(incident.affectedEnvironments)[0] ||
        listFromSet(incident.affectedComponents)[0],
      tenantID: listFromSet(incident.affectedTenants)[0],
      environment: listFromSet(incident.affectedEnvironments)[0],
      repository: listFromSet(incident.affectedRepos)[0],
      summary: incidentSummary(incident),
      caseSummary: caseSummary(incident),
      statusNarrative: statusNarrative(incident, statusForIncident(incident)),
      likelyCause: incident.likelyCause,
      recommendedAction: incident.recommendedAction,
      remediationChecklist: buildChecklist(incident.id),
      eventCount: incident.events.length,
      denyCount: incident.denyCount,
      allowCount: incident.allowCount,
      errorCount: incident.errorCount,
      openedAt: incident.firstSeenAt,
      updatedAt: incident.lastSeenAt,
      lastActivityAt: incident.lastSeenAt,
      lastOperatorUpdateAt: undefined,
      owner: undefined,
      assignment: {},
      resolution: { refs: [] },
      lifecycle: {
        state: lifecycleStateForIncident(incident),
        assignment: {},
        resolution: { refs: [] },
        notes: [],
        history: [],
        newActivityDetected: false,
      },
      newActivityDetected: false,
      notes: [],
      history: [],
      firstSeenAt: incident.firstSeenAt,
      lastSeenAt: incident.lastSeenAt,
      primaryReason: incident.primaryReason,
      reasonCodes: listFromSet(incident.relatedReasons),
      relatedReasons: listFromSet(incident.relatedReasons),
      findingRefs: incident.events.flatMap((event) => [`event:${event.id}`, ...(event.decision_hash ? [`decision:${event.decision_hash}`] : [])]).slice(0, 16),
      guidanceRefs: [],
      scorecardRefs: [],
      metricLinks: [],
      affectedRepos: listFromSet(incident.affectedRepos),
      affectedEnvironments: listFromSet(incident.affectedEnvironments),
      affectedTenants: listFromSet(incident.affectedTenants),
      affectedNamespaces: listFromSet(incident.affectedNamespaces),
      affectedWorkloads: listFromSet(incident.affectedWorkloads),
      affectedImages: listFromSet(incident.affectedImages),
      affectedComponents: listFromSet(incident.affectedComponents),
      evidenceRefs: listFromSet(incident.evidenceRefs).slice(0, 8),
      evidencePack: {
        requestIDs: listFromSet(incident.requestIDs).slice(0, 8),
        digests: listFromSet(incident.digests).slice(0, 8),
        bundles: listFromSet(incident.bundles).slice(0, 6),
        exceptions: listFromSet(incident.exceptions).slice(0, 6),
        vulnerabilities: listFromSet(incident.vulnerabilities).slice(0, 6),
      },
      governanceImpacts: buildGovernanceImpacts(incident),
      labels: [slug(incident.category)],
      timeline: buildTimeline(incident.events),
      events: incident.events.sort((left, right) => toTimestamp(eventTimestamp(right)) - toTimestamp(eventTimestamp(left))),
    }))
    .sort((left, right) => {
      const severityRank: Record<IncidentSeverity, number> = { critical: 4, high: 3, medium: 2, low: 1 };
      const severityDelta = severityRank[right.severity] - severityRank[left.severity];
      if (severityDelta !== 0) {
        return severityDelta;
      }
      return (toTimestamp(right.lastSeenAt) - toTimestamp(left.lastSeenAt)) || (right.eventCount - left.eventCount);
    });
}
