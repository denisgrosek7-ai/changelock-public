package audit

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestPostgresStoreVulnerabilityOpsRoundTrip(t *testing.T) {
	dsn := strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_POSTGRES_TEST_DSN"), os.Getenv("CHANGELOCK_POSTGRES_DSN"), "postgres://changelock:changelock@127.0.0.1:5433/changelock?sslmode=disable"))
	if dsn == "" {
		t.Skip("postgres vulnerability ops test dsn is not configured")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Skipf("postgres unavailable for vulnops integration test: %v", err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		t.Skipf("postgres unavailable for vulnops integration test: %v", err)
	}

	store, err := NewPostgresStore(ctx, dsn)
	if err != nil {
		t.Fatalf("NewPostgresStore() error = %v", err)
	}
	defer store.Close()
	if err := store.Migrate(ctx); err != nil {
		t.Fatalf("Migrate() error = %v", err)
	}

	digest := "sha256:pg-vulnops-" + NewRequestID()
	if _, err := store.Ingest(ctx, Event{
		RequestID:   NewRequestID(),
		Component:   "runtime-agent",
		EventType:   EventTypeRuntimeDriftResult,
		Decision:    DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
		Workload:    "checkout",
		Repo:        "my-org/checkout",
		Image:       "ghcr.io/example/checkout:1.0.0",
		Digest:      digest,
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	if _, err := store.IngestSBOM(ctx, SBOMIngestRequest{
		ImageDigest: digest,
		ImageRef:    "ghcr.io/example/checkout:1.0.0",
		SBOMFormat:  SBOMFormatSPDXJSON,
		SBOM: []byte(`{
		  "packages": [
		    {
		      "name": "openssl",
		      "versionInfo": "3.0.14-r0",
		      "licenseConcluded": "Apache-2.0",
		      "primaryPackagePurpose": "LIBRARY",
		      "externalRefs": [{"referenceType":"purl","referenceLocator":"pkg:apk/alpine/openssl@3.0.14-r0"}]
		    }
		  ]
		}`),
	}); err != nil {
		t.Fatalf("IngestSBOM() error = %v", err)
	}

	if _, err := store.RecordVulnerabilityScan(ctx, VulnerabilityScanRequest{
		ImageDigest: digest,
		ImageRef:    "ghcr.io/example/checkout:1.0.0",
		Scanner:     "trivy",
		ScanMode:    VulnerabilityScanModeManual,
		StartedAt:   time.Now().UTC(),
		CompletedAt: ptrTime(time.Now().UTC()),
		Status:      VulnerabilityScanStatusCompleted,
		Findings: []VulnerabilityFindingInput{{
			CVEID:          "CVE-2026-8888",
			Severity:       "HIGH",
			PackageName:    "openssl",
			PackageVersion: "3.0.14-r0",
			PURL:           "pkg:apk/alpine/openssl@3.0.14-r0",
			Title:          "openssl issue",
			Description:    "integration test",
			Source:         "trivy-db",
		}},
	}); err != nil {
		t.Fatalf("RecordVulnerabilityScan() error = %v", err)
	}

	components, err := store.SearchSBOMComponents(ctx, SBOMComponentSearchFilter{ComponentName: "openssl", ImageDigest: digest, Limit: 10})
	if err != nil {
		t.Fatalf("SearchSBOMComponents() error = %v", err)
	}
	if len(components) != 1 {
		t.Fatalf("unexpected components %#v", components)
	}

	if _, err := store.CreateVulnerabilityDecision(ctx, VulnerabilityDecisionCreateRequest{
		ImageDigest:   digest,
		CVEID:         "CVE-2026-8888",
		Decision:      VulnerabilityDecisionNotAffected,
		Justification: "postgres suppression test",
	}, "security@example.com"); err != nil {
		t.Fatalf("CreateVulnerabilityDecision() error = %v", err)
	}

	active, err := store.ListActiveVulnerabilities(ctx, VulnerabilityActiveFilter{ImageDigest: digest, Limit: 10})
	if err != nil {
		t.Fatalf("ListActiveVulnerabilities() error = %v", err)
	}
	if len(active) != 0 {
		t.Fatalf("expected suppressed findings to be hidden, got %#v", active)
	}

	activeWithSuppressed, err := store.ListActiveVulnerabilities(ctx, VulnerabilityActiveFilter{ImageDigest: digest, IncludeSuppressed: true, Limit: 10})
	if err != nil {
		t.Fatalf("ListActiveVulnerabilities() include suppressed error = %v", err)
	}
	if len(activeWithSuppressed) != 1 || activeWithSuppressed[0].Decision == nil {
		t.Fatalf("expected suppressed finding with decision, got %#v", activeWithSuppressed)
	}

	timeline, err := store.VulnerabilityTimeline(ctx, VulnerabilityTimelineFilter{ImageDigest: digest, CVEID: "CVE-2026-8888", WindowDays: 30})
	if err != nil {
		t.Fatalf("VulnerabilityTimeline() error = %v", err)
	}
	if len(timeline.Items) != 1 || timeline.Items[0].Decision == nil {
		t.Fatalf("unexpected timeline %#v", timeline)
	}
}
