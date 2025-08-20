package crypto

import (
	"testing"
	"github.com/universal-secrets-manager/usm/core/crypto/asym"
	"github.com/universal-secrets-manager/usm/core/crypto/kdf"
	"github.com/universal-secrets-manager/usm/core/crypto/sign"
)

func TestEncryptDecrypt(t *testing.T) {
	// This is a simple test to verify that the encrypt/decrypt functions work
	key := make([]byte, 32)
	value := []byte("test secret")
	aad := []byte("test aad")

	// Encrypt the value
	ct, nonce, err := Encrypt(value, aad, key)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// Decrypt the value
	decrypted, err := Decrypt(ct, nonce, aad, key)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	// Verify that the decrypted value matches the original
	if string(decrypted) != string(value) {
		t.Errorf("Decrypted value does not match original. Got %s, want %s", string(decrypted), string(value))
	}
}

func TestCompleteCryptoFlow(t *testing.T) {
	// This test verifies a complete crypto flow:
	// 1. Generate X25519 key pairs for two parties
	// 2. Derive a shared secret
	// 3. Use scrypt to derive a key from a passphrase
	// 4. Use the derived key for AES-GCM encryption
	// 5. Sign the encrypted data
	// 6. Verify the signature

	// Step 1: Generate X25519 key pairs
	_, privateKey1, err := asym.GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate first key pair: %v", err)
	}

	publicKey2, _, err := asym.GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate second key pair: %v", err)
	}

	// Step 2: Derive a shared secret
	sharedSecret, err := asym.DeriveSharedSecret(privateKey1, publicKey2)
	if err != nil {
		t.Fatalf("Failed to derive shared secret: %v", err)
	}

	// Step 3: Use scrypt to derive a key from the shared secret
	salt := make([]byte, 32)
	params := kdf.DefaultParams()
	aesKey, err := kdf.DeriveKey(sharedSecret, salt, params)
	if err != nil {
		t.Fatalf("Failed to derive AES key: %v", err)
	}

	// Ensure the key is 32 bytes for AES-256
	if len(aesKey) != 32 {
		t.Fatalf("Derived key has incorrect length. Got %d, want %d", len(aesKey), 32)
	}

	// Step 4: Use the derived key for AES-GCM encryption
	value := []byte("test secret value")
	aad := []byte("test additional authenticated data")

	ct, nonce, err := Encrypt(value, aad, aesKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// Step 5: Sign the encrypted data
	signerPublicKey, signerPrivateKey, err := sign.GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate signer key pair: %v", err)
	}

	// Combine ct, nonce, and aad for signing
	signData := append(append(ct, nonce...), aad...)
	signature := sign.Sign(signerPrivateKey, signData)

	// Step 6: Verify the signature
	if !sign.Verify(signerPublicKey, signData, signature) {
		t.Error("Failed to verify signature")
	}

	// Verify decryption works
	decrypted, err := Decrypt(ct, nonce, aad, aesKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if string(decrypted) != string(value) {
		t.Errorf("Decrypted value does not match original. Got %s, want %s", string(decrypted), string(value))
	}
}