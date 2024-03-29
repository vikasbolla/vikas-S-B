package solutions

func HourlyGlass(arr [][]int32) int32 {
	var maxSum int32 = 0
	var sum int32
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			sum = (arr[row][col] + arr[row][col+1] + arr[row][col+2] + arr[row+1][col+1] + arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2])
		}
		if sum > maxSum {
			maxSum = sum
		} else {
			continue
		}
	}
	return maxSum
}
