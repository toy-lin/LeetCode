package lowest_common_ancestor_bst

//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
//例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
//
//
//
//
//
//示例 1:
//
//输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
//解释: 节点 2 和节点 8 的最近公共祖先是 6。
//示例 2:
//
//输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//输出: 2
//解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
//
//说明:
//
//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉搜索树中。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ancestor *TreeNode

//题解：可以采用后续遍历的方法。步骤如下：
//1.遍历左子树，返回匹配节点的数量。
//2.遍历右子树，返回匹配节点的数量。
//3.如果左右子树各有1个匹配，则当前节点就是最近的祖先节点。
//4.如果左右子树只有1个匹配，查看当前节点是否匹配，若匹配则当前节点是最近的祖先节点。
//5.返回当前节点接收到的匹配总数。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor = nil
	lrd(root, p, q)
	return ancestor
}

func lrd(node, p, q *TreeNode) (matchCount int) {
	if node == nil {
		return
	}
	matchLeft := lrd(node.Left, p, q)
	matchRight := lrd(node.Right, p, q)
	if matchLeft == 2 || matchRight == 2 {
		return 2
	}
	if matchLeft == 1 && matchRight == 1 {
		ancestor = node
		return 2
	}
	if node == p || node == q {
		if matchLeft == 1 || matchRight == 1 {
			ancestor = node
			return 2
		}
		return 1
	}
	return matchLeft + matchRight
}
