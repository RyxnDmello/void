package void

import (
	"fmt"
	"os"
	"strings"
)

func mustResolveCommand(pathToCommand map[string]*Command) (*Command, []string, []string) {
	params := os.Args[1:]

	if len(params) == 0 {
		panic(CommandNotFound)
	}

	for i := len(params); i > 0; i-- {
		path := strings.Join(params[0:i], " ")

		if command := pathToCommand[path]; command != nil {
			flagParams, argumentParams := mustSortCommandParams(params[i:])
			return command, flagParams, argumentParams
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
		panic(fmt.Sprintf("Too Many Arguments Provided For Command '%s'", cmd.name))
	}

	index := 0

	arguments := make(map[string]any)

	for i := range required {
		index++
		arguments[required[i].name] = required[i].setValue(params[index]).getValue()
	}

	for i := range optional {
		if index < len(params) {
			index++
			optional[i].setValue(params[i])
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
				panic(fmt.Sprintf("Command '%s' Requires Arguments: %s", cmd.name, variadic.name))
			}

			if value := variadic.getValue(); value != nil {
				arguments[variadic.name] = value
				return arguments
			}
		}
	}

	return arguments
}

func mustResolveFlags(cmd *Command, params []string) map[string]any {
	flagToValues := mustSortFlagParams(params)

	required, optional := cmd.getFlags()

	if len(flagToValues) < len(required) {
		names := make([]string, 0)

		for i := range required {
			names = append(names, required[i].name)
		}

		panic(fmt.Sprintf("Command '%s' Requires Flags: %v", cmd.name, strings.Join(names, ", ")))
	}

	if len(flagToValues) > len(required)+len(optional) {
		panic(fmt.Sprintf("Too Many Flags Provided For Command '%s'", cmd.name))
	}

	flags := make(map[string]any)

	for i := range required {
		flag := required[i]

		values := flagToValues[flag.name]

		if len(values) == 0 {
			panic(fmt.Sprintf("Command '%s' Requires Flag: %s", cmd.name, flag.name))
		}

		if flag.variadic {
			flags[flag.name] = flag.setValue(values).getValue()
			continue
		}

		if len(values) > 1 {
			panic(fmt.Sprintf("Command '%s' With Flag: %s Only Accepts A Single Value", cmd.name, flag.name))
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
			panic(fmt.Sprintf("Command '%s' With Flag: %s Only Accepts A Single Value", cmd.name, flag.name))
		}

		flags[flag.name] = flag.setValue(values[0]).getValue()
	}

	return flags
}

func mustSortCommandParams(params []string) ([]string, []string) {
	for i := range params {
		if strings.HasPrefix(params[i], "-") {
			return params[i:], params[0:i]
		}
	}

	return nil, nil
}

func mustSortFlagParams(params []string) map[string][]string {
	index := 0

	flagToValues := make(map[string][]string)

	for index < len(params) {
		name := params[index]

		if !strings.HasPrefix(name, "-") {
			panic(fmt.Sprintf("%s Not A Valid Flag", name))
		}

		if parted := strings.Split(name, "="); len(parted) > 2 {
			panic(fmt.Sprintf("%s Not A Valid Flag", name))
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
