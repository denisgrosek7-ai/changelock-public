package signing

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	ModeDisabled     = "disabled"
	ModeSoftware     = "software"
	ModeVaultTransit = "vault-transit"

	PurposeExceptions    = "exceptions"
	PurposeSyncSnapshots = "sync-snapshots"

	StateVerified   = "verified"
	StateUnverified = "unverified"
	StateFailed     = "failed"
	StateDisabled   = "disabled"

	AlgorithmHMACSHA256 = "hmac-sha256"
	AlgorithmSHA2256    = "sha2-256"
)

var supportedPurposes = map[string]struct{}{
	PurposeExceptions:    {},
	PurposeSyncSnapshots: {},
}

type Envelope struct {
	Provider      string    `json:"provider"`
	KeyID         string    `json:"key_id,omitempty"`
	Algorithm     string    `json:"algorithm"`
	Purpose       string    `json:"purpose"`
	PayloadDigest string    `json:"payload_digest"`
	Signature     string    `json:"signature"`
	SignedAt      time.Time `json:"signed_at"`
}

type VerificationResult struct {
	State  string `json:"state"`
	Reason string `json:"reason,omitempty"`
}

type Provider interface {
	Mode() string
	Sign(ctx context.Context, purpose string, payload []byte) (Envelope, error)
	Verify(ctx context.Context, purpose string, payload []byte, envelope Envelope) (VerificationResult, error)
}

type Config struct {
	Mode         string
	Purposes     map[string]struct{}
	KeyID        string
	Algorithm    string
	VerifyOnRead bool

	SoftwareSecret string

	VaultAddr        string
	VaultToken       string
	VaultTransitPath string
	VaultTransitKey  string
}

type Runtime struct {
	Config   Config
	Provider Provider
}

type ProviderOptions struct {
	HTTPClient *http.Client
	Now        func() time.Time
}

func ParseEnvConfig(getenv func(string) string) (Config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}

	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_SIGNER_MODE"), ModeDisabled)))
	switch mode {
	case ModeDisabled, ModeSoftware, ModeVaultTransit:
	default:
		return Config{}, fmt.Errorf("unsupported CHANGELOCK_SIGNER_MODE: %s", mode)
	}

	verifyOnRead := true
	if raw := strings.TrimSpace(getenv("CHANGELOCK_SIGNER_VERIFY_ON_READ")); raw != "" {
		value, err := parseBool(raw)
		if err != nil {
			return Config{}, fmt.Errorf("invalid CHANGELOCK_SIGNER_VERIFY_ON_READ: %w", err)
		}
		verifyOnRead = value
	}

	purposes, err := parsePurposes(firstNonEmpty(getenv("CHANGELOCK_SIGNER_PURPOSES"), PurposeExceptions+","+PurposeSyncSnapshots))
	if err != nil {
		return Config{}, err
	}

	config := Config{
		Mode:             mode,
		Purposes:         purposes,
		KeyID:            strings.TrimSpace(getenv("CHANGELOCK_SIGNER_KEY_ID")),
		Algorithm:        strings.ToLower(strings.TrimSpace(getenv("CHANGELOCK_SIGNER_ALGORITHM"))),
		VerifyOnRead:     verifyOnRead,
		SoftwareSecret:   strings.TrimSpace(getenv("CHANGELOCK_SIGNER_SOFTWARE_SECRET")),
		VaultAddr:        strings.TrimRight(strings.TrimSpace(getenv("CHANGELOCK_VAULT_ADDR")), "/"),
		VaultToken:       strings.TrimSpace(getenv("CHANGELOCK_VAULT_TOKEN")),
		VaultTransitPath: strings.Trim(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_VAULT_TRANSIT_PATH"), "transit")), "/"),
		VaultTransitKey:  strings.TrimSpace(getenv("CHANGELOCK_VAULT_TRANSIT_KEY")),
	}

	switch config.Mode {
	case ModeDisabled:
		return config, nil
	case ModeSoftware:
		if config.SoftwareSecret == "" {
			return Config{}, errors.New("CHANGELOCK_SIGNER_SOFTWARE_SECRET is required when CHANGELOCK_SIGNER_MODE=software")
		}
		if config.Algorithm == "" {
			config.Algorithm = AlgorithmHMACSHA256
		}
		if config.Algorithm != AlgorithmHMACSHA256 {
			return Config{}, fmt.Errorf("unsupported CHANGELOCK_SIGNER_ALGORITHM for software mode: %s", config.Algorithm)
		}
		if config.KeyID == "" {
			config.KeyID = "software-default"
		}
	case ModeVaultTransit:
		if config.VaultAddr == "" {
			return Config{}, errors.New("CHANGELOCK_VAULT_ADDR is required when CHANGELOCK_SIGNER_MODE=vault-transit")
		}
		if config.VaultToken == "" {
			return Config{}, errors.New("CHANGELOCK_VAULT_TOKEN is required when CHANGELOCK_SIGNER_MODE=vault-transit")
		}
		if config.VaultTransitKey == "" {
			return Config{}, errors.New("CHANGELOCK_VAULT_TRANSIT_KEY is required when CHANGELOCK_SIGNER_MODE=vault-transit")
		}
		if config.Algorithm == "" {
			config.Algorithm = AlgorithmSHA2256
		}
		if config.KeyID == "" {
			config.KeyID = config.VaultTransitKey
		}
	}

	return config, nil
}

