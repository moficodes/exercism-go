package paasio

import (
	"io"
	"sync"
)

// counter is subtype.

// It has a embedded type of sync.Mutex that protect the n and nops
type counter struct {
	sync.Mutex
	n    int64
	nops int
}

// readCounter implements ReadCounter

// Embedded struct type of counter
type readCounter struct {
	r io.Reader
	counter
}

// writeCounter implements WriteCounter

// Embedded struct type of counter
type writeCounter struct {
	w io.Writer
	counter
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func (c *counter) count() (int64, int) {
	c.Lock()
	defer c.Unlock()
	return c.n, c.nops
}

func (c *counter) addBytes(n int) {
	c.Lock()
	defer c.Unlock()
	c.n += int64(n)
	c.nops++
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.r.Read(p)
	rc.addBytes(n)
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.count()
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.addBytes(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	// count() is a method in counter struct
	// we can call like this because of type promotion
	return wc.count()
}

// NewReadCounter takes a reader and returns a ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r, counter{}}
}

// NewWriteCounter takes a writer and returns a WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w, counter{}}
}

//NewReadWriteCounter takes a readerwriter and returns a ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(rw), NewWriteCounter(rw)}
}
