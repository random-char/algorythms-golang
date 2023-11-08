package sorting

func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		j := findIndexOfSmallestEl(array, i)

		tmp := array[i]
		array[i] = array[j]
		array[j] = tmp
	}
}

func findIndexOfSmallestEl(array []int, from int) int {
	index := from

	for i := from; i < len(array); i++ {
		if array[i] < array[index] {
			index = i
		}
	}

	return index
}
