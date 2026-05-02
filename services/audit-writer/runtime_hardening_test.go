package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
)

func TestRuntimeHardeningEvaluateApplyRollbackRecommendationsAndHandoff(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	findingsReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/findings?tenant_id=acme&environment=prod&limit=20", nil)
	findingsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	findingsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(findingsRec, findingsReq)
	if findingsRec.Code != http.StatusOK {
		t.Fatalf("expected runtime findings 200, got %d: %s", findingsRec.Code, findingsRec.Body.String())
	}

	var findings runtimeFindingsResponse
	if err := json.NewDecoder(findingsRec.Body).Decode(&findings); err != nil {
		t.Fatalf("decode runtime findings: %v", err)
	}
	binaryFinding := findRuntimeFinding(t, findings.Items, runtimeFindingUnknownBinaryExec, "edge-gateway")

	evaluateReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/evaluate?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`"}`),
	)
	evaluateReq.Header.Set("Authorization", "Bearer operator-demo-token")
	evaluateReq.Header.Set("Content-Type", "application/json")
	evaluateRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(evaluateRec, evaluateReq)
	if evaluateRec.Code != http.StatusOK {
		t.Fatalf("expected hardening evaluate 200, got %d: %s", evaluateRec.Code, evaluateRec.Body.String())
	}

	var evaluated hardeningEvaluationResponse
	if err := json.NewDecoder(evaluateRec.Body).Decode(&evaluated); err != nil {
		t.Fatalf("decode hardening evaluate: %v", err)
	}
	if !evaluated.Assessment.ForensicFirst || evaluated.Assessment.RecommendedHardeningClass != hardeningModeForensicPreserving {
		t.Fatalf("expected forensic-first preserving class, got %#v", evaluated.Assessment)
	}
	if evaluated.PolicyDecision.ApprovalMode != recommendationApprovalHumanReview {
		t.Fatalf("expected approval-required policy decision, got %#v", evaluated.PolicyDecision)
	}
	if len(evaluated.Actions) < 2 || evaluated.Actions[0].ActionType != hardeningActionRequestForensics {
		t.Fatalf("expected forensic-first action ordering, got %#v", evaluated.Actions)
	}

	forensicFirstReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/forensic-first?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`"}`),
	)
	forensicFirstReq.Header.Set("Authorization", "Bearer operator-demo-token")
	forensicFirstReq.Header.Set("Content-Type", "application/json")
	forensicFirstRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(forensicFirstRec, forensicFirstReq)
	if forensicFirstRec.Code != http.StatusOK {
		t.Fatalf("expected hardening forensic-first 200, got %d: %s", forensicFirstRec.Code, forensicFirstRec.Body.String())
	}

	var pending hardeningExecutionResponse
	if err := json.NewDecoder(forensicFirstRec.Body).Decode(&pending); err != nil {
		t.Fatalf("decode hardening forensic-first response: %v", err)
	}
	if pending.Execution.ExecutionResult != "forensic_snapshot_requested_containment_pending_approval" {
		t.Fatalf("expected forensic snapshot then pending containment, got %#v", pending.Execution)
	}
	if len(pending.Execution.ForensicRefs) == 0 || pending.Posture.CurrentMode != hardeningModePendingApproval {
		t.Fatalf("expected pending posture with forensic lineage, got %#v", pending)
	}

	quarantineReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/quarantine?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-9k-1"}`),
	)
	quarantineReq.Header.Set("Authorization", "Bearer operator-demo-token")
	quarantineReq.Header.Set("Content-Type", "application/json")
	quarantineRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(quarantineRec, quarantineReq)
	if quarantineRec.Code != http.StatusOK {
		t.Fatalf("expected hardening quarantine 200, got %d: %s", quarantineRec.Code, quarantineRec.Body.String())
	}

	var applied hardeningExecutionResponse
	if err := json.NewDecoder(quarantineRec.Body).Decode(&applied); err != nil {
		t.Fatalf("decode hardening quarantine response: %v", err)
	}
	if applied.Execution.ExecutionResult != "soft_isolation_applied" {
		t.Fatalf("expected soft isolation result, got %#v", applied.Execution)
	}
	if !containsString(applied.Posture.ActiveRestrictions, hardeningActionApplyNetworkQuarantine) || !applied.Posture.RollbackReady {
		t.Fatalf("expected active quarantine restriction with rollback readiness, got %#v", applied.Posture)
	}

	postureReq := httptest.NewRequest(http.MethodGet, "/v1/hardening/posture?tenant_id=acme&environment=prod&limit=10", nil)
	postureReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	postureRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(postureRec, postureReq)
	if postureRec.Code != http.StatusOK {
		t.Fatalf("expected hardening posture 200, got %d: %s", postureRec.Code, postureRec.Body.String())
	}

	var posture hardeningPostureResponse
	if err := json.NewDecoder(postureRec.Body).Decode(&posture); err != nil {
		t.Fatalf("decode hardening posture: %v", err)
	}
	edgePosture := findDefensePosture(t, posture.Items, "edge-gateway")
	if edgePosture.CurrentMode != hardeningModeSoftIsolation {
		t.Fatalf("expected soft isolation posture for edge-gateway, got %#v", edgePosture)
	}

	recommendationReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&environment=prod&source_type=hardening_signal", nil)
	recommendationReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	recommendationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recommendationRec, recommendationReq)
	if recommendationRec.Code != http.StatusOK {
		t.Fatalf("expected hardening recommendations 200, got %d: %s", recommendationRec.Code, recommendationRec.Body.String())
	}

	var recommendations recommendationListResponse
	if err := json.NewDecoder(recommendationRec.Body).Decode(&recommendations); err != nil {
		t.Fatalf("decode hardening recommendations: %v", err)
	}
	if len(recommendations.Recommendations) == 0 || recommendations.Recommendations[0].SourceType != "hardening_signal" {
		t.Fatalf("expected hardening_signal recommendation, got %#v", recommendations)
	}

	rollbackReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/rollback?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"execution_id":"`+applied.Execution.ExecutionID+`"}`),
	)
	rollbackReq.Header.Set("Authorization", "Bearer operator-demo-token")
	rollbackReq.Header.Set("Content-Type", "application/json")
	rollbackRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rollbackRec, rollbackReq)
	if rollbackRec.Code != http.StatusOK {
		t.Fatalf("expected hardening rollback 200, got %d: %s", rollbackRec.Code, rollbackRec.Body.String())
	}

	var rollback hardeningExecutionResponse
	if err := json.NewDecoder(rollbackRec.Body).Decode(&rollback); err != nil {
		t.Fatalf("decode hardening rollback response: %v", err)
	}
	if rollback.Execution.ExecutionResult != "rollback_applied" || rollback.Posture.CurrentMode != hardeningModeObserveOnly {
		t.Fatalf("expected rollback into observe-only posture, got %#v", rollback)
	}
	if rollback.PolicyDecision.LeastInvasiveRank != hardeningActionRank(hardeningActionRollbackRestrictions) || rollback.PolicyDecision.LeastInvasiveRank >= hardeningActionRank(hardeningActionApplyNetworkQuarantine) {
		t.Fatalf("expected rollback to have an explicit least-invasive rank ahead of quarantine, got %#v", rollback.PolicyDecision)
	}

	sealReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/seal?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"audience":"internal","include_hardening":true}`),
	)
	sealReq.Header.Set("Authorization", "Bearer operator-demo-token")
	sealReq.Header.Set("Content-Type", "application/json")
	sealRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(sealRec, sealReq)
	if sealRec.Code != http.StatusOK {
		t.Fatalf("expected hardening-inclusive handoff seal 200, got %d: %s", sealRec.Code, sealRec.Body.String())
	}

	var sealed handoffSealResponse
	if err := json.NewDecoder(sealRec.Body).Decode(&sealed); err != nil {
		t.Fatalf("decode hardening-inclusive handoff: %v", err)
	}
	if !handoffManifestHasArtifact(sealed.Manifest, "evidence/hardening_context.json") {
		t.Fatalf("expected hardening context artifact in sealed handoff, got %#v", sealed.Manifest.Artifacts)
	}
}

