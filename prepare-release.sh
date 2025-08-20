#!/bin/bash

# This script prepares the CLI for release by creating a separate go.mod file
# that doesn't use replace directives

set -e

echo "Preparing CLI for release..."

# Save the original go.mod file
cp cli/usm/go.mod cli/usm/go.mod.release.bak

# Remove the replace directive
sed -i '/replace github.com\/universal-secrets-manager\/usm\/core\/crypto => \.\.\/\.\.\/core\/crypto/d' cli/usm/go.mod

# Update the requirement to use the same version as the main module
go mod tidy

echo "CLI is ready for release. Don't forget to restore the original go.mod after building:"
echo "cp cli/usm/go.mod.release.bak cli/usm/go.mod"