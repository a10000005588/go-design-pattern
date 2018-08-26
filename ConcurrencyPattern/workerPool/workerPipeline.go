package main

import (
	"fmt"
	"log"
	"sync"
)

// Request holds Data and Handler
type Request struct {
	// WE can use Data to pass a string
	//   the string may be the (string, int or struct) type.
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

// NewStringRequest create Request along with the s string.
func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: s,
		// Handler id defined by using a closure
		//   closure mainly wrap the wg.Done so that
		//   if the goroutine goto the latest step,
		//   it can trigger the Handler to end itself
		Handler: func(i interface{}) {
			defer wg.Done()
			// We will cast the i and check the casting is ok or not.
			s, ok := i.(string)
			// If cast failed then print s.
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			// In here, we usually do something to the s.
			//   but we just print it out.
			fmt.Println(s)
		},
	}
	return myRequest
}

func main() {
	bufferSize := 100
	// Create a dispatcher which is responsible for
	// 1. LaunchWorker(w WorkerLaucher)
	// 2. MakeRequest(Request)
	// 3. Stop()

	// The dispatcher will launch as many instances of the pipeline as we want to
	//   route the incoming request to any available workers !
	// If none of the workers takes the request,
	//   the request is lost.
	var dispatcher = NewDispatcher(bufferSize)

	// We launch three workers of our defined pipeline.
	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{
			// prefixS will be the workerID and the suffiexS which is combined by worker through the pipeline.
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		// We launch the workers by calling the LaunchWorker in the dispatch interface.
		dispatcher.LaunchWorker(i, w)
		// A worker is a pipeline.
	}
	// Then, three workers are running concurrently
	//   and waiting for the request to tackle with.

	requests := 10
	// We should need a WaitGroup to properly synchronize the app.
	//   so that it doesn't "exit" too early !
	var wg sync.WaitGroup
	wg.Add(requests)

	// For 10 requests, we will need to wait for 10 wg.Dones()
	for i := 0; i < requests; i++ {
		// Create Request which is waiting for dispatched.
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}
