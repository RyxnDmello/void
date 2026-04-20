package void

type Context struct {
	path      string
	command   *Command
	flags     []string
	arguments []string
}

func NewContext(path string, command *Command) *Context {
	if command == nil {
		panic("A Command Must Be Assigned To The Void Context")
	}

	return &Context{
		path:      path,
		command:   command,
		flags:     command.flags,
		arguments: command.arguments,
	}
}

func (context *Context) Execute() error {
	if context.command.preExecute != nil {
		err := context.command.preExecute(context, context.arguments, context.flags)

		if err != nil {
			return err
		}
	}

	if context.command.execute != nil {
		err := context.command.execute(context, context.arguments, context.flags)

		if err != nil {
			return err
		}
	}

	if context.command.postExecute != nil {
		err := context.command.postExecute(context, context.arguments, context.flags)

		if err != nil {
			return err
		}
	}

	return nil
}
