#!/bin/bash

# Run all tests in the core/crypto package and its subpackages

echo "Running tests for core/crypto package..."

# Run tests for the main crypto package
echo "Testing main crypto package..."
cd core/crypto && go test -v

# Run tests for each subpackage
echo "Testing aead subpackage..."
cd aead && go test -v

echo "Testing kdf subpackage..."
cd ../kdf && go test -v

echo "Testing asym subpackage..."
cd ../asym && go test -v

echo "Testing sign subpackage..."
cd ../sign && go test -v

echo "Testing securemem subpackage..."
cd ../securemem && go test -v

echo "Testing file subpackage..."
cd ../file && go test -v

echo "All tests completed!"