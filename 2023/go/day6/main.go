package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var top int

func main() {
	file, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	time, dist := strings.Fields(lines[0])[1:], strings.Fields(lines[1])[1:]

	timeInt := convStr(time)
	distInt := convStr(dist)

	fmt.Println(timeInt)
	fmt.Println(distInt)

	multiples := 1

	for i, v := range timeInt {
		for n := 1; n < v; n++ {
			if n*(v-n) > distInt[i] {
				multiples *= (top - n)
				break
			} else {
				top = v - n
			}
		}
	}

	fmt.Println(multiples)

}

func convStr(nums []string) []int {
	var ints []int
	for _, n := range nums {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, num)
	}
	return ints
}
