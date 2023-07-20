package tree

type BinNode[T interface{}] struct {
	value T
	left  *BinNode[T]
	right *BinNode[T]
}

type TNode[T interface{}] struct {
	value    T
	children []BinNode[T]
}
