package runtime

import (
	"fmt"
	"os"
	"strings"

	"github.com/ryxndmello/void/pkg/command"
	"github.com/ryxndmello/void/pkg/errors"
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
	paramFlags := make(map[string][]string)

	index := 0

	for index < len(params) {
		name := params[index]

		if !strings.HasPrefix(name, "-") {
			panic(fmt.Sprintf("%s Not A Valid Flag", name))
		}

		parted := strings.Split(name, "=")

		if len(parted) > 2 {
			panic(fmt.Sprintf("%s Not A Valid Flag", name))
		}

		if len(parted) == 2 {
			name = parted[0]
			values := strings.Split(parted[1], ",")
			paramFlags[name] = append(paramFlags[name], values...)

			index++
			continue
		}

		if index+1 < len(params) && !strings.HasPrefix(params[index+1], "-") {
			values := strings.Split(params[index+1], ",")
			paramFlags[name] = append(paramFlags[name], values...)
			index += 2
			continue
		}

		paramFlags[name] = append(paramFlags[name], "true")
		index++
	}

	required, optional := cmd.Flags()

	if len(paramFlags) < len(required) {
		names := make([]string, 0)

		for i := range required {
			names = append(names, required[i].Name())
		}

		panic(fmt.Sprintf("Command '%s' Requires Flags: %v", cmd.Name(), strings.Join(names, ", ")))
	}

	if len(paramFlags) > len(required)+len(optional) {
		panic("Too Many Flags Provided")
	}

	flags := make(map[string]any)

	index = 0

	for i := range required {
		flag := required[i]

		paramFlag := paramFlags[flag.Name()]

		if len(paramFlag) == 0 {
			panic(fmt.Sprintf("Command '%s' Requires Flag: %s", cmd.Name(), flag.Name()))
		}

		if flag.IsVariadic() {
			flag.SetValue(paramFlag)
			flags[flag.Name()] = flag.Value()
			continue
		}

		if len(paramFlag) > 1 {
			panic(fmt.Sprintf("Command '%s' With Flag: %s Only Accepts A Single Value", cmd.Name(), flag.Name()))
		}

		flag.SetValue(paramFlag[0])
		flags[flag.Name()] = flag.Value()
	}

	for i := range optional {
		flag := optional[i]

		paramFlag := paramFlags[flag.Name()]

		if len(paramFlag) == 0 {
			if flag.Value() != nil {
				flags[flag.Name()] = flag.Value()
			}

			continue
		}

		if flag.IsVariadic() {
			flag.SetValue(paramFlag)
			flags[flag.Name()] = flag.Value()
			continue
		}

		if len(paramFlag) > 1 {
			panic(fmt.Sprintf("Command '%s' With Flag: %s Only Accepts A Single Value", cmd.Name(), flag.Name()))
		}

		flag.SetValue(paramFlag[0])
		flags[flag.Name()] = flag.Value()
	}

	return flags
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
