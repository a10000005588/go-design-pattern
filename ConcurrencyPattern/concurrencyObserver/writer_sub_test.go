package conobserver

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestStdoutPrinter(t *testing.T) {}

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func TestWriter(t *testing.T) {
	// it will be printed on the "io.Writer" interface
	msg := "Hello"
	// we should synchronize with the Subscribers
	//   in order to avoid race conditions on tests
	var wg sync.WaitGroup
	// one Notify() method will need to wait for one call to
	//   the Done() method
	wg.Add(1)
	// NewWriterSubscriber return a Subscriber interface
	sub := NewWriterSubscriber(0, nil)
	sub.Notify(msg)

	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		t.Fatal(err)
	}

	wg.Wait()
	sub.Close()
}
