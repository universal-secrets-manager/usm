package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto/asym"
)

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share secrets with a user",
	Long:  "Manage sharing of secrets with other users.",
}

var shareAddCmd = &cobra.Command{
	Use:   "add user@org",
	Short: "Add a recipient to share secrets with",
	Long:  "Add a recipient to the team to share secrets.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]

		// For now, we'll just print a message
		// In a real implementation, this would:
		// 1. Generate a new X25519 key pair for the recipient
		// 2. Add the recipient's public key to the team file
		// 3. Re-encrypt all secrets for the new recipient

		// Generate a new X25519 key pair for the recipient
		publicKey, privateKey, err := asym.GenerateKeyPair()
		if err != nil {
			fmt.Printf("Failed to generate key pair: %v\n", err)
			return
		}

		// For now, we'll just save the private key to a file
		// In a real implementation, this would be sent securely to the recipient
		privateKeyFile := fmt.Sprintf("%s.private.key", email)
		err = os.WriteFile(privateKeyFile, privateKey, 0600)
		if err != nil {
			fmt.Printf("Failed to save private key: %v\n", err)
			return
		}

		// Create a simple team file structure
		teamFile := fmt.Sprintf("team:\n  - email: %s\n    public_key: %x\n", email, publicKey)

		// Save the team file
		err = os.WriteFile(".usm.team", []byte(teamFile), 0644)
		if err != nil {
			fmt.Printf("Failed to save team file: %v\n", err)
			return
		}

		fmt.Printf("Recipient '%s' added successfully\n", email)
		fmt.Printf("Private key saved to %s\n", privateKeyFile)
		fmt.Printf("Team file saved to .usm.team\n")
	},
}

func init() {
	shareCmd.AddCommand(shareAddCmd)
	rootCmd.AddCommand(shareCmd)
}