package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	"container/heap"
	_ "embed"
	"flag"
	"fmt"
	"math"
)

//go:embed in.txt
var input string

// point is still [x, y].
type point [2]int
type direction point

const (
	wall  = '#'
	path  = '.'
	start = 'S'
	end   = 'E'
)

var (
	up    = direction{0, -1}
	left  = direction{-1, 0}
	down  = direction{0, 1}
	right = direction{1, 0}
)

// -----------------------------------

func (p point) add(o point) point {
	return point{p[0] + o[0], p[1] + o[1]}
}

func (p point) addDirection(d direction) point {
	return point{p[0] + d[0], p[1] + d[1]}
}

func (p point) on(grid [][]rune) rune {
	return grid[p[1]][p[0]]
}

func (d direction) getOffsets() [3]direction {
	switch d {
	case up, down:
		return [3]direction{
			d, left, right,
		}
	case left, right:
		return [3]direction{
			d, up, down,
		}
	}
	panic("Invalid direction")
}

func (d direction) toRune() rune {
	switch d {
	case up:
		return '^'
	case left:
		return '<'
	case right:
		return '>'
	case down:
		return 'v'
	}
	panic("Invalid direction")
}

//----------------------------------------------------

type directedPoint struct {
	pos point
	dir direction
}

type routeState struct {
	reindeer directedPoint
	path     []directedPoint
	cost     int
}

//----------------------------------------------------

type routePQ []routeState

func (r routePQ) Len() int           { return len(r) }
func (r routePQ) Less(i, j int) bool { return r[i].cost < r[j].cost }
func (r routePQ) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r *routePQ) Push(x any) {
	*r = append(*r, x.(routeState))
}

func (r *routePQ) Pop() any {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}

//----------------------------------------------------

func parseInput(input string) (grid [][]rune, startPt directedPoint, endPt point) {
	grid = arrays.Map(util.Lines(input), func(ln string) []rune {
		return []rune(ln)
	})
	for y, row := range grid {
		for x, c := range row {
			switch c {
			case start:
				startPt = directedPoint{point{x, y}, right}
			case end:
				endPt = point{x, y}
			case wall, path:
				continue
			}
			grid[y][x] = path
		}
	}
	return
}

func findBestRoute(maze [][]rune, startPt directedPoint, endPt point) (minCost, totalSteps int) {
	minCost = math.MaxInt
	queue := routePQ{
		{startPt, []directedPoint{startPt}, 0},
	}
	heap.Init(&queue)
	visited := make(map[directedPoint]int)
	bestRoutes := make(map[int][]directedPoint)

	for len(queue) > 0 {
		state := heap.Pop(&queue).(routeState)
		if len(state.path) > 10_000 {
			continue
		}

		// because we are working with a min heap, we can immediately stop
		// when the cost is higher.
		if state.cost > minCost {
			break
		}

		if state.reindeer.pos == endPt {
			if state.cost <= minCost {
				bestRouteTmp := bestRoutes[state.cost]
				bestRouteTmp = append(bestRouteTmp, state.path...)
				bestRoutes[state.cost] = bestRouteTmp
				minCost = state.cost
				continue
			}
		}

		for _, d := range state.reindeer.dir.getOffsets() {
			newPos := state.reindeer.pos.addDirection(d)
			if newPos.on(maze) == wall {
				continue
			}
			candidate := directedPoint{
				pos: newPos,
				dir: d,
			}
			cost := state.cost + 1
			if candidate.dir != state.reindeer.dir {
				cost += 1000
			}
			if previous, found := visited[candidate]; found && previous < cost {
				continue
			}
			visited[candidate] = cost

			newPath := make([]directedPoint, len(state.path))
			copy(newPath, state.path)

			heap.Push(&queue, routeState{
				reindeer: candidate,
				path:     append(newPath, candidate),
				cost:     cost})
		}
	}
	buffer := make(map[point]int)
	for _, v := range bestRoutes[minCost] {
		buffer[v.pos]++
	}
	return minCost, len(buffer)
}

func printMaze(maze [][]rune, route routeState) {
	fmt.Printf("Route Found: %d\n", route.cost)

	getStep := func(p point) (rune, bool) {
		for i, step := range route.path {
			if step.pos == p {
				switch i {
				case 0:
					return 'S', true
				case len(route.path) - 1:
					return 'E', true
				default:
					return step.dir.toRune(), true
				}
			}
		}

		return 0, false
	}

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			current := point{x, y}
			value := current.on(maze)

			if r, found := getStep(current); found {
				value = r
			}

			fmt.Printf("%s", string(value))

		}
		fmt.Println()
	}

	fmt.Println()
}

func part1(input string) string {
	maze, startPt, endPt := parseInput(input)
	bestScore, _ := findBestRoute(maze, startPt, endPt)
	return conv.ToString(bestScore)
}

func part2(input string) string {
	maze, startPt, endPt := parseInput(input)
	_, tiles := findBestRoute(maze, startPt, endPt)
	return conv.ToString(tiles)
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
