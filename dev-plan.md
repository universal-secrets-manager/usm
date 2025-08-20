# 🧩 Master Prompt – Build Universal Secrets Manager (USM)

You are tasked with building **Universal Secrets Manager (USM)**: a professional, production-grade open-source developer tool that replaces `.env` files with a secure, standardized, cross-platform secrets manager.

### 🎯 Core Goals

1. Provide a **universal encrypted secrets format** (a `.secrets` file).
2. Create a **CLI tool (`usm`)** to manage secrets locally and in CI/CD.
3. Offer **language SDKs** (Node.js, Python, PHP, Go) for easy integration.
4. Support **team collaboration** (sharing, encryption, versioning).
5. Allow **optional sync with cloud providers** (AWS, GCP, Vault, etc.).
6. Be **local-first, offline-capable**, and **cross-platform** (Windows, macOS, Linux).

---

## 🛠️ Features to Implement

### 1. CLI Tool (`usm`)

* Written in **Go** (for portability and single-binary builds).
* Subcommands:

  * `usm init` → Create a `.secrets` file (AES-256 encrypted).
  * `usm set KEY=VALUE` → Add or update a secret.
  * `usm get KEY` → Retrieve a secret securely.
  * `usm list` → Show all keys (not values).
  * `usm delete KEY` → Remove a secret.
  * `usm share USER_EMAIL` → Share encrypted secrets with another dev.
  * `usm rotate KEY` → Rotate secret value and update versions.
  * `usm sync PROVIDER` → Push/pull secrets from AWS/GCP/Vault.

### 2. File Format (`.secrets`)

* Human-readable YAML/JSON with **encrypted values**. Example:

  ```yaml
  version: 1
  secrets:
    DB_URL: ENC(aes256:base64string)
    API_KEY: ENC(aes256:base64string)
  metadata:
    created: 2025-08-20
    updated: 2025-08-20
  ```
* Encrypted with a **project master key** (stored in `.usmkey`).

### 3. Language SDKs

* **Node.js SDK (npm package)** → `import usm from 'usm'`
* **Python SDK (PyPI package)** → `import usm`
* **PHP SDK (Composer package)** → `use Usm;`
* **Go SDK** → `import "github.com/usm/secrets"`
* All SDKs should:

  * Auto-detect `.secrets` file.
  * Decrypt on demand.
  * Cache securely in memory.

### 4. CI/CD Integration

* GitHub Action: `uses: usm/action@v1` → Injects secrets into workflow env.
* GitLab CI + Jenkins plugin → reads `.secrets` and exports vars.

### 5. Security & Compliance

* AES-256 encryption for values.
* RSA/ECC for team sharing (each dev has a keypair).
* Versioning + audit logs.
* Secrets never stored in plaintext anywhere.

### 6. Team Collaboration

* Each dev has a **public/private keypair** managed by `usm`.
* Sharing = encrypt secret for recipient’s public key.
* Rotation notifications via CLI.

### 7. Optional Features

* Local GUI dashboard (Electron or Tauri).
* Browser extension to inject secrets into local dev servers.
* Enterprise mode: central server + SSO integration.

---

## 📦 Technical Stack

* **CLI**: Go (compiled single-binary).
* **Encryption**: AES-256-GCM for symmetric, RSA/ECC for sharing.
* **File Format**: YAML with encrypted values.
* **SDKs**:

  * Node.js → TypeScript
  * Python → pure Python
  * PHP → pure PHP
  * Go → native library
* **CI/CD**: YAML plugins for GitHub Actions, GitLab, Jenkins.

---

## 📅 Development Phases

### Phase 1 – MVP (CLI + File Format)

* Implement `.secrets` file format.
* CLI commands: `init`, `set`, `get`, `list`.
* AES encryption with `.usmkey`.
* Node.js + Python SDKs (basic read/get).

### Phase 2 – Team Features

* Implement `share`, `rotate`, `delete`.
* RSA/ECC keypairs for each user.
* Add Go + PHP SDKs.

### Phase 3 – CI/CD & Sync

* GitHub Action + GitLab integration.
* `usm sync aws/gcp/vault`.
* Versioning + audit logs.

### Phase 4 – Advanced Features

* Local GUI manager.
* Enterprise features (SSO, policy enforcement).
* Browser extension for local dev.

---

## ✅ Deliverables

* `usm` CLI (cross-platform binaries).
* `.secrets` file format (spec + docs).
* SDKs: Node.js, Python, PHP, Go.
* CI/CD integrations.
* Documentation website (mkdocs or docusaurus).
* Example projects (Node, Python, PHP, Go apps using USM).

---

### 📢 Instructions for AI

* Write **clean, modular, production-grade code**.
* Follow **best practices for security** (no hardcoded keys, strong crypto).
* Ensure **cross-platform support** (Windows, Linux, macOS).
* Generate **documentation + tests** for each module.
* Keep it **open-source ready** (MIT license, contribution guide, README).

