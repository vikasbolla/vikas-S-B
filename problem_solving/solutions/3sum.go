package solutions

import (
	"reflect"
	"sort"
)

func ThreeSum(nums []int) [][]int {

	// i := 0
	// j := i + 1
		var triplets [][]int
		for i := 0; i < len(nums)/2; i++ {
			for j := i + 1; j < len(nums); j++ {
				for k := j + 1; k < len(nums); k++ {
					sum := nums[i] + nums[j] + nums[k]
					if sum == 0 {
						triplet := []int{nums[i], nums[j], nums[k]}
						sort.Ints(triplet)
						if !contains(triplet, triplets) {
							triplets = append(triplets, triplet)
						}
					}
				}
			}
		//	j := i + 1

		//findDistinctTriplets(triplets)
	}
	return triplets
}

// func findDistinctTriplets(triplets [][]int) bool {
// 	for _, triplet := range triplets {
// 		if contains(triplet) {

// 		}
// 	}
// 	return false
// }

func contains(triplet []int, triplets [][]int) bool {
	for _, triplet2 := range triplets {
		if reflect.DeepEqual(triplet, triplet2) {
			return true
		} else {
			continue
		}
	}
	return false
}

// func compareTwoTriplets(triplet, triplet2 []int) bool {
// 	count := 0
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			if triplet[i] == triplet2[j] {
// 				count++
// 			}
// 			if count == 3 {
// 				return true
// 			} else {
// 				continue
// 			}
// 		}
// 	}
// 	return count == 3
// }
