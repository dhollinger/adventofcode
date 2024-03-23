package main

import "fmt"

func part2() {
	grid := parseFile()
	el, ec := empty(grid)
	points := findPoints(grid)
	distS := make([]int, 0)
	weight := 1000000

	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			x := 0

			for k := min(p1.x, p2.x); k < max(p1.x, p2.x); k++ {
				if _, exists := el[k]; exists {
					x += weight
				} else {
					x++
				}
			}

			for k := min(p1.y, p2.y); k < max(p1.y, p2.y); k++ {
				if _, exists := ec[k]; exists {
					x += weight
				} else {
					x++
				}
			}

			distS = append(distS, x)
		}
	}

	total := 0
	for _, dist := range distS {
		total += dist
	}

	fmt.Println(total)

}

func empty(grid [][]rune) (map[int]struct{}, map[int]struct{}) {
	size := len(grid)
	ec := make(map[int]struct{})
	el := make(map[int]struct{})

	for i := 0; i < size; i++ {
		l := true
		c := true
		for j := 0; j < size; j++ {
			l = l && grid[i][j] != '#'
			c = c && grid[j][i] != '#'
		}
		if l {
			el[i] = struct{}{}
		}
		if c {
			ec[i] = struct{}{}
		}
	}

	return el, ec
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