func TestHardeningActionRankIncludesRollbackRestrictions(t *testing.T) {
	rank := hardeningActionRank(hardeningActionRollbackRestrictions)
	if rank != 2 {
		t.Fatalf("expected rollback restrictions rank 2, got %d", rank)
	}
	if rank >= hardeningActionRank(hardeningActionApplyNetworkQuarantine) {
		t.Fatalf("expected rollback restrictions to rank ahead of quarantine, got rollback=%d quarantine=%d", rank, hardeningActionRank(hardeningActionApplyNetworkQuarantine))
	}
}

func TestBuildHardeningAssessmentEscalatesOnUnresolvedHighPriorityIncident(t *testing.T) {
	assessment := buildHardeningAssessment(
		runtimeIntegrityFinding{
			FindingID:   "finding-runtime-hardening-1",
			FindingType: runtimeFindingOutboundDrift,
			Severity:    "medium",
			SubjectRef:  "cluster/ns/workload",
		},
		runtimeWorkloadView{
			State: runtimeIntegrityState{CurrentSandboxClass: runtimeSandboxClassStandard},
		},
		[]investigationIncident{{
			State:    incidentStateOpen,
			Severity: "medium",
			Priority: "high",
		}},
		nil,
	)
	if assessment.Criticality != "critical" {
		t.Fatalf("expected unresolved high-priority incident to escalate criticality, got %#v", assessment)
	}
}

