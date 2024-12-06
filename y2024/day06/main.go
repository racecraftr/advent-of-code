package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
)

//go:embed in.txt
var input string

type Set[T comparable] map[T]bool

// point stores it as [x, y]
type point [2]int

// dirs stores it as [dx, dy]
var dirs = [][2]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func parseInput(input string) ([][]rune, point) {
	lines := strings.Split(input, "\n")
	mtx := make([][]rune, len(lines))
	var p point
	for i, line := range lines {
		arr := []rune(line)
		if slices.Contains(arr, '^') {
			p = point{slices.Index(arr, '^'), i}
		}
		mtx[i] = arr
	}
	return mtx, p
}

func getSeen(input string) Set[point] {
	mtx, p := parseInput(input)
	dirIdx := 0
	set := Set[point]{p: true}
	x, y := p[0], p[1]
	for grid.IsValidPos(mtx, x, y) {
		currentDir := dirs[dirIdx]
		dx, dy := currentDir[0], currentDir[1]
		if grid.IsValidPos(mtx, x+dx, y+dy) && mtx[y+dy][x+dx] == '#' {
			dirIdx = (dirIdx + 1) % 4
		} else {
			set[point{x, y}] = true
			x, y = x+dx, y+dy
		}
	}
	return set
}

func part1(input string) string {

	return conv.ToString(len(getSeen(input)))
}

func part2(input string) string {
	mtx, startPos := parseInput(input)
	n, m := len(mtx), len(mtx[0])

	ans := 0
	for i := range n {
		for j := range m {
			if mtx[i][j] == '#' {
				continue
			}
			x, y := startPos[0], startPos[1]
			seen := make(Set[[3]int])
			currDir := 0
			mtx[i][j] = '#'
			for grid.IsValidPos(mtx, x, y) {
				if mtx[y][x] == '#' {
					x, y = x-dirs[currDir][0], y-dirs[currDir][1]
					currDir = (currDir + 1) % len(dirs)
					continue
				}
				key := [3]int{x, y, currDir}
				if seen[key] {
					ans++
					break
				}

				seen[key] = true
				x, y = x+dirs[currDir][0], y+dirs[currDir][1]
			}
			mtx[i][j] = '.'
		}
	}
	return conv.ToString(ans)
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
