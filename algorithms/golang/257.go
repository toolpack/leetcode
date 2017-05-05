package main

import (
	"fmt"
)

func main() {
	binaryTreePaths(nil)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	subBinaryTreePaths(root, "")
	return data
}

var data []string

func subBinaryTreePaths(root *TreeNode, answer string) string {
	if root == nil {
		return ""
	}
	if answer != "" {
		answer = answer + "->" + fmt.Sprintf("%v", root.Val)
	} else {
		answer = fmt.Sprintf("%v", root.Val)
	}
	if root.Left == nil && root.Right == nil {
		data = append(data, answer)
		return ""
	}

	if root.Left != nil {
		subBinaryTreePaths(root.Left, answer)
	}
	if root.Right != nil {
		subBinaryTreePaths(root.Right, answer)
	}
	return ""
}
