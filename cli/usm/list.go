package main

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secrets",
	Long:  "List all secret keys in the secrets file.",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// List secrets from the dev profile
		devSecrets, exists := secretsFile.Secrets["dev"]
		if !exists {
			fmt.Println("No secrets found for dev profile")
			return
		}

		// Collect and sort the keys
		keys := make([]string, 0, len(devSecrets))
		for key := range devSecrets {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		// Print the keys
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}