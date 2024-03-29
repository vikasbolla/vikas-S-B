package solutions

func ProductExceptSelf(nums []int) []int {

	result := make([]int, len(nums))

	for i := range result {
		result[i] = 1
	}
	leftProduct := 1
	for i := 0; i < len(nums); i++ {
		result[i] = result[i] * leftProduct
		leftProduct = leftProduct * nums[i]
	}
	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return result
}
