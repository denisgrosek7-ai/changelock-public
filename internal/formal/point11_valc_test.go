package formal

import (
	"encoding/json"
	"strings"
	"testing"
)

func point11ValCActiveDependencySnapshot() Point11ValCDependencySnapshot {
	valB := activePoint11ValBFoundation()
	return SnapshotPoint11ValCDependencyFromComputedValB(valB, Point11ValCValBReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
}

func reviewRequiredPoint11ValCDependencySnapshot() Point11ValCDependencySnapshot {
	valB := activePoint11ValBFoundation()
	valB.CurrentState = Point11ValBStateReviewRequired
	valB.DependencyState = Point11ValBDependencyStateReviewRequired
	valB.ReviewPrerequisites = []string{"valb_repo_visibility_review_prerequisite"}
	return SnapshotPoint11ValCDependencyFromComputedValB(valB, Point11ValCValBReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
}

func activePoint11ValCFoundation() Point11ValCFoundation {
	model := Point11ValCFoundationModel()
	model.Dependency = point11ValCActiveDependencySnapshot()
	return ComputePoint11ValCFoundation(model)
}

func point11ValCPassToken() string {
	return "point_11_" + "pass"
}

func TestPoint11ValCDependencyState(t *testing.T) {
	t.Run("happy path valb dependency active", func(t *testing.T) {
		snapshot := point11ValCActiveDependencySnapshot()
		if got := EvaluatePoint11ValCDependencyState(snapshot); got != Point11ValCDependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", snapshot)
		}
	})

	t.Run("copied valb projection disclaimer propagates exactly", func(t *testing.T) {
		valB := activePoint11ValBFoundation()
		valB.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_valb_disclaimer"
		valB.ClaimTypeDefinition.ProjectionDisclaimer = "projection_only not_canonical_truth component_claim_type_disclaimer"
		snapshot := SnapshotPoint11ValCDependencyFromComputedValB(valB, Point11ValCValBReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
		})
		if snapshot.ProjectionDisclaimer != valB.ProjectionDisclaimer {
			t.Fatalf("expected aggregate valb projection disclaimer, got snapshot=%q valb=%q", snapshot.ProjectionDisclaimer, valB.ProjectionDisclaimer)
		}
	})

	t.Run("malformed valb aggregate projection disclaimer blocks", func(t *testing.T) {
		valB := activePoint11ValBFoundation()
		valB.ProjectionDisclaimer = "canonical_truth"
		valB.ClaimTypeDefinition.ProjectionDisclaimer = "projection_only not_canonical_truth component_claim_type_disclaimer"
		snapshot := SnapshotPoint11ValCDependencyFromComputedValB(valB, Point11ValCValBReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
		})
		if got := EvaluatePoint11ValCDependencyState(snapshot); got != Point11ValCDependencyStateBlocked {
			t.Fatalf("expected malformed aggregate disclaimer to block dependency, got %#v", snapshot)
		}
	})

	testCases := []struct {
		name      string
		mutate    func(*Point11ValCDependencySnapshot)
		wantState string
	}{
		{name: "blocked valb claim type blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBClaimTypeState = Point11ValBClaimTypeStateBlocked
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "blocked valb issuance request blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBIssuanceRequestState = Point11ValBIssuanceRequestStateBlocked
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "blocked valb issued claim blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBIssuedClaimState = Point11ValBIssuedClaimStateBlocked
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "blocked valb registry blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBRegistryState = Point11ValBRegistryStateBlocked
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "blocked valb verification blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBVerificationState = Point11ValBVerificationStateBlocked
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "blocked valb cross domain intake blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCrossDomainIntakeState = Point11ValBCrossDomainIntakeStateBlocked
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "whitespace retagged valb current state blocks raw exact", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCurrentState = " " + Point11ValBStateActive
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "tab newline retagged valb dependency state blocks raw exact", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBDependencyState = "\t" + Point11ValBDependencyStateActive + "\n"
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "valb dependency review required propagates", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCurrentState = Point11ValBStateReviewRequired
			model.ValBDependencyState = Point11ValBDependencyStateReviewRequired
			model.ReviewPrerequisites = []string{"valb_repo_visibility_review_prerequisite"}
		}, wantState: Point11ValCDependencyStateReviewRequired},
		{name: "valb point11 pass marker blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBPoint11PassEmitted = true
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "valb authority marker blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCreatesAuthorityClaims = true
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "valb publication marker blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCreatesPublicationSideEffects = true
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "valb signing marker blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCreatesSigningSideEffects = true
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "valb anchoring marker blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCreatesAnchoringSideEffects = true
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "valb external api marker blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCreatesExternalAPISideEffects = true
		}, wantState: Point11ValCDependencyStateBlocked},
		{name: "valb production marker blocks", mutate: func(model *Point11ValCDependencySnapshot) {
			model.ValBCreatesProductionSideEffects = true
		}, wantState: Point11ValCDependencyStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := point11ValCActiveDependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint11ValCDependencyState(model); got != testCase.wantState {
				t.Fatalf("expected dependency state %q, got %#v", testCase.wantState, model)
			}
		})
	}

	t.Run("retagged valb inherited state records exact dependency reason", func(t *testing.T) {
		for _, mutate := range []func(*Point11ValCDependencySnapshot){
			func(model *Point11ValCDependencySnapshot) { model.ValBCurrentState = " " + Point11ValBStateActive },
			func(model *Point11ValCDependencySnapshot) {
				model.ValBDependencyState = "\t" + Point11ValBDependencyStateActive + "\n"
			},
		} {
			model := point11ValCActiveDependencySnapshot()
			mutate(&model)
			state, reasons := point11ValCDependencyStateAndReasons(model)
			if state != Point11ValCDependencyStateBlocked {
				t.Fatalf("expected retagged valb inherited state to block, got %q for %#v", state, model)
			}
			if !point11Val0ContainsTrimmed(reasons, "valb_dependency_not_active") {
				t.Fatalf("expected exact valb dependency reason, got %#v", reasons)
			}
		}
	})
}

