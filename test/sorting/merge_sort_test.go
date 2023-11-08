package sorting_test

import (
	"algorythms-golang/pkg/sorting"
	"testing"
)

func TestMergeSort(t *testing.T) {
	AbstractSortingTester(t, sorting.MergeSort)
}
