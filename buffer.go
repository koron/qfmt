package qfmt

import "sync"

type buffer []byte

func (b *buffer) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(p), nil
}

func (b *buffer) WriteString(s string) (n int, err error) {
	*b = append(*b, s...)
	return len(s), nil
}

func (b *buffer) reset() {
	*b = (*b)[:0]
}

var bufPool = sync.Pool{
	New: func() interface{} { return new(buffer) },
}

func buf_get() *buffer {
	return bufPool.Get().(*buffer)
}

func buf_put(b *buffer) {
	b.reset()
	bufPool.Put(b)
}
