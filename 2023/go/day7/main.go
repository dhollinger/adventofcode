package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	HIGHCARD = iota
	PAIR
	TWOPAIRS
	THREEOFAKIND
	FULLHOUSE
	FOUROFAKIND
	FIVEOFAKIND
)

type hand string

func index(c byte) int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	panic("invalid card")
}

func (h hand) kind() int {
	var multi = make([]int, 15)
	for i := 0; i < len(h); i++ {
		multi[index(h[i])]++
	}

	var indexJ = index('J')
	var nbJoker = multi[indexJ]
	multi[indexJ] = 0

	slices.SortFunc(multi, func(a, b int) int { return b - a })

	multi[0] += nbJoker

	if multi[0] == 5 {
		return FIVEOFAKIND
	}
	if multi[0] == 4 {
		return FOUROFAKIND
	}
	if multi[0] == 3 && multi[1] == 2 {
		return FULLHOUSE
	}
	if multi[0] == 3 {
		return THREEOFAKIND
	}
	if multi[0] == 2 && multi[1] == 2 {
		return TWOPAIRS
	}
	if multi[0] == 2 {
		return PAIR
	}
	return HIGHCARD
}

func compareString(s1, s2 hand) int {
	if s1 == s2 {
		return 0
	}

	var e1, e2 string
	for i := 0; i < len(s1); i++ {
		if s1[i] == 'J' {
			e1 = e1 + "A"
		} else {
			e1 = e1 + string('A'+uint8(index(s1[i])))
		}

		if s2[i] == 'J' {
			e2 = e2 + "A"
		} else {
			e2 = e2 + string('A'+uint8(index(s2[i])))
		}
	}

	if e1 > e2 {
		return 1
	}

	return -1
}

func compare(h1, h2 hand) int {
	if h1 == h2 {
		return 0
	}

	var k1 = h1.kind()
	var k2 = h2.kind()
	if k1 > k2 {
		return 1
	}
	if k1 < k2 {
		return -1
	}

	return compareString(h1, h2)
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSuffix(string(file), "\n")

	lines := strings.Split(input, "\n")

	var hands []hand
	var bid = make(map[hand]int)
	for _, line := range lines {
		s := strings.Fields(line)
		hands = append(hands, hand(s[0]))
		v, _ := strconv.Atoi(s[1])
		bid[hand(s[0])] = v
	}

	slices.SortFunc(hands, func(a, b hand) int { return compare(a, b) })

	var result int
	for i, h := range hands {
		result += bid[h] * (i + 1)
	}

	fmt.Println(result)
}
