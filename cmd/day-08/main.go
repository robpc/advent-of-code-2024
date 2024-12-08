package main

import (
	"errors"
	"fmt"
	"robpc/advent-of-code-2024/internal/my"
	"strings"
)

const DAY = 8

type Pos [2]int

func (p1 Pos) Add(p2 Pos) Pos {
	return Pos{p1[0] + p2[0], p1[1] + p2[1]}
}

func (p1 Pos) Sub(p2 Pos) Pos {
	return Pos{p1[0] - p2[0], p1[1] + -p2[1]}
}

func (p Pos) Copy() Pos {
	return Pos{p[0], p[1]}
}

func (p Pos) Cell(g [][]string) (string, error) {
	if p[0] < 0 || p[0] >= len(g) || p[1] < 0 || p[1] >= len(g[0]) {
		return "", errors.New("out of bounds")
	}
	return g[p[0]][p[1]], nil
}

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

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	attennaCodes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var grid Grid = my.StringGrid(input, "", "\n")

	antenna := make(map[string][]Pos)

	for i, row := range grid {
		for j, cell := range row {
			if strings.Contains(attennaCodes, cell) {
				if antenna[cell] == nil {
					antenna[cell] = []Pos{}
				}

				antenna[cell] = append(antenna[cell], Pos{i, j})
			}
		}
	}

	{
		n := 0

		antinodes := grid.Copy()

		for a, arr := range antenna {
			for i, p1 := range arr {
				if len(arr)-1 > i {
					rest := arr[i+1:]
					fmt.Println(a, i, p1, rest)
					for _, p2 := range rest {
						an1 := p1.Add(p1.Sub(p2))
						an2 := p2.Add(p2.Sub(p1))
						antinodes.Set(an1, "#")
						antinodes.Set(an2, "#")
						fmt.Println("-", p1, p2, an1, an2)
					}
				}
			}
		}

		for _, row := range antinodes {
			for _, cell := range row {
				if cell == "#" {
					n++
				}
			}
		}

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(grid)
		fmt.Println(antinodes)
		fmt.Println(n)
	}

	{
		n := 0

		antinodes := grid.Copy()

		for a, arr := range antenna {
			for i, p1 := range arr {
				if len(arr)-1 > i {
					rest := arr[i+1:]
					fmt.Println(a, i, p1, rest)
					for _, p2 := range rest {
						// an1 := p1.Add(p1.Sub(p2))
						// an2 := p2.Add(p2.Sub(p1))
						// antinodes.Set(an1, "#")
						// antinodes.Set(an2, "#")
						// fmt.Println("-", p1, p2, an1, an2)
						d1 := p1.Sub(p2)
						an1 := p1.Copy()
						for {
							an1 = an1.Add(d1)
							err := antinodes.Set(an1, "#")
							fmt.Println("an1", an1, err)
							if err != nil {
								break
							}
						}
						d2 := p2.Sub(p1)
						an2 := p2.Copy()
						for {
							an2 = an2.Add(d2)
							err := antinodes.Set(an2, "#")
							fmt.Println("an2", an2, err)
							if err != nil {
								break
							}
						}
					}
				}
			}
		}

		for _, row := range antinodes {
			for _, cell := range row {
				if cell != "." {
					n++
				}
			}
		}

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(grid)
		fmt.Println(antinodes)
		fmt.Println(n)
	}

}
