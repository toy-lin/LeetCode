package palindrome_num

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var result int
	var origin = x
	for {
		if x == 0 {
			return origin == result
		}
		result = 10*result + x%10
		x = x / 10
	}
}
