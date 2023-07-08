package main

import (
	"fmt"
)

func mergeAndReturn(nums1 []int, m int, nums2 []int, n int) []int {
	var buf = make([]int, m+n)
	var i = 0
	var j = 0
	var idx = 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			buf[idx] = nums1[i]
			i += 1
		} else if nums1[i] > nums2[j] {
			buf[idx] = nums2[j]
			j += 1
		} else {
			buf[idx] = nums1[i]
			idx += 1
			buf[idx] = nums2[j]
			i += 1
			j += 1
		}
		idx += 1
	}
	for i < m {
		buf[idx] = nums1[i]
		i += 1
		idx += 1
	}
	for j < n {
		buf[idx] = nums2[j]
		j += 1
		idx += 1
	}
	return buf
}

func main() {
	var n1 = []int{1, 2, 4, 5}
	var n2 = []int{3, 7}
	var buf = mergeAndReturn(n1, len(n1), n2, len(n2))
	fmt.Println("res %v", buf)
}
