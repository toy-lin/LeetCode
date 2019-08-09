package subsets

import "math"

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/subsets
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func subsets(nums []int) [][]int {
	var n = len(nums)
	var result [][]int
	//通过遍历000 - 111二进制的方法遍历所有子集
	all := int(math.Pow(2, float64(n)))
	for i := 0; i < all; i++ { //遍历所有集合
		result = append(result, []int{})
		bs := int(math.Ceil(math.Log2(float64(i + 1))))
		for bit := 0; bit < bs; bit++ { //遍历所有二进制位
			j := i >> uint(bit)
			if j%2 == 1 {
				result[i] = append(result[i], nums[bit])
			}
		}
	}
	return result
}
