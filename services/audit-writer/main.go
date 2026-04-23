package main

import (
	"context"
	"crypto/subtle"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/metrics"
	"github.com/denisgrosek/changelock/internal/signing"
)

type server struct {
	store                    audit.Store
	backend                  string
	allowedOrigins           map[string]struct{}
	requestTimeout           time.Duration
	authConfig               auth.Config
	vulnOps                  *vulnOpsRuntime
	syncRuntime              *syncRuntime
	signing                  *signingRuntime
	internalToken            string
	readbackGrantSecretValue string
}

type ingestResponse struct {
	Status     string    `json:"status"`
	ID         int64     `json:"id"`
	RequestID  string    `json:"request_id"`
	ReceivedAt time.Time `json:"received_at"`
}

type eventsResponse struct {
	Events []audit.StoredEvent `json:"events"`
}

type exceptionsResponse struct {
	Exceptions []audit.PolicyException `json:"exceptions"`
}

type exceptionResponse struct {
	Status    string                `json:"status"`
	Exception audit.PolicyException `json:"exception"`
}

type exceptionActionResponse struct {
	Status    string                `json:"status"`
	Exception audit.PolicyException `json:"exception"`
}

type authInfoResponse struct {
	Authenticated bool   `json:"authenticated"`
	AuthMode      string `json:"auth_mode"`
	Subject       string `json:"subject,omitempty"`
	Role          string `json:"role,omitempty"`
	TokenID       string `json:"token_id,omitempty"`
	IdentityType  string `json:"identity_type,omitempty"`
	Email         string `json:"email,omitempty"`
	TenantID      string `json:"tenant_id,omitempty"`
	GlobalScope   bool   `json:"global_scope,omitempty"`
}

