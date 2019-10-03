package permutations

//给定一个没有重复数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 未完成 -------------------------

func permute(nums []int) (result [][]int) {
	var n = len(nums)
	if n < 1 {
		return
	}
	if n < 2 {
		result = append(result, nums)
		return
	}
	for i := 0; i < n; i++ {
		var ns []int
		ns = append(ns, nums[0:i]...)
		ns = append(ns, nums[i+1:]...)
		result = permute(ns)

		for j := 0; j < len(result); j++ {
		}
	}
	return
}
