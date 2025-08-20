package aead

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// AESGCM implements AES-GCM encryption and decryption.
type AESGCM struct {
	aead cipher.AEAD
}

// NewAESGCM creates a new AESGCM instance.
func NewAESGCM(key []byte) (*AESGCM, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &AESGCM{aead: aead}, nil
}

// Encrypt encrypts plaintext with additional data.
func (g *AESGCM) Encrypt(plaintext, additionalData []byte) (ciphertext, nonce []byte, err error) {
	nonce = make([]byte, g.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext = g.aead.Seal(nil, nonce, plaintext, additionalData)
	return ciphertext, nonce, nil
}

// Decrypt decrypts ciphertext with nonce and additional data.
func (g *AESGCM) Decrypt(ciphertext, nonce, additionalData []byte) (plaintext []byte, err error) {
	plaintext, err = g.aead.Open(nil, nonce, ciphertext, additionalData)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}
	return plaintext, nil
}