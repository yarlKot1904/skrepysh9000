package main

import (
	"log"
	"skrepysh-agent/internal"
)

func main() {
	if err := internal.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
