package three_sum

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	var nums = []int{1, -1, -1, 0}
	fmt.Println(threeSum(nums))

	nums = []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(nums))
}
