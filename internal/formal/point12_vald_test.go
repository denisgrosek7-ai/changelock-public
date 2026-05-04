package formal

import "testing"

func activePoint12ValDFoundation() Point12ValDFoundation {
	return ComputePoint12ValDFoundation(Point12ValDFoundationModel())
}

func activePoint12ValDComparisonFoundation() Point12ValDFoundation {
	model := Point12ValDFoundationModel()
	model.Dependency.ValBReplayRequest.ReplayMode = point12Val0ReplayModeComparisonMode
	model.Dependency.ValBReplayRequest.CurrentPolicyRef = "policy_ref_point12_vald_current"
	model.Dependency.ValBReplayRequest.CurrentPolicyVersion = "policy_version_point12_vald_current"
	model.Dependency.ValBReplayRequest.CurrentPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	model.Dependency.ValBReplayRequest.CurrentEngineVersion = "engine_version_point12_vald_current"
	model.Dependency.ValBReplayRequest.CurrentEngineHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
	model.Dependency.ValBReplayRequest.CurrentSchemaVersion = "schema_version_point12_vald_current"
	model.Dependency.ValBReplayRequest.CurrentSchemaHash = "sha256:cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
	model.Dependency.ValBReplayRequest.CurrentEvidenceRefs = []string{"evidence_ref_point12_vald_current"}
	model.Dependency.ValBReplayRequest.CurrentEvidenceHashRefs = []string{"sha256:dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"}
	model.Dependency.ValBReplayRequest.CurrentClaimRefs = []string{"claim_ref_point12_vald_current"}
	model.Dependency.ValBReplayRequest.CurrentGovernanceEventRefs = []string{"governance_event_point12_vald_current"}
	model.Dependency.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
	model.Dependency.ValBReplayResult.DecisionDriftClassification = point12ValBDriftDueToPolicy
	model.Dependency.ValBReplayResult.DecisionDriftExplanation = "policy and context drift caused a different decision"
	model.Query.QueryKind = point12ValDQueryKindWhyChanged
	model.Explanation.ExplanationKind = point12ValDQueryKindWhyChanged
	model.Explanation.WhyChangedSummary = "The decision changed because policy, engine, schema, evidence, claim, governance, and tenant scope drifted."
	model.Explanation.DriftReasons = []string{point12ValBDriftDueToPolicy}
	model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
	return ComputePoint12ValDFoundation(model)
}

func recomputePoint12ValDLocalHashes(model *Point12ValDFoundation) {
	model.ProofChain.ProjectionHash = point12ValDComputedProjectionHash(model.ProofChain)
	model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
	model.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.SupportProfile)
}

func addPoint12ValDSecondEvidencePair(model *Point12ValDFoundation, includeCanonicalEdges bool) {
	evidenceRef := "evidence_ref_point12_vald_002"
	evidenceHash := "evidence_hash_point12_vald_002"
	model.Dependency.ValCAuditExportBundle.EvidenceRefs = append(model.Dependency.ValCAuditExportBundle.EvidenceRefs, evidenceRef)
	model.Dependency.ValCAuditExportBundle.EvidenceHashRefs = append(model.Dependency.ValCAuditExportBundle.EvidenceHashRefs, evidenceHash)
	model.Dependency.ValCOfflineBundle.EvidenceRefs = append(model.Dependency.ValCOfflineBundle.EvidenceRefs, evidenceRef)
	model.Dependency.ValCOfflineBundle.EvidenceHashRefs = append(model.Dependency.ValCOfflineBundle.EvidenceHashRefs, evidenceHash)
	model.ProofChain.EvidenceRefs = append(model.ProofChain.EvidenceRefs, evidenceRef)
	model.ProofChain.EvidenceHashRefs = append(model.ProofChain.EvidenceHashRefs, evidenceHash)
	model.ProofChain.SourceEvidenceSpineRefs = append(model.ProofChain.SourceEvidenceSpineRefs, evidenceRef)
	model.Explanation.BasedOnRefs = point12ValDExpectedExplanationRefs(model.ProofChain)
	model.Explanation.BasedOnHashes = point12ValDExpectedExplanationHashes(model.ProofChain)
	model.SupportProfile.SupportingEvidenceRefs = append(model.SupportProfile.SupportingEvidenceRefs, evidenceRef)
	model.SupportProfile.SupportingEvidenceHashRefs = append(model.SupportProfile.SupportingEvidenceHashRefs, evidenceHash)
	if includeCanonicalEdges {
		model.ProofChain.LineageEdges = append(model.ProofChain.LineageEdges,
			Point12ValDLineageEdge{
				EdgeID:           "lineage_edge_point12_vald_source_002",
				EdgeType:         point12ValDLineageEdgeTypeSourceToEvidence,
				FromRef:          "source_point12_vald_002",
				ToRef:            evidenceRef,
				FromHash:         evidenceHash,
				ToHash:           evidenceHash,
				TenantScope:      model.ProofChain.TenantScope,
				EvidenceSpineRef: evidenceRef,
				SourceTimestamp:  "2026-05-04T08:00:22Z",
				TargetTimestamp:  "2026-05-04T08:00:23Z",
				AdvisoryOnly:     true,
				EdgeState:        Point12ValDLineageEdgeStateActive,
				Explanation:      "second source evidence captured in canonical spine",
			},
			Point12ValDLineageEdge{
				EdgeID:           "lineage_edge_point12_vald_artifact_002",
				EdgeType:         point12ValDLineageEdgeTypeEvidenceToArtifact,
				FromRef:          evidenceRef,
				ToRef:            model.ProofChain.ArtifactRef,
				FromHash:         evidenceHash,
				ToHash:           model.ProofChain.ArtifactHash,
				TenantScope:      model.ProofChain.TenantScope,
				EvidenceSpineRef: evidenceRef,
				SourceTimestamp:  "2026-05-04T08:00:24Z",
				TargetTimestamp:  "2026-05-04T08:00:25Z",
				AdvisoryOnly:     true,
				EdgeState:        Point12ValDLineageEdgeStateActive,
				Explanation:      "second evidence bound to artifact hash",
			},
		)
	}
}

