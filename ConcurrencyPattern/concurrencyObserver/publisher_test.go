package conobserver

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// mockSubscriber must implement Subscriber interface.
// need have 1. Close 2. Notify
type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

// The main problem of TestPublisher is to deal with the race condition
//   when access the subscriber list.
func TestPublisher(t *testing.T) {
	// The message we used to send to the Subscribers.
	// Each subscriber use the channel returned by the "addSubscriberCh"
	//   must receive this message.
	msg := "Hello"

	p := NewPublisher()

	var wg sync.WaitGroup
	// Define the mockSubscriber to add it to the "publisher list" of known subscribers.
	// mockSubscriber has two func, need to define them.
	//   1. notifyTestingFunc func(msg interface{})
	//   2. closeTestingFunc  func()
	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			// In the end triggers a wg.Done()
			defer wg.Done()
			// msg is coming as an interface, need type assertion.
			s, ok := msg.(string)
			if !ok {
				t.Fatal(errors.New("Could not assert result"))
			}
			fmt.Printf("Subscriber get new message from publisher: %s\n", s)
			if s != msg {
				t.Fail()
			}

		},
		closeTestingFunc: func() {
			wg.Done()
		},
	}
	// start() contains select which also contain four condition :
	//  1. <-p.in
	//  2. <-p.addSubCh
	//  3. <-p.removeSubCh
	//  4. <-p.stop
	//  when we goroutine the 'p.start', it will be block
	//    until there is someone give value to one of the channel(condition) above.
	go p.start()

	/* 注意 若start還沒初始化好p.各種channel 以下會發生deadlock... */
	// 等待p.start() 初始化好select內的channel...
	time.Sleep(time.Duration(1) * time.Second)

	// NewPublisher 回傳的是 Publisher interface
	//   而不是 publisher struct ，所以無法直接存取 publisher strcut內的資訊
	//   只能使用publisher 實作Publisher interface的方法。。。
	// add a subscriber into the publisher's [addSubCh] channel.
	p.AddSubscriberCh() <- sub

	// 回傳了p.addSubCh 但。。。Select沒有順利執行到
	//   可能是start goroutine尚未啟動完畢，這邊需要synchronized一下
	// 丟入的是 *mockSubscriber型態，channel只接受 Subscriebr
	//   不過 Subscriber是interface
	//   有實作(Notify,Close)的struct 還是可以傳入只有接受Subscriber interface。。。

	// 只有一位subscriber 故只會等待一個p.PublishingCh()觸發 wg.Done()
	//   定義在notifyTestFunc內，且sub.Notify會呼叫notifyTestFunc()
	wg.Add(1)
	// then broadcasting a msg to the subscriber.
	p.PublishingCh() <- msg
	wg.Wait()

	// pubCon is a "concrete instance" of publisher
	pubCon := p.(*publisher)
	// The number of subscribers must be 1 after calling the "AddSubscriberCh" channel
	if len(pubCon.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}

	wg.Add(1)
	p.RemoveSubscriberCh() <- sub
	wg.Wait()
	// Number of subscriber is restored to zero
	if len(pubCon.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}

	p.Stop() <- nil
}
