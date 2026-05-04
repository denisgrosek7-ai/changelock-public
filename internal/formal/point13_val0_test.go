package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

var (
	point13Val0ActiveFoundationBaselineJSON []byte
	point13Val0ActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint13Val0Foundation(model Point13Val0Foundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint13Val0Foundation(payload []byte) Point13Val0Foundation {
	var clone Point13Val0Foundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func uncachedActivePoint13Val0Foundation() Point13Val0Foundation {
	return ComputePoint13Val0Foundation(Point13Val0FoundationModel())
}

func activePoint13Val0Foundation() Point13Val0Foundation {
	point13Val0ActiveFoundationBaselineOnce.Do(func() {
		point13Val0ActiveFoundationBaselineJSON = mustMarshalPoint13Val0Foundation(uncachedActivePoint13Val0Foundation())
	})
	return clonePoint13Val0Foundation(point13Val0ActiveFoundationBaselineJSON)
}

func withProductionImpactingAIBoundary(model *Point13Val0Foundation) {
	model.AIEvidenceCandidatePilotBoundary.ProductionImpactingActionRequested = true
	model.AIEvidenceCandidatePilotBoundary.HumanApprovalRef = "approval_event_point13_val0_001"
	model.AIEvidenceCandidatePilotBoundary.ReasonRef = "reason_point13_val0_prod_change_001"
	model.AIEvidenceCandidatePilotBoundary.ExpiryWindowRef = "expiry_window_point13_val0_24h"
	model.AIEvidenceCandidatePilotBoundary.SandboxResultRef = "sandbox_result_point13_val0_001"
	model.AIEvidenceCandidatePilotBoundary.RollbackPlanRef = "rollback_plan_point13_val0_001"
	model.AIEvidenceCandidatePilotBoundary.AuditEventRef = "audit_point13_val0_ai_boundary_001"
	model.AIEvidenceCandidatePilotBoundary.PostActionVerificationPlanRef = "post_action_verification_plan_point13_val0_001"
}

func TestPoint13Val0FoundationFixtureIsolation(t *testing.T) {
	t.Run("raw production path still computes", func(t *testing.T) {
		model := uncachedActivePoint13Val0Foundation()
		if model.CurrentState != Point13Val0StateActive {
			t.Fatalf("expected raw production path to compute active foundation, got %#v", model)
		}
	})

	t.Run("cached fixture mutation does not contaminate next clone", func(t *testing.T) {
		mutated := activePoint13Val0Foundation()
		mutated.Dependency.AIGovernanceBackfillVerified = false
		mutated.PilotReadiness.PilotOwnerRef = ""
		mutated.CustomerOnboardingBoundary.CustomerArtifactPromotedToCanonical = true

		fresh := activePoint13Val0Foundation()
		if !fresh.Dependency.AIGovernanceBackfillVerified {
			t.Fatalf("expected dependency verification to remain true on fresh clone, got %#v", fresh.Dependency)
		}
		if fresh.PilotReadiness.PilotOwnerRef == "" {
			t.Fatalf("expected pilot owner ref on fresh clone, got %#v", fresh.PilotReadiness)
		}
		if fresh.CustomerOnboardingBoundary.CustomerArtifactPromotedToCanonical {
			t.Fatalf("expected customer artifact promotion flag false on fresh clone, got %#v", fresh.CustomerOnboardingBoundary)
		}
	})
}

func TestPoint13Val0DependencyState(t *testing.T) {
	t.Run("valid point12 pass confirmed dependency active", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		if model.DependencyState != Point13Val0StateActive || model.CurrentState != Point13Val0StateActive {
			t.Fatalf("expected active dependency and foundation, got %#v", model)
		}
	})

	t.Run("missing point12 final closure blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.Dependency.Point12CurrentState = ""
		model = ComputePoint13Val0Foundation(model)
		if model.DependencyState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected missing point12 final closure to block, got %#v", model)
		}
	})

	t.Run("non pass confirmed point12 closure blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.Dependency.Point12CurrentState = Point12ValEStateActive
		model.Dependency.Point12.CurrentState = Point12ValEStateActive
		model = ComputePoint13Val0Foundation(model)
		if model.DependencyState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected non pass confirmed point12 closure to block, got %#v", model)
		}
	})

	t.Run("missing ai governance backfill prerequisite returns review required", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.Dependency.AIGovernanceBackfillVerified = false
		model = ComputePoint13Val0Foundation(model)
		if model.DependencyState != Point13Val0StateReviewRequired || model.CurrentState != Point13Val0StateReviewRequired {
			t.Fatalf("expected missing ai governance verification to require review, got %#v", model)
		}
	})

	t.Run("local readiness cannot override dependency failure", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.Dependency.Point12ReviewerResult = point12ValEReviewerResultPass
		model.Dependency.Point12.PassClosureManifest.ReviewerResult = point12ValEReviewerResultPass
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState == Point13Val0StateActive {
			t.Fatalf("expected dependency failure to override local readiness, got %#v", model)
		}
	})

	t.Run("point 13 pass cannot appear in val0", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		payload := string(mustMarshalPoint13Val0Foundation(model))
		if strings.Contains(payload, point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected no point_13_pass token in point13 val0 payload, got %s", payload)
		}
		if model.Dependency.Point12PassToken != point12ValEPoint12PassToken {
			t.Fatalf("expected point12 pass token to exist only as dependency evidence, got %#v", model.Dependency)
		}
	})
}

