# Contributing to Universal Secrets Manager (USM)

First off, thank you for considering contributing to USM! It's people like you that make USM such a great tool.

## Code of Conduct

This project and everyone participating in it is governed by the [USM Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [INSERT CONTACT METHOD].

## How Can I Contribute?

### Reporting Bugs

Bugs are tracked as [GitHub issues](https://github.com/universal-secrets-manager/usm/issues). Before creating a bug report, please check the existing issues to avoid duplicates.

When you are creating a bug report, please include as many details as possible:

*   **Use a clear and descriptive title** for the issue to identify the problem.
*   **Describe the exact steps which reproduce the problem** in as many details as possible.
*   **Provide specific examples to demonstrate the steps**.
*   **Describe the behavior you observed after following the steps** and point out what exactly is the problem with that behavior.
*   **Explain which behavior you expected to see instead and why.**
*   **Include details about your configuration and environment**.

### Suggesting Enhancements

Feature requests are welcome. But take a moment to find out whether your idea fits with the scope and aims of the project. It's up to you to make a strong case to convince the project's developers of the merits of this feature.

When you are creating an enhancement suggestion, please include as many details as possible:

*   **Use a clear and descriptive title** for the issue to identify the suggestion.
*   **Provide a step-by-step description of the suggested enhancement** in as many details as possible.
*   **Provide specific examples to demonstrate the steps**.
*   **Describe the current behavior and explain which behavior you expected to see instead** and why.
*   **Explain why this enhancement would be useful** to most USM users.

### Pull Requests

*   Fill in [the required template](.github/PULL_REQUEST_TEMPLATE.md)
*   Do not include issue numbers in the PR title
*   Include screenshots and animated GIFs in your pull request whenever possible.
*   Follow the [Go](https://golang.org/doc/effective_go.html), [TypeScript](https://google.github.io/styleguide/tsguide.html), [Python (PEP 8)](https://pep8.org/), and [PHP](https://www.php-fig.org/psr/) styleguides.
*   Include thoughtfully-worded, well-structured tests.
*   Document new code based on existing documentation style.
*   End all files with a newline.

## Development Setup

See [Developer Guide](./docs/developer-guide.md) for instructions on setting up the development environment.

## Styleguides

### Git Commit Messages

*   Use the present tense ("Add feature" not "Added feature")
*   Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
*   Limit the first line to 72 characters or less
*   Reference issues and pull requests liberally after the first line
*   When only changing documentation, include `[ci skip]` in the commit title

### Go Styleguide

All Go code must adhere to [Effective Go](https://golang.org/doc/effective_go.html) and pass `golint` and `go vet`.

### TypeScript Styleguide

All TypeScript code must adhere to [Google TypeScript Style Guide](https://google.github.io/styleguide/tsguide.html).

### Python Styleguide

All Python code must adhere to [PEP 8](https://pep8.org/).

### PHP Styleguide

All PHP code must adhere to [PSR-12](https://www.php-fig.org/psr/psr-12/).

## Additional Notes

### Issue and Pull Request Labels

This section lists the labels we use to help us track and manage issues and pull requests.

*   `bug` - Issues that are bugs.
*   `enhancement` - Issues that are feature requests.
*   `documentation` - Issues or pull requests related to documentation.
*   `good first issue` - Issues that are good for newcomers.
*   `help wanted` - Issues that need assistance.