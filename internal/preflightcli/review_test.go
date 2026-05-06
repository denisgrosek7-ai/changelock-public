package preflightcli

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReviewProviderBlockingFindingFails(t *testing.T) {
	repoRoot := t.TempDir()
	writeReviewFile(t, repoRoot, "internal/formal/point14_valc.go", "package formal\n\nfunc demo() {}\n")
	writeReviewFile(t, repoRoot, "internal/formal/point14_valc_test.go", "package formal\n\nfunc TestDemo(t *testing.T) {}\n")
	diff := strings.TrimSpace(`
diff --git a/internal/formal/point14_valc.go b/internal/formal/point14_valc.go
index 1111111..2222222 100644
--- a/internal/formal/point14_valc.go
+++ b/internal/formal/point14_valc.go
@@ -933,0 +934,2 @@
+	if model.PublicPrivateBoundary == "tenant_private" {
+		return Point14ValCStateActive
+	}
`)

	app := newTestApp(t, Runtime{
		RunCommand: func(_ context.Context, name string, args ...string) (CommandExecution, error) {
			switch {
			case name == "git" && sameArgs(args, "rev-parse", "--show-toplevel"):
				return CommandExecution{Stdout: repoRoot + "\n", ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--name-only", "--diff-filter=ACMR"):
				return CommandExecution{Stdout: "internal/formal/point14_valc.go\ninternal/formal/point14_valc_test.go\n", ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--unified=0", "--no-ext-diff"):
				return CommandExecution{Stdout: diff + "\n", ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--check"):
				return CommandExecution{ExitCode: 0}, nil
			case name == "gofmt":
				return CommandExecution{ExitCode: 0}, nil
			case name == "fake-reviewer":
				if len(args) != 4 || args[0] != "--input" || args[2] != "--output" {
					t.Fatalf("unexpected provider args: %#v", args)
				}
				body, err := os.ReadFile(args[1])
				if err != nil {
					t.Fatalf("read provider request: %v", err)
				}
				var request ReviewProviderRequest
				if err := json.Unmarshal(body, &request); err != nil {
					t.Fatalf("decode provider request: %v", err)
				}
				if request.ReviewMode != "staged" || len(request.ChangedFiles) != 2 {
					t.Fatalf("unexpected provider request: %#v", request)
				}
				if len(request.Files) != 2 {
					t.Fatalf("expected full changed-file snapshots in provider request, got %#v", request.Files)
				}
				for _, file := range request.Files {
					if strings.TrimSpace(file.Content) == "" {
						t.Fatalf("expected provider request to include full changed-file content, got %#v", request.Files)
					}
				}
				response := ReviewProviderResponse{
					SchemaVersion: reviewResponseSchema,
					Findings: []ReviewProviderFinding{
						{
							FindingID: "finding-1",
							RuleID:    "formal.visibility_pair_consistency",
							Severity:  "P1",
							Summary:   "Visibility/public-private pair can become internally inconsistent",
							Detail:    "public_notice_bounded with tenant_private remains active in the diffed branch",
							File:      "internal/formal/point14_valc.go",
							StartLine: 934,
							EndLine:   935,
						},
					},
				}
				encoded, err := json.Marshal(response)
				if err != nil {
					t.Fatalf("marshal provider response: %v", err)
				}
				if err := os.WriteFile(args[3], encoded, 0o644); err != nil {
					t.Fatalf("write provider response: %v", err)
				}
				return CommandExecution{ExitCode: 0}, nil
			default:
				t.Fatalf("unexpected command %q args=%#v", name, args)
				return CommandExecution{}, nil
			}
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"review", "--staged", "--provider-bin", "fake-reviewer", "--output", "json"}, stdout, stderr)
	if code != ExitFailed {
		t.Fatalf("expected review to fail on P1 finding, got %d stderr=%q", code, stderr.String())
	}

	result := decodeJSONResult(t, stdout)
	if result.Command != "review" || result.OverallResult != StatusFail {
		t.Fatalf("unexpected review result: %#v", result)
	}
	foundFinding := false
	for _, check := range result.Checks {
		if check.Name == "review-finding" {
			foundFinding = true
			if check.Status != StatusFail {
				t.Fatalf("expected blocking review finding, got %#v", check)
			}
			if !strings.Contains(check.Summary, "P1") {
				t.Fatalf("expected severity in summary, got %#v", check)
			}
		}
	}
	if !foundFinding {
		t.Fatalf("expected review-finding check, got %#v", result.Checks)
	}
	if len(result.Diagnostics) == 0 || result.Diagnostics[len(result.Diagnostics)-1].TargetFile == "" || result.Diagnostics[len(result.Diagnostics)-1].Range == nil {
		t.Fatalf("expected provider finding diagnostics with file/range, got %#v", result.Diagnostics)
	}
}

func TestReviewBuiltInFormatFailureBlocksPush(t *testing.T) {
	repoRoot := t.TempDir()
	writeReviewFile(t, repoRoot, "internal/formal/point14_valc.go", "package formal\nfunc bad( ){}\n")
	diff := strings.TrimSpace(`
diff --git a/internal/formal/point14_valc.go b/internal/formal/point14_valc.go
index 1111111..2222222 100644
--- a/internal/formal/point14_valc.go
+++ b/internal/formal/point14_valc.go
@@ -1,0 +1,2 @@
+package formal
+func bad( ){}
`)

	app := newTestApp(t, Runtime{
		RunCommand: func(_ context.Context, name string, args ...string) (CommandExecution, error) {
			switch {
			case name == "git" && sameArgs(args, "rev-parse", "--show-toplevel"):
				return CommandExecution{Stdout: repoRoot + "\n", ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--name-only", "--diff-filter=ACMR"):
				return CommandExecution{Stdout: "internal/formal/point14_valc.go\n", ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--unified=0", "--no-ext-diff"):
				return CommandExecution{Stdout: diff + "\n", ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--check"):
				return CommandExecution{ExitCode: 0}, nil
			case name == "gofmt":
				return CommandExecution{Stdout: filepath.Join(repoRoot, "internal/formal/point14_valc.go") + "\n", ExitCode: 0}, nil
			default:
				t.Fatalf("unexpected command %q args=%#v", name, args)
				return CommandExecution{}, nil
			}
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"review", "--staged", "--output", "json"}, stdout, stderr)
	if code != ExitFailed {
		t.Fatalf("expected built-in review format failure, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.OverallResult != StatusFail {
		t.Fatalf("expected FAIL, got %#v", result)
	}
	foundFormat := false
	for _, check := range result.Checks {
		if check.Name == "review-format" {
			foundFormat = true
			if check.Status != StatusFail {
				t.Fatalf("expected review-format failure, got %#v", check)
			}
		}
	}
	if !foundFormat {
		t.Fatalf("expected review-format check, got %#v", result.Checks)
	}
}

func TestReviewNoChangedFilesReturnsInfo(t *testing.T) {
	repoRoot := t.TempDir()
	app := newTestApp(t, Runtime{
		RunCommand: func(_ context.Context, name string, args ...string) (CommandExecution, error) {
			switch {
			case name == "git" && sameArgs(args, "rev-parse", "--show-toplevel"):
				return CommandExecution{Stdout: repoRoot + "\n", ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--name-only", "--diff-filter=ACMR"):
				return CommandExecution{ExitCode: 0}, nil
			case name == "git" && sameArgs(args, "diff", "--cached", "--unified=0", "--no-ext-diff"):
				return CommandExecution{ExitCode: 0}, nil
			default:
				t.Fatalf("unexpected command %q args=%#v", name, args)
				return CommandExecution{}, nil
			}
		},
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"review", "--staged", "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success on empty review scope, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.OverallResult != StatusInfo {
		t.Fatalf("expected INFO result for empty scope, got %#v", result)
	}
}

func TestReviewDiffHeuristicChecksIgnoresDocumentationAndRuleStrings(t *testing.T) {
	repoRoot := t.TempDir()
	writeReviewFile(t, repoRoot, "docs/shift-left-integration.md", "- added `t.Skip(` markers\n- added `TODO` / `FIXME` markers\n")
	writeReviewFile(t, repoRoot, "internal/preflightcli/review.go", "package preflightcli\n\nfunc demo() {\n\t_ = []string{\"t.Skip(\", \"TODO\", \"FIXME\"}\n}\n")
	writeReviewFile(t, repoRoot, "internal/preflightcli/review_test.go", "package preflightcli\n\nfunc demoFixture() {\n\twriteReviewFile(nil, \"\", \"fixture\", \"t.Skip(\\\"debug\\\")\")\n}\n")
	app := &App{}
	scope := reviewScope{
		Repository: repoRoot,
		UnifiedDiff: strings.Join([]string{
			"diff --git a/docs/shift-left-integration.md b/docs/shift-left-integration.md",
			"index 1111111..2222222 100644",
			"--- a/docs/shift-left-integration.md",
			"+++ b/docs/shift-left-integration.md",
			"@@ -150,0 +151,2 @@",
			"+- added `t.Skip(` markers",
			"+- added `TODO` / `FIXME` markers",
			"diff --git a/internal/preflightcli/review.go b/internal/preflightcli/review.go",
			"index 3333333..4444444 100644",
			"--- a/internal/preflightcli/review.go",
			"+++ b/internal/preflightcli/review.go",
			"@@ -320,0 +321,2 @@",
			"+if details := reviewPatternDetails(lines, []string{\"t.Skip(\"}); len(details) > 0 {",
			"+if details := reviewPatternDetails(lines, []string{\"TODO\", \"FIXME\"}); len(details) > 0 {",
			"diff --git a/internal/preflightcli/review_test.go b/internal/preflightcli/review_test.go",
			"index 5555555..6666666 100644",
			"--- a/internal/preflightcli/review_test.go",
			"+++ b/internal/preflightcli/review_test.go",
			"@@ -1,0 +1,1 @@",
			"+\twriteReviewFile(nil, \"\", \"fixture\", \"t.Skip(\\\"debug\\\")\")",
		}, "\n"),
	}

	checks := app.reviewDiffHeuristicChecks(scope)
	if len(checks) == 0 {
		t.Fatalf("expected heuristic checks")
	}
	for _, check := range checks {
		switch check.Name {
		case "review-test-skip", "review-deferred-marker":
			if check.Status != StatusPass {
				t.Fatalf("expected %s to ignore documentation and rule strings, got %#v", check.Name, check)
			}
		}
	}
}

func TestReviewDiffHeuristicChecksFlagsRealSkipAndDeferredMarkers(t *testing.T) {
	repoRoot := t.TempDir()
	writeReviewFile(t, repoRoot, "internal/formal/point14_vale_test.go", "package formal\n\nimport \"testing\"\n\nfunc TestDemo(t *testing.T) {\n\t// TODO: tighten final authority coverage\n\tt.Skip(\"debugging\")\n}\n")
	writeReviewFile(t, repoRoot, "scripts/hooks/changelock-pre-commit.sh", "# FIXME: temporary hook marker\n")
	app := &App{}
	scope := reviewScope{
		Repository: repoRoot,
		UnifiedDiff: strings.TrimSpace(`
diff --git a/internal/formal/point14_vale_test.go b/internal/formal/point14_vale_test.go
index 1111111..2222222 100644
--- a/internal/formal/point14_vale_test.go
+++ b/internal/formal/point14_vale_test.go
@@ -10,0 +11,2 @@
+	// TODO: tighten final authority coverage
+	t.Skip("debugging")
diff --git a/scripts/hooks/changelock-pre-commit.sh b/scripts/hooks/changelock-pre-commit.sh
index 3333333..4444444 100644
--- a/scripts/hooks/changelock-pre-commit.sh
+++ b/scripts/hooks/changelock-pre-commit.sh
@@ -1,0 +1,1 @@
+# FIXME: temporary hook marker
`),
	}

	checks := app.reviewDiffHeuristicChecks(scope)
	if len(checks) == 0 {
		t.Fatalf("expected heuristic checks")
	}

	foundSkip := false
	foundDeferred := false
	for _, check := range checks {
		switch check.Name {
		case "review-test-skip":
			foundSkip = true
			if check.Status != StatusFail || len(check.Details) == 0 {
				t.Fatalf("expected real test skip to fail, got %#v", check)
			}
		case "review-deferred-marker":
			foundDeferred = true
			if check.Status != StatusWarning || len(check.Details) == 0 {
				t.Fatalf("expected real deferred marker to warn, got %#v", check)
			}
		}
	}
	if !foundSkip || !foundDeferred {
		t.Fatalf("expected both heuristic checks, got %#v", checks)
	}
}

func writeReviewFile(t *testing.T, repoRoot, relativePath, contents string) string {
	t.Helper()
	path := filepath.Join(repoRoot, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir review file: %v", err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write review file: %v", err)
	}
	return path
}

func sameArgs(got []string, want ...string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
