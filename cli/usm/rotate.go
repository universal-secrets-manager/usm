package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto"
	"github.com/universal-secrets-manager/usm/core/crypto/file"
)

var rotateCmd = &cobra.Command{
	Use:   "rotate [KEY]",
	Short: "Rotate a secret or key",
	Long:  "Rotate a secret value, project file key, or recipient keys.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a key to rotate")
			return
		}

		key := args[0]

		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// Check if the secret exists
		devSecrets, exists := secretsFile.Secrets["dev"]
		if !exists {
			fmt.Printf("No secrets found for dev profile\n")
			return
		}

		_, exists = devSecrets[key]
		if !exists {
			fmt.Printf("Secret '%s' not found\n", key)
			return
		}

		// For now, we'll use a fixed project key
		// In a real implementation, this would be loaded from a .usmkey file
		projectKey := &crypto.ProjectKey{
			Key: make([]byte, 32),
		}

		// Generate a new file key
		fileKey, err := crypto.GenerateFileKey()
		if err != nil {
			fmt.Printf("Failed to generate file key: %v\n", err)
			return
		}

		// Get the current encrypted secret
		currentSecret := devSecrets[key]

		// Create the encrypted secret structure for decryption
		encryptedSecret := &crypto.EncryptedSecret{
			ProjectKeyEnc: currentSecret.FkEnc,
			FkNonce:       currentSecret.FkNonce,
			AAD:           currentSecret.AAD,
			ValueNonce:    currentSecret.Nonce,
			ValueTag:      currentSecret.Tag,
			ValueCT:       currentSecret.CT,
		}

		// Create AAD (Additional Authenticated Data)
		aad := []byte(fmt.Sprintf("dev:%s", key))

		// Decrypt the current secret value
		value, err := crypto.DecryptSecret(encryptedSecret, projectKey, aad)
		if err != nil {
			fmt.Printf("Failed to decrypt secret: %v\n", err)
			return
		}

		// Re-encrypt the secret value with the new file key
		newEncryptedSecret, err := crypto.EncryptSecret(value, projectKey, fileKey, aad)
		if err != nil {
			fmt.Printf("Failed to encrypt secret: %v\n", err)
			return
		}

		// Update the secret in the secrets file
		secretsFile.Secrets["dev"][key] = file.EncryptedSecret{
			FkEnc:    newEncryptedSecret.ProjectKeyEnc,
			FkNonce:  newEncryptedSecret.FkNonce,
			AAD:      newEncryptedSecret.AAD,
			Nonce:    newEncryptedSecret.ValueNonce,
			Tag:      newEncryptedSecret.ValueTag,
			CT:       newEncryptedSecret.ValueCT,
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
		logAuditEntry("ROTATE", key, "dev")

		fmt.Printf("Secret '%s' rotated successfully\n", key)
	},
}

func init() {
	rootCmd.AddCommand(rotateCmd)
}