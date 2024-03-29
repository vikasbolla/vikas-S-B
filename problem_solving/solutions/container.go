package solutions

func FindMaxAreaOfContainer(arr []int) int {
	var area int
	for left := 0; left < len(arr); left++ {
		for right := left + 1; right < len(arr); right++ {
			curreArea := min(arr[left], arr[right]) * (right - left)
			area = max(area, curreArea)
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
