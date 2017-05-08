package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// if root.Left == nil&&
	//    root.Right==nil{
	//      return true
	//    }
	return cal(root) != -1
}

func cal(root *TreeNode) (l int) {
	if root == nil {
		return 0
	}
	var left, right int
	left = cal(root.Left)
	if left == -1 {
		return -1
	}

	right = cal(root.Right)
	if right == -1 {
		return -1
	}

	if left-right > 1 {
		return -1
	}
	if right-left > 1 {
		return -1
	}

	ret := 0
	if left >= right {
		ret = left + 1
	}
	if right > left {
		ret = right + 1
	}
	return ret
}
