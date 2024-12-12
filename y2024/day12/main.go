package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	_ "embed"
	"flag"
	"fmt"
)

type point struct {
	r, c int
}

type Set[T comparable] map[T]bool

//go:embed in.txt
var input string

var dirs = []point{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

var cornerToOrtho = map[point][]point{
	{1, 1}:   {{1, 0}, {0, 1}},
	{1, -1}:  {{1, 0}, {0, -1}},
	{-1, 1}:  {{-1, 0}, {0, 1}},
	{-1, -1}: {{-1, 0}, {0, -1}},
}


func parseInput(input string) [][]rune {
	return arrays.Map(util.Lines(input), func(line string) []rune {
		return []rune(line)
	})
}

func explore(curr point, mtx [][]rune, visited Set[point], c rune, area, perimeter *int) {

	neightborCt := 0

	visited[curr] = true
	for _, dir := range dirs {
		next := point{curr.r + dir.r, curr.c + dir.c}
		if grid.IsValidPos(mtx, next.c, next.r) && mtx[next.r][next.c] == c {
			neightborCt ++
			if !visited[next] {
				(*area)++
				explore(next, mtx, visited, c, area, perimeter)
			}
		}
	}
	(*perimeter) += 4 - neightborCt
}


func exploreV2(curr point, mtx [][]rune, visited map[point]bool, currPlant rune, area, corners *int) {
	visited[curr] = true

	for _, dir := range dirs {
		next := point{r: curr.r + dir.r, c: curr.c + dir.c}
		if grid.IsValidPos(mtx, next.c, next.r) && mtx[next.r][next.c] == currPlant {
			if !visited[next] {
				(*area)++
				exploreV2(next, mtx, visited, currPlant, area, corners)
			}
		}
	}

	for corner, pair := range cornerToOrtho {
		c := point{r: curr.r + corner.r, c: curr.c + corner.c}
		i1 := point{r: curr.r + pair[0].r, c: curr.c + pair[0].c}
		i2 := point{r: curr.r + pair[1].r, c: curr.c + pair[1].c}

		if !match(i1, curr, mtx) && !match(i2, curr, mtx) {
			(*corners)++
		}
		if match(i1, curr, mtx) && match(i2, curr, mtx) && !match(curr, c, mtx) {
			(*corners)++
		}
	}
}

func match(i1, i2 point, mtx [][]rune) bool {

	if grid.IsValidPos(mtx, i1.c, i1.r) && !grid.IsValidPos(mtx, i2.c, i2.r) {
		return true
	} else if grid.IsValidPos(mtx, i1.c, i1.r) && grid.IsValidPos(mtx, i2.c, i2.r) {
		p1, p2 := mtx[i1.r][i1.c], mtx[i2.r][i2.c]
		return p1 == p2
	} else {
		return false
	}
}

func part1(input string) string {
	cost := 0
	mtx := parseInput(input)
	visited := make(Set[point])
	for y, row := range mtx{
		for x, c := range row {
			p := point{y, x}
			if !visited[p] {
				var area, perimeter int
				explore(p, mtx, visited, c, &area, &perimeter)
				cost += (area + 1) * perimeter
			}
		}
	}
	return conv.ToString(cost)
}

func part2(input string) string {
	cost := 0
	mtx := parseInput(input)
	visited := make(Set[point])
	for y, row := range mtx{
		for x, c := range row {
			p := point{y, x}
			if !visited[p] {
				var area, perimeter int
				exploreV2(p, mtx, visited, c, &area, &perimeter)
				cost += (area + 1) * perimeter
			}
		}
	}
	return conv.ToString(cost)
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