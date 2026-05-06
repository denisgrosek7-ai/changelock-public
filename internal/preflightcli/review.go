package preflightcli

import (
	"context"
	"encoding/json"
	"fmt"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	reviewRequestSchema  = "1.code_review_request.v1"
	reviewResponseSchema = "1.code_review_response.v1"
)

type reviewOptions struct {
	Staged        bool
	UpstreamRef   string
	Output        string
	ProviderBin   string
	BlockSeverity string
}

type reviewScope struct {
	Mode         string
	BaseRef      string
	Repository   string
	ChangedFiles []string
	UnifiedDiff  string
}

type ReviewProviderRequest struct {
	SchemaVersion string               `json:"schema_version"`
	ReviewMode    string               `json:"review_mode"`
	BaseRef       string               `json:"base_ref,omitempty"`
	Repository    string               `json:"repository"`
	BlockSeverity string               `json:"block_severity"`
	ChangedFiles  []string             `json:"changed_files"`
	UnifiedDiff   string               `json:"unified_diff"`
	Files         []ReviewProviderFile `json:"files,omitempty"`
	Limitations   []string             `json:"limitations,omitempty"`
	Metadata      map[string]string    `json:"metadata,omitempty"`
}

type ReviewProviderFile struct {
	Path         string `json:"path"`
	AbsolutePath string `json:"absolute_path,omitempty"`
	Content      string `json:"content,omitempty"`
}

type ReviewProviderResponse struct {
	SchemaVersion string                  `json:"schema_version"`
	Findings      []ReviewProviderFinding `json:"findings,omitempty"`
}

type ReviewProviderFinding struct {
	FindingID string `json:"finding_id,omitempty"`
	RuleID    string `json:"rule_id,omitempty"`
	Severity  string `json:"severity"`
	Summary   string `json:"summary"`
	Detail    string `json:"detail,omitempty"`
	File      string `json:"file,omitempty"`
	StartLine int    `json:"start_line,omitempty"`
	EndLine   int    `json:"end_line,omitempty"`
}

type reviewAddedLine struct {
	File string
	Line int
	Text string
}

type reviewGoLineAnnotations struct {
	stringLines         map[int]struct{}
	deferredCommentLine map[int]struct{}
}

func (a *App) parseReviewOptions(args []string) (reviewOptions, error) {
	fs := newFlagSet("review")
	options := reviewOptions{
		Output:        a.config.Output,
		ProviderBin:   a.config.ReviewProviderBin,
		BlockSeverity: normalizeReviewSeverity(firstNonEmpty(a.config.ReviewBlockSeverity, "P2")),
	}
	fs.BoolVar(&options.Staged, "staged", false, "review staged changes from git index")
	fs.StringVar(&options.UpstreamRef, "upstream-ref", "", "review git diff against <ref>...HEAD")
	fs.StringVar(&options.ProviderBin, "provider-bin", options.ProviderBin, "optional external review provider binary")
	fs.StringVar(&options.BlockSeverity, "block-severity", options.BlockSeverity, "minimum provider severity that blocks: P0|P1|P2|P3")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	if err := fs.Parse(args); err != nil {
		return reviewOptions{}, usageError{message: err.Error()}
	}
	if options.Staged && strings.TrimSpace(options.UpstreamRef) != "" {
		return reviewOptions{}, usageError{message: "review accepts either --staged or --upstream-ref, not both"}
	}
	if !options.Staged && strings.TrimSpace(options.UpstreamRef) == "" {
		options.Staged = true
	}
	options.ProviderBin = strings.TrimSpace(options.ProviderBin)
	options.BlockSeverity = normalizeReviewSeverity(options.BlockSeverity)
	if !validReviewSeverity(options.BlockSeverity) {
		return reviewOptions{}, usageError{message: fmt.Sprintf("unsupported review block severity %q", options.BlockSeverity)}
	}
	return options, nil
}

