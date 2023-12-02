package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile = "input.txt"
)

var calibrationPairs []string
var calibrationInts []int

func main() {
	calibrationNumber := 0
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := re.FindAllString(scanner.Text(), -1)
		calibrationPairs = append(calibrationPairs, values[0]+values[len(values)-1])
	}

	for i := range calibrationPairs {
		number, err := strconv.Atoi(calibrationPairs[i])
		if err != nil {
			log.Fatal(err)
		}
		calibrationInts = append(calibrationInts, number)
	}

	for i := range calibrationInts {
		calibrationNumber = calibrationNumber + calibrationInts[i]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(calibrationNumber)
}
