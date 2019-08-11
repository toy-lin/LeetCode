package max_path_sum_bin_tree

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 0,
	}
	fmt.Println(maxPathSum(root))
}
