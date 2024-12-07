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

// original solution, updated to work for both parts!
// huge shoutout to u/cranebirdidk on Reddit for finding the mistake
func (e *Equation) canMake(base int) bool {
	perms := int(math.Pow(float64(base), float64(
		len(e.operands)-1,
	)))

outer:
	for p := range perms {
		perm, n := p, e.operands[0]
		for i, v := range e.operands[1:] {
			switch operators[perm%base] {
			case "+":
				n += v
			case "*":
				n *= v
			case "||":
				n = concat(n, v)
			}

			// mistake was here.
			// originally, I was checking to ensure that n == e.num.
			// however, I should have checked to see if I was checking the last operand.
			if n == e.num && i == len(e.operands[1:])-1 {
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
		if eq.canMake(3) {
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
