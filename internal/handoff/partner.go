package handoff

import (
	"strings"
	"time"
)

const (
	PartnerIntakeSchemaVersion    = "4.enterprise_partner_intake.v1"
	PartnerDashboardSchemaVersion = "4.enterprise_partner_dashboard.v1"

	IntakeStateReceived    = "received"
	IntakeStateVerified    = "verified"
	IntakeStateAccepted    = "accepted"
	IntakeStateExpired     = "expired"
	IntakeStateSuperseded  = "superseded"
	IntakeStateRejected    = "rejected"
	IntakeStateUnderReview = "under_review"
)

type IntakeInput struct {
	PartnerID                string   `json:"partner_id"`
	Organization             string   `json:"organization,omitempty"`
	TrustDomain              string   `json:"trust_domain,omitempty"`
	HandoffRef               string   `json:"handoff_ref,omitempty"`
	VerificationStatus       string   `json:"verification_status,omitempty"`
	FreshnessState           string   `json:"freshness_state,omitempty"`
	PolicyCompatibility      string   `json:"policy_compatibility,omitempty"`
	ExceptionState           string   `json:"exception_state,omitempty"`
	IncidentDisclosureStatus string   `json:"incident_disclosure_status,omitempty"`
	PartnerVisibleEvidence   []string `json:"partner_visible_evidence,omitempty"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	ReasonCodes              []string `json:"reason_codes,omitempty"`
}

type DashboardProjection struct {
	SchemaVersion            string   `json:"schema_version"`
	PartnerID                string   `json:"partner_id"`
	CurrentState             string   `json:"current_state"`
	TrustSummary             string   `json:"trust_summary"`
	VerificationFreshness    string   `json:"verification_freshness,omitempty"`
	SharedResponsibilityView []string `json:"shared_responsibility_view,omitempty"`
	PartnerVisibleEvidence   []string `json:"partner_visible_evidence,omitempty"`
	SensitiveSignalsRedacted bool     `json:"sensitive_signals_redacted"`
	Limitations              []string `json:"limitations,omitempty"`
}

type IntakeRecord struct {
	SchemaVersion           string              `json:"schema_version"`
	PartnerID               string              `json:"partner_id"`
	Organization            string              `json:"organization,omitempty"`
	TrustDomain             string              `json:"trust_domain,omitempty"`
	HandoffRef              string              `json:"handoff_ref,omitempty"`
	CurrentState            string              `json:"current_state"`
	VerificationStatus      string              `json:"verification_status,omitempty"`
	FreshnessState          string              `json:"freshness_state,omitempty"`
	PolicyCompatibility     string              `json:"policy_compatibility,omitempty"`
	ExceptionState          string              `json:"exception_state,omitempty"`
	IncidentDisclosureState string              `json:"incident_disclosure_state,omitempty"`
	TrustScore              int                 `json:"trust_score"`
	ReasonCodes             []string            `json:"reason_codes,omitempty"`
	EvidenceRefs            []string            `json:"evidence_refs,omitempty"`
	Dashboard               DashboardProjection `json:"dashboard"`
	ObservedAt              time.Time           `json:"observed_at"`
	Limitations             []string            `json:"limitations,omitempty"`
}

func EvaluateIntake(input IntakeInput, now func() time.Time) IntakeRecord {
	if now == nil {
		now = time.Now
	}
	input.PartnerID = strings.TrimSpace(input.PartnerID)
	input.Organization = strings.TrimSpace(input.Organization)
	input.TrustDomain = strings.TrimSpace(input.TrustDomain)
	input.HandoffRef = strings.TrimSpace(input.HandoffRef)
	input.VerificationStatus = strings.ToLower(strings.TrimSpace(input.VerificationStatus))
	input.FreshnessState = strings.ToLower(strings.TrimSpace(input.FreshnessState))
	input.PolicyCompatibility = strings.ToLower(strings.TrimSpace(input.PolicyCompatibility))
	input.ExceptionState = strings.ToLower(strings.TrimSpace(input.ExceptionState))
	input.IncidentDisclosureStatus = strings.ToLower(strings.TrimSpace(input.IncidentDisclosureStatus))

	score := 82
	state := IntakeStateReceived
	reasons := append([]string{}, input.ReasonCodes...)

	switch {
	case input.VerificationStatus == "rejected" || input.VerificationStatus == "failed":
		state = IntakeStateRejected
		score -= 55
		reasons = append(reasons, "verifier_rejected_partner_handoff")
	case input.FreshnessState == "expired":
		state = IntakeStateExpired
		score -= 30
		reasons = append(reasons, "proof_freshness_expired")
	case input.PolicyCompatibility == "mismatch":
		state = IntakeStateUnderReview
		score -= 20
		reasons = append(reasons, "policy_compatibility_mismatch")
	case input.VerificationStatus == "verified" && input.FreshnessState != "expired" && input.PolicyCompatibility != "mismatch":
		state = IntakeStateAccepted
		reasons = append(reasons, "partner_handoff_verified")
	default:
		state = IntakeStateUnderReview
	}
	if input.ExceptionState == "active" {
		score -= 12
		reasons = append(reasons, "partner_exception_active")
	}
	if input.IncidentDisclosureStatus == "withheld" {
		score -= 8
		reasons = append(reasons, "incident_disclosure_restricted")
	}
	if input.FreshnessState == "" {
		input.FreshnessState = "fresh"
	}
	if input.PolicyCompatibility == "" {
		input.PolicyCompatibility = "compatible"
	}
	if input.VerificationStatus == "" {
		input.VerificationStatus = "under_review"
	}

	dashboard := DashboardProjection{
		SchemaVersion:         PartnerDashboardSchemaVersion,
		PartnerID:             input.PartnerID,
		CurrentState:          state,
		TrustSummary:          dashboardSummary(state, score),
		VerificationFreshness: input.FreshnessState,
		SharedResponsibilityView: []string{
			"partner view shows bounded trust state, freshness, and disclosure posture without exposing internal sensitive signals",
			"local verification continuity remains authoritative for acceptance, expiry, and rejection transitions",
		},
		PartnerVisibleEvidence:   uniqueStrings(input.PartnerVisibleEvidence),
		SensitiveSignalsRedacted: true,
		Limitations: []string{
			"Partner-safe dashboards are bounded projections and intentionally exclude internal sensitive runtime and investigation signals.",
		},
	}

	return IntakeRecord{
		SchemaVersion:           PartnerIntakeSchemaVersion,
		PartnerID:               input.PartnerID,
		Organization:            input.Organization,
		TrustDomain:             input.TrustDomain,
		HandoffRef:              input.HandoffRef,
		CurrentState:            state,
		VerificationStatus:      input.VerificationStatus,
		FreshnessState:          input.FreshnessState,
		PolicyCompatibility:     input.PolicyCompatibility,
		ExceptionState:          input.ExceptionState,
		IncidentDisclosureState: input.IncidentDisclosureStatus,
		TrustScore:              clamp(score, 0, 100),
		ReasonCodes:             uniqueStrings(reasons),
		EvidenceRefs:            uniqueStrings(input.EvidenceRefs),
		Dashboard:               dashboard,
		ObservedAt:              now().UTC(),
		Limitations: []string{
			"Partner trust scoring remains bounded and explainable; it is not a blind global truth score.",
			"External trust intake remains tied to local verifier evidence and local compatibility review.",
		},
	}
}

func dashboardSummary(state string, score int) string {
	switch state {
	case IntakeStateAccepted:
		return "partner handoff accepted with bounded verification continuity"
	case IntakeStateRejected:
		return "partner handoff rejected by local verification or compatibility policy"
	case IntakeStateExpired:
		return "partner trust degraded because freshness guarantees expired"
	default:
		if score >= 60 {
			return "partner handoff remains under review with bounded trust posture"
		}
		return "partner trust remains restricted pending verification and compatibility review"
	}
}

func clamp(value, minValue, maxValue int) int {
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		items = append(items, value)
	}
	if len(items) == 0 {
		return nil
	}
	return items
}
