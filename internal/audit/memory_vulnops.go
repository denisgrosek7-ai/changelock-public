package audit

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

func (s *MemoryStore) IngestSBOM(_ context.Context, request SBOMIngestRequest) (SBOMIngestResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	request, err := NormalizeSBOMIngestRequest(request)
	if err != nil {
		return SBOMIngestResult{}, err
	}
	hash := sbomHash(request.SBOM)
	key := request.ImageDigest + "|" + hash
	if existing, ok := s.sbomDocuments[key]; ok {
		componentCount := 0
		for _, component := range s.sbomComponents {
			if component.ImageDigest == request.ImageDigest {
				componentCount++
			}
		}
		return SBOMIngestResult{
			DocumentStored:     false,
			DocumentID:         existing.ID,
			ImageDigest:        request.ImageDigest,
			SBOMHash:           hash,
			ComponentsIngested: componentCount,
		}, nil
	}

	components, err := ParseSBOMComponents(request.SBOMFormat, request.SBOM, request.ImageDigest)
	if err != nil {
		return SBOMIngestResult{}, err
	}

	now := s.now().UTC()
	document := SBOMDocument{
		ID:          s.nextSBOMDocumentID,
		ImageDigest: request.ImageDigest,
		ImageRef:    request.ImageRef,
		SBOMFormat:  request.SBOMFormat,
		SourceRef:   request.SourceRef,
		SBOMHash:    hash,
		RawSBOM:     append(jsonRawCopy(nil), request.SBOM...),
		CreatedAt:   now,
	}
	s.nextSBOMDocumentID++
	s.sbomDocuments[key] = document

	for i := range components {
		components[i].ID = s.nextSBOMComponentID
		components[i].CreatedAt = now
		components[i].Metadata = normalizeMetadata(components[i].Metadata)
		s.nextSBOMComponentID++
		s.sbomComponents = append(s.sbomComponents, components[i])
	}

	return SBOMIngestResult{
		DocumentStored:     true,
		DocumentID:         document.ID,
		ImageDigest:        request.ImageDigest,
		SBOMHash:           hash,
		ComponentsIngested: len(components),
	}, nil
}

func (s *MemoryStore) GetSBOMImage(_ context.Context, imageDigest string, limit int) (SBOMImageResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	imageDigest = strings.TrimSpace(imageDigest)
	if imageDigest == "" {
		return SBOMImageResponse{}, fmt.Errorf("%w: image_digest is required", ErrInvalidFilter)
	}
	if limit <= 0 {
		limit = 100
	}
	if limit > 1000 {
		limit = 1000
	}

	var document *SBOMDocument
	for _, candidate := range s.sbomDocuments {
		if candidate.ImageDigest != imageDigest {
			continue
		}
		if document == nil || candidate.CreatedAt.After(document.CreatedAt) {
			copy := cloneSBOMDocument(candidate)
			document = &copy
		}
	}
	if document == nil {
		return SBOMImageResponse{}, ErrExceptionNotFound
	}

	components := make([]SBOMComponent, 0)
	for _, component := range s.sbomComponents {
		if component.ImageDigest != imageDigest {
			continue
		}
		components = append(components, cloneSBOMComponent(component))
	}
	sort.Slice(components, func(i, j int) bool {
		if components[i].ComponentName == components[j].ComponentName {
			return components[i].ComponentVersion < components[j].ComponentVersion
		}
		return components[i].ComponentName < components[j].ComponentName
	})
	componentCount := len(components)
	if len(components) > limit {
		components = components[:limit]
	}
	return SBOMImageResponse{Document: *document, ComponentCount: componentCount, Components: components}, nil
}

