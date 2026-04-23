package main

import (
	"context"
	"net/http"
	"sort"
	"time"

	claimscore "github.com/denisgrosek/changelock/internal/claims"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/signing"
)

const (
	publicProofValEReplayCorrectnessReviewSchema = "point2.measured_public_proof.vale.replay_correctness_review.v1"
	publicProofValESigningTrustReviewSchema      = "point2.measured_public_proof.vale.signing_trust_review.v1"
	publicProofValETransparencyReviewSchema      = "point2.measured_public_proof.vale.transparency_review.v1"
	publicProofValERedactionReviewSchema         = "point2.measured_public_proof.vale.redaction_review.v1"
	publicProofValECompatibilityReviewSchema     = "point2.measured_public_proof.vale.compatibility_review.v1"
	publicProofValEIssuanceReviewSchema          = "point2.measured_public_proof.vale.issuance_review.v1"
	publicProofValEFailureStateReviewSchema      = "point2.measured_public_proof.vale.failure_state_review.v1"
	publicProofValEProofsSchema                  = "point2.measured_public_proof.vale.proofs.v1"
)

type publicProofValEReplayCorrectnessReviewResponse struct {
	SchemaVersion string                                              `json:"schema_version"`
	GeneratedAt   time.Time                                           `json:"generated_at"`
	CurrentState  string                                              `json:"current_state"`
	Items         []claimscore.PublicProofReplayCorrectnessReviewItem `json:"items,omitempty"`
	RouteRefs     []string                                            `json:"route_refs,omitempty"`
	Limitations   []string                                            `json:"limitations,omitempty"`
}

type publicProofValESigningTrustReviewResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Items         []claimscore.PublicProofSigningTrustReviewItem `json:"items,omitempty"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type publicProofValETransparencyReviewResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Items         []claimscore.PublicProofTransparencyReviewItem `json:"items,omitempty"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type publicProofValERedactionReviewResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Items         []claimscore.PublicProofRedactionReviewItem `json:"items,omitempty"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type publicProofValECompatibilityReviewResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Items         []claimscore.PublicProofCompatibilityReviewItem `json:"items,omitempty"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type publicProofValEIssuanceReviewResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Items         []claimscore.PublicProofIssuanceReviewItem `json:"items,omitempty"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type publicProofValEFailureStateReviewResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Items         []claimscore.PublicProofFailureStateReviewItem `json:"items,omitempty"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type publicProofValEProofsResponse struct {
	SchemaVersion                string    `json:"schema_version"`
	GeneratedAt                  time.Time `json:"generated_at"`
	CurrentState                 string    `json:"current_state"`
	Phase6State                  string    `json:"phase6_state"`
	Val0State                    string    `json:"val0_state"`
	ValAState                    string    `json:"val_a_state"`
	ValBState                    string    `json:"val_b_state"`
	ValCState                    string    `json:"val_c_state"`
	ValDState                    string    `json:"val_d_state"`
	ReplayCorrectnessReviewState string    `json:"replay_correctness_review_state"`
	SigningTrustReviewState      string    `json:"signing_trust_review_state"`
	TransparencyReviewState      string    `json:"transparency_review_state"`
	RedactionReviewState         string    `json:"redaction_review_state"`
	CompatibilityReviewState     string    `json:"compatibility_review_state"`
	IssuanceReviewState          string    `json:"issuance_review_state"`
	FailureStateReviewState      string    `json:"failure_state_review_state"`
	SurfaceRefs                  []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                 []string  `json:"evidence_refs,omitempty"`
	DeferredScope                []string  `json:"deferred_scope,omitempty"`
	Limitations                  []string  `json:"limitations,omitempty"`
	IntegrationSummary           []string  `json:"integration_summary,omitempty"`
}

func (s server) publicProofValEReplayCorrectnessReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValEReplayCorrectnessReview(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValESigningTrustReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValESigningTrustReview(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValETransparencyReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValETransparencyReview(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValERedactionReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValERedactionReview(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValECompatibilityReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValECompatibilityReview(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValEIssuanceReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValEIssuanceReview(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValEFailureStateReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValEFailureStateReview(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) publicProofValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	asOf, err := parsePhase6AsOf(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err := s.buildPublicProofValEProofs(r.Context(), asOf)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildPublicProofValEReplayCorrectnessReview(ctx context.Context, asOf time.Time) (publicProofValEReplayCorrectnessReviewResponse, error) {
	items, err := s.publicProofValEReplayCorrectnessReviewItems(ctx, asOf)
	if err != nil {
		return publicProofValEReplayCorrectnessReviewResponse{}, err
	}
	return publicProofValEReplayCorrectnessReviewResponse{
		SchemaVersion: publicProofValEReplayCorrectnessReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValEReplayCorrectnessReviewState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/vald/release-issuance-gate",
			"/v1/public/proof-expansion/vale/replay-correctness-review",
		},
		Limitations: []string{
			"Val E replay correctness review is a final bounded signoff over methodology-compatible and tolerance-bounded replay posture; it does not claim universal replay parity.",
		},
	}, nil
}

func (s server) buildPublicProofValESigningTrustReview(ctx context.Context, asOf time.Time) (publicProofValESigningTrustReviewResponse, error) {
	items, err := s.publicProofValESigningTrustReviewItems(ctx, asOf)
	if err != nil {
		return publicProofValESigningTrustReviewResponse{}, err
	}
	return publicProofValESigningTrustReviewResponse{
		SchemaVersion: publicProofValESigningTrustReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValESigningTrustReviewState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/signing-authority",
			"/v1/public/proof-expansion/vala/sealing-discipline",
			"/v1/public/proof-expansion/valb/signature-verification",
			"/v1/public/proof-expansion/vale/signing-trust-review",
		},
		Limitations: []string{
			"Val E signing and trust review remains bound to declared trust roots, signing purpose, timestamp linkage, and historical verification posture.",
		},
	}, nil
}

func (s server) buildPublicProofValETransparencyReview(ctx context.Context, asOf time.Time) (publicProofValETransparencyReviewResponse, error) {
	items, err := s.publicProofValETransparencyReviewItems(ctx, asOf)
	if err != nil {
		return publicProofValETransparencyReviewResponse{}, err
	}
	return publicProofValETransparencyReviewResponse{
		SchemaVersion: publicProofValETransparencyReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValETransparencyReviewState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/valb/transparency-chain",
			"/v1/public/proof-expansion/valc/claim-lineage",
			"/v1/public/proof-expansion/vald/claim-lifecycle",
			"/v1/public/proof-expansion/vale/transparency-review",
		},
		Limitations: []string{
			"Val E transparency review confirms digest-bound anchoring and visible lineage only; it does not replace later incident or publication operations with a new truth store.",
		},
	}, nil
}

func (s server) buildPublicProofValERedactionReview(ctx context.Context, asOf time.Time) (publicProofValERedactionReviewResponse, error) {
	items, err := s.publicProofValERedactionReviewItems(ctx, asOf)
	if err != nil {
		return publicProofValERedactionReviewResponse{}, err
	}
	return publicProofValERedactionReviewResponse{
		SchemaVersion: publicProofValERedactionReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValERedactionReviewState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/redaction-tiers",
			"/v1/public/proof-expansion/valc/public-proof-portal",
			"/v1/public/proof-expansion/valc/partner-proof-portal",
			"/v1/public/proof-expansion/valc/download-projections",
			"/v1/public/proof-expansion/vale/redaction-review",
		},
		Limitations: []string{
			"Val E redaction review signs off only on declared public-safe and partner-scoped projections; it does not authorize internal_full disclosure through partner or public surfaces.",
		},
	}, nil
}

func (s server) buildPublicProofValECompatibilityReview(ctx context.Context, asOf time.Time) (publicProofValECompatibilityReviewResponse, error) {
	items, err := s.publicProofValECompatibilityReviewItems(ctx, asOf)
	if err != nil {
		return publicProofValECompatibilityReviewResponse{}, err
	}
	return publicProofValECompatibilityReviewResponse{
		SchemaVersion: publicProofValECompatibilityReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValECompatibilityReviewState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/compatibility-baseline",
			"/v1/public/proof-expansion/valb/verifier-capability",
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/vale/compatibility-review",
		},
		Limitations: []string{
			"Val E compatibility review remains schema-bound, verifier-bound, methodology-bound, and replay-bound instead of implying universal interoperability.",
		},
	}, nil
}

func (s server) buildPublicProofValEIssuanceReview(ctx context.Context, asOf time.Time) (publicProofValEIssuanceReviewResponse, error) {
	items, err := s.publicProofValEIssuanceReviewItems(ctx, asOf)
	if err != nil {
		return publicProofValEIssuanceReviewResponse{}, err
	}
	return publicProofValEIssuanceReviewResponse{
		SchemaVersion: publicProofValEIssuanceReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValEIssuanceReviewState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/vald/release-issuance-gate",
			"/v1/public/proof-expansion/vald/claim-lifecycle",
			"/v1/public/proof-expansion/vald/publication-decisions",
			"/v1/public/proof-expansion/vald/correction-workflow",
			"/v1/public/proof-expansion/vale/issuance-review",
		},
		Limitations: []string{
			"Val E issuance review is the final signoff over existing release, lifecycle, and correction gates; it does not create a new mutable issuance database.",
		},
	}, nil
}

func (s server) buildPublicProofValEFailureStateReview(ctx context.Context, asOf time.Time) (publicProofValEFailureStateReviewResponse, error) {
	items, err := s.publicProofValEFailureStateReviewItems(ctx, asOf)
	if err != nil {
		return publicProofValEFailureStateReviewResponse{}, err
	}
	return publicProofValEFailureStateReviewResponse{
		SchemaVersion: publicProofValEFailureStateReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  claimscore.EvaluateMeasuredPublicProofValEFailureStateReviewState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/public/proof-expansion/val0/compatibility-baseline",
			"/v1/public/proof-expansion/vald/claim-lifecycle",
			"/v1/public/proof-expansion/vald/correction-workflow",
			"/v1/public/proof-expansion/vale/failure-state-review",
		},
		Limitations: []string{
			"Val E failure-state review keeps proof failures, claim_not_reissued, restriction, supersession, withdrawal, and stale posture verifier-visible instead of silently hiding degraded publication state.",
		},
	}, nil
}

