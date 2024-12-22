package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	"adventOfCode/util/mathy"
	_ "embed"
	"flag"
	"fmt"
)

// for some reason, this code does not work.

/*
x 0   1   2
+---+---+---+  y
| 7 | 8 | 9 |  3
+---+---+---+
| 4 | 5 | 6 |  2
+---+---+---+
| 1 | 2 | 3 |  1
+---+---+---+
    | 0 | A |  0
    +---+---+
*/

var numPad = map[rune]grid.Point{
	'0': {1, 0},
	'A': {2, 0},

	'1': {0, 1},
	'2': {1, 1},
	'3': {2, 1},

	'4': {0, 2},
	'5': {1, 2},
	'6': {2, 2},

	'7': {0, 3},
	'8': {1, 3},
	'9': {2, 3},
}

/*
x 0     1   2
    +---+---+ y
    | ^ | A | 1
+---+---+---+
| < | v | > | 0
+---+---+---+
*/

var dirPad = map[rune]grid.Point{
	'^': {1, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
	'A': {2, 0},
}

//go:embed in.txt
var input string

// numpad <- dirPad <- dirPad <- dirPad

func numPadPresses(seq string, start rune) string {
	current := numPad[start]
	output := ""
	for _, r := range []rune(seq) {
		dest := numPad[r]
		diffX, diffY := dest[0]-current[0], dest[1]-current[1]

		h, v := "", ""
		for range mathy.IntAbs(diffX) {
			if diffX >= 0 {
				h += ">"
			} else {
				h += "<"
			}
		}

		for range mathy.IntAbs(diffY) {
			if diffY >= 0 {
				v += "^" // remember: going down is increasing in y direction
			} else {
				v += "v"
			}
		}

		if current[1] == 0 && dest[0] == 0 {
			output += v + h
		} else if current[0] == 0 && dest[0] == 0 {
			output += h + v
		} else if diffX < 0 {
			output += h + v
		} else if diffX >= 0 {
			output += v + h
		}

		current = dest
		output += "A"
	}
	return output
}

func dirPadPresses(input string, start rune) string {
	current := dirPad[start]
	output := ""
	for _, r := range input {
		dest := dirPad[r]
		diffX, diffY := dest[0]-current[0], dest[1]-current[1]
		h, v := "", ""
		for range mathy.IntAbs(diffX) {
			if diffX >= 0 {
				h += ">"
			} else {
				h += "<"
			}
		}

		for range mathy.IntAbs(diffY) {
			if diffY >= 0 {
				v += "^"
			} else {
				v += "v"
			}
		}

		if current[0] == 0 && dest[1] == 1 {
			output += h + v
		} else if current[1] == 1 && dest[0] == 0 {
			output += v + h
		} else if diffX < 0 {
			output += h + v
		} else if diffX >= 0 {
			output += v + h
		}
		current = dest
		output += "A"
	}
	return output
}

func getSeq(lines []string, robots int) int {
	count := 0
	cache := make(map[string][]int)
	for _, ln := range lines {
		seq := numPadPresses(ln, 'A')
		num := countAfterRobots(seq, robots, 1, cache)
		count += conv.ToInt(ln[:3]) * num
	}
	return count
}

func countAfterRobots(input string, maxRobots, robot int, cache map[string][]int) int {
	if val, ok := cache[input]; ok {
		if val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		cache[input] = make([]int, maxRobots)
	}

	seq := dirPadPresses(input, 'A')
	cache[input][0] = len(seq)

	if robot == maxRobots {
		return len(seq)
	}

	splitSeq := getIndividualSteps(seq)
	count := 0
	for _, s := range splitSeq {
		c := countAfterRobots(s, maxRobots, robot+1, cache)
		if _, ok := cache[s]; !ok {
			cache[s] = make([]int, maxRobots)
		}
		cache[s][0] = c
		count += c
	}

	cache[input][robot-1] = count
	return count
}

func getIndividualSteps(input string) []string {
	var res []string
	current := ""
	for _, c := range input {
		current += string(c)
		if c == 'A' {
			res = append(res, current)
			current = ""
		}
	}
	return res
}

func part1(input string) string {
	return conv.ToString(getSeq(util.Lines(input), 2))
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
		_ = part1("029A")
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
