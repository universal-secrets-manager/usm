#!/bin/bash

# Comprehensive test script for USM CLI with all implemented features

echo "Running comprehensive tests for USM CLI..."

# Clean up any existing files
rm -f .secrets.yml .usm.team *.private.key .usm.audit.log

# Test init command
echo "Testing init command..."
./usm.exe init
if [ ! -f .secrets.yml ]; then
  echo "FAIL: init command failed"
  exit 1
fi
echo "PASS: init command"

# Test set command
echo "Testing set command..."
./usm.exe set API_KEY=12345
if ! grep -q "API_KEY" .secrets.yml; then
  echo "FAIL: set command failed"
  exit 1
fi
echo "PASS: set command"

# Test list command
echo "Testing list command..."
output=$(./usm.exe list)
if [ "$output" != "API_KEY" ]; then
  echo "FAIL: list command failed"
  exit 1
fi
echo "PASS: list command"

# Test get command
echo "Testing get command..."
output=$(./usm.exe get API_KEY)
if [ "$output" != "12345" ]; then
  echo "FAIL: get command failed"
  exit 1
fi
echo "PASS: get command"

# Test rotate command
echo "Testing rotate command..."
./usm.exe rotate API_KEY
if ! grep -q "API_KEY" .secrets.yml; then
  echo "FAIL: rotate command failed"
  exit 1
fi
echo "PASS: rotate command"

# Verify we can still get the secret after rotating
output=$(./usm.exe get API_KEY)
if [ "$output" != "12345" ]; then
  echo "FAIL: get command failed after rotation"
  exit 1
fi
echo "PASS: get command after rotation"

# Test share command
echo "Testing share command..."
./usm.exe share add alice@example.com
if [ ! -f alice@example.com.private.key ] || [ ! -f .usm.team ]; then
  echo "FAIL: share command failed"
  exit 1
fi
echo "PASS: share command"

# Test recipients list command
echo "Testing recipients list command..."
output=$(./usm.exe recipients list)
if [ "$output" != "Recipients:
  - alice@example.com" ]; then
  echo "FAIL: recipients list command failed"
  exit 1
fi
echo "PASS: recipients list command"

# Test recipients add command
echo "Testing recipients add command..."
./usm.exe recipients add bob@example.com
if [ ! -f bob@example.com.private.key ]; then
  echo "FAIL: recipients add command failed"
  exit 1
fi
echo "PASS: recipients add command"

# Test recipients list command again
output=$(./usm.exe recipients list)
if [ "$output" != "Recipients:
  - alice@example.com
  - bob@example.com" ]; then
  echo "FAIL: recipients list command failed after adding recipient"
  exit 1
fi
echo "PASS: recipients list command after adding recipient"

# Test recipients remove command
echo "Testing recipients remove command..."
./usm.exe recipients rm alice@example.com
if [ -f alice@example.com.private.key ]; then
  echo "FAIL: recipients remove command failed"
  exit 1
fi
echo "PASS: recipients remove command"

# Test recipients list command again
output=$(./usm.exe recipients list)
if [ "$output" != "Recipients:
  - bob@example.com" ]; then
  echo "FAIL: recipients list command failed after removing recipient"
  exit 1
fi
echo "PASS: recipients list command after removing recipient"

# Test profiles list command
echo "Testing profiles list command..."
output=$(./usm.exe profiles list)
if [ "$output" != "Profiles:
  - dev" ]; then
  echo "FAIL: profiles list command failed"
  exit 1
fi
echo "PASS: profiles list command"

# Test profiles add command
echo "Testing profiles add command..."
./usm.exe profiles add staging
echo "PASS: profiles add command"

# Test profiles list command again
output=$(./usm.exe profiles list)
if [ "$output" != "Profiles:
  - dev
  - staging" ]; then
  echo "FAIL: profiles list command failed after adding profile"
  exit 1
fi
echo "PASS: profiles list command after adding profile"

# Test profiles remove command
echo "Testing profiles remove command..."
./usm.exe profiles rm staging
echo "PASS: profiles remove command"

# Test profiles list command again
output=$(./usm.exe profiles list)
if [ "$output" != "Profiles:
  - dev" ]; then
  echo "FAIL: profiles list command failed after removing profile"
  exit 1
fi
echo "PASS: profiles list command after removing profile"

# Test audit show command
echo "Testing audit show command..."
output=$(./usm.exe audit show)
if [ -z "$output" ]; then
  echo "FAIL: audit show command failed"
  exit 1
fi
echo "PASS: audit show command"

# Test delete command
echo "Testing delete command..."
./usm.exe delete API_KEY
output=$(./usm.exe list)
if [ "$output" != "" ]; then
  echo "FAIL: delete command failed"
  exit 1
fi
echo "PASS: delete command"

echo "All comprehensive tests passed!"