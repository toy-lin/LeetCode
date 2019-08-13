package lowest_common_ancestor_bst

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	root := TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
		},
	}

	//fmt.Println(lowestCommonAncestor(&root, root.Left.Left, root.Left.Right))

	root = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Println(lowestCommonAncestor(&root, root.Left.Right, root.Right))
}
