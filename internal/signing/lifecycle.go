package signing

const (
	KeyClassSealing               = "sealing"
	KeyClassExceptionApproval     = "exception_approval"
	KeyClassSyncSnapshot          = "sync_snapshot"
	KeyClassFederationHandoff     = "federation_handoff"
	KeyClassValidationCertificate = "validation_certificate"
	KeyClassVerificationRoot      = "verification_root"
	KeyClassDevelopmentEphemeral  = "development_ephemeral"

	KeyStateProvisioned       = "provisioned"
	KeyStateActive            = "active"
	KeyStateRotatePending     = "rotate_pending"
	KeyStateRetiredVerifyOnly = "retired_verify_only"
	KeyStateRevoked           = "revoked"
	KeyStateDestroyed         = "destroyed"

	TrustBoundaryDisabled          = "disabled"
	TrustBoundaryApplicationLocal  = "application_local"
	TrustBoundaryExternalManaged   = "external_managed"
	TrustBoundaryDevelopmentScoped = "development_scoped"
)

type ProviderDescriptor struct {
	ProviderMode                   string   `json:"provider_mode"`
	TrustBoundary                  string   `json:"trust_boundary"`
	ActiveLifecycleState           string   `json:"active_lifecycle_state"`
	KeyExportability               string   `json:"key_exportability"`
	AuditLogAvailability           string   `json:"audit_log_availability"`
	SupportsHistoricalVerification bool     `json:"supports_historical_verification"`
	SupportsRotation               bool     `json:"supports_rotation"`
	SupportsVerifyOnlyRetirement   bool     `json:"supports_verify_only_retirement"`
	SupportsRevocation             bool     `json:"supports_revocation"`
	KeyClasses                     []string `json:"key_classes,omitempty"`
	CapabilityMatrix               []string `json:"capability_matrix,omitempty"`
	Limitations                    []string `json:"limitations,omitempty"`
}

func (r *Runtime) DescribeProvider() ProviderDescriptor {
	if r == nil {
		return ProviderDescriptor{
			ProviderMode:         ModeDisabled,
			TrustBoundary:        TrustBoundaryDisabled,
			ActiveLifecycleState: KeyStateProvisioned,
			KeyExportability:     "not_applicable",
			AuditLogAvailability: "not_applicable",
			KeyClasses:           defaultKeyClasses(),
			Limitations:          []string{"No signer runtime is configured, so production trust material is not active."},
		}
	}

	switch r.Config.Mode {
	case ModeSoftware:
		return ProviderDescriptor{
			ProviderMode:                   ModeSoftware,
			TrustBoundary:                  TrustBoundaryApplicationLocal,
			ActiveLifecycleState:           KeyStateActive,
			KeyExportability:               "exportable_application_secret",
			AuditLogAvailability:           "application_audit_only",
			SupportsHistoricalVerification: true,
			SupportsRotation:               true,
			SupportsVerifyOnlyRetirement:   true,
			SupportsRevocation:             true,
			KeyClasses:                     append(defaultKeyClasses(), KeyClassDevelopmentEphemeral),
			CapabilityMatrix: []string{
				"sign",
				"verify",
				"purpose_scoped_signing",
				"verify_only_retirement_supported_by_trust_set_history",
			},
			Limitations: []string{
				"Software signing is suitable for development, testing, or explicitly accepted lower-trust deployments. It is not a substitute for externally managed production trust custody.",
			},
		}
	case ModeVaultTransit:
		return ProviderDescriptor{
			ProviderMode:                   ModeVaultTransit,
			TrustBoundary:                  TrustBoundaryExternalManaged,
			ActiveLifecycleState:           KeyStateActive,
			KeyExportability:               "non_exportable_remote_key_material",
			AuditLogAvailability:           "provider_and_application_audit",
			SupportsHistoricalVerification: true,
			SupportsRotation:               true,
			SupportsVerifyOnlyRetirement:   true,
			SupportsRevocation:             true,
			KeyClasses:                     defaultKeyClasses(),
			CapabilityMatrix: []string{
				"sign",
				"verify",
				"provider_managed_key_custody",
				"provider_backed_rotation_window",
				"verify_only_retirement_supported_by_trust_set_history",
			},
			Limitations: []string{
				"Vault transit remains the only externalized production signer provider currently implemented in code. Wider KMS/HSM provider coverage is still future work.",
			},
		}
	default:
		return ProviderDescriptor{
			ProviderMode:         ModeDisabled,
			TrustBoundary:        TrustBoundaryDisabled,
			ActiveLifecycleState: KeyStateProvisioned,
			KeyExportability:     "not_applicable",
			AuditLogAvailability: "not_applicable",
			KeyClasses:           defaultKeyClasses(),
			Limitations:          []string{"Signer runtime is disabled or unsupported for this deployment profile."},
		}
	}
}

func defaultKeyClasses() []string {
	return []string{
		KeyClassSealing,
		KeyClassExceptionApproval,
		KeyClassSyncSnapshot,
		KeyClassFederationHandoff,
		KeyClassValidationCertificate,
		KeyClassVerificationRoot,
	}
}
