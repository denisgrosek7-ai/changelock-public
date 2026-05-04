package formal

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func activePoint12ValCDependencySnapshot() Point12ValCDependencySnapshot {
	valB := activePoint12ValBFoundation()
	return SnapshotPoint12ValCDependencyFromComputedValB(valB, point12ValCDependencyReviewContextModel())
}

func syncPoint12ValCFoundationToDependency(model *Point12ValCFoundation) {
	model.ExportBundle.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.ExportBundle.ManifestID = model.Dependency.ValAManifest.ManifestID
	model.ExportBundle.ReplayResultID = model.Dependency.ValBReplayResult.ReplayResultID
	model.ExportBundle.DecisionID = model.Dependency.ValBReplayRequest.DecisionID
	model.ExportBundle.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.ExportBundle.ArtifactRef = model.Dependency.ValBReplayRequest.ArtifactRef
	model.ExportBundle.ArtifactHash = model.Dependency.ValBReplayRequest.ArtifactHash
	model.ExportBundle.EvidenceRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceRefs...)
	model.ExportBundle.EvidenceHashRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceHashRefs...)
	model.ExportBundle.PolicyRef = model.Dependency.ValBReplayRequest.PolicyRef
	model.ExportBundle.PolicyVersion = model.Dependency.ValBReplayRequest.PolicyVersion
	model.ExportBundle.PolicyHash = model.Dependency.ValBReplayRequest.PolicyHash
	model.ExportBundle.EngineVersion = model.Dependency.ValBReplayRequest.EngineVersion
	model.ExportBundle.EngineHash = model.Dependency.ValBReplayRequest.EngineHash
	model.ExportBundle.SchemaVersion = model.Dependency.ValBReplayRequest.SchemaVersion
	model.ExportBundle.SchemaHash = model.Dependency.ValBReplayRequest.SchemaHash
	model.ExportBundle.ClaimRefs = append([]string{}, model.Dependency.ValBReplayRequest.ClaimRefs...)
	model.ExportBundle.GovernanceEventRefs = append([]string{}, model.Dependency.ValBReplayRequest.GovernanceEventRefs...)
	model.ExportBundle.CompatibilityProfileRef = model.Dependency.ValBReplayRequest.CompatibilityProfileRef
	model.ExportBundle.RedactionManifestRef = model.Dependency.ValBReplayRequest.RedactionManifestRef
	model.ExportBundle.ManifestPayloadHash = model.Dependency.ValBReplayRequest.ManifestPayloadHash
	model.ExportBundle.SignatureMetadataRef = model.Dependency.ValAManifest.SignatureMetadataRef
	model.ExportBundle.RetentionClassRef = model.Dependency.ValAManifest.RetentionClassRef

	model.RedactionManifest.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.RedactionManifest.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.RedactionManifest.RedactionManifestID = model.Dependency.ValBReplayRequest.RedactionManifestRef
	model.RedactionManifest.RedactionPolicyRef = model.Dependency.ValBReplayRequest.PolicyRef
	model.RedactionManifest.RedactionPolicyVersion = model.Dependency.ValBReplayRequest.PolicyVersion

	model.OfflineBundle.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.OfflineBundle.ManifestID = model.Dependency.ValAManifest.ManifestID
	model.OfflineBundle.ReplayRequestID = model.Dependency.ValBReplayRequest.ReplayRequestID
	model.OfflineBundle.ReplayResultID = model.Dependency.ValBReplayResult.ReplayResultID
	model.OfflineBundle.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.OfflineBundle.ArtifactRef = model.Dependency.ValBReplayRequest.ArtifactRef
	model.OfflineBundle.ArtifactHash = model.Dependency.ValBReplayRequest.ArtifactHash
	model.OfflineBundle.EvidenceRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceRefs...)
	model.OfflineBundle.EvidenceHashRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceHashRefs...)
	model.OfflineBundle.PolicyRef = model.Dependency.ValBReplayRequest.PolicyRef
	model.OfflineBundle.PolicyVersion = model.Dependency.ValBReplayRequest.PolicyVersion
	model.OfflineBundle.PolicyHash = model.Dependency.ValBReplayRequest.PolicyHash
	model.OfflineBundle.EngineVersion = model.Dependency.ValBReplayRequest.EngineVersion
	model.OfflineBundle.EngineHash = model.Dependency.ValBReplayRequest.EngineHash
	model.OfflineBundle.SchemaVersion = model.Dependency.ValBReplayRequest.SchemaVersion
	model.OfflineBundle.SchemaHash = model.Dependency.ValBReplayRequest.SchemaHash
	model.OfflineBundle.ManifestPayloadHash = model.Dependency.ValBReplayRequest.ManifestPayloadHash
	model.OfflineBundle.SignatureMetadataRef = model.Dependency.ValAManifest.SignatureMetadataRef
	model.OfflineBundle.DetachedSignatureRef = model.Dependency.ValAManifest.DetachedSignatureRef
	model.OfflineBundle.CompatibilityProfileRef = model.Dependency.ValBReplayRequest.CompatibilityProfileRef
	model.OfflineBundle.RedactionManifestRef = model.Dependency.ValBReplayRequest.RedactionManifestRef
	model.OfflineBundle.RetentionClassRef = model.Dependency.ValAManifest.RetentionClassRef

	model.PublicPrivateBoundary.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.PublicPrivateBoundary.ExportID = model.ExportBundle.ExportID
	model.PublicPrivateBoundary.OfflineBundleID = model.OfflineBundle.OfflineBundleID
}

