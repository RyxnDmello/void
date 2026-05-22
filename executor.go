package void

func execute(command *Command, context *Context) error {
	err := command.execute(context)

	if err != nil {
		return err
	}

	err = command.preExecute(context)

	if err != nil {
		return err
	}

	err = command.postExecute(context)

	if err != nil {
		return err
	}

	return nil
}
