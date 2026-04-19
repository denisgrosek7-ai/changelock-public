const { execFile } = require("child_process");
const fs = require("fs");
const os = require("os");
const path = require("path");

function getVSCode() {
  return require("vscode");
}

function isYAMLDocument(document) {
  if (!document || typeof document.fileName !== "string") {
    return false;
  }
  return document.fileName.endsWith(".yaml") || document.fileName.endsWith(".yml") || document.languageId === "yaml";
}

function extractDigestPinnedImage(text) {
  const expression = /^\s*image:\s*["']?([^"'#\s]+)["']?\s*$/gm;
  let match;
  while ((match = expression.exec(text)) !== null) {
    const value = String(match[1] || "").trim();
    if (value.includes("@sha256:")) {
      return value;
    }
  }
  return "";
}

function resolveWorkspacePath(basePath, configuredPath) {
  if (!configuredPath) {
    return "";
  }
  if (path.isAbsolute(configuredPath)) {
    return configuredPath;
  }
  return path.join(basePath, configuredPath);
}

function buildCLIArgs(options) {
  const args = ["preflight", "--output", "json"];
  for (const file of options.files || []) {
    args.push("--file", file);
  }
  if (options.policyDir) {
    args.push("--policy-dir", options.policyDir);
  }
  if (options.bundleDir) {
    args.push("--bundle-dir", options.bundleDir);
  }
  if (options.scanner) {
    args.push("--scanner", options.scanner);
  }
  if (options.tenant) {
    args.push("--tenant", options.tenant);
  }
  if (options.repository) {
    args.push("--repository", options.repository);
  }
  if (options.offline) {
    args.push("--offline");
  } else if (options.apiUrl) {
    args.push("--api-url", options.apiUrl);
  }
  if (options.image) {
    args.push("--image", options.image);
  }
  return args;
}

function buildGuidanceArgs(resultPath) {
  return ["guidance", "--input", resultPath, "--format", "markdown"];
}

function normalizeDiagnostics(result) {
  if (!result || !Array.isArray(result.diagnostics)) {
    return [];
  }
  return result.diagnostics.filter((item) => item && typeof item === "object");
}

function diagnosticText(issue) {
  const lines = [];
  if (issue.summary) {
    lines.push(issue.summary);
  } else if (issue.message) {
    lines.push(issue.message);
  }
  if (issue.fix_hint) {
    lines.push(`Fix: ${issue.fix_hint}`);
  }
  if (issue.docs_ref) {
    lines.push(`Docs: ${issue.docs_ref}`);
  }
  if (issue.evaluation_state && issue.evaluation_state !== "pass") {
    lines.push(`State: ${issue.evaluation_state}`);
  }
  return lines.join("\n");
}

function diagnosticRange(issue, vscode) {
  const range = issue.range || {};
  const startLine = Math.max(Number(range.start_line || 1) - 1, 0);
  const startColumn = Math.max(Number(range.start_column || 1) - 1, 0);
  const endLine = Math.max(Number(range.end_line || range.start_line || 1) - 1, startLine);
  const endColumn = Math.max(Number(range.end_column || range.start_column || 1) - 1, startColumn);
  return new vscode.Range(startLine, startColumn, endLine, endColumn);
}

function diagnosticSeverity(issue, vscode) {
  switch (issue.severity) {
    case "error":
      return vscode.DiagnosticSeverity.Error;
    case "warning":
      return vscode.DiagnosticSeverity.Warning;
    default:
      return vscode.DiagnosticSeverity.Information;
  }
}

function toVSCodeDiagnostic(issue, vscode) {
  const diagnostic = new vscode.Diagnostic(diagnosticRange(issue, vscode), diagnosticText(issue), diagnosticSeverity(issue, vscode));
  diagnostic.code = issue.reason_code || issue.rule_id || issue.check_id;
  diagnostic.source = "ChangeLock";
  diagnostic.changelock = issue;
  return diagnostic;
}

function groupByWorkspace(documents, vscode) {
  const grouped = new Map();
  for (const document of documents) {
    const folder = vscode.workspace.getWorkspaceFolder(document.uri);
    const key = folder ? folder.uri.fsPath : path.dirname(document.fileName);
    if (!grouped.has(key)) {
      grouped.set(key, { folder, documents: [] });
    }
    grouped.get(key).documents.push(document);
  }
  return grouped;
}

