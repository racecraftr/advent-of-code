package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var wsRegex = regexp.MustCompile("\\s+")

//go:embed in.txt
var input string

func part1(input string) string {
	valid := 0
	lines := strings.Split(input, "\n")
	for _, ln := range lines {
		words := make(map[string]bool)
		for _, word := range wsRegex.Split(ln, -1) {
			if words[word] {
				goto loopEnd
			}
			words[word] = true
		}
		valid++
	loopEnd:
	}

	return conv.ToString(valid)
}

func part2(input string) string {
	valid := 0
	lines := strings.Split(input, "\n")
	for _, ln := range lines {
		words := make(map[string]bool)
		for _, word := range wsRegex.Split(ln, -1) {
			sortedWord := sortString(word)
			if words[sortedWord] {
				goto loopEnd
			}
			words[sortedWord] = true
		}
		valid++
	loopEnd:
	}

	return conv.ToString(valid)
}

func sortString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
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
