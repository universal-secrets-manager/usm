# Quickstart

Get started with USM in a few simple steps.

## Installation

### CLI

Download the latest binary from [Releases](https://github.com/universal-secrets-manager/usm/releases) or install via Go:

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