func TestRuntimeHardeningRecoveryRequiresCleanStateAndSupportsTrustedRecovery(t *testing.T) {
	fixture := forensicsTestFixture(t)

	findingsReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/findings?tenant_id=acme&environment=prod&limit=20", nil)
	findingsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	findingsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(findingsRec, findingsReq)
	if findingsRec.Code != http.StatusOK {
		t.Fatalf("expected runtime findings 200, got %d: %s", findingsRec.Code, findingsRec.Body.String())
	}

	var findings runtimeFindingsResponse
	if err := json.NewDecoder(findingsRec.Body).Decode(&findings); err != nil {
		t.Fatalf("decode runtime findings: %v", err)
	}
	binaryFinding := findRuntimeFinding(t, findings.Items, runtimeFindingUnknownBinaryExec, "edge-gateway")

	quarantineReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/quarantine?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-9k-2"}`),
	)
	quarantineReq.Header.Set("Authorization", "Bearer operator-demo-token")
	quarantineReq.Header.Set("Content-Type", "application/json")
	quarantineRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(quarantineRec, quarantineReq)
	if quarantineRec.Code != http.StatusOK {
		t.Fatalf("expected hardening quarantine 200, got %d: %s", quarantineRec.Code, quarantineRec.Body.String())
	}

	var applied hardeningExecutionResponse
	if err := json.NewDecoder(quarantineRec.Body).Decode(&applied); err != nil {
		t.Fatalf("decode hardening quarantine response: %v", err)
	}

	recoverDirtyReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/recover?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"execution_id":"`+applied.Execution.ExecutionID+`"}`),
	)
	recoverDirtyReq.Header.Set("Authorization", "Bearer operator-demo-token")
	recoverDirtyReq.Header.Set("Content-Type", "application/json")
	recoverDirtyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recoverDirtyRec, recoverDirtyReq)
	if recoverDirtyRec.Code != http.StatusConflict {
		t.Fatalf("expected recover to require clean verification on dirty subject, got %d: %s", recoverDirtyRec.Code, recoverDirtyRec.Body.String())
	}

	cleanSubjectRef := seedCleanHardeningSubject(t, fixture.store)
	recoverCleanReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/recover?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"subject_ref":"`+cleanSubjectRef+`"}`),
	)
	recoverCleanReq.Header.Set("Authorization", "Bearer operator-demo-token")
	recoverCleanReq.Header.Set("Content-Type", "application/json")
	recoverCleanRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recoverCleanRec, recoverCleanReq)
	if recoverCleanRec.Code != http.StatusOK {
		t.Fatalf("expected clean trusted recovery 200, got %d: %s", recoverCleanRec.Code, recoverCleanRec.Body.String())
	}

	var recovered hardeningExecutionResponse
	if err := json.NewDecoder(recoverCleanRec.Body).Decode(&recovered); err != nil {
		t.Fatalf("decode clean trusted recovery response: %v", err)
	}
	if recovered.Execution.ExecutionResult != "trusted_recovery_completed" || recovered.Posture.CurrentMode != hardeningModeObserveOnly {
		t.Fatalf("expected trusted recovery completion into observe-only posture, got %#v", recovered)
	}
}

