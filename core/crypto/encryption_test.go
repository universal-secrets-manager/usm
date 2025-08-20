package crypto

import (
	"crypto/rand"
	"testing"
)

func TestEncryptDecryptSecret(t *testing.T) {
	// Generate a project key
	projectKey, err := GenerateProjectKey()
	if err != nil {
		t.Fatalf("Failed to generate project key: %v", err)
	}

	// Generate a file key
	fileKey, err := GenerateFileKey()
	if err != nil {
		t.Fatalf("Failed to generate file key: %v", err)
	}

	// Create a secret value and AAD
	value := []byte("test secret value")
	aad := []byte("test additional authenticated data")

	// Encrypt the secret
	encryptedSecret, err := EncryptSecret(value, projectKey, fileKey, aad)
	if err != nil {
		t.Fatalf("Failed to encrypt secret: %v", err)
	}

	// Verify that all fields are populated
	if encryptedSecret.ProjectKeyEnc == nil {
		t.Error("ProjectKeyEnc is nil")
	}
	if encryptedSecret.FkNonce == nil {
		t.Error("FkNonce is nil")
	}
	if encryptedSecret.AAD == nil {
		t.Error("AAD is nil")
	}
	if encryptedSecret.ValueNonce == nil {
		t.Error("ValueNonce is nil")
	}
	if encryptedSecret.ValueCT == nil {
		t.Error("ValueCT is nil")
	}

	// Decrypt the secret
	decryptedValue, err := DecryptSecret(encryptedSecret, projectKey, aad)
	if err != nil {
		t.Fatalf("Failed to decrypt secret: %v", err)
	}

	// Verify that the decrypted value matches the original
	if string(decryptedValue) != string(value) {
		t.Errorf("Decrypted value does not match original. Got %s, want %s", string(decryptedValue), string(value))
	}

	// Try to decrypt with wrong AAD (should fail)
	_, err = DecryptSecret(encryptedSecret, projectKey, []byte("wrong aad"))
	if err == nil {
		t.Error("Decryption should have failed with wrong AAD")
	}
}

func TestDeriveProjectKeyFromPassphrase(t *testing.T) {
	// Create a passphrase and salt
	passphrase := []byte("test passphrase")
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		t.Fatalf("Failed to generate salt: %v", err)
	}

	// Derive a project key from the passphrase
	projectKey, err := DeriveProjectKeyFromPassphrase(passphrase, salt)
	if err != nil {
		t.Fatalf("Failed to derive project key from passphrase: %v", err)
	}

	// Verify that the project key is not nil
	if projectKey == nil {
		t.Error("Project key is nil")
	}

	// Verify that the project key has the correct length
	if len(projectKey.Key) != 32 {
		t.Errorf("Project key has incorrect length. Got %d, want %d", len(projectKey.Key), 32)
	}

	// Derive the project key again with the same parameters
	projectKey2, err := DeriveProjectKeyFromPassphrase(passphrase, salt)
	if err != nil {
		t.Fatalf("Failed to derive project key from passphrase: %v", err)
	}

	// Verify that both project keys are equal
	if string(projectKey.Key) != string(projectKey2.Key) {
		t.Error("Derived project keys do not match")
	}
}