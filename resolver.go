package void

import (
	"fmt"
	"os"
	"strings"
)

func mustResolveCommand(commands []*Command) (*Command, []string, []string) {
	pathToCommand := mustValidateCommands(commands)

	params := os.Args[1:]

	if len(params) == 0 {
		panic(
			NewUsageError(CommandRequired, nil).
				WithMessage(
					"No command was provided.",
				).
				WithExpected(
					"A command path must be provided for execution.",
				).
				WithSuggestion(
					"Run the application with a valid command.",
				),
		)
	}

	for i := len(params); i > 0; i-- {
		path := strings.Join(params[0:i], " ")

		if command := pathToCommand[path]; command != nil {
			flagParams, argumentParams := mustPartitionCommandParams(params[i:])
			return command, flagParams, argumentParams
		}
	}

	panic(
		NewUsageError(UnknownCommand, nil).
			WithMessage(
				"Unknown command.",
			).
			WithReceived(
				fmt.Sprintf("Command path: %s", strings.Join(params, " ")),
			).
			WithSuggestion(
				"Run the application with a valid command.",
			),
	)
}

func mustResolveArguments(command *Command, params []string) map[string]any {
	required, optional, variadic := command.getArguments()

	if len(params) < len(required) {
		var requiredArguments strings.Builder

		for i := range required {
			fmt.Fprintf(&requiredArguments, "\n\t• %s", required[i].name)
		}

		panic(
			NewUsageError(MissingRequiredCommandArguments, nil).
				WithCommand(command).
				WithMessage(
					"Required arguments are missing.",
				).
				WithExpected(
					fmt.Sprintf("Required arguments:%s", requiredArguments.String()),
				).
				WithReceived(
					fmt.Sprintf("Received %d argument(s)", len(params)),
				).
				WithSuggestion(
					"Provide all required arguments.",
				),
		)
	}

	if variadic == nil && len(params) > len(required)+len(optional) {
		panic(
			NewUsageError(TooManyCommandArguments, nil).
				WithCommand(command).
				WithMessage(
					"Too many arguments were provided.",
				).
				WithExpected(
					fmt.Sprintf("Expected at most %d argument(s).", len(required)+len(optional)),
				).
				WithReceived(
					fmt.Sprintf("Received %d argument(s).", len(params)),
				).
				WithSuggestion(
					"Remove unsupported arguments from the command input.",
				),
		)
	}

	index := 0

	arguments := make(map[string]any)

	for i := range required {
		arguments[required[i].name] = required[i].setValue(params[index]).getValue()
		index++
	}

	for i := range optional {
		if index < len(params) {
			optional[i].setValue(params[index])
			index++
		}

		if value := optional[i].getValue(); value != nil {
			arguments[optional[i].name] = value
		}
	}

	if variadic != nil {
		if index < len(params) {
			arguments[variadic.name] = variadic.setValue(params[index:]).getValue()
			return arguments
		}

		if len(optional) == 0 {
			if variadic.required {
				panic(
					NewUsageError(MissingRequiredCommandArguments, nil).
						WithCommand(command).
						WithMessage(
							"Required arguments are missing.",
						).
						WithExpected(
							fmt.Sprintf("Required argument: %s", variadic.name),
						).
						WithSuggestion(
							"Provide all required command arguments.",
						),
				)
			}

			if value := variadic.getValue(); value != nil {
				arguments[variadic.name] = value
				return arguments
			}
		}
	}

	return arguments
}

func mustResolveFlags(command *Command, params []string) map[string]any {
	flagToValues := mustGroupFlagParams(params)

	required, optional := command.getFlags()

	if len(flagToValues) < len(required) {
		var flags strings.Builder

		for i := range required {
			fmt.Fprintf(&flags, "\n\t• %s", required[i].name)
		}

		panic(
			NewUsageError(MissingRequiredCommandFlags, nil).
				WithCommand(command).
				WithMessage(
					"Required flags are missing.",
				).
				WithExpected(
					fmt.Sprintf("Required flags:%s", flags.String()),
				).
				WithReceived(
					fmt.Sprintf("Received %d flag(s).", len(flagToValues)),
				).
				WithSuggestion(
					"Provide all required command flags.",
				),
		)
	}

	if len(flagToValues) > len(required)+len(optional) {
		panic(
			NewUsageError(TooManyCommandFlags, nil).
				WithCommand(command).
				WithMessage(
					"Too many flags were provided.",
				).
				WithExpected(
					fmt.Sprintf("Expected at most %d flag(s).", len(required)+len(optional)),
				).
				WithReceived(
					fmt.Sprintf("Received %d flag(s).", len(flagToValues)),
				).
				WithSuggestion(
					"Remove unsupported flags from the command input.",
				),
		)
	}

	flags := make(map[string]any)

	for i := range required {
		flag := required[i]

		values := flagToValues[flag.name]

		if len(values) == 0 {
			panic(
				NewUsageError(MissingRequiredCommandFlags, nil).
					WithFlag(flag).
					WithCommand(command).
					WithMessage(
						"A required flag was not provided.",
					).
					WithExpected(
						fmt.Sprintf("Required flag: %s", flag.name),
					).
					WithSuggestion(
						"Provide all required command flags.",
					),
			)
		}

		if flag.variadic {
			flags[flag.name] = flag.setValue(values).getValue()
			continue
		}

		if len(values) > 1 {
			panic(
				NewUsageError(TooManyCommandFlagValues, nil).
					WithFlag(flag).
					WithCommand(command).
					WithMessage(
						"Too many values were provided for the flag.",
					).
					WithExpected(
						"Expected a single flag value.",
					).
					WithReceived(
						fmt.Sprintf("Received %d values for flag '%s'.", len(values), flag.name),
					).
					WithSuggestion(
						"Provide only a single value for the flag.",
					),
			)
		}

		flags[flag.name] = flag.setValue(values[0]).getValue()
	}

	for i := range optional {
		flag := optional[i]

		values := flagToValues[flag.name]

		if len(values) == 0 {
			if value := flag.getValue(); value != nil {
				flags[flag.name] = value
			}

			continue
		}

		if flag.variadic {
			flags[flag.name] = flag.setValue(values).getValue()
			continue
		}

		if len(values) > 1 {
			panic(
				NewUsageError(TooManyCommandFlagValues, nil).
					WithFlag(flag).
					WithCommand(command).
					WithMessage(
						"Too many values were provided for the flag.",
					).
					WithExpected(
						"Expected a single flag value.",
					).
					WithReceived(
						fmt.Sprintf("Received %d values for flag '%s'.", len(values), flag.name),
					).
					WithSuggestion(
						"Provide only a single value for the flag.",
					),
			)
		}

		flags[flag.name] = flag.setValue(values[0]).getValue()
	}

	return flags
}

func mustValidateCommands(commands []*Command) map[string]*Command {
	pathToCommand := make(map[string]*Command)

	for i := range commands {
		path := strings.Join(commands[i].path, " ")

		if pathToCommand[path] != nil {
			panic(
				NewDefinitionError(DuplicateCommandRoute, nil).
					WithMessage(
						"Duplicate command paths detected.",
					).
					WithExpected(
						"Each command must resolve to a unique command path.",
					).
					WithReceived(
						fmt.Sprintf("Duplicate command path: %s", path),
					).
					WithSuggestion(
						"Ensure no two commands resolve to the same command path.",
					),
			)
		}

		commands[i].mustValidate()

		pathToCommand[path] = commands[i]
	}

	return pathToCommand
}

func mustPartitionCommandParams(params []string) ([]string, []string) {
	for i := range params {
		if strings.HasPrefix(params[i], "-") {
			return params[i:], params[0:i]
		}
	}

	return []string{}, params
}

func mustGroupFlagParams(params []string) map[string][]string {
	index := 0

	flagToValues := make(map[string][]string)

	for index < len(params) {
		name := params[index]

		if !strings.HasPrefix(name, "-") {
			panic(
				NewUsageError(InvalidFlag, nil).
					WithMessage(
						"An invalid flag was provided.",
					).
					WithReceived(
						fmt.Sprintf("Invalid flag: %s", name),
					).
					WithSuggestion(
						"Expected flag formats:\n\t• --flag=value\n\t• --flag=value1,value2,value3\n\t• --flag value1,value2,value3\n\t• --flag value1 --flag value2 --flag value3",
					),
			)
		}

		if parted := strings.Split(name, "="); len(parted) > 2 {
			panic(
				NewUsageError(InvalidFlagFormat, nil).
					WithMessage(
						"Invalid flag format detected.",
					).
					WithReceived(
						fmt.Sprintf("Invalid flag format: %s", name),
					).
					WithSuggestion(
						"Expected flag formats:\n\t• --flag=value\n\t• --flag=value1,value2,value3\n\t• --flag value1,value2,value3\n\t• --flag value1 --flag value2 --flag value3",
					),
			)
		} else if len(parted) == 2 {
			name = parted[0]
			values := strings.Split(parted[1], ",")
			flagToValues[name] = append(flagToValues[name], values...)
			index++
			continue
		}

		if index+1 < len(params) && !strings.HasPrefix(params[index+1], "-") {
			values := strings.Split(params[index+1], ",")
			flagToValues[name] = append(flagToValues[name], values...)
			index += 2
			continue
		}

		flagToValues[name] = append(flagToValues[name], name)
		index++
	}

	return flagToValues
}
