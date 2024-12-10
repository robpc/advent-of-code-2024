package main

import (
	"flag"
	"fmt"
	"math"
	"robpc/advent-of-code-2024/internal/my"
	"slices"
	"strconv"
)

const DAY = 9

type Block []int

func (b Block) String() string {
	out := []string{}
	for c := range slices.Chunk(b, 64) {
		bin := ""
		for _, i := range c {
			if i == -1 {
				bin += "...."
			} else {
				bin += fmt.Sprintf("%04x", i)
			}

			bin += " "
		}
		out = append(out, bin)
	}

	s := ""
	repeat := 0
	for i, o := range out {
		if i > 0 && out[i-1] == out[i] {
			repeat++
			continue
		} else {
			if repeat > 0 {
				s += fmt.Sprintf("x%d", repeat)
			}
			repeat = 0
		}

		if i == 0 {
			s += o
		} else {
			s += "\n" + o
		}
	}

	return s
}

func SwapOne(s string, dest int, src int) string {
	if src < dest {
		return s[:src] + string(s[dest]) + s[src+1:dest] + string(s[src]) + s[dest+1:]
	}
	return s[:dest] + string(s[src]) + s[dest+1:src] + string(s[dest]) + s[src+1:]
}

func Subset(set []int, sub []int) (int, int) {
	idx := 0
	for {
		if idx > len(set)-len(sub) {
			// fmt.Println("Subset OOB 1", len(set), "|", idx, idx+len(sub))
			break
		}

		if set[idx] != sub[0] {
			idx++
			continue
		}

		isGood := true
		for i := range len(sub) {
			if sub[i] != set[idx+i] {
				// var b Block = set[idx : idx+len(sub)]
				// fmt.Println("Subset Bad at", i, idx, idx+len(sub), "|", b)
				isGood = false
				idx += i
				break
			}
		}

		if isGood {
			// fmt.Println("Subset Good", idx, idx+len(sub), "|", b)
			return idx, idx + len(sub)
		}
	}
	return -1, -1
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

	id := 0

	var blocks []int
	for i, s := range input {
		isEven := math.Mod(float64(i), 2) == 0

		n, _ := strconv.Atoi(string(s))

		if isEven {
			s := make([]int, n)
			for i := range s {
				s[i] = id
			}
			blocks = append(blocks, s...)
			id++
		} else {
			s := make([]int, n)
			for i := range s {
				s[i] = -1
			}
			blocks = append(blocks, s...)
		}
	}

	{
		n := 0

		res := make([]int, len(blocks))
		copy(res, blocks)

		for x := range res {
			i := len(res) - 1 - x
			id := res[i]

			if id < 0 {
				continue
			}

			r := slices.Index[[]int](res, -1)
			if r < 0 || r > i {
				break
			}
			res[i] = -1
			res[r] = id
		}

		for i, id := range res {
			if id < 0 {
				continue
			}

			n += i * id
		}

		fmt.Printf("\n[Day %02v: Part 1]\n", DAY)
		fmt.Println(n)
	}

	{
		n := 0

		res := make(Block, len(blocks))
		copy(res, blocks)

		// unmoved := []string{}

		currId := id - 1

		for currId >= 0 {
			idx := slices.Index(res, currId)
			if idx < 0 {
				break
			}

			var lidx int = len(res)
			for i, x := range res[idx:] {
				if x != currId {
					lidx = idx + i
					break
				}
			}

			d := make(Block, lidx-idx)
			for i := range d {
				d[i] = -1
			}

			// fmt.Println("## Looking at", currId)
			fIdx, fEnd := Subset(res[:idx], d)
			if fIdx >= 0 {
				// fmt.Printf("> Moving [%04x] %d-%d to %d-%d\n", currId, idx, lidx, fIdx, fEnd)
				s := make([]int, len(res[idx:lidx]))
				copy(s, res[idx:lidx])
				res = append(res[:fIdx], append(s, append(res[fEnd:idx], append(d, res[lidx:]...)...)...)...)
			} else {
				// unmoved = append(unmoved, fmt.Sprintf("%04x", currId))
				// fmt.Println(res)
				// fmt.Printf("! No space found: [%04x] %d-%d\n", currId, idx, lidx)
				// break
			}

			currId--
		}

		for i, id := range res {
			if id < 0 {
				continue
			}

			n += i * id
		}
		// fmt.Printf("! No space found: %s\n", strings.Join(unmoved, ", "))
		// fmt.Println(res)

		fmt.Printf("\n[Day %02v: Part 2]\n", DAY)
		fmt.Println(n)
	}

}
