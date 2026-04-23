package claims

import (
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	MeasuredPublicProofValAArtifactSchemaStateActive     = "measured_public_proof_vala_artifact_schema_active"
	MeasuredPublicProofValAArtifactSchemaStatePartial    = "measured_public_proof_vala_artifact_schema_partial"
	MeasuredPublicProofValAArtifactSchemaStateIncomplete = "measured_public_proof_vala_artifact_schema_incomplete"

	MeasuredPublicProofValASealingDisciplineStateActive     = "measured_public_proof_vala_sealing_discipline_active"
	MeasuredPublicProofValASealingDisciplineStatePartial    = "measured_public_proof_vala_sealing_discipline_partial"
	MeasuredPublicProofValASealingDisciplineStateIncomplete = "measured_public_proof_vala_sealing_discipline_incomplete"

	MeasuredPublicProofValAEnvironmentBindingStateActive     = "measured_public_proof_vala_environment_binding_active"
	MeasuredPublicProofValAEnvironmentBindingStatePartial    = "measured_public_proof_vala_environment_binding_partial"
	MeasuredPublicProofValAEnvironmentBindingStateIncomplete = "measured_public_proof_vala_environment_binding_incomplete"

	MeasuredPublicProofValADownloadablePackStateActive     = "measured_public_proof_vala_downloadable_packs_active"
	MeasuredPublicProofValADownloadablePackStatePartial    = "measured_public_proof_vala_downloadable_packs_partial"
	MeasuredPublicProofValADownloadablePackStateIncomplete = "measured_public_proof_vala_downloadable_packs_incomplete"

	MeasuredPublicProofValAStateIncomplete  = "measured_public_proof_vala_incomplete"
	MeasuredPublicProofValAStateSubstantial = "measured_public_proof_vala_substantially_ready"
	MeasuredPublicProofValAStateActive      = "measured_public_proof_vala_active"

	PublicProofArtifactTypeBenchmarkPack = "sealed_benchmark_pack"
	PublicProofArtifactTypeProofPack     = "sealed_proof_pack"
)

