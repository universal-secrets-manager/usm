package sign

import (
	"crypto/ed25519"
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
	
	// Verify that the keys are of the correct type
	if len(publicKey) != ed25519.PublicKeySize {
		t.Errorf("Public key has incorrect size. Got %d, want %d", len(publicKey), ed25519.PublicKeySize)
	}
	
	if len(privateKey) != ed25519.PrivateKeySize {
		t.Errorf("Private key has incorrect size. Got %d, want %d", len(privateKey), ed25519.PrivateKeySize)
	}
}

func TestSignAndVerify(t *testing.T) {
	// Generate a key pair
	publicKey, privateKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Create a message to sign
	message := []byte("test message")

	// Sign the message
	signature := Sign(privateKey, message)
	
	// Verify that the signature is not nil
	if signature == nil {
		t.Error("Signature is nil")
	}
	
	// Verify the signature
	if !Verify(publicKey, message, signature) {
		t.Error("Failed to verify signature")
	}
	
	// Verify that the signature is invalid for a different message
	differentMessage := []byte("different message")
	if Verify(publicKey, differentMessage, signature) {
		t.Error("Signature should be invalid for a different message")
	}
}