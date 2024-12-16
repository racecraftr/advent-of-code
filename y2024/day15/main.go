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

type point [2]int

type Set[T comparable] map[T]bool

var dirs = map[rune]point{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

const (
	boundX = 50
	boundY
	boundX2 = 100
	boundY2
)

func (s Set[T]) replace(old, new T) bool {
	if !s[old] {
		return false
	}
	delete(s, old)
	s[new] = true
	return true
}

func parseInput(input string) (boxes, walls Set[point], robotPos point, instructions string) {
	parts := strings.Split(input, "\n\n")
	instructions = strings.ReplaceAll(parts[1], "\n", "")

	boxes = make(Set[point])
	walls = make(Set[point])

	for y, row := range util.Lines(parts[0]) {
		for x, c := range row {
			switch c {
			case '@':
				robotPos = point{x, y}
			case 'O':
				boxes[point{x, y}] = true
			case '#':
				walls[point{x, y}] = true
			}
		}
	}
	return
}

func parseInput2(input string) (boxes, walls Set[point], robotPos point, instructions string) {
	parts := strings.Split(input, "\n\n")
	instructions = strings.ReplaceAll(parts[1], "\n", "")

	boxes = make(Set[point])
	walls = make(Set[point])

	for y, row := range util.Lines(parts[0]) {
		for x, c := range row {
			switch c {
			case '@':
				robotPos = point{x * 2, y}
			case 'O':
				boxes[point{x * 2, y}] = true
			case '#':
				walls[point{x * 2, y}] = true
				walls[point{x*2 + 1, y}] = true
			}
		}
	}
	return
}

func render(boxes, walls Set[point], bounds [2]int) {
	grid := make([][]rune, bounds[0])
	for y := range grid {
		grid[y] = make([]rune, bounds[1])
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}
	for wall := range walls {
		x, y := wall[0], wall[1]
		grid[y][x] = '#'
	}

	for box := range boxes {
		x, y := box[0], box[1]
		grid[y][x] = 'O'
	}

	for _, ln := range grid {
		fmt.Println(string(ln))
	}
}

func part1(input string) string {
	boxes, walls, robotPos, instructions := parseInput(input)
outer:
	for _, c := range instructions {
		dir := dirs[c]
		newPos := point{robotPos[0] + dir[0], robotPos[1] + dir[1]}
		if walls[newPos] { // robot does not move
			continue outer
		}
		if boxes[newPos] { // update boxes here.
			pushPos := point{newPos[0], newPos[1]}
			for boxes[pushPos] {
				pushPos = point{pushPos[0] + dir[0], pushPos[1] + dir[1]}
			}
			if walls[pushPos] {
				continue outer
			}
			// now, some swapping.
			delete(boxes, newPos)
			boxes[pushPos] = true
		}
		robotPos = newPos
	}

	render(boxes, walls, point{boundX, boundY})
	sum := 0
	for box := range boxes {
		sum += box[1]*100 + box[0]
	}
	return conv.ToString(sum)
}

func (p point) copy() point {
	return point{p[0], p[1]}
}

func part2(input string) string {
	panic("I'm not able to solve this, fuck this shit")
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
