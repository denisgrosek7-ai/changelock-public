package evidence

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ModeDisabled      = "disabled"
	ModeBundleOnly    = "bundle-only"
	ModeRekorRequired = "rekor-required"

	StateVerified   = "verified"
	StateUnverified = "unverified"
	StateFailed     = "failed"
	StateDisabled   = "disabled"
)

type Config struct {
	Mode             string
	TLogURL          string
	RequireInclusion bool
	OfflineBundleOK  bool
	VerifyOnRead     bool
	VerifyOnDeploy   bool
}

type Bundle struct {
	ID                 string     `json:"id,omitempty"`
	Purpose            string     `json:"purpose,omitempty"`
	Provider           string     `json:"provider,omitempty"`
	LogURL             string     `json:"log_url,omitempty"`
	LogID              string     `json:"log_id,omitempty"`
	LogEntryID         string     `json:"log_entry_id,omitempty"`
	BundleRef          string     `json:"bundle_ref,omitempty"`
	BundleHash         string     `json:"bundle_hash,omitempty"`
	PayloadDigest      string     `json:"payload_digest,omitempty"`
	SignedAt           *time.Time `json:"signed_at,omitempty"`
	IntegratedTime     *time.Time `json:"integrated_time,omitempty"`
	InclusionVerified  bool       `json:"inclusion_verified,omitempty"`
	VerificationState  string     `json:"verification_state,omitempty"`
	VerificationReason string     `json:"verification_reason,omitempty"`
	LastVerifiedAt     *time.Time `json:"last_verified_at,omitempty"`
}

type TLogLookupResult struct {
	EntryID        string
	LogID          string
	BodyHash       string
	IntegratedTime *time.Time
}

type TLogClient interface {
	Lookup(ctx context.Context, logURL string, entryID string) (TLogLookupResult, error)
}

type httpTLogClient struct {
	client *http.Client
}

func LoadConfigFromEnv(getenv func(string) string) (Config, error) {
	if getenv == nil {
		return Config{}, errors.New("getenv is required")
	}

	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_EVIDENCE_MODE"), ModeDisabled)))
	switch mode {
	case ModeDisabled, ModeBundleOnly, ModeRekorRequired:
	default:
		return Config{}, fmt.Errorf("unsupported CHANGELOCK_EVIDENCE_MODE: %s", mode)
	}

	requireInclusion, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_TLOG_REQUIRE_INCLUSION"), "true"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_TLOG_REQUIRE_INCLUSION: %w", err)
	}

	offlineBundleOK, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_TLOG_OFFLINE_BUNDLE_OK"), "false"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_TLOG_OFFLINE_BUNDLE_OK: %w", err)
	}

	verifyOnRead, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_EVIDENCE_VERIFY_ON_READ"), "false"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_EVIDENCE_VERIFY_ON_READ: %w", err)
	}

	verifyOnDeploy, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_EVIDENCE_VERIFY_ON_DEPLOY"), "false"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_EVIDENCE_VERIFY_ON_DEPLOY: %w", err)
	}

	cfg := Config{
		Mode:             mode,
		TLogURL:          strings.TrimRight(strings.TrimSpace(getenv("CHANGELOCK_TLOG_URL")), "/"),
		RequireInclusion: requireInclusion,
		OfflineBundleOK:  offlineBundleOK,
		VerifyOnRead:     verifyOnRead,
		VerifyOnDeploy:   verifyOnDeploy,
	}

	if cfg.Mode == ModeRekorRequired && cfg.TLogURL == "" && !cfg.OfflineBundleOK {
		return Config{}, errors.New("CHANGELOCK_TLOG_URL is required when CHANGELOCK_EVIDENCE_MODE=rekor-required and offline bundles are not allowed")
	}

	return cfg, nil
}

func NewHTTPTLogClient(timeout time.Duration) TLogClient {
	if timeout <= 0 {
		timeout = 5 * time.Second
	}
	return &httpTLogClient{
		client: &http.Client{Timeout: timeout},
	}
}

func CloneBundle(bundle *Bundle) *Bundle {
	if bundle == nil {
		return nil
	}
	clone := *bundle
	if bundle.SignedAt != nil {
		signedAt := bundle.SignedAt.UTC()
		clone.SignedAt = &signedAt
	}
	if bundle.IntegratedTime != nil {
		integratedAt := bundle.IntegratedTime.UTC()
		clone.IntegratedTime = &integratedAt
	}
	if bundle.LastVerifiedAt != nil {
		lastVerifiedAt := bundle.LastVerifiedAt.UTC()
		clone.LastVerifiedAt = &lastVerifiedAt
	}
	return &clone
}

func ApplyVerificationResult(bundle *Bundle, state string, reason string, verifiedAt *time.Time) *Bundle {
	clone := CloneBundle(bundle)
	if clone == nil {
		return nil
	}
	clone.VerificationState = strings.TrimSpace(state)
	clone.VerificationReason = strings.TrimSpace(reason)
	if verifiedAt == nil {
		clone.LastVerifiedAt = nil
	} else {
		timestamp := verifiedAt.UTC()
		clone.LastVerifiedAt = &timestamp
	}
	return clone
}

