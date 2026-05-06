package formal

import "testing"

func point15Val0ValidFreshnessTaxonomyModel() Point15Val0EvidenceFreshnessTaxonomy {
	return point15Val0FreshnessTaxonomyModel()
}

func point15Val0ValidDowngradeTaxonomyModel() Point15Val0DowngradeTaxonomy {
	return point15Val0DowngradeTaxonomyModel()
}

func point15Val0ValidFreshnessEvidenceContextModel() Point15Val0FreshnessEvidenceContext {
	return point15Val0FreshnessEvidenceContextModel(point15Val0DependencySnapshotModel())
}

func point15Val0ValidTimestampDisciplineModel() Point15Val0TimestampDiscipline {
	return point15Val0TimestampDisciplineModel(point15Val0DependencySnapshotModel())
}

func point15Val0ValidAuthorityBoundaryModel() Point15Val0AuthorityBoundary {
	return point15Val0AuthorityBoundaryModel(point15Val0DependencySnapshotModel())
}

func point15Val0ValidNoOverclaimGuardModel() Point15Val0NoOverclaimGuard {
	return point15Val0NoOverclaimGuardModel()
}

func TestPoint15Val0DependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15Val0DependencySnapshot)
		want   string
	}{
		{"active when point14 vale closure is clean", func(model *Point15Val0DependencySnapshot) {}, Point15Val0StateActive},
		{"blocks when point14 vale missing", func(model *Point15Val0DependencySnapshot) { model.Point14ValECurrentState = "" }, Point15Val0StateBlocked},
		{"blocks when point14 vale blocked", func(model *Point15Val0DependencySnapshot) { model.Point14ValECurrentState = Point14ValEStateBlocked }, Point15Val0StateBlocked},
		{"blocks when point14 vale review required", func(model *Point15Val0DependencySnapshot) {
			model.Point14ValECurrentState = Point14ValEStateReviewRequired
		}, Point15Val0StateBlocked},
		{"blocks when point14 vale incomplete", func(model *Point15Val0DependencySnapshot) { model.Point14ValECurrentState = Point14ValEStateIncomplete }, Point15Val0StateBlocked},
		{"blocks when point14 vale not merged", func(model *Point15Val0DependencySnapshot) { model.Point14ValEMerged = false }, Point15Val0StateBlocked},
		{"blocks when point14 vale ci not green", func(model *Point15Val0DependencySnapshot) { model.Point14ValECIGreen = false }, Point15Val0StateBlocked},
		{"blocks when point14 vale not reviewed on main", func(model *Point15Val0DependencySnapshot) {
			model.Point14ValEReviewedOnMain = false
		}, Point15Val0StateBlocked},
		{"blocks when embedded point14 snapshot is not computed from upstream", func(model *Point15Val0DependencySnapshot) {
			model.Point14ValEComputedFromUpstream = true
			model.Point14ValE.Dependency.SnapshotFromComputedOutput = false
		}, Point15Val0StateBlocked},
		{"blocks when point14 pass token absent", func(model *Point15Val0DependencySnapshot) { model.Point14PassToken = "" }, Point15Val0StateBlocked},
		{"blocks when point14 pass token not from vale", func(model *Point15Val0DependencySnapshot) { model.Point14PassManifestWaveID = point14ValDWaveID }, Point15Val0StateBlocked},
		{"blocks when point15 pass already appears", func(model *Point15Val0DependencySnapshot) { model.Point15PassSeen = true }, Point15Val0StateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0DependencySnapshotModel()
			tc.mutate(&model)
			if got := EvaluatePoint15Val0DependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15Val0EvidenceFreshnessTaxonomyState(t *testing.T) {
	tests := []struct {
		name    string
		status  string
		state   string
		outcome string
		want    string
	}{
		{"fresh active", point15Val0FreshnessFresh, Point15Val0StateActive, point15Val0DowngradeRetainActive, Point15Val0StateActive},
		{"stale review required", point15Val0FreshnessStale, Point15Val0StateReviewRequired, point15Val0DowngradeReview, Point15Val0StateReviewRequired},
		{"expired blocked", point15Val0FreshnessExpired, Point15Val0StateBlocked, point15Val0DowngradeBlocked, Point15Val0StateBlocked},
		{"missing incomplete", point15Val0FreshnessMissing, Point15Val0StateIncomplete, point15Val0DowngradeIncomplete, Point15Val0StateIncomplete},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0ValidFreshnessTaxonomyModel()
			model.FreshnessStatus = tc.status
			model.MappedState = tc.state
			model.MappedDowngradeOutcome = tc.outcome
			if got := EvaluatePoint15Val0EvidenceFreshnessTaxonomyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}

	rejects := []struct {
		name   string
		mutate func(*Point15Val0EvidenceFreshnessTaxonomy)
	}{
		{"rejects unknown freshness status", func(model *Point15Val0EvidenceFreshnessTaxonomy) { model.FreshnessStatus = "freshish" }},
		{"rejects empty freshness status", func(model *Point15Val0EvidenceFreshnessTaxonomy) { model.FreshnessStatus = "" }},
		{"rejects generic non empty status", func(model *Point15Val0EvidenceFreshnessTaxonomy) { model.FreshnessStatus = "active" }},
		{"rejects pass preserving alias", func(model *Point15Val0EvidenceFreshnessTaxonomy) {
			model.MappedDowngradeOutcome = "retain_pass_if_stale"
		}},
	}
	for _, tc := range rejects {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0ValidFreshnessTaxonomyModel()
			tc.mutate(&model)
			if got := EvaluatePoint15Val0EvidenceFreshnessTaxonomyState(model); got != Point15Val0StateBlocked {
				t.Fatalf("expected %s, got %s", Point15Val0StateBlocked, got)
			}
		})
	}
}

