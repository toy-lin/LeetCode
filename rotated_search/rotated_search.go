package rotated_search

//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//示例 2:
//
//输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func search(nums []int, target int) (result int) {
	result = -1
	var n = len(nums)
	if n == 0 {
		return -1
	}
	var k = n - 1
	var i = k / 2
	var j = 0
	for {
		if target == nums[i] {
			result = i
			return
		}
		if j >= k {
			return
		}
		if target < nums[i] {
			if target == nums[j] {
				return j
			}
			if (nums[i] > nums[j] && target > nums[j]) || nums[i] < nums[j] {
				k = i - 1
			} else {
				j = i + 1
			}
			i = (j + k) / 2
			continue
		}
		if target == nums[k] {
			return k
		}
		if (nums[i] < nums[k] && target < nums[k]) || nums[i] > nums[k] {
			j = i + 1
		} else {
			k = i - 1
		}
		i = (i + k) / 2
		continue
	}
}