func TestPoint12ValDPreflightPrerequisites(t *testing.T) {
	t.Run("vala schema hash drift remains blocked after local recomputation", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.SchemaHash = "sha256:abababababababababababababababababababababababababababababababab"
		model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
		model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState == Point12ValAManifestIntegrityStateActive {
			t.Fatalf("expected schema hash drift to block/tamper vala manifest integrity, got %#v", model)
		}
	})

	t.Run("valc export exact binding mutation remains not export ready", func(t *testing.T) {
		model := Point12ValCFoundationModel()
		model.ExportBundle.PolicyHash = "sha256:abababababababababababababababababababababababababababababababab"
		model.ExportBundle.ExportState = Point12ValCExportStateReady
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState == Point12ValCExportStateReady || model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected mutated valc export binding to avoid full-ready state, got %#v", model)
		}
	})

	t.Run("valc offline exact binding mutation remains not active", func(t *testing.T) {
		model := Point12ValCFoundationModel()
		model.OfflineBundle.ManifestPayloadHash = "sha256:cdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcd"
		model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		model = ComputePoint12ValCFoundation(model)
		if model.OfflineBundleState == Point12ValCOfflineBundleStateActive && model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected mutated valc offline binding to avoid active/full-ready state, got %#v", model)
		}
	})
}

func TestPoint12ValDDependencyState(t *testing.T) {
	t.Run("valid computed valc output remains active", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		if model.DependencyState != Point12ValDDependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", model)
		}
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValDDependencySnapshot)
		want   string
	}{
		{name: "missing computed provenance blocks", mutate: func(model *Point12ValDDependencySnapshot) { model.SnapshotFromComputedOutput = false }, want: Point12ValDDependencyStateBlocked},
		{name: "partial advisory export requires review", mutate: func(model *Point12ValDDependencySnapshot) {
			model.ValCExportState = Point12ValCExportStatePartialAdvisory
		}, want: Point12ValDDependencyStateReviewRequired},
		{name: "tenant scope mismatch export blocks", mutate: func(model *Point12ValDDependencySnapshot) {
			model.ValCExportState = Point12ValCExportStateTenantMismatch
		}, want: Point12ValDDependencyStateBlocked},
		{name: "public private boundary violation blocks", mutate: func(model *Point12ValDDependencySnapshot) {
			model.ValCPublicPrivateBoundaryState = Point12ValCPublicPrivateBoundaryStateBlocked
		}, want: Point12ValDDependencyStateBlocked},
		{name: "offline external api blocks", mutate: func(model *Point12ValDDependencySnapshot) { model.ValCExternalAPIUsed = true }, want: Point12ValDDependencyStateBlocked},
		{name: "premature point12 pass blocks", mutate: func(model *Point12ValDDependencySnapshot) { model.ValCPrematurePoint12PassSeen = true }, want: Point12ValDDependencyStateBlocked},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint12ValDFoundation().Dependency
			tc.mutate(&model)
			if got := EvaluatePoint12ValDDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s for %#v", tc.want, got, model)
			}
		})
	}
}

func TestPoint12ValDBindingMatrixState(t *testing.T) {
	t.Run("binding matrix covers downstream models", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		requiredModels := map[string]bool{
			"Point12ValDDependencySnapshot":                       false,
			"Point12ValDProofChainProjection":                     false,
			"Point12ValDLineageEdge":                              false,
			"Point12ValDProofChainQuery":                          false,
			"Point12ValDExplanationResult":                        false,
			"Point12ValDFinancialInsuranceEvidenceSupportProfile": false,
			"Point12ValDPortalCompatibilityContract":              false,
		}
		for _, entry := range model.BindingMatrix.BoundFields {
			requiredModels[entry.DownstreamModel] = true
		}
		for modelName, seen := range requiredModels {
			if !seen {
				t.Fatalf("expected binding matrix coverage for %s in %#v", modelName, model.BindingMatrix)
			}
		}
		if model.BindingMatrixState != Point12ValDBindingMatrixStateActive {
			t.Fatalf("expected active binding matrix state, got %#v", model)
		}
	})

	t.Run("exact required fields require validators", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.BindingMatrix.BoundFields[0].ValidationRequired = false
		if got := EvaluatePoint12ValDBindingMatrixState(model.BindingMatrix); got != Point12ValDBindingMatrixStateBlocked {
			t.Fatalf("expected blocked binding matrix without validators, got %#v", model.BindingMatrix)
		}
	})

	t.Run("intentionally not bound requires reason", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		for i := range model.BindingMatrix.BoundFields {
			if model.BindingMatrix.BoundFields[i].BindingClass == point12ValDBindingClassIntentionallyNotBound {
				model.BindingMatrix.BoundFields[i].Reason = ""
				break
			}
		}
		if got := EvaluatePoint12ValDBindingMatrixState(model.BindingMatrix); got != Point12ValDBindingMatrixStateReviewRequired {
			t.Fatalf("expected review required when intentionally-not-bound reason missing, got %#v", model.BindingMatrix)
		}
	})

	t.Run("missing downstream model entry requires review", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		filtered := model.BindingMatrix.BoundFields[:0]
		for _, entry := range model.BindingMatrix.BoundFields {
			if entry.DownstreamModel != "Point12ValDPortalCompatibilityContract" {
				filtered = append(filtered, entry)
			}
		}
		model.BindingMatrix.BoundFields = filtered
		if got := EvaluatePoint12ValDBindingMatrixState(model.BindingMatrix); got != Point12ValDBindingMatrixStateReviewRequired {
			t.Fatalf("expected review required when model coverage missing, got %#v", model.BindingMatrix)
		}
	})

	t.Run("non empty exact binding without upstream source is insufficient", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.BindingMatrix.BoundFields[0].UpstreamSource = ""
		if got := EvaluatePoint12ValDBindingMatrixState(model.BindingMatrix); got != Point12ValDBindingMatrixStateBlocked {
			t.Fatalf("expected blocked binding matrix without upstream source, got %#v", model.BindingMatrix)
		}
	})
}