function readSettings(vscode, scopeUri) {
  const config = vscode.workspace.getConfiguration("changelock", scopeUri);
  return {
    cliPath: config.get("cliPath", "changelock-cli"),
    apiUrl: config.get("apiUrl", ""),
    token: config.get("token", ""),
    tokenEnvVar: config.get("tokenEnvVar", "CHANGELOCK_CLI_TOKEN"),
    offline: config.get("offline", false),
    policyDir: config.get("policyDir", "deploy/kyverno"),
    bundleDir: config.get("bundleDir", "policies"),
    scanner: config.get("scanner", "auto"),
    tenant: config.get("tenant", "acme"),
    repository: config.get("repository", ""),
    runOnSave: config.get("runOnSave", false),
  };
}

function execCLI(cliPath, args, cwd, env) {
  return new Promise((resolve, reject) => {
    execFile(cliPath, args, { cwd, env }, (error, stdout, stderr) => {
      if (error && !stdout) {
        reject(new Error(stderr || error.message));
        return;
      }
      try {
        const parsed = JSON.parse(stdout);
        resolve(parsed);
      } catch (parseError) {
        reject(new Error(stderr || parseError.message));
      }
    });
  });
}

function execCLIText(cliPath, args, cwd, env) {
  return new Promise((resolve, reject) => {
    execFile(cliPath, args, { cwd, env }, (error, stdout, stderr) => {
      if (error) {
        reject(new Error(stderr || error.message));
        return;
      }
      resolve(stdout);
    });
  });
}

function writeGuidanceInput(result) {
  const dir = fs.mkdtempSync(path.join(os.tmpdir(), "changelock-guidance-"));
  const file = path.join(dir, "result.json");
  fs.writeFileSync(file, JSON.stringify(result), "utf8");
  return { dir, file };
}

function docsPathForIssue(rootPath, issue) {
  if (!issue || !issue.docs_ref) {
    return null;
  }
  const docsPath = path.resolve(rootPath, issue.docs_ref);
  if (!fs.existsSync(docsPath)) {
    return null;
  }
  return docsPath;
}

async function openDocs(documentUri, docsRef) {
  const vscode = getVSCode();
  const folder = vscode.workspace.getWorkspaceFolder(documentUri);
  const rootPath = folder ? folder.uri.fsPath : path.dirname(documentUri.fsPath);
  const docsPath = docsPathForIssue(rootPath, { docs_ref: docsRef });
  if (!docsPath) {
    vscode.window.showWarningMessage(`ChangeLock docs path not found: ${docsRef}`);
    return;
  }
  await vscode.window.showTextDocument(vscode.Uri.file(docsPath));
}

