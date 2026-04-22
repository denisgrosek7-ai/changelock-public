package preflightcli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/policy"
	"github.com/denisgrosek/changelock/internal/verify"
)

var defaultAllowedOIDCIssuers = []string{
	"https://token.actions.githubusercontent.com",
}

type Config struct {
	Output               string
	APIURL               string
	Token                string
	Timeout              time.Duration
	Offline              bool
	ProductionConfigPath string
	PolicyDir            string
	KyvernoPolicyDir     string
	FailureSeverity      string
	Scanner              string
	KyvernoBin           string
	CosignBin            string
	TrivyBin             string
	GrypeBin             string
}

type App struct {
	config  Config
	runtime Runtime
}

type usageError struct {
	message string
}

func (e usageError) Error() string {
	return e.message
}

func NewApp(getenv func(string) string, runtime Runtime) (*App, error) {
	config, err := loadConfig(getenv)
	if err != nil {
		return nil, err
	}
	if runtime.RunCommand == nil || runtime.VerifyArtifact == nil || runtime.ScanImage == nil {
		return nil, errors.New("preflight runtime is incomplete")
	}
	if runtime.HTTPClient == nil {
		runtime.HTTPClient = &http.Client{Timeout: apiTimeout(config)}
	}
	return &App{config: config, runtime: runtime}, nil
}

func loadConfig(getenv func(string) string) (Config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}
	output := firstNonEmpty(getenv("CHANGELOCK_CLI_OUTPUT"), "human")
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "human", "json":
	default:
		return Config{}, fmt.Errorf("unsupported CHANGELOCK_CLI_OUTPUT %q", output)
	}

	timeout := 2 * time.Minute
	if raw := strings.TrimSpace(getenv("CHANGELOCK_CLI_TIMEOUT")); raw != "" {
		parsed, err := time.ParseDuration(raw)
		if err != nil || parsed <= 0 {
			return Config{}, fmt.Errorf("invalid CHANGELOCK_CLI_TIMEOUT %q", raw)
		}
		timeout = parsed
	}

	offline, err := parseBool(getenv("CHANGELOCK_CLI_OFFLINE"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_CLI_OFFLINE: %w", err)
	}

	failureSeverity := normalizeSeverity(firstNonEmpty(getenv("CHANGELOCK_VULN_FAIL_SEVERITY"), "CRITICAL"))
	if !validSeverity(failureSeverity) {
		return Config{}, fmt.Errorf("unsupported vulnerability failure severity %q", failureSeverity)
	}

	scanner := strings.ToLower(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_CLI_SCANNER"), "auto")))
	switch scanner {
	case "auto", "trivy", "grype":
	default:
		return Config{}, fmt.Errorf("unsupported CHANGELOCK_CLI_SCANNER %q", scanner)
	}

	return Config{
		Output:               output,
		APIURL:               strings.TrimSpace(getenv("CHANGELOCK_CLI_API_URL")),
		Token:                strings.TrimSpace(getenv("CHANGELOCK_CLI_TOKEN")),
		Timeout:              timeout,
		Offline:              offline,
		ProductionConfigPath: strings.TrimSpace(getenv("CHANGELOCK_CLI_CONFIG")),
		PolicyDir:            firstNonEmpty(getenv("CHANGELOCK_CLI_POLICY_DIR"), "policies"),
		KyvernoPolicyDir:     firstNonEmpty(getenv("CHANGELOCK_CLI_KYVERNO_POLICY_DIR"), "deploy/kyverno"),
		FailureSeverity:      failureSeverity,
		Scanner:              scanner,
		KyvernoBin:           firstNonEmpty(getenv("CHANGELOCK_CLI_KYVERNO_BIN"), "kyverno"),
		CosignBin:            firstNonEmpty(getenv("CHANGELOCK_CLI_COSIGN_BIN"), "cosign"),
		TrivyBin:             firstNonEmpty(getenv("CHANGELOCK_CLI_TRIVY_BIN"), "trivy"),
		GrypeBin:             firstNonEmpty(getenv("CHANGELOCK_CLI_GRYPE_BIN"), "grype"),
	}, nil
}

func (a *App) Run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		return a.writeError(stderr, usageError{message: a.usage()})
	}
	if a.config.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, a.config.Timeout)
		defer cancel()
	}

	command := strings.TrimSpace(args[0])
	switch command {
	case "version":
		return a.runVersion(stdout)
	case "manifest":
		result, err := a.runManifest(ctx, args[1:])
		return a.finish(result, err, stdout, stderr)
	case "image":
		result, err := a.runImage(ctx, args[1:])
		return a.finish(result, err, stdout, stderr)
	case "scan":
		result, err := a.runScan(ctx, args[1:])
		return a.finish(result, err, stdout, stderr)
	case "preflight":
		result, err := a.runPreflight(ctx, args[1:])
		return a.finish(result, err, stdout, stderr)
	case "check":
		result, err := a.runCheck(ctx, args[1:])
		return a.finish(result, err, stdout, stderr)
	case "preview":
		result, err := a.runPreview(ctx, args[1:])
		return a.finish(result, err, stdout, stderr)
	case "inspect":
		return a.runInspect(ctx, args[1:], stdout, stderr)
	case "explain":
		return a.runExplain(ctx, args[1:], stdout, stderr)
	case "readiness":
		result, err := a.runReadiness(ctx, args[1:])
		return a.finish(result, err, stdout, stderr)
	case "support":
		return a.runSupport(ctx, args[1:], stdout, stderr)
	case "upgrade-readiness":
		return a.runUpgradeReadiness(ctx, args[1:], stdout, stderr)
	case "phase5-summary":
		return a.runPhase5Summary(ctx, args[1:], stdout, stderr)
	case "diagnostics":
		return a.runDiagnostics(args[1:], stdout, stderr)
	case "guidance":
		return a.runGuidance(args[1:], stdout, stderr)
	case "help", "-h", "--help":
		_, _ = io.WriteString(stdout, a.usage())
		return ExitUsage
	default:
		return a.writeError(stderr, usageError{message: fmt.Sprintf("unknown command %q\n\n%s", command, a.usage())})
	}
}

