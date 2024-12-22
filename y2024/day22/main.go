package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
)

const PRUNE = 16777216

//go:embed in.txt
var input string

func parseInput(input string) []int {
	return conv.ToIntArr(util.Lines(input))
}

func step(secretNum int) int {
	secretNum = ((secretNum * 64) ^ secretNum) % PRUNE
	secretNum = ((secretNum / 32) ^ secretNum) % PRUNE
	secretNum = ((secretNum * 2048) ^ secretNum) % PRUNE
	return secretNum
}

func part1(input string) string {
	secretNums := parseInput(input)
	sum := 0
	for _, secretNum := range secretNums {
		for range 2000 {
			secretNum = step(secretNum)
		}
		sum += secretNum
	}
	return conv.ToString(sum)
}

func part2(input string) string {
	secretNums := parseInput(input)
	changeMap := map[[4]int]int{}
	for _, secretNum := range secretNums {
		changes := make([]int, 2000)
		currentPrice := secretNum % 10

		currentChangeMap := map[[4]int]int{}

		for i := range 2000 {
			secretNum = step(secretNum)
			newPrice := secretNum % 10
			diff := newPrice - currentPrice

			changes[i] = diff
			if i >= 3 {
				ch := [4]int(changes[i-3 : i+1])
				if _, ok := currentChangeMap[ch]; !ok {
					currentChangeMap[ch] = newPrice
				}
			}

			currentPrice = newPrice
		}
		for changeSeq, bananas := range currentChangeMap {
			if _, ok := changeMap[changeSeq]; !ok {
				changeMap[changeSeq] = 0
			}
			changeMap[changeSeq] += bananas
		}
	}
	maxBananas := -1
	for _, v := range changeMap {
		if v > maxBananas {
			maxBananas = v
		}
	}
	return conv.ToString(maxBananas)
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
