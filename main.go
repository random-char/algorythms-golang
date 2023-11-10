package main

import (
	"algorythms-golang/pkg/adt/hash"
	"algorythms-golang/pkg/adt/list"
	"fmt"
)

func main() {
	ht := hash.NewHashTable()

	data := map[int]int{}
	for i := 0; i < 1000; i++ {
		data[i] = i
	}

	for k, v := range data {
		ht.Put(k, v)
		fmt.Println(ht.ToString())
	}

	fmt.Println(ht.Get(100))
	return

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
