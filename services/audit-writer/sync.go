package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/signing"
)

const syncClusterHeader = "X-Changelock-Cluster-Id"

type exceptionSyncStore interface {
	ReplaceApprovedExceptions(ctx context.Context, exceptions []audit.SyncedException) error
}

type clusterBinding struct {
	Clusters []string `json:"clusters"`
	Tenants  []string `json:"tenants,omitempty"`
}

type syncConfig struct {
	Mode             string
	ClusterID        string
	HubURL           string
	Token            string
	PollInterval     time.Duration
	FailMode         string
	CacheDir         string
	RequireClusterID bool
	ClusterBindings  map[string]clusterBinding
}

type syncCacheFile struct {
	Snapshot             audit.ExceptionSyncSnapshot `json:"snapshot"`
	LastSuccessfulSyncAt time.Time                   `json:"last_successful_sync_at"`
}

type syncRuntime struct {
	config      syncConfig
	client      *http.Client
	forwardSink *audit.HTTPSink
	signing     *signingRuntime

	mu     sync.RWMutex
	status audit.SyncStatus
}

func loadSyncRuntimeFromEnv() (*syncRuntime, error) {
	config, err := loadSyncConfigFromEnv(os.Getenv)
	if err != nil {
		return nil, err
	}
	return newSyncRuntime(config), nil
}

func loadSyncConfigFromEnv(getenv func(string) string) (syncConfig, error) {
	if getenv == nil {
		getenv = os.Getenv
	}

	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_SYNC_MODE"), audit.SyncModeDisabled)))
	switch mode {
	case audit.SyncModeDisabled, audit.SyncModeHub, audit.SyncModeSpoke:
	default:
		return syncConfig{}, fmt.Errorf("unsupported CHANGELOCK_SYNC_MODE: %s", mode)
	}

	requireClusterID, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID"), "true"))
	if err != nil {
		return syncConfig{}, fmt.Errorf("invalid CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID: %w", err)
	}

	failMode := strings.ToLower(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_SYNC_FAIL_MODE"), audit.SyncFailModeLastKnownGood)))
	switch failMode {
	case audit.SyncFailModeLastKnownGood, audit.SyncFailModeDeny:
	default:
		return syncConfig{}, fmt.Errorf("unsupported CHANGELOCK_SYNC_FAIL_MODE: %s", failMode)
	}

	pollInterval := 30 * time.Second
	if raw := strings.TrimSpace(getenv("CHANGELOCK_SYNC_POLL_INTERVAL")); raw != "" {
		parsed, err := time.ParseDuration(raw)
		if err != nil || parsed <= 0 {
			return syncConfig{}, fmt.Errorf("invalid CHANGELOCK_SYNC_POLL_INTERVAL: %s", raw)
		}
		pollInterval = parsed
	}

	cacheDir := strings.TrimSpace(getenv("CHANGELOCK_SYNC_CACHE_DIR"))
	if cacheDir == "" {
		cacheDir = filepath.Clean(".changelock-sync")
	}

	clusterID := strings.TrimSpace(getenv("CHANGELOCK_CLUSTER_ID"))
	hubURL := strings.TrimRight(strings.TrimSpace(getenv("CHANGELOCK_SYNC_HUB_URL")), "/")
	token := strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_SYNC_TOKEN"), getenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")))

	bindings := map[string]clusterBinding{}
	rawBindings := strings.TrimSpace(getenv("CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON"))
	if rawBindings != "" {
		if err := json.Unmarshal([]byte(rawBindings), &bindings); err != nil {
			return syncConfig{}, fmt.Errorf("invalid CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON: %w", err)
		}
		for principal, binding := range bindings {
			principal = strings.TrimSpace(principal)
			if principal == "" {
				return syncConfig{}, errors.New("CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON requires non-empty principal keys")
			}
			if len(binding.Clusters) == 0 {
				return syncConfig{}, fmt.Errorf("cluster binding %q requires at least one cluster", principal)
			}
			for _, cluster := range binding.Clusters {
				if strings.TrimSpace(cluster) == "" {
					return syncConfig{}, fmt.Errorf("cluster binding %q contains empty cluster id", principal)
				}
			}
		}
	}

	config := syncConfig{
		Mode:             mode,
		ClusterID:        clusterID,
		HubURL:           hubURL,
		Token:            token,
		PollInterval:     pollInterval,
		FailMode:         failMode,
		CacheDir:         cacheDir,
		RequireClusterID: requireClusterID,
		ClusterBindings:  bindings,
	}

	switch mode {
	case audit.SyncModeDisabled:
		return config, nil
	case audit.SyncModeHub:
		if requireClusterID && len(bindings) == 0 {
			return syncConfig{}, errors.New("CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON is required when CHANGELOCK_SYNC_MODE=hub and CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID=true")
		}
	case audit.SyncModeSpoke:
		if clusterID == "" {
			return syncConfig{}, errors.New("CHANGELOCK_CLUSTER_ID is required when CHANGELOCK_SYNC_MODE=spoke")
		}
		if hubURL == "" {
			return syncConfig{}, errors.New("CHANGELOCK_SYNC_HUB_URL is required when CHANGELOCK_SYNC_MODE=spoke")
		}
		if token == "" {
			return syncConfig{}, errors.New("CHANGELOCK_SYNC_TOKEN or CHANGELOCK_INTERNAL_SERVICE_TOKEN is required when CHANGELOCK_SYNC_MODE=spoke")
		}
	}

	return config, nil
}

