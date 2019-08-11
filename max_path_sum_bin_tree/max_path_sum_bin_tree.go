package max_path_sum_bin_tree

import "math"

//给定一个非空二叉树，返回其最大路径和。
//
//本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//示例 1:
//
//输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
//输出: 6
//示例 2:
//
//输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt32
	result := maxGain(root)
	return int(math.Max(float64(maxSum), float64(result)))
}

func maxGain(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// get max compare to 0 to drop negative gain
	left := int(math.Max(float64(maxGain(node.Left)), 0))
	right := int(math.Max(float64(maxGain(node.Right)), 0))
	sumNewPath := left + right + node.Val
	maxSum = int(math.Max(float64(sumNewPath), float64(maxSum)))
	return node.Val + int(math.Max(float64(left), float64(right)))
}
