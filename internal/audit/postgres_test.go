package audit

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestAppendConditions(t *testing.T) {
	tests := []struct {
		name       string
		where      string
		conditions []string
		want       string
	}{
		{
			name:       "start new where clause",
			where:      "",
			conditions: []string{"decision = 'DENY'"},
			want:       " WHERE decision = 'DENY'",
		},
		{
			name:       "append to existing where clause",
			where:      " WHERE tenant_id = $1",
			conditions: []string{"decision = 'DENY'", "event_type = 'runtime_drift_result'"},
			want:       " WHERE tenant_id = $1 AND decision = 'DENY' AND event_type = 'runtime_drift_result'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendConditions(tt.where, tt.conditions...); got != tt.want {
				t.Fatalf("appendConditions() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestScanPolicyExceptionHandlesNullScopeFields(t *testing.T) {
	now := time.Date(2026, 4, 14, 12, 0, 0, 0, time.UTC)
	row := fakePolicyExceptionRow{
		values: []any{
			int64(1),
			"EX-2026-001",
			ExceptionTypeBreakGlass,
			ExceptionStatusApproved,
			nil,
			"prod",
			"acme-prod",
			nil,
			nil,
			nil,
			"P0 production restore",
			"INC-1234",
			nil,
			nil,
			"security-lead",
			now,
			nil,
			nil,
			nil,
			now,
			now.Add(2 * time.Hour),
			true,
			now,
			nil,
			json.RawMessage(`{}`),
		},
	}

	exception, err := scanPolicyException(row)
	if err != nil {
		t.Fatalf("scanPolicyException() error = %v", err)
	}
	if exception.TenantID != "" || exception.Repo != "" || exception.ImageDigest != "" || exception.CVEID != "" {
		t.Fatalf("expected nullable scope fields to be empty strings, got %#v", exception)
	}
	if exception.Environment != "prod" || exception.Namespace != "acme-prod" {
		t.Fatalf("unexpected non-null fields %#v", exception)
	}
}

type fakePolicyExceptionRow struct {
	values []any
}

func (r fakePolicyExceptionRow) Scan(dest ...any) error {
	if len(dest) != len(r.values) {
		return fmt.Errorf("scan arg mismatch: got %d want %d", len(dest), len(r.values))
	}
	for i, target := range dest {
		value := r.values[i]
		switch d := target.(type) {
		case *int64:
			if value == nil {
				*d = 0
			} else {
				*d = value.(int64)
			}
		case *string:
			if value == nil {
				*d = ""
			} else {
				*d = value.(string)
			}
		case *sql.NullString:
			if value == nil {
				*d = sql.NullString{}
			} else {
				*d = sql.NullString{String: value.(string), Valid: true}
			}
		case *sql.NullTime:
			if value == nil {
				*d = sql.NullTime{}
			} else {
				*d = sql.NullTime{Time: value.(time.Time), Valid: true}
			}
		case *time.Time:
			*d = value.(time.Time)
		case *bool:
			*d = value.(bool)
		case *json.RawMessage:
			if value == nil {
				*d = nil
			} else {
				switch encoded := value.(type) {
				case json.RawMessage:
					*d = append((*d)[:0], encoded...)
				case []byte:
					*d = append((*d)[:0], encoded...)
				default:
					return fmt.Errorf("unsupported json raw message source %T", value)
				}
			}
		case *[]byte:
			if value == nil {
				*d = nil
			} else {
				switch encoded := value.(type) {
				case []byte:
					*d = append((*d)[:0], encoded...)
				case json.RawMessage:
					*d = append((*d)[:0], encoded...)
				default:
					return fmt.Errorf("unsupported byte slice source %T", value)
				}
			}
		default:
			return fmt.Errorf("unsupported scan target %T", target)
		}
	}
	return nil
}
