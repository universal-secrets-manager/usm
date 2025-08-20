package main

import (
	"fmt"
	"log"
	"github.com/universal-secrets-manager/usm/core/crypto"
)

func main() {
	// Example of using the crypto package to encrypt and decrypt a value
	
	// Create a 32-byte key for AES-256
	key := make([]byte, 32)
	
	// The value to encrypt
	value := []byte("This is a secret value")
	
	// Additional authenticated data
	aad := []byte("example-aad")
	
	// Encrypt the value
	fmt.Println("Encrypting value...")
	ct, nonce, err := crypto.Encrypt(value, aad, key)
	if err != nil {
		log.Fatalf("Failed to encrypt: %v", err)
	}
	
	fmt.Printf("Ciphertext: %x\n", ct)
	fmt.Printf("Nonce: %x\n", nonce)
	
	// Decrypt the value
	fmt.Println("Decrypting value...")
	decrypted, err := crypto.Decrypt(ct, nonce, aad, key)
	if err != nil {
		log.Fatalf("Failed to decrypt: %v", err)
	}
	
	fmt.Printf("Decrypted value: %s\n", decrypted)
	
	// Verify that the decrypted value matches the original
	if string(decrypted) == string(value) {
		fmt.Println("Encryption/decryption successful!")
	} else {
		fmt.Println("Encryption/decryption failed!")
	}
}