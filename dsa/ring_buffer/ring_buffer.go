package ringbuffer

type RingBuffer struct {
}

func New() *RingBuffer {
	return &RingBuffer{}
}

func (b *RingBuffer) Push(value int) {

}

func (b *RingBuffer) Pop() (val int, ok bool) {

}

func (b *RingBuffer) Get(idx int) (val int, ok bool) {

}