func TestPoint13Val0PilotReadinessState(t *testing.T) {
	t.Run("valid pilot readiness active without point13 pass", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		if model.PilotReadinessState != Point13Val0StateActive || strings.Contains(string(mustMarshalPoint13Val0Foundation(model)), point13Val0BlockedPoint13PassToken) {
			t.Fatalf("expected active pilot readiness without point13 pass, got %#v", model)
		}
	})

	t.Run("pilot success does not mean production approval", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		if !model.PilotReadiness.PilotSuccessDoesNotMeanProductionApproval || model.PilotReadiness.ProductionApprovalImplied {
			t.Fatalf("expected pilot success to stay non-authoritative, got %#v", model.PilotReadiness)
		}
	})

	t.Run("missing pilot owner blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.PilotReadiness.PilotOwnerRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.PilotReadinessState != Point13Val0StateBlocked {
			t.Fatalf("expected missing pilot owner to block, got %#v", model)
		}
	})

	t.Run("missing tenant scope blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.PilotReadiness.PilotTenantScope = ""
		model = ComputePoint13Val0Foundation(model)
		if model.PilotReadinessState != Point13Val0StateBlocked {
			t.Fatalf("expected missing tenant scope to block, got %#v", model)
		}
	})

	t.Run("first repo intake outside declared boundary blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.CustomerOnboardingBoundary.FirstRepoIntakeBoundary = "boundary_point13_wrong_repo_scope_001"
		model = ComputePoint13Val0Foundation(model)
		if model.CustomerOnboardingState != Point13Val0StateBlocked {
			t.Fatalf("expected intake boundary mismatch to block, got %#v", model)
		}
	})
}

func TestPoint13Val0CustomerOnboardingBoundaryState(t *testing.T) {
	t.Run("customer upload cannot become canonical without governance event", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.CustomerOnboardingBoundary.CustomerArtifactPromotedToCanonical = true
		model = ComputePoint13Val0Foundation(model)
		if model.CustomerOnboardingState != Point13Val0StateBlocked {
			t.Fatalf("expected canonical promotion without governance event to block, got %#v", model)
		}
	})

	t.Run("customer artifact remains candidate support material until canonical validation", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		if model.CustomerOnboardingBoundary.CustomerArtifactClassification != point13Val0CustomerArtifactCandidateOnly ||
			!model.CustomerOnboardingBoundary.CustomerUploadIsCandidateOnly ||
			model.CustomerOnboardingState != Point13Val0StateActive {
			t.Fatalf("expected customer artifact to remain candidate only, got %#v", model)
		}
	})

	t.Run("customer onboarding cannot bypass evidence or context identity", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.CustomerOnboardingBoundary.EvidenceIdentityRef = "artifact_point13_wrong_001"
		model = ComputePoint13Val0Foundation(model)
		if model.CustomerOnboardingState != Point13Val0StateBlocked {
			t.Fatalf("expected evidence identity bypass to block, got %#v", model)
		}
	})
}

