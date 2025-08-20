# Prerequisites

- Go 1.21+
- Node.js 18+ (for Node.js SDK and examples)
- Python 3.8+ (for Python SDK and examples)
- PHP 8.1+ (for PHP SDK and examples)

# Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/universal-secrets-manager/usm.git
   cd usm
   ```

2. Install pre-commit hooks (optional but recommended):
   ```bash
   pip install pre-commit
   pre-commit install
   ```

# Build

- CLI: `cd cli/usm && go build`
- SDKs: Each SDK has its own build process (see respective READMEs).

# Test

- CLI: `cd cli/usm && go test ./...`
- Node SDK: `cd sdks/node && npm test`
- Python SDK: `cd sdks/python && pytest`
- PHP SDK: `cd sdks/php && ./vendor/bin/phpunit`
- Go SDK: `cd sdks/go && go test ./...`

# Contributing

We welcome contributions! Please read our [CONTRIBUTING.md](./CONTRIBUTING.md) and [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md) before submitting a pull request.

By contributing, you agree to license your contributions under the terms of the MIT License.