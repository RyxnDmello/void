package void

func Execute(context *Context, command *Command) error {
	err := command.getPreExecute(context)

	if err != nil {
		return err
	}

	err = command.getExecute(context)

	if err != nil {
		return err
	}

	err = command.getPostExecute(context)

	if err != nil {
		return err
	}

	return nil
}