func main() {
	migrateOnly := flag.Bool("migrate-only", false, "apply database migrations and exit")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	store, backend, err := newStoreFromEnv(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	authConfig, err := loadAuthConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	vulnOps, err := loadVulnOpsRuntimeFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	syncRuntime, err := loadSyncRuntimeFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	signingRuntime, err := loadSigningRuntimeFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	vexConfig, err := loadVEXConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	if syncRuntime != nil {
		syncRuntime.signing = signingRuntime
	}
	if imported, err := importVEXDirectory(ctx, store, vexConfig, "vex-importer"); err != nil {
		log.Fatal(err)
	} else if imported > 0 {
		log.Printf("audit-writer imported %d VEX statements from %s", imported, vexConfig.ImportDir)
	}

	if *migrateOnly {
		log.Printf("audit-writer migrations applied using %s backend", backend)
		return
	}
	if vulnOps != nil {
		vulnOps.start(context.Background(), store)
	}
	if syncRuntime != nil {
		syncRuntime.start(context.Background(), store)
	}

	addr := ":" + envOrDefault("PORT", "8094")
	log.Printf("audit-writer listening on %s using %s backend", addr, backend)
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           newHandlerWithDeps(store, backend, authConfig, vulnOps, syncRuntime, signingRuntime),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}

func newHandler(store audit.Store, backend string) http.Handler {
	authConfig, err := loadAuthConfigFromEnv()
	if err != nil {
		panic(err)
	}
	vulnOps, err := loadVulnOpsRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	syncRuntime, err := loadSyncRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	signingRuntime, err := loadSigningRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	return newHandlerWithDeps(store, backend, authConfig, vulnOps, syncRuntime, signingRuntime)
}

func newHandlerWithAuth(store audit.Store, backend string, authConfig auth.Config) http.Handler {
	vulnOps, err := loadVulnOpsRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	syncRuntime, err := loadSyncRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	signingRuntime, err := loadSigningRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	return newHandlerWithRuntimesAndSigning(store, backend, authConfig, vulnOps, syncRuntime, signingRuntime)
}

func newHandlerWithDeps(store audit.Store, backend string, authConfig auth.Config, vulnOps *vulnOpsRuntime, syncRuntime *syncRuntime, signingRuntime *signingRuntime) http.Handler {
	return newHandlerWithRuntimesAndSigning(store, backend, authConfig, vulnOps, syncRuntime, signingRuntime)
}

func newHandlerWithRuntimes(store audit.Store, backend string, authConfig auth.Config, vulnOps *vulnOpsRuntime, syncRuntime *syncRuntime) http.Handler {
	signingRuntime, err := loadSigningRuntimeFromEnv()
	if err != nil {
		panic(err)
	}
	return newHandlerWithRuntimesAndSigning(store, backend, authConfig, vulnOps, syncRuntime, signingRuntime)
}

func newHandlerWithRuntimesAndSigning(store audit.Store, backend string, authConfig auth.Config, vulnOps *vulnOpsRuntime, syncRuntime *syncRuntime, signingRuntime *signingRuntime) http.Handler {
	if syncRuntime != nil {
		syncRuntime.signing = signingRuntime
	}
	srv := server{
		store:                    store,
		backend:                  backend,
		allowedOrigins:           allowedOriginsFromEnv(),
		requestTimeout:           envDurationOrDefault("CHANGELOCK_REPORTS_TIMEOUT", 5*time.Second),
		authConfig:               authConfig,
		vulnOps:                  vulnOps,
		syncRuntime:              syncRuntime,
		signing:                  signingRuntime,
		internalToken:            strings.TrimSpace(os.Getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")),
		readbackGrantSecretValue: readbackGrantSecret(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", srv.healthHandler)
	mux.HandleFunc("/ready", srv.readyHandler)
	mux.HandleFunc("/v1/sync/status", srv.syncStatusHandler)
	mux.HandleFunc("/v1/sync/exceptions", srv.syncExceptionsHandler)
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/v1/ingest", srv.ingestHandler)
	mux.HandleFunc("/v1/auth/me", srv.authMeHandler)
	mux.HandleFunc("/v1/sbom/ingest", srv.sbomIngestHandler)
	mux.HandleFunc("/v1/sbom/images/", srv.sbomImageHandler)
	mux.HandleFunc("/v1/sbom/components/search", srv.sbomComponentsSearchHandler)
	mux.HandleFunc("/v1/exceptions", srv.exceptionsHandler)
	mux.HandleFunc("/v1/exceptions/request", srv.requestExceptionHandler)
	mux.HandleFunc("/v1/exceptions/", srv.exceptionByIDHandler)
	mux.HandleFunc("/v1/exceptions/validate", srv.validateExceptionHandler)
	mux.HandleFunc("/v1/analytics/trends", srv.trendsHandler)
	mux.HandleFunc("/v1/analytics/delta", srv.analyticsDeltaHandler)
	mux.HandleFunc("/v1/analytics/anomalies", srv.analyticsAnomaliesHandler)
	mux.HandleFunc("/v1/analytics/scorecards", srv.analyticsScorecardsHandler)
	mux.HandleFunc("/v1/analytics/segments", srv.analyticsSegmentsHandler)
	mux.HandleFunc("/v1/analytics/top-violators", srv.topViolatorsHandler)
	mux.HandleFunc("/v1/analytics/drift-stats", srv.driftStatsHandler)
	mux.HandleFunc("/v1/topology/services", srv.topologyServicesHandler)
	mux.HandleFunc("/v1/topology/graph", srv.topologyGraphHandler)
	mux.HandleFunc("/v1/topology/blast-radius", srv.topologyBlastRadiusHandler)
	mux.HandleFunc("/v1/topology/delta", srv.topologyDeltaHandler)
	mux.HandleFunc("/v1/topology/heatmap", srv.topologyHeatmapHandler)
	mux.HandleFunc("/v1/topology/quarantine-simulation", srv.topologyQuarantineSimulationHandler)
	mux.HandleFunc("/v1/forensics/state", srv.forensicsStateHandler)
	mux.HandleFunc("/v1/forensics/delta", srv.forensicsDeltaHandler)
	mux.HandleFunc("/v1/forensics/timeline", srv.forensicsTimelineHandler)
	mux.HandleFunc("/v1/forensics/vex-flashback", srv.forensicsVEXFlashbackHandler)
	mux.HandleFunc("/v1/forensics/replay", srv.forensicsReplayHandler)
	mux.HandleFunc("/v1/handoff/seal", srv.handoffSealHandler)
	mux.HandleFunc("/v1/handoff/verify", srv.handoffVerifyHandler)
	mux.HandleFunc("/v1/handoff/quality-gates", srv.handoffQualityGatesHandler)
	mux.HandleFunc("/v1/handoff/", srv.handoffByIDHandler)
	mux.HandleFunc("/v1/federation/peers", srv.federationPeersHandler)
	mux.HandleFunc("/v1/federation/peers/", srv.federationPeerByIDHandler)
	mux.HandleFunc("/v1/federation/proof-request", srv.federationProofRequestHandler)
	mux.HandleFunc("/v1/federation/proof-verify", srv.federationProofVerifyHandler)
	mux.HandleFunc("/v1/federation/proof-history", srv.federationProofHistoryHandler)
	mux.HandleFunc("/v1/federation/policy-state", srv.federationPolicyStateHandler)
	mux.HandleFunc("/v1/federation/policy-sync", srv.federationPolicyStateHandler)
	mux.HandleFunc("/v1/federation/anchors", srv.federationAnchorsHandler)
	mux.HandleFunc("/v1/federation/global-view", srv.federationGlobalViewHandler)
	mux.HandleFunc("/v1/federation/resilience", srv.federationResilienceHandler)
	mux.HandleFunc("/v1/validation/scenarios", srv.validationScenariosHandler)
	mux.HandleFunc("/v1/validation/execute", srv.validationExecuteHandler)
	mux.HandleFunc("/v1/validation/executions", srv.validationExecutionsHandler)
	mux.HandleFunc("/v1/validation/executions/", srv.validationExecutionByIDHandler)
	mux.HandleFunc("/v1/validation/verdicts/", srv.validationVerdictByIDHandler)
	mux.HandleFunc("/v1/validation/certificates/", srv.validationCertificateByIDHandler)
	mux.HandleFunc("/v1/validation/regression/run", srv.validationRegressionRunHandler)
	mux.HandleFunc("/v1/validation/chaos/run", srv.validationChaosRunHandler)
	mux.HandleFunc("/v1/validation/compatibility/run", srv.validationCompatibilityRunHandler)
	mux.HandleFunc("/v1/validation/readiness", srv.validationReadinessHandler)
	mux.HandleFunc("/v1/validation/harness/scenarios", srv.validationHarnessScenariosHandler)
	mux.HandleFunc("/v1/validation/harness/score", srv.validationHarnessScoreHandler)
	mux.HandleFunc("/v1/validation/harness/runs", srv.validationHarnessRunsHandler)
	mux.HandleFunc("/v1/validation/harness/runs/", srv.validationHarnessRunByIDHandler)
	mux.HandleFunc("/v1/validation/harness/what-if", srv.validationHarnessWhatIfHandler)
	mux.HandleFunc("/v1/vulnerabilities/active", srv.activeVulnerabilitiesHandler)
	mux.HandleFunc("/v1/vulnerabilities/net", srv.vulnerabilityNetHandler)
	mux.HandleFunc("/v1/vulnerabilities/blast-radius", srv.vulnerabilityBlastRadiusHandler)
	mux.HandleFunc("/v1/vulnerabilities/timeline", srv.vulnerabilityTimelineHandler)
	mux.HandleFunc("/v1/vulnerabilities/rescan", srv.vulnerabilityRescanHandler)
	mux.HandleFunc("/v1/vulnerabilities/decisions", srv.vulnerabilityDecisionsHandler)
	mux.HandleFunc("/v1/vulnerabilities/decisions/", srv.vulnerabilityDecisionByIDHandler)
	mux.HandleFunc("/v1/ai/defense-gap-assessments", srv.defenseGapAssessmentsHandler)
	mux.HandleFunc("/v1/ai/policy-replay", srv.policyReplayAssessmentsHandler)
	mux.HandleFunc("/v1/ai/systemic-weaknesses", srv.systemicWeaknessesHandler)
	mux.HandleFunc("/v1/ai/executive-defense-report", srv.executiveDefenseReportHandler)
	mux.HandleFunc("/v1/readback/grants", srv.readbackGrantHandler)
	mux.HandleFunc("/v1/readback/defense-gap/", srv.readbackDefenseGapHandler)
	mux.HandleFunc("/v1/readback/policy-replay/", srv.readbackPolicyReplayHandler)
	mux.HandleFunc("/v1/readback/systemic-weakness/", srv.readbackSystemicWeaknessHandler)
	mux.HandleFunc("/v1/recommendation-actions", srv.recommendationActionsHandler)
	mux.HandleFunc("/v1/recommendation-actions/", srv.recommendationActionsHandler)
	mux.HandleFunc("/v1/recommendations", srv.recommendationsHandler)
	mux.HandleFunc("/v1/recommendations/", srv.recommendationByIDHandler)
	mux.HandleFunc("/v1/foundation/execution", srv.executionFoundationHandler)
	mux.HandleFunc("/v1/foundation/execution/contracts", srv.executionFoundationContractsHandler)
	mux.HandleFunc("/v1/foundation/execution/benchmarks", srv.executionFoundationBenchmarksHandler)
	mux.HandleFunc("/v1/foundation/execution/benchmarks/harness", srv.executionFoundationBenchmarkHarnessHandler)
	mux.HandleFunc("/v1/foundation/execution/benchmarks/evaluate", srv.executionFoundationBenchmarkEvaluateHandler)
	mux.HandleFunc("/v1/foundation/execution/async", srv.executionFoundationAsyncHandler)
	mux.HandleFunc("/v1/foundation/execution/async/tasks", srv.executionFoundationAsyncTasksHandler)
	mux.HandleFunc("/v1/foundation/execution/async/tasks/", srv.executionFoundationAsyncTaskByIDHandler)
	mux.HandleFunc("/v1/foundation/execution/traces", srv.executionFoundationTracesHandler)
	mux.HandleFunc("/v1/foundation/execution/trust", srv.executionFoundationTrustHandler)
	mux.HandleFunc("/v1/foundation/execution/trust/rotation-drill", srv.executionFoundationTrustRotationDrillHandler)
	mux.HandleFunc("/v1/foundation/execution/trust/rotation-drills", srv.executionFoundationTrustRotationDrillsHandler)
	mux.HandleFunc("/v1/foundation/execution/proofs", srv.executionFoundationProofsHandler)
	mux.HandleFunc("/v1/integrations/identity", srv.integrationIdentityHandler)
	mux.HandleFunc("/v1/integrations/tickets/catalog", srv.integrationTicketCatalogHandler)
	mux.HandleFunc("/v1/integrations/tickets/prepare", srv.integrationTicketPrepareHandler)
	mux.HandleFunc("/v1/integrations/siem/export", srv.integrationSIEMExportHandler)
	mux.HandleFunc("/v1/integrations/evidence/export", srv.integrationEvidenceExportHandler)
	mux.HandleFunc("/v1/command-center/timeline", srv.securityTimelineHandler)
	mux.HandleFunc("/v1/command-center/search", srv.commandCenterSearchHandler)
	mux.HandleFunc("/v1/command-center/notifications", srv.commandCenterNotificationsHandler)
	mux.HandleFunc("/v1/scorecard/metrics/", srv.scorecardMetricIncidentsHandler)
	mux.HandleFunc("/v1/incidents/package", srv.incidentPackageHandler)
	mux.HandleFunc("/v1/incidents", srv.incidentsHandler)
	mux.HandleFunc("/v1/incidents/", srv.incidentByIDHandler)
	mux.HandleFunc("/r/defense-gap/", srv.readbackDefenseGapHandler)
	mux.HandleFunc("/r/policy-replay/", srv.readbackPolicyReplayHandler)
	mux.HandleFunc("/r/systemic-weakness/", srv.readbackSystemicWeaknessHandler)
	mux.HandleFunc("/s/", srv.readbackShareHandler)
	mux.HandleFunc("/v1/vex/status", srv.vexStatusHandler)
	mux.HandleFunc("/v1/vex/ingest", srv.vexIngestHandler)
	mux.HandleFunc("/v1/vex", srv.vexStatementsHandler)
	mux.HandleFunc("/v1/vex/", srv.vexStatementByIDHandler)
	mux.HandleFunc("/v1/signing-identities/status", srv.signingIdentityStatusHandler)
	mux.HandleFunc("/v1/signing-identities/findings", srv.signingIdentityFindingsHandler)
	mux.HandleFunc("/v1/signing-identities/evaluate", srv.signingIdentityEvaluateHandler)
	mux.HandleFunc("/v1/signing-identities/policies", srv.signingIdentityPoliciesHandler)
	mux.HandleFunc("/v1/signing-identities/policies/", srv.signingIdentityPolicyByIDHandler)
	mux.HandleFunc("/v1/signing-identities/", srv.signingIdentityObservationByIDHandler)
	mux.HandleFunc("/v1/signing-identities", srv.signingIdentityObservationsHandler)
	mux.HandleFunc("/v1/scorecards/findings", srv.scorecardFindingsHandler)
	mux.HandleFunc("/v1/scorecards", srv.scorecardHandler)
	mux.HandleFunc("/v1/trust-badges", srv.trustBadgesHandler)
	mux.HandleFunc("/v1/trust/published", srv.publishedTrustViewHandler)
	mux.HandleFunc("/v1/public/specs/handoff", srv.publicHandoffSpecHandler)
	mux.HandleFunc("/v1/public/specs/proof-verification", srv.publicProofVerificationSpecHandler)
	mux.HandleFunc("/v1/public/specs/validation-certificate", srv.publicValidationCertificateSpecHandler)
	mux.HandleFunc("/v1/public/specs/federation-proof-exchange", srv.publicFederationExchangeSpecHandler)
	mux.HandleFunc("/v1/public/specs/explainability-boundaries", srv.publicExplainabilityBoundariesHandler)
	mux.HandleFunc("/v1/public/schemas", srv.publicSchemaIndexHandler)
	mux.HandleFunc("/v1/public/schemas/", srv.publicSchemaExportHandler)
	mux.HandleFunc("/v1/public/verifier/profiles", srv.publicVerifierProfilesHandler)
	mux.HandleFunc("/v1/public/verifier/offline-guide", srv.publicOfflineGuideHandler)
	mux.HandleFunc("/v1/public/verifier/reference-pack", srv.publicVerifierReferencePackHandler)
	mux.HandleFunc("/v1/public/samples/handoff", srv.publicHandoffSampleHandler)
	mux.HandleFunc("/v1/public/samples/proof-verification", srv.publicProofVerificationSampleHandler)
	mux.HandleFunc("/v1/public/samples/validation-certificate", srv.publicValidationCertificateSampleHandler)
	mux.HandleFunc("/v1/public/samples/federation-proof-exchange", srv.publicFederationExchangeSampleHandler)
	mux.HandleFunc("/v1/public/conformance-pack", srv.publicConformancePackHandler)
	mux.HandleFunc("/v1/public/reference-architectures", srv.publicReferenceArchitecturesHandler)
	mux.HandleFunc("/v1/public/reference-architectures/sector-profiles", srv.publicSectorProfilesHandler)
	mux.HandleFunc("/v1/public/maturity-map", srv.publicMaturityMapHandler)
	mux.HandleFunc("/v1/public/decision-guides", srv.publicDecisionGuidesHandler)
	mux.HandleFunc("/v1/public/decision-guides/matrix", srv.publicDeploymentDecisionMatrixHandler)
	mux.HandleFunc("/v1/public/benchmarks/methodology", srv.publicBenchmarkMethodologyHandler)
	mux.HandleFunc("/v1/public/benchmarks/set", srv.publicBenchmarkSetHandler)
	mux.HandleFunc("/v1/public/analytics/publication-discipline", srv.publicAnalyticsPublicationDisciplineHandler)
	mux.HandleFunc("/v1/public/case-studies", srv.publicCaseStudyPacksHandler)
	mux.HandleFunc("/v1/public/trust-program/badges", srv.publicTrustBadgeProgramHandler)
	mux.HandleFunc("/v1/public/trust-program/badges/verify", srv.publicTrustBadgeVerificationHandler)
	mux.HandleFunc("/v1/public/trust-program/verifier-program", srv.publicVerifierProgramHandler)
	mux.HandleFunc("/v1/public/trust-program/claims-governance", srv.publicClaimsGovernanceHandler)
	mux.HandleFunc("/v1/public/trust-program/marks", srv.publicTrustMarkLifecycleHandler)
	mux.HandleFunc("/v1/public/trust-program/marks/", srv.publicTrustMarkByIDHandler)
	mux.HandleFunc("/v1/public/transparency/anchor", srv.publicTransparencyAnchorHandler)
	mux.HandleFunc("/v1/public/proof-portal", srv.publicProofPortalHandler)
	mux.HandleFunc("/v1/public/benchmarks/packs", srv.publicBenchmarkPacksHandler)
	mux.HandleFunc("/v1/public/reference/conformance", srv.publicReferenceConformanceHandler)
	mux.HandleFunc("/v1/public/verifier/sdk", srv.publicVerifierSDKHandler)
	mux.HandleFunc("/v1/public/claims/summary", srv.publicClaimsSummaryHandler)
	mux.HandleFunc("/v1/public/auditor/workflows", srv.publicAuditorWorkflowHandler)
	mux.HandleFunc("/v1/public/phase6/proofs", srv.publicPhase6ProofsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/val0/claim-registry-model", srv.publicProofVal0ClaimRegistryHandler)
	mux.HandleFunc("/v1/public/proof-expansion/val0/redaction-tiers", srv.publicProofVal0RedactionTiersHandler)
	mux.HandleFunc("/v1/public/proof-expansion/val0/signing-authority", srv.publicProofVal0SigningAuthorityHandler)
	mux.HandleFunc("/v1/public/proof-expansion/val0/compatibility-baseline", srv.publicProofVal0CompatibilityHandler)
	mux.HandleFunc("/v1/public/proof-expansion/val0/proofs", srv.publicProofVal0ProofsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vala/sealed-artifact-schema", srv.publicProofValAArtifactSchemaHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vala/sealing-discipline", srv.publicProofValASealingDisciplineHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vala/environment-binding", srv.publicProofValAEnvironmentBindingHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vala/downloadable-packs", srv.publicProofValADownloadablePacksHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vala/downloadable-packs/", srv.publicProofValAPackByIDHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vala/proofs", srv.publicProofValAProofsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valb/transparency-chain", srv.publicProofValBTransparencyChainHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valb/verifier-capability", srv.publicProofValBVerifierCapabilityHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valb/signature-verification", srv.publicProofValBSignatureVerificationHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valb/replay-verification", srv.publicProofValBReplayVerificationHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valb/proofs", srv.publicProofValBProofsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valc/public-proof-portal", srv.publicProofValCPublicPortalHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valc/partner-proof-portal", srv.publicProofValCPartnerPortalHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valc/claim-lineage", srv.publicProofValCClaimLineageHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valc/download-projections", srv.publicProofValCDownloadProjectionsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/valc/proofs", srv.publicProofValCProofsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vald/release-issuance-gate", srv.publicProofValDReleaseIssuanceHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vald/claim-lifecycle", srv.publicProofValDClaimLifecycleHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vald/publication-decisions", srv.publicProofValDPublicationDecisionsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vald/correction-workflow", srv.publicProofValDCorrectionWorkflowHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vald/proofs", srv.publicProofValDProofsHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/replay-correctness-review", srv.publicProofValEReplayCorrectnessReviewHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/signing-trust-review", srv.publicProofValESigningTrustReviewHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/transparency-review", srv.publicProofValETransparencyReviewHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/redaction-review", srv.publicProofValERedactionReviewHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/compatibility-review", srv.publicProofValECompatibilityReviewHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/issuance-review", srv.publicProofValEIssuanceReviewHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/failure-state-review", srv.publicProofValEFailureStateReviewHandler)
	mux.HandleFunc("/v1/public/proof-expansion/vale/proofs", srv.publicProofValEProofsHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/entry-gate", srv.phase7EntryGateHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/contracts", srv.phase7ContractsHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/developer-presence", srv.phase7DeveloperPresenceHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/developer/workbench", srv.phase7DeveloperWorkbenchHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/developer/context", srv.phase7DeveloperContextHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/developer/pre-commit", srv.phase7DeveloperPreCommitHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/oss-network", srv.phase7OSSNetworkHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/oss/connectors", srv.phase7OSSConnectorsHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/oss/observations", srv.phase7OSSObservationsHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/oss/review-flow", srv.phase7OSSReviewFlowHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/oss/reviewed-signals", srv.phase7OSSReviewedSignalsHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/distribution", srv.phase7DistributionHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/distribution/marketplace-readiness", srv.phase7MarketplaceReadinessHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/distribution/msp-isolation", srv.phase7MSPIsolationHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/distribution/partner-export", srv.phase7PartnerExportHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/proofs", srv.phase7ProofsHandler)
	mux.HandleFunc("/v1/ecosystem/phase7/final-summary", srv.phase7FinalSummaryHandler)
	mux.HandleFunc("/v1/formal/phase8/entry-gate", srv.phase8EntryGateHandler)
	mux.HandleFunc("/v1/formal/phase8/contracts", srv.phase8ContractsHandler)
	mux.HandleFunc("/v1/formal/phase8/formal-discipline", srv.phase8FormalDisciplineHandler)
	mux.HandleFunc("/v1/formal/phase8/compliance-codification", srv.phase8ComplianceCodificationHandler)
	mux.HandleFunc("/v1/formal/phase8/governed-autonomy", srv.phase8GovernedAutonomyHandler)
	mux.HandleFunc("/v1/formal/phase8/proofs", srv.phase8ProofsHandler)
	mux.HandleFunc("/v1/formal/phase8/compliance/policy-profiles", srv.phase8PolicyProfilesHandler)
	mux.HandleFunc("/v1/formal/phase8/compliance/regulatory-mappings", srv.phase8RegulatoryMappingsHandler)
	mux.HandleFunc("/v1/formal/phase8/compliance/verifier-surfaces", srv.phase8VerifierSurfacesHandler)
	mux.HandleFunc("/v1/formal/phase8/compliance/certification-workflow", srv.phase8CertificationWorkflowHandler)
	mux.HandleFunc("/v1/formal/phase8/compliance/evidence-automation", srv.phase8EvidenceAutomationHandler)
	mux.HandleFunc("/v1/formal/phase8/governance/consensus-review", srv.phase8ConsensusReviewHandler)
	mux.HandleFunc("/v1/formal/phase8/governance/policy-suggestions", srv.phase8PolicySuggestionsHandler)
	mux.HandleFunc("/v1/formal/phase8/governance/authority-routing", srv.phase8AuthorityRoutingHandler)
	mux.HandleFunc("/v1/formal/phase8/governance/ai-guardrails", srv.phase8AIGuardrailsHandler)
	mux.HandleFunc("/v1/formal/phase8/governance/model-risk", srv.phase8ModelRiskHandler)
	mux.HandleFunc("/v1/formal/phase8/institutional/risk-quantification", srv.phase8RiskQuantificationHandler)
	mux.HandleFunc("/v1/formal/phase8/institutional/insurance-exports", srv.phase8InsuranceExportsHandler)
	mux.HandleFunc("/v1/formal/phase8/institutional/incident-attribution", srv.phase8IncidentAttributionHandler)
	mux.HandleFunc("/v1/formal/phase8/institutional/actuarial-benchmarks", srv.phase8ActuarialBenchmarksHandler)
	mux.HandleFunc("/v1/formal/phase8/final-summary", srv.phase8FinalSummaryHandler)
	mux.HandleFunc("/v1/integrations/identity-fabric", srv.identityFabricHandler)
	mux.HandleFunc("/v1/integrations/itsm-lifecycle", srv.itsmLifecycleHandler)
	mux.HandleFunc("/v1/integrations/itsm-lifecycle/flows", srv.itsmLifecycleFlowsHandler)
	mux.HandleFunc("/v1/integrations/siem-sync", srv.siemSyncHandler)
	mux.HandleFunc("/v1/integrations/siem-sync/evaluate", srv.siemSyncHandler)
	mux.HandleFunc("/v1/integrations/safety", srv.integrationSafetyHandler)
	mux.HandleFunc("/v1/integrations/safety/health", srv.integrationSafetyHealthHandler)
	mux.HandleFunc("/v1/incidents/collaboration", srv.incidentCollaborationHandler)
	mux.HandleFunc("/v1/b2b/suppliers/onboarding", srv.b2bSupplierOnboardingHandler)
	mux.HandleFunc("/v1/b2b/sealed-proof/acceptance", srv.b2bSealedProofAcceptanceHandler)
	mux.HandleFunc("/v1/b2b/disclosure-profiles", srv.b2bDisclosureProfilesHandler)
	mux.HandleFunc("/v1/b2b/customer-bundles", srv.b2bCustomerBundleHandler)
	mux.HandleFunc("/v1/b2b/consortium-readiness", srv.b2bConsortiumReadinessHandler)
	mux.HandleFunc("/v1/trust-hub/governance", srv.trustHubGovernanceHandler)
	mux.HandleFunc("/v1/trust-hub/analytics", srv.trustHubAnalyticsHandler)
	mux.HandleFunc("/v1/trust-hub/clearance", srv.trustHubClearanceHandler)
	mux.HandleFunc("/v1/trust-hub/boundaries", srv.trustHubBoundariesHandler)
	mux.HandleFunc("/v1/audit/reports", srv.auditReportsHandler)
	mux.HandleFunc("/v1/audit/exports", srv.auditExportsHandler)
	mux.HandleFunc("/v1/ai/insights", srv.aiInsightsHandler)
	mux.HandleFunc("/v1/ai/vex-drafts", srv.aiVEXDraftsHandler)
	mux.HandleFunc("/v1/ai/break-glass-guidance", srv.aiBreakGlassGuidanceHandler)
	mux.HandleFunc("/v1/ai/guidance/", srv.aiGuidanceByIDHandler)
	mux.HandleFunc("/v1/ai/guidance", srv.aiGuidanceHandler)
	mux.HandleFunc("/v1/reports/events", srv.eventsHandler)
	mux.HandleFunc("/v1/reports/self-audit", srv.selfAuditSummaryHandler)
	mux.HandleFunc("/v1/reports/summary", srv.summaryHandler)
	mux.HandleFunc("/v1/reports/denies", srv.deniesHandler)
	mux.HandleFunc("/v1/reports/runtime-drift", srv.runtimeDriftHandler)
	mux.HandleFunc("/v1/runtime/desired-state", srv.runtimeDesiredStateHandler)
	mux.HandleFunc("/v1/runtime/active-state", srv.runtimeActiveStateHandler)
	mux.HandleFunc("/v1/runtime/quarantine", srv.runtimeQuarantineHandler)
	mux.HandleFunc("/v1/runtime/closed-loop/status", srv.runtimeClosedLoopStatusHandler)
	mux.HandleFunc("/v1/runtime/drift/status", srv.runtimeDriftStatusHandler)
	mux.HandleFunc("/v1/runtime/drift/", srv.runtimeDriftByIDHandler)
	mux.HandleFunc("/v1/runtime/drift", srv.runtimeDriftFindingsHandler)
	mux.HandleFunc("/v1/runtime/integrity", srv.runtimeIntegrityHandler)
	mux.HandleFunc("/v1/runtime/posture", srv.runtimePostureHandler)
	mux.HandleFunc("/v1/runtime/posture-linkage", srv.runtimePostureLinkageHandler)
	mux.HandleFunc("/v1/runtime/response-policy", srv.runtimeResponsePolicyHandler)
	mux.HandleFunc("/v1/runtime/response-tuning", srv.runtimeResponsePolicyHandler)
	mux.HandleFunc("/v1/runtime/boundaries", srv.runtimeBoundaryDisciplineHandler)
	mux.HandleFunc("/v1/runtime/workloads", srv.runtimeWorkloadsHandler)
	mux.HandleFunc("/v1/runtime/substrate-truth", srv.runtimeSubstrateTruthHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/entry-gate", srv.runtimeSubstrateDepthEntryGateHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vala/event-schema", srv.runtimeSubstrateValAEventSchemaHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vala/support-matrix", srv.runtimeSubstrateValASupportMatrixHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vala/observability", srv.runtimeSubstrateValAObservabilityHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vala/proofs", srv.runtimeSubstrateValAProofsHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valb/correlation-model", srv.runtimeSubstrateValBCorrelationModelHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valb/process-image-linkage", srv.runtimeSubstrateValBProcessImageHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valb/provenance-linkage", srv.runtimeSubstrateValBProvenanceHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valb/drift-catalog", srv.runtimeSubstrateValBDriftCatalogHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valb/proofs", srv.runtimeSubstrateValBProofsHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valc/enforcement-taxonomy", srv.runtimeSubstrateValCEnforcementTaxonomyHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valc/action-catalog", srv.runtimeSubstrateValCActionCatalogHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valc/policy-hook-mapping", srv.runtimeSubstrateValCPolicyHookMappingHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valc/decision-audit", srv.runtimeSubstrateValCDecisionAuditHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/valc/proofs", srv.runtimeSubstrateValCProofsHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vald/execution-class-matrix", srv.runtimeSubstrateValDExecutionClassMatrixHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vald/signal-coverage", srv.runtimeSubstrateValDSignalCoverageHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vald/enforcement-availability", srv.runtimeSubstrateValDEnforcementAvailabilityHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vald/overhead-visibility", srv.runtimeSubstrateValDOverheadVisibilityHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vald/proofs", srv.runtimeSubstrateValDProofsHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vale/latency-pack", srv.runtimeSubstrateValELatencyPackHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vale/false-positive-budget", srv.runtimeSubstrateValEFalsePositiveBudgetHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vale/replayable-benchmark-pack", srv.runtimeSubstrateValEReplayableBenchmarkPackHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vale/performance-gate", srv.runtimeSubstrateValEPerformanceGateHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/vale/proofs", srv.runtimeSubstrateValEProofsHandler)
	mux.HandleFunc("/v1/runtime/substrate-depth/complete", srv.runtimeSubstratePoint1CompleteHandler)
	mux.HandleFunc("/v1/runtime/trusted-execution-profiles", srv.runtimeTrustedExecutionProfilesHandler)
	mux.HandleFunc("/v1/runtime/attestation/verify", srv.runtimeAttestationVerifyHandler)
	mux.HandleFunc("/v1/runtime/attestation/verifications", srv.runtimeAttestationVerificationsHandler)
	mux.HandleFunc("/v1/runtime/response/simulate", srv.runtimeResponseSimulationHandler)
	mux.HandleFunc("/v1/runtime/response/rollback-drill", srv.runtimeRollbackDrillHandler)
	mux.HandleFunc("/v1/runtime/phase2/proofs", srv.runtimePhase2ProofsHandler)
	mux.HandleFunc("/v1/intelligence/vulnerability-relevance", srv.intelligenceVulnerabilityRelevanceHandler)
	mux.HandleFunc("/v1/intelligence/supply-chain/patterns", srv.intelligenceSupplyChainPatternsHandler)
	mux.HandleFunc("/v1/intelligence/strategic/simulate", srv.intelligenceStrategicSimulationHandler)
	mux.HandleFunc("/v1/intelligence/strategic/query", srv.intelligenceStrategicQueryHandler)
	mux.HandleFunc("/v1/intelligence/phase3/proofs", srv.intelligencePhase3ProofsHandler)
	mux.HandleFunc("/v1/enterprise/workflow/lifecycle", srv.enterpriseWorkflowLifecycleHandler)
	mux.HandleFunc("/v1/enterprise/workflow/connectors/reconcile", srv.enterpriseConnectorReconciliationHandler)
	mux.HandleFunc("/v1/enterprise/partner-trust/intake", srv.enterprisePartnerIntakeHandler)
	mux.HandleFunc("/v1/enterprise/partner-trust/dashboard", srv.enterprisePartnerDashboardHandler)
	mux.HandleFunc("/v1/enterprise/governance/compliance-mapping", srv.enterpriseComplianceMappingHandler)
	mux.HandleFunc("/v1/enterprise/governance/policy-drift", srv.enterprisePolicyDriftHandler)
	mux.HandleFunc("/v1/enterprise/governance/executive-report", srv.enterpriseExecutiveReportHandler)
	mux.HandleFunc("/v1/enterprise/phase4/proofs", srv.enterprisePhase4ProofsHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/authority-boundaries", srv.enterpriseWorkflowAuthorityVal0BoundaryHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/state-machine", srv.enterpriseWorkflowAuthorityVal0StateMachineHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/external-projection-rules", srv.enterpriseWorkflowAuthorityVal0ProjectionHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/approval-contract", srv.enterpriseWorkflowAuthorityVal0ApprovalContractHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/exception-lifecycle", srv.enterpriseWorkflowAuthorityVal0ExceptionLifecycleHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/closure-validation", srv.enterpriseWorkflowAuthorityVal0ClosureValidationHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/separation-of-duties", srv.enterpriseWorkflowAuthorityVal0SeparationHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/time-authority", srv.enterpriseWorkflowAuthorityVal0TimeAuthorityHandler)
	mux.HandleFunc("/v1/enterprise/workflow-authority/val0/proofs", srv.enterpriseWorkflowAuthorityVal0ProofsHandler)
	mux.HandleFunc("/v1/execution/coverage/matrix", srv.executionCoverageMatrixHandler)
	mux.HandleFunc("/v1/execution/coverage", srv.executionCoverageHandler)
	mux.HandleFunc("/v1/execution/vm-lineage", srv.executionVMLineageHandler)
	mux.HandleFunc("/v1/execution/ephemeral", srv.executionEphemeralHandler)
	mux.HandleFunc("/v1/execution/ambient-readiness", srv.executionAmbientReadinessHandler)
	mux.HandleFunc("/v1/execution/confidential-readiness", srv.executionConfidentialReadinessHandler)
	mux.HandleFunc("/v1/execution/compliance-readiness", srv.executionComplianceReadinessHandler)
	mux.HandleFunc("/v1/runtime/findings/", srv.runtimeFindingByIDHandler)
	mux.HandleFunc("/v1/runtime/findings", srv.runtimeFindingsHandler)
	mux.HandleFunc("/v1/runtime/rule-packs/", srv.runtimeRulePackByIDHandler)
	mux.HandleFunc("/v1/runtime/rule-packs", srv.runtimeRulePacksHandler)
	mux.HandleFunc("/v1/runtime/profiles/", srv.runtimeProfileBySubjectHandler)
	mux.HandleFunc("/v1/runtime/enforcement/evaluate", srv.runtimeEnforcementEvaluateHandler)
	mux.HandleFunc("/v1/runtime/enforcement", srv.runtimeEnforcementHandler)
	mux.HandleFunc("/v1/runtime/forensic-snapshot", srv.runtimeForensicSnapshotHandler)
	mux.HandleFunc("/v1/runtime/restart-trusted", srv.runtimeRestartTrustedHandler)
	mux.HandleFunc("/v1/hardening/posture", srv.hardeningPostureHandler)
	mux.HandleFunc("/v1/hardening/actions/", srv.hardeningActionByIDHandler)
	mux.HandleFunc("/v1/hardening/actions", srv.hardeningActionsHandler)
	mux.HandleFunc("/v1/hardening/evaluate", srv.hardeningEvaluateHandler)
	mux.HandleFunc("/v1/hardening/apply", srv.hardeningApplyHandler)
	mux.HandleFunc("/v1/hardening/rollback", srv.hardeningRollbackHandler)
	mux.HandleFunc("/v1/hardening/recover", srv.hardeningRecoverHandler)
	mux.HandleFunc("/v1/hardening/quarantine", srv.hardeningQuarantineHandler)
	mux.HandleFunc("/v1/hardening/divert-traffic", srv.hardeningDivertTrafficHandler)
	mux.HandleFunc("/v1/hardening/forensic-first", srv.hardeningForensicFirstHandler)
	mux.HandleFunc("/v1/reports/exceptions", srv.exceptionsReportHandler)
	mux.HandleFunc("/v1/self-audit/summary", srv.selfAuditSummaryHandler)
	mux.HandleFunc("/v1/self-audit/events", srv.selfAuditEventsHandler)
	return metrics.InstrumentHTTP("audit-writer", srv.wrap(mux))
}

func loadAuthConfigFromEnv() (auth.Config, error) {
	return auth.ParseEnvConfig(os.Getenv)
}

func newStoreFromEnv(ctx context.Context) (audit.Store, string, error) {
	storeKind := strings.ToLower(strings.TrimSpace(os.Getenv("CHANGELOCK_AUDIT_STORE")))
	dsn := strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_POSTGRES_DSN"), os.Getenv("DATABASE_URL")))

	switch storeKind {
	case "", "auto":
		if dsn == "" {
			return audit.NewMemoryStore(), "memory", nil
		}
		return newPostgresStore(ctx, dsn)
	case "memory":
		return audit.NewMemoryStore(), "memory", nil
	case "postgres":
		if dsn == "" {
			return nil, "", errors.New("CHANGELOCK_POSTGRES_DSN is required when CHANGELOCK_AUDIT_STORE=postgres")
		}
		return newPostgresStore(ctx, dsn)
	default:
		return nil, "", errors.New("unsupported CHANGELOCK_AUDIT_STORE: " + storeKind)
	}
}

func newPostgresStore(ctx context.Context, dsn string) (audit.Store, string, error) {
	store, err := audit.NewPostgresStore(ctx, dsn)
	if err != nil {
		return nil, "", err
	}
	if err := store.Migrate(ctx); err != nil {
		store.Close()
		return nil, "", err
	}
	return store, "postgres", nil
}

func (s server) healthHandler(w http.ResponseWriter, _ *http.Request) {
	httpjson.Write(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"backend": s.backend,
	})
}

func (s server) readyHandler(w http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.store.Ping(ctx); err != nil {
		httpjson.Write(w, http.StatusServiceUnavailable, map[string]string{
			"status":  "error",
			"backend": s.backend,
			"error":   err.Error(),
		})
		return
	}

	httpjson.Write(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"backend": s.backend,
	})
}

func (s server) ingestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	principal, err := s.requireIngestPrincipal(r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	r = r.WithContext(auth.WithPrincipal(r.Context(), principal))

	var event audit.Event
	if err := httpjson.Decode(r, &event); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if event.RequestID == "" {
		event.RequestID = requestIDFromHeader(r)
	}
	clusterID, err := s.resolveInboundClusterID(r, principal)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if clusterID != "" {
		event.ClusterID = clusterID
	} else if s.syncRuntime != nil && s.syncRuntime.config.Mode == audit.SyncModeSpoke && strings.TrimSpace(event.ClusterID) == "" {
		event.ClusterID = s.syncRuntime.config.ClusterID
	}
	event = audit.NormalizeEvent(event, time.Now)

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	record, err := s.store.Ingest(ctx, event)
	if err != nil {
		metrics.IncAuditStoreWriteFailure("audit-writer", s.backend)
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidEvent) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	metrics.IncAuditStoreWriteSuccess("audit-writer", s.backend)
	asyncForwardState := "not_applicable"
	if s.syncRuntime != nil && s.syncRuntime.config.Mode == audit.SyncModeSpoke {
		if task, err := s.enqueueSyncForwardTask(ctx, event.RequestID, principal.Subject, event); err != nil {
			log.Printf("audit-writer async sync forward enqueue failed: %v", err)
			asyncForwardState = "enqueue_failed"
		} else if task.TaskID != "" {
			asyncForwardState = "queued"
		}
	}
	trace := audit.NormalizeExecutionTraceRecord(audit.ExecutionTraceRecord{
		TraceID:       event.TraceID,
		Component:     "audit-writer",
		Operation:     "audit_ingest",
		TenantID:      event.TenantID,
		Environment:   event.Environment,
		EventID:       event.EventID,
		DecisionID:    event.DecisionID,
		CorrelationID: event.CorrelationID,
		Status:        "completed",
		StartedAt:     event.Timestamp,
		EndedAt:       record.ReceivedAt,
		Attributes: map[string]string{
			"async_sync_forward": asyncForwardState,
			"event_type":         event.EventType,
		},
		Notes: []string{
			"Canonical evidence write completed before any optional background forwarding.",
		},
	}, time.Now)
	if err := s.persistExecutionTrace(ctx, event.RequestID, principal.Subject, trace); err != nil {
		log.Printf("audit-writer execution trace persist failed: %v", err)
	}

	httpjson.Write(w, http.StatusCreated, ingestResponse{
		Status:     "stored",
		ID:         record.ID,
		RequestID:  record.RequestID,
		ReceivedAt: record.ReceivedAt,
	})
}

func (s server) authMeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	principal, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	httpjson.Write(w, http.StatusOK, authInfoResponse{
		Authenticated: principal.Authenticated,
		AuthMode:      principal.AuthMode,
		Subject:       principal.Subject,
		Role:          principal.Role,
		TokenID:       principal.TokenID,
		IdentityType:  principal.IdentityType,
		Email:         principal.Email,
		TenantID:      principal.TenantID,
		GlobalScope:   principal.GlobalScope,
	})
}

func (s server) eventsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	events, err := s.store.ListEvents(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, eventsResponse{Events: events})
}

func (s server) exceptionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		r, err := applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parseExceptionFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()

		exceptions, err := s.store.ListExceptions(ctx, filter)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrInvalidException) {
				status = http.StatusBadRequest
			} else if errors.Is(err, context.DeadlineExceeded) {
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}

		httpjson.Write(w, http.StatusOK, exceptionsResponse{Exceptions: exceptions})
	case http.MethodPost:
		if reason := s.exceptionMutationBlockedReason(); reason != "" {
			httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
			return
		}
		principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		var request audit.ExceptionCreateRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request, err := applyPrincipalTenantToExceptionRequest(principal, request)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()

		if strings.TrimSpace(request.ApprovedBy) == "" {
			request.ApprovedBy = principal.Subject
		}
		exception, err := s.store.CreateException(ctx, request)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrInvalidException) {
				status = http.StatusBadRequest
			} else if errors.Is(err, context.DeadlineExceeded) {
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		exception, err = s.signAndPersistException(ctx, exception)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionApproved, audit.DecisionAllow, exception, "direct emergency exception created as approved")
		httpjson.Write(w, http.StatusCreated, exceptionResponse{
			Status:    "created",
			Exception: exception,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) requestExceptionHandler(w http.ResponseWriter, r *http.Request) {
	if reason := s.exceptionMutationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
		return
	}
	principal, r, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request audit.ExceptionCreateRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request, err := applyPrincipalTenantToExceptionRequest(principal, request)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	exception, err := s.store.RequestException(ctx, request, principal.Subject, principal.Role)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionRequested, audit.DecisionAllow, exception, "exception requested for approval")
	httpjson.Write(w, http.StatusCreated, exceptionActionResponse{Status: "requested", Exception: exception})
}

