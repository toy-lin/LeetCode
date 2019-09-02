package midian_two_sorted_arrays

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	var arr1 = []int{1, 3}
	var arr2 = []int{2}

	fmt.Println(findMedianSortedArrays(arr1, arr2))

	arr1 = []int{1, 2}
	arr2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

	arr1 = []int{1, 2}
	arr2 = []int{-1, 3}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

	arr1 = []int{1, 2}
	arr2 = []int{1, 2, 3}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

	arr1 = []int{1, 2}
	arr2 = []int{3, 4, 5, 6}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

	arr1 = []int{2, 2, 2}
	arr2 = []int{2, 2, 2, 2}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

	arr1 = []int{1, 2, 3, 4}
	arr2 = []int{1, 2, 3, 4}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

	arr1 = []int{1, 2, 2, 2}
	arr2 = []int{1, 2, 3, 4}
	fmt.Println(findMedianSortedArrays(arr1, arr2))
}
