package queue

import adtError "algorythms-golang/pkg/adt/error"

type DequeArray[T int] struct {
	array  []T
	start  int //index of first element
	end    int //index of next to last element
	length int
}

func NewDequeArray[T int]() DequeArray[T] {
	return DequeArray[T]{
		array:  make([]T, 10),
		start:  0,
		end:    0,
		length: 0,
	}
}

func (d *DequeArray[T]) Empty() bool {
	return d.length == 0
}

func (d *DequeArray[T]) Get(i int) (T, error) {
	var result T

	if i >= d.length {
		return result, &adtError.OutOfBoundsError{}
	}

	result = d.array[(d.start+i)%len(d.array)]
	return result, nil
}

func (d *DequeArray[T]) PushBack(value T) {
	d.increaseIfNeeded()

	d.array[d.end] = value
	d.end = (d.end + 1) % (len(d.array) + 1)
	d.length++
}

func (d *DequeArray[T]) PushFront(value T) {
	d.increaseIfNeeded()

	d.start = (d.start - 1 + len(d.array)) % len(d.array)
	d.array[d.start] = value
	d.length++
}

func (d *DequeArray[T]) PopBack() (T, error) {
	var result T

	if d.Empty() {
		return result, &adtError.EmptyListError{}
	}

	d.end = (d.end - 1 + len(d.array)) % len(d.array)
	result = d.array[d.end]
	d.length--

	d.decreaseIfNeeded()

	return result, nil
}

func (d *DequeArray[T]) PopFront() (T, error) {
	var result T

	if d.Empty() {
		return result, &adtError.EmptyListError{}
	}

	result = d.array[d.start]
	d.start = (d.start + 1) % len(d.array)
	d.length--

	d.decreaseIfNeeded()

	return result, nil
}

func (d *DequeArray[T]) increaseIfNeeded() {
	if len(d.array) == d.length {
		d.resize(len(d.array) * 2)
	}
}

func (d *DequeArray[T]) decreaseIfNeeded() {
	if len(d.array)/4 > d.length {
		d.resize(len(d.array) / 2)
	}
}

func (d *DequeArray[T]) resize(size int) {
	newArray := make([]T, size)

	for i := 0; i < d.length; i++ {
		newArray[i] = d.array[(d.start+i)%len(d.array)]
	}

	d.array = newArray
	d.start = 0
	d.end = d.length
}