func (s server) exceptionByIDHandler(w http.ResponseWriter, r *http.Request) {
	exceptionID, action, err := exceptionActionFromPath(r.URL.Path)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	switch {
	case action == "" && r.Method == http.MethodDelete:
		if reason := s.exceptionMutationBlockedReason(); reason != "" {
			httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
			return
		}
		principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
		if !ok {
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		existing, err := s.store.GetException(ctx, exceptionID)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, audit.ErrExceptionNotFound) {
				status = http.StatusNotFound
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}
		if err := ensureExceptionTenantAccess(principal, existing); err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}

		exception, err := s.store.RevokeException(ctx, exceptionID)
		if err != nil {
			status := http.StatusInternalServerError
			switch {
			case errors.Is(err, audit.ErrInvalidException):
				status = http.StatusBadRequest
			case errors.Is(err, audit.ErrExceptionNotFound):
				status = http.StatusNotFound
			case errors.Is(err, context.DeadlineExceeded):
				status = http.StatusGatewayTimeout
			}
			httpjson.Write(w, status, map[string]string{"error": err.Error()})
			return
		}

		s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionRevoked, audit.DecisionAllow, exception, "exception revoked")
		httpjson.Write(w, http.StatusOK, exceptionResponse{
			Status:    "revoked",
			Exception: exception,
		})
	case action == "approve" && r.Method == http.MethodPost:
		s.approveExceptionHandler(w, r, exceptionID)
	case action == "reject" && r.Method == http.MethodPost:
		s.rejectExceptionHandler(w, r, exceptionID)
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) approveExceptionHandler(w http.ResponseWriter, r *http.Request, exceptionID string) {
	if reason := s.exceptionMutationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
		return
	}
	principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}

	var request audit.ExceptionActionRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	existing, err := s.store.GetException(ctx, exceptionID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrExceptionNotFound) {
			status = http.StatusNotFound
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	if err := ensureExceptionTenantAccess(principal, existing); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	exception, err := s.store.ApproveException(ctx, exceptionID, principal.Subject, principal.Role)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, audit.ErrInvalidException):
			status = http.StatusBadRequest
		case errors.Is(err, audit.ErrExceptionNotFound):
			status = http.StatusNotFound
		case errors.Is(err, context.DeadlineExceeded):
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	exception, err = s.signAndPersistException(ctx, exception)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	reason := "exception approved"
	if normalized := audit.NormalizeExceptionActionRequest(request); normalized.Reason != "" {
		reason = normalized.Reason
	}
	s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionApproved, audit.DecisionAllow, exception, reason)
	httpjson.Write(w, http.StatusOK, exceptionActionResponse{Status: "approved", Exception: exception})
}

