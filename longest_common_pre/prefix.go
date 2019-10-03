package longest_common_pre

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func longestCommonPrefix(strs []string) (pre string) {
	var n = len(strs)
	if n < 1 {
		return
	}
	if n == 1 {
		return strs[0]
	}
	var index int
	for {
		for i := 1; i < n; i++ {
			if index >= len(strs[i]) || index >= len(strs[0]) || strs[i][index] != strs[0][index] {
				return
			}
		}
		index++
		pre = strs[0][:index]
	}
}
