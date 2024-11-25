package main

import (
	"adventOfCode/util"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
)

//go:embed in.txt
var input string

const LINE_LEN = 8

func transposeBytes(input string) [][]byte {
	lines := strings.Split(input, "\n")
	bytes := make([][]byte, len(lines))

	for i, line := range lines {
		bytes[i] = []byte(line)
	}

	bytes = util.Transpose(bytes)

	return bytes
}

func part1(input string) string{
	bytes := transposeBytes(input)
	res := ""
	for _, byteArr := range bytes {
		maxbyte, maxcount := byte(0), 0
		counts := make([]int, 26)

		for _, b := range byteArr {
			idx := b - 'a'
			counts[idx]++
			if count := counts[idx]; count > maxcount {
				maxcount = count
				maxbyte = b
			}
		}

		res += string(maxbyte)
	}

	return res
}

func part2(input string) string {
	bytes := transposeBytes(input)
	res := ""
	for _, byteArr := range bytes {
		counts := make([]int, 26)

		for _, b := range byteArr {
			idx := b - 'a'
			counts[idx]++
		}

		mincount := math.MaxInt32
		minbyte := byte(0)

		for i, count := range counts {
			if count > 1 && count < mincount {
				minbyte = 'a' + byte(i)
				mincount = count
			}
		}

		res += string(minbyte)
	}

	return res
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
