package preflightcli

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/verify"
	internalvulnops "github.com/denisgrosek/changelock/internal/vulnops"
)

var ErrCommandNotFound = errors.New("command not found")

type VersionInfo struct {
	Version string
	Commit  string
	Date    string
}

type CommandExecution struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

type ScanFinding struct {
	CVEID          string `json:"cve_id"`
	Severity       string `json:"severity"`
	PackageName    string `json:"package_name,omitempty"`
	PackageVersion string `json:"package_version,omitempty"`
	FixedVersion   string `json:"fixed_version,omitempty"`
}

type ScanSummary struct {
	Scanner  string         `json:"scanner"`
	Image    string         `json:"image"`
	Counts   map[string]int `json:"counts"`
	Findings []ScanFinding  `json:"findings,omitempty"`
}

type Runtime struct {
	RunCommand     func(ctx context.Context, name string, args ...string) (CommandExecution, error)
	VerifyArtifact func(ctx context.Context, cosignBin string, request verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error)
	ScanImage      func(ctx context.Context, config Config, image string) (ScanSummary, error)
	HTTPClient     *http.Client
	VersionInfo    VersionInfo
}

func DefaultRuntime(version VersionInfo) Runtime {
	return Runtime{
		RunCommand:     runCommand,
		VerifyArtifact: verifyArtifact,
		ScanImage:      scanImage,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		VersionInfo: version,
	}
}

func runCommand(ctx context.Context, name string, args ...string) (CommandExecution, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return CommandExecution{}, fmt.Errorf("%w: %s", ErrCommandNotFound, name)
		}
		if exitErr, ok := err.(*exec.ExitError); ok {
			return CommandExecution{
				Stdout:   stdout.String(),
				Stderr:   stderr.String(),
				ExitCode: exitErr.ExitCode(),
			}, nil
		}
		return CommandExecution{
			Stdout: stdout.String(),
			Stderr: stderr.String(),
		}, err
	}
	return CommandExecution{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: 0,
	}, nil
}

func verifyArtifact(ctx context.Context, cosignBin string, request verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
	return verify.NewCosignVerifier(cosignBin).VerifyArtifact(ctx, request)
}

func scanImage(ctx context.Context, config Config, image string) (ScanSummary, error) {
	scannerName := strings.ToLower(strings.TrimSpace(config.Scanner))
	if scannerName == "" || scannerName == "auto" {
		switch {
		case lookPath(config.TrivyBin) == nil:
			scannerName = internalvulnops.ScannerTrivy
		case lookPath(config.GrypeBin) == nil:
			scannerName = internalvulnops.ScannerGrype
		default:
			return ScanSummary{}, fmt.Errorf("%w: no vulnerability scanner found in PATH", ErrCommandNotFound)
		}
	}

	scanner, err := internalvulnops.NewScanner(internalvulnops.Config{
		Enabled:   true,
		Scanner:   scannerName,
		TrivyPath: config.TrivyBin,
		GrypePath: config.GrypeBin,
	})
	if err != nil {
		return ScanSummary{}, err
	}

	result, err := scanner.ScanDigest(ctx, audit.ActiveDigestRef{
		ImageRef:    image,
		ImageDigest: audit.DigestFromImage(image),
	})
	if err != nil {
		return ScanSummary{}, err
	}

	summary := ScanSummary{
		Scanner: scannerName,
		Image:   image,
		Counts: map[string]int{
			"CRITICAL": 0,
			"HIGH":     0,
			"MEDIUM":   0,
			"LOW":      0,
			"UNKNOWN":  0,
		},
		Findings: make([]ScanFinding, 0, len(result.Findings)),
	}
	for _, finding := range result.Findings {
		severity := strings.ToUpper(strings.TrimSpace(finding.Severity))
		if severity == "" {
			severity = "UNKNOWN"
		}
		if _, ok := summary.Counts[severity]; !ok {
			severity = "UNKNOWN"
		}
		summary.Counts[severity]++
		summary.Findings = append(summary.Findings, ScanFinding{
			CVEID:          finding.CVEID,
			Severity:       severity,
			PackageName:    finding.PackageName,
			PackageVersion: finding.PackageVersion,
			FixedVersion:   finding.FixedVersion,
		})
	}
	return summary, nil
}

func lookPath(binary string) error {
	if strings.TrimSpace(binary) == "" {
		return fmt.Errorf("%w: empty binary name", ErrCommandNotFound)
	}
	_, err := exec.LookPath(binary)
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return fmt.Errorf("%w: %s", ErrCommandNotFound, binary)
		}
		return err
	}
	return nil
}
