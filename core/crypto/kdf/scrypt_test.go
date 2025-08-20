package kdf

import (
	"testing"
)

func TestDeriveKey(t *testing.T) {
	// Create test parameters
	params := DefaultParams()
	
	// Create a passphrase and salt
	passphrase := []byte("test passphrase")
	salt := make([]byte, 32)
	
	// Derive a key
	key, err := DeriveKey(passphrase, salt, params)
	if err != nil {
		t.Fatalf("Failed to derive key: %v", err)
	}
	
	// Verify that the key is not nil
	if key == nil {
		t.Error("Derived key is nil")
	}
	
	// Verify that the key has the expected length (32 bytes for AES-256)
	if len(key) != 32 {
		t.Errorf("Derived key has incorrect length. Got %d, want %d", len(key), 32)
	}
	
	// Derive the key again with the same parameters
	key2, err := DeriveKey(passphrase, salt, params)
	if err != nil {
		t.Fatalf("Failed to derive key: %v", err)
	}
	
	// Verify that both keys are equal
	if string(key) != string(key2) {
		t.Error("Derived keys do not match")
	}
}