func (s *MemoryStore) SearchSBOMComponents(_ context.Context, filter SBOMComponentSearchFilter) ([]SBOMComponent, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeSBOMComponentSearchFilter(filter)
	if err != nil {
		return nil, err
	}

	results := make([]SBOMComponent, 0)
	nameNeedle := strings.ToLower(filter.ComponentName)
	purlNeedle := strings.ToLower(filter.PURL)
	for _, component := range s.sbomComponents {
		if filter.ImageDigest != "" && component.ImageDigest != filter.ImageDigest {
			continue
		}
		if !matchesDigestScopeFilters(s.workloadRefsForDigestLocked(component.ImageDigest), filter.TenantID, "") {
			continue
		}
		if nameNeedle != "" && !strings.Contains(strings.ToLower(component.ComponentName), nameNeedle) {
			continue
		}
		if purlNeedle != "" && !strings.Contains(strings.ToLower(component.PURL), purlNeedle) {
			continue
		}
		results = append(results, cloneSBOMComponent(component))
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].ComponentName == results[j].ComponentName {
			if results[i].ComponentVersion == results[j].ComponentVersion {
				return results[i].ImageDigest < results[j].ImageDigest
			}
			return results[i].ComponentVersion < results[j].ComponentVersion
		}
		return results[i].ComponentName < results[j].ComponentName
	})
	if len(results) > filter.Limit {
		results = results[:filter.Limit]
	}
	return results, nil
}

func (s *MemoryStore) RecordVulnerabilityScan(_ context.Context, request VulnerabilityScanRequest) (VulnerabilityScanIngestResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	request, err := NormalizeVulnerabilityScanRequest(request, s.now)
	if err != nil {
		return VulnerabilityScanIngestResult{}, err
	}

	now := s.now().UTC()
	hadPriorSuccessfulRun := false
	for _, run := range s.scanRuns {
		if run.ImageDigest == request.ImageDigest && run.Status == VulnerabilityScanStatusCompleted {
			hadPriorSuccessfulRun = true
			break
		}
	}

	run := VulnerabilityScanRun{
		ID:          s.nextScanRunID,
		ImageDigest: request.ImageDigest,
		ImageRef:    request.ImageRef,
		Scanner:     request.Scanner,
		ScanMode:    request.ScanMode,
		StartedAt:   request.StartedAt.UTC(),
		CompletedAt: normalizeTimePointer(request.CompletedAt),
		Status:      request.Status,
		Summary:     normalizeMetadata(request.Summary),
		SourceRef:   request.SourceRef,
		CreatedAt:   now,
	}
	s.nextScanRunID++
	s.scanRuns = append(s.scanRuns, run)

	seenKeys := map[string]struct{}{}
	findings := make([]VulnerabilityFinding, 0, len(request.Findings))
	newFindings := []VulnerabilityFinding{}
	for _, findingInput := range request.Findings {
		key := vulnerabilityFindingKey(request.ImageDigest, findingInput)
		seenKeys[key] = struct{}{}

		existing, ok := s.findings[key]
		isNew := !ok || existing.Status == VulnerabilityFindingStatusResolved
		if !ok {
			existing = VulnerabilityFinding{
				ID:          s.nextFindingID,
				ImageDigest: request.ImageDigest,
				FirstSeenAt: now,
			}
			s.nextFindingID++
		}
		existing.ImageRef = firstNonEmpty(request.ImageRef, existing.ImageRef)
		existing.ScanRunID = run.ID
		existing.CVEID = strings.TrimSpace(strings.ToUpper(findingInput.CVEID))
		existing.Severity = findingInput.Severity
		existing.PackageName = findingInput.PackageName
		existing.PackageVersion = findingInput.PackageVersion
		existing.FixedVersion = findingInput.FixedVersion
		existing.PURL = findingInput.PURL
		existing.Status = VulnerabilityFindingStatusOpen
		existing.Title = findingInput.Title
		existing.Description = findingInput.Description
		existing.Source = findingInput.Source
		existing.Metadata = normalizeMetadata(findingInput.Metadata)
		existing.LastSeenAt = now
		s.findings[key] = existing
		finding := cloneVulnerabilityFinding(existing)
		findings = append(findings, finding)
		if isNew {
			newFindings = append(newFindings, finding)
		}
	}

	for key, finding := range s.findings {
		if finding.ImageDigest != request.ImageDigest {
			continue
		}
		if _, ok := seenKeys[key]; ok {
			continue
		}
		if finding.Status != VulnerabilityFindingStatusResolved {
			finding.Status = VulnerabilityFindingStatusResolved
			finding.ScanRunID = run.ID
			finding.LastSeenAt = now
			s.findings[key] = finding
		}
	}

	return VulnerabilityScanIngestResult{
		Run:                   run,
		Findings:              findings,
		NewFindings:           newFindings,
		HadPriorSuccessfulRun: hadPriorSuccessfulRun,
	}, nil
}

