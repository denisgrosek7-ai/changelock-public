package auth

import "sort"

type Description struct {
	Mode                     string              `json:"mode"`
	IdentityProvider         string              `json:"identity_provider"`
	RoleClaim                string              `json:"role_claim,omitempty"`
	TenantClaim              string              `json:"tenant_claim,omitempty"`
	SubjectClaim             string              `json:"subject_claim,omitempty"`
	EmailClaim               string              `json:"email_claim,omitempty"`
	Issuer                   string              `json:"issuer,omitempty"`
	JWKSURL                  string              `json:"jwks_url,omitempty"`
	RequireTenantScope       bool                `json:"require_tenant_scope,omitempty"`
	AllowGlobalSecurityAdmin bool                `json:"allow_global_security_admin,omitempty"`
	RoleBindings             map[string][]string `json:"role_bindings,omitempty"`
	StaticTokenRoleCounts    map[string]int      `json:"static_token_role_counts,omitempty"`
	AuditActorAttribution    bool                `json:"audit_actor_attribution"`
	ApprovalActorAttribution bool                `json:"approval_actor_attribution"`
}

func (c Config) Describe() Description {
	description := Description{
		Mode:                     c.Mode,
		AuditActorAttribution:    true,
		ApprovalActorAttribution: true,
	}
	switch c.Mode {
	case ModeDisabled:
		description.IdentityProvider = "disabled"
	case ModeStaticToken:
		description.IdentityProvider = "static_token"
		counts := map[string]int{}
		for _, token := range c.tokens {
			counts[token.Role]++
		}
		description.StaticTokenRoleCounts = counts
	case ModeOIDCJWT:
		description.IdentityProvider = "oidc"
		if c.oidc == nil {
			return description
		}
		description.RoleClaim = c.oidc.roleClaim
		description.TenantClaim = c.oidc.tenantClaim
		description.SubjectClaim = c.oidc.subjectClaim
		description.EmailClaim = c.oidc.emailClaim
		description.Issuer = c.oidc.issuer
		if c.oidc.jwks != nil {
			description.JWKSURL = c.oidc.jwks.url
		}
		description.RequireTenantScope = c.oidc.requireTenantScope
		description.AllowGlobalSecurityAdmin = c.oidc.allowGlobalSecurityAdmin
		description.RoleBindings = describeRoleBindings(c.oidc.roleBindings)
	}
	return description
}

func describeRoleBindings(values map[string]map[string]struct{}) map[string][]string {
	if len(values) == 0 {
		return nil
	}
	result := make(map[string][]string, len(values))
	for role, bindings := range values {
		items := make([]string, 0, len(bindings))
		for value := range bindings {
			items = append(items, value)
		}
		sort.Strings(items)
		result[role] = items
	}
	return result
}
