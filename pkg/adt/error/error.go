package error

type EmptyListError struct{}

func (e *EmptyListError) Error() string {
	return "Empty list"
}

type OutOfBoundsError struct{}

func (e *OutOfBoundsError) Error() string {
	return "List out of bounds"
}

type EmptyHeapError struct{}

func (e *EmptyHeapError) Error() string {
	return "Heap is empty"
}

type HeapIsFullError struct{}

func (e *HeapIsFullError) Error() string {
	return "Heap is full"
}
