package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result []int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	path(root, 0)
	min := result[0]
	for _, v := range result {
		if v < min {
			min = v
		}
	}
	result = []int{}
	return min
}

func path(root *TreeNode, pre int) {
	if root.Left == nil &&
		root.Right == nil {
		result = append(result, pre+1)
		return
	}
	if root.Left != nil {
		path(root.Left, pre+1)
	}
	if root.Right != nil {
		path(root.Right, pre+1)
	}
	return
}
