package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"adventOfCode/util/maps"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
)

//go:embed in.txt
var input string

type adjList map[string][]string

func intersection(a, b []string) (res []string) {
	set := make(map[string]bool)
	for _, s := range b {
		set[s] = true
	}
	for _, s := range a {
		if set[s] {
			res = append(res, s)
		}
	}
	return
}

func BronKerbosch(R, P, X []string, maxLenClique *[]string, maxLen *int, al adjList) {
	if len(P) == 0 && len(X) == 0 && *maxLen < len(R) {
		rCopy := slices.Clone(R)
		*maxLenClique = rCopy
		*maxLen = len(rCopy)
		return
	}

	pCopy := slices.Clone(P)
	for _, v := range pCopy {
		newR := append(R, v)
		neighbours := al[v]

		newP := intersection(P, neighbours)
		newX := intersection(X, neighbours)

		BronKerbosch(newR, newP, newX, maxLenClique, maxLen, al)

		vIdx := slices.Index(P, v)
		P = slices.Delete(P, vIdx, vIdx+1)

		X = append(X, v)
	}
}

func parseInput(input string) adjList {
	res := make(adjList)
	for _, ln := range util.Lines(input) {
		parts := strings.Split(ln, "-")
		a, b := parts[0], parts[1]
		res[a] = append(res[a], b)
		res[b] = append(res[b], a)
	}
	return res
}

func part1(input string) string {
	graph := parseInput(input)
	cliques := make(adjList)
	for n1 := range graph {
		if n1[0] != 't' {
			continue
		}
		for _, n2 := range graph[n1] {
			for _, n3 := range graph[n2] {
				if slices.Contains(graph[n3], n1) {
					clique := []string{n1, n2, n3}
					slices.Sort(clique)
					key := strings.Join(clique, "")
					cliques[key] = clique
				}
			}
		}
	}
	return conv.ToString(len(cliques))
}

func part2(input string) string {
	graph := parseInput(input)
	nodesList := maps.Keys(graph)

	var maxLenClique []string
	var maxLen int

	BronKerbosch([]string{}, nodesList, []string{}, &maxLenClique, &maxLen, graph)

	slices.Sort(maxLenClique)
	return strings.Join(maxLenClique, ",")
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
