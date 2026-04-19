package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/signing"
)

func newTestSoftwareSigningRuntime(t *testing.T, secret string) *signingRuntime {
	t.Helper()

	runtime, err := signing.NewRuntime(signing.Config{
		Mode:             signing.ModeSoftware,
		Purposes:         map[string]struct{}{signing.PurposeExceptions: {}, signing.PurposeSyncSnapshots: {}},
		KeyID:            "test-signing-key",
		Algorithm:        signing.AlgorithmHMACSHA256,
		VerifyOnRead:     true,
		SoftwareSecret:   secret,
		VaultTransitPath: "transit",
	}, signing.ProviderOptions{
		Now: func() time.Time {
			return time.Date(2026, 4, 16, 10, 0, 0, 0, time.UTC)
		},
	})
	if err != nil {
		t.Fatalf("NewRuntime() error = %v", err)
	}
	return &signingRuntime{runtime: runtime}
}

func TestCreateApprovedExceptionSignsEvidenceWhenEnabled(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandlerWithRuntimesAndSigning(
		store,
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "secret-a"),
	)

	req := httptest.NewRequest(http.MethodPost, "/v1/exceptions", bytes.NewBufferString(`{
	  "exception_id":"EX-SIGNED-001",
	  "exception_type":"BREAK_GLASS",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod",
	  "reason":"signed evidence test",
	  "ticket_id":"INC-SIGNED-1",
	  "ttl_hours":1
	}`))
	req.Header.Set("Authorization", "Bearer security-admin-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", rec.Code, rec.Body.String())
	}

	var response exceptionResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if response.Exception.Signature == nil || response.Exception.Signature.Provider != signing.ModeSoftware {
		t.Fatalf("expected signed exception response, got %#v", response.Exception)
	}

	stored, err := store.GetException(context.Background(), "EX-SIGNED-001")
	if err != nil {
		t.Fatalf("GetException() error = %v", err)
	}
	if stored.Signature == nil || stored.Signature.Purpose != signing.PurposeExceptions {
		t.Fatalf("expected stored signature, got %#v", stored)
	}
}

func TestValidateExceptionFailsForTamperedSignedEvidence(t *testing.T) {
	store := audit.NewMemoryStore()
	exception, err := store.CreateException(context.Background(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-TAMPERED-001",
		ExceptionType: audit.ExceptionTypeDigestBypass,
		ImageDigest:   "sha256:abc123",
		Reason:        "tamper test",
		TicketID:      "INC-TAMPER-1",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	})
	if err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}

	signRuntime := newTestSoftwareSigningRuntime(t, "secret-a")
	envelope, err := signRuntime.signException(context.Background(), exception)
	if err != nil {
		t.Fatalf("signException() error = %v", err)
	}
	envelope.Signature = envelope.Signature + "-tampered"
	if _, err := store.SetExceptionSignature(context.Background(), exception.ExceptionID, envelope); err != nil {
		t.Fatalf("SetExceptionSignature() error = %v", err)
	}

	handler := newHandlerWithRuntimesAndSigning(
		store,
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		signRuntime,
	)
	req := httptest.NewRequest(http.MethodPost, "/v1/exceptions/validate", bytes.NewBufferString(`{
	  "exception_id":"EX-TAMPERED-001",
	  "image_digest":"sha256:abc123"
	}`))
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var result audit.ExceptionValidationResult
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatalf("decode result: %v", err)
	}
	if result.Valid || result.VerificationState != signing.StateFailed || result.Exception == nil {
		t.Fatalf("expected failed verification, got %#v", result)
	}
}

func TestSyncExceptionsEndpointSignsSnapshotWhenEnabled(t *testing.T) {
	store := audit.NewMemoryStore()
	if _, err := store.CreateException(context.Background(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-SNAPSHOT-001",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Environment:   "prod",
		Reason:        "snapshot sign test",
		TicketID:      "INC-SNAPSHOT-1",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}

	syncRuntime := newSyncRuntime(syncConfig{
		Mode:             audit.SyncModeHub,
		RequireClusterID: true,
		ClusterBindings: map[string]clusterBinding{
			"service-internal-demo": {Clusters: []string{"cluster-a"}},
		},
	})
	handler := newHandlerWithRuntimesAndSigning(
		store,
		"memory",
		mustStaticAuthConfig(t),
		nil,
		syncRuntime,
		newTestSoftwareSigningRuntime(t, "secret-a"),
	)

	req := httptest.NewRequest(http.MethodGet, "/v1/sync/exceptions", nil)
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	req.Header.Set(syncClusterHeader, "cluster-a")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var snapshot audit.ExceptionSyncSnapshot
	if err := json.NewDecoder(rec.Body).Decode(&snapshot); err != nil {
		t.Fatalf("decode snapshot: %v", err)
	}
	if snapshot.Signature == nil || snapshot.Signature.Purpose != signing.PurposeSyncSnapshots {
		t.Fatalf("expected signed snapshot, got %#v", snapshot)
	}
}

func TestSpokeSyncRejectsSnapshotWhenSignatureVerificationFails(t *testing.T) {
	signerA := newTestSoftwareSigningRuntime(t, "secret-a")
	signerB := newTestSoftwareSigningRuntime(t, "secret-b")

	expiresAt := time.Now().UTC().Add(time.Hour)
	snapshot := audit.ExceptionSyncSnapshot{
		ClusterID:   "cluster-a",
		GeneratedAt: time.Now().UTC(),
		Exceptions: []audit.SyncedException{{
			ExceptionID:   "EX-SYNC-001",
			ExceptionType: audit.ExceptionTypeBreakGlass,
			TenantID:      "acme",
			Environment:   "prod",
			Namespace:     "acme-prod",
			Reason:        "signed sync test",
			TicketID:      "INC-SYNC-1",
			ApprovedBy:    "security@example.com",
			CreatedAt:     time.Now().UTC(),
			ExpiresAt:     expiresAt,
		}},
	}
	snapshot.Revision = audit.ComputeExceptionSyncRevision(snapshot.Exceptions)
	envelope, err := signerA.signSyncSnapshot(context.Background(), snapshot)
	if err != nil {
		t.Fatalf("signSyncSnapshot() error = %v", err)
	}
	snapshot.Signature = envelope

	hub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(snapshot)
	}))
	defer hub.Close()

	syncRuntime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "cluster-a",
		HubURL:       hub.URL,
		Token:        "service-internal-demo-token",
		PollInterval: time.Minute,
		FailMode:     audit.SyncFailModeLastKnownGood,
		CacheDir:     t.TempDir(),
	})
	syncRuntime.signing = signerB
	store := audit.NewMemoryStore()
	syncRuntime.syncOnce(context.Background(), store)

	status := syncRuntime.statusSnapshot()
	if status.Health != audit.SyncHealthError || status.VerificationState != signing.StateFailed {
		t.Fatalf("expected verification failure status, got %#v", status)
	}
}
