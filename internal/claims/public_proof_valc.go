package claims

import "strings"

const (
	MeasuredPublicProofValCPublicPortalStateActive     = "measured_public_proof_valc_public_portal_active"
	MeasuredPublicProofValCPublicPortalStatePartial    = "measured_public_proof_valc_public_portal_partial"
	MeasuredPublicProofValCPublicPortalStateIncomplete = "measured_public_proof_valc_public_portal_incomplete"

	MeasuredPublicProofValCPartnerPortalStateActive     = "measured_public_proof_valc_partner_portal_active"
	MeasuredPublicProofValCPartnerPortalStatePartial    = "measured_public_proof_valc_partner_portal_partial"
	MeasuredPublicProofValCPartnerPortalStateIncomplete = "measured_public_proof_valc_partner_portal_incomplete"

	MeasuredPublicProofValCClaimLineageStateActive     = "measured_public_proof_valc_claim_lineage_active"
	MeasuredPublicProofValCClaimLineageStatePartial    = "measured_public_proof_valc_claim_lineage_partial"
	MeasuredPublicProofValCClaimLineageStateIncomplete = "measured_public_proof_valc_claim_lineage_incomplete"

	MeasuredPublicProofValCDownloadProjectionStateActive     = "measured_public_proof_valc_download_projection_active"
	MeasuredPublicProofValCDownloadProjectionStatePartial    = "measured_public_proof_valc_download_projection_partial"
	MeasuredPublicProofValCDownloadProjectionStateIncomplete = "measured_public_proof_valc_download_projection_incomplete"

	MeasuredPublicProofValCStateIncomplete  = "measured_public_proof_valc_incomplete"
	MeasuredPublicProofValCStateSubstantial = "measured_public_proof_valc_substantially_ready"
	MeasuredPublicProofValCStateActive      = "measured_public_proof_valc_active"
)

type PublicProofPortalProjectionItem struct {
	ClaimID         string   `json:"claim_id"`
	ArtifactID      string   `json:"artifact_id"`
	CurrentState    string   `json:"current_state"`
	ClaimClass      string   `json:"claim_class"`
	Scope           string   `json:"scope"`
	VisibilityState string   `json:"visibility_state"`
	FreshnessState  string   `json:"freshness_state"`
	MethodologyRef  string   `json:"methodology_ref"`
	DownloadRef     string   `json:"download_ref"`
	VerificationRef string   `json:"verification_ref"`
	ReplayRef       string   `json:"replay_ref"`
	LineageRef      string   `json:"lineage_ref"`
	EvidenceRefs    []string `json:"evidence_refs,omitempty"`
	StatusNotes     []string `json:"status_notes,omitempty"`
	Limitations     []string `json:"limitations,omitempty"`
}

type PublicProofClaimLineageItem struct {
	ClaimID           string   `json:"claim_id"`
	ArtifactID        string   `json:"artifact_id"`
	CurrentState      string   `json:"current_state"`
	FreshnessState    string   `json:"freshness_state"`
	PublicationScope  string   `json:"publication_scope"`
	VisibilityState   string   `json:"visibility_state"`
	SupersessionState string   `json:"supersession_state"`
	SupersededBy      string   `json:"superseded_by,omitempty"`
	ArtifactRefs      []string `json:"artifact_refs,omitempty"`
	TransparencyRefs  []string `json:"transparency_refs,omitempty"`
	VerifierRefs      []string `json:"verifier_refs,omitempty"`
	MethodologyRefs   []string `json:"methodology_refs,omitempty"`
	EvidenceRefs      []string `json:"evidence_refs,omitempty"`
	Limitations       []string `json:"limitations,omitempty"`
}

type PublicProofDownloadProjectionItem struct {
	ArtifactID         string   `json:"artifact_id"`
	ClaimID            string   `json:"claim_id"`
	CurrentState       string   `json:"current_state"`
	RedactionTier      string   `json:"redaction_tier"`
	PublicationScope   string   `json:"publication_scope"`
	VisibilityState    string   `json:"visibility_state"`
	DownloadRef        string   `json:"download_ref"`
	TimestampRef       string   `json:"timestamp_ref"`
	PayloadDigest      string   `json:"payload_digest"`
	ReplayAvailability string   `json:"replay_availability"`
	AllowedScopes      []string `json:"allowed_scopes,omitempty"`
	EvidenceRefs       []string `json:"evidence_refs,omitempty"`
	Limitations        []string `json:"limitations,omitempty"`
}

func EvaluateMeasuredPublicProofValCPublicPortalState(items []PublicProofPortalProjectionItem) string {
	return evaluateMeasuredPublicProofValCPortalState(items, ScopePublic, VisibilityPublicSafe, MeasuredPublicProofValCPublicPortalStateIncomplete, MeasuredPublicProofValCPublicPortalStatePartial, MeasuredPublicProofValCPublicPortalStateActive)
}

