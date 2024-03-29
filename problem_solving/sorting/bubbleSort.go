package sorting

//Bubble Sort is a simple sorting algorithm that repeatedly steps through the list,
//compares adjacent elements, and swaps them if they are in the wrong order.
//This process is repeated until the entire list is sorted. The algorithm gets its
//name because smaller elements “bubble” to the top of the list with each iteration.
func BubbleSort(arr []int) []int { //[10,30,40,20,50]
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
