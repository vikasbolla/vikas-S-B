package main

import "fmt"

func fincTwoSum(arr []int, target int) []int {
	var result []int
	indexvalmap := make(map[int]int)
	// for i, v := range arr {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if v+arr[j] == target {
	// 			result = append(result, i, j)
	// 			break
	// 		}
	// 	}
	// }
	for i, v := range arr {
		if index, ok := indexvalmap[target-v]; ok {
			return []int{index, i}
		}
		//indexvalmap = map[int]int{i: target - v}v
		indexvalmap[v] = i
		fmt.Println(indexvalmap)

		// if indexvalmap[i]+indexvalmap[i+1] == target {
		// 	result = append(result, indexvalmap[i], indexvalmap[i+1])
		// }
	}

	return result
}
