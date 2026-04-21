package void

import (
	"fmt"
)

type Void struct {
	commands map[string]*Command
}

func NewApp() *Void {
	return &Void{
		commands: make(map[string]*Command),
	}
}

func (void *Void) AddCommand(command *Command) *Void {
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

	path, command, _ := resolveCommand(void)

	return NewContext(path, command).Execute()
}
