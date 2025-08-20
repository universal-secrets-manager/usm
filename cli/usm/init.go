package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto/file"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new secrets file",
	Long:  "Initialize a new secrets file with a project key.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if .secrets.yml already exists
		if _, err := os.Stat(".secrets.yml"); err == nil {
			fmt.Println("Secrets file already exists")
			return
		}

		// Create a new secrets file
		secretsFile := file.NewSecretsFile()

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

		fmt.Println("Secrets file initialized successfully")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}