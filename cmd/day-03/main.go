package main

import (
	"fmt"
	"regexp"
	my "robpc/advent-of-code-2024/internal/my"
	"strconv"
)

const DAY = 3

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	{
		rex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := rex.FindAllStringSubmatch(input, -1)

		n := 0
		for _, m := range matches {
			// match[str, group1a, group1a]
			x1, _ := strconv.Atoi(m[1])
			x2, _ := strconv.Atoi(m[2])

			n += x1 * x2
		}

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		rex := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)
		matches := rex.FindAllStringSubmatch(input, -1)

		do := true

		n := 0
		for _, s := range matches {
			if s[0] == "do()" {
				do = true
			} else if s[0] == "don't()" {
				do = false
			} else if do {
				// match[str, group1, group2a, group2a]
				x1, _ := strconv.Atoi(s[2])
				x2, _ := strconv.Atoi(s[3])

				n += x1 * x2
			}
		}

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}
