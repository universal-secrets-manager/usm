# Quickstart

Get started with USM in a few simple steps.

## Installation

### CLI

#### Option 1: Download Pre-built Binary (Recommended)

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

#### Option 2: Install via Go

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

## Initialize a Project

```bash
usm init
```

This creates a `.secrets.yml` file and a `.usmkey` file.

## Set a Secret

```bash
usm set DB_URL=postgresql://localhost:5432/mydb
```

## Get a Secret

```bash
usm get DB_URL
```

## List Secrets

```bash
usm list
```