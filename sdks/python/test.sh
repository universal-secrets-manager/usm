#!/bin/bash

# Test script for Python SDK

echo "Testing Python SDK..."

# Test installation
echo "Testing installation..."
pip install -e .
if [ $? -ne 0 ]; then
  echo "FAIL: installation failed"
  exit 1
fi
echo "PASS: installation"

# Test tests
echo "Testing tests..."
python -m pytest
if [ $? -ne 0 ]; then
  echo "FAIL: tests failed"
  exit 1
fi
echo "PASS: tests"

# Test example
echo "Testing example..."
python example.py
if [ $? -ne 0 ]; then
  echo "FAIL: example failed"
  exit 1
fi
echo "PASS: example"

echo "All tests passed!"