func parseBool(value string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "t", "yes", "y", "on":
		return true, nil
	case "0", "false", "f", "no", "n", "off":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean %q", value)
	}
}

func newSyncRuntime(config syncConfig) *syncRuntime {
	runtime := &syncRuntime{
		config: config,
		client: &http.Client{Timeout: 5 * time.Second},
		status: audit.SyncStatus{
			SyncMode:          config.Mode,
			Mode:              config.Mode,
			ClusterID:         config.ClusterID,
			HubURL:            config.HubURL,
			FailMode:          config.FailMode,
			Health:            audit.SyncHealthDisabled,
			SignerMode:        signing.ModeDisabled,
			StaleAfterSeconds: int64((2 * config.PollInterval) / time.Second),
		},
	}
	switch config.Mode {
	case audit.SyncModeHub:
		runtime.status.Health = audit.SyncHealthHealthy
	case audit.SyncModeSpoke:
		runtime.status.Health = audit.SyncHealthError
		runtime.forwardSink = audit.NewHTTPSinkWithConfig(config.HubURL, 2*time.Second, config.Token, config.ClusterID)
	}
	return runtime
}

func (s *syncRuntime) start(ctx context.Context, store audit.Store) {
	if s == nil || s.config.Mode != audit.SyncModeSpoke {
		return
	}
	replacementStore, ok := store.(exceptionSyncStore)
	if !ok {
		s.markFailure(errors.New("store does not support approved exception replacement"), false)
		return
	}
	if err := s.loadCache(ctx, replacementStore); err != nil {
		s.markFailure(err, false)
	}

	go func() {
		s.syncOnce(context.Background(), replacementStore)
		ticker := time.NewTicker(s.config.PollInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.syncOnce(context.Background(), replacementStore)
			}
		}
	}()
}

