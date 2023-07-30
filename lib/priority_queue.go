package lib

import (
	"container/heap"
	"sort"
)

const (
	//超过Threshold的键值对才加到大根堆中
	Threshold = 100
	//大根堆长度
	SortQueueLimit = 100
)

type DataItem struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

// 定义有限长度的队列（最大容量为100）
type SortableQueue []DataItem

// Len returns the length of the queue.
func (q SortableQueue) Len() int {
	return len(q)
}

// Less returns true if the item at index i is greater than the item at index j (for max heap).
func (q SortableQueue) Less(i, j int) bool {
	return q[i].Value > q[j].Value
}

// Swap swaps the items at index i and j.
func (q SortableQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

type PriorityQueue struct {
	q SortableQueue
}

func NewPriorityQueue(sortableQueue SortableQueue) *PriorityQueue {
	q := PriorityQueue{sortableQueue}
	heap.Init(&q)
	return &q
}

// Len returns the length of the queue.
func (q *PriorityQueue) Len() int {
	return len(q.q)
}

// Less returns true if the item at index i is greater than the item at index j (for max heap).
func (q *PriorityQueue) Less(i, j int) bool {
	return q.q[i].Value > q.q[j].Value
}

// Swap swaps the items at index i and j.
func (q *PriorityQueue) Swap(i, j int) {
	q.q[i], q.q[j] = q.q[j], q.q[i]
}

func (q *PriorityQueue) Push(x interface{}) {
	if q.Len() >= SortQueueLimit {
		heap.Pop(q) // Remove the smallest element (root) in case of overflow
	}
	item := x.(DataItem)
	q.q = append(q.q, item)
	if x.(DataItem).Value > Threshold {
		heap.Fix(q, len(q.q)-1)
	}
}

func (q *PriorityQueue) Sort() {
	if len(q.q) <= SortQueueLimit {
		sort.Sort(q)
	} else {
		sort.Sort(q.q[:SortQueueLimit])
	}
}

func (q *PriorityQueue) GetRank(rank int) DataItem {
	return q.q[rank]
}

func (q *PriorityQueue) GetSortedQueue() []int {
	sortedQueue := make([]int, 0, SortQueueLimit)
	for i, data := range q.q {
		if i >= Threshold {
			break
		}
		sortedQueue = append(sortedQueue, data.Value)
	}
	return sortedQueue
}

// up moves the element at index i up to its correct position to maintain the heap property.
//func (q *PriorityQueue) up(i int) {
//	for i > 0 {
//		parent := (i - 1) / 2
//		if q.Less(i, parent) {
//			break
//		}
//		q.Swap(i, parent)
//		i = parent
//	}
//}

// Pop removes and returns the last item from the SortableQueue (not used in this case).
func (q *PriorityQueue) Pop() interface{} {
	old := q.q
	x := old[0]
	q.q = old[1:]
	return x
}
