package void

import (
	"fmt"

	"github.com/ryxndmello/void/command"
	"github.com/ryxndmello/void/internal"
)

type Void struct {
	commands map[string]*command.Command
}

func NewApp() *Void {
	return &Void{
		commands: make(map[string]*command.Command),
	}
}

func (void *Void) Commands() map[string]*command.Command {
	return void.commands
}

func (void *Void) AddCommand(command *command.Command) *Void {
	path := command.Path()

	exists := void.commands[path]

	if exists != nil {
		panic("Command With Path Already Exists")
	}

	void.commands[path] = command

	return void
}

func (void *Void) Run() error {
	defer func() {
		if recover := recover(); recover != nil {
			fmt.Println(recover)
		}
	}()

	path, command, flags, arguments := internal.Parse(void.commands)

	return internal.NewContext(path, command, flags, arguments).Execute()
}
