package vex

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestNormalizeCreateRequestValidatesCanonicalStatement(t *testing.T) {
	future := time.Date(2026, 4, 20, 12, 0, 0, 0, time.UTC)
	now := func() time.Time {
		return time.Date(2026, 4, 16, 12, 0, 0, 0, time.UTC)
	}

	request, err := NormalizeCreateRequest(CreateRequest{
		SourceFormat:    SourceFormatAPI,
		VulnerabilityID: " cve-2026-1234 ",
		Scope: Scope{
			ImageDigest: " sha256:abc ",
			PackageName: " openssl ",
			TenantID:    " acme ",
		},
		Status:        " Not_Affected ",
		Justification: " component is unreachable ",
		ExpiresAt:     &future,
		Metadata:      json.RawMessage(`{"legacy_decision":"NOT_AFFECTED"}`),
	}, now)
	if err != nil {
		t.Fatalf("NormalizeCreateRequest() error = %v", err)
	}

	if request.VulnerabilityID != "CVE-2026-1234" {
		t.Fatalf("unexpected vulnerability id %q", request.VulnerabilityID)
	}
	if request.Scope.ImageDigest != "sha256:abc" || request.Scope.PackageName != "openssl" || request.Scope.TenantID != "acme" {
		t.Fatalf("unexpected normalized scope %#v", request.Scope)
	}
	if request.Status != StatusNotAffected {
		t.Fatalf("unexpected status %q", request.Status)
	}
	if string(request.Metadata) != `{"legacy_decision":"NOT_AFFECTED"}` {
		t.Fatalf("unexpected metadata %s", string(request.Metadata))
	}
}