func TestPoint11ValCEnforcementInputState(t *testing.T) {
	t.Run("valid enforcement input active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.EnforcementInputState != Point11ValCEnforcementInputStateActive {
			t.Fatalf("expected active enforcement input, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCGovernanceEnforcementInput)
	}{
		{name: "missing enforcement id blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.EnforcementID = ""
		}},
		{name: "enforcement unknown blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.EnforcementID = "enforcement_unknown"
		}},
		{name: "enforcement revoked blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.EnforcementID = "enforcement_revoked"
		}},
		{name: "enforcement invalid blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.EnforcementID = "enforcement_invalid"
		}},
		{name: "global tenant scope blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.TenantScope = "global"
		}},
		{name: "unscoped tenant scope blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.TenantScope = "unscoped"
		}},
		{name: "all tenants tenant scope blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.TenantScope = "all-tenants"
		}},
		{name: "wildcard tenant scope blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.TenantScope = "wildcard"
		}},
		{name: "cross tenant tenant scope blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.TenantScope = "cross-tenant"
		}},
		{name: "missing policy basis ref blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.PolicyBasisRef = ""
		}},
		{name: "missing claim refs when claims are required blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.ClaimsRequired = true
			model.ClaimRefs = nil
		}},
		{name: "inactive claim verification blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.ClaimVerificationState = Point11ValBVerificationStateBlocked
		}},
		{name: "missing registry ref blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.RegistryRef = ""
		}},
		{name: "unsupported requested action blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.RequestedAction = "unsupported_action"
		}},
		{name: "real side effect requested action blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.RequestedAction = point11ValCRequestedActionDeploy
		}},
		{name: "forbidden requested surface blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.RequestedSurface = "forbidden_surface"
		}},
		{name: "overclaim requested outcome blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.RequestedOutcome = "production approved"
		}},
		{name: "missing evidence refs blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.EvidenceRefs = nil
		}},
		{name: "missing governance event for customer public export surface blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.RequestedSurface = point11Val0PublicationSurfaceExport
			model.GovernanceEventRef = ""
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *Point11ValCGovernanceEnforcementInput) {
			model.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model.EnforcementInput)
			model = ComputePoint11ValCFoundation(model)
			if model.EnforcementInputState != Point11ValCEnforcementInputStateBlocked {
				t.Fatalf("expected blocked enforcement input state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValCEnforcementResultState(t *testing.T) {
	t.Run("valid enforcement result active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.EnforcementResultState != Point11ValCEnforcementResultStateActive {
			t.Fatalf("expected active enforcement result, got %#v", model)
		}
	})

	t.Run("review required state requires review required reason", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.EnforcementResult.ExceptionDecisionState = Point11ValCExceptionDecisionStateReviewRequired
		model.EnforcementResult.ReviewRequiredReason = ""
		model = ComputePoint11ValCFoundation(model)
		if model.EnforcementResultState != Point11ValCEnforcementResultStateBlocked {
			t.Fatalf("expected missing review required reason to block, got %#v", model)
		}
	})

	t.Run("diagnostics include policy claim abac exception and evidence reason", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		reasons := strings.Join(model.Diagnostics.EnforcementResultReasons, " ")
		for _, required := range []string{"policy_reason_present", "claim_reason_present", "abac_reason_present", "exception_reason_present", "evidence_reason_present"} {
			if !strings.Contains(reasons, required) {
				t.Fatalf("expected enforcement diagnostics to include %q, got %#v", required, model.Diagnostics.EnforcementResultReasons)
			}
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCPolicyDecisionEnforcementResult)
	}{
		{name: "invalid policy blocks result", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.PolicyResultState = point11ValCCheckStateBlocked
			model.BlockedActionReason = "policy invalid"
		}},
		{name: "whitespace retagged active policy result blocks raw exact", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.PolicyResultState = " " + point11ValCPolicyStateActive
			model.BlockedActionReason = "policy invalid"
		}},
		{name: "invalid claim blocks result", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.ClaimResultState = point11ValCCheckStateBlocked
			model.BlockedActionReason = "claim invalid"
		}},
		{name: "abac deny blocks result", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.ABACDecisionState = Point11ValCABACDecisionStateBlocked
			model.BlockedActionReason = "abac deny"
		}},
		{name: "invalid exception blocks result", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.ExceptionDecisionState = Point11ValCExceptionDecisionStateBlocked
			model.BlockedActionReason = "exception invalid"
		}},
		{name: "invalid emergency blocks result", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.EmergencyDecisionState = Point11ValCExceptionDecisionStateBlocked
			model.BlockedActionReason = "emergency invalid"
		}},
		{name: "evidence mismatch blocks result", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.EvidenceState = point11ValCCheckStateBlocked
			model.BlockedActionReason = "evidence invalid"
		}},
		{name: "blocked state requires blocked action reason", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.PolicyResultState = point11ValCCheckStateBlocked
			model.BlockedActionReason = ""
		}},
		{name: "allow outcome cannot contain production compliance certification overclaim", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.EnforcementOutcome = "production approved"
			model.BlockedActionReason = "overclaim blocked"
		}},
		{name: "split source of truth diagnostics block result", mutate: func(model *Point11ValCPolicyDecisionEnforcementResult) {
			model.Diagnostics = []string{"source of", "truth"}
			model.BlockedActionReason = "overclaim blocked"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model.EnforcementResult)
			model = ComputePoint11ValCFoundation(model)
			if model.EnforcementResultState != Point11ValCEnforcementResultStateBlocked {
				t.Fatalf("expected blocked enforcement result state, got %#v", model)
			}
		})
	}

	t.Run("whitespace retagged active policy result records exact reason", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.EnforcementResult.PolicyResultState = " " + point11ValCPolicyStateActive
		model.EnforcementResult.BlockedActionReason = "policy invalid"
		model = ComputePoint11ValCFoundation(model)
		if model.EnforcementResultState != Point11ValCEnforcementResultStateBlocked || model.CurrentState != Point11ValCStateBlocked {
			t.Fatalf("expected retagged policy result to block exactly, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.Diagnostics.EnforcementResultReasons, "enforcement_result_policy_invalid") {
			t.Fatalf("expected exact enforcement result policy reason, got %#v", model.Diagnostics.EnforcementResultReasons)
		}
	})
}