async function runChecksForDocuments(collection, issueCache, resultCache, documents, explicit) {
  const vscode = getVSCode();
  const supported = documents.filter(isYAMLDocument).filter((document) => !document.isUntitled);
  if (!supported.length) {
    if (explicit) {
      vscode.window.showInformationMessage("ChangeLock checks currently run on saved YAML manifests only.");
    }
    return;
  }
  for (const document of supported) {
    if (document.isDirty) {
      if (explicit) {
        vscode.window.showWarningMessage(`Save ${path.basename(document.fileName)} before running ChangeLock checks.`);
      }
      return;
    }
  }

  const grouped = groupByWorkspace(supported, vscode);
  let totalBlocking = 0;
  let totalAdvisory = 0;

  for (const [, group] of grouped) {
    const scopeUri = group.folder ? group.folder.uri : supported[0].uri;
    const settings = readSettings(vscode, scopeUri);
    const rootPath = group.folder ? group.folder.uri.fsPath : path.dirname(group.documents[0].fileName);
    const firstImage = group.documents.length === 1 ? extractDigestPinnedImage(group.documents[0].getText()) : "";
    const args = buildCLIArgs({
      files: group.documents.map((document) => document.fileName),
      policyDir: resolveWorkspacePath(rootPath, settings.policyDir),
      bundleDir: resolveWorkspacePath(rootPath, settings.bundleDir),
      scanner: settings.scanner,
      tenant: settings.tenant,
      repository: settings.repository,
      apiUrl: settings.apiUrl,
      offline: settings.offline,
      image: firstImage,
    });
    const env = { ...process.env };
    const token = settings.token || process.env[settings.tokenEnvVar] || "";
    if (token) {
      env.CHANGELOCK_CLI_TOKEN = token;
    }

    let result;
    try {
      result = await execCLI(settings.cliPath, args, rootPath, env);
    } catch (error) {
      const message = error instanceof Error ? error.message : "ChangeLock CLI execution failed.";
      if (explicit) {
        vscode.window.showErrorMessage(message);
      }
      return;
    }
    resultCache.set(rootPath, result);

    const diagnosticsByFile = new Map();
    const diagnostics = normalizeDiagnostics(result);
    for (const issue of diagnostics) {
      const targetFile = issue.target_file;
      if (!targetFile) {
        continue;
      }
      if (!diagnosticsByFile.has(targetFile)) {
        diagnosticsByFile.set(targetFile, []);
      }
      diagnosticsByFile.get(targetFile).push(issue);
      if (issue.blocking) {
        totalBlocking += 1;
      } else if (issue.evaluation_state !== "pass") {
        totalAdvisory += 1;
      }
    }
    for (const document of group.documents) {
      const fileDiagnostics = (diagnosticsByFile.get(document.fileName) || []).map((issue) => toVSCodeDiagnostic(issue, vscode));
      collection.set(document.uri, fileDiagnostics);
      issueCache.set(document.uri.toString(), diagnosticsByFile.get(document.fileName) || []);
    }
  }

  vscode.window.setStatusBarMessage(`ChangeLock: ${totalBlocking} blocking, ${totalAdvisory} advisory diagnostic(s)`, 4000);
}

class ChangeLockCodeActionProvider {
  provideCodeActions(document, _range, context) {
    const vscode = getVSCode();
    const actions = [];
    const seenDocs = new Set();
    for (const diagnostic of context.diagnostics || []) {
      const issue = diagnostic.changelock;
      if (!issue) {
        continue;
      }
      if (issue.docs_ref && !seenDocs.has(issue.docs_ref)) {
        seenDocs.add(issue.docs_ref);
        const action = new vscode.CodeAction("Open ChangeLock docs", vscode.CodeActionKind.QuickFix);
        action.command = {
          command: "changelock.openDocs",
          title: "Open ChangeLock docs",
          arguments: [document.uri, issue.docs_ref],
        };
        action.diagnostics = [diagnostic];
        actions.push(action);
      }
    }
    const rerun = new vscode.CodeAction("Rerun ChangeLock checks", vscode.CodeActionKind.QuickFix);
    rerun.command = {
      command: "changelock.runCurrentFileChecks",
      title: "Rerun ChangeLock checks",
    };
    rerun.diagnostics = context.diagnostics;
    actions.push(rerun);
    const guidance = new vscode.CodeAction("Show ChangeLock guidance", vscode.CodeActionKind.QuickFix);
    guidance.command = {
      command: "changelock.showCurrentFileGuidance",
      title: "Show ChangeLock guidance",
    };
    guidance.diagnostics = context.diagnostics;
    actions.push(guidance);
    return actions;
  }
}

class ChangeLockHoverProvider {
  constructor(issueCache) {
    this.issueCache = issueCache;
  }

  provideHover(document, position) {
    const vscode = getVSCode();
    const issues = this.issueCache.get(document.uri.toString()) || [];
    const matching = issues.filter((issue) => {
      const range = issue.range || {};
      const startLine = Math.max(Number(range.start_line || 1) - 1, 0);
      const endLine = Math.max(Number(range.end_line || range.start_line || 1) - 1, startLine);
      return position.line >= startLine && position.line <= endLine;
    });
    if (!matching.length) {
      return undefined;
    }
    const markdown = new vscode.MarkdownString();
    markdown.isTrusted = false;
    for (const issue of matching) {
      markdown.appendMarkdown(`**${issue.reason_code || issue.rule_id || issue.check_id}**\n\n`);
      markdown.appendText(diagnosticText(issue));
      markdown.appendMarkdown("\n\n");
    }
    return new vscode.Hover(markdown);
  }
}

