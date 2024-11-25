package main

import (
	"adventOfCode/util"
	"math"
	"strconv"
	"strings"
)

func dirChar(dir int) byte {
	switch dir {
	case 0:
		return 'N'
	case 1:
		return 'E'
	case 2:
		return 'S'
	case 3:
		return 'W'
	}
	panic("invalid direction value")
}

func part1() {
	x, y := 0, 0

	/*
		0 = north
		1 = east
		2 = south
		3 = west
	*/
	dir := 0
	content := util.GetContentLocal()
	steps := strings.Split(content, ", ")
	for _, step := range steps {
		turn := step[0]
		incr, err := strconv.Atoi(string(step[1:]))
		util.Check(err)

		switch turn {
		case 'L':
			dir--
		case 'R':
			dir++
		}

		if dir < 0 {
			dir = 3
		}

		if dir > 3 {
			dir = 0
		}

		switch dir {
		case 0:
			y += incr
		case 1:
			x += incr
		case 2:
			y -= incr
		case 3:
			x -= incr
		}
	}
	println(int(math.Abs(float64(x)) + math.Abs(float64(y))))
}

type Pos struct {
	x, y int
}

func part2() {
	x, y := 0, 0
	dx, dy := 0, 0

	stor := map[Pos]bool{}

	/*
		0 = north
		1 = east
		2 = south
		3 = west
	*/
	dir := 0
	content := util.GetContentLocal()
	steps := strings.Split(content, ", ")
	for _, step := range steps {
		turn := step[0]
		incr, err := strconv.Atoi(string(step[1:]))
		util.Check(err)

		switch turn {
		case 'L':
			dir--
		case 'R':
			dir++
		}

		if dir < 0 {
			dir = 3
		}

		if dir > 3 {
			dir = 0
		}

		switch dir {
		case 0:
			dx, dy = 0, 1
		case 1:
			dx, dy = 1, 0
		case 2:
			dx, dy = 0, -1
		case 3:
			dx, dy = -1, 0
		}

		for range incr {
			x, y = x+dx, y+dy
			p := Pos{x, y}
			if ok := stor[p]; ok {
				println(int(math.Abs(float64(x)) + math.Abs(float64(y))))
				return
			}
			stor[p] = true
		}

	}
}

func main() {
	part2()
}
