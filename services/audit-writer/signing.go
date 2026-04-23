package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/signing"
)

type signingRuntime struct {
	runtime *signing.Runtime
}

func loadSigningRuntimeFromEnv() (*signingRuntime, error) {
	config, err := signing.ParseEnvConfig(os.Getenv)
	if err != nil {
		return nil, err
	}
	runtime, err := signing.NewRuntime(config, signing.ProviderOptions{})
	if err != nil {
		return nil, err
	}
	return &signingRuntime{runtime: runtime}, nil
}

func (s *signingRuntime) mode() string {
	if s == nil || s.runtime == nil || strings.TrimSpace(s.runtime.Config.Mode) == "" {
		return signing.ModeDisabled
	}
	return s.runtime.Config.Mode
}

func (s *signingRuntime) supportsPurpose(purpose string) bool {
	return s != nil && s.runtime != nil && s.runtime.SupportsPurpose(purpose)
}

func (s *signingRuntime) enabledForPurpose(purpose string) bool {
	return s != nil && s.runtime != nil && s.runtime.Enabled() && s.runtime.SupportsPurpose(purpose)
}

func (s *signingRuntime) verifyOnRead(purpose string) bool {
	return s.enabledForPurpose(purpose) && s.runtime.Config.VerifyOnRead
}

func (s *signingRuntime) signException(ctx context.Context, exception audit.PolicyException) (*signing.Envelope, error) {
	if !s.enabledForPurpose(signing.PurposeExceptions) {
		return nil, nil
	}
	payload, err := audit.CanonicalExceptionEvidence(exception)
	if err != nil {
		return nil, err
	}
	envelope, err := s.runtime.Sign(ctx, signing.PurposeExceptions, payload)
	if err != nil {
		return nil, err
	}
	return &envelope, nil
}

func (s *signingRuntime) verifyException(ctx context.Context, exception audit.PolicyException) (signing.VerificationResult, error) {
	if !s.enabledForPurpose(signing.PurposeExceptions) {
		return signing.VerificationResult{State: signing.StateDisabled}, nil
	}
	if exception.Signature == nil {
		if s.verifyOnRead(signing.PurposeExceptions) {
			return signing.VerificationResult{State: signing.StateFailed, Reason: "approved exception evidence signature is missing"}, nil
		}
		return signing.VerificationResult{State: signing.StateUnverified, Reason: "approved exception evidence signature is missing"}, nil
	}
	if !s.runtime.Config.VerifyOnRead {
		return signing.VerificationResult{State: signing.StateUnverified, Reason: "verify on read is disabled"}, nil
	}
	payload, err := audit.CanonicalExceptionEvidence(exception)
	if err != nil {
		return signing.VerificationResult{}, err
	}
	return s.runtime.Verify(ctx, signing.PurposeExceptions, payload, *exception.Signature)
}

func (s *signingRuntime) signSyncSnapshot(ctx context.Context, snapshot audit.ExceptionSyncSnapshot) (*signing.Envelope, error) {
	if !s.enabledForPurpose(signing.PurposeSyncSnapshots) {
		return nil, nil
	}
	payload, err := audit.CanonicalExceptionSyncSnapshot(snapshot)
	if err != nil {
		return nil, err
	}
	envelope, err := s.runtime.Sign(ctx, signing.PurposeSyncSnapshots, payload)
	if err != nil {
		return nil, err
	}
	return &envelope, nil
}

func (s *signingRuntime) verifySyncSnapshot(ctx context.Context, snapshot audit.ExceptionSyncSnapshot) (signing.VerificationResult, error) {
	if !s.enabledForPurpose(signing.PurposeSyncSnapshots) {
		return signing.VerificationResult{State: signing.StateDisabled}, nil
	}
	if snapshot.Signature == nil {
		if s.verifyOnRead(signing.PurposeSyncSnapshots) {
			return signing.VerificationResult{State: signing.StateFailed, Reason: "sync snapshot signature is missing"}, nil
		}
		return signing.VerificationResult{State: signing.StateUnverified, Reason: "sync snapshot signature is missing"}, nil
	}
	if !s.runtime.Config.VerifyOnRead {
		return signing.VerificationResult{State: signing.StateUnverified, Reason: "verify on read is disabled"}, nil
	}
	payload, err := audit.CanonicalExceptionSyncSnapshot(snapshot)
	if err != nil {
		return signing.VerificationResult{}, err
	}
	return s.runtime.Verify(ctx, signing.PurposeSyncSnapshots, payload, *snapshot.Signature)
}

func (s *signingRuntime) signPublicProofArtifact(ctx context.Context, payload []byte) (*signing.Envelope, error) {
	if !s.enabledForPurpose(signing.PurposePublicProofArtifact) {
		return nil, nil
	}
	envelope, err := s.runtime.Sign(ctx, signing.PurposePublicProofArtifact, payload)
	if err != nil {
		return nil, err
	}
	return &envelope, nil
}

func (s *signingRuntime) verifyPublicProofArtifact(ctx context.Context, payload []byte, envelope signing.Envelope) (signing.VerificationResult, error) {
	if !s.enabledForPurpose(signing.PurposePublicProofArtifact) {
		return signing.VerificationResult{State: signing.StateDisabled}, nil
	}
	return s.runtime.Verify(ctx, signing.PurposePublicProofArtifact, payload, envelope)
}

func (s server) signAndPersistException(ctx context.Context, exception audit.PolicyException) (audit.PolicyException, error) {
	if s.signing == nil || !s.signing.enabledForPurpose(signing.PurposeExceptions) {
		exception.VerificationState = signing.StateDisabled
		exception.VerificationReason = ""
		return exception, nil
	}
	envelope, err := s.signing.signException(ctx, exception)
	if err != nil {
		return audit.PolicyException{}, fmt.Errorf("sign approved exception evidence: %w", err)
	}
	if envelope == nil {
		exception.VerificationState = signing.StateDisabled
		return exception, nil
	}
	signedException, err := s.store.SetExceptionSignature(ctx, exception.ExceptionID, envelope)
	if err != nil {
		return audit.PolicyException{}, fmt.Errorf("persist approved exception signature: %w", err)
	}
	return signedException, nil
}
