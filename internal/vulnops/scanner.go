package vulnops

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
)

const (
	DefaultScanInterval = 6 * time.Hour
	ScannerTrivy        = "trivy"
	ScannerGrype        = "grype"
)

type Config struct {
	Enabled           bool
	SBOMIngestEnabled bool
	ScanInterval      time.Duration
	Scanner           string
	TrivyPath         string
	GrypePath         string
}

type Result struct {
	ImageDigest string
	ImageRef    string
	Scanner     string
	StartedAt   time.Time
	CompletedAt time.Time
	Status      string
	Summary     json.RawMessage
	SourceRef   string
	Findings    []audit.VulnerabilityFindingInput
}

type Scanner interface {
	ScanDigest(ctx context.Context, target audit.ActiveDigestRef) (Result, error)
}

func ConfigFromEnv() (Config, error) {
	config := Config{
		Enabled:           parseBoolEnv("CHANGELOCK_VULNOPS_ENABLED", false),
		SBOMIngestEnabled: parseBoolEnv("CHANGELOCK_SBOM_INGEST_ENABLED", true),
		Scanner:           normalizeScannerName(firstNonEmpty(os.Getenv("CHANGELOCK_VULNOPS_SCANNER"), ScannerTrivy)),
		TrivyPath:         strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_VULNOPS_TRIVY_PATH"), ScannerTrivy)),
		GrypePath:         strings.TrimSpace(firstNonEmpty(os.Getenv("CHANGELOCK_VULNOPS_GRYPE_PATH"), ScannerGrype)),
		ScanInterval:      parseDurationEnv("CHANGELOCK_VULNOPS_SCAN_INTERVAL", DefaultScanInterval),
	}
	if !config.Enabled {
		return config, nil
	}
	switch config.Scanner {
	case ScannerTrivy, ScannerGrype:
	default:
		return Config{}, fmt.Errorf("unsupported CHANGELOCK_VULNOPS_SCANNER: %s", config.Scanner)
	}
	if config.ScanInterval <= 0 {
		return Config{}, errors.New("CHANGELOCK_VULNOPS_SCAN_INTERVAL must be positive when vulnerability ops are enabled")
	}
	return config, nil
}

func NewScanner(config Config) (Scanner, error) {
	if !config.Enabled {
		return nil, nil
	}
	switch config.Scanner {
	case ScannerTrivy:
		return trivyScanner{path: config.TrivyPath}, nil
	case ScannerGrype:
		return grypeScanner{path: config.GrypePath}, nil
	default:
		return nil, fmt.Errorf("unsupported vulnerability scanner: %s", config.Scanner)
	}
}

type trivyScanner struct {
	path string
}

func (s trivyScanner) ScanDigest(ctx context.Context, target audit.ActiveDigestRef) (Result, error) {
	targetRef, err := targetReference(target)
	if err != nil {
		return Result{}, err
	}

	startedAt := time.Now().UTC()
	cmd := exec.CommandContext(ctx, s.path, "image", "--quiet", "--scanners", "vuln", "--format", "json", targetRef)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return Result{}, formatScanError("trivy", targetRef, err, stderr.String())
	}

	report, err := parseTrivyReport(stdout.Bytes())
	if err != nil {
		return Result{}, err
	}
	return Result{
		ImageDigest: target.ImageDigest,
		ImageRef:    target.ImageRef,
		Scanner:     ScannerTrivy,
		StartedAt:   startedAt,
		CompletedAt: time.Now().UTC(),
		Status:      audit.VulnerabilityScanStatusCompleted,
		Summary:     report.summaryJSON(),
		SourceRef:   targetRef,
		Findings:    report.findings,
	}, nil
}

type grypeScanner struct {
	path string
}

func (s grypeScanner) ScanDigest(ctx context.Context, target audit.ActiveDigestRef) (Result, error) {
	targetRef, err := targetReference(target)
	if err != nil {
		return Result{}, err
	}

	startedAt := time.Now().UTC()
	cmd := exec.CommandContext(ctx, s.path, targetRef, "-o", "json")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return Result{}, formatScanError("grype", targetRef, err, stderr.String())
	}

	report, err := parseGrypeReport(stdout.Bytes())
	if err != nil {
		return Result{}, err
	}
	return Result{
		ImageDigest: target.ImageDigest,
		ImageRef:    target.ImageRef,
		Scanner:     ScannerGrype,
		StartedAt:   startedAt,
		CompletedAt: time.Now().UTC(),
		Status:      audit.VulnerabilityScanStatusCompleted,
		Summary:     report.summaryJSON(),
		SourceRef:   targetRef,
		Findings:    report.findings,
	}, nil
}

type parsedReport struct {
	findings []audit.VulnerabilityFindingInput
	summary  map[string]int
}

func (p parsedReport) summaryJSON() json.RawMessage {
	summary := map[string]int{
		"critical": p.summary["CRITICAL"],
		"high":     p.summary["HIGH"],
		"medium":   p.summary["MEDIUM"],
		"low":      p.summary["LOW"],
		"unknown":  p.summary["UNKNOWN"],
		"total":    len(p.findings),
	}
	encoded, _ := json.Marshal(summary)
	return encoded
}

type trivyReport struct {
	Results []struct {
		Vulnerabilities []struct {
			ID               string `json:"VulnerabilityID"`
			PackageName      string `json:"PkgName"`
			InstalledVersion string `json:"InstalledVersion"`
			FixedVersion     string `json:"FixedVersion"`
			Title            string `json:"Title"`
			Description      string `json:"Description"`
			Severity         string `json:"Severity"`
			PURL             string `json:"PkgIdentifier.PURL"`
		} `json:"Vulnerabilities"`
	} `json:"Results"`
}

