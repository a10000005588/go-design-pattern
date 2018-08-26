package conobserver

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Subscriber interface {
	Notify(interface{}) error
	Close()
}

type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	// if no writer is specified, the default io.Writer interface is stdout
	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		id: id,
		// [in] is a chennel to accept the msg from publisher
		in:     make(chan interface{}),
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}

// if [in] channel is closed, the for range loop will stop
//   that particular behavior Goroutine will finish.
func (s *writerSubscriber) Close() {
	close(s.in)
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			// %#v gives us most of the information about any type when formatting to a string.
			// for example, a closed channel, it will return "send on a closed channel"
			err = fmt.Errorf("%#v", rec)
		}
	}()

	// when communicating with channel, there are two behavior we usually control:
	//  1. waiting time (for channel, how many time it can wait?)
	//  2. when the channel is closed ?
	select {
	case s.in <- msg:
		fmt.Printf("Subscriber get new message from publisher: %s\n", msg)
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout\n")
	}

	return

}
