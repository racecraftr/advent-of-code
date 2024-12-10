package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	_ "embed"
	"flag"
	"fmt"
)

//go:embed in.txt
var input string

// point stores as x, y
type point [2]int

type set[T comparable] map[T]bool

// dirs are stored as (dx, dy)
var directions = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func parseInput(input string) ([][]rune, []point) {
	lines := util.Lines(input)
	mtx := make([][]rune, len(lines))
	trailheads := make([]point, 0)
	for i, ln := range lines {
		mtx[i] = []rune(ln)
		for j, c := range ln {
			if c == '0' {
				trailheads = append(trailheads, point{j, i})
			}
		}
	}
	return mtx, trailheads
}

func bfs(mtx [][]rune, p point, allowDupes bool) int {
	visited := make(set[point])
	count := 0
	queue := []point{p}
	for len(queue) > 0 {
		deQ := queue[0]
		queue = queue[1:]

		x, y := deQ[0], deQ[1]
		c := mtx[y][x]

		if visited[deQ] && !allowDupes {
			continue
		}

		if c == '9' {
			count++
			visited[deQ] = true
			continue
		}

		for _, dir := range directions {
			dx, dy := dir[0], dir[1]
			if grid.IsValidPos(mtx, x+dx, y+dy) &&
				mtx[y+dy][x+dx] == c+1 {
				queue = append(queue, point{x + dx, y + dy})
			}
		}
	}

	return count
}

func part1(input string) string {
	mtx, trailheads := parseInput(input)
	count := 0
	for _, p := range trailheads {
		count += bfs(mtx, p, false)
	}
	return conv.ToString(count)
}

func part2(input string) string {
	mtx, trailheads := parseInput(input)
	count := 0
	for _, p := range trailheads {
		count += bfs(mtx, p, true)
	}
	return conv.ToString(count)
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
