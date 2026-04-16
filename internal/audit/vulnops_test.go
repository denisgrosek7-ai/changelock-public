package audit

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestMemoryStoreIngestSBOMStoresDocumentsAndSearchableComponents(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()

	spdxPayload := []byte(`{
	  "packages": [
	    {
	      "name": "openssl",
	      "versionInfo": "3.0.14-r0",
	      "licenseConcluded": "Apache-2.0",
	      "primaryPackagePurpose": "LIBRARY",
	      "SPDXID": "SPDXRef-openssl",
	      "externalRefs": [{"referenceType":"purl","referenceLocator":"pkg:apk/alpine/openssl@3.0.14-r0"}]
	    },
	    {
	      "name": "busybox",
	      "versionInfo": "1.36.1-r2",
	      "licenseDeclared": "GPL-2.0-only",
	      "primaryPackagePurpose": "APPLICATION"
	    }
	  ]
	}`)
	result, err := store.IngestSBOM(ctx, SBOMIngestRequest{
		ImageDigest: "sha256:sbom1",
		ImageRef:    "ghcr.io/example/app:1.0.0",
		SBOMFormat:  SBOMFormatSPDXJSON,
		SourceRef:   "artifact://sbom/spdx.json",
		SBOM:        spdxPayload,
	})
	if err != nil {
		t.Fatalf("IngestSBOM() error = %v", err)
	}
	if !result.DocumentStored || result.ComponentsIngested != 2 {
		t.Fatalf("unexpected ingest result %#v", result)
	}

	image, err := store.GetSBOMImage(ctx, "sha256:sbom1", 100)
	if err != nil {
		t.Fatalf("GetSBOMImage() error = %v", err)
	}
	if image.ComponentCount != 2 || image.Document.SBOMFormat != SBOMFormatSPDXJSON {
		t.Fatalf("unexpected sbom image response %#v", image)
	}

	components, err := store.SearchSBOMComponents(ctx, SBOMComponentSearchFilter{ComponentName: "openssl", Limit: 10})
	if err != nil {
		t.Fatalf("SearchSBOMComponents() error = %v", err)
	}
	if len(components) != 1 || components[0].ComponentName != "openssl" || !strings.Contains(components[0].PURL, "openssl") {
		t.Fatalf("unexpected component search %#v", components)
	}
}

func TestMemoryStoreIngestCycloneDXAndRejectMalformedSBOM(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()

	cdxPayload := []byte(`{
	  "metadata": {
	    "component": {
	      "type": "container",
	      "name": "example-image",
	      "version": "1.0.0"
	    }
	  },
	  "components": [
	    {
	      "type": "library",
	      "name": "log4j-core",
	      "version": "2.14.1",
	      "purl": "pkg:maven/org.apache.logging.log4j/log4j-core@2.14.1",
	      "licenses": [{"license":{"id":"Apache-2.0"}}]
	    }
	  ]
	}`)
	result, err := store.IngestSBOM(ctx, SBOMIngestRequest{
		ImageDigest: "sha256:cdx1",
		ImageRef:    "ghcr.io/example/app:2.0.0",
		SBOMFormat:  SBOMFormatCycloneDXJSON,
		SBOM:        cdxPayload,
	})
	if err != nil {
		t.Fatalf("IngestSBOM() cyclonedx error = %v", err)
	}
	if result.ComponentsIngested != 2 {
		t.Fatalf("expected metadata component plus library, got %#v", result)
	}

	if _, err := store.IngestSBOM(ctx, SBOMIngestRequest{
		ImageDigest: "sha256:bad1",
		SBOMFormat:  SBOMFormatSPDXJSON,
		SBOM:        []byte(`{"packages":[`),
	}); err == nil {
		t.Fatal("expected malformed sbom error")
	}
}

