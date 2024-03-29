package searching

func BinarySearch(nums []int, target int) int { // [10,20,30,40,50,60],20
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchUsingRecursive(nums []int, target, left, right int) int {

	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return BinarySearchUsingRecursive(nums, target, mid+1, right)
	} else {
		return BinarySearchUsingRecursive(nums, target, left, mid-1)
	}
}