func TestPoint12ValDLineageEdgeState(t *testing.T) {
	base := activePoint12ValDFoundation()
	t.Run("valid source to evidence edge active", func(t *testing.T) {
		edge := base.ProofChain.LineageEdges[0]
		if got := EvaluatePoint12ValDLineageEdgeState(edge, base.ProofChain); got != Point12ValDLineageEdgeStateActive {
			t.Fatalf("expected active edge, got %#v", edge)
		}
	})

	t.Run("valid evidence to artifact edge active", func(t *testing.T) {
		edge := base.ProofChain.LineageEdges[1]
		if got := EvaluatePoint12ValDLineageEdgeState(edge, base.ProofChain); got != Point12ValDLineageEdgeStateActive {
			t.Fatalf("expected active edge, got %#v", edge)
		}
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValDLineageEdge)
		want   string
	}{
		{name: "missing from ref blocks", mutate: func(model *Point12ValDLineageEdge) { model.FromRef = "" }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "missing to ref blocks", mutate: func(model *Point12ValDLineageEdge) { model.ToRef = "" }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "missing hash blocks", mutate: func(model *Point12ValDLineageEdge) { model.FromHash = "" }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "query state rejected as lineage edge state", mutate: func(model *Point12ValDLineageEdge) { model.EdgeState = Point12ValDQueryStateActive }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "explanation state rejected as lineage edge state", mutate: func(model *Point12ValDLineageEdge) { model.EdgeState = Point12ValDExplanationStateActive }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "support profile state rejected as lineage edge state", mutate: func(model *Point12ValDLineageEdge) { model.EdgeState = Point12ValDSupportProfileStateActive }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "portal compatibility state rejected as lineage edge state", mutate: func(model *Point12ValDLineageEdge) { model.EdgeState = Point12ValDPortalCompatibilityStateActive }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "canonical looking junk state rejected as lineage edge state", mutate: func(model *Point12ValDLineageEdge) { model.EdgeState = "point12_vald_lineage_edge_active_but_wrong" }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "inferred decisive edge blocks", mutate: func(model *Point12ValDLineageEdge) { model.Inferred = true; model.Decisive = true }, want: Point12ValDLineageEdgeStateBlocked},
		{name: "advisory inferred non decisive edge review", mutate: func(model *Point12ValDLineageEdge) {
			model.Inferred = true
			model.Decisive = false
			model.AdvisoryOnly = false
		}, want: Point12ValDLineageEdgeStateReviewRequired},
		{name: "cross tenant edge blocks", mutate: func(model *Point12ValDLineageEdge) { model.TenantScope = "tenant_scope_other" }, want: Point12ValDLineageEdgeStateBlocked},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			edge := base.ProofChain.LineageEdges[0]
			tc.mutate(&edge)
			if got := EvaluatePoint12ValDLineageEdgeState(edge, base.ProofChain); got != tc.want {
				t.Fatalf("expected %s, got %s for %#v", tc.want, got, edge)
			}
		})
	}

	t.Run("agent advisory edge cannot certify or emit pass", func(t *testing.T) {
		edge := Point12ValDLineageEdge{
			EdgeID:                 "lineage_edge_point12_vald_agent_001",
			EdgeType:               point12ValDLineageEdgeTypeAgentFindingAdvisory,
			FromRef:                "agent_lineage_point12_vald_001",
			ToRef:                  base.ProofChain.ArtifactRef,
			FromHash:               base.ProofChain.ArtifactHash,
			ToHash:                 base.ProofChain.ArtifactHash,
			TenantScope:            base.ProofChain.TenantScope,
			EvidenceSpineRef:       base.ProofChain.EvidenceRefs[0],
			SourceTimestamp:        "2026-05-04T08:20:00Z",
			TargetTimestamp:        "2026-05-04T08:20:01Z",
			AdvisoryOnly:           true,
			EdgeState:              Point12ValDLineageEdgeStateActive,
			AgentID:                "agent_lineage_point12_val0_001",
			AgentType:              "analysis_recommendation",
			PermissionManifestHash: "sha256:6666666666666666666666666666666666666666666666666666666666666666",
			InputEvidenceRefs:      []string{base.ProofChain.EvidenceRefs[0]},
			AuditID:                "audit_point12_vald_agent_001",
			RecommendationID:       "recommendation_point12_vald_001",
			LineageInputOnly:       true,
			ClaimsCertification:    true,
			EmitsPrematurePass:     true,
		}
		if got := EvaluatePoint12ValDLineageEdgeState(edge, base.ProofChain); got != Point12ValDLineageEdgeStateBlocked {
			t.Fatalf("expected blocked agent advisory edge with certification/pass claims, got %#v", edge)
		}
	})
}

