package main

import (
	"fmt"
	my "robpc/advent-of-code-2024/internal/my"
	"slices"
)

const DAY = 1

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	grid := my.IntGridFromString(input, "   ", "\n")
	left, right := my.UnzipIntGrid(grid)

	slices.Sort(left)
	slices.Sort(right)

	{
		var distance []int

		for i := range left {
			distance = append(distance, my.Abs(left[i]-right[i]))
		}

		var n int = my.Sum(distance)

		fmt.Printf("\nDay %02v: Part 1\n", DAY)
		fmt.Println(n)
	}

	{
		var similarity []int

		for i := range left {
			x := left[i]
			c := 0
			for _, e := range right {
				if e == x {
					c++
				}
			}
			similarity = append(similarity, x*c)
		}

		var n int = my.Sum(similarity)

		fmt.Printf("\nDay %02v: Part 2\n", DAY)
		fmt.Println(n)
	}
}
