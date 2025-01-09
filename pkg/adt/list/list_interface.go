package list

type List[T int] interface {
	Get(int) (T, error)
	PushBack(T)
	PushFront(T)
	PopBack() (T, error)
	PopFront() (T, error)
}
