package main

import (
	"errors"
	"fmt"
	"robpc/advent-of-code-2024/internal/my"
	"strings"
)

const DAY = 6

type Pos [2]int

type Grid [][]string

func (g Grid) Set(p Pos, s string) {
	g[p[0]][p[1]] = s
}
func (g Grid) Copy() Grid {
	var next Grid
	for _, row := range g {
		r := make([]string, len(row))
		copy(r, row)
		next = append(next, r)
	}
	return next
}
func (g Grid) String() string {
	var s string
	for _, l := range g {
		s += fmt.Sprintln(strings.Join(l, ""))
	}
	return s
}

func (p1 Pos) Add(p2 Pos) Pos {
	return Pos{p1[0] + p2[0], p1[1] + p2[1]}
}

func (p Pos) Cell(g [][]string) (string, error) {
	if p[0] < 0 || p[0] >= len(g) || p[1] < 0 || p[1] >= len(g[0]) {
		return "", errors.New("out of bounds")
	}
	return g[p[0]][p[1]], nil
}

func patrol(grid Grid) (int, bool) {
	var pos = Pos{}

	for i, row := range grid {
		for j, cell := range row {
			if strings.Contains("^>V<", cell) {
				pos = Pos{i, j}
			}
		}
	}

	n := 0

	hasTurned := false
	dupeTurns := 0
	for {
		dir := grid[pos[0]][pos[1]]

		mov := Pos{}
		if dir == "^" {
			mov = Pos{-1, 0}
		} else if dir == ">" {
			mov = Pos{0, 1}
		} else if dir == "V" {
			mov = Pos{1, 0}
		} else if dir == "<" {
			mov = Pos{0, -1}
		}
		dest := pos.Add(mov)

		next, err := dest.Cell(grid)

		if err != nil {
			grid.Set(pos, "X")
			n++
			break
		}

		if next == "+" {
			nextNext, _ := dest.Add(mov).Cell(grid)
			if strings.Contains("#O", nextNext) {
				dupeTurns++
			}
			if dupeTurns >= 4 {
				break
			}
		}

		if strings.Contains("#O", next) {
			if dir == "^" {
				grid.Set(pos, ">")
			} else if dir == ">" {
				grid.Set(pos, "V")
			} else if dir == "V" {
				grid.Set(pos, "<")
			} else if dir == "<" {
				grid.Set(pos, "^")
			}
			hasTurned = true
		} else {
			if hasTurned {
				grid.Set(pos, "+")
			} else if strings.Contains("><", dir) {
				grid.Set(pos, "-")
			} else if strings.Contains("^V", dir) {
				grid.Set(pos, "|")
			}
			grid.Set(dest, dir)
			pos = dest
			if next == "." {
				n++
			}
			hasTurned = false
		}

		// fmt.Print(grid, "\n")
	}

	return n, dupeTurns >= 4
}

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	var grid Grid = my.StringGrid(input, "", "\n")

	{
		myGrid := grid.Copy()

		n, _ := patrol(myGrid)

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		n := 0

		for i, row := range grid {
			for j, cell := range row {
				if cell == "." {
					g := grid.Copy()
					g.Set(Pos{i, j}, "O")
					_, loops := patrol(g)
					if loops {
						// fmt.Println(g)
						n++
					}
				}
			}
		}

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}
