package verify

import (
	"context"

	"github.com/denisgrosek/changelock/internal/evidence"
)

const DefaultPredicateType = "slsaprovenance"

type ArtifactVerifier interface {
	VerifyArtifact(ctx context.Context, request ArtifactVerificationRequest) (ArtifactVerification, error)
}

type ArtifactVerificationRequest struct {
	Image                   string               `json:"image" yaml:"image"`
	ExpectedRepository      string               `json:"expectedRepository,omitempty" yaml:"expectedRepository,omitempty"`
	ExpectedRef             string               `json:"expectedRef,omitempty" yaml:"expectedRef,omitempty"`
	ExpectedCommitSHA       string               `json:"expectedCommitSHA,omitempty" yaml:"expectedCommitSHA,omitempty"`
	AllowedSignerIdentities []string             `json:"allowedSignerIdentities,omitempty" yaml:"allowedSignerIdentities,omitempty"`
	AllowedOIDCIssuers      []string             `json:"allowedOidcIssuers,omitempty" yaml:"allowedOidcIssuers,omitempty"`
	PredicateType           string               `json:"predicateType,omitempty" yaml:"predicateType,omitempty"`
	SupplyChain             *SupplyChainEvidence `json:"supplyChain,omitempty" yaml:"supplyChain,omitempty"`
	EvidenceBundle          *evidence.Bundle     `json:"evidenceBundle,omitempty" yaml:"evidenceBundle,omitempty"`
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
	MatchedIdentity          string               `json:"matchedIdentity,omitempty" yaml:"matchedIdentity,omitempty"`
	SignatureClaimsCount     int                  `json:"signatureClaimsCount,omitempty" yaml:"signatureClaimsCount,omitempty"`
	AttestationCount         int                  `json:"attestationCount,omitempty" yaml:"attestationCount,omitempty"`
	AttestationPredicateType string               `json:"attestationPredicateType,omitempty" yaml:"attestationPredicateType,omitempty"`
	AttestationSubjectName   string               `json:"attestationSubjectName,omitempty" yaml:"attestationSubjectName,omitempty"`
	AttestationSubjectDigest string               `json:"attestationSubjectDigest,omitempty" yaml:"attestationSubjectDigest,omitempty"`
	SupplyChain              *SupplyChainEvidence `json:"supplyChain,omitempty" yaml:"supplyChain,omitempty"`
	Bundle                   *evidence.Bundle     `json:"bundle,omitempty" yaml:"bundle,omitempty"`
	TransparencyLogState     string               `json:"transparencyLogState,omitempty" yaml:"transparencyLogState,omitempty"`
	TransparencyLogReason    string               `json:"transparencyLogReason,omitempty" yaml:"transparencyLogReason,omitempty"`
}

type SupplyChainEvidence struct {
	SBOMFormat                         string                `json:"sbomFormat,omitempty" yaml:"sbomFormat,omitempty"`
	SBOMDigestRef                      string                `json:"sbomDigestRef,omitempty" yaml:"sbomDigestRef,omitempty"`
	SBOMHash                           string                `json:"sbomHash,omitempty" yaml:"sbomHash,omitempty"`
	SBOMArtifactRef                    string                `json:"sbomArtifactRef,omitempty" yaml:"sbomArtifactRef,omitempty"`
	VulnerabilityScanStatus            string                `json:"vulnerabilityScanStatus,omitempty" yaml:"vulnerabilityScanStatus,omitempty"`
	VulnerabilityScanTool              string                `json:"vulnerabilityScanTool,omitempty" yaml:"vulnerabilityScanTool,omitempty"`
	VulnerabilityScanSeverityThreshold string                `json:"vulnerabilityScanSeverityThreshold,omitempty" yaml:"vulnerabilityScanSeverityThreshold,omitempty"`
	VulnerabilitySummary               *VulnerabilitySummary `json:"vulnerabilitySummary,omitempty" yaml:"vulnerabilitySummary,omitempty"`
	VulnerabilityReportRef             string                `json:"vulnerabilityReportRef,omitempty" yaml:"vulnerabilityReportRef,omitempty"`
}

type VulnerabilitySummary struct {
	Critical int `json:"critical,omitempty" yaml:"critical,omitempty"`
	High     int `json:"high,omitempty" yaml:"high,omitempty"`
	Medium   int `json:"medium,omitempty" yaml:"medium,omitempty"`
	Low      int `json:"low,omitempty" yaml:"low,omitempty"`
	Unknown  int `json:"unknown,omitempty" yaml:"unknown,omitempty"`
	Total    int `json:"total,omitempty" yaml:"total,omitempty"`
}

func (e *SupplyChainEvidence) MatchesDigest(digest string) bool {
	if e == nil {
		return false
	}

	expected := firstNonEmpty(digest, digestFromImage(e.SBOMDigestRef))
	if expected == "" {
		return false
	}

	actual := firstNonEmpty(digestFromImage(e.SBOMDigestRef), e.SBOMDigestRef)
	return actual == expected
}
