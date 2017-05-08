package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	cnt := 0
	for len(q) > 0 {
		cnt++
		nq := []*TreeNode{}
		for _, v := range q {
			if v.Left != nil {
				nq = append(nq, v.Left)
			}
			if v.Right != nil {
				nq = append(nq, v.Right)
			}
		}
		q = nq
	}
	return cnt
}
