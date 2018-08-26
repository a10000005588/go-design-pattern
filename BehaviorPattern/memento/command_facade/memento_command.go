package main

import "fmt"

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

//--------------------------------------------------------------------
// Command pattern will contain a specific function.Command
// 透過Command pattern 可以將“執行命令”封裝起來，不單單只是state
type Command interface {
	GetValue() interface{}
}

//--------------------------------------------------------------------
// Memento will store the type of Command struct
type Memento struct {
	memento Command
}

//--------------------------------------------------------------------
// Originator responsible for creating the Command struct.
// And dispatch the Command to the memento struct.
type originator struct {
	Command Command
}

// NewMemento Dispatch the Command inside the originator to the memento
func (o *originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

// Originator extract the Command in memento and
// store it to the Command in originators.
func (o *originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

//--------------------------------------------------------------------
// careTake store many Memento (each memento store a Command)
type careTaker struct {
	mementoStack []Memento
}

/* only the careTaker struct can access the Push and Pop method */
func (c *careTaker) Push(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *careTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		memento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[0 : len(c.mementoStack)-1]
		return memento
	}

	return Memento{}
}

//--------------------------------------------------------------------
// Facade patten is responsible for wrap the complex logic
// into a simple interface.
type MementoFacade struct {
	originator originator
	careTaker  careTaker
}

// via MementoFacade, we can save a Command in careTaker
func (m *MementoFacade) SaveSettings(s Command) {
	// originator record the newest Command.
	m.originator.Command = s
	m.careTaker.Push(m.originator.NewMemento())
}

// with RestoreSettings, we can extract the newset Command
// inside the memento array of careTaker.
func (m *MementoFacade) RestoreSettings() Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Pop())
	return m.originator.Command
}

//--------------------------------------------------------------------

func main() {
	m := MementoFacade{}

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	assertAndPrint(m.RestoreSettings())
	assertAndPrint(m.RestoreSettings())
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
