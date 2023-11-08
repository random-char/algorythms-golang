package heap

import (
	adtError "algorythms-golang/pkg/adt/error"
)

type MaxHeap struct {
	array  []int
	length int
}

func NewHeap() MaxHeap {
	return MaxHeap{
		array:  make([]int, 10),
		length: 0,
	}
}

func (h *MaxHeap) Insert(value int) error {
	if h.length == len(h.array) {
		return &adtError.HeapIsFullError{}
	}

	h.array[h.length] = value
	h.fixUp(h.length)
	h.length++

	return nil
}

func (h *MaxHeap) RemoveMax() (int, error) {
	if h.length == 0 {
		return 0, &adtError.EmptyHeapError{}
	}

	h.length--
	swap(h.array, h.length, 0)
	h.fixDown(0)

	return h.array[h.length], nil
}

func (h *MaxHeap) fixUp(index int) {
	for index > 0 {
		parentIndex := getParentIndex(index)

		if h.array[parentIndex] > h.array[index] {
			break
		}

		swap(h.array, parentIndex, index)

		index = parentIndex
	}
}

func (h *MaxHeap) fixDown(index int) {
	for index < h.length {
		childIndex := getLeftChildIndex(index)
		if childIndex >= h.length {
			return
		}

		if childIndex+1 != h.length && h.array[childIndex] < h.array[childIndex+1] {
			childIndex++
		}

		if h.array[index] < h.array[childIndex] {
			swap(h.array, index, childIndex)
		}

		index = childIndex
	}
}

func swap(array []int, i, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}

func getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func getLeftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}
