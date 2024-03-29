package solutions

func TwoSum(arr []int, target int) []int {
	indexMap := make(map[int]int)

	for i, val := range arr {
		complement := target - val
		if index, ok := indexMap[complement]; ok {
			return []int{index, i}
		}
		indexMap[val] = i
	}
	return []int{}
}
