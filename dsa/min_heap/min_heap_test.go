package minheap

import "testing"

func TestMinHeap(t *testing.T) {
	heap := New()
	var tmp int
	var ok bool

	expectEqual(t, heap.Len(), 0)

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(9)
	heap.Insert(7)

	expectEqual(t, heap.Len(), 8)

	tmp, ok = heap.Delete()
	expectEqual(t, ok, true)
	expectEqual(t, tmp, 1)
	tmp, ok = heap.Delete()
	expectEqual(t, tmp, 3)
	tmp, ok = heap.Delete()
	expectEqual(t, tmp, 4)
	tmp, ok = heap.Delete()
	expectEqual(t, tmp, 5)

	expectEqual(t, heap.Len(), 4)

	tmp, ok = heap.Delete()
	expectEqual(t, tmp, 7)
	tmp, ok = heap.Delete()
	expectEqual(t, tmp, 9)
	tmp, ok = heap.Delete()
	expectEqual(t, tmp, 69)
	tmp, ok = heap.Delete()
	expectEqual(t, tmp, 420)
	expectEqual(t, heap.Len(), 0)
	tmp, ok = heap.Delete()
	expectEqual(t, ok, false)
}

func expectEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	if actual != expected {
		t.Errorf("Wrong Value. expected %v, got %v", expected, actual)

	}
}
