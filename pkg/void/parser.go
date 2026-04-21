package void

import (
	"os"
	"strings"
)

func resolveCommand(void *Void) (string, *Command, []string) {
	args := os.Args[1:]

	commands := void.commands

	for i := len(args); i > 0; i-- {
		path := strings.Join(args[0:i], " ")

		if command := commands[path]; command != nil {
			return path, command, args[i:]
		}
	}

	panic(CommandNotFound)
}
