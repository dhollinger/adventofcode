package main

import (
	"log"
	"os"
	"slices"
)

func main() {
	args := os.Args

	if slices.Contains(args, "part1") {
		part1()
	} else if slices.Contains(args, "part2") {
		part2()
	} else {
		log.Fatal("Invalid function")
	}

}
