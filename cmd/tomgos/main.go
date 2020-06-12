package main

import (
	"log"

	"github.com/ridhoperdana/tomgos/cmd"
)

func main() {
	if err := cmd.RootCMD.Execute(); err != nil {
		log.Fatal("Fail init tomgos with error: ", err)
	}
}
