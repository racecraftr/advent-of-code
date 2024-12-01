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

func part1(input string) string {
	lines := strings.Split(input, "\n")
	arr := make([]int, len(lines))
	for i, v := range lines {
		arr[i] = conv.ToInt(v)
	}

	idx, steps := 0, 0
	for idx < len(arr) && idx >= 0 {
		prevIdx := idx
		idx += arr[idx]
		arr[prevIdx]++
		steps++
	}
	return conv.ToString(steps)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	arr := make([]int, len(lines))
	for i, v := range lines {
		arr[i] = conv.ToInt(v)
	}

	idx, steps := 0, 0
	for idx < len(arr) && idx >= 0 {
		prevIdx := idx
		idx += arr[idx]
		if arr[prevIdx] >= 3 {
			arr[prevIdx]--
		} else {
			arr[prevIdx]++
		}
		steps++
	}
	return conv.ToString(steps)
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