func (s *syncRuntime) statusSnapshot() audit.SyncStatus {
	if s == nil {
		status := audit.SyncStatus{
			SyncMode: audit.SyncModeDisabled,
			Mode:     audit.SyncModeDisabled,
			Health:   audit.SyncHealthDisabled,
		}
		return deriveSyncStatus(status, syncConfig{}, time.Now().UTC())
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	status := s.status
	return deriveSyncStatus(status, s.config, time.Now().UTC())
}

func (s *syncRuntime) mutationBlockedReason() string {
	if s == nil || s.config.Mode != audit.SyncModeSpoke {
		return ""
	}
	return "spoke sync mode is read-only for exceptions; manage approvals on the hub"
}

func (s *syncRuntime) exceptionValidationBlockReason() string {
	if s == nil || s.config.Mode != audit.SyncModeSpoke {
		return ""
	}
	status := s.statusSnapshot()
	switch s.config.FailMode {
	case audit.SyncFailModeDeny:
		if status.Health != audit.SyncHealthHealthy {
			return "exception sync is not healthy and deny mode is enabled"
		}
		return ""
	default:
		if status.CachePresent || status.Health == audit.SyncHealthHealthy || status.Health == audit.SyncHealthStale {
			return ""
		}
		return "exception sync is unavailable and no last-known-good snapshot is loaded"
	}
}

func (s *syncRuntime) forwardEvent(ctx context.Context, event audit.Event) error {
	if s == nil || s.config.Mode != audit.SyncModeSpoke || s.forwardSink == nil {
		return nil
	}
	if strings.TrimSpace(event.ClusterID) == "" {
		event.ClusterID = s.config.ClusterID
	}
	return s.forwardSink.Write(ctx, event)
}

func (s *syncRuntime) syncOnce(ctx context.Context, store exceptionSyncStore) {
	if s == nil {
		return
	}
	revision, _ := s.currentRevision()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.config.HubURL+"/v1/sync/exceptions", nil)
	if err != nil {
		s.markFailure(err, s.statusSnapshot().CachePresent)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set(syncClusterHeader, s.config.ClusterID)
	req.Header.Set("Authorization", "Bearer "+s.config.Token)
	if revision != "" {
		req.Header.Set("If-None-Match", fmt.Sprintf("%q", revision))
	}

	response, err := s.client.Do(req)
	if err != nil {
		s.markFailure(err, s.statusSnapshot().CachePresent)
		return
	}
	defer response.Body.Close()

	now := time.Now().UTC()
	switch response.StatusCode {
	case http.StatusNotModified:
		s.markSuccess(revision, now, s.statusSnapshot().CachePresent)
		return
	case http.StatusOK:
	default:
		s.markFailure(fmt.Errorf("exception sync request failed with status %d", response.StatusCode), s.statusSnapshot().CachePresent)
		return
	}

	var snapshot audit.ExceptionSyncSnapshot
	if err := json.NewDecoder(response.Body).Decode(&snapshot); err != nil {
		s.markFailure(err, s.statusSnapshot().CachePresent)
		return
	}
	verification, err := s.verifySnapshot(ctx, snapshot)
	if err != nil {
		s.recordVerification(signing.VerificationResult{State: signing.StateFailed, Reason: err.Error()})
		s.markFailure(err, s.statusSnapshot().CachePresent)
		return
	}
	s.recordVerification(verification)
	if s.signing != nil && s.signing.verifyOnRead(signing.PurposeSyncSnapshots) && verification.State != signing.StateVerified {
		s.markFailure(errors.New(firstNonEmpty(verification.Reason, "sync snapshot verification failed")), s.statusSnapshot().CachePresent)
		return
	}
	if snapshot.ClusterID != "" && snapshot.ClusterID != s.config.ClusterID {
		s.markFailure(fmt.Errorf("hub snapshot cluster_id %q does not match local cluster_id %q", snapshot.ClusterID, s.config.ClusterID), s.statusSnapshot().CachePresent)
		return
	}
	if snapshot.Revision == "" {
		snapshot.Revision = audit.ComputeExceptionSyncRevision(snapshot.Exceptions)
	}
	if err := store.ReplaceApprovedExceptions(ctx, snapshot.Exceptions); err != nil {
		s.markFailure(err, s.statusSnapshot().CachePresent)
		return
	}
	if err := s.writeCache(snapshot, now); err != nil {
		s.markFailure(err, true)
		return
	}
	s.markSuccess(snapshot.Revision, now, true)
}

func (s *syncRuntime) loadCache(ctx context.Context, store exceptionSyncStore) error {
	path := s.cacheFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s.updateStatus(func(status *audit.SyncStatus) {
				status.CachePresent = false
			})
			return nil
		}
		return err
	}

	var cache syncCacheFile
	if err := json.Unmarshal(data, &cache); err != nil {
		return err
	}
	if cache.Snapshot.Revision == "" {
		cache.Snapshot.Revision = audit.ComputeExceptionSyncRevision(cache.Snapshot.Exceptions)
	}
	verification, err := s.verifySnapshot(ctx, cache.Snapshot)
	if err != nil {
		s.recordVerification(signing.VerificationResult{State: signing.StateFailed, Reason: err.Error()})
		return err
	}
	s.recordVerification(verification)
	if s.signing != nil && s.signing.verifyOnRead(signing.PurposeSyncSnapshots) && verification.State != signing.StateVerified {
		return errors.New(firstNonEmpty(verification.Reason, "sync snapshot verification failed"))
	}
	if err := store.ReplaceApprovedExceptions(ctx, cache.Snapshot.Exceptions); err != nil {
		return err
	}
	s.updateStatus(func(status *audit.SyncStatus) {
		status.CachePresent = true
		status.CurrentRevision = cache.Snapshot.Revision
		if !cache.LastSuccessfulSyncAt.IsZero() {
			last := cache.LastSuccessfulSyncAt.UTC()
			status.LastSuccessfulSyncAt = &last
		}
	})
	return nil
}