func EvaluateMeasuredPublicProofValCPartnerPortalState(items []PublicProofPortalProjectionItem) string {
	return evaluateMeasuredPublicProofValCPortalState(items, ScopePartner, VisibilityPartnerSafe, MeasuredPublicProofValCPartnerPortalStateIncomplete, MeasuredPublicProofValCPartnerPortalStatePartial, MeasuredPublicProofValCPartnerPortalStateActive)
}

func evaluateMeasuredPublicProofValCPortalState(items []PublicProofPortalProjectionItem, expectedScope, expectedVisibility, incompleteState, partialState, activeState string) string {
	if len(items) == 0 {
		return incompleteState
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ClaimClass) == "" {
			return incompleteState
		}
		if strings.TrimSpace(item.Scope) == "" || strings.TrimSpace(item.VisibilityState) == "" || strings.TrimSpace(item.FreshnessState) == "" || strings.TrimSpace(item.MethodologyRef) == "" || strings.TrimSpace(item.DownloadRef) == "" || strings.TrimSpace(item.VerificationRef) == "" || strings.TrimSpace(item.ReplayRef) == "" || strings.TrimSpace(item.LineageRef) == "" || len(item.EvidenceRefs) == 0 {
			return partialState
		}
		if strings.TrimSpace(item.Scope) != expectedScope || strings.TrimSpace(item.VisibilityState) != expectedVisibility || strings.TrimSpace(item.CurrentState) != "portal_projection_ready" {
			hasPartial = true
		}
	}
	if hasPartial {
		return partialState
	}
	return activeState
}

func EvaluateMeasuredPublicProofValCClaimLineageState(items []PublicProofClaimLineageItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValCClaimLineageStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.FreshnessState) == "" {
			return MeasuredPublicProofValCClaimLineageStateIncomplete
		}
		if strings.TrimSpace(item.PublicationScope) == "" || strings.TrimSpace(item.VisibilityState) == "" || strings.TrimSpace(item.SupersessionState) == "" || len(item.ArtifactRefs) == 0 || len(item.TransparencyRefs) == 0 || len(item.VerifierRefs) == 0 || len(item.MethodologyRefs) == 0 || len(item.EvidenceRefs) == 0 {
			return MeasuredPublicProofValCClaimLineageStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "lineage_ready" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValCClaimLineageStatePartial
	}
	return MeasuredPublicProofValCClaimLineageStateActive
}

func EvaluateMeasuredPublicProofValCDownloadProjectionState(items []PublicProofDownloadProjectionItem) string {
	if len(items) == 0 {
		return MeasuredPublicProofValCDownloadProjectionStateIncomplete
	}
	hasPartial := false
	for _, item := range items {
		if strings.TrimSpace(item.ArtifactID) == "" || strings.TrimSpace(item.ClaimID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.RedactionTier) == "" {
			return MeasuredPublicProofValCDownloadProjectionStateIncomplete
		}
		if strings.TrimSpace(item.PublicationScope) == "" || strings.TrimSpace(item.VisibilityState) == "" || strings.TrimSpace(item.DownloadRef) == "" || strings.TrimSpace(item.TimestampRef) == "" || strings.TrimSpace(item.PayloadDigest) == "" || strings.TrimSpace(item.ReplayAvailability) == "" || len(item.AllowedScopes) == 0 || len(item.EvidenceRefs) == 0 {
			return MeasuredPublicProofValCDownloadProjectionStatePartial
		}
		if strings.TrimSpace(item.CurrentState) != "download_projection_ready" {
			hasPartial = true
		}
	}
	if hasPartial {
		return MeasuredPublicProofValCDownloadProjectionStatePartial
	}
	return MeasuredPublicProofValCDownloadProjectionStateActive
}

func EvaluateMeasuredPublicProofValCState(valBState, publicPortalState, partnerPortalState, claimLineageState, downloadProjectionState string) string {
	if strings.TrimSpace(valBState) != MeasuredPublicProofValBStateActive {
		return MeasuredPublicProofValCStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(publicPortalState),
		strings.TrimSpace(partnerPortalState),
		strings.TrimSpace(claimLineageState),
		strings.TrimSpace(downloadProjectionState),
	} {
		switch state {
		case MeasuredPublicProofValCPublicPortalStateActive,
			MeasuredPublicProofValCPartnerPortalStateActive,
			MeasuredPublicProofValCClaimLineageStateActive,
			MeasuredPublicProofValCDownloadProjectionStateActive:
		case MeasuredPublicProofValCPublicPortalStatePartial,
			MeasuredPublicProofValCPartnerPortalStatePartial,
			MeasuredPublicProofValCClaimLineageStatePartial,
			MeasuredPublicProofValCDownloadProjectionStatePartial:
			hasPartial = true
		default:
			return MeasuredPublicProofValCStateIncomplete
		}
	}
	if hasPartial {
		return MeasuredPublicProofValCStateSubstantial
	}
	return MeasuredPublicProofValCStateActive
}
