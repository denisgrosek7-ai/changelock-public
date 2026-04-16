package audit

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *PostgresStore) IngestSBOM(ctx context.Context, request SBOMIngestRequest) (SBOMIngestResult, error) {
	request, err := NormalizeSBOMIngestRequest(request)
	if err != nil {
		return SBOMIngestResult{}, err
	}
	hash := sbomHash(request.SBOM)
	components, err := ParseSBOMComponents(request.SBOMFormat, request.SBOM, request.ImageDigest)
	if err != nil {
		return SBOMIngestResult{}, err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return SBOMIngestResult{}, err
	}
	defer tx.Rollback(ctx)

	var existingID int64
	err = tx.QueryRow(ctx, `SELECT id FROM sbom_documents WHERE image_digest = $1 AND sbom_hash = $2 ORDER BY created_at DESC LIMIT 1`, request.ImageDigest, hash).Scan(&existingID)
	if err == nil {
		var componentCount int
		if err := tx.QueryRow(ctx, `SELECT COUNT(*) FROM sbom_components WHERE image_digest = $1`, request.ImageDigest).Scan(&componentCount); err != nil {
			return SBOMIngestResult{}, err
		}
		if err := tx.Commit(ctx); err != nil {
			return SBOMIngestResult{}, err
		}
		return SBOMIngestResult{DocumentStored: false, DocumentID: existingID, ImageDigest: request.ImageDigest, SBOMHash: hash, ComponentsIngested: componentCount}, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return SBOMIngestResult{}, err
	}

	var documentID int64
	err = tx.QueryRow(ctx, `
INSERT INTO sbom_documents (image_digest, image_ref, sbom_format, source_ref, sbom_hash, raw_sbom)
VALUES ($1, $2, $3, $4, $5, $6::jsonb)
RETURNING id
`, request.ImageDigest, nullableString(request.ImageRef), request.SBOMFormat, nullableString(request.SourceRef), nullableString(hash), string(request.SBOM)).Scan(&documentID)
	if err != nil {
		return SBOMIngestResult{}, err
	}

	if _, err := tx.Exec(ctx, `DELETE FROM sbom_components WHERE image_digest = $1`, request.ImageDigest); err != nil {
		return SBOMIngestResult{}, err
	}

	for _, component := range components {
		metadata, err := nullableJSON(component.Metadata)
		if err != nil {
			return SBOMIngestResult{}, err
		}
		if _, err := tx.Exec(ctx, `
INSERT INTO sbom_components (image_digest, component_name, component_version, component_type, license, purl, metadata)
VALUES ($1, $2, $3, $4, $5, $6, $7::jsonb)
`, component.ImageDigest, component.ComponentName, nullableString(component.ComponentVersion), nullableString(component.ComponentType), nullableString(component.License), nullableString(component.PURL), metadata); err != nil {
			return SBOMIngestResult{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return SBOMIngestResult{}, err
	}
	return SBOMIngestResult{DocumentStored: true, DocumentID: documentID, ImageDigest: request.ImageDigest, SBOMHash: hash, ComponentsIngested: len(components)}, nil
}

func (s *PostgresStore) GetSBOMImage(ctx context.Context, imageDigest string, limit int) (SBOMImageResponse, error) {
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

	var document SBOMDocument
	var imageRef, sourceRef, sbomHash sql.NullString
	if err := s.pool.QueryRow(ctx, `
SELECT id, image_digest, image_ref, sbom_format, source_ref, sbom_hash, raw_sbom, created_at
FROM sbom_documents
WHERE image_digest = $1
ORDER BY created_at DESC
LIMIT 1
`, imageDigest).Scan(&document.ID, &document.ImageDigest, &imageRef, &document.SBOMFormat, &sourceRef, &sbomHash, &document.RawSBOM, &document.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SBOMImageResponse{}, ErrExceptionNotFound
		}
		return SBOMImageResponse{}, err
	}
	document.ImageRef = nullableStringValue(imageRef)
	document.SourceRef = nullableStringValue(sourceRef)
	document.SBOMHash = nullableStringValue(sbomHash)

	var componentCount int
	if err := s.pool.QueryRow(ctx, `SELECT COUNT(*) FROM sbom_components WHERE image_digest = $1`, imageDigest).Scan(&componentCount); err != nil {
		return SBOMImageResponse{}, err
	}

	rows, err := s.pool.Query(ctx, `
SELECT id, image_digest, component_name, component_version, component_type, license, purl, metadata, created_at
FROM sbom_components
WHERE image_digest = $1
ORDER BY component_name, component_version
LIMIT $2
`, imageDigest, limit)
	if err != nil {
		return SBOMImageResponse{}, err
	}
	defer rows.Close()

	components, err := pgx.CollectRows(rows, scanSBOMComponent)
	if err != nil {
		return SBOMImageResponse{}, err
	}
	return SBOMImageResponse{Document: document, ComponentCount: componentCount, Components: components}, nil
}

func (s *PostgresStore) SearchSBOMComponents(ctx context.Context, filter SBOMComponentSearchFilter) ([]SBOMComponent, error) {
	filter, err := NormalizeSBOMComponentSearchFilter(filter)
	if err != nil {
		return nil, err
	}
	conditions := []string{}
	args := []any{}
	appendCondition := func(value, expression string) {
		args = append(args, value)
		conditions = append(conditions, fmt.Sprintf("%s $%d", expression, len(args)))
	}
	if filter.ImageDigest != "" {
		appendCondition(filter.ImageDigest, "image_digest =")
	}
	if filter.ComponentName != "" {
		args = append(args, "%"+strings.ToLower(filter.ComponentName)+"%")
		conditions = append(conditions, fmt.Sprintf("LOWER(component_name) LIKE $%d", len(args)))
	}
	if filter.PURL != "" {
		args = append(args, "%"+strings.ToLower(filter.PURL)+"%")
		conditions = append(conditions, fmt.Sprintf("LOWER(purl) LIKE $%d", len(args)))
	}
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}
	args = append(args, filter.Limit)
	rows, err := s.pool.Query(ctx, `
SELECT id, image_digest, component_name, component_version, component_type, license, purl, metadata, created_at
FROM sbom_components`+whereClause+`
ORDER BY component_name, component_version, image_digest
LIMIT $`+fmt.Sprint(len(args)), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, scanSBOMComponent)
}

