package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := New[int](10)
	var tmp int
	var ok bool

	stack.Push(5)
	stack.Push(7)
	stack.Push(9)

	tmp, ok = stack.Pop()
	expectEqual(t, tmp, 9)
	expectEqual(t, stack.Len(), uint(2))

	stack.Push(11)
	tmp, ok = stack.Pop()
	expectEqual(t, tmp, 11)
	tmp, ok = stack.Pop()
	expectEqual(t, tmp, 7)
	tmp, ok = stack.Peek()
	expectEqual(t, tmp, 5)
	tmp, ok = stack.Pop()
	expectEqual(t, tmp, 5)
	tmp, ok = stack.Pop()
	expectEqual(t, ok, false)

	stack.Push(69)
	tmp, ok = stack.Peek()
	expectEqual(t, tmp, 69)
	expectEqual(t, stack.Len(), uint(1))
}

func expectEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	if actual != expected {
		t.Errorf("Wrong Value. expected %v, got %v", expected, actual)

	}
}
