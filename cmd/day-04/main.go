package main

import (
	"fmt"
	my "robpc/advent-of-code-2024/internal/my"
	"strings"
)

const DAY = 4

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	lines := strings.Split(input, "\n")

	var grid Grid

	for _, l := range lines {
		grid = append(grid, []byte(l))
	}

	{
		n := grid.Count("XMAS", false)

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		n := grid.Count("MAS", true)

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}

type Grid [][]byte

func (g Grid) Count(word string, isPartTwo bool) int {
	n := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if isPartTwo {
				if i > 0 && i < len(g)-1 && j > 0 && j < len(g[i])-1 {
					x := g.countThreeLetterFromMiddle(i, j, word)
					if x == 2 {
						n++
						// fmt.Print("\n@", i, j, "\n", string(g[i-1][j-1:j+2]), "\n", string(g[i][j-1:j+2]), "\n", string(g[i+1][j-1:j+2]), "\n-")
					}
				}
			} else {
				n += g.countFrom(i, j, word)
			}
		}
	}
	return n
}

func (g Grid) countFrom(row, col int, word string) int {
	// Check all 8 directions
	n := 0
	for _, dir := range []struct{ dx, dy int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		if g.searchInDirection(row, col, word, dir.dx, dir.dy) {
			n++
		}
	}
	return n
}

func (g Grid) countThreeLetterFromMiddle(row, col int, word string) int {
	n := 0
	if g.searchInDirection(row+1, col-1, word, -1, 1) {
		n++
	}
	if g.searchInDirection(row-1, col+1, word, 1, -1) {
		n++
	}
	if g.searchInDirection(row-1, col-1, word, 1, 1) {
		n++
	}
	if g.searchInDirection(row+1, col+1, word, -1, -1) {
		n++
	}
	return n
}

func (g Grid) Search(word string) bool {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g.searchFrom(i, j, word) {
				return true
			}
		}
	}
	return false
}

func (g Grid) searchFrom(row, col int, word string) bool {
	// Check all 8 directions
	for _, dir := range []struct{ dx, dy int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		if g.searchInDirection(row, col, word, dir.dx, dir.dy) {
			return true
		}
	}
	return false
}

func (g Grid) searchInDirection(row, col int, word string, dx, dy int) bool {
	for i := 0; i < len(word); i++ {
		if row < 0 || row >= len(g) || col < 0 || col >= len(g[row]) || g[row][col] != word[i] {
			return false
		}
		row += dx
		col += dy
	}
	return true
}