func activePoint12ValCFoundation() Point12ValCFoundation {
	model := Point12ValCFoundationModel()
	model.Dependency = activePoint12ValCDependencySnapshot()
	syncPoint12ValCFoundationToDependency(&model)
	return ComputePoint12ValCFoundation(model)
}

func activePoint12ValCFoundationFromValB(valB Point12ValBFoundation) Point12ValCFoundation {
	model := Point12ValCFoundationModel()
	model.Dependency = SnapshotPoint12ValCDependencyFromComputedValB(valB, point12ValCDependencyReviewContextModel())
	syncPoint12ValCFoundationToDependency(&model)
	return ComputePoint12ValCFoundation(model)
}

func readPoint12ValCSource(t *testing.T) string {
	t.Helper()
	for _, path := range []string{"point12_valc.go", "internal/formal/point12_valc.go"} {
		body, err := os.ReadFile(path)
		if err == nil {
			return string(body)
		}
	}
	t.Fatal("failed to read point12_valc.go source")
	return ""
}

func TestPoint12ValCDependencyState(t *testing.T) {
	t.Run("valid computed valb output allows valc to proceed", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.DependencyState != Point12ValCDependencyStateActive {
			t.Fatalf("expected active valc dependency state, got %#v", model)
		}
		if model.CurrentState != Point12ValCStateActive {
			t.Fatalf("expected active valc state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12ValCDependencySnapshot)
		want   string
	}{
		{name: "missing valb dependency blocks", mutate: func(model *Point12ValCDependencySnapshot) { *model = Point12ValCDependencySnapshot{} }, want: Point12ValCDependencyStateBlocked},
		{name: "fallback regenerated valb snapshot blocks", mutate: func(model *Point12ValCDependencySnapshot) { model.SnapshotFromComputedOutput = false }, want: Point12ValCDependencyStateBlocked},
		{name: "unsupported manifest integrity check requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBManifestIntegrityResult = point12ValBCheckResultUnsupported
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "unsupported signature metadata check requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBSignatureMetadataResult = point12ValBCheckResultUnsupported
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "unsupported compatibility check requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBCompatibilityResult = point12ValBCheckResultUnsupported
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "tampered manifest integrity check blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBManifestIntegrityResult = point12ValBCheckResultTampered
		}, want: Point12ValCDependencyStateBlocked},
		{name: "blocked signature metadata check blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBSignatureMetadataResult = point12ValBCheckResultBlocked
		}, want: Point12ValCDependencyStateBlocked},
		{name: "unknown manifest integrity check does not become active", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBManifestIntegrityResult = "manifest_integrity_check_result_unknown_001"
		}, want: Point12ValCDependencyStateBlocked},
		{name: "valb tamper detected blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultTamperDetected
		}, want: Point12ValCDependencyStateBlocked},
		{name: "valb unsupported version requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "valb insufficient evidence requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultInsufficientEvidence
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "valb redacted limitations requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultRedactedLimitations
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "premature point12 pass in dependency blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBPrematurePoint12PassSeen = true
		}, want: Point12ValCDependencyStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCDependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint12ValCDependencyState(model); got != testCase.want {
				t.Fatalf("expected dependency state %q, got %#v", testCase.want, model)
			}
		})
	}

	t.Run("unsupported dependency enters review path and cannot remain export_ready", func(t *testing.T) {
		model := Point12ValCFoundationModel()
		model.Dependency = activePoint12ValCDependencySnapshot()
		model.Dependency.ValBManifestIntegrityResult = point12ValBCheckResultUnsupported
		syncPoint12ValCFoundationToDependency(&model)
		model = ComputePoint12ValCFoundation(model)
		if model.DependencyState != Point12ValCDependencyStateReviewRequired {
			t.Fatalf("expected dependency review-required for unsupported manifest check, got %#v", model)
		}
		if model.ExportState == Point12ValCExportStateReady || model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected unsupported dependency to avoid full export_ready/active path, got %#v", model)
		}
	})

	t.Run("tampered dependency still blocks full valc foundation", func(t *testing.T) {
		model := Point12ValCFoundationModel()
		model.Dependency = activePoint12ValCDependencySnapshot()
		model.Dependency.ValBManifestIntegrityResult = point12ValBCheckResultTampered
		syncPoint12ValCFoundationToDependency(&model)
		model = ComputePoint12ValCFoundation(model)
		if model.DependencyState != Point12ValCDependencyStateBlocked || model.CurrentState != Point12ValCStateBlocked {
			t.Fatalf("expected tampered dependency to block valc foundation, got %#v", model)
		}
	})
}

