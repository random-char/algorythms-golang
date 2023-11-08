package sorting_test

import (
	"testing"
)

func AbstractSortingTester(t *testing.T, sortingFunc func([]int)) {
	arrays := [][]int{
		{2, 4, 1, 6, 8, 9, 0, 3, 5, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		{58, 70, 91, 68, 69, 31, 36, 72, 19, 51, 76, 96, 75, 15, 35, 85, 84, 40, 83, 34, 64, 61, 45, 27, 88, 23, 60, 21, 99, 43, 94, 53, 46, 65, 32, 24, 39, 5, 71, 14, 81, 33, 44, 90, 4, 1, 13, 98, 50, 30, 79, 9, 97, 2, 29, 56, 89, 78, 52, 57, 54, 12, 62, 10, 93, 7, 77, 80, 95, 49, 67, 6, 16, 87, 55, 82, 3, 38, 74, 17, 42, 8, 41, 48, 25, 20, 22, 28, 37, 66, 59, 63, 73, 26, 18, 11, 47, 86, 92},
	}

	for _, array := range arrays {
		sortingFunc(array)

		for i := 1; i < len(array); i++ {
			if array[i-1] > array[i] {
				t.Errorf("wrong order: %d before %d", array[i-1], array[i])
			}
		}
	}
}
