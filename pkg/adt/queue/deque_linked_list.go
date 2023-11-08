package queue

import adtError "algorythms-golang/pkg/adt/error"

type DoubleEndedNode[T int] struct {
	value T
	next  *DoubleEndedNode[T]
	prev  *DoubleEndedNode[T]
}

type DequeLinkedList[T int] struct {
	head   *DoubleEndedNode[T]
	tail   *DoubleEndedNode[T]
	length int
}

func NewDequeLinkedList[T int]() DequeLinkedList[T] {
	return DequeLinkedList[T]{
		length: 0,
	}
}

func (d *DequeLinkedList[T]) Get(index int) (T, error) {
	var result T

	if index >= d.length {
		return result, &adtError.OutOfBoundsError{}
	}

	count := 0
	current := d.head

	for count < index && current != nil && current.next != nil {
		current = current.next
		count++
	}

	if count == index {
		return current.value, nil
	} else {
		return 0, &adtError.OutOfBoundsError{}
	}
}

func (d *DequeLinkedList[T]) PushFront(value T) {
	newNode := DoubleEndedNode[T]{
		value: value,
	}

	if d.length == 0 {
		d.head = &newNode
		d.tail = &newNode
	} else {
		newNode.next = d.head
		d.head.prev = &newNode
		d.head = &newNode
	}

	d.length++
}

func (d *DequeLinkedList[T]) PushBack(value T) {
	newNode := DoubleEndedNode[T]{
		value: value,
	}

	if d.length == 0 {
		d.head = &newNode
		d.tail = &newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = &newNode
		d.tail = &newNode
	}

	d.length++
}

func (d *DequeLinkedList[T]) PopFront() (T, error) {
	var result T
	var err error

	if d.length == 0 {
		err = &adtError.EmptyListError{}
	} else {
		result = d.head.value
		d.head = d.head.next
		d.length--
	}

	return result, err
}

func (d *DequeLinkedList[T]) PopBack() (T, error) {
	var result T
	var err error

	if d.length == 0 {
		err = &adtError.EmptyListError{}
	} else {
		result = d.tail.value
		d.tail = d.tail.prev
		d.length--
	}

	return result, err
}
