package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// LocateSecretsFile locates the .secrets.yml file in the current directory or its parents.
func LocateSecretsFile() (string, error) {
	return locateFile(".secrets.yml")
}

// LocateUSMKeyFile locates the .usmkey file in the current directory or its parents.
func LocateUSMKeyFile() (string, error) {
	return locateFile(".usmkey")
}

// LocateUSMTeamFile locates the .usm/team.yml file in the current directory or its parents.
func LocateUSMTeamFile() (string, error) {
	return locateFile(filepath.Join(".usm", "team.yml"))
}

// locateFile locates a file in the current directory or its parents.
func locateFile(filename string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	for {
		path := filepath.Join(dir, filename)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}

		// Move to parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// We've reached the root directory
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("could not locate %s file", filename)
}