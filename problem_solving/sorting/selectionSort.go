package sorting

//Selection Sort is a simple sorting algorithm that repeatedly finds the minimum (or maximum)
// element from the unsorted portion of the list and swaps it with the leftmost unsorted element.
// This process is repeated until the entire list is sorted.
func SelectionSort(arr []int) []int { //[50,30,20,40,10], [20,30,50,40,10]
	for i, _ := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}

		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
	return arr
}