func (a *App) finish(result Result, err error, stdout, stderr io.Writer) int {
	if err != nil {
		return a.writeError(stderr, err)
	}
	if renderErr := renderResult(stdout, a.config.Output, result); renderErr != nil {
		return a.writeError(stderr, renderErr)
	}
	return exitCodeForResult(result)
}

func (a *App) writeError(stderr io.Writer, err error) int {
	if err == nil {
		return ExitSuccess
	}
	_, _ = fmt.Fprintln(stderr, err.Error())
	var usage usageError
	if errors.As(err, &usage) {
		return ExitUsage
	}
	return ExitExecution
}

func (a *App) runVersion(stdout io.Writer) int {
	version := firstNonEmpty(a.runtime.VersionInfo.Version, "dev")
	commit := firstNonEmpty(a.runtime.VersionInfo.Commit, "unknown")
	date := firstNonEmpty(a.runtime.VersionInfo.Date, "unknown")
	_, _ = fmt.Fprintf(stdout, "changelock-cli version=%s commit=%s date=%s\n", version, commit, date)
	return ExitSuccess
}

func (a *App) runManifest(ctx context.Context, args []string) (Result, error) {
	options, err := a.parseManifestOptions(args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.Output
	checks, err := a.manifestChecks(ctx, options)
	if err != nil {
		return Result{}, err
	}
	result := Result{
		Command: "manifest",
		Mode:    ModeLocalOnly,
		Inputs: map[string]string{
			"policy_dir": options.PolicyDir,
		},
		Checks: checks,
	}
	return finalizeResult(result), nil
}

func (a *App) runImage(ctx context.Context, args []string) (Result, error) {
	options, err := a.parseImageOptions(args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.Output
	checks, err := a.imageChecks(ctx, options)
	if err != nil {
		return Result{}, err
	}
	result := Result{
		Command: "image",
		Mode:    executionMode(options.Offline, options.APIURL),
		Inputs: map[string]string{
			"image":      options.Image,
			"tenant":     options.Tenant,
			"repository": options.Repository,
			"bundle_dir": options.BundleDir,
		},
		Checks: checks,
	}
	return finalizeResult(result), nil
}

func (a *App) runScan(ctx context.Context, args []string) (Result, error) {
	options, err := a.parseScanOptions(args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.Output
	checks, err := a.scanChecks(ctx, options)
	if err != nil {
		return Result{}, err
	}
	result := Result{
		Command: "scan",
		Mode:    executionMode(options.Offline, options.APIURL),
		Inputs: map[string]string{
			"image":         options.Image,
			"fail_severity": options.FailSeverity,
			"scanner":       options.Scanner,
			"tenant":        options.Tenant,
			"repository":    options.Repository,
			"environment":   options.Environment,
		},
		Checks: checks,
	}
	return finalizeResult(result), nil
}

func (a *App) runPreflight(ctx context.Context, args []string) (Result, error) {
	options, err := a.parsePreflightOptions(args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.manifestOptions.Output
	if len(options.manifestOptions.Files) == 0 && len(options.manifestOptions.Dirs) == 0 && strings.TrimSpace(options.imageOptions.Image) == "" {
		return Result{}, usageError{message: "preflight requires at least one manifest or image input"}
	}
	result := Result{
		Command: "preflight",
		Mode:    executionMode(options.imageOptions.Offline, options.imageOptions.APIURL),
		Inputs: map[string]string{
			"image":       options.imageOptions.Image,
			"tenant":      options.imageOptions.Tenant,
			"repository":  options.imageOptions.Repository,
			"environment": options.imageOptions.Environment,
			"namespace":   options.imageOptions.Namespace,
			"bundle_dir":  options.imageOptions.BundleDir,
			"policy_dir":  options.manifestOptions.PolicyDir,
		},
	}

	if options.imageOptions.APIURL != "" && !options.imageOptions.Offline {
		authCheck := a.remoteAuthCheck(ctx, options)
		result.add(authCheck)
	}

	if len(options.manifestOptions.Files) > 0 || len(options.manifestOptions.Dirs) > 0 {
		checks, err := a.manifestChecks(ctx, manifestOptions{
			Files:      options.manifestOptions.Files,
			Dirs:       options.manifestOptions.Dirs,
			PolicyDir:  options.manifestOptions.PolicyDir,
			KyvernoBin: options.manifestOptions.KyvernoBin,
		})
		if err != nil {
			return Result{}, err
		}
		result.add(checks...)
	}

	if options.imageOptions.Image != "" {
		imageChecks, err := a.imageChecks(ctx, imageOptions{
			Image:       options.imageOptions.Image,
			Tenant:      options.imageOptions.Tenant,
			Repository:  options.imageOptions.Repository,
			Environment: options.imageOptions.Environment,
			Namespace:   options.imageOptions.Namespace,
			BundleDir:   options.imageOptions.BundleDir,
			CosignBin:   options.imageOptions.CosignBin,
			APIURL:      options.imageOptions.APIURL,
			Token:       options.imageOptions.Token,
			Offline:     options.imageOptions.Offline,
			ExpectedRef: options.imageOptions.ExpectedRef,
			CommitSHA:   options.imageOptions.CommitSHA,
			OIDCIssuers: options.imageOptions.OIDCIssuers,
		})
		if err != nil {
			return Result{}, err
		}
		result.add(imageChecks...)

		scanChecks, err := a.scanChecks(ctx, scanOptions{
			Image:        options.scanOptions.Image,
			Tenant:       options.scanOptions.Tenant,
			Repository:   options.scanOptions.Repository,
			Environment:  options.scanOptions.Environment,
			FailSeverity: options.scanOptions.FailSeverity,
			Scanner:      options.scanOptions.Scanner,
			APIURL:       options.scanOptions.APIURL,
			Token:        options.scanOptions.Token,
			Offline:      options.scanOptions.Offline,
		})
		if err != nil {
			return Result{}, err
		}
		result.add(scanChecks...)
	}

	return finalizeResult(result), nil
}

type manifestOptions struct {
	Files      []string
	Dirs       []string
	PolicyDir  string
	KyvernoBin string
	Output     string
}

func (a *App) parseManifestOptions(args []string) (manifestOptions, error) {
	fs := newFlagSet("manifest")
	var files, dirs stringSlice
	options := manifestOptions{
		PolicyDir:  a.config.KyvernoPolicyDir,
		KyvernoBin: a.config.KyvernoBin,
		Output:     a.config.Output,
	}
	fs.Var(&files, "file", "manifest file to validate")
	fs.Var(&dirs, "dir", "manifest directory to validate recursively")
	fs.StringVar(&options.PolicyDir, "policy-dir", options.PolicyDir, "kyverno policy directory")
	fs.StringVar(&options.KyvernoBin, "kyverno-bin", options.KyvernoBin, "kyverno binary path")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	if err := fs.Parse(args); err != nil {
		return manifestOptions{}, usageError{message: err.Error()}
	}
	options.Files = append(options.Files, files...)
	options.Dirs = append(options.Dirs, dirs...)
	for _, arg := range fs.Args() {
		options.Files = append(options.Files, arg)
	}
	if len(options.Files) == 0 && len(options.Dirs) == 0 {
		return manifestOptions{}, usageError{message: "manifest requires at least one --file, --dir, or positional manifest path"}
	}
	return options, nil
}

type imageOptions struct {
	Image       string
	Tenant      string
	Repository  string
	Environment string
	Namespace   string
	BundleDir   string
	CosignBin   string
	APIURL      string
	Token       string
	Offline     bool
	ExpectedRef string
	CommitSHA   string
	OIDCIssuers []string
	Output      string
}

func (a *App) parseImageOptions(args []string) (imageOptions, error) {
	fs := newFlagSet("image")
	var issuers stringSlice
	options := imageOptions{
		Tenant:    "acme",
		BundleDir: a.config.PolicyDir,
		CosignBin: a.config.CosignBin,
		APIURL:    a.config.APIURL,
		Token:     a.config.Token,
		Offline:   a.config.Offline,
		Output:    a.config.Output,
	}
	fs.StringVar(&options.Image, "image", "", "image reference to evaluate")
	fs.StringVar(&options.Tenant, "tenant", options.Tenant, "tenant name for bundle loading")
	fs.StringVar(&options.Repository, "repository", "", "repository identity, for example my-org/acme-app")
	fs.StringVar(&options.Environment, "environment", "", "environment name")
	fs.StringVar(&options.Namespace, "namespace", "", "kubernetes namespace")
	fs.StringVar(&options.BundleDir, "bundle-dir", options.BundleDir, "ChangeLock policy bundle directory")
	fs.StringVar(&options.CosignBin, "cosign-bin", options.CosignBin, "cosign binary path")
	fs.StringVar(&options.APIURL, "api-url", options.APIURL, "optional ChangeLock API URL")
	fs.StringVar(&options.Token, "token", options.Token, "optional bearer token for ChangeLock API")
	fs.BoolVar(&options.Offline, "offline", options.Offline, "disable API-assisted checks")
	fs.StringVar(&options.ExpectedRef, "workflow-ref", "", "expected workflow ref")
	fs.StringVar(&options.CommitSHA, "commit-sha", "", "expected workflow commit sha")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	fs.Var(&issuers, "oidc-issuer", "allowed OIDC issuer for cosign verification")
	if err := fs.Parse(args); err != nil {
		return imageOptions{}, usageError{message: err.Error()}
	}
	if options.Image == "" && len(fs.Args()) > 0 {
		options.Image = fs.Args()[0]
	}
	if strings.TrimSpace(options.Image) == "" {
		return imageOptions{}, usageError{message: "image command requires an image reference"}
	}
	options.OIDCIssuers = issuers
	return options, nil
}

type scanOptions struct {
	Image        string
	Tenant       string
	Repository   string
	Environment  string
	FailSeverity string
	Scanner      string
	APIURL       string
	Token        string
	Offline      bool
	Output       string
}

func (a *App) parseScanOptions(args []string) (scanOptions, error) {
	fs := newFlagSet("scan")
	options := scanOptions{
		Tenant:       "acme",
		FailSeverity: a.config.FailureSeverity,
		Scanner:      a.config.Scanner,
		APIURL:       a.config.APIURL,
		Token:        a.config.Token,
		Offline:      a.config.Offline,
		Output:       a.config.Output,
	}
	fs.StringVar(&options.Image, "image", "", "image reference to scan")
	fs.StringVar(&options.Tenant, "tenant", options.Tenant, "tenant name")
	fs.StringVar(&options.Repository, "repository", "", "repository identity")
	fs.StringVar(&options.Environment, "environment", "", "environment name")
	fs.StringVar(&options.FailSeverity, "fail-severity", options.FailSeverity, "severity threshold that fails the scan")
	fs.StringVar(&options.Scanner, "scanner", options.Scanner, "scanner selection: auto|trivy|grype")
	fs.StringVar(&options.APIURL, "api-url", options.APIURL, "optional ChangeLock API URL")
	fs.StringVar(&options.Token, "token", options.Token, "optional bearer token for ChangeLock API")
	fs.BoolVar(&options.Offline, "offline", options.Offline, "disable API-assisted checks")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	if err := fs.Parse(args); err != nil {
		return scanOptions{}, usageError{message: err.Error()}
	}
	if options.Image == "" && len(fs.Args()) > 0 {
		options.Image = fs.Args()[0]
	}
	if strings.TrimSpace(options.Image) == "" {
		return scanOptions{}, usageError{message: "scan command requires an image reference"}
	}
	options.FailSeverity = normalizeSeverity(options.FailSeverity)
	if !validSeverity(options.FailSeverity) {
		return scanOptions{}, usageError{message: fmt.Sprintf("unsupported fail severity %q", options.FailSeverity)}
	}
	return options, nil
}

type preflightOptions struct {
	manifestOptions
	imageOptions
	scanOptions
}

func (a *App) parsePreflightOptions(args []string) (preflightOptions, error) {
	fs := newFlagSet("preflight")
	var files, dirs, issuers stringSlice
	options := preflightOptions{
		manifestOptions: manifestOptions{
			PolicyDir:  a.config.KyvernoPolicyDir,
			KyvernoBin: a.config.KyvernoBin,
			Output:     a.config.Output,
		},
		imageOptions: imageOptions{
			Tenant:    "acme",
			BundleDir: a.config.PolicyDir,
			CosignBin: a.config.CosignBin,
			APIURL:    a.config.APIURL,
			Token:     a.config.Token,
			Offline:   a.config.Offline,
			Output:    a.config.Output,
		},
		scanOptions: scanOptions{
			Tenant:       "acme",
			FailSeverity: a.config.FailureSeverity,
			Scanner:      a.config.Scanner,
			APIURL:       a.config.APIURL,
			Token:        a.config.Token,
			Offline:      a.config.Offline,
			Output:       a.config.Output,
		},
	}
	fs.Var(&files, "file", "manifest file to validate")
	fs.Var(&dirs, "dir", "manifest directory to validate recursively")
	fs.StringVar(&options.manifestOptions.PolicyDir, "policy-dir", options.manifestOptions.PolicyDir, "kyverno policy directory")
	fs.StringVar(&options.manifestOptions.KyvernoBin, "kyverno-bin", options.manifestOptions.KyvernoBin, "kyverno binary path")
	fs.StringVar(&options.imageOptions.Image, "image", "", "image reference to assess")
	fs.StringVar(&options.imageOptions.Tenant, "tenant", options.imageOptions.Tenant, "tenant name")
	fs.StringVar(&options.imageOptions.Repository, "repository", "", "repository identity")
	fs.StringVar(&options.imageOptions.Environment, "environment", "", "environment name")
	fs.StringVar(&options.imageOptions.Namespace, "namespace", "", "namespace")
	fs.StringVar(&options.imageOptions.BundleDir, "bundle-dir", options.imageOptions.BundleDir, "ChangeLock policy bundle directory")
	fs.StringVar(&options.imageOptions.CosignBin, "cosign-bin", options.imageOptions.CosignBin, "cosign binary path")
	fs.StringVar(&options.scanOptions.FailSeverity, "fail-severity", options.scanOptions.FailSeverity, "scan failure severity threshold")
	fs.StringVar(&options.scanOptions.Scanner, "scanner", options.scanOptions.Scanner, "scanner selection: auto|trivy|grype")
	fs.StringVar(&options.imageOptions.APIURL, "api-url", options.imageOptions.APIURL, "optional ChangeLock API URL")
	fs.StringVar(&options.imageOptions.Token, "token", options.imageOptions.Token, "optional bearer token for ChangeLock API")
	fs.BoolVar(&options.imageOptions.Offline, "offline", options.imageOptions.Offline, "disable API-assisted checks")
	fs.StringVar(&options.imageOptions.ExpectedRef, "workflow-ref", "", "expected workflow ref")
	fs.StringVar(&options.imageOptions.CommitSHA, "commit-sha", "", "expected workflow commit sha")
	fs.StringVar(&options.manifestOptions.Output, "output", options.manifestOptions.Output, "output mode: human|json")
	fs.Var(&issuers, "oidc-issuer", "allowed OIDC issuer for cosign verification")
	if err := fs.Parse(args); err != nil {
		return preflightOptions{}, usageError{message: err.Error()}
	}
	options.manifestOptions.Files = files
	options.manifestOptions.Dirs = dirs
	if options.imageOptions.Image == "" && len(fs.Args()) > 0 {
		options.imageOptions.Image = fs.Args()[0]
	}
	options.imageOptions.OIDCIssuers = issuers
	options.imageOptions.Output = options.manifestOptions.Output
	options.scanOptions.Image = options.imageOptions.Image
	options.scanOptions.Tenant = options.imageOptions.Tenant
	options.scanOptions.Repository = options.imageOptions.Repository
	options.scanOptions.Environment = options.imageOptions.Environment
	options.scanOptions.APIURL = options.imageOptions.APIURL
	options.scanOptions.Token = options.imageOptions.Token
	options.scanOptions.Offline = options.imageOptions.Offline
	options.scanOptions.Output = options.manifestOptions.Output
	options.scanOptions.FailSeverity = normalizeSeverity(options.scanOptions.FailSeverity)
	if !validSeverity(options.scanOptions.FailSeverity) {
		return preflightOptions{}, usageError{message: fmt.Sprintf("unsupported fail severity %q", options.scanOptions.FailSeverity)}
	}
	return options, nil
}

func (a *App) manifestChecks(ctx context.Context, options manifestOptions) ([]CheckResult, error) {
	resources, err := collectManifestFiles(options.Files, options.Dirs)
	if err != nil {
		return nil, usageError{message: err.Error()}
	}
	if len(resources) == 0 {
		return nil, usageError{message: "no manifest files found"}
	}

	checks := make([]CheckResult, 0, len(resources))
	for _, resource := range resources {
		commandResult, execErr := a.runtime.RunCommand(ctx, options.KyvernoBin, "apply", options.PolicyDir, "--resource", resource)
		check := CheckResult{
			Name:   "manifest",
			Mode:   ModeLocal,
			Target: resource,
			Metadata: map[string]any{
				"resource_kind": "manifest",
			},
		}
		if execErr != nil {
			check.Status = StatusError
			check.Summary = execErr.Error()
			check.Details = truncateLines(commandResult.Stderr, 8)
			checks = append(checks, check)
			continue
		}
		if commandResult.ExitCode == 0 {
			check.Status = StatusPass
			check.Summary = "Kyverno accepted the manifest against the local policy set"
			check.Details = truncateLines(firstNonEmpty(commandResult.Stdout, commandResult.Stderr), 4)
		} else {
			check.Status = StatusFail
			check.Summary = "Kyverno reported policy violations"
			check.Details = truncateLines(firstNonEmpty(commandResult.Stdout, commandResult.Stderr), 8)
		}
		checks = append(checks, check)
	}
	return checks, nil
}

func (a *App) imageChecks(ctx context.Context, options imageOptions) ([]CheckResult, error) {
	bundle, err := policy.LoadBundle(options.BundleDir, options.Tenant)
	if err != nil {
		return []CheckResult{{
			Name:    "image-policy",
			Status:  StatusError,
			Target:  options.Image,
			Summary: "unable to load local ChangeLock policy bundle",
			Details: []string{err.Error()},
		}}, nil
	}

	repository := normalizeRepository(firstNonEmpty(options.Repository, inferRepositoryFromImage(options.Image), firstTenantRepository(bundle)))
	verificationRequest := verify.ArtifactVerificationRequest{
		Image:                   options.Image,
		ExpectedRepository:      repository,
		ExpectedRef:             options.ExpectedRef,
		ExpectedCommitSHA:       options.CommitSHA,
		AllowedSignerIdentities: bundle.Artifact.Spec.AllowedSignerIdentities,
		AllowedOIDCIssuers:      firstNonEmptySlice(options.OIDCIssuers, defaultAllowedOIDCIssuers),
	}

	checks := []CheckResult{{
		Name:    "image-digest",
		Mode:    ModeLocal,
		Target:  options.Image,
		Status:  statusForBool(strings.Contains(options.Image, "@sha256:"), StatusPass, StatusFail),
		Summary: summaryForBool(strings.Contains(options.Image, "@sha256:"), "image is digest-pinned", "image is not digest-pinned"),
		Metadata: map[string]any{
			"digest_pinned": strings.Contains(options.Image, "@sha256:"),
		},
	}}

	var verification *verify.ArtifactVerification
	if !strings.Contains(options.Image, "@sha256:") {
		checks = append(checks, CheckResult{
			Name:    "image-trust",
			Mode:    ModeLocal,
			Target:  options.Image,
			Status:  StatusFail,
			Summary: "cosign verification requires a digest-pinned image reference",
		})
	} else {
		result, verifyErr := a.runtime.VerifyArtifact(ctx, options.CosignBin, verificationRequest)
		if verifyErr != nil {
			checks = append(checks, CheckResult{
				Name:    "image-trust",
				Mode:    ModeLocal,
				Target:  options.Image,
				Status:  StatusError,
				Summary: verifyErr.Error(),
				Metadata: map[string]any{
					"signature_valid":   false,
					"attestation_valid": false,
				},
			})
		} else {
			verification = &result
			status := StatusPass
			summary := "signature and attestation verification passed"
			if !result.SignatureValid || !result.AttestationValid || len(result.Reasons) > 0 {
				status = StatusFail
				summary = "signature and attestation verification did not satisfy ChangeLock expectations"
			}
			checks = append(checks, CheckResult{
				Name:    "image-trust",
				Mode:    ModeLocal,
				Target:  options.Image,
				Status:  status,
				Summary: summary,
				Details: truncateLines(strings.Join(result.Reasons, "\n"), 8),
				Metadata: map[string]any{
					"signature_valid":   result.SignatureValid,
					"attestation_valid": result.AttestationValid,
					"verified_identity": result.VerifiedIdentity,
				},
			})
		}
	}

	evaluation := policy.EvaluateArtifact(bundle, policy.ArtifactEvaluationRequest{
		Tenant:         options.Tenant,
		Repository:     repository,
		Environment:    options.Environment,
		Namespace:      options.Namespace,
		Image:          options.Image,
		Registry:       stripTagAndDigest(options.Image),
		DigestPinned:   strings.Contains(options.Image, "@sha256:"),
		HasProvenance:  verification != nil && verification.AttestationValid,
		HasSignature:   verification != nil && verification.SignatureValid,
		SignerIdentity: verifiedIdentity(verification),
		WorkflowFile:   verifiedWorkflow(verification),
		Subject:        firstNonEmpty(verifiedSubject(verification), repoSubject(repository)),
		Verification:   verification,
	})
	policyCheck := CheckResult{
		Name:    "image-policy",
		Mode:    ModeLocal,
		Target:  options.Image,
		Status:  StatusPass,
		Summary: "local ChangeLock artifact policy would allow this image",
		Metadata: map[string]any{
			"decision": evaluation.Decision,
		},
	}
	if evaluation.Decision != "ALLOW" {
		policyCheck.Status = StatusFail
		policyCheck.Summary = "local ChangeLock artifact policy would deny this image"
		policyCheck.Details = append(policyCheck.Details, evaluation.Reasons...)
	}
	checks = append(checks, policyCheck)

	checks = append(checks, a.remoteImageContextCheck(ctx, options, audit.DigestFromImage(options.Image))...)
	return checks, nil
}

func (a *App) scanChecks(ctx context.Context, options scanOptions) ([]CheckResult, error) {
	summary, err := a.runtime.ScanImage(ctx, Config{
		Scanner:  options.Scanner,
		TrivyBin: a.config.TrivyBin,
		GrypeBin: a.config.GrypeBin,
		Timeout:  a.config.Timeout,
	}, options.Image)
	if err != nil {
		return []CheckResult{{
			Name:    "scan",
			Mode:    ModeLocal,
			Target:  options.Image,
			Status:  StatusError,
			Summary: err.Error(),
		}}, nil
	}

	breaching := findingsAtOrAbove(summary.Findings, options.FailSeverity)
	scanCheck := CheckResult{
		Name:    "scan",
		Mode:    ModeLocal,
		Target:  options.Image,
		Status:  StatusPass,
		Summary: fmt.Sprintf("%s scan found no findings at or above %s", summary.Scanner, options.FailSeverity),
		Metadata: map[string]any{
			"scanner":         summary.Scanner,
			"counts":          summary.Counts,
			"breaching_count": len(breaching),
			"fail_severity":   options.FailSeverity,
		},
	}
	if len(breaching) > 0 {
		scanCheck.Status = StatusFail
		scanCheck.Summary = fmt.Sprintf("%s scan found %d findings at or above %s", summary.Scanner, len(breaching), options.FailSeverity)
	}
	scanCheck.Details = summarizeFindings(breaching, 5)

	checks := []CheckResult{scanCheck}
	checks = append(checks, a.remoteScanContextCheck(ctx, options, audit.DigestFromImage(options.Image), breaching)...)
	return checks, nil
}

func (a *App) remoteAuthCheck(ctx context.Context, options preflightOptions) CheckResult {
	client := NewAPIClient(Config{
		APIURL:  options.imageOptions.APIURL,
		Token:   options.imageOptions.Token,
		Timeout: a.config.Timeout,
		Offline: options.imageOptions.Offline,
	}, a.runtime.HTTPClient)
	if client == nil {
		return CheckResult{Name: "remote-auth", Mode: ModeRemote, Status: StatusSkip, Summary: "API-assisted checks disabled or no API URL configured"}
	}
	info, err := client.AuthMe(ctx)
	if err != nil {
		return CheckResult{Name: "remote-auth", Mode: ModeRemote, Status: StatusError, Summary: err.Error()}
	}
	scope := firstNonEmpty(info.TenantID, "global")
	return CheckResult{
		Name:    "remote-auth",
		Mode:    ModeRemote,
		Status:  StatusPass,
		Summary: fmt.Sprintf("authenticated as %s (%s scope)", firstNonEmpty(info.Role, info.AuthMode), scope),
		Metadata: map[string]any{
			"auth_mode": info.AuthMode,
			"role":      info.Role,
			"tenant_id": info.TenantID,
		},
	}
}

func (a *App) remoteImageContextCheck(ctx context.Context, options imageOptions, digest string) []CheckResult {
	client := NewAPIClient(Config{
		APIURL:  options.APIURL,
		Token:   options.Token,
		Timeout: a.config.Timeout,
		Offline: options.Offline,
	}, a.runtime.HTTPClient)
	if client == nil {
		return []CheckResult{{
			Name:    "remote-image-context",
			Mode:    ModeRemote,
			Target:  options.Image,
			Status:  StatusSkip,
			Summary: "API-assisted exception lookup disabled",
		}}
	}
	if digest == "" {
		return []CheckResult{{
			Name:    "remote-image-context",
			Mode:    ModeRemote,
			Target:  options.Image,
			Status:  StatusSkip,
			Summary: "digest-specific exception lookup skipped because image is not digest-pinned",
		}}
	}
	exceptions, err := client.ListExceptions(ctx, audit.ExceptionFilter{
		Status:      audit.ExceptionStatusApproved,
		ImageDigest: digest,
		TenantID:    options.Tenant,
		Environment: options.Environment,
		Repo:        options.Repository,
		Limit:       20,
	})
	if err != nil {
		return []CheckResult{{
			Name:    "remote-image-context",
			Mode:    ModeRemote,
			Target:  options.Image,
			Status:  StatusError,
			Summary: err.Error(),
		}}
	}
	details := make([]string, 0, len(exceptions))
	for _, exception := range exceptions {
		details = append(details, fmt.Sprintf("%s %s (%s)", exception.ExceptionID, exception.ExceptionType, exception.Status))
	}
	summary := "no approved digest exceptions found in ChangeLock"
	if len(exceptions) > 0 {
		summary = fmt.Sprintf("found %d approved digest exception(s) in ChangeLock", len(exceptions))
	}
	return []CheckResult{{
		Name:    "remote-image-context",
		Mode:    ModeRemote,
		Target:  options.Image,
		Status:  StatusPass,
		Summary: summary,
		Details: details,
		Metadata: map[string]any{
			"context_kind":    "exception-match",
			"exception_count": len(exceptions),
		},
	}}
}

func (a *App) remoteScanContextCheck(ctx context.Context, options scanOptions, digest string, findings []ScanFinding) []CheckResult {
	client := NewAPIClient(Config{
		APIURL:  options.APIURL,
		Token:   options.Token,
		Timeout: a.config.Timeout,
		Offline: options.Offline,
	}, a.runtime.HTTPClient)
	if client == nil {
		return []CheckResult{{
			Name:    "remote-scan-context",
			Mode:    ModeRemote,
			Target:  options.Image,
			Status:  StatusSkip,
			Summary: "API-assisted exception lookup disabled",
		}}
	}
	if len(findings) == 0 {
		return []CheckResult{{
			Name:    "remote-scan-context",
			Mode:    ModeRemote,
			Target:  options.Image,
			Status:  StatusSkip,
			Summary: "no threshold-breaching findings to match against ChangeLock exceptions",
			Metadata: map[string]any{
				"context_kind": "threshold-clear",
			},
		}}
	}
	if digest != "" {
		netResponse, err := client.VulnerabilityNet(ctx, digest, options.Tenant, options.Environment, options.FailSeverity)
		if err != nil {
			return []CheckResult{{
				Name:    "remote-scan-context",
				Mode:    ModeRemote,
				Target:  options.Image,
				Status:  StatusError,
				Summary: err.Error(),
				Metadata: map[string]any{
					"context_kind": "vex-net",
				},
			}}
		}
		details := make([]string, 0, min(5, len(netResponse.Findings)))
		for _, finding := range netResponse.Findings {
			disposition := "actionable"
			if finding.VEX != nil && strings.TrimSpace(finding.VEX.Status) != "" {
				disposition = finding.VEX.Status
			}
			details = append(details, fmt.Sprintf("%s %s %s", firstNonEmpty(finding.Severity, "UNKNOWN"), finding.CVEID, disposition))
			if len(details) == 5 {
				break
			}
		}
		status := StatusPass
		summary := fmt.Sprintf("net actionable vulnerability context shows %d remaining finding(s) after VEX merge", netResponse.ActionableCount)
		if netResponse.ThresholdBreached {
			status = StatusFail
			summary = fmt.Sprintf("net actionable vulnerability context still breaches %s after VEX merge", options.FailSeverity)
		}
		return []CheckResult{{
			Name:    "remote-scan-context",
			Mode:    ModeRemote,
			Target:  options.Image,
			Status:  status,
			Summary: summary,
			Details: details,
			Metadata: map[string]any{
				"context_kind":              "vex-net",
				"raw_count":                 netResponse.RawCount,
				"resolved_by_vex_count":     netResponse.ResolvedByVEXCount,
				"actionable_count":          netResponse.ActionableCount,
				"under_investigation_count": netResponse.UnderInvestigationCount,
				"threshold_breached":        netResponse.ThresholdBreached,
				"severity_threshold":        netResponse.SeverityThreshold,
			},
		}}
	}

	matches := []string{}
	seen := map[string]struct{}{}
	for _, finding := range findings {
		if _, ok := seen[finding.CVEID]; ok {
			continue
		}
		seen[finding.CVEID] = struct{}{}
		exceptions, err := client.ListExceptions(ctx, audit.ExceptionFilter{
			Status:      audit.ExceptionStatusApproved,
			CVEID:       finding.CVEID,
			ImageDigest: digest,
			TenantID:    options.Tenant,
			Environment: options.Environment,
			Repo:        options.Repository,
			Limit:       5,
		})
		if err != nil {
			return []CheckResult{{
				Name:    "remote-scan-context",
				Mode:    ModeRemote,
				Target:  options.Image,
				Status:  StatusError,
				Summary: err.Error(),
			}}
		}
		for _, exception := range exceptions {
			matches = append(matches, fmt.Sprintf("%s -> %s", finding.CVEID, exception.ExceptionID))
		}
	}

	summary := "no approved CVE exceptions found in ChangeLock"
	if len(matches) > 0 {
		summary = fmt.Sprintf("found %d approved CVE exception match(es) in ChangeLock", len(matches))
	}
	return []CheckResult{{
		Name:    "remote-scan-context",
		Mode:    ModeRemote,
		Target:  options.Image,
		Status:  StatusSkip,
		Summary: "net actionable VEX context skipped because image is not digest-pinned; approved CVE exception matches are partial context only",
		Details: matches,
		Metadata: map[string]any{
			"context_kind":          "vex-net-partial",
			"exception_match_count": len(matches),
			"partial_summary":       summary,
		},
	}}
}

func newFlagSet(name string) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs
}

type stringSlice []string

func (s *stringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSlice) Set(value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return errors.New("value cannot be empty")
	}
	*s = append(*s, value)
	return nil
}

func (s *stringSlice) Get() any {
	return []string(*s)
}

func collectManifestFiles(files, dirs []string) ([]string, error) {
	seen := map[string]struct{}{}
	results := []string{}
	addFile := func(path string) error {
		info, err := os.Stat(path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fmt.Errorf("%s is a directory; use --dir instead", path)
		}
		if !isYAML(path) {
			return nil
		}
		absolute, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		if _, ok := seen[absolute]; ok {
			return nil
		}
		seen[absolute] = struct{}{}
		results = append(results, absolute)
		return nil
	}
	for _, file := range files {
		if err := addFile(file); err != nil {
			return nil, err
		}
	}
	for _, dir := range dirs {
		if err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if !isYAML(path) {
				return nil
			}
			return addFile(path)
		}); err != nil {
			return nil, err
		}
	}
	sort.Strings(results)
	return results, nil
}

