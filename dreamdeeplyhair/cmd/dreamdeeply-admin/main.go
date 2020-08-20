package main

import (
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error %s", err.Error())
		os.Exit(1)
	}
}

func run() error {
	return nil
}
