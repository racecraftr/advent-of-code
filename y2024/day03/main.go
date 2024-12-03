package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

//go:embed in.txt
var input string

func calcMul(line string) int {
	sum := 0
	var re = regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	mulStatements := re.FindAllString(line, -1)
	for _, statement := range mulStatements {
		strs := regexp.MustCompile("\\d+").FindAllString(statement, -1)
		sum += conv.ToInt(strs[0]) * conv.ToInt(strs[1])
	}
	return sum
}

func part1(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, ln := range lines {
		sum += calcMul(ln)
	}
	return conv.ToString(sum)
}

func part2(input string) string {
	re := regexp.MustCompile("(mul\\(\\d+,\\d+\\))|(do\\(\\))|(don't\\(\\))")
	statements := re.FindAllString(input, -1)
	sum, enabled := 0, true
	for _, statement := range statements {
		if strings.HasPrefix(statement, "mul") && enabled {
			strs := regexp.MustCompile("\\d+").FindAllString(statement, -1)
			sum += conv.ToInt(strs[0]) * conv.ToInt(strs[1])
			continue
		}
		if strings.HasPrefix(statement, "don't") {
			enabled = false
			continue
		}
		if strings.HasPrefix(statement, "do") {
			enabled = true
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
