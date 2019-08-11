package majority_ele

func majorityElement(nums []int) (result int) {
	n := len(nums)
	candidate := 0
	for i := 0; i < n; i++ {
		if candidate == 0 {
			result = nums[i]
			candidate = 1
			continue
		}
		if result == nums[i] {
			candidate += 1
		} else {
			candidate -= 1
		}
	}
	return
}
