package main

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed in.txt
var input string

func inputToArr(input string) []int {
	sArr := strings.Split(input, " ")
	arr := make([]int, len(sArr))
	for i, s := range sArr {
		arr[i] = conv.ToInt(s)
	}
	return arr
}

type TreeNode struct {
	children []*TreeNode
	metadata []int
}

func makeTree(nums []int) (node *TreeNode, recursiveValuesHandled int) {
	if len(nums) == 0 {
		return nil, 0
	}
	if len(nums) == 2 {
		return &TreeNode{nil, []int{}}, 2
	}

	childrenCount := nums[0]
	metadataCount := nums[1]

	newNode := TreeNode{}

	valuesHandled := 2
	for i := 2; childrenCount > 0 || metadataCount > 0; {
		if childrenCount > 0 {
			// recursively make child
			child, subValuesHandled := makeTree(nums[i:])

			newNode.children = append(newNode.children, child)
			valuesHandled += subValuesHandled

			childrenCount--

			i += subValuesHandled
		} else {
			newNode.metadata = append(newNode.metadata, nums[i])
			valuesHandled++

			metadataCount--
			i++
		}
	}

	return &newNode, valuesHandled
}

func sumMetadata(tree *TreeNode) int {

	sum := 0

	for _, v := range tree.metadata {
		sum += v
	}

	for _, child := range tree.children {
		sum += sumMetadata(child)
	}

	return sum
}

func sumValue(tree *TreeNode) int {
	sum := 0

	if len(tree.children) == 0 {
		for _, v := range tree.metadata {
			sum += v
		}
		return sum
	}

	for _, val := range tree.metadata {
		idx := val - 1
		if idx < len(tree.children) {
			sum += sumValue(tree.children[idx])
		}
	}

	return sum
}

func part1(input string) string {
	tree, _ := makeTree(inputToArr(input))
	return conv.ToString(sumMetadata(tree))
}

func part2(input string) string {
	tree, _ := makeTree(inputToArr(input))
	return conv.ToString(sumValue(tree))
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