func TestPoint12ValDProofChainProjection(t *testing.T) {
	t.Run("valid proof chain projection active", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		if model.ProofChainState != Point12ValDProofChainStateActive {
			t.Fatalf("expected active proof chain state, got %#v", model)
		}
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValDFoundation)
	}{
		{name: "missing evidence lineage edge blocks", mutate: func(model *Point12ValDFoundation) {
			filtered := model.ProofChain.LineageEdges[:0]
			for _, edge := range model.ProofChain.LineageEdges {
				if edge.EdgeType != point12ValDLineageEdgeTypeSourceToEvidence {
					filtered = append(filtered, edge)
				}
			}
			model.ProofChain.LineageEdges = filtered
		}},
		{name: "missing source evidence spine ref blocks", mutate: func(model *Point12ValDFoundation) { model.ProofChain.SourceEvidenceSpineRefs = nil }},
		{name: "inferred decisive edge blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[0].Inferred = true
			model.ProofChain.LineageEdges[0].Decisive = true
		}},
		{name: "cross tenant edge blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[0].TenantScope = "tenant_scope_other"
		}},
		{name: "cross namespace lineage edge state blocks proof chain", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[0].EdgeState = Point12ValDQueryStateActive
		}},
		{name: "source to evidence wrong to ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[0].ToRef = "evidence_ref_point12_vald_wrong"
		}},
		{name: "evidence to artifact wrong from ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[1].FromRef = "evidence_ref_point12_vald_wrong"
		}},
		{name: "evidence to artifact wrong from hash blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[1].FromHash = "sha256:dededededededededededededededededededededededededededededededede"
		}},
		{name: "artifact to decision wrong to ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[2].ToRef = "decision_ref_point12_vald_wrong"
		}},
		{name: "decision to proof pack wrong from ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[3].FromRef = "decision_ref_point12_vald_wrong"
		}},
		{name: "proof pack to manifest wrong to ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[4].ToRef = "manifest_ref_point12_vald_wrong"
		}},
		{name: "manifest to replay wrong to ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[5].ToRef = "replay_result_point12_vald_wrong"
		}},
		{name: "replay to export wrong to ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[6].ToRef = "export_point12_vald_wrong"
		}},
		{name: "export to offline bundle wrong to ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[7].ToRef = "offline_bundle_point12_vald_wrong"
		}},
		{name: "redaction to export wrong from ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[8].FromRef = "redaction_manifest_point12_vald_wrong"
		}},
		{name: "claim to decision wrong claim ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[9].FromRef = "claim_ref_point12_vald_wrong"
		}},
		{name: "governance to decision wrong governance ref blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[10].FromRef = "governance_event_point12_vald_wrong"
		}},
		{name: "evidence to artifact swapped endpoints block", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[1].FromRef = model.ProofChain.ArtifactRef
			model.ProofChain.LineageEdges[1].ToRef = model.ProofChain.EvidenceRefs[0]
		}},
		{name: "replay to export swapped endpoints block", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[6].FromRef = model.ProofChain.ExportID
			model.ProofChain.LineageEdges[6].ToRef = model.ProofChain.ReplayResultID
		}},
		{name: "redaction to export swapped endpoints block", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[8].FromRef = model.ProofChain.ExportID
			model.ProofChain.LineageEdges[8].ToRef = model.ProofChain.RedactionManifestID
		}},
		{name: "agent advisory edge cannot satisfy source to evidence", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.LineageEdges[0].EdgeType = point12ValDLineageEdgeTypeAgentFindingAdvisory
			model.ProofChain.LineageEdges[0].AgentID = "agent_lineage_point12_val0_001"
			model.ProofChain.LineageEdges[0].AgentType = "analysis_recommendation"
			model.ProofChain.LineageEdges[0].PermissionManifestHash = "sha256:6666666666666666666666666666666666666666666666666666666666666666"
			model.ProofChain.LineageEdges[0].InputEvidenceRefs = []string{model.ProofChain.EvidenceRefs[0]}
			model.ProofChain.LineageEdges[0].AuditID = "audit_point12_vald_agent_001"
			model.ProofChain.LineageEdges[0].RecommendationID = "recommendation_point12_vald_001"
			model.ProofChain.LineageEdges[0].LineageInputOnly = true
			model.ProofChain.LineageEdges[0].AdvisoryOnly = true
		}},
		{name: "artifact hash mismatch blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.ArtifactHash = "sha256:abababababababababababababababababababababababababababababababab"
		}},
		{name: "evidence hash mismatch blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.EvidenceHashRefs = []string{"sha256:cdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcd"}
		}},
		{name: "policy hash mismatch blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.PolicyHash = "sha256:eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		}},
		{name: "engine hash mismatch blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.EngineHash = "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
		}},
		{name: "schema hash mismatch blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.SchemaHash = "sha256:1111111111111111111111111111111111111111111111111111111111111111"
		}},
		{name: "claim refs mismatch blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.ClaimRefs = []string{"claim_ref_point12_vald_999"}
		}},
		{name: "governance refs mismatch blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.GovernanceEventRefs = []string{"governance_event_point12_vald_999"}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			model := Point12ValDFoundationModel()
			tc.mutate(&model)
			recomputePoint12ValDLocalHashes(&model)
			model = ComputePoint12ValDFoundation(model)
			if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
				t.Fatalf("expected mutated proof chain to avoid active state, got %#v", model)
			}
		})
	}

	t.Run("projection hash recomputation cannot hide upstream drift", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.ProofChain.SchemaHash = "sha256:2222222222222222222222222222222222222222222222222222222222222222"
		model.ProofChain.ProjectionHash = point12ValDComputedProjectionHash(model.ProofChain)
		model = ComputePoint12ValDFoundation(model)
		if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected local projection hash recomputation to fail closed, got %#v", model)
		}
	})

	t.Run("projection hash recomputation cannot hide wrong lineage endpoint", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.ProofChain.LineageEdges[6].ToRef = "export_point12_vald_wrong"
		model.ProofChain.ProjectionHash = point12ValDComputedProjectionHash(model.ProofChain)
		model = ComputePoint12ValDFoundation(model)
		if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected wrong lineage endpoint to stay blocked after projection hash recomputation, got %#v", model)
		}
	})

	t.Run("projection hash recomputation cannot hide wrong lineage hash", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.ProofChain.LineageEdges[1].FromHash = "sha256:3434343434343434343434343434343434343434343434343434343434343434"
		model.ProofChain.ProjectionHash = point12ValDComputedProjectionHash(model.ProofChain)
		model = ComputePoint12ValDFoundation(model)
		if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected wrong lineage hash to stay blocked after projection hash recomputation, got %#v", model)
		}
	})

	t.Run("multi evidence exact lineage bindings stay active when all expected edges are present", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		addPoint12ValDSecondEvidencePair(&model, true)
		recomputePoint12ValDLocalHashes(&model)
		model = ComputePoint12ValDFoundation(model)
		if model.ProofChainState != Point12ValDProofChainStateActive || model.CurrentState != Point12ValDStateActive {
			t.Fatalf("expected multi evidence proof chain to stay active with exact lineage coverage, got %#v", model)
		}
	})

	t.Run("missing one evidence edge from required evidence refs blocks", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		addPoint12ValDSecondEvidencePair(&model, false)
		recomputePoint12ValDLocalHashes(&model)
		model = ComputePoint12ValDFoundation(model)
		if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected missing second evidence edge coverage to fail closed, got %#v", model)
		}
	})

	t.Run("duplicate evidence edge cannot satisfy another required evidence ref", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		addPoint12ValDSecondEvidencePair(&model, false)
		model.ProofChain.LineageEdges = append(model.ProofChain.LineageEdges,
			Point12ValDLineageEdge{
				EdgeID:           "lineage_edge_point12_vald_source_duplicate_002",
				EdgeType:         point12ValDLineageEdgeTypeSourceToEvidence,
				FromRef:          "source_point12_vald_duplicate_002",
				ToRef:            model.Dependency.ValCAuditExportBundle.EvidenceRefs[0],
				FromHash:         model.Dependency.ValCAuditExportBundle.EvidenceHashRefs[0],
				ToHash:           model.Dependency.ValCAuditExportBundle.EvidenceHashRefs[0],
				TenantScope:      model.ProofChain.TenantScope,
				EvidenceSpineRef: model.Dependency.ValCAuditExportBundle.EvidenceRefs[0],
				SourceTimestamp:  "2026-05-04T08:00:26Z",
				TargetTimestamp:  "2026-05-04T08:00:27Z",
				AdvisoryOnly:     true,
				EdgeState:        Point12ValDLineageEdgeStateActive,
				Explanation:      "duplicate source edge cannot satisfy second evidence ref",
			},
			Point12ValDLineageEdge{
				EdgeID:           "lineage_edge_point12_vald_artifact_duplicate_002",
				EdgeType:         point12ValDLineageEdgeTypeEvidenceToArtifact,
				FromRef:          model.Dependency.ValCAuditExportBundle.EvidenceRefs[0],
				ToRef:            model.ProofChain.ArtifactRef,
				FromHash:         model.Dependency.ValCAuditExportBundle.EvidenceHashRefs[0],
				ToHash:           model.ProofChain.ArtifactHash,
				TenantScope:      model.ProofChain.TenantScope,
				EvidenceSpineRef: model.Dependency.ValCAuditExportBundle.EvidenceRefs[0],
				SourceTimestamp:  "2026-05-04T08:00:28Z",
				TargetTimestamp:  "2026-05-04T08:00:29Z",
				AdvisoryOnly:     true,
				EdgeState:        Point12ValDLineageEdgeStateActive,
				Explanation:      "duplicate artifact edge cannot satisfy second evidence ref",
			},
		)
		recomputePoint12ValDLocalHashes(&model)
		model = ComputePoint12ValDFoundation(model)
		if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected duplicate evidence edges to fail closed for missing second evidence binding, got %#v", model)
		}
	})

	t.Run("unrelated evidence edge with correct edge type cannot satisfy required evidence binding", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		addPoint12ValDSecondEvidencePair(&model, true)
		model.ProofChain.LineageEdges = append(model.ProofChain.LineageEdges, Point12ValDLineageEdge{
			EdgeID:           "lineage_edge_point12_vald_artifact_unrelated_003",
			EdgeType:         point12ValDLineageEdgeTypeEvidenceToArtifact,
			FromRef:          "evidence_ref_point12_vald_unrelated",
			ToRef:            model.ProofChain.ArtifactRef,
			FromHash:         "sha256:5656565656565656565656565656565656565656565656565656565656565656",
			ToHash:           model.ProofChain.ArtifactHash,
			TenantScope:      model.ProofChain.TenantScope,
			EvidenceSpineRef: "evidence_ref_point12_vald_unrelated",
			SourceTimestamp:  "2026-05-04T08:00:30Z",
			TargetTimestamp:  "2026-05-04T08:00:31Z",
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "unrelated edge cannot satisfy canonical evidence binding",
		})
		recomputePoint12ValDLocalHashes(&model)
		model = ComputePoint12ValDFoundation(model)
		if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected unrelated evidence edge to fail closed or remain non-satisfying, got %#v", model)
		}
	})
}

