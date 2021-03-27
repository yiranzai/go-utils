package leetcode

import (
	"testing"
)

//TestGenerateList test generate list
func Test_GenerateList(t *testing.T) {
	var head *ListNode

	head = GenerateList([]int{1, 2, 3, 4, 5})
	LoopEqualList(t, head, []int{1, 2, 3, 4, 5})

	head = GenerateList([]int{0, 1, 2})
	LoopEqualList(t, head, []int{0, 1, 2})

	head = GenerateList([]int{})
	LoopEqualList(t, head, []int{})
}

//TestGenerateTree test generate tree
func Test_GenerateTree(t *testing.T) {

	DeepEqualTree(t,
		GenerateTree([]interface{}{}),
		GenerateTree([]interface{}{}),
	)

	DeepEqualTree(t,
		GenerateTree([]interface{}{0}),
		GenerateTree([]interface{}{0}),
	)

	DeepEqualTree(t,
		GenerateTree([]interface{}{0, -10, 9, 20, nil, nil, 15, 7}),
		GenerateTree([]interface{}{0, -10, 9, 20, nil, nil, 15, 7}),
	)
}
