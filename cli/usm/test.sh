#!/bin/bash

# Test script for USM CLI

echo "Testing USM CLI..."

# Clean up any existing .secrets.yml file
rm -f .secrets.yml

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

# Test delete command
echo "Testing delete command..."
./usm.exe delete API_KEY
output=$(./usm.exe list)
if [ "$output" != "" ]; then
  echo "FAIL: delete command failed"
  exit 1
fi
echo "PASS: delete command"

echo "All tests passed!"