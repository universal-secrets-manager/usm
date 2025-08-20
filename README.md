# Universal Secrets Manager (USM)

A professional, production-grade open-source developer tool that replaces `.env` files with a secure, standardized, cross-platform secrets manager.

## 🎯 Core Goals

1. Provide a **universal encrypted secrets format** (a `.secrets.yml` file).
2. Create a **CLI tool (`usm`)** to manage secrets locally and in CI/CD.
3. Offer **language SDKs** (Node.js, Python, PHP, Go) for easy integration.
4. Support **team collaboration** (sharing, encryption, versioning).
5. Allow **optional sync with cloud providers** (AWS, GCP, Vault, etc.).
6. Be **local-first, offline-capable**, and **cross-platform** (Windows, macOS, Linux).

## 🚀 Quick Start

```bash
# Install the CLI
go install github.com/universal-secrets-manager/usm/cli/usm@latest

# Initialize a new secrets file
usm init

# Set a secret
usm set DB_URL=postgresql://localhost:5432/mydb

# Get a secret
usm get DB_URL

# List all keys
usm list
```

## 🛠️ Features

- **Encrypted Secrets**: AES-256-GCM for data, X25519 for sharing.
- **Team Collaboration**: Share secrets securely with teammates.
- **Language SDKs**: Node.js, Python, PHP, Go.
- **CI/CD Integration**: GitHub Actions, GitLab CI, Jenkins.
- **Cloud Sync (Optional)**: Sync with AWS, GCP, HashiCorp Vault.
- **Offline First**: Works without internet.
- **Cross-Platform**: Windows, macOS, Linux.

## 📦 Installation

### CLI

#### Download Pre-built Binaries (Recommended)

Download the latest binary for your platform from [Releases](https://github.com/universal-secrets-manager/usm/releases):

- **Linux**: `usm-linux-amd64`
- **macOS**: `usm-darwin-amd64`
- **Windows**: `usm-windows-amd64.exe`

Make the binary executable and move it to a directory in your PATH:

```bash
# Linux/macOS
chmod +x usm-*
sudo mv usm-* /usr/local/bin/usm

# Windows (in Command Prompt as Administrator)
move usm-windows-amd64.exe C:\Windows\System32\usm.exe
```

#### Install via Go

```bash
go install github.com/universal-secrets-manager/usm/cli/usm@latest
```

### SDKs

#### Node.js

```bash
npm install @usm/secrets
```

#### Python

```bash
pip install usm-secrets
```

#### PHP

```bash
composer require usm/secrets
```

#### Go

```bash
go get github.com/universal-secrets-manager/usm/sdks/go
```

## 📚 Documentation

See the [official documentation](./docs) for detailed usage, API references, and security information.

## 🤝 Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for development setup and guidelines.

## 📄 License

[MIT](./LICENSE)