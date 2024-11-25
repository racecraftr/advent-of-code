package main

import (
	"adventOfCode/util"
	"strconv"
)

func part1() {
	x := 0
	lines := util.GetLinesLocal()
	for _, ln := range lines {
		n, err := strconv.Atoi(ln)
		util.Check(err)
		x += n
	}
	println(x)
}

func part2() {
	nums := map[int]bool{0: true}
	x := 0
	lines := util.GetLinesLocal()
	for _, ln := range lines {
		n, err := strconv.Atoi(ln)
		util.Check(err)
		x += n
		if nums[x] {
			println(x)
			return
		}
		nums[x] = true
	}
}

func main() {
	part2()
}
