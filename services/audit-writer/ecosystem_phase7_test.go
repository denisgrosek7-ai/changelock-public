package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	ecosystemcore "github.com/denisgrosek/changelock/internal/ecosystem"
)

func TestPhase7EntryGateAndContractsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	entryReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/entry-gate", nil)
	entryRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(entryRec, entryReq)
	if entryRec.Code != http.StatusOK {
		t.Fatalf("expected entry gate 200, got %d: %s", entryRec.Code, entryRec.Body.String())
	}

	var entry phase7EntryGateResponse
	if err := json.NewDecoder(entryRec.Body).Decode(&entry); err != nil {
		t.Fatalf("decode entry gate: %v", err)
	}
	if entry.CurrentState != ecosystemcore.EntryGateStateReady || entry.CanonicalWorkspace == "" {
		t.Fatalf("expected ready entry gate with canonical workspace, got %#v", entry)
	}
	if !containsString(entry.ContractRefs, "phase7.signal_contract_matrix") {
		t.Fatalf("expected entry gate contract refs, got %#v", entry.ContractRefs)
	}

	contractsReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/contracts", nil)
	contractsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(contractsRec, contractsReq)
	if contractsRec.Code != http.StatusOK {
		t.Fatalf("expected contracts 200, got %d: %s", contractsRec.Code, contractsRec.Body.String())
	}

	var contracts phase7ContractsResponse
	if err := json.NewDecoder(contractsRec.Body).Decode(&contracts); err != nil {
		t.Fatalf("decode contracts: %v", err)
	}
	if contracts.CurrentState != ecosystemcore.FoundationStateActive {
		t.Fatalf("expected active contracts foundation, got %#v", contracts)
	}
	if contracts.Coverage.SignalContracts < 8 || contracts.Coverage.AuthoritySurfaces < 8 || contracts.Coverage.DataBoundaries < 4 {
		t.Fatalf("expected non-trivial phase7 coverage, got %#v", contracts.Coverage)
	}
	if !hasSignalContract(contracts.SignalContracts, "developer.ide_trust_advisory") {
		t.Fatalf("expected developer signal contract in contracts matrix, got %#v", contracts.SignalContracts)
	}
	if !hasBoundary(contracts.DataBoundaries, "distribution.partner_export") {
		t.Fatalf("expected partner export boundary, got %#v", contracts.DataBoundaries)
	}
}

func TestPhase7DeveloperAndOSSHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	developerReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/developer-presence", nil)
	developerRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(developerRec, developerReq)
	if developerRec.Code != http.StatusOK {
		t.Fatalf("expected developer presence 200, got %d: %s", developerRec.Code, developerRec.Body.String())
	}

	var developer phase7DeveloperPresenceResponse
	if err := json.NewDecoder(developerRec.Body).Decode(&developer); err != nil {
		t.Fatalf("decode developer presence: %v", err)
	}
	if developer.CurrentState != ecosystemcore.DeveloperPresenceStateActive {
		t.Fatalf("expected active developer presence, got %#v", developer)
	}
	if !containsString(developer.OutputSemantics, ecosystemcore.SignalClassObservedFact) || !containsString(developer.OutputSemantics, "uncertainty") {
		t.Fatalf("expected bounded developer output semantics, got %#v", developer.OutputSemantics)
	}
	if !containsString(developer.SignalRefs, "developer.pre_commit_dependency_trust") {
		t.Fatalf("expected pre-commit signal ref, got %#v", developer.SignalRefs)
	}

	ossReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/oss-network", nil)
	ossRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(ossRec, ossReq)
	if ossRec.Code != http.StatusOK {
		t.Fatalf("expected oss network 200, got %d: %s", ossRec.Code, ossRec.Body.String())
	}

	var oss phase7OSSNetworkResponse
	if err := json.NewDecoder(ossRec.Body).Decode(&oss); err != nil {
		t.Fatalf("decode oss network: %v", err)
	}
	if oss.CurrentState != ecosystemcore.OSSPresenceStateActive {
		t.Fatalf("expected active oss presence, got %#v", oss)
	}
	if oss.ObservationPipeline.CurrentState != "observation_pipeline_active" || oss.ClaimPipeline.CurrentState != "claim_pipeline_active" {
		t.Fatalf("expected active separated pipelines, got %#v", oss)
	}
	if oss.AutomatedPRState != phase7DeferredExpandedScopeTag {
		t.Fatalf("expected automated PR to remain deferred from core pass, got %#v", oss)
	}
	if len(oss.Connectors) < 3 || len(oss.ReviewedSignals) == 0 {
		t.Fatalf("expected registry connectors and reviewed signals, got %#v", oss)
	}
}

func TestPhase7DistributionAndProofsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	distributionReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/distribution", nil)
	distributionRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(distributionRec, distributionReq)
	if distributionRec.Code != http.StatusOK {
		t.Fatalf("expected distribution presence 200, got %d: %s", distributionRec.Code, distributionRec.Body.String())
	}

	var distribution phase7DistributionResponse
	if err := json.NewDecoder(distributionRec.Body).Decode(&distribution); err != nil {
		t.Fatalf("decode distribution presence: %v", err)
	}
	if distribution.CurrentState != ecosystemcore.DistributionPresenceStateActive {
		t.Fatalf("expected active distribution presence, got %#v", distribution)
	}
	if distribution.MSPIsolation.TenantIsolation != "strict_tenant_isolation_verified" {
		t.Fatalf("expected strict tenant isolation, got %#v", distribution.MSPIsolation)
	}
	if containsString(distribution.PartnerSurface.AllowedOperations, "cross_tenant_read") {
		t.Fatalf("expected partner surface to keep cross-tenant reads forbidden, got %#v", distribution.PartnerSurface)
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/proofs", nil)
	proofsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected phase7 proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}

	var proofs phase7ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode phase7 proofs: %v", err)
	}
	if proofs.CurrentState != ecosystemcore.Phase7StateActive {
		t.Fatalf("expected active phase7 core proofs, got %#v", proofs)
	}
	if proofs.CoverageScope != phase7CoverageScopeCorePass {
		t.Fatalf("expected core-pass coverage scope, got %#v", proofs)
	}
	if !containsString(proofs.ExpandedScopeDeferred, "automated_pr_discipline") {
		t.Fatalf("expected deferred expanded scope in final proofs, got %#v", proofs.ExpandedScopeDeferred)
	}
}

func hasSignalContract(items []ecosystemcore.SignalContract, signalID string) bool {
	for _, item := range items {
		if item.SignalID == signalID {
			return true
		}
	}
	return false
}

func hasBoundary(items []ecosystemcore.DataBoundary, surfaceID string) bool {
	for _, item := range items {
		if item.SurfaceID == surfaceID {
			return true
		}
	}
	return false
}
