package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto/file"
)

var deleteCmd = &cobra.Command{
	Use:   "delete KEY",
	Short: "Delete a secret",
	Long:  "Delete a secret from the secrets file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// Delete the secret from the dev profile
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

		delete(devSecrets, key)
		secretsFile.Secrets["dev"] = devSecrets

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
		logAuditEntry("DELETE", key, "dev")

		fmt.Printf("Secret '%s' deleted successfully\n", key)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}