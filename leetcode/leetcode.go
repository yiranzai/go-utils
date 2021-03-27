package leetcode

import (
	"reflect"
	"testing"

	"gotest.tools/v3/assert"
)

//TreeNode Leetcode Tree Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//ListNode Leetcode List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * GenerateTree generate leetcode tree
 * e.g: []int{0, 1, 2, 2, 3, 3, nil, nil, 4, 4}
 *	           1
 *  		  / \
 * 		     2  2
 *          / \
 *	       3  3
 *	     / \
 *      4  4
 */
func GenerateTree(treeList []interface{}) *TreeNode {
	var root *TreeNode
	if len(treeList) <= 1 {
		return root
	}
	treeNodeList := make([]*TreeNode, len(treeList))
	for i := 1; i < len(treeList); i++ {
		v := reflect.ValueOf(treeList[i])
		if v.Kind() != reflect.Int {
			continue
		}
		node := &TreeNode{Val: reflect.ValueOf(treeList[i]).Interface().(int)}
		treeNodeList[i] = node
		rootIndex := i >> 1
		if rootIndex == 0 {
			continue
		}
		if i%2 == 1 {
			treeNodeList[rootIndex].Right = node
		} else {
			treeNodeList[rootIndex].Left = node
		}
	}
	return treeNodeList[1]
}

//GenerateList generate leetcode list
func GenerateList(list []int) *ListNode {
	var head *ListNode
	node := head
	for _, v := range list {
		t := &ListNode{Val: v}
		if head == nil {
			head = t
			node = t
			continue
		} else {
			node.Next = t
			node = node.Next
		}
	}
	return head
}

//LoopEqualList test diff list and array
func LoopEqualList(t *testing.T, head *ListNode, list []int) {
	for _, v := range list {
		assert.Equal(t, v, head.Val)
		head = head.Next
	}
	assert.Equal(t, true, reflect.ValueOf(head).IsValid())
}

//DeepEqualTree test recursion diff tree
func DeepEqualTree(t *testing.T, head, right *TreeNode) {
	assert.Equal(t, reflect.ValueOf(head).IsValid(), reflect.ValueOf(right).IsValid())
	assert.Equal(t, head == nil, right == nil)
	if reflect.ValueOf(head).IsValid() == false || reflect.ValueOf(right).IsValid() == false || head == nil || right == nil {
		return
	}
	assert.Equal(t, head.Val, right.Val)
	DeepEqualTree(t, head.Left, right.Left)
	DeepEqualTree(t, head.Right, right.Right)
}
