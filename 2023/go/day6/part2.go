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
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	time, dist := strings.Fields(lines[0])[1:], strings.Fields(lines[1])[1:]

	tm := concatNum(time)
	dt := concatNum(dist)

	timeInt, err := strconv.Atoi(tm)
	if err != nil {
		log.Fatal(err)
	}
	distInt, err := strconv.Atoi(dt)
	if err != nil {
		log.Fatal(err)
	}

	multiples := 1

	for n := 1; n < timeInt; n++ {
		if n*(timeInt-n) > distInt {
			multiples *= (top - n)
			break
		} else {
			top = timeInt - n
		}
	}

	fmt.Println(multiples)

}

func concatNum(nums []string) string {
	str := ""
	for _, v := range nums {
		str += v
	}
	return str
}