func TestPoint12ValCAuditExportBundle(t *testing.T) {
	t.Run("valid audit ready export metadata remains bounded and active", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.ExportState != Point12ValCExportStateReady {
			t.Fatalf("expected export_ready, got %#v", model)
		}
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal valc foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected bounded export metadata to not emit point12 pass, got %s", body)
		}
	})

	testCases := []struct {
		name      string
		mutate    func(*Point12ValCFoundation)
		wantState string
	}{
		{name: "missing export id blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.ExportID = "" }, wantState: Point12ValCExportStateBlocked},
		{name: "missing tenant scope blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.TenantScope = "" }, wantState: Point12ValCExportStateTenantMismatch},
		{name: "missing proof pack manifest replay refs block", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ProofPackID = ""
			model.ExportBundle.ManifestID = ""
			model.ExportBundle.ReplayResultID = ""
		}, wantState: Point12ValCExportStateBlocked},
		{name: "missing projection disclaimer blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.ProjectionDisclaimer = "" }, wantState: Point12ValCExportStateBlocked},
		{name: "missing retention owner and disposal path returns retention missing", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.RetentionOwnerRef = ""
			model.ExportBundle.DisposalPathRef = ""
		}, wantState: Point12ValCExportStateRetentionMissing},
		{name: "public private classification missing blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.PublicPrivateClassification = "" }, wantState: Point12ValCExportStateBlocked},
		{name: "advisory only false blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.AdvisoryOnly = false }, wantState: Point12ValCExportStateBlocked},
		{name: "insufficient evidence cannot become export ready", mutate: func(model *Point12ValCFoundation) {
			model.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultInsufficientEvidence
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked},
		{name: "verifier export requires offline bundle ref", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ExportKind = point12ValCExportKindVerifierPackageMetadata
			model.ExportBundle.OfflineBundleRef = ""
		}, wantState: Point12ValCExportStateBlocked},
		{name: "export cannot accept point12 pass fixture", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.CustomerVisibleSummary = "point_12_pass"
		}, wantState: Point12ValCExportStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCFoundation()
			testCase.mutate(&model)
			model = ComputePoint12ValCFoundation(model)
			if model.ExportState != testCase.wantState {
				t.Fatalf("expected export state %q, got %#v", testCase.wantState, model)
			}
		})
	}

	t.Run("unsupported version may remain bounded review required", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		model.ExportBundle.ExportState = Point12ValCExportStateUnsupported
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateUnsupported || model.CurrentState != Point12ValCStateReviewRequired {
			t.Fatalf("expected unsupported export taxonomy with overall review-required state, got %#v", model)
		}
	})
}

