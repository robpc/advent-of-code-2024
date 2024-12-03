package main

import (
	"fmt"
	"log"
	my "robpc/advent-of-code-2024/internal/my"
	"slices"
	"strconv"
	"strings"
)

const DAY = 1

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	split1 := my.SplitLines(input)

	var left []int
	var right []int

	for _, element := range split1 {
		s := strings.Split(element, "   ")
		l, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

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