func parseTrivyReport(payload []byte) (parsedReport, error) {
	var report struct {
		Results []struct {
			Vulnerabilities []struct {
				VulnerabilityID  string `json:"VulnerabilityID"`
				PkgName          string `json:"PkgName"`
				InstalledVersion string `json:"InstalledVersion"`
				FixedVersion     string `json:"FixedVersion"`
				Title            string `json:"Title"`
				Description      string `json:"Description"`
				Severity         string `json:"Severity"`
				PkgIdentifier    struct {
					PURL string `json:"PURL"`
				} `json:"PkgIdentifier"`
				DataSource struct {
					Name string `json:"Name"`
				} `json:"DataSource"`
			} `json:"Vulnerabilities"`
		} `json:"Results"`
	}
	if err := json.Unmarshal(payload, &report); err != nil {
		return parsedReport{}, fmt.Errorf("invalid trivy json output: %w", err)
	}

	result := parsedReport{
		findings: []audit.VulnerabilityFindingInput{},
		summary:  map[string]int{},
	}
	for _, section := range report.Results {
		for _, vulnerability := range section.Vulnerabilities {
			cveID := strings.TrimSpace(strings.ToUpper(vulnerability.VulnerabilityID))
			if cveID == "" {
				continue
			}
			severity := normalizeSeverity(vulnerability.Severity)
			result.summary[severity]++
			result.findings = append(result.findings, audit.VulnerabilityFindingInput{
				CVEID:          cveID,
				Severity:       severity,
				PackageName:    strings.TrimSpace(vulnerability.PkgName),
				PackageVersion: strings.TrimSpace(vulnerability.InstalledVersion),
				FixedVersion:   strings.TrimSpace(vulnerability.FixedVersion),
				PURL:           strings.TrimSpace(vulnerability.PkgIdentifier.PURL),
				Title:          strings.TrimSpace(vulnerability.Title),
				Description:    strings.TrimSpace(vulnerability.Description),
				Source:         firstNonEmpty(vulnerability.DataSource.Name, ScannerTrivy),
			})
		}
	}
	return result, nil
}

func parseGrypeReport(payload []byte) (parsedReport, error) {
	var report struct {
		Matches []struct {
			Vulnerability struct {
				ID          string `json:"id"`
				Severity    string `json:"severity"`
				Description string `json:"description"`
				DataSource  string `json:"dataSource"`
				Fix         struct {
					Versions []string `json:"versions"`
				} `json:"fix"`
			} `json:"vulnerability"`
			Artifact struct {
				Name    string `json:"name"`
				Version string `json:"version"`
				PURL    string `json:"purl"`
			} `json:"artifact"`
		} `json:"matches"`
	}
	if err := json.Unmarshal(payload, &report); err != nil {
		return parsedReport{}, fmt.Errorf("invalid grype json output: %w", err)
	}

	result := parsedReport{
		findings: []audit.VulnerabilityFindingInput{},
		summary:  map[string]int{},
	}
	for _, match := range report.Matches {
		cveID := strings.TrimSpace(strings.ToUpper(match.Vulnerability.ID))
		if cveID == "" {
			continue
		}
		severity := normalizeSeverity(match.Vulnerability.Severity)
		result.summary[severity]++
		result.findings = append(result.findings, audit.VulnerabilityFindingInput{
			CVEID:          cveID,
			Severity:       severity,
			PackageName:    strings.TrimSpace(match.Artifact.Name),
			PackageVersion: strings.TrimSpace(match.Artifact.Version),
			FixedVersion:   firstNonEmpty(match.Vulnerability.Fix.Versions...),
			PURL:           strings.TrimSpace(match.Artifact.PURL),
			Title:          cveID,
			Description:    strings.TrimSpace(match.Vulnerability.Description),
			Source:         firstNonEmpty(match.Vulnerability.DataSource, ScannerGrype),
		})
	}
	return result, nil
}

func targetReference(target audit.ActiveDigestRef) (string, error) {
	imageRef := strings.TrimSpace(target.ImageRef)
	digest := strings.TrimSpace(target.ImageDigest)
	switch {
	case imageRef != "" && strings.Contains(imageRef, "@"):
		return imageRef, nil
	case imageRef != "" && digest != "":
		return imageRef + "@" + digest, nil
	case imageRef != "":
		return imageRef, nil
	default:
		return "", fmt.Errorf("image_ref is required to scan digest %s", digest)
	}
}

func normalizeScannerName(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func normalizeSeverity(value string) string {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "CRITICAL":
		return "CRITICAL"
	case "HIGH":
		return "HIGH"
	case "MEDIUM":
		return "MEDIUM"
	case "LOW":
		return "LOW"
	default:
		return "UNKNOWN"
	}
}

func parseBoolEnv(key string, fallback bool) bool {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback
	}
	switch strings.ToLower(raw) {
	case "1", "true", "yes", "on":
		return true
	case "0", "false", "no", "off":
		return false
	default:
		return fallback
	}
}

func parseDurationEnv(key string, fallback time.Duration) time.Duration {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(raw)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}

func formatScanError(scanner, target string, err error, stderr string) error {
	message := strings.TrimSpace(stderr)
	if message == "" {
		message = err.Error()
	}
	return fmt.Errorf("%s scan failed for %s: %s", scanner, target, message)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			return trimmed
		}
	}
	return ""
}
