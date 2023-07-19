package sorting

import "testing"

var test = struct {
	input    []int
	expected []int
}{
	input:    []int{9, 3, 7, 4, 69, 420, 42},
	expected: []int{3, 4, 7, 9, 42, 420},
}

func TestInsertionSort(t *testing.T) {
	testArray(t, InsertionSort(test.input), test.expected)
}

func TestBubbleSort(t *testing.T) {
	testArray(t, BubbleSort(test.input), test.expected)
}

func TestQuickSort(t *testing.T) {
	testArray(t, QuickSort(test.input), test.expected)
}
func TestMergeSort(t *testing.T) {
	testArray(t, MergeSort(test.input), test.expected)
}

func testArray(t *testing.T, actual, expected []int) {
	if len(expected) != len(actual) {
		t.Errorf("Wrong length. got %d, expected %d", len(actual), len(expected))
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Wrong value at index %d. got %d, expected %d", i, actual[i], expected[i])
		}
	}
}
