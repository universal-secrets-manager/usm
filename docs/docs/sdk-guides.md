# SDK Guides

## Node.js

```typescript
import { USM } from "@usm/secrets";

const usm = await USM.load();
const dbUrl = await usm.get("DB_URL");
```

## Python

```python
from usm import load

usm = load()
db_url = usm.get("DB_URL")
```

## PHP

```php
<?php
require_once 'vendor/autoload.php';

use Usm\USM;

$usm = USM::load();
$dbUrl = $usm->get('DB_URL');
```

## Go

```go
package main

import (
    "github.com/universal-secrets-manager/usm/sdks/go"
)

func main() {
    usm, err := usm.Load()
    if err != nil {
        panic(err)
    }
    
    dbUrl, err := usm.Get("DB_URL")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(dbUrl)
}
```