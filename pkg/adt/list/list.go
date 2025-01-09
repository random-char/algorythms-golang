package list

import adtError "algorythms-golang/pkg/adt/error"

type ArrayList[T int] struct {
	array  []T
	length int
}

func NewList[T int]() ArrayList[T] {
	return ArrayList[T]{
		array:  make([]T, 1),
		length: 0,
	}
}

func (l *ArrayList[T]) Get(i int) (T, error) {
	if i >= l.length {
		return 0, &adtError.OutOfBoundsError{}
	}

	return l.array[i], nil
}

func (l *ArrayList[T]) PushBack(value T) {
	l.increaseIfNeeded()

	l.array[l.length] = value
	l.length++
}

func (l *ArrayList[T]) PushFront(value T) {
	l.increaseIfNeeded()

	l.shiftRightFrom(0)
	l.array[0] = value
	l.length++
}

func (l *ArrayList[T]) PopBack() (T, error) {
	var result T

	if l.length == 0 {
		return result, &adtError.EmptyListError{}
	}

	result = l.array[l.length-1]
	l.array[l.length-1] = 0
	l.length--
	l.decreaseIfNeeded()

	return result, nil
}

func (l *ArrayList[T]) PopFront() (T, error) {
	var result T

	if l.length == 0 {
		return result, &adtError.EmptyListError{}
	}

	result = l.array[0]
	l.shiftLeftFrom(0)
	l.length--
	l.decreaseIfNeeded()

	return result, nil
}

func (l *ArrayList[T]) shiftLeftFrom(position int) {
	for i := position; i < l.length-1; i++ {
		l.array[i] = l.array[i+1]
	}

	l.array[l.length-1] = 0
}

func (l *ArrayList[T]) shiftRightFrom(position int) {
	for i := l.length; i > position; i-- {
		l.array[i] = l.array[i-1]
	}

	l.array[position] = 0
}

func (l *ArrayList[T]) increaseIfNeeded() {
	if len(l.array) == l.length {
		l.resize(l.length * 2)
	}
}

func (l *ArrayList[T]) decreaseIfNeeded() {
	if len(l.array)/4 >= l.length {
		l.resize(len(l.array) / 2)
	}
}

func (l *ArrayList[T]) resize(size int) {
	newArray := make([]T, size)
	for i := 0; i < l.length; i++ {
		newArray[i] = l.array[i]
	}

	l.array = newArray
}
