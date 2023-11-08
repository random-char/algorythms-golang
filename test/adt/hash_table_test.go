package adt

import (
	"algorythms-golang/pkg/adt/hash"
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := hash.NewHashTable()

	ht.Put(1, 1)
	ht.ToString()
	ht.Put(100, 100)
	ht.ToString()
	ht.Put(10, 10)
	ht.ToString()
	ht.Put(11, 11)
	ht.ToString()
}
