package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		newq := []*TreeNode{}
		level := []int{}
		for _, v := range queue {
			if v.Left != nil {
				newq = append(newq, v.Left)
			}
			if v.Right != nil {
				newq = append(newq, v.Right)
			}
			level = append(level, v.Val)
		}
		if len(level) != 0 {
			r = append([][]int{level}, r...)
		}
		queue = newq
	}
	return r
}