func TestPoint12ValDProofChainQueryState(t *testing.T) {
	t.Run("why decision query active", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		if model.QueryState != Point12ValDQueryStateActive {
			t.Fatalf("expected active query state, got %#v", model)
		}
	})

	t.Run("why changed requires comparison context", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.Query.QueryKind = point12ValDQueryKindWhyChanged
		if got := EvaluatePoint12ValDProofChainQueryState(model.Query, model.ProofChain, model.Dependency); got != Point12ValDQueryStateReviewRequired {
			t.Fatalf("expected review required why_changed query without comparison context, got %#v", model.Query)
		}
	})

	t.Run("why changed active with comparison context", func(t *testing.T) {
		model := activePoint12ValDComparisonFoundation()
		if model.QueryState != Point12ValDQueryStateActive {
			t.Fatalf("expected active why_changed query, got %#v", model)
		}
	})

	t.Run("explain mismatch requires mismatch details", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.Query.QueryKind = point12ValDQueryKindExplainMismatch
		model.Query.IncludeMismatchDetails = false
		if got := EvaluatePoint12ValDProofChainQueryState(model.Query, model.ProofChain, model.Dependency); got != Point12ValDQueryStateReviewRequired {
			t.Fatalf("expected review required mismatch query without details, got %#v", model.Query)
		}
	})

	t.Run("explain missing evidence requires missing evidence context", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.Query.QueryKind = point12ValDQueryKindExplainMissingEvidence
		if got := EvaluatePoint12ValDProofChainQueryState(model.Query, model.ProofChain, model.Dependency); got != Point12ValDQueryStateReviewRequired {
			t.Fatalf("expected review required missing-evidence query without missing-evidence context, got %#v", model.Query)
		}
	})

	t.Run("explain redaction limitations requires redaction flag", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.Query.QueryKind = point12ValDQueryKindExplainRedactionLimitations
		model.Query.IncludeRedactionLimitations = false
		if got := EvaluatePoint12ValDProofChainQueryState(model.Query, model.ProofChain, model.Dependency); got != Point12ValDQueryStateReviewRequired {
			t.Fatalf("expected review required redaction query without flag, got %#v", model.Query)
		}
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValDProofChainQuery)
	}{
		{name: "allow external api blocks", mutate: func(model *Point12ValDProofChainQuery) { model.AllowExternalAPI = true }},
		{name: "allow mutation blocks", mutate: func(model *Point12ValDProofChainQuery) { model.AllowMutation = true }},
		{name: "unknown query kind blocks", mutate: func(model *Point12ValDProofChainQuery) { model.QueryKind = "query_kind_unknown" }},
		{name: "query cannot emit point12 pass", mutate: func(model *Point12ValDProofChainQuery) { model.RequestedExplanation = "point_12_pass" }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint12ValDFoundation()
			tc.mutate(&model.Query)
			if got := EvaluatePoint12ValDProofChainQueryState(model.Query, model.ProofChain, model.Dependency); got != Point12ValDQueryStateBlocked {
				t.Fatalf("expected blocked query state, got %#v", model.Query)
			}
		})
	}
}