func (s *MemoryStore) ListActiveVulnerabilities(_ context.Context, filter VulnerabilityActiveFilter) ([]VulnerabilityFinding, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter = NormalizeVulnerabilityActiveFilter(filter)
	now := s.now().UTC()
	scope := vexScopeFromFilter(filter)
	results := make([]VulnerabilityFinding, 0)
	for _, finding := range s.findings {
		if finding.Status != VulnerabilityFindingStatusOpen {
			continue
		}
		if filter.ImageDigest != "" && finding.ImageDigest != filter.ImageDigest {
			continue
		}
		if filter.CVEID != "" && finding.CVEID != filter.CVEID {
			continue
		}
		if filter.Severity != "" && strings.ToUpper(finding.Severity) != filter.Severity {
			continue
		}
		if filter.ComponentName != "" && !strings.Contains(strings.ToLower(firstNonEmpty(finding.PackageName, finding.PURL)), strings.ToLower(filter.ComponentName)) {
			continue
		}
		if !matchesDigestScopeFilters(s.workloadRefsForDigestLocked(finding.ImageDigest), filter.TenantID, filter.Environment) {
			continue
		}
		statement := s.currentVEXStatementLocked(finding, scope, now)
		if statement != nil && internalvex.SuppressesFinding(*statement) && !filter.IncludeSuppressed {
			continue
		}
		item := cloneVulnerabilityFindingWithVEX(finding)
		if statement != nil {
			item.VEX = statementToVEXMatch(*statement)
			item.Decision = statementToLegacyDecision(*statement)
		}
		results = append(results, item)
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].Severity == results[j].Severity {
			if results[i].CVEID == results[j].CVEID {
				return results[i].ImageDigest < results[j].ImageDigest
			}
			return results[i].CVEID < results[j].CVEID
		}
		return severityRank(results[i].Severity) > severityRank(results[j].Severity)
	})
	if len(results) > filter.Limit {
		results = results[:filter.Limit]
	}
	return results, nil
}

func (s *MemoryStore) VulnerabilityBlastRadius(_ context.Context, filter VulnerabilityBlastRadiusFilter) (VulnerabilityBlastRadiusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeVulnerabilityBlastRadiusFilter(filter)
	if err != nil {
		return VulnerabilityBlastRadiusResponse{}, err
	}

	digestMap := map[string]*VulnerabilityBlastRadiusItem{}
	addDigest := func(imageDigest, imageRef string, finding *VulnerabilityFinding) {
		workloads := s.workloadRefsForDigestLocked(imageDigest)
		if !matchesDigestScopeFilters(workloads, filter.TenantID, "") {
			return
		}
		item := digestMap[imageDigest]
		if item == nil {
			item = &VulnerabilityBlastRadiusItem{
				ImageDigest: imageDigest,
				ImageRef:    imageRef,
				Findings:    []VulnerabilityFinding{},
				Workloads:   workloads,
			}
			digestMap[imageDigest] = item
		}
		if item.ImageRef == "" {
			item.ImageRef = imageRef
		}
		if finding != nil {
			item.Findings = append(item.Findings, cloneVulnerabilityFinding(*finding))
		}
	}

	if filter.CVEID != "" {
		now := s.now().UTC()
		for _, finding := range s.findings {
			if finding.CVEID != filter.CVEID || finding.Status != VulnerabilityFindingStatusOpen {
				continue
			}
			decision := s.currentDecisionLocked(finding.ImageDigest, finding.CVEID, now)
			copy := cloneVulnerabilityFinding(finding)
			if decision != nil && activeDecisionApplies(decision, now) {
				decisionCopy := cloneVulnerabilityDecision(*decision)
				copy.Decision = &decisionCopy
			}
			addDigest(finding.ImageDigest, finding.ImageRef, &copy)
		}
	}
	if filter.ComponentName != "" || filter.PURL != "" {
		nameNeedle := strings.ToLower(filter.ComponentName)
		purlNeedle := strings.ToLower(filter.PURL)
		for _, component := range s.sbomComponents {
			if nameNeedle != "" && !strings.Contains(strings.ToLower(component.ComponentName), nameNeedle) {
				continue
			}
			if purlNeedle != "" && !strings.Contains(strings.ToLower(component.PURL), purlNeedle) {
				continue
			}
			addDigest(component.ImageDigest, s.imageRefForDigestLocked(component.ImageDigest), nil)
		}
	}

	items := make([]VulnerabilityBlastRadiusItem, 0, len(digestMap))
	for _, item := range digestMap {
		items = append(items, *item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ImageDigest < items[j].ImageDigest })
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return VulnerabilityBlastRadiusResponse{
		Items: items,
		AppliedFilters: map[string]string{
			"cve_id":         filter.CVEID,
			"component_name": filter.ComponentName,
			"purl":           filter.PURL,
		},
	}, nil
}

func (s *MemoryStore) VulnerabilityTimeline(_ context.Context, filter VulnerabilityTimelineFilter) (VulnerabilityTimelineResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := NormalizeVulnerabilityTimelineFilter(filter)
	if err != nil {
		return VulnerabilityTimelineResponse{}, err
	}
	windowStart := s.now().UTC().AddDate(0, 0, -filter.WindowDays)
	now := s.now().UTC()

	items := []VulnerabilityTimelineEntry{}
	for _, finding := range s.findings {
		if finding.ImageDigest != filter.ImageDigest || finding.CVEID != filter.CVEID {
			continue
		}
		if !matchesDigestScopeFilters(s.workloadRefsForDigestLocked(finding.ImageDigest), filter.TenantID, "") {
			continue
		}
		if finding.LastSeenAt.Before(windowStart) {
			continue
		}
		entry := VulnerabilityTimelineEntry{
			ImageDigest:    finding.ImageDigest,
			CVEID:          finding.CVEID,
			PackageName:    finding.PackageName,
			PackageVersion: finding.PackageVersion,
			Severity:       finding.Severity,
			Status:         finding.Status,
			FirstSeenAt:    finding.FirstSeenAt,
			LastSeenAt:     finding.LastSeenAt,
		}
		if statement := s.currentVEXStatementLocked(finding, internalvex.EvaluationScope{TenantID: filter.TenantID}, now); statement != nil {
			entry.Decision = statementToLegacyDecision(*statement)
		}
		items = append(items, entry)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].FirstSeenAt.Before(items[j].FirstSeenAt) })
	return VulnerabilityTimelineResponse{
		Items: items,
		AppliedFilters: map[string]string{
			"image_digest": filter.ImageDigest,
			"cve_id":       filter.CVEID,
			"window_days":  filterWindowString(filter.WindowDays),
		},
	}, nil
}

func (s *MemoryStore) ListVulnerabilityDecisions(_ context.Context, filter VulnerabilityDecisionFilter) ([]VulnerabilityDecision, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter = NormalizeVulnerabilityDecisionFilter(filter)
	decisions := make([]VulnerabilityDecision, 0)
	for _, statement := range s.vexStatements {
		decision := statementToLegacyDecision(statement)
		if filter.ImageDigest != "" && decision.ImageDigest != filter.ImageDigest {
			continue
		}
		if filter.CVEID != "" && decision.CVEID != filter.CVEID {
			continue
		}
		if filter.Active != nil && decision.Active != *filter.Active {
			continue
		}
		if !matchesDigestScopeFilters(s.workloadRefsForDigestLocked(decision.ImageDigest), filter.TenantID, "") {
			continue
		}
		decisions = append(decisions, cloneVulnerabilityDecision(*decision))
	}
	sort.Slice(decisions, func(i, j int) bool { return decisions[i].CreatedAt.After(decisions[j].CreatedAt) })
	if len(decisions) > filter.Limit {
		decisions = decisions[:filter.Limit]
	}
	return decisions, nil
}

