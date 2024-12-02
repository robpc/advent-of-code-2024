package main

import (
	"fmt"
	my "robpc/advent-of-code-2024/internal/my"
)

const DAY = 2

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	fmt.Printf("Day %02v: Part 1\n", DAY)
	fmt.Println(input)

	// fmt.Printf("Day %02v: Part 2\n", DAY)
	// fmt.Println(input)

}
