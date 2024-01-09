package sorting_test

import (
	"algorythms-golang/pkg/sorting"
	"testing"
)

func TestHeapsort(t *testing.T) {
	AbstractSortingTester(t, sorting.Heapsort)
}
