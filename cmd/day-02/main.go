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

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	reports := strings.Split(input, "\n")

	n := 0
	// var deltas [][]int
	for _, r := range reports {
		var l []int
		for _, i := range strings.Split(r, " ") {
			x, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			l = append(l, x)
		}

		if isSafe(l) {
			n++
		}
	}

	fmt.Printf("Day %02v: Part 1\n", DAY)
	fmt.Println(n)

	n2 := 0
	for _, r := range reports {
		var l []int
		for _, i := range strings.Split(r, " ") {
			x, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			l = append(l, x)
		}

		if isSafe(l) {
			n2++
		} else {
			safe := false
			for i := range l {
				safe = safe || isSafe(removeOne(l, i))
			}
			if safe {
				n2++
			}
		}
	}
	fmt.Printf("Day %02v: Part 2\n", DAY)
	fmt.Println(n2)

}