func (s *syncRuntime) verifySnapshot(ctx context.Context, snapshot audit.ExceptionSyncSnapshot) (signing.VerificationResult, error) {
	if s == nil || s.signing == nil {
		return signing.VerificationResult{State: signing.StateDisabled}, nil
	}
	return s.signing.verifySyncSnapshot(ctx, snapshot)
}

func (s *syncRuntime) writeCache(snapshot audit.ExceptionSyncSnapshot, lastSuccessful time.Time) error {
	if err := os.MkdirAll(s.config.CacheDir, 0o755); err != nil {
		return err
	}
	payload, err := json.MarshalIndent(syncCacheFile{
		Snapshot:             snapshot,
		LastSuccessfulSyncAt: lastSuccessful.UTC(),
	}, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.cacheFilePath(), payload, 0o644)
}

func (s *syncRuntime) cacheFilePath() string {
	return filepath.Join(s.config.CacheDir, "approved-exceptions.json")
}

func (s *syncRuntime) currentRevision() (string, bool) {
	status := s.statusSnapshot()
	if strings.TrimSpace(status.CurrentRevision) == "" {
		return "", false
	}
	return status.CurrentRevision, true
}

func (s *syncRuntime) markSuccess(revision string, at time.Time, cachePresent bool) {
	at = at.UTC()
	s.updateStatus(func(status *audit.SyncStatus) {
		status.SyncMode = s.config.Mode
		status.Health = audit.SyncHealthHealthy
		status.CurrentRevision = revision
		status.RevisionETag = quotedETag(revision)
		status.LastSuccessfulSyncAt = &at
		status.LastAttemptAt = &at
		status.LastError = ""
		status.CachePresent = cachePresent
		status.SignerMode = signing.ModeDisabled
		if s.signing != nil {
			status.SignerMode = s.signing.mode()
		}
	})
}

func (s *syncRuntime) markFailure(err error, cachePresent bool) {
	now := time.Now().UTC()
	s.updateStatus(func(status *audit.SyncStatus) {
		status.SyncMode = s.config.Mode
		status.LastAttemptAt = &now
		status.LastError = strings.TrimSpace(err.Error())
		status.CachePresent = cachePresent
		status.SignerMode = signing.ModeDisabled
		if s.signing != nil {
			status.SignerMode = s.signing.mode()
		}
		if cachePresent && s.config.FailMode == audit.SyncFailModeLastKnownGood {
			status.Health = audit.SyncHealthStale
			return
		}
		status.Health = audit.SyncHealthError
	})
}

func (s *syncRuntime) updateStatus(update func(*audit.SyncStatus)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	update(&s.status)
}

func (s *syncRuntime) recordVerification(result signing.VerificationResult) {
	s.updateStatus(func(status *audit.SyncStatus) {
		status.SignerMode = signing.ModeDisabled
		if s.signing != nil {
			status.SignerMode = s.signing.mode()
		}
		status.VerificationState = result.State
		status.VerificationReason = result.Reason
	})
}

