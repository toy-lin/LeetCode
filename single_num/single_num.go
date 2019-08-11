package single_num

func singleNumber(nums []int) (result int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		result ^= nums[i]
	}
	return
}
