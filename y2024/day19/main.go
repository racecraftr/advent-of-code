package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed in.txt
var input string

func parseInput(input string) (available, designs []string) {
	parts := strings.Split(input, "\n\n")
	available = strings.Split(parts[0], ", ")
	designs = util.Lines(parts[1])
	return
}

func part1(input string) string {
	patterns, designs := parseInput(input)

	cache := map[string]bool{}
	var isPossible func(string) int
	isPossible = func(design string) int {
		if cache[design] {
			return 1
		}
		for _, pattern := range patterns {
			if len(pattern) > len(design) {
				continue
			}
			if pattern == design {
				cache[design] = true
				return 1
			}
			if design[:len(pattern)] == pattern && isPossible(design[len(pattern):]) == 1 {
				cache[design] = true
				return 1
			}
		}
		return 0
	}

	return conv.ToString(arrays.Sum(arrays.Map(designs, isPossible)))
}

func part2(input string) string {
	patterns, designs := parseInput(input)

	cache := map[string]int{}
	var numCombos func(string) int
	numCombos = func(design string) int {
		res := 0
		if v, ok := cache[design]; ok {
			return v
		}
		for _, pattern := range patterns {
			if len(pattern) > len(design) {
				continue
			}
			if pattern == design {
				res += 1
			}
			if design[:len(pattern)] == pattern {
				res += numCombos(design[len(pattern):])
			}
		}
		cache[design] = res
		return res
	}

	return conv.ToString(arrays.Sum(arrays.Map(designs, numCombos)))
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