func TestPoint11ValCABACDecisionState(t *testing.T) {
	t.Run("valid abac allow active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.ABACDecisionState != Point11ValCABACDecisionStateActive {
			t.Fatalf("expected active abac decision, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCABACEnforcementDecision)
	}{
		{name: "unknown attributes block", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.UnknownAttributes = []string{"attr:unknown"}
		}},
		{name: "denied attributes override allowed attributes", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.DeniedAttributes = []string{"deny:customer_export"}
		}},
		{name: "missing deny over allow explanation blocks", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.Explanation = "missing precedence text"
			model.Diagnostics = []string{"no precedence visible"}
		}},
		{name: "global tenant scope blocks", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.TenantScope = "global"
		}},
		{name: "missing policy profile blocks", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.PolicyProfileRef = ""
		}},
		{name: "missing audit id blocks", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.AuditID = ""
		}},
		{name: "exception cannot override deny unless active scoped unexpired revocable governance approved", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.DeniedAttributes = []string{"deny:customer_export"}
			model.ExceptionRefs = []string{"exception_point11_valc_scope_override_001"}
			model.ExceptionState = Point11ValCExceptionDecisionStateBlocked
		}},
		{name: "abac allow cannot override invalid policy claim evidence", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.PolicyState = point11ValCCheckStateBlocked
		}},
		{name: "requested action with real side effect blocks", mutate: func(model *Point11ValCABACEnforcementDecision) {
			model.RequestedAction = point11ValCRequestedActionExecuteEnforcement
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model.ABACDecision)
			model = ComputePoint11ValCFoundation(model)
			if model.ABACDecisionState != Point11ValCABACDecisionStateBlocked {
				t.Fatalf("expected blocked abac decision state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValCExceptionDecisionState(t *testing.T) {
	t.Run("valid exception emergency decision active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.ExceptionDecisionState != Point11ValCExceptionDecisionStateActive {
			t.Fatalf("expected active exception decision, got %#v", model)
		}
	})

	t.Run("temporary override candidate yields review required", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.ExceptionDecision.ExceptionType = point11ValCExceptionTypeTemporaryOverrideCandidate
		model = ComputePoint11ValCFoundation(model)
		if model.ExceptionDecisionState != Point11ValCExceptionDecisionStateReviewRequired {
			t.Fatalf("expected temporary override candidate to require review, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCExceptionEmergencyDecision)
	}{
		{name: "missing exception ref blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.ExceptionRef = ""
		}},
		{name: "exception unknown blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.ExceptionRef = "exception_unknown"
		}},
		{name: "exception revoked blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.ExceptionRef = "exception_revoked"
		}},
		{name: "exception invalid blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.ExceptionRef = "exception_invalid"
		}},
		{name: "missing emergency ref blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.EmergencyRef = ""
		}},
		{name: "emergency unknown blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.EmergencyRef = "emergency_unknown"
		}},
		{name: "emergency revoked blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.EmergencyRef = "emergency_revoked"
		}},
		{name: "emergency invalid blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.EmergencyRef = "emergency_invalid"
		}},
		{name: "global tenant scope blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.TenantScope = "global"
		}},
		{name: "issuer equals approver blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.ApproverRef = model.IssuerRef
		}},
		{name: "missing authority basis ref blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.AuthorityBasisRef = ""
		}},
		{name: "missing governance event ref blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.GovernanceEventRef = ""
		}},
		{name: "expired exception blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.ExpiresAt = "2000-01-01T00:00:00Z"
		}},
		{name: "missing revocation path blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.RevocationPathRef = ""
		}},
		{name: "missing monitoring requirement blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.MonitoringRequirementRef = ""
		}},
		{name: "missing rollback review condition blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.RollbackOrReviewConditionRef = ""
		}},
		{name: "permanent silent exception blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.PermanentSilentException = true
		}},
		{name: "emergency production approval wording blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.Reason = "production approved emergency exception"
		}},
		{name: "emergency canonical mutation flag blocks", mutate: func(model *Point11ValCExceptionEmergencyDecision) {
			model.MutatesCanonicalEvidenceSpine = true
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model.ExceptionDecision)
			model = ComputePoint11ValCFoundation(model)
			if model.ExceptionDecisionState != Point11ValCExceptionDecisionStateBlocked {
				t.Fatalf("expected blocked exception decision state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValCPrecedenceState(t *testing.T) {
	t.Run("valid precedence result active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.PrecedenceState != Point11ValCPrecedenceStateActive {
			t.Fatalf("expected active precedence state, got %#v", model)
		}
	})

	t.Run("local policy deny with valid scoped exception produces review required candidate", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.Precedence.LocalPolicyResult = point11ValCLocalPolicyResultDeny
		model.Precedence.ABACResult = point11ValCABACResultDeny
		model.Precedence.ExceptionResult = Point11ValCExceptionDecisionStateReviewRequired
		model.Precedence.GovernanceEventResolved = false
		model = ComputePoint11ValCFoundation(model)
		if model.PrecedenceState != Point11ValCPrecedenceStateReviewRequired {
			t.Fatalf("expected valid scoped exception override candidate to require review, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCOverridePrecedence)
		want   string
	}{
		{name: "invalid policy cannot be overridden by claim", mutate: func(model *Point11ValCOverridePrecedence) {
			model.LocalPolicyResult = point11ValCLocalPolicyResultInvalid
			model.ExceptionResult = Point11ValCExceptionDecisionStateActive
		}, want: Point11ValCPrecedenceStateBlocked},
		{name: "invalid claim cannot be fixed by exception", mutate: func(model *Point11ValCOverridePrecedence) {
			model.ClaimResultState = point11ValCCheckStateBlocked
			model.ExceptionResult = Point11ValCExceptionDecisionStateActive
		}, want: Point11ValCPrecedenceStateBlocked},
		{name: "expired exception cannot override", mutate: func(model *Point11ValCOverridePrecedence) {
			model.LocalPolicyResult = point11ValCLocalPolicyResultDeny
			model.ABACResult = point11ValCABACResultDeny
			model.ExceptionResult = Point11ValCExceptionDecisionStateReviewRequired
			model.ExceptionExpired = true
		}, want: Point11ValCPrecedenceStateBlocked},
		{name: "revoked exception cannot override", mutate: func(model *Point11ValCOverridePrecedence) {
			model.LocalPolicyResult = point11ValCLocalPolicyResultDeny
			model.ABACResult = point11ValCABACResultDeny
			model.ExceptionResult = Point11ValCExceptionDecisionStateReviewRequired
			model.ExceptionRevoked = true
		}, want: Point11ValCPrecedenceStateBlocked},
		{name: "emergency produces review required only when scoped time bound monitored revocable", mutate: func(model *Point11ValCOverridePrecedence) {
			model.LocalPolicyResult = point11ValCLocalPolicyResultDeny
			model.ABACResult = point11ValCABACResultDeny
			model.ExceptionScoped = false
			model.ExceptionGovernanceApproved = false
			model.EmergencyResult = Point11ValCExceptionDecisionStateReviewRequired
			model.GovernanceEventResolved = false
		}, want: Point11ValCPrecedenceStateReviewRequired},
		{name: "remote claim cannot override local policy", mutate: func(model *Point11ValCOverridePrecedence) {
			model.RemoteClaimOverridesLocal = true
		}, want: Point11ValCPrecedenceStateBlocked},
		{name: "abac deny beats allow", mutate: func(model *Point11ValCOverridePrecedence) {
			model.ABACResult = point11ValCABACResultDeny
			model.ExceptionScoped = false
			model.EmergencyScoped = false
		}, want: Point11ValCPrecedenceStateBlocked},
		{name: "blocked beats review required", mutate: func(model *Point11ValCOverridePrecedence) {
			model.LocalPolicyResult = point11ValCLocalPolicyResultInvalid
			model.EmergencyResult = Point11ValCExceptionDecisionStateReviewRequired
		}, want: Point11ValCPrecedenceStateBlocked},
		{name: "review required beats active allow unless valid governance event resolves it", mutate: func(model *Point11ValCOverridePrecedence) {
			model.LocalPolicyResult = point11ValCLocalPolicyResultReview
			model.GovernanceEventResolved = false
		}, want: Point11ValCPrecedenceStateReviewRequired},
		{name: "overclaim wording in override diagnostics blocks", mutate: func(model *Point11ValCOverridePrecedence) {
			model.Diagnostics = []string{"official authority"}
		}, want: Point11ValCPrecedenceStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model.Precedence)
			model = ComputePoint11ValCFoundation(model)
			if model.PrecedenceState != testCase.want {
				t.Fatalf("expected precedence state %q, got %#v", testCase.want, model)
			}
		})
	}
}

