package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
)

//go:embed in.txt
var input string

func part1(input string) string {
	panic("Unimplemented")
}

func part2(input string) string {
	panic("Unimplemented")
}

func parseInput(input string) []int {
	sArr := util.SplitSpace(input)
	arr := make([]int, len(sArr))

	for i, s := range sArr {
		arr[i] = conv.ToInt(s)
	}
	return arr
}

func memoryReallocation() {

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