func TestPoint12ValCRedactionManifestAndImpact(t *testing.T) {
	t.Run("valid non decisive redaction remains active with no decision impact", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive {
			t.Fatalf("expected active redaction manifest state, got %#v", model)
		}
		if model.RedactionImpactState != Point12ValCRedactionImpactNoDecisionImpact {
			t.Fatalf("expected no_decision_impact, got %#v", model)
		}
	})

	t.Run("missing redaction reason blocks when redacted fields exist", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"customer_email"}
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionReasons = nil
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected blocked redaction manifest, got %#v", model)
		}
	})

	t.Run("missing redaction approval event blocks where required", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"customer_email"}
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApprovalEventRef = ""
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected blocked redaction manifest, got %#v", model)
		}
	})

	t.Run("disallowed claims may contain production approved as denylist content", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive {
			t.Fatalf("expected denylist-only disallowed claim to remain active, got %#v", model)
		}
	})

	t.Run("minimum safe claim production approved blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.MinimumSafeClaimAfterRedaction = "production approved"
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected forbidden minimum safe claim to block, got %#v", model)
		}
	})

	t.Run("customer visible exported surviving replay claims with compliance guaranteed block", func(t *testing.T) {
		for _, mutate := range []func(*Point12ValCRedactionManifest){
			func(model *Point12ValCRedactionManifest) {
				model.CustomerVisibleClaimsAfterRedaction = []string{"compliance guaranteed"}
			},
			func(model *Point12ValCRedactionManifest) {
				model.ExportedClaimsAfterRedaction = []string{"compliance guaranteed"}
			},
			func(model *Point12ValCRedactionManifest) {
				model.SurvivingClaimsAfterRedaction = []string{"compliance guaranteed"}
			},
			func(model *Point12ValCRedactionManifest) {
				model.ReplayResultClaims = []string{"compliance guaranteed"}
			},
		} {
			model := activePoint12ValCFoundation()
			mutate(&model.RedactionManifest)
			model = ComputePoint12ValCFoundation(model)
			if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
				t.Fatalf("expected forbidden surviving/export claim to block, got %#v", model)
			}
		}
	})

	t.Run("internal redaction summary may describe removed forbidden claim", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model.RedactionManifest.RedactionSummary = "internal summary: disallowed production approved claim removed during redaction"
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive {
			t.Fatalf("expected internal diagnostic summary to remain active, got %#v", model)
		}
	})

	t.Run("redaction cannot hide decisive missing evidence", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"decisive_evidence_hash"}
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionAffectsDecision = true
		model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultSameDecision
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected decisive hidden evidence to block, got %#v", model)
		}
	})

	t.Run("decisive evidence removed cannot be no decision impact", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"decisive_evidence_hash"}
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionAffectsReplay = true
		model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.Limitations = []string{"decisive evidence removed"}
		model.RedactionImpactVerdict.DecisiveEvidenceRemoved = true
		model.RedactionImpactVerdict.AffectsReplay = true
		model.RedactionImpactVerdict.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionImpactVerdict.RedactionImpactState = Point12ValCRedactionImpactNoDecisionImpact
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionImpactState != Point12ValCRedactionImpactReviewRequired {
			t.Fatalf("expected review required redaction impact verdict, got %#v", model)
		}
	})

	t.Run("replay affecting redaction can produce redacted limitations with required limits", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"customer_identifier"}
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionAffectsReplay = true
		model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.Limitations = []string{"replay bounded by redaction"}
		model.RedactionImpactVerdict.AffectsReplay = true
		model.RedactionImpactVerdict.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionImpactVerdict.Limitations = []string{"replay bounded by redaction"}
		model.RedactionImpactVerdict.RequiresPartialAdvisoryExport = true
		model.RedactionImpactVerdict.RedactionImpactState = Point12ValCRedactionImpactRedactedLimits
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive || model.RedactionImpactState != Point12ValCRedactionImpactRedactedLimits {
			t.Fatalf("expected active redaction manifest with redacted limitations impact, got %#v", model)
		}
	})
}

