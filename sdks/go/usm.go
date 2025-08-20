package usm

import (
	"os"

	"github.com/universal-secrets-manager/usm/core/crypto/file"
)

// USM represents a loaded secrets file.
type USM struct {
	secretsFile *file.SecretsFile
	cache       map[string]string
	cacheExpiry int64 // Unix timestamp in seconds
}

// Load loads a USM secrets file.
func Load(filePath string) (*USM, error) {
	if filePath == "" {
		var err error
		filePath, err = locateSecretsFile()
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	secretsFile, err := file.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return &USM{
		secretsFile: secretsFile,
		cache:       make(map[string]string),
		cacheExpiry: 300, // 5 minutes
	}, nil
}

// Get retrieves and decrypts a secret.
func (u *USM) Get(key string) (string, error) {
	// Implementation to get and decrypt a secret
	// This is a placeholder
	return "decrypted_value_for_" + key, nil
}

// locateSecretsFile locates the .secrets.yml file in the current or parent directories.
func locateSecretsFile() (string, error) {
	return file.LocateSecretsFile()
}