func deriveSyncStatus(status audit.SyncStatus, config syncConfig, now time.Time) audit.SyncStatus {
	if strings.TrimSpace(status.SyncMode) == "" {
		status.SyncMode = status.Mode
	}
	if strings.TrimSpace(status.Mode) == "" {
		status.Mode = status.SyncMode
	}
	if status.CurrentRevision != "" && status.RevisionETag == "" {
		status.RevisionETag = quotedETag(status.CurrentRevision)
	}
	if status.Mode == audit.SyncModeDisabled || status.SyncMode == audit.SyncModeDisabled {
		status.Health = audit.SyncHealthDisabled
	}
	if status.Health == audit.SyncHealthHealthy && config.Mode == audit.SyncModeSpoke && status.LastSuccessfulSyncAt != nil {
		staleAfter := staleWindow(config)
		if staleAfter > 0 && now.Sub(status.LastSuccessfulSyncAt.UTC()) > staleAfter {
			status.Health = audit.SyncHealthStale
		}
	}
	status.Summary = syncHealthSummary(status, config, now)
	return status
}

func staleWindow(config syncConfig) time.Duration {
	if config.PollInterval <= 0 {
		return 0
	}
	return 2 * config.PollInterval
}

func quotedETag(revision string) string {
	revision = strings.TrimSpace(revision)
	if revision == "" {
		return ""
	}
	return fmt.Sprintf("%q", revision)
}

func syncHealthSummary(status audit.SyncStatus, config syncConfig, now time.Time) string {
	switch status.Health {
	case audit.SyncHealthDisabled:
		return "sync is disabled"
	case audit.SyncHealthHealthy:
		if status.Mode == audit.SyncModeHub {
			return "hub sync endpoints are enabled and ready to serve approved exception snapshots"
		}
		return "spoke sync state is current and usable"
	case audit.SyncHealthStale:
		if status.LastSuccessfulSyncAt != nil && staleWindow(config) > 0 && now.Sub(status.LastSuccessfulSyncAt.UTC()) > staleWindow(config) {
			return "last-known-good cache is usable, but the sync freshness window has been exceeded"
		}
		return "hub is unavailable; continuing with last-known-good cached approved exceptions"
	case audit.SyncHealthError:
		if status.LastError != "" {
			return status.LastError
		}
		if config.Mode == audit.SyncModeSpoke && config.FailMode == audit.SyncFailModeDeny {
			return "sync is unavailable and deny mode blocks exception-based allowance"
		}
		if config.Mode == audit.SyncModeSpoke && !status.CachePresent {
			return "sync is unavailable and no last-known-good cache is loaded"
		}
		return "sync is unavailable"
	default:
		return ""
	}
}

func (s *syncRuntime) authorizeClusterPrincipal(principal auth.Principal, clusterID string) (clusterBinding, error) {
	clusterID = strings.TrimSpace(clusterID)
	if s == nil {
		return clusterBinding{}, errors.New("sync runtime is not configured")
	}
	if s.config.RequireClusterID && clusterID == "" {
		return clusterBinding{}, &auth.AccessError{Status: http.StatusForbidden, Message: "cluster identity is required"}
	}

	lookupKeys := []string{strings.TrimSpace(principal.TokenID), strings.TrimSpace(principal.Subject)}
	var binding clusterBinding
	found := false
	for _, key := range lookupKeys {
		if key == "" {
			continue
		}
		candidate, ok := s.config.ClusterBindings[key]
		if ok {
			binding = candidate
			found = true
			break
		}
	}
	if !found {
		return clusterBinding{}, &auth.AccessError{Status: http.StatusForbidden, Message: "cluster sync principal is not authorized"}
	}
	if !slices.Contains(binding.Clusters, clusterID) {
		return clusterBinding{}, &auth.AccessError{Status: http.StatusForbidden, Message: "cluster sync principal is not authorized for this cluster"}
	}
	return binding, nil
}

func filterSyncedExceptionsForBinding(exceptions []audit.PolicyException, binding clusterBinding) []audit.SyncedException {
	results := make([]audit.SyncedException, 0, len(exceptions))
	for _, exception := range exceptions {
		if len(binding.Tenants) > 0 && !slices.Contains(binding.Tenants, strings.TrimSpace(exception.TenantID)) {
			continue
		}
		results = append(results, audit.SyncedExceptionFromPolicyException(exception))
	}
	return results
}
