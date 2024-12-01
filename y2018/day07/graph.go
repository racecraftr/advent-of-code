package main

import (
	"adventOfCode/util/conv"
	"sort"
)

const startRune = 'T' // based off of my input

type Graph struct {
	edges map[rune]map[rune]bool
}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[rune]map[rune]bool),
	}
}

func (g *Graph) AddEdge(src rune, dest rune) {
	if g.edges[src] == nil {
		g.edges[src] = make(map[rune]bool)
	}
	g.edges[src][dest] = true
}

func (g *Graph) Traverse() string {
	current, visited, res := startRune, map[rune]bool{}, ""
	g.traverse(current, &visited, &res)
	return res
}

func (g *Graph) traverse(current rune, visited *map[rune]bool, res *string) {
	if (*visited)[current] {
		return
	}

	*res += conv.ToString(current)

	neighbors := g.edges[current]
	neighborArr := []rune{}
	for neighbor := range neighbors {
		neighborArr = append(neighborArr, neighbor)
	}

	sort.Slice(neighborArr, func(i, j int) bool {
		return neighborArr[i] < neighborArr[j]
	})

	for _, neighbor := range neighborArr {
		g.traverse(neighbor, visited, res)
	}
}
