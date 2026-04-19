package signing

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type VaultTransitProvider struct {
	addr        string
	token       string
	transitPath string
	key         string
	keyID       string
	algorithm   string
	client      *http.Client
	now         func() time.Time
}

func NewVaultTransitProvider(config Config, options ProviderOptions) (*VaultTransitProvider, error) {
	if config.Mode != ModeVaultTransit {
		return nil, fmt.Errorf("vault transit provider requires %s mode", ModeVaultTransit)
	}
	if config.VaultAddr == "" || config.VaultToken == "" || config.VaultTransitKey == "" {
		return nil, errors.New("vault transit signer configuration is incomplete")
	}
	client := options.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: 5 * time.Second}
	}
	now := options.Now
	if now == nil {
		now = time.Now
	}
	return &VaultTransitProvider{
		addr:        strings.TrimRight(config.VaultAddr, "/"),
		token:       config.VaultToken,
		transitPath: strings.Trim(config.VaultTransitPath, "/"),
		key:         config.VaultTransitKey,
		keyID:       config.KeyID,
		algorithm:   config.Algorithm,
		client:      client,
		now:         now,
	}, nil
}

func (p *VaultTransitProvider) Mode() string {
	return ModeVaultTransit
}

func (p *VaultTransitProvider) Sign(ctx context.Context, purpose string, payload []byte) (Envelope, error) {
	if strings.TrimSpace(purpose) == "" {
		return Envelope{}, errors.New("signing purpose is required")
	}

	var response struct {
		Data struct {
			Signature string `json:"signature"`
		} `json:"data"`
		Errors []string `json:"errors"`
	}
	if err := p.doJSON(ctx, http.MethodPost, p.signURL(), map[string]any{
		"input":          base64.StdEncoding.EncodeToString(payload),
		"hash_algorithm": p.algorithm,
	}, &response); err != nil {
		return Envelope{}, err
	}
	if strings.TrimSpace(response.Data.Signature) == "" {
		return Envelope{}, errors.New("vault transit response did not include a signature")
	}

	digest := digestPayload(payload)
	return Envelope{
		Provider:      p.Mode(),
		KeyID:         p.keyID,
		Algorithm:     p.algorithm,
		Purpose:       strings.TrimSpace(purpose),
		PayloadDigest: digest,
		Signature:     response.Data.Signature,
		SignedAt:      p.now().UTC(),
	}, nil
}

func (p *VaultTransitProvider) Verify(ctx context.Context, purpose string, payload []byte, envelope Envelope) (VerificationResult, error) {
	if envelope.Provider != p.Mode() {
		return VerificationResult{State: StateFailed, Reason: "signature provider does not match configured signer"}, nil
	}
	if envelope.Algorithm != p.algorithm {
		return VerificationResult{State: StateFailed, Reason: "signature algorithm does not match configured signer"}, nil
	}
	if strings.TrimSpace(purpose) == "" || envelope.Purpose != strings.TrimSpace(purpose) {
		return VerificationResult{State: StateFailed, Reason: "signature purpose does not match verification purpose"}, nil
	}
	if envelope.PayloadDigest != digestPayload(payload) {
		return VerificationResult{State: StateFailed, Reason: "payload digest does not match signature envelope"}, nil
	}

	var response struct {
		Data struct {
			Valid bool `json:"valid"`
		} `json:"data"`
		Errors []string `json:"errors"`
	}
	if err := p.doJSON(ctx, http.MethodPost, p.verifyURL(), map[string]any{
		"input":          base64.StdEncoding.EncodeToString(payload),
		"signature":      envelope.Signature,
		"hash_algorithm": p.algorithm,
	}, &response); err != nil {
		return VerificationResult{}, err
	}
	if !response.Data.Valid {
		return VerificationResult{State: StateFailed, Reason: "vault transit signature verification failed"}, nil
	}
	return VerificationResult{State: StateVerified}, nil
}

func (p *VaultTransitProvider) signURL() string {
	return fmt.Sprintf("%s/v1/%s/sign/%s", p.addr, p.transitPath, p.key)
}

func (p *VaultTransitProvider) verifyURL() string {
	return fmt.Sprintf("%s/v1/%s/verify/%s", p.addr, p.transitPath, p.key)
}

func (p *VaultTransitProvider) doJSON(ctx context.Context, method, endpoint string, requestBody any, out any) error {
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	request, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Vault-Token", p.token)

	response, err := p.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		if len(body) == 0 {
			return fmt.Errorf("vault transit returned status %d", response.StatusCode)
		}
		return fmt.Errorf("vault transit returned status %d: %s", response.StatusCode, strings.TrimSpace(string(body)))
	}
	if err := json.Unmarshal(body, out); err != nil {
		return err
	}
	return nil
}
