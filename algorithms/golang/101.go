package main

import (
	"fmt"
)

func main() {
	n7 := &TreeNode{3, nil, nil}
	// n6 := &TreeNode{4, nil, nil}
	// n5 := &TreeNode{4, nil, nil}
	n4 := &TreeNode{3, nil, nil}
	n3 := &TreeNode{2, n7, nil}
	n2 := &TreeNode{2, nil, n4}
	n1 := &TreeNode{1, n2, n3}
	fmt.Println(isSymmetric(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*TreeNode{root}
	count := 1
	for len(q) > 0 {
		if count == 0 {
			return true
		} else {
			count = 0
		}
		nq := []*TreeNode{}
		for k, v := range q {
			if k < len(q)/2 {
				if q[k] == nil && q[len(q)-1-k] != nil {
					return false
				}
				if q[k] != nil && q[len(q)-1-k] == nil {
					return false
				}
				if q[k] == nil && q[len(q)-1-k] == nil {
					continue
				}
				if q[k].Val != q[len(q)-1-k].Val {
					return false
				}
			}
			if v != nil {
				nq = append(nq, v.Left)
				nq = append(nq, v.Right)
				if v.Left != nil {
					count++
				}
				if v.Right != nil {
					count++
				}
			}
		}
		q = nq
	}
	return true
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil && r != nil {
		return false
	}
	if l != nil && r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return (helper(l.Left, r.Right) && helper(l.Right, r.Left))
}
