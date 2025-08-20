package usm

import (
	"testing"
)

func TestLoad(t *testing.T) {
	// This is a placeholder test
	// In a real test, you would mock the file system
	usm, err := Load("./test/fixtures/.secrets.yml")
	if err != nil {
		t.Fatalf("Failed to load secrets file: %v", err)
	}

	if usm == nil {
		t.Error("Expected USM instance, got nil")
	}
}

func TestGet(t *testing.T) {
	// This is a placeholder test
	usm, err := Load("./test/fixtures/.secrets.yml")
	if err != nil {
		t.Fatalf("Failed to load secrets file: %v", err)
	}

	value, err := usm.Get("TEST_KEY")
	if err != nil {
		t.Fatalf("Failed to get secret: %v", err)
	}

	if value != "decrypted_value_for_TEST_KEY" {
		t.Errorf("Expected 'decrypted_value_for_TEST_KEY', got '%s'", value)
	}
}