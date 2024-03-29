package solutions

func ReverseArray(a []int32) []int32 {
	// Write your code here
	var revArr []int32
	for i := len(a) - 1; i >= 0; i-- {
		revArr = append(revArr, a[i])
	}
	return revArr
}
