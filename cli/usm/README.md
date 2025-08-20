# USM CLI Tool

The USM CLI tool is a command-line interface for managing encrypted secrets.

## Installation

To install the USM CLI tool, you can build it from source:

```bash
cd cli/usm
go build -o usm .
```

## Usage

### Initialize a new secrets file

```bash
usm init
```

This creates a new `.secrets.yml` file with default settings.

### Set a secret

```bash
usm set KEY=VALUE
```

This sets a secret value in the secrets file.

### Get a secret

```bash
usm get KEY
```

This retrieves a secret value from the secrets file.

### List all secrets

```bash
usm list
```

This lists all secret keys in the secrets file.

### Delete a secret

```bash
usm delete KEY
```

This deletes a secret from the secrets file.

## Commands

- `init` - Initialize a new secrets file
- `set` - Set a secret
- `get` - Get a secret
- `list` - List all secrets
- `delete` - Delete a secret
- `rotate` - Rotate a secret or key
- `share` - Share secrets with a user
- `recipients` - Manage recipients
- `profiles` - Manage profiles
- `sync` - Synchronize with cloud providers
- `audit` - Show audit log

## Building

To build the CLI tool, run:

```bash
go build -o usm .
```

## Testing

To run the tests, run:

```bash
./test.sh
```