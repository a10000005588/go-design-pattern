package memento

import "fmt"

// There are three type in Memento : 1.Memento, 2. State, 3. Originator

// State is the core business object
// It is any object we want to track.
type State struct {
	Description string
}

// contain a single value called state which has State type.
// our state will be containerized within this type before storing them
// into the Care Taker type.

// 這裡不直接宣告一個state變數，而是再透過memento strcut來包起來
// 目的是為了要originator 和 careTake 從business object中 抽離出來
type memento struct {
	state State
}

// 這裡有個tip 為何不要直接就存 originator的state就好？？
// 因為 memnto存的是特定的state.
// originator存的是當前載入的state
type originator struct {
	state State
}

// 產生一個mememto struct 存有著originator的state.
func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

//--------------------------------------------------------------------
// store the memnto list.
type careTaker struct {
	mementoList []memento
}

// add the memnto to its memntoList
func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	// 如果索引i大於當前mementoList容量 以及 < 0...
	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("Index not found\n")
	}
	return c.mementoList[i], nil
}
