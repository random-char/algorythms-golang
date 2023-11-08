package main

import (
	"algorythms-golang/pkg/adt/hash"
	"algorythms-golang/pkg/adt/list"
	"fmt"
)

func main() {
	ht := hash.NewHashTable()

	ht.Put(1, 1)
	ht.ToString()
	ht.Put(100, 123)
	ht.ToString()
	ht.Put(10, 10)
	ht.ToString()
	ht.Put(11, 11)
	ht.ToString()

	fmt.Println(ht.Get(100))

	list := list.NewLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	fmt.Printf("%+v\n", list)
	value, err := list.PopBack()
	if err != nil {
		fmt.Println("got error popping value")
	}

	if value != 2 {
		fmt.Println("wrong value popped")
	}

	value, err = list.PopBack()
	if err != nil {
		fmt.Println("got error popping value")
	}

	if value != 1 {
		fmt.Println("wrong value popped")
	}

	_, err = list.PopBack()
	if err == nil {
		fmt.Println("expected error, got nil")
	}
}
