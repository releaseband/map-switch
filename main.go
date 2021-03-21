package main

import (
	"github.com/releaseband/map-switch/cli/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}