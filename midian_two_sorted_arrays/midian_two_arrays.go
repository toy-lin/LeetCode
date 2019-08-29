package midian_two_sorted_arrays

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var m = len(nums1)
	var n = len(nums2)
	if n < m {
		var temp = nums1
		nums1 = nums2
		nums2 = temp
		m = len(nums1)
		n = len(nums2)
	}
	var midM = (m - 1) / 2
	var midN = (n - 1) / 2

	if m == 0 {
		if n%2 == 1 {
			return float64(nums2[midN])
		}
		return float64(nums2[midN]+nums2[midN+1]) / 2
	}

	if m == 1 {
		var num = nums1[0]
		if n == 1 {
			return float64(num+nums2[0]) / 2
		}
		if n%2 == 1 {
			var nMid = nums2[midN]
			var nLeft = nums2[midN-1]
			var nRight = nums2[midN+1]
			if num < nLeft {
				return float64(nLeft+nMid) / 2
			}
			if num > nRight {
				return float64(nRight+nMid) / 2
			}
			return float64(nMid+num) / 2
		}
		var nLeft = nums2[midN]
		var nRight = nums2[midN+1]
		if num < nLeft {
			return float64(nLeft)
		}
		if num > nRight {
			return float64(nRight)
		}
		return float64(num)
	}

	if m == 2 {
		if n%2 == 1 {
			var nMid = nums2[midN]
			var nLeft = nums2[midN-1]
			var nRight = nums2[midN+1]
			if nums1[0] > nRight {
				return float64(nRight)
			}
			if nums1[1] < nLeft {
				return float64(nLeft)
			}
			if nums1[0] > nMid {
				return float64(nums1[0])
			}
			if nums1[1] < nMid {
				return float64(nums1[1])
			}
			return float64(nMid)
		}
		if n == 2 {
			var nLeft = nums2[midN]
			var nRight = nums2[midN+1]
			if nums1[0] > nRight {
				return float64(nRight+nums1[0]) / 2
			}
			if nums1[1] < nLeft {
				return float64(nLeft+nums1[1]) / 2
			}
			var r = nums1[1]
			if r > nRight {
				r = nRight
			}
			var l = nums1[0]
			if l < nLeft {
				l = nLeft
			}
			return float64(l+r) / 2
		}
		var nLeftL = nums2[midN-1]
		var nLeft = nums2[midN]
		var nRight = nums2[midN+1]
		var nRightR = nums2[midN+2]

		if nums1[0] > nRightR {
			return float64(nRight+nRightR) / 2
		}
		if nums1[1] < nLeftL {
			return float64(nLeftL+nLeft) / 2
		}
		if nums1[0] < nLeftL {
			var o = nums1[1]
			if o > nRight {
				o = nRight
			} else if o < nLeftL {
				o = nLeftL
			}
			return float64(nLeft+o) / 2
		}
		if nums1[1] > nRightR {
			var o = nums1[0]
			if o > nRightR {
				o = nRightR
			} else if o < nLeft {
				o = nLeft
			}
			return float64(nRight+o) / 2
		}
		var oLeft = nums1[0]
		var oRight = nums1[1]
		if oLeft < nLeft {
			oLeft = nLeft
		}
		if oRight > nRight {
			oRight = nRight
		}
		return float64(oLeft+oRight) / 2
	}

	if nums1[midM] < nums2[midN] {
		return findMedianSortedArrays(nums1[midM:], nums2[:n-midM])
	}
	return findMedianSortedArrays(nums2[midM:], nums1[:m-midM])
}