func NewRuntime(config Config, options ProviderOptions) (*Runtime, error) {
	if config.Mode == "" {
		config.Mode = ModeDisabled
	}
	if config.Mode == ModeDisabled {
		return &Runtime{Config: config}, nil
	}

	var (
		provider Provider
		err      error
	)
	switch config.Mode {
	case ModeSoftware:
		provider, err = NewSoftwareProvider(config, options)
	case ModeVaultTransit:
		provider, err = NewVaultTransitProvider(config, options)
	default:
		err = fmt.Errorf("unsupported signer mode %s", config.Mode)
	}
	if err != nil {
		return nil, err
	}
	return &Runtime{Config: config, Provider: provider}, nil
}

func (r *Runtime) Enabled() bool {
	return r != nil && r.Provider != nil && r.Config.Mode != ModeDisabled
}

func (r *Runtime) SupportsPurpose(purpose string) bool {
	if r == nil {
		return false
	}
	_, ok := r.Config.Purposes[strings.TrimSpace(purpose)]
	return ok
}

func (r *Runtime) Sign(ctx context.Context, purpose string, payload []byte) (Envelope, error) {
	if !r.Enabled() {
		return Envelope{}, errors.New("signing is disabled")
	}
	if !r.SupportsPurpose(purpose) {
		return Envelope{}, fmt.Errorf("signer does not support purpose %s", strings.TrimSpace(purpose))
	}
	return r.Provider.Sign(ctx, purpose, payload)
}

func (r *Runtime) Verify(ctx context.Context, purpose string, payload []byte, envelope Envelope) (VerificationResult, error) {
	if !r.Enabled() {
		return VerificationResult{State: StateDisabled}, nil
	}
	if !r.SupportsPurpose(purpose) {
		return VerificationResult{State: StateDisabled, Reason: "signing purpose is disabled"}, nil
	}
	return r.Provider.Verify(ctx, purpose, payload, envelope)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func parsePurposes(raw string) (map[string]struct{}, error) {
	results := map[string]struct{}{}
	for _, item := range strings.Split(raw, ",") {
		purpose := strings.TrimSpace(strings.ToLower(item))
		if purpose == "" {
			continue
		}
		if _, ok := supportedPurposes[purpose]; !ok {
			return nil, fmt.Errorf("unsupported CHANGELOCK_SIGNER_PURPOSES entry: %s", purpose)
		}
		results[purpose] = struct{}{}
	}
	if len(results) == 0 {
		return nil, errors.New("CHANGELOCK_SIGNER_PURPOSES must include at least one supported purpose")
	}
	return results, nil
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

func digestPayload(payload []byte) string {
	sum := sha256.Sum256(payload)
	return "sha256:" + hex.EncodeToString(sum[:])
}
