package rm_dup

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	var nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
