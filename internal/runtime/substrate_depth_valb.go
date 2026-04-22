package runtime

import "strings"

const (
	RuntimeSubstrateCorrelationStateSupported   = "supported_correlation"
	RuntimeSubstrateCorrelationStatePartial     = "partial_correlation"
	RuntimeSubstrateCorrelationStateUnsupported = "unsupported_correlation"

	RuntimeSubstrateDriftExpected     = "expected"
	RuntimeSubstrateDriftLowRisk      = "low_risk_drift"
	RuntimeSubstrateDriftSuspicious   = "suspicious_drift"
	RuntimeSubstrateDriftHardMismatch = "hard_mismatch"

	RuntimeSubstrateValBCorrelationModelStateActive     = "runtime_substrate_valb_correlation_model_active"
	RuntimeSubstrateValBCorrelationModelStatePartial    = "runtime_substrate_valb_correlation_model_partial"
	RuntimeSubstrateValBCorrelationModelStateIncomplete = "runtime_substrate_valb_correlation_model_incomplete"

	RuntimeSubstrateValBProcessImageStateActive     = "runtime_substrate_valb_process_image_active"
	RuntimeSubstrateValBProcessImageStatePartial    = "runtime_substrate_valb_process_image_partial"
	RuntimeSubstrateValBProcessImageStateIncomplete = "runtime_substrate_valb_process_image_incomplete"

	RuntimeSubstrateValBProvenanceStateActive     = "runtime_substrate_valb_provenance_active"
	RuntimeSubstrateValBProvenanceStatePartial    = "runtime_substrate_valb_provenance_partial"
	RuntimeSubstrateValBProvenanceStateIncomplete = "runtime_substrate_valb_provenance_incomplete"

	RuntimeSubstrateValBDriftCatalogStateActive     = "runtime_substrate_valb_drift_catalog_active"
	RuntimeSubstrateValBDriftCatalogStatePartial    = "runtime_substrate_valb_drift_catalog_partial"
	RuntimeSubstrateValBDriftCatalogStateIncomplete = "runtime_substrate_valb_drift_catalog_incomplete"

	RuntimeSubstrateValBStateIncomplete  = "runtime_substrate_valb_incomplete"
	RuntimeSubstrateValBStateSubstantial = "runtime_substrate_valb_substantially_ready"
	RuntimeSubstrateValBStateActive      = "runtime_substrate_valb_active"
)

type RuntimeSubstrateCorrelationModel struct {
	CurrentState            string   `json:"current_state"`
	RequiredEvidenceInputs  []string `json:"required_evidence_inputs,omitempty"`
	DirectMatchSemantics    []string `json:"direct_match_semantics,omitempty"`
	CorrelationStates       []string `json:"correlation_states,omitempty"`
	DriftClasses            []string `json:"drift_classes,omitempty"`
	UnsupportedExpectations []string `json:"unsupported_expectations,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type RuntimeSubstrateProcessImageLinkage struct {
	SubjectRef                string           `json:"subject_ref"`
	EventID                   string           `json:"event_id"`
	Process                   ProcessIdentity  `json:"process"`
	Workload                  WorkloadIdentity `json:"workload"`
	Repository                string           `json:"repository,omitempty"`
	ArtifactDigests           []string         `json:"artifact_digests,omitempty"`
	AttestationSubjectDigests []string         `json:"attestation_subject_digests,omitempty"`
	CurrentState              string           `json:"current_state"`
	DriftClass                string           `json:"drift_class"`
	UnsupportedFields         []string         `json:"unsupported_fields,omitempty"`
	EvidenceRefs              []string         `json:"evidence_refs,omitempty"`
	Reasons                   []string         `json:"reasons,omitempty"`
}

type RuntimeSubstrateProvenanceLinkage struct {
	SubjectRef                string   `json:"subject_ref"`
	Repository                string   `json:"repository,omitempty"`
	WorkloadImageDigest       string   `json:"workload_image_digest,omitempty"`
	ArtifactDigests           []string `json:"artifact_digests,omitempty"`
	SignerIdentities          []string `json:"signer_identities,omitempty"`
	AttestationSubjectDigests []string `json:"attestation_subject_digests,omitempty"`
	AttestationPredicates     []string `json:"attestation_predicates,omitempty"`
	AttestationProvider       string   `json:"attestation_provider,omitempty"`
	AttestationState          string   `json:"attestation_state,omitempty"`
	CurrentState              string   `json:"current_state"`
	DriftClass                string   `json:"drift_class"`
	UnsupportedFields         []string `json:"unsupported_fields,omitempty"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	Reasons                   []string `json:"reasons,omitempty"`
}