func TestPoint11ValCMonitoringState(t *testing.T) {
	t.Run("valid monitoring linked emergency active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.MonitoringState != Point11ValCMonitoringStateActive {
			t.Fatalf("expected active monitoring state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCMonitoringLinkedEmergencyHandling)
		want   string
	}{
		{name: "missing monitoring link blocks emergency active", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.MonitoringLinkID = ""
		}, want: Point11ValCMonitoringStateBlocked},
		{name: "stale monitoring signal requires review and never active", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.SignalFreshness = point11ValCMonitoringSignalStale
		}, want: Point11ValCMonitoringStateReviewRequired},
		{name: "expired review deadline blocks", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.ReviewDeadline = "2000-01-01T00:00:00Z"
		}, want: Point11ValCMonitoringStateBlocked},
		{name: "missing expiry enforcement blocks", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.ExpiryEnforcementState = point11ValCCheckStateBlocked
		}, want: Point11ValCMonitoringStateBlocked},
		{name: "missing revocation check blocks", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.RevocationCheckState = point11ValCCheckStateBlocked
		}, want: Point11ValCMonitoringStateBlocked},
		{name: "missing rollback check blocks", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.RollbackCheckState = point11ValCCheckStateBlocked
		}, want: Point11ValCMonitoringStateBlocked},
		{name: "missing escalation ref for high risk emergency blocks", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.HighRiskEmergency = true
			model.EscalationRef = ""
		}, want: Point11ValCMonitoringStateBlocked},
		{name: "monitoring creates no action side effects", mutate: func(model *Point11ValCMonitoringLinkedEmergencyHandling) {
			model.CreatesActionSideEffects = true
		}, want: Point11ValCMonitoringStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model.Monitoring)
			model = ComputePoint11ValCFoundation(model)
			if model.MonitoringState != testCase.want {
				t.Fatalf("expected monitoring state %q, got %#v", testCase.want, model)
			}
		})
	}
}