func TestPoint12ValDExplanationState(t *testing.T) {
	t.Run("why decision explanation active with context", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		if model.ExplanationState != Point12ValDExplanationStateActive {
			t.Fatalf("expected active explanation state, got %#v", model)
		}
	})

	t.Run("why decision with missing decisive evidence requires missing explanation", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.Dependency.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultInsufficientEvidence
		model.Explanation.MissingEvidenceExplanations = nil
		model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
		model = ComputePoint12ValDFoundation(model)
		if model.ExplanationState != Point12ValDExplanationStateReviewRequired {
			t.Fatalf("expected review required explanation for missing decisive evidence, got %#v", model)
		}
	})

	t.Run("why changed without drift reasons requires review", func(t *testing.T) {
		model := activePoint12ValDComparisonFoundation()
		model.Explanation.DriftReasons = nil
		model.Explanation.WhyChangedSummary = ""
		model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
		model = ComputePoint12ValDFoundation(model)
		if model.ExplanationState != Point12ValDExplanationStateReviewRequired {
			t.Fatalf("expected review required why_changed explanation without drift reasons, got %#v", model)
		}
	})

	t.Run("mismatch explanation missing expected vs actual blocks review", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.Query.QueryKind = point12ValDQueryKindExplainMismatch
		model.Explanation.ExplanationKind = point12ValDQueryKindExplainMismatch
		model.Explanation.ExpectedRefs = nil
		model.Explanation.ActualRefs = nil
		model.Explanation.ExpectedHashes = nil
		model.Explanation.ActualHashes = nil
		model.Explanation.MismatchExplanations = []string{"policy hash mismatch"}
		model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
		model = ComputePoint12ValDFoundation(model)
		if model.ExplanationState != Point12ValDExplanationStateReviewRequired {
			t.Fatalf("expected review required mismatch explanation without expected/actual, got %#v", model)
		}
	})

	t.Run("customer visible overclaim blocks", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.Explanation.CustomerVisibleStatement = "This is production approved."
		model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
		model = ComputePoint12ValDFoundation(model)
		if model.ExplanationState != Point12ValDExplanationStateBlocked {
			t.Fatalf("expected blocked overclaim explanation, got %#v", model)
		}
	})

	t.Run("internal diagnostic may mention removed forbidden claim", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.Explanation.InternalDiagnosticSummary = "internal diagnostic: removed production approved claim from blocked surface"
		model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
		model = ComputePoint12ValDFoundation(model)
		if model.ExplanationState != Point12ValDExplanationStateActive {
			t.Fatalf("expected active explanation with internal diagnostic wording, got %#v", model)
		}
	})
}

