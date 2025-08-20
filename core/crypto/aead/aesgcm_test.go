package aead

import (
	"testing"
)

func TestAESGCM(t *testing.T) {
	// Create a 32-byte key for AES-256
	key := make([]byte, 32)
	
	// Create a new AESGCM instance
	aesgcm, err := NewAESGCM(key)
	if err != nil {
		t.Fatalf("Failed to create AESGCM: %v", err)
	}
	
	// Create a plaintext and additional data
	plaintext := []byte("test plaintext")
	additionalData := []byte("test additional data")
	
	// Encrypt the plaintext
	ciphertext, nonce, err := aesgcm.Encrypt(plaintext, additionalData)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}
	
	// Verify that ciphertext and nonce are not nil
	if ciphertext == nil {
		t.Error("Ciphertext is nil")
	}
	
	if nonce == nil {
		t.Error("Nonce is nil")
	}
	
	// Decrypt the ciphertext
	decrypted, err := aesgcm.Decrypt(ciphertext, nonce, additionalData)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	
	// Verify that the decrypted plaintext matches the original
	if string(decrypted) != string(plaintext) {
		t.Errorf("Decrypted plaintext does not match original. Got %s, want %s", string(decrypted), string(plaintext))
	}
	
	// Try to decrypt with wrong additional data (should fail)
	_, err = aesgcm.Decrypt(ciphertext, nonce, []byte("wrong additional data"))
	if err == nil {
		t.Error("Decryption should have failed with wrong additional data")
	}
}