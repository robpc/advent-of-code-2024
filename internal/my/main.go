package my

import (
	"log"
	"os"
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

func SplitLines(s string) []string {
	return strings.Split(s, "\n")
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
