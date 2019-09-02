package three_sum

import (
	"math"
	"sort"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSum(nums []int) (result [][]int) {
	var n = len(nums)
	if n < 3 {
		return
	}
	sort.Ints(nums)

	var j, k, s int
	var jt, kt int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j = i + 1
		k = n - 1

		jt = math.MinInt32
		kt = math.MinInt32
		for {
			if k <= j {
				break
			}
			if nums[j] == jt {
				j++
				continue
			}
			if nums[k] == kt {
				k--
				continue
			}
			s = nums[i] + nums[j] + nums[k]
			if s < 0 {
				j++
				continue
			}
			if s > 0 {
				k--
				continue
			}
			if s == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				jt = nums[j]
				kt = nums[k]
				j++
				k--
				continue
			}
		}
	}
	return
}