func (s server) rejectExceptionHandler(w http.ResponseWriter, r *http.Request, exceptionID string) {
	if reason := s.exceptionMutationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": reason})
		return
	}
	principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}

	var request audit.ExceptionActionRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request = audit.NormalizeExceptionActionRequest(request)
	if request.Reason == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "reason is required"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	existing, err := s.store.GetException(ctx, exceptionID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrExceptionNotFound) {
			status = http.StatusNotFound
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	if err := ensureExceptionTenantAccess(principal, existing); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	exception, err := s.store.RejectException(ctx, exceptionID, request.Reason, principal.Subject, principal.Role)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, audit.ErrInvalidException):
			status = http.StatusBadRequest
		case errors.Is(err, audit.ErrExceptionNotFound):
			status = http.StatusNotFound
		case errors.Is(err, context.DeadlineExceeded):
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	s.writeLifecycleAuditEvent(ctx, r, principal.Subject, audit.EventTypeExceptionRejected, audit.DecisionDeny, exception, request.Reason)
	httpjson.Write(w, http.StatusOK, exceptionActionResponse{Status: "rejected", Exception: exception})
}

func (s server) validateExceptionHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleService, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var request audit.ExceptionValidationRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if principal.Role != auth.RoleService {
		updatedRequest, err := applyPrincipalTenantToExceptionValidation(principal, request)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		request = updatedRequest
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	if reason := s.exceptionValidationBlockedReason(); reason != "" {
		httpjson.Write(w, http.StatusOK, audit.ExceptionValidationResult{Valid: false, Reason: reason})
		return
	}

	result, err := s.store.ValidateException(ctx, request)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	if result.Exception != nil {
		verification, err := s.signing.verifyException(ctx, *result.Exception)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		result.VerificationState = verification.State
		result.VerificationReason = verification.Reason
		result.Exception.VerificationState = verification.State
		result.Exception.VerificationReason = verification.Reason
		if s.signing.verifyOnRead(signing.PurposeExceptions) && verification.State != signing.StateVerified {
			result.Valid = false
			if result.Reason == "" {
				result.Reason = firstNonEmpty(verification.Reason, "exception evidence verification failed")
			}
		}
	} else {
		result.VerificationState = signing.StateDisabled
	}

	httpjson.Write(w, http.StatusOK, result)
}

