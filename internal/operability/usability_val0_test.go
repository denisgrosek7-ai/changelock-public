package operability

import (
	"testing"

	workflowauthority "github.com/denisgrosek/changelock/internal/workflow"
)

func TestProductionUsabilityVal0StateFailsClosedWhenRequiredContractMissing(t *testing.T) {
	got := EvaluateProductionUsabilityVal0State(
		workflowauthority.EnterpriseWorkflowAuthorityValDStateActive,
		ProductionUsabilityVal0ConfigIntegrityStateActive,
		ProductionUsabilityVal0ExplainabilityStateActive,
		ProductionUsabilityVal0StatusModelStateActive,
		ProductionUsabilityVal0OperationContractStateActive,
		ProductionUsabilityVal0DecisionQualityStateActive,
		ProductionUsabilityVal0NotificationStateActive,
		ProductionUsabilityVal0PermissionRedactionStateActive,
		ProductionUsabilityVal0RecoveryStateActive,
		ProductionUsabilityVal0ActionModeStateIncomplete,
	)
	if got != ProductionUsabilityVal0StateIncomplete {
		t.Fatalf("expected incomplete state with missing action mode contract, got %q", got)
	}
}

func TestProductionUsabilityVal0ConfigIntegrityFailsClosedForInvalidOrUnsupportedState(t *testing.T) {
	model := ProductionUsabilityVal0ConfigIntegrity()
	model.CurrentValidationResult = ProductionUsabilityValidationInvalid
	if got := EvaluateProductionUsabilityVal0ConfigIntegrityState(model); got == ProductionUsabilityVal0ConfigIntegrityStateActive {
		t.Fatalf("expected non-active config integrity state for invalid config, got %q", got)
	}

	model = ProductionUsabilityVal0ConfigIntegrity()
	model.CurrentCompatibility = ProductionUsabilityCompatibilityUnsupported
	if got := EvaluateProductionUsabilityVal0ConfigIntegrityState(model); got == ProductionUsabilityVal0ConfigIntegrityStateActive {
		t.Fatalf("expected non-active config integrity state for unsupported config, got %q", got)
	}
}

func TestProductionUsabilityVal0ConfigIntegrityFailsClosedForTypoCompatibilityValue(t *testing.T) {
	model := ProductionUsabilityVal0ConfigIntegrity()
	model.CurrentCompatibility = "currnet"
	if got := EvaluateProductionUsabilityVal0ConfigIntegrityState(model); got == ProductionUsabilityVal0ConfigIntegrityStateActive {
		t.Fatalf("expected non-active config integrity state for typo compatibility value, got %q", got)
	}
}

func TestProductionUsabilityVal0ConfigIntegrityFailsClosedForTypoValidationResult(t *testing.T) {
	model := ProductionUsabilityVal0ConfigIntegrity()
	model.CurrentValidationResult = "vlaid"
	if got := EvaluateProductionUsabilityVal0ConfigIntegrityState(model); got == ProductionUsabilityVal0ConfigIntegrityStateActive {
		t.Fatalf("expected non-active config integrity state for typo validation result, got %q", got)
	}
}

func TestProductionUsabilityVal0ConfigIntegrityPassesForSupportedEnums(t *testing.T) {
	model := ProductionUsabilityVal0ConfigIntegrity()
	if got := EvaluateProductionUsabilityVal0ConfigIntegrityState(model); got != ProductionUsabilityVal0ConfigIntegrityStateActive {
		t.Fatalf("expected active config integrity state for supported enums, got %q", got)
	}
}

func TestProductionUsabilityVal0UnknownFieldsDoNotSilentlyPassUnlessExplicitlyAllowed(t *testing.T) {
	model := ProductionUsabilityVal0ConfigIntegrity()
	model.CurrentUnknownFieldPolicy = "silent_ignore"
	if got := EvaluateProductionUsabilityVal0ConfigIntegrityState(model); got == ProductionUsabilityVal0ConfigIntegrityStateActive {
		t.Fatalf("expected non-active config integrity state for silent unknown field policy, got %q", got)
	}
}

func TestProductionUsabilityVal0ExplainabilityRequiresRequiredFields(t *testing.T) {
	model := ProductionUsabilityVal0ExplainabilityContract()
	model.RequiredFields = []string{"reason_code", "severity"}
	if got := EvaluateProductionUsabilityVal0ExplainabilityState(model); got == ProductionUsabilityVal0ExplainabilityStateActive {
		t.Fatalf("expected non-active explainability state without full required fields, got %q", got)
	}
}

func TestProductionUsabilityVal0PartnerAndPublicVisibilityStayRedactedWithoutLeakage(t *testing.T) {
	model := ProductionUsabilityVal0PermissionRedactionContract()
	if got := EvaluateProductionUsabilityVal0PermissionRedactionState(model); got != ProductionUsabilityVal0PermissionRedactionStateActive {
		t.Fatalf("expected active permission redaction contract, got %q", got)
	}
}