function activate(context) {
  const vscode = getVSCode();
  const collection = vscode.languages.createDiagnosticCollection("changelock");
  const issueCache = new Map();
  const resultCache = new Map();
  const guidanceOutput = vscode.window.createOutputChannel("ChangeLock Guidance");
  context.subscriptions.push(collection);
  context.subscriptions.push(guidanceOutput);

  context.subscriptions.push(vscode.commands.registerCommand("changelock.runCurrentFileChecks", async () => {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
      vscode.window.showInformationMessage("Open a YAML manifest to run ChangeLock checks.");
      return;
    }
    await runChecksForDocuments(collection, issueCache, resultCache, [editor.document], true);
  }));

  context.subscriptions.push(vscode.commands.registerCommand("changelock.runWorkspaceChecks", async () => {
    const documents = await vscode.workspace.findFiles("**/*.{yaml,yml}", "**/{.git,node_modules}/**");
    const openable = documents.map((uri) => ({ uri, fileName: uri.fsPath, languageId: "yaml", isUntitled: false, isDirty: false, getText: () => fs.readFileSync(uri.fsPath, "utf8") }));
    await runChecksForDocuments(collection, issueCache, resultCache, openable, true);
  }));

  context.subscriptions.push(vscode.commands.registerCommand("changelock.showCurrentFileGuidance", async () => {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
      vscode.window.showInformationMessage("Open a YAML manifest to show ChangeLock guidance.");
      return;
    }
    await runChecksForDocuments(collection, issueCache, resultCache, [editor.document], true);
    const folder = vscode.workspace.getWorkspaceFolder(editor.document.uri);
    const rootPath = folder ? folder.uri.fsPath : path.dirname(editor.document.fileName);
    const result = resultCache.get(rootPath);
    if (!result) {
      vscode.window.showWarningMessage("ChangeLock guidance is unavailable because no preflight result was produced.");
      return;
    }
    const settings = readSettings(vscode, editor.document.uri);
    const env = { ...process.env };
    const token = settings.token || process.env[settings.tokenEnvVar] || "";
    if (token) {
      env.CHANGELOCK_CLI_TOKEN = token;
    }
    const temp = writeGuidanceInput(result);
    try {
      const markdown = await execCLIText(settings.cliPath, buildGuidanceArgs(temp.file), rootPath, env);
      guidanceOutput.clear();
      guidanceOutput.append(markdown.trim() || "No contextual guidance was generated.");
      guidanceOutput.show(true);
    } catch (error) {
      const message = error instanceof Error ? error.message : "ChangeLock guidance rendering failed.";
      vscode.window.showErrorMessage(message);
    } finally {
      fs.rmSync(temp.dir, { recursive: true, force: true });
    }
  }));

  context.subscriptions.push(vscode.commands.registerCommand("changelock.clearDiagnostics", () => {
    collection.clear();
    issueCache.clear();
    resultCache.clear();
    vscode.window.setStatusBarMessage("ChangeLock diagnostics cleared.", 3000);
  }));

  context.subscriptions.push(vscode.commands.registerCommand("changelock.openDocs", openDocs));
  context.subscriptions.push(vscode.languages.registerCodeActionsProvider({ language: "yaml" }, new ChangeLockCodeActionProvider()));
  context.subscriptions.push(vscode.languages.registerHoverProvider({ language: "yaml" }, new ChangeLockHoverProvider(issueCache)));
  context.subscriptions.push(vscode.workspace.onDidSaveTextDocument(async (document) => {
    if (!isYAMLDocument(document)) {
      return;
    }
    const settings = readSettings(vscode, document.uri);
    if (!settings.runOnSave) {
      return;
    }
    await runChecksForDocuments(collection, issueCache, resultCache, [document], false);
  }));
}

function deactivate() {}

module.exports = {
  activate,
  deactivate,
  buildCLIArgs,
  diagnosticText,
  extractDigestPinnedImage,
  buildGuidanceArgs,
  normalizeDiagnostics,
  resolveWorkspacePath,
};
