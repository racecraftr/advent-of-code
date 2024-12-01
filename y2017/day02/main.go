package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strings"
)

//go:embed in.txt
var input string

var wsRegex = regexp.MustCompile("\\s+")

func part1(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, ln := range lines {
		lnMin, lnMax := math.MaxInt, -1
		for _, nStr := range wsRegex.Split(ln, -1) {
			n := conv.ToInt(nStr)
			lnMin, lnMax = min(lnMin, n), max(lnMax, n)
		}
		sum += lnMax - lnMin
	}
	return conv.ToString(sum)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, ln := range lines {
		intStrArr := wsRegex.Split(ln, -1)
		intArr := make([]int, len(intStrArr))
		for i, s := range intStrArr {
			intArr[i] = conv.ToInt(s)
		}
		for i := 0; i < len(intArr); i++ {
			for j := 0; j < i; j++ {
				num, denom := max(intArr[i], intArr[j]), min(intArr[i], intArr[j])
				if num%denom == 0 {
					sum += num / denom
					goto loopEnd
				}
			}
		}
	loopEnd:
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