func (a *App) runReview(ctx context.Context, args []string) (Result, error) {
	options, err := a.parseReviewOptions(args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.Output

	repoRoot, err := a.gitRepoRoot(ctx)
	if err != nil {
		return Result{}, err
	}
	scope, err := a.collectReviewScope(ctx, repoRoot, options)
	if err != nil {
		return Result{}, err
	}

	result := Result{
		Command: "review",
		Mode:    ModeLocalOnly,
		Inputs: map[string]string{
			"scope":      scope.Mode,
			"base_ref":   scope.BaseRef,
			"repo_root":  repoRoot,
			"provider":   options.ProviderBin,
			"block_from": options.BlockSeverity,
		},
		Limitations: []string{
			"Review gate is read-only and bounded to the selected git diff scope.",
			"Runtime agents and canonical production decision paths remain unchanged.",
		},
	}
	if len(scope.ChangedFiles) == 0 {
		result.add(CheckResult{
			Name:    "review-scope",
			Mode:    ModeLocal,
			Status:  StatusInfo,
			Summary: "no changed files matched the requested review scope",
		})
		return finalizeResult(result), nil
	}

	result.add(CheckResult{
		Name:    "review-scope",
		Mode:    ModeLocal,
		Status:  StatusPass,
		Summary: fmt.Sprintf("reviewing %d changed file(s) in %s scope", len(scope.ChangedFiles), scope.Mode),
		Details: truncateReviewDetails(scope.ChangedFiles, 12),
		Metadata: map[string]any{
			"context_kind": "code_review_scope",
		},
	})

	builtInChecks, err := a.localReviewChecks(ctx, scope)
	if err != nil {
		return Result{}, err
	}
	result.add(builtInChecks...)

	providerChecks, err := a.providerReviewChecks(ctx, scope, options)
	if err != nil {
		result.add(CheckResult{
			Name:    "review-provider",
			Mode:    ModeLocal,
			Status:  StatusError,
			Summary: err.Error(),
		})
		return finalizeResult(result), nil
	}
	result.add(providerChecks...)

	return finalizeResult(result), nil
}

func (a *App) gitRepoRoot(ctx context.Context) (string, error) {
	result, err := a.runtime.RunCommand(ctx, "git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", err
	}
	if result.ExitCode != 0 {
		return "", usageError{message: firstNonEmpty(strings.TrimSpace(result.Stderr), "unable to determine git repository root")}
	}
	root := strings.TrimSpace(result.Stdout)
	if root == "" {
		return "", usageError{message: "git repository root is empty"}
	}
	return root, nil
}

func (a *App) collectReviewScope(ctx context.Context, repoRoot string, options reviewOptions) (reviewScope, error) {
	scope := reviewScope{
		Repository: repoRoot,
	}
	var fileArgs []string
	var diffArgs []string
	if options.Staged {
		scope.Mode = "staged"
		scope.BaseRef = "INDEX"
		fileArgs = []string{"diff", "--cached", "--name-only", "--diff-filter=ACMR"}
		diffArgs = []string{"diff", "--cached", "--unified=0", "--no-ext-diff"}
	} else {
		scope.Mode = "upstream"
		scope.BaseRef = strings.TrimSpace(options.UpstreamRef)
		rangeRef := scope.BaseRef + "...HEAD"
		fileArgs = []string{"diff", "--name-only", "--diff-filter=ACMR", rangeRef}
		diffArgs = []string{"diff", "--unified=0", "--no-ext-diff", rangeRef}
	}

	fileResult, err := a.runtime.RunCommand(ctx, "git", fileArgs...)
	if err != nil {
		return reviewScope{}, err
	}
	if fileResult.ExitCode != 0 {
		return reviewScope{}, usageError{message: firstNonEmpty(strings.TrimSpace(fileResult.Stderr), "unable to collect review scope")}
	}
	scope.ChangedFiles = collectReviewFiles(fileResult.Stdout)

	diffResult, err := a.runtime.RunCommand(ctx, "git", diffArgs...)
	if err != nil {
		return reviewScope{}, err
	}
	if diffResult.ExitCode != 0 {
		return reviewScope{}, usageError{message: firstNonEmpty(strings.TrimSpace(diffResult.Stderr), "unable to collect review diff")}
	}
	scope.UnifiedDiff = diffResult.Stdout
	return scope, nil
}

func collectReviewFiles(raw string) []string {
	files := []string{}
	seen := map[string]struct{}{}
	for _, line := range strings.Split(raw, "\n") {
		file := strings.TrimSpace(line)
		if file == "" {
			continue
		}
		if _, ok := seen[file]; ok {
			continue
		}
		seen[file] = struct{}{}
		files = append(files, file)
	}
	sort.Strings(files)
	return files
}

func (a *App) localReviewChecks(ctx context.Context, scope reviewScope) ([]CheckResult, error) {
	checks := []CheckResult{a.reviewDiffIntegrityCheck(ctx, scope)}
	goFiles := filterReviewFilesByExt(scope.Repository, scope.ChangedFiles, ".go")
	if len(goFiles) == 0 {
		checks = append(checks, CheckResult{
			Name:    "review-format",
			Mode:    ModeLocal,
			Status:  StatusSkip,
			Summary: "no changed Go files required local formatting review",
		})
		return append(checks, a.reviewDiffHeuristicChecks(scope)...), nil
	}
	checks = append(checks, a.reviewGoFormatCheck(ctx, goFiles))
	checks = append(checks, a.reviewDiffHeuristicChecks(scope)...)
	return checks, nil
}

func (a *App) reviewDiffIntegrityCheck(ctx context.Context, scope reviewScope) CheckResult {
	args := []string{"diff"}
	if scope.Mode == "staged" {
		args = append(args, "--cached", "--check")
	} else {
		args = append(args, "--check", scope.BaseRef+"...HEAD")
	}
	result, err := a.runtime.RunCommand(ctx, "git", args...)
	check := CheckResult{
		Name:    "review-diff",
		Mode:    ModeLocal,
		Status:  StatusPass,
		Summary: "git diff --check found no whitespace or merge marker issues in the review scope",
	}
	if err != nil {
		check.Status = StatusError
		check.Summary = err.Error()
		return check
	}
	if result.ExitCode != 0 {
		check.Status = StatusFail
		check.Summary = "git diff --check reported whitespace or merge marker issues"
		check.Details = truncateLines(firstNonEmpty(result.Stdout, result.Stderr), 12)
	}
	return check
}

func (a *App) reviewGoFormatCheck(ctx context.Context, files []string) CheckResult {
	args := append([]string{"-l"}, files...)
	result, err := a.runtime.RunCommand(ctx, "gofmt", args...)
	check := CheckResult{
		Name:    "review-format",
		Mode:    ModeLocal,
		Status:  StatusPass,
		Summary: "gofmt formatting matches the changed Go files",
	}
	if err != nil {
		check.Status = StatusError
		check.Summary = err.Error()
		return check
	}
	unformatted := collectReviewFiles(result.Stdout)
	if len(unformatted) > 0 {
		check.Status = StatusFail
		check.Summary = "changed Go files require gofmt before push"
		check.Details = truncateReviewDetails(unformatted, 12)
	}
	return check
}

func (a *App) reviewDiffHeuristicChecks(scope reviewScope) []CheckResult {
	lines := parseReviewAddedLines(scope.UnifiedDiff)
	checks := []CheckResult{}

	if details := reviewTestSkipDetails(scope.Repository, lines); len(details) > 0 {
		checks = append(checks, CheckResult{
			Name:    "review-test-skip",
			Mode:    ModeLocal,
			Status:  StatusFail,
			Summary: "added test skip markers require review before push",
			Details: truncateReviewDetails(details, 12),
		})
	} else {
		checks = append(checks, CheckResult{
			Name:    "review-test-skip",
			Mode:    ModeLocal,
			Status:  StatusPass,
			Summary: "no added test skip markers were detected in the review diff",
		})
	}

	if details := reviewDeferredMarkerDetails(scope.Repository, lines); len(details) > 0 {
		checks = append(checks, CheckResult{
			Name:    "review-deferred-marker",
			Mode:    ModeLocal,
			Status:  StatusWarning,
			Summary: "added deferred markers were detected in the review diff",
			Details: truncateReviewDetails(details, 12),
		})
	} else {
		checks = append(checks, CheckResult{
			Name:    "review-deferred-marker",
			Mode:    ModeLocal,
			Status:  StatusPass,
			Summary: "no added TODO or FIXME markers were detected in the review diff",
		})
	}

	if details := formalPointFilesWithoutChangedTests(scope.ChangedFiles); len(details) > 0 {
		checks = append(checks, CheckResult{
			Name:    "review-formal-test-coverage",
			Mode:    ModeLocal,
			Status:  StatusWarning,
			Summary: "formal point production files changed without matching test-file changes",
			Details: truncateReviewDetails(details, 12),
		})
	} else {
		checks = append(checks, CheckResult{
			Name:    "review-formal-test-coverage",
			Mode:    ModeLocal,
			Status:  StatusPass,
			Summary: "formal point production-file changes were accompanied by matching test-file changes when present",
		})
	}

	return checks
}

func formalPointFilesWithoutChangedTests(files []string) []string {
	changed := map[string]struct{}{}
	for _, file := range files {
		changed[file] = struct{}{}
	}
	missing := []string{}
	for _, file := range files {
		if !strings.HasPrefix(file, "internal/formal/point") || strings.HasSuffix(file, "_test.go") || filepath.Ext(file) != ".go" {
			continue
		}
		testFile := strings.TrimSuffix(file, ".go") + "_test.go"
		if _, ok := changed[testFile]; !ok {
			missing = append(missing, fmt.Sprintf("%s -> expected related test delta in %s", file, testFile))
		}
	}
	sort.Strings(missing)
	return missing
}

func (a *App) providerReviewChecks(ctx context.Context, scope reviewScope, options reviewOptions) ([]CheckResult, error) {
	if strings.TrimSpace(options.ProviderBin) == "" {
		return []CheckResult{{
			Name:    "review-provider",
			Mode:    ModeLocal,
			Status:  StatusSkip,
			Summary: "no external review provider configured; deterministic local review only",
		}}, nil
	}

	request := ReviewProviderRequest{
		SchemaVersion: reviewRequestSchema,
		ReviewMode:    scope.Mode,
		BaseRef:       scope.BaseRef,
		Repository:    scope.Repository,
		BlockSeverity: options.BlockSeverity,
		ChangedFiles:  append([]string{}, scope.ChangedFiles...),
		UnifiedDiff:   scope.UnifiedDiff,
		Files:         reviewProviderFiles(scope.Repository, scope.ChangedFiles),
		Limitations: []string{
			"Provider is read-only and must not mutate repository state.",
			"Provider findings are bounded to the supplied diff and file snapshots.",
		},
		Metadata: map[string]string{
			"review_command": "changelock-cli review",
		},
	}

	tempDir, err := os.MkdirTemp("", "changelock-review-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	inputPath := filepath.Join(tempDir, "request.json")
	outputPath := filepath.Join(tempDir, "response.json")
	body, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(inputPath, body, 0o600); err != nil {
		return nil, err
	}

	result, err := a.runtime.RunCommand(ctx, options.ProviderBin, "--input", inputPath, "--output", outputPath)
	if err != nil {
		return nil, err
	}
	if result.ExitCode != 0 {
		return nil, fmt.Errorf("review provider exited with status %d: %s", result.ExitCode, firstNonEmpty(strings.TrimSpace(result.Stderr), strings.TrimSpace(result.Stdout)))
	}

	responseBody, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, err
	}
	var response ReviewProviderResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("decode review provider response: %w", err)
	}
	if len(response.Findings) == 0 {
		return []CheckResult{{
			Name:    "review-provider",
			Mode:    ModeLocal,
			Status:  StatusPass,
			Summary: "external review provider returned no findings in the requested scope",
		}}, nil
	}

	sort.SliceStable(response.Findings, func(i, j int) bool {
		if reviewSeverityRank(response.Findings[i].Severity) != reviewSeverityRank(response.Findings[j].Severity) {
			return reviewSeverityRank(response.Findings[i].Severity) < reviewSeverityRank(response.Findings[j].Severity)
		}
		if response.Findings[i].File != response.Findings[j].File {
			return response.Findings[i].File < response.Findings[j].File
		}
		return response.Findings[i].StartLine < response.Findings[j].StartLine
	})

	checks := make([]CheckResult, 0, len(response.Findings))
	for _, finding := range response.Findings {
		severity := normalizeReviewSeverity(finding.Severity)
		if !validReviewSeverity(severity) {
			severity = "P2"
		}
		targetFile := normalizeReviewFilePath(scope.Repository, finding.File)
		status := StatusWarning
		if reviewSeverityBlocks(severity, options.BlockSeverity) {
			status = StatusFail
		}
		details := []string{}
		if text := strings.TrimSpace(finding.Detail); text != "" {
			details = append(details, text)
		}
		if targetFile != "" && finding.StartLine > 0 {
			details = append(details, fmt.Sprintf("%s:%d", targetFile, finding.StartLine))
		}
		checks = append(checks, CheckResult{
			Name:    "review-finding",
			Mode:    ModeLocal,
			Status:  status,
			Target:  targetFile,
			Summary: fmt.Sprintf("%s %s", severity, firstNonEmpty(strings.TrimSpace(finding.Summary), strings.TrimSpace(finding.RuleID))),
			Details: details,
			Metadata: map[string]any{
				"finding_severity": severity,
				"target_file":      targetFile,
				"start_line":       finding.StartLine,
				"end_line":         maxInt(finding.EndLine, finding.StartLine),
				"rule_id":          strings.TrimSpace(finding.RuleID),
				"provider":         options.ProviderBin,
			},
		})
	}
	return checks, nil
}

