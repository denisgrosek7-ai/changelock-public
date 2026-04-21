package main

import (
	"context"
	"errors"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	runtimeRulePackSchemaVersion     = "3a.runtime_rule_pack.v1"
	runtimeExplainabilitySchema      = "3a.runtime_explainability.v1"
	runtimePostureSchemaVersion      = "3a.runtime_posture.v1"
	runtimeRulePackStatusEnforceable = "enforceable"
	runtimeRulePackStatusDetectOnly  = "detect_only"
	runtimeSchedulingAllowStandard   = "allow_standard"
	runtimeSchedulingRestricted      = "schedule_restricted"
	runtimeSchedulingIsolatedReview  = "schedule_isolated_review"
	runtimeMismatchAttestation       = "attestation_missing_for_verified_desired_state"
	runtimeMismatchSigner            = "expected_signer_evidence_missing"
	runtimeMismatchDesiredState      = "desired_state_not_verified"
	runtimeMismatchIdentity          = "runtime_identity_drift"
	runtimeMismatchSBOM              = "runtime_sbom_drift"
	runtimeMismatchCriticalFindings  = "critical_runtime_findings_present"
)

type runtimeRulePack struct {
	SchemaVersion     string   `json:"schema_version"`
	PackID            string   `json:"pack_id"`
	DisplayName       string   `json:"display_name"`
	FindingTypes      []string `json:"finding_types"`
	TriggerCondition  string   `json:"trigger_condition"`
	EvidenceModel     []string `json:"evidence_model,omitempty"`
	DefaultSeverity   string   `json:"default_severity"`
	ConfidenceModel   string   `json:"confidence_model"`
	DefaultNextAction string   `json:"default_next_action"`
	ApprovalModel     string   `json:"approval_model"`
	RollbackPosture   string   `json:"rollback_posture"`
	ForensicLinkage   string   `json:"forensic_linkage"`
	ExecutionStatus   string   `json:"execution_status"`
	Limitations       []string `json:"limitations,omitempty"`
}

type runtimeRulePackListResponse struct {
	SchemaVersion string            `json:"schema_version"`
	Items         []runtimeRulePack `json:"items"`
	Limitations   []string          `json:"limitations,omitempty"`
}

