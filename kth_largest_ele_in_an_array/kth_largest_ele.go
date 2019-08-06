package kth_largest_ele_in_an_array

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	temp := make([]int, k)
	//copy(temp, nums[:k])
	for i := 0; i < k; i++ {
		temp[i] = nums[i]
	}
	h := intHeap(temp)
	heap.Init(&h)

	for i := k; i < len(nums); i++ {
		if nums[i] > h[0] {
			h[0] = nums[i]
			heap.Fix(&h, 0)
		}
	}
	return h[0]
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return (h)[i] < (h)[j]
}

func (h intHeap) Swap(i, j int) {
	(h)[i], (h)[j] = (h)[j], (h)[i]
}

func (intHeap) Push(x interface{}) {}
func (intHeap) Pop() interface{}   { return nil }
