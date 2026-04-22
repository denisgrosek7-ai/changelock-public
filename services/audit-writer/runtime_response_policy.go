package main

import (
	"net/http"
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	runtimeResponsePolicySchemaVersion = "3a.runtime_response_policy.v1"

	runtimeResponseModeBoundedAutonomous = "bounded_autonomous"
	runtimeResponseModeApprovalGated     = "approval_gated"

	runtimeSafetyLimitAdvisoryOnly    = "runtime_response_safety.v1:advisory_only"
	runtimeSafetyLimitWorkloadOnly    = "runtime_response_safety.v1:workload_only"
	runtimeSafetyLimitOperatorReview  = "runtime_response_safety.v1:operator_review"
	runtimeSafetyLimitTrustedRecovery = "runtime_response_safety.v1:trusted_recovery"
)

type runtimeResponseActionPolicy struct {
	Action            string `json:"action"`
	ResponseMode      string `json:"response_mode"`
	ApprovalRequired  bool   `json:"approval_required"`
	ConfidenceLevel   string `json:"confidence_level"`
	ForensicFirst     bool   `json:"forensic_first"`
	RollbackRequired  bool   `json:"rollback_required"`
	TTL               string `json:"ttl"`
	LeastInvasiveRank int    `json:"least_invasive_rank"`
	SafetyLimitRef    string `json:"safety_limit_ref"`
}

type runtimeResponseConfidenceThreshold struct {
	Action            string `json:"action"`
	MinimumConfidence string `json:"minimum_confidence"`
}

type runtimeResponseTTL struct {
	Action string `json:"action"`
	TTL    string `json:"ttl"`
}

type runtimeBlastRadiusSafetyLimit struct {
	SafetyLimitRef                string   `json:"safety_limit_ref"`
	AppliesWhen                   string   `json:"applies_when"`
	MaxAutonomousBlastRadiusScore int      `json:"max_autonomous_blast_radius_score,omitempty"`
	ContainmentScope              string   `json:"containment_scope"`
	ApprovalMode                  string   `json:"approval_mode"`
	Notes                         []string `json:"notes,omitempty"`
}

type runtimeForensicFirstPolicy struct {
	SnapshotFirstRequiredFor             []string `json:"snapshot_first_required_for"`
	SoftIsolationAllowedBeforeRecovery   []string `json:"soft_isolation_allowed_before_recovery"`
	RecoveryRequiresCleanVerificationFor []string `json:"recovery_requires_clean_verification_for"`
	Limitations                          []string `json:"limitations,omitempty"`
}

type runtimeResponsePolicyResponse struct {
	SchemaVersion            string                               `json:"schema_version"`
	DefaultResponseMode      string                               `json:"default_response_mode"`
	LeastInvasiveFirst       bool                                 `json:"least_invasive_first"`
	AutonomousActions        []string                             `json:"autonomous_actions"`
	ApprovalRequiredActions  []string                             `json:"approval_required_actions"`
	ConfidenceThresholds     []runtimeResponseConfidenceThreshold `json:"confidence_thresholds"`
	ForensicFirstRequiredFor []string                             `json:"forensic_first_required_for"`
	RollbackRequiredFor      []string                             `json:"rollback_required_for"`
	DefaultTTL               []runtimeResponseTTL                 `json:"default_ttl"`
	BlastRadiusSafetyLimits  []runtimeBlastRadiusSafetyLimit      `json:"blast_radius_safety_limits"`
	ActionPolicies           []runtimeResponseActionPolicy        `json:"action_policies"`
	ForensicFirstPolicy      runtimeForensicFirstPolicy           `json:"forensic_first_policy"`
	Limitations              []string                             `json:"limitations,omitempty"`
}

func (s server) runtimeResponsePolicyHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, runtimeResponsePolicyCatalog())
}

