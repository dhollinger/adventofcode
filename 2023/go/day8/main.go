package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	L = iota
	R
)

type Node struct {
	left, right string
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	split := strings.TrimSuffix(string(file), "\n")
	lines := strings.Split(split, "\n")
	directions := strings.Split(lines[0], "")

	desertMap := make(map[string]Node)

	var startNodes []string

	for _, line := range lines[2:] {
		s, _, _ := reMatch("(\\w+)", line)
		node, left, right := s[0], s[1], s[2]
		desertMap[node] = Node{left, right}
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
	}

	for i := range desertMap {
		if string(i[len(i)-1]) == "A" {
			startNodes = append(startNodes, i)
		}
	}

	stepsTaken := []int{}
	for _, currentNode := range startNodes {
		steps := findSteps(desertMap, currentNode, directions)
		stepsTaken = append(stepsTaken, steps)
	}
	fmt.Println(lcmN(stepsTaken))

}

func findSteps(net map[string]Node, startNode string, directions []string) int {
	steps := 0
	for {
		currentNode := startNode
		for _, dir := range directions {
			steps++
			node := net[currentNode]
			if dir == "L" {
				currentNode = node.left
			} else {
				currentNode = node.right
			}
			if strings.HasSuffix(currentNode, "Z") {
				return steps
			}
		}
		startNode = currentNode
	}
}

func getDir(dir string) int {
	if dir == "L" {
		return L
	}
	return R
}

func trimParenth(s string) string {
	s = strings.Trim(s, "(")
	s = strings.Trim(s, ")")
	return s
}

func reMatch(pattern string, s string) ([]string, [][]int, [][]string) {
	r, _ := regexp.Compile(pattern)
	return r.FindAllString(s, -1), r.FindAllStringIndex(s, -1), r.FindAllStringSubmatch(s, -1)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	// a * b = lcm(a, b) * gcd(a, b)
	return (a * b) / gcd(a, b)
}

func lcmN(n []int) int {
	if len(n) == 2 {
		return lcm(n[0], n[1])
	}
	return lcm(n[0], lcmN(n[1:]))
}
