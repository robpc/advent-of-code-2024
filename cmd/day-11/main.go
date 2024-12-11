package main

import (
	"flag"
	"fmt"
	"robpc/advent-of-code-2024/internal/my"
	"strconv"
)

const DAY = 11

func blink(stones []int) []int {
	blink := []int{}
	for _, x := range stones {
		if x == 0 {
			blink = append(blink, 1)
			continue
		}

		s := strconv.Itoa(x)
		if my.Mod(len(s), 2) == 0 {
			i := len(s) / 2
			s1 := s[:i]
			s2 := s[i:]

			i1 := my.ToInt(s1)
			i2 := my.ToInt(s2)

			blink = append(blink, i1, i2)
			continue
		}

		blink = append(blink, x*2024)
	}
	return blink
}

func single(x int) []int {
	if x == 0 {
		return []int{1}
	}

	s := strconv.Itoa(x)
	if my.Mod(len(s), 2) == 0 {
		i := len(s) / 2
		s1 := s[:i]
		s2 := s[i:]

		i1 := my.ToInt(s1)
		i2 := my.ToInt(s2)

		return []int{i1, i2}
	}

	return []int{x * 2024}
}

var stats map[string]int = map[string]int{"total": 0, "hits": 0, "misses": 0}
var cache map[string]int = map[string]int{}

func get(stones []int, i int) int {
	key := fmt.Sprint(stones, i)
	n, ok := cache[key]
	if !ok {
		// n = oneOf(stones, i)

		n = 0
		if i == 0 {
			n = len(stones)
		} else {
			for _, x := range stones {
				n += get(single(x), i-1)
			}
		}

		cache[key] = n
	}

	if ok {
		stats["hits"] += 1
	} else {
		stats["misses"] += 1
	}
	stats["total"] += 1

	return n

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

	list := my.StringIntSplit(input, " ")

	{
		n := 0

		stones := make([]int, len(list))
		copy(stones, list)

		for range 25 {
			stones = blink(stones)
		}

		n = len(stones)

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		n := 0

		stones := make([]int, len(list))
		copy(stones, list)

		n = get(stones, 75)

		fmt.Println("Cache Stats:", stats, "Items:", len(cache))

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}
