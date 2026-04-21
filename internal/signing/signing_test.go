package signing

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestParseEnvConfigValidation(t *testing.T) {
	t.Run("software requires secret", func(t *testing.T) {
		_, err := ParseEnvConfig(func(key string) string {
			switch key {
			case "CHANGELOCK_SIGNER_MODE":
				return ModeSoftware
			default:
				return ""
			}
		})
		if err == nil || !strings.Contains(err.Error(), "CHANGELOCK_SIGNER_SOFTWARE_SECRET") {
			t.Fatalf("expected software secret error, got %v", err)
		}
	})

	t.Run("vault transit requires addr token and key", func(t *testing.T) {
		_, err := ParseEnvConfig(func(key string) string {
			switch key {
			case "CHANGELOCK_SIGNER_MODE":
				return ModeVaultTransit
			case "CHANGELOCK_VAULT_ADDR":
				return "https://vault.example.com"
			default:
				return ""
			}
		})
		if err == nil || !strings.Contains(err.Error(), "CHANGELOCK_VAULT_TOKEN") {
			t.Fatalf("expected vault token error, got %v", err)
		}
	})

	t.Run("unsupported purpose rejected", func(t *testing.T) {
		_, err := ParseEnvConfig(func(key string) string {
			switch key {
			case "CHANGELOCK_SIGNER_MODE":
				return ModeSoftware
			case "CHANGELOCK_SIGNER_SOFTWARE_SECRET":
				return "secret"
			case "CHANGELOCK_SIGNER_PURPOSES":
				return "bogus"
			default:
				return ""
			}
		})
		if err == nil || !strings.Contains(err.Error(), "unsupported CHANGELOCK_SIGNER_PURPOSES") {
			t.Fatalf("expected unsupported purpose error, got %v", err)
		}
	})
}

func TestSoftwareProviderRoundTrip(t *testing.T) {
	runtime, err := NewRuntime(Config{
		Mode:           ModeSoftware,
		Purposes:       map[string]struct{}{PurposeExceptions: {}},
		KeyID:          "software-key",
		Algorithm:      AlgorithmHMACSHA256,
		VerifyOnRead:   true,
		SoftwareSecret: "super-secret",
	}, ProviderOptions{Now: func() time.Time { return time.Unix(1700000000, 0).UTC() }})
	if err != nil {
		t.Fatalf("NewRuntime() error = %v", err)
	}
	payload := []byte(`{"exception_id":"EX-1"}`)

	envelope, err := runtime.Provider.Sign(context.Background(), PurposeExceptions, payload)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	if envelope.Provider != ModeSoftware || envelope.Purpose != PurposeExceptions {
		t.Fatalf("unexpected envelope %#v", envelope)
	}

	result, err := runtime.Provider.Verify(context.Background(), PurposeExceptions, payload, envelope)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if result.State != StateVerified {
		t.Fatalf("expected verified, got %#v", result)
	}

	tampered, err := runtime.Provider.Verify(context.Background(), PurposeExceptions, []byte(`{"exception_id":"EX-2"}`), envelope)
	if err != nil {
		t.Fatalf("Verify() tampered error = %v", err)
	}
	if tampered.State != StateFailed {
		t.Fatalf("expected failed for tampered payload, got %#v", tampered)
	}
}

