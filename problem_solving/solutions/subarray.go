package solutions

func FindLengthOfTargetSumSubArray(arr []int, target int) { //arr [5,8,-4,-1,3,2]  target=10
	i, j := 0, len(arr)-1
	sum := 0
	for i < len(arr) {
		sum = sum + arr[i]
	}
}
