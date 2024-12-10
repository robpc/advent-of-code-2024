package main

import (
	"errors"
	"flag"
	"fmt"
	"robpc/advent-of-code-2024/internal/my"
	"slices"
	"strconv"
	"strings"
)

const DAY = 10

type Pos [2]int

func (p Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p[0], p[1])
}

func (p1 Pos) Add(p2 Pos) Pos {
	return Pos{p1[0] + p2[0], p1[1] + p2[1]}
}

func (p1 Pos) Sub(p2 Pos) Pos {
	return Pos{p1[0] - p2[0], p1[1] + -p2[1]}
}

func (p Pos) Copy() Pos {
	return Pos{p[0], p[1]}
}

var CardinalDirections = []Pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Grid [][]string

func (g Grid) Set(p Pos, s string) error {
	if p[0] < 0 || p[0] >= len(g) || p[1] < 0 || p[1] >= len(g[0]) {
		return errors.New("out of bounds")
	}
	g[p[0]][p[1]] = s

	return nil
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
func (g Grid) Cell(p Pos) (string, error) {
	if p[0] < 0 || p[0] >= len(g) || p[1] < 0 || p[1] >= len(g[0]) {
		return "", errors.New("out of bounds")
	}
	return g[p[0]][p[1]], nil
}

func GetTrailheadScore(grid Grid, pos Pos) []Pos {
	cell, _ := grid.Cell(pos)
	height, _ := strconv.Atoi(cell)

	tops := []Pos{}
	for _, dir := range CardinalDirections {
		p := pos.Add(dir)
		c, err := grid.Cell(p)

		if err != nil {
			continue
		}

		h, _ := strconv.Atoi(c)

		if height == 9 {
			return []Pos{pos}
		}

		if h == height+1 {
			subs := GetTrailheadScore(grid, p)
			for _, s := range subs {
				if !slices.Contains(tops, s) {
					tops = append(tops, s)
				}
			}
		}
	}
	return tops
}

func GetTrailheadRating(grid Grid, pos Pos) []Pos {
	cell, _ := grid.Cell(pos)
	height, _ := strconv.Atoi(cell)

	tops := []Pos{}
	for _, dir := range CardinalDirections {
		p := pos.Add(dir)
		c, err := grid.Cell(p)

		if err != nil {
			continue
		}

		h, _ := strconv.Atoi(c)

		if height == 9 {
			return []Pos{pos}
		}

		if h == height+1 {
			subs := GetTrailheadRating(grid, p)
			tops = append(tops, subs...)
		}
	}
	return tops
}

type Path []Pos

func FindUphillPaths(grid Grid, pos Pos, trail Path) []Path {
	cell, _ := grid.Cell(pos)
	height, _ := strconv.Atoi(cell)

	group := []Path{}
	for _, dir := range CardinalDirections {
		p := pos.Add(dir)
		c, err := grid.Cell(p)

		if err != nil {
			continue
		}

		h, _ := strconv.Atoi(c)

		if height == 9 {
			t := append(trail, p)
			return []Path{t}
		}

		if h == height+1 {
			t := append(trail, p)
			subs := FindUphillPaths(grid, p, t)
			group = append(group, subs...)
		}
	}
	return group
}

func main() {
	var puzzle = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/puzzle.txt", DAY))
	var example = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/example.txt", DAY))

	useExampleInput := flag.Bool("example", false, "use example input")
	flag.Parse()

	var input string
	if *useExampleInput {
		fmt.Println("Using example input")
		input = example
	} else {
		fmt.Println("Using puzzle input")
		input = puzzle
	}

	var grid Grid = my.StringGrid(input, "", "\n")

	{
		n := 0

		for i, row := range grid {
			for j, cell := range row {
				if cell == "0" {
					tops := GetTrailheadScore(grid, Pos{i, j})

					fmt.Print(Pos{i, j}, cell, "=", tops)

					n += len(tops)
				}
			}
		}

		fmt.Println(grid)

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		n := 0
		for i, row := range grid {
			for j, cell := range row {
				if cell == "0" {
					tops := GetTrailheadRating(grid, Pos{i, j})

					fmt.Println(Pos{i, j}, cell, "=", tops)

					n += len(tops)
				}
			}
		}

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}