func reviewProviderFiles(repoRoot string, files []string) []ReviewProviderFile {
	snapshots := make([]ReviewProviderFile, 0, len(files))
	for _, file := range files {
		absolute := normalizeReviewFilePath(repoRoot, file)
		snapshot := ReviewProviderFile{
			Path:         file,
			AbsolutePath: absolute,
		}
		if shouldEmbedReviewFile(file) {
			if body, err := os.ReadFile(absolute); err == nil {
				snapshot.Content = string(body)
			}
		}
		snapshots = append(snapshots, snapshot)
	}
	return snapshots
}

func shouldEmbedReviewFile(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".go", ".py", ".js", ".jsx", ".ts", ".tsx", ".sh", ".md", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func filterReviewFilesByExt(repoRoot string, files []string, extension string) []string {
	filtered := []string{}
	for _, file := range files {
		if strings.EqualFold(filepath.Ext(file), extension) {
			filtered = append(filtered, normalizeReviewFilePath(repoRoot, file))
		}
	}
	sort.Strings(filtered)
	return filtered
}

func normalizeReviewFilePath(repoRoot, path string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return ""
	}
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	if strings.TrimSpace(repoRoot) == "" {
		return filepath.Clean(path)
	}
	return filepath.Join(repoRoot, filepath.Clean(path))
}

