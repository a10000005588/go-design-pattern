package observer

import "fmt"

// Observer interface define an observer must implement the Notify() so that they can receive the message from publisher.
type Observer interface {
	Notify(string)
}

// Publisher will maintain observer list.
type Publisher struct {
	ObserversList []Observer
}

func (s *Publisher) AddObserver(o Observer) {
	s.ObserversList = append(s.ObserversList, o)
}

func (s *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range s.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	s.ObserversList = append(s.ObserversList[:indexToRemove], s.ObserversList[indexToRemove+1:]...)
}

// NotifyObservers is the method that the publisher notify every observer.
func (s *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, observer := range s.ObserversList {
		observer.Notify(m)
	}
}