func runtimeResponsePolicyCatalog() runtimeResponsePolicyResponse {
	actionPolicies := []runtimeResponseActionPolicy{
		runtimeResponseActionPolicyForContract(runtimeActionObserveOnly),
		runtimeResponseActionPolicyForContract(runtimeActionAlert),
		runtimeResponseActionPolicyForContract(runtimeActionCaptureForensics),
		runtimeResponseActionPolicyForContract(runtimeActionRecommendQuarantine),
		runtimeResponseActionPolicyForContract(runtimeActionApplyNetworkIsolation),
		runtimeResponseActionPolicyForContract(runtimeActionRestartTrusted),
	}
	sort.Slice(actionPolicies, func(i, j int) bool {
		if actionPolicies[i].LeastInvasiveRank == actionPolicies[j].LeastInvasiveRank {
			return actionPolicies[i].Action < actionPolicies[j].Action
		}
		return actionPolicies[i].LeastInvasiveRank < actionPolicies[j].LeastInvasiveRank
	})
	confidenceThresholds := []runtimeResponseConfidenceThreshold{}
	defaultTTLs := []runtimeResponseTTL{}
	autonomousActions := []string{}
	approvalActions := []string{}
	forensicRequired := []string{}
	rollbackRequired := []string{}
	for _, item := range actionPolicies {
		confidenceThresholds = append(confidenceThresholds, runtimeResponseConfidenceThreshold{
			Action:            item.Action,
			MinimumConfidence: item.ConfidenceLevel,
		})
		defaultTTLs = append(defaultTTLs, runtimeResponseTTL{
			Action: item.Action,
			TTL:    item.TTL,
		})
		if item.ApprovalRequired {
			approvalActions = append(approvalActions, item.Action)
		} else {
			autonomousActions = append(autonomousActions, item.Action)
		}
		if item.ForensicFirst {
			forensicRequired = append(forensicRequired, item.Action)
		}
		if item.RollbackRequired {
			rollbackRequired = append(rollbackRequired, item.Action)
		}
	}
	return runtimeResponsePolicyResponse{
		SchemaVersion:           runtimeResponsePolicySchemaVersion,
		DefaultResponseMode:     "bounded_defense_loop",
		LeastInvasiveFirst:      true,
		AutonomousActions:       autonomousActions,
		ApprovalRequiredActions: approvalActions,
		ConfidenceThresholds:    confidenceThresholds,
		ForensicFirstRequiredFor: []string{
			runtimeFindingUnknownBinaryExec,
			runtimeFindingUnsignedBinaryExec,
			runtimeFindingPrivilegeDrift,
			runtimeFindingFilesystemMutation,
			runtimeFindingMemoryExecAnomaly,
			runtimeFindingAttestationMismatch,
			runtimeFindingSBOMMismatch,
			runtimeFindingProfileDeviation,
		},
		RollbackRequiredFor: rollbackRequired,
		DefaultTTL:          defaultTTLs,
		BlastRadiusSafetyLimits: []runtimeBlastRadiusSafetyLimit{
			{
				SafetyLimitRef:                runtimeSafetyLimitAdvisoryOnly,
				AppliesWhen:                   "observe_only, alert, or forensic capture does not widen containment scope",
				MaxAutonomousBlastRadiusScore: 100,
				ContainmentScope:              "advisory_only",
				ApprovalMode:                  recommendationApprovalAutoSafe,
				Notes: []string{
					"Advisory and forensic capture actions remain autonomous only while they do not expand workload containment.",
				},
			},
			{
				SafetyLimitRef:                runtimeSafetyLimitWorkloadOnly,
				AppliesWhen:                   "bounded workload-only containment with low or moderate blast radius",
				MaxAutonomousBlastRadiusScore: 59,
				ContainmentScope:              "workload_only",
				ApprovalMode:                  recommendationApprovalHumanReview,
				Notes: []string{
					"Containment must stay scoped to the affected workload and keep TTL plus rollback metadata attached.",
				},
			},
			{
				SafetyLimitRef:                runtimeSafetyLimitOperatorReview,
				AppliesWhen:                   "service reach, critical topology paths, or elevated blast radius require manual review before broader restriction",
				MaxAutonomousBlastRadiusScore: 0,
				ContainmentScope:              "service_only_or_review",
				ApprovalMode:                  recommendationApprovalHumanReview,
				Notes: []string{
					"Runtime response must not widen from bounded workload containment into larger service impact without operator approval.",
				},
			},
			{
				SafetyLimitRef:                runtimeSafetyLimitTrustedRecovery,
				AppliesWhen:                   "trusted recovery or restart changes execution state after verification checks",
				MaxAutonomousBlastRadiusScore: 0,
				ContainmentScope:              "trusted_recovery",
				ApprovalMode:                  recommendationApprovalHumanReview,
				Notes: []string{
					"Trusted recovery stays blocked until clean verification succeeds and temporary restrictions are ready to clear.",
				},
			},
		},
		ActionPolicies: actionPolicies,
		ForensicFirstPolicy: runtimeForensicFirstPolicy{
			SnapshotFirstRequiredFor: []string{
				runtimeFindingUnknownBinaryExec,
				runtimeFindingUnsignedBinaryExec,
				runtimeFindingPrivilegeDrift,
				runtimeFindingFilesystemMutation,
				runtimeFindingMemoryExecAnomaly,
				runtimeFindingAttestationMismatch,
				runtimeFindingSBOMMismatch,
				runtimeFindingProfileDeviation,
			},
			SoftIsolationAllowedBeforeRecovery: []string{
				runtimeFindingOutboundDrift,
				runtimeFindingTopologyExpansion,
				runtimeFindingUnknownBinaryExec,
				runtimeFindingUnsignedBinaryExec,
			},
			RecoveryRequiresCleanVerificationFor: []string{
				runtimeActionRestartTrusted,
				hardeningModeTrustedRecovery,
			},
			Limitations: []string{
				"Forensic-first policy is bounded to the runtime evidence and response layers already present in ChangeLock; it does not imply full memory acquisition or kernel-wide snapshot capability for every substrate.",
			},
		},
		Limitations: []string{
			"Runtime response policy documents bounded response semantics, approval gates, TTLs, rollback posture, and blast-radius safety references for the current runtime and hardening slices.",
			"Policy metadata does not claim unrestricted autonomous remediation, universal pre-execution blocking, or coverage beyond the evidence-backed runtime surfaces already implemented.",
		},
	}
}