type RuntimeSubstrateDriftRecord struct {
	SubjectRef   string   `json:"subject_ref"`
	SourceKind   string   `json:"source_kind"`
	SourceRef    string   `json:"source_ref,omitempty"`
	CurrentState string   `json:"current_state"`
	DriftClass   string   `json:"drift_class"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
	Reasons      []string `json:"reasons,omitempty"`
}

func RuntimeSubstrateValBCorrelationModel() RuntimeSubstrateCorrelationModel {
	model := RuntimeSubstrateCorrelationModel{
		RequiredEvidenceInputs: []string{
			"process.process_path",
			"process.binary_digest",
			"workload.image_digest",
			"artifact.digest",
			"artifact.signer_identity",
			"artifact.attestation_subject_digest",
			"phase2.attestation.current_state",
			"canonical_evidence_refs",
		},
		DirectMatchSemantics: []string{
			"binary_digest_direct_match_is_only_claimed_when_process.binary_digest matches artifact.digest or artifact.attestation_subject_digest",
			"workload_image_digest_correlation_is bounded to signed or attested provenance digests and must not be confused with generic memory-integrity truth",
			"missing binary_digest or provenance digests must stay partial or unsupported rather than inferred as expected",
		},
		CorrelationStates: []string{
			RuntimeSubstrateCorrelationStateSupported,
			RuntimeSubstrateCorrelationStatePartial,
			RuntimeSubstrateCorrelationStateUnsupported,
		},
		DriftClasses: []string{
			RuntimeSubstrateDriftExpected,
			RuntimeSubstrateDriftLowRisk,
			RuntimeSubstrateDriftSuspicious,
			RuntimeSubstrateDriftHardMismatch,
		},
		UnsupportedExpectations: []string{
			"unsupported correlation remains explicit when binary digests, signer linkage, or attestation linkage are unavailable",
			"val_b does not claim binary provenance truth for subjects lacking canonical artifact or attestation evidence",
		},
		Limitations: []string{
			"Val B correlates runtime binary and workload image context to canonical artifact and attestation evidence where supportable; it does not claim whole-system provenance truth.",
			"Direct digest match semantics remain bounded to canonical artifact digests and attestation subject digests and do not turn workload image context into generic binary-equality claims.",
		},
	}
	model.CurrentState = EvaluateRuntimeSubstrateValBCorrelationModelState(model)
	return model
}

func RuntimeSubstrateValBRemainingDeferredScope() []string {
	return []string{
		"enforcement_taxonomy_baseline",
		"execution_class_matrix_depth",
		"performance_and_proof_pack",
	}
}

func EvaluateRuntimeSubstrateValBCorrelationModelState(model RuntimeSubstrateCorrelationModel) string {
	if len(model.RequiredEvidenceInputs) == 0 || len(model.CorrelationStates) == 0 || len(model.DriftClasses) == 0 {
		return RuntimeSubstrateValBCorrelationModelStateIncomplete
	}
	partial := false
	for _, state := range []string{
		RuntimeSubstrateCorrelationStateSupported,
		RuntimeSubstrateCorrelationStatePartial,
		RuntimeSubstrateCorrelationStateUnsupported,
	} {
		if !containsString(model.CorrelationStates, state) {
			return RuntimeSubstrateValBCorrelationModelStateIncomplete
		}
	}
	for _, drift := range []string{
		RuntimeSubstrateDriftExpected,
		RuntimeSubstrateDriftLowRisk,
		RuntimeSubstrateDriftSuspicious,
		RuntimeSubstrateDriftHardMismatch,
	} {
		if !containsString(model.DriftClasses, drift) {
			return RuntimeSubstrateValBCorrelationModelStateIncomplete
		}
	}
	if len(model.DirectMatchSemantics) < 2 || len(model.UnsupportedExpectations) == 0 {
		partial = true
	}
	if partial {
		return RuntimeSubstrateValBCorrelationModelStatePartial
	}
	return RuntimeSubstrateValBCorrelationModelStateActive
}

func EvaluateRuntimeSubstrateValBProcessImageState(items []RuntimeSubstrateProcessImageLinkage) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.SubjectRef) == "" || strings.TrimSpace(item.EventID) == "" {
					return false
				}
				if !isRuntimeSubstrateCorrelationState(item.CurrentState) || !isRuntimeSubstrateDriftClass(item.DriftClass) {
					return false
				}
				if len(item.Reasons) == 0 {
					return false
				}
				if item.CurrentState == RuntimeSubstrateCorrelationStateUnsupported && len(item.UnsupportedFields) == 0 {
					return false
				}
			}
			return true
		},
		func() bool {
			for _, item := range items {
				if item.CurrentState != RuntimeSubstrateCorrelationStateUnsupported {
					return true
				}
			}
			return false
		},
		RuntimeSubstrateValBProcessImageStateIncomplete,
		RuntimeSubstrateValBProcessImageStatePartial,
		RuntimeSubstrateValBProcessImageStateActive,
	)
}

func EvaluateRuntimeSubstrateValBProvenanceState(items []RuntimeSubstrateProvenanceLinkage) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.SubjectRef) == "" {
					return false
				}
				if !isRuntimeSubstrateCorrelationState(item.CurrentState) || !isRuntimeSubstrateDriftClass(item.DriftClass) {
					return false
				}
				if len(item.Reasons) == 0 {
					return false
				}
				if item.CurrentState == RuntimeSubstrateCorrelationStateUnsupported && len(item.UnsupportedFields) == 0 {
					return false
				}
			}
			return true
		},
		func() bool {
			for _, item := range items {
				if item.CurrentState != RuntimeSubstrateCorrelationStateUnsupported {
					return true
				}
			}
			return false
		},
		RuntimeSubstrateValBProvenanceStateIncomplete,
		RuntimeSubstrateValBProvenanceStatePartial,
		RuntimeSubstrateValBProvenanceStateActive,
	)
}

func EvaluateRuntimeSubstrateValBDriftCatalogState(items []RuntimeSubstrateDriftRecord) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.SubjectRef) == "" || strings.TrimSpace(item.SourceKind) == "" || strings.TrimSpace(item.Summary) == "" {
					return false
				}
				if !isRuntimeSubstrateCorrelationState(item.CurrentState) || !isRuntimeSubstrateDriftClass(item.DriftClass) {
					return false
				}
				if len(item.Reasons) == 0 {
					return false
				}
			}
			return true
		},
		func() bool { return true },
		RuntimeSubstrateValBDriftCatalogStateIncomplete,
		RuntimeSubstrateValBDriftCatalogStatePartial,
		RuntimeSubstrateValBDriftCatalogStateActive,
	)
}

func EvaluateRuntimeSubstrateValBState(valAState, modelState, processState, provenanceState, driftState string) string {
	valAState = strings.TrimSpace(valAState)
	modelState = strings.TrimSpace(modelState)
	processState = strings.TrimSpace(processState)
	provenanceState = strings.TrimSpace(provenanceState)
	driftState = strings.TrimSpace(driftState)

	if valAState != RuntimeSubstrateValAStateActive {
		return RuntimeSubstrateValBStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{modelState, processState, provenanceState, driftState} {
		switch state {
		case RuntimeSubstrateValBCorrelationModelStateActive,
			RuntimeSubstrateValBProcessImageStateActive,
			RuntimeSubstrateValBProvenanceStateActive,
			RuntimeSubstrateValBDriftCatalogStateActive:
		case RuntimeSubstrateValBCorrelationModelStatePartial,
			RuntimeSubstrateValBProcessImageStatePartial,
			RuntimeSubstrateValBProvenanceStatePartial,
			RuntimeSubstrateValBDriftCatalogStatePartial:
			hasPartial = true
		default:
			return RuntimeSubstrateValBStateIncomplete
		}
	}
	if hasPartial {
		return RuntimeSubstrateValBStateSubstantial
	}
	return RuntimeSubstrateValBStateActive
}

func evaluateRuntimeSubstrateCorrelationItems(count int, valid func() bool, hasActionable func() bool, incompleteState, partialState, activeState string) string {
	if count == 0 {
		return incompleteState
	}
	if !valid() {
		return incompleteState
	}
	if !hasActionable() {
		return partialState
	}
	return activeState
}

func isRuntimeSubstrateCorrelationState(value string) bool {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateCorrelationStateSupported, RuntimeSubstrateCorrelationStatePartial, RuntimeSubstrateCorrelationStateUnsupported:
		return true
	default:
		return false
	}
}

func isRuntimeSubstrateDriftClass(value string) bool {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateDriftExpected, RuntimeSubstrateDriftLowRisk, RuntimeSubstrateDriftSuspicious, RuntimeSubstrateDriftHardMismatch:
		return true
	default:
		return false
	}
}
