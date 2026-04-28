package runtime

import (
	"fmt"
	"os"
	"strings"

	"github.com/ryxndmello/void/command"
	"github.com/ryxndmello/void/errors"
)

func resolveCommand(cmds map[string]*command.Command) (string, *command.Command, []string, []string) {
	params := os.Args[1:]

	for i := len(params); i > 0; i-- {
		path := strings.Join(params[0:i], " ")

		if command := cmds[path]; command != nil {
			argParams, flagParams := sortParams(params[i:])
			return path, command, argParams, flagParams
		}
	}

	panic(errors.CommandNotFound)
}

func resolveArguments(cmd *command.Command, params []string) map[string]any {
	required, optional, variadic := cmd.Arguments()

	if len(params) < len(required) {
		names := make([]string, 0)

		for i := range required {
			names = append(names, required[i].Name())
		}

		panic(fmt.Sprintf("Command '%s' Requires Arguments: %v", cmd.Name(), strings.Join(names, ", ")))
	}

	if variadic == nil && len(params) > len(required)+len(optional) {
		panic("Too Many Arguments Provided")
	}

	arguments := make(map[string]any)

	index := 0

	for i := range required {
		required[i].SetValue(params[index])
		arguments[required[i].Name()] = required[i].Value()
		index++
	}

	for i := range optional {
		if index < len(params) {
			optional[i].SetValue(params[index])
			arguments[optional[i].Name()] = optional[i].Value()
			index++
			continue
		}

		if optional[i].Value() != nil {
			arguments[optional[i].Name()] = optional[i].Value()
		}
	}

	if variadic != nil {
		if index < len(params) {
			variadic.SetValue(params[index:])
			arguments[variadic.Name()] = variadic.Value()
			return arguments
		}

		if len(optional) == 0 {
			if variadic.IsRequired() {
				panic(fmt.Sprintf("Command '%s' Requires Arguments: %s", cmd.Name(), variadic.Name()))
			}

			if variadic.Value() != nil {
				arguments[variadic.Name()] = variadic.Value()
				return arguments
			}
		}
	}

	return arguments
}

func resolveFlags(cmd *command.Command, params []string) map[string]any {
	flags := cmd.Flags()

	if len(params) > 0 && len(flags) == 0 {
		panic("Too Many Flags Provided")
	}

	return make(map[string]any)
}

func sortParams(params []string) ([]string, []string) {
	var argParams, flagParams []string

	for i := range params {
		if strings.HasPrefix(params[i], "-") {
			flagParams = params[i:]
			break
		}

		argParams = append(argParams, params[i])
	}

	return argParams, flagParams
}
