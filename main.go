package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("requires single argument containing filename of module masses to read")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("failed to open file:", os.Args[1])
	}

	modules, err := modulesFromReader(f)
	if err != nil {
		log.Fatalln("failed to read module data:", err)
	}

	var totalFuel int64
	for _, module := range modules {
		totalFuel += totalFuelForModule(module)
	}
	log.Println("total fuel:", totalFuel)
}