func TestProductionUsabilityVal0PermissionRedactionFailsClosedForUnsupportedEvidenceVisibility(t *testing.T) {
	model := ProductionUsabilityVal0PermissionRedactionContract()
	for idx := range model.Items {
		if model.Items[idx].VisibilityScope == ProductionUsabilityVisibilityDeveloper {
			model.Items[idx].EvidenceVisibility = "metdata_only"
			break
		}
	}
	if got := EvaluateProductionUsabilityVal0PermissionRedactionState(model); got == ProductionUsabilityVal0PermissionRedactionStateActive {
		t.Fatalf("expected non-active permission redaction state for unsupported evidence visibility, got %q", got)
	}
}

func TestProductionUsabilityVal0PermissionRedactionPassesForSupportedEvidenceVisibility(t *testing.T) {
	model := ProductionUsabilityVal0PermissionRedactionContract()
	if got := EvaluateProductionUsabilityVal0PermissionRedactionState(model); got != ProductionUsabilityVal0PermissionRedactionStateActive {
		t.Fatalf("expected active permission redaction state for supported evidence visibility, got %q", got)
	}
}

func TestProductionUsabilityVal0PartnerOrPublicScopeCannotExposeFullEvidence(t *testing.T) {
	model := ProductionUsabilityVal0PermissionRedactionContract()
	for idx := range model.Items {
		if model.Items[idx].VisibilityScope == ProductionUsabilityVisibilityPartner {
			model.Items[idx].EvidenceVisibility = ProductionUsabilityEvidenceFull
			break
		}
	}
	if got := EvaluateProductionUsabilityVal0PermissionRedactionState(model); got == ProductionUsabilityVal0PermissionRedactionStateActive {
		t.Fatalf("expected non-active permission redaction state when partner scope gets full evidence visibility, got %q", got)
	}
}

func TestProductionUsabilityVal0RedactionDoesNotConvertFailIntoPass(t *testing.T) {
	model := ProductionUsabilityVal0PermissionRedactionContract()
	model.PreservesFailureSemantics = false
	if got := EvaluateProductionUsabilityVal0PermissionRedactionState(model); got == ProductionUsabilityVal0PermissionRedactionStateActive {
		t.Fatalf("expected non-active permission redaction state when failure semantics are not preserved, got %q", got)
	}
}

func TestProductionUsabilityVal0StatusModelRequiresDistinctStates(t *testing.T) {
	model := ProductionUsabilityVal0OperationalStatusModel()
	if got := EvaluateProductionUsabilityVal0StatusModelState(model); got != ProductionUsabilityVal0StatusModelStateActive {
		t.Fatalf("expected active status model, got %q", got)
	}
}

func TestProductionUsabilityVal0StaleOrPartialProjectionCannotClaimCanonicalTruth(t *testing.T) {
	model := ProductionUsabilityVal0OperationalStatusModel()
	for idx := range model.Items {
		if model.Items[idx].Status == ProductionUsabilityStatusStale {
			model.Items[idx].ClaimsCanonicalTruth = true
			break
		}
	}
	if got := EvaluateProductionUsabilityVal0StatusModelState(model); got == ProductionUsabilityVal0StatusModelStateActive {
		t.Fatalf("expected non-active status model when stale projection claims canonical truth, got %q", got)
	}
}

func TestProductionUsabilityVal0MutatingOperationsAreNotAssumedRetrySafe(t *testing.T) {
	model := ProductionUsabilityVal0OperationContractModel()
	for idx := range model.Items {
		if model.Items[idx].OperationType == ProductionUsabilityOperationNonIdempotentMutate {
			model.Items[idx].RetrySafety = ProductionUsabilityRetrySafe
			break
		}
	}
	if got := EvaluateProductionUsabilityVal0OperationContractState(model); got == ProductionUsabilityVal0OperationContractStateActive {
		t.Fatalf("expected non-active operation contract when non-idempotent mutation is retry-safe, got %q", got)
	}
}

func TestProductionUsabilityVal0PreviewDryRunAuditOnlyStayNonMutating(t *testing.T) {
	model := ProductionUsabilityVal0ActionModeTaxonomy()
	for idx := range model.Items {
		if model.Items[idx].ActionMode == ProductionUsabilityActionModeDryRun {
			model.Items[idx].MutatesCanonicalState = true
			break
		}
	}
	if got := EvaluateProductionUsabilityVal0ActionModeState(model); got == ProductionUsabilityVal0ActionModeStateActive {
		t.Fatalf("expected non-active action mode taxonomy when dry_run mutates canonical state, got %q", got)
	}
}

