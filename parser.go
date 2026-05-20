package void

func parse(pathToCommand map[string]*Command) (*Command, map[string]any, map[string]any) {
	command, flagParams, argumentParams := mustResolveCommand(pathToCommand)
	arguments := mustResolveArguments(command, argumentParams)
	flags := mustResolveFlags(command, flagParams)
	return command, arguments, flags
}