func TestPoint15Val0DowngradeTaxonomyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15Val0DowngradeTaxonomy)
		want   string
	}{
		{"stale downgrades and cannot retain pass", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessStale
			model.DowngradeOutcome = point15Val0DowngradeReview
		}, Point15Val0StateReviewRequired},
		{"expired blocks", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessExpired
			model.DowngradeOutcome = point15Val0DowngradeBlocked
		}, Point15Val0StateBlocked},
		{"revoked blocks", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessRevoked
			model.DowngradeOutcome = point15Val0DowngradeBlocked
		}, Point15Val0StateBlocked},
		{"superseded without lineage blocks", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessSuperseded
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.SupersessionLineageRef = ""
		}, Point15Val0StateBlocked},
		{"superseded with lineage requires review", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessSuperseded
			model.DowngradeOutcome = point15Val0DowngradeReview
			model.SupersessionLineageRef = "supersession_lineage_point15_val0_001"
		}, Point15Val0StateReviewRequired},
		{"drifted requires review", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessDrifted
			model.DowngradeOutcome = point15Val0DowngradeReview
		}, Point15Val0StateReviewRequired},
		{"decisive drift blocks", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessDrifted
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.DriftIsDecisive = true
		}, Point15Val0StateBlocked},
		{"missing freshness proof returns incomplete", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessMissing
			model.DowngradeOutcome = point15Val0DowngradeIncomplete
			model.FreshnessProofPresent = false
		}, Point15Val0StateIncomplete},
		{"decisive missing freshness proof blocks", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessMissing
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.FreshnessProofPresent = false
			model.MissingFreshnessProofDecisive = true
		}, Point15Val0StateBlocked},
		{"unsupported blocks", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessUnsupported
			model.DowngradeOutcome = point15Val0DowngradeBlocked
		}, Point15Val0StateBlocked},
		{"tampered blocks", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessTampered
			model.DowngradeOutcome = point15Val0DowngradeBlocked
		}, Point15Val0StateBlocked},
		{"stale cannot retain pass", func(model *Point15Val0DowngradeTaxonomy) {
			model.FreshnessStatus = point15Val0FreshnessStale
			model.DowngradeOutcome = point15Val0DowngradeReview
			model.RetainsPass = true
		}, Point15Val0StateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0ValidDowngradeTaxonomyModel()
			tc.mutate(&model)
			if got := EvaluatePoint15Val0DowngradeTaxonomyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15Val0TimestampDisciplineState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15Val0TimestampDiscipline)
		want   string
	}{
		{"server utc canonical times pass", func(model *Point15Val0TimestampDiscipline) {}, Point15Val0StateActive},
		{"approved customer controlled time source passes", func(model *Point15Val0TimestampDiscipline) {
			model.ObservedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.EvaluatedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.ValidatedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.ReferenceNowTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.ReviewerApprovedTimeSource = point14Val0TimeSourceApprovedCustomerTime
		}, Point15Val0StateActive},
		{"client local time cannot create freshness validity", func(model *Point15Val0TimestampDiscipline) {
			model.ClientLocalCreatesCanonical = true
		}, Point15Val0StateBlocked},
		{"source event at alone cannot create canonical validity", func(model *Point15Val0TimestampDiscipline) {
			model.SourceEventCreatesCanonical = true
		}, Point15Val0StateBlocked},
		{"future active evidence blocks", func(model *Point15Val0TimestampDiscipline) {
			model.ValidatedAt = "2026-05-06T18:35:00Z"
		}, Point15Val0StateBlocked},
		{"backdated approval requires review", func(model *Point15Val0TimestampDiscipline) {
			model.ReviewerApprovedAt = "2026-05-06T18:24:00Z"
		}, Point15Val0StateReviewRequired},
		{"impossible ordering blocks", func(model *Point15Val0TimestampDiscipline) {
			model.ObservedAt = "2026-05-06T18:26:00Z"
		}, Point15Val0StateBlocked},
		{"missing evaluated at returns incomplete", func(model *Point15Val0TimestampDiscipline) {
			model.EvaluatedAt = ""
		}, Point15Val0StateIncomplete},
		{"missing validated at returns incomplete", func(model *Point15Val0TimestampDiscipline) {
			model.ValidatedAt = ""
		}, Point15Val0StateIncomplete},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0ValidTimestampDisciplineModel()
			tc.mutate(&model)
			if got := EvaluatePoint15Val0TimestampDisciplineState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15Val0AuthorityBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15Val0AuthorityBoundary)
		want   string
	}{
		{"agent recommendation remains advisory only", func(model *Point15Val0AuthorityBoundary) {}, Point15Val0StateActive},
		{"external source must remain input only", func(model *Point15Val0AuthorityBoundary) {
			model.ExternalSourceInputOnly = false
		}, Point15Val0StateBlocked},
		{"scheduler cannot approve pass", func(model *Point15Val0AuthorityBoundary) { model.SchedulerPassAllowed = true }, Point15Val0StateBlocked},
		{"dashboard cannot mark evidence fresh", func(model *Point15Val0AuthorityBoundary) { model.DashboardFreshnessAllowed = true }, Point15Val0StateBlocked},
		{"agent freshness authority blocks", func(model *Point15Val0AuthorityBoundary) { model.AgentFreshnessAllowed = true }, Point15Val0StateBlocked},
		{"connector signal remains input only", func(model *Point15Val0AuthorityBoundary) { model.ConnectorFreshnessAuthorityAllowed = true }, Point15Val0StateBlocked},
		{"portal auditor customer projection cannot mutate freshness state", func(model *Point15Val0AuthorityBoundary) {
			model.CustomerProjectionMutatesFreshness = true
		}, Point15Val0StateBlocked},
		{"auditor projection cannot mutate freshness state", func(model *Point15Val0AuthorityBoundary) {
			model.AuditorProjectionMutatesFreshness = true
		}, Point15Val0StateBlocked},
		{"portal projection cannot mutate freshness state", func(model *Point15Val0AuthorityBoundary) {
			model.PortalProjectionMutatesFreshness = true
		}, Point15Val0StateBlocked},
		{"no production canonical mutation path", func(model *Point15Val0AuthorityBoundary) {
			model.CanonicalMutationAllowed = true
		}, Point15Val0StateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0ValidAuthorityBoundaryModel()
			tc.mutate(&model)
			if got := EvaluatePoint15Val0AuthorityBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15Val0FreshnessEvidenceContextState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15Val0FreshnessEvidenceContext)
		want   string
	}{
		{"freshness evidence context happy path active", func(model *Point15Val0FreshnessEvidenceContext) {}, Point15Val0StateActive},
		{"cross tenant freshness evidence blocks", func(model *Point15Val0FreshnessEvidenceContext) {
			model.ReferencedTenantScope = "tenant_point15_val0_other"
		}, Point15Val0StateBlocked},
		{"missing tenant scope incomplete", func(model *Point15Val0FreshnessEvidenceContext) { model.TenantScope = "" }, Point15Val0StateIncomplete},
		{"missing evidence id incomplete", func(model *Point15Val0FreshnessEvidenceContext) { model.EvidenceID = "" }, Point15Val0StateIncomplete},
		{"missing evidence hash incomplete", func(model *Point15Val0FreshnessEvidenceContext) { model.EvidenceHash = "" }, Point15Val0StateIncomplete},
		{"missing version incomplete", func(model *Point15Val0FreshnessEvidenceContext) { model.PolicyVersion = "" }, Point15Val0StateIncomplete},
		{"similar evidence names paths do not imply identity", func(model *Point15Val0FreshnessEvidenceContext) {
			model.IdentityInferredFromNameOrPath = true
		}, Point15Val0StateBlocked},
		{"missing freshness proof returns incomplete", func(model *Point15Val0FreshnessEvidenceContext) {
			model.FreshnessProofRef = ""
		}, Point15Val0StateIncomplete},
		{"expired freshness context blocks", func(model *Point15Val0FreshnessEvidenceContext) {
			model.FreshnessStatus = point15Val0FreshnessExpired
			model.DowngradeOutcome = point15Val0DowngradeBlocked
		}, Point15Val0StateBlocked},
		{"stale freshness context requires review", func(model *Point15Val0FreshnessEvidenceContext) {
			model.FreshnessStatus = point15Val0FreshnessStale
			model.DowngradeOutcome = point15Val0DowngradeReview
		}, Point15Val0StateReviewRequired},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0ValidFreshnessEvidenceContextModel()
			tc.mutate(&model)
			if got := EvaluatePoint15Val0FreshnessEvidenceContextState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15Val0NoOverclaimGuardState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15Val0NoOverclaimGuard)
		want   string
	}{
		{"forbidden wording blocks", func(model *Point15Val0NoOverclaimGuard) {
			model.ObservedTexts = []string{"continuous assurance guaranteed"}
		}, Point15Val0StateBlocked},
		{"safe bounded wording passes", func(model *Point15Val0NoOverclaimGuard) {}, Point15Val0StateActive},
		{"internal blocked diagnostics remain classified", func(model *Point15Val0NoOverclaimGuard) {
			model.InternalDiagnosticTexts = []string{"production approved"}
			model.InternalDiagnosticsClassifiedBlocked = true
		}, Point15Val0StateActive},
		{"unclassified internal blocked diagnostics block", func(model *Point15Val0NoOverclaimGuard) {
			model.InternalDiagnosticTexts = []string{"production approved"}
		}, Point15Val0StateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15Val0ValidNoOverclaimGuardModel()
			tc.mutate(&model)
			if got := EvaluatePoint15Val0NoOverclaimGuardState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15Val0FreshnessDisciplineFoundationHappyPath(t *testing.T) {
	model := Point15Val0FoundationModel()
	got := ComputePoint15Val0FreshnessDisciplineFoundation(model)
	if got.CurrentState != Point15Val0StateActive ||
		got.DependencyState != Point15Val0StateActive ||
		got.FreshnessTaxonomyState != Point15Val0StateActive ||
		got.DowngradeTaxonomyState != Point15Val0StateActive ||
		got.EvidenceContextState != Point15Val0StateActive ||
		got.TenantBoundaryState != Point15Val0StateActive ||
		got.TimestampDisciplineState != Point15Val0StateActive ||
		got.AuthorityBoundaryState != Point15Val0StateActive ||
		got.NoOverclaimState != Point15Val0StateActive {
		t.Fatalf("expected full point15 val0 foundation active, got %#v", got)
	}
}

func TestPoint15Val0FreshnessDisciplineFoundationStatusMismatchBlocks(t *testing.T) {
	model := Point15Val0FoundationModel()
	model.EvidenceContext.FreshnessStatus = point15Val0FreshnessExpired
	model.EvidenceContext.DowngradeOutcome = point15Val0DowngradeBlocked

	got := ComputePoint15Val0FreshnessDisciplineFoundation(model)
	if got.CurrentState != Point15Val0StateBlocked || got.EvidenceContextState != Point15Val0StateBlocked {
		t.Fatalf("expected expired evidence context mismatch to block foundation, got %#v", got)
	}
}

func TestPoint15Val0FreshnessDisciplineFoundationTimestampMismatchBlocks(t *testing.T) {
	model := Point15Val0FoundationModel()
	model.TimestampDiscipline.FreshnessStatus = point15Val0FreshnessStale

	got := ComputePoint15Val0FreshnessDisciplineFoundation(model)
	if got.CurrentState != Point15Val0StateBlocked || got.TimestampDisciplineState != Point15Val0StateBlocked {
		t.Fatalf("expected timestamp freshness mismatch to block foundation, got %#v", got)
	}
}

func TestPoint15Val0FreshnessDisciplineFoundationDowngradeMismatchBlocks(t *testing.T) {
	model := Point15Val0FoundationModel()
	model.DowngradeTaxonomy.DowngradeOutcome = point15Val0DowngradeBlocked

	got := ComputePoint15Val0FreshnessDisciplineFoundation(model)
	if got.CurrentState != Point15Val0StateBlocked || got.EvidenceContextState != Point15Val0StateBlocked {
		t.Fatalf("expected downgrade mismatch to block foundation, got %#v", got)
	}
}

func TestPoint10ThroughPoint15Val0CurrentSweep(t *testing.T) {
	point14ValE := ComputePoint14ValEFoundation(Point14ValEFoundationModel())
	if point14ValE.CurrentState != Point14ValEStatePassConfirmed ||
		!point14ValE.Point14PassAllowed ||
		point14ValE.Point14PassToken != point14Val0BlockedPassToken {
		t.Fatalf("expected point14 vale pass confirmed before point15, got %#v", point14ValE)
	}

	point15Val0 := ComputePoint15Val0FreshnessDisciplineFoundation(Point15Val0FoundationModel())
	if point15Val0.CurrentState != Point15Val0StateActive ||
		point15Val0.DependencyState != Point15Val0StateActive ||
		point15Val0.Dependency.Point14ValECurrentState != point14ValE.CurrentState ||
		point15Val0.Dependency.Point14PassToken != point14ValE.Point14PassToken {
		t.Fatalf("expected point15 val0 active and exact-bound to point14 vale closure, got %#v", point15Val0)
	}
}
