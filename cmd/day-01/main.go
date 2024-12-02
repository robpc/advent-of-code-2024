package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) string {
  data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	return string(data)
}

func sum(arr []int) int {
  sum := 0
  for _,x := range arr {
    sum += x
  }
  return sum
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func main() {
    fmt.Println("Hello, World!")
    var input = readFile("./inputs/day-01/input.txt")

    split1 := strings.Split(input, "\n")

    var left []int
    var right []int

    for _,element := range split1 {
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

    var distance []int

    for i := range left {
      distance = append(distance, abs(left[i] - right[i]))
    }

    var n int = sum(distance)
 
    fmt.Println("Day 01: Part 1")
    fmt.Println(n)

    var similarity []int

    for i := range left {
      x := left[i]
      c := 0
      for _,e := range right {
        if e == x {
          c++
        }
      }
      similarity = append(similarity, x * c)
    }

    var n2 int = sum(similarity)
 
    fmt.Println("Day 01: Part 2")
    fmt.Println(n2)

}