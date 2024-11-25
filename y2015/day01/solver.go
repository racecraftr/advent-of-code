package main

import "adventOfCode/util"

func part1() {
	in := util.GetContentLocal()
	n := 0
	for _, c := range in {
		switch c {
		case '(':
			n++
		case ')':
			n--
		}
	}
	println(n)
}

func part2() {
	in := util.GetContentLocal()
	n := 0
	for i, c := range in {
		switch c {
		case '(':
			n++
		case ')':
			n--
		}
		if n == -1 {
			println(i + 1)
			return
		}
	}
}

func main() {
	part2()
}