func runtimeResponseActionPolicyForContract(action string) runtimeResponseActionPolicy {
	approvalMode := runtimeApprovalModeForAction(action)
	return runtimeResponseActionPolicy{
		Action:            action,
		ResponseMode:      runtimeResponseModeForApprovalMode(approvalMode),
		ApprovalRequired:  approvalMode == recommendationApprovalHumanReview,
		ConfidenceLevel:   runtimeConfidenceThresholdForAction(action),
		ForensicFirst:     runtimeActionRequiresForensicFirst(action, runtimeIntegrityFinding{}),
		RollbackRequired:  runtimeRollbackRequired(action),
		TTL:               runtimeTTLForAction(action),
		LeastInvasiveRank: runtimeLeastInvasiveRank(action),
		SafetyLimitRef:    runtimeSafetyLimitRef(action, nil, approvalMode),
	}
}

func runtimeResponseModeForApprovalMode(approvalMode string) string {
	if approvalMode == recommendationApprovalHumanReview {
		return runtimeResponseModeApprovalGated
	}
	return runtimeResponseModeBoundedAutonomous
}

func runtimeConfidenceThresholdForAction(action string) string {
	switch action {
	case runtimeActionObserveOnly, runtimeActionAlert:
		return runtimeConfidenceLow
	case runtimeActionCaptureForensics, runtimeActionRecommendQuarantine:
		return runtimeConfidenceMedium
	case runtimeActionApplyNetworkIsolation, runtimeActionRestartTrusted:
		return runtimeConfidenceHigh
	default:
		return runtimeConfidenceMedium
	}
}

func runtimeTTLForAction(action string) string {
	switch action {
	case runtimeActionCaptureForensics:
		return hardeningModeForensicTTL
	case runtimeActionRecommendQuarantine, runtimeActionApplyNetworkIsolation:
		return hardeningModeSoftIsolationTTL
	case runtimeActionRestartTrusted:
		return hardeningModeRecoveryTTL
	default:
		return "0s"
	}
}

