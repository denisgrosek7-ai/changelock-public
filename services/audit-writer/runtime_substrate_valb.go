package main

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"time"

	attestationruntime "github.com/denisgrosek/changelock/internal/attestation"
	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	runtimeSubstrateValBCorrelationModelSchema = "runtime.substrate.valb.correlation_model.v1"
	runtimeSubstrateValBProcessImageSchema     = "runtime.substrate.valb.process_image_linkage.v1"
	runtimeSubstrateValBProvenanceSchema       = "runtime.substrate.valb.provenance_linkage.v1"
	runtimeSubstrateValBDriftCatalogSchema     = "runtime.substrate.valb.drift_catalog.v1"
	runtimeSubstrateValBProofsSchema           = "runtime.substrate.valb.proofs.v1"
	runtimeSubstrateValBCoverageScope          = "binary_and_provenance_correlation"
)

type runtimeSubstrateValBCorrelationModelResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Model         runtimesubstrate.RuntimeSubstrateCorrelationModel `json:"model"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type runtimeSubstrateValBProcessImageResponse struct {
	SchemaVersion string                                                 `json:"schema_version"`
	GeneratedAt   time.Time                                              `json:"generated_at"`
	CurrentState  string                                                 `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateProcessImageLinkage `json:"items,omitempty"`
	RouteRefs     []string                                               `json:"route_refs,omitempty"`
	Limitations   []string                                               `json:"limitations,omitempty"`
}

type runtimeSubstrateValBProvenanceResponse struct {
	SchemaVersion string                                               `json:"schema_version"`
	GeneratedAt   time.Time                                            `json:"generated_at"`
	CurrentState  string                                               `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateProvenanceLinkage `json:"items,omitempty"`
	RouteRefs     []string                                             `json:"route_refs,omitempty"`
	Limitations   []string                                             `json:"limitations,omitempty"`
}

type runtimeSubstrateValBDriftCatalogResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateDriftRecord `json:"items,omitempty"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type runtimeSubstrateValBProofsResponse struct {
	SchemaVersion          string                                                 `json:"schema_version"`
	GeneratedAt            time.Time                                              `json:"generated_at"`
	CurrentState           string                                                 `json:"current_state"`
	CoverageScope          string                                                 `json:"coverage_scope"`
	ValAState              string                                                 `json:"val_a_state"`
	CorrelationModelState  string                                                 `json:"correlation_model_state"`
	ProcessImageState      string                                                 `json:"process_image_state"`
	ProvenanceState        string                                                 `json:"provenance_state"`
	DriftCatalogState      string                                                 `json:"drift_catalog_state"`
	ProcessImageItems      []runtimesubstrate.RuntimeSubstrateProcessImageLinkage `json:"process_image_items,omitempty"`
	ProvenanceItems        []runtimesubstrate.RuntimeSubstrateProvenanceLinkage   `json:"provenance_items,omitempty"`
	DriftItems             []runtimesubstrate.RuntimeSubstrateDriftRecord         `json:"drift_items,omitempty"`
	RemainingDeferredScope []string                                               `json:"remaining_deferred_scope,omitempty"`
	RouteRefs              []string                                               `json:"route_refs,omitempty"`
	Limitations            []string                                               `json:"limitations,omitempty"`
}

type runtimeSubstrateValBBundle struct {
	ValAState         string
	ProcessImage      []runtimesubstrate.RuntimeSubstrateProcessImageLinkage
	Provenance        []runtimesubstrate.RuntimeSubstrateProvenanceLinkage
	DriftCatalog      []runtimesubstrate.RuntimeSubstrateDriftRecord
	ProcessState      string
	ProvenanceState   string
	DriftCatalogState string
}

func (s server) runtimeSubstrateValBCorrelationModelHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValBCorrelationModel())
}

