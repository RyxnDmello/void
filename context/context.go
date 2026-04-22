package context

import (
	"github.com/ryxndmello/void/command"
)

type Context struct {
	path      string
	command   *command.Command
	flags     []string
	arguments []string
}

func NewContext(path string, command *command.Command) *Context {
	if command == nil {
		panic("A Command Must Be Assigned To The Void Context")
	}

	return &Context{
		path:      path,
		command:   command,
		flags:     command.Flags(),
		arguments: command.Arguments(),
	}
}

func (ctx *Context) Execute() error {
	if preExecute := ctx.command.PreExecute(); preExecute != nil {
		err := preExecute(ctx.path, ctx.command, ctx.arguments, ctx.flags)

		if err != nil {
			return err
		}
	}

	if execute := ctx.command.Execute(); execute != nil {
		err := execute(ctx.path, ctx.command, ctx.arguments, ctx.flags)

		if err != nil {
			return err
		}
	}

	if postExecute := ctx.command.PostExecute(); postExecute != nil {
		err := postExecute(ctx.path, ctx.command, ctx.arguments, ctx.flags)

		if err != nil {
			return err
		}
	}

	return nil
}
