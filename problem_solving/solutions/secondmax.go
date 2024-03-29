package solutions

func SecondMax(height []int32) int32 {
	max := height[0]
	secondMax := height[1]
	for _, val := range height {
		if val > max {
			secondMax = max
			max = val
		}
		if val > secondMax && val != max {
			secondMax = val
		}
	}

	return secondMax
}
