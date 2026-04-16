package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

var (
	errTenantScopeViolation = &auth.AccessError{Status: 403, Message: "requested resource is outside tenant scope"}
)

func tenantScoped(principal auth.Principal) bool {
	return principal.AuthMode != auth.ModeDisabled && principal.Role != auth.RoleService && principal.TenantID != "" && !principal.GlobalScope
}

func applyPrincipalTenantToRequest(principal auth.Principal, r *http.Request) (*http.Request, error) {
	if !tenantScoped(principal) {
		return r, nil
	}

	query := r.URL.Query()
	tenantID, err := coerceTenantScope(principal, query.Get("tenant_id"))
	if err != nil {
		return nil, err
	}
	query.Set("tenant_id", tenantID)

	cloned := r.Clone(r.Context())
	cloned.URL.RawQuery = query.Encode()
	return cloned, nil
}

func coerceTenantScope(principal auth.Principal, tenantID string) (string, error) {
	tenantID = strings.TrimSpace(tenantID)
	if !tenantScoped(principal) {
		return tenantID, nil
	}
	if tenantID == "" {
		return principal.TenantID, nil
	}
	if tenantID != principal.TenantID {
		return "", errTenantScopeViolation
	}
	return tenantID, nil
}

func applyPrincipalTenantToExceptionRequest(principal auth.Principal, request audit.ExceptionCreateRequest) (audit.ExceptionCreateRequest, error) {
	tenantID, err := coerceTenantScope(principal, request.TenantID)
	if err != nil {
		return audit.ExceptionCreateRequest{}, err
	}
	request.TenantID = tenantID
	return request, nil
}

func applyPrincipalTenantToExceptionValidation(principal auth.Principal, request audit.ExceptionValidationRequest) (audit.ExceptionValidationRequest, error) {
	tenantID, err := coerceTenantScope(principal, request.TenantID)
	if err != nil {
		return audit.ExceptionValidationRequest{}, err
	}
	request.TenantID = tenantID
	return request, nil
}

func ensureExceptionTenantAccess(principal auth.Principal, exception audit.PolicyException) error {
	if !tenantScoped(principal) {
		return nil
	}
	return ensureTenantMatch(principal, exception.TenantID)
}

func ensureTenantMatch(principal auth.Principal, tenantID string) error {
	if !tenantScoped(principal) {
		return nil
	}
	if strings.TrimSpace(tenantID) == "" || strings.TrimSpace(tenantID) != principal.TenantID {
		return errTenantScopeViolation
	}
	return nil
}

func ensureDigestTenantAccess(ctx context.Context, store audit.Store, principal auth.Principal, imageDigest string) error {
	if !tenantScoped(principal) {
		return nil
	}
	imageDigest = strings.TrimSpace(imageDigest)
	if imageDigest == "" {
		return errTenantScopeViolation
	}
	scopes, err := store.LookupDigestScopes(ctx, imageDigest, 50)
	if err != nil {
		return err
	}
	for _, scope := range scopes {
		if strings.TrimSpace(scope.TenantID) == principal.TenantID {
			return nil
		}
	}
	return errTenantScopeViolation
}
