package main

import (
	"adventOfCode/util"
	"strconv"
	"strings"
)

type InstType int

const (
	off InstType = iota
	on
	toggle
)

type Instruction struct {
	instType InstType
	x1       int
	y1       int
	x2       int
	y2       int
}

func instTypeFromStr(s string) InstType {
	switch s {
	case "off":
		return off
	case "on":
		return on
	case "toggle":
		return toggle
	}
	panic("Invalid string to parse: " + s)
}

func parseCoord(coord string) (int, int) {
	parts := strings.Split(coord, ",")
	if len(parts) != 2 {
		panic("Invalid number of comma-delimited parts")
	}

	x, err := strconv.Atoi(parts[0])
	util.Check(err)

	y, err := strconv.Atoi(parts[1])
	util.Check(err)

	return x, y
}

func parseLine(line string) *Instruction {
	line = strings.ReplaceAll(line, " through ", " ")
	line = strings.ReplaceAll(line, "turn ", "")
	parts := strings.Split(line, " ")
	inst, coord1, coord2 := parts[0], parts[1], parts[2]

	instType := instTypeFromStr(inst)
	x1, y1 := parseCoord(coord1)
	x2, y2 := parseCoord(coord2)

	return &Instruction{instType, x1, y1, x2, y2}
}

func part1() {
	lights := make([][]bool, 1000)
	for i := range lights {
		lights[i] = make([]bool, 1000)
	}

	lines := util.GetLinesLocal()
	count := 0
	for _, line := range lines {
		instruction := parseLine(line)
		for i := instruction.y1; i <= instruction.y2; i++ {
			for j := instruction.x1; j <= instruction.x2; j++ {
				switch instruction.instType {
				case off:
					{
						if lights[i][j] {
							count--
						}
						lights[i][j] = false
					}
				case on:
					{
						if !lights[i][j] {
							count++
						}
						lights[i][j] = true
					}
				case toggle:
					{
						if lights[i][j] {
							count--
						} else {
							count++
						}
						lights[i][j] = !lights[i][j]
					}
				}
			}
		}
	}
	println(count)
}

func part2() {
	lights := make([][]int, 1000)
	for i := range lights {
		lights[i] = make([]int, 1000)
	}

	lines := util.GetLinesLocal()
	for _, line := range lines {
		instruction := parseLine(line)
		for i := instruction.y1; i <= instruction.y2; i++ {
			for j := instruction.x1; j <= instruction.x2; j++ {
				switch instruction.instType {
				case off:
					{
						lights[i][j] = max(lights[i][j]-1, 0)
					}
				case on:
					{
						lights[i][j] += 1
					}
				case toggle:
					{
						lights[i][j] += 2
					}
				}
			}
		}
	}

	count := 0
	for _, row := range lights {
		for _, n := range row {
			count += n
		}
	}
	println(count)
}

func main() {
	part2()
}