func TestPoint11ValCDashboardState(t *testing.T) {
	t.Run("valid bounded dashboard read model active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.DashboardState != Point11ValCDashboardStateActive {
			t.Fatalf("expected active dashboard state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCGovernanceDashboardReadModel)
	}{
		{name: "missing dashboard view id blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.DashboardViewID = ""
		}},
		{name: "dashboard unknown blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.DashboardViewID = "dashboard_unknown"
		}},
		{name: "dashboard revoked blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.DashboardViewID = "dashboard_revoked"
		}},
		{name: "dashboard invalid blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.DashboardViewID = "dashboard_invalid"
		}},
		{name: "malformed source decision ref blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.SourceDecisionRefs = []string{"decision_revoked"}
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.ProjectionDisclaimer = "canonical_truth"
		}},
		{name: "creates publication side effects blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.CreatesPublicationSideEffects = true
		}},
		{name: "creates authority claim blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.CreatesAuthorityClaim = true
		}},
		{name: "mutates canonical state blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.MutatesCanonicalState = true
		}},
		{name: "dashboard certification wording blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.RenderedState = "certified governance dashboard"
		}},
		{name: "dashboard production approval wording blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.RenderedState = "production approved dashboard"
		}},
		{name: "dashboard compliance guarantee wording blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.RenderedState = "compliance guaranteed dashboard"
		}},
		{name: "dashboard source of truth wording blocks", mutate: func(model *Point11ValCGovernanceDashboardReadModel) {
			model.RenderedState = "source of truth governance dashboard"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model.Dashboard)
			model = ComputePoint11ValCFoundation(model)
			if model.DashboardState != Point11ValCDashboardStateBlocked {
				t.Fatalf("expected blocked dashboard state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValCAggregateState(t *testing.T) {
	t.Run("aggregate active only when all components active", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.CurrentState != Point11ValCStateActive {
			t.Fatalf("expected active valc foundation, got %#v", model)
		}
	})

	t.Run("dependency review required yields aggregate review required if no local blockers", func(t *testing.T) {
		model := Point11ValCFoundationModel()
		model.Dependency = reviewRequiredPoint11ValCDependencySnapshot()
		model = ComputePoint11ValCFoundation(model)
		if model.CurrentState != Point11ValCStateReviewRequired {
			t.Fatalf("expected dependency review required to propagate, got %#v", model)
		}
	})

	t.Run("exception review required yields aggregate review required if no local blockers", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.ExceptionDecision.ExceptionType = point11ValCExceptionTypeTemporaryOverrideCandidate
		model = ComputePoint11ValCFoundation(model)
		if model.CurrentState != Point11ValCStateReviewRequired {
			t.Fatalf("expected exception review required to propagate, got %#v", model)
		}
	})

	t.Run("monitoring review required yields aggregate review required if no local blockers", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.Monitoring.SignalFreshness = point11ValCMonitoringSignalStale
		model = ComputePoint11ValCFoundation(model)
		if model.CurrentState != Point11ValCStateReviewRequired {
			t.Fatalf("expected monitoring review required to propagate, got %#v", model)
		}
	})

	t.Run("any local component blocked yields aggregate blocked", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.Dashboard.DashboardViewID = ""
		model = ComputePoint11ValCFoundation(model)
		if model.CurrentState != Point11ValCStateBlocked {
			t.Fatalf("expected local dashboard blocker to block aggregate, got %#v", model)
		}
	})

	t.Run("blocked overrides review required", func(t *testing.T) {
		model := Point11ValCFoundationModel()
		model.Dependency = reviewRequiredPoint11ValCDependencySnapshot()
		model.Monitoring.SignalFreshness = point11ValCMonitoringSignalStale
		model.EnforcementInput.EnforcementID = ""
		model = ComputePoint11ValCFoundation(model)
		if model.CurrentState != Point11ValCStateBlocked {
			t.Fatalf("expected local blocked state to override review required, got %#v", model)
		}
	})

	t.Run("diagnostics include component blocking reason", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.EnforcementInput.EnforcementID = ""
		model = ComputePoint11ValCFoundation(model)
		if !point11Val0ContainsTrimmed(model.Diagnostics.BlockingReasons, "enforcement_input_blocked") {
			t.Fatalf("expected blocking reasons to include enforcement input blocker, got %#v", model.Diagnostics)
		}
	})

	t.Run("aggregate does not emit point11 pass marker", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		if strings.Contains(string(body), point11ValCPassToken()) {
			t.Fatalf("unexpected pass token in valc aggregate output: %s", string(body))
		}
	})

	t.Run("aggregate does not create legal regulatory certification authority", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.CreatesAuthorityClaims {
			t.Fatalf("expected no authority claims on active model, got %#v", model)
		}
	})

	t.Run("aggregate does not create publication side effects", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.CreatesPublicationSideEffects {
			t.Fatalf("expected no publication side effects on active model, got %#v", model)
		}
	})

	t.Run("aggregate does not create real enforcement side effects", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		if model.CreatesRealEnforcementSideEffects {
			t.Fatalf("expected no real enforcement side effects on active model, got %#v", model)
		}
	})

	t.Run("aggregate blocks padded active monitoring state instead of normalizing", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.MonitoringState = "\t" + Point11ValCMonitoringStateActive + "\n"
		if got := EvaluatePoint11ValCFoundationState(model); got != Point11ValCStateBlocked {
			t.Fatalf("expected padded monitoring state to block aggregate, got %q for %#v", got, model)
		}
		if !point11Val0ContainsTrimmed(point11ValCBlockingReasons(model), "monitoring_blocked") {
			t.Fatalf("expected exact monitoring blocking reason, got %#v", point11ValCBlockingReasons(model))
		}
	})
}

