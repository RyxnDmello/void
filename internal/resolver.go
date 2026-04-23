package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/ryxndmello/void/command"
	"github.com/ryxndmello/void/errors"
)

func resolveCommand(commands map[string]*command.Command) (string, *command.Command, []string) {
	parameters := os.Args[1:]

	for i := len(parameters); i > 0; i-- {
		path := strings.Join(parameters[0:i], " ")

		if command := commands[path]; command != nil {
			return path, command, parameters[i:]
		}
	}

	panic(errors.CommandNotFound)
}

func resolveArguments(cmd *command.Command, parameters []string) map[string]*command.Argument {
	required, optional := cmd.Arguments()

	// for _, r := range required {
	// 	fmt.Println(r.Name())
	// }

	if len(parameters) < len(required) {
		names := make([]string, 0)

		for i := range required {
			names = append(names, required[i].Name())
		}

		panic(fmt.Sprintf("Command '%s' Requires Arguments: %v", cmd.Name(), names))
	}

	if len(parameters) > len(required)+len(optional) {
		panic("Too Many Arguments Provided")
	}

	arguments := make(map[string]*command.Argument)

	index := 0

	for i := range required {
		required[i].SetValue(parameters[index])
		arguments[required[i].Name()] = required[i]
		index++
	}

	for i := range optional {
		if index < len(parameters) {
			optional[i].SetValue(parameters[index])
			arguments[optional[i].Name()] = optional[i]
			index++
			continue
		}

		if optional[i].Value() != nil {
			arguments[optional[i].Name()] = optional[i]
		}
	}

	// for i := range arguments {
	// 	fmt.Println(arguments[i].Name(), "=", arguments[i].Value())
	// }

	// fmt.Println()

	return arguments
}