func (s server) trendsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseTrendsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.store.Trends(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	analyticsFilter, err := parseAnalyticsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	response, err = s.buildAnalyticsTrendsResponse(ctx, analyticsFilter, response)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	response.SchemaVersion = audit.TrendsSchemaVersion

	httpjson.Write(w, http.StatusOK, response)
}

func (s server) topViolatorsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseTopViolatorsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.store.TopViolators(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	response.SchemaVersion = audit.TopViolatorsSchemaVersion

	httpjson.Write(w, http.StatusOK, response)
}

func (s server) driftStatsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseDriftStatsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.store.DriftStats(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	response.SchemaVersion = audit.DriftStatsSchemaVersion

	httpjson.Write(w, http.StatusOK, response)
}

func (s server) summaryHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	summary, err := s.store.Summary(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, summary)
}

func (s server) deniesHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	query := r.URL.Query()
	query.Set("decision", audit.DecisionDeny)
	r.URL.RawQuery = query.Encode()
	s.eventsHandler(w, r)
}

func (s server) runtimeDriftHandler(w http.ResponseWriter, r *http.Request) {
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}
	query := r.URL.Query()
	query.Set("event_type", audit.EventTypeRuntimeDriftResult)
	r.URL.RawQuery = query.Encode()
	s.eventsHandler(w, r)
}

