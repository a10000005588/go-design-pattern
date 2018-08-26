package main

import (
	"fmt"
	"net/http"
)

// Command is a interface the regulate the action of the command.
type Command interface {
	Execute()
}

// ConsoleOutput dealing with the message
type ConsoleOutput struct {
	message string
}

// Execute is the function that implemented by ConsoleOutput.
func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

// CreateCommands will return a type which has implemented Execute() in the Command interface.
func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

// CommandQueue will store a list of Command interface.
type CommandQueue struct {
	queue []Command
}

// AddCommand will deal with the incoming Command and append it to the CommandQueue.AddCommand
// if the number of command exceeds 3, then doing the action for that commmand.
func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}

		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))

	client := http.Client{}
	client.Do(nil)
}
