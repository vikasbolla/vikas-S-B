package sorting

//Insertion sort is another simple sorting algorithm that works by dividing the array into two parts:
// a sorted portion and an unsorted portion. The algorithm iterates through the unsorted portion and
//gradually inserts each element into its correct position in the sorted portion.
func InsertionSort(arr []int) []int { //[64, 34, 25, 12, 22, 11, 90]
	for i := 1; i < len(arr); i++ {
		key := arr[i] //34
		j := i - 1 //0
		for j >= 0 && arr[j] > key {  //64>34 --> true, false
			arr[j+1] = arr[j] //[64,64,25,12,22,11,90]
			j--  //j=-1
		}
		arr[j+1] = key  // [32,64,25,......]
	}
	return arr
}