func CanonicalDigest(value any) string {
	if value == nil {
		return ""
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(encoded)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func EvaluateBundle(ctx context.Context, cfg Config, bundle *Bundle, expectedPayloadDigest string, client TLogClient, now func() time.Time) (string, string, *time.Time) {
	if now == nil {
		now = time.Now
	}
	if strings.TrimSpace(cfg.Mode) == "" || cfg.Mode == ModeDisabled {
		return StateDisabled, "transparency evidence verification disabled", nil
	}

	expectedPayloadDigest = normalizeDigest(expectedPayloadDigest)
	if bundle == nil {
		if cfg.Mode == ModeRekorRequired {
			return StateFailed, "transparency evidence bundle is required", nil
		}
		return StateUnverified, "no transparency evidence bundle recorded", nil
	}

	bundlePayloadDigest := normalizeDigest(bundle.PayloadDigest)
	if expectedPayloadDigest != "" && bundlePayloadDigest == "" {
		return StateFailed, "evidence bundle payload_digest is required", nil
	}
	if expectedPayloadDigest != "" && bundlePayloadDigest != expectedPayloadDigest {
		return StateFailed, fmt.Sprintf("evidence bundle payload_digest %q does not match expected digest %q", bundlePayloadDigest, expectedPayloadDigest), nil
	}

	verifiedAt := timePointer(now().UTC())
	if cfg.OfflineBundleOK && bundle.InclusionVerified && strings.TrimSpace(bundle.BundleHash) != "" {
		if cfg.RequireInclusion && bundle.IntegratedTime == nil {
			return StateFailed, "transparency evidence bundle is missing integrated_time", nil
		}
		return StateVerified, "offline transparency evidence bundle accepted", verifiedAt
	}

	if bundle.LogEntryID == "" {
		if cfg.Mode == ModeRekorRequired {
			return StateFailed, "transparency log entry id is required", nil
		}
		return StateUnverified, "transparency log entry id is not recorded", nil
	}

	logURL := strings.TrimRight(firstNonEmpty(bundle.LogURL, cfg.TLogURL), "/")
	if logURL == "" {
		if cfg.Mode == ModeRekorRequired {
			return StateFailed, "transparency log URL is required", nil
		}
		return StateUnverified, "transparency log URL is not configured", nil
	}
	if client == nil {
		client = NewHTTPTLogClient(5 * time.Second)
	}

	result, err := client.Lookup(ctx, logURL, bundle.LogEntryID)
	if err != nil {
		if cfg.Mode == ModeRekorRequired {
			return StateFailed, "transparency log lookup failed: " + err.Error(), nil
		}
		return StateUnverified, "transparency log lookup unavailable: " + err.Error(), nil
	}

	if bundle.BundleHash != "" && normalizeDigest(bundle.BundleHash) != result.BodyHash {
		return StateFailed, fmt.Sprintf("transparency log body hash %q does not match evidence bundle hash %q", result.BodyHash, normalizeDigest(bundle.BundleHash)), nil
	}
	if cfg.RequireInclusion && result.IntegratedTime == nil {
		return StateFailed, "transparency log entry is missing integrated_time", nil
	}

	return StateVerified, "transparency log entry verified", verifiedAt
}

func (c *httpTLogClient) Lookup(ctx context.Context, logURL string, entryID string) (TLogLookupResult, error) {
	logURL = strings.TrimRight(strings.TrimSpace(logURL), "/")
	entryID = strings.TrimSpace(entryID)
	if logURL == "" {
		return TLogLookupResult{}, errors.New("transparency log URL is required")
	}
	if entryID == "" {
		return TLogLookupResult{}, errors.New("transparency log entry id is required")
	}

	endpoint := logURL + "/api/v1/log/entries/" + url.PathEscape(entryID)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return TLogLookupResult{}, err
	}
	request.Header.Set("Accept", "application/json")

	response, err := c.client.Do(request)
	if err != nil {
		return TLogLookupResult{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return TLogLookupResult{}, fmt.Errorf("transparency log returned status %d", response.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		return TLogLookupResult{}, err
	}

	entry, err := decodeLookupBody(body, entryID)
	if err != nil {
		return TLogLookupResult{}, err
	}

	bodyHash, err := hashLogBody(entry.Body)
	if err != nil {
		return TLogLookupResult{}, err
	}
	return TLogLookupResult{
		EntryID:        entryID,
		LogID:          entry.LogID,
		BodyHash:       bodyHash,
		IntegratedTime: entry.IntegratedTime,
	}, nil
}

type lookupPayload struct {
	Body           string
	LogID          string
	IntegratedTime *time.Time
}

func decodeLookupBody(body []byte, entryID string) (lookupPayload, error) {
	var raw map[string]any
	if err := json.Unmarshal(body, &raw); err != nil {
		return lookupPayload{}, err
	}

	candidate := raw
	if nested, ok := raw[entryID]; ok {
		record, ok := nested.(map[string]any)
		if !ok {
			return lookupPayload{}, errors.New("transparency log entry response is malformed")
		}
		candidate = record
	}

	entryBody := readString(candidate["body"])
	if entryBody == "" {
		return lookupPayload{}, errors.New("transparency log entry body is missing")
	}

	var integratedTime *time.Time
	if value := readInt64(candidate["integratedTime"]); value > 0 {
		timestamp := time.Unix(value, 0).UTC()
		integratedTime = &timestamp
	}

	return lookupPayload{
		Body:           entryBody,
		LogID:          strings.TrimSpace(readString(candidate["logID"])),
		IntegratedTime: integratedTime,
	}, nil
}

func hashLogBody(body string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(body))
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(decoded)
	return "sha256:" + hex.EncodeToString(sum[:]), nil
}

func normalizeDigest(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	if value == "" {
		return ""
	}
	if strings.HasPrefix(value, "sha256:") {
		return value
	}
	return "sha256:" + value
}

func readString(value any) string {
	typed, _ := value.(string)
	return typed
}

func readInt64(value any) int64 {
	switch typed := value.(type) {
	case float64:
		return int64(typed)
	case int64:
		return typed
	case int:
		return int64(typed)
	default:
		return 0
	}
}

func parseBool(value string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("unsupported boolean value %q", value)
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func timePointer(value time.Time) *time.Time {
	if value.IsZero() {
		return nil
	}
	timestamp := value.UTC()
	return &timestamp
}
