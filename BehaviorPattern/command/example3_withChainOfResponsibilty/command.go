package main

import (
	"fmt"
	"time"
)

// Command is an interface that has an action called Info
type Command interface {
	Info() string
}

// ChainLogger is also an interface that has an action called Next
// Next function will be implemented by the struct
type ChainLogger interface {
	Next(Command)
}

// TimePassed is a struct stored start time
type TimePassed struct {
	start time.Time
}

// Info is implemented by TimePassed struct and return the total passed time since the struct was been created.
func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

// Logger is a struct that store an variable called NextChain which type is an interface called ChainLogger.
type Logger struct {
	NextChain ChainLogger
}

// Next will check whether the NextChain of the struct Logger is nil or not.
// If the Logger is not nil, then pass the Command to the next logger.
func (f *Logger) Next(c Command) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}

func main() {
	// create an new instance of Logger
	second := new(Logger)
	// then append the second Logger to the Next function of the first Logger.
	first := Logger{NextChain: second}
	// Create an command (struct) called TimePassed which has implementd the Info() in Command interface.
	command := &TimePassed{start: time.Now()}

	first.Next(command)
}
