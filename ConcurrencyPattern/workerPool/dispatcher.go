package main

import "time"

// Dispatcher launches workers in parallel
//   and handle all the possible incoming channels.
type Dispatcher interface {
	// Injected WorkerLauncher type in this interface.
	// Dispatcher must use LauncherWorker method of any of the WorkerLaucher type
	//   to "initialize a pipeline"
	// In this way we can reuse Dispatcher interface to launch many types of WorkerLaunchers.
	LaunchWorker(id int, w WorkerLauncher)
	// Dispatcher interface expose this MakeReqeust
	//   to inject a new Requst into the worker pool.
	MakeRequest(Request)
	// Finally, user must call Stop() when all Goroutines must be finished to avoid Goroutine leaks.
	Stop()
}

// Dispatcher strcut stores a input Channel that takes the Request type.
// This is going to be "the single point" of entrance for requests in any pipelines.
type dispatcher struct {
	inCh chan Request
}

// LaunchWorker is to initialize Worker's LaunchWorker by taking dispatcher's channel which is deal with Request type.
//   need at least an ID and a channel for incoming Requests.
func (d *dispatcher) LaunchWorker(id int, w WorkerLauncher) {
	// In here, we can save running worker ID
	//   to control which ones are up or down.
	// The idea here (create a Dispatcher interface and using the "LauchWorker" to hide the implementation)
	//   is to hide launching implementation detail for users.
	//   which is "Facade Pattern".
	w.LaunchWorker(d.inCh)
}

// Stop is to close the incoming Requsts channel.
// When closing the incoming channel,
//   each for-range loop within the Goroutine breaks
//   and the Goroutine is also finished.
// In this case, when closing a shared channel,
//   it will provoke the same reaction, but in every listening Goroutine,
//   so all pipeline will be stopped.
func (d *dispatcher) Stop() {
	close(d.inCh)
}

// MakeRequest is just pass the Request to the channel which is transfer the incoming Request.
func (d *dispatcher) MakeRequest(r Request) {
	select {
	// It will select Goroutine
	//   until someone takes the request in the opposite side of the channel.
	case d.inCh <- r:
	// This is a receiving condition
	//   which will be trigger after 5s if there are no one take the Request.
	case <-time.After(time.Second * 5):
		// We can return an error here.
		return
	}
}

// NewDispatcher method handles the creation of dispatcher
//   instead of create it manully so that we can avoid some mistakes.
// Dispatcher will transfer the request to workers(pipeline).
func NewDispatcher(b int) Dispatcher {
	// Using "Singleton Pattern" to return the same instance
	return &dispatcher{
		// b refers to the buffer size in the channel.
		inCh: make(chan Request, b),
	}
}
