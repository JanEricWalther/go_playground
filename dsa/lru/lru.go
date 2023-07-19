package lru

type LRU struct {
}

func New(capacity int) *LRU {
	return &LRU{}
}

func (lru *LRU) Get(key string) (val int, ok bool) {

}

func (lru *LRU) Update(key string, value int) {

}
