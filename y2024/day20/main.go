package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/grid"
	"adventOfCode/util/maps"
	"adventOfCode/util/pq"
	_ "embed"
	"flag"
	"fmt"
)

//go:embed in.txt
var input string

const wall = '#'

type Set[T comparable] map[T]bool

type racer struct {
	pos   grid.Point
	dir   grid.Point
	cost  int
	rPath map[grid.Point]int
}

func parseInput(input string) (maze [][]rune, paths Set[grid.Point], start, end grid.Point) {
	paths = make(Set[grid.Point])
	lines := util.Lines(input)
	maze = make([][]rune, len(lines))
	for y, row := range lines {
		maze[y] = []rune(row)
		for x, c := range row {
			p := grid.Point{x, y}
			switch c {
			case 'S':
				start = p
			case 'E':
				end = p
			case wall:
				continue
			}
			paths[p] = true
		}
	}
	return
}

func allowedDirs(dir grid.Point) []grid.Point {
	switch dir {
	case grid.Up, grid.Down:
		return []grid.Point{dir, grid.Left, grid.Right}
	case grid.Left, grid.Right:
		return []grid.Point{dir, grid.Up, grid.Down}
	}
	return []grid.Point{}
}

func (r racer) getNext(paths Set[grid.Point], visited Set[grid.Point]) []racer {
	var possible []racer
	for _, dir := range allowedDirs(r.dir) {
		newPos := r.pos.Add(dir)
		if !paths[newPos] || visited[newPos] {
			continue
		}
		possible = append(possible, racer{
			pos:   newPos,
			dir:   dir,
			cost:  r.cost + 1,
			rPath: maps.Copy(r.rPath),
		})
	}
	return possible
}

func dijkstras(paths Set[grid.Point], start, end grid.Point) (int, map[grid.Point]int) {
	priorityQueue := pq.NewPQ(func(a racer, b racer) bool {
		return a.cost < b.cost
	})
	priorityQueue.Push(racer{start, grid.Up, 0, make(map[grid.Point]int)})
	visited := make(Set[grid.Point])

	for priorityQueue.Len() > 0 {
		popped := priorityQueue.Pop()

		if visited[popped.pos] {
			continue
		}
		visited[popped.pos] = true

		popped.rPath[popped.pos] = popped.cost
		if popped.pos == end {
			return popped.cost, popped.rPath
		}

		for _, next := range popped.getNext(paths, visited) {
			priorityQueue.Push(next)
		}
	}

	return -1, make(map[grid.Point]int)
}

type cheat struct {
	start grid.Point
	end   grid.Point
}

func calcCheats(dPath map[grid.Point]int, leastGain, maxDist int) int {
	uniqueCheats := make(map[cheat]int)
	count := 0
	for p1, d1 := range dPath {
		for p2, d2 := range dPath {
			if p1 == p2 {
				continue
			}
			if d2-d1-p1.Dist(p2) >= leastGain && p1.Dist(p2) <= maxDist {
				uniqueCheats[cheat{p1, p2}] = d2 - d1
				count++
			}
		}
	}
	return count
}

func part1(input string) string {
	_, paths, start, end := parseInput(input)
	cost, dPath := dijkstras(paths, start, end)
	fmt.Println(cost)
	return conv.ToString(calcCheats(dPath, 100, 2))
}

func part2(input string) string {
	_, paths, start, end := parseInput(input)
	cost, dPath := dijkstras(paths, start, end)
	fmt.Println(cost)
	return conv.ToString(calcCheats(dPath, 100, 20))
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
