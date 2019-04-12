//go:generate wire

package main

import "log"

func main() {
	t := InitTester()

	err := t.Run()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Tests passed")
	}
}