func TestPoint13Val0SupportEscalationBoundaryState(t *testing.T) {
	t.Run("support escalation cannot bypass evidence spine", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationBoundary.EvidenceSpineBypassAttempted = true
		model = ComputePoint13Val0Foundation(model)
		if model.SupportEscalationState != Point13Val0StateBlocked {
			t.Fatalf("expected evidence spine bypass attempt to block, got %#v", model)
		}
	})

	t.Run("support escalation cannot override core decision", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationBoundary.CoreDecisionOverrideAttempted = true
		model = ComputePoint13Val0Foundation(model)
		if model.SupportEscalationState != Point13Val0StateBlocked {
			t.Fatalf("expected core decision override attempt to block, got %#v", model)
		}
	})

	t.Run("support escalation cannot approve production mutation", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationBoundary.ProductionMutationApprovalAttempted = true
		model = ComputePoint13Val0Foundation(model)
		if model.SupportEscalationState != Point13Val0StateBlocked {
			t.Fatalf("expected production mutation approval attempt to block, got %#v", model)
		}
	})

	t.Run("missing support owner or audit event blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationBoundary.EscalationOwnerRef = ""
		model.SupportEscalationBoundary.AuditEventRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.SupportEscalationState != Point13Val0StateBlocked {
			t.Fatalf("expected missing support owner and audit event to block, got %#v", model)
		}
	})
}

func TestPoint13Val0OffboardingRetentionBoundaryState(t *testing.T) {
	t.Run("missing retention owner blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.OffboardingRetentionBoundary.RetentionOwnerRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.OffboardingRetentionState != Point13Val0StateBlocked {
			t.Fatalf("expected missing retention owner to block, got %#v", model)
		}
	})

	t.Run("missing disposal path blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.OffboardingRetentionBoundary.DisposalPathRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.OffboardingRetentionState != Point13Val0StateBlocked {
			t.Fatalf("expected missing disposal path to block, got %#v", model)
		}
	})

	t.Run("missing tenant scope blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.OffboardingRetentionBoundary.TenantScope = ""
		model = ComputePoint13Val0Foundation(model)
		if model.OffboardingRetentionState != Point13Val0StateBlocked {
			t.Fatalf("expected missing offboarding tenant scope to block, got %#v", model)
		}
	})

	t.Run("support and pilot artifacts cannot become canonical without governance event", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.OffboardingRetentionBoundary.PilotArtifactPromotedToCanonical = true
		model.OffboardingRetentionBoundary.SupportArtifactPromotedToCanonical = true
		model = ComputePoint13Val0Foundation(model)
		if model.OffboardingRetentionState != Point13Val0StateBlocked {
			t.Fatalf("expected canonical promotion without governance event to block, got %#v", model)
		}
	})
}

