package main

import (
	"fmt"
)

func main() {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{1, nil, nil}
	fmt.Println(isSameTree(n1, n2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	v1 := helper(p, []int{})
	v2 := helper(q, []int{})
	if len(v1) != len(v2) {
		return false
	}
	for k, v := range v1 {
		if v != v2[k] {
			return false
		}
	}
	return true
}

func helper(root *TreeNode, pre []int) []int {
	if root == nil {
		return []int{}
	}
	r := []int{}
	r = append(pre, root.Val)
	if root.Left != nil {
		r = helper(root.Left, r)
	}
	if root.Right != nil {
		r = helper(root.Right, r)
	}
	return r
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
