package conobserver

import (
	"fmt"
)

type Publisher interface {
	start()
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop() chan<- interface{}
}

type publisher struct {
	// Subscirber[] is empty
	// addSubCh, removeSubCh, in, stop are all <nil> channel

	subscribers []Subscriber
	// a channel to communicate with when we want to add a subscriber.
	addSubCh chan Subscriber
	// addSubCh chan Subscriber
	// a channel to ....when we want to remove a subscriber.
	removeSubCh chan Subscriber
	// [in] channel handle the incoming msg that must broadcast to all subscribers.
	in chan interface{}
	// [stop] channel is to kill all Goroutines.
	stop chan interface{}
}

// NewPublisher will new a publisher struct has implemented all the method in Publisher interface.
// 為各個channel 預設1大小的buffer channel，
//   避免start各個channel接收時，在傳送給各個channel造成deadlock住的狀況發生
func NewPublisher() Publisher {
	p := publisher{}
	p.addSubCh = make(chan Subscriber, 1)
	p.removeSubCh = make(chan Subscriber, 1)
	p.in = make(chan interface{}, 1)
	p.stop = make(chan interface{}, 1)
	// will return a initialized publisher struct.

	return &p
}

func (p *publisher) AddSubscriberCh() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return p.removeSubCh
}

func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}

func (p *publisher) Stop() chan<- interface{} {
	return p.stop
}

// The reason why we don't let user access each channel in publisher instead of using "proxy function"
//   is that the user needn't to deal with complexity of concurrent structure associated with our library
//   which they all be implement in proxy function.

// so...user can focus on their bussiness logic while maximizing performance

func (p *publisher) start() {
	for {
		// only one of a condition will be selected.
		select {
		// receives a message to publish to subscribers.
		case msg := <-p.in:
			fmt.Println("broadcast the msg to all subscriber")
			for _, ch := range p.subscribers {
				// if using go here, will occur race condition.
				// avoid the channel being close.

				// Notify each subscriber with the msg.
				// Notify will trigger the mockSubscirber's notifyTestingFunc(msg)
				// 將訊息塞給每一位subscriber
				// 透過subscriber寄放在publisher的notify電話筒
				// 在電話筒內傳msg給subscriber
				ch.Notify(msg)
			}
		// when a value arrives to the channel to add channels.
		case sub := <-p.addSubCh:
			fmt.Println("add a subscriber to channel")
			p.subscribers = append(p.subscribers, sub)
		// when a value arrives at the remove channel,
		// we need to search the subscriber in slice.
		case sub := <-p.removeSubCh:
			fmt.Println("remove a subscriber")

			for i, candidate := range p.subscribers {
				// find the subscriber interface
				if candidate == sub {
					// then remove it.
					p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		// stop all the goroutines.
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)
			close(p.stop)

			fmt.Println("p.stop is executed")
			return
		}
	}
}
