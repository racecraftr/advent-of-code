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
	sum := 0
	for i := 0; i < len(input); i++ {
		if c := input[i]; c == input[(i+1)%len(input)] {
			sum += int(c - '0')
		}
	}
	return conv.ToString(sum)
}

func part2(input string) string {
	sum, l := 0, len(input)
	if l%2 != 0 {
		panic("invalid length")
	}
	halfIdx := l / 2
	for i := 0; i <= halfIdx; i++ {
		if c := input[i]; c == input[(i+halfIdx)%l] {
			sum += int(c-'0') * 2
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
