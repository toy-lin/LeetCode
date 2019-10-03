package reverse_words

import "strings"

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//示例 1:
//
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var Space = " "[0]

func reverseWords(s string) (result string) {
	var n = len(s)
	if n < 1 {
		return
	}
	var indexStart int
	var end int
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if s[i] == Space || i == n-1 {
			// reverse words
			if s[i] == Space {
				end = i - 1
			} else {
				end = i
			}

			for j := end; j >= indexStart; j-- {
				sb.WriteByte(s[j])
			}

			if s[i] == Space {
				sb.WriteByte(Space)
			}

			// reset states
			indexStart = i + 1
			continue
		}
	}
	result = sb.String()
	return
}
