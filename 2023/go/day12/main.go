package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	parseFile()
}

func parseFile() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var diagSym []string
	var diagInt []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed := strings.Fields(scanner.Text())
		diagSym = append(diagSym, parsed[0])
		diagInt = append(diagInt, parsed[1])
	}

	var diag [][]int

	for _, i := range diagInt {
		splitInt := strings.Split(i, ",")
		for _, j := range splitInt {
		}
	}

	fmt.Println(diagSym)
	fmt.Println(diagInt)
}
