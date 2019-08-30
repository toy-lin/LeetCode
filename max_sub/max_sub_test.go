package max_sub

import (
	"fmt"
	"testing"
	"time"
)

func TestF(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))

	nums = []int{-2, -1}
	fmt.Println(maxSubArray(nums))

	nums = []int{-2, -3, -1}
	fmt.Println(maxSubArray(nums))

	nums = []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func TestFF(t *testing.T) {
	nums := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}

	var times = 1000000
	start := time.Now()
	for i := 0; i < times; i++ {
		maxSubArray(nums)
	}
	fmt.Println(time.Now().Sub(start))
}
