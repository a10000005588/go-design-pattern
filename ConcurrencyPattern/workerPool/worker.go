package main

import (
	"fmt"
	"strings"
)

// WorkerLauncher handles the Request type.
// Any type implement this interface will have to receive a channel of Request type to satisfy it.
type WorkerLauncher interface {
	// [in] channel of Request type is the
	// single entrance point to the pipeline.
	LaunchWorker(in chan Request)
}

// PreffixSuffixWorker implement the WorkerLauncher
type PreffixSuffixWorker struct {
	id int
	// a string to prefix
	prefixS string
	// a string to suffix the incoming data of Request type.
	suffixS string
}

// Above section is about the steps of pipeline.
// Rememner, each step will accept a channel of incoming Request data,
//   and return a channel of same type.

// The first step in our pipeline is uppercase.
// uppercase will make the incoming word to be uppercase
// return a output Request channel.
func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	// create a channel which pass the Request struct.
	out := make(chan Request)

	go func() {
		// msg is Request type
		for msg := range in {

			s, ok := msg.Data.(string)
			if !ok {
				// stop this Request by sending nil
				//   if the casting is failed.
				msg.Handler(nil)
				continue
			}
			// be aware that the string is store in the msg(Request.Data).
			//   and will be sent as interface{}
			//   that means we need to cast it in next step.
			msg.Data = strings.ToUpper(s)

			out <- msg
		}

		close(out)
	}()

	return out
}

// second step of pipeline.
func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

// The last step in pipeline.
// Notice that this step isn't return a channel.
//   because we have used "Future Pattern: a future handler function"
//   (meaning execute the handler in the future)
//   to execute the final result in the pipeline with msg.Handler()
func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercaseStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			// Execute the final result by triggering the Handler in Request.
			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, uppercaseStringWithSuffix))
		}
	}()
}

// LaunchWorker is the worker that takes the request from the dispatcher.
// If each step in pipeline has done.
//   then LauchWorker deal with the request "from Request channel"
//   and send it to the pipeline.
func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}
