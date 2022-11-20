package cpu

// The debugger needs some way to control the cpu and request information about it
// It'll do this through commands

type Command interface {
	Command() Command
}

type command struct{}

func (e *command) Command() Command { return e }

type InterruptCommand struct {
	command
	Vector uint16
}
