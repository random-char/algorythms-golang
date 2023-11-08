package adt

import (
	"algorythms-golang/pkg/adt/queue"
	"testing"
)

func TestPushAndPopWithDequeLinkedList(t *testing.T) {
	list := queue.NewDequeLinkedList[int]()
	PushAndPopWithAbstractList[int](t, &list)
}

func TestGetElementDequeLinkedList(t *testing.T) {
	list := queue.NewDequeLinkedList[int]()
	GetElementFromAbstractList[int](t, &list)
}

func TestPushOneSideAndPopFromOtherDequeLinkedList(t *testing.T) {
	list := queue.NewDequeLinkedList[int]()
	PushOneSideAndPopFromOther[int](t, &list)
}
