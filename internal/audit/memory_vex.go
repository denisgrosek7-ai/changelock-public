package audit

import (
	"context"
	"sort"
	"strings"

	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

func (s *MemoryStore) ListVEXStatements(_ context.Context, filter internalvex.Filter) ([]internalvex.Statement, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter, err := internalvex.NormalizeFilter(filter)
	if err != nil {
		return nil, err
	}
	results := make([]internalvex.Statement, 0, len(s.vexStatements))
	for _, statement := range s.vexStatements {
		if filter.VulnerabilityID != "" && statement.VulnerabilityID != filter.VulnerabilityID {
			continue
		}
		if filter.ImageDigest != "" && statement.Scope.ImageDigest != filter.ImageDigest {
			continue
		}
		if filter.PackageName != "" && !strings.EqualFold(statement.Scope.PackageName, filter.PackageName) {
			continue
		}
		if filter.PURL != "" && statement.Scope.PURL != filter.PURL {
			continue
		}
		if filter.Repo != "" && statement.Scope.Repo != filter.Repo {
			continue
		}
		if filter.Workload != "" && statement.Scope.Workload != filter.Workload {
			continue
		}
		if filter.TenantID != "" && statement.Scope.TenantID != filter.TenantID {
			continue
		}
		if filter.ClusterID != "" && statement.Scope.ClusterID != filter.ClusterID {
			continue
		}
		if filter.Environment != "" && statement.Scope.Environment != filter.Environment {
			continue
		}
		if filter.Namespace != "" && statement.Scope.Namespace != filter.Namespace {
			continue
		}
		if filter.SourceFormat != "" && statement.SourceFormat != filter.SourceFormat {
			continue
		}
		if filter.SourceRef != "" && statement.SourceRef != filter.SourceRef {
			continue
		}
		if filter.Status != "" && statement.Status != filter.Status {
			continue
		}
		if filter.Active != nil {
			active := statement.Active && statement.RevokedAt == nil && (statement.ExpiresAt == nil || statement.ExpiresAt.After(s.now().UTC()))
			if active != *filter.Active {
				continue
			}
		}
		results = append(results, cloneVEXStatement(statement))
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].UpdatedAt.After(results[j].UpdatedAt)
	})
	if len(results) > filter.Limit {
		results = results[:filter.Limit]
	}
	return results, nil
}

func (s *MemoryStore) GetVEXStatement(_ context.Context, statementID int64) (internalvex.Statement, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	statement, ok := s.vexStatements[statementID]
	if !ok {
		return internalvex.Statement{}, ErrExceptionNotFound
	}
	return cloneVEXStatement(statement), nil
}

func (s *MemoryStore) CreateVEXStatement(_ context.Context, request internalvex.CreateRequest, createdBy string) (internalvex.Statement, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.createVEXStatementLocked(request, createdBy)
}

func (s *MemoryStore) createVEXStatementLocked(request internalvex.CreateRequest, createdBy string) (internalvex.Statement, error) {

	createdBy = strings.TrimSpace(createdBy)
	if createdBy == "" {
		return internalvex.Statement{}, ErrInvalidException
	}
	request, err := internalvex.NormalizeCreateRequest(request, s.now)
	if err != nil {
		return internalvex.Statement{}, err
	}
	now := s.now().UTC()
	key := internalvex.StatementIdentityKey(internalvex.Statement{
		SourceFormat:    request.SourceFormat,
		SourceRef:       request.SourceRef,
		VulnerabilityID: request.VulnerabilityID,
		Scope:           request.Scope,
	})
	for id, existing := range s.vexStatements {
		if existing.StatementKey != key {
			continue
		}
		existing.SourceFormat = request.SourceFormat
		existing.SourceRef = request.SourceRef
		existing.VulnerabilityID = request.VulnerabilityID
		existing.Scope = request.Scope
		existing.Status = request.Status
		existing.Justification = request.Justification
		existing.ActionStatement = request.ActionStatement
		existing.ImpactStatement = request.ImpactStatement
		existing.FixedVersion = request.FixedVersion
		existing.ExpiresAt = normalizeTimePointer(request.ExpiresAt)
		existing.Metadata = cloneLegacyMetadata(request.Metadata)
		existing.Active = true
		existing.RevokedAt = nil
		existing.RevokedBy = ""
		existing.UpdatedBy = createdBy
		existing.UpdatedAt = now
		s.vexStatements[id] = existing
		return cloneVEXStatement(existing), nil
	}

	statement := internalvex.Statement{
		ID:              s.nextVEXStatementID,
		StatementKey:    key,
		SourceFormat:    request.SourceFormat,
		SourceRef:       request.SourceRef,
		VulnerabilityID: request.VulnerabilityID,
		Scope:           request.Scope,
		Status:          request.Status,
		Justification:   request.Justification,
		ActionStatement: request.ActionStatement,
		ImpactStatement: request.ImpactStatement,
		FixedVersion:    request.FixedVersion,
		CreatedBy:       createdBy,
		UpdatedBy:       createdBy,
		ExpiresAt:       normalizeTimePointer(request.ExpiresAt),
		Active:          true,
		Metadata:        cloneLegacyMetadata(request.Metadata),
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	s.nextVEXStatementID++
	s.vexStatements[statement.ID] = statement
	return cloneVEXStatement(statement), nil
}

func (s *MemoryStore) RevokeVEXStatement(_ context.Context, statementID int64, revokedBy string) (internalvex.Statement, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.revokeVEXStatementLocked(statementID, revokedBy)
}

func (s *MemoryStore) revokeVEXStatementLocked(statementID int64, revokedBy string) (internalvex.Statement, error) {

	statement, ok := s.vexStatements[statementID]
	if !ok {
		return internalvex.Statement{}, ErrExceptionNotFound
	}
	now := s.now().UTC()
	statement.Active = false
	statement.RevokedAt = timePointer(now)
	statement.RevokedBy = strings.TrimSpace(revokedBy)
	statement.UpdatedBy = strings.TrimSpace(revokedBy)
	statement.UpdatedAt = now
	s.vexStatements[statementID] = statement
	return cloneVEXStatement(statement), nil
}

func cloneVEXStatement(statement internalvex.Statement) internalvex.Statement {
	if statement.ExpiresAt != nil {
		expiresAt := statement.ExpiresAt.UTC()
		statement.ExpiresAt = &expiresAt
	}
	if statement.RevokedAt != nil {
		revokedAt := statement.RevokedAt.UTC()
		statement.RevokedAt = &revokedAt
	}
	statement.Metadata = cloneLegacyMetadata(statement.Metadata)
	return statement
}