func (s server) exceptionsReportHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseExceptionFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	report, err := s.store.ExceptionReport(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidException) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}

	httpjson.Write(w, http.StatusOK, report)
}

func (s server) syncStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	if _, _, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin); !ok {
		return
	}

	status := deriveSyncStatus(audit.SyncStatus{
		SyncMode: audit.SyncModeDisabled,
		Mode:     audit.SyncModeDisabled,
		Health:   audit.SyncHealthDisabled,
	}, syncConfig{}, time.Now().UTC())
	if s.syncRuntime != nil {
		status = s.syncRuntime.statusSnapshot()
		if s.syncRuntime.config.Mode == audit.SyncModeHub {
			revision, err := s.currentHubExceptionRevision(r.Context())
			if err == nil {
				status.CurrentRevision = revision
				status.RevisionETag = quotedETag(revision)
			}
		}
		status = deriveSyncStatus(status, s.syncRuntime.config, time.Now().UTC())
	}
	if s.signing != nil {
		status.SignerMode = s.signing.mode()
	}
	httpjson.Write(w, http.StatusOK, status)
}

func (s server) syncExceptionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	if s.syncRuntime == nil || s.syncRuntime.config.Mode != audit.SyncModeHub {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "sync hub endpoint is disabled"})
		return
	}
	principal, _, ok := s.authorize(w, r, auth.RoleService)
	if !ok {
		return
	}
	clusterID := strings.TrimSpace(r.Header.Get(syncClusterHeader))
	binding, err := s.syncRuntime.authorizeClusterPrincipal(principal, clusterID)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	filtered, revision, err := s.currentHubSyncedExceptions(r.Context(), binding)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	etag := fmt.Sprintf("%q", revision)
	if strings.TrimSpace(r.Header.Get("If-None-Match")) == etag {
		w.Header().Set("ETag", etag)
		w.WriteHeader(http.StatusNotModified)
		return
	}

	response := audit.ExceptionSyncSnapshot{
		ClusterID:   clusterID,
		Revision:    revision,
		GeneratedAt: time.Now().UTC(),
		Exceptions:  filtered,
	}
	envelope, err := s.signing.signSyncSnapshot(r.Context(), response)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	response.Signature = envelope
	w.Header().Set("ETag", etag)
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) authorize(w http.ResponseWriter, r *http.Request, roles ...string) (auth.Principal, *http.Request, bool) {
	principal, err := s.authConfig.Require(r, roles...)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return auth.Principal{}, r, false
	}
	r = r.WithContext(auth.WithPrincipal(r.Context(), principal))
	return principal, r, true
}

