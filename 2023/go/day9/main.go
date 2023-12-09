package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var seqSlice []int

func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		iseq := intconv(line)

		total += prevLineNum(iseq)
	}

	fmt.Println(total)
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func intconv(line string) []int {
	var numSeq []int
	numStr := strings.Fields(line)
	for _, num := range numStr {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		numSeq = append(numSeq, n)
	}
	return numSeq
}

func lcd(src []int) int {
	res := make([]int, len(src)-1)
	hasNonZero := false
	for i := 1; i < len(src); i++ {
		res[i-1] = src[i] - src[i-1]
		if res[i] != 0 {
			hasNonZero = true
		}
	}

	if !hasNonZero {
		return src[len(src)-1]
	}

	return src[len(src)-1] + lcd(res)
}

func prevLineNum(src []int) int {
	res := make([]int, len(src)-1)
	hasNonZero := false
	for i := 1; i < len(src); i++ {
		res[i-1] = src[i] - src[i-1]
		if res[i-1] != 0 {
			hasNonZero = true
		}
	}

	if !hasNonZero {
		return src[0]
	}

	return src[0] - prevLineNum(res)

}
