package circularbuffer

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	Capacity int
	Len      int
	Start    int
	buffer   []byte
}

func NewBuffer(size int) *Buffer {
	return &Buffer{Capacity: size, Len: 0, Start: 0, buffer: make([]byte, size)}
}

func (b *Buffer) ReadByte() (byte, error) {
	var zero byte
	if b.Len == 0 {
		return zero, errors.New("Empty buffer")
	}

	zero = b.buffer[b.Start]
	b.Start = (b.Start + 1) % b.Capacity
	b.Len--
	return zero, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.Len == b.Capacity {
		return errors.New("Full buffer")
	}

	b.buffer[(b.Start+b.Len)%b.Capacity] = c
	b.Len++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		b.ReadByte()
		b.WriteByte(c)
	}
}

func (b *Buffer) Reset() {
	b.Start = 0
	b.Len = 0
}