func TestPoint12ValDSupportProfileState(t *testing.T) {
	for _, profileType := range []string{point12Val0ProfileTypeFinancialReview, point12Val0ProfileTypeInsuranceReview, point12ValDProfileTypeAuditReview} {
		t.Run("valid "+profileType+" profile active", func(t *testing.T) {
			model := activePoint12ValDFoundation()
			model.SupportProfile.ProfileType = profileType
			model.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.SupportProfile)
			model = ComputePoint12ValDFoundation(model)
			if model.SupportProfileState != Point12ValDSupportProfileStateActive {
				t.Fatalf("expected active support profile, got %#v", model)
			}
		})
	}

	cases := []struct {
		name   string
		mutate func(*Point12ValDFinancialInsuranceEvidenceSupportProfile)
	}{
		{name: "required customer review false blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) { model.RequiredCustomerReview = false }},
		{name: "legal review false blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) {
			model.LegalReviewRequiredForExternalUse = false
		}},
		{name: "no premium guarantee false blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) { model.NoPremiumGuarantee = false }},
		{name: "no rating claim false blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) { model.NoRatingClaim = false }},
		{name: "no compliance guarantee false blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) { model.NoComplianceGuarantee = false }},
		{name: "no financial guarantee false blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) { model.NoFinancialGuarantee = false }},
		{name: "no legal protection guarantee false blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) {
			model.NoLegalProtectionGuarantee = false
		}},
		{name: "lowers insurance premium blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) {
			model.SupportStatement = "This lowers insurance premium."
		}},
		{name: "proves dora compliance blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) {
			model.SupportStatement = "This proves DORA compliance."
		}},
		{name: "increases credit rating blocks", mutate: func(model *Point12ValDFinancialInsuranceEvidenceSupportProfile) {
			model.SupportStatement = "This increases credit rating."
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint12ValDFoundation()
			tc.mutate(&model.SupportProfile)
			model.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.SupportProfile)
			model = ComputePoint12ValDFoundation(model)
			if model.SupportProfileState != Point12ValDSupportProfileStateBlocked {
				t.Fatalf("expected blocked support profile, got %#v", model)
			}
		})
	}

	t.Run("blocked wording refs may contain forbidden wording", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		model.SupportProfile.BlockedWordingRefs = []string{"production approved", "financial guarantee"}
		model.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.SupportProfile)
		model = ComputePoint12ValDFoundation(model)
		if model.SupportProfileState != Point12ValDSupportProfileStateActive {
			t.Fatalf("expected active support profile with blocked wording refs, got %#v", model)
		}
	})
}

func TestPoint12ValDPortalCompatibilityState(t *testing.T) {
	t.Run("valid portal compatibility contract active", func(t *testing.T) {
		model := activePoint12ValDFoundation()
		if model.PortalCompatibilityState != Point12ValDPortalCompatibilityStateActive {
			t.Fatalf("expected active portal compatibility state, got %#v", model)
		}
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValDPortalCompatibilityContract)
	}{
		{name: "read only false blocks", mutate: func(model *Point12ValDPortalCompatibilityContract) { model.ReadOnly = false }},
		{name: "notes annotation only false blocks", mutate: func(model *Point12ValDPortalCompatibilityContract) { model.NotesAnnotationOnly = false }},
		{name: "evidence mutation allowed blocks", mutate: func(model *Point12ValDPortalCompatibilityContract) { model.EvidenceMutationAllowed = true }},
		{name: "decision mutation allowed blocks", mutate: func(model *Point12ValDPortalCompatibilityContract) { model.DecisionMutationAllowed = true }},
		{name: "certification allowed blocks", mutate: func(model *Point12ValDPortalCompatibilityContract) { model.CertificationAllowed = true }},
		{name: "point pass allowed blocks", mutate: func(model *Point12ValDPortalCompatibilityContract) { model.PointPassAllowed = true }},
		{name: "missing projection disclaimer blocks", mutate: func(model *Point12ValDPortalCompatibilityContract) { model.RequiredProjectionDisclaimer = "" }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint12ValDFoundation()
			tc.mutate(&model.PortalCompatibility)
			model = ComputePoint12ValDFoundation(model)
			if model.PortalCompatibilityState != Point12ValDPortalCompatibilityStateBlocked {
				t.Fatalf("expected blocked portal contract, got %#v", model)
			}
		})
	}
}