func isYAML(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return extension == ".yaml" || extension == ".yml"
}

func parseBool(value string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "0", "false", "no", "off":
		return false, nil
	case "1", "true", "yes", "on":
		return true, nil
	default:
		return false, fmt.Errorf("unsupported boolean %q", value)
	}
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
		return strings.ToUpper(strings.TrimSpace(value))
	}
}

func validSeverity(value string) bool {
	switch normalizeSeverity(value) {
	case "CRITICAL", "HIGH", "MEDIUM", "LOW":
		return true
	default:
		return false
	}
}

func findingsAtOrAbove(findings []ScanFinding, threshold string) []ScanFinding {
	thresholdRank := severityRank(threshold)
	filtered := make([]ScanFinding, 0, len(findings))
	for _, finding := range findings {
		if severityRank(finding.Severity) <= thresholdRank {
			filtered = append(filtered, finding)
		}
	}
	sort.Slice(filtered, func(i, j int) bool {
		if severityRank(filtered[i].Severity) != severityRank(filtered[j].Severity) {
			return severityRank(filtered[i].Severity) < severityRank(filtered[j].Severity)
		}
		return filtered[i].CVEID < filtered[j].CVEID
	})
	return filtered
}

func severityRank(value string) int {
	switch normalizeSeverity(value) {
	case "CRITICAL":
		return 0
	case "HIGH":
		return 1
	case "MEDIUM":
		return 2
	case "LOW":
		return 3
	default:
		return 4
	}
}

