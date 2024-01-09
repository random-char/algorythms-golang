package sorting

import (
	"algorythms-golang/pkg/adt/heap"
)

func Heapsort(array []int) []int {
	h := heap.Heapify(array)
	return h.GetSortedMaxArray()
}
