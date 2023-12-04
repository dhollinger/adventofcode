package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

func main() {
	powerSum := 0
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maxColor := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		split1 := strings.Split(scanner.Text(), ":")
		// game := split1[0]
		// gameNum := gameNumInt(strings.Split(game, " ")[1])
		rounds := strings.Split(split1[1], ";")
		for i := range rounds {
			colors := strings.Split(rounds[i], ",")
			for i := range colors {
				maxColor = checkColor(colors[i], maxColor)
			}
		}

		gamePower := maxColor["red"] * maxColor["green"] * maxColor["blue"]
		powerSum = powerSum + gamePower

		fmt.Println(powerSum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkColor(color string, maxVal map[string]int) map[string]int {
	color = strings.TrimSpace(color)
	data := strings.Split(color, " ")
	num, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	switch data[1] {
	case "red":
		if num > maxVal["red"] {
			maxVal["red"] = num
			return maxVal
		}
	case "green":
		if num > maxVal["green"] {
			maxVal["green"] = num
			return maxVal
		}
	case "blue":
		if num > maxVal["blue"] {
			maxVal["blue"] = num
			return maxVal
		}
	}
	return maxVal
}

func gameNumInt(num string) int {
	numi, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return numi
}
