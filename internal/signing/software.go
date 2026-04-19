package signing

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

type SoftwareProvider struct {
	key       []byte
	keyID     string
	algorithm string
	now       func() time.Time
}

func NewSoftwareProvider(config Config, options ProviderOptions) (*SoftwareProvider, error) {
	if config.Mode != ModeSoftware {
		return nil, fmt.Errorf("software provider requires %s mode", ModeSoftware)
	}
	if config.SoftwareSecret == "" {
		return nil, errors.New("software signer secret is required")
	}
	now := options.Now
	if now == nil {
		now = time.Now
	}
	return &SoftwareProvider{
		key:       []byte(config.SoftwareSecret),
		keyID:     config.KeyID,
		algorithm: config.Algorithm,
		now:       now,
	}, nil
}

func (p *SoftwareProvider) Mode() string {
	return ModeSoftware
}

func (p *SoftwareProvider) Sign(_ context.Context, purpose string, payload []byte) (Envelope, error) {
	if stringsTrim(purpose) == "" {
		return Envelope{}, errors.New("signing purpose is required")
	}
	digest := sha256.Sum256(payload)
	mac := hmac.New(sha256.New, p.key)
	_, _ = mac.Write(payload)
	signature := mac.Sum(nil)
	return Envelope{
		Provider:      p.Mode(),
		KeyID:         p.keyID,
		Algorithm:     p.algorithm,
		Purpose:       stringsTrim(purpose),
		PayloadDigest: "sha256:" + hex.EncodeToString(digest[:]),
		Signature:     base64.RawURLEncoding.EncodeToString(signature),
		SignedAt:      p.now().UTC(),
	}, nil
}

func (p *SoftwareProvider) Verify(_ context.Context, purpose string, payload []byte, envelope Envelope) (VerificationResult, error) {
	if stringsTrim(purpose) == "" {
		return VerificationResult{State: StateFailed, Reason: "signing purpose is required"}, nil
	}
	if envelope.Provider != p.Mode() {
		return VerificationResult{State: StateFailed, Reason: "signature provider does not match configured signer"}, nil
	}
	if envelope.Algorithm != p.algorithm {
		return VerificationResult{State: StateFailed, Reason: "signature algorithm does not match configured signer"}, nil
	}
	if envelope.Purpose != stringsTrim(purpose) {
		return VerificationResult{State: StateFailed, Reason: "signature purpose does not match verification purpose"}, nil
	}
	digest := sha256.Sum256(payload)
	expectedDigest := "sha256:" + hex.EncodeToString(digest[:])
	if envelope.PayloadDigest != expectedDigest {
		return VerificationResult{State: StateFailed, Reason: "payload digest does not match signature envelope"}, nil
	}
	signature, err := base64.RawURLEncoding.DecodeString(envelope.Signature)
	if err != nil {
		return VerificationResult{State: StateFailed, Reason: "signature encoding is invalid"}, nil
	}
	mac := hmac.New(sha256.New, p.key)
	_, _ = mac.Write(payload)
	expectedSignature := mac.Sum(nil)
	if subtle.ConstantTimeCompare(signature, expectedSignature) != 1 {
		return VerificationResult{State: StateFailed, Reason: "signature verification failed"}, nil
	}
	return VerificationResult{State: StateVerified}, nil
}

func stringsTrim(value string) string {
	return strings.TrimSpace(value)
}
