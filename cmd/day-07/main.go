package main

import (
	"fmt"
	"log"
	"math"
	"robpc/advent-of-code-2024/internal/my"
	"strconv"
)

const DAY = 7

func genOperators(arr []string, ops []string, n int) [][]string {
	if n == 0 {
		return [][]string{arr}
	}

	var r [][]string
	for _, o := range ops {
		var b []string
		b = make([]string, len(arr))
		copy(b, arr)
		b = append(b, o)
		r = append(r, genOperators(b, ops, n-1)...)
	}
	return r
}

func calc(values []int, ops []string) int {

	if len(values) == 1 {
		return values[0]
	}

	a := values[0]
	b := values[1]
	o := ops[0]

	if o == "+" {
		nv := append([]int{a + b}, values[2:]...)
		no := ops[1:]
		// fmt.Println(nv, no)
		return calc(nv, no)
	} else if o == "*" {
		nv := append([]int{a * b}, values[2:]...)
		no := ops[1:]
		// fmt.Println(nv, no)
		return calc(nv, no)
	} else if o == "||" {
		s := strconv.Itoa(b)
		digits := len(s)
		a1 := a * int(math.Pow(10, float64(digits)))
		nv := append([]int{a1 + b}, values[2:]...)
		no := ops[1:]
		// fmt.Println(nv, no, a, int(math.Pow(10, float64(digits))), digits, b)
		return calc(nv, no)
	}

	log.Fatal("bob")
	return 0
}

func main() {
	var input = my.ReadFile(fmt.Sprintf("./inputs/day-%02v/input.txt", DAY))

	var temp1 [][]string = my.StringGrid(input, ": ", "\n")
	// fmt.Println(temp1)

	operators := []string{"+", "*"}
	operators2 := []string{"+", "*", "||"}
	{
		n := 0

		for _, a := range temp1 {
			// fmt.Println("a", a)
			total := my.ToInt(a[0])
			vals := my.StringIntSplit(a[1], " ")
			ops := genOperators([]string{}, operators, len(vals))
			// fmt.Println(total, vals, ops)
			for _, o := range ops {
				t := calc(vals, o)
				if t == total {
					n += t
					fmt.Println("yay", t, vals, o)
					break
				}
			}
		}

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		n := 0

		for _, a := range temp1 {
			// fmt.Println("a", a)
			total := my.ToInt(a[0])
			vals := my.StringIntSplit(a[1], " ")
			ops := genOperators([]string{}, operators2, len(vals)-1)
			// fmt.Println(total, vals, ops)
			for _, o := range ops {
				t := calc(vals, o)
				if t == total {
					n += t
					// fmt.Println("yay", t, vals, o)
					break
				}
			}
		}

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}
