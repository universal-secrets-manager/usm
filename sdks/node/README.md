# Node.js SDK

This directory contains the Node.js SDK for USM.

## Installation

```bash
npm install @usm/secrets
```

## Usage

```typescript
import { USM } from "@usm/secrets";

const usm = await USM.load();
const dbUrl = await usm.get("DB_URL");
```

## API

### `load(filePath?: string): Promise<USM>`

Loads a USM secrets file. If no file path is provided, it will search for a `.secrets.yml` file in the current directory and its parents.

### `USM.get(key: string): Promise<string>`

Retrieves a secret value by key. Returns a promise that resolves to the decrypted secret value.

## Development

### Building

```bash
npm run build
```

### Testing

```bash
npm test
```

### Formatting

```bash
npm run format
```

## Example

```javascript
const { load } = require('@usm/secrets');

async function main() {
  try {
    const usm = await load();
    const dbUrl = await usm.get('DB_URL');
    console.log(`DB URL: ${dbUrl}`);
  } catch (error) {
    console.error('Error:', error.message);
  }
}

main();
```