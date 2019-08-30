package max_sub

import (
	"math"
	"sync"
)

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// ------------------------------------------------
// 多协程分治解法，但是启动协程的时间和堆分配的时间加起来比算法的本身的时间还长。
// 跑10万次同一个数组的时间，多协程分治算法用700多毫秒，不分协程用50毫秒。
func maxSubArray(nums []int) (max int) {
	max, _, _ = mSA(nums)
	return
}

func mSA(nums []int) (max, h, t int) {
	var n = len(nums)
	if n == 1 {
		max = nums[0]
		h = max
		t = max
		return
	}
	if n <= 8 {
		max, h, t = maxSubHeadTail(nums)
	} else {
		var mid = (n - 1) / 2
		var ml, mr int
		var rh, lt int

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			ml, h, lt = maxSubHeadTail(nums[:mid])
			wg.Done()
		}()
		mr, rh, t = maxSubHeadTail(nums[mid:])

		wg.Wait()
		max = ml
		if mr > max {
			max = mr
		}
		if rh+lt > max {
			max = rh + lt
		}
	}

	return
}

func maxSubHeadTail(nums []int) (max int, h int, t int) {
	var n = len(nums)
	var acc int
	max = math.MinInt32
	t = math.MinInt32
	h = nums[0]
	var hAss int
	for i := 0; i < n; i++ {
		hAss += nums[i]
		if hAss > h {
			h = hAss
		}

		acc += nums[i]
		if acc > max {
			max = acc
		}
		t = acc
		if acc < 0 {
			acc = 0
		}
	}
	return
}

// ------------------------------------------------

// O(n)解法
func maxSubArray1(nums []int) int {
	var n = len(nums)
	var acc int
	var max = math.MinInt32
	for i := 0; i < n; i++ {
		acc += nums[i]
		if acc > max {
			max = acc
		}
		if acc < 0 {
			acc = 0
		}
	}
	return max
}
