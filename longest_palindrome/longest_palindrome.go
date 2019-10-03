package longest_palindrome

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 未完成 ------------------------------------------

func longestPalindrome(s string) (result string) {
	var n = len(s)
	if n < 2 {
		return s
	}
	var sPrime string
	var maxLen int
	for i := n - 1; i >= 0; i-- {
		sPrime += string(s[i])
	}

	var startIndex int
	var subLen int
	var tempB bool
	for i := 0; i < n; i++ { // 外层循环，原字符串下标
		for j := 0; j < n; j++ { // 内层循环，反转字符串的下标
			if s[i] != sPrime[j] {
				// 判断上一个子串是否回文
				subLen = j - startIndex
				if subLen > maxLen {
					tempB = true
					for k := 0; k < subLen/2-1; k++ {
						if sPrime[k] != sPrime[subLen-1-k] {
							tempB = false
							break
						}
					}
					if tempB {
						maxLen = subLen
						result = sPrime[startIndex:j]
					}
				}

				// 清除状态
				startIndex = -1
				continue
			}
			if startIndex < 0 {
				startIndex = j
			}

		}
	}
	return
}
