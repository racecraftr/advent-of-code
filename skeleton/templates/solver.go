package main

import (
	"adventOfCode/util"
	_ "embed"
	"flag"
	"fmt"
)

//go:embed in.txt
var input string

func part1(intput string) string {
	panic("Unimplemented")
}

func part2(input string) string {
	panic("Unimplemented")
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