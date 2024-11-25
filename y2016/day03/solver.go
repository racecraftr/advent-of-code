package main

import (
	"adventOfCode/util"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	lines := util.GetLinesLocal()
	valid := 0

	for _, ln := range lines {
		line := strings.TrimSpace(ln)
		r := regexp.MustCompile(" +")
		parts := r.Split(line, 3)
		nums := make([]int, 3)

		for i, part := range parts {
			// println(part)
			n, err := strconv.Atoi(part)
			util.Check(err)
			nums[i] = n
		}

		sort.Ints(nums)

		if nums[0]+nums[1] > nums[2] {
			valid++
		}
	}

	println(valid)
}

func part2() {
	lines := util.GetLinesLocal()
	ints := make([][]int, 0)
	valid := 0

	for _, ln := range lines {
		line := strings.TrimSpace(ln)
		r := regexp.MustCompile(" +")
		parts := r.Split(line, 3)
		nums := make([]int, 3)

		for i, part := range parts {
			// println(part)
			n, err := strconv.Atoi(part)
			util.Check(err)
			nums[i] = n
		}

		ints = append(ints, nums)
		if len(ints) == 3 {
			for i := range 3 {
				col := []int{ints[0][i], ints[1][i], ints[2][i]}
				sort.Ints(col)
				if col[0]+col[1] > col[2] {
					valid++
				}
			}
			ints = make([][]int, 0)
		}
	}

	println(valid)
}

func main() {
	part2()
}
