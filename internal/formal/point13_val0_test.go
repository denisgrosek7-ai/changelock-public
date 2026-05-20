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

func point13Val0StringSliceContains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
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

	t.Run("padded inherited point12 pass token fails raw exact dependency gate", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		padded := " " + point12ValEPoint12PassToken + " "
		model.Dependency.Point12PassToken = padded
		model.Dependency.Point12.Point12PassToken = padded
		got, reasons := point13Val0DependencyStateAndReasons(model.Dependency)
		if got != Point13Val0StateBlocked || !point13Val0StringSliceContains(reasons, "point12_pass_evidence_missing") {
			t.Fatalf("expected padded point12 pass token to block with exact reason, got state=%s reasons=%v", got, reasons)
		}
		model = ComputePoint13Val0Foundation(model)
		if model.DependencyState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected padded inherited point12 pass token to block foundation, got %#v", model)
		}
	})

	t.Run("tab newline inherited point12 tenant scope fails raw exact dependency gate", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		retagged := model.Dependency.Point12TenantScope + "\n"
		model.Dependency.Point12TenantScope = retagged
		model.Dependency.Point12.PassClosureManifest.TenantScope = retagged
		got, reasons := point13Val0DependencyStateAndReasons(model.Dependency)
		if got != Point13Val0StateBlocked || !point13Val0StringSliceContains(reasons, "dependency_snapshot_identity_invalid") {
			t.Fatalf("expected retagged tenant scope to block with exact reason, got state=%s reasons=%v", got, reasons)
		}
		model = ComputePoint13Val0Foundation(model)
		if model.DependencyState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected retagged inherited tenant scope to block foundation, got %#v", model)
		}
	})

	t.Run("stale embedded point12 summary cannot hide mutated profile context", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.Dependency.Point12.Dependency.Val0.Manifest.ProfileContext.CurrentProfileHash = ""
		got, reasons := point13Val0DependencyStateAndReasons(model.Dependency)
		if got != Point13Val0StateBlocked || !point13Val0StringSliceContains(reasons, "point12_recomputed_snapshot_mismatch") {
			t.Fatalf("expected stale embedded point12 recompute mismatch to block with exact reason, got state=%s reasons=%v", got, reasons)
		}
		model = ComputePoint13Val0Foundation(model)
		if model.DependencyState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected stale embedded point12 profile context to block foundation, got %#v", model)
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

	t.Run("padded pilot tenant scope blocks raw exact sibling path", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.PilotReadiness.PilotTenantScope = model.Dependency.Point12TenantScope + " "
		model = ComputePoint13Val0Foundation(model)
		if model.PilotReadinessState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected padded pilot tenant scope to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "pilot_readiness:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact pilot readiness blocking reason, got %#v", model.BlockingReasons)
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

	t.Run("leading whitespace intake boundary blocks raw exact sibling path", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.CustomerOnboardingBoundary.FirstRepoIntakeBoundary = " " + model.PilotReadiness.EvidenceHandlingBoundary
		model = ComputePoint13Val0Foundation(model)
		if model.CustomerOnboardingState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected whitespace retagged intake boundary to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "customer_onboarding:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact customer onboarding blocking reason, got %#v", model.BlockingReasons)
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

	t.Run("tab newline support tenant scope blocks raw exact sibling path", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationBoundary.TenantScope = model.Dependency.Point12TenantScope + "\t"
		model = ComputePoint13Val0Foundation(model)
		if model.SupportEscalationState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected tab retagged support tenant scope to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "support_escalation:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact support escalation blocking reason, got %#v", model.BlockingReasons)
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

	t.Run("newline offboarding tenant scope blocks raw exact sibling path", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.OffboardingRetentionBoundary.TenantScope = model.Dependency.Point12TenantScope + "\n"
		model = ComputePoint13Val0Foundation(model)
		if model.OffboardingRetentionState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected newline retagged offboarding tenant scope to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "offboarding_retention:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact offboarding retention blocking reason, got %#v", model.BlockingReasons)
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

	t.Run("forbidden wording cannot be laundered through allowed list", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.AllowedCustomerFacingWording = []string{"production approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected forbidden allowed wording list mutation to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("unhyphenated regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"regulator approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected unhyphenated regulator approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("unicode dash regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"regulator\u2011approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected unicode dash regulator approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("slash regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"regulator/approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected slash-separated regulator approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("dot production approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"production.approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected dot-separated production approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("underscore production approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"production_approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected underscore production approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("intra bucket filler production approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"production is approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected filler production approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("confusable regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"regulat\u043er approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected confusable regulator approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("zero width regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"regulator appro\u200dved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected zero-width regulator approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("zero width separator production approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"production\u200bapproved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected zero-width separator production approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("math bold production approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"\U0001d429\U0001d42b\U0001d428\U0001d41d\U0001d42e\U0001d41c\U0001d42d\U0001d422\U0001d428\U0001d427 \U0001d41a\U0001d429\U0001d429\U0001d42b\U0001d428\U0001d42f\U0001d41e\U0001d41d"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected math bold production approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("greek nu regulator approval wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"regulator appro\u03bded"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected greek nu regulator approval wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("greek upsilon production wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"prod\u03c5ction approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected greek upsilon production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("small cap u production wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"prod\U00001d1cction approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected small-cap u production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin upsilon production wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"prod\u028action approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin upsilon production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("greek delta approved wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"production approve\u03b4"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected greek delta production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("small cap t global truth wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"global \U00001d1bruth"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected small-cap t global truth wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin alpha global truth wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"glob\u0251l truth"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin alpha global truth wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin iota official authority wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"off\u0269cial authority"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin iota official authority wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("dental click global truth wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"g\u01c0obal truth"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected dental-click global truth wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("armenian oh official authority wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"\u0585fficial authority"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected armenian-oh official authority wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("greek eta production wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"productio\u03b7 approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected greek eta production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin eng production wording blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"productio\u014b approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin eng production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("underscore machine token remains non-boundary safe wording", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.InternalDiagnosticTexts = []string{"internal_production_approved_metric"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateActive ||
			point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected underscore machine token not to become a forbidden phrase, got %#v", model)
		}
	})

	t.Run("zero width split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"deployment"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"appro\u2060ved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected zero-width split deployment approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("word fragment split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"produc"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"tion approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected word-fragment split production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("right leg u split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"prod\uab4e"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"ction approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected right-leg u split production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin upsilon split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"prod\u028a"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"ction approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin upsilon split production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("greek nu split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"production"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"appro\u03bded"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected greek nu split production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("greek delta split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"production"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"approve\u03b4"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected greek delta split production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("small cap t split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"global"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"\U00001d1bruth"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected small-cap t split global truth wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin alpha split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"glob\u0251l"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"truth"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin alpha split global truth wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin iota split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"off\u0269cial"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"authority"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin iota split official authority wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("dental click split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"g\u01c0obal"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"truth"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected dental-click split global truth wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("armenian oh split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"\u0585fficial"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"authority"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected armenian-oh split official authority wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("armenian vo split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"productio\u0578"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected armenian vo split production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("latin n with long right leg split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"productio\u019e"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected latin n with long right leg split production approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("split forbidden wording across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"deployment"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected split deployment approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("split regulator approval across customer export surfaces blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"regulator"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected split regulator approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("split forbidden wording involving support surface blocks", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"deployment"}
		model.NoOverclaimCustomerWording.ObservedSupportFacingTexts = []string{"approved"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateBlocked || model.CurrentState != Point13Val0StateBlocked {
			t.Fatalf("expected support-involved split deployment approved wording to block, got %#v", model)
		}
		if !point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected exact no-overclaim blocking reason, got %#v", model.BlockingReasons)
		}
	})

	t.Run("all allowed split safe wording does not false positive", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.NoOverclaimCustomerWording.ObservedCustomerFacingTexts = []string{"pilot readiness support"}
		model.NoOverclaimCustomerWording.ObservedExportFacingTexts = []string{"evidence candidate"}
		model.NoOverclaimCustomerWording.ObservedSupportFacingTexts = []string{"operational onboarding boundary"}
		model = ComputePoint13Val0Foundation(model)
		if model.NoOverclaimState != Point13Val0StateActive {
			t.Fatalf("expected allowed split-safe wording to remain active, got %#v", model)
		}
		if point13Val0StringSliceContains(model.BlockingReasons, "no_overclaim:"+Point13Val0StateBlocked) {
			t.Fatalf("expected no no-overclaim blocking reason for allowed split-safe wording, got %#v", model.BlockingReasons)
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

func TestPoint13Val0AggregateStateRawExactComponents(t *testing.T) {
	t.Run("padded component state blocks aggregate", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.PilotReadinessState = " " + Point13Val0StateActive + " "
		if got := EvaluatePoint13Val0State(model); got != Point13Val0StateBlocked {
			t.Fatalf("expected padded component state to block aggregate, got %s", got)
		}
		reasons := point13Val0BlockingReasons(model)
		if !point13Val0StringSliceContains(reasons, "pilot_readiness:invalid_state") {
			t.Fatalf("expected invalid pilot readiness state reason, got %#v", reasons)
		}
	})

	t.Run("non dependency review required propagates review required", func(t *testing.T) {
		model := activePoint13Val0Foundation()
		model.SupportEscalationState = Point13Val0StateReviewRequired
		if got := EvaluatePoint13Val0State(model); got != Point13Val0StateReviewRequired {
			t.Fatalf("expected support review state to propagate, got %s", got)
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
