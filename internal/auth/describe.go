package auth

import "sort"

type RoleBindingDescription struct {
	ChangelockRole string   `json:"changelock_role"`
	BindingValues  []string `json:"binding_values,omitempty"`
	Subjects       []string `json:"subjects,omitempty"`
}

type Description struct {
	Mode                       string                   `json:"mode"`
	IdentityProvider           string                   `json:"identity_provider"`
	RoleClaim                  string                   `json:"role_claim,omitempty"`
	TenantClaim                string                   `json:"tenant_claim,omitempty"`
	SubjectClaim               string                   `json:"subject_claim,omitempty"`
	EmailClaim                 string                   `json:"email_claim,omitempty"`
	Issuer                     string                   `json:"issuer,omitempty"`
	JWKSURL                    string                   `json:"jwks_url,omitempty"`
	RequireTenantScope         bool                     `json:"require_tenant_scope,omitempty"`
	AllowGlobalSecurityAdmin   bool                     `json:"allow_global_security_admin,omitempty"`
	TenantScopeRequired        bool                     `json:"tenant_scope_required"`
	GlobalSecurityAdminAllowed bool                     `json:"global_security_admin_allowed"`
	RoleBindings               map[string][]string      `json:"role_bindings,omitempty"`
	RoleBindingDetails         []RoleBindingDescription `json:"role_binding_details,omitempty"`
	StaticTokenRoleCounts      map[string]int           `json:"static_token_role_counts,omitempty"`
	AuditActorAttribution      bool                     `json:"audit_actor_attribution"`
	ApprovalActorAttribution   bool                     `json:"approval_actor_attribution"`
	Limitations                []string                 `json:"limitations,omitempty"`
}

func (c Config) Describe() Description {
	description := Description{
		Mode:                     c.Mode,
		IdentityProvider:         c.Mode,
		AuditActorAttribution:    true,
		ApprovalActorAttribution: true,
	}

	switch c.Mode {
	case ModeDisabled:
		description.IdentityProvider = "disabled"
		description.Limitations = []string{
			"Disabled auth mode is local-only and does not represent an enterprise identity fabric integration.",
		}
	case ModeStaticToken:
		description.IdentityProvider = "static_token"
		counts := map[string]int{}
		for _, token := range c.tokens {
			counts[token.Role]++
		}
		description.StaticTokenRoleCounts = counts
		description.RoleBindings = describeStaticTokenRoleMap(c.tokens)
		description.RoleBindingDetails = describeStaticTokenBindings(c.tokens)
		description.Limitations = []string{
			"Static token mode is bounded to locally configured bearer subjects and does not prove upstream identity-provider health or delegation state.",
		}
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
		description.TenantScopeRequired = c.oidc.requireTenantScope
		description.GlobalSecurityAdminAllowed = c.oidc.allowGlobalSecurityAdmin
		description.RoleBindings = describeRoleBindings(c.oidc.roleBindings)
		description.RoleBindingDetails = describeOIDCRoleBindings(c.oidc.roleBindings)
		description.Limitations = []string{
			"OIDC description summarizes configured claims and role bindings, but it does not prove current issuer or JWKS availability.",
		}
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

func describeStaticTokenRoleMap(tokens []TokenConfig) map[string][]string {
	if len(tokens) == 0 {
		return nil
	}
	result := map[string][]string{}
	for _, token := range tokens {
		if token.Role == "" || token.Subject == "" {
			continue
		}
		result[token.Role] = append(result[token.Role], token.Subject)
	}
	for role, subjects := range result {
		result[role] = uniqueSortedDescriptions(subjects)
	}
	return result
}

func describeStaticTokenBindings(tokens []TokenConfig) []RoleBindingDescription {
	roleMap := describeStaticTokenRoleMap(tokens)
	if len(roleMap) == 0 {
		return nil
	}
	roles := make([]string, 0, len(roleMap))
	for role := range roleMap {
		roles = append(roles, role)
	}
	sort.Strings(roles)

	descriptions := make([]RoleBindingDescription, 0, len(roles))
	for _, role := range roles {
		descriptions = append(descriptions, RoleBindingDescription{
			ChangelockRole: role,
			Subjects:       roleMap[role],
		})
	}
	return descriptions
}

func describeOIDCRoleBindings(bindings map[string]map[string]struct{}) []RoleBindingDescription {
	if len(bindings) == 0 {
		return nil
	}
	roles := make([]string, 0, len(bindings))
	for role := range bindings {
		roles = append(roles, role)
	}
	sort.Strings(roles)

	descriptions := make([]RoleBindingDescription, 0, len(roles))
	for _, role := range roles {
		values := make([]string, 0, len(bindings[role]))
		for value := range bindings[role] {
			values = append(values, value)
		}
		descriptions = append(descriptions, RoleBindingDescription{
			ChangelockRole: role,
			BindingValues:  uniqueSortedDescriptions(values),
		})
	}
	return descriptions
}

func uniqueSortedDescriptions(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	unique := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		unique = append(unique, value)
	}
	sort.Strings(unique)
	return unique
}
