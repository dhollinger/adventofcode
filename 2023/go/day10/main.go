package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	row int
	col int
}

func main() {
	file, err := os.Open("sample2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]rune, 0)
	var lines []string
	var starting Position

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for row, line := range lines {
		grid = append(grid, make([]rune, 0))
		for col, c := range line {
			grid[row] = append(grid[row], c)

			if c == 'S' {
				starting.row = row
				starting.col = col
			}
		}
	}

	grid[starting.row][starting.col] = findConn(grid, starting)

	res := pt2(grid, starting)

	fmt.Println(res)

}

func p1(grid [][]rune, start Position) int {
	dist := 0
	enteredFrom := 'E'
	loopTiles := make(map[[2]int]bool)
	cur := start

	for {
		dist++
		loopTiles[[2]int{cur.row, cur.col}] = true

		switch grid[cur.row][cur.col] {
		case '|':
			if enteredFrom == 'N' {
				cur.row += 1
			} else {
				cur.row -= 1
				enteredFrom = 'S'
			}
			break
		case '-':
			if enteredFrom == 'E' {
				cur.col -= 1
			} else {
				cur.col += 1
				enteredFrom = 'W'
			}
			break
		case '7':
			if enteredFrom == 'S' {
				cur.col -= 1
				enteredFrom = 'E'
			} else {
				cur.row += 1
				enteredFrom = 'N'
			}
			break
		case 'F':
			if enteredFrom == 'S' {
				cur.col += 1
				enteredFrom = 'W'
			} else {
				cur.row += 1
				enteredFrom = 'N'
			}
			break
		case 'J':
			if enteredFrom == 'N' {
				cur.col -= 1
				enteredFrom = 'E'
			} else {
				cur.row -= 1
				enteredFrom = 'S'
			}
			break
		case 'L':
			if enteredFrom == 'N' {
				cur.col += 1
				enteredFrom = 'W'
			} else {
				cur.row -= 1
				enteredFrom = 'S'
			}
			break
		}

		if cur.row == start.row && cur.col == start.col {
			break
		}
	}

	for r, _ := range grid {
		for c, _ := range grid[r] {
			if _, ok := loopTiles[[2]int{r, c}]; !ok {
				grid[r][c] = '.'
			}
		}
	}

	return dist
}

func findConn(grid [][]rune, p Position) rune {
	connections := ""
	var N, E, S, W rune
	if p.row > 0 {
		N = grid[p.row-1][p.col]
	}
	S = grid[p.row+1][p.col]
	E = grid[p.row][p.col+1]
	if p.col > 0 {
		W = grid[p.row][p.col-1]
	}

	if N == '|' || N == '7' || N == 'F' {
		connections += "N"
	}
	if E == '7' || E == '-' || E == 'J' {
		connections += "E"
	}
	if S == '|' || S == 'L' || S == 'J' {
		connections += "S"
	}
	if W == '-' || W == 'F' || W == 'L' {
		connections += "W"
	}

	m := map[string]rune{
		"NW": 'J',
		"NE": 'L',
		"NS": '|',
		"ES": 'F',
		"SW": '7',
		"EW": '-',
	}

	return m[connections]
}
