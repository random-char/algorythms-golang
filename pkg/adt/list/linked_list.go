package list

import adtError "algorythms-golang/pkg/adt/error"

type LinkedList[T int] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T int] struct {
	value T
	next  *Node[T]
}

func NewLinkedList[T int]() LinkedList[T] {
	return LinkedList[T]{}
}

func (ll *LinkedList[T]) Get(index int) (T, error) {
	count := 0
	current := ll.head

	for count < index && current != nil && current.next != nil {
		current = current.next
		count++
	}

	if count == index {
		return current.value, nil
	} else {
		return 0, &adtError.EmptyListError{}
	}
}

func (ll *LinkedList[T]) PushBack(value T) {
	newNode := Node[T]{
		value: value,
	}

	if ll.tail != nil {
		ll.tail.next = &newNode
		ll.tail = ll.tail.next
	} else {
		ll.head = &newNode
		ll.tail = &newNode
	}
}

func (ll *LinkedList[T]) PushFront(value T) {
	newNode := Node[T]{
		value: value,
		next:  ll.head,
	}

	ll.head = &newNode

	if ll.tail == nil {
		ll.tail = &newNode
	}
}

func (ll *LinkedList[T]) PopBack() (T, error) {
	var result T

	if ll.head == nil {
		return result, &adtError.EmptyListError{}
	}

	currentNode := ll.head
	if currentNode.next == nil {
		ll.head = nil
		ll.tail = nil
		return currentNode.value, nil
	}

	nextNode := currentNode.next
	for nextNode.next != nil {
		currentNode = nextNode
		nextNode = nextNode.next
	}

	result = nextNode.value
	currentNode.next = nil

	return result, nil
}

func (ll *LinkedList[T]) PopFront() (T, error) {
	var result T

	if ll.head == nil {
		return result, &adtError.EmptyListError{}
	}

	result = ll.head.value
	ll.head = ll.head.next

	if ll.head == nil {
		ll.tail = nil
	}

	return result, nil
}
