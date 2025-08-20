package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto/asym"
	"gopkg.in/yaml.v3"
)

var recipientsCmd = &cobra.Command{
	Use:   "recipients",
	Short: "Manage recipients",
	Long:  "List, add, or remove recipients.",
}

var recipientsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recipients",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if team file exists
		if _, err := os.Stat(".usm.team"); os.IsNotExist(err) {
			fmt.Println("No recipients found")
			return
		}

		// Read the team file
		data, err := os.ReadFile(".usm.team")
		if err != nil {
			fmt.Printf("Failed to read team file: %v\n", err)
			return
		}

		// Parse the team file
		var team struct {
			Team []struct {
				Email     string `yaml:"email"`
				PublicKey string `yaml:"public_key"`
			} `yaml:"team"`
		}

		err = yaml.Unmarshal(data, &team)
		if err != nil {
			fmt.Printf("Failed to parse team file: %v\n", err)
			return
		}

		// List recipients
		if len(team.Team) == 0 {
			fmt.Println("No recipients found")
			return
		}

		fmt.Println("Recipients:")
		for _, recipient := range team.Team {
			fmt.Printf("  - %s\n", recipient.Email)
		}
	},
}

var recipientsAddCmd = &cobra.Command{
	Use:   "add user@org",
	Short: "Add a recipient",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]

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

		// Read existing team file or create new one
		var team struct {
			Team []struct {
				Email     string `yaml:"email"`
				PublicKey string `yaml:"public_key"`
			} `yaml:"team"`
		}

		if _, err := os.Stat(".usm.team"); err == nil {
			// Read existing team file
			data, err := os.ReadFile(".usm.team")
			if err != nil {
				fmt.Printf("Failed to read team file: %v\n", err)
				return
			}

			err = yaml.Unmarshal(data, &team)
			if err != nil {
				fmt.Printf("Failed to parse team file: %v\n", err)
				return
			}
		}

		// Add new recipient
		team.Team = append(team.Team, struct {
			Email     string `yaml:"email"`
			PublicKey string `yaml:"public_key"`
		}{
			Email:     email,
			PublicKey: fmt.Sprintf("%x", publicKey),
		})

		// Save the team file
		data, err := yaml.Marshal(team)
		if err != nil {
			fmt.Printf("Failed to marshal team file: %v\n", err)
			return
		}

		err = os.WriteFile(".usm.team", data, 0644)
		if err != nil {
			fmt.Printf("Failed to save team file: %v\n", err)
			return
		}

		fmt.Printf("Recipient '%s' added successfully\n", email)
		fmt.Printf("Private key saved to %s\n", privateKeyFile)
	},
}

var recipientsRemoveCmd = &cobra.Command{
	Use:   "rm user@org",
	Short: "Remove a recipient",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]

		// Check if team file exists
		if _, err := os.Stat(".usm.team"); os.IsNotExist(err) {
			fmt.Println("No recipients found")
			return
		}

		// Read the team file
		data, err := os.ReadFile(".usm.team")
		if err != nil {
			fmt.Printf("Failed to read team file: %v\n", err)
			return
		}

		// Parse the team file
		var team struct {
			Team []struct {
				Email     string `yaml:"email"`
				PublicKey string `yaml:"public_key"`
			} `yaml:"team"`
		}

		err = yaml.Unmarshal(data, &team)
		if err != nil {
			fmt.Printf("Failed to parse team file: %v\n", err)
			return
		}

		// Find and remove recipient
		found := false
		for i, recipient := range team.Team {
			if recipient.Email == email {
				// Remove recipient from slice
				team.Team = append(team.Team[:i], team.Team[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Recipient '%s' not found\n", email)
			return
		}

		// Save the team file
		data, err = yaml.Marshal(team)
		if err != nil {
			fmt.Printf("Failed to marshal team file: %v\n", err)
			return
		}

		err = os.WriteFile(".usm.team", data, 0644)
		if err != nil {
			fmt.Printf("Failed to save team file: %v\n", err)
			return
		}

		// Remove the private key file
		privateKeyFile := fmt.Sprintf("%s.private.key", email)
		if _, err := os.Stat(privateKeyFile); err == nil {
			err = os.Remove(privateKeyFile)
			if err != nil {
				fmt.Printf("Failed to remove private key file: %v\n", err)
				// Don't return here, as the main operation (removing from team file) was successful
			}
		}

		fmt.Printf("Recipient '%s' removed successfully\n", email)
	},
}

func init() {
	recipientsCmd.AddCommand(recipientsListCmd)
	recipientsCmd.AddCommand(recipientsAddCmd)
	recipientsCmd.AddCommand(recipientsRemoveCmd)
	rootCmd.AddCommand(recipientsCmd)
}