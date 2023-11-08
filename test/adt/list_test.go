package adt

import (
	"algorythms-golang/pkg/adt/list"
	"testing"
)

func TestPushAndPopWithList(t *testing.T) {
	list := list.NewList[int]()
	PushAndPopWithAbstractList[int](t, &list)
}

func TestGetElementList(t *testing.T) {
	list := list.NewList[int]()
	GetElementFromAbstractList[int](t, &list)
}

func TestPushOneSideAndPopFromOtherList(t *testing.T) {
	list := list.NewList[int]()
	PushOneSideAndPopFromOther[int](t, &list)
}