func runtimeRollbackRequired(action string) bool {
	switch action {
	case runtimeActionRecommendQuarantine, runtimeActionApplyNetworkIsolation, runtimeActionRestartTrusted:
		return true
	default:
		return false
	}
}

func runtimeLeastInvasiveRank(action string) int {
	switch action {
	case runtimeActionObserveOnly:
		return 1
	case runtimeActionAlert:
		return 2
	case runtimeActionCaptureForensics:
		return 3
	case runtimeActionRecommendQuarantine:
		return 4
	case runtimeActionApplyNetworkIsolation:
		return 5
	case runtimeActionRestartTrusted:
		return 6
	case runtimeActionEscalateManualReview:
		return 7
	default:
		return 50
	}
}

func runtimeActionRequiresForensicFirst(action string, finding runtimeIntegrityFinding) bool {
	if action == runtimeActionCaptureForensics {
		return true
	}
	if action == runtimeActionRestartTrusted {
		return true
	}
	switch action {
	case runtimeActionApplyNetworkIsolation, runtimeActionRecommendQuarantine, runtimeActionRestartTrusted:
		switch finding.FindingType {
		case runtimeFindingUnknownBinaryExec,
			runtimeFindingUnsignedBinaryExec,
			runtimeFindingPrivilegeDrift,
			runtimeFindingFilesystemMutation,
			runtimeFindingMemoryExecAnomaly,
			runtimeFindingAttestationMismatch,
			runtimeFindingSBOMMismatch,
			runtimeFindingProfileDeviation:
			return true
		}
		if runtimeSeverityRank(finding.Severity) >= runtimeSeverityRank("high") {
			return true
		}
	}
	return false
}

func runtimeSafetyLimitRef(action string, topology *runtimeEnforcementTopologyContext, approvalMode string) string {
	switch action {
	case runtimeActionRestartTrusted:
		return runtimeSafetyLimitTrustedRecovery
	case runtimeActionObserveOnly, runtimeActionAlert, runtimeActionCaptureForensics:
		return runtimeSafetyLimitAdvisoryOnly
	}
	if topology != nil && topology.BlastRadiusScore >= 60 {
		return runtimeSafetyLimitOperatorReview
	}
	if approvalMode == recommendationApprovalHumanReview {
		return runtimeSafetyLimitWorkloadOnly
	}
	return runtimeSafetyLimitAdvisoryOnly
}

func hardeningLeastInvasiveRank(actions []string) int {
	rank := 0
	for _, action := range actions {
		candidate := hardeningActionRank(action)
		if candidate <= 0 {
			continue
		}
		if rank == 0 || candidate < rank {
			rank = candidate
		}
	}
	if rank == 0 {
		return 50
	}
	return rank
}

func hardeningActionRank(action string) int {
	switch strings.TrimSpace(action) {
	case hardeningActionRequestForensics:
		return 1
	case hardeningActionRollbackRestrictions:
		return 2
	case hardeningActionRequireHumanReview:
		return 3
	case hardeningActionApplyNetworkQuarantine:
		return 4
	case hardeningActionRemoveFromTraffic:
		return 5
	case hardeningActionDivertIngress:
		return 6
	case hardeningActionTightenRuntimeProfile, hardeningActionBlockExecClass:
		return 7
	case hardeningActionRestartTrusted:
		return 8
	default:
		return 50
	}
}

func hardeningSafetyLimitRef(assessment hardeningAssessment, decision hardeningPolicyDecision) string {
	if containsString(decision.AllowedActions, hardeningActionRestartTrusted) || assessment.RecommendedHardeningClass == hardeningModeTrustedRecovery {
		return runtimeSafetyLimitTrustedRecovery
	}
	if assessment.BlastRadiusScore >= 60 || assessment.Criticality == "critical" {
		return runtimeSafetyLimitOperatorReview
	}
	if decision.ApprovalMode == recommendationApprovalHumanReview {
		return runtimeSafetyLimitWorkloadOnly
	}
	return runtimeSafetyLimitAdvisoryOnly
}