func (s server) runtimeSubstrateValBProcessImageHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(req)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildRuntimeSubstrateValBProcessImage(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) runtimeSubstrateValBProvenanceHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(req)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildRuntimeSubstrateValBProvenance(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) runtimeSubstrateValBDriftCatalogHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(req)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildRuntimeSubstrateValBDriftCatalog(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) runtimeSubstrateValBProofsHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(req)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildRuntimeSubstrateValBProofs(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildRuntimeSubstrateValBCorrelationModel() runtimeSubstrateValBCorrelationModelResponse {
	model := runtimesubstrate.RuntimeSubstrateValBCorrelationModel()
	return runtimeSubstrateValBCorrelationModelResponse{
		SchemaVersion: runtimeSubstrateValBCorrelationModelSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/proofs",
			"/v1/runtime/substrate-depth/valb/process-image-linkage",
			"/v1/runtime/substrate-depth/valb/provenance-linkage",
			"/v1/runtime/substrate-depth/valb/proofs",
		},
		Limitations: []string{
			"Val B correlation model remains bounded to canonical runtime observations, artifact evidence, and phase2 attestation linkage where supportable.",
		},
	}
}

func (s server) buildRuntimeSubstrateValBProcessImage(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValBProcessImageResponse, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeSubstrateValBProcessImageResponse{}, err
	}
	bundle := runtimeSubstrateValBBundleFromSnapshot(snapshot)
	return runtimeSubstrateValBProcessImageResponse{
		SchemaVersion: runtimeSubstrateValBProcessImageSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  bundle.ProcessState,
		Items:         bundle.ProcessImage,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/observability",
			"/v1/runtime/substrate-depth/valb/provenance-linkage",
			"/v1/runtime/substrate-depth/valb/drift-catalog",
		},
		Limitations: []string{
			"Process-image linkage is only claimed where exec observations carry bounded process path and binary digest context; missing digest support remains partial or unsupported.",
			"Direct digest matches remain bounded to canonical artifact digests or attestation subject digests and do not equate workload image context with generic memory truth.",
		},
	}, nil
}

func (s server) buildRuntimeSubstrateValBProvenance(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValBProvenanceResponse, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeSubstrateValBProvenanceResponse{}, err
	}
	bundle := runtimeSubstrateValBBundleFromSnapshot(snapshot)
	return runtimeSubstrateValBProvenanceResponse{
		SchemaVersion: runtimeSubstrateValBProvenanceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  bundle.ProvenanceState,
		Items:         bundle.Provenance,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/proofs",
			"/v1/runtime/substrate-depth/valb/process-image-linkage",
			"/v1/runtime/substrate-depth/valb/drift-catalog",
		},
		Limitations: []string{
			"Provenance linkage correlates workload image and canonical artifact evidence where present; it does not claim universal subject coverage for workloads lacking signer or attestation evidence.",
			"Phase2 attestation linkage hardens trust state visibility but does not widen Val B into enforcement or confidential-computing authority claims.",
		},
	}, nil
}

func (s server) buildRuntimeSubstrateValBDriftCatalog(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValBDriftCatalogResponse, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeSubstrateValBDriftCatalogResponse{}, err
	}
	bundle := runtimeSubstrateValBBundleFromSnapshot(snapshot)
	return runtimeSubstrateValBDriftCatalogResponse{
		SchemaVersion: runtimeSubstrateValBDriftCatalogSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  bundle.DriftCatalogState,
		Items:         bundle.DriftCatalog,
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/valb/process-image-linkage",
			"/v1/runtime/substrate-depth/valb/provenance-linkage",
			"/v1/runtime/substrate-depth/valb/proofs",
		},
		Limitations: []string{
			"Val B drift classes explain expected, low-risk, suspicious, and hard mismatch outcomes only for canonical runtime and provenance evidence in scope.",
		},
	}, nil
}

