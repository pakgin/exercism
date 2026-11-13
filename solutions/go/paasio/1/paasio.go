package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	io.Reader
	mu     sync.Mutex
	nbytes int
	nops   int
}

type writeCounter struct {
	io.Writer
	mu     sync.Mutex
	nbytes int
	nops   int
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{Writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{Reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		NewReadCounter(readwriter),
		NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	bytes, err := rc.Reader.Read(p)
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.nbytes += bytes
	rc.nops++
	return bytes, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return int64(rc.nbytes), rc.nops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	bytes, err := wc.Writer.Write(p)
	wc.mu.Lock()
	defer wc.mu.Unlock()
	wc.nbytes += bytes
	wc.nops++
	return bytes, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return int64(wc.nbytes), wc.nops
}
