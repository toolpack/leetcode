package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var total int

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	total = sum
	return run(root, 0)
}

func run(root *TreeNode, cur int) bool {
	if root.Left == nil && root.Right == nil {
		if cur+root.Val == total {
			return true
		}
	}

	if root.Left != nil {
		if run(root.Left, cur+root.Val) {
			return true
		}
	}

	if root.Right != nil {
		if run(root.Right, cur+root.Val) {
			return true
		}
	}
	return false
}
