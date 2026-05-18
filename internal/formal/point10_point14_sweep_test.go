package formal

import (
	"testing"

	"github.com/denisgrosek/changelock/internal/operability"
)

func TestPoint10ThroughPoint14CurrentSweep(t *testing.T) {
	point10 := operability.ComputeDeploymentMultiTenantValEFoundation(operability.DeploymentMultiTenantValEFoundationModel())
	if point10.Point10State != operability.DeploymentMultiTenantPoint10StatePass ||
		point10.NoOverclaimState != operability.DeploymentMultiTenantValENoOverclaimStateActive ||
		point10.ProjectionBoundaryState != operability.DeploymentMultiTenantValEProjectionBoundaryStateActive ||
		point10.Point10PassRuleState != operability.DeploymentMultiTenantValEPoint10PassRuleStateActive {
		t.Fatalf("expected point10 chain foundation active/pass, got %#v", point10)
	}

	point11 := point14Val0ActivePoint11ValDFoundation()
	if point11.CurrentState != Point11ValDStateActive ||
		point11.DependencyState != Point11ValDDependencyStateActive ||
		point11.PublicationReviewState != Point11ValDPublicationReviewStateActive ||
		point11.NoOverclaimReviewState != Point11ValDNoOverclaimReviewStateActive ||
		point11.FinalPassGateState != Point11ValDFinalPassGateStateActive {
		t.Fatalf("expected point11 chain foundation active for point14 inherited dependency closure, got %#v", point11)
	}

	point12 := ComputePoint12ValEFoundation(Point12ValEFoundationModel())
	if point12.CurrentState != Point12ValEStatePassConfirmed ||
		point12.DependencyState != Point12ValEStateActive ||
		point12.PassClosureManifestState != Point12ValEStateActive ||
		point12.PassClosureManifest.ReviewerResult != point12ValEReviewerResultPassConfirmed {
		t.Fatalf("expected point12 chain closure pass confirmed, got %#v", point12)
	}

	point13 := ComputePoint13ValEFoundation(Point13ValEFoundationModel())
	if point13.CurrentState != Point13ValEStatePassConfirmed ||
		point13.DependencyState != Point13ValEStateActive ||
		point13.PassClosureManifestState != Point13ValEStateActive ||
		!point13.Point13PassAllowed ||
		point13.Point13PassToken != point13ValEPoint13PassToken {
		t.Fatalf("expected point13 chain closure pass confirmed, got %#v", point13)
	}

	point14Val0 := ComputePoint14Val0Foundation(Point14Val0FoundationModel())
	if point14Val0.CurrentState != Point14Val0StateActive ||
		point14Val0.Dependency.Point13ValECurrentState != point13.CurrentState ||
		point14Val0.Dependency.InheritedPoint11FinalPassGateState != point11.FinalPassGateState {
		t.Fatalf("expected point14 val0 active and exact-bound to point13/point11, got %#v", point14Val0)
	}

	point14ValA := ComputePoint14ValAFoundation(Point14ValAFoundationModel())
	if point14ValA.CurrentState != Point14ValAStateActive ||
		point14ValA.Dependency.Point14Val0CurrentState != point14Val0.CurrentState ||
		point14ValA.Dependency.InheritedPoint11FinalPassGateState != point11.FinalPassGateState {
		t.Fatalf("expected point14 vala active and exact-bound to val0/point11, got %#v", point14ValA)
	}

	point14ValB := ComputePoint14ValBFoundation(Point14ValBFoundationModel())
	if point14ValB.CurrentState != Point14ValBStateActive ||
		point14ValB.Dependency.Point14ValACurrentState != point14ValA.CurrentState ||
		point14ValB.Dependency.InheritedPoint11FinalPassGateState != point11.FinalPassGateState ||
		point14ValB.Dependency.Point14PassSeen {
		t.Fatalf("expected point14 valb active with no premature point14 pass token path, got %#v", point14ValB)
	}

	point14ValC := ComputePoint14ValCFoundation(Point14ValCFoundationModel())
	if point14ValC.CurrentState != Point14ValCStateActive ||
		point14ValC.Dependency.Point14ValBCurrentState != point14ValB.CurrentState ||
		point14ValC.Dependency.InheritedPoint11FinalPassGateState != point11.FinalPassGateState ||
		point14ValC.Dependency.Point14PassSeen {
		t.Fatalf("expected point14 valc active with no premature point14 pass token path, got %#v", point14ValC)
	}

	point14ValD := ComputePoint14ValDFoundation(Point14ValDFoundationModel())
	if point14ValD.CurrentState != Point14ValDStateActive ||
		point14ValD.Dependency.Point14ValCCurrentState != point14ValC.CurrentState ||
		point14ValD.Dependency.InheritedPoint11FinalPassGateState != point11.FinalPassGateState ||
		point14ValD.Dependency.Point14PassSeen {
		t.Fatalf("expected point14 vald active with no premature point14 pass token path, got %#v", point14ValD)
	}

	point14ValE := ComputePoint14ValEFoundation(Point14ValEFoundationModel())
	if point14ValE.CurrentState != Point14ValEStatePassConfirmed ||
		point14ValE.Dependency.Point14ValDCurrentState != point14ValD.CurrentState ||
		point14ValE.Dependency.InheritedPoint11FinalPassGateState != point11.FinalPassGateState ||
		!point14ValE.Point14PassAllowed ||
		point14ValE.Point14PassToken != point14Val0BlockedPassToken {
		t.Fatalf("expected point14 vale pass confirmed with final point14 pass token, got %#v", point14ValE)
	}
}
