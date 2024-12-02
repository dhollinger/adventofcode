package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	leftList  []int
	rightList []int
	diffList  []int
	simList   []int
)

func main() {
	readFile, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), "   ")
		leftList = append(leftList, convert(line[0]))
		rightList = append(rightList, convert(line[1]))
	}

	readFile.Close()

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := 0; i < len(leftList); i++ {
		var distance int
		if leftList[i] < rightList[i] {
			distance = rightList[i] - leftList[i]
		} else {
			distance = leftList[i] - rightList[i]
		}
		diffList = append(diffList, distance)
	}

	totalDist := 0

	for _, num := range diffList {
		totalDist = totalDist + num
	}

	fmt.Printf("Total Distance: %d\n", totalDist)

	// Part Two

	count := 0

	for _, num := range leftList {
		for _, num2 := range rightList {
			if num == num2 {
				count++
			}
		}
		simNum := num * count
		simList = append(simList, simNum)
		count = 0
	}

	totalSim := 0
	for _, num := range simList {
		totalSim = totalSim + num
	}

	fmt.Printf("Total Similarity Value: %d", totalSim)
}

func convert(value string) int {
	num, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		panic(err)
	}

	return num
}
