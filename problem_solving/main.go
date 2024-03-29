package main

import (
	"fmt"
	"pblmsolving/linkedList"
	"pblmsolving/searching"
	"pblmsolving/solutions"
	"pblmsolving/sorting"
	"pblmsolving/twopointers"
)

func main() {
	nums := []int{0, 0, 0, 0}
	arr := []int32{1, 8, 6, 2, 5, 4, 8, 3, 7}
	arr2 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	arr3 := [][]int32{{1, 2, 3, 0, 0},
		{0, 0, 0, 0, 0},
		{2, 1, 4, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0}}
	solutions.ReverseArray(arr)
	solutions.HourlyGlass(arr3)
	solutions.SecondMax(arr)
	solutions.FindMaxAreaOfContainer(arr2)
	solutions.ThreeSum(nums)
	solutions.IsAnagram("nagaram", "anagram")
	solutions.TwoSum([]int{1, 4, 4, 5}, 8)
	solutions.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	solutions.TopKFrequentElements([]int{1, 1, 1, 2, 2, 3}, 2)
	solutions.ProductExceptSelf([]int{1, 2, 3, 4})
	fmt.Println()
	fmt.Printf("element found at %d using linearSearch", searching.LinearSearch([]int{10, 20, 30, 40, 50, 60}, 40))
	fmt.Println()
	fmt.Printf("element found at %d using binarysearch", searching.BinarySearch([]int{10, 20, 30, 40, 50, 60}, 20))
	fmt.Println()
	fmt.Printf("element found at %d using binarysearchRescusion", searching.BinarySearchUsingRecursive([]int{10, 20, 30, 40, 50, 60}, 20, 0, 5))
	fmt.Println()
	fmt.Println("Array after sorting using bubblesort is", sorting.BubbleSort([]int{10, 40, 30, 20, 50}))
	fmt.Println()
	fmt.Println("Array after sorting using SelectionSort is", sorting.SelectionSort([]int{50, 30, 40, 20, 10}))
	fmt.Println()
	fmt.Println("Array after sorting using InsertionSort is", sorting.InsertionSort([]int{64, 34, 25, 12, 22, 11, 90}))
	fmt.Println()
	fmt.Println("Array after sorting using MergeSort is", sorting.MergeSort([]int{32, 11, 25, 40, 45, 20, 10, 5}))

	list := linkedList.LinkedList{}
	doublist := linkedList.DoublyLinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Display()
	doublist.Insert(10)
	doublist.Insert(20)
	doublist.Insert(30)
	doublist.DisplayFromHead()
	doublist.DisplayFromTail()
	linkedList.PrintConvertedList(linkedList.ArrayToLinkedList([]int{1, 2, 3, 4, 5, 6}))
	var head *linkedList.Node
	head = linkedList.AddNodeAtEnd(head, 200)
	head = linkedList.AddNodeAtEnd(head, 300)
	head = linkedList.AddNodeAtBegging(head, 100)
	head = linkedList.AddNodeAtEnd(head, 400)
	linkedList.PrintConvertedList(head)

	head = linkedList.DeleteNode(head, 200)
	linkedList.PrintConvertedList(head)
	head = linkedList.AddNodeAtEnd(head, 200)
	linkedList.PrintConvertedList(head)
	head = linkedList.DeleteNode(head, 300)
	linkedList.PrintConvertedList(head)
	head = linkedList.InsertNodeAfter(head, 100, 300)
	linkedList.PrintConvertedList(head)
	head = linkedList.InsertNodeBefore(head, 100, 50)
	linkedList.PrintConvertedList(head)
	head = linkedList.InsertNodeBefore(head, 400, 350)
	linkedList.PrintConvertedList(head)
	head = linkedList.InsertNodeBefore2(head, 200, 450)
	linkedList.PrintConvertedList(head)
	twopointers.ValidPalindrome("`l;`` 1o1 ??;l`")
	twopointers.TwoSum([]int{2, 7, 11, 15}, 9)
	twopointers.ThreeSum([]int{-2, 0, 1, 1, 2})
}