func seedCleanHardeningSubject(t *testing.T, store audit.Store) string {
	t.Helper()
	subjectRef := runtimeSubjectRef("local", "acme-prod", "Deployment", "clean-worker")
	now := time.Now().UTC()
	events := []audit.Event{
		{
			RequestID:                "runtime-clean-desired",
			Timestamp:                now.Add(-2 * time.Hour),
			Component:                "runtime-agent",
			EventType:                audit.EventTypeRuntimeDesiredStateRecorded,
			Decision:                 audit.DecisionAllow,
			TenantID:                 "acme",
			ClusterID:                "local",
			Repo:                     "acme/platform-clean",
			Environment:              "prod",
			Namespace:                "acme-prod",
			WorkloadKind:             "Deployment",
			Workload:                 "clean-worker",
			ServiceAccount:           "clean-sa",
			Digest:                   "sha256:clean-v1",
			DesiredStateSourceRef:    "deploy:clean-worker:v1",
			DesiredStateApprovalID:   "runtime-approval-clean",
			DesiredStateVerification: "verified",
			Reasons:                  []string{"runtime desired state approved"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity:           "signer-clean",
					Issuer:                   "root-clean",
					AttestationPredicate:     "https://slsa.dev/provenance/v1",
					AttestationSubjectDigest: "sha256:clean-v1",
					SBOMHash:                 "sbom-clean-v1",
				},
				Runtime: &audit.RuntimeEvidence{
					ApprovedDigest:         "sha256:clean-v1",
					ServiceAccountExpected: "clean-sa",
					ApprovedContainers: []audit.RuntimeApprovedContainer{{
						Name:           "clean-worker",
						ApprovedDigest: "sha256:clean-v1",
						Runtime: audit.RuntimeSecurityConstraints{
							RunAsNonRoot:           true,
							ReadOnlyRootFilesystem: true,
							DropAllCapabilities:    true,
							SeccompRuntimeDefault:  true,
							DenyPrivileged:         true,
						},
					}},
				},
			},
		},
		{
			RequestID:                "runtime-clean-active",
			Timestamp:                now.Add(-1 * time.Hour),
			Component:                "runtime-agent",
			EventType:                audit.EventTypeRuntimeActiveStateObserved,
			Decision:                 audit.DecisionAllow,
			TenantID:                 "acme",
			ClusterID:                "local",
			Repo:                     "acme/platform-clean",
			Environment:              "prod",
			Namespace:                "acme-prod",
			WorkloadKind:             "Deployment",
			Workload:                 "clean-worker",
			ServiceAccount:           "clean-sa",
			Digest:                   "sha256:clean-v1",
			ReconciliationStatus:     "in_sync",
			DesiredStateSourceRef:    "deploy:clean-worker:v1",
			DesiredStateApprovalID:   "runtime-approval-clean",
			DesiredStateVerification: "verified",
			Reasons:                  []string{"observed healthy runtime state"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity:           "signer-clean",
					Issuer:                   "root-clean",
					AttestationPredicate:     "https://slsa.dev/provenance/v1",
					AttestationSubjectDigest: "sha256:clean-v1",
					SBOMHash:                 "sbom-clean-v1",
				},
				Runtime: &audit.RuntimeEvidence{
					ApprovedDigest:         "sha256:clean-v1",
					RunningDigest:          "sha256:clean-v1",
					ServiceAccountExpected: "clean-sa",
					ServiceAccountObserved: "clean-sa",
				},
			},
		},
		{
			RequestID:      "hardening-clean-record",
			Timestamp:      now.Add(-30 * time.Minute),
			Component:      hardeningComponent,
			EventType:      audit.EventTypeHardeningActionApplied,
			Decision:       audit.DecisionAllow,
			TenantID:       "acme",
			ClusterID:      "local",
			Repo:           "acme/platform-clean",
			Environment:    "prod",
			Namespace:      "acme-prod",
			WorkloadKind:   "Deployment",
			Workload:       "clean-worker",
			ServiceAccount: "clean-sa",
			Digest:         "sha256:clean-v1",
			DriftResult:    hardeningModeSoftIsolation,
			DriftSeverity:  "high",
			Reasons:        []string{"soft_isolation_applied", "seeded for trusted recovery test"},
			RuntimeIntegrity: canonicalJSONMust(hardeningEventPayload{
				Trigger: &hardeningTrigger{
					TriggerID:     recommendationID("hardening-trigger", subjectRef, "sbom_runtime_mismatch_signal"),
					SourceFinding: recommendationID("runtime-finding", subjectRef, runtimeFindingSBOMMismatch),
					TriggerType:   runtimeFindingSBOMMismatch,
					Timestamp:     now.Add(-30 * time.Minute),
					SubjectRef:    subjectRef,
					Severity:      "high",
					Confidence:    runtimeConfidenceHigh,
					EvidenceRefs:  []string{"hardening-clean-record", "sha256:clean-v1"},
				},
				Assessment: &hardeningAssessment{
					AssessmentID:              recommendationID("hardening-assessment", subjectRef, runtimeFindingSBOMMismatch),
					TriggerRef:                recommendationID("hardening-trigger", subjectRef, runtimeFindingSBOMMismatch),
					SubjectRef:                subjectRef,
					BlastRadiusScore:          12,
					Criticality:               "standard",
					CurrentSandboxClass:       runtimeSandboxClassRestricted,
					ForensicFirst:             false,
					RecommendedHardeningClass: hardeningModeSoftIsolation,
					ReasonCodes:               []string{"seeded_clean_subject"},
				},
				PolicyDecision: &hardeningPolicyDecision{
					DecisionID:          recommendationID("hardening-policy", subjectRef, runtimeFindingSBOMMismatch),
					AssessmentRef:       recommendationID("hardening-assessment", subjectRef, runtimeFindingSBOMMismatch),
					PolicyRef:           "runtime_closed_loop_hardening.v1:seed",
					AllowedActions:      []string{hardeningActionApplyNetworkQuarantine},
					ApprovalMode:        recommendationApprovalAutoSafe,
					TTL:                 hardeningModeSoftIsolationTTL,
					RollbackRequired:    true,
					ForensicRequirement: "linked_when_available",
					DecisionSummary:     "Seeded clean subject hardening record.",
				},
				Execution: &hardeningExecutionRecord{
					ExecutionID:        recommendationID("hardening-execution", subjectRef, "seed"),
					SubjectRef:         subjectRef,
					TriggerRef:         recommendationID("hardening-trigger", subjectRef, runtimeFindingSBOMMismatch),
					DecisionRef:        recommendationID("hardening-policy", subjectRef, runtimeFindingSBOMMismatch),
					ActionsApplied:     []hardeningAction{{ActionID: recommendationID("hardening-action", subjectRef, hardeningActionApplyNetworkQuarantine), ActionType: hardeningActionApplyNetworkQuarantine, SubjectRef: subjectRef, Scope: "workload_only", Parameters: map[string]any{"ttl": hardeningModeSoftIsolationTTL}, IsImmediate: true, IsReversible: true}},
					ExecutedAt:         now.Add(-30 * time.Minute),
					ExecutionResult:    "soft_isolation_applied",
					RollbackPlan:       []string{"Remove temporary quarantine after clean verification."},
					ForensicRefs:       nil,
					IncidentRefs:       nil,
					RecommendationRefs: []string{hardeningRecommendationID(subjectRef, "seed")},
					ExpiresAt:          timePointer(now.Add(15 * time.Minute)),
				},
				Posture: &defensePostureState{
					SubjectRef:         subjectRef,
					CurrentMode:        hardeningModeSoftIsolation,
					ActiveRestrictions: []string{hardeningActionApplyNetworkQuarantine},
					TriggerSummary:     "Seeded clean subject hardening posture.",
					ForensicStatus:     "not_requested",
					RollbackReady:      true,
					ExpiresAt:          timePointer(now.Add(15 * time.Minute)),
				},
			}),
		},
	}
	for _, event := range events {
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}
	return subjectRef
}

func findDefensePosture(t *testing.T, items []defensePostureState, workload string) defensePostureState {
	t.Helper()
	for _, item := range items {
		if strings.Contains(item.SubjectRef, workload) {
			return item
		}
	}
	t.Fatalf("defense posture for %s not found", workload)
	return defensePostureState{}
}
