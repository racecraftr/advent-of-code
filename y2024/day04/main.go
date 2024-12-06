package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed in.txt
var input string

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	runes := make([][]rune, len(lines))
	for i, v := range lines {
		runes[i] = []rune(v)
	}
	return runes
}

var dirs = [8][2]int{
	{0, 1},
	{0, -1},

	{1, 0},
	{-1, 0},

	{1, 1},
	{1, -1},

	{-1, 1},
	{-1, -1},
}

func posValid(grid [][]rune, x, y int) bool {
	return x >= 0 && x < len(grid[0]) &&
		y >= 0 && y < len(grid)
}

func part1(input string) string {
	total := 0
	grid := parseInput(input)
	const strCheck = "MAS"
	for y, ln := range grid {
		for x, c := range ln {
			if c != 'X' {
				continue
			}
			for _, dir := range dirs {
				cx, cy := x+dir[0], y+dir[1]   // order doesn't matter
				for i := range len(strCheck) { // 4 nested for loops. holy shit.
					if !posValid(grid, cx, cy) ||
						grid[cy][cx] != rune(strCheck[i]) {
						goto loopEnd
					}
					cx += dir[0]
					cy += dir[1]
				}
				total++
			loopEnd:
			}
		}
	}
	return conv.ToString(total)
}

func part2(input string) string {
	total := 0
	grid := parseInput(input)
	for y, ln := range grid {
		for x, c := range ln {
			if c != 'A' ||
				x == 0 || y == 0 ||
				x == len(grid[0])-1 || y == len(grid)-1 {
				continue
			}
			tl, tr, bl, br :=
				grid[y-1][x-1],
				grid[y-1][x+1],
				grid[y+1][x-1],
				grid[y+1][x+1]

			diag1 := string(tl) + string(br)
			diag2 := string(tr) + string(bl)
			if (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM") {
				total++
			}
		}
	}
	return conv.ToString(total)
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
