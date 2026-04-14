package verify

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type FixtureVerifier struct {
	entries []FixtureEntry
}

type FixtureDocument struct {
	Artifacts []FixtureEntry `yaml:"artifacts"`
}

type FixtureEntry struct {
	Image              string               `yaml:"image"`
	ExpectedRepository string               `yaml:"expectedRepository,omitempty"`
	ExpectedRef        string               `yaml:"expectedRef,omitempty"`
	ExpectedCommitSHA  string               `yaml:"expectedCommitSHA,omitempty"`
	Error              string               `yaml:"error,omitempty"`
	Result             ArtifactVerification `yaml:"result,omitempty"`
}

func NewFixtureVerifier(path string) (*FixtureVerifier, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read verifier fixture: %w", err)
	}

	var document FixtureDocument
	if err := yaml.Unmarshal(data, &document); err != nil {
		return nil, fmt.Errorf("decode verifier fixture: %w", err)
	}

	return &FixtureVerifier{entries: document.Artifacts}, nil
}

func (v *FixtureVerifier) VerifyArtifact(_ context.Context, request ArtifactVerificationRequest) (ArtifactVerification, error) {
	if v == nil {
		return ArtifactVerification{}, errors.New("fixture verifier is nil")
	}

	for _, entry := range v.entries {
		if !entry.matches(request) {
			continue
		}
		if entry.Error != "" {
			return ArtifactVerification{}, errors.New(entry.Error)
		}

		result := entry.Result
		if result.VerifiedDigest == "" {
			result.VerifiedDigest = digestFromImage(request.Image)
		}
		if result.VerifiedRepo == "" {
			result.VerifiedRepo = request.ExpectedRepository
		}
		if result.VerifiedRef == "" {
			result.VerifiedRef = request.ExpectedRef
		}
		if result.VerifiedCommitSHA == "" {
			result.VerifiedCommitSHA = request.ExpectedCommitSHA
		}
		if result.VerifiedSubject == "" && result.VerifiedRepo != "" {
			result.VerifiedSubject = repoSubject(result.VerifiedRepo)
		}
		mergeSupplyChainEvidence(&result.Evidence.SupplyChain, request.SupplyChain)
		return result, nil
	}

	return ArtifactVerification{}, fmt.Errorf("no verifier fixture matched %q", request.Image)
}

func (e FixtureEntry) matches(request ArtifactVerificationRequest) bool {
	if e.Image != request.Image {
		return false
	}
	if e.ExpectedRepository != "" && e.ExpectedRepository != request.ExpectedRepository {
		return false
	}
	if e.ExpectedRef != "" && e.ExpectedRef != request.ExpectedRef {
		return false
	}
	if e.ExpectedCommitSHA != "" && e.ExpectedCommitSHA != request.ExpectedCommitSHA {
		return false
	}
	return true
}
