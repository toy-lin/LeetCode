package three_sum_closest

import (
	"math"
	"sort"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
//例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum-closest
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSumClosest(nums []int, target int) (cSum int) {
	var n = len(nums)
	if n < 3 {
		return
	}
	sort.Ints(nums)

	var cDis = math.MaxInt32
	var i, j, k int
	var sum int
	var dis int
	for i = 0; i < n-2; i++ {
		j = i + 1
		k = n - 1

		for {
			if k <= j {
				break
			}
			sum = nums[i] + nums[j] + nums[k]
			dis = abs(sum - target)

			if dis < cDis {
				cSum = sum
				cDis = dis
				if dis == 0 {
					return
				}
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
