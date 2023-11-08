package adt

import (
	"algorythms-golang/pkg/adt/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := heap.NewHeap()
	heap.RemoveMax()
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(4)
	heap.RemoveMax()
	heap.Insert(2)
}
