# Threat Model

## Assets

- Secrets at rest
- Keys at rest
- Keys in memory
- CI environments

## Adversaries

- Repository reader
- Stolen laptop
- Malicious insider
- Compromised CI

## Mitigations

- GCM AEAD
- Zeroization
- OS keyrings
- Masking in logs
- Short-lived process environment
- SIGTERM traps

## Out of Scope

- Running process memory scrape
- Fully compromised developer box