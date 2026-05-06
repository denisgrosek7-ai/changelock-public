#!/usr/bin/env node

import fs from 'node:fs';
import fsp from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';

const REQUEST_SCHEMA = '1.code_review_request.v1';
const RESPONSE_SCHEMA = '1.code_review_response.v1';
const DEFAULT_MODEL = 'gpt-5.4-mini';
const DEFAULT_API_URL = 'https://api.openai.com/v1/responses';

async function main() {
  const args = parseArgs(process.argv.slice(2));
  const request = JSON.parse(await fsp.readFile(args.input, 'utf8'));
  if (request.schema_version !== REQUEST_SCHEMA) {
    throw new Error(`unsupported request schema ${JSON.stringify(request.schema_version)}`);
  }

  loadDotEnvIfPresent(request.repository);

  const apiKey = firstNonEmpty(
    process.env.CHANGELOCK_REVIEW_OPENAI_API_KEY,
    process.env.OPENAI_API_KEY,
  );
  if (!apiKey) {
    throw new Error('missing CHANGELOCK_REVIEW_OPENAI_API_KEY or OPENAI_API_KEY');
  }

  const model = firstNonEmpty(process.env.CHANGELOCK_REVIEW_OPENAI_MODEL, DEFAULT_MODEL);
  const apiURL = firstNonEmpty(process.env.CHANGELOCK_REVIEW_OPENAI_BASE_URL, DEFAULT_API_URL);
  const prompt = buildPrompt(request);
  const payload = {
    model,
    input: [
      {
        role: 'system',
        content: [
          {
            type: 'input_text',
            text: systemPrompt(),
          },
        ],
      },
      {
        role: 'user',
        content: [
          {
            type: 'input_text',
            text: prompt,
          },
        ],
      },
    ],
    text: {
      format: {
        type: 'json_schema',
        name: 'code_review_response',
        strict: true,
        schema: responseSchema(),
      },
    },
  };

  const response = await fetch(apiURL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${apiKey}`,
    },
    body: JSON.stringify(payload),
  });

  const raw = await response.text();
  if (!response.ok) {
    throw new Error(`OpenAI review request failed (${response.status}): ${raw}`);
  }

  const body = JSON.parse(raw);
  const text = extractOutputText(body);
  if (!text) {
    throw new Error('OpenAI review response contained no text output');
  }

  const parsed = JSON.parse(text);
  const normalized = normalizeResponse(parsed);
  await fsp.writeFile(args.output, `${JSON.stringify(normalized, null, 2)}\n`, 'utf8');
}

function parseArgs(argv) {
  let input = '';
  let output = '';
  for (let index = 0; index < argv.length; index += 1) {
    const arg = argv[index];
    if (arg === '--input') {
      input = argv[index + 1] || '';
      index += 1;
      continue;
    }
    if (arg === '--output') {
      output = argv[index + 1] || '';
      index += 1;
      continue;
    }
  }
  if (!input || !output) {
    throw new Error('usage: changelock-openai-review.mjs --input request.json --output response.json');
  }
  return { input, output };
}

function loadDotEnvIfPresent(repository) {
  const envPath = path.join(repository || process.cwd(), '.env');
  try {
    const body = fs.readFileSync(envPath, 'utf8');
    for (const line of body.split(/\r?\n/)) {
      const trimmed = line.trim();
      if (!trimmed || trimmed.startsWith('#')) {
        continue;
      }
      const equals = trimmed.indexOf('=');
      if (equals <= 0) {
        continue;
      }
      const key = trimmed.slice(0, equals).trim();
      if (!key || process.env[key]) {
        continue;
      }
      let value = trimmed.slice(equals + 1).trim();
      if (
        (value.startsWith('"') && value.endsWith('"')) ||
        (value.startsWith("'") && value.endsWith("'"))
      ) {
        value = value.slice(1, -1);
      }
      process.env[key] = value;
    }
  } catch {
    // Optional local env only.
  }
}

function systemPrompt() {
  return [
    'You are a strict code review gate for ChangeLock.',
    'Review only the supplied diff and changed file snapshots.',
    'Return only real findings, not style nits.',
    'Focus on semantic defects, fail-closed regressions, hidden fallback-to-active paths, authority leaks, mutation leaks, tenant/privacy leaks, timestamp ordering bugs, test weakening, and mismatches between production logic and regression coverage.',
    'If there are no real findings, return an empty findings array.',
    'Do not propose fixes inside the findings JSON.',
  ].join(' ');
}

function buildPrompt(request) {
  const files = Array.isArray(request.files) ? request.files : [];
  const changedFiles = Array.isArray(request.changed_files) ? request.changed_files : [];
  const fileSections = files.map((file) => {
    const parts = [`FILE: ${file.path}`];
    if (file.content) {
      parts.push('CONTENT START');
      parts.push(file.content);
      parts.push('CONTENT END');
    }
    return parts.join('\n');
  });

  return [
    `Review mode: ${request.review_mode || 'unknown'}`,
    `Base ref: ${request.base_ref || 'none'}`,
    `Repository: ${request.repository || 'unknown'}`,
    `Block severity: ${request.block_severity || 'P2'}`,
    'Changed files:',
    changedFiles.join('\n'),
    '',
    'Unified diff:',
    request.unified_diff || '',
    '',
    'Changed file snapshots:',
    fileSections.join('\n\n'),
  ].join('\n');
}

function responseSchema() {
  return {
    type: 'object',
    additionalProperties: false,
    properties: {
      schema_version: { type: 'string' },
      findings: {
        type: 'array',
        items: {
          type: 'object',
          additionalProperties: false,
          properties: {
            finding_id: { type: 'string' },
            rule_id: { type: 'string' },
            severity: { type: 'string', enum: ['P0', 'P1', 'P2', 'P3'] },
            summary: { type: 'string' },
            detail: { type: 'string' },
            file: { type: 'string' },
            start_line: { type: 'integer', minimum: 1 },
            end_line: { type: 'integer', minimum: 1 },
          },
          required: ['severity', 'summary', 'file', 'start_line', 'end_line'],
        },
      },
    },
    required: ['schema_version', 'findings'],
  };
}

function extractOutputText(body) {
  if (typeof body.output_text === 'string' && body.output_text.trim()) {
    return body.output_text;
  }
  if (!Array.isArray(body.output)) {
    return '';
  }
  const texts = [];
  for (const item of body.output) {
    if (!Array.isArray(item.content)) {
      continue;
    }
    for (const content of item.content) {
      if (typeof content.text === 'string' && content.text.trim()) {
        texts.push(content.text);
      }
    }
  }
  return texts.join('\n').trim();
}

function normalizeResponse(response) {
  const findings = Array.isArray(response.findings) ? response.findings : [];
  return {
    schema_version: RESPONSE_SCHEMA,
    findings: findings.map((finding, index) => ({
      finding_id: firstNonEmpty(finding.finding_id, `finding-${index + 1}`),
      rule_id: firstNonEmpty(finding.rule_id, 'code_review_finding'),
      severity: normalizeSeverity(finding.severity),
      summary: String(finding.summary || '').trim(),
      detail: String(finding.detail || '').trim(),
      file: String(finding.file || '').trim(),
      start_line: positiveInteger(finding.start_line),
      end_line: positiveInteger(finding.end_line, positiveInteger(finding.start_line)),
    })).filter((finding) => finding.summary && finding.file && finding.start_line > 0 && finding.end_line > 0),
  };
}

function normalizeSeverity(value) {
  const severity = String(value || '').trim().toUpperCase();
  if (['P0', 'P1', 'P2', 'P3'].includes(severity)) {
    return severity;
  }
  return 'P2';
}

function positiveInteger(value, fallback = 1) {
  const number = Number.parseInt(String(value ?? ''), 10);
  if (Number.isInteger(number) && number > 0) {
    return number;
  }
  return fallback;
}

function firstNonEmpty(...values) {
  for (const value of values) {
    if (typeof value === 'string' && value.trim()) {
      return value.trim();
    }
  }
  return '';
}

main().catch((error) => {
  process.stderr.write(`${error.message}\n`);
  process.exitCode = 1;
});