func (s server) buildRuntimeSubstrateValBProofs(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValBProofsResponse, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeSubstrateValBProofsResponse{}, err
	}
	model := runtimesubstrate.RuntimeSubstrateValBCorrelationModel()
	bundle := runtimeSubstrateValBBundleFromSnapshot(snapshot)
	return runtimeSubstrateValBProofsResponse{
		SchemaVersion:          runtimeSubstrateValBProofsSchema,
		GeneratedAt:            publicSampleTime(),
		CurrentState:           runtimesubstrate.EvaluateRuntimeSubstrateValBState(bundle.ValAState, model.CurrentState, bundle.ProcessState, bundle.ProvenanceState, bundle.DriftCatalogState),
		CoverageScope:          runtimeSubstrateValBCoverageScope,
		ValAState:              bundle.ValAState,
		CorrelationModelState:  model.CurrentState,
		ProcessImageState:      bundle.ProcessState,
		ProvenanceState:        bundle.ProvenanceState,
		DriftCatalogState:      bundle.DriftCatalogState,
		ProcessImageItems:      bundle.ProcessImage,
		ProvenanceItems:        bundle.Provenance,
		DriftItems:             bundle.DriftCatalog,
		RemainingDeferredScope: runtimesubstrate.RuntimeSubstrateValBRemainingDeferredScope(),
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/vala/proofs",
			"/v1/runtime/substrate-depth/valb/correlation-model",
			"/v1/runtime/substrate-depth/valb/process-image-linkage",
			"/v1/runtime/substrate-depth/valb/provenance-linkage",
			"/v1/runtime/substrate-depth/valb/drift-catalog",
		},
		Limitations: []string{
			"Val B proofs remain fail-closed on top of an active Val A runtime baseline and explicit canonical provenance linkage.",
			"Unsupported correlation cases stay visible instead of being flattened into expected or mismatch claims.",
		},
	}, nil
}

func runtimeSubstrateValBBundleFromSnapshot(snapshot runtimeSnapshot) runtimeSubstrateValBBundle {
	model := runtimesubstrate.RuntimeSubstrateValBCorrelationModel()
	events := runtimeSubstrateValAObservedEventsFromSnapshot(snapshot)
	valAState := runtimesubstrate.EvaluateRuntimeSubstrateValAState(
		runtimesubstrate.RuntimeSubstrateEntryGateRuntimeBaseline().CurrentState,
		runtimesubstrate.RuntimeSubstrateValAEventSchema().CurrentState,
		runtimesubstrate.EvaluateRuntimeSubstrateValASupportMatrixState(runtimesubstrate.RuntimeSubstrateValASupportMatrix()),
		runtimesubstrate.EvaluateRuntimeSubstrateValAObservabilityState(events),
	)
	processItems := runtimeSubstrateValBProcessImageLinkagesFromSnapshot(snapshot)
	provenanceItems := runtimeSubstrateValBProvenanceLinkagesFromSnapshot(snapshot)
	driftItems := runtimeSubstrateValBDriftCatalogFromLinkages(processItems, provenanceItems)
	_ = model
	return runtimeSubstrateValBBundle{
		ValAState:         valAState,
		ProcessImage:      processItems,
		Provenance:        provenanceItems,
		DriftCatalog:      driftItems,
		ProcessState:      runtimesubstrate.EvaluateRuntimeSubstrateValBProcessImageState(processItems),
		ProvenanceState:   runtimesubstrate.EvaluateRuntimeSubstrateValBProvenanceState(provenanceItems),
		DriftCatalogState: runtimesubstrate.EvaluateRuntimeSubstrateValBDriftCatalogState(driftItems),
	}
}

func runtimeSubstrateValBProcessImageLinkagesFromSnapshot(snapshot runtimeSnapshot) []runtimesubstrate.RuntimeSubstrateProcessImageLinkage {
	items := []runtimesubstrate.RuntimeSubstrateProcessImageLinkage{}
	for _, subject := range snapshot.sortedSubjects() {
		for _, observation := range subject.Observations {
			event, ok := runtimeSubstrateObservedEventFromObservation(observation, subject)
			if !ok || event.EventFamily != runtimesubstrate.RuntimeSubstrateEventFamilyExecLifecycle {
				continue
			}
			currentState, driftClass, unsupportedFields, reasons := runtimeSubstrateValBProcessImageClassification(event, subject)
			items = append(items, runtimesubstrate.RuntimeSubstrateProcessImageLinkage{
				SubjectRef:                subject.SubjectRef,
				EventID:                   event.EventID,
				Process:                   event.Process,
				Workload:                  event.Workload,
				Repository:                subject.Repo,
				ArtifactDigests:           append([]string{}, subject.ArtifactDigests...),
				AttestationSubjectDigests: append([]string{}, subject.AttestationSubjectDigests...),
				CurrentState:              currentState,
				DriftClass:                driftClass,
				UnsupportedFields:         unsupportedFields,
				EvidenceRefs:              uniqueStrings(append(append([]string{}, event.EvidenceRefs...), runtimeSnapshotSubjectEvidenceRefs(subject)...)),
				Reasons:                   reasons,
			})
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].SubjectRef == items[j].SubjectRef {
			return items[i].EventID < items[j].EventID
		}
		return items[i].SubjectRef < items[j].SubjectRef
	})
	return items
}