func (s *PostgresStore) RecordVulnerabilityScan(ctx context.Context, request VulnerabilityScanRequest) (VulnerabilityScanIngestResult, error) {
	request, err := NormalizeVulnerabilityScanRequest(request, time.Now)
	if err != nil {
		return VulnerabilityScanIngestResult{}, err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return VulnerabilityScanIngestResult{}, err
	}
	defer tx.Rollback(ctx)

	var priorSuccessfulCount int64
	if err := tx.QueryRow(ctx, `SELECT COUNT(*) FROM vulnerability_scan_runs WHERE image_digest = $1 AND status = $2`, request.ImageDigest, VulnerabilityScanStatusCompleted).Scan(&priorSuccessfulCount); err != nil {
		return VulnerabilityScanIngestResult{}, err
	}

	var summary any
	if summary, err = nullableJSON(request.Summary); err != nil {
		return VulnerabilityScanIngestResult{}, err
	}
	var run VulnerabilityScanRun
	var completedAt sql.NullTime
	var imageRef, sourceRef sql.NullString
	err = tx.QueryRow(ctx, `
INSERT INTO vulnerability_scan_runs (image_digest, image_ref, scanner, scan_mode, started_at, completed_at, status, summary, source_ref)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8::jsonb, $9)
RETURNING id, image_digest, image_ref, scanner, scan_mode, started_at, completed_at, status, summary, source_ref, created_at
`, request.ImageDigest, nullableString(request.ImageRef), request.Scanner, request.ScanMode, request.StartedAt, request.CompletedAt, request.Status, summary, nullableString(request.SourceRef)).
		Scan(&run.ID, &run.ImageDigest, &imageRef, &run.Scanner, &run.ScanMode, &run.StartedAt, &completedAt, &run.Status, &run.Summary, &sourceRef, &run.CreatedAt)
	if err != nil {
		return VulnerabilityScanIngestResult{}, err
	}
	run.ImageRef = nullableStringValue(imageRef)
	run.SourceRef = nullableStringValue(sourceRef)
	run.CompletedAt = normalizeSQLTimePointer(completedAt)

	rows, err := tx.Query(ctx, `
SELECT id, image_digest, image_ref, scan_run_id, cve_id, severity, package_name, package_version, fixed_version, purl, status, title, description, source, metadata, first_seen_at, last_seen_at
FROM vulnerability_findings
WHERE image_digest = $1
`, request.ImageDigest)
	if err != nil {
		return VulnerabilityScanIngestResult{}, err
	}
	existingFindings, err := pgx.CollectRows(rows, scanVulnerabilityFinding)
	if err != nil {
		return VulnerabilityScanIngestResult{}, err
	}
	existingByKey := map[string]VulnerabilityFinding{}
	for _, finding := range existingFindings {
		key := vulnerabilityFindingKey(request.ImageDigest, VulnerabilityFindingInput{
			CVEID:          finding.CVEID,
			PackageName:    finding.PackageName,
			PackageVersion: finding.PackageVersion,
			PURL:           finding.PURL,
		})
		existingByKey[key] = finding
	}

	now := time.Now().UTC()
	seenKeys := map[string]struct{}{}
	findings := make([]VulnerabilityFinding, 0, len(request.Findings))
	newFindings := []VulnerabilityFinding{}
	for _, input := range request.Findings {
		key := vulnerabilityFindingKey(request.ImageDigest, input)
		seenKeys[key] = struct{}{}
		existing, ok := existingByKey[key]
		isNew := !ok || existing.Status == VulnerabilityFindingStatusResolved
		if _, err := tx.Exec(ctx, `
INSERT INTO vulnerability_findings (
  image_digest, image_ref, scan_run_id, finding_key, cve_id, severity, package_name, package_version, fixed_version,
  purl, status, title, description, source, metadata, first_seen_at, last_seen_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15::jsonb, $16, $17)
ON CONFLICT (finding_key) DO UPDATE SET
  image_digest = EXCLUDED.image_digest,
  image_ref = EXCLUDED.image_ref,
  scan_run_id = EXCLUDED.scan_run_id,
  severity = EXCLUDED.severity,
  package_name = EXCLUDED.package_name,
  package_version = EXCLUDED.package_version,
  fixed_version = EXCLUDED.fixed_version,
  purl = EXCLUDED.purl,
  status = EXCLUDED.status,
  title = EXCLUDED.title,
  description = EXCLUDED.description,
  source = EXCLUDED.source,
  metadata = EXCLUDED.metadata,
  last_seen_at = EXCLUDED.last_seen_at
`, request.ImageDigest, nullableString(firstNonEmpty(request.ImageRef, existing.ImageRef)), run.ID, key, input.CVEID, nullableString(input.Severity), nullableString(input.PackageName), nullableString(input.PackageVersion), nullableString(input.FixedVersion), nullableString(input.PURL), VulnerabilityFindingStatusOpen, nullableString(input.Title), nullableString(input.Description), nullableString(input.Source), string(normalizeMetadata(input.Metadata)), coalesceFindingFirstSeen(existing.FirstSeenAt, now, ok), now); err != nil {
			return VulnerabilityScanIngestResult{}, err
		}

		finding, err := s.lookupFindingInTx(ctx, tx, key)
		if err != nil {
			return VulnerabilityScanIngestResult{}, err
		}
		findings = append(findings, finding)
		if isNew {
			newFindings = append(newFindings, finding)
		}
	}

	for key, existing := range existingByKey {
		if _, ok := seenKeys[key]; ok {
			continue
		}
		if existing.Status == VulnerabilityFindingStatusResolved {
			continue
		}
		if _, err := tx.Exec(ctx, `
UPDATE vulnerability_findings
SET status = $2, scan_run_id = $3, last_seen_at = $4
WHERE finding_key = $1
`, key, VulnerabilityFindingStatusResolved, run.ID, now); err != nil {
			return VulnerabilityScanIngestResult{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return VulnerabilityScanIngestResult{}, err
	}
	return VulnerabilityScanIngestResult{
		Run:                   run,
		Findings:              findings,
		NewFindings:           newFindings,
		HadPriorSuccessfulRun: priorSuccessfulCount > 0,
	}, nil
}

func (s *PostgresStore) ListActiveVulnerabilities(ctx context.Context, filter VulnerabilityActiveFilter) ([]VulnerabilityFinding, error) {
	filter = NormalizeVulnerabilityActiveFilter(filter)
	query, args := buildActiveVulnerabilitiesQuery(filter)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, scanVulnerabilityFindingWithDecision)
}

func (s *PostgresStore) VulnerabilityBlastRadius(ctx context.Context, filter VulnerabilityBlastRadiusFilter) (VulnerabilityBlastRadiusResponse, error) {
	filter, err := NormalizeVulnerabilityBlastRadiusFilter(filter)
	if err != nil {
		return VulnerabilityBlastRadiusResponse{}, err
	}

	items := map[string]*VulnerabilityBlastRadiusItem{}
	now := time.Now().UTC()
	if filter.CVEID != "" {
		findings, err := s.ListActiveVulnerabilities(ctx, VulnerabilityActiveFilter{
			CVEID:             filter.CVEID,
			Limit:             filter.Limit,
			IncludeSuppressed: true,
		})
		if err != nil {
			return VulnerabilityBlastRadiusResponse{}, err
		}
		for _, finding := range findings {
			if finding.Status != VulnerabilityFindingStatusOpen {
				continue
			}
			item := items[finding.ImageDigest]
			if item == nil {
				item = &VulnerabilityBlastRadiusItem{
					ImageDigest: finding.ImageDigest,
					ImageRef:    finding.ImageRef,
					Findings:    []VulnerabilityFinding{},
					Workloads:   []ActiveWorkloadRef{},
				}
				items[finding.ImageDigest] = item
			}
			if item.ImageRef == "" {
				item.ImageRef = finding.ImageRef
			}
			if finding.Decision != nil && !activeDecisionApplies(finding.Decision, now) {
				finding.Decision = nil
			}
			item.Findings = append(item.Findings, finding)
		}
	}

	if filter.ComponentName != "" || filter.PURL != "" {
		components, err := s.SearchSBOMComponents(ctx, SBOMComponentSearchFilter{
			ComponentName: filter.ComponentName,
			PURL:          filter.PURL,
			Limit:         filter.Limit,
		})
		if err != nil {
			return VulnerabilityBlastRadiusResponse{}, err
		}
		for _, component := range components {
			item := items[component.ImageDigest]
			if item == nil {
				item = &VulnerabilityBlastRadiusItem{
					ImageDigest: component.ImageDigest,
					ImageRef:    s.lookupImageRefForDigest(ctx, component.ImageDigest),
					Findings:    []VulnerabilityFinding{},
					Workloads:   []ActiveWorkloadRef{},
				}
				items[component.ImageDigest] = item
			}
		}
	}

	for digest, item := range items {
		workloads, err := s.lookupWorkloadsForDigest(ctx, digest, filter.Limit)
		if err != nil {
			return VulnerabilityBlastRadiusResponse{}, err
		}
		item.Workloads = workloads
	}

	results := make([]VulnerabilityBlastRadiusItem, 0, len(items))
	for _, item := range items {
		results = append(results, *item)
	}
	sort.Slice(results, func(i, j int) bool { return results[i].ImageDigest < results[j].ImageDigest })
	if len(results) > filter.Limit {
		results = results[:filter.Limit]
	}
	return VulnerabilityBlastRadiusResponse{
		Items: results,
		AppliedFilters: map[string]string{
			"cve_id":         filter.CVEID,
			"component_name": filter.ComponentName,
			"purl":           filter.PURL,
		},
	}, nil
}

func (s *PostgresStore) VulnerabilityTimeline(ctx context.Context, filter VulnerabilityTimelineFilter) (VulnerabilityTimelineResponse, error) {
	filter, err := NormalizeVulnerabilityTimelineFilter(filter)
	if err != nil {
		return VulnerabilityTimelineResponse{}, err
	}
	rows, err := s.pool.Query(ctx, `
SELECT
  f.image_digest,
  f.cve_id,
  f.package_name,
  f.package_version,
  f.severity,
  f.status,
  f.first_seen_at,
  f.last_seen_at,
  d.id,
  d.image_digest,
  d.cve_id,
  d.decision,
  d.justification,
  d.decided_by,
  d.expires_at,
  d.active,
  d.metadata,
  d.created_at,
  d.updated_at
FROM vulnerability_findings f
LEFT JOIN LATERAL (
  SELECT *
  FROM vulnerability_decisions
  WHERE image_digest = f.image_digest
    AND cve_id = f.cve_id
    AND active = TRUE
    AND (expires_at IS NULL OR expires_at > now())
  ORDER BY created_at DESC
  LIMIT 1
) d ON TRUE
WHERE f.image_digest = $1
  AND f.cve_id = $2
  AND f.last_seen_at >= now() - ($3 * interval '1 day')
ORDER BY f.first_seen_at ASC
`, filter.ImageDigest, filter.CVEID, filter.WindowDays)
	if err != nil {
		return VulnerabilityTimelineResponse{}, err
	}
	defer rows.Close()

	items := []VulnerabilityTimelineEntry{}
	for rows.Next() {
		entry, err := scanVulnerabilityTimelineEntry(rows)
		if err != nil {
			return VulnerabilityTimelineResponse{}, err
		}
		items = append(items, entry)
	}
	if rows.Err() != nil {
		return VulnerabilityTimelineResponse{}, rows.Err()
	}
	return VulnerabilityTimelineResponse{
		Items: items,
		AppliedFilters: map[string]string{
			"image_digest": filter.ImageDigest,
			"cve_id":       filter.CVEID,
			"window_days":  fmt.Sprint(filter.WindowDays),
		},
	}, nil
}

func (s *PostgresStore) ListVulnerabilityDecisions(ctx context.Context, filter VulnerabilityDecisionFilter) ([]VulnerabilityDecision, error) {
	filter = NormalizeVulnerabilityDecisionFilter(filter)
	conditions := []string{}
	args := []any{}
	appendCondition := func(value any, expression string) {
		args = append(args, value)
		conditions = append(conditions, fmt.Sprintf("%s $%d", expression, len(args)))
	}
	if filter.ImageDigest != "" {
		appendCondition(filter.ImageDigest, "image_digest =")
	}
	if filter.CVEID != "" {
		appendCondition(filter.CVEID, "cve_id =")
	}
	if filter.Active != nil {
		appendCondition(*filter.Active, "active =")
	}
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}
	args = append(args, filter.Limit)
	rows, err := s.pool.Query(ctx, `
SELECT id, image_digest, cve_id, decision, justification, decided_by, expires_at, active, metadata, created_at, updated_at
FROM vulnerability_decisions`+whereClause+`
ORDER BY created_at DESC
LIMIT $`+fmt.Sprint(len(args)), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, scanVulnerabilityDecision)
}

