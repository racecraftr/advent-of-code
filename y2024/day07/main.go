package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
)

//go:embed in.txt
var input string

type Equation struct {
	num      int
	operands []int
}

var operators = [3]string{"+", "*", "||"}

func parseInput(input string) []*Equation {
	return arrays.Map(strings.Split(input, "\n"), func(line string) *Equation {
		parts := strings.Split(line, ": ")
		num := conv.ToInt(parts[0])
		operands := conv.ToIntArr(strings.Split(parts[1], " "))
		return &Equation{
			num, operands,
		}
	})
}
func concat(i1, i2 int) int {
	n := 1
	for ; i2%n != i2; n *= 10 {
	}
	return i1*n + i2
}

// original solution, did not work for part 2 for some reason lol
// works fine for part 1 :)
func (e *Equation) canMake(base int) bool {
	perms := int(math.Pow(float64(base), float64(
		len(e.operands)-1,
	)))

outer:
	for p := range perms {
		perm, n := p, e.operands[0]
		for _, v := range e.operands[1:] {
			switch operators[perm%base] {
			case "+":
				n += v
			case "*":
				n *= v
			case "||":
				n = concat(n, v)
			}

			if n == e.num {
				return true
			}
			if n > e.num {
				continue outer
			}
			perm /= base
		}
	}
	return false
}

// from https://github.com/mnml/aoc/blob/main/2024/07/1.go. Idk how this works for part 2
// but not canMake. :(
func (e *Equation) isPossible(base int) bool {
	var f func(int, int) bool
	f = func(idx, total int) bool {
		if total > e.num {
			return false
		}
		if idx == len(e.operands) {
			return total == e.num
		}

		c := false
		v := e.operands[idx]
		for i := range base {
			switch operators[i] {
			case "+":
				c = c || f(idx+1, total+v)
			case "*":
				c = c || f(idx+1, total*v)
			case "||":
				c = c || f(idx+1, concat(total, v))
			}
		}
		return c
	}

	return f(1, e.operands[0])
}

func part1(input string) string {
	eqs := parseInput(input)
	sum := 0
	for _, eq := range eqs {
		if eq.canMake(2) {
			sum += eq.num
		}
	}
	return conv.ToString(sum)
}

func part2(input string) string {
	eqs := parseInput(input)
	sum := 0
	for _, eq := range eqs {
		if eq.isPossible(3) {
			sum += eq.num
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
