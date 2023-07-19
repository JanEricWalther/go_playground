package lru

import "testing"

func TestLRU(t *testing.T) {
	lru := New(3)
	var val int
	var ok bool

	val, ok = lru.Get("foo")
	expectEqual(t, ok, false)
	lru.Update("foo", 69)
	val, ok = lru.Get("foo")
	expectEqual(t, val, 69)

	lru.Update("bar", 420)
	val, ok = lru.Get("bar")
	expectEqual(t, val, 420)

	lru.Update("baz", 1337)
	val, ok = lru.Get("baz")
	expectEqual(t, val, 1337)

	lru.Update("nice meme", 69420)
	val, ok = lru.Get("nice meme")
	expectEqual(t, val, 69420)

	val, ok = lru.Get("foo")
	expectEqual(t, ok, false)
	val, ok = lru.Get("bar")
	expectEqual(t, val, 420)

	lru.Update("foo", 69)
	val, ok = lru.Get("foo")
	expectEqual(t, val, 69)
	val, ok = lru.Get("bar")
	expectEqual(t, val, 420)

	val, ok = lru.Get("baz")
	expectEqual(t, ok, false)
}

func expectEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()

	if actual != expected {
		t.Errorf("Wrong Value. expected %v, got %v", expected, actual)

	}
}
