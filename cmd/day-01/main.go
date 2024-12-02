package main

import (
	"fmt"
	"log"
	my "robpc/advent-of-code-2024/internal/my"
	"slices"
	"strconv"
	"strings"
)



func main() {
    fmt.Println("Hello, World!")
    var input = my.ReadFile("./inputs/day-01/input.txt")

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
      distance = append(distance, my.Abs(left[i] - right[i]))
    }

    var n int = my.Sum(distance)
 
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

    var n2 int = my.Sum(similarity)
 
    fmt.Println("Day 01: Part 2")
    fmt.Println(n2)

}