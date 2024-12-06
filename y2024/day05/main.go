package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
)

type Pair [2]int

type Set[T comparable] map[T]bool

//go:embed in.txt
var input string

func parseInput(input string) (Set[Pair], [][]int) {
	rules := make(Set[Pair])
	var orders [][]int

	parts := strings.Split(input, "\n\n")
	for _, rule := range strings.Split(parts[0], "\n") {
		arr := conv.ToIntArr(strings.Split(rule, "|"))
		rules[Pair{arr[0], arr[1]}] = true
	}

	for _, order := range strings.Split(parts[1], "\n") {
		orders = append(orders, conv.ToIntArr(strings.Split(order, ",")))
	}

	return rules, orders
}

func solve(input string) (int, int) {
	rules, orders := parseInput(input)

	correct, incorrect := 0, 0
	for _, order := range orders {

		sorted := sort(order, rules)
		middle := sorted[len(sorted)/2]
		if slices.Equal(order, sorted) {
			correct += middle
		} else {
			incorrect += middle
		}
	}

	return correct, incorrect
}

func sort(order []int, rules Set[Pair]) []int {
	clone := slices.Clone(order)
	slices.SortStableFunc(clone, func(a, b int) int {
		if v, ok := rules[Pair{a, b}]; ok && v {
			return -1
		}
		return 1
	})

	return clone
}

func part1(input string) string {
	correct, _ := solve(input)
	return conv.ToString(correct)
}

func part2(input string) string {
	_, incorrect := solve(input)
	return conv.ToString(incorrect)
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
