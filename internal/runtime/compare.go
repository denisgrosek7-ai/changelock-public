package runtime

import "sort"

func Compare(approved ApprovedWorkloadState, observed ObservedWorkloadState) ComparisonResult {
	result := ComparisonResult{
		Namespace: firstNonEmpty(observed.Namespace, approved.Namespace),
		Workload:  firstNonEmpty(observed.Workload, approved.Workload),
	}

	evidence := &DriftEvidence{}
	classSet := map[DriftClass]struct{}{}
	approvedByName := indexApprovedContainers(approved.Containers)
	observedByName := indexObservedContainers(observed.Containers)

	for _, approvedContainer := range approved.Containers {
		result.Image = firstNonEmpty(result.Image, approvedContainer.Image)
		result.ApprovedDigest = firstNonEmpty(result.ApprovedDigest, approvedContainer.ApprovedDigest)

		observedContainer, ok := observedByName[approvedContainer.Name]
		if !ok {
			classSet[DriftClassImage] = struct{}{}
			evidence.MissingContainers = append(evidence.MissingContainers, approvedContainer.Name)
			result.Reasons = append(result.Reasons, approvedContainer.Name+": approved container is missing from runtime state")
			continue
		}

		result.Image = firstNonEmpty(result.Image, observedContainer.Image)
		result.RunningDigest = firstNonEmpty(result.RunningDigest, observedContainer.RunningDigest)

		if approvedContainer.ApprovedDigest != "" && approvedContainer.ApprovedDigest != observedContainer.RunningDigest {
			classSet[DriftClassImage] = struct{}{}
			evidence.ImageMismatches = append(evidence.ImageMismatches, ImageMismatch{
				Container:      approvedContainer.Name,
				ApprovedImage:  approvedContainer.Image,
				RunningImage:   observedContainer.Image,
				ApprovedDigest: approvedContainer.ApprovedDigest,
				RunningDigest:  observedContainer.RunningDigest,
			})
			result.Reasons = append(result.Reasons, approvedContainer.Name+": running digest does not match approved digest")
		}

		securityMismatches := compareSecurityContext(approvedContainer.Name, approvedContainer.Runtime, observedContainer.Runtime)
		if len(securityMismatches) > 0 {
			classSet[DriftClassSecurityContext] = struct{}{}
			evidence.SecurityContextMismatches = append(evidence.SecurityContextMismatches, securityMismatches...)
			for _, mismatch := range securityMismatches {
				result.Reasons = append(result.Reasons, mismatch.Container+": security context drift for "+mismatch.Field)
			}
		}
	}

	unexpectedContainers := []string{}
	for name, observedContainer := range observedByName {
		if _, ok := approvedByName[name]; ok {
			continue
		}
		classSet[DriftClassImage] = struct{}{}
		unexpectedContainers = append(unexpectedContainers, name)
		result.Image = firstNonEmpty(result.Image, observedContainer.Image)
		result.RunningDigest = firstNonEmpty(result.RunningDigest, observedContainer.RunningDigest)
		result.Reasons = append(result.Reasons, name+": unexpected runtime container is present")
	}
	sort.Strings(unexpectedContainers)
	evidence.UnexpectedContainers = append(evidence.UnexpectedContainers, unexpectedContainers...)

	if approved.ExpectedConfigHash != "" && approved.ExpectedConfigHash != observed.ActualConfigHash {
		classSet[DriftClassConfig] = struct{}{}
		evidence.ConfigExpectation = approved.ExpectedConfigHash
		evidence.ConfigObserved = observed.ActualConfigHash
		result.Reasons = append(result.Reasons, "runtime config hash does not match approved config hash")
	}

	if len(classSet) == 0 {
		result.Result = string(DriftClassNoDrift)
	} else {
		result.Classes = orderedClasses(classSet)
		if len(result.Classes) == 1 {
			result.Result = result.Classes[0]
		} else {
			result.Result = string(DriftClassMultiple)
		}
		result.Evidence = trimEvidence(evidence)
	}

	if result.Image == "" {
		result.Image = firstContainerImage(approved.Containers, observed.Containers)
	}
	if result.RunningDigest == "" {
		result.RunningDigest = firstRunningDigest(observed.Containers)
	}
	if result.ApprovedDigest == "" {
		result.ApprovedDigest = firstApprovedDigest(approved.Containers)
	}

	return result
}

