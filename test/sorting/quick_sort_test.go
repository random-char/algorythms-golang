package sorting_test

import (
	"algorythms-golang/pkg/sorting"
	"testing"
)

func TestQuickSort(t *testing.T) {
	AbstractSortingTester(t, sorting.QuickSort)
}
