package kth_largest_ele_in_an_array

import "sort"

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return 0
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	return temp[len(temp)-k]
}
