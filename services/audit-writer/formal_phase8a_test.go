package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
)

func TestPhase8PolicyProfilesAndMappingsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/compliance/policy-profiles", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 policy profiles 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var profiles phase8PolicyProfilesResponse
	if err := json.NewDecoder(rec.Body).Decode(&profiles); err != nil {
		t.Fatalf("decode phase8 policy profiles: %v", err)
	}
	if profiles.CurrentState != phase8PolicyProfilesStateActive {
		t.Fatalf("expected active phase8 policy profile slice, got %#v", profiles)
	}
	if len(profiles.Profiles) == 0 || profiles.Profiles[0].MachineCheckableCoverageRate == "" {
		t.Fatalf("expected populated policy profiles with machine-checkable coverage, got %#v", profiles)
	}
	if !containsString(profiles.ActivationBoundaries, "conflicting_jurisdiction_profiles_require_manual_resolution") {
		t.Fatalf("expected manual resolution boundary, got %#v", profiles.ActivationBoundaries)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/compliance/regulatory-mappings", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 regulatory mappings 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var mappings phase8RegulatoryMappingsResponse
	if err := json.NewDecoder(rec.Body).Decode(&mappings); err != nil {
		t.Fatalf("decode phase8 regulatory mappings: %v", err)
	}
	if mappings.CurrentState != phase8RegulatoryMappingsStateActive {
		t.Fatalf("expected active phase8 regulatory mapping slice, got %#v", mappings)
	}
	if mappings.MappingState != "regulatory_mapping_pack_active" || mappings.ConflictHandlingState != "control_conflict_handling_visible" {
		t.Fatalf("expected active mapping states, got %#v", mappings)
	}
	if len(mappings.Mappings) == 0 || len(mappings.Mappings[0].ControlConflictMarkers) == 0 || len(mappings.Mappings[0].CompensatingControlSemantics) == 0 {
		t.Fatalf("expected populated mapping semantics, got %#v", mappings)
	}
}

func TestPhase8VerifierSurfacesAndCertificationWorkflowHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/compliance/verifier-surfaces", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 verifier surfaces 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var surfaces phase8VerifierSurfacesResponse
	if err := json.NewDecoder(rec.Body).Decode(&surfaces); err != nil {
		t.Fatalf("decode phase8 verifier surfaces: %v", err)
	}
	if surfaces.CurrentState != phase8VerifierSurfacesStateActive {
		t.Fatalf("expected active phase8 verifier surface slice, got %#v", surfaces)
	}
	if !containsString(surfaces.AllowedAudienceClasses, formalcore.AudienceRegulator) || !containsString(surfaces.AllowedAudienceClasses, formalcore.AudienceCertification) {
		t.Fatalf("expected regulator and certification audiences only, got %#v", surfaces.AllowedAudienceClasses)
	}
	if !containsString(surfaces.ForbiddenAudiences, formalcore.AudienceInsurer) || !containsString(surfaces.ForbiddenAudiences, "public") {
		t.Fatalf("expected insurer and public audiences forbidden, got %#v", surfaces.ForbiddenAudiences)
	}
	for _, surface := range surfaces.Surfaces {
		if surface.AudienceClass == formalcore.AudienceInsurer {
			t.Fatalf("expected verifier slice to stay free of insurer audience, got %#v", surfaces.Surfaces)
		}
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/compliance/certification-workflow", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 certification workflow 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var workflow phase8CertificationWorkflowResponse
	if err := json.NewDecoder(rec.Body).Decode(&workflow); err != nil {
		t.Fatalf("decode phase8 certification workflow: %v", err)
	}
	if workflow.CurrentState != phase8CertificationWorkflowStateActive {
		t.Fatalf("expected active phase8 certification workflow slice, got %#v", workflow)
	}
	if workflow.ReleaseBoundaryState != "support_artifact_not_certification_issuance" || workflow.SnapshotState != "assessment_snapshot_and_evidence_freeze_visible" {
		t.Fatalf("expected bounded certification workflow states, got %#v", workflow)
	}
	if !containsString(workflow.Lifecycle.States, "approved_for_release") || !containsString(workflow.Lifecycle.States, "challenged") {
		t.Fatalf("expected lifecycle states to remain visible, got %#v", workflow.Lifecycle)
	}
	if !containsString(workflow.RequiredApprovalActions, "certification_support_release_approval") {
		t.Fatalf("expected certification release approval to remain non-delegable, got %#v", workflow.RequiredApprovalActions)
	}
	if len(workflow.EvidencePacks) == 0 || !workflow.EvidencePacks[0].EvidenceFreezeForSnapshot {
		t.Fatalf("expected evidence freeze for certification snapshot, got %#v", workflow.EvidencePacks)
	}
}

func TestPhase8EvidenceAutomationHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/compliance/evidence-automation", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 evidence automation 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var automation phase8EvidenceAutomationResponse
	if err := json.NewDecoder(rec.Body).Decode(&automation); err != nil {
		t.Fatalf("decode phase8 evidence automation: %v", err)
	}
	if automation.CurrentState != phase8EvidenceAutomationStateActive {
		t.Fatalf("expected active phase8 evidence automation slice, got %#v", automation)
	}
	if automation.PackCompletenessScore != "bounded_machine_checkable_plus_manual_interpretation_visible" {
		t.Fatalf("expected bounded completeness score, got %#v", automation)
	}
	if len(automation.IncludedArtifacts) == 0 || !containsString(automation.ChallengeLinked, "certification_support_pack_baseline") {
		t.Fatalf("expected populated included artifacts and challenge links, got %#v", automation)
	}

	foundInsurerDeferred := false
	foundNoCertificationIssuance := false
	for _, item := range automation.NotInPackReasons {
		if item.ItemID == "insurance_facing_evidence_exports" && item.DeferredScope {
			foundInsurerDeferred = true
		}
		if item.ItemID == "direct_certification_issuance" && !item.DeferredScope {
			foundNoCertificationIssuance = true
		}
	}
	if !foundInsurerDeferred || !foundNoCertificationIssuance {
		t.Fatalf("expected insurer exports deferred and direct certification issuance excluded, got %#v", automation.NotInPackReasons)
	}
}
