package my

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	return string(data)
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