func summarizeFindings(findings []ScanFinding, limit int) []string {
	if len(findings) == 0 {
		return nil
	}
	if limit > 0 && len(findings) > limit {
		findings = findings[:limit]
	}
	lines := make([]string, 0, len(findings))
	for _, finding := range findings {
		lines = append(lines, fmt.Sprintf("%s %s %s %s", finding.Severity, finding.CVEID, strings.TrimSpace(finding.PackageName), strings.TrimSpace(finding.FixedVersion)))
	}
	return lines
}

func stripTagAndDigest(image string) string {
	image = strings.TrimSpace(image)
	if before, _, ok := strings.Cut(image, "@"); ok {
		image = before
	}
	lastSlash := strings.LastIndex(image, "/")
	lastColon := strings.LastIndex(image, ":")
	if lastColon > lastSlash {
		return image[:lastColon]
	}
	return image
}

func inferRepositoryFromImage(image string) string {
	ref := stripTagAndDigest(image)
	parts := strings.Split(ref, "/")
	if len(parts) >= 2 && (strings.Contains(parts[0], ".") || strings.Contains(parts[0], ":") || parts[0] == "localhost") {
		return strings.Join(parts[1:], "/")
	}
	return ref
}

func normalizeRepository(value string) string {
	return strings.Trim(strings.TrimSpace(value), "/")
}

