package asym

import (
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	// Generate a key pair
	publicKey, privateKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Verify that both keys are not nil
	if publicKey == nil {
		t.Error("Public key is nil")
	}

	if privateKey == nil {
		t.Error("Private key is nil")
	}
}

func TestDeriveSharedSecret(t *testing.T) {
	// Generate two key pairs
	publicKey1, privateKey1, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate first key pair: %v", err)
	}

	publicKey2, privateKey2, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate second key pair: %v", err)
	}

	// Derive shared secrets
	sharedSecret1, err := DeriveSharedSecret(privateKey1, publicKey2)
	if err != nil {
		t.Fatalf("Failed to derive first shared secret: %v", err)
	}

	sharedSecret2, err := DeriveSharedSecret(privateKey2, publicKey1)
	if err != nil {
		t.Fatalf("Failed to derive second shared secret: %v", err)
	}

	// Verify that both shared secrets are equal
	if string(sharedSecret1) != string(sharedSecret2) {
		t.Error("Shared secrets do not match")
	}
}