func TestVaultTransitProviderRequestAndVerify(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Vault-Token"); got != "vault-token" {
			t.Fatalf("unexpected vault token %q", got)
		}
		switch r.URL.Path {
		case "/v1/transit/sign/changelock":
			var request map[string]any
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Fatalf("decode sign request: %v", err)
			}
			if request["hash_algorithm"] != AlgorithmSHA2256 {
				t.Fatalf("unexpected sign request %#v", request)
			}
			_ = json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"signature": "vault:v1:c2lnbmVk",
				},
			})
		case "/v1/transit/verify/changelock":
			var request map[string]any
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Fatalf("decode verify request: %v", err)
			}
			if request["signature"] != "vault:v1:c2lnbmVk" {
				t.Fatalf("unexpected verify request %#v", request)
			}
			_ = json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"valid": true,
				},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	runtime, err := NewRuntime(Config{
		Mode:             ModeVaultTransit,
		Purposes:         map[string]struct{}{PurposeSyncSnapshots: {}},
		KeyID:            "transit/changelock",
		Algorithm:        AlgorithmSHA2256,
		VerifyOnRead:     true,
		VaultAddr:        server.URL,
		VaultToken:       "vault-token",
		VaultTransitPath: "transit",
		VaultTransitKey:  "changelock",
	}, ProviderOptions{HTTPClient: server.Client(), Now: func() time.Time { return time.Unix(1700000000, 0).UTC() }})
	if err != nil {
		t.Fatalf("NewRuntime() error = %v", err)
	}

	payload := []byte(`{"revision":"abc123"}`)
	envelope, err := runtime.Provider.Sign(context.Background(), PurposeSyncSnapshots, payload)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	result, err := runtime.Provider.Verify(context.Background(), PurposeSyncSnapshots, payload, envelope)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if result.State != StateVerified {
		t.Fatalf("expected verified result, got %#v", result)
	}
}

func TestNoImplicitFallbackFromVaultToSoftware(t *testing.T) {
	_, err := NewRuntime(Config{
		Mode:             ModeVaultTransit,
		Purposes:         map[string]struct{}{PurposeExceptions: {}},
		VerifyOnRead:     true,
		VaultAddr:        "https://vault.example.com",
		VaultTransitPath: "transit",
		VaultTransitKey:  "changelock",
	}, ProviderOptions{})
	if err == nil || !strings.Contains(err.Error(), "vault transit signer configuration is incomplete") {
		t.Fatalf("expected vault configuration error, got %v", err)
	}
}

func TestRuntimeDescribeProvider(t *testing.T) {
	t.Run("software", func(t *testing.T) {
		runtime, err := NewRuntime(Config{
			Mode:           ModeSoftware,
			Purposes:       map[string]struct{}{PurposeExceptions: {}},
			KeyID:          "software-key",
			Algorithm:      AlgorithmHMACSHA256,
			VerifyOnRead:   true,
			SoftwareSecret: "super-secret",
		}, ProviderOptions{})
		if err != nil {
			t.Fatalf("NewRuntime() error = %v", err)
		}
		descriptor := runtime.DescribeProvider()
		if descriptor.ProviderMode != ModeSoftware || descriptor.TrustBoundary != TrustBoundaryApplicationLocal || !descriptor.SupportsRotation {
			t.Fatalf("unexpected software provider descriptor %#v", descriptor)
		}
	})

	t.Run("vault transit", func(t *testing.T) {
		runtime, err := NewRuntime(Config{
			Mode:             ModeVaultTransit,
			Purposes:         map[string]struct{}{PurposeSyncSnapshots: {}},
			KeyID:            "transit/changelock",
			Algorithm:        AlgorithmSHA2256,
			VerifyOnRead:     true,
			VaultAddr:        "https://vault.example.com",
			VaultToken:       "vault-token",
			VaultTransitPath: "transit",
			VaultTransitKey:  "changelock",
		}, ProviderOptions{HTTPClient: http.DefaultClient})
		if err != nil {
			t.Fatalf("NewRuntime() error = %v", err)
		}
		descriptor := runtime.DescribeProvider()
		if descriptor.ProviderMode != ModeVaultTransit || descriptor.TrustBoundary != TrustBoundaryExternalManaged || descriptor.KeyExportability != "non_exportable_remote_key_material" {
			t.Fatalf("unexpected vault transit descriptor %#v", descriptor)
		}
	})

	t.Run("disabled", func(t *testing.T) {
		runtime, err := NewRuntime(Config{Mode: ModeDisabled}, ProviderOptions{})
		if err != nil {
			t.Fatalf("NewRuntime() error = %v", err)
		}
		descriptor := runtime.DescribeProvider()
		if descriptor.ProviderMode != ModeDisabled || descriptor.TrustBoundary != TrustBoundaryDisabled {
			t.Fatalf("unexpected disabled descriptor %#v", descriptor)
		}
	})
}
