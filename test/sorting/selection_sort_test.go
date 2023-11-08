package sorting_test

import (
	"algorythms-golang/pkg/sorting"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	AbstractSortingTester(t, sorting.SelectionSort)
}
