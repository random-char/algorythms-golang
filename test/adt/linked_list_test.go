package adt

import (
	"algorythms-golang/pkg/adt/list"
	"testing"
)

func TestPushAndPopWithLinkedList(t *testing.T) {
	list := list.NewLinkedList[int]()
	PushAndPopWithAbstractList[int](t, &list)
}

func TestGetElementLinkedList(t *testing.T) {
	list := list.NewLinkedList[int]()
	GetElementFromAbstractList[int](t, &list)
}

func TestPushOneSideAndPopFromOtherLinkedList(t *testing.T) {
	list := list.NewLinkedList[int]()
	PushOneSideAndPopFromOther[int](t, &list)
}
