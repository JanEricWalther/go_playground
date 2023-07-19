package minheap

type MinHeap struct {
	length int
	data   []Node
}

func New() *MinHeap {
	return &MinHeap{
		length: 0,
		data:   make([]Node, 0),
	}
}

func (h *MinHeap) Insert(val int) {
	node := Node{value: val}
	h.data = append(h.data, node)
	h.heapifyUp(h.length)
	h.length += 1
}

func (h *MinHeap) Delete() (val int, ok bool) {
	if h.length == 0 {
		return val, false
	}
	val = h.data[0].value
	ok = true

	h.data = h.data[:h.length]
	h.length -= 1
	h.data[0] = h.data[h.length]
	h.heapifyDown(0)
	return
}

func (h *MinHeap) Peek() (val int, ok bool) {
	if h.length == 0 {
		return val, false
	}
	return h.data[0].value, true

}

func (h *MinHeap) Len() int {
	return h.length
}

func (h *MinHeap) heapifyUp(idx int) {

	if idx == 0 {
		return
	}
	node := h.data[idx]
	parentIdx := h.parent(idx)
	parentNode := h.data[parentIdx]

	if node.value >= parentNode.value {
		return
	}
	h.data[idx] = parentNode
	h.data[parentIdx] = node
	h.heapifyUp(parentIdx)

}

func (h *MinHeap) heapifyDown(idx int) {
	if idx >= h.length {
		return
	}
	left := h.leftChild(idx)
	right := h.rightChild(idx)
	if left >= h.length {
		return
	}
	node := h.data[idx]
	leftNode := h.data[left]
	rightNode := h.data[right]

	if node.value <= leftNode.value && node.value <= rightNode.value {
		return
	}

	if leftNode.value < rightNode.value {
		h.data[idx] = leftNode
		h.data[left] = node
		h.heapifyDown(left)
		return
	}
	h.data[idx] = rightNode
	h.data[right] = node
	h.heapifyDown(right)
}

func (h *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}

func (h *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

type Node struct {
	value int
}
