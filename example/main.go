//go:generate wire

package main

import (
	"log"
)

func main() {
	server, err := InitServer()

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Listen())
}