func (s *PostgresStore) CreateVulnerabilityDecision(ctx context.Context, request VulnerabilityDecisionCreateRequest, decidedBy string) (VulnerabilityDecision, error) {
	request, err := NormalizeVulnerabilityDecisionCreateRequest(request, time.Now)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	decidedBy = strings.TrimSpace(decidedBy)
	if decidedBy == "" {
		return VulnerabilityDecision{}, fmt.Errorf("%w: decided_by is required", ErrInvalidException)
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	defer tx.Rollback(ctx)

	now := time.Now().UTC()
	if _, err := tx.Exec(ctx, `
UPDATE vulnerability_decisions
SET active = FALSE, updated_at = $3
WHERE image_digest = $1 AND cve_id = $2 AND active = TRUE
`, request.ImageDigest, request.CVEID, now); err != nil {
		return VulnerabilityDecision{}, err
	}

	metadata, err := nullableJSON(request.Metadata)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	rows, err := tx.Query(ctx, `
INSERT INTO vulnerability_decisions (image_digest, cve_id, decision, justification, decided_by, expires_at, active, metadata, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, TRUE, $7::jsonb, $8, $8)
RETURNING id, image_digest, cve_id, decision, justification, decided_by, expires_at, active, metadata, created_at, updated_at
`, request.ImageDigest, request.CVEID, request.Decision, request.Justification, decidedBy, request.ExpiresAt, metadata, now)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	defer rows.Close()
	decision, err := pgx.CollectOneRow(rows, scanVulnerabilityDecision)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return VulnerabilityDecision{}, err
	}
	return decision, nil
}

