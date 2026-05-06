package formal

import "testing"

func TestPoint14ValCDependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValCDependencySnapshot)
		want   string
	}{
		{
			name: "missing point14 valb blocks",
			mutate: func(model *Point14ValCDependencySnapshot) {
				model.Point14ValBCurrentState = ""
			},
			want: Point14ValCStateBlocked,
		},
		{
			name: "point14 valb blocked blocks",
			mutate: func(model *Point14ValCDependencySnapshot) {
				model.Point14ValBCurrentState = Point14ValBStateBlocked
			},
			want: Point14ValCStateBlocked,
		},
		{
			name: "point14 valb review required prevents active",
			mutate: func(model *Point14ValCDependencySnapshot) {
				model.Point14ValBCurrentState = Point14ValBStateReviewRequired
			},
			want: Point14ValCStateBlocked,
		},
		{
			name: "point14 valb incomplete prevents active",
			mutate: func(model *Point14ValCDependencySnapshot) {
				model.Point14ValBCurrentState = Point14ValBStateIncomplete
			},
			want: Point14ValCStateBlocked,
		},
		{
			name: "premature point14 pass blocks",
			mutate: func(model *Point14ValCDependencySnapshot) {
				model.Point14PassSeen = true
			},
			want: Point14ValCStateBlocked,
		},
		{
			name: "local valc readiness cannot override missing valb closure",
			mutate: func(model *Point14ValCDependencySnapshot) {
				model.Point14ValBCurrentState = ""
				model.InheritedPoint14ValACurrentState = Point14ValAStateActive
				model.InheritedPoint13ValEPassAllowed = true
				model.InheritedPoint13ValEPassToken = point13ValEPoint13PassToken
			},
			want: Point14ValCStateBlocked,
		},
		{
			name: "embedded point14 valb mismatch blocks",
			mutate: func(model *Point14ValCDependencySnapshot) {
				model.Point14ValB.CurrentState = Point14ValBStateBlocked
			},
			want: Point14ValCStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCDependencySnapshotModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValCDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCCorrectionNoticeBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalCorrectionNoticeBoundary)
		want   string
	}{
		{"valid bounded correction notice active", func(model *ExternalCorrectionNoticeBoundary) {}, Point14ValCStateActive},
		{"missing correction notice id blocks", func(model *ExternalCorrectionNoticeBoundary) { model.CorrectionNoticeID = "" }, Point14ValCStateBlocked},
		{"missing dispute ref blocks", func(model *ExternalCorrectionNoticeBoundary) { model.DisputeRef = "" }, Point14ValCStateBlocked},
		{"missing conflict set ref blocks", func(model *ExternalCorrectionNoticeBoundary) { model.ConflictSetRef = "" }, Point14ValCStateBlocked},
		{"missing corrected signal refs blocks", func(model *ExternalCorrectionNoticeBoundary) { model.CorrectedSignalRefs = nil }, Point14ValCStateBlocked},
		{"missing affected evidence refs blocks", func(model *ExternalCorrectionNoticeBoundary) { model.AffectedEvidenceRefs = nil }, Point14ValCStateBlocked},
		{"missing correction reason blocks", func(model *ExternalCorrectionNoticeBoundary) { model.CorrectionReason = "" }, Point14ValCStateBlocked},
		{"missing limitations blocks", func(model *ExternalCorrectionNoticeBoundary) { model.CorrectionLimitations = nil }, Point14ValCStateBlocked},
		{"missing governance event ref blocks", func(model *ExternalCorrectionNoticeBoundary) { model.GovernanceEventRef = "" }, Point14ValCStateBlocked},
		{"missing audit ref blocks", func(model *ExternalCorrectionNoticeBoundary) { model.AuditRef = "" }, Point14ValCStateBlocked},
		{"correction auto published blocks", func(model *ExternalCorrectionNoticeBoundary) { model.CorrectionState = "correction_auto_published" }, Point14ValCStateBlocked},
		{"correction certified blocks", func(model *ExternalCorrectionNoticeBoundary) { model.CertifiesCompliance = true }, Point14ValCStateBlocked},
		{"correction production approved blocks", func(model *ExternalCorrectionNoticeBoundary) { model.ApprovesProduction = true }, Point14ValCStateBlocked},
		{"correction public badge blocks", func(model *ExternalCorrectionNoticeBoundary) { model.CreatesPublicBadge = true }, Point14ValCStateBlocked},
		{"correction cannot mutate canonical evidence", func(model *ExternalCorrectionNoticeBoundary) { model.CanonicalMutationRequested = true }, Point14ValCStateBlocked},
		{"correction cannot override canonical decision", func(model *ExternalCorrectionNoticeBoundary) { model.OverridesCanonicalDecision = true }, Point14ValCStateBlocked},
		{"correction cannot emit pass", func(model *ExternalCorrectionNoticeBoundary) { model.EmitsPass = true }, Point14ValCStateBlocked},
		{"correction cannot strengthen claim", func(model *ExternalCorrectionNoticeBoundary) { model.StrengthensClaim = true }, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCCorrectionNoticeBoundaryModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValCCorrectionNoticeBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCRevocationRequestBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalRevocationRequestBoundary)
		want   string
	}{
		{"valid revocation request active as bounded request only", func(model *ExternalRevocationRequestBoundary) {}, Point14ValCStateActive},
		{"missing revocation request id blocks", func(model *ExternalRevocationRequestBoundary) { model.RevocationRequestID = "" }, Point14ValCStateBlocked},
		{"missing target refs blocks", func(model *ExternalRevocationRequestBoundary) {
			model.TargetClaimRefs = nil
			model.TargetSignalRefs = nil
		}, Point14ValCStateBlocked},
		{"missing revocation reason blocks", func(model *ExternalRevocationRequestBoundary) { model.RevocationReason = "" }, Point14ValCStateBlocked},
		{"missing evidence refs blocks", func(model *ExternalRevocationRequestBoundary) { model.EvidenceRefs = nil }, Point14ValCStateBlocked},
		{"missing governance event ref blocks approved state", func(model *ExternalRevocationRequestBoundary) {
			model.RevocationState = point14ValCRevocationApprovedGovernance
			model.GovernanceEventRef = ""
		}, Point14ValCStateBlocked},
		{"revocation auto executed blocks", func(model *ExternalRevocationRequestBoundary) { model.RevocationState = "revocation_auto_executed" }, Point14ValCStateBlocked},
		{"revocation external authority blocks", func(model *ExternalRevocationRequestBoundary) { model.ExternalAuthorityGranted = true }, Point14ValCStateBlocked},
		{"revocation public badge blocks", func(model *ExternalRevocationRequestBoundary) { model.PublicBadgeAllowed = true }, Point14ValCStateBlocked},
		{"revocation pass blocks", func(model *ExternalRevocationRequestBoundary) { model.PassAllowed = true }, Point14ValCStateBlocked},
		{"revocation cannot mutate canonical claim evidence directly", func(model *ExternalRevocationRequestBoundary) { model.CanonicalMutationRequested = true }, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCRevocationRequestBoundaryModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValCRevocationRequestBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCSupersessionRecordBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalSupersessionRecordBoundary)
		want   string
	}{
		{"valid supersession record active", func(model *ExternalSupersessionRecordBoundary) {}, Point14ValCStateActive},
		{"missing previous signal ref blocks", func(model *ExternalSupersessionRecordBoundary) { model.PreviousSignalRef = "" }, Point14ValCStateBlocked},
		{"missing replacement signal ref blocks", func(model *ExternalSupersessionRecordBoundary) { model.ReplacementSignalRef = "" }, Point14ValCStateBlocked},
		{"missing supersession reason blocks", func(model *ExternalSupersessionRecordBoundary) { model.SupersessionReason = "" }, Point14ValCStateBlocked},
		{"missing evidence refs blocks", func(model *ExternalSupersessionRecordBoundary) { model.EvidenceRefs = nil }, Point14ValCStateBlocked},
		{"missing governance event ref blocks", func(model *ExternalSupersessionRecordBoundary) { model.GovernanceEventRef = "" }, Point14ValCStateBlocked},
		{"silent replacement blocks", func(model *ExternalSupersessionRecordBoundary) { model.SilentReplacement = true }, Point14ValCStateBlocked},
		{"previous evidence deletion hiding blocks", func(model *ExternalSupersessionRecordBoundary) { model.DeletesHistory = true }, Point14ValCStateBlocked},
		{"replacement hash recomputation hiding drift blocks", func(model *ExternalSupersessionRecordBoundary) { model.ReplacementHashRecomputed = true }, Point14ValCStateBlocked},
		{"missing supersession audit trace blocks", func(model *ExternalSupersessionRecordBoundary) { model.AuditRef = "" }, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCSupersessionRecordBoundaryModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValCSupersessionRecordBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCPublicationApprovalBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalPublicationApprovalBoundary)
		want   string
	}{
		{"valid bounded publication approval active", func(model *ExternalPublicationApprovalBoundary) {}, Point14ValCStateActive},
		{"missing publication approval id blocks", func(model *ExternalPublicationApprovalBoundary) { model.PublicationApprovalID = "" }, Point14ValCStateBlocked},
		{"missing referenced correction revocation supersession ref blocks", func(model *ExternalPublicationApprovalBoundary) {
			model.CorrectionNoticeRef = ""
			model.RevocationRequestRef = ""
			model.SupersessionRecordRef = ""
		}, Point14ValCStateBlocked},
		{"missing approver role blocks", func(model *ExternalPublicationApprovalBoundary) { model.ApproverRole = "" }, Point14ValCStateBlocked},
		{"missing approver ref blocks", func(model *ExternalPublicationApprovalBoundary) { model.ApproverRef = "" }, Point14ValCStateBlocked},
		{"missing approval reason blocks", func(model *ExternalPublicationApprovalBoundary) { model.ApprovalReason = "" }, Point14ValCStateBlocked},
		{"missing approval scope blocks", func(model *ExternalPublicationApprovalBoundary) { model.ApprovalScope = "" }, Point14ValCStateBlocked},
		{"missing audit ref blocks", func(model *ExternalPublicationApprovalBoundary) { model.AuditRef = "" }, Point14ValCStateBlocked},
		{"missing governance event ref blocks", func(model *ExternalPublicationApprovalBoundary) { model.GovernanceEventRef = "" }, Point14ValCStateBlocked},
		{"client local time as approved at source blocks", func(model *ExternalPublicationApprovalBoundary) {
			model.ApprovedTimeSource = point14Val0TimeSourceClientLocal
		}, Point14ValCStateBlocked},
		{"publication auto approved blocks", func(model *ExternalPublicationApprovalBoundary) { model.AutomaticApprovalRequested = true }, Point14ValCStateBlocked},
		{"publication production approved blocks", func(model *ExternalPublicationApprovalBoundary) { model.ApprovesProduction = true }, Point14ValCStateBlocked},
		{"publication certified blocks", func(model *ExternalPublicationApprovalBoundary) { model.Certifies = true }, Point14ValCStateBlocked},
		{"publication public badge blocks", func(model *ExternalPublicationApprovalBoundary) { model.CreatesPublicBadge = true }, Point14ValCStateBlocked},
		{"publication global truth blocks", func(model *ExternalPublicationApprovalBoundary) { model.GlobalTruthRequested = true }, Point14ValCStateBlocked},
		{"publication approval cannot create pass", func(model *ExternalPublicationApprovalBoundary) { model.CreatesPass = true }, Point14ValCStateBlocked},
		{"future dated publication approval review required", func(model *ExternalPublicationApprovalBoundary) { model.RecordedAt = "2026-05-06T08:04:00Z" }, Point14ValCStateReviewRequired},
		{"backdated correction approval before dispute open review required", func(model *ExternalPublicationApprovalBoundary) { model.ApprovedAt = "2026-05-06T07:55:00Z" }, Point14ValCStateReviewRequired},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCPublicationApprovalBoundaryModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValCPublicationApprovalBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCPublicationVisibilityBoundaryState(t *testing.T) {
	dependency := point14ValCDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalPublicationVisibilityBoundary)
		want   string
	}{
		{"private tenant only visibility active", func(model *ExternalPublicationVisibilityBoundary) {}, Point14ValCStateActive},
		{"customer shared bounded active with limitations", func(model *ExternalPublicationVisibilityBoundary) {
			model.VisibilityClassification = point14ValCVisibilityCustomerBounded
			model.PublicPrivateBoundary = point14ValCPublicationBoundaryCustomer
		}, Point14ValCStateActive},
		{"auditor shared bounded active with limitations", func(model *ExternalPublicationVisibilityBoundary) {
			model.VisibilityClassification = point14ValCVisibilityAuditorBounded
			model.PublicPrivateBoundary = point14ValCPublicationBoundaryAuditor
		}, Point14ValCStateActive},
		{"public notice bounded active only with limitations redaction privacy boundary", func(model *ExternalPublicationVisibilityBoundary) {
			model.VisibilityClassification = point14ValCVisibilityPublicBounded
			model.PublicPrivateBoundary = point14ValCPublicationBoundaryPublic
			model.RedactionRefs = []string{"redaction_ref_point14_valc_001"}
		}, Point14ValCStateActive},
		{"publication blocked blocks", func(model *ExternalPublicationVisibilityBoundary) {
			model.VisibilityClassification = point14ValCVisibilityBlocked
		}, Point14ValCStateBlocked},
		{"missing visibility classification blocks", func(model *ExternalPublicationVisibilityBoundary) { model.VisibilityClassification = "" }, Point14ValCStateBlocked},
		{"public notice without limitations blocks", func(model *ExternalPublicationVisibilityBoundary) {
			model.VisibilityClassification = point14ValCVisibilityPublicBounded
			model.PublicPrivateBoundary = point14ValCPublicationBoundaryPublic
			model.LimitationRefs = nil
			model.RedactionRefs = []string{"redaction_ref_point14_valc_001"}
		}, Point14ValCStateBlocked},
		{"public notice exposing tenant private data blocks", func(model *ExternalPublicationVisibilityBoundary) { model.TenantPrivateDataExposed = true }, Point14ValCStateBlocked},
		{"public notice implying official authority blocks", func(model *ExternalPublicationVisibilityBoundary) { model.ImpliesPublicAuthority = true }, Point14ValCStateBlocked},
		{"public notice omitting private limitations that change meaning blocks", func(model *ExternalPublicationVisibilityBoundary) {
			model.MeaningChangingPrivateLimitationOmitted = true
		}, Point14ValCStateBlocked},
		{"public notice bounded with tenant private boundary blocks", func(model *ExternalPublicationVisibilityBoundary) {
			model.VisibilityClassification = point14ValCVisibilityPublicBounded
			model.PublicPrivateBoundary = point14ValCPublicationBoundaryPrivate
			model.RedactionRefs = []string{"redaction_ref_point14_valc_001"}
		}, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCPublicationVisibilityBoundaryModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValCPublicationVisibilityBoundaryState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCTenantPrivacyPublicationGuardState(t *testing.T) {
	dependency := point14ValCDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*TenantPrivacyPublicationGuard)
		want   string
	}{
		{"tenant privacy active", func(model *TenantPrivacyPublicationGuard) {}, Point14ValCStateActive},
		{"cross tenant publication blocks clb0", func(model *TenantPrivacyPublicationGuard) { model.CrossTenantPublication = true }, Point14ValCStateBlocked},
		{"tenant private data exposed blocks clb0", func(model *TenantPrivacyPublicationGuard) { model.TenantPrivateDataExposed = true }, Point14ValCStateBlocked},
		{"global publication cannot attach tenant private evidence without governed redaction", func(model *TenantPrivacyPublicationGuard) {
			model.PublicationTargetScope = point14ValCVisibilityPublicBounded
			model.PublicPrivateClassification = point14ValCPublicationBoundaryPublic
			model.RedactionRefs = nil
		}, Point14ValCStateBlocked},
		{"public private classification missing blocks", func(model *TenantPrivacyPublicationGuard) { model.PublicPrivateClassification = "" }, Point14ValCStateBlocked},
		{"redacted publication must preserve limitations", func(model *TenantPrivacyPublicationGuard) { model.LimitationsVisible = false }, Point14ValCStateBlocked},
		{"blocked publication target scope blocks", func(model *TenantPrivacyPublicationGuard) {
			model.PublicationTargetScope = point14ValCVisibilityBlocked
		}, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCTenantPrivacyPublicationGuardModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValCTenantPrivacyPublicationGuardState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCCorrectionRedactionLimitationGuardState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*CorrectionRedactionLimitationGuard)
		want   string
	}{
		{"redaction limitation active", func(model *CorrectionRedactionLimitationGuard) {}, Point14ValCStateActive},
		{"missing redaction manifest when redacted blocks", func(model *CorrectionRedactionLimitationGuard) { model.RedactionManifestRef = "" }, Point14ValCStateBlocked},
		{"missing limitation refs blocks", func(model *CorrectionRedactionLimitationGuard) { model.LimitationRefs = nil }, Point14ValCStateBlocked},
		{"decisive missing evidence hidden blocks", func(model *CorrectionRedactionLimitationGuard) { model.DecisiveMissingEvidenceHidden = true }, Point14ValCStateBlocked},
		{"redaction strengthens claim blocks", func(model *CorrectionRedactionLimitationGuard) { model.RedactionStrengthensClaim = true }, Point14ValCStateBlocked},
		{"limitation omitted blocks", func(model *CorrectionRedactionLimitationGuard) { model.LimitationOmitted = true }, Point14ValCStateBlocked},
		{"surviving text misleading blocks", func(model *CorrectionRedactionLimitationGuard) { model.SurvivingTextMisleading = true }, Point14ValCStateBlocked},
		{"redaction cannot convert incomplete review required into active", func(model *CorrectionRedactionLimitationGuard) {
			model.SourcePublicationState = Point14ValCStateIncomplete
			model.ResolvedPublicationState = Point14ValCStateActive
		}, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCCorrectionRedactionLimitationGuardModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValCCorrectionRedactionLimitationGuardState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCGovernanceTraceState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*CorrectionPublicationGovernanceTrace)
		want   string
	}{
		{"valid governance trace active", func(model *CorrectionPublicationGovernanceTrace) {}, Point14ValCStateActive},
		{"missing governance event ref blocks", func(model *CorrectionPublicationGovernanceTrace) { model.GovernanceEventRef = "" }, Point14ValCStateBlocked},
		{"missing owner blocks", func(model *CorrectionPublicationGovernanceTrace) { model.Owner = "" }, Point14ValCStateBlocked},
		{"missing approver role blocks", func(model *CorrectionPublicationGovernanceTrace) { model.ApproverRole = "" }, Point14ValCStateBlocked},
		{"missing audit ref blocks", func(model *CorrectionPublicationGovernanceTrace) { model.AuditRef = "" }, Point14ValCStateBlocked},
		{"missing evidence refs blocks", func(model *CorrectionPublicationGovernanceTrace) { model.EvidenceRefs = nil }, Point14ValCStateBlocked},
		{"missing decision reason blocks", func(model *CorrectionPublicationGovernanceTrace) { model.DecisionReason = "" }, Point14ValCStateBlocked},
		{"missing timestamp blocks", func(model *CorrectionPublicationGovernanceTrace) { model.Timestamp = "" }, Point14ValCStateBlocked},
		{"client local time as canonical timestamp blocks", func(model *CorrectionPublicationGovernanceTrace) { model.TimeSource = point14Val0TimeSourceClientLocal }, Point14ValCStateBlocked},
		{"backdated correction approval before dispute open review required", func(model *CorrectionPublicationGovernanceTrace) { model.Timestamp = "2026-05-06T07:59:00Z" }, Point14ValCStateReviewRequired},
		{"governance trace cannot approve production or certify compliance", func(model *CorrectionPublicationGovernanceTrace) { model.ApprovesProduction = true }, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCGovernanceTraceModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValCGovernanceTraceState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCAgentCorrectionPublicationBoundaryState(t *testing.T) {
	dependency := point14ValCDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*AgentCorrectionPublicationBoundary)
		want   string
	}{
		{"agent recommendation may be advisory input", func(model *AgentCorrectionPublicationBoundary) {}, Point14ValCStateActive},
		{"agent cannot approve correction", func(model *AgentCorrectionPublicationBoundary) { model.CanApproveCorrection = true }, Point14ValCStateBlocked},
		{"agent cannot approve revocation", func(model *AgentCorrectionPublicationBoundary) { model.CanApproveRevocation = true }, Point14ValCStateBlocked},
		{"agent cannot approve publication", func(model *AgentCorrectionPublicationBoundary) { model.CanApprovePublication = true }, Point14ValCStateBlocked},
		{"agent cannot publish notice", func(model *AgentCorrectionPublicationBoundary) { model.CanPublishNotice = true }, Point14ValCStateBlocked},
		{"agent cannot satisfy governance trace alone", func(model *AgentCorrectionPublicationBoundary) { model.CanSatisfyGovernanceTraceAlone = true }, Point14ValCStateBlocked},
		{"ai agent authority flags block globally", func(model *AgentCorrectionPublicationBoundary) { model.ExternalAuthorityAllowed = true }, Point14ValCStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValCAgentCorrectionPublicationBoundaryModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValCAgentCorrectionPublicationBoundaryState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValCAuthorityAndWordingGuards(t *testing.T) {
	t.Run("correction auto published marker blocks", func(t *testing.T) {
		model := point14ValCNoExternalAuthorityPublicationGuardModel()
		model.ObservedAuthorityMarkers = []string{"correction_auto_published"}
		if got := EvaluatePoint14ValCNoExternalAuthorityPublicationGuardState(model); got != Point14ValCStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("agent approved publication blocks", func(t *testing.T) {
		model := point14ValCNoExternalAuthorityPublicationGuardModel()
		model.AgentApprovedPublication = true
		if got := EvaluatePoint14ValCNoExternalAuthorityPublicationGuardState(model); got != Point14ValCStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("forbidden publication wording blocks", func(t *testing.T) {
		model := point14ValCNoOverclaimPublicationWordingModel()
		model.ObservedPublicationTexts = []string{"publication proves safety"}
		if got := EvaluatePoint14ValCNoOverclaimPublicationWordingState(model); got != Point14ValCStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("safe bounded wording passes", func(t *testing.T) {
		model := point14ValCNoOverclaimPublicationWordingModel()
		if got := EvaluatePoint14ValCNoOverclaimPublicationWordingState(model); got != Point14ValCStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("forbidden wording allowed only in internal blocked diagnostic context", func(t *testing.T) {
		model := point14ValCNoOverclaimPublicationWordingModel()
		model.InternalDiagnosticTexts = []string{"scanner PASS"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint14ValCNoOverclaimPublicationWordingState(model); got != Point14ValCStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
}

func TestPoint14ValCFoundationAggregation(t *testing.T) {
	t.Run("any blocked component yields blocked", func(t *testing.T) {
		model := Point14ValCFoundationModel()
		model.CorrectionNoticeBoundary.CanonicalMutationRequested = true
		got := ComputePoint14ValCFoundation(model)
		if got.CurrentState != Point14ValCStateBlocked {
			t.Fatalf("expected blocked, got %#v", got)
		}
	})
	t.Run("any review required and no blocked yields review required", func(t *testing.T) {
		model := Point14ValCFoundationModel()
		model.PublicationApprovalBoundary.RecordedAt = "2026-05-06T08:04:00Z"
		got := ComputePoint14ValCFoundation(model)
		if got.CurrentState != Point14ValCStateReviewRequired {
			t.Fatalf("expected review required, got %#v", got)
		}
	})
	t.Run("incomplete only when no blocked review required exists", func(t *testing.T) {
		model := Point14ValCFoundationModel()
		model.CorrectionNoticeBoundary.CorrectionState = point14ValCCorrectionEvidenceRequired
		got := ComputePoint14ValCFoundation(model)
		if got.CurrentState != Point14ValCStateIncomplete {
			t.Fatalf("expected incomplete, got %#v", got)
		}
	})
	t.Run("active only when all components active", func(t *testing.T) {
		model := ComputePoint14ValCFoundation(Point14ValCFoundationModel())
		if model.CurrentState != Point14ValCStateActive {
			t.Fatalf("expected active, got %#v", model)
		}
	})
}