func TestMemoryStoreVulnerabilityScanLifecycleAndDecisionSuppression(t *testing.T) {
	store := NewMemoryStore()
	now := time.Date(2026, 4, 16, 11, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return now }
	ctx := context.Background()

	if _, err := store.Ingest(ctx, Event{
		Component:   "runtime-agent",
		EventType:   EventTypeRuntimeDriftResult,
		Decision:    DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
		Workload:    "checkout",
		Repo:        "my-org/checkout",
		Image:       "ghcr.io/example/checkout:1.0.0",
		Digest:      "sha256:digest1",
	}); err != nil {
		t.Fatalf("Ingest() runtime event error = %v", err)
	}

	firstRun, err := store.RecordVulnerabilityScan(ctx, VulnerabilityScanRequest{
		ImageDigest: "sha256:digest1",
		ImageRef:    "ghcr.io/example/checkout:1.0.0",
		Scanner:     "trivy",
		ScanMode:    VulnerabilityScanModeManual,
		StartedAt:   now,
		CompletedAt: ptrTime(now.Add(2 * time.Minute)),
		Status:      VulnerabilityScanStatusCompleted,
		Findings:    nil,
	})
	if err != nil {
		t.Fatalf("RecordVulnerabilityScan() initial error = %v", err)
	}
	if firstRun.HadPriorSuccessfulRun {
		t.Fatalf("expected first run not to have prior successful run")
	}

	secondRun, err := store.RecordVulnerabilityScan(ctx, VulnerabilityScanRequest{
		ImageDigest: "sha256:digest1",
		ImageRef:    "ghcr.io/example/checkout:1.0.0",
		Scanner:     "trivy",
		ScanMode:    VulnerabilityScanModePeriodic,
		StartedAt:   now.Add(10 * time.Minute),
		CompletedAt: ptrTime(now.Add(12 * time.Minute)),
		Status:      VulnerabilityScanStatusCompleted,
		Findings: []VulnerabilityFindingInput{
			{
				CVEID:          "CVE-2026-9999",
				Severity:       "HIGH",
				PackageName:    "openssl",
				PackageVersion: "3.0.14-r0",
				PURL:           "pkg:apk/alpine/openssl@3.0.14-r0",
				Title:          "openssl issue",
				Description:    "new issue",
				Source:         "trivy-db",
			},
		},
	})
	if err != nil {
		t.Fatalf("RecordVulnerabilityScan() second error = %v", err)
	}
	if !secondRun.HadPriorSuccessfulRun || len(secondRun.NewFindings) != 1 {
		t.Fatalf("expected drift-like new finding, got %#v", secondRun)
	}

	active, err := store.ListActiveVulnerabilities(ctx, VulnerabilityActiveFilter{ImageDigest: "sha256:digest1", Limit: 10})
	if err != nil {
		t.Fatalf("ListActiveVulnerabilities() error = %v", err)
	}
	if len(active) != 1 || active[0].CVEID != "CVE-2026-9999" {
		t.Fatalf("unexpected active findings %#v", active)
	}

	blastRadius, err := store.VulnerabilityBlastRadius(ctx, VulnerabilityBlastRadiusFilter{CVEID: "CVE-2026-9999", Limit: 10})
	if err != nil {
		t.Fatalf("VulnerabilityBlastRadius() error = %v", err)
	}
	if len(blastRadius.Items) != 1 || len(blastRadius.Items[0].Workloads) != 1 || blastRadius.Items[0].Workloads[0].Workload != "checkout" {
		t.Fatalf("unexpected blast radius %#v", blastRadius)
	}

	timeline, err := store.VulnerabilityTimeline(ctx, VulnerabilityTimelineFilter{
		ImageDigest: "sha256:digest1",
		CVEID:       "CVE-2026-9999",
		WindowDays:  30,
	})
	if err != nil {
		t.Fatalf("VulnerabilityTimeline() error = %v", err)
	}
	if len(timeline.Items) != 1 || timeline.Items[0].Status != VulnerabilityFindingStatusOpen {
		t.Fatalf("unexpected timeline %#v", timeline)
	}

	decision, err := store.CreateVulnerabilityDecision(ctx, VulnerabilityDecisionCreateRequest{
		ImageDigest:   "sha256:digest1",
		CVEID:         "CVE-2026-9999",
		Decision:      VulnerabilityDecisionNotAffected,
		Justification: "not reachable in this image path",
	}, "security@example.com")
	if err != nil {
		t.Fatalf("CreateVulnerabilityDecision() error = %v", err)
	}

	activeSuppressed, err := store.ListActiveVulnerabilities(ctx, VulnerabilityActiveFilter{ImageDigest: "sha256:digest1", Limit: 10})
	if err != nil {
		t.Fatalf("ListActiveVulnerabilities() suppressed error = %v", err)
	}
	if len(activeSuppressed) != 0 {
		t.Fatalf("expected NOT_AFFECTED decision to suppress default listing, got %#v", activeSuppressed)
	}

	activeWithSuppressed, err := store.ListActiveVulnerabilities(ctx, VulnerabilityActiveFilter{
		ImageDigest:       "sha256:digest1",
		IncludeSuppressed: true,
		Limit:             10,
	})
	if err != nil {
		t.Fatalf("ListActiveVulnerabilities() include suppressed error = %v", err)
	}
	if len(activeWithSuppressed) != 1 || activeWithSuppressed[0].Decision == nil || activeWithSuppressed[0].Decision.ID != decision.ID {
		t.Fatalf("expected decision on suppressed listing, got %#v", activeWithSuppressed)
	}

	if _, err := store.DeactivateVulnerabilityDecision(ctx, decision.ID); err != nil {
		t.Fatalf("DeactivateVulnerabilityDecision() error = %v", err)
	}
	activeAgain, err := store.ListActiveVulnerabilities(ctx, VulnerabilityActiveFilter{ImageDigest: "sha256:digest1", Limit: 10})
	if err != nil {
		t.Fatalf("ListActiveVulnerabilities() after deactivate error = %v", err)
	}
	if len(activeAgain) != 1 {
		t.Fatalf("expected active finding after decision deactivation, got %#v", activeAgain)
	}
}

func ptrTime(value time.Time) *time.Time {
	return &value
}
