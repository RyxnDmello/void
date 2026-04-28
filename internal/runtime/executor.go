package runtime

import (
	"github.com/ryxndmello/void/command"
	"github.com/ryxndmello/void/internal/context"
)

func Execute(context *context.Context, command *command.Command) error {
	err := command.PreExecute(context)

	if err != nil {
		return err
	}

	err = command.Execute(context)

	if err != nil {
		return err
	}

	err = command.PostExecute(context)

	if err != nil {
		return err
	}

	return nil
}
