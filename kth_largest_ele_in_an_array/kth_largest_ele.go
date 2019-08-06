package kth_largest_ele_in_an_array

import (
	"container/heap"
)

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//示例 1:
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//示例 2:
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
