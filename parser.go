package void

import (
	"os"
	"strings"

	"github.com/ryxndmello/void/command"
	"github.com/ryxndmello/void/errors"
)

func resolveCommand(commands map[string]*command.Command) (string, *command.Command, []string) {
	args := os.Args[1:]

	for i := len(args); i > 0; i-- {
		path := strings.Join(args[0:i], " ")

		if command := commands[path]; command != nil {
			return path, command, args[i:]
		}
	}

	panic(errors.CommandNotFound)
}
