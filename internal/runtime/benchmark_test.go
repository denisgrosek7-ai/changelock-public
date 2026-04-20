package runtime

import (
	"fmt"
	"testing"
)

var runtimeBenchmarkProfiles = []struct {
	name           string
	containerCount int
}{
	{name: "small", containerCount: 1},
	{name: "medium", containerCount: 10},
	{name: "large", containerCount: 100},
}

func BenchmarkCompare(b *testing.B) {
	for _, profile := range runtimeBenchmarkProfiles {
		profile := profile
		approved, observed := benchmarkRuntimeStates(profile.containerCount)
		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result := Compare(approved, observed)
				if result.Result == "" {
					b.Fatal("expected comparison result")
				}
			}
		})
	}
}

func benchmarkRuntimeStates(containerCount int) (ApprovedWorkloadState, ObservedWorkloadState) {
	approved := ApprovedWorkloadState{
		TenantID:           "acme",
		ClusterID:          "local",
		Namespace:          "acme-prod",
		WorkloadKind:       "Deployment",
		Workload:           "edge-gateway",
		ServiceAccountName: "edge-sa",
		ExpectedConfigHash: "cfg-approved",
		Containers:         make([]ApprovedContainerState, 0, containerCount),
	}
	observed := ObservedWorkloadState{
		ClusterID:          "local",
		Namespace:          "acme-prod",
		WorkloadKind:       "Deployment",
		Workload:           "edge-gateway",
		ServiceAccountName: "edge-sa",
		ActualConfigHash:   "cfg-approved",
		Containers:         make([]ObservedContainerState, 0, containerCount),
	}

	for i := 0; i < containerCount; i++ {
		name := fmt.Sprintf("container-%03d", i)
		approved.Containers = append(approved.Containers, ApprovedContainerState{
			Name:           name,
			Image:          "ghcr.io/my-org/acme-app@sha256:abc123",
			ApprovedDigest: "sha256:abc123",
			Runtime: SecurityConstraints{
				RunAsNonRoot:             true,
				ReadOnlyRootFilesystem:   true,
				AllowPrivilegeEscalation: false,
				DropAllCapabilities:      true,
				SeccompRuntimeDefault:    true,
				DenyPrivileged:           true,
			},
		})
		runningDigest := "sha256:abc123"
		if i == containerCount-1 {
			runningDigest = "sha256:def456"
		}
		observed.Containers = append(observed.Containers, ObservedContainerState{
			Name:          name,
			Image:         "ghcr.io/my-org/acme-app@sha256:abc123",
			RunningDigest: runningDigest,
			Runtime: SecurityPosture{
				RunAsNonRoot:             true,
				ReadOnlyRootFilesystem:   true,
				AllowPrivilegeEscalation: false,
				DropAllCapabilities:      true,
				SeccompRuntimeDefault:    true,
				Privileged:               false,
			},
		})
	}

	return approved, observed
}
