# Python SDK

This directory contains the Python SDK for USM.

## Installation

```bash
pip install usm-secrets
```

## Usage

```python
from usm import load

usm = load()
db_url = usm.get("DB_URL")
```

## API

### `load(file_path=None)`

Loads a USM secrets file. If no file path is provided, it will search for a `.secrets.yml` file in the current directory and its parents.

### `USM.get(key)`

Retrieves a secret value by key. Returns the decrypted secret value.

## Development

### Installing in Development Mode

```bash
pip install -e .
```

### Running Tests

```bash
python -m pytest
```

## Example

```python
from usm import load

def main():
    try:
        usm = load()
        db_url = usm.get('DB_URL')
        print(f"DB URL: {db_url}")
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()
```