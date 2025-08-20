package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"github.com/universal-secrets-manager/usm/core/crypto"
	"github.com/universal-secrets-manager/usm/core/crypto/file"
)

var syncCmd = &cobra.Command{
	Use:   "sync PROVIDER",
	Short: "Synchronize with cloud providers",
	Long:  "Synchronize secrets with AWS, GCP, or Vault.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]

		// Load the secrets file
		secretsFile, err := loadSecretsFile()
		if err != nil {
			fmt.Printf("Failed to load secrets file: %v\n", err)
			return
		}

		switch provider {
		case "aws":
			syncAWS(secretsFile)
		default:
			fmt.Printf("Unsupported provider: %s\n", provider)
			fmt.Println("Supported providers: aws")
		}
	},
}

func syncAWS(secretsFile *file.SecretsFile) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("Failed to load AWS config: %v\n", err)
		return
	}

	// Create Secrets Manager client
	client := secretsmanager.NewFromConfig(cfg)

	// For each secret in the dev profile, sync it to AWS
	devSecrets, exists := secretsFile.Secrets["dev"]
	if !exists {
		fmt.Println("No secrets found for dev profile")
		return
	}

	// For now, we'll use a fixed project key
	// In a real implementation, this would be loaded from a .usmkey file
	projectKey := &crypto.ProjectKey{
		Key: make([]byte, 32),
	}

	for key, secret := range devSecrets {
		// Create the encrypted secret structure for decryption
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
			fmt.Printf("Failed to decrypt secret '%s': %v\n", key, err)
			continue
		}

		// Create or update the secret in AWS Secrets Manager
		secretName := fmt.Sprintf("usm/%s", key)
		secretString := string(value)

		// Check if the secret already exists
		_, err = client.DescribeSecret(context.TODO(), &secretsmanager.DescribeSecretInput{
			SecretId: &secretName,
		})

		if err != nil {
			// Secret doesn't exist, create it
			_, err = client.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
				Name:         &secretName,
				SecretString: &secretString,
			})
			if err != nil {
				fmt.Printf("Failed to create secret '%s' in AWS: %v\n", key, err)
				continue
			}
			fmt.Printf("Created secret '%s' in AWS\n", key)
		} else {
			// Secret exists, update it
			_, err = client.PutSecretValue(context.TODO(), &secretsmanager.PutSecretValueInput{
				SecretId:     &secretName,
				SecretString: &secretString,
			})
			if err != nil {
				fmt.Printf("Failed to update secret '%s' in AWS: %v\n", key, err)
				continue
			}
			fmt.Printf("Updated secret '%s' in AWS\n", key)
		}
	}

	fmt.Println("AWS sync completed")
}

func init() {
	rootCmd.AddCommand(syncCmd)
}