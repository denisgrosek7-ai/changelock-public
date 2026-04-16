package vex

import (
	"encoding/json"
	"fmt"
	"strings"
)

type csafNote struct {
	Category string `json:"category"`
	Text     string `json:"text"`
}

type csafRemediation struct {
	Details string `json:"details"`
}

type cyclonedxProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func DetectFormat(payload json.RawMessage) (string, error) {
	var document map[string]json.RawMessage
	if err := json.Unmarshal(payload, &document); err != nil {
		return "", fmt.Errorf("invalid vex payload: %w", err)
	}
	if raw, ok := document["bomFormat"]; ok {
		var bomFormat string
		if err := json.Unmarshal(raw, &bomFormat); err == nil && strings.EqualFold(strings.TrimSpace(bomFormat), "CycloneDX") {
			return SourceFormatCycloneDX, nil
		}
	}
	if _, ok := document["document"]; ok {
		if _, ok := document["vulnerabilities"]; ok {
			return SourceFormatCSAF, nil
		}
	}
	return "", fmt.Errorf("unsupported vex payload format")
}

func ParseIngestRequest(request IngestRequest) ([]CreateRequest, string, error) {
	request.Format = normalizeSourceFormat(request.Format)
	request.SourceRef = strings.TrimSpace(request.SourceRef)
	request.Scope = normalizeScope(request.Scope)
	request.Payload = normalizeJSON(request.Payload)
	if len(request.Payload) == 0 {
		return nil, "", fmt.Errorf("%w: payload is required", ErrInvalidStatement)
	}
	if request.Format == "" {
		format, err := DetectFormat(request.Payload)
		if err != nil {
			return nil, "", err
		}
		request.Format = format
	}

	var (
		statements []CreateRequest
		err        error
	)
	switch request.Format {
	case SourceFormatCSAF:
		statements, err = parseCSAF(request)
	case SourceFormatCycloneDX:
		statements, err = parseCycloneDX(request)
	default:
		err = fmt.Errorf("%w: unsupported ingest format %q", ErrInvalidStatement, request.Format)
	}
	if err != nil {
		return nil, "", err
	}
	for i := range statements {
		statements[i].SourceFormat = request.Format
		statements[i].SourceRef = firstNonEmpty(statements[i].SourceRef, request.SourceRef)
		if statements[i].ExpiresAt == nil && request.ExpiresAt != nil {
			expiresAt := request.ExpiresAt.UTC()
			statements[i].ExpiresAt = &expiresAt
		}
		if statements[i].Scope.ImageDigest == "" {
			statements[i].Scope.ImageDigest = request.Scope.ImageDigest
		}
		if statements[i].Scope.PackageName == "" {
			statements[i].Scope.PackageName = request.Scope.PackageName
		}
		if statements[i].Scope.PURL == "" {
			statements[i].Scope.PURL = request.Scope.PURL
		}
		if statements[i].Scope.Repo == "" {
			statements[i].Scope.Repo = request.Scope.Repo
		}
		if statements[i].Scope.Workload == "" {
			statements[i].Scope.Workload = request.Scope.Workload
		}
		if statements[i].Scope.TenantID == "" {
			statements[i].Scope.TenantID = request.Scope.TenantID
		}
		if statements[i].Scope.ClusterID == "" {
			statements[i].Scope.ClusterID = request.Scope.ClusterID
		}
		if statements[i].Scope.Environment == "" {
			statements[i].Scope.Environment = request.Scope.Environment
		}
		if statements[i].Scope.Namespace == "" {
			statements[i].Scope.Namespace = request.Scope.Namespace
		}
		normalized, err := NormalizeCreateRequest(statements[i], nil)
		if err != nil {
			return nil, "", err
		}
		statements[i] = normalized
	}
	return statements, request.Format, nil
}