func TestPoint12ValCOfflineBundleAndBoundary(t *testing.T) {
	t.Run("valid offline verification metadata remains active", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.OfflineBundleState != Point12ValCOfflineBundleStateActive {
			t.Fatalf("expected active offline bundle state, got %#v", model)
		}
	})

	testCases := []struct {
		name      string
		mutate    func(*Point12ValCFoundation)
		wantState string
	}{
		{name: "no external api required false blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.NoExternalAPIRequired = false
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "external api used true blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ExternalAPIUsed = true
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "missing manifest proof replay refs block", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ManifestID = ""
			model.OfflineBundle.ProofPackID = ""
			model.OfflineBundle.ReplayResultID = ""
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "unsupported verifier version returns unsupported", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.RequestedVerifierVersion = "verifier_version_point12_valc_002"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateUnsupported
		}, wantState: Point12ValCOfflineBundleStateUnsupported},
		{name: "tenant mismatch blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.TenantScope = "tenant_scope_point12_cross_001"
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "cross tenant evidence blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.EvidenceRefs = []string{"evidence:cross-tenant-pack-001"}
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "redacted decisive evidence forces limitations or blocks", mutate: func(model *Point12ValCFoundation) {
			model.RedactionManifest.RedactedFields = []string{"decisive_evidence_hash"}
			model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
			model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
			model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
			model.RedactionManifest.RedactionAffectsReplay = true
			model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultBlockedReplay
			model.RedactionManifest.PartialOrAdvisoryOnly = true
			model.RedactionManifest.Limitations = []string{"decisive evidence removed"}
			model.RedactionImpactVerdict.DecisiveEvidenceRemoved = true
			model.RedactionImpactVerdict.AffectsReplay = true
			model.RedactionImpactVerdict.PostRedactionResult = Point12Val0ReplayResultBlockedReplay
			model.RedactionImpactVerdict.Limitations = []string{"decisive evidence removed"}
			model.RedactionImpactVerdict.RequiresPartialAdvisoryExport = true
			model.RedactionImpactVerdict.RedactionImpactState = Point12ValCRedactionImpactBlockedReplay
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "offline bundle cannot accept point12 pass fixture", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.CustomerVisibleExplanation = "point_12_pass"
		}, wantState: Point12ValCOfflineBundleStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCFoundation()
			testCase.mutate(&model)
			model = ComputePoint12ValCFoundation(model)
			if model.OfflineBundleState != testCase.wantState {
				t.Fatalf("expected offline state %q, got %#v", testCase.wantState, model)
			}
		})
	}

	t.Run("all exported fields classified keeps boundary active", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateActive {
			t.Fatalf("expected active public/private boundary, got %#v", model)
		}
	})

	t.Run("missing field classification blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.PrivateFields = []string{"artifact_hash"}
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected boundary classification failure to block, got %#v", model)
		}
	})

	t.Run("private field in customer visible output blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportAudience = point12ValCExportAudienceCustomer
		model.PublicPrivateBoundary.CustomerVisibleFields = []string{"artifact_hash"}
		model.ExportBundle.CustomerVisibleSummary = "customer summary"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected private customer visible field to block, got %#v", model)
		}
	})

	t.Run("unknown audience blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.AllowedAudience = "unknown audience"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected unknown audience to block boundary, got %#v", model)
		}
	})

	t.Run("redaction summary cannot leak private field into public output", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportAudience = point12ValCExportAudienceCustomer
		model.RedactionManifest.RedactionSummary = "internal summary: artifact_hash removed during redaction"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected private field leak through redaction summary to block, got %#v", model)
		}
	})
}