func TestPoint13Val0NoOverclaimCustomerWordingState(t *testing.T) {
	for _, phrase := range point13Val0ForbiddenClaims() {
		t.Run("forbidden wording blocks "+phrase, func(t *testing.T) {
			model := activePoint13Val0Foundation()
			model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{phrase}
			model = ComputePoint13Val0Foundation(model)
			if model.NoOverclaimState != Point13Val0StateBlocked {
				t.Fatalf("expected forbidden phrase %q to block, got %#v", phrase, model)
			}
		})
	}

	t.Run("safe wording remains allowed", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{
			"pilot readiness support",
			"evidence candidate",
			"advisory recommendation",
		}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateActive {
			t.Fatalf("expected safe wording to remain allowed, got %#v", model)
		}
	})

	t.Run("forbidden wording allowed only in classified internal diagnostics", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.InternalDiagnosticTexts = []string{"blocked phrase: production approved"}
		model.NoOverclaimCustomerWording.InternalDiagnosticsClassifiedBlocked = true
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateActive {
			t.Fatalf("expected classified internal diagnostics to remain active, got %#v", model)
		}
	})
}

func TestPoint13Val0AIPilotBoundaryState(t *testing.T) {
	t.Run("all allowed ai outputs remain advisory candidate only", func(t *testing.T) {
		for _, agentType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13Val0Foundation()
			model.AIEvidenceCandidatePilotBoundary.AIOutputType = agentType
			model = ComputePoint13Val0Foundation(model)
			if model.AIPilotBoundaryState != Point13Val0StateActive {
				t.Fatalf("expected allowed AI output type %q to remain active advisory candidate, got %#v", agentType, model)
			}
		}
	})

	t.Run("all allowed ai outputs block on deployment authorized", func(t *testing.T) {
		for _, agentType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13Val0Foundation()
			model.AIEvidenceCandidatePilotBoundary.AIOutputType = agentType
			model.AIEvidenceCandidatePilotBoundary.DeploymentAuthorized = true
			model = ComputePoint13Val0Foundation(model)
			if model.AIPilotBoundaryState != Point13Val0StateBlocked {
				t.Fatalf("expected allowed AI output type %q with deployment authority to block, got %#v", agentType, model)
			}
		}
	})

	t.Run("all allowed ai outputs block on production readiness claimed", func(t *testing.T) {
		for _, agentType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			model := activePoint13Val0Foundation()
			model.AIEvidenceCandidatePilotBoundary.AIOutputType = agentType
			model.AIEvidenceCandidatePilotBoundary.ProductionReadinessClaimed = true
			model = ComputePoint13Val0Foundation(model)
			if model.AIPilotBoundaryState != Point13Val0StateBlocked {
				t.Fatalf("expected allowed AI output type %q with production readiness claim to block, got %#v", agentType, model)
			}
		}
	})

	t.Run("blocked ai taxonomy values are rejected", func(t *testing.T) {
		for _, agentType := range point12Val0BlockedAIEvidenceCandidateTypes() {
			model := activePoint13Val0Foundation()
			model.AIEvidenceCandidatePilotBoundary.AIOutputType = agentType
			model = ComputePoint13Val0Foundation(model)
			if model.AIPilotBoundaryState != Point13Val0StateBlocked {
				t.Fatalf("expected blocked AI output type %q to block, got %#v", agentType, model)
			}
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point13Val0AIEvidenceCandidatePilotBoundary)
	}{
		{name: "ai finding cannot create pass", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_FINDING"
			model.PassAllowed = true
		}},
		{name: "ai recommendation cannot create approval", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_RECOMMENDATION"
			model.ApprovalGranted = true
		}},
		{name: "ai patch proposal cannot become deployment", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_PATCH_PROPOSAL"
			model.DeploymentAuthorized = true
		}},
		{name: "ai sandbox result cannot become production readiness", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_SANDBOX_RESULT"
			model.ProductionReadinessClaimed = true
		}},
		{name: "ai approval request is not approval", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_APPROVAL_REQUEST"
			model.ApprovalRequestCreatesApproval = true
		}},
		{name: "ai finding cannot authorize deployment", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_FINDING"
			model.DeploymentAuthorized = true
		}},
		{name: "ai recommendation cannot authorize deployment", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_RECOMMENDATION"
			model.DeploymentAuthorized = true
		}},
		{name: "ai finding cannot claim production readiness", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_FINDING"
			model.ProductionReadinessClaimed = true
		}},
		{name: "ai recommendation cannot claim production readiness", mutate: func(model *Point13Val0AIEvidenceCandidatePilotBoundary) {
			model.AIOutputType = "AI_RECOMMENDATION"
			model.ProductionReadinessClaimed = true
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint13Val0Foundation()
			tc.mutate(&model.AIEvidenceCandidatePilotBoundary)
			model = ComputePoint13Val0Foundation(model)
			if model.AIPilotBoundaryState != Point13Val0StateBlocked {
				t.Fatalf("expected AI pilot boundary mutation to block, got %#v", model)
			}
		})
	}

	t.Run("production impacting ai assisted action without human approval blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		withProductionImpactingAIBoundary(&model)
		model.AIEvidenceCandidatePilotBoundary.HumanApprovalRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.AIPilotBoundaryState != Point13Val0StateBlocked {
			t.Fatalf("expected missing human approval to block, got %#v", model)
		}
	})

	t.Run("production impacting ai assisted action without rollback plan blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		withProductionImpactingAIBoundary(&model)
		model.AIEvidenceCandidatePilotBoundary.RollbackPlanRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.AIPilotBoundaryState != Point13Val0StateBlocked {
			t.Fatalf("expected missing rollback plan to block, got %#v", model)
		}
	})

	t.Run("production impacting ai assisted action without audit event blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		withProductionImpactingAIBoundary(&model)
		model.AIEvidenceCandidatePilotBoundary.AuditEventRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.AIPilotBoundaryState != Point13Val0StateBlocked {
			t.Fatalf("expected missing audit event to block, got %#v", model)
		}
	})

	t.Run("production impacting ai assisted action without expiry blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		withProductionImpactingAIBoundary(&model)
		model.AIEvidenceCandidatePilotBoundary.ExpiryWindowRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.AIPilotBoundaryState != Point13Val0StateBlocked {
			t.Fatalf("expected missing expiry window to block, got %#v", model)
		}
	})
}

