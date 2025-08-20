package kdf

import (
	"golang.org/x/crypto/scrypt"
)

// ScryptParams holds the parameters for scrypt.
type ScryptParams struct {
	N int
	R int
	P int
}

// DefaultParams returns the default scrypt parameters.
func DefaultParams() ScryptParams {
	return ScryptParams{
		N: 32768,
		R: 8,
		P: 1,
	}
}

// DeriveKey derives a key from a passphrase and salt using scrypt.
func DeriveKey(passphrase, salt []byte, params ScryptParams) ([]byte, error) {
	return scrypt.Key(passphrase, salt, params.N, params.R, params.P, 32)
}