func TestPoint11ValCSemanticAntiGreenRefs(t *testing.T) {
	t.Run("enforcement unknown blocks even though it has a valid prefix", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.EnforcementInput.EnforcementID = "enforcement_unknown"
		model = ComputePoint11ValCFoundation(model)
		if model.EnforcementInputState != Point11ValCEnforcementInputStateBlocked {
			t.Fatalf("expected enforcement_unknown to block, got %#v", model)
		}
	})

	t.Run("decision revoked blocks even though it has a valid prefix", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.EnforcementInput.DecisionID = "decision_revoked"
		model = ComputePoint11ValCFoundation(model)
		if model.EnforcementInputState != Point11ValCEnforcementInputStateBlocked {
			t.Fatalf("expected decision_revoked to block, got %#v", model)
		}
	})

	t.Run("abac invalid blocks even though it has a valid prefix", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.ABACDecision.ABACDecisionID = "abac_invalid"
		model = ComputePoint11ValCFoundation(model)
		if model.ABACDecisionState != Point11ValCABACDecisionStateBlocked {
			t.Fatalf("expected abac_invalid to block, got %#v", model)
		}
	})

	t.Run("exception expired blocks even though it has a valid prefix", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.ExceptionDecision.ExceptionRef = "exception_expired"
		model = ComputePoint11ValCFoundation(model)
		if model.ExceptionDecisionState != Point11ValCExceptionDecisionStateBlocked {
			t.Fatalf("expected exception_expired to block, got %#v", model)
		}
	})

	t.Run("emergency placeholder blocks even though it has a valid prefix", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.ExceptionDecision.EmergencyRef = "emergency_placeholder"
		model = ComputePoint11ValCFoundation(model)
		if model.ExceptionDecisionState != Point11ValCExceptionDecisionStateBlocked {
			t.Fatalf("expected emergency_placeholder to block, got %#v", model)
		}
	})

	t.Run("monitoring stale blocks even though it has a valid prefix", func(t *testing.T) {
		model := activePoint11ValCFoundation()
		model.Monitoring.MonitoringLinkID = "monitoring_stale"
		model = ComputePoint11ValCFoundation(model)
		if model.MonitoringState != Point11ValCMonitoringStateBlocked {
			t.Fatalf("expected monitoring_stale to block, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValCFoundation)
	}{
		{name: "revoked invalid marker blocks enforcement ref", mutate: func(model *Point11ValCFoundation) {
			model.EnforcementInput.EnforcementID = "revoked/invalid marker"
		}},
		{name: "revoked invalid marker blocks decision ref", mutate: func(model *Point11ValCFoundation) {
			model.EnforcementInput.DecisionID = "revoked/invalid marker"
		}},
		{name: "revoked invalid marker blocks exception ref", mutate: func(model *Point11ValCFoundation) {
			model.ExceptionDecision.ExceptionRef = "revoked/invalid marker"
		}},
		{name: "revoked invalid marker blocks emergency ref", mutate: func(model *Point11ValCFoundation) {
			model.ExceptionDecision.EmergencyRef = "revoked/invalid marker"
		}},
		{name: "revoked invalid marker blocks monitoring ref", mutate: func(model *Point11ValCFoundation) {
			model.Monitoring.MonitoringLinkID = "revoked/invalid marker"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValCFoundation()
			testCase.mutate(&model)
			model = ComputePoint11ValCFoundation(model)
			if model.CurrentState != Point11ValCStateBlocked {
				t.Fatalf("expected revoked/invalid marker to block valc, got %#v", model)
			}
		})
	}
}
