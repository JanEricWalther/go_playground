package tree

import "testing"

func TestBFS(t *testing.T) {
	found := Bfs(tree1(), 45)
	if found != true {
		t.Errorf("expected %v to be `true`", found)
	}

	found = Bfs(tree1(), 7)
	if found != true {
		t.Errorf("expected %v to be `true`", found)
	}

	found = Bfs(tree1(), 69)
	if found != false {
		t.Errorf("expected %v to be `false`", found)
	}
}

func TestDFS(t *testing.T) {
	found := Dfs(tree1(), 45)
	if found != true {
		t.Errorf("expected %v to be `true`", found)
	}

	found = Dfs(tree1(), 7)
	if found != true {
		t.Errorf("expected %v to be `true`", found)
	}

	found = Dfs(tree1(), 69)
	if found != false {
		t.Errorf("expected %v to be `false`", found)
	}
}

func TestCompare(t *testing.T) {
	same := Compare(tree1(), tree1())
	if same != true {
		t.Errorf("expected %v to be `true`", same)
	}

	same = Compare(tree1(), tree2())
	if same != false {
		t.Errorf("expected %v to be `false`", same)
	}
}

func TestTraverseInOrder(t *testing.T) {
	checkPath(t, TraverseInOrder[int](tree1()), []int{
		5, 7, 10, 15, 20, 29, 30, 45, 50, 100,
	})
}

func TestTraversePostOrder(t *testing.T) {
	checkPath(t, TraversePostOrder[int](tree1()), []int{
		7, 5, 15, 10, 29, 45, 30, 100, 50, 20,
	})
}

func TestTraversePreOrder(t *testing.T) {
	checkPath(t, TraverseInOrder[int](tree1()), []int{
		20, 10, 5, 7, 15, 50, 30, 29, 45, 100,
	})
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

func tree1() BinNode[int] {
	var tree BinNode[int]
	tree = BinNode[int]{
		value: 20,
		right: &BinNode[int]{
			value: 50,
			right: &BinNode[int]{
				value: 100,
				right: nil,
				left:  nil,
			},
			left: &BinNode[int]{
				value: 30,
				right: &BinNode[int]{
					value: 45,
					right: nil,
					left:  nil,
				},
				left: &BinNode[int]{
					value: 29,
					right: nil,
					left:  nil,
				},
			},
		},
		left: &BinNode[int]{
			value: 10,
			right: &BinNode[int]{
				value: 15,
				right: nil,
				left:  nil,
			},
			left: &BinNode[int]{
				value: 5,
				right: &BinNode[int]{
					value: 7,
					right: nil,
					left:  nil,
				},
				left: nil,
			},
		},
	}

	return tree
}

func tree2() BinNode[int] {
	var tree BinNode[int]
	tree = BinNode[int]{
		value: 20,
		right: &BinNode[int]{
			value: 50,
			right: nil,
			left: &BinNode[int]{
				value: 30,
				right: &BinNode[int]{
					value: 45,
					right: &BinNode[int]{
						value: 49,
						left:  nil,
						right: nil,
					},
					left: nil,
				},
				left: &BinNode[int]{
					value: 29,
					right: nil,
					left: &BinNode[int]{
						value: 21,
						right: nil,
						left:  nil,
					},
				},
			},
		},
		left: &BinNode[int]{
			value: 10,
			right: &BinNode[int]{
				value: 15,
				right: nil,
				left:  nil,
			},
			left: &BinNode[int]{
				value: 5,
				right: &BinNode[int]{
					value: 7,
					right: nil,
					left:  nil,
				},
				left: nil,
			},
		},
	}
	return tree
}
