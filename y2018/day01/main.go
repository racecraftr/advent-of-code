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
	n := 0
	for _, ln := range lines {
		n += conv.ToInt(ln)
	}
	return fmt.Sprintf("%v", n)
}

func part2(input string) string {
	var n int
	set := map[int]bool{}

	for {
		for _, ln := range strings.Split(input, "\n") {
			n += conv.ToInt(ln)

			if set[n] {
				return conv.ToString(n)
			}
			set[n] = true
		}
	}
	panic("should not reach here")
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(ans)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(ans)
		fmt.Println("Output:", ans)
	}
}
