package search

import (
	"math/rand"
	"testing"
)

var tests = []struct {
	input    int
	expected bool
}{
	{69, true},
	{1336, false},
	{69420, true},
	{69421, false},
	{1, true},
	{0, false},
}
var arr = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

func TestLinearSearch(t *testing.T) {
	for _, test := range tests {
		if LinearSearch(arr, test.input) != test.expected {
			t.Errorf("%d: expected %v, got %v", test.input, test.expected, !test.expected)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, test := range tests {
		if BinarySearch(arr, test.input) != test.expected {
			t.Errorf("%d: expected %v, got %v", test.input, test.expected, !test.expected)
		}
	}
}

func TestTwoChrystalBalls(t *testing.T) {
	idx := rand.Intn(10_000)
	arr := make([]bool, 10_000)

	for i := idx; i < 10_000; i += 1 {
		arr[i] = true
	}
	tmp := TwoCrystalBalls(arr)
	if tmp != idx {
		t.Errorf("Expected Return to be %d, got %d", idx, tmp)
	}
	arr2 := make([]bool, 500)
	tmp = TwoCrystalBalls(arr2)
	if tmp != -1 {
		t.Errorf("Expected no valid answer. Got %d", tmp)
	}
}
