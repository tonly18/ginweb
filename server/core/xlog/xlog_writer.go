package xlog

import (
	"bufio"
	"sync"
	"time"
)

type Writer struct {
	out *bufio.Writer
	mu  sync.Mutex
}

func NewWriter(out *bufio.Writer) *Writer {
	writer := &Writer{
		out: out,
		mu:  sync.Mutex{},
	}
	go writer.daemon()

	return writer
}

func (w *Writer) Write(data []byte) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.out.Write(data)
}

func (w *Writer) daemon() {
	for range time.NewTicker(time.Second * 5).C {
		w.flush()
	}
}

func (w *Writer) flush() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.out == nil {
		return nil
	}

	return w.out.Flush()
}
