package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

func (s server) runtimeDesiredStateHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin, auth.RoleService)
	if !ok {
		return
	}
	r = authorizedRequest
	if principal.Role != auth.RoleService {
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	events, err := s.listRuntimeEvents(r)
	if err != nil {
		s.writeRuntimeEventError(w, err)
		return
	}

	items := audit.DeriveRuntimeDesiredStates(events, audit.RuntimeDesiredStateFilter{
		ClusterID:    r.URL.Query().Get("cluster_id"),
		TenantID:     r.URL.Query().Get("tenant_id"),
		Namespace:    r.URL.Query().Get("namespace"),
		WorkloadKind: r.URL.Query().Get("workload_kind"),
		Workload:     r.URL.Query().Get("workload"),
		Limit:        runtimeLimit(r),
	})
	httpjson.Write(w, http.StatusOK, map[string]any{"items": items})
}

func (s server) runtimeActiveStateHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin, auth.RoleService)
	if !ok {
		return
	}
	r = authorizedRequest
	if principal.Role != auth.RoleService {
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	events, err := s.listRuntimeEvents(r)
	if err != nil {
		s.writeRuntimeEventError(w, err)
		return
	}
	items := audit.DeriveRuntimeActiveStates(events, audit.RuntimeActiveStateFilter{
		ClusterID:            r.URL.Query().Get("cluster_id"),
		TenantID:             r.URL.Query().Get("tenant_id"),
		Namespace:            r.URL.Query().Get("namespace"),
		WorkloadKind:         r.URL.Query().Get("workload_kind"),
		Workload:             r.URL.Query().Get("workload"),
		ReconciliationStatus: r.URL.Query().Get("reconciliation_status"),
		QuarantineType:       r.URL.Query().Get("quarantine_type"),
		Limit:                runtimeLimit(r),
	})
	httpjson.Write(w, http.StatusOK, map[string]any{"items": items})
}

func (s server) runtimeClosedLoopStatusHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin, auth.RoleService)
	if !ok {
		return
	}
	r = authorizedRequest
	if principal.Role != auth.RoleService {
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	events, err := s.listRuntimeEvents(r)
	if err != nil {
		s.writeRuntimeEventError(w, err)
		return
	}
	items := audit.DeriveRuntimeActiveStates(events, audit.RuntimeActiveStateFilter{
		ClusterID:            r.URL.Query().Get("cluster_id"),
		TenantID:             r.URL.Query().Get("tenant_id"),
		Namespace:            r.URL.Query().Get("namespace"),
		WorkloadKind:         r.URL.Query().Get("workload_kind"),
		Workload:             r.URL.Query().Get("workload"),
		ReconciliationStatus: r.URL.Query().Get("reconciliation_status"),
		QuarantineType:       r.URL.Query().Get("quarantine_type"),
		Limit:                max(runtimeLimit(r), 250),
	})
	httpjson.Write(w, http.StatusOK, audit.DeriveRuntimeClosedLoopStatus(items))
}

func (s server) runtimeQuarantineHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin, auth.RoleService)
	if !ok {
		return
	}
	r = authorizedRequest
	if principal.Role != auth.RoleService {
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	events, err := s.listRuntimeEvents(r)
	if err != nil {
		s.writeRuntimeEventError(w, err)
		return
	}
	items := audit.DeriveRuntimeActiveStates(events, audit.RuntimeActiveStateFilter{
		ClusterID:            r.URL.Query().Get("cluster_id"),
		TenantID:             r.URL.Query().Get("tenant_id"),
		Namespace:            r.URL.Query().Get("namespace"),
		WorkloadKind:         r.URL.Query().Get("workload_kind"),
		Workload:             r.URL.Query().Get("workload"),
		ReconciliationStatus: "quarantined",
		QuarantineType:       r.URL.Query().Get("quarantine_type"),
		Limit:                runtimeLimit(r),
	})
	httpjson.Write(w, http.StatusOK, map[string]any{"items": items})
}

func (s server) runtimeDriftFindingsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	findings, err := s.deriveRuntimeDriftFindings(r)
	if err != nil {
		s.writeRuntimeEventError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, map[string]any{"items": findings})
}

func (s server) runtimeDriftByIDHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/runtime/drift/")
	if id == "" || id == "status" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "runtime drift finding not found"})
		return
	}

	findings, err := s.deriveRuntimeDriftFindings(r)
	if err != nil {
		s.writeRuntimeEventError(w, err)
		return
	}
	for _, finding := range findings {
		if finding.ID == id {
			httpjson.Write(w, http.StatusOK, finding)
			return
		}
	}
	httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "runtime drift finding not found"})
}

func (s server) runtimeDriftStatusHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	findings, err := s.deriveRuntimeDriftFindings(r)
	if err != nil {
		s.writeRuntimeEventError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, audit.DeriveRuntimeDriftStatus(findings))
}

func (s server) deriveRuntimeDriftFindings(r *http.Request) ([]audit.RuntimeDriftFinding, error) {
	events, err := s.listRuntimeEvents(r)
	if err != nil {
		return nil, err
	}
	return audit.DeriveRuntimeDriftFindings(events, audit.RuntimeDriftFilter{
		ClusterID:    r.URL.Query().Get("cluster_id"),
		TenantID:     r.URL.Query().Get("tenant_id"),
		Namespace:    r.URL.Query().Get("namespace"),
		WorkloadKind: r.URL.Query().Get("workload_kind"),
		Workload:     r.URL.Query().Get("workload"),
		Severity:     r.URL.Query().Get("severity"),
		Status:       r.URL.Query().Get("status"),
		Limit:        runtimeLimit(r),
	}), nil
}

func (s server) listRuntimeEvents(r *http.Request) ([]audit.StoredEvent, error) {
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	return s.store.ListEvents(ctx, audit.EventFilter{
		Component: "runtime-agent",
		ClusterID: strings.TrimSpace(r.URL.Query().Get("cluster_id")),
		TenantID:  strings.TrimSpace(r.URL.Query().Get("tenant_id")),
		Limit:     max(runtimeLimit(r), 250),
	})
}

func (s server) writeRuntimeEventError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if errors.Is(err, audit.ErrInvalidFilter) {
		status = http.StatusBadRequest
	} else if errors.Is(err, context.DeadlineExceeded) {
		status = http.StatusGatewayTimeout
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}

func runtimeLimit(r *http.Request) int {
	raw := strings.TrimSpace(r.URL.Query().Get("limit"))
	if raw == "" {
		return 100
	}
	value, err := strconv.Atoi(raw)
	if err != nil || value <= 0 {
		return 100
	}
	if value > 500 {
		return 500
	}
	return value
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