func TestPoint12ValCNoOverclaimAndTaxonomy(t *testing.T) {
	t.Run("forbidden wording in export output blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportOutputClaims = []string{"production approved"}
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateBlocked {
			t.Fatalf("expected export overclaim to block, got %#v", model)
		}
	})

	t.Run("forbidden wording in customer visible summary blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.CustomerVisibleSummary = "compliance guaranteed"
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateBlocked {
			t.Fatalf("expected customer summary overclaim to block, got %#v", model)
		}
	})

	t.Run("customer facing limitations cannot overclaim", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportAudience = point12ValCExportAudienceCustomer
		model.ExportBundle.Limitations = []string{"certified"}
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateBlocked {
			t.Fatalf("expected customer-facing limitation overclaim to block, got %#v", model)
		}
	})

	t.Run("projection only is not final pass", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportState = Point12ValCExportStateProjectionOnly
		model = ComputePoint12ValCFoundation(model)
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal valc foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected projection_only to remain non-pass, got %s", body)
		}
	})

	t.Run("partial advisory export is review required not pass", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultInsufficientEvidence
		model.ExportBundle.ExportState = Point12ValCExportStatePartialAdvisory
		model.ExportBundle.Limitations = []string{"insufficient evidence for full export"}
		model = ComputePoint12ValCFoundation(model)
		if model.CurrentState != Point12ValCStateReviewRequired {
			t.Fatalf("expected partial advisory export to remain review required, got %#v", model)
		}
	})
}

func TestPoint12ValCRegressionGuards(t *testing.T) {
	t.Run("val0 computed provenance fix preserved through vala valb valc chain", func(t *testing.T) {
		valD := activePoint11ValDFoundation()
		val0 := Point12Val0FoundationModel()
		val0.Dependency = SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, Point12Val0Point11ReviewContext{
			SnapshotFromComputedOutput: false,
		})
		val0 = ComputePoint12Val0Foundation(val0)
		valA := activePoint12ValAFoundationFromVal0(val0)
		valB := activePoint12ValBFoundationFromValA(valA)
		model := activePoint12ValCFoundationFromValB(valB)
		if model.DependencyState != Point12ValCDependencyStateBlocked {
			t.Fatalf("expected non-computed upstream provenance to stay blocked through valc, got %#v", model)
		}
	})

	t.Run("vala manifest tamper behavior preserved", func(t *testing.T) {
		valA := activePoint12ValAFoundation()
		valA.ManifestIntegrityState = Point12ValAManifestIntegrityStateTampered
		valB := activePoint12ValBFoundationFromValA(valA)
		model := activePoint12ValCFoundationFromValB(valB)
		if model.DependencyState != Point12ValCDependencyStateBlocked {
			t.Fatalf("expected tampered vala manifest to block valc dependency, got %#v", model)
		}
	})

	t.Run("valb original context cannot silently use current policy preserved", func(t *testing.T) {
		valB := activePoint12ValBFoundation()
		valB.ReplayRequest.CurrentPolicyRef = "policy_point12_current_001"
		valB.ReplayRequest.CurrentPolicyVersion = "policy_version_point12_current_001"
		valB.ReplayRequest.CurrentPolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
		valB = ComputePoint12ValBFoundation(valB)
		model := activePoint12ValCFoundationFromValB(valB)
		if model.DependencyState != Point12ValCDependencyStateBlocked {
			t.Fatalf("expected blocked valc dependency from invalid original_context replay, got %#v", model)
		}
	})

	t.Run("no real signing or external api side effects introduced", func(t *testing.T) {
		body := readPoint12ValCSource(t)
		for _, forbidden := range []string{
			"http.Get",
			"http.Post",
			"fetch(",
			"Sign(",
			"GenerateKey",
			"crypto/rsa",
			"crypto/ecdsa",
			"crypto/ed25519",
		} {
			if strings.Contains(body, forbidden) {
				t.Fatalf("unexpected valc source boundary violation %q", forbidden)
			}
		}
	})
}
