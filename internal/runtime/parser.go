package runtime

import (
	"github.com/ryxndmello/void/command"
)

func Parse(commands map[string]*command.Command) (string, *command.Command, map[string]any, map[string]any) {
	path, command, argParams, flagParams := resolveCommand(commands)

	flags := resolveFlags(command, flagParams)
	arguments := resolveArguments(command, argParams)

	return path, command, arguments, flags
}
