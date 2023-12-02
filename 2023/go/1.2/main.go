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

var (
	numbers = map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fileLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, replaceString(scanner.Text()))
	}

	re := regexp.MustCompile(`\d`)
	var calibrationPairs []string
	for i := range fileLines {
		values := re.FindAllString(fileLines[i], -1)
		calibrationPairs = append(calibrationPairs, values[0]+values[len(values)-1])
	}

	var calibrationInts []int
	for i := range calibrationPairs {
		number, err := strconv.Atoi(calibrationPairs[i])
		if err != nil {
			log.Fatal(err)
		}
		calibrationInts = append(calibrationInts, number)
	}

	calibrationNumber := 0
	for i := range calibrationInts {
		calibrationNumber = calibrationNumber + calibrationInts[i]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(calibrationNumber)
}

func replaceString(line string) string {
	line2 := line
	for k, v := range numbers {
		re := regexp.MustCompile(k)
		line2 = string(re.ReplaceAll([]byte(line2), []byte(v)))
	}
	return line2
}
