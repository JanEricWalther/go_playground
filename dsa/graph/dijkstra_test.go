package graph

import "testing"

func TestDijkstraList(t *testing.T) {
	checkPath(t, DijkstraList(list1(), 0, 6), []int{
		0, 1, 4, 5, 6,
	})
}