func parseCSAF(request IngestRequest) ([]CreateRequest, error) {
	type product struct {
		ProductID string `json:"product_id"`
		Name      string `json:"name"`
		Helper    struct {
			PURL string `json:"purl"`
		} `json:"product_identification_helper"`
	}
	type vulnerability struct {
		CVE           string `json:"cve"`
		IDs           []struct {
			Text string `json:"text"`
		} `json:"ids"`
		Notes         []csafNote        `json:"notes"`
		Remediations  []csafRemediation `json:"remediations"`
		ProductStatus struct {
			KnownNotAffected   []string `json:"known_not_affected"`
			KnownAffected      []string `json:"known_affected"`
			Fixed              []string `json:"fixed"`
			UnderInvestigation []string `json:"under_investigation"`
		} `json:"product_status"`
	}
	var document struct {
		ProductTree struct {
			FullProductNames []product `json:"full_product_names"`
		} `json:"product_tree"`
		Vulnerabilities []vulnerability `json:"vulnerabilities"`
	}
	if err := json.Unmarshal(request.Payload, &document); err != nil {
		return nil, fmt.Errorf("invalid csaf vex payload: %w", err)
	}
	products := map[string]Scope{}
	for _, entry := range document.ProductTree.FullProductNames {
		scope := Scope{
			ImageDigest: extractDigest(entry.Name),
			PURL:        strings.TrimSpace(entry.Helper.PURL),
			PackageName: strings.TrimSpace(entry.Name),
		}
		products[strings.TrimSpace(entry.ProductID)] = normalizeScope(scope)
	}

	results := []CreateRequest{}
	for _, item := range document.Vulnerabilities {
		vulnerabilityID := normalizeVulnerabilityID(firstNonEmpty(item.CVE, firstText(item.IDs)))
		if vulnerabilityID == "" {
			return nil, fmt.Errorf("%w: csaf vulnerability entry is missing cve/id", ErrInvalidStatement)
		}
		impact := joinNotes(item.Notes)
		action := firstRemediation(item.Remediations)
		appendStatements := func(status string, productIDs []string) error {
			for _, productID := range productIDs {
				scope := products[strings.TrimSpace(productID)]
				if !scopeHasMatchTarget(scope) {
					scope = request.Scope
				}
				results = append(results, CreateRequest{
					SourceFormat:    SourceFormatCSAF,
					SourceRef:       request.SourceRef,
					VulnerabilityID: vulnerabilityID,
					Scope:           scope,
					Status:          status,
					Justification:   noteByCategory(item.Notes, "details", "description", "summary"),
					ActionStatement: action,
					ImpactStatement: impact,
					ExpiresAt:       request.ExpiresAt,
				})
			}
			return nil
		}
		if err := appendStatements(StatusNotAffected, item.ProductStatus.KnownNotAffected); err != nil {
			return nil, err
		}
		if err := appendStatements(StatusAffected, item.ProductStatus.KnownAffected); err != nil {
			return nil, err
		}
		if err := appendStatements(StatusFixed, item.ProductStatus.Fixed); err != nil {
			return nil, err
		}
		if err := appendStatements(StatusUnderInvestigation, item.ProductStatus.UnderInvestigation); err != nil {
			return nil, err
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("%w: csaf payload did not produce any supported vex statements", ErrInvalidStatement)
	}
	return results, nil
}

func parseCycloneDX(request IngestRequest) ([]CreateRequest, error) {
	type component struct {
		BOMRef     string     `json:"bom-ref"`
		Name       string     `json:"name"`
		PURL       string     `json:"purl"`
		Version    string     `json:"version"`
		Properties []cyclonedxProperty `json:"properties"`
	}
	type analysis struct {
		State         string   `json:"state"`
		Justification string   `json:"justification"`
		Detail        string   `json:"detail"`
		Response      []string `json:"response"`
	}
	type vulnerability struct {
		ID             string `json:"id"`
		Affects        []struct {
			Ref string `json:"ref"`
		} `json:"affects"`
		Analysis       analysis `json:"analysis"`
		Recommendation string   `json:"recommendation"`
	}
	var document struct {
		BOMFormat      string          `json:"bomFormat"`
		Metadata       struct {
			Component component `json:"component"`
		} `json:"metadata"`
		Components     []component     `json:"components"`
		Vulnerabilities []vulnerability `json:"vulnerabilities"`
	}
	if err := json.Unmarshal(request.Payload, &document); err != nil {
		return nil, fmt.Errorf("invalid cyclonedx vex payload: %w", err)
	}
	if !strings.EqualFold(strings.TrimSpace(document.BOMFormat), "CycloneDX") {
		return nil, fmt.Errorf("%w: cyclonedx bomFormat is required", ErrInvalidStatement)
	}
	components := map[string]Scope{}
	registerComponent := func(entry component) {
		imageDigest := extractDigest(firstNonEmpty(entry.BOMRef, propertyValue(entry.Properties, "changelock:image-digest"), propertyValue(entry.Properties, "image-digest")))
		components[strings.TrimSpace(entry.BOMRef)] = normalizeScope(Scope{
			ImageDigest: imageDigest,
			PackageName: strings.TrimSpace(entry.Name),
			PURL:        strings.TrimSpace(entry.PURL),
		})
	}
	if strings.TrimSpace(document.Metadata.Component.BOMRef) != "" || strings.TrimSpace(document.Metadata.Component.Name) != "" || strings.TrimSpace(document.Metadata.Component.PURL) != "" {
		registerComponent(document.Metadata.Component)
	}
	for _, entry := range document.Components {
		registerComponent(entry)
	}

	results := []CreateRequest{}
	for _, item := range document.Vulnerabilities {
		vulnerabilityID := normalizeVulnerabilityID(item.ID)
		if vulnerabilityID == "" {
			return nil, fmt.Errorf("%w: cyclonedx vulnerability id is required", ErrInvalidStatement)
		}
		status, err := mapCycloneDXState(item.Analysis.State)
		if err != nil {
			return nil, err
		}
		refs := item.Affects
		if len(refs) == 0 {
			scope := request.Scope
			if !scopeHasMatchTarget(scope) {
				scope = components[strings.TrimSpace(document.Metadata.Component.BOMRef)]
			}
			results = append(results, CreateRequest{
				SourceFormat:    SourceFormatCycloneDX,
				SourceRef:       request.SourceRef,
				VulnerabilityID: vulnerabilityID,
				Scope:           scope,
				Status:          status,
				Justification:   strings.TrimSpace(item.Analysis.Justification),
				ActionStatement: strings.TrimSpace(firstNonEmpty(item.Recommendation, strings.Join(item.Analysis.Response, ", "))),
				ImpactStatement: strings.TrimSpace(item.Analysis.Detail),
				ExpiresAt:       request.ExpiresAt,
			})
			continue
		}
		for _, affected := range refs {
			scope := components[strings.TrimSpace(affected.Ref)]
			if !scopeHasMatchTarget(scope) {
				scope = request.Scope
			}
			results = append(results, CreateRequest{
				SourceFormat:    SourceFormatCycloneDX,
				SourceRef:       request.SourceRef,
				VulnerabilityID: vulnerabilityID,
				Scope:           scope,
				Status:          status,
				Justification:   strings.TrimSpace(item.Analysis.Justification),
				ActionStatement: strings.TrimSpace(firstNonEmpty(item.Recommendation, strings.Join(item.Analysis.Response, ", "))),
				ImpactStatement: strings.TrimSpace(item.Analysis.Detail),
				ExpiresAt:       request.ExpiresAt,
			})
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("%w: cyclonedx payload did not produce any supported vex statements", ErrInvalidStatement)
	}
	return results, nil
}

func mapCycloneDXState(state string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(state)) {
	case "not_affected", "false_positive":
		return StatusNotAffected, nil
	case "exploitable", "affected":
		return StatusAffected, nil
	case "resolved", "fixed":
		return StatusFixed, nil
	case "under_investigation", "in_triage":
		return StatusUnderInvestigation, nil
	default:
		return "", fmt.Errorf("%w: unsupported cyclonedx analysis.state %q", ErrInvalidStatement, state)
	}
}

func extractDigest(values ...string) string {
	for _, value := range values {
		text := strings.TrimSpace(value)
		if text == "" {
			continue
		}
		if idx := strings.Index(text, "sha256:"); idx >= 0 {
			return text[idx:]
		}
	}
	return ""
}

func firstText(values []struct{ Text string `json:"text"` }) string {
	for _, value := range values {
		if strings.TrimSpace(value.Text) != "" {
			return strings.TrimSpace(value.Text)
		}
	}
	return ""
}

func joinNotes(notes []csafNote) string {
	parts := make([]string, 0, len(notes))
	for _, note := range notes {
		if strings.TrimSpace(note.Text) != "" {
			parts = append(parts, strings.TrimSpace(note.Text))
		}
	}
	return strings.Join(parts, " ")
}

func noteByCategory(notes []csafNote, categories ...string) string {
	for _, category := range categories {
		for _, note := range notes {
			if strings.EqualFold(strings.TrimSpace(note.Category), category) && strings.TrimSpace(note.Text) != "" {
				return strings.TrimSpace(note.Text)
			}
		}
	}
	return ""
}

func firstRemediation(remediations []csafRemediation) string {
	for _, remediation := range remediations {
		if strings.TrimSpace(remediation.Details) != "" {
			return strings.TrimSpace(remediation.Details)
		}
	}
	return ""
}

func propertyValue(properties []cyclonedxProperty, names ...string) string {
	for _, property := range properties {
		for _, name := range names {
			if strings.EqualFold(strings.TrimSpace(property.Name), name) && strings.TrimSpace(property.Value) != "" {
				return strings.TrimSpace(property.Value)
			}
		}
	}
	return ""
}
