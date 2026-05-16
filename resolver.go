package void

import (
	"fmt"
	"os"
	"strings"
)

func mustResolveCommand(cmds map[string]*Command) (string, *Command, []string, []string) {
	params := os.Args[1:]

	for i := len(params); i > 0; i-- {
		path := strings.Join(params[0:i], " ")

		if command := cmds[path]; command != nil {
			argParams, flagParams := sortParams(params[i:])
			return path, command, argParams, flagParams
		}
	}

	panic(CommandNotFound)
}

func mustResolveArguments(cmd *Command, params []string) map[string]any {
	required, optional, variadic := cmd.getArgs()

	if len(params) < len(required) {
		names := make([]string, 0)

		for i := range required {
			names = append(names, required[i].name)
		}

		panic(fmt.Sprintf("Command '%s' Requires Arguments: %v", cmd.name, strings.Join(names, ", ")))
	}

	if variadic == nil && len(params) > len(required)+len(optional) {
		panic("Too Many Arguments Provided")
	}

	arguments := make(map[string]any)

	index := 0

	for i := range required {
		required[i].setValue(params[index])
		arguments[required[i].name] = required[i].getValue()
		index++
	}

	for i := range optional {
		if index < len(params) {
			optional[i].setValue(params[index])
			arguments[optional[i].name] = optional[i].getValue()
			index++
			continue
		}

		if optional[i].getValue() != nil {
			arguments[optional[i].name] = optional[i].getValue()
		}
	}

	if variadic != nil {
		if index < len(params) {
			variadic.setValue(params[index:])
			arguments[variadic.name] = variadic.getValue()
			return arguments
		}

		if len(optional) == 0 {
			if variadic.required {
				panic(fmt.Sprintf("Command '%s' Requires Arguments: %s", cmd.name, variadic.name))
			}

			if variadic.getValue() != nil {
				arguments[variadic.name] = variadic.getValue()
				return arguments
			}
		}
	}

	return arguments
}

func mustResolveFlags(cmd *Command, params []string) map[string]any {
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
			array := strings.Split(params[index+1], ",")

			values := make([]string, 0)

			for i := range array {
				values = append(values, strings.Trim(array[i], "\""))
			}

			paramFlags[name] = append(paramFlags[name], values...)
			index += 2
			continue
		}

		paramFlags[name] = append(paramFlags[name], "true")
		index++
	}

	required, optional := cmd.getFlags()

	if len(paramFlags) < len(required) {
		names := make([]string, 0)

		for i := range required {
			names = append(names, required[i].name)
		}

		panic(fmt.Sprintf("Command '%s' Requires Flags: %v", cmd.name, strings.Join(names, ", ")))
	}

	if len(paramFlags) > len(required)+len(optional) {
		panic("Too Many Flags Provided")
	}

	flags := make(map[string]any)

	index = 0

	for i := range required {
		flag := required[i]

		paramFlag := paramFlags[flag.name]

		if len(paramFlag) == 0 {
			panic(fmt.Sprintf("Command '%s' Requires Flag: %s", cmd.name, flag.name))
		}

		if flag.variadic {
			flag.setValue(paramFlag)
			flags[flag.name] = flag.getValue()
			continue
		}

		if len(paramFlag) > 1 {
			panic(fmt.Sprintf("Command '%s' With Flag: %s Only Accepts A Single Value", cmd.name, flag.name))
		}

		flag.setValue(paramFlag[0])
		flags[flag.name] = flag.getValue()
	}

	for i := range optional {
		flag := optional[i]

		paramFlag := paramFlags[flag.name]

		if len(paramFlag) == 0 {
			if flag.getValue() != nil {
				flags[flag.name] = flag.getValue()
			}

			continue
		}

		if flag.variadic {
			flag.setValue(paramFlag)
			flags[flag.name] = flag.getValue()
			continue
		}

		if len(paramFlag) > 1 {
			panic(fmt.Sprintf("Command '%s' With Flag: %s Only Accepts A Single Value", cmd.name, flag.name))
		}

		flag.setValue(paramFlag[0])
		flags[flag.name] = flag.getValue()
	}

	return flags
}

func sortParams(params []string) ([]string, []string) {
	var argParams, flagParams []string

	for i := range params {
		if strings.HasPrefix(params[i], "-") {
			argParams = params[0:i]
			flagParams = params[i:]
			break
		}
	}

	return argParams, flagParams
}