func parseReviewAddedLines(diff string) []reviewAddedLine {
	lines := []reviewAddedLine{}
	currentFile := ""
	currentLine := 0
	for _, raw := range strings.Split(diff, "\n") {
		switch {
		case strings.HasPrefix(raw, "+++ b/"):
			currentFile = strings.TrimPrefix(raw, "+++ b/")
		case strings.HasPrefix(raw, "+++ /dev/null"):
			currentFile = ""
		case strings.HasPrefix(raw, "@@"):
			currentLine = parseReviewHunkStart(raw)
		case strings.HasPrefix(raw, "+") && !strings.HasPrefix(raw, "+++"):
			if currentFile != "" && currentLine > 0 {
				lines = append(lines, reviewAddedLine{
					File: currentFile,
					Line: currentLine,
					Text: strings.TrimPrefix(raw, "+"),
				})
			}
			if currentLine > 0 {
				currentLine++
			}
		case strings.HasPrefix(raw, "-") && !strings.HasPrefix(raw, "---"):
		case strings.HasPrefix(raw, " "):
			if currentLine > 0 {
				currentLine++
			}
		}
	}
	return lines
}

func parseReviewHunkStart(raw string) int {
	plus := strings.Index(raw, "+")
	if plus == -1 {
		return 0
	}
	rest := raw[plus+1:]
	end := strings.Index(rest, " ")
	if end == -1 {
		end = len(rest)
	}
	chunk := rest[:end]
	startText, _, _ := strings.Cut(chunk, ",")
	start, err := strconv.Atoi(strings.TrimSpace(startText))
	if err != nil || start <= 0 {
		return 0
	}
	return start
}

