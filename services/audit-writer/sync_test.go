package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

func mustStaticAuthConfig(t *testing.T) auth.Config {
	t.Helper()

	cfg, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	return cfg
}

func TestLoadSyncConfigFromEnvRejectsInvalidConfig(t *testing.T) {
	t.Run("invalid mode", func(t *testing.T) {
		_, err := loadSyncConfigFromEnv(func(key string) string {
			if key == "CHANGELOCK_SYNC_MODE" {
				return "bogus"
			}
			return ""
		})
		if err == nil || !strings.Contains(err.Error(), "unsupported CHANGELOCK_SYNC_MODE") {
			t.Fatalf("expected invalid mode error, got %v", err)
		}
	})

	t.Run("spoke requires cluster id", func(t *testing.T) {
		_, err := loadSyncConfigFromEnv(func(key string) string {
			switch key {
			case "CHANGELOCK_SYNC_MODE":
				return audit.SyncModeSpoke
			case "CHANGELOCK_SYNC_HUB_URL":
				return "https://hub.example.com"
			case "CHANGELOCK_SYNC_TOKEN":
				return "secret"
			default:
				return ""
			}
		})
		if err == nil || !strings.Contains(err.Error(), "CHANGELOCK_CLUSTER_ID is required") {
			t.Fatalf("expected missing cluster id error, got %v", err)
		}
	})

	t.Run("hub requires bindings when cluster identity is enforced", func(t *testing.T) {
		_, err := loadSyncConfigFromEnv(func(key string) string {
			switch key {
			case "CHANGELOCK_SYNC_MODE":
				return audit.SyncModeHub
			case "CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID":
				return "true"
			default:
				return ""
			}
		})
		if err == nil || !strings.Contains(err.Error(), "CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON is required") {
			t.Fatalf("expected missing bindings error, got %v", err)
		}
	})
}

func TestSyncExceptionsEndpointRejectsUnauthorizedClusterRequest(t *testing.T) {
	store := audit.NewMemoryStore()
	syncRuntime := newSyncRuntime(syncConfig{
		Mode:             audit.SyncModeHub,
		RequireClusterID: true,
		ClusterBindings: map[string]clusterBinding{
			"service-internal-demo": {Clusters: []string{"cluster-a"}},
		},
	})
	handler := newHandlerWithRuntimes(store, "memory", mustStaticAuthConfig(t), nil, syncRuntime)

	req := httptest.NewRequest(http.MethodGet, "/v1/sync/exceptions", nil)
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	req.Header.Set(syncClusterHeader, "cluster-b")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestSyncExceptionsEndpointReturnsApprovedSnapshotAndETag(t *testing.T) {
	store := audit.NewMemoryStore()
	if _, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-ACME-APPROVED",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Environment:   "prod",
		Reason:        "approved",
		TicketID:      "INC-APPROVED",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	if _, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-GLOBEX-APPROVED",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		TenantID:      "globex",
		Environment:   "prod",
		Reason:        "approved",
		TicketID:      "INC-GLOBEX",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	if _, err := store.RequestException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-PENDING",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Environment:   "prod",
		Reason:        "pending",
		TicketID:      "INC-PENDING",
		TTLHours:      1,
	}, "operator@example.com", auth.RoleOperator); err != nil {
		t.Fatalf("RequestException() error = %v", err)
	}

	syncRuntime := newSyncRuntime(syncConfig{
		Mode:             audit.SyncModeHub,
		RequireClusterID: true,
		ClusterBindings: map[string]clusterBinding{
			"service-internal-demo": {
				Clusters: []string{"cluster-a"},
				Tenants:  []string{"acme"},
			},
		},
	})
	handler := newHandlerWithRuntimes(store, "memory", mustStaticAuthConfig(t), nil, syncRuntime)

	req := httptest.NewRequest(http.MethodGet, "/v1/sync/exceptions", nil)
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	req.Header.Set(syncClusterHeader, "cluster-a")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	etag := rec.Header().Get("ETag")
	if etag == "" {
		t.Fatal("expected ETag header")
	}

	var snapshot audit.ExceptionSyncSnapshot
	if err := json.NewDecoder(rec.Body).Decode(&snapshot); err != nil {
		t.Fatalf("decode snapshot: %v", err)
	}
	if snapshot.ClusterID != "cluster-a" {
		t.Fatalf("expected cluster-a snapshot, got %#v", snapshot)
	}
	if len(snapshot.Exceptions) != 1 || snapshot.Exceptions[0].ExceptionID != "EX-ACME-APPROVED" {
		t.Fatalf("unexpected exception snapshot %#v", snapshot.Exceptions)
	}

	notModifiedReq := httptest.NewRequest(http.MethodGet, "/v1/sync/exceptions", nil)
	notModifiedReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	notModifiedReq.Header.Set(syncClusterHeader, "cluster-a")
	notModifiedReq.Header.Set("If-None-Match", etag)
	notModifiedRec := httptest.NewRecorder()
	handler.ServeHTTP(notModifiedRec, notModifiedReq)

	if notModifiedRec.Code != http.StatusNotModified {
		t.Fatalf("expected 304, got %d: %s", notModifiedRec.Code, notModifiedRec.Body.String())
	}
}

