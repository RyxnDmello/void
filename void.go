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

func (void *Void) SetCommands(commands ...*Command) *Void {
	for i := range commands {
		path := commands[i].getPath()

		if void.commands[path] != nil {
			panic(fmt.Sprintf("Command With Path '%s' Already Exists", path))
		}

		commands[i].mustValidate()

		void.commands[path] = commands[i]
	}

	return void
}

func (void *Void) Run() error {
	defer func() {
		if recover := recover(); recover != nil {
			fmt.Println(recover)
		}
	}()

	path, command, arguments, flags := parse(void.commands)

	context := NewContext(path, arguments, flags)

	return Execute(context, command)
}
