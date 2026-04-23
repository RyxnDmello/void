package internal

import (
	"github.com/ryxndmello/void/command"
)

func Parse(commands map[string]*command.Command) (string, *command.Command, map[string]string, map[string]*command.Argument) {
	path, command, args := resolveCommand(commands)

	arguments := resolveArguments(command, args)

	return path, command, nil, arguments
}
