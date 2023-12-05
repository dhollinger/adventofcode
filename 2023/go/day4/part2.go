package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func part2() {
	lines := countLines("input.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var played []int

	for i := 0; i < lines; i++ {
		played = append(played, 0)
	}

	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		played[index] += 1

		game := strings.Split(scanner.Text(), ": ")[1]

		winning := strings.Split(game, " | ")[0]

		ours := strings.Split(game, " | ")[1]

		winnum := numSlice(winning)
		ournum := numSlice(ours)

		intersection := sliceutils.Intersection(winnum, ournum)

		for w := range intersection {
			played[index+w+1] += played[index]
		}

		index++
	}
	total := 0
	for i := range played {
		total += played[i]
	}
	fmt.Println(total)
}

func countLines(file string) int {
	lines := 0
	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	scan := bufio.NewScanner(content)
	for scan.Scan() {
		lines++
	}
	return lines
}

func numSlice(nss string) []int {
	numstrsl := strings.Fields(nss)

	var numintsl []int

	for _, n := range numstrsl {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numintsl = append(numintsl, num)
	}

	return numintsl
}
