package void

import (
	"fmt"

	"github.com/ryxndmello/void/internal/context"
	"github.com/ryxndmello/void/internal/runtime"
	"github.com/ryxndmello/void/pkg/command"
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

func (void *Void) SetCommands(commands ...*command.Command) *Void {
	for i := range commands {
		path := commands[i].Path()

		if command := void.commands[path]; command != nil {
			panic(fmt.Sprintf("Command With Path '%s' Already Exists", path))
		}

		commands[i].MustValidate()

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

	path, command, arguments, flags := runtime.Parse(void.commands)

	context := context.NewContext(path, arguments, flags)

	return runtime.Execute(context, command)
}
