package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fullArr := append(nums1, nums2...)
	sort.Ints(fullArr)
	var mid int
	if len(fullArr) > 2 {
		mid = len(fullArr) / 2
	} else {
		mid = len(fullArr)
	}
	if mid > 0 && mid%2 == 0 {
		return float64(fullArr[mid-1]+fullArr[mid]) / 2
	}
	return float64(fullArr[mid])
}

func main() {
	nums1 := []int{}
	nums2 := []int{2, 3}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
