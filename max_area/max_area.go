package max_area

func maxArea(heights []int) (max int) {
	var il, ir int
	var n = len(heights)
	ir = n - 1
	var low int
	for {
		if ir <= il {
			break
		}

		low = heights[il]
		if low > heights[ir] {
			low = heights[ir]
		}
		low = low * (ir - il)
		if max < low {
			max = low
		}

		if heights[il] < heights[ir] {
			il++
		} else {
			ir--
		}
	}
	return
}