func reviewPatternDetails(lines []reviewAddedLine, patterns []string) []string {
	details := []string{}
	for _, line := range lines {
		normalized := strings.ToLower(line.Text)
		for _, pattern := range patterns {
			if strings.Contains(normalized, strings.ToLower(pattern)) {
				details = append(details, fmt.Sprintf("%s:%d %s", line.File, line.Line, strings.TrimSpace(line.Text)))
				break
			}
		}
	}
	return details
}

func reviewTestSkipDetails(repoRoot string, lines []reviewAddedLine) []string {
	testLines := []reviewAddedLine{}
	for _, line := range lines {
		if strings.HasSuffix(line.File, "_test.go") && !reviewGoLineInString(repoRoot, line) && reviewExecutableSkipLine(line.Text) {
			testLines = append(testLines, line)
		}
	}
	return reviewPatternDetails(testLines, []string{"t.Skip("})
}

func reviewDeferredMarkerDetails(repoRoot string, lines []reviewAddedLine) []string {
	details := []string{}
	for _, line := range lines {
		if !reviewDeferredMarkerLine(repoRoot, line) {
			continue
		}
		details = append(details, fmt.Sprintf("%s:%d %s", line.File, line.Line, strings.TrimSpace(line.Text)))
	}
	return details
}

func reviewDeferredMarkerLine(repoRoot string, line reviewAddedLine) bool {
	if strings.HasSuffix(line.File, ".go") {
		return reviewGoLineHasDeferredComment(repoRoot, line)
	}
	return reviewCommentDeferredMarker(line.Text)
}

