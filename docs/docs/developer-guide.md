# Developer Guide

This guide provides comprehensive information for developers who want to use USM in their projects.

## Table of Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Working with Profiles](#working-with-profiles)
- [Team Collaboration](#team-collaboration)
- [Cloud Sync](#cloud-sync)
- [Security Best Practices](#security-best-practices)

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

## Basic Usage

### Initialize a Project

```bash
usm init
```

This creates a `.secrets.yml` file and a `.usmkey` file.

### Set a Secret

```bash
usm set DB_URL=postgresql://localhost:5432/mydb
```

### Get a Secret

```bash
usm get DB_URL
```

### List Secrets

```bash
usm list
```

### Delete a Secret

```bash
usm delete DB_URL
```

## Working with Profiles

USM supports multiple profiles for different environments (e.g., dev, staging, prod).

### Create a Profile

```bash
usm profiles add staging
```

### List Profiles

```bash
usm profiles list
```

### Set a Secret in a Specific Profile

```bash
usm set --profile staging DB_URL=postgresql://staging.example.com:5432/mydb
```

### Get a Secret from a Specific Profile

```bash
usm get --profile staging DB_URL
```

## Team Collaboration

USM makes it easy to share secrets securely with your team.

### Add a Recipient

```bash
usm recipients add alice@example.com
```

### List Recipients

```bash
usm recipients list
```

### Remove a Recipient

```bash
usm recipients rm alice@example.com
```

When you add a recipient, USM will re-encrypt all secrets so that the new recipient can access them.

## Cloud Sync

USM can sync secrets with cloud providers like AWS, GCP, and HashiCorp Vault.

### Sync with AWS Secrets Manager

```bash
usm sync aws
```

### Sync with GCP Secret Manager

```bash
usm sync gcp
```

### Sync with HashiCorp Vault

```bash
usm sync vault
```

## Security Best Practices

1. **Never commit secrets to version control**: USM automatically adds `.secrets.yml` to your `.gitignore` file.

2. **Rotate secrets regularly**: Use `usm rotate` to rotate secrets.

3. **Limit access**: Only add recipients who absolutely need access to the secrets.

4. **Use strong passwords**: When setting up USM, use a strong password for the master key.

5. **Keep USM updated**: Regularly update USM to get the latest security fixes.