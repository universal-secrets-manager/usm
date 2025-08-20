package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"sort"
	"time"

	"gopkg.in/yaml.v3"
)

// SecretsFile represents the structure of a .secrets.yml file.
type SecretsFile struct {
	USM      int      `yaml:"usm"`
	KDF      KDF      `yaml:"kdf"`
	Crypto   Crypto   `yaml:"crypto"`
	Profiles []string `yaml:"profiles"`
	Sign     Sign     `yaml:"sign"`
	Secrets  Secrets  `yaml:"secrets"`
	Metadata Metadata `yaml:"metadata"`
}

// KDF represents the key derivation function parameters.
type KDF struct {
	Name string `yaml:"name"`
	N    int    `yaml:"n"`
	R    int    `yaml:"r"`
	P    int    `yaml:"p"`
	Salt []byte `yaml:"salt"`
}

// Crypto represents the cryptographic parameters.
type Crypto struct {
	Cipher      string `yaml:"cipher"`
	NonceBytes  int    `yaml:"nonce_bytes"`
}

// Sign represents the signature information.
type Sign struct {
	Alg   string `yaml:"alg"`
	KeyID string `yaml:"keyid"`
	Sig   []byte `yaml:"sig"`
}

// Secrets represents the secrets by profile.
type Secrets map[string]map[string]EncryptedSecret

// EncryptedSecret represents an encrypted secret value.
type EncryptedSecret struct {
	FkEnc    []byte `yaml:"fk_enc"`
	FkNonce  []byte `yaml:"fk_nonce"`
	AAD      []byte `yaml:"aad"`
	Nonce    []byte `yaml:"nonce"`
	Tag      []byte `yaml:"tag"`
	CT       []byte `yaml:"ct"`
}

// MarshalYAML implements yaml.Marshaler for EncryptedSecret
func (es EncryptedSecret) MarshalYAML() (interface{}, error) {
	return struct {
		FkEnc   string `yaml:"fk_enc"`
		FkNonce string `yaml:"fk_nonce"`
		AAD     string `yaml:"aad"`
		Nonce   string `yaml:"nonce"`
		Tag     string `yaml:"tag"`
		CT      string `yaml:"ct"`
	}{
		FkEnc:   base64.StdEncoding.EncodeToString(es.FkEnc),
		FkNonce: base64.StdEncoding.EncodeToString(es.FkNonce),
		AAD:     base64.StdEncoding.EncodeToString(es.AAD),
		Nonce:   base64.StdEncoding.EncodeToString(es.Nonce),
		Tag:     base64.StdEncoding.EncodeToString(es.Tag),
		CT:      base64.StdEncoding.EncodeToString(es.CT),
	}, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for EncryptedSecret
func (es *EncryptedSecret) UnmarshalYAML(value *yaml.Node) error {
	var aux struct {
		FkEnc   string `yaml:"fk_enc"`
		FkNonce string `yaml:"fk_nonce"`
		AAD     string `yaml:"aad"`
		Nonce   string `yaml:"nonce"`
		Tag     string `yaml:"tag"`
		CT      string `yaml:"ct"`
	}
	
	if err := value.Decode(&aux); err != nil {
		return err
	}
	
	var err error
	if es.FkEnc, err = base64.StdEncoding.DecodeString(aux.FkEnc); err != nil {
		return err
	}
	if es.FkNonce, err = base64.StdEncoding.DecodeString(aux.FkNonce); err != nil {
		return err
	}
	if es.AAD, err = base64.StdEncoding.DecodeString(aux.AAD); err != nil {
		return err
	}
	if es.Nonce, err = base64.StdEncoding.DecodeString(aux.Nonce); err != nil {
		return err
	}
	if es.Tag, err = base64.StdEncoding.DecodeString(aux.Tag); err != nil {
		return err
	}
	if es.CT, err = base64.StdEncoding.DecodeString(aux.CT); err != nil {
		return err
	}
	
	return nil
}

// Metadata represents the metadata of the secrets file.
type Metadata struct {
	Created string `yaml:"created"`
	Updated string `yaml:"updated"`
	Version int    `yaml:"version"`
}

// Marshal marshals a SecretsFile to YAML with stable ordering.
func Marshal(sf *SecretsFile) ([]byte, error) {
	// Create a new SecretsFile with sorted profiles and secrets for deterministic output
	sortedSF := &SecretsFile{
		USM:      sf.USM,
		KDF:      sf.KDF,
		Crypto:   sf.Crypto,
		Profiles: make([]string, len(sf.Profiles)),
		Sign:     sf.Sign,
		Secrets:  make(Secrets),
		Metadata: sf.Metadata,
	}

	// Copy and sort profiles
	copy(sortedSF.Profiles, sf.Profiles)
	sort.Strings(sortedSF.Profiles)

	// Copy and sort secrets
	for profile, secrets := range sf.Secrets {
		sortedSecrets := make(map[string]EncryptedSecret)
		for key, secret := range secrets {
			sortedSecrets[key] = secret
		}
		sortedSF.Secrets[profile] = sortedSecrets
	}

	// Marshal to YAML
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err := encoder.Encode(sortedSF)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal secrets file: %w", err)
	}

	return buf.Bytes(), nil
}

// Unmarshal unmarshals YAML to a SecretsFile.
func Unmarshal(data []byte) (*SecretsFile, error) {
	var sf SecretsFile
	err := yaml.Unmarshal(data, &sf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secrets file: %w", err)
	}
	return &sf, nil
}

// NewSecretsFile creates a new SecretsFile with default values.
func NewSecretsFile() *SecretsFile {
	return &SecretsFile{
		USM: 1,
		KDF: KDF{
			Name: "scrypt",
			N:    32768,
			R:    8,
			P:    1,
		},
		Crypto: Crypto{
			Cipher:     "aes-256-gcm",
			NonceBytes: 12,
		},
		Profiles: []string{"dev"},
		Secrets:  make(Secrets),
		Metadata: Metadata{
			Created: time.Now().UTC().Format(time.RFC3339),
			Updated: time.Now().UTC().Format(time.RFC3339),
			Version: 1,
		},
	}
}