type PublicSealedArtifactSchema struct {
	CurrentState             string   `json:"current_state"`
	SchemaID                 string   `json:"schema_id"`
	ArtifactTypes            []string `json:"artifact_types,omitempty"`
	RequiredFields           []string `json:"required_fields,omitempty"`
	RequiredDigestAlgorithms []string `json:"required_digest_algorithms,omitempty"`
	RequiredPackagingFiles   []string `json:"required_packaging_files,omitempty"`
	RequiredSignatureFields  []string `json:"required_signature_fields,omitempty"`
	RequiredMeasurementRefs  []string `json:"required_measurement_refs,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type PublicProofSealingDiscipline struct {
	CurrentState         string                     `json:"current_state"`
	SigningPurpose       string                     `json:"signing_purpose"`
	PurposeEnabled       bool                       `json:"purpose_enabled"`
	Provider             signing.ProviderDescriptor `json:"provider"`
	RequiredTrustFields  []string                   `json:"required_trust_fields,omitempty"`
	TimestampPolicy      []string                   `json:"timestamp_policy,omitempty"`
	PackagingRules       []string                   `json:"packaging_rules,omitempty"`
	RevocationRules      []string                   `json:"revocation_rules,omitempty"`
	SupportedKeyClasses  []string                   `json:"supported_key_classes,omitempty"`
	RequiredArtifactRefs []string                   `json:"required_artifact_refs,omitempty"`
	Limitations          []string                   `json:"limitations,omitempty"`
}

type PublicProofEnvironmentBindingItem struct {
	ArtifactID         string   `json:"artifact_id"`
	CurrentState       string   `json:"current_state"`
	ClaimClass         string   `json:"claim_class"`
	RedactionTier      string   `json:"redaction_tier"`
	EnvironmentClass   string   `json:"environment_class"`
	ExecutionProfile   string   `json:"execution_profile"`
	WorkloadShape      string   `json:"workload_shape"`
	BuildIdentity      string   `json:"build_identity"`
	HarnessVersion     string   `json:"harness_version"`
	MethodologyRef     string   `json:"methodology_ref"`
	CompatibilityScope string   `json:"compatibility_scope"`
	ProvenanceInputs   []string `json:"provenance_inputs,omitempty"`
	ReplayBoundaries   []string `json:"replay_boundaries,omitempty"`
	UnsupportedReplay  []string `json:"unsupported_replay_cases,omitempty"`
	Limitations        []string `json:"limitations,omitempty"`
}

type PublicSealedProofPackFile struct {
	Path         string `json:"path"`
	MediaType    string `json:"media_type"`
	Role         string `json:"role"`
	SHA256       string `json:"sha256"`
	AdvisoryOnly bool   `json:"advisory_only,omitempty"`
}

type PublicSealedProofPack struct {
	ArtifactID            string                      `json:"artifact_id"`
	ArtifactSchemaVersion string                      `json:"artifact_schema_version"`
	ArtifactType          string                      `json:"artifact_type"`
	CurrentState          string                      `json:"current_state"`
	ClaimID               string                      `json:"claim_id"`
	ClaimClass            string                      `json:"claim_class"`
	RedactionTier         string                      `json:"redaction_tier"`
	EnvironmentClass      string                      `json:"environment_class"`
	ExecutionProfile      string                      `json:"execution_profile"`
	WorkloadShape         string                      `json:"workload_shape"`
	BuildIdentity         string                      `json:"build_identity"`
	HarnessVersion        string                      `json:"harness_version"`
	MethodologyRef        string                      `json:"methodology_ref"`
	IssuedAt              time.Time                   `json:"issued_at"`
	ValidThrough          time.Time                   `json:"valid_through"`
	MeasurementSource     string                      `json:"measurement_source"`
	EvidenceRefs          []string                    `json:"evidence_refs,omitempty"`
	DownloadRef           string                      `json:"download_ref"`
	PayloadDigest         string                      `json:"payload_digest"`
	SignatureEnvelope     *signing.Envelope           `json:"signature_envelope,omitempty"`
	TrustRootID           string                      `json:"trust_root_id"`
	KeyVersion            string                      `json:"key_version"`
	TimestampRef          string                      `json:"timestamp_ref"`
	PackagingFiles        []PublicSealedProofPackFile `json:"packaging_files,omitempty"`
	MetricSummaries       []string                    `json:"metric_summaries,omitempty"`
	Limitations           []string                    `json:"limitations,omitempty"`
}

func MeasuredPublicProofValAArtifactSchema() PublicSealedArtifactSchema {
	model := PublicSealedArtifactSchema{
		SchemaID:      "public.proof.sealed_artifact.v1",
		ArtifactTypes: []string{PublicProofArtifactTypeBenchmarkPack, PublicProofArtifactTypeProofPack},
		RequiredFields: []string{
			"artifact_id",
			"artifact_schema_version",
			"artifact_type",
			"claim_id",
			"claim_class",
			"redaction_tier",
			"environment_class",
			"execution_profile",
			"workload_shape",
			"build_identity",
			"harness_version",
			"methodology_ref",
			"issued_at",
			"valid_through",
			"measurement_source",
			"payload_digest",
			"signature_envelope",
			"trust_root_id",
			"key_version",
			"timestamp_ref",
			"evidence_refs",
			"packaging_files",
		},
		RequiredDigestAlgorithms: []string{"sha256"},
		RequiredPackagingFiles: []string{
			"manifest.json",
			"payload/measurement_summary.json",
			"environment/binding.json",
			"signature/envelope.json",
			"timestamp/receipt.json",
		},
		RequiredSignatureFields: []string{
			"provider",
			"key_id",
			"algorithm",
			"purpose",
			"payload_digest",
			"signature",
			"signed_at",
		},
		RequiredMeasurementRefs: []string{
			"measurement_basis",
			"measurement_source",
			"methodology_ref",
			"environment_class",
		},
		Limitations: []string{
			"Val A defines sealed proof artifact structure and sample pack packaging before later Point 2 waves add transparency anchoring and verifier-side signature validation.",
		},
	}
	model.CurrentState = EvaluateMeasuredPublicProofValAArtifactSchemaState(model)
	return model
}

func MeasuredPublicProofValASealingDiscipline(provider signing.ProviderDescriptor, purposeEnabled bool) PublicProofSealingDiscipline {
	model := PublicProofSealingDiscipline{
		SigningPurpose: signing.PurposePublicProofArtifact,
		PurposeEnabled: purposeEnabled,
		Provider:       provider,
		RequiredTrustFields: []string{
			"provider_mode",
			"trust_boundary",
			"key_id",
			"key_version",
			"trust_root_id",
			"active_lifecycle_state",
		},
		TimestampPolicy: []string{
			"sealed public proof artifacts require issuer-visible timestamp metadata before activation",
			"timestamp receipts remain integrity and freshness linkage, not a replacement for methodology boundaries",
		},
		PackagingRules: []string{
			"artifacts must publish manifest, payload summary, environment binding, signature envelope, and timestamp receipt together",
			"downloadable pack projection remains evidence-spine-linked and does not become a new truth store",
		},
		RevocationRules: []string{
			"revoked signer or trust-root state blocks new sealed artifact issuance",
			"previously issued artifacts transition through restricted, superseded, or withdrawn states instead of silent mutation",
		},
		SupportedKeyClasses:  []string{signing.KeyClassSealing, signing.KeyClassVerificationRoot},
		RequiredArtifactRefs: []string{"/v1/public/phase6/proofs", "/v1/public/proof-expansion/val0/proofs"},
		Limitations: uniqueStrings(append([]string{
			"Val A sealing discipline requires purpose-scoped signer enablement before downloadable packs can become fully sealed.",
		}, provider.Limitations...)),
	}
	model.CurrentState = EvaluateMeasuredPublicProofValASealingDisciplineState(model)
	return model
}

func EvaluateMeasuredPublicProofValAArtifactSchemaState(model PublicSealedArtifactSchema) string {
	if strings.TrimSpace(model.SchemaID) == "" || len(model.ArtifactTypes) == 0 || len(model.RequiredFields) == 0 {
		return MeasuredPublicProofValAArtifactSchemaStateIncomplete
	}
	if len(model.RequiredDigestAlgorithms) == 0 || len(model.RequiredPackagingFiles) == 0 || len(model.RequiredSignatureFields) == 0 || len(model.RequiredMeasurementRefs) == 0 {
		return MeasuredPublicProofValAArtifactSchemaStatePartial
	}
	return MeasuredPublicProofValAArtifactSchemaStateActive
}

func EvaluateMeasuredPublicProofValASealingDisciplineState(model PublicProofSealingDiscipline) string {
	if strings.TrimSpace(model.SigningPurpose) == "" || strings.TrimSpace(model.Provider.ProviderMode) == "" || strings.TrimSpace(model.Provider.TrustBoundary) == "" {
		return MeasuredPublicProofValASealingDisciplineStateIncomplete
	}
	if len(model.RequiredTrustFields) == 0 || len(model.TimestampPolicy) == 0 || len(model.PackagingRules) == 0 || len(model.RevocationRules) == 0 || len(model.RequiredArtifactRefs) == 0 {
		return MeasuredPublicProofValASealingDisciplineStatePartial
	}
	if !model.PurposeEnabled || model.Provider.ProviderMode == signing.ModeDisabled || !containsTrimmedString(model.Provider.KeyClasses, signing.KeyClassSealing) {
		return MeasuredPublicProofValASealingDisciplineStatePartial
	}
	return MeasuredPublicProofValASealingDisciplineStateActive
}

func EvaluateMeasuredPublicProofValAEnvironmentBindingState(items []PublicProofEnvironmentBindingItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValAEnvironmentBindingStateIncomplete
	}
	for _, item := range items {
		if strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ClaimClass) == "" {
			return MeasuredPublicProofValAEnvironmentBindingStatePartial
		}
		if strings.TrimSpace(item.RedactionTier) == "" || strings.TrimSpace(item.EnvironmentClass) == "" || strings.TrimSpace(item.ExecutionProfile) == "" || strings.TrimSpace(item.WorkloadShape) == "" || strings.TrimSpace(item.BuildIdentity) == "" || strings.TrimSpace(item.HarnessVersion) == "" || strings.TrimSpace(item.MethodologyRef) == "" || strings.TrimSpace(item.CompatibilityScope) == "" {
			return MeasuredPublicProofValAEnvironmentBindingStatePartial
		}
		if len(item.ProvenanceInputs) == 0 || len(item.ReplayBoundaries) == 0 || len(item.UnsupportedReplay) == 0 {
			return MeasuredPublicProofValAEnvironmentBindingStatePartial
		}
	}
	return MeasuredPublicProofValAEnvironmentBindingStateActive
}

func EvaluateMeasuredPublicProofValADownloadablePackState(items []PublicSealedProofPack) string {
	if len(items) == 0 {
		return MeasuredPublicProofValADownloadablePackStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.ArtifactSchemaVersion) == "" || strings.TrimSpace(item.ArtifactType) == "" || strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ClaimClass) == "" {
			return MeasuredPublicProofValADownloadablePackStateIncomplete
		}
		if strings.TrimSpace(item.RedactionTier) == "" || strings.TrimSpace(item.EnvironmentClass) == "" || strings.TrimSpace(item.ExecutionProfile) == "" || strings.TrimSpace(item.WorkloadShape) == "" || strings.TrimSpace(item.BuildIdentity) == "" || strings.TrimSpace(item.HarnessVersion) == "" || strings.TrimSpace(item.MethodologyRef) == "" {
			return MeasuredPublicProofValADownloadablePackStatePartial
		}
		if item.IssuedAt.IsZero() || item.ValidThrough.IsZero() || strings.TrimSpace(item.MeasurementSource) == "" || len(item.EvidenceRefs) == 0 || strings.TrimSpace(item.DownloadRef) == "" || strings.TrimSpace(item.PayloadDigest) == "" || strings.TrimSpace(item.TrustRootID) == "" || strings.TrimSpace(item.KeyVersion) == "" || strings.TrimSpace(item.TimestampRef) == "" || len(item.PackagingFiles) == 0 || len(item.MetricSummaries) == 0 {
			return MeasuredPublicProofValADownloadablePackStatePartial
		}
		if item.SignatureEnvelope == nil || strings.TrimSpace(item.SignatureEnvelope.Provider) == "" || strings.TrimSpace(item.SignatureEnvelope.Purpose) != signing.PurposePublicProofArtifact || strings.TrimSpace(item.SignatureEnvelope.PayloadDigest) != strings.TrimSpace(item.PayloadDigest) || item.SignatureEnvelope.SignedAt.IsZero() {
			hasPartial = true
			continue
		}
		for _, file := range item.PackagingFiles {
			if strings.TrimSpace(file.Path) == "" || strings.TrimSpace(file.MediaType) == "" || strings.TrimSpace(file.Role) == "" || !strings.HasPrefix(strings.TrimSpace(file.SHA256), "sha256:") {
				return MeasuredPublicProofValADownloadablePackStatePartial
			}
		}
	}
	if hasPartial {
		return MeasuredPublicProofValADownloadablePackStatePartial
	}
	return MeasuredPublicProofValADownloadablePackStateActive
}

func EvaluateMeasuredPublicProofValAState(val0State, artifactSchemaState, sealingDisciplineState, environmentBindingState, downloadablePackState string) string {
	if strings.TrimSpace(val0State) != MeasuredPublicProofVal0StateActive {
		return MeasuredPublicProofValAStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(artifactSchemaState),
		strings.TrimSpace(sealingDisciplineState),
		strings.TrimSpace(environmentBindingState),
		strings.TrimSpace(downloadablePackState),
	} {
		switch state {
		case MeasuredPublicProofValAArtifactSchemaStateActive,
			MeasuredPublicProofValASealingDisciplineStateActive,
			MeasuredPublicProofValAEnvironmentBindingStateActive,
			MeasuredPublicProofValADownloadablePackStateActive:
		case MeasuredPublicProofValAArtifactSchemaStatePartial,
			MeasuredPublicProofValASealingDisciplineStatePartial,
			MeasuredPublicProofValAEnvironmentBindingStatePartial,
			MeasuredPublicProofValADownloadablePackStatePartial:
			hasPartial = true
		default:
			return MeasuredPublicProofValAStateIncomplete
		}
	}
	if hasPartial {
		return MeasuredPublicProofValAStateSubstantial
	}
	return MeasuredPublicProofValAStateActive
}

func containsTrimmedString(values []string, want string) bool {
	want = strings.TrimSpace(want)
	for _, value := range values {
		if strings.TrimSpace(value) == want {
			return true
		}
	}
	return false
}
