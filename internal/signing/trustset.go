package signing

import (
	"context"
	"strings"
)

type TrustSetMember struct {
	MemberID       string   `json:"member_id"`
	Runtime        *Runtime `json:"-"`
	LifecycleState string   `json:"lifecycle_state"`
}

type TrustSet struct {
	Members []TrustSetMember `json:"members,omitempty"`
}

type VerificationPath struct {
	MemberID       string `json:"member_id"`
	ProviderMode   string `json:"provider_mode"`
	KeyID          string `json:"key_id,omitempty"`
	LifecycleState string `json:"lifecycle_state"`
}

func (s TrustSet) Verify(ctx context.Context, purpose string, payload []byte, envelope Envelope) (VerificationResult, VerificationPath, error) {
	var bestResult VerificationResult
	for _, member := range s.Members {
		if member.Runtime == nil || !member.Runtime.Enabled() {
			continue
		}
		if member.Runtime.Config.Mode != envelope.Provider {
			continue
		}
		memberKeyID := strings.TrimSpace(member.Runtime.Config.KeyID)
		envelopeKeyID := strings.TrimSpace(envelope.KeyID)
		if memberKeyID != "" && envelopeKeyID != "" && memberKeyID != envelopeKeyID {
			continue
		}
		switch member.LifecycleState {
		case KeyStateRevoked, KeyStateDestroyed:
			bestResult = VerificationResult{State: StateFailed, Reason: "matching trust-set member is revoked or destroyed"}
			continue
		}
		result, err := member.Runtime.Verify(ctx, purpose, payload, envelope)
		if err != nil {
			return VerificationResult{}, VerificationPath{}, err
		}
		if result.State == StateVerified {
			return result, VerificationPath{
				MemberID:       member.MemberID,
				ProviderMode:   member.Runtime.Config.Mode,
				KeyID:          member.Runtime.Config.KeyID,
				LifecycleState: member.LifecycleState,
			}, nil
		}
		bestResult = result
	}
	if bestResult.State == "" {
		bestResult = VerificationResult{State: StateFailed, Reason: "no trust-set member matched the envelope provider and key identity"}
	}
	return bestResult, VerificationPath{}, nil
}