func runtimeSubstrateValBProcessImageClassification(event runtimesubstrate.RuntimeSubstrateObservedEvent, subject *runtimeSnapshotSubject) (string, string, []string, []string) {
	provenanceDigests := uniqueStrings(append([]string{}, append(subject.ArtifactDigests, subject.AttestationSubjectDigests...)...))
	switch {
	case strings.TrimSpace(event.Process.ProcessPath) == "":
		return runtimesubstrate.RuntimeSubstrateCorrelationStateUnsupported, runtimesubstrate.RuntimeSubstrateDriftLowRisk, []string{"process.process_path"}, []string{"process_path_missing_for_exec_correlation"}
	case strings.TrimSpace(event.Process.BinaryDigest) == "":
		if len(provenanceDigests) > 0 || len(subject.ExpectedSigners) > 0 || subject.LatestAttestation != nil {
			return runtimesubstrate.RuntimeSubstrateCorrelationStatePartial, runtimesubstrate.RuntimeSubstrateDriftLowRisk, []string{"process.binary_digest"}, []string{"binary_digest_missing_for_direct_digest_match", "workload_image_context_present"}
		}
		return runtimesubstrate.RuntimeSubstrateCorrelationStateUnsupported, runtimesubstrate.RuntimeSubstrateDriftLowRisk, []string{"process.binary_digest", "artifact.digest", "artifact.attestation_subject_digest"}, []string{"direct_digest_match_not_supportable"}
	case len(provenanceDigests) == 0:
		if len(subject.ExpectedSigners) > 0 || strings.TrimSpace(subject.Repo) != "" || subject.LatestAttestation != nil {
			reasons := []string{"provenance_context_present_without_direct_digest"}
			if subject.LatestAttestation != nil {
				reasons = append(reasons, "phase2_attestation_state_"+strings.TrimSpace(subject.LatestAttestation.CurrentState))
			}
			return runtimesubstrate.RuntimeSubstrateCorrelationStatePartial, runtimesubstrate.RuntimeSubstrateDriftSuspicious, []string{"artifact.digest", "artifact.attestation_subject_digest"}, uniqueStrings(reasons)
		}
		return runtimesubstrate.RuntimeSubstrateCorrelationStateUnsupported, runtimesubstrate.RuntimeSubstrateDriftSuspicious, []string{"artifact.digest", "artifact.signer_identity", "artifact.attestation_subject_digest"}, []string{"no_provenance_context_for_binary_digest"}
	case containsString(provenanceDigests, event.Process.BinaryDigest):
		reasons := []string{"binary_digest_matches_provenance_digest"}
		if strings.TrimSpace(subject.ImageDigest) != "" {
			reasons = append(reasons, "workload_image_context_present")
		}
		return runtimesubstrate.RuntimeSubstrateCorrelationStateSupported, runtimesubstrate.RuntimeSubstrateDriftExpected, nil, uniqueStrings(reasons)
	default:
		reasons := []string{"binary_digest_mismatches_provenance_digest"}
		if subject.LatestAttestation != nil {
			reasons = append(reasons, "phase2_attestation_state_"+strings.TrimSpace(subject.LatestAttestation.CurrentState))
		}
		return runtimesubstrate.RuntimeSubstrateCorrelationStateSupported, runtimesubstrate.RuntimeSubstrateDriftHardMismatch, nil, uniqueStrings(reasons)
	}
}