func (s server) buildPublicProofValEProofs(ctx context.Context, asOf time.Time) (publicProofValEProofsResponse, error) {
	phase6, err := buildPhase6Proofs(asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	val0, err := s.buildPublicProofVal0Proofs(asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	valA, err := s.buildPublicProofValAProofs(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	valB, err := s.buildPublicProofValBProofs(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	valC, err := s.buildPublicProofValCProofs(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	valD, err := s.buildPublicProofValDProofs(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	replay, err := s.buildPublicProofValEReplayCorrectnessReview(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	signingTrust, err := s.buildPublicProofValESigningTrustReview(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	transparency, err := s.buildPublicProofValETransparencyReview(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	redaction, err := s.buildPublicProofValERedactionReview(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	compatibility, err := s.buildPublicProofValECompatibilityReview(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	issuance, err := s.buildPublicProofValEIssuanceReview(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	failureStates, err := s.buildPublicProofValEFailureStateReview(ctx, asOf)
	if err != nil {
		return publicProofValEProofsResponse{}, err
	}
	currentState := claimscore.EvaluateMeasuredPublicProofValEState(
		valD.CurrentState,
		replay.CurrentState,
		signingTrust.CurrentState,
		transparency.CurrentState,
		redaction.CurrentState,
		compatibility.CurrentState,
		issuance.CurrentState,
		failureStates.CurrentState,
	)
	return publicProofValEProofsResponse{
		SchemaVersion:                publicProofValEProofsSchema,
		GeneratedAt:                  publicSampleTime(),
		CurrentState:                 currentState,
		Phase6State:                  phase6.CurrentState,
		Val0State:                    val0.CurrentState,
		ValAState:                    valA.CurrentState,
		ValBState:                    valB.CurrentState,
		ValCState:                    valC.CurrentState,
		ValDState:                    valD.CurrentState,
		ReplayCorrectnessReviewState: replay.CurrentState,
		SigningTrustReviewState:      signingTrust.CurrentState,
		TransparencyReviewState:      transparency.CurrentState,
		RedactionReviewState:         redaction.CurrentState,
		CompatibilityReviewState:     compatibility.CurrentState,
		IssuanceReviewState:          issuance.CurrentState,
		FailureStateReviewState:      failureStates.CurrentState,
		SurfaceRefs: []string{
			"/v1/public/proof-expansion/vald/proofs",
			"/v1/public/proof-expansion/vale/replay-correctness-review",
			"/v1/public/proof-expansion/vale/signing-trust-review",
			"/v1/public/proof-expansion/vale/transparency-review",
			"/v1/public/proof-expansion/vale/redaction-review",
			"/v1/public/proof-expansion/vale/compatibility-review",
			"/v1/public/proof-expansion/vale/issuance-review",
			"/v1/public/proof-expansion/vale/failure-state-review",
			"/v1/public/proof-expansion/vale/proofs",
		},
		EvidenceRefs: []string{
			"/v1/public/proof-expansion/val0/proofs",
			"/v1/public/proof-expansion/vala/downloadable-packs",
			"/v1/public/proof-expansion/valb/transparency-chain",
			"/v1/public/proof-expansion/valb/signature-verification",
			"/v1/public/proof-expansion/valb/replay-verification",
			"/v1/public/proof-expansion/valc/claim-lineage",
			"/v1/public/proof-expansion/valc/download-projections",
			"/v1/public/proof-expansion/vald/release-issuance-gate",
			"/v1/public/proof-expansion/vald/claim-lifecycle",
			"/v1/public/proof-expansion/vald/publication-decisions",
			"/v1/public/proof-expansion/vald/correction-workflow",
		},
		DeferredScope: nil,
		Limitations: []string{
			"Val E closes Point 2 as a final proof-gate review over declared methodology, environment, trust-root, freshness, redaction, and governance boundaries.",
			"Final proof-gate approval remains bounded and does not convert signed, replayable, or anchored artifacts into universal or absolute truth claims.",
		},
		IntegrationSummary: []string{
			"Val E performs the final replay, signing, transparency, redaction, compatibility, issuance, and failure-state review over the existing Point 2 evidence spine.",
			"Point 2 closes only because final proof-gate signoff stays fail-closed on active Val D and keeps publication scope separate from lifecycle status.",
		},
	}, nil
}

type publicProofValEBaseData struct {
	ValDBase                publicProofValDBaseProjectionData
	SigningAuthority        claimscore.PublicProofSigningAuthorityModel
	Compatibility           claimscore.PublicProofCompatibilityBaseline
	VerifierCapability      claimscore.PublicProofVerifierCapability
	RedactionByTierID       map[string]claimscore.PublicProofRedactionTier
	ReleaseIssuanceByID     map[string]claimscore.PublicProofReleaseIssuanceItem
	ClaimLifecycleByID      map[string]claimscore.PublicProofClaimLifecycleItem
	PublicationDecisionByID map[string]claimscore.PublicProofPublicationDecisionItem
	CorrectionWorkflowByID  map[string]claimscore.PublicProofCorrectionWorkflowItem
}

func (s server) publicProofValEBaseData(ctx context.Context, asOf time.Time) (publicProofValEBaseData, error) {
	valDBase, err := s.publicProofValDBaseData(ctx, asOf)
	if err != nil {
		return publicProofValEBaseData{}, err
	}
	verifierCapability, err := s.buildPublicProofValBVerifierCapability(asOf)
	if err != nil {
		return publicProofValEBaseData{}, err
	}
	releaseIssuanceItems, err := s.publicProofValDReleaseIssuanceItems(ctx, asOf)
	if err != nil {
		return publicProofValEBaseData{}, err
	}
	claimLifecycleItems, err := s.publicProofValDClaimLifecycleItems(ctx, asOf)
	if err != nil {
		return publicProofValEBaseData{}, err
	}
	publicationItems, err := s.publicProofValDPublicationDecisionItems(ctx, asOf)
	if err != nil {
		return publicProofValEBaseData{}, err
	}
	correctionItems, err := s.publicProofValDCorrectionWorkflowItems(ctx, asOf)
	if err != nil {
		return publicProofValEBaseData{}, err
	}
	redactionByTierID := map[string]claimscore.PublicProofRedactionTier{}
	for _, item := range claimscore.MeasuredPublicProofVal0RedactionTiers() {
		redactionByTierID[item.TierID] = item
	}
	data := publicProofValEBaseData{
		ValDBase:                valDBase,
		SigningAuthority:        claimscore.MeasuredPublicProofVal0SigningAuthority(publicProofVal0ProviderDescriptor(s.signing)),
		Compatibility:           claimscore.MeasuredPublicProofVal0CompatibilityBaseline(),
		VerifierCapability:      verifierCapability.Model,
		RedactionByTierID:       redactionByTierID,
		ReleaseIssuanceByID:     map[string]claimscore.PublicProofReleaseIssuanceItem{},
		ClaimLifecycleByID:      map[string]claimscore.PublicProofClaimLifecycleItem{},
		PublicationDecisionByID: map[string]claimscore.PublicProofPublicationDecisionItem{},
		CorrectionWorkflowByID:  map[string]claimscore.PublicProofCorrectionWorkflowItem{},
	}
	for _, item := range releaseIssuanceItems {
		data.ReleaseIssuanceByID[item.ArtifactID] = item
	}
	for _, item := range claimLifecycleItems {
		data.ClaimLifecycleByID[item.ArtifactID] = item
	}
	for _, item := range publicationItems {
		data.PublicationDecisionByID[item.ArtifactID] = item
	}
	for _, item := range correctionItems {
		data.CorrectionWorkflowByID[item.ArtifactID] = item
	}
	return data, nil
}

func (s server) publicProofValEReplayCorrectnessReviewItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofReplayCorrectnessReviewItem, error) {
	data, err := s.publicProofValEBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofReplayCorrectnessReviewItem, 0, len(data.ValDBase.ValCBase.BuiltPacks))
	for _, built := range data.ValDBase.ValCBase.BuiltPacks {
		pack := built.Pack
		replayItem, hasReplay := data.ValDBase.ValCBase.ReplayByArtifactID[pack.ArtifactID]
		releaseItem, hasRelease := data.ReleaseIssuanceByID[pack.ArtifactID]
		currentState := "replay_review_ready"
		reviewOutcome := "approved"
		toleranceDecision := "within_declared_bands"
		if !hasReplay || !hasRelease || replayItem.CurrentState != "replay_ready" || (replayItem.EvaluationState != "comparison_verified" && replayItem.EvaluationState != "replay_verified") || len(replayItem.ToleranceBands) == 0 || len(replayItem.UnsupportedReplayCases) == 0 || len(replayItem.SupportedEnvironmentClasses) == 0 {
			currentState = "replay_review_limited"
			reviewOutcome = "approved_with_limitations"
			toleranceDecision = "tolerance_not_confirmed"
		}
		items = append(items, claimscore.PublicProofReplayCorrectnessReviewItem{
			ClaimID:                     pack.ClaimID,
			ArtifactID:                  pack.ArtifactID,
			CurrentState:                currentState,
			ReviewOutcome:               reviewOutcome,
			ReplayState:                 replayItem.EvaluationState,
			ComparisonMode:              replayItem.ComparisonMode,
			MethodologyRef:              firstNonEmpty(replayItem.MethodologyRef, pack.MethodologyRef),
			EvaluationRef:               replayItem.EvaluationRef,
			ToleranceDecision:           toleranceDecision,
			SupportedEnvironmentClasses: replayItem.SupportedEnvironmentClasses,
			ToleranceBands:              replayItem.ToleranceBands,
			UnsupportedReplayCases:      replayItem.UnsupportedReplayCases,
			ReviewRefs: []string{
				"/v1/public/proof-expansion/valb/replay-verification",
				"/v1/public/proof-expansion/vald/release-issuance-gate",
				"/v1/public/proof-expansion/vale/replay-correctness-review",
			},
			EvidenceRefs: uniqueStrings(append(append([]string{pack.DownloadRef, pack.TimestampRef}, replayItem.EvidenceRefs...), releaseItem.VerificationRefs...)),
			Limitations:  uniqueStrings(append(pack.Limitations, replayItem.Limitations...)),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValESigningTrustReviewItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofSigningTrustReviewItem, error) {
	data, err := s.publicProofValEBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	signingPurposeState := "purpose_disabled"
	if s.publicProofValASigningPurposeEnabled() {
		signingPurposeState = "purpose_enabled"
	}
	historicalVerificationState := "historical_verification_unavailable"
	if data.SigningAuthority.Provider.SupportsHistoricalVerification {
		historicalVerificationState = "historical_verification_ready"
	}
	keyRotationState := "rotation_limited"
	if data.SigningAuthority.Provider.SupportsRotation && data.SigningAuthority.Provider.SupportsVerifyOnlyRetirement {
		keyRotationState = "rotation_ready"
	}
	revocationState := "revocation_limited"
	if data.SigningAuthority.Provider.SupportsRevocation {
		revocationState = "revocation_ready"
	}
	items := make([]claimscore.PublicProofSigningTrustReviewItem, 0, len(data.ValDBase.ValCBase.BuiltPacks))
	for _, built := range data.ValDBase.ValCBase.BuiltPacks {
		pack := built.Pack
		signatureItem, hasSignature := data.ValDBase.ValCBase.SignatureByArtifactID[pack.ArtifactID]
		timestampState := "timestamp_unbound"
		if pack.TimestampRef != "" {
			timestampState = "timestamp_bound"
		}
		currentState := "signing_trust_review_ready"
		reviewOutcome := "approved"
		if !hasSignature ||
			signatureItem.VerificationState != signing.StateVerified ||
			signatureItem.TrustRootState != "trusted" ||
			signingPurposeState != "purpose_enabled" ||
			historicalVerificationState != "historical_verification_ready" ||
			keyRotationState != "rotation_ready" ||
			revocationState != "revocation_ready" ||
			timestampState != "timestamp_bound" ||
			data.SigningAuthority.Provider.ProviderMode == signing.ModeDisabled {
			currentState = "signing_trust_review_limited"
			reviewOutcome = "approved_with_limitations"
		}
		items = append(items, claimscore.PublicProofSigningTrustReviewItem{
			ClaimID:                     pack.ClaimID,
			ArtifactID:                  pack.ArtifactID,
			CurrentState:                currentState,
			ReviewOutcome:               reviewOutcome,
			VerificationState:           signatureItem.VerificationState,
			TrustRootState:              signatureItem.TrustRootState,
			SigningPurposeState:         signingPurposeState,
			HistoricalVerificationState: historicalVerificationState,
			KeyRotationState:            keyRotationState,
			RevocationState:             revocationState,
			TimestampState:              timestampState,
			SignerMode:                  data.SigningAuthority.Provider.ProviderMode,
			TrustRootID:                 pack.TrustRootID,
			KeyVersion:                  pack.KeyVersion,
			ReviewRefs: []string{
				"/v1/public/proof-expansion/val0/signing-authority",
				"/v1/public/proof-expansion/valb/signature-verification",
				"/v1/public/proof-expansion/vale/signing-trust-review",
			},
			EvidenceRefs: uniqueStrings(append(signatureItem.EvidenceRefs, pack.TimestampRef)),
			FailureStates: uniqueStrings(append(
				append([]string{}, signatureItem.FailureStates...),
				data.Compatibility.FailureStates...,
			)),
			Limitations: uniqueStrings(append(pack.Limitations, signatureItem.Limitations...)),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValETransparencyReviewItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofTransparencyReviewItem, error) {
	data, err := s.publicProofValEBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofTransparencyReviewItem, 0, len(data.ValDBase.ValCBase.BuiltPacks))
	for _, built := range data.ValDBase.ValCBase.BuiltPacks {
		pack := built.Pack
		entry, hasEntry := data.ValDBase.ValCBase.TransparencyByArtifactID[pack.ArtifactID]
		lineage, hasLineage := data.ValDBase.LineageByArtifactID[pack.ArtifactID]
		currentState := "transparency_review_ready"
		reviewOutcome := "approved"
		supersessionVisibility := "visibility_missing"
		if hasLineage {
			if lineage.SupersessionState == "not_superseded" {
				supersessionVisibility = "visible_not_superseded"
			} else {
				supersessionVisibility = "visible_" + lineage.SupersessionState
			}
		}
		if !hasEntry || !hasLineage || entry.CurrentState != "anchored_projection_ready" || lineage.CurrentState != "lineage_ready" || supersessionVisibility != "visible_not_superseded" {
			currentState = "transparency_review_limited"
			reviewOutcome = "approved_with_limitations"
		}
		items = append(items, claimscore.PublicProofTransparencyReviewItem{
			ClaimID:                pack.ClaimID,
			ArtifactID:             pack.ArtifactID,
			CurrentState:           currentState,
			ReviewOutcome:          reviewOutcome,
			TransparencyState:      entry.CurrentState,
			EntryHashState:         publicProofValEEntryHashState(entry),
			AnchorState:            publicProofValEAnchorState(entry),
			SupersessionVisibility: supersessionVisibility,
			AnchorRef:              entry.ParentAnchorRef,
			EntryID:                entry.EntryID,
			EntryHash:              entry.EntryHash,
			LineageRef:             "/v1/public/proof-expansion/valc/claim-lineage",
			ReviewRefs: []string{
				"/v1/public/proof-expansion/valb/transparency-chain",
				"/v1/public/proof-expansion/valc/claim-lineage",
				"/v1/public/proof-expansion/vale/transparency-review",
			},
			EvidenceRefs: uniqueStrings(append(entry.TransparencyRefs, lineage.EvidenceRefs...)),
			FailureStates: []string{
				"anchoring_unavailable",
				claimscore.PublicProofStatusSuperseded,
				claimscore.PublicProofStatusWithdrawn,
			},
			Limitations: uniqueStrings(append(pack.Limitations, entry.Limitations...)),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValERedactionReviewItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofRedactionReviewItem, error) {
	data, err := s.publicProofValEBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofRedactionReviewItem, 0, len(data.ValDBase.ValCBase.BuiltPacks))
	for _, built := range data.ValDBase.ValCBase.BuiltPacks {
		pack := built.Pack
		tier, hasTier := data.RedactionByTierID[pack.RedactionTier]
		portalItem, hasPortal := publicProofValDPortalProjection(data.ValDBase, pack)
		publicationItem, hasPublication := data.PublicationDecisionByID[pack.ArtifactID]
		publicationScope := publicProofValCPrimaryScope(pack.RedactionTier)
		if hasPublication && publicationItem.PublicationScope != "" {
			publicationScope = publicationItem.PublicationScope
		}
		currentState := "redaction_review_ready"
		reviewOutcome := "approved"
		redactionDecision := "tier_scope_mismatch"
		if hasTier && containsString(tier.AllowedScopes, publicationScope) {
			redactionDecision = "tier_scope_match"
		}
		projectionDiscipline := "projection_boundary_unclear"
		if hasTier && tier.PortalPolicy != "" {
			projectionDiscipline = "projection_only_enforced"
		}
		if !hasTier || !hasPortal || !hasPublication || portalItem.CurrentState != "portal_projection_ready" || redactionDecision != "tier_scope_match" || projectionDiscipline != "projection_only_enforced" {
			currentState = "redaction_review_limited"
			reviewOutcome = "approved_with_limitations"
		}
		items = append(items, claimscore.PublicProofRedactionReviewItem{
			ClaimID:               pack.ClaimID,
			ArtifactID:            pack.ArtifactID,
			CurrentState:          currentState,
			ReviewOutcome:         reviewOutcome,
			RedactionTier:         pack.RedactionTier,
			PublicationScope:      publicationScope,
			PortalProjectionState: portalItem.CurrentState,
			RedactionDecision:     redactionDecision,
			ProjectionDiscipline:  projectionDiscipline,
			RemovedFields:         tier.RemovedFields,
			NeverPublishedFields:  tier.NeverPublishedFields,
			ReviewRefs: []string{
				"/v1/public/proof-expansion/val0/redaction-tiers",
				publicProofValDPortalRef(pack),
				"/v1/public/proof-expansion/valc/download-projections",
				"/v1/public/proof-expansion/vale/redaction-review",
			},
			EvidenceRefs: uniqueStrings(append(portalItem.EvidenceRefs, pack.EvidenceRefs...)),
			FailureStates: []string{
				"claim_restricted",
				"claim_withdrawn",
				"claim_superseded",
			},
			Limitations: uniqueStrings(append(pack.Limitations, tier.Limitations...)),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValECompatibilityReviewItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofCompatibilityReviewItem, error) {
	data, err := s.publicProofValEBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	verifierState := claimscore.EvaluateMeasuredPublicProofValBVerifierCapabilityState(data.VerifierCapability)
	items := make([]claimscore.PublicProofCompatibilityReviewItem, 0, len(data.ValDBase.ValCBase.BuiltPacks))
	for _, built := range data.ValDBase.ValCBase.BuiltPacks {
		pack := built.Pack
		signatureItem, hasSignature := data.ValDBase.ValCBase.SignatureByArtifactID[pack.ArtifactID]
		replayItem, hasReplay := data.ValDBase.ValCBase.ReplayByArtifactID[pack.ArtifactID]
		verifierCompatibility := "unsupported"
		if verifierState == claimscore.MeasuredPublicProofValBVerifierCapabilityStateActive {
			verifierCompatibility = "supported"
		}
		replayCompatibility := "unsupported"
		if hasReplay && (replayItem.EvaluationState == "comparison_verified" || replayItem.EvaluationState == "replay_verified") {
			replayCompatibility = "supported"
		}
		methodologyCompatibility := "missing"
		if firstNonEmpty(replayItem.MethodologyRef, pack.MethodologyRef) != "" {
			methodologyCompatibility = "bounded_supported"
		}
		currentState := "compatibility_review_ready"
		reviewOutcome := "approved"
		if !hasSignature || !hasReplay || signatureItem.SchemaCompatibility != "supported" || verifierCompatibility != "supported" || replayCompatibility != "supported" || methodologyCompatibility != "bounded_supported" {
			currentState = "compatibility_review_limited"
			reviewOutcome = "approved_with_limitations"
		}
		items = append(items, claimscore.PublicProofCompatibilityReviewItem{
			ClaimID:                  pack.ClaimID,
			ArtifactID:               pack.ArtifactID,
			CurrentState:             currentState,
			ReviewOutcome:            reviewOutcome,
			SchemaCompatibility:      signatureItem.SchemaCompatibility,
			VerifierCompatibility:    verifierCompatibility,
			DeprecationState:         "not_deprecated",
			ReplayCompatibility:      replayCompatibility,
			MethodologyCompatibility: methodologyCompatibility,
			SupportedSchemaLines: uniqueStrings(append(
				append([]string{pack.ArtifactSchemaVersion}, data.Compatibility.SupportedArtifactSchemas...),
				data.VerifierCapability.SupportedSchemaLines...,
			)),
			UnsupportedCases: uniqueStrings(append([]string{}, data.Compatibility.UnsupportedCases...)),
			ReviewRefs: []string{
				"/v1/public/proof-expansion/val0/compatibility-baseline",
				"/v1/public/proof-expansion/valb/verifier-capability",
				"/v1/public/proof-expansion/valb/signature-verification",
				"/v1/public/proof-expansion/valb/replay-verification",
				"/v1/public/proof-expansion/vale/compatibility-review",
			},
			EvidenceRefs: uniqueStrings(append(append(signatureItem.EvidenceRefs, replayItem.EvidenceRefs...), pack.EvidenceRefs...)),
			FailureStates: uniqueStrings(append(
				append([]string{}, data.Compatibility.FailureStates...),
				signatureItem.FailureStates...,
			)),
			Limitations: uniqueStrings(append(append(pack.Limitations, replayItem.Limitations...), data.Compatibility.Limitations...)),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValEIssuanceReviewItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofIssuanceReviewItem, error) {
	data, err := s.publicProofValEBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	items := make([]claimscore.PublicProofIssuanceReviewItem, 0, len(data.ValDBase.ValCBase.BuiltPacks))
	for _, built := range data.ValDBase.ValCBase.BuiltPacks {
		pack := built.Pack
		releaseItem, hasRelease := data.ReleaseIssuanceByID[pack.ArtifactID]
		lifecycleItem, hasLifecycle := data.ClaimLifecycleByID[pack.ArtifactID]
		publicationItem, hasPublication := data.PublicationDecisionByID[pack.ArtifactID]
		correctionItem, hasCorrection := data.CorrectionWorkflowByID[pack.ArtifactID]
		overrideState := "override_permitted"
		if hasPublication && !publicationItem.OverridePermitted {
			overrideState = "override_not_permitted"
		}
		currentState := "issuance_review_ready"
		reviewOutcome := "approved"
		if !hasRelease || !hasLifecycle || !hasPublication || !hasCorrection ||
			releaseItem.CurrentState != "issuance_gate_ready" ||
			lifecycleItem.CurrentState != "claim_lifecycle_governed" ||
			publicationItem.CurrentState != "publication_decision_ready" ||
			correctionItem.CurrentState != "correction_workflow_ready" ||
			overrideState != "override_not_permitted" {
			currentState = "issuance_review_limited"
			reviewOutcome = "approved_with_limitations"
		}
		items = append(items, claimscore.PublicProofIssuanceReviewItem{
			ClaimID:                 pack.ClaimID,
			ArtifactID:              pack.ArtifactID,
			CurrentState:            currentState,
			ReviewOutcome:           reviewOutcome,
			ReleaseIssuanceState:    releaseItem.CurrentState,
			ClaimLifecycleStatus:    lifecycleItem.ClaimStatus,
			PublicationDecision:     publicationItem.PublicationStatus,
			CorrectionWorkflowState: correctionItem.CurrentState,
			OverrideState:           overrideState,
			ReviewRefs: []string{
				"/v1/public/proof-expansion/vald/release-issuance-gate",
				"/v1/public/proof-expansion/vald/claim-lifecycle",
				"/v1/public/proof-expansion/vald/publication-decisions",
				"/v1/public/proof-expansion/vald/correction-workflow",
				"/v1/public/proof-expansion/vale/issuance-review",
			},
			EvidenceRefs: uniqueStrings(append(
				append(append([]string{}, releaseItem.AuditRefs...), releaseItem.VerificationRefs...),
				append(publicationItem.DecisionAuditRefs, correctionItem.AuditRefs...)...,
			)),
			FailureStates: uniqueStrings(append(
				append(append([]string{}, releaseItem.FailureStates...), publicationItem.FailureStates...),
				correctionItem.FailureStates...,
			)),
			Limitations: uniqueStrings(append(append(pack.Limitations, releaseItem.Limitations...), correctionItem.Limitations...)),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func (s server) publicProofValEFailureStateReviewItems(ctx context.Context, asOf time.Time) ([]claimscore.PublicProofFailureStateReviewItem, error) {
	data, err := s.publicProofValEBaseData(ctx, asOf)
	if err != nil {
		return nil, err
	}
	supportedFailureStates := []string{
		claimscore.PublicProofStatusProofFailed,
		claimscore.PublicProofStatusClaimNotReissued,
		claimscore.PublicProofStatusRestricted,
		claimscore.PublicProofStatusSuperseded,
		claimscore.PublicProofStatusWithdrawn,
		claimscore.PublicProofStatusStale,
	}
	items := make([]claimscore.PublicProofFailureStateReviewItem, 0, len(data.ValDBase.ValCBase.BuiltPacks))
	for _, built := range data.ValDBase.ValCBase.BuiltPacks {
		pack := built.Pack
		releaseItem, hasRelease := data.ReleaseIssuanceByID[pack.ArtifactID]
		lifecycleItem, hasLifecycle := data.ClaimLifecycleByID[pack.ArtifactID]
		publicationItem, hasPublication := data.PublicationDecisionByID[pack.ArtifactID]
		correctionItem, hasCorrection := data.CorrectionWorkflowByID[pack.ArtifactID]
		currentState := "failure_state_review_ready"
		reviewOutcome := "approved"
		if !hasRelease || !hasLifecycle || !hasPublication || !hasCorrection || len(data.Compatibility.FailureStates) == 0 || lifecycleItem.CurrentState != "claim_lifecycle_governed" {
			currentState = "failure_state_review_limited"
			reviewOutcome = "approved_with_limitations"
		}
		items = append(items, claimscore.PublicProofFailureStateReviewItem{
			ClaimID:                     pack.ClaimID,
			ArtifactID:                  pack.ArtifactID,
			CurrentState:                currentState,
			ReviewOutcome:               reviewOutcome,
			FailureVisibilityState:      "verifier_visible_failures_modeled",
			RestrictionVisibilityState:  "restriction_visible",
			WithdrawalVisibilityState:   "withdrawal_visible_not_triggered",
			SupersessionVisibilityState: "supersession_visible_not_triggered",
			ReissueFailureState:         "claim_not_reissued_modeled",
			SupportedFailureStates:      supportedFailureStates,
			ReviewRefs: []string{
				"/v1/public/proof-expansion/vald/claim-lifecycle",
				"/v1/public/proof-expansion/vald/publication-decisions",
				"/v1/public/proof-expansion/vald/correction-workflow",
				"/v1/public/proof-expansion/vale/failure-state-review",
			},
			EvidenceRefs: uniqueStrings(append(
				append(lifecycleItem.VerifierNoticeRefs, publicationItem.DecisionAuditRefs...),
				correctionItem.AuditRefs...,
			)),
			FailureStates: uniqueStrings(append(
				append(append([]string{}, data.Compatibility.FailureStates...), releaseItem.FailureStates...),
				append(publicationItem.FailureStates, correctionItem.FailureStates...)...,
			)),
			Limitations: uniqueStrings(append(append(pack.Limitations, lifecycleItem.Limitations...), correctionItem.Limitations...)),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ArtifactID < items[j].ArtifactID })
	return items, nil
}

func publicProofValEEntryHashState(entry claimscore.PublicProofTransparencyEntry) string {
	if entry.EntryHash == "" {
		return "entry_hash_missing"
	}
	return "digest_bound"
}

func publicProofValEAnchorState(entry claimscore.PublicProofTransparencyEntry) string {
	if entry.ParentAnchorRef == "" {
		return "anchor_unavailable"
	}
	return "anchor_active"
}
