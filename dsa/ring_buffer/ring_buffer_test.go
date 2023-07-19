package ringbuffer

import "testing"

func TestRingBuffer(t *testing.T) {
	buffer := New()
	var val int
	var ok bool

	buffer.Push(5)
	val, ok = buffer.Pop()
	expectEqual(t, val, 5)
	val, ok = buffer.Pop()
	expectEqual(t, ok, false)

	buffer.Push(42)
	buffer.Push(69)
	val, ok = buffer.Pop()
	expectEqual(t, val, 42)
	val, ok = buffer.Pop()
	expectEqual(t, val, 69)
	val, ok = buffer.Pop()
	expectEqual(t, ok, false)

	buffer.Push(42)
	buffer.Push(69)
	buffer.Push(12)
	val, ok = buffer.Get(2)
	expectEqual(t, val, 12)
	val, ok = buffer.Get(1)
	expectEqual(t, val, 69)
	val, ok = buffer.Get(0)
	expectEqual(t, val, 42)

}

func expectEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()

	if actual != expected {
		t.Errorf("Wrong Value. expected %v, got %v", expected, actual)

	}
}
