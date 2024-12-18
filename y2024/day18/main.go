package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	"adventOfCode/util/pq"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

type Set[T comparable] map[T]bool

type weightedPoint struct {
	p    grid.Point
	cost int
}

const byteChar = '#'

var dirs = [4]grid.Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

//go:embed in.txt
var input string

func parseInput(input string, maxLines int) [][]rune {
	mtx := make([][]rune, 71)
	for i := range mtx {
		mtx[i] = make([]rune, 71)
		for j := range mtx[i] {
			mtx[i][j] = '.'
		}
	}
	lines := util.Lines(input)
	if maxLines < 0 || maxLines > len(lines) {
		maxLines = len(lines)
	}
	for _, ln := range util.Lines(input)[:maxLines] {
		bytePos := conv.ToIntArr(strings.Split(ln, ","))
		x, y := bytePos[0], bytePos[1]
		mtx[y][x] = byteChar
	}
	return mtx
}

func dijkstra(mtx [][]rune) int {
	end := grid.Point{70, 70}
	queue := pq.NewPQ(func(p weightedPoint, o weightedPoint) bool {
		return p.cost < o.cost
	})
	queue.Push(weightedPoint{
		p:    grid.Point{0, 0},
		cost: 0,
	})
	visited := Set[grid.Point]{}
	for queue.Len() > 0 {
		popped := queue.Pop()
		if popped.p == end {
			return popped.cost
		}

		if visited[popped.p] {
			continue
		}
		visited[popped.p] = true

		if ok, v := grid.At(mtx, popped.p); !ok || (ok && v == byteChar) {
			continue
		}

		for _, dir := range dirs {
			queue.Push(weightedPoint{
				p:    popped.p.Add(dir),
				cost: popped.cost + 1,
			})
		}
	}
	return -1
}

func part1(input string) string {
	mtx := parseInput(input, 1024)
	return conv.ToString(dijkstra(mtx))
}

func part2(input string) string {
	res := "-1, -1"
	mtx := parseInput(input, -1)
	lines := util.Lines(input)
	for i := len(lines) - 1; i >= 0; i-- {
		ln := lines[i]
		bytePt := conv.ToIntArr(strings.Split(ln, ","))
		x, y := bytePt[0], bytePt[1]
		mtx[y][x] = '.'
		if v := dijkstra(mtx); v >= 0 {
			res = ln
			break
		}
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
