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
	e2, e3 := 0, 0
	for _, ln := range lines {
		count := charCount(ln)
		for _, v := range count {
			if v == 2 {
				e2++
				break
			}
		}

		for _, v := range count {
			if v == 3 {
				e3++
			}
		}
	}

	return conv.ToString(e2 * e3)
}

func charCount(s string) map[rune]int {
	res := make(map[rune]int)
	for _, c := range s {
		res[c]++
	}
	return res
}

func part2(input string) string {

	visited := []string{}
	lines := strings.Split(input, "\n")
	for _, ln := range lines {
		for _, v := range visited {
			diffstr := ""
			for i, c := range ln {
				if c == rune(v[i]) {
					diffstr += string(c)
				}
			}
			if len(diffstr) == len(ln)-1 {
				println(ln)
				println(v)
				return diffstr
			}
		}
		visited = append(visited, ln)
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
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
