package runtime

import (
	"github.com/ryxndmello/void/internal/context"
	"github.com/ryxndmello/void/pkg/command"
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