func TestPoint13Val0MutationClosure(t *testing.T) {
	t.Run("mutate tenant scope blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.PilotReadiness.PilotTenantScope = "tenant_scope_wrong"
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected tenant scope drift to block, got %#v", model)
		}
	})

	t.Run("mutate evidence refs blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationBoundary.EvidenceRefs = []string{"evidence:point13-support-wrong-001"}
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected evidence ref drift to block, got %#v", model)
		}
	})

	t.Run("mutate governance event refs blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.CustomerOnboardingBoundary.CustomerArtifactPromotedToCanonical = true
		model.CustomerOnboardingBoundary.CanonicalGovernanceEventRef = "governance placeholder"
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected governance event drift to block, got %#v", model)
		}
	})

	t.Run("mutate support audit event ref blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationBoundary.AuditEventRef = ""
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected support audit event drift to block, got %#v", model)
		}
	})

	t.Run("mutate retention disposal path blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.OffboardingRetentionBoundary.DisposalPathRef = "disposal path invalid"
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected retention disposal path drift to block, got %#v", model)
		}
	})

	t.Run("mutate ai boundary from advisory to production approved blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.AIEvidenceCandidatePilotBoundary.ApprovalGranted = true
		model.AIEvidenceCandidatePilotBoundary.ProductionMutationAllowed = true
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected production approved ai boundary mutation to block, got %#v", model)
		}
	})

	t.Run("recomputing after dependency mutation does not hide drift", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.Dependency.Point12CurrentState = Point12ValEStateActive
		model.Dependency.Point12.CurrentState = Point12ValEStateActive
		model = ComputePoint13Val0Foundation(model)
		if model.CurrentState == Point13Val0StateActive {
			t.Fatalf("expected recomputation after dependency drift to remain non-active, got %#v", model)
		}
	})
}
