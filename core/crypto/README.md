# Go Crypto Core Package

This package provides the core cryptographic functions for USM.

## Installation

```bash
go get github.com/universal-secrets-manager/usm/core/crypto
```

## Subpackages

- `aead`: AES-GCM implementation
- `kdf`: scrypt implementation
- `asym`: X25519 ECDH implementation
- `sign`: Ed25519 signature implementation
- `securemem`: Secure memory zeroization helpers
- `file`: File format marshaling with stable YAML ordering

## Example Usage

```go
package main

import (
	"fmt"
	"log"
	"github.com/universal-secrets-manager/usm/core/crypto"
)

func main() {
	// Create a 32-byte key for AES-256
	key := make([]byte, 32)
	
	// The value to encrypt
	value := []byte("This is a secret value")
	
	// Additional authenticated data
	aad := []byte("example-aad")
	
	// Encrypt the value
	ct, nonce, err := crypto.Encrypt(value, aad, key)
	if err != nil {
		log.Fatalf("Failed to encrypt: %v", err)
	}
	
	// Decrypt the value
	decrypted, err := crypto.Decrypt(ct, nonce, aad, key)
	if err != nil {
		log.Fatalf("Failed to decrypt: %v", err)
	}
	
	fmt.Printf("Decrypted value: %s\n", decrypted)
}
```