func TestSpokeCacheReloadKeepsLastKnownGoodValidationWorking(t *testing.T) {
	cacheDir := t.TempDir()
	expiresAt := time.Now().UTC().Add(time.Hour)
	snapshot := audit.ExceptionSyncSnapshot{
		ClusterID: "cluster-a",
		Revision: audit.ComputeExceptionSyncRevision([]audit.SyncedException{{
			ExceptionID:   "EX-CACHED",
			ExceptionType: audit.ExceptionTypeBreakGlass,
			TenantID:      "acme",
			Environment:   "prod",
			Namespace:     "acme-prod",
			Reason:        "cached approval",
			TicketID:      "INC-CACHED",
			ApprovedBy:    "security@example.com",
			CreatedAt:     time.Now().UTC(),
			ExpiresAt:     expiresAt,
		}}),
		GeneratedAt: time.Now().UTC(),
		Exceptions: []audit.SyncedException{{
			ExceptionID:   "EX-CACHED",
			ExceptionType: audit.ExceptionTypeBreakGlass,
			TenantID:      "acme",
			Environment:   "prod",
			Namespace:     "acme-prod",
			Reason:        "cached approval",
			TicketID:      "INC-CACHED",
			ApprovedBy:    "security@example.com",
			CreatedAt:     time.Now().UTC(),
			ExpiresAt:     expiresAt,
		}},
	}

	syncRuntime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "cluster-a",
		HubURL:       "https://hub.example.com",
		Token:        "secret",
		PollInterval: time.Minute,
		FailMode:     audit.SyncFailModeLastKnownGood,
		CacheDir:     cacheDir,
	})
	if err := syncRuntime.writeCache(snapshot, time.Now().UTC()); err != nil {
		t.Fatalf("writeCache() error = %v", err)
	}
	if _, err := os.Stat(syncRuntime.cacheFilePath()); err != nil {
		t.Fatalf("expected cache file on disk, got %v", err)
	}

	store := audit.NewMemoryStore()
	if err := syncRuntime.loadCache(t.Context(), store); err != nil {
		t.Fatalf("loadCache() error = %v", err)
	}
	syncRuntime.markFailure(errors.New("hub unreachable"), true)

	handler := newHandlerWithRuntimes(store, "memory", mustStaticAuthConfig(t), nil, syncRuntime)

	validateReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/validate", bytes.NewBufferString(`{
	  "exception_id":"EX-CACHED",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod"
	}`))
	validateReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	validateReq.Header.Set("Content-Type", "application/json")
	validateRec := httptest.NewRecorder()
	handler.ServeHTTP(validateRec, validateReq)

	if validateRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", validateRec.Code, validateRec.Body.String())
	}

	var validation audit.ExceptionValidationResult
	if err := json.NewDecoder(validateRec.Body).Decode(&validation); err != nil {
		t.Fatalf("decode validation: %v", err)
	}
	if !validation.Valid || validation.Exception == nil || validation.Exception.ExceptionID != "EX-CACHED" {
		t.Fatalf("expected last-known-good validation, got %#v", validation)
	}

	statusReq := httptest.NewRequest(http.MethodGet, "/v1/sync/status", nil)
	statusReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	statusRec := httptest.NewRecorder()
	handler.ServeHTTP(statusRec, statusReq)
	if statusRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", statusRec.Code, statusRec.Body.String())
	}

	var status audit.SyncStatus
	if err := json.NewDecoder(statusRec.Body).Decode(&status); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if status.Health != audit.SyncHealthStale || !status.CachePresent || status.CurrentRevision == "" || status.LastSuccessfulSyncAt == nil {
		t.Fatalf("unexpected sync status %#v", status)
	}
	if status.SyncMode != audit.SyncModeSpoke || status.RevisionETag == "" || !strings.Contains(status.Summary, "last-known-good") {
		t.Fatalf("expected explicit stale status surface, got %#v", status)
	}
}

func TestSpokeDenyModeBlocksExceptionValidationWhenSyncIsUnhealthy(t *testing.T) {
	store := audit.NewMemoryStore()
	if _, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-DENY",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Environment:   "prod",
		Namespace:     "acme-prod",
		Reason:        "deny mode test",
		TicketID:      "INC-DENY",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}

	syncRuntime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "cluster-a",
		HubURL:       "https://hub.example.com",
		Token:        "secret",
		PollInterval: time.Minute,
		FailMode:     audit.SyncFailModeDeny,
		CacheDir:     t.TempDir(),
	})
	syncRuntime.markFailure(errors.New("hub unreachable"), false)

	handler := newHandlerWithRuntimes(store, "memory", mustStaticAuthConfig(t), nil, syncRuntime)
	req := httptest.NewRequest(http.MethodPost, "/v1/exceptions/validate", bytes.NewBufferString(`{
	  "exception_id":"EX-DENY",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod"
	}`))
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var validation audit.ExceptionValidationResult
	if err := json.NewDecoder(rec.Body).Decode(&validation); err != nil {
		t.Fatalf("decode validation: %v", err)
	}
	if validation.Valid || !strings.Contains(validation.Reason, "deny mode") {
		t.Fatalf("expected deny mode block, got %#v", validation)
	}
}

func TestSpokeStartupWithoutCacheAndUnreachableHubIsError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "hub unavailable", http.StatusServiceUnavailable)
	}))
	server.Close()

	syncRuntime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "cluster-a",
		HubURL:       server.URL,
		Token:        "secret",
		PollInterval: time.Minute,
		FailMode:     audit.SyncFailModeLastKnownGood,
		CacheDir:     t.TempDir(),
	})

	store := audit.NewMemoryStore()
	if err := syncRuntime.loadCache(context.Background(), store); err != nil {
		t.Fatalf("loadCache() error = %v", err)
	}
	syncRuntime.syncOnce(context.Background(), store)

	status := syncRuntime.statusSnapshot()
	if status.Health != audit.SyncHealthError || status.CachePresent || strings.TrimSpace(status.LastError) == "" {
		t.Fatalf("unexpected startup status %#v", status)
	}
}

func TestHubIngestPreservesClusterIDAndSupportsClusterFilter(t *testing.T) {
	store := audit.NewMemoryStore()
	syncRuntime := newSyncRuntime(syncConfig{
		Mode:             audit.SyncModeHub,
		RequireClusterID: true,
		ClusterBindings: map[string]clusterBinding{
			"service-internal-demo": {Clusters: []string{"cluster-a"}},
		},
	})
	handler := newHandlerWithRuntimes(store, "memory", mustStaticAuthConfig(t), nil, syncRuntime)

	ingestReq := httptest.NewRequest(http.MethodPost, "/v1/ingest", bytes.NewBufferString(`{
	  "component":"deploy-gate",
	  "event_type":"deploy_gate_decision",
	  "decision":"DENY",
	  "tenant_id":"acme",
	  "environment":"prod"
	}`))
	ingestReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	ingestReq.Header.Set(syncClusterHeader, "cluster-a")
	ingestReq.Header.Set("Content-Type", "application/json")
	ingestRec := httptest.NewRecorder()
	handler.ServeHTTP(ingestRec, ingestReq)
	if ingestRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", ingestRec.Code, ingestRec.Body.String())
	}

	eventsReq := httptest.NewRequest(http.MethodGet, "/v1/reports/events?cluster_id=cluster-a", nil)
	eventsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	eventsRec := httptest.NewRecorder()
	handler.ServeHTTP(eventsRec, eventsReq)
	if eventsRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", eventsRec.Code, eventsRec.Body.String())
	}

	var events eventsResponse
	if err := json.NewDecoder(eventsRec.Body).Decode(&events); err != nil {
		t.Fatalf("decode events: %v", err)
	}
	if len(events.Events) != 1 || events.Events[0].ClusterID != "cluster-a" {
		t.Fatalf("unexpected cluster-aware events %#v", events.Events)
	}
}

func TestHubIngestRejectsUnauthorizedClusterIdentity(t *testing.T) {
	store := audit.NewMemoryStore()
	syncRuntime := newSyncRuntime(syncConfig{
		Mode:             audit.SyncModeHub,
		RequireClusterID: true,
		ClusterBindings: map[string]clusterBinding{
			"service-internal-demo": {Clusters: []string{"cluster-a"}},
		},
	})
	handler := newHandlerWithRuntimes(store, "memory", mustStaticAuthConfig(t), nil, syncRuntime)

	req := httptest.NewRequest(http.MethodPost, "/v1/ingest", bytes.NewBufferString(`{
	  "component":"deploy-gate",
	  "event_type":"deploy_gate_decision",
	  "decision":"DENY"
	}`))
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	req.Header.Set(syncClusterHeader, "cluster-b")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestTenantAndClusterScopingRemainIntersected(t *testing.T) {
	store := audit.NewMemoryStore()
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		ClusterID:   "cluster-a",
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
	})
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		ClusterID:   "cluster-a",
		TenantID:    "globex",
		Environment: "prod",
		Repo:        "my-org/globex-app",
	})

	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithRuntimes(store, "memory", cfg, nil, newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}))

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/events?cluster_id=cluster-a", nil)
	req.Header.Set("Authorization", "Bearer "+signer.token(t, map[string]any{
		"sub":       "viewer@example.com",
		"groups":    []string{"changelock-viewers"},
		"tenant_id": "acme",
	}))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var events eventsResponse
	if err := json.NewDecoder(rec.Body).Decode(&events); err != nil {
		t.Fatalf("decode events: %v", err)
	}
	if len(events.Events) != 1 || events.Events[0].TenantID != "acme" || events.Events[0].ClusterID != "cluster-a" {
		t.Fatalf("unexpected tenant+cluster-scoped events %#v", events.Events)
	}
}

func TestSyncStatusSnapshotReflectsHealthModes(t *testing.T) {
	if status := (*syncRuntime)(nil).statusSnapshot(); status.Health != audit.SyncHealthDisabled || status.Mode != audit.SyncModeDisabled {
		t.Fatalf("expected disabled status, got %#v", status)
	}
	if status := (*syncRuntime)(nil).statusSnapshot(); status.SyncMode != audit.SyncModeDisabled || status.Summary == "" {
		t.Fatalf("expected explicit disabled sync contract, got %#v", status)
	}

	hubRuntime := newSyncRuntime(syncConfig{Mode: audit.SyncModeHub, PollInterval: time.Minute})
	if status := hubRuntime.statusSnapshot(); status.Health != audit.SyncHealthHealthy || status.SyncMode != audit.SyncModeHub || !strings.Contains(status.Summary, "hub sync endpoints") {
		t.Fatalf("expected healthy hub status, got %#v", status)
	}

	staleRuntime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "cluster-a",
		HubURL:       "https://hub.example.com",
		Token:        "secret",
		PollInterval: time.Minute,
		FailMode:     audit.SyncFailModeLastKnownGood,
	})
	staleRuntime.markFailure(errors.New("hub unavailable"), true)
	if status := staleRuntime.statusSnapshot(); status.Health != audit.SyncHealthStale || !strings.Contains(status.Summary, "last-known-good") {
		t.Fatalf("expected stale status, got %#v", status)
	}

	errorRuntime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "cluster-a",
		HubURL:       "https://hub.example.com",
		Token:        "secret",
		PollInterval: time.Minute,
		FailMode:     audit.SyncFailModeDeny,
	})
	errorRuntime.markFailure(errors.New("hub unavailable"), false)
	if status := errorRuntime.statusSnapshot(); status.Health != audit.SyncHealthError || !strings.Contains(status.Summary, "hub unavailable") {
		t.Fatalf("expected error status, got %#v", status)
	}
}

func TestSyncStatusSnapshotMarksFreshnessExpiryAsStale(t *testing.T) {
	runtime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "cluster-a",
		HubURL:       "https://hub.example.com",
		Token:        "secret",
		PollInterval: time.Second,
		FailMode:     audit.SyncFailModeLastKnownGood,
	})
	old := time.Now().UTC().Add(-5 * time.Second)
	runtime.updateStatus(func(status *audit.SyncStatus) {
		status.Health = audit.SyncHealthHealthy
		status.CachePresent = true
		status.LastSuccessfulSyncAt = &old
		status.CurrentRevision = "rev-1"
	})

	status := runtime.statusSnapshot()
	if status.Health != audit.SyncHealthStale {
		t.Fatalf("expected stale status, got %#v", status)
	}
	if !strings.Contains(status.Summary, "freshness window") {
		t.Fatalf("expected freshness summary, got %#v", status)
	}
}

func TestSyncStatusEndpointExposesExplicitHealthStates(t *testing.T) {
	tests := []struct {
		name       string
		runtime    *syncRuntime
		wantHealth string
		wantMode   string
	}{
		{
			name:       "disabled",
			runtime:    nil,
			wantHealth: audit.SyncHealthDisabled,
			wantMode:   audit.SyncModeDisabled,
		},
		{
			name:       "healthy",
			runtime:    newSyncRuntime(syncConfig{Mode: audit.SyncModeHub, PollInterval: time.Minute}),
			wantHealth: audit.SyncHealthHealthy,
			wantMode:   audit.SyncModeHub,
		},
		{
			name: "stale",
			runtime: func() *syncRuntime {
				runtime := newSyncRuntime(syncConfig{
					Mode:         audit.SyncModeSpoke,
					ClusterID:    "cluster-a",
					HubURL:       "https://hub.example.com",
					Token:        "secret",
					PollInterval: time.Minute,
					FailMode:     audit.SyncFailModeLastKnownGood,
				})
				runtime.markFailure(errors.New("hub unavailable"), true)
				return runtime
			}(),
			wantHealth: audit.SyncHealthStale,
			wantMode:   audit.SyncModeSpoke,
		},
		{
			name: "error",
			runtime: func() *syncRuntime {
				runtime := newSyncRuntime(syncConfig{
					Mode:         audit.SyncModeSpoke,
					ClusterID:    "cluster-a",
					HubURL:       "https://hub.example.com",
					Token:        "secret",
					PollInterval: time.Minute,
					FailMode:     audit.SyncFailModeDeny,
				})
				runtime.markFailure(errors.New("hub unavailable"), false)
				return runtime
			}(),
			wantHealth: audit.SyncHealthError,
			wantMode:   audit.SyncModeSpoke,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := newHandlerWithRuntimes(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t), nil, tc.runtime)
			req := httptest.NewRequest(http.MethodGet, "/v1/sync/status", nil)
			req.Header.Set("Authorization", "Bearer viewer-demo-token")
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
			}
			var status audit.SyncStatus
			if err := json.NewDecoder(rec.Body).Decode(&status); err != nil {
				t.Fatalf("decode status: %v", err)
			}
			if status.Health != tc.wantHealth || status.SyncMode != tc.wantMode || status.Mode != tc.wantMode {
				t.Fatalf("unexpected sync status %#v", status)
			}
			if strings.TrimSpace(status.Summary) == "" {
				t.Fatalf("expected status summary, got %#v", status)
			}
		})
	}
}
