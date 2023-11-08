package sorting

func QuickSort(array []int) {
	quickSortRecursive(array, 0, len(array)-1)
}

func quickSortRecursive(array []int, firstIndex, lastIndex int) {
	if firstIndex >= lastIndex {
		return
	}

	pivot := array[firstIndex]
	pos := lastIndex

	for i := lastIndex; i > firstIndex; i-- {
		if array[i] > pivot {
			swap(array, pos, i)
			pos--
		}
	}

	swap(array, firstIndex, pos)
	quickSortRecursive(array, firstIndex, pos-1)
	quickSortRecursive(array, pos+1, lastIndex)
}

func swap(array []int, i, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}
