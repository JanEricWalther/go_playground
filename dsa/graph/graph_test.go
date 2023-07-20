package graph

import "testing"

func TestBfsGraphList(t *testing.T) {
	checkPath(t, BfsGraphList(list1(), 0, 6), []int{
		0, 1, 4, 5, 6,
	})

	checkPath(t, BfsGraphList(list2(), 0, 6), []int{
		0, 1, 4, 5, 6,
	})

	tmp := BfsGraphList(list2(), 6, 0)
	if tmp != nil {
		t.Errorf("Expected nil, got %+v", tmp)
	}
}

func TestBfsGraphMatrix(t *testing.T) {
	checkPath(t, BfsGraphMatrix(matrix(), 0, 6), []int{
		0, 1, 4, 5, 6,
	})

	tmp := BfsGraphMatrix(matrix(), 6, 0)
	if tmp != nil {
		t.Errorf("Expected nil, got %+v", tmp)
	}
}

func TestDfsGraphList(t *testing.T) {
	checkPath(t, DfsGraphList(list2(), 0, 6), []int{
		0, 1, 4, 5, 6,
	})

	tmp := DfsGraphList(list2(), 6, 0)
	if tmp != nil {
		t.Errorf("Expected nil, got %+v", tmp)
	}
}

func checkPath(t *testing.T, actual, expected []int) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Errorf("Wrong length: got %d, expected %d", len(actual), len(expected))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Wrong value at index %d: got %d, expected %d", i, actual[i], expected[i])
		}
	}
}

func list1() GraphList {
	//      (1) --- (4) ---- (5)
	//    /  |       |       /|
	// (0)   | ------|------- |
	//    \  |/      |        |
	//      (2) --- (3) ---- (6)

	list := make(GraphList, 0)
	list = append(list, []GNode{
		{to: 1, weight: 3},
		{to: 2, weight: 1},
	})
	list = append(list, []GNode{
		{to: 0, weight: 3},
		{to: 2, weight: 4},
		{to: 4, weight: 1},
	})
	list = append(list, []GNode{
		{to: 1, weight: 4},
		{to: 3, weight: 7},
		{to: 0, weight: 1},
	})
	list = append(list, []GNode{
		{to: 2, weight: 7},
		{to: 4, weight: 5},
		{to: 6, weight: 1},
	})
	list = append(list, []GNode{
		{to: 1, weight: 1},
		{to: 3, weight: 5},
		{to: 5, weight: 2},
	})
	list = append(list, []GNode{
		{to: 6, weight: 1},
		{to: 4, weight: 2},
		{to: 2, weight: 18},
	})
	list = append(list, []GNode{
		{to: 3, weight: 1},
		{to: 5, weight: 1},
	})

	return list
}

func list2() GraphList {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	list := make(GraphList, 0)

	list = append(list, []GNode{
		{to: 1, weight: 3},
		{to: 2, weight: 1},
	})

	list = append(list, []GNode{
		{to: 4, weight: 1},
	})
	list = append(list, []GNode{
		{to: 3, weight: 7},
	})
	list = append(list, []GNode{})
	list = append(list, []GNode{
		{to: 1, weight: 1},
		{to: 3, weight: 5},
		{to: 5, weight: 2},
	})
	list = append(list, []GNode{
		{to: 2, weight: 18},
		{to: 6, weight: 1},
	})
	list = append(list, []GNode{
		{to: 3, weight: 1},
	})

	return list
}

func matrix() GraphMatrix {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	return [][]int{
		{0, 3, 1, 0, 0, 0, 0}, // 0
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}
}
