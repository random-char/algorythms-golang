package sorting

func MergeSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}

	mid := length / 2

	//copy to new arrays not to update initial array
	l := make([]int, mid)
	r := make([]int, length-mid)
	copy(l, array[:mid])
	copy(r, array[mid:])

	MergeSort(l)
	MergeSort(r)

	lHead, rHead := 0, 0

	var inserting int
	for i := 0; i < length; i++ {
		if rHead >= len(r) || (lHead < len(l) && l[lHead] < r[rHead]) {
			inserting = l[lHead]
			lHead++
		} else {
			inserting = r[rHead]
			rHead++
		}

		array[i] = inserting
	}
}
