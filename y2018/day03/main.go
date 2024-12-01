package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

//go:embed in.txt
var input string

func part1(input string) string {
	lines := strings.Split(input, "\n")
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	res := 0
	for _, ln := range lines {
		_, x, y, width, height := parseLn(ln)

		for i := y; i < y+height; i++ {
			for j := x; j < x+width; j++ {
				grid[i][j]++
				if grid[i][j] == 2 {
					res++
				}
			}
		}
	}
	return conv.ToString(res)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	// stores all unique values.
	uniques := make(map[int]bool)
	for i := range len(lines) {
		uniques[i+1] = true
	}

	// grid[i][j] stores an array of ins containing ids
	grid := make([][][]int, 1000)
	for i := range grid {
		grid[i] = make([][]int, 1000)
		for j := range grid {
			grid[i][j] = []int{}
		}
	}
	for _, ln := range lines {
		id, x, y, width, height := parseLn(ln)

		for i := y; i < y+height; i++ {
			for j := x; j < x+width; j++ {
				grid[i][j] = append(grid[i][j], id)
				if len(grid[i][j]) > 1 {
					for _, n := range grid[i][j] {
						delete(uniques, n)
					}
				}
			}
		}
	}
	var res int
	for v := range uniques {
		res = v
	}
	return conv.ToString(res)
}

func parseLn(line string) (id, x, y, width, height int) {
	trimmedLn := line[1:] // remove #
	splitRegex := regexp.MustCompile(" @ |,|: |x")
	parts := splitRegex.Split(trimmedLn, -1)

	id = conv.ToInt(parts[0])
	x = conv.ToInt(parts[1])
	y = conv.ToInt(parts[2])
	width = conv.ToInt(parts[3])
	height = conv.ToInt(parts[4])
	return
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
