package main

import (
	"bufio"
	"log"
	"os"
)

type point struct {
	x, y int
}

func main() {
	part2()
	// grid := parseFile()
	//
	// grid = addLines(grid)
	// grid = addColumns(grid)
	// points := findPoints(grid)
	//
	// dists := make([]int, 0)
	//
	// for i := 0; i < len(points); i++ {
	// 	p1 := points[i]
	// 	for j := i + 1; j < len(points); j++ {
	// 		p2 := points[j]
	// 		x := abs(p1.x-p2.x) + abs(p1.y-p2.y)
	// 		dists = append(dists, x)
	// 	}
	// }
	//
	// total := 0
	// for _, dist := range dists {
	// 	total += dist
	// }
	//
	// fmt.Println(total)
	//
}

func findPoints(grid [][]rune) []point {
	var p []point

	for i, l := range grid {
		for j, c := range l {
			if c == '#' {
				p = append(p, point{i, j})
			}
		}
	}

	return p
}

func parseFile() [][]rune {
	grid := make([][]rune, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}

func addLines(grid [][]rune) [][]rune {
	size := len(grid)
	var nm [][]rune

	for i := 0; i < size; i++ {
		r := true
		for j := 0; j < size; j++ {
			r = r && grid[i][j] != '#'
		}
		nm = append(nm, grid[i])
		if r {
			nm = append(nm, make([]rune, size))
		}
	}

	return nm
}

func addColumns(grid [][]rune) [][]rune {
	size := len(grid)
	var nm [][]rune

	for i := 0; i < size; i++ {
		nm = append(nm, make([]rune, 0, len(grid[0])))
	}

	for i := 0; i < len(grid[0]); i++ {
		c := true
		for j := 0; j < size; j++ {
			c = c && grid[j][i] != '#'
			nm[j] = append(nm[j], grid[j][i])
		}

		if c {
			for j := 0; j < size; j++ {
				nm[j] = append(nm[j], '.')
			}
		}
	}

	return nm
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
