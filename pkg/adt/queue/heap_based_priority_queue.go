package queue

import "algorythms-golang/pkg/adt/heap"

type PriorityQueue struct {
	heap heap.MaxHeap
}

func (pq *PriorityQueue) Enque(value int) {
	pq.heap.Insert(value)
}

func (pq *PriorityQueue) DequeMax() (int, error) {
	return pq.heap.RemoveMax()
}
