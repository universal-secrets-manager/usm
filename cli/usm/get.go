package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto"
)

var getCmd = &cobra.Command{
	Use:   "get KEY",
	Short: "Get a secret",
	Long:  "Get a secret value from the secrets file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// Get the secret value from the dev profile
		devSecrets, exists := secretsFile.Secrets["dev"]
		if !exists {
			fmt.Printf("No secrets found for dev profile\n")
			return
		}

		secret, exists := devSecrets[key]
		if !exists {
			fmt.Printf("Secret '%s' not found\n", key)
			return
		}

		// For now, we'll use a fixed project key
		// In a real implementation, this would be loaded from a .usmkey file
		projectKey := &crypto.ProjectKey{
			Key: make([]byte, 32),
		}

		// Create the encrypted secret structure
		encryptedSecret := &crypto.EncryptedSecret{
			ProjectKeyEnc: secret.FkEnc,
			FkNonce:       secret.FkNonce,
			AAD:           secret.AAD,
			ValueNonce:    secret.Nonce,
			ValueTag:      secret.Tag,
			ValueCT:       secret.CT,
		}

		// Create AAD (Additional Authenticated Data)
		aad := []byte(fmt.Sprintf("dev:%s", key))

		// Decrypt the secret value
		value, err := crypto.DecryptSecret(encryptedSecret, projectKey, aad)
		if err != nil {
			fmt.Printf("Failed to decrypt secret: %v\n", err)
			return
		}

		fmt.Printf("%s\n", string(value))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}