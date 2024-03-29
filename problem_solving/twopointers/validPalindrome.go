package twopointers

import (
	"fmt"
	"regexp"
	"strings"
)

//remove non-alpha char
//convert to lower case
//check if palin

func ValidPalindrome(str string) bool { //"A man, a plan, a canal: Panama"
	reg, err := regexp.Compile("[^a-zA-z0-9]+")
	nonAlphanumericRegex, err2 := regexp.Compile(`[^\p{L}\p{N} ]+`)
	if err != nil || err2 != nil {
		panic(err)
	}
	str = reg.ReplaceAllString(str, "")
	str = nonAlphanumericRegex.ReplaceAllString(str, "")
	str = strings.ToLower(str)
	i, j := 0, len(str)-1
	if str == "" {
		return true
	}
	for i < j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	fmt.Println(str)
	return true
}
