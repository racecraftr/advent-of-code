package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/mathy"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
)

//go:embed in.txt
var input string

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	rows := make([][]int, len(lines))
	for i, ln := range lines {
		parts := util.SplitSpace(ln)
		row := make([]int, len(parts))
		for j, part := range parts {
			row[j] = conv.ToInt(part)
		}
		rows[i] = row
	}
	return rows
}

func safe(row []int) bool {
	decreasing := row[0] > row[1]
	for i := range len(row) - 1 {
		diff := mathy.IntAbs(row[i] - row[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
		currDec := row[i] > row[i+1]
		if currDec != decreasing {
			return false
		}
	}
	return true
}

func part1(input string) string {
	rows := parseInput(input)
	safeReps := 0
	for _, row := range rows {
		if safe(row) {
			safeReps++
		}
	}
	return conv.ToString(safeReps)
}

func part2(input string) string {
	rows := parseInput(input)
	safeReps := 0
outer:
	for _, row := range rows {
		for i := range row {
			newRow := slices.Clone(row)
			newRow = append(newRow[:i], newRow[i+1:]...)
			if safe(newRow) {
				safeReps++
				continue outer
			}
		}
	}
	return conv.ToString(safeReps)
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
