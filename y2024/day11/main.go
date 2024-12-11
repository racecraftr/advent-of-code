package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"math"
)

//go:embed in.txt
var input string


func parseInput(input string) []int {
	return arrays.Map(util.SplitSpace(input), func(s string) int {
		return conv.ToInt(s)
	})
}

func numDigits(i int) int {
	return (int(math.Log10(float64(i)))) + 1
}

func powTen(pow int) int {
	n := 1
	for range pow {
		n *= 10
	}
	return n;
}

func blink(stones map[int]int) map[int]int{
	newStones := map[int]int{}

	add := func(key, incr int) {
		if _, ok := newStones[key]; !ok {
			newStones[key] = 0
		}
		newStones[key] += incr
	}

	for stone, count := range stones {
		if stone == 0 {
			add(1, count)
		} else if digits := numDigits(stone); digits % 2 == 0 {
			filter := powTen(digits / 2)
			left, right := stone/filter, stone%filter
			add(left, count)
			add(right, count)
		} else {
			add(stone * 2024, count)
		}
	}
	return newStones
}

func part1(input string) string {
	cache := map[int]int{}

	for _, v := range parseInput(input) {
		cache[v] = 1
	}

	fmt.Printf("%v\n", cache)

	for range 25 {
		cache = blink(cache)
	}

	sum := 0
	for _, v := range cache {
		sum += v
	}
	return conv.ToString(sum)
}

func part2(input string) string {
	cache := map[int]int{}

	for _, v := range parseInput(input) {
		cache[v] = 1
	}

	fmt.Printf("%v\n", cache)

	for range 75 {
		cache = blink(cache)
	}

	sum := 0
	for _, v := range cache {
		sum += v
	}
	return conv.ToString(sum)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans) // o_<
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}