package main

import (
	"fmt"
	"log"
	my "robpc/advent-of-code-2024/internal/my"
	"strconv"
	"strings"
)

const DAY = 2

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x > 0 {
		return 1
	}
	return -1
}

func isSafe(l []int) bool {
	var d []int
	for idx, n := range l {
		if idx == 0 {
			continue
		}
		p := l[idx-1]
		d = append(d, n-p)
	}

	var safe = false
	for i, x := range d {
		if i == 0 {
			safe = my.Abs(x) >= 1 && my.Abs(x) <= 3
		} else {
			p := d[i-1]
			safe = safe && my.Abs(x) >= 1 && my.Abs(x) <= 3 && sign(x) == sign(p)
		}
	}
	return safe
}

func removeOne(arr []int, n int) []int {
	var r []int
	r = append(r, arr[:n]...)
	r = append(r, arr[n+1:]...)
	return r
}

func columnsToInt(s string) []int {
	var arr []int
	for _, i := range strings.Split(s, " ") {
		x, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, x)
	}
	return arr
}

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	reports := my.SplitLines(input)

	{
		n := 0
		for _, r := range reports {
			l := columnsToInt(r)

			if isSafe(l) {
				n++
			}
		}

		fmt.Printf("\nDay %02v: Part 1\n", DAY)
		fmt.Println(n)
	}

	{
		n := 0
		for _, r := range reports {
			l := columnsToInt(r)

			safe := isSafe(l)

			if !safe {
				for i := range l {
					safe = safe || isSafe(removeOne(l, i))
				}
			}

			if safe {
				n++
			}
		}
		fmt.Printf("\nDay %02v: Part 2\n", DAY)
		fmt.Println(n)
	}
}