func TestProductionUsabilityVal0AcknowledgementDoesNotEqualCanonicalClosure(t *testing.T) {
	model := ProductionUsabilityVal0NotificationTaxonomyContract()
	model.ResolvedEqualsClosure = true
	if got := EvaluateProductionUsabilityVal0NotificationState(model); got == ProductionUsabilityVal0NotificationStateActive {
		t.Fatalf("expected non-active notification taxonomy when resolved equals closure, got %q", got)
	}
}

func TestProductionUsabilityVal0NoiseSuppressionDoesNotHideCriticalBlockers(t *testing.T) {
	model := ProductionUsabilityVal0NotificationTaxonomyContract()
	model.CriticalSuppressionAllowed = true
	if got := EvaluateProductionUsabilityVal0NotificationState(model); got == ProductionUsabilityVal0NotificationStateActive {
		t.Fatalf("expected non-active notification taxonomy when critical suppression is allowed, got %q", got)
	}
}

func TestProductionUsabilityVal0RecoveryGuidanceDoesNotSuggestUnsafeRetryForRetryUnsafeOperations(t *testing.T) {
	model := ProductionUsabilityVal0RecoveryUXContract()
	model.UnsafeRetryDenied = false
	if got := EvaluateProductionUsabilityVal0RecoveryState(model); got == ProductionUsabilityVal0RecoveryStateActive {
		t.Fatalf("expected non-active recovery contract when unsafe retry is not denied, got %q", got)
	}
}

func TestProductionUsabilityVal0ProofPassesOnlyAsFoundationWhilePoint4RemainsNotComplete(t *testing.T) {
	got := EvaluateProductionUsabilityVal0ProofsState(
		workflowauthority.EnterpriseWorkflowAuthorityValDStateActive,
		ProductionUsabilityVal0ConfigIntegrityStateActive,
		ProductionUsabilityVal0ExplainabilityStateActive,
		ProductionUsabilityVal0StatusModelStateActive,
		ProductionUsabilityVal0OperationContractStateActive,
		ProductionUsabilityVal0DecisionQualityStateActive,
		ProductionUsabilityVal0NotificationStateActive,
		ProductionUsabilityVal0PermissionRedactionStateActive,
		ProductionUsabilityVal0RecoveryStateActive,
		ProductionUsabilityVal0ActionModeStateActive,
		[]string{
			"/v1/production/usability-operability-recovery/val0/config-integrity",
			"/v1/production/usability-operability-recovery/val0/explainability-contract",
			"/v1/production/usability-operability-recovery/val0/status-model",
			"/v1/production/usability-operability-recovery/val0/operation-contracts",
			"/v1/production/usability-operability-recovery/val0/decision-quality",
			"/v1/production/usability-operability-recovery/val0/notification-taxonomy",
			"/v1/production/usability-operability-recovery/val0/permission-redaction",
			"/v1/production/usability-operability-recovery/val0/recovery-contract",
			"/v1/production/usability-operability-recovery/val0/action-modes",
			"/v1/production/usability-operability-recovery/val0/proofs",
		},
		[]string{
			"/v1/enterprise/workflow-authority/vald/proofs",
			"/v1/enterprise/phase4/proofs",
			"config_integrity_contract",
			"explainability_contract",
			"status_model",
			"operation_contracts",
		},
		[]string{"Val 0 is a foundation slice only."},
		[]string{"Point 4 full PASS remains closed until later waves land."},
	)
	if got != ProductionUsabilityVal0StateActive {
		t.Fatalf("expected active Val 0 proofs state for complete foundation, got %q", got)
	}
	if ProductionUsabilityPoint4StateNotComplete == ProductionUsabilityVal0StateActive {
		t.Fatalf("point 4 state must remain distinct from active Val 0 foundation")
	}
}

func TestProductionUsabilityVal0MissingProofMetadataKeepsVal0NonActive(t *testing.T) {
	got := EvaluateProductionUsabilityVal0ProofsState(
		workflowauthority.EnterpriseWorkflowAuthorityValDStateActive,
		ProductionUsabilityVal0ConfigIntegrityStateActive,
		ProductionUsabilityVal0ExplainabilityStateActive,
		ProductionUsabilityVal0StatusModelStateActive,
		ProductionUsabilityVal0OperationContractStateActive,
		ProductionUsabilityVal0DecisionQualityStateActive,
		ProductionUsabilityVal0NotificationStateActive,
		ProductionUsabilityVal0PermissionRedactionStateActive,
		ProductionUsabilityVal0RecoveryStateActive,
		ProductionUsabilityVal0ActionModeStateActive,
		nil,
		[]string{"evidence"},
		[]string{"limitation"},
		[]string{"Point 4 full PASS remains closed until later waves land."},
	)
	if got == ProductionUsabilityVal0StateActive {
		t.Fatalf("expected non-active Val 0 proofs state without metadata, got %q", got)
	}
}
