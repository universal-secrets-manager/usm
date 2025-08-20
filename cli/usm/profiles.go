package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto/file"
)

var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "Manage profiles",
	Long:  "List, add, or remove profiles.",
}

var profilesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List profiles",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// List profiles
		if len(secretsFile.Profiles) == 0 {
			fmt.Println("No profiles found")
			return
		}

		fmt.Println("Profiles:")
		for _, profile := range secretsFile.Profiles {
			fmt.Printf("  - %s\n", profile)
		}
	},
}

var profilesAddCmd = &cobra.Command{
	Use:   "add PROFILE",
	Short: "Add a profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// Check if profile already exists
		for _, p := range secretsFile.Profiles {
			if p == profile {
				fmt.Printf("Profile '%s' already exists\n", profile)
				return
			}
		}

		// Add profile
		secretsFile.Profiles = append(secretsFile.Profiles, profile)

		// Initialize secrets map for the new profile
		if secretsFile.Secrets[profile] == nil {
			secretsFile.Secrets[profile] = make(map[string]file.EncryptedSecret)
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

		fmt.Printf("Profile '%s' added successfully\n", profile)
	},
}

var profilesRemoveCmd = &cobra.Command{
	Use:   "rm PROFILE",
	Short: "Remove a profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		// Find and remove profile
		found := false
		for i, p := range secretsFile.Profiles {
			if p == profile {
				// Remove profile from slice
				secretsFile.Profiles = append(secretsFile.Profiles[:i], secretsFile.Profiles[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Profile '%s' not found\n", profile)
			return
		}

		// Remove secrets for the profile
		delete(secretsFile.Secrets, profile)

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

		fmt.Printf("Profile '%s' removed successfully\n", profile)
	},
}

func init() {
	profilesCmd.AddCommand(profilesListCmd)
	profilesCmd.AddCommand(profilesAddCmd)
	profilesCmd.AddCommand(profilesRemoveCmd)
	rootCmd.AddCommand(profilesCmd)
}