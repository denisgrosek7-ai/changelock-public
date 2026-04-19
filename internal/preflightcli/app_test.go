package preflightcli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/verify"
)

func TestUnknownCommandReturnsUsage(t *testing.T) {
	app := newTestApp(t, Runtime{})
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	code := app.Run(context.Background(), []string{"unknown"}, stdout, stderr)
	if code != ExitUsage {
		t.Fatalf("expected usage exit code, got %d", code)
	}
	if !strings.Contains(stderr.String(), "unknown command") {
		t.Fatalf("expected usage message, got %q", stderr.String())
	}
}

func TestPreflightRequiresInput(t *testing.T) {
	app := newTestApp(t, Runtime{})
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	code := app.Run(context.Background(), []string{"preflight"}, stdout, stderr)
	if code != ExitUsage {
		t.Fatalf("expected usage exit code, got %d", code)
	}
	if !strings.Contains(stderr.String(), "requires at least one manifest or image input") {
		t.Fatalf("unexpected usage output %q", stderr.String())
	}
}

func TestManifestSingleFileJSONOutput(t *testing.T) {
	manifestPath := writeTempYAML(t, "deployment.yaml")
	app := newTestApp(t, Runtime{
		RunCommand: func(_ context.Context, name string, args ...string) (CommandExecution, error) {
			if name != "kyverno" {
				t.Fatalf("unexpected binary %q", name)
			}
			return CommandExecution{Stdout: "policy ok", ExitCode: 0}, nil
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"manifest", "--file", manifestPath, "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}

	result := decodeJSONResult(t, stdout)
	if result.Command != "manifest" {
		t.Fatalf("unexpected command %q", result.Command)
	}
	if result.Mode != ModeLocalOnly {
		t.Fatalf("expected local-only mode, got %q", result.Mode)
	}
	if result.OverallResult != StatusPass {
		t.Fatalf("expected PASS, got %s", result.OverallResult)
	}
	if result.ExitCode != ExitSuccess {
		t.Fatalf("expected exit code 0, got %d", result.ExitCode)
	}
	if len(result.Checks) != 1 {
		t.Fatalf("expected one check, got %d", len(result.Checks))
	}
	if result.Checks[0].Mode != ModeLocal {
		t.Fatalf("expected local manifest check, got %q", result.Checks[0].Mode)
	}
	if result.Checks[0].Target != manifestPath {
		t.Fatalf("unexpected target %q", result.Checks[0].Target)
	}
}

func TestManifestMultiFileInput(t *testing.T) {
	dir := t.TempDir()
	first := writeYAMLFile(t, dir, "a.yaml")
	second := writeYAMLFile(t, dir, "nested/b.yaml")

	app := newTestApp(t, Runtime{
		RunCommand: func(_ context.Context, _ string, args ...string) (CommandExecution, error) {
			target := args[len(args)-1]
			return CommandExecution{Stdout: "ok " + filepath.Base(target), ExitCode: 0}, nil
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"manifest", "--file", second, "--file", first, "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}

	result := decodeJSONResult(t, stdout)
	if len(result.Checks) != 2 {
		t.Fatalf("expected two checks, got %d", len(result.Checks))
	}
	targets := []string{result.Checks[0].Target, result.Checks[1].Target}
	if !slices.Equal(targets, []string{first, second}) {
		t.Fatalf("unexpected target order: %+v", targets)
	}
}

func TestManifestDirectoryInput(t *testing.T) {
	dir := t.TempDir()
	first := writeYAMLFile(t, dir, "a.yaml")
	second := writeYAMLFile(t, dir, "nested/b.yml")
	if err := os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("ignore"), 0o644); err != nil {
		t.Fatalf("write notes.txt: %v", err)
	}

	app := newTestApp(t, Runtime{
		RunCommand: func(_ context.Context, _ string, _ ...string) (CommandExecution, error) {
			return CommandExecution{Stdout: "ok", ExitCode: 0}, nil
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"manifest", "--dir", dir, "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}

	result := decodeJSONResult(t, stdout)
	if len(result.Checks) != 2 {
		t.Fatalf("expected two YAML checks, got %d", len(result.Checks))
	}
	targets := []string{result.Checks[0].Target, result.Checks[1].Target}
	if !slices.Equal(targets, []string{first, second}) {
		t.Fatalf("unexpected targets %+v", targets)
	}
}

func TestPreflightAggregateExitCodes(t *testing.T) {
	manifestPath := writeTempYAML(t, "deployment.yaml")
	cases := []struct {
		name        string
		verifyErr   error
		scanSummary ScanSummary
		scanErr     error
		wantStatus  Status
		wantExit    int
	}{
		{
			name: "pass plus skip",
			scanSummary: ScanSummary{
				Scanner: "trivy",
				Image:   "ghcr.io/my-org/acme-app@sha256:1234",
				Counts:  map[string]int{"CRITICAL": 0, "HIGH": 0},
			},
			wantStatus: StatusPass,
			wantExit:   ExitSuccess,
		},
		{
			name: "fail plus skip",
			scanSummary: ScanSummary{
				Scanner: "trivy",
				Image:   "ghcr.io/my-org/acme-app@sha256:1234",
				Counts:  map[string]int{"CRITICAL": 1, "HIGH": 0},
				Findings: []ScanFinding{
					{CVEID: "CVE-2026-0001", Severity: "CRITICAL", PackageName: "openssl"},
				},
			},
			wantStatus: StatusFail,
			wantExit:   ExitFailed,
		},
		{
			name:      "error plus skip",
			verifyErr: fmt.Errorf("%w: cosign", ErrCommandNotFound),
			scanSummary: ScanSummary{
				Scanner: "trivy",
				Image:   "ghcr.io/my-org/acme-app@sha256:1234",
				Counts:  map[string]int{"CRITICAL": 0, "HIGH": 0},
			},
			wantStatus: StatusError,
			wantExit:   ExitExecution,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			app := newTestApp(t, Runtime{
				RunCommand: func(_ context.Context, _ string, _ ...string) (CommandExecution, error) {
					return CommandExecution{Stdout: "ok", ExitCode: 0}, nil
				},
				VerifyArtifact: func(_ context.Context, _ string, _ verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
					if tc.verifyErr != nil {
						return verify.ArtifactVerification{}, tc.verifyErr
					}
					return verify.ArtifactVerification{
						SignatureValid:   true,
						AttestationValid: true,
						VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
						VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
						VerifiedSubject:  "repo:my-org/acme-app",
						VerifiedRepo:     "my-org/acme-app",
					}, nil
				},
				ScanImage: func(_ context.Context, _ Config, _ string) (ScanSummary, error) {
					if tc.scanErr != nil {
						return ScanSummary{}, tc.scanErr
					}
					return tc.scanSummary, nil
				},
			})

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			code := app.Run(context.Background(), []string{
				"preflight",
				"--file", manifestPath,
				"--image", "ghcr.io/my-org/acme-app@sha256:1234",
				"--offline",
				"--output", "json",
			}, stdout, stderr)
			if code != tc.wantExit {
				t.Fatalf("expected exit %d, got %d stderr=%q", tc.wantExit, code, stderr.String())
			}

			result := decodeJSONResult(t, stdout)
			if result.OverallResult != tc.wantStatus {
				t.Fatalf("expected overall %s, got %s", tc.wantStatus, result.OverallResult)
			}
			if result.ExitCode != tc.wantExit {
				t.Fatalf("expected JSON exit %d, got %d", tc.wantExit, result.ExitCode)
			}
		})
	}
}

func TestRenderJSONOutputShape(t *testing.T) {
	result := Result{
		Command: "preflight",
		Mode:    ModeOffline,
		Checks: []CheckResult{
			{Name: "manifest", Mode: ModeLocal, Status: StatusPass, Summary: "ok", Target: "/tmp/a.yaml"},
			{Name: "image-trust", Mode: ModeLocal, Status: StatusFail, Summary: "denied", Target: "img"},
			{Name: "remote-auth", Mode: ModeRemote, Status: StatusSkip, Summary: "offline"},
			{Name: "scan", Mode: ModeLocal, Status: StatusError, Summary: "scanner missing"},
		},
	}
	buffer := &bytes.Buffer{}
	if err := renderResult(buffer, "json", result); err != nil {
		t.Fatalf("renderResult: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(buffer.Bytes(), &payload); err != nil {
		t.Fatalf("json unmarshal: %v", err)
	}
	for _, key := range []string{"command", "mode", "overall_result", "exit_code", "checks"} {
		if _, ok := payload[key]; !ok {
			t.Fatalf("expected key %q in payload: %+v", key, payload)
		}
	}
	for _, key := range []string{"diagnostics", "diagnostic_summary"} {
		if _, ok := payload[key]; !ok {
			t.Fatalf("expected key %q in payload: %+v", key, payload)
		}
	}
	checks, ok := payload["checks"].([]any)
	if !ok || len(checks) != 4 {
		t.Fatalf("unexpected checks payload: %#v", payload["checks"])
	}
	first, ok := checks[0].(map[string]any)
	if !ok {
		t.Fatalf("unexpected first check payload: %#v", checks[0])
	}
	for _, key := range []string{"name", "mode", "status", "summary", "target"} {
		if _, ok := first[key]; !ok {
			t.Fatalf("expected key %q in first check: %+v", key, first)
		}
	}
}

func TestOfflineModeSkipsRemoteChecks(t *testing.T) {
	app := newTestApp(t, Runtime{
		VerifyArtifact: func(_ context.Context, _ string, _ verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
			return verify.ArtifactVerification{
				SignatureValid:   true,
				AttestationValid: true,
			}, nil
		},
		ScanImage: func(_ context.Context, _ Config, image string) (ScanSummary, error) {
			return ScanSummary{Scanner: "trivy", Image: image, Counts: map[string]int{}}, nil
		},
	})

	result, err := app.runPreflight(context.Background(), []string{"--image", "ghcr.io/my-org/acme-app@sha256:abcd", "--offline"})
	if err != nil {
		t.Fatalf("runPreflight error: %v", err)
	}
	if result.Mode != ModeOffline {
		t.Fatalf("expected offline mode, got %q", result.Mode)
	}
	foundSkip := false
	for _, check := range result.Checks {
		if check.Mode == ModeRemote && check.Status == StatusSkip {
			foundSkip = true
		}
	}
	if !foundSkip {
		t.Fatalf("expected skipped remote check, got %+v", result.Checks)
	}
}

func TestConfiguredOnlineUnreachableAPIIsError(t *testing.T) {
	app := newTestApp(t, Runtime{
		VerifyArtifact: func(_ context.Context, _ string, _ verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
			return verify.ArtifactVerification{
				SignatureValid:   true,
				AttestationValid: true,
			}, nil
		},
		ScanImage: func(_ context.Context, _ Config, image string) (ScanSummary, error) {
			return ScanSummary{Scanner: "trivy", Image: image, Counts: map[string]int{}}, nil
		},
		HTTPClient: &http.Client{
			Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
				return nil, errors.New("dial tcp 127.0.0.1:8094: connect: connection refused")
			}),
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{
		"preflight",
		"--image", "ghcr.io/my-org/acme-app@sha256:abcd",
		"--api-url", "https://changelock.example",
		"--output", "json",
	}, stdout, stderr)
	if code != ExitExecution {
		t.Fatalf("expected execution exit code, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.Mode != ModeAPIAssisted {
		t.Fatalf("expected api-assisted mode, got %q", result.Mode)
	}
	foundError := false
	for _, check := range result.Checks {
		if check.Name == "remote-auth" && check.Mode == ModeRemote && check.Status == StatusError {
			foundError = true
		}
	}
	if !foundError {
		t.Fatalf("expected remote-auth error, got %+v", result.Checks)
	}
}

func TestManifestMissingDependencyReturnsExecutionError(t *testing.T) {
	manifestPath := writeTempYAML(t, "deployment.yaml")
	app := newTestApp(t, Runtime{
		RunCommand: func(_ context.Context, _ string, _ ...string) (CommandExecution, error) {
			return CommandExecution{}, fmt.Errorf("%w: kyverno", ErrCommandNotFound)
		},
	})
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"manifest", "--file", manifestPath, "--output", "json"}, stdout, stderr)
	if code != ExitExecution {
		t.Fatalf("expected execution exit code, got %d", code)
	}
	result := decodeJSONResult(t, stdout)
	if result.OverallResult != StatusError {
		t.Fatalf("expected ERROR, got %s", result.OverallResult)
	}
	if result.Checks[0].Status != StatusError {
		t.Fatalf("expected manifest check ERROR, got %s", result.Checks[0].Status)
	}
}

func TestImageMissingDependencyReturnsExecutionError(t *testing.T) {
	app := newTestApp(t, Runtime{
		VerifyArtifact: func(_ context.Context, _ string, _ verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
			return verify.ArtifactVerification{}, fmt.Errorf("%w: cosign", ErrCommandNotFound)
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{
		"image",
		"--image", "ghcr.io/my-org/acme-app@sha256:abcd",
		"--repository", "my-org/acme-app",
		"--output", "json",
	}, stdout, stderr)
	if code != ExitExecution {
		t.Fatalf("expected execution exit code, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.OverallResult != StatusError {
		t.Fatalf("expected ERROR, got %s", result.OverallResult)
	}
}

func TestScanMissingDependencyReturnsExecutionError(t *testing.T) {
	app := newTestApp(t, Runtime{
		ScanImage: func(_ context.Context, _ Config, _ string) (ScanSummary, error) {
			return ScanSummary{}, fmt.Errorf("%w: no vulnerability scanner found in PATH", ErrCommandNotFound)
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"scan", "--image", "ghcr.io/my-org/acme-app@sha256:abcd", "--output", "json"}, stdout, stderr)
	if code != ExitExecution {
		t.Fatalf("expected execution exit code, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.OverallResult != StatusError {
		t.Fatalf("expected ERROR, got %s", result.OverallResult)
	}
}

func TestRemoteExceptionLookupUsesAuthHeader(t *testing.T) {
	var authHeaders []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaders = append(authHeaders, r.Header.Get("Authorization"))
		switch r.URL.Path {
		case "/v1/exceptions":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"exceptions":[{"id":1,"exception_id":"EX-1","exception_type":"DIGEST_BYPASS","status":"APPROVED","reason":"demo","ticket_id":"INC-1","created_at":"2026-01-01T00:00:00Z","expires_at":"2026-01-02T00:00:00Z","active":true}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	app := newTestApp(t, Runtime{
		VerifyArtifact: func(_ context.Context, _ string, _ verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
			return verify.ArtifactVerification{
				SignatureValid:   true,
				AttestationValid: true,
				VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
				VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
				VerifiedSubject:  "repo:my-org/acme-app",
				VerifiedRepo:     "my-org/acme-app",
			}, nil
		},
		HTTPClient: server.Client(),
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{
		"image",
		"--image", "ghcr.io/my-org/acme-app@sha256:abcd",
		"--repository", "my-org/acme-app",
		"--api-url", server.URL,
		"--token", "secret-token",
		"--output", "json",
	}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}
	if len(authHeaders) == 0 || authHeaders[0] != "Bearer secret-token" {
		t.Fatalf("expected bearer auth header, got %+v", authHeaders)
	}
	result := decodeJSONResult(t, stdout)
	foundRemote := false
	for _, check := range result.Checks {
		if check.Name == "remote-image-context" && check.Status == StatusPass {
			foundRemote = true
			if len(check.Details) != 1 || !strings.Contains(check.Details[0], "EX-1") {
				t.Fatalf("unexpected remote details %+v", check.Details)
			}
		}
	}
	if !foundRemote {
		t.Fatalf("expected remote-image-context check, got %+v", result.Checks)
	}
}

func TestScanThresholdFailure(t *testing.T) {
	app := newTestApp(t, Runtime{
		ScanImage: func(_ context.Context, _ Config, image string) (ScanSummary, error) {
			return ScanSummary{
				Scanner: "trivy",
				Image:   image,
				Counts:  map[string]int{"CRITICAL": 1, "HIGH": 1},
				Findings: []ScanFinding{
					{CVEID: "CVE-2026-0001", Severity: "CRITICAL", PackageName: "openssl"},
				},
			}, nil
		},
	})
	result, err := app.runScan(context.Background(), []string{"--image", "ghcr.io/my-org/acme-app@sha256:abcd", "--offline"})
	if err != nil {
		t.Fatalf("runScan error: %v", err)
	}
	if result.OverallResult != StatusFail {
		t.Fatalf("expected FAIL, got %s", result.OverallResult)
	}
}

func TestNoChecksRunIsExecutionError(t *testing.T) {
	result := finalizeResult(Result{Command: "preflight", Mode: ModeLocalOnly})
	if result.OverallResult != StatusError {
		t.Fatalf("expected ERROR overall result, got %s", result.OverallResult)
	}
	if result.ExitCode != ExitExecution {
		t.Fatalf("expected execution exit code, got %d", result.ExitCode)
	}
}

func newTestApp(t *testing.T, runtime Runtime) *App {
	t.Helper()
	if runtime.RunCommand == nil {
		runtime.RunCommand = func(_ context.Context, _ string, _ ...string) (CommandExecution, error) {
			return CommandExecution{ExitCode: 0}, nil
		}
	}
	if runtime.VerifyArtifact == nil {
		runtime.VerifyArtifact = func(_ context.Context, _ string, _ verify.ArtifactVerificationRequest) (verify.ArtifactVerification, error) {
			return verify.ArtifactVerification{}, nil
		}
	}
	if runtime.ScanImage == nil {
		runtime.ScanImage = func(_ context.Context, _ Config, image string) (ScanSummary, error) {
			return ScanSummary{Scanner: "trivy", Image: image, Counts: map[string]int{}}, nil
		}
	}
	if runtime.HTTPClient == nil {
		runtime.HTTPClient = http.DefaultClient
	}

	app, err := NewApp(func(key string) string {
		switch key {
		case "CHANGELOCK_CLI_OUTPUT":
			return "human"
		case "CHANGELOCK_CLI_TIMEOUT":
			return "10s"
		case "CHANGELOCK_CLI_POLICY_DIR":
			return filepath.Join("..", "..", "policies")
		case "CHANGELOCK_CLI_KYVERNO_POLICY_DIR":
			return filepath.Join("..", "..", "deploy", "kyverno")
		case "CHANGELOCK_CLI_SCANNER":
			return "auto"
		default:
			return ""
		}
	}, runtime)
	if err != nil {
		t.Fatalf("NewApp error: %v", err)
	}
	return app
}

func decodeJSONResult(t *testing.T, buffer *bytes.Buffer) Result {
	t.Helper()
	var result Result
	if err := json.Unmarshal(buffer.Bytes(), &result); err != nil {
		t.Fatalf("unmarshal json result: %v body=%q", err, buffer.String())
	}
	return result
}

func writeTempYAML(t *testing.T, name string) string {
	t.Helper()
	return writeYAMLFile(t, t.TempDir(), name)
}

func writeYAMLFile(t *testing.T, dir, relativePath string) string {
	t.Helper()
	path := filepath.Join(dir, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdirs for yaml file: %v", err)
	}
	if err := os.WriteFile(path, []byte("apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: demo\n"), 0o644); err != nil {
		t.Fatalf("write yaml file: %v", err)
	}
	absolute, err := filepath.Abs(path)
	if err != nil {
		t.Fatalf("abs path: %v", err)
	}
	return absolute
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}