func parseFilter(r *http.Request) (audit.EventFilter, error) {
	query := r.URL.Query()
	filter := audit.EventFilter{
		Decision:    query.Get("decision"),
		EventType:   query.Get("event_type"),
		Component:   query.Get("component"),
		ClusterID:   query.Get("cluster_id"),
		Repo:        query.Get("repo"),
		Environment: query.Get("environment"),
		TenantID:    query.Get("tenant_id"),
	}
	if rawLimit := strings.TrimSpace(query.Get("limit")); rawLimit != "" {
		limit, err := strconv.Atoi(rawLimit)
		if err != nil {
			return audit.EventFilter{}, errors.New("limit must be an integer")
		}
		filter.Limit = limit
	}
	return audit.NormalizeFilter(filter)
}

func parseExceptionFilter(r *http.Request) (audit.ExceptionFilter, error) {
	query := r.URL.Query()
	filter := audit.ExceptionFilter{
		Status:        query.Get("status"),
		ExceptionType: query.Get("exception_type"),
		TenantID:      query.Get("tenant_id"),
		Environment:   query.Get("environment"),
		Namespace:     query.Get("namespace"),
		Repo:          query.Get("repo"),
		ImageDigest:   query.Get("image_digest"),
		CVEID:         query.Get("cve_id"),
	}

	if rawActive := strings.TrimSpace(query.Get("active")); rawActive != "" {
		active, err := strconv.ParseBool(rawActive)
		if err != nil {
			return audit.ExceptionFilter{}, errors.New("active must be a boolean")
		}
		filter.Active = &active
	}

	if rawLimit := strings.TrimSpace(query.Get("limit")); rawLimit != "" {
		limit, err := strconv.Atoi(rawLimit)
		if err != nil {
			return audit.ExceptionFilter{}, errors.New("limit must be an integer")
		}
		filter.Limit = limit
	}

	return audit.NormalizeExceptionFilter(filter)
}

func (s server) requireIngestPrincipal(r *http.Request) (auth.Principal, error) {
	if s.authConfig.Mode != auth.ModeDisabled {
		return s.authConfig.Require(r, auth.RoleService)
	}
	if strings.TrimSpace(s.internalToken) == "" {
		return auth.Principal{}, auth.ErrMissingBearerToken
	}
	token, err := bearerTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		return auth.Principal{}, err
	}
	if subtle.ConstantTimeCompare([]byte(token), []byte(s.internalToken)) != 1 {
		return auth.Principal{}, auth.ErrInvalidBearerToken
	}
	return auth.Principal{
		Authenticated: true,
		AuthMode:      auth.ModeDisabled,
		Subject:       "internal-service",
		Role:          auth.RoleService,
		TokenID:       "internal-service",
		IdentityType:  auth.IdentityTypeService,
		GlobalScope:   true,
	}, nil
}

func bearerTokenFromHeader(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", auth.ErrMissingBearerToken
	}
	parts := strings.Fields(value)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
		return "", auth.ErrMalformedAuthorization
	}
	return strings.TrimSpace(parts[1]), nil
}

func (s server) resolveInboundClusterID(r *http.Request, principal auth.Principal) (string, error) {
	clusterID := strings.TrimSpace(r.Header.Get(syncClusterHeader))
	if clusterID == "" {
		return "", nil
	}
	if s.syncRuntime == nil || s.syncRuntime.config.Mode != audit.SyncModeHub {
		return clusterID, nil
	}
	if _, err := s.syncRuntime.authorizeClusterPrincipal(principal, clusterID); err != nil {
		return "", err
	}
	return clusterID, nil
}

