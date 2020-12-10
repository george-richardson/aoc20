package day10_adapter_array

import (
	"sort"
)

func findAdapterDifferential(input []int) int {
	sort.Ints(input)
	ones, threes := 0, 1
	if input[0] == 1 {
		ones++
	} else {
		threes++
	}
	for i := 1; i < len(input); i++ {
		if input[i]-input[i-1] == 1 {
			ones++
		} else {
			threes++
		}
	}

	return ones * threes
}

type treeNode struct {
	visited  bool
	children []int
	value    int
}

func countAdapterPermutations(input []int) (count int) {
	sort.Ints(input)
	input = append([]int{0}, append(input, input[len(input)-1]+3)...)

	count = 1
	tree := make(map[int]*treeNode)
	for i := 0; i < len(input); i++ {
		tree[input[i]] = &treeNode{}
		for j := i + 1; j < len(input) && j <= i+3; j++ {
			if input[j]-input[i] <= 3 {
				tree[input[i]].children = append(tree[input[i]].children, input[j])
			} else {
				continue
			}
		}
	}
	return recurseCountTree(tree, 0)
}

func recurseCountTree(tree map[int]*treeNode, start int) int {
	node := tree[start]
	if !node.visited {
		node.visited = true
		if len(node.children) == 0 {
			node.value = 1
			return node.value
		} else {
			for _, i := range node.children {
				node.value += recurseCountTree(tree, i)
			}
		}
	}
	return node.value
}
