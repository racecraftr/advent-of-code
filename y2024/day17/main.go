package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

//go:embed in.txt
var input string

func parseInput(input string) (a, b, c int, ops []int) {
	parts := strings.Split(input, "\n\n")
	re := regexp.MustCompile("\\d+")
	regStrs := re.FindAllString(parts[0], -1)
	a = conv.ToInt(regStrs[0])
	b = conv.ToInt(regStrs[1])
	c = conv.ToInt(regStrs[2])

	ops = arrays.Map(re.FindAllString(parts[1], -1), func(s string) int {
		return conv.ToInt(s)
	})
	return
}

func solve(a, b, c int, ops []int) []int {
	var output []int

	combo := func(operand int) int {
		switch operand {
		case 0, 1, 2, 3:
			return operand
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		}
		panic("Invalid combo operator")
	}

outer:
	for i := 0; i < len(ops)-1; {
		operator := ops[i]
		operand := ops[i+1]
		comboOp := combo(operand)
		switch operator {
		case 0:
			a = a >> comboOp
		case 1:
			b ^= operand
		case 2:
			b = comboOp % 8
		case 3:
			{
				if a != 0 {
					i = operand
					continue outer
				}
			}
		case 4:
			b = b ^ c
		case 5:
			{
				output = append(output, comboOp%8)
			}
		case 6:
			b = a >> comboOp
		case 7:
			c = a >> comboOp
		}
		i += 2
	}

	return output
}

func part1(input string) string {
	a, b, c, ops := parseInput(input) // ok then
	return strings.Join(arrays.Map(solve(a, b, c, ops), func(n int) string {
		return conv.ToString(n)
	}), ",")
}

func arraysAreEqual[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func part2(input string) string {
	_, _, _, ops := parseInput(input)
	var bestInput func(int, int) int
	bestInput = func(cursor int, sofar int) int {
		for i := range 8 {
			candidate := sofar*8 + i
			if arraysAreEqual(solve(candidate, 0, 0, ops), ops[cursor:]) {
				if cursor == 0 {
					return candidate
				}
				if ret := bestInput(cursor-1, candidate); ret >= 0 {
					return ret
				}
			}
		}
		return -1
	}
	res := bestInput(len(ops)-1, 0)
	fmt.Printf("%v\n", ops)
	fmt.Printf("%v\n", solve(res, 0, 0, ops))
	return conv.ToString(bestInput(len(ops)-1, 0))
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
