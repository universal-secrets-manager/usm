package asym

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/curve25519"
)

// GenerateKeyPair generates a new X25519 key pair.
func GenerateKeyPair() (publicKey, privateKey []byte, err error) {
	privateKey = make([]byte, curve25519.ScalarSize)
	if _, err := rand.Read(privateKey); err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	// Clamp the private key as required by Curve25519
	privateKey[0] &= 248
	privateKey[31] &= 127
	privateKey[31] |= 64

	publicKey, err = curve25519.X25519(privateKey, curve25519.Basepoint)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate public key: %w", err)
	}

	return publicKey, privateKey, nil
}

// DeriveSharedSecret derives a shared secret from a private key and a public key.
func DeriveSharedSecret(privateKey, publicKey []byte) ([]byte, error) {
	sharedSecret, err := curve25519.X25519(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive shared secret: %w", err)
	}
	return sharedSecret, nil
}