// Package crypto provides the core cryptographic functions for USM.
package crypto

import (
	"fmt"
	"github.com/universal-secrets-manager/usm/core/crypto/aead"
)

// PFK represents a Project File Key.
type PFK struct {
	Key []byte
}

// FK represents a File Key.
type FK struct {
	Key []byte
}

// Recipient represents a recipient of encrypted secrets.
type Recipient struct {
	Type string
	ID   string
	Pub  []byte
}

// Signature represents a signature for a secrets file.
type Signature struct {
	Alg   string
	KeyID string
	Sig   []byte
}

// Encrypt encrypts a value with additional authenticated data (AAD).
func Encrypt(value []byte, aad []byte, key []byte) (ct, nonce []byte, err error) {
	aesgcm, err := aead.NewAESGCM(key)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create AES-GCM: %w", err)
	}

	ct, nonce, err = aesgcm.Encrypt(value, aad)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encrypt: %w", err)
	}

	return ct, nonce, nil
}

// Decrypt decrypts a ciphertext with nonce and AAD.
func Decrypt(ct, nonce, aad, key []byte) ([]byte, error) {
	aesgcm, err := aead.NewAESGCM(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM: %w", err)
	}

	plaintext, err := aesgcm.Decrypt(ct, nonce, aad)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}