func TestNormalizeCreateRequestRejectsUnsupportedOrUnsafeStatements(t *testing.T) {
	now := func() time.Time {
		return time.Date(2026, 4, 16, 12, 0, 0, 0, time.UTC)
	}
	past := time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		request CreateRequest
	}{
		{
			name: "missing match scope",
			request: CreateRequest{
				SourceFormat:    SourceFormatAPI,
				VulnerabilityID: "CVE-2026-1000",
				Status:          StatusNotAffected,
			},
		},
		{
			name: "unsupported status",
			request: CreateRequest{
				SourceFormat:    SourceFormatAPI,
				VulnerabilityID: "CVE-2026-1000",
				Scope:           Scope{ImageDigest: "sha256:abc"},
				Status:          "suppressed",
			},
		},
		{
			name: "past expiry",
			request: CreateRequest{
				SourceFormat:    SourceFormatAPI,
				VulnerabilityID: "CVE-2026-1000",
				Scope:           Scope{ImageDigest: "sha256:abc"},
				Status:          StatusNotAffected,
				ExpiresAt:       &past,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NormalizeCreateRequest(tt.request, now); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestParseIngestRequestParsesCSAFSubset(t *testing.T) {
	payload := json.RawMessage(`{
		"document":{"title":"example"},
		"product_tree":{
			"full_product_names":[
				{
					"product_id":"prod-a",
					"name":"ghcr.io/example/api@sha256:abc123",
					"product_identification_helper":{"purl":"pkg:oci/example/api@sha256:abc123"}
				}
			]
		},
		"vulnerabilities":[
			{
				"cve":"CVE-2026-1111",
				"notes":[
					{"category":"description","text":"component is not reachable in the request path"}
				],
				"remediations":[{"details":"monitor next upstream release"}],
				"product_status":{"known_not_affected":["prod-a"]}
			}
		]
	}`)

	statements, format, err := ParseIngestRequest(IngestRequest{
		Payload: payload,
		Scope: Scope{
			TenantID: "acme",
		},
	})
	if err != nil {
		t.Fatalf("ParseIngestRequest() error = %v", err)
	}
	if format != SourceFormatCSAF {
		t.Fatalf("unexpected format %q", format)
	}
	if len(statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(statements))
	}
	if statements[0].Status != StatusNotAffected {
		t.Fatalf("unexpected status %q", statements[0].Status)
	}
	if statements[0].Scope.ImageDigest != "sha256:abc123" {
		t.Fatalf("unexpected digest scope %#v", statements[0].Scope)
	}
	if statements[0].Scope.TenantID != "acme" {
		t.Fatalf("expected inherited tenant scope, got %#v", statements[0].Scope)
	}
	if !strings.Contains(statements[0].ImpactStatement, "not reachable") {
		t.Fatalf("unexpected impact statement %#v", statements[0])
	}
}

func TestParseIngestRequestParsesCycloneDXSubset(t *testing.T) {
	payload := json.RawMessage(`{
		"bomFormat":"CycloneDX",
		"metadata":{
			"component":{
				"bom-ref":"ghcr.io/example/api@sha256:def456",
				"name":"openssl",
				"purl":"pkg:apk/alpine/openssl@3.0.14-r0"
			}
		},
		"vulnerabilities":[
			{
				"id":"CVE-2026-2222",
				"analysis":{
					"state":"in_triage",
					"justification":"requires runtime reachability review",
					"detail":"shared object is present but usage is unclear"
				},
				"affects":[{"ref":"ghcr.io/example/api@sha256:def456"}],
				"recommendation":"inspect runtime code path"
			}
		]
	}`)

	statements, format, err := ParseIngestRequest(IngestRequest{Payload: payload})
	if err != nil {
		t.Fatalf("ParseIngestRequest() error = %v", err)
	}
	if format != SourceFormatCycloneDX {
		t.Fatalf("unexpected format %q", format)
	}
	if len(statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(statements))
	}
	if statements[0].Status != StatusUnderInvestigation {
		t.Fatalf("unexpected status %q", statements[0].Status)
	}
	if statements[0].Scope.ImageDigest != "sha256:def456" {
		t.Fatalf("unexpected scope %#v", statements[0].Scope)
	}
	if statements[0].Scope.PURL != "pkg:apk/alpine/openssl@3.0.14-r0" {
		t.Fatalf("unexpected purl scope %#v", statements[0].Scope)
	}
}

func TestMatchesUsesDigestAndScopeNarrowing(t *testing.T) {
	now := time.Date(2026, 4, 16, 12, 0, 0, 0, time.UTC)
	statement := Statement{
		Active:          true,
		VulnerabilityID: "CVE-2026-3333",
		Scope: Scope{
			ImageDigest: "sha256:abc",
			PackageName: "openssl",
			TenantID:    "acme",
			ClusterID:   "prod-eu",
		},
		Status: StatusNotAffected,
	}
	finding := FindingRef{
		VulnerabilityID: "CVE-2026-3333",
		ImageDigest:     "sha256:abc",
		PackageName:     "OpenSSL",
	}
	scope := EvaluationScope{
		TenantID:  "acme",
		ClusterID: "prod-eu",
	}

	if !Matches(statement, finding, scope, now) {
		t.Fatal("expected statement to match scoped finding")
	}
	if Matches(statement, finding, EvaluationScope{TenantID: "globex", ClusterID: "prod-eu"}, now) {
		t.Fatal("expected tenant mismatch to prevent match")
	}
	if Matches(statement, FindingRef{
		VulnerabilityID: "CVE-2026-3333",
		ImageDigest:     "sha256:other",
		PackageName:     "openssl",
	}, scope, now) {
		t.Fatal("expected digest mismatch to prevent match")
	}
}

func TestMatchesRejectsExpiredOrRevokedStatements(t *testing.T) {
	now := time.Date(2026, 4, 16, 12, 0, 0, 0, time.UTC)
	expiredAt := now.Add(-time.Hour)
	revokedAt := now.Add(-30 * time.Minute)
	finding := FindingRef{
		VulnerabilityID: "CVE-2026-4444",
		ImageDigest:     "sha256:abc",
	}
	scope := EvaluationScope{}

	if Matches(Statement{
		Active:          true,
		VulnerabilityID: "CVE-2026-4444",
		Scope:           Scope{ImageDigest: "sha256:abc"},
		Status:          StatusNotAffected,
		ExpiresAt:       &expiredAt,
	}, finding, scope, now) {
		t.Fatal("expected expired statement not to match")
	}
	if Matches(Statement{
		Active:          true,
		VulnerabilityID: "CVE-2026-4444",
		Scope:           Scope{ImageDigest: "sha256:abc"},
		Status:          StatusNotAffected,
		RevokedAt:       &revokedAt,
	}, finding, scope, now) {
		t.Fatal("expected revoked statement not to match")
	}
}

func TestSuppressesFindingOnlyForNotAffected(t *testing.T) {
	if !SuppressesFinding(Statement{Status: StatusNotAffected}) {
		t.Fatal("expected not_affected to suppress finding")
	}
	if SuppressesFinding(Statement{Status: StatusUnderInvestigation}) {
		t.Fatal("expected under_investigation not to suppress finding")
	}
	if SuppressesFinding(Statement{Status: StatusFixed}) {
		t.Fatal("expected fixed not to suppress finding")
	}
}
