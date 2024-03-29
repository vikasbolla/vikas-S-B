package solutions

import "container/heap"

type Item struct {
	Value     int
	frequency int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency > pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item

}

func TopKFrequentElements(nums []int, k int) []int {
	freqMap := make(map[int]int, k)
	var result []int
	for _, v := range nums {
		freqMap[v]++
	}
	pq := make(PriorityQueue, len(freqMap))
	i := 0
	for num, freq := range freqMap {
		pq[i] = Item{Value: num, frequency: freq}
		i++
	}
	heap.Init(&pq)
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(Item)
		result = append(result, item.Value)
	}
	return result
}
