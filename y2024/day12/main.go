package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	_ "embed"
	"flag"
	"fmt"
)

type point [2]int

type Set[T comparable] map[T]bool

//go:embed in.txt
var input string

var dirs = [][2]int {
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func parseInput(input string) [][]rune {
	return arrays.Map(util.Lines(input), func(line string) []rune {
		return []rune(line)
	})
}

func part1(input string) string {
	mtx := parseInput(input)
	visited := make(Set[point])

	sum := 0
	for y, row := range mtx {

		for x, c := range row {

			p := point{x, y}

			if visited[p] {
				continue
			}

			visited[p] = true

			perimSquares := map[point]int{}
			areaSquares := make(Set[point])

			queue := []point{p}

			for len(queue) > 0 {
				deQ := queue[0]
				queue = queue[1:]

				if areaSquares[deQ] {
					continue
				}

				cx, cy := deQ[0], deQ[1]

				for _, dir := range dirs {

					dx, dy := dir[0], dir[1]
					nx, ny := cx + dx, cy + dy

					newPoint := point{nx, ny}

					if grid.IsValidPos(mtx, nx, ny) && mtx[ny][nx] == c {

						queue = append(queue, newPoint)
						areaSquares[point{cx, cy}] = true
						visited[newPoint] = true

					} else {

						if v, ok := perimSquares[newPoint]; ok {
							perimSquares[newPoint] = v + 1
						} else {
							perimSquares[newPoint] = 1
						}

					}
				}
			}

			perim := 0
			area := max(len(areaSquares), 1)
			for _, n := range perimSquares {
				perim += n
			}

			sum += area * perim
			fmt.Printf("Region of %c has area %d and perimeter %d\n", c, area, perim)
		}
	}

	return conv.ToString(sum)
}

// this does not work. someone help me please lol
func countSides(areaSquares Set[point], perimSquares map[point]int) int{
	// we check the direct diagonal for each area square.
	/*
	if the diagonal is already inside areaSquares, don't count it.
	if the diagonal is not in the perimeter squares, count it.
	if the diagonal is inside the perimeter squares and the value is greater than one, count it.

	it should be noted that the value (the amount of area squares surronding the perimeter square)
	determines the number of corners.

	if it's 2: there is  1 corner  associated with it.
	if it's 3: there are 2 corners associated with it.
	if it's 4: there are 4 corners associated with it.
	*/

	diagDirs := [][2]int {
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	corners := map[point]int{}

	for areaSquare := range areaSquares {
		ax, ay := areaSquare[0], areaSquare[1]
		for _, dir := range diagDirs {
			cx, cy := ax + dir[0], ay + dir[1]
			cpoint := point{cx, cy}

			// skip over all squares in areasquare.
			if areaSquares[cpoint] {
				continue
			}

			if v, ok := perimSquares[cpoint]; !ok || (ok && v >= 2) {
				numCorners := 1
				switch v {
				case 2: numCorners = 1
				case 3: numCorners = 2 // for dips, there are two corners.
				case 4: numCorners = 4 // for holes, there are four corners.
				}
				corners[cpoint] = numCorners
			}
		}
	}

	sides := 0
	for _, v := range corners {
		sides += v
	}
	return sides
}

func part2(input string) string {
	mtx := parseInput(input)
	visited := make(Set[point])

	sum := 0
	for y, row := range mtx {

		for x, c := range row {

			p := point{x, y}

			if visited[p] {
				continue
			}

			visited[p] = true

			perimSquares := map[point]int{}
			areaSquares := make(Set[point])

			queue := []point{p}

			for len(queue) > 0 {
				deQ := queue[0]
				queue = queue[1:]

				if areaSquares[deQ] {
					continue
				}

				cx, cy := deQ[0], deQ[1]

				for _, dir := range dirs {

					dx, dy := dir[0], dir[1]
					nx, ny := cx + dx, cy + dy

					newPoint := point{nx, ny}

					if grid.IsValidPos(mtx, nx, ny) && mtx[ny][nx] == c {

						queue = append(queue, newPoint)
						areaSquares[point{cx, cy}] = true
						visited[newPoint] = true

					} else {

						if v, ok := perimSquares[newPoint]; ok {
							perimSquares[newPoint] = v + 1
						} else {
							perimSquares[newPoint] = 1
						}

					}
				}
			}

			areaSquares[p] = true

			area := len(areaSquares)
			sides := countSides(areaSquares, perimSquares)

			sum += area * sides
			fmt.Printf("Region of %c has area %d and %d sides\n", c, area, sides)
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