func repoSubject(repository string) string {
	repository = normalizeRepository(repository)
	if repository == "" {
		return ""
	}
	return "repo:" + repository
}

func firstTenantRepository(bundle *policy.Bundle) string {
	if bundle == nil || len(bundle.Tenant.Spec.Repositories) == 0 {
		return ""
	}
	return bundle.Tenant.Spec.Repositories[0]
}

func verifiedIdentity(result *verify.ArtifactVerification) string {
	if result == nil {
		return ""
	}
	return result.VerifiedIdentity
}

func verifiedWorkflow(result *verify.ArtifactVerification) string {
	if result == nil {
		return ""
	}
	return result.VerifiedWorkflow
}

func verifiedSubject(result *verify.ArtifactVerification) string {
	if result == nil {
		return ""
	}
	return result.VerifiedSubject
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func firstNonEmptySlice(primary, fallback []string) []string {
	if len(primary) > 0 {
		return primary
	}
	return fallback
}

func statusForBool(condition bool, whenTrue, whenFalse Status) Status {
	if condition {
		return whenTrue
	}
	return whenFalse
}

func executionMode(offline bool, apiURL string) string {
	switch {
	case offline:
		return ModeOffline
	case strings.TrimSpace(apiURL) != "":
		return ModeAPIAssisted
	default:
		return ModeLocalOnly
	}
}

func summaryForBool(condition bool, whenTrue, whenFalse string) string {
	if condition {
		return whenTrue
	}
	return whenFalse
}

func (a *App) usage() string {
	return strings.TrimSpace(`
Usage:
  changelock-cli preflight [flags]
  changelock-cli check --config path
  changelock-cli preview --config path
  changelock-cli inspect --config path
  changelock-cli explain --config path --topic sync
  changelock-cli readiness --config path --profile production
  changelock-cli support --config path --profile production
  changelock-cli upgrade-readiness --config path --target-version 5.6.0
  changelock-cli phase5-summary --config path --profile production
  changelock-cli manifest [--file path | --dir path]
  changelock-cli image --image <ref>
  changelock-cli scan --image <ref>
  changelock-cli diagnostics --input result.json --format markdown
  changelock-cli guidance --input result.json --format markdown
  changelock-cli version

Exit codes:
  0 all executed checks passed
  1 one or more checks failed
  2 usage or input error
  3 execution or dependency error
`) + "\n"
}