func runtimeSubstrateValBProvenanceLinkagesFromSnapshot(snapshot runtimeSnapshot) []runtimesubstrate.RuntimeSubstrateProvenanceLinkage {
	items := []runtimesubstrate.RuntimeSubstrateProvenanceLinkage{}
	for _, subject := range snapshot.sortedSubjects() {
		currentState, driftClass, unsupportedFields, reasons := runtimeSubstrateValBProvenanceClassification(subject)
		item := runtimesubstrate.RuntimeSubstrateProvenanceLinkage{
			SubjectRef:                subject.SubjectRef,
			Repository:                subject.Repo,
			WorkloadImageDigest:       subject.ImageDigest,
			ArtifactDigests:           append([]string{}, subject.ArtifactDigests...),
			SignerIdentities:          append([]string{}, subject.ExpectedSigners...),
			AttestationSubjectDigests: append([]string{}, subject.AttestationSubjectDigests...),
			AttestationPredicates:     append([]string{}, subject.AttestationPredicates...),
			CurrentState:              currentState,
			DriftClass:                driftClass,
			UnsupportedFields:         unsupportedFields,
			EvidenceRefs:              runtimeSnapshotSubjectEvidenceRefs(subject),
			Reasons:                   reasons,
		}
		if subject.LatestAttestation != nil {
			item.AttestationProvider = strings.TrimSpace(subject.LatestAttestation.Provider)
			item.AttestationState = strings.TrimSpace(subject.LatestAttestation.CurrentState)
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].SubjectRef < items[j].SubjectRef
	})
	return items
}

func runtimeSubstrateValBProvenanceClassification(subject *runtimeSnapshotSubject) (string, string, []string, []string) {
	directDigests := uniqueStrings(append([]string{}, append(subject.ArtifactDigests, subject.AttestationSubjectDigests...)...))
	switch {
	case strings.TrimSpace(subject.ImageDigest) == "" && len(directDigests) == 0 && len(subject.ExpectedSigners) == 0 && subject.LatestAttestation == nil:
		return runtimesubstrate.RuntimeSubstrateCorrelationStateUnsupported, runtimesubstrate.RuntimeSubstrateDriftLowRisk, []string{"workload.image_digest", "artifact.digest", "artifact.signer_identity", "phase2.attestation.current_state"}, []string{"no_provenance_evidence_present_for_subject"}
	case strings.TrimSpace(subject.ImageDigest) != "" && len(directDigests) > 0 && containsString(directDigests, subject.ImageDigest):
		reasons := []string{"workload_image_digest_matches_provenance_digest"}
		if subject.LatestAttestation != nil {
			switch strings.TrimSpace(subject.LatestAttestation.CurrentState) {
			case attestationruntime.VerdictVerified:
				reasons = append(reasons, "phase2_attestation_verified")
				return runtimesubstrate.RuntimeSubstrateCorrelationStateSupported, runtimesubstrate.RuntimeSubstrateDriftExpected, nil, uniqueStrings(reasons)
			case attestationruntime.VerdictDegraded:
				reasons = append(reasons, "phase2_attestation_degraded")
				return runtimesubstrate.RuntimeSubstrateCorrelationStateSupported, runtimesubstrate.RuntimeSubstrateDriftLowRisk, nil, uniqueStrings(reasons)
			default:
				reasons = append(reasons, "phase2_attestation_not_trusted")
				return runtimesubstrate.RuntimeSubstrateCorrelationStateSupported, runtimesubstrate.RuntimeSubstrateDriftSuspicious, nil, uniqueStrings(reasons)
			}
		}
		return runtimesubstrate.RuntimeSubstrateCorrelationStateSupported, runtimesubstrate.RuntimeSubstrateDriftExpected, nil, uniqueStrings(reasons)
	case strings.TrimSpace(subject.ImageDigest) != "" && len(directDigests) > 0:
		reasons := []string{"workload_image_digest_mismatches_provenance_digest"}
		if subject.LatestAttestation != nil {
			reasons = append(reasons, "phase2_attestation_state_"+strings.TrimSpace(subject.LatestAttestation.CurrentState))
		}
		return runtimesubstrate.RuntimeSubstrateCorrelationStateSupported, runtimesubstrate.RuntimeSubstrateDriftHardMismatch, nil, uniqueStrings(reasons)
	case len(subject.ExpectedSigners) > 0 || subject.LatestAttestation != nil || strings.TrimSpace(subject.Repo) != "":
		reasons := []string{"signed_provenance_present_without_direct_digest_match"}
		if subject.LatestAttestation != nil {
			reasons = append(reasons, "phase2_attestation_state_"+strings.TrimSpace(subject.LatestAttestation.CurrentState))
		}
		return runtimesubstrate.RuntimeSubstrateCorrelationStatePartial, runtimesubstrate.RuntimeSubstrateDriftLowRisk, []string{"artifact.digest", "artifact.attestation_subject_digest"}, uniqueStrings(reasons)
	case strings.TrimSpace(subject.ImageDigest) != "":
		return runtimesubstrate.RuntimeSubstrateCorrelationStatePartial, runtimesubstrate.RuntimeSubstrateDriftSuspicious, []string{"artifact.digest", "artifact.signer_identity", "artifact.attestation_subject_digest"}, []string{"runtime_image_digest_present_without_provenance_evidence"}
	default:
		return runtimesubstrate.RuntimeSubstrateCorrelationStateUnsupported, runtimesubstrate.RuntimeSubstrateDriftLowRisk, []string{"artifact.digest", "artifact.signer_identity", "artifact.attestation_subject_digest"}, []string{"no_provenance_evidence_present_for_subject"}
	}
}

