package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	ossTrustNetworkValAStatusSchema = "point9.oss_trust_network.vala.status.v1"
	ossTrustNetworkValAProofsSchema = "point9.oss_trust_network.vala.proofs.v1"
)

type ossTrustNetworkValAStatusResponse struct {
	SchemaVersion string                              `json:"schema_version"`
	GeneratedAt   time.Time                           `json:"generated_at"`
	CurrentState  string                              `json:"current_state"`
	Model         operability.OSSTrustNetworkValACore `json:"model"`
	RouteRefs     []string                            `json:"route_refs,omitempty"`
	Limitations   []string                            `json:"limitations,omitempty"`
}

type ossTrustNetworkValAProofsResponse struct {
	SchemaVersion                string    `json:"schema_version"`
	GeneratedAt                  time.Time `json:"generated_at"`
	CurrentState                 string    `json:"current_state"`
	Point9State                  string    `json:"point_9_state"`
	DependencyState              string    `json:"dependency_state"`
	Val0CurrentState             string    `json:"val0_current_state"`
	Val0Point9State              string    `json:"val0_point_9_state"`
	Val0DependencyState          string    `json:"val0_dependency_state"`
	Val0NoOverclaimState         string    `json:"val0_no_overclaim_state"`
	Val0Point8State              string    `json:"val0_point_8_state"`
	Val0Point8PassAllowed        bool      `json:"val0_point_8_pass_allowed"`
	Val0Point8PassReason         string    `json:"val0_point_8_pass_reason"`
	Val0Point8ClosureState       string    `json:"val0_point_8_closure_state"`
	ReleaseTrustIntakeState      string    `json:"release_trust_intake_state"`
	SigningSignalState           string    `json:"signing_signal_state"`
	MaintainerAttestationState   string    `json:"maintainer_attestation_state"`
	ProvenanceMaterialState      string    `json:"provenance_material_state"`
	RegistryDescriptorState      string    `json:"registry_descriptor_state"`
	RegistryMetadataState        string    `json:"registry_metadata_state"`
	TypoSquattingWarningState    string    `json:"typo_squatting_warning_state"`
	DriftSignalState             string    `json:"drift_signal_state"`
	NoOverclaimState             string    `json:"no_overclaim_state"`
	ReleaseTrustReviewState      string    `json:"release_trust_review_state"`
	ReleaseTrustFreshness        string    `json:"release_trust_freshness"`
	SigningState                 string    `json:"signing_state"`
	MaintainerAttestation        string    `json:"maintainer_attestation"`
	ProvenanceState              string    `json:"provenance_state"`
	RegistryDescriptor           string    `json:"registry_descriptor"`
	RegistryMetadataFreshness    string    `json:"registry_metadata_freshness"`
	TypoSquattingReviewState     string    `json:"typo_squatting_review_state"`
	DriftClass                   string    `json:"drift_class"`
	DriftState                   string    `json:"drift_state"`
	SupportedSigningStates       []string  `json:"supported_signing_states,omitempty"`
	SupportedAttestationStates   []string  `json:"supported_attestation_states,omitempty"`
	SupportedProvenanceStates    []string  `json:"supported_provenance_states,omitempty"`
	SupportedRegistryDescriptors []string  `json:"supported_registry_descriptors,omitempty"`
	SupportedDriftClasses        []string  `json:"supported_drift_classes,omitempty"`
	SurfaceRefs                  []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                 []string  `json:"evidence_refs,omitempty"`
	BlockingReasons              []string  `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete         []string  `json:"why_point_9_not_complete,omitempty"`
	Limitations                  []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer         string    `json:"projection_disclaimer"`
	IntegrationSummary           []string  `json:"integration_summary,omitempty"`
}

func ossTrustNetworkValAAllSurfaceRefs() []string {
	return operability.OSSTrustNetworkValAProofSurfaceRefs()
}

func ossTrustNetworkValASupportedSigningStates() []string {
	return []string{
		operability.OSSTrustNetworkValASigningStatePresent,
		operability.OSSTrustNetworkValASigningStateVerified,
		operability.OSSTrustNetworkValASigningStateMissing,
		operability.OSSTrustNetworkValASigningStateMismatch,
		operability.OSSTrustNetworkValASigningStateRevoked,
		operability.OSSTrustNetworkValASigningStateUnsupported,
		operability.OSSTrustNetworkValASigningStateUnknown,
	}
}

func ossTrustNetworkValASupportedAttestationStates() []string {
	return []string{
		operability.OSSTrustNetworkValAAttestationStateAttested,
		operability.OSSTrustNetworkValAAttestationStateMissing,
		operability.OSSTrustNetworkValAAttestationStateStale,
		operability.OSSTrustNetworkValAAttestationStateRevoked,
		operability.OSSTrustNetworkValAAttestationStateDelegated,
		operability.OSSTrustNetworkValAAttestationStateUnsupported,
		operability.OSSTrustNetworkValAAttestationStateUnknown,
	}
}

func ossTrustNetworkValASupportedProvenanceStates() []string {
	return []string{
		operability.OSSTrustNetworkValAProvenanceStateVerified,
		operability.OSSTrustNetworkValAProvenanceStatePresentUnverified,
		operability.OSSTrustNetworkValAProvenanceStateMissing,
		operability.OSSTrustNetworkValAProvenanceStateMismatch,
		operability.OSSTrustNetworkValAProvenanceStateStale,
		operability.OSSTrustNetworkValAProvenanceStateUnsupported,
		operability.OSSTrustNetworkValAProvenanceStateUnknown,
	}
}

func ossTrustNetworkValASupportedDriftClasses() []string {
	return []string{
		operability.OSSTrustNetworkValADriftClassMaintainer,
		operability.OSSTrustNetworkValADriftClassProvenance,
		operability.OSSTrustNetworkValADriftClassSigning,
		operability.OSSTrustNetworkValADriftClassRegistryMetadata,
		operability.OSSTrustNetworkValADriftClassSuspiciousRelease,
	}
}

func buildOSSTrustNetworkValAModel() operability.OSSTrustNetworkValACore {
	model := operability.OSSTrustNetworkValACoreModel()
	return operability.ComputeOSSTrustNetworkValACore(model)
}

func (s server) ossTrustNetworkValAStatusHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValAStatus())
}

func (s server) ossTrustNetworkValAProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildOSSTrustNetworkValAProofs())
}

func buildOSSTrustNetworkValAStatus() ossTrustNetworkValAStatusResponse {
	model := buildOSSTrustNetworkValAModel()
	limitations := []string{
		"Val A defines bounded release trust and registry core only and does not implement shared reviewed intelligence workflows, dashboards, final closure, or Točka 10.",
		"Registry descriptors remain descriptor-only in Val A and do not perform live network trust fetches or create canonical truth.",
	}
	return ossTrustNetworkValAStatusResponse{
		SchemaVersion: ossTrustNetworkValAStatusSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs:     ossTrustNetworkValAAllSurfaceRefs(),
		Limitations:   limitations,
	}
}

func buildOSSTrustNetworkValAProofs() ossTrustNetworkValAProofsResponse {
	model := buildOSSTrustNetworkValAModel()
	limitations := []string{
		"Val A keeps Točka 9 incomplete and reserves any final pass semantics for later integrated closure waves only.",
		"Release, signing, maintainer, provenance, registry, typo-warning, and drift outputs remain bounded advisory signals rather than canonical truth, certification, or approval authority.",
		"Val B through Val E and Točka 10 remain out of scope here.",
	}
	currentState := operability.EvaluateOSSTrustNetworkValAProofsState(model, limitations)
	return ossTrustNetworkValAProofsResponse{
		SchemaVersion:                ossTrustNetworkValAProofsSchema,
		GeneratedAt:                  publicSampleTime(),
		CurrentState:                 currentState,
		Point9State:                  model.Point9State,
		DependencyState:              model.DependencyState,
		Val0CurrentState:             model.Dependency.Val0CurrentState,
		Val0Point9State:              model.Dependency.Val0Point9State,
		Val0DependencyState:          model.Dependency.Val0DependencyState,
		Val0NoOverclaimState:         model.Dependency.Val0NoOverclaimState,
		Val0Point8State:              model.Dependency.Val0Point8State,
		Val0Point8PassAllowed:        model.Dependency.Val0Point8PassAllowed,
		Val0Point8PassReason:         model.Dependency.Val0Point8PassReason,
		Val0Point8ClosureState:       model.Dependency.Val0Point8ClosureState,
		ReleaseTrustIntakeState:      model.ReleaseTrustIntakeState,
		SigningSignalState:           model.SigningSignalState,
		MaintainerAttestationState:   model.MaintainerAttestationState,
		ProvenanceMaterialState:      model.ProvenanceMaterialState,
		RegistryDescriptorState:      model.RegistryDescriptorState,
		RegistryMetadataState:        model.RegistryMetadataState,
		TypoSquattingWarningState:    model.TypoSquattingWarningState,
		DriftSignalState:             model.DriftSignalState,
		NoOverclaimState:             model.NoOverclaimState,
		ReleaseTrustReviewState:      model.ReleaseTrustIntake.ReviewState,
		ReleaseTrustFreshness:        model.ReleaseTrustIntake.FreshnessState,
		SigningState:                 model.SigningSignal.SigningState,
		MaintainerAttestation:        model.MaintainerAttestation.AttestationState,
		ProvenanceState:              model.ProvenanceMaterial.ProvenanceState,
		RegistryDescriptor:           model.RegistryDescriptor.RequestedRegistryDescriptor,
		RegistryMetadataFreshness:    model.RegistryMetadata.MetadataFreshness,
		TypoSquattingReviewState:     model.TypoSquattingWarning.ReviewState,
		DriftClass:                   model.DriftSignal.DriftClass,
		DriftState:                   model.DriftSignal.DriftState,
		SupportedSigningStates:       ossTrustNetworkValASupportedSigningStates(),
		SupportedAttestationStates:   ossTrustNetworkValASupportedAttestationStates(),
		SupportedProvenanceStates:    ossTrustNetworkValASupportedProvenanceStates(),
		SupportedRegistryDescriptors: model.RegistryDescriptor.SupportedRegistryDescriptors,
		SupportedDriftClasses:        ossTrustNetworkValASupportedDriftClasses(),
		SurfaceRefs:                  model.ProofSurfaceRefs,
		EvidenceRefs:                 model.EvidenceRefs,
		BlockingReasons:              model.BlockingReasons,
		WhyPoint9NotComplete:         model.WhyPoint9NotComplete,
		Limitations:                  limitations,
		ProjectionDisclaimer:         model.ProjectionDisclaimer,
		IntegrationSummary: []string{
			"Val A adds bounded release trust intake, signing, maintainer attestation, provenance, registry descriptor, registry metadata, typo-warning, and drift disciplines on top of exact Val 0 dependency.",
			"Registry and OSS network inputs remain bounded projection surfaces and do not become canonical truth, package certification, or final Point 9 closure here.",
		},
	}
}
