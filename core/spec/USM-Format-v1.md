# USM Format Specification v1

## Header

```yaml
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
```

## Secrets Payload (by profile)

```yaml
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
```

## Key Derivation and Encryption

Deterministic AAD: `${repo_path}:${profile}:${key}`.

Key derivation: Passphrase → scrypt → MasterKey; MasterKey unwraps Project File Key (PFK). PFK encrypts FKs; FKs encrypt values.

## Rotation

- Rotate value → new FK
- Rotate PFK → rewrap FKs
- Rotate recipients → re-encrypt FK for new pubkeys

## Audit

Append-only log (hash-chained) in `.usm/audit.log`.