package main

import (
	"adventOfCode/util"
	"strconv"
)

var keypad = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

/*
    1
  2 3 4
5 6 7 8 9
  A B C
   	D
*/

var stupidKeypad = [][]byte{
	{0, 0, '1', 0, 0},
	{0, '2', '3', '4', 0},
	{'5', '6', '7', '8', '9'},
	{0, 'A', 'B', 'C', 0},
	{0, 0, 'D', 0, 0},
}

func part1() {
	lines := util.GetLinesLocal()
	res := ""
	for _, line := range lines {
		x, y := 1, 1
		for _, c := range line {
			switch c {
			case 'U':
				y = max(0, y-1)
			case 'D':
				y = min(2, y+1)
			case 'L':
				x = max(0, x-1)
			case 'R':
				x = min(2, x+1)
			}
		}
		val := keypad[y][x]
		res += strconv.Itoa(val)
	}

	println(res)
}

func part2() {
	lines := util.GetLinesLocal()
	res := ""
	println(len(stupidKeypad))
	println(len(stupidKeypad[0]))
	for _, line := range lines {
		x, y := 1, 1
		for _, c := range line {
			switch c {
			case 'U':
				{
					y = max(0, y-1)
					if stupidKeypad[y][x] == 0 {
						y++
					}
				}
			case 'D':
				{
					y = min(4, y+1)
					if stupidKeypad[y][x] == 0 {
						y--
					}
				}
			case 'L':
				{
					x = max(0, x-1)
					if stupidKeypad[y][x] == 0 {
						x++
					}
				}
			case 'R':
				{
					x = min(4, x+1)
					if stupidKeypad[y][x] == 0 {
						x--
					}
				}
			}
		}
		val := stupidKeypad[y][x]
		res += string(val)
	}

	println(res)
}

func main() {
	part2()
}
