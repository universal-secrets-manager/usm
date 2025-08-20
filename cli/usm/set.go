package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto"
	"github.com/universal-secrets-manager/usm/core/crypto/file"
)

var setCmd = &cobra.Command{
	Use:   "set KEY=VALUE",
	Short: "Set a secret",
	Long:  "Set a secret value in the secrets file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the KEY=VALUE argument
		parts := strings.SplitN(args[0], "=", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format. Use KEY=VALUE")
			return
		}
		key := parts[0]
		value := parts[1]

		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// For now, we'll use a fixed project key
		// In a real implementation, this would be loaded from a .usmkey file
		projectKey := &crypto.ProjectKey{
			Key: make([]byte, 32),
		}

		// Generate a file key
		fileKey, err := crypto.GenerateFileKey()
		if err != nil {
			fmt.Printf("Failed to generate file key: %v\n", err)
			return
		}

		// Create AAD (Additional Authenticated Data)
		aad := []byte(fmt.Sprintf("dev:%s", key))

		// Encrypt the secret value
		encryptedSecret, err := crypto.EncryptSecret([]byte(value), projectKey, fileKey, aad)
		if err != nil {
			fmt.Printf("Failed to encrypt secret: %v\n", err)
			return
		}

		// Set the secret value in the dev profile
		if secretsFile.Secrets["dev"] == nil {
			secretsFile.Secrets["dev"] = make(map[string]file.EncryptedSecret)
		}
		
		secretsFile.Secrets["dev"][key] = file.EncryptedSecret{
			FkEnc:    encryptedSecret.ProjectKeyEnc,
			FkNonce:  encryptedSecret.FkNonce,
			AAD:      encryptedSecret.AAD,
			Nonce:    encryptedSecret.ValueNonce,
			Tag:      encryptedSecret.ValueTag,
			CT:       encryptedSecret.ValueCT,
		}

		// Marshal the secrets file
		data, err := file.Marshal(secretsFile)
		if err != nil {
			fmt.Printf("Failed to marshal secrets file: %v\n", err)
			return
		}

		// Write the secrets file to disk
		err = os.WriteFile(".secrets.yml", data, 0644)
		if err != nil {
			fmt.Printf("Failed to write secrets file: %v\n", err)
			return
		}

		// Log audit entry
		logAuditEntry("SET", key, "dev")

		fmt.Printf("Secret '%s' set successfully\n", key)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

// loadSecretsFile loads the secrets file from disk
func loadSecretsFile() (*file.SecretsFile, error) {
	// Read the secrets file
	data, err := os.ReadFile(".secrets.yml")
	if err != nil {
		return nil, fmt.Errorf("failed to read secrets file: %w", err)
	}

	// Unmarshal the secrets file
	secretsFile, err := file.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secrets file: %w", err)
	}

	return secretsFile, nil
}