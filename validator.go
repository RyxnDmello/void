package void

import (
	"fmt"
	"slices"
	"strings"
)

func mustValidateCommands(commands []*Command) {
	names := make([]string, 0)

	for i := range commands {
		commands[i].mustValidate()

		if slices.Contains(names, commands[i].name) {
			panic(
				NewDefinitionError(DuplicateCommandPath, nil).
					WithMessage(
						"Duplicate command paths detected.",
					).
					WithExpected(
						"Each command must resolve to a unique command path.",
					).
					WithReceived(
						fmt.Sprintf("Duplicate command path: %s", strings.Join(commands[i].path, " ")),
					).
					WithSuggestion(
						"Ensure no two commands resolve to the same command path.",
					),
			)
		}

		mustValidateFlags(commands[i])
		mustValidateArguments(commands[i])

		names = append(names, commands[i].name)
	}
}

func mustValidateArguments(command *Command) {
	names := make([]string, 0)

	for i := range command.arguments {
		command.arguments[i].mustValidate()

		if slices.Contains(names, command.arguments[i].name) {
			panic(
				NewDefinitionError(DuplicateCommandArgument, nil).
					WithCommand(command).
					WithArgument(command.arguments[i]).
					WithMessage(
						"Duplicate command arguments detected.",
					).
					WithExpected(
						"Each command argument must have a unique name.",
					).
					WithReceived(
						fmt.Sprintf("Duplicate argument: %s", command.arguments[i].name),
					).
					WithSuggestion(
						"Rename or remove duplicate command arguments.",
					),
			)
		}

		names = append(names, command.arguments[i].name)
	}
}

func mustValidateFlags(command *Command) {
	names := make([]string, 0)

	for i := range command.flags {
		command.flags[i].mustValidate()

		if slices.Contains(names, command.flags[i].name) {
			panic(
				NewDefinitionError(DuplicateCommandFlag, nil).
					WithCommand(command).
					WithFlag(command.flags[i]).
					WithMessage(
						"Duplicate command flags detected.",
					).
					WithExpected(
						"Each command flag must have a unique name.",
					).
					WithReceived(
						fmt.Sprintf("Duplicate flag: %s", command.arguments[i].name),
					).
					WithSuggestion(
						"Rename or remove duplicate command flag.",
					),
			)
		}

		names = append(names, command.flags[i].name)
	}
}