func TestPoint12ValDMutationClosure(t *testing.T) {
	exportMutations := []struct {
		name   string
		mutate func(*Point12ValDFoundation)
	}{
		{name: "proof chain tenant scope mutation blocks", mutate: func(model *Point12ValDFoundation) { model.ProofChain.TenantScope = "tenant_scope_other" }},
		{name: "proof chain artifact hash mutation blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.ArtifactHash = "sha256:1010101010101010101010101010101010101010101010101010101010101010"
		}},
		{name: "proof chain evidence hash mutation blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.EvidenceHashRefs = []string{"sha256:2020202020202020202020202020202020202020202020202020202020202020"}
		}},
		{name: "proof chain policy hash mutation blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.PolicyHash = "sha256:3030303030303030303030303030303030303030303030303030303030303030"
		}},
		{name: "proof chain engine hash mutation blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.EngineHash = "sha256:4040404040404040404040404040404040404040404040404040404040404040"
		}},
		{name: "proof chain schema hash mutation blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.SchemaHash = "sha256:5050505050505050505050505050505050505050505050505050505050505050"
		}},
		{name: "proof chain manifest payload mutation blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.ManifestPayloadHash = "sha256:6060606060606060606060606060606060606060606060606060606060606060"
		}},
		{name: "proof chain redaction manifest mutation blocks", mutate: func(model *Point12ValDFoundation) {
			model.ProofChain.RedactionManifestID = "redaction_manifest_point12_vald_999"
		}},
	}
	for _, tc := range exportMutations {
		t.Run(tc.name, func(t *testing.T) {
			model := Point12ValDFoundationModel()
			tc.mutate(&model)
			recomputePoint12ValDLocalHashes(&model)
			model = ComputePoint12ValDFoundation(model)
			if model.ProofChainState == Point12ValDProofChainStateActive || model.CurrentState == Point12ValDStateActive {
				t.Fatalf("expected mutation to fail closed, got %#v", model)
			}
		})
	}

	t.Run("explanation based on hashes mutation with recomputed hash blocks", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.Explanation.BasedOnHashes = []string{"sha256:7070707070707070707070707070707070707070707070707070707070707070"}
		model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
		model = ComputePoint12ValDFoundation(model)
		if model.ExplanationState == Point12ValDExplanationStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected explanation hash recomputation to fail closed, got %#v", model)
		}
	})

	t.Run("support profile evidence hash mutation with recomputed hash blocks", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.SupportProfile.SupportingEvidenceHashRefs = []string{"sha256:8080808080808080808080808080808080808080808080808080808080808080"}
		model.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.SupportProfile)
		model = ComputePoint12ValDFoundation(model)
		if model.SupportProfileState == Point12ValDSupportProfileStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected support profile hash recomputation to fail closed, got %#v", model)
		}
	})

	t.Run("portal contract binding mutation blocks", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.PortalCompatibility.ProofPackID = "proof_pack_point12_vald_999"
		model.PortalCompatibility.ReplayResultID = "replay_result_point12_vald_999"
		model.PortalCompatibility.ExportID = "export_point12_vald_999"
		model = ComputePoint12ValDFoundation(model)
		if model.PortalCompatibilityState == Point12ValDPortalCompatibilityStateActive || model.CurrentState == Point12ValDStateActive {
			t.Fatalf("expected portal contract drift to fail closed, got %#v", model)
		}
	})
}

func TestPoint12ValDRegressionGuards(t *testing.T) {
	t.Run("valc unsupported dependency remains review required and not export ready", func(t *testing.T) {
		model := Point12ValCFoundationModel()
		model.Dependency.ValBManifestIntegrityResult = point12ValBCheckResultUnsupported
		model = ComputePoint12ValCFoundation(model)
		if model.DependencyState != Point12ValCDependencyStateReviewRequired {
			t.Fatalf("expected valc dependency review required, got %#v", model)
		}
		if model.ExportState == Point12ValCExportStateReady || model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected valc unsupported dependency to avoid full export_ready/active, got %#v", model)
		}
	})

	t.Run("point12 pass token remains rejected", func(t *testing.T) {
		model := Point12ValDFoundationModel()
		model.Explanation.CustomerVisibleStatement = "point_12_pass"
		model.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Explanation)
		model = ComputePoint12ValDFoundation(model)
		if model.ExplanationState != Point12ValDExplanationStateBlocked {
			t.Fatalf("expected point_12_pass token to be rejected, got %#v", model)
		}
	})
}
