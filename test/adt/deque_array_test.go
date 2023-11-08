package adt

import (
	"algorythms-golang/pkg/adt/queue"
	"testing"
)

func TestPushAndPopWithDequeArray(t *testing.T) {
	list := queue.NewDequeArray[int]()
	PushAndPopWithAbstractList[int](t, &list)
}

func TestGetElementDequeArray(t *testing.T) {
	list := queue.NewDequeArray[int]()
	GetElementFromAbstractList[int](t, &list)
}

func TestPushOneSideAndPopFromOtherDequeArray(t *testing.T) {
	list := queue.NewDequeArray[int]()
	PushOneSideAndPopFromOther[int](t, &list)
}
