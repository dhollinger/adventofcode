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
	reds      = 12
	greens    = 13
	blues     = 14
)

var gameList []string
var gameBool []bool

func main() {
	gameSum := 0
	var impossible bool
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split1 := strings.Split(scanner.Text(), ":")
		game := split1[0]
		gameNum := gameNumInt(strings.Split(game, " ")[1])
		rounds := strings.Split(split1[1], ";")
		for i := range rounds {
			colors := strings.Split(rounds[i], ",")
			impossible = gamePossible(colors)
			if impossible {
				break
			}
		}
		if !impossible {
			gameList = append(gameList, game)
			gameSum = gameSum + gameNum
		}
	}

	fmt.Println(gameSum)
	fmt.Println(gameList)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkColor(color string) bool {
	var compare int
	color = strings.TrimSpace(color)
	data := strings.Split(color, " ")
	num, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	switch data[1] {
	case "red":
		compare = 12
	case "green":
		compare = 13
	case "blue":
		compare = 14
	}
	if num > compare {
		return true
	} else {
		return false
	}
}

func gameNumInt(num string) int {
	numi, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return numi
}

func gamePossible(colors []string) bool {
	for i := range colors {
		if checkColor(colors[i]) {
			return true
		}
	}
	return false
}