func compareSecurityContext(containerName string, expected SecurityConstraints, actual SecurityPosture) []SecurityContextMismatch {
	mismatches := []SecurityContextMismatch{}

	if expected.RunAsNonRoot && !actual.RunAsNonRoot {
		mismatches = append(mismatches, SecurityContextMismatch{Container: containerName, Field: "runAsNonRoot", Expected: true, Actual: actual.RunAsNonRoot})
	}
	if expected.ReadOnlyRootFilesystem && !actual.ReadOnlyRootFilesystem {
		mismatches = append(mismatches, SecurityContextMismatch{Container: containerName, Field: "readOnlyRootFilesystem", Expected: true, Actual: actual.ReadOnlyRootFilesystem})
	}
	if !expected.AllowPrivilegeEscalation && actual.AllowPrivilegeEscalation {
		mismatches = append(mismatches, SecurityContextMismatch{Container: containerName, Field: "allowPrivilegeEscalation", Expected: false, Actual: actual.AllowPrivilegeEscalation})
	}
	if expected.DropAllCapabilities && !actual.DropAllCapabilities {
		mismatches = append(mismatches, SecurityContextMismatch{Container: containerName, Field: "dropAllCapabilities", Expected: true, Actual: actual.DropAllCapabilities})
	}
	if expected.SeccompRuntimeDefault && !actual.SeccompRuntimeDefault {
		mismatches = append(mismatches, SecurityContextMismatch{Container: containerName, Field: "seccompRuntimeDefault", Expected: true, Actual: actual.SeccompRuntimeDefault})
	}
	if expected.DenyPrivileged && actual.Privileged {
		mismatches = append(mismatches, SecurityContextMismatch{Container: containerName, Field: "privileged", Expected: false, Actual: actual.Privileged})
	}

	return mismatches
}

func orderedClasses(classes map[DriftClass]struct{}) []string {
	ordered := []DriftClass{DriftClassImage, DriftClassConfig, DriftClassSecurityContext}
	values := make([]string, 0, len(classes))
	for _, class := range ordered {
		if _, ok := classes[class]; ok {
			values = append(values, string(class))
		}
	}
	return values
}

func trimEvidence(evidence *DriftEvidence) *DriftEvidence {
	if evidence == nil {
		return nil
	}
	if len(evidence.ImageMismatches) == 0 && evidence.ConfigExpectation == "" && evidence.ConfigObserved == "" && len(evidence.SecurityContextMismatches) == 0 && len(evidence.MissingContainers) == 0 && len(evidence.UnexpectedContainers) == 0 {
		return nil
	}
	return evidence
}

func indexApprovedContainers(containers []ApprovedContainerState) map[string]ApprovedContainerState {
	index := make(map[string]ApprovedContainerState, len(containers))
	for _, container := range containers {
		index[container.Name] = container
	}
	return index
}

func indexObservedContainers(containers []ObservedContainerState) map[string]ObservedContainerState {
	index := make(map[string]ObservedContainerState, len(containers))
	for _, container := range containers {
		index[container.Name] = container
	}
	return index
}

func firstContainerImage(approved []ApprovedContainerState, observed []ObservedContainerState) string {
	for _, container := range approved {
		if container.Image != "" {
			return container.Image
		}
	}
	for _, container := range observed {
		if container.Image != "" {
			return container.Image
		}
	}
	return ""
}

func firstApprovedDigest(containers []ApprovedContainerState) string {
	for _, container := range containers {
		if container.ApprovedDigest != "" {
			return container.ApprovedDigest
		}
	}
	return ""
}

func firstRunningDigest(containers []ObservedContainerState) string {
	for _, container := range containers {
		if container.RunningDigest != "" {
			return container.RunningDigest
		}
	}
	return ""
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
