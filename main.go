//go:generate wire

package main

import (
	"flag"
	"github.com/robojones/gqltest/config"
	"log"
)

func main() {
	flag.Parse()
	testdir := flag.Arg(0)

	t, err := InitTester(config.WD(testdir))

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