type runtimeExplainability struct {
	SchemaVersion string                            `json:"schema_version"`
	Trigger       string                            `json:"trigger"`
	TriggerSource string                            `json:"trigger_source,omitempty"`
	EvidenceRefs  []string                          `json:"evidence_refs,omitempty"`
	TrustContext  runtimeExplainabilityTrustContext `json:"trust_context"`
	ResponsePath  runtimeExplainabilityResponsePath `json:"response_path"`
	Forensics     runtimeExplainabilityForensics    `json:"forensics"`
	Topology      []string                          `json:"topology_context,omitempty"`
	NextSteps     []string                          `json:"next_steps,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type runtimeExplainabilityTrustContext struct {
	DesiredStateVerification string   `json:"desired_state_verification_state,omitempty"`
	AttestationInputs        []string `json:"attestation_inputs,omitempty"`
	ExpectedSigners          []string `json:"expected_signers,omitempty"`
	SandboxClass             string   `json:"sandbox_class,omitempty"`
	SBOMStatus               string   `json:"sbom_verification_status,omitempty"`
	CurrentPosture           string   `json:"current_enforcement_posture,omitempty"`
	IdentityStatus           string   `json:"identity_status,omitempty"`
}

type runtimeExplainabilityResponsePath struct {
	PolicyRef         string   `json:"policy_ref,omitempty"`
	RecommendedAction string   `json:"recommended_action,omitempty"`
	SelectedAction    string   `json:"selected_action,omitempty"`
	ApprovalMode      string   `json:"approval_mode,omitempty"`
	ApprovalRequired  bool     `json:"approval_required"`
	RollbackHints     []string `json:"rollback_hints,omitempty"`
}

type runtimeExplainabilityForensics struct {
	ForensicContextURI string                `json:"forensic_context_uri,omitempty"`
	ReadbackRefs       []advisoryReadbackRef `json:"readback_refs,omitempty"`
}

type runtimePostureTrustState struct {
	DesiredStateVerification  string   `json:"desired_state_verification_state,omitempty"`
	AttestationInputs         []string `json:"attestation_inputs,omitempty"`
	ExpectedSigners           []string `json:"expected_signers,omitempty"`
	SBOMStatus                string   `json:"sbom_verification_status,omitempty"`
	SandboxClass              string   `json:"sandbox_class,omitempty"`
	CurrentEnforcementPosture string   `json:"current_enforcement_posture,omitempty"`
}

type runtimePostureMismatch struct {
	Code        string `json:"code"`
	Severity    string `json:"severity"`
	Summary     string `json:"summary"`
	EvidenceRef string `json:"evidence_ref,omitempty"`
}

type runtimeSchedulingGuidance struct {
	Decision          string   `json:"decision"`
	ReasonCodes       []string `json:"reason_codes,omitempty"`
	ApprovalMode      string   `json:"approval_mode"`
	RecommendedAction string   `json:"recommended_action,omitempty"`
}

type runtimePostureState struct {
	SchemaVersion      string                    `json:"schema_version"`
	SubjectRef         string                    `json:"subject_ref"`
	RuntimeModuleReady bool                      `json:"runtime_module_ready"`
	ReadinessSignals   []string                  `json:"readiness_signals,omitempty"`
	ExpectedTrustState runtimePostureTrustState  `json:"expected_trust_state"`
	ActualTrustState   runtimePostureTrustState  `json:"actual_trust_state"`
	Mismatches         []runtimePostureMismatch  `json:"mismatches,omitempty"`
	SchedulingGuidance runtimeSchedulingGuidance `json:"scheduling_guidance"`
	EvidenceRefs       []string                  `json:"evidence_refs,omitempty"`
	LastVerifiedAt     time.Time                 `json:"last_verified_at"`
	Limitations        []string                  `json:"limitations,omitempty"`
}

type runtimePostureListResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Items         []runtimePostureState `json:"items"`
	Limitations   []string              `json:"limitations,omitempty"`
}

func (s server) runtimeRulePacksHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, runtimeRulePackListResponse{
		SchemaVersion: runtimeRulePackSchemaVersion,
		Items:         runtimeRulePackCatalog(),
		Limitations: []string{
			"Rule-pack metadata documents stable runtime semantics and bounded response posture for the current backend implementation; it does not claim kernel-level coverage beyond the evidence paths already present in ChangeLock.",
		},
	})
}

func (s server) runtimeRulePackByIDHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	packID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/runtime/rule-packs/"))
	if packID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "runtime rule pack not found"})
		return
	}
	pack, ok := runtimeRulePackByID(packID)
	if !ok {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "runtime rule pack not found"})
		return
	}
	httpjson.Write(w, http.StatusOK, pack)
}

func (s server) runtimePostureHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildRuntimePostureStates(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, runtimePostureListResponse{
		SchemaVersion: runtimePostureSchemaVersion,
		Items:         items,
		Limitations:   limitations,
	})
}

func (s server) buildRuntimePostureStates(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimePostureState, []string, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	items := make([]runtimePostureState, 0, len(snapshot.subjects))
	for _, subject := range snapshot.sortedSubjects() {
		profile, err := s.profileFromSubject(ctx, filter, subject)
		if err != nil {
			return nil, nil, err
		}
		findings, err := s.findingsForSubject(ctx, filter, subject, profile)
		if err != nil {
			return nil, nil, err
		}
		sbom := s.subjectSBOMVerification(subject)
		sandbox := s.buildRuntimeSandboxDecision(subject, findings, sbom)
		state := s.buildRuntimeIntegrityState(subject, findings, sandbox, sbom)
		items = append(items, buildRuntimePostureState(subject, findings, sbom, sandbox, state))
	}
	sort.Slice(items, func(i, j int) bool {
		if len(items[i].Mismatches) == len(items[j].Mismatches) {
			return items[i].SubjectRef < items[j].SubjectRef
		}
		return len(items[i].Mismatches) > len(items[j].Mismatches)
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, []string{
		"Runtime posture links desired-state verification, signer/attestation evidence, SBOM state, and current enforcement posture into a bounded workload scheduling signal.",
		"Current posture guidance is workload-scoped. Node-level confidential substrate, enclave, or kernel attestation claims are not inferred when the corresponding evidence does not exist in scope.",
	}, nil
}

func runtimeRulePackCatalog() []runtimeRulePack {
	items := []runtimeRulePack{
		{
			SchemaVersion:     runtimeRulePackSchemaVersion,
			PackID:            "binary_execution_integrity",
			DisplayName:       "Binary Execution Integrity",
			FindingTypes:      []string{runtimeFindingUnknownBinaryExec, runtimeFindingUnsignedBinaryExec},
			TriggerCondition:  "Unexpected or unsigned executable paths are observed in runtime evidence for a workload.",
			EvidenceModel:     []string{"runtime observation", "desired-state digest", "forensic context"},
			DefaultSeverity:   "critical",
			ConfidenceModel:   "high when execution signal is explicit; medium when only drift correlation exists",
			DefaultNextAction: runtimeActionApplyNetworkIsolation,
			ApprovalModel:     recommendationApprovalHumanReview,
			RollbackPosture:   "temporary isolation is reversible after clean verification and forensics review",
			ForensicLinkage:   "forensic snapshot should be linked before destructive recovery",
			ExecutionStatus:   runtimeRulePackStatusEnforceable,
			Limitations: []string{
				"This pack relies on evidence-backed runtime execution signals and does not claim universal pre-CPU prevention semantics.",
			},
		},
		{
			SchemaVersion:     runtimeRulePackSchemaVersion,
			PackID:            "runtime_identity_and_attestation",
			DisplayName:       "Runtime Identity and Attestation",
			FindingTypes:      []string{runtimeFindingIdentityDrift, runtimeFindingContainerIDMismatch, runtimeFindingAttestationMismatch},
			TriggerCondition:  "Observed runtime identity, digest lineage, signer evidence, or attestation posture no longer matches expected trusted inputs.",
			EvidenceModel:     []string{"artifact verification evidence", "runtime drift state", "desired-state verification"},
			DefaultSeverity:   "high",
			ConfidenceModel:   "medium to high depending on desired-state verification and artifact evidence completeness",
			DefaultNextAction: runtimeActionCaptureForensics,
			ApprovalModel:     recommendationApprovalHumanReview,
			RollbackPosture:   "trusted recovery or isolation rollback requires clean verification",
			ForensicLinkage:   "always retain forensic linkage when attestation mismatch gates recovery or scheduling",
			ExecutionStatus:   runtimeRulePackStatusEnforceable,
		},
		{
			SchemaVersion:     runtimeRulePackSchemaVersion,
			PackID:            "sbom_runtime_integrity",
			DisplayName:       "SBOM and Library Integrity",
			FindingTypes:      []string{runtimeFindingUnexpectedLibrary, runtimeFindingSBOMMismatch},
			TriggerCondition:  "Observed runtime library or digest state diverges from SBOM-linked or approved artifact expectations.",
			EvidenceModel:     []string{"runtime SBOM verification", "observed library refs", "approved and observed digests"},
			DefaultSeverity:   "high",
			ConfidenceModel:   "high for digest mismatches; medium when library evidence is partial",
			DefaultNextAction: runtimeActionRestartTrusted,
			ApprovalModel:     recommendationApprovalHumanReview,
			RollbackPosture:   "rollback is allowed only after trusted image or clean workload state is verified",
			ForensicLinkage:   "link forensics before restart when mismatch may hide tampering",
			ExecutionStatus:   runtimeRulePackStatusEnforceable,
		},
		{
			SchemaVersion:     runtimeRulePackSchemaVersion,
			PackID:            "outbound_and_topology_expansion",
			DisplayName:       "Outbound and Topology Expansion",
			FindingTypes:      []string{runtimeFindingOutboundDrift, runtimeFindingTopologyExpansion},
			TriggerCondition:  "Unexpected outbound activity or widened blast radius appears in the current runtime scope.",
			EvidenceModel:     []string{"runtime network observation", "topology blast-radius context", "incident linkage"},
			DefaultSeverity:   "medium",
			ConfidenceModel:   "medium for isolated egress; high when paired with elevated topology expansion",
			DefaultNextAction: runtimeActionRecommendQuarantine,
			ApprovalModel:     recommendationApprovalHumanReview,
			RollbackPosture:   "network containment should remain bounded and reversible with TTL",
			ForensicLinkage:   "topology and forensics should remain attached to containment review",
			ExecutionStatus:   runtimeRulePackStatusEnforceable,
		},
		{
			SchemaVersion:     runtimeRulePackSchemaVersion,
			PackID:            "privilege_and_profile_drift",
			DisplayName:       "Privilege and Profile Drift",
			FindingTypes:      []string{runtimeFindingPrivilegeDrift, runtimeFindingProfileDeviation},
			TriggerCondition:  "Privilege envelope or multi-signal runtime behavior deviates from the expected workload profile.",
			EvidenceModel:     []string{"runtime drift evidence", "active findings", "desired-state privilege profile"},
			DefaultSeverity:   "high",
			ConfidenceModel:   "high for explicit privilege drift; medium for aggregate profile deviation",
			DefaultNextAction: runtimeActionCaptureForensics,
			ApprovalModel:     recommendationApprovalHumanReview,
			RollbackPosture:   "tightened runtime restrictions should be reversible after stability is re-established",
			ForensicLinkage:   "forensic-first sequencing applies before stronger containment or recovery",
			ExecutionStatus:   runtimeRulePackStatusEnforceable,
		},
		{
			SchemaVersion:     runtimeRulePackSchemaVersion,
			PackID:            "filesystem_and_memory_execution",
			DisplayName:       "Filesystem and Memory Execution",
			FindingTypes:      []string{runtimeFindingFilesystemMutation, runtimeFindingMemoryExecAnomaly},
			TriggerCondition:  "Sensitive filesystem mutations or executable-memory anomalies appear in runtime evidence.",
			EvidenceModel:     []string{"runtime observation", "forensic context", "drift evidence"},
			DefaultSeverity:   "critical",
			ConfidenceModel:   "high when runtime evidence explicitly identifies the mutation or executable mapping",
			DefaultNextAction: runtimeActionCaptureForensics,
			ApprovalModel:     recommendationApprovalHumanReview,
			RollbackPosture:   "containment and recovery must preserve forensic state before reversal",
			ForensicLinkage:   "forensic capture is mandatory before destructive recovery",
			ExecutionStatus:   runtimeRulePackStatusEnforceable,
			Limitations: []string{
				"Memory anomaly handling is bounded to runtime evidence and executable mapping signals. It does not claim full live memory scanning coverage.",
			},
		},
	}
	sort.Slice(items, func(i, j int) bool { return items[i].PackID < items[j].PackID })
	return items
}

func runtimeRulePackByID(packID string) (runtimeRulePack, bool) {
	for _, item := range runtimeRulePackCatalog() {
		if item.PackID == strings.TrimSpace(packID) {
			return item, true
		}
	}
	return runtimeRulePack{}, false
}

func runtimeRulePackForFinding(findingType string) runtimeRulePack {
	for _, item := range runtimeRulePackCatalog() {
		if containsString(item.FindingTypes, findingType) {
			return item
		}
	}
	return runtimeRulePack{
		SchemaVersion:     runtimeRulePackSchemaVersion,
		PackID:            "runtime_integrity_fallback",
		DisplayName:       "Runtime Integrity Fallback",
		FindingTypes:      []string{findingType},
		TriggerCondition:  "A runtime integrity signal mapped into the fallback rule-pack path.",
		DefaultSeverity:   "medium",
		ConfidenceModel:   "bounded by the underlying evidence signal",
		DefaultNextAction: runtimeActionAlert,
		ApprovalModel:     recommendationApprovalAutoSafe,
		RollbackPosture:   "manual review required",
		ForensicLinkage:   "link forensic context when available",
		ExecutionStatus:   runtimeRulePackStatusDetectOnly,
	}
}

func buildRuntimePostureState(subject *runtimeSnapshotSubject, findings []runtimeIntegrityFinding, sbom runtimeSBOMVerificationResult, sandbox runtimeSandboxDecision, state runtimeIntegrityState) runtimePostureState {
	mismatches := []runtimePostureMismatch{}
	readiness := []string{}
	if subject.DesiredState != nil {
		readiness = append(readiness, "desired_state_present")
	}
	if subject.ActiveState != nil {
		readiness = append(readiness, "active_state_present")
	}
	if len(subject.ExpectedSigners) > 0 {
		readiness = append(readiness, "expected_signers_present")
	}
	if len(subject.TrustInputs) > 0 {
		readiness = append(readiness, "trust_inputs_present")
	}
	if len(findings) > 0 {
		readiness = append(readiness, "runtime_findings_present")
	}
	if subject.DesiredState == nil || subject.DesiredState.DesiredStateVerification != "verified" {
		mismatches = append(mismatches, runtimePostureMismatch{
			Code:     runtimeMismatchDesiredState,
			Severity: "medium",
			Summary:  "Desired state verification is not in verified posture for this workload.",
		})
	}
	if subject.DesiredState != nil && subject.DesiredState.DesiredStateVerification == "verified" && !containsString(mapKeys(subject.TrustInputs), "attestation_provenance") {
		mismatches = append(mismatches, runtimePostureMismatch{
			Code:        runtimeMismatchAttestation,
			Severity:    "high",
			Summary:     "Desired state is verified, but attestation-backed runtime trust input is missing from the current scope.",
			EvidenceRef: firstNonEmptyString(uniqueStrings(mapKeys(subject.EvidenceRefs))),
		})
	}
	if len(subject.ExpectedSigners) > 0 && !containsString(mapKeys(subject.TrustInputs), "signed_artifact") {
		mismatches = append(mismatches, runtimePostureMismatch{
			Code:        runtimeMismatchSigner,
			Severity:    "high",
			Summary:     "The workload expects signer-backed trust inputs, but no signed-artifact evidence is currently linked.",
			EvidenceRef: firstNonEmptyString(uniqueStrings(mapKeys(subject.EvidenceRefs))),
		})
	}
	if containsRuntimeFinding(findings, runtimeFindingIdentityDrift) || containsRuntimeFinding(findings, runtimeFindingAttestationMismatch) {
		mismatches = append(mismatches, runtimePostureMismatch{
			Code:     runtimeMismatchIdentity,
			Severity: "high",
			Summary:  "Runtime identity posture diverges from the expected signer, service-account, or attestation-linked trust path.",
		})
	}
	if sbom.Status == runtimeSBOMStatusDrift {
		mismatches = append(mismatches, runtimePostureMismatch{
			Code:        runtimeMismatchSBOM,
			Severity:    "high",
			Summary:     "Observed runtime digest no longer matches the approved SBOM-linked artifact.",
			EvidenceRef: firstNonEmptyString(sbom.UnexpectedArtifactRefs),
		})
	}
	if hasCriticalRuntimeFinding(findings) {
		mismatches = append(mismatches, runtimePostureMismatch{
			Code:     runtimeMismatchCriticalFindings,
			Severity: "critical",
			Summary:  "One or more critical runtime findings remain active in the current scope.",
		})
	}
	scheduling := runtimeSchedulingGuidance{
		Decision:          runtimeSchedulingAllowStandard,
		ApprovalMode:      recommendationApprovalAutoSafe,
		ReasonCodes:       uniqueStrings(append([]string{}, sandbox.ReasonCodes...)),
		RecommendedAction: runtimeActionAlert,
	}
	switch sandbox.AssignedSandboxClass {
	case runtimeSandboxClassIsolatedReview:
		scheduling.Decision = runtimeSchedulingIsolatedReview
		scheduling.ApprovalMode = recommendationApprovalHumanReview
	case runtimeSandboxClassHardened, runtimeSandboxClassRestricted:
		scheduling.Decision = runtimeSchedulingRestricted
		scheduling.ApprovalMode = recommendationApprovalHumanReview
	}
	if len(findings) > 0 {
		scheduling.RecommendedAction = findings[0].RecommendedAction
	}
	return runtimePostureState{
		SchemaVersion:      runtimePostureSchemaVersion,
		SubjectRef:         subject.SubjectRef,
		RuntimeModuleReady: subject.DesiredState != nil && subject.ActiveState != nil,
		ReadinessSignals:   uniqueStrings(readiness),
		ExpectedTrustState: runtimePostureTrustState{
			DesiredStateVerification: func() string {
				if subject.DesiredState == nil {
					return ""
				}
				return subject.DesiredState.DesiredStateVerification
			}(),
			AttestationInputs: uniqueStrings(mapKeys(subject.TrustInputs)),
			ExpectedSigners:   uniqueStrings(subject.ExpectedSigners),
			SBOMStatus:        sbom.Status,
			SandboxClass:      sandbox.AssignedSandboxClass,
		},
		ActualTrustState: runtimePostureTrustState{
			DesiredStateVerification:  runtimeFindingDesiredStateVerification(subject),
			AttestationInputs:         uniqueStrings(mapKeys(subject.TrustInputs)),
			ExpectedSigners:           uniqueStrings(subject.ExpectedSigners),
			SBOMStatus:                sbom.Status,
			SandboxClass:              state.CurrentSandboxClass,
			CurrentEnforcementPosture: state.CurrentEnforcementPosture,
		},
		Mismatches:         uniqueRuntimePostureMismatches(mismatches),
		SchedulingGuidance: scheduling,
		EvidenceRefs:       uniqueStrings(mapKeys(subject.EvidenceRefs)),
		LastVerifiedAt:     state.LastVerifiedAt,
		Limitations: []string{
			"Posture state is workload-scoped and derives from canonical desired-state, active-state, artifact, and runtime evidence already present in ChangeLock.",
		},
	}
}

func uniqueRuntimePostureMismatches(values []runtimePostureMismatch) []runtimePostureMismatch {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]runtimePostureMismatch{}
	for _, item := range values {
		if strings.TrimSpace(item.Code) == "" {
			continue
		}
		if current, ok := seen[item.Code]; !ok || runtimeSeverityRank(item.Severity) > runtimeSeverityRank(current.Severity) {
			seen[item.Code] = item
		}
	}
	items := make([]runtimePostureMismatch, 0, len(seen))
	for _, item := range seen {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Severity == items[j].Severity {
			return items[i].Code < items[j].Code
		}
		return runtimeSeverityRank(items[i].Severity) > runtimeSeverityRank(items[j].Severity)
	})
	return items
}

func firstNonEmptyString(values []string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func runtimeExplainabilityForFinding(finding runtimeIntegrityFinding, sbom runtimeSBOMVerificationResult, sandbox runtimeSandboxDecision, state runtimeIntegrityState, topology *runtimeEnforcementTopologyContext, expectedSigners []string) runtimeExplainability {
	return runtimeExplainability{
		SchemaVersion: runtimeExplainabilitySchema,
		Trigger:       finding.Summary,
		TriggerSource: finding.FindingType,
		EvidenceRefs:  append([]string{}, finding.EvidenceRefs...),
		TrustContext: runtimeExplainabilityTrustContext{
			DesiredStateVerification: firstNonEmpty(finding.Explainability.TrustContext.DesiredStateVerification, ""),
			AttestationInputs:        append([]string{}, sandbox.AttestationInputs...),
			ExpectedSigners:          uniqueStrings(expectedSigners),
			SandboxClass:             sandbox.AssignedSandboxClass,
			SBOMStatus:               sbom.Status,
			CurrentPosture:           state.CurrentEnforcementPosture,
			IdentityStatus:           state.IdentityStatus,
		},
		ResponsePath: runtimeExplainabilityResponsePath{
			PolicyRef:         finding.MatchedPolicyRule,
			RecommendedAction: finding.RecommendedAction,
			ApprovalMode:      runtimeApprovalModeForAction(finding.RecommendedAction),
			ApprovalRequired:  runtimeApprovalModeForAction(finding.RecommendedAction) == recommendationApprovalHumanReview,
			RollbackHints:     runtimeRollbackHints(finding.RecommendedAction),
		},
		Forensics: runtimeExplainabilityForensics{
			ForensicContextURI: finding.ForensicContextURI,
			ReadbackRefs:       append([]advisoryReadbackRef(nil), finding.ReadbackRefs...),
		},
		Topology:    runtimeTopologyExplainability(topology),
		NextSteps:   runtimeNextStepsForAction(finding.RecommendedAction),
		Limitations: append([]string{}, finding.Limitations...),
	}
}

func runtimeExplainabilityForDecision(decision runtimeEnforcementDecision, finding runtimeIntegrityFinding, sandbox runtimeSandboxDecision, state runtimeIntegrityState, expectedSigners []string, sbom runtimeSBOMVerificationResult) runtimeExplainability {
	return runtimeExplainability{
		SchemaVersion: runtimeExplainabilitySchema,
		Trigger:       finding.Summary,
		TriggerSource: finding.FindingType,
		EvidenceRefs:  append([]string{}, decision.EvidenceRefs...),
		TrustContext: runtimeExplainabilityTrustContext{
			DesiredStateVerification: finding.Explainability.TrustContext.DesiredStateVerification,
			AttestationInputs:        append([]string{}, sandbox.AttestationInputs...),
			ExpectedSigners:          uniqueStrings(expectedSigners),
			SandboxClass:             sandbox.AssignedSandboxClass,
			SBOMStatus:               sbom.Status,
			CurrentPosture:           state.CurrentEnforcementPosture,
			IdentityStatus:           state.IdentityStatus,
		},
		ResponsePath: runtimeExplainabilityResponsePath{
			PolicyRef:         decision.PolicyRef,
			RecommendedAction: finding.RecommendedAction,
			SelectedAction:    decision.Action,
			ApprovalMode:      decision.ApprovalMode,
			ApprovalRequired:  decision.ApprovalMode == recommendationApprovalHumanReview,
			RollbackHints:     runtimeRollbackHints(decision.Action),
		},
		Forensics: runtimeExplainabilityForensics{
			ForensicContextURI: decision.ForensicContextURI,
			ReadbackRefs:       append([]advisoryReadbackRef(nil), finding.ReadbackRefs...),
		},
		Topology:    runtimeTopologyExplainability(decision.TopologyContext),
		NextSteps:   runtimeNextStepsForAction(decision.Action),
		Limitations: append([]string{}, decision.Limitations...),
	}
}

func runtimeTopologyExplainability(topology *runtimeEnforcementTopologyContext) []string {
	if topology == nil {
		return nil
	}
	values := []string{}
	if topology.PrimaryService != "" {
		values = append(values, "primary_service:"+topology.PrimaryService)
	}
	if topology.BlastRadiusScore > 0 {
		values = append(values, "blast_radius_score:"+strconv.Itoa(topology.BlastRadiusScore))
	}
	if topology.CriticalReachCount > 0 {
		values = append(values, "critical_reach_count:"+strconv.Itoa(topology.CriticalReachCount))
	}
	values = append(values, topology.TopRiskPathSummaries...)
	return uniqueStrings(values)
}

func runtimeApprovalModeForAction(action string) string {
	switch action {
	case runtimeActionApplyNetworkIsolation, runtimeActionRestartTrusted, runtimeActionRecommendQuarantine:
		return recommendationApprovalHumanReview
	default:
		return recommendationApprovalAutoSafe
	}
}

func runtimeRollbackHints(action string) []string {
	switch action {
	case runtimeActionApplyNetworkIsolation, runtimeActionRecommendQuarantine:
		return []string{
			"Keep containment bounded with TTL and remove temporary restrictions only after clean verification confirms the trigger is no longer active.",
		}
	case runtimeActionRestartTrusted:
		return []string{
			"Trusted recovery should only clear active restrictions after the workload is back on an approved digest and the original finding no longer reproduces.",
		}
	default:
		return []string{
			"Review the linked evidence and operator guidance before widening runtime posture.",
		}
	}
}

func runtimeNextStepsForAction(action string) []string {
	switch action {
	case runtimeActionCaptureForensics:
		return []string{
			"Preserve forensic state before choosing stronger containment or recovery.",
			"Review the linked evidence and confirm whether the signal reproduces against the same workload scope.",
		}
	case runtimeActionApplyNetworkIsolation, runtimeActionRecommendQuarantine:
		return []string{
			"Keep containment bounded and review blast-radius impact before widening isolation.",
			"Confirm the underlying runtime signal no longer reproduces before rollback.",
		}
	case runtimeActionRestartTrusted:
		return []string{
			"Verify the workload returns on an approved digest and that the original runtime mismatch no longer appears.",
			"Only clear temporary restrictions after clean verification succeeds.",
		}
	default:
		return []string{
			"Review the linked runtime finding and operator guidance before changing runtime posture.",
		}
	}
}

func runtimeFindingRulePackID(findingType string) string {
	return runtimeRulePackForFinding(findingType).PackID
}

func runtimeDerivedFinding(rulePackID, findingType, severity, subjectRef, profileRef, summary, policyRule, recommendedAction, forensicURI, desiredStateVerification string, evidenceRefs []string, readbackRefs []advisoryReadbackRef, limitations []string) runtimeIntegrityFinding {
	return runtimeIntegrityFinding{
		FindingID:          recommendationID("runtime-finding", subjectRef, findingType),
		RulePackRef:        rulePackID,
		FindingType:        findingType,
		Severity:           severity,
		SubjectRef:         subjectRef,
		ProfileRef:         profileRef,
		Status:             runtimeFindingStatusActive,
		Summary:            summary,
		MatchedPolicyRule:  policyRule,
		EvidenceRefs:       uniqueStrings(evidenceRefs),
		ReadbackRefs:       readbackRefs,
		ForensicContextURI: forensicURI,
		Confidence:         runtimeConfidenceMedium,
		RecommendedAction:  recommendedAction,
		Limitations:        uniqueStrings(limitations),
		Explainability: runtimeExplainability{
			SchemaVersion: runtimeExplainabilitySchema,
			TrustContext: runtimeExplainabilityTrustContext{
				DesiredStateVerification: desiredStateVerification,
			},
		},
	}
}

func runtimeFindingDesiredStateVerification(subject *runtimeSnapshotSubject) string {
	if subject.DesiredState == nil {
		return ""
	}
	return subject.DesiredState.DesiredStateVerification
}

func runtimeDerivedContextFindings(subject *runtimeSnapshotSubject, profile runtimeIntegrityProfile, baseFindings []runtimeIntegrityFinding, readbackRefs []advisoryReadbackRef, topology *runtimeEnforcementTopologyContext) []runtimeIntegrityFinding {
	items := []runtimeIntegrityFinding{}
	evidenceRefs := uniqueStrings(mapKeys(subject.EvidenceRefs))
	desiredStateVerification := runtimeFindingDesiredStateVerification(subject)
	forensicURI := runtimeForensicContextURI(runtimeIntegrityFilter{TenantID: subject.TenantID, Environment: subject.Environment}, subject.SubjectRef, runtimeLastVerified(subject))
	trustInputs := uniqueStrings(mapKeys(subject.TrustInputs))
	if desiredStateVerification == "verified" && !containsString(trustInputs, "attestation_provenance") {
		items = append(items, runtimeDerivedFinding(
			"runtime_identity_and_attestation",
			runtimeFindingAttestationMismatch,
			"high",
			subject.SubjectRef,
			profile.ProfileID,
			"Desired state is verified, but runtime trust evidence no longer proves the expected attestation posture for this workload.",
			"runtime_attestation_linkage",
			runtimeActionCaptureForensics,
			forensicURI,
			desiredStateVerification,
			evidenceRefs,
			readbackRefs,
			[]string{
				"This finding marks an evidence-backed trust mismatch. It does not claim live scheduler or node attestation unless explicit substrate evidence is present.",
			},
		))
	}
	if topology != nil && topology.BlastRadiusScore >= 60 && (containsRuntimeFinding(baseFindings, runtimeFindingOutboundDrift) || containsRuntimeFinding(baseFindings, runtimeFindingUnknownBinaryExec) || containsRuntimeFinding(baseFindings, runtimeFindingUnsignedBinaryExec) || containsRuntimeFinding(baseFindings, runtimeFindingMemoryExecAnomaly)) {
		items = append(items, runtimeDerivedFinding(
			"outbound_and_topology_expansion",
			runtimeFindingTopologyExpansion,
			"high",
			subject.SubjectRef,
			profile.ProfileID,
			"Runtime signals coincide with widened blast-radius context, so the workload now presents a larger containment surface than the baseline profile.",
			"topology_aware_containment",
			runtimeActionRecommendQuarantine,
			forensicURI,
			desiredStateVerification,
			append(evidenceRefs, topology.TopRiskPathSummaries...),
			readbackRefs,
			[]string{
				"Topology expansion is advisory and derived from the current topology model; it does not claim full live network truth.",
			},
		))
	}
	activeTypes := map[string]struct{}{}
	for _, item := range baseFindings {
		if item.Status == runtimeFindingStatusRemediated {
			continue
		}
		activeTypes[item.FindingType] = struct{}{}
	}
	if len(activeTypes) >= 2 {
		severity := "medium"
		if hasHighRuntimeFinding(baseFindings) {
			severity = "high"
		}
		items = append(items, runtimeDerivedFinding(
			"privilege_and_profile_drift",
			runtimeFindingProfileDeviation,
			severity,
			subject.SubjectRef,
			profile.ProfileID,
			"Multiple active runtime signals now deviate from the expected profile, indicating a broader behavioral shift rather than one isolated mismatch.",
			"runtime_profile_behavior_baseline",
			runtimeActionCaptureForensics,
			forensicURI,
			desiredStateVerification,
			evidenceRefs,
			readbackRefs,
			[]string{
				"Profile deviation is an aggregate signal over multiple active findings and remains bounded to evidence already present in scope.",
			},
		))
	}
	return items
}

func runtimeFindingRulePackRef(findingType string) string {
	return runtimeRulePackForFinding(findingType).PackID
}

func ensureRuntimeRulePackExists(findingType string) error {
	if runtimeFindingRulePackRef(findingType) == "" {
		return errors.New("runtime rule pack missing")
	}
	return nil
}