func reviewCommentDeferredMarker(text string) bool {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return false
	}
	upper := strings.ToUpper(trimmed)
	if !strings.Contains(upper, "TODO") && !strings.Contains(upper, "FIXME") {
		return false
	}
	commentPrefixes := []string{"//", "#", "/*", "*", "<!--", "--", ";"}
	for _, prefix := range commentPrefixes {
		if strings.HasPrefix(trimmed, prefix) {
			return true
		}
	}
	return strings.HasPrefix(upper, "TODO:") || strings.HasPrefix(upper, "FIXME:")
}

func reviewExecutableSkipLine(text string) bool {
	trimmed := strings.TrimSpace(text)
	index := strings.Index(trimmed, "t.Skip(")
	if index == -1 {
		return false
	}
	prefix := trimmed[:index]
	return !strings.ContainsAny(prefix, "\"'`") && !strings.Contains(prefix, "//")
}

func reviewGoLineInString(repoRoot string, line reviewAddedLine) bool {
	annotations := reviewGoFileAnnotations(repoRoot, line.File)
	if annotations == nil {
		return false
	}
	_, ok := annotations.stringLines[line.Line]
	return ok
}

func reviewGoLineHasDeferredComment(repoRoot string, line reviewAddedLine) bool {
	annotations := reviewGoFileAnnotations(repoRoot, line.File)
	if annotations == nil {
		return false
	}
	_, ok := annotations.deferredCommentLine[line.Line]
	return ok
}

func reviewGoFileAnnotations(repoRoot, file string) *reviewGoLineAnnotations {
	absolute := normalizeReviewFilePath(repoRoot, file)
	body, err := os.ReadFile(absolute)
	if err != nil {
		return nil
	}
	fset := token.NewFileSet()
	tokenFile := fset.AddFile(absolute, -1, len(body))
	var s scanner.Scanner
	s.Init(tokenFile, body, nil, scanner.ScanComments)

	annotations := &reviewGoLineAnnotations{
		stringLines:         map[int]struct{}{},
		deferredCommentLine: map[int]struct{}{},
	}
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			return annotations
		}
		position := fset.Position(pos)
		startLine := position.Line
		endLine := startLine + strings.Count(lit, "\n")
		switch tok {
		case token.STRING, token.CHAR:
			for lineNumber := startLine; lineNumber <= endLine; lineNumber++ {
				annotations.stringLines[lineNumber] = struct{}{}
			}
		case token.COMMENT:
			if !reviewCommentDeferredMarker(lit) {
				continue
			}
			for lineNumber := startLine; lineNumber <= endLine; lineNumber++ {
				annotations.deferredCommentLine[lineNumber] = struct{}{}
			}
		}
	}
}

func truncateReviewDetails(values []string, limit int) []string {
	if len(values) == 0 {
		return nil
	}
	if limit > 0 && len(values) > limit {
		values = append(append([]string{}, values[:limit]...), "...output truncated...")
	}
	return append([]string{}, values...)
}

func normalizeReviewSeverity(value string) string {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "P0":
		return "P0"
	case "P1":
		return "P1"
	case "P2":
		return "P2"
	case "P3":
		return "P3"
	default:
		return strings.ToUpper(strings.TrimSpace(value))
	}
}

func validReviewSeverity(value string) bool {
	switch normalizeReviewSeverity(value) {
	case "P0", "P1", "P2", "P3":
		return true
	default:
		return false
	}
}

func reviewSeverityRank(value string) int {
	switch normalizeReviewSeverity(value) {
	case "P0":
		return 0
	case "P1":
		return 1
	case "P2":
		return 2
	case "P3":
		return 3
	default:
		return 99
	}
}

func reviewSeverityBlocks(findingSeverity, threshold string) bool {
	return reviewSeverityRank(findingSeverity) <= reviewSeverityRank(threshold)
}

func maxInt(left, right int) int {
	if left > right {
		return left
	}
	return right
}
