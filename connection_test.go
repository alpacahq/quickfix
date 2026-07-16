package quickfix

import (
	"errors"
	"io"
	"sync"
	"testing"
	"time"
)

type stubLog struct {
	mu     sync.Mutex
	events []string
}

func (l *stubLog) OnIncoming([]byte)           {}
func (l *stubLog) OnOutgoing([]byte)           {}
func (l *stubLog) OnEventf(string, ...interface{}) {}
func (l *stubLog) OnEvent(s string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.events = append(l.events, s)
}

func (l *stubLog) eventCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.events)
}

type failWriter struct {
	mu     sync.Mutex
	writes int
	closed bool
}

func (w *failWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.writes++
	return 0, errors.New("write tcp: broken pipe")
}

func (w *failWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.closed = true
	return nil
}

func (w *failWriter) writeCount() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.writes
}

func (w *failWriter) wasClosed() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.closed
}

var _ io.WriteCloser = (*failWriter)(nil)

func TestWriteLoopStopsWritingAfterBrokenPipe(t *testing.T) {
	msgOut := make(chan []byte)
	log := &stubLog{}
	w := &failWriter{}

	done := make(chan struct{})
	go func() {
		writeLoop(w, msgOut, log)
		close(done)
	}()

	// Flood the write loop with messages after the socket is dead.
	for i := 0; i < 50; i++ {
		msgOut <- []byte("8=FIX.4.2|35=8|")
	}
	close(msgOut)

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("writeLoop did not exit after messageOut was closed")
	}

	if got := w.writeCount(); got != 1 {
		t.Fatalf("expected exactly 1 Write attempt after broken pipe, got %d", got)
	}
	if got := log.eventCount(); got != 1 {
		t.Fatalf("expected exactly 1 broken-pipe log event, got %d", got)
	}
	if !w.wasClosed() {
		t.Fatal("expected connection to be closed after write failure")
	}
}
