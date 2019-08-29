package midian_two_sorted_arrays

import "sort"

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

	if m == 0 { // 边界条件
		if n%2 == 1 {
			return float64(nums2[midN])
		}
		return float64(nums2[midN]+nums2[midN+1]) / 2
	}

	if m == 1 || m == 2 { // 边界条件
		if n < 3 { // n小于3的情况下，取nums2所有元素和nums1的元素进行排序
			for i := 0; i < n; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else if n%2 == 1 { // n大于2且为奇数的情况下，取nums2中间3位和nums1的元素进行排序
			for i := midN - 1; i < midN+2; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else { // 其他情况下，取nums2的中间4位和nums1的元素进行排序
			for i := midN - 1; i < midN+3; i++ {
				nums1 = append(nums1, nums2[i])
			}
		}
		sort.Ints(nums1)
		m = len(nums1)
		midM = (m - 1) / 2

		// 递到头了，开始归
		if len(nums1)%2 == 1 {
			return float64(nums1[midM])
		} else {
			return float64(nums1[midM]+nums1[midM+1]) / 2
		}
	}

	if nums1[midM] < nums2[midN] {
		return findMedianSortedArrays(nums1[midM:], nums2[:n-midM])
	}
	return findMedianSortedArrays(nums2[midM:], nums1[:m-midM])
}
