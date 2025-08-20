package main

import (
	"fmt"
	"log"

	usm "github.com/universal-secrets-manager/usm/sdks/go"
)

func main() {
	usm, err := usm.Load("")
	if err != nil {
		log.Fatal(err)
	}

	dbUrl, err := usm.Get("DB_URL")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DB_URL: %s\n", dbUrl)
}