func (s *MemoryStore) GetVulnerabilityDecision(_ context.Context, decisionID int64) (VulnerabilityDecision, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	statement, ok := s.vexStatements[decisionID]
	if !ok {
		return VulnerabilityDecision{}, ErrExceptionNotFound
	}
	decision := statementToLegacyDecision(statement)
	return cloneVulnerabilityDecision(*decision), nil
}

func (s *MemoryStore) CreateVulnerabilityDecision(_ context.Context, request VulnerabilityDecisionCreateRequest, decidedBy string) (VulnerabilityDecision, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	createRequest, err := legacyDecisionCreateRequestToVEX(request)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	decidedBy = strings.TrimSpace(decidedBy)
	if decidedBy == "" {
		return VulnerabilityDecision{}, fmt.Errorf("%w: decided_by is required", ErrInvalidException)
	}
	statement, err := s.createVEXStatementLocked(createRequest, decidedBy)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	decision := statementToLegacyDecision(statement)
	return cloneVulnerabilityDecision(*decision), nil
}

func (s *MemoryStore) DeactivateVulnerabilityDecision(_ context.Context, decisionID int64) (VulnerabilityDecision, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	statement, err := s.revokeVEXStatementLocked(decisionID, "")
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	decision := statementToLegacyDecision(statement)
	return cloneVulnerabilityDecision(*decision), nil
}

func (s *MemoryStore) ListActiveDigests(_ context.Context, windowDays int, limit int) ([]ActiveDigestRef, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if windowDays <= 0 {
		windowDays = 30
	}
	if limit <= 0 {
		limit = 100
	}
	now := s.now().UTC()
	windowStart := now.AddDate(0, 0, -windowDays)
	items := map[string]*ActiveDigestRef{}
	for _, record := range s.records {
		if record.ReceivedAt.Before(windowStart) {
			continue
		}
		digest := strings.TrimSpace(record.Digest)
		if digest == "" {
			digest = DigestFromImage(record.Image)
		}
		if digest == "" {
			continue
		}
		item := items[digest]
		if item == nil {
			item = &ActiveDigestRef{ImageDigest: digest, ImageRef: strings.TrimSpace(record.Image), Repo: strings.TrimSpace(record.Repo)}
			items[digest] = item
		}
		scope := ActiveWorkloadRef{
			TenantID:    record.TenantID,
			Environment: record.Environment,
			Namespace:   record.Namespace,
			Workload:    record.Workload,
			Repo:        record.Repo,
			Image:       record.Image,
			Digest:      digest,
		}
		appendScopeIfMissing(item, scope)
	}
	results := make([]ActiveDigestRef, 0, len(items))
	for _, item := range items {
		results = append(results, *item)
	}
	sort.Slice(results, func(i, j int) bool { return results[i].ImageDigest < results[j].ImageDigest })
	if len(results) > limit {
		results = results[:limit]
	}
	return results, nil
}

func (s *MemoryStore) LookupDigestScopes(_ context.Context, imageDigest string, limit int) ([]ActiveWorkloadRef, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	workloads := s.workloadRefsForDigestLocked(strings.TrimSpace(imageDigest))
	if limit > 0 && len(workloads) > limit {
		workloads = workloads[:limit]
	}
	return workloads, nil
}

