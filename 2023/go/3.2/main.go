package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Number struct {
	num    string
	x      int
	y      int
	length int
}

type Symbol struct {
	symbol string
	x      int
	y      int
	nums   []int
}

type Numbers []Number

type Symbols []Symbol

var lines []string

var nums Numbers
var syms Symbols

var ints []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := range lines {
		re := regexp.MustCompile(`\d+`)
		indices := re.FindAllIndex([]byte(lines[i]), -1)
		for _, values := range indices {
			start := values[0]
			end := values[1]

			number := Number{
				num:    string(lines[i][start:end]),
				x:      i + 1,
				y:      values[0],
				length: values[1] - 1,
			}

			nums = append(nums, number)
		}
	}

	for i := range lines {
		re := regexp.MustCompile(`\*`)
		indices := re.FindAllIndex([]byte(lines[i]), -1)
		for _, sym := range indices {
			symbol := Symbol{
				symbol: string(lines[i][sym[0]]),
				x:      i + 1,
				y:      sym[0],
			}

			syms = append(syms, symbol)
		}
	}

	for ind, val := range syms {
		for _, num := range nums {
			if num.x == val.x-1 || num.x == val.x || num.x == val.x+1 {
				if (num.y == val.y-1 || num.length == val.y-1) || (num.y == val.y || num.length == val.y) || (num.y == val.y+1 || num.length == val.y+1) {
					i, err := strconv.Atoi(num.num)
					if err != nil {
						log.Fatal(err)
					}
					syms[ind].nums = append(syms[ind].nums, i)
				}
			}
		}
	}

	for _, val := range syms {
		if len(val.nums) == 2 {
			ints = append(ints, val.nums[0]*val.nums[1])
		}
	}

	total := 0

	for i := range ints {
		total = total + ints[i]
	}

	fmt.Println(total)
}
