package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	card    string
	winning []int
	numbers []int
}

type Cards []Card

func main() {
	var cards Cards
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ": ")
		split := strings.Split(card[1], " | ")
		numbers := sortNums(split[1])
		winning := sortNums(split[0])
		cardInfo := Card{
			card:    card[0],
			winning: winning,
			numbers: numbers,
		}
		cards = append(cards, cardInfo)
	}

	total := 0

	for _, val := range cards {
		points := 0
		for _, int := range val.numbers {
			if slices.Contains(val.winning, int) && points == 0 {
				points = 1
			} else if slices.Contains(val.winning, int) {
				points = points * 2
			}
		}
		total = total + points
	}

	fmt.Println(total)
}

func sortNums(numStr string) []int {
	var ints []int
	numSlice := strings.Fields(numStr)
	for i := range numSlice {
		num, err := strconv.Atoi(numSlice[i])
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, num)
	}
	sort.Ints(ints)
	return ints
}
