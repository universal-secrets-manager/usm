A professional, production-grade plan for building a universal, encrypted secrets system that replaces .env and unifies local dev, CI/CD, and cloud sync. Structured as Cursor-ready tasks to generate the repo file-by-file.

0) High-Level Architecture

Monorepo layout (single Git repo):

cli/usm — Go CLI (single binary).

core/spec — Format specs, threat model, ADRs.

core/crypto — Go crypto library shared by CLI and Go SDK.

sdks/node — TypeScript SDK.

sdks/python — Python SDK.

sdks/php — PHP SDK.

sdks/go — Go SDK (thin wrapper over core/crypto).

integrations/github-action — Action to inject secrets.

integrations/gitlab — GitLab template & helper script.

integrations/jenkins — Jenkins plugin wrapper script.

examples/* — Minimal apps consuming USM via each SDK.

docs — Docusaurus site (or MkDocs if preferred).

Secrets file: .secrets.yml (default); supports profiles (dev, staging, prod) and recipient-based envelope encryption.

Key material: .usmkey (project file key), .usm/keys/* (per-user keys), .usm/team.yml (team recipients).

Crypto: AES‑256‑GCM for data; X25519 (ECDH) + HKDF for envelope keys; scrypt for passphrase KDF; Ed25519 for signing.

Distribution: GoReleaser for multi-OS binaries; npm/PyPI/Packagist modules; GitHub Action marketplace.

1) Repository Scaffolding (Cursor Task Block)

Task 1.1 — Create repo root files

README.md — Product overview; quickstart; security notes.

LICENSE — MIT.

.gitignore — Go, Node, Python, PHP, build artifacts.

.gitattributes — Ensure LF normalization, linguist hints.

.editorconfig — Consistent whitespace.

CODEOWNERS — Assign maintainers.

CONTRIBUTING.md — Dev setup, commit style, DCO/CLA.

SECURITY.md — Vulnerability disclosure policy.

Makefile — Top-level build/test/format targets.

.pre-commit-config.yaml — Format/lint hooks across langs.

.github/workflows/ci.yml — Matrix build & tests (Go/Node/Python/PHP).

.github/ISSUE_TEMPLATE/* & .github/PULL_REQUEST_TEMPLATE.md.

Task 1.2 — Monorepo directories

Create: cli/usm, core/spec, core/crypto, sdks/{node,python,php,go}, integrations/{github-action,gitlab,jenkins}, examples/{node,python,php,go}, docs.

2) Format Specification & Threat Model

Spec goals

Human-diffable YAML for structure; values encrypted.

Support multiple environments (profiles) and scopes (app, service, ci).

Envelope encryption: a random File Key (FK) encrypts each secret value; FK encrypted to Recipients (team/users/CI) via X25519 + HKDF.

Auth: Ed25519 signature covers the entire file to detect tampering.

File: core/spec/USM-Format-v1.md

Header:

usm: 1
kdf:
  name: scrypt
  n: 32768
  r: 8
  p: 1
  salt: <base64>
crypto:
  cipher: aes-256-gcm
  nonce_bytes: 12
recipients:
  - type: x25519
    id: team:engineering
    pub: <base64>
  - type: x25519
    id: user:alice@company.com
    pub: <base64>
profiles: [dev, staging, prod]
sign:
  alg: ed25519
  keyid: repo:default
  sig: <base64>

Secrets payload (by profile):

secrets:
  dev:
    DB_URL: ENC{
      fk_enc: <base64>
      aad: <base64>            # AAD binds key name+profile+path
      nonce: <base64>
      tag: <base64>
      ct: <base64>
    }
    API_KEY: ENC{...}
  prod: { ... }
metadata:
  created: 2025-08-20T00:00:00Z
  updated: 2025-08-20T00:00:00Z
  version: 1

Deterministic AAD: ${repo_path}:${profile}:${key}.

Key derivation: Passphrase → scrypt → MasterKey; MasterKey unwraps Project File Key (PFK). PFK encrypts FKs; FKs encrypt values.

Rotation: rotate value → new FK; rotate PFK → rewrap FKs; rotate recipients → re-encrypt FK for new pubkeys.

Audit: append-only log (hash-chained) in .usm/audit.log.

File: core/spec/Threat-Model.md

Assets: secrets at rest, keys at rest, keys in memory, CI envs.

Adversaries: repo reader, stolen laptop, malicious insider, compromised CI.

Mitigations: GCM AEAD; zeroization; OS keyrings; masking in logs; short‑lived process env; SIGTERM traps.

Out of scope: running process memory scrape; fully compromised developer box.

File: core/spec/ADRs/*.md

ADR-0001: Use YAML for structure + encrypted values.

ADR-0002: X25519 + HKDF for envelope; Ed25519 for signatures.

ADR-0003: scrypt parameters and tunables.

3) Go Crypto Core (core/crypto)

Task 3.1 — Package skeleton

core/crypto/crypto.go: interfaces

Encrypt(value []byte, aad []byte) (ct, nonce, tag []byte, err error)

Decrypt(ct, nonce, tag, aad []byte) ([]byte, err)

Key structs: PFK, FK, Recipient, Signature.

aead/aesgcm.go: AES‑GCM implementation.

kdf/scrypt.go: scrypt; params struct; test vectors.

asym/x25519.go: ECDH; HKDF for wrapping FKs.

sign/ed25519.go: sign/verify.

securemem/zero.go: best‑effort zeroization helpers.

Task 3.2 — File I/O

file/format.go: marshal/unmarshal YAML, preserve ordering, stable formatting for reproducible sigs.

file/paths.go: locate .secrets.yml, .usm/ directory cross‑platform.

Task 3.3 — Tests

Unit tests for AEAD, KDF, round‑trip, invalid tags.

Fuzz tests on parse/decrypt paths.

4) CLI (cli/usm) — UX & Commands

CLI goals: zero‑friction, safe defaults, works offline, JSON output capable.

Task 4.1 — Main skeleton

Use spf13/cobra.

Commands: init, set, get, list, delete, rotate, share, recipients, profiles, sync (scaffold only in MVP), audit.

Task 4.2 — Command behaviors

init:

Generate .secrets.yml, .usmkey (PFK wrapped by scrypt from passphrase) and .usm/team.yml.

Offer --no-passphrase to store PFK in OS keyring (Keychain/DPAPI/libsecret).

set KEY=VALUE [--profile prod] [--stdin] [--from-file path].

get KEY [--profile prod] [--json] [--no-export] — prints to stdout; exit non‑zero if missing.

list [--profile dev] — keys only; --with-meta to include timestamps.

delete KEY [--profile ...] — tombstone with audit entry.

rotate [KEY|--pfk|--recipients] — value or key rotation.

share add user@org --pub <file|base64> — updates team.yml.

recipients list|add|rm — team/admin helpers.

profiles add staging — create new profile.

audit show --since 7d — verify hash chain & print events.

Task 4.3 — Safety

Mask secrets in logs; prevent --verbose from printing values.

Refuse to run on dirty git tree for destructive ops unless --force.

.env import: usm import-env .env --profile dev.

Task 4.4 — Non‑interactive mode

All commands accept flags only; no prompts in CI.

Task 4.5 — Build & release

goreleaser.yml for darwin/amd64+arm64, linux, windows.

5) SDKs

Common behavior

Auto‑discover .secrets.yml via CWD → parents.

Respect USM_PROFILE env var (default dev).

Allow override via USM_FILE path.

In‑memory cache with TTL; no on‑disk plaintext.

Task 5.1 — Node (TypeScript) sdks/node

package.json, tsconfig.json, src/index.ts.

API:

import { USM } from "@usm/secrets";
const usm = await USM.load();
const db = await usm.get("DB_URL");

Provide ESM+CJS builds, type defs.

Tests with Vitest.

Task 5.2 — Python sdks/python

pyproject.toml (PEP 621), src/usm/__init__.py.

API:

from usm import load
usm = load()
db = usm.get("DB_URL")

Tests with pytest.

Task 5.3 — PHP sdks/php

composer.json, src/USM.php.

PSR‑4 autoloading; unit tests via PHPUnit.

Task 5.4 — Go sdks/go

Thin wrapper over core/crypto + file loader.

6) CI/CD Integrations

Task 6.1 — GitHub Action integrations/github-action

action.yml (composite):

Inputs: profile, file, export (bool).

Steps: download usm release; run usm get --profile ... --json and export to ${{ env }}.

Mask values via ::add-mask::.

Example usage in README.md.

Task 6.2 — GitLab integrations/gitlab

usm.gitlab.yml include; helper script usm-export.sh.

Task 6.3 — Jenkins integrations/jenkins

usm.groovy shared library example and shell wrapper.

7) Examples

examples/node — express app reading DB_URL via SDK.

examples/python — flask app.

examples/php — slim app.

examples/go — small HTTP server.

Each example includes a Makefile target make run that ensures usm is present and exports the necessary env at runtime without writing plaintext files.

8) Documentation Site (docs)

Docusaurus with sections:

Quickstart (CLI)

Format Spec

Security Model

SDK Guides

CI/CD Recipes

Migration from .env

Enterprise features (preview)

Include copy‑paste snippets and security callouts.

9) Roadmap & Milestones

M1 — MVP (2 weeks): CLI (init/set/get/list), Node+Python SDKs, basic docs, examples.

M2 — Team & Rotation (2–3 weeks): share/recipients/rotate, audit log, Go+PHP SDKs.

M3 — CI/CD & Releases (1–2 weeks): GitHub Action, GoReleaser, signed releases.

M4 — Cloud Sync (3–4 weeks): AWS/GCP/Vault providers (opt‑in).

M5 — Enterprise Preview: policy engine, SSO, RBAC, approval flows.

10) Security Baselines (Do Not Skip)

Secrets never touch disk in plaintext.

Zero out buffers after use; short‑lived processes for env export.

Refuse to print values unless explicitly --plain (default off); CI uses JSON→masking.

Sign .secrets.yml with Ed25519; verify on load.

OS keyring support; fallback to passphrase KDF.

Rate-limit get in CI to avoid logs spam.

11) Cursor Task Prompts (Copy/Paste Into Cursor)

T1 — Initialize repo root files Create the following files with professional content: README.md, LICENSE (MIT), .gitignore (Go+Node+Python+PHP), .gitattributes, .editorconfig, CODEOWNERS, CONTRIBUTING.md, SECURITY.md, Makefile, .pre-commit-config.yaml, .github/workflows/ci.yml, issue/PR templates. Ensure CI runs Go, Node, Python, and PHP unit tests in a matrix.

T2 — Spec & Threat Model Create core/spec/USM-Format-v1.md, core/spec/Threat-Model.md, and three ADRs under core/spec/ADRs/. Follow the architecture in this blueprint. Include diagrams in Mermaid.

T3 — Go crypto core Scaffold core/crypto with subpackages aead, kdf, asym, sign, securemem, file. Implement AES‑GCM, scrypt,