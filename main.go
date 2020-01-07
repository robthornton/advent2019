package main

import (
	"log"
	"os"

	"github.com/robthornton/advent2019/orbit"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("must supply a filename to load")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	m := orbit.NewFromReader(f)

	log.Println(m.FewestTransfers("YOU", "SAN"))
}