func (s *MemoryStore) currentDecisionLocked(imageDigest, cveID string, now time.Time) *VulnerabilityDecision {
	finding := VulnerabilityFinding{ImageDigest: imageDigest, CVEID: cveID}
	statement := s.currentVEXStatementLocked(finding, internalvex.EvaluationScope{}, now)
	if statement == nil {
		return nil
	}
	return statementToLegacyDecision(*statement)
}

func (s *MemoryStore) currentVEXStatementLocked(finding VulnerabilityFinding, scope internalvex.EvaluationScope, now time.Time) *internalvex.Statement {
	var current *internalvex.Statement
	for _, statement := range s.vexStatements {
		if !internalvex.Matches(statement, findingRefFromFinding(finding), scope, now) {
			continue
		}
		if current == nil || statement.UpdatedAt.After(current.UpdatedAt) {
			copy := statement
			current = &copy
		}
	}
	return current
}

func (s *MemoryStore) workloadRefsForDigestLocked(digest string) []ActiveWorkloadRef {
	seen := map[string]struct{}{}
	workloads := []ActiveWorkloadRef{}
	for _, record := range s.records {
		recordDigest := strings.TrimSpace(record.Digest)
		if recordDigest == "" {
			recordDigest = DigestFromImage(record.Image)
		}
		if recordDigest != digest {
			continue
		}
		scope := ActiveWorkloadRef{
			TenantID:    record.TenantID,
			Environment: record.Environment,
			Namespace:   record.Namespace,
			Workload:    record.Workload,
			Repo:        record.Repo,
			Image:       record.Image,
			Digest:      digest,
		}
		key := strings.Join([]string{scope.TenantID, scope.Environment, scope.Namespace, scope.Workload, scope.Repo, scope.Image}, "|")
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		workloads = append(workloads, scope)
	}
	sort.Slice(workloads, func(i, j int) bool {
		if workloads[i].Namespace == workloads[j].Namespace {
			return workloads[i].Workload < workloads[j].Workload
		}
		return workloads[i].Namespace < workloads[j].Namespace
	})
	return workloads
}

func (s *MemoryStore) imageRefForDigestLocked(digest string) string {
	for _, record := range s.records {
		recordDigest := strings.TrimSpace(record.Digest)
		if recordDigest == "" {
			recordDigest = DigestFromImage(record.Image)
		}
		if recordDigest == digest && strings.TrimSpace(record.Image) != "" {
			return strings.TrimSpace(record.Image)
		}
	}
	for _, document := range s.sbomDocuments {
		if document.ImageDigest == digest && document.ImageRef != "" {
			return document.ImageRef
		}
	}
	return ""
}

func appendScopeIfMissing(item *ActiveDigestRef, scope ActiveWorkloadRef) {
	key := strings.Join([]string{scope.TenantID, scope.Environment, scope.Namespace, scope.Workload, scope.Repo, scope.Image}, "|")
	for _, existing := range item.Scopes {
		existingKey := strings.Join([]string{existing.TenantID, existing.Environment, existing.Namespace, existing.Workload, existing.Repo, existing.Image}, "|")
		if existingKey == key {
			return
		}
	}
	item.Scopes = append(item.Scopes, scope)
}

func matchesDigestScopeFilters(workloads []ActiveWorkloadRef, tenantID, environment string) bool {
	if tenantID == "" && environment == "" {
		return true
	}
	for _, workload := range workloads {
		if tenantID != "" && workload.TenantID != tenantID {
			continue
		}
		if environment != "" && workload.Environment != environment {
			continue
		}
		return true
	}
	return false
}

func normalizeTimePointer(value *time.Time) *time.Time {
	if value == nil {
		return nil
	}
	copy := value.UTC()
	return &copy
}

func severityRank(value string) int {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "CRITICAL":
		return 5
	case "HIGH":
		return 4
	case "MEDIUM":
		return 3
	case "LOW":
		return 2
	case "UNKNOWN":
		return 1
	default:
		return 0
	}
}

func jsonRawCopy(dst json.RawMessage) json.RawMessage {
	if dst == nil {
		return json.RawMessage{}
	}
	return append(json.RawMessage(nil), dst...)
}
