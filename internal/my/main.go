package my

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
//	defer timer("sum")()
func Timer(name string) func() {
	fmt.Printf("==== %s ====\n", name)
	start := time.Now()

	return func() {
		fmt.Printf("==== took %v ====\n", time.Since(start))
	}
}

func Mod(n int, m int) float64 {
	return math.Mod(float64(n), 2)
}

func Sum(arr []int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	return string(data)
}

func StringIntSplit(input string, sep string) []int {
	var arr []int
	for _, i := range strings.Split(input, sep) {
		x, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, x)
	}
	return arr
}

func StringGrid(input string, xSep string, ySep string) [][]string {
	var grid [][]string = [][]string{}

	lineArray := strings.Split(input, ySep)

	for _, line := range lineArray {
		grid = append(grid, strings.Split(line, xSep))
	}

	return grid
}

func IntGridFromString(input string, xSep string, ySep string) [][]int {
	var grid [][]int = [][]int{}

	lineArray := strings.Split(input, ySep)

	for _, line := range lineArray {
		grid = append(grid, StringIntSplit(line, xSep))
	}

	return grid
}

func UnzipIntGrid(grid [][]int) ([]int, []int) {
	var a []int
	var b []int
	for _, r := range grid {
		a = append(a, r[0])
		b = append(b, r[1])
	}

	return a, b
}