func runtimeSubstrateValBDriftCatalogFromLinkages(processItems []runtimesubstrate.RuntimeSubstrateProcessImageLinkage, provenanceItems []runtimesubstrate.RuntimeSubstrateProvenanceLinkage) []runtimesubstrate.RuntimeSubstrateDriftRecord {
	items := make([]runtimesubstrate.RuntimeSubstrateDriftRecord, 0, len(processItems)+len(provenanceItems))
	for _, item := range processItems {
		items = append(items, runtimesubstrate.RuntimeSubstrateDriftRecord{
			SubjectRef:   item.SubjectRef,
			SourceKind:   "process_image_linkage",
			SourceRef:    item.EventID,
			CurrentState: item.CurrentState,
			DriftClass:   item.DriftClass,
			Summary:      runtimeSubstrateValBDriftSummary("process_image_linkage", item.DriftClass),
			EvidenceRefs: append([]string{}, item.EvidenceRefs...),
			Reasons:      append([]string{}, item.Reasons...),
		})
	}
	for _, item := range provenanceItems {
		items = append(items, runtimesubstrate.RuntimeSubstrateDriftRecord{
			SubjectRef:   item.SubjectRef,
			SourceKind:   "provenance_linkage",
			SourceRef:    item.SubjectRef,
			CurrentState: item.CurrentState,
			DriftClass:   item.DriftClass,
			Summary:      runtimeSubstrateValBDriftSummary("provenance_linkage", item.DriftClass),
			EvidenceRefs: append([]string{}, item.EvidenceRefs...),
			Reasons:      append([]string{}, item.Reasons...),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].SubjectRef == items[j].SubjectRef {
			return items[i].SourceKind < items[j].SourceKind
		}
		return items[i].SubjectRef < items[j].SubjectRef
	})
	return items
}

func runtimeSubstrateValBDriftSummary(sourceKind, driftClass string) string {
	switch strings.TrimSpace(driftClass) {
	case runtimesubstrate.RuntimeSubstrateDriftExpected:
		return sourceKind + ": expected runtime linkage matches canonical provenance evidence"
	case runtimesubstrate.RuntimeSubstrateDriftLowRisk:
		return sourceKind + ": bounded linkage is present but direct digest support remains incomplete"
	case runtimesubstrate.RuntimeSubstrateDriftSuspicious:
		return sourceKind + ": runtime evidence exists without enough supporting provenance linkage"
	default:
		return sourceKind + ": runtime evidence conflicts with canonical provenance linkage"
	}
}

func runtimeSnapshotSubjectEvidenceRefs(subject *runtimeSnapshotSubject) []string {
	if subject == nil || len(subject.EvidenceRefs) == 0 {
		return nil
	}
	refs := make([]string, 0, len(subject.EvidenceRefs))
	for ref := range subject.EvidenceRefs {
		refs = append(refs, ref)
	}
	sort.Strings(refs)
	return refs
}
