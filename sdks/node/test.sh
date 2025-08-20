#!/bin/bash

# Test script for Node.js SDK

echo "Testing Node.js SDK..."

# Test build
echo "Testing build..."
npm run build
if [ $? -ne 0 ]; then
  echo "FAIL: build failed"
  exit 1
fi
echo "PASS: build"

# Test tests
echo "Testing tests..."
npm test
if [ $? -ne 0 ]; then
  echo "FAIL: tests failed"
  exit 1
fi
echo "PASS: tests"

# Test example
echo "Testing example..."
cd test/fixtures && node ../../example.js
if [ $? -ne 0 ]; then
  echo "FAIL: example failed"
  exit 1
fi
echo "PASS: example"

echo "All tests passed!"