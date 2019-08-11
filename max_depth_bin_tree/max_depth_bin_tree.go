package max_depth_bin_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (result int) {
	if root == nil {
		return 0
	}
	result = maxDepth(root.Left)
	right := maxDepth(root.Right)
	if right > result {
		result = right
	}
	result += 1
	return
}
