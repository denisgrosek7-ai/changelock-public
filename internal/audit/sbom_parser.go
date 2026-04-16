package audit

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ParseSBOMComponents(format string, raw json.RawMessage, imageDigest string) ([]SBOMComponent, error) {
	switch normalizeSBOMFormat(format) {
	case SBOMFormatSPDXJSON:
		return parseSPDXComponents(raw, imageDigest)
	case SBOMFormatCycloneDXJSON:
		return parseCycloneDXComponents(raw, imageDigest)
	default:
		return nil, fmt.Errorf("%w: unsupported sbom_format %q", ErrInvalidException, format)
	}
}

type spdxDocument struct {
	Packages []struct {
		Name                  string `json:"name"`
		VersionInfo           string `json:"versionInfo"`
		LicenseConcluded      string `json:"licenseConcluded"`
		LicenseDeclared       string `json:"licenseDeclared"`
		PrimaryPackagePurpose string `json:"primaryPackagePurpose"`
		SPDXID                string `json:"SPDXID"`
		ExternalRefs          []struct {
			ReferenceType    string `json:"referenceType"`
			ReferenceLocator string `json:"referenceLocator"`
		} `json:"externalRefs"`
	} `json:"packages"`
}

func parseSPDXComponents(raw json.RawMessage, imageDigest string) ([]SBOMComponent, error) {
	var document spdxDocument
	if err := json.Unmarshal(raw, &document); err != nil {
		return nil, fmt.Errorf("%w: invalid SPDX JSON: %v", ErrInvalidException, err)
	}

	seen := map[string]struct{}{}
	components := make([]SBOMComponent, 0, len(document.Packages))
	for _, pkg := range document.Packages {
		name := strings.TrimSpace(pkg.Name)
		if name == "" {
			continue
		}
		purl := ""
		for _, ref := range pkg.ExternalRefs {
			if strings.EqualFold(strings.TrimSpace(ref.ReferenceType), "purl") {
				purl = strings.TrimSpace(ref.ReferenceLocator)
				break
			}
		}
		component := SBOMComponent{
			ImageDigest:      imageDigest,
			ComponentName:    name,
			ComponentVersion: strings.TrimSpace(pkg.VersionInfo),
			ComponentType:    strings.TrimSpace(pkg.PrimaryPackagePurpose),
			License:          firstNonEmptyValue(strings.TrimSpace(pkg.LicenseConcluded), strings.TrimSpace(pkg.LicenseDeclared)),
			PURL:             purl,
		}
		key := strings.Join([]string{component.ComponentName, component.ComponentVersion, component.PURL}, "|")
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		metadata, _ := json.Marshal(map[string]string{
			"spdx_id": pkg.SPDXID,
		})
		component.Metadata = normalizeMetadata(metadata)
		components = append(components, component)
	}
	return components, nil
}

type cycloneDXDocument struct {
	Metadata struct {
		Component *cycloneDXComponent `json:"component"`
	} `json:"metadata"`
	Components []cycloneDXComponent `json:"components"`
}

type cycloneDXComponent struct {
	Type       string               `json:"type"`
	Name       string               `json:"name"`
	Version    string               `json:"version"`
	PURL       string               `json:"purl"`
	Licenses   []cycloneDXLicense   `json:"licenses"`
	Components []cycloneDXComponent `json:"components"`
}

type cycloneDXLicense struct {
	License *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"license"`
}

func parseCycloneDXComponents(raw json.RawMessage, imageDigest string) ([]SBOMComponent, error) {
	var document cycloneDXDocument
	if err := json.Unmarshal(raw, &document); err != nil {
		return nil, fmt.Errorf("%w: invalid CycloneDX JSON: %v", ErrInvalidException, err)
	}

	seen := map[string]struct{}{}
	components := []SBOMComponent{}
	var walk func(component cycloneDXComponent)
	walk = func(component cycloneDXComponent) {
		name := strings.TrimSpace(component.Name)
		if name != "" {
			item := SBOMComponent{
				ImageDigest:      imageDigest,
				ComponentName:    name,
				ComponentVersion: strings.TrimSpace(component.Version),
				ComponentType:    strings.TrimSpace(component.Type),
				License:          cycloneDXLicenseValue(component.Licenses),
				PURL:             strings.TrimSpace(component.PURL),
			}
			key := strings.Join([]string{item.ComponentName, item.ComponentVersion, item.PURL}, "|")
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				components = append(components, item)
			}
		}
		for _, child := range component.Components {
			walk(child)
		}
	}

	if document.Metadata.Component != nil {
		walk(*document.Metadata.Component)
	}
	for _, component := range document.Components {
		walk(component)
	}
	return components, nil
}

func cycloneDXLicenseValue(licenses []cycloneDXLicense) string {
	for _, entry := range licenses {
		if entry.License == nil {
			continue
		}
		if value := strings.TrimSpace(entry.License.ID); value != "" {
			return value
		}
		if value := strings.TrimSpace(entry.License.Name); value != "" {
			return value
		}
	}
	return ""
}

func firstNonEmptyValue(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
