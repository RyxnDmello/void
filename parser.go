package void

func parse(commands []*Command) (*Command, map[string]any, map[string]any) {
	command, flagParams, argumentParams := mustResolveCommand(commands)
	arguments := mustResolveArguments(command, argumentParams)
	flags := mustResolveFlags(command, flagParams)
	return command, arguments, flags
}
