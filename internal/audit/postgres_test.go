package audit

import "testing"

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