func (s server) exceptionMutationBlockedReason() string {
	if s.syncRuntime == nil {
		return ""
	}
	return s.syncRuntime.mutationBlockedReason()
}

func (s server) exceptionValidationBlockedReason() string {
	if s.syncRuntime == nil {
		return ""
	}
	return s.syncRuntime.exceptionValidationBlockReason()
}

func (s server) currentHubSyncedExceptions(ctx context.Context, binding clusterBinding) ([]audit.SyncedException, string, error) {
	filter := audit.ExceptionFilter{
		Status: audit.ExceptionStatusApproved,
		Limit:  500,
	}
	var exceptions []audit.PolicyException
	if len(binding.Tenants) == 0 {
		listed, err := s.store.ListExceptions(ctx, filter)
		if err != nil {
			return nil, "", err
		}
		exceptions = listed
	} else {
		byID := map[string]audit.PolicyException{}
		for _, tenantID := range binding.Tenants {
			filter.TenantID = tenantID
			listed, err := s.store.ListExceptions(ctx, filter)
			if err != nil {
				return nil, "", err
			}
			for _, exception := range listed {
				byID[exception.ExceptionID] = exception
			}
		}
		exceptions = make([]audit.PolicyException, 0, len(byID))
		for _, exception := range byID {
			exceptions = append(exceptions, exception)
		}
	}

	filtered := filterSyncedExceptionsForBinding(exceptions, binding)
	revision := audit.ComputeExceptionSyncRevision(filtered)
	return filtered, revision, nil
}

func (s server) currentHubExceptionRevision(ctx context.Context) (string, error) {
	exceptions, err := s.store.ListExceptions(ctx, audit.ExceptionFilter{
		Status: audit.ExceptionStatusApproved,
		Limit:  500,
	})
	if err != nil {
		return "", err
	}
	synced := make([]audit.SyncedException, 0, len(exceptions))
	for _, exception := range exceptions {
		synced = append(synced, audit.SyncedExceptionFromPolicyException(exception))
	}
	return audit.ComputeExceptionSyncRevision(synced), nil
}

func parseTrendsFilter(r *http.Request) (audit.TrendsFilter, error) {
	query := r.URL.Query()
	filter := audit.TrendsFilter{
		WindowDays:  30,
		Granularity: query.Get("granularity"),
		ClusterID:   query.Get("cluster_id"),
		TenantID:    query.Get("tenant_id"),
		Environment: query.Get("environment"),
		Repo:        query.Get("repo"),
		EventType:   query.Get("event_type"),
	}
	if raw := strings.TrimSpace(query.Get("window_days")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.TrendsFilter{}, errors.New("window_days must be an integer")
		}
		filter.WindowDays = value
	}
	return audit.NormalizeTrendsFilter(filter)
}

func parseTopViolatorsFilter(r *http.Request) (audit.TopViolatorsFilter, error) {
	query := r.URL.Query()
	filter := audit.TopViolatorsFilter{
		WindowDays:  30,
		Limit:       5,
		Dimension:   query.Get("dimension"),
		ClusterID:   query.Get("cluster_id"),
		TenantID:    query.Get("tenant_id"),
		Environment: query.Get("environment"),
		Repo:        query.Get("repo"),
	}
	if raw := strings.TrimSpace(query.Get("window_days")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.TopViolatorsFilter{}, errors.New("window_days must be an integer")
		}
		filter.WindowDays = value
	}
	if raw := strings.TrimSpace(query.Get("limit")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.TopViolatorsFilter{}, errors.New("limit must be an integer")
		}
		filter.Limit = value
	}
	return audit.NormalizeTopViolatorsFilter(filter)
}

func parseDriftStatsFilter(r *http.Request) (audit.DriftStatsFilter, error) {
	query := r.URL.Query()
	filter := audit.DriftStatsFilter{
		WindowDays:  30,
		ClusterID:   query.Get("cluster_id"),
		TenantID:    query.Get("tenant_id"),
		Environment: query.Get("environment"),
		Repo:        query.Get("repo"),
		Namespace:   query.Get("namespace"),
		Workload:    query.Get("workload"),
	}
	if raw := strings.TrimSpace(query.Get("window_days")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.DriftStatsFilter{}, errors.New("window_days must be an integer")
		}
		filter.WindowDays = value
	}
	return audit.NormalizeDriftStatsFilter(filter)
}

func exceptionActionFromPath(path string) (string, string, error) {
	raw := strings.TrimPrefix(path, "/v1/exceptions/")
	raw = strings.TrimSpace(strings.Trim(raw, "/"))
	if raw == "" || raw == "validate" {
		return "", "", errors.New("exception_id path segment is required")
	}
	parts := strings.Split(raw, "/")
	if len(parts) > 2 {
		return "", "", errors.New("invalid exception path")
	}
	value, err := url.PathUnescape(parts[0])
	if err != nil {
		return "", "", errors.New("invalid exception_id path segment")
	}
	if strings.TrimSpace(value) == "" {
		return "", "", errors.New("exception_id path segment is required")
	}
	action := ""
	if len(parts) == 2 {
		action = strings.TrimSpace(parts[1])
	}
	return value, action, nil
}

func requestIDFromHeader(r *http.Request) string {
	if requestID := strings.TrimSpace(r.Header.Get("X-Request-Id")); requestID != "" {
		return requestID
	}
	return audit.NewRequestID()
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func envDurationOrDefault(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}

func firstNonEmpty(values ...string) string {
	return audit.FirstNonEmpty(values...)
}

func allowedOriginsFromEnv() map[string]struct{} {
	raw := strings.TrimSpace(os.Getenv("CHANGELOCK_CORS_ALLOW_ORIGINS"))
	if raw == "" {
		raw = strings.Join([]string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"http://localhost:3000",
			"http://127.0.0.1:3000",
		}, ",")
	}

	allowed := map[string]struct{}{}
	for _, origin := range strings.Split(raw, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowed[origin] = struct{}{}
		}
	}
	return allowed
}

func (s server) wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.applySecurityHeaders(w, r)
		if s.handleCORS(w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s server) applySecurityHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if r.URL.Path == "/health" || r.URL.Path == "/ready" || strings.HasPrefix(r.URL.Path, "/v1/") {
		w.Header().Set("Cache-Control", "no-store, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
	}
}

func (s server) handleCORS(w http.ResponseWriter, r *http.Request) bool {
	origin := strings.TrimSpace(r.Header.Get("Origin"))
	if origin != "" {
		w.Header().Add("Vary", "Origin")
	}
	w.Header().Add("Vary", "Access-Control-Request-Method")
	w.Header().Add("Vary", "Access-Control-Request-Headers")

	if origin != "" {
		if _, ok := s.allowedOrigins[origin]; ok {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Request-Id")
			w.Header().Set("Access-Control-Max-Age", "600")
		} else if r.Method == http.MethodOptions {
			httpjson.Write(w, http.StatusForbidden, map[string]string{"error": "origin not allowed"})
			return true
		}
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}

func (s server) writeLifecycleAuditEvent(ctx context.Context, r *http.Request, actor, eventType, decision string, exception audit.PolicyException, reason string) {
	if actor == "" {
		actor = firstNonEmpty(exception.ApprovedBy, exception.RequestedBy, exception.RejectedBy)
	}
	_, _ = s.store.Ingest(ctx, audit.Event{
		RequestID:                requestIDFromHeader(r),
		Component:                "audit-writer",
		EventType:                eventType,
		Actor:                    actor,
		TenantID:                 exception.TenantID,
		Repo:                     exception.Repo,
		Environment:              exception.Environment,
		Namespace:                exception.Namespace,
		Digest:                   exception.ImageDigest,
		CVEID:                    exception.CVEID,
		Decision:                 decision,
		Reasons:                  []string{reason},
		IsException:              true,
		ExceptionID:              exception.ExceptionID,
		ExceptionType:            exception.ExceptionType,
		ExceptionStatus:          exception.Status,
		ExceptionReason:          exception.Reason,
		ExceptionTicketID:        exception.TicketID,
		ExceptionRequestedBy:     exception.RequestedBy,
		ExceptionRequestedAt:     exception.RequestedAt,
		ExceptionApprovedBy:      exception.ApprovedBy,
		ExceptionApprovedAt:      exception.ApprovedAt,
		ExceptionRejectedBy:      exception.RejectedBy,
		ExceptionRejectedAt:      exception.RejectedAt,
		ExceptionRejectionReason: exception.RejectionReason,
		ExceptionExpiresAt:       &exception.ExpiresAt,
	})
}
