package string_mul

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//说明：
//
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/multiply-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var Zero = "0"[0]

func multiply(num1 string, num2 string) (result string) {
	var n = len(num1)
	var j byte
	var sub byte
	var tail string
	result = "0"
	for i := n - 1; i >= 0; i-- { // 123
		sub = num1[i] - Zero
		for j = 0; j < sub; j++ {
			result = add(result, num2+tail)
		}
		tail += "0"
	}
	return
}

func add(num1, num2 string) (result string) {
	var maxN = len(num1)
	if maxN < len(num2) {
		var t = num1
		num1 = num2
		num2 = t
		maxN = len(num1)
	}
	var n2 = len(num2)
	var b = false
	var next byte
	for i := 1; i <= maxN; i++ {
		if i > n2 {
			next = num1[maxN-i] - Zero
		} else {
			next = num1[maxN-i] - Zero + num2[n2-i] - Zero
		}
		if b {
			next += 1
		}
		if next >= 10 {
			next = next - 10
			b = true
			result = string(next+Zero) + result
		} else {
			b = false
			result = string(next+Zero) + result
		}

	}
	if b {
		result = "1" + result
	} else {
		var index int
		for i := 0; i < len(result); i++ {
			if result[i] == Zero {
				index = i
			} else {
				break
			}
		}
		result = result[index:]
	}
	return
}
