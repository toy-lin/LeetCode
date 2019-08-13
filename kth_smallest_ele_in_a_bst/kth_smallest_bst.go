package kth_smallest_ele_in_a_bst

//给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
//
//说明：
//你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
//
//示例 1:
//
//输入: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//   2
//输出: 1
//示例 2:
//
//输入: root = [5,3,6,2,4,null,null,1], k = 3
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1
//输出: 3
//进阶：
//如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var K int
var index int
var kSmallest int

//题解：二叉搜索树的每个节点的左子树所有节点的值必定小于这个节点，右子树所有节点的值必定大于这个节点。
//基于二叉搜索树上述的特点，采用中序遍历可以解决第K小或第K大的问题。
//
//首先对数进行中序遍历，当经过中间节点时，全局索引+1。当全局所以和k相等时，该节点即为题目要求的节点。
//最坏情况下，需要遍历整棵树，时间复杂度为O(n)。
// 由于是递归调用，故使用最大空间为树的最大深度，可能达到O(n)，平均空间复杂度应为O(logN)。
func kthSmallest(root *TreeNode, k int) int {
	K = k
	index = 0
	ldr(root)
	return kSmallest
}

func ldr(node *TreeNode) {
	if node == nil {
		return
	}
	ldr(node.Left)
	index += 1
	if index == K {
		kSmallest = node.Val
		return
	} else if index >= K {
		return
	}
	ldr(node.Right)
}
