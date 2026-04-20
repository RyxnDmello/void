package void

import (
	"fmt"
	"os"
	"strings"
)

type Void struct {
	commands map[string]*Command
}

func NewApp() *Void {
	return &Void{
		commands: make(map[string]*Command),
	}
}

func (tree *Void) AddCommand(command *Command) *Void {
	path := command.Path()

	exists := tree.commands[path]

	if exists != nil {
		panic("Command With Path Already Exists")
	}

	tree.commands[path] = command

	return tree
}

func (tree *Void) Run() error {
	defer func() {
		if recover := recover(); recover != nil {
			fmt.Println(recover)
		}
	}()

	args := os.Args[1:]

	var path strings.Builder

	for _, arg := range args {
		path.WriteString(arg)

		command := tree.commands[path.String()]

		if command != nil {
			return NewContext(path.String(), command).Execute()
		}
	}

	return CommandNotFound
}
