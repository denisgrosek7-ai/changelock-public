package main

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"time"

	claimscore "github.com/denisgrosek/changelock/internal/claims"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	publicProofValCPublicPortalSchema        = "point2.measured_public_proof.valc.public_portal.v1"
	publicProofValCPartnerPortalSchema       = "point2.measured_public_proof.valc.partner_portal.v1"
	publicProofValCClaimLineageSchema        = "point2.measured_public_proof.valc.claim_lineage.v1"
	publicProofValCDownloadProjectionsSchema = "point2.measured_public_proof.valc.download_projections.v1"
	publicProofValCProofsSchema              = "point2.measured_public_proof.valc.proofs.v1"
)

type publicProofValCPortalResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	Scope         string                                       `json:"scope"`
	CurrentState  string                                       `json:"current_state"`
	Items         []claimscore.PublicProofPortalProjectionItem `json:"items,omitempty"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type publicProofValCClaimLineageResponse struct {
	SchemaVersion string                                   `json:"schema_version"`
	GeneratedAt   time.Time                                `json:"generated_at"`
	CurrentState  string                                   `json:"current_state"`
	Items         []claimscore.PublicProofClaimLineageItem `json:"items,omitempty"`
	RouteRefs     []string                                 `json:"route_refs,omitempty"`
	Limitations   []string                                 `json:"limitations,omitempty"`
}

type publicProofValCDownloadProjectionsResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Items         []claimscore.PublicProofDownloadProjectionItem `json:"items,omitempty"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type publicProofValCProofsResponse struct {
	SchemaVersion           string    `json:"schema_version"`
	GeneratedAt             time.Time `json:"generated_at"`
	CurrentState            string    `json:"current_state"`
	Phase6State             string    `json:"phase6_state"`
	Val0State               string    `json:"val0_state"`
	ValAState               string    `json:"val_a_state"`
	ValBState               string    `json:"val_b_state"`
	PublicPortalState       string    `json:"public_portal_state"`
	PartnerPortalState      string    `json:"partner_portal_state"`
	ClaimLineageState       string    `json:"claim_lineage_state"`
	DownloadProjectionState string    `json:"download_projection_state"`
	SurfaceRefs             []string  `json:"surface_refs,omitempty"`
	EvidenceRefs            []string  `json:"evidence_refs,omitempty"`
	DeferredScope           []string  `json:"deferred_scope,omitempty"`
	Limitations             []string  `json:"limitations,omitempty"`
	IntegrationSummary      []string  `json:"integration_summary,omitempty"`
}

func (s server) publicProofValCPublicPortalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValCPortal(r.Context(), asOf, phase6PublicScopePublic)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValCPartnerPortalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValCPortal(r.Context(), asOf, phase6PublicScopePartner)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValCClaimLineageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValCClaimLineage(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValCDownloadProjectionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValCDownloadProjections(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValCProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValCProofs(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildPublicProofValCPortal(ctx context.Context, asOf time.Time, scope string) (publicProofValCPortalResponse, error) {
	items, err := s.publicProofValCPortalItems(ctx, asOf, scope)
	if err != nil {
		return publicProofValCPortalResponse{}, err
	}
	currentState := claimscore.EvaluateMeasuredPublicProofValCPublicPortalState(items)
	schemaVersion := publicProofValCPublicPortalSchema
	limitations := []string{
		"Val C public portal remains a public-safe projection over sealed artifacts, claims, and verifier state; it does not expose partner-only evidence or tenant-sensitive raw data.",
	}
	routeRefs := []string{
		"/v1/public/phase6/proof-portal?scope=public",
		"/v1/public/claims/summary?scope=public",
		"/v1/public/proof-expansion/vala/downloadable-packs",
		"/v1/public/proof-expansion/valb/proofs",
		"/v1/public/proof-expansion/valc/public-proof-portal",
	}
	if scope == phase6PublicScopePartner {
		currentState = claimscore.EvaluateMeasuredPublicProofValCPartnerPortalState(items)
		schemaVersion = publicProofValCPartnerPortalSchema
		limitations = []string{
			"Val C partner portal remains partner-scoped and bounded by explicit redaction tiers; it is not an internal_full evidence view.",
		}
		routeRefs = []string{
			"/v1/public/phase6/proof-portal?scope=partner",
			"/v1/public/claims/summary?scope=partner",
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/valb/proofs",
			"/v1/public/proof-expansion/valc/partner-proof-portal",
		}
	}
	return publicProofValCPortalResponse{
		SchemaVersion: schemaVersion,
		GeneratedAt:   publicSampleTime(),
		Scope:         scope,
		CurrentState:  currentState,
		Items:         items,
		RouteRefs:     routeRefs,
		Limitations:   limitations,
	}, nil
}

func (s server) buildPublicProofValCClaimLineage(ctx context.Context, asOf time.Time) (publicProofValCClaimLineageResponse, error) {
	items, err := s.publicProofValCClaimLineageItems(ctx, asOf)
	if err != nil {
		return publicProofValCClaimLineageResponse{}, err
	}
	return publicProofValCClaimLineageResponse{
		SchemaVersion: publicProofValCClaimLineageSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValCClaimLineageState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/valb/transparency-chain",
			"/v1/public/proof-expansion/valb/signature-verification",
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/valc/claim-lineage",
		},
		Limitations: []string{
			"Val C lineage remains a read-only linkage view over current claim and artifact state; supersession and withdrawal governance remain deferred to later Point 2 waves.",
		},
	}, nil
}

func (s server) buildPublicProofValCDownloadProjections(ctx context.Context, asOf time.Time) (publicProofValCDownloadProjectionsResponse, error) {
	items, err := s.publicProofValCDownloadProjectionItems(ctx, asOf)
	if err != nil {
		return publicProofValCDownloadProjectionsResponse{}, err
	}
	return publicProofValCDownloadProjectionsResponse{
		SchemaVersion: publicProofValCDownloadProjectionsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValCDownloadProjectionState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/valc/download-projections",
		},
		Limitations: []string{
			"Val C download projections remain catalog views over already sealed Val A artifacts and do not introduce automated issuance or revocation workflow.",
		},
	}, nil
}

func (s server) buildPublicProofValCProofs(ctx context.Context, asOf time.Time) (publicProofValCProofsResponse, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	val0, err := s.buildPublicProofVal0Proofs(asOf)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	valA, err := s.buildPublicProofValAProofs(ctx, asOf)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	valB, err := s.buildPublicProofValBProofs(ctx, asOf)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	publicPortal, err := s.buildPublicProofValCPortal(ctx, asOf, phase6PublicScopePublic)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	partnerPortal, err := s.buildPublicProofValCPortal(ctx, asOf, phase6PublicScopePartner)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	lineage, err := s.buildPublicProofValCClaimLineage(ctx, asOf)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	downloads, err := s.buildPublicProofValCDownloadProjections(ctx, asOf)
	if err != nil {
		return publicProofValCProofsResponse{}, err
	}
	currentState := claimscore.EvaluateMeasuredPublicProofValCState(
		valB.CurrentState,
		publicPortal.CurrentState,
		partnerPortal.CurrentState,
		lineage.CurrentState,
		downloads.CurrentState,
	)
	return publicProofValCProofsResponse{
		SchemaVersion:           publicProofValCProofsSchema,
		GeneratedAt:             publicSampleTime(),
		CurrentState:            currentState,
		Phase6State:             phase6.CurrentState,
		Val0State:               val0.CurrentState,
		ValAState:               valA.CurrentState,
		ValBState:               valB.CurrentState,
		PublicPortalState:       publicPortal.CurrentState,
		PartnerPortalState:      partnerPortal.CurrentState,
		ClaimLineageState:       lineage.CurrentState,
		DownloadProjectionState: downloads.CurrentState,
		SurfaceRefs: []string{
			"/v1/public/proof-expansion/valb/proofs",
			"/v1/public/proof-expansion/valc/public-proof-portal",
			"/v1/public/proof-expansion/valc/partner-proof-portal",
			"/v1/public/proof-expansion/valc/claim-lineage",
			"/v1/public/proof-expansion/valc/download-projections",
			"/v1/public/proof-expansion/valc/proofs",
		},
		EvidenceRefs: []string{
			"/v1/public/phase6/proof-portal?scope=public",
			"/v1/public/claims/summary?scope=public",
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/valb/signature-verification",
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/valb/transparency-chain",
		},
		DeferredScope: []string{
			"point2_vald_automated_issuance_and_revocation_gate",
			"point2_vale_final_proof_gate",
		},
		Limitations: []string{
			"Val C closes public-safe and partner-scoped portal projection only; automated issuance, restriction, supersession, and withdrawal workflow remain deferred.",
			"Portal surfaces remain viewers over the evidence spine and do not replace verifier result state, methodology boundary, or claim lifecycle authority.",
		},
		IntegrationSummary: []string{
			"Val C exposes public-safe and partner-scoped portal projections over sealed artifacts, claim status, freshness, and verifier posture.",
			"Val C adds lineage and download catalog views without creating a new truth base or widening disclosure beyond explicit redaction tiers.",
		},
	}, nil
}

func (s server) publicProofValCPortalItems(ctx context.Context, asOf time.Time, scope string) ([]claimscore.PublicProofPortalProjectionItem, error) {
	phase6Portal, err := buildPhase6ProofPortal(scope, asOf)
	if err != nil {
		return nil, err
	}
	phase6Claims, err := buildPhase6ClaimsSummary(scope, asOf)
	if err != nil {
		return nil, err
	}
	data, err := s.publicProofValCBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofPortalProjectionItem, 0, len(data.BuiltPacks))
	for _, built := range data.BuiltPacks {
		pack := built.Pack
		if scope == phase6PublicScopePublic && pack.RedactionTier != claimscore.RedactionTierPublicSafe {
			continue
		}
		decision := publicProofValCDecision(pack, scope, asOf)
		if !claimscore.AllowsScope(decision, scope) {
			continue
		}
		currentState := "portal_projection_ready"
		transparencyEntry, hasTransparency := data.TransparencyByArtifactID[pack.ArtifactID]
		signatureItem, hasSignature := data.SignatureByArtifactID[pack.ArtifactID]
		replayItem, hasReplay := data.ReplayByArtifactID[pack.ArtifactID]
		if !hasTransparency || !hasSignature || !hasReplay || signatureItem.VerificationState != "verified" || (replayItem.EvaluationState != "comparison_verified" && replayItem.EvaluationState != "replay_verified") || phase6Portal.CurrentState == phase6PortalStateIncomplete || phase6Claims.CurrentState == phase6ClaimsStateIncomplete || decision.CurrentState == claimscore.StateBlocked {
			currentState = "portal_projection_limited"
		}
		items = append(items, claimscore.PublicProofPortalProjectionItem{
			ClaimID:         pack.ClaimID,
			ArtifactID:      pack.ArtifactID,
			CurrentState:    currentState,
			ClaimClass:      pack.ClaimClass,
			Scope:           scope,
			VisibilityState: publicProofValCPortalVisibility(scope),
			FreshnessState:  decision.FreshnessState,
			MethodologyRef:  pack.MethodologyRef,
			DownloadRef:     pack.DownloadRef,
			VerificationRef: "/v1/public/proof-expansion/valb/signature-verification",
			ReplayRef:       "/v1/public/proof-expansion/valb/replay-verification",
			LineageRef:      "/v1/public/proof-expansion/valc/claim-lineage",
			EvidenceRefs:    uniqueStrings(append(append([]string{}, pack.EvidenceRefs...), transparencyEntry.TransparencyRefs...)),
			StatusNotes: uniqueStrings([]string{
				"phase6_portal=" + phase6Portal.CurrentState,
				"phase6_claims=" + phase6Claims.CurrentState,
				"claim_state=" + decision.CurrentState,
				"signature_verification=" + signatureItem.VerificationState,
				"replay_verification=" + replayItem.EvaluationState,
			}),
			Limitations: uniqueStrings(append(pack.Limitations, "Portal projection stays bounded to "+scope+" scope and declared redaction tier.")),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValCClaimLineageItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofClaimLineageItem, error) {
	data, err := s.publicProofValCBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofClaimLineageItem, 0, len(data.BuiltPacks))
	for _, built := range data.BuiltPacks {
		pack := built.Pack
		scope := publicProofValCPrimaryScope(pack.RedactionTier)
		visibility := publicProofValCPrimaryVisibility(pack.RedactionTier)
		decision := publicProofValCDecision(pack, scope, asOf)
		currentState := "lineage_ready"
		transparencyEntry, hasTransparency := data.TransparencyByArtifactID[pack.ArtifactID]
		signatureItem, hasSignature := data.SignatureByArtifactID[pack.ArtifactID]
		replayItem, hasReplay := data.ReplayByArtifactID[pack.ArtifactID]
		if !hasTransparency || !hasSignature || !hasReplay || signatureItem.VerificationState != "verified" {
			currentState = "lineage_limited"
		}
		items = append(items, claimscore.PublicProofClaimLineageItem{
			ClaimID:           pack.ClaimID,
			ArtifactID:        pack.ArtifactID,
			CurrentState:      currentState,
			FreshnessState:    decision.FreshnessState,
			PublicationScope:  scope,
			VisibilityState:   visibility,
			SupersessionState: "not_superseded",
			ArtifactRefs:      uniqueStrings([]string{pack.DownloadRef, pack.TimestampRef}),
			TransparencyRefs:  transparencyEntry.TransparencyRefs,
			VerifierRefs: uniqueStrings([]string{
				"/v1/public/proof-expansion/valb/signature-verification",
				"/v1/public/proof-expansion/valb/replay-verification",
				"/v1/public/proof-expansion/valb/proofs",
			}),
			MethodologyRefs: uniqueStrings([]string{pack.MethodologyRef, replayItem.MethodologyRef}),
			EvidenceRefs:    pack.EvidenceRefs,
			Limitations:     uniqueStrings(append(pack.Limitations, "Lineage view remains current-state-only until Point 2 / Val D adds supersession and withdrawal workflow.")),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValCDownloadProjectionItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofDownloadProjectionItem, error) {
	data, err := s.publicProofValCBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofDownloadProjectionItem, 0, len(data.BuiltPacks))
	for _, built := range data.BuiltPacks {
		pack := built.Pack
		signatureItem, hasSignature := data.SignatureByArtifactID[pack.ArtifactID]
		replayItem, hasReplay := data.ReplayByArtifactID[pack.ArtifactID]
		currentState := "download_projection_ready"
		replayAvailability := publicProofValCReplayAvailability(replayItem)
		if !hasSignature || !hasReplay || signatureItem.VerificationState != "verified" || replayAvailability == "" || pack.CurrentState != "sealed_artifact_ready" {
			currentState = "download_projection_limited"
		}
		items = append(items, claimscore.PublicProofDownloadProjectionItem{
			ArtifactID:         pack.ArtifactID,
			ClaimID:            pack.ClaimID,
			CurrentState:       currentState,
			RedactionTier:      pack.RedactionTier,
			PublicationScope:   publicProofValCPrimaryScope(pack.RedactionTier),
			VisibilityState:    publicProofValCPrimaryVisibility(pack.RedactionTier),
			DownloadRef:        pack.DownloadRef,
			TimestampRef:       pack.TimestampRef,
			PayloadDigest:      pack.PayloadDigest,
			ReplayAvailability: replayAvailability,
			AllowedScopes:      publicProofValCAllowedScopes(pack.RedactionTier),
			EvidenceRefs:       pack.EvidenceRefs,
			Limitations:        uniqueStrings(append(pack.Limitations, "Download projection remains bounded to declared redaction tier and replay posture.")),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

type publicProofValCBaseProjectionData struct {
	BuiltPacks               []publicProofValABuiltPack
	TransparencyByArtifactID map[string]claimscore.PublicProofTransparencyEntry
	SignatureByArtifactID    map[string]claimscore.PublicProofSignatureVerificationItem
	ReplayByArtifactID       map[string]claimscore.PublicProofReplayVerificationItem
}

func (s server) publicProofValCBaseData(ctx context.Context, asOf time.Time) (publicProofValCBaseProjectionData, error) {
	builtPacks, err := s.publicProofValABuiltPacks(ctx, asOf)
	if err != nil {
		return publicProofValCBaseProjectionData{}, err
	}
	transparency, err := s.publicProofValBTransparencyChainModel(ctx, asOf)
	if err != nil {
		return publicProofValCBaseProjectionData{}, err
	}
	signatureItems, err := s.publicProofValBSignatureVerificationItems(ctx, asOf)
	if err != nil {
		return publicProofValCBaseProjectionData{}, err
	}
	replayItems, _, err := s.publicProofValBReplayVerificationItems(ctx, asOf)
	if err != nil {
		return publicProofValCBaseProjectionData{}, err
	}
	data := publicProofValCBaseProjectionData{
		BuiltPacks:               builtPacks,
		TransparencyByArtifactID: map[string]claimscore.PublicProofTransparencyEntry{},
		SignatureByArtifactID:    map[string]claimscore.PublicProofSignatureVerificationItem{},
		ReplayByArtifactID:       map[string]claimscore.PublicProofReplayVerificationItem{},
	}
	for _, entry := range transparency.Entries {
		data.TransparencyByArtifactID[entry.ArtifactID] = entry
	}
	for _, item := range signatureItems {
		data.SignatureByArtifactID[item.ArtifactID] = item
	}
	for _, item := range replayItems {
		data.ReplayByArtifactID[item.ArtifactID] = item
	}
	return data, nil
}

func publicProofValCDecision(pack claimscore.PublicSealedProofPack, scope string, asOf time.Time) claimscore.Decision {
	return claimscore.Evaluate(claimscore.Input{
		ClaimID:                  pack.ClaimID,
		ClaimClass:               pack.ClaimClass,
		PublicationClass:         pack.ClaimClass,
		Scope:                    scope,
		VerifiedAt:               pack.IssuedAt,
		ValidUntil:               pack.ValidThrough,
		EvidenceRefs:             pack.EvidenceRefs,
		ProofRefs:                []string{pack.DownloadRef},
		VerifierRefs:             []string{"/v1/public/proof-expansion/valb/signature-verification", "/v1/public/proof-expansion/valb/replay-verification"},
		MethodologyRef:           pack.MethodologyRef,
		MethodologyRefs:          []string{pack.MethodologyRef},
		SupportsIndependentCheck: pack.ClaimClass == claimscore.PublicProofClaimClassVerification,
		PartnerVisibleOnly:       pack.RedactionTier == claimscore.RedactionTierPartnerScoped,
	}, asOf)
}

func publicProofValCPortalVisibility(scope string) string {
	if scope == phase6PublicScopePartner {
		return claimscore.VisibilityPartnerSafe
	}
	return claimscore.VisibilityPublicSafe
}

func publicProofValCPrimaryScope(redactionTier string) string {
	if redactionTier == claimscore.RedactionTierPartnerScoped {
		return claimscore.ScopePartner
	}
	return claimscore.ScopePublic
}

func publicProofValCPrimaryVisibility(redactionTier string) string {
	if redactionTier == claimscore.RedactionTierPartnerScoped {
		return claimscore.VisibilityPartnerSafe
	}
	return claimscore.VisibilityPublicSafe
}

func publicProofValCAllowedScopes(redactionTier string) []string {
	if redactionTier == claimscore.RedactionTierPartnerScoped {
		return []string{claimscore.ScopePartner, claimscore.ScopeAuditor, claimscore.ScopeInternal}
	}
	return []string{claimscore.ScopePublic, claimscore.ScopePartner, claimscore.ScopeAuditor, claimscore.ScopeInternal}
}

func publicProofValCReplayAvailability(item claimscore.PublicProofReplayVerificationItem) string {
	switch strings.TrimSpace(item.EvaluationState) {
	case "comparison_verified":
		return "bounded_replay_available"
	case "replay_verified":
		return "reference_replay_available"
	default:
		return ""
	}
}
