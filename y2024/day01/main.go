package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"
)

//go:embed in.txt
var input string

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	mtx := make([][]int, len(lines))
	for i, ln := range lines {
		parts := util.SplitSpace(ln)
		mtx[i] = []int{
			conv.ToInt(parts[0]), conv.ToInt(parts[1]),
		}
	}

	// now the columns will be rows.
	mtx = util.Transpose(mtx)
	return mtx
}

func part1(input string) string {
	mtx := parseInput(input)
	sort.Ints(mtx[0])
	sort.Ints(mtx[1])
	sum := 0
	for i := range mtx[0] {
		diff := mtx[0][i] - mtx[1][i]
		absDiff := int(math.Abs(float64(diff)))
		//fmt.Printf("values %v and %v have a distance of %v\n", mtx[0][i], mtx[1][i], diff)
		sum += absDiff
	}
	return conv.ToString(sum)
}

func part2(input string) string {
	mtx := parseInput(input)
	appearances := make(map[int]int)
	for _, v := range mtx[1] {
		if _, ok := appearances[v]; !ok {
			appearances[v] = 0
		}
		appearances[v]++
	}

	simScore := 0
	for _, v := range mtx[0] {
		if apps, ok := appearances[v]; ok {
			simScore += v * apps
		}
	}
	return conv.ToString(simScore)
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
