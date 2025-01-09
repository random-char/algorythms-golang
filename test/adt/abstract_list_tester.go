package adt

import (
	"algorythms-golang/pkg/adt/list"
	"testing"
)

func PushAndPopWithAbstractList[T int](t *testing.T, list list.List[T]) {
	var val T
	var err error

	list.PushBack(1)
	list.PushBack(2)

	val, err = list.PopBack()
	if err != nil {
		t.Fatal("got error during popping val")
	}

	if val != 2 {
		t.Fatal("wrong val popped")
	}

	val, err = list.PopBack()
	if err != nil {
		t.Fatal("got error during popping val")
	}

	if val != 1 {
		t.Fatal("wrong val popped")
	}

	_, err = list.PopBack()
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	list.PushFront(2)
	list.PushFront(1)

	val, err = list.PopFront()
	if err != nil {
		t.Fatal("got error during popping val")
	}

	if val != 1 {
		t.Fatal("wrong val popped")
	}

	val, err = list.PopFront()
	if err != nil {
		t.Fatal("got error during popping val")
	}

	if val != 2 {
		t.Fatal("wrong val popped")
	}

	_, err = list.PopBack()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func GetElementFromAbstractList[T int](t *testing.T, list list.List[T]) {
	var val T
	var err error

	list.PushBack(2)
	list.PushBack(3)
	list.PushFront(1)

	val, err = list.Get(0)
	if err != nil {
		t.Fatal("got error during getting val")
	}
	if val != 1 {
		t.Fatalf("got %d, expected 1", val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Fatal("got error during getting val")
	}
	if val != 2 {
		t.Fatalf("got %d, expected 1", val)
	}
}

func PushOneSideAndPopFromOther[T int](t *testing.T, list list.List[T]) {
	for i := 1; i <= 20; i++ {
		list.PushBack(T(i))
	}

	for i := 1; i <= 20; i++ {
		val, err := list.PopFront()
		if val != T(i) {
			t.Fatalf("got %d, expected %d", val, i)
		}
		if err != nil {
			t.Fatal("got error during getting val")
		}
	}
}
