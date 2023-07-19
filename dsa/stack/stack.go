package stack

type stackable interface {
	uint | int | string
}

type Stack[T stackable] struct {
	length   uint
	capacity uint
	data     []T
}

func New[T stackable](capacity uint) *Stack[T] {
	return &Stack[T]{
		length:   0,
		capacity: capacity,
		data:     make([]T, capacity, capacity),
	}
}

func (s *Stack[T]) Push(value T) {
	if s.length >= s.capacity {
		s.grow()
	}
	s.data[s.length] = value
	s.length += 1
}

func (s *Stack[T]) Peek() (value T, ok bool) {
	ok = false
	if s.length < 1 {
		return
	}
	return s.data[s.length-1], true
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	ok = false
	if s.length == 0 {
		return
	}
	s.length -= 1
	return s.data[s.length], true
}

func (s *Stack[T]) Len() uint {
	return s.length
}

func (s *Stack[T]) grow() {
	s.capacity *= 2
	tmp := make([]T, s.length, s.capacity)
	copy(tmp, s.data)
	s.data = tmp
}
