package verify

import "context"

const DefaultPredicateType = "slsaprovenance"

type ArtifactVerifier interface {
	VerifyArtifact(ctx context.Context, request ArtifactVerificationRequest) (ArtifactVerification, error)
}

type ArtifactVerificationRequest struct {
	Image                   string   `json:"image" yaml:"image"`
	ExpectedRepository      string   `json:"expectedRepository,omitempty" yaml:"expectedRepository,omitempty"`
	ExpectedRef             string   `json:"expectedRef,omitempty" yaml:"expectedRef,omitempty"`
	ExpectedCommitSHA       string   `json:"expectedCommitSHA,omitempty" yaml:"expectedCommitSHA,omitempty"`
	AllowedSignerIdentities []string `json:"allowedSignerIdentities,omitempty" yaml:"allowedSignerIdentities,omitempty"`
	AllowedOIDCIssuers      []string `json:"allowedOidcIssuers,omitempty" yaml:"allowedOidcIssuers,omitempty"`
	PredicateType           string   `json:"predicateType,omitempty" yaml:"predicateType,omitempty"`
}

type ArtifactVerification struct {
	SignatureValid    bool                 `json:"signatureValid" yaml:"signatureValid"`
	AttestationValid  bool                 `json:"attestationValid" yaml:"attestationValid"`
	VerifiedIdentity  string               `json:"verifiedIdentity,omitempty" yaml:"verifiedIdentity,omitempty"`
	VerifiedIssuer    string               `json:"verifiedIssuer,omitempty" yaml:"verifiedIssuer,omitempty"`
	VerifiedSubject   string               `json:"verifiedSubject,omitempty" yaml:"verifiedSubject,omitempty"`
	VerifiedRepo      string               `json:"verifiedRepo,omitempty" yaml:"verifiedRepo,omitempty"`
	VerifiedWorkflow  string               `json:"verifiedWorkflow,omitempty" yaml:"verifiedWorkflow,omitempty"`
	VerifiedRef       string               `json:"verifiedRef,omitempty" yaml:"verifiedRef,omitempty"`
	VerifiedCommitSHA string               `json:"verifiedCommitSHA,omitempty" yaml:"verifiedCommitSHA,omitempty"`
	VerifiedDigest    string               `json:"verifiedDigest,omitempty" yaml:"verifiedDigest,omitempty"`
	Reasons           []string             `json:"reasons,omitempty" yaml:"reasons,omitempty"`
	Evidence          VerificationEvidence `json:"evidence,omitempty" yaml:"evidence,omitempty"`
}

type VerificationEvidence struct {
	MatchedIdentity          string `json:"matchedIdentity,omitempty" yaml:"matchedIdentity,omitempty"`
	SignatureClaimsCount     int    `json:"signatureClaimsCount,omitempty" yaml:"signatureClaimsCount,omitempty"`
	AttestationCount         int    `json:"attestationCount,omitempty" yaml:"attestationCount,omitempty"`
	AttestationPredicateType string `json:"attestationPredicateType,omitempty" yaml:"attestationPredicateType,omitempty"`
	AttestationSubjectName   string `json:"attestationSubjectName,omitempty" yaml:"attestationSubjectName,omitempty"`
	AttestationSubjectDigest string `json:"attestationSubjectDigest,omitempty" yaml:"attestationSubjectDigest,omitempty"`
}
