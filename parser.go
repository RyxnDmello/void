package void

func parse(commands map[string]*Command) (string, *Command, map[string]any, map[string]any) {
	path, command, argumentParams, flagParams := mustResolveCommand(commands)

	flags := mustResolveFlags(command, flagParams)
	arguments := mustResolveArguments(command, argumentParams)

	return path, command, arguments, flags
}
