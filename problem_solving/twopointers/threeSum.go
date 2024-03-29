package twopointers

import "sort"

func ThreeSum(nums []int) [][]int { //[-1,0,1,2,-1,-4]  op:= [[-1,-1,2],[-1,0,1]]
	out := make([][]int, 0, len(nums)/3)
	sort.Ints(nums)
	var left, right, sum int
	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}
		left, right = i+1, len(nums)-1
		for left < right {
			sum = a + nums[left] + nums[right]
			if sum == 0 {
				out = append(out, []int{a, nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
			if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return out
}

// func ThreeSum(nums []int) [][]int { //[-1,0,1,2,-1,-4]  op:= [[-1,-1,2],[-1,0,1]]
// 	//var j, k int
// 	result := make([][]int, 0, 3)
// 	//i := 0
// 	for i := 0; i < len(nums); i++ {
// 		j := i + 1
// 		k := j + 1
// 		for j < k && k < len(nums) {
// 			if nums[i]+nums[j]+nums[k] == 0 {
// 				triplet := []int{nums[i], nums[j], nums[k]}
// 				sort.Ints(triplet)
// 				if !contains(triplet, result) {
// 					result = append(result, triplet)
// 				}
// 				if k == len(nums)-1 {
// 					k = j + 2
// 					j++
// 					continue
// 				}
// 				k++
// 			} else {
// 				if k == len(nums)-1 {
// 					k = j + 2
// 					j++
// 					continue
// 				}
// 				k++
// 			}

// 		}

// 	}
// 	return result

// }

// func contains(triplet []int, triplets [][]int) bool {
// 	for _, triplet2 := range triplets {
// 		if reflect.DeepEqual(triplet, triplet2) {
// 			return true
// 		} else {
// 			continue
// 		}
// 	}
// 	return false
// }
