package main

import (
	"fmt"
	"strings"
)

func subStrings(str string) int {
	//var allSubstrings []string
	var substr string
	maxLength := 0
	for i, _ := range str {
		substr = substr + string(str[i])
		for j := i + 1; j < len(str); j++ {
			if !strings.Contains(substr, string(str[j])) {
				substr = substr + string(str[j])
			} else {
				break
			}
		}
		if len(substr) > maxLength {
			maxLength = len(substr)
		}
		//allSubstrings = append(allSubstrings, substr)
		substr = ""
	}
	//fmt.Println(allSubstrings)
	//maxLength := findMaxLengthOfSubstrings(allSubstrings)
	return maxLength
}

// func findMaxLengthOfSubstrings(substrings []string) int {
// 	maxLength := 0
// 	for _, substring := range substrings {
// 		if len(substring) > maxLength {
// 			maxLength = len(substring)
// 		}
// 	}
// 	return maxLength
// }

func main() {
	s := "bbbbb"
	maxLen := subStrings(s)
	fmt.Println(maxLen)
	arr := []int{7, 3, 5}
	target := 8
	result := fincTwoSum(arr, target)
	fmt.Println(result)
}
