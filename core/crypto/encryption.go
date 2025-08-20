package crypto

import (
	"crypto/rand"
	"fmt"
	"github.com/universal-secrets-manager/usm/core/crypto/aead"
	"github.com/universal-secrets-manager/usm/core/crypto/kdf"
	"github.com/universal-secrets-manager/usm/core/crypto/securemem"
	"golang.org/x/crypto/scrypt"
)

// ProjectKey represents a project key used for encrypting file keys
type ProjectKey struct {
	Key []byte
}

// FileKey represents a file key used for encrypting secret values
type FileKey struct {
	Key []byte
}

// EncryptedSecret represents an encrypted secret value
type EncryptedSecret struct {
	ProjectKeyEnc []byte // Encrypted file key
	FkNonce       []byte // Nonce for file key encryption
	FkTag         []byte // Tag for file key encryption
	AAD           []byte // Additional authenticated data
	ValueNonce    []byte // Nonce for value encryption
	ValueTag      []byte // Tag for value encryption
	ValueCT       []byte // Ciphertext of the secret value
}

// GenerateProjectKey generates a new project key
func GenerateProjectKey() (*ProjectKey, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("failed to generate project key: %w", err)
	}
	return &ProjectKey{Key: key}, nil
}

// GenerateFileKey generates a new file key
func GenerateFileKey() (*FileKey, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("failed to generate file key: %w", err)
	}
	return &FileKey{Key: key}, nil
}

// EncryptSecret encrypts a secret value using the project key and file key
func EncryptSecret(value []byte, projectKey *ProjectKey, fileKey *FileKey, aad []byte) (*EncryptedSecret, error) {
	// Encrypt the file key with the project key
	aesgcm, err := aead.NewAESGCM(projectKey.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM for project key: %w", err)
	}

	fkCT, fkNonce, err := aesgcm.Encrypt(fileKey.Key, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt file key: %w", err)
	}

	// Encrypt the secret value with the file key
	valueAESGCM, err := aead.NewAESGCM(fileKey.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM for file key: %w", err)
	}

	valueCT, valueNonce, err := valueAESGCM.Encrypt(value, aad)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt secret value: %w", err)
	}

	// Zeroize the file key after use
	securemem.Zeroize(fileKey.Key)

	return &EncryptedSecret{
		ProjectKeyEnc: fkCT,
		FkNonce:       fkNonce,
		FkTag:         nil, // Tag is included in CT for GCM
		AAD:           aad,
		ValueNonce:    valueNonce,
		ValueTag:      nil, // Tag is included in CT for GCM
		ValueCT:       valueCT,
	}, nil
}

// DecryptSecret decrypts a secret value using the project key
func DecryptSecret(encryptedSecret *EncryptedSecret, projectKey *ProjectKey, aad []byte) ([]byte, error) {
	// Decrypt the file key with the project key
	aesgcm, err := aead.NewAESGCM(projectKey.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM for project key: %w", err)
	}

	fileKeyBytes, err := aesgcm.Decrypt(encryptedSecret.ProjectKeyEnc, encryptedSecret.FkNonce, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt file key: %w", err)
	}

	// Decrypt the secret value with the file key
	fileKey := &FileKey{Key: fileKeyBytes}
	defer securemem.Zeroize(fileKey.Key)

	valueAESGCM, err := aead.NewAESGCM(fileKey.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM for file key: %w", err)
	}

	value, err := valueAESGCM.Decrypt(encryptedSecret.ValueCT, encryptedSecret.ValueNonce, aad)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt secret value: %w", err)
	}

	return value, nil
}

// DeriveProjectKeyFromPassphrase derives a project key from a passphrase using scrypt
func DeriveProjectKeyFromPassphrase(passphrase, salt []byte) (*ProjectKey, error) {
	// Use scrypt to derive a key from the passphrase
	params := kdf.DefaultParams()
	key, err := scrypt.Key(passphrase, salt, params.N, params.R, params.P, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to derive key from passphrase: %w", err)
	}

	return &ProjectKey{Key: key}, nil
}