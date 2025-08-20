# Makefile

# Default target
.PHONY: help
help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $1, $2}'

# Build targets
.PHONY: build-cli
build-cli: ## Build the CLI tool
	cd cli/usm && go build -o ../../usm .

.PHONY: build-release
build-release: ## Build release binaries for all platforms
	GOOS=linux GOARCH=amd64 go build -o usm-linux-amd64 cli/usm/main.go
	GOOS=darwin GOARCH=amd64 go build -o usm-darwin-amd64 cli/usm/main.go
	GOOS=windows GOARCH=amd64 go build -o usm-windows-amd64.exe cli/usm/main.go

.PHONY: build-sdks
build-sdks: ## Build all SDKs
	@echo "Building Node.js SDK..."
	cd sdks/node && npm run build
	@echo "Building Python SDK..."
	# Python builds are typically handled by pip, but you can add specific build steps here if needed
	@echo "Building PHP SDK..."
	# PHP builds are typically handled by Composer, but you can add specific build steps here if needed
	@echo "Building Go SDK..."
	# Go SDK is a library, no build step needed

# Test targets
.PHONY: test-cli
test-cli: ## Run CLI tests
	cd cli/usm && go test ./...

.PHONY: test-sdks
test-sdks: ## Run all SDK tests
	@echo "Testing Node.js SDK..."
	cd sdks/node && npm test
	@echo "Testing Python SDK..."
	cd sdks/python && pytest
	@echo "Testing PHP SDK..."
	cd sdks/php && ./vendor/bin/phpunit
	@echo "Testing Go SDK..."
	cd sdks/go && go test ./...

.PHONY: test
test: test-cli test-sdks ## Run all tests

# Format targets
.PHONY: format-go
format-go: ## Format Go code
	cd cli/usm && go fmt ./...
	cd core/crypto && go fmt ./...
	cd sdks/go && go fmt ./...

.PHONY: format-node
format-node: ## Format Node.js code
	cd sdks/node && npm run format

.PHONY: format-python
format-python: ## Format Python code
	cd sdks/python && black .
	cd sdks/python && isort .

.PHONY: format-php
format-php: ## Format PHP code
	cd sdks/php && ./vendor/bin/phpcbf

.PHONY: format
format: format-go format-node format-python format-php ## Format all code

# Clean target
.PHONY: clean
clean: ## Clean build artifacts
	rm -f usm
	rm -f usm.exe
	rm -f usm-*