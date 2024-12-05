package main

import (
	"fmt"
	my "robpc/advent-of-code-2024/internal/my"
	"slices"
	"strconv"
	"strings"
)

const DAY = 5

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	sets := strings.Split(input, "\n\n")

	rulesInput := strings.Split(sets[0], "\n")

	var rules [][2]int = [][2]int{}
	for _, line := range rulesInput {
		r := strings.Split(line, "|")
		before, _ := strconv.Atoi(r[0])
		after, _ := strconv.Atoi(r[1])

		rules = append(rules, [2]int{before, after})
	}

	updatesInput := strings.Split(sets[1], "\n")

	var updates [][]int = [][]int{}
	for i, line := range updatesInput {
		r := strings.Split(line, ",")
		updates = append(updates, []int{})
		for _, s := range r {
			n, _ := strconv.Atoi(s)
			updates[i] = append(updates[i], n)
		}
	}

	var goodUpdates = [][]int{}
	var badUpdates = [][]int{}

	for _, update := range updates {
		lookup := map[int]int{}
		for i, n := range update {
			lookup[n] = i
		}

		isGood := true
		for _, rule := range rules {
			before, bOk := lookup[rule[0]]
			after, aOk := lookup[rule[1]]

			if bOk && aOk && before > after {
				isGood = false
				break
			}
		}

		if isGood {
			goodUpdates = append(goodUpdates, update)
		} else {
			badUpdates = append(badUpdates, update)
		}
	}

	ruleLookup := map[int][]int{}
	for _, r := range rules {
		ruleLookup[r[0]] = append(ruleLookup[r[0]], r[1])
	}

	fixedUpdates := [][]int{}
	for _, original := range badUpdates {
		fixed := []int{}
		rest := make([]int, len(original))
		copy(rest, original)

		for len(rest) > 0 {
			n := rest[0]
			rest = rest[1:]

			afters := ruleLookup[n]

			place := 0
			for i, x := range rest {
				if !slices.Contains(afters, x) {
					place = i + 1
				}
			}

			if place == 0 {
				fixed = append(fixed, n)
			} else if place == len(rest) {
				rest = append(rest, n)
			} else {
				rest = append(rest[:place], append([]int{n}, rest[place:]...)...)
			}
		}

		fixedUpdates = append(fixedUpdates, fixed)
	}

	{
		n := 0

		for _, u := range goodUpdates {
			i := len(u) / 2
			n += u[i]
		}

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		n := 0

		for _, u := range fixedUpdates {
			i := len(u) / 2
			n += u[i]
		}

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}
