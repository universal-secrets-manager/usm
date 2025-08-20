package file

import (
	"testing"
)

func TestMarshalUnmarshal(t *testing.T) {
	// Create a new secrets file
	sf := NewSecretsFile()

	// Marshal it
	data, err := Marshal(sf)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal it
	sf2, err := Unmarshal(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify that the unmarshaled file matches the original
	if sf2.USM != sf.USM {
		t.Errorf("USM version mismatch. Got %d, want %d", sf2.USM, sf.USM)
	}
}