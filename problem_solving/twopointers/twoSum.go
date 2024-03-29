package twopointers

func TwoSum(numbers []int, target int) []int { //[2,7,11,15]  9
	result := make([]int, 2)
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return result
}
