package merge_sorted_array

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	var a = []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{4, 5, 6}, 3)
	fmt.Println(a)

	a = []int{2, 0}
	merge(a, 1, []int{1}, 1)
	fmt.Println(a)
}
