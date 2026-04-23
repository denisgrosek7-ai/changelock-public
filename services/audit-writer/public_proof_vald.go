package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	claimscore "github.com/denisgrosek/changelock/internal/claims"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	publicProofValDReleaseIssuanceSchema     = "point2.measured_public_proof.vald.release_issuance_gate.v1"
	publicProofValDClaimLifecycleSchema      = "point2.measured_public_proof.vald.claim_lifecycle.v1"
	publicProofValDPublicationDecisionSchema = "point2.measured_public_proof.vald.publication_decisions.v1"
	publicProofValDCorrectionWorkflowSchema  = "point2.measured_public_proof.vald.correction_workflow.v1"
	publicProofValDProofsSchema              = "point2.measured_public_proof.vald.proofs.v1"
)

type publicProofValDReleaseIssuanceResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Items         []claimscore.PublicProofReleaseIssuanceItem `json:"items,omitempty"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type publicProofValDClaimLifecycleResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Items         []claimscore.PublicProofClaimLifecycleItem `json:"items,omitempty"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type publicProofValDPublicationDecisionResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Items         []claimscore.PublicProofPublicationDecisionItem `json:"items,omitempty"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type publicProofValDCorrectionWorkflowResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Items         []claimscore.PublicProofCorrectionWorkflowItem `json:"items,omitempty"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type publicProofValDProofsResponse struct {
	SchemaVersion            string    `json:"schema_version"`
	GeneratedAt              time.Time `json:"generated_at"`
	CurrentState             string    `json:"current_state"`
	Phase6State              string    `json:"phase6_state"`
	Val0State                string    `json:"val0_state"`
	ValAState                string    `json:"val_a_state"`
	ValBState                string    `json:"val_b_state"`
	ValCState                string    `json:"val_c_state"`
	ReleaseIssuanceState     string    `json:"release_issuance_state"`
	ClaimLifecycleState      string    `json:"claim_lifecycle_state"`
	PublicationDecisionState string    `json:"publication_decision_state"`
	CorrectionWorkflowState  string    `json:"correction_workflow_state"`
	SurfaceRefs              []string  `json:"surface_refs,omitempty"`
	EvidenceRefs             []string  `json:"evidence_refs,omitempty"`
	DeferredScope            []string  `json:"deferred_scope,omitempty"`
	Limitations              []string  `json:"limitations,omitempty"`
	IntegrationSummary       []string  `json:"integration_summary,omitempty"`
}

func (s server) publicProofValDReleaseIssuanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValDReleaseIssuance(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValDClaimLifecycleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValDClaimLifecycle(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValDPublicationDecisionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValDPublicationDecisions(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValDCorrectionWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValDCorrectionWorkflow(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValDProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValDProofs(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildPublicProofValDReleaseIssuance(ctx context.Context, asOf time.Time) (publicProofValDReleaseIssuanceResponse, error) {
	items, err := s.publicProofValDReleaseIssuanceItems(ctx, asOf)
	if err != nil {
		return publicProofValDReleaseIssuanceResponse{}, err
	}
	return publicProofValDReleaseIssuanceResponse{
		SchemaVersion: publicProofValDReleaseIssuanceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValDReleaseIssuanceState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/valb/proofs",
			"/v1/public/proof-expansion/valc/claim-lineage",
			"/v1/public/proof-expansion/vald/release-issuance-gate",
		},
		Limitations: []string{
			"Val D release issuance remains a read-only gate projection over sealed artifacts, verification, and lineage outputs; it does not create a mutable issuance store.",
		},
	}, nil
}

func (s server) buildPublicProofValDClaimLifecycle(ctx context.Context, asOf time.Time) (publicProofValDClaimLifecycleResponse, error) {
	items, err := s.publicProofValDClaimLifecycleItems(ctx, asOf)
	if err != nil {
		return publicProofValDClaimLifecycleResponse{}, err
	}
	return publicProofValDClaimLifecycleResponse{
		SchemaVersion: publicProofValDClaimLifecycleSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValDClaimLifecycleState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/claim-registry-model",
			"/v1/public/proof-expansion/valc/claim-lineage",
			"/v1/public/proof-expansion/vald/claim-lifecycle",
		},
		Limitations: []string{
			"Val D claim lifecycle exposes restriction, supersession, withdrawal, and claim-not-reissued governance as verifier-visible states; it does not yet automate final proof publication approval.",
		},
	}, nil
}

func (s server) buildPublicProofValDPublicationDecisions(ctx context.Context, asOf time.Time) (publicProofValDPublicationDecisionResponse, error) {
	items, err := s.publicProofValDPublicationDecisionItems(ctx, asOf)
	if err != nil {
		return publicProofValDPublicationDecisionResponse{}, err
	}
	return publicProofValDPublicationDecisionResponse{
		SchemaVersion: publicProofValDPublicationDecisionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValDPublicationDecisionState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/redaction-tiers",
			"/v1/public/proof-expansion/valc/public-proof-portal",
			"/v1/public/proof-expansion/valc/partner-proof-portal",
			"/v1/public/proof-expansion/vald/publication-decisions",
		},
		Limitations: []string{
			"Val D publication decisions expose bounded approval and restriction outcomes only; they do not yet add a fully automated external publication service.",
		},
	}, nil
}

func (s server) buildPublicProofValDCorrectionWorkflow(ctx context.Context, asOf time.Time) (publicProofValDCorrectionWorkflowResponse, error) {
	items, err := s.publicProofValDCorrectionWorkflowItems(ctx, asOf)
	if err != nil {
		return publicProofValDCorrectionWorkflowResponse{}, err
	}
	return publicProofValDCorrectionWorkflowResponse{
		SchemaVersion: publicProofValDCorrectionWorkflowSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValDCorrectionWorkflowState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/valb/transparency-chain",
			"/v1/public/proof-expansion/valb/proofs",
			"/v1/public/proof-expansion/valc/claim-lineage",
			"/v1/public/proof-expansion/vald/correction-workflow",
		},
		Limitations: []string{
			"Val D correction workflow keeps withdrawal, restriction, and supersession actions visible as bounded workflow steps; final proof-gate signoff remains Val E work.",
		},
	}, nil
}

func (s server) buildPublicProofValDProofs(ctx context.Context, asOf time.Time) (publicProofValDProofsResponse, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	val0, err := s.buildPublicProofVal0Proofs(asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	valA, err := s.buildPublicProofValAProofs(ctx, asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	valB, err := s.buildPublicProofValBProofs(ctx, asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	valC, err := s.buildPublicProofValCProofs(ctx, asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	releaseIssuance, err := s.buildPublicProofValDReleaseIssuance(ctx, asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	claimLifecycle, err := s.buildPublicProofValDClaimLifecycle(ctx, asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	publicationDecisions, err := s.buildPublicProofValDPublicationDecisions(ctx, asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	correctionWorkflow, err := s.buildPublicProofValDCorrectionWorkflow(ctx, asOf)
	if err != nil {
		return publicProofValDProofsResponse{}, err
	}
	currentState := claimscore.EvaluateMeasuredPublicProofValDState(
		valC.CurrentState,
		releaseIssuance.CurrentState,
		claimLifecycle.CurrentState,
		publicationDecisions.CurrentState,
		correctionWorkflow.CurrentState,
	)
	return publicProofValDProofsResponse{
		SchemaVersion:            publicProofValDProofsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             currentState,
		Phase6State:              phase6.CurrentState,
		Val0State:                val0.CurrentState,
		ValAState:                valA.CurrentState,
		ValBState:                valB.CurrentState,
		ValCState:                valC.CurrentState,
		ReleaseIssuanceState:     releaseIssuance.CurrentState,
		ClaimLifecycleState:      claimLifecycle.CurrentState,
		PublicationDecisionState: publicationDecisions.CurrentState,
		CorrectionWorkflowState:  correctionWorkflow.CurrentState,
		SurfaceRefs: []string{
			"/v1/public/proof-expansion/valc/proofs",
			"/v1/public/proof-expansion/vald/release-issuance-gate",
			"/v1/public/proof-expansion/vald/claim-lifecycle",
			"/v1/public/proof-expansion/vald/publication-decisions",
			"/v1/public/proof-expansion/vald/correction-workflow",
			"/v1/public/proof-expansion/vald/proofs",
		},
		EvidenceRefs: []string{
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/valb/transparency-chain",
			"/v1/public/proof-expansion/valb/signature-verification",
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/valc/claim-lineage",
			"/v1/public/proof-expansion/valc/download-projections",
		},
		DeferredScope: []string{
			"point2_vale_final_proof_gate",
		},
		Limitations: []string{
			"Val D closes release-bound issuance, lifecycle, and correction governance over existing artifact and verification surfaces, but it does not yet perform final replay, signing, redaction, and compatibility signoff.",
			"Governance surfaces remain bounded projections over sealed artifacts, verification, and lineage outputs rather than a new mutable truth base.",
		},
		IntegrationSummary: []string{
			"Val D binds release issuance decisions to existing sealed artifacts, verifier outputs, redaction tiers, and claim lineage so new proof claims do not silently inherit prior public posture.",
			"Val D makes restriction, withdrawal, supersession, and correction workflow states visible before Point 2 / Val E adds the final proof gate review.",
		},
	}, nil
}

type publicProofValDBaseProjectionData struct {
	ValCBase                  publicProofValCBaseProjectionData
	PublicPortalByArtifactID  map[string]claimscore.PublicProofPortalProjectionItem
	PartnerPortalByArtifactID map[string]claimscore.PublicProofPortalProjectionItem
	LineageByArtifactID       map[string]claimscore.PublicProofClaimLineageItem
	DownloadByArtifactID      map[string]claimscore.PublicProofDownloadProjectionItem
}

func (s server) publicProofValDBaseData(ctx context.Context, asOf time.Time) (publicProofValDBaseProjectionData, error) {
	valCBase, err := s.publicProofValCBaseData(ctx, asOf)
	if err != nil {
		return publicProofValDBaseProjectionData{}, err
	}
	publicPortalItems, err := s.publicProofValCPortalItems(ctx, asOf, phase6PublicScopePublic)
	if err != nil {
		return publicProofValDBaseProjectionData{}, err
	}
	partnerPortalItems, err := s.publicProofValCPortalItems(ctx, asOf, phase6PublicScopePartner)
	if err != nil {
		return publicProofValDBaseProjectionData{}, err
	}
	lineageItems, err := s.publicProofValCClaimLineageItems(ctx, asOf)
	if err != nil {
		return publicProofValDBaseProjectionData{}, err
	}
	downloadItems, err := s.publicProofValCDownloadProjectionItems(ctx, asOf)
	if err != nil {
		return publicProofValDBaseProjectionData{}, err
	}
	data := publicProofValDBaseProjectionData{
		ValCBase:                  valCBase,
		PublicPortalByArtifactID:  map[string]claimscore.PublicProofPortalProjectionItem{},
		PartnerPortalByArtifactID: map[string]claimscore.PublicProofPortalProjectionItem{},
		LineageByArtifactID:       map[string]claimscore.PublicProofClaimLineageItem{},
		DownloadByArtifactID:      map[string]claimscore.PublicProofDownloadProjectionItem{},
	}
	for _, item := range publicPortalItems {
		data.PublicPortalByArtifactID[item.ArtifactID] = item
	}
	for _, item := range partnerPortalItems {
		data.PartnerPortalByArtifactID[item.ArtifactID] = item
	}
	for _, item := range lineageItems {
		data.LineageByArtifactID[item.ArtifactID] = item
	}
	for _, item := range downloadItems {
		data.DownloadByArtifactID[item.ArtifactID] = item
	}
	return data, nil
}

func (s server) publicProofValDReleaseIssuanceItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofReleaseIssuanceItem, error) {
	data, err := s.publicProofValDBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofReleaseIssuanceItem, 0, len(data.ValCBase.BuiltPacks))
	for _, built := range data.ValCBase.BuiltPacks {
		pack := built.Pack
		signatureItem, hasSignature := data.ValCBase.SignatureByArtifactID[pack.ArtifactID]
		replayItem, hasReplay := data.ValCBase.ReplayByArtifactID[pack.ArtifactID]
		lineageItem, hasLineage := data.LineageByArtifactID[pack.ArtifactID]
		portalItem, hasPortal := publicProofValDPortalProjection(data, pack)
		currentState := "issuance_gate_ready"
		if !hasSignature || !hasReplay || !hasLineage || !hasPortal || signatureItem.VerificationState != "verified" || (replayItem.EvaluationState != "comparison_verified" && replayItem.EvaluationState != "replay_verified") || lineageItem.CurrentState != "lineage_ready" || portalItem.CurrentState != "portal_projection_ready" {
			currentState = "issuance_gate_limited"
		}
		requiredChecks := publicProofValDRequiredChecks(pack)
		satisfiedChecks := publicProofValDSatisfiedChecks(pack, signatureItem, replayItem, lineageItem, portalItem)
		items = append(items, claimscore.PublicProofReleaseIssuanceItem{
			ClaimID:             pack.ClaimID,
			ArtifactID:          pack.ArtifactID,
			CurrentState:        currentState,
			ReleaseID:           publicProofValDReleaseID(pack, asOf),
			BuildIdentity:       pack.BuildIdentity,
			ReleaseChannel:      "public_proof_expansion",
			PriorReleaseRef:     "/v1/public/phase6/proofs?as_of=" + publicProofValDEncodedAsOf(asOf),
			ReissueDecision:     publicProofValDReissueDecision(pack),
			PublicationDecision: publicProofValDPublicationStatus(pack),
			RequiredChecks:      requiredChecks,
			SatisfiedChecks:     satisfiedChecks,
			VerificationRefs: uniqueStrings([]string{
				"/v1/public/proof-expansion/valb/signature-verification",
				"/v1/public/proof-expansion/valb/replay-verification",
				"/v1/public/proof-expansion/valb/proofs",
			}),
			AuditRefs: uniqueStrings([]string{
				"/v1/public/proof-expansion/valb/transparency-chain",
				"/v1/public/proof-expansion/valc/claim-lineage",
				portalItem.LineageRef,
			}),
			FailureStates: []string{
				claimscore.PublicProofStatusProofFailed,
				claimscore.PublicProofStatusClaimNotReissued,
				claimscore.PublicProofStatusRestricted,
				claimscore.PublicProofStatusWithdrawn,
			},
			Limitations: uniqueStrings(append(pack.Limitations, "Release issuance remains release-bound and cannot silently inherit a prior public claim without declared reissue posture.")),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValDClaimLifecycleItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofClaimLifecycleItem, error) {
	data, err := s.publicProofValDBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofClaimLifecycleItem, 0, len(data.ValCBase.BuiltPacks))
	for _, built := range data.ValCBase.BuiltPacks {
		pack := built.Pack
		scope := publicProofValCPrimaryScope(pack.RedactionTier)
		decision := publicProofValCDecision(pack, scope, asOf)
		lineageItem, hasLineage := data.LineageByArtifactID[pack.ArtifactID]
		_, hasPortal := publicProofValDPortalProjection(data, pack)
		currentState := "claim_lifecycle_governed"
		if !hasLineage || !hasPortal || decision.CurrentState == claimscore.StateBlocked {
			currentState = "claim_lifecycle_limited"
		}
		items = append(items, claimscore.PublicProofClaimLifecycleItem{
			ClaimID:           pack.ClaimID,
			ArtifactID:        pack.ArtifactID,
			CurrentState:      currentState,
			ClaimStatus:       publicProofValDClaimStatus(pack),
			ReissueState:      "reissued_for_current_release",
			FreshnessState:    decision.FreshnessState,
			PublicationScope:  scope,
			RestrictionState:  publicProofValDRestrictionState(pack),
			WithdrawalState:   "not_withdrawn",
			SupersessionState: "not_superseded",
			SupportedLifecycleStates: []string{
				claimscore.PublicProofStatusProven,
				claimscore.PublicProofStatusPartiallyProven,
				claimscore.PublicProofStatusProofPending,
				claimscore.PublicProofStatusProofFailed,
				claimscore.PublicProofStatusClaimNotReissued,
				claimscore.PublicProofStatusRestricted,
				claimscore.PublicProofStatusSuperseded,
				claimscore.PublicProofStatusWithdrawn,
				claimscore.PublicProofStatusStale,
			},
			VerifierNoticeRefs: []string{
				"/v1/public/proof-expansion/valb/signature-verification",
				"/v1/public/proof-expansion/valb/replay-verification",
				"/v1/public/proof-expansion/valb/proofs",
			},
			PortalRefs: uniqueStrings([]string{
				publicProofValDPortalRef(pack),
				"/v1/public/proof-expansion/valc/claim-lineage",
			}),
			EvidenceRefs: uniqueStrings(append(append([]string{}, pack.EvidenceRefs...), lineageItem.ArtifactRefs...)),
			Limitations:  uniqueStrings(append(pack.Limitations, "Lifecycle view remains verifier-visible and keeps restricted, superseded, withdrawn, and claim_not_reissued states explicit.")),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValDPublicationDecisionItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofPublicationDecisionItem, error) {
	data, err := s.publicProofValDBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofPublicationDecisionItem, 0, len(data.ValCBase.BuiltPacks))
	for _, built := range data.ValCBase.BuiltPacks {
		pack := built.Pack
		portalItem, hasPortal := publicProofValDPortalProjection(data, pack)
		lineageItem, hasLineage := data.LineageByArtifactID[pack.ArtifactID]
		downloadItem, hasDownload := data.DownloadByArtifactID[pack.ArtifactID]
		currentState := "publication_decision_ready"
		if !hasPortal || !hasLineage || !hasDownload || portalItem.CurrentState != "portal_projection_ready" || lineageItem.CurrentState != "lineage_ready" || downloadItem.CurrentState != "download_projection_ready" {
			currentState = "publication_decision_limited"
		}
		items = append(items, claimscore.PublicProofPublicationDecisionItem{
			ClaimID:           pack.ClaimID,
			ArtifactID:        pack.ArtifactID,
			CurrentState:      currentState,
			PublicationStatus: publicProofValDPublicationStatus(pack),
			ApprovalBoundary:  "signature_replay_and_redaction_review_required",
			RedactionTier:     pack.RedactionTier,
			PublicationScope:  publicProofValCPrimaryScope(pack.RedactionTier),
			AutomationState:   publicProofValDAutomationState(pack),
			OverridePermitted: false,
			DecisionAuditRefs: uniqueStrings([]string{
				"/v1/public/proof-expansion/valb/proofs",
				"/v1/public/proof-expansion/valc/claim-lineage",
				"/v1/public/proof-expansion/valc/download-projections",
			}),
			ProjectionRefs: uniqueStrings([]string{
				publicProofValDPortalRef(pack),
				downloadItem.DownloadRef,
			}),
			FailureStates: []string{
				claimscore.PublicProofStatusProofFailed,
				claimscore.PublicProofStatusClaimNotReissued,
				claimscore.PublicProofStatusRestricted,
				claimscore.PublicProofStatusWithdrawn,
			},
			Limitations: uniqueStrings(append(pack.Limitations, "Publication decisions stay bounded by redaction tier, verifier state, and declared portal scope instead of ad hoc overrides.")),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValDCorrectionWorkflowItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofCorrectionWorkflowItem, error) {
	data, err := s.publicProofValDBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofCorrectionWorkflowItem, 0, len(data.ValCBase.BuiltPacks))
	for _, built := range data.ValCBase.BuiltPacks {
		pack := built.Pack
		_, hasPortal := publicProofValDPortalProjection(data, pack)
		_, hasLineage := data.LineageByArtifactID[pack.ArtifactID]
		currentState := "correction_workflow_ready"
		if !hasPortal || !hasLineage {
			currentState = "correction_workflow_limited"
		}
		items = append(items, claimscore.PublicProofCorrectionWorkflowItem{
			ClaimID:               pack.ClaimID,
			ArtifactID:            pack.ArtifactID,
			CurrentState:          currentState,
			TriggerClass:          "release_regression_or_incident",
			TriggerState:          "monitoring_ready",
			RestrictionActionRef:  "/v1/public/proof-expansion/vald/publication-decisions",
			WithdrawalActionRef:   "/v1/public/proof-expansion/vald/claim-lifecycle",
			SupersessionActionRef: "/v1/public/proof-expansion/valc/claim-lineage",
			CorrectionNoticeRef:   publicProofValDPortalRef(pack),
			AuditRefs: uniqueStrings([]string{
				"/v1/public/proof-expansion/valb/transparency-chain",
				"/v1/public/proof-expansion/valb/proofs",
				"/v1/public/proof-expansion/valc/claim-lineage",
				publicProofValDPortalRef(pack),
			}),
			VerifierRefs: []string{
				"/v1/public/proof-expansion/valb/signature-verification",
				"/v1/public/proof-expansion/valb/replay-verification",
			},
			PortalNoticeRefs: uniqueStrings([]string{
				publicProofValDPortalRef(pack),
				"/v1/public/proof-expansion/valc/download-projections",
			}),
			FailureStates: []string{
				claimscore.PublicProofStatusProofFailed,
				claimscore.PublicProofStatusRestricted,
				claimscore.PublicProofStatusWithdrawn,
				claimscore.PublicProofStatusSuperseded,
			},
			Limitations: uniqueStrings(append(pack.Limitations, "Correction workflow keeps restriction, withdrawal, and supersession steps visible without mutating historical lineage.")),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func publicProofValDPortalProjection(data publicProofValDBaseProjectionData, pack claimscore.PublicSealedProofPack) (claimscore.PublicProofPortalProjectionItem, bool) {
	if pack.RedactionTier == claimscore.RedactionTierPublicSafe {
		item, ok := data.PublicPortalByArtifactID[pack.ArtifactID]
		return item, ok
	}
	item, ok := data.PartnerPortalByArtifactID[pack.ArtifactID]
	return item, ok
}

func publicProofValDPortalRef(pack claimscore.PublicSealedProofPack) string {
	if pack.RedactionTier == claimscore.RedactionTierPublicSafe {
		return "/v1/public/proof-expansion/valc/public-proof-portal"
	}
	return "/v1/public/proof-expansion/valc/partner-proof-portal"
}

func publicProofValDReleaseID(pack claimscore.PublicSealedProofPack, asOf time.Time) string {
	suffix := "runtime"
	if pack.ArtifactID == "point2_verification_public_pack" {
		suffix = "verification"
	}
	return fmt.Sprintf("point2-%s-%s", asOf.UTC().Format("20060102t150405z"), suffix)
}

func publicProofValDReissueDecision(pack claimscore.PublicSealedProofPack) string {
	if pack.RedactionTier == claimscore.RedactionTierPartnerScoped {
		return "restricted_reissue_ready"
	}
	return "reissued"
}

func publicProofValDPublicationStatus(pack claimscore.PublicSealedProofPack) string {
	if pack.RedactionTier == claimscore.RedactionTierPartnerScoped {
		return "restricted_partner_scoped_reissue"
	}
	return "approved_public_safe_reissue"
}

func publicProofValDAutomationState(pack claimscore.PublicSealedProofPack) string {
	if pack.RedactionTier == claimscore.RedactionTierPartnerScoped {
		return "auto_issue_restricted_no_override"
	}
	return "auto_issue_ready_no_override"
}

func publicProofValDClaimStatus(pack claimscore.PublicSealedProofPack) string {
	if pack.RedactionTier == claimscore.RedactionTierPartnerScoped {
		return claimscore.PublicProofStatusRestricted
	}
	return claimscore.PublicProofStatusProven
}

func publicProofValDRestrictionState(pack claimscore.PublicSealedProofPack) string {
	if pack.RedactionTier == claimscore.RedactionTierPartnerScoped {
		return "restricted_to_partner_scope"
	}
	return "unrestricted_public_safe"
}

func publicProofValDRequiredChecks(pack claimscore.PublicSealedProofPack) []string {
	if pack.ArtifactID == "point2_verification_public_pack" {
		return []string{
			"signature_verified",
			"replay_verified",
			"partner_redaction_review",
			"claim_lineage_bound",
		}
	}
	return []string{
		"signature_verified",
		"replay_verified",
		"transparency_verified",
		"public_portal_ready",
	}
}

func publicProofValDSatisfiedChecks(pack claimscore.PublicSealedProofPack, signatureItem claimscore.PublicProofSignatureVerificationItem, replayItem claimscore.PublicProofReplayVerificationItem, lineageItem claimscore.PublicProofClaimLineageItem, portalItem claimscore.PublicProofPortalProjectionItem) []string {
	checks := make([]string, 0, 4)
	if signatureItem.VerificationState == "verified" {
		checks = append(checks, "signature_verified")
	}
	if replayItem.EvaluationState == "comparison_verified" || replayItem.EvaluationState == "replay_verified" {
		checks = append(checks, "replay_verified")
	}
	if lineageItem.CurrentState == "lineage_ready" {
		checks = append(checks, "claim_lineage_bound")
		checks = append(checks, "transparency_verified")
	}
	if portalItem.CurrentState == "portal_projection_ready" {
		if pack.ArtifactID == "point2_verification_public_pack" {
			checks = append(checks, "partner_redaction_review")
		} else {
			checks = append(checks, "public_portal_ready")
		}
	}
	return uniqueStrings(checks)
}

func publicProofValDEncodedAsOf(asOf time.Time) string {
	return url.QueryEscape(strings.TrimSpace(asOf.UTC().Format(time.RFC3339)))
}