func (s *PostgresStore) DeactivateVulnerabilityDecision(ctx context.Context, decisionID int64) (VulnerabilityDecision, error) {
	rows, err := s.pool.Query(ctx, `
UPDATE vulnerability_decisions
SET active = FALSE, updated_at = now()
WHERE id = $1
RETURNING id, image_digest, cve_id, decision, justification, decided_by, expires_at, active, metadata, created_at, updated_at
`, decisionID)
	if err != nil {
		return VulnerabilityDecision{}, err
	}
	defer rows.Close()
	decision, err := pgx.CollectOneRow(rows, scanVulnerabilityDecision)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return VulnerabilityDecision{}, ErrExceptionNotFound
		}
		return VulnerabilityDecision{}, err
	}
	return decision, nil
}

func (s *PostgresStore) ListActiveDigests(ctx context.Context, windowDays int, limit int) ([]ActiveDigestRef, error) {
	if windowDays <= 0 {
		windowDays = 30
	}
	if limit <= 0 {
		limit = 100
	}

	rows, err := s.pool.Query(ctx, `
SELECT digest, MAX(NULLIF(image, '')) AS image_ref, MAX(NULLIF(repo, '')) AS repo
FROM audit_events
WHERE digest IS NOT NULL
  AND digest <> ''
  AND received_at >= now() - ($1::text || ' days')::interval
GROUP BY digest
ORDER BY digest
LIMIT $2
`, windowDays, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []ActiveDigestRef{}
	for rows.Next() {
		var item ActiveDigestRef
		var imageRef, repo sql.NullString
		if err := rows.Scan(&item.ImageDigest, &imageRef, &repo); err != nil {
			return nil, err
		}
		item.ImageRef = nullableStringValue(imageRef)
		item.Repo = nullableStringValue(repo)
		scopes, err := s.lookupWorkloadsForDigest(ctx, item.ImageDigest, 20)
		if err != nil {
			return nil, err
		}
		item.Scopes = scopes
		results = append(results, item)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return results, nil
}

func (s *PostgresStore) lookupFindingInTx(ctx context.Context, tx pgx.Tx, findingKey string) (VulnerabilityFinding, error) {
	rows, err := tx.Query(ctx, `
SELECT id, image_digest, image_ref, scan_run_id, cve_id, severity, package_name, package_version, fixed_version, purl, status, title, description, source, metadata, first_seen_at, last_seen_at
FROM vulnerability_findings
WHERE finding_key = $1
`, findingKey)
	if err != nil {
		return VulnerabilityFinding{}, err
	}
	defer rows.Close()
	return pgx.CollectOneRow(rows, scanVulnerabilityFinding)
}

func buildActiveVulnerabilitiesQuery(filter VulnerabilityActiveFilter) (string, []any) {
	args := []any{}
	conditions := []string{"f.status = 'OPEN'"}
	appendCondition := func(value any, expression string) {
		args = append(args, value)
		conditions = append(conditions, fmt.Sprintf("%s $%d", expression, len(args)))
	}
	if filter.Severity != "" {
		appendCondition(filter.Severity, "UPPER(f.severity) =")
	}
	if filter.CVEID != "" {
		appendCondition(filter.CVEID, "f.cve_id =")
	}
	if filter.ImageDigest != "" {
		appendCondition(filter.ImageDigest, "f.image_digest =")
	}
	if filter.ComponentName != "" {
		args = append(args, "%"+strings.ToLower(filter.ComponentName)+"%")
		conditions = append(conditions, fmt.Sprintf("(LOWER(f.package_name) LIKE $%d OR LOWER(COALESCE(f.purl, '')) LIKE $%d)", len(args), len(args)))
	}
	if filter.TenantID != "" {
		args = append(args, filter.TenantID)
		conditions = append(conditions, fmt.Sprintf(`EXISTS (
  SELECT 1 FROM audit_events ae
  WHERE ae.digest = f.image_digest AND ae.tenant_id = $%d
)`, len(args)))
	}
	if filter.Environment != "" {
		args = append(args, filter.Environment)
		conditions = append(conditions, fmt.Sprintf(`EXISTS (
  SELECT 1 FROM audit_events ae
  WHERE ae.digest = f.image_digest AND ae.environment = $%d
)`, len(args)))
	}
	if !filter.IncludeSuppressed {
		conditions = append(conditions, `COALESCE(d.decision, '') <> 'NOT_AFFECTED'`)
	}
	args = append(args, filter.Limit)
	return `
SELECT
  f.id, f.image_digest, COALESCE(f.image_ref, ''), f.scan_run_id, f.cve_id, COALESCE(f.severity, ''), COALESCE(f.package_name, ''), COALESCE(f.package_version, ''),
  COALESCE(f.fixed_version, ''), COALESCE(f.purl, ''), f.status, COALESCE(f.title, ''), COALESCE(f.description, ''), COALESCE(f.source, ''), f.metadata, f.first_seen_at, f.last_seen_at,
  d.id, d.image_digest, d.cve_id, d.decision, d.justification, d.decided_by, d.expires_at, d.active, d.metadata, d.created_at, d.updated_at
FROM vulnerability_findings f
LEFT JOIN LATERAL (
  SELECT *
  FROM vulnerability_decisions
  WHERE image_digest = f.image_digest
    AND cve_id = f.cve_id
    AND active = TRUE
    AND (expires_at IS NULL OR expires_at > now())
  ORDER BY created_at DESC
  LIMIT 1
) d ON TRUE
WHERE ` + strings.Join(conditions, " AND ") + `
ORDER BY
  CASE UPPER(COALESCE(f.severity, ''))
    WHEN 'CRITICAL' THEN 5
    WHEN 'HIGH' THEN 4
    WHEN 'MEDIUM' THEN 3
    WHEN 'LOW' THEN 2
    WHEN 'UNKNOWN' THEN 1
    ELSE 0
  END DESC,
  f.cve_id,
  f.image_digest
LIMIT $` + fmt.Sprint(len(args)), args
}

func (s *PostgresStore) lookupWorkloadsForDigest(ctx context.Context, digest string, limit int) ([]ActiveWorkloadRef, error) {
	if limit <= 0 {
		limit = 20
	}
	rows, err := s.pool.Query(ctx, `
SELECT DISTINCT tenant_id, environment, namespace, workload, repo, image, digest
FROM audit_events
WHERE digest = $1
ORDER BY namespace, workload, repo
LIMIT $2
`, digest, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	workloads := []ActiveWorkloadRef{}
	for rows.Next() {
		var workload ActiveWorkloadRef
		if err := rows.Scan(&workload.TenantID, &workload.Environment, &workload.Namespace, &workload.Workload, &workload.Repo, &workload.Image, &workload.Digest); err != nil {
			return nil, err
		}
		workloads = append(workloads, workload)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return workloads, nil
}

func (s *PostgresStore) lookupImageRefForDigest(ctx context.Context, digest string) string {
	var image sql.NullString
	if err := s.pool.QueryRow(ctx, `
SELECT COALESCE(MAX(NULLIF(image, '')), '')
FROM audit_events
WHERE digest = $1
`, digest).Scan(&image); err == nil {
		return nullableStringValue(image)
	}
	return ""
}

func scanSBOMComponent(row pgx.CollectableRow) (SBOMComponent, error) {
	var component SBOMComponent
	var version, componentType, license, purl sql.NullString
	if err := row.Scan(&component.ID, &component.ImageDigest, &component.ComponentName, &version, &componentType, &license, &purl, &component.Metadata, &component.CreatedAt); err != nil {
		return SBOMComponent{}, err
	}
	component.ComponentVersion = nullableStringValue(version)
	component.ComponentType = nullableStringValue(componentType)
	component.License = nullableStringValue(license)
	component.PURL = nullableStringValue(purl)
	return component, nil
}

func scanVulnerabilityFinding(row pgx.CollectableRow) (VulnerabilityFinding, error) {
	var finding VulnerabilityFinding
	var imageRef, severity, packageName, packageVersion, fixedVersion, purl, title, description, source sql.NullString
	if err := row.Scan(
		&finding.ID,
		&finding.ImageDigest,
		&imageRef,
		&finding.ScanRunID,
		&finding.CVEID,
		&severity,
		&packageName,
		&packageVersion,
		&fixedVersion,
		&purl,
		&finding.Status,
		&title,
		&description,
		&source,
		&finding.Metadata,
		&finding.FirstSeenAt,
		&finding.LastSeenAt,
	); err != nil {
		return VulnerabilityFinding{}, err
	}
	finding.ImageRef = nullableStringValue(imageRef)
	finding.Severity = nullableStringValue(severity)
	finding.PackageName = nullableStringValue(packageName)
	finding.PackageVersion = nullableStringValue(packageVersion)
	finding.FixedVersion = nullableStringValue(fixedVersion)
	finding.PURL = nullableStringValue(purl)
	finding.Title = nullableStringValue(title)
	finding.Description = nullableStringValue(description)
	finding.Source = nullableStringValue(source)
	return finding, nil
}

func scanVulnerabilityFindingWithDecision(row pgx.CollectableRow) (VulnerabilityFinding, error) {
	var finding VulnerabilityFinding
	var imageRef, severity, packageName, packageVersion, fixedVersion, purl, title, description, source sql.NullString
	decision, err := scanOptionalDecisionColumns(&finding, row,
		&finding.ID,
		&finding.ImageDigest,
		&imageRef,
		&finding.ScanRunID,
		&finding.CVEID,
		&severity,
		&packageName,
		&packageVersion,
		&fixedVersion,
		&purl,
		&finding.Status,
		&title,
		&description,
		&source,
		&finding.Metadata,
		&finding.FirstSeenAt,
		&finding.LastSeenAt,
	)
	if err != nil {
		return VulnerabilityFinding{}, err
	}
	finding.ImageRef = nullableStringValue(imageRef)
	finding.Severity = nullableStringValue(severity)
	finding.PackageName = nullableStringValue(packageName)
	finding.PackageVersion = nullableStringValue(packageVersion)
	finding.FixedVersion = nullableStringValue(fixedVersion)
	finding.PURL = nullableStringValue(purl)
	finding.Title = nullableStringValue(title)
	finding.Description = nullableStringValue(description)
	finding.Source = nullableStringValue(source)
	finding.Decision = decision
	return finding, nil
}

func scanOptionalDecision(row pgx.CollectableRow) (*VulnerabilityDecision, error) {
	return scanOptionalDecisionColumns(nil, row)
}

func scanOptionalDecisionColumns(_ *VulnerabilityFinding, row pgx.CollectableRow, leading ...any) (*VulnerabilityDecision, error) {
	var decisionID sql.NullInt64
	var imageDigest, cveID, decisionValue, justification, decidedBy sql.NullString
	var expiresAt sql.NullTime
	var active sql.NullBool
	var metadata json.RawMessage
	var createdAt, updatedAt sql.NullTime
	scanArgs := append(append([]any{}, leading...),
		&decisionID,
		&imageDigest,
		&cveID,
		&decisionValue,
		&justification,
		&decidedBy,
		&expiresAt,
		&active,
		&metadata,
		&createdAt,
		&updatedAt,
	)
	if err := row.Scan(scanArgs...); err != nil {
		return nil, err
	}
	if !decisionID.Valid {
		return nil, nil
	}
	decision := VulnerabilityDecision{
		ID:            decisionID.Int64,
		ImageDigest:   nullableStringValue(imageDigest),
		CVEID:         nullableStringValue(cveID),
		Decision:      nullableStringValue(decisionValue),
		Justification: nullableStringValue(justification),
		DecidedBy:     nullableStringValue(decidedBy),
		ExpiresAt:     normalizeSQLTimePointer(expiresAt),
		Active:        active.Valid && active.Bool,
		Metadata:      metadata,
	}
	if createdAt.Valid {
		decision.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		decision.UpdatedAt = updatedAt.Time
	}
	return &decision, nil
}

func scanVulnerabilityTimelineEntry(row pgx.CollectableRow) (VulnerabilityTimelineEntry, error) {
	var entry VulnerabilityTimelineEntry
	var packageName, packageVersion, severity sql.NullString
	decision, err := scanOptionalDecisionColumns(nil, row,
		&entry.ImageDigest,
		&entry.CVEID,
		&packageName,
		&packageVersion,
		&severity,
		&entry.Status,
		&entry.FirstSeenAt,
		&entry.LastSeenAt,
	)
	if err != nil {
		return VulnerabilityTimelineEntry{}, err
	}
	entry.PackageName = nullableStringValue(packageName)
	entry.PackageVersion = nullableStringValue(packageVersion)
	entry.Severity = nullableStringValue(severity)
	entry.Decision = decision
	return entry, nil
}

func scanVulnerabilityDecision(row pgx.CollectableRow) (VulnerabilityDecision, error) {
	var decision VulnerabilityDecision
	var expiresAt sql.NullTime
	if err := row.Scan(&decision.ID, &decision.ImageDigest, &decision.CVEID, &decision.Decision, &decision.Justification, &decision.DecidedBy, &expiresAt, &decision.Active, &decision.Metadata, &decision.CreatedAt, &decision.UpdatedAt); err != nil {
		return VulnerabilityDecision{}, err
	}
	decision.ExpiresAt = normalizeSQLTimePointer(expiresAt)
	return decision, nil
}

func normalizeSQLTimePointer(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}
	copy := value.Time.UTC()
	return &copy
}

func coalesceFindingFirstSeen(existing time.Time, now time.Time, ok bool) time.Time {
	if ok && !existing.IsZero() {
		return existing
	}
	return now
}
