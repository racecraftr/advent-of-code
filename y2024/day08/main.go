package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed in.txt
var input string

type Set[T comparable] map[T]bool

// each point is stored as [x, y].
type point [2]int

func (p point) unwrap() (int, int) {
	return p[0], p[1]
}

func parseInput(input string) ([][]rune, map[rune][]point) {

	lines := strings.Split(input, "\n")
	mtx := make([][]rune, len(lines))
	antennaMap := make(map[rune][]point)

	for i, ln := range lines {
		mtx[i] = []rune(ln)
		for j, c := range ln {
			if c != '.' {
				antennaMap[c] = append(antennaMap[c], point{j, i})
			}
		}
	}

	return mtx, antennaMap
}

func part1(input string) string {
	mtx, antennaMap := parseInput(input)

	set := make(Set[point]) // all antinode points

	for _, arr := range antennaMap {
		for i := range arr {
			for j := range i {
				x1, y1 := arr[i].unwrap()
				x2, y2 := arr[j].unwrap()
				xDiff := x2 - x1
				yDiff := y2 - y1
				set[point{x2 + xDiff, y2 + yDiff}] = true
				set[point{x1 - xDiff, y1 - yDiff}] = true
			}
		}
	}

	sum := 0
	for p := range set {
		if grid.IsValidPos(mtx, p[0], p[1]) {
			sum++
		}
	}

	return conv.ToString(sum)
}

func part2(input string) string {
	mtx, antennaMap := parseInput(input)

	set := make(Set[point]) // all antinode points

	for _, arr := range antennaMap {
		for i := range arr {
			for j := range i {
				x1, y1 := arr[i].unwrap()
				x2, y2 := arr[j].unwrap()
				xDiff := x2 - x1
				yDiff := y2 - y1

				nx, ny := x1, y1
				for grid.IsValidPos(mtx, nx, ny) {
					set[point{nx, ny}] = true
					nx -= xDiff
					ny -= yDiff
				}

				nx, ny = x1, y1
				for grid.IsValidPos(mtx, nx, ny) {
					set[point{nx, ny}] = true
					nx += xDiff
					ny += yDiff
				}
				//set[point{x2 + xDiff, y2 + yDiff}] = true
				//set[point{x1 - xDiff, y1 - yDiff}] = true
			}
		}
	}

	sum := 0
	for p := range set {
		if grid.IsValidPos(mtx, p[0], p[1]) {
			sum++
		}
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
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
