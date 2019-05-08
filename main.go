//go:generate wire

package main

import (
	"github.com/robojones/gqltest/config"
	"log"
)

func main() {
	t, err := InitTester(config.WorkinDirectoryName(""))

	if err != nil {
		log.Fatal(err)
	}

	err = t.Run()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Tests passed")
	}
}
