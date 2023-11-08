package hash

import (
	"errors"
	"fmt"
)

type HashTable struct {
	array       []HashTableList
	arrayLength int
	count       int
	hashFunc    func(int) int
}

func NewHashTable() *HashTable {
	length := 10

	array := make([]HashTableList, length)
	for i := 0; i < length; i++ {
		array[i] = HashTableList{}
	}

	ht := HashTable{
		arrayLength: length,
		array:       array,
	}

	ht.generateAndSetNewHashFunction()

	return &ht
}

func (ht *HashTable) ToString() string {
	s := "-------------------------------"
	for i, hashTableList := range ht.array {
		s += fmt.Sprintf("%d)->", i)

		node := hashTableList.head
		for {
			if node == nil {
				break
			}

			s += fmt.Sprintf("|%d|->", node.value)
			node = node.next
		}

		s += "\n"
	}

	return s + "-------------------------------\n\n"
}

type HashTableList struct {
	head *HashTableNode
}

type HashTableNode struct {
	key   int
	value int
	next  *HashTableNode
}

func (ll *HashTableList) PushBack(key, value int) {
	newNode := HashTableNode{
		key:   key,
		value: value,
	}

	if ll.head == nil {
		ll.head = &newNode
	} else {
		node := ll.head
		for {
			if node.next == nil {
				node.next = &newNode
				return
			} else {
				node = node.next
			}
		}
	}
}

func (ll *HashTableList) Get(key int) (int, error) {
	current := ll.head

	for current != nil && current.key != key {
		current = current.next
	}

	if current != nil {
		return current.value, nil
	} else {
		return 0, errors.New("not found")
	}
}

func (ht *HashTable) Put(key, value int) {
	keyHash := ht.hashFunc(key)
	ht.array[keyHash].PushBack(key, value)
	ht.count++
}

func (ht *HashTable) Get(key int) (int, error) {
	keyHash := ht.hashFunc(key)
	return ht.array[keyHash].Get(key)
}

func (ht *HashTable) rehashIfNeeded() {
	//change size
	//change hashing function
	//recalculate hashes and store in new array
}

func (ht *HashTable) generateAndSetNewHashFunction() {
	f := func(i int) int {
		return i % len(ht.array)
	}

	ht.hashFunc = f
}
