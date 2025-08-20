# Go SDK

This directory contains the Go SDK for USM.

## Installation

```bash
go get github.com/universal-secrets-manager/usm/sdks/go